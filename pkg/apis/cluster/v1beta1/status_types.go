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
)

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

func (c *OdooCluster) SetStatusConditionOperatorNamespaceErrored() {
	condition := c.NewStatusCondition(
		OdooClusterStatusConditionTypeErrored,
		corev1.ConditionTrue,
		"OperatorNamespace",
		"The operator namespace is not accesible for secrets loaning.")
	c.SetStatusCondition(*condition)
}

func (c *OdooCluster) SetStatusConditionSecretLoaningNotFoundErrored() {
	condition := c.NewStatusCondition(
		OdooClusterStatusConditionTypeErrored,
		corev1.ConditionTrue,
		"SecretLoaningNotFound",
		"The app secret was not found in the operator namespace for loaning.")
	c.SetStatusCondition(*condition)
}

func (c *OdooCluster) SetStatusConditionSecretLoaningAdminPasswdNotFoundErrored() {
	condition := c.NewStatusCondition(
		OdooClusterStatusConditionTypeErrored,
		corev1.ConditionTrue,
		"SecretLoaningAdminPasswdNotFound",
		"The app secret did not contain the expected `adminpasswd` key for loaning.")
	c.SetStatusCondition(*condition)
}

func (c *OdooCluster) SetStatusConditionSecretLoaningSuccessAppSecretLoaned() {
	condition := c.NewStatusCondition(
		OdooClusterStatusConditionTypeAppSecretLoaned,
		corev1.ConditionTrue,
		"LoaningSuccess",
		"The app secret has been loaned from the operator namespace.")
	c.SetStatusCondition(*condition)
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
