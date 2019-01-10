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

import (
	"k8s.io/api/core/v1"
)

// DBNamespaceSpec defines the desired state of DBNamespace
type DBNamespaceSpec struct {
	User      string             `json:"user"`
	Password  string             `json:"password"`
	Admin     DBAdminCredentials `json:"dbAdmin"`
	UserQuota v1.ResourceList    `json:"userQuota,omitempty"`
}

// OdooClusterSpec defines the desired state of OdooCluster
type OdooClusterSpec struct {
	Tracks   []OdooTrackSpec        `json:"tracks"`
	Image    OdooImageRepoSpec      `json:"image"`
	Database DBNamespaceSpec        `json:"database"`
	Config   map[string]ConfigValue `json:"config"`
	// +optional
	AppSecret string `json:"appsecret,omitempty"`
	// +optional
	Resources OdooResource `json:"resources,omitempty"`
	// +optional
	NodeSelector *v1.NodeSelector `json:"nodes,omitempty"`

	// MailServer  bool `json:"mailServer"`
	// OnlyOffice  bool `json:"onlyOffice"`
	// Mattermost  bool `json:"mattermost"`
	// Nuxeo       bool `json:"nuxeo"`
	// BpmnEngine  bool `json:"bpmnEngine"`
	// OpenProject bool `json:"openProject"`
}

// OdooTrackSpec defines the desired state of OdooTrack
type OdooTrackSpec struct {
	Name OdooTracknameType `json:"name"`
	// +optional
	Config map[string]ConfigValue `json:"config,omitempty"`
}

// OdooVersionSpec defines the desired state of OdooVersion
type OdooVersionSpec struct {
	Version string            `json:"name"`
	Track   OdooTracknameType `json:"track"`
	Bugfix  bool              `json:"bugfix"`
	// +optional
	Config map[string]ConfigValue `json:"config,omitempty"`
}
