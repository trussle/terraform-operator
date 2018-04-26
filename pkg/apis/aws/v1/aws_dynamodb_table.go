
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
	RangeKey	string	`json:"range_key"`
	WriteCapacity	int	`json:"write_capacity"`
	Ttl	Generic	`json:"ttl"`
	Tags	map[string]Generic	`json:"tags"`
	ReadCapacity	int	`json:"read_capacity"`
	StreamEnabled	bool	`json:"stream_enabled"`
	LocalSecondaryIndex	Generic	`json:"local_secondary_index"`
	Name	string	`json:"name"`
	HashKey	string	`json:"hash_key"`
	Attribute	Generic	`json:"attribute"`
	GlobalSecondaryIndex	Generic	`json:"global_secondary_index"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDynamodbTableList is a list of AwsDynamodbTable resources
type AwsDynamodbTableList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDynamodbTable	`json:"items"`
}

