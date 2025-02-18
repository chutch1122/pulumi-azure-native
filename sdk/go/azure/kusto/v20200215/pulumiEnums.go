


package v20200215

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
	AzureSkuName_Standard_E2a_v4            = AzureSkuName("Standard_E2a_v4")
	AzureSkuName_Standard_E4a_v4            = AzureSkuName("Standard_E4a_v4")
	AzureSkuName_Standard_E8a_v4            = AzureSkuName("Standard_E8a_v4")
	AzureSkuName_Standard_E16a_v4           = AzureSkuName("Standard_E16a_v4")
	AzureSkuName_Standard_E8as_v4_1TB_PS    = AzureSkuName("Standard_E8as_v4+1TB_PS")
	AzureSkuName_Standard_E8as_v4_2TB_PS    = AzureSkuName("Standard_E8as_v4+2TB_PS")
	AzureSkuName_Standard_E16as_v4_3TB_PS   = AzureSkuName("Standard_E16as_v4+3TB_PS")
	AzureSkuName_Standard_E16as_v4_4TB_PS   = AzureSkuName("Standard_E16as_v4+4TB_PS")
	AzureSkuName_Dev_No_SLA_Standard_E2a_v4 = AzureSkuName("Dev(No SLA)_Standard_E2a_v4")
)

type AzureSkuTier string

const (
	AzureSkuTierBasic    = AzureSkuTier("Basic")
	AzureSkuTierStandard = AzureSkuTier("Standard")
)

type ClusterPrincipalRole string

const (
	ClusterPrincipalRoleAllDatabasesAdmin  = ClusterPrincipalRole("AllDatabasesAdmin")
	ClusterPrincipalRoleAllDatabasesViewer = ClusterPrincipalRole("AllDatabasesViewer")
)

type Compression string

const (
	CompressionNone = Compression("None")
	CompressionGZip = Compression("GZip")
)

type DatabasePrincipalRole string

const (
	DatabasePrincipalRoleAdmin               = DatabasePrincipalRole("Admin")
	DatabasePrincipalRoleIngestor            = DatabasePrincipalRole("Ingestor")
	DatabasePrincipalRoleMonitor             = DatabasePrincipalRole("Monitor")
	DatabasePrincipalRoleUser                = DatabasePrincipalRole("User")
	DatabasePrincipalRoleUnrestrictedViewers = DatabasePrincipalRole("UnrestrictedViewers")
	DatabasePrincipalRoleViewer              = DatabasePrincipalRole("Viewer")
)

type DefaultPrincipalsModificationKind string

const (
	DefaultPrincipalsModificationKindUnion   = DefaultPrincipalsModificationKind("Union")
	DefaultPrincipalsModificationKindReplace = DefaultPrincipalsModificationKind("Replace")
	DefaultPrincipalsModificationKindNone    = DefaultPrincipalsModificationKind("None")
)

type EventGridDataFormat string

const (
	EventGridDataFormatMULTIJSON  = EventGridDataFormat("MULTIJSON")
	EventGridDataFormatJSON       = EventGridDataFormat("JSON")
	EventGridDataFormatCSV        = EventGridDataFormat("CSV")
	EventGridDataFormatTSV        = EventGridDataFormat("TSV")
	EventGridDataFormatSCSV       = EventGridDataFormat("SCSV")
	EventGridDataFormatSOHSV      = EventGridDataFormat("SOHSV")
	EventGridDataFormatPSV        = EventGridDataFormat("PSV")
	EventGridDataFormatTXT        = EventGridDataFormat("TXT")
	EventGridDataFormatRAW        = EventGridDataFormat("RAW")
	EventGridDataFormatSINGLEJSON = EventGridDataFormat("SINGLEJSON")
	EventGridDataFormatAVRO       = EventGridDataFormat("AVRO")
	EventGridDataFormatTSVE       = EventGridDataFormat("TSVE")
	EventGridDataFormatPARQUET    = EventGridDataFormat("PARQUET")
	EventGridDataFormatORC        = EventGridDataFormat("ORC")
)

type EventHubDataFormat string

const (
	EventHubDataFormatMULTIJSON  = EventHubDataFormat("MULTIJSON")
	EventHubDataFormatJSON       = EventHubDataFormat("JSON")
	EventHubDataFormatCSV        = EventHubDataFormat("CSV")
	EventHubDataFormatTSV        = EventHubDataFormat("TSV")
	EventHubDataFormatSCSV       = EventHubDataFormat("SCSV")
	EventHubDataFormatSOHSV      = EventHubDataFormat("SOHSV")
	EventHubDataFormatPSV        = EventHubDataFormat("PSV")
	EventHubDataFormatTXT        = EventHubDataFormat("TXT")
	EventHubDataFormatRAW        = EventHubDataFormat("RAW")
	EventHubDataFormatSINGLEJSON = EventHubDataFormat("SINGLEJSON")
	EventHubDataFormatAVRO       = EventHubDataFormat("AVRO")
	EventHubDataFormatTSVE       = EventHubDataFormat("TSVE")
	EventHubDataFormatPARQUET    = EventHubDataFormat("PARQUET")
	EventHubDataFormatORC        = EventHubDataFormat("ORC")
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

type IotHubDataFormat string

const (
	IotHubDataFormatMULTIJSON  = IotHubDataFormat("MULTIJSON")
	IotHubDataFormatJSON       = IotHubDataFormat("JSON")
	IotHubDataFormatCSV        = IotHubDataFormat("CSV")
	IotHubDataFormatTSV        = IotHubDataFormat("TSV")
	IotHubDataFormatSCSV       = IotHubDataFormat("SCSV")
	IotHubDataFormatSOHSV      = IotHubDataFormat("SOHSV")
	IotHubDataFormatPSV        = IotHubDataFormat("PSV")
	IotHubDataFormatTXT        = IotHubDataFormat("TXT")
	IotHubDataFormatRAW        = IotHubDataFormat("RAW")
	IotHubDataFormatSINGLEJSON = IotHubDataFormat("SINGLEJSON")
	IotHubDataFormatAVRO       = IotHubDataFormat("AVRO")
	IotHubDataFormatTSVE       = IotHubDataFormat("TSVE")
	IotHubDataFormatPARQUET    = IotHubDataFormat("PARQUET")
	IotHubDataFormatORC        = IotHubDataFormat("ORC")
)

type Kind string

const (
	KindReadWrite         = Kind("ReadWrite")
	KindReadOnlyFollowing = Kind("ReadOnlyFollowing")
)

type LanguageExtensionName string

const (
	LanguageExtensionNamePYTHON = LanguageExtensionName("PYTHON")
	LanguageExtensionNameR      = LanguageExtensionName("R")
)

type PrincipalType string

const (
	PrincipalTypeApp   = PrincipalType("App")
	PrincipalTypeGroup = PrincipalType("Group")
	PrincipalTypeUser  = PrincipalType("User")
)

func init() {
	pulumi.RegisterOutputType(IdentityTypeOutput{})
	pulumi.RegisterOutputType(IdentityTypePtrOutput{})
}
