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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OdooVersionSpec defines the desired state of OdooVersion
type OdooVersionSpec struct {
	Version string `json:"name"`
	Bugfix  bool   `json:"bugfix"`
	// +optional
	Config *string `json:"config,omitempty"`
}

// OdooVersionStatus defines the observed state of OdooVersion
type OdooVersionStatus struct {
	// Current service state of apiService.
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []OdooVersionStatusCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// OdooVersionStatusCondition defines an observable OdooVersionStatus condition
type OdooVersionStatusCondition struct {
	// Type of the OdooVersionStatus condition.
	Type OdooVersionStatusConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=OdooVersionStatusConditionType"`
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

// OdooVersionStatusConditionType ...
type OdooVersionStatusConditionType string

const (
	// OdooVersionStatusConditionTypeDeployed ...
	OdooVersionStatusConditionTypeDeployed OdooVersionStatusConditionType = "Deployed"
	// OdooVersionStatusConditionTypeApplied ...
	OdooVersionStatusConditionTypeApplied OdooVersionStatusConditionType = "Applied"
	// OdooVersionStatusConditionTypeReconciled ...
	OdooVersionStatusConditionTypeReconciled OdooVersionStatusConditionType = "Reconciled"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OdooVersion is the Schema for the odooversions API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type OdooVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OdooVersionSpec   `json:"spec,omitempty"`
	Status OdooVersionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OdooVersionList contains a list of OdooVersion
type OdooVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OdooVersion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OdooVersion{}, &OdooVersionList{})
}
