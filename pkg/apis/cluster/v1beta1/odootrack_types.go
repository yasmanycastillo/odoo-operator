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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OdooTrackSpec defines the desired state of OdooTrack
type OdooTrackSpec struct {
	Name         OdooTracknameType `json:"name"`
	StartVersion OdooVersionSpec   `json:"startversion"`
	// +optional
	Config *string `json:"config,omitempty"`
}

// OdooTracknameType can be one of the supported tracks
type OdooTracknameType string

const (
	// OdooTracknameTypeMaster ...
	OdooTracknameTypeMaster OdooTracknameType = "master"
	// OdooTracknameType12 ...
	OdooTracknameType12 OdooTracknameType = "12.0"
	// OdooTracknameType11 ...
	OdooTracknameType11 OdooTracknameType = "11.0"
	// OdooTracknameType10 ...
	OdooTracknameType10 OdooTracknameType = "10.0"
)

// OdooTrackStatus defines the observed state of OdooTrack
type OdooTrackStatus struct {
	CurrentHead OdooVersion `json:"currenthead"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OdooTrack is the Schema for the odootracks API
// +k8s:openapi-gen=true
type OdooTrack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OdooTrackSpec   `json:"spec,omitempty"`
	Status OdooTrackStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OdooTrackList contains a list of OdooTrack
type OdooTrackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OdooTrack `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OdooTrack{}, &OdooTrackList{})
}
