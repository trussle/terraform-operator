
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDynamodbTableItem struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDynamodbTableItemSpec	`json:"spec"`
}

type AwsDynamodbTableItemList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDynamodbTableItem	`json:"items"`
}

type AwsDynamodbTableItemSpec struct {
	RangeKey	string	`json:"range_key"`
	Item	string	`json:"item"`
	TableName	string	`json:"table_name"`
	HashKey	string	`json:"hash_key"`
}
