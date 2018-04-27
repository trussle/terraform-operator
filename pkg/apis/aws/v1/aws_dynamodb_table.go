
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
	WriteCapacity	int	`json:"write_capacity"`
	LocalSecondaryIndex	AwsDynamodbTableLocalSecondaryIndex	`json:"local_secondary_index"`
	GlobalSecondaryIndex	AwsDynamodbTableGlobalSecondaryIndex	`json:"global_secondary_index"`
	HashKey	string	`json:"hash_key"`
	StreamEnabled	bool	`json:"stream_enabled"`
	Tags	map[string]string	`json:"tags"`
	RangeKey	string	`json:"range_key"`
	ReadCapacity	int	`json:"read_capacity"`
	Attribute	AwsDynamodbTableAttribute	`json:"attribute"`
	Ttl	AwsDynamodbTableTtl	`json:"ttl"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDynamodbTableList is a list of AwsDynamodbTable resources
type AwsDynamodbTableList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDynamodbTable	`json:"items"`
}


// AwsDynamodbTableGlobalSecondaryIndex is a AwsDynamodbTableGlobalSecondaryIndex
type AwsDynamodbTableGlobalSecondaryIndex struct {
	WriteCapacity	int	`json:"write_capacity"`
	ReadCapacity	int	`json:"read_capacity"`
	HashKey	string	`json:"hash_key"`
	RangeKey	string	`json:"range_key"`
	ProjectionType	string	`json:"projection_type"`
	NonKeyAttributes	[]string	`json:"non_key_attributes"`
	Name	string	`json:"name"`
}

// AwsDynamodbTableServerSideEncryption is a AwsDynamodbTableServerSideEncryption
type AwsDynamodbTableServerSideEncryption struct {
	Enabled	bool	`json:"enabled"`
}

// AwsDynamodbTableAttribute is a AwsDynamodbTableAttribute
type AwsDynamodbTableAttribute struct {
	Name	string	`json:"name"`
	Type	string	`json:"type"`
}

// AwsDynamodbTableTtl is a AwsDynamodbTableTtl
type AwsDynamodbTableTtl struct {
	Enabled	bool	`json:"enabled"`
	AttributeName	string	`json:"attribute_name"`
}

// AwsDynamodbTableLocalSecondaryIndex is a AwsDynamodbTableLocalSecondaryIndex
type AwsDynamodbTableLocalSecondaryIndex struct {
	Name	string	`json:"name"`
	RangeKey	string	`json:"range_key"`
	ProjectionType	string	`json:"projection_type"`
	NonKeyAttributes	[]string	`json:"non_key_attributes"`
}
