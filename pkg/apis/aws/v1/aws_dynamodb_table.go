
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
	HashKey	string	`json:"hash_key"`
	RangeKey	string	`json:"range_key"`
	StreamEnabled	bool	`json:"stream_enabled"`
	WriteCapacity	int	`json:"write_capacity"`
	Attribute	string	`json:"attribute"`
	LocalSecondaryIndex	string	`json:"local_secondary_index"`
	ReadCapacity	int	`json:"read_capacity"`
	GlobalSecondaryIndex	string	`json:"global_secondary_index"`
	Tags	map[string]???	`json:"tags"`
	Ttl	string	`json:"ttl"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDynamodbTableList is a list of AwsDynamodbTable resources
type AwsDynamodbTableList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDynamodbTable	`json:"items"`
}


// mHEfnwcN is a mHEfnwcN
type mHEfnwcN struct {
	Enabled	bool	`json:"enabled"`
}
