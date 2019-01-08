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
 * a commercial license.  Buying such a license is mandatory as soon as you
 * develop commercial activities involving the Odoo-Operator software without
 * disclosing the source code of your own applications. These activities
 * include: Offering paid services to a customer as an ASP, shipping Odoo-
 * Operator with a closed source product.
 *
 */

package v1beta1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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

// OdooImageRepoSpec defines the cluster-level image configuration
type OdooImageRepoSpec struct {
	Registry string `json:"registry"`
	Repo     string `json:"repo"`
	// +optional
	Secret string `json:"secret,omitempty"`
}

// OdooResource defines the desired resource spec
type OdooResource struct {
	Tiers   []OdooTierSpec   `json:"tiers,omitempty"`
	Volumes []OdooVolumeSpec `json:"volumes,omitempty"`
}

// OdooTierSpec specs a tier
type OdooTierSpec struct {
	Name string `json:"name"`
	// +optional
	Replicas int32 `json:"replicas,omitempty"`
	// +optional
	QOS *v1.PodQOSClass `json:"qos,omitempty"`
	// +optional
	Resources *v1.ResourceRequirements `json:"resources,omitempty"`
}

// OdooVolumeSpec specs a volume
type OdooVolumeSpec struct {
	Name string `json:"name"`
	// +optional
	PVCSpec v1.PersistentVolumeClaimSpec `json:"pvcspec"`
}

// OdooClusterStatus defines the observed state of OdooCluster
type OdooClusterStatus struct {
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []OdooClusterStatusCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// OdooClusterStatusCondition defines an observable OdooClusterStatus condition
type OdooClusterStatusCondition struct {
	// Type of the OdooClusterStatus condition.
	Type OdooClusterStatusConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=OdooClusterStatusConditionType"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/v1.ConditionStatus"`
	// The last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,3,opt,name=lastTransitionTime"`
	// The reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
	// A human readable message indicating details about the transition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,5,opt,name=message"`
}

// OdooClusterStatusConditionType ...
type OdooClusterStatusConditionType string

const (
	// OdooClusterStatusConditionTypeCreated ...
	OdooClusterStatusConditionTypeCreated OdooClusterStatusConditionType = "Created"
	// OdooClusterStatusConditionTypeReconciled ...
	OdooClusterStatusConditionTypeReconciled OdooClusterStatusConditionType = "Reconciled"
	// OdooClusterStatusConditionTypeAppSecretLoaned ...
	OdooClusterStatusConditionTypeAppSecretLoaned OdooClusterStatusConditionType = "AppSecretLoaned"
	// OdooClusterStatusConditionTypePullSecretLoaned ...
	OdooClusterStatusConditionTypePullSecretLoaned OdooClusterStatusConditionType = "PullSecretLoaned"
	// OdooClusterStatusConditionTypeErrored ...
	OdooClusterStatusConditionTypeErrored OdooClusterStatusConditionType = "Errored"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OdooCluster is the Schema for the odooclusters API
// +k8s:openapi-gen=true
type OdooCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OdooClusterSpec   `json:"spec,omitempty"`
	Status OdooClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OdooClusterList contains a list of OdooCluster
type OdooClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OdooCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OdooCluster{}, &OdooClusterList{})
}
