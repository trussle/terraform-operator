
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Name	string	`json:"name"`
	Attribute	Attribute	`json:"attribute"`
	RangeKey	string	`json:"range_key"`
	ReadCapacity	int	`json:"read_capacity"`
	LocalSecondaryIndex	LocalSecondaryIndex	`json:"local_secondary_index"`
	HashKey	string	`json:"hash_key"`
	Ttl	Ttl	`json:"ttl"`
	Tags	map[string]string	`json:"tags"`
	WriteCapacity	int	`json:"write_capacity"`
	GlobalSecondaryIndex	GlobalSecondaryIndex	`json:"global_secondary_index"`
	StreamEnabled	bool	`json:"stream_enabled"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDynamodbTableList is a list of AwsDynamodbTable resources
type AwsDynamodbTableList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDynamodbTable	`json:"items"`
}


// ServerSideEncryption is a ServerSideEncryption
type ServerSideEncryption struct {
	Enabled	bool	`json:"enabled"`
}

// Attribute is a Attribute
type Attribute struct {
	Name	string	`json:"name"`
	Type	string	`json:"type"`
}

// LocalSecondaryIndex is a LocalSecondaryIndex
type LocalSecondaryIndex struct {
	Name	string	`json:"name"`
	RangeKey	string	`json:"range_key"`
	ProjectionType	string	`json:"projection_type"`
	NonKeyAttributes	[]string	`json:"non_key_attributes"`
}

// Ttl is a Ttl
type Ttl struct {
	AttributeName	string	`json:"attribute_name"`
	Enabled	bool	`json:"enabled"`
}

// GlobalSecondaryIndex is a GlobalSecondaryIndex
type GlobalSecondaryIndex struct {
	HashKey	string	`json:"hash_key"`
	RangeKey	string	`json:"range_key"`
	ProjectionType	string	`json:"projection_type"`
	NonKeyAttributes	[]string	`json:"non_key_attributes"`
	Name	string	`json:"name"`
	WriteCapacity	int	`json:"write_capacity"`
	ReadCapacity	int	`json:"read_capacity"`
}
