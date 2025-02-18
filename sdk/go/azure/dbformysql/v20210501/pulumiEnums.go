


package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CreateMode string

const (
	CreateModeDefault            = CreateMode("Default")
	CreateModePointInTimeRestore = CreateMode("PointInTimeRestore")
	CreateModeReplica            = CreateMode("Replica")
	CreateModeGeoRestore         = CreateMode("GeoRestore")
)

type DataEncryptionType string

const (
	DataEncryptionTypeAzureKeyVault = DataEncryptionType("AzureKeyVault")
	DataEncryptionTypeSystemManaged = DataEncryptionType("SystemManaged")
)

func (DataEncryptionType) ElementType() reflect.Type {
	return reflect.TypeOf((*DataEncryptionType)(nil)).Elem()
}

func (e DataEncryptionType) ToDataEncryptionTypeOutput() DataEncryptionTypeOutput {
	return pulumi.ToOutput(e).(DataEncryptionTypeOutput)
}

func (e DataEncryptionType) ToDataEncryptionTypeOutputWithContext(ctx context.Context) DataEncryptionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataEncryptionTypeOutput)
}

func (e DataEncryptionType) ToDataEncryptionTypePtrOutput() DataEncryptionTypePtrOutput {
	return e.ToDataEncryptionTypePtrOutputWithContext(context.Background())
}

func (e DataEncryptionType) ToDataEncryptionTypePtrOutputWithContext(ctx context.Context) DataEncryptionTypePtrOutput {
	return DataEncryptionType(e).ToDataEncryptionTypeOutputWithContext(ctx).ToDataEncryptionTypePtrOutputWithContext(ctx)
}

func (e DataEncryptionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataEncryptionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataEncryptionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataEncryptionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataEncryptionTypeOutput struct{ *pulumi.OutputState }

func (DataEncryptionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataEncryptionType)(nil)).Elem()
}

func (o DataEncryptionTypeOutput) ToDataEncryptionTypeOutput() DataEncryptionTypeOutput {
	return o
}

func (o DataEncryptionTypeOutput) ToDataEncryptionTypeOutputWithContext(ctx context.Context) DataEncryptionTypeOutput {
	return o
}

func (o DataEncryptionTypeOutput) ToDataEncryptionTypePtrOutput() DataEncryptionTypePtrOutput {
	return o.ToDataEncryptionTypePtrOutputWithContext(context.Background())
}

func (o DataEncryptionTypeOutput) ToDataEncryptionTypePtrOutputWithContext(ctx context.Context) DataEncryptionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataEncryptionType) *DataEncryptionType {
		return &v
	}).(DataEncryptionTypePtrOutput)
}

func (o DataEncryptionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataEncryptionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataEncryptionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataEncryptionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataEncryptionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataEncryptionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataEncryptionTypePtrOutput struct{ *pulumi.OutputState }

func (DataEncryptionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataEncryptionType)(nil)).Elem()
}

func (o DataEncryptionTypePtrOutput) ToDataEncryptionTypePtrOutput() DataEncryptionTypePtrOutput {
	return o
}

func (o DataEncryptionTypePtrOutput) ToDataEncryptionTypePtrOutputWithContext(ctx context.Context) DataEncryptionTypePtrOutput {
	return o
}

func (o DataEncryptionTypePtrOutput) Elem() DataEncryptionTypeOutput {
	return o.ApplyT(func(v *DataEncryptionType) DataEncryptionType {
		if v != nil {
			return *v
		}
		var ret DataEncryptionType
		return ret
	}).(DataEncryptionTypeOutput)
}

func (o DataEncryptionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataEncryptionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataEncryptionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DataEncryptionTypeInput interface {
	pulumi.Input

	ToDataEncryptionTypeOutput() DataEncryptionTypeOutput
	ToDataEncryptionTypeOutputWithContext(context.Context) DataEncryptionTypeOutput
}

var dataEncryptionTypePtrType = reflect.TypeOf((**DataEncryptionType)(nil)).Elem()

type DataEncryptionTypePtrInput interface {
	pulumi.Input

	ToDataEncryptionTypePtrOutput() DataEncryptionTypePtrOutput
	ToDataEncryptionTypePtrOutputWithContext(context.Context) DataEncryptionTypePtrOutput
}

type dataEncryptionTypePtr string

func DataEncryptionTypePtr(v string) DataEncryptionTypePtrInput {
	return (*dataEncryptionTypePtr)(&v)
}

func (*dataEncryptionTypePtr) ElementType() reflect.Type {
	return dataEncryptionTypePtrType
}

func (in *dataEncryptionTypePtr) ToDataEncryptionTypePtrOutput() DataEncryptionTypePtrOutput {
	return pulumi.ToOutput(in).(DataEncryptionTypePtrOutput)
}

func (in *dataEncryptionTypePtr) ToDataEncryptionTypePtrOutputWithContext(ctx context.Context) DataEncryptionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataEncryptionTypePtrOutput)
}

type EnableStatusEnum string

const (
	EnableStatusEnumEnabled  = EnableStatusEnum("Enabled")
	EnableStatusEnumDisabled = EnableStatusEnum("Disabled")
)

type HighAvailabilityMode string

const (
	HighAvailabilityModeDisabled      = HighAvailabilityMode("Disabled")
	HighAvailabilityModeZoneRedundant = HighAvailabilityMode("ZoneRedundant")
	HighAvailabilityModeSameZone      = HighAvailabilityMode("SameZone")
)

type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeUserAssigned = ManagedServiceIdentityType("UserAssigned")
)

func (ManagedServiceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityType)(nil)).Elem()
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ManagedServiceIdentityTypeOutput)
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypeOutputWithContext(ctx context.Context) ManagedServiceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedServiceIdentityTypeOutput)
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return e.ToManagedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ManagedServiceIdentityType) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return ManagedServiceIdentityType(e).ToManagedServiceIdentityTypeOutputWithContext(ctx).ToManagedServiceIdentityTypePtrOutputWithContext(ctx)
}

func (e ManagedServiceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedServiceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedServiceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedServiceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedServiceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityType)(nil)).Elem()
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput {
	return o
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypeOutputWithContext(ctx context.Context) ManagedServiceIdentityTypeOutput {
	return o
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return o.ToManagedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentityType) *ManagedServiceIdentityType {
		return &v
	}).(ManagedServiceIdentityTypePtrOutput)
}

func (o ManagedServiceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedServiceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedServiceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedServiceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityType)(nil)).Elem()
}

func (o ManagedServiceIdentityTypePtrOutput) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return o
}

func (o ManagedServiceIdentityTypePtrOutput) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return o
}

func (o ManagedServiceIdentityTypePtrOutput) Elem() ManagedServiceIdentityTypeOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityType) ManagedServiceIdentityType {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityType
		return ret
	}).(ManagedServiceIdentityTypeOutput)
}

func (o ManagedServiceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedServiceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ManagedServiceIdentityTypeInput interface {
	pulumi.Input

	ToManagedServiceIdentityTypeOutput() ManagedServiceIdentityTypeOutput
	ToManagedServiceIdentityTypeOutputWithContext(context.Context) ManagedServiceIdentityTypeOutput
}

var managedServiceIdentityTypePtrType = reflect.TypeOf((**ManagedServiceIdentityType)(nil)).Elem()

type ManagedServiceIdentityTypePtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput
	ToManagedServiceIdentityTypePtrOutputWithContext(context.Context) ManagedServiceIdentityTypePtrOutput
}

type managedServiceIdentityTypePtr string

func ManagedServiceIdentityTypePtr(v string) ManagedServiceIdentityTypePtrInput {
	return (*managedServiceIdentityTypePtr)(&v)
}

func (*managedServiceIdentityTypePtr) ElementType() reflect.Type {
	return managedServiceIdentityTypePtrType
}

func (in *managedServiceIdentityTypePtr) ToManagedServiceIdentityTypePtrOutput() ManagedServiceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ManagedServiceIdentityTypePtrOutput)
}

func (in *managedServiceIdentityTypePtr) ToManagedServiceIdentityTypePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedServiceIdentityTypePtrOutput)
}

type ReplicationRole string

const (
	ReplicationRoleNone    = ReplicationRole("None")
	ReplicationRoleSource  = ReplicationRole("Source")
	ReplicationRoleReplica = ReplicationRole("Replica")
)

type ServerVersion string

const (
	ServerVersion_5_7    = ServerVersion("5.7")
	ServerVersion_8_0_21 = ServerVersion("8.0.21")
)

type SkuTier string

const (
	SkuTierBurstable       = SkuTier("Burstable")
	SkuTierGeneralPurpose  = SkuTier("GeneralPurpose")
	SkuTierMemoryOptimized = SkuTier("MemoryOptimized")
)

func init() {
	pulumi.RegisterOutputType(DataEncryptionTypeOutput{})
	pulumi.RegisterOutputType(DataEncryptionTypePtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityTypePtrOutput{})
}
