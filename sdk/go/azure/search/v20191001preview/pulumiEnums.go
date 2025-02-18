


package v20191001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EndpointAccess string

const (
	EndpointAccessPublic  = EndpointAccess("Public")
	EndpointAccessPrivate = EndpointAccess("Private")
)

func (EndpointAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointAccess)(nil)).Elem()
}

func (e EndpointAccess) ToEndpointAccessOutput() EndpointAccessOutput {
	return pulumi.ToOutput(e).(EndpointAccessOutput)
}

func (e EndpointAccess) ToEndpointAccessOutputWithContext(ctx context.Context) EndpointAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EndpointAccessOutput)
}

func (e EndpointAccess) ToEndpointAccessPtrOutput() EndpointAccessPtrOutput {
	return e.ToEndpointAccessPtrOutputWithContext(context.Background())
}

func (e EndpointAccess) ToEndpointAccessPtrOutputWithContext(ctx context.Context) EndpointAccessPtrOutput {
	return EndpointAccess(e).ToEndpointAccessOutputWithContext(ctx).ToEndpointAccessPtrOutputWithContext(ctx)
}

func (e EndpointAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EndpointAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EndpointAccessOutput struct{ *pulumi.OutputState }

func (EndpointAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointAccess)(nil)).Elem()
}

func (o EndpointAccessOutput) ToEndpointAccessOutput() EndpointAccessOutput {
	return o
}

func (o EndpointAccessOutput) ToEndpointAccessOutputWithContext(ctx context.Context) EndpointAccessOutput {
	return o
}

func (o EndpointAccessOutput) ToEndpointAccessPtrOutput() EndpointAccessPtrOutput {
	return o.ToEndpointAccessPtrOutputWithContext(context.Background())
}

func (o EndpointAccessOutput) ToEndpointAccessPtrOutputWithContext(ctx context.Context) EndpointAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointAccess) *EndpointAccess {
		return &v
	}).(EndpointAccessPtrOutput)
}

func (o EndpointAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EndpointAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EndpointAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EndpointAccessPtrOutput struct{ *pulumi.OutputState }

func (EndpointAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointAccess)(nil)).Elem()
}

func (o EndpointAccessPtrOutput) ToEndpointAccessPtrOutput() EndpointAccessPtrOutput {
	return o
}

func (o EndpointAccessPtrOutput) ToEndpointAccessPtrOutputWithContext(ctx context.Context) EndpointAccessPtrOutput {
	return o
}

func (o EndpointAccessPtrOutput) Elem() EndpointAccessOutput {
	return o.ApplyT(func(v *EndpointAccess) EndpointAccess {
		if v != nil {
			return *v
		}
		var ret EndpointAccess
		return ret
	}).(EndpointAccessOutput)
}

func (o EndpointAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EndpointAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EndpointAccessInput interface {
	pulumi.Input

	ToEndpointAccessOutput() EndpointAccessOutput
	ToEndpointAccessOutputWithContext(context.Context) EndpointAccessOutput
}

var endpointAccessPtrType = reflect.TypeOf((**EndpointAccess)(nil)).Elem()

type EndpointAccessPtrInput interface {
	pulumi.Input

	ToEndpointAccessPtrOutput() EndpointAccessPtrOutput
	ToEndpointAccessPtrOutputWithContext(context.Context) EndpointAccessPtrOutput
}

type endpointAccessPtr string

func EndpointAccessPtr(v string) EndpointAccessPtrInput {
	return (*endpointAccessPtr)(&v)
}

func (*endpointAccessPtr) ElementType() reflect.Type {
	return endpointAccessPtrType
}

func (in *endpointAccessPtr) ToEndpointAccessPtrOutput() EndpointAccessPtrOutput {
	return pulumi.ToOutput(in).(EndpointAccessPtrOutput)
}

func (in *endpointAccessPtr) ToEndpointAccessPtrOutputWithContext(ctx context.Context) EndpointAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EndpointAccessPtrOutput)
}

type HostingMode string

const (
	HostingModeDefault     = HostingMode("default")
	HostingModeHighDensity = HostingMode("highDensity")
)

func (HostingMode) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingMode)(nil)).Elem()
}

func (e HostingMode) ToHostingModeOutput() HostingModeOutput {
	return pulumi.ToOutput(e).(HostingModeOutput)
}

func (e HostingMode) ToHostingModeOutputWithContext(ctx context.Context) HostingModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HostingModeOutput)
}

func (e HostingMode) ToHostingModePtrOutput() HostingModePtrOutput {
	return e.ToHostingModePtrOutputWithContext(context.Background())
}

func (e HostingMode) ToHostingModePtrOutputWithContext(ctx context.Context) HostingModePtrOutput {
	return HostingMode(e).ToHostingModeOutputWithContext(ctx).ToHostingModePtrOutputWithContext(ctx)
}

func (e HostingMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostingMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostingMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HostingMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HostingModeOutput struct{ *pulumi.OutputState }

func (HostingModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingMode)(nil)).Elem()
}

func (o HostingModeOutput) ToHostingModeOutput() HostingModeOutput {
	return o
}

func (o HostingModeOutput) ToHostingModeOutputWithContext(ctx context.Context) HostingModeOutput {
	return o
}

func (o HostingModeOutput) ToHostingModePtrOutput() HostingModePtrOutput {
	return o.ToHostingModePtrOutputWithContext(context.Background())
}

func (o HostingModeOutput) ToHostingModePtrOutputWithContext(ctx context.Context) HostingModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostingMode) *HostingMode {
		return &v
	}).(HostingModePtrOutput)
}

func (o HostingModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HostingModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostingMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HostingModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostingModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostingMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HostingModePtrOutput struct{ *pulumi.OutputState }

func (HostingModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingMode)(nil)).Elem()
}

func (o HostingModePtrOutput) ToHostingModePtrOutput() HostingModePtrOutput {
	return o
}

func (o HostingModePtrOutput) ToHostingModePtrOutputWithContext(ctx context.Context) HostingModePtrOutput {
	return o
}

func (o HostingModePtrOutput) Elem() HostingModeOutput {
	return o.ApplyT(func(v *HostingMode) HostingMode {
		if v != nil {
			return *v
		}
		var ret HostingMode
		return ret
	}).(HostingModeOutput)
}

func (o HostingModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostingModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HostingMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HostingModeInput interface {
	pulumi.Input

	ToHostingModeOutput() HostingModeOutput
	ToHostingModeOutputWithContext(context.Context) HostingModeOutput
}

var hostingModePtrType = reflect.TypeOf((**HostingMode)(nil)).Elem()

type HostingModePtrInput interface {
	pulumi.Input

	ToHostingModePtrOutput() HostingModePtrOutput
	ToHostingModePtrOutputWithContext(context.Context) HostingModePtrOutput
}

type hostingModePtr string

func HostingModePtr(v string) HostingModePtrInput {
	return (*hostingModePtr)(&v)
}

func (*hostingModePtr) ElementType() reflect.Type {
	return hostingModePtrType
}

func (in *hostingModePtr) ToHostingModePtrOutput() HostingModePtrOutput {
	return pulumi.ToOutput(in).(HostingModePtrOutput)
}

func (in *hostingModePtr) ToHostingModePtrOutputWithContext(ctx context.Context) HostingModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HostingModePtrOutput)
}

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

type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusPending      = PrivateLinkServiceConnectionStatus("Pending")
	PrivateLinkServiceConnectionStatusApproved     = PrivateLinkServiceConnectionStatus("Approved")
	PrivateLinkServiceConnectionStatusRejected     = PrivateLinkServiceConnectionStatus("Rejected")
	PrivateLinkServiceConnectionStatusDisconnected = PrivateLinkServiceConnectionStatus("Disconnected")
)

func (PrivateLinkServiceConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatus)(nil)).Elem()
}

func (e PrivateLinkServiceConnectionStatus) ToPrivateLinkServiceConnectionStatusOutput() PrivateLinkServiceConnectionStatusOutput {
	return pulumi.ToOutput(e).(PrivateLinkServiceConnectionStatusOutput)
}

func (e PrivateLinkServiceConnectionStatus) ToPrivateLinkServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrivateLinkServiceConnectionStatusOutput)
}

func (e PrivateLinkServiceConnectionStatus) ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput {
	return e.ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (e PrivateLinkServiceConnectionStatus) ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusPtrOutput {
	return PrivateLinkServiceConnectionStatus(e).ToPrivateLinkServiceConnectionStatusOutputWithContext(ctx).ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx)
}

func (e PrivateLinkServiceConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateLinkServiceConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateLinkServiceConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateLinkServiceConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrivateLinkServiceConnectionStatusOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatusOutput) ToPrivateLinkServiceConnectionStatusOutput() PrivateLinkServiceConnectionStatusOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatusOutput) ToPrivateLinkServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatusOutput) ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput {
	return o.ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatusOutput) ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStatus) *PrivateLinkServiceConnectionStatus {
		return &v
	}).(PrivateLinkServiceConnectionStatusPtrOutput)
}

func (o PrivateLinkServiceConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateLinkServiceConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateLinkServiceConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) Elem() PrivateLinkServiceConnectionStatusOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatus) PrivateLinkServiceConnectionStatus {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStatus
		return ret
	}).(PrivateLinkServiceConnectionStatusOutput)
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrivateLinkServiceConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrivateLinkServiceConnectionStatusInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatusOutput() PrivateLinkServiceConnectionStatusOutput
	ToPrivateLinkServiceConnectionStatusOutputWithContext(context.Context) PrivateLinkServiceConnectionStatusOutput
}

var privateLinkServiceConnectionStatusPtrType = reflect.TypeOf((**PrivateLinkServiceConnectionStatus)(nil)).Elem()

type PrivateLinkServiceConnectionStatusPtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput
	ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatusPtrOutput
}

type privateLinkServiceConnectionStatusPtr string

func PrivateLinkServiceConnectionStatusPtr(v string) PrivateLinkServiceConnectionStatusPtrInput {
	return (*privateLinkServiceConnectionStatusPtr)(&v)
}

func (*privateLinkServiceConnectionStatusPtr) ElementType() reflect.Type {
	return privateLinkServiceConnectionStatusPtrType
}

func (in *privateLinkServiceConnectionStatusPtr) ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput {
	return pulumi.ToOutput(in).(PrivateLinkServiceConnectionStatusPtrOutput)
}

func (in *privateLinkServiceConnectionStatusPtr) ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrivateLinkServiceConnectionStatusPtrOutput)
}

type SkuName string

const (
	SkuNameFree                  = SkuName("free")
	SkuNameBasic                 = SkuName("basic")
	SkuNameStandard              = SkuName("standard")
	SkuNameStandard2             = SkuName("standard2")
	SkuNameStandard3             = SkuName("standard3")
	SkuName_Storage_optimized_l1 = SkuName("storage_optimized_l1")
	SkuName_Storage_optimized_l2 = SkuName("storage_optimized_l2")
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

func init() {
	pulumi.RegisterOutputType(EndpointAccessOutput{})
	pulumi.RegisterOutputType(EndpointAccessPtrOutput{})
	pulumi.RegisterOutputType(HostingModeOutput{})
	pulumi.RegisterOutputType(HostingModePtrOutput{})
	pulumi.RegisterOutputType(IdentityTypeOutput{})
	pulumi.RegisterOutputType(IdentityTypePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
}
