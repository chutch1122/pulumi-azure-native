


package logic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AS2AcknowledgementConnectionSettings struct {
	IgnoreCertificateNameMismatch bool `pulumi:"ignoreCertificateNameMismatch"`
	KeepHttpConnectionAlive       bool `pulumi:"keepHttpConnectionAlive"`
	SupportHttpStatusCodeContinue bool `pulumi:"supportHttpStatusCodeContinue"`
	UnfoldHttpHeaders             bool `pulumi:"unfoldHttpHeaders"`
}





type AS2AcknowledgementConnectionSettingsInput interface {
	pulumi.Input

	ToAS2AcknowledgementConnectionSettingsOutput() AS2AcknowledgementConnectionSettingsOutput
	ToAS2AcknowledgementConnectionSettingsOutputWithContext(context.Context) AS2AcknowledgementConnectionSettingsOutput
}

type AS2AcknowledgementConnectionSettingsArgs struct {
	IgnoreCertificateNameMismatch pulumi.BoolInput `pulumi:"ignoreCertificateNameMismatch"`
	KeepHttpConnectionAlive       pulumi.BoolInput `pulumi:"keepHttpConnectionAlive"`
	SupportHttpStatusCodeContinue pulumi.BoolInput `pulumi:"supportHttpStatusCodeContinue"`
	UnfoldHttpHeaders             pulumi.BoolInput `pulumi:"unfoldHttpHeaders"`
}

func (AS2AcknowledgementConnectionSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2AcknowledgementConnectionSettings)(nil)).Elem()
}

func (i AS2AcknowledgementConnectionSettingsArgs) ToAS2AcknowledgementConnectionSettingsOutput() AS2AcknowledgementConnectionSettingsOutput {
	return i.ToAS2AcknowledgementConnectionSettingsOutputWithContext(context.Background())
}

func (i AS2AcknowledgementConnectionSettingsArgs) ToAS2AcknowledgementConnectionSettingsOutputWithContext(ctx context.Context) AS2AcknowledgementConnectionSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2AcknowledgementConnectionSettingsOutput)
}

func (i AS2AcknowledgementConnectionSettingsArgs) ToAS2AcknowledgementConnectionSettingsPtrOutput() AS2AcknowledgementConnectionSettingsPtrOutput {
	return i.ToAS2AcknowledgementConnectionSettingsPtrOutputWithContext(context.Background())
}

func (i AS2AcknowledgementConnectionSettingsArgs) ToAS2AcknowledgementConnectionSettingsPtrOutputWithContext(ctx context.Context) AS2AcknowledgementConnectionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2AcknowledgementConnectionSettingsOutput).ToAS2AcknowledgementConnectionSettingsPtrOutputWithContext(ctx)
}









type AS2AcknowledgementConnectionSettingsPtrInput interface {
	pulumi.Input

	ToAS2AcknowledgementConnectionSettingsPtrOutput() AS2AcknowledgementConnectionSettingsPtrOutput
	ToAS2AcknowledgementConnectionSettingsPtrOutputWithContext(context.Context) AS2AcknowledgementConnectionSettingsPtrOutput
}

type as2acknowledgementConnectionSettingsPtrType AS2AcknowledgementConnectionSettingsArgs

func AS2AcknowledgementConnectionSettingsPtr(v *AS2AcknowledgementConnectionSettingsArgs) AS2AcknowledgementConnectionSettingsPtrInput {
	return (*as2acknowledgementConnectionSettingsPtrType)(v)
}

func (*as2acknowledgementConnectionSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2AcknowledgementConnectionSettings)(nil)).Elem()
}

func (i *as2acknowledgementConnectionSettingsPtrType) ToAS2AcknowledgementConnectionSettingsPtrOutput() AS2AcknowledgementConnectionSettingsPtrOutput {
	return i.ToAS2AcknowledgementConnectionSettingsPtrOutputWithContext(context.Background())
}

func (i *as2acknowledgementConnectionSettingsPtrType) ToAS2AcknowledgementConnectionSettingsPtrOutputWithContext(ctx context.Context) AS2AcknowledgementConnectionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2AcknowledgementConnectionSettingsPtrOutput)
}

type AS2AcknowledgementConnectionSettingsOutput struct{ *pulumi.OutputState }

func (AS2AcknowledgementConnectionSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2AcknowledgementConnectionSettings)(nil)).Elem()
}

func (o AS2AcknowledgementConnectionSettingsOutput) ToAS2AcknowledgementConnectionSettingsOutput() AS2AcknowledgementConnectionSettingsOutput {
	return o
}

func (o AS2AcknowledgementConnectionSettingsOutput) ToAS2AcknowledgementConnectionSettingsOutputWithContext(ctx context.Context) AS2AcknowledgementConnectionSettingsOutput {
	return o
}

func (o AS2AcknowledgementConnectionSettingsOutput) ToAS2AcknowledgementConnectionSettingsPtrOutput() AS2AcknowledgementConnectionSettingsPtrOutput {
	return o.ToAS2AcknowledgementConnectionSettingsPtrOutputWithContext(context.Background())
}

func (o AS2AcknowledgementConnectionSettingsOutput) ToAS2AcknowledgementConnectionSettingsPtrOutputWithContext(ctx context.Context) AS2AcknowledgementConnectionSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2AcknowledgementConnectionSettings) *AS2AcknowledgementConnectionSettings {
		return &v
	}).(AS2AcknowledgementConnectionSettingsPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsOutput) IgnoreCertificateNameMismatch() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2AcknowledgementConnectionSettings) bool { return v.IgnoreCertificateNameMismatch }).(pulumi.BoolOutput)
}

func (o AS2AcknowledgementConnectionSettingsOutput) KeepHttpConnectionAlive() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2AcknowledgementConnectionSettings) bool { return v.KeepHttpConnectionAlive }).(pulumi.BoolOutput)
}

func (o AS2AcknowledgementConnectionSettingsOutput) SupportHttpStatusCodeContinue() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2AcknowledgementConnectionSettings) bool { return v.SupportHttpStatusCodeContinue }).(pulumi.BoolOutput)
}

func (o AS2AcknowledgementConnectionSettingsOutput) UnfoldHttpHeaders() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2AcknowledgementConnectionSettings) bool { return v.UnfoldHttpHeaders }).(pulumi.BoolOutput)
}

type AS2AcknowledgementConnectionSettingsPtrOutput struct{ *pulumi.OutputState }

func (AS2AcknowledgementConnectionSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2AcknowledgementConnectionSettings)(nil)).Elem()
}

func (o AS2AcknowledgementConnectionSettingsPtrOutput) ToAS2AcknowledgementConnectionSettingsPtrOutput() AS2AcknowledgementConnectionSettingsPtrOutput {
	return o
}

func (o AS2AcknowledgementConnectionSettingsPtrOutput) ToAS2AcknowledgementConnectionSettingsPtrOutputWithContext(ctx context.Context) AS2AcknowledgementConnectionSettingsPtrOutput {
	return o
}

func (o AS2AcknowledgementConnectionSettingsPtrOutput) Elem() AS2AcknowledgementConnectionSettingsOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettings) AS2AcknowledgementConnectionSettings {
		if v != nil {
			return *v
		}
		var ret AS2AcknowledgementConnectionSettings
		return ret
	}).(AS2AcknowledgementConnectionSettingsOutput)
}

func (o AS2AcknowledgementConnectionSettingsPtrOutput) IgnoreCertificateNameMismatch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.IgnoreCertificateNameMismatch
	}).(pulumi.BoolPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsPtrOutput) KeepHttpConnectionAlive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.KeepHttpConnectionAlive
	}).(pulumi.BoolPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsPtrOutput) SupportHttpStatusCodeContinue() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.SupportHttpStatusCodeContinue
	}).(pulumi.BoolPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsPtrOutput) UnfoldHttpHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.UnfoldHttpHeaders
	}).(pulumi.BoolPtrOutput)
}

type AS2AcknowledgementConnectionSettingsResponse struct {
	IgnoreCertificateNameMismatch bool `pulumi:"ignoreCertificateNameMismatch"`
	KeepHttpConnectionAlive       bool `pulumi:"keepHttpConnectionAlive"`
	SupportHttpStatusCodeContinue bool `pulumi:"supportHttpStatusCodeContinue"`
	UnfoldHttpHeaders             bool `pulumi:"unfoldHttpHeaders"`
}

type AS2AcknowledgementConnectionSettingsResponseOutput struct{ *pulumi.OutputState }

func (AS2AcknowledgementConnectionSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2AcknowledgementConnectionSettingsResponse)(nil)).Elem()
}

func (o AS2AcknowledgementConnectionSettingsResponseOutput) ToAS2AcknowledgementConnectionSettingsResponseOutput() AS2AcknowledgementConnectionSettingsResponseOutput {
	return o
}

func (o AS2AcknowledgementConnectionSettingsResponseOutput) ToAS2AcknowledgementConnectionSettingsResponseOutputWithContext(ctx context.Context) AS2AcknowledgementConnectionSettingsResponseOutput {
	return o
}

func (o AS2AcknowledgementConnectionSettingsResponseOutput) IgnoreCertificateNameMismatch() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2AcknowledgementConnectionSettingsResponse) bool { return v.IgnoreCertificateNameMismatch }).(pulumi.BoolOutput)
}

func (o AS2AcknowledgementConnectionSettingsResponseOutput) KeepHttpConnectionAlive() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2AcknowledgementConnectionSettingsResponse) bool { return v.KeepHttpConnectionAlive }).(pulumi.BoolOutput)
}

func (o AS2AcknowledgementConnectionSettingsResponseOutput) SupportHttpStatusCodeContinue() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2AcknowledgementConnectionSettingsResponse) bool { return v.SupportHttpStatusCodeContinue }).(pulumi.BoolOutput)
}

func (o AS2AcknowledgementConnectionSettingsResponseOutput) UnfoldHttpHeaders() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2AcknowledgementConnectionSettingsResponse) bool { return v.UnfoldHttpHeaders }).(pulumi.BoolOutput)
}

type AS2AcknowledgementConnectionSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2AcknowledgementConnectionSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2AcknowledgementConnectionSettingsResponse)(nil)).Elem()
}

func (o AS2AcknowledgementConnectionSettingsResponsePtrOutput) ToAS2AcknowledgementConnectionSettingsResponsePtrOutput() AS2AcknowledgementConnectionSettingsResponsePtrOutput {
	return o
}

func (o AS2AcknowledgementConnectionSettingsResponsePtrOutput) ToAS2AcknowledgementConnectionSettingsResponsePtrOutputWithContext(ctx context.Context) AS2AcknowledgementConnectionSettingsResponsePtrOutput {
	return o
}

func (o AS2AcknowledgementConnectionSettingsResponsePtrOutput) Elem() AS2AcknowledgementConnectionSettingsResponseOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettingsResponse) AS2AcknowledgementConnectionSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AS2AcknowledgementConnectionSettingsResponse
		return ret
	}).(AS2AcknowledgementConnectionSettingsResponseOutput)
}

func (o AS2AcknowledgementConnectionSettingsResponsePtrOutput) IgnoreCertificateNameMismatch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IgnoreCertificateNameMismatch
	}).(pulumi.BoolPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsResponsePtrOutput) KeepHttpConnectionAlive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.KeepHttpConnectionAlive
	}).(pulumi.BoolPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsResponsePtrOutput) SupportHttpStatusCodeContinue() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.SupportHttpStatusCodeContinue
	}).(pulumi.BoolPtrOutput)
}

func (o AS2AcknowledgementConnectionSettingsResponsePtrOutput) UnfoldHttpHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2AcknowledgementConnectionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.UnfoldHttpHeaders
	}).(pulumi.BoolPtrOutput)
}

type AS2AgreementContent struct {
	ReceiveAgreement AS2OneWayAgreement `pulumi:"receiveAgreement"`
	SendAgreement    AS2OneWayAgreement `pulumi:"sendAgreement"`
}





type AS2AgreementContentInput interface {
	pulumi.Input

	ToAS2AgreementContentOutput() AS2AgreementContentOutput
	ToAS2AgreementContentOutputWithContext(context.Context) AS2AgreementContentOutput
}

type AS2AgreementContentArgs struct {
	ReceiveAgreement AS2OneWayAgreementInput `pulumi:"receiveAgreement"`
	SendAgreement    AS2OneWayAgreementInput `pulumi:"sendAgreement"`
}

func (AS2AgreementContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2AgreementContent)(nil)).Elem()
}

func (i AS2AgreementContentArgs) ToAS2AgreementContentOutput() AS2AgreementContentOutput {
	return i.ToAS2AgreementContentOutputWithContext(context.Background())
}

func (i AS2AgreementContentArgs) ToAS2AgreementContentOutputWithContext(ctx context.Context) AS2AgreementContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2AgreementContentOutput)
}

func (i AS2AgreementContentArgs) ToAS2AgreementContentPtrOutput() AS2AgreementContentPtrOutput {
	return i.ToAS2AgreementContentPtrOutputWithContext(context.Background())
}

func (i AS2AgreementContentArgs) ToAS2AgreementContentPtrOutputWithContext(ctx context.Context) AS2AgreementContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2AgreementContentOutput).ToAS2AgreementContentPtrOutputWithContext(ctx)
}









type AS2AgreementContentPtrInput interface {
	pulumi.Input

	ToAS2AgreementContentPtrOutput() AS2AgreementContentPtrOutput
	ToAS2AgreementContentPtrOutputWithContext(context.Context) AS2AgreementContentPtrOutput
}

type as2agreementContentPtrType AS2AgreementContentArgs

func AS2AgreementContentPtr(v *AS2AgreementContentArgs) AS2AgreementContentPtrInput {
	return (*as2agreementContentPtrType)(v)
}

func (*as2agreementContentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2AgreementContent)(nil)).Elem()
}

func (i *as2agreementContentPtrType) ToAS2AgreementContentPtrOutput() AS2AgreementContentPtrOutput {
	return i.ToAS2AgreementContentPtrOutputWithContext(context.Background())
}

func (i *as2agreementContentPtrType) ToAS2AgreementContentPtrOutputWithContext(ctx context.Context) AS2AgreementContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2AgreementContentPtrOutput)
}

type AS2AgreementContentOutput struct{ *pulumi.OutputState }

func (AS2AgreementContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2AgreementContent)(nil)).Elem()
}

func (o AS2AgreementContentOutput) ToAS2AgreementContentOutput() AS2AgreementContentOutput {
	return o
}

func (o AS2AgreementContentOutput) ToAS2AgreementContentOutputWithContext(ctx context.Context) AS2AgreementContentOutput {
	return o
}

func (o AS2AgreementContentOutput) ToAS2AgreementContentPtrOutput() AS2AgreementContentPtrOutput {
	return o.ToAS2AgreementContentPtrOutputWithContext(context.Background())
}

func (o AS2AgreementContentOutput) ToAS2AgreementContentPtrOutputWithContext(ctx context.Context) AS2AgreementContentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2AgreementContent) *AS2AgreementContent {
		return &v
	}).(AS2AgreementContentPtrOutput)
}

func (o AS2AgreementContentOutput) ReceiveAgreement() AS2OneWayAgreementOutput {
	return o.ApplyT(func(v AS2AgreementContent) AS2OneWayAgreement { return v.ReceiveAgreement }).(AS2OneWayAgreementOutput)
}

func (o AS2AgreementContentOutput) SendAgreement() AS2OneWayAgreementOutput {
	return o.ApplyT(func(v AS2AgreementContent) AS2OneWayAgreement { return v.SendAgreement }).(AS2OneWayAgreementOutput)
}

type AS2AgreementContentPtrOutput struct{ *pulumi.OutputState }

func (AS2AgreementContentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2AgreementContent)(nil)).Elem()
}

func (o AS2AgreementContentPtrOutput) ToAS2AgreementContentPtrOutput() AS2AgreementContentPtrOutput {
	return o
}

func (o AS2AgreementContentPtrOutput) ToAS2AgreementContentPtrOutputWithContext(ctx context.Context) AS2AgreementContentPtrOutput {
	return o
}

func (o AS2AgreementContentPtrOutput) Elem() AS2AgreementContentOutput {
	return o.ApplyT(func(v *AS2AgreementContent) AS2AgreementContent {
		if v != nil {
			return *v
		}
		var ret AS2AgreementContent
		return ret
	}).(AS2AgreementContentOutput)
}

func (o AS2AgreementContentPtrOutput) ReceiveAgreement() AS2OneWayAgreementPtrOutput {
	return o.ApplyT(func(v *AS2AgreementContent) *AS2OneWayAgreement {
		if v == nil {
			return nil
		}
		return &v.ReceiveAgreement
	}).(AS2OneWayAgreementPtrOutput)
}

func (o AS2AgreementContentPtrOutput) SendAgreement() AS2OneWayAgreementPtrOutput {
	return o.ApplyT(func(v *AS2AgreementContent) *AS2OneWayAgreement {
		if v == nil {
			return nil
		}
		return &v.SendAgreement
	}).(AS2OneWayAgreementPtrOutput)
}

type AS2AgreementContentResponse struct {
	ReceiveAgreement AS2OneWayAgreementResponse `pulumi:"receiveAgreement"`
	SendAgreement    AS2OneWayAgreementResponse `pulumi:"sendAgreement"`
}

type AS2AgreementContentResponseOutput struct{ *pulumi.OutputState }

func (AS2AgreementContentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2AgreementContentResponse)(nil)).Elem()
}

func (o AS2AgreementContentResponseOutput) ToAS2AgreementContentResponseOutput() AS2AgreementContentResponseOutput {
	return o
}

func (o AS2AgreementContentResponseOutput) ToAS2AgreementContentResponseOutputWithContext(ctx context.Context) AS2AgreementContentResponseOutput {
	return o
}

func (o AS2AgreementContentResponseOutput) ReceiveAgreement() AS2OneWayAgreementResponseOutput {
	return o.ApplyT(func(v AS2AgreementContentResponse) AS2OneWayAgreementResponse { return v.ReceiveAgreement }).(AS2OneWayAgreementResponseOutput)
}

func (o AS2AgreementContentResponseOutput) SendAgreement() AS2OneWayAgreementResponseOutput {
	return o.ApplyT(func(v AS2AgreementContentResponse) AS2OneWayAgreementResponse { return v.SendAgreement }).(AS2OneWayAgreementResponseOutput)
}

type AS2AgreementContentResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2AgreementContentResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2AgreementContentResponse)(nil)).Elem()
}

func (o AS2AgreementContentResponsePtrOutput) ToAS2AgreementContentResponsePtrOutput() AS2AgreementContentResponsePtrOutput {
	return o
}

func (o AS2AgreementContentResponsePtrOutput) ToAS2AgreementContentResponsePtrOutputWithContext(ctx context.Context) AS2AgreementContentResponsePtrOutput {
	return o
}

func (o AS2AgreementContentResponsePtrOutput) Elem() AS2AgreementContentResponseOutput {
	return o.ApplyT(func(v *AS2AgreementContentResponse) AS2AgreementContentResponse {
		if v != nil {
			return *v
		}
		var ret AS2AgreementContentResponse
		return ret
	}).(AS2AgreementContentResponseOutput)
}

func (o AS2AgreementContentResponsePtrOutput) ReceiveAgreement() AS2OneWayAgreementResponsePtrOutput {
	return o.ApplyT(func(v *AS2AgreementContentResponse) *AS2OneWayAgreementResponse {
		if v == nil {
			return nil
		}
		return &v.ReceiveAgreement
	}).(AS2OneWayAgreementResponsePtrOutput)
}

func (o AS2AgreementContentResponsePtrOutput) SendAgreement() AS2OneWayAgreementResponsePtrOutput {
	return o.ApplyT(func(v *AS2AgreementContentResponse) *AS2OneWayAgreementResponse {
		if v == nil {
			return nil
		}
		return &v.SendAgreement
	}).(AS2OneWayAgreementResponsePtrOutput)
}

type AS2EnvelopeSettings struct {
	AutogenerateFileName                    bool   `pulumi:"autogenerateFileName"`
	FileNameTemplate                        string `pulumi:"fileNameTemplate"`
	MessageContentType                      string `pulumi:"messageContentType"`
	SuspendMessageOnFileNameGenerationError bool   `pulumi:"suspendMessageOnFileNameGenerationError"`
	TransmitFileNameInMimeHeader            bool   `pulumi:"transmitFileNameInMimeHeader"`
}





type AS2EnvelopeSettingsInput interface {
	pulumi.Input

	ToAS2EnvelopeSettingsOutput() AS2EnvelopeSettingsOutput
	ToAS2EnvelopeSettingsOutputWithContext(context.Context) AS2EnvelopeSettingsOutput
}

type AS2EnvelopeSettingsArgs struct {
	AutogenerateFileName                    pulumi.BoolInput   `pulumi:"autogenerateFileName"`
	FileNameTemplate                        pulumi.StringInput `pulumi:"fileNameTemplate"`
	MessageContentType                      pulumi.StringInput `pulumi:"messageContentType"`
	SuspendMessageOnFileNameGenerationError pulumi.BoolInput   `pulumi:"suspendMessageOnFileNameGenerationError"`
	TransmitFileNameInMimeHeader            pulumi.BoolInput   `pulumi:"transmitFileNameInMimeHeader"`
}

func (AS2EnvelopeSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2EnvelopeSettings)(nil)).Elem()
}

func (i AS2EnvelopeSettingsArgs) ToAS2EnvelopeSettingsOutput() AS2EnvelopeSettingsOutput {
	return i.ToAS2EnvelopeSettingsOutputWithContext(context.Background())
}

func (i AS2EnvelopeSettingsArgs) ToAS2EnvelopeSettingsOutputWithContext(ctx context.Context) AS2EnvelopeSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2EnvelopeSettingsOutput)
}

func (i AS2EnvelopeSettingsArgs) ToAS2EnvelopeSettingsPtrOutput() AS2EnvelopeSettingsPtrOutput {
	return i.ToAS2EnvelopeSettingsPtrOutputWithContext(context.Background())
}

func (i AS2EnvelopeSettingsArgs) ToAS2EnvelopeSettingsPtrOutputWithContext(ctx context.Context) AS2EnvelopeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2EnvelopeSettingsOutput).ToAS2EnvelopeSettingsPtrOutputWithContext(ctx)
}









type AS2EnvelopeSettingsPtrInput interface {
	pulumi.Input

	ToAS2EnvelopeSettingsPtrOutput() AS2EnvelopeSettingsPtrOutput
	ToAS2EnvelopeSettingsPtrOutputWithContext(context.Context) AS2EnvelopeSettingsPtrOutput
}

type as2envelopeSettingsPtrType AS2EnvelopeSettingsArgs

func AS2EnvelopeSettingsPtr(v *AS2EnvelopeSettingsArgs) AS2EnvelopeSettingsPtrInput {
	return (*as2envelopeSettingsPtrType)(v)
}

func (*as2envelopeSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2EnvelopeSettings)(nil)).Elem()
}

func (i *as2envelopeSettingsPtrType) ToAS2EnvelopeSettingsPtrOutput() AS2EnvelopeSettingsPtrOutput {
	return i.ToAS2EnvelopeSettingsPtrOutputWithContext(context.Background())
}

func (i *as2envelopeSettingsPtrType) ToAS2EnvelopeSettingsPtrOutputWithContext(ctx context.Context) AS2EnvelopeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2EnvelopeSettingsPtrOutput)
}

type AS2EnvelopeSettingsOutput struct{ *pulumi.OutputState }

func (AS2EnvelopeSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2EnvelopeSettings)(nil)).Elem()
}

func (o AS2EnvelopeSettingsOutput) ToAS2EnvelopeSettingsOutput() AS2EnvelopeSettingsOutput {
	return o
}

func (o AS2EnvelopeSettingsOutput) ToAS2EnvelopeSettingsOutputWithContext(ctx context.Context) AS2EnvelopeSettingsOutput {
	return o
}

func (o AS2EnvelopeSettingsOutput) ToAS2EnvelopeSettingsPtrOutput() AS2EnvelopeSettingsPtrOutput {
	return o.ToAS2EnvelopeSettingsPtrOutputWithContext(context.Background())
}

func (o AS2EnvelopeSettingsOutput) ToAS2EnvelopeSettingsPtrOutputWithContext(ctx context.Context) AS2EnvelopeSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2EnvelopeSettings) *AS2EnvelopeSettings {
		return &v
	}).(AS2EnvelopeSettingsPtrOutput)
}

func (o AS2EnvelopeSettingsOutput) AutogenerateFileName() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2EnvelopeSettings) bool { return v.AutogenerateFileName }).(pulumi.BoolOutput)
}

func (o AS2EnvelopeSettingsOutput) FileNameTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v AS2EnvelopeSettings) string { return v.FileNameTemplate }).(pulumi.StringOutput)
}

func (o AS2EnvelopeSettingsOutput) MessageContentType() pulumi.StringOutput {
	return o.ApplyT(func(v AS2EnvelopeSettings) string { return v.MessageContentType }).(pulumi.StringOutput)
}

func (o AS2EnvelopeSettingsOutput) SuspendMessageOnFileNameGenerationError() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2EnvelopeSettings) bool { return v.SuspendMessageOnFileNameGenerationError }).(pulumi.BoolOutput)
}

func (o AS2EnvelopeSettingsOutput) TransmitFileNameInMimeHeader() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2EnvelopeSettings) bool { return v.TransmitFileNameInMimeHeader }).(pulumi.BoolOutput)
}

type AS2EnvelopeSettingsPtrOutput struct{ *pulumi.OutputState }

func (AS2EnvelopeSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2EnvelopeSettings)(nil)).Elem()
}

func (o AS2EnvelopeSettingsPtrOutput) ToAS2EnvelopeSettingsPtrOutput() AS2EnvelopeSettingsPtrOutput {
	return o
}

func (o AS2EnvelopeSettingsPtrOutput) ToAS2EnvelopeSettingsPtrOutputWithContext(ctx context.Context) AS2EnvelopeSettingsPtrOutput {
	return o
}

func (o AS2EnvelopeSettingsPtrOutput) Elem() AS2EnvelopeSettingsOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettings) AS2EnvelopeSettings {
		if v != nil {
			return *v
		}
		var ret AS2EnvelopeSettings
		return ret
	}).(AS2EnvelopeSettingsOutput)
}

func (o AS2EnvelopeSettingsPtrOutput) AutogenerateFileName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.AutogenerateFileName
	}).(pulumi.BoolPtrOutput)
}

func (o AS2EnvelopeSettingsPtrOutput) FileNameTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return &v.FileNameTemplate
	}).(pulumi.StringPtrOutput)
}

func (o AS2EnvelopeSettingsPtrOutput) MessageContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return &v.MessageContentType
	}).(pulumi.StringPtrOutput)
}

func (o AS2EnvelopeSettingsPtrOutput) SuspendMessageOnFileNameGenerationError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.SuspendMessageOnFileNameGenerationError
	}).(pulumi.BoolPtrOutput)
}

func (o AS2EnvelopeSettingsPtrOutput) TransmitFileNameInMimeHeader() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.TransmitFileNameInMimeHeader
	}).(pulumi.BoolPtrOutput)
}

type AS2EnvelopeSettingsResponse struct {
	AutogenerateFileName                    bool   `pulumi:"autogenerateFileName"`
	FileNameTemplate                        string `pulumi:"fileNameTemplate"`
	MessageContentType                      string `pulumi:"messageContentType"`
	SuspendMessageOnFileNameGenerationError bool   `pulumi:"suspendMessageOnFileNameGenerationError"`
	TransmitFileNameInMimeHeader            bool   `pulumi:"transmitFileNameInMimeHeader"`
}

type AS2EnvelopeSettingsResponseOutput struct{ *pulumi.OutputState }

func (AS2EnvelopeSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2EnvelopeSettingsResponse)(nil)).Elem()
}

func (o AS2EnvelopeSettingsResponseOutput) ToAS2EnvelopeSettingsResponseOutput() AS2EnvelopeSettingsResponseOutput {
	return o
}

func (o AS2EnvelopeSettingsResponseOutput) ToAS2EnvelopeSettingsResponseOutputWithContext(ctx context.Context) AS2EnvelopeSettingsResponseOutput {
	return o
}

func (o AS2EnvelopeSettingsResponseOutput) AutogenerateFileName() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2EnvelopeSettingsResponse) bool { return v.AutogenerateFileName }).(pulumi.BoolOutput)
}

func (o AS2EnvelopeSettingsResponseOutput) FileNameTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v AS2EnvelopeSettingsResponse) string { return v.FileNameTemplate }).(pulumi.StringOutput)
}

func (o AS2EnvelopeSettingsResponseOutput) MessageContentType() pulumi.StringOutput {
	return o.ApplyT(func(v AS2EnvelopeSettingsResponse) string { return v.MessageContentType }).(pulumi.StringOutput)
}

func (o AS2EnvelopeSettingsResponseOutput) SuspendMessageOnFileNameGenerationError() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2EnvelopeSettingsResponse) bool { return v.SuspendMessageOnFileNameGenerationError }).(pulumi.BoolOutput)
}

func (o AS2EnvelopeSettingsResponseOutput) TransmitFileNameInMimeHeader() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2EnvelopeSettingsResponse) bool { return v.TransmitFileNameInMimeHeader }).(pulumi.BoolOutput)
}

type AS2EnvelopeSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2EnvelopeSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2EnvelopeSettingsResponse)(nil)).Elem()
}

func (o AS2EnvelopeSettingsResponsePtrOutput) ToAS2EnvelopeSettingsResponsePtrOutput() AS2EnvelopeSettingsResponsePtrOutput {
	return o
}

func (o AS2EnvelopeSettingsResponsePtrOutput) ToAS2EnvelopeSettingsResponsePtrOutputWithContext(ctx context.Context) AS2EnvelopeSettingsResponsePtrOutput {
	return o
}

func (o AS2EnvelopeSettingsResponsePtrOutput) Elem() AS2EnvelopeSettingsResponseOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettingsResponse) AS2EnvelopeSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AS2EnvelopeSettingsResponse
		return ret
	}).(AS2EnvelopeSettingsResponseOutput)
}

func (o AS2EnvelopeSettingsResponsePtrOutput) AutogenerateFileName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.AutogenerateFileName
	}).(pulumi.BoolPtrOutput)
}

func (o AS2EnvelopeSettingsResponsePtrOutput) FileNameTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FileNameTemplate
	}).(pulumi.StringPtrOutput)
}

func (o AS2EnvelopeSettingsResponsePtrOutput) MessageContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.MessageContentType
	}).(pulumi.StringPtrOutput)
}

func (o AS2EnvelopeSettingsResponsePtrOutput) SuspendMessageOnFileNameGenerationError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.SuspendMessageOnFileNameGenerationError
	}).(pulumi.BoolPtrOutput)
}

func (o AS2EnvelopeSettingsResponsePtrOutput) TransmitFileNameInMimeHeader() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2EnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.TransmitFileNameInMimeHeader
	}).(pulumi.BoolPtrOutput)
}

type AS2ErrorSettings struct {
	ResendIfMDNNotReceived  bool `pulumi:"resendIfMDNNotReceived"`
	SuspendDuplicateMessage bool `pulumi:"suspendDuplicateMessage"`
}





type AS2ErrorSettingsInput interface {
	pulumi.Input

	ToAS2ErrorSettingsOutput() AS2ErrorSettingsOutput
	ToAS2ErrorSettingsOutputWithContext(context.Context) AS2ErrorSettingsOutput
}

type AS2ErrorSettingsArgs struct {
	ResendIfMDNNotReceived  pulumi.BoolInput `pulumi:"resendIfMDNNotReceived"`
	SuspendDuplicateMessage pulumi.BoolInput `pulumi:"suspendDuplicateMessage"`
}

func (AS2ErrorSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2ErrorSettings)(nil)).Elem()
}

func (i AS2ErrorSettingsArgs) ToAS2ErrorSettingsOutput() AS2ErrorSettingsOutput {
	return i.ToAS2ErrorSettingsOutputWithContext(context.Background())
}

func (i AS2ErrorSettingsArgs) ToAS2ErrorSettingsOutputWithContext(ctx context.Context) AS2ErrorSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2ErrorSettingsOutput)
}

func (i AS2ErrorSettingsArgs) ToAS2ErrorSettingsPtrOutput() AS2ErrorSettingsPtrOutput {
	return i.ToAS2ErrorSettingsPtrOutputWithContext(context.Background())
}

func (i AS2ErrorSettingsArgs) ToAS2ErrorSettingsPtrOutputWithContext(ctx context.Context) AS2ErrorSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2ErrorSettingsOutput).ToAS2ErrorSettingsPtrOutputWithContext(ctx)
}









type AS2ErrorSettingsPtrInput interface {
	pulumi.Input

	ToAS2ErrorSettingsPtrOutput() AS2ErrorSettingsPtrOutput
	ToAS2ErrorSettingsPtrOutputWithContext(context.Context) AS2ErrorSettingsPtrOutput
}

type as2errorSettingsPtrType AS2ErrorSettingsArgs

func AS2ErrorSettingsPtr(v *AS2ErrorSettingsArgs) AS2ErrorSettingsPtrInput {
	return (*as2errorSettingsPtrType)(v)
}

func (*as2errorSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2ErrorSettings)(nil)).Elem()
}

func (i *as2errorSettingsPtrType) ToAS2ErrorSettingsPtrOutput() AS2ErrorSettingsPtrOutput {
	return i.ToAS2ErrorSettingsPtrOutputWithContext(context.Background())
}

func (i *as2errorSettingsPtrType) ToAS2ErrorSettingsPtrOutputWithContext(ctx context.Context) AS2ErrorSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2ErrorSettingsPtrOutput)
}

type AS2ErrorSettingsOutput struct{ *pulumi.OutputState }

func (AS2ErrorSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2ErrorSettings)(nil)).Elem()
}

func (o AS2ErrorSettingsOutput) ToAS2ErrorSettingsOutput() AS2ErrorSettingsOutput {
	return o
}

func (o AS2ErrorSettingsOutput) ToAS2ErrorSettingsOutputWithContext(ctx context.Context) AS2ErrorSettingsOutput {
	return o
}

func (o AS2ErrorSettingsOutput) ToAS2ErrorSettingsPtrOutput() AS2ErrorSettingsPtrOutput {
	return o.ToAS2ErrorSettingsPtrOutputWithContext(context.Background())
}

func (o AS2ErrorSettingsOutput) ToAS2ErrorSettingsPtrOutputWithContext(ctx context.Context) AS2ErrorSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2ErrorSettings) *AS2ErrorSettings {
		return &v
	}).(AS2ErrorSettingsPtrOutput)
}

func (o AS2ErrorSettingsOutput) ResendIfMDNNotReceived() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2ErrorSettings) bool { return v.ResendIfMDNNotReceived }).(pulumi.BoolOutput)
}

func (o AS2ErrorSettingsOutput) SuspendDuplicateMessage() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2ErrorSettings) bool { return v.SuspendDuplicateMessage }).(pulumi.BoolOutput)
}

type AS2ErrorSettingsPtrOutput struct{ *pulumi.OutputState }

func (AS2ErrorSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2ErrorSettings)(nil)).Elem()
}

func (o AS2ErrorSettingsPtrOutput) ToAS2ErrorSettingsPtrOutput() AS2ErrorSettingsPtrOutput {
	return o
}

func (o AS2ErrorSettingsPtrOutput) ToAS2ErrorSettingsPtrOutputWithContext(ctx context.Context) AS2ErrorSettingsPtrOutput {
	return o
}

func (o AS2ErrorSettingsPtrOutput) Elem() AS2ErrorSettingsOutput {
	return o.ApplyT(func(v *AS2ErrorSettings) AS2ErrorSettings {
		if v != nil {
			return *v
		}
		var ret AS2ErrorSettings
		return ret
	}).(AS2ErrorSettingsOutput)
}

func (o AS2ErrorSettingsPtrOutput) ResendIfMDNNotReceived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ErrorSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.ResendIfMDNNotReceived
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ErrorSettingsPtrOutput) SuspendDuplicateMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ErrorSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.SuspendDuplicateMessage
	}).(pulumi.BoolPtrOutput)
}

type AS2ErrorSettingsResponse struct {
	ResendIfMDNNotReceived  bool `pulumi:"resendIfMDNNotReceived"`
	SuspendDuplicateMessage bool `pulumi:"suspendDuplicateMessage"`
}

type AS2ErrorSettingsResponseOutput struct{ *pulumi.OutputState }

func (AS2ErrorSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2ErrorSettingsResponse)(nil)).Elem()
}

func (o AS2ErrorSettingsResponseOutput) ToAS2ErrorSettingsResponseOutput() AS2ErrorSettingsResponseOutput {
	return o
}

func (o AS2ErrorSettingsResponseOutput) ToAS2ErrorSettingsResponseOutputWithContext(ctx context.Context) AS2ErrorSettingsResponseOutput {
	return o
}

func (o AS2ErrorSettingsResponseOutput) ResendIfMDNNotReceived() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2ErrorSettingsResponse) bool { return v.ResendIfMDNNotReceived }).(pulumi.BoolOutput)
}

func (o AS2ErrorSettingsResponseOutput) SuspendDuplicateMessage() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2ErrorSettingsResponse) bool { return v.SuspendDuplicateMessage }).(pulumi.BoolOutput)
}

type AS2ErrorSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2ErrorSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2ErrorSettingsResponse)(nil)).Elem()
}

func (o AS2ErrorSettingsResponsePtrOutput) ToAS2ErrorSettingsResponsePtrOutput() AS2ErrorSettingsResponsePtrOutput {
	return o
}

func (o AS2ErrorSettingsResponsePtrOutput) ToAS2ErrorSettingsResponsePtrOutputWithContext(ctx context.Context) AS2ErrorSettingsResponsePtrOutput {
	return o
}

func (o AS2ErrorSettingsResponsePtrOutput) Elem() AS2ErrorSettingsResponseOutput {
	return o.ApplyT(func(v *AS2ErrorSettingsResponse) AS2ErrorSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AS2ErrorSettingsResponse
		return ret
	}).(AS2ErrorSettingsResponseOutput)
}

func (o AS2ErrorSettingsResponsePtrOutput) ResendIfMDNNotReceived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ErrorSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.ResendIfMDNNotReceived
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ErrorSettingsResponsePtrOutput) SuspendDuplicateMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ErrorSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.SuspendDuplicateMessage
	}).(pulumi.BoolPtrOutput)
}

type AS2MdnSettings struct {
	DispositionNotificationTo  *string `pulumi:"dispositionNotificationTo"`
	MdnText                    *string `pulumi:"mdnText"`
	MicHashingAlgorithm        string  `pulumi:"micHashingAlgorithm"`
	NeedMDN                    bool    `pulumi:"needMDN"`
	ReceiptDeliveryUrl         *string `pulumi:"receiptDeliveryUrl"`
	SendInboundMDNToMessageBox bool    `pulumi:"sendInboundMDNToMessageBox"`
	SendMDNAsynchronously      bool    `pulumi:"sendMDNAsynchronously"`
	SignMDN                    bool    `pulumi:"signMDN"`
	SignOutboundMDNIfOptional  bool    `pulumi:"signOutboundMDNIfOptional"`
}





type AS2MdnSettingsInput interface {
	pulumi.Input

	ToAS2MdnSettingsOutput() AS2MdnSettingsOutput
	ToAS2MdnSettingsOutputWithContext(context.Context) AS2MdnSettingsOutput
}

type AS2MdnSettingsArgs struct {
	DispositionNotificationTo  pulumi.StringPtrInput `pulumi:"dispositionNotificationTo"`
	MdnText                    pulumi.StringPtrInput `pulumi:"mdnText"`
	MicHashingAlgorithm        pulumi.StringInput    `pulumi:"micHashingAlgorithm"`
	NeedMDN                    pulumi.BoolInput      `pulumi:"needMDN"`
	ReceiptDeliveryUrl         pulumi.StringPtrInput `pulumi:"receiptDeliveryUrl"`
	SendInboundMDNToMessageBox pulumi.BoolInput      `pulumi:"sendInboundMDNToMessageBox"`
	SendMDNAsynchronously      pulumi.BoolInput      `pulumi:"sendMDNAsynchronously"`
	SignMDN                    pulumi.BoolInput      `pulumi:"signMDN"`
	SignOutboundMDNIfOptional  pulumi.BoolInput      `pulumi:"signOutboundMDNIfOptional"`
}

func (AS2MdnSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2MdnSettings)(nil)).Elem()
}

func (i AS2MdnSettingsArgs) ToAS2MdnSettingsOutput() AS2MdnSettingsOutput {
	return i.ToAS2MdnSettingsOutputWithContext(context.Background())
}

func (i AS2MdnSettingsArgs) ToAS2MdnSettingsOutputWithContext(ctx context.Context) AS2MdnSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2MdnSettingsOutput)
}

func (i AS2MdnSettingsArgs) ToAS2MdnSettingsPtrOutput() AS2MdnSettingsPtrOutput {
	return i.ToAS2MdnSettingsPtrOutputWithContext(context.Background())
}

func (i AS2MdnSettingsArgs) ToAS2MdnSettingsPtrOutputWithContext(ctx context.Context) AS2MdnSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2MdnSettingsOutput).ToAS2MdnSettingsPtrOutputWithContext(ctx)
}









type AS2MdnSettingsPtrInput interface {
	pulumi.Input

	ToAS2MdnSettingsPtrOutput() AS2MdnSettingsPtrOutput
	ToAS2MdnSettingsPtrOutputWithContext(context.Context) AS2MdnSettingsPtrOutput
}

type as2mdnSettingsPtrType AS2MdnSettingsArgs

func AS2MdnSettingsPtr(v *AS2MdnSettingsArgs) AS2MdnSettingsPtrInput {
	return (*as2mdnSettingsPtrType)(v)
}

func (*as2mdnSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2MdnSettings)(nil)).Elem()
}

func (i *as2mdnSettingsPtrType) ToAS2MdnSettingsPtrOutput() AS2MdnSettingsPtrOutput {
	return i.ToAS2MdnSettingsPtrOutputWithContext(context.Background())
}

func (i *as2mdnSettingsPtrType) ToAS2MdnSettingsPtrOutputWithContext(ctx context.Context) AS2MdnSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2MdnSettingsPtrOutput)
}

type AS2MdnSettingsOutput struct{ *pulumi.OutputState }

func (AS2MdnSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2MdnSettings)(nil)).Elem()
}

func (o AS2MdnSettingsOutput) ToAS2MdnSettingsOutput() AS2MdnSettingsOutput {
	return o
}

func (o AS2MdnSettingsOutput) ToAS2MdnSettingsOutputWithContext(ctx context.Context) AS2MdnSettingsOutput {
	return o
}

func (o AS2MdnSettingsOutput) ToAS2MdnSettingsPtrOutput() AS2MdnSettingsPtrOutput {
	return o.ToAS2MdnSettingsPtrOutputWithContext(context.Background())
}

func (o AS2MdnSettingsOutput) ToAS2MdnSettingsPtrOutputWithContext(ctx context.Context) AS2MdnSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2MdnSettings) *AS2MdnSettings {
		return &v
	}).(AS2MdnSettingsPtrOutput)
}

func (o AS2MdnSettingsOutput) DispositionNotificationTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2MdnSettings) *string { return v.DispositionNotificationTo }).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsOutput) MdnText() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2MdnSettings) *string { return v.MdnText }).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsOutput) MicHashingAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v AS2MdnSettings) string { return v.MicHashingAlgorithm }).(pulumi.StringOutput)
}

func (o AS2MdnSettingsOutput) NeedMDN() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2MdnSettings) bool { return v.NeedMDN }).(pulumi.BoolOutput)
}

func (o AS2MdnSettingsOutput) ReceiptDeliveryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2MdnSettings) *string { return v.ReceiptDeliveryUrl }).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsOutput) SendInboundMDNToMessageBox() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2MdnSettings) bool { return v.SendInboundMDNToMessageBox }).(pulumi.BoolOutput)
}

func (o AS2MdnSettingsOutput) SendMDNAsynchronously() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2MdnSettings) bool { return v.SendMDNAsynchronously }).(pulumi.BoolOutput)
}

func (o AS2MdnSettingsOutput) SignMDN() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2MdnSettings) bool { return v.SignMDN }).(pulumi.BoolOutput)
}

func (o AS2MdnSettingsOutput) SignOutboundMDNIfOptional() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2MdnSettings) bool { return v.SignOutboundMDNIfOptional }).(pulumi.BoolOutput)
}

type AS2MdnSettingsPtrOutput struct{ *pulumi.OutputState }

func (AS2MdnSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2MdnSettings)(nil)).Elem()
}

func (o AS2MdnSettingsPtrOutput) ToAS2MdnSettingsPtrOutput() AS2MdnSettingsPtrOutput {
	return o
}

func (o AS2MdnSettingsPtrOutput) ToAS2MdnSettingsPtrOutputWithContext(ctx context.Context) AS2MdnSettingsPtrOutput {
	return o
}

func (o AS2MdnSettingsPtrOutput) Elem() AS2MdnSettingsOutput {
	return o.ApplyT(func(v *AS2MdnSettings) AS2MdnSettings {
		if v != nil {
			return *v
		}
		var ret AS2MdnSettings
		return ret
	}).(AS2MdnSettingsOutput)
}

func (o AS2MdnSettingsPtrOutput) DispositionNotificationTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettings) *string {
		if v == nil {
			return nil
		}
		return v.DispositionNotificationTo
	}).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsPtrOutput) MdnText() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettings) *string {
		if v == nil {
			return nil
		}
		return v.MdnText
	}).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsPtrOutput) MicHashingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettings) *string {
		if v == nil {
			return nil
		}
		return &v.MicHashingAlgorithm
	}).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsPtrOutput) NeedMDN() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.NeedMDN
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsPtrOutput) ReceiptDeliveryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettings) *string {
		if v == nil {
			return nil
		}
		return v.ReceiptDeliveryUrl
	}).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsPtrOutput) SendInboundMDNToMessageBox() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.SendInboundMDNToMessageBox
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsPtrOutput) SendMDNAsynchronously() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.SendMDNAsynchronously
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsPtrOutput) SignMDN() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.SignMDN
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsPtrOutput) SignOutboundMDNIfOptional() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.SignOutboundMDNIfOptional
	}).(pulumi.BoolPtrOutput)
}

type AS2MdnSettingsResponse struct {
	DispositionNotificationTo  *string `pulumi:"dispositionNotificationTo"`
	MdnText                    *string `pulumi:"mdnText"`
	MicHashingAlgorithm        string  `pulumi:"micHashingAlgorithm"`
	NeedMDN                    bool    `pulumi:"needMDN"`
	ReceiptDeliveryUrl         *string `pulumi:"receiptDeliveryUrl"`
	SendInboundMDNToMessageBox bool    `pulumi:"sendInboundMDNToMessageBox"`
	SendMDNAsynchronously      bool    `pulumi:"sendMDNAsynchronously"`
	SignMDN                    bool    `pulumi:"signMDN"`
	SignOutboundMDNIfOptional  bool    `pulumi:"signOutboundMDNIfOptional"`
}

type AS2MdnSettingsResponseOutput struct{ *pulumi.OutputState }

func (AS2MdnSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2MdnSettingsResponse)(nil)).Elem()
}

func (o AS2MdnSettingsResponseOutput) ToAS2MdnSettingsResponseOutput() AS2MdnSettingsResponseOutput {
	return o
}

func (o AS2MdnSettingsResponseOutput) ToAS2MdnSettingsResponseOutputWithContext(ctx context.Context) AS2MdnSettingsResponseOutput {
	return o
}

func (o AS2MdnSettingsResponseOutput) DispositionNotificationTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2MdnSettingsResponse) *string { return v.DispositionNotificationTo }).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsResponseOutput) MdnText() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2MdnSettingsResponse) *string { return v.MdnText }).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsResponseOutput) MicHashingAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v AS2MdnSettingsResponse) string { return v.MicHashingAlgorithm }).(pulumi.StringOutput)
}

func (o AS2MdnSettingsResponseOutput) NeedMDN() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2MdnSettingsResponse) bool { return v.NeedMDN }).(pulumi.BoolOutput)
}

func (o AS2MdnSettingsResponseOutput) ReceiptDeliveryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2MdnSettingsResponse) *string { return v.ReceiptDeliveryUrl }).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsResponseOutput) SendInboundMDNToMessageBox() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2MdnSettingsResponse) bool { return v.SendInboundMDNToMessageBox }).(pulumi.BoolOutput)
}

func (o AS2MdnSettingsResponseOutput) SendMDNAsynchronously() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2MdnSettingsResponse) bool { return v.SendMDNAsynchronously }).(pulumi.BoolOutput)
}

func (o AS2MdnSettingsResponseOutput) SignMDN() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2MdnSettingsResponse) bool { return v.SignMDN }).(pulumi.BoolOutput)
}

func (o AS2MdnSettingsResponseOutput) SignOutboundMDNIfOptional() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2MdnSettingsResponse) bool { return v.SignOutboundMDNIfOptional }).(pulumi.BoolOutput)
}

type AS2MdnSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2MdnSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2MdnSettingsResponse)(nil)).Elem()
}

func (o AS2MdnSettingsResponsePtrOutput) ToAS2MdnSettingsResponsePtrOutput() AS2MdnSettingsResponsePtrOutput {
	return o
}

func (o AS2MdnSettingsResponsePtrOutput) ToAS2MdnSettingsResponsePtrOutputWithContext(ctx context.Context) AS2MdnSettingsResponsePtrOutput {
	return o
}

func (o AS2MdnSettingsResponsePtrOutput) Elem() AS2MdnSettingsResponseOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) AS2MdnSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AS2MdnSettingsResponse
		return ret
	}).(AS2MdnSettingsResponseOutput)
}

func (o AS2MdnSettingsResponsePtrOutput) DispositionNotificationTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DispositionNotificationTo
	}).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsResponsePtrOutput) MdnText() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.MdnText
	}).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsResponsePtrOutput) MicHashingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.MicHashingAlgorithm
	}).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsResponsePtrOutput) NeedMDN() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.NeedMDN
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsResponsePtrOutput) ReceiptDeliveryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReceiptDeliveryUrl
	}).(pulumi.StringPtrOutput)
}

func (o AS2MdnSettingsResponsePtrOutput) SendInboundMDNToMessageBox() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.SendInboundMDNToMessageBox
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsResponsePtrOutput) SendMDNAsynchronously() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.SendMDNAsynchronously
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsResponsePtrOutput) SignMDN() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.SignMDN
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MdnSettingsResponsePtrOutput) SignOutboundMDNIfOptional() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MdnSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.SignOutboundMDNIfOptional
	}).(pulumi.BoolPtrOutput)
}

type AS2MessageConnectionSettings struct {
	IgnoreCertificateNameMismatch bool `pulumi:"ignoreCertificateNameMismatch"`
	KeepHttpConnectionAlive       bool `pulumi:"keepHttpConnectionAlive"`
	SupportHttpStatusCodeContinue bool `pulumi:"supportHttpStatusCodeContinue"`
	UnfoldHttpHeaders             bool `pulumi:"unfoldHttpHeaders"`
}





type AS2MessageConnectionSettingsInput interface {
	pulumi.Input

	ToAS2MessageConnectionSettingsOutput() AS2MessageConnectionSettingsOutput
	ToAS2MessageConnectionSettingsOutputWithContext(context.Context) AS2MessageConnectionSettingsOutput
}

type AS2MessageConnectionSettingsArgs struct {
	IgnoreCertificateNameMismatch pulumi.BoolInput `pulumi:"ignoreCertificateNameMismatch"`
	KeepHttpConnectionAlive       pulumi.BoolInput `pulumi:"keepHttpConnectionAlive"`
	SupportHttpStatusCodeContinue pulumi.BoolInput `pulumi:"supportHttpStatusCodeContinue"`
	UnfoldHttpHeaders             pulumi.BoolInput `pulumi:"unfoldHttpHeaders"`
}

func (AS2MessageConnectionSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2MessageConnectionSettings)(nil)).Elem()
}

func (i AS2MessageConnectionSettingsArgs) ToAS2MessageConnectionSettingsOutput() AS2MessageConnectionSettingsOutput {
	return i.ToAS2MessageConnectionSettingsOutputWithContext(context.Background())
}

func (i AS2MessageConnectionSettingsArgs) ToAS2MessageConnectionSettingsOutputWithContext(ctx context.Context) AS2MessageConnectionSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2MessageConnectionSettingsOutput)
}

func (i AS2MessageConnectionSettingsArgs) ToAS2MessageConnectionSettingsPtrOutput() AS2MessageConnectionSettingsPtrOutput {
	return i.ToAS2MessageConnectionSettingsPtrOutputWithContext(context.Background())
}

func (i AS2MessageConnectionSettingsArgs) ToAS2MessageConnectionSettingsPtrOutputWithContext(ctx context.Context) AS2MessageConnectionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2MessageConnectionSettingsOutput).ToAS2MessageConnectionSettingsPtrOutputWithContext(ctx)
}









type AS2MessageConnectionSettingsPtrInput interface {
	pulumi.Input

	ToAS2MessageConnectionSettingsPtrOutput() AS2MessageConnectionSettingsPtrOutput
	ToAS2MessageConnectionSettingsPtrOutputWithContext(context.Context) AS2MessageConnectionSettingsPtrOutput
}

type as2messageConnectionSettingsPtrType AS2MessageConnectionSettingsArgs

func AS2MessageConnectionSettingsPtr(v *AS2MessageConnectionSettingsArgs) AS2MessageConnectionSettingsPtrInput {
	return (*as2messageConnectionSettingsPtrType)(v)
}

func (*as2messageConnectionSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2MessageConnectionSettings)(nil)).Elem()
}

func (i *as2messageConnectionSettingsPtrType) ToAS2MessageConnectionSettingsPtrOutput() AS2MessageConnectionSettingsPtrOutput {
	return i.ToAS2MessageConnectionSettingsPtrOutputWithContext(context.Background())
}

func (i *as2messageConnectionSettingsPtrType) ToAS2MessageConnectionSettingsPtrOutputWithContext(ctx context.Context) AS2MessageConnectionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2MessageConnectionSettingsPtrOutput)
}

type AS2MessageConnectionSettingsOutput struct{ *pulumi.OutputState }

func (AS2MessageConnectionSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2MessageConnectionSettings)(nil)).Elem()
}

func (o AS2MessageConnectionSettingsOutput) ToAS2MessageConnectionSettingsOutput() AS2MessageConnectionSettingsOutput {
	return o
}

func (o AS2MessageConnectionSettingsOutput) ToAS2MessageConnectionSettingsOutputWithContext(ctx context.Context) AS2MessageConnectionSettingsOutput {
	return o
}

func (o AS2MessageConnectionSettingsOutput) ToAS2MessageConnectionSettingsPtrOutput() AS2MessageConnectionSettingsPtrOutput {
	return o.ToAS2MessageConnectionSettingsPtrOutputWithContext(context.Background())
}

func (o AS2MessageConnectionSettingsOutput) ToAS2MessageConnectionSettingsPtrOutputWithContext(ctx context.Context) AS2MessageConnectionSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2MessageConnectionSettings) *AS2MessageConnectionSettings {
		return &v
	}).(AS2MessageConnectionSettingsPtrOutput)
}

func (o AS2MessageConnectionSettingsOutput) IgnoreCertificateNameMismatch() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2MessageConnectionSettings) bool { return v.IgnoreCertificateNameMismatch }).(pulumi.BoolOutput)
}

func (o AS2MessageConnectionSettingsOutput) KeepHttpConnectionAlive() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2MessageConnectionSettings) bool { return v.KeepHttpConnectionAlive }).(pulumi.BoolOutput)
}

func (o AS2MessageConnectionSettingsOutput) SupportHttpStatusCodeContinue() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2MessageConnectionSettings) bool { return v.SupportHttpStatusCodeContinue }).(pulumi.BoolOutput)
}

func (o AS2MessageConnectionSettingsOutput) UnfoldHttpHeaders() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2MessageConnectionSettings) bool { return v.UnfoldHttpHeaders }).(pulumi.BoolOutput)
}

type AS2MessageConnectionSettingsPtrOutput struct{ *pulumi.OutputState }

func (AS2MessageConnectionSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2MessageConnectionSettings)(nil)).Elem()
}

func (o AS2MessageConnectionSettingsPtrOutput) ToAS2MessageConnectionSettingsPtrOutput() AS2MessageConnectionSettingsPtrOutput {
	return o
}

func (o AS2MessageConnectionSettingsPtrOutput) ToAS2MessageConnectionSettingsPtrOutputWithContext(ctx context.Context) AS2MessageConnectionSettingsPtrOutput {
	return o
}

func (o AS2MessageConnectionSettingsPtrOutput) Elem() AS2MessageConnectionSettingsOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettings) AS2MessageConnectionSettings {
		if v != nil {
			return *v
		}
		var ret AS2MessageConnectionSettings
		return ret
	}).(AS2MessageConnectionSettingsOutput)
}

func (o AS2MessageConnectionSettingsPtrOutput) IgnoreCertificateNameMismatch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.IgnoreCertificateNameMismatch
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MessageConnectionSettingsPtrOutput) KeepHttpConnectionAlive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.KeepHttpConnectionAlive
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MessageConnectionSettingsPtrOutput) SupportHttpStatusCodeContinue() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.SupportHttpStatusCodeContinue
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MessageConnectionSettingsPtrOutput) UnfoldHttpHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.UnfoldHttpHeaders
	}).(pulumi.BoolPtrOutput)
}

type AS2MessageConnectionSettingsResponse struct {
	IgnoreCertificateNameMismatch bool `pulumi:"ignoreCertificateNameMismatch"`
	KeepHttpConnectionAlive       bool `pulumi:"keepHttpConnectionAlive"`
	SupportHttpStatusCodeContinue bool `pulumi:"supportHttpStatusCodeContinue"`
	UnfoldHttpHeaders             bool `pulumi:"unfoldHttpHeaders"`
}

type AS2MessageConnectionSettingsResponseOutput struct{ *pulumi.OutputState }

func (AS2MessageConnectionSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2MessageConnectionSettingsResponse)(nil)).Elem()
}

func (o AS2MessageConnectionSettingsResponseOutput) ToAS2MessageConnectionSettingsResponseOutput() AS2MessageConnectionSettingsResponseOutput {
	return o
}

func (o AS2MessageConnectionSettingsResponseOutput) ToAS2MessageConnectionSettingsResponseOutputWithContext(ctx context.Context) AS2MessageConnectionSettingsResponseOutput {
	return o
}

func (o AS2MessageConnectionSettingsResponseOutput) IgnoreCertificateNameMismatch() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2MessageConnectionSettingsResponse) bool { return v.IgnoreCertificateNameMismatch }).(pulumi.BoolOutput)
}

func (o AS2MessageConnectionSettingsResponseOutput) KeepHttpConnectionAlive() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2MessageConnectionSettingsResponse) bool { return v.KeepHttpConnectionAlive }).(pulumi.BoolOutput)
}

func (o AS2MessageConnectionSettingsResponseOutput) SupportHttpStatusCodeContinue() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2MessageConnectionSettingsResponse) bool { return v.SupportHttpStatusCodeContinue }).(pulumi.BoolOutput)
}

func (o AS2MessageConnectionSettingsResponseOutput) UnfoldHttpHeaders() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2MessageConnectionSettingsResponse) bool { return v.UnfoldHttpHeaders }).(pulumi.BoolOutput)
}

type AS2MessageConnectionSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2MessageConnectionSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2MessageConnectionSettingsResponse)(nil)).Elem()
}

func (o AS2MessageConnectionSettingsResponsePtrOutput) ToAS2MessageConnectionSettingsResponsePtrOutput() AS2MessageConnectionSettingsResponsePtrOutput {
	return o
}

func (o AS2MessageConnectionSettingsResponsePtrOutput) ToAS2MessageConnectionSettingsResponsePtrOutputWithContext(ctx context.Context) AS2MessageConnectionSettingsResponsePtrOutput {
	return o
}

func (o AS2MessageConnectionSettingsResponsePtrOutput) Elem() AS2MessageConnectionSettingsResponseOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettingsResponse) AS2MessageConnectionSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AS2MessageConnectionSettingsResponse
		return ret
	}).(AS2MessageConnectionSettingsResponseOutput)
}

func (o AS2MessageConnectionSettingsResponsePtrOutput) IgnoreCertificateNameMismatch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IgnoreCertificateNameMismatch
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MessageConnectionSettingsResponsePtrOutput) KeepHttpConnectionAlive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.KeepHttpConnectionAlive
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MessageConnectionSettingsResponsePtrOutput) SupportHttpStatusCodeContinue() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.SupportHttpStatusCodeContinue
	}).(pulumi.BoolPtrOutput)
}

func (o AS2MessageConnectionSettingsResponsePtrOutput) UnfoldHttpHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2MessageConnectionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.UnfoldHttpHeaders
	}).(pulumi.BoolPtrOutput)
}

type AS2OneWayAgreement struct {
	ProtocolSettings         AS2ProtocolSettings `pulumi:"protocolSettings"`
	ReceiverBusinessIdentity BusinessIdentity    `pulumi:"receiverBusinessIdentity"`
	SenderBusinessIdentity   BusinessIdentity    `pulumi:"senderBusinessIdentity"`
}





type AS2OneWayAgreementInput interface {
	pulumi.Input

	ToAS2OneWayAgreementOutput() AS2OneWayAgreementOutput
	ToAS2OneWayAgreementOutputWithContext(context.Context) AS2OneWayAgreementOutput
}

type AS2OneWayAgreementArgs struct {
	ProtocolSettings         AS2ProtocolSettingsInput `pulumi:"protocolSettings"`
	ReceiverBusinessIdentity BusinessIdentityInput    `pulumi:"receiverBusinessIdentity"`
	SenderBusinessIdentity   BusinessIdentityInput    `pulumi:"senderBusinessIdentity"`
}

func (AS2OneWayAgreementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2OneWayAgreement)(nil)).Elem()
}

func (i AS2OneWayAgreementArgs) ToAS2OneWayAgreementOutput() AS2OneWayAgreementOutput {
	return i.ToAS2OneWayAgreementOutputWithContext(context.Background())
}

func (i AS2OneWayAgreementArgs) ToAS2OneWayAgreementOutputWithContext(ctx context.Context) AS2OneWayAgreementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2OneWayAgreementOutput)
}

func (i AS2OneWayAgreementArgs) ToAS2OneWayAgreementPtrOutput() AS2OneWayAgreementPtrOutput {
	return i.ToAS2OneWayAgreementPtrOutputWithContext(context.Background())
}

func (i AS2OneWayAgreementArgs) ToAS2OneWayAgreementPtrOutputWithContext(ctx context.Context) AS2OneWayAgreementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2OneWayAgreementOutput).ToAS2OneWayAgreementPtrOutputWithContext(ctx)
}









type AS2OneWayAgreementPtrInput interface {
	pulumi.Input

	ToAS2OneWayAgreementPtrOutput() AS2OneWayAgreementPtrOutput
	ToAS2OneWayAgreementPtrOutputWithContext(context.Context) AS2OneWayAgreementPtrOutput
}

type as2oneWayAgreementPtrType AS2OneWayAgreementArgs

func AS2OneWayAgreementPtr(v *AS2OneWayAgreementArgs) AS2OneWayAgreementPtrInput {
	return (*as2oneWayAgreementPtrType)(v)
}

func (*as2oneWayAgreementPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2OneWayAgreement)(nil)).Elem()
}

func (i *as2oneWayAgreementPtrType) ToAS2OneWayAgreementPtrOutput() AS2OneWayAgreementPtrOutput {
	return i.ToAS2OneWayAgreementPtrOutputWithContext(context.Background())
}

func (i *as2oneWayAgreementPtrType) ToAS2OneWayAgreementPtrOutputWithContext(ctx context.Context) AS2OneWayAgreementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2OneWayAgreementPtrOutput)
}

type AS2OneWayAgreementOutput struct{ *pulumi.OutputState }

func (AS2OneWayAgreementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2OneWayAgreement)(nil)).Elem()
}

func (o AS2OneWayAgreementOutput) ToAS2OneWayAgreementOutput() AS2OneWayAgreementOutput {
	return o
}

func (o AS2OneWayAgreementOutput) ToAS2OneWayAgreementOutputWithContext(ctx context.Context) AS2OneWayAgreementOutput {
	return o
}

func (o AS2OneWayAgreementOutput) ToAS2OneWayAgreementPtrOutput() AS2OneWayAgreementPtrOutput {
	return o.ToAS2OneWayAgreementPtrOutputWithContext(context.Background())
}

func (o AS2OneWayAgreementOutput) ToAS2OneWayAgreementPtrOutputWithContext(ctx context.Context) AS2OneWayAgreementPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2OneWayAgreement) *AS2OneWayAgreement {
		return &v
	}).(AS2OneWayAgreementPtrOutput)
}

func (o AS2OneWayAgreementOutput) ProtocolSettings() AS2ProtocolSettingsOutput {
	return o.ApplyT(func(v AS2OneWayAgreement) AS2ProtocolSettings { return v.ProtocolSettings }).(AS2ProtocolSettingsOutput)
}

func (o AS2OneWayAgreementOutput) ReceiverBusinessIdentity() BusinessIdentityOutput {
	return o.ApplyT(func(v AS2OneWayAgreement) BusinessIdentity { return v.ReceiverBusinessIdentity }).(BusinessIdentityOutput)
}

func (o AS2OneWayAgreementOutput) SenderBusinessIdentity() BusinessIdentityOutput {
	return o.ApplyT(func(v AS2OneWayAgreement) BusinessIdentity { return v.SenderBusinessIdentity }).(BusinessIdentityOutput)
}

type AS2OneWayAgreementPtrOutput struct{ *pulumi.OutputState }

func (AS2OneWayAgreementPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2OneWayAgreement)(nil)).Elem()
}

func (o AS2OneWayAgreementPtrOutput) ToAS2OneWayAgreementPtrOutput() AS2OneWayAgreementPtrOutput {
	return o
}

func (o AS2OneWayAgreementPtrOutput) ToAS2OneWayAgreementPtrOutputWithContext(ctx context.Context) AS2OneWayAgreementPtrOutput {
	return o
}

func (o AS2OneWayAgreementPtrOutput) Elem() AS2OneWayAgreementOutput {
	return o.ApplyT(func(v *AS2OneWayAgreement) AS2OneWayAgreement {
		if v != nil {
			return *v
		}
		var ret AS2OneWayAgreement
		return ret
	}).(AS2OneWayAgreementOutput)
}

func (o AS2OneWayAgreementPtrOutput) ProtocolSettings() AS2ProtocolSettingsPtrOutput {
	return o.ApplyT(func(v *AS2OneWayAgreement) *AS2ProtocolSettings {
		if v == nil {
			return nil
		}
		return &v.ProtocolSettings
	}).(AS2ProtocolSettingsPtrOutput)
}

func (o AS2OneWayAgreementPtrOutput) ReceiverBusinessIdentity() BusinessIdentityPtrOutput {
	return o.ApplyT(func(v *AS2OneWayAgreement) *BusinessIdentity {
		if v == nil {
			return nil
		}
		return &v.ReceiverBusinessIdentity
	}).(BusinessIdentityPtrOutput)
}

func (o AS2OneWayAgreementPtrOutput) SenderBusinessIdentity() BusinessIdentityPtrOutput {
	return o.ApplyT(func(v *AS2OneWayAgreement) *BusinessIdentity {
		if v == nil {
			return nil
		}
		return &v.SenderBusinessIdentity
	}).(BusinessIdentityPtrOutput)
}

type AS2OneWayAgreementResponse struct {
	ProtocolSettings         AS2ProtocolSettingsResponse `pulumi:"protocolSettings"`
	ReceiverBusinessIdentity BusinessIdentityResponse    `pulumi:"receiverBusinessIdentity"`
	SenderBusinessIdentity   BusinessIdentityResponse    `pulumi:"senderBusinessIdentity"`
}

type AS2OneWayAgreementResponseOutput struct{ *pulumi.OutputState }

func (AS2OneWayAgreementResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2OneWayAgreementResponse)(nil)).Elem()
}

func (o AS2OneWayAgreementResponseOutput) ToAS2OneWayAgreementResponseOutput() AS2OneWayAgreementResponseOutput {
	return o
}

func (o AS2OneWayAgreementResponseOutput) ToAS2OneWayAgreementResponseOutputWithContext(ctx context.Context) AS2OneWayAgreementResponseOutput {
	return o
}

func (o AS2OneWayAgreementResponseOutput) ProtocolSettings() AS2ProtocolSettingsResponseOutput {
	return o.ApplyT(func(v AS2OneWayAgreementResponse) AS2ProtocolSettingsResponse { return v.ProtocolSettings }).(AS2ProtocolSettingsResponseOutput)
}

func (o AS2OneWayAgreementResponseOutput) ReceiverBusinessIdentity() BusinessIdentityResponseOutput {
	return o.ApplyT(func(v AS2OneWayAgreementResponse) BusinessIdentityResponse { return v.ReceiverBusinessIdentity }).(BusinessIdentityResponseOutput)
}

func (o AS2OneWayAgreementResponseOutput) SenderBusinessIdentity() BusinessIdentityResponseOutput {
	return o.ApplyT(func(v AS2OneWayAgreementResponse) BusinessIdentityResponse { return v.SenderBusinessIdentity }).(BusinessIdentityResponseOutput)
}

type AS2OneWayAgreementResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2OneWayAgreementResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2OneWayAgreementResponse)(nil)).Elem()
}

func (o AS2OneWayAgreementResponsePtrOutput) ToAS2OneWayAgreementResponsePtrOutput() AS2OneWayAgreementResponsePtrOutput {
	return o
}

func (o AS2OneWayAgreementResponsePtrOutput) ToAS2OneWayAgreementResponsePtrOutputWithContext(ctx context.Context) AS2OneWayAgreementResponsePtrOutput {
	return o
}

func (o AS2OneWayAgreementResponsePtrOutput) Elem() AS2OneWayAgreementResponseOutput {
	return o.ApplyT(func(v *AS2OneWayAgreementResponse) AS2OneWayAgreementResponse {
		if v != nil {
			return *v
		}
		var ret AS2OneWayAgreementResponse
		return ret
	}).(AS2OneWayAgreementResponseOutput)
}

func (o AS2OneWayAgreementResponsePtrOutput) ProtocolSettings() AS2ProtocolSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AS2OneWayAgreementResponse) *AS2ProtocolSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.ProtocolSettings
	}).(AS2ProtocolSettingsResponsePtrOutput)
}

func (o AS2OneWayAgreementResponsePtrOutput) ReceiverBusinessIdentity() BusinessIdentityResponsePtrOutput {
	return o.ApplyT(func(v *AS2OneWayAgreementResponse) *BusinessIdentityResponse {
		if v == nil {
			return nil
		}
		return &v.ReceiverBusinessIdentity
	}).(BusinessIdentityResponsePtrOutput)
}

func (o AS2OneWayAgreementResponsePtrOutput) SenderBusinessIdentity() BusinessIdentityResponsePtrOutput {
	return o.ApplyT(func(v *AS2OneWayAgreementResponse) *BusinessIdentityResponse {
		if v == nil {
			return nil
		}
		return &v.SenderBusinessIdentity
	}).(BusinessIdentityResponsePtrOutput)
}

type AS2ProtocolSettings struct {
	AcknowledgementConnectionSettings AS2AcknowledgementConnectionSettings `pulumi:"acknowledgementConnectionSettings"`
	EnvelopeSettings                  AS2EnvelopeSettings                  `pulumi:"envelopeSettings"`
	ErrorSettings                     AS2ErrorSettings                     `pulumi:"errorSettings"`
	MdnSettings                       AS2MdnSettings                       `pulumi:"mdnSettings"`
	MessageConnectionSettings         AS2MessageConnectionSettings         `pulumi:"messageConnectionSettings"`
	SecuritySettings                  AS2SecuritySettings                  `pulumi:"securitySettings"`
	ValidationSettings                AS2ValidationSettings                `pulumi:"validationSettings"`
}





type AS2ProtocolSettingsInput interface {
	pulumi.Input

	ToAS2ProtocolSettingsOutput() AS2ProtocolSettingsOutput
	ToAS2ProtocolSettingsOutputWithContext(context.Context) AS2ProtocolSettingsOutput
}

type AS2ProtocolSettingsArgs struct {
	AcknowledgementConnectionSettings AS2AcknowledgementConnectionSettingsInput `pulumi:"acknowledgementConnectionSettings"`
	EnvelopeSettings                  AS2EnvelopeSettingsInput                  `pulumi:"envelopeSettings"`
	ErrorSettings                     AS2ErrorSettingsInput                     `pulumi:"errorSettings"`
	MdnSettings                       AS2MdnSettingsInput                       `pulumi:"mdnSettings"`
	MessageConnectionSettings         AS2MessageConnectionSettingsInput         `pulumi:"messageConnectionSettings"`
	SecuritySettings                  AS2SecuritySettingsInput                  `pulumi:"securitySettings"`
	ValidationSettings                AS2ValidationSettingsInput                `pulumi:"validationSettings"`
}

func (AS2ProtocolSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2ProtocolSettings)(nil)).Elem()
}

func (i AS2ProtocolSettingsArgs) ToAS2ProtocolSettingsOutput() AS2ProtocolSettingsOutput {
	return i.ToAS2ProtocolSettingsOutputWithContext(context.Background())
}

func (i AS2ProtocolSettingsArgs) ToAS2ProtocolSettingsOutputWithContext(ctx context.Context) AS2ProtocolSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2ProtocolSettingsOutput)
}

func (i AS2ProtocolSettingsArgs) ToAS2ProtocolSettingsPtrOutput() AS2ProtocolSettingsPtrOutput {
	return i.ToAS2ProtocolSettingsPtrOutputWithContext(context.Background())
}

func (i AS2ProtocolSettingsArgs) ToAS2ProtocolSettingsPtrOutputWithContext(ctx context.Context) AS2ProtocolSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2ProtocolSettingsOutput).ToAS2ProtocolSettingsPtrOutputWithContext(ctx)
}









type AS2ProtocolSettingsPtrInput interface {
	pulumi.Input

	ToAS2ProtocolSettingsPtrOutput() AS2ProtocolSettingsPtrOutput
	ToAS2ProtocolSettingsPtrOutputWithContext(context.Context) AS2ProtocolSettingsPtrOutput
}

type as2protocolSettingsPtrType AS2ProtocolSettingsArgs

func AS2ProtocolSettingsPtr(v *AS2ProtocolSettingsArgs) AS2ProtocolSettingsPtrInput {
	return (*as2protocolSettingsPtrType)(v)
}

func (*as2protocolSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2ProtocolSettings)(nil)).Elem()
}

func (i *as2protocolSettingsPtrType) ToAS2ProtocolSettingsPtrOutput() AS2ProtocolSettingsPtrOutput {
	return i.ToAS2ProtocolSettingsPtrOutputWithContext(context.Background())
}

func (i *as2protocolSettingsPtrType) ToAS2ProtocolSettingsPtrOutputWithContext(ctx context.Context) AS2ProtocolSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2ProtocolSettingsPtrOutput)
}

type AS2ProtocolSettingsOutput struct{ *pulumi.OutputState }

func (AS2ProtocolSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2ProtocolSettings)(nil)).Elem()
}

func (o AS2ProtocolSettingsOutput) ToAS2ProtocolSettingsOutput() AS2ProtocolSettingsOutput {
	return o
}

func (o AS2ProtocolSettingsOutput) ToAS2ProtocolSettingsOutputWithContext(ctx context.Context) AS2ProtocolSettingsOutput {
	return o
}

func (o AS2ProtocolSettingsOutput) ToAS2ProtocolSettingsPtrOutput() AS2ProtocolSettingsPtrOutput {
	return o.ToAS2ProtocolSettingsPtrOutputWithContext(context.Background())
}

func (o AS2ProtocolSettingsOutput) ToAS2ProtocolSettingsPtrOutputWithContext(ctx context.Context) AS2ProtocolSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2ProtocolSettings) *AS2ProtocolSettings {
		return &v
	}).(AS2ProtocolSettingsPtrOutput)
}

func (o AS2ProtocolSettingsOutput) AcknowledgementConnectionSettings() AS2AcknowledgementConnectionSettingsOutput {
	return o.ApplyT(func(v AS2ProtocolSettings) AS2AcknowledgementConnectionSettings {
		return v.AcknowledgementConnectionSettings
	}).(AS2AcknowledgementConnectionSettingsOutput)
}

func (o AS2ProtocolSettingsOutput) EnvelopeSettings() AS2EnvelopeSettingsOutput {
	return o.ApplyT(func(v AS2ProtocolSettings) AS2EnvelopeSettings { return v.EnvelopeSettings }).(AS2EnvelopeSettingsOutput)
}

func (o AS2ProtocolSettingsOutput) ErrorSettings() AS2ErrorSettingsOutput {
	return o.ApplyT(func(v AS2ProtocolSettings) AS2ErrorSettings { return v.ErrorSettings }).(AS2ErrorSettingsOutput)
}

func (o AS2ProtocolSettingsOutput) MdnSettings() AS2MdnSettingsOutput {
	return o.ApplyT(func(v AS2ProtocolSettings) AS2MdnSettings { return v.MdnSettings }).(AS2MdnSettingsOutput)
}

func (o AS2ProtocolSettingsOutput) MessageConnectionSettings() AS2MessageConnectionSettingsOutput {
	return o.ApplyT(func(v AS2ProtocolSettings) AS2MessageConnectionSettings { return v.MessageConnectionSettings }).(AS2MessageConnectionSettingsOutput)
}

func (o AS2ProtocolSettingsOutput) SecuritySettings() AS2SecuritySettingsOutput {
	return o.ApplyT(func(v AS2ProtocolSettings) AS2SecuritySettings { return v.SecuritySettings }).(AS2SecuritySettingsOutput)
}

func (o AS2ProtocolSettingsOutput) ValidationSettings() AS2ValidationSettingsOutput {
	return o.ApplyT(func(v AS2ProtocolSettings) AS2ValidationSettings { return v.ValidationSettings }).(AS2ValidationSettingsOutput)
}

type AS2ProtocolSettingsPtrOutput struct{ *pulumi.OutputState }

func (AS2ProtocolSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2ProtocolSettings)(nil)).Elem()
}

func (o AS2ProtocolSettingsPtrOutput) ToAS2ProtocolSettingsPtrOutput() AS2ProtocolSettingsPtrOutput {
	return o
}

func (o AS2ProtocolSettingsPtrOutput) ToAS2ProtocolSettingsPtrOutputWithContext(ctx context.Context) AS2ProtocolSettingsPtrOutput {
	return o
}

func (o AS2ProtocolSettingsPtrOutput) Elem() AS2ProtocolSettingsOutput {
	return o.ApplyT(func(v *AS2ProtocolSettings) AS2ProtocolSettings {
		if v != nil {
			return *v
		}
		var ret AS2ProtocolSettings
		return ret
	}).(AS2ProtocolSettingsOutput)
}

func (o AS2ProtocolSettingsPtrOutput) AcknowledgementConnectionSettings() AS2AcknowledgementConnectionSettingsPtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettings) *AS2AcknowledgementConnectionSettings {
		if v == nil {
			return nil
		}
		return &v.AcknowledgementConnectionSettings
	}).(AS2AcknowledgementConnectionSettingsPtrOutput)
}

func (o AS2ProtocolSettingsPtrOutput) EnvelopeSettings() AS2EnvelopeSettingsPtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettings) *AS2EnvelopeSettings {
		if v == nil {
			return nil
		}
		return &v.EnvelopeSettings
	}).(AS2EnvelopeSettingsPtrOutput)
}

func (o AS2ProtocolSettingsPtrOutput) ErrorSettings() AS2ErrorSettingsPtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettings) *AS2ErrorSettings {
		if v == nil {
			return nil
		}
		return &v.ErrorSettings
	}).(AS2ErrorSettingsPtrOutput)
}

func (o AS2ProtocolSettingsPtrOutput) MdnSettings() AS2MdnSettingsPtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettings) *AS2MdnSettings {
		if v == nil {
			return nil
		}
		return &v.MdnSettings
	}).(AS2MdnSettingsPtrOutput)
}

func (o AS2ProtocolSettingsPtrOutput) MessageConnectionSettings() AS2MessageConnectionSettingsPtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettings) *AS2MessageConnectionSettings {
		if v == nil {
			return nil
		}
		return &v.MessageConnectionSettings
	}).(AS2MessageConnectionSettingsPtrOutput)
}

func (o AS2ProtocolSettingsPtrOutput) SecuritySettings() AS2SecuritySettingsPtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettings) *AS2SecuritySettings {
		if v == nil {
			return nil
		}
		return &v.SecuritySettings
	}).(AS2SecuritySettingsPtrOutput)
}

func (o AS2ProtocolSettingsPtrOutput) ValidationSettings() AS2ValidationSettingsPtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettings) *AS2ValidationSettings {
		if v == nil {
			return nil
		}
		return &v.ValidationSettings
	}).(AS2ValidationSettingsPtrOutput)
}

type AS2ProtocolSettingsResponse struct {
	AcknowledgementConnectionSettings AS2AcknowledgementConnectionSettingsResponse `pulumi:"acknowledgementConnectionSettings"`
	EnvelopeSettings                  AS2EnvelopeSettingsResponse                  `pulumi:"envelopeSettings"`
	ErrorSettings                     AS2ErrorSettingsResponse                     `pulumi:"errorSettings"`
	MdnSettings                       AS2MdnSettingsResponse                       `pulumi:"mdnSettings"`
	MessageConnectionSettings         AS2MessageConnectionSettingsResponse         `pulumi:"messageConnectionSettings"`
	SecuritySettings                  AS2SecuritySettingsResponse                  `pulumi:"securitySettings"`
	ValidationSettings                AS2ValidationSettingsResponse                `pulumi:"validationSettings"`
}

type AS2ProtocolSettingsResponseOutput struct{ *pulumi.OutputState }

func (AS2ProtocolSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2ProtocolSettingsResponse)(nil)).Elem()
}

func (o AS2ProtocolSettingsResponseOutput) ToAS2ProtocolSettingsResponseOutput() AS2ProtocolSettingsResponseOutput {
	return o
}

func (o AS2ProtocolSettingsResponseOutput) ToAS2ProtocolSettingsResponseOutputWithContext(ctx context.Context) AS2ProtocolSettingsResponseOutput {
	return o
}

func (o AS2ProtocolSettingsResponseOutput) AcknowledgementConnectionSettings() AS2AcknowledgementConnectionSettingsResponseOutput {
	return o.ApplyT(func(v AS2ProtocolSettingsResponse) AS2AcknowledgementConnectionSettingsResponse {
		return v.AcknowledgementConnectionSettings
	}).(AS2AcknowledgementConnectionSettingsResponseOutput)
}

func (o AS2ProtocolSettingsResponseOutput) EnvelopeSettings() AS2EnvelopeSettingsResponseOutput {
	return o.ApplyT(func(v AS2ProtocolSettingsResponse) AS2EnvelopeSettingsResponse { return v.EnvelopeSettings }).(AS2EnvelopeSettingsResponseOutput)
}

func (o AS2ProtocolSettingsResponseOutput) ErrorSettings() AS2ErrorSettingsResponseOutput {
	return o.ApplyT(func(v AS2ProtocolSettingsResponse) AS2ErrorSettingsResponse { return v.ErrorSettings }).(AS2ErrorSettingsResponseOutput)
}

func (o AS2ProtocolSettingsResponseOutput) MdnSettings() AS2MdnSettingsResponseOutput {
	return o.ApplyT(func(v AS2ProtocolSettingsResponse) AS2MdnSettingsResponse { return v.MdnSettings }).(AS2MdnSettingsResponseOutput)
}

func (o AS2ProtocolSettingsResponseOutput) MessageConnectionSettings() AS2MessageConnectionSettingsResponseOutput {
	return o.ApplyT(func(v AS2ProtocolSettingsResponse) AS2MessageConnectionSettingsResponse {
		return v.MessageConnectionSettings
	}).(AS2MessageConnectionSettingsResponseOutput)
}

func (o AS2ProtocolSettingsResponseOutput) SecuritySettings() AS2SecuritySettingsResponseOutput {
	return o.ApplyT(func(v AS2ProtocolSettingsResponse) AS2SecuritySettingsResponse { return v.SecuritySettings }).(AS2SecuritySettingsResponseOutput)
}

func (o AS2ProtocolSettingsResponseOutput) ValidationSettings() AS2ValidationSettingsResponseOutput {
	return o.ApplyT(func(v AS2ProtocolSettingsResponse) AS2ValidationSettingsResponse { return v.ValidationSettings }).(AS2ValidationSettingsResponseOutput)
}

type AS2ProtocolSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2ProtocolSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2ProtocolSettingsResponse)(nil)).Elem()
}

func (o AS2ProtocolSettingsResponsePtrOutput) ToAS2ProtocolSettingsResponsePtrOutput() AS2ProtocolSettingsResponsePtrOutput {
	return o
}

func (o AS2ProtocolSettingsResponsePtrOutput) ToAS2ProtocolSettingsResponsePtrOutputWithContext(ctx context.Context) AS2ProtocolSettingsResponsePtrOutput {
	return o
}

func (o AS2ProtocolSettingsResponsePtrOutput) Elem() AS2ProtocolSettingsResponseOutput {
	return o.ApplyT(func(v *AS2ProtocolSettingsResponse) AS2ProtocolSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AS2ProtocolSettingsResponse
		return ret
	}).(AS2ProtocolSettingsResponseOutput)
}

func (o AS2ProtocolSettingsResponsePtrOutput) AcknowledgementConnectionSettings() AS2AcknowledgementConnectionSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettingsResponse) *AS2AcknowledgementConnectionSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.AcknowledgementConnectionSettings
	}).(AS2AcknowledgementConnectionSettingsResponsePtrOutput)
}

func (o AS2ProtocolSettingsResponsePtrOutput) EnvelopeSettings() AS2EnvelopeSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettingsResponse) *AS2EnvelopeSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.EnvelopeSettings
	}).(AS2EnvelopeSettingsResponsePtrOutput)
}

func (o AS2ProtocolSettingsResponsePtrOutput) ErrorSettings() AS2ErrorSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettingsResponse) *AS2ErrorSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.ErrorSettings
	}).(AS2ErrorSettingsResponsePtrOutput)
}

func (o AS2ProtocolSettingsResponsePtrOutput) MdnSettings() AS2MdnSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettingsResponse) *AS2MdnSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.MdnSettings
	}).(AS2MdnSettingsResponsePtrOutput)
}

func (o AS2ProtocolSettingsResponsePtrOutput) MessageConnectionSettings() AS2MessageConnectionSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettingsResponse) *AS2MessageConnectionSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.MessageConnectionSettings
	}).(AS2MessageConnectionSettingsResponsePtrOutput)
}

func (o AS2ProtocolSettingsResponsePtrOutput) SecuritySettings() AS2SecuritySettingsResponsePtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettingsResponse) *AS2SecuritySettingsResponse {
		if v == nil {
			return nil
		}
		return &v.SecuritySettings
	}).(AS2SecuritySettingsResponsePtrOutput)
}

func (o AS2ProtocolSettingsResponsePtrOutput) ValidationSettings() AS2ValidationSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AS2ProtocolSettingsResponse) *AS2ValidationSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.ValidationSettings
	}).(AS2ValidationSettingsResponsePtrOutput)
}

type AS2SecuritySettings struct {
	EnableNRRForInboundDecodedMessages  bool    `pulumi:"enableNRRForInboundDecodedMessages"`
	EnableNRRForInboundEncodedMessages  bool    `pulumi:"enableNRRForInboundEncodedMessages"`
	EnableNRRForInboundMDN              bool    `pulumi:"enableNRRForInboundMDN"`
	EnableNRRForOutboundDecodedMessages bool    `pulumi:"enableNRRForOutboundDecodedMessages"`
	EnableNRRForOutboundEncodedMessages bool    `pulumi:"enableNRRForOutboundEncodedMessages"`
	EnableNRRForOutboundMDN             bool    `pulumi:"enableNRRForOutboundMDN"`
	EncryptionCertificateName           *string `pulumi:"encryptionCertificateName"`
	OverrideGroupSigningCertificate     bool    `pulumi:"overrideGroupSigningCertificate"`
	Sha2AlgorithmFormat                 *string `pulumi:"sha2AlgorithmFormat"`
	SigningCertificateName              *string `pulumi:"signingCertificateName"`
}





type AS2SecuritySettingsInput interface {
	pulumi.Input

	ToAS2SecuritySettingsOutput() AS2SecuritySettingsOutput
	ToAS2SecuritySettingsOutputWithContext(context.Context) AS2SecuritySettingsOutput
}

type AS2SecuritySettingsArgs struct {
	EnableNRRForInboundDecodedMessages  pulumi.BoolInput      `pulumi:"enableNRRForInboundDecodedMessages"`
	EnableNRRForInboundEncodedMessages  pulumi.BoolInput      `pulumi:"enableNRRForInboundEncodedMessages"`
	EnableNRRForInboundMDN              pulumi.BoolInput      `pulumi:"enableNRRForInboundMDN"`
	EnableNRRForOutboundDecodedMessages pulumi.BoolInput      `pulumi:"enableNRRForOutboundDecodedMessages"`
	EnableNRRForOutboundEncodedMessages pulumi.BoolInput      `pulumi:"enableNRRForOutboundEncodedMessages"`
	EnableNRRForOutboundMDN             pulumi.BoolInput      `pulumi:"enableNRRForOutboundMDN"`
	EncryptionCertificateName           pulumi.StringPtrInput `pulumi:"encryptionCertificateName"`
	OverrideGroupSigningCertificate     pulumi.BoolInput      `pulumi:"overrideGroupSigningCertificate"`
	Sha2AlgorithmFormat                 pulumi.StringPtrInput `pulumi:"sha2AlgorithmFormat"`
	SigningCertificateName              pulumi.StringPtrInput `pulumi:"signingCertificateName"`
}

func (AS2SecuritySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2SecuritySettings)(nil)).Elem()
}

func (i AS2SecuritySettingsArgs) ToAS2SecuritySettingsOutput() AS2SecuritySettingsOutput {
	return i.ToAS2SecuritySettingsOutputWithContext(context.Background())
}

func (i AS2SecuritySettingsArgs) ToAS2SecuritySettingsOutputWithContext(ctx context.Context) AS2SecuritySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2SecuritySettingsOutput)
}

func (i AS2SecuritySettingsArgs) ToAS2SecuritySettingsPtrOutput() AS2SecuritySettingsPtrOutput {
	return i.ToAS2SecuritySettingsPtrOutputWithContext(context.Background())
}

func (i AS2SecuritySettingsArgs) ToAS2SecuritySettingsPtrOutputWithContext(ctx context.Context) AS2SecuritySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2SecuritySettingsOutput).ToAS2SecuritySettingsPtrOutputWithContext(ctx)
}









type AS2SecuritySettingsPtrInput interface {
	pulumi.Input

	ToAS2SecuritySettingsPtrOutput() AS2SecuritySettingsPtrOutput
	ToAS2SecuritySettingsPtrOutputWithContext(context.Context) AS2SecuritySettingsPtrOutput
}

type as2securitySettingsPtrType AS2SecuritySettingsArgs

func AS2SecuritySettingsPtr(v *AS2SecuritySettingsArgs) AS2SecuritySettingsPtrInput {
	return (*as2securitySettingsPtrType)(v)
}

func (*as2securitySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2SecuritySettings)(nil)).Elem()
}

func (i *as2securitySettingsPtrType) ToAS2SecuritySettingsPtrOutput() AS2SecuritySettingsPtrOutput {
	return i.ToAS2SecuritySettingsPtrOutputWithContext(context.Background())
}

func (i *as2securitySettingsPtrType) ToAS2SecuritySettingsPtrOutputWithContext(ctx context.Context) AS2SecuritySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2SecuritySettingsPtrOutput)
}

type AS2SecuritySettingsOutput struct{ *pulumi.OutputState }

func (AS2SecuritySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2SecuritySettings)(nil)).Elem()
}

func (o AS2SecuritySettingsOutput) ToAS2SecuritySettingsOutput() AS2SecuritySettingsOutput {
	return o
}

func (o AS2SecuritySettingsOutput) ToAS2SecuritySettingsOutputWithContext(ctx context.Context) AS2SecuritySettingsOutput {
	return o
}

func (o AS2SecuritySettingsOutput) ToAS2SecuritySettingsPtrOutput() AS2SecuritySettingsPtrOutput {
	return o.ToAS2SecuritySettingsPtrOutputWithContext(context.Background())
}

func (o AS2SecuritySettingsOutput) ToAS2SecuritySettingsPtrOutputWithContext(ctx context.Context) AS2SecuritySettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2SecuritySettings) *AS2SecuritySettings {
		return &v
	}).(AS2SecuritySettingsPtrOutput)
}

func (o AS2SecuritySettingsOutput) EnableNRRForInboundDecodedMessages() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2SecuritySettings) bool { return v.EnableNRRForInboundDecodedMessages }).(pulumi.BoolOutput)
}

func (o AS2SecuritySettingsOutput) EnableNRRForInboundEncodedMessages() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2SecuritySettings) bool { return v.EnableNRRForInboundEncodedMessages }).(pulumi.BoolOutput)
}

func (o AS2SecuritySettingsOutput) EnableNRRForInboundMDN() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2SecuritySettings) bool { return v.EnableNRRForInboundMDN }).(pulumi.BoolOutput)
}

func (o AS2SecuritySettingsOutput) EnableNRRForOutboundDecodedMessages() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2SecuritySettings) bool { return v.EnableNRRForOutboundDecodedMessages }).(pulumi.BoolOutput)
}

func (o AS2SecuritySettingsOutput) EnableNRRForOutboundEncodedMessages() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2SecuritySettings) bool { return v.EnableNRRForOutboundEncodedMessages }).(pulumi.BoolOutput)
}

func (o AS2SecuritySettingsOutput) EnableNRRForOutboundMDN() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2SecuritySettings) bool { return v.EnableNRRForOutboundMDN }).(pulumi.BoolOutput)
}

func (o AS2SecuritySettingsOutput) EncryptionCertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettings) *string { return v.EncryptionCertificateName }).(pulumi.StringPtrOutput)
}

func (o AS2SecuritySettingsOutput) OverrideGroupSigningCertificate() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2SecuritySettings) bool { return v.OverrideGroupSigningCertificate }).(pulumi.BoolOutput)
}

func (o AS2SecuritySettingsOutput) Sha2AlgorithmFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettings) *string { return v.Sha2AlgorithmFormat }).(pulumi.StringPtrOutput)
}

func (o AS2SecuritySettingsOutput) SigningCertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettings) *string { return v.SigningCertificateName }).(pulumi.StringPtrOutput)
}

type AS2SecuritySettingsPtrOutput struct{ *pulumi.OutputState }

func (AS2SecuritySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2SecuritySettings)(nil)).Elem()
}

func (o AS2SecuritySettingsPtrOutput) ToAS2SecuritySettingsPtrOutput() AS2SecuritySettingsPtrOutput {
	return o
}

func (o AS2SecuritySettingsPtrOutput) ToAS2SecuritySettingsPtrOutputWithContext(ctx context.Context) AS2SecuritySettingsPtrOutput {
	return o
}

func (o AS2SecuritySettingsPtrOutput) Elem() AS2SecuritySettingsOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) AS2SecuritySettings {
		if v != nil {
			return *v
		}
		var ret AS2SecuritySettings
		return ret
	}).(AS2SecuritySettingsOutput)
}

func (o AS2SecuritySettingsPtrOutput) EnableNRRForInboundDecodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *bool {
		if v == nil {
			return nil
		}
		return &v.EnableNRRForInboundDecodedMessages
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsPtrOutput) EnableNRRForInboundEncodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *bool {
		if v == nil {
			return nil
		}
		return &v.EnableNRRForInboundEncodedMessages
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsPtrOutput) EnableNRRForInboundMDN() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *bool {
		if v == nil {
			return nil
		}
		return &v.EnableNRRForInboundMDN
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsPtrOutput) EnableNRRForOutboundDecodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *bool {
		if v == nil {
			return nil
		}
		return &v.EnableNRRForOutboundDecodedMessages
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsPtrOutput) EnableNRRForOutboundEncodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *bool {
		if v == nil {
			return nil
		}
		return &v.EnableNRRForOutboundEncodedMessages
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsPtrOutput) EnableNRRForOutboundMDN() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *bool {
		if v == nil {
			return nil
		}
		return &v.EnableNRRForOutboundMDN
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsPtrOutput) EncryptionCertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionCertificateName
	}).(pulumi.StringPtrOutput)
}

func (o AS2SecuritySettingsPtrOutput) OverrideGroupSigningCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *bool {
		if v == nil {
			return nil
		}
		return &v.OverrideGroupSigningCertificate
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsPtrOutput) Sha2AlgorithmFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *string {
		if v == nil {
			return nil
		}
		return v.Sha2AlgorithmFormat
	}).(pulumi.StringPtrOutput)
}

func (o AS2SecuritySettingsPtrOutput) SigningCertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettings) *string {
		if v == nil {
			return nil
		}
		return v.SigningCertificateName
	}).(pulumi.StringPtrOutput)
}

type AS2SecuritySettingsResponse struct {
	EnableNRRForInboundDecodedMessages  bool    `pulumi:"enableNRRForInboundDecodedMessages"`
	EnableNRRForInboundEncodedMessages  bool    `pulumi:"enableNRRForInboundEncodedMessages"`
	EnableNRRForInboundMDN              bool    `pulumi:"enableNRRForInboundMDN"`
	EnableNRRForOutboundDecodedMessages bool    `pulumi:"enableNRRForOutboundDecodedMessages"`
	EnableNRRForOutboundEncodedMessages bool    `pulumi:"enableNRRForOutboundEncodedMessages"`
	EnableNRRForOutboundMDN             bool    `pulumi:"enableNRRForOutboundMDN"`
	EncryptionCertificateName           *string `pulumi:"encryptionCertificateName"`
	OverrideGroupSigningCertificate     bool    `pulumi:"overrideGroupSigningCertificate"`
	Sha2AlgorithmFormat                 *string `pulumi:"sha2AlgorithmFormat"`
	SigningCertificateName              *string `pulumi:"signingCertificateName"`
}

type AS2SecuritySettingsResponseOutput struct{ *pulumi.OutputState }

func (AS2SecuritySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2SecuritySettingsResponse)(nil)).Elem()
}

func (o AS2SecuritySettingsResponseOutput) ToAS2SecuritySettingsResponseOutput() AS2SecuritySettingsResponseOutput {
	return o
}

func (o AS2SecuritySettingsResponseOutput) ToAS2SecuritySettingsResponseOutputWithContext(ctx context.Context) AS2SecuritySettingsResponseOutput {
	return o
}

func (o AS2SecuritySettingsResponseOutput) EnableNRRForInboundDecodedMessages() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) bool { return v.EnableNRRForInboundDecodedMessages }).(pulumi.BoolOutput)
}

func (o AS2SecuritySettingsResponseOutput) EnableNRRForInboundEncodedMessages() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) bool { return v.EnableNRRForInboundEncodedMessages }).(pulumi.BoolOutput)
}

func (o AS2SecuritySettingsResponseOutput) EnableNRRForInboundMDN() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) bool { return v.EnableNRRForInboundMDN }).(pulumi.BoolOutput)
}

func (o AS2SecuritySettingsResponseOutput) EnableNRRForOutboundDecodedMessages() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) bool { return v.EnableNRRForOutboundDecodedMessages }).(pulumi.BoolOutput)
}

func (o AS2SecuritySettingsResponseOutput) EnableNRRForOutboundEncodedMessages() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) bool { return v.EnableNRRForOutboundEncodedMessages }).(pulumi.BoolOutput)
}

func (o AS2SecuritySettingsResponseOutput) EnableNRRForOutboundMDN() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) bool { return v.EnableNRRForOutboundMDN }).(pulumi.BoolOutput)
}

func (o AS2SecuritySettingsResponseOutput) EncryptionCertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) *string { return v.EncryptionCertificateName }).(pulumi.StringPtrOutput)
}

func (o AS2SecuritySettingsResponseOutput) OverrideGroupSigningCertificate() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) bool { return v.OverrideGroupSigningCertificate }).(pulumi.BoolOutput)
}

func (o AS2SecuritySettingsResponseOutput) Sha2AlgorithmFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) *string { return v.Sha2AlgorithmFormat }).(pulumi.StringPtrOutput)
}

func (o AS2SecuritySettingsResponseOutput) SigningCertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2SecuritySettingsResponse) *string { return v.SigningCertificateName }).(pulumi.StringPtrOutput)
}

type AS2SecuritySettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2SecuritySettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2SecuritySettingsResponse)(nil)).Elem()
}

func (o AS2SecuritySettingsResponsePtrOutput) ToAS2SecuritySettingsResponsePtrOutput() AS2SecuritySettingsResponsePtrOutput {
	return o
}

func (o AS2SecuritySettingsResponsePtrOutput) ToAS2SecuritySettingsResponsePtrOutputWithContext(ctx context.Context) AS2SecuritySettingsResponsePtrOutput {
	return o
}

func (o AS2SecuritySettingsResponsePtrOutput) Elem() AS2SecuritySettingsResponseOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) AS2SecuritySettingsResponse {
		if v != nil {
			return *v
		}
		var ret AS2SecuritySettingsResponse
		return ret
	}).(AS2SecuritySettingsResponseOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) EnableNRRForInboundDecodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.EnableNRRForInboundDecodedMessages
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) EnableNRRForInboundEncodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.EnableNRRForInboundEncodedMessages
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) EnableNRRForInboundMDN() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.EnableNRRForInboundMDN
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) EnableNRRForOutboundDecodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.EnableNRRForOutboundDecodedMessages
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) EnableNRRForOutboundEncodedMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.EnableNRRForOutboundEncodedMessages
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) EnableNRRForOutboundMDN() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.EnableNRRForOutboundMDN
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) EncryptionCertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionCertificateName
	}).(pulumi.StringPtrOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) OverrideGroupSigningCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.OverrideGroupSigningCertificate
	}).(pulumi.BoolPtrOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) Sha2AlgorithmFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sha2AlgorithmFormat
	}).(pulumi.StringPtrOutput)
}

func (o AS2SecuritySettingsResponsePtrOutput) SigningCertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2SecuritySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SigningCertificateName
	}).(pulumi.StringPtrOutput)
}

type AS2ValidationSettings struct {
	CheckCertificateRevocationListOnReceive bool    `pulumi:"checkCertificateRevocationListOnReceive"`
	CheckCertificateRevocationListOnSend    bool    `pulumi:"checkCertificateRevocationListOnSend"`
	CheckDuplicateMessage                   bool    `pulumi:"checkDuplicateMessage"`
	CompressMessage                         bool    `pulumi:"compressMessage"`
	EncryptMessage                          bool    `pulumi:"encryptMessage"`
	EncryptionAlgorithm                     string  `pulumi:"encryptionAlgorithm"`
	InterchangeDuplicatesValidityDays       int     `pulumi:"interchangeDuplicatesValidityDays"`
	OverrideMessageProperties               bool    `pulumi:"overrideMessageProperties"`
	SignMessage                             bool    `pulumi:"signMessage"`
	SigningAlgorithm                        *string `pulumi:"signingAlgorithm"`
}





type AS2ValidationSettingsInput interface {
	pulumi.Input

	ToAS2ValidationSettingsOutput() AS2ValidationSettingsOutput
	ToAS2ValidationSettingsOutputWithContext(context.Context) AS2ValidationSettingsOutput
}

type AS2ValidationSettingsArgs struct {
	CheckCertificateRevocationListOnReceive pulumi.BoolInput      `pulumi:"checkCertificateRevocationListOnReceive"`
	CheckCertificateRevocationListOnSend    pulumi.BoolInput      `pulumi:"checkCertificateRevocationListOnSend"`
	CheckDuplicateMessage                   pulumi.BoolInput      `pulumi:"checkDuplicateMessage"`
	CompressMessage                         pulumi.BoolInput      `pulumi:"compressMessage"`
	EncryptMessage                          pulumi.BoolInput      `pulumi:"encryptMessage"`
	EncryptionAlgorithm                     pulumi.StringInput    `pulumi:"encryptionAlgorithm"`
	InterchangeDuplicatesValidityDays       pulumi.IntInput       `pulumi:"interchangeDuplicatesValidityDays"`
	OverrideMessageProperties               pulumi.BoolInput      `pulumi:"overrideMessageProperties"`
	SignMessage                             pulumi.BoolInput      `pulumi:"signMessage"`
	SigningAlgorithm                        pulumi.StringPtrInput `pulumi:"signingAlgorithm"`
}

func (AS2ValidationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2ValidationSettings)(nil)).Elem()
}

func (i AS2ValidationSettingsArgs) ToAS2ValidationSettingsOutput() AS2ValidationSettingsOutput {
	return i.ToAS2ValidationSettingsOutputWithContext(context.Background())
}

func (i AS2ValidationSettingsArgs) ToAS2ValidationSettingsOutputWithContext(ctx context.Context) AS2ValidationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2ValidationSettingsOutput)
}

func (i AS2ValidationSettingsArgs) ToAS2ValidationSettingsPtrOutput() AS2ValidationSettingsPtrOutput {
	return i.ToAS2ValidationSettingsPtrOutputWithContext(context.Background())
}

func (i AS2ValidationSettingsArgs) ToAS2ValidationSettingsPtrOutputWithContext(ctx context.Context) AS2ValidationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2ValidationSettingsOutput).ToAS2ValidationSettingsPtrOutputWithContext(ctx)
}









type AS2ValidationSettingsPtrInput interface {
	pulumi.Input

	ToAS2ValidationSettingsPtrOutput() AS2ValidationSettingsPtrOutput
	ToAS2ValidationSettingsPtrOutputWithContext(context.Context) AS2ValidationSettingsPtrOutput
}

type as2validationSettingsPtrType AS2ValidationSettingsArgs

func AS2ValidationSettingsPtr(v *AS2ValidationSettingsArgs) AS2ValidationSettingsPtrInput {
	return (*as2validationSettingsPtrType)(v)
}

func (*as2validationSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2ValidationSettings)(nil)).Elem()
}

func (i *as2validationSettingsPtrType) ToAS2ValidationSettingsPtrOutput() AS2ValidationSettingsPtrOutput {
	return i.ToAS2ValidationSettingsPtrOutputWithContext(context.Background())
}

func (i *as2validationSettingsPtrType) ToAS2ValidationSettingsPtrOutputWithContext(ctx context.Context) AS2ValidationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AS2ValidationSettingsPtrOutput)
}

type AS2ValidationSettingsOutput struct{ *pulumi.OutputState }

func (AS2ValidationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2ValidationSettings)(nil)).Elem()
}

func (o AS2ValidationSettingsOutput) ToAS2ValidationSettingsOutput() AS2ValidationSettingsOutput {
	return o
}

func (o AS2ValidationSettingsOutput) ToAS2ValidationSettingsOutputWithContext(ctx context.Context) AS2ValidationSettingsOutput {
	return o
}

func (o AS2ValidationSettingsOutput) ToAS2ValidationSettingsPtrOutput() AS2ValidationSettingsPtrOutput {
	return o.ToAS2ValidationSettingsPtrOutputWithContext(context.Background())
}

func (o AS2ValidationSettingsOutput) ToAS2ValidationSettingsPtrOutputWithContext(ctx context.Context) AS2ValidationSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AS2ValidationSettings) *AS2ValidationSettings {
		return &v
	}).(AS2ValidationSettingsPtrOutput)
}

func (o AS2ValidationSettingsOutput) CheckCertificateRevocationListOnReceive() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2ValidationSettings) bool { return v.CheckCertificateRevocationListOnReceive }).(pulumi.BoolOutput)
}

func (o AS2ValidationSettingsOutput) CheckCertificateRevocationListOnSend() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2ValidationSettings) bool { return v.CheckCertificateRevocationListOnSend }).(pulumi.BoolOutput)
}

func (o AS2ValidationSettingsOutput) CheckDuplicateMessage() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2ValidationSettings) bool { return v.CheckDuplicateMessage }).(pulumi.BoolOutput)
}

func (o AS2ValidationSettingsOutput) CompressMessage() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2ValidationSettings) bool { return v.CompressMessage }).(pulumi.BoolOutput)
}

func (o AS2ValidationSettingsOutput) EncryptMessage() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2ValidationSettings) bool { return v.EncryptMessage }).(pulumi.BoolOutput)
}

func (o AS2ValidationSettingsOutput) EncryptionAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v AS2ValidationSettings) string { return v.EncryptionAlgorithm }).(pulumi.StringOutput)
}

func (o AS2ValidationSettingsOutput) InterchangeDuplicatesValidityDays() pulumi.IntOutput {
	return o.ApplyT(func(v AS2ValidationSettings) int { return v.InterchangeDuplicatesValidityDays }).(pulumi.IntOutput)
}

func (o AS2ValidationSettingsOutput) OverrideMessageProperties() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2ValidationSettings) bool { return v.OverrideMessageProperties }).(pulumi.BoolOutput)
}

func (o AS2ValidationSettingsOutput) SignMessage() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2ValidationSettings) bool { return v.SignMessage }).(pulumi.BoolOutput)
}

func (o AS2ValidationSettingsOutput) SigningAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettings) *string { return v.SigningAlgorithm }).(pulumi.StringPtrOutput)
}

type AS2ValidationSettingsPtrOutput struct{ *pulumi.OutputState }

func (AS2ValidationSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2ValidationSettings)(nil)).Elem()
}

func (o AS2ValidationSettingsPtrOutput) ToAS2ValidationSettingsPtrOutput() AS2ValidationSettingsPtrOutput {
	return o
}

func (o AS2ValidationSettingsPtrOutput) ToAS2ValidationSettingsPtrOutputWithContext(ctx context.Context) AS2ValidationSettingsPtrOutput {
	return o
}

func (o AS2ValidationSettingsPtrOutput) Elem() AS2ValidationSettingsOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) AS2ValidationSettings {
		if v != nil {
			return *v
		}
		var ret AS2ValidationSettings
		return ret
	}).(AS2ValidationSettingsOutput)
}

func (o AS2ValidationSettingsPtrOutput) CheckCertificateRevocationListOnReceive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.CheckCertificateRevocationListOnReceive
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsPtrOutput) CheckCertificateRevocationListOnSend() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.CheckCertificateRevocationListOnSend
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsPtrOutput) CheckDuplicateMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.CheckDuplicateMessage
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsPtrOutput) CompressMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.CompressMessage
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsPtrOutput) EncryptMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.EncryptMessage
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsPtrOutput) EncryptionAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *string {
		if v == nil {
			return nil
		}
		return &v.EncryptionAlgorithm
	}).(pulumi.StringPtrOutput)
}

func (o AS2ValidationSettingsPtrOutput) InterchangeDuplicatesValidityDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *int {
		if v == nil {
			return nil
		}
		return &v.InterchangeDuplicatesValidityDays
	}).(pulumi.IntPtrOutput)
}

func (o AS2ValidationSettingsPtrOutput) OverrideMessageProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.OverrideMessageProperties
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsPtrOutput) SignMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.SignMessage
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsPtrOutput) SigningAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettings) *string {
		if v == nil {
			return nil
		}
		return v.SigningAlgorithm
	}).(pulumi.StringPtrOutput)
}

type AS2ValidationSettingsResponse struct {
	CheckCertificateRevocationListOnReceive bool    `pulumi:"checkCertificateRevocationListOnReceive"`
	CheckCertificateRevocationListOnSend    bool    `pulumi:"checkCertificateRevocationListOnSend"`
	CheckDuplicateMessage                   bool    `pulumi:"checkDuplicateMessage"`
	CompressMessage                         bool    `pulumi:"compressMessage"`
	EncryptMessage                          bool    `pulumi:"encryptMessage"`
	EncryptionAlgorithm                     string  `pulumi:"encryptionAlgorithm"`
	InterchangeDuplicatesValidityDays       int     `pulumi:"interchangeDuplicatesValidityDays"`
	OverrideMessageProperties               bool    `pulumi:"overrideMessageProperties"`
	SignMessage                             bool    `pulumi:"signMessage"`
	SigningAlgorithm                        *string `pulumi:"signingAlgorithm"`
}

type AS2ValidationSettingsResponseOutput struct{ *pulumi.OutputState }

func (AS2ValidationSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AS2ValidationSettingsResponse)(nil)).Elem()
}

func (o AS2ValidationSettingsResponseOutput) ToAS2ValidationSettingsResponseOutput() AS2ValidationSettingsResponseOutput {
	return o
}

func (o AS2ValidationSettingsResponseOutput) ToAS2ValidationSettingsResponseOutputWithContext(ctx context.Context) AS2ValidationSettingsResponseOutput {
	return o
}

func (o AS2ValidationSettingsResponseOutput) CheckCertificateRevocationListOnReceive() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) bool { return v.CheckCertificateRevocationListOnReceive }).(pulumi.BoolOutput)
}

func (o AS2ValidationSettingsResponseOutput) CheckCertificateRevocationListOnSend() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) bool { return v.CheckCertificateRevocationListOnSend }).(pulumi.BoolOutput)
}

func (o AS2ValidationSettingsResponseOutput) CheckDuplicateMessage() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) bool { return v.CheckDuplicateMessage }).(pulumi.BoolOutput)
}

func (o AS2ValidationSettingsResponseOutput) CompressMessage() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) bool { return v.CompressMessage }).(pulumi.BoolOutput)
}

func (o AS2ValidationSettingsResponseOutput) EncryptMessage() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) bool { return v.EncryptMessage }).(pulumi.BoolOutput)
}

func (o AS2ValidationSettingsResponseOutput) EncryptionAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) string { return v.EncryptionAlgorithm }).(pulumi.StringOutput)
}

func (o AS2ValidationSettingsResponseOutput) InterchangeDuplicatesValidityDays() pulumi.IntOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) int { return v.InterchangeDuplicatesValidityDays }).(pulumi.IntOutput)
}

func (o AS2ValidationSettingsResponseOutput) OverrideMessageProperties() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) bool { return v.OverrideMessageProperties }).(pulumi.BoolOutput)
}

func (o AS2ValidationSettingsResponseOutput) SignMessage() pulumi.BoolOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) bool { return v.SignMessage }).(pulumi.BoolOutput)
}

func (o AS2ValidationSettingsResponseOutput) SigningAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AS2ValidationSettingsResponse) *string { return v.SigningAlgorithm }).(pulumi.StringPtrOutput)
}

type AS2ValidationSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AS2ValidationSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AS2ValidationSettingsResponse)(nil)).Elem()
}

func (o AS2ValidationSettingsResponsePtrOutput) ToAS2ValidationSettingsResponsePtrOutput() AS2ValidationSettingsResponsePtrOutput {
	return o
}

func (o AS2ValidationSettingsResponsePtrOutput) ToAS2ValidationSettingsResponsePtrOutputWithContext(ctx context.Context) AS2ValidationSettingsResponsePtrOutput {
	return o
}

func (o AS2ValidationSettingsResponsePtrOutput) Elem() AS2ValidationSettingsResponseOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) AS2ValidationSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AS2ValidationSettingsResponse
		return ret
	}).(AS2ValidationSettingsResponseOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) CheckCertificateRevocationListOnReceive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CheckCertificateRevocationListOnReceive
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) CheckCertificateRevocationListOnSend() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CheckCertificateRevocationListOnSend
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) CheckDuplicateMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CheckDuplicateMessage
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) CompressMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CompressMessage
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) EncryptMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.EncryptMessage
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) EncryptionAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EncryptionAlgorithm
	}).(pulumi.StringPtrOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) InterchangeDuplicatesValidityDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.InterchangeDuplicatesValidityDays
	}).(pulumi.IntPtrOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) OverrideMessageProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.OverrideMessageProperties
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) SignMessage() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.SignMessage
	}).(pulumi.BoolPtrOutput)
}

func (o AS2ValidationSettingsResponsePtrOutput) SigningAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AS2ValidationSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SigningAlgorithm
	}).(pulumi.StringPtrOutput)
}

type AgreementContent struct {
	AS2     *AS2AgreementContent     `pulumi:"aS2"`
	Edifact *EdifactAgreementContent `pulumi:"edifact"`
	X12     *X12AgreementContent     `pulumi:"x12"`
}





type AgreementContentInput interface {
	pulumi.Input

	ToAgreementContentOutput() AgreementContentOutput
	ToAgreementContentOutputWithContext(context.Context) AgreementContentOutput
}

type AgreementContentArgs struct {
	AS2     AS2AgreementContentPtrInput     `pulumi:"aS2"`
	Edifact EdifactAgreementContentPtrInput `pulumi:"edifact"`
	X12     X12AgreementContentPtrInput     `pulumi:"x12"`
}

func (AgreementContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AgreementContent)(nil)).Elem()
}

func (i AgreementContentArgs) ToAgreementContentOutput() AgreementContentOutput {
	return i.ToAgreementContentOutputWithContext(context.Background())
}

func (i AgreementContentArgs) ToAgreementContentOutputWithContext(ctx context.Context) AgreementContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgreementContentOutput)
}

type AgreementContentOutput struct{ *pulumi.OutputState }

func (AgreementContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgreementContent)(nil)).Elem()
}

func (o AgreementContentOutput) ToAgreementContentOutput() AgreementContentOutput {
	return o
}

func (o AgreementContentOutput) ToAgreementContentOutputWithContext(ctx context.Context) AgreementContentOutput {
	return o
}

func (o AgreementContentOutput) AS2() AS2AgreementContentPtrOutput {
	return o.ApplyT(func(v AgreementContent) *AS2AgreementContent { return v.AS2 }).(AS2AgreementContentPtrOutput)
}

func (o AgreementContentOutput) Edifact() EdifactAgreementContentPtrOutput {
	return o.ApplyT(func(v AgreementContent) *EdifactAgreementContent { return v.Edifact }).(EdifactAgreementContentPtrOutput)
}

func (o AgreementContentOutput) X12() X12AgreementContentPtrOutput {
	return o.ApplyT(func(v AgreementContent) *X12AgreementContent { return v.X12 }).(X12AgreementContentPtrOutput)
}

type AgreementContentResponse struct {
	AS2     *AS2AgreementContentResponse     `pulumi:"aS2"`
	Edifact *EdifactAgreementContentResponse `pulumi:"edifact"`
	X12     *X12AgreementContentResponse     `pulumi:"x12"`
}

type AgreementContentResponseOutput struct{ *pulumi.OutputState }

func (AgreementContentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgreementContentResponse)(nil)).Elem()
}

func (o AgreementContentResponseOutput) ToAgreementContentResponseOutput() AgreementContentResponseOutput {
	return o
}

func (o AgreementContentResponseOutput) ToAgreementContentResponseOutputWithContext(ctx context.Context) AgreementContentResponseOutput {
	return o
}

func (o AgreementContentResponseOutput) AS2() AS2AgreementContentResponsePtrOutput {
	return o.ApplyT(func(v AgreementContentResponse) *AS2AgreementContentResponse { return v.AS2 }).(AS2AgreementContentResponsePtrOutput)
}

func (o AgreementContentResponseOutput) Edifact() EdifactAgreementContentResponsePtrOutput {
	return o.ApplyT(func(v AgreementContentResponse) *EdifactAgreementContentResponse { return v.Edifact }).(EdifactAgreementContentResponsePtrOutput)
}

func (o AgreementContentResponseOutput) X12() X12AgreementContentResponsePtrOutput {
	return o.ApplyT(func(v AgreementContentResponse) *X12AgreementContentResponse { return v.X12 }).(X12AgreementContentResponsePtrOutput)
}

type ApiDeploymentParameterMetadataResponse struct {
	Description *string `pulumi:"description"`
	DisplayName *string `pulumi:"displayName"`
	IsRequired  *bool   `pulumi:"isRequired"`
	Type        *string `pulumi:"type"`
	Visibility  *string `pulumi:"visibility"`
}

type ApiDeploymentParameterMetadataResponseOutput struct{ *pulumi.OutputState }

func (ApiDeploymentParameterMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiDeploymentParameterMetadataResponse)(nil)).Elem()
}

func (o ApiDeploymentParameterMetadataResponseOutput) ToApiDeploymentParameterMetadataResponseOutput() ApiDeploymentParameterMetadataResponseOutput {
	return o
}

func (o ApiDeploymentParameterMetadataResponseOutput) ToApiDeploymentParameterMetadataResponseOutputWithContext(ctx context.Context) ApiDeploymentParameterMetadataResponseOutput {
	return o
}

func (o ApiDeploymentParameterMetadataResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiDeploymentParameterMetadataResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ApiDeploymentParameterMetadataResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiDeploymentParameterMetadataResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ApiDeploymentParameterMetadataResponseOutput) IsRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApiDeploymentParameterMetadataResponse) *bool { return v.IsRequired }).(pulumi.BoolPtrOutput)
}

func (o ApiDeploymentParameterMetadataResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiDeploymentParameterMetadataResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ApiDeploymentParameterMetadataResponseOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiDeploymentParameterMetadataResponse) *string { return v.Visibility }).(pulumi.StringPtrOutput)
}

type ApiDeploymentParameterMetadataResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiDeploymentParameterMetadataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDeploymentParameterMetadataResponse)(nil)).Elem()
}

func (o ApiDeploymentParameterMetadataResponsePtrOutput) ToApiDeploymentParameterMetadataResponsePtrOutput() ApiDeploymentParameterMetadataResponsePtrOutput {
	return o
}

func (o ApiDeploymentParameterMetadataResponsePtrOutput) ToApiDeploymentParameterMetadataResponsePtrOutputWithContext(ctx context.Context) ApiDeploymentParameterMetadataResponsePtrOutput {
	return o
}

func (o ApiDeploymentParameterMetadataResponsePtrOutput) Elem() ApiDeploymentParameterMetadataResponseOutput {
	return o.ApplyT(func(v *ApiDeploymentParameterMetadataResponse) ApiDeploymentParameterMetadataResponse {
		if v != nil {
			return *v
		}
		var ret ApiDeploymentParameterMetadataResponse
		return ret
	}).(ApiDeploymentParameterMetadataResponseOutput)
}

func (o ApiDeploymentParameterMetadataResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiDeploymentParameterMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ApiDeploymentParameterMetadataResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiDeploymentParameterMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ApiDeploymentParameterMetadataResponsePtrOutput) IsRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiDeploymentParameterMetadataResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsRequired
	}).(pulumi.BoolPtrOutput)
}

func (o ApiDeploymentParameterMetadataResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiDeploymentParameterMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ApiDeploymentParameterMetadataResponsePtrOutput) Visibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiDeploymentParameterMetadataResponse) *string {
		if v == nil {
			return nil
		}
		return v.Visibility
	}).(pulumi.StringPtrOutput)
}

type ApiDeploymentParameterMetadataSetResponse struct {
	PackageContentLink         *ApiDeploymentParameterMetadataResponse `pulumi:"packageContentLink"`
	RedisCacheConnectionString *ApiDeploymentParameterMetadataResponse `pulumi:"redisCacheConnectionString"`
}

type ApiDeploymentParameterMetadataSetResponseOutput struct{ *pulumi.OutputState }

func (ApiDeploymentParameterMetadataSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiDeploymentParameterMetadataSetResponse)(nil)).Elem()
}

func (o ApiDeploymentParameterMetadataSetResponseOutput) ToApiDeploymentParameterMetadataSetResponseOutput() ApiDeploymentParameterMetadataSetResponseOutput {
	return o
}

func (o ApiDeploymentParameterMetadataSetResponseOutput) ToApiDeploymentParameterMetadataSetResponseOutputWithContext(ctx context.Context) ApiDeploymentParameterMetadataSetResponseOutput {
	return o
}

func (o ApiDeploymentParameterMetadataSetResponseOutput) PackageContentLink() ApiDeploymentParameterMetadataResponsePtrOutput {
	return o.ApplyT(func(v ApiDeploymentParameterMetadataSetResponse) *ApiDeploymentParameterMetadataResponse {
		return v.PackageContentLink
	}).(ApiDeploymentParameterMetadataResponsePtrOutput)
}

func (o ApiDeploymentParameterMetadataSetResponseOutput) RedisCacheConnectionString() ApiDeploymentParameterMetadataResponsePtrOutput {
	return o.ApplyT(func(v ApiDeploymentParameterMetadataSetResponse) *ApiDeploymentParameterMetadataResponse {
		return v.RedisCacheConnectionString
	}).(ApiDeploymentParameterMetadataResponsePtrOutput)
}

type ApiDeploymentParameterMetadataSetResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiDeploymentParameterMetadataSetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDeploymentParameterMetadataSetResponse)(nil)).Elem()
}

func (o ApiDeploymentParameterMetadataSetResponsePtrOutput) ToApiDeploymentParameterMetadataSetResponsePtrOutput() ApiDeploymentParameterMetadataSetResponsePtrOutput {
	return o
}

func (o ApiDeploymentParameterMetadataSetResponsePtrOutput) ToApiDeploymentParameterMetadataSetResponsePtrOutputWithContext(ctx context.Context) ApiDeploymentParameterMetadataSetResponsePtrOutput {
	return o
}

func (o ApiDeploymentParameterMetadataSetResponsePtrOutput) Elem() ApiDeploymentParameterMetadataSetResponseOutput {
	return o.ApplyT(func(v *ApiDeploymentParameterMetadataSetResponse) ApiDeploymentParameterMetadataSetResponse {
		if v != nil {
			return *v
		}
		var ret ApiDeploymentParameterMetadataSetResponse
		return ret
	}).(ApiDeploymentParameterMetadataSetResponseOutput)
}

func (o ApiDeploymentParameterMetadataSetResponsePtrOutput) PackageContentLink() ApiDeploymentParameterMetadataResponsePtrOutput {
	return o.ApplyT(func(v *ApiDeploymentParameterMetadataSetResponse) *ApiDeploymentParameterMetadataResponse {
		if v == nil {
			return nil
		}
		return v.PackageContentLink
	}).(ApiDeploymentParameterMetadataResponsePtrOutput)
}

func (o ApiDeploymentParameterMetadataSetResponsePtrOutput) RedisCacheConnectionString() ApiDeploymentParameterMetadataResponsePtrOutput {
	return o.ApplyT(func(v *ApiDeploymentParameterMetadataSetResponse) *ApiDeploymentParameterMetadataResponse {
		if v == nil {
			return nil
		}
		return v.RedisCacheConnectionString
	}).(ApiDeploymentParameterMetadataResponsePtrOutput)
}

type ApiResourceBackendServiceResponse struct {
	ServiceUrl *string `pulumi:"serviceUrl"`
}

type ApiResourceBackendServiceResponseOutput struct{ *pulumi.OutputState }

func (ApiResourceBackendServiceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiResourceBackendServiceResponse)(nil)).Elem()
}

func (o ApiResourceBackendServiceResponseOutput) ToApiResourceBackendServiceResponseOutput() ApiResourceBackendServiceResponseOutput {
	return o
}

func (o ApiResourceBackendServiceResponseOutput) ToApiResourceBackendServiceResponseOutputWithContext(ctx context.Context) ApiResourceBackendServiceResponseOutput {
	return o
}

func (o ApiResourceBackendServiceResponseOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceBackendServiceResponse) *string { return v.ServiceUrl }).(pulumi.StringPtrOutput)
}

type ApiResourceDefinitionsResponse struct {
	ModifiedSwaggerUrl *string `pulumi:"modifiedSwaggerUrl"`
	OriginalSwaggerUrl *string `pulumi:"originalSwaggerUrl"`
}

type ApiResourceDefinitionsResponseOutput struct{ *pulumi.OutputState }

func (ApiResourceDefinitionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiResourceDefinitionsResponse)(nil)).Elem()
}

func (o ApiResourceDefinitionsResponseOutput) ToApiResourceDefinitionsResponseOutput() ApiResourceDefinitionsResponseOutput {
	return o
}

func (o ApiResourceDefinitionsResponseOutput) ToApiResourceDefinitionsResponseOutputWithContext(ctx context.Context) ApiResourceDefinitionsResponseOutput {
	return o
}

func (o ApiResourceDefinitionsResponseOutput) ModifiedSwaggerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceDefinitionsResponse) *string { return v.ModifiedSwaggerUrl }).(pulumi.StringPtrOutput)
}

func (o ApiResourceDefinitionsResponseOutput) OriginalSwaggerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceDefinitionsResponse) *string { return v.OriginalSwaggerUrl }).(pulumi.StringPtrOutput)
}

type ApiResourceGeneralInformationResponse struct {
	Description   *string `pulumi:"description"`
	DisplayName   *string `pulumi:"displayName"`
	IconUrl       *string `pulumi:"iconUrl"`
	ReleaseTag    *string `pulumi:"releaseTag"`
	TermsOfUseUrl *string `pulumi:"termsOfUseUrl"`
	Tier          *string `pulumi:"tier"`
}

type ApiResourceGeneralInformationResponseOutput struct{ *pulumi.OutputState }

func (ApiResourceGeneralInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiResourceGeneralInformationResponse)(nil)).Elem()
}

func (o ApiResourceGeneralInformationResponseOutput) ToApiResourceGeneralInformationResponseOutput() ApiResourceGeneralInformationResponseOutput {
	return o
}

func (o ApiResourceGeneralInformationResponseOutput) ToApiResourceGeneralInformationResponseOutputWithContext(ctx context.Context) ApiResourceGeneralInformationResponseOutput {
	return o
}

func (o ApiResourceGeneralInformationResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceGeneralInformationResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ApiResourceGeneralInformationResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceGeneralInformationResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ApiResourceGeneralInformationResponseOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceGeneralInformationResponse) *string { return v.IconUrl }).(pulumi.StringPtrOutput)
}

func (o ApiResourceGeneralInformationResponseOutput) ReleaseTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceGeneralInformationResponse) *string { return v.ReleaseTag }).(pulumi.StringPtrOutput)
}

func (o ApiResourceGeneralInformationResponseOutput) TermsOfUseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceGeneralInformationResponse) *string { return v.TermsOfUseUrl }).(pulumi.StringPtrOutput)
}

func (o ApiResourceGeneralInformationResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceGeneralInformationResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ApiResourceMetadataResponse struct {
	ApiType              *string                                    `pulumi:"apiType"`
	BrandColor           *string                                    `pulumi:"brandColor"`
	ConnectionType       *string                                    `pulumi:"connectionType"`
	DeploymentParameters *ApiDeploymentParameterMetadataSetResponse `pulumi:"deploymentParameters"`
	HideKey              *string                                    `pulumi:"hideKey"`
	ProvisioningState    *string                                    `pulumi:"provisioningState"`
	Source               *string                                    `pulumi:"source"`
	Tags                 map[string]string                          `pulumi:"tags"`
	WsdlImportMethod     *string                                    `pulumi:"wsdlImportMethod"`
	WsdlService          *WsdlServiceResponse                       `pulumi:"wsdlService"`
}

type ApiResourceMetadataResponseOutput struct{ *pulumi.OutputState }

func (ApiResourceMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiResourceMetadataResponse)(nil)).Elem()
}

func (o ApiResourceMetadataResponseOutput) ToApiResourceMetadataResponseOutput() ApiResourceMetadataResponseOutput {
	return o
}

func (o ApiResourceMetadataResponseOutput) ToApiResourceMetadataResponseOutputWithContext(ctx context.Context) ApiResourceMetadataResponseOutput {
	return o
}

func (o ApiResourceMetadataResponseOutput) ApiType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceMetadataResponse) *string { return v.ApiType }).(pulumi.StringPtrOutput)
}

func (o ApiResourceMetadataResponseOutput) BrandColor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceMetadataResponse) *string { return v.BrandColor }).(pulumi.StringPtrOutput)
}

func (o ApiResourceMetadataResponseOutput) ConnectionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceMetadataResponse) *string { return v.ConnectionType }).(pulumi.StringPtrOutput)
}

func (o ApiResourceMetadataResponseOutput) DeploymentParameters() ApiDeploymentParameterMetadataSetResponsePtrOutput {
	return o.ApplyT(func(v ApiResourceMetadataResponse) *ApiDeploymentParameterMetadataSetResponse {
		return v.DeploymentParameters
	}).(ApiDeploymentParameterMetadataSetResponsePtrOutput)
}

func (o ApiResourceMetadataResponseOutput) HideKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceMetadataResponse) *string { return v.HideKey }).(pulumi.StringPtrOutput)
}

func (o ApiResourceMetadataResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceMetadataResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApiResourceMetadataResponseOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceMetadataResponse) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func (o ApiResourceMetadataResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ApiResourceMetadataResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApiResourceMetadataResponseOutput) WsdlImportMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourceMetadataResponse) *string { return v.WsdlImportMethod }).(pulumi.StringPtrOutput)
}

func (o ApiResourceMetadataResponseOutput) WsdlService() WsdlServiceResponsePtrOutput {
	return o.ApplyT(func(v ApiResourceMetadataResponse) *WsdlServiceResponse { return v.WsdlService }).(WsdlServiceResponsePtrOutput)
}

type ApiResourcePoliciesResponse struct {
	Content     *string `pulumi:"content"`
	ContentLink *string `pulumi:"contentLink"`
}

type ApiResourcePoliciesResponseOutput struct{ *pulumi.OutputState }

func (ApiResourcePoliciesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiResourcePoliciesResponse)(nil)).Elem()
}

func (o ApiResourcePoliciesResponseOutput) ToApiResourcePoliciesResponseOutput() ApiResourcePoliciesResponseOutput {
	return o
}

func (o ApiResourcePoliciesResponseOutput) ToApiResourcePoliciesResponseOutputWithContext(ctx context.Context) ApiResourcePoliciesResponseOutput {
	return o
}

func (o ApiResourcePoliciesResponseOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourcePoliciesResponse) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o ApiResourcePoliciesResponseOutput) ContentLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiResourcePoliciesResponse) *string { return v.ContentLink }).(pulumi.StringPtrOutput)
}

type AssemblyProperties struct {
	AssemblyCulture        *string      `pulumi:"assemblyCulture"`
	AssemblyName           string       `pulumi:"assemblyName"`
	AssemblyPublicKeyToken *string      `pulumi:"assemblyPublicKeyToken"`
	AssemblyVersion        *string      `pulumi:"assemblyVersion"`
	ChangedTime            *string      `pulumi:"changedTime"`
	Content                interface{}  `pulumi:"content"`
	ContentLink            *ContentLink `pulumi:"contentLink"`
	ContentType            *string      `pulumi:"contentType"`
	CreatedTime            *string      `pulumi:"createdTime"`
	Metadata               interface{}  `pulumi:"metadata"`
}





type AssemblyPropertiesInput interface {
	pulumi.Input

	ToAssemblyPropertiesOutput() AssemblyPropertiesOutput
	ToAssemblyPropertiesOutputWithContext(context.Context) AssemblyPropertiesOutput
}

type AssemblyPropertiesArgs struct {
	AssemblyCulture        pulumi.StringPtrInput `pulumi:"assemblyCulture"`
	AssemblyName           pulumi.StringInput    `pulumi:"assemblyName"`
	AssemblyPublicKeyToken pulumi.StringPtrInput `pulumi:"assemblyPublicKeyToken"`
	AssemblyVersion        pulumi.StringPtrInput `pulumi:"assemblyVersion"`
	ChangedTime            pulumi.StringPtrInput `pulumi:"changedTime"`
	Content                pulumi.Input          `pulumi:"content"`
	ContentLink            ContentLinkPtrInput   `pulumi:"contentLink"`
	ContentType            pulumi.StringPtrInput `pulumi:"contentType"`
	CreatedTime            pulumi.StringPtrInput `pulumi:"createdTime"`
	Metadata               pulumi.Input          `pulumi:"metadata"`
}

func (AssemblyPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssemblyProperties)(nil)).Elem()
}

func (i AssemblyPropertiesArgs) ToAssemblyPropertiesOutput() AssemblyPropertiesOutput {
	return i.ToAssemblyPropertiesOutputWithContext(context.Background())
}

func (i AssemblyPropertiesArgs) ToAssemblyPropertiesOutputWithContext(ctx context.Context) AssemblyPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssemblyPropertiesOutput)
}

type AssemblyPropertiesOutput struct{ *pulumi.OutputState }

func (AssemblyPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssemblyProperties)(nil)).Elem()
}

func (o AssemblyPropertiesOutput) ToAssemblyPropertiesOutput() AssemblyPropertiesOutput {
	return o
}

func (o AssemblyPropertiesOutput) ToAssemblyPropertiesOutputWithContext(ctx context.Context) AssemblyPropertiesOutput {
	return o
}

func (o AssemblyPropertiesOutput) AssemblyCulture() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssemblyProperties) *string { return v.AssemblyCulture }).(pulumi.StringPtrOutput)
}

func (o AssemblyPropertiesOutput) AssemblyName() pulumi.StringOutput {
	return o.ApplyT(func(v AssemblyProperties) string { return v.AssemblyName }).(pulumi.StringOutput)
}

func (o AssemblyPropertiesOutput) AssemblyPublicKeyToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssemblyProperties) *string { return v.AssemblyPublicKeyToken }).(pulumi.StringPtrOutput)
}

func (o AssemblyPropertiesOutput) AssemblyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssemblyProperties) *string { return v.AssemblyVersion }).(pulumi.StringPtrOutput)
}

func (o AssemblyPropertiesOutput) ChangedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssemblyProperties) *string { return v.ChangedTime }).(pulumi.StringPtrOutput)
}

func (o AssemblyPropertiesOutput) Content() pulumi.AnyOutput {
	return o.ApplyT(func(v AssemblyProperties) interface{} { return v.Content }).(pulumi.AnyOutput)
}

func (o AssemblyPropertiesOutput) ContentLink() ContentLinkPtrOutput {
	return o.ApplyT(func(v AssemblyProperties) *ContentLink { return v.ContentLink }).(ContentLinkPtrOutput)
}

func (o AssemblyPropertiesOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssemblyProperties) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o AssemblyPropertiesOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssemblyProperties) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o AssemblyPropertiesOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v AssemblyProperties) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

type AssemblyPropertiesResponse struct {
	AssemblyCulture        *string              `pulumi:"assemblyCulture"`
	AssemblyName           string               `pulumi:"assemblyName"`
	AssemblyPublicKeyToken *string              `pulumi:"assemblyPublicKeyToken"`
	AssemblyVersion        *string              `pulumi:"assemblyVersion"`
	ChangedTime            *string              `pulumi:"changedTime"`
	Content                interface{}          `pulumi:"content"`
	ContentLink            *ContentLinkResponse `pulumi:"contentLink"`
	ContentType            *string              `pulumi:"contentType"`
	CreatedTime            *string              `pulumi:"createdTime"`
	Metadata               interface{}          `pulumi:"metadata"`
}

type AssemblyPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AssemblyPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssemblyPropertiesResponse)(nil)).Elem()
}

func (o AssemblyPropertiesResponseOutput) ToAssemblyPropertiesResponseOutput() AssemblyPropertiesResponseOutput {
	return o
}

func (o AssemblyPropertiesResponseOutput) ToAssemblyPropertiesResponseOutputWithContext(ctx context.Context) AssemblyPropertiesResponseOutput {
	return o
}

func (o AssemblyPropertiesResponseOutput) AssemblyCulture() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssemblyPropertiesResponse) *string { return v.AssemblyCulture }).(pulumi.StringPtrOutput)
}

func (o AssemblyPropertiesResponseOutput) AssemblyName() pulumi.StringOutput {
	return o.ApplyT(func(v AssemblyPropertiesResponse) string { return v.AssemblyName }).(pulumi.StringOutput)
}

func (o AssemblyPropertiesResponseOutput) AssemblyPublicKeyToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssemblyPropertiesResponse) *string { return v.AssemblyPublicKeyToken }).(pulumi.StringPtrOutput)
}

func (o AssemblyPropertiesResponseOutput) AssemblyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssemblyPropertiesResponse) *string { return v.AssemblyVersion }).(pulumi.StringPtrOutput)
}

func (o AssemblyPropertiesResponseOutput) ChangedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssemblyPropertiesResponse) *string { return v.ChangedTime }).(pulumi.StringPtrOutput)
}

func (o AssemblyPropertiesResponseOutput) Content() pulumi.AnyOutput {
	return o.ApplyT(func(v AssemblyPropertiesResponse) interface{} { return v.Content }).(pulumi.AnyOutput)
}

func (o AssemblyPropertiesResponseOutput) ContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v AssemblyPropertiesResponse) *ContentLinkResponse { return v.ContentLink }).(ContentLinkResponsePtrOutput)
}

func (o AssemblyPropertiesResponseOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssemblyPropertiesResponse) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o AssemblyPropertiesResponseOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssemblyPropertiesResponse) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o AssemblyPropertiesResponseOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v AssemblyPropertiesResponse) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

type AzureResourceErrorInfoResponse struct {
	Code    string                           `pulumi:"code"`
	Details []AzureResourceErrorInfoResponse `pulumi:"details"`
	Message string                           `pulumi:"message"`
}

type AzureResourceErrorInfoResponseOutput struct{ *pulumi.OutputState }

func (AzureResourceErrorInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceErrorInfoResponse)(nil)).Elem()
}

func (o AzureResourceErrorInfoResponseOutput) ToAzureResourceErrorInfoResponseOutput() AzureResourceErrorInfoResponseOutput {
	return o
}

func (o AzureResourceErrorInfoResponseOutput) ToAzureResourceErrorInfoResponseOutputWithContext(ctx context.Context) AzureResourceErrorInfoResponseOutput {
	return o
}

func (o AzureResourceErrorInfoResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceErrorInfoResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o AzureResourceErrorInfoResponseOutput) Details() AzureResourceErrorInfoResponseArrayOutput {
	return o.ApplyT(func(v AzureResourceErrorInfoResponse) []AzureResourceErrorInfoResponse { return v.Details }).(AzureResourceErrorInfoResponseArrayOutput)
}

func (o AzureResourceErrorInfoResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceErrorInfoResponse) string { return v.Message }).(pulumi.StringOutput)
}

type AzureResourceErrorInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureResourceErrorInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureResourceErrorInfoResponse)(nil)).Elem()
}

func (o AzureResourceErrorInfoResponsePtrOutput) ToAzureResourceErrorInfoResponsePtrOutput() AzureResourceErrorInfoResponsePtrOutput {
	return o
}

func (o AzureResourceErrorInfoResponsePtrOutput) ToAzureResourceErrorInfoResponsePtrOutputWithContext(ctx context.Context) AzureResourceErrorInfoResponsePtrOutput {
	return o
}

func (o AzureResourceErrorInfoResponsePtrOutput) Elem() AzureResourceErrorInfoResponseOutput {
	return o.ApplyT(func(v *AzureResourceErrorInfoResponse) AzureResourceErrorInfoResponse {
		if v != nil {
			return *v
		}
		var ret AzureResourceErrorInfoResponse
		return ret
	}).(AzureResourceErrorInfoResponseOutput)
}

func (o AzureResourceErrorInfoResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureResourceErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o AzureResourceErrorInfoResponsePtrOutput) Details() AzureResourceErrorInfoResponseArrayOutput {
	return o.ApplyT(func(v *AzureResourceErrorInfoResponse) []AzureResourceErrorInfoResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(AzureResourceErrorInfoResponseArrayOutput)
}

func (o AzureResourceErrorInfoResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureResourceErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

type AzureResourceErrorInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureResourceErrorInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureResourceErrorInfoResponse)(nil)).Elem()
}

func (o AzureResourceErrorInfoResponseArrayOutput) ToAzureResourceErrorInfoResponseArrayOutput() AzureResourceErrorInfoResponseArrayOutput {
	return o
}

func (o AzureResourceErrorInfoResponseArrayOutput) ToAzureResourceErrorInfoResponseArrayOutputWithContext(ctx context.Context) AzureResourceErrorInfoResponseArrayOutput {
	return o
}

func (o AzureResourceErrorInfoResponseArrayOutput) Index(i pulumi.IntInput) AzureResourceErrorInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureResourceErrorInfoResponse {
		return vs[0].([]AzureResourceErrorInfoResponse)[vs[1].(int)]
	}).(AzureResourceErrorInfoResponseOutput)
}

type B2BPartnerContent struct {
	BusinessIdentities []BusinessIdentity `pulumi:"businessIdentities"`
}





type B2BPartnerContentInput interface {
	pulumi.Input

	ToB2BPartnerContentOutput() B2BPartnerContentOutput
	ToB2BPartnerContentOutputWithContext(context.Context) B2BPartnerContentOutput
}

type B2BPartnerContentArgs struct {
	BusinessIdentities BusinessIdentityArrayInput `pulumi:"businessIdentities"`
}

func (B2BPartnerContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*B2BPartnerContent)(nil)).Elem()
}

func (i B2BPartnerContentArgs) ToB2BPartnerContentOutput() B2BPartnerContentOutput {
	return i.ToB2BPartnerContentOutputWithContext(context.Background())
}

func (i B2BPartnerContentArgs) ToB2BPartnerContentOutputWithContext(ctx context.Context) B2BPartnerContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(B2BPartnerContentOutput)
}

func (i B2BPartnerContentArgs) ToB2BPartnerContentPtrOutput() B2BPartnerContentPtrOutput {
	return i.ToB2BPartnerContentPtrOutputWithContext(context.Background())
}

func (i B2BPartnerContentArgs) ToB2BPartnerContentPtrOutputWithContext(ctx context.Context) B2BPartnerContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(B2BPartnerContentOutput).ToB2BPartnerContentPtrOutputWithContext(ctx)
}









type B2BPartnerContentPtrInput interface {
	pulumi.Input

	ToB2BPartnerContentPtrOutput() B2BPartnerContentPtrOutput
	ToB2BPartnerContentPtrOutputWithContext(context.Context) B2BPartnerContentPtrOutput
}

type b2bpartnerContentPtrType B2BPartnerContentArgs

func B2BPartnerContentPtr(v *B2BPartnerContentArgs) B2BPartnerContentPtrInput {
	return (*b2bpartnerContentPtrType)(v)
}

func (*b2bpartnerContentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**B2BPartnerContent)(nil)).Elem()
}

func (i *b2bpartnerContentPtrType) ToB2BPartnerContentPtrOutput() B2BPartnerContentPtrOutput {
	return i.ToB2BPartnerContentPtrOutputWithContext(context.Background())
}

func (i *b2bpartnerContentPtrType) ToB2BPartnerContentPtrOutputWithContext(ctx context.Context) B2BPartnerContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(B2BPartnerContentPtrOutput)
}

type B2BPartnerContentOutput struct{ *pulumi.OutputState }

func (B2BPartnerContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*B2BPartnerContent)(nil)).Elem()
}

func (o B2BPartnerContentOutput) ToB2BPartnerContentOutput() B2BPartnerContentOutput {
	return o
}

func (o B2BPartnerContentOutput) ToB2BPartnerContentOutputWithContext(ctx context.Context) B2BPartnerContentOutput {
	return o
}

func (o B2BPartnerContentOutput) ToB2BPartnerContentPtrOutput() B2BPartnerContentPtrOutput {
	return o.ToB2BPartnerContentPtrOutputWithContext(context.Background())
}

func (o B2BPartnerContentOutput) ToB2BPartnerContentPtrOutputWithContext(ctx context.Context) B2BPartnerContentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v B2BPartnerContent) *B2BPartnerContent {
		return &v
	}).(B2BPartnerContentPtrOutput)
}

func (o B2BPartnerContentOutput) BusinessIdentities() BusinessIdentityArrayOutput {
	return o.ApplyT(func(v B2BPartnerContent) []BusinessIdentity { return v.BusinessIdentities }).(BusinessIdentityArrayOutput)
}

type B2BPartnerContentPtrOutput struct{ *pulumi.OutputState }

func (B2BPartnerContentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**B2BPartnerContent)(nil)).Elem()
}

func (o B2BPartnerContentPtrOutput) ToB2BPartnerContentPtrOutput() B2BPartnerContentPtrOutput {
	return o
}

func (o B2BPartnerContentPtrOutput) ToB2BPartnerContentPtrOutputWithContext(ctx context.Context) B2BPartnerContentPtrOutput {
	return o
}

func (o B2BPartnerContentPtrOutput) Elem() B2BPartnerContentOutput {
	return o.ApplyT(func(v *B2BPartnerContent) B2BPartnerContent {
		if v != nil {
			return *v
		}
		var ret B2BPartnerContent
		return ret
	}).(B2BPartnerContentOutput)
}

func (o B2BPartnerContentPtrOutput) BusinessIdentities() BusinessIdentityArrayOutput {
	return o.ApplyT(func(v *B2BPartnerContent) []BusinessIdentity {
		if v == nil {
			return nil
		}
		return v.BusinessIdentities
	}).(BusinessIdentityArrayOutput)
}

type B2BPartnerContentResponse struct {
	BusinessIdentities []BusinessIdentityResponse `pulumi:"businessIdentities"`
}

type B2BPartnerContentResponseOutput struct{ *pulumi.OutputState }

func (B2BPartnerContentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*B2BPartnerContentResponse)(nil)).Elem()
}

func (o B2BPartnerContentResponseOutput) ToB2BPartnerContentResponseOutput() B2BPartnerContentResponseOutput {
	return o
}

func (o B2BPartnerContentResponseOutput) ToB2BPartnerContentResponseOutputWithContext(ctx context.Context) B2BPartnerContentResponseOutput {
	return o
}

func (o B2BPartnerContentResponseOutput) BusinessIdentities() BusinessIdentityResponseArrayOutput {
	return o.ApplyT(func(v B2BPartnerContentResponse) []BusinessIdentityResponse { return v.BusinessIdentities }).(BusinessIdentityResponseArrayOutput)
}

type B2BPartnerContentResponsePtrOutput struct{ *pulumi.OutputState }

func (B2BPartnerContentResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**B2BPartnerContentResponse)(nil)).Elem()
}

func (o B2BPartnerContentResponsePtrOutput) ToB2BPartnerContentResponsePtrOutput() B2BPartnerContentResponsePtrOutput {
	return o
}

func (o B2BPartnerContentResponsePtrOutput) ToB2BPartnerContentResponsePtrOutputWithContext(ctx context.Context) B2BPartnerContentResponsePtrOutput {
	return o
}

func (o B2BPartnerContentResponsePtrOutput) Elem() B2BPartnerContentResponseOutput {
	return o.ApplyT(func(v *B2BPartnerContentResponse) B2BPartnerContentResponse {
		if v != nil {
			return *v
		}
		var ret B2BPartnerContentResponse
		return ret
	}).(B2BPartnerContentResponseOutput)
}

func (o B2BPartnerContentResponsePtrOutput) BusinessIdentities() BusinessIdentityResponseArrayOutput {
	return o.ApplyT(func(v *B2BPartnerContentResponse) []BusinessIdentityResponse {
		if v == nil {
			return nil
		}
		return v.BusinessIdentities
	}).(BusinessIdentityResponseArrayOutput)
}

type BatchConfigurationProperties struct {
	BatchGroupName  string               `pulumi:"batchGroupName"`
	ChangedTime     *string              `pulumi:"changedTime"`
	CreatedTime     *string              `pulumi:"createdTime"`
	Metadata        interface{}          `pulumi:"metadata"`
	ReleaseCriteria BatchReleaseCriteria `pulumi:"releaseCriteria"`
}





type BatchConfigurationPropertiesInput interface {
	pulumi.Input

	ToBatchConfigurationPropertiesOutput() BatchConfigurationPropertiesOutput
	ToBatchConfigurationPropertiesOutputWithContext(context.Context) BatchConfigurationPropertiesOutput
}

type BatchConfigurationPropertiesArgs struct {
	BatchGroupName  pulumi.StringInput        `pulumi:"batchGroupName"`
	ChangedTime     pulumi.StringPtrInput     `pulumi:"changedTime"`
	CreatedTime     pulumi.StringPtrInput     `pulumi:"createdTime"`
	Metadata        pulumi.Input              `pulumi:"metadata"`
	ReleaseCriteria BatchReleaseCriteriaInput `pulumi:"releaseCriteria"`
}

func (BatchConfigurationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchConfigurationProperties)(nil)).Elem()
}

func (i BatchConfigurationPropertiesArgs) ToBatchConfigurationPropertiesOutput() BatchConfigurationPropertiesOutput {
	return i.ToBatchConfigurationPropertiesOutputWithContext(context.Background())
}

func (i BatchConfigurationPropertiesArgs) ToBatchConfigurationPropertiesOutputWithContext(ctx context.Context) BatchConfigurationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchConfigurationPropertiesOutput)
}

type BatchConfigurationPropertiesOutput struct{ *pulumi.OutputState }

func (BatchConfigurationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchConfigurationProperties)(nil)).Elem()
}

func (o BatchConfigurationPropertiesOutput) ToBatchConfigurationPropertiesOutput() BatchConfigurationPropertiesOutput {
	return o
}

func (o BatchConfigurationPropertiesOutput) ToBatchConfigurationPropertiesOutputWithContext(ctx context.Context) BatchConfigurationPropertiesOutput {
	return o
}

func (o BatchConfigurationPropertiesOutput) BatchGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v BatchConfigurationProperties) string { return v.BatchGroupName }).(pulumi.StringOutput)
}

func (o BatchConfigurationPropertiesOutput) ChangedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchConfigurationProperties) *string { return v.ChangedTime }).(pulumi.StringPtrOutput)
}

func (o BatchConfigurationPropertiesOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchConfigurationProperties) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o BatchConfigurationPropertiesOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v BatchConfigurationProperties) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o BatchConfigurationPropertiesOutput) ReleaseCriteria() BatchReleaseCriteriaOutput {
	return o.ApplyT(func(v BatchConfigurationProperties) BatchReleaseCriteria { return v.ReleaseCriteria }).(BatchReleaseCriteriaOutput)
}

type BatchConfigurationPropertiesResponse struct {
	BatchGroupName  string                       `pulumi:"batchGroupName"`
	ChangedTime     *string                      `pulumi:"changedTime"`
	CreatedTime     *string                      `pulumi:"createdTime"`
	Metadata        interface{}                  `pulumi:"metadata"`
	ReleaseCriteria BatchReleaseCriteriaResponse `pulumi:"releaseCriteria"`
}

type BatchConfigurationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (BatchConfigurationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchConfigurationPropertiesResponse)(nil)).Elem()
}

func (o BatchConfigurationPropertiesResponseOutput) ToBatchConfigurationPropertiesResponseOutput() BatchConfigurationPropertiesResponseOutput {
	return o
}

func (o BatchConfigurationPropertiesResponseOutput) ToBatchConfigurationPropertiesResponseOutputWithContext(ctx context.Context) BatchConfigurationPropertiesResponseOutput {
	return o
}

func (o BatchConfigurationPropertiesResponseOutput) BatchGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v BatchConfigurationPropertiesResponse) string { return v.BatchGroupName }).(pulumi.StringOutput)
}

func (o BatchConfigurationPropertiesResponseOutput) ChangedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchConfigurationPropertiesResponse) *string { return v.ChangedTime }).(pulumi.StringPtrOutput)
}

func (o BatchConfigurationPropertiesResponseOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BatchConfigurationPropertiesResponse) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o BatchConfigurationPropertiesResponseOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v BatchConfigurationPropertiesResponse) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o BatchConfigurationPropertiesResponseOutput) ReleaseCriteria() BatchReleaseCriteriaResponseOutput {
	return o.ApplyT(func(v BatchConfigurationPropertiesResponse) BatchReleaseCriteriaResponse { return v.ReleaseCriteria }).(BatchReleaseCriteriaResponseOutput)
}

type BatchReleaseCriteria struct {
	BatchSize    *int                       `pulumi:"batchSize"`
	MessageCount *int                       `pulumi:"messageCount"`
	Recurrence   *WorkflowTriggerRecurrence `pulumi:"recurrence"`
}





type BatchReleaseCriteriaInput interface {
	pulumi.Input

	ToBatchReleaseCriteriaOutput() BatchReleaseCriteriaOutput
	ToBatchReleaseCriteriaOutputWithContext(context.Context) BatchReleaseCriteriaOutput
}

type BatchReleaseCriteriaArgs struct {
	BatchSize    pulumi.IntPtrInput                `pulumi:"batchSize"`
	MessageCount pulumi.IntPtrInput                `pulumi:"messageCount"`
	Recurrence   WorkflowTriggerRecurrencePtrInput `pulumi:"recurrence"`
}

func (BatchReleaseCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchReleaseCriteria)(nil)).Elem()
}

func (i BatchReleaseCriteriaArgs) ToBatchReleaseCriteriaOutput() BatchReleaseCriteriaOutput {
	return i.ToBatchReleaseCriteriaOutputWithContext(context.Background())
}

func (i BatchReleaseCriteriaArgs) ToBatchReleaseCriteriaOutputWithContext(ctx context.Context) BatchReleaseCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BatchReleaseCriteriaOutput)
}

type BatchReleaseCriteriaOutput struct{ *pulumi.OutputState }

func (BatchReleaseCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchReleaseCriteria)(nil)).Elem()
}

func (o BatchReleaseCriteriaOutput) ToBatchReleaseCriteriaOutput() BatchReleaseCriteriaOutput {
	return o
}

func (o BatchReleaseCriteriaOutput) ToBatchReleaseCriteriaOutputWithContext(ctx context.Context) BatchReleaseCriteriaOutput {
	return o
}

func (o BatchReleaseCriteriaOutput) BatchSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BatchReleaseCriteria) *int { return v.BatchSize }).(pulumi.IntPtrOutput)
}

func (o BatchReleaseCriteriaOutput) MessageCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BatchReleaseCriteria) *int { return v.MessageCount }).(pulumi.IntPtrOutput)
}

func (o BatchReleaseCriteriaOutput) Recurrence() WorkflowTriggerRecurrencePtrOutput {
	return o.ApplyT(func(v BatchReleaseCriteria) *WorkflowTriggerRecurrence { return v.Recurrence }).(WorkflowTriggerRecurrencePtrOutput)
}

type BatchReleaseCriteriaResponse struct {
	BatchSize    *int                               `pulumi:"batchSize"`
	MessageCount *int                               `pulumi:"messageCount"`
	Recurrence   *WorkflowTriggerRecurrenceResponse `pulumi:"recurrence"`
}

type BatchReleaseCriteriaResponseOutput struct{ *pulumi.OutputState }

func (BatchReleaseCriteriaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BatchReleaseCriteriaResponse)(nil)).Elem()
}

func (o BatchReleaseCriteriaResponseOutput) ToBatchReleaseCriteriaResponseOutput() BatchReleaseCriteriaResponseOutput {
	return o
}

func (o BatchReleaseCriteriaResponseOutput) ToBatchReleaseCriteriaResponseOutputWithContext(ctx context.Context) BatchReleaseCriteriaResponseOutput {
	return o
}

func (o BatchReleaseCriteriaResponseOutput) BatchSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BatchReleaseCriteriaResponse) *int { return v.BatchSize }).(pulumi.IntPtrOutput)
}

func (o BatchReleaseCriteriaResponseOutput) MessageCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BatchReleaseCriteriaResponse) *int { return v.MessageCount }).(pulumi.IntPtrOutput)
}

func (o BatchReleaseCriteriaResponseOutput) Recurrence() WorkflowTriggerRecurrenceResponsePtrOutput {
	return o.ApplyT(func(v BatchReleaseCriteriaResponse) *WorkflowTriggerRecurrenceResponse { return v.Recurrence }).(WorkflowTriggerRecurrenceResponsePtrOutput)
}

type BusinessIdentity struct {
	Qualifier string `pulumi:"qualifier"`
	Value     string `pulumi:"value"`
}





type BusinessIdentityInput interface {
	pulumi.Input

	ToBusinessIdentityOutput() BusinessIdentityOutput
	ToBusinessIdentityOutputWithContext(context.Context) BusinessIdentityOutput
}

type BusinessIdentityArgs struct {
	Qualifier pulumi.StringInput `pulumi:"qualifier"`
	Value     pulumi.StringInput `pulumi:"value"`
}

func (BusinessIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BusinessIdentity)(nil)).Elem()
}

func (i BusinessIdentityArgs) ToBusinessIdentityOutput() BusinessIdentityOutput {
	return i.ToBusinessIdentityOutputWithContext(context.Background())
}

func (i BusinessIdentityArgs) ToBusinessIdentityOutputWithContext(ctx context.Context) BusinessIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BusinessIdentityOutput)
}

func (i BusinessIdentityArgs) ToBusinessIdentityPtrOutput() BusinessIdentityPtrOutput {
	return i.ToBusinessIdentityPtrOutputWithContext(context.Background())
}

func (i BusinessIdentityArgs) ToBusinessIdentityPtrOutputWithContext(ctx context.Context) BusinessIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BusinessIdentityOutput).ToBusinessIdentityPtrOutputWithContext(ctx)
}









type BusinessIdentityPtrInput interface {
	pulumi.Input

	ToBusinessIdentityPtrOutput() BusinessIdentityPtrOutput
	ToBusinessIdentityPtrOutputWithContext(context.Context) BusinessIdentityPtrOutput
}

type businessIdentityPtrType BusinessIdentityArgs

func BusinessIdentityPtr(v *BusinessIdentityArgs) BusinessIdentityPtrInput {
	return (*businessIdentityPtrType)(v)
}

func (*businessIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BusinessIdentity)(nil)).Elem()
}

func (i *businessIdentityPtrType) ToBusinessIdentityPtrOutput() BusinessIdentityPtrOutput {
	return i.ToBusinessIdentityPtrOutputWithContext(context.Background())
}

func (i *businessIdentityPtrType) ToBusinessIdentityPtrOutputWithContext(ctx context.Context) BusinessIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BusinessIdentityPtrOutput)
}





type BusinessIdentityArrayInput interface {
	pulumi.Input

	ToBusinessIdentityArrayOutput() BusinessIdentityArrayOutput
	ToBusinessIdentityArrayOutputWithContext(context.Context) BusinessIdentityArrayOutput
}

type BusinessIdentityArray []BusinessIdentityInput

func (BusinessIdentityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BusinessIdentity)(nil)).Elem()
}

func (i BusinessIdentityArray) ToBusinessIdentityArrayOutput() BusinessIdentityArrayOutput {
	return i.ToBusinessIdentityArrayOutputWithContext(context.Background())
}

func (i BusinessIdentityArray) ToBusinessIdentityArrayOutputWithContext(ctx context.Context) BusinessIdentityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BusinessIdentityArrayOutput)
}

type BusinessIdentityOutput struct{ *pulumi.OutputState }

func (BusinessIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BusinessIdentity)(nil)).Elem()
}

func (o BusinessIdentityOutput) ToBusinessIdentityOutput() BusinessIdentityOutput {
	return o
}

func (o BusinessIdentityOutput) ToBusinessIdentityOutputWithContext(ctx context.Context) BusinessIdentityOutput {
	return o
}

func (o BusinessIdentityOutput) ToBusinessIdentityPtrOutput() BusinessIdentityPtrOutput {
	return o.ToBusinessIdentityPtrOutputWithContext(context.Background())
}

func (o BusinessIdentityOutput) ToBusinessIdentityPtrOutputWithContext(ctx context.Context) BusinessIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BusinessIdentity) *BusinessIdentity {
		return &v
	}).(BusinessIdentityPtrOutput)
}

func (o BusinessIdentityOutput) Qualifier() pulumi.StringOutput {
	return o.ApplyT(func(v BusinessIdentity) string { return v.Qualifier }).(pulumi.StringOutput)
}

func (o BusinessIdentityOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v BusinessIdentity) string { return v.Value }).(pulumi.StringOutput)
}

type BusinessIdentityPtrOutput struct{ *pulumi.OutputState }

func (BusinessIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BusinessIdentity)(nil)).Elem()
}

func (o BusinessIdentityPtrOutput) ToBusinessIdentityPtrOutput() BusinessIdentityPtrOutput {
	return o
}

func (o BusinessIdentityPtrOutput) ToBusinessIdentityPtrOutputWithContext(ctx context.Context) BusinessIdentityPtrOutput {
	return o
}

func (o BusinessIdentityPtrOutput) Elem() BusinessIdentityOutput {
	return o.ApplyT(func(v *BusinessIdentity) BusinessIdentity {
		if v != nil {
			return *v
		}
		var ret BusinessIdentity
		return ret
	}).(BusinessIdentityOutput)
}

func (o BusinessIdentityPtrOutput) Qualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BusinessIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Qualifier
	}).(pulumi.StringPtrOutput)
}

func (o BusinessIdentityPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BusinessIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type BusinessIdentityArrayOutput struct{ *pulumi.OutputState }

func (BusinessIdentityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BusinessIdentity)(nil)).Elem()
}

func (o BusinessIdentityArrayOutput) ToBusinessIdentityArrayOutput() BusinessIdentityArrayOutput {
	return o
}

func (o BusinessIdentityArrayOutput) ToBusinessIdentityArrayOutputWithContext(ctx context.Context) BusinessIdentityArrayOutput {
	return o
}

func (o BusinessIdentityArrayOutput) Index(i pulumi.IntInput) BusinessIdentityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BusinessIdentity {
		return vs[0].([]BusinessIdentity)[vs[1].(int)]
	}).(BusinessIdentityOutput)
}

type BusinessIdentityResponse struct {
	Qualifier string `pulumi:"qualifier"`
	Value     string `pulumi:"value"`
}

type BusinessIdentityResponseOutput struct{ *pulumi.OutputState }

func (BusinessIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BusinessIdentityResponse)(nil)).Elem()
}

func (o BusinessIdentityResponseOutput) ToBusinessIdentityResponseOutput() BusinessIdentityResponseOutput {
	return o
}

func (o BusinessIdentityResponseOutput) ToBusinessIdentityResponseOutputWithContext(ctx context.Context) BusinessIdentityResponseOutput {
	return o
}

func (o BusinessIdentityResponseOutput) Qualifier() pulumi.StringOutput {
	return o.ApplyT(func(v BusinessIdentityResponse) string { return v.Qualifier }).(pulumi.StringOutput)
}

func (o BusinessIdentityResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v BusinessIdentityResponse) string { return v.Value }).(pulumi.StringOutput)
}

type BusinessIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (BusinessIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BusinessIdentityResponse)(nil)).Elem()
}

func (o BusinessIdentityResponsePtrOutput) ToBusinessIdentityResponsePtrOutput() BusinessIdentityResponsePtrOutput {
	return o
}

func (o BusinessIdentityResponsePtrOutput) ToBusinessIdentityResponsePtrOutputWithContext(ctx context.Context) BusinessIdentityResponsePtrOutput {
	return o
}

func (o BusinessIdentityResponsePtrOutput) Elem() BusinessIdentityResponseOutput {
	return o.ApplyT(func(v *BusinessIdentityResponse) BusinessIdentityResponse {
		if v != nil {
			return *v
		}
		var ret BusinessIdentityResponse
		return ret
	}).(BusinessIdentityResponseOutput)
}

func (o BusinessIdentityResponsePtrOutput) Qualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BusinessIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Qualifier
	}).(pulumi.StringPtrOutput)
}

func (o BusinessIdentityResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BusinessIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type BusinessIdentityResponseArrayOutput struct{ *pulumi.OutputState }

func (BusinessIdentityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BusinessIdentityResponse)(nil)).Elem()
}

func (o BusinessIdentityResponseArrayOutput) ToBusinessIdentityResponseArrayOutput() BusinessIdentityResponseArrayOutput {
	return o
}

func (o BusinessIdentityResponseArrayOutput) ToBusinessIdentityResponseArrayOutputWithContext(ctx context.Context) BusinessIdentityResponseArrayOutput {
	return o
}

func (o BusinessIdentityResponseArrayOutput) Index(i pulumi.IntInput) BusinessIdentityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BusinessIdentityResponse {
		return vs[0].([]BusinessIdentityResponse)[vs[1].(int)]
	}).(BusinessIdentityResponseOutput)
}

type ContentHashResponse struct {
	Algorithm *string `pulumi:"algorithm"`
	Value     *string `pulumi:"value"`
}

type ContentHashResponseOutput struct{ *pulumi.OutputState }

func (ContentHashResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentHashResponse)(nil)).Elem()
}

func (o ContentHashResponseOutput) ToContentHashResponseOutput() ContentHashResponseOutput {
	return o
}

func (o ContentHashResponseOutput) ToContentHashResponseOutputWithContext(ctx context.Context) ContentHashResponseOutput {
	return o
}

func (o ContentHashResponseOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentHashResponse) *string { return v.Algorithm }).(pulumi.StringPtrOutput)
}

func (o ContentHashResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentHashResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ContentHashResponsePtrOutput struct{ *pulumi.OutputState }

func (ContentHashResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentHashResponse)(nil)).Elem()
}

func (o ContentHashResponsePtrOutput) ToContentHashResponsePtrOutput() ContentHashResponsePtrOutput {
	return o
}

func (o ContentHashResponsePtrOutput) ToContentHashResponsePtrOutputWithContext(ctx context.Context) ContentHashResponsePtrOutput {
	return o
}

func (o ContentHashResponsePtrOutput) Elem() ContentHashResponseOutput {
	return o.ApplyT(func(v *ContentHashResponse) ContentHashResponse {
		if v != nil {
			return *v
		}
		var ret ContentHashResponse
		return ret
	}).(ContentHashResponseOutput)
}

func (o ContentHashResponsePtrOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHashResponse) *string {
		if v == nil {
			return nil
		}
		return v.Algorithm
	}).(pulumi.StringPtrOutput)
}

func (o ContentHashResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHashResponse) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
}

type ContentLink struct {
	Uri *string `pulumi:"uri"`
}





type ContentLinkInput interface {
	pulumi.Input

	ToContentLinkOutput() ContentLinkOutput
	ToContentLinkOutputWithContext(context.Context) ContentLinkOutput
}

type ContentLinkArgs struct {
	Uri pulumi.StringPtrInput `pulumi:"uri"`
}

func (ContentLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentLink)(nil)).Elem()
}

func (i ContentLinkArgs) ToContentLinkOutput() ContentLinkOutput {
	return i.ToContentLinkOutputWithContext(context.Background())
}

func (i ContentLinkArgs) ToContentLinkOutputWithContext(ctx context.Context) ContentLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkOutput)
}

func (i ContentLinkArgs) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return i.ToContentLinkPtrOutputWithContext(context.Background())
}

func (i ContentLinkArgs) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkOutput).ToContentLinkPtrOutputWithContext(ctx)
}









type ContentLinkPtrInput interface {
	pulumi.Input

	ToContentLinkPtrOutput() ContentLinkPtrOutput
	ToContentLinkPtrOutputWithContext(context.Context) ContentLinkPtrOutput
}

type contentLinkPtrType ContentLinkArgs

func ContentLinkPtr(v *ContentLinkArgs) ContentLinkPtrInput {
	return (*contentLinkPtrType)(v)
}

func (*contentLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLink)(nil)).Elem()
}

func (i *contentLinkPtrType) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return i.ToContentLinkPtrOutputWithContext(context.Background())
}

func (i *contentLinkPtrType) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkPtrOutput)
}

type ContentLinkOutput struct{ *pulumi.OutputState }

func (ContentLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentLink)(nil)).Elem()
}

func (o ContentLinkOutput) ToContentLinkOutput() ContentLinkOutput {
	return o
}

func (o ContentLinkOutput) ToContentLinkOutputWithContext(ctx context.Context) ContentLinkOutput {
	return o
}

func (o ContentLinkOutput) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return o.ToContentLinkPtrOutputWithContext(context.Background())
}

func (o ContentLinkOutput) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentLink) *ContentLink {
		return &v
	}).(ContentLinkPtrOutput)
}

func (o ContentLinkOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentLink) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ContentLinkPtrOutput struct{ *pulumi.OutputState }

func (ContentLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLink)(nil)).Elem()
}

func (o ContentLinkPtrOutput) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return o
}

func (o ContentLinkPtrOutput) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return o
}

func (o ContentLinkPtrOutput) Elem() ContentLinkOutput {
	return o.ApplyT(func(v *ContentLink) ContentLink {
		if v != nil {
			return *v
		}
		var ret ContentLink
		return ret
	}).(ContentLinkOutput)
}

func (o ContentLinkPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLink) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type ContentLinkResponse struct {
	ContentHash    ContentHashResponse `pulumi:"contentHash"`
	ContentSize    float64             `pulumi:"contentSize"`
	ContentVersion string              `pulumi:"contentVersion"`
	Metadata       interface{}         `pulumi:"metadata"`
	Uri            *string             `pulumi:"uri"`
}

type ContentLinkResponseOutput struct{ *pulumi.OutputState }

func (ContentLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentLinkResponse)(nil)).Elem()
}

func (o ContentLinkResponseOutput) ToContentLinkResponseOutput() ContentLinkResponseOutput {
	return o
}

func (o ContentLinkResponseOutput) ToContentLinkResponseOutputWithContext(ctx context.Context) ContentLinkResponseOutput {
	return o
}

func (o ContentLinkResponseOutput) ContentHash() ContentHashResponseOutput {
	return o.ApplyT(func(v ContentLinkResponse) ContentHashResponse { return v.ContentHash }).(ContentHashResponseOutput)
}

func (o ContentLinkResponseOutput) ContentSize() pulumi.Float64Output {
	return o.ApplyT(func(v ContentLinkResponse) float64 { return v.ContentSize }).(pulumi.Float64Output)
}

func (o ContentLinkResponseOutput) ContentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ContentLinkResponse) string { return v.ContentVersion }).(pulumi.StringOutput)
}

func (o ContentLinkResponseOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v ContentLinkResponse) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o ContentLinkResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentLinkResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ContentLinkResponsePtrOutput struct{ *pulumi.OutputState }

func (ContentLinkResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLinkResponse)(nil)).Elem()
}

func (o ContentLinkResponsePtrOutput) ToContentLinkResponsePtrOutput() ContentLinkResponsePtrOutput {
	return o
}

func (o ContentLinkResponsePtrOutput) ToContentLinkResponsePtrOutputWithContext(ctx context.Context) ContentLinkResponsePtrOutput {
	return o
}

func (o ContentLinkResponsePtrOutput) Elem() ContentLinkResponseOutput {
	return o.ApplyT(func(v *ContentLinkResponse) ContentLinkResponse {
		if v != nil {
			return *v
		}
		var ret ContentLinkResponse
		return ret
	}).(ContentLinkResponseOutput)
}

func (o ContentLinkResponsePtrOutput) ContentHash() ContentHashResponsePtrOutput {
	return o.ApplyT(func(v *ContentLinkResponse) *ContentHashResponse {
		if v == nil {
			return nil
		}
		return &v.ContentHash
	}).(ContentHashResponsePtrOutput)
}

func (o ContentLinkResponsePtrOutput) ContentSize() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ContentLinkResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.ContentSize
	}).(pulumi.Float64PtrOutput)
}

func (o ContentLinkResponsePtrOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLinkResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ContentVersion
	}).(pulumi.StringPtrOutput)
}

func (o ContentLinkResponsePtrOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *ContentLinkResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.AnyOutput)
}

func (o ContentLinkResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLinkResponse) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type EdifactAcknowledgementSettings struct {
	AcknowledgementControlNumberLowerBound int     `pulumi:"acknowledgementControlNumberLowerBound"`
	AcknowledgementControlNumberPrefix     *string `pulumi:"acknowledgementControlNumberPrefix"`
	AcknowledgementControlNumberSuffix     *string `pulumi:"acknowledgementControlNumberSuffix"`
	AcknowledgementControlNumberUpperBound int     `pulumi:"acknowledgementControlNumberUpperBound"`
	BatchFunctionalAcknowledgements        bool    `pulumi:"batchFunctionalAcknowledgements"`
	BatchTechnicalAcknowledgements         bool    `pulumi:"batchTechnicalAcknowledgements"`
	NeedFunctionalAcknowledgement          bool    `pulumi:"needFunctionalAcknowledgement"`
	NeedLoopForValidMessages               bool    `pulumi:"needLoopForValidMessages"`
	NeedTechnicalAcknowledgement           bool    `pulumi:"needTechnicalAcknowledgement"`
	RolloverAcknowledgementControlNumber   bool    `pulumi:"rolloverAcknowledgementControlNumber"`
	SendSynchronousAcknowledgement         bool    `pulumi:"sendSynchronousAcknowledgement"`
}





type EdifactAcknowledgementSettingsInput interface {
	pulumi.Input

	ToEdifactAcknowledgementSettingsOutput() EdifactAcknowledgementSettingsOutput
	ToEdifactAcknowledgementSettingsOutputWithContext(context.Context) EdifactAcknowledgementSettingsOutput
}

type EdifactAcknowledgementSettingsArgs struct {
	AcknowledgementControlNumberLowerBound pulumi.IntInput       `pulumi:"acknowledgementControlNumberLowerBound"`
	AcknowledgementControlNumberPrefix     pulumi.StringPtrInput `pulumi:"acknowledgementControlNumberPrefix"`
	AcknowledgementControlNumberSuffix     pulumi.StringPtrInput `pulumi:"acknowledgementControlNumberSuffix"`
	AcknowledgementControlNumberUpperBound pulumi.IntInput       `pulumi:"acknowledgementControlNumberUpperBound"`
	BatchFunctionalAcknowledgements        pulumi.BoolInput      `pulumi:"batchFunctionalAcknowledgements"`
	BatchTechnicalAcknowledgements         pulumi.BoolInput      `pulumi:"batchTechnicalAcknowledgements"`
	NeedFunctionalAcknowledgement          pulumi.BoolInput      `pulumi:"needFunctionalAcknowledgement"`
	NeedLoopForValidMessages               pulumi.BoolInput      `pulumi:"needLoopForValidMessages"`
	NeedTechnicalAcknowledgement           pulumi.BoolInput      `pulumi:"needTechnicalAcknowledgement"`
	RolloverAcknowledgementControlNumber   pulumi.BoolInput      `pulumi:"rolloverAcknowledgementControlNumber"`
	SendSynchronousAcknowledgement         pulumi.BoolInput      `pulumi:"sendSynchronousAcknowledgement"`
}

func (EdifactAcknowledgementSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactAcknowledgementSettings)(nil)).Elem()
}

func (i EdifactAcknowledgementSettingsArgs) ToEdifactAcknowledgementSettingsOutput() EdifactAcknowledgementSettingsOutput {
	return i.ToEdifactAcknowledgementSettingsOutputWithContext(context.Background())
}

func (i EdifactAcknowledgementSettingsArgs) ToEdifactAcknowledgementSettingsOutputWithContext(ctx context.Context) EdifactAcknowledgementSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactAcknowledgementSettingsOutput)
}

func (i EdifactAcknowledgementSettingsArgs) ToEdifactAcknowledgementSettingsPtrOutput() EdifactAcknowledgementSettingsPtrOutput {
	return i.ToEdifactAcknowledgementSettingsPtrOutputWithContext(context.Background())
}

func (i EdifactAcknowledgementSettingsArgs) ToEdifactAcknowledgementSettingsPtrOutputWithContext(ctx context.Context) EdifactAcknowledgementSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactAcknowledgementSettingsOutput).ToEdifactAcknowledgementSettingsPtrOutputWithContext(ctx)
}









type EdifactAcknowledgementSettingsPtrInput interface {
	pulumi.Input

	ToEdifactAcknowledgementSettingsPtrOutput() EdifactAcknowledgementSettingsPtrOutput
	ToEdifactAcknowledgementSettingsPtrOutputWithContext(context.Context) EdifactAcknowledgementSettingsPtrOutput
}

type edifactAcknowledgementSettingsPtrType EdifactAcknowledgementSettingsArgs

func EdifactAcknowledgementSettingsPtr(v *EdifactAcknowledgementSettingsArgs) EdifactAcknowledgementSettingsPtrInput {
	return (*edifactAcknowledgementSettingsPtrType)(v)
}

func (*edifactAcknowledgementSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactAcknowledgementSettings)(nil)).Elem()
}

func (i *edifactAcknowledgementSettingsPtrType) ToEdifactAcknowledgementSettingsPtrOutput() EdifactAcknowledgementSettingsPtrOutput {
	return i.ToEdifactAcknowledgementSettingsPtrOutputWithContext(context.Background())
}

func (i *edifactAcknowledgementSettingsPtrType) ToEdifactAcknowledgementSettingsPtrOutputWithContext(ctx context.Context) EdifactAcknowledgementSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactAcknowledgementSettingsPtrOutput)
}

type EdifactAcknowledgementSettingsOutput struct{ *pulumi.OutputState }

func (EdifactAcknowledgementSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactAcknowledgementSettings)(nil)).Elem()
}

func (o EdifactAcknowledgementSettingsOutput) ToEdifactAcknowledgementSettingsOutput() EdifactAcknowledgementSettingsOutput {
	return o
}

func (o EdifactAcknowledgementSettingsOutput) ToEdifactAcknowledgementSettingsOutputWithContext(ctx context.Context) EdifactAcknowledgementSettingsOutput {
	return o
}

func (o EdifactAcknowledgementSettingsOutput) ToEdifactAcknowledgementSettingsPtrOutput() EdifactAcknowledgementSettingsPtrOutput {
	return o.ToEdifactAcknowledgementSettingsPtrOutputWithContext(context.Background())
}

func (o EdifactAcknowledgementSettingsOutput) ToEdifactAcknowledgementSettingsPtrOutputWithContext(ctx context.Context) EdifactAcknowledgementSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactAcknowledgementSettings) *EdifactAcknowledgementSettings {
		return &v
	}).(EdifactAcknowledgementSettingsPtrOutput)
}

func (o EdifactAcknowledgementSettingsOutput) AcknowledgementControlNumberLowerBound() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) int { return v.AcknowledgementControlNumberLowerBound }).(pulumi.IntOutput)
}

func (o EdifactAcknowledgementSettingsOutput) AcknowledgementControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) *string { return v.AcknowledgementControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o EdifactAcknowledgementSettingsOutput) AcknowledgementControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) *string { return v.AcknowledgementControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o EdifactAcknowledgementSettingsOutput) AcknowledgementControlNumberUpperBound() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) int { return v.AcknowledgementControlNumberUpperBound }).(pulumi.IntOutput)
}

func (o EdifactAcknowledgementSettingsOutput) BatchFunctionalAcknowledgements() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) bool { return v.BatchFunctionalAcknowledgements }).(pulumi.BoolOutput)
}

func (o EdifactAcknowledgementSettingsOutput) BatchTechnicalAcknowledgements() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) bool { return v.BatchTechnicalAcknowledgements }).(pulumi.BoolOutput)
}

func (o EdifactAcknowledgementSettingsOutput) NeedFunctionalAcknowledgement() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) bool { return v.NeedFunctionalAcknowledgement }).(pulumi.BoolOutput)
}

func (o EdifactAcknowledgementSettingsOutput) NeedLoopForValidMessages() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) bool { return v.NeedLoopForValidMessages }).(pulumi.BoolOutput)
}

func (o EdifactAcknowledgementSettingsOutput) NeedTechnicalAcknowledgement() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) bool { return v.NeedTechnicalAcknowledgement }).(pulumi.BoolOutput)
}

func (o EdifactAcknowledgementSettingsOutput) RolloverAcknowledgementControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) bool { return v.RolloverAcknowledgementControlNumber }).(pulumi.BoolOutput)
}

func (o EdifactAcknowledgementSettingsOutput) SendSynchronousAcknowledgement() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettings) bool { return v.SendSynchronousAcknowledgement }).(pulumi.BoolOutput)
}

type EdifactAcknowledgementSettingsPtrOutput struct{ *pulumi.OutputState }

func (EdifactAcknowledgementSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactAcknowledgementSettings)(nil)).Elem()
}

func (o EdifactAcknowledgementSettingsPtrOutput) ToEdifactAcknowledgementSettingsPtrOutput() EdifactAcknowledgementSettingsPtrOutput {
	return o
}

func (o EdifactAcknowledgementSettingsPtrOutput) ToEdifactAcknowledgementSettingsPtrOutputWithContext(ctx context.Context) EdifactAcknowledgementSettingsPtrOutput {
	return o
}

func (o EdifactAcknowledgementSettingsPtrOutput) Elem() EdifactAcknowledgementSettingsOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) EdifactAcknowledgementSettings {
		if v != nil {
			return *v
		}
		var ret EdifactAcknowledgementSettings
		return ret
	}).(EdifactAcknowledgementSettingsOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) AcknowledgementControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *int {
		if v == nil {
			return nil
		}
		return &v.AcknowledgementControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) AcknowledgementControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *string {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) AcknowledgementControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *string {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) AcknowledgementControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *int {
		if v == nil {
			return nil
		}
		return &v.AcknowledgementControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) BatchFunctionalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.BatchFunctionalAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) BatchTechnicalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.BatchTechnicalAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) NeedFunctionalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.NeedFunctionalAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) NeedLoopForValidMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.NeedLoopForValidMessages
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) NeedTechnicalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.NeedTechnicalAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) RolloverAcknowledgementControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.RolloverAcknowledgementControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsPtrOutput) SendSynchronousAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.SendSynchronousAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

type EdifactAcknowledgementSettingsResponse struct {
	AcknowledgementControlNumberLowerBound int     `pulumi:"acknowledgementControlNumberLowerBound"`
	AcknowledgementControlNumberPrefix     *string `pulumi:"acknowledgementControlNumberPrefix"`
	AcknowledgementControlNumberSuffix     *string `pulumi:"acknowledgementControlNumberSuffix"`
	AcknowledgementControlNumberUpperBound int     `pulumi:"acknowledgementControlNumberUpperBound"`
	BatchFunctionalAcknowledgements        bool    `pulumi:"batchFunctionalAcknowledgements"`
	BatchTechnicalAcknowledgements         bool    `pulumi:"batchTechnicalAcknowledgements"`
	NeedFunctionalAcknowledgement          bool    `pulumi:"needFunctionalAcknowledgement"`
	NeedLoopForValidMessages               bool    `pulumi:"needLoopForValidMessages"`
	NeedTechnicalAcknowledgement           bool    `pulumi:"needTechnicalAcknowledgement"`
	RolloverAcknowledgementControlNumber   bool    `pulumi:"rolloverAcknowledgementControlNumber"`
	SendSynchronousAcknowledgement         bool    `pulumi:"sendSynchronousAcknowledgement"`
}

type EdifactAcknowledgementSettingsResponseOutput struct{ *pulumi.OutputState }

func (EdifactAcknowledgementSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactAcknowledgementSettingsResponse)(nil)).Elem()
}

func (o EdifactAcknowledgementSettingsResponseOutput) ToEdifactAcknowledgementSettingsResponseOutput() EdifactAcknowledgementSettingsResponseOutput {
	return o
}

func (o EdifactAcknowledgementSettingsResponseOutput) ToEdifactAcknowledgementSettingsResponseOutputWithContext(ctx context.Context) EdifactAcknowledgementSettingsResponseOutput {
	return o
}

func (o EdifactAcknowledgementSettingsResponseOutput) AcknowledgementControlNumberLowerBound() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) int { return v.AcknowledgementControlNumberLowerBound }).(pulumi.IntOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) AcknowledgementControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) *string { return v.AcknowledgementControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) AcknowledgementControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) *string { return v.AcknowledgementControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) AcknowledgementControlNumberUpperBound() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) int { return v.AcknowledgementControlNumberUpperBound }).(pulumi.IntOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) BatchFunctionalAcknowledgements() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) bool { return v.BatchFunctionalAcknowledgements }).(pulumi.BoolOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) BatchTechnicalAcknowledgements() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) bool { return v.BatchTechnicalAcknowledgements }).(pulumi.BoolOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) NeedFunctionalAcknowledgement() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) bool { return v.NeedFunctionalAcknowledgement }).(pulumi.BoolOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) NeedLoopForValidMessages() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) bool { return v.NeedLoopForValidMessages }).(pulumi.BoolOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) NeedTechnicalAcknowledgement() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) bool { return v.NeedTechnicalAcknowledgement }).(pulumi.BoolOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) RolloverAcknowledgementControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) bool { return v.RolloverAcknowledgementControlNumber }).(pulumi.BoolOutput)
}

func (o EdifactAcknowledgementSettingsResponseOutput) SendSynchronousAcknowledgement() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactAcknowledgementSettingsResponse) bool { return v.SendSynchronousAcknowledgement }).(pulumi.BoolOutput)
}

type EdifactAcknowledgementSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (EdifactAcknowledgementSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactAcknowledgementSettingsResponse)(nil)).Elem()
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) ToEdifactAcknowledgementSettingsResponsePtrOutput() EdifactAcknowledgementSettingsResponsePtrOutput {
	return o
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) ToEdifactAcknowledgementSettingsResponsePtrOutputWithContext(ctx context.Context) EdifactAcknowledgementSettingsResponsePtrOutput {
	return o
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) Elem() EdifactAcknowledgementSettingsResponseOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) EdifactAcknowledgementSettingsResponse {
		if v != nil {
			return *v
		}
		var ret EdifactAcknowledgementSettingsResponse
		return ret
	}).(EdifactAcknowledgementSettingsResponseOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) AcknowledgementControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.AcknowledgementControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) AcknowledgementControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) AcknowledgementControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) AcknowledgementControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.AcknowledgementControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) BatchFunctionalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.BatchFunctionalAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) BatchTechnicalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.BatchTechnicalAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) NeedFunctionalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.NeedFunctionalAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) NeedLoopForValidMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.NeedLoopForValidMessages
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) NeedTechnicalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.NeedTechnicalAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) RolloverAcknowledgementControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.RolloverAcknowledgementControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactAcknowledgementSettingsResponsePtrOutput) SendSynchronousAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactAcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.SendSynchronousAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

type EdifactAgreementContent struct {
	ReceiveAgreement EdifactOneWayAgreement `pulumi:"receiveAgreement"`
	SendAgreement    EdifactOneWayAgreement `pulumi:"sendAgreement"`
}





type EdifactAgreementContentInput interface {
	pulumi.Input

	ToEdifactAgreementContentOutput() EdifactAgreementContentOutput
	ToEdifactAgreementContentOutputWithContext(context.Context) EdifactAgreementContentOutput
}

type EdifactAgreementContentArgs struct {
	ReceiveAgreement EdifactOneWayAgreementInput `pulumi:"receiveAgreement"`
	SendAgreement    EdifactOneWayAgreementInput `pulumi:"sendAgreement"`
}

func (EdifactAgreementContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactAgreementContent)(nil)).Elem()
}

func (i EdifactAgreementContentArgs) ToEdifactAgreementContentOutput() EdifactAgreementContentOutput {
	return i.ToEdifactAgreementContentOutputWithContext(context.Background())
}

func (i EdifactAgreementContentArgs) ToEdifactAgreementContentOutputWithContext(ctx context.Context) EdifactAgreementContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactAgreementContentOutput)
}

func (i EdifactAgreementContentArgs) ToEdifactAgreementContentPtrOutput() EdifactAgreementContentPtrOutput {
	return i.ToEdifactAgreementContentPtrOutputWithContext(context.Background())
}

func (i EdifactAgreementContentArgs) ToEdifactAgreementContentPtrOutputWithContext(ctx context.Context) EdifactAgreementContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactAgreementContentOutput).ToEdifactAgreementContentPtrOutputWithContext(ctx)
}









type EdifactAgreementContentPtrInput interface {
	pulumi.Input

	ToEdifactAgreementContentPtrOutput() EdifactAgreementContentPtrOutput
	ToEdifactAgreementContentPtrOutputWithContext(context.Context) EdifactAgreementContentPtrOutput
}

type edifactAgreementContentPtrType EdifactAgreementContentArgs

func EdifactAgreementContentPtr(v *EdifactAgreementContentArgs) EdifactAgreementContentPtrInput {
	return (*edifactAgreementContentPtrType)(v)
}

func (*edifactAgreementContentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactAgreementContent)(nil)).Elem()
}

func (i *edifactAgreementContentPtrType) ToEdifactAgreementContentPtrOutput() EdifactAgreementContentPtrOutput {
	return i.ToEdifactAgreementContentPtrOutputWithContext(context.Background())
}

func (i *edifactAgreementContentPtrType) ToEdifactAgreementContentPtrOutputWithContext(ctx context.Context) EdifactAgreementContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactAgreementContentPtrOutput)
}

type EdifactAgreementContentOutput struct{ *pulumi.OutputState }

func (EdifactAgreementContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactAgreementContent)(nil)).Elem()
}

func (o EdifactAgreementContentOutput) ToEdifactAgreementContentOutput() EdifactAgreementContentOutput {
	return o
}

func (o EdifactAgreementContentOutput) ToEdifactAgreementContentOutputWithContext(ctx context.Context) EdifactAgreementContentOutput {
	return o
}

func (o EdifactAgreementContentOutput) ToEdifactAgreementContentPtrOutput() EdifactAgreementContentPtrOutput {
	return o.ToEdifactAgreementContentPtrOutputWithContext(context.Background())
}

func (o EdifactAgreementContentOutput) ToEdifactAgreementContentPtrOutputWithContext(ctx context.Context) EdifactAgreementContentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactAgreementContent) *EdifactAgreementContent {
		return &v
	}).(EdifactAgreementContentPtrOutput)
}

func (o EdifactAgreementContentOutput) ReceiveAgreement() EdifactOneWayAgreementOutput {
	return o.ApplyT(func(v EdifactAgreementContent) EdifactOneWayAgreement { return v.ReceiveAgreement }).(EdifactOneWayAgreementOutput)
}

func (o EdifactAgreementContentOutput) SendAgreement() EdifactOneWayAgreementOutput {
	return o.ApplyT(func(v EdifactAgreementContent) EdifactOneWayAgreement { return v.SendAgreement }).(EdifactOneWayAgreementOutput)
}

type EdifactAgreementContentPtrOutput struct{ *pulumi.OutputState }

func (EdifactAgreementContentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactAgreementContent)(nil)).Elem()
}

func (o EdifactAgreementContentPtrOutput) ToEdifactAgreementContentPtrOutput() EdifactAgreementContentPtrOutput {
	return o
}

func (o EdifactAgreementContentPtrOutput) ToEdifactAgreementContentPtrOutputWithContext(ctx context.Context) EdifactAgreementContentPtrOutput {
	return o
}

func (o EdifactAgreementContentPtrOutput) Elem() EdifactAgreementContentOutput {
	return o.ApplyT(func(v *EdifactAgreementContent) EdifactAgreementContent {
		if v != nil {
			return *v
		}
		var ret EdifactAgreementContent
		return ret
	}).(EdifactAgreementContentOutput)
}

func (o EdifactAgreementContentPtrOutput) ReceiveAgreement() EdifactOneWayAgreementPtrOutput {
	return o.ApplyT(func(v *EdifactAgreementContent) *EdifactOneWayAgreement {
		if v == nil {
			return nil
		}
		return &v.ReceiveAgreement
	}).(EdifactOneWayAgreementPtrOutput)
}

func (o EdifactAgreementContentPtrOutput) SendAgreement() EdifactOneWayAgreementPtrOutput {
	return o.ApplyT(func(v *EdifactAgreementContent) *EdifactOneWayAgreement {
		if v == nil {
			return nil
		}
		return &v.SendAgreement
	}).(EdifactOneWayAgreementPtrOutput)
}

type EdifactAgreementContentResponse struct {
	ReceiveAgreement EdifactOneWayAgreementResponse `pulumi:"receiveAgreement"`
	SendAgreement    EdifactOneWayAgreementResponse `pulumi:"sendAgreement"`
}

type EdifactAgreementContentResponseOutput struct{ *pulumi.OutputState }

func (EdifactAgreementContentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactAgreementContentResponse)(nil)).Elem()
}

func (o EdifactAgreementContentResponseOutput) ToEdifactAgreementContentResponseOutput() EdifactAgreementContentResponseOutput {
	return o
}

func (o EdifactAgreementContentResponseOutput) ToEdifactAgreementContentResponseOutputWithContext(ctx context.Context) EdifactAgreementContentResponseOutput {
	return o
}

func (o EdifactAgreementContentResponseOutput) ReceiveAgreement() EdifactOneWayAgreementResponseOutput {
	return o.ApplyT(func(v EdifactAgreementContentResponse) EdifactOneWayAgreementResponse { return v.ReceiveAgreement }).(EdifactOneWayAgreementResponseOutput)
}

func (o EdifactAgreementContentResponseOutput) SendAgreement() EdifactOneWayAgreementResponseOutput {
	return o.ApplyT(func(v EdifactAgreementContentResponse) EdifactOneWayAgreementResponse { return v.SendAgreement }).(EdifactOneWayAgreementResponseOutput)
}

type EdifactAgreementContentResponsePtrOutput struct{ *pulumi.OutputState }

func (EdifactAgreementContentResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactAgreementContentResponse)(nil)).Elem()
}

func (o EdifactAgreementContentResponsePtrOutput) ToEdifactAgreementContentResponsePtrOutput() EdifactAgreementContentResponsePtrOutput {
	return o
}

func (o EdifactAgreementContentResponsePtrOutput) ToEdifactAgreementContentResponsePtrOutputWithContext(ctx context.Context) EdifactAgreementContentResponsePtrOutput {
	return o
}

func (o EdifactAgreementContentResponsePtrOutput) Elem() EdifactAgreementContentResponseOutput {
	return o.ApplyT(func(v *EdifactAgreementContentResponse) EdifactAgreementContentResponse {
		if v != nil {
			return *v
		}
		var ret EdifactAgreementContentResponse
		return ret
	}).(EdifactAgreementContentResponseOutput)
}

func (o EdifactAgreementContentResponsePtrOutput) ReceiveAgreement() EdifactOneWayAgreementResponsePtrOutput {
	return o.ApplyT(func(v *EdifactAgreementContentResponse) *EdifactOneWayAgreementResponse {
		if v == nil {
			return nil
		}
		return &v.ReceiveAgreement
	}).(EdifactOneWayAgreementResponsePtrOutput)
}

func (o EdifactAgreementContentResponsePtrOutput) SendAgreement() EdifactOneWayAgreementResponsePtrOutput {
	return o.ApplyT(func(v *EdifactAgreementContentResponse) *EdifactOneWayAgreementResponse {
		if v == nil {
			return nil
		}
		return &v.SendAgreement
	}).(EdifactOneWayAgreementResponsePtrOutput)
}

type EdifactDelimiterOverride struct {
	ComponentSeparator             int                     `pulumi:"componentSeparator"`
	DataElementSeparator           int                     `pulumi:"dataElementSeparator"`
	DecimalPointIndicator          EdifactDecimalIndicator `pulumi:"decimalPointIndicator"`
	MessageAssociationAssignedCode *string                 `pulumi:"messageAssociationAssignedCode"`
	MessageId                      *string                 `pulumi:"messageId"`
	MessageRelease                 *string                 `pulumi:"messageRelease"`
	MessageVersion                 *string                 `pulumi:"messageVersion"`
	ReleaseIndicator               int                     `pulumi:"releaseIndicator"`
	RepetitionSeparator            int                     `pulumi:"repetitionSeparator"`
	SegmentTerminator              int                     `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix        SegmentTerminatorSuffix `pulumi:"segmentTerminatorSuffix"`
	TargetNamespace                *string                 `pulumi:"targetNamespace"`
}





type EdifactDelimiterOverrideInput interface {
	pulumi.Input

	ToEdifactDelimiterOverrideOutput() EdifactDelimiterOverrideOutput
	ToEdifactDelimiterOverrideOutputWithContext(context.Context) EdifactDelimiterOverrideOutput
}

type EdifactDelimiterOverrideArgs struct {
	ComponentSeparator             pulumi.IntInput              `pulumi:"componentSeparator"`
	DataElementSeparator           pulumi.IntInput              `pulumi:"dataElementSeparator"`
	DecimalPointIndicator          EdifactDecimalIndicatorInput `pulumi:"decimalPointIndicator"`
	MessageAssociationAssignedCode pulumi.StringPtrInput        `pulumi:"messageAssociationAssignedCode"`
	MessageId                      pulumi.StringPtrInput        `pulumi:"messageId"`
	MessageRelease                 pulumi.StringPtrInput        `pulumi:"messageRelease"`
	MessageVersion                 pulumi.StringPtrInput        `pulumi:"messageVersion"`
	ReleaseIndicator               pulumi.IntInput              `pulumi:"releaseIndicator"`
	RepetitionSeparator            pulumi.IntInput              `pulumi:"repetitionSeparator"`
	SegmentTerminator              pulumi.IntInput              `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix        SegmentTerminatorSuffixInput `pulumi:"segmentTerminatorSuffix"`
	TargetNamespace                pulumi.StringPtrInput        `pulumi:"targetNamespace"`
}

func (EdifactDelimiterOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactDelimiterOverride)(nil)).Elem()
}

func (i EdifactDelimiterOverrideArgs) ToEdifactDelimiterOverrideOutput() EdifactDelimiterOverrideOutput {
	return i.ToEdifactDelimiterOverrideOutputWithContext(context.Background())
}

func (i EdifactDelimiterOverrideArgs) ToEdifactDelimiterOverrideOutputWithContext(ctx context.Context) EdifactDelimiterOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactDelimiterOverrideOutput)
}





type EdifactDelimiterOverrideArrayInput interface {
	pulumi.Input

	ToEdifactDelimiterOverrideArrayOutput() EdifactDelimiterOverrideArrayOutput
	ToEdifactDelimiterOverrideArrayOutputWithContext(context.Context) EdifactDelimiterOverrideArrayOutput
}

type EdifactDelimiterOverrideArray []EdifactDelimiterOverrideInput

func (EdifactDelimiterOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactDelimiterOverride)(nil)).Elem()
}

func (i EdifactDelimiterOverrideArray) ToEdifactDelimiterOverrideArrayOutput() EdifactDelimiterOverrideArrayOutput {
	return i.ToEdifactDelimiterOverrideArrayOutputWithContext(context.Background())
}

func (i EdifactDelimiterOverrideArray) ToEdifactDelimiterOverrideArrayOutputWithContext(ctx context.Context) EdifactDelimiterOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactDelimiterOverrideArrayOutput)
}

type EdifactDelimiterOverrideOutput struct{ *pulumi.OutputState }

func (EdifactDelimiterOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactDelimiterOverride)(nil)).Elem()
}

func (o EdifactDelimiterOverrideOutput) ToEdifactDelimiterOverrideOutput() EdifactDelimiterOverrideOutput {
	return o
}

func (o EdifactDelimiterOverrideOutput) ToEdifactDelimiterOverrideOutputWithContext(ctx context.Context) EdifactDelimiterOverrideOutput {
	return o
}

func (o EdifactDelimiterOverrideOutput) ComponentSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) int { return v.ComponentSeparator }).(pulumi.IntOutput)
}

func (o EdifactDelimiterOverrideOutput) DataElementSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) int { return v.DataElementSeparator }).(pulumi.IntOutput)
}

func (o EdifactDelimiterOverrideOutput) DecimalPointIndicator() EdifactDecimalIndicatorOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) EdifactDecimalIndicator { return v.DecimalPointIndicator }).(EdifactDecimalIndicatorOutput)
}

func (o EdifactDelimiterOverrideOutput) MessageAssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) *string { return v.MessageAssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactDelimiterOverrideOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o EdifactDelimiterOverrideOutput) MessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) *string { return v.MessageRelease }).(pulumi.StringPtrOutput)
}

func (o EdifactDelimiterOverrideOutput) MessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) *string { return v.MessageVersion }).(pulumi.StringPtrOutput)
}

func (o EdifactDelimiterOverrideOutput) ReleaseIndicator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) int { return v.ReleaseIndicator }).(pulumi.IntOutput)
}

func (o EdifactDelimiterOverrideOutput) RepetitionSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) int { return v.RepetitionSeparator }).(pulumi.IntOutput)
}

func (o EdifactDelimiterOverrideOutput) SegmentTerminator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) int { return v.SegmentTerminator }).(pulumi.IntOutput)
}

func (o EdifactDelimiterOverrideOutput) SegmentTerminatorSuffix() SegmentTerminatorSuffixOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) SegmentTerminatorSuffix { return v.SegmentTerminatorSuffix }).(SegmentTerminatorSuffixOutput)
}

func (o EdifactDelimiterOverrideOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverride) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

type EdifactDelimiterOverrideArrayOutput struct{ *pulumi.OutputState }

func (EdifactDelimiterOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactDelimiterOverride)(nil)).Elem()
}

func (o EdifactDelimiterOverrideArrayOutput) ToEdifactDelimiterOverrideArrayOutput() EdifactDelimiterOverrideArrayOutput {
	return o
}

func (o EdifactDelimiterOverrideArrayOutput) ToEdifactDelimiterOverrideArrayOutputWithContext(ctx context.Context) EdifactDelimiterOverrideArrayOutput {
	return o
}

func (o EdifactDelimiterOverrideArrayOutput) Index(i pulumi.IntInput) EdifactDelimiterOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactDelimiterOverride {
		return vs[0].([]EdifactDelimiterOverride)[vs[1].(int)]
	}).(EdifactDelimiterOverrideOutput)
}

type EdifactDelimiterOverrideResponse struct {
	ComponentSeparator             int     `pulumi:"componentSeparator"`
	DataElementSeparator           int     `pulumi:"dataElementSeparator"`
	DecimalPointIndicator          string  `pulumi:"decimalPointIndicator"`
	MessageAssociationAssignedCode *string `pulumi:"messageAssociationAssignedCode"`
	MessageId                      *string `pulumi:"messageId"`
	MessageRelease                 *string `pulumi:"messageRelease"`
	MessageVersion                 *string `pulumi:"messageVersion"`
	ReleaseIndicator               int     `pulumi:"releaseIndicator"`
	RepetitionSeparator            int     `pulumi:"repetitionSeparator"`
	SegmentTerminator              int     `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix        string  `pulumi:"segmentTerminatorSuffix"`
	TargetNamespace                *string `pulumi:"targetNamespace"`
}

type EdifactDelimiterOverrideResponseOutput struct{ *pulumi.OutputState }

func (EdifactDelimiterOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactDelimiterOverrideResponse)(nil)).Elem()
}

func (o EdifactDelimiterOverrideResponseOutput) ToEdifactDelimiterOverrideResponseOutput() EdifactDelimiterOverrideResponseOutput {
	return o
}

func (o EdifactDelimiterOverrideResponseOutput) ToEdifactDelimiterOverrideResponseOutputWithContext(ctx context.Context) EdifactDelimiterOverrideResponseOutput {
	return o
}

func (o EdifactDelimiterOverrideResponseOutput) ComponentSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) int { return v.ComponentSeparator }).(pulumi.IntOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) DataElementSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) int { return v.DataElementSeparator }).(pulumi.IntOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) DecimalPointIndicator() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) string { return v.DecimalPointIndicator }).(pulumi.StringOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) MessageAssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) *string { return v.MessageAssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) MessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) *string { return v.MessageRelease }).(pulumi.StringPtrOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) MessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) *string { return v.MessageVersion }).(pulumi.StringPtrOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) ReleaseIndicator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) int { return v.ReleaseIndicator }).(pulumi.IntOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) RepetitionSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) int { return v.RepetitionSeparator }).(pulumi.IntOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) SegmentTerminator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) int { return v.SegmentTerminator }).(pulumi.IntOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) SegmentTerminatorSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) string { return v.SegmentTerminatorSuffix }).(pulumi.StringOutput)
}

func (o EdifactDelimiterOverrideResponseOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactDelimiterOverrideResponse) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

type EdifactDelimiterOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (EdifactDelimiterOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactDelimiterOverrideResponse)(nil)).Elem()
}

func (o EdifactDelimiterOverrideResponseArrayOutput) ToEdifactDelimiterOverrideResponseArrayOutput() EdifactDelimiterOverrideResponseArrayOutput {
	return o
}

func (o EdifactDelimiterOverrideResponseArrayOutput) ToEdifactDelimiterOverrideResponseArrayOutputWithContext(ctx context.Context) EdifactDelimiterOverrideResponseArrayOutput {
	return o
}

func (o EdifactDelimiterOverrideResponseArrayOutput) Index(i pulumi.IntInput) EdifactDelimiterOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactDelimiterOverrideResponse {
		return vs[0].([]EdifactDelimiterOverrideResponse)[vs[1].(int)]
	}).(EdifactDelimiterOverrideResponseOutput)
}

type EdifactEnvelopeOverride struct {
	ApplicationPassword            *string `pulumi:"applicationPassword"`
	AssociationAssignedCode        *string `pulumi:"associationAssignedCode"`
	ControllingAgencyCode          *string `pulumi:"controllingAgencyCode"`
	FunctionalGroupId              *string `pulumi:"functionalGroupId"`
	GroupHeaderMessageRelease      *string `pulumi:"groupHeaderMessageRelease"`
	GroupHeaderMessageVersion      *string `pulumi:"groupHeaderMessageVersion"`
	MessageAssociationAssignedCode *string `pulumi:"messageAssociationAssignedCode"`
	MessageId                      *string `pulumi:"messageId"`
	MessageRelease                 *string `pulumi:"messageRelease"`
	MessageVersion                 *string `pulumi:"messageVersion"`
	ReceiverApplicationId          *string `pulumi:"receiverApplicationId"`
	ReceiverApplicationQualifier   *string `pulumi:"receiverApplicationQualifier"`
	SenderApplicationId            *string `pulumi:"senderApplicationId"`
	SenderApplicationQualifier     *string `pulumi:"senderApplicationQualifier"`
	TargetNamespace                *string `pulumi:"targetNamespace"`
}





type EdifactEnvelopeOverrideInput interface {
	pulumi.Input

	ToEdifactEnvelopeOverrideOutput() EdifactEnvelopeOverrideOutput
	ToEdifactEnvelopeOverrideOutputWithContext(context.Context) EdifactEnvelopeOverrideOutput
}

type EdifactEnvelopeOverrideArgs struct {
	ApplicationPassword            pulumi.StringPtrInput `pulumi:"applicationPassword"`
	AssociationAssignedCode        pulumi.StringPtrInput `pulumi:"associationAssignedCode"`
	ControllingAgencyCode          pulumi.StringPtrInput `pulumi:"controllingAgencyCode"`
	FunctionalGroupId              pulumi.StringPtrInput `pulumi:"functionalGroupId"`
	GroupHeaderMessageRelease      pulumi.StringPtrInput `pulumi:"groupHeaderMessageRelease"`
	GroupHeaderMessageVersion      pulumi.StringPtrInput `pulumi:"groupHeaderMessageVersion"`
	MessageAssociationAssignedCode pulumi.StringPtrInput `pulumi:"messageAssociationAssignedCode"`
	MessageId                      pulumi.StringPtrInput `pulumi:"messageId"`
	MessageRelease                 pulumi.StringPtrInput `pulumi:"messageRelease"`
	MessageVersion                 pulumi.StringPtrInput `pulumi:"messageVersion"`
	ReceiverApplicationId          pulumi.StringPtrInput `pulumi:"receiverApplicationId"`
	ReceiverApplicationQualifier   pulumi.StringPtrInput `pulumi:"receiverApplicationQualifier"`
	SenderApplicationId            pulumi.StringPtrInput `pulumi:"senderApplicationId"`
	SenderApplicationQualifier     pulumi.StringPtrInput `pulumi:"senderApplicationQualifier"`
	TargetNamespace                pulumi.StringPtrInput `pulumi:"targetNamespace"`
}

func (EdifactEnvelopeOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactEnvelopeOverride)(nil)).Elem()
}

func (i EdifactEnvelopeOverrideArgs) ToEdifactEnvelopeOverrideOutput() EdifactEnvelopeOverrideOutput {
	return i.ToEdifactEnvelopeOverrideOutputWithContext(context.Background())
}

func (i EdifactEnvelopeOverrideArgs) ToEdifactEnvelopeOverrideOutputWithContext(ctx context.Context) EdifactEnvelopeOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactEnvelopeOverrideOutput)
}





type EdifactEnvelopeOverrideArrayInput interface {
	pulumi.Input

	ToEdifactEnvelopeOverrideArrayOutput() EdifactEnvelopeOverrideArrayOutput
	ToEdifactEnvelopeOverrideArrayOutputWithContext(context.Context) EdifactEnvelopeOverrideArrayOutput
}

type EdifactEnvelopeOverrideArray []EdifactEnvelopeOverrideInput

func (EdifactEnvelopeOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactEnvelopeOverride)(nil)).Elem()
}

func (i EdifactEnvelopeOverrideArray) ToEdifactEnvelopeOverrideArrayOutput() EdifactEnvelopeOverrideArrayOutput {
	return i.ToEdifactEnvelopeOverrideArrayOutputWithContext(context.Background())
}

func (i EdifactEnvelopeOverrideArray) ToEdifactEnvelopeOverrideArrayOutputWithContext(ctx context.Context) EdifactEnvelopeOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactEnvelopeOverrideArrayOutput)
}

type EdifactEnvelopeOverrideOutput struct{ *pulumi.OutputState }

func (EdifactEnvelopeOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactEnvelopeOverride)(nil)).Elem()
}

func (o EdifactEnvelopeOverrideOutput) ToEdifactEnvelopeOverrideOutput() EdifactEnvelopeOverrideOutput {
	return o
}

func (o EdifactEnvelopeOverrideOutput) ToEdifactEnvelopeOverrideOutputWithContext(ctx context.Context) EdifactEnvelopeOverrideOutput {
	return o
}

func (o EdifactEnvelopeOverrideOutput) ApplicationPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.ApplicationPassword }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) AssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.AssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) ControllingAgencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.ControllingAgencyCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.FunctionalGroupId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) GroupHeaderMessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.GroupHeaderMessageRelease }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) GroupHeaderMessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.GroupHeaderMessageVersion }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) MessageAssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.MessageAssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) MessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.MessageRelease }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) MessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.MessageVersion }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) ReceiverApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.ReceiverApplicationId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) ReceiverApplicationQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.ReceiverApplicationQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.SenderApplicationId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) SenderApplicationQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.SenderApplicationQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverride) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

type EdifactEnvelopeOverrideArrayOutput struct{ *pulumi.OutputState }

func (EdifactEnvelopeOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactEnvelopeOverride)(nil)).Elem()
}

func (o EdifactEnvelopeOverrideArrayOutput) ToEdifactEnvelopeOverrideArrayOutput() EdifactEnvelopeOverrideArrayOutput {
	return o
}

func (o EdifactEnvelopeOverrideArrayOutput) ToEdifactEnvelopeOverrideArrayOutputWithContext(ctx context.Context) EdifactEnvelopeOverrideArrayOutput {
	return o
}

func (o EdifactEnvelopeOverrideArrayOutput) Index(i pulumi.IntInput) EdifactEnvelopeOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactEnvelopeOverride {
		return vs[0].([]EdifactEnvelopeOverride)[vs[1].(int)]
	}).(EdifactEnvelopeOverrideOutput)
}

type EdifactEnvelopeOverrideResponse struct {
	ApplicationPassword            *string `pulumi:"applicationPassword"`
	AssociationAssignedCode        *string `pulumi:"associationAssignedCode"`
	ControllingAgencyCode          *string `pulumi:"controllingAgencyCode"`
	FunctionalGroupId              *string `pulumi:"functionalGroupId"`
	GroupHeaderMessageRelease      *string `pulumi:"groupHeaderMessageRelease"`
	GroupHeaderMessageVersion      *string `pulumi:"groupHeaderMessageVersion"`
	MessageAssociationAssignedCode *string `pulumi:"messageAssociationAssignedCode"`
	MessageId                      *string `pulumi:"messageId"`
	MessageRelease                 *string `pulumi:"messageRelease"`
	MessageVersion                 *string `pulumi:"messageVersion"`
	ReceiverApplicationId          *string `pulumi:"receiverApplicationId"`
	ReceiverApplicationQualifier   *string `pulumi:"receiverApplicationQualifier"`
	SenderApplicationId            *string `pulumi:"senderApplicationId"`
	SenderApplicationQualifier     *string `pulumi:"senderApplicationQualifier"`
	TargetNamespace                *string `pulumi:"targetNamespace"`
}

type EdifactEnvelopeOverrideResponseOutput struct{ *pulumi.OutputState }

func (EdifactEnvelopeOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactEnvelopeOverrideResponse)(nil)).Elem()
}

func (o EdifactEnvelopeOverrideResponseOutput) ToEdifactEnvelopeOverrideResponseOutput() EdifactEnvelopeOverrideResponseOutput {
	return o
}

func (o EdifactEnvelopeOverrideResponseOutput) ToEdifactEnvelopeOverrideResponseOutputWithContext(ctx context.Context) EdifactEnvelopeOverrideResponseOutput {
	return o
}

func (o EdifactEnvelopeOverrideResponseOutput) ApplicationPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.ApplicationPassword }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) AssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.AssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) ControllingAgencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.ControllingAgencyCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.FunctionalGroupId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) GroupHeaderMessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.GroupHeaderMessageRelease }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) GroupHeaderMessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.GroupHeaderMessageVersion }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) MessageAssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.MessageAssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) MessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.MessageRelease }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) MessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.MessageVersion }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) ReceiverApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.ReceiverApplicationId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) ReceiverApplicationQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.ReceiverApplicationQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.SenderApplicationId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) SenderApplicationQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.SenderApplicationQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeOverrideResponseOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeOverrideResponse) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

type EdifactEnvelopeOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (EdifactEnvelopeOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactEnvelopeOverrideResponse)(nil)).Elem()
}

func (o EdifactEnvelopeOverrideResponseArrayOutput) ToEdifactEnvelopeOverrideResponseArrayOutput() EdifactEnvelopeOverrideResponseArrayOutput {
	return o
}

func (o EdifactEnvelopeOverrideResponseArrayOutput) ToEdifactEnvelopeOverrideResponseArrayOutputWithContext(ctx context.Context) EdifactEnvelopeOverrideResponseArrayOutput {
	return o
}

func (o EdifactEnvelopeOverrideResponseArrayOutput) Index(i pulumi.IntInput) EdifactEnvelopeOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactEnvelopeOverrideResponse {
		return vs[0].([]EdifactEnvelopeOverrideResponse)[vs[1].(int)]
	}).(EdifactEnvelopeOverrideResponseOutput)
}

type EdifactEnvelopeSettings struct {
	ApplicationReferenceId                       *string `pulumi:"applicationReferenceId"`
	ApplyDelimiterStringAdvice                   bool    `pulumi:"applyDelimiterStringAdvice"`
	CommunicationAgreementId                     *string `pulumi:"communicationAgreementId"`
	CreateGroupingSegments                       bool    `pulumi:"createGroupingSegments"`
	EnableDefaultGroupHeaders                    bool    `pulumi:"enableDefaultGroupHeaders"`
	FunctionalGroupId                            *string `pulumi:"functionalGroupId"`
	GroupApplicationPassword                     *string `pulumi:"groupApplicationPassword"`
	GroupApplicationReceiverId                   *string `pulumi:"groupApplicationReceiverId"`
	GroupApplicationReceiverQualifier            *string `pulumi:"groupApplicationReceiverQualifier"`
	GroupApplicationSenderId                     *string `pulumi:"groupApplicationSenderId"`
	GroupApplicationSenderQualifier              *string `pulumi:"groupApplicationSenderQualifier"`
	GroupAssociationAssignedCode                 *string `pulumi:"groupAssociationAssignedCode"`
	GroupControlNumberLowerBound                 float64 `pulumi:"groupControlNumberLowerBound"`
	GroupControlNumberPrefix                     *string `pulumi:"groupControlNumberPrefix"`
	GroupControlNumberSuffix                     *string `pulumi:"groupControlNumberSuffix"`
	GroupControlNumberUpperBound                 float64 `pulumi:"groupControlNumberUpperBound"`
	GroupControllingAgencyCode                   *string `pulumi:"groupControllingAgencyCode"`
	GroupMessageRelease                          *string `pulumi:"groupMessageRelease"`
	GroupMessageVersion                          *string `pulumi:"groupMessageVersion"`
	InterchangeControlNumberLowerBound           float64 `pulumi:"interchangeControlNumberLowerBound"`
	InterchangeControlNumberPrefix               *string `pulumi:"interchangeControlNumberPrefix"`
	InterchangeControlNumberSuffix               *string `pulumi:"interchangeControlNumberSuffix"`
	InterchangeControlNumberUpperBound           float64 `pulumi:"interchangeControlNumberUpperBound"`
	IsTestInterchange                            bool    `pulumi:"isTestInterchange"`
	OverwriteExistingTransactionSetControlNumber bool    `pulumi:"overwriteExistingTransactionSetControlNumber"`
	ProcessingPriorityCode                       *string `pulumi:"processingPriorityCode"`
	ReceiverInternalIdentification               *string `pulumi:"receiverInternalIdentification"`
	ReceiverInternalSubIdentification            *string `pulumi:"receiverInternalSubIdentification"`
	ReceiverReverseRoutingAddress                *string `pulumi:"receiverReverseRoutingAddress"`
	RecipientReferencePasswordQualifier          *string `pulumi:"recipientReferencePasswordQualifier"`
	RecipientReferencePasswordValue              *string `pulumi:"recipientReferencePasswordValue"`
	RolloverGroupControlNumber                   bool    `pulumi:"rolloverGroupControlNumber"`
	RolloverInterchangeControlNumber             bool    `pulumi:"rolloverInterchangeControlNumber"`
	RolloverTransactionSetControlNumber          bool    `pulumi:"rolloverTransactionSetControlNumber"`
	SenderInternalIdentification                 *string `pulumi:"senderInternalIdentification"`
	SenderInternalSubIdentification              *string `pulumi:"senderInternalSubIdentification"`
	SenderReverseRoutingAddress                  *string `pulumi:"senderReverseRoutingAddress"`
	TransactionSetControlNumberLowerBound        float64 `pulumi:"transactionSetControlNumberLowerBound"`
	TransactionSetControlNumberPrefix            *string `pulumi:"transactionSetControlNumberPrefix"`
	TransactionSetControlNumberSuffix            *string `pulumi:"transactionSetControlNumberSuffix"`
	TransactionSetControlNumberUpperBound        float64 `pulumi:"transactionSetControlNumberUpperBound"`
}





type EdifactEnvelopeSettingsInput interface {
	pulumi.Input

	ToEdifactEnvelopeSettingsOutput() EdifactEnvelopeSettingsOutput
	ToEdifactEnvelopeSettingsOutputWithContext(context.Context) EdifactEnvelopeSettingsOutput
}

type EdifactEnvelopeSettingsArgs struct {
	ApplicationReferenceId                       pulumi.StringPtrInput `pulumi:"applicationReferenceId"`
	ApplyDelimiterStringAdvice                   pulumi.BoolInput      `pulumi:"applyDelimiterStringAdvice"`
	CommunicationAgreementId                     pulumi.StringPtrInput `pulumi:"communicationAgreementId"`
	CreateGroupingSegments                       pulumi.BoolInput      `pulumi:"createGroupingSegments"`
	EnableDefaultGroupHeaders                    pulumi.BoolInput      `pulumi:"enableDefaultGroupHeaders"`
	FunctionalGroupId                            pulumi.StringPtrInput `pulumi:"functionalGroupId"`
	GroupApplicationPassword                     pulumi.StringPtrInput `pulumi:"groupApplicationPassword"`
	GroupApplicationReceiverId                   pulumi.StringPtrInput `pulumi:"groupApplicationReceiverId"`
	GroupApplicationReceiverQualifier            pulumi.StringPtrInput `pulumi:"groupApplicationReceiverQualifier"`
	GroupApplicationSenderId                     pulumi.StringPtrInput `pulumi:"groupApplicationSenderId"`
	GroupApplicationSenderQualifier              pulumi.StringPtrInput `pulumi:"groupApplicationSenderQualifier"`
	GroupAssociationAssignedCode                 pulumi.StringPtrInput `pulumi:"groupAssociationAssignedCode"`
	GroupControlNumberLowerBound                 pulumi.Float64Input   `pulumi:"groupControlNumberLowerBound"`
	GroupControlNumberPrefix                     pulumi.StringPtrInput `pulumi:"groupControlNumberPrefix"`
	GroupControlNumberSuffix                     pulumi.StringPtrInput `pulumi:"groupControlNumberSuffix"`
	GroupControlNumberUpperBound                 pulumi.Float64Input   `pulumi:"groupControlNumberUpperBound"`
	GroupControllingAgencyCode                   pulumi.StringPtrInput `pulumi:"groupControllingAgencyCode"`
	GroupMessageRelease                          pulumi.StringPtrInput `pulumi:"groupMessageRelease"`
	GroupMessageVersion                          pulumi.StringPtrInput `pulumi:"groupMessageVersion"`
	InterchangeControlNumberLowerBound           pulumi.Float64Input   `pulumi:"interchangeControlNumberLowerBound"`
	InterchangeControlNumberPrefix               pulumi.StringPtrInput `pulumi:"interchangeControlNumberPrefix"`
	InterchangeControlNumberSuffix               pulumi.StringPtrInput `pulumi:"interchangeControlNumberSuffix"`
	InterchangeControlNumberUpperBound           pulumi.Float64Input   `pulumi:"interchangeControlNumberUpperBound"`
	IsTestInterchange                            pulumi.BoolInput      `pulumi:"isTestInterchange"`
	OverwriteExistingTransactionSetControlNumber pulumi.BoolInput      `pulumi:"overwriteExistingTransactionSetControlNumber"`
	ProcessingPriorityCode                       pulumi.StringPtrInput `pulumi:"processingPriorityCode"`
	ReceiverInternalIdentification               pulumi.StringPtrInput `pulumi:"receiverInternalIdentification"`
	ReceiverInternalSubIdentification            pulumi.StringPtrInput `pulumi:"receiverInternalSubIdentification"`
	ReceiverReverseRoutingAddress                pulumi.StringPtrInput `pulumi:"receiverReverseRoutingAddress"`
	RecipientReferencePasswordQualifier          pulumi.StringPtrInput `pulumi:"recipientReferencePasswordQualifier"`
	RecipientReferencePasswordValue              pulumi.StringPtrInput `pulumi:"recipientReferencePasswordValue"`
	RolloverGroupControlNumber                   pulumi.BoolInput      `pulumi:"rolloverGroupControlNumber"`
	RolloverInterchangeControlNumber             pulumi.BoolInput      `pulumi:"rolloverInterchangeControlNumber"`
	RolloverTransactionSetControlNumber          pulumi.BoolInput      `pulumi:"rolloverTransactionSetControlNumber"`
	SenderInternalIdentification                 pulumi.StringPtrInput `pulumi:"senderInternalIdentification"`
	SenderInternalSubIdentification              pulumi.StringPtrInput `pulumi:"senderInternalSubIdentification"`
	SenderReverseRoutingAddress                  pulumi.StringPtrInput `pulumi:"senderReverseRoutingAddress"`
	TransactionSetControlNumberLowerBound        pulumi.Float64Input   `pulumi:"transactionSetControlNumberLowerBound"`
	TransactionSetControlNumberPrefix            pulumi.StringPtrInput `pulumi:"transactionSetControlNumberPrefix"`
	TransactionSetControlNumberSuffix            pulumi.StringPtrInput `pulumi:"transactionSetControlNumberSuffix"`
	TransactionSetControlNumberUpperBound        pulumi.Float64Input   `pulumi:"transactionSetControlNumberUpperBound"`
}

func (EdifactEnvelopeSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactEnvelopeSettings)(nil)).Elem()
}

func (i EdifactEnvelopeSettingsArgs) ToEdifactEnvelopeSettingsOutput() EdifactEnvelopeSettingsOutput {
	return i.ToEdifactEnvelopeSettingsOutputWithContext(context.Background())
}

func (i EdifactEnvelopeSettingsArgs) ToEdifactEnvelopeSettingsOutputWithContext(ctx context.Context) EdifactEnvelopeSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactEnvelopeSettingsOutput)
}

func (i EdifactEnvelopeSettingsArgs) ToEdifactEnvelopeSettingsPtrOutput() EdifactEnvelopeSettingsPtrOutput {
	return i.ToEdifactEnvelopeSettingsPtrOutputWithContext(context.Background())
}

func (i EdifactEnvelopeSettingsArgs) ToEdifactEnvelopeSettingsPtrOutputWithContext(ctx context.Context) EdifactEnvelopeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactEnvelopeSettingsOutput).ToEdifactEnvelopeSettingsPtrOutputWithContext(ctx)
}









type EdifactEnvelopeSettingsPtrInput interface {
	pulumi.Input

	ToEdifactEnvelopeSettingsPtrOutput() EdifactEnvelopeSettingsPtrOutput
	ToEdifactEnvelopeSettingsPtrOutputWithContext(context.Context) EdifactEnvelopeSettingsPtrOutput
}

type edifactEnvelopeSettingsPtrType EdifactEnvelopeSettingsArgs

func EdifactEnvelopeSettingsPtr(v *EdifactEnvelopeSettingsArgs) EdifactEnvelopeSettingsPtrInput {
	return (*edifactEnvelopeSettingsPtrType)(v)
}

func (*edifactEnvelopeSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactEnvelopeSettings)(nil)).Elem()
}

func (i *edifactEnvelopeSettingsPtrType) ToEdifactEnvelopeSettingsPtrOutput() EdifactEnvelopeSettingsPtrOutput {
	return i.ToEdifactEnvelopeSettingsPtrOutputWithContext(context.Background())
}

func (i *edifactEnvelopeSettingsPtrType) ToEdifactEnvelopeSettingsPtrOutputWithContext(ctx context.Context) EdifactEnvelopeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactEnvelopeSettingsPtrOutput)
}

type EdifactEnvelopeSettingsOutput struct{ *pulumi.OutputState }

func (EdifactEnvelopeSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactEnvelopeSettings)(nil)).Elem()
}

func (o EdifactEnvelopeSettingsOutput) ToEdifactEnvelopeSettingsOutput() EdifactEnvelopeSettingsOutput {
	return o
}

func (o EdifactEnvelopeSettingsOutput) ToEdifactEnvelopeSettingsOutputWithContext(ctx context.Context) EdifactEnvelopeSettingsOutput {
	return o
}

func (o EdifactEnvelopeSettingsOutput) ToEdifactEnvelopeSettingsPtrOutput() EdifactEnvelopeSettingsPtrOutput {
	return o.ToEdifactEnvelopeSettingsPtrOutputWithContext(context.Background())
}

func (o EdifactEnvelopeSettingsOutput) ToEdifactEnvelopeSettingsPtrOutputWithContext(ctx context.Context) EdifactEnvelopeSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactEnvelopeSettings) *EdifactEnvelopeSettings {
		return &v
	}).(EdifactEnvelopeSettingsPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) ApplicationReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.ApplicationReferenceId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) ApplyDelimiterStringAdvice() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) bool { return v.ApplyDelimiterStringAdvice }).(pulumi.BoolOutput)
}

func (o EdifactEnvelopeSettingsOutput) CommunicationAgreementId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.CommunicationAgreementId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) CreateGroupingSegments() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) bool { return v.CreateGroupingSegments }).(pulumi.BoolOutput)
}

func (o EdifactEnvelopeSettingsOutput) EnableDefaultGroupHeaders() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) bool { return v.EnableDefaultGroupHeaders }).(pulumi.BoolOutput)
}

func (o EdifactEnvelopeSettingsOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.FunctionalGroupId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupApplicationPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupApplicationPassword }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupApplicationReceiverId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupApplicationReceiverId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupApplicationReceiverQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupApplicationReceiverQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupApplicationSenderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupApplicationSenderId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupApplicationSenderQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupApplicationSenderQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupAssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupAssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupControlNumberLowerBound() pulumi.Float64Output {
	return o.ApplyT(func(v EdifactEnvelopeSettings) float64 { return v.GroupControlNumberLowerBound }).(pulumi.Float64Output)
}

func (o EdifactEnvelopeSettingsOutput) GroupControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupControlNumberUpperBound() pulumi.Float64Output {
	return o.ApplyT(func(v EdifactEnvelopeSettings) float64 { return v.GroupControlNumberUpperBound }).(pulumi.Float64Output)
}

func (o EdifactEnvelopeSettingsOutput) GroupControllingAgencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupControllingAgencyCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupMessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupMessageRelease }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) GroupMessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.GroupMessageVersion }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) InterchangeControlNumberLowerBound() pulumi.Float64Output {
	return o.ApplyT(func(v EdifactEnvelopeSettings) float64 { return v.InterchangeControlNumberLowerBound }).(pulumi.Float64Output)
}

func (o EdifactEnvelopeSettingsOutput) InterchangeControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.InterchangeControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) InterchangeControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.InterchangeControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) InterchangeControlNumberUpperBound() pulumi.Float64Output {
	return o.ApplyT(func(v EdifactEnvelopeSettings) float64 { return v.InterchangeControlNumberUpperBound }).(pulumi.Float64Output)
}

func (o EdifactEnvelopeSettingsOutput) IsTestInterchange() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) bool { return v.IsTestInterchange }).(pulumi.BoolOutput)
}

func (o EdifactEnvelopeSettingsOutput) OverwriteExistingTransactionSetControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) bool { return v.OverwriteExistingTransactionSetControlNumber }).(pulumi.BoolOutput)
}

func (o EdifactEnvelopeSettingsOutput) ProcessingPriorityCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.ProcessingPriorityCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) ReceiverInternalIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.ReceiverInternalIdentification }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) ReceiverInternalSubIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.ReceiverInternalSubIdentification }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) ReceiverReverseRoutingAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.ReceiverReverseRoutingAddress }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) RecipientReferencePasswordQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.RecipientReferencePasswordQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) RecipientReferencePasswordValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.RecipientReferencePasswordValue }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) RolloverGroupControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) bool { return v.RolloverGroupControlNumber }).(pulumi.BoolOutput)
}

func (o EdifactEnvelopeSettingsOutput) RolloverInterchangeControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) bool { return v.RolloverInterchangeControlNumber }).(pulumi.BoolOutput)
}

func (o EdifactEnvelopeSettingsOutput) RolloverTransactionSetControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) bool { return v.RolloverTransactionSetControlNumber }).(pulumi.BoolOutput)
}

func (o EdifactEnvelopeSettingsOutput) SenderInternalIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.SenderInternalIdentification }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) SenderInternalSubIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.SenderInternalSubIdentification }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) SenderReverseRoutingAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.SenderReverseRoutingAddress }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) TransactionSetControlNumberLowerBound() pulumi.Float64Output {
	return o.ApplyT(func(v EdifactEnvelopeSettings) float64 { return v.TransactionSetControlNumberLowerBound }).(pulumi.Float64Output)
}

func (o EdifactEnvelopeSettingsOutput) TransactionSetControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.TransactionSetControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) TransactionSetControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettings) *string { return v.TransactionSetControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsOutput) TransactionSetControlNumberUpperBound() pulumi.Float64Output {
	return o.ApplyT(func(v EdifactEnvelopeSettings) float64 { return v.TransactionSetControlNumberUpperBound }).(pulumi.Float64Output)
}

type EdifactEnvelopeSettingsPtrOutput struct{ *pulumi.OutputState }

func (EdifactEnvelopeSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactEnvelopeSettings)(nil)).Elem()
}

func (o EdifactEnvelopeSettingsPtrOutput) ToEdifactEnvelopeSettingsPtrOutput() EdifactEnvelopeSettingsPtrOutput {
	return o
}

func (o EdifactEnvelopeSettingsPtrOutput) ToEdifactEnvelopeSettingsPtrOutputWithContext(ctx context.Context) EdifactEnvelopeSettingsPtrOutput {
	return o
}

func (o EdifactEnvelopeSettingsPtrOutput) Elem() EdifactEnvelopeSettingsOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) EdifactEnvelopeSettings {
		if v != nil {
			return *v
		}
		var ret EdifactEnvelopeSettings
		return ret
	}).(EdifactEnvelopeSettingsOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) ApplicationReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationReferenceId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) ApplyDelimiterStringAdvice() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.ApplyDelimiterStringAdvice
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) CommunicationAgreementId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.CommunicationAgreementId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) CreateGroupingSegments() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.CreateGroupingSegments
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) EnableDefaultGroupHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.EnableDefaultGroupHeaders
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.FunctionalGroupId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupApplicationPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationPassword
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupApplicationReceiverId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationReceiverId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupApplicationReceiverQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationReceiverQualifier
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupApplicationSenderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationSenderId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupApplicationSenderQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationSenderQualifier
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupAssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupAssociationAssignedCode
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupControlNumberLowerBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *float64 {
		if v == nil {
			return nil
		}
		return &v.GroupControlNumberLowerBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupControlNumberUpperBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *float64 {
		if v == nil {
			return nil
		}
		return &v.GroupControlNumberUpperBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupControllingAgencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupControllingAgencyCode
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupMessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupMessageRelease
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) GroupMessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.GroupMessageVersion
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) InterchangeControlNumberLowerBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *float64 {
		if v == nil {
			return nil
		}
		return &v.InterchangeControlNumberLowerBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) InterchangeControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) InterchangeControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) InterchangeControlNumberUpperBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *float64 {
		if v == nil {
			return nil
		}
		return &v.InterchangeControlNumberUpperBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) IsTestInterchange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.IsTestInterchange
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) OverwriteExistingTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.OverwriteExistingTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) ProcessingPriorityCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.ProcessingPriorityCode
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) ReceiverInternalIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.ReceiverInternalIdentification
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) ReceiverInternalSubIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.ReceiverInternalSubIdentification
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) ReceiverReverseRoutingAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.ReceiverReverseRoutingAddress
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) RecipientReferencePasswordQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.RecipientReferencePasswordQualifier
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) RecipientReferencePasswordValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.RecipientReferencePasswordValue
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) RolloverGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.RolloverGroupControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) RolloverInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.RolloverInterchangeControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) RolloverTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.RolloverTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) SenderInternalIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.SenderInternalIdentification
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) SenderInternalSubIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.SenderInternalSubIdentification
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) SenderReverseRoutingAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.SenderReverseRoutingAddress
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) TransactionSetControlNumberLowerBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *float64 {
		if v == nil {
			return nil
		}
		return &v.TransactionSetControlNumberLowerBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) TransactionSetControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) TransactionSetControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsPtrOutput) TransactionSetControlNumberUpperBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettings) *float64 {
		if v == nil {
			return nil
		}
		return &v.TransactionSetControlNumberUpperBound
	}).(pulumi.Float64PtrOutput)
}

type EdifactEnvelopeSettingsResponse struct {
	ApplicationReferenceId                       *string `pulumi:"applicationReferenceId"`
	ApplyDelimiterStringAdvice                   bool    `pulumi:"applyDelimiterStringAdvice"`
	CommunicationAgreementId                     *string `pulumi:"communicationAgreementId"`
	CreateGroupingSegments                       bool    `pulumi:"createGroupingSegments"`
	EnableDefaultGroupHeaders                    bool    `pulumi:"enableDefaultGroupHeaders"`
	FunctionalGroupId                            *string `pulumi:"functionalGroupId"`
	GroupApplicationPassword                     *string `pulumi:"groupApplicationPassword"`
	GroupApplicationReceiverId                   *string `pulumi:"groupApplicationReceiverId"`
	GroupApplicationReceiverQualifier            *string `pulumi:"groupApplicationReceiverQualifier"`
	GroupApplicationSenderId                     *string `pulumi:"groupApplicationSenderId"`
	GroupApplicationSenderQualifier              *string `pulumi:"groupApplicationSenderQualifier"`
	GroupAssociationAssignedCode                 *string `pulumi:"groupAssociationAssignedCode"`
	GroupControlNumberLowerBound                 float64 `pulumi:"groupControlNumberLowerBound"`
	GroupControlNumberPrefix                     *string `pulumi:"groupControlNumberPrefix"`
	GroupControlNumberSuffix                     *string `pulumi:"groupControlNumberSuffix"`
	GroupControlNumberUpperBound                 float64 `pulumi:"groupControlNumberUpperBound"`
	GroupControllingAgencyCode                   *string `pulumi:"groupControllingAgencyCode"`
	GroupMessageRelease                          *string `pulumi:"groupMessageRelease"`
	GroupMessageVersion                          *string `pulumi:"groupMessageVersion"`
	InterchangeControlNumberLowerBound           float64 `pulumi:"interchangeControlNumberLowerBound"`
	InterchangeControlNumberPrefix               *string `pulumi:"interchangeControlNumberPrefix"`
	InterchangeControlNumberSuffix               *string `pulumi:"interchangeControlNumberSuffix"`
	InterchangeControlNumberUpperBound           float64 `pulumi:"interchangeControlNumberUpperBound"`
	IsTestInterchange                            bool    `pulumi:"isTestInterchange"`
	OverwriteExistingTransactionSetControlNumber bool    `pulumi:"overwriteExistingTransactionSetControlNumber"`
	ProcessingPriorityCode                       *string `pulumi:"processingPriorityCode"`
	ReceiverInternalIdentification               *string `pulumi:"receiverInternalIdentification"`
	ReceiverInternalSubIdentification            *string `pulumi:"receiverInternalSubIdentification"`
	ReceiverReverseRoutingAddress                *string `pulumi:"receiverReverseRoutingAddress"`
	RecipientReferencePasswordQualifier          *string `pulumi:"recipientReferencePasswordQualifier"`
	RecipientReferencePasswordValue              *string `pulumi:"recipientReferencePasswordValue"`
	RolloverGroupControlNumber                   bool    `pulumi:"rolloverGroupControlNumber"`
	RolloverInterchangeControlNumber             bool    `pulumi:"rolloverInterchangeControlNumber"`
	RolloverTransactionSetControlNumber          bool    `pulumi:"rolloverTransactionSetControlNumber"`
	SenderInternalIdentification                 *string `pulumi:"senderInternalIdentification"`
	SenderInternalSubIdentification              *string `pulumi:"senderInternalSubIdentification"`
	SenderReverseRoutingAddress                  *string `pulumi:"senderReverseRoutingAddress"`
	TransactionSetControlNumberLowerBound        float64 `pulumi:"transactionSetControlNumberLowerBound"`
	TransactionSetControlNumberPrefix            *string `pulumi:"transactionSetControlNumberPrefix"`
	TransactionSetControlNumberSuffix            *string `pulumi:"transactionSetControlNumberSuffix"`
	TransactionSetControlNumberUpperBound        float64 `pulumi:"transactionSetControlNumberUpperBound"`
}

type EdifactEnvelopeSettingsResponseOutput struct{ *pulumi.OutputState }

func (EdifactEnvelopeSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactEnvelopeSettingsResponse)(nil)).Elem()
}

func (o EdifactEnvelopeSettingsResponseOutput) ToEdifactEnvelopeSettingsResponseOutput() EdifactEnvelopeSettingsResponseOutput {
	return o
}

func (o EdifactEnvelopeSettingsResponseOutput) ToEdifactEnvelopeSettingsResponseOutputWithContext(ctx context.Context) EdifactEnvelopeSettingsResponseOutput {
	return o
}

func (o EdifactEnvelopeSettingsResponseOutput) ApplicationReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.ApplicationReferenceId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) ApplyDelimiterStringAdvice() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) bool { return v.ApplyDelimiterStringAdvice }).(pulumi.BoolOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) CommunicationAgreementId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.CommunicationAgreementId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) CreateGroupingSegments() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) bool { return v.CreateGroupingSegments }).(pulumi.BoolOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) EnableDefaultGroupHeaders() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) bool { return v.EnableDefaultGroupHeaders }).(pulumi.BoolOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.FunctionalGroupId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupApplicationPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupApplicationPassword }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupApplicationReceiverId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupApplicationReceiverId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupApplicationReceiverQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupApplicationReceiverQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupApplicationSenderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupApplicationSenderId }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupApplicationSenderQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupApplicationSenderQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupAssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupAssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupControlNumberLowerBound() pulumi.Float64Output {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) float64 { return v.GroupControlNumberLowerBound }).(pulumi.Float64Output)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupControlNumberUpperBound() pulumi.Float64Output {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) float64 { return v.GroupControlNumberUpperBound }).(pulumi.Float64Output)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupControllingAgencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupControllingAgencyCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupMessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupMessageRelease }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) GroupMessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.GroupMessageVersion }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) InterchangeControlNumberLowerBound() pulumi.Float64Output {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) float64 { return v.InterchangeControlNumberLowerBound }).(pulumi.Float64Output)
}

func (o EdifactEnvelopeSettingsResponseOutput) InterchangeControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.InterchangeControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) InterchangeControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.InterchangeControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) InterchangeControlNumberUpperBound() pulumi.Float64Output {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) float64 { return v.InterchangeControlNumberUpperBound }).(pulumi.Float64Output)
}

func (o EdifactEnvelopeSettingsResponseOutput) IsTestInterchange() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) bool { return v.IsTestInterchange }).(pulumi.BoolOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) OverwriteExistingTransactionSetControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) bool { return v.OverwriteExistingTransactionSetControlNumber }).(pulumi.BoolOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) ProcessingPriorityCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.ProcessingPriorityCode }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) ReceiverInternalIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.ReceiverInternalIdentification }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) ReceiverInternalSubIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.ReceiverInternalSubIdentification }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) ReceiverReverseRoutingAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.ReceiverReverseRoutingAddress }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) RecipientReferencePasswordQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.RecipientReferencePasswordQualifier }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) RecipientReferencePasswordValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.RecipientReferencePasswordValue }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) RolloverGroupControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) bool { return v.RolloverGroupControlNumber }).(pulumi.BoolOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) RolloverInterchangeControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) bool { return v.RolloverInterchangeControlNumber }).(pulumi.BoolOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) RolloverTransactionSetControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) bool { return v.RolloverTransactionSetControlNumber }).(pulumi.BoolOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) SenderInternalIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.SenderInternalIdentification }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) SenderInternalSubIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.SenderInternalSubIdentification }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) SenderReverseRoutingAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.SenderReverseRoutingAddress }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) TransactionSetControlNumberLowerBound() pulumi.Float64Output {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) float64 { return v.TransactionSetControlNumberLowerBound }).(pulumi.Float64Output)
}

func (o EdifactEnvelopeSettingsResponseOutput) TransactionSetControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.TransactionSetControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) TransactionSetControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) *string { return v.TransactionSetControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponseOutput) TransactionSetControlNumberUpperBound() pulumi.Float64Output {
	return o.ApplyT(func(v EdifactEnvelopeSettingsResponse) float64 { return v.TransactionSetControlNumberUpperBound }).(pulumi.Float64Output)
}

type EdifactEnvelopeSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (EdifactEnvelopeSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactEnvelopeSettingsResponse)(nil)).Elem()
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) ToEdifactEnvelopeSettingsResponsePtrOutput() EdifactEnvelopeSettingsResponsePtrOutput {
	return o
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) ToEdifactEnvelopeSettingsResponsePtrOutputWithContext(ctx context.Context) EdifactEnvelopeSettingsResponsePtrOutput {
	return o
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) Elem() EdifactEnvelopeSettingsResponseOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) EdifactEnvelopeSettingsResponse {
		if v != nil {
			return *v
		}
		var ret EdifactEnvelopeSettingsResponse
		return ret
	}).(EdifactEnvelopeSettingsResponseOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) ApplicationReferenceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationReferenceId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) ApplyDelimiterStringAdvice() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.ApplyDelimiterStringAdvice
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) CommunicationAgreementId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.CommunicationAgreementId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) CreateGroupingSegments() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CreateGroupingSegments
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) EnableDefaultGroupHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.EnableDefaultGroupHeaders
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.FunctionalGroupId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupApplicationPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationPassword
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupApplicationReceiverId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationReceiverId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupApplicationReceiverQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationReceiverQualifier
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupApplicationSenderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationSenderId
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupApplicationSenderQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupApplicationSenderQualifier
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupAssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupAssociationAssignedCode
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupControlNumberLowerBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.GroupControlNumberLowerBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupControlNumberUpperBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.GroupControlNumberUpperBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupControllingAgencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupControllingAgencyCode
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupMessageRelease() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupMessageRelease
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) GroupMessageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupMessageVersion
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) InterchangeControlNumberLowerBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.InterchangeControlNumberLowerBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) InterchangeControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) InterchangeControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.InterchangeControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) InterchangeControlNumberUpperBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.InterchangeControlNumberUpperBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) IsTestInterchange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsTestInterchange
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) OverwriteExistingTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.OverwriteExistingTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) ProcessingPriorityCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProcessingPriorityCode
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) ReceiverInternalIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReceiverInternalIdentification
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) ReceiverInternalSubIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReceiverInternalSubIdentification
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) ReceiverReverseRoutingAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReceiverReverseRoutingAddress
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) RecipientReferencePasswordQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.RecipientReferencePasswordQualifier
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) RecipientReferencePasswordValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.RecipientReferencePasswordValue
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) RolloverGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.RolloverGroupControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) RolloverInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.RolloverInterchangeControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) RolloverTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.RolloverTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) SenderInternalIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SenderInternalIdentification
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) SenderInternalSubIdentification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SenderInternalSubIdentification
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) SenderReverseRoutingAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SenderReverseRoutingAddress
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) TransactionSetControlNumberLowerBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.TransactionSetControlNumberLowerBound
	}).(pulumi.Float64PtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) TransactionSetControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) TransactionSetControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactEnvelopeSettingsResponsePtrOutput) TransactionSetControlNumberUpperBound() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *EdifactEnvelopeSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.TransactionSetControlNumberUpperBound
	}).(pulumi.Float64PtrOutput)
}

type EdifactFramingSettings struct {
	CharacterEncoding               *string                 `pulumi:"characterEncoding"`
	CharacterSet                    string                  `pulumi:"characterSet"`
	ComponentSeparator              int                     `pulumi:"componentSeparator"`
	DataElementSeparator            int                     `pulumi:"dataElementSeparator"`
	DecimalPointIndicator           EdifactDecimalIndicator `pulumi:"decimalPointIndicator"`
	ProtocolVersion                 int                     `pulumi:"protocolVersion"`
	ReleaseIndicator                int                     `pulumi:"releaseIndicator"`
	RepetitionSeparator             int                     `pulumi:"repetitionSeparator"`
	SegmentTerminator               int                     `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix         SegmentTerminatorSuffix `pulumi:"segmentTerminatorSuffix"`
	ServiceCodeListDirectoryVersion *string                 `pulumi:"serviceCodeListDirectoryVersion"`
}





type EdifactFramingSettingsInput interface {
	pulumi.Input

	ToEdifactFramingSettingsOutput() EdifactFramingSettingsOutput
	ToEdifactFramingSettingsOutputWithContext(context.Context) EdifactFramingSettingsOutput
}

type EdifactFramingSettingsArgs struct {
	CharacterEncoding               pulumi.StringPtrInput        `pulumi:"characterEncoding"`
	CharacterSet                    pulumi.StringInput           `pulumi:"characterSet"`
	ComponentSeparator              pulumi.IntInput              `pulumi:"componentSeparator"`
	DataElementSeparator            pulumi.IntInput              `pulumi:"dataElementSeparator"`
	DecimalPointIndicator           EdifactDecimalIndicatorInput `pulumi:"decimalPointIndicator"`
	ProtocolVersion                 pulumi.IntInput              `pulumi:"protocolVersion"`
	ReleaseIndicator                pulumi.IntInput              `pulumi:"releaseIndicator"`
	RepetitionSeparator             pulumi.IntInput              `pulumi:"repetitionSeparator"`
	SegmentTerminator               pulumi.IntInput              `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix         SegmentTerminatorSuffixInput `pulumi:"segmentTerminatorSuffix"`
	ServiceCodeListDirectoryVersion pulumi.StringPtrInput        `pulumi:"serviceCodeListDirectoryVersion"`
}

func (EdifactFramingSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactFramingSettings)(nil)).Elem()
}

func (i EdifactFramingSettingsArgs) ToEdifactFramingSettingsOutput() EdifactFramingSettingsOutput {
	return i.ToEdifactFramingSettingsOutputWithContext(context.Background())
}

func (i EdifactFramingSettingsArgs) ToEdifactFramingSettingsOutputWithContext(ctx context.Context) EdifactFramingSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactFramingSettingsOutput)
}

func (i EdifactFramingSettingsArgs) ToEdifactFramingSettingsPtrOutput() EdifactFramingSettingsPtrOutput {
	return i.ToEdifactFramingSettingsPtrOutputWithContext(context.Background())
}

func (i EdifactFramingSettingsArgs) ToEdifactFramingSettingsPtrOutputWithContext(ctx context.Context) EdifactFramingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactFramingSettingsOutput).ToEdifactFramingSettingsPtrOutputWithContext(ctx)
}









type EdifactFramingSettingsPtrInput interface {
	pulumi.Input

	ToEdifactFramingSettingsPtrOutput() EdifactFramingSettingsPtrOutput
	ToEdifactFramingSettingsPtrOutputWithContext(context.Context) EdifactFramingSettingsPtrOutput
}

type edifactFramingSettingsPtrType EdifactFramingSettingsArgs

func EdifactFramingSettingsPtr(v *EdifactFramingSettingsArgs) EdifactFramingSettingsPtrInput {
	return (*edifactFramingSettingsPtrType)(v)
}

func (*edifactFramingSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactFramingSettings)(nil)).Elem()
}

func (i *edifactFramingSettingsPtrType) ToEdifactFramingSettingsPtrOutput() EdifactFramingSettingsPtrOutput {
	return i.ToEdifactFramingSettingsPtrOutputWithContext(context.Background())
}

func (i *edifactFramingSettingsPtrType) ToEdifactFramingSettingsPtrOutputWithContext(ctx context.Context) EdifactFramingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactFramingSettingsPtrOutput)
}

type EdifactFramingSettingsOutput struct{ *pulumi.OutputState }

func (EdifactFramingSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactFramingSettings)(nil)).Elem()
}

func (o EdifactFramingSettingsOutput) ToEdifactFramingSettingsOutput() EdifactFramingSettingsOutput {
	return o
}

func (o EdifactFramingSettingsOutput) ToEdifactFramingSettingsOutputWithContext(ctx context.Context) EdifactFramingSettingsOutput {
	return o
}

func (o EdifactFramingSettingsOutput) ToEdifactFramingSettingsPtrOutput() EdifactFramingSettingsPtrOutput {
	return o.ToEdifactFramingSettingsPtrOutputWithContext(context.Background())
}

func (o EdifactFramingSettingsOutput) ToEdifactFramingSettingsPtrOutputWithContext(ctx context.Context) EdifactFramingSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactFramingSettings) *EdifactFramingSettings {
		return &v
	}).(EdifactFramingSettingsPtrOutput)
}

func (o EdifactFramingSettingsOutput) CharacterEncoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettings) *string { return v.CharacterEncoding }).(pulumi.StringPtrOutput)
}

func (o EdifactFramingSettingsOutput) CharacterSet() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactFramingSettings) string { return v.CharacterSet }).(pulumi.StringOutput)
}

func (o EdifactFramingSettingsOutput) ComponentSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactFramingSettings) int { return v.ComponentSeparator }).(pulumi.IntOutput)
}

func (o EdifactFramingSettingsOutput) DataElementSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactFramingSettings) int { return v.DataElementSeparator }).(pulumi.IntOutput)
}

func (o EdifactFramingSettingsOutput) DecimalPointIndicator() EdifactDecimalIndicatorOutput {
	return o.ApplyT(func(v EdifactFramingSettings) EdifactDecimalIndicator { return v.DecimalPointIndicator }).(EdifactDecimalIndicatorOutput)
}

func (o EdifactFramingSettingsOutput) ProtocolVersion() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactFramingSettings) int { return v.ProtocolVersion }).(pulumi.IntOutput)
}

func (o EdifactFramingSettingsOutput) ReleaseIndicator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactFramingSettings) int { return v.ReleaseIndicator }).(pulumi.IntOutput)
}

func (o EdifactFramingSettingsOutput) RepetitionSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactFramingSettings) int { return v.RepetitionSeparator }).(pulumi.IntOutput)
}

func (o EdifactFramingSettingsOutput) SegmentTerminator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactFramingSettings) int { return v.SegmentTerminator }).(pulumi.IntOutput)
}

func (o EdifactFramingSettingsOutput) SegmentTerminatorSuffix() SegmentTerminatorSuffixOutput {
	return o.ApplyT(func(v EdifactFramingSettings) SegmentTerminatorSuffix { return v.SegmentTerminatorSuffix }).(SegmentTerminatorSuffixOutput)
}

func (o EdifactFramingSettingsOutput) ServiceCodeListDirectoryVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettings) *string { return v.ServiceCodeListDirectoryVersion }).(pulumi.StringPtrOutput)
}

type EdifactFramingSettingsPtrOutput struct{ *pulumi.OutputState }

func (EdifactFramingSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactFramingSettings)(nil)).Elem()
}

func (o EdifactFramingSettingsPtrOutput) ToEdifactFramingSettingsPtrOutput() EdifactFramingSettingsPtrOutput {
	return o
}

func (o EdifactFramingSettingsPtrOutput) ToEdifactFramingSettingsPtrOutputWithContext(ctx context.Context) EdifactFramingSettingsPtrOutput {
	return o
}

func (o EdifactFramingSettingsPtrOutput) Elem() EdifactFramingSettingsOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) EdifactFramingSettings {
		if v != nil {
			return *v
		}
		var ret EdifactFramingSettings
		return ret
	}).(EdifactFramingSettingsOutput)
}

func (o EdifactFramingSettingsPtrOutput) CharacterEncoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *string {
		if v == nil {
			return nil
		}
		return v.CharacterEncoding
	}).(pulumi.StringPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) CharacterSet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *string {
		if v == nil {
			return nil
		}
		return &v.CharacterSet
	}).(pulumi.StringPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) ComponentSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *int {
		if v == nil {
			return nil
		}
		return &v.ComponentSeparator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) DataElementSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *int {
		if v == nil {
			return nil
		}
		return &v.DataElementSeparator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) DecimalPointIndicator() EdifactDecimalIndicatorPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *EdifactDecimalIndicator {
		if v == nil {
			return nil
		}
		return &v.DecimalPointIndicator
	}).(EdifactDecimalIndicatorPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) ProtocolVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *int {
		if v == nil {
			return nil
		}
		return &v.ProtocolVersion
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) ReleaseIndicator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *int {
		if v == nil {
			return nil
		}
		return &v.ReleaseIndicator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) RepetitionSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *int {
		if v == nil {
			return nil
		}
		return &v.RepetitionSeparator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) SegmentTerminator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *int {
		if v == nil {
			return nil
		}
		return &v.SegmentTerminator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) SegmentTerminatorSuffix() SegmentTerminatorSuffixPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *SegmentTerminatorSuffix {
		if v == nil {
			return nil
		}
		return &v.SegmentTerminatorSuffix
	}).(SegmentTerminatorSuffixPtrOutput)
}

func (o EdifactFramingSettingsPtrOutput) ServiceCodeListDirectoryVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettings) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCodeListDirectoryVersion
	}).(pulumi.StringPtrOutput)
}

type EdifactFramingSettingsResponse struct {
	CharacterEncoding               *string `pulumi:"characterEncoding"`
	CharacterSet                    string  `pulumi:"characterSet"`
	ComponentSeparator              int     `pulumi:"componentSeparator"`
	DataElementSeparator            int     `pulumi:"dataElementSeparator"`
	DecimalPointIndicator           string  `pulumi:"decimalPointIndicator"`
	ProtocolVersion                 int     `pulumi:"protocolVersion"`
	ReleaseIndicator                int     `pulumi:"releaseIndicator"`
	RepetitionSeparator             int     `pulumi:"repetitionSeparator"`
	SegmentTerminator               int     `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix         string  `pulumi:"segmentTerminatorSuffix"`
	ServiceCodeListDirectoryVersion *string `pulumi:"serviceCodeListDirectoryVersion"`
}

type EdifactFramingSettingsResponseOutput struct{ *pulumi.OutputState }

func (EdifactFramingSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactFramingSettingsResponse)(nil)).Elem()
}

func (o EdifactFramingSettingsResponseOutput) ToEdifactFramingSettingsResponseOutput() EdifactFramingSettingsResponseOutput {
	return o
}

func (o EdifactFramingSettingsResponseOutput) ToEdifactFramingSettingsResponseOutputWithContext(ctx context.Context) EdifactFramingSettingsResponseOutput {
	return o
}

func (o EdifactFramingSettingsResponseOutput) CharacterEncoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) *string { return v.CharacterEncoding }).(pulumi.StringPtrOutput)
}

func (o EdifactFramingSettingsResponseOutput) CharacterSet() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) string { return v.CharacterSet }).(pulumi.StringOutput)
}

func (o EdifactFramingSettingsResponseOutput) ComponentSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) int { return v.ComponentSeparator }).(pulumi.IntOutput)
}

func (o EdifactFramingSettingsResponseOutput) DataElementSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) int { return v.DataElementSeparator }).(pulumi.IntOutput)
}

func (o EdifactFramingSettingsResponseOutput) DecimalPointIndicator() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) string { return v.DecimalPointIndicator }).(pulumi.StringOutput)
}

func (o EdifactFramingSettingsResponseOutput) ProtocolVersion() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) int { return v.ProtocolVersion }).(pulumi.IntOutput)
}

func (o EdifactFramingSettingsResponseOutput) ReleaseIndicator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) int { return v.ReleaseIndicator }).(pulumi.IntOutput)
}

func (o EdifactFramingSettingsResponseOutput) RepetitionSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) int { return v.RepetitionSeparator }).(pulumi.IntOutput)
}

func (o EdifactFramingSettingsResponseOutput) SegmentTerminator() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) int { return v.SegmentTerminator }).(pulumi.IntOutput)
}

func (o EdifactFramingSettingsResponseOutput) SegmentTerminatorSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) string { return v.SegmentTerminatorSuffix }).(pulumi.StringOutput)
}

func (o EdifactFramingSettingsResponseOutput) ServiceCodeListDirectoryVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactFramingSettingsResponse) *string { return v.ServiceCodeListDirectoryVersion }).(pulumi.StringPtrOutput)
}

type EdifactFramingSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (EdifactFramingSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactFramingSettingsResponse)(nil)).Elem()
}

func (o EdifactFramingSettingsResponsePtrOutput) ToEdifactFramingSettingsResponsePtrOutput() EdifactFramingSettingsResponsePtrOutput {
	return o
}

func (o EdifactFramingSettingsResponsePtrOutput) ToEdifactFramingSettingsResponsePtrOutputWithContext(ctx context.Context) EdifactFramingSettingsResponsePtrOutput {
	return o
}

func (o EdifactFramingSettingsResponsePtrOutput) Elem() EdifactFramingSettingsResponseOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) EdifactFramingSettingsResponse {
		if v != nil {
			return *v
		}
		var ret EdifactFramingSettingsResponse
		return ret
	}).(EdifactFramingSettingsResponseOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) CharacterEncoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.CharacterEncoding
	}).(pulumi.StringPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) CharacterSet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CharacterSet
	}).(pulumi.StringPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) ComponentSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ComponentSeparator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) DataElementSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.DataElementSeparator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) DecimalPointIndicator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DecimalPointIndicator
	}).(pulumi.StringPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) ProtocolVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ProtocolVersion
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) ReleaseIndicator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ReleaseIndicator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) RepetitionSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.RepetitionSeparator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) SegmentTerminator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.SegmentTerminator
	}).(pulumi.IntPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) SegmentTerminatorSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SegmentTerminatorSuffix
	}).(pulumi.StringPtrOutput)
}

func (o EdifactFramingSettingsResponsePtrOutput) ServiceCodeListDirectoryVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactFramingSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCodeListDirectoryVersion
	}).(pulumi.StringPtrOutput)
}

type EdifactMessageFilter struct {
	MessageFilterType string `pulumi:"messageFilterType"`
}





type EdifactMessageFilterInput interface {
	pulumi.Input

	ToEdifactMessageFilterOutput() EdifactMessageFilterOutput
	ToEdifactMessageFilterOutputWithContext(context.Context) EdifactMessageFilterOutput
}

type EdifactMessageFilterArgs struct {
	MessageFilterType pulumi.StringInput `pulumi:"messageFilterType"`
}

func (EdifactMessageFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactMessageFilter)(nil)).Elem()
}

func (i EdifactMessageFilterArgs) ToEdifactMessageFilterOutput() EdifactMessageFilterOutput {
	return i.ToEdifactMessageFilterOutputWithContext(context.Background())
}

func (i EdifactMessageFilterArgs) ToEdifactMessageFilterOutputWithContext(ctx context.Context) EdifactMessageFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactMessageFilterOutput)
}

func (i EdifactMessageFilterArgs) ToEdifactMessageFilterPtrOutput() EdifactMessageFilterPtrOutput {
	return i.ToEdifactMessageFilterPtrOutputWithContext(context.Background())
}

func (i EdifactMessageFilterArgs) ToEdifactMessageFilterPtrOutputWithContext(ctx context.Context) EdifactMessageFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactMessageFilterOutput).ToEdifactMessageFilterPtrOutputWithContext(ctx)
}









type EdifactMessageFilterPtrInput interface {
	pulumi.Input

	ToEdifactMessageFilterPtrOutput() EdifactMessageFilterPtrOutput
	ToEdifactMessageFilterPtrOutputWithContext(context.Context) EdifactMessageFilterPtrOutput
}

type edifactMessageFilterPtrType EdifactMessageFilterArgs

func EdifactMessageFilterPtr(v *EdifactMessageFilterArgs) EdifactMessageFilterPtrInput {
	return (*edifactMessageFilterPtrType)(v)
}

func (*edifactMessageFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactMessageFilter)(nil)).Elem()
}

func (i *edifactMessageFilterPtrType) ToEdifactMessageFilterPtrOutput() EdifactMessageFilterPtrOutput {
	return i.ToEdifactMessageFilterPtrOutputWithContext(context.Background())
}

func (i *edifactMessageFilterPtrType) ToEdifactMessageFilterPtrOutputWithContext(ctx context.Context) EdifactMessageFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactMessageFilterPtrOutput)
}

type EdifactMessageFilterOutput struct{ *pulumi.OutputState }

func (EdifactMessageFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactMessageFilter)(nil)).Elem()
}

func (o EdifactMessageFilterOutput) ToEdifactMessageFilterOutput() EdifactMessageFilterOutput {
	return o
}

func (o EdifactMessageFilterOutput) ToEdifactMessageFilterOutputWithContext(ctx context.Context) EdifactMessageFilterOutput {
	return o
}

func (o EdifactMessageFilterOutput) ToEdifactMessageFilterPtrOutput() EdifactMessageFilterPtrOutput {
	return o.ToEdifactMessageFilterPtrOutputWithContext(context.Background())
}

func (o EdifactMessageFilterOutput) ToEdifactMessageFilterPtrOutputWithContext(ctx context.Context) EdifactMessageFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactMessageFilter) *EdifactMessageFilter {
		return &v
	}).(EdifactMessageFilterPtrOutput)
}

func (o EdifactMessageFilterOutput) MessageFilterType() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactMessageFilter) string { return v.MessageFilterType }).(pulumi.StringOutput)
}

type EdifactMessageFilterPtrOutput struct{ *pulumi.OutputState }

func (EdifactMessageFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactMessageFilter)(nil)).Elem()
}

func (o EdifactMessageFilterPtrOutput) ToEdifactMessageFilterPtrOutput() EdifactMessageFilterPtrOutput {
	return o
}

func (o EdifactMessageFilterPtrOutput) ToEdifactMessageFilterPtrOutputWithContext(ctx context.Context) EdifactMessageFilterPtrOutput {
	return o
}

func (o EdifactMessageFilterPtrOutput) Elem() EdifactMessageFilterOutput {
	return o.ApplyT(func(v *EdifactMessageFilter) EdifactMessageFilter {
		if v != nil {
			return *v
		}
		var ret EdifactMessageFilter
		return ret
	}).(EdifactMessageFilterOutput)
}

func (o EdifactMessageFilterPtrOutput) MessageFilterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactMessageFilter) *string {
		if v == nil {
			return nil
		}
		return &v.MessageFilterType
	}).(pulumi.StringPtrOutput)
}

type EdifactMessageFilterResponse struct {
	MessageFilterType string `pulumi:"messageFilterType"`
}

type EdifactMessageFilterResponseOutput struct{ *pulumi.OutputState }

func (EdifactMessageFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactMessageFilterResponse)(nil)).Elem()
}

func (o EdifactMessageFilterResponseOutput) ToEdifactMessageFilterResponseOutput() EdifactMessageFilterResponseOutput {
	return o
}

func (o EdifactMessageFilterResponseOutput) ToEdifactMessageFilterResponseOutputWithContext(ctx context.Context) EdifactMessageFilterResponseOutput {
	return o
}

func (o EdifactMessageFilterResponseOutput) MessageFilterType() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactMessageFilterResponse) string { return v.MessageFilterType }).(pulumi.StringOutput)
}

type EdifactMessageFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (EdifactMessageFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactMessageFilterResponse)(nil)).Elem()
}

func (o EdifactMessageFilterResponsePtrOutput) ToEdifactMessageFilterResponsePtrOutput() EdifactMessageFilterResponsePtrOutput {
	return o
}

func (o EdifactMessageFilterResponsePtrOutput) ToEdifactMessageFilterResponsePtrOutputWithContext(ctx context.Context) EdifactMessageFilterResponsePtrOutput {
	return o
}

func (o EdifactMessageFilterResponsePtrOutput) Elem() EdifactMessageFilterResponseOutput {
	return o.ApplyT(func(v *EdifactMessageFilterResponse) EdifactMessageFilterResponse {
		if v != nil {
			return *v
		}
		var ret EdifactMessageFilterResponse
		return ret
	}).(EdifactMessageFilterResponseOutput)
}

func (o EdifactMessageFilterResponsePtrOutput) MessageFilterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactMessageFilterResponse) *string {
		if v == nil {
			return nil
		}
		return &v.MessageFilterType
	}).(pulumi.StringPtrOutput)
}

type EdifactMessageIdentifier struct {
	MessageId string `pulumi:"messageId"`
}





type EdifactMessageIdentifierInput interface {
	pulumi.Input

	ToEdifactMessageIdentifierOutput() EdifactMessageIdentifierOutput
	ToEdifactMessageIdentifierOutputWithContext(context.Context) EdifactMessageIdentifierOutput
}

type EdifactMessageIdentifierArgs struct {
	MessageId pulumi.StringInput `pulumi:"messageId"`
}

func (EdifactMessageIdentifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactMessageIdentifier)(nil)).Elem()
}

func (i EdifactMessageIdentifierArgs) ToEdifactMessageIdentifierOutput() EdifactMessageIdentifierOutput {
	return i.ToEdifactMessageIdentifierOutputWithContext(context.Background())
}

func (i EdifactMessageIdentifierArgs) ToEdifactMessageIdentifierOutputWithContext(ctx context.Context) EdifactMessageIdentifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactMessageIdentifierOutput)
}





type EdifactMessageIdentifierArrayInput interface {
	pulumi.Input

	ToEdifactMessageIdentifierArrayOutput() EdifactMessageIdentifierArrayOutput
	ToEdifactMessageIdentifierArrayOutputWithContext(context.Context) EdifactMessageIdentifierArrayOutput
}

type EdifactMessageIdentifierArray []EdifactMessageIdentifierInput

func (EdifactMessageIdentifierArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactMessageIdentifier)(nil)).Elem()
}

func (i EdifactMessageIdentifierArray) ToEdifactMessageIdentifierArrayOutput() EdifactMessageIdentifierArrayOutput {
	return i.ToEdifactMessageIdentifierArrayOutputWithContext(context.Background())
}

func (i EdifactMessageIdentifierArray) ToEdifactMessageIdentifierArrayOutputWithContext(ctx context.Context) EdifactMessageIdentifierArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactMessageIdentifierArrayOutput)
}

type EdifactMessageIdentifierOutput struct{ *pulumi.OutputState }

func (EdifactMessageIdentifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactMessageIdentifier)(nil)).Elem()
}

func (o EdifactMessageIdentifierOutput) ToEdifactMessageIdentifierOutput() EdifactMessageIdentifierOutput {
	return o
}

func (o EdifactMessageIdentifierOutput) ToEdifactMessageIdentifierOutputWithContext(ctx context.Context) EdifactMessageIdentifierOutput {
	return o
}

func (o EdifactMessageIdentifierOutput) MessageId() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactMessageIdentifier) string { return v.MessageId }).(pulumi.StringOutput)
}

type EdifactMessageIdentifierArrayOutput struct{ *pulumi.OutputState }

func (EdifactMessageIdentifierArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactMessageIdentifier)(nil)).Elem()
}

func (o EdifactMessageIdentifierArrayOutput) ToEdifactMessageIdentifierArrayOutput() EdifactMessageIdentifierArrayOutput {
	return o
}

func (o EdifactMessageIdentifierArrayOutput) ToEdifactMessageIdentifierArrayOutputWithContext(ctx context.Context) EdifactMessageIdentifierArrayOutput {
	return o
}

func (o EdifactMessageIdentifierArrayOutput) Index(i pulumi.IntInput) EdifactMessageIdentifierOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactMessageIdentifier {
		return vs[0].([]EdifactMessageIdentifier)[vs[1].(int)]
	}).(EdifactMessageIdentifierOutput)
}

type EdifactMessageIdentifierResponse struct {
	MessageId string `pulumi:"messageId"`
}

type EdifactMessageIdentifierResponseOutput struct{ *pulumi.OutputState }

func (EdifactMessageIdentifierResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactMessageIdentifierResponse)(nil)).Elem()
}

func (o EdifactMessageIdentifierResponseOutput) ToEdifactMessageIdentifierResponseOutput() EdifactMessageIdentifierResponseOutput {
	return o
}

func (o EdifactMessageIdentifierResponseOutput) ToEdifactMessageIdentifierResponseOutputWithContext(ctx context.Context) EdifactMessageIdentifierResponseOutput {
	return o
}

func (o EdifactMessageIdentifierResponseOutput) MessageId() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactMessageIdentifierResponse) string { return v.MessageId }).(pulumi.StringOutput)
}

type EdifactMessageIdentifierResponseArrayOutput struct{ *pulumi.OutputState }

func (EdifactMessageIdentifierResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactMessageIdentifierResponse)(nil)).Elem()
}

func (o EdifactMessageIdentifierResponseArrayOutput) ToEdifactMessageIdentifierResponseArrayOutput() EdifactMessageIdentifierResponseArrayOutput {
	return o
}

func (o EdifactMessageIdentifierResponseArrayOutput) ToEdifactMessageIdentifierResponseArrayOutputWithContext(ctx context.Context) EdifactMessageIdentifierResponseArrayOutput {
	return o
}

func (o EdifactMessageIdentifierResponseArrayOutput) Index(i pulumi.IntInput) EdifactMessageIdentifierResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactMessageIdentifierResponse {
		return vs[0].([]EdifactMessageIdentifierResponse)[vs[1].(int)]
	}).(EdifactMessageIdentifierResponseOutput)
}

type EdifactOneWayAgreement struct {
	ProtocolSettings         EdifactProtocolSettings `pulumi:"protocolSettings"`
	ReceiverBusinessIdentity BusinessIdentity        `pulumi:"receiverBusinessIdentity"`
	SenderBusinessIdentity   BusinessIdentity        `pulumi:"senderBusinessIdentity"`
}





type EdifactOneWayAgreementInput interface {
	pulumi.Input

	ToEdifactOneWayAgreementOutput() EdifactOneWayAgreementOutput
	ToEdifactOneWayAgreementOutputWithContext(context.Context) EdifactOneWayAgreementOutput
}

type EdifactOneWayAgreementArgs struct {
	ProtocolSettings         EdifactProtocolSettingsInput `pulumi:"protocolSettings"`
	ReceiverBusinessIdentity BusinessIdentityInput        `pulumi:"receiverBusinessIdentity"`
	SenderBusinessIdentity   BusinessIdentityInput        `pulumi:"senderBusinessIdentity"`
}

func (EdifactOneWayAgreementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactOneWayAgreement)(nil)).Elem()
}

func (i EdifactOneWayAgreementArgs) ToEdifactOneWayAgreementOutput() EdifactOneWayAgreementOutput {
	return i.ToEdifactOneWayAgreementOutputWithContext(context.Background())
}

func (i EdifactOneWayAgreementArgs) ToEdifactOneWayAgreementOutputWithContext(ctx context.Context) EdifactOneWayAgreementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactOneWayAgreementOutput)
}

func (i EdifactOneWayAgreementArgs) ToEdifactOneWayAgreementPtrOutput() EdifactOneWayAgreementPtrOutput {
	return i.ToEdifactOneWayAgreementPtrOutputWithContext(context.Background())
}

func (i EdifactOneWayAgreementArgs) ToEdifactOneWayAgreementPtrOutputWithContext(ctx context.Context) EdifactOneWayAgreementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactOneWayAgreementOutput).ToEdifactOneWayAgreementPtrOutputWithContext(ctx)
}









type EdifactOneWayAgreementPtrInput interface {
	pulumi.Input

	ToEdifactOneWayAgreementPtrOutput() EdifactOneWayAgreementPtrOutput
	ToEdifactOneWayAgreementPtrOutputWithContext(context.Context) EdifactOneWayAgreementPtrOutput
}

type edifactOneWayAgreementPtrType EdifactOneWayAgreementArgs

func EdifactOneWayAgreementPtr(v *EdifactOneWayAgreementArgs) EdifactOneWayAgreementPtrInput {
	return (*edifactOneWayAgreementPtrType)(v)
}

func (*edifactOneWayAgreementPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactOneWayAgreement)(nil)).Elem()
}

func (i *edifactOneWayAgreementPtrType) ToEdifactOneWayAgreementPtrOutput() EdifactOneWayAgreementPtrOutput {
	return i.ToEdifactOneWayAgreementPtrOutputWithContext(context.Background())
}

func (i *edifactOneWayAgreementPtrType) ToEdifactOneWayAgreementPtrOutputWithContext(ctx context.Context) EdifactOneWayAgreementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactOneWayAgreementPtrOutput)
}

type EdifactOneWayAgreementOutput struct{ *pulumi.OutputState }

func (EdifactOneWayAgreementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactOneWayAgreement)(nil)).Elem()
}

func (o EdifactOneWayAgreementOutput) ToEdifactOneWayAgreementOutput() EdifactOneWayAgreementOutput {
	return o
}

func (o EdifactOneWayAgreementOutput) ToEdifactOneWayAgreementOutputWithContext(ctx context.Context) EdifactOneWayAgreementOutput {
	return o
}

func (o EdifactOneWayAgreementOutput) ToEdifactOneWayAgreementPtrOutput() EdifactOneWayAgreementPtrOutput {
	return o.ToEdifactOneWayAgreementPtrOutputWithContext(context.Background())
}

func (o EdifactOneWayAgreementOutput) ToEdifactOneWayAgreementPtrOutputWithContext(ctx context.Context) EdifactOneWayAgreementPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactOneWayAgreement) *EdifactOneWayAgreement {
		return &v
	}).(EdifactOneWayAgreementPtrOutput)
}

func (o EdifactOneWayAgreementOutput) ProtocolSettings() EdifactProtocolSettingsOutput {
	return o.ApplyT(func(v EdifactOneWayAgreement) EdifactProtocolSettings { return v.ProtocolSettings }).(EdifactProtocolSettingsOutput)
}

func (o EdifactOneWayAgreementOutput) ReceiverBusinessIdentity() BusinessIdentityOutput {
	return o.ApplyT(func(v EdifactOneWayAgreement) BusinessIdentity { return v.ReceiverBusinessIdentity }).(BusinessIdentityOutput)
}

func (o EdifactOneWayAgreementOutput) SenderBusinessIdentity() BusinessIdentityOutput {
	return o.ApplyT(func(v EdifactOneWayAgreement) BusinessIdentity { return v.SenderBusinessIdentity }).(BusinessIdentityOutput)
}

type EdifactOneWayAgreementPtrOutput struct{ *pulumi.OutputState }

func (EdifactOneWayAgreementPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactOneWayAgreement)(nil)).Elem()
}

func (o EdifactOneWayAgreementPtrOutput) ToEdifactOneWayAgreementPtrOutput() EdifactOneWayAgreementPtrOutput {
	return o
}

func (o EdifactOneWayAgreementPtrOutput) ToEdifactOneWayAgreementPtrOutputWithContext(ctx context.Context) EdifactOneWayAgreementPtrOutput {
	return o
}

func (o EdifactOneWayAgreementPtrOutput) Elem() EdifactOneWayAgreementOutput {
	return o.ApplyT(func(v *EdifactOneWayAgreement) EdifactOneWayAgreement {
		if v != nil {
			return *v
		}
		var ret EdifactOneWayAgreement
		return ret
	}).(EdifactOneWayAgreementOutput)
}

func (o EdifactOneWayAgreementPtrOutput) ProtocolSettings() EdifactProtocolSettingsPtrOutput {
	return o.ApplyT(func(v *EdifactOneWayAgreement) *EdifactProtocolSettings {
		if v == nil {
			return nil
		}
		return &v.ProtocolSettings
	}).(EdifactProtocolSettingsPtrOutput)
}

func (o EdifactOneWayAgreementPtrOutput) ReceiverBusinessIdentity() BusinessIdentityPtrOutput {
	return o.ApplyT(func(v *EdifactOneWayAgreement) *BusinessIdentity {
		if v == nil {
			return nil
		}
		return &v.ReceiverBusinessIdentity
	}).(BusinessIdentityPtrOutput)
}

func (o EdifactOneWayAgreementPtrOutput) SenderBusinessIdentity() BusinessIdentityPtrOutput {
	return o.ApplyT(func(v *EdifactOneWayAgreement) *BusinessIdentity {
		if v == nil {
			return nil
		}
		return &v.SenderBusinessIdentity
	}).(BusinessIdentityPtrOutput)
}

type EdifactOneWayAgreementResponse struct {
	ProtocolSettings         EdifactProtocolSettingsResponse `pulumi:"protocolSettings"`
	ReceiverBusinessIdentity BusinessIdentityResponse        `pulumi:"receiverBusinessIdentity"`
	SenderBusinessIdentity   BusinessIdentityResponse        `pulumi:"senderBusinessIdentity"`
}

type EdifactOneWayAgreementResponseOutput struct{ *pulumi.OutputState }

func (EdifactOneWayAgreementResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactOneWayAgreementResponse)(nil)).Elem()
}

func (o EdifactOneWayAgreementResponseOutput) ToEdifactOneWayAgreementResponseOutput() EdifactOneWayAgreementResponseOutput {
	return o
}

func (o EdifactOneWayAgreementResponseOutput) ToEdifactOneWayAgreementResponseOutputWithContext(ctx context.Context) EdifactOneWayAgreementResponseOutput {
	return o
}

func (o EdifactOneWayAgreementResponseOutput) ProtocolSettings() EdifactProtocolSettingsResponseOutput {
	return o.ApplyT(func(v EdifactOneWayAgreementResponse) EdifactProtocolSettingsResponse { return v.ProtocolSettings }).(EdifactProtocolSettingsResponseOutput)
}

func (o EdifactOneWayAgreementResponseOutput) ReceiverBusinessIdentity() BusinessIdentityResponseOutput {
	return o.ApplyT(func(v EdifactOneWayAgreementResponse) BusinessIdentityResponse { return v.ReceiverBusinessIdentity }).(BusinessIdentityResponseOutput)
}

func (o EdifactOneWayAgreementResponseOutput) SenderBusinessIdentity() BusinessIdentityResponseOutput {
	return o.ApplyT(func(v EdifactOneWayAgreementResponse) BusinessIdentityResponse { return v.SenderBusinessIdentity }).(BusinessIdentityResponseOutput)
}

type EdifactOneWayAgreementResponsePtrOutput struct{ *pulumi.OutputState }

func (EdifactOneWayAgreementResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactOneWayAgreementResponse)(nil)).Elem()
}

func (o EdifactOneWayAgreementResponsePtrOutput) ToEdifactOneWayAgreementResponsePtrOutput() EdifactOneWayAgreementResponsePtrOutput {
	return o
}

func (o EdifactOneWayAgreementResponsePtrOutput) ToEdifactOneWayAgreementResponsePtrOutputWithContext(ctx context.Context) EdifactOneWayAgreementResponsePtrOutput {
	return o
}

func (o EdifactOneWayAgreementResponsePtrOutput) Elem() EdifactOneWayAgreementResponseOutput {
	return o.ApplyT(func(v *EdifactOneWayAgreementResponse) EdifactOneWayAgreementResponse {
		if v != nil {
			return *v
		}
		var ret EdifactOneWayAgreementResponse
		return ret
	}).(EdifactOneWayAgreementResponseOutput)
}

func (o EdifactOneWayAgreementResponsePtrOutput) ProtocolSettings() EdifactProtocolSettingsResponsePtrOutput {
	return o.ApplyT(func(v *EdifactOneWayAgreementResponse) *EdifactProtocolSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.ProtocolSettings
	}).(EdifactProtocolSettingsResponsePtrOutput)
}

func (o EdifactOneWayAgreementResponsePtrOutput) ReceiverBusinessIdentity() BusinessIdentityResponsePtrOutput {
	return o.ApplyT(func(v *EdifactOneWayAgreementResponse) *BusinessIdentityResponse {
		if v == nil {
			return nil
		}
		return &v.ReceiverBusinessIdentity
	}).(BusinessIdentityResponsePtrOutput)
}

func (o EdifactOneWayAgreementResponsePtrOutput) SenderBusinessIdentity() BusinessIdentityResponsePtrOutput {
	return o.ApplyT(func(v *EdifactOneWayAgreementResponse) *BusinessIdentityResponse {
		if v == nil {
			return nil
		}
		return &v.SenderBusinessIdentity
	}).(BusinessIdentityResponsePtrOutput)
}

type EdifactProcessingSettings struct {
	CreateEmptyXmlTagsForTrailingSeparators bool `pulumi:"createEmptyXmlTagsForTrailingSeparators"`
	MaskSecurityInfo                        bool `pulumi:"maskSecurityInfo"`
	PreserveInterchange                     bool `pulumi:"preserveInterchange"`
	SuspendInterchangeOnError               bool `pulumi:"suspendInterchangeOnError"`
	UseDotAsDecimalSeparator                bool `pulumi:"useDotAsDecimalSeparator"`
}





type EdifactProcessingSettingsInput interface {
	pulumi.Input

	ToEdifactProcessingSettingsOutput() EdifactProcessingSettingsOutput
	ToEdifactProcessingSettingsOutputWithContext(context.Context) EdifactProcessingSettingsOutput
}

type EdifactProcessingSettingsArgs struct {
	CreateEmptyXmlTagsForTrailingSeparators pulumi.BoolInput `pulumi:"createEmptyXmlTagsForTrailingSeparators"`
	MaskSecurityInfo                        pulumi.BoolInput `pulumi:"maskSecurityInfo"`
	PreserveInterchange                     pulumi.BoolInput `pulumi:"preserveInterchange"`
	SuspendInterchangeOnError               pulumi.BoolInput `pulumi:"suspendInterchangeOnError"`
	UseDotAsDecimalSeparator                pulumi.BoolInput `pulumi:"useDotAsDecimalSeparator"`
}

func (EdifactProcessingSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactProcessingSettings)(nil)).Elem()
}

func (i EdifactProcessingSettingsArgs) ToEdifactProcessingSettingsOutput() EdifactProcessingSettingsOutput {
	return i.ToEdifactProcessingSettingsOutputWithContext(context.Background())
}

func (i EdifactProcessingSettingsArgs) ToEdifactProcessingSettingsOutputWithContext(ctx context.Context) EdifactProcessingSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactProcessingSettingsOutput)
}

func (i EdifactProcessingSettingsArgs) ToEdifactProcessingSettingsPtrOutput() EdifactProcessingSettingsPtrOutput {
	return i.ToEdifactProcessingSettingsPtrOutputWithContext(context.Background())
}

func (i EdifactProcessingSettingsArgs) ToEdifactProcessingSettingsPtrOutputWithContext(ctx context.Context) EdifactProcessingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactProcessingSettingsOutput).ToEdifactProcessingSettingsPtrOutputWithContext(ctx)
}









type EdifactProcessingSettingsPtrInput interface {
	pulumi.Input

	ToEdifactProcessingSettingsPtrOutput() EdifactProcessingSettingsPtrOutput
	ToEdifactProcessingSettingsPtrOutputWithContext(context.Context) EdifactProcessingSettingsPtrOutput
}

type edifactProcessingSettingsPtrType EdifactProcessingSettingsArgs

func EdifactProcessingSettingsPtr(v *EdifactProcessingSettingsArgs) EdifactProcessingSettingsPtrInput {
	return (*edifactProcessingSettingsPtrType)(v)
}

func (*edifactProcessingSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactProcessingSettings)(nil)).Elem()
}

func (i *edifactProcessingSettingsPtrType) ToEdifactProcessingSettingsPtrOutput() EdifactProcessingSettingsPtrOutput {
	return i.ToEdifactProcessingSettingsPtrOutputWithContext(context.Background())
}

func (i *edifactProcessingSettingsPtrType) ToEdifactProcessingSettingsPtrOutputWithContext(ctx context.Context) EdifactProcessingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactProcessingSettingsPtrOutput)
}

type EdifactProcessingSettingsOutput struct{ *pulumi.OutputState }

func (EdifactProcessingSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactProcessingSettings)(nil)).Elem()
}

func (o EdifactProcessingSettingsOutput) ToEdifactProcessingSettingsOutput() EdifactProcessingSettingsOutput {
	return o
}

func (o EdifactProcessingSettingsOutput) ToEdifactProcessingSettingsOutputWithContext(ctx context.Context) EdifactProcessingSettingsOutput {
	return o
}

func (o EdifactProcessingSettingsOutput) ToEdifactProcessingSettingsPtrOutput() EdifactProcessingSettingsPtrOutput {
	return o.ToEdifactProcessingSettingsPtrOutputWithContext(context.Background())
}

func (o EdifactProcessingSettingsOutput) ToEdifactProcessingSettingsPtrOutputWithContext(ctx context.Context) EdifactProcessingSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactProcessingSettings) *EdifactProcessingSettings {
		return &v
	}).(EdifactProcessingSettingsPtrOutput)
}

func (o EdifactProcessingSettingsOutput) CreateEmptyXmlTagsForTrailingSeparators() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactProcessingSettings) bool { return v.CreateEmptyXmlTagsForTrailingSeparators }).(pulumi.BoolOutput)
}

func (o EdifactProcessingSettingsOutput) MaskSecurityInfo() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactProcessingSettings) bool { return v.MaskSecurityInfo }).(pulumi.BoolOutput)
}

func (o EdifactProcessingSettingsOutput) PreserveInterchange() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactProcessingSettings) bool { return v.PreserveInterchange }).(pulumi.BoolOutput)
}

func (o EdifactProcessingSettingsOutput) SuspendInterchangeOnError() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactProcessingSettings) bool { return v.SuspendInterchangeOnError }).(pulumi.BoolOutput)
}

func (o EdifactProcessingSettingsOutput) UseDotAsDecimalSeparator() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactProcessingSettings) bool { return v.UseDotAsDecimalSeparator }).(pulumi.BoolOutput)
}

type EdifactProcessingSettingsPtrOutput struct{ *pulumi.OutputState }

func (EdifactProcessingSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactProcessingSettings)(nil)).Elem()
}

func (o EdifactProcessingSettingsPtrOutput) ToEdifactProcessingSettingsPtrOutput() EdifactProcessingSettingsPtrOutput {
	return o
}

func (o EdifactProcessingSettingsPtrOutput) ToEdifactProcessingSettingsPtrOutputWithContext(ctx context.Context) EdifactProcessingSettingsPtrOutput {
	return o
}

func (o EdifactProcessingSettingsPtrOutput) Elem() EdifactProcessingSettingsOutput {
	return o.ApplyT(func(v *EdifactProcessingSettings) EdifactProcessingSettings {
		if v != nil {
			return *v
		}
		var ret EdifactProcessingSettings
		return ret
	}).(EdifactProcessingSettingsOutput)
}

func (o EdifactProcessingSettingsPtrOutput) CreateEmptyXmlTagsForTrailingSeparators() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.CreateEmptyXmlTagsForTrailingSeparators
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsPtrOutput) MaskSecurityInfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.MaskSecurityInfo
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsPtrOutput) PreserveInterchange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.PreserveInterchange
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsPtrOutput) SuspendInterchangeOnError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.SuspendInterchangeOnError
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsPtrOutput) UseDotAsDecimalSeparator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.UseDotAsDecimalSeparator
	}).(pulumi.BoolPtrOutput)
}

type EdifactProcessingSettingsResponse struct {
	CreateEmptyXmlTagsForTrailingSeparators bool `pulumi:"createEmptyXmlTagsForTrailingSeparators"`
	MaskSecurityInfo                        bool `pulumi:"maskSecurityInfo"`
	PreserveInterchange                     bool `pulumi:"preserveInterchange"`
	SuspendInterchangeOnError               bool `pulumi:"suspendInterchangeOnError"`
	UseDotAsDecimalSeparator                bool `pulumi:"useDotAsDecimalSeparator"`
}

type EdifactProcessingSettingsResponseOutput struct{ *pulumi.OutputState }

func (EdifactProcessingSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactProcessingSettingsResponse)(nil)).Elem()
}

func (o EdifactProcessingSettingsResponseOutput) ToEdifactProcessingSettingsResponseOutput() EdifactProcessingSettingsResponseOutput {
	return o
}

func (o EdifactProcessingSettingsResponseOutput) ToEdifactProcessingSettingsResponseOutputWithContext(ctx context.Context) EdifactProcessingSettingsResponseOutput {
	return o
}

func (o EdifactProcessingSettingsResponseOutput) CreateEmptyXmlTagsForTrailingSeparators() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactProcessingSettingsResponse) bool { return v.CreateEmptyXmlTagsForTrailingSeparators }).(pulumi.BoolOutput)
}

func (o EdifactProcessingSettingsResponseOutput) MaskSecurityInfo() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactProcessingSettingsResponse) bool { return v.MaskSecurityInfo }).(pulumi.BoolOutput)
}

func (o EdifactProcessingSettingsResponseOutput) PreserveInterchange() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactProcessingSettingsResponse) bool { return v.PreserveInterchange }).(pulumi.BoolOutput)
}

func (o EdifactProcessingSettingsResponseOutput) SuspendInterchangeOnError() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactProcessingSettingsResponse) bool { return v.SuspendInterchangeOnError }).(pulumi.BoolOutput)
}

func (o EdifactProcessingSettingsResponseOutput) UseDotAsDecimalSeparator() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactProcessingSettingsResponse) bool { return v.UseDotAsDecimalSeparator }).(pulumi.BoolOutput)
}

type EdifactProcessingSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (EdifactProcessingSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactProcessingSettingsResponse)(nil)).Elem()
}

func (o EdifactProcessingSettingsResponsePtrOutput) ToEdifactProcessingSettingsResponsePtrOutput() EdifactProcessingSettingsResponsePtrOutput {
	return o
}

func (o EdifactProcessingSettingsResponsePtrOutput) ToEdifactProcessingSettingsResponsePtrOutputWithContext(ctx context.Context) EdifactProcessingSettingsResponsePtrOutput {
	return o
}

func (o EdifactProcessingSettingsResponsePtrOutput) Elem() EdifactProcessingSettingsResponseOutput {
	return o.ApplyT(func(v *EdifactProcessingSettingsResponse) EdifactProcessingSettingsResponse {
		if v != nil {
			return *v
		}
		var ret EdifactProcessingSettingsResponse
		return ret
	}).(EdifactProcessingSettingsResponseOutput)
}

func (o EdifactProcessingSettingsResponsePtrOutput) CreateEmptyXmlTagsForTrailingSeparators() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CreateEmptyXmlTagsForTrailingSeparators
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsResponsePtrOutput) MaskSecurityInfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.MaskSecurityInfo
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsResponsePtrOutput) PreserveInterchange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.PreserveInterchange
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsResponsePtrOutput) SuspendInterchangeOnError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.SuspendInterchangeOnError
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactProcessingSettingsResponsePtrOutput) UseDotAsDecimalSeparator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.UseDotAsDecimalSeparator
	}).(pulumi.BoolPtrOutput)
}

type EdifactProtocolSettings struct {
	AcknowledgementSettings   EdifactAcknowledgementSettings `pulumi:"acknowledgementSettings"`
	EdifactDelimiterOverrides []EdifactDelimiterOverride     `pulumi:"edifactDelimiterOverrides"`
	EnvelopeOverrides         []EdifactEnvelopeOverride      `pulumi:"envelopeOverrides"`
	EnvelopeSettings          EdifactEnvelopeSettings        `pulumi:"envelopeSettings"`
	FramingSettings           EdifactFramingSettings         `pulumi:"framingSettings"`
	MessageFilter             EdifactMessageFilter           `pulumi:"messageFilter"`
	MessageFilterList         []EdifactMessageIdentifier     `pulumi:"messageFilterList"`
	ProcessingSettings        EdifactProcessingSettings      `pulumi:"processingSettings"`
	SchemaReferences          []EdifactSchemaReference       `pulumi:"schemaReferences"`
	ValidationOverrides       []EdifactValidationOverride    `pulumi:"validationOverrides"`
	ValidationSettings        EdifactValidationSettings      `pulumi:"validationSettings"`
}





type EdifactProtocolSettingsInput interface {
	pulumi.Input

	ToEdifactProtocolSettingsOutput() EdifactProtocolSettingsOutput
	ToEdifactProtocolSettingsOutputWithContext(context.Context) EdifactProtocolSettingsOutput
}

type EdifactProtocolSettingsArgs struct {
	AcknowledgementSettings   EdifactAcknowledgementSettingsInput `pulumi:"acknowledgementSettings"`
	EdifactDelimiterOverrides EdifactDelimiterOverrideArrayInput  `pulumi:"edifactDelimiterOverrides"`
	EnvelopeOverrides         EdifactEnvelopeOverrideArrayInput   `pulumi:"envelopeOverrides"`
	EnvelopeSettings          EdifactEnvelopeSettingsInput        `pulumi:"envelopeSettings"`
	FramingSettings           EdifactFramingSettingsInput         `pulumi:"framingSettings"`
	MessageFilter             EdifactMessageFilterInput           `pulumi:"messageFilter"`
	MessageFilterList         EdifactMessageIdentifierArrayInput  `pulumi:"messageFilterList"`
	ProcessingSettings        EdifactProcessingSettingsInput      `pulumi:"processingSettings"`
	SchemaReferences          EdifactSchemaReferenceArrayInput    `pulumi:"schemaReferences"`
	ValidationOverrides       EdifactValidationOverrideArrayInput `pulumi:"validationOverrides"`
	ValidationSettings        EdifactValidationSettingsInput      `pulumi:"validationSettings"`
}

func (EdifactProtocolSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactProtocolSettings)(nil)).Elem()
}

func (i EdifactProtocolSettingsArgs) ToEdifactProtocolSettingsOutput() EdifactProtocolSettingsOutput {
	return i.ToEdifactProtocolSettingsOutputWithContext(context.Background())
}

func (i EdifactProtocolSettingsArgs) ToEdifactProtocolSettingsOutputWithContext(ctx context.Context) EdifactProtocolSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactProtocolSettingsOutput)
}

func (i EdifactProtocolSettingsArgs) ToEdifactProtocolSettingsPtrOutput() EdifactProtocolSettingsPtrOutput {
	return i.ToEdifactProtocolSettingsPtrOutputWithContext(context.Background())
}

func (i EdifactProtocolSettingsArgs) ToEdifactProtocolSettingsPtrOutputWithContext(ctx context.Context) EdifactProtocolSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactProtocolSettingsOutput).ToEdifactProtocolSettingsPtrOutputWithContext(ctx)
}









type EdifactProtocolSettingsPtrInput interface {
	pulumi.Input

	ToEdifactProtocolSettingsPtrOutput() EdifactProtocolSettingsPtrOutput
	ToEdifactProtocolSettingsPtrOutputWithContext(context.Context) EdifactProtocolSettingsPtrOutput
}

type edifactProtocolSettingsPtrType EdifactProtocolSettingsArgs

func EdifactProtocolSettingsPtr(v *EdifactProtocolSettingsArgs) EdifactProtocolSettingsPtrInput {
	return (*edifactProtocolSettingsPtrType)(v)
}

func (*edifactProtocolSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactProtocolSettings)(nil)).Elem()
}

func (i *edifactProtocolSettingsPtrType) ToEdifactProtocolSettingsPtrOutput() EdifactProtocolSettingsPtrOutput {
	return i.ToEdifactProtocolSettingsPtrOutputWithContext(context.Background())
}

func (i *edifactProtocolSettingsPtrType) ToEdifactProtocolSettingsPtrOutputWithContext(ctx context.Context) EdifactProtocolSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactProtocolSettingsPtrOutput)
}

type EdifactProtocolSettingsOutput struct{ *pulumi.OutputState }

func (EdifactProtocolSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactProtocolSettings)(nil)).Elem()
}

func (o EdifactProtocolSettingsOutput) ToEdifactProtocolSettingsOutput() EdifactProtocolSettingsOutput {
	return o
}

func (o EdifactProtocolSettingsOutput) ToEdifactProtocolSettingsOutputWithContext(ctx context.Context) EdifactProtocolSettingsOutput {
	return o
}

func (o EdifactProtocolSettingsOutput) ToEdifactProtocolSettingsPtrOutput() EdifactProtocolSettingsPtrOutput {
	return o.ToEdifactProtocolSettingsPtrOutputWithContext(context.Background())
}

func (o EdifactProtocolSettingsOutput) ToEdifactProtocolSettingsPtrOutputWithContext(ctx context.Context) EdifactProtocolSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactProtocolSettings) *EdifactProtocolSettings {
		return &v
	}).(EdifactProtocolSettingsPtrOutput)
}

func (o EdifactProtocolSettingsOutput) AcknowledgementSettings() EdifactAcknowledgementSettingsOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) EdifactAcknowledgementSettings { return v.AcknowledgementSettings }).(EdifactAcknowledgementSettingsOutput)
}

func (o EdifactProtocolSettingsOutput) EdifactDelimiterOverrides() EdifactDelimiterOverrideArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) []EdifactDelimiterOverride { return v.EdifactDelimiterOverrides }).(EdifactDelimiterOverrideArrayOutput)
}

func (o EdifactProtocolSettingsOutput) EnvelopeOverrides() EdifactEnvelopeOverrideArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) []EdifactEnvelopeOverride { return v.EnvelopeOverrides }).(EdifactEnvelopeOverrideArrayOutput)
}

func (o EdifactProtocolSettingsOutput) EnvelopeSettings() EdifactEnvelopeSettingsOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) EdifactEnvelopeSettings { return v.EnvelopeSettings }).(EdifactEnvelopeSettingsOutput)
}

func (o EdifactProtocolSettingsOutput) FramingSettings() EdifactFramingSettingsOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) EdifactFramingSettings { return v.FramingSettings }).(EdifactFramingSettingsOutput)
}

func (o EdifactProtocolSettingsOutput) MessageFilter() EdifactMessageFilterOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) EdifactMessageFilter { return v.MessageFilter }).(EdifactMessageFilterOutput)
}

func (o EdifactProtocolSettingsOutput) MessageFilterList() EdifactMessageIdentifierArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) []EdifactMessageIdentifier { return v.MessageFilterList }).(EdifactMessageIdentifierArrayOutput)
}

func (o EdifactProtocolSettingsOutput) ProcessingSettings() EdifactProcessingSettingsOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) EdifactProcessingSettings { return v.ProcessingSettings }).(EdifactProcessingSettingsOutput)
}

func (o EdifactProtocolSettingsOutput) SchemaReferences() EdifactSchemaReferenceArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) []EdifactSchemaReference { return v.SchemaReferences }).(EdifactSchemaReferenceArrayOutput)
}

func (o EdifactProtocolSettingsOutput) ValidationOverrides() EdifactValidationOverrideArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) []EdifactValidationOverride { return v.ValidationOverrides }).(EdifactValidationOverrideArrayOutput)
}

func (o EdifactProtocolSettingsOutput) ValidationSettings() EdifactValidationSettingsOutput {
	return o.ApplyT(func(v EdifactProtocolSettings) EdifactValidationSettings { return v.ValidationSettings }).(EdifactValidationSettingsOutput)
}

type EdifactProtocolSettingsPtrOutput struct{ *pulumi.OutputState }

func (EdifactProtocolSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactProtocolSettings)(nil)).Elem()
}

func (o EdifactProtocolSettingsPtrOutput) ToEdifactProtocolSettingsPtrOutput() EdifactProtocolSettingsPtrOutput {
	return o
}

func (o EdifactProtocolSettingsPtrOutput) ToEdifactProtocolSettingsPtrOutputWithContext(ctx context.Context) EdifactProtocolSettingsPtrOutput {
	return o
}

func (o EdifactProtocolSettingsPtrOutput) Elem() EdifactProtocolSettingsOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) EdifactProtocolSettings {
		if v != nil {
			return *v
		}
		var ret EdifactProtocolSettings
		return ret
	}).(EdifactProtocolSettingsOutput)
}

func (o EdifactProtocolSettingsPtrOutput) AcknowledgementSettings() EdifactAcknowledgementSettingsPtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) *EdifactAcknowledgementSettings {
		if v == nil {
			return nil
		}
		return &v.AcknowledgementSettings
	}).(EdifactAcknowledgementSettingsPtrOutput)
}

func (o EdifactProtocolSettingsPtrOutput) EdifactDelimiterOverrides() EdifactDelimiterOverrideArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) []EdifactDelimiterOverride {
		if v == nil {
			return nil
		}
		return v.EdifactDelimiterOverrides
	}).(EdifactDelimiterOverrideArrayOutput)
}

func (o EdifactProtocolSettingsPtrOutput) EnvelopeOverrides() EdifactEnvelopeOverrideArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) []EdifactEnvelopeOverride {
		if v == nil {
			return nil
		}
		return v.EnvelopeOverrides
	}).(EdifactEnvelopeOverrideArrayOutput)
}

func (o EdifactProtocolSettingsPtrOutput) EnvelopeSettings() EdifactEnvelopeSettingsPtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) *EdifactEnvelopeSettings {
		if v == nil {
			return nil
		}
		return &v.EnvelopeSettings
	}).(EdifactEnvelopeSettingsPtrOutput)
}

func (o EdifactProtocolSettingsPtrOutput) FramingSettings() EdifactFramingSettingsPtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) *EdifactFramingSettings {
		if v == nil {
			return nil
		}
		return &v.FramingSettings
	}).(EdifactFramingSettingsPtrOutput)
}

func (o EdifactProtocolSettingsPtrOutput) MessageFilter() EdifactMessageFilterPtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) *EdifactMessageFilter {
		if v == nil {
			return nil
		}
		return &v.MessageFilter
	}).(EdifactMessageFilterPtrOutput)
}

func (o EdifactProtocolSettingsPtrOutput) MessageFilterList() EdifactMessageIdentifierArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) []EdifactMessageIdentifier {
		if v == nil {
			return nil
		}
		return v.MessageFilterList
	}).(EdifactMessageIdentifierArrayOutput)
}

func (o EdifactProtocolSettingsPtrOutput) ProcessingSettings() EdifactProcessingSettingsPtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) *EdifactProcessingSettings {
		if v == nil {
			return nil
		}
		return &v.ProcessingSettings
	}).(EdifactProcessingSettingsPtrOutput)
}

func (o EdifactProtocolSettingsPtrOutput) SchemaReferences() EdifactSchemaReferenceArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) []EdifactSchemaReference {
		if v == nil {
			return nil
		}
		return v.SchemaReferences
	}).(EdifactSchemaReferenceArrayOutput)
}

func (o EdifactProtocolSettingsPtrOutput) ValidationOverrides() EdifactValidationOverrideArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) []EdifactValidationOverride {
		if v == nil {
			return nil
		}
		return v.ValidationOverrides
	}).(EdifactValidationOverrideArrayOutput)
}

func (o EdifactProtocolSettingsPtrOutput) ValidationSettings() EdifactValidationSettingsPtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettings) *EdifactValidationSettings {
		if v == nil {
			return nil
		}
		return &v.ValidationSettings
	}).(EdifactValidationSettingsPtrOutput)
}

type EdifactProtocolSettingsResponse struct {
	AcknowledgementSettings   EdifactAcknowledgementSettingsResponse `pulumi:"acknowledgementSettings"`
	EdifactDelimiterOverrides []EdifactDelimiterOverrideResponse     `pulumi:"edifactDelimiterOverrides"`
	EnvelopeOverrides         []EdifactEnvelopeOverrideResponse      `pulumi:"envelopeOverrides"`
	EnvelopeSettings          EdifactEnvelopeSettingsResponse        `pulumi:"envelopeSettings"`
	FramingSettings           EdifactFramingSettingsResponse         `pulumi:"framingSettings"`
	MessageFilter             EdifactMessageFilterResponse           `pulumi:"messageFilter"`
	MessageFilterList         []EdifactMessageIdentifierResponse     `pulumi:"messageFilterList"`
	ProcessingSettings        EdifactProcessingSettingsResponse      `pulumi:"processingSettings"`
	SchemaReferences          []EdifactSchemaReferenceResponse       `pulumi:"schemaReferences"`
	ValidationOverrides       []EdifactValidationOverrideResponse    `pulumi:"validationOverrides"`
	ValidationSettings        EdifactValidationSettingsResponse      `pulumi:"validationSettings"`
}

type EdifactProtocolSettingsResponseOutput struct{ *pulumi.OutputState }

func (EdifactProtocolSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactProtocolSettingsResponse)(nil)).Elem()
}

func (o EdifactProtocolSettingsResponseOutput) ToEdifactProtocolSettingsResponseOutput() EdifactProtocolSettingsResponseOutput {
	return o
}

func (o EdifactProtocolSettingsResponseOutput) ToEdifactProtocolSettingsResponseOutputWithContext(ctx context.Context) EdifactProtocolSettingsResponseOutput {
	return o
}

func (o EdifactProtocolSettingsResponseOutput) AcknowledgementSettings() EdifactAcknowledgementSettingsResponseOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) EdifactAcknowledgementSettingsResponse {
		return v.AcknowledgementSettings
	}).(EdifactAcknowledgementSettingsResponseOutput)
}

func (o EdifactProtocolSettingsResponseOutput) EdifactDelimiterOverrides() EdifactDelimiterOverrideResponseArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) []EdifactDelimiterOverrideResponse {
		return v.EdifactDelimiterOverrides
	}).(EdifactDelimiterOverrideResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponseOutput) EnvelopeOverrides() EdifactEnvelopeOverrideResponseArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) []EdifactEnvelopeOverrideResponse { return v.EnvelopeOverrides }).(EdifactEnvelopeOverrideResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponseOutput) EnvelopeSettings() EdifactEnvelopeSettingsResponseOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) EdifactEnvelopeSettingsResponse { return v.EnvelopeSettings }).(EdifactEnvelopeSettingsResponseOutput)
}

func (o EdifactProtocolSettingsResponseOutput) FramingSettings() EdifactFramingSettingsResponseOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) EdifactFramingSettingsResponse { return v.FramingSettings }).(EdifactFramingSettingsResponseOutput)
}

func (o EdifactProtocolSettingsResponseOutput) MessageFilter() EdifactMessageFilterResponseOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) EdifactMessageFilterResponse { return v.MessageFilter }).(EdifactMessageFilterResponseOutput)
}

func (o EdifactProtocolSettingsResponseOutput) MessageFilterList() EdifactMessageIdentifierResponseArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) []EdifactMessageIdentifierResponse { return v.MessageFilterList }).(EdifactMessageIdentifierResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponseOutput) ProcessingSettings() EdifactProcessingSettingsResponseOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) EdifactProcessingSettingsResponse { return v.ProcessingSettings }).(EdifactProcessingSettingsResponseOutput)
}

func (o EdifactProtocolSettingsResponseOutput) SchemaReferences() EdifactSchemaReferenceResponseArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) []EdifactSchemaReferenceResponse { return v.SchemaReferences }).(EdifactSchemaReferenceResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponseOutput) ValidationOverrides() EdifactValidationOverrideResponseArrayOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) []EdifactValidationOverrideResponse {
		return v.ValidationOverrides
	}).(EdifactValidationOverrideResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponseOutput) ValidationSettings() EdifactValidationSettingsResponseOutput {
	return o.ApplyT(func(v EdifactProtocolSettingsResponse) EdifactValidationSettingsResponse { return v.ValidationSettings }).(EdifactValidationSettingsResponseOutput)
}

type EdifactProtocolSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (EdifactProtocolSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactProtocolSettingsResponse)(nil)).Elem()
}

func (o EdifactProtocolSettingsResponsePtrOutput) ToEdifactProtocolSettingsResponsePtrOutput() EdifactProtocolSettingsResponsePtrOutput {
	return o
}

func (o EdifactProtocolSettingsResponsePtrOutput) ToEdifactProtocolSettingsResponsePtrOutputWithContext(ctx context.Context) EdifactProtocolSettingsResponsePtrOutput {
	return o
}

func (o EdifactProtocolSettingsResponsePtrOutput) Elem() EdifactProtocolSettingsResponseOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) EdifactProtocolSettingsResponse {
		if v != nil {
			return *v
		}
		var ret EdifactProtocolSettingsResponse
		return ret
	}).(EdifactProtocolSettingsResponseOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) AcknowledgementSettings() EdifactAcknowledgementSettingsResponsePtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) *EdifactAcknowledgementSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.AcknowledgementSettings
	}).(EdifactAcknowledgementSettingsResponsePtrOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) EdifactDelimiterOverrides() EdifactDelimiterOverrideResponseArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) []EdifactDelimiterOverrideResponse {
		if v == nil {
			return nil
		}
		return v.EdifactDelimiterOverrides
	}).(EdifactDelimiterOverrideResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) EnvelopeOverrides() EdifactEnvelopeOverrideResponseArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) []EdifactEnvelopeOverrideResponse {
		if v == nil {
			return nil
		}
		return v.EnvelopeOverrides
	}).(EdifactEnvelopeOverrideResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) EnvelopeSettings() EdifactEnvelopeSettingsResponsePtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) *EdifactEnvelopeSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.EnvelopeSettings
	}).(EdifactEnvelopeSettingsResponsePtrOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) FramingSettings() EdifactFramingSettingsResponsePtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) *EdifactFramingSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.FramingSettings
	}).(EdifactFramingSettingsResponsePtrOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) MessageFilter() EdifactMessageFilterResponsePtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) *EdifactMessageFilterResponse {
		if v == nil {
			return nil
		}
		return &v.MessageFilter
	}).(EdifactMessageFilterResponsePtrOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) MessageFilterList() EdifactMessageIdentifierResponseArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) []EdifactMessageIdentifierResponse {
		if v == nil {
			return nil
		}
		return v.MessageFilterList
	}).(EdifactMessageIdentifierResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) ProcessingSettings() EdifactProcessingSettingsResponsePtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) *EdifactProcessingSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.ProcessingSettings
	}).(EdifactProcessingSettingsResponsePtrOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) SchemaReferences() EdifactSchemaReferenceResponseArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) []EdifactSchemaReferenceResponse {
		if v == nil {
			return nil
		}
		return v.SchemaReferences
	}).(EdifactSchemaReferenceResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) ValidationOverrides() EdifactValidationOverrideResponseArrayOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) []EdifactValidationOverrideResponse {
		if v == nil {
			return nil
		}
		return v.ValidationOverrides
	}).(EdifactValidationOverrideResponseArrayOutput)
}

func (o EdifactProtocolSettingsResponsePtrOutput) ValidationSettings() EdifactValidationSettingsResponsePtrOutput {
	return o.ApplyT(func(v *EdifactProtocolSettingsResponse) *EdifactValidationSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.ValidationSettings
	}).(EdifactValidationSettingsResponsePtrOutput)
}

type EdifactSchemaReference struct {
	AssociationAssignedCode    *string `pulumi:"associationAssignedCode"`
	MessageId                  string  `pulumi:"messageId"`
	MessageRelease             string  `pulumi:"messageRelease"`
	MessageVersion             string  `pulumi:"messageVersion"`
	SchemaName                 string  `pulumi:"schemaName"`
	SenderApplicationId        *string `pulumi:"senderApplicationId"`
	SenderApplicationQualifier *string `pulumi:"senderApplicationQualifier"`
}





type EdifactSchemaReferenceInput interface {
	pulumi.Input

	ToEdifactSchemaReferenceOutput() EdifactSchemaReferenceOutput
	ToEdifactSchemaReferenceOutputWithContext(context.Context) EdifactSchemaReferenceOutput
}

type EdifactSchemaReferenceArgs struct {
	AssociationAssignedCode    pulumi.StringPtrInput `pulumi:"associationAssignedCode"`
	MessageId                  pulumi.StringInput    `pulumi:"messageId"`
	MessageRelease             pulumi.StringInput    `pulumi:"messageRelease"`
	MessageVersion             pulumi.StringInput    `pulumi:"messageVersion"`
	SchemaName                 pulumi.StringInput    `pulumi:"schemaName"`
	SenderApplicationId        pulumi.StringPtrInput `pulumi:"senderApplicationId"`
	SenderApplicationQualifier pulumi.StringPtrInput `pulumi:"senderApplicationQualifier"`
}

func (EdifactSchemaReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactSchemaReference)(nil)).Elem()
}

func (i EdifactSchemaReferenceArgs) ToEdifactSchemaReferenceOutput() EdifactSchemaReferenceOutput {
	return i.ToEdifactSchemaReferenceOutputWithContext(context.Background())
}

func (i EdifactSchemaReferenceArgs) ToEdifactSchemaReferenceOutputWithContext(ctx context.Context) EdifactSchemaReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactSchemaReferenceOutput)
}





type EdifactSchemaReferenceArrayInput interface {
	pulumi.Input

	ToEdifactSchemaReferenceArrayOutput() EdifactSchemaReferenceArrayOutput
	ToEdifactSchemaReferenceArrayOutputWithContext(context.Context) EdifactSchemaReferenceArrayOutput
}

type EdifactSchemaReferenceArray []EdifactSchemaReferenceInput

func (EdifactSchemaReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactSchemaReference)(nil)).Elem()
}

func (i EdifactSchemaReferenceArray) ToEdifactSchemaReferenceArrayOutput() EdifactSchemaReferenceArrayOutput {
	return i.ToEdifactSchemaReferenceArrayOutputWithContext(context.Background())
}

func (i EdifactSchemaReferenceArray) ToEdifactSchemaReferenceArrayOutputWithContext(ctx context.Context) EdifactSchemaReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactSchemaReferenceArrayOutput)
}

type EdifactSchemaReferenceOutput struct{ *pulumi.OutputState }

func (EdifactSchemaReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactSchemaReference)(nil)).Elem()
}

func (o EdifactSchemaReferenceOutput) ToEdifactSchemaReferenceOutput() EdifactSchemaReferenceOutput {
	return o
}

func (o EdifactSchemaReferenceOutput) ToEdifactSchemaReferenceOutputWithContext(ctx context.Context) EdifactSchemaReferenceOutput {
	return o
}

func (o EdifactSchemaReferenceOutput) AssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReference) *string { return v.AssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactSchemaReferenceOutput) MessageId() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactSchemaReference) string { return v.MessageId }).(pulumi.StringOutput)
}

func (o EdifactSchemaReferenceOutput) MessageRelease() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactSchemaReference) string { return v.MessageRelease }).(pulumi.StringOutput)
}

func (o EdifactSchemaReferenceOutput) MessageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactSchemaReference) string { return v.MessageVersion }).(pulumi.StringOutput)
}

func (o EdifactSchemaReferenceOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactSchemaReference) string { return v.SchemaName }).(pulumi.StringOutput)
}

func (o EdifactSchemaReferenceOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReference) *string { return v.SenderApplicationId }).(pulumi.StringPtrOutput)
}

func (o EdifactSchemaReferenceOutput) SenderApplicationQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReference) *string { return v.SenderApplicationQualifier }).(pulumi.StringPtrOutput)
}

type EdifactSchemaReferenceArrayOutput struct{ *pulumi.OutputState }

func (EdifactSchemaReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactSchemaReference)(nil)).Elem()
}

func (o EdifactSchemaReferenceArrayOutput) ToEdifactSchemaReferenceArrayOutput() EdifactSchemaReferenceArrayOutput {
	return o
}

func (o EdifactSchemaReferenceArrayOutput) ToEdifactSchemaReferenceArrayOutputWithContext(ctx context.Context) EdifactSchemaReferenceArrayOutput {
	return o
}

func (o EdifactSchemaReferenceArrayOutput) Index(i pulumi.IntInput) EdifactSchemaReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactSchemaReference {
		return vs[0].([]EdifactSchemaReference)[vs[1].(int)]
	}).(EdifactSchemaReferenceOutput)
}

type EdifactSchemaReferenceResponse struct {
	AssociationAssignedCode    *string `pulumi:"associationAssignedCode"`
	MessageId                  string  `pulumi:"messageId"`
	MessageRelease             string  `pulumi:"messageRelease"`
	MessageVersion             string  `pulumi:"messageVersion"`
	SchemaName                 string  `pulumi:"schemaName"`
	SenderApplicationId        *string `pulumi:"senderApplicationId"`
	SenderApplicationQualifier *string `pulumi:"senderApplicationQualifier"`
}

type EdifactSchemaReferenceResponseOutput struct{ *pulumi.OutputState }

func (EdifactSchemaReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactSchemaReferenceResponse)(nil)).Elem()
}

func (o EdifactSchemaReferenceResponseOutput) ToEdifactSchemaReferenceResponseOutput() EdifactSchemaReferenceResponseOutput {
	return o
}

func (o EdifactSchemaReferenceResponseOutput) ToEdifactSchemaReferenceResponseOutputWithContext(ctx context.Context) EdifactSchemaReferenceResponseOutput {
	return o
}

func (o EdifactSchemaReferenceResponseOutput) AssociationAssignedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReferenceResponse) *string { return v.AssociationAssignedCode }).(pulumi.StringPtrOutput)
}

func (o EdifactSchemaReferenceResponseOutput) MessageId() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactSchemaReferenceResponse) string { return v.MessageId }).(pulumi.StringOutput)
}

func (o EdifactSchemaReferenceResponseOutput) MessageRelease() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactSchemaReferenceResponse) string { return v.MessageRelease }).(pulumi.StringOutput)
}

func (o EdifactSchemaReferenceResponseOutput) MessageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactSchemaReferenceResponse) string { return v.MessageVersion }).(pulumi.StringOutput)
}

func (o EdifactSchemaReferenceResponseOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactSchemaReferenceResponse) string { return v.SchemaName }).(pulumi.StringOutput)
}

func (o EdifactSchemaReferenceResponseOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReferenceResponse) *string { return v.SenderApplicationId }).(pulumi.StringPtrOutput)
}

func (o EdifactSchemaReferenceResponseOutput) SenderApplicationQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EdifactSchemaReferenceResponse) *string { return v.SenderApplicationQualifier }).(pulumi.StringPtrOutput)
}

type EdifactSchemaReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (EdifactSchemaReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactSchemaReferenceResponse)(nil)).Elem()
}

func (o EdifactSchemaReferenceResponseArrayOutput) ToEdifactSchemaReferenceResponseArrayOutput() EdifactSchemaReferenceResponseArrayOutput {
	return o
}

func (o EdifactSchemaReferenceResponseArrayOutput) ToEdifactSchemaReferenceResponseArrayOutputWithContext(ctx context.Context) EdifactSchemaReferenceResponseArrayOutput {
	return o
}

func (o EdifactSchemaReferenceResponseArrayOutput) Index(i pulumi.IntInput) EdifactSchemaReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactSchemaReferenceResponse {
		return vs[0].([]EdifactSchemaReferenceResponse)[vs[1].(int)]
	}).(EdifactSchemaReferenceResponseOutput)
}

type EdifactValidationOverride struct {
	AllowLeadingAndTrailingSpacesAndZeroes bool   `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	EnforceCharacterSet                    bool   `pulumi:"enforceCharacterSet"`
	MessageId                              string `pulumi:"messageId"`
	TrailingSeparatorPolicy                string `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes  bool   `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateEDITypes                       bool   `pulumi:"validateEDITypes"`
	ValidateXSDTypes                       bool   `pulumi:"validateXSDTypes"`
}





type EdifactValidationOverrideInput interface {
	pulumi.Input

	ToEdifactValidationOverrideOutput() EdifactValidationOverrideOutput
	ToEdifactValidationOverrideOutputWithContext(context.Context) EdifactValidationOverrideOutput
}

type EdifactValidationOverrideArgs struct {
	AllowLeadingAndTrailingSpacesAndZeroes pulumi.BoolInput   `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	EnforceCharacterSet                    pulumi.BoolInput   `pulumi:"enforceCharacterSet"`
	MessageId                              pulumi.StringInput `pulumi:"messageId"`
	TrailingSeparatorPolicy                pulumi.StringInput `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes  pulumi.BoolInput   `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateEDITypes                       pulumi.BoolInput   `pulumi:"validateEDITypes"`
	ValidateXSDTypes                       pulumi.BoolInput   `pulumi:"validateXSDTypes"`
}

func (EdifactValidationOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactValidationOverride)(nil)).Elem()
}

func (i EdifactValidationOverrideArgs) ToEdifactValidationOverrideOutput() EdifactValidationOverrideOutput {
	return i.ToEdifactValidationOverrideOutputWithContext(context.Background())
}

func (i EdifactValidationOverrideArgs) ToEdifactValidationOverrideOutputWithContext(ctx context.Context) EdifactValidationOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactValidationOverrideOutput)
}





type EdifactValidationOverrideArrayInput interface {
	pulumi.Input

	ToEdifactValidationOverrideArrayOutput() EdifactValidationOverrideArrayOutput
	ToEdifactValidationOverrideArrayOutputWithContext(context.Context) EdifactValidationOverrideArrayOutput
}

type EdifactValidationOverrideArray []EdifactValidationOverrideInput

func (EdifactValidationOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactValidationOverride)(nil)).Elem()
}

func (i EdifactValidationOverrideArray) ToEdifactValidationOverrideArrayOutput() EdifactValidationOverrideArrayOutput {
	return i.ToEdifactValidationOverrideArrayOutputWithContext(context.Background())
}

func (i EdifactValidationOverrideArray) ToEdifactValidationOverrideArrayOutputWithContext(ctx context.Context) EdifactValidationOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactValidationOverrideArrayOutput)
}

type EdifactValidationOverrideOutput struct{ *pulumi.OutputState }

func (EdifactValidationOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactValidationOverride)(nil)).Elem()
}

func (o EdifactValidationOverrideOutput) ToEdifactValidationOverrideOutput() EdifactValidationOverrideOutput {
	return o
}

func (o EdifactValidationOverrideOutput) ToEdifactValidationOverrideOutputWithContext(ctx context.Context) EdifactValidationOverrideOutput {
	return o
}

func (o EdifactValidationOverrideOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationOverride) bool { return v.AllowLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolOutput)
}

func (o EdifactValidationOverrideOutput) EnforceCharacterSet() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationOverride) bool { return v.EnforceCharacterSet }).(pulumi.BoolOutput)
}

func (o EdifactValidationOverrideOutput) MessageId() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactValidationOverride) string { return v.MessageId }).(pulumi.StringOutput)
}

func (o EdifactValidationOverrideOutput) TrailingSeparatorPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactValidationOverride) string { return v.TrailingSeparatorPolicy }).(pulumi.StringOutput)
}

func (o EdifactValidationOverrideOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationOverride) bool { return v.TrimLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolOutput)
}

func (o EdifactValidationOverrideOutput) ValidateEDITypes() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationOverride) bool { return v.ValidateEDITypes }).(pulumi.BoolOutput)
}

func (o EdifactValidationOverrideOutput) ValidateXSDTypes() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationOverride) bool { return v.ValidateXSDTypes }).(pulumi.BoolOutput)
}

type EdifactValidationOverrideArrayOutput struct{ *pulumi.OutputState }

func (EdifactValidationOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactValidationOverride)(nil)).Elem()
}

func (o EdifactValidationOverrideArrayOutput) ToEdifactValidationOverrideArrayOutput() EdifactValidationOverrideArrayOutput {
	return o
}

func (o EdifactValidationOverrideArrayOutput) ToEdifactValidationOverrideArrayOutputWithContext(ctx context.Context) EdifactValidationOverrideArrayOutput {
	return o
}

func (o EdifactValidationOverrideArrayOutput) Index(i pulumi.IntInput) EdifactValidationOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactValidationOverride {
		return vs[0].([]EdifactValidationOverride)[vs[1].(int)]
	}).(EdifactValidationOverrideOutput)
}

type EdifactValidationOverrideResponse struct {
	AllowLeadingAndTrailingSpacesAndZeroes bool   `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	EnforceCharacterSet                    bool   `pulumi:"enforceCharacterSet"`
	MessageId                              string `pulumi:"messageId"`
	TrailingSeparatorPolicy                string `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes  bool   `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateEDITypes                       bool   `pulumi:"validateEDITypes"`
	ValidateXSDTypes                       bool   `pulumi:"validateXSDTypes"`
}

type EdifactValidationOverrideResponseOutput struct{ *pulumi.OutputState }

func (EdifactValidationOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactValidationOverrideResponse)(nil)).Elem()
}

func (o EdifactValidationOverrideResponseOutput) ToEdifactValidationOverrideResponseOutput() EdifactValidationOverrideResponseOutput {
	return o
}

func (o EdifactValidationOverrideResponseOutput) ToEdifactValidationOverrideResponseOutputWithContext(ctx context.Context) EdifactValidationOverrideResponseOutput {
	return o
}

func (o EdifactValidationOverrideResponseOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationOverrideResponse) bool { return v.AllowLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolOutput)
}

func (o EdifactValidationOverrideResponseOutput) EnforceCharacterSet() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationOverrideResponse) bool { return v.EnforceCharacterSet }).(pulumi.BoolOutput)
}

func (o EdifactValidationOverrideResponseOutput) MessageId() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactValidationOverrideResponse) string { return v.MessageId }).(pulumi.StringOutput)
}

func (o EdifactValidationOverrideResponseOutput) TrailingSeparatorPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactValidationOverrideResponse) string { return v.TrailingSeparatorPolicy }).(pulumi.StringOutput)
}

func (o EdifactValidationOverrideResponseOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationOverrideResponse) bool { return v.TrimLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolOutput)
}

func (o EdifactValidationOverrideResponseOutput) ValidateEDITypes() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationOverrideResponse) bool { return v.ValidateEDITypes }).(pulumi.BoolOutput)
}

func (o EdifactValidationOverrideResponseOutput) ValidateXSDTypes() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationOverrideResponse) bool { return v.ValidateXSDTypes }).(pulumi.BoolOutput)
}

type EdifactValidationOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (EdifactValidationOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EdifactValidationOverrideResponse)(nil)).Elem()
}

func (o EdifactValidationOverrideResponseArrayOutput) ToEdifactValidationOverrideResponseArrayOutput() EdifactValidationOverrideResponseArrayOutput {
	return o
}

func (o EdifactValidationOverrideResponseArrayOutput) ToEdifactValidationOverrideResponseArrayOutputWithContext(ctx context.Context) EdifactValidationOverrideResponseArrayOutput {
	return o
}

func (o EdifactValidationOverrideResponseArrayOutput) Index(i pulumi.IntInput) EdifactValidationOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EdifactValidationOverrideResponse {
		return vs[0].([]EdifactValidationOverrideResponse)[vs[1].(int)]
	}).(EdifactValidationOverrideResponseOutput)
}

type EdifactValidationSettings struct {
	AllowLeadingAndTrailingSpacesAndZeroes    bool   `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	CheckDuplicateGroupControlNumber          bool   `pulumi:"checkDuplicateGroupControlNumber"`
	CheckDuplicateInterchangeControlNumber    bool   `pulumi:"checkDuplicateInterchangeControlNumber"`
	CheckDuplicateTransactionSetControlNumber bool   `pulumi:"checkDuplicateTransactionSetControlNumber"`
	InterchangeControlNumberValidityDays      int    `pulumi:"interchangeControlNumberValidityDays"`
	TrailingSeparatorPolicy                   string `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes     bool   `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateCharacterSet                      bool   `pulumi:"validateCharacterSet"`
	ValidateEDITypes                          bool   `pulumi:"validateEDITypes"`
	ValidateXSDTypes                          bool   `pulumi:"validateXSDTypes"`
}





type EdifactValidationSettingsInput interface {
	pulumi.Input

	ToEdifactValidationSettingsOutput() EdifactValidationSettingsOutput
	ToEdifactValidationSettingsOutputWithContext(context.Context) EdifactValidationSettingsOutput
}

type EdifactValidationSettingsArgs struct {
	AllowLeadingAndTrailingSpacesAndZeroes    pulumi.BoolInput   `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	CheckDuplicateGroupControlNumber          pulumi.BoolInput   `pulumi:"checkDuplicateGroupControlNumber"`
	CheckDuplicateInterchangeControlNumber    pulumi.BoolInput   `pulumi:"checkDuplicateInterchangeControlNumber"`
	CheckDuplicateTransactionSetControlNumber pulumi.BoolInput   `pulumi:"checkDuplicateTransactionSetControlNumber"`
	InterchangeControlNumberValidityDays      pulumi.IntInput    `pulumi:"interchangeControlNumberValidityDays"`
	TrailingSeparatorPolicy                   pulumi.StringInput `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes     pulumi.BoolInput   `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateCharacterSet                      pulumi.BoolInput   `pulumi:"validateCharacterSet"`
	ValidateEDITypes                          pulumi.BoolInput   `pulumi:"validateEDITypes"`
	ValidateXSDTypes                          pulumi.BoolInput   `pulumi:"validateXSDTypes"`
}

func (EdifactValidationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactValidationSettings)(nil)).Elem()
}

func (i EdifactValidationSettingsArgs) ToEdifactValidationSettingsOutput() EdifactValidationSettingsOutput {
	return i.ToEdifactValidationSettingsOutputWithContext(context.Background())
}

func (i EdifactValidationSettingsArgs) ToEdifactValidationSettingsOutputWithContext(ctx context.Context) EdifactValidationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactValidationSettingsOutput)
}

func (i EdifactValidationSettingsArgs) ToEdifactValidationSettingsPtrOutput() EdifactValidationSettingsPtrOutput {
	return i.ToEdifactValidationSettingsPtrOutputWithContext(context.Background())
}

func (i EdifactValidationSettingsArgs) ToEdifactValidationSettingsPtrOutputWithContext(ctx context.Context) EdifactValidationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactValidationSettingsOutput).ToEdifactValidationSettingsPtrOutputWithContext(ctx)
}









type EdifactValidationSettingsPtrInput interface {
	pulumi.Input

	ToEdifactValidationSettingsPtrOutput() EdifactValidationSettingsPtrOutput
	ToEdifactValidationSettingsPtrOutputWithContext(context.Context) EdifactValidationSettingsPtrOutput
}

type edifactValidationSettingsPtrType EdifactValidationSettingsArgs

func EdifactValidationSettingsPtr(v *EdifactValidationSettingsArgs) EdifactValidationSettingsPtrInput {
	return (*edifactValidationSettingsPtrType)(v)
}

func (*edifactValidationSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactValidationSettings)(nil)).Elem()
}

func (i *edifactValidationSettingsPtrType) ToEdifactValidationSettingsPtrOutput() EdifactValidationSettingsPtrOutput {
	return i.ToEdifactValidationSettingsPtrOutputWithContext(context.Background())
}

func (i *edifactValidationSettingsPtrType) ToEdifactValidationSettingsPtrOutputWithContext(ctx context.Context) EdifactValidationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdifactValidationSettingsPtrOutput)
}

type EdifactValidationSettingsOutput struct{ *pulumi.OutputState }

func (EdifactValidationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactValidationSettings)(nil)).Elem()
}

func (o EdifactValidationSettingsOutput) ToEdifactValidationSettingsOutput() EdifactValidationSettingsOutput {
	return o
}

func (o EdifactValidationSettingsOutput) ToEdifactValidationSettingsOutputWithContext(ctx context.Context) EdifactValidationSettingsOutput {
	return o
}

func (o EdifactValidationSettingsOutput) ToEdifactValidationSettingsPtrOutput() EdifactValidationSettingsPtrOutput {
	return o.ToEdifactValidationSettingsPtrOutputWithContext(context.Background())
}

func (o EdifactValidationSettingsOutput) ToEdifactValidationSettingsPtrOutputWithContext(ctx context.Context) EdifactValidationSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactValidationSettings) *EdifactValidationSettings {
		return &v
	}).(EdifactValidationSettingsPtrOutput)
}

func (o EdifactValidationSettingsOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationSettings) bool { return v.AllowLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolOutput)
}

func (o EdifactValidationSettingsOutput) CheckDuplicateGroupControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationSettings) bool { return v.CheckDuplicateGroupControlNumber }).(pulumi.BoolOutput)
}

func (o EdifactValidationSettingsOutput) CheckDuplicateInterchangeControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationSettings) bool { return v.CheckDuplicateInterchangeControlNumber }).(pulumi.BoolOutput)
}

func (o EdifactValidationSettingsOutput) CheckDuplicateTransactionSetControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationSettings) bool { return v.CheckDuplicateTransactionSetControlNumber }).(pulumi.BoolOutput)
}

func (o EdifactValidationSettingsOutput) InterchangeControlNumberValidityDays() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactValidationSettings) int { return v.InterchangeControlNumberValidityDays }).(pulumi.IntOutput)
}

func (o EdifactValidationSettingsOutput) TrailingSeparatorPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactValidationSettings) string { return v.TrailingSeparatorPolicy }).(pulumi.StringOutput)
}

func (o EdifactValidationSettingsOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationSettings) bool { return v.TrimLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolOutput)
}

func (o EdifactValidationSettingsOutput) ValidateCharacterSet() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationSettings) bool { return v.ValidateCharacterSet }).(pulumi.BoolOutput)
}

func (o EdifactValidationSettingsOutput) ValidateEDITypes() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationSettings) bool { return v.ValidateEDITypes }).(pulumi.BoolOutput)
}

func (o EdifactValidationSettingsOutput) ValidateXSDTypes() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationSettings) bool { return v.ValidateXSDTypes }).(pulumi.BoolOutput)
}

type EdifactValidationSettingsPtrOutput struct{ *pulumi.OutputState }

func (EdifactValidationSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactValidationSettings)(nil)).Elem()
}

func (o EdifactValidationSettingsPtrOutput) ToEdifactValidationSettingsPtrOutput() EdifactValidationSettingsPtrOutput {
	return o
}

func (o EdifactValidationSettingsPtrOutput) ToEdifactValidationSettingsPtrOutputWithContext(ctx context.Context) EdifactValidationSettingsPtrOutput {
	return o
}

func (o EdifactValidationSettingsPtrOutput) Elem() EdifactValidationSettingsOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) EdifactValidationSettings {
		if v != nil {
			return *v
		}
		var ret EdifactValidationSettings
		return ret
	}).(EdifactValidationSettingsOutput)
}

func (o EdifactValidationSettingsPtrOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.AllowLeadingAndTrailingSpacesAndZeroes
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsPtrOutput) CheckDuplicateGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.CheckDuplicateGroupControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsPtrOutput) CheckDuplicateInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.CheckDuplicateInterchangeControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsPtrOutput) CheckDuplicateTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.CheckDuplicateTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsPtrOutput) InterchangeControlNumberValidityDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *int {
		if v == nil {
			return nil
		}
		return &v.InterchangeControlNumberValidityDays
	}).(pulumi.IntPtrOutput)
}

func (o EdifactValidationSettingsPtrOutput) TrailingSeparatorPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *string {
		if v == nil {
			return nil
		}
		return &v.TrailingSeparatorPolicy
	}).(pulumi.StringPtrOutput)
}

func (o EdifactValidationSettingsPtrOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.TrimLeadingAndTrailingSpacesAndZeroes
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsPtrOutput) ValidateCharacterSet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.ValidateCharacterSet
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsPtrOutput) ValidateEDITypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.ValidateEDITypes
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsPtrOutput) ValidateXSDTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.ValidateXSDTypes
	}).(pulumi.BoolPtrOutput)
}

type EdifactValidationSettingsResponse struct {
	AllowLeadingAndTrailingSpacesAndZeroes    bool   `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	CheckDuplicateGroupControlNumber          bool   `pulumi:"checkDuplicateGroupControlNumber"`
	CheckDuplicateInterchangeControlNumber    bool   `pulumi:"checkDuplicateInterchangeControlNumber"`
	CheckDuplicateTransactionSetControlNumber bool   `pulumi:"checkDuplicateTransactionSetControlNumber"`
	InterchangeControlNumberValidityDays      int    `pulumi:"interchangeControlNumberValidityDays"`
	TrailingSeparatorPolicy                   string `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes     bool   `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateCharacterSet                      bool   `pulumi:"validateCharacterSet"`
	ValidateEDITypes                          bool   `pulumi:"validateEDITypes"`
	ValidateXSDTypes                          bool   `pulumi:"validateXSDTypes"`
}

type EdifactValidationSettingsResponseOutput struct{ *pulumi.OutputState }

func (EdifactValidationSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactValidationSettingsResponse)(nil)).Elem()
}

func (o EdifactValidationSettingsResponseOutput) ToEdifactValidationSettingsResponseOutput() EdifactValidationSettingsResponseOutput {
	return o
}

func (o EdifactValidationSettingsResponseOutput) ToEdifactValidationSettingsResponseOutputWithContext(ctx context.Context) EdifactValidationSettingsResponseOutput {
	return o
}

func (o EdifactValidationSettingsResponseOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) bool { return v.AllowLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolOutput)
}

func (o EdifactValidationSettingsResponseOutput) CheckDuplicateGroupControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) bool { return v.CheckDuplicateGroupControlNumber }).(pulumi.BoolOutput)
}

func (o EdifactValidationSettingsResponseOutput) CheckDuplicateInterchangeControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) bool { return v.CheckDuplicateInterchangeControlNumber }).(pulumi.BoolOutput)
}

func (o EdifactValidationSettingsResponseOutput) CheckDuplicateTransactionSetControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) bool { return v.CheckDuplicateTransactionSetControlNumber }).(pulumi.BoolOutput)
}

func (o EdifactValidationSettingsResponseOutput) InterchangeControlNumberValidityDays() pulumi.IntOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) int { return v.InterchangeControlNumberValidityDays }).(pulumi.IntOutput)
}

func (o EdifactValidationSettingsResponseOutput) TrailingSeparatorPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) string { return v.TrailingSeparatorPolicy }).(pulumi.StringOutput)
}

func (o EdifactValidationSettingsResponseOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) bool { return v.TrimLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolOutput)
}

func (o EdifactValidationSettingsResponseOutput) ValidateCharacterSet() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) bool { return v.ValidateCharacterSet }).(pulumi.BoolOutput)
}

func (o EdifactValidationSettingsResponseOutput) ValidateEDITypes() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) bool { return v.ValidateEDITypes }).(pulumi.BoolOutput)
}

func (o EdifactValidationSettingsResponseOutput) ValidateXSDTypes() pulumi.BoolOutput {
	return o.ApplyT(func(v EdifactValidationSettingsResponse) bool { return v.ValidateXSDTypes }).(pulumi.BoolOutput)
}

type EdifactValidationSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (EdifactValidationSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactValidationSettingsResponse)(nil)).Elem()
}

func (o EdifactValidationSettingsResponsePtrOutput) ToEdifactValidationSettingsResponsePtrOutput() EdifactValidationSettingsResponsePtrOutput {
	return o
}

func (o EdifactValidationSettingsResponsePtrOutput) ToEdifactValidationSettingsResponsePtrOutputWithContext(ctx context.Context) EdifactValidationSettingsResponsePtrOutput {
	return o
}

func (o EdifactValidationSettingsResponsePtrOutput) Elem() EdifactValidationSettingsResponseOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) EdifactValidationSettingsResponse {
		if v != nil {
			return *v
		}
		var ret EdifactValidationSettingsResponse
		return ret
	}).(EdifactValidationSettingsResponseOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.AllowLeadingAndTrailingSpacesAndZeroes
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) CheckDuplicateGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CheckDuplicateGroupControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) CheckDuplicateInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CheckDuplicateInterchangeControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) CheckDuplicateTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CheckDuplicateTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) InterchangeControlNumberValidityDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.InterchangeControlNumberValidityDays
	}).(pulumi.IntPtrOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) TrailingSeparatorPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TrailingSeparatorPolicy
	}).(pulumi.StringPtrOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.TrimLeadingAndTrailingSpacesAndZeroes
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) ValidateCharacterSet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.ValidateCharacterSet
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) ValidateEDITypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.ValidateEDITypes
	}).(pulumi.BoolPtrOutput)
}

func (o EdifactValidationSettingsResponsePtrOutput) ValidateXSDTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdifactValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.ValidateXSDTypes
	}).(pulumi.BoolPtrOutput)
}

type ExpressionResponse struct {
	Error          *AzureResourceErrorInfoResponse `pulumi:"error"`
	Subexpressions []ExpressionResponse            `pulumi:"subexpressions"`
	Text           *string                         `pulumi:"text"`
	Value          interface{}                     `pulumi:"value"`
}

type ExpressionResponseOutput struct{ *pulumi.OutputState }

func (ExpressionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressionResponse)(nil)).Elem()
}

func (o ExpressionResponseOutput) ToExpressionResponseOutput() ExpressionResponseOutput {
	return o
}

func (o ExpressionResponseOutput) ToExpressionResponseOutputWithContext(ctx context.Context) ExpressionResponseOutput {
	return o
}

func (o ExpressionResponseOutput) Error() AzureResourceErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v ExpressionResponse) *AzureResourceErrorInfoResponse { return v.Error }).(AzureResourceErrorInfoResponsePtrOutput)
}

func (o ExpressionResponseOutput) Subexpressions() ExpressionResponseArrayOutput {
	return o.ApplyT(func(v ExpressionResponse) []ExpressionResponse { return v.Subexpressions }).(ExpressionResponseArrayOutput)
}

func (o ExpressionResponseOutput) Text() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressionResponse) *string { return v.Text }).(pulumi.StringPtrOutput)
}

func (o ExpressionResponseOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v ExpressionResponse) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type ExpressionResponseArrayOutput struct{ *pulumi.OutputState }

func (ExpressionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressionResponse)(nil)).Elem()
}

func (o ExpressionResponseArrayOutput) ToExpressionResponseArrayOutput() ExpressionResponseArrayOutput {
	return o
}

func (o ExpressionResponseArrayOutput) ToExpressionResponseArrayOutputWithContext(ctx context.Context) ExpressionResponseArrayOutput {
	return o
}

func (o ExpressionResponseArrayOutput) Index(i pulumi.IntInput) ExpressionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExpressionResponse {
		return vs[0].([]ExpressionResponse)[vs[1].(int)]
	}).(ExpressionResponseOutput)
}

type ExpressionRootResponse struct {
	Error          *AzureResourceErrorInfoResponse `pulumi:"error"`
	Path           *string                         `pulumi:"path"`
	Subexpressions []ExpressionResponse            `pulumi:"subexpressions"`
	Text           *string                         `pulumi:"text"`
	Value          interface{}                     `pulumi:"value"`
}

type ExpressionRootResponseOutput struct{ *pulumi.OutputState }

func (ExpressionRootResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressionRootResponse)(nil)).Elem()
}

func (o ExpressionRootResponseOutput) ToExpressionRootResponseOutput() ExpressionRootResponseOutput {
	return o
}

func (o ExpressionRootResponseOutput) ToExpressionRootResponseOutputWithContext(ctx context.Context) ExpressionRootResponseOutput {
	return o
}

func (o ExpressionRootResponseOutput) Error() AzureResourceErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v ExpressionRootResponse) *AzureResourceErrorInfoResponse { return v.Error }).(AzureResourceErrorInfoResponsePtrOutput)
}

func (o ExpressionRootResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressionRootResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ExpressionRootResponseOutput) Subexpressions() ExpressionResponseArrayOutput {
	return o.ApplyT(func(v ExpressionRootResponse) []ExpressionResponse { return v.Subexpressions }).(ExpressionResponseArrayOutput)
}

func (o ExpressionRootResponseOutput) Text() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressionRootResponse) *string { return v.Text }).(pulumi.StringPtrOutput)
}

func (o ExpressionRootResponseOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v ExpressionRootResponse) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type ExpressionRootResponseArrayOutput struct{ *pulumi.OutputState }

func (ExpressionRootResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressionRootResponse)(nil)).Elem()
}

func (o ExpressionRootResponseArrayOutput) ToExpressionRootResponseArrayOutput() ExpressionRootResponseArrayOutput {
	return o
}

func (o ExpressionRootResponseArrayOutput) ToExpressionRootResponseArrayOutputWithContext(ctx context.Context) ExpressionRootResponseArrayOutput {
	return o
}

func (o ExpressionRootResponseArrayOutput) Index(i pulumi.IntInput) ExpressionRootResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExpressionRootResponse {
		return vs[0].([]ExpressionRootResponse)[vs[1].(int)]
	}).(ExpressionRootResponseOutput)
}

type FlowAccessControlConfiguration struct {
	Actions            *FlowAccessControlConfigurationPolicy `pulumi:"actions"`
	Contents           *FlowAccessControlConfigurationPolicy `pulumi:"contents"`
	Triggers           *FlowAccessControlConfigurationPolicy `pulumi:"triggers"`
	WorkflowManagement *FlowAccessControlConfigurationPolicy `pulumi:"workflowManagement"`
}





type FlowAccessControlConfigurationInput interface {
	pulumi.Input

	ToFlowAccessControlConfigurationOutput() FlowAccessControlConfigurationOutput
	ToFlowAccessControlConfigurationOutputWithContext(context.Context) FlowAccessControlConfigurationOutput
}

type FlowAccessControlConfigurationArgs struct {
	Actions            FlowAccessControlConfigurationPolicyPtrInput `pulumi:"actions"`
	Contents           FlowAccessControlConfigurationPolicyPtrInput `pulumi:"contents"`
	Triggers           FlowAccessControlConfigurationPolicyPtrInput `pulumi:"triggers"`
	WorkflowManagement FlowAccessControlConfigurationPolicyPtrInput `pulumi:"workflowManagement"`
}

func (FlowAccessControlConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowAccessControlConfiguration)(nil)).Elem()
}

func (i FlowAccessControlConfigurationArgs) ToFlowAccessControlConfigurationOutput() FlowAccessControlConfigurationOutput {
	return i.ToFlowAccessControlConfigurationOutputWithContext(context.Background())
}

func (i FlowAccessControlConfigurationArgs) ToFlowAccessControlConfigurationOutputWithContext(ctx context.Context) FlowAccessControlConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowAccessControlConfigurationOutput)
}

func (i FlowAccessControlConfigurationArgs) ToFlowAccessControlConfigurationPtrOutput() FlowAccessControlConfigurationPtrOutput {
	return i.ToFlowAccessControlConfigurationPtrOutputWithContext(context.Background())
}

func (i FlowAccessControlConfigurationArgs) ToFlowAccessControlConfigurationPtrOutputWithContext(ctx context.Context) FlowAccessControlConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowAccessControlConfigurationOutput).ToFlowAccessControlConfigurationPtrOutputWithContext(ctx)
}









type FlowAccessControlConfigurationPtrInput interface {
	pulumi.Input

	ToFlowAccessControlConfigurationPtrOutput() FlowAccessControlConfigurationPtrOutput
	ToFlowAccessControlConfigurationPtrOutputWithContext(context.Context) FlowAccessControlConfigurationPtrOutput
}

type flowAccessControlConfigurationPtrType FlowAccessControlConfigurationArgs

func FlowAccessControlConfigurationPtr(v *FlowAccessControlConfigurationArgs) FlowAccessControlConfigurationPtrInput {
	return (*flowAccessControlConfigurationPtrType)(v)
}

func (*flowAccessControlConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowAccessControlConfiguration)(nil)).Elem()
}

func (i *flowAccessControlConfigurationPtrType) ToFlowAccessControlConfigurationPtrOutput() FlowAccessControlConfigurationPtrOutput {
	return i.ToFlowAccessControlConfigurationPtrOutputWithContext(context.Background())
}

func (i *flowAccessControlConfigurationPtrType) ToFlowAccessControlConfigurationPtrOutputWithContext(ctx context.Context) FlowAccessControlConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowAccessControlConfigurationPtrOutput)
}

type FlowAccessControlConfigurationOutput struct{ *pulumi.OutputState }

func (FlowAccessControlConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowAccessControlConfiguration)(nil)).Elem()
}

func (o FlowAccessControlConfigurationOutput) ToFlowAccessControlConfigurationOutput() FlowAccessControlConfigurationOutput {
	return o
}

func (o FlowAccessControlConfigurationOutput) ToFlowAccessControlConfigurationOutputWithContext(ctx context.Context) FlowAccessControlConfigurationOutput {
	return o
}

func (o FlowAccessControlConfigurationOutput) ToFlowAccessControlConfigurationPtrOutput() FlowAccessControlConfigurationPtrOutput {
	return o.ToFlowAccessControlConfigurationPtrOutputWithContext(context.Background())
}

func (o FlowAccessControlConfigurationOutput) ToFlowAccessControlConfigurationPtrOutputWithContext(ctx context.Context) FlowAccessControlConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FlowAccessControlConfiguration) *FlowAccessControlConfiguration {
		return &v
	}).(FlowAccessControlConfigurationPtrOutput)
}

func (o FlowAccessControlConfigurationOutput) Actions() FlowAccessControlConfigurationPolicyPtrOutput {
	return o.ApplyT(func(v FlowAccessControlConfiguration) *FlowAccessControlConfigurationPolicy { return v.Actions }).(FlowAccessControlConfigurationPolicyPtrOutput)
}

func (o FlowAccessControlConfigurationOutput) Contents() FlowAccessControlConfigurationPolicyPtrOutput {
	return o.ApplyT(func(v FlowAccessControlConfiguration) *FlowAccessControlConfigurationPolicy { return v.Contents }).(FlowAccessControlConfigurationPolicyPtrOutput)
}

func (o FlowAccessControlConfigurationOutput) Triggers() FlowAccessControlConfigurationPolicyPtrOutput {
	return o.ApplyT(func(v FlowAccessControlConfiguration) *FlowAccessControlConfigurationPolicy { return v.Triggers }).(FlowAccessControlConfigurationPolicyPtrOutput)
}

func (o FlowAccessControlConfigurationOutput) WorkflowManagement() FlowAccessControlConfigurationPolicyPtrOutput {
	return o.ApplyT(func(v FlowAccessControlConfiguration) *FlowAccessControlConfigurationPolicy {
		return v.WorkflowManagement
	}).(FlowAccessControlConfigurationPolicyPtrOutput)
}

type FlowAccessControlConfigurationPtrOutput struct{ *pulumi.OutputState }

func (FlowAccessControlConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowAccessControlConfiguration)(nil)).Elem()
}

func (o FlowAccessControlConfigurationPtrOutput) ToFlowAccessControlConfigurationPtrOutput() FlowAccessControlConfigurationPtrOutput {
	return o
}

func (o FlowAccessControlConfigurationPtrOutput) ToFlowAccessControlConfigurationPtrOutputWithContext(ctx context.Context) FlowAccessControlConfigurationPtrOutput {
	return o
}

func (o FlowAccessControlConfigurationPtrOutput) Elem() FlowAccessControlConfigurationOutput {
	return o.ApplyT(func(v *FlowAccessControlConfiguration) FlowAccessControlConfiguration {
		if v != nil {
			return *v
		}
		var ret FlowAccessControlConfiguration
		return ret
	}).(FlowAccessControlConfigurationOutput)
}

func (o FlowAccessControlConfigurationPtrOutput) Actions() FlowAccessControlConfigurationPolicyPtrOutput {
	return o.ApplyT(func(v *FlowAccessControlConfiguration) *FlowAccessControlConfigurationPolicy {
		if v == nil {
			return nil
		}
		return v.Actions
	}).(FlowAccessControlConfigurationPolicyPtrOutput)
}

func (o FlowAccessControlConfigurationPtrOutput) Contents() FlowAccessControlConfigurationPolicyPtrOutput {
	return o.ApplyT(func(v *FlowAccessControlConfiguration) *FlowAccessControlConfigurationPolicy {
		if v == nil {
			return nil
		}
		return v.Contents
	}).(FlowAccessControlConfigurationPolicyPtrOutput)
}

func (o FlowAccessControlConfigurationPtrOutput) Triggers() FlowAccessControlConfigurationPolicyPtrOutput {
	return o.ApplyT(func(v *FlowAccessControlConfiguration) *FlowAccessControlConfigurationPolicy {
		if v == nil {
			return nil
		}
		return v.Triggers
	}).(FlowAccessControlConfigurationPolicyPtrOutput)
}

func (o FlowAccessControlConfigurationPtrOutput) WorkflowManagement() FlowAccessControlConfigurationPolicyPtrOutput {
	return o.ApplyT(func(v *FlowAccessControlConfiguration) *FlowAccessControlConfigurationPolicy {
		if v == nil {
			return nil
		}
		return v.WorkflowManagement
	}).(FlowAccessControlConfigurationPolicyPtrOutput)
}

type FlowAccessControlConfigurationPolicy struct {
	AllowedCallerIpAddresses   []IpAddressRange                  `pulumi:"allowedCallerIpAddresses"`
	OpenAuthenticationPolicies *OpenAuthenticationAccessPolicies `pulumi:"openAuthenticationPolicies"`
}





type FlowAccessControlConfigurationPolicyInput interface {
	pulumi.Input

	ToFlowAccessControlConfigurationPolicyOutput() FlowAccessControlConfigurationPolicyOutput
	ToFlowAccessControlConfigurationPolicyOutputWithContext(context.Context) FlowAccessControlConfigurationPolicyOutput
}

type FlowAccessControlConfigurationPolicyArgs struct {
	AllowedCallerIpAddresses   IpAddressRangeArrayInput                 `pulumi:"allowedCallerIpAddresses"`
	OpenAuthenticationPolicies OpenAuthenticationAccessPoliciesPtrInput `pulumi:"openAuthenticationPolicies"`
}

func (FlowAccessControlConfigurationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowAccessControlConfigurationPolicy)(nil)).Elem()
}

func (i FlowAccessControlConfigurationPolicyArgs) ToFlowAccessControlConfigurationPolicyOutput() FlowAccessControlConfigurationPolicyOutput {
	return i.ToFlowAccessControlConfigurationPolicyOutputWithContext(context.Background())
}

func (i FlowAccessControlConfigurationPolicyArgs) ToFlowAccessControlConfigurationPolicyOutputWithContext(ctx context.Context) FlowAccessControlConfigurationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowAccessControlConfigurationPolicyOutput)
}

func (i FlowAccessControlConfigurationPolicyArgs) ToFlowAccessControlConfigurationPolicyPtrOutput() FlowAccessControlConfigurationPolicyPtrOutput {
	return i.ToFlowAccessControlConfigurationPolicyPtrOutputWithContext(context.Background())
}

func (i FlowAccessControlConfigurationPolicyArgs) ToFlowAccessControlConfigurationPolicyPtrOutputWithContext(ctx context.Context) FlowAccessControlConfigurationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowAccessControlConfigurationPolicyOutput).ToFlowAccessControlConfigurationPolicyPtrOutputWithContext(ctx)
}









type FlowAccessControlConfigurationPolicyPtrInput interface {
	pulumi.Input

	ToFlowAccessControlConfigurationPolicyPtrOutput() FlowAccessControlConfigurationPolicyPtrOutput
	ToFlowAccessControlConfigurationPolicyPtrOutputWithContext(context.Context) FlowAccessControlConfigurationPolicyPtrOutput
}

type flowAccessControlConfigurationPolicyPtrType FlowAccessControlConfigurationPolicyArgs

func FlowAccessControlConfigurationPolicyPtr(v *FlowAccessControlConfigurationPolicyArgs) FlowAccessControlConfigurationPolicyPtrInput {
	return (*flowAccessControlConfigurationPolicyPtrType)(v)
}

func (*flowAccessControlConfigurationPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowAccessControlConfigurationPolicy)(nil)).Elem()
}

func (i *flowAccessControlConfigurationPolicyPtrType) ToFlowAccessControlConfigurationPolicyPtrOutput() FlowAccessControlConfigurationPolicyPtrOutput {
	return i.ToFlowAccessControlConfigurationPolicyPtrOutputWithContext(context.Background())
}

func (i *flowAccessControlConfigurationPolicyPtrType) ToFlowAccessControlConfigurationPolicyPtrOutputWithContext(ctx context.Context) FlowAccessControlConfigurationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowAccessControlConfigurationPolicyPtrOutput)
}

type FlowAccessControlConfigurationPolicyOutput struct{ *pulumi.OutputState }

func (FlowAccessControlConfigurationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowAccessControlConfigurationPolicy)(nil)).Elem()
}

func (o FlowAccessControlConfigurationPolicyOutput) ToFlowAccessControlConfigurationPolicyOutput() FlowAccessControlConfigurationPolicyOutput {
	return o
}

func (o FlowAccessControlConfigurationPolicyOutput) ToFlowAccessControlConfigurationPolicyOutputWithContext(ctx context.Context) FlowAccessControlConfigurationPolicyOutput {
	return o
}

func (o FlowAccessControlConfigurationPolicyOutput) ToFlowAccessControlConfigurationPolicyPtrOutput() FlowAccessControlConfigurationPolicyPtrOutput {
	return o.ToFlowAccessControlConfigurationPolicyPtrOutputWithContext(context.Background())
}

func (o FlowAccessControlConfigurationPolicyOutput) ToFlowAccessControlConfigurationPolicyPtrOutputWithContext(ctx context.Context) FlowAccessControlConfigurationPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FlowAccessControlConfigurationPolicy) *FlowAccessControlConfigurationPolicy {
		return &v
	}).(FlowAccessControlConfigurationPolicyPtrOutput)
}

func (o FlowAccessControlConfigurationPolicyOutput) AllowedCallerIpAddresses() IpAddressRangeArrayOutput {
	return o.ApplyT(func(v FlowAccessControlConfigurationPolicy) []IpAddressRange { return v.AllowedCallerIpAddresses }).(IpAddressRangeArrayOutput)
}

func (o FlowAccessControlConfigurationPolicyOutput) OpenAuthenticationPolicies() OpenAuthenticationAccessPoliciesPtrOutput {
	return o.ApplyT(func(v FlowAccessControlConfigurationPolicy) *OpenAuthenticationAccessPolicies {
		return v.OpenAuthenticationPolicies
	}).(OpenAuthenticationAccessPoliciesPtrOutput)
}

type FlowAccessControlConfigurationPolicyPtrOutput struct{ *pulumi.OutputState }

func (FlowAccessControlConfigurationPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowAccessControlConfigurationPolicy)(nil)).Elem()
}

func (o FlowAccessControlConfigurationPolicyPtrOutput) ToFlowAccessControlConfigurationPolicyPtrOutput() FlowAccessControlConfigurationPolicyPtrOutput {
	return o
}

func (o FlowAccessControlConfigurationPolicyPtrOutput) ToFlowAccessControlConfigurationPolicyPtrOutputWithContext(ctx context.Context) FlowAccessControlConfigurationPolicyPtrOutput {
	return o
}

func (o FlowAccessControlConfigurationPolicyPtrOutput) Elem() FlowAccessControlConfigurationPolicyOutput {
	return o.ApplyT(func(v *FlowAccessControlConfigurationPolicy) FlowAccessControlConfigurationPolicy {
		if v != nil {
			return *v
		}
		var ret FlowAccessControlConfigurationPolicy
		return ret
	}).(FlowAccessControlConfigurationPolicyOutput)
}

func (o FlowAccessControlConfigurationPolicyPtrOutput) AllowedCallerIpAddresses() IpAddressRangeArrayOutput {
	return o.ApplyT(func(v *FlowAccessControlConfigurationPolicy) []IpAddressRange {
		if v == nil {
			return nil
		}
		return v.AllowedCallerIpAddresses
	}).(IpAddressRangeArrayOutput)
}

func (o FlowAccessControlConfigurationPolicyPtrOutput) OpenAuthenticationPolicies() OpenAuthenticationAccessPoliciesPtrOutput {
	return o.ApplyT(func(v *FlowAccessControlConfigurationPolicy) *OpenAuthenticationAccessPolicies {
		if v == nil {
			return nil
		}
		return v.OpenAuthenticationPolicies
	}).(OpenAuthenticationAccessPoliciesPtrOutput)
}

type FlowAccessControlConfigurationPolicyResponse struct {
	AllowedCallerIpAddresses   []IpAddressRangeResponse                  `pulumi:"allowedCallerIpAddresses"`
	OpenAuthenticationPolicies *OpenAuthenticationAccessPoliciesResponse `pulumi:"openAuthenticationPolicies"`
}

type FlowAccessControlConfigurationPolicyResponseOutput struct{ *pulumi.OutputState }

func (FlowAccessControlConfigurationPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowAccessControlConfigurationPolicyResponse)(nil)).Elem()
}

func (o FlowAccessControlConfigurationPolicyResponseOutput) ToFlowAccessControlConfigurationPolicyResponseOutput() FlowAccessControlConfigurationPolicyResponseOutput {
	return o
}

func (o FlowAccessControlConfigurationPolicyResponseOutput) ToFlowAccessControlConfigurationPolicyResponseOutputWithContext(ctx context.Context) FlowAccessControlConfigurationPolicyResponseOutput {
	return o
}

func (o FlowAccessControlConfigurationPolicyResponseOutput) AllowedCallerIpAddresses() IpAddressRangeResponseArrayOutput {
	return o.ApplyT(func(v FlowAccessControlConfigurationPolicyResponse) []IpAddressRangeResponse {
		return v.AllowedCallerIpAddresses
	}).(IpAddressRangeResponseArrayOutput)
}

func (o FlowAccessControlConfigurationPolicyResponseOutput) OpenAuthenticationPolicies() OpenAuthenticationAccessPoliciesResponsePtrOutput {
	return o.ApplyT(func(v FlowAccessControlConfigurationPolicyResponse) *OpenAuthenticationAccessPoliciesResponse {
		return v.OpenAuthenticationPolicies
	}).(OpenAuthenticationAccessPoliciesResponsePtrOutput)
}

type FlowAccessControlConfigurationPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (FlowAccessControlConfigurationPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowAccessControlConfigurationPolicyResponse)(nil)).Elem()
}

func (o FlowAccessControlConfigurationPolicyResponsePtrOutput) ToFlowAccessControlConfigurationPolicyResponsePtrOutput() FlowAccessControlConfigurationPolicyResponsePtrOutput {
	return o
}

func (o FlowAccessControlConfigurationPolicyResponsePtrOutput) ToFlowAccessControlConfigurationPolicyResponsePtrOutputWithContext(ctx context.Context) FlowAccessControlConfigurationPolicyResponsePtrOutput {
	return o
}

func (o FlowAccessControlConfigurationPolicyResponsePtrOutput) Elem() FlowAccessControlConfigurationPolicyResponseOutput {
	return o.ApplyT(func(v *FlowAccessControlConfigurationPolicyResponse) FlowAccessControlConfigurationPolicyResponse {
		if v != nil {
			return *v
		}
		var ret FlowAccessControlConfigurationPolicyResponse
		return ret
	}).(FlowAccessControlConfigurationPolicyResponseOutput)
}

func (o FlowAccessControlConfigurationPolicyResponsePtrOutput) AllowedCallerIpAddresses() IpAddressRangeResponseArrayOutput {
	return o.ApplyT(func(v *FlowAccessControlConfigurationPolicyResponse) []IpAddressRangeResponse {
		if v == nil {
			return nil
		}
		return v.AllowedCallerIpAddresses
	}).(IpAddressRangeResponseArrayOutput)
}

func (o FlowAccessControlConfigurationPolicyResponsePtrOutput) OpenAuthenticationPolicies() OpenAuthenticationAccessPoliciesResponsePtrOutput {
	return o.ApplyT(func(v *FlowAccessControlConfigurationPolicyResponse) *OpenAuthenticationAccessPoliciesResponse {
		if v == nil {
			return nil
		}
		return v.OpenAuthenticationPolicies
	}).(OpenAuthenticationAccessPoliciesResponsePtrOutput)
}

type FlowAccessControlConfigurationResponse struct {
	Actions            *FlowAccessControlConfigurationPolicyResponse `pulumi:"actions"`
	Contents           *FlowAccessControlConfigurationPolicyResponse `pulumi:"contents"`
	Triggers           *FlowAccessControlConfigurationPolicyResponse `pulumi:"triggers"`
	WorkflowManagement *FlowAccessControlConfigurationPolicyResponse `pulumi:"workflowManagement"`
}

type FlowAccessControlConfigurationResponseOutput struct{ *pulumi.OutputState }

func (FlowAccessControlConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowAccessControlConfigurationResponse)(nil)).Elem()
}

func (o FlowAccessControlConfigurationResponseOutput) ToFlowAccessControlConfigurationResponseOutput() FlowAccessControlConfigurationResponseOutput {
	return o
}

func (o FlowAccessControlConfigurationResponseOutput) ToFlowAccessControlConfigurationResponseOutputWithContext(ctx context.Context) FlowAccessControlConfigurationResponseOutput {
	return o
}

func (o FlowAccessControlConfigurationResponseOutput) Actions() FlowAccessControlConfigurationPolicyResponsePtrOutput {
	return o.ApplyT(func(v FlowAccessControlConfigurationResponse) *FlowAccessControlConfigurationPolicyResponse {
		return v.Actions
	}).(FlowAccessControlConfigurationPolicyResponsePtrOutput)
}

func (o FlowAccessControlConfigurationResponseOutput) Contents() FlowAccessControlConfigurationPolicyResponsePtrOutput {
	return o.ApplyT(func(v FlowAccessControlConfigurationResponse) *FlowAccessControlConfigurationPolicyResponse {
		return v.Contents
	}).(FlowAccessControlConfigurationPolicyResponsePtrOutput)
}

func (o FlowAccessControlConfigurationResponseOutput) Triggers() FlowAccessControlConfigurationPolicyResponsePtrOutput {
	return o.ApplyT(func(v FlowAccessControlConfigurationResponse) *FlowAccessControlConfigurationPolicyResponse {
		return v.Triggers
	}).(FlowAccessControlConfigurationPolicyResponsePtrOutput)
}

func (o FlowAccessControlConfigurationResponseOutput) WorkflowManagement() FlowAccessControlConfigurationPolicyResponsePtrOutput {
	return o.ApplyT(func(v FlowAccessControlConfigurationResponse) *FlowAccessControlConfigurationPolicyResponse {
		return v.WorkflowManagement
	}).(FlowAccessControlConfigurationPolicyResponsePtrOutput)
}

type FlowAccessControlConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (FlowAccessControlConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowAccessControlConfigurationResponse)(nil)).Elem()
}

func (o FlowAccessControlConfigurationResponsePtrOutput) ToFlowAccessControlConfigurationResponsePtrOutput() FlowAccessControlConfigurationResponsePtrOutput {
	return o
}

func (o FlowAccessControlConfigurationResponsePtrOutput) ToFlowAccessControlConfigurationResponsePtrOutputWithContext(ctx context.Context) FlowAccessControlConfigurationResponsePtrOutput {
	return o
}

func (o FlowAccessControlConfigurationResponsePtrOutput) Elem() FlowAccessControlConfigurationResponseOutput {
	return o.ApplyT(func(v *FlowAccessControlConfigurationResponse) FlowAccessControlConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret FlowAccessControlConfigurationResponse
		return ret
	}).(FlowAccessControlConfigurationResponseOutput)
}

func (o FlowAccessControlConfigurationResponsePtrOutput) Actions() FlowAccessControlConfigurationPolicyResponsePtrOutput {
	return o.ApplyT(func(v *FlowAccessControlConfigurationResponse) *FlowAccessControlConfigurationPolicyResponse {
		if v == nil {
			return nil
		}
		return v.Actions
	}).(FlowAccessControlConfigurationPolicyResponsePtrOutput)
}

func (o FlowAccessControlConfigurationResponsePtrOutput) Contents() FlowAccessControlConfigurationPolicyResponsePtrOutput {
	return o.ApplyT(func(v *FlowAccessControlConfigurationResponse) *FlowAccessControlConfigurationPolicyResponse {
		if v == nil {
			return nil
		}
		return v.Contents
	}).(FlowAccessControlConfigurationPolicyResponsePtrOutput)
}

func (o FlowAccessControlConfigurationResponsePtrOutput) Triggers() FlowAccessControlConfigurationPolicyResponsePtrOutput {
	return o.ApplyT(func(v *FlowAccessControlConfigurationResponse) *FlowAccessControlConfigurationPolicyResponse {
		if v == nil {
			return nil
		}
		return v.Triggers
	}).(FlowAccessControlConfigurationPolicyResponsePtrOutput)
}

func (o FlowAccessControlConfigurationResponsePtrOutput) WorkflowManagement() FlowAccessControlConfigurationPolicyResponsePtrOutput {
	return o.ApplyT(func(v *FlowAccessControlConfigurationResponse) *FlowAccessControlConfigurationPolicyResponse {
		if v == nil {
			return nil
		}
		return v.WorkflowManagement
	}).(FlowAccessControlConfigurationPolicyResponsePtrOutput)
}

type FlowEndpoints struct {
	AccessEndpointIpAddresses []IpAddress `pulumi:"accessEndpointIpAddresses"`
	OutgoingIpAddresses       []IpAddress `pulumi:"outgoingIpAddresses"`
}





type FlowEndpointsInput interface {
	pulumi.Input

	ToFlowEndpointsOutput() FlowEndpointsOutput
	ToFlowEndpointsOutputWithContext(context.Context) FlowEndpointsOutput
}

type FlowEndpointsArgs struct {
	AccessEndpointIpAddresses IpAddressArrayInput `pulumi:"accessEndpointIpAddresses"`
	OutgoingIpAddresses       IpAddressArrayInput `pulumi:"outgoingIpAddresses"`
}

func (FlowEndpointsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowEndpoints)(nil)).Elem()
}

func (i FlowEndpointsArgs) ToFlowEndpointsOutput() FlowEndpointsOutput {
	return i.ToFlowEndpointsOutputWithContext(context.Background())
}

func (i FlowEndpointsArgs) ToFlowEndpointsOutputWithContext(ctx context.Context) FlowEndpointsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowEndpointsOutput)
}

func (i FlowEndpointsArgs) ToFlowEndpointsPtrOutput() FlowEndpointsPtrOutput {
	return i.ToFlowEndpointsPtrOutputWithContext(context.Background())
}

func (i FlowEndpointsArgs) ToFlowEndpointsPtrOutputWithContext(ctx context.Context) FlowEndpointsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowEndpointsOutput).ToFlowEndpointsPtrOutputWithContext(ctx)
}









type FlowEndpointsPtrInput interface {
	pulumi.Input

	ToFlowEndpointsPtrOutput() FlowEndpointsPtrOutput
	ToFlowEndpointsPtrOutputWithContext(context.Context) FlowEndpointsPtrOutput
}

type flowEndpointsPtrType FlowEndpointsArgs

func FlowEndpointsPtr(v *FlowEndpointsArgs) FlowEndpointsPtrInput {
	return (*flowEndpointsPtrType)(v)
}

func (*flowEndpointsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowEndpoints)(nil)).Elem()
}

func (i *flowEndpointsPtrType) ToFlowEndpointsPtrOutput() FlowEndpointsPtrOutput {
	return i.ToFlowEndpointsPtrOutputWithContext(context.Background())
}

func (i *flowEndpointsPtrType) ToFlowEndpointsPtrOutputWithContext(ctx context.Context) FlowEndpointsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowEndpointsPtrOutput)
}

type FlowEndpointsOutput struct{ *pulumi.OutputState }

func (FlowEndpointsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowEndpoints)(nil)).Elem()
}

func (o FlowEndpointsOutput) ToFlowEndpointsOutput() FlowEndpointsOutput {
	return o
}

func (o FlowEndpointsOutput) ToFlowEndpointsOutputWithContext(ctx context.Context) FlowEndpointsOutput {
	return o
}

func (o FlowEndpointsOutput) ToFlowEndpointsPtrOutput() FlowEndpointsPtrOutput {
	return o.ToFlowEndpointsPtrOutputWithContext(context.Background())
}

func (o FlowEndpointsOutput) ToFlowEndpointsPtrOutputWithContext(ctx context.Context) FlowEndpointsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FlowEndpoints) *FlowEndpoints {
		return &v
	}).(FlowEndpointsPtrOutput)
}

func (o FlowEndpointsOutput) AccessEndpointIpAddresses() IpAddressArrayOutput {
	return o.ApplyT(func(v FlowEndpoints) []IpAddress { return v.AccessEndpointIpAddresses }).(IpAddressArrayOutput)
}

func (o FlowEndpointsOutput) OutgoingIpAddresses() IpAddressArrayOutput {
	return o.ApplyT(func(v FlowEndpoints) []IpAddress { return v.OutgoingIpAddresses }).(IpAddressArrayOutput)
}

type FlowEndpointsPtrOutput struct{ *pulumi.OutputState }

func (FlowEndpointsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowEndpoints)(nil)).Elem()
}

func (o FlowEndpointsPtrOutput) ToFlowEndpointsPtrOutput() FlowEndpointsPtrOutput {
	return o
}

func (o FlowEndpointsPtrOutput) ToFlowEndpointsPtrOutputWithContext(ctx context.Context) FlowEndpointsPtrOutput {
	return o
}

func (o FlowEndpointsPtrOutput) Elem() FlowEndpointsOutput {
	return o.ApplyT(func(v *FlowEndpoints) FlowEndpoints {
		if v != nil {
			return *v
		}
		var ret FlowEndpoints
		return ret
	}).(FlowEndpointsOutput)
}

func (o FlowEndpointsPtrOutput) AccessEndpointIpAddresses() IpAddressArrayOutput {
	return o.ApplyT(func(v *FlowEndpoints) []IpAddress {
		if v == nil {
			return nil
		}
		return v.AccessEndpointIpAddresses
	}).(IpAddressArrayOutput)
}

func (o FlowEndpointsPtrOutput) OutgoingIpAddresses() IpAddressArrayOutput {
	return o.ApplyT(func(v *FlowEndpoints) []IpAddress {
		if v == nil {
			return nil
		}
		return v.OutgoingIpAddresses
	}).(IpAddressArrayOutput)
}

type FlowEndpointsConfiguration struct {
	Connector *FlowEndpoints `pulumi:"connector"`
	Workflow  *FlowEndpoints `pulumi:"workflow"`
}





type FlowEndpointsConfigurationInput interface {
	pulumi.Input

	ToFlowEndpointsConfigurationOutput() FlowEndpointsConfigurationOutput
	ToFlowEndpointsConfigurationOutputWithContext(context.Context) FlowEndpointsConfigurationOutput
}

type FlowEndpointsConfigurationArgs struct {
	Connector FlowEndpointsPtrInput `pulumi:"connector"`
	Workflow  FlowEndpointsPtrInput `pulumi:"workflow"`
}

func (FlowEndpointsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowEndpointsConfiguration)(nil)).Elem()
}

func (i FlowEndpointsConfigurationArgs) ToFlowEndpointsConfigurationOutput() FlowEndpointsConfigurationOutput {
	return i.ToFlowEndpointsConfigurationOutputWithContext(context.Background())
}

func (i FlowEndpointsConfigurationArgs) ToFlowEndpointsConfigurationOutputWithContext(ctx context.Context) FlowEndpointsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowEndpointsConfigurationOutput)
}

func (i FlowEndpointsConfigurationArgs) ToFlowEndpointsConfigurationPtrOutput() FlowEndpointsConfigurationPtrOutput {
	return i.ToFlowEndpointsConfigurationPtrOutputWithContext(context.Background())
}

func (i FlowEndpointsConfigurationArgs) ToFlowEndpointsConfigurationPtrOutputWithContext(ctx context.Context) FlowEndpointsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowEndpointsConfigurationOutput).ToFlowEndpointsConfigurationPtrOutputWithContext(ctx)
}









type FlowEndpointsConfigurationPtrInput interface {
	pulumi.Input

	ToFlowEndpointsConfigurationPtrOutput() FlowEndpointsConfigurationPtrOutput
	ToFlowEndpointsConfigurationPtrOutputWithContext(context.Context) FlowEndpointsConfigurationPtrOutput
}

type flowEndpointsConfigurationPtrType FlowEndpointsConfigurationArgs

func FlowEndpointsConfigurationPtr(v *FlowEndpointsConfigurationArgs) FlowEndpointsConfigurationPtrInput {
	return (*flowEndpointsConfigurationPtrType)(v)
}

func (*flowEndpointsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowEndpointsConfiguration)(nil)).Elem()
}

func (i *flowEndpointsConfigurationPtrType) ToFlowEndpointsConfigurationPtrOutput() FlowEndpointsConfigurationPtrOutput {
	return i.ToFlowEndpointsConfigurationPtrOutputWithContext(context.Background())
}

func (i *flowEndpointsConfigurationPtrType) ToFlowEndpointsConfigurationPtrOutputWithContext(ctx context.Context) FlowEndpointsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowEndpointsConfigurationPtrOutput)
}

type FlowEndpointsConfigurationOutput struct{ *pulumi.OutputState }

func (FlowEndpointsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowEndpointsConfiguration)(nil)).Elem()
}

func (o FlowEndpointsConfigurationOutput) ToFlowEndpointsConfigurationOutput() FlowEndpointsConfigurationOutput {
	return o
}

func (o FlowEndpointsConfigurationOutput) ToFlowEndpointsConfigurationOutputWithContext(ctx context.Context) FlowEndpointsConfigurationOutput {
	return o
}

func (o FlowEndpointsConfigurationOutput) ToFlowEndpointsConfigurationPtrOutput() FlowEndpointsConfigurationPtrOutput {
	return o.ToFlowEndpointsConfigurationPtrOutputWithContext(context.Background())
}

func (o FlowEndpointsConfigurationOutput) ToFlowEndpointsConfigurationPtrOutputWithContext(ctx context.Context) FlowEndpointsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FlowEndpointsConfiguration) *FlowEndpointsConfiguration {
		return &v
	}).(FlowEndpointsConfigurationPtrOutput)
}

func (o FlowEndpointsConfigurationOutput) Connector() FlowEndpointsPtrOutput {
	return o.ApplyT(func(v FlowEndpointsConfiguration) *FlowEndpoints { return v.Connector }).(FlowEndpointsPtrOutput)
}

func (o FlowEndpointsConfigurationOutput) Workflow() FlowEndpointsPtrOutput {
	return o.ApplyT(func(v FlowEndpointsConfiguration) *FlowEndpoints { return v.Workflow }).(FlowEndpointsPtrOutput)
}

type FlowEndpointsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (FlowEndpointsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowEndpointsConfiguration)(nil)).Elem()
}

func (o FlowEndpointsConfigurationPtrOutput) ToFlowEndpointsConfigurationPtrOutput() FlowEndpointsConfigurationPtrOutput {
	return o
}

func (o FlowEndpointsConfigurationPtrOutput) ToFlowEndpointsConfigurationPtrOutputWithContext(ctx context.Context) FlowEndpointsConfigurationPtrOutput {
	return o
}

func (o FlowEndpointsConfigurationPtrOutput) Elem() FlowEndpointsConfigurationOutput {
	return o.ApplyT(func(v *FlowEndpointsConfiguration) FlowEndpointsConfiguration {
		if v != nil {
			return *v
		}
		var ret FlowEndpointsConfiguration
		return ret
	}).(FlowEndpointsConfigurationOutput)
}

func (o FlowEndpointsConfigurationPtrOutput) Connector() FlowEndpointsPtrOutput {
	return o.ApplyT(func(v *FlowEndpointsConfiguration) *FlowEndpoints {
		if v == nil {
			return nil
		}
		return v.Connector
	}).(FlowEndpointsPtrOutput)
}

func (o FlowEndpointsConfigurationPtrOutput) Workflow() FlowEndpointsPtrOutput {
	return o.ApplyT(func(v *FlowEndpointsConfiguration) *FlowEndpoints {
		if v == nil {
			return nil
		}
		return v.Workflow
	}).(FlowEndpointsPtrOutput)
}

type FlowEndpointsConfigurationResponse struct {
	Connector *FlowEndpointsResponse `pulumi:"connector"`
	Workflow  *FlowEndpointsResponse `pulumi:"workflow"`
}

type FlowEndpointsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (FlowEndpointsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowEndpointsConfigurationResponse)(nil)).Elem()
}

func (o FlowEndpointsConfigurationResponseOutput) ToFlowEndpointsConfigurationResponseOutput() FlowEndpointsConfigurationResponseOutput {
	return o
}

func (o FlowEndpointsConfigurationResponseOutput) ToFlowEndpointsConfigurationResponseOutputWithContext(ctx context.Context) FlowEndpointsConfigurationResponseOutput {
	return o
}

func (o FlowEndpointsConfigurationResponseOutput) Connector() FlowEndpointsResponsePtrOutput {
	return o.ApplyT(func(v FlowEndpointsConfigurationResponse) *FlowEndpointsResponse { return v.Connector }).(FlowEndpointsResponsePtrOutput)
}

func (o FlowEndpointsConfigurationResponseOutput) Workflow() FlowEndpointsResponsePtrOutput {
	return o.ApplyT(func(v FlowEndpointsConfigurationResponse) *FlowEndpointsResponse { return v.Workflow }).(FlowEndpointsResponsePtrOutput)
}

type FlowEndpointsConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (FlowEndpointsConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowEndpointsConfigurationResponse)(nil)).Elem()
}

func (o FlowEndpointsConfigurationResponsePtrOutput) ToFlowEndpointsConfigurationResponsePtrOutput() FlowEndpointsConfigurationResponsePtrOutput {
	return o
}

func (o FlowEndpointsConfigurationResponsePtrOutput) ToFlowEndpointsConfigurationResponsePtrOutputWithContext(ctx context.Context) FlowEndpointsConfigurationResponsePtrOutput {
	return o
}

func (o FlowEndpointsConfigurationResponsePtrOutput) Elem() FlowEndpointsConfigurationResponseOutput {
	return o.ApplyT(func(v *FlowEndpointsConfigurationResponse) FlowEndpointsConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret FlowEndpointsConfigurationResponse
		return ret
	}).(FlowEndpointsConfigurationResponseOutput)
}

func (o FlowEndpointsConfigurationResponsePtrOutput) Connector() FlowEndpointsResponsePtrOutput {
	return o.ApplyT(func(v *FlowEndpointsConfigurationResponse) *FlowEndpointsResponse {
		if v == nil {
			return nil
		}
		return v.Connector
	}).(FlowEndpointsResponsePtrOutput)
}

func (o FlowEndpointsConfigurationResponsePtrOutput) Workflow() FlowEndpointsResponsePtrOutput {
	return o.ApplyT(func(v *FlowEndpointsConfigurationResponse) *FlowEndpointsResponse {
		if v == nil {
			return nil
		}
		return v.Workflow
	}).(FlowEndpointsResponsePtrOutput)
}

type FlowEndpointsResponse struct {
	AccessEndpointIpAddresses []IpAddressResponse `pulumi:"accessEndpointIpAddresses"`
	OutgoingIpAddresses       []IpAddressResponse `pulumi:"outgoingIpAddresses"`
}

type FlowEndpointsResponseOutput struct{ *pulumi.OutputState }

func (FlowEndpointsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlowEndpointsResponse)(nil)).Elem()
}

func (o FlowEndpointsResponseOutput) ToFlowEndpointsResponseOutput() FlowEndpointsResponseOutput {
	return o
}

func (o FlowEndpointsResponseOutput) ToFlowEndpointsResponseOutputWithContext(ctx context.Context) FlowEndpointsResponseOutput {
	return o
}

func (o FlowEndpointsResponseOutput) AccessEndpointIpAddresses() IpAddressResponseArrayOutput {
	return o.ApplyT(func(v FlowEndpointsResponse) []IpAddressResponse { return v.AccessEndpointIpAddresses }).(IpAddressResponseArrayOutput)
}

func (o FlowEndpointsResponseOutput) OutgoingIpAddresses() IpAddressResponseArrayOutput {
	return o.ApplyT(func(v FlowEndpointsResponse) []IpAddressResponse { return v.OutgoingIpAddresses }).(IpAddressResponseArrayOutput)
}

type FlowEndpointsResponsePtrOutput struct{ *pulumi.OutputState }

func (FlowEndpointsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowEndpointsResponse)(nil)).Elem()
}

func (o FlowEndpointsResponsePtrOutput) ToFlowEndpointsResponsePtrOutput() FlowEndpointsResponsePtrOutput {
	return o
}

func (o FlowEndpointsResponsePtrOutput) ToFlowEndpointsResponsePtrOutputWithContext(ctx context.Context) FlowEndpointsResponsePtrOutput {
	return o
}

func (o FlowEndpointsResponsePtrOutput) Elem() FlowEndpointsResponseOutput {
	return o.ApplyT(func(v *FlowEndpointsResponse) FlowEndpointsResponse {
		if v != nil {
			return *v
		}
		var ret FlowEndpointsResponse
		return ret
	}).(FlowEndpointsResponseOutput)
}

func (o FlowEndpointsResponsePtrOutput) AccessEndpointIpAddresses() IpAddressResponseArrayOutput {
	return o.ApplyT(func(v *FlowEndpointsResponse) []IpAddressResponse {
		if v == nil {
			return nil
		}
		return v.AccessEndpointIpAddresses
	}).(IpAddressResponseArrayOutput)
}

func (o FlowEndpointsResponsePtrOutput) OutgoingIpAddresses() IpAddressResponseArrayOutput {
	return o.ApplyT(func(v *FlowEndpointsResponse) []IpAddressResponse {
		if v == nil {
			return nil
		}
		return v.OutgoingIpAddresses
	}).(IpAddressResponseArrayOutput)
}

type IntegrationAccountMapPropertiesParametersSchema struct {
	Ref *string `pulumi:"ref"`
}





type IntegrationAccountMapPropertiesParametersSchemaInput interface {
	pulumi.Input

	ToIntegrationAccountMapPropertiesParametersSchemaOutput() IntegrationAccountMapPropertiesParametersSchemaOutput
	ToIntegrationAccountMapPropertiesParametersSchemaOutputWithContext(context.Context) IntegrationAccountMapPropertiesParametersSchemaOutput
}

type IntegrationAccountMapPropertiesParametersSchemaArgs struct {
	Ref pulumi.StringPtrInput `pulumi:"ref"`
}

func (IntegrationAccountMapPropertiesParametersSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountMapPropertiesParametersSchema)(nil)).Elem()
}

func (i IntegrationAccountMapPropertiesParametersSchemaArgs) ToIntegrationAccountMapPropertiesParametersSchemaOutput() IntegrationAccountMapPropertiesParametersSchemaOutput {
	return i.ToIntegrationAccountMapPropertiesParametersSchemaOutputWithContext(context.Background())
}

func (i IntegrationAccountMapPropertiesParametersSchemaArgs) ToIntegrationAccountMapPropertiesParametersSchemaOutputWithContext(ctx context.Context) IntegrationAccountMapPropertiesParametersSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountMapPropertiesParametersSchemaOutput)
}

func (i IntegrationAccountMapPropertiesParametersSchemaArgs) ToIntegrationAccountMapPropertiesParametersSchemaPtrOutput() IntegrationAccountMapPropertiesParametersSchemaPtrOutput {
	return i.ToIntegrationAccountMapPropertiesParametersSchemaPtrOutputWithContext(context.Background())
}

func (i IntegrationAccountMapPropertiesParametersSchemaArgs) ToIntegrationAccountMapPropertiesParametersSchemaPtrOutputWithContext(ctx context.Context) IntegrationAccountMapPropertiesParametersSchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountMapPropertiesParametersSchemaOutput).ToIntegrationAccountMapPropertiesParametersSchemaPtrOutputWithContext(ctx)
}









type IntegrationAccountMapPropertiesParametersSchemaPtrInput interface {
	pulumi.Input

	ToIntegrationAccountMapPropertiesParametersSchemaPtrOutput() IntegrationAccountMapPropertiesParametersSchemaPtrOutput
	ToIntegrationAccountMapPropertiesParametersSchemaPtrOutputWithContext(context.Context) IntegrationAccountMapPropertiesParametersSchemaPtrOutput
}

type integrationAccountMapPropertiesParametersSchemaPtrType IntegrationAccountMapPropertiesParametersSchemaArgs

func IntegrationAccountMapPropertiesParametersSchemaPtr(v *IntegrationAccountMapPropertiesParametersSchemaArgs) IntegrationAccountMapPropertiesParametersSchemaPtrInput {
	return (*integrationAccountMapPropertiesParametersSchemaPtrType)(v)
}

func (*integrationAccountMapPropertiesParametersSchemaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountMapPropertiesParametersSchema)(nil)).Elem()
}

func (i *integrationAccountMapPropertiesParametersSchemaPtrType) ToIntegrationAccountMapPropertiesParametersSchemaPtrOutput() IntegrationAccountMapPropertiesParametersSchemaPtrOutput {
	return i.ToIntegrationAccountMapPropertiesParametersSchemaPtrOutputWithContext(context.Background())
}

func (i *integrationAccountMapPropertiesParametersSchemaPtrType) ToIntegrationAccountMapPropertiesParametersSchemaPtrOutputWithContext(ctx context.Context) IntegrationAccountMapPropertiesParametersSchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountMapPropertiesParametersSchemaPtrOutput)
}

type IntegrationAccountMapPropertiesParametersSchemaOutput struct{ *pulumi.OutputState }

func (IntegrationAccountMapPropertiesParametersSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountMapPropertiesParametersSchema)(nil)).Elem()
}

func (o IntegrationAccountMapPropertiesParametersSchemaOutput) ToIntegrationAccountMapPropertiesParametersSchemaOutput() IntegrationAccountMapPropertiesParametersSchemaOutput {
	return o
}

func (o IntegrationAccountMapPropertiesParametersSchemaOutput) ToIntegrationAccountMapPropertiesParametersSchemaOutputWithContext(ctx context.Context) IntegrationAccountMapPropertiesParametersSchemaOutput {
	return o
}

func (o IntegrationAccountMapPropertiesParametersSchemaOutput) ToIntegrationAccountMapPropertiesParametersSchemaPtrOutput() IntegrationAccountMapPropertiesParametersSchemaPtrOutput {
	return o.ToIntegrationAccountMapPropertiesParametersSchemaPtrOutputWithContext(context.Background())
}

func (o IntegrationAccountMapPropertiesParametersSchemaOutput) ToIntegrationAccountMapPropertiesParametersSchemaPtrOutputWithContext(ctx context.Context) IntegrationAccountMapPropertiesParametersSchemaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntegrationAccountMapPropertiesParametersSchema) *IntegrationAccountMapPropertiesParametersSchema {
		return &v
	}).(IntegrationAccountMapPropertiesParametersSchemaPtrOutput)
}

func (o IntegrationAccountMapPropertiesParametersSchemaOutput) Ref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationAccountMapPropertiesParametersSchema) *string { return v.Ref }).(pulumi.StringPtrOutput)
}

type IntegrationAccountMapPropertiesParametersSchemaPtrOutput struct{ *pulumi.OutputState }

func (IntegrationAccountMapPropertiesParametersSchemaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountMapPropertiesParametersSchema)(nil)).Elem()
}

func (o IntegrationAccountMapPropertiesParametersSchemaPtrOutput) ToIntegrationAccountMapPropertiesParametersSchemaPtrOutput() IntegrationAccountMapPropertiesParametersSchemaPtrOutput {
	return o
}

func (o IntegrationAccountMapPropertiesParametersSchemaPtrOutput) ToIntegrationAccountMapPropertiesParametersSchemaPtrOutputWithContext(ctx context.Context) IntegrationAccountMapPropertiesParametersSchemaPtrOutput {
	return o
}

func (o IntegrationAccountMapPropertiesParametersSchemaPtrOutput) Elem() IntegrationAccountMapPropertiesParametersSchemaOutput {
	return o.ApplyT(func(v *IntegrationAccountMapPropertiesParametersSchema) IntegrationAccountMapPropertiesParametersSchema {
		if v != nil {
			return *v
		}
		var ret IntegrationAccountMapPropertiesParametersSchema
		return ret
	}).(IntegrationAccountMapPropertiesParametersSchemaOutput)
}

func (o IntegrationAccountMapPropertiesParametersSchemaPtrOutput) Ref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountMapPropertiesParametersSchema) *string {
		if v == nil {
			return nil
		}
		return v.Ref
	}).(pulumi.StringPtrOutput)
}

type IntegrationAccountMapPropertiesResponseParametersSchema struct {
	Ref *string `pulumi:"ref"`
}

type IntegrationAccountMapPropertiesResponseParametersSchemaOutput struct{ *pulumi.OutputState }

func (IntegrationAccountMapPropertiesResponseParametersSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountMapPropertiesResponseParametersSchema)(nil)).Elem()
}

func (o IntegrationAccountMapPropertiesResponseParametersSchemaOutput) ToIntegrationAccountMapPropertiesResponseParametersSchemaOutput() IntegrationAccountMapPropertiesResponseParametersSchemaOutput {
	return o
}

func (o IntegrationAccountMapPropertiesResponseParametersSchemaOutput) ToIntegrationAccountMapPropertiesResponseParametersSchemaOutputWithContext(ctx context.Context) IntegrationAccountMapPropertiesResponseParametersSchemaOutput {
	return o
}

func (o IntegrationAccountMapPropertiesResponseParametersSchemaOutput) Ref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationAccountMapPropertiesResponseParametersSchema) *string { return v.Ref }).(pulumi.StringPtrOutput)
}

type IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput struct{ *pulumi.OutputState }

func (IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountMapPropertiesResponseParametersSchema)(nil)).Elem()
}

func (o IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput) ToIntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput() IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput {
	return o
}

func (o IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput) ToIntegrationAccountMapPropertiesResponseParametersSchemaPtrOutputWithContext(ctx context.Context) IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput {
	return o
}

func (o IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput) Elem() IntegrationAccountMapPropertiesResponseParametersSchemaOutput {
	return o.ApplyT(func(v *IntegrationAccountMapPropertiesResponseParametersSchema) IntegrationAccountMapPropertiesResponseParametersSchema {
		if v != nil {
			return *v
		}
		var ret IntegrationAccountMapPropertiesResponseParametersSchema
		return ret
	}).(IntegrationAccountMapPropertiesResponseParametersSchemaOutput)
}

func (o IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput) Ref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountMapPropertiesResponseParametersSchema) *string {
		if v == nil {
			return nil
		}
		return v.Ref
	}).(pulumi.StringPtrOutput)
}

type IntegrationAccountSku struct {
	Name string `pulumi:"name"`
}





type IntegrationAccountSkuInput interface {
	pulumi.Input

	ToIntegrationAccountSkuOutput() IntegrationAccountSkuOutput
	ToIntegrationAccountSkuOutputWithContext(context.Context) IntegrationAccountSkuOutput
}

type IntegrationAccountSkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (IntegrationAccountSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountSku)(nil)).Elem()
}

func (i IntegrationAccountSkuArgs) ToIntegrationAccountSkuOutput() IntegrationAccountSkuOutput {
	return i.ToIntegrationAccountSkuOutputWithContext(context.Background())
}

func (i IntegrationAccountSkuArgs) ToIntegrationAccountSkuOutputWithContext(ctx context.Context) IntegrationAccountSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountSkuOutput)
}

func (i IntegrationAccountSkuArgs) ToIntegrationAccountSkuPtrOutput() IntegrationAccountSkuPtrOutput {
	return i.ToIntegrationAccountSkuPtrOutputWithContext(context.Background())
}

func (i IntegrationAccountSkuArgs) ToIntegrationAccountSkuPtrOutputWithContext(ctx context.Context) IntegrationAccountSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountSkuOutput).ToIntegrationAccountSkuPtrOutputWithContext(ctx)
}









type IntegrationAccountSkuPtrInput interface {
	pulumi.Input

	ToIntegrationAccountSkuPtrOutput() IntegrationAccountSkuPtrOutput
	ToIntegrationAccountSkuPtrOutputWithContext(context.Context) IntegrationAccountSkuPtrOutput
}

type integrationAccountSkuPtrType IntegrationAccountSkuArgs

func IntegrationAccountSkuPtr(v *IntegrationAccountSkuArgs) IntegrationAccountSkuPtrInput {
	return (*integrationAccountSkuPtrType)(v)
}

func (*integrationAccountSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountSku)(nil)).Elem()
}

func (i *integrationAccountSkuPtrType) ToIntegrationAccountSkuPtrOutput() IntegrationAccountSkuPtrOutput {
	return i.ToIntegrationAccountSkuPtrOutputWithContext(context.Background())
}

func (i *integrationAccountSkuPtrType) ToIntegrationAccountSkuPtrOutputWithContext(ctx context.Context) IntegrationAccountSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountSkuPtrOutput)
}

type IntegrationAccountSkuOutput struct{ *pulumi.OutputState }

func (IntegrationAccountSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountSku)(nil)).Elem()
}

func (o IntegrationAccountSkuOutput) ToIntegrationAccountSkuOutput() IntegrationAccountSkuOutput {
	return o
}

func (o IntegrationAccountSkuOutput) ToIntegrationAccountSkuOutputWithContext(ctx context.Context) IntegrationAccountSkuOutput {
	return o
}

func (o IntegrationAccountSkuOutput) ToIntegrationAccountSkuPtrOutput() IntegrationAccountSkuPtrOutput {
	return o.ToIntegrationAccountSkuPtrOutputWithContext(context.Background())
}

func (o IntegrationAccountSkuOutput) ToIntegrationAccountSkuPtrOutputWithContext(ctx context.Context) IntegrationAccountSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntegrationAccountSku) *IntegrationAccountSku {
		return &v
	}).(IntegrationAccountSkuPtrOutput)
}

func (o IntegrationAccountSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v IntegrationAccountSku) string { return v.Name }).(pulumi.StringOutput)
}

type IntegrationAccountSkuPtrOutput struct{ *pulumi.OutputState }

func (IntegrationAccountSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountSku)(nil)).Elem()
}

func (o IntegrationAccountSkuPtrOutput) ToIntegrationAccountSkuPtrOutput() IntegrationAccountSkuPtrOutput {
	return o
}

func (o IntegrationAccountSkuPtrOutput) ToIntegrationAccountSkuPtrOutputWithContext(ctx context.Context) IntegrationAccountSkuPtrOutput {
	return o
}

func (o IntegrationAccountSkuPtrOutput) Elem() IntegrationAccountSkuOutput {
	return o.ApplyT(func(v *IntegrationAccountSku) IntegrationAccountSku {
		if v != nil {
			return *v
		}
		var ret IntegrationAccountSku
		return ret
	}).(IntegrationAccountSkuOutput)
}

func (o IntegrationAccountSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type IntegrationAccountSkuResponse struct {
	Name string `pulumi:"name"`
}

type IntegrationAccountSkuResponseOutput struct{ *pulumi.OutputState }

func (IntegrationAccountSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationAccountSkuResponse)(nil)).Elem()
}

func (o IntegrationAccountSkuResponseOutput) ToIntegrationAccountSkuResponseOutput() IntegrationAccountSkuResponseOutput {
	return o
}

func (o IntegrationAccountSkuResponseOutput) ToIntegrationAccountSkuResponseOutputWithContext(ctx context.Context) IntegrationAccountSkuResponseOutput {
	return o
}

func (o IntegrationAccountSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v IntegrationAccountSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type IntegrationAccountSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (IntegrationAccountSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountSkuResponse)(nil)).Elem()
}

func (o IntegrationAccountSkuResponsePtrOutput) ToIntegrationAccountSkuResponsePtrOutput() IntegrationAccountSkuResponsePtrOutput {
	return o
}

func (o IntegrationAccountSkuResponsePtrOutput) ToIntegrationAccountSkuResponsePtrOutputWithContext(ctx context.Context) IntegrationAccountSkuResponsePtrOutput {
	return o
}

func (o IntegrationAccountSkuResponsePtrOutput) Elem() IntegrationAccountSkuResponseOutput {
	return o.ApplyT(func(v *IntegrationAccountSkuResponse) IntegrationAccountSkuResponse {
		if v != nil {
			return *v
		}
		var ret IntegrationAccountSkuResponse
		return ret
	}).(IntegrationAccountSkuResponseOutput)
}

func (o IntegrationAccountSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type IntegrationServiceEnvironmenEncryptionConfiguration struct {
	EncryptionKeyReference *IntegrationServiceEnvironmenEncryptionKeyReference `pulumi:"encryptionKeyReference"`
}





type IntegrationServiceEnvironmenEncryptionConfigurationInput interface {
	pulumi.Input

	ToIntegrationServiceEnvironmenEncryptionConfigurationOutput() IntegrationServiceEnvironmenEncryptionConfigurationOutput
	ToIntegrationServiceEnvironmenEncryptionConfigurationOutputWithContext(context.Context) IntegrationServiceEnvironmenEncryptionConfigurationOutput
}

type IntegrationServiceEnvironmenEncryptionConfigurationArgs struct {
	EncryptionKeyReference IntegrationServiceEnvironmenEncryptionKeyReferencePtrInput `pulumi:"encryptionKeyReference"`
}

func (IntegrationServiceEnvironmenEncryptionConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironmenEncryptionConfiguration)(nil)).Elem()
}

func (i IntegrationServiceEnvironmenEncryptionConfigurationArgs) ToIntegrationServiceEnvironmenEncryptionConfigurationOutput() IntegrationServiceEnvironmenEncryptionConfigurationOutput {
	return i.ToIntegrationServiceEnvironmenEncryptionConfigurationOutputWithContext(context.Background())
}

func (i IntegrationServiceEnvironmenEncryptionConfigurationArgs) ToIntegrationServiceEnvironmenEncryptionConfigurationOutputWithContext(ctx context.Context) IntegrationServiceEnvironmenEncryptionConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmenEncryptionConfigurationOutput)
}

func (i IntegrationServiceEnvironmenEncryptionConfigurationArgs) ToIntegrationServiceEnvironmenEncryptionConfigurationPtrOutput() IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput {
	return i.ToIntegrationServiceEnvironmenEncryptionConfigurationPtrOutputWithContext(context.Background())
}

func (i IntegrationServiceEnvironmenEncryptionConfigurationArgs) ToIntegrationServiceEnvironmenEncryptionConfigurationPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmenEncryptionConfigurationOutput).ToIntegrationServiceEnvironmenEncryptionConfigurationPtrOutputWithContext(ctx)
}









type IntegrationServiceEnvironmenEncryptionConfigurationPtrInput interface {
	pulumi.Input

	ToIntegrationServiceEnvironmenEncryptionConfigurationPtrOutput() IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput
	ToIntegrationServiceEnvironmenEncryptionConfigurationPtrOutputWithContext(context.Context) IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput
}

type integrationServiceEnvironmenEncryptionConfigurationPtrType IntegrationServiceEnvironmenEncryptionConfigurationArgs

func IntegrationServiceEnvironmenEncryptionConfigurationPtr(v *IntegrationServiceEnvironmenEncryptionConfigurationArgs) IntegrationServiceEnvironmenEncryptionConfigurationPtrInput {
	return (*integrationServiceEnvironmenEncryptionConfigurationPtrType)(v)
}

func (*integrationServiceEnvironmenEncryptionConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmenEncryptionConfiguration)(nil)).Elem()
}

func (i *integrationServiceEnvironmenEncryptionConfigurationPtrType) ToIntegrationServiceEnvironmenEncryptionConfigurationPtrOutput() IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput {
	return i.ToIntegrationServiceEnvironmenEncryptionConfigurationPtrOutputWithContext(context.Background())
}

func (i *integrationServiceEnvironmenEncryptionConfigurationPtrType) ToIntegrationServiceEnvironmenEncryptionConfigurationPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput)
}

type IntegrationServiceEnvironmenEncryptionConfigurationOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmenEncryptionConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironmenEncryptionConfiguration)(nil)).Elem()
}

func (o IntegrationServiceEnvironmenEncryptionConfigurationOutput) ToIntegrationServiceEnvironmenEncryptionConfigurationOutput() IntegrationServiceEnvironmenEncryptionConfigurationOutput {
	return o
}

func (o IntegrationServiceEnvironmenEncryptionConfigurationOutput) ToIntegrationServiceEnvironmenEncryptionConfigurationOutputWithContext(ctx context.Context) IntegrationServiceEnvironmenEncryptionConfigurationOutput {
	return o
}

func (o IntegrationServiceEnvironmenEncryptionConfigurationOutput) ToIntegrationServiceEnvironmenEncryptionConfigurationPtrOutput() IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput {
	return o.ToIntegrationServiceEnvironmenEncryptionConfigurationPtrOutputWithContext(context.Background())
}

func (o IntegrationServiceEnvironmenEncryptionConfigurationOutput) ToIntegrationServiceEnvironmenEncryptionConfigurationPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntegrationServiceEnvironmenEncryptionConfiguration) *IntegrationServiceEnvironmenEncryptionConfiguration {
		return &v
	}).(IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput)
}

func (o IntegrationServiceEnvironmenEncryptionConfigurationOutput) EncryptionKeyReference() IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmenEncryptionConfiguration) *IntegrationServiceEnvironmenEncryptionKeyReference {
		return v.EncryptionKeyReference
	}).(IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput)
}

type IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmenEncryptionConfiguration)(nil)).Elem()
}

func (o IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput) ToIntegrationServiceEnvironmenEncryptionConfigurationPtrOutput() IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput {
	return o
}

func (o IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput) ToIntegrationServiceEnvironmenEncryptionConfigurationPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput {
	return o
}

func (o IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput) Elem() IntegrationServiceEnvironmenEncryptionConfigurationOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmenEncryptionConfiguration) IntegrationServiceEnvironmenEncryptionConfiguration {
		if v != nil {
			return *v
		}
		var ret IntegrationServiceEnvironmenEncryptionConfiguration
		return ret
	}).(IntegrationServiceEnvironmenEncryptionConfigurationOutput)
}

func (o IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput) EncryptionKeyReference() IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmenEncryptionConfiguration) *IntegrationServiceEnvironmenEncryptionKeyReference {
		if v == nil {
			return nil
		}
		return v.EncryptionKeyReference
	}).(IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput)
}

type IntegrationServiceEnvironmenEncryptionConfigurationResponse struct {
	EncryptionKeyReference *IntegrationServiceEnvironmenEncryptionKeyReferenceResponse `pulumi:"encryptionKeyReference"`
}

type IntegrationServiceEnvironmenEncryptionConfigurationResponseOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmenEncryptionConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironmenEncryptionConfigurationResponse)(nil)).Elem()
}

func (o IntegrationServiceEnvironmenEncryptionConfigurationResponseOutput) ToIntegrationServiceEnvironmenEncryptionConfigurationResponseOutput() IntegrationServiceEnvironmenEncryptionConfigurationResponseOutput {
	return o
}

func (o IntegrationServiceEnvironmenEncryptionConfigurationResponseOutput) ToIntegrationServiceEnvironmenEncryptionConfigurationResponseOutputWithContext(ctx context.Context) IntegrationServiceEnvironmenEncryptionConfigurationResponseOutput {
	return o
}

func (o IntegrationServiceEnvironmenEncryptionConfigurationResponseOutput) EncryptionKeyReference() IntegrationServiceEnvironmenEncryptionKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmenEncryptionConfigurationResponse) *IntegrationServiceEnvironmenEncryptionKeyReferenceResponse {
		return v.EncryptionKeyReference
	}).(IntegrationServiceEnvironmenEncryptionKeyReferenceResponsePtrOutput)
}

type IntegrationServiceEnvironmenEncryptionConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmenEncryptionConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmenEncryptionConfigurationResponse)(nil)).Elem()
}

func (o IntegrationServiceEnvironmenEncryptionConfigurationResponsePtrOutput) ToIntegrationServiceEnvironmenEncryptionConfigurationResponsePtrOutput() IntegrationServiceEnvironmenEncryptionConfigurationResponsePtrOutput {
	return o
}

func (o IntegrationServiceEnvironmenEncryptionConfigurationResponsePtrOutput) ToIntegrationServiceEnvironmenEncryptionConfigurationResponsePtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmenEncryptionConfigurationResponsePtrOutput {
	return o
}

func (o IntegrationServiceEnvironmenEncryptionConfigurationResponsePtrOutput) Elem() IntegrationServiceEnvironmenEncryptionConfigurationResponseOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmenEncryptionConfigurationResponse) IntegrationServiceEnvironmenEncryptionConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret IntegrationServiceEnvironmenEncryptionConfigurationResponse
		return ret
	}).(IntegrationServiceEnvironmenEncryptionConfigurationResponseOutput)
}

func (o IntegrationServiceEnvironmenEncryptionConfigurationResponsePtrOutput) EncryptionKeyReference() IntegrationServiceEnvironmenEncryptionKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmenEncryptionConfigurationResponse) *IntegrationServiceEnvironmenEncryptionKeyReferenceResponse {
		if v == nil {
			return nil
		}
		return v.EncryptionKeyReference
	}).(IntegrationServiceEnvironmenEncryptionKeyReferenceResponsePtrOutput)
}

type IntegrationServiceEnvironmenEncryptionKeyReference struct {
	KeyName    *string            `pulumi:"keyName"`
	KeyVault   *ResourceReference `pulumi:"keyVault"`
	KeyVersion *string            `pulumi:"keyVersion"`
}





type IntegrationServiceEnvironmenEncryptionKeyReferenceInput interface {
	pulumi.Input

	ToIntegrationServiceEnvironmenEncryptionKeyReferenceOutput() IntegrationServiceEnvironmenEncryptionKeyReferenceOutput
	ToIntegrationServiceEnvironmenEncryptionKeyReferenceOutputWithContext(context.Context) IntegrationServiceEnvironmenEncryptionKeyReferenceOutput
}

type IntegrationServiceEnvironmenEncryptionKeyReferenceArgs struct {
	KeyName    pulumi.StringPtrInput     `pulumi:"keyName"`
	KeyVault   ResourceReferencePtrInput `pulumi:"keyVault"`
	KeyVersion pulumi.StringPtrInput     `pulumi:"keyVersion"`
}

func (IntegrationServiceEnvironmenEncryptionKeyReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironmenEncryptionKeyReference)(nil)).Elem()
}

func (i IntegrationServiceEnvironmenEncryptionKeyReferenceArgs) ToIntegrationServiceEnvironmenEncryptionKeyReferenceOutput() IntegrationServiceEnvironmenEncryptionKeyReferenceOutput {
	return i.ToIntegrationServiceEnvironmenEncryptionKeyReferenceOutputWithContext(context.Background())
}

func (i IntegrationServiceEnvironmenEncryptionKeyReferenceArgs) ToIntegrationServiceEnvironmenEncryptionKeyReferenceOutputWithContext(ctx context.Context) IntegrationServiceEnvironmenEncryptionKeyReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmenEncryptionKeyReferenceOutput)
}

func (i IntegrationServiceEnvironmenEncryptionKeyReferenceArgs) ToIntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput() IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput {
	return i.ToIntegrationServiceEnvironmenEncryptionKeyReferencePtrOutputWithContext(context.Background())
}

func (i IntegrationServiceEnvironmenEncryptionKeyReferenceArgs) ToIntegrationServiceEnvironmenEncryptionKeyReferencePtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmenEncryptionKeyReferenceOutput).ToIntegrationServiceEnvironmenEncryptionKeyReferencePtrOutputWithContext(ctx)
}









type IntegrationServiceEnvironmenEncryptionKeyReferencePtrInput interface {
	pulumi.Input

	ToIntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput() IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput
	ToIntegrationServiceEnvironmenEncryptionKeyReferencePtrOutputWithContext(context.Context) IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput
}

type integrationServiceEnvironmenEncryptionKeyReferencePtrType IntegrationServiceEnvironmenEncryptionKeyReferenceArgs

func IntegrationServiceEnvironmenEncryptionKeyReferencePtr(v *IntegrationServiceEnvironmenEncryptionKeyReferenceArgs) IntegrationServiceEnvironmenEncryptionKeyReferencePtrInput {
	return (*integrationServiceEnvironmenEncryptionKeyReferencePtrType)(v)
}

func (*integrationServiceEnvironmenEncryptionKeyReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmenEncryptionKeyReference)(nil)).Elem()
}

func (i *integrationServiceEnvironmenEncryptionKeyReferencePtrType) ToIntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput() IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput {
	return i.ToIntegrationServiceEnvironmenEncryptionKeyReferencePtrOutputWithContext(context.Background())
}

func (i *integrationServiceEnvironmenEncryptionKeyReferencePtrType) ToIntegrationServiceEnvironmenEncryptionKeyReferencePtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput)
}

type IntegrationServiceEnvironmenEncryptionKeyReferenceOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmenEncryptionKeyReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironmenEncryptionKeyReference)(nil)).Elem()
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferenceOutput) ToIntegrationServiceEnvironmenEncryptionKeyReferenceOutput() IntegrationServiceEnvironmenEncryptionKeyReferenceOutput {
	return o
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferenceOutput) ToIntegrationServiceEnvironmenEncryptionKeyReferenceOutputWithContext(ctx context.Context) IntegrationServiceEnvironmenEncryptionKeyReferenceOutput {
	return o
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferenceOutput) ToIntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput() IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput {
	return o.ToIntegrationServiceEnvironmenEncryptionKeyReferencePtrOutputWithContext(context.Background())
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferenceOutput) ToIntegrationServiceEnvironmenEncryptionKeyReferencePtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntegrationServiceEnvironmenEncryptionKeyReference) *IntegrationServiceEnvironmenEncryptionKeyReference {
		return &v
	}).(IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput)
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferenceOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmenEncryptionKeyReference) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferenceOutput) KeyVault() ResourceReferencePtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmenEncryptionKeyReference) *ResourceReference { return v.KeyVault }).(ResourceReferencePtrOutput)
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferenceOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmenEncryptionKeyReference) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmenEncryptionKeyReference)(nil)).Elem()
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput) ToIntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput() IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput {
	return o
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput) ToIntegrationServiceEnvironmenEncryptionKeyReferencePtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput {
	return o
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput) Elem() IntegrationServiceEnvironmenEncryptionKeyReferenceOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmenEncryptionKeyReference) IntegrationServiceEnvironmenEncryptionKeyReference {
		if v != nil {
			return *v
		}
		var ret IntegrationServiceEnvironmenEncryptionKeyReference
		return ret
	}).(IntegrationServiceEnvironmenEncryptionKeyReferenceOutput)
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmenEncryptionKeyReference) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput) KeyVault() ResourceReferencePtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmenEncryptionKeyReference) *ResourceReference {
		if v == nil {
			return nil
		}
		return v.KeyVault
	}).(ResourceReferencePtrOutput)
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmenEncryptionKeyReference) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type IntegrationServiceEnvironmenEncryptionKeyReferenceResponse struct {
	KeyName    *string                    `pulumi:"keyName"`
	KeyVault   *ResourceReferenceResponse `pulumi:"keyVault"`
	KeyVersion *string                    `pulumi:"keyVersion"`
}

type IntegrationServiceEnvironmenEncryptionKeyReferenceResponseOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmenEncryptionKeyReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironmenEncryptionKeyReferenceResponse)(nil)).Elem()
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferenceResponseOutput) ToIntegrationServiceEnvironmenEncryptionKeyReferenceResponseOutput() IntegrationServiceEnvironmenEncryptionKeyReferenceResponseOutput {
	return o
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferenceResponseOutput) ToIntegrationServiceEnvironmenEncryptionKeyReferenceResponseOutputWithContext(ctx context.Context) IntegrationServiceEnvironmenEncryptionKeyReferenceResponseOutput {
	return o
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferenceResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmenEncryptionKeyReferenceResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferenceResponseOutput) KeyVault() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmenEncryptionKeyReferenceResponse) *ResourceReferenceResponse {
		return v.KeyVault
	}).(ResourceReferenceResponsePtrOutput)
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferenceResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmenEncryptionKeyReferenceResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type IntegrationServiceEnvironmenEncryptionKeyReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmenEncryptionKeyReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmenEncryptionKeyReferenceResponse)(nil)).Elem()
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferenceResponsePtrOutput) ToIntegrationServiceEnvironmenEncryptionKeyReferenceResponsePtrOutput() IntegrationServiceEnvironmenEncryptionKeyReferenceResponsePtrOutput {
	return o
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferenceResponsePtrOutput) ToIntegrationServiceEnvironmenEncryptionKeyReferenceResponsePtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmenEncryptionKeyReferenceResponsePtrOutput {
	return o
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferenceResponsePtrOutput) Elem() IntegrationServiceEnvironmenEncryptionKeyReferenceResponseOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmenEncryptionKeyReferenceResponse) IntegrationServiceEnvironmenEncryptionKeyReferenceResponse {
		if v != nil {
			return *v
		}
		var ret IntegrationServiceEnvironmenEncryptionKeyReferenceResponse
		return ret
	}).(IntegrationServiceEnvironmenEncryptionKeyReferenceResponseOutput)
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferenceResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmenEncryptionKeyReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferenceResponsePtrOutput) KeyVault() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmenEncryptionKeyReferenceResponse) *ResourceReferenceResponse {
		if v == nil {
			return nil
		}
		return v.KeyVault
	}).(ResourceReferenceResponsePtrOutput)
}

func (o IntegrationServiceEnvironmenEncryptionKeyReferenceResponsePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmenEncryptionKeyReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type IntegrationServiceEnvironmentAccessEndpoint struct {
	Type *string `pulumi:"type"`
}





type IntegrationServiceEnvironmentAccessEndpointInput interface {
	pulumi.Input

	ToIntegrationServiceEnvironmentAccessEndpointOutput() IntegrationServiceEnvironmentAccessEndpointOutput
	ToIntegrationServiceEnvironmentAccessEndpointOutputWithContext(context.Context) IntegrationServiceEnvironmentAccessEndpointOutput
}

type IntegrationServiceEnvironmentAccessEndpointArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (IntegrationServiceEnvironmentAccessEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironmentAccessEndpoint)(nil)).Elem()
}

func (i IntegrationServiceEnvironmentAccessEndpointArgs) ToIntegrationServiceEnvironmentAccessEndpointOutput() IntegrationServiceEnvironmentAccessEndpointOutput {
	return i.ToIntegrationServiceEnvironmentAccessEndpointOutputWithContext(context.Background())
}

func (i IntegrationServiceEnvironmentAccessEndpointArgs) ToIntegrationServiceEnvironmentAccessEndpointOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentAccessEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmentAccessEndpointOutput)
}

func (i IntegrationServiceEnvironmentAccessEndpointArgs) ToIntegrationServiceEnvironmentAccessEndpointPtrOutput() IntegrationServiceEnvironmentAccessEndpointPtrOutput {
	return i.ToIntegrationServiceEnvironmentAccessEndpointPtrOutputWithContext(context.Background())
}

func (i IntegrationServiceEnvironmentAccessEndpointArgs) ToIntegrationServiceEnvironmentAccessEndpointPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentAccessEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmentAccessEndpointOutput).ToIntegrationServiceEnvironmentAccessEndpointPtrOutputWithContext(ctx)
}









type IntegrationServiceEnvironmentAccessEndpointPtrInput interface {
	pulumi.Input

	ToIntegrationServiceEnvironmentAccessEndpointPtrOutput() IntegrationServiceEnvironmentAccessEndpointPtrOutput
	ToIntegrationServiceEnvironmentAccessEndpointPtrOutputWithContext(context.Context) IntegrationServiceEnvironmentAccessEndpointPtrOutput
}

type integrationServiceEnvironmentAccessEndpointPtrType IntegrationServiceEnvironmentAccessEndpointArgs

func IntegrationServiceEnvironmentAccessEndpointPtr(v *IntegrationServiceEnvironmentAccessEndpointArgs) IntegrationServiceEnvironmentAccessEndpointPtrInput {
	return (*integrationServiceEnvironmentAccessEndpointPtrType)(v)
}

func (*integrationServiceEnvironmentAccessEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmentAccessEndpoint)(nil)).Elem()
}

func (i *integrationServiceEnvironmentAccessEndpointPtrType) ToIntegrationServiceEnvironmentAccessEndpointPtrOutput() IntegrationServiceEnvironmentAccessEndpointPtrOutput {
	return i.ToIntegrationServiceEnvironmentAccessEndpointPtrOutputWithContext(context.Background())
}

func (i *integrationServiceEnvironmentAccessEndpointPtrType) ToIntegrationServiceEnvironmentAccessEndpointPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentAccessEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmentAccessEndpointPtrOutput)
}

type IntegrationServiceEnvironmentAccessEndpointOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentAccessEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironmentAccessEndpoint)(nil)).Elem()
}

func (o IntegrationServiceEnvironmentAccessEndpointOutput) ToIntegrationServiceEnvironmentAccessEndpointOutput() IntegrationServiceEnvironmentAccessEndpointOutput {
	return o
}

func (o IntegrationServiceEnvironmentAccessEndpointOutput) ToIntegrationServiceEnvironmentAccessEndpointOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentAccessEndpointOutput {
	return o
}

func (o IntegrationServiceEnvironmentAccessEndpointOutput) ToIntegrationServiceEnvironmentAccessEndpointPtrOutput() IntegrationServiceEnvironmentAccessEndpointPtrOutput {
	return o.ToIntegrationServiceEnvironmentAccessEndpointPtrOutputWithContext(context.Background())
}

func (o IntegrationServiceEnvironmentAccessEndpointOutput) ToIntegrationServiceEnvironmentAccessEndpointPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentAccessEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntegrationServiceEnvironmentAccessEndpoint) *IntegrationServiceEnvironmentAccessEndpoint {
		return &v
	}).(IntegrationServiceEnvironmentAccessEndpointPtrOutput)
}

func (o IntegrationServiceEnvironmentAccessEndpointOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentAccessEndpoint) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IntegrationServiceEnvironmentAccessEndpointPtrOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentAccessEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmentAccessEndpoint)(nil)).Elem()
}

func (o IntegrationServiceEnvironmentAccessEndpointPtrOutput) ToIntegrationServiceEnvironmentAccessEndpointPtrOutput() IntegrationServiceEnvironmentAccessEndpointPtrOutput {
	return o
}

func (o IntegrationServiceEnvironmentAccessEndpointPtrOutput) ToIntegrationServiceEnvironmentAccessEndpointPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentAccessEndpointPtrOutput {
	return o
}

func (o IntegrationServiceEnvironmentAccessEndpointPtrOutput) Elem() IntegrationServiceEnvironmentAccessEndpointOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentAccessEndpoint) IntegrationServiceEnvironmentAccessEndpoint {
		if v != nil {
			return *v
		}
		var ret IntegrationServiceEnvironmentAccessEndpoint
		return ret
	}).(IntegrationServiceEnvironmentAccessEndpointOutput)
}

func (o IntegrationServiceEnvironmentAccessEndpointPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentAccessEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type IntegrationServiceEnvironmentAccessEndpointResponse struct {
	Type *string `pulumi:"type"`
}

type IntegrationServiceEnvironmentAccessEndpointResponseOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentAccessEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironmentAccessEndpointResponse)(nil)).Elem()
}

func (o IntegrationServiceEnvironmentAccessEndpointResponseOutput) ToIntegrationServiceEnvironmentAccessEndpointResponseOutput() IntegrationServiceEnvironmentAccessEndpointResponseOutput {
	return o
}

func (o IntegrationServiceEnvironmentAccessEndpointResponseOutput) ToIntegrationServiceEnvironmentAccessEndpointResponseOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentAccessEndpointResponseOutput {
	return o
}

func (o IntegrationServiceEnvironmentAccessEndpointResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentAccessEndpointResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IntegrationServiceEnvironmentAccessEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentAccessEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmentAccessEndpointResponse)(nil)).Elem()
}

func (o IntegrationServiceEnvironmentAccessEndpointResponsePtrOutput) ToIntegrationServiceEnvironmentAccessEndpointResponsePtrOutput() IntegrationServiceEnvironmentAccessEndpointResponsePtrOutput {
	return o
}

func (o IntegrationServiceEnvironmentAccessEndpointResponsePtrOutput) ToIntegrationServiceEnvironmentAccessEndpointResponsePtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentAccessEndpointResponsePtrOutput {
	return o
}

func (o IntegrationServiceEnvironmentAccessEndpointResponsePtrOutput) Elem() IntegrationServiceEnvironmentAccessEndpointResponseOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentAccessEndpointResponse) IntegrationServiceEnvironmentAccessEndpointResponse {
		if v != nil {
			return *v
		}
		var ret IntegrationServiceEnvironmentAccessEndpointResponse
		return ret
	}).(IntegrationServiceEnvironmentAccessEndpointResponseOutput)
}

func (o IntegrationServiceEnvironmentAccessEndpointResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentAccessEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type IntegrationServiceEnvironmentManagedApiDeploymentParameters struct {
	ContentLinkDefinition *ContentLink `pulumi:"contentLinkDefinition"`
}





type IntegrationServiceEnvironmentManagedApiDeploymentParametersInput interface {
	pulumi.Input

	ToIntegrationServiceEnvironmentManagedApiDeploymentParametersOutput() IntegrationServiceEnvironmentManagedApiDeploymentParametersOutput
	ToIntegrationServiceEnvironmentManagedApiDeploymentParametersOutputWithContext(context.Context) IntegrationServiceEnvironmentManagedApiDeploymentParametersOutput
}

type IntegrationServiceEnvironmentManagedApiDeploymentParametersArgs struct {
	ContentLinkDefinition ContentLinkPtrInput `pulumi:"contentLinkDefinition"`
}

func (IntegrationServiceEnvironmentManagedApiDeploymentParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironmentManagedApiDeploymentParameters)(nil)).Elem()
}

func (i IntegrationServiceEnvironmentManagedApiDeploymentParametersArgs) ToIntegrationServiceEnvironmentManagedApiDeploymentParametersOutput() IntegrationServiceEnvironmentManagedApiDeploymentParametersOutput {
	return i.ToIntegrationServiceEnvironmentManagedApiDeploymentParametersOutputWithContext(context.Background())
}

func (i IntegrationServiceEnvironmentManagedApiDeploymentParametersArgs) ToIntegrationServiceEnvironmentManagedApiDeploymentParametersOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentManagedApiDeploymentParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmentManagedApiDeploymentParametersOutput)
}

func (i IntegrationServiceEnvironmentManagedApiDeploymentParametersArgs) ToIntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput() IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput {
	return i.ToIntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutputWithContext(context.Background())
}

func (i IntegrationServiceEnvironmentManagedApiDeploymentParametersArgs) ToIntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmentManagedApiDeploymentParametersOutput).ToIntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutputWithContext(ctx)
}









type IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrInput interface {
	pulumi.Input

	ToIntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput() IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput
	ToIntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutputWithContext(context.Context) IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput
}

type integrationServiceEnvironmentManagedApiDeploymentParametersPtrType IntegrationServiceEnvironmentManagedApiDeploymentParametersArgs

func IntegrationServiceEnvironmentManagedApiDeploymentParametersPtr(v *IntegrationServiceEnvironmentManagedApiDeploymentParametersArgs) IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrInput {
	return (*integrationServiceEnvironmentManagedApiDeploymentParametersPtrType)(v)
}

func (*integrationServiceEnvironmentManagedApiDeploymentParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmentManagedApiDeploymentParameters)(nil)).Elem()
}

func (i *integrationServiceEnvironmentManagedApiDeploymentParametersPtrType) ToIntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput() IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput {
	return i.ToIntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutputWithContext(context.Background())
}

func (i *integrationServiceEnvironmentManagedApiDeploymentParametersPtrType) ToIntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput)
}

type IntegrationServiceEnvironmentManagedApiDeploymentParametersOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentManagedApiDeploymentParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironmentManagedApiDeploymentParameters)(nil)).Elem()
}

func (o IntegrationServiceEnvironmentManagedApiDeploymentParametersOutput) ToIntegrationServiceEnvironmentManagedApiDeploymentParametersOutput() IntegrationServiceEnvironmentManagedApiDeploymentParametersOutput {
	return o
}

func (o IntegrationServiceEnvironmentManagedApiDeploymentParametersOutput) ToIntegrationServiceEnvironmentManagedApiDeploymentParametersOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentManagedApiDeploymentParametersOutput {
	return o
}

func (o IntegrationServiceEnvironmentManagedApiDeploymentParametersOutput) ToIntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput() IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput {
	return o.ToIntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutputWithContext(context.Background())
}

func (o IntegrationServiceEnvironmentManagedApiDeploymentParametersOutput) ToIntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntegrationServiceEnvironmentManagedApiDeploymentParameters) *IntegrationServiceEnvironmentManagedApiDeploymentParameters {
		return &v
	}).(IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput)
}

func (o IntegrationServiceEnvironmentManagedApiDeploymentParametersOutput) ContentLinkDefinition() ContentLinkPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentManagedApiDeploymentParameters) *ContentLink {
		return v.ContentLinkDefinition
	}).(ContentLinkPtrOutput)
}

type IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmentManagedApiDeploymentParameters)(nil)).Elem()
}

func (o IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput) ToIntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput() IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput {
	return o
}

func (o IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput) ToIntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput {
	return o
}

func (o IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput) Elem() IntegrationServiceEnvironmentManagedApiDeploymentParametersOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentManagedApiDeploymentParameters) IntegrationServiceEnvironmentManagedApiDeploymentParameters {
		if v != nil {
			return *v
		}
		var ret IntegrationServiceEnvironmentManagedApiDeploymentParameters
		return ret
	}).(IntegrationServiceEnvironmentManagedApiDeploymentParametersOutput)
}

func (o IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput) ContentLinkDefinition() ContentLinkPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentManagedApiDeploymentParameters) *ContentLink {
		if v == nil {
			return nil
		}
		return v.ContentLinkDefinition
	}).(ContentLinkPtrOutput)
}

type IntegrationServiceEnvironmentManagedApiDeploymentParametersResponse struct {
	ContentLinkDefinition *ContentLinkResponse `pulumi:"contentLinkDefinition"`
}

type IntegrationServiceEnvironmentManagedApiDeploymentParametersResponseOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentManagedApiDeploymentParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironmentManagedApiDeploymentParametersResponse)(nil)).Elem()
}

func (o IntegrationServiceEnvironmentManagedApiDeploymentParametersResponseOutput) ToIntegrationServiceEnvironmentManagedApiDeploymentParametersResponseOutput() IntegrationServiceEnvironmentManagedApiDeploymentParametersResponseOutput {
	return o
}

func (o IntegrationServiceEnvironmentManagedApiDeploymentParametersResponseOutput) ToIntegrationServiceEnvironmentManagedApiDeploymentParametersResponseOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentManagedApiDeploymentParametersResponseOutput {
	return o
}

func (o IntegrationServiceEnvironmentManagedApiDeploymentParametersResponseOutput) ContentLinkDefinition() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentManagedApiDeploymentParametersResponse) *ContentLinkResponse {
		return v.ContentLinkDefinition
	}).(ContentLinkResponsePtrOutput)
}

type IntegrationServiceEnvironmentManagedApiDeploymentParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentManagedApiDeploymentParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmentManagedApiDeploymentParametersResponse)(nil)).Elem()
}

func (o IntegrationServiceEnvironmentManagedApiDeploymentParametersResponsePtrOutput) ToIntegrationServiceEnvironmentManagedApiDeploymentParametersResponsePtrOutput() IntegrationServiceEnvironmentManagedApiDeploymentParametersResponsePtrOutput {
	return o
}

func (o IntegrationServiceEnvironmentManagedApiDeploymentParametersResponsePtrOutput) ToIntegrationServiceEnvironmentManagedApiDeploymentParametersResponsePtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentManagedApiDeploymentParametersResponsePtrOutput {
	return o
}

func (o IntegrationServiceEnvironmentManagedApiDeploymentParametersResponsePtrOutput) Elem() IntegrationServiceEnvironmentManagedApiDeploymentParametersResponseOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentManagedApiDeploymentParametersResponse) IntegrationServiceEnvironmentManagedApiDeploymentParametersResponse {
		if v != nil {
			return *v
		}
		var ret IntegrationServiceEnvironmentManagedApiDeploymentParametersResponse
		return ret
	}).(IntegrationServiceEnvironmentManagedApiDeploymentParametersResponseOutput)
}

func (o IntegrationServiceEnvironmentManagedApiDeploymentParametersResponsePtrOutput) ContentLinkDefinition() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentManagedApiDeploymentParametersResponse) *ContentLinkResponse {
		if v == nil {
			return nil
		}
		return v.ContentLinkDefinition
	}).(ContentLinkResponsePtrOutput)
}

type IntegrationServiceEnvironmentProperties struct {
	EncryptionConfiguration         *IntegrationServiceEnvironmenEncryptionConfiguration `pulumi:"encryptionConfiguration"`
	EndpointsConfiguration          *FlowEndpointsConfiguration                          `pulumi:"endpointsConfiguration"`
	IntegrationServiceEnvironmentId *string                                              `pulumi:"integrationServiceEnvironmentId"`
	NetworkConfiguration            *NetworkConfiguration                                `pulumi:"networkConfiguration"`
	ProvisioningState               *string                                              `pulumi:"provisioningState"`
	State                           *string                                              `pulumi:"state"`
}





type IntegrationServiceEnvironmentPropertiesInput interface {
	pulumi.Input

	ToIntegrationServiceEnvironmentPropertiesOutput() IntegrationServiceEnvironmentPropertiesOutput
	ToIntegrationServiceEnvironmentPropertiesOutputWithContext(context.Context) IntegrationServiceEnvironmentPropertiesOutput
}

type IntegrationServiceEnvironmentPropertiesArgs struct {
	EncryptionConfiguration         IntegrationServiceEnvironmenEncryptionConfigurationPtrInput `pulumi:"encryptionConfiguration"`
	EndpointsConfiguration          FlowEndpointsConfigurationPtrInput                          `pulumi:"endpointsConfiguration"`
	IntegrationServiceEnvironmentId pulumi.StringPtrInput                                       `pulumi:"integrationServiceEnvironmentId"`
	NetworkConfiguration            NetworkConfigurationPtrInput                                `pulumi:"networkConfiguration"`
	ProvisioningState               pulumi.StringPtrInput                                       `pulumi:"provisioningState"`
	State                           pulumi.StringPtrInput                                       `pulumi:"state"`
}

func (IntegrationServiceEnvironmentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironmentProperties)(nil)).Elem()
}

func (i IntegrationServiceEnvironmentPropertiesArgs) ToIntegrationServiceEnvironmentPropertiesOutput() IntegrationServiceEnvironmentPropertiesOutput {
	return i.ToIntegrationServiceEnvironmentPropertiesOutputWithContext(context.Background())
}

func (i IntegrationServiceEnvironmentPropertiesArgs) ToIntegrationServiceEnvironmentPropertiesOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmentPropertiesOutput)
}

func (i IntegrationServiceEnvironmentPropertiesArgs) ToIntegrationServiceEnvironmentPropertiesPtrOutput() IntegrationServiceEnvironmentPropertiesPtrOutput {
	return i.ToIntegrationServiceEnvironmentPropertiesPtrOutputWithContext(context.Background())
}

func (i IntegrationServiceEnvironmentPropertiesArgs) ToIntegrationServiceEnvironmentPropertiesPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmentPropertiesOutput).ToIntegrationServiceEnvironmentPropertiesPtrOutputWithContext(ctx)
}









type IntegrationServiceEnvironmentPropertiesPtrInput interface {
	pulumi.Input

	ToIntegrationServiceEnvironmentPropertiesPtrOutput() IntegrationServiceEnvironmentPropertiesPtrOutput
	ToIntegrationServiceEnvironmentPropertiesPtrOutputWithContext(context.Context) IntegrationServiceEnvironmentPropertiesPtrOutput
}

type integrationServiceEnvironmentPropertiesPtrType IntegrationServiceEnvironmentPropertiesArgs

func IntegrationServiceEnvironmentPropertiesPtr(v *IntegrationServiceEnvironmentPropertiesArgs) IntegrationServiceEnvironmentPropertiesPtrInput {
	return (*integrationServiceEnvironmentPropertiesPtrType)(v)
}

func (*integrationServiceEnvironmentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmentProperties)(nil)).Elem()
}

func (i *integrationServiceEnvironmentPropertiesPtrType) ToIntegrationServiceEnvironmentPropertiesPtrOutput() IntegrationServiceEnvironmentPropertiesPtrOutput {
	return i.ToIntegrationServiceEnvironmentPropertiesPtrOutputWithContext(context.Background())
}

func (i *integrationServiceEnvironmentPropertiesPtrType) ToIntegrationServiceEnvironmentPropertiesPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmentPropertiesPtrOutput)
}

type IntegrationServiceEnvironmentPropertiesOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironmentProperties)(nil)).Elem()
}

func (o IntegrationServiceEnvironmentPropertiesOutput) ToIntegrationServiceEnvironmentPropertiesOutput() IntegrationServiceEnvironmentPropertiesOutput {
	return o
}

func (o IntegrationServiceEnvironmentPropertiesOutput) ToIntegrationServiceEnvironmentPropertiesOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentPropertiesOutput {
	return o
}

func (o IntegrationServiceEnvironmentPropertiesOutput) ToIntegrationServiceEnvironmentPropertiesPtrOutput() IntegrationServiceEnvironmentPropertiesPtrOutput {
	return o.ToIntegrationServiceEnvironmentPropertiesPtrOutputWithContext(context.Background())
}

func (o IntegrationServiceEnvironmentPropertiesOutput) ToIntegrationServiceEnvironmentPropertiesPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntegrationServiceEnvironmentProperties) *IntegrationServiceEnvironmentProperties {
		return &v
	}).(IntegrationServiceEnvironmentPropertiesPtrOutput)
}

func (o IntegrationServiceEnvironmentPropertiesOutput) EncryptionConfiguration() IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentProperties) *IntegrationServiceEnvironmenEncryptionConfiguration {
		return v.EncryptionConfiguration
	}).(IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput)
}

func (o IntegrationServiceEnvironmentPropertiesOutput) EndpointsConfiguration() FlowEndpointsConfigurationPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentProperties) *FlowEndpointsConfiguration {
		return v.EndpointsConfiguration
	}).(FlowEndpointsConfigurationPtrOutput)
}

func (o IntegrationServiceEnvironmentPropertiesOutput) IntegrationServiceEnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentProperties) *string { return v.IntegrationServiceEnvironmentId }).(pulumi.StringPtrOutput)
}

func (o IntegrationServiceEnvironmentPropertiesOutput) NetworkConfiguration() NetworkConfigurationPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentProperties) *NetworkConfiguration { return v.NetworkConfiguration }).(NetworkConfigurationPtrOutput)
}

func (o IntegrationServiceEnvironmentPropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o IntegrationServiceEnvironmentPropertiesOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentProperties) *string { return v.State }).(pulumi.StringPtrOutput)
}

type IntegrationServiceEnvironmentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmentProperties)(nil)).Elem()
}

func (o IntegrationServiceEnvironmentPropertiesPtrOutput) ToIntegrationServiceEnvironmentPropertiesPtrOutput() IntegrationServiceEnvironmentPropertiesPtrOutput {
	return o
}

func (o IntegrationServiceEnvironmentPropertiesPtrOutput) ToIntegrationServiceEnvironmentPropertiesPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentPropertiesPtrOutput {
	return o
}

func (o IntegrationServiceEnvironmentPropertiesPtrOutput) Elem() IntegrationServiceEnvironmentPropertiesOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentProperties) IntegrationServiceEnvironmentProperties {
		if v != nil {
			return *v
		}
		var ret IntegrationServiceEnvironmentProperties
		return ret
	}).(IntegrationServiceEnvironmentPropertiesOutput)
}

func (o IntegrationServiceEnvironmentPropertiesPtrOutput) EncryptionConfiguration() IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentProperties) *IntegrationServiceEnvironmenEncryptionConfiguration {
		if v == nil {
			return nil
		}
		return v.EncryptionConfiguration
	}).(IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput)
}

func (o IntegrationServiceEnvironmentPropertiesPtrOutput) EndpointsConfiguration() FlowEndpointsConfigurationPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentProperties) *FlowEndpointsConfiguration {
		if v == nil {
			return nil
		}
		return v.EndpointsConfiguration
	}).(FlowEndpointsConfigurationPtrOutput)
}

func (o IntegrationServiceEnvironmentPropertiesPtrOutput) IntegrationServiceEnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.IntegrationServiceEnvironmentId
	}).(pulumi.StringPtrOutput)
}

func (o IntegrationServiceEnvironmentPropertiesPtrOutput) NetworkConfiguration() NetworkConfigurationPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentProperties) *NetworkConfiguration {
		if v == nil {
			return nil
		}
		return v.NetworkConfiguration
	}).(NetworkConfigurationPtrOutput)
}

func (o IntegrationServiceEnvironmentPropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o IntegrationServiceEnvironmentPropertiesPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type IntegrationServiceEnvironmentPropertiesResponse struct {
	EncryptionConfiguration         *IntegrationServiceEnvironmenEncryptionConfigurationResponse `pulumi:"encryptionConfiguration"`
	EndpointsConfiguration          *FlowEndpointsConfigurationResponse                          `pulumi:"endpointsConfiguration"`
	IntegrationServiceEnvironmentId *string                                                      `pulumi:"integrationServiceEnvironmentId"`
	NetworkConfiguration            *NetworkConfigurationResponse                                `pulumi:"networkConfiguration"`
	ProvisioningState               *string                                                      `pulumi:"provisioningState"`
	State                           *string                                                      `pulumi:"state"`
}

type IntegrationServiceEnvironmentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironmentPropertiesResponse)(nil)).Elem()
}

func (o IntegrationServiceEnvironmentPropertiesResponseOutput) ToIntegrationServiceEnvironmentPropertiesResponseOutput() IntegrationServiceEnvironmentPropertiesResponseOutput {
	return o
}

func (o IntegrationServiceEnvironmentPropertiesResponseOutput) ToIntegrationServiceEnvironmentPropertiesResponseOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentPropertiesResponseOutput {
	return o
}

func (o IntegrationServiceEnvironmentPropertiesResponseOutput) EncryptionConfiguration() IntegrationServiceEnvironmenEncryptionConfigurationResponsePtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentPropertiesResponse) *IntegrationServiceEnvironmenEncryptionConfigurationResponse {
		return v.EncryptionConfiguration
	}).(IntegrationServiceEnvironmenEncryptionConfigurationResponsePtrOutput)
}

func (o IntegrationServiceEnvironmentPropertiesResponseOutput) EndpointsConfiguration() FlowEndpointsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentPropertiesResponse) *FlowEndpointsConfigurationResponse {
		return v.EndpointsConfiguration
	}).(FlowEndpointsConfigurationResponsePtrOutput)
}

func (o IntegrationServiceEnvironmentPropertiesResponseOutput) IntegrationServiceEnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentPropertiesResponse) *string {
		return v.IntegrationServiceEnvironmentId
	}).(pulumi.StringPtrOutput)
}

func (o IntegrationServiceEnvironmentPropertiesResponseOutput) NetworkConfiguration() NetworkConfigurationResponsePtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentPropertiesResponse) *NetworkConfigurationResponse {
		return v.NetworkConfiguration
	}).(NetworkConfigurationResponsePtrOutput)
}

func (o IntegrationServiceEnvironmentPropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentPropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o IntegrationServiceEnvironmentPropertiesResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentPropertiesResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type IntegrationServiceEnvironmentSku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
}





type IntegrationServiceEnvironmentSkuInput interface {
	pulumi.Input

	ToIntegrationServiceEnvironmentSkuOutput() IntegrationServiceEnvironmentSkuOutput
	ToIntegrationServiceEnvironmentSkuOutputWithContext(context.Context) IntegrationServiceEnvironmentSkuOutput
}

type IntegrationServiceEnvironmentSkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
}

func (IntegrationServiceEnvironmentSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironmentSku)(nil)).Elem()
}

func (i IntegrationServiceEnvironmentSkuArgs) ToIntegrationServiceEnvironmentSkuOutput() IntegrationServiceEnvironmentSkuOutput {
	return i.ToIntegrationServiceEnvironmentSkuOutputWithContext(context.Background())
}

func (i IntegrationServiceEnvironmentSkuArgs) ToIntegrationServiceEnvironmentSkuOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmentSkuOutput)
}

func (i IntegrationServiceEnvironmentSkuArgs) ToIntegrationServiceEnvironmentSkuPtrOutput() IntegrationServiceEnvironmentSkuPtrOutput {
	return i.ToIntegrationServiceEnvironmentSkuPtrOutputWithContext(context.Background())
}

func (i IntegrationServiceEnvironmentSkuArgs) ToIntegrationServiceEnvironmentSkuPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmentSkuOutput).ToIntegrationServiceEnvironmentSkuPtrOutputWithContext(ctx)
}









type IntegrationServiceEnvironmentSkuPtrInput interface {
	pulumi.Input

	ToIntegrationServiceEnvironmentSkuPtrOutput() IntegrationServiceEnvironmentSkuPtrOutput
	ToIntegrationServiceEnvironmentSkuPtrOutputWithContext(context.Context) IntegrationServiceEnvironmentSkuPtrOutput
}

type integrationServiceEnvironmentSkuPtrType IntegrationServiceEnvironmentSkuArgs

func IntegrationServiceEnvironmentSkuPtr(v *IntegrationServiceEnvironmentSkuArgs) IntegrationServiceEnvironmentSkuPtrInput {
	return (*integrationServiceEnvironmentSkuPtrType)(v)
}

func (*integrationServiceEnvironmentSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmentSku)(nil)).Elem()
}

func (i *integrationServiceEnvironmentSkuPtrType) ToIntegrationServiceEnvironmentSkuPtrOutput() IntegrationServiceEnvironmentSkuPtrOutput {
	return i.ToIntegrationServiceEnvironmentSkuPtrOutputWithContext(context.Background())
}

func (i *integrationServiceEnvironmentSkuPtrType) ToIntegrationServiceEnvironmentSkuPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmentSkuPtrOutput)
}

type IntegrationServiceEnvironmentSkuOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironmentSku)(nil)).Elem()
}

func (o IntegrationServiceEnvironmentSkuOutput) ToIntegrationServiceEnvironmentSkuOutput() IntegrationServiceEnvironmentSkuOutput {
	return o
}

func (o IntegrationServiceEnvironmentSkuOutput) ToIntegrationServiceEnvironmentSkuOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentSkuOutput {
	return o
}

func (o IntegrationServiceEnvironmentSkuOutput) ToIntegrationServiceEnvironmentSkuPtrOutput() IntegrationServiceEnvironmentSkuPtrOutput {
	return o.ToIntegrationServiceEnvironmentSkuPtrOutputWithContext(context.Background())
}

func (o IntegrationServiceEnvironmentSkuOutput) ToIntegrationServiceEnvironmentSkuPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IntegrationServiceEnvironmentSku) *IntegrationServiceEnvironmentSku {
		return &v
	}).(IntegrationServiceEnvironmentSkuPtrOutput)
}

func (o IntegrationServiceEnvironmentSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o IntegrationServiceEnvironmentSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type IntegrationServiceEnvironmentSkuPtrOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmentSku)(nil)).Elem()
}

func (o IntegrationServiceEnvironmentSkuPtrOutput) ToIntegrationServiceEnvironmentSkuPtrOutput() IntegrationServiceEnvironmentSkuPtrOutput {
	return o
}

func (o IntegrationServiceEnvironmentSkuPtrOutput) ToIntegrationServiceEnvironmentSkuPtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentSkuPtrOutput {
	return o
}

func (o IntegrationServiceEnvironmentSkuPtrOutput) Elem() IntegrationServiceEnvironmentSkuOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentSku) IntegrationServiceEnvironmentSku {
		if v != nil {
			return *v
		}
		var ret IntegrationServiceEnvironmentSku
		return ret
	}).(IntegrationServiceEnvironmentSkuOutput)
}

func (o IntegrationServiceEnvironmentSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o IntegrationServiceEnvironmentSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type IntegrationServiceEnvironmentSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
}

type IntegrationServiceEnvironmentSkuResponseOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationServiceEnvironmentSkuResponse)(nil)).Elem()
}

func (o IntegrationServiceEnvironmentSkuResponseOutput) ToIntegrationServiceEnvironmentSkuResponseOutput() IntegrationServiceEnvironmentSkuResponseOutput {
	return o
}

func (o IntegrationServiceEnvironmentSkuResponseOutput) ToIntegrationServiceEnvironmentSkuResponseOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentSkuResponseOutput {
	return o
}

func (o IntegrationServiceEnvironmentSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o IntegrationServiceEnvironmentSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationServiceEnvironmentSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type IntegrationServiceEnvironmentSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironmentSkuResponse)(nil)).Elem()
}

func (o IntegrationServiceEnvironmentSkuResponsePtrOutput) ToIntegrationServiceEnvironmentSkuResponsePtrOutput() IntegrationServiceEnvironmentSkuResponsePtrOutput {
	return o
}

func (o IntegrationServiceEnvironmentSkuResponsePtrOutput) ToIntegrationServiceEnvironmentSkuResponsePtrOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentSkuResponsePtrOutput {
	return o
}

func (o IntegrationServiceEnvironmentSkuResponsePtrOutput) Elem() IntegrationServiceEnvironmentSkuResponseOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentSkuResponse) IntegrationServiceEnvironmentSkuResponse {
		if v != nil {
			return *v
		}
		var ret IntegrationServiceEnvironmentSkuResponse
		return ret
	}).(IntegrationServiceEnvironmentSkuResponseOutput)
}

func (o IntegrationServiceEnvironmentSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o IntegrationServiceEnvironmentSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironmentSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type IpAddress struct {
	Address *string `pulumi:"address"`
}





type IpAddressInput interface {
	pulumi.Input

	ToIpAddressOutput() IpAddressOutput
	ToIpAddressOutputWithContext(context.Context) IpAddressOutput
}

type IpAddressArgs struct {
	Address pulumi.StringPtrInput `pulumi:"address"`
}

func (IpAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddress)(nil)).Elem()
}

func (i IpAddressArgs) ToIpAddressOutput() IpAddressOutput {
	return i.ToIpAddressOutputWithContext(context.Background())
}

func (i IpAddressArgs) ToIpAddressOutputWithContext(ctx context.Context) IpAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressOutput)
}





type IpAddressArrayInput interface {
	pulumi.Input

	ToIpAddressArrayOutput() IpAddressArrayOutput
	ToIpAddressArrayOutputWithContext(context.Context) IpAddressArrayOutput
}

type IpAddressArray []IpAddressInput

func (IpAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAddress)(nil)).Elem()
}

func (i IpAddressArray) ToIpAddressArrayOutput() IpAddressArrayOutput {
	return i.ToIpAddressArrayOutputWithContext(context.Background())
}

func (i IpAddressArray) ToIpAddressArrayOutputWithContext(ctx context.Context) IpAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressArrayOutput)
}

type IpAddressOutput struct{ *pulumi.OutputState }

func (IpAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddress)(nil)).Elem()
}

func (o IpAddressOutput) ToIpAddressOutput() IpAddressOutput {
	return o
}

func (o IpAddressOutput) ToIpAddressOutputWithContext(ctx context.Context) IpAddressOutput {
	return o
}

func (o IpAddressOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpAddress) *string { return v.Address }).(pulumi.StringPtrOutput)
}

type IpAddressArrayOutput struct{ *pulumi.OutputState }

func (IpAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAddress)(nil)).Elem()
}

func (o IpAddressArrayOutput) ToIpAddressArrayOutput() IpAddressArrayOutput {
	return o
}

func (o IpAddressArrayOutput) ToIpAddressArrayOutputWithContext(ctx context.Context) IpAddressArrayOutput {
	return o
}

func (o IpAddressArrayOutput) Index(i pulumi.IntInput) IpAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpAddress {
		return vs[0].([]IpAddress)[vs[1].(int)]
	}).(IpAddressOutput)
}

type IpAddressRange struct {
	AddressRange *string `pulumi:"addressRange"`
}





type IpAddressRangeInput interface {
	pulumi.Input

	ToIpAddressRangeOutput() IpAddressRangeOutput
	ToIpAddressRangeOutputWithContext(context.Context) IpAddressRangeOutput
}

type IpAddressRangeArgs struct {
	AddressRange pulumi.StringPtrInput `pulumi:"addressRange"`
}

func (IpAddressRangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddressRange)(nil)).Elem()
}

func (i IpAddressRangeArgs) ToIpAddressRangeOutput() IpAddressRangeOutput {
	return i.ToIpAddressRangeOutputWithContext(context.Background())
}

func (i IpAddressRangeArgs) ToIpAddressRangeOutputWithContext(ctx context.Context) IpAddressRangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressRangeOutput)
}





type IpAddressRangeArrayInput interface {
	pulumi.Input

	ToIpAddressRangeArrayOutput() IpAddressRangeArrayOutput
	ToIpAddressRangeArrayOutputWithContext(context.Context) IpAddressRangeArrayOutput
}

type IpAddressRangeArray []IpAddressRangeInput

func (IpAddressRangeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAddressRange)(nil)).Elem()
}

func (i IpAddressRangeArray) ToIpAddressRangeArrayOutput() IpAddressRangeArrayOutput {
	return i.ToIpAddressRangeArrayOutputWithContext(context.Background())
}

func (i IpAddressRangeArray) ToIpAddressRangeArrayOutputWithContext(ctx context.Context) IpAddressRangeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressRangeArrayOutput)
}

type IpAddressRangeOutput struct{ *pulumi.OutputState }

func (IpAddressRangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddressRange)(nil)).Elem()
}

func (o IpAddressRangeOutput) ToIpAddressRangeOutput() IpAddressRangeOutput {
	return o
}

func (o IpAddressRangeOutput) ToIpAddressRangeOutputWithContext(ctx context.Context) IpAddressRangeOutput {
	return o
}

func (o IpAddressRangeOutput) AddressRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpAddressRange) *string { return v.AddressRange }).(pulumi.StringPtrOutput)
}

type IpAddressRangeArrayOutput struct{ *pulumi.OutputState }

func (IpAddressRangeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAddressRange)(nil)).Elem()
}

func (o IpAddressRangeArrayOutput) ToIpAddressRangeArrayOutput() IpAddressRangeArrayOutput {
	return o
}

func (o IpAddressRangeArrayOutput) ToIpAddressRangeArrayOutputWithContext(ctx context.Context) IpAddressRangeArrayOutput {
	return o
}

func (o IpAddressRangeArrayOutput) Index(i pulumi.IntInput) IpAddressRangeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpAddressRange {
		return vs[0].([]IpAddressRange)[vs[1].(int)]
	}).(IpAddressRangeOutput)
}

type IpAddressRangeResponse struct {
	AddressRange *string `pulumi:"addressRange"`
}

type IpAddressRangeResponseOutput struct{ *pulumi.OutputState }

func (IpAddressRangeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddressRangeResponse)(nil)).Elem()
}

func (o IpAddressRangeResponseOutput) ToIpAddressRangeResponseOutput() IpAddressRangeResponseOutput {
	return o
}

func (o IpAddressRangeResponseOutput) ToIpAddressRangeResponseOutputWithContext(ctx context.Context) IpAddressRangeResponseOutput {
	return o
}

func (o IpAddressRangeResponseOutput) AddressRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpAddressRangeResponse) *string { return v.AddressRange }).(pulumi.StringPtrOutput)
}

type IpAddressRangeResponseArrayOutput struct{ *pulumi.OutputState }

func (IpAddressRangeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAddressRangeResponse)(nil)).Elem()
}

func (o IpAddressRangeResponseArrayOutput) ToIpAddressRangeResponseArrayOutput() IpAddressRangeResponseArrayOutput {
	return o
}

func (o IpAddressRangeResponseArrayOutput) ToIpAddressRangeResponseArrayOutputWithContext(ctx context.Context) IpAddressRangeResponseArrayOutput {
	return o
}

func (o IpAddressRangeResponseArrayOutput) Index(i pulumi.IntInput) IpAddressRangeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpAddressRangeResponse {
		return vs[0].([]IpAddressRangeResponse)[vs[1].(int)]
	}).(IpAddressRangeResponseOutput)
}

type IpAddressResponse struct {
	Address *string `pulumi:"address"`
}

type IpAddressResponseOutput struct{ *pulumi.OutputState }

func (IpAddressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddressResponse)(nil)).Elem()
}

func (o IpAddressResponseOutput) ToIpAddressResponseOutput() IpAddressResponseOutput {
	return o
}

func (o IpAddressResponseOutput) ToIpAddressResponseOutputWithContext(ctx context.Context) IpAddressResponseOutput {
	return o
}

func (o IpAddressResponseOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpAddressResponse) *string { return v.Address }).(pulumi.StringPtrOutput)
}

type IpAddressResponseArrayOutput struct{ *pulumi.OutputState }

func (IpAddressResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAddressResponse)(nil)).Elem()
}

func (o IpAddressResponseArrayOutput) ToIpAddressResponseArrayOutput() IpAddressResponseArrayOutput {
	return o
}

func (o IpAddressResponseArrayOutput) ToIpAddressResponseArrayOutputWithContext(ctx context.Context) IpAddressResponseArrayOutput {
	return o
}

func (o IpAddressResponseArrayOutput) Index(i pulumi.IntInput) IpAddressResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpAddressResponse {
		return vs[0].([]IpAddressResponse)[vs[1].(int)]
	}).(IpAddressResponseOutput)
}

type KeyVaultKeyReference struct {
	KeyName    string                       `pulumi:"keyName"`
	KeyVault   KeyVaultKeyReferenceKeyVault `pulumi:"keyVault"`
	KeyVersion *string                      `pulumi:"keyVersion"`
}





type KeyVaultKeyReferenceInput interface {
	pulumi.Input

	ToKeyVaultKeyReferenceOutput() KeyVaultKeyReferenceOutput
	ToKeyVaultKeyReferenceOutputWithContext(context.Context) KeyVaultKeyReferenceOutput
}

type KeyVaultKeyReferenceArgs struct {
	KeyName    pulumi.StringInput                `pulumi:"keyName"`
	KeyVault   KeyVaultKeyReferenceKeyVaultInput `pulumi:"keyVault"`
	KeyVersion pulumi.StringPtrInput             `pulumi:"keyVersion"`
}

func (KeyVaultKeyReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReference)(nil)).Elem()
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferenceOutput() KeyVaultKeyReferenceOutput {
	return i.ToKeyVaultKeyReferenceOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferenceOutputWithContext(ctx context.Context) KeyVaultKeyReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceOutput)
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return i.ToKeyVaultKeyReferencePtrOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceOutput).ToKeyVaultKeyReferencePtrOutputWithContext(ctx)
}









type KeyVaultKeyReferencePtrInput interface {
	pulumi.Input

	ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput
	ToKeyVaultKeyReferencePtrOutputWithContext(context.Context) KeyVaultKeyReferencePtrOutput
}

type keyVaultKeyReferencePtrType KeyVaultKeyReferenceArgs

func KeyVaultKeyReferencePtr(v *KeyVaultKeyReferenceArgs) KeyVaultKeyReferencePtrInput {
	return (*keyVaultKeyReferencePtrType)(v)
}

func (*keyVaultKeyReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReference)(nil)).Elem()
}

func (i *keyVaultKeyReferencePtrType) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return i.ToKeyVaultKeyReferencePtrOutputWithContext(context.Background())
}

func (i *keyVaultKeyReferencePtrType) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferencePtrOutput)
}

type KeyVaultKeyReferenceOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReference)(nil)).Elem()
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferenceOutput() KeyVaultKeyReferenceOutput {
	return o
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferenceOutputWithContext(ctx context.Context) KeyVaultKeyReferenceOutput {
	return o
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return o.ToKeyVaultKeyReferencePtrOutputWithContext(context.Background())
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultKeyReference) *KeyVaultKeyReference {
		return &v
	}).(KeyVaultKeyReferencePtrOutput)
}

func (o KeyVaultKeyReferenceOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultKeyReference) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o KeyVaultKeyReferenceOutput) KeyVault() KeyVaultKeyReferenceKeyVaultOutput {
	return o.ApplyT(func(v KeyVaultKeyReference) KeyVaultKeyReferenceKeyVault { return v.KeyVault }).(KeyVaultKeyReferenceKeyVaultOutput)
}

func (o KeyVaultKeyReferenceOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyReference) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type KeyVaultKeyReferencePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReference)(nil)).Elem()
}

func (o KeyVaultKeyReferencePtrOutput) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return o
}

func (o KeyVaultKeyReferencePtrOutput) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return o
}

func (o KeyVaultKeyReferencePtrOutput) Elem() KeyVaultKeyReferenceOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) KeyVaultKeyReference {
		if v != nil {
			return *v
		}
		var ret KeyVaultKeyReference
		return ret
	}).(KeyVaultKeyReferenceOutput)
}

func (o KeyVaultKeyReferencePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) *string {
		if v == nil {
			return nil
		}
		return &v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferencePtrOutput) KeyVault() KeyVaultKeyReferenceKeyVaultPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) *KeyVaultKeyReferenceKeyVault {
		if v == nil {
			return nil
		}
		return &v.KeyVault
	}).(KeyVaultKeyReferenceKeyVaultPtrOutput)
}

func (o KeyVaultKeyReferencePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type KeyVaultKeyReferenceKeyVault struct {
	Id *string `pulumi:"id"`
}





type KeyVaultKeyReferenceKeyVaultInput interface {
	pulumi.Input

	ToKeyVaultKeyReferenceKeyVaultOutput() KeyVaultKeyReferenceKeyVaultOutput
	ToKeyVaultKeyReferenceKeyVaultOutputWithContext(context.Context) KeyVaultKeyReferenceKeyVaultOutput
}

type KeyVaultKeyReferenceKeyVaultArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (KeyVaultKeyReferenceKeyVaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReferenceKeyVault)(nil)).Elem()
}

func (i KeyVaultKeyReferenceKeyVaultArgs) ToKeyVaultKeyReferenceKeyVaultOutput() KeyVaultKeyReferenceKeyVaultOutput {
	return i.ToKeyVaultKeyReferenceKeyVaultOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceKeyVaultArgs) ToKeyVaultKeyReferenceKeyVaultOutputWithContext(ctx context.Context) KeyVaultKeyReferenceKeyVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceKeyVaultOutput)
}

func (i KeyVaultKeyReferenceKeyVaultArgs) ToKeyVaultKeyReferenceKeyVaultPtrOutput() KeyVaultKeyReferenceKeyVaultPtrOutput {
	return i.ToKeyVaultKeyReferenceKeyVaultPtrOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceKeyVaultArgs) ToKeyVaultKeyReferenceKeyVaultPtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceKeyVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceKeyVaultOutput).ToKeyVaultKeyReferenceKeyVaultPtrOutputWithContext(ctx)
}









type KeyVaultKeyReferenceKeyVaultPtrInput interface {
	pulumi.Input

	ToKeyVaultKeyReferenceKeyVaultPtrOutput() KeyVaultKeyReferenceKeyVaultPtrOutput
	ToKeyVaultKeyReferenceKeyVaultPtrOutputWithContext(context.Context) KeyVaultKeyReferenceKeyVaultPtrOutput
}

type keyVaultKeyReferenceKeyVaultPtrType KeyVaultKeyReferenceKeyVaultArgs

func KeyVaultKeyReferenceKeyVaultPtr(v *KeyVaultKeyReferenceKeyVaultArgs) KeyVaultKeyReferenceKeyVaultPtrInput {
	return (*keyVaultKeyReferenceKeyVaultPtrType)(v)
}

func (*keyVaultKeyReferenceKeyVaultPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReferenceKeyVault)(nil)).Elem()
}

func (i *keyVaultKeyReferenceKeyVaultPtrType) ToKeyVaultKeyReferenceKeyVaultPtrOutput() KeyVaultKeyReferenceKeyVaultPtrOutput {
	return i.ToKeyVaultKeyReferenceKeyVaultPtrOutputWithContext(context.Background())
}

func (i *keyVaultKeyReferenceKeyVaultPtrType) ToKeyVaultKeyReferenceKeyVaultPtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceKeyVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceKeyVaultPtrOutput)
}

type KeyVaultKeyReferenceKeyVaultOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceKeyVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReferenceKeyVault)(nil)).Elem()
}

func (o KeyVaultKeyReferenceKeyVaultOutput) ToKeyVaultKeyReferenceKeyVaultOutput() KeyVaultKeyReferenceKeyVaultOutput {
	return o
}

func (o KeyVaultKeyReferenceKeyVaultOutput) ToKeyVaultKeyReferenceKeyVaultOutputWithContext(ctx context.Context) KeyVaultKeyReferenceKeyVaultOutput {
	return o
}

func (o KeyVaultKeyReferenceKeyVaultOutput) ToKeyVaultKeyReferenceKeyVaultPtrOutput() KeyVaultKeyReferenceKeyVaultPtrOutput {
	return o.ToKeyVaultKeyReferenceKeyVaultPtrOutputWithContext(context.Background())
}

func (o KeyVaultKeyReferenceKeyVaultOutput) ToKeyVaultKeyReferenceKeyVaultPtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceKeyVaultPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultKeyReferenceKeyVault) *KeyVaultKeyReferenceKeyVault {
		return &v
	}).(KeyVaultKeyReferenceKeyVaultPtrOutput)
}

func (o KeyVaultKeyReferenceKeyVaultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceKeyVault) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type KeyVaultKeyReferenceKeyVaultPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceKeyVaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReferenceKeyVault)(nil)).Elem()
}

func (o KeyVaultKeyReferenceKeyVaultPtrOutput) ToKeyVaultKeyReferenceKeyVaultPtrOutput() KeyVaultKeyReferenceKeyVaultPtrOutput {
	return o
}

func (o KeyVaultKeyReferenceKeyVaultPtrOutput) ToKeyVaultKeyReferenceKeyVaultPtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceKeyVaultPtrOutput {
	return o
}

func (o KeyVaultKeyReferenceKeyVaultPtrOutput) Elem() KeyVaultKeyReferenceKeyVaultOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceKeyVault) KeyVaultKeyReferenceKeyVault {
		if v != nil {
			return *v
		}
		var ret KeyVaultKeyReferenceKeyVault
		return ret
	}).(KeyVaultKeyReferenceKeyVaultOutput)
}

func (o KeyVaultKeyReferenceKeyVaultPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceKeyVault) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type KeyVaultKeyReferenceResponse struct {
	KeyName    string                               `pulumi:"keyName"`
	KeyVault   KeyVaultKeyReferenceResponseKeyVault `pulumi:"keyVault"`
	KeyVersion *string                              `pulumi:"keyVersion"`
}

type KeyVaultKeyReferenceResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReferenceResponse)(nil)).Elem()
}

func (o KeyVaultKeyReferenceResponseOutput) ToKeyVaultKeyReferenceResponseOutput() KeyVaultKeyReferenceResponseOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseOutput) ToKeyVaultKeyReferenceResponseOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponseOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponse) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o KeyVaultKeyReferenceResponseOutput) KeyVault() KeyVaultKeyReferenceResponseKeyVaultOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponse) KeyVaultKeyReferenceResponseKeyVault { return v.KeyVault }).(KeyVaultKeyReferenceResponseKeyVaultOutput)
}

func (o KeyVaultKeyReferenceResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type KeyVaultKeyReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReferenceResponse)(nil)).Elem()
}

func (o KeyVaultKeyReferenceResponsePtrOutput) ToKeyVaultKeyReferenceResponsePtrOutput() KeyVaultKeyReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultKeyReferenceResponsePtrOutput) ToKeyVaultKeyReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultKeyReferenceResponsePtrOutput) Elem() KeyVaultKeyReferenceResponseOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) KeyVaultKeyReferenceResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultKeyReferenceResponse
		return ret
	}).(KeyVaultKeyReferenceResponseOutput)
}

func (o KeyVaultKeyReferenceResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferenceResponsePtrOutput) KeyVault() KeyVaultKeyReferenceResponseKeyVaultPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) *KeyVaultKeyReferenceResponseKeyVault {
		if v == nil {
			return nil
		}
		return &v.KeyVault
	}).(KeyVaultKeyReferenceResponseKeyVaultPtrOutput)
}

func (o KeyVaultKeyReferenceResponsePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type KeyVaultKeyReferenceResponseKeyVault struct {
	Id   *string `pulumi:"id"`
	Name string  `pulumi:"name"`
	Type string  `pulumi:"type"`
}

type KeyVaultKeyReferenceResponseKeyVaultOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceResponseKeyVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReferenceResponseKeyVault)(nil)).Elem()
}

func (o KeyVaultKeyReferenceResponseKeyVaultOutput) ToKeyVaultKeyReferenceResponseKeyVaultOutput() KeyVaultKeyReferenceResponseKeyVaultOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseKeyVaultOutput) ToKeyVaultKeyReferenceResponseKeyVaultOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponseKeyVaultOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseKeyVaultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponseKeyVault) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferenceResponseKeyVaultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponseKeyVault) string { return v.Name }).(pulumi.StringOutput)
}

func (o KeyVaultKeyReferenceResponseKeyVaultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponseKeyVault) string { return v.Type }).(pulumi.StringOutput)
}

type KeyVaultKeyReferenceResponseKeyVaultPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceResponseKeyVaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReferenceResponseKeyVault)(nil)).Elem()
}

func (o KeyVaultKeyReferenceResponseKeyVaultPtrOutput) ToKeyVaultKeyReferenceResponseKeyVaultPtrOutput() KeyVaultKeyReferenceResponseKeyVaultPtrOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseKeyVaultPtrOutput) ToKeyVaultKeyReferenceResponseKeyVaultPtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponseKeyVaultPtrOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseKeyVaultPtrOutput) Elem() KeyVaultKeyReferenceResponseKeyVaultOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponseKeyVault) KeyVaultKeyReferenceResponseKeyVault {
		if v != nil {
			return *v
		}
		var ret KeyVaultKeyReferenceResponseKeyVault
		return ret
	}).(KeyVaultKeyReferenceResponseKeyVaultOutput)
}

func (o KeyVaultKeyReferenceResponseKeyVaultPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponseKeyVault) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferenceResponseKeyVaultPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponseKeyVault) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferenceResponseKeyVaultPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponseKeyVault) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type KeyVaultKeyResponse struct {
	Attributes *KeyVaultKeyResponseAttributes `pulumi:"attributes"`
	Kid        *string                        `pulumi:"kid"`
}

type KeyVaultKeyResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyResponse)(nil)).Elem()
}

func (o KeyVaultKeyResponseOutput) ToKeyVaultKeyResponseOutput() KeyVaultKeyResponseOutput {
	return o
}

func (o KeyVaultKeyResponseOutput) ToKeyVaultKeyResponseOutputWithContext(ctx context.Context) KeyVaultKeyResponseOutput {
	return o
}

func (o KeyVaultKeyResponseOutput) Attributes() KeyVaultKeyResponseAttributesPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyResponse) *KeyVaultKeyResponseAttributes { return v.Attributes }).(KeyVaultKeyResponseAttributesPtrOutput)
}

func (o KeyVaultKeyResponseOutput) Kid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyResponse) *string { return v.Kid }).(pulumi.StringPtrOutput)
}

type KeyVaultKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyVaultKeyResponse)(nil)).Elem()
}

func (o KeyVaultKeyResponseArrayOutput) ToKeyVaultKeyResponseArrayOutput() KeyVaultKeyResponseArrayOutput {
	return o
}

func (o KeyVaultKeyResponseArrayOutput) ToKeyVaultKeyResponseArrayOutputWithContext(ctx context.Context) KeyVaultKeyResponseArrayOutput {
	return o
}

func (o KeyVaultKeyResponseArrayOutput) Index(i pulumi.IntInput) KeyVaultKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyVaultKeyResponse {
		return vs[0].([]KeyVaultKeyResponse)[vs[1].(int)]
	}).(KeyVaultKeyResponseOutput)
}

type KeyVaultKeyResponseAttributes struct {
	Created *float64 `pulumi:"created"`
	Enabled *bool    `pulumi:"enabled"`
	Updated *float64 `pulumi:"updated"`
}

type KeyVaultKeyResponseAttributesOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyResponseAttributesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyResponseAttributes)(nil)).Elem()
}

func (o KeyVaultKeyResponseAttributesOutput) ToKeyVaultKeyResponseAttributesOutput() KeyVaultKeyResponseAttributesOutput {
	return o
}

func (o KeyVaultKeyResponseAttributesOutput) ToKeyVaultKeyResponseAttributesOutputWithContext(ctx context.Context) KeyVaultKeyResponseAttributesOutput {
	return o
}

func (o KeyVaultKeyResponseAttributesOutput) Created() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KeyVaultKeyResponseAttributes) *float64 { return v.Created }).(pulumi.Float64PtrOutput)
}

func (o KeyVaultKeyResponseAttributesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyResponseAttributes) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o KeyVaultKeyResponseAttributesOutput) Updated() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KeyVaultKeyResponseAttributes) *float64 { return v.Updated }).(pulumi.Float64PtrOutput)
}

type KeyVaultKeyResponseAttributesPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyResponseAttributesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyResponseAttributes)(nil)).Elem()
}

func (o KeyVaultKeyResponseAttributesPtrOutput) ToKeyVaultKeyResponseAttributesPtrOutput() KeyVaultKeyResponseAttributesPtrOutput {
	return o
}

func (o KeyVaultKeyResponseAttributesPtrOutput) ToKeyVaultKeyResponseAttributesPtrOutputWithContext(ctx context.Context) KeyVaultKeyResponseAttributesPtrOutput {
	return o
}

func (o KeyVaultKeyResponseAttributesPtrOutput) Elem() KeyVaultKeyResponseAttributesOutput {
	return o.ApplyT(func(v *KeyVaultKeyResponseAttributes) KeyVaultKeyResponseAttributes {
		if v != nil {
			return *v
		}
		var ret KeyVaultKeyResponseAttributes
		return ret
	}).(KeyVaultKeyResponseAttributesOutput)
}

func (o KeyVaultKeyResponseAttributesPtrOutput) Created() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyResponseAttributes) *float64 {
		if v == nil {
			return nil
		}
		return v.Created
	}).(pulumi.Float64PtrOutput)
}

func (o KeyVaultKeyResponseAttributesPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyResponseAttributes) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o KeyVaultKeyResponseAttributesPtrOutput) Updated() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyResponseAttributes) *float64 {
		if v == nil {
			return nil
		}
		return v.Updated
	}).(pulumi.Float64PtrOutput)
}

type KeyVaultReference struct {
	Id   *string `pulumi:"id"`
	Name *string `pulumi:"name"`
}





type KeyVaultReferenceInput interface {
	pulumi.Input

	ToKeyVaultReferenceOutput() KeyVaultReferenceOutput
	ToKeyVaultReferenceOutputWithContext(context.Context) KeyVaultReferenceOutput
}

type KeyVaultReferenceArgs struct {
	Id   pulumi.StringPtrInput `pulumi:"id"`
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (KeyVaultReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultReference)(nil)).Elem()
}

func (i KeyVaultReferenceArgs) ToKeyVaultReferenceOutput() KeyVaultReferenceOutput {
	return i.ToKeyVaultReferenceOutputWithContext(context.Background())
}

func (i KeyVaultReferenceArgs) ToKeyVaultReferenceOutputWithContext(ctx context.Context) KeyVaultReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultReferenceOutput)
}

type KeyVaultReferenceOutput struct{ *pulumi.OutputState }

func (KeyVaultReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultReference)(nil)).Elem()
}

func (o KeyVaultReferenceOutput) ToKeyVaultReferenceOutput() KeyVaultReferenceOutput {
	return o
}

func (o KeyVaultReferenceOutput) ToKeyVaultReferenceOutputWithContext(ctx context.Context) KeyVaultReferenceOutput {
	return o
}

func (o KeyVaultReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o KeyVaultReferenceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultReference) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ManagedServiceIdentity struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ManagedServiceIdentityInput interface {
	pulumi.Input

	ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput
	ToManagedServiceIdentityOutputWithContext(context.Context) ManagedServiceIdentityOutput
}

type ManagedServiceIdentityArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
}

func (ManagedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return i.ToManagedServiceIdentityOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput)
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput).ToManagedServiceIdentityPtrOutputWithContext(ctx)
}









type ManagedServiceIdentityPtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput
	ToManagedServiceIdentityPtrOutputWithContext(context.Context) ManagedServiceIdentityPtrOutput
}

type managedServiceIdentityPtrType ManagedServiceIdentityArgs

func ManagedServiceIdentityPtr(v *ManagedServiceIdentityArgs) ManagedServiceIdentityPtrInput {
	return (*managedServiceIdentityPtrType)(v)
}

func (*managedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityPtrOutput)
}

type ManagedServiceIdentityOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentity) *ManagedServiceIdentity {
		return &v
	}).(ManagedServiceIdentityPtrOutput)
}

func (o ManagedServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) Elem() ManagedServiceIdentityOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) ManagedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentity
		return ret
	}).(ManagedServiceIdentityOutput)
}

func (o ManagedServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ManagedServiceIdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   string                                  `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}

type ManagedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ManagedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) Elem() ManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) ManagedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityResponse
		return ret
	}).(ManagedServiceIdentityResponseOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type NetworkConfiguration struct {
	AccessEndpoint             *IntegrationServiceEnvironmentAccessEndpoint `pulumi:"accessEndpoint"`
	Subnets                    []ResourceReference                          `pulumi:"subnets"`
	VirtualNetworkAddressSpace *string                                      `pulumi:"virtualNetworkAddressSpace"`
}





type NetworkConfigurationInput interface {
	pulumi.Input

	ToNetworkConfigurationOutput() NetworkConfigurationOutput
	ToNetworkConfigurationOutputWithContext(context.Context) NetworkConfigurationOutput
}

type NetworkConfigurationArgs struct {
	AccessEndpoint             IntegrationServiceEnvironmentAccessEndpointPtrInput `pulumi:"accessEndpoint"`
	Subnets                    ResourceReferenceArrayInput                         `pulumi:"subnets"`
	VirtualNetworkAddressSpace pulumi.StringPtrInput                               `pulumi:"virtualNetworkAddressSpace"`
}

func (NetworkConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkConfiguration)(nil)).Elem()
}

func (i NetworkConfigurationArgs) ToNetworkConfigurationOutput() NetworkConfigurationOutput {
	return i.ToNetworkConfigurationOutputWithContext(context.Background())
}

func (i NetworkConfigurationArgs) ToNetworkConfigurationOutputWithContext(ctx context.Context) NetworkConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConfigurationOutput)
}

func (i NetworkConfigurationArgs) ToNetworkConfigurationPtrOutput() NetworkConfigurationPtrOutput {
	return i.ToNetworkConfigurationPtrOutputWithContext(context.Background())
}

func (i NetworkConfigurationArgs) ToNetworkConfigurationPtrOutputWithContext(ctx context.Context) NetworkConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConfigurationOutput).ToNetworkConfigurationPtrOutputWithContext(ctx)
}









type NetworkConfigurationPtrInput interface {
	pulumi.Input

	ToNetworkConfigurationPtrOutput() NetworkConfigurationPtrOutput
	ToNetworkConfigurationPtrOutputWithContext(context.Context) NetworkConfigurationPtrOutput
}

type networkConfigurationPtrType NetworkConfigurationArgs

func NetworkConfigurationPtr(v *NetworkConfigurationArgs) NetworkConfigurationPtrInput {
	return (*networkConfigurationPtrType)(v)
}

func (*networkConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkConfiguration)(nil)).Elem()
}

func (i *networkConfigurationPtrType) ToNetworkConfigurationPtrOutput() NetworkConfigurationPtrOutput {
	return i.ToNetworkConfigurationPtrOutputWithContext(context.Background())
}

func (i *networkConfigurationPtrType) ToNetworkConfigurationPtrOutputWithContext(ctx context.Context) NetworkConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConfigurationPtrOutput)
}

type NetworkConfigurationOutput struct{ *pulumi.OutputState }

func (NetworkConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkConfiguration)(nil)).Elem()
}

func (o NetworkConfigurationOutput) ToNetworkConfigurationOutput() NetworkConfigurationOutput {
	return o
}

func (o NetworkConfigurationOutput) ToNetworkConfigurationOutputWithContext(ctx context.Context) NetworkConfigurationOutput {
	return o
}

func (o NetworkConfigurationOutput) ToNetworkConfigurationPtrOutput() NetworkConfigurationPtrOutput {
	return o.ToNetworkConfigurationPtrOutputWithContext(context.Background())
}

func (o NetworkConfigurationOutput) ToNetworkConfigurationPtrOutputWithContext(ctx context.Context) NetworkConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkConfiguration) *NetworkConfiguration {
		return &v
	}).(NetworkConfigurationPtrOutput)
}

func (o NetworkConfigurationOutput) AccessEndpoint() IntegrationServiceEnvironmentAccessEndpointPtrOutput {
	return o.ApplyT(func(v NetworkConfiguration) *IntegrationServiceEnvironmentAccessEndpoint { return v.AccessEndpoint }).(IntegrationServiceEnvironmentAccessEndpointPtrOutput)
}

func (o NetworkConfigurationOutput) Subnets() ResourceReferenceArrayOutput {
	return o.ApplyT(func(v NetworkConfiguration) []ResourceReference { return v.Subnets }).(ResourceReferenceArrayOutput)
}

func (o NetworkConfigurationOutput) VirtualNetworkAddressSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkConfiguration) *string { return v.VirtualNetworkAddressSpace }).(pulumi.StringPtrOutput)
}

type NetworkConfigurationPtrOutput struct{ *pulumi.OutputState }

func (NetworkConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkConfiguration)(nil)).Elem()
}

func (o NetworkConfigurationPtrOutput) ToNetworkConfigurationPtrOutput() NetworkConfigurationPtrOutput {
	return o
}

func (o NetworkConfigurationPtrOutput) ToNetworkConfigurationPtrOutputWithContext(ctx context.Context) NetworkConfigurationPtrOutput {
	return o
}

func (o NetworkConfigurationPtrOutput) Elem() NetworkConfigurationOutput {
	return o.ApplyT(func(v *NetworkConfiguration) NetworkConfiguration {
		if v != nil {
			return *v
		}
		var ret NetworkConfiguration
		return ret
	}).(NetworkConfigurationOutput)
}

func (o NetworkConfigurationPtrOutput) AccessEndpoint() IntegrationServiceEnvironmentAccessEndpointPtrOutput {
	return o.ApplyT(func(v *NetworkConfiguration) *IntegrationServiceEnvironmentAccessEndpoint {
		if v == nil {
			return nil
		}
		return v.AccessEndpoint
	}).(IntegrationServiceEnvironmentAccessEndpointPtrOutput)
}

func (o NetworkConfigurationPtrOutput) Subnets() ResourceReferenceArrayOutput {
	return o.ApplyT(func(v *NetworkConfiguration) []ResourceReference {
		if v == nil {
			return nil
		}
		return v.Subnets
	}).(ResourceReferenceArrayOutput)
}

func (o NetworkConfigurationPtrOutput) VirtualNetworkAddressSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkAddressSpace
	}).(pulumi.StringPtrOutput)
}

type NetworkConfigurationResponse struct {
	AccessEndpoint             *IntegrationServiceEnvironmentAccessEndpointResponse `pulumi:"accessEndpoint"`
	Subnets                    []ResourceReferenceResponse                          `pulumi:"subnets"`
	VirtualNetworkAddressSpace *string                                              `pulumi:"virtualNetworkAddressSpace"`
}

type NetworkConfigurationResponseOutput struct{ *pulumi.OutputState }

func (NetworkConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkConfigurationResponse)(nil)).Elem()
}

func (o NetworkConfigurationResponseOutput) ToNetworkConfigurationResponseOutput() NetworkConfigurationResponseOutput {
	return o
}

func (o NetworkConfigurationResponseOutput) ToNetworkConfigurationResponseOutputWithContext(ctx context.Context) NetworkConfigurationResponseOutput {
	return o
}

func (o NetworkConfigurationResponseOutput) AccessEndpoint() IntegrationServiceEnvironmentAccessEndpointResponsePtrOutput {
	return o.ApplyT(func(v NetworkConfigurationResponse) *IntegrationServiceEnvironmentAccessEndpointResponse {
		return v.AccessEndpoint
	}).(IntegrationServiceEnvironmentAccessEndpointResponsePtrOutput)
}

func (o NetworkConfigurationResponseOutput) Subnets() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v NetworkConfigurationResponse) []ResourceReferenceResponse { return v.Subnets }).(ResourceReferenceResponseArrayOutput)
}

func (o NetworkConfigurationResponseOutput) VirtualNetworkAddressSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkConfigurationResponse) *string { return v.VirtualNetworkAddressSpace }).(pulumi.StringPtrOutput)
}

type NetworkConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkConfigurationResponse)(nil)).Elem()
}

func (o NetworkConfigurationResponsePtrOutput) ToNetworkConfigurationResponsePtrOutput() NetworkConfigurationResponsePtrOutput {
	return o
}

func (o NetworkConfigurationResponsePtrOutput) ToNetworkConfigurationResponsePtrOutputWithContext(ctx context.Context) NetworkConfigurationResponsePtrOutput {
	return o
}

func (o NetworkConfigurationResponsePtrOutput) Elem() NetworkConfigurationResponseOutput {
	return o.ApplyT(func(v *NetworkConfigurationResponse) NetworkConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret NetworkConfigurationResponse
		return ret
	}).(NetworkConfigurationResponseOutput)
}

func (o NetworkConfigurationResponsePtrOutput) AccessEndpoint() IntegrationServiceEnvironmentAccessEndpointResponsePtrOutput {
	return o.ApplyT(func(v *NetworkConfigurationResponse) *IntegrationServiceEnvironmentAccessEndpointResponse {
		if v == nil {
			return nil
		}
		return v.AccessEndpoint
	}).(IntegrationServiceEnvironmentAccessEndpointResponsePtrOutput)
}

func (o NetworkConfigurationResponsePtrOutput) Subnets() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *NetworkConfigurationResponse) []ResourceReferenceResponse {
		if v == nil {
			return nil
		}
		return v.Subnets
	}).(ResourceReferenceResponseArrayOutput)
}

func (o NetworkConfigurationResponsePtrOutput) VirtualNetworkAddressSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkAddressSpace
	}).(pulumi.StringPtrOutput)
}

type OpenAuthenticationAccessPolicies struct {
	Policies map[string]OpenAuthenticationAccessPolicy `pulumi:"policies"`
}





type OpenAuthenticationAccessPoliciesInput interface {
	pulumi.Input

	ToOpenAuthenticationAccessPoliciesOutput() OpenAuthenticationAccessPoliciesOutput
	ToOpenAuthenticationAccessPoliciesOutputWithContext(context.Context) OpenAuthenticationAccessPoliciesOutput
}

type OpenAuthenticationAccessPoliciesArgs struct {
	Policies OpenAuthenticationAccessPolicyMapInput `pulumi:"policies"`
}

func (OpenAuthenticationAccessPoliciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenAuthenticationAccessPolicies)(nil)).Elem()
}

func (i OpenAuthenticationAccessPoliciesArgs) ToOpenAuthenticationAccessPoliciesOutput() OpenAuthenticationAccessPoliciesOutput {
	return i.ToOpenAuthenticationAccessPoliciesOutputWithContext(context.Background())
}

func (i OpenAuthenticationAccessPoliciesArgs) ToOpenAuthenticationAccessPoliciesOutputWithContext(ctx context.Context) OpenAuthenticationAccessPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenAuthenticationAccessPoliciesOutput)
}

func (i OpenAuthenticationAccessPoliciesArgs) ToOpenAuthenticationAccessPoliciesPtrOutput() OpenAuthenticationAccessPoliciesPtrOutput {
	return i.ToOpenAuthenticationAccessPoliciesPtrOutputWithContext(context.Background())
}

func (i OpenAuthenticationAccessPoliciesArgs) ToOpenAuthenticationAccessPoliciesPtrOutputWithContext(ctx context.Context) OpenAuthenticationAccessPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenAuthenticationAccessPoliciesOutput).ToOpenAuthenticationAccessPoliciesPtrOutputWithContext(ctx)
}









type OpenAuthenticationAccessPoliciesPtrInput interface {
	pulumi.Input

	ToOpenAuthenticationAccessPoliciesPtrOutput() OpenAuthenticationAccessPoliciesPtrOutput
	ToOpenAuthenticationAccessPoliciesPtrOutputWithContext(context.Context) OpenAuthenticationAccessPoliciesPtrOutput
}

type openAuthenticationAccessPoliciesPtrType OpenAuthenticationAccessPoliciesArgs

func OpenAuthenticationAccessPoliciesPtr(v *OpenAuthenticationAccessPoliciesArgs) OpenAuthenticationAccessPoliciesPtrInput {
	return (*openAuthenticationAccessPoliciesPtrType)(v)
}

func (*openAuthenticationAccessPoliciesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenAuthenticationAccessPolicies)(nil)).Elem()
}

func (i *openAuthenticationAccessPoliciesPtrType) ToOpenAuthenticationAccessPoliciesPtrOutput() OpenAuthenticationAccessPoliciesPtrOutput {
	return i.ToOpenAuthenticationAccessPoliciesPtrOutputWithContext(context.Background())
}

func (i *openAuthenticationAccessPoliciesPtrType) ToOpenAuthenticationAccessPoliciesPtrOutputWithContext(ctx context.Context) OpenAuthenticationAccessPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenAuthenticationAccessPoliciesPtrOutput)
}

type OpenAuthenticationAccessPoliciesOutput struct{ *pulumi.OutputState }

func (OpenAuthenticationAccessPoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenAuthenticationAccessPolicies)(nil)).Elem()
}

func (o OpenAuthenticationAccessPoliciesOutput) ToOpenAuthenticationAccessPoliciesOutput() OpenAuthenticationAccessPoliciesOutput {
	return o
}

func (o OpenAuthenticationAccessPoliciesOutput) ToOpenAuthenticationAccessPoliciesOutputWithContext(ctx context.Context) OpenAuthenticationAccessPoliciesOutput {
	return o
}

func (o OpenAuthenticationAccessPoliciesOutput) ToOpenAuthenticationAccessPoliciesPtrOutput() OpenAuthenticationAccessPoliciesPtrOutput {
	return o.ToOpenAuthenticationAccessPoliciesPtrOutputWithContext(context.Background())
}

func (o OpenAuthenticationAccessPoliciesOutput) ToOpenAuthenticationAccessPoliciesPtrOutputWithContext(ctx context.Context) OpenAuthenticationAccessPoliciesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OpenAuthenticationAccessPolicies) *OpenAuthenticationAccessPolicies {
		return &v
	}).(OpenAuthenticationAccessPoliciesPtrOutput)
}

func (o OpenAuthenticationAccessPoliciesOutput) Policies() OpenAuthenticationAccessPolicyMapOutput {
	return o.ApplyT(func(v OpenAuthenticationAccessPolicies) map[string]OpenAuthenticationAccessPolicy { return v.Policies }).(OpenAuthenticationAccessPolicyMapOutput)
}

type OpenAuthenticationAccessPoliciesPtrOutput struct{ *pulumi.OutputState }

func (OpenAuthenticationAccessPoliciesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenAuthenticationAccessPolicies)(nil)).Elem()
}

func (o OpenAuthenticationAccessPoliciesPtrOutput) ToOpenAuthenticationAccessPoliciesPtrOutput() OpenAuthenticationAccessPoliciesPtrOutput {
	return o
}

func (o OpenAuthenticationAccessPoliciesPtrOutput) ToOpenAuthenticationAccessPoliciesPtrOutputWithContext(ctx context.Context) OpenAuthenticationAccessPoliciesPtrOutput {
	return o
}

func (o OpenAuthenticationAccessPoliciesPtrOutput) Elem() OpenAuthenticationAccessPoliciesOutput {
	return o.ApplyT(func(v *OpenAuthenticationAccessPolicies) OpenAuthenticationAccessPolicies {
		if v != nil {
			return *v
		}
		var ret OpenAuthenticationAccessPolicies
		return ret
	}).(OpenAuthenticationAccessPoliciesOutput)
}

func (o OpenAuthenticationAccessPoliciesPtrOutput) Policies() OpenAuthenticationAccessPolicyMapOutput {
	return o.ApplyT(func(v *OpenAuthenticationAccessPolicies) map[string]OpenAuthenticationAccessPolicy {
		if v == nil {
			return nil
		}
		return v.Policies
	}).(OpenAuthenticationAccessPolicyMapOutput)
}

type OpenAuthenticationAccessPoliciesResponse struct {
	Policies map[string]OpenAuthenticationAccessPolicyResponse `pulumi:"policies"`
}

type OpenAuthenticationAccessPoliciesResponseOutput struct{ *pulumi.OutputState }

func (OpenAuthenticationAccessPoliciesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenAuthenticationAccessPoliciesResponse)(nil)).Elem()
}

func (o OpenAuthenticationAccessPoliciesResponseOutput) ToOpenAuthenticationAccessPoliciesResponseOutput() OpenAuthenticationAccessPoliciesResponseOutput {
	return o
}

func (o OpenAuthenticationAccessPoliciesResponseOutput) ToOpenAuthenticationAccessPoliciesResponseOutputWithContext(ctx context.Context) OpenAuthenticationAccessPoliciesResponseOutput {
	return o
}

func (o OpenAuthenticationAccessPoliciesResponseOutput) Policies() OpenAuthenticationAccessPolicyResponseMapOutput {
	return o.ApplyT(func(v OpenAuthenticationAccessPoliciesResponse) map[string]OpenAuthenticationAccessPolicyResponse {
		return v.Policies
	}).(OpenAuthenticationAccessPolicyResponseMapOutput)
}

type OpenAuthenticationAccessPoliciesResponsePtrOutput struct{ *pulumi.OutputState }

func (OpenAuthenticationAccessPoliciesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenAuthenticationAccessPoliciesResponse)(nil)).Elem()
}

func (o OpenAuthenticationAccessPoliciesResponsePtrOutput) ToOpenAuthenticationAccessPoliciesResponsePtrOutput() OpenAuthenticationAccessPoliciesResponsePtrOutput {
	return o
}

func (o OpenAuthenticationAccessPoliciesResponsePtrOutput) ToOpenAuthenticationAccessPoliciesResponsePtrOutputWithContext(ctx context.Context) OpenAuthenticationAccessPoliciesResponsePtrOutput {
	return o
}

func (o OpenAuthenticationAccessPoliciesResponsePtrOutput) Elem() OpenAuthenticationAccessPoliciesResponseOutput {
	return o.ApplyT(func(v *OpenAuthenticationAccessPoliciesResponse) OpenAuthenticationAccessPoliciesResponse {
		if v != nil {
			return *v
		}
		var ret OpenAuthenticationAccessPoliciesResponse
		return ret
	}).(OpenAuthenticationAccessPoliciesResponseOutput)
}

func (o OpenAuthenticationAccessPoliciesResponsePtrOutput) Policies() OpenAuthenticationAccessPolicyResponseMapOutput {
	return o.ApplyT(func(v *OpenAuthenticationAccessPoliciesResponse) map[string]OpenAuthenticationAccessPolicyResponse {
		if v == nil {
			return nil
		}
		return v.Policies
	}).(OpenAuthenticationAccessPolicyResponseMapOutput)
}

type OpenAuthenticationAccessPolicy struct {
	Claims []OpenAuthenticationPolicyClaim `pulumi:"claims"`
	Type   *string                         `pulumi:"type"`
}





type OpenAuthenticationAccessPolicyInput interface {
	pulumi.Input

	ToOpenAuthenticationAccessPolicyOutput() OpenAuthenticationAccessPolicyOutput
	ToOpenAuthenticationAccessPolicyOutputWithContext(context.Context) OpenAuthenticationAccessPolicyOutput
}

type OpenAuthenticationAccessPolicyArgs struct {
	Claims OpenAuthenticationPolicyClaimArrayInput `pulumi:"claims"`
	Type   pulumi.StringPtrInput                   `pulumi:"type"`
}

func (OpenAuthenticationAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenAuthenticationAccessPolicy)(nil)).Elem()
}

func (i OpenAuthenticationAccessPolicyArgs) ToOpenAuthenticationAccessPolicyOutput() OpenAuthenticationAccessPolicyOutput {
	return i.ToOpenAuthenticationAccessPolicyOutputWithContext(context.Background())
}

func (i OpenAuthenticationAccessPolicyArgs) ToOpenAuthenticationAccessPolicyOutputWithContext(ctx context.Context) OpenAuthenticationAccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenAuthenticationAccessPolicyOutput)
}





type OpenAuthenticationAccessPolicyMapInput interface {
	pulumi.Input

	ToOpenAuthenticationAccessPolicyMapOutput() OpenAuthenticationAccessPolicyMapOutput
	ToOpenAuthenticationAccessPolicyMapOutputWithContext(context.Context) OpenAuthenticationAccessPolicyMapOutput
}

type OpenAuthenticationAccessPolicyMap map[string]OpenAuthenticationAccessPolicyInput

func (OpenAuthenticationAccessPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OpenAuthenticationAccessPolicy)(nil)).Elem()
}

func (i OpenAuthenticationAccessPolicyMap) ToOpenAuthenticationAccessPolicyMapOutput() OpenAuthenticationAccessPolicyMapOutput {
	return i.ToOpenAuthenticationAccessPolicyMapOutputWithContext(context.Background())
}

func (i OpenAuthenticationAccessPolicyMap) ToOpenAuthenticationAccessPolicyMapOutputWithContext(ctx context.Context) OpenAuthenticationAccessPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenAuthenticationAccessPolicyMapOutput)
}

type OpenAuthenticationAccessPolicyOutput struct{ *pulumi.OutputState }

func (OpenAuthenticationAccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenAuthenticationAccessPolicy)(nil)).Elem()
}

func (o OpenAuthenticationAccessPolicyOutput) ToOpenAuthenticationAccessPolicyOutput() OpenAuthenticationAccessPolicyOutput {
	return o
}

func (o OpenAuthenticationAccessPolicyOutput) ToOpenAuthenticationAccessPolicyOutputWithContext(ctx context.Context) OpenAuthenticationAccessPolicyOutput {
	return o
}

func (o OpenAuthenticationAccessPolicyOutput) Claims() OpenAuthenticationPolicyClaimArrayOutput {
	return o.ApplyT(func(v OpenAuthenticationAccessPolicy) []OpenAuthenticationPolicyClaim { return v.Claims }).(OpenAuthenticationPolicyClaimArrayOutput)
}

func (o OpenAuthenticationAccessPolicyOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenAuthenticationAccessPolicy) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type OpenAuthenticationAccessPolicyMapOutput struct{ *pulumi.OutputState }

func (OpenAuthenticationAccessPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OpenAuthenticationAccessPolicy)(nil)).Elem()
}

func (o OpenAuthenticationAccessPolicyMapOutput) ToOpenAuthenticationAccessPolicyMapOutput() OpenAuthenticationAccessPolicyMapOutput {
	return o
}

func (o OpenAuthenticationAccessPolicyMapOutput) ToOpenAuthenticationAccessPolicyMapOutputWithContext(ctx context.Context) OpenAuthenticationAccessPolicyMapOutput {
	return o
}

func (o OpenAuthenticationAccessPolicyMapOutput) MapIndex(k pulumi.StringInput) OpenAuthenticationAccessPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OpenAuthenticationAccessPolicy {
		return vs[0].(map[string]OpenAuthenticationAccessPolicy)[vs[1].(string)]
	}).(OpenAuthenticationAccessPolicyOutput)
}

type OpenAuthenticationAccessPolicyResponse struct {
	Claims []OpenAuthenticationPolicyClaimResponse `pulumi:"claims"`
	Type   *string                                 `pulumi:"type"`
}

type OpenAuthenticationAccessPolicyResponseOutput struct{ *pulumi.OutputState }

func (OpenAuthenticationAccessPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenAuthenticationAccessPolicyResponse)(nil)).Elem()
}

func (o OpenAuthenticationAccessPolicyResponseOutput) ToOpenAuthenticationAccessPolicyResponseOutput() OpenAuthenticationAccessPolicyResponseOutput {
	return o
}

func (o OpenAuthenticationAccessPolicyResponseOutput) ToOpenAuthenticationAccessPolicyResponseOutputWithContext(ctx context.Context) OpenAuthenticationAccessPolicyResponseOutput {
	return o
}

func (o OpenAuthenticationAccessPolicyResponseOutput) Claims() OpenAuthenticationPolicyClaimResponseArrayOutput {
	return o.ApplyT(func(v OpenAuthenticationAccessPolicyResponse) []OpenAuthenticationPolicyClaimResponse {
		return v.Claims
	}).(OpenAuthenticationPolicyClaimResponseArrayOutput)
}

func (o OpenAuthenticationAccessPolicyResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenAuthenticationAccessPolicyResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type OpenAuthenticationAccessPolicyResponseMapOutput struct{ *pulumi.OutputState }

func (OpenAuthenticationAccessPolicyResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OpenAuthenticationAccessPolicyResponse)(nil)).Elem()
}

func (o OpenAuthenticationAccessPolicyResponseMapOutput) ToOpenAuthenticationAccessPolicyResponseMapOutput() OpenAuthenticationAccessPolicyResponseMapOutput {
	return o
}

func (o OpenAuthenticationAccessPolicyResponseMapOutput) ToOpenAuthenticationAccessPolicyResponseMapOutputWithContext(ctx context.Context) OpenAuthenticationAccessPolicyResponseMapOutput {
	return o
}

func (o OpenAuthenticationAccessPolicyResponseMapOutput) MapIndex(k pulumi.StringInput) OpenAuthenticationAccessPolicyResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OpenAuthenticationAccessPolicyResponse {
		return vs[0].(map[string]OpenAuthenticationAccessPolicyResponse)[vs[1].(string)]
	}).(OpenAuthenticationAccessPolicyResponseOutput)
}

type OpenAuthenticationPolicyClaim struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type OpenAuthenticationPolicyClaimInput interface {
	pulumi.Input

	ToOpenAuthenticationPolicyClaimOutput() OpenAuthenticationPolicyClaimOutput
	ToOpenAuthenticationPolicyClaimOutputWithContext(context.Context) OpenAuthenticationPolicyClaimOutput
}

type OpenAuthenticationPolicyClaimArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (OpenAuthenticationPolicyClaimArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenAuthenticationPolicyClaim)(nil)).Elem()
}

func (i OpenAuthenticationPolicyClaimArgs) ToOpenAuthenticationPolicyClaimOutput() OpenAuthenticationPolicyClaimOutput {
	return i.ToOpenAuthenticationPolicyClaimOutputWithContext(context.Background())
}

func (i OpenAuthenticationPolicyClaimArgs) ToOpenAuthenticationPolicyClaimOutputWithContext(ctx context.Context) OpenAuthenticationPolicyClaimOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenAuthenticationPolicyClaimOutput)
}





type OpenAuthenticationPolicyClaimArrayInput interface {
	pulumi.Input

	ToOpenAuthenticationPolicyClaimArrayOutput() OpenAuthenticationPolicyClaimArrayOutput
	ToOpenAuthenticationPolicyClaimArrayOutputWithContext(context.Context) OpenAuthenticationPolicyClaimArrayOutput
}

type OpenAuthenticationPolicyClaimArray []OpenAuthenticationPolicyClaimInput

func (OpenAuthenticationPolicyClaimArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OpenAuthenticationPolicyClaim)(nil)).Elem()
}

func (i OpenAuthenticationPolicyClaimArray) ToOpenAuthenticationPolicyClaimArrayOutput() OpenAuthenticationPolicyClaimArrayOutput {
	return i.ToOpenAuthenticationPolicyClaimArrayOutputWithContext(context.Background())
}

func (i OpenAuthenticationPolicyClaimArray) ToOpenAuthenticationPolicyClaimArrayOutputWithContext(ctx context.Context) OpenAuthenticationPolicyClaimArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenAuthenticationPolicyClaimArrayOutput)
}

type OpenAuthenticationPolicyClaimOutput struct{ *pulumi.OutputState }

func (OpenAuthenticationPolicyClaimOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenAuthenticationPolicyClaim)(nil)).Elem()
}

func (o OpenAuthenticationPolicyClaimOutput) ToOpenAuthenticationPolicyClaimOutput() OpenAuthenticationPolicyClaimOutput {
	return o
}

func (o OpenAuthenticationPolicyClaimOutput) ToOpenAuthenticationPolicyClaimOutputWithContext(ctx context.Context) OpenAuthenticationPolicyClaimOutput {
	return o
}

func (o OpenAuthenticationPolicyClaimOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenAuthenticationPolicyClaim) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OpenAuthenticationPolicyClaimOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenAuthenticationPolicyClaim) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type OpenAuthenticationPolicyClaimArrayOutput struct{ *pulumi.OutputState }

func (OpenAuthenticationPolicyClaimArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OpenAuthenticationPolicyClaim)(nil)).Elem()
}

func (o OpenAuthenticationPolicyClaimArrayOutput) ToOpenAuthenticationPolicyClaimArrayOutput() OpenAuthenticationPolicyClaimArrayOutput {
	return o
}

func (o OpenAuthenticationPolicyClaimArrayOutput) ToOpenAuthenticationPolicyClaimArrayOutputWithContext(ctx context.Context) OpenAuthenticationPolicyClaimArrayOutput {
	return o
}

func (o OpenAuthenticationPolicyClaimArrayOutput) Index(i pulumi.IntInput) OpenAuthenticationPolicyClaimOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OpenAuthenticationPolicyClaim {
		return vs[0].([]OpenAuthenticationPolicyClaim)[vs[1].(int)]
	}).(OpenAuthenticationPolicyClaimOutput)
}

type OpenAuthenticationPolicyClaimResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type OpenAuthenticationPolicyClaimResponseOutput struct{ *pulumi.OutputState }

func (OpenAuthenticationPolicyClaimResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenAuthenticationPolicyClaimResponse)(nil)).Elem()
}

func (o OpenAuthenticationPolicyClaimResponseOutput) ToOpenAuthenticationPolicyClaimResponseOutput() OpenAuthenticationPolicyClaimResponseOutput {
	return o
}

func (o OpenAuthenticationPolicyClaimResponseOutput) ToOpenAuthenticationPolicyClaimResponseOutputWithContext(ctx context.Context) OpenAuthenticationPolicyClaimResponseOutput {
	return o
}

func (o OpenAuthenticationPolicyClaimResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenAuthenticationPolicyClaimResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OpenAuthenticationPolicyClaimResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenAuthenticationPolicyClaimResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type OpenAuthenticationPolicyClaimResponseArrayOutput struct{ *pulumi.OutputState }

func (OpenAuthenticationPolicyClaimResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OpenAuthenticationPolicyClaimResponse)(nil)).Elem()
}

func (o OpenAuthenticationPolicyClaimResponseArrayOutput) ToOpenAuthenticationPolicyClaimResponseArrayOutput() OpenAuthenticationPolicyClaimResponseArrayOutput {
	return o
}

func (o OpenAuthenticationPolicyClaimResponseArrayOutput) ToOpenAuthenticationPolicyClaimResponseArrayOutputWithContext(ctx context.Context) OpenAuthenticationPolicyClaimResponseArrayOutput {
	return o
}

func (o OpenAuthenticationPolicyClaimResponseArrayOutput) Index(i pulumi.IntInput) OpenAuthenticationPolicyClaimResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OpenAuthenticationPolicyClaimResponse {
		return vs[0].([]OpenAuthenticationPolicyClaimResponse)[vs[1].(int)]
	}).(OpenAuthenticationPolicyClaimResponseOutput)
}

type PartnerContent struct {
	B2b *B2BPartnerContent `pulumi:"b2b"`
}





type PartnerContentInput interface {
	pulumi.Input

	ToPartnerContentOutput() PartnerContentOutput
	ToPartnerContentOutputWithContext(context.Context) PartnerContentOutput
}

type PartnerContentArgs struct {
	B2b B2BPartnerContentPtrInput `pulumi:"b2b"`
}

func (PartnerContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerContent)(nil)).Elem()
}

func (i PartnerContentArgs) ToPartnerContentOutput() PartnerContentOutput {
	return i.ToPartnerContentOutputWithContext(context.Background())
}

func (i PartnerContentArgs) ToPartnerContentOutputWithContext(ctx context.Context) PartnerContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerContentOutput)
}

type PartnerContentOutput struct{ *pulumi.OutputState }

func (PartnerContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerContent)(nil)).Elem()
}

func (o PartnerContentOutput) ToPartnerContentOutput() PartnerContentOutput {
	return o
}

func (o PartnerContentOutput) ToPartnerContentOutputWithContext(ctx context.Context) PartnerContentOutput {
	return o
}

func (o PartnerContentOutput) B2b() B2BPartnerContentPtrOutput {
	return o.ApplyT(func(v PartnerContent) *B2BPartnerContent { return v.B2b }).(B2BPartnerContentPtrOutput)
}

type PartnerContentResponse struct {
	B2b *B2BPartnerContentResponse `pulumi:"b2b"`
}

type PartnerContentResponseOutput struct{ *pulumi.OutputState }

func (PartnerContentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerContentResponse)(nil)).Elem()
}

func (o PartnerContentResponseOutput) ToPartnerContentResponseOutput() PartnerContentResponseOutput {
	return o
}

func (o PartnerContentResponseOutput) ToPartnerContentResponseOutputWithContext(ctx context.Context) PartnerContentResponseOutput {
	return o
}

func (o PartnerContentResponseOutput) B2b() B2BPartnerContentResponsePtrOutput {
	return o.ApplyT(func(v PartnerContentResponse) *B2BPartnerContentResponse { return v.B2b }).(B2BPartnerContentResponsePtrOutput)
}

type RecurrenceSchedule struct {
	Hours              []int                          `pulumi:"hours"`
	Minutes            []int                          `pulumi:"minutes"`
	MonthDays          []int                          `pulumi:"monthDays"`
	MonthlyOccurrences []RecurrenceScheduleOccurrence `pulumi:"monthlyOccurrences"`
	WeekDays           []DaysOfWeek                   `pulumi:"weekDays"`
}





type RecurrenceScheduleInput interface {
	pulumi.Input

	ToRecurrenceScheduleOutput() RecurrenceScheduleOutput
	ToRecurrenceScheduleOutputWithContext(context.Context) RecurrenceScheduleOutput
}

type RecurrenceScheduleArgs struct {
	Hours              pulumi.IntArrayInput                   `pulumi:"hours"`
	Minutes            pulumi.IntArrayInput                   `pulumi:"minutes"`
	MonthDays          pulumi.IntArrayInput                   `pulumi:"monthDays"`
	MonthlyOccurrences RecurrenceScheduleOccurrenceArrayInput `pulumi:"monthlyOccurrences"`
	WeekDays           DaysOfWeekArrayInput                   `pulumi:"weekDays"`
}

func (RecurrenceScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceSchedule)(nil)).Elem()
}

func (i RecurrenceScheduleArgs) ToRecurrenceScheduleOutput() RecurrenceScheduleOutput {
	return i.ToRecurrenceScheduleOutputWithContext(context.Background())
}

func (i RecurrenceScheduleArgs) ToRecurrenceScheduleOutputWithContext(ctx context.Context) RecurrenceScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrenceScheduleOutput)
}

func (i RecurrenceScheduleArgs) ToRecurrenceSchedulePtrOutput() RecurrenceSchedulePtrOutput {
	return i.ToRecurrenceSchedulePtrOutputWithContext(context.Background())
}

func (i RecurrenceScheduleArgs) ToRecurrenceSchedulePtrOutputWithContext(ctx context.Context) RecurrenceSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrenceScheduleOutput).ToRecurrenceSchedulePtrOutputWithContext(ctx)
}









type RecurrenceSchedulePtrInput interface {
	pulumi.Input

	ToRecurrenceSchedulePtrOutput() RecurrenceSchedulePtrOutput
	ToRecurrenceSchedulePtrOutputWithContext(context.Context) RecurrenceSchedulePtrOutput
}

type recurrenceSchedulePtrType RecurrenceScheduleArgs

func RecurrenceSchedulePtr(v *RecurrenceScheduleArgs) RecurrenceSchedulePtrInput {
	return (*recurrenceSchedulePtrType)(v)
}

func (*recurrenceSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RecurrenceSchedule)(nil)).Elem()
}

func (i *recurrenceSchedulePtrType) ToRecurrenceSchedulePtrOutput() RecurrenceSchedulePtrOutput {
	return i.ToRecurrenceSchedulePtrOutputWithContext(context.Background())
}

func (i *recurrenceSchedulePtrType) ToRecurrenceSchedulePtrOutputWithContext(ctx context.Context) RecurrenceSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrenceSchedulePtrOutput)
}

type RecurrenceScheduleOutput struct{ *pulumi.OutputState }

func (RecurrenceScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceSchedule)(nil)).Elem()
}

func (o RecurrenceScheduleOutput) ToRecurrenceScheduleOutput() RecurrenceScheduleOutput {
	return o
}

func (o RecurrenceScheduleOutput) ToRecurrenceScheduleOutputWithContext(ctx context.Context) RecurrenceScheduleOutput {
	return o
}

func (o RecurrenceScheduleOutput) ToRecurrenceSchedulePtrOutput() RecurrenceSchedulePtrOutput {
	return o.ToRecurrenceSchedulePtrOutputWithContext(context.Background())
}

func (o RecurrenceScheduleOutput) ToRecurrenceSchedulePtrOutputWithContext(ctx context.Context) RecurrenceSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecurrenceSchedule) *RecurrenceSchedule {
		return &v
	}).(RecurrenceSchedulePtrOutput)
}

func (o RecurrenceScheduleOutput) Hours() pulumi.IntArrayOutput {
	return o.ApplyT(func(v RecurrenceSchedule) []int { return v.Hours }).(pulumi.IntArrayOutput)
}

func (o RecurrenceScheduleOutput) Minutes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v RecurrenceSchedule) []int { return v.Minutes }).(pulumi.IntArrayOutput)
}

func (o RecurrenceScheduleOutput) MonthDays() pulumi.IntArrayOutput {
	return o.ApplyT(func(v RecurrenceSchedule) []int { return v.MonthDays }).(pulumi.IntArrayOutput)
}

func (o RecurrenceScheduleOutput) MonthlyOccurrences() RecurrenceScheduleOccurrenceArrayOutput {
	return o.ApplyT(func(v RecurrenceSchedule) []RecurrenceScheduleOccurrence { return v.MonthlyOccurrences }).(RecurrenceScheduleOccurrenceArrayOutput)
}

func (o RecurrenceScheduleOutput) WeekDays() DaysOfWeekArrayOutput {
	return o.ApplyT(func(v RecurrenceSchedule) []DaysOfWeek { return v.WeekDays }).(DaysOfWeekArrayOutput)
}

type RecurrenceSchedulePtrOutput struct{ *pulumi.OutputState }

func (RecurrenceSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecurrenceSchedule)(nil)).Elem()
}

func (o RecurrenceSchedulePtrOutput) ToRecurrenceSchedulePtrOutput() RecurrenceSchedulePtrOutput {
	return o
}

func (o RecurrenceSchedulePtrOutput) ToRecurrenceSchedulePtrOutputWithContext(ctx context.Context) RecurrenceSchedulePtrOutput {
	return o
}

func (o RecurrenceSchedulePtrOutput) Elem() RecurrenceScheduleOutput {
	return o.ApplyT(func(v *RecurrenceSchedule) RecurrenceSchedule {
		if v != nil {
			return *v
		}
		var ret RecurrenceSchedule
		return ret
	}).(RecurrenceScheduleOutput)
}

func (o RecurrenceSchedulePtrOutput) Hours() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *RecurrenceSchedule) []int {
		if v == nil {
			return nil
		}
		return v.Hours
	}).(pulumi.IntArrayOutput)
}

func (o RecurrenceSchedulePtrOutput) Minutes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *RecurrenceSchedule) []int {
		if v == nil {
			return nil
		}
		return v.Minutes
	}).(pulumi.IntArrayOutput)
}

func (o RecurrenceSchedulePtrOutput) MonthDays() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *RecurrenceSchedule) []int {
		if v == nil {
			return nil
		}
		return v.MonthDays
	}).(pulumi.IntArrayOutput)
}

func (o RecurrenceSchedulePtrOutput) MonthlyOccurrences() RecurrenceScheduleOccurrenceArrayOutput {
	return o.ApplyT(func(v *RecurrenceSchedule) []RecurrenceScheduleOccurrence {
		if v == nil {
			return nil
		}
		return v.MonthlyOccurrences
	}).(RecurrenceScheduleOccurrenceArrayOutput)
}

func (o RecurrenceSchedulePtrOutput) WeekDays() DaysOfWeekArrayOutput {
	return o.ApplyT(func(v *RecurrenceSchedule) []DaysOfWeek {
		if v == nil {
			return nil
		}
		return v.WeekDays
	}).(DaysOfWeekArrayOutput)
}

type RecurrenceScheduleOccurrence struct {
	Day        *DayOfWeek `pulumi:"day"`
	Occurrence *int       `pulumi:"occurrence"`
}





type RecurrenceScheduleOccurrenceInput interface {
	pulumi.Input

	ToRecurrenceScheduleOccurrenceOutput() RecurrenceScheduleOccurrenceOutput
	ToRecurrenceScheduleOccurrenceOutputWithContext(context.Context) RecurrenceScheduleOccurrenceOutput
}

type RecurrenceScheduleOccurrenceArgs struct {
	Day        DayOfWeekPtrInput  `pulumi:"day"`
	Occurrence pulumi.IntPtrInput `pulumi:"occurrence"`
}

func (RecurrenceScheduleOccurrenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceScheduleOccurrence)(nil)).Elem()
}

func (i RecurrenceScheduleOccurrenceArgs) ToRecurrenceScheduleOccurrenceOutput() RecurrenceScheduleOccurrenceOutput {
	return i.ToRecurrenceScheduleOccurrenceOutputWithContext(context.Background())
}

func (i RecurrenceScheduleOccurrenceArgs) ToRecurrenceScheduleOccurrenceOutputWithContext(ctx context.Context) RecurrenceScheduleOccurrenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrenceScheduleOccurrenceOutput)
}





type RecurrenceScheduleOccurrenceArrayInput interface {
	pulumi.Input

	ToRecurrenceScheduleOccurrenceArrayOutput() RecurrenceScheduleOccurrenceArrayOutput
	ToRecurrenceScheduleOccurrenceArrayOutputWithContext(context.Context) RecurrenceScheduleOccurrenceArrayOutput
}

type RecurrenceScheduleOccurrenceArray []RecurrenceScheduleOccurrenceInput

func (RecurrenceScheduleOccurrenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecurrenceScheduleOccurrence)(nil)).Elem()
}

func (i RecurrenceScheduleOccurrenceArray) ToRecurrenceScheduleOccurrenceArrayOutput() RecurrenceScheduleOccurrenceArrayOutput {
	return i.ToRecurrenceScheduleOccurrenceArrayOutputWithContext(context.Background())
}

func (i RecurrenceScheduleOccurrenceArray) ToRecurrenceScheduleOccurrenceArrayOutputWithContext(ctx context.Context) RecurrenceScheduleOccurrenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecurrenceScheduleOccurrenceArrayOutput)
}

type RecurrenceScheduleOccurrenceOutput struct{ *pulumi.OutputState }

func (RecurrenceScheduleOccurrenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceScheduleOccurrence)(nil)).Elem()
}

func (o RecurrenceScheduleOccurrenceOutput) ToRecurrenceScheduleOccurrenceOutput() RecurrenceScheduleOccurrenceOutput {
	return o
}

func (o RecurrenceScheduleOccurrenceOutput) ToRecurrenceScheduleOccurrenceOutputWithContext(ctx context.Context) RecurrenceScheduleOccurrenceOutput {
	return o
}

func (o RecurrenceScheduleOccurrenceOutput) Day() DayOfWeekPtrOutput {
	return o.ApplyT(func(v RecurrenceScheduleOccurrence) *DayOfWeek { return v.Day }).(DayOfWeekPtrOutput)
}

func (o RecurrenceScheduleOccurrenceOutput) Occurrence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RecurrenceScheduleOccurrence) *int { return v.Occurrence }).(pulumi.IntPtrOutput)
}

type RecurrenceScheduleOccurrenceArrayOutput struct{ *pulumi.OutputState }

func (RecurrenceScheduleOccurrenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecurrenceScheduleOccurrence)(nil)).Elem()
}

func (o RecurrenceScheduleOccurrenceArrayOutput) ToRecurrenceScheduleOccurrenceArrayOutput() RecurrenceScheduleOccurrenceArrayOutput {
	return o
}

func (o RecurrenceScheduleOccurrenceArrayOutput) ToRecurrenceScheduleOccurrenceArrayOutputWithContext(ctx context.Context) RecurrenceScheduleOccurrenceArrayOutput {
	return o
}

func (o RecurrenceScheduleOccurrenceArrayOutput) Index(i pulumi.IntInput) RecurrenceScheduleOccurrenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecurrenceScheduleOccurrence {
		return vs[0].([]RecurrenceScheduleOccurrence)[vs[1].(int)]
	}).(RecurrenceScheduleOccurrenceOutput)
}

type RecurrenceScheduleOccurrenceResponse struct {
	Day        *string `pulumi:"day"`
	Occurrence *int    `pulumi:"occurrence"`
}

type RecurrenceScheduleOccurrenceResponseOutput struct{ *pulumi.OutputState }

func (RecurrenceScheduleOccurrenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceScheduleOccurrenceResponse)(nil)).Elem()
}

func (o RecurrenceScheduleOccurrenceResponseOutput) ToRecurrenceScheduleOccurrenceResponseOutput() RecurrenceScheduleOccurrenceResponseOutput {
	return o
}

func (o RecurrenceScheduleOccurrenceResponseOutput) ToRecurrenceScheduleOccurrenceResponseOutputWithContext(ctx context.Context) RecurrenceScheduleOccurrenceResponseOutput {
	return o
}

func (o RecurrenceScheduleOccurrenceResponseOutput) Day() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecurrenceScheduleOccurrenceResponse) *string { return v.Day }).(pulumi.StringPtrOutput)
}

func (o RecurrenceScheduleOccurrenceResponseOutput) Occurrence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RecurrenceScheduleOccurrenceResponse) *int { return v.Occurrence }).(pulumi.IntPtrOutput)
}

type RecurrenceScheduleOccurrenceResponseArrayOutput struct{ *pulumi.OutputState }

func (RecurrenceScheduleOccurrenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecurrenceScheduleOccurrenceResponse)(nil)).Elem()
}

func (o RecurrenceScheduleOccurrenceResponseArrayOutput) ToRecurrenceScheduleOccurrenceResponseArrayOutput() RecurrenceScheduleOccurrenceResponseArrayOutput {
	return o
}

func (o RecurrenceScheduleOccurrenceResponseArrayOutput) ToRecurrenceScheduleOccurrenceResponseArrayOutputWithContext(ctx context.Context) RecurrenceScheduleOccurrenceResponseArrayOutput {
	return o
}

func (o RecurrenceScheduleOccurrenceResponseArrayOutput) Index(i pulumi.IntInput) RecurrenceScheduleOccurrenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecurrenceScheduleOccurrenceResponse {
		return vs[0].([]RecurrenceScheduleOccurrenceResponse)[vs[1].(int)]
	}).(RecurrenceScheduleOccurrenceResponseOutput)
}

type RecurrenceScheduleResponse struct {
	Hours              []int                                  `pulumi:"hours"`
	Minutes            []int                                  `pulumi:"minutes"`
	MonthDays          []int                                  `pulumi:"monthDays"`
	MonthlyOccurrences []RecurrenceScheduleOccurrenceResponse `pulumi:"monthlyOccurrences"`
	WeekDays           []string                               `pulumi:"weekDays"`
}

type RecurrenceScheduleResponseOutput struct{ *pulumi.OutputState }

func (RecurrenceScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceScheduleResponse)(nil)).Elem()
}

func (o RecurrenceScheduleResponseOutput) ToRecurrenceScheduleResponseOutput() RecurrenceScheduleResponseOutput {
	return o
}

func (o RecurrenceScheduleResponseOutput) ToRecurrenceScheduleResponseOutputWithContext(ctx context.Context) RecurrenceScheduleResponseOutput {
	return o
}

func (o RecurrenceScheduleResponseOutput) Hours() pulumi.IntArrayOutput {
	return o.ApplyT(func(v RecurrenceScheduleResponse) []int { return v.Hours }).(pulumi.IntArrayOutput)
}

func (o RecurrenceScheduleResponseOutput) Minutes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v RecurrenceScheduleResponse) []int { return v.Minutes }).(pulumi.IntArrayOutput)
}

func (o RecurrenceScheduleResponseOutput) MonthDays() pulumi.IntArrayOutput {
	return o.ApplyT(func(v RecurrenceScheduleResponse) []int { return v.MonthDays }).(pulumi.IntArrayOutput)
}

func (o RecurrenceScheduleResponseOutput) MonthlyOccurrences() RecurrenceScheduleOccurrenceResponseArrayOutput {
	return o.ApplyT(func(v RecurrenceScheduleResponse) []RecurrenceScheduleOccurrenceResponse { return v.MonthlyOccurrences }).(RecurrenceScheduleOccurrenceResponseArrayOutput)
}

func (o RecurrenceScheduleResponseOutput) WeekDays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RecurrenceScheduleResponse) []string { return v.WeekDays }).(pulumi.StringArrayOutput)
}

type RecurrenceScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (RecurrenceScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecurrenceScheduleResponse)(nil)).Elem()
}

func (o RecurrenceScheduleResponsePtrOutput) ToRecurrenceScheduleResponsePtrOutput() RecurrenceScheduleResponsePtrOutput {
	return o
}

func (o RecurrenceScheduleResponsePtrOutput) ToRecurrenceScheduleResponsePtrOutputWithContext(ctx context.Context) RecurrenceScheduleResponsePtrOutput {
	return o
}

func (o RecurrenceScheduleResponsePtrOutput) Elem() RecurrenceScheduleResponseOutput {
	return o.ApplyT(func(v *RecurrenceScheduleResponse) RecurrenceScheduleResponse {
		if v != nil {
			return *v
		}
		var ret RecurrenceScheduleResponse
		return ret
	}).(RecurrenceScheduleResponseOutput)
}

func (o RecurrenceScheduleResponsePtrOutput) Hours() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *RecurrenceScheduleResponse) []int {
		if v == nil {
			return nil
		}
		return v.Hours
	}).(pulumi.IntArrayOutput)
}

func (o RecurrenceScheduleResponsePtrOutput) Minutes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *RecurrenceScheduleResponse) []int {
		if v == nil {
			return nil
		}
		return v.Minutes
	}).(pulumi.IntArrayOutput)
}

func (o RecurrenceScheduleResponsePtrOutput) MonthDays() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *RecurrenceScheduleResponse) []int {
		if v == nil {
			return nil
		}
		return v.MonthDays
	}).(pulumi.IntArrayOutput)
}

func (o RecurrenceScheduleResponsePtrOutput) MonthlyOccurrences() RecurrenceScheduleOccurrenceResponseArrayOutput {
	return o.ApplyT(func(v *RecurrenceScheduleResponse) []RecurrenceScheduleOccurrenceResponse {
		if v == nil {
			return nil
		}
		return v.MonthlyOccurrences
	}).(RecurrenceScheduleOccurrenceResponseArrayOutput)
}

func (o RecurrenceScheduleResponsePtrOutput) WeekDays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RecurrenceScheduleResponse) []string {
		if v == nil {
			return nil
		}
		return v.WeekDays
	}).(pulumi.StringArrayOutput)
}

type ResourceReference struct {
	Id *string `pulumi:"id"`
}





type ResourceReferenceInput interface {
	pulumi.Input

	ToResourceReferenceOutput() ResourceReferenceOutput
	ToResourceReferenceOutputWithContext(context.Context) ResourceReferenceOutput
}

type ResourceReferenceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ResourceReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReference)(nil)).Elem()
}

func (i ResourceReferenceArgs) ToResourceReferenceOutput() ResourceReferenceOutput {
	return i.ToResourceReferenceOutputWithContext(context.Background())
}

func (i ResourceReferenceArgs) ToResourceReferenceOutputWithContext(ctx context.Context) ResourceReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceOutput)
}

func (i ResourceReferenceArgs) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return i.ToResourceReferencePtrOutputWithContext(context.Background())
}

func (i ResourceReferenceArgs) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceOutput).ToResourceReferencePtrOutputWithContext(ctx)
}









type ResourceReferencePtrInput interface {
	pulumi.Input

	ToResourceReferencePtrOutput() ResourceReferencePtrOutput
	ToResourceReferencePtrOutputWithContext(context.Context) ResourceReferencePtrOutput
}

type resourceReferencePtrType ResourceReferenceArgs

func ResourceReferencePtr(v *ResourceReferenceArgs) ResourceReferencePtrInput {
	return (*resourceReferencePtrType)(v)
}

func (*resourceReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceReference)(nil)).Elem()
}

func (i *resourceReferencePtrType) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return i.ToResourceReferencePtrOutputWithContext(context.Background())
}

func (i *resourceReferencePtrType) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferencePtrOutput)
}





type ResourceReferenceArrayInput interface {
	pulumi.Input

	ToResourceReferenceArrayOutput() ResourceReferenceArrayOutput
	ToResourceReferenceArrayOutputWithContext(context.Context) ResourceReferenceArrayOutput
}

type ResourceReferenceArray []ResourceReferenceInput

func (ResourceReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceReference)(nil)).Elem()
}

func (i ResourceReferenceArray) ToResourceReferenceArrayOutput() ResourceReferenceArrayOutput {
	return i.ToResourceReferenceArrayOutputWithContext(context.Background())
}

func (i ResourceReferenceArray) ToResourceReferenceArrayOutputWithContext(ctx context.Context) ResourceReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceArrayOutput)
}

type ResourceReferenceOutput struct{ *pulumi.OutputState }

func (ResourceReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReference)(nil)).Elem()
}

func (o ResourceReferenceOutput) ToResourceReferenceOutput() ResourceReferenceOutput {
	return o
}

func (o ResourceReferenceOutput) ToResourceReferenceOutputWithContext(ctx context.Context) ResourceReferenceOutput {
	return o
}

func (o ResourceReferenceOutput) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return o.ToResourceReferencePtrOutputWithContext(context.Background())
}

func (o ResourceReferenceOutput) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceReference) *ResourceReference {
		return &v
	}).(ResourceReferencePtrOutput)
}

func (o ResourceReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ResourceReferencePtrOutput struct{ *pulumi.OutputState }

func (ResourceReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceReference)(nil)).Elem()
}

func (o ResourceReferencePtrOutput) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return o
}

func (o ResourceReferencePtrOutput) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return o
}

func (o ResourceReferencePtrOutput) Elem() ResourceReferenceOutput {
	return o.ApplyT(func(v *ResourceReference) ResourceReference {
		if v != nil {
			return *v
		}
		var ret ResourceReference
		return ret
	}).(ResourceReferenceOutput)
}

func (o ResourceReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceReference) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ResourceReferenceArrayOutput struct{ *pulumi.OutputState }

func (ResourceReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceReference)(nil)).Elem()
}

func (o ResourceReferenceArrayOutput) ToResourceReferenceArrayOutput() ResourceReferenceArrayOutput {
	return o
}

func (o ResourceReferenceArrayOutput) ToResourceReferenceArrayOutputWithContext(ctx context.Context) ResourceReferenceArrayOutput {
	return o
}

func (o ResourceReferenceArrayOutput) Index(i pulumi.IntInput) ResourceReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceReference {
		return vs[0].([]ResourceReference)[vs[1].(int)]
	}).(ResourceReferenceOutput)
}

type ResourceReferenceResponse struct {
	Id   *string `pulumi:"id"`
	Name string  `pulumi:"name"`
	Type string  `pulumi:"type"`
}

type ResourceReferenceResponseOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutputWithContext(ctx context.Context) ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ResourceReferenceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceReferenceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceReferenceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceReferenceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ResourceReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponsePtrOutput) ToResourceReferenceResponsePtrOutput() ResourceReferenceResponsePtrOutput {
	return o
}

func (o ResourceReferenceResponsePtrOutput) ToResourceReferenceResponsePtrOutputWithContext(ctx context.Context) ResourceReferenceResponsePtrOutput {
	return o
}

func (o ResourceReferenceResponsePtrOutput) Elem() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v *ResourceReferenceResponse) ResourceReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ResourceReferenceResponse
		return ret
	}).(ResourceReferenceResponseOutput)
}

func (o ResourceReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ResourceReferenceResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceReferenceResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ResourceReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponseArrayOutput) ToResourceReferenceResponseArrayOutput() ResourceReferenceResponseArrayOutput {
	return o
}

func (o ResourceReferenceResponseArrayOutput) ToResourceReferenceResponseArrayOutputWithContext(ctx context.Context) ResourceReferenceResponseArrayOutput {
	return o
}

func (o ResourceReferenceResponseArrayOutput) Index(i pulumi.IntInput) ResourceReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceReferenceResponse {
		return vs[0].([]ResourceReferenceResponse)[vs[1].(int)]
	}).(ResourceReferenceResponseOutput)
}

type RosettaNetPipAcknowledgmentOfReceiptSettings struct {
	IsNonRepudiationRequired   bool `pulumi:"isNonRepudiationRequired"`
	TimeToAcknowledgeInSeconds int  `pulumi:"timeToAcknowledgeInSeconds"`
}





type RosettaNetPipAcknowledgmentOfReceiptSettingsInput interface {
	pulumi.Input

	ToRosettaNetPipAcknowledgmentOfReceiptSettingsOutput() RosettaNetPipAcknowledgmentOfReceiptSettingsOutput
	ToRosettaNetPipAcknowledgmentOfReceiptSettingsOutputWithContext(context.Context) RosettaNetPipAcknowledgmentOfReceiptSettingsOutput
}

type RosettaNetPipAcknowledgmentOfReceiptSettingsArgs struct {
	IsNonRepudiationRequired   pulumi.BoolInput `pulumi:"isNonRepudiationRequired"`
	TimeToAcknowledgeInSeconds pulumi.IntInput  `pulumi:"timeToAcknowledgeInSeconds"`
}

func (RosettaNetPipAcknowledgmentOfReceiptSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RosettaNetPipAcknowledgmentOfReceiptSettings)(nil)).Elem()
}

func (i RosettaNetPipAcknowledgmentOfReceiptSettingsArgs) ToRosettaNetPipAcknowledgmentOfReceiptSettingsOutput() RosettaNetPipAcknowledgmentOfReceiptSettingsOutput {
	return i.ToRosettaNetPipAcknowledgmentOfReceiptSettingsOutputWithContext(context.Background())
}

func (i RosettaNetPipAcknowledgmentOfReceiptSettingsArgs) ToRosettaNetPipAcknowledgmentOfReceiptSettingsOutputWithContext(ctx context.Context) RosettaNetPipAcknowledgmentOfReceiptSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RosettaNetPipAcknowledgmentOfReceiptSettingsOutput)
}

type RosettaNetPipAcknowledgmentOfReceiptSettingsOutput struct{ *pulumi.OutputState }

func (RosettaNetPipAcknowledgmentOfReceiptSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RosettaNetPipAcknowledgmentOfReceiptSettings)(nil)).Elem()
}

func (o RosettaNetPipAcknowledgmentOfReceiptSettingsOutput) ToRosettaNetPipAcknowledgmentOfReceiptSettingsOutput() RosettaNetPipAcknowledgmentOfReceiptSettingsOutput {
	return o
}

func (o RosettaNetPipAcknowledgmentOfReceiptSettingsOutput) ToRosettaNetPipAcknowledgmentOfReceiptSettingsOutputWithContext(ctx context.Context) RosettaNetPipAcknowledgmentOfReceiptSettingsOutput {
	return o
}

func (o RosettaNetPipAcknowledgmentOfReceiptSettingsOutput) IsNonRepudiationRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v RosettaNetPipAcknowledgmentOfReceiptSettings) bool { return v.IsNonRepudiationRequired }).(pulumi.BoolOutput)
}

func (o RosettaNetPipAcknowledgmentOfReceiptSettingsOutput) TimeToAcknowledgeInSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v RosettaNetPipAcknowledgmentOfReceiptSettings) int { return v.TimeToAcknowledgeInSeconds }).(pulumi.IntOutput)
}

type RosettaNetPipAcknowledgmentOfReceiptSettingsResponse struct {
	IsNonRepudiationRequired   bool `pulumi:"isNonRepudiationRequired"`
	TimeToAcknowledgeInSeconds int  `pulumi:"timeToAcknowledgeInSeconds"`
}

type RosettaNetPipAcknowledgmentOfReceiptSettingsResponseOutput struct{ *pulumi.OutputState }

func (RosettaNetPipAcknowledgmentOfReceiptSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RosettaNetPipAcknowledgmentOfReceiptSettingsResponse)(nil)).Elem()
}

func (o RosettaNetPipAcknowledgmentOfReceiptSettingsResponseOutput) ToRosettaNetPipAcknowledgmentOfReceiptSettingsResponseOutput() RosettaNetPipAcknowledgmentOfReceiptSettingsResponseOutput {
	return o
}

func (o RosettaNetPipAcknowledgmentOfReceiptSettingsResponseOutput) ToRosettaNetPipAcknowledgmentOfReceiptSettingsResponseOutputWithContext(ctx context.Context) RosettaNetPipAcknowledgmentOfReceiptSettingsResponseOutput {
	return o
}

func (o RosettaNetPipAcknowledgmentOfReceiptSettingsResponseOutput) IsNonRepudiationRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v RosettaNetPipAcknowledgmentOfReceiptSettingsResponse) bool { return v.IsNonRepudiationRequired }).(pulumi.BoolOutput)
}

func (o RosettaNetPipAcknowledgmentOfReceiptSettingsResponseOutput) TimeToAcknowledgeInSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v RosettaNetPipAcknowledgmentOfReceiptSettingsResponse) int { return v.TimeToAcknowledgeInSeconds }).(pulumi.IntOutput)
}

type RosettaNetPipActivityBehavior struct {
	ActionType                       RosettaNetActionType              `pulumi:"actionType"`
	IsAuthorizationRequired          bool                              `pulumi:"isAuthorizationRequired"`
	IsSecuredTransportRequired       bool                              `pulumi:"isSecuredTransportRequired"`
	NonRepudiationOfOriginAndContent bool                              `pulumi:"nonRepudiationOfOriginAndContent"`
	PersistentConfidentialityScope   RosettaNetPipConfidentialityScope `pulumi:"persistentConfidentialityScope"`
	ResponseType                     RosettaNetResponseType            `pulumi:"responseType"`
	RetryCount                       int                               `pulumi:"retryCount"`
	TimeToPerformInSeconds           int                               `pulumi:"timeToPerformInSeconds"`
}





type RosettaNetPipActivityBehaviorInput interface {
	pulumi.Input

	ToRosettaNetPipActivityBehaviorOutput() RosettaNetPipActivityBehaviorOutput
	ToRosettaNetPipActivityBehaviorOutputWithContext(context.Context) RosettaNetPipActivityBehaviorOutput
}

type RosettaNetPipActivityBehaviorArgs struct {
	ActionType                       RosettaNetActionTypeInput              `pulumi:"actionType"`
	IsAuthorizationRequired          pulumi.BoolInput                       `pulumi:"isAuthorizationRequired"`
	IsSecuredTransportRequired       pulumi.BoolInput                       `pulumi:"isSecuredTransportRequired"`
	NonRepudiationOfOriginAndContent pulumi.BoolInput                       `pulumi:"nonRepudiationOfOriginAndContent"`
	PersistentConfidentialityScope   RosettaNetPipConfidentialityScopeInput `pulumi:"persistentConfidentialityScope"`
	ResponseType                     RosettaNetResponseTypeInput            `pulumi:"responseType"`
	RetryCount                       pulumi.IntInput                        `pulumi:"retryCount"`
	TimeToPerformInSeconds           pulumi.IntInput                        `pulumi:"timeToPerformInSeconds"`
}

func (RosettaNetPipActivityBehaviorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RosettaNetPipActivityBehavior)(nil)).Elem()
}

func (i RosettaNetPipActivityBehaviorArgs) ToRosettaNetPipActivityBehaviorOutput() RosettaNetPipActivityBehaviorOutput {
	return i.ToRosettaNetPipActivityBehaviorOutputWithContext(context.Background())
}

func (i RosettaNetPipActivityBehaviorArgs) ToRosettaNetPipActivityBehaviorOutputWithContext(ctx context.Context) RosettaNetPipActivityBehaviorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RosettaNetPipActivityBehaviorOutput)
}

type RosettaNetPipActivityBehaviorOutput struct{ *pulumi.OutputState }

func (RosettaNetPipActivityBehaviorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RosettaNetPipActivityBehavior)(nil)).Elem()
}

func (o RosettaNetPipActivityBehaviorOutput) ToRosettaNetPipActivityBehaviorOutput() RosettaNetPipActivityBehaviorOutput {
	return o
}

func (o RosettaNetPipActivityBehaviorOutput) ToRosettaNetPipActivityBehaviorOutputWithContext(ctx context.Context) RosettaNetPipActivityBehaviorOutput {
	return o
}

func (o RosettaNetPipActivityBehaviorOutput) ActionType() RosettaNetActionTypeOutput {
	return o.ApplyT(func(v RosettaNetPipActivityBehavior) RosettaNetActionType { return v.ActionType }).(RosettaNetActionTypeOutput)
}

func (o RosettaNetPipActivityBehaviorOutput) IsAuthorizationRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v RosettaNetPipActivityBehavior) bool { return v.IsAuthorizationRequired }).(pulumi.BoolOutput)
}

func (o RosettaNetPipActivityBehaviorOutput) IsSecuredTransportRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v RosettaNetPipActivityBehavior) bool { return v.IsSecuredTransportRequired }).(pulumi.BoolOutput)
}

func (o RosettaNetPipActivityBehaviorOutput) NonRepudiationOfOriginAndContent() pulumi.BoolOutput {
	return o.ApplyT(func(v RosettaNetPipActivityBehavior) bool { return v.NonRepudiationOfOriginAndContent }).(pulumi.BoolOutput)
}

func (o RosettaNetPipActivityBehaviorOutput) PersistentConfidentialityScope() RosettaNetPipConfidentialityScopeOutput {
	return o.ApplyT(func(v RosettaNetPipActivityBehavior) RosettaNetPipConfidentialityScope {
		return v.PersistentConfidentialityScope
	}).(RosettaNetPipConfidentialityScopeOutput)
}

func (o RosettaNetPipActivityBehaviorOutput) ResponseType() RosettaNetResponseTypeOutput {
	return o.ApplyT(func(v RosettaNetPipActivityBehavior) RosettaNetResponseType { return v.ResponseType }).(RosettaNetResponseTypeOutput)
}

func (o RosettaNetPipActivityBehaviorOutput) RetryCount() pulumi.IntOutput {
	return o.ApplyT(func(v RosettaNetPipActivityBehavior) int { return v.RetryCount }).(pulumi.IntOutput)
}

func (o RosettaNetPipActivityBehaviorOutput) TimeToPerformInSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v RosettaNetPipActivityBehavior) int { return v.TimeToPerformInSeconds }).(pulumi.IntOutput)
}

type RosettaNetPipActivityBehaviorResponse struct {
	ActionType                       string `pulumi:"actionType"`
	IsAuthorizationRequired          bool   `pulumi:"isAuthorizationRequired"`
	IsSecuredTransportRequired       bool   `pulumi:"isSecuredTransportRequired"`
	NonRepudiationOfOriginAndContent bool   `pulumi:"nonRepudiationOfOriginAndContent"`
	PersistentConfidentialityScope   string `pulumi:"persistentConfidentialityScope"`
	ResponseType                     string `pulumi:"responseType"`
	RetryCount                       int    `pulumi:"retryCount"`
	TimeToPerformInSeconds           int    `pulumi:"timeToPerformInSeconds"`
}

type RosettaNetPipActivityBehaviorResponseOutput struct{ *pulumi.OutputState }

func (RosettaNetPipActivityBehaviorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RosettaNetPipActivityBehaviorResponse)(nil)).Elem()
}

func (o RosettaNetPipActivityBehaviorResponseOutput) ToRosettaNetPipActivityBehaviorResponseOutput() RosettaNetPipActivityBehaviorResponseOutput {
	return o
}

func (o RosettaNetPipActivityBehaviorResponseOutput) ToRosettaNetPipActivityBehaviorResponseOutputWithContext(ctx context.Context) RosettaNetPipActivityBehaviorResponseOutput {
	return o
}

func (o RosettaNetPipActivityBehaviorResponseOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v RosettaNetPipActivityBehaviorResponse) string { return v.ActionType }).(pulumi.StringOutput)
}

func (o RosettaNetPipActivityBehaviorResponseOutput) IsAuthorizationRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v RosettaNetPipActivityBehaviorResponse) bool { return v.IsAuthorizationRequired }).(pulumi.BoolOutput)
}

func (o RosettaNetPipActivityBehaviorResponseOutput) IsSecuredTransportRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v RosettaNetPipActivityBehaviorResponse) bool { return v.IsSecuredTransportRequired }).(pulumi.BoolOutput)
}

func (o RosettaNetPipActivityBehaviorResponseOutput) NonRepudiationOfOriginAndContent() pulumi.BoolOutput {
	return o.ApplyT(func(v RosettaNetPipActivityBehaviorResponse) bool { return v.NonRepudiationOfOriginAndContent }).(pulumi.BoolOutput)
}

func (o RosettaNetPipActivityBehaviorResponseOutput) PersistentConfidentialityScope() pulumi.StringOutput {
	return o.ApplyT(func(v RosettaNetPipActivityBehaviorResponse) string { return v.PersistentConfidentialityScope }).(pulumi.StringOutput)
}

func (o RosettaNetPipActivityBehaviorResponseOutput) ResponseType() pulumi.StringOutput {
	return o.ApplyT(func(v RosettaNetPipActivityBehaviorResponse) string { return v.ResponseType }).(pulumi.StringOutput)
}

func (o RosettaNetPipActivityBehaviorResponseOutput) RetryCount() pulumi.IntOutput {
	return o.ApplyT(func(v RosettaNetPipActivityBehaviorResponse) int { return v.RetryCount }).(pulumi.IntOutput)
}

func (o RosettaNetPipActivityBehaviorResponseOutput) TimeToPerformInSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v RosettaNetPipActivityBehaviorResponse) int { return v.TimeToPerformInSeconds }).(pulumi.IntOutput)
}

type RosettaNetPipActivitySettings struct {
	AcknowledgmentOfReceiptSettings RosettaNetPipAcknowledgmentOfReceiptSettings `pulumi:"acknowledgmentOfReceiptSettings"`
	ActivityBehavior                RosettaNetPipActivityBehavior                `pulumi:"activityBehavior"`
	ActivityType                    RosettaNetPipActivityType                    `pulumi:"activityType"`
}





type RosettaNetPipActivitySettingsInput interface {
	pulumi.Input

	ToRosettaNetPipActivitySettingsOutput() RosettaNetPipActivitySettingsOutput
	ToRosettaNetPipActivitySettingsOutputWithContext(context.Context) RosettaNetPipActivitySettingsOutput
}

type RosettaNetPipActivitySettingsArgs struct {
	AcknowledgmentOfReceiptSettings RosettaNetPipAcknowledgmentOfReceiptSettingsInput `pulumi:"acknowledgmentOfReceiptSettings"`
	ActivityBehavior                RosettaNetPipActivityBehaviorInput                `pulumi:"activityBehavior"`
	ActivityType                    RosettaNetPipActivityTypeInput                    `pulumi:"activityType"`
}

func (RosettaNetPipActivitySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RosettaNetPipActivitySettings)(nil)).Elem()
}

func (i RosettaNetPipActivitySettingsArgs) ToRosettaNetPipActivitySettingsOutput() RosettaNetPipActivitySettingsOutput {
	return i.ToRosettaNetPipActivitySettingsOutputWithContext(context.Background())
}

func (i RosettaNetPipActivitySettingsArgs) ToRosettaNetPipActivitySettingsOutputWithContext(ctx context.Context) RosettaNetPipActivitySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RosettaNetPipActivitySettingsOutput)
}

type RosettaNetPipActivitySettingsOutput struct{ *pulumi.OutputState }

func (RosettaNetPipActivitySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RosettaNetPipActivitySettings)(nil)).Elem()
}

func (o RosettaNetPipActivitySettingsOutput) ToRosettaNetPipActivitySettingsOutput() RosettaNetPipActivitySettingsOutput {
	return o
}

func (o RosettaNetPipActivitySettingsOutput) ToRosettaNetPipActivitySettingsOutputWithContext(ctx context.Context) RosettaNetPipActivitySettingsOutput {
	return o
}

func (o RosettaNetPipActivitySettingsOutput) AcknowledgmentOfReceiptSettings() RosettaNetPipAcknowledgmentOfReceiptSettingsOutput {
	return o.ApplyT(func(v RosettaNetPipActivitySettings) RosettaNetPipAcknowledgmentOfReceiptSettings {
		return v.AcknowledgmentOfReceiptSettings
	}).(RosettaNetPipAcknowledgmentOfReceiptSettingsOutput)
}

func (o RosettaNetPipActivitySettingsOutput) ActivityBehavior() RosettaNetPipActivityBehaviorOutput {
	return o.ApplyT(func(v RosettaNetPipActivitySettings) RosettaNetPipActivityBehavior { return v.ActivityBehavior }).(RosettaNetPipActivityBehaviorOutput)
}

func (o RosettaNetPipActivitySettingsOutput) ActivityType() RosettaNetPipActivityTypeOutput {
	return o.ApplyT(func(v RosettaNetPipActivitySettings) RosettaNetPipActivityType { return v.ActivityType }).(RosettaNetPipActivityTypeOutput)
}

type RosettaNetPipActivitySettingsResponse struct {
	AcknowledgmentOfReceiptSettings RosettaNetPipAcknowledgmentOfReceiptSettingsResponse `pulumi:"acknowledgmentOfReceiptSettings"`
	ActivityBehavior                RosettaNetPipActivityBehaviorResponse                `pulumi:"activityBehavior"`
	ActivityType                    string                                               `pulumi:"activityType"`
}

type RosettaNetPipActivitySettingsResponseOutput struct{ *pulumi.OutputState }

func (RosettaNetPipActivitySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RosettaNetPipActivitySettingsResponse)(nil)).Elem()
}

func (o RosettaNetPipActivitySettingsResponseOutput) ToRosettaNetPipActivitySettingsResponseOutput() RosettaNetPipActivitySettingsResponseOutput {
	return o
}

func (o RosettaNetPipActivitySettingsResponseOutput) ToRosettaNetPipActivitySettingsResponseOutputWithContext(ctx context.Context) RosettaNetPipActivitySettingsResponseOutput {
	return o
}

func (o RosettaNetPipActivitySettingsResponseOutput) AcknowledgmentOfReceiptSettings() RosettaNetPipAcknowledgmentOfReceiptSettingsResponseOutput {
	return o.ApplyT(func(v RosettaNetPipActivitySettingsResponse) RosettaNetPipAcknowledgmentOfReceiptSettingsResponse {
		return v.AcknowledgmentOfReceiptSettings
	}).(RosettaNetPipAcknowledgmentOfReceiptSettingsResponseOutput)
}

func (o RosettaNetPipActivitySettingsResponseOutput) ActivityBehavior() RosettaNetPipActivityBehaviorResponseOutput {
	return o.ApplyT(func(v RosettaNetPipActivitySettingsResponse) RosettaNetPipActivityBehaviorResponse {
		return v.ActivityBehavior
	}).(RosettaNetPipActivityBehaviorResponseOutput)
}

func (o RosettaNetPipActivitySettingsResponseOutput) ActivityType() pulumi.StringOutput {
	return o.ApplyT(func(v RosettaNetPipActivitySettingsResponse) string { return v.ActivityType }).(pulumi.StringOutput)
}

type RosettaNetPipBusinessDocument struct {
	Description *string `pulumi:"description"`
	Name        string  `pulumi:"name"`
	Version     string  `pulumi:"version"`
}





type RosettaNetPipBusinessDocumentInput interface {
	pulumi.Input

	ToRosettaNetPipBusinessDocumentOutput() RosettaNetPipBusinessDocumentOutput
	ToRosettaNetPipBusinessDocumentOutputWithContext(context.Context) RosettaNetPipBusinessDocumentOutput
}

type RosettaNetPipBusinessDocumentArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Name        pulumi.StringInput    `pulumi:"name"`
	Version     pulumi.StringInput    `pulumi:"version"`
}

func (RosettaNetPipBusinessDocumentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RosettaNetPipBusinessDocument)(nil)).Elem()
}

func (i RosettaNetPipBusinessDocumentArgs) ToRosettaNetPipBusinessDocumentOutput() RosettaNetPipBusinessDocumentOutput {
	return i.ToRosettaNetPipBusinessDocumentOutputWithContext(context.Background())
}

func (i RosettaNetPipBusinessDocumentArgs) ToRosettaNetPipBusinessDocumentOutputWithContext(ctx context.Context) RosettaNetPipBusinessDocumentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RosettaNetPipBusinessDocumentOutput)
}

type RosettaNetPipBusinessDocumentOutput struct{ *pulumi.OutputState }

func (RosettaNetPipBusinessDocumentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RosettaNetPipBusinessDocument)(nil)).Elem()
}

func (o RosettaNetPipBusinessDocumentOutput) ToRosettaNetPipBusinessDocumentOutput() RosettaNetPipBusinessDocumentOutput {
	return o
}

func (o RosettaNetPipBusinessDocumentOutput) ToRosettaNetPipBusinessDocumentOutputWithContext(ctx context.Context) RosettaNetPipBusinessDocumentOutput {
	return o
}

func (o RosettaNetPipBusinessDocumentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RosettaNetPipBusinessDocument) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RosettaNetPipBusinessDocumentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RosettaNetPipBusinessDocument) string { return v.Name }).(pulumi.StringOutput)
}

func (o RosettaNetPipBusinessDocumentOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v RosettaNetPipBusinessDocument) string { return v.Version }).(pulumi.StringOutput)
}

type RosettaNetPipBusinessDocumentResponse struct {
	Description *string `pulumi:"description"`
	Name        string  `pulumi:"name"`
	Version     string  `pulumi:"version"`
}

type RosettaNetPipBusinessDocumentResponseOutput struct{ *pulumi.OutputState }

func (RosettaNetPipBusinessDocumentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RosettaNetPipBusinessDocumentResponse)(nil)).Elem()
}

func (o RosettaNetPipBusinessDocumentResponseOutput) ToRosettaNetPipBusinessDocumentResponseOutput() RosettaNetPipBusinessDocumentResponseOutput {
	return o
}

func (o RosettaNetPipBusinessDocumentResponseOutput) ToRosettaNetPipBusinessDocumentResponseOutputWithContext(ctx context.Context) RosettaNetPipBusinessDocumentResponseOutput {
	return o
}

func (o RosettaNetPipBusinessDocumentResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RosettaNetPipBusinessDocumentResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RosettaNetPipBusinessDocumentResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RosettaNetPipBusinessDocumentResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RosettaNetPipBusinessDocumentResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v RosettaNetPipBusinessDocumentResponse) string { return v.Version }).(pulumi.StringOutput)
}

type RosettaNetPipRoleSettings struct {
	Action                string                        `pulumi:"action"`
	BusinessDocument      RosettaNetPipBusinessDocument `pulumi:"businessDocument"`
	Description           *string                       `pulumi:"description"`
	Role                  string                        `pulumi:"role"`
	RoleType              RosettaNetPipRoleType         `pulumi:"roleType"`
	Service               string                        `pulumi:"service"`
	ServiceClassification string                        `pulumi:"serviceClassification"`
}





type RosettaNetPipRoleSettingsInput interface {
	pulumi.Input

	ToRosettaNetPipRoleSettingsOutput() RosettaNetPipRoleSettingsOutput
	ToRosettaNetPipRoleSettingsOutputWithContext(context.Context) RosettaNetPipRoleSettingsOutput
}

type RosettaNetPipRoleSettingsArgs struct {
	Action                pulumi.StringInput                 `pulumi:"action"`
	BusinessDocument      RosettaNetPipBusinessDocumentInput `pulumi:"businessDocument"`
	Description           pulumi.StringPtrInput              `pulumi:"description"`
	Role                  pulumi.StringInput                 `pulumi:"role"`
	RoleType              RosettaNetPipRoleTypeInput         `pulumi:"roleType"`
	Service               pulumi.StringInput                 `pulumi:"service"`
	ServiceClassification pulumi.StringInput                 `pulumi:"serviceClassification"`
}

func (RosettaNetPipRoleSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RosettaNetPipRoleSettings)(nil)).Elem()
}

func (i RosettaNetPipRoleSettingsArgs) ToRosettaNetPipRoleSettingsOutput() RosettaNetPipRoleSettingsOutput {
	return i.ToRosettaNetPipRoleSettingsOutputWithContext(context.Background())
}

func (i RosettaNetPipRoleSettingsArgs) ToRosettaNetPipRoleSettingsOutputWithContext(ctx context.Context) RosettaNetPipRoleSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RosettaNetPipRoleSettingsOutput)
}

type RosettaNetPipRoleSettingsOutput struct{ *pulumi.OutputState }

func (RosettaNetPipRoleSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RosettaNetPipRoleSettings)(nil)).Elem()
}

func (o RosettaNetPipRoleSettingsOutput) ToRosettaNetPipRoleSettingsOutput() RosettaNetPipRoleSettingsOutput {
	return o
}

func (o RosettaNetPipRoleSettingsOutput) ToRosettaNetPipRoleSettingsOutputWithContext(ctx context.Context) RosettaNetPipRoleSettingsOutput {
	return o
}

func (o RosettaNetPipRoleSettingsOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v RosettaNetPipRoleSettings) string { return v.Action }).(pulumi.StringOutput)
}

func (o RosettaNetPipRoleSettingsOutput) BusinessDocument() RosettaNetPipBusinessDocumentOutput {
	return o.ApplyT(func(v RosettaNetPipRoleSettings) RosettaNetPipBusinessDocument { return v.BusinessDocument }).(RosettaNetPipBusinessDocumentOutput)
}

func (o RosettaNetPipRoleSettingsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RosettaNetPipRoleSettings) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RosettaNetPipRoleSettingsOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v RosettaNetPipRoleSettings) string { return v.Role }).(pulumi.StringOutput)
}

func (o RosettaNetPipRoleSettingsOutput) RoleType() RosettaNetPipRoleTypeOutput {
	return o.ApplyT(func(v RosettaNetPipRoleSettings) RosettaNetPipRoleType { return v.RoleType }).(RosettaNetPipRoleTypeOutput)
}

func (o RosettaNetPipRoleSettingsOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v RosettaNetPipRoleSettings) string { return v.Service }).(pulumi.StringOutput)
}

func (o RosettaNetPipRoleSettingsOutput) ServiceClassification() pulumi.StringOutput {
	return o.ApplyT(func(v RosettaNetPipRoleSettings) string { return v.ServiceClassification }).(pulumi.StringOutput)
}

type RosettaNetPipRoleSettingsResponse struct {
	Action                string                                `pulumi:"action"`
	BusinessDocument      RosettaNetPipBusinessDocumentResponse `pulumi:"businessDocument"`
	Description           *string                               `pulumi:"description"`
	Role                  string                                `pulumi:"role"`
	RoleType              string                                `pulumi:"roleType"`
	Service               string                                `pulumi:"service"`
	ServiceClassification string                                `pulumi:"serviceClassification"`
}

type RosettaNetPipRoleSettingsResponseOutput struct{ *pulumi.OutputState }

func (RosettaNetPipRoleSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RosettaNetPipRoleSettingsResponse)(nil)).Elem()
}

func (o RosettaNetPipRoleSettingsResponseOutput) ToRosettaNetPipRoleSettingsResponseOutput() RosettaNetPipRoleSettingsResponseOutput {
	return o
}

func (o RosettaNetPipRoleSettingsResponseOutput) ToRosettaNetPipRoleSettingsResponseOutputWithContext(ctx context.Context) RosettaNetPipRoleSettingsResponseOutput {
	return o
}

func (o RosettaNetPipRoleSettingsResponseOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v RosettaNetPipRoleSettingsResponse) string { return v.Action }).(pulumi.StringOutput)
}

func (o RosettaNetPipRoleSettingsResponseOutput) BusinessDocument() RosettaNetPipBusinessDocumentResponseOutput {
	return o.ApplyT(func(v RosettaNetPipRoleSettingsResponse) RosettaNetPipBusinessDocumentResponse {
		return v.BusinessDocument
	}).(RosettaNetPipBusinessDocumentResponseOutput)
}

func (o RosettaNetPipRoleSettingsResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RosettaNetPipRoleSettingsResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RosettaNetPipRoleSettingsResponseOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v RosettaNetPipRoleSettingsResponse) string { return v.Role }).(pulumi.StringOutput)
}

func (o RosettaNetPipRoleSettingsResponseOutput) RoleType() pulumi.StringOutput {
	return o.ApplyT(func(v RosettaNetPipRoleSettingsResponse) string { return v.RoleType }).(pulumi.StringOutput)
}

func (o RosettaNetPipRoleSettingsResponseOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v RosettaNetPipRoleSettingsResponse) string { return v.Service }).(pulumi.StringOutput)
}

func (o RosettaNetPipRoleSettingsResponseOutput) ServiceClassification() pulumi.StringOutput {
	return o.ApplyT(func(v RosettaNetPipRoleSettingsResponse) string { return v.ServiceClassification }).(pulumi.StringOutput)
}

type SkuResponse struct {
	Name string                     `pulumi:"name"`
	Plan *ResourceReferenceResponse `pulumi:"plan"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Plan() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v SkuResponse) *ResourceReferenceResponse { return v.Plan }).(ResourceReferenceResponsePtrOutput)
}

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

type WorkflowParameter struct {
	Description *string     `pulumi:"description"`
	Metadata    interface{} `pulumi:"metadata"`
	Type        *string     `pulumi:"type"`
	Value       interface{} `pulumi:"value"`
}





type WorkflowParameterInput interface {
	pulumi.Input

	ToWorkflowParameterOutput() WorkflowParameterOutput
	ToWorkflowParameterOutputWithContext(context.Context) WorkflowParameterOutput
}

type WorkflowParameterArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Metadata    pulumi.Input          `pulumi:"metadata"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
	Value       pulumi.Input          `pulumi:"value"`
}

func (WorkflowParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowParameter)(nil)).Elem()
}

func (i WorkflowParameterArgs) ToWorkflowParameterOutput() WorkflowParameterOutput {
	return i.ToWorkflowParameterOutputWithContext(context.Background())
}

func (i WorkflowParameterArgs) ToWorkflowParameterOutputWithContext(ctx context.Context) WorkflowParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowParameterOutput)
}





type WorkflowParameterMapInput interface {
	pulumi.Input

	ToWorkflowParameterMapOutput() WorkflowParameterMapOutput
	ToWorkflowParameterMapOutputWithContext(context.Context) WorkflowParameterMapOutput
}

type WorkflowParameterMap map[string]WorkflowParameterInput

func (WorkflowParameterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkflowParameter)(nil)).Elem()
}

func (i WorkflowParameterMap) ToWorkflowParameterMapOutput() WorkflowParameterMapOutput {
	return i.ToWorkflowParameterMapOutputWithContext(context.Background())
}

func (i WorkflowParameterMap) ToWorkflowParameterMapOutputWithContext(ctx context.Context) WorkflowParameterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowParameterMapOutput)
}

type WorkflowParameterOutput struct{ *pulumi.OutputState }

func (WorkflowParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowParameter)(nil)).Elem()
}

func (o WorkflowParameterOutput) ToWorkflowParameterOutput() WorkflowParameterOutput {
	return o
}

func (o WorkflowParameterOutput) ToWorkflowParameterOutputWithContext(ctx context.Context) WorkflowParameterOutput {
	return o
}

func (o WorkflowParameterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowParameter) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WorkflowParameterOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkflowParameter) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o WorkflowParameterOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowParameter) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o WorkflowParameterOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkflowParameter) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type WorkflowParameterMapOutput struct{ *pulumi.OutputState }

func (WorkflowParameterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkflowParameter)(nil)).Elem()
}

func (o WorkflowParameterMapOutput) ToWorkflowParameterMapOutput() WorkflowParameterMapOutput {
	return o
}

func (o WorkflowParameterMapOutput) ToWorkflowParameterMapOutputWithContext(ctx context.Context) WorkflowParameterMapOutput {
	return o
}

func (o WorkflowParameterMapOutput) MapIndex(k pulumi.StringInput) WorkflowParameterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WorkflowParameter {
		return vs[0].(map[string]WorkflowParameter)[vs[1].(string)]
	}).(WorkflowParameterOutput)
}

type WorkflowParameterResponse struct {
	Description *string     `pulumi:"description"`
	Metadata    interface{} `pulumi:"metadata"`
	Type        *string     `pulumi:"type"`
	Value       interface{} `pulumi:"value"`
}

type WorkflowParameterResponseOutput struct{ *pulumi.OutputState }

func (WorkflowParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowParameterResponse)(nil)).Elem()
}

func (o WorkflowParameterResponseOutput) ToWorkflowParameterResponseOutput() WorkflowParameterResponseOutput {
	return o
}

func (o WorkflowParameterResponseOutput) ToWorkflowParameterResponseOutputWithContext(ctx context.Context) WorkflowParameterResponseOutput {
	return o
}

func (o WorkflowParameterResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowParameterResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WorkflowParameterResponseOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkflowParameterResponse) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o WorkflowParameterResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowParameterResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o WorkflowParameterResponseOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkflowParameterResponse) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type WorkflowParameterResponseMapOutput struct{ *pulumi.OutputState }

func (WorkflowParameterResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkflowParameterResponse)(nil)).Elem()
}

func (o WorkflowParameterResponseMapOutput) ToWorkflowParameterResponseMapOutput() WorkflowParameterResponseMapOutput {
	return o
}

func (o WorkflowParameterResponseMapOutput) ToWorkflowParameterResponseMapOutputWithContext(ctx context.Context) WorkflowParameterResponseMapOutput {
	return o
}

func (o WorkflowParameterResponseMapOutput) MapIndex(k pulumi.StringInput) WorkflowParameterResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WorkflowParameterResponse {
		return vs[0].(map[string]WorkflowParameterResponse)[vs[1].(string)]
	}).(WorkflowParameterResponseOutput)
}

type WorkflowTriggerListCallbackUrlQueriesResponse struct {
	ApiVersion *string `pulumi:"apiVersion"`
	Se         *string `pulumi:"se"`
	Sig        *string `pulumi:"sig"`
	Sp         *string `pulumi:"sp"`
	Sv         *string `pulumi:"sv"`
}

type WorkflowTriggerListCallbackUrlQueriesResponseOutput struct{ *pulumi.OutputState }

func (WorkflowTriggerListCallbackUrlQueriesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowTriggerListCallbackUrlQueriesResponse)(nil)).Elem()
}

func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) ToWorkflowTriggerListCallbackUrlQueriesResponseOutput() WorkflowTriggerListCallbackUrlQueriesResponseOutput {
	return o
}

func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) ToWorkflowTriggerListCallbackUrlQueriesResponseOutputWithContext(ctx context.Context) WorkflowTriggerListCallbackUrlQueriesResponseOutput {
	return o
}

func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerListCallbackUrlQueriesResponse) *string { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) Se() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerListCallbackUrlQueriesResponse) *string { return v.Se }).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) Sig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerListCallbackUrlQueriesResponse) *string { return v.Sig }).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) Sp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerListCallbackUrlQueriesResponse) *string { return v.Sp }).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) Sv() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerListCallbackUrlQueriesResponse) *string { return v.Sv }).(pulumi.StringPtrOutput)
}

type WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTriggerListCallbackUrlQueriesResponse)(nil)).Elem()
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) ToWorkflowTriggerListCallbackUrlQueriesResponsePtrOutput() WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) ToWorkflowTriggerListCallbackUrlQueriesResponsePtrOutputWithContext(ctx context.Context) WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) Elem() WorkflowTriggerListCallbackUrlQueriesResponseOutput {
	return o.ApplyT(func(v *WorkflowTriggerListCallbackUrlQueriesResponse) WorkflowTriggerListCallbackUrlQueriesResponse {
		if v != nil {
			return *v
		}
		var ret WorkflowTriggerListCallbackUrlQueriesResponse
		return ret
	}).(WorkflowTriggerListCallbackUrlQueriesResponseOutput)
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerListCallbackUrlQueriesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApiVersion
	}).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) Se() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerListCallbackUrlQueriesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Se
	}).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) Sig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerListCallbackUrlQueriesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sig
	}).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) Sp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerListCallbackUrlQueriesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sp
	}).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) Sv() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerListCallbackUrlQueriesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sv
	}).(pulumi.StringPtrOutput)
}

type WorkflowTriggerRecurrence struct {
	EndTime   *string             `pulumi:"endTime"`
	Frequency *string             `pulumi:"frequency"`
	Interval  *int                `pulumi:"interval"`
	Schedule  *RecurrenceSchedule `pulumi:"schedule"`
	StartTime *string             `pulumi:"startTime"`
	TimeZone  *string             `pulumi:"timeZone"`
}





type WorkflowTriggerRecurrenceInput interface {
	pulumi.Input

	ToWorkflowTriggerRecurrenceOutput() WorkflowTriggerRecurrenceOutput
	ToWorkflowTriggerRecurrenceOutputWithContext(context.Context) WorkflowTriggerRecurrenceOutput
}

type WorkflowTriggerRecurrenceArgs struct {
	EndTime   pulumi.StringPtrInput      `pulumi:"endTime"`
	Frequency pulumi.StringPtrInput      `pulumi:"frequency"`
	Interval  pulumi.IntPtrInput         `pulumi:"interval"`
	Schedule  RecurrenceSchedulePtrInput `pulumi:"schedule"`
	StartTime pulumi.StringPtrInput      `pulumi:"startTime"`
	TimeZone  pulumi.StringPtrInput      `pulumi:"timeZone"`
}

func (WorkflowTriggerRecurrenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowTriggerRecurrence)(nil)).Elem()
}

func (i WorkflowTriggerRecurrenceArgs) ToWorkflowTriggerRecurrenceOutput() WorkflowTriggerRecurrenceOutput {
	return i.ToWorkflowTriggerRecurrenceOutputWithContext(context.Background())
}

func (i WorkflowTriggerRecurrenceArgs) ToWorkflowTriggerRecurrenceOutputWithContext(ctx context.Context) WorkflowTriggerRecurrenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTriggerRecurrenceOutput)
}

func (i WorkflowTriggerRecurrenceArgs) ToWorkflowTriggerRecurrencePtrOutput() WorkflowTriggerRecurrencePtrOutput {
	return i.ToWorkflowTriggerRecurrencePtrOutputWithContext(context.Background())
}

func (i WorkflowTriggerRecurrenceArgs) ToWorkflowTriggerRecurrencePtrOutputWithContext(ctx context.Context) WorkflowTriggerRecurrencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTriggerRecurrenceOutput).ToWorkflowTriggerRecurrencePtrOutputWithContext(ctx)
}









type WorkflowTriggerRecurrencePtrInput interface {
	pulumi.Input

	ToWorkflowTriggerRecurrencePtrOutput() WorkflowTriggerRecurrencePtrOutput
	ToWorkflowTriggerRecurrencePtrOutputWithContext(context.Context) WorkflowTriggerRecurrencePtrOutput
}

type workflowTriggerRecurrencePtrType WorkflowTriggerRecurrenceArgs

func WorkflowTriggerRecurrencePtr(v *WorkflowTriggerRecurrenceArgs) WorkflowTriggerRecurrencePtrInput {
	return (*workflowTriggerRecurrencePtrType)(v)
}

func (*workflowTriggerRecurrencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTriggerRecurrence)(nil)).Elem()
}

func (i *workflowTriggerRecurrencePtrType) ToWorkflowTriggerRecurrencePtrOutput() WorkflowTriggerRecurrencePtrOutput {
	return i.ToWorkflowTriggerRecurrencePtrOutputWithContext(context.Background())
}

func (i *workflowTriggerRecurrencePtrType) ToWorkflowTriggerRecurrencePtrOutputWithContext(ctx context.Context) WorkflowTriggerRecurrencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTriggerRecurrencePtrOutput)
}

type WorkflowTriggerRecurrenceOutput struct{ *pulumi.OutputState }

func (WorkflowTriggerRecurrenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowTriggerRecurrence)(nil)).Elem()
}

func (o WorkflowTriggerRecurrenceOutput) ToWorkflowTriggerRecurrenceOutput() WorkflowTriggerRecurrenceOutput {
	return o
}

func (o WorkflowTriggerRecurrenceOutput) ToWorkflowTriggerRecurrenceOutputWithContext(ctx context.Context) WorkflowTriggerRecurrenceOutput {
	return o
}

func (o WorkflowTriggerRecurrenceOutput) ToWorkflowTriggerRecurrencePtrOutput() WorkflowTriggerRecurrencePtrOutput {
	return o.ToWorkflowTriggerRecurrencePtrOutputWithContext(context.Background())
}

func (o WorkflowTriggerRecurrenceOutput) ToWorkflowTriggerRecurrencePtrOutputWithContext(ctx context.Context) WorkflowTriggerRecurrencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkflowTriggerRecurrence) *WorkflowTriggerRecurrence {
		return &v
	}).(WorkflowTriggerRecurrencePtrOutput)
}

func (o WorkflowTriggerRecurrenceOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerRecurrence) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerRecurrenceOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerRecurrence) *string { return v.Frequency }).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerRecurrenceOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerRecurrence) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o WorkflowTriggerRecurrenceOutput) Schedule() RecurrenceSchedulePtrOutput {
	return o.ApplyT(func(v WorkflowTriggerRecurrence) *RecurrenceSchedule { return v.Schedule }).(RecurrenceSchedulePtrOutput)
}

func (o WorkflowTriggerRecurrenceOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerRecurrence) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerRecurrenceOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerRecurrence) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type WorkflowTriggerRecurrencePtrOutput struct{ *pulumi.OutputState }

func (WorkflowTriggerRecurrencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTriggerRecurrence)(nil)).Elem()
}

func (o WorkflowTriggerRecurrencePtrOutput) ToWorkflowTriggerRecurrencePtrOutput() WorkflowTriggerRecurrencePtrOutput {
	return o
}

func (o WorkflowTriggerRecurrencePtrOutput) ToWorkflowTriggerRecurrencePtrOutputWithContext(ctx context.Context) WorkflowTriggerRecurrencePtrOutput {
	return o
}

func (o WorkflowTriggerRecurrencePtrOutput) Elem() WorkflowTriggerRecurrenceOutput {
	return o.ApplyT(func(v *WorkflowTriggerRecurrence) WorkflowTriggerRecurrence {
		if v != nil {
			return *v
		}
		var ret WorkflowTriggerRecurrence
		return ret
	}).(WorkflowTriggerRecurrenceOutput)
}

func (o WorkflowTriggerRecurrencePtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerRecurrence) *string {
		if v == nil {
			return nil
		}
		return v.EndTime
	}).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerRecurrencePtrOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerRecurrence) *string {
		if v == nil {
			return nil
		}
		return v.Frequency
	}).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerRecurrencePtrOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerRecurrence) *int {
		if v == nil {
			return nil
		}
		return v.Interval
	}).(pulumi.IntPtrOutput)
}

func (o WorkflowTriggerRecurrencePtrOutput) Schedule() RecurrenceSchedulePtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerRecurrence) *RecurrenceSchedule {
		if v == nil {
			return nil
		}
		return v.Schedule
	}).(RecurrenceSchedulePtrOutput)
}

func (o WorkflowTriggerRecurrencePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerRecurrence) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerRecurrencePtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerRecurrence) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

type WorkflowTriggerRecurrenceResponse struct {
	EndTime   *string                     `pulumi:"endTime"`
	Frequency *string                     `pulumi:"frequency"`
	Interval  *int                        `pulumi:"interval"`
	Schedule  *RecurrenceScheduleResponse `pulumi:"schedule"`
	StartTime *string                     `pulumi:"startTime"`
	TimeZone  *string                     `pulumi:"timeZone"`
}

type WorkflowTriggerRecurrenceResponseOutput struct{ *pulumi.OutputState }

func (WorkflowTriggerRecurrenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowTriggerRecurrenceResponse)(nil)).Elem()
}

func (o WorkflowTriggerRecurrenceResponseOutput) ToWorkflowTriggerRecurrenceResponseOutput() WorkflowTriggerRecurrenceResponseOutput {
	return o
}

func (o WorkflowTriggerRecurrenceResponseOutput) ToWorkflowTriggerRecurrenceResponseOutputWithContext(ctx context.Context) WorkflowTriggerRecurrenceResponseOutput {
	return o
}

func (o WorkflowTriggerRecurrenceResponseOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerRecurrenceResponse) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerRecurrenceResponseOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerRecurrenceResponse) *string { return v.Frequency }).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerRecurrenceResponseOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerRecurrenceResponse) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o WorkflowTriggerRecurrenceResponseOutput) Schedule() RecurrenceScheduleResponsePtrOutput {
	return o.ApplyT(func(v WorkflowTriggerRecurrenceResponse) *RecurrenceScheduleResponse { return v.Schedule }).(RecurrenceScheduleResponsePtrOutput)
}

func (o WorkflowTriggerRecurrenceResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerRecurrenceResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerRecurrenceResponseOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerRecurrenceResponse) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type WorkflowTriggerRecurrenceResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkflowTriggerRecurrenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTriggerRecurrenceResponse)(nil)).Elem()
}

func (o WorkflowTriggerRecurrenceResponsePtrOutput) ToWorkflowTriggerRecurrenceResponsePtrOutput() WorkflowTriggerRecurrenceResponsePtrOutput {
	return o
}

func (o WorkflowTriggerRecurrenceResponsePtrOutput) ToWorkflowTriggerRecurrenceResponsePtrOutputWithContext(ctx context.Context) WorkflowTriggerRecurrenceResponsePtrOutput {
	return o
}

func (o WorkflowTriggerRecurrenceResponsePtrOutput) Elem() WorkflowTriggerRecurrenceResponseOutput {
	return o.ApplyT(func(v *WorkflowTriggerRecurrenceResponse) WorkflowTriggerRecurrenceResponse {
		if v != nil {
			return *v
		}
		var ret WorkflowTriggerRecurrenceResponse
		return ret
	}).(WorkflowTriggerRecurrenceResponseOutput)
}

func (o WorkflowTriggerRecurrenceResponsePtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerRecurrenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndTime
	}).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerRecurrenceResponsePtrOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerRecurrenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Frequency
	}).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerRecurrenceResponsePtrOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerRecurrenceResponse) *int {
		if v == nil {
			return nil
		}
		return v.Interval
	}).(pulumi.IntPtrOutput)
}

func (o WorkflowTriggerRecurrenceResponsePtrOutput) Schedule() RecurrenceScheduleResponsePtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerRecurrenceResponse) *RecurrenceScheduleResponse {
		if v == nil {
			return nil
		}
		return v.Schedule
	}).(RecurrenceScheduleResponsePtrOutput)
}

func (o WorkflowTriggerRecurrenceResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerRecurrenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerRecurrenceResponsePtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerRecurrenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

type WsdlServiceResponse struct {
	EndpointQualifiedNames []string `pulumi:"endpointQualifiedNames"`
	QualifiedName          *string  `pulumi:"qualifiedName"`
}

type WsdlServiceResponseOutput struct{ *pulumi.OutputState }

func (WsdlServiceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WsdlServiceResponse)(nil)).Elem()
}

func (o WsdlServiceResponseOutput) ToWsdlServiceResponseOutput() WsdlServiceResponseOutput {
	return o
}

func (o WsdlServiceResponseOutput) ToWsdlServiceResponseOutputWithContext(ctx context.Context) WsdlServiceResponseOutput {
	return o
}

func (o WsdlServiceResponseOutput) EndpointQualifiedNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WsdlServiceResponse) []string { return v.EndpointQualifiedNames }).(pulumi.StringArrayOutput)
}

func (o WsdlServiceResponseOutput) QualifiedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsdlServiceResponse) *string { return v.QualifiedName }).(pulumi.StringPtrOutput)
}

type WsdlServiceResponsePtrOutput struct{ *pulumi.OutputState }

func (WsdlServiceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WsdlServiceResponse)(nil)).Elem()
}

func (o WsdlServiceResponsePtrOutput) ToWsdlServiceResponsePtrOutput() WsdlServiceResponsePtrOutput {
	return o
}

func (o WsdlServiceResponsePtrOutput) ToWsdlServiceResponsePtrOutputWithContext(ctx context.Context) WsdlServiceResponsePtrOutput {
	return o
}

func (o WsdlServiceResponsePtrOutput) Elem() WsdlServiceResponseOutput {
	return o.ApplyT(func(v *WsdlServiceResponse) WsdlServiceResponse {
		if v != nil {
			return *v
		}
		var ret WsdlServiceResponse
		return ret
	}).(WsdlServiceResponseOutput)
}

func (o WsdlServiceResponsePtrOutput) EndpointQualifiedNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WsdlServiceResponse) []string {
		if v == nil {
			return nil
		}
		return v.EndpointQualifiedNames
	}).(pulumi.StringArrayOutput)
}

func (o WsdlServiceResponsePtrOutput) QualifiedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsdlServiceResponse) *string {
		if v == nil {
			return nil
		}
		return v.QualifiedName
	}).(pulumi.StringPtrOutput)
}

type X12AcknowledgementSettings struct {
	AcknowledgementControlNumberLowerBound int     `pulumi:"acknowledgementControlNumberLowerBound"`
	AcknowledgementControlNumberPrefix     *string `pulumi:"acknowledgementControlNumberPrefix"`
	AcknowledgementControlNumberSuffix     *string `pulumi:"acknowledgementControlNumberSuffix"`
	AcknowledgementControlNumberUpperBound int     `pulumi:"acknowledgementControlNumberUpperBound"`
	BatchFunctionalAcknowledgements        bool    `pulumi:"batchFunctionalAcknowledgements"`
	BatchImplementationAcknowledgements    bool    `pulumi:"batchImplementationAcknowledgements"`
	BatchTechnicalAcknowledgements         bool    `pulumi:"batchTechnicalAcknowledgements"`
	FunctionalAcknowledgementVersion       *string `pulumi:"functionalAcknowledgementVersion"`
	ImplementationAcknowledgementVersion   *string `pulumi:"implementationAcknowledgementVersion"`
	NeedFunctionalAcknowledgement          bool    `pulumi:"needFunctionalAcknowledgement"`
	NeedImplementationAcknowledgement      bool    `pulumi:"needImplementationAcknowledgement"`
	NeedLoopForValidMessages               bool    `pulumi:"needLoopForValidMessages"`
	NeedTechnicalAcknowledgement           bool    `pulumi:"needTechnicalAcknowledgement"`
	RolloverAcknowledgementControlNumber   bool    `pulumi:"rolloverAcknowledgementControlNumber"`
	SendSynchronousAcknowledgement         bool    `pulumi:"sendSynchronousAcknowledgement"`
}





type X12AcknowledgementSettingsInput interface {
	pulumi.Input

	ToX12AcknowledgementSettingsOutput() X12AcknowledgementSettingsOutput
	ToX12AcknowledgementSettingsOutputWithContext(context.Context) X12AcknowledgementSettingsOutput
}

type X12AcknowledgementSettingsArgs struct {
	AcknowledgementControlNumberLowerBound pulumi.IntInput       `pulumi:"acknowledgementControlNumberLowerBound"`
	AcknowledgementControlNumberPrefix     pulumi.StringPtrInput `pulumi:"acknowledgementControlNumberPrefix"`
	AcknowledgementControlNumberSuffix     pulumi.StringPtrInput `pulumi:"acknowledgementControlNumberSuffix"`
	AcknowledgementControlNumberUpperBound pulumi.IntInput       `pulumi:"acknowledgementControlNumberUpperBound"`
	BatchFunctionalAcknowledgements        pulumi.BoolInput      `pulumi:"batchFunctionalAcknowledgements"`
	BatchImplementationAcknowledgements    pulumi.BoolInput      `pulumi:"batchImplementationAcknowledgements"`
	BatchTechnicalAcknowledgements         pulumi.BoolInput      `pulumi:"batchTechnicalAcknowledgements"`
	FunctionalAcknowledgementVersion       pulumi.StringPtrInput `pulumi:"functionalAcknowledgementVersion"`
	ImplementationAcknowledgementVersion   pulumi.StringPtrInput `pulumi:"implementationAcknowledgementVersion"`
	NeedFunctionalAcknowledgement          pulumi.BoolInput      `pulumi:"needFunctionalAcknowledgement"`
	NeedImplementationAcknowledgement      pulumi.BoolInput      `pulumi:"needImplementationAcknowledgement"`
	NeedLoopForValidMessages               pulumi.BoolInput      `pulumi:"needLoopForValidMessages"`
	NeedTechnicalAcknowledgement           pulumi.BoolInput      `pulumi:"needTechnicalAcknowledgement"`
	RolloverAcknowledgementControlNumber   pulumi.BoolInput      `pulumi:"rolloverAcknowledgementControlNumber"`
	SendSynchronousAcknowledgement         pulumi.BoolInput      `pulumi:"sendSynchronousAcknowledgement"`
}

func (X12AcknowledgementSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12AcknowledgementSettings)(nil)).Elem()
}

func (i X12AcknowledgementSettingsArgs) ToX12AcknowledgementSettingsOutput() X12AcknowledgementSettingsOutput {
	return i.ToX12AcknowledgementSettingsOutputWithContext(context.Background())
}

func (i X12AcknowledgementSettingsArgs) ToX12AcknowledgementSettingsOutputWithContext(ctx context.Context) X12AcknowledgementSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12AcknowledgementSettingsOutput)
}

func (i X12AcknowledgementSettingsArgs) ToX12AcknowledgementSettingsPtrOutput() X12AcknowledgementSettingsPtrOutput {
	return i.ToX12AcknowledgementSettingsPtrOutputWithContext(context.Background())
}

func (i X12AcknowledgementSettingsArgs) ToX12AcknowledgementSettingsPtrOutputWithContext(ctx context.Context) X12AcknowledgementSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12AcknowledgementSettingsOutput).ToX12AcknowledgementSettingsPtrOutputWithContext(ctx)
}









type X12AcknowledgementSettingsPtrInput interface {
	pulumi.Input

	ToX12AcknowledgementSettingsPtrOutput() X12AcknowledgementSettingsPtrOutput
	ToX12AcknowledgementSettingsPtrOutputWithContext(context.Context) X12AcknowledgementSettingsPtrOutput
}

type x12acknowledgementSettingsPtrType X12AcknowledgementSettingsArgs

func X12AcknowledgementSettingsPtr(v *X12AcknowledgementSettingsArgs) X12AcknowledgementSettingsPtrInput {
	return (*x12acknowledgementSettingsPtrType)(v)
}

func (*x12acknowledgementSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12AcknowledgementSettings)(nil)).Elem()
}

func (i *x12acknowledgementSettingsPtrType) ToX12AcknowledgementSettingsPtrOutput() X12AcknowledgementSettingsPtrOutput {
	return i.ToX12AcknowledgementSettingsPtrOutputWithContext(context.Background())
}

func (i *x12acknowledgementSettingsPtrType) ToX12AcknowledgementSettingsPtrOutputWithContext(ctx context.Context) X12AcknowledgementSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12AcknowledgementSettingsPtrOutput)
}

type X12AcknowledgementSettingsOutput struct{ *pulumi.OutputState }

func (X12AcknowledgementSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12AcknowledgementSettings)(nil)).Elem()
}

func (o X12AcknowledgementSettingsOutput) ToX12AcknowledgementSettingsOutput() X12AcknowledgementSettingsOutput {
	return o
}

func (o X12AcknowledgementSettingsOutput) ToX12AcknowledgementSettingsOutputWithContext(ctx context.Context) X12AcknowledgementSettingsOutput {
	return o
}

func (o X12AcknowledgementSettingsOutput) ToX12AcknowledgementSettingsPtrOutput() X12AcknowledgementSettingsPtrOutput {
	return o.ToX12AcknowledgementSettingsPtrOutputWithContext(context.Background())
}

func (o X12AcknowledgementSettingsOutput) ToX12AcknowledgementSettingsPtrOutputWithContext(ctx context.Context) X12AcknowledgementSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12AcknowledgementSettings) *X12AcknowledgementSettings {
		return &v
	}).(X12AcknowledgementSettingsPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) AcknowledgementControlNumberLowerBound() pulumi.IntOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) int { return v.AcknowledgementControlNumberLowerBound }).(pulumi.IntOutput)
}

func (o X12AcknowledgementSettingsOutput) AcknowledgementControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *string { return v.AcknowledgementControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) AcknowledgementControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *string { return v.AcknowledgementControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) AcknowledgementControlNumberUpperBound() pulumi.IntOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) int { return v.AcknowledgementControlNumberUpperBound }).(pulumi.IntOutput)
}

func (o X12AcknowledgementSettingsOutput) BatchFunctionalAcknowledgements() pulumi.BoolOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) bool { return v.BatchFunctionalAcknowledgements }).(pulumi.BoolOutput)
}

func (o X12AcknowledgementSettingsOutput) BatchImplementationAcknowledgements() pulumi.BoolOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) bool { return v.BatchImplementationAcknowledgements }).(pulumi.BoolOutput)
}

func (o X12AcknowledgementSettingsOutput) BatchTechnicalAcknowledgements() pulumi.BoolOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) bool { return v.BatchTechnicalAcknowledgements }).(pulumi.BoolOutput)
}

func (o X12AcknowledgementSettingsOutput) FunctionalAcknowledgementVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *string { return v.FunctionalAcknowledgementVersion }).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) ImplementationAcknowledgementVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) *string { return v.ImplementationAcknowledgementVersion }).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsOutput) NeedFunctionalAcknowledgement() pulumi.BoolOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) bool { return v.NeedFunctionalAcknowledgement }).(pulumi.BoolOutput)
}

func (o X12AcknowledgementSettingsOutput) NeedImplementationAcknowledgement() pulumi.BoolOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) bool { return v.NeedImplementationAcknowledgement }).(pulumi.BoolOutput)
}

func (o X12AcknowledgementSettingsOutput) NeedLoopForValidMessages() pulumi.BoolOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) bool { return v.NeedLoopForValidMessages }).(pulumi.BoolOutput)
}

func (o X12AcknowledgementSettingsOutput) NeedTechnicalAcknowledgement() pulumi.BoolOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) bool { return v.NeedTechnicalAcknowledgement }).(pulumi.BoolOutput)
}

func (o X12AcknowledgementSettingsOutput) RolloverAcknowledgementControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) bool { return v.RolloverAcknowledgementControlNumber }).(pulumi.BoolOutput)
}

func (o X12AcknowledgementSettingsOutput) SendSynchronousAcknowledgement() pulumi.BoolOutput {
	return o.ApplyT(func(v X12AcknowledgementSettings) bool { return v.SendSynchronousAcknowledgement }).(pulumi.BoolOutput)
}

type X12AcknowledgementSettingsPtrOutput struct{ *pulumi.OutputState }

func (X12AcknowledgementSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12AcknowledgementSettings)(nil)).Elem()
}

func (o X12AcknowledgementSettingsPtrOutput) ToX12AcknowledgementSettingsPtrOutput() X12AcknowledgementSettingsPtrOutput {
	return o
}

func (o X12AcknowledgementSettingsPtrOutput) ToX12AcknowledgementSettingsPtrOutputWithContext(ctx context.Context) X12AcknowledgementSettingsPtrOutput {
	return o
}

func (o X12AcknowledgementSettingsPtrOutput) Elem() X12AcknowledgementSettingsOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) X12AcknowledgementSettings {
		if v != nil {
			return *v
		}
		var ret X12AcknowledgementSettings
		return ret
	}).(X12AcknowledgementSettingsOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) AcknowledgementControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *int {
		if v == nil {
			return nil
		}
		return &v.AcknowledgementControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) AcknowledgementControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *string {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) AcknowledgementControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *string {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) AcknowledgementControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *int {
		if v == nil {
			return nil
		}
		return &v.AcknowledgementControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) BatchFunctionalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.BatchFunctionalAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) BatchImplementationAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.BatchImplementationAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) BatchTechnicalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.BatchTechnicalAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) FunctionalAcknowledgementVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *string {
		if v == nil {
			return nil
		}
		return v.FunctionalAcknowledgementVersion
	}).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) ImplementationAcknowledgementVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *string {
		if v == nil {
			return nil
		}
		return v.ImplementationAcknowledgementVersion
	}).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) NeedFunctionalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.NeedFunctionalAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) NeedImplementationAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.NeedImplementationAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) NeedLoopForValidMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.NeedLoopForValidMessages
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) NeedTechnicalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.NeedTechnicalAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) RolloverAcknowledgementControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.RolloverAcknowledgementControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsPtrOutput) SendSynchronousAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.SendSynchronousAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

type X12AcknowledgementSettingsResponse struct {
	AcknowledgementControlNumberLowerBound int     `pulumi:"acknowledgementControlNumberLowerBound"`
	AcknowledgementControlNumberPrefix     *string `pulumi:"acknowledgementControlNumberPrefix"`
	AcknowledgementControlNumberSuffix     *string `pulumi:"acknowledgementControlNumberSuffix"`
	AcknowledgementControlNumberUpperBound int     `pulumi:"acknowledgementControlNumberUpperBound"`
	BatchFunctionalAcknowledgements        bool    `pulumi:"batchFunctionalAcknowledgements"`
	BatchImplementationAcknowledgements    bool    `pulumi:"batchImplementationAcknowledgements"`
	BatchTechnicalAcknowledgements         bool    `pulumi:"batchTechnicalAcknowledgements"`
	FunctionalAcknowledgementVersion       *string `pulumi:"functionalAcknowledgementVersion"`
	ImplementationAcknowledgementVersion   *string `pulumi:"implementationAcknowledgementVersion"`
	NeedFunctionalAcknowledgement          bool    `pulumi:"needFunctionalAcknowledgement"`
	NeedImplementationAcknowledgement      bool    `pulumi:"needImplementationAcknowledgement"`
	NeedLoopForValidMessages               bool    `pulumi:"needLoopForValidMessages"`
	NeedTechnicalAcknowledgement           bool    `pulumi:"needTechnicalAcknowledgement"`
	RolloverAcknowledgementControlNumber   bool    `pulumi:"rolloverAcknowledgementControlNumber"`
	SendSynchronousAcknowledgement         bool    `pulumi:"sendSynchronousAcknowledgement"`
}

type X12AcknowledgementSettingsResponseOutput struct{ *pulumi.OutputState }

func (X12AcknowledgementSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12AcknowledgementSettingsResponse)(nil)).Elem()
}

func (o X12AcknowledgementSettingsResponseOutput) ToX12AcknowledgementSettingsResponseOutput() X12AcknowledgementSettingsResponseOutput {
	return o
}

func (o X12AcknowledgementSettingsResponseOutput) ToX12AcknowledgementSettingsResponseOutputWithContext(ctx context.Context) X12AcknowledgementSettingsResponseOutput {
	return o
}

func (o X12AcknowledgementSettingsResponseOutput) AcknowledgementControlNumberLowerBound() pulumi.IntOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) int { return v.AcknowledgementControlNumberLowerBound }).(pulumi.IntOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) AcknowledgementControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *string { return v.AcknowledgementControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) AcknowledgementControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *string { return v.AcknowledgementControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) AcknowledgementControlNumberUpperBound() pulumi.IntOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) int { return v.AcknowledgementControlNumberUpperBound }).(pulumi.IntOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) BatchFunctionalAcknowledgements() pulumi.BoolOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) bool { return v.BatchFunctionalAcknowledgements }).(pulumi.BoolOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) BatchImplementationAcknowledgements() pulumi.BoolOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) bool { return v.BatchImplementationAcknowledgements }).(pulumi.BoolOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) BatchTechnicalAcknowledgements() pulumi.BoolOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) bool { return v.BatchTechnicalAcknowledgements }).(pulumi.BoolOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) FunctionalAcknowledgementVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *string { return v.FunctionalAcknowledgementVersion }).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) ImplementationAcknowledgementVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) *string { return v.ImplementationAcknowledgementVersion }).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) NeedFunctionalAcknowledgement() pulumi.BoolOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) bool { return v.NeedFunctionalAcknowledgement }).(pulumi.BoolOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) NeedImplementationAcknowledgement() pulumi.BoolOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) bool { return v.NeedImplementationAcknowledgement }).(pulumi.BoolOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) NeedLoopForValidMessages() pulumi.BoolOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) bool { return v.NeedLoopForValidMessages }).(pulumi.BoolOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) NeedTechnicalAcknowledgement() pulumi.BoolOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) bool { return v.NeedTechnicalAcknowledgement }).(pulumi.BoolOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) RolloverAcknowledgementControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) bool { return v.RolloverAcknowledgementControlNumber }).(pulumi.BoolOutput)
}

func (o X12AcknowledgementSettingsResponseOutput) SendSynchronousAcknowledgement() pulumi.BoolOutput {
	return o.ApplyT(func(v X12AcknowledgementSettingsResponse) bool { return v.SendSynchronousAcknowledgement }).(pulumi.BoolOutput)
}

type X12AcknowledgementSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (X12AcknowledgementSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12AcknowledgementSettingsResponse)(nil)).Elem()
}

func (o X12AcknowledgementSettingsResponsePtrOutput) ToX12AcknowledgementSettingsResponsePtrOutput() X12AcknowledgementSettingsResponsePtrOutput {
	return o
}

func (o X12AcknowledgementSettingsResponsePtrOutput) ToX12AcknowledgementSettingsResponsePtrOutputWithContext(ctx context.Context) X12AcknowledgementSettingsResponsePtrOutput {
	return o
}

func (o X12AcknowledgementSettingsResponsePtrOutput) Elem() X12AcknowledgementSettingsResponseOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) X12AcknowledgementSettingsResponse {
		if v != nil {
			return *v
		}
		var ret X12AcknowledgementSettingsResponse
		return ret
	}).(X12AcknowledgementSettingsResponseOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) AcknowledgementControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.AcknowledgementControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) AcknowledgementControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) AcknowledgementControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AcknowledgementControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) AcknowledgementControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.AcknowledgementControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) BatchFunctionalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.BatchFunctionalAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) BatchImplementationAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.BatchImplementationAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) BatchTechnicalAcknowledgements() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.BatchTechnicalAcknowledgements
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) FunctionalAcknowledgementVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.FunctionalAcknowledgementVersion
	}).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) ImplementationAcknowledgementVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ImplementationAcknowledgementVersion
	}).(pulumi.StringPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) NeedFunctionalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.NeedFunctionalAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) NeedImplementationAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.NeedImplementationAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) NeedLoopForValidMessages() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.NeedLoopForValidMessages
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) NeedTechnicalAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.NeedTechnicalAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) RolloverAcknowledgementControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.RolloverAcknowledgementControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12AcknowledgementSettingsResponsePtrOutput) SendSynchronousAcknowledgement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12AcknowledgementSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.SendSynchronousAcknowledgement
	}).(pulumi.BoolPtrOutput)
}

type X12AgreementContent struct {
	ReceiveAgreement X12OneWayAgreement `pulumi:"receiveAgreement"`
	SendAgreement    X12OneWayAgreement `pulumi:"sendAgreement"`
}





type X12AgreementContentInput interface {
	pulumi.Input

	ToX12AgreementContentOutput() X12AgreementContentOutput
	ToX12AgreementContentOutputWithContext(context.Context) X12AgreementContentOutput
}

type X12AgreementContentArgs struct {
	ReceiveAgreement X12OneWayAgreementInput `pulumi:"receiveAgreement"`
	SendAgreement    X12OneWayAgreementInput `pulumi:"sendAgreement"`
}

func (X12AgreementContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12AgreementContent)(nil)).Elem()
}

func (i X12AgreementContentArgs) ToX12AgreementContentOutput() X12AgreementContentOutput {
	return i.ToX12AgreementContentOutputWithContext(context.Background())
}

func (i X12AgreementContentArgs) ToX12AgreementContentOutputWithContext(ctx context.Context) X12AgreementContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12AgreementContentOutput)
}

func (i X12AgreementContentArgs) ToX12AgreementContentPtrOutput() X12AgreementContentPtrOutput {
	return i.ToX12AgreementContentPtrOutputWithContext(context.Background())
}

func (i X12AgreementContentArgs) ToX12AgreementContentPtrOutputWithContext(ctx context.Context) X12AgreementContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12AgreementContentOutput).ToX12AgreementContentPtrOutputWithContext(ctx)
}









type X12AgreementContentPtrInput interface {
	pulumi.Input

	ToX12AgreementContentPtrOutput() X12AgreementContentPtrOutput
	ToX12AgreementContentPtrOutputWithContext(context.Context) X12AgreementContentPtrOutput
}

type x12agreementContentPtrType X12AgreementContentArgs

func X12AgreementContentPtr(v *X12AgreementContentArgs) X12AgreementContentPtrInput {
	return (*x12agreementContentPtrType)(v)
}

func (*x12agreementContentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12AgreementContent)(nil)).Elem()
}

func (i *x12agreementContentPtrType) ToX12AgreementContentPtrOutput() X12AgreementContentPtrOutput {
	return i.ToX12AgreementContentPtrOutputWithContext(context.Background())
}

func (i *x12agreementContentPtrType) ToX12AgreementContentPtrOutputWithContext(ctx context.Context) X12AgreementContentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12AgreementContentPtrOutput)
}

type X12AgreementContentOutput struct{ *pulumi.OutputState }

func (X12AgreementContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12AgreementContent)(nil)).Elem()
}

func (o X12AgreementContentOutput) ToX12AgreementContentOutput() X12AgreementContentOutput {
	return o
}

func (o X12AgreementContentOutput) ToX12AgreementContentOutputWithContext(ctx context.Context) X12AgreementContentOutput {
	return o
}

func (o X12AgreementContentOutput) ToX12AgreementContentPtrOutput() X12AgreementContentPtrOutput {
	return o.ToX12AgreementContentPtrOutputWithContext(context.Background())
}

func (o X12AgreementContentOutput) ToX12AgreementContentPtrOutputWithContext(ctx context.Context) X12AgreementContentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12AgreementContent) *X12AgreementContent {
		return &v
	}).(X12AgreementContentPtrOutput)
}

func (o X12AgreementContentOutput) ReceiveAgreement() X12OneWayAgreementOutput {
	return o.ApplyT(func(v X12AgreementContent) X12OneWayAgreement { return v.ReceiveAgreement }).(X12OneWayAgreementOutput)
}

func (o X12AgreementContentOutput) SendAgreement() X12OneWayAgreementOutput {
	return o.ApplyT(func(v X12AgreementContent) X12OneWayAgreement { return v.SendAgreement }).(X12OneWayAgreementOutput)
}

type X12AgreementContentPtrOutput struct{ *pulumi.OutputState }

func (X12AgreementContentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12AgreementContent)(nil)).Elem()
}

func (o X12AgreementContentPtrOutput) ToX12AgreementContentPtrOutput() X12AgreementContentPtrOutput {
	return o
}

func (o X12AgreementContentPtrOutput) ToX12AgreementContentPtrOutputWithContext(ctx context.Context) X12AgreementContentPtrOutput {
	return o
}

func (o X12AgreementContentPtrOutput) Elem() X12AgreementContentOutput {
	return o.ApplyT(func(v *X12AgreementContent) X12AgreementContent {
		if v != nil {
			return *v
		}
		var ret X12AgreementContent
		return ret
	}).(X12AgreementContentOutput)
}

func (o X12AgreementContentPtrOutput) ReceiveAgreement() X12OneWayAgreementPtrOutput {
	return o.ApplyT(func(v *X12AgreementContent) *X12OneWayAgreement {
		if v == nil {
			return nil
		}
		return &v.ReceiveAgreement
	}).(X12OneWayAgreementPtrOutput)
}

func (o X12AgreementContentPtrOutput) SendAgreement() X12OneWayAgreementPtrOutput {
	return o.ApplyT(func(v *X12AgreementContent) *X12OneWayAgreement {
		if v == nil {
			return nil
		}
		return &v.SendAgreement
	}).(X12OneWayAgreementPtrOutput)
}

type X12AgreementContentResponse struct {
	ReceiveAgreement X12OneWayAgreementResponse `pulumi:"receiveAgreement"`
	SendAgreement    X12OneWayAgreementResponse `pulumi:"sendAgreement"`
}

type X12AgreementContentResponseOutput struct{ *pulumi.OutputState }

func (X12AgreementContentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12AgreementContentResponse)(nil)).Elem()
}

func (o X12AgreementContentResponseOutput) ToX12AgreementContentResponseOutput() X12AgreementContentResponseOutput {
	return o
}

func (o X12AgreementContentResponseOutput) ToX12AgreementContentResponseOutputWithContext(ctx context.Context) X12AgreementContentResponseOutput {
	return o
}

func (o X12AgreementContentResponseOutput) ReceiveAgreement() X12OneWayAgreementResponseOutput {
	return o.ApplyT(func(v X12AgreementContentResponse) X12OneWayAgreementResponse { return v.ReceiveAgreement }).(X12OneWayAgreementResponseOutput)
}

func (o X12AgreementContentResponseOutput) SendAgreement() X12OneWayAgreementResponseOutput {
	return o.ApplyT(func(v X12AgreementContentResponse) X12OneWayAgreementResponse { return v.SendAgreement }).(X12OneWayAgreementResponseOutput)
}

type X12AgreementContentResponsePtrOutput struct{ *pulumi.OutputState }

func (X12AgreementContentResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12AgreementContentResponse)(nil)).Elem()
}

func (o X12AgreementContentResponsePtrOutput) ToX12AgreementContentResponsePtrOutput() X12AgreementContentResponsePtrOutput {
	return o
}

func (o X12AgreementContentResponsePtrOutput) ToX12AgreementContentResponsePtrOutputWithContext(ctx context.Context) X12AgreementContentResponsePtrOutput {
	return o
}

func (o X12AgreementContentResponsePtrOutput) Elem() X12AgreementContentResponseOutput {
	return o.ApplyT(func(v *X12AgreementContentResponse) X12AgreementContentResponse {
		if v != nil {
			return *v
		}
		var ret X12AgreementContentResponse
		return ret
	}).(X12AgreementContentResponseOutput)
}

func (o X12AgreementContentResponsePtrOutput) ReceiveAgreement() X12OneWayAgreementResponsePtrOutput {
	return o.ApplyT(func(v *X12AgreementContentResponse) *X12OneWayAgreementResponse {
		if v == nil {
			return nil
		}
		return &v.ReceiveAgreement
	}).(X12OneWayAgreementResponsePtrOutput)
}

func (o X12AgreementContentResponsePtrOutput) SendAgreement() X12OneWayAgreementResponsePtrOutput {
	return o.ApplyT(func(v *X12AgreementContentResponse) *X12OneWayAgreementResponse {
		if v == nil {
			return nil
		}
		return &v.SendAgreement
	}).(X12OneWayAgreementResponsePtrOutput)
}

type X12DelimiterOverrides struct {
	ComponentSeparator         int                     `pulumi:"componentSeparator"`
	DataElementSeparator       int                     `pulumi:"dataElementSeparator"`
	MessageId                  *string                 `pulumi:"messageId"`
	ProtocolVersion            *string                 `pulumi:"protocolVersion"`
	ReplaceCharacter           int                     `pulumi:"replaceCharacter"`
	ReplaceSeparatorsInPayload bool                    `pulumi:"replaceSeparatorsInPayload"`
	SegmentTerminator          int                     `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix    SegmentTerminatorSuffix `pulumi:"segmentTerminatorSuffix"`
	TargetNamespace            *string                 `pulumi:"targetNamespace"`
}





type X12DelimiterOverridesInput interface {
	pulumi.Input

	ToX12DelimiterOverridesOutput() X12DelimiterOverridesOutput
	ToX12DelimiterOverridesOutputWithContext(context.Context) X12DelimiterOverridesOutput
}

type X12DelimiterOverridesArgs struct {
	ComponentSeparator         pulumi.IntInput              `pulumi:"componentSeparator"`
	DataElementSeparator       pulumi.IntInput              `pulumi:"dataElementSeparator"`
	MessageId                  pulumi.StringPtrInput        `pulumi:"messageId"`
	ProtocolVersion            pulumi.StringPtrInput        `pulumi:"protocolVersion"`
	ReplaceCharacter           pulumi.IntInput              `pulumi:"replaceCharacter"`
	ReplaceSeparatorsInPayload pulumi.BoolInput             `pulumi:"replaceSeparatorsInPayload"`
	SegmentTerminator          pulumi.IntInput              `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix    SegmentTerminatorSuffixInput `pulumi:"segmentTerminatorSuffix"`
	TargetNamespace            pulumi.StringPtrInput        `pulumi:"targetNamespace"`
}

func (X12DelimiterOverridesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12DelimiterOverrides)(nil)).Elem()
}

func (i X12DelimiterOverridesArgs) ToX12DelimiterOverridesOutput() X12DelimiterOverridesOutput {
	return i.ToX12DelimiterOverridesOutputWithContext(context.Background())
}

func (i X12DelimiterOverridesArgs) ToX12DelimiterOverridesOutputWithContext(ctx context.Context) X12DelimiterOverridesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12DelimiterOverridesOutput)
}





type X12DelimiterOverridesArrayInput interface {
	pulumi.Input

	ToX12DelimiterOverridesArrayOutput() X12DelimiterOverridesArrayOutput
	ToX12DelimiterOverridesArrayOutputWithContext(context.Context) X12DelimiterOverridesArrayOutput
}

type X12DelimiterOverridesArray []X12DelimiterOverridesInput

func (X12DelimiterOverridesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12DelimiterOverrides)(nil)).Elem()
}

func (i X12DelimiterOverridesArray) ToX12DelimiterOverridesArrayOutput() X12DelimiterOverridesArrayOutput {
	return i.ToX12DelimiterOverridesArrayOutputWithContext(context.Background())
}

func (i X12DelimiterOverridesArray) ToX12DelimiterOverridesArrayOutputWithContext(ctx context.Context) X12DelimiterOverridesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12DelimiterOverridesArrayOutput)
}

type X12DelimiterOverridesOutput struct{ *pulumi.OutputState }

func (X12DelimiterOverridesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12DelimiterOverrides)(nil)).Elem()
}

func (o X12DelimiterOverridesOutput) ToX12DelimiterOverridesOutput() X12DelimiterOverridesOutput {
	return o
}

func (o X12DelimiterOverridesOutput) ToX12DelimiterOverridesOutputWithContext(ctx context.Context) X12DelimiterOverridesOutput {
	return o
}

func (o X12DelimiterOverridesOutput) ComponentSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v X12DelimiterOverrides) int { return v.ComponentSeparator }).(pulumi.IntOutput)
}

func (o X12DelimiterOverridesOutput) DataElementSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v X12DelimiterOverrides) int { return v.DataElementSeparator }).(pulumi.IntOutput)
}

func (o X12DelimiterOverridesOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverrides) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o X12DelimiterOverridesOutput) ProtocolVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverrides) *string { return v.ProtocolVersion }).(pulumi.StringPtrOutput)
}

func (o X12DelimiterOverridesOutput) ReplaceCharacter() pulumi.IntOutput {
	return o.ApplyT(func(v X12DelimiterOverrides) int { return v.ReplaceCharacter }).(pulumi.IntOutput)
}

func (o X12DelimiterOverridesOutput) ReplaceSeparatorsInPayload() pulumi.BoolOutput {
	return o.ApplyT(func(v X12DelimiterOverrides) bool { return v.ReplaceSeparatorsInPayload }).(pulumi.BoolOutput)
}

func (o X12DelimiterOverridesOutput) SegmentTerminator() pulumi.IntOutput {
	return o.ApplyT(func(v X12DelimiterOverrides) int { return v.SegmentTerminator }).(pulumi.IntOutput)
}

func (o X12DelimiterOverridesOutput) SegmentTerminatorSuffix() SegmentTerminatorSuffixOutput {
	return o.ApplyT(func(v X12DelimiterOverrides) SegmentTerminatorSuffix { return v.SegmentTerminatorSuffix }).(SegmentTerminatorSuffixOutput)
}

func (o X12DelimiterOverridesOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverrides) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

type X12DelimiterOverridesArrayOutput struct{ *pulumi.OutputState }

func (X12DelimiterOverridesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12DelimiterOverrides)(nil)).Elem()
}

func (o X12DelimiterOverridesArrayOutput) ToX12DelimiterOverridesArrayOutput() X12DelimiterOverridesArrayOutput {
	return o
}

func (o X12DelimiterOverridesArrayOutput) ToX12DelimiterOverridesArrayOutputWithContext(ctx context.Context) X12DelimiterOverridesArrayOutput {
	return o
}

func (o X12DelimiterOverridesArrayOutput) Index(i pulumi.IntInput) X12DelimiterOverridesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12DelimiterOverrides {
		return vs[0].([]X12DelimiterOverrides)[vs[1].(int)]
	}).(X12DelimiterOverridesOutput)
}

type X12DelimiterOverridesResponse struct {
	ComponentSeparator         int     `pulumi:"componentSeparator"`
	DataElementSeparator       int     `pulumi:"dataElementSeparator"`
	MessageId                  *string `pulumi:"messageId"`
	ProtocolVersion            *string `pulumi:"protocolVersion"`
	ReplaceCharacter           int     `pulumi:"replaceCharacter"`
	ReplaceSeparatorsInPayload bool    `pulumi:"replaceSeparatorsInPayload"`
	SegmentTerminator          int     `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix    string  `pulumi:"segmentTerminatorSuffix"`
	TargetNamespace            *string `pulumi:"targetNamespace"`
}

type X12DelimiterOverridesResponseOutput struct{ *pulumi.OutputState }

func (X12DelimiterOverridesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12DelimiterOverridesResponse)(nil)).Elem()
}

func (o X12DelimiterOverridesResponseOutput) ToX12DelimiterOverridesResponseOutput() X12DelimiterOverridesResponseOutput {
	return o
}

func (o X12DelimiterOverridesResponseOutput) ToX12DelimiterOverridesResponseOutputWithContext(ctx context.Context) X12DelimiterOverridesResponseOutput {
	return o
}

func (o X12DelimiterOverridesResponseOutput) ComponentSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v X12DelimiterOverridesResponse) int { return v.ComponentSeparator }).(pulumi.IntOutput)
}

func (o X12DelimiterOverridesResponseOutput) DataElementSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v X12DelimiterOverridesResponse) int { return v.DataElementSeparator }).(pulumi.IntOutput)
}

func (o X12DelimiterOverridesResponseOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverridesResponse) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o X12DelimiterOverridesResponseOutput) ProtocolVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverridesResponse) *string { return v.ProtocolVersion }).(pulumi.StringPtrOutput)
}

func (o X12DelimiterOverridesResponseOutput) ReplaceCharacter() pulumi.IntOutput {
	return o.ApplyT(func(v X12DelimiterOverridesResponse) int { return v.ReplaceCharacter }).(pulumi.IntOutput)
}

func (o X12DelimiterOverridesResponseOutput) ReplaceSeparatorsInPayload() pulumi.BoolOutput {
	return o.ApplyT(func(v X12DelimiterOverridesResponse) bool { return v.ReplaceSeparatorsInPayload }).(pulumi.BoolOutput)
}

func (o X12DelimiterOverridesResponseOutput) SegmentTerminator() pulumi.IntOutput {
	return o.ApplyT(func(v X12DelimiterOverridesResponse) int { return v.SegmentTerminator }).(pulumi.IntOutput)
}

func (o X12DelimiterOverridesResponseOutput) SegmentTerminatorSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v X12DelimiterOverridesResponse) string { return v.SegmentTerminatorSuffix }).(pulumi.StringOutput)
}

func (o X12DelimiterOverridesResponseOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12DelimiterOverridesResponse) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

type X12DelimiterOverridesResponseArrayOutput struct{ *pulumi.OutputState }

func (X12DelimiterOverridesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12DelimiterOverridesResponse)(nil)).Elem()
}

func (o X12DelimiterOverridesResponseArrayOutput) ToX12DelimiterOverridesResponseArrayOutput() X12DelimiterOverridesResponseArrayOutput {
	return o
}

func (o X12DelimiterOverridesResponseArrayOutput) ToX12DelimiterOverridesResponseArrayOutputWithContext(ctx context.Context) X12DelimiterOverridesResponseArrayOutput {
	return o
}

func (o X12DelimiterOverridesResponseArrayOutput) Index(i pulumi.IntInput) X12DelimiterOverridesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12DelimiterOverridesResponse {
		return vs[0].([]X12DelimiterOverridesResponse)[vs[1].(int)]
	}).(X12DelimiterOverridesResponseOutput)
}

type X12EnvelopeOverride struct {
	DateFormat               string  `pulumi:"dateFormat"`
	FunctionalIdentifierCode *string `pulumi:"functionalIdentifierCode"`
	HeaderVersion            string  `pulumi:"headerVersion"`
	MessageId                string  `pulumi:"messageId"`
	ProtocolVersion          string  `pulumi:"protocolVersion"`
	ReceiverApplicationId    string  `pulumi:"receiverApplicationId"`
	ResponsibleAgencyCode    string  `pulumi:"responsibleAgencyCode"`
	SenderApplicationId      string  `pulumi:"senderApplicationId"`
	TargetNamespace          string  `pulumi:"targetNamespace"`
	TimeFormat               string  `pulumi:"timeFormat"`
}





type X12EnvelopeOverrideInput interface {
	pulumi.Input

	ToX12EnvelopeOverrideOutput() X12EnvelopeOverrideOutput
	ToX12EnvelopeOverrideOutputWithContext(context.Context) X12EnvelopeOverrideOutput
}

type X12EnvelopeOverrideArgs struct {
	DateFormat               pulumi.StringInput    `pulumi:"dateFormat"`
	FunctionalIdentifierCode pulumi.StringPtrInput `pulumi:"functionalIdentifierCode"`
	HeaderVersion            pulumi.StringInput    `pulumi:"headerVersion"`
	MessageId                pulumi.StringInput    `pulumi:"messageId"`
	ProtocolVersion          pulumi.StringInput    `pulumi:"protocolVersion"`
	ReceiverApplicationId    pulumi.StringInput    `pulumi:"receiverApplicationId"`
	ResponsibleAgencyCode    pulumi.StringInput    `pulumi:"responsibleAgencyCode"`
	SenderApplicationId      pulumi.StringInput    `pulumi:"senderApplicationId"`
	TargetNamespace          pulumi.StringInput    `pulumi:"targetNamespace"`
	TimeFormat               pulumi.StringInput    `pulumi:"timeFormat"`
}

func (X12EnvelopeOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12EnvelopeOverride)(nil)).Elem()
}

func (i X12EnvelopeOverrideArgs) ToX12EnvelopeOverrideOutput() X12EnvelopeOverrideOutput {
	return i.ToX12EnvelopeOverrideOutputWithContext(context.Background())
}

func (i X12EnvelopeOverrideArgs) ToX12EnvelopeOverrideOutputWithContext(ctx context.Context) X12EnvelopeOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12EnvelopeOverrideOutput)
}





type X12EnvelopeOverrideArrayInput interface {
	pulumi.Input

	ToX12EnvelopeOverrideArrayOutput() X12EnvelopeOverrideArrayOutput
	ToX12EnvelopeOverrideArrayOutputWithContext(context.Context) X12EnvelopeOverrideArrayOutput
}

type X12EnvelopeOverrideArray []X12EnvelopeOverrideInput

func (X12EnvelopeOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12EnvelopeOverride)(nil)).Elem()
}

func (i X12EnvelopeOverrideArray) ToX12EnvelopeOverrideArrayOutput() X12EnvelopeOverrideArrayOutput {
	return i.ToX12EnvelopeOverrideArrayOutputWithContext(context.Background())
}

func (i X12EnvelopeOverrideArray) ToX12EnvelopeOverrideArrayOutputWithContext(ctx context.Context) X12EnvelopeOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12EnvelopeOverrideArrayOutput)
}

type X12EnvelopeOverrideOutput struct{ *pulumi.OutputState }

func (X12EnvelopeOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12EnvelopeOverride)(nil)).Elem()
}

func (o X12EnvelopeOverrideOutput) ToX12EnvelopeOverrideOutput() X12EnvelopeOverrideOutput {
	return o
}

func (o X12EnvelopeOverrideOutput) ToX12EnvelopeOverrideOutputWithContext(ctx context.Context) X12EnvelopeOverrideOutput {
	return o
}

func (o X12EnvelopeOverrideOutput) DateFormat() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) string { return v.DateFormat }).(pulumi.StringOutput)
}

func (o X12EnvelopeOverrideOutput) FunctionalIdentifierCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) *string { return v.FunctionalIdentifierCode }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeOverrideOutput) HeaderVersion() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) string { return v.HeaderVersion }).(pulumi.StringOutput)
}

func (o X12EnvelopeOverrideOutput) MessageId() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) string { return v.MessageId }).(pulumi.StringOutput)
}

func (o X12EnvelopeOverrideOutput) ProtocolVersion() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) string { return v.ProtocolVersion }).(pulumi.StringOutput)
}

func (o X12EnvelopeOverrideOutput) ReceiverApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) string { return v.ReceiverApplicationId }).(pulumi.StringOutput)
}

func (o X12EnvelopeOverrideOutput) ResponsibleAgencyCode() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) string { return v.ResponsibleAgencyCode }).(pulumi.StringOutput)
}

func (o X12EnvelopeOverrideOutput) SenderApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) string { return v.SenderApplicationId }).(pulumi.StringOutput)
}

func (o X12EnvelopeOverrideOutput) TargetNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) string { return v.TargetNamespace }).(pulumi.StringOutput)
}

func (o X12EnvelopeOverrideOutput) TimeFormat() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeOverride) string { return v.TimeFormat }).(pulumi.StringOutput)
}

type X12EnvelopeOverrideArrayOutput struct{ *pulumi.OutputState }

func (X12EnvelopeOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12EnvelopeOverride)(nil)).Elem()
}

func (o X12EnvelopeOverrideArrayOutput) ToX12EnvelopeOverrideArrayOutput() X12EnvelopeOverrideArrayOutput {
	return o
}

func (o X12EnvelopeOverrideArrayOutput) ToX12EnvelopeOverrideArrayOutputWithContext(ctx context.Context) X12EnvelopeOverrideArrayOutput {
	return o
}

func (o X12EnvelopeOverrideArrayOutput) Index(i pulumi.IntInput) X12EnvelopeOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12EnvelopeOverride {
		return vs[0].([]X12EnvelopeOverride)[vs[1].(int)]
	}).(X12EnvelopeOverrideOutput)
}

type X12EnvelopeOverrideResponse struct {
	DateFormat               string  `pulumi:"dateFormat"`
	FunctionalIdentifierCode *string `pulumi:"functionalIdentifierCode"`
	HeaderVersion            string  `pulumi:"headerVersion"`
	MessageId                string  `pulumi:"messageId"`
	ProtocolVersion          string  `pulumi:"protocolVersion"`
	ReceiverApplicationId    string  `pulumi:"receiverApplicationId"`
	ResponsibleAgencyCode    string  `pulumi:"responsibleAgencyCode"`
	SenderApplicationId      string  `pulumi:"senderApplicationId"`
	TargetNamespace          string  `pulumi:"targetNamespace"`
	TimeFormat               string  `pulumi:"timeFormat"`
}

type X12EnvelopeOverrideResponseOutput struct{ *pulumi.OutputState }

func (X12EnvelopeOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12EnvelopeOverrideResponse)(nil)).Elem()
}

func (o X12EnvelopeOverrideResponseOutput) ToX12EnvelopeOverrideResponseOutput() X12EnvelopeOverrideResponseOutput {
	return o
}

func (o X12EnvelopeOverrideResponseOutput) ToX12EnvelopeOverrideResponseOutputWithContext(ctx context.Context) X12EnvelopeOverrideResponseOutput {
	return o
}

func (o X12EnvelopeOverrideResponseOutput) DateFormat() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) string { return v.DateFormat }).(pulumi.StringOutput)
}

func (o X12EnvelopeOverrideResponseOutput) FunctionalIdentifierCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) *string { return v.FunctionalIdentifierCode }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeOverrideResponseOutput) HeaderVersion() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) string { return v.HeaderVersion }).(pulumi.StringOutput)
}

func (o X12EnvelopeOverrideResponseOutput) MessageId() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) string { return v.MessageId }).(pulumi.StringOutput)
}

func (o X12EnvelopeOverrideResponseOutput) ProtocolVersion() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) string { return v.ProtocolVersion }).(pulumi.StringOutput)
}

func (o X12EnvelopeOverrideResponseOutput) ReceiverApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) string { return v.ReceiverApplicationId }).(pulumi.StringOutput)
}

func (o X12EnvelopeOverrideResponseOutput) ResponsibleAgencyCode() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) string { return v.ResponsibleAgencyCode }).(pulumi.StringOutput)
}

func (o X12EnvelopeOverrideResponseOutput) SenderApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) string { return v.SenderApplicationId }).(pulumi.StringOutput)
}

func (o X12EnvelopeOverrideResponseOutput) TargetNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) string { return v.TargetNamespace }).(pulumi.StringOutput)
}

func (o X12EnvelopeOverrideResponseOutput) TimeFormat() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeOverrideResponse) string { return v.TimeFormat }).(pulumi.StringOutput)
}

type X12EnvelopeOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (X12EnvelopeOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12EnvelopeOverrideResponse)(nil)).Elem()
}

func (o X12EnvelopeOverrideResponseArrayOutput) ToX12EnvelopeOverrideResponseArrayOutput() X12EnvelopeOverrideResponseArrayOutput {
	return o
}

func (o X12EnvelopeOverrideResponseArrayOutput) ToX12EnvelopeOverrideResponseArrayOutputWithContext(ctx context.Context) X12EnvelopeOverrideResponseArrayOutput {
	return o
}

func (o X12EnvelopeOverrideResponseArrayOutput) Index(i pulumi.IntInput) X12EnvelopeOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12EnvelopeOverrideResponse {
		return vs[0].([]X12EnvelopeOverrideResponse)[vs[1].(int)]
	}).(X12EnvelopeOverrideResponseOutput)
}

type X12EnvelopeSettings struct {
	ControlStandardsId                           int     `pulumi:"controlStandardsId"`
	ControlVersionNumber                         string  `pulumi:"controlVersionNumber"`
	EnableDefaultGroupHeaders                    bool    `pulumi:"enableDefaultGroupHeaders"`
	FunctionalGroupId                            *string `pulumi:"functionalGroupId"`
	GroupControlNumberLowerBound                 int     `pulumi:"groupControlNumberLowerBound"`
	GroupControlNumberUpperBound                 int     `pulumi:"groupControlNumberUpperBound"`
	GroupHeaderAgencyCode                        string  `pulumi:"groupHeaderAgencyCode"`
	GroupHeaderDateFormat                        string  `pulumi:"groupHeaderDateFormat"`
	GroupHeaderTimeFormat                        string  `pulumi:"groupHeaderTimeFormat"`
	GroupHeaderVersion                           string  `pulumi:"groupHeaderVersion"`
	InterchangeControlNumberLowerBound           int     `pulumi:"interchangeControlNumberLowerBound"`
	InterchangeControlNumberUpperBound           int     `pulumi:"interchangeControlNumberUpperBound"`
	OverwriteExistingTransactionSetControlNumber bool    `pulumi:"overwriteExistingTransactionSetControlNumber"`
	ReceiverApplicationId                        string  `pulumi:"receiverApplicationId"`
	RolloverGroupControlNumber                   bool    `pulumi:"rolloverGroupControlNumber"`
	RolloverInterchangeControlNumber             bool    `pulumi:"rolloverInterchangeControlNumber"`
	RolloverTransactionSetControlNumber          bool    `pulumi:"rolloverTransactionSetControlNumber"`
	SenderApplicationId                          string  `pulumi:"senderApplicationId"`
	TransactionSetControlNumberLowerBound        int     `pulumi:"transactionSetControlNumberLowerBound"`
	TransactionSetControlNumberPrefix            *string `pulumi:"transactionSetControlNumberPrefix"`
	TransactionSetControlNumberSuffix            *string `pulumi:"transactionSetControlNumberSuffix"`
	TransactionSetControlNumberUpperBound        int     `pulumi:"transactionSetControlNumberUpperBound"`
	UsageIndicator                               string  `pulumi:"usageIndicator"`
	UseControlStandardsIdAsRepetitionCharacter   bool    `pulumi:"useControlStandardsIdAsRepetitionCharacter"`
}





type X12EnvelopeSettingsInput interface {
	pulumi.Input

	ToX12EnvelopeSettingsOutput() X12EnvelopeSettingsOutput
	ToX12EnvelopeSettingsOutputWithContext(context.Context) X12EnvelopeSettingsOutput
}

type X12EnvelopeSettingsArgs struct {
	ControlStandardsId                           pulumi.IntInput       `pulumi:"controlStandardsId"`
	ControlVersionNumber                         pulumi.StringInput    `pulumi:"controlVersionNumber"`
	EnableDefaultGroupHeaders                    pulumi.BoolInput      `pulumi:"enableDefaultGroupHeaders"`
	FunctionalGroupId                            pulumi.StringPtrInput `pulumi:"functionalGroupId"`
	GroupControlNumberLowerBound                 pulumi.IntInput       `pulumi:"groupControlNumberLowerBound"`
	GroupControlNumberUpperBound                 pulumi.IntInput       `pulumi:"groupControlNumberUpperBound"`
	GroupHeaderAgencyCode                        pulumi.StringInput    `pulumi:"groupHeaderAgencyCode"`
	GroupHeaderDateFormat                        pulumi.StringInput    `pulumi:"groupHeaderDateFormat"`
	GroupHeaderTimeFormat                        pulumi.StringInput    `pulumi:"groupHeaderTimeFormat"`
	GroupHeaderVersion                           pulumi.StringInput    `pulumi:"groupHeaderVersion"`
	InterchangeControlNumberLowerBound           pulumi.IntInput       `pulumi:"interchangeControlNumberLowerBound"`
	InterchangeControlNumberUpperBound           pulumi.IntInput       `pulumi:"interchangeControlNumberUpperBound"`
	OverwriteExistingTransactionSetControlNumber pulumi.BoolInput      `pulumi:"overwriteExistingTransactionSetControlNumber"`
	ReceiverApplicationId                        pulumi.StringInput    `pulumi:"receiverApplicationId"`
	RolloverGroupControlNumber                   pulumi.BoolInput      `pulumi:"rolloverGroupControlNumber"`
	RolloverInterchangeControlNumber             pulumi.BoolInput      `pulumi:"rolloverInterchangeControlNumber"`
	RolloverTransactionSetControlNumber          pulumi.BoolInput      `pulumi:"rolloverTransactionSetControlNumber"`
	SenderApplicationId                          pulumi.StringInput    `pulumi:"senderApplicationId"`
	TransactionSetControlNumberLowerBound        pulumi.IntInput       `pulumi:"transactionSetControlNumberLowerBound"`
	TransactionSetControlNumberPrefix            pulumi.StringPtrInput `pulumi:"transactionSetControlNumberPrefix"`
	TransactionSetControlNumberSuffix            pulumi.StringPtrInput `pulumi:"transactionSetControlNumberSuffix"`
	TransactionSetControlNumberUpperBound        pulumi.IntInput       `pulumi:"transactionSetControlNumberUpperBound"`
	UsageIndicator                               pulumi.StringInput    `pulumi:"usageIndicator"`
	UseControlStandardsIdAsRepetitionCharacter   pulumi.BoolInput      `pulumi:"useControlStandardsIdAsRepetitionCharacter"`
}

func (X12EnvelopeSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12EnvelopeSettings)(nil)).Elem()
}

func (i X12EnvelopeSettingsArgs) ToX12EnvelopeSettingsOutput() X12EnvelopeSettingsOutput {
	return i.ToX12EnvelopeSettingsOutputWithContext(context.Background())
}

func (i X12EnvelopeSettingsArgs) ToX12EnvelopeSettingsOutputWithContext(ctx context.Context) X12EnvelopeSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12EnvelopeSettingsOutput)
}

func (i X12EnvelopeSettingsArgs) ToX12EnvelopeSettingsPtrOutput() X12EnvelopeSettingsPtrOutput {
	return i.ToX12EnvelopeSettingsPtrOutputWithContext(context.Background())
}

func (i X12EnvelopeSettingsArgs) ToX12EnvelopeSettingsPtrOutputWithContext(ctx context.Context) X12EnvelopeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12EnvelopeSettingsOutput).ToX12EnvelopeSettingsPtrOutputWithContext(ctx)
}









type X12EnvelopeSettingsPtrInput interface {
	pulumi.Input

	ToX12EnvelopeSettingsPtrOutput() X12EnvelopeSettingsPtrOutput
	ToX12EnvelopeSettingsPtrOutputWithContext(context.Context) X12EnvelopeSettingsPtrOutput
}

type x12envelopeSettingsPtrType X12EnvelopeSettingsArgs

func X12EnvelopeSettingsPtr(v *X12EnvelopeSettingsArgs) X12EnvelopeSettingsPtrInput {
	return (*x12envelopeSettingsPtrType)(v)
}

func (*x12envelopeSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12EnvelopeSettings)(nil)).Elem()
}

func (i *x12envelopeSettingsPtrType) ToX12EnvelopeSettingsPtrOutput() X12EnvelopeSettingsPtrOutput {
	return i.ToX12EnvelopeSettingsPtrOutputWithContext(context.Background())
}

func (i *x12envelopeSettingsPtrType) ToX12EnvelopeSettingsPtrOutputWithContext(ctx context.Context) X12EnvelopeSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12EnvelopeSettingsPtrOutput)
}

type X12EnvelopeSettingsOutput struct{ *pulumi.OutputState }

func (X12EnvelopeSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12EnvelopeSettings)(nil)).Elem()
}

func (o X12EnvelopeSettingsOutput) ToX12EnvelopeSettingsOutput() X12EnvelopeSettingsOutput {
	return o
}

func (o X12EnvelopeSettingsOutput) ToX12EnvelopeSettingsOutputWithContext(ctx context.Context) X12EnvelopeSettingsOutput {
	return o
}

func (o X12EnvelopeSettingsOutput) ToX12EnvelopeSettingsPtrOutput() X12EnvelopeSettingsPtrOutput {
	return o.ToX12EnvelopeSettingsPtrOutputWithContext(context.Background())
}

func (o X12EnvelopeSettingsOutput) ToX12EnvelopeSettingsPtrOutputWithContext(ctx context.Context) X12EnvelopeSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12EnvelopeSettings) *X12EnvelopeSettings {
		return &v
	}).(X12EnvelopeSettingsPtrOutput)
}

func (o X12EnvelopeSettingsOutput) ControlStandardsId() pulumi.IntOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) int { return v.ControlStandardsId }).(pulumi.IntOutput)
}

func (o X12EnvelopeSettingsOutput) ControlVersionNumber() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) string { return v.ControlVersionNumber }).(pulumi.StringOutput)
}

func (o X12EnvelopeSettingsOutput) EnableDefaultGroupHeaders() pulumi.BoolOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) bool { return v.EnableDefaultGroupHeaders }).(pulumi.BoolOutput)
}

func (o X12EnvelopeSettingsOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *string { return v.FunctionalGroupId }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsOutput) GroupControlNumberLowerBound() pulumi.IntOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) int { return v.GroupControlNumberLowerBound }).(pulumi.IntOutput)
}

func (o X12EnvelopeSettingsOutput) GroupControlNumberUpperBound() pulumi.IntOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) int { return v.GroupControlNumberUpperBound }).(pulumi.IntOutput)
}

func (o X12EnvelopeSettingsOutput) GroupHeaderAgencyCode() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) string { return v.GroupHeaderAgencyCode }).(pulumi.StringOutput)
}

func (o X12EnvelopeSettingsOutput) GroupHeaderDateFormat() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) string { return v.GroupHeaderDateFormat }).(pulumi.StringOutput)
}

func (o X12EnvelopeSettingsOutput) GroupHeaderTimeFormat() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) string { return v.GroupHeaderTimeFormat }).(pulumi.StringOutput)
}

func (o X12EnvelopeSettingsOutput) GroupHeaderVersion() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) string { return v.GroupHeaderVersion }).(pulumi.StringOutput)
}

func (o X12EnvelopeSettingsOutput) InterchangeControlNumberLowerBound() pulumi.IntOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) int { return v.InterchangeControlNumberLowerBound }).(pulumi.IntOutput)
}

func (o X12EnvelopeSettingsOutput) InterchangeControlNumberUpperBound() pulumi.IntOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) int { return v.InterchangeControlNumberUpperBound }).(pulumi.IntOutput)
}

func (o X12EnvelopeSettingsOutput) OverwriteExistingTransactionSetControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) bool { return v.OverwriteExistingTransactionSetControlNumber }).(pulumi.BoolOutput)
}

func (o X12EnvelopeSettingsOutput) ReceiverApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) string { return v.ReceiverApplicationId }).(pulumi.StringOutput)
}

func (o X12EnvelopeSettingsOutput) RolloverGroupControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) bool { return v.RolloverGroupControlNumber }).(pulumi.BoolOutput)
}

func (o X12EnvelopeSettingsOutput) RolloverInterchangeControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) bool { return v.RolloverInterchangeControlNumber }).(pulumi.BoolOutput)
}

func (o X12EnvelopeSettingsOutput) RolloverTransactionSetControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) bool { return v.RolloverTransactionSetControlNumber }).(pulumi.BoolOutput)
}

func (o X12EnvelopeSettingsOutput) SenderApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) string { return v.SenderApplicationId }).(pulumi.StringOutput)
}

func (o X12EnvelopeSettingsOutput) TransactionSetControlNumberLowerBound() pulumi.IntOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) int { return v.TransactionSetControlNumberLowerBound }).(pulumi.IntOutput)
}

func (o X12EnvelopeSettingsOutput) TransactionSetControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *string { return v.TransactionSetControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsOutput) TransactionSetControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) *string { return v.TransactionSetControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsOutput) TransactionSetControlNumberUpperBound() pulumi.IntOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) int { return v.TransactionSetControlNumberUpperBound }).(pulumi.IntOutput)
}

func (o X12EnvelopeSettingsOutput) UsageIndicator() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) string { return v.UsageIndicator }).(pulumi.StringOutput)
}

func (o X12EnvelopeSettingsOutput) UseControlStandardsIdAsRepetitionCharacter() pulumi.BoolOutput {
	return o.ApplyT(func(v X12EnvelopeSettings) bool { return v.UseControlStandardsIdAsRepetitionCharacter }).(pulumi.BoolOutput)
}

type X12EnvelopeSettingsPtrOutput struct{ *pulumi.OutputState }

func (X12EnvelopeSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12EnvelopeSettings)(nil)).Elem()
}

func (o X12EnvelopeSettingsPtrOutput) ToX12EnvelopeSettingsPtrOutput() X12EnvelopeSettingsPtrOutput {
	return o
}

func (o X12EnvelopeSettingsPtrOutput) ToX12EnvelopeSettingsPtrOutputWithContext(ctx context.Context) X12EnvelopeSettingsPtrOutput {
	return o
}

func (o X12EnvelopeSettingsPtrOutput) Elem() X12EnvelopeSettingsOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) X12EnvelopeSettings {
		if v != nil {
			return *v
		}
		var ret X12EnvelopeSettings
		return ret
	}).(X12EnvelopeSettingsOutput)
}

func (o X12EnvelopeSettingsPtrOutput) ControlStandardsId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *int {
		if v == nil {
			return nil
		}
		return &v.ControlStandardsId
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) ControlVersionNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return &v.ControlVersionNumber
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) EnableDefaultGroupHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.EnableDefaultGroupHeaders
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.FunctionalGroupId
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) GroupControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *int {
		if v == nil {
			return nil
		}
		return &v.GroupControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) GroupControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *int {
		if v == nil {
			return nil
		}
		return &v.GroupControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) GroupHeaderAgencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return &v.GroupHeaderAgencyCode
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) GroupHeaderDateFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return &v.GroupHeaderDateFormat
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) GroupHeaderTimeFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return &v.GroupHeaderTimeFormat
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) GroupHeaderVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return &v.GroupHeaderVersion
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) InterchangeControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *int {
		if v == nil {
			return nil
		}
		return &v.InterchangeControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) InterchangeControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *int {
		if v == nil {
			return nil
		}
		return &v.InterchangeControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) OverwriteExistingTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.OverwriteExistingTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) ReceiverApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return &v.ReceiverApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) RolloverGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.RolloverGroupControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) RolloverInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.RolloverInterchangeControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) RolloverTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.RolloverTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return &v.SenderApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) TransactionSetControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *int {
		if v == nil {
			return nil
		}
		return &v.TransactionSetControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) TransactionSetControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) TransactionSetControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) TransactionSetControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *int {
		if v == nil {
			return nil
		}
		return &v.TransactionSetControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) UsageIndicator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *string {
		if v == nil {
			return nil
		}
		return &v.UsageIndicator
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsPtrOutput) UseControlStandardsIdAsRepetitionCharacter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.UseControlStandardsIdAsRepetitionCharacter
	}).(pulumi.BoolPtrOutput)
}

type X12EnvelopeSettingsResponse struct {
	ControlStandardsId                           int     `pulumi:"controlStandardsId"`
	ControlVersionNumber                         string  `pulumi:"controlVersionNumber"`
	EnableDefaultGroupHeaders                    bool    `pulumi:"enableDefaultGroupHeaders"`
	FunctionalGroupId                            *string `pulumi:"functionalGroupId"`
	GroupControlNumberLowerBound                 int     `pulumi:"groupControlNumberLowerBound"`
	GroupControlNumberUpperBound                 int     `pulumi:"groupControlNumberUpperBound"`
	GroupHeaderAgencyCode                        string  `pulumi:"groupHeaderAgencyCode"`
	GroupHeaderDateFormat                        string  `pulumi:"groupHeaderDateFormat"`
	GroupHeaderTimeFormat                        string  `pulumi:"groupHeaderTimeFormat"`
	GroupHeaderVersion                           string  `pulumi:"groupHeaderVersion"`
	InterchangeControlNumberLowerBound           int     `pulumi:"interchangeControlNumberLowerBound"`
	InterchangeControlNumberUpperBound           int     `pulumi:"interchangeControlNumberUpperBound"`
	OverwriteExistingTransactionSetControlNumber bool    `pulumi:"overwriteExistingTransactionSetControlNumber"`
	ReceiverApplicationId                        string  `pulumi:"receiverApplicationId"`
	RolloverGroupControlNumber                   bool    `pulumi:"rolloverGroupControlNumber"`
	RolloverInterchangeControlNumber             bool    `pulumi:"rolloverInterchangeControlNumber"`
	RolloverTransactionSetControlNumber          bool    `pulumi:"rolloverTransactionSetControlNumber"`
	SenderApplicationId                          string  `pulumi:"senderApplicationId"`
	TransactionSetControlNumberLowerBound        int     `pulumi:"transactionSetControlNumberLowerBound"`
	TransactionSetControlNumberPrefix            *string `pulumi:"transactionSetControlNumberPrefix"`
	TransactionSetControlNumberSuffix            *string `pulumi:"transactionSetControlNumberSuffix"`
	TransactionSetControlNumberUpperBound        int     `pulumi:"transactionSetControlNumberUpperBound"`
	UsageIndicator                               string  `pulumi:"usageIndicator"`
	UseControlStandardsIdAsRepetitionCharacter   bool    `pulumi:"useControlStandardsIdAsRepetitionCharacter"`
}

type X12EnvelopeSettingsResponseOutput struct{ *pulumi.OutputState }

func (X12EnvelopeSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12EnvelopeSettingsResponse)(nil)).Elem()
}

func (o X12EnvelopeSettingsResponseOutput) ToX12EnvelopeSettingsResponseOutput() X12EnvelopeSettingsResponseOutput {
	return o
}

func (o X12EnvelopeSettingsResponseOutput) ToX12EnvelopeSettingsResponseOutputWithContext(ctx context.Context) X12EnvelopeSettingsResponseOutput {
	return o
}

func (o X12EnvelopeSettingsResponseOutput) ControlStandardsId() pulumi.IntOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) int { return v.ControlStandardsId }).(pulumi.IntOutput)
}

func (o X12EnvelopeSettingsResponseOutput) ControlVersionNumber() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) string { return v.ControlVersionNumber }).(pulumi.StringOutput)
}

func (o X12EnvelopeSettingsResponseOutput) EnableDefaultGroupHeaders() pulumi.BoolOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) bool { return v.EnableDefaultGroupHeaders }).(pulumi.BoolOutput)
}

func (o X12EnvelopeSettingsResponseOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *string { return v.FunctionalGroupId }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) GroupControlNumberLowerBound() pulumi.IntOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) int { return v.GroupControlNumberLowerBound }).(pulumi.IntOutput)
}

func (o X12EnvelopeSettingsResponseOutput) GroupControlNumberUpperBound() pulumi.IntOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) int { return v.GroupControlNumberUpperBound }).(pulumi.IntOutput)
}

func (o X12EnvelopeSettingsResponseOutput) GroupHeaderAgencyCode() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) string { return v.GroupHeaderAgencyCode }).(pulumi.StringOutput)
}

func (o X12EnvelopeSettingsResponseOutput) GroupHeaderDateFormat() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) string { return v.GroupHeaderDateFormat }).(pulumi.StringOutput)
}

func (o X12EnvelopeSettingsResponseOutput) GroupHeaderTimeFormat() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) string { return v.GroupHeaderTimeFormat }).(pulumi.StringOutput)
}

func (o X12EnvelopeSettingsResponseOutput) GroupHeaderVersion() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) string { return v.GroupHeaderVersion }).(pulumi.StringOutput)
}

func (o X12EnvelopeSettingsResponseOutput) InterchangeControlNumberLowerBound() pulumi.IntOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) int { return v.InterchangeControlNumberLowerBound }).(pulumi.IntOutput)
}

func (o X12EnvelopeSettingsResponseOutput) InterchangeControlNumberUpperBound() pulumi.IntOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) int { return v.InterchangeControlNumberUpperBound }).(pulumi.IntOutput)
}

func (o X12EnvelopeSettingsResponseOutput) OverwriteExistingTransactionSetControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) bool { return v.OverwriteExistingTransactionSetControlNumber }).(pulumi.BoolOutput)
}

func (o X12EnvelopeSettingsResponseOutput) ReceiverApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) string { return v.ReceiverApplicationId }).(pulumi.StringOutput)
}

func (o X12EnvelopeSettingsResponseOutput) RolloverGroupControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) bool { return v.RolloverGroupControlNumber }).(pulumi.BoolOutput)
}

func (o X12EnvelopeSettingsResponseOutput) RolloverInterchangeControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) bool { return v.RolloverInterchangeControlNumber }).(pulumi.BoolOutput)
}

func (o X12EnvelopeSettingsResponseOutput) RolloverTransactionSetControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) bool { return v.RolloverTransactionSetControlNumber }).(pulumi.BoolOutput)
}

func (o X12EnvelopeSettingsResponseOutput) SenderApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) string { return v.SenderApplicationId }).(pulumi.StringOutput)
}

func (o X12EnvelopeSettingsResponseOutput) TransactionSetControlNumberLowerBound() pulumi.IntOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) int { return v.TransactionSetControlNumberLowerBound }).(pulumi.IntOutput)
}

func (o X12EnvelopeSettingsResponseOutput) TransactionSetControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *string { return v.TransactionSetControlNumberPrefix }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) TransactionSetControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) *string { return v.TransactionSetControlNumberSuffix }).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponseOutput) TransactionSetControlNumberUpperBound() pulumi.IntOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) int { return v.TransactionSetControlNumberUpperBound }).(pulumi.IntOutput)
}

func (o X12EnvelopeSettingsResponseOutput) UsageIndicator() pulumi.StringOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) string { return v.UsageIndicator }).(pulumi.StringOutput)
}

func (o X12EnvelopeSettingsResponseOutput) UseControlStandardsIdAsRepetitionCharacter() pulumi.BoolOutput {
	return o.ApplyT(func(v X12EnvelopeSettingsResponse) bool { return v.UseControlStandardsIdAsRepetitionCharacter }).(pulumi.BoolOutput)
}

type X12EnvelopeSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (X12EnvelopeSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12EnvelopeSettingsResponse)(nil)).Elem()
}

func (o X12EnvelopeSettingsResponsePtrOutput) ToX12EnvelopeSettingsResponsePtrOutput() X12EnvelopeSettingsResponsePtrOutput {
	return o
}

func (o X12EnvelopeSettingsResponsePtrOutput) ToX12EnvelopeSettingsResponsePtrOutputWithContext(ctx context.Context) X12EnvelopeSettingsResponsePtrOutput {
	return o
}

func (o X12EnvelopeSettingsResponsePtrOutput) Elem() X12EnvelopeSettingsResponseOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) X12EnvelopeSettingsResponse {
		if v != nil {
			return *v
		}
		var ret X12EnvelopeSettingsResponse
		return ret
	}).(X12EnvelopeSettingsResponseOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) ControlStandardsId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ControlStandardsId
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) ControlVersionNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ControlVersionNumber
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) EnableDefaultGroupHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.EnableDefaultGroupHeaders
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) FunctionalGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.FunctionalGroupId
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) GroupControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.GroupControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) GroupControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.GroupControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) GroupHeaderAgencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.GroupHeaderAgencyCode
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) GroupHeaderDateFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.GroupHeaderDateFormat
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) GroupHeaderTimeFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.GroupHeaderTimeFormat
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) GroupHeaderVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.GroupHeaderVersion
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) InterchangeControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.InterchangeControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) InterchangeControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.InterchangeControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) OverwriteExistingTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.OverwriteExistingTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) ReceiverApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ReceiverApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) RolloverGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.RolloverGroupControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) RolloverInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.RolloverInterchangeControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) RolloverTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.RolloverTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SenderApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) TransactionSetControlNumberLowerBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.TransactionSetControlNumberLowerBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) TransactionSetControlNumberPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberPrefix
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) TransactionSetControlNumberSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TransactionSetControlNumberSuffix
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) TransactionSetControlNumberUpperBound() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.TransactionSetControlNumberUpperBound
	}).(pulumi.IntPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) UsageIndicator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UsageIndicator
	}).(pulumi.StringPtrOutput)
}

func (o X12EnvelopeSettingsResponsePtrOutput) UseControlStandardsIdAsRepetitionCharacter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12EnvelopeSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.UseControlStandardsIdAsRepetitionCharacter
	}).(pulumi.BoolPtrOutput)
}

type X12FramingSettings struct {
	CharacterSet               string                  `pulumi:"characterSet"`
	ComponentSeparator         int                     `pulumi:"componentSeparator"`
	DataElementSeparator       int                     `pulumi:"dataElementSeparator"`
	ReplaceCharacter           int                     `pulumi:"replaceCharacter"`
	ReplaceSeparatorsInPayload bool                    `pulumi:"replaceSeparatorsInPayload"`
	SegmentTerminator          int                     `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix    SegmentTerminatorSuffix `pulumi:"segmentTerminatorSuffix"`
}





type X12FramingSettingsInput interface {
	pulumi.Input

	ToX12FramingSettingsOutput() X12FramingSettingsOutput
	ToX12FramingSettingsOutputWithContext(context.Context) X12FramingSettingsOutput
}

type X12FramingSettingsArgs struct {
	CharacterSet               pulumi.StringInput           `pulumi:"characterSet"`
	ComponentSeparator         pulumi.IntInput              `pulumi:"componentSeparator"`
	DataElementSeparator       pulumi.IntInput              `pulumi:"dataElementSeparator"`
	ReplaceCharacter           pulumi.IntInput              `pulumi:"replaceCharacter"`
	ReplaceSeparatorsInPayload pulumi.BoolInput             `pulumi:"replaceSeparatorsInPayload"`
	SegmentTerminator          pulumi.IntInput              `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix    SegmentTerminatorSuffixInput `pulumi:"segmentTerminatorSuffix"`
}

func (X12FramingSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12FramingSettings)(nil)).Elem()
}

func (i X12FramingSettingsArgs) ToX12FramingSettingsOutput() X12FramingSettingsOutput {
	return i.ToX12FramingSettingsOutputWithContext(context.Background())
}

func (i X12FramingSettingsArgs) ToX12FramingSettingsOutputWithContext(ctx context.Context) X12FramingSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12FramingSettingsOutput)
}

func (i X12FramingSettingsArgs) ToX12FramingSettingsPtrOutput() X12FramingSettingsPtrOutput {
	return i.ToX12FramingSettingsPtrOutputWithContext(context.Background())
}

func (i X12FramingSettingsArgs) ToX12FramingSettingsPtrOutputWithContext(ctx context.Context) X12FramingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12FramingSettingsOutput).ToX12FramingSettingsPtrOutputWithContext(ctx)
}









type X12FramingSettingsPtrInput interface {
	pulumi.Input

	ToX12FramingSettingsPtrOutput() X12FramingSettingsPtrOutput
	ToX12FramingSettingsPtrOutputWithContext(context.Context) X12FramingSettingsPtrOutput
}

type x12framingSettingsPtrType X12FramingSettingsArgs

func X12FramingSettingsPtr(v *X12FramingSettingsArgs) X12FramingSettingsPtrInput {
	return (*x12framingSettingsPtrType)(v)
}

func (*x12framingSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12FramingSettings)(nil)).Elem()
}

func (i *x12framingSettingsPtrType) ToX12FramingSettingsPtrOutput() X12FramingSettingsPtrOutput {
	return i.ToX12FramingSettingsPtrOutputWithContext(context.Background())
}

func (i *x12framingSettingsPtrType) ToX12FramingSettingsPtrOutputWithContext(ctx context.Context) X12FramingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12FramingSettingsPtrOutput)
}

type X12FramingSettingsOutput struct{ *pulumi.OutputState }

func (X12FramingSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12FramingSettings)(nil)).Elem()
}

func (o X12FramingSettingsOutput) ToX12FramingSettingsOutput() X12FramingSettingsOutput {
	return o
}

func (o X12FramingSettingsOutput) ToX12FramingSettingsOutputWithContext(ctx context.Context) X12FramingSettingsOutput {
	return o
}

func (o X12FramingSettingsOutput) ToX12FramingSettingsPtrOutput() X12FramingSettingsPtrOutput {
	return o.ToX12FramingSettingsPtrOutputWithContext(context.Background())
}

func (o X12FramingSettingsOutput) ToX12FramingSettingsPtrOutputWithContext(ctx context.Context) X12FramingSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12FramingSettings) *X12FramingSettings {
		return &v
	}).(X12FramingSettingsPtrOutput)
}

func (o X12FramingSettingsOutput) CharacterSet() pulumi.StringOutput {
	return o.ApplyT(func(v X12FramingSettings) string { return v.CharacterSet }).(pulumi.StringOutput)
}

func (o X12FramingSettingsOutput) ComponentSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v X12FramingSettings) int { return v.ComponentSeparator }).(pulumi.IntOutput)
}

func (o X12FramingSettingsOutput) DataElementSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v X12FramingSettings) int { return v.DataElementSeparator }).(pulumi.IntOutput)
}

func (o X12FramingSettingsOutput) ReplaceCharacter() pulumi.IntOutput {
	return o.ApplyT(func(v X12FramingSettings) int { return v.ReplaceCharacter }).(pulumi.IntOutput)
}

func (o X12FramingSettingsOutput) ReplaceSeparatorsInPayload() pulumi.BoolOutput {
	return o.ApplyT(func(v X12FramingSettings) bool { return v.ReplaceSeparatorsInPayload }).(pulumi.BoolOutput)
}

func (o X12FramingSettingsOutput) SegmentTerminator() pulumi.IntOutput {
	return o.ApplyT(func(v X12FramingSettings) int { return v.SegmentTerminator }).(pulumi.IntOutput)
}

func (o X12FramingSettingsOutput) SegmentTerminatorSuffix() SegmentTerminatorSuffixOutput {
	return o.ApplyT(func(v X12FramingSettings) SegmentTerminatorSuffix { return v.SegmentTerminatorSuffix }).(SegmentTerminatorSuffixOutput)
}

type X12FramingSettingsPtrOutput struct{ *pulumi.OutputState }

func (X12FramingSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12FramingSettings)(nil)).Elem()
}

func (o X12FramingSettingsPtrOutput) ToX12FramingSettingsPtrOutput() X12FramingSettingsPtrOutput {
	return o
}

func (o X12FramingSettingsPtrOutput) ToX12FramingSettingsPtrOutputWithContext(ctx context.Context) X12FramingSettingsPtrOutput {
	return o
}

func (o X12FramingSettingsPtrOutput) Elem() X12FramingSettingsOutput {
	return o.ApplyT(func(v *X12FramingSettings) X12FramingSettings {
		if v != nil {
			return *v
		}
		var ret X12FramingSettings
		return ret
	}).(X12FramingSettingsOutput)
}

func (o X12FramingSettingsPtrOutput) CharacterSet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12FramingSettings) *string {
		if v == nil {
			return nil
		}
		return &v.CharacterSet
	}).(pulumi.StringPtrOutput)
}

func (o X12FramingSettingsPtrOutput) ComponentSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12FramingSettings) *int {
		if v == nil {
			return nil
		}
		return &v.ComponentSeparator
	}).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsPtrOutput) DataElementSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12FramingSettings) *int {
		if v == nil {
			return nil
		}
		return &v.DataElementSeparator
	}).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsPtrOutput) ReplaceCharacter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12FramingSettings) *int {
		if v == nil {
			return nil
		}
		return &v.ReplaceCharacter
	}).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsPtrOutput) ReplaceSeparatorsInPayload() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12FramingSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.ReplaceSeparatorsInPayload
	}).(pulumi.BoolPtrOutput)
}

func (o X12FramingSettingsPtrOutput) SegmentTerminator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12FramingSettings) *int {
		if v == nil {
			return nil
		}
		return &v.SegmentTerminator
	}).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsPtrOutput) SegmentTerminatorSuffix() SegmentTerminatorSuffixPtrOutput {
	return o.ApplyT(func(v *X12FramingSettings) *SegmentTerminatorSuffix {
		if v == nil {
			return nil
		}
		return &v.SegmentTerminatorSuffix
	}).(SegmentTerminatorSuffixPtrOutput)
}

type X12FramingSettingsResponse struct {
	CharacterSet               string `pulumi:"characterSet"`
	ComponentSeparator         int    `pulumi:"componentSeparator"`
	DataElementSeparator       int    `pulumi:"dataElementSeparator"`
	ReplaceCharacter           int    `pulumi:"replaceCharacter"`
	ReplaceSeparatorsInPayload bool   `pulumi:"replaceSeparatorsInPayload"`
	SegmentTerminator          int    `pulumi:"segmentTerminator"`
	SegmentTerminatorSuffix    string `pulumi:"segmentTerminatorSuffix"`
}

type X12FramingSettingsResponseOutput struct{ *pulumi.OutputState }

func (X12FramingSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12FramingSettingsResponse)(nil)).Elem()
}

func (o X12FramingSettingsResponseOutput) ToX12FramingSettingsResponseOutput() X12FramingSettingsResponseOutput {
	return o
}

func (o X12FramingSettingsResponseOutput) ToX12FramingSettingsResponseOutputWithContext(ctx context.Context) X12FramingSettingsResponseOutput {
	return o
}

func (o X12FramingSettingsResponseOutput) CharacterSet() pulumi.StringOutput {
	return o.ApplyT(func(v X12FramingSettingsResponse) string { return v.CharacterSet }).(pulumi.StringOutput)
}

func (o X12FramingSettingsResponseOutput) ComponentSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v X12FramingSettingsResponse) int { return v.ComponentSeparator }).(pulumi.IntOutput)
}

func (o X12FramingSettingsResponseOutput) DataElementSeparator() pulumi.IntOutput {
	return o.ApplyT(func(v X12FramingSettingsResponse) int { return v.DataElementSeparator }).(pulumi.IntOutput)
}

func (o X12FramingSettingsResponseOutput) ReplaceCharacter() pulumi.IntOutput {
	return o.ApplyT(func(v X12FramingSettingsResponse) int { return v.ReplaceCharacter }).(pulumi.IntOutput)
}

func (o X12FramingSettingsResponseOutput) ReplaceSeparatorsInPayload() pulumi.BoolOutput {
	return o.ApplyT(func(v X12FramingSettingsResponse) bool { return v.ReplaceSeparatorsInPayload }).(pulumi.BoolOutput)
}

func (o X12FramingSettingsResponseOutput) SegmentTerminator() pulumi.IntOutput {
	return o.ApplyT(func(v X12FramingSettingsResponse) int { return v.SegmentTerminator }).(pulumi.IntOutput)
}

func (o X12FramingSettingsResponseOutput) SegmentTerminatorSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v X12FramingSettingsResponse) string { return v.SegmentTerminatorSuffix }).(pulumi.StringOutput)
}

type X12FramingSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (X12FramingSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12FramingSettingsResponse)(nil)).Elem()
}

func (o X12FramingSettingsResponsePtrOutput) ToX12FramingSettingsResponsePtrOutput() X12FramingSettingsResponsePtrOutput {
	return o
}

func (o X12FramingSettingsResponsePtrOutput) ToX12FramingSettingsResponsePtrOutputWithContext(ctx context.Context) X12FramingSettingsResponsePtrOutput {
	return o
}

func (o X12FramingSettingsResponsePtrOutput) Elem() X12FramingSettingsResponseOutput {
	return o.ApplyT(func(v *X12FramingSettingsResponse) X12FramingSettingsResponse {
		if v != nil {
			return *v
		}
		var ret X12FramingSettingsResponse
		return ret
	}).(X12FramingSettingsResponseOutput)
}

func (o X12FramingSettingsResponsePtrOutput) CharacterSet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12FramingSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CharacterSet
	}).(pulumi.StringPtrOutput)
}

func (o X12FramingSettingsResponsePtrOutput) ComponentSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12FramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ComponentSeparator
	}).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsResponsePtrOutput) DataElementSeparator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12FramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.DataElementSeparator
	}).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsResponsePtrOutput) ReplaceCharacter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12FramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ReplaceCharacter
	}).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsResponsePtrOutput) ReplaceSeparatorsInPayload() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12FramingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.ReplaceSeparatorsInPayload
	}).(pulumi.BoolPtrOutput)
}

func (o X12FramingSettingsResponsePtrOutput) SegmentTerminator() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12FramingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.SegmentTerminator
	}).(pulumi.IntPtrOutput)
}

func (o X12FramingSettingsResponsePtrOutput) SegmentTerminatorSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12FramingSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SegmentTerminatorSuffix
	}).(pulumi.StringPtrOutput)
}

type X12MessageFilter struct {
	MessageFilterType string `pulumi:"messageFilterType"`
}





type X12MessageFilterInput interface {
	pulumi.Input

	ToX12MessageFilterOutput() X12MessageFilterOutput
	ToX12MessageFilterOutputWithContext(context.Context) X12MessageFilterOutput
}

type X12MessageFilterArgs struct {
	MessageFilterType pulumi.StringInput `pulumi:"messageFilterType"`
}

func (X12MessageFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12MessageFilter)(nil)).Elem()
}

func (i X12MessageFilterArgs) ToX12MessageFilterOutput() X12MessageFilterOutput {
	return i.ToX12MessageFilterOutputWithContext(context.Background())
}

func (i X12MessageFilterArgs) ToX12MessageFilterOutputWithContext(ctx context.Context) X12MessageFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12MessageFilterOutput)
}

func (i X12MessageFilterArgs) ToX12MessageFilterPtrOutput() X12MessageFilterPtrOutput {
	return i.ToX12MessageFilterPtrOutputWithContext(context.Background())
}

func (i X12MessageFilterArgs) ToX12MessageFilterPtrOutputWithContext(ctx context.Context) X12MessageFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12MessageFilterOutput).ToX12MessageFilterPtrOutputWithContext(ctx)
}









type X12MessageFilterPtrInput interface {
	pulumi.Input

	ToX12MessageFilterPtrOutput() X12MessageFilterPtrOutput
	ToX12MessageFilterPtrOutputWithContext(context.Context) X12MessageFilterPtrOutput
}

type x12messageFilterPtrType X12MessageFilterArgs

func X12MessageFilterPtr(v *X12MessageFilterArgs) X12MessageFilterPtrInput {
	return (*x12messageFilterPtrType)(v)
}

func (*x12messageFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12MessageFilter)(nil)).Elem()
}

func (i *x12messageFilterPtrType) ToX12MessageFilterPtrOutput() X12MessageFilterPtrOutput {
	return i.ToX12MessageFilterPtrOutputWithContext(context.Background())
}

func (i *x12messageFilterPtrType) ToX12MessageFilterPtrOutputWithContext(ctx context.Context) X12MessageFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12MessageFilterPtrOutput)
}

type X12MessageFilterOutput struct{ *pulumi.OutputState }

func (X12MessageFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12MessageFilter)(nil)).Elem()
}

func (o X12MessageFilterOutput) ToX12MessageFilterOutput() X12MessageFilterOutput {
	return o
}

func (o X12MessageFilterOutput) ToX12MessageFilterOutputWithContext(ctx context.Context) X12MessageFilterOutput {
	return o
}

func (o X12MessageFilterOutput) ToX12MessageFilterPtrOutput() X12MessageFilterPtrOutput {
	return o.ToX12MessageFilterPtrOutputWithContext(context.Background())
}

func (o X12MessageFilterOutput) ToX12MessageFilterPtrOutputWithContext(ctx context.Context) X12MessageFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12MessageFilter) *X12MessageFilter {
		return &v
	}).(X12MessageFilterPtrOutput)
}

func (o X12MessageFilterOutput) MessageFilterType() pulumi.StringOutput {
	return o.ApplyT(func(v X12MessageFilter) string { return v.MessageFilterType }).(pulumi.StringOutput)
}

type X12MessageFilterPtrOutput struct{ *pulumi.OutputState }

func (X12MessageFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12MessageFilter)(nil)).Elem()
}

func (o X12MessageFilterPtrOutput) ToX12MessageFilterPtrOutput() X12MessageFilterPtrOutput {
	return o
}

func (o X12MessageFilterPtrOutput) ToX12MessageFilterPtrOutputWithContext(ctx context.Context) X12MessageFilterPtrOutput {
	return o
}

func (o X12MessageFilterPtrOutput) Elem() X12MessageFilterOutput {
	return o.ApplyT(func(v *X12MessageFilter) X12MessageFilter {
		if v != nil {
			return *v
		}
		var ret X12MessageFilter
		return ret
	}).(X12MessageFilterOutput)
}

func (o X12MessageFilterPtrOutput) MessageFilterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12MessageFilter) *string {
		if v == nil {
			return nil
		}
		return &v.MessageFilterType
	}).(pulumi.StringPtrOutput)
}

type X12MessageFilterResponse struct {
	MessageFilterType string `pulumi:"messageFilterType"`
}

type X12MessageFilterResponseOutput struct{ *pulumi.OutputState }

func (X12MessageFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12MessageFilterResponse)(nil)).Elem()
}

func (o X12MessageFilterResponseOutput) ToX12MessageFilterResponseOutput() X12MessageFilterResponseOutput {
	return o
}

func (o X12MessageFilterResponseOutput) ToX12MessageFilterResponseOutputWithContext(ctx context.Context) X12MessageFilterResponseOutput {
	return o
}

func (o X12MessageFilterResponseOutput) MessageFilterType() pulumi.StringOutput {
	return o.ApplyT(func(v X12MessageFilterResponse) string { return v.MessageFilterType }).(pulumi.StringOutput)
}

type X12MessageFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (X12MessageFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12MessageFilterResponse)(nil)).Elem()
}

func (o X12MessageFilterResponsePtrOutput) ToX12MessageFilterResponsePtrOutput() X12MessageFilterResponsePtrOutput {
	return o
}

func (o X12MessageFilterResponsePtrOutput) ToX12MessageFilterResponsePtrOutputWithContext(ctx context.Context) X12MessageFilterResponsePtrOutput {
	return o
}

func (o X12MessageFilterResponsePtrOutput) Elem() X12MessageFilterResponseOutput {
	return o.ApplyT(func(v *X12MessageFilterResponse) X12MessageFilterResponse {
		if v != nil {
			return *v
		}
		var ret X12MessageFilterResponse
		return ret
	}).(X12MessageFilterResponseOutput)
}

func (o X12MessageFilterResponsePtrOutput) MessageFilterType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12MessageFilterResponse) *string {
		if v == nil {
			return nil
		}
		return &v.MessageFilterType
	}).(pulumi.StringPtrOutput)
}

type X12MessageIdentifier struct {
	MessageId string `pulumi:"messageId"`
}





type X12MessageIdentifierInput interface {
	pulumi.Input

	ToX12MessageIdentifierOutput() X12MessageIdentifierOutput
	ToX12MessageIdentifierOutputWithContext(context.Context) X12MessageIdentifierOutput
}

type X12MessageIdentifierArgs struct {
	MessageId pulumi.StringInput `pulumi:"messageId"`
}

func (X12MessageIdentifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12MessageIdentifier)(nil)).Elem()
}

func (i X12MessageIdentifierArgs) ToX12MessageIdentifierOutput() X12MessageIdentifierOutput {
	return i.ToX12MessageIdentifierOutputWithContext(context.Background())
}

func (i X12MessageIdentifierArgs) ToX12MessageIdentifierOutputWithContext(ctx context.Context) X12MessageIdentifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12MessageIdentifierOutput)
}





type X12MessageIdentifierArrayInput interface {
	pulumi.Input

	ToX12MessageIdentifierArrayOutput() X12MessageIdentifierArrayOutput
	ToX12MessageIdentifierArrayOutputWithContext(context.Context) X12MessageIdentifierArrayOutput
}

type X12MessageIdentifierArray []X12MessageIdentifierInput

func (X12MessageIdentifierArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12MessageIdentifier)(nil)).Elem()
}

func (i X12MessageIdentifierArray) ToX12MessageIdentifierArrayOutput() X12MessageIdentifierArrayOutput {
	return i.ToX12MessageIdentifierArrayOutputWithContext(context.Background())
}

func (i X12MessageIdentifierArray) ToX12MessageIdentifierArrayOutputWithContext(ctx context.Context) X12MessageIdentifierArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12MessageIdentifierArrayOutput)
}

type X12MessageIdentifierOutput struct{ *pulumi.OutputState }

func (X12MessageIdentifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12MessageIdentifier)(nil)).Elem()
}

func (o X12MessageIdentifierOutput) ToX12MessageIdentifierOutput() X12MessageIdentifierOutput {
	return o
}

func (o X12MessageIdentifierOutput) ToX12MessageIdentifierOutputWithContext(ctx context.Context) X12MessageIdentifierOutput {
	return o
}

func (o X12MessageIdentifierOutput) MessageId() pulumi.StringOutput {
	return o.ApplyT(func(v X12MessageIdentifier) string { return v.MessageId }).(pulumi.StringOutput)
}

type X12MessageIdentifierArrayOutput struct{ *pulumi.OutputState }

func (X12MessageIdentifierArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12MessageIdentifier)(nil)).Elem()
}

func (o X12MessageIdentifierArrayOutput) ToX12MessageIdentifierArrayOutput() X12MessageIdentifierArrayOutput {
	return o
}

func (o X12MessageIdentifierArrayOutput) ToX12MessageIdentifierArrayOutputWithContext(ctx context.Context) X12MessageIdentifierArrayOutput {
	return o
}

func (o X12MessageIdentifierArrayOutput) Index(i pulumi.IntInput) X12MessageIdentifierOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12MessageIdentifier {
		return vs[0].([]X12MessageIdentifier)[vs[1].(int)]
	}).(X12MessageIdentifierOutput)
}

type X12MessageIdentifierResponse struct {
	MessageId string `pulumi:"messageId"`
}

type X12MessageIdentifierResponseOutput struct{ *pulumi.OutputState }

func (X12MessageIdentifierResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12MessageIdentifierResponse)(nil)).Elem()
}

func (o X12MessageIdentifierResponseOutput) ToX12MessageIdentifierResponseOutput() X12MessageIdentifierResponseOutput {
	return o
}

func (o X12MessageIdentifierResponseOutput) ToX12MessageIdentifierResponseOutputWithContext(ctx context.Context) X12MessageIdentifierResponseOutput {
	return o
}

func (o X12MessageIdentifierResponseOutput) MessageId() pulumi.StringOutput {
	return o.ApplyT(func(v X12MessageIdentifierResponse) string { return v.MessageId }).(pulumi.StringOutput)
}

type X12MessageIdentifierResponseArrayOutput struct{ *pulumi.OutputState }

func (X12MessageIdentifierResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12MessageIdentifierResponse)(nil)).Elem()
}

func (o X12MessageIdentifierResponseArrayOutput) ToX12MessageIdentifierResponseArrayOutput() X12MessageIdentifierResponseArrayOutput {
	return o
}

func (o X12MessageIdentifierResponseArrayOutput) ToX12MessageIdentifierResponseArrayOutputWithContext(ctx context.Context) X12MessageIdentifierResponseArrayOutput {
	return o
}

func (o X12MessageIdentifierResponseArrayOutput) Index(i pulumi.IntInput) X12MessageIdentifierResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12MessageIdentifierResponse {
		return vs[0].([]X12MessageIdentifierResponse)[vs[1].(int)]
	}).(X12MessageIdentifierResponseOutput)
}

type X12OneWayAgreement struct {
	ProtocolSettings         X12ProtocolSettings `pulumi:"protocolSettings"`
	ReceiverBusinessIdentity BusinessIdentity    `pulumi:"receiverBusinessIdentity"`
	SenderBusinessIdentity   BusinessIdentity    `pulumi:"senderBusinessIdentity"`
}





type X12OneWayAgreementInput interface {
	pulumi.Input

	ToX12OneWayAgreementOutput() X12OneWayAgreementOutput
	ToX12OneWayAgreementOutputWithContext(context.Context) X12OneWayAgreementOutput
}

type X12OneWayAgreementArgs struct {
	ProtocolSettings         X12ProtocolSettingsInput `pulumi:"protocolSettings"`
	ReceiverBusinessIdentity BusinessIdentityInput    `pulumi:"receiverBusinessIdentity"`
	SenderBusinessIdentity   BusinessIdentityInput    `pulumi:"senderBusinessIdentity"`
}

func (X12OneWayAgreementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12OneWayAgreement)(nil)).Elem()
}

func (i X12OneWayAgreementArgs) ToX12OneWayAgreementOutput() X12OneWayAgreementOutput {
	return i.ToX12OneWayAgreementOutputWithContext(context.Background())
}

func (i X12OneWayAgreementArgs) ToX12OneWayAgreementOutputWithContext(ctx context.Context) X12OneWayAgreementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12OneWayAgreementOutput)
}

func (i X12OneWayAgreementArgs) ToX12OneWayAgreementPtrOutput() X12OneWayAgreementPtrOutput {
	return i.ToX12OneWayAgreementPtrOutputWithContext(context.Background())
}

func (i X12OneWayAgreementArgs) ToX12OneWayAgreementPtrOutputWithContext(ctx context.Context) X12OneWayAgreementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12OneWayAgreementOutput).ToX12OneWayAgreementPtrOutputWithContext(ctx)
}









type X12OneWayAgreementPtrInput interface {
	pulumi.Input

	ToX12OneWayAgreementPtrOutput() X12OneWayAgreementPtrOutput
	ToX12OneWayAgreementPtrOutputWithContext(context.Context) X12OneWayAgreementPtrOutput
}

type x12oneWayAgreementPtrType X12OneWayAgreementArgs

func X12OneWayAgreementPtr(v *X12OneWayAgreementArgs) X12OneWayAgreementPtrInput {
	return (*x12oneWayAgreementPtrType)(v)
}

func (*x12oneWayAgreementPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12OneWayAgreement)(nil)).Elem()
}

func (i *x12oneWayAgreementPtrType) ToX12OneWayAgreementPtrOutput() X12OneWayAgreementPtrOutput {
	return i.ToX12OneWayAgreementPtrOutputWithContext(context.Background())
}

func (i *x12oneWayAgreementPtrType) ToX12OneWayAgreementPtrOutputWithContext(ctx context.Context) X12OneWayAgreementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12OneWayAgreementPtrOutput)
}

type X12OneWayAgreementOutput struct{ *pulumi.OutputState }

func (X12OneWayAgreementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12OneWayAgreement)(nil)).Elem()
}

func (o X12OneWayAgreementOutput) ToX12OneWayAgreementOutput() X12OneWayAgreementOutput {
	return o
}

func (o X12OneWayAgreementOutput) ToX12OneWayAgreementOutputWithContext(ctx context.Context) X12OneWayAgreementOutput {
	return o
}

func (o X12OneWayAgreementOutput) ToX12OneWayAgreementPtrOutput() X12OneWayAgreementPtrOutput {
	return o.ToX12OneWayAgreementPtrOutputWithContext(context.Background())
}

func (o X12OneWayAgreementOutput) ToX12OneWayAgreementPtrOutputWithContext(ctx context.Context) X12OneWayAgreementPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12OneWayAgreement) *X12OneWayAgreement {
		return &v
	}).(X12OneWayAgreementPtrOutput)
}

func (o X12OneWayAgreementOutput) ProtocolSettings() X12ProtocolSettingsOutput {
	return o.ApplyT(func(v X12OneWayAgreement) X12ProtocolSettings { return v.ProtocolSettings }).(X12ProtocolSettingsOutput)
}

func (o X12OneWayAgreementOutput) ReceiverBusinessIdentity() BusinessIdentityOutput {
	return o.ApplyT(func(v X12OneWayAgreement) BusinessIdentity { return v.ReceiverBusinessIdentity }).(BusinessIdentityOutput)
}

func (o X12OneWayAgreementOutput) SenderBusinessIdentity() BusinessIdentityOutput {
	return o.ApplyT(func(v X12OneWayAgreement) BusinessIdentity { return v.SenderBusinessIdentity }).(BusinessIdentityOutput)
}

type X12OneWayAgreementPtrOutput struct{ *pulumi.OutputState }

func (X12OneWayAgreementPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12OneWayAgreement)(nil)).Elem()
}

func (o X12OneWayAgreementPtrOutput) ToX12OneWayAgreementPtrOutput() X12OneWayAgreementPtrOutput {
	return o
}

func (o X12OneWayAgreementPtrOutput) ToX12OneWayAgreementPtrOutputWithContext(ctx context.Context) X12OneWayAgreementPtrOutput {
	return o
}

func (o X12OneWayAgreementPtrOutput) Elem() X12OneWayAgreementOutput {
	return o.ApplyT(func(v *X12OneWayAgreement) X12OneWayAgreement {
		if v != nil {
			return *v
		}
		var ret X12OneWayAgreement
		return ret
	}).(X12OneWayAgreementOutput)
}

func (o X12OneWayAgreementPtrOutput) ProtocolSettings() X12ProtocolSettingsPtrOutput {
	return o.ApplyT(func(v *X12OneWayAgreement) *X12ProtocolSettings {
		if v == nil {
			return nil
		}
		return &v.ProtocolSettings
	}).(X12ProtocolSettingsPtrOutput)
}

func (o X12OneWayAgreementPtrOutput) ReceiverBusinessIdentity() BusinessIdentityPtrOutput {
	return o.ApplyT(func(v *X12OneWayAgreement) *BusinessIdentity {
		if v == nil {
			return nil
		}
		return &v.ReceiverBusinessIdentity
	}).(BusinessIdentityPtrOutput)
}

func (o X12OneWayAgreementPtrOutput) SenderBusinessIdentity() BusinessIdentityPtrOutput {
	return o.ApplyT(func(v *X12OneWayAgreement) *BusinessIdentity {
		if v == nil {
			return nil
		}
		return &v.SenderBusinessIdentity
	}).(BusinessIdentityPtrOutput)
}

type X12OneWayAgreementResponse struct {
	ProtocolSettings         X12ProtocolSettingsResponse `pulumi:"protocolSettings"`
	ReceiverBusinessIdentity BusinessIdentityResponse    `pulumi:"receiverBusinessIdentity"`
	SenderBusinessIdentity   BusinessIdentityResponse    `pulumi:"senderBusinessIdentity"`
}

type X12OneWayAgreementResponseOutput struct{ *pulumi.OutputState }

func (X12OneWayAgreementResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12OneWayAgreementResponse)(nil)).Elem()
}

func (o X12OneWayAgreementResponseOutput) ToX12OneWayAgreementResponseOutput() X12OneWayAgreementResponseOutput {
	return o
}

func (o X12OneWayAgreementResponseOutput) ToX12OneWayAgreementResponseOutputWithContext(ctx context.Context) X12OneWayAgreementResponseOutput {
	return o
}

func (o X12OneWayAgreementResponseOutput) ProtocolSettings() X12ProtocolSettingsResponseOutput {
	return o.ApplyT(func(v X12OneWayAgreementResponse) X12ProtocolSettingsResponse { return v.ProtocolSettings }).(X12ProtocolSettingsResponseOutput)
}

func (o X12OneWayAgreementResponseOutput) ReceiverBusinessIdentity() BusinessIdentityResponseOutput {
	return o.ApplyT(func(v X12OneWayAgreementResponse) BusinessIdentityResponse { return v.ReceiverBusinessIdentity }).(BusinessIdentityResponseOutput)
}

func (o X12OneWayAgreementResponseOutput) SenderBusinessIdentity() BusinessIdentityResponseOutput {
	return o.ApplyT(func(v X12OneWayAgreementResponse) BusinessIdentityResponse { return v.SenderBusinessIdentity }).(BusinessIdentityResponseOutput)
}

type X12OneWayAgreementResponsePtrOutput struct{ *pulumi.OutputState }

func (X12OneWayAgreementResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12OneWayAgreementResponse)(nil)).Elem()
}

func (o X12OneWayAgreementResponsePtrOutput) ToX12OneWayAgreementResponsePtrOutput() X12OneWayAgreementResponsePtrOutput {
	return o
}

func (o X12OneWayAgreementResponsePtrOutput) ToX12OneWayAgreementResponsePtrOutputWithContext(ctx context.Context) X12OneWayAgreementResponsePtrOutput {
	return o
}

func (o X12OneWayAgreementResponsePtrOutput) Elem() X12OneWayAgreementResponseOutput {
	return o.ApplyT(func(v *X12OneWayAgreementResponse) X12OneWayAgreementResponse {
		if v != nil {
			return *v
		}
		var ret X12OneWayAgreementResponse
		return ret
	}).(X12OneWayAgreementResponseOutput)
}

func (o X12OneWayAgreementResponsePtrOutput) ProtocolSettings() X12ProtocolSettingsResponsePtrOutput {
	return o.ApplyT(func(v *X12OneWayAgreementResponse) *X12ProtocolSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.ProtocolSettings
	}).(X12ProtocolSettingsResponsePtrOutput)
}

func (o X12OneWayAgreementResponsePtrOutput) ReceiverBusinessIdentity() BusinessIdentityResponsePtrOutput {
	return o.ApplyT(func(v *X12OneWayAgreementResponse) *BusinessIdentityResponse {
		if v == nil {
			return nil
		}
		return &v.ReceiverBusinessIdentity
	}).(BusinessIdentityResponsePtrOutput)
}

func (o X12OneWayAgreementResponsePtrOutput) SenderBusinessIdentity() BusinessIdentityResponsePtrOutput {
	return o.ApplyT(func(v *X12OneWayAgreementResponse) *BusinessIdentityResponse {
		if v == nil {
			return nil
		}
		return &v.SenderBusinessIdentity
	}).(BusinessIdentityResponsePtrOutput)
}

type X12ProcessingSettings struct {
	ConvertImpliedDecimal                   bool `pulumi:"convertImpliedDecimal"`
	CreateEmptyXmlTagsForTrailingSeparators bool `pulumi:"createEmptyXmlTagsForTrailingSeparators"`
	MaskSecurityInfo                        bool `pulumi:"maskSecurityInfo"`
	PreserveInterchange                     bool `pulumi:"preserveInterchange"`
	SuspendInterchangeOnError               bool `pulumi:"suspendInterchangeOnError"`
	UseDotAsDecimalSeparator                bool `pulumi:"useDotAsDecimalSeparator"`
}





type X12ProcessingSettingsInput interface {
	pulumi.Input

	ToX12ProcessingSettingsOutput() X12ProcessingSettingsOutput
	ToX12ProcessingSettingsOutputWithContext(context.Context) X12ProcessingSettingsOutput
}

type X12ProcessingSettingsArgs struct {
	ConvertImpliedDecimal                   pulumi.BoolInput `pulumi:"convertImpliedDecimal"`
	CreateEmptyXmlTagsForTrailingSeparators pulumi.BoolInput `pulumi:"createEmptyXmlTagsForTrailingSeparators"`
	MaskSecurityInfo                        pulumi.BoolInput `pulumi:"maskSecurityInfo"`
	PreserveInterchange                     pulumi.BoolInput `pulumi:"preserveInterchange"`
	SuspendInterchangeOnError               pulumi.BoolInput `pulumi:"suspendInterchangeOnError"`
	UseDotAsDecimalSeparator                pulumi.BoolInput `pulumi:"useDotAsDecimalSeparator"`
}

func (X12ProcessingSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ProcessingSettings)(nil)).Elem()
}

func (i X12ProcessingSettingsArgs) ToX12ProcessingSettingsOutput() X12ProcessingSettingsOutput {
	return i.ToX12ProcessingSettingsOutputWithContext(context.Background())
}

func (i X12ProcessingSettingsArgs) ToX12ProcessingSettingsOutputWithContext(ctx context.Context) X12ProcessingSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ProcessingSettingsOutput)
}

func (i X12ProcessingSettingsArgs) ToX12ProcessingSettingsPtrOutput() X12ProcessingSettingsPtrOutput {
	return i.ToX12ProcessingSettingsPtrOutputWithContext(context.Background())
}

func (i X12ProcessingSettingsArgs) ToX12ProcessingSettingsPtrOutputWithContext(ctx context.Context) X12ProcessingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ProcessingSettingsOutput).ToX12ProcessingSettingsPtrOutputWithContext(ctx)
}









type X12ProcessingSettingsPtrInput interface {
	pulumi.Input

	ToX12ProcessingSettingsPtrOutput() X12ProcessingSettingsPtrOutput
	ToX12ProcessingSettingsPtrOutputWithContext(context.Context) X12ProcessingSettingsPtrOutput
}

type x12processingSettingsPtrType X12ProcessingSettingsArgs

func X12ProcessingSettingsPtr(v *X12ProcessingSettingsArgs) X12ProcessingSettingsPtrInput {
	return (*x12processingSettingsPtrType)(v)
}

func (*x12processingSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12ProcessingSettings)(nil)).Elem()
}

func (i *x12processingSettingsPtrType) ToX12ProcessingSettingsPtrOutput() X12ProcessingSettingsPtrOutput {
	return i.ToX12ProcessingSettingsPtrOutputWithContext(context.Background())
}

func (i *x12processingSettingsPtrType) ToX12ProcessingSettingsPtrOutputWithContext(ctx context.Context) X12ProcessingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ProcessingSettingsPtrOutput)
}

type X12ProcessingSettingsOutput struct{ *pulumi.OutputState }

func (X12ProcessingSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ProcessingSettings)(nil)).Elem()
}

func (o X12ProcessingSettingsOutput) ToX12ProcessingSettingsOutput() X12ProcessingSettingsOutput {
	return o
}

func (o X12ProcessingSettingsOutput) ToX12ProcessingSettingsOutputWithContext(ctx context.Context) X12ProcessingSettingsOutput {
	return o
}

func (o X12ProcessingSettingsOutput) ToX12ProcessingSettingsPtrOutput() X12ProcessingSettingsPtrOutput {
	return o.ToX12ProcessingSettingsPtrOutputWithContext(context.Background())
}

func (o X12ProcessingSettingsOutput) ToX12ProcessingSettingsPtrOutputWithContext(ctx context.Context) X12ProcessingSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12ProcessingSettings) *X12ProcessingSettings {
		return &v
	}).(X12ProcessingSettingsPtrOutput)
}

func (o X12ProcessingSettingsOutput) ConvertImpliedDecimal() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ProcessingSettings) bool { return v.ConvertImpliedDecimal }).(pulumi.BoolOutput)
}

func (o X12ProcessingSettingsOutput) CreateEmptyXmlTagsForTrailingSeparators() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ProcessingSettings) bool { return v.CreateEmptyXmlTagsForTrailingSeparators }).(pulumi.BoolOutput)
}

func (o X12ProcessingSettingsOutput) MaskSecurityInfo() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ProcessingSettings) bool { return v.MaskSecurityInfo }).(pulumi.BoolOutput)
}

func (o X12ProcessingSettingsOutput) PreserveInterchange() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ProcessingSettings) bool { return v.PreserveInterchange }).(pulumi.BoolOutput)
}

func (o X12ProcessingSettingsOutput) SuspendInterchangeOnError() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ProcessingSettings) bool { return v.SuspendInterchangeOnError }).(pulumi.BoolOutput)
}

func (o X12ProcessingSettingsOutput) UseDotAsDecimalSeparator() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ProcessingSettings) bool { return v.UseDotAsDecimalSeparator }).(pulumi.BoolOutput)
}

type X12ProcessingSettingsPtrOutput struct{ *pulumi.OutputState }

func (X12ProcessingSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12ProcessingSettings)(nil)).Elem()
}

func (o X12ProcessingSettingsPtrOutput) ToX12ProcessingSettingsPtrOutput() X12ProcessingSettingsPtrOutput {
	return o
}

func (o X12ProcessingSettingsPtrOutput) ToX12ProcessingSettingsPtrOutputWithContext(ctx context.Context) X12ProcessingSettingsPtrOutput {
	return o
}

func (o X12ProcessingSettingsPtrOutput) Elem() X12ProcessingSettingsOutput {
	return o.ApplyT(func(v *X12ProcessingSettings) X12ProcessingSettings {
		if v != nil {
			return *v
		}
		var ret X12ProcessingSettings
		return ret
	}).(X12ProcessingSettingsOutput)
}

func (o X12ProcessingSettingsPtrOutput) ConvertImpliedDecimal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.ConvertImpliedDecimal
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsPtrOutput) CreateEmptyXmlTagsForTrailingSeparators() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.CreateEmptyXmlTagsForTrailingSeparators
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsPtrOutput) MaskSecurityInfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.MaskSecurityInfo
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsPtrOutput) PreserveInterchange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.PreserveInterchange
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsPtrOutput) SuspendInterchangeOnError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.SuspendInterchangeOnError
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsPtrOutput) UseDotAsDecimalSeparator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.UseDotAsDecimalSeparator
	}).(pulumi.BoolPtrOutput)
}

type X12ProcessingSettingsResponse struct {
	ConvertImpliedDecimal                   bool `pulumi:"convertImpliedDecimal"`
	CreateEmptyXmlTagsForTrailingSeparators bool `pulumi:"createEmptyXmlTagsForTrailingSeparators"`
	MaskSecurityInfo                        bool `pulumi:"maskSecurityInfo"`
	PreserveInterchange                     bool `pulumi:"preserveInterchange"`
	SuspendInterchangeOnError               bool `pulumi:"suspendInterchangeOnError"`
	UseDotAsDecimalSeparator                bool `pulumi:"useDotAsDecimalSeparator"`
}

type X12ProcessingSettingsResponseOutput struct{ *pulumi.OutputState }

func (X12ProcessingSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ProcessingSettingsResponse)(nil)).Elem()
}

func (o X12ProcessingSettingsResponseOutput) ToX12ProcessingSettingsResponseOutput() X12ProcessingSettingsResponseOutput {
	return o
}

func (o X12ProcessingSettingsResponseOutput) ToX12ProcessingSettingsResponseOutputWithContext(ctx context.Context) X12ProcessingSettingsResponseOutput {
	return o
}

func (o X12ProcessingSettingsResponseOutput) ConvertImpliedDecimal() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ProcessingSettingsResponse) bool { return v.ConvertImpliedDecimal }).(pulumi.BoolOutput)
}

func (o X12ProcessingSettingsResponseOutput) CreateEmptyXmlTagsForTrailingSeparators() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ProcessingSettingsResponse) bool { return v.CreateEmptyXmlTagsForTrailingSeparators }).(pulumi.BoolOutput)
}

func (o X12ProcessingSettingsResponseOutput) MaskSecurityInfo() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ProcessingSettingsResponse) bool { return v.MaskSecurityInfo }).(pulumi.BoolOutput)
}

func (o X12ProcessingSettingsResponseOutput) PreserveInterchange() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ProcessingSettingsResponse) bool { return v.PreserveInterchange }).(pulumi.BoolOutput)
}

func (o X12ProcessingSettingsResponseOutput) SuspendInterchangeOnError() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ProcessingSettingsResponse) bool { return v.SuspendInterchangeOnError }).(pulumi.BoolOutput)
}

func (o X12ProcessingSettingsResponseOutput) UseDotAsDecimalSeparator() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ProcessingSettingsResponse) bool { return v.UseDotAsDecimalSeparator }).(pulumi.BoolOutput)
}

type X12ProcessingSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (X12ProcessingSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12ProcessingSettingsResponse)(nil)).Elem()
}

func (o X12ProcessingSettingsResponsePtrOutput) ToX12ProcessingSettingsResponsePtrOutput() X12ProcessingSettingsResponsePtrOutput {
	return o
}

func (o X12ProcessingSettingsResponsePtrOutput) ToX12ProcessingSettingsResponsePtrOutputWithContext(ctx context.Context) X12ProcessingSettingsResponsePtrOutput {
	return o
}

func (o X12ProcessingSettingsResponsePtrOutput) Elem() X12ProcessingSettingsResponseOutput {
	return o.ApplyT(func(v *X12ProcessingSettingsResponse) X12ProcessingSettingsResponse {
		if v != nil {
			return *v
		}
		var ret X12ProcessingSettingsResponse
		return ret
	}).(X12ProcessingSettingsResponseOutput)
}

func (o X12ProcessingSettingsResponsePtrOutput) ConvertImpliedDecimal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.ConvertImpliedDecimal
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsResponsePtrOutput) CreateEmptyXmlTagsForTrailingSeparators() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CreateEmptyXmlTagsForTrailingSeparators
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsResponsePtrOutput) MaskSecurityInfo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.MaskSecurityInfo
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsResponsePtrOutput) PreserveInterchange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.PreserveInterchange
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsResponsePtrOutput) SuspendInterchangeOnError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.SuspendInterchangeOnError
	}).(pulumi.BoolPtrOutput)
}

func (o X12ProcessingSettingsResponsePtrOutput) UseDotAsDecimalSeparator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ProcessingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.UseDotAsDecimalSeparator
	}).(pulumi.BoolPtrOutput)
}

type X12ProtocolSettings struct {
	AcknowledgementSettings X12AcknowledgementSettings `pulumi:"acknowledgementSettings"`
	EnvelopeOverrides       []X12EnvelopeOverride      `pulumi:"envelopeOverrides"`
	EnvelopeSettings        X12EnvelopeSettings        `pulumi:"envelopeSettings"`
	FramingSettings         X12FramingSettings         `pulumi:"framingSettings"`
	MessageFilter           X12MessageFilter           `pulumi:"messageFilter"`
	MessageFilterList       []X12MessageIdentifier     `pulumi:"messageFilterList"`
	ProcessingSettings      X12ProcessingSettings      `pulumi:"processingSettings"`
	SchemaReferences        []X12SchemaReference       `pulumi:"schemaReferences"`
	SecuritySettings        X12SecuritySettings        `pulumi:"securitySettings"`
	ValidationOverrides     []X12ValidationOverride    `pulumi:"validationOverrides"`
	ValidationSettings      X12ValidationSettings      `pulumi:"validationSettings"`
	X12DelimiterOverrides   []X12DelimiterOverrides    `pulumi:"x12DelimiterOverrides"`
}





type X12ProtocolSettingsInput interface {
	pulumi.Input

	ToX12ProtocolSettingsOutput() X12ProtocolSettingsOutput
	ToX12ProtocolSettingsOutputWithContext(context.Context) X12ProtocolSettingsOutput
}

type X12ProtocolSettingsArgs struct {
	AcknowledgementSettings X12AcknowledgementSettingsInput `pulumi:"acknowledgementSettings"`
	EnvelopeOverrides       X12EnvelopeOverrideArrayInput   `pulumi:"envelopeOverrides"`
	EnvelopeSettings        X12EnvelopeSettingsInput        `pulumi:"envelopeSettings"`
	FramingSettings         X12FramingSettingsInput         `pulumi:"framingSettings"`
	MessageFilter           X12MessageFilterInput           `pulumi:"messageFilter"`
	MessageFilterList       X12MessageIdentifierArrayInput  `pulumi:"messageFilterList"`
	ProcessingSettings      X12ProcessingSettingsInput      `pulumi:"processingSettings"`
	SchemaReferences        X12SchemaReferenceArrayInput    `pulumi:"schemaReferences"`
	SecuritySettings        X12SecuritySettingsInput        `pulumi:"securitySettings"`
	ValidationOverrides     X12ValidationOverrideArrayInput `pulumi:"validationOverrides"`
	ValidationSettings      X12ValidationSettingsInput      `pulumi:"validationSettings"`
	X12DelimiterOverrides   X12DelimiterOverridesArrayInput `pulumi:"x12DelimiterOverrides"`
}

func (X12ProtocolSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ProtocolSettings)(nil)).Elem()
}

func (i X12ProtocolSettingsArgs) ToX12ProtocolSettingsOutput() X12ProtocolSettingsOutput {
	return i.ToX12ProtocolSettingsOutputWithContext(context.Background())
}

func (i X12ProtocolSettingsArgs) ToX12ProtocolSettingsOutputWithContext(ctx context.Context) X12ProtocolSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ProtocolSettingsOutput)
}

func (i X12ProtocolSettingsArgs) ToX12ProtocolSettingsPtrOutput() X12ProtocolSettingsPtrOutput {
	return i.ToX12ProtocolSettingsPtrOutputWithContext(context.Background())
}

func (i X12ProtocolSettingsArgs) ToX12ProtocolSettingsPtrOutputWithContext(ctx context.Context) X12ProtocolSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ProtocolSettingsOutput).ToX12ProtocolSettingsPtrOutputWithContext(ctx)
}









type X12ProtocolSettingsPtrInput interface {
	pulumi.Input

	ToX12ProtocolSettingsPtrOutput() X12ProtocolSettingsPtrOutput
	ToX12ProtocolSettingsPtrOutputWithContext(context.Context) X12ProtocolSettingsPtrOutput
}

type x12protocolSettingsPtrType X12ProtocolSettingsArgs

func X12ProtocolSettingsPtr(v *X12ProtocolSettingsArgs) X12ProtocolSettingsPtrInput {
	return (*x12protocolSettingsPtrType)(v)
}

func (*x12protocolSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12ProtocolSettings)(nil)).Elem()
}

func (i *x12protocolSettingsPtrType) ToX12ProtocolSettingsPtrOutput() X12ProtocolSettingsPtrOutput {
	return i.ToX12ProtocolSettingsPtrOutputWithContext(context.Background())
}

func (i *x12protocolSettingsPtrType) ToX12ProtocolSettingsPtrOutputWithContext(ctx context.Context) X12ProtocolSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ProtocolSettingsPtrOutput)
}

type X12ProtocolSettingsOutput struct{ *pulumi.OutputState }

func (X12ProtocolSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ProtocolSettings)(nil)).Elem()
}

func (o X12ProtocolSettingsOutput) ToX12ProtocolSettingsOutput() X12ProtocolSettingsOutput {
	return o
}

func (o X12ProtocolSettingsOutput) ToX12ProtocolSettingsOutputWithContext(ctx context.Context) X12ProtocolSettingsOutput {
	return o
}

func (o X12ProtocolSettingsOutput) ToX12ProtocolSettingsPtrOutput() X12ProtocolSettingsPtrOutput {
	return o.ToX12ProtocolSettingsPtrOutputWithContext(context.Background())
}

func (o X12ProtocolSettingsOutput) ToX12ProtocolSettingsPtrOutputWithContext(ctx context.Context) X12ProtocolSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12ProtocolSettings) *X12ProtocolSettings {
		return &v
	}).(X12ProtocolSettingsPtrOutput)
}

func (o X12ProtocolSettingsOutput) AcknowledgementSettings() X12AcknowledgementSettingsOutput {
	return o.ApplyT(func(v X12ProtocolSettings) X12AcknowledgementSettings { return v.AcknowledgementSettings }).(X12AcknowledgementSettingsOutput)
}

func (o X12ProtocolSettingsOutput) EnvelopeOverrides() X12EnvelopeOverrideArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettings) []X12EnvelopeOverride { return v.EnvelopeOverrides }).(X12EnvelopeOverrideArrayOutput)
}

func (o X12ProtocolSettingsOutput) EnvelopeSettings() X12EnvelopeSettingsOutput {
	return o.ApplyT(func(v X12ProtocolSettings) X12EnvelopeSettings { return v.EnvelopeSettings }).(X12EnvelopeSettingsOutput)
}

func (o X12ProtocolSettingsOutput) FramingSettings() X12FramingSettingsOutput {
	return o.ApplyT(func(v X12ProtocolSettings) X12FramingSettings { return v.FramingSettings }).(X12FramingSettingsOutput)
}

func (o X12ProtocolSettingsOutput) MessageFilter() X12MessageFilterOutput {
	return o.ApplyT(func(v X12ProtocolSettings) X12MessageFilter { return v.MessageFilter }).(X12MessageFilterOutput)
}

func (o X12ProtocolSettingsOutput) MessageFilterList() X12MessageIdentifierArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettings) []X12MessageIdentifier { return v.MessageFilterList }).(X12MessageIdentifierArrayOutput)
}

func (o X12ProtocolSettingsOutput) ProcessingSettings() X12ProcessingSettingsOutput {
	return o.ApplyT(func(v X12ProtocolSettings) X12ProcessingSettings { return v.ProcessingSettings }).(X12ProcessingSettingsOutput)
}

func (o X12ProtocolSettingsOutput) SchemaReferences() X12SchemaReferenceArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettings) []X12SchemaReference { return v.SchemaReferences }).(X12SchemaReferenceArrayOutput)
}

func (o X12ProtocolSettingsOutput) SecuritySettings() X12SecuritySettingsOutput {
	return o.ApplyT(func(v X12ProtocolSettings) X12SecuritySettings { return v.SecuritySettings }).(X12SecuritySettingsOutput)
}

func (o X12ProtocolSettingsOutput) ValidationOverrides() X12ValidationOverrideArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettings) []X12ValidationOverride { return v.ValidationOverrides }).(X12ValidationOverrideArrayOutput)
}

func (o X12ProtocolSettingsOutput) ValidationSettings() X12ValidationSettingsOutput {
	return o.ApplyT(func(v X12ProtocolSettings) X12ValidationSettings { return v.ValidationSettings }).(X12ValidationSettingsOutput)
}

func (o X12ProtocolSettingsOutput) X12DelimiterOverrides() X12DelimiterOverridesArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettings) []X12DelimiterOverrides { return v.X12DelimiterOverrides }).(X12DelimiterOverridesArrayOutput)
}

type X12ProtocolSettingsPtrOutput struct{ *pulumi.OutputState }

func (X12ProtocolSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12ProtocolSettings)(nil)).Elem()
}

func (o X12ProtocolSettingsPtrOutput) ToX12ProtocolSettingsPtrOutput() X12ProtocolSettingsPtrOutput {
	return o
}

func (o X12ProtocolSettingsPtrOutput) ToX12ProtocolSettingsPtrOutputWithContext(ctx context.Context) X12ProtocolSettingsPtrOutput {
	return o
}

func (o X12ProtocolSettingsPtrOutput) Elem() X12ProtocolSettingsOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) X12ProtocolSettings {
		if v != nil {
			return *v
		}
		var ret X12ProtocolSettings
		return ret
	}).(X12ProtocolSettingsOutput)
}

func (o X12ProtocolSettingsPtrOutput) AcknowledgementSettings() X12AcknowledgementSettingsPtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) *X12AcknowledgementSettings {
		if v == nil {
			return nil
		}
		return &v.AcknowledgementSettings
	}).(X12AcknowledgementSettingsPtrOutput)
}

func (o X12ProtocolSettingsPtrOutput) EnvelopeOverrides() X12EnvelopeOverrideArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) []X12EnvelopeOverride {
		if v == nil {
			return nil
		}
		return v.EnvelopeOverrides
	}).(X12EnvelopeOverrideArrayOutput)
}

func (o X12ProtocolSettingsPtrOutput) EnvelopeSettings() X12EnvelopeSettingsPtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) *X12EnvelopeSettings {
		if v == nil {
			return nil
		}
		return &v.EnvelopeSettings
	}).(X12EnvelopeSettingsPtrOutput)
}

func (o X12ProtocolSettingsPtrOutput) FramingSettings() X12FramingSettingsPtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) *X12FramingSettings {
		if v == nil {
			return nil
		}
		return &v.FramingSettings
	}).(X12FramingSettingsPtrOutput)
}

func (o X12ProtocolSettingsPtrOutput) MessageFilter() X12MessageFilterPtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) *X12MessageFilter {
		if v == nil {
			return nil
		}
		return &v.MessageFilter
	}).(X12MessageFilterPtrOutput)
}

func (o X12ProtocolSettingsPtrOutput) MessageFilterList() X12MessageIdentifierArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) []X12MessageIdentifier {
		if v == nil {
			return nil
		}
		return v.MessageFilterList
	}).(X12MessageIdentifierArrayOutput)
}

func (o X12ProtocolSettingsPtrOutput) ProcessingSettings() X12ProcessingSettingsPtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) *X12ProcessingSettings {
		if v == nil {
			return nil
		}
		return &v.ProcessingSettings
	}).(X12ProcessingSettingsPtrOutput)
}

func (o X12ProtocolSettingsPtrOutput) SchemaReferences() X12SchemaReferenceArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) []X12SchemaReference {
		if v == nil {
			return nil
		}
		return v.SchemaReferences
	}).(X12SchemaReferenceArrayOutput)
}

func (o X12ProtocolSettingsPtrOutput) SecuritySettings() X12SecuritySettingsPtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) *X12SecuritySettings {
		if v == nil {
			return nil
		}
		return &v.SecuritySettings
	}).(X12SecuritySettingsPtrOutput)
}

func (o X12ProtocolSettingsPtrOutput) ValidationOverrides() X12ValidationOverrideArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) []X12ValidationOverride {
		if v == nil {
			return nil
		}
		return v.ValidationOverrides
	}).(X12ValidationOverrideArrayOutput)
}

func (o X12ProtocolSettingsPtrOutput) ValidationSettings() X12ValidationSettingsPtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) *X12ValidationSettings {
		if v == nil {
			return nil
		}
		return &v.ValidationSettings
	}).(X12ValidationSettingsPtrOutput)
}

func (o X12ProtocolSettingsPtrOutput) X12DelimiterOverrides() X12DelimiterOverridesArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettings) []X12DelimiterOverrides {
		if v == nil {
			return nil
		}
		return v.X12DelimiterOverrides
	}).(X12DelimiterOverridesArrayOutput)
}

type X12ProtocolSettingsResponse struct {
	AcknowledgementSettings X12AcknowledgementSettingsResponse `pulumi:"acknowledgementSettings"`
	EnvelopeOverrides       []X12EnvelopeOverrideResponse      `pulumi:"envelopeOverrides"`
	EnvelopeSettings        X12EnvelopeSettingsResponse        `pulumi:"envelopeSettings"`
	FramingSettings         X12FramingSettingsResponse         `pulumi:"framingSettings"`
	MessageFilter           X12MessageFilterResponse           `pulumi:"messageFilter"`
	MessageFilterList       []X12MessageIdentifierResponse     `pulumi:"messageFilterList"`
	ProcessingSettings      X12ProcessingSettingsResponse      `pulumi:"processingSettings"`
	SchemaReferences        []X12SchemaReferenceResponse       `pulumi:"schemaReferences"`
	SecuritySettings        X12SecuritySettingsResponse        `pulumi:"securitySettings"`
	ValidationOverrides     []X12ValidationOverrideResponse    `pulumi:"validationOverrides"`
	ValidationSettings      X12ValidationSettingsResponse      `pulumi:"validationSettings"`
	X12DelimiterOverrides   []X12DelimiterOverridesResponse    `pulumi:"x12DelimiterOverrides"`
}

type X12ProtocolSettingsResponseOutput struct{ *pulumi.OutputState }

func (X12ProtocolSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ProtocolSettingsResponse)(nil)).Elem()
}

func (o X12ProtocolSettingsResponseOutput) ToX12ProtocolSettingsResponseOutput() X12ProtocolSettingsResponseOutput {
	return o
}

func (o X12ProtocolSettingsResponseOutput) ToX12ProtocolSettingsResponseOutputWithContext(ctx context.Context) X12ProtocolSettingsResponseOutput {
	return o
}

func (o X12ProtocolSettingsResponseOutput) AcknowledgementSettings() X12AcknowledgementSettingsResponseOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) X12AcknowledgementSettingsResponse {
		return v.AcknowledgementSettings
	}).(X12AcknowledgementSettingsResponseOutput)
}

func (o X12ProtocolSettingsResponseOutput) EnvelopeOverrides() X12EnvelopeOverrideResponseArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) []X12EnvelopeOverrideResponse { return v.EnvelopeOverrides }).(X12EnvelopeOverrideResponseArrayOutput)
}

func (o X12ProtocolSettingsResponseOutput) EnvelopeSettings() X12EnvelopeSettingsResponseOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) X12EnvelopeSettingsResponse { return v.EnvelopeSettings }).(X12EnvelopeSettingsResponseOutput)
}

func (o X12ProtocolSettingsResponseOutput) FramingSettings() X12FramingSettingsResponseOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) X12FramingSettingsResponse { return v.FramingSettings }).(X12FramingSettingsResponseOutput)
}

func (o X12ProtocolSettingsResponseOutput) MessageFilter() X12MessageFilterResponseOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) X12MessageFilterResponse { return v.MessageFilter }).(X12MessageFilterResponseOutput)
}

func (o X12ProtocolSettingsResponseOutput) MessageFilterList() X12MessageIdentifierResponseArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) []X12MessageIdentifierResponse { return v.MessageFilterList }).(X12MessageIdentifierResponseArrayOutput)
}

func (o X12ProtocolSettingsResponseOutput) ProcessingSettings() X12ProcessingSettingsResponseOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) X12ProcessingSettingsResponse { return v.ProcessingSettings }).(X12ProcessingSettingsResponseOutput)
}

func (o X12ProtocolSettingsResponseOutput) SchemaReferences() X12SchemaReferenceResponseArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) []X12SchemaReferenceResponse { return v.SchemaReferences }).(X12SchemaReferenceResponseArrayOutput)
}

func (o X12ProtocolSettingsResponseOutput) SecuritySettings() X12SecuritySettingsResponseOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) X12SecuritySettingsResponse { return v.SecuritySettings }).(X12SecuritySettingsResponseOutput)
}

func (o X12ProtocolSettingsResponseOutput) ValidationOverrides() X12ValidationOverrideResponseArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) []X12ValidationOverrideResponse { return v.ValidationOverrides }).(X12ValidationOverrideResponseArrayOutput)
}

func (o X12ProtocolSettingsResponseOutput) ValidationSettings() X12ValidationSettingsResponseOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) X12ValidationSettingsResponse { return v.ValidationSettings }).(X12ValidationSettingsResponseOutput)
}

func (o X12ProtocolSettingsResponseOutput) X12DelimiterOverrides() X12DelimiterOverridesResponseArrayOutput {
	return o.ApplyT(func(v X12ProtocolSettingsResponse) []X12DelimiterOverridesResponse { return v.X12DelimiterOverrides }).(X12DelimiterOverridesResponseArrayOutput)
}

type X12ProtocolSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (X12ProtocolSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12ProtocolSettingsResponse)(nil)).Elem()
}

func (o X12ProtocolSettingsResponsePtrOutput) ToX12ProtocolSettingsResponsePtrOutput() X12ProtocolSettingsResponsePtrOutput {
	return o
}

func (o X12ProtocolSettingsResponsePtrOutput) ToX12ProtocolSettingsResponsePtrOutputWithContext(ctx context.Context) X12ProtocolSettingsResponsePtrOutput {
	return o
}

func (o X12ProtocolSettingsResponsePtrOutput) Elem() X12ProtocolSettingsResponseOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) X12ProtocolSettingsResponse {
		if v != nil {
			return *v
		}
		var ret X12ProtocolSettingsResponse
		return ret
	}).(X12ProtocolSettingsResponseOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) AcknowledgementSettings() X12AcknowledgementSettingsResponsePtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) *X12AcknowledgementSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.AcknowledgementSettings
	}).(X12AcknowledgementSettingsResponsePtrOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) EnvelopeOverrides() X12EnvelopeOverrideResponseArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) []X12EnvelopeOverrideResponse {
		if v == nil {
			return nil
		}
		return v.EnvelopeOverrides
	}).(X12EnvelopeOverrideResponseArrayOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) EnvelopeSettings() X12EnvelopeSettingsResponsePtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) *X12EnvelopeSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.EnvelopeSettings
	}).(X12EnvelopeSettingsResponsePtrOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) FramingSettings() X12FramingSettingsResponsePtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) *X12FramingSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.FramingSettings
	}).(X12FramingSettingsResponsePtrOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) MessageFilter() X12MessageFilterResponsePtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) *X12MessageFilterResponse {
		if v == nil {
			return nil
		}
		return &v.MessageFilter
	}).(X12MessageFilterResponsePtrOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) MessageFilterList() X12MessageIdentifierResponseArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) []X12MessageIdentifierResponse {
		if v == nil {
			return nil
		}
		return v.MessageFilterList
	}).(X12MessageIdentifierResponseArrayOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) ProcessingSettings() X12ProcessingSettingsResponsePtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) *X12ProcessingSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.ProcessingSettings
	}).(X12ProcessingSettingsResponsePtrOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) SchemaReferences() X12SchemaReferenceResponseArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) []X12SchemaReferenceResponse {
		if v == nil {
			return nil
		}
		return v.SchemaReferences
	}).(X12SchemaReferenceResponseArrayOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) SecuritySettings() X12SecuritySettingsResponsePtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) *X12SecuritySettingsResponse {
		if v == nil {
			return nil
		}
		return &v.SecuritySettings
	}).(X12SecuritySettingsResponsePtrOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) ValidationOverrides() X12ValidationOverrideResponseArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) []X12ValidationOverrideResponse {
		if v == nil {
			return nil
		}
		return v.ValidationOverrides
	}).(X12ValidationOverrideResponseArrayOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) ValidationSettings() X12ValidationSettingsResponsePtrOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) *X12ValidationSettingsResponse {
		if v == nil {
			return nil
		}
		return &v.ValidationSettings
	}).(X12ValidationSettingsResponsePtrOutput)
}

func (o X12ProtocolSettingsResponsePtrOutput) X12DelimiterOverrides() X12DelimiterOverridesResponseArrayOutput {
	return o.ApplyT(func(v *X12ProtocolSettingsResponse) []X12DelimiterOverridesResponse {
		if v == nil {
			return nil
		}
		return v.X12DelimiterOverrides
	}).(X12DelimiterOverridesResponseArrayOutput)
}

type X12SchemaReference struct {
	MessageId           string  `pulumi:"messageId"`
	SchemaName          string  `pulumi:"schemaName"`
	SchemaVersion       string  `pulumi:"schemaVersion"`
	SenderApplicationId *string `pulumi:"senderApplicationId"`
}





type X12SchemaReferenceInput interface {
	pulumi.Input

	ToX12SchemaReferenceOutput() X12SchemaReferenceOutput
	ToX12SchemaReferenceOutputWithContext(context.Context) X12SchemaReferenceOutput
}

type X12SchemaReferenceArgs struct {
	MessageId           pulumi.StringInput    `pulumi:"messageId"`
	SchemaName          pulumi.StringInput    `pulumi:"schemaName"`
	SchemaVersion       pulumi.StringInput    `pulumi:"schemaVersion"`
	SenderApplicationId pulumi.StringPtrInput `pulumi:"senderApplicationId"`
}

func (X12SchemaReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12SchemaReference)(nil)).Elem()
}

func (i X12SchemaReferenceArgs) ToX12SchemaReferenceOutput() X12SchemaReferenceOutput {
	return i.ToX12SchemaReferenceOutputWithContext(context.Background())
}

func (i X12SchemaReferenceArgs) ToX12SchemaReferenceOutputWithContext(ctx context.Context) X12SchemaReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12SchemaReferenceOutput)
}





type X12SchemaReferenceArrayInput interface {
	pulumi.Input

	ToX12SchemaReferenceArrayOutput() X12SchemaReferenceArrayOutput
	ToX12SchemaReferenceArrayOutputWithContext(context.Context) X12SchemaReferenceArrayOutput
}

type X12SchemaReferenceArray []X12SchemaReferenceInput

func (X12SchemaReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12SchemaReference)(nil)).Elem()
}

func (i X12SchemaReferenceArray) ToX12SchemaReferenceArrayOutput() X12SchemaReferenceArrayOutput {
	return i.ToX12SchemaReferenceArrayOutputWithContext(context.Background())
}

func (i X12SchemaReferenceArray) ToX12SchemaReferenceArrayOutputWithContext(ctx context.Context) X12SchemaReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12SchemaReferenceArrayOutput)
}

type X12SchemaReferenceOutput struct{ *pulumi.OutputState }

func (X12SchemaReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12SchemaReference)(nil)).Elem()
}

func (o X12SchemaReferenceOutput) ToX12SchemaReferenceOutput() X12SchemaReferenceOutput {
	return o
}

func (o X12SchemaReferenceOutput) ToX12SchemaReferenceOutputWithContext(ctx context.Context) X12SchemaReferenceOutput {
	return o
}

func (o X12SchemaReferenceOutput) MessageId() pulumi.StringOutput {
	return o.ApplyT(func(v X12SchemaReference) string { return v.MessageId }).(pulumi.StringOutput)
}

func (o X12SchemaReferenceOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v X12SchemaReference) string { return v.SchemaName }).(pulumi.StringOutput)
}

func (o X12SchemaReferenceOutput) SchemaVersion() pulumi.StringOutput {
	return o.ApplyT(func(v X12SchemaReference) string { return v.SchemaVersion }).(pulumi.StringOutput)
}

func (o X12SchemaReferenceOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SchemaReference) *string { return v.SenderApplicationId }).(pulumi.StringPtrOutput)
}

type X12SchemaReferenceArrayOutput struct{ *pulumi.OutputState }

func (X12SchemaReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12SchemaReference)(nil)).Elem()
}

func (o X12SchemaReferenceArrayOutput) ToX12SchemaReferenceArrayOutput() X12SchemaReferenceArrayOutput {
	return o
}

func (o X12SchemaReferenceArrayOutput) ToX12SchemaReferenceArrayOutputWithContext(ctx context.Context) X12SchemaReferenceArrayOutput {
	return o
}

func (o X12SchemaReferenceArrayOutput) Index(i pulumi.IntInput) X12SchemaReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12SchemaReference {
		return vs[0].([]X12SchemaReference)[vs[1].(int)]
	}).(X12SchemaReferenceOutput)
}

type X12SchemaReferenceResponse struct {
	MessageId           string  `pulumi:"messageId"`
	SchemaName          string  `pulumi:"schemaName"`
	SchemaVersion       string  `pulumi:"schemaVersion"`
	SenderApplicationId *string `pulumi:"senderApplicationId"`
}

type X12SchemaReferenceResponseOutput struct{ *pulumi.OutputState }

func (X12SchemaReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12SchemaReferenceResponse)(nil)).Elem()
}

func (o X12SchemaReferenceResponseOutput) ToX12SchemaReferenceResponseOutput() X12SchemaReferenceResponseOutput {
	return o
}

func (o X12SchemaReferenceResponseOutput) ToX12SchemaReferenceResponseOutputWithContext(ctx context.Context) X12SchemaReferenceResponseOutput {
	return o
}

func (o X12SchemaReferenceResponseOutput) MessageId() pulumi.StringOutput {
	return o.ApplyT(func(v X12SchemaReferenceResponse) string { return v.MessageId }).(pulumi.StringOutput)
}

func (o X12SchemaReferenceResponseOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v X12SchemaReferenceResponse) string { return v.SchemaName }).(pulumi.StringOutput)
}

func (o X12SchemaReferenceResponseOutput) SchemaVersion() pulumi.StringOutput {
	return o.ApplyT(func(v X12SchemaReferenceResponse) string { return v.SchemaVersion }).(pulumi.StringOutput)
}

func (o X12SchemaReferenceResponseOutput) SenderApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SchemaReferenceResponse) *string { return v.SenderApplicationId }).(pulumi.StringPtrOutput)
}

type X12SchemaReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (X12SchemaReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12SchemaReferenceResponse)(nil)).Elem()
}

func (o X12SchemaReferenceResponseArrayOutput) ToX12SchemaReferenceResponseArrayOutput() X12SchemaReferenceResponseArrayOutput {
	return o
}

func (o X12SchemaReferenceResponseArrayOutput) ToX12SchemaReferenceResponseArrayOutputWithContext(ctx context.Context) X12SchemaReferenceResponseArrayOutput {
	return o
}

func (o X12SchemaReferenceResponseArrayOutput) Index(i pulumi.IntInput) X12SchemaReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12SchemaReferenceResponse {
		return vs[0].([]X12SchemaReferenceResponse)[vs[1].(int)]
	}).(X12SchemaReferenceResponseOutput)
}

type X12SecuritySettings struct {
	AuthorizationQualifier string  `pulumi:"authorizationQualifier"`
	AuthorizationValue     *string `pulumi:"authorizationValue"`
	PasswordValue          *string `pulumi:"passwordValue"`
	SecurityQualifier      string  `pulumi:"securityQualifier"`
}





type X12SecuritySettingsInput interface {
	pulumi.Input

	ToX12SecuritySettingsOutput() X12SecuritySettingsOutput
	ToX12SecuritySettingsOutputWithContext(context.Context) X12SecuritySettingsOutput
}

type X12SecuritySettingsArgs struct {
	AuthorizationQualifier pulumi.StringInput    `pulumi:"authorizationQualifier"`
	AuthorizationValue     pulumi.StringPtrInput `pulumi:"authorizationValue"`
	PasswordValue          pulumi.StringPtrInput `pulumi:"passwordValue"`
	SecurityQualifier      pulumi.StringInput    `pulumi:"securityQualifier"`
}

func (X12SecuritySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12SecuritySettings)(nil)).Elem()
}

func (i X12SecuritySettingsArgs) ToX12SecuritySettingsOutput() X12SecuritySettingsOutput {
	return i.ToX12SecuritySettingsOutputWithContext(context.Background())
}

func (i X12SecuritySettingsArgs) ToX12SecuritySettingsOutputWithContext(ctx context.Context) X12SecuritySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12SecuritySettingsOutput)
}

func (i X12SecuritySettingsArgs) ToX12SecuritySettingsPtrOutput() X12SecuritySettingsPtrOutput {
	return i.ToX12SecuritySettingsPtrOutputWithContext(context.Background())
}

func (i X12SecuritySettingsArgs) ToX12SecuritySettingsPtrOutputWithContext(ctx context.Context) X12SecuritySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12SecuritySettingsOutput).ToX12SecuritySettingsPtrOutputWithContext(ctx)
}









type X12SecuritySettingsPtrInput interface {
	pulumi.Input

	ToX12SecuritySettingsPtrOutput() X12SecuritySettingsPtrOutput
	ToX12SecuritySettingsPtrOutputWithContext(context.Context) X12SecuritySettingsPtrOutput
}

type x12securitySettingsPtrType X12SecuritySettingsArgs

func X12SecuritySettingsPtr(v *X12SecuritySettingsArgs) X12SecuritySettingsPtrInput {
	return (*x12securitySettingsPtrType)(v)
}

func (*x12securitySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12SecuritySettings)(nil)).Elem()
}

func (i *x12securitySettingsPtrType) ToX12SecuritySettingsPtrOutput() X12SecuritySettingsPtrOutput {
	return i.ToX12SecuritySettingsPtrOutputWithContext(context.Background())
}

func (i *x12securitySettingsPtrType) ToX12SecuritySettingsPtrOutputWithContext(ctx context.Context) X12SecuritySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12SecuritySettingsPtrOutput)
}

type X12SecuritySettingsOutput struct{ *pulumi.OutputState }

func (X12SecuritySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12SecuritySettings)(nil)).Elem()
}

func (o X12SecuritySettingsOutput) ToX12SecuritySettingsOutput() X12SecuritySettingsOutput {
	return o
}

func (o X12SecuritySettingsOutput) ToX12SecuritySettingsOutputWithContext(ctx context.Context) X12SecuritySettingsOutput {
	return o
}

func (o X12SecuritySettingsOutput) ToX12SecuritySettingsPtrOutput() X12SecuritySettingsPtrOutput {
	return o.ToX12SecuritySettingsPtrOutputWithContext(context.Background())
}

func (o X12SecuritySettingsOutput) ToX12SecuritySettingsPtrOutputWithContext(ctx context.Context) X12SecuritySettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12SecuritySettings) *X12SecuritySettings {
		return &v
	}).(X12SecuritySettingsPtrOutput)
}

func (o X12SecuritySettingsOutput) AuthorizationQualifier() pulumi.StringOutput {
	return o.ApplyT(func(v X12SecuritySettings) string { return v.AuthorizationQualifier }).(pulumi.StringOutput)
}

func (o X12SecuritySettingsOutput) AuthorizationValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SecuritySettings) *string { return v.AuthorizationValue }).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsOutput) PasswordValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SecuritySettings) *string { return v.PasswordValue }).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsOutput) SecurityQualifier() pulumi.StringOutput {
	return o.ApplyT(func(v X12SecuritySettings) string { return v.SecurityQualifier }).(pulumi.StringOutput)
}

type X12SecuritySettingsPtrOutput struct{ *pulumi.OutputState }

func (X12SecuritySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12SecuritySettings)(nil)).Elem()
}

func (o X12SecuritySettingsPtrOutput) ToX12SecuritySettingsPtrOutput() X12SecuritySettingsPtrOutput {
	return o
}

func (o X12SecuritySettingsPtrOutput) ToX12SecuritySettingsPtrOutputWithContext(ctx context.Context) X12SecuritySettingsPtrOutput {
	return o
}

func (o X12SecuritySettingsPtrOutput) Elem() X12SecuritySettingsOutput {
	return o.ApplyT(func(v *X12SecuritySettings) X12SecuritySettings {
		if v != nil {
			return *v
		}
		var ret X12SecuritySettings
		return ret
	}).(X12SecuritySettingsOutput)
}

func (o X12SecuritySettingsPtrOutput) AuthorizationQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12SecuritySettings) *string {
		if v == nil {
			return nil
		}
		return &v.AuthorizationQualifier
	}).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsPtrOutput) AuthorizationValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12SecuritySettings) *string {
		if v == nil {
			return nil
		}
		return v.AuthorizationValue
	}).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsPtrOutput) PasswordValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12SecuritySettings) *string {
		if v == nil {
			return nil
		}
		return v.PasswordValue
	}).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsPtrOutput) SecurityQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12SecuritySettings) *string {
		if v == nil {
			return nil
		}
		return &v.SecurityQualifier
	}).(pulumi.StringPtrOutput)
}

type X12SecuritySettingsResponse struct {
	AuthorizationQualifier string  `pulumi:"authorizationQualifier"`
	AuthorizationValue     *string `pulumi:"authorizationValue"`
	PasswordValue          *string `pulumi:"passwordValue"`
	SecurityQualifier      string  `pulumi:"securityQualifier"`
}

type X12SecuritySettingsResponseOutput struct{ *pulumi.OutputState }

func (X12SecuritySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12SecuritySettingsResponse)(nil)).Elem()
}

func (o X12SecuritySettingsResponseOutput) ToX12SecuritySettingsResponseOutput() X12SecuritySettingsResponseOutput {
	return o
}

func (o X12SecuritySettingsResponseOutput) ToX12SecuritySettingsResponseOutputWithContext(ctx context.Context) X12SecuritySettingsResponseOutput {
	return o
}

func (o X12SecuritySettingsResponseOutput) AuthorizationQualifier() pulumi.StringOutput {
	return o.ApplyT(func(v X12SecuritySettingsResponse) string { return v.AuthorizationQualifier }).(pulumi.StringOutput)
}

func (o X12SecuritySettingsResponseOutput) AuthorizationValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SecuritySettingsResponse) *string { return v.AuthorizationValue }).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsResponseOutput) PasswordValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v X12SecuritySettingsResponse) *string { return v.PasswordValue }).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsResponseOutput) SecurityQualifier() pulumi.StringOutput {
	return o.ApplyT(func(v X12SecuritySettingsResponse) string { return v.SecurityQualifier }).(pulumi.StringOutput)
}

type X12SecuritySettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (X12SecuritySettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12SecuritySettingsResponse)(nil)).Elem()
}

func (o X12SecuritySettingsResponsePtrOutput) ToX12SecuritySettingsResponsePtrOutput() X12SecuritySettingsResponsePtrOutput {
	return o
}

func (o X12SecuritySettingsResponsePtrOutput) ToX12SecuritySettingsResponsePtrOutputWithContext(ctx context.Context) X12SecuritySettingsResponsePtrOutput {
	return o
}

func (o X12SecuritySettingsResponsePtrOutput) Elem() X12SecuritySettingsResponseOutput {
	return o.ApplyT(func(v *X12SecuritySettingsResponse) X12SecuritySettingsResponse {
		if v != nil {
			return *v
		}
		var ret X12SecuritySettingsResponse
		return ret
	}).(X12SecuritySettingsResponseOutput)
}

func (o X12SecuritySettingsResponsePtrOutput) AuthorizationQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12SecuritySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AuthorizationQualifier
	}).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsResponsePtrOutput) AuthorizationValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12SecuritySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuthorizationValue
	}).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsResponsePtrOutput) PasswordValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12SecuritySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.PasswordValue
	}).(pulumi.StringPtrOutput)
}

func (o X12SecuritySettingsResponsePtrOutput) SecurityQualifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12SecuritySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecurityQualifier
	}).(pulumi.StringPtrOutput)
}

type X12ValidationOverride struct {
	AllowLeadingAndTrailingSpacesAndZeroes bool   `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	MessageId                              string `pulumi:"messageId"`
	TrailingSeparatorPolicy                string `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes  bool   `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateCharacterSet                   bool   `pulumi:"validateCharacterSet"`
	ValidateEDITypes                       bool   `pulumi:"validateEDITypes"`
	ValidateXSDTypes                       bool   `pulumi:"validateXSDTypes"`
}





type X12ValidationOverrideInput interface {
	pulumi.Input

	ToX12ValidationOverrideOutput() X12ValidationOverrideOutput
	ToX12ValidationOverrideOutputWithContext(context.Context) X12ValidationOverrideOutput
}

type X12ValidationOverrideArgs struct {
	AllowLeadingAndTrailingSpacesAndZeroes pulumi.BoolInput   `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	MessageId                              pulumi.StringInput `pulumi:"messageId"`
	TrailingSeparatorPolicy                pulumi.StringInput `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes  pulumi.BoolInput   `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateCharacterSet                   pulumi.BoolInput   `pulumi:"validateCharacterSet"`
	ValidateEDITypes                       pulumi.BoolInput   `pulumi:"validateEDITypes"`
	ValidateXSDTypes                       pulumi.BoolInput   `pulumi:"validateXSDTypes"`
}

func (X12ValidationOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ValidationOverride)(nil)).Elem()
}

func (i X12ValidationOverrideArgs) ToX12ValidationOverrideOutput() X12ValidationOverrideOutput {
	return i.ToX12ValidationOverrideOutputWithContext(context.Background())
}

func (i X12ValidationOverrideArgs) ToX12ValidationOverrideOutputWithContext(ctx context.Context) X12ValidationOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ValidationOverrideOutput)
}





type X12ValidationOverrideArrayInput interface {
	pulumi.Input

	ToX12ValidationOverrideArrayOutput() X12ValidationOverrideArrayOutput
	ToX12ValidationOverrideArrayOutputWithContext(context.Context) X12ValidationOverrideArrayOutput
}

type X12ValidationOverrideArray []X12ValidationOverrideInput

func (X12ValidationOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12ValidationOverride)(nil)).Elem()
}

func (i X12ValidationOverrideArray) ToX12ValidationOverrideArrayOutput() X12ValidationOverrideArrayOutput {
	return i.ToX12ValidationOverrideArrayOutputWithContext(context.Background())
}

func (i X12ValidationOverrideArray) ToX12ValidationOverrideArrayOutputWithContext(ctx context.Context) X12ValidationOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ValidationOverrideArrayOutput)
}

type X12ValidationOverrideOutput struct{ *pulumi.OutputState }

func (X12ValidationOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ValidationOverride)(nil)).Elem()
}

func (o X12ValidationOverrideOutput) ToX12ValidationOverrideOutput() X12ValidationOverrideOutput {
	return o
}

func (o X12ValidationOverrideOutput) ToX12ValidationOverrideOutputWithContext(ctx context.Context) X12ValidationOverrideOutput {
	return o
}

func (o X12ValidationOverrideOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationOverride) bool { return v.AllowLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolOutput)
}

func (o X12ValidationOverrideOutput) MessageId() pulumi.StringOutput {
	return o.ApplyT(func(v X12ValidationOverride) string { return v.MessageId }).(pulumi.StringOutput)
}

func (o X12ValidationOverrideOutput) TrailingSeparatorPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v X12ValidationOverride) string { return v.TrailingSeparatorPolicy }).(pulumi.StringOutput)
}

func (o X12ValidationOverrideOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationOverride) bool { return v.TrimLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolOutput)
}

func (o X12ValidationOverrideOutput) ValidateCharacterSet() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationOverride) bool { return v.ValidateCharacterSet }).(pulumi.BoolOutput)
}

func (o X12ValidationOverrideOutput) ValidateEDITypes() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationOverride) bool { return v.ValidateEDITypes }).(pulumi.BoolOutput)
}

func (o X12ValidationOverrideOutput) ValidateXSDTypes() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationOverride) bool { return v.ValidateXSDTypes }).(pulumi.BoolOutput)
}

type X12ValidationOverrideArrayOutput struct{ *pulumi.OutputState }

func (X12ValidationOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12ValidationOverride)(nil)).Elem()
}

func (o X12ValidationOverrideArrayOutput) ToX12ValidationOverrideArrayOutput() X12ValidationOverrideArrayOutput {
	return o
}

func (o X12ValidationOverrideArrayOutput) ToX12ValidationOverrideArrayOutputWithContext(ctx context.Context) X12ValidationOverrideArrayOutput {
	return o
}

func (o X12ValidationOverrideArrayOutput) Index(i pulumi.IntInput) X12ValidationOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12ValidationOverride {
		return vs[0].([]X12ValidationOverride)[vs[1].(int)]
	}).(X12ValidationOverrideOutput)
}

type X12ValidationOverrideResponse struct {
	AllowLeadingAndTrailingSpacesAndZeroes bool   `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	MessageId                              string `pulumi:"messageId"`
	TrailingSeparatorPolicy                string `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes  bool   `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateCharacterSet                   bool   `pulumi:"validateCharacterSet"`
	ValidateEDITypes                       bool   `pulumi:"validateEDITypes"`
	ValidateXSDTypes                       bool   `pulumi:"validateXSDTypes"`
}

type X12ValidationOverrideResponseOutput struct{ *pulumi.OutputState }

func (X12ValidationOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ValidationOverrideResponse)(nil)).Elem()
}

func (o X12ValidationOverrideResponseOutput) ToX12ValidationOverrideResponseOutput() X12ValidationOverrideResponseOutput {
	return o
}

func (o X12ValidationOverrideResponseOutput) ToX12ValidationOverrideResponseOutputWithContext(ctx context.Context) X12ValidationOverrideResponseOutput {
	return o
}

func (o X12ValidationOverrideResponseOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationOverrideResponse) bool { return v.AllowLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolOutput)
}

func (o X12ValidationOverrideResponseOutput) MessageId() pulumi.StringOutput {
	return o.ApplyT(func(v X12ValidationOverrideResponse) string { return v.MessageId }).(pulumi.StringOutput)
}

func (o X12ValidationOverrideResponseOutput) TrailingSeparatorPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v X12ValidationOverrideResponse) string { return v.TrailingSeparatorPolicy }).(pulumi.StringOutput)
}

func (o X12ValidationOverrideResponseOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationOverrideResponse) bool { return v.TrimLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolOutput)
}

func (o X12ValidationOverrideResponseOutput) ValidateCharacterSet() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationOverrideResponse) bool { return v.ValidateCharacterSet }).(pulumi.BoolOutput)
}

func (o X12ValidationOverrideResponseOutput) ValidateEDITypes() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationOverrideResponse) bool { return v.ValidateEDITypes }).(pulumi.BoolOutput)
}

func (o X12ValidationOverrideResponseOutput) ValidateXSDTypes() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationOverrideResponse) bool { return v.ValidateXSDTypes }).(pulumi.BoolOutput)
}

type X12ValidationOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (X12ValidationOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]X12ValidationOverrideResponse)(nil)).Elem()
}

func (o X12ValidationOverrideResponseArrayOutput) ToX12ValidationOverrideResponseArrayOutput() X12ValidationOverrideResponseArrayOutput {
	return o
}

func (o X12ValidationOverrideResponseArrayOutput) ToX12ValidationOverrideResponseArrayOutputWithContext(ctx context.Context) X12ValidationOverrideResponseArrayOutput {
	return o
}

func (o X12ValidationOverrideResponseArrayOutput) Index(i pulumi.IntInput) X12ValidationOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) X12ValidationOverrideResponse {
		return vs[0].([]X12ValidationOverrideResponse)[vs[1].(int)]
	}).(X12ValidationOverrideResponseOutput)
}

type X12ValidationSettings struct {
	AllowLeadingAndTrailingSpacesAndZeroes    bool   `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	CheckDuplicateGroupControlNumber          bool   `pulumi:"checkDuplicateGroupControlNumber"`
	CheckDuplicateInterchangeControlNumber    bool   `pulumi:"checkDuplicateInterchangeControlNumber"`
	CheckDuplicateTransactionSetControlNumber bool   `pulumi:"checkDuplicateTransactionSetControlNumber"`
	InterchangeControlNumberValidityDays      int    `pulumi:"interchangeControlNumberValidityDays"`
	TrailingSeparatorPolicy                   string `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes     bool   `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateCharacterSet                      bool   `pulumi:"validateCharacterSet"`
	ValidateEDITypes                          bool   `pulumi:"validateEDITypes"`
	ValidateXSDTypes                          bool   `pulumi:"validateXSDTypes"`
}





type X12ValidationSettingsInput interface {
	pulumi.Input

	ToX12ValidationSettingsOutput() X12ValidationSettingsOutput
	ToX12ValidationSettingsOutputWithContext(context.Context) X12ValidationSettingsOutput
}

type X12ValidationSettingsArgs struct {
	AllowLeadingAndTrailingSpacesAndZeroes    pulumi.BoolInput   `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	CheckDuplicateGroupControlNumber          pulumi.BoolInput   `pulumi:"checkDuplicateGroupControlNumber"`
	CheckDuplicateInterchangeControlNumber    pulumi.BoolInput   `pulumi:"checkDuplicateInterchangeControlNumber"`
	CheckDuplicateTransactionSetControlNumber pulumi.BoolInput   `pulumi:"checkDuplicateTransactionSetControlNumber"`
	InterchangeControlNumberValidityDays      pulumi.IntInput    `pulumi:"interchangeControlNumberValidityDays"`
	TrailingSeparatorPolicy                   pulumi.StringInput `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes     pulumi.BoolInput   `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateCharacterSet                      pulumi.BoolInput   `pulumi:"validateCharacterSet"`
	ValidateEDITypes                          pulumi.BoolInput   `pulumi:"validateEDITypes"`
	ValidateXSDTypes                          pulumi.BoolInput   `pulumi:"validateXSDTypes"`
}

func (X12ValidationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ValidationSettings)(nil)).Elem()
}

func (i X12ValidationSettingsArgs) ToX12ValidationSettingsOutput() X12ValidationSettingsOutput {
	return i.ToX12ValidationSettingsOutputWithContext(context.Background())
}

func (i X12ValidationSettingsArgs) ToX12ValidationSettingsOutputWithContext(ctx context.Context) X12ValidationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ValidationSettingsOutput)
}

func (i X12ValidationSettingsArgs) ToX12ValidationSettingsPtrOutput() X12ValidationSettingsPtrOutput {
	return i.ToX12ValidationSettingsPtrOutputWithContext(context.Background())
}

func (i X12ValidationSettingsArgs) ToX12ValidationSettingsPtrOutputWithContext(ctx context.Context) X12ValidationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ValidationSettingsOutput).ToX12ValidationSettingsPtrOutputWithContext(ctx)
}









type X12ValidationSettingsPtrInput interface {
	pulumi.Input

	ToX12ValidationSettingsPtrOutput() X12ValidationSettingsPtrOutput
	ToX12ValidationSettingsPtrOutputWithContext(context.Context) X12ValidationSettingsPtrOutput
}

type x12validationSettingsPtrType X12ValidationSettingsArgs

func X12ValidationSettingsPtr(v *X12ValidationSettingsArgs) X12ValidationSettingsPtrInput {
	return (*x12validationSettingsPtrType)(v)
}

func (*x12validationSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**X12ValidationSettings)(nil)).Elem()
}

func (i *x12validationSettingsPtrType) ToX12ValidationSettingsPtrOutput() X12ValidationSettingsPtrOutput {
	return i.ToX12ValidationSettingsPtrOutputWithContext(context.Background())
}

func (i *x12validationSettingsPtrType) ToX12ValidationSettingsPtrOutputWithContext(ctx context.Context) X12ValidationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(X12ValidationSettingsPtrOutput)
}

type X12ValidationSettingsOutput struct{ *pulumi.OutputState }

func (X12ValidationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ValidationSettings)(nil)).Elem()
}

func (o X12ValidationSettingsOutput) ToX12ValidationSettingsOutput() X12ValidationSettingsOutput {
	return o
}

func (o X12ValidationSettingsOutput) ToX12ValidationSettingsOutputWithContext(ctx context.Context) X12ValidationSettingsOutput {
	return o
}

func (o X12ValidationSettingsOutput) ToX12ValidationSettingsPtrOutput() X12ValidationSettingsPtrOutput {
	return o.ToX12ValidationSettingsPtrOutputWithContext(context.Background())
}

func (o X12ValidationSettingsOutput) ToX12ValidationSettingsPtrOutputWithContext(ctx context.Context) X12ValidationSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12ValidationSettings) *X12ValidationSettings {
		return &v
	}).(X12ValidationSettingsPtrOutput)
}

func (o X12ValidationSettingsOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationSettings) bool { return v.AllowLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolOutput)
}

func (o X12ValidationSettingsOutput) CheckDuplicateGroupControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationSettings) bool { return v.CheckDuplicateGroupControlNumber }).(pulumi.BoolOutput)
}

func (o X12ValidationSettingsOutput) CheckDuplicateInterchangeControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationSettings) bool { return v.CheckDuplicateInterchangeControlNumber }).(pulumi.BoolOutput)
}

func (o X12ValidationSettingsOutput) CheckDuplicateTransactionSetControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationSettings) bool { return v.CheckDuplicateTransactionSetControlNumber }).(pulumi.BoolOutput)
}

func (o X12ValidationSettingsOutput) InterchangeControlNumberValidityDays() pulumi.IntOutput {
	return o.ApplyT(func(v X12ValidationSettings) int { return v.InterchangeControlNumberValidityDays }).(pulumi.IntOutput)
}

func (o X12ValidationSettingsOutput) TrailingSeparatorPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v X12ValidationSettings) string { return v.TrailingSeparatorPolicy }).(pulumi.StringOutput)
}

func (o X12ValidationSettingsOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationSettings) bool { return v.TrimLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolOutput)
}

func (o X12ValidationSettingsOutput) ValidateCharacterSet() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationSettings) bool { return v.ValidateCharacterSet }).(pulumi.BoolOutput)
}

func (o X12ValidationSettingsOutput) ValidateEDITypes() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationSettings) bool { return v.ValidateEDITypes }).(pulumi.BoolOutput)
}

func (o X12ValidationSettingsOutput) ValidateXSDTypes() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationSettings) bool { return v.ValidateXSDTypes }).(pulumi.BoolOutput)
}

type X12ValidationSettingsPtrOutput struct{ *pulumi.OutputState }

func (X12ValidationSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12ValidationSettings)(nil)).Elem()
}

func (o X12ValidationSettingsPtrOutput) ToX12ValidationSettingsPtrOutput() X12ValidationSettingsPtrOutput {
	return o
}

func (o X12ValidationSettingsPtrOutput) ToX12ValidationSettingsPtrOutputWithContext(ctx context.Context) X12ValidationSettingsPtrOutput {
	return o
}

func (o X12ValidationSettingsPtrOutput) Elem() X12ValidationSettingsOutput {
	return o.ApplyT(func(v *X12ValidationSettings) X12ValidationSettings {
		if v != nil {
			return *v
		}
		var ret X12ValidationSettings
		return ret
	}).(X12ValidationSettingsOutput)
}

func (o X12ValidationSettingsPtrOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.AllowLeadingAndTrailingSpacesAndZeroes
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsPtrOutput) CheckDuplicateGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.CheckDuplicateGroupControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsPtrOutput) CheckDuplicateInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.CheckDuplicateInterchangeControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsPtrOutput) CheckDuplicateTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.CheckDuplicateTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsPtrOutput) InterchangeControlNumberValidityDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *int {
		if v == nil {
			return nil
		}
		return &v.InterchangeControlNumberValidityDays
	}).(pulumi.IntPtrOutput)
}

func (o X12ValidationSettingsPtrOutput) TrailingSeparatorPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *string {
		if v == nil {
			return nil
		}
		return &v.TrailingSeparatorPolicy
	}).(pulumi.StringPtrOutput)
}

func (o X12ValidationSettingsPtrOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.TrimLeadingAndTrailingSpacesAndZeroes
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsPtrOutput) ValidateCharacterSet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.ValidateCharacterSet
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsPtrOutput) ValidateEDITypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.ValidateEDITypes
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsPtrOutput) ValidateXSDTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.ValidateXSDTypes
	}).(pulumi.BoolPtrOutput)
}

type X12ValidationSettingsResponse struct {
	AllowLeadingAndTrailingSpacesAndZeroes    bool   `pulumi:"allowLeadingAndTrailingSpacesAndZeroes"`
	CheckDuplicateGroupControlNumber          bool   `pulumi:"checkDuplicateGroupControlNumber"`
	CheckDuplicateInterchangeControlNumber    bool   `pulumi:"checkDuplicateInterchangeControlNumber"`
	CheckDuplicateTransactionSetControlNumber bool   `pulumi:"checkDuplicateTransactionSetControlNumber"`
	InterchangeControlNumberValidityDays      int    `pulumi:"interchangeControlNumberValidityDays"`
	TrailingSeparatorPolicy                   string `pulumi:"trailingSeparatorPolicy"`
	TrimLeadingAndTrailingSpacesAndZeroes     bool   `pulumi:"trimLeadingAndTrailingSpacesAndZeroes"`
	ValidateCharacterSet                      bool   `pulumi:"validateCharacterSet"`
	ValidateEDITypes                          bool   `pulumi:"validateEDITypes"`
	ValidateXSDTypes                          bool   `pulumi:"validateXSDTypes"`
}

type X12ValidationSettingsResponseOutput struct{ *pulumi.OutputState }

func (X12ValidationSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12ValidationSettingsResponse)(nil)).Elem()
}

func (o X12ValidationSettingsResponseOutput) ToX12ValidationSettingsResponseOutput() X12ValidationSettingsResponseOutput {
	return o
}

func (o X12ValidationSettingsResponseOutput) ToX12ValidationSettingsResponseOutputWithContext(ctx context.Context) X12ValidationSettingsResponseOutput {
	return o
}

func (o X12ValidationSettingsResponseOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) bool { return v.AllowLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolOutput)
}

func (o X12ValidationSettingsResponseOutput) CheckDuplicateGroupControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) bool { return v.CheckDuplicateGroupControlNumber }).(pulumi.BoolOutput)
}

func (o X12ValidationSettingsResponseOutput) CheckDuplicateInterchangeControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) bool { return v.CheckDuplicateInterchangeControlNumber }).(pulumi.BoolOutput)
}

func (o X12ValidationSettingsResponseOutput) CheckDuplicateTransactionSetControlNumber() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) bool { return v.CheckDuplicateTransactionSetControlNumber }).(pulumi.BoolOutput)
}

func (o X12ValidationSettingsResponseOutput) InterchangeControlNumberValidityDays() pulumi.IntOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) int { return v.InterchangeControlNumberValidityDays }).(pulumi.IntOutput)
}

func (o X12ValidationSettingsResponseOutput) TrailingSeparatorPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) string { return v.TrailingSeparatorPolicy }).(pulumi.StringOutput)
}

func (o X12ValidationSettingsResponseOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) bool { return v.TrimLeadingAndTrailingSpacesAndZeroes }).(pulumi.BoolOutput)
}

func (o X12ValidationSettingsResponseOutput) ValidateCharacterSet() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) bool { return v.ValidateCharacterSet }).(pulumi.BoolOutput)
}

func (o X12ValidationSettingsResponseOutput) ValidateEDITypes() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) bool { return v.ValidateEDITypes }).(pulumi.BoolOutput)
}

func (o X12ValidationSettingsResponseOutput) ValidateXSDTypes() pulumi.BoolOutput {
	return o.ApplyT(func(v X12ValidationSettingsResponse) bool { return v.ValidateXSDTypes }).(pulumi.BoolOutput)
}

type X12ValidationSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (X12ValidationSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12ValidationSettingsResponse)(nil)).Elem()
}

func (o X12ValidationSettingsResponsePtrOutput) ToX12ValidationSettingsResponsePtrOutput() X12ValidationSettingsResponsePtrOutput {
	return o
}

func (o X12ValidationSettingsResponsePtrOutput) ToX12ValidationSettingsResponsePtrOutputWithContext(ctx context.Context) X12ValidationSettingsResponsePtrOutput {
	return o
}

func (o X12ValidationSettingsResponsePtrOutput) Elem() X12ValidationSettingsResponseOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) X12ValidationSettingsResponse {
		if v != nil {
			return *v
		}
		var ret X12ValidationSettingsResponse
		return ret
	}).(X12ValidationSettingsResponseOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) AllowLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.AllowLeadingAndTrailingSpacesAndZeroes
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) CheckDuplicateGroupControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CheckDuplicateGroupControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) CheckDuplicateInterchangeControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CheckDuplicateInterchangeControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) CheckDuplicateTransactionSetControlNumber() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CheckDuplicateTransactionSetControlNumber
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) InterchangeControlNumberValidityDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.InterchangeControlNumberValidityDays
	}).(pulumi.IntPtrOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) TrailingSeparatorPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TrailingSeparatorPolicy
	}).(pulumi.StringPtrOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) TrimLeadingAndTrailingSpacesAndZeroes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.TrimLeadingAndTrailingSpacesAndZeroes
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) ValidateCharacterSet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.ValidateCharacterSet
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) ValidateEDITypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.ValidateEDITypes
	}).(pulumi.BoolPtrOutput)
}

func (o X12ValidationSettingsResponsePtrOutput) ValidateXSDTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *X12ValidationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.ValidateXSDTypes
	}).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AS2AcknowledgementConnectionSettingsOutput{})
	pulumi.RegisterOutputType(AS2AcknowledgementConnectionSettingsPtrOutput{})
	pulumi.RegisterOutputType(AS2AcknowledgementConnectionSettingsResponseOutput{})
	pulumi.RegisterOutputType(AS2AcknowledgementConnectionSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AS2AgreementContentOutput{})
	pulumi.RegisterOutputType(AS2AgreementContentPtrOutput{})
	pulumi.RegisterOutputType(AS2AgreementContentResponseOutput{})
	pulumi.RegisterOutputType(AS2AgreementContentResponsePtrOutput{})
	pulumi.RegisterOutputType(AS2EnvelopeSettingsOutput{})
	pulumi.RegisterOutputType(AS2EnvelopeSettingsPtrOutput{})
	pulumi.RegisterOutputType(AS2EnvelopeSettingsResponseOutput{})
	pulumi.RegisterOutputType(AS2EnvelopeSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AS2ErrorSettingsOutput{})
	pulumi.RegisterOutputType(AS2ErrorSettingsPtrOutput{})
	pulumi.RegisterOutputType(AS2ErrorSettingsResponseOutput{})
	pulumi.RegisterOutputType(AS2ErrorSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AS2MdnSettingsOutput{})
	pulumi.RegisterOutputType(AS2MdnSettingsPtrOutput{})
	pulumi.RegisterOutputType(AS2MdnSettingsResponseOutput{})
	pulumi.RegisterOutputType(AS2MdnSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AS2MessageConnectionSettingsOutput{})
	pulumi.RegisterOutputType(AS2MessageConnectionSettingsPtrOutput{})
	pulumi.RegisterOutputType(AS2MessageConnectionSettingsResponseOutput{})
	pulumi.RegisterOutputType(AS2MessageConnectionSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AS2OneWayAgreementOutput{})
	pulumi.RegisterOutputType(AS2OneWayAgreementPtrOutput{})
	pulumi.RegisterOutputType(AS2OneWayAgreementResponseOutput{})
	pulumi.RegisterOutputType(AS2OneWayAgreementResponsePtrOutput{})
	pulumi.RegisterOutputType(AS2ProtocolSettingsOutput{})
	pulumi.RegisterOutputType(AS2ProtocolSettingsPtrOutput{})
	pulumi.RegisterOutputType(AS2ProtocolSettingsResponseOutput{})
	pulumi.RegisterOutputType(AS2ProtocolSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AS2SecuritySettingsOutput{})
	pulumi.RegisterOutputType(AS2SecuritySettingsPtrOutput{})
	pulumi.RegisterOutputType(AS2SecuritySettingsResponseOutput{})
	pulumi.RegisterOutputType(AS2SecuritySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AS2ValidationSettingsOutput{})
	pulumi.RegisterOutputType(AS2ValidationSettingsPtrOutput{})
	pulumi.RegisterOutputType(AS2ValidationSettingsResponseOutput{})
	pulumi.RegisterOutputType(AS2ValidationSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AgreementContentOutput{})
	pulumi.RegisterOutputType(AgreementContentResponseOutput{})
	pulumi.RegisterOutputType(ApiDeploymentParameterMetadataResponseOutput{})
	pulumi.RegisterOutputType(ApiDeploymentParameterMetadataResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiDeploymentParameterMetadataSetResponseOutput{})
	pulumi.RegisterOutputType(ApiDeploymentParameterMetadataSetResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiResourceBackendServiceResponseOutput{})
	pulumi.RegisterOutputType(ApiResourceDefinitionsResponseOutput{})
	pulumi.RegisterOutputType(ApiResourceGeneralInformationResponseOutput{})
	pulumi.RegisterOutputType(ApiResourceMetadataResponseOutput{})
	pulumi.RegisterOutputType(ApiResourcePoliciesResponseOutput{})
	pulumi.RegisterOutputType(AssemblyPropertiesOutput{})
	pulumi.RegisterOutputType(AssemblyPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AzureResourceErrorInfoResponseOutput{})
	pulumi.RegisterOutputType(AzureResourceErrorInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureResourceErrorInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(B2BPartnerContentOutput{})
	pulumi.RegisterOutputType(B2BPartnerContentPtrOutput{})
	pulumi.RegisterOutputType(B2BPartnerContentResponseOutput{})
	pulumi.RegisterOutputType(B2BPartnerContentResponsePtrOutput{})
	pulumi.RegisterOutputType(BatchConfigurationPropertiesOutput{})
	pulumi.RegisterOutputType(BatchConfigurationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(BatchReleaseCriteriaOutput{})
	pulumi.RegisterOutputType(BatchReleaseCriteriaResponseOutput{})
	pulumi.RegisterOutputType(BusinessIdentityOutput{})
	pulumi.RegisterOutputType(BusinessIdentityPtrOutput{})
	pulumi.RegisterOutputType(BusinessIdentityArrayOutput{})
	pulumi.RegisterOutputType(BusinessIdentityResponseOutput{})
	pulumi.RegisterOutputType(BusinessIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(BusinessIdentityResponseArrayOutput{})
	pulumi.RegisterOutputType(ContentHashResponseOutput{})
	pulumi.RegisterOutputType(ContentHashResponsePtrOutput{})
	pulumi.RegisterOutputType(ContentLinkOutput{})
	pulumi.RegisterOutputType(ContentLinkPtrOutput{})
	pulumi.RegisterOutputType(ContentLinkResponseOutput{})
	pulumi.RegisterOutputType(ContentLinkResponsePtrOutput{})
	pulumi.RegisterOutputType(EdifactAcknowledgementSettingsOutput{})
	pulumi.RegisterOutputType(EdifactAcknowledgementSettingsPtrOutput{})
	pulumi.RegisterOutputType(EdifactAcknowledgementSettingsResponseOutput{})
	pulumi.RegisterOutputType(EdifactAcknowledgementSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(EdifactAgreementContentOutput{})
	pulumi.RegisterOutputType(EdifactAgreementContentPtrOutput{})
	pulumi.RegisterOutputType(EdifactAgreementContentResponseOutput{})
	pulumi.RegisterOutputType(EdifactAgreementContentResponsePtrOutput{})
	pulumi.RegisterOutputType(EdifactDelimiterOverrideOutput{})
	pulumi.RegisterOutputType(EdifactDelimiterOverrideArrayOutput{})
	pulumi.RegisterOutputType(EdifactDelimiterOverrideResponseOutput{})
	pulumi.RegisterOutputType(EdifactDelimiterOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(EdifactEnvelopeOverrideOutput{})
	pulumi.RegisterOutputType(EdifactEnvelopeOverrideArrayOutput{})
	pulumi.RegisterOutputType(EdifactEnvelopeOverrideResponseOutput{})
	pulumi.RegisterOutputType(EdifactEnvelopeOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(EdifactEnvelopeSettingsOutput{})
	pulumi.RegisterOutputType(EdifactEnvelopeSettingsPtrOutput{})
	pulumi.RegisterOutputType(EdifactEnvelopeSettingsResponseOutput{})
	pulumi.RegisterOutputType(EdifactEnvelopeSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(EdifactFramingSettingsOutput{})
	pulumi.RegisterOutputType(EdifactFramingSettingsPtrOutput{})
	pulumi.RegisterOutputType(EdifactFramingSettingsResponseOutput{})
	pulumi.RegisterOutputType(EdifactFramingSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(EdifactMessageFilterOutput{})
	pulumi.RegisterOutputType(EdifactMessageFilterPtrOutput{})
	pulumi.RegisterOutputType(EdifactMessageFilterResponseOutput{})
	pulumi.RegisterOutputType(EdifactMessageFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(EdifactMessageIdentifierOutput{})
	pulumi.RegisterOutputType(EdifactMessageIdentifierArrayOutput{})
	pulumi.RegisterOutputType(EdifactMessageIdentifierResponseOutput{})
	pulumi.RegisterOutputType(EdifactMessageIdentifierResponseArrayOutput{})
	pulumi.RegisterOutputType(EdifactOneWayAgreementOutput{})
	pulumi.RegisterOutputType(EdifactOneWayAgreementPtrOutput{})
	pulumi.RegisterOutputType(EdifactOneWayAgreementResponseOutput{})
	pulumi.RegisterOutputType(EdifactOneWayAgreementResponsePtrOutput{})
	pulumi.RegisterOutputType(EdifactProcessingSettingsOutput{})
	pulumi.RegisterOutputType(EdifactProcessingSettingsPtrOutput{})
	pulumi.RegisterOutputType(EdifactProcessingSettingsResponseOutput{})
	pulumi.RegisterOutputType(EdifactProcessingSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(EdifactProtocolSettingsOutput{})
	pulumi.RegisterOutputType(EdifactProtocolSettingsPtrOutput{})
	pulumi.RegisterOutputType(EdifactProtocolSettingsResponseOutput{})
	pulumi.RegisterOutputType(EdifactProtocolSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(EdifactSchemaReferenceOutput{})
	pulumi.RegisterOutputType(EdifactSchemaReferenceArrayOutput{})
	pulumi.RegisterOutputType(EdifactSchemaReferenceResponseOutput{})
	pulumi.RegisterOutputType(EdifactSchemaReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(EdifactValidationOverrideOutput{})
	pulumi.RegisterOutputType(EdifactValidationOverrideArrayOutput{})
	pulumi.RegisterOutputType(EdifactValidationOverrideResponseOutput{})
	pulumi.RegisterOutputType(EdifactValidationOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(EdifactValidationSettingsOutput{})
	pulumi.RegisterOutputType(EdifactValidationSettingsPtrOutput{})
	pulumi.RegisterOutputType(EdifactValidationSettingsResponseOutput{})
	pulumi.RegisterOutputType(EdifactValidationSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ExpressionResponseOutput{})
	pulumi.RegisterOutputType(ExpressionResponseArrayOutput{})
	pulumi.RegisterOutputType(ExpressionRootResponseOutput{})
	pulumi.RegisterOutputType(ExpressionRootResponseArrayOutput{})
	pulumi.RegisterOutputType(FlowAccessControlConfigurationOutput{})
	pulumi.RegisterOutputType(FlowAccessControlConfigurationPtrOutput{})
	pulumi.RegisterOutputType(FlowAccessControlConfigurationPolicyOutput{})
	pulumi.RegisterOutputType(FlowAccessControlConfigurationPolicyPtrOutput{})
	pulumi.RegisterOutputType(FlowAccessControlConfigurationPolicyResponseOutput{})
	pulumi.RegisterOutputType(FlowAccessControlConfigurationPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(FlowAccessControlConfigurationResponseOutput{})
	pulumi.RegisterOutputType(FlowAccessControlConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(FlowEndpointsOutput{})
	pulumi.RegisterOutputType(FlowEndpointsPtrOutput{})
	pulumi.RegisterOutputType(FlowEndpointsConfigurationOutput{})
	pulumi.RegisterOutputType(FlowEndpointsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(FlowEndpointsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(FlowEndpointsConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(FlowEndpointsResponseOutput{})
	pulumi.RegisterOutputType(FlowEndpointsResponsePtrOutput{})
	pulumi.RegisterOutputType(IntegrationAccountMapPropertiesParametersSchemaOutput{})
	pulumi.RegisterOutputType(IntegrationAccountMapPropertiesParametersSchemaPtrOutput{})
	pulumi.RegisterOutputType(IntegrationAccountMapPropertiesResponseParametersSchemaOutput{})
	pulumi.RegisterOutputType(IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput{})
	pulumi.RegisterOutputType(IntegrationAccountSkuOutput{})
	pulumi.RegisterOutputType(IntegrationAccountSkuPtrOutput{})
	pulumi.RegisterOutputType(IntegrationAccountSkuResponseOutput{})
	pulumi.RegisterOutputType(IntegrationAccountSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmenEncryptionConfigurationOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmenEncryptionConfigurationPtrOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmenEncryptionConfigurationResponseOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmenEncryptionConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmenEncryptionKeyReferenceOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmenEncryptionKeyReferencePtrOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmenEncryptionKeyReferenceResponseOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmenEncryptionKeyReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentAccessEndpointOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentAccessEndpointPtrOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentAccessEndpointResponseOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentAccessEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentManagedApiDeploymentParametersOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentManagedApiDeploymentParametersPtrOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentManagedApiDeploymentParametersResponseOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentManagedApiDeploymentParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentPropertiesOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentSkuOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentSkuPtrOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentSkuResponseOutput{})
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(IpAddressOutput{})
	pulumi.RegisterOutputType(IpAddressArrayOutput{})
	pulumi.RegisterOutputType(IpAddressRangeOutput{})
	pulumi.RegisterOutputType(IpAddressRangeArrayOutput{})
	pulumi.RegisterOutputType(IpAddressRangeResponseOutput{})
	pulumi.RegisterOutputType(IpAddressRangeResponseArrayOutput{})
	pulumi.RegisterOutputType(IpAddressResponseOutput{})
	pulumi.RegisterOutputType(IpAddressResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceKeyVaultOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceKeyVaultPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceResponseKeyVaultOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceResponseKeyVaultPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyResponseAttributesOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyResponseAttributesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultReferenceOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkConfigurationOutput{})
	pulumi.RegisterOutputType(NetworkConfigurationPtrOutput{})
	pulumi.RegisterOutputType(NetworkConfigurationResponseOutput{})
	pulumi.RegisterOutputType(NetworkConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(OpenAuthenticationAccessPoliciesOutput{})
	pulumi.RegisterOutputType(OpenAuthenticationAccessPoliciesPtrOutput{})
	pulumi.RegisterOutputType(OpenAuthenticationAccessPoliciesResponseOutput{})
	pulumi.RegisterOutputType(OpenAuthenticationAccessPoliciesResponsePtrOutput{})
	pulumi.RegisterOutputType(OpenAuthenticationAccessPolicyOutput{})
	pulumi.RegisterOutputType(OpenAuthenticationAccessPolicyMapOutput{})
	pulumi.RegisterOutputType(OpenAuthenticationAccessPolicyResponseOutput{})
	pulumi.RegisterOutputType(OpenAuthenticationAccessPolicyResponseMapOutput{})
	pulumi.RegisterOutputType(OpenAuthenticationPolicyClaimOutput{})
	pulumi.RegisterOutputType(OpenAuthenticationPolicyClaimArrayOutput{})
	pulumi.RegisterOutputType(OpenAuthenticationPolicyClaimResponseOutput{})
	pulumi.RegisterOutputType(OpenAuthenticationPolicyClaimResponseArrayOutput{})
	pulumi.RegisterOutputType(PartnerContentOutput{})
	pulumi.RegisterOutputType(PartnerContentResponseOutput{})
	pulumi.RegisterOutputType(RecurrenceScheduleOutput{})
	pulumi.RegisterOutputType(RecurrenceSchedulePtrOutput{})
	pulumi.RegisterOutputType(RecurrenceScheduleOccurrenceOutput{})
	pulumi.RegisterOutputType(RecurrenceScheduleOccurrenceArrayOutput{})
	pulumi.RegisterOutputType(RecurrenceScheduleOccurrenceResponseOutput{})
	pulumi.RegisterOutputType(RecurrenceScheduleOccurrenceResponseArrayOutput{})
	pulumi.RegisterOutputType(RecurrenceScheduleResponseOutput{})
	pulumi.RegisterOutputType(RecurrenceScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceReferenceOutput{})
	pulumi.RegisterOutputType(ResourceReferencePtrOutput{})
	pulumi.RegisterOutputType(ResourceReferenceArrayOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponseOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(RosettaNetPipAcknowledgmentOfReceiptSettingsOutput{})
	pulumi.RegisterOutputType(RosettaNetPipAcknowledgmentOfReceiptSettingsResponseOutput{})
	pulumi.RegisterOutputType(RosettaNetPipActivityBehaviorOutput{})
	pulumi.RegisterOutputType(RosettaNetPipActivityBehaviorResponseOutput{})
	pulumi.RegisterOutputType(RosettaNetPipActivitySettingsOutput{})
	pulumi.RegisterOutputType(RosettaNetPipActivitySettingsResponseOutput{})
	pulumi.RegisterOutputType(RosettaNetPipBusinessDocumentOutput{})
	pulumi.RegisterOutputType(RosettaNetPipBusinessDocumentResponseOutput{})
	pulumi.RegisterOutputType(RosettaNetPipRoleSettingsOutput{})
	pulumi.RegisterOutputType(RosettaNetPipRoleSettingsResponseOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(WorkflowParameterOutput{})
	pulumi.RegisterOutputType(WorkflowParameterMapOutput{})
	pulumi.RegisterOutputType(WorkflowParameterResponseOutput{})
	pulumi.RegisterOutputType(WorkflowParameterResponseMapOutput{})
	pulumi.RegisterOutputType(WorkflowTriggerListCallbackUrlQueriesResponseOutput{})
	pulumi.RegisterOutputType(WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkflowTriggerRecurrenceOutput{})
	pulumi.RegisterOutputType(WorkflowTriggerRecurrencePtrOutput{})
	pulumi.RegisterOutputType(WorkflowTriggerRecurrenceResponseOutput{})
	pulumi.RegisterOutputType(WorkflowTriggerRecurrenceResponsePtrOutput{})
	pulumi.RegisterOutputType(WsdlServiceResponseOutput{})
	pulumi.RegisterOutputType(WsdlServiceResponsePtrOutput{})
	pulumi.RegisterOutputType(X12AcknowledgementSettingsOutput{})
	pulumi.RegisterOutputType(X12AcknowledgementSettingsPtrOutput{})
	pulumi.RegisterOutputType(X12AcknowledgementSettingsResponseOutput{})
	pulumi.RegisterOutputType(X12AcknowledgementSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(X12AgreementContentOutput{})
	pulumi.RegisterOutputType(X12AgreementContentPtrOutput{})
	pulumi.RegisterOutputType(X12AgreementContentResponseOutput{})
	pulumi.RegisterOutputType(X12AgreementContentResponsePtrOutput{})
	pulumi.RegisterOutputType(X12DelimiterOverridesOutput{})
	pulumi.RegisterOutputType(X12DelimiterOverridesArrayOutput{})
	pulumi.RegisterOutputType(X12DelimiterOverridesResponseOutput{})
	pulumi.RegisterOutputType(X12DelimiterOverridesResponseArrayOutput{})
	pulumi.RegisterOutputType(X12EnvelopeOverrideOutput{})
	pulumi.RegisterOutputType(X12EnvelopeOverrideArrayOutput{})
	pulumi.RegisterOutputType(X12EnvelopeOverrideResponseOutput{})
	pulumi.RegisterOutputType(X12EnvelopeOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(X12EnvelopeSettingsOutput{})
	pulumi.RegisterOutputType(X12EnvelopeSettingsPtrOutput{})
	pulumi.RegisterOutputType(X12EnvelopeSettingsResponseOutput{})
	pulumi.RegisterOutputType(X12EnvelopeSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(X12FramingSettingsOutput{})
	pulumi.RegisterOutputType(X12FramingSettingsPtrOutput{})
	pulumi.RegisterOutputType(X12FramingSettingsResponseOutput{})
	pulumi.RegisterOutputType(X12FramingSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(X12MessageFilterOutput{})
	pulumi.RegisterOutputType(X12MessageFilterPtrOutput{})
	pulumi.RegisterOutputType(X12MessageFilterResponseOutput{})
	pulumi.RegisterOutputType(X12MessageFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(X12MessageIdentifierOutput{})
	pulumi.RegisterOutputType(X12MessageIdentifierArrayOutput{})
	pulumi.RegisterOutputType(X12MessageIdentifierResponseOutput{})
	pulumi.RegisterOutputType(X12MessageIdentifierResponseArrayOutput{})
	pulumi.RegisterOutputType(X12OneWayAgreementOutput{})
	pulumi.RegisterOutputType(X12OneWayAgreementPtrOutput{})
	pulumi.RegisterOutputType(X12OneWayAgreementResponseOutput{})
	pulumi.RegisterOutputType(X12OneWayAgreementResponsePtrOutput{})
	pulumi.RegisterOutputType(X12ProcessingSettingsOutput{})
	pulumi.RegisterOutputType(X12ProcessingSettingsPtrOutput{})
	pulumi.RegisterOutputType(X12ProcessingSettingsResponseOutput{})
	pulumi.RegisterOutputType(X12ProcessingSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(X12ProtocolSettingsOutput{})
	pulumi.RegisterOutputType(X12ProtocolSettingsPtrOutput{})
	pulumi.RegisterOutputType(X12ProtocolSettingsResponseOutput{})
	pulumi.RegisterOutputType(X12ProtocolSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(X12SchemaReferenceOutput{})
	pulumi.RegisterOutputType(X12SchemaReferenceArrayOutput{})
	pulumi.RegisterOutputType(X12SchemaReferenceResponseOutput{})
	pulumi.RegisterOutputType(X12SchemaReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(X12SecuritySettingsOutput{})
	pulumi.RegisterOutputType(X12SecuritySettingsPtrOutput{})
	pulumi.RegisterOutputType(X12SecuritySettingsResponseOutput{})
	pulumi.RegisterOutputType(X12SecuritySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(X12ValidationOverrideOutput{})
	pulumi.RegisterOutputType(X12ValidationOverrideArrayOutput{})
	pulumi.RegisterOutputType(X12ValidationOverrideResponseOutput{})
	pulumi.RegisterOutputType(X12ValidationOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(X12ValidationSettingsOutput{})
	pulumi.RegisterOutputType(X12ValidationSettingsPtrOutput{})
	pulumi.RegisterOutputType(X12ValidationSettingsResponseOutput{})
	pulumi.RegisterOutputType(X12ValidationSettingsResponsePtrOutput{})
}
