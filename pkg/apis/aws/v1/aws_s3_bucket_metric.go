
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketMetric struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsS3BucketMetricSpec	`json:"spec"`
}

type AwsS3BucketMetricList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsS3BucketMetric	`json:"items"`
}

type AwsS3BucketMetricSpec struct {
	Filter	[]interface{}	`json:"filter"`
	Name	string	`json:"name"`
	Bucket	string	`json:"bucket"`
}
