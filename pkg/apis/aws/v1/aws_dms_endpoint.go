
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	Username	string	`json:"username"`
	EndpointId	string	`json:"endpoint_id"`
	Port	int	`json:"port"`
	EngineName	string	`json:"engine_name"`
	Password	string	`json:"password"`
	EndpointType	string	`json:"endpoint_type"`
	DatabaseName	string	`json:"database_name"`
	ServiceAccessRole	string	`json:"service_access_role"`
	ServerName	string	`json:"server_name"`
	Tags	map[string]interface{}	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDmsEndpointList is a list of AwsDmsEndpoint resources
type AwsDmsEndpointList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDmsEndpoint	`json:"items"`
}

