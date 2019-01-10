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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/blaggacao/ridecell-operator/pkg/components"
)

// OdooInstance Status ...
func (c *OdooInstance) GetStatus() components.Status {
	return c.Status
}

func (c *OdooInstance) SetStatus(status components.Status) {
	c.Status = status.(OdooInstanceStatus)
}

func (c *OdooInstance) SetErrorStatus(errorMsg string) {
	// c.Status.Status = StatusError
	// c.Status.Message = errorMsg
}

func (_ *OdooInstance) NewStatusCondition(
	condType OdooInstanceStatusConditionType, status corev1.ConditionStatus,
	reason, message string) *OdooInstanceStatusCondition {
	return &OdooInstanceStatusCondition{condType, StatusCondition{
		Status:             status,
		LastTransitionTime: metav1.Now(),
		Reason:             reason,
		Message:            message,
	}}
}

func (c *OdooInstance) GetStatusCondition(
	condType OdooInstanceStatusConditionType) *OdooInstanceStatusCondition {
	for i := range c.Status.Conditions {
		cond := c.Status.Conditions[i]
		if cond.Type == condType {
			return &cond
		}
	}
	return nil
}

func (c *OdooInstance) SetStatusCondition(condition OdooInstanceStatusCondition) {
	currentCond := c.GetStatusCondition(condition.Type)
	if currentCond != nil && currentCond.Status == condition.Status && currentCond.Reason == condition.Reason {
		return
	}
	// Do not update lastTransitionTime if the status of the condition doesn't change.
	if currentCond != nil && currentCond.Status == condition.Status {
		condition.LastTransitionTime = currentCond.LastTransitionTime
	}
	newConditions := c.filterOutCondition(condition.Type)
	c.Status.Conditions = append(newConditions, condition)
}

func (c *OdooInstance) RemoveStatusCondition(condType OdooInstanceStatusConditionType) {
	c.Status.Conditions = c.filterOutCondition(condType)
}

func (c *OdooInstance) filterOutCondition(condType OdooInstanceStatusConditionType) []OdooInstanceStatusCondition {
	var newConditions []OdooInstanceStatusCondition
	for _, cond := range c.Status.Conditions {
		if cond.Type == condType {
			continue
		}
		newConditions = append(newConditions, cond)
	}
	return newConditions
}
