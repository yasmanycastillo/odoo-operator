/*
 * This file is part of the Odoo-Operator (R) project.
 * Copyright (c) 2018-2018 XOE Corp. SAS
 * Authors: David Arnold, et al.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *
 * ALTERNATIVE LICENCING OPTION
 *
 * You can be released from the requirements of the license by purchasing
 * a commercial license. Buying such a license is mandatory as soon as you
 * develop commercial activities involving the Odoo-Operator software without
 * disclosing the source code of your own applications. These activities
 * include: Offering paid services to a customer as an ASP, shipping Odoo-
 * Operator with a closed source product.
 *
 */

package odooinstance

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/xoe-labs/odoo-operator/pkg/finalizer"

	clusterv1beta1 "github.com/xoe-labs/odoo-operator/pkg/apis/cluster/v1beta1"
	instancev1beta1 "github.com/xoe-labs/odoo-operator/pkg/apis/instance/v1beta1"
	clustercontroller "github.com/xoe-labs/odoo-operator/pkg/controller/odoocluster/odoocluster_controller"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new OdooInstance Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
// USER ACTION REQUIRED: update cmd/manager/main.go to call this instance.Add(mgr) to install this Controller
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileOdooInstance{Client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("odooinstance-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to OdooInstance
	err = c.Watch(&source.Kind{Type: &instancev1beta1.OdooInstance{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create
	// Uncomment watch a Deployment created by OdooInstance - change this for objects you create
	// err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForOwner{
	// 	IsController: true,
	// 	OwnerType:    &instancev1beta1.OdooInstance{},
	// })
	// if err != nil {
	// 	return err
	// }

	return nil
}

var _ reconcile.Reconciler = &ReconcileOdooInstance{}

const (
	// FinalizerKey ...
	FinalizerKey = "cleanup.odooinstance.odoo.io"
)

// ReconcileOdooInstance reconciles a OdooInstance object
type ReconcileOdooInstance struct {
	client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a OdooInstance object and makes changes based on the state read
// and what is in the OdooInstance.Spec
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=instance.odoo.io,resources=odooinstances,verbs=get;list;watch;create;update;patch;delete
func (r *ReconcileOdooInstance) Reconcile(request reconcile.Request) (reconcile.Result, error) {

	// Fetch the OdooInstance instance
	instance := &instancev1beta1.OdooInstance{}
	err := r.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Fetch the OdooCluster instance, error if not found
	clusterinstance := &clusterv1beta1.OdooCluster{}
	err = r.Get(context.TODO(), request.NamespacedName, clusterinstance)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Printf("%s/%s Operation: %s. (Controller: OdooInstance) Error: No OdooCluster found in this namespace\n", instance.Namespace, instance.Name, "validate")
			return reconcile.Result{}, err
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}
	var parentInstance *instancev1beta1.OdooInstance
	// Recursively get parent instance, if set
	if instance.Spec.ParentName != nil {
		parentInstance := &instancev1beta1.OdooInstance{}
		err := r.Get(context.TODO(), request.NamespacedName, parentInstance)
		if err != nil {
			if errors.IsNotFound(err) {
				log.Printf("%s/%s Operation: %s. (Controller: OdooInstance) Error: No parent OdooInstance found for name %s\n", instance.Namespace, instance.Name, "validate", instance.Spec.ParentName)
				return reconcile.Result{}, err
			}
			// Error reading the object - requeue the request.
			return reconcile.Result{}, err
		}
	}
	// Find the Track of the OdooCluster instance by it's name ...
	foundTrack := false
	var track clusterv1beta1.TrackSpec
	for _, _track := range clusterinstance.Spec.Tracks {
		if _track.Name == instance.Spec.TrackName {
			track = _track
			foundTrack = true
		}
	}
	// ... or the one of it's parent OdooInstance
	if foundTrack != true && instance.Spec.ParentName != nil {
		for _, _track := range clusterinstance.Spec.Tracks {
			if _track.Name == parentInstance.Spec.TrackName {
				track = _track
				foundTrack = true
			}
		}
	}

	// create | delete | update | copy
	operation := "create"

	// Check a parent name is set ("child" or "testing" instance)
	if instance.Spec.ParentName != nil {
		operation = "copy"
	}

	// Check if it's a deletion operation
	if instance.GetDeletionTimestamp() != nil {
		operation = "delete"
	}

	if foundTrack != true {
		log.Printf("%s/%s Operation: %s. (Controller: OdooInstance) Error: no track with name %s exists\n", instance.Namespace, instance.Name, "validate", instance.Spec.TrackName)
		return reconcile.Result{}, fmt.Errorf("instance controller: no track with name %s exists %s", instance.Spec.TrackName)
	}

	// Define the desired Job object
	initializer := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      strings.ToLower(fmt.Sprintf("%s-%s-initializer", clusterinstance.Name, instance.Name)),
			Namespace: instance.Namespace,
		},
		Spec: batchv1.JobSpec{
			Completions:           func(a int32) *int32 { return &a }(1),
			BackoffLimit:          func(a int32) *int32 { return &a }(1),
			ActiveDeadlineSeconds: func(a int64) *int64 { return &a }(360),
			Template:              corev1.PodTemplateSpec{},
		},
	}
	clustercontroller.setPodTemplateSpec(&initializer.Spec.Template, clusterinstance, &track)
	initializer.Spec.Template.Labels["host"] = strings.ToLower(fmt.Sprintf("%s", instance.Spec.HostName))
	containerArgs := []string{"dodoo-initializer", "--new-database", strings.ToLower(fmt.Sprintf("%s", instance.Spec.HostName)), "--config", clustercontroller.appConfigsPath}
	if instance.Spec.Modules != nil {
		containerArgs = append(containerArgs, "--modules", strings.Join(instance.Spec.Modules, ","))
	}
	if instance.Spec.Demo != nil {
		containerArgs = append(containerArgs, "--demo")
	} else {
		containerArgs = append(containerArgs, "--no-demo")
	}
	for _, container := range initializer.Spec.Template.Spec.Containers {
		container.Name = strings.ToLower(fmt.Sprintf("%s-%s-initializer", container.Name, instance.Name))
		container.Args = containerArgs
		container.Ports = []corev1.ContainerPort{}
	}

	if err := controllerutil.SetControllerReference(instance, initializer, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if the Job already exists
	// Cave: Job existance signals instance existence
	found := &batchv1.Job{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: initializer.Name, Namespace: initializer.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		log.Printf("Creating OdooInstance %s/%s\n", instance.Namespace, instance.Name)
		err = r.Create(context.TODO(), initializer)
		if err != nil {
			return reconcile.Result{}, err
		}
		log.Printf("%s/%s setting finalizer: %s\n", instance.Namespace, instance.Name, FinalizerKey)
		_, err = finalizers.AddFinalizers(instance, sets.NewString(FinalizerKey))
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Check if finalizers are set
	hasFinalizer, err := finalizers.HasFinalizer(instance, FinalizerKey)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Marked for deletion: tear down!
	if instance.GetDeletionTimestamp() != nil {
		operation = "delete"
		if hasFinalizer { // Define the desired Job object
			dropper := &batchv1.Job{
				ObjectMeta: metav1.ObjectMeta{
					Name:      strings.ToLower(fmt.Sprintf("%s-%s-dropper", clusterinstance.Name, instance.Name)),
					Namespace: instance.Namespace,
				},
				Spec: batchv1.JobSpec{
					Completions:           func(a int32) *int32 { return &a }(1),
					BackoffLimit:          func(a int32) *int32 { return &a }(1),
					ActiveDeadlineSeconds: func(a int64) *int64 { return &a }(360),
					Template:              corev1.PodTemplateSpec{},
				},
			}
			clustercontroller.setPodTemplateSpec(&dropper.Spec.Template, clusterinstance, &track)
			dropper.Spec.Template.Labels["host"] = strings.ToLower(fmt.Sprintf("%s", instance.Spec.HostName))
			containerArgs := []string{"dodoo-dropper", "--database", strings.ToLower(fmt.Sprintf("%s", instance.Spec.HostName)), "--config", clustercontroller.appConfigsPath, "--if-exists"}
			for _, container := range dropper.Spec.Template.Spec.Containers {
				container.Name = strings.ToLower(fmt.Sprintf("%s-%s-dropper", container.Name, instance.Name))
				container.Args = containerArgs
				container.Ports = []corev1.ContainerPort{}
			}
			if err := controllerutil.SetControllerReference(instance, dropper, r.scheme); err != nil {
				return reconcile.Result{}, err
			}
			log.Printf("Removing OdooInstance %s/%s\n", instance.Namespace, instance.Name)
			err = r.Create(context.TODO(), dropper)
			if err != nil {
				return reconcile.Result{}, err
			}
			log.Printf("%s/%s removing finalizer: %s\n", instance.Namespace, instance.Name, FinalizerKey)
			finalizers.RemoveFinalizers(instance, sets.NewString(FinalizerKey))
		}
		log.Printf("%s/%s reconciled. Operation: %s. (Controller: OdooInstance)\n", instance.Namespace, instance.Name, operation)
		return reconcile.Result{}, nil
	}

	// Update the found object and write the result back if there are any changes
	if !reflect.DeepEqual(initializer.Spec, found.Spec) {
		found.Spec = initializer.Spec
		log.Printf("Updating OdooInstance %s/%s\n", initializer.Namespace, initializer.Name)
		err = r.Update(context.TODO(), found)
		if err != nil {
			return reconcile.Result{}, err
		}
	}
	return reconcile.Result{}, nil
}
