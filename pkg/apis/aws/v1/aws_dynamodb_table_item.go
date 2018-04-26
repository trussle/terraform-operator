
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDynamodbTableItem describes a AwsDynamodbTableItem resource
type AwsDynamodbTableItem struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDynamodbTableItemSpec	`json:"spec"`
}


// AwsDynamodbTableItemSpec is the spec for a AwsDynamodbTableItem Resource
type AwsDynamodbTableItemSpec struct {
	TableName	string	`json:"table_name"`
	HashKey	string	`json:"hash_key"`
	RangeKey	string	`json:"range_key"`
	Item	string	`json:"item"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDynamodbTableItemList is a list of AwsDynamodbTableItem resources
type AwsDynamodbTableItemList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDynamodbTableItem	`json:"items"`
}

