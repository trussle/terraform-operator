
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsEndpoint struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDmsEndpointSpec	`json:"spec"`
}

type AwsDmsEndpointList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDmsEndpoint	`json:"items"`
}

type AwsDmsEndpointSpec struct {
	Password	string	`json:"password"`
	Tags	map[string]interface{}	`json:"tags"`
	Port	int	`json:"port"`
	DatabaseName	string	`json:"database_name"`
	EndpointId	string	`json:"endpoint_id"`
	ServiceAccessRole	string	`json:"service_access_role"`
	EngineName	string	`json:"engine_name"`
	EndpointType	string	`json:"endpoint_type"`
	ServerName	string	`json:"server_name"`
	Username	string	`json:"username"`
}
