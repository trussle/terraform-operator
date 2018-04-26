
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
	Description	string	`json:"description"`
	Thumbnails	Generic	`json:"thumbnails"`
	Video	Generic	`json:"video"`
	VideoCodecOptions	map[string]Generic	`json:"video_codec_options"`
	Container	string	`json:"container"`
	Audio	Generic	`json:"audio"`
	AudioCodecOptions	Generic	`json:"audio_codec_options"`
	VideoWatermarks	Generic	`json:"video_watermarks"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElastictranscoderPresetList is a list of AwsElastictranscoderPreset resources
type AwsElastictranscoderPresetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElastictranscoderPreset	`json:"items"`
}

