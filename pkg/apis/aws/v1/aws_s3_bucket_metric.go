
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Bucket	string	`json:"bucket"`
	Filter	[]zgaczAIn	`json:"filter"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketMetricList is a list of AwsS3BucketMetric resources
type AwsS3BucketMetricList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsS3BucketMetric	`json:"items"`
}


// zgaczAIn is a zgaczAIn
type zgaczAIn struct {
	Tags	map[string]???	`json:"tags"`
	Prefix	string	`json:"prefix"`
}
