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

// OdooInstanceSpec defines the desired state of OdooInstance
type OdooInstanceSpec struct {
	Cluster string `json:"cluster"`
	// The host name of this instance (mutable)
	Hostname string `json:"hostname"`
	// The version of this instance (immutable)
	// Either Version or ParentHostname is required
	// +optional
	Version string `json:"version,omitempty"`
	// +optional
	ParentHostname *string `json:"parentHostname,omitempty"`
	// +optional
	Demo *bool `json:"demo,omitempty"`
	// +optional
	InitModules []string `json:"initModules,omitempty"`
	// +optional
	InitSQL string `json:"initSQL,omitempty"`
}
