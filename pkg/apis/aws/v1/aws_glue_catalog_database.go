
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGlueCatalogDatabase struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGlueCatalogDatabaseSpec	`json:"spec"`
}

type AwsGlueCatalogDatabaseList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGlueCatalogDatabase	`json:"items"`
}

type AwsGlueCatalogDatabaseSpec struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	LocationUri	string	`json:"location_uri"`
	Parameters	map[string]interface{}	`json:"parameters"`
}
