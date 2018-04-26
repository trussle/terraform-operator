
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketMetric describes a AwsS3BucketMetric resource
type AwsS3BucketMetric struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsS3BucketMetricSpec	`json:"spec"`
}


// AwsS3BucketMetricSpec is the spec for a AwsS3BucketMetric Resource
type AwsS3BucketMetricSpec struct {
	Name	string	`json:"name"`
	Bucket	string	`json:"bucket"`
	Filter	[]Generic	`json:"filter"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketMetricList is a list of AwsS3BucketMetric resources
type AwsS3BucketMetricList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsS3BucketMetric	`json:"items"`
}

