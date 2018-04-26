
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRedshiftSubnetGroup describes a AwsRedshiftSubnetGroup resource
type AwsRedshiftSubnetGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRedshiftSubnetGroupSpec	`json:"spec"`
}


// AwsRedshiftSubnetGroupSpec is the spec for a AwsRedshiftSubnetGroup Resource
type AwsRedshiftSubnetGroupSpec struct {
	Tags	map[string]Generic	`json:"tags"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	SubnetIds	Generic	`json:"subnet_ids"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRedshiftSubnetGroupList is a list of AwsRedshiftSubnetGroup resources
type AwsRedshiftSubnetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRedshiftSubnetGroup	`json:"items"`
}

