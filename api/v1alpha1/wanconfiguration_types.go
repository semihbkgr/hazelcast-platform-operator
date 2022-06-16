package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// WanConfigurationSpec defines the desired state of WanConfiguration
type WanConfigurationSpec struct {
	// MapResourceName is the name of Map custom resource which WAN replication will be applied to.
	// +kubebuilder:validation:MinLength:=1
	MapResourceName string `json:"mapResourceName"`

	// ClusterName is the clusterName field of the target Hazelcast resource.
	// +kubebuilder:validation:MinLength:=1
	TargetClusterName string `json:"targetClusterName"`

	// Endpoints is the target cluster endpoints.
	// +kubebuilder:validation:MinLength:=1
	Endpoints string `json:"endpoints"`

	// Queue is the configuration for WAN events queue.
	Queue QueueSetting `json:"queue,omitempty"`

	// Batch is the configuration for WAN events batch.
	Batch BatchSetting `json:"batch,omitempty"`

	// Acknowledgement is the configuration for the condition when the next batch of WAN events are sent.
	Acknowledgement AcknowledgementSetting `json:"acknowledgement,omitempty"`
}

// QueueSetting defines the configuration for Hazelcast WAN queue
type QueueSetting struct {
	// Capacity is the total capacity of WAN queue.
	// +kubebuilder:default:=10000
	Capacity int32 `json:"capacity,omitempty"`

	// FullBehavior represents the behavior of the new arrival when the queue is full.
	// +kubebuilder:validation:Enum=DISCARD_AFTER_MUTATION;THROW_EXCEPTION;THROW_EXCEPTION_ONLY_IF_REPLICATION_ACTIVE
	// +kubebuilder:default:=DISCARD_AFTER_MUTATION
	FullBehavior FullBehaviorSetting `json:"fullBehavior,omitempty"`
}

type FullBehaviorSetting string

const (
	DISCARD_AFTER_MUTATION FullBehaviorSetting = "DISCARD_AFTER_MUTATION"

	THROW_EXCEPTION FullBehaviorSetting = "THROW_EXCEPTION"

	THROW_EXCEPTION_ONLY_IF_REPLICATION_ACTIVE = "THROW_EXCEPTION_ONLY_IF_REPLICATION_ACTIVE"
)

type BatchSetting struct {
	// Size represents the maximum batch size.
	// +kubebuilder:default:=500
	Size int32 `json:"size,omitempty"`

	// MaximumDelay represents the maximum delay in milliseconds.
	// If the batch size is not reached, the events will be sent after
	// the maximum delay.
	// +kubebuilder:default:=1000
	MaximumDelay int32 `json:"maximumDelay,omitempty"`
}

type AcknowledgementSetting struct {
	// Type represents how a batch of replication events is considered successfully replicated.
	// +kubebuilder:validation:Enum=ACK_ON_OPERATION_COMPLETE;ACK_ON_RECEIPT
	// +kubebuilder:default:=ACK_ON_OPERATION_COMPLETE
	Type AcknowledgementType `json:"type,omitempty"`

	// Timeout represents the time the source cluster waits for the acknowledgement.
	// After timeout, the events will be resent.
	// +kubebuilder:default:=60000
	Timeout int32 `json:"timeout,omitempty"`
}

type AcknowledgementType string

const (
	ACK_ON_RECEIPT AcknowledgementType = "ACK_ON_RECEIPT"

	ACK_ON_OPERATION_COMPLETE AcknowledgementType = "ACK_ON_OPERATION_COMPLETE"
)

// WanConfigurationStatus defines the observed state of WanConfiguration
type WanConfigurationStatus struct {
	// State represents the current Hazelcast WAN Replication state
	State string `json:"state,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// WanConfiguration is the Schema for the wanconfigurations API
type WanConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WanConfigurationSpec   `json:"spec,omitempty"`
	Status WanConfigurationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// WanConfigurationList contains a list of WanConfiguration
type WanConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WanConfiguration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WanConfiguration{}, &WanConfigurationList{})
}