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

// NewOdooClusterStatusCondition creates a new odoocluster condition.
func NewOdooClusterStatusCondition(condType clusterv1beta1.OdooClusterStatusConditionType, status corev1.ConditionStatus, reason, message string) *clusterv1beta1.OdooClusterStatusCondition {
	return &clusterv1beta1.OdooClusterStatusCondition{
		Type:               condType,
		Status:             status,
		LastTransitionTime: metav1.Now(),
		Reason:             reason,
		Message:            message,
	}
}

// GetOdooClusterStatusCondition returns the condition with the provided type.
func GetOdooClusterStatusCondition(status clusterv1beta1.OdooClusterStatus, condType clusterv1beta1.OdooClusterStatusConditionType) *clusterv1beta1.OdooClusterStatusCondition {
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == condType {
			return &c
		}
	}
	return nil
}

// SetOdooClusterStatusCondition updates the odoocluster to include the provided condition. If the condition that
// we are about to add already exists and has the same status and reason then we are not going to update.
func SetOdooClusterStatusCondition(status *clusterv1beta1.OdooClusterStatus, condition clusterv1beta1.OdooClusterStatusCondition) {
	currentCond := GetOdooClusterStatusCondition(*status, condition.Type)
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

// RemoveOdooClusterStatusCondition removes the odoocluster condition with the provided type.
func RemoveOdooClusterStatusCondition(status *clusterv1beta1.OdooClusterStatus, condType clusterv1beta1.OdooClusterStatusConditionType) {
	status.Conditions = filterOutCondition(status.Conditions, condType)
}

// filterOutCondition returns a new slice of odoocluster conditions without conditions with the provided type.
func filterOutCondition(conditions []clusterv1beta1.OdooClusterStatusCondition, condType clusterv1beta1.OdooClusterStatusConditionType) []clusterv1beta1.OdooClusterStatusCondition {
	var newConditions []clusterv1beta1.OdooClusterStatusCondition
	for _, c := range conditions {
		if c.Type == condType {
			continue
		}
		newConditions = append(newConditions, c)
	}
	return newConditions
}
