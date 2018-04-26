
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
	KmsKeyId	string	`json:"kms_key_id"`
	Acl	string	`json:"acl"`
	ContentBase64	string	`json:"content_base64"`
	Source	string	`json:"source"`
	Tags	map[string]???	`json:"tags"`
	Bucket	string	`json:"bucket"`
	ContentLanguage	string	`json:"content_language"`
	Key	string	`json:"key"`
	Content	string	`json:"content"`
	CacheControl	string	`json:"cache_control"`
	ContentDisposition	string	`json:"content_disposition"`
	WebsiteRedirect	string	`json:"website_redirect"`
	ContentEncoding	string	`json:"content_encoding"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketObjectList is a list of AwsS3BucketObject resources
type AwsS3BucketObjectList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsS3BucketObject	`json:"items"`
}

