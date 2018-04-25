
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDynamodbTable describes a AwsDynamodbTable resource
type AwsDynamodbTable struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDynamodbTableSpec	`json:"spec"`
}


// AwsDynamodbTableSpec is the spec for a AwsDynamodbTable Resource
type AwsDynamodbTableSpec struct {
	StreamEnabled	bool	`json:"stream_enabled"`
	LocalSecondaryIndex	string	`json:"local_secondary_index"`
	HashKey	string	`json:"hash_key"`
	ReadCapacity	int	`json:"read_capacity"`
	Attribute	string	`json:"attribute"`
	Name	string	`json:"name"`
	WriteCapacity	int	`json:"write_capacity"`
	GlobalSecondaryIndex	string	`json:"global_secondary_index"`
	Tags	map[string]interface{}	`json:"tags"`
	RangeKey	string	`json:"range_key"`
	Ttl	string	`json:"ttl"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDynamodbTableList is a list of AwsDynamodbTable resources
type AwsDynamodbTableList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDynamodbTable	`json:"items"`
}

