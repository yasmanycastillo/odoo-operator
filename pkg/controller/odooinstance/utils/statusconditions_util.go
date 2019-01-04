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

	instancev1beta1 "github.com/xoe-labs/odoo-operator/pkg/apis/instance/v1beta1"
)

// NewOdooInstanceStatusCondition creates a new odooinstance condition.
func NewOdooInstanceStatusCondition(condType instancev1beta1.OdooInstanceStatusConditionType, status corev1.ConditionStatus, reason, message string) *instancev1beta1.OdooInstanceStatusCondition {
	return &instancev1beta1.OdooInstanceStatusCondition{
		Type:               condType,
		Status:             status,
		LastTransitionTime: metav1.Now(),
		Reason:             reason,
		Message:            message,
	}
}

// GetOdooInstanceStatusCondition returns the condition with the provided type.
func GetOdooInstanceStatusCondition(status instancev1beta1.OdooInstanceStatus, condType instancev1beta1.OdooInstanceStatusConditionType) *instancev1beta1.OdooInstanceStatusCondition {
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}
	return nil
}

// SetOdooInstanceStatusCondition updates the odooinstance to include the provided condition. If the condition that
// we are about to add already exists and has the same status and reason then we are not going to update.
func SetOdooInstanceStatusCondition(status *instancev1beta1.OdooInstanceStatus, condition instancev1beta1.OdooInstanceStatusCondition) {
	currentCond := GetOdooInstanceStatusCondition(*status, condition.Type)
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

// RemoveOdooInstanceStatusCondition removes the odooinstance condition with the provided type.
func RemoveOdooInstanceStatusCondition(status *instancev1beta1.OdooInstanceStatus, condType instancev1beta1.OdooInstanceStatusConditionType) {
	status.Conditions = filterOutCondition(status.Conditions, condType)
}

// filterOutCondition returns a new slice of odooinstance conditions without conditions with the provided type.
func filterOutCondition(conditions []instancev1beta1.OdooInstanceStatusCondition, condType instancev1beta1.OdooInstanceStatusConditionType) []instancev1beta1.OdooInstanceStatusCondition {
	var newConditions []instancev1beta1.OdooInstanceStatusCondition
	for _, c := range conditions {
		if c.Type == condType {
			continue
		}
		newConditions = append(newConditions, c)
	}
	return newConditions
}
