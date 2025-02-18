


package v20190907

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureSkuName string

const (
	AzureSkuName_Standard_DS13_v2_1TB_PS    = AzureSkuName("Standard_DS13_v2+1TB_PS")
	AzureSkuName_Standard_DS13_v2_2TB_PS    = AzureSkuName("Standard_DS13_v2+2TB_PS")
	AzureSkuName_Standard_DS14_v2_3TB_PS    = AzureSkuName("Standard_DS14_v2+3TB_PS")
	AzureSkuName_Standard_DS14_v2_4TB_PS    = AzureSkuName("Standard_DS14_v2+4TB_PS")
	AzureSkuName_Standard_D13_v2            = AzureSkuName("Standard_D13_v2")
	AzureSkuName_Standard_D14_v2            = AzureSkuName("Standard_D14_v2")
	AzureSkuName_Standard_L8s               = AzureSkuName("Standard_L8s")
	AzureSkuName_Standard_L16s              = AzureSkuName("Standard_L16s")
	AzureSkuName_Standard_D11_v2            = AzureSkuName("Standard_D11_v2")
	AzureSkuName_Standard_D12_v2            = AzureSkuName("Standard_D12_v2")
	AzureSkuName_Standard_L4s               = AzureSkuName("Standard_L4s")
	AzureSkuName_Dev_No_SLA_Standard_D11_v2 = AzureSkuName("Dev(No SLA)_Standard_D11_v2")
)

type AzureSkuTier string

const (
	AzureSkuTierBasic    = AzureSkuTier("Basic")
	AzureSkuTierStandard = AzureSkuTier("Standard")
)

type DataFormat string

const (
	DataFormatMULTIJSON  = DataFormat("MULTIJSON")
	DataFormatJSON       = DataFormat("JSON")
	DataFormatCSV        = DataFormat("CSV")
	DataFormatTSV        = DataFormat("TSV")
	DataFormatSCSV       = DataFormat("SCSV")
	DataFormatSOHSV      = DataFormat("SOHSV")
	DataFormatPSV        = DataFormat("PSV")
	DataFormatTXT        = DataFormat("TXT")
	DataFormatRAW        = DataFormat("RAW")
	DataFormatSINGLEJSON = DataFormat("SINGLEJSON")
	DataFormatAVRO       = DataFormat("AVRO")
	DataFormatTSVE       = DataFormat("TSVE")
)

type DefaultPrincipalsModificationKind string

const (
	DefaultPrincipalsModificationKindUnion   = DefaultPrincipalsModificationKind("Union")
	DefaultPrincipalsModificationKindReplace = DefaultPrincipalsModificationKind("Replace")
	DefaultPrincipalsModificationKindNone    = DefaultPrincipalsModificationKind("None")
)

type IdentityType string

const (
	IdentityTypeNone           = IdentityType("None")
	IdentityTypeSystemAssigned = IdentityType("SystemAssigned")
)

func (IdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (e IdentityType) ToIdentityTypeOutput() IdentityTypeOutput {
	return pulumi.ToOutput(e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return e.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (e IdentityType) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return IdentityType(e).ToIdentityTypeOutputWithContext(ctx).ToIdentityTypePtrOutputWithContext(ctx)
}

func (e IdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IdentityTypeOutput struct{ *pulumi.OutputState }

func (IdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (o IdentityTypeOutput) ToIdentityTypeOutput() IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityType) *IdentityType {
		return &v
	}).(IdentityTypePtrOutput)
}

func (o IdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IdentityTypePtrOutput struct{ *pulumi.OutputState }

func (IdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityType)(nil)).Elem()
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) Elem() IdentityTypeOutput {
	return o.ApplyT(func(v *IdentityType) IdentityType {
		if v != nil {
			return *v
		}
		var ret IdentityType
		return ret
	}).(IdentityTypeOutput)
}

func (o IdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IdentityTypeInput interface {
	pulumi.Input

	ToIdentityTypeOutput() IdentityTypeOutput
	ToIdentityTypeOutputWithContext(context.Context) IdentityTypeOutput
}

var identityTypePtrType = reflect.TypeOf((**IdentityType)(nil)).Elem()

type IdentityTypePtrInput interface {
	pulumi.Input

	ToIdentityTypePtrOutput() IdentityTypePtrOutput
	ToIdentityTypePtrOutputWithContext(context.Context) IdentityTypePtrOutput
}

type identityTypePtr string

func IdentityTypePtr(v string) IdentityTypePtrInput {
	return (*identityTypePtr)(&v)
}

func (*identityTypePtr) ElementType() reflect.Type {
	return identityTypePtrType
}

func (in *identityTypePtr) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return pulumi.ToOutput(in).(IdentityTypePtrOutput)
}

func (in *identityTypePtr) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IdentityTypePtrOutput)
}

type Kind string

const (
	KindReadWrite         = Kind("ReadWrite")
	KindReadOnlyFollowing = Kind("ReadOnlyFollowing")
)

func init() {
	pulumi.RegisterOutputType(IdentityTypeOutput{})
	pulumi.RegisterOutputType(IdentityTypePtrOutput{})
}
