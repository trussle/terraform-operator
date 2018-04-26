
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Video	Video	`json:"video"`
	Container	string	`json:"container"`
	Description	string	`json:"description"`
	Thumbnails	Thumbnails	`json:"thumbnails"`
	VideoWatermarks	VideoWatermarks	`json:"video_watermarks"`
	VideoCodecOptions	map[string]string	`json:"video_codec_options"`
	Audio	Audio	`json:"audio"`
	AudioCodecOptions	AudioCodecOptions	`json:"audio_codec_options"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElastictranscoderPresetList is a list of AwsElastictranscoderPreset resources
type AwsElastictranscoderPresetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElastictranscoderPreset	`json:"items"`
}


// Audio is a Audio
type Audio struct {
	AudioPackingMode	string	`json:"audio_packing_mode"`
	BitRate	string	`json:"bit_rate"`
	Channels	string	`json:"channels"`
	Codec	string	`json:"codec"`
	SampleRate	string	`json:"sample_rate"`
}

// AudioCodecOptions is a AudioCodecOptions
type AudioCodecOptions struct {
	Profile	string	`json:"profile"`
	Signed	string	`json:"signed"`
	BitDepth	string	`json:"bit_depth"`
	BitOrder	string	`json:"bit_order"`
}

// Video is a Video
type Video struct {
	FrameRate	string	`json:"frame_rate"`
	KeyframesMaxDist	string	`json:"keyframes_max_dist"`
	MaxFrameRate	string	`json:"max_frame_rate"`
	MaxHeight	string	`json:"max_height"`
	PaddingPolicy	string	`json:"padding_policy"`
	Resolution	string	`json:"resolution"`
	SizingPolicy	string	`json:"sizing_policy"`
	AspectRatio	string	`json:"aspect_ratio"`
	BitRate	string	`json:"bit_rate"`
	Codec	string	`json:"codec"`
	DisplayAspectRatio	string	`json:"display_aspect_ratio"`
	FixedGop	string	`json:"fixed_gop"`
	MaxWidth	string	`json:"max_width"`
}

// Thumbnails is a Thumbnails
type Thumbnails struct {
	Format	string	`json:"format"`
	Interval	string	`json:"interval"`
	MaxHeight	string	`json:"max_height"`
	MaxWidth	string	`json:"max_width"`
	PaddingPolicy	string	`json:"padding_policy"`
	Resolution	string	`json:"resolution"`
	SizingPolicy	string	`json:"sizing_policy"`
	AspectRatio	string	`json:"aspect_ratio"`
}

// VideoWatermarks is a VideoWatermarks
type VideoWatermarks struct {
	SizingPolicy	string	`json:"sizing_policy"`
	Target	string	`json:"target"`
	VerticalAlign	string	`json:"vertical_align"`
	HorizontalAlign	string	`json:"horizontal_align"`
	HorizontalOffset	string	`json:"horizontal_offset"`
	Id	string	`json:"id"`
	VerticalOffset	string	`json:"vertical_offset"`
	MaxHeight	string	`json:"max_height"`
	MaxWidth	string	`json:"max_width"`
	Opacity	string	`json:"opacity"`
}
