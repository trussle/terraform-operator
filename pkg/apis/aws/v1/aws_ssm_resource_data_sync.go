
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmResourceDataSync describes a AwsSsmResourceDataSync resource
type AwsSsmResourceDataSync struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmResourceDataSyncSpec	`json:"spec"`
}


// AwsSsmResourceDataSyncSpec is the spec for a AwsSsmResourceDataSync Resource
type AwsSsmResourceDataSyncSpec struct {
	S3Destination	[]AwsSsmResourceDataSyncS3Destination	`json:"s3_destination"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmResourceDataSyncList is a list of AwsSsmResourceDataSync resources
type AwsSsmResourceDataSyncList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmResourceDataSync	`json:"items"`
}


// AwsSsmResourceDataSyncS3Destination is a AwsSsmResourceDataSyncS3Destination
type AwsSsmResourceDataSyncS3Destination struct {
	Prefix	string	`json:"prefix"`
	Region	string	`json:"region"`
	SyncFormat	string	`json:"sync_format"`
	KmsKeyArn	string	`json:"kms_key_arn"`
	BucketName	string	`json:"bucket_name"`
}
