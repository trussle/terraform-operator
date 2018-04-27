
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
	VideoWatermarks	AwsElastictranscoderPresetVideoWatermarks	`json:"video_watermarks"`
	VideoCodecOptions	map[string]string	`json:"video_codec_options"`
	Description	string	`json:"description"`
	Thumbnails	AwsElastictranscoderPresetThumbnails	`json:"thumbnails"`
	Audio	AwsElastictranscoderPresetAudio	`json:"audio"`
	AudioCodecOptions	AwsElastictranscoderPresetAudioCodecOptions	`json:"audio_codec_options"`
	Container	string	`json:"container"`
	Video	AwsElastictranscoderPresetVideo	`json:"video"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElastictranscoderPresetList is a list of AwsElastictranscoderPreset resources
type AwsElastictranscoderPresetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElastictranscoderPreset	`json:"items"`
}


// AwsElastictranscoderPresetVideoWatermarks is a AwsElastictranscoderPresetVideoWatermarks
type AwsElastictranscoderPresetVideoWatermarks struct {
	VerticalAlign	string	`json:"vertical_align"`
	VerticalOffset	string	`json:"vertical_offset"`
	HorizontalAlign	string	`json:"horizontal_align"`
	Id	string	`json:"id"`
	Opacity	string	`json:"opacity"`
	Target	string	`json:"target"`
	HorizontalOffset	string	`json:"horizontal_offset"`
	MaxHeight	string	`json:"max_height"`
	MaxWidth	string	`json:"max_width"`
	SizingPolicy	string	`json:"sizing_policy"`
}

// AwsElastictranscoderPresetThumbnails is a AwsElastictranscoderPresetThumbnails
type AwsElastictranscoderPresetThumbnails struct {
	MaxHeight	string	`json:"max_height"`
	MaxWidth	string	`json:"max_width"`
	PaddingPolicy	string	`json:"padding_policy"`
	Resolution	string	`json:"resolution"`
	SizingPolicy	string	`json:"sizing_policy"`
	AspectRatio	string	`json:"aspect_ratio"`
	Format	string	`json:"format"`
	Interval	string	`json:"interval"`
}

// AwsElastictranscoderPresetAudio is a AwsElastictranscoderPresetAudio
type AwsElastictranscoderPresetAudio struct {
	AudioPackingMode	string	`json:"audio_packing_mode"`
	BitRate	string	`json:"bit_rate"`
	Channels	string	`json:"channels"`
	Codec	string	`json:"codec"`
	SampleRate	string	`json:"sample_rate"`
}

// AwsElastictranscoderPresetAudioCodecOptions is a AwsElastictranscoderPresetAudioCodecOptions
type AwsElastictranscoderPresetAudioCodecOptions struct {
	BitDepth	string	`json:"bit_depth"`
	BitOrder	string	`json:"bit_order"`
	Profile	string	`json:"profile"`
	Signed	string	`json:"signed"`
}

// AwsElastictranscoderPresetVideo is a AwsElastictranscoderPresetVideo
type AwsElastictranscoderPresetVideo struct {
	MaxHeight	string	`json:"max_height"`
	Resolution	string	`json:"resolution"`
	AspectRatio	string	`json:"aspect_ratio"`
	DisplayAspectRatio	string	`json:"display_aspect_ratio"`
	FixedGop	string	`json:"fixed_gop"`
	FrameRate	string	`json:"frame_rate"`
	KeyframesMaxDist	string	`json:"keyframes_max_dist"`
	MaxFrameRate	string	`json:"max_frame_rate"`
	SizingPolicy	string	`json:"sizing_policy"`
	BitRate	string	`json:"bit_rate"`
	Codec	string	`json:"codec"`
	MaxWidth	string	`json:"max_width"`
	PaddingPolicy	string	`json:"padding_policy"`
}
