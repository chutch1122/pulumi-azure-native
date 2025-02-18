


package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AutoExecuteStatus string

const (
	AutoExecuteStatusEnabled  = AutoExecuteStatus("Enabled")
	AutoExecuteStatusDisabled = AutoExecuteStatus("Disabled")
	AutoExecuteStatusDefault  = AutoExecuteStatus("Default")
)

func (AutoExecuteStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoExecuteStatus)(nil)).Elem()
}

func (e AutoExecuteStatus) ToAutoExecuteStatusOutput() AutoExecuteStatusOutput {
	return pulumi.ToOutput(e).(AutoExecuteStatusOutput)
}

func (e AutoExecuteStatus) ToAutoExecuteStatusOutputWithContext(ctx context.Context) AutoExecuteStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutoExecuteStatusOutput)
}

func (e AutoExecuteStatus) ToAutoExecuteStatusPtrOutput() AutoExecuteStatusPtrOutput {
	return e.ToAutoExecuteStatusPtrOutputWithContext(context.Background())
}

func (e AutoExecuteStatus) ToAutoExecuteStatusPtrOutputWithContext(ctx context.Context) AutoExecuteStatusPtrOutput {
	return AutoExecuteStatus(e).ToAutoExecuteStatusOutputWithContext(ctx).ToAutoExecuteStatusPtrOutputWithContext(ctx)
}

func (e AutoExecuteStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoExecuteStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoExecuteStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutoExecuteStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutoExecuteStatusOutput struct{ *pulumi.OutputState }

func (AutoExecuteStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoExecuteStatus)(nil)).Elem()
}

func (o AutoExecuteStatusOutput) ToAutoExecuteStatusOutput() AutoExecuteStatusOutput {
	return o
}

func (o AutoExecuteStatusOutput) ToAutoExecuteStatusOutputWithContext(ctx context.Context) AutoExecuteStatusOutput {
	return o
}

func (o AutoExecuteStatusOutput) ToAutoExecuteStatusPtrOutput() AutoExecuteStatusPtrOutput {
	return o.ToAutoExecuteStatusPtrOutputWithContext(context.Background())
}

func (o AutoExecuteStatusOutput) ToAutoExecuteStatusPtrOutputWithContext(ctx context.Context) AutoExecuteStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoExecuteStatus) *AutoExecuteStatus {
		return &v
	}).(AutoExecuteStatusPtrOutput)
}

func (o AutoExecuteStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutoExecuteStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoExecuteStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutoExecuteStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoExecuteStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoExecuteStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutoExecuteStatusPtrOutput struct{ *pulumi.OutputState }

func (AutoExecuteStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoExecuteStatus)(nil)).Elem()
}

func (o AutoExecuteStatusPtrOutput) ToAutoExecuteStatusPtrOutput() AutoExecuteStatusPtrOutput {
	return o
}

func (o AutoExecuteStatusPtrOutput) ToAutoExecuteStatusPtrOutputWithContext(ctx context.Context) AutoExecuteStatusPtrOutput {
	return o
}

func (o AutoExecuteStatusPtrOutput) Elem() AutoExecuteStatusOutput {
	return o.ApplyT(func(v *AutoExecuteStatus) AutoExecuteStatus {
		if v != nil {
			return *v
		}
		var ret AutoExecuteStatus
		return ret
	}).(AutoExecuteStatusOutput)
}

func (o AutoExecuteStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoExecuteStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutoExecuteStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AutoExecuteStatusInput interface {
	pulumi.Input

	ToAutoExecuteStatusOutput() AutoExecuteStatusOutput
	ToAutoExecuteStatusOutputWithContext(context.Context) AutoExecuteStatusOutput
}

var autoExecuteStatusPtrType = reflect.TypeOf((**AutoExecuteStatus)(nil)).Elem()

type AutoExecuteStatusPtrInput interface {
	pulumi.Input

	ToAutoExecuteStatusPtrOutput() AutoExecuteStatusPtrOutput
	ToAutoExecuteStatusPtrOutputWithContext(context.Context) AutoExecuteStatusPtrOutput
}

type autoExecuteStatusPtr string

func AutoExecuteStatusPtr(v string) AutoExecuteStatusPtrInput {
	return (*autoExecuteStatusPtr)(&v)
}

func (*autoExecuteStatusPtr) ElementType() reflect.Type {
	return autoExecuteStatusPtrType
}

func (in *autoExecuteStatusPtr) ToAutoExecuteStatusPtrOutput() AutoExecuteStatusPtrOutput {
	return pulumi.ToOutput(in).(AutoExecuteStatusPtrOutput)
}

func (in *autoExecuteStatusPtr) ToAutoExecuteStatusPtrOutputWithContext(ctx context.Context) AutoExecuteStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutoExecuteStatusPtrOutput)
}

type BlobAuditingPolicyState string

const (
	BlobAuditingPolicyStateEnabled  = BlobAuditingPolicyState("Enabled")
	BlobAuditingPolicyStateDisabled = BlobAuditingPolicyState("Disabled")
)

func (BlobAuditingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobAuditingPolicyState)(nil)).Elem()
}

func (e BlobAuditingPolicyState) ToBlobAuditingPolicyStateOutput() BlobAuditingPolicyStateOutput {
	return pulumi.ToOutput(e).(BlobAuditingPolicyStateOutput)
}

func (e BlobAuditingPolicyState) ToBlobAuditingPolicyStateOutputWithContext(ctx context.Context) BlobAuditingPolicyStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BlobAuditingPolicyStateOutput)
}

func (e BlobAuditingPolicyState) ToBlobAuditingPolicyStatePtrOutput() BlobAuditingPolicyStatePtrOutput {
	return e.ToBlobAuditingPolicyStatePtrOutputWithContext(context.Background())
}

func (e BlobAuditingPolicyState) ToBlobAuditingPolicyStatePtrOutputWithContext(ctx context.Context) BlobAuditingPolicyStatePtrOutput {
	return BlobAuditingPolicyState(e).ToBlobAuditingPolicyStateOutputWithContext(ctx).ToBlobAuditingPolicyStatePtrOutputWithContext(ctx)
}

func (e BlobAuditingPolicyState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlobAuditingPolicyState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlobAuditingPolicyState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BlobAuditingPolicyState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BlobAuditingPolicyStateOutput struct{ *pulumi.OutputState }

func (BlobAuditingPolicyStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobAuditingPolicyState)(nil)).Elem()
}

func (o BlobAuditingPolicyStateOutput) ToBlobAuditingPolicyStateOutput() BlobAuditingPolicyStateOutput {
	return o
}

func (o BlobAuditingPolicyStateOutput) ToBlobAuditingPolicyStateOutputWithContext(ctx context.Context) BlobAuditingPolicyStateOutput {
	return o
}

func (o BlobAuditingPolicyStateOutput) ToBlobAuditingPolicyStatePtrOutput() BlobAuditingPolicyStatePtrOutput {
	return o.ToBlobAuditingPolicyStatePtrOutputWithContext(context.Background())
}

func (o BlobAuditingPolicyStateOutput) ToBlobAuditingPolicyStatePtrOutputWithContext(ctx context.Context) BlobAuditingPolicyStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlobAuditingPolicyState) *BlobAuditingPolicyState {
		return &v
	}).(BlobAuditingPolicyStatePtrOutput)
}

func (o BlobAuditingPolicyStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BlobAuditingPolicyStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlobAuditingPolicyState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BlobAuditingPolicyStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlobAuditingPolicyStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlobAuditingPolicyState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BlobAuditingPolicyStatePtrOutput struct{ *pulumi.OutputState }

func (BlobAuditingPolicyStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobAuditingPolicyState)(nil)).Elem()
}

func (o BlobAuditingPolicyStatePtrOutput) ToBlobAuditingPolicyStatePtrOutput() BlobAuditingPolicyStatePtrOutput {
	return o
}

func (o BlobAuditingPolicyStatePtrOutput) ToBlobAuditingPolicyStatePtrOutputWithContext(ctx context.Context) BlobAuditingPolicyStatePtrOutput {
	return o
}

func (o BlobAuditingPolicyStatePtrOutput) Elem() BlobAuditingPolicyStateOutput {
	return o.ApplyT(func(v *BlobAuditingPolicyState) BlobAuditingPolicyState {
		if v != nil {
			return *v
		}
		var ret BlobAuditingPolicyState
		return ret
	}).(BlobAuditingPolicyStateOutput)
}

func (o BlobAuditingPolicyStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlobAuditingPolicyStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BlobAuditingPolicyState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BlobAuditingPolicyStateInput interface {
	pulumi.Input

	ToBlobAuditingPolicyStateOutput() BlobAuditingPolicyStateOutput
	ToBlobAuditingPolicyStateOutputWithContext(context.Context) BlobAuditingPolicyStateOutput
}

var blobAuditingPolicyStatePtrType = reflect.TypeOf((**BlobAuditingPolicyState)(nil)).Elem()

type BlobAuditingPolicyStatePtrInput interface {
	pulumi.Input

	ToBlobAuditingPolicyStatePtrOutput() BlobAuditingPolicyStatePtrOutput
	ToBlobAuditingPolicyStatePtrOutputWithContext(context.Context) BlobAuditingPolicyStatePtrOutput
}

type blobAuditingPolicyStatePtr string

func BlobAuditingPolicyStatePtr(v string) BlobAuditingPolicyStatePtrInput {
	return (*blobAuditingPolicyStatePtr)(&v)
}

func (*blobAuditingPolicyStatePtr) ElementType() reflect.Type {
	return blobAuditingPolicyStatePtrType
}

func (in *blobAuditingPolicyStatePtr) ToBlobAuditingPolicyStatePtrOutput() BlobAuditingPolicyStatePtrOutput {
	return pulumi.ToOutput(in).(BlobAuditingPolicyStatePtrOutput)
}

func (in *blobAuditingPolicyStatePtr) ToBlobAuditingPolicyStatePtrOutputWithContext(ctx context.Context) BlobAuditingPolicyStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BlobAuditingPolicyStatePtrOutput)
}

type IdentityType string

const (
	IdentityTypeNone                         = IdentityType("None")
	IdentityTypeSystemAssigned               = IdentityType("SystemAssigned")
	IdentityTypeUserAssigned                 = IdentityType("UserAssigned")
	IdentityType_SystemAssigned_UserAssigned = IdentityType("SystemAssigned,UserAssigned")
)

type ManagedInstanceLicenseType string

const (
	ManagedInstanceLicenseTypeLicenseIncluded = ManagedInstanceLicenseType("LicenseIncluded")
	ManagedInstanceLicenseTypeBasePrice       = ManagedInstanceLicenseType("BasePrice")
)

type ManagedInstanceProxyOverride string

const (
	ManagedInstanceProxyOverrideProxy    = ManagedInstanceProxyOverride("Proxy")
	ManagedInstanceProxyOverrideRedirect = ManagedInstanceProxyOverride("Redirect")
	ManagedInstanceProxyOverrideDefault  = ManagedInstanceProxyOverride("Default")
)

type ManagedServerCreateMode string

const (
	ManagedServerCreateModeDefault            = ManagedServerCreateMode("Default")
	ManagedServerCreateModePointInTimeRestore = ManagedServerCreateMode("PointInTimeRestore")
)

type ReadOnlyEndpointFailoverPolicy string

const (
	ReadOnlyEndpointFailoverPolicyDisabled = ReadOnlyEndpointFailoverPolicy("Disabled")
	ReadOnlyEndpointFailoverPolicyEnabled  = ReadOnlyEndpointFailoverPolicy("Enabled")
)

type ReadWriteEndpointFailoverPolicy string

const (
	ReadWriteEndpointFailoverPolicyManual    = ReadWriteEndpointFailoverPolicy("Manual")
	ReadWriteEndpointFailoverPolicyAutomatic = ReadWriteEndpointFailoverPolicy("Automatic")
)

type ServerKeyType string

const (
	ServerKeyTypeServiceManaged = ServerKeyType("ServiceManaged")
	ServerKeyTypeAzureKeyVault  = ServerKeyType("AzureKeyVault")
)

type SyncConflictResolutionPolicy string

const (
	SyncConflictResolutionPolicyHubWin    = SyncConflictResolutionPolicy("HubWin")
	SyncConflictResolutionPolicyMemberWin = SyncConflictResolutionPolicy("MemberWin")
)

type SyncDirection string

const (
	SyncDirectionBidirectional     = SyncDirection("Bidirectional")
	SyncDirectionOneWayMemberToHub = SyncDirection("OneWayMemberToHub")
	SyncDirectionOneWayHubToMember = SyncDirection("OneWayHubToMember")
)

type SyncMemberDbType string

const (
	SyncMemberDbTypeAzureSqlDatabase  = SyncMemberDbType("AzureSqlDatabase")
	SyncMemberDbTypeSqlServerDatabase = SyncMemberDbType("SqlServerDatabase")
)

func init() {
	pulumi.RegisterOutputType(AutoExecuteStatusOutput{})
	pulumi.RegisterOutputType(AutoExecuteStatusPtrOutput{})
	pulumi.RegisterOutputType(BlobAuditingPolicyStateOutput{})
	pulumi.RegisterOutputType(BlobAuditingPolicyStatePtrOutput{})
}
