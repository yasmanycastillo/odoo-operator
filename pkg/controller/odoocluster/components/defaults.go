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
	"github.com/blaggacao/ridecell-operator/pkg/components"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	clusterv1beta1 "github.com/xoe-labs/odoo-operator/pkg/apis/cluster/v1beta1"
)

var configDefaults map[string]clusterv1beta1.ConfigValue

type defaultComponent struct{}

func NewDefaults() *defaultComponent {
	return &defaultComponent{}
}

func (_ *defaultComponent) WatchTypes() []runtime.Object {
	return []runtime.Object{}
}

func (_ *defaultComponent) IsReconcilable(_ *components.ComponentContext) bool {
	return true
}

func (comp *defaultComponent) Reconcile(ctx *components.ComponentContext) (reconcile.Result, error) {
	instance := ctx.Top.(*clusterv1beta1.OdooCluster)

	// Fill in defaults.
	// if instance.Spec.Tracks == nil {
	// 	instance.Spec.Tracks = instance.Spec.Tracks
	// }

	// if instance.Spec.ResourceQuotaSpec == nil {
	// 	instance.Spec.ResourceQuotaSpec = instance.Spec.ResourceQuotaSpec
	// }

	// Fill in static default config values.
	if instance.Spec.Config == nil {
		instance.Spec.Config = map[string]clusterv1beta1.ConfigValue{}
	}

	for key, value := range configDefaults {
		_, ok := instance.Spec.Config[key]
		if !ok {
			instance.Spec.Config[key] = value
		}
	}

	return reconcile.Result{}, nil
}

func defConfig(sec, key string, value interface{}) {
	if sec != "" {
		_, ok := configDefaults[sec]
		if !ok {
			configDefaults[sec] = clusterv1beta1.ConfigValue{
				Section: map[string]clusterv1beta1.ConfigValue{},
			}
		}
		section, _ := configDefaults[sec]
		section.Section[key] = *getValue(value)
	}
	configDefaults[key] = *getValue(value)
}
func getValue(value interface{}) *clusterv1beta1.ConfigValue {
	boolVal, ok := value.(bool)
	if ok {
		return &clusterv1beta1.ConfigValue{Bool: &boolVal}
	}
	intVal, ok := value.(int)
	if ok {
		return &clusterv1beta1.ConfigValue{Int: &intVal}
	}
	floatVal, ok := value.(float64)
	if ok {
		return &clusterv1beta1.ConfigValue{Float: &floatVal}
	}
	stringVal, ok := value.(string)
	if ok {
		return &clusterv1beta1.ConfigValue{String: &stringVal}
	}
	panic("Unknown type")
}
func init() {
	configDefaults = map[string]clusterv1beta1.ConfigValue{}
	// Default config for prod environment.
	defConfig("options", "list_db", false)
	defConfig("options", "unaccent", true)
	defConfig("options", "proxy_mode", true)
	defConfig("options", "dbfilter", "^%h$")
	defConfig("options", "data_dir", "/mnt/odoo/data")
	defConfig("options", "backupfolder", "/mnt/odoo/backup")
	defConfig("options", "publisher_warranty_url", "http://services.openerp.com/publisher-warranty/")
}
