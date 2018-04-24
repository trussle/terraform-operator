
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElastictranscoderPreset struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElastictranscoderPresetSpec	`json:"spec"`
}

type AwsElastictranscoderPresetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElastictranscoderPreset	`json:"items"`
}

type AwsElastictranscoderPresetSpec struct {
	Audio	interface{}	`json:"audio"`
	AudioCodecOptions	interface{}	`json:"audio_codec_options"`
	Container	string	`json:"container"`
	Thumbnails	interface{}	`json:"thumbnails"`
	Video	interface{}	`json:"video"`
	VideoWatermarks	interface{}	`json:"video_watermarks"`
	VideoCodecOptions	map[string]interface{}	`json:"video_codec_options"`
	Description	string	`json:"description"`
}
