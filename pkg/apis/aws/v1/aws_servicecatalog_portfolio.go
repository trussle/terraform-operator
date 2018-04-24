
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsServicecatalogPortfolio struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsServicecatalogPortfolioSpec	`json:"spec"`
}

type AwsServicecatalogPortfolioList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsServicecatalogPortfolio	`json:"items"`
}

type AwsServicecatalogPortfolioSpec struct {
	Name	string	`json:"name"`
	ProviderName	string	`json:"provider_name"`
	Tags	map[string]interface{}	`json:"tags"`
}
