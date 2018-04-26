
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGlueCatalogDatabase describes a AwsGlueCatalogDatabase resource
type AwsGlueCatalogDatabase struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGlueCatalogDatabaseSpec	`json:"spec"`
}


// AwsGlueCatalogDatabaseSpec is the spec for a AwsGlueCatalogDatabase Resource
type AwsGlueCatalogDatabaseSpec struct {
	LocationUri	string	`json:"location_uri"`
	Parameters	map[string]Generic	`json:"parameters"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGlueCatalogDatabaseList is a list of AwsGlueCatalogDatabase resources
type AwsGlueCatalogDatabaseList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGlueCatalogDatabase	`json:"items"`
}

