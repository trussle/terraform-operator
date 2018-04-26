
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDmsEndpoint describes a AwsDmsEndpoint resource
type AwsDmsEndpoint struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDmsEndpointSpec	`json:"spec"`
}


// AwsDmsEndpointSpec is the spec for a AwsDmsEndpoint Resource
type AwsDmsEndpointSpec struct {
	EndpointId	string	`json:"endpoint_id"`
	Password	string	`json:"password"`
	EngineName	string	`json:"engine_name"`
	Username	string	`json:"username"`
	DatabaseName	string	`json:"database_name"`
	ServerName	string	`json:"server_name"`
	ServiceAccessRole	string	`json:"service_access_role"`
	EndpointType	string	`json:"endpoint_type"`
	Port	int	`json:"port"`
	Tags	map[string]???	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDmsEndpointList is a list of AwsDmsEndpoint resources
type AwsDmsEndpointList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDmsEndpoint	`json:"items"`
}

