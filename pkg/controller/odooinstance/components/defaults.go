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

package components

import (
	// "fmt"

	"github.com/Ridecell/ridecell-operator/pkg/components"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	instancev1beta1 "github.com/xoe-labs/odoo-operator/pkg/apis/instance/v1beta1"
)

type defaultsComponent struct {
}

func NewDefaults() *defaultsComponent {
	return &defaultsComponent{}
}

func (_ *defaultsComponent) WatchTypes() []runtime.Object {
	return []runtime.Object{}
}

func (_ *defaultsComponent) IsReconcilable(_ *components.ComponentContext) bool {
	// Always "reconcile" (apply) Defaults as the first controller action
	return true
}

func (comp *defaultsComponent) Reconcile(ctx *components.ComponentContext) (reconcile.Result, error) {
	instance := ctx.Top.(*instancev1beta1.OdooInstance)

	// Fill in defaults.
	if instance.Spec.Hostname == "" {
		instance.Spec.Hostname = instance.Name + ".xoe.cloud"
	}

	return reconcile.Result{}, nil
}
