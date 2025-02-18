


package videoanalyzer

type AccessPolicyEccAlgo string

const (
	// ES265
	AccessPolicyEccAlgoES256 = AccessPolicyEccAlgo("ES256")
	// ES384
	AccessPolicyEccAlgoES384 = AccessPolicyEccAlgo("ES384")
	// ES512
	AccessPolicyEccAlgoES512 = AccessPolicyEccAlgo("ES512")
)

type AccessPolicyRole string

const (
	// Reader role allows for read-only operations to be performed through the client APIs.
	AccessPolicyRoleReader = AccessPolicyRole("Reader")
)

type AccessPolicyRsaAlgo string

const (
	// RS256
	AccessPolicyRsaAlgoRS256 = AccessPolicyRsaAlgo("RS256")
	// RS384
	AccessPolicyRsaAlgoRS384 = AccessPolicyRsaAlgo("RS384")
	// RS512
	AccessPolicyRsaAlgoRS512 = AccessPolicyRsaAlgo("RS512")
)

type AccountEncryptionKeyType string

const (
	// The Account Key is encrypted with a System Key.
	AccountEncryptionKeyTypeSystemKey = AccountEncryptionKeyType("SystemKey")
	// The Account Key is encrypted with a Customer Key.
	AccountEncryptionKeyTypeCustomerKey = AccountEncryptionKeyType("CustomerKey")
)

type EncoderSystemPresetType string

const (
	// Produces an MP4 file where the video is encoded with H.264 codec at a picture height of 540 pixels, and at a maximum bitrate of 2000 Kbps. Encoded video has the same average frame rate as the input. The aspect ratio of the input is preserved. If the input content has audio, then it is encoded with AAC-LC codec at 96 Kbps
	EncoderSystemPresetType_SingleLayer_540p_H264_AAC = EncoderSystemPresetType("SingleLayer_540p_H264_AAC")
	// Produces an MP4 file where the video is encoded with H.264 codec at a picture height of 720 pixels, and at a maximum bitrate of 3500 Kbps. Encoded video has the same average frame rate as the input. The aspect ratio of the input is preserved. If the input content has audio, then it is encoded with AAC-LC codec at 96 Kbps
	EncoderSystemPresetType_SingleLayer_720p_H264_AAC = EncoderSystemPresetType("SingleLayer_720p_H264_AAC")
	// Produces an MP4 file where the video is encoded with H.264 codec at a picture height of 1080 pixels, and at a maximum bitrate of 6000 Kbps. Encoded video has the same average frame rate as the input. The aspect ratio of the input is preserved. If the input content has audio, then it is encoded with AAC-LC codec at 128 Kbps
	EncoderSystemPresetType_SingleLayer_1080p_H264_AAC = EncoderSystemPresetType("SingleLayer_1080p_H264_AAC")
	// Produces an MP4 file where the video is encoded with H.264 codec at a picture height of 2160 pixels, and at a maximum bitrate of 16000 Kbps. Encoded video has the same average frame rate as the input. The aspect ratio of the input is preserved. If the input content has audio, then it is encoded with AAC-LC codec at 128 Kbps
	EncoderSystemPresetType_SingleLayer_2160p_H264_AAC = EncoderSystemPresetType("SingleLayer_2160p_H264_AAC")
)

type Kind string

const (
	// Live pipeline topology resource.
	KindLive = Kind("Live")
	// Batch pipeline topology resource.
	KindBatch = Kind("Batch")
)

type ParameterType string

const (
	// The parameter's value is a string.
	ParameterTypeString = ParameterType("String")
	// The parameter's value is a string that holds sensitive information.
	ParameterTypeSecretString = ParameterType("SecretString")
	// The parameter's value is a 32-bit signed integer.
	ParameterTypeInt = ParameterType("Int")
	// The parameter's value is a 64-bit double-precision floating point.
	ParameterTypeDouble = ParameterType("Double")
	// The parameter's value is a boolean value that is either true or false.
	ParameterTypeBool = ParameterType("Bool")
)

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending  = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected = PrivateEndpointServiceConnectionStatus("Rejected")
)

type RtspTransport string

const (
	// HTTP transport. RTSP messages are exchanged over long running HTTP requests and RTP packets are interleaved within the HTTP channel.
	RtspTransportHttp = RtspTransport("Http")
	// TCP transport. RTSP is used directly over TCP and RTP packets are interleaved within the TCP channel.
	RtspTransportTcp = RtspTransport("Tcp")
)

type SkuName string

const (
	// Represents the Live S1 SKU name. Using this SKU you can create live pipelines to capture, record, and stream live video from RTSP-capable cameras at bitrate settings from 0.5 Kbps to 3000 Kbps.
	SkuName_Live_S1 = SkuName("Live_S1")
	// Represents the Batch S1 SKU name. Using this SKU you can create pipeline jobs to process recorded content.
	SkuName_Batch_S1 = SkuName("Batch_S1")
)

type VideoScaleMode string

const (
	// Pads the video with black horizontal stripes (letterbox) or black vertical stripes (pillar-box) so the video is resized to the specified dimensions while not altering the content aspect ratio.
	VideoScaleModePad = VideoScaleMode("Pad")
	// Preserves the same aspect ratio as the input video. If only one video dimension is provided, the second dimension is calculated based on the input video aspect ratio. When 2 dimensions are provided, the video is resized to fit the most constraining dimension, considering the input video size and aspect ratio.
	VideoScaleModePreserveAspectRatio = VideoScaleMode("PreserveAspectRatio")
	// Stretches the original video so it resized to the specified dimensions.
	VideoScaleModeStretch = VideoScaleMode("Stretch")
)

func init() {
}
