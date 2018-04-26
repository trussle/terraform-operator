
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGlueConnection describes a AwsGlueConnection resource
type AwsGlueConnection struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGlueConnectionSpec	`json:"spec"`
}


// AwsGlueConnectionSpec is the spec for a AwsGlueConnection Resource
type AwsGlueConnectionSpec struct {
	Description	string	`json:"description"`
	MatchCriteria	[]string	`json:"match_criteria"`
	Name	string	`json:"name"`
	PhysicalConnectionRequirements	[]JYWRncGK	`json:"physical_connection_requirements"`
	ConnectionProperties	map[string]???	`json:"connection_properties"`
	ConnectionType	string	`json:"connection_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGlueConnectionList is a list of AwsGlueConnection resources
type AwsGlueConnectionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGlueConnection	`json:"items"`
}


// JYWRncGK is a JYWRncGK
type JYWRncGK struct {
	SecurityGroupIdList	[]string	`json:"security_group_id_list"`
	SubnetId	string	`json:"subnet_id"`
}
