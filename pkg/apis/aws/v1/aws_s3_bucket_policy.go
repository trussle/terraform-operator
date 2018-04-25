
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketPolicy describes a AwsS3BucketPolicy resource
type AwsS3BucketPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsS3BucketPolicySpec	`json:"spec"`
}


// AwsS3BucketPolicySpec is the spec for a AwsS3BucketPolicy Resource
type AwsS3BucketPolicySpec struct {
	Bucket	string	`json:"bucket"`
	Policy	string	`json:"policy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketPolicyList is a list of AwsS3BucketPolicy resources
type AwsS3BucketPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsS3BucketPolicy	`json:"items"`
}

