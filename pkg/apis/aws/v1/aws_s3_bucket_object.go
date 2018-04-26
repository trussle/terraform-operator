
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketObject describes a AwsS3BucketObject resource
type AwsS3BucketObject struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsS3BucketObjectSpec	`json:"spec"`
}


// AwsS3BucketObjectSpec is the spec for a AwsS3BucketObject Resource
type AwsS3BucketObjectSpec struct {
	ContentDisposition	string	`json:"content_disposition"`
	ContentLanguage	string	`json:"content_language"`
	Source	string	`json:"source"`
	ContentBase64	string	`json:"content_base64"`
	Tags	map[string]string	`json:"tags"`
	CacheControl	string	`json:"cache_control"`
	KmsKeyId	string	`json:"kms_key_id"`
	WebsiteRedirect	string	`json:"website_redirect"`
	Key	string	`json:"key"`
	ContentEncoding	string	`json:"content_encoding"`
	Bucket	string	`json:"bucket"`
	Content	string	`json:"content"`
	Acl	string	`json:"acl"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketObjectList is a list of AwsS3BucketObject resources
type AwsS3BucketObjectList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsS3BucketObject	`json:"items"`
}

