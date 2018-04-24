
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDynamodbTable struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDynamodbTableSpec	`json:"spec"`
}

type AwsDynamodbTableList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDynamodbTable	`json:"items"`
}

type AwsDynamodbTableSpec struct {
	WriteCapacity	int	`json:"write_capacity"`
	GlobalSecondaryIndex	interface{}	`json:"global_secondary_index"`
	Tags	map[string]interface{}	`json:"tags"`
	LocalSecondaryIndex	interface{}	`json:"local_secondary_index"`
	Name	string	`json:"name"`
	RangeKey	string	`json:"range_key"`
	Attribute	interface{}	`json:"attribute"`
	Ttl	interface{}	`json:"ttl"`
	HashKey	string	`json:"hash_key"`
	ReadCapacity	int	`json:"read_capacity"`
	StreamEnabled	bool	`json:"stream_enabled"`
}
