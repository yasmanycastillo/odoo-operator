/*
Copyright 2016 The Kubernetes Authors.
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

package utils

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	clusterv1beta1 "github.com/xoe-labs/odoo-operator/pkg/apis/cluster/v1beta1"
)

// NewOdooVersionStatusCondition creates a new OdooVersion condition.
func NewOdooVersionStatusCondition(condType clusterv1beta1.OdooVersionStatusConditionType, status corev1.ConditionStatus, reason, message string) *clusterv1beta1.OdooVersionStatusCondition {
	return &clusterv1beta1.OdooVersionStatusCondition{
		Type:               condType,
		Status:             status,
		LastTransitionTime: metav1.Now(),
		Reason:             reason,
		Message:            message,
	}
}

// GetOdooVersionStatusCondition returns the condition with the provided type.
func GetOdooVersionStatusCondition(status clusterv1beta1.OdooVersionStatus, condType clusterv1beta1.OdooVersionStatusConditionType) *clusterv1beta1.OdooVersionStatusCondition {
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}
	return nil
}

// SetOdooVersionStatusCondition updates the OdooVersion to include the provided condition. If the condition that
// we are about to add already exists and has the same status and reason then we are not going to update.
func SetOdooVersionStatusCondition(status *clusterv1beta1.OdooVersionStatus, condition clusterv1beta1.OdooVersionStatusCondition) {
	currentCond := GetOdooVersionStatusCondition(*status, condition.Type)
	if currentCond != nil && currentCond.Status == condition.Status && currentCond.Reason == condition.Reason {
		return
	}
	// Do not update lastTransitionTime if the status of the condition doesn't change.
	if currentCond != nil && currentCond.Status == condition.Status {
		condition.LastTransitionTime = currentCond.LastTransitionTime
	}
	newConditions := filterOutCondition(status.Conditions, condition.Type)
	status.Conditions = append(newConditions, condition)
}

// RemoveOdooVersionStatusCondition removes the OdooVersion condition with the provided type.
func RemoveOdooVersionStatusCondition(status *clusterv1beta1.OdooVersionStatus, condType clusterv1beta1.OdooVersionStatusConditionType) {
	status.Conditions = filterOutCondition(status.Conditions, condType)
}

// filterOutCondition returns a new slice of OdooVersion conditions without conditions with the provided type.
func filterOutCondition(conditions []clusterv1beta1.OdooVersionStatusCondition, condType clusterv1beta1.OdooVersionStatusConditionType) []clusterv1beta1.OdooVersionStatusCondition {
	var newConditions []clusterv1beta1.OdooVersionStatusCondition
	for _, c := range conditions {
		if c.Type == condType {
			continue
		}
		newConditions = append(newConditions, c)
	}
	return newConditions
}
