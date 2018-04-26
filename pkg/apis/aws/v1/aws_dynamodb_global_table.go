
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDynamodbGlobalTable describes a AwsDynamodbGlobalTable resource
type AwsDynamodbGlobalTable struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDynamodbGlobalTableSpec	`json:"spec"`
}


// AwsDynamodbGlobalTableSpec is the spec for a AwsDynamodbGlobalTable Resource
type AwsDynamodbGlobalTableSpec struct {
	Name	string	`json:"name"`
	Replica	string	`json:"replica"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDynamodbGlobalTableList is a list of AwsDynamodbGlobalTable resources
type AwsDynamodbGlobalTableList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDynamodbGlobalTable	`json:"items"`
}

