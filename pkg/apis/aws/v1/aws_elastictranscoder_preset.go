
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElastictranscoderPreset describes a AwsElastictranscoderPreset resource
type AwsElastictranscoderPreset struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElastictranscoderPresetSpec	`json:"spec"`
}


// AwsElastictranscoderPresetSpec is the spec for a AwsElastictranscoderPreset Resource
type AwsElastictranscoderPresetSpec struct {
	Video	string	`json:"video"`
	VideoWatermarks	string	`json:"video_watermarks"`
	Audio	string	`json:"audio"`
	AudioCodecOptions	string	`json:"audio_codec_options"`
	Description	string	`json:"description"`
	Thumbnails	string	`json:"thumbnails"`
	VideoCodecOptions	map[string]interface{}	`json:"video_codec_options"`
	Container	string	`json:"container"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElastictranscoderPresetList is a list of AwsElastictranscoderPreset resources
type AwsElastictranscoderPresetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElastictranscoderPreset	`json:"items"`
}

