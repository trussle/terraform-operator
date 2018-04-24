
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketObject struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsS3BucketObjectSpec	`json:"spec"`
}

type AwsS3BucketObjectList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsS3BucketObject	`json:"items"`
}

type AwsS3BucketObjectSpec struct {
	Acl	string	`json:"acl"`
	ContentEncoding	string	`json:"content_encoding"`
	ContentLanguage	string	`json:"content_language"`
	ContentBase64	string	`json:"content_base64"`
	KmsKeyId	string	`json:"kms_key_id"`
	Tags	map[string]interface{}	`json:"tags"`
	CacheControl	string	`json:"cache_control"`
	ContentDisposition	string	`json:"content_disposition"`
	Source	string	`json:"source"`
	WebsiteRedirect	string	`json:"website_redirect"`
	Bucket	string	`json:"bucket"`
	Key	string	`json:"key"`
	Content	string	`json:"content"`
}
