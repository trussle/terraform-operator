
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmActivation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmActivationSpec	`json:"spec"`
}

type AwsSsmActivationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmActivation	`json:"items"`
}

type AwsSsmActivationSpec struct {
	IamRole	string	`json:"iam_role"`
	RegistrationLimit	int	`json:"registration_limit"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	ExpirationDate	string	`json:"expiration_date"`
}
