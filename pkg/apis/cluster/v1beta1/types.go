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

package v1beta1

// Gross workaround for limitations the Kubernetes code generator and interface{}.
// If you want to see the weird inner workings of the hack, looking marshall.go.
type ConfigValue struct {
	Bool    *bool                  `json:"bool,omitempty"`
	Int     *int                   `json:"int,omitempty"`
	Float   *float64               `json:"float,omitempty"`
	String  *string                `json:"string,omitempty"`
	Section map[string]ConfigValue `json:"section,omitempty"`
}

// OdooImageSpec defines an Image and (optionally) it's registry credentials
type OdooImageSpec struct {
	Repository string            `json:"repository"`
	Image      string            `json:"image"`
	Trackname  OdooTracknameType `json:"track"`
	Version    string            `json:"version"`
	// +optional
	Secret string `json:"secret,omitempty"`
}
