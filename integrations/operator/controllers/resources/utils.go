/*
Copyright 2022 Gravitational, Inc.

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

package resources

import (
	"context"
	"fmt"

	"github.com/gravitational/trace"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kclient "sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/gravitational/teleport/api/types"
)

const (
	ConditionReasonFailedToDecode         = "FailedToDecode"
	ConditionReasonOriginLabelNotMatching = "OriginLabelNotMatching"
	ConditionReasonOriginLabelMatching    = "OriginLabelMatching"
	ConditionReasonNewResource            = "NewResource"
	ConditionReasonNoError                = "NoError"
	ConditionReasonTeleportError          = "TeleportError"
	ConditionReasonTeleportClientError    = "TeleportClientError"
	ConditionTypeTeleportResourceOwned    = "TeleportResourceOwned"
	ConditionTypeSuccessfullyReconciled   = "SuccessfullyReconciled"
	ConditionTypeValidStructure           = "ValidStructure"
	ConditionTypeTeleportClient           = "TeleportClient"
)

var newResourceCondition = metav1.Condition{
	Type:    ConditionTypeTeleportResourceOwned,
	Status:  metav1.ConditionTrue,
	Reason:  ConditionReasonNewResource,
	Message: "No existing Teleport resource found with that name. The created resource is owned by the operator.",
}

type ownedResource interface {
	GetMetadata() types.Metadata
}

// isResourceOriginKubernetes reads a teleport resource metadata, searches for the origin label and checks its
// value is kubernetes.
func isResourceOriginKubernetes(resource ownedResource) bool {
	label := resource.GetMetadata().Labels[types.OriginLabel]
	return label == types.OriginKubernetes
}

// checkOwnership takes an existing resource and validates the operator owns it.
// It returns an ownership condition and a boolean representing if the resource is
// owned by the operator. The ownedResource must be non-nil.
func checkOwnership(existingResource ownedResource) (metav1.Condition, bool) {
	if !isResourceOriginKubernetes(existingResource) {
		// Existing Teleport resource does not belong to us, bailing out

		condition := metav1.Condition{
			Type:    ConditionTypeTeleportResourceOwned,
			Status:  metav1.ConditionFalse,
			Reason:  ConditionReasonOriginLabelNotMatching,
			Message: "A resource with the same name already exists in Teleport and does not have the Kubernetes origin label. Refusing to reconcile.",
		}
		return condition, false
	}

	condition := metav1.Condition{
		Type:    ConditionTypeTeleportResourceOwned,
		Status:  metav1.ConditionTrue,
		Reason:  ConditionReasonOriginLabelMatching,
		Message: "Teleport resource has the Kubernetes origin label.",
	}
	return condition, true
}

// getReconciliationConditionFromError takes an error returned by a call to Teleport and returns a
// metav1.Condition describing how the Teleport resource reconciliation went. This is used to provide feedback to
// the user about the controller's ability to reconcile the resource.
func getReconciliationConditionFromError(err error, ignoreNotFound bool) metav1.Condition {
	if err == nil || trace.IsNotFound(err) && ignoreNotFound {
		return metav1.Condition{
			Type:    ConditionTypeSuccessfullyReconciled,
			Status:  metav1.ConditionTrue,
			Reason:  ConditionReasonNoError,
			Message: "Teleport resource was successfully reconciled, no error was returned by Teleport.",
		}
	}
	return metav1.Condition{
		Type:    ConditionTypeSuccessfullyReconciled,
		Status:  metav1.ConditionFalse,
		Reason:  ConditionReasonTeleportError,
		Message: fmt.Sprintf("Teleport returned the error: %s", err),
	}
}

// getStructureConditionFromError takes a conversion error from k8s apimachinery's runtime.UnstructuredConverter
// and returns a metav1.Condition describing how the status conversion went. This is used to provide feedback to
// the user about the controller's ability to reconcile the resource.
func getStructureConditionFromError(err error) metav1.Condition {
	if err != nil {
		return metav1.Condition{
			Type:    ConditionTypeValidStructure,
			Status:  metav1.ConditionFalse,
			Reason:  ConditionReasonFailedToDecode,
			Message: fmt.Sprintf("Failed to decode Kubernetes CR: %s", err),
		}
	}
	return metav1.Condition{
		Type:    ConditionTypeValidStructure,
		Status:  metav1.ConditionTrue,
		Reason:  ConditionReasonNoError,
		Message: "Kubernetes CR was successfully decoded.",
	}
}

// getTeleportClientConditionFromError takes an error returned by a call to Teleport ClientAccessor and returns a
// metav1.Condition describing how the Teleport client creation went. This is used to provide feedback to
// the user about the controller's ability to reconcile the resource.
func getTeleportClientConditionFromError(err error) metav1.Condition {
	if err != nil {
		return metav1.Condition{
			Type:    ConditionTypeTeleportClient,
			Status:  metav1.ConditionFalse,
			Reason:  ConditionReasonTeleportClientError,
			Message: fmt.Sprintf("Failed to create Teleport client: %s", err),
		}
	}
	return metav1.Condition{
		Type:    ConditionTypeTeleportClient,
		Status:  metav1.ConditionTrue,
		Reason:  ConditionReasonNoError,
		Message: "Teleport client creation was successful.",
	}
}

// updateStatusConfig is a configuration struct for silentUpdateStatus.
type updateStatusConfig struct {
	ctx         context.Context
	client      kclient.Client
	k8sResource interface {
		kclient.Object
		StatusConditions() *[]metav1.Condition
	}
	condition metav1.Condition
}

// updateStatus updates the resource status but swallows the error if the update fails.
func updateStatus(config updateStatusConfig) error {
	// If the condition is empty, we don't want to update the status.
	if config.condition == (metav1.Condition{}) {
		return nil
	}
	log := ctrllog.FromContext(config.ctx)
	meta.SetStatusCondition(
		config.k8sResource.StatusConditions(),
		config.condition,
	)
	statusErr := config.client.Status().Update(config.ctx, config.k8sResource)
	if statusErr != nil {
		log.Error(statusErr, "failed to report error in status conditions")
	}
	return trace.Wrap(statusErr)
}
