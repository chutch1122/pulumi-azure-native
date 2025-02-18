


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AacAudio struct {
	Bitrate      *int    `pulumi:"bitrate"`
	Channels     *int    `pulumi:"channels"`
	Label        *string `pulumi:"label"`
	OdataType    string  `pulumi:"odataType"`
	Profile      *string `pulumi:"profile"`
	SamplingRate *int    `pulumi:"samplingRate"`
}

type AacAudioResponse struct {
	Bitrate      *int    `pulumi:"bitrate"`
	Channels     *int    `pulumi:"channels"`
	Label        *string `pulumi:"label"`
	OdataType    string  `pulumi:"odataType"`
	Profile      *string `pulumi:"profile"`
	SamplingRate *int    `pulumi:"samplingRate"`
}

type AkamaiAccessControl struct {
	AkamaiSignatureHeaderAuthenticationKeyList []AkamaiSignatureHeaderAuthenticationKey `pulumi:"akamaiSignatureHeaderAuthenticationKeyList"`
}





type AkamaiAccessControlInput interface {
	pulumi.Input

	ToAkamaiAccessControlOutput() AkamaiAccessControlOutput
	ToAkamaiAccessControlOutputWithContext(context.Context) AkamaiAccessControlOutput
}

type AkamaiAccessControlArgs struct {
	AkamaiSignatureHeaderAuthenticationKeyList AkamaiSignatureHeaderAuthenticationKeyArrayInput `pulumi:"akamaiSignatureHeaderAuthenticationKeyList"`
}

func (AkamaiAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiAccessControl)(nil)).Elem()
}

func (i AkamaiAccessControlArgs) ToAkamaiAccessControlOutput() AkamaiAccessControlOutput {
	return i.ToAkamaiAccessControlOutputWithContext(context.Background())
}

func (i AkamaiAccessControlArgs) ToAkamaiAccessControlOutputWithContext(ctx context.Context) AkamaiAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiAccessControlOutput)
}

func (i AkamaiAccessControlArgs) ToAkamaiAccessControlPtrOutput() AkamaiAccessControlPtrOutput {
	return i.ToAkamaiAccessControlPtrOutputWithContext(context.Background())
}

func (i AkamaiAccessControlArgs) ToAkamaiAccessControlPtrOutputWithContext(ctx context.Context) AkamaiAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiAccessControlOutput).ToAkamaiAccessControlPtrOutputWithContext(ctx)
}









type AkamaiAccessControlPtrInput interface {
	pulumi.Input

	ToAkamaiAccessControlPtrOutput() AkamaiAccessControlPtrOutput
	ToAkamaiAccessControlPtrOutputWithContext(context.Context) AkamaiAccessControlPtrOutput
}

type akamaiAccessControlPtrType AkamaiAccessControlArgs

func AkamaiAccessControlPtr(v *AkamaiAccessControlArgs) AkamaiAccessControlPtrInput {
	return (*akamaiAccessControlPtrType)(v)
}

func (*akamaiAccessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AkamaiAccessControl)(nil)).Elem()
}

func (i *akamaiAccessControlPtrType) ToAkamaiAccessControlPtrOutput() AkamaiAccessControlPtrOutput {
	return i.ToAkamaiAccessControlPtrOutputWithContext(context.Background())
}

func (i *akamaiAccessControlPtrType) ToAkamaiAccessControlPtrOutputWithContext(ctx context.Context) AkamaiAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiAccessControlPtrOutput)
}

type AkamaiAccessControlOutput struct{ *pulumi.OutputState }

func (AkamaiAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiAccessControl)(nil)).Elem()
}

func (o AkamaiAccessControlOutput) ToAkamaiAccessControlOutput() AkamaiAccessControlOutput {
	return o
}

func (o AkamaiAccessControlOutput) ToAkamaiAccessControlOutputWithContext(ctx context.Context) AkamaiAccessControlOutput {
	return o
}

func (o AkamaiAccessControlOutput) ToAkamaiAccessControlPtrOutput() AkamaiAccessControlPtrOutput {
	return o.ToAkamaiAccessControlPtrOutputWithContext(context.Background())
}

func (o AkamaiAccessControlOutput) ToAkamaiAccessControlPtrOutputWithContext(ctx context.Context) AkamaiAccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AkamaiAccessControl) *AkamaiAccessControl {
		return &v
	}).(AkamaiAccessControlPtrOutput)
}

func (o AkamaiAccessControlOutput) AkamaiSignatureHeaderAuthenticationKeyList() AkamaiSignatureHeaderAuthenticationKeyArrayOutput {
	return o.ApplyT(func(v AkamaiAccessControl) []AkamaiSignatureHeaderAuthenticationKey {
		return v.AkamaiSignatureHeaderAuthenticationKeyList
	}).(AkamaiSignatureHeaderAuthenticationKeyArrayOutput)
}

type AkamaiAccessControlPtrOutput struct{ *pulumi.OutputState }

func (AkamaiAccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AkamaiAccessControl)(nil)).Elem()
}

func (o AkamaiAccessControlPtrOutput) ToAkamaiAccessControlPtrOutput() AkamaiAccessControlPtrOutput {
	return o
}

func (o AkamaiAccessControlPtrOutput) ToAkamaiAccessControlPtrOutputWithContext(ctx context.Context) AkamaiAccessControlPtrOutput {
	return o
}

func (o AkamaiAccessControlPtrOutput) Elem() AkamaiAccessControlOutput {
	return o.ApplyT(func(v *AkamaiAccessControl) AkamaiAccessControl {
		if v != nil {
			return *v
		}
		var ret AkamaiAccessControl
		return ret
	}).(AkamaiAccessControlOutput)
}

func (o AkamaiAccessControlPtrOutput) AkamaiSignatureHeaderAuthenticationKeyList() AkamaiSignatureHeaderAuthenticationKeyArrayOutput {
	return o.ApplyT(func(v *AkamaiAccessControl) []AkamaiSignatureHeaderAuthenticationKey {
		if v == nil {
			return nil
		}
		return v.AkamaiSignatureHeaderAuthenticationKeyList
	}).(AkamaiSignatureHeaderAuthenticationKeyArrayOutput)
}

type AkamaiAccessControlResponse struct {
	AkamaiSignatureHeaderAuthenticationKeyList []AkamaiSignatureHeaderAuthenticationKeyResponse `pulumi:"akamaiSignatureHeaderAuthenticationKeyList"`
}

type AkamaiAccessControlResponseOutput struct{ *pulumi.OutputState }

func (AkamaiAccessControlResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiAccessControlResponse)(nil)).Elem()
}

func (o AkamaiAccessControlResponseOutput) ToAkamaiAccessControlResponseOutput() AkamaiAccessControlResponseOutput {
	return o
}

func (o AkamaiAccessControlResponseOutput) ToAkamaiAccessControlResponseOutputWithContext(ctx context.Context) AkamaiAccessControlResponseOutput {
	return o
}

func (o AkamaiAccessControlResponseOutput) AkamaiSignatureHeaderAuthenticationKeyList() AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput {
	return o.ApplyT(func(v AkamaiAccessControlResponse) []AkamaiSignatureHeaderAuthenticationKeyResponse {
		return v.AkamaiSignatureHeaderAuthenticationKeyList
	}).(AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput)
}

type AkamaiAccessControlResponsePtrOutput struct{ *pulumi.OutputState }

func (AkamaiAccessControlResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AkamaiAccessControlResponse)(nil)).Elem()
}

func (o AkamaiAccessControlResponsePtrOutput) ToAkamaiAccessControlResponsePtrOutput() AkamaiAccessControlResponsePtrOutput {
	return o
}

func (o AkamaiAccessControlResponsePtrOutput) ToAkamaiAccessControlResponsePtrOutputWithContext(ctx context.Context) AkamaiAccessControlResponsePtrOutput {
	return o
}

func (o AkamaiAccessControlResponsePtrOutput) Elem() AkamaiAccessControlResponseOutput {
	return o.ApplyT(func(v *AkamaiAccessControlResponse) AkamaiAccessControlResponse {
		if v != nil {
			return *v
		}
		var ret AkamaiAccessControlResponse
		return ret
	}).(AkamaiAccessControlResponseOutput)
}

func (o AkamaiAccessControlResponsePtrOutput) AkamaiSignatureHeaderAuthenticationKeyList() AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput {
	return o.ApplyT(func(v *AkamaiAccessControlResponse) []AkamaiSignatureHeaderAuthenticationKeyResponse {
		if v == nil {
			return nil
		}
		return v.AkamaiSignatureHeaderAuthenticationKeyList
	}).(AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput)
}

type AkamaiSignatureHeaderAuthenticationKey struct {
	Base64Key  *string `pulumi:"base64Key"`
	Expiration *string `pulumi:"expiration"`
	Identifier *string `pulumi:"identifier"`
}





type AkamaiSignatureHeaderAuthenticationKeyInput interface {
	pulumi.Input

	ToAkamaiSignatureHeaderAuthenticationKeyOutput() AkamaiSignatureHeaderAuthenticationKeyOutput
	ToAkamaiSignatureHeaderAuthenticationKeyOutputWithContext(context.Context) AkamaiSignatureHeaderAuthenticationKeyOutput
}

type AkamaiSignatureHeaderAuthenticationKeyArgs struct {
	Base64Key  pulumi.StringPtrInput `pulumi:"base64Key"`
	Expiration pulumi.StringPtrInput `pulumi:"expiration"`
	Identifier pulumi.StringPtrInput `pulumi:"identifier"`
}

func (AkamaiSignatureHeaderAuthenticationKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKey)(nil)).Elem()
}

func (i AkamaiSignatureHeaderAuthenticationKeyArgs) ToAkamaiSignatureHeaderAuthenticationKeyOutput() AkamaiSignatureHeaderAuthenticationKeyOutput {
	return i.ToAkamaiSignatureHeaderAuthenticationKeyOutputWithContext(context.Background())
}

func (i AkamaiSignatureHeaderAuthenticationKeyArgs) ToAkamaiSignatureHeaderAuthenticationKeyOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiSignatureHeaderAuthenticationKeyOutput)
}





type AkamaiSignatureHeaderAuthenticationKeyArrayInput interface {
	pulumi.Input

	ToAkamaiSignatureHeaderAuthenticationKeyArrayOutput() AkamaiSignatureHeaderAuthenticationKeyArrayOutput
	ToAkamaiSignatureHeaderAuthenticationKeyArrayOutputWithContext(context.Context) AkamaiSignatureHeaderAuthenticationKeyArrayOutput
}

type AkamaiSignatureHeaderAuthenticationKeyArray []AkamaiSignatureHeaderAuthenticationKeyInput

func (AkamaiSignatureHeaderAuthenticationKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AkamaiSignatureHeaderAuthenticationKey)(nil)).Elem()
}

func (i AkamaiSignatureHeaderAuthenticationKeyArray) ToAkamaiSignatureHeaderAuthenticationKeyArrayOutput() AkamaiSignatureHeaderAuthenticationKeyArrayOutput {
	return i.ToAkamaiSignatureHeaderAuthenticationKeyArrayOutputWithContext(context.Background())
}

func (i AkamaiSignatureHeaderAuthenticationKeyArray) ToAkamaiSignatureHeaderAuthenticationKeyArrayOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AkamaiSignatureHeaderAuthenticationKeyArrayOutput)
}

type AkamaiSignatureHeaderAuthenticationKeyOutput struct{ *pulumi.OutputState }

func (AkamaiSignatureHeaderAuthenticationKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKey)(nil)).Elem()
}

func (o AkamaiSignatureHeaderAuthenticationKeyOutput) ToAkamaiSignatureHeaderAuthenticationKeyOutput() AkamaiSignatureHeaderAuthenticationKeyOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyOutput) ToAkamaiSignatureHeaderAuthenticationKeyOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyOutput) Base64Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AkamaiSignatureHeaderAuthenticationKey) *string { return v.Base64Key }).(pulumi.StringPtrOutput)
}

func (o AkamaiSignatureHeaderAuthenticationKeyOutput) Expiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AkamaiSignatureHeaderAuthenticationKey) *string { return v.Expiration }).(pulumi.StringPtrOutput)
}

func (o AkamaiSignatureHeaderAuthenticationKeyOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AkamaiSignatureHeaderAuthenticationKey) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

type AkamaiSignatureHeaderAuthenticationKeyArrayOutput struct{ *pulumi.OutputState }

func (AkamaiSignatureHeaderAuthenticationKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AkamaiSignatureHeaderAuthenticationKey)(nil)).Elem()
}

func (o AkamaiSignatureHeaderAuthenticationKeyArrayOutput) ToAkamaiSignatureHeaderAuthenticationKeyArrayOutput() AkamaiSignatureHeaderAuthenticationKeyArrayOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyArrayOutput) ToAkamaiSignatureHeaderAuthenticationKeyArrayOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyArrayOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyArrayOutput) Index(i pulumi.IntInput) AkamaiSignatureHeaderAuthenticationKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AkamaiSignatureHeaderAuthenticationKey {
		return vs[0].([]AkamaiSignatureHeaderAuthenticationKey)[vs[1].(int)]
	}).(AkamaiSignatureHeaderAuthenticationKeyOutput)
}

type AkamaiSignatureHeaderAuthenticationKeyResponse struct {
	Base64Key  *string `pulumi:"base64Key"`
	Expiration *string `pulumi:"expiration"`
	Identifier *string `pulumi:"identifier"`
}

type AkamaiSignatureHeaderAuthenticationKeyResponseOutput struct{ *pulumi.OutputState }

func (AkamaiSignatureHeaderAuthenticationKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AkamaiSignatureHeaderAuthenticationKeyResponse)(nil)).Elem()
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseOutput) ToAkamaiSignatureHeaderAuthenticationKeyResponseOutput() AkamaiSignatureHeaderAuthenticationKeyResponseOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseOutput) ToAkamaiSignatureHeaderAuthenticationKeyResponseOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyResponseOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseOutput) Base64Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AkamaiSignatureHeaderAuthenticationKeyResponse) *string { return v.Base64Key }).(pulumi.StringPtrOutput)
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseOutput) Expiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AkamaiSignatureHeaderAuthenticationKeyResponse) *string { return v.Expiration }).(pulumi.StringPtrOutput)
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AkamaiSignatureHeaderAuthenticationKeyResponse) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

type AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AkamaiSignatureHeaderAuthenticationKeyResponse)(nil)).Elem()
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput) ToAkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput() AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput) ToAkamaiSignatureHeaderAuthenticationKeyResponseArrayOutputWithContext(ctx context.Context) AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput {
	return o
}

func (o AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput) Index(i pulumi.IntInput) AkamaiSignatureHeaderAuthenticationKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AkamaiSignatureHeaderAuthenticationKeyResponse {
		return vs[0].([]AkamaiSignatureHeaderAuthenticationKeyResponse)[vs[1].(int)]
	}).(AkamaiSignatureHeaderAuthenticationKeyResponseOutput)
}

type Audio struct {
	Bitrate      *int    `pulumi:"bitrate"`
	Channels     *int    `pulumi:"channels"`
	Label        *string `pulumi:"label"`
	OdataType    string  `pulumi:"odataType"`
	SamplingRate *int    `pulumi:"samplingRate"`
}

type AudioAnalyzerPreset struct {
	AudioLanguage *string `pulumi:"audioLanguage"`
	OdataType     string  `pulumi:"odataType"`
}

type AudioAnalyzerPresetResponse struct {
	AudioLanguage *string `pulumi:"audioLanguage"`
	OdataType     string  `pulumi:"odataType"`
}

type AudioOverlay struct {
	AudioGainLevel  *float64 `pulumi:"audioGainLevel"`
	End             *string  `pulumi:"end"`
	FadeInDuration  *string  `pulumi:"fadeInDuration"`
	FadeOutDuration *string  `pulumi:"fadeOutDuration"`
	InputLabel      *string  `pulumi:"inputLabel"`
	OdataType       string   `pulumi:"odataType"`
	Start           *string  `pulumi:"start"`
}

type AudioOverlayResponse struct {
	AudioGainLevel  *float64 `pulumi:"audioGainLevel"`
	End             *string  `pulumi:"end"`
	FadeInDuration  *string  `pulumi:"fadeInDuration"`
	FadeOutDuration *string  `pulumi:"fadeOutDuration"`
	InputLabel      *string  `pulumi:"inputLabel"`
	OdataType       string   `pulumi:"odataType"`
	Start           *string  `pulumi:"start"`
}

type AudioResponse struct {
	Bitrate      *int    `pulumi:"bitrate"`
	Channels     *int    `pulumi:"channels"`
	Label        *string `pulumi:"label"`
	OdataType    string  `pulumi:"odataType"`
	SamplingRate *int    `pulumi:"samplingRate"`
}

type BuiltInStandardEncoderPreset struct {
	OdataType  string `pulumi:"odataType"`
	PresetName string `pulumi:"presetName"`
}

type BuiltInStandardEncoderPresetResponse struct {
	OdataType  string `pulumi:"odataType"`
	PresetName string `pulumi:"presetName"`
}

type CbcsDrmConfiguration struct {
	FairPlay  *StreamingPolicyFairPlayConfiguration  `pulumi:"fairPlay"`
	PlayReady *StreamingPolicyPlayReadyConfiguration `pulumi:"playReady"`
	Widevine  *StreamingPolicyWidevineConfiguration  `pulumi:"widevine"`
}





type CbcsDrmConfigurationInput interface {
	pulumi.Input

	ToCbcsDrmConfigurationOutput() CbcsDrmConfigurationOutput
	ToCbcsDrmConfigurationOutputWithContext(context.Context) CbcsDrmConfigurationOutput
}

type CbcsDrmConfigurationArgs struct {
	FairPlay  StreamingPolicyFairPlayConfigurationPtrInput  `pulumi:"fairPlay"`
	PlayReady StreamingPolicyPlayReadyConfigurationPtrInput `pulumi:"playReady"`
	Widevine  StreamingPolicyWidevineConfigurationPtrInput  `pulumi:"widevine"`
}

func (CbcsDrmConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CbcsDrmConfiguration)(nil)).Elem()
}

func (i CbcsDrmConfigurationArgs) ToCbcsDrmConfigurationOutput() CbcsDrmConfigurationOutput {
	return i.ToCbcsDrmConfigurationOutputWithContext(context.Background())
}

func (i CbcsDrmConfigurationArgs) ToCbcsDrmConfigurationOutputWithContext(ctx context.Context) CbcsDrmConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CbcsDrmConfigurationOutput)
}

func (i CbcsDrmConfigurationArgs) ToCbcsDrmConfigurationPtrOutput() CbcsDrmConfigurationPtrOutput {
	return i.ToCbcsDrmConfigurationPtrOutputWithContext(context.Background())
}

func (i CbcsDrmConfigurationArgs) ToCbcsDrmConfigurationPtrOutputWithContext(ctx context.Context) CbcsDrmConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CbcsDrmConfigurationOutput).ToCbcsDrmConfigurationPtrOutputWithContext(ctx)
}









type CbcsDrmConfigurationPtrInput interface {
	pulumi.Input

	ToCbcsDrmConfigurationPtrOutput() CbcsDrmConfigurationPtrOutput
	ToCbcsDrmConfigurationPtrOutputWithContext(context.Context) CbcsDrmConfigurationPtrOutput
}

type cbcsDrmConfigurationPtrType CbcsDrmConfigurationArgs

func CbcsDrmConfigurationPtr(v *CbcsDrmConfigurationArgs) CbcsDrmConfigurationPtrInput {
	return (*cbcsDrmConfigurationPtrType)(v)
}

func (*cbcsDrmConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CbcsDrmConfiguration)(nil)).Elem()
}

func (i *cbcsDrmConfigurationPtrType) ToCbcsDrmConfigurationPtrOutput() CbcsDrmConfigurationPtrOutput {
	return i.ToCbcsDrmConfigurationPtrOutputWithContext(context.Background())
}

func (i *cbcsDrmConfigurationPtrType) ToCbcsDrmConfigurationPtrOutputWithContext(ctx context.Context) CbcsDrmConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CbcsDrmConfigurationPtrOutput)
}

type CbcsDrmConfigurationOutput struct{ *pulumi.OutputState }

func (CbcsDrmConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CbcsDrmConfiguration)(nil)).Elem()
}

func (o CbcsDrmConfigurationOutput) ToCbcsDrmConfigurationOutput() CbcsDrmConfigurationOutput {
	return o
}

func (o CbcsDrmConfigurationOutput) ToCbcsDrmConfigurationOutputWithContext(ctx context.Context) CbcsDrmConfigurationOutput {
	return o
}

func (o CbcsDrmConfigurationOutput) ToCbcsDrmConfigurationPtrOutput() CbcsDrmConfigurationPtrOutput {
	return o.ToCbcsDrmConfigurationPtrOutputWithContext(context.Background())
}

func (o CbcsDrmConfigurationOutput) ToCbcsDrmConfigurationPtrOutputWithContext(ctx context.Context) CbcsDrmConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CbcsDrmConfiguration) *CbcsDrmConfiguration {
		return &v
	}).(CbcsDrmConfigurationPtrOutput)
}

func (o CbcsDrmConfigurationOutput) FairPlay() StreamingPolicyFairPlayConfigurationPtrOutput {
	return o.ApplyT(func(v CbcsDrmConfiguration) *StreamingPolicyFairPlayConfiguration { return v.FairPlay }).(StreamingPolicyFairPlayConfigurationPtrOutput)
}

func (o CbcsDrmConfigurationOutput) PlayReady() StreamingPolicyPlayReadyConfigurationPtrOutput {
	return o.ApplyT(func(v CbcsDrmConfiguration) *StreamingPolicyPlayReadyConfiguration { return v.PlayReady }).(StreamingPolicyPlayReadyConfigurationPtrOutput)
}

func (o CbcsDrmConfigurationOutput) Widevine() StreamingPolicyWidevineConfigurationPtrOutput {
	return o.ApplyT(func(v CbcsDrmConfiguration) *StreamingPolicyWidevineConfiguration { return v.Widevine }).(StreamingPolicyWidevineConfigurationPtrOutput)
}

type CbcsDrmConfigurationPtrOutput struct{ *pulumi.OutputState }

func (CbcsDrmConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CbcsDrmConfiguration)(nil)).Elem()
}

func (o CbcsDrmConfigurationPtrOutput) ToCbcsDrmConfigurationPtrOutput() CbcsDrmConfigurationPtrOutput {
	return o
}

func (o CbcsDrmConfigurationPtrOutput) ToCbcsDrmConfigurationPtrOutputWithContext(ctx context.Context) CbcsDrmConfigurationPtrOutput {
	return o
}

func (o CbcsDrmConfigurationPtrOutput) Elem() CbcsDrmConfigurationOutput {
	return o.ApplyT(func(v *CbcsDrmConfiguration) CbcsDrmConfiguration {
		if v != nil {
			return *v
		}
		var ret CbcsDrmConfiguration
		return ret
	}).(CbcsDrmConfigurationOutput)
}

func (o CbcsDrmConfigurationPtrOutput) FairPlay() StreamingPolicyFairPlayConfigurationPtrOutput {
	return o.ApplyT(func(v *CbcsDrmConfiguration) *StreamingPolicyFairPlayConfiguration {
		if v == nil {
			return nil
		}
		return v.FairPlay
	}).(StreamingPolicyFairPlayConfigurationPtrOutput)
}

func (o CbcsDrmConfigurationPtrOutput) PlayReady() StreamingPolicyPlayReadyConfigurationPtrOutput {
	return o.ApplyT(func(v *CbcsDrmConfiguration) *StreamingPolicyPlayReadyConfiguration {
		if v == nil {
			return nil
		}
		return v.PlayReady
	}).(StreamingPolicyPlayReadyConfigurationPtrOutput)
}

func (o CbcsDrmConfigurationPtrOutput) Widevine() StreamingPolicyWidevineConfigurationPtrOutput {
	return o.ApplyT(func(v *CbcsDrmConfiguration) *StreamingPolicyWidevineConfiguration {
		if v == nil {
			return nil
		}
		return v.Widevine
	}).(StreamingPolicyWidevineConfigurationPtrOutput)
}

type CbcsDrmConfigurationResponse struct {
	FairPlay  *StreamingPolicyFairPlayConfigurationResponse  `pulumi:"fairPlay"`
	PlayReady *StreamingPolicyPlayReadyConfigurationResponse `pulumi:"playReady"`
	Widevine  *StreamingPolicyWidevineConfigurationResponse  `pulumi:"widevine"`
}

type CbcsDrmConfigurationResponseOutput struct{ *pulumi.OutputState }

func (CbcsDrmConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CbcsDrmConfigurationResponse)(nil)).Elem()
}

func (o CbcsDrmConfigurationResponseOutput) ToCbcsDrmConfigurationResponseOutput() CbcsDrmConfigurationResponseOutput {
	return o
}

func (o CbcsDrmConfigurationResponseOutput) ToCbcsDrmConfigurationResponseOutputWithContext(ctx context.Context) CbcsDrmConfigurationResponseOutput {
	return o
}

func (o CbcsDrmConfigurationResponseOutput) FairPlay() StreamingPolicyFairPlayConfigurationResponsePtrOutput {
	return o.ApplyT(func(v CbcsDrmConfigurationResponse) *StreamingPolicyFairPlayConfigurationResponse { return v.FairPlay }).(StreamingPolicyFairPlayConfigurationResponsePtrOutput)
}

func (o CbcsDrmConfigurationResponseOutput) PlayReady() StreamingPolicyPlayReadyConfigurationResponsePtrOutput {
	return o.ApplyT(func(v CbcsDrmConfigurationResponse) *StreamingPolicyPlayReadyConfigurationResponse {
		return v.PlayReady
	}).(StreamingPolicyPlayReadyConfigurationResponsePtrOutput)
}

func (o CbcsDrmConfigurationResponseOutput) Widevine() StreamingPolicyWidevineConfigurationResponsePtrOutput {
	return o.ApplyT(func(v CbcsDrmConfigurationResponse) *StreamingPolicyWidevineConfigurationResponse { return v.Widevine }).(StreamingPolicyWidevineConfigurationResponsePtrOutput)
}

type CbcsDrmConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (CbcsDrmConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CbcsDrmConfigurationResponse)(nil)).Elem()
}

func (o CbcsDrmConfigurationResponsePtrOutput) ToCbcsDrmConfigurationResponsePtrOutput() CbcsDrmConfigurationResponsePtrOutput {
	return o
}

func (o CbcsDrmConfigurationResponsePtrOutput) ToCbcsDrmConfigurationResponsePtrOutputWithContext(ctx context.Context) CbcsDrmConfigurationResponsePtrOutput {
	return o
}

func (o CbcsDrmConfigurationResponsePtrOutput) Elem() CbcsDrmConfigurationResponseOutput {
	return o.ApplyT(func(v *CbcsDrmConfigurationResponse) CbcsDrmConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret CbcsDrmConfigurationResponse
		return ret
	}).(CbcsDrmConfigurationResponseOutput)
}

func (o CbcsDrmConfigurationResponsePtrOutput) FairPlay() StreamingPolicyFairPlayConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *CbcsDrmConfigurationResponse) *StreamingPolicyFairPlayConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.FairPlay
	}).(StreamingPolicyFairPlayConfigurationResponsePtrOutput)
}

func (o CbcsDrmConfigurationResponsePtrOutput) PlayReady() StreamingPolicyPlayReadyConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *CbcsDrmConfigurationResponse) *StreamingPolicyPlayReadyConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.PlayReady
	}).(StreamingPolicyPlayReadyConfigurationResponsePtrOutput)
}

func (o CbcsDrmConfigurationResponsePtrOutput) Widevine() StreamingPolicyWidevineConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *CbcsDrmConfigurationResponse) *StreamingPolicyWidevineConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Widevine
	}).(StreamingPolicyWidevineConfigurationResponsePtrOutput)
}

type CencDrmConfiguration struct {
	PlayReady *StreamingPolicyPlayReadyConfiguration `pulumi:"playReady"`
	Widevine  *StreamingPolicyWidevineConfiguration  `pulumi:"widevine"`
}





type CencDrmConfigurationInput interface {
	pulumi.Input

	ToCencDrmConfigurationOutput() CencDrmConfigurationOutput
	ToCencDrmConfigurationOutputWithContext(context.Context) CencDrmConfigurationOutput
}

type CencDrmConfigurationArgs struct {
	PlayReady StreamingPolicyPlayReadyConfigurationPtrInput `pulumi:"playReady"`
	Widevine  StreamingPolicyWidevineConfigurationPtrInput  `pulumi:"widevine"`
}

func (CencDrmConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CencDrmConfiguration)(nil)).Elem()
}

func (i CencDrmConfigurationArgs) ToCencDrmConfigurationOutput() CencDrmConfigurationOutput {
	return i.ToCencDrmConfigurationOutputWithContext(context.Background())
}

func (i CencDrmConfigurationArgs) ToCencDrmConfigurationOutputWithContext(ctx context.Context) CencDrmConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CencDrmConfigurationOutput)
}

func (i CencDrmConfigurationArgs) ToCencDrmConfigurationPtrOutput() CencDrmConfigurationPtrOutput {
	return i.ToCencDrmConfigurationPtrOutputWithContext(context.Background())
}

func (i CencDrmConfigurationArgs) ToCencDrmConfigurationPtrOutputWithContext(ctx context.Context) CencDrmConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CencDrmConfigurationOutput).ToCencDrmConfigurationPtrOutputWithContext(ctx)
}









type CencDrmConfigurationPtrInput interface {
	pulumi.Input

	ToCencDrmConfigurationPtrOutput() CencDrmConfigurationPtrOutput
	ToCencDrmConfigurationPtrOutputWithContext(context.Context) CencDrmConfigurationPtrOutput
}

type cencDrmConfigurationPtrType CencDrmConfigurationArgs

func CencDrmConfigurationPtr(v *CencDrmConfigurationArgs) CencDrmConfigurationPtrInput {
	return (*cencDrmConfigurationPtrType)(v)
}

func (*cencDrmConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CencDrmConfiguration)(nil)).Elem()
}

func (i *cencDrmConfigurationPtrType) ToCencDrmConfigurationPtrOutput() CencDrmConfigurationPtrOutput {
	return i.ToCencDrmConfigurationPtrOutputWithContext(context.Background())
}

func (i *cencDrmConfigurationPtrType) ToCencDrmConfigurationPtrOutputWithContext(ctx context.Context) CencDrmConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CencDrmConfigurationPtrOutput)
}

type CencDrmConfigurationOutput struct{ *pulumi.OutputState }

func (CencDrmConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CencDrmConfiguration)(nil)).Elem()
}

func (o CencDrmConfigurationOutput) ToCencDrmConfigurationOutput() CencDrmConfigurationOutput {
	return o
}

func (o CencDrmConfigurationOutput) ToCencDrmConfigurationOutputWithContext(ctx context.Context) CencDrmConfigurationOutput {
	return o
}

func (o CencDrmConfigurationOutput) ToCencDrmConfigurationPtrOutput() CencDrmConfigurationPtrOutput {
	return o.ToCencDrmConfigurationPtrOutputWithContext(context.Background())
}

func (o CencDrmConfigurationOutput) ToCencDrmConfigurationPtrOutputWithContext(ctx context.Context) CencDrmConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CencDrmConfiguration) *CencDrmConfiguration {
		return &v
	}).(CencDrmConfigurationPtrOutput)
}

func (o CencDrmConfigurationOutput) PlayReady() StreamingPolicyPlayReadyConfigurationPtrOutput {
	return o.ApplyT(func(v CencDrmConfiguration) *StreamingPolicyPlayReadyConfiguration { return v.PlayReady }).(StreamingPolicyPlayReadyConfigurationPtrOutput)
}

func (o CencDrmConfigurationOutput) Widevine() StreamingPolicyWidevineConfigurationPtrOutput {
	return o.ApplyT(func(v CencDrmConfiguration) *StreamingPolicyWidevineConfiguration { return v.Widevine }).(StreamingPolicyWidevineConfigurationPtrOutput)
}

type CencDrmConfigurationPtrOutput struct{ *pulumi.OutputState }

func (CencDrmConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CencDrmConfiguration)(nil)).Elem()
}

func (o CencDrmConfigurationPtrOutput) ToCencDrmConfigurationPtrOutput() CencDrmConfigurationPtrOutput {
	return o
}

func (o CencDrmConfigurationPtrOutput) ToCencDrmConfigurationPtrOutputWithContext(ctx context.Context) CencDrmConfigurationPtrOutput {
	return o
}

func (o CencDrmConfigurationPtrOutput) Elem() CencDrmConfigurationOutput {
	return o.ApplyT(func(v *CencDrmConfiguration) CencDrmConfiguration {
		if v != nil {
			return *v
		}
		var ret CencDrmConfiguration
		return ret
	}).(CencDrmConfigurationOutput)
}

func (o CencDrmConfigurationPtrOutput) PlayReady() StreamingPolicyPlayReadyConfigurationPtrOutput {
	return o.ApplyT(func(v *CencDrmConfiguration) *StreamingPolicyPlayReadyConfiguration {
		if v == nil {
			return nil
		}
		return v.PlayReady
	}).(StreamingPolicyPlayReadyConfigurationPtrOutput)
}

func (o CencDrmConfigurationPtrOutput) Widevine() StreamingPolicyWidevineConfigurationPtrOutput {
	return o.ApplyT(func(v *CencDrmConfiguration) *StreamingPolicyWidevineConfiguration {
		if v == nil {
			return nil
		}
		return v.Widevine
	}).(StreamingPolicyWidevineConfigurationPtrOutput)
}

type CencDrmConfigurationResponse struct {
	PlayReady *StreamingPolicyPlayReadyConfigurationResponse `pulumi:"playReady"`
	Widevine  *StreamingPolicyWidevineConfigurationResponse  `pulumi:"widevine"`
}

type CencDrmConfigurationResponseOutput struct{ *pulumi.OutputState }

func (CencDrmConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CencDrmConfigurationResponse)(nil)).Elem()
}

func (o CencDrmConfigurationResponseOutput) ToCencDrmConfigurationResponseOutput() CencDrmConfigurationResponseOutput {
	return o
}

func (o CencDrmConfigurationResponseOutput) ToCencDrmConfigurationResponseOutputWithContext(ctx context.Context) CencDrmConfigurationResponseOutput {
	return o
}

func (o CencDrmConfigurationResponseOutput) PlayReady() StreamingPolicyPlayReadyConfigurationResponsePtrOutput {
	return o.ApplyT(func(v CencDrmConfigurationResponse) *StreamingPolicyPlayReadyConfigurationResponse {
		return v.PlayReady
	}).(StreamingPolicyPlayReadyConfigurationResponsePtrOutput)
}

func (o CencDrmConfigurationResponseOutput) Widevine() StreamingPolicyWidevineConfigurationResponsePtrOutput {
	return o.ApplyT(func(v CencDrmConfigurationResponse) *StreamingPolicyWidevineConfigurationResponse { return v.Widevine }).(StreamingPolicyWidevineConfigurationResponsePtrOutput)
}

type CencDrmConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (CencDrmConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CencDrmConfigurationResponse)(nil)).Elem()
}

func (o CencDrmConfigurationResponsePtrOutput) ToCencDrmConfigurationResponsePtrOutput() CencDrmConfigurationResponsePtrOutput {
	return o
}

func (o CencDrmConfigurationResponsePtrOutput) ToCencDrmConfigurationResponsePtrOutputWithContext(ctx context.Context) CencDrmConfigurationResponsePtrOutput {
	return o
}

func (o CencDrmConfigurationResponsePtrOutput) Elem() CencDrmConfigurationResponseOutput {
	return o.ApplyT(func(v *CencDrmConfigurationResponse) CencDrmConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret CencDrmConfigurationResponse
		return ret
	}).(CencDrmConfigurationResponseOutput)
}

func (o CencDrmConfigurationResponsePtrOutput) PlayReady() StreamingPolicyPlayReadyConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *CencDrmConfigurationResponse) *StreamingPolicyPlayReadyConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.PlayReady
	}).(StreamingPolicyPlayReadyConfigurationResponsePtrOutput)
}

func (o CencDrmConfigurationResponsePtrOutput) Widevine() StreamingPolicyWidevineConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *CencDrmConfigurationResponse) *StreamingPolicyWidevineConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Widevine
	}).(StreamingPolicyWidevineConfigurationResponsePtrOutput)
}

type CommonEncryptionCbcs struct {
	ClearTracks      []TrackSelection            `pulumi:"clearTracks"`
	ContentKeys      *StreamingPolicyContentKeys `pulumi:"contentKeys"`
	Drm              *CbcsDrmConfiguration       `pulumi:"drm"`
	EnabledProtocols *EnabledProtocols           `pulumi:"enabledProtocols"`
}





type CommonEncryptionCbcsInput interface {
	pulumi.Input

	ToCommonEncryptionCbcsOutput() CommonEncryptionCbcsOutput
	ToCommonEncryptionCbcsOutputWithContext(context.Context) CommonEncryptionCbcsOutput
}

type CommonEncryptionCbcsArgs struct {
	ClearTracks      TrackSelectionArrayInput           `pulumi:"clearTracks"`
	ContentKeys      StreamingPolicyContentKeysPtrInput `pulumi:"contentKeys"`
	Drm              CbcsDrmConfigurationPtrInput       `pulumi:"drm"`
	EnabledProtocols EnabledProtocolsPtrInput           `pulumi:"enabledProtocols"`
}

func (CommonEncryptionCbcsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CommonEncryptionCbcs)(nil)).Elem()
}

func (i CommonEncryptionCbcsArgs) ToCommonEncryptionCbcsOutput() CommonEncryptionCbcsOutput {
	return i.ToCommonEncryptionCbcsOutputWithContext(context.Background())
}

func (i CommonEncryptionCbcsArgs) ToCommonEncryptionCbcsOutputWithContext(ctx context.Context) CommonEncryptionCbcsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonEncryptionCbcsOutput)
}

func (i CommonEncryptionCbcsArgs) ToCommonEncryptionCbcsPtrOutput() CommonEncryptionCbcsPtrOutput {
	return i.ToCommonEncryptionCbcsPtrOutputWithContext(context.Background())
}

func (i CommonEncryptionCbcsArgs) ToCommonEncryptionCbcsPtrOutputWithContext(ctx context.Context) CommonEncryptionCbcsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonEncryptionCbcsOutput).ToCommonEncryptionCbcsPtrOutputWithContext(ctx)
}









type CommonEncryptionCbcsPtrInput interface {
	pulumi.Input

	ToCommonEncryptionCbcsPtrOutput() CommonEncryptionCbcsPtrOutput
	ToCommonEncryptionCbcsPtrOutputWithContext(context.Context) CommonEncryptionCbcsPtrOutput
}

type commonEncryptionCbcsPtrType CommonEncryptionCbcsArgs

func CommonEncryptionCbcsPtr(v *CommonEncryptionCbcsArgs) CommonEncryptionCbcsPtrInput {
	return (*commonEncryptionCbcsPtrType)(v)
}

func (*commonEncryptionCbcsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonEncryptionCbcs)(nil)).Elem()
}

func (i *commonEncryptionCbcsPtrType) ToCommonEncryptionCbcsPtrOutput() CommonEncryptionCbcsPtrOutput {
	return i.ToCommonEncryptionCbcsPtrOutputWithContext(context.Background())
}

func (i *commonEncryptionCbcsPtrType) ToCommonEncryptionCbcsPtrOutputWithContext(ctx context.Context) CommonEncryptionCbcsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonEncryptionCbcsPtrOutput)
}

type CommonEncryptionCbcsOutput struct{ *pulumi.OutputState }

func (CommonEncryptionCbcsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommonEncryptionCbcs)(nil)).Elem()
}

func (o CommonEncryptionCbcsOutput) ToCommonEncryptionCbcsOutput() CommonEncryptionCbcsOutput {
	return o
}

func (o CommonEncryptionCbcsOutput) ToCommonEncryptionCbcsOutputWithContext(ctx context.Context) CommonEncryptionCbcsOutput {
	return o
}

func (o CommonEncryptionCbcsOutput) ToCommonEncryptionCbcsPtrOutput() CommonEncryptionCbcsPtrOutput {
	return o.ToCommonEncryptionCbcsPtrOutputWithContext(context.Background())
}

func (o CommonEncryptionCbcsOutput) ToCommonEncryptionCbcsPtrOutputWithContext(ctx context.Context) CommonEncryptionCbcsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CommonEncryptionCbcs) *CommonEncryptionCbcs {
		return &v
	}).(CommonEncryptionCbcsPtrOutput)
}

func (o CommonEncryptionCbcsOutput) ClearTracks() TrackSelectionArrayOutput {
	return o.ApplyT(func(v CommonEncryptionCbcs) []TrackSelection { return v.ClearTracks }).(TrackSelectionArrayOutput)
}

func (o CommonEncryptionCbcsOutput) ContentKeys() StreamingPolicyContentKeysPtrOutput {
	return o.ApplyT(func(v CommonEncryptionCbcs) *StreamingPolicyContentKeys { return v.ContentKeys }).(StreamingPolicyContentKeysPtrOutput)
}

func (o CommonEncryptionCbcsOutput) Drm() CbcsDrmConfigurationPtrOutput {
	return o.ApplyT(func(v CommonEncryptionCbcs) *CbcsDrmConfiguration { return v.Drm }).(CbcsDrmConfigurationPtrOutput)
}

func (o CommonEncryptionCbcsOutput) EnabledProtocols() EnabledProtocolsPtrOutput {
	return o.ApplyT(func(v CommonEncryptionCbcs) *EnabledProtocols { return v.EnabledProtocols }).(EnabledProtocolsPtrOutput)
}

type CommonEncryptionCbcsPtrOutput struct{ *pulumi.OutputState }

func (CommonEncryptionCbcsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonEncryptionCbcs)(nil)).Elem()
}

func (o CommonEncryptionCbcsPtrOutput) ToCommonEncryptionCbcsPtrOutput() CommonEncryptionCbcsPtrOutput {
	return o
}

func (o CommonEncryptionCbcsPtrOutput) ToCommonEncryptionCbcsPtrOutputWithContext(ctx context.Context) CommonEncryptionCbcsPtrOutput {
	return o
}

func (o CommonEncryptionCbcsPtrOutput) Elem() CommonEncryptionCbcsOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcs) CommonEncryptionCbcs {
		if v != nil {
			return *v
		}
		var ret CommonEncryptionCbcs
		return ret
	}).(CommonEncryptionCbcsOutput)
}

func (o CommonEncryptionCbcsPtrOutput) ClearTracks() TrackSelectionArrayOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcs) []TrackSelection {
		if v == nil {
			return nil
		}
		return v.ClearTracks
	}).(TrackSelectionArrayOutput)
}

func (o CommonEncryptionCbcsPtrOutput) ContentKeys() StreamingPolicyContentKeysPtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcs) *StreamingPolicyContentKeys {
		if v == nil {
			return nil
		}
		return v.ContentKeys
	}).(StreamingPolicyContentKeysPtrOutput)
}

func (o CommonEncryptionCbcsPtrOutput) Drm() CbcsDrmConfigurationPtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcs) *CbcsDrmConfiguration {
		if v == nil {
			return nil
		}
		return v.Drm
	}).(CbcsDrmConfigurationPtrOutput)
}

func (o CommonEncryptionCbcsPtrOutput) EnabledProtocols() EnabledProtocolsPtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcs) *EnabledProtocols {
		if v == nil {
			return nil
		}
		return v.EnabledProtocols
	}).(EnabledProtocolsPtrOutput)
}

type CommonEncryptionCbcsResponse struct {
	ClearTracks      []TrackSelectionResponse            `pulumi:"clearTracks"`
	ContentKeys      *StreamingPolicyContentKeysResponse `pulumi:"contentKeys"`
	Drm              *CbcsDrmConfigurationResponse       `pulumi:"drm"`
	EnabledProtocols *EnabledProtocolsResponse           `pulumi:"enabledProtocols"`
}

type CommonEncryptionCbcsResponseOutput struct{ *pulumi.OutputState }

func (CommonEncryptionCbcsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommonEncryptionCbcsResponse)(nil)).Elem()
}

func (o CommonEncryptionCbcsResponseOutput) ToCommonEncryptionCbcsResponseOutput() CommonEncryptionCbcsResponseOutput {
	return o
}

func (o CommonEncryptionCbcsResponseOutput) ToCommonEncryptionCbcsResponseOutputWithContext(ctx context.Context) CommonEncryptionCbcsResponseOutput {
	return o
}

func (o CommonEncryptionCbcsResponseOutput) ClearTracks() TrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v CommonEncryptionCbcsResponse) []TrackSelectionResponse { return v.ClearTracks }).(TrackSelectionResponseArrayOutput)
}

func (o CommonEncryptionCbcsResponseOutput) ContentKeys() StreamingPolicyContentKeysResponsePtrOutput {
	return o.ApplyT(func(v CommonEncryptionCbcsResponse) *StreamingPolicyContentKeysResponse { return v.ContentKeys }).(StreamingPolicyContentKeysResponsePtrOutput)
}

func (o CommonEncryptionCbcsResponseOutput) Drm() CbcsDrmConfigurationResponsePtrOutput {
	return o.ApplyT(func(v CommonEncryptionCbcsResponse) *CbcsDrmConfigurationResponse { return v.Drm }).(CbcsDrmConfigurationResponsePtrOutput)
}

func (o CommonEncryptionCbcsResponseOutput) EnabledProtocols() EnabledProtocolsResponsePtrOutput {
	return o.ApplyT(func(v CommonEncryptionCbcsResponse) *EnabledProtocolsResponse { return v.EnabledProtocols }).(EnabledProtocolsResponsePtrOutput)
}

type CommonEncryptionCbcsResponsePtrOutput struct{ *pulumi.OutputState }

func (CommonEncryptionCbcsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonEncryptionCbcsResponse)(nil)).Elem()
}

func (o CommonEncryptionCbcsResponsePtrOutput) ToCommonEncryptionCbcsResponsePtrOutput() CommonEncryptionCbcsResponsePtrOutput {
	return o
}

func (o CommonEncryptionCbcsResponsePtrOutput) ToCommonEncryptionCbcsResponsePtrOutputWithContext(ctx context.Context) CommonEncryptionCbcsResponsePtrOutput {
	return o
}

func (o CommonEncryptionCbcsResponsePtrOutput) Elem() CommonEncryptionCbcsResponseOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcsResponse) CommonEncryptionCbcsResponse {
		if v != nil {
			return *v
		}
		var ret CommonEncryptionCbcsResponse
		return ret
	}).(CommonEncryptionCbcsResponseOutput)
}

func (o CommonEncryptionCbcsResponsePtrOutput) ClearTracks() TrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcsResponse) []TrackSelectionResponse {
		if v == nil {
			return nil
		}
		return v.ClearTracks
	}).(TrackSelectionResponseArrayOutput)
}

func (o CommonEncryptionCbcsResponsePtrOutput) ContentKeys() StreamingPolicyContentKeysResponsePtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcsResponse) *StreamingPolicyContentKeysResponse {
		if v == nil {
			return nil
		}
		return v.ContentKeys
	}).(StreamingPolicyContentKeysResponsePtrOutput)
}

func (o CommonEncryptionCbcsResponsePtrOutput) Drm() CbcsDrmConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcsResponse) *CbcsDrmConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Drm
	}).(CbcsDrmConfigurationResponsePtrOutput)
}

func (o CommonEncryptionCbcsResponsePtrOutput) EnabledProtocols() EnabledProtocolsResponsePtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCbcsResponse) *EnabledProtocolsResponse {
		if v == nil {
			return nil
		}
		return v.EnabledProtocols
	}).(EnabledProtocolsResponsePtrOutput)
}

type CommonEncryptionCenc struct {
	ClearTracks      []TrackSelection            `pulumi:"clearTracks"`
	ContentKeys      *StreamingPolicyContentKeys `pulumi:"contentKeys"`
	Drm              *CencDrmConfiguration       `pulumi:"drm"`
	EnabledProtocols *EnabledProtocols           `pulumi:"enabledProtocols"`
}





type CommonEncryptionCencInput interface {
	pulumi.Input

	ToCommonEncryptionCencOutput() CommonEncryptionCencOutput
	ToCommonEncryptionCencOutputWithContext(context.Context) CommonEncryptionCencOutput
}

type CommonEncryptionCencArgs struct {
	ClearTracks      TrackSelectionArrayInput           `pulumi:"clearTracks"`
	ContentKeys      StreamingPolicyContentKeysPtrInput `pulumi:"contentKeys"`
	Drm              CencDrmConfigurationPtrInput       `pulumi:"drm"`
	EnabledProtocols EnabledProtocolsPtrInput           `pulumi:"enabledProtocols"`
}

func (CommonEncryptionCencArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CommonEncryptionCenc)(nil)).Elem()
}

func (i CommonEncryptionCencArgs) ToCommonEncryptionCencOutput() CommonEncryptionCencOutput {
	return i.ToCommonEncryptionCencOutputWithContext(context.Background())
}

func (i CommonEncryptionCencArgs) ToCommonEncryptionCencOutputWithContext(ctx context.Context) CommonEncryptionCencOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonEncryptionCencOutput)
}

func (i CommonEncryptionCencArgs) ToCommonEncryptionCencPtrOutput() CommonEncryptionCencPtrOutput {
	return i.ToCommonEncryptionCencPtrOutputWithContext(context.Background())
}

func (i CommonEncryptionCencArgs) ToCommonEncryptionCencPtrOutputWithContext(ctx context.Context) CommonEncryptionCencPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonEncryptionCencOutput).ToCommonEncryptionCencPtrOutputWithContext(ctx)
}









type CommonEncryptionCencPtrInput interface {
	pulumi.Input

	ToCommonEncryptionCencPtrOutput() CommonEncryptionCencPtrOutput
	ToCommonEncryptionCencPtrOutputWithContext(context.Context) CommonEncryptionCencPtrOutput
}

type commonEncryptionCencPtrType CommonEncryptionCencArgs

func CommonEncryptionCencPtr(v *CommonEncryptionCencArgs) CommonEncryptionCencPtrInput {
	return (*commonEncryptionCencPtrType)(v)
}

func (*commonEncryptionCencPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonEncryptionCenc)(nil)).Elem()
}

func (i *commonEncryptionCencPtrType) ToCommonEncryptionCencPtrOutput() CommonEncryptionCencPtrOutput {
	return i.ToCommonEncryptionCencPtrOutputWithContext(context.Background())
}

func (i *commonEncryptionCencPtrType) ToCommonEncryptionCencPtrOutputWithContext(ctx context.Context) CommonEncryptionCencPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommonEncryptionCencPtrOutput)
}

type CommonEncryptionCencOutput struct{ *pulumi.OutputState }

func (CommonEncryptionCencOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommonEncryptionCenc)(nil)).Elem()
}

func (o CommonEncryptionCencOutput) ToCommonEncryptionCencOutput() CommonEncryptionCencOutput {
	return o
}

func (o CommonEncryptionCencOutput) ToCommonEncryptionCencOutputWithContext(ctx context.Context) CommonEncryptionCencOutput {
	return o
}

func (o CommonEncryptionCencOutput) ToCommonEncryptionCencPtrOutput() CommonEncryptionCencPtrOutput {
	return o.ToCommonEncryptionCencPtrOutputWithContext(context.Background())
}

func (o CommonEncryptionCencOutput) ToCommonEncryptionCencPtrOutputWithContext(ctx context.Context) CommonEncryptionCencPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CommonEncryptionCenc) *CommonEncryptionCenc {
		return &v
	}).(CommonEncryptionCencPtrOutput)
}

func (o CommonEncryptionCencOutput) ClearTracks() TrackSelectionArrayOutput {
	return o.ApplyT(func(v CommonEncryptionCenc) []TrackSelection { return v.ClearTracks }).(TrackSelectionArrayOutput)
}

func (o CommonEncryptionCencOutput) ContentKeys() StreamingPolicyContentKeysPtrOutput {
	return o.ApplyT(func(v CommonEncryptionCenc) *StreamingPolicyContentKeys { return v.ContentKeys }).(StreamingPolicyContentKeysPtrOutput)
}

func (o CommonEncryptionCencOutput) Drm() CencDrmConfigurationPtrOutput {
	return o.ApplyT(func(v CommonEncryptionCenc) *CencDrmConfiguration { return v.Drm }).(CencDrmConfigurationPtrOutput)
}

func (o CommonEncryptionCencOutput) EnabledProtocols() EnabledProtocolsPtrOutput {
	return o.ApplyT(func(v CommonEncryptionCenc) *EnabledProtocols { return v.EnabledProtocols }).(EnabledProtocolsPtrOutput)
}

type CommonEncryptionCencPtrOutput struct{ *pulumi.OutputState }

func (CommonEncryptionCencPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonEncryptionCenc)(nil)).Elem()
}

func (o CommonEncryptionCencPtrOutput) ToCommonEncryptionCencPtrOutput() CommonEncryptionCencPtrOutput {
	return o
}

func (o CommonEncryptionCencPtrOutput) ToCommonEncryptionCencPtrOutputWithContext(ctx context.Context) CommonEncryptionCencPtrOutput {
	return o
}

func (o CommonEncryptionCencPtrOutput) Elem() CommonEncryptionCencOutput {
	return o.ApplyT(func(v *CommonEncryptionCenc) CommonEncryptionCenc {
		if v != nil {
			return *v
		}
		var ret CommonEncryptionCenc
		return ret
	}).(CommonEncryptionCencOutput)
}

func (o CommonEncryptionCencPtrOutput) ClearTracks() TrackSelectionArrayOutput {
	return o.ApplyT(func(v *CommonEncryptionCenc) []TrackSelection {
		if v == nil {
			return nil
		}
		return v.ClearTracks
	}).(TrackSelectionArrayOutput)
}

func (o CommonEncryptionCencPtrOutput) ContentKeys() StreamingPolicyContentKeysPtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCenc) *StreamingPolicyContentKeys {
		if v == nil {
			return nil
		}
		return v.ContentKeys
	}).(StreamingPolicyContentKeysPtrOutput)
}

func (o CommonEncryptionCencPtrOutput) Drm() CencDrmConfigurationPtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCenc) *CencDrmConfiguration {
		if v == nil {
			return nil
		}
		return v.Drm
	}).(CencDrmConfigurationPtrOutput)
}

func (o CommonEncryptionCencPtrOutput) EnabledProtocols() EnabledProtocolsPtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCenc) *EnabledProtocols {
		if v == nil {
			return nil
		}
		return v.EnabledProtocols
	}).(EnabledProtocolsPtrOutput)
}

type CommonEncryptionCencResponse struct {
	ClearTracks      []TrackSelectionResponse            `pulumi:"clearTracks"`
	ContentKeys      *StreamingPolicyContentKeysResponse `pulumi:"contentKeys"`
	Drm              *CencDrmConfigurationResponse       `pulumi:"drm"`
	EnabledProtocols *EnabledProtocolsResponse           `pulumi:"enabledProtocols"`
}

type CommonEncryptionCencResponseOutput struct{ *pulumi.OutputState }

func (CommonEncryptionCencResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommonEncryptionCencResponse)(nil)).Elem()
}

func (o CommonEncryptionCencResponseOutput) ToCommonEncryptionCencResponseOutput() CommonEncryptionCencResponseOutput {
	return o
}

func (o CommonEncryptionCencResponseOutput) ToCommonEncryptionCencResponseOutputWithContext(ctx context.Context) CommonEncryptionCencResponseOutput {
	return o
}

func (o CommonEncryptionCencResponseOutput) ClearTracks() TrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v CommonEncryptionCencResponse) []TrackSelectionResponse { return v.ClearTracks }).(TrackSelectionResponseArrayOutput)
}

func (o CommonEncryptionCencResponseOutput) ContentKeys() StreamingPolicyContentKeysResponsePtrOutput {
	return o.ApplyT(func(v CommonEncryptionCencResponse) *StreamingPolicyContentKeysResponse { return v.ContentKeys }).(StreamingPolicyContentKeysResponsePtrOutput)
}

func (o CommonEncryptionCencResponseOutput) Drm() CencDrmConfigurationResponsePtrOutput {
	return o.ApplyT(func(v CommonEncryptionCencResponse) *CencDrmConfigurationResponse { return v.Drm }).(CencDrmConfigurationResponsePtrOutput)
}

func (o CommonEncryptionCencResponseOutput) EnabledProtocols() EnabledProtocolsResponsePtrOutput {
	return o.ApplyT(func(v CommonEncryptionCencResponse) *EnabledProtocolsResponse { return v.EnabledProtocols }).(EnabledProtocolsResponsePtrOutput)
}

type CommonEncryptionCencResponsePtrOutput struct{ *pulumi.OutputState }

func (CommonEncryptionCencResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonEncryptionCencResponse)(nil)).Elem()
}

func (o CommonEncryptionCencResponsePtrOutput) ToCommonEncryptionCencResponsePtrOutput() CommonEncryptionCencResponsePtrOutput {
	return o
}

func (o CommonEncryptionCencResponsePtrOutput) ToCommonEncryptionCencResponsePtrOutputWithContext(ctx context.Context) CommonEncryptionCencResponsePtrOutput {
	return o
}

func (o CommonEncryptionCencResponsePtrOutput) Elem() CommonEncryptionCencResponseOutput {
	return o.ApplyT(func(v *CommonEncryptionCencResponse) CommonEncryptionCencResponse {
		if v != nil {
			return *v
		}
		var ret CommonEncryptionCencResponse
		return ret
	}).(CommonEncryptionCencResponseOutput)
}

func (o CommonEncryptionCencResponsePtrOutput) ClearTracks() TrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v *CommonEncryptionCencResponse) []TrackSelectionResponse {
		if v == nil {
			return nil
		}
		return v.ClearTracks
	}).(TrackSelectionResponseArrayOutput)
}

func (o CommonEncryptionCencResponsePtrOutput) ContentKeys() StreamingPolicyContentKeysResponsePtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCencResponse) *StreamingPolicyContentKeysResponse {
		if v == nil {
			return nil
		}
		return v.ContentKeys
	}).(StreamingPolicyContentKeysResponsePtrOutput)
}

func (o CommonEncryptionCencResponsePtrOutput) Drm() CencDrmConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCencResponse) *CencDrmConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Drm
	}).(CencDrmConfigurationResponsePtrOutput)
}

func (o CommonEncryptionCencResponsePtrOutput) EnabledProtocols() EnabledProtocolsResponsePtrOutput {
	return o.ApplyT(func(v *CommonEncryptionCencResponse) *EnabledProtocolsResponse {
		if v == nil {
			return nil
		}
		return v.EnabledProtocols
	}).(EnabledProtocolsResponsePtrOutput)
}

type ContentKeyPolicyClearKeyConfiguration struct {
	OdataType string `pulumi:"odataType"`
}

type ContentKeyPolicyClearKeyConfigurationResponse struct {
	OdataType string `pulumi:"odataType"`
}

type ContentKeyPolicyFairPlayConfiguration struct {
	Ask                   string  `pulumi:"ask"`
	FairPlayPfx           string  `pulumi:"fairPlayPfx"`
	FairPlayPfxPassword   string  `pulumi:"fairPlayPfxPassword"`
	OdataType             string  `pulumi:"odataType"`
	RentalAndLeaseKeyType string  `pulumi:"rentalAndLeaseKeyType"`
	RentalDuration        float64 `pulumi:"rentalDuration"`
}

type ContentKeyPolicyFairPlayConfigurationResponse struct {
	Ask                   string  `pulumi:"ask"`
	FairPlayPfx           string  `pulumi:"fairPlayPfx"`
	FairPlayPfxPassword   string  `pulumi:"fairPlayPfxPassword"`
	OdataType             string  `pulumi:"odataType"`
	RentalAndLeaseKeyType string  `pulumi:"rentalAndLeaseKeyType"`
	RentalDuration        float64 `pulumi:"rentalDuration"`
}

type ContentKeyPolicyOpenRestriction struct {
	OdataType string `pulumi:"odataType"`
}

type ContentKeyPolicyOpenRestrictionResponse struct {
	OdataType string `pulumi:"odataType"`
}

type ContentKeyPolicyOption struct {
	Configuration interface{} `pulumi:"configuration"`
	Name          *string     `pulumi:"name"`
	Restriction   interface{} `pulumi:"restriction"`
}





type ContentKeyPolicyOptionInput interface {
	pulumi.Input

	ToContentKeyPolicyOptionOutput() ContentKeyPolicyOptionOutput
	ToContentKeyPolicyOptionOutputWithContext(context.Context) ContentKeyPolicyOptionOutput
}

type ContentKeyPolicyOptionArgs struct {
	Configuration pulumi.Input          `pulumi:"configuration"`
	Name          pulumi.StringPtrInput `pulumi:"name"`
	Restriction   pulumi.Input          `pulumi:"restriction"`
}

func (ContentKeyPolicyOptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyOption)(nil)).Elem()
}

func (i ContentKeyPolicyOptionArgs) ToContentKeyPolicyOptionOutput() ContentKeyPolicyOptionOutput {
	return i.ToContentKeyPolicyOptionOutputWithContext(context.Background())
}

func (i ContentKeyPolicyOptionArgs) ToContentKeyPolicyOptionOutputWithContext(ctx context.Context) ContentKeyPolicyOptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyOptionOutput)
}





type ContentKeyPolicyOptionArrayInput interface {
	pulumi.Input

	ToContentKeyPolicyOptionArrayOutput() ContentKeyPolicyOptionArrayOutput
	ToContentKeyPolicyOptionArrayOutputWithContext(context.Context) ContentKeyPolicyOptionArrayOutput
}

type ContentKeyPolicyOptionArray []ContentKeyPolicyOptionInput

func (ContentKeyPolicyOptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentKeyPolicyOption)(nil)).Elem()
}

func (i ContentKeyPolicyOptionArray) ToContentKeyPolicyOptionArrayOutput() ContentKeyPolicyOptionArrayOutput {
	return i.ToContentKeyPolicyOptionArrayOutputWithContext(context.Background())
}

func (i ContentKeyPolicyOptionArray) ToContentKeyPolicyOptionArrayOutputWithContext(ctx context.Context) ContentKeyPolicyOptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyOptionArrayOutput)
}

type ContentKeyPolicyOptionOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyOptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyOption)(nil)).Elem()
}

func (o ContentKeyPolicyOptionOutput) ToContentKeyPolicyOptionOutput() ContentKeyPolicyOptionOutput {
	return o
}

func (o ContentKeyPolicyOptionOutput) ToContentKeyPolicyOptionOutputWithContext(ctx context.Context) ContentKeyPolicyOptionOutput {
	return o
}

func (o ContentKeyPolicyOptionOutput) Configuration() pulumi.AnyOutput {
	return o.ApplyT(func(v ContentKeyPolicyOption) interface{} { return v.Configuration }).(pulumi.AnyOutput)
}

func (o ContentKeyPolicyOptionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyOption) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyOptionOutput) Restriction() pulumi.AnyOutput {
	return o.ApplyT(func(v ContentKeyPolicyOption) interface{} { return v.Restriction }).(pulumi.AnyOutput)
}

type ContentKeyPolicyOptionArrayOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyOptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentKeyPolicyOption)(nil)).Elem()
}

func (o ContentKeyPolicyOptionArrayOutput) ToContentKeyPolicyOptionArrayOutput() ContentKeyPolicyOptionArrayOutput {
	return o
}

func (o ContentKeyPolicyOptionArrayOutput) ToContentKeyPolicyOptionArrayOutputWithContext(ctx context.Context) ContentKeyPolicyOptionArrayOutput {
	return o
}

func (o ContentKeyPolicyOptionArrayOutput) Index(i pulumi.IntInput) ContentKeyPolicyOptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContentKeyPolicyOption {
		return vs[0].([]ContentKeyPolicyOption)[vs[1].(int)]
	}).(ContentKeyPolicyOptionOutput)
}

type ContentKeyPolicyOptionResponse struct {
	Configuration  interface{} `pulumi:"configuration"`
	Name           *string     `pulumi:"name"`
	PolicyOptionId string      `pulumi:"policyOptionId"`
	Restriction    interface{} `pulumi:"restriction"`
}

type ContentKeyPolicyOptionResponseOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyOptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentKeyPolicyOptionResponse)(nil)).Elem()
}

func (o ContentKeyPolicyOptionResponseOutput) ToContentKeyPolicyOptionResponseOutput() ContentKeyPolicyOptionResponseOutput {
	return o
}

func (o ContentKeyPolicyOptionResponseOutput) ToContentKeyPolicyOptionResponseOutputWithContext(ctx context.Context) ContentKeyPolicyOptionResponseOutput {
	return o
}

func (o ContentKeyPolicyOptionResponseOutput) Configuration() pulumi.AnyOutput {
	return o.ApplyT(func(v ContentKeyPolicyOptionResponse) interface{} { return v.Configuration }).(pulumi.AnyOutput)
}

func (o ContentKeyPolicyOptionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentKeyPolicyOptionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ContentKeyPolicyOptionResponseOutput) PolicyOptionId() pulumi.StringOutput {
	return o.ApplyT(func(v ContentKeyPolicyOptionResponse) string { return v.PolicyOptionId }).(pulumi.StringOutput)
}

func (o ContentKeyPolicyOptionResponseOutput) Restriction() pulumi.AnyOutput {
	return o.ApplyT(func(v ContentKeyPolicyOptionResponse) interface{} { return v.Restriction }).(pulumi.AnyOutput)
}

type ContentKeyPolicyOptionResponseArrayOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyOptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentKeyPolicyOptionResponse)(nil)).Elem()
}

func (o ContentKeyPolicyOptionResponseArrayOutput) ToContentKeyPolicyOptionResponseArrayOutput() ContentKeyPolicyOptionResponseArrayOutput {
	return o
}

func (o ContentKeyPolicyOptionResponseArrayOutput) ToContentKeyPolicyOptionResponseArrayOutputWithContext(ctx context.Context) ContentKeyPolicyOptionResponseArrayOutput {
	return o
}

func (o ContentKeyPolicyOptionResponseArrayOutput) Index(i pulumi.IntInput) ContentKeyPolicyOptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContentKeyPolicyOptionResponse {
		return vs[0].([]ContentKeyPolicyOptionResponse)[vs[1].(int)]
	}).(ContentKeyPolicyOptionResponseOutput)
}

type ContentKeyPolicyPlayReadyConfiguration struct {
	Licenses           []ContentKeyPolicyPlayReadyLicense `pulumi:"licenses"`
	OdataType          string                             `pulumi:"odataType"`
	ResponseCustomData *string                            `pulumi:"responseCustomData"`
}

type ContentKeyPolicyPlayReadyConfigurationResponse struct {
	Licenses           []ContentKeyPolicyPlayReadyLicenseResponse `pulumi:"licenses"`
	OdataType          string                                     `pulumi:"odataType"`
	ResponseCustomData *string                                    `pulumi:"responseCustomData"`
}

type ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader struct {
	OdataType string `pulumi:"odataType"`
}

type ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeaderResponse struct {
	OdataType string `pulumi:"odataType"`
}

type ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifier struct {
	KeyId     string `pulumi:"keyId"`
	OdataType string `pulumi:"odataType"`
}

type ContentKeyPolicyPlayReadyContentEncryptionKeyFromKeyIdentifierResponse struct {
	KeyId     string `pulumi:"keyId"`
	OdataType string `pulumi:"odataType"`
}

type ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction struct {
	BestEffort        bool `pulumi:"bestEffort"`
	ConfigurationData int  `pulumi:"configurationData"`
}

type ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse struct {
	BestEffort        bool `pulumi:"bestEffort"`
	ConfigurationData int  `pulumi:"configurationData"`
}

type ContentKeyPolicyPlayReadyLicense struct {
	AllowTestDevices       bool                                `pulumi:"allowTestDevices"`
	BeginDate              *string                             `pulumi:"beginDate"`
	ContentKeyLocation     interface{}                         `pulumi:"contentKeyLocation"`
	ContentType            string                              `pulumi:"contentType"`
	ExpirationDate         *string                             `pulumi:"expirationDate"`
	GracePeriod            *string                             `pulumi:"gracePeriod"`
	LicenseType            string                              `pulumi:"licenseType"`
	PlayRight              *ContentKeyPolicyPlayReadyPlayRight `pulumi:"playRight"`
	RelativeBeginDate      *string                             `pulumi:"relativeBeginDate"`
	RelativeExpirationDate *string                             `pulumi:"relativeExpirationDate"`
}

type ContentKeyPolicyPlayReadyLicenseResponse struct {
	AllowTestDevices       bool                                        `pulumi:"allowTestDevices"`
	BeginDate              *string                                     `pulumi:"beginDate"`
	ContentKeyLocation     interface{}                                 `pulumi:"contentKeyLocation"`
	ContentType            string                                      `pulumi:"contentType"`
	ExpirationDate         *string                                     `pulumi:"expirationDate"`
	GracePeriod            *string                                     `pulumi:"gracePeriod"`
	LicenseType            string                                      `pulumi:"licenseType"`
	PlayRight              *ContentKeyPolicyPlayReadyPlayRightResponse `pulumi:"playRight"`
	RelativeBeginDate      *string                                     `pulumi:"relativeBeginDate"`
	RelativeExpirationDate *string                                     `pulumi:"relativeExpirationDate"`
}

type ContentKeyPolicyPlayReadyPlayRight struct {
	AgcAndColorStripeRestriction                       *int                                                          `pulumi:"agcAndColorStripeRestriction"`
	AllowPassingVideoContentToUnknownOutput            string                                                        `pulumi:"allowPassingVideoContentToUnknownOutput"`
	AnalogVideoOpl                                     *int                                                          `pulumi:"analogVideoOpl"`
	CompressedDigitalAudioOpl                          *int                                                          `pulumi:"compressedDigitalAudioOpl"`
	CompressedDigitalVideoOpl                          *int                                                          `pulumi:"compressedDigitalVideoOpl"`
	DigitalVideoOnlyContentRestriction                 bool                                                          `pulumi:"digitalVideoOnlyContentRestriction"`
	ExplicitAnalogTelevisionOutputRestriction          *ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestriction `pulumi:"explicitAnalogTelevisionOutputRestriction"`
	FirstPlayExpiration                                *string                                                       `pulumi:"firstPlayExpiration"`
	ImageConstraintForAnalogComponentVideoRestriction  bool                                                          `pulumi:"imageConstraintForAnalogComponentVideoRestriction"`
	ImageConstraintForAnalogComputerMonitorRestriction bool                                                          `pulumi:"imageConstraintForAnalogComputerMonitorRestriction"`
	ScmsRestriction                                    *int                                                          `pulumi:"scmsRestriction"`
	UncompressedDigitalAudioOpl                        *int                                                          `pulumi:"uncompressedDigitalAudioOpl"`
	UncompressedDigitalVideoOpl                        *int                                                          `pulumi:"uncompressedDigitalVideoOpl"`
}

type ContentKeyPolicyPlayReadyPlayRightResponse struct {
	AgcAndColorStripeRestriction                       *int                                                                  `pulumi:"agcAndColorStripeRestriction"`
	AllowPassingVideoContentToUnknownOutput            string                                                                `pulumi:"allowPassingVideoContentToUnknownOutput"`
	AnalogVideoOpl                                     *int                                                                  `pulumi:"analogVideoOpl"`
	CompressedDigitalAudioOpl                          *int                                                                  `pulumi:"compressedDigitalAudioOpl"`
	CompressedDigitalVideoOpl                          *int                                                                  `pulumi:"compressedDigitalVideoOpl"`
	DigitalVideoOnlyContentRestriction                 bool                                                                  `pulumi:"digitalVideoOnlyContentRestriction"`
	ExplicitAnalogTelevisionOutputRestriction          *ContentKeyPolicyPlayReadyExplicitAnalogTelevisionRestrictionResponse `pulumi:"explicitAnalogTelevisionOutputRestriction"`
	FirstPlayExpiration                                *string                                                               `pulumi:"firstPlayExpiration"`
	ImageConstraintForAnalogComponentVideoRestriction  bool                                                                  `pulumi:"imageConstraintForAnalogComponentVideoRestriction"`
	ImageConstraintForAnalogComputerMonitorRestriction bool                                                                  `pulumi:"imageConstraintForAnalogComputerMonitorRestriction"`
	ScmsRestriction                                    *int                                                                  `pulumi:"scmsRestriction"`
	UncompressedDigitalAudioOpl                        *int                                                                  `pulumi:"uncompressedDigitalAudioOpl"`
	UncompressedDigitalVideoOpl                        *int                                                                  `pulumi:"uncompressedDigitalVideoOpl"`
}

type ContentKeyPolicyRsaTokenKey struct {
	Exponent  string `pulumi:"exponent"`
	Modulus   string `pulumi:"modulus"`
	OdataType string `pulumi:"odataType"`
}

type ContentKeyPolicyRsaTokenKeyResponse struct {
	Exponent  string `pulumi:"exponent"`
	Modulus   string `pulumi:"modulus"`
	OdataType string `pulumi:"odataType"`
}

type ContentKeyPolicySymmetricTokenKey struct {
	KeyValue  string `pulumi:"keyValue"`
	OdataType string `pulumi:"odataType"`
}

type ContentKeyPolicySymmetricTokenKeyResponse struct {
	KeyValue  string `pulumi:"keyValue"`
	OdataType string `pulumi:"odataType"`
}

type ContentKeyPolicyTokenClaim struct {
	ClaimType  *string `pulumi:"claimType"`
	ClaimValue *string `pulumi:"claimValue"`
}

type ContentKeyPolicyTokenClaimResponse struct {
	ClaimType  *string `pulumi:"claimType"`
	ClaimValue *string `pulumi:"claimValue"`
}

type ContentKeyPolicyTokenRestriction struct {
	AlternateVerificationKeys      []interface{}                `pulumi:"alternateVerificationKeys"`
	Audience                       string                       `pulumi:"audience"`
	Issuer                         string                       `pulumi:"issuer"`
	OdataType                      string                       `pulumi:"odataType"`
	OpenIdConnectDiscoveryDocument *string                      `pulumi:"openIdConnectDiscoveryDocument"`
	PrimaryVerificationKey         interface{}                  `pulumi:"primaryVerificationKey"`
	RequiredClaims                 []ContentKeyPolicyTokenClaim `pulumi:"requiredClaims"`
	RestrictionTokenType           string                       `pulumi:"restrictionTokenType"`
}

type ContentKeyPolicyTokenRestrictionResponse struct {
	AlternateVerificationKeys      []interface{}                        `pulumi:"alternateVerificationKeys"`
	Audience                       string                               `pulumi:"audience"`
	Issuer                         string                               `pulumi:"issuer"`
	OdataType                      string                               `pulumi:"odataType"`
	OpenIdConnectDiscoveryDocument *string                              `pulumi:"openIdConnectDiscoveryDocument"`
	PrimaryVerificationKey         interface{}                          `pulumi:"primaryVerificationKey"`
	RequiredClaims                 []ContentKeyPolicyTokenClaimResponse `pulumi:"requiredClaims"`
	RestrictionTokenType           string                               `pulumi:"restrictionTokenType"`
}

type ContentKeyPolicyUnknownConfiguration struct {
	OdataType string `pulumi:"odataType"`
}

type ContentKeyPolicyUnknownConfigurationResponse struct {
	OdataType string `pulumi:"odataType"`
}

type ContentKeyPolicyUnknownRestriction struct {
	OdataType string `pulumi:"odataType"`
}

type ContentKeyPolicyUnknownRestrictionResponse struct {
	OdataType string `pulumi:"odataType"`
}

type ContentKeyPolicyWidevineConfiguration struct {
	OdataType        string `pulumi:"odataType"`
	WidevineTemplate string `pulumi:"widevineTemplate"`
}

type ContentKeyPolicyWidevineConfigurationResponse struct {
	OdataType        string `pulumi:"odataType"`
	WidevineTemplate string `pulumi:"widevineTemplate"`
}

type ContentKeyPolicyX509CertificateTokenKey struct {
	OdataType string `pulumi:"odataType"`
	RawBody   string `pulumi:"rawBody"`
}

type ContentKeyPolicyX509CertificateTokenKeyResponse struct {
	OdataType string `pulumi:"odataType"`
	RawBody   string `pulumi:"rawBody"`
}

type CopyAudio struct {
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
}

type CopyAudioResponse struct {
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
}

type CopyVideo struct {
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
}

type CopyVideoResponse struct {
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
}

type CrossSiteAccessPolicies struct {
	ClientAccessPolicy *string `pulumi:"clientAccessPolicy"`
	CrossDomainPolicy  *string `pulumi:"crossDomainPolicy"`
}





type CrossSiteAccessPoliciesInput interface {
	pulumi.Input

	ToCrossSiteAccessPoliciesOutput() CrossSiteAccessPoliciesOutput
	ToCrossSiteAccessPoliciesOutputWithContext(context.Context) CrossSiteAccessPoliciesOutput
}

type CrossSiteAccessPoliciesArgs struct {
	ClientAccessPolicy pulumi.StringPtrInput `pulumi:"clientAccessPolicy"`
	CrossDomainPolicy  pulumi.StringPtrInput `pulumi:"crossDomainPolicy"`
}

func (CrossSiteAccessPoliciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CrossSiteAccessPolicies)(nil)).Elem()
}

func (i CrossSiteAccessPoliciesArgs) ToCrossSiteAccessPoliciesOutput() CrossSiteAccessPoliciesOutput {
	return i.ToCrossSiteAccessPoliciesOutputWithContext(context.Background())
}

func (i CrossSiteAccessPoliciesArgs) ToCrossSiteAccessPoliciesOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossSiteAccessPoliciesOutput)
}

func (i CrossSiteAccessPoliciesArgs) ToCrossSiteAccessPoliciesPtrOutput() CrossSiteAccessPoliciesPtrOutput {
	return i.ToCrossSiteAccessPoliciesPtrOutputWithContext(context.Background())
}

func (i CrossSiteAccessPoliciesArgs) ToCrossSiteAccessPoliciesPtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossSiteAccessPoliciesOutput).ToCrossSiteAccessPoliciesPtrOutputWithContext(ctx)
}









type CrossSiteAccessPoliciesPtrInput interface {
	pulumi.Input

	ToCrossSiteAccessPoliciesPtrOutput() CrossSiteAccessPoliciesPtrOutput
	ToCrossSiteAccessPoliciesPtrOutputWithContext(context.Context) CrossSiteAccessPoliciesPtrOutput
}

type crossSiteAccessPoliciesPtrType CrossSiteAccessPoliciesArgs

func CrossSiteAccessPoliciesPtr(v *CrossSiteAccessPoliciesArgs) CrossSiteAccessPoliciesPtrInput {
	return (*crossSiteAccessPoliciesPtrType)(v)
}

func (*crossSiteAccessPoliciesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CrossSiteAccessPolicies)(nil)).Elem()
}

func (i *crossSiteAccessPoliciesPtrType) ToCrossSiteAccessPoliciesPtrOutput() CrossSiteAccessPoliciesPtrOutput {
	return i.ToCrossSiteAccessPoliciesPtrOutputWithContext(context.Background())
}

func (i *crossSiteAccessPoliciesPtrType) ToCrossSiteAccessPoliciesPtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CrossSiteAccessPoliciesPtrOutput)
}

type CrossSiteAccessPoliciesOutput struct{ *pulumi.OutputState }

func (CrossSiteAccessPoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CrossSiteAccessPolicies)(nil)).Elem()
}

func (o CrossSiteAccessPoliciesOutput) ToCrossSiteAccessPoliciesOutput() CrossSiteAccessPoliciesOutput {
	return o
}

func (o CrossSiteAccessPoliciesOutput) ToCrossSiteAccessPoliciesOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesOutput {
	return o
}

func (o CrossSiteAccessPoliciesOutput) ToCrossSiteAccessPoliciesPtrOutput() CrossSiteAccessPoliciesPtrOutput {
	return o.ToCrossSiteAccessPoliciesPtrOutputWithContext(context.Background())
}

func (o CrossSiteAccessPoliciesOutput) ToCrossSiteAccessPoliciesPtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CrossSiteAccessPolicies) *CrossSiteAccessPolicies {
		return &v
	}).(CrossSiteAccessPoliciesPtrOutput)
}

func (o CrossSiteAccessPoliciesOutput) ClientAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CrossSiteAccessPolicies) *string { return v.ClientAccessPolicy }).(pulumi.StringPtrOutput)
}

func (o CrossSiteAccessPoliciesOutput) CrossDomainPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CrossSiteAccessPolicies) *string { return v.CrossDomainPolicy }).(pulumi.StringPtrOutput)
}

type CrossSiteAccessPoliciesPtrOutput struct{ *pulumi.OutputState }

func (CrossSiteAccessPoliciesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CrossSiteAccessPolicies)(nil)).Elem()
}

func (o CrossSiteAccessPoliciesPtrOutput) ToCrossSiteAccessPoliciesPtrOutput() CrossSiteAccessPoliciesPtrOutput {
	return o
}

func (o CrossSiteAccessPoliciesPtrOutput) ToCrossSiteAccessPoliciesPtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesPtrOutput {
	return o
}

func (o CrossSiteAccessPoliciesPtrOutput) Elem() CrossSiteAccessPoliciesOutput {
	return o.ApplyT(func(v *CrossSiteAccessPolicies) CrossSiteAccessPolicies {
		if v != nil {
			return *v
		}
		var ret CrossSiteAccessPolicies
		return ret
	}).(CrossSiteAccessPoliciesOutput)
}

func (o CrossSiteAccessPoliciesPtrOutput) ClientAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CrossSiteAccessPolicies) *string {
		if v == nil {
			return nil
		}
		return v.ClientAccessPolicy
	}).(pulumi.StringPtrOutput)
}

func (o CrossSiteAccessPoliciesPtrOutput) CrossDomainPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CrossSiteAccessPolicies) *string {
		if v == nil {
			return nil
		}
		return v.CrossDomainPolicy
	}).(pulumi.StringPtrOutput)
}

type CrossSiteAccessPoliciesResponse struct {
	ClientAccessPolicy *string `pulumi:"clientAccessPolicy"`
	CrossDomainPolicy  *string `pulumi:"crossDomainPolicy"`
}

type CrossSiteAccessPoliciesResponseOutput struct{ *pulumi.OutputState }

func (CrossSiteAccessPoliciesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CrossSiteAccessPoliciesResponse)(nil)).Elem()
}

func (o CrossSiteAccessPoliciesResponseOutput) ToCrossSiteAccessPoliciesResponseOutput() CrossSiteAccessPoliciesResponseOutput {
	return o
}

func (o CrossSiteAccessPoliciesResponseOutput) ToCrossSiteAccessPoliciesResponseOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesResponseOutput {
	return o
}

func (o CrossSiteAccessPoliciesResponseOutput) ClientAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CrossSiteAccessPoliciesResponse) *string { return v.ClientAccessPolicy }).(pulumi.StringPtrOutput)
}

func (o CrossSiteAccessPoliciesResponseOutput) CrossDomainPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CrossSiteAccessPoliciesResponse) *string { return v.CrossDomainPolicy }).(pulumi.StringPtrOutput)
}

type CrossSiteAccessPoliciesResponsePtrOutput struct{ *pulumi.OutputState }

func (CrossSiteAccessPoliciesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CrossSiteAccessPoliciesResponse)(nil)).Elem()
}

func (o CrossSiteAccessPoliciesResponsePtrOutput) ToCrossSiteAccessPoliciesResponsePtrOutput() CrossSiteAccessPoliciesResponsePtrOutput {
	return o
}

func (o CrossSiteAccessPoliciesResponsePtrOutput) ToCrossSiteAccessPoliciesResponsePtrOutputWithContext(ctx context.Context) CrossSiteAccessPoliciesResponsePtrOutput {
	return o
}

func (o CrossSiteAccessPoliciesResponsePtrOutput) Elem() CrossSiteAccessPoliciesResponseOutput {
	return o.ApplyT(func(v *CrossSiteAccessPoliciesResponse) CrossSiteAccessPoliciesResponse {
		if v != nil {
			return *v
		}
		var ret CrossSiteAccessPoliciesResponse
		return ret
	}).(CrossSiteAccessPoliciesResponseOutput)
}

func (o CrossSiteAccessPoliciesResponsePtrOutput) ClientAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CrossSiteAccessPoliciesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientAccessPolicy
	}).(pulumi.StringPtrOutput)
}

func (o CrossSiteAccessPoliciesResponsePtrOutput) CrossDomainPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CrossSiteAccessPoliciesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CrossDomainPolicy
	}).(pulumi.StringPtrOutput)
}

type DefaultKey struct {
	Label      *string `pulumi:"label"`
	PolicyName *string `pulumi:"policyName"`
}





type DefaultKeyInput interface {
	pulumi.Input

	ToDefaultKeyOutput() DefaultKeyOutput
	ToDefaultKeyOutputWithContext(context.Context) DefaultKeyOutput
}

type DefaultKeyArgs struct {
	Label      pulumi.StringPtrInput `pulumi:"label"`
	PolicyName pulumi.StringPtrInput `pulumi:"policyName"`
}

func (DefaultKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultKey)(nil)).Elem()
}

func (i DefaultKeyArgs) ToDefaultKeyOutput() DefaultKeyOutput {
	return i.ToDefaultKeyOutputWithContext(context.Background())
}

func (i DefaultKeyArgs) ToDefaultKeyOutputWithContext(ctx context.Context) DefaultKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultKeyOutput)
}

func (i DefaultKeyArgs) ToDefaultKeyPtrOutput() DefaultKeyPtrOutput {
	return i.ToDefaultKeyPtrOutputWithContext(context.Background())
}

func (i DefaultKeyArgs) ToDefaultKeyPtrOutputWithContext(ctx context.Context) DefaultKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultKeyOutput).ToDefaultKeyPtrOutputWithContext(ctx)
}









type DefaultKeyPtrInput interface {
	pulumi.Input

	ToDefaultKeyPtrOutput() DefaultKeyPtrOutput
	ToDefaultKeyPtrOutputWithContext(context.Context) DefaultKeyPtrOutput
}

type defaultKeyPtrType DefaultKeyArgs

func DefaultKeyPtr(v *DefaultKeyArgs) DefaultKeyPtrInput {
	return (*defaultKeyPtrType)(v)
}

func (*defaultKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultKey)(nil)).Elem()
}

func (i *defaultKeyPtrType) ToDefaultKeyPtrOutput() DefaultKeyPtrOutput {
	return i.ToDefaultKeyPtrOutputWithContext(context.Background())
}

func (i *defaultKeyPtrType) ToDefaultKeyPtrOutputWithContext(ctx context.Context) DefaultKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultKeyPtrOutput)
}

type DefaultKeyOutput struct{ *pulumi.OutputState }

func (DefaultKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultKey)(nil)).Elem()
}

func (o DefaultKeyOutput) ToDefaultKeyOutput() DefaultKeyOutput {
	return o
}

func (o DefaultKeyOutput) ToDefaultKeyOutputWithContext(ctx context.Context) DefaultKeyOutput {
	return o
}

func (o DefaultKeyOutput) ToDefaultKeyPtrOutput() DefaultKeyPtrOutput {
	return o.ToDefaultKeyPtrOutputWithContext(context.Background())
}

func (o DefaultKeyOutput) ToDefaultKeyPtrOutputWithContext(ctx context.Context) DefaultKeyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefaultKey) *DefaultKey {
		return &v
	}).(DefaultKeyPtrOutput)
}

func (o DefaultKeyOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefaultKey) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o DefaultKeyOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefaultKey) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

type DefaultKeyPtrOutput struct{ *pulumi.OutputState }

func (DefaultKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultKey)(nil)).Elem()
}

func (o DefaultKeyPtrOutput) ToDefaultKeyPtrOutput() DefaultKeyPtrOutput {
	return o
}

func (o DefaultKeyPtrOutput) ToDefaultKeyPtrOutputWithContext(ctx context.Context) DefaultKeyPtrOutput {
	return o
}

func (o DefaultKeyPtrOutput) Elem() DefaultKeyOutput {
	return o.ApplyT(func(v *DefaultKey) DefaultKey {
		if v != nil {
			return *v
		}
		var ret DefaultKey
		return ret
	}).(DefaultKeyOutput)
}

func (o DefaultKeyPtrOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultKey) *string {
		if v == nil {
			return nil
		}
		return v.Label
	}).(pulumi.StringPtrOutput)
}

func (o DefaultKeyPtrOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultKey) *string {
		if v == nil {
			return nil
		}
		return v.PolicyName
	}).(pulumi.StringPtrOutput)
}

type DefaultKeyResponse struct {
	Label      *string `pulumi:"label"`
	PolicyName *string `pulumi:"policyName"`
}

type DefaultKeyResponseOutput struct{ *pulumi.OutputState }

func (DefaultKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultKeyResponse)(nil)).Elem()
}

func (o DefaultKeyResponseOutput) ToDefaultKeyResponseOutput() DefaultKeyResponseOutput {
	return o
}

func (o DefaultKeyResponseOutput) ToDefaultKeyResponseOutputWithContext(ctx context.Context) DefaultKeyResponseOutput {
	return o
}

func (o DefaultKeyResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefaultKeyResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o DefaultKeyResponseOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefaultKeyResponse) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

type DefaultKeyResponsePtrOutput struct{ *pulumi.OutputState }

func (DefaultKeyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultKeyResponse)(nil)).Elem()
}

func (o DefaultKeyResponsePtrOutput) ToDefaultKeyResponsePtrOutput() DefaultKeyResponsePtrOutput {
	return o
}

func (o DefaultKeyResponsePtrOutput) ToDefaultKeyResponsePtrOutputWithContext(ctx context.Context) DefaultKeyResponsePtrOutput {
	return o
}

func (o DefaultKeyResponsePtrOutput) Elem() DefaultKeyResponseOutput {
	return o.ApplyT(func(v *DefaultKeyResponse) DefaultKeyResponse {
		if v != nil {
			return *v
		}
		var ret DefaultKeyResponse
		return ret
	}).(DefaultKeyResponseOutput)
}

func (o DefaultKeyResponsePtrOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultKeyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Label
	}).(pulumi.StringPtrOutput)
}

func (o DefaultKeyResponsePtrOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultKeyResponse) *string {
		if v == nil {
			return nil
		}
		return v.PolicyName
	}).(pulumi.StringPtrOutput)
}

type Deinterlace struct {
	Mode   *string `pulumi:"mode"`
	Parity *string `pulumi:"parity"`
}

type DeinterlaceResponse struct {
	Mode   *string `pulumi:"mode"`
	Parity *string `pulumi:"parity"`
}

type EnabledProtocols struct {
	Dash            bool `pulumi:"dash"`
	Download        bool `pulumi:"download"`
	Hls             bool `pulumi:"hls"`
	SmoothStreaming bool `pulumi:"smoothStreaming"`
}





type EnabledProtocolsInput interface {
	pulumi.Input

	ToEnabledProtocolsOutput() EnabledProtocolsOutput
	ToEnabledProtocolsOutputWithContext(context.Context) EnabledProtocolsOutput
}

type EnabledProtocolsArgs struct {
	Dash            pulumi.BoolInput `pulumi:"dash"`
	Download        pulumi.BoolInput `pulumi:"download"`
	Hls             pulumi.BoolInput `pulumi:"hls"`
	SmoothStreaming pulumi.BoolInput `pulumi:"smoothStreaming"`
}

func (EnabledProtocolsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledProtocols)(nil)).Elem()
}

func (i EnabledProtocolsArgs) ToEnabledProtocolsOutput() EnabledProtocolsOutput {
	return i.ToEnabledProtocolsOutputWithContext(context.Background())
}

func (i EnabledProtocolsArgs) ToEnabledProtocolsOutputWithContext(ctx context.Context) EnabledProtocolsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledProtocolsOutput)
}

func (i EnabledProtocolsArgs) ToEnabledProtocolsPtrOutput() EnabledProtocolsPtrOutput {
	return i.ToEnabledProtocolsPtrOutputWithContext(context.Background())
}

func (i EnabledProtocolsArgs) ToEnabledProtocolsPtrOutputWithContext(ctx context.Context) EnabledProtocolsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledProtocolsOutput).ToEnabledProtocolsPtrOutputWithContext(ctx)
}









type EnabledProtocolsPtrInput interface {
	pulumi.Input

	ToEnabledProtocolsPtrOutput() EnabledProtocolsPtrOutput
	ToEnabledProtocolsPtrOutputWithContext(context.Context) EnabledProtocolsPtrOutput
}

type enabledProtocolsPtrType EnabledProtocolsArgs

func EnabledProtocolsPtr(v *EnabledProtocolsArgs) EnabledProtocolsPtrInput {
	return (*enabledProtocolsPtrType)(v)
}

func (*enabledProtocolsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledProtocols)(nil)).Elem()
}

func (i *enabledProtocolsPtrType) ToEnabledProtocolsPtrOutput() EnabledProtocolsPtrOutput {
	return i.ToEnabledProtocolsPtrOutputWithContext(context.Background())
}

func (i *enabledProtocolsPtrType) ToEnabledProtocolsPtrOutputWithContext(ctx context.Context) EnabledProtocolsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledProtocolsPtrOutput)
}

type EnabledProtocolsOutput struct{ *pulumi.OutputState }

func (EnabledProtocolsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledProtocols)(nil)).Elem()
}

func (o EnabledProtocolsOutput) ToEnabledProtocolsOutput() EnabledProtocolsOutput {
	return o
}

func (o EnabledProtocolsOutput) ToEnabledProtocolsOutputWithContext(ctx context.Context) EnabledProtocolsOutput {
	return o
}

func (o EnabledProtocolsOutput) ToEnabledProtocolsPtrOutput() EnabledProtocolsPtrOutput {
	return o.ToEnabledProtocolsPtrOutputWithContext(context.Background())
}

func (o EnabledProtocolsOutput) ToEnabledProtocolsPtrOutputWithContext(ctx context.Context) EnabledProtocolsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnabledProtocols) *EnabledProtocols {
		return &v
	}).(EnabledProtocolsPtrOutput)
}

func (o EnabledProtocolsOutput) Dash() pulumi.BoolOutput {
	return o.ApplyT(func(v EnabledProtocols) bool { return v.Dash }).(pulumi.BoolOutput)
}

func (o EnabledProtocolsOutput) Download() pulumi.BoolOutput {
	return o.ApplyT(func(v EnabledProtocols) bool { return v.Download }).(pulumi.BoolOutput)
}

func (o EnabledProtocolsOutput) Hls() pulumi.BoolOutput {
	return o.ApplyT(func(v EnabledProtocols) bool { return v.Hls }).(pulumi.BoolOutput)
}

func (o EnabledProtocolsOutput) SmoothStreaming() pulumi.BoolOutput {
	return o.ApplyT(func(v EnabledProtocols) bool { return v.SmoothStreaming }).(pulumi.BoolOutput)
}

type EnabledProtocolsPtrOutput struct{ *pulumi.OutputState }

func (EnabledProtocolsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledProtocols)(nil)).Elem()
}

func (o EnabledProtocolsPtrOutput) ToEnabledProtocolsPtrOutput() EnabledProtocolsPtrOutput {
	return o
}

func (o EnabledProtocolsPtrOutput) ToEnabledProtocolsPtrOutputWithContext(ctx context.Context) EnabledProtocolsPtrOutput {
	return o
}

func (o EnabledProtocolsPtrOutput) Elem() EnabledProtocolsOutput {
	return o.ApplyT(func(v *EnabledProtocols) EnabledProtocols {
		if v != nil {
			return *v
		}
		var ret EnabledProtocols
		return ret
	}).(EnabledProtocolsOutput)
}

func (o EnabledProtocolsPtrOutput) Dash() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledProtocols) *bool {
		if v == nil {
			return nil
		}
		return &v.Dash
	}).(pulumi.BoolPtrOutput)
}

func (o EnabledProtocolsPtrOutput) Download() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledProtocols) *bool {
		if v == nil {
			return nil
		}
		return &v.Download
	}).(pulumi.BoolPtrOutput)
}

func (o EnabledProtocolsPtrOutput) Hls() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledProtocols) *bool {
		if v == nil {
			return nil
		}
		return &v.Hls
	}).(pulumi.BoolPtrOutput)
}

func (o EnabledProtocolsPtrOutput) SmoothStreaming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledProtocols) *bool {
		if v == nil {
			return nil
		}
		return &v.SmoothStreaming
	}).(pulumi.BoolPtrOutput)
}

type EnabledProtocolsResponse struct {
	Dash            bool `pulumi:"dash"`
	Download        bool `pulumi:"download"`
	Hls             bool `pulumi:"hls"`
	SmoothStreaming bool `pulumi:"smoothStreaming"`
}

type EnabledProtocolsResponseOutput struct{ *pulumi.OutputState }

func (EnabledProtocolsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledProtocolsResponse)(nil)).Elem()
}

func (o EnabledProtocolsResponseOutput) ToEnabledProtocolsResponseOutput() EnabledProtocolsResponseOutput {
	return o
}

func (o EnabledProtocolsResponseOutput) ToEnabledProtocolsResponseOutputWithContext(ctx context.Context) EnabledProtocolsResponseOutput {
	return o
}

func (o EnabledProtocolsResponseOutput) Dash() pulumi.BoolOutput {
	return o.ApplyT(func(v EnabledProtocolsResponse) bool { return v.Dash }).(pulumi.BoolOutput)
}

func (o EnabledProtocolsResponseOutput) Download() pulumi.BoolOutput {
	return o.ApplyT(func(v EnabledProtocolsResponse) bool { return v.Download }).(pulumi.BoolOutput)
}

func (o EnabledProtocolsResponseOutput) Hls() pulumi.BoolOutput {
	return o.ApplyT(func(v EnabledProtocolsResponse) bool { return v.Hls }).(pulumi.BoolOutput)
}

func (o EnabledProtocolsResponseOutput) SmoothStreaming() pulumi.BoolOutput {
	return o.ApplyT(func(v EnabledProtocolsResponse) bool { return v.SmoothStreaming }).(pulumi.BoolOutput)
}

type EnabledProtocolsResponsePtrOutput struct{ *pulumi.OutputState }

func (EnabledProtocolsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledProtocolsResponse)(nil)).Elem()
}

func (o EnabledProtocolsResponsePtrOutput) ToEnabledProtocolsResponsePtrOutput() EnabledProtocolsResponsePtrOutput {
	return o
}

func (o EnabledProtocolsResponsePtrOutput) ToEnabledProtocolsResponsePtrOutputWithContext(ctx context.Context) EnabledProtocolsResponsePtrOutput {
	return o
}

func (o EnabledProtocolsResponsePtrOutput) Elem() EnabledProtocolsResponseOutput {
	return o.ApplyT(func(v *EnabledProtocolsResponse) EnabledProtocolsResponse {
		if v != nil {
			return *v
		}
		var ret EnabledProtocolsResponse
		return ret
	}).(EnabledProtocolsResponseOutput)
}

func (o EnabledProtocolsResponsePtrOutput) Dash() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledProtocolsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Dash
	}).(pulumi.BoolPtrOutput)
}

func (o EnabledProtocolsResponsePtrOutput) Download() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledProtocolsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Download
	}).(pulumi.BoolPtrOutput)
}

func (o EnabledProtocolsResponsePtrOutput) Hls() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledProtocolsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Hls
	}).(pulumi.BoolPtrOutput)
}

func (o EnabledProtocolsResponsePtrOutput) SmoothStreaming() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledProtocolsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.SmoothStreaming
	}).(pulumi.BoolPtrOutput)
}

type EnvelopeEncryption struct {
	ClearTracks                     []TrackSelection            `pulumi:"clearTracks"`
	ContentKeys                     *StreamingPolicyContentKeys `pulumi:"contentKeys"`
	CustomKeyAcquisitionUrlTemplate *string                     `pulumi:"customKeyAcquisitionUrlTemplate"`
	EnabledProtocols                *EnabledProtocols           `pulumi:"enabledProtocols"`
}





type EnvelopeEncryptionInput interface {
	pulumi.Input

	ToEnvelopeEncryptionOutput() EnvelopeEncryptionOutput
	ToEnvelopeEncryptionOutputWithContext(context.Context) EnvelopeEncryptionOutput
}

type EnvelopeEncryptionArgs struct {
	ClearTracks                     TrackSelectionArrayInput           `pulumi:"clearTracks"`
	ContentKeys                     StreamingPolicyContentKeysPtrInput `pulumi:"contentKeys"`
	CustomKeyAcquisitionUrlTemplate pulumi.StringPtrInput              `pulumi:"customKeyAcquisitionUrlTemplate"`
	EnabledProtocols                EnabledProtocolsPtrInput           `pulumi:"enabledProtocols"`
}

func (EnvelopeEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvelopeEncryption)(nil)).Elem()
}

func (i EnvelopeEncryptionArgs) ToEnvelopeEncryptionOutput() EnvelopeEncryptionOutput {
	return i.ToEnvelopeEncryptionOutputWithContext(context.Background())
}

func (i EnvelopeEncryptionArgs) ToEnvelopeEncryptionOutputWithContext(ctx context.Context) EnvelopeEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvelopeEncryptionOutput)
}

func (i EnvelopeEncryptionArgs) ToEnvelopeEncryptionPtrOutput() EnvelopeEncryptionPtrOutput {
	return i.ToEnvelopeEncryptionPtrOutputWithContext(context.Background())
}

func (i EnvelopeEncryptionArgs) ToEnvelopeEncryptionPtrOutputWithContext(ctx context.Context) EnvelopeEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvelopeEncryptionOutput).ToEnvelopeEncryptionPtrOutputWithContext(ctx)
}









type EnvelopeEncryptionPtrInput interface {
	pulumi.Input

	ToEnvelopeEncryptionPtrOutput() EnvelopeEncryptionPtrOutput
	ToEnvelopeEncryptionPtrOutputWithContext(context.Context) EnvelopeEncryptionPtrOutput
}

type envelopeEncryptionPtrType EnvelopeEncryptionArgs

func EnvelopeEncryptionPtr(v *EnvelopeEncryptionArgs) EnvelopeEncryptionPtrInput {
	return (*envelopeEncryptionPtrType)(v)
}

func (*envelopeEncryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvelopeEncryption)(nil)).Elem()
}

func (i *envelopeEncryptionPtrType) ToEnvelopeEncryptionPtrOutput() EnvelopeEncryptionPtrOutput {
	return i.ToEnvelopeEncryptionPtrOutputWithContext(context.Background())
}

func (i *envelopeEncryptionPtrType) ToEnvelopeEncryptionPtrOutputWithContext(ctx context.Context) EnvelopeEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvelopeEncryptionPtrOutput)
}

type EnvelopeEncryptionOutput struct{ *pulumi.OutputState }

func (EnvelopeEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvelopeEncryption)(nil)).Elem()
}

func (o EnvelopeEncryptionOutput) ToEnvelopeEncryptionOutput() EnvelopeEncryptionOutput {
	return o
}

func (o EnvelopeEncryptionOutput) ToEnvelopeEncryptionOutputWithContext(ctx context.Context) EnvelopeEncryptionOutput {
	return o
}

func (o EnvelopeEncryptionOutput) ToEnvelopeEncryptionPtrOutput() EnvelopeEncryptionPtrOutput {
	return o.ToEnvelopeEncryptionPtrOutputWithContext(context.Background())
}

func (o EnvelopeEncryptionOutput) ToEnvelopeEncryptionPtrOutputWithContext(ctx context.Context) EnvelopeEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvelopeEncryption) *EnvelopeEncryption {
		return &v
	}).(EnvelopeEncryptionPtrOutput)
}

func (o EnvelopeEncryptionOutput) ClearTracks() TrackSelectionArrayOutput {
	return o.ApplyT(func(v EnvelopeEncryption) []TrackSelection { return v.ClearTracks }).(TrackSelectionArrayOutput)
}

func (o EnvelopeEncryptionOutput) ContentKeys() StreamingPolicyContentKeysPtrOutput {
	return o.ApplyT(func(v EnvelopeEncryption) *StreamingPolicyContentKeys { return v.ContentKeys }).(StreamingPolicyContentKeysPtrOutput)
}

func (o EnvelopeEncryptionOutput) CustomKeyAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvelopeEncryption) *string { return v.CustomKeyAcquisitionUrlTemplate }).(pulumi.StringPtrOutput)
}

func (o EnvelopeEncryptionOutput) EnabledProtocols() EnabledProtocolsPtrOutput {
	return o.ApplyT(func(v EnvelopeEncryption) *EnabledProtocols { return v.EnabledProtocols }).(EnabledProtocolsPtrOutput)
}

type EnvelopeEncryptionPtrOutput struct{ *pulumi.OutputState }

func (EnvelopeEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvelopeEncryption)(nil)).Elem()
}

func (o EnvelopeEncryptionPtrOutput) ToEnvelopeEncryptionPtrOutput() EnvelopeEncryptionPtrOutput {
	return o
}

func (o EnvelopeEncryptionPtrOutput) ToEnvelopeEncryptionPtrOutputWithContext(ctx context.Context) EnvelopeEncryptionPtrOutput {
	return o
}

func (o EnvelopeEncryptionPtrOutput) Elem() EnvelopeEncryptionOutput {
	return o.ApplyT(func(v *EnvelopeEncryption) EnvelopeEncryption {
		if v != nil {
			return *v
		}
		var ret EnvelopeEncryption
		return ret
	}).(EnvelopeEncryptionOutput)
}

func (o EnvelopeEncryptionPtrOutput) ClearTracks() TrackSelectionArrayOutput {
	return o.ApplyT(func(v *EnvelopeEncryption) []TrackSelection {
		if v == nil {
			return nil
		}
		return v.ClearTracks
	}).(TrackSelectionArrayOutput)
}

func (o EnvelopeEncryptionPtrOutput) ContentKeys() StreamingPolicyContentKeysPtrOutput {
	return o.ApplyT(func(v *EnvelopeEncryption) *StreamingPolicyContentKeys {
		if v == nil {
			return nil
		}
		return v.ContentKeys
	}).(StreamingPolicyContentKeysPtrOutput)
}

func (o EnvelopeEncryptionPtrOutput) CustomKeyAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvelopeEncryption) *string {
		if v == nil {
			return nil
		}
		return v.CustomKeyAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

func (o EnvelopeEncryptionPtrOutput) EnabledProtocols() EnabledProtocolsPtrOutput {
	return o.ApplyT(func(v *EnvelopeEncryption) *EnabledProtocols {
		if v == nil {
			return nil
		}
		return v.EnabledProtocols
	}).(EnabledProtocolsPtrOutput)
}

type EnvelopeEncryptionResponse struct {
	ClearTracks                     []TrackSelectionResponse            `pulumi:"clearTracks"`
	ContentKeys                     *StreamingPolicyContentKeysResponse `pulumi:"contentKeys"`
	CustomKeyAcquisitionUrlTemplate *string                             `pulumi:"customKeyAcquisitionUrlTemplate"`
	EnabledProtocols                *EnabledProtocolsResponse           `pulumi:"enabledProtocols"`
}

type EnvelopeEncryptionResponseOutput struct{ *pulumi.OutputState }

func (EnvelopeEncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvelopeEncryptionResponse)(nil)).Elem()
}

func (o EnvelopeEncryptionResponseOutput) ToEnvelopeEncryptionResponseOutput() EnvelopeEncryptionResponseOutput {
	return o
}

func (o EnvelopeEncryptionResponseOutput) ToEnvelopeEncryptionResponseOutputWithContext(ctx context.Context) EnvelopeEncryptionResponseOutput {
	return o
}

func (o EnvelopeEncryptionResponseOutput) ClearTracks() TrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v EnvelopeEncryptionResponse) []TrackSelectionResponse { return v.ClearTracks }).(TrackSelectionResponseArrayOutput)
}

func (o EnvelopeEncryptionResponseOutput) ContentKeys() StreamingPolicyContentKeysResponsePtrOutput {
	return o.ApplyT(func(v EnvelopeEncryptionResponse) *StreamingPolicyContentKeysResponse { return v.ContentKeys }).(StreamingPolicyContentKeysResponsePtrOutput)
}

func (o EnvelopeEncryptionResponseOutput) CustomKeyAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvelopeEncryptionResponse) *string { return v.CustomKeyAcquisitionUrlTemplate }).(pulumi.StringPtrOutput)
}

func (o EnvelopeEncryptionResponseOutput) EnabledProtocols() EnabledProtocolsResponsePtrOutput {
	return o.ApplyT(func(v EnvelopeEncryptionResponse) *EnabledProtocolsResponse { return v.EnabledProtocols }).(EnabledProtocolsResponsePtrOutput)
}

type EnvelopeEncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EnvelopeEncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvelopeEncryptionResponse)(nil)).Elem()
}

func (o EnvelopeEncryptionResponsePtrOutput) ToEnvelopeEncryptionResponsePtrOutput() EnvelopeEncryptionResponsePtrOutput {
	return o
}

func (o EnvelopeEncryptionResponsePtrOutput) ToEnvelopeEncryptionResponsePtrOutputWithContext(ctx context.Context) EnvelopeEncryptionResponsePtrOutput {
	return o
}

func (o EnvelopeEncryptionResponsePtrOutput) Elem() EnvelopeEncryptionResponseOutput {
	return o.ApplyT(func(v *EnvelopeEncryptionResponse) EnvelopeEncryptionResponse {
		if v != nil {
			return *v
		}
		var ret EnvelopeEncryptionResponse
		return ret
	}).(EnvelopeEncryptionResponseOutput)
}

func (o EnvelopeEncryptionResponsePtrOutput) ClearTracks() TrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v *EnvelopeEncryptionResponse) []TrackSelectionResponse {
		if v == nil {
			return nil
		}
		return v.ClearTracks
	}).(TrackSelectionResponseArrayOutput)
}

func (o EnvelopeEncryptionResponsePtrOutput) ContentKeys() StreamingPolicyContentKeysResponsePtrOutput {
	return o.ApplyT(func(v *EnvelopeEncryptionResponse) *StreamingPolicyContentKeysResponse {
		if v == nil {
			return nil
		}
		return v.ContentKeys
	}).(StreamingPolicyContentKeysResponsePtrOutput)
}

func (o EnvelopeEncryptionResponsePtrOutput) CustomKeyAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvelopeEncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomKeyAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

func (o EnvelopeEncryptionResponsePtrOutput) EnabledProtocols() EnabledProtocolsResponsePtrOutput {
	return o.ApplyT(func(v *EnvelopeEncryptionResponse) *EnabledProtocolsResponse {
		if v == nil {
			return nil
		}
		return v.EnabledProtocols
	}).(EnabledProtocolsResponsePtrOutput)
}

type Filters struct {
	Crop        *Rectangle    `pulumi:"crop"`
	Deinterlace *Deinterlace  `pulumi:"deinterlace"`
	Overlays    []interface{} `pulumi:"overlays"`
	Rotation    *string       `pulumi:"rotation"`
}

type FiltersResponse struct {
	Crop        *RectangleResponse   `pulumi:"crop"`
	Deinterlace *DeinterlaceResponse `pulumi:"deinterlace"`
	Overlays    []interface{}        `pulumi:"overlays"`
	Rotation    *string              `pulumi:"rotation"`
}

type H264Layer struct {
	AdaptiveBFrame  *bool   `pulumi:"adaptiveBFrame"`
	BFrames         *int    `pulumi:"bFrames"`
	Bitrate         *int    `pulumi:"bitrate"`
	BufferWindow    *string `pulumi:"bufferWindow"`
	EntropyMode     *string `pulumi:"entropyMode"`
	FrameRate       *string `pulumi:"frameRate"`
	Height          *string `pulumi:"height"`
	Label           *string `pulumi:"label"`
	Level           *string `pulumi:"level"`
	MaxBitrate      *int    `pulumi:"maxBitrate"`
	OdataType       string  `pulumi:"odataType"`
	Profile         *string `pulumi:"profile"`
	ReferenceFrames *int    `pulumi:"referenceFrames"`
	Slices          *int    `pulumi:"slices"`
	Width           *string `pulumi:"width"`
}

type H264LayerResponse struct {
	AdaptiveBFrame  *bool   `pulumi:"adaptiveBFrame"`
	BFrames         *int    `pulumi:"bFrames"`
	Bitrate         *int    `pulumi:"bitrate"`
	BufferWindow    *string `pulumi:"bufferWindow"`
	EntropyMode     *string `pulumi:"entropyMode"`
	FrameRate       *string `pulumi:"frameRate"`
	Height          *string `pulumi:"height"`
	Label           *string `pulumi:"label"`
	Level           *string `pulumi:"level"`
	MaxBitrate      *int    `pulumi:"maxBitrate"`
	OdataType       string  `pulumi:"odataType"`
	Profile         *string `pulumi:"profile"`
	ReferenceFrames *int    `pulumi:"referenceFrames"`
	Slices          *int    `pulumi:"slices"`
	Width           *string `pulumi:"width"`
}

type H264Video struct {
	Complexity           *string     `pulumi:"complexity"`
	KeyFrameInterval     *string     `pulumi:"keyFrameInterval"`
	Label                *string     `pulumi:"label"`
	Layers               []H264Layer `pulumi:"layers"`
	OdataType            string      `pulumi:"odataType"`
	SceneChangeDetection *bool       `pulumi:"sceneChangeDetection"`
	StretchMode          *string     `pulumi:"stretchMode"`
}

type H264VideoResponse struct {
	Complexity           *string             `pulumi:"complexity"`
	KeyFrameInterval     *string             `pulumi:"keyFrameInterval"`
	Label                *string             `pulumi:"label"`
	Layers               []H264LayerResponse `pulumi:"layers"`
	OdataType            string              `pulumi:"odataType"`
	SceneChangeDetection *bool               `pulumi:"sceneChangeDetection"`
	StretchMode          *string             `pulumi:"stretchMode"`
}

type Hls struct {
	FragmentsPerTsSegment *int `pulumi:"fragmentsPerTsSegment"`
}





type HlsInput interface {
	pulumi.Input

	ToHlsOutput() HlsOutput
	ToHlsOutputWithContext(context.Context) HlsOutput
}

type HlsArgs struct {
	FragmentsPerTsSegment pulumi.IntPtrInput `pulumi:"fragmentsPerTsSegment"`
}

func (HlsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Hls)(nil)).Elem()
}

func (i HlsArgs) ToHlsOutput() HlsOutput {
	return i.ToHlsOutputWithContext(context.Background())
}

func (i HlsArgs) ToHlsOutputWithContext(ctx context.Context) HlsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HlsOutput)
}

func (i HlsArgs) ToHlsPtrOutput() HlsPtrOutput {
	return i.ToHlsPtrOutputWithContext(context.Background())
}

func (i HlsArgs) ToHlsPtrOutputWithContext(ctx context.Context) HlsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HlsOutput).ToHlsPtrOutputWithContext(ctx)
}









type HlsPtrInput interface {
	pulumi.Input

	ToHlsPtrOutput() HlsPtrOutput
	ToHlsPtrOutputWithContext(context.Context) HlsPtrOutput
}

type hlsPtrType HlsArgs

func HlsPtr(v *HlsArgs) HlsPtrInput {
	return (*hlsPtrType)(v)
}

func (*hlsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Hls)(nil)).Elem()
}

func (i *hlsPtrType) ToHlsPtrOutput() HlsPtrOutput {
	return i.ToHlsPtrOutputWithContext(context.Background())
}

func (i *hlsPtrType) ToHlsPtrOutputWithContext(ctx context.Context) HlsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HlsPtrOutput)
}

type HlsOutput struct{ *pulumi.OutputState }

func (HlsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Hls)(nil)).Elem()
}

func (o HlsOutput) ToHlsOutput() HlsOutput {
	return o
}

func (o HlsOutput) ToHlsOutputWithContext(ctx context.Context) HlsOutput {
	return o
}

func (o HlsOutput) ToHlsPtrOutput() HlsPtrOutput {
	return o.ToHlsPtrOutputWithContext(context.Background())
}

func (o HlsOutput) ToHlsPtrOutputWithContext(ctx context.Context) HlsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Hls) *Hls {
		return &v
	}).(HlsPtrOutput)
}

func (o HlsOutput) FragmentsPerTsSegment() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Hls) *int { return v.FragmentsPerTsSegment }).(pulumi.IntPtrOutput)
}

type HlsPtrOutput struct{ *pulumi.OutputState }

func (HlsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Hls)(nil)).Elem()
}

func (o HlsPtrOutput) ToHlsPtrOutput() HlsPtrOutput {
	return o
}

func (o HlsPtrOutput) ToHlsPtrOutputWithContext(ctx context.Context) HlsPtrOutput {
	return o
}

func (o HlsPtrOutput) Elem() HlsOutput {
	return o.ApplyT(func(v *Hls) Hls {
		if v != nil {
			return *v
		}
		var ret Hls
		return ret
	}).(HlsOutput)
}

func (o HlsPtrOutput) FragmentsPerTsSegment() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Hls) *int {
		if v == nil {
			return nil
		}
		return v.FragmentsPerTsSegment
	}).(pulumi.IntPtrOutput)
}

type HlsResponse struct {
	FragmentsPerTsSegment *int `pulumi:"fragmentsPerTsSegment"`
}

type HlsResponseOutput struct{ *pulumi.OutputState }

func (HlsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HlsResponse)(nil)).Elem()
}

func (o HlsResponseOutput) ToHlsResponseOutput() HlsResponseOutput {
	return o
}

func (o HlsResponseOutput) ToHlsResponseOutputWithContext(ctx context.Context) HlsResponseOutput {
	return o
}

func (o HlsResponseOutput) FragmentsPerTsSegment() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HlsResponse) *int { return v.FragmentsPerTsSegment }).(pulumi.IntPtrOutput)
}

type HlsResponsePtrOutput struct{ *pulumi.OutputState }

func (HlsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HlsResponse)(nil)).Elem()
}

func (o HlsResponsePtrOutput) ToHlsResponsePtrOutput() HlsResponsePtrOutput {
	return o
}

func (o HlsResponsePtrOutput) ToHlsResponsePtrOutputWithContext(ctx context.Context) HlsResponsePtrOutput {
	return o
}

func (o HlsResponsePtrOutput) Elem() HlsResponseOutput {
	return o.ApplyT(func(v *HlsResponse) HlsResponse {
		if v != nil {
			return *v
		}
		var ret HlsResponse
		return ret
	}).(HlsResponseOutput)
}

func (o HlsResponsePtrOutput) FragmentsPerTsSegment() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HlsResponse) *int {
		if v == nil {
			return nil
		}
		return v.FragmentsPerTsSegment
	}).(pulumi.IntPtrOutput)
}

type IPAccessControl struct {
	Allow []IPRange `pulumi:"allow"`
}





type IPAccessControlInput interface {
	pulumi.Input

	ToIPAccessControlOutput() IPAccessControlOutput
	ToIPAccessControlOutputWithContext(context.Context) IPAccessControlOutput
}

type IPAccessControlArgs struct {
	Allow IPRangeArrayInput `pulumi:"allow"`
}

func (IPAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAccessControl)(nil)).Elem()
}

func (i IPAccessControlArgs) ToIPAccessControlOutput() IPAccessControlOutput {
	return i.ToIPAccessControlOutputWithContext(context.Background())
}

func (i IPAccessControlArgs) ToIPAccessControlOutputWithContext(ctx context.Context) IPAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPAccessControlOutput)
}

func (i IPAccessControlArgs) ToIPAccessControlPtrOutput() IPAccessControlPtrOutput {
	return i.ToIPAccessControlPtrOutputWithContext(context.Background())
}

func (i IPAccessControlArgs) ToIPAccessControlPtrOutputWithContext(ctx context.Context) IPAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPAccessControlOutput).ToIPAccessControlPtrOutputWithContext(ctx)
}









type IPAccessControlPtrInput interface {
	pulumi.Input

	ToIPAccessControlPtrOutput() IPAccessControlPtrOutput
	ToIPAccessControlPtrOutputWithContext(context.Context) IPAccessControlPtrOutput
}

type ipaccessControlPtrType IPAccessControlArgs

func IPAccessControlPtr(v *IPAccessControlArgs) IPAccessControlPtrInput {
	return (*ipaccessControlPtrType)(v)
}

func (*ipaccessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IPAccessControl)(nil)).Elem()
}

func (i *ipaccessControlPtrType) ToIPAccessControlPtrOutput() IPAccessControlPtrOutput {
	return i.ToIPAccessControlPtrOutputWithContext(context.Background())
}

func (i *ipaccessControlPtrType) ToIPAccessControlPtrOutputWithContext(ctx context.Context) IPAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPAccessControlPtrOutput)
}

type IPAccessControlOutput struct{ *pulumi.OutputState }

func (IPAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAccessControl)(nil)).Elem()
}

func (o IPAccessControlOutput) ToIPAccessControlOutput() IPAccessControlOutput {
	return o
}

func (o IPAccessControlOutput) ToIPAccessControlOutputWithContext(ctx context.Context) IPAccessControlOutput {
	return o
}

func (o IPAccessControlOutput) ToIPAccessControlPtrOutput() IPAccessControlPtrOutput {
	return o.ToIPAccessControlPtrOutputWithContext(context.Background())
}

func (o IPAccessControlOutput) ToIPAccessControlPtrOutputWithContext(ctx context.Context) IPAccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IPAccessControl) *IPAccessControl {
		return &v
	}).(IPAccessControlPtrOutput)
}

func (o IPAccessControlOutput) Allow() IPRangeArrayOutput {
	return o.ApplyT(func(v IPAccessControl) []IPRange { return v.Allow }).(IPRangeArrayOutput)
}

type IPAccessControlPtrOutput struct{ *pulumi.OutputState }

func (IPAccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPAccessControl)(nil)).Elem()
}

func (o IPAccessControlPtrOutput) ToIPAccessControlPtrOutput() IPAccessControlPtrOutput {
	return o
}

func (o IPAccessControlPtrOutput) ToIPAccessControlPtrOutputWithContext(ctx context.Context) IPAccessControlPtrOutput {
	return o
}

func (o IPAccessControlPtrOutput) Elem() IPAccessControlOutput {
	return o.ApplyT(func(v *IPAccessControl) IPAccessControl {
		if v != nil {
			return *v
		}
		var ret IPAccessControl
		return ret
	}).(IPAccessControlOutput)
}

func (o IPAccessControlPtrOutput) Allow() IPRangeArrayOutput {
	return o.ApplyT(func(v *IPAccessControl) []IPRange {
		if v == nil {
			return nil
		}
		return v.Allow
	}).(IPRangeArrayOutput)
}

type IPAccessControlResponse struct {
	Allow []IPRangeResponse `pulumi:"allow"`
}

type IPAccessControlResponseOutput struct{ *pulumi.OutputState }

func (IPAccessControlResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAccessControlResponse)(nil)).Elem()
}

func (o IPAccessControlResponseOutput) ToIPAccessControlResponseOutput() IPAccessControlResponseOutput {
	return o
}

func (o IPAccessControlResponseOutput) ToIPAccessControlResponseOutputWithContext(ctx context.Context) IPAccessControlResponseOutput {
	return o
}

func (o IPAccessControlResponseOutput) Allow() IPRangeResponseArrayOutput {
	return o.ApplyT(func(v IPAccessControlResponse) []IPRangeResponse { return v.Allow }).(IPRangeResponseArrayOutput)
}

type IPAccessControlResponsePtrOutput struct{ *pulumi.OutputState }

func (IPAccessControlResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPAccessControlResponse)(nil)).Elem()
}

func (o IPAccessControlResponsePtrOutput) ToIPAccessControlResponsePtrOutput() IPAccessControlResponsePtrOutput {
	return o
}

func (o IPAccessControlResponsePtrOutput) ToIPAccessControlResponsePtrOutputWithContext(ctx context.Context) IPAccessControlResponsePtrOutput {
	return o
}

func (o IPAccessControlResponsePtrOutput) Elem() IPAccessControlResponseOutput {
	return o.ApplyT(func(v *IPAccessControlResponse) IPAccessControlResponse {
		if v != nil {
			return *v
		}
		var ret IPAccessControlResponse
		return ret
	}).(IPAccessControlResponseOutput)
}

func (o IPAccessControlResponsePtrOutput) Allow() IPRangeResponseArrayOutput {
	return o.ApplyT(func(v *IPAccessControlResponse) []IPRangeResponse {
		if v == nil {
			return nil
		}
		return v.Allow
	}).(IPRangeResponseArrayOutput)
}

type IPRange struct {
	Address            *string `pulumi:"address"`
	Name               *string `pulumi:"name"`
	SubnetPrefixLength *int    `pulumi:"subnetPrefixLength"`
}





type IPRangeInput interface {
	pulumi.Input

	ToIPRangeOutput() IPRangeOutput
	ToIPRangeOutputWithContext(context.Context) IPRangeOutput
}

type IPRangeArgs struct {
	Address            pulumi.StringPtrInput `pulumi:"address"`
	Name               pulumi.StringPtrInput `pulumi:"name"`
	SubnetPrefixLength pulumi.IntPtrInput    `pulumi:"subnetPrefixLength"`
}

func (IPRangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRange)(nil)).Elem()
}

func (i IPRangeArgs) ToIPRangeOutput() IPRangeOutput {
	return i.ToIPRangeOutputWithContext(context.Background())
}

func (i IPRangeArgs) ToIPRangeOutputWithContext(ctx context.Context) IPRangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRangeOutput)
}





type IPRangeArrayInput interface {
	pulumi.Input

	ToIPRangeArrayOutput() IPRangeArrayOutput
	ToIPRangeArrayOutputWithContext(context.Context) IPRangeArrayOutput
}

type IPRangeArray []IPRangeInput

func (IPRangeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRange)(nil)).Elem()
}

func (i IPRangeArray) ToIPRangeArrayOutput() IPRangeArrayOutput {
	return i.ToIPRangeArrayOutputWithContext(context.Background())
}

func (i IPRangeArray) ToIPRangeArrayOutputWithContext(ctx context.Context) IPRangeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRangeArrayOutput)
}

type IPRangeOutput struct{ *pulumi.OutputState }

func (IPRangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRange)(nil)).Elem()
}

func (o IPRangeOutput) ToIPRangeOutput() IPRangeOutput {
	return o
}

func (o IPRangeOutput) ToIPRangeOutputWithContext(ctx context.Context) IPRangeOutput {
	return o
}

func (o IPRangeOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRange) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o IPRangeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRange) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IPRangeOutput) SubnetPrefixLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IPRange) *int { return v.SubnetPrefixLength }).(pulumi.IntPtrOutput)
}

type IPRangeArrayOutput struct{ *pulumi.OutputState }

func (IPRangeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRange)(nil)).Elem()
}

func (o IPRangeArrayOutput) ToIPRangeArrayOutput() IPRangeArrayOutput {
	return o
}

func (o IPRangeArrayOutput) ToIPRangeArrayOutputWithContext(ctx context.Context) IPRangeArrayOutput {
	return o
}

func (o IPRangeArrayOutput) Index(i pulumi.IntInput) IPRangeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRange {
		return vs[0].([]IPRange)[vs[1].(int)]
	}).(IPRangeOutput)
}

type IPRangeResponse struct {
	Address            *string `pulumi:"address"`
	Name               *string `pulumi:"name"`
	SubnetPrefixLength *int    `pulumi:"subnetPrefixLength"`
}

type IPRangeResponseOutput struct{ *pulumi.OutputState }

func (IPRangeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRangeResponse)(nil)).Elem()
}

func (o IPRangeResponseOutput) ToIPRangeResponseOutput() IPRangeResponseOutput {
	return o
}

func (o IPRangeResponseOutput) ToIPRangeResponseOutputWithContext(ctx context.Context) IPRangeResponseOutput {
	return o
}

func (o IPRangeResponseOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRangeResponse) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o IPRangeResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRangeResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IPRangeResponseOutput) SubnetPrefixLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IPRangeResponse) *int { return v.SubnetPrefixLength }).(pulumi.IntPtrOutput)
}

type IPRangeResponseArrayOutput struct{ *pulumi.OutputState }

func (IPRangeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRangeResponse)(nil)).Elem()
}

func (o IPRangeResponseArrayOutput) ToIPRangeResponseArrayOutput() IPRangeResponseArrayOutput {
	return o
}

func (o IPRangeResponseArrayOutput) ToIPRangeResponseArrayOutputWithContext(ctx context.Context) IPRangeResponseArrayOutput {
	return o
}

func (o IPRangeResponseArrayOutput) Index(i pulumi.IntInput) IPRangeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRangeResponse {
		return vs[0].([]IPRangeResponse)[vs[1].(int)]
	}).(IPRangeResponseOutput)
}

type Image struct {
	KeyFrameInterval *string `pulumi:"keyFrameInterval"`
	Label            *string `pulumi:"label"`
	OdataType        string  `pulumi:"odataType"`
	Range            *string `pulumi:"range"`
	Start            *string `pulumi:"start"`
	Step             *string `pulumi:"step"`
	StretchMode      *string `pulumi:"stretchMode"`
}

type ImageFormat struct {
	FilenamePattern *string `pulumi:"filenamePattern"`
	OdataType       string  `pulumi:"odataType"`
}

type ImageFormatResponse struct {
	FilenamePattern *string `pulumi:"filenamePattern"`
	OdataType       string  `pulumi:"odataType"`
}

type ImageResponse struct {
	KeyFrameInterval *string `pulumi:"keyFrameInterval"`
	Label            *string `pulumi:"label"`
	OdataType        string  `pulumi:"odataType"`
	Range            *string `pulumi:"range"`
	Start            *string `pulumi:"start"`
	Step             *string `pulumi:"step"`
	StretchMode      *string `pulumi:"stretchMode"`
}

type JobErrorDetailResponse struct {
	Code    string `pulumi:"code"`
	Message string `pulumi:"message"`
}

type JobErrorDetailResponseOutput struct{ *pulumi.OutputState }

func (JobErrorDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobErrorDetailResponse)(nil)).Elem()
}

func (o JobErrorDetailResponseOutput) ToJobErrorDetailResponseOutput() JobErrorDetailResponseOutput {
	return o
}

func (o JobErrorDetailResponseOutput) ToJobErrorDetailResponseOutputWithContext(ctx context.Context) JobErrorDetailResponseOutput {
	return o
}

func (o JobErrorDetailResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorDetailResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o JobErrorDetailResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorDetailResponse) string { return v.Message }).(pulumi.StringOutput)
}

type JobErrorDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (JobErrorDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobErrorDetailResponse)(nil)).Elem()
}

func (o JobErrorDetailResponseArrayOutput) ToJobErrorDetailResponseArrayOutput() JobErrorDetailResponseArrayOutput {
	return o
}

func (o JobErrorDetailResponseArrayOutput) ToJobErrorDetailResponseArrayOutputWithContext(ctx context.Context) JobErrorDetailResponseArrayOutput {
	return o
}

func (o JobErrorDetailResponseArrayOutput) Index(i pulumi.IntInput) JobErrorDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobErrorDetailResponse {
		return vs[0].([]JobErrorDetailResponse)[vs[1].(int)]
	}).(JobErrorDetailResponseOutput)
}

type JobErrorResponse struct {
	Category string                   `pulumi:"category"`
	Code     string                   `pulumi:"code"`
	Details  []JobErrorDetailResponse `pulumi:"details"`
	Message  string                   `pulumi:"message"`
	Retry    string                   `pulumi:"retry"`
}

type JobErrorResponseOutput struct{ *pulumi.OutputState }

func (JobErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobErrorResponse)(nil)).Elem()
}

func (o JobErrorResponseOutput) ToJobErrorResponseOutput() JobErrorResponseOutput {
	return o
}

func (o JobErrorResponseOutput) ToJobErrorResponseOutputWithContext(ctx context.Context) JobErrorResponseOutput {
	return o
}

func (o JobErrorResponseOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorResponse) string { return v.Category }).(pulumi.StringOutput)
}

func (o JobErrorResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o JobErrorResponseOutput) Details() JobErrorDetailResponseArrayOutput {
	return o.ApplyT(func(v JobErrorResponse) []JobErrorDetailResponse { return v.Details }).(JobErrorDetailResponseArrayOutput)
}

func (o JobErrorResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o JobErrorResponseOutput) Retry() pulumi.StringOutput {
	return o.ApplyT(func(v JobErrorResponse) string { return v.Retry }).(pulumi.StringOutput)
}

type JobInputAsset struct {
	AssetName string   `pulumi:"assetName"`
	Files     []string `pulumi:"files"`
	Label     *string  `pulumi:"label"`
	OdataType string   `pulumi:"odataType"`
}

type JobInputAssetResponse struct {
	AssetName string   `pulumi:"assetName"`
	Files     []string `pulumi:"files"`
	Label     *string  `pulumi:"label"`
	OdataType string   `pulumi:"odataType"`
}

type JobInputClip struct {
	Files     []string `pulumi:"files"`
	Label     *string  `pulumi:"label"`
	OdataType string   `pulumi:"odataType"`
}

type JobInputClipResponse struct {
	Files     []string `pulumi:"files"`
	Label     *string  `pulumi:"label"`
	OdataType string   `pulumi:"odataType"`
}

type JobInputHttp struct {
	BaseUri   *string  `pulumi:"baseUri"`
	Files     []string `pulumi:"files"`
	Label     *string  `pulumi:"label"`
	OdataType string   `pulumi:"odataType"`
}

type JobInputHttpResponse struct {
	BaseUri   *string  `pulumi:"baseUri"`
	Files     []string `pulumi:"files"`
	Label     *string  `pulumi:"label"`
	OdataType string   `pulumi:"odataType"`
}

type JobInputs struct {
	Inputs    []interface{} `pulumi:"inputs"`
	Label     *string       `pulumi:"label"`
	OdataType string        `pulumi:"odataType"`
}

type JobInputsResponse struct {
	Inputs    []interface{} `pulumi:"inputs"`
	Label     *string       `pulumi:"label"`
	OdataType string        `pulumi:"odataType"`
}

type JobOutputAsset struct {
	AssetName string `pulumi:"assetName"`
	OdataType string `pulumi:"odataType"`
}





type JobOutputAssetInput interface {
	pulumi.Input

	ToJobOutputAssetOutput() JobOutputAssetOutput
	ToJobOutputAssetOutputWithContext(context.Context) JobOutputAssetOutput
}

type JobOutputAssetArgs struct {
	AssetName pulumi.StringInput `pulumi:"assetName"`
	OdataType pulumi.StringInput `pulumi:"odataType"`
}

func (JobOutputAssetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobOutputAsset)(nil)).Elem()
}

func (i JobOutputAssetArgs) ToJobOutputAssetOutput() JobOutputAssetOutput {
	return i.ToJobOutputAssetOutputWithContext(context.Background())
}

func (i JobOutputAssetArgs) ToJobOutputAssetOutputWithContext(ctx context.Context) JobOutputAssetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutputAssetOutput)
}





type JobOutputAssetArrayInput interface {
	pulumi.Input

	ToJobOutputAssetArrayOutput() JobOutputAssetArrayOutput
	ToJobOutputAssetArrayOutputWithContext(context.Context) JobOutputAssetArrayOutput
}

type JobOutputAssetArray []JobOutputAssetInput

func (JobOutputAssetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobOutputAsset)(nil)).Elem()
}

func (i JobOutputAssetArray) ToJobOutputAssetArrayOutput() JobOutputAssetArrayOutput {
	return i.ToJobOutputAssetArrayOutputWithContext(context.Background())
}

func (i JobOutputAssetArray) ToJobOutputAssetArrayOutputWithContext(ctx context.Context) JobOutputAssetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutputAssetArrayOutput)
}

type JobOutputAssetOutput struct{ *pulumi.OutputState }

func (JobOutputAssetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobOutputAsset)(nil)).Elem()
}

func (o JobOutputAssetOutput) ToJobOutputAssetOutput() JobOutputAssetOutput {
	return o
}

func (o JobOutputAssetOutput) ToJobOutputAssetOutputWithContext(ctx context.Context) JobOutputAssetOutput {
	return o
}

func (o JobOutputAssetOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAsset) string { return v.AssetName }).(pulumi.StringOutput)
}

func (o JobOutputAssetOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAsset) string { return v.OdataType }).(pulumi.StringOutput)
}

type JobOutputAssetArrayOutput struct{ *pulumi.OutputState }

func (JobOutputAssetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobOutputAsset)(nil)).Elem()
}

func (o JobOutputAssetArrayOutput) ToJobOutputAssetArrayOutput() JobOutputAssetArrayOutput {
	return o
}

func (o JobOutputAssetArrayOutput) ToJobOutputAssetArrayOutputWithContext(ctx context.Context) JobOutputAssetArrayOutput {
	return o
}

func (o JobOutputAssetArrayOutput) Index(i pulumi.IntInput) JobOutputAssetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobOutputAsset {
		return vs[0].([]JobOutputAsset)[vs[1].(int)]
	}).(JobOutputAssetOutput)
}

type JobOutputAssetResponse struct {
	AssetName string           `pulumi:"assetName"`
	Error     JobErrorResponse `pulumi:"error"`
	OdataType string           `pulumi:"odataType"`
	Progress  int              `pulumi:"progress"`
	State     string           `pulumi:"state"`
}

type JobOutputAssetResponseOutput struct{ *pulumi.OutputState }

func (JobOutputAssetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobOutputAssetResponse)(nil)).Elem()
}

func (o JobOutputAssetResponseOutput) ToJobOutputAssetResponseOutput() JobOutputAssetResponseOutput {
	return o
}

func (o JobOutputAssetResponseOutput) ToJobOutputAssetResponseOutputWithContext(ctx context.Context) JobOutputAssetResponseOutput {
	return o
}

func (o JobOutputAssetResponseOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) string { return v.AssetName }).(pulumi.StringOutput)
}

func (o JobOutputAssetResponseOutput) Error() JobErrorResponseOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) JobErrorResponse { return v.Error }).(JobErrorResponseOutput)
}

func (o JobOutputAssetResponseOutput) OdataType() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) string { return v.OdataType }).(pulumi.StringOutput)
}

func (o JobOutputAssetResponseOutput) Progress() pulumi.IntOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) int { return v.Progress }).(pulumi.IntOutput)
}

func (o JobOutputAssetResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v JobOutputAssetResponse) string { return v.State }).(pulumi.StringOutput)
}

type JobOutputAssetResponseArrayOutput struct{ *pulumi.OutputState }

func (JobOutputAssetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobOutputAssetResponse)(nil)).Elem()
}

func (o JobOutputAssetResponseArrayOutput) ToJobOutputAssetResponseArrayOutput() JobOutputAssetResponseArrayOutput {
	return o
}

func (o JobOutputAssetResponseArrayOutput) ToJobOutputAssetResponseArrayOutputWithContext(ctx context.Context) JobOutputAssetResponseArrayOutput {
	return o
}

func (o JobOutputAssetResponseArrayOutput) Index(i pulumi.IntInput) JobOutputAssetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobOutputAssetResponse {
		return vs[0].([]JobOutputAssetResponse)[vs[1].(int)]
	}).(JobOutputAssetResponseOutput)
}

type JpgFormat struct {
	FilenamePattern *string `pulumi:"filenamePattern"`
	OdataType       string  `pulumi:"odataType"`
}

type JpgFormatResponse struct {
	FilenamePattern *string `pulumi:"filenamePattern"`
	OdataType       string  `pulumi:"odataType"`
}

type JpgImage struct {
	KeyFrameInterval *string    `pulumi:"keyFrameInterval"`
	Label            *string    `pulumi:"label"`
	Layers           []JpgLayer `pulumi:"layers"`
	OdataType        string     `pulumi:"odataType"`
	Range            *string    `pulumi:"range"`
	Start            *string    `pulumi:"start"`
	Step             *string    `pulumi:"step"`
	StretchMode      *string    `pulumi:"stretchMode"`
}

type JpgImageResponse struct {
	KeyFrameInterval *string            `pulumi:"keyFrameInterval"`
	Label            *string            `pulumi:"label"`
	Layers           []JpgLayerResponse `pulumi:"layers"`
	OdataType        string             `pulumi:"odataType"`
	Range            *string            `pulumi:"range"`
	Start            *string            `pulumi:"start"`
	Step             *string            `pulumi:"step"`
	StretchMode      *string            `pulumi:"stretchMode"`
}

type JpgLayer struct {
	Height    *string `pulumi:"height"`
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
	Quality   *int    `pulumi:"quality"`
	Width     *string `pulumi:"width"`
}

type JpgLayerResponse struct {
	Height    *string `pulumi:"height"`
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
	Quality   *int    `pulumi:"quality"`
	Width     *string `pulumi:"width"`
}

type LiveEventEncoding struct {
	EncodingType *LiveEventEncodingType `pulumi:"encodingType"`
	PresetName   *string                `pulumi:"presetName"`
}





type LiveEventEncodingInput interface {
	pulumi.Input

	ToLiveEventEncodingOutput() LiveEventEncodingOutput
	ToLiveEventEncodingOutputWithContext(context.Context) LiveEventEncodingOutput
}

type LiveEventEncodingArgs struct {
	EncodingType LiveEventEncodingTypePtrInput `pulumi:"encodingType"`
	PresetName   pulumi.StringPtrInput         `pulumi:"presetName"`
}

func (LiveEventEncodingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEncoding)(nil)).Elem()
}

func (i LiveEventEncodingArgs) ToLiveEventEncodingOutput() LiveEventEncodingOutput {
	return i.ToLiveEventEncodingOutputWithContext(context.Background())
}

func (i LiveEventEncodingArgs) ToLiveEventEncodingOutputWithContext(ctx context.Context) LiveEventEncodingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEncodingOutput)
}

func (i LiveEventEncodingArgs) ToLiveEventEncodingPtrOutput() LiveEventEncodingPtrOutput {
	return i.ToLiveEventEncodingPtrOutputWithContext(context.Background())
}

func (i LiveEventEncodingArgs) ToLiveEventEncodingPtrOutputWithContext(ctx context.Context) LiveEventEncodingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEncodingOutput).ToLiveEventEncodingPtrOutputWithContext(ctx)
}









type LiveEventEncodingPtrInput interface {
	pulumi.Input

	ToLiveEventEncodingPtrOutput() LiveEventEncodingPtrOutput
	ToLiveEventEncodingPtrOutputWithContext(context.Context) LiveEventEncodingPtrOutput
}

type liveEventEncodingPtrType LiveEventEncodingArgs

func LiveEventEncodingPtr(v *LiveEventEncodingArgs) LiveEventEncodingPtrInput {
	return (*liveEventEncodingPtrType)(v)
}

func (*liveEventEncodingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventEncoding)(nil)).Elem()
}

func (i *liveEventEncodingPtrType) ToLiveEventEncodingPtrOutput() LiveEventEncodingPtrOutput {
	return i.ToLiveEventEncodingPtrOutputWithContext(context.Background())
}

func (i *liveEventEncodingPtrType) ToLiveEventEncodingPtrOutputWithContext(ctx context.Context) LiveEventEncodingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEncodingPtrOutput)
}

type LiveEventEncodingOutput struct{ *pulumi.OutputState }

func (LiveEventEncodingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEncoding)(nil)).Elem()
}

func (o LiveEventEncodingOutput) ToLiveEventEncodingOutput() LiveEventEncodingOutput {
	return o
}

func (o LiveEventEncodingOutput) ToLiveEventEncodingOutputWithContext(ctx context.Context) LiveEventEncodingOutput {
	return o
}

func (o LiveEventEncodingOutput) ToLiveEventEncodingPtrOutput() LiveEventEncodingPtrOutput {
	return o.ToLiveEventEncodingPtrOutputWithContext(context.Background())
}

func (o LiveEventEncodingOutput) ToLiveEventEncodingPtrOutputWithContext(ctx context.Context) LiveEventEncodingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventEncoding) *LiveEventEncoding {
		return &v
	}).(LiveEventEncodingPtrOutput)
}

func (o LiveEventEncodingOutput) EncodingType() LiveEventEncodingTypePtrOutput {
	return o.ApplyT(func(v LiveEventEncoding) *LiveEventEncodingType { return v.EncodingType }).(LiveEventEncodingTypePtrOutput)
}

func (o LiveEventEncodingOutput) PresetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncoding) *string { return v.PresetName }).(pulumi.StringPtrOutput)
}

type LiveEventEncodingPtrOutput struct{ *pulumi.OutputState }

func (LiveEventEncodingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventEncoding)(nil)).Elem()
}

func (o LiveEventEncodingPtrOutput) ToLiveEventEncodingPtrOutput() LiveEventEncodingPtrOutput {
	return o
}

func (o LiveEventEncodingPtrOutput) ToLiveEventEncodingPtrOutputWithContext(ctx context.Context) LiveEventEncodingPtrOutput {
	return o
}

func (o LiveEventEncodingPtrOutput) Elem() LiveEventEncodingOutput {
	return o.ApplyT(func(v *LiveEventEncoding) LiveEventEncoding {
		if v != nil {
			return *v
		}
		var ret LiveEventEncoding
		return ret
	}).(LiveEventEncodingOutput)
}

func (o LiveEventEncodingPtrOutput) EncodingType() LiveEventEncodingTypePtrOutput {
	return o.ApplyT(func(v *LiveEventEncoding) *LiveEventEncodingType {
		if v == nil {
			return nil
		}
		return v.EncodingType
	}).(LiveEventEncodingTypePtrOutput)
}

func (o LiveEventEncodingPtrOutput) PresetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncoding) *string {
		if v == nil {
			return nil
		}
		return v.PresetName
	}).(pulumi.StringPtrOutput)
}

type LiveEventEncodingResponse struct {
	EncodingType *string `pulumi:"encodingType"`
	PresetName   *string `pulumi:"presetName"`
}

type LiveEventEncodingResponseOutput struct{ *pulumi.OutputState }

func (LiveEventEncodingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEncodingResponse)(nil)).Elem()
}

func (o LiveEventEncodingResponseOutput) ToLiveEventEncodingResponseOutput() LiveEventEncodingResponseOutput {
	return o
}

func (o LiveEventEncodingResponseOutput) ToLiveEventEncodingResponseOutputWithContext(ctx context.Context) LiveEventEncodingResponseOutput {
	return o
}

func (o LiveEventEncodingResponseOutput) EncodingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncodingResponse) *string { return v.EncodingType }).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingResponseOutput) PresetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEncodingResponse) *string { return v.PresetName }).(pulumi.StringPtrOutput)
}

type LiveEventEncodingResponsePtrOutput struct{ *pulumi.OutputState }

func (LiveEventEncodingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventEncodingResponse)(nil)).Elem()
}

func (o LiveEventEncodingResponsePtrOutput) ToLiveEventEncodingResponsePtrOutput() LiveEventEncodingResponsePtrOutput {
	return o
}

func (o LiveEventEncodingResponsePtrOutput) ToLiveEventEncodingResponsePtrOutputWithContext(ctx context.Context) LiveEventEncodingResponsePtrOutput {
	return o
}

func (o LiveEventEncodingResponsePtrOutput) Elem() LiveEventEncodingResponseOutput {
	return o.ApplyT(func(v *LiveEventEncodingResponse) LiveEventEncodingResponse {
		if v != nil {
			return *v
		}
		var ret LiveEventEncodingResponse
		return ret
	}).(LiveEventEncodingResponseOutput)
}

func (o LiveEventEncodingResponsePtrOutput) EncodingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncodingResponse) *string {
		if v == nil {
			return nil
		}
		return v.EncodingType
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventEncodingResponsePtrOutput) PresetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventEncodingResponse) *string {
		if v == nil {
			return nil
		}
		return v.PresetName
	}).(pulumi.StringPtrOutput)
}

type LiveEventEndpoint struct {
	Protocol *string `pulumi:"protocol"`
	Url      *string `pulumi:"url"`
}





type LiveEventEndpointInput interface {
	pulumi.Input

	ToLiveEventEndpointOutput() LiveEventEndpointOutput
	ToLiveEventEndpointOutputWithContext(context.Context) LiveEventEndpointOutput
}

type LiveEventEndpointArgs struct {
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
	Url      pulumi.StringPtrInput `pulumi:"url"`
}

func (LiveEventEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEndpoint)(nil)).Elem()
}

func (i LiveEventEndpointArgs) ToLiveEventEndpointOutput() LiveEventEndpointOutput {
	return i.ToLiveEventEndpointOutputWithContext(context.Background())
}

func (i LiveEventEndpointArgs) ToLiveEventEndpointOutputWithContext(ctx context.Context) LiveEventEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEndpointOutput)
}





type LiveEventEndpointArrayInput interface {
	pulumi.Input

	ToLiveEventEndpointArrayOutput() LiveEventEndpointArrayOutput
	ToLiveEventEndpointArrayOutputWithContext(context.Context) LiveEventEndpointArrayOutput
}

type LiveEventEndpointArray []LiveEventEndpointInput

func (LiveEventEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventEndpoint)(nil)).Elem()
}

func (i LiveEventEndpointArray) ToLiveEventEndpointArrayOutput() LiveEventEndpointArrayOutput {
	return i.ToLiveEventEndpointArrayOutputWithContext(context.Background())
}

func (i LiveEventEndpointArray) ToLiveEventEndpointArrayOutputWithContext(ctx context.Context) LiveEventEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventEndpointArrayOutput)
}

type LiveEventEndpointOutput struct{ *pulumi.OutputState }

func (LiveEventEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEndpoint)(nil)).Elem()
}

func (o LiveEventEndpointOutput) ToLiveEventEndpointOutput() LiveEventEndpointOutput {
	return o
}

func (o LiveEventEndpointOutput) ToLiveEventEndpointOutputWithContext(ctx context.Context) LiveEventEndpointOutput {
	return o
}

func (o LiveEventEndpointOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEndpoint) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o LiveEventEndpointOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEndpoint) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type LiveEventEndpointArrayOutput struct{ *pulumi.OutputState }

func (LiveEventEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventEndpoint)(nil)).Elem()
}

func (o LiveEventEndpointArrayOutput) ToLiveEventEndpointArrayOutput() LiveEventEndpointArrayOutput {
	return o
}

func (o LiveEventEndpointArrayOutput) ToLiveEventEndpointArrayOutputWithContext(ctx context.Context) LiveEventEndpointArrayOutput {
	return o
}

func (o LiveEventEndpointArrayOutput) Index(i pulumi.IntInput) LiveEventEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiveEventEndpoint {
		return vs[0].([]LiveEventEndpoint)[vs[1].(int)]
	}).(LiveEventEndpointOutput)
}

type LiveEventEndpointResponse struct {
	Protocol *string `pulumi:"protocol"`
	Url      *string `pulumi:"url"`
}

type LiveEventEndpointResponseOutput struct{ *pulumi.OutputState }

func (LiveEventEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventEndpointResponse)(nil)).Elem()
}

func (o LiveEventEndpointResponseOutput) ToLiveEventEndpointResponseOutput() LiveEventEndpointResponseOutput {
	return o
}

func (o LiveEventEndpointResponseOutput) ToLiveEventEndpointResponseOutputWithContext(ctx context.Context) LiveEventEndpointResponseOutput {
	return o
}

func (o LiveEventEndpointResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEndpointResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o LiveEventEndpointResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventEndpointResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type LiveEventEndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (LiveEventEndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveEventEndpointResponse)(nil)).Elem()
}

func (o LiveEventEndpointResponseArrayOutput) ToLiveEventEndpointResponseArrayOutput() LiveEventEndpointResponseArrayOutput {
	return o
}

func (o LiveEventEndpointResponseArrayOutput) ToLiveEventEndpointResponseArrayOutputWithContext(ctx context.Context) LiveEventEndpointResponseArrayOutput {
	return o
}

func (o LiveEventEndpointResponseArrayOutput) Index(i pulumi.IntInput) LiveEventEndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiveEventEndpointResponse {
		return vs[0].([]LiveEventEndpointResponse)[vs[1].(int)]
	}).(LiveEventEndpointResponseOutput)
}

type LiveEventInputType struct {
	AccessToken              *string                `pulumi:"accessToken"`
	Endpoints                []LiveEventEndpoint    `pulumi:"endpoints"`
	KeyFrameIntervalDuration *string                `pulumi:"keyFrameIntervalDuration"`
	StreamingProtocol        LiveEventInputProtocol `pulumi:"streamingProtocol"`
}





type LiveEventInputTypeInput interface {
	pulumi.Input

	ToLiveEventInputTypeOutput() LiveEventInputTypeOutput
	ToLiveEventInputTypeOutputWithContext(context.Context) LiveEventInputTypeOutput
}

type LiveEventInputTypeArgs struct {
	AccessToken              pulumi.StringPtrInput       `pulumi:"accessToken"`
	Endpoints                LiveEventEndpointArrayInput `pulumi:"endpoints"`
	KeyFrameIntervalDuration pulumi.StringPtrInput       `pulumi:"keyFrameIntervalDuration"`
	StreamingProtocol        LiveEventInputProtocolInput `pulumi:"streamingProtocol"`
}

func (LiveEventInputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputType)(nil)).Elem()
}

func (i LiveEventInputTypeArgs) ToLiveEventInputTypeOutput() LiveEventInputTypeOutput {
	return i.ToLiveEventInputTypeOutputWithContext(context.Background())
}

func (i LiveEventInputTypeArgs) ToLiveEventInputTypeOutputWithContext(ctx context.Context) LiveEventInputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventInputTypeOutput)
}

type LiveEventInputTypeOutput struct{ *pulumi.OutputState }

func (LiveEventInputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputType)(nil)).Elem()
}

func (o LiveEventInputTypeOutput) ToLiveEventInputTypeOutput() LiveEventInputTypeOutput {
	return o
}

func (o LiveEventInputTypeOutput) ToLiveEventInputTypeOutputWithContext(ctx context.Context) LiveEventInputTypeOutput {
	return o
}

func (o LiveEventInputTypeOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputType) *string { return v.AccessToken }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputTypeOutput) Endpoints() LiveEventEndpointArrayOutput {
	return o.ApplyT(func(v LiveEventInputType) []LiveEventEndpoint { return v.Endpoints }).(LiveEventEndpointArrayOutput)
}

func (o LiveEventInputTypeOutput) KeyFrameIntervalDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputType) *string { return v.KeyFrameIntervalDuration }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputTypeOutput) StreamingProtocol() LiveEventInputProtocolOutput {
	return o.ApplyT(func(v LiveEventInputType) LiveEventInputProtocol { return v.StreamingProtocol }).(LiveEventInputProtocolOutput)
}

type LiveEventInputResponse struct {
	AccessToken              *string                     `pulumi:"accessToken"`
	Endpoints                []LiveEventEndpointResponse `pulumi:"endpoints"`
	KeyFrameIntervalDuration *string                     `pulumi:"keyFrameIntervalDuration"`
	StreamingProtocol        string                      `pulumi:"streamingProtocol"`
}

type LiveEventInputResponseOutput struct{ *pulumi.OutputState }

func (LiveEventInputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventInputResponse)(nil)).Elem()
}

func (o LiveEventInputResponseOutput) ToLiveEventInputResponseOutput() LiveEventInputResponseOutput {
	return o
}

func (o LiveEventInputResponseOutput) ToLiveEventInputResponseOutputWithContext(ctx context.Context) LiveEventInputResponseOutput {
	return o
}

func (o LiveEventInputResponseOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputResponse) *string { return v.AccessToken }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputResponseOutput) Endpoints() LiveEventEndpointResponseArrayOutput {
	return o.ApplyT(func(v LiveEventInputResponse) []LiveEventEndpointResponse { return v.Endpoints }).(LiveEventEndpointResponseArrayOutput)
}

func (o LiveEventInputResponseOutput) KeyFrameIntervalDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventInputResponse) *string { return v.KeyFrameIntervalDuration }).(pulumi.StringPtrOutput)
}

func (o LiveEventInputResponseOutput) StreamingProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LiveEventInputResponse) string { return v.StreamingProtocol }).(pulumi.StringOutput)
}

type LiveEventPreview struct {
	AccessControl       *LiveEventPreviewAccessControl `pulumi:"accessControl"`
	AlternativeMediaId  *string                        `pulumi:"alternativeMediaId"`
	Endpoints           []LiveEventEndpoint            `pulumi:"endpoints"`
	PreviewLocator      *string                        `pulumi:"previewLocator"`
	StreamingPolicyName *string                        `pulumi:"streamingPolicyName"`
}





type LiveEventPreviewInput interface {
	pulumi.Input

	ToLiveEventPreviewOutput() LiveEventPreviewOutput
	ToLiveEventPreviewOutputWithContext(context.Context) LiveEventPreviewOutput
}

type LiveEventPreviewArgs struct {
	AccessControl       LiveEventPreviewAccessControlPtrInput `pulumi:"accessControl"`
	AlternativeMediaId  pulumi.StringPtrInput                 `pulumi:"alternativeMediaId"`
	Endpoints           LiveEventEndpointArrayInput           `pulumi:"endpoints"`
	PreviewLocator      pulumi.StringPtrInput                 `pulumi:"previewLocator"`
	StreamingPolicyName pulumi.StringPtrInput                 `pulumi:"streamingPolicyName"`
}

func (LiveEventPreviewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreview)(nil)).Elem()
}

func (i LiveEventPreviewArgs) ToLiveEventPreviewOutput() LiveEventPreviewOutput {
	return i.ToLiveEventPreviewOutputWithContext(context.Background())
}

func (i LiveEventPreviewArgs) ToLiveEventPreviewOutputWithContext(ctx context.Context) LiveEventPreviewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewOutput)
}

func (i LiveEventPreviewArgs) ToLiveEventPreviewPtrOutput() LiveEventPreviewPtrOutput {
	return i.ToLiveEventPreviewPtrOutputWithContext(context.Background())
}

func (i LiveEventPreviewArgs) ToLiveEventPreviewPtrOutputWithContext(ctx context.Context) LiveEventPreviewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewOutput).ToLiveEventPreviewPtrOutputWithContext(ctx)
}









type LiveEventPreviewPtrInput interface {
	pulumi.Input

	ToLiveEventPreviewPtrOutput() LiveEventPreviewPtrOutput
	ToLiveEventPreviewPtrOutputWithContext(context.Context) LiveEventPreviewPtrOutput
}

type liveEventPreviewPtrType LiveEventPreviewArgs

func LiveEventPreviewPtr(v *LiveEventPreviewArgs) LiveEventPreviewPtrInput {
	return (*liveEventPreviewPtrType)(v)
}

func (*liveEventPreviewPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreview)(nil)).Elem()
}

func (i *liveEventPreviewPtrType) ToLiveEventPreviewPtrOutput() LiveEventPreviewPtrOutput {
	return i.ToLiveEventPreviewPtrOutputWithContext(context.Background())
}

func (i *liveEventPreviewPtrType) ToLiveEventPreviewPtrOutputWithContext(ctx context.Context) LiveEventPreviewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewPtrOutput)
}

type LiveEventPreviewOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreview)(nil)).Elem()
}

func (o LiveEventPreviewOutput) ToLiveEventPreviewOutput() LiveEventPreviewOutput {
	return o
}

func (o LiveEventPreviewOutput) ToLiveEventPreviewOutputWithContext(ctx context.Context) LiveEventPreviewOutput {
	return o
}

func (o LiveEventPreviewOutput) ToLiveEventPreviewPtrOutput() LiveEventPreviewPtrOutput {
	return o.ToLiveEventPreviewPtrOutputWithContext(context.Background())
}

func (o LiveEventPreviewOutput) ToLiveEventPreviewPtrOutputWithContext(ctx context.Context) LiveEventPreviewPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventPreview) *LiveEventPreview {
		return &v
	}).(LiveEventPreviewPtrOutput)
}

func (o LiveEventPreviewOutput) AccessControl() LiveEventPreviewAccessControlPtrOutput {
	return o.ApplyT(func(v LiveEventPreview) *LiveEventPreviewAccessControl { return v.AccessControl }).(LiveEventPreviewAccessControlPtrOutput)
}

func (o LiveEventPreviewOutput) AlternativeMediaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventPreview) *string { return v.AlternativeMediaId }).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewOutput) Endpoints() LiveEventEndpointArrayOutput {
	return o.ApplyT(func(v LiveEventPreview) []LiveEventEndpoint { return v.Endpoints }).(LiveEventEndpointArrayOutput)
}

func (o LiveEventPreviewOutput) PreviewLocator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventPreview) *string { return v.PreviewLocator }).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewOutput) StreamingPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventPreview) *string { return v.StreamingPolicyName }).(pulumi.StringPtrOutput)
}

type LiveEventPreviewPtrOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreview)(nil)).Elem()
}

func (o LiveEventPreviewPtrOutput) ToLiveEventPreviewPtrOutput() LiveEventPreviewPtrOutput {
	return o
}

func (o LiveEventPreviewPtrOutput) ToLiveEventPreviewPtrOutputWithContext(ctx context.Context) LiveEventPreviewPtrOutput {
	return o
}

func (o LiveEventPreviewPtrOutput) Elem() LiveEventPreviewOutput {
	return o.ApplyT(func(v *LiveEventPreview) LiveEventPreview {
		if v != nil {
			return *v
		}
		var ret LiveEventPreview
		return ret
	}).(LiveEventPreviewOutput)
}

func (o LiveEventPreviewPtrOutput) AccessControl() LiveEventPreviewAccessControlPtrOutput {
	return o.ApplyT(func(v *LiveEventPreview) *LiveEventPreviewAccessControl {
		if v == nil {
			return nil
		}
		return v.AccessControl
	}).(LiveEventPreviewAccessControlPtrOutput)
}

func (o LiveEventPreviewPtrOutput) AlternativeMediaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventPreview) *string {
		if v == nil {
			return nil
		}
		return v.AlternativeMediaId
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewPtrOutput) Endpoints() LiveEventEndpointArrayOutput {
	return o.ApplyT(func(v *LiveEventPreview) []LiveEventEndpoint {
		if v == nil {
			return nil
		}
		return v.Endpoints
	}).(LiveEventEndpointArrayOutput)
}

func (o LiveEventPreviewPtrOutput) PreviewLocator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventPreview) *string {
		if v == nil {
			return nil
		}
		return v.PreviewLocator
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewPtrOutput) StreamingPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventPreview) *string {
		if v == nil {
			return nil
		}
		return v.StreamingPolicyName
	}).(pulumi.StringPtrOutput)
}

type LiveEventPreviewAccessControl struct {
	Ip *IPAccessControl `pulumi:"ip"`
}





type LiveEventPreviewAccessControlInput interface {
	pulumi.Input

	ToLiveEventPreviewAccessControlOutput() LiveEventPreviewAccessControlOutput
	ToLiveEventPreviewAccessControlOutputWithContext(context.Context) LiveEventPreviewAccessControlOutput
}

type LiveEventPreviewAccessControlArgs struct {
	Ip IPAccessControlPtrInput `pulumi:"ip"`
}

func (LiveEventPreviewAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreviewAccessControl)(nil)).Elem()
}

func (i LiveEventPreviewAccessControlArgs) ToLiveEventPreviewAccessControlOutput() LiveEventPreviewAccessControlOutput {
	return i.ToLiveEventPreviewAccessControlOutputWithContext(context.Background())
}

func (i LiveEventPreviewAccessControlArgs) ToLiveEventPreviewAccessControlOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewAccessControlOutput)
}

func (i LiveEventPreviewAccessControlArgs) ToLiveEventPreviewAccessControlPtrOutput() LiveEventPreviewAccessControlPtrOutput {
	return i.ToLiveEventPreviewAccessControlPtrOutputWithContext(context.Background())
}

func (i LiveEventPreviewAccessControlArgs) ToLiveEventPreviewAccessControlPtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewAccessControlOutput).ToLiveEventPreviewAccessControlPtrOutputWithContext(ctx)
}









type LiveEventPreviewAccessControlPtrInput interface {
	pulumi.Input

	ToLiveEventPreviewAccessControlPtrOutput() LiveEventPreviewAccessControlPtrOutput
	ToLiveEventPreviewAccessControlPtrOutputWithContext(context.Context) LiveEventPreviewAccessControlPtrOutput
}

type liveEventPreviewAccessControlPtrType LiveEventPreviewAccessControlArgs

func LiveEventPreviewAccessControlPtr(v *LiveEventPreviewAccessControlArgs) LiveEventPreviewAccessControlPtrInput {
	return (*liveEventPreviewAccessControlPtrType)(v)
}

func (*liveEventPreviewAccessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreviewAccessControl)(nil)).Elem()
}

func (i *liveEventPreviewAccessControlPtrType) ToLiveEventPreviewAccessControlPtrOutput() LiveEventPreviewAccessControlPtrOutput {
	return i.ToLiveEventPreviewAccessControlPtrOutputWithContext(context.Background())
}

func (i *liveEventPreviewAccessControlPtrType) ToLiveEventPreviewAccessControlPtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventPreviewAccessControlPtrOutput)
}

type LiveEventPreviewAccessControlOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreviewAccessControl)(nil)).Elem()
}

func (o LiveEventPreviewAccessControlOutput) ToLiveEventPreviewAccessControlOutput() LiveEventPreviewAccessControlOutput {
	return o
}

func (o LiveEventPreviewAccessControlOutput) ToLiveEventPreviewAccessControlOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlOutput {
	return o
}

func (o LiveEventPreviewAccessControlOutput) ToLiveEventPreviewAccessControlPtrOutput() LiveEventPreviewAccessControlPtrOutput {
	return o.ToLiveEventPreviewAccessControlPtrOutputWithContext(context.Background())
}

func (o LiveEventPreviewAccessControlOutput) ToLiveEventPreviewAccessControlPtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveEventPreviewAccessControl) *LiveEventPreviewAccessControl {
		return &v
	}).(LiveEventPreviewAccessControlPtrOutput)
}

func (o LiveEventPreviewAccessControlOutput) Ip() IPAccessControlPtrOutput {
	return o.ApplyT(func(v LiveEventPreviewAccessControl) *IPAccessControl { return v.Ip }).(IPAccessControlPtrOutput)
}

type LiveEventPreviewAccessControlPtrOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewAccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreviewAccessControl)(nil)).Elem()
}

func (o LiveEventPreviewAccessControlPtrOutput) ToLiveEventPreviewAccessControlPtrOutput() LiveEventPreviewAccessControlPtrOutput {
	return o
}

func (o LiveEventPreviewAccessControlPtrOutput) ToLiveEventPreviewAccessControlPtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlPtrOutput {
	return o
}

func (o LiveEventPreviewAccessControlPtrOutput) Elem() LiveEventPreviewAccessControlOutput {
	return o.ApplyT(func(v *LiveEventPreviewAccessControl) LiveEventPreviewAccessControl {
		if v != nil {
			return *v
		}
		var ret LiveEventPreviewAccessControl
		return ret
	}).(LiveEventPreviewAccessControlOutput)
}

func (o LiveEventPreviewAccessControlPtrOutput) Ip() IPAccessControlPtrOutput {
	return o.ApplyT(func(v *LiveEventPreviewAccessControl) *IPAccessControl {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(IPAccessControlPtrOutput)
}

type LiveEventPreviewAccessControlResponse struct {
	Ip *IPAccessControlResponse `pulumi:"ip"`
}

type LiveEventPreviewAccessControlResponseOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewAccessControlResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreviewAccessControlResponse)(nil)).Elem()
}

func (o LiveEventPreviewAccessControlResponseOutput) ToLiveEventPreviewAccessControlResponseOutput() LiveEventPreviewAccessControlResponseOutput {
	return o
}

func (o LiveEventPreviewAccessControlResponseOutput) ToLiveEventPreviewAccessControlResponseOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlResponseOutput {
	return o
}

func (o LiveEventPreviewAccessControlResponseOutput) Ip() IPAccessControlResponsePtrOutput {
	return o.ApplyT(func(v LiveEventPreviewAccessControlResponse) *IPAccessControlResponse { return v.Ip }).(IPAccessControlResponsePtrOutput)
}

type LiveEventPreviewAccessControlResponsePtrOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewAccessControlResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreviewAccessControlResponse)(nil)).Elem()
}

func (o LiveEventPreviewAccessControlResponsePtrOutput) ToLiveEventPreviewAccessControlResponsePtrOutput() LiveEventPreviewAccessControlResponsePtrOutput {
	return o
}

func (o LiveEventPreviewAccessControlResponsePtrOutput) ToLiveEventPreviewAccessControlResponsePtrOutputWithContext(ctx context.Context) LiveEventPreviewAccessControlResponsePtrOutput {
	return o
}

func (o LiveEventPreviewAccessControlResponsePtrOutput) Elem() LiveEventPreviewAccessControlResponseOutput {
	return o.ApplyT(func(v *LiveEventPreviewAccessControlResponse) LiveEventPreviewAccessControlResponse {
		if v != nil {
			return *v
		}
		var ret LiveEventPreviewAccessControlResponse
		return ret
	}).(LiveEventPreviewAccessControlResponseOutput)
}

func (o LiveEventPreviewAccessControlResponsePtrOutput) Ip() IPAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *LiveEventPreviewAccessControlResponse) *IPAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(IPAccessControlResponsePtrOutput)
}

type LiveEventPreviewResponse struct {
	AccessControl       *LiveEventPreviewAccessControlResponse `pulumi:"accessControl"`
	AlternativeMediaId  *string                                `pulumi:"alternativeMediaId"`
	Endpoints           []LiveEventEndpointResponse            `pulumi:"endpoints"`
	PreviewLocator      *string                                `pulumi:"previewLocator"`
	StreamingPolicyName *string                                `pulumi:"streamingPolicyName"`
}

type LiveEventPreviewResponseOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveEventPreviewResponse)(nil)).Elem()
}

func (o LiveEventPreviewResponseOutput) ToLiveEventPreviewResponseOutput() LiveEventPreviewResponseOutput {
	return o
}

func (o LiveEventPreviewResponseOutput) ToLiveEventPreviewResponseOutputWithContext(ctx context.Context) LiveEventPreviewResponseOutput {
	return o
}

func (o LiveEventPreviewResponseOutput) AccessControl() LiveEventPreviewAccessControlResponsePtrOutput {
	return o.ApplyT(func(v LiveEventPreviewResponse) *LiveEventPreviewAccessControlResponse { return v.AccessControl }).(LiveEventPreviewAccessControlResponsePtrOutput)
}

func (o LiveEventPreviewResponseOutput) AlternativeMediaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventPreviewResponse) *string { return v.AlternativeMediaId }).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewResponseOutput) Endpoints() LiveEventEndpointResponseArrayOutput {
	return o.ApplyT(func(v LiveEventPreviewResponse) []LiveEventEndpointResponse { return v.Endpoints }).(LiveEventEndpointResponseArrayOutput)
}

func (o LiveEventPreviewResponseOutput) PreviewLocator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventPreviewResponse) *string { return v.PreviewLocator }).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewResponseOutput) StreamingPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveEventPreviewResponse) *string { return v.StreamingPolicyName }).(pulumi.StringPtrOutput)
}

type LiveEventPreviewResponsePtrOutput struct{ *pulumi.OutputState }

func (LiveEventPreviewResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEventPreviewResponse)(nil)).Elem()
}

func (o LiveEventPreviewResponsePtrOutput) ToLiveEventPreviewResponsePtrOutput() LiveEventPreviewResponsePtrOutput {
	return o
}

func (o LiveEventPreviewResponsePtrOutput) ToLiveEventPreviewResponsePtrOutputWithContext(ctx context.Context) LiveEventPreviewResponsePtrOutput {
	return o
}

func (o LiveEventPreviewResponsePtrOutput) Elem() LiveEventPreviewResponseOutput {
	return o.ApplyT(func(v *LiveEventPreviewResponse) LiveEventPreviewResponse {
		if v != nil {
			return *v
		}
		var ret LiveEventPreviewResponse
		return ret
	}).(LiveEventPreviewResponseOutput)
}

func (o LiveEventPreviewResponsePtrOutput) AccessControl() LiveEventPreviewAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *LiveEventPreviewResponse) *LiveEventPreviewAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.AccessControl
	}).(LiveEventPreviewAccessControlResponsePtrOutput)
}

func (o LiveEventPreviewResponsePtrOutput) AlternativeMediaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventPreviewResponse) *string {
		if v == nil {
			return nil
		}
		return v.AlternativeMediaId
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewResponsePtrOutput) Endpoints() LiveEventEndpointResponseArrayOutput {
	return o.ApplyT(func(v *LiveEventPreviewResponse) []LiveEventEndpointResponse {
		if v == nil {
			return nil
		}
		return v.Endpoints
	}).(LiveEventEndpointResponseArrayOutput)
}

func (o LiveEventPreviewResponsePtrOutput) PreviewLocator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventPreviewResponse) *string {
		if v == nil {
			return nil
		}
		return v.PreviewLocator
	}).(pulumi.StringPtrOutput)
}

func (o LiveEventPreviewResponsePtrOutput) StreamingPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEventPreviewResponse) *string {
		if v == nil {
			return nil
		}
		return v.StreamingPolicyName
	}).(pulumi.StringPtrOutput)
}

type Mp4Format struct {
	FilenamePattern *string      `pulumi:"filenamePattern"`
	OdataType       string       `pulumi:"odataType"`
	OutputFiles     []OutputFile `pulumi:"outputFiles"`
}

type Mp4FormatResponse struct {
	FilenamePattern *string              `pulumi:"filenamePattern"`
	OdataType       string               `pulumi:"odataType"`
	OutputFiles     []OutputFileResponse `pulumi:"outputFiles"`
}

type MultiBitrateFormat struct {
	FilenamePattern *string      `pulumi:"filenamePattern"`
	OdataType       string       `pulumi:"odataType"`
	OutputFiles     []OutputFile `pulumi:"outputFiles"`
}

type MultiBitrateFormatResponse struct {
	FilenamePattern *string              `pulumi:"filenamePattern"`
	OdataType       string               `pulumi:"odataType"`
	OutputFiles     []OutputFileResponse `pulumi:"outputFiles"`
}

type NoEncryption struct {
	EnabledProtocols *EnabledProtocols `pulumi:"enabledProtocols"`
}





type NoEncryptionInput interface {
	pulumi.Input

	ToNoEncryptionOutput() NoEncryptionOutput
	ToNoEncryptionOutputWithContext(context.Context) NoEncryptionOutput
}

type NoEncryptionArgs struct {
	EnabledProtocols EnabledProtocolsPtrInput `pulumi:"enabledProtocols"`
}

func (NoEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NoEncryption)(nil)).Elem()
}

func (i NoEncryptionArgs) ToNoEncryptionOutput() NoEncryptionOutput {
	return i.ToNoEncryptionOutputWithContext(context.Background())
}

func (i NoEncryptionArgs) ToNoEncryptionOutputWithContext(ctx context.Context) NoEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoEncryptionOutput)
}

func (i NoEncryptionArgs) ToNoEncryptionPtrOutput() NoEncryptionPtrOutput {
	return i.ToNoEncryptionPtrOutputWithContext(context.Background())
}

func (i NoEncryptionArgs) ToNoEncryptionPtrOutputWithContext(ctx context.Context) NoEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoEncryptionOutput).ToNoEncryptionPtrOutputWithContext(ctx)
}









type NoEncryptionPtrInput interface {
	pulumi.Input

	ToNoEncryptionPtrOutput() NoEncryptionPtrOutput
	ToNoEncryptionPtrOutputWithContext(context.Context) NoEncryptionPtrOutput
}

type noEncryptionPtrType NoEncryptionArgs

func NoEncryptionPtr(v *NoEncryptionArgs) NoEncryptionPtrInput {
	return (*noEncryptionPtrType)(v)
}

func (*noEncryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NoEncryption)(nil)).Elem()
}

func (i *noEncryptionPtrType) ToNoEncryptionPtrOutput() NoEncryptionPtrOutput {
	return i.ToNoEncryptionPtrOutputWithContext(context.Background())
}

func (i *noEncryptionPtrType) ToNoEncryptionPtrOutputWithContext(ctx context.Context) NoEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoEncryptionPtrOutput)
}

type NoEncryptionOutput struct{ *pulumi.OutputState }

func (NoEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NoEncryption)(nil)).Elem()
}

func (o NoEncryptionOutput) ToNoEncryptionOutput() NoEncryptionOutput {
	return o
}

func (o NoEncryptionOutput) ToNoEncryptionOutputWithContext(ctx context.Context) NoEncryptionOutput {
	return o
}

func (o NoEncryptionOutput) ToNoEncryptionPtrOutput() NoEncryptionPtrOutput {
	return o.ToNoEncryptionPtrOutputWithContext(context.Background())
}

func (o NoEncryptionOutput) ToNoEncryptionPtrOutputWithContext(ctx context.Context) NoEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NoEncryption) *NoEncryption {
		return &v
	}).(NoEncryptionPtrOutput)
}

func (o NoEncryptionOutput) EnabledProtocols() EnabledProtocolsPtrOutput {
	return o.ApplyT(func(v NoEncryption) *EnabledProtocols { return v.EnabledProtocols }).(EnabledProtocolsPtrOutput)
}

type NoEncryptionPtrOutput struct{ *pulumi.OutputState }

func (NoEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NoEncryption)(nil)).Elem()
}

func (o NoEncryptionPtrOutput) ToNoEncryptionPtrOutput() NoEncryptionPtrOutput {
	return o
}

func (o NoEncryptionPtrOutput) ToNoEncryptionPtrOutputWithContext(ctx context.Context) NoEncryptionPtrOutput {
	return o
}

func (o NoEncryptionPtrOutput) Elem() NoEncryptionOutput {
	return o.ApplyT(func(v *NoEncryption) NoEncryption {
		if v != nil {
			return *v
		}
		var ret NoEncryption
		return ret
	}).(NoEncryptionOutput)
}

func (o NoEncryptionPtrOutput) EnabledProtocols() EnabledProtocolsPtrOutput {
	return o.ApplyT(func(v *NoEncryption) *EnabledProtocols {
		if v == nil {
			return nil
		}
		return v.EnabledProtocols
	}).(EnabledProtocolsPtrOutput)
}

type NoEncryptionResponse struct {
	EnabledProtocols *EnabledProtocolsResponse `pulumi:"enabledProtocols"`
}

type NoEncryptionResponseOutput struct{ *pulumi.OutputState }

func (NoEncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NoEncryptionResponse)(nil)).Elem()
}

func (o NoEncryptionResponseOutput) ToNoEncryptionResponseOutput() NoEncryptionResponseOutput {
	return o
}

func (o NoEncryptionResponseOutput) ToNoEncryptionResponseOutputWithContext(ctx context.Context) NoEncryptionResponseOutput {
	return o
}

func (o NoEncryptionResponseOutput) EnabledProtocols() EnabledProtocolsResponsePtrOutput {
	return o.ApplyT(func(v NoEncryptionResponse) *EnabledProtocolsResponse { return v.EnabledProtocols }).(EnabledProtocolsResponsePtrOutput)
}

type NoEncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (NoEncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NoEncryptionResponse)(nil)).Elem()
}

func (o NoEncryptionResponsePtrOutput) ToNoEncryptionResponsePtrOutput() NoEncryptionResponsePtrOutput {
	return o
}

func (o NoEncryptionResponsePtrOutput) ToNoEncryptionResponsePtrOutputWithContext(ctx context.Context) NoEncryptionResponsePtrOutput {
	return o
}

func (o NoEncryptionResponsePtrOutput) Elem() NoEncryptionResponseOutput {
	return o.ApplyT(func(v *NoEncryptionResponse) NoEncryptionResponse {
		if v != nil {
			return *v
		}
		var ret NoEncryptionResponse
		return ret
	}).(NoEncryptionResponseOutput)
}

func (o NoEncryptionResponsePtrOutput) EnabledProtocols() EnabledProtocolsResponsePtrOutput {
	return o.ApplyT(func(v *NoEncryptionResponse) *EnabledProtocolsResponse {
		if v == nil {
			return nil
		}
		return v.EnabledProtocols
	}).(EnabledProtocolsResponsePtrOutput)
}

type OutputFile struct {
	Labels []string `pulumi:"labels"`
}

type OutputFileResponse struct {
	Labels []string `pulumi:"labels"`
}

type PngFormat struct {
	FilenamePattern *string `pulumi:"filenamePattern"`
	OdataType       string  `pulumi:"odataType"`
}

type PngFormatResponse struct {
	FilenamePattern *string `pulumi:"filenamePattern"`
	OdataType       string  `pulumi:"odataType"`
}

type PngImage struct {
	KeyFrameInterval *string    `pulumi:"keyFrameInterval"`
	Label            *string    `pulumi:"label"`
	Layers           []PngLayer `pulumi:"layers"`
	OdataType        string     `pulumi:"odataType"`
	Range            *string    `pulumi:"range"`
	Start            *string    `pulumi:"start"`
	Step             *string    `pulumi:"step"`
	StretchMode      *string    `pulumi:"stretchMode"`
}

type PngImageResponse struct {
	KeyFrameInterval *string            `pulumi:"keyFrameInterval"`
	Label            *string            `pulumi:"label"`
	Layers           []PngLayerResponse `pulumi:"layers"`
	OdataType        string             `pulumi:"odataType"`
	Range            *string            `pulumi:"range"`
	Start            *string            `pulumi:"start"`
	Step             *string            `pulumi:"step"`
	StretchMode      *string            `pulumi:"stretchMode"`
}

type PngLayer struct {
	Height    *string `pulumi:"height"`
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
	Width     *string `pulumi:"width"`
}

type PngLayerResponse struct {
	Height    *string `pulumi:"height"`
	Label     *string `pulumi:"label"`
	OdataType string  `pulumi:"odataType"`
	Width     *string `pulumi:"width"`
}

type Rectangle struct {
	Height *string `pulumi:"height"`
	Left   *string `pulumi:"left"`
	Top    *string `pulumi:"top"`
	Width  *string `pulumi:"width"`
}

type RectangleResponse struct {
	Height *string `pulumi:"height"`
	Left   *string `pulumi:"left"`
	Top    *string `pulumi:"top"`
	Width  *string `pulumi:"width"`
}

type StandardEncoderPreset struct {
	Codecs    []interface{} `pulumi:"codecs"`
	Filters   *Filters      `pulumi:"filters"`
	Formats   []interface{} `pulumi:"formats"`
	OdataType string        `pulumi:"odataType"`
}

type StandardEncoderPresetResponse struct {
	Codecs    []interface{}    `pulumi:"codecs"`
	Filters   *FiltersResponse `pulumi:"filters"`
	Formats   []interface{}    `pulumi:"formats"`
	OdataType string           `pulumi:"odataType"`
}

type StorageAccount struct {
	Id   *string `pulumi:"id"`
	Type string  `pulumi:"type"`
}





type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(context.Context) StorageAccountOutput
}

type StorageAccountArgs struct {
	Id   pulumi.StringPtrInput `pulumi:"id"`
	Type pulumi.StringInput    `pulumi:"type"`
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (i StorageAccountArgs) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}





type StorageAccountArrayInput interface {
	pulumi.Input

	ToStorageAccountArrayOutput() StorageAccountArrayOutput
	ToStorageAccountArrayOutputWithContext(context.Context) StorageAccountArrayOutput
}

type StorageAccountArray []StorageAccountInput

func (StorageAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccount)(nil)).Elem()
}

func (i StorageAccountArray) ToStorageAccountArrayOutput() StorageAccountArrayOutput {
	return i.ToStorageAccountArrayOutputWithContext(context.Background())
}

func (i StorageAccountArray) ToStorageAccountArrayOutputWithContext(ctx context.Context) StorageAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountArrayOutput)
}

type StorageAccountOutput struct{ *pulumi.OutputState }

func (StorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (o StorageAccountOutput) ToStorageAccountOutput() StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccount) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccount) string { return v.Type }).(pulumi.StringOutput)
}

type StorageAccountArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccount)(nil)).Elem()
}

func (o StorageAccountArrayOutput) ToStorageAccountArrayOutput() StorageAccountArrayOutput {
	return o
}

func (o StorageAccountArrayOutput) ToStorageAccountArrayOutputWithContext(ctx context.Context) StorageAccountArrayOutput {
	return o
}

func (o StorageAccountArrayOutput) Index(i pulumi.IntInput) StorageAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccount {
		return vs[0].([]StorageAccount)[vs[1].(int)]
	}).(StorageAccountOutput)
}

type StorageAccountResponse struct {
	Id   *string `pulumi:"id"`
	Type string  `pulumi:"type"`
}

type StorageAccountResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutput() StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutputWithContext(ctx context.Context) StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o StorageAccountResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountResponse) string { return v.Type }).(pulumi.StringOutput)
}

type StorageAccountResponseArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseArrayOutput) ToStorageAccountResponseArrayOutput() StorageAccountResponseArrayOutput {
	return o
}

func (o StorageAccountResponseArrayOutput) ToStorageAccountResponseArrayOutputWithContext(ctx context.Context) StorageAccountResponseArrayOutput {
	return o
}

func (o StorageAccountResponseArrayOutput) Index(i pulumi.IntInput) StorageAccountResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccountResponse {
		return vs[0].([]StorageAccountResponse)[vs[1].(int)]
	}).(StorageAccountResponseOutput)
}

type StreamingEndpointAccessControl struct {
	Akamai *AkamaiAccessControl `pulumi:"akamai"`
	Ip     *IPAccessControl     `pulumi:"ip"`
}





type StreamingEndpointAccessControlInput interface {
	pulumi.Input

	ToStreamingEndpointAccessControlOutput() StreamingEndpointAccessControlOutput
	ToStreamingEndpointAccessControlOutputWithContext(context.Context) StreamingEndpointAccessControlOutput
}

type StreamingEndpointAccessControlArgs struct {
	Akamai AkamaiAccessControlPtrInput `pulumi:"akamai"`
	Ip     IPAccessControlPtrInput     `pulumi:"ip"`
}

func (StreamingEndpointAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingEndpointAccessControl)(nil)).Elem()
}

func (i StreamingEndpointAccessControlArgs) ToStreamingEndpointAccessControlOutput() StreamingEndpointAccessControlOutput {
	return i.ToStreamingEndpointAccessControlOutputWithContext(context.Background())
}

func (i StreamingEndpointAccessControlArgs) ToStreamingEndpointAccessControlOutputWithContext(ctx context.Context) StreamingEndpointAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointAccessControlOutput)
}

func (i StreamingEndpointAccessControlArgs) ToStreamingEndpointAccessControlPtrOutput() StreamingEndpointAccessControlPtrOutput {
	return i.ToStreamingEndpointAccessControlPtrOutputWithContext(context.Background())
}

func (i StreamingEndpointAccessControlArgs) ToStreamingEndpointAccessControlPtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointAccessControlOutput).ToStreamingEndpointAccessControlPtrOutputWithContext(ctx)
}









type StreamingEndpointAccessControlPtrInput interface {
	pulumi.Input

	ToStreamingEndpointAccessControlPtrOutput() StreamingEndpointAccessControlPtrOutput
	ToStreamingEndpointAccessControlPtrOutputWithContext(context.Context) StreamingEndpointAccessControlPtrOutput
}

type streamingEndpointAccessControlPtrType StreamingEndpointAccessControlArgs

func StreamingEndpointAccessControlPtr(v *StreamingEndpointAccessControlArgs) StreamingEndpointAccessControlPtrInput {
	return (*streamingEndpointAccessControlPtrType)(v)
}

func (*streamingEndpointAccessControlPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingEndpointAccessControl)(nil)).Elem()
}

func (i *streamingEndpointAccessControlPtrType) ToStreamingEndpointAccessControlPtrOutput() StreamingEndpointAccessControlPtrOutput {
	return i.ToStreamingEndpointAccessControlPtrOutputWithContext(context.Background())
}

func (i *streamingEndpointAccessControlPtrType) ToStreamingEndpointAccessControlPtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingEndpointAccessControlPtrOutput)
}

type StreamingEndpointAccessControlOutput struct{ *pulumi.OutputState }

func (StreamingEndpointAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingEndpointAccessControl)(nil)).Elem()
}

func (o StreamingEndpointAccessControlOutput) ToStreamingEndpointAccessControlOutput() StreamingEndpointAccessControlOutput {
	return o
}

func (o StreamingEndpointAccessControlOutput) ToStreamingEndpointAccessControlOutputWithContext(ctx context.Context) StreamingEndpointAccessControlOutput {
	return o
}

func (o StreamingEndpointAccessControlOutput) ToStreamingEndpointAccessControlPtrOutput() StreamingEndpointAccessControlPtrOutput {
	return o.ToStreamingEndpointAccessControlPtrOutputWithContext(context.Background())
}

func (o StreamingEndpointAccessControlOutput) ToStreamingEndpointAccessControlPtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingEndpointAccessControl) *StreamingEndpointAccessControl {
		return &v
	}).(StreamingEndpointAccessControlPtrOutput)
}

func (o StreamingEndpointAccessControlOutput) Akamai() AkamaiAccessControlPtrOutput {
	return o.ApplyT(func(v StreamingEndpointAccessControl) *AkamaiAccessControl { return v.Akamai }).(AkamaiAccessControlPtrOutput)
}

func (o StreamingEndpointAccessControlOutput) Ip() IPAccessControlPtrOutput {
	return o.ApplyT(func(v StreamingEndpointAccessControl) *IPAccessControl { return v.Ip }).(IPAccessControlPtrOutput)
}

type StreamingEndpointAccessControlPtrOutput struct{ *pulumi.OutputState }

func (StreamingEndpointAccessControlPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingEndpointAccessControl)(nil)).Elem()
}

func (o StreamingEndpointAccessControlPtrOutput) ToStreamingEndpointAccessControlPtrOutput() StreamingEndpointAccessControlPtrOutput {
	return o
}

func (o StreamingEndpointAccessControlPtrOutput) ToStreamingEndpointAccessControlPtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlPtrOutput {
	return o
}

func (o StreamingEndpointAccessControlPtrOutput) Elem() StreamingEndpointAccessControlOutput {
	return o.ApplyT(func(v *StreamingEndpointAccessControl) StreamingEndpointAccessControl {
		if v != nil {
			return *v
		}
		var ret StreamingEndpointAccessControl
		return ret
	}).(StreamingEndpointAccessControlOutput)
}

func (o StreamingEndpointAccessControlPtrOutput) Akamai() AkamaiAccessControlPtrOutput {
	return o.ApplyT(func(v *StreamingEndpointAccessControl) *AkamaiAccessControl {
		if v == nil {
			return nil
		}
		return v.Akamai
	}).(AkamaiAccessControlPtrOutput)
}

func (o StreamingEndpointAccessControlPtrOutput) Ip() IPAccessControlPtrOutput {
	return o.ApplyT(func(v *StreamingEndpointAccessControl) *IPAccessControl {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(IPAccessControlPtrOutput)
}

type StreamingEndpointAccessControlResponse struct {
	Akamai *AkamaiAccessControlResponse `pulumi:"akamai"`
	Ip     *IPAccessControlResponse     `pulumi:"ip"`
}

type StreamingEndpointAccessControlResponseOutput struct{ *pulumi.OutputState }

func (StreamingEndpointAccessControlResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingEndpointAccessControlResponse)(nil)).Elem()
}

func (o StreamingEndpointAccessControlResponseOutput) ToStreamingEndpointAccessControlResponseOutput() StreamingEndpointAccessControlResponseOutput {
	return o
}

func (o StreamingEndpointAccessControlResponseOutput) ToStreamingEndpointAccessControlResponseOutputWithContext(ctx context.Context) StreamingEndpointAccessControlResponseOutput {
	return o
}

func (o StreamingEndpointAccessControlResponseOutput) Akamai() AkamaiAccessControlResponsePtrOutput {
	return o.ApplyT(func(v StreamingEndpointAccessControlResponse) *AkamaiAccessControlResponse { return v.Akamai }).(AkamaiAccessControlResponsePtrOutput)
}

func (o StreamingEndpointAccessControlResponseOutput) Ip() IPAccessControlResponsePtrOutput {
	return o.ApplyT(func(v StreamingEndpointAccessControlResponse) *IPAccessControlResponse { return v.Ip }).(IPAccessControlResponsePtrOutput)
}

type StreamingEndpointAccessControlResponsePtrOutput struct{ *pulumi.OutputState }

func (StreamingEndpointAccessControlResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingEndpointAccessControlResponse)(nil)).Elem()
}

func (o StreamingEndpointAccessControlResponsePtrOutput) ToStreamingEndpointAccessControlResponsePtrOutput() StreamingEndpointAccessControlResponsePtrOutput {
	return o
}

func (o StreamingEndpointAccessControlResponsePtrOutput) ToStreamingEndpointAccessControlResponsePtrOutputWithContext(ctx context.Context) StreamingEndpointAccessControlResponsePtrOutput {
	return o
}

func (o StreamingEndpointAccessControlResponsePtrOutput) Elem() StreamingEndpointAccessControlResponseOutput {
	return o.ApplyT(func(v *StreamingEndpointAccessControlResponse) StreamingEndpointAccessControlResponse {
		if v != nil {
			return *v
		}
		var ret StreamingEndpointAccessControlResponse
		return ret
	}).(StreamingEndpointAccessControlResponseOutput)
}

func (o StreamingEndpointAccessControlResponsePtrOutput) Akamai() AkamaiAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *StreamingEndpointAccessControlResponse) *AkamaiAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.Akamai
	}).(AkamaiAccessControlResponsePtrOutput)
}

func (o StreamingEndpointAccessControlResponsePtrOutput) Ip() IPAccessControlResponsePtrOutput {
	return o.ApplyT(func(v *StreamingEndpointAccessControlResponse) *IPAccessControlResponse {
		if v == nil {
			return nil
		}
		return v.Ip
	}).(IPAccessControlResponsePtrOutput)
}

type StreamingLocatorContentKey struct {
	Id     string           `pulumi:"id"`
	Label  *string          `pulumi:"label"`
	Tracks []TrackSelection `pulumi:"tracks"`
	Value  *string          `pulumi:"value"`
}





type StreamingLocatorContentKeyInput interface {
	pulumi.Input

	ToStreamingLocatorContentKeyOutput() StreamingLocatorContentKeyOutput
	ToStreamingLocatorContentKeyOutputWithContext(context.Context) StreamingLocatorContentKeyOutput
}

type StreamingLocatorContentKeyArgs struct {
	Id     pulumi.StringInput       `pulumi:"id"`
	Label  pulumi.StringPtrInput    `pulumi:"label"`
	Tracks TrackSelectionArrayInput `pulumi:"tracks"`
	Value  pulumi.StringPtrInput    `pulumi:"value"`
}

func (StreamingLocatorContentKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingLocatorContentKey)(nil)).Elem()
}

func (i StreamingLocatorContentKeyArgs) ToStreamingLocatorContentKeyOutput() StreamingLocatorContentKeyOutput {
	return i.ToStreamingLocatorContentKeyOutputWithContext(context.Background())
}

func (i StreamingLocatorContentKeyArgs) ToStreamingLocatorContentKeyOutputWithContext(ctx context.Context) StreamingLocatorContentKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingLocatorContentKeyOutput)
}





type StreamingLocatorContentKeyArrayInput interface {
	pulumi.Input

	ToStreamingLocatorContentKeyArrayOutput() StreamingLocatorContentKeyArrayOutput
	ToStreamingLocatorContentKeyArrayOutputWithContext(context.Context) StreamingLocatorContentKeyArrayOutput
}

type StreamingLocatorContentKeyArray []StreamingLocatorContentKeyInput

func (StreamingLocatorContentKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamingLocatorContentKey)(nil)).Elem()
}

func (i StreamingLocatorContentKeyArray) ToStreamingLocatorContentKeyArrayOutput() StreamingLocatorContentKeyArrayOutput {
	return i.ToStreamingLocatorContentKeyArrayOutputWithContext(context.Background())
}

func (i StreamingLocatorContentKeyArray) ToStreamingLocatorContentKeyArrayOutputWithContext(ctx context.Context) StreamingLocatorContentKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingLocatorContentKeyArrayOutput)
}

type StreamingLocatorContentKeyOutput struct{ *pulumi.OutputState }

func (StreamingLocatorContentKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingLocatorContentKey)(nil)).Elem()
}

func (o StreamingLocatorContentKeyOutput) ToStreamingLocatorContentKeyOutput() StreamingLocatorContentKeyOutput {
	return o
}

func (o StreamingLocatorContentKeyOutput) ToStreamingLocatorContentKeyOutputWithContext(ctx context.Context) StreamingLocatorContentKeyOutput {
	return o
}

func (o StreamingLocatorContentKeyOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StreamingLocatorContentKey) string { return v.Id }).(pulumi.StringOutput)
}

func (o StreamingLocatorContentKeyOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingLocatorContentKey) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o StreamingLocatorContentKeyOutput) Tracks() TrackSelectionArrayOutput {
	return o.ApplyT(func(v StreamingLocatorContentKey) []TrackSelection { return v.Tracks }).(TrackSelectionArrayOutput)
}

func (o StreamingLocatorContentKeyOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingLocatorContentKey) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type StreamingLocatorContentKeyArrayOutput struct{ *pulumi.OutputState }

func (StreamingLocatorContentKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamingLocatorContentKey)(nil)).Elem()
}

func (o StreamingLocatorContentKeyArrayOutput) ToStreamingLocatorContentKeyArrayOutput() StreamingLocatorContentKeyArrayOutput {
	return o
}

func (o StreamingLocatorContentKeyArrayOutput) ToStreamingLocatorContentKeyArrayOutputWithContext(ctx context.Context) StreamingLocatorContentKeyArrayOutput {
	return o
}

func (o StreamingLocatorContentKeyArrayOutput) Index(i pulumi.IntInput) StreamingLocatorContentKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StreamingLocatorContentKey {
		return vs[0].([]StreamingLocatorContentKey)[vs[1].(int)]
	}).(StreamingLocatorContentKeyOutput)
}

type StreamingLocatorContentKeyResponse struct {
	Id         string                   `pulumi:"id"`
	Label      *string                  `pulumi:"label"`
	PolicyName string                   `pulumi:"policyName"`
	Tracks     []TrackSelectionResponse `pulumi:"tracks"`
	Type       string                   `pulumi:"type"`
	Value      *string                  `pulumi:"value"`
}

type StreamingLocatorContentKeyResponseOutput struct{ *pulumi.OutputState }

func (StreamingLocatorContentKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingLocatorContentKeyResponse)(nil)).Elem()
}

func (o StreamingLocatorContentKeyResponseOutput) ToStreamingLocatorContentKeyResponseOutput() StreamingLocatorContentKeyResponseOutput {
	return o
}

func (o StreamingLocatorContentKeyResponseOutput) ToStreamingLocatorContentKeyResponseOutputWithContext(ctx context.Context) StreamingLocatorContentKeyResponseOutput {
	return o
}

func (o StreamingLocatorContentKeyResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StreamingLocatorContentKeyResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o StreamingLocatorContentKeyResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingLocatorContentKeyResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o StreamingLocatorContentKeyResponseOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v StreamingLocatorContentKeyResponse) string { return v.PolicyName }).(pulumi.StringOutput)
}

func (o StreamingLocatorContentKeyResponseOutput) Tracks() TrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v StreamingLocatorContentKeyResponse) []TrackSelectionResponse { return v.Tracks }).(TrackSelectionResponseArrayOutput)
}

func (o StreamingLocatorContentKeyResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StreamingLocatorContentKeyResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o StreamingLocatorContentKeyResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingLocatorContentKeyResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type StreamingLocatorContentKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (StreamingLocatorContentKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamingLocatorContentKeyResponse)(nil)).Elem()
}

func (o StreamingLocatorContentKeyResponseArrayOutput) ToStreamingLocatorContentKeyResponseArrayOutput() StreamingLocatorContentKeyResponseArrayOutput {
	return o
}

func (o StreamingLocatorContentKeyResponseArrayOutput) ToStreamingLocatorContentKeyResponseArrayOutputWithContext(ctx context.Context) StreamingLocatorContentKeyResponseArrayOutput {
	return o
}

func (o StreamingLocatorContentKeyResponseArrayOutput) Index(i pulumi.IntInput) StreamingLocatorContentKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StreamingLocatorContentKeyResponse {
		return vs[0].([]StreamingLocatorContentKeyResponse)[vs[1].(int)]
	}).(StreamingLocatorContentKeyResponseOutput)
}

type StreamingPathResponse struct {
	EncryptionScheme  string   `pulumi:"encryptionScheme"`
	Paths             []string `pulumi:"paths"`
	StreamingProtocol string   `pulumi:"streamingProtocol"`
}

type StreamingPathResponseOutput struct{ *pulumi.OutputState }

func (StreamingPathResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPathResponse)(nil)).Elem()
}

func (o StreamingPathResponseOutput) ToStreamingPathResponseOutput() StreamingPathResponseOutput {
	return o
}

func (o StreamingPathResponseOutput) ToStreamingPathResponseOutputWithContext(ctx context.Context) StreamingPathResponseOutput {
	return o
}

func (o StreamingPathResponseOutput) EncryptionScheme() pulumi.StringOutput {
	return o.ApplyT(func(v StreamingPathResponse) string { return v.EncryptionScheme }).(pulumi.StringOutput)
}

func (o StreamingPathResponseOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v StreamingPathResponse) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

func (o StreamingPathResponseOutput) StreamingProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v StreamingPathResponse) string { return v.StreamingProtocol }).(pulumi.StringOutput)
}

type StreamingPathResponseArrayOutput struct{ *pulumi.OutputState }

func (StreamingPathResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamingPathResponse)(nil)).Elem()
}

func (o StreamingPathResponseArrayOutput) ToStreamingPathResponseArrayOutput() StreamingPathResponseArrayOutput {
	return o
}

func (o StreamingPathResponseArrayOutput) ToStreamingPathResponseArrayOutputWithContext(ctx context.Context) StreamingPathResponseArrayOutput {
	return o
}

func (o StreamingPathResponseArrayOutput) Index(i pulumi.IntInput) StreamingPathResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StreamingPathResponse {
		return vs[0].([]StreamingPathResponse)[vs[1].(int)]
	}).(StreamingPathResponseOutput)
}

type StreamingPolicyContentKey struct {
	Label      *string          `pulumi:"label"`
	PolicyName *string          `pulumi:"policyName"`
	Tracks     []TrackSelection `pulumi:"tracks"`
}





type StreamingPolicyContentKeyInput interface {
	pulumi.Input

	ToStreamingPolicyContentKeyOutput() StreamingPolicyContentKeyOutput
	ToStreamingPolicyContentKeyOutputWithContext(context.Context) StreamingPolicyContentKeyOutput
}

type StreamingPolicyContentKeyArgs struct {
	Label      pulumi.StringPtrInput    `pulumi:"label"`
	PolicyName pulumi.StringPtrInput    `pulumi:"policyName"`
	Tracks     TrackSelectionArrayInput `pulumi:"tracks"`
}

func (StreamingPolicyContentKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyContentKey)(nil)).Elem()
}

func (i StreamingPolicyContentKeyArgs) ToStreamingPolicyContentKeyOutput() StreamingPolicyContentKeyOutput {
	return i.ToStreamingPolicyContentKeyOutputWithContext(context.Background())
}

func (i StreamingPolicyContentKeyArgs) ToStreamingPolicyContentKeyOutputWithContext(ctx context.Context) StreamingPolicyContentKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyContentKeyOutput)
}





type StreamingPolicyContentKeyArrayInput interface {
	pulumi.Input

	ToStreamingPolicyContentKeyArrayOutput() StreamingPolicyContentKeyArrayOutput
	ToStreamingPolicyContentKeyArrayOutputWithContext(context.Context) StreamingPolicyContentKeyArrayOutput
}

type StreamingPolicyContentKeyArray []StreamingPolicyContentKeyInput

func (StreamingPolicyContentKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamingPolicyContentKey)(nil)).Elem()
}

func (i StreamingPolicyContentKeyArray) ToStreamingPolicyContentKeyArrayOutput() StreamingPolicyContentKeyArrayOutput {
	return i.ToStreamingPolicyContentKeyArrayOutputWithContext(context.Background())
}

func (i StreamingPolicyContentKeyArray) ToStreamingPolicyContentKeyArrayOutputWithContext(ctx context.Context) StreamingPolicyContentKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyContentKeyArrayOutput)
}

type StreamingPolicyContentKeyOutput struct{ *pulumi.OutputState }

func (StreamingPolicyContentKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyContentKey)(nil)).Elem()
}

func (o StreamingPolicyContentKeyOutput) ToStreamingPolicyContentKeyOutput() StreamingPolicyContentKeyOutput {
	return o
}

func (o StreamingPolicyContentKeyOutput) ToStreamingPolicyContentKeyOutputWithContext(ctx context.Context) StreamingPolicyContentKeyOutput {
	return o
}

func (o StreamingPolicyContentKeyOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyContentKey) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o StreamingPolicyContentKeyOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyContentKey) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

func (o StreamingPolicyContentKeyOutput) Tracks() TrackSelectionArrayOutput {
	return o.ApplyT(func(v StreamingPolicyContentKey) []TrackSelection { return v.Tracks }).(TrackSelectionArrayOutput)
}

type StreamingPolicyContentKeyArrayOutput struct{ *pulumi.OutputState }

func (StreamingPolicyContentKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamingPolicyContentKey)(nil)).Elem()
}

func (o StreamingPolicyContentKeyArrayOutput) ToStreamingPolicyContentKeyArrayOutput() StreamingPolicyContentKeyArrayOutput {
	return o
}

func (o StreamingPolicyContentKeyArrayOutput) ToStreamingPolicyContentKeyArrayOutputWithContext(ctx context.Context) StreamingPolicyContentKeyArrayOutput {
	return o
}

func (o StreamingPolicyContentKeyArrayOutput) Index(i pulumi.IntInput) StreamingPolicyContentKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StreamingPolicyContentKey {
		return vs[0].([]StreamingPolicyContentKey)[vs[1].(int)]
	}).(StreamingPolicyContentKeyOutput)
}

type StreamingPolicyContentKeyResponse struct {
	Label      *string                  `pulumi:"label"`
	PolicyName *string                  `pulumi:"policyName"`
	Tracks     []TrackSelectionResponse `pulumi:"tracks"`
}

type StreamingPolicyContentKeyResponseOutput struct{ *pulumi.OutputState }

func (StreamingPolicyContentKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyContentKeyResponse)(nil)).Elem()
}

func (o StreamingPolicyContentKeyResponseOutput) ToStreamingPolicyContentKeyResponseOutput() StreamingPolicyContentKeyResponseOutput {
	return o
}

func (o StreamingPolicyContentKeyResponseOutput) ToStreamingPolicyContentKeyResponseOutputWithContext(ctx context.Context) StreamingPolicyContentKeyResponseOutput {
	return o
}

func (o StreamingPolicyContentKeyResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyContentKeyResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o StreamingPolicyContentKeyResponseOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyContentKeyResponse) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

func (o StreamingPolicyContentKeyResponseOutput) Tracks() TrackSelectionResponseArrayOutput {
	return o.ApplyT(func(v StreamingPolicyContentKeyResponse) []TrackSelectionResponse { return v.Tracks }).(TrackSelectionResponseArrayOutput)
}

type StreamingPolicyContentKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (StreamingPolicyContentKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StreamingPolicyContentKeyResponse)(nil)).Elem()
}

func (o StreamingPolicyContentKeyResponseArrayOutput) ToStreamingPolicyContentKeyResponseArrayOutput() StreamingPolicyContentKeyResponseArrayOutput {
	return o
}

func (o StreamingPolicyContentKeyResponseArrayOutput) ToStreamingPolicyContentKeyResponseArrayOutputWithContext(ctx context.Context) StreamingPolicyContentKeyResponseArrayOutput {
	return o
}

func (o StreamingPolicyContentKeyResponseArrayOutput) Index(i pulumi.IntInput) StreamingPolicyContentKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StreamingPolicyContentKeyResponse {
		return vs[0].([]StreamingPolicyContentKeyResponse)[vs[1].(int)]
	}).(StreamingPolicyContentKeyResponseOutput)
}

type StreamingPolicyContentKeys struct {
	DefaultKey         *DefaultKey                 `pulumi:"defaultKey"`
	KeyToTrackMappings []StreamingPolicyContentKey `pulumi:"keyToTrackMappings"`
}





type StreamingPolicyContentKeysInput interface {
	pulumi.Input

	ToStreamingPolicyContentKeysOutput() StreamingPolicyContentKeysOutput
	ToStreamingPolicyContentKeysOutputWithContext(context.Context) StreamingPolicyContentKeysOutput
}

type StreamingPolicyContentKeysArgs struct {
	DefaultKey         DefaultKeyPtrInput                  `pulumi:"defaultKey"`
	KeyToTrackMappings StreamingPolicyContentKeyArrayInput `pulumi:"keyToTrackMappings"`
}

func (StreamingPolicyContentKeysArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyContentKeys)(nil)).Elem()
}

func (i StreamingPolicyContentKeysArgs) ToStreamingPolicyContentKeysOutput() StreamingPolicyContentKeysOutput {
	return i.ToStreamingPolicyContentKeysOutputWithContext(context.Background())
}

func (i StreamingPolicyContentKeysArgs) ToStreamingPolicyContentKeysOutputWithContext(ctx context.Context) StreamingPolicyContentKeysOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyContentKeysOutput)
}

func (i StreamingPolicyContentKeysArgs) ToStreamingPolicyContentKeysPtrOutput() StreamingPolicyContentKeysPtrOutput {
	return i.ToStreamingPolicyContentKeysPtrOutputWithContext(context.Background())
}

func (i StreamingPolicyContentKeysArgs) ToStreamingPolicyContentKeysPtrOutputWithContext(ctx context.Context) StreamingPolicyContentKeysPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyContentKeysOutput).ToStreamingPolicyContentKeysPtrOutputWithContext(ctx)
}









type StreamingPolicyContentKeysPtrInput interface {
	pulumi.Input

	ToStreamingPolicyContentKeysPtrOutput() StreamingPolicyContentKeysPtrOutput
	ToStreamingPolicyContentKeysPtrOutputWithContext(context.Context) StreamingPolicyContentKeysPtrOutput
}

type streamingPolicyContentKeysPtrType StreamingPolicyContentKeysArgs

func StreamingPolicyContentKeysPtr(v *StreamingPolicyContentKeysArgs) StreamingPolicyContentKeysPtrInput {
	return (*streamingPolicyContentKeysPtrType)(v)
}

func (*streamingPolicyContentKeysPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyContentKeys)(nil)).Elem()
}

func (i *streamingPolicyContentKeysPtrType) ToStreamingPolicyContentKeysPtrOutput() StreamingPolicyContentKeysPtrOutput {
	return i.ToStreamingPolicyContentKeysPtrOutputWithContext(context.Background())
}

func (i *streamingPolicyContentKeysPtrType) ToStreamingPolicyContentKeysPtrOutputWithContext(ctx context.Context) StreamingPolicyContentKeysPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyContentKeysPtrOutput)
}

type StreamingPolicyContentKeysOutput struct{ *pulumi.OutputState }

func (StreamingPolicyContentKeysOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyContentKeys)(nil)).Elem()
}

func (o StreamingPolicyContentKeysOutput) ToStreamingPolicyContentKeysOutput() StreamingPolicyContentKeysOutput {
	return o
}

func (o StreamingPolicyContentKeysOutput) ToStreamingPolicyContentKeysOutputWithContext(ctx context.Context) StreamingPolicyContentKeysOutput {
	return o
}

func (o StreamingPolicyContentKeysOutput) ToStreamingPolicyContentKeysPtrOutput() StreamingPolicyContentKeysPtrOutput {
	return o.ToStreamingPolicyContentKeysPtrOutputWithContext(context.Background())
}

func (o StreamingPolicyContentKeysOutput) ToStreamingPolicyContentKeysPtrOutputWithContext(ctx context.Context) StreamingPolicyContentKeysPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingPolicyContentKeys) *StreamingPolicyContentKeys {
		return &v
	}).(StreamingPolicyContentKeysPtrOutput)
}

func (o StreamingPolicyContentKeysOutput) DefaultKey() DefaultKeyPtrOutput {
	return o.ApplyT(func(v StreamingPolicyContentKeys) *DefaultKey { return v.DefaultKey }).(DefaultKeyPtrOutput)
}

func (o StreamingPolicyContentKeysOutput) KeyToTrackMappings() StreamingPolicyContentKeyArrayOutput {
	return o.ApplyT(func(v StreamingPolicyContentKeys) []StreamingPolicyContentKey { return v.KeyToTrackMappings }).(StreamingPolicyContentKeyArrayOutput)
}

type StreamingPolicyContentKeysPtrOutput struct{ *pulumi.OutputState }

func (StreamingPolicyContentKeysPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyContentKeys)(nil)).Elem()
}

func (o StreamingPolicyContentKeysPtrOutput) ToStreamingPolicyContentKeysPtrOutput() StreamingPolicyContentKeysPtrOutput {
	return o
}

func (o StreamingPolicyContentKeysPtrOutput) ToStreamingPolicyContentKeysPtrOutputWithContext(ctx context.Context) StreamingPolicyContentKeysPtrOutput {
	return o
}

func (o StreamingPolicyContentKeysPtrOutput) Elem() StreamingPolicyContentKeysOutput {
	return o.ApplyT(func(v *StreamingPolicyContentKeys) StreamingPolicyContentKeys {
		if v != nil {
			return *v
		}
		var ret StreamingPolicyContentKeys
		return ret
	}).(StreamingPolicyContentKeysOutput)
}

func (o StreamingPolicyContentKeysPtrOutput) DefaultKey() DefaultKeyPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyContentKeys) *DefaultKey {
		if v == nil {
			return nil
		}
		return v.DefaultKey
	}).(DefaultKeyPtrOutput)
}

func (o StreamingPolicyContentKeysPtrOutput) KeyToTrackMappings() StreamingPolicyContentKeyArrayOutput {
	return o.ApplyT(func(v *StreamingPolicyContentKeys) []StreamingPolicyContentKey {
		if v == nil {
			return nil
		}
		return v.KeyToTrackMappings
	}).(StreamingPolicyContentKeyArrayOutput)
}

type StreamingPolicyContentKeysResponse struct {
	DefaultKey         *DefaultKeyResponse                 `pulumi:"defaultKey"`
	KeyToTrackMappings []StreamingPolicyContentKeyResponse `pulumi:"keyToTrackMappings"`
}

type StreamingPolicyContentKeysResponseOutput struct{ *pulumi.OutputState }

func (StreamingPolicyContentKeysResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyContentKeysResponse)(nil)).Elem()
}

func (o StreamingPolicyContentKeysResponseOutput) ToStreamingPolicyContentKeysResponseOutput() StreamingPolicyContentKeysResponseOutput {
	return o
}

func (o StreamingPolicyContentKeysResponseOutput) ToStreamingPolicyContentKeysResponseOutputWithContext(ctx context.Context) StreamingPolicyContentKeysResponseOutput {
	return o
}

func (o StreamingPolicyContentKeysResponseOutput) DefaultKey() DefaultKeyResponsePtrOutput {
	return o.ApplyT(func(v StreamingPolicyContentKeysResponse) *DefaultKeyResponse { return v.DefaultKey }).(DefaultKeyResponsePtrOutput)
}

func (o StreamingPolicyContentKeysResponseOutput) KeyToTrackMappings() StreamingPolicyContentKeyResponseArrayOutput {
	return o.ApplyT(func(v StreamingPolicyContentKeysResponse) []StreamingPolicyContentKeyResponse {
		return v.KeyToTrackMappings
	}).(StreamingPolicyContentKeyResponseArrayOutput)
}

type StreamingPolicyContentKeysResponsePtrOutput struct{ *pulumi.OutputState }

func (StreamingPolicyContentKeysResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyContentKeysResponse)(nil)).Elem()
}

func (o StreamingPolicyContentKeysResponsePtrOutput) ToStreamingPolicyContentKeysResponsePtrOutput() StreamingPolicyContentKeysResponsePtrOutput {
	return o
}

func (o StreamingPolicyContentKeysResponsePtrOutput) ToStreamingPolicyContentKeysResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyContentKeysResponsePtrOutput {
	return o
}

func (o StreamingPolicyContentKeysResponsePtrOutput) Elem() StreamingPolicyContentKeysResponseOutput {
	return o.ApplyT(func(v *StreamingPolicyContentKeysResponse) StreamingPolicyContentKeysResponse {
		if v != nil {
			return *v
		}
		var ret StreamingPolicyContentKeysResponse
		return ret
	}).(StreamingPolicyContentKeysResponseOutput)
}

func (o StreamingPolicyContentKeysResponsePtrOutput) DefaultKey() DefaultKeyResponsePtrOutput {
	return o.ApplyT(func(v *StreamingPolicyContentKeysResponse) *DefaultKeyResponse {
		if v == nil {
			return nil
		}
		return v.DefaultKey
	}).(DefaultKeyResponsePtrOutput)
}

func (o StreamingPolicyContentKeysResponsePtrOutput) KeyToTrackMappings() StreamingPolicyContentKeyResponseArrayOutput {
	return o.ApplyT(func(v *StreamingPolicyContentKeysResponse) []StreamingPolicyContentKeyResponse {
		if v == nil {
			return nil
		}
		return v.KeyToTrackMappings
	}).(StreamingPolicyContentKeyResponseArrayOutput)
}

type StreamingPolicyFairPlayConfiguration struct {
	AllowPersistentLicense              bool    `pulumi:"allowPersistentLicense"`
	CustomLicenseAcquisitionUrlTemplate *string `pulumi:"customLicenseAcquisitionUrlTemplate"`
}





type StreamingPolicyFairPlayConfigurationInput interface {
	pulumi.Input

	ToStreamingPolicyFairPlayConfigurationOutput() StreamingPolicyFairPlayConfigurationOutput
	ToStreamingPolicyFairPlayConfigurationOutputWithContext(context.Context) StreamingPolicyFairPlayConfigurationOutput
}

type StreamingPolicyFairPlayConfigurationArgs struct {
	AllowPersistentLicense              pulumi.BoolInput      `pulumi:"allowPersistentLicense"`
	CustomLicenseAcquisitionUrlTemplate pulumi.StringPtrInput `pulumi:"customLicenseAcquisitionUrlTemplate"`
}

func (StreamingPolicyFairPlayConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyFairPlayConfiguration)(nil)).Elem()
}

func (i StreamingPolicyFairPlayConfigurationArgs) ToStreamingPolicyFairPlayConfigurationOutput() StreamingPolicyFairPlayConfigurationOutput {
	return i.ToStreamingPolicyFairPlayConfigurationOutputWithContext(context.Background())
}

func (i StreamingPolicyFairPlayConfigurationArgs) ToStreamingPolicyFairPlayConfigurationOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyFairPlayConfigurationOutput)
}

func (i StreamingPolicyFairPlayConfigurationArgs) ToStreamingPolicyFairPlayConfigurationPtrOutput() StreamingPolicyFairPlayConfigurationPtrOutput {
	return i.ToStreamingPolicyFairPlayConfigurationPtrOutputWithContext(context.Background())
}

func (i StreamingPolicyFairPlayConfigurationArgs) ToStreamingPolicyFairPlayConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyFairPlayConfigurationOutput).ToStreamingPolicyFairPlayConfigurationPtrOutputWithContext(ctx)
}









type StreamingPolicyFairPlayConfigurationPtrInput interface {
	pulumi.Input

	ToStreamingPolicyFairPlayConfigurationPtrOutput() StreamingPolicyFairPlayConfigurationPtrOutput
	ToStreamingPolicyFairPlayConfigurationPtrOutputWithContext(context.Context) StreamingPolicyFairPlayConfigurationPtrOutput
}

type streamingPolicyFairPlayConfigurationPtrType StreamingPolicyFairPlayConfigurationArgs

func StreamingPolicyFairPlayConfigurationPtr(v *StreamingPolicyFairPlayConfigurationArgs) StreamingPolicyFairPlayConfigurationPtrInput {
	return (*streamingPolicyFairPlayConfigurationPtrType)(v)
}

func (*streamingPolicyFairPlayConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyFairPlayConfiguration)(nil)).Elem()
}

func (i *streamingPolicyFairPlayConfigurationPtrType) ToStreamingPolicyFairPlayConfigurationPtrOutput() StreamingPolicyFairPlayConfigurationPtrOutput {
	return i.ToStreamingPolicyFairPlayConfigurationPtrOutputWithContext(context.Background())
}

func (i *streamingPolicyFairPlayConfigurationPtrType) ToStreamingPolicyFairPlayConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyFairPlayConfigurationPtrOutput)
}

type StreamingPolicyFairPlayConfigurationOutput struct{ *pulumi.OutputState }

func (StreamingPolicyFairPlayConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyFairPlayConfiguration)(nil)).Elem()
}

func (o StreamingPolicyFairPlayConfigurationOutput) ToStreamingPolicyFairPlayConfigurationOutput() StreamingPolicyFairPlayConfigurationOutput {
	return o
}

func (o StreamingPolicyFairPlayConfigurationOutput) ToStreamingPolicyFairPlayConfigurationOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationOutput {
	return o
}

func (o StreamingPolicyFairPlayConfigurationOutput) ToStreamingPolicyFairPlayConfigurationPtrOutput() StreamingPolicyFairPlayConfigurationPtrOutput {
	return o.ToStreamingPolicyFairPlayConfigurationPtrOutputWithContext(context.Background())
}

func (o StreamingPolicyFairPlayConfigurationOutput) ToStreamingPolicyFairPlayConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingPolicyFairPlayConfiguration) *StreamingPolicyFairPlayConfiguration {
		return &v
	}).(StreamingPolicyFairPlayConfigurationPtrOutput)
}

func (o StreamingPolicyFairPlayConfigurationOutput) AllowPersistentLicense() pulumi.BoolOutput {
	return o.ApplyT(func(v StreamingPolicyFairPlayConfiguration) bool { return v.AllowPersistentLicense }).(pulumi.BoolOutput)
}

func (o StreamingPolicyFairPlayConfigurationOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyFairPlayConfiguration) *string { return v.CustomLicenseAcquisitionUrlTemplate }).(pulumi.StringPtrOutput)
}

type StreamingPolicyFairPlayConfigurationPtrOutput struct{ *pulumi.OutputState }

func (StreamingPolicyFairPlayConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyFairPlayConfiguration)(nil)).Elem()
}

func (o StreamingPolicyFairPlayConfigurationPtrOutput) ToStreamingPolicyFairPlayConfigurationPtrOutput() StreamingPolicyFairPlayConfigurationPtrOutput {
	return o
}

func (o StreamingPolicyFairPlayConfigurationPtrOutput) ToStreamingPolicyFairPlayConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationPtrOutput {
	return o
}

func (o StreamingPolicyFairPlayConfigurationPtrOutput) Elem() StreamingPolicyFairPlayConfigurationOutput {
	return o.ApplyT(func(v *StreamingPolicyFairPlayConfiguration) StreamingPolicyFairPlayConfiguration {
		if v != nil {
			return *v
		}
		var ret StreamingPolicyFairPlayConfiguration
		return ret
	}).(StreamingPolicyFairPlayConfigurationOutput)
}

func (o StreamingPolicyFairPlayConfigurationPtrOutput) AllowPersistentLicense() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyFairPlayConfiguration) *bool {
		if v == nil {
			return nil
		}
		return &v.AllowPersistentLicense
	}).(pulumi.BoolPtrOutput)
}

func (o StreamingPolicyFairPlayConfigurationPtrOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyFairPlayConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.CustomLicenseAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

type StreamingPolicyFairPlayConfigurationResponse struct {
	AllowPersistentLicense              bool    `pulumi:"allowPersistentLicense"`
	CustomLicenseAcquisitionUrlTemplate *string `pulumi:"customLicenseAcquisitionUrlTemplate"`
}

type StreamingPolicyFairPlayConfigurationResponseOutput struct{ *pulumi.OutputState }

func (StreamingPolicyFairPlayConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyFairPlayConfigurationResponse)(nil)).Elem()
}

func (o StreamingPolicyFairPlayConfigurationResponseOutput) ToStreamingPolicyFairPlayConfigurationResponseOutput() StreamingPolicyFairPlayConfigurationResponseOutput {
	return o
}

func (o StreamingPolicyFairPlayConfigurationResponseOutput) ToStreamingPolicyFairPlayConfigurationResponseOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationResponseOutput {
	return o
}

func (o StreamingPolicyFairPlayConfigurationResponseOutput) AllowPersistentLicense() pulumi.BoolOutput {
	return o.ApplyT(func(v StreamingPolicyFairPlayConfigurationResponse) bool { return v.AllowPersistentLicense }).(pulumi.BoolOutput)
}

func (o StreamingPolicyFairPlayConfigurationResponseOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyFairPlayConfigurationResponse) *string {
		return v.CustomLicenseAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

type StreamingPolicyFairPlayConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (StreamingPolicyFairPlayConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyFairPlayConfigurationResponse)(nil)).Elem()
}

func (o StreamingPolicyFairPlayConfigurationResponsePtrOutput) ToStreamingPolicyFairPlayConfigurationResponsePtrOutput() StreamingPolicyFairPlayConfigurationResponsePtrOutput {
	return o
}

func (o StreamingPolicyFairPlayConfigurationResponsePtrOutput) ToStreamingPolicyFairPlayConfigurationResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyFairPlayConfigurationResponsePtrOutput {
	return o
}

func (o StreamingPolicyFairPlayConfigurationResponsePtrOutput) Elem() StreamingPolicyFairPlayConfigurationResponseOutput {
	return o.ApplyT(func(v *StreamingPolicyFairPlayConfigurationResponse) StreamingPolicyFairPlayConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret StreamingPolicyFairPlayConfigurationResponse
		return ret
	}).(StreamingPolicyFairPlayConfigurationResponseOutput)
}

func (o StreamingPolicyFairPlayConfigurationResponsePtrOutput) AllowPersistentLicense() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyFairPlayConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.AllowPersistentLicense
	}).(pulumi.BoolPtrOutput)
}

func (o StreamingPolicyFairPlayConfigurationResponsePtrOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyFairPlayConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomLicenseAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

type StreamingPolicyPlayReadyConfiguration struct {
	CustomLicenseAcquisitionUrlTemplate *string `pulumi:"customLicenseAcquisitionUrlTemplate"`
	PlayReadyCustomAttributes           *string `pulumi:"playReadyCustomAttributes"`
}





type StreamingPolicyPlayReadyConfigurationInput interface {
	pulumi.Input

	ToStreamingPolicyPlayReadyConfigurationOutput() StreamingPolicyPlayReadyConfigurationOutput
	ToStreamingPolicyPlayReadyConfigurationOutputWithContext(context.Context) StreamingPolicyPlayReadyConfigurationOutput
}

type StreamingPolicyPlayReadyConfigurationArgs struct {
	CustomLicenseAcquisitionUrlTemplate pulumi.StringPtrInput `pulumi:"customLicenseAcquisitionUrlTemplate"`
	PlayReadyCustomAttributes           pulumi.StringPtrInput `pulumi:"playReadyCustomAttributes"`
}

func (StreamingPolicyPlayReadyConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyPlayReadyConfiguration)(nil)).Elem()
}

func (i StreamingPolicyPlayReadyConfigurationArgs) ToStreamingPolicyPlayReadyConfigurationOutput() StreamingPolicyPlayReadyConfigurationOutput {
	return i.ToStreamingPolicyPlayReadyConfigurationOutputWithContext(context.Background())
}

func (i StreamingPolicyPlayReadyConfigurationArgs) ToStreamingPolicyPlayReadyConfigurationOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyPlayReadyConfigurationOutput)
}

func (i StreamingPolicyPlayReadyConfigurationArgs) ToStreamingPolicyPlayReadyConfigurationPtrOutput() StreamingPolicyPlayReadyConfigurationPtrOutput {
	return i.ToStreamingPolicyPlayReadyConfigurationPtrOutputWithContext(context.Background())
}

func (i StreamingPolicyPlayReadyConfigurationArgs) ToStreamingPolicyPlayReadyConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyPlayReadyConfigurationOutput).ToStreamingPolicyPlayReadyConfigurationPtrOutputWithContext(ctx)
}









type StreamingPolicyPlayReadyConfigurationPtrInput interface {
	pulumi.Input

	ToStreamingPolicyPlayReadyConfigurationPtrOutput() StreamingPolicyPlayReadyConfigurationPtrOutput
	ToStreamingPolicyPlayReadyConfigurationPtrOutputWithContext(context.Context) StreamingPolicyPlayReadyConfigurationPtrOutput
}

type streamingPolicyPlayReadyConfigurationPtrType StreamingPolicyPlayReadyConfigurationArgs

func StreamingPolicyPlayReadyConfigurationPtr(v *StreamingPolicyPlayReadyConfigurationArgs) StreamingPolicyPlayReadyConfigurationPtrInput {
	return (*streamingPolicyPlayReadyConfigurationPtrType)(v)
}

func (*streamingPolicyPlayReadyConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyPlayReadyConfiguration)(nil)).Elem()
}

func (i *streamingPolicyPlayReadyConfigurationPtrType) ToStreamingPolicyPlayReadyConfigurationPtrOutput() StreamingPolicyPlayReadyConfigurationPtrOutput {
	return i.ToStreamingPolicyPlayReadyConfigurationPtrOutputWithContext(context.Background())
}

func (i *streamingPolicyPlayReadyConfigurationPtrType) ToStreamingPolicyPlayReadyConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyPlayReadyConfigurationPtrOutput)
}

type StreamingPolicyPlayReadyConfigurationOutput struct{ *pulumi.OutputState }

func (StreamingPolicyPlayReadyConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyPlayReadyConfiguration)(nil)).Elem()
}

func (o StreamingPolicyPlayReadyConfigurationOutput) ToStreamingPolicyPlayReadyConfigurationOutput() StreamingPolicyPlayReadyConfigurationOutput {
	return o
}

func (o StreamingPolicyPlayReadyConfigurationOutput) ToStreamingPolicyPlayReadyConfigurationOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationOutput {
	return o
}

func (o StreamingPolicyPlayReadyConfigurationOutput) ToStreamingPolicyPlayReadyConfigurationPtrOutput() StreamingPolicyPlayReadyConfigurationPtrOutput {
	return o.ToStreamingPolicyPlayReadyConfigurationPtrOutputWithContext(context.Background())
}

func (o StreamingPolicyPlayReadyConfigurationOutput) ToStreamingPolicyPlayReadyConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingPolicyPlayReadyConfiguration) *StreamingPolicyPlayReadyConfiguration {
		return &v
	}).(StreamingPolicyPlayReadyConfigurationPtrOutput)
}

func (o StreamingPolicyPlayReadyConfigurationOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyPlayReadyConfiguration) *string { return v.CustomLicenseAcquisitionUrlTemplate }).(pulumi.StringPtrOutput)
}

func (o StreamingPolicyPlayReadyConfigurationOutput) PlayReadyCustomAttributes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyPlayReadyConfiguration) *string { return v.PlayReadyCustomAttributes }).(pulumi.StringPtrOutput)
}

type StreamingPolicyPlayReadyConfigurationPtrOutput struct{ *pulumi.OutputState }

func (StreamingPolicyPlayReadyConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyPlayReadyConfiguration)(nil)).Elem()
}

func (o StreamingPolicyPlayReadyConfigurationPtrOutput) ToStreamingPolicyPlayReadyConfigurationPtrOutput() StreamingPolicyPlayReadyConfigurationPtrOutput {
	return o
}

func (o StreamingPolicyPlayReadyConfigurationPtrOutput) ToStreamingPolicyPlayReadyConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationPtrOutput {
	return o
}

func (o StreamingPolicyPlayReadyConfigurationPtrOutput) Elem() StreamingPolicyPlayReadyConfigurationOutput {
	return o.ApplyT(func(v *StreamingPolicyPlayReadyConfiguration) StreamingPolicyPlayReadyConfiguration {
		if v != nil {
			return *v
		}
		var ret StreamingPolicyPlayReadyConfiguration
		return ret
	}).(StreamingPolicyPlayReadyConfigurationOutput)
}

func (o StreamingPolicyPlayReadyConfigurationPtrOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyPlayReadyConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.CustomLicenseAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

func (o StreamingPolicyPlayReadyConfigurationPtrOutput) PlayReadyCustomAttributes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyPlayReadyConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.PlayReadyCustomAttributes
	}).(pulumi.StringPtrOutput)
}

type StreamingPolicyPlayReadyConfigurationResponse struct {
	CustomLicenseAcquisitionUrlTemplate *string `pulumi:"customLicenseAcquisitionUrlTemplate"`
	PlayReadyCustomAttributes           *string `pulumi:"playReadyCustomAttributes"`
}

type StreamingPolicyPlayReadyConfigurationResponseOutput struct{ *pulumi.OutputState }

func (StreamingPolicyPlayReadyConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyPlayReadyConfigurationResponse)(nil)).Elem()
}

func (o StreamingPolicyPlayReadyConfigurationResponseOutput) ToStreamingPolicyPlayReadyConfigurationResponseOutput() StreamingPolicyPlayReadyConfigurationResponseOutput {
	return o
}

func (o StreamingPolicyPlayReadyConfigurationResponseOutput) ToStreamingPolicyPlayReadyConfigurationResponseOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationResponseOutput {
	return o
}

func (o StreamingPolicyPlayReadyConfigurationResponseOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyPlayReadyConfigurationResponse) *string {
		return v.CustomLicenseAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

func (o StreamingPolicyPlayReadyConfigurationResponseOutput) PlayReadyCustomAttributes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyPlayReadyConfigurationResponse) *string { return v.PlayReadyCustomAttributes }).(pulumi.StringPtrOutput)
}

type StreamingPolicyPlayReadyConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (StreamingPolicyPlayReadyConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyPlayReadyConfigurationResponse)(nil)).Elem()
}

func (o StreamingPolicyPlayReadyConfigurationResponsePtrOutput) ToStreamingPolicyPlayReadyConfigurationResponsePtrOutput() StreamingPolicyPlayReadyConfigurationResponsePtrOutput {
	return o
}

func (o StreamingPolicyPlayReadyConfigurationResponsePtrOutput) ToStreamingPolicyPlayReadyConfigurationResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyPlayReadyConfigurationResponsePtrOutput {
	return o
}

func (o StreamingPolicyPlayReadyConfigurationResponsePtrOutput) Elem() StreamingPolicyPlayReadyConfigurationResponseOutput {
	return o.ApplyT(func(v *StreamingPolicyPlayReadyConfigurationResponse) StreamingPolicyPlayReadyConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret StreamingPolicyPlayReadyConfigurationResponse
		return ret
	}).(StreamingPolicyPlayReadyConfigurationResponseOutput)
}

func (o StreamingPolicyPlayReadyConfigurationResponsePtrOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyPlayReadyConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomLicenseAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

func (o StreamingPolicyPlayReadyConfigurationResponsePtrOutput) PlayReadyCustomAttributes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyPlayReadyConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PlayReadyCustomAttributes
	}).(pulumi.StringPtrOutput)
}

type StreamingPolicyWidevineConfiguration struct {
	CustomLicenseAcquisitionUrlTemplate *string `pulumi:"customLicenseAcquisitionUrlTemplate"`
}





type StreamingPolicyWidevineConfigurationInput interface {
	pulumi.Input

	ToStreamingPolicyWidevineConfigurationOutput() StreamingPolicyWidevineConfigurationOutput
	ToStreamingPolicyWidevineConfigurationOutputWithContext(context.Context) StreamingPolicyWidevineConfigurationOutput
}

type StreamingPolicyWidevineConfigurationArgs struct {
	CustomLicenseAcquisitionUrlTemplate pulumi.StringPtrInput `pulumi:"customLicenseAcquisitionUrlTemplate"`
}

func (StreamingPolicyWidevineConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyWidevineConfiguration)(nil)).Elem()
}

func (i StreamingPolicyWidevineConfigurationArgs) ToStreamingPolicyWidevineConfigurationOutput() StreamingPolicyWidevineConfigurationOutput {
	return i.ToStreamingPolicyWidevineConfigurationOutputWithContext(context.Background())
}

func (i StreamingPolicyWidevineConfigurationArgs) ToStreamingPolicyWidevineConfigurationOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyWidevineConfigurationOutput)
}

func (i StreamingPolicyWidevineConfigurationArgs) ToStreamingPolicyWidevineConfigurationPtrOutput() StreamingPolicyWidevineConfigurationPtrOutput {
	return i.ToStreamingPolicyWidevineConfigurationPtrOutputWithContext(context.Background())
}

func (i StreamingPolicyWidevineConfigurationArgs) ToStreamingPolicyWidevineConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyWidevineConfigurationOutput).ToStreamingPolicyWidevineConfigurationPtrOutputWithContext(ctx)
}









type StreamingPolicyWidevineConfigurationPtrInput interface {
	pulumi.Input

	ToStreamingPolicyWidevineConfigurationPtrOutput() StreamingPolicyWidevineConfigurationPtrOutput
	ToStreamingPolicyWidevineConfigurationPtrOutputWithContext(context.Context) StreamingPolicyWidevineConfigurationPtrOutput
}

type streamingPolicyWidevineConfigurationPtrType StreamingPolicyWidevineConfigurationArgs

func StreamingPolicyWidevineConfigurationPtr(v *StreamingPolicyWidevineConfigurationArgs) StreamingPolicyWidevineConfigurationPtrInput {
	return (*streamingPolicyWidevineConfigurationPtrType)(v)
}

func (*streamingPolicyWidevineConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyWidevineConfiguration)(nil)).Elem()
}

func (i *streamingPolicyWidevineConfigurationPtrType) ToStreamingPolicyWidevineConfigurationPtrOutput() StreamingPolicyWidevineConfigurationPtrOutput {
	return i.ToStreamingPolicyWidevineConfigurationPtrOutputWithContext(context.Background())
}

func (i *streamingPolicyWidevineConfigurationPtrType) ToStreamingPolicyWidevineConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyWidevineConfigurationPtrOutput)
}

type StreamingPolicyWidevineConfigurationOutput struct{ *pulumi.OutputState }

func (StreamingPolicyWidevineConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyWidevineConfiguration)(nil)).Elem()
}

func (o StreamingPolicyWidevineConfigurationOutput) ToStreamingPolicyWidevineConfigurationOutput() StreamingPolicyWidevineConfigurationOutput {
	return o
}

func (o StreamingPolicyWidevineConfigurationOutput) ToStreamingPolicyWidevineConfigurationOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationOutput {
	return o
}

func (o StreamingPolicyWidevineConfigurationOutput) ToStreamingPolicyWidevineConfigurationPtrOutput() StreamingPolicyWidevineConfigurationPtrOutput {
	return o.ToStreamingPolicyWidevineConfigurationPtrOutputWithContext(context.Background())
}

func (o StreamingPolicyWidevineConfigurationOutput) ToStreamingPolicyWidevineConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StreamingPolicyWidevineConfiguration) *StreamingPolicyWidevineConfiguration {
		return &v
	}).(StreamingPolicyWidevineConfigurationPtrOutput)
}

func (o StreamingPolicyWidevineConfigurationOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyWidevineConfiguration) *string { return v.CustomLicenseAcquisitionUrlTemplate }).(pulumi.StringPtrOutput)
}

type StreamingPolicyWidevineConfigurationPtrOutput struct{ *pulumi.OutputState }

func (StreamingPolicyWidevineConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyWidevineConfiguration)(nil)).Elem()
}

func (o StreamingPolicyWidevineConfigurationPtrOutput) ToStreamingPolicyWidevineConfigurationPtrOutput() StreamingPolicyWidevineConfigurationPtrOutput {
	return o
}

func (o StreamingPolicyWidevineConfigurationPtrOutput) ToStreamingPolicyWidevineConfigurationPtrOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationPtrOutput {
	return o
}

func (o StreamingPolicyWidevineConfigurationPtrOutput) Elem() StreamingPolicyWidevineConfigurationOutput {
	return o.ApplyT(func(v *StreamingPolicyWidevineConfiguration) StreamingPolicyWidevineConfiguration {
		if v != nil {
			return *v
		}
		var ret StreamingPolicyWidevineConfiguration
		return ret
	}).(StreamingPolicyWidevineConfigurationOutput)
}

func (o StreamingPolicyWidevineConfigurationPtrOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyWidevineConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.CustomLicenseAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

type StreamingPolicyWidevineConfigurationResponse struct {
	CustomLicenseAcquisitionUrlTemplate *string `pulumi:"customLicenseAcquisitionUrlTemplate"`
}

type StreamingPolicyWidevineConfigurationResponseOutput struct{ *pulumi.OutputState }

func (StreamingPolicyWidevineConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamingPolicyWidevineConfigurationResponse)(nil)).Elem()
}

func (o StreamingPolicyWidevineConfigurationResponseOutput) ToStreamingPolicyWidevineConfigurationResponseOutput() StreamingPolicyWidevineConfigurationResponseOutput {
	return o
}

func (o StreamingPolicyWidevineConfigurationResponseOutput) ToStreamingPolicyWidevineConfigurationResponseOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationResponseOutput {
	return o
}

func (o StreamingPolicyWidevineConfigurationResponseOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamingPolicyWidevineConfigurationResponse) *string {
		return v.CustomLicenseAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

type StreamingPolicyWidevineConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (StreamingPolicyWidevineConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicyWidevineConfigurationResponse)(nil)).Elem()
}

func (o StreamingPolicyWidevineConfigurationResponsePtrOutput) ToStreamingPolicyWidevineConfigurationResponsePtrOutput() StreamingPolicyWidevineConfigurationResponsePtrOutput {
	return o
}

func (o StreamingPolicyWidevineConfigurationResponsePtrOutput) ToStreamingPolicyWidevineConfigurationResponsePtrOutputWithContext(ctx context.Context) StreamingPolicyWidevineConfigurationResponsePtrOutput {
	return o
}

func (o StreamingPolicyWidevineConfigurationResponsePtrOutput) Elem() StreamingPolicyWidevineConfigurationResponseOutput {
	return o.ApplyT(func(v *StreamingPolicyWidevineConfigurationResponse) StreamingPolicyWidevineConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret StreamingPolicyWidevineConfigurationResponse
		return ret
	}).(StreamingPolicyWidevineConfigurationResponseOutput)
}

func (o StreamingPolicyWidevineConfigurationResponsePtrOutput) CustomLicenseAcquisitionUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingPolicyWidevineConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomLicenseAcquisitionUrlTemplate
	}).(pulumi.StringPtrOutput)
}

type TrackPropertyCondition struct {
	Operation string  `pulumi:"operation"`
	Property  string  `pulumi:"property"`
	Value     *string `pulumi:"value"`
}





type TrackPropertyConditionInput interface {
	pulumi.Input

	ToTrackPropertyConditionOutput() TrackPropertyConditionOutput
	ToTrackPropertyConditionOutputWithContext(context.Context) TrackPropertyConditionOutput
}

type TrackPropertyConditionArgs struct {
	Operation pulumi.StringInput    `pulumi:"operation"`
	Property  pulumi.StringInput    `pulumi:"property"`
	Value     pulumi.StringPtrInput `pulumi:"value"`
}

func (TrackPropertyConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackPropertyCondition)(nil)).Elem()
}

func (i TrackPropertyConditionArgs) ToTrackPropertyConditionOutput() TrackPropertyConditionOutput {
	return i.ToTrackPropertyConditionOutputWithContext(context.Background())
}

func (i TrackPropertyConditionArgs) ToTrackPropertyConditionOutputWithContext(ctx context.Context) TrackPropertyConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackPropertyConditionOutput)
}





type TrackPropertyConditionArrayInput interface {
	pulumi.Input

	ToTrackPropertyConditionArrayOutput() TrackPropertyConditionArrayOutput
	ToTrackPropertyConditionArrayOutputWithContext(context.Context) TrackPropertyConditionArrayOutput
}

type TrackPropertyConditionArray []TrackPropertyConditionInput

func (TrackPropertyConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrackPropertyCondition)(nil)).Elem()
}

func (i TrackPropertyConditionArray) ToTrackPropertyConditionArrayOutput() TrackPropertyConditionArrayOutput {
	return i.ToTrackPropertyConditionArrayOutputWithContext(context.Background())
}

func (i TrackPropertyConditionArray) ToTrackPropertyConditionArrayOutputWithContext(ctx context.Context) TrackPropertyConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackPropertyConditionArrayOutput)
}

type TrackPropertyConditionOutput struct{ *pulumi.OutputState }

func (TrackPropertyConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackPropertyCondition)(nil)).Elem()
}

func (o TrackPropertyConditionOutput) ToTrackPropertyConditionOutput() TrackPropertyConditionOutput {
	return o
}

func (o TrackPropertyConditionOutput) ToTrackPropertyConditionOutputWithContext(ctx context.Context) TrackPropertyConditionOutput {
	return o
}

func (o TrackPropertyConditionOutput) Operation() pulumi.StringOutput {
	return o.ApplyT(func(v TrackPropertyCondition) string { return v.Operation }).(pulumi.StringOutput)
}

func (o TrackPropertyConditionOutput) Property() pulumi.StringOutput {
	return o.ApplyT(func(v TrackPropertyCondition) string { return v.Property }).(pulumi.StringOutput)
}

func (o TrackPropertyConditionOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrackPropertyCondition) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type TrackPropertyConditionArrayOutput struct{ *pulumi.OutputState }

func (TrackPropertyConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrackPropertyCondition)(nil)).Elem()
}

func (o TrackPropertyConditionArrayOutput) ToTrackPropertyConditionArrayOutput() TrackPropertyConditionArrayOutput {
	return o
}

func (o TrackPropertyConditionArrayOutput) ToTrackPropertyConditionArrayOutputWithContext(ctx context.Context) TrackPropertyConditionArrayOutput {
	return o
}

func (o TrackPropertyConditionArrayOutput) Index(i pulumi.IntInput) TrackPropertyConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrackPropertyCondition {
		return vs[0].([]TrackPropertyCondition)[vs[1].(int)]
	}).(TrackPropertyConditionOutput)
}

type TrackPropertyConditionResponse struct {
	Operation string  `pulumi:"operation"`
	Property  string  `pulumi:"property"`
	Value     *string `pulumi:"value"`
}

type TrackPropertyConditionResponseOutput struct{ *pulumi.OutputState }

func (TrackPropertyConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackPropertyConditionResponse)(nil)).Elem()
}

func (o TrackPropertyConditionResponseOutput) ToTrackPropertyConditionResponseOutput() TrackPropertyConditionResponseOutput {
	return o
}

func (o TrackPropertyConditionResponseOutput) ToTrackPropertyConditionResponseOutputWithContext(ctx context.Context) TrackPropertyConditionResponseOutput {
	return o
}

func (o TrackPropertyConditionResponseOutput) Operation() pulumi.StringOutput {
	return o.ApplyT(func(v TrackPropertyConditionResponse) string { return v.Operation }).(pulumi.StringOutput)
}

func (o TrackPropertyConditionResponseOutput) Property() pulumi.StringOutput {
	return o.ApplyT(func(v TrackPropertyConditionResponse) string { return v.Property }).(pulumi.StringOutput)
}

func (o TrackPropertyConditionResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrackPropertyConditionResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type TrackPropertyConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (TrackPropertyConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrackPropertyConditionResponse)(nil)).Elem()
}

func (o TrackPropertyConditionResponseArrayOutput) ToTrackPropertyConditionResponseArrayOutput() TrackPropertyConditionResponseArrayOutput {
	return o
}

func (o TrackPropertyConditionResponseArrayOutput) ToTrackPropertyConditionResponseArrayOutputWithContext(ctx context.Context) TrackPropertyConditionResponseArrayOutput {
	return o
}

func (o TrackPropertyConditionResponseArrayOutput) Index(i pulumi.IntInput) TrackPropertyConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrackPropertyConditionResponse {
		return vs[0].([]TrackPropertyConditionResponse)[vs[1].(int)]
	}).(TrackPropertyConditionResponseOutput)
}

type TrackSelection struct {
	TrackSelections []TrackPropertyCondition `pulumi:"trackSelections"`
}





type TrackSelectionInput interface {
	pulumi.Input

	ToTrackSelectionOutput() TrackSelectionOutput
	ToTrackSelectionOutputWithContext(context.Context) TrackSelectionOutput
}

type TrackSelectionArgs struct {
	TrackSelections TrackPropertyConditionArrayInput `pulumi:"trackSelections"`
}

func (TrackSelectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackSelection)(nil)).Elem()
}

func (i TrackSelectionArgs) ToTrackSelectionOutput() TrackSelectionOutput {
	return i.ToTrackSelectionOutputWithContext(context.Background())
}

func (i TrackSelectionArgs) ToTrackSelectionOutputWithContext(ctx context.Context) TrackSelectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackSelectionOutput)
}





type TrackSelectionArrayInput interface {
	pulumi.Input

	ToTrackSelectionArrayOutput() TrackSelectionArrayOutput
	ToTrackSelectionArrayOutputWithContext(context.Context) TrackSelectionArrayOutput
}

type TrackSelectionArray []TrackSelectionInput

func (TrackSelectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrackSelection)(nil)).Elem()
}

func (i TrackSelectionArray) ToTrackSelectionArrayOutput() TrackSelectionArrayOutput {
	return i.ToTrackSelectionArrayOutputWithContext(context.Background())
}

func (i TrackSelectionArray) ToTrackSelectionArrayOutputWithContext(ctx context.Context) TrackSelectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackSelectionArrayOutput)
}

type TrackSelectionOutput struct{ *pulumi.OutputState }

func (TrackSelectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackSelection)(nil)).Elem()
}

func (o TrackSelectionOutput) ToTrackSelectionOutput() TrackSelectionOutput {
	return o
}

func (o TrackSelectionOutput) ToTrackSelectionOutputWithContext(ctx context.Context) TrackSelectionOutput {
	return o
}

func (o TrackSelectionOutput) TrackSelections() TrackPropertyConditionArrayOutput {
	return o.ApplyT(func(v TrackSelection) []TrackPropertyCondition { return v.TrackSelections }).(TrackPropertyConditionArrayOutput)
}

type TrackSelectionArrayOutput struct{ *pulumi.OutputState }

func (TrackSelectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrackSelection)(nil)).Elem()
}

func (o TrackSelectionArrayOutput) ToTrackSelectionArrayOutput() TrackSelectionArrayOutput {
	return o
}

func (o TrackSelectionArrayOutput) ToTrackSelectionArrayOutputWithContext(ctx context.Context) TrackSelectionArrayOutput {
	return o
}

func (o TrackSelectionArrayOutput) Index(i pulumi.IntInput) TrackSelectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrackSelection {
		return vs[0].([]TrackSelection)[vs[1].(int)]
	}).(TrackSelectionOutput)
}

type TrackSelectionResponse struct {
	TrackSelections []TrackPropertyConditionResponse `pulumi:"trackSelections"`
}

type TrackSelectionResponseOutput struct{ *pulumi.OutputState }

func (TrackSelectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackSelectionResponse)(nil)).Elem()
}

func (o TrackSelectionResponseOutput) ToTrackSelectionResponseOutput() TrackSelectionResponseOutput {
	return o
}

func (o TrackSelectionResponseOutput) ToTrackSelectionResponseOutputWithContext(ctx context.Context) TrackSelectionResponseOutput {
	return o
}

func (o TrackSelectionResponseOutput) TrackSelections() TrackPropertyConditionResponseArrayOutput {
	return o.ApplyT(func(v TrackSelectionResponse) []TrackPropertyConditionResponse { return v.TrackSelections }).(TrackPropertyConditionResponseArrayOutput)
}

type TrackSelectionResponseArrayOutput struct{ *pulumi.OutputState }

func (TrackSelectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrackSelectionResponse)(nil)).Elem()
}

func (o TrackSelectionResponseArrayOutput) ToTrackSelectionResponseArrayOutput() TrackSelectionResponseArrayOutput {
	return o
}

func (o TrackSelectionResponseArrayOutput) ToTrackSelectionResponseArrayOutputWithContext(ctx context.Context) TrackSelectionResponseArrayOutput {
	return o
}

func (o TrackSelectionResponseArrayOutput) Index(i pulumi.IntInput) TrackSelectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrackSelectionResponse {
		return vs[0].([]TrackSelectionResponse)[vs[1].(int)]
	}).(TrackSelectionResponseOutput)
}

type TransformOutputType struct {
	OnError          *string     `pulumi:"onError"`
	Preset           interface{} `pulumi:"preset"`
	RelativePriority *string     `pulumi:"relativePriority"`
}





type TransformOutputTypeInput interface {
	pulumi.Input

	ToTransformOutputTypeOutput() TransformOutputTypeOutput
	ToTransformOutputTypeOutputWithContext(context.Context) TransformOutputTypeOutput
}

type TransformOutputTypeArgs struct {
	OnError          pulumi.StringPtrInput `pulumi:"onError"`
	Preset           pulumi.Input          `pulumi:"preset"`
	RelativePriority pulumi.StringPtrInput `pulumi:"relativePriority"`
}

func (TransformOutputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransformOutputType)(nil)).Elem()
}

func (i TransformOutputTypeArgs) ToTransformOutputTypeOutput() TransformOutputTypeOutput {
	return i.ToTransformOutputTypeOutputWithContext(context.Background())
}

func (i TransformOutputTypeArgs) ToTransformOutputTypeOutputWithContext(ctx context.Context) TransformOutputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformOutputTypeOutput)
}





type TransformOutputTypeArrayInput interface {
	pulumi.Input

	ToTransformOutputTypeArrayOutput() TransformOutputTypeArrayOutput
	ToTransformOutputTypeArrayOutputWithContext(context.Context) TransformOutputTypeArrayOutput
}

type TransformOutputTypeArray []TransformOutputTypeInput

func (TransformOutputTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TransformOutputType)(nil)).Elem()
}

func (i TransformOutputTypeArray) ToTransformOutputTypeArrayOutput() TransformOutputTypeArrayOutput {
	return i.ToTransformOutputTypeArrayOutputWithContext(context.Background())
}

func (i TransformOutputTypeArray) ToTransformOutputTypeArrayOutputWithContext(ctx context.Context) TransformOutputTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformOutputTypeArrayOutput)
}

type TransformOutputTypeOutput struct{ *pulumi.OutputState }

func (TransformOutputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransformOutputType)(nil)).Elem()
}

func (o TransformOutputTypeOutput) ToTransformOutputTypeOutput() TransformOutputTypeOutput {
	return o
}

func (o TransformOutputTypeOutput) ToTransformOutputTypeOutputWithContext(ctx context.Context) TransformOutputTypeOutput {
	return o
}

func (o TransformOutputTypeOutput) OnError() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransformOutputType) *string { return v.OnError }).(pulumi.StringPtrOutput)
}

func (o TransformOutputTypeOutput) Preset() pulumi.AnyOutput {
	return o.ApplyT(func(v TransformOutputType) interface{} { return v.Preset }).(pulumi.AnyOutput)
}

func (o TransformOutputTypeOutput) RelativePriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransformOutputType) *string { return v.RelativePriority }).(pulumi.StringPtrOutput)
}

type TransformOutputTypeArrayOutput struct{ *pulumi.OutputState }

func (TransformOutputTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TransformOutputType)(nil)).Elem()
}

func (o TransformOutputTypeArrayOutput) ToTransformOutputTypeArrayOutput() TransformOutputTypeArrayOutput {
	return o
}

func (o TransformOutputTypeArrayOutput) ToTransformOutputTypeArrayOutputWithContext(ctx context.Context) TransformOutputTypeArrayOutput {
	return o
}

func (o TransformOutputTypeArrayOutput) Index(i pulumi.IntInput) TransformOutputTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TransformOutputType {
		return vs[0].([]TransformOutputType)[vs[1].(int)]
	}).(TransformOutputTypeOutput)
}

type TransformOutputResponse struct {
	OnError          *string     `pulumi:"onError"`
	Preset           interface{} `pulumi:"preset"`
	RelativePriority *string     `pulumi:"relativePriority"`
}

type TransformOutputResponseOutput struct{ *pulumi.OutputState }

func (TransformOutputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransformOutputResponse)(nil)).Elem()
}

func (o TransformOutputResponseOutput) ToTransformOutputResponseOutput() TransformOutputResponseOutput {
	return o
}

func (o TransformOutputResponseOutput) ToTransformOutputResponseOutputWithContext(ctx context.Context) TransformOutputResponseOutput {
	return o
}

func (o TransformOutputResponseOutput) OnError() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransformOutputResponse) *string { return v.OnError }).(pulumi.StringPtrOutput)
}

func (o TransformOutputResponseOutput) Preset() pulumi.AnyOutput {
	return o.ApplyT(func(v TransformOutputResponse) interface{} { return v.Preset }).(pulumi.AnyOutput)
}

func (o TransformOutputResponseOutput) RelativePriority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransformOutputResponse) *string { return v.RelativePriority }).(pulumi.StringPtrOutput)
}

type TransformOutputResponseArrayOutput struct{ *pulumi.OutputState }

func (TransformOutputResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TransformOutputResponse)(nil)).Elem()
}

func (o TransformOutputResponseArrayOutput) ToTransformOutputResponseArrayOutput() TransformOutputResponseArrayOutput {
	return o
}

func (o TransformOutputResponseArrayOutput) ToTransformOutputResponseArrayOutputWithContext(ctx context.Context) TransformOutputResponseArrayOutput {
	return o
}

func (o TransformOutputResponseArrayOutput) Index(i pulumi.IntInput) TransformOutputResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TransformOutputResponse {
		return vs[0].([]TransformOutputResponse)[vs[1].(int)]
	}).(TransformOutputResponseOutput)
}

type TransportStreamFormat struct {
	FilenamePattern *string      `pulumi:"filenamePattern"`
	OdataType       string       `pulumi:"odataType"`
	OutputFiles     []OutputFile `pulumi:"outputFiles"`
}

type TransportStreamFormatResponse struct {
	FilenamePattern *string              `pulumi:"filenamePattern"`
	OdataType       string               `pulumi:"odataType"`
	OutputFiles     []OutputFileResponse `pulumi:"outputFiles"`
}

type Video struct {
	KeyFrameInterval *string `pulumi:"keyFrameInterval"`
	Label            *string `pulumi:"label"`
	OdataType        string  `pulumi:"odataType"`
	StretchMode      *string `pulumi:"stretchMode"`
}

type VideoAnalyzerPreset struct {
	AudioInsightsOnly *bool   `pulumi:"audioInsightsOnly"`
	AudioLanguage     *string `pulumi:"audioLanguage"`
	OdataType         string  `pulumi:"odataType"`
}

type VideoAnalyzerPresetResponse struct {
	AudioInsightsOnly *bool   `pulumi:"audioInsightsOnly"`
	AudioLanguage     *string `pulumi:"audioLanguage"`
	OdataType         string  `pulumi:"odataType"`
}

type VideoOverlay struct {
	AudioGainLevel  *float64   `pulumi:"audioGainLevel"`
	CropRectangle   *Rectangle `pulumi:"cropRectangle"`
	End             *string    `pulumi:"end"`
	FadeInDuration  *string    `pulumi:"fadeInDuration"`
	FadeOutDuration *string    `pulumi:"fadeOutDuration"`
	InputLabel      *string    `pulumi:"inputLabel"`
	OdataType       string     `pulumi:"odataType"`
	Opacity         *float64   `pulumi:"opacity"`
	Position        *Rectangle `pulumi:"position"`
	Start           *string    `pulumi:"start"`
}

type VideoOverlayResponse struct {
	AudioGainLevel  *float64           `pulumi:"audioGainLevel"`
	CropRectangle   *RectangleResponse `pulumi:"cropRectangle"`
	End             *string            `pulumi:"end"`
	FadeInDuration  *string            `pulumi:"fadeInDuration"`
	FadeOutDuration *string            `pulumi:"fadeOutDuration"`
	InputLabel      *string            `pulumi:"inputLabel"`
	OdataType       string             `pulumi:"odataType"`
	Opacity         *float64           `pulumi:"opacity"`
	Position        *RectangleResponse `pulumi:"position"`
	Start           *string            `pulumi:"start"`
}

type VideoResponse struct {
	KeyFrameInterval *string `pulumi:"keyFrameInterval"`
	Label            *string `pulumi:"label"`
	OdataType        string  `pulumi:"odataType"`
	StretchMode      *string `pulumi:"stretchMode"`
}

func init() {
	pulumi.RegisterOutputType(AkamaiAccessControlOutput{})
	pulumi.RegisterOutputType(AkamaiAccessControlPtrOutput{})
	pulumi.RegisterOutputType(AkamaiAccessControlResponseOutput{})
	pulumi.RegisterOutputType(AkamaiAccessControlResponsePtrOutput{})
	pulumi.RegisterOutputType(AkamaiSignatureHeaderAuthenticationKeyOutput{})
	pulumi.RegisterOutputType(AkamaiSignatureHeaderAuthenticationKeyArrayOutput{})
	pulumi.RegisterOutputType(AkamaiSignatureHeaderAuthenticationKeyResponseOutput{})
	pulumi.RegisterOutputType(AkamaiSignatureHeaderAuthenticationKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(CbcsDrmConfigurationOutput{})
	pulumi.RegisterOutputType(CbcsDrmConfigurationPtrOutput{})
	pulumi.RegisterOutputType(CbcsDrmConfigurationResponseOutput{})
	pulumi.RegisterOutputType(CbcsDrmConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(CencDrmConfigurationOutput{})
	pulumi.RegisterOutputType(CencDrmConfigurationPtrOutput{})
	pulumi.RegisterOutputType(CencDrmConfigurationResponseOutput{})
	pulumi.RegisterOutputType(CencDrmConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(CommonEncryptionCbcsOutput{})
	pulumi.RegisterOutputType(CommonEncryptionCbcsPtrOutput{})
	pulumi.RegisterOutputType(CommonEncryptionCbcsResponseOutput{})
	pulumi.RegisterOutputType(CommonEncryptionCbcsResponsePtrOutput{})
	pulumi.RegisterOutputType(CommonEncryptionCencOutput{})
	pulumi.RegisterOutputType(CommonEncryptionCencPtrOutput{})
	pulumi.RegisterOutputType(CommonEncryptionCencResponseOutput{})
	pulumi.RegisterOutputType(CommonEncryptionCencResponsePtrOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyOptionOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyOptionArrayOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyOptionResponseOutput{})
	pulumi.RegisterOutputType(ContentKeyPolicyOptionResponseArrayOutput{})
	pulumi.RegisterOutputType(CrossSiteAccessPoliciesOutput{})
	pulumi.RegisterOutputType(CrossSiteAccessPoliciesPtrOutput{})
	pulumi.RegisterOutputType(CrossSiteAccessPoliciesResponseOutput{})
	pulumi.RegisterOutputType(CrossSiteAccessPoliciesResponsePtrOutput{})
	pulumi.RegisterOutputType(DefaultKeyOutput{})
	pulumi.RegisterOutputType(DefaultKeyPtrOutput{})
	pulumi.RegisterOutputType(DefaultKeyResponseOutput{})
	pulumi.RegisterOutputType(DefaultKeyResponsePtrOutput{})
	pulumi.RegisterOutputType(EnabledProtocolsOutput{})
	pulumi.RegisterOutputType(EnabledProtocolsPtrOutput{})
	pulumi.RegisterOutputType(EnabledProtocolsResponseOutput{})
	pulumi.RegisterOutputType(EnabledProtocolsResponsePtrOutput{})
	pulumi.RegisterOutputType(EnvelopeEncryptionOutput{})
	pulumi.RegisterOutputType(EnvelopeEncryptionPtrOutput{})
	pulumi.RegisterOutputType(EnvelopeEncryptionResponseOutput{})
	pulumi.RegisterOutputType(EnvelopeEncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(HlsOutput{})
	pulumi.RegisterOutputType(HlsPtrOutput{})
	pulumi.RegisterOutputType(HlsResponseOutput{})
	pulumi.RegisterOutputType(HlsResponsePtrOutput{})
	pulumi.RegisterOutputType(IPAccessControlOutput{})
	pulumi.RegisterOutputType(IPAccessControlPtrOutput{})
	pulumi.RegisterOutputType(IPAccessControlResponseOutput{})
	pulumi.RegisterOutputType(IPAccessControlResponsePtrOutput{})
	pulumi.RegisterOutputType(IPRangeOutput{})
	pulumi.RegisterOutputType(IPRangeArrayOutput{})
	pulumi.RegisterOutputType(IPRangeResponseOutput{})
	pulumi.RegisterOutputType(IPRangeResponseArrayOutput{})
	pulumi.RegisterOutputType(JobErrorDetailResponseOutput{})
	pulumi.RegisterOutputType(JobErrorDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(JobErrorResponseOutput{})
	pulumi.RegisterOutputType(JobOutputAssetOutput{})
	pulumi.RegisterOutputType(JobOutputAssetArrayOutput{})
	pulumi.RegisterOutputType(JobOutputAssetResponseOutput{})
	pulumi.RegisterOutputType(JobOutputAssetResponseArrayOutput{})
	pulumi.RegisterOutputType(LiveEventEncodingOutput{})
	pulumi.RegisterOutputType(LiveEventEncodingPtrOutput{})
	pulumi.RegisterOutputType(LiveEventEncodingResponseOutput{})
	pulumi.RegisterOutputType(LiveEventEncodingResponsePtrOutput{})
	pulumi.RegisterOutputType(LiveEventEndpointOutput{})
	pulumi.RegisterOutputType(LiveEventEndpointArrayOutput{})
	pulumi.RegisterOutputType(LiveEventEndpointResponseOutput{})
	pulumi.RegisterOutputType(LiveEventEndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(LiveEventInputTypeOutput{})
	pulumi.RegisterOutputType(LiveEventInputResponseOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewPtrOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewAccessControlOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewAccessControlPtrOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewAccessControlResponseOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewAccessControlResponsePtrOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewResponseOutput{})
	pulumi.RegisterOutputType(LiveEventPreviewResponsePtrOutput{})
	pulumi.RegisterOutputType(NoEncryptionOutput{})
	pulumi.RegisterOutputType(NoEncryptionPtrOutput{})
	pulumi.RegisterOutputType(NoEncryptionResponseOutput{})
	pulumi.RegisterOutputType(NoEncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountOutput{})
	pulumi.RegisterOutputType(StorageAccountArrayOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseArrayOutput{})
	pulumi.RegisterOutputType(StreamingEndpointAccessControlOutput{})
	pulumi.RegisterOutputType(StreamingEndpointAccessControlPtrOutput{})
	pulumi.RegisterOutputType(StreamingEndpointAccessControlResponseOutput{})
	pulumi.RegisterOutputType(StreamingEndpointAccessControlResponsePtrOutput{})
	pulumi.RegisterOutputType(StreamingLocatorContentKeyOutput{})
	pulumi.RegisterOutputType(StreamingLocatorContentKeyArrayOutput{})
	pulumi.RegisterOutputType(StreamingLocatorContentKeyResponseOutput{})
	pulumi.RegisterOutputType(StreamingLocatorContentKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(StreamingPathResponseOutput{})
	pulumi.RegisterOutputType(StreamingPathResponseArrayOutput{})
	pulumi.RegisterOutputType(StreamingPolicyContentKeyOutput{})
	pulumi.RegisterOutputType(StreamingPolicyContentKeyArrayOutput{})
	pulumi.RegisterOutputType(StreamingPolicyContentKeyResponseOutput{})
	pulumi.RegisterOutputType(StreamingPolicyContentKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(StreamingPolicyContentKeysOutput{})
	pulumi.RegisterOutputType(StreamingPolicyContentKeysPtrOutput{})
	pulumi.RegisterOutputType(StreamingPolicyContentKeysResponseOutput{})
	pulumi.RegisterOutputType(StreamingPolicyContentKeysResponsePtrOutput{})
	pulumi.RegisterOutputType(StreamingPolicyFairPlayConfigurationOutput{})
	pulumi.RegisterOutputType(StreamingPolicyFairPlayConfigurationPtrOutput{})
	pulumi.RegisterOutputType(StreamingPolicyFairPlayConfigurationResponseOutput{})
	pulumi.RegisterOutputType(StreamingPolicyFairPlayConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(StreamingPolicyPlayReadyConfigurationOutput{})
	pulumi.RegisterOutputType(StreamingPolicyPlayReadyConfigurationPtrOutput{})
	pulumi.RegisterOutputType(StreamingPolicyPlayReadyConfigurationResponseOutput{})
	pulumi.RegisterOutputType(StreamingPolicyPlayReadyConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(StreamingPolicyWidevineConfigurationOutput{})
	pulumi.RegisterOutputType(StreamingPolicyWidevineConfigurationPtrOutput{})
	pulumi.RegisterOutputType(StreamingPolicyWidevineConfigurationResponseOutput{})
	pulumi.RegisterOutputType(StreamingPolicyWidevineConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(TrackPropertyConditionOutput{})
	pulumi.RegisterOutputType(TrackPropertyConditionArrayOutput{})
	pulumi.RegisterOutputType(TrackPropertyConditionResponseOutput{})
	pulumi.RegisterOutputType(TrackPropertyConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(TrackSelectionOutput{})
	pulumi.RegisterOutputType(TrackSelectionArrayOutput{})
	pulumi.RegisterOutputType(TrackSelectionResponseOutput{})
	pulumi.RegisterOutputType(TrackSelectionResponseArrayOutput{})
	pulumi.RegisterOutputType(TransformOutputTypeOutput{})
	pulumi.RegisterOutputType(TransformOutputTypeArrayOutput{})
	pulumi.RegisterOutputType(TransformOutputResponseOutput{})
	pulumi.RegisterOutputType(TransformOutputResponseArrayOutput{})
}
