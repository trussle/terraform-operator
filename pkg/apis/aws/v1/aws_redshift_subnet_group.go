
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRedshiftSubnetGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRedshiftSubnetGroupSpec	`json:"spec"`
}

type AwsRedshiftSubnetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRedshiftSubnetGroup	`json:"items"`
}

type AwsRedshiftSubnetGroupSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	SubnetIds	interface{}	`json:"subnet_ids"`
}
