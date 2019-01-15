/*
Copyright 2018 Ridecell, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package components

import (
	"context"
	"net/http"
	"reflect"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
	"sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

func NewReconciler(name string, mgr manager.Manager, top runtime.Object, templates http.FileSystem, components []Component) (*componentReconciler, error) {
	logger := log.Log.WithName(name)
	cr := &componentReconciler{
		name:       name,
		top:        top,
		templates:  templates,
		components: components,
		manager:    mgr,
	}

	// Create the controller.
	c, err := controller.New(name, mgr, controller.Options{Reconciler: cr})
	if err != nil {
		logger.Error(err, "unable to create controller")
		return nil, err
	}

	// Watch for changes in the Top object.
	err = c.Watch(&source.Kind{Type: cr.top}, &handler.EnqueueRequestForObject{})
	if err != nil {
		logger.Error(err, "unable to create top-level watch")
		return nil, err
	}

	// Watch for changes in owned objects requested by components.
	watchedTypes := map[reflect.Type]bool{}
	for _, comp := range cr.components {
		for _, watchObj := range comp.WatchTypes() {
			watchType := reflect.TypeOf(watchObj).Elem()
			_, ok := watchedTypes[watchType]
			if ok {
				// Already watching.
				continue
			}
			watchedTypes[watchType] = true

			err = c.Watch(&source.Kind{Type: watchObj}, &handler.EnqueueRequestForOwner{
				IsController: true,
				OwnerType:    cr.top,
			})
			if err != nil {
				logger.Error(err, "unable to create watch")
				return nil, err
			}
			gatherer, ok := comp.(GathererComponent)
			if !ok {
				continue
			}
			toRequests := func(object handler.MapObject) []reconcile.Request {
				// Pull the metav1.Object out of the runtime.Object
				metaObj, err := meta.Accessor(top)
				if err != nil {
					logger.Error(err, "unable to create watch")
				}
				return []reconcile.Request{
					{NamespacedName: types.NamespacedName{
						Namespace: metaObj.GetNamespace(),
						Name:      metaObj.GetName(),
					}},
				}
			}
			err = c.Watch(
				&source.Kind{Type: watchObj},
				&handler.EnqueueRequestsFromMapFunc{
					ToRequests: handler.ToRequestsFunc(toRequests),
				},
				gatherer.WatchPredicateFuncs(),
			)
			if err != nil {
				logger.Error(err, "unable to create watch")
				return nil, err
			}
		}
	}

	return cr, nil
}

func (cr *componentReconciler) newContext(request reconcile.Request) (*ComponentContext, error) {
	reqCtx := context.TODO()

	// Fetch the current value of the top object for this reconcile.
	top := cr.top.DeepCopyObject()
	err := cr.client.Get(reqCtx, request.NamespacedName, top)
	if err != nil {
		return nil, err
	}

	ctx := &ComponentContext{
		logger:    log.Log.WithName(cr.name).WithValues("request", request.NamespacedName),
		templates: cr.templates,
		Context:   reqCtx,
		Top:       top,
	}
	err = cr.manager.SetFields(ctx)
	if err != nil {
		ctx.logger.Error(err, "error calling manager.SetFields")
		return nil, err
	}
	return ctx, nil
}

func (cr *componentReconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {

	// Build a reconciler context to pass around.
	ctx, err := cr.newContext(request)
	if err != nil {
		if errors.IsNotFound(err) {
			// Top object not found, likely already deleted.
			return reconcile.Result{}, nil
		}
		// Some other fetch error, try again on the next tick.
		return reconcile.Result{Requeue: true}, err
	}
	ctx.logger.Info("reconciling")

	// Make a clean copy of the top object to diff against later. This is used for
	// diffing because the status subresource might not always be available.
	cleanTop := ctx.Top.DeepCopyObject()

	// Reconcile all the components.
	result, err := cr.reconcileComponents(ctx)
	if err != nil {
		ctx.logger.Error(err, "Reconciliation error")
		ctx.Top.(Statuser).SetErrorStatus(err.Error())
	}

	// Check if an update to the status subresource is required.
	if !reflect.DeepEqual(ctx.Top.(Statuser).GetStatus(), cleanTop.(Statuser).GetStatus()) {
		// Update the top object status.
		ctx.logger.V(1).Info("updating status")
		err = cr.client.Status().Update(ctx.Context, ctx.Top)
		if err != nil {
			// Something went wrong, we definitely want to rerun, unless ...
			oldRequeue := result.Requeue
			result.Requeue = true
			if errors.IsNotFound(err) {
				// Older Kubernetes which doesn't support status subobjects, so use a GET+UPDATE
				// because the controller-runtime client doesn't support PATCH calls.
				freshTop := cr.top.DeepCopyObject()
				err = cr.client.Get(ctx.Context, request.NamespacedName, freshTop)
				if err != nil {
					// What?
					return result, err
				}
				freshTop.(Statuser).SetStatus(ctx.Top.(Statuser).GetStatus())
				err = cr.client.Update(ctx.Context, freshTop)
				if err != nil {
					// Update failed, probably another update got there first.
					return result, err
				} else {
					// Update worked, so no error for the final return.
					result.Requeue = oldRequeue
					err = nil
				}
			}
		}
	}

	return result, err
}

func (cr *componentReconciler) reconcileComponents(ctx *ComponentContext) (reconcile.Result, error) {
	ready := []Component{}
	for _, component := range cr.components {
		ctx.logger.V(1).Info("checking if available to reconcile", "component", component)
		if component.IsReconcilable(ctx) {
			ctx.logger.V(1).Info("is available to reconcile", "component", component)
			ready = append(ready, component)
		}
	}
	res := reconcile.Result{}
	ctxLogger := ctx.logger
	for _, component := range ready {
		ctx.logger = ctxLogger.WithValues("component", component)
		innerRes, err := component.Reconcile(ctx)
		// Update result. This should be checked before the err!=nil because sometimes
		// we want to requeue immediately on error.
		if innerRes.Requeue {
			res.Requeue = true
		}
		if innerRes.RequeueAfter != 0 && (res.RequeueAfter == 0 || res.RequeueAfter > innerRes.RequeueAfter) {
			res.RequeueAfter = innerRes.RequeueAfter
		}
		if err != nil {
			return res, err
		}
	}
	ctx.logger = ctxLogger
	return res, nil
}

// componentReconciler implements inject.Client.
// A client will be automatically injected.
var _ inject.Client = &componentReconciler{}

// InjectClient injects the client.
func (v *componentReconciler) InjectClient(c client.Client) error {
	v.client = c
	return nil
}
