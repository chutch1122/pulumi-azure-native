


package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AgreementType string

const (
	AgreementTypeNotSpecified = AgreementType("NotSpecified")
	AgreementTypeAS2          = AgreementType("AS2")
	AgreementTypeX12          = AgreementType("X12")
	AgreementTypeEdifact      = AgreementType("Edifact")
)

func (AgreementType) ElementType() reflect.Type {
	return reflect.TypeOf((*AgreementType)(nil)).Elem()
}

func (e AgreementType) ToAgreementTypeOutput() AgreementTypeOutput {
	return pulumi.ToOutput(e).(AgreementTypeOutput)
}

func (e AgreementType) ToAgreementTypeOutputWithContext(ctx context.Context) AgreementTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AgreementTypeOutput)
}

func (e AgreementType) ToAgreementTypePtrOutput() AgreementTypePtrOutput {
	return e.ToAgreementTypePtrOutputWithContext(context.Background())
}

func (e AgreementType) ToAgreementTypePtrOutputWithContext(ctx context.Context) AgreementTypePtrOutput {
	return AgreementType(e).ToAgreementTypeOutputWithContext(ctx).ToAgreementTypePtrOutputWithContext(ctx)
}

func (e AgreementType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AgreementType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AgreementType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AgreementType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AgreementTypeOutput struct{ *pulumi.OutputState }

func (AgreementTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgreementType)(nil)).Elem()
}

func (o AgreementTypeOutput) ToAgreementTypeOutput() AgreementTypeOutput {
	return o
}

func (o AgreementTypeOutput) ToAgreementTypeOutputWithContext(ctx context.Context) AgreementTypeOutput {
	return o
}

func (o AgreementTypeOutput) ToAgreementTypePtrOutput() AgreementTypePtrOutput {
	return o.ToAgreementTypePtrOutputWithContext(context.Background())
}

func (o AgreementTypeOutput) ToAgreementTypePtrOutputWithContext(ctx context.Context) AgreementTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AgreementType) *AgreementType {
		return &v
	}).(AgreementTypePtrOutput)
}

func (o AgreementTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AgreementTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AgreementType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AgreementTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AgreementTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AgreementType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AgreementTypePtrOutput struct{ *pulumi.OutputState }

func (AgreementTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgreementType)(nil)).Elem()
}

func (o AgreementTypePtrOutput) ToAgreementTypePtrOutput() AgreementTypePtrOutput {
	return o
}

func (o AgreementTypePtrOutput) ToAgreementTypePtrOutputWithContext(ctx context.Context) AgreementTypePtrOutput {
	return o
}

func (o AgreementTypePtrOutput) Elem() AgreementTypeOutput {
	return o.ApplyT(func(v *AgreementType) AgreementType {
		if v != nil {
			return *v
		}
		var ret AgreementType
		return ret
	}).(AgreementTypeOutput)
}

func (o AgreementTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AgreementTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AgreementType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AgreementTypeInput interface {
	pulumi.Input

	ToAgreementTypeOutput() AgreementTypeOutput
	ToAgreementTypeOutputWithContext(context.Context) AgreementTypeOutput
}

var agreementTypePtrType = reflect.TypeOf((**AgreementType)(nil)).Elem()

type AgreementTypePtrInput interface {
	pulumi.Input

	ToAgreementTypePtrOutput() AgreementTypePtrOutput
	ToAgreementTypePtrOutputWithContext(context.Context) AgreementTypePtrOutput
}

type agreementTypePtr string

func AgreementTypePtr(v string) AgreementTypePtrInput {
	return (*agreementTypePtr)(&v)
}

func (*agreementTypePtr) ElementType() reflect.Type {
	return agreementTypePtrType
}

func (in *agreementTypePtr) ToAgreementTypePtrOutput() AgreementTypePtrOutput {
	return pulumi.ToOutput(in).(AgreementTypePtrOutput)
}

func (in *agreementTypePtr) ToAgreementTypePtrOutputWithContext(ctx context.Context) AgreementTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AgreementTypePtrOutput)
}

type EdifactCharacterSet string

const (
	EdifactCharacterSetNotSpecified = EdifactCharacterSet("NotSpecified")
	EdifactCharacterSetUNOB         = EdifactCharacterSet("UNOB")
	EdifactCharacterSetUNOA         = EdifactCharacterSet("UNOA")
	EdifactCharacterSetUNOC         = EdifactCharacterSet("UNOC")
	EdifactCharacterSetUNOD         = EdifactCharacterSet("UNOD")
	EdifactCharacterSetUNOE         = EdifactCharacterSet("UNOE")
	EdifactCharacterSetUNOF         = EdifactCharacterSet("UNOF")
	EdifactCharacterSetUNOG         = EdifactCharacterSet("UNOG")
	EdifactCharacterSetUNOH         = EdifactCharacterSet("UNOH")
	EdifactCharacterSetUNOI         = EdifactCharacterSet("UNOI")
	EdifactCharacterSetUNOJ         = EdifactCharacterSet("UNOJ")
	EdifactCharacterSetUNOK         = EdifactCharacterSet("UNOK")
	EdifactCharacterSetUNOX         = EdifactCharacterSet("UNOX")
	EdifactCharacterSetUNOY         = EdifactCharacterSet("UNOY")
	EdifactCharacterSetKECA         = EdifactCharacterSet("KECA")
)

func (EdifactCharacterSet) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactCharacterSet)(nil)).Elem()
}

func (e EdifactCharacterSet) ToEdifactCharacterSetOutput() EdifactCharacterSetOutput {
	return pulumi.ToOutput(e).(EdifactCharacterSetOutput)
}

func (e EdifactCharacterSet) ToEdifactCharacterSetOutputWithContext(ctx context.Context) EdifactCharacterSetOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EdifactCharacterSetOutput)
}

func (e EdifactCharacterSet) ToEdifactCharacterSetPtrOutput() EdifactCharacterSetPtrOutput {
	return e.ToEdifactCharacterSetPtrOutputWithContext(context.Background())
}

func (e EdifactCharacterSet) ToEdifactCharacterSetPtrOutputWithContext(ctx context.Context) EdifactCharacterSetPtrOutput {
	return EdifactCharacterSet(e).ToEdifactCharacterSetOutputWithContext(ctx).ToEdifactCharacterSetPtrOutputWithContext(ctx)
}

func (e EdifactCharacterSet) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EdifactCharacterSet) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EdifactCharacterSet) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EdifactCharacterSet) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EdifactCharacterSetOutput struct{ *pulumi.OutputState }

func (EdifactCharacterSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactCharacterSet)(nil)).Elem()
}

func (o EdifactCharacterSetOutput) ToEdifactCharacterSetOutput() EdifactCharacterSetOutput {
	return o
}

func (o EdifactCharacterSetOutput) ToEdifactCharacterSetOutputWithContext(ctx context.Context) EdifactCharacterSetOutput {
	return o
}

func (o EdifactCharacterSetOutput) ToEdifactCharacterSetPtrOutput() EdifactCharacterSetPtrOutput {
	return o.ToEdifactCharacterSetPtrOutputWithContext(context.Background())
}

func (o EdifactCharacterSetOutput) ToEdifactCharacterSetPtrOutputWithContext(ctx context.Context) EdifactCharacterSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactCharacterSet) *EdifactCharacterSet {
		return &v
	}).(EdifactCharacterSetPtrOutput)
}

func (o EdifactCharacterSetOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EdifactCharacterSetOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EdifactCharacterSet) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EdifactCharacterSetOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EdifactCharacterSetOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EdifactCharacterSet) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EdifactCharacterSetPtrOutput struct{ *pulumi.OutputState }

func (EdifactCharacterSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactCharacterSet)(nil)).Elem()
}

func (o EdifactCharacterSetPtrOutput) ToEdifactCharacterSetPtrOutput() EdifactCharacterSetPtrOutput {
	return o
}

func (o EdifactCharacterSetPtrOutput) ToEdifactCharacterSetPtrOutputWithContext(ctx context.Context) EdifactCharacterSetPtrOutput {
	return o
}

func (o EdifactCharacterSetPtrOutput) Elem() EdifactCharacterSetOutput {
	return o.ApplyT(func(v *EdifactCharacterSet) EdifactCharacterSet {
		if v != nil {
			return *v
		}
		var ret EdifactCharacterSet
		return ret
	}).(EdifactCharacterSetOutput)
}

func (o EdifactCharacterSetPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EdifactCharacterSetPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EdifactCharacterSet) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EdifactCharacterSetInput interface {
	pulumi.Input

	ToEdifactCharacterSetOutput() EdifactCharacterSetOutput
	ToEdifactCharacterSetOutputWithContext(context.Context) EdifactCharacterSetOutput
}

var edifactCharacterSetPtrType = reflect.TypeOf((**EdifactCharacterSet)(nil)).Elem()

type EdifactCharacterSetPtrInput interface {
	pulumi.Input

	ToEdifactCharacterSetPtrOutput() EdifactCharacterSetPtrOutput
	ToEdifactCharacterSetPtrOutputWithContext(context.Context) EdifactCharacterSetPtrOutput
}

type edifactCharacterSetPtr string

func EdifactCharacterSetPtr(v string) EdifactCharacterSetPtrInput {
	return (*edifactCharacterSetPtr)(&v)
}

func (*edifactCharacterSetPtr) ElementType() reflect.Type {
	return edifactCharacterSetPtrType
}

func (in *edifactCharacterSetPtr) ToEdifactCharacterSetPtrOutput() EdifactCharacterSetPtrOutput {
	return pulumi.ToOutput(in).(EdifactCharacterSetPtrOutput)
}

func (in *edifactCharacterSetPtr) ToEdifactCharacterSetPtrOutputWithContext(ctx context.Context) EdifactCharacterSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EdifactCharacterSetPtrOutput)
}

type EdifactDecimalIndicator string

const (
	EdifactDecimalIndicatorNotSpecified = EdifactDecimalIndicator("NotSpecified")
	EdifactDecimalIndicatorComma        = EdifactDecimalIndicator("Comma")
	EdifactDecimalIndicatorDecimal      = EdifactDecimalIndicator("Decimal")
)

func (EdifactDecimalIndicator) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactDecimalIndicator)(nil)).Elem()
}

func (e EdifactDecimalIndicator) ToEdifactDecimalIndicatorOutput() EdifactDecimalIndicatorOutput {
	return pulumi.ToOutput(e).(EdifactDecimalIndicatorOutput)
}

func (e EdifactDecimalIndicator) ToEdifactDecimalIndicatorOutputWithContext(ctx context.Context) EdifactDecimalIndicatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EdifactDecimalIndicatorOutput)
}

func (e EdifactDecimalIndicator) ToEdifactDecimalIndicatorPtrOutput() EdifactDecimalIndicatorPtrOutput {
	return e.ToEdifactDecimalIndicatorPtrOutputWithContext(context.Background())
}

func (e EdifactDecimalIndicator) ToEdifactDecimalIndicatorPtrOutputWithContext(ctx context.Context) EdifactDecimalIndicatorPtrOutput {
	return EdifactDecimalIndicator(e).ToEdifactDecimalIndicatorOutputWithContext(ctx).ToEdifactDecimalIndicatorPtrOutputWithContext(ctx)
}

func (e EdifactDecimalIndicator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EdifactDecimalIndicator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EdifactDecimalIndicator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EdifactDecimalIndicator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EdifactDecimalIndicatorOutput struct{ *pulumi.OutputState }

func (EdifactDecimalIndicatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdifactDecimalIndicator)(nil)).Elem()
}

func (o EdifactDecimalIndicatorOutput) ToEdifactDecimalIndicatorOutput() EdifactDecimalIndicatorOutput {
	return o
}

func (o EdifactDecimalIndicatorOutput) ToEdifactDecimalIndicatorOutputWithContext(ctx context.Context) EdifactDecimalIndicatorOutput {
	return o
}

func (o EdifactDecimalIndicatorOutput) ToEdifactDecimalIndicatorPtrOutput() EdifactDecimalIndicatorPtrOutput {
	return o.ToEdifactDecimalIndicatorPtrOutputWithContext(context.Background())
}

func (o EdifactDecimalIndicatorOutput) ToEdifactDecimalIndicatorPtrOutputWithContext(ctx context.Context) EdifactDecimalIndicatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EdifactDecimalIndicator) *EdifactDecimalIndicator {
		return &v
	}).(EdifactDecimalIndicatorPtrOutput)
}

func (o EdifactDecimalIndicatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EdifactDecimalIndicatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EdifactDecimalIndicator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EdifactDecimalIndicatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EdifactDecimalIndicatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EdifactDecimalIndicator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EdifactDecimalIndicatorPtrOutput struct{ *pulumi.OutputState }

func (EdifactDecimalIndicatorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdifactDecimalIndicator)(nil)).Elem()
}

func (o EdifactDecimalIndicatorPtrOutput) ToEdifactDecimalIndicatorPtrOutput() EdifactDecimalIndicatorPtrOutput {
	return o
}

func (o EdifactDecimalIndicatorPtrOutput) ToEdifactDecimalIndicatorPtrOutputWithContext(ctx context.Context) EdifactDecimalIndicatorPtrOutput {
	return o
}

func (o EdifactDecimalIndicatorPtrOutput) Elem() EdifactDecimalIndicatorOutput {
	return o.ApplyT(func(v *EdifactDecimalIndicator) EdifactDecimalIndicator {
		if v != nil {
			return *v
		}
		var ret EdifactDecimalIndicator
		return ret
	}).(EdifactDecimalIndicatorOutput)
}

func (o EdifactDecimalIndicatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EdifactDecimalIndicatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EdifactDecimalIndicator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EdifactDecimalIndicatorInput interface {
	pulumi.Input

	ToEdifactDecimalIndicatorOutput() EdifactDecimalIndicatorOutput
	ToEdifactDecimalIndicatorOutputWithContext(context.Context) EdifactDecimalIndicatorOutput
}

var edifactDecimalIndicatorPtrType = reflect.TypeOf((**EdifactDecimalIndicator)(nil)).Elem()

type EdifactDecimalIndicatorPtrInput interface {
	pulumi.Input

	ToEdifactDecimalIndicatorPtrOutput() EdifactDecimalIndicatorPtrOutput
	ToEdifactDecimalIndicatorPtrOutputWithContext(context.Context) EdifactDecimalIndicatorPtrOutput
}

type edifactDecimalIndicatorPtr string

func EdifactDecimalIndicatorPtr(v string) EdifactDecimalIndicatorPtrInput {
	return (*edifactDecimalIndicatorPtr)(&v)
}

func (*edifactDecimalIndicatorPtr) ElementType() reflect.Type {
	return edifactDecimalIndicatorPtrType
}

func (in *edifactDecimalIndicatorPtr) ToEdifactDecimalIndicatorPtrOutput() EdifactDecimalIndicatorPtrOutput {
	return pulumi.ToOutput(in).(EdifactDecimalIndicatorPtrOutput)
}

func (in *edifactDecimalIndicatorPtr) ToEdifactDecimalIndicatorPtrOutputWithContext(ctx context.Context) EdifactDecimalIndicatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EdifactDecimalIndicatorPtrOutput)
}

type EncryptionAlgorithm string

const (
	EncryptionAlgorithmNotSpecified = EncryptionAlgorithm("NotSpecified")
	EncryptionAlgorithmNone         = EncryptionAlgorithm("None")
	EncryptionAlgorithmDES3         = EncryptionAlgorithm("DES3")
	EncryptionAlgorithmRC2          = EncryptionAlgorithm("RC2")
	EncryptionAlgorithmAES128       = EncryptionAlgorithm("AES128")
	EncryptionAlgorithmAES192       = EncryptionAlgorithm("AES192")
	EncryptionAlgorithmAES256       = EncryptionAlgorithm("AES256")
)

func (EncryptionAlgorithm) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionAlgorithm)(nil)).Elem()
}

func (e EncryptionAlgorithm) ToEncryptionAlgorithmOutput() EncryptionAlgorithmOutput {
	return pulumi.ToOutput(e).(EncryptionAlgorithmOutput)
}

func (e EncryptionAlgorithm) ToEncryptionAlgorithmOutputWithContext(ctx context.Context) EncryptionAlgorithmOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EncryptionAlgorithmOutput)
}

func (e EncryptionAlgorithm) ToEncryptionAlgorithmPtrOutput() EncryptionAlgorithmPtrOutput {
	return e.ToEncryptionAlgorithmPtrOutputWithContext(context.Background())
}

func (e EncryptionAlgorithm) ToEncryptionAlgorithmPtrOutputWithContext(ctx context.Context) EncryptionAlgorithmPtrOutput {
	return EncryptionAlgorithm(e).ToEncryptionAlgorithmOutputWithContext(ctx).ToEncryptionAlgorithmPtrOutputWithContext(ctx)
}

func (e EncryptionAlgorithm) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionAlgorithm) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EncryptionAlgorithm) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EncryptionAlgorithm) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EncryptionAlgorithmOutput struct{ *pulumi.OutputState }

func (EncryptionAlgorithmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionAlgorithm)(nil)).Elem()
}

func (o EncryptionAlgorithmOutput) ToEncryptionAlgorithmOutput() EncryptionAlgorithmOutput {
	return o
}

func (o EncryptionAlgorithmOutput) ToEncryptionAlgorithmOutputWithContext(ctx context.Context) EncryptionAlgorithmOutput {
	return o
}

func (o EncryptionAlgorithmOutput) ToEncryptionAlgorithmPtrOutput() EncryptionAlgorithmPtrOutput {
	return o.ToEncryptionAlgorithmPtrOutputWithContext(context.Background())
}

func (o EncryptionAlgorithmOutput) ToEncryptionAlgorithmPtrOutputWithContext(ctx context.Context) EncryptionAlgorithmPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionAlgorithm) *EncryptionAlgorithm {
		return &v
	}).(EncryptionAlgorithmPtrOutput)
}

func (o EncryptionAlgorithmOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EncryptionAlgorithmOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionAlgorithm) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EncryptionAlgorithmOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionAlgorithmOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EncryptionAlgorithm) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EncryptionAlgorithmPtrOutput struct{ *pulumi.OutputState }

func (EncryptionAlgorithmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionAlgorithm)(nil)).Elem()
}

func (o EncryptionAlgorithmPtrOutput) ToEncryptionAlgorithmPtrOutput() EncryptionAlgorithmPtrOutput {
	return o
}

func (o EncryptionAlgorithmPtrOutput) ToEncryptionAlgorithmPtrOutputWithContext(ctx context.Context) EncryptionAlgorithmPtrOutput {
	return o
}

func (o EncryptionAlgorithmPtrOutput) Elem() EncryptionAlgorithmOutput {
	return o.ApplyT(func(v *EncryptionAlgorithm) EncryptionAlgorithm {
		if v != nil {
			return *v
		}
		var ret EncryptionAlgorithm
		return ret
	}).(EncryptionAlgorithmOutput)
}

func (o EncryptionAlgorithmPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncryptionAlgorithmPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EncryptionAlgorithm) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EncryptionAlgorithmInput interface {
	pulumi.Input

	ToEncryptionAlgorithmOutput() EncryptionAlgorithmOutput
	ToEncryptionAlgorithmOutputWithContext(context.Context) EncryptionAlgorithmOutput
}

var encryptionAlgorithmPtrType = reflect.TypeOf((**EncryptionAlgorithm)(nil)).Elem()

type EncryptionAlgorithmPtrInput interface {
	pulumi.Input

	ToEncryptionAlgorithmPtrOutput() EncryptionAlgorithmPtrOutput
	ToEncryptionAlgorithmPtrOutputWithContext(context.Context) EncryptionAlgorithmPtrOutput
}

type encryptionAlgorithmPtr string

func EncryptionAlgorithmPtr(v string) EncryptionAlgorithmPtrInput {
	return (*encryptionAlgorithmPtr)(&v)
}

func (*encryptionAlgorithmPtr) ElementType() reflect.Type {
	return encryptionAlgorithmPtrType
}

func (in *encryptionAlgorithmPtr) ToEncryptionAlgorithmPtrOutput() EncryptionAlgorithmPtrOutput {
	return pulumi.ToOutput(in).(EncryptionAlgorithmPtrOutput)
}

func (in *encryptionAlgorithmPtr) ToEncryptionAlgorithmPtrOutputWithContext(ctx context.Context) EncryptionAlgorithmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EncryptionAlgorithmPtrOutput)
}

type HashingAlgorithm string

const (
	HashingAlgorithmNotSpecified = HashingAlgorithm("NotSpecified")
	HashingAlgorithmNone         = HashingAlgorithm("None")
	HashingAlgorithmSHA2256      = HashingAlgorithm("SHA2256")
	HashingAlgorithmSHA2384      = HashingAlgorithm("SHA2384")
	HashingAlgorithmSHA2512      = HashingAlgorithm("SHA2512")
)

func (HashingAlgorithm) ElementType() reflect.Type {
	return reflect.TypeOf((*HashingAlgorithm)(nil)).Elem()
}

func (e HashingAlgorithm) ToHashingAlgorithmOutput() HashingAlgorithmOutput {
	return pulumi.ToOutput(e).(HashingAlgorithmOutput)
}

func (e HashingAlgorithm) ToHashingAlgorithmOutputWithContext(ctx context.Context) HashingAlgorithmOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HashingAlgorithmOutput)
}

func (e HashingAlgorithm) ToHashingAlgorithmPtrOutput() HashingAlgorithmPtrOutput {
	return e.ToHashingAlgorithmPtrOutputWithContext(context.Background())
}

func (e HashingAlgorithm) ToHashingAlgorithmPtrOutputWithContext(ctx context.Context) HashingAlgorithmPtrOutput {
	return HashingAlgorithm(e).ToHashingAlgorithmOutputWithContext(ctx).ToHashingAlgorithmPtrOutputWithContext(ctx)
}

func (e HashingAlgorithm) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HashingAlgorithm) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HashingAlgorithm) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HashingAlgorithm) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HashingAlgorithmOutput struct{ *pulumi.OutputState }

func (HashingAlgorithmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HashingAlgorithm)(nil)).Elem()
}

func (o HashingAlgorithmOutput) ToHashingAlgorithmOutput() HashingAlgorithmOutput {
	return o
}

func (o HashingAlgorithmOutput) ToHashingAlgorithmOutputWithContext(ctx context.Context) HashingAlgorithmOutput {
	return o
}

func (o HashingAlgorithmOutput) ToHashingAlgorithmPtrOutput() HashingAlgorithmPtrOutput {
	return o.ToHashingAlgorithmPtrOutputWithContext(context.Background())
}

func (o HashingAlgorithmOutput) ToHashingAlgorithmPtrOutputWithContext(ctx context.Context) HashingAlgorithmPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HashingAlgorithm) *HashingAlgorithm {
		return &v
	}).(HashingAlgorithmPtrOutput)
}

func (o HashingAlgorithmOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HashingAlgorithmOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HashingAlgorithm) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HashingAlgorithmOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HashingAlgorithmOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HashingAlgorithm) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HashingAlgorithmPtrOutput struct{ *pulumi.OutputState }

func (HashingAlgorithmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HashingAlgorithm)(nil)).Elem()
}

func (o HashingAlgorithmPtrOutput) ToHashingAlgorithmPtrOutput() HashingAlgorithmPtrOutput {
	return o
}

func (o HashingAlgorithmPtrOutput) ToHashingAlgorithmPtrOutputWithContext(ctx context.Context) HashingAlgorithmPtrOutput {
	return o
}

func (o HashingAlgorithmPtrOutput) Elem() HashingAlgorithmOutput {
	return o.ApplyT(func(v *HashingAlgorithm) HashingAlgorithm {
		if v != nil {
			return *v
		}
		var ret HashingAlgorithm
		return ret
	}).(HashingAlgorithmOutput)
}

func (o HashingAlgorithmPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HashingAlgorithmPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HashingAlgorithm) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HashingAlgorithmInput interface {
	pulumi.Input

	ToHashingAlgorithmOutput() HashingAlgorithmOutput
	ToHashingAlgorithmOutputWithContext(context.Context) HashingAlgorithmOutput
}

var hashingAlgorithmPtrType = reflect.TypeOf((**HashingAlgorithm)(nil)).Elem()

type HashingAlgorithmPtrInput interface {
	pulumi.Input

	ToHashingAlgorithmPtrOutput() HashingAlgorithmPtrOutput
	ToHashingAlgorithmPtrOutputWithContext(context.Context) HashingAlgorithmPtrOutput
}

type hashingAlgorithmPtr string

func HashingAlgorithmPtr(v string) HashingAlgorithmPtrInput {
	return (*hashingAlgorithmPtr)(&v)
}

func (*hashingAlgorithmPtr) ElementType() reflect.Type {
	return hashingAlgorithmPtrType
}

func (in *hashingAlgorithmPtr) ToHashingAlgorithmPtrOutput() HashingAlgorithmPtrOutput {
	return pulumi.ToOutput(in).(HashingAlgorithmPtrOutput)
}

func (in *hashingAlgorithmPtr) ToHashingAlgorithmPtrOutputWithContext(ctx context.Context) HashingAlgorithmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HashingAlgorithmPtrOutput)
}

type MapType string

const (
	MapTypeNotSpecified = MapType("NotSpecified")
	MapTypeXslt         = MapType("Xslt")
)

func (MapType) ElementType() reflect.Type {
	return reflect.TypeOf((*MapType)(nil)).Elem()
}

func (e MapType) ToMapTypeOutput() MapTypeOutput {
	return pulumi.ToOutput(e).(MapTypeOutput)
}

func (e MapType) ToMapTypeOutputWithContext(ctx context.Context) MapTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MapTypeOutput)
}

func (e MapType) ToMapTypePtrOutput() MapTypePtrOutput {
	return e.ToMapTypePtrOutputWithContext(context.Background())
}

func (e MapType) ToMapTypePtrOutputWithContext(ctx context.Context) MapTypePtrOutput {
	return MapType(e).ToMapTypeOutputWithContext(ctx).ToMapTypePtrOutputWithContext(ctx)
}

func (e MapType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MapType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MapType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MapType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MapTypeOutput struct{ *pulumi.OutputState }

func (MapTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MapType)(nil)).Elem()
}

func (o MapTypeOutput) ToMapTypeOutput() MapTypeOutput {
	return o
}

func (o MapTypeOutput) ToMapTypeOutputWithContext(ctx context.Context) MapTypeOutput {
	return o
}

func (o MapTypeOutput) ToMapTypePtrOutput() MapTypePtrOutput {
	return o.ToMapTypePtrOutputWithContext(context.Background())
}

func (o MapTypeOutput) ToMapTypePtrOutputWithContext(ctx context.Context) MapTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MapType) *MapType {
		return &v
	}).(MapTypePtrOutput)
}

func (o MapTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MapTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MapType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MapTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MapTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MapType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MapTypePtrOutput struct{ *pulumi.OutputState }

func (MapTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MapType)(nil)).Elem()
}

func (o MapTypePtrOutput) ToMapTypePtrOutput() MapTypePtrOutput {
	return o
}

func (o MapTypePtrOutput) ToMapTypePtrOutputWithContext(ctx context.Context) MapTypePtrOutput {
	return o
}

func (o MapTypePtrOutput) Elem() MapTypeOutput {
	return o.ApplyT(func(v *MapType) MapType {
		if v != nil {
			return *v
		}
		var ret MapType
		return ret
	}).(MapTypeOutput)
}

func (o MapTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MapTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MapType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MapTypeInput interface {
	pulumi.Input

	ToMapTypeOutput() MapTypeOutput
	ToMapTypeOutputWithContext(context.Context) MapTypeOutput
}

var mapTypePtrType = reflect.TypeOf((**MapType)(nil)).Elem()

type MapTypePtrInput interface {
	pulumi.Input

	ToMapTypePtrOutput() MapTypePtrOutput
	ToMapTypePtrOutputWithContext(context.Context) MapTypePtrOutput
}

type mapTypePtr string

func MapTypePtr(v string) MapTypePtrInput {
	return (*mapTypePtr)(&v)
}

func (*mapTypePtr) ElementType() reflect.Type {
	return mapTypePtrType
}

func (in *mapTypePtr) ToMapTypePtrOutput() MapTypePtrOutput {
	return pulumi.ToOutput(in).(MapTypePtrOutput)
}

func (in *mapTypePtr) ToMapTypePtrOutputWithContext(ctx context.Context) MapTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MapTypePtrOutput)
}

type MessageFilterType string

const (
	MessageFilterTypeNotSpecified = MessageFilterType("NotSpecified")
	MessageFilterTypeInclude      = MessageFilterType("Include")
	MessageFilterTypeExclude      = MessageFilterType("Exclude")
)

func (MessageFilterType) ElementType() reflect.Type {
	return reflect.TypeOf((*MessageFilterType)(nil)).Elem()
}

func (e MessageFilterType) ToMessageFilterTypeOutput() MessageFilterTypeOutput {
	return pulumi.ToOutput(e).(MessageFilterTypeOutput)
}

func (e MessageFilterType) ToMessageFilterTypeOutputWithContext(ctx context.Context) MessageFilterTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MessageFilterTypeOutput)
}

func (e MessageFilterType) ToMessageFilterTypePtrOutput() MessageFilterTypePtrOutput {
	return e.ToMessageFilterTypePtrOutputWithContext(context.Background())
}

func (e MessageFilterType) ToMessageFilterTypePtrOutputWithContext(ctx context.Context) MessageFilterTypePtrOutput {
	return MessageFilterType(e).ToMessageFilterTypeOutputWithContext(ctx).ToMessageFilterTypePtrOutputWithContext(ctx)
}

func (e MessageFilterType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MessageFilterType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MessageFilterType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MessageFilterType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MessageFilterTypeOutput struct{ *pulumi.OutputState }

func (MessageFilterTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MessageFilterType)(nil)).Elem()
}

func (o MessageFilterTypeOutput) ToMessageFilterTypeOutput() MessageFilterTypeOutput {
	return o
}

func (o MessageFilterTypeOutput) ToMessageFilterTypeOutputWithContext(ctx context.Context) MessageFilterTypeOutput {
	return o
}

func (o MessageFilterTypeOutput) ToMessageFilterTypePtrOutput() MessageFilterTypePtrOutput {
	return o.ToMessageFilterTypePtrOutputWithContext(context.Background())
}

func (o MessageFilterTypeOutput) ToMessageFilterTypePtrOutputWithContext(ctx context.Context) MessageFilterTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MessageFilterType) *MessageFilterType {
		return &v
	}).(MessageFilterTypePtrOutput)
}

func (o MessageFilterTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MessageFilterTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MessageFilterType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MessageFilterTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MessageFilterTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MessageFilterType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MessageFilterTypePtrOutput struct{ *pulumi.OutputState }

func (MessageFilterTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MessageFilterType)(nil)).Elem()
}

func (o MessageFilterTypePtrOutput) ToMessageFilterTypePtrOutput() MessageFilterTypePtrOutput {
	return o
}

func (o MessageFilterTypePtrOutput) ToMessageFilterTypePtrOutputWithContext(ctx context.Context) MessageFilterTypePtrOutput {
	return o
}

func (o MessageFilterTypePtrOutput) Elem() MessageFilterTypeOutput {
	return o.ApplyT(func(v *MessageFilterType) MessageFilterType {
		if v != nil {
			return *v
		}
		var ret MessageFilterType
		return ret
	}).(MessageFilterTypeOutput)
}

func (o MessageFilterTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MessageFilterTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MessageFilterType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MessageFilterTypeInput interface {
	pulumi.Input

	ToMessageFilterTypeOutput() MessageFilterTypeOutput
	ToMessageFilterTypeOutputWithContext(context.Context) MessageFilterTypeOutput
}

var messageFilterTypePtrType = reflect.TypeOf((**MessageFilterType)(nil)).Elem()

type MessageFilterTypePtrInput interface {
	pulumi.Input

	ToMessageFilterTypePtrOutput() MessageFilterTypePtrOutput
	ToMessageFilterTypePtrOutputWithContext(context.Context) MessageFilterTypePtrOutput
}

type messageFilterTypePtr string

func MessageFilterTypePtr(v string) MessageFilterTypePtrInput {
	return (*messageFilterTypePtr)(&v)
}

func (*messageFilterTypePtr) ElementType() reflect.Type {
	return messageFilterTypePtrType
}

func (in *messageFilterTypePtr) ToMessageFilterTypePtrOutput() MessageFilterTypePtrOutput {
	return pulumi.ToOutput(in).(MessageFilterTypePtrOutput)
}

func (in *messageFilterTypePtr) ToMessageFilterTypePtrOutputWithContext(ctx context.Context) MessageFilterTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MessageFilterTypePtrOutput)
}

type PartnerType string

const (
	PartnerTypeNotSpecified = PartnerType("NotSpecified")
	PartnerTypeB2B          = PartnerType("B2B")
)

func (PartnerType) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerType)(nil)).Elem()
}

func (e PartnerType) ToPartnerTypeOutput() PartnerTypeOutput {
	return pulumi.ToOutput(e).(PartnerTypeOutput)
}

func (e PartnerType) ToPartnerTypeOutputWithContext(ctx context.Context) PartnerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PartnerTypeOutput)
}

func (e PartnerType) ToPartnerTypePtrOutput() PartnerTypePtrOutput {
	return e.ToPartnerTypePtrOutputWithContext(context.Background())
}

func (e PartnerType) ToPartnerTypePtrOutputWithContext(ctx context.Context) PartnerTypePtrOutput {
	return PartnerType(e).ToPartnerTypeOutputWithContext(ctx).ToPartnerTypePtrOutputWithContext(ctx)
}

func (e PartnerType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PartnerType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PartnerType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PartnerType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PartnerTypeOutput struct{ *pulumi.OutputState }

func (PartnerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerType)(nil)).Elem()
}

func (o PartnerTypeOutput) ToPartnerTypeOutput() PartnerTypeOutput {
	return o
}

func (o PartnerTypeOutput) ToPartnerTypeOutputWithContext(ctx context.Context) PartnerTypeOutput {
	return o
}

func (o PartnerTypeOutput) ToPartnerTypePtrOutput() PartnerTypePtrOutput {
	return o.ToPartnerTypePtrOutputWithContext(context.Background())
}

func (o PartnerTypeOutput) ToPartnerTypePtrOutputWithContext(ctx context.Context) PartnerTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PartnerType) *PartnerType {
		return &v
	}).(PartnerTypePtrOutput)
}

func (o PartnerTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PartnerTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PartnerType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PartnerTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PartnerTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PartnerType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PartnerTypePtrOutput struct{ *pulumi.OutputState }

func (PartnerTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerType)(nil)).Elem()
}

func (o PartnerTypePtrOutput) ToPartnerTypePtrOutput() PartnerTypePtrOutput {
	return o
}

func (o PartnerTypePtrOutput) ToPartnerTypePtrOutputWithContext(ctx context.Context) PartnerTypePtrOutput {
	return o
}

func (o PartnerTypePtrOutput) Elem() PartnerTypeOutput {
	return o.ApplyT(func(v *PartnerType) PartnerType {
		if v != nil {
			return *v
		}
		var ret PartnerType
		return ret
	}).(PartnerTypeOutput)
}

func (o PartnerTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PartnerTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PartnerType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PartnerTypeInput interface {
	pulumi.Input

	ToPartnerTypeOutput() PartnerTypeOutput
	ToPartnerTypeOutputWithContext(context.Context) PartnerTypeOutput
}

var partnerTypePtrType = reflect.TypeOf((**PartnerType)(nil)).Elem()

type PartnerTypePtrInput interface {
	pulumi.Input

	ToPartnerTypePtrOutput() PartnerTypePtrOutput
	ToPartnerTypePtrOutputWithContext(context.Context) PartnerTypePtrOutput
}

type partnerTypePtr string

func PartnerTypePtr(v string) PartnerTypePtrInput {
	return (*partnerTypePtr)(&v)
}

func (*partnerTypePtr) ElementType() reflect.Type {
	return partnerTypePtrType
}

func (in *partnerTypePtr) ToPartnerTypePtrOutput() PartnerTypePtrOutput {
	return pulumi.ToOutput(in).(PartnerTypePtrOutput)
}

func (in *partnerTypePtr) ToPartnerTypePtrOutputWithContext(ctx context.Context) PartnerTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PartnerTypePtrOutput)
}

type SchemaType string

const (
	SchemaTypeNotSpecified = SchemaType("NotSpecified")
	SchemaTypeXml          = SchemaType("Xml")
)

func (SchemaType) ElementType() reflect.Type {
	return reflect.TypeOf((*SchemaType)(nil)).Elem()
}

func (e SchemaType) ToSchemaTypeOutput() SchemaTypeOutput {
	return pulumi.ToOutput(e).(SchemaTypeOutput)
}

func (e SchemaType) ToSchemaTypeOutputWithContext(ctx context.Context) SchemaTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SchemaTypeOutput)
}

func (e SchemaType) ToSchemaTypePtrOutput() SchemaTypePtrOutput {
	return e.ToSchemaTypePtrOutputWithContext(context.Background())
}

func (e SchemaType) ToSchemaTypePtrOutputWithContext(ctx context.Context) SchemaTypePtrOutput {
	return SchemaType(e).ToSchemaTypeOutputWithContext(ctx).ToSchemaTypePtrOutputWithContext(ctx)
}

func (e SchemaType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SchemaType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SchemaType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SchemaType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SchemaTypeOutput struct{ *pulumi.OutputState }

func (SchemaTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SchemaType)(nil)).Elem()
}

func (o SchemaTypeOutput) ToSchemaTypeOutput() SchemaTypeOutput {
	return o
}

func (o SchemaTypeOutput) ToSchemaTypeOutputWithContext(ctx context.Context) SchemaTypeOutput {
	return o
}

func (o SchemaTypeOutput) ToSchemaTypePtrOutput() SchemaTypePtrOutput {
	return o.ToSchemaTypePtrOutputWithContext(context.Background())
}

func (o SchemaTypeOutput) ToSchemaTypePtrOutputWithContext(ctx context.Context) SchemaTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SchemaType) *SchemaType {
		return &v
	}).(SchemaTypePtrOutput)
}

func (o SchemaTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SchemaTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SchemaType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SchemaTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SchemaTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SchemaType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SchemaTypePtrOutput struct{ *pulumi.OutputState }

func (SchemaTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SchemaType)(nil)).Elem()
}

func (o SchemaTypePtrOutput) ToSchemaTypePtrOutput() SchemaTypePtrOutput {
	return o
}

func (o SchemaTypePtrOutput) ToSchemaTypePtrOutputWithContext(ctx context.Context) SchemaTypePtrOutput {
	return o
}

func (o SchemaTypePtrOutput) Elem() SchemaTypeOutput {
	return o.ApplyT(func(v *SchemaType) SchemaType {
		if v != nil {
			return *v
		}
		var ret SchemaType
		return ret
	}).(SchemaTypeOutput)
}

func (o SchemaTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SchemaTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SchemaType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SchemaTypeInput interface {
	pulumi.Input

	ToSchemaTypeOutput() SchemaTypeOutput
	ToSchemaTypeOutputWithContext(context.Context) SchemaTypeOutput
}

var schemaTypePtrType = reflect.TypeOf((**SchemaType)(nil)).Elem()

type SchemaTypePtrInput interface {
	pulumi.Input

	ToSchemaTypePtrOutput() SchemaTypePtrOutput
	ToSchemaTypePtrOutputWithContext(context.Context) SchemaTypePtrOutput
}

type schemaTypePtr string

func SchemaTypePtr(v string) SchemaTypePtrInput {
	return (*schemaTypePtr)(&v)
}

func (*schemaTypePtr) ElementType() reflect.Type {
	return schemaTypePtrType
}

func (in *schemaTypePtr) ToSchemaTypePtrOutput() SchemaTypePtrOutput {
	return pulumi.ToOutput(in).(SchemaTypePtrOutput)
}

func (in *schemaTypePtr) ToSchemaTypePtrOutputWithContext(ctx context.Context) SchemaTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SchemaTypePtrOutput)
}

type SegmentTerminatorSuffix string

const (
	SegmentTerminatorSuffixNotSpecified = SegmentTerminatorSuffix("NotSpecified")
	SegmentTerminatorSuffixNone         = SegmentTerminatorSuffix("None")
	SegmentTerminatorSuffixCR           = SegmentTerminatorSuffix("CR")
	SegmentTerminatorSuffixLF           = SegmentTerminatorSuffix("LF")
	SegmentTerminatorSuffixCRLF         = SegmentTerminatorSuffix("CRLF")
)

func (SegmentTerminatorSuffix) ElementType() reflect.Type {
	return reflect.TypeOf((*SegmentTerminatorSuffix)(nil)).Elem()
}

func (e SegmentTerminatorSuffix) ToSegmentTerminatorSuffixOutput() SegmentTerminatorSuffixOutput {
	return pulumi.ToOutput(e).(SegmentTerminatorSuffixOutput)
}

func (e SegmentTerminatorSuffix) ToSegmentTerminatorSuffixOutputWithContext(ctx context.Context) SegmentTerminatorSuffixOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SegmentTerminatorSuffixOutput)
}

func (e SegmentTerminatorSuffix) ToSegmentTerminatorSuffixPtrOutput() SegmentTerminatorSuffixPtrOutput {
	return e.ToSegmentTerminatorSuffixPtrOutputWithContext(context.Background())
}

func (e SegmentTerminatorSuffix) ToSegmentTerminatorSuffixPtrOutputWithContext(ctx context.Context) SegmentTerminatorSuffixPtrOutput {
	return SegmentTerminatorSuffix(e).ToSegmentTerminatorSuffixOutputWithContext(ctx).ToSegmentTerminatorSuffixPtrOutputWithContext(ctx)
}

func (e SegmentTerminatorSuffix) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SegmentTerminatorSuffix) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SegmentTerminatorSuffix) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SegmentTerminatorSuffix) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SegmentTerminatorSuffixOutput struct{ *pulumi.OutputState }

func (SegmentTerminatorSuffixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SegmentTerminatorSuffix)(nil)).Elem()
}

func (o SegmentTerminatorSuffixOutput) ToSegmentTerminatorSuffixOutput() SegmentTerminatorSuffixOutput {
	return o
}

func (o SegmentTerminatorSuffixOutput) ToSegmentTerminatorSuffixOutputWithContext(ctx context.Context) SegmentTerminatorSuffixOutput {
	return o
}

func (o SegmentTerminatorSuffixOutput) ToSegmentTerminatorSuffixPtrOutput() SegmentTerminatorSuffixPtrOutput {
	return o.ToSegmentTerminatorSuffixPtrOutputWithContext(context.Background())
}

func (o SegmentTerminatorSuffixOutput) ToSegmentTerminatorSuffixPtrOutputWithContext(ctx context.Context) SegmentTerminatorSuffixPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SegmentTerminatorSuffix) *SegmentTerminatorSuffix {
		return &v
	}).(SegmentTerminatorSuffixPtrOutput)
}

func (o SegmentTerminatorSuffixOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SegmentTerminatorSuffixOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SegmentTerminatorSuffix) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SegmentTerminatorSuffixOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SegmentTerminatorSuffixOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SegmentTerminatorSuffix) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SegmentTerminatorSuffixPtrOutput struct{ *pulumi.OutputState }

func (SegmentTerminatorSuffixPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SegmentTerminatorSuffix)(nil)).Elem()
}

func (o SegmentTerminatorSuffixPtrOutput) ToSegmentTerminatorSuffixPtrOutput() SegmentTerminatorSuffixPtrOutput {
	return o
}

func (o SegmentTerminatorSuffixPtrOutput) ToSegmentTerminatorSuffixPtrOutputWithContext(ctx context.Context) SegmentTerminatorSuffixPtrOutput {
	return o
}

func (o SegmentTerminatorSuffixPtrOutput) Elem() SegmentTerminatorSuffixOutput {
	return o.ApplyT(func(v *SegmentTerminatorSuffix) SegmentTerminatorSuffix {
		if v != nil {
			return *v
		}
		var ret SegmentTerminatorSuffix
		return ret
	}).(SegmentTerminatorSuffixOutput)
}

func (o SegmentTerminatorSuffixPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SegmentTerminatorSuffixPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SegmentTerminatorSuffix) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SegmentTerminatorSuffixInput interface {
	pulumi.Input

	ToSegmentTerminatorSuffixOutput() SegmentTerminatorSuffixOutput
	ToSegmentTerminatorSuffixOutputWithContext(context.Context) SegmentTerminatorSuffixOutput
}

var segmentTerminatorSuffixPtrType = reflect.TypeOf((**SegmentTerminatorSuffix)(nil)).Elem()

type SegmentTerminatorSuffixPtrInput interface {
	pulumi.Input

	ToSegmentTerminatorSuffixPtrOutput() SegmentTerminatorSuffixPtrOutput
	ToSegmentTerminatorSuffixPtrOutputWithContext(context.Context) SegmentTerminatorSuffixPtrOutput
}

type segmentTerminatorSuffixPtr string

func SegmentTerminatorSuffixPtr(v string) SegmentTerminatorSuffixPtrInput {
	return (*segmentTerminatorSuffixPtr)(&v)
}

func (*segmentTerminatorSuffixPtr) ElementType() reflect.Type {
	return segmentTerminatorSuffixPtrType
}

func (in *segmentTerminatorSuffixPtr) ToSegmentTerminatorSuffixPtrOutput() SegmentTerminatorSuffixPtrOutput {
	return pulumi.ToOutput(in).(SegmentTerminatorSuffixPtrOutput)
}

func (in *segmentTerminatorSuffixPtr) ToSegmentTerminatorSuffixPtrOutputWithContext(ctx context.Context) SegmentTerminatorSuffixPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SegmentTerminatorSuffixPtrOutput)
}

type SkuName string

const (
	SkuNameNotSpecified = SkuName("NotSpecified")
	SkuNameFree         = SkuName("Free")
	SkuNameShared       = SkuName("Shared")
	SkuNameBasic        = SkuName("Basic")
	SkuNameStandard     = SkuName("Standard")
	SkuNamePremium      = SkuName("Premium")
)

func (SkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (e SkuName) ToSkuNameOutput() SkuNameOutput {
	return pulumi.ToOutput(e).(SkuNameOutput)
}

func (e SkuName) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuNameOutput)
}

func (e SkuName) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return e.ToSkuNamePtrOutputWithContext(context.Background())
}

func (e SkuName) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return SkuName(e).ToSkuNameOutputWithContext(ctx).ToSkuNamePtrOutputWithContext(ctx)
}

func (e SkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuNameOutput struct{ *pulumi.OutputState }

func (SkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (o SkuNameOutput) ToSkuNameOutput() SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o.ToSkuNamePtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuName) *SkuName {
		return &v
	}).(SkuNamePtrOutput)
}

func (o SkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuNamePtrOutput struct{ *pulumi.OutputState }

func (SkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuName)(nil)).Elem()
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) Elem() SkuNameOutput {
	return o.ApplyT(func(v *SkuName) SkuName {
		if v != nil {
			return *v
		}
		var ret SkuName
		return ret
	}).(SkuNameOutput)
}

func (o SkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuNameInput interface {
	pulumi.Input

	ToSkuNameOutput() SkuNameOutput
	ToSkuNameOutputWithContext(context.Context) SkuNameOutput
}

var skuNamePtrType = reflect.TypeOf((**SkuName)(nil)).Elem()

type SkuNamePtrInput interface {
	pulumi.Input

	ToSkuNamePtrOutput() SkuNamePtrOutput
	ToSkuNamePtrOutputWithContext(context.Context) SkuNamePtrOutput
}

type skuNamePtr string

func SkuNamePtr(v string) SkuNamePtrInput {
	return (*skuNamePtr)(&v)
}

func (*skuNamePtr) ElementType() reflect.Type {
	return skuNamePtrType
}

func (in *skuNamePtr) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return pulumi.ToOutput(in).(SkuNamePtrOutput)
}

func (in *skuNamePtr) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuNamePtrOutput)
}

type TrailingSeparatorPolicy string

const (
	TrailingSeparatorPolicyNotSpecified = TrailingSeparatorPolicy("NotSpecified")
	TrailingSeparatorPolicyNotAllowed   = TrailingSeparatorPolicy("NotAllowed")
	TrailingSeparatorPolicyOptional     = TrailingSeparatorPolicy("Optional")
	TrailingSeparatorPolicyMandatory    = TrailingSeparatorPolicy("Mandatory")
)

func (TrailingSeparatorPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*TrailingSeparatorPolicy)(nil)).Elem()
}

func (e TrailingSeparatorPolicy) ToTrailingSeparatorPolicyOutput() TrailingSeparatorPolicyOutput {
	return pulumi.ToOutput(e).(TrailingSeparatorPolicyOutput)
}

func (e TrailingSeparatorPolicy) ToTrailingSeparatorPolicyOutputWithContext(ctx context.Context) TrailingSeparatorPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TrailingSeparatorPolicyOutput)
}

func (e TrailingSeparatorPolicy) ToTrailingSeparatorPolicyPtrOutput() TrailingSeparatorPolicyPtrOutput {
	return e.ToTrailingSeparatorPolicyPtrOutputWithContext(context.Background())
}

func (e TrailingSeparatorPolicy) ToTrailingSeparatorPolicyPtrOutputWithContext(ctx context.Context) TrailingSeparatorPolicyPtrOutput {
	return TrailingSeparatorPolicy(e).ToTrailingSeparatorPolicyOutputWithContext(ctx).ToTrailingSeparatorPolicyPtrOutputWithContext(ctx)
}

func (e TrailingSeparatorPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TrailingSeparatorPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TrailingSeparatorPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TrailingSeparatorPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TrailingSeparatorPolicyOutput struct{ *pulumi.OutputState }

func (TrailingSeparatorPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrailingSeparatorPolicy)(nil)).Elem()
}

func (o TrailingSeparatorPolicyOutput) ToTrailingSeparatorPolicyOutput() TrailingSeparatorPolicyOutput {
	return o
}

func (o TrailingSeparatorPolicyOutput) ToTrailingSeparatorPolicyOutputWithContext(ctx context.Context) TrailingSeparatorPolicyOutput {
	return o
}

func (o TrailingSeparatorPolicyOutput) ToTrailingSeparatorPolicyPtrOutput() TrailingSeparatorPolicyPtrOutput {
	return o.ToTrailingSeparatorPolicyPtrOutputWithContext(context.Background())
}

func (o TrailingSeparatorPolicyOutput) ToTrailingSeparatorPolicyPtrOutputWithContext(ctx context.Context) TrailingSeparatorPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TrailingSeparatorPolicy) *TrailingSeparatorPolicy {
		return &v
	}).(TrailingSeparatorPolicyPtrOutput)
}

func (o TrailingSeparatorPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TrailingSeparatorPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TrailingSeparatorPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TrailingSeparatorPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TrailingSeparatorPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TrailingSeparatorPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TrailingSeparatorPolicyPtrOutput struct{ *pulumi.OutputState }

func (TrailingSeparatorPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrailingSeparatorPolicy)(nil)).Elem()
}

func (o TrailingSeparatorPolicyPtrOutput) ToTrailingSeparatorPolicyPtrOutput() TrailingSeparatorPolicyPtrOutput {
	return o
}

func (o TrailingSeparatorPolicyPtrOutput) ToTrailingSeparatorPolicyPtrOutputWithContext(ctx context.Context) TrailingSeparatorPolicyPtrOutput {
	return o
}

func (o TrailingSeparatorPolicyPtrOutput) Elem() TrailingSeparatorPolicyOutput {
	return o.ApplyT(func(v *TrailingSeparatorPolicy) TrailingSeparatorPolicy {
		if v != nil {
			return *v
		}
		var ret TrailingSeparatorPolicy
		return ret
	}).(TrailingSeparatorPolicyOutput)
}

func (o TrailingSeparatorPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TrailingSeparatorPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TrailingSeparatorPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TrailingSeparatorPolicyInput interface {
	pulumi.Input

	ToTrailingSeparatorPolicyOutput() TrailingSeparatorPolicyOutput
	ToTrailingSeparatorPolicyOutputWithContext(context.Context) TrailingSeparatorPolicyOutput
}

var trailingSeparatorPolicyPtrType = reflect.TypeOf((**TrailingSeparatorPolicy)(nil)).Elem()

type TrailingSeparatorPolicyPtrInput interface {
	pulumi.Input

	ToTrailingSeparatorPolicyPtrOutput() TrailingSeparatorPolicyPtrOutput
	ToTrailingSeparatorPolicyPtrOutputWithContext(context.Context) TrailingSeparatorPolicyPtrOutput
}

type trailingSeparatorPolicyPtr string

func TrailingSeparatorPolicyPtr(v string) TrailingSeparatorPolicyPtrInput {
	return (*trailingSeparatorPolicyPtr)(&v)
}

func (*trailingSeparatorPolicyPtr) ElementType() reflect.Type {
	return trailingSeparatorPolicyPtrType
}

func (in *trailingSeparatorPolicyPtr) ToTrailingSeparatorPolicyPtrOutput() TrailingSeparatorPolicyPtrOutput {
	return pulumi.ToOutput(in).(TrailingSeparatorPolicyPtrOutput)
}

func (in *trailingSeparatorPolicyPtr) ToTrailingSeparatorPolicyPtrOutputWithContext(ctx context.Context) TrailingSeparatorPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TrailingSeparatorPolicyPtrOutput)
}

type UsageIndicator string

const (
	UsageIndicatorNotSpecified = UsageIndicator("NotSpecified")
	UsageIndicatorTest         = UsageIndicator("Test")
	UsageIndicatorInformation  = UsageIndicator("Information")
	UsageIndicatorProduction   = UsageIndicator("Production")
)

func (UsageIndicator) ElementType() reflect.Type {
	return reflect.TypeOf((*UsageIndicator)(nil)).Elem()
}

func (e UsageIndicator) ToUsageIndicatorOutput() UsageIndicatorOutput {
	return pulumi.ToOutput(e).(UsageIndicatorOutput)
}

func (e UsageIndicator) ToUsageIndicatorOutputWithContext(ctx context.Context) UsageIndicatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UsageIndicatorOutput)
}

func (e UsageIndicator) ToUsageIndicatorPtrOutput() UsageIndicatorPtrOutput {
	return e.ToUsageIndicatorPtrOutputWithContext(context.Background())
}

func (e UsageIndicator) ToUsageIndicatorPtrOutputWithContext(ctx context.Context) UsageIndicatorPtrOutput {
	return UsageIndicator(e).ToUsageIndicatorOutputWithContext(ctx).ToUsageIndicatorPtrOutputWithContext(ctx)
}

func (e UsageIndicator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UsageIndicator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UsageIndicator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UsageIndicator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UsageIndicatorOutput struct{ *pulumi.OutputState }

func (UsageIndicatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UsageIndicator)(nil)).Elem()
}

func (o UsageIndicatorOutput) ToUsageIndicatorOutput() UsageIndicatorOutput {
	return o
}

func (o UsageIndicatorOutput) ToUsageIndicatorOutputWithContext(ctx context.Context) UsageIndicatorOutput {
	return o
}

func (o UsageIndicatorOutput) ToUsageIndicatorPtrOutput() UsageIndicatorPtrOutput {
	return o.ToUsageIndicatorPtrOutputWithContext(context.Background())
}

func (o UsageIndicatorOutput) ToUsageIndicatorPtrOutputWithContext(ctx context.Context) UsageIndicatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UsageIndicator) *UsageIndicator {
		return &v
	}).(UsageIndicatorPtrOutput)
}

func (o UsageIndicatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UsageIndicatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UsageIndicator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UsageIndicatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UsageIndicatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UsageIndicator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UsageIndicatorPtrOutput struct{ *pulumi.OutputState }

func (UsageIndicatorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UsageIndicator)(nil)).Elem()
}

func (o UsageIndicatorPtrOutput) ToUsageIndicatorPtrOutput() UsageIndicatorPtrOutput {
	return o
}

func (o UsageIndicatorPtrOutput) ToUsageIndicatorPtrOutputWithContext(ctx context.Context) UsageIndicatorPtrOutput {
	return o
}

func (o UsageIndicatorPtrOutput) Elem() UsageIndicatorOutput {
	return o.ApplyT(func(v *UsageIndicator) UsageIndicator {
		if v != nil {
			return *v
		}
		var ret UsageIndicator
		return ret
	}).(UsageIndicatorOutput)
}

func (o UsageIndicatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UsageIndicatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UsageIndicator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type UsageIndicatorInput interface {
	pulumi.Input

	ToUsageIndicatorOutput() UsageIndicatorOutput
	ToUsageIndicatorOutputWithContext(context.Context) UsageIndicatorOutput
}

var usageIndicatorPtrType = reflect.TypeOf((**UsageIndicator)(nil)).Elem()

type UsageIndicatorPtrInput interface {
	pulumi.Input

	ToUsageIndicatorPtrOutput() UsageIndicatorPtrOutput
	ToUsageIndicatorPtrOutputWithContext(context.Context) UsageIndicatorPtrOutput
}

type usageIndicatorPtr string

func UsageIndicatorPtr(v string) UsageIndicatorPtrInput {
	return (*usageIndicatorPtr)(&v)
}

func (*usageIndicatorPtr) ElementType() reflect.Type {
	return usageIndicatorPtrType
}

func (in *usageIndicatorPtr) ToUsageIndicatorPtrOutput() UsageIndicatorPtrOutput {
	return pulumi.ToOutput(in).(UsageIndicatorPtrOutput)
}

func (in *usageIndicatorPtr) ToUsageIndicatorPtrOutputWithContext(ctx context.Context) UsageIndicatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UsageIndicatorPtrOutput)
}

type X12CharacterSet string

const (
	X12CharacterSetNotSpecified = X12CharacterSet("NotSpecified")
	X12CharacterSetBasic        = X12CharacterSet("Basic")
	X12CharacterSetExtended     = X12CharacterSet("Extended")
	X12CharacterSetUTF8         = X12CharacterSet("UTF8")
)

func (X12CharacterSet) ElementType() reflect.Type {
	return reflect.TypeOf((*X12CharacterSet)(nil)).Elem()
}

func (e X12CharacterSet) ToX12CharacterSetOutput() X12CharacterSetOutput {
	return pulumi.ToOutput(e).(X12CharacterSetOutput)
}

func (e X12CharacterSet) ToX12CharacterSetOutputWithContext(ctx context.Context) X12CharacterSetOutput {
	return pulumi.ToOutputWithContext(ctx, e).(X12CharacterSetOutput)
}

func (e X12CharacterSet) ToX12CharacterSetPtrOutput() X12CharacterSetPtrOutput {
	return e.ToX12CharacterSetPtrOutputWithContext(context.Background())
}

func (e X12CharacterSet) ToX12CharacterSetPtrOutputWithContext(ctx context.Context) X12CharacterSetPtrOutput {
	return X12CharacterSet(e).ToX12CharacterSetOutputWithContext(ctx).ToX12CharacterSetPtrOutputWithContext(ctx)
}

func (e X12CharacterSet) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e X12CharacterSet) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e X12CharacterSet) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e X12CharacterSet) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type X12CharacterSetOutput struct{ *pulumi.OutputState }

func (X12CharacterSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12CharacterSet)(nil)).Elem()
}

func (o X12CharacterSetOutput) ToX12CharacterSetOutput() X12CharacterSetOutput {
	return o
}

func (o X12CharacterSetOutput) ToX12CharacterSetOutputWithContext(ctx context.Context) X12CharacterSetOutput {
	return o
}

func (o X12CharacterSetOutput) ToX12CharacterSetPtrOutput() X12CharacterSetPtrOutput {
	return o.ToX12CharacterSetPtrOutputWithContext(context.Background())
}

func (o X12CharacterSetOutput) ToX12CharacterSetPtrOutputWithContext(ctx context.Context) X12CharacterSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12CharacterSet) *X12CharacterSet {
		return &v
	}).(X12CharacterSetPtrOutput)
}

func (o X12CharacterSetOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o X12CharacterSetOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e X12CharacterSet) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o X12CharacterSetOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o X12CharacterSetOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e X12CharacterSet) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type X12CharacterSetPtrOutput struct{ *pulumi.OutputState }

func (X12CharacterSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12CharacterSet)(nil)).Elem()
}

func (o X12CharacterSetPtrOutput) ToX12CharacterSetPtrOutput() X12CharacterSetPtrOutput {
	return o
}

func (o X12CharacterSetPtrOutput) ToX12CharacterSetPtrOutputWithContext(ctx context.Context) X12CharacterSetPtrOutput {
	return o
}

func (o X12CharacterSetPtrOutput) Elem() X12CharacterSetOutput {
	return o.ApplyT(func(v *X12CharacterSet) X12CharacterSet {
		if v != nil {
			return *v
		}
		var ret X12CharacterSet
		return ret
	}).(X12CharacterSetOutput)
}

func (o X12CharacterSetPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o X12CharacterSetPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *X12CharacterSet) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type X12CharacterSetInput interface {
	pulumi.Input

	ToX12CharacterSetOutput() X12CharacterSetOutput
	ToX12CharacterSetOutputWithContext(context.Context) X12CharacterSetOutput
}

var x12characterSetPtrType = reflect.TypeOf((**X12CharacterSet)(nil)).Elem()

type X12CharacterSetPtrInput interface {
	pulumi.Input

	ToX12CharacterSetPtrOutput() X12CharacterSetPtrOutput
	ToX12CharacterSetPtrOutputWithContext(context.Context) X12CharacterSetPtrOutput
}

type x12characterSetPtr string

func X12CharacterSetPtr(v string) X12CharacterSetPtrInput {
	return (*x12characterSetPtr)(&v)
}

func (*x12characterSetPtr) ElementType() reflect.Type {
	return x12characterSetPtrType
}

func (in *x12characterSetPtr) ToX12CharacterSetPtrOutput() X12CharacterSetPtrOutput {
	return pulumi.ToOutput(in).(X12CharacterSetPtrOutput)
}

func (in *x12characterSetPtr) ToX12CharacterSetPtrOutputWithContext(ctx context.Context) X12CharacterSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(X12CharacterSetPtrOutput)
}

type X12DateFormat string

const (
	X12DateFormatNotSpecified = X12DateFormat("NotSpecified")
	X12DateFormatCCYYMMDD     = X12DateFormat("CCYYMMDD")
	X12DateFormatYYMMDD       = X12DateFormat("YYMMDD")
)

func (X12DateFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*X12DateFormat)(nil)).Elem()
}

func (e X12DateFormat) ToX12DateFormatOutput() X12DateFormatOutput {
	return pulumi.ToOutput(e).(X12DateFormatOutput)
}

func (e X12DateFormat) ToX12DateFormatOutputWithContext(ctx context.Context) X12DateFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(X12DateFormatOutput)
}

func (e X12DateFormat) ToX12DateFormatPtrOutput() X12DateFormatPtrOutput {
	return e.ToX12DateFormatPtrOutputWithContext(context.Background())
}

func (e X12DateFormat) ToX12DateFormatPtrOutputWithContext(ctx context.Context) X12DateFormatPtrOutput {
	return X12DateFormat(e).ToX12DateFormatOutputWithContext(ctx).ToX12DateFormatPtrOutputWithContext(ctx)
}

func (e X12DateFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e X12DateFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e X12DateFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e X12DateFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type X12DateFormatOutput struct{ *pulumi.OutputState }

func (X12DateFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12DateFormat)(nil)).Elem()
}

func (o X12DateFormatOutput) ToX12DateFormatOutput() X12DateFormatOutput {
	return o
}

func (o X12DateFormatOutput) ToX12DateFormatOutputWithContext(ctx context.Context) X12DateFormatOutput {
	return o
}

func (o X12DateFormatOutput) ToX12DateFormatPtrOutput() X12DateFormatPtrOutput {
	return o.ToX12DateFormatPtrOutputWithContext(context.Background())
}

func (o X12DateFormatOutput) ToX12DateFormatPtrOutputWithContext(ctx context.Context) X12DateFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12DateFormat) *X12DateFormat {
		return &v
	}).(X12DateFormatPtrOutput)
}

func (o X12DateFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o X12DateFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e X12DateFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o X12DateFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o X12DateFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e X12DateFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type X12DateFormatPtrOutput struct{ *pulumi.OutputState }

func (X12DateFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12DateFormat)(nil)).Elem()
}

func (o X12DateFormatPtrOutput) ToX12DateFormatPtrOutput() X12DateFormatPtrOutput {
	return o
}

func (o X12DateFormatPtrOutput) ToX12DateFormatPtrOutputWithContext(ctx context.Context) X12DateFormatPtrOutput {
	return o
}

func (o X12DateFormatPtrOutput) Elem() X12DateFormatOutput {
	return o.ApplyT(func(v *X12DateFormat) X12DateFormat {
		if v != nil {
			return *v
		}
		var ret X12DateFormat
		return ret
	}).(X12DateFormatOutput)
}

func (o X12DateFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o X12DateFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *X12DateFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type X12DateFormatInput interface {
	pulumi.Input

	ToX12DateFormatOutput() X12DateFormatOutput
	ToX12DateFormatOutputWithContext(context.Context) X12DateFormatOutput
}

var x12dateFormatPtrType = reflect.TypeOf((**X12DateFormat)(nil)).Elem()

type X12DateFormatPtrInput interface {
	pulumi.Input

	ToX12DateFormatPtrOutput() X12DateFormatPtrOutput
	ToX12DateFormatPtrOutputWithContext(context.Context) X12DateFormatPtrOutput
}

type x12dateFormatPtr string

func X12DateFormatPtr(v string) X12DateFormatPtrInput {
	return (*x12dateFormatPtr)(&v)
}

func (*x12dateFormatPtr) ElementType() reflect.Type {
	return x12dateFormatPtrType
}

func (in *x12dateFormatPtr) ToX12DateFormatPtrOutput() X12DateFormatPtrOutput {
	return pulumi.ToOutput(in).(X12DateFormatPtrOutput)
}

func (in *x12dateFormatPtr) ToX12DateFormatPtrOutputWithContext(ctx context.Context) X12DateFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(X12DateFormatPtrOutput)
}

type X12TimeFormat string

const (
	X12TimeFormatNotSpecified = X12TimeFormat("NotSpecified")
	X12TimeFormatHHMM         = X12TimeFormat("HHMM")
	X12TimeFormatHHMMSS       = X12TimeFormat("HHMMSS")
	X12TimeFormatHHMMSSdd     = X12TimeFormat("HHMMSSdd")
	X12TimeFormatHHMMSSd      = X12TimeFormat("HHMMSSd")
)

func (X12TimeFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*X12TimeFormat)(nil)).Elem()
}

func (e X12TimeFormat) ToX12TimeFormatOutput() X12TimeFormatOutput {
	return pulumi.ToOutput(e).(X12TimeFormatOutput)
}

func (e X12TimeFormat) ToX12TimeFormatOutputWithContext(ctx context.Context) X12TimeFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(X12TimeFormatOutput)
}

func (e X12TimeFormat) ToX12TimeFormatPtrOutput() X12TimeFormatPtrOutput {
	return e.ToX12TimeFormatPtrOutputWithContext(context.Background())
}

func (e X12TimeFormat) ToX12TimeFormatPtrOutputWithContext(ctx context.Context) X12TimeFormatPtrOutput {
	return X12TimeFormat(e).ToX12TimeFormatOutputWithContext(ctx).ToX12TimeFormatPtrOutputWithContext(ctx)
}

func (e X12TimeFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e X12TimeFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e X12TimeFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e X12TimeFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type X12TimeFormatOutput struct{ *pulumi.OutputState }

func (X12TimeFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*X12TimeFormat)(nil)).Elem()
}

func (o X12TimeFormatOutput) ToX12TimeFormatOutput() X12TimeFormatOutput {
	return o
}

func (o X12TimeFormatOutput) ToX12TimeFormatOutputWithContext(ctx context.Context) X12TimeFormatOutput {
	return o
}

func (o X12TimeFormatOutput) ToX12TimeFormatPtrOutput() X12TimeFormatPtrOutput {
	return o.ToX12TimeFormatPtrOutputWithContext(context.Background())
}

func (o X12TimeFormatOutput) ToX12TimeFormatPtrOutputWithContext(ctx context.Context) X12TimeFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v X12TimeFormat) *X12TimeFormat {
		return &v
	}).(X12TimeFormatPtrOutput)
}

func (o X12TimeFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o X12TimeFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e X12TimeFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o X12TimeFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o X12TimeFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e X12TimeFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type X12TimeFormatPtrOutput struct{ *pulumi.OutputState }

func (X12TimeFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**X12TimeFormat)(nil)).Elem()
}

func (o X12TimeFormatPtrOutput) ToX12TimeFormatPtrOutput() X12TimeFormatPtrOutput {
	return o
}

func (o X12TimeFormatPtrOutput) ToX12TimeFormatPtrOutputWithContext(ctx context.Context) X12TimeFormatPtrOutput {
	return o
}

func (o X12TimeFormatPtrOutput) Elem() X12TimeFormatOutput {
	return o.ApplyT(func(v *X12TimeFormat) X12TimeFormat {
		if v != nil {
			return *v
		}
		var ret X12TimeFormat
		return ret
	}).(X12TimeFormatOutput)
}

func (o X12TimeFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o X12TimeFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *X12TimeFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type X12TimeFormatInput interface {
	pulumi.Input

	ToX12TimeFormatOutput() X12TimeFormatOutput
	ToX12TimeFormatOutputWithContext(context.Context) X12TimeFormatOutput
}

var x12timeFormatPtrType = reflect.TypeOf((**X12TimeFormat)(nil)).Elem()

type X12TimeFormatPtrInput interface {
	pulumi.Input

	ToX12TimeFormatPtrOutput() X12TimeFormatPtrOutput
	ToX12TimeFormatPtrOutputWithContext(context.Context) X12TimeFormatPtrOutput
}

type x12timeFormatPtr string

func X12TimeFormatPtr(v string) X12TimeFormatPtrInput {
	return (*x12timeFormatPtr)(&v)
}

func (*x12timeFormatPtr) ElementType() reflect.Type {
	return x12timeFormatPtrType
}

func (in *x12timeFormatPtr) ToX12TimeFormatPtrOutput() X12TimeFormatPtrOutput {
	return pulumi.ToOutput(in).(X12TimeFormatPtrOutput)
}

func (in *x12timeFormatPtr) ToX12TimeFormatPtrOutputWithContext(ctx context.Context) X12TimeFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(X12TimeFormatPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AgreementTypeOutput{})
	pulumi.RegisterOutputType(AgreementTypePtrOutput{})
	pulumi.RegisterOutputType(EdifactCharacterSetOutput{})
	pulumi.RegisterOutputType(EdifactCharacterSetPtrOutput{})
	pulumi.RegisterOutputType(EdifactDecimalIndicatorOutput{})
	pulumi.RegisterOutputType(EdifactDecimalIndicatorPtrOutput{})
	pulumi.RegisterOutputType(EncryptionAlgorithmOutput{})
	pulumi.RegisterOutputType(EncryptionAlgorithmPtrOutput{})
	pulumi.RegisterOutputType(HashingAlgorithmOutput{})
	pulumi.RegisterOutputType(HashingAlgorithmPtrOutput{})
	pulumi.RegisterOutputType(MapTypeOutput{})
	pulumi.RegisterOutputType(MapTypePtrOutput{})
	pulumi.RegisterOutputType(MessageFilterTypeOutput{})
	pulumi.RegisterOutputType(MessageFilterTypePtrOutput{})
	pulumi.RegisterOutputType(PartnerTypeOutput{})
	pulumi.RegisterOutputType(PartnerTypePtrOutput{})
	pulumi.RegisterOutputType(SchemaTypeOutput{})
	pulumi.RegisterOutputType(SchemaTypePtrOutput{})
	pulumi.RegisterOutputType(SegmentTerminatorSuffixOutput{})
	pulumi.RegisterOutputType(SegmentTerminatorSuffixPtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
	pulumi.RegisterOutputType(TrailingSeparatorPolicyOutput{})
	pulumi.RegisterOutputType(TrailingSeparatorPolicyPtrOutput{})
	pulumi.RegisterOutputType(UsageIndicatorOutput{})
	pulumi.RegisterOutputType(UsageIndicatorPtrOutput{})
	pulumi.RegisterOutputType(X12CharacterSetOutput{})
	pulumi.RegisterOutputType(X12CharacterSetPtrOutput{})
	pulumi.RegisterOutputType(X12DateFormatOutput{})
	pulumi.RegisterOutputType(X12DateFormatPtrOutput{})
	pulumi.RegisterOutputType(X12TimeFormatOutput{})
	pulumi.RegisterOutputType(X12TimeFormatPtrOutput{})
}
