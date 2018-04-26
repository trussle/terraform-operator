
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSqsQueue describes a AwsSqsQueue resource
type AwsSqsQueue struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSqsQueueSpec	`json:"spec"`
}


// AwsSqsQueueSpec is the spec for a AwsSqsQueue Resource
type AwsSqsQueueSpec struct {
	MessageRetentionSeconds	int	`json:"message_retention_seconds"`
	RedrivePolicy	string	`json:"redrive_policy"`
	FifoQueue	bool	`json:"fifo_queue"`
	ContentBasedDeduplication	bool	`json:"content_based_deduplication"`
	MaxMessageSize	int	`json:"max_message_size"`
	Tags	map[string]Generic	`json:"tags"`
	DelaySeconds	int	`json:"delay_seconds"`
	ReceiveWaitTimeSeconds	int	`json:"receive_wait_time_seconds"`
	KmsMasterKeyId	string	`json:"kms_master_key_id"`
	NamePrefix	string	`json:"name_prefix"`
	VisibilityTimeoutSeconds	int	`json:"visibility_timeout_seconds"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSqsQueueList is a list of AwsSqsQueue resources
type AwsSqsQueueList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSqsQueue	`json:"items"`
}

