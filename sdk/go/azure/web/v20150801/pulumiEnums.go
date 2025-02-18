


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessControlEntryAction string

const (
	AccessControlEntryActionPermit = AccessControlEntryAction("Permit")
	AccessControlEntryActionDeny   = AccessControlEntryAction("Deny")
)

func (AccessControlEntryAction) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessControlEntryAction)(nil)).Elem()
}

func (e AccessControlEntryAction) ToAccessControlEntryActionOutput() AccessControlEntryActionOutput {
	return pulumi.ToOutput(e).(AccessControlEntryActionOutput)
}

func (e AccessControlEntryAction) ToAccessControlEntryActionOutputWithContext(ctx context.Context) AccessControlEntryActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessControlEntryActionOutput)
}

func (e AccessControlEntryAction) ToAccessControlEntryActionPtrOutput() AccessControlEntryActionPtrOutput {
	return e.ToAccessControlEntryActionPtrOutputWithContext(context.Background())
}

func (e AccessControlEntryAction) ToAccessControlEntryActionPtrOutputWithContext(ctx context.Context) AccessControlEntryActionPtrOutput {
	return AccessControlEntryAction(e).ToAccessControlEntryActionOutputWithContext(ctx).ToAccessControlEntryActionPtrOutputWithContext(ctx)
}

func (e AccessControlEntryAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessControlEntryAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessControlEntryAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessControlEntryAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessControlEntryActionOutput struct{ *pulumi.OutputState }

func (AccessControlEntryActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessControlEntryAction)(nil)).Elem()
}

func (o AccessControlEntryActionOutput) ToAccessControlEntryActionOutput() AccessControlEntryActionOutput {
	return o
}

func (o AccessControlEntryActionOutput) ToAccessControlEntryActionOutputWithContext(ctx context.Context) AccessControlEntryActionOutput {
	return o
}

func (o AccessControlEntryActionOutput) ToAccessControlEntryActionPtrOutput() AccessControlEntryActionPtrOutput {
	return o.ToAccessControlEntryActionPtrOutputWithContext(context.Background())
}

func (o AccessControlEntryActionOutput) ToAccessControlEntryActionPtrOutputWithContext(ctx context.Context) AccessControlEntryActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessControlEntryAction) *AccessControlEntryAction {
		return &v
	}).(AccessControlEntryActionPtrOutput)
}

func (o AccessControlEntryActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessControlEntryActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessControlEntryAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessControlEntryActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessControlEntryActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessControlEntryAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessControlEntryActionPtrOutput struct{ *pulumi.OutputState }

func (AccessControlEntryActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessControlEntryAction)(nil)).Elem()
}

func (o AccessControlEntryActionPtrOutput) ToAccessControlEntryActionPtrOutput() AccessControlEntryActionPtrOutput {
	return o
}

func (o AccessControlEntryActionPtrOutput) ToAccessControlEntryActionPtrOutputWithContext(ctx context.Context) AccessControlEntryActionPtrOutput {
	return o
}

func (o AccessControlEntryActionPtrOutput) Elem() AccessControlEntryActionOutput {
	return o.ApplyT(func(v *AccessControlEntryAction) AccessControlEntryAction {
		if v != nil {
			return *v
		}
		var ret AccessControlEntryAction
		return ret
	}).(AccessControlEntryActionOutput)
}

func (o AccessControlEntryActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessControlEntryActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessControlEntryAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AccessControlEntryActionInput interface {
	pulumi.Input

	ToAccessControlEntryActionOutput() AccessControlEntryActionOutput
	ToAccessControlEntryActionOutputWithContext(context.Context) AccessControlEntryActionOutput
}

var accessControlEntryActionPtrType = reflect.TypeOf((**AccessControlEntryAction)(nil)).Elem()

type AccessControlEntryActionPtrInput interface {
	pulumi.Input

	ToAccessControlEntryActionPtrOutput() AccessControlEntryActionPtrOutput
	ToAccessControlEntryActionPtrOutputWithContext(context.Context) AccessControlEntryActionPtrOutput
}

type accessControlEntryActionPtr string

func AccessControlEntryActionPtr(v string) AccessControlEntryActionPtrInput {
	return (*accessControlEntryActionPtr)(&v)
}

func (*accessControlEntryActionPtr) ElementType() reflect.Type {
	return accessControlEntryActionPtrType
}

func (in *accessControlEntryActionPtr) ToAccessControlEntryActionPtrOutput() AccessControlEntryActionPtrOutput {
	return pulumi.ToOutput(in).(AccessControlEntryActionPtrOutput)
}

func (in *accessControlEntryActionPtr) ToAccessControlEntryActionPtrOutputWithContext(ctx context.Context) AccessControlEntryActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessControlEntryActionPtrOutput)
}

type AutoHealActionType string

const (
	AutoHealActionTypeRecycle      = AutoHealActionType("Recycle")
	AutoHealActionTypeLogEvent     = AutoHealActionType("LogEvent")
	AutoHealActionTypeCustomAction = AutoHealActionType("CustomAction")
)

func (AutoHealActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealActionType)(nil)).Elem()
}

func (e AutoHealActionType) ToAutoHealActionTypeOutput() AutoHealActionTypeOutput {
	return pulumi.ToOutput(e).(AutoHealActionTypeOutput)
}

func (e AutoHealActionType) ToAutoHealActionTypeOutputWithContext(ctx context.Context) AutoHealActionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutoHealActionTypeOutput)
}

func (e AutoHealActionType) ToAutoHealActionTypePtrOutput() AutoHealActionTypePtrOutput {
	return e.ToAutoHealActionTypePtrOutputWithContext(context.Background())
}

func (e AutoHealActionType) ToAutoHealActionTypePtrOutputWithContext(ctx context.Context) AutoHealActionTypePtrOutput {
	return AutoHealActionType(e).ToAutoHealActionTypeOutputWithContext(ctx).ToAutoHealActionTypePtrOutputWithContext(ctx)
}

func (e AutoHealActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoHealActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoHealActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutoHealActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutoHealActionTypeOutput struct{ *pulumi.OutputState }

func (AutoHealActionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealActionType)(nil)).Elem()
}

func (o AutoHealActionTypeOutput) ToAutoHealActionTypeOutput() AutoHealActionTypeOutput {
	return o
}

func (o AutoHealActionTypeOutput) ToAutoHealActionTypeOutputWithContext(ctx context.Context) AutoHealActionTypeOutput {
	return o
}

func (o AutoHealActionTypeOutput) ToAutoHealActionTypePtrOutput() AutoHealActionTypePtrOutput {
	return o.ToAutoHealActionTypePtrOutputWithContext(context.Background())
}

func (o AutoHealActionTypeOutput) ToAutoHealActionTypePtrOutputWithContext(ctx context.Context) AutoHealActionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealActionType) *AutoHealActionType {
		return &v
	}).(AutoHealActionTypePtrOutput)
}

func (o AutoHealActionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutoHealActionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoHealActionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutoHealActionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoHealActionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoHealActionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutoHealActionTypePtrOutput struct{ *pulumi.OutputState }

func (AutoHealActionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealActionType)(nil)).Elem()
}

func (o AutoHealActionTypePtrOutput) ToAutoHealActionTypePtrOutput() AutoHealActionTypePtrOutput {
	return o
}

func (o AutoHealActionTypePtrOutput) ToAutoHealActionTypePtrOutputWithContext(ctx context.Context) AutoHealActionTypePtrOutput {
	return o
}

func (o AutoHealActionTypePtrOutput) Elem() AutoHealActionTypeOutput {
	return o.ApplyT(func(v *AutoHealActionType) AutoHealActionType {
		if v != nil {
			return *v
		}
		var ret AutoHealActionType
		return ret
	}).(AutoHealActionTypeOutput)
}

func (o AutoHealActionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoHealActionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutoHealActionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AutoHealActionTypeInput interface {
	pulumi.Input

	ToAutoHealActionTypeOutput() AutoHealActionTypeOutput
	ToAutoHealActionTypeOutputWithContext(context.Context) AutoHealActionTypeOutput
}

var autoHealActionTypePtrType = reflect.TypeOf((**AutoHealActionType)(nil)).Elem()

type AutoHealActionTypePtrInput interface {
	pulumi.Input

	ToAutoHealActionTypePtrOutput() AutoHealActionTypePtrOutput
	ToAutoHealActionTypePtrOutputWithContext(context.Context) AutoHealActionTypePtrOutput
}

type autoHealActionTypePtr string

func AutoHealActionTypePtr(v string) AutoHealActionTypePtrInput {
	return (*autoHealActionTypePtr)(&v)
}

func (*autoHealActionTypePtr) ElementType() reflect.Type {
	return autoHealActionTypePtrType
}

func (in *autoHealActionTypePtr) ToAutoHealActionTypePtrOutput() AutoHealActionTypePtrOutput {
	return pulumi.ToOutput(in).(AutoHealActionTypePtrOutput)
}

func (in *autoHealActionTypePtr) ToAutoHealActionTypePtrOutputWithContext(ctx context.Context) AutoHealActionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutoHealActionTypePtrOutput)
}

type AzureResourceType string

const (
	AzureResourceTypeWebsite        = AzureResourceType("Website")
	AzureResourceTypeTrafficManager = AzureResourceType("TrafficManager")
)

func (AzureResourceType) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceType)(nil)).Elem()
}

func (e AzureResourceType) ToAzureResourceTypeOutput() AzureResourceTypeOutput {
	return pulumi.ToOutput(e).(AzureResourceTypeOutput)
}

func (e AzureResourceType) ToAzureResourceTypeOutputWithContext(ctx context.Context) AzureResourceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureResourceTypeOutput)
}

func (e AzureResourceType) ToAzureResourceTypePtrOutput() AzureResourceTypePtrOutput {
	return e.ToAzureResourceTypePtrOutputWithContext(context.Background())
}

func (e AzureResourceType) ToAzureResourceTypePtrOutputWithContext(ctx context.Context) AzureResourceTypePtrOutput {
	return AzureResourceType(e).ToAzureResourceTypeOutputWithContext(ctx).ToAzureResourceTypePtrOutputWithContext(ctx)
}

func (e AzureResourceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureResourceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureResourceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureResourceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureResourceTypeOutput struct{ *pulumi.OutputState }

func (AzureResourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceType)(nil)).Elem()
}

func (o AzureResourceTypeOutput) ToAzureResourceTypeOutput() AzureResourceTypeOutput {
	return o
}

func (o AzureResourceTypeOutput) ToAzureResourceTypeOutputWithContext(ctx context.Context) AzureResourceTypeOutput {
	return o
}

func (o AzureResourceTypeOutput) ToAzureResourceTypePtrOutput() AzureResourceTypePtrOutput {
	return o.ToAzureResourceTypePtrOutputWithContext(context.Background())
}

func (o AzureResourceTypeOutput) ToAzureResourceTypePtrOutputWithContext(ctx context.Context) AzureResourceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureResourceType) *AzureResourceType {
		return &v
	}).(AzureResourceTypePtrOutput)
}

func (o AzureResourceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureResourceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureResourceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureResourceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureResourceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureResourceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureResourceTypePtrOutput struct{ *pulumi.OutputState }

func (AzureResourceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureResourceType)(nil)).Elem()
}

func (o AzureResourceTypePtrOutput) ToAzureResourceTypePtrOutput() AzureResourceTypePtrOutput {
	return o
}

func (o AzureResourceTypePtrOutput) ToAzureResourceTypePtrOutputWithContext(ctx context.Context) AzureResourceTypePtrOutput {
	return o
}

func (o AzureResourceTypePtrOutput) Elem() AzureResourceTypeOutput {
	return o.ApplyT(func(v *AzureResourceType) AzureResourceType {
		if v != nil {
			return *v
		}
		var ret AzureResourceType
		return ret
	}).(AzureResourceTypeOutput)
}

func (o AzureResourceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureResourceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureResourceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureResourceTypeInput interface {
	pulumi.Input

	ToAzureResourceTypeOutput() AzureResourceTypeOutput
	ToAzureResourceTypeOutputWithContext(context.Context) AzureResourceTypeOutput
}

var azureResourceTypePtrType = reflect.TypeOf((**AzureResourceType)(nil)).Elem()

type AzureResourceTypePtrInput interface {
	pulumi.Input

	ToAzureResourceTypePtrOutput() AzureResourceTypePtrOutput
	ToAzureResourceTypePtrOutputWithContext(context.Context) AzureResourceTypePtrOutput
}

type azureResourceTypePtr string

func AzureResourceTypePtr(v string) AzureResourceTypePtrInput {
	return (*azureResourceTypePtr)(&v)
}

func (*azureResourceTypePtr) ElementType() reflect.Type {
	return azureResourceTypePtrType
}

func (in *azureResourceTypePtr) ToAzureResourceTypePtrOutput() AzureResourceTypePtrOutput {
	return pulumi.ToOutput(in).(AzureResourceTypePtrOutput)
}

func (in *azureResourceTypePtr) ToAzureResourceTypePtrOutputWithContext(ctx context.Context) AzureResourceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureResourceTypePtrOutput)
}

type BackupRestoreOperationType string

const (
	BackupRestoreOperationTypeDefault    = BackupRestoreOperationType("Default")
	BackupRestoreOperationTypeClone      = BackupRestoreOperationType("Clone")
	BackupRestoreOperationTypeRelocation = BackupRestoreOperationType("Relocation")
)

type BuiltInAuthenticationProvider string

const (
	BuiltInAuthenticationProviderAzureActiveDirectory = BuiltInAuthenticationProvider("AzureActiveDirectory")
	BuiltInAuthenticationProviderFacebook             = BuiltInAuthenticationProvider("Facebook")
	BuiltInAuthenticationProviderGoogle               = BuiltInAuthenticationProvider("Google")
	BuiltInAuthenticationProviderMicrosoftAccount     = BuiltInAuthenticationProvider("MicrosoftAccount")
	BuiltInAuthenticationProviderTwitter              = BuiltInAuthenticationProvider("Twitter")
)

func (BuiltInAuthenticationProvider) ElementType() reflect.Type {
	return reflect.TypeOf((*BuiltInAuthenticationProvider)(nil)).Elem()
}

func (e BuiltInAuthenticationProvider) ToBuiltInAuthenticationProviderOutput() BuiltInAuthenticationProviderOutput {
	return pulumi.ToOutput(e).(BuiltInAuthenticationProviderOutput)
}

func (e BuiltInAuthenticationProvider) ToBuiltInAuthenticationProviderOutputWithContext(ctx context.Context) BuiltInAuthenticationProviderOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BuiltInAuthenticationProviderOutput)
}

func (e BuiltInAuthenticationProvider) ToBuiltInAuthenticationProviderPtrOutput() BuiltInAuthenticationProviderPtrOutput {
	return e.ToBuiltInAuthenticationProviderPtrOutputWithContext(context.Background())
}

func (e BuiltInAuthenticationProvider) ToBuiltInAuthenticationProviderPtrOutputWithContext(ctx context.Context) BuiltInAuthenticationProviderPtrOutput {
	return BuiltInAuthenticationProvider(e).ToBuiltInAuthenticationProviderOutputWithContext(ctx).ToBuiltInAuthenticationProviderPtrOutputWithContext(ctx)
}

func (e BuiltInAuthenticationProvider) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BuiltInAuthenticationProvider) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BuiltInAuthenticationProvider) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BuiltInAuthenticationProvider) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BuiltInAuthenticationProviderOutput struct{ *pulumi.OutputState }

func (BuiltInAuthenticationProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuiltInAuthenticationProvider)(nil)).Elem()
}

func (o BuiltInAuthenticationProviderOutput) ToBuiltInAuthenticationProviderOutput() BuiltInAuthenticationProviderOutput {
	return o
}

func (o BuiltInAuthenticationProviderOutput) ToBuiltInAuthenticationProviderOutputWithContext(ctx context.Context) BuiltInAuthenticationProviderOutput {
	return o
}

func (o BuiltInAuthenticationProviderOutput) ToBuiltInAuthenticationProviderPtrOutput() BuiltInAuthenticationProviderPtrOutput {
	return o.ToBuiltInAuthenticationProviderPtrOutputWithContext(context.Background())
}

func (o BuiltInAuthenticationProviderOutput) ToBuiltInAuthenticationProviderPtrOutputWithContext(ctx context.Context) BuiltInAuthenticationProviderPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BuiltInAuthenticationProvider) *BuiltInAuthenticationProvider {
		return &v
	}).(BuiltInAuthenticationProviderPtrOutput)
}

func (o BuiltInAuthenticationProviderOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BuiltInAuthenticationProviderOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BuiltInAuthenticationProvider) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BuiltInAuthenticationProviderOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BuiltInAuthenticationProviderOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BuiltInAuthenticationProvider) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BuiltInAuthenticationProviderPtrOutput struct{ *pulumi.OutputState }

func (BuiltInAuthenticationProviderPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BuiltInAuthenticationProvider)(nil)).Elem()
}

func (o BuiltInAuthenticationProviderPtrOutput) ToBuiltInAuthenticationProviderPtrOutput() BuiltInAuthenticationProviderPtrOutput {
	return o
}

func (o BuiltInAuthenticationProviderPtrOutput) ToBuiltInAuthenticationProviderPtrOutputWithContext(ctx context.Context) BuiltInAuthenticationProviderPtrOutput {
	return o
}

func (o BuiltInAuthenticationProviderPtrOutput) Elem() BuiltInAuthenticationProviderOutput {
	return o.ApplyT(func(v *BuiltInAuthenticationProvider) BuiltInAuthenticationProvider {
		if v != nil {
			return *v
		}
		var ret BuiltInAuthenticationProvider
		return ret
	}).(BuiltInAuthenticationProviderOutput)
}

func (o BuiltInAuthenticationProviderPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BuiltInAuthenticationProviderPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BuiltInAuthenticationProvider) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BuiltInAuthenticationProviderInput interface {
	pulumi.Input

	ToBuiltInAuthenticationProviderOutput() BuiltInAuthenticationProviderOutput
	ToBuiltInAuthenticationProviderOutputWithContext(context.Context) BuiltInAuthenticationProviderOutput
}

var builtInAuthenticationProviderPtrType = reflect.TypeOf((**BuiltInAuthenticationProvider)(nil)).Elem()

type BuiltInAuthenticationProviderPtrInput interface {
	pulumi.Input

	ToBuiltInAuthenticationProviderPtrOutput() BuiltInAuthenticationProviderPtrOutput
	ToBuiltInAuthenticationProviderPtrOutputWithContext(context.Context) BuiltInAuthenticationProviderPtrOutput
}

type builtInAuthenticationProviderPtr string

func BuiltInAuthenticationProviderPtr(v string) BuiltInAuthenticationProviderPtrInput {
	return (*builtInAuthenticationProviderPtr)(&v)
}

func (*builtInAuthenticationProviderPtr) ElementType() reflect.Type {
	return builtInAuthenticationProviderPtrType
}

func (in *builtInAuthenticationProviderPtr) ToBuiltInAuthenticationProviderPtrOutput() BuiltInAuthenticationProviderPtrOutput {
	return pulumi.ToOutput(in).(BuiltInAuthenticationProviderPtrOutput)
}

func (in *builtInAuthenticationProviderPtr) ToBuiltInAuthenticationProviderPtrOutputWithContext(ctx context.Context) BuiltInAuthenticationProviderPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BuiltInAuthenticationProviderPtrOutput)
}

type ComputeModeOptions string

const (
	ComputeModeOptionsShared    = ComputeModeOptions("Shared")
	ComputeModeOptionsDedicated = ComputeModeOptions("Dedicated")
	ComputeModeOptionsDynamic   = ComputeModeOptions("Dynamic")
)

func (ComputeModeOptions) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeModeOptions)(nil)).Elem()
}

func (e ComputeModeOptions) ToComputeModeOptionsOutput() ComputeModeOptionsOutput {
	return pulumi.ToOutput(e).(ComputeModeOptionsOutput)
}

func (e ComputeModeOptions) ToComputeModeOptionsOutputWithContext(ctx context.Context) ComputeModeOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ComputeModeOptionsOutput)
}

func (e ComputeModeOptions) ToComputeModeOptionsPtrOutput() ComputeModeOptionsPtrOutput {
	return e.ToComputeModeOptionsPtrOutputWithContext(context.Background())
}

func (e ComputeModeOptions) ToComputeModeOptionsPtrOutputWithContext(ctx context.Context) ComputeModeOptionsPtrOutput {
	return ComputeModeOptions(e).ToComputeModeOptionsOutputWithContext(ctx).ToComputeModeOptionsPtrOutputWithContext(ctx)
}

func (e ComputeModeOptions) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeModeOptions) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeModeOptions) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ComputeModeOptions) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ComputeModeOptionsOutput struct{ *pulumi.OutputState }

func (ComputeModeOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeModeOptions)(nil)).Elem()
}

func (o ComputeModeOptionsOutput) ToComputeModeOptionsOutput() ComputeModeOptionsOutput {
	return o
}

func (o ComputeModeOptionsOutput) ToComputeModeOptionsOutputWithContext(ctx context.Context) ComputeModeOptionsOutput {
	return o
}

func (o ComputeModeOptionsOutput) ToComputeModeOptionsPtrOutput() ComputeModeOptionsPtrOutput {
	return o.ToComputeModeOptionsPtrOutputWithContext(context.Background())
}

func (o ComputeModeOptionsOutput) ToComputeModeOptionsPtrOutputWithContext(ctx context.Context) ComputeModeOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeModeOptions) *ComputeModeOptions {
		return &v
	}).(ComputeModeOptionsPtrOutput)
}

func (o ComputeModeOptionsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ComputeModeOptionsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComputeModeOptions) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ComputeModeOptionsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComputeModeOptionsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComputeModeOptions) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ComputeModeOptionsPtrOutput struct{ *pulumi.OutputState }

func (ComputeModeOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeModeOptions)(nil)).Elem()
}

func (o ComputeModeOptionsPtrOutput) ToComputeModeOptionsPtrOutput() ComputeModeOptionsPtrOutput {
	return o
}

func (o ComputeModeOptionsPtrOutput) ToComputeModeOptionsPtrOutputWithContext(ctx context.Context) ComputeModeOptionsPtrOutput {
	return o
}

func (o ComputeModeOptionsPtrOutput) Elem() ComputeModeOptionsOutput {
	return o.ApplyT(func(v *ComputeModeOptions) ComputeModeOptions {
		if v != nil {
			return *v
		}
		var ret ComputeModeOptions
		return ret
	}).(ComputeModeOptionsOutput)
}

func (o ComputeModeOptionsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComputeModeOptionsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ComputeModeOptions) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ComputeModeOptionsInput interface {
	pulumi.Input

	ToComputeModeOptionsOutput() ComputeModeOptionsOutput
	ToComputeModeOptionsOutputWithContext(context.Context) ComputeModeOptionsOutput
}

var computeModeOptionsPtrType = reflect.TypeOf((**ComputeModeOptions)(nil)).Elem()

type ComputeModeOptionsPtrInput interface {
	pulumi.Input

	ToComputeModeOptionsPtrOutput() ComputeModeOptionsPtrOutput
	ToComputeModeOptionsPtrOutputWithContext(context.Context) ComputeModeOptionsPtrOutput
}

type computeModeOptionsPtr string

func ComputeModeOptionsPtr(v string) ComputeModeOptionsPtrInput {
	return (*computeModeOptionsPtr)(&v)
}

func (*computeModeOptionsPtr) ElementType() reflect.Type {
	return computeModeOptionsPtrType
}

func (in *computeModeOptionsPtr) ToComputeModeOptionsPtrOutput() ComputeModeOptionsPtrOutput {
	return pulumi.ToOutput(in).(ComputeModeOptionsPtrOutput)
}

func (in *computeModeOptionsPtr) ToComputeModeOptionsPtrOutputWithContext(ctx context.Context) ComputeModeOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ComputeModeOptionsPtrOutput)
}

type CustomHostNameDnsRecordType string

const (
	CustomHostNameDnsRecordTypeCName = CustomHostNameDnsRecordType("CName")
	CustomHostNameDnsRecordTypeA     = CustomHostNameDnsRecordType("A")
)

func (CustomHostNameDnsRecordType) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomHostNameDnsRecordType)(nil)).Elem()
}

func (e CustomHostNameDnsRecordType) ToCustomHostNameDnsRecordTypeOutput() CustomHostNameDnsRecordTypeOutput {
	return pulumi.ToOutput(e).(CustomHostNameDnsRecordTypeOutput)
}

func (e CustomHostNameDnsRecordType) ToCustomHostNameDnsRecordTypeOutputWithContext(ctx context.Context) CustomHostNameDnsRecordTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CustomHostNameDnsRecordTypeOutput)
}

func (e CustomHostNameDnsRecordType) ToCustomHostNameDnsRecordTypePtrOutput() CustomHostNameDnsRecordTypePtrOutput {
	return e.ToCustomHostNameDnsRecordTypePtrOutputWithContext(context.Background())
}

func (e CustomHostNameDnsRecordType) ToCustomHostNameDnsRecordTypePtrOutputWithContext(ctx context.Context) CustomHostNameDnsRecordTypePtrOutput {
	return CustomHostNameDnsRecordType(e).ToCustomHostNameDnsRecordTypeOutputWithContext(ctx).ToCustomHostNameDnsRecordTypePtrOutputWithContext(ctx)
}

func (e CustomHostNameDnsRecordType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CustomHostNameDnsRecordType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CustomHostNameDnsRecordType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CustomHostNameDnsRecordType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CustomHostNameDnsRecordTypeOutput struct{ *pulumi.OutputState }

func (CustomHostNameDnsRecordTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomHostNameDnsRecordType)(nil)).Elem()
}

func (o CustomHostNameDnsRecordTypeOutput) ToCustomHostNameDnsRecordTypeOutput() CustomHostNameDnsRecordTypeOutput {
	return o
}

func (o CustomHostNameDnsRecordTypeOutput) ToCustomHostNameDnsRecordTypeOutputWithContext(ctx context.Context) CustomHostNameDnsRecordTypeOutput {
	return o
}

func (o CustomHostNameDnsRecordTypeOutput) ToCustomHostNameDnsRecordTypePtrOutput() CustomHostNameDnsRecordTypePtrOutput {
	return o.ToCustomHostNameDnsRecordTypePtrOutputWithContext(context.Background())
}

func (o CustomHostNameDnsRecordTypeOutput) ToCustomHostNameDnsRecordTypePtrOutputWithContext(ctx context.Context) CustomHostNameDnsRecordTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomHostNameDnsRecordType) *CustomHostNameDnsRecordType {
		return &v
	}).(CustomHostNameDnsRecordTypePtrOutput)
}

func (o CustomHostNameDnsRecordTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CustomHostNameDnsRecordTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CustomHostNameDnsRecordType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CustomHostNameDnsRecordTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CustomHostNameDnsRecordTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CustomHostNameDnsRecordType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CustomHostNameDnsRecordTypePtrOutput struct{ *pulumi.OutputState }

func (CustomHostNameDnsRecordTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomHostNameDnsRecordType)(nil)).Elem()
}

func (o CustomHostNameDnsRecordTypePtrOutput) ToCustomHostNameDnsRecordTypePtrOutput() CustomHostNameDnsRecordTypePtrOutput {
	return o
}

func (o CustomHostNameDnsRecordTypePtrOutput) ToCustomHostNameDnsRecordTypePtrOutputWithContext(ctx context.Context) CustomHostNameDnsRecordTypePtrOutput {
	return o
}

func (o CustomHostNameDnsRecordTypePtrOutput) Elem() CustomHostNameDnsRecordTypeOutput {
	return o.ApplyT(func(v *CustomHostNameDnsRecordType) CustomHostNameDnsRecordType {
		if v != nil {
			return *v
		}
		var ret CustomHostNameDnsRecordType
		return ret
	}).(CustomHostNameDnsRecordTypeOutput)
}

func (o CustomHostNameDnsRecordTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CustomHostNameDnsRecordTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CustomHostNameDnsRecordType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CustomHostNameDnsRecordTypeInput interface {
	pulumi.Input

	ToCustomHostNameDnsRecordTypeOutput() CustomHostNameDnsRecordTypeOutput
	ToCustomHostNameDnsRecordTypeOutputWithContext(context.Context) CustomHostNameDnsRecordTypeOutput
}

var customHostNameDnsRecordTypePtrType = reflect.TypeOf((**CustomHostNameDnsRecordType)(nil)).Elem()

type CustomHostNameDnsRecordTypePtrInput interface {
	pulumi.Input

	ToCustomHostNameDnsRecordTypePtrOutput() CustomHostNameDnsRecordTypePtrOutput
	ToCustomHostNameDnsRecordTypePtrOutputWithContext(context.Context) CustomHostNameDnsRecordTypePtrOutput
}

type customHostNameDnsRecordTypePtr string

func CustomHostNameDnsRecordTypePtr(v string) CustomHostNameDnsRecordTypePtrInput {
	return (*customHostNameDnsRecordTypePtr)(&v)
}

func (*customHostNameDnsRecordTypePtr) ElementType() reflect.Type {
	return customHostNameDnsRecordTypePtrType
}

func (in *customHostNameDnsRecordTypePtr) ToCustomHostNameDnsRecordTypePtrOutput() CustomHostNameDnsRecordTypePtrOutput {
	return pulumi.ToOutput(in).(CustomHostNameDnsRecordTypePtrOutput)
}

func (in *customHostNameDnsRecordTypePtr) ToCustomHostNameDnsRecordTypePtrOutputWithContext(ctx context.Context) CustomHostNameDnsRecordTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CustomHostNameDnsRecordTypePtrOutput)
}

type DatabaseServerType string

const (
	DatabaseServerTypeMySql     = DatabaseServerType("MySql")
	DatabaseServerTypeSQLServer = DatabaseServerType("SQLServer")
	DatabaseServerTypeSQLAzure  = DatabaseServerType("SQLAzure")
	DatabaseServerTypeCustom    = DatabaseServerType("Custom")
)

func (DatabaseServerType) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseServerType)(nil)).Elem()
}

func (e DatabaseServerType) ToDatabaseServerTypeOutput() DatabaseServerTypeOutput {
	return pulumi.ToOutput(e).(DatabaseServerTypeOutput)
}

func (e DatabaseServerType) ToDatabaseServerTypeOutputWithContext(ctx context.Context) DatabaseServerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DatabaseServerTypeOutput)
}

func (e DatabaseServerType) ToDatabaseServerTypePtrOutput() DatabaseServerTypePtrOutput {
	return e.ToDatabaseServerTypePtrOutputWithContext(context.Background())
}

func (e DatabaseServerType) ToDatabaseServerTypePtrOutputWithContext(ctx context.Context) DatabaseServerTypePtrOutput {
	return DatabaseServerType(e).ToDatabaseServerTypeOutputWithContext(ctx).ToDatabaseServerTypePtrOutputWithContext(ctx)
}

func (e DatabaseServerType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseServerType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseServerType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DatabaseServerType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DatabaseServerTypeOutput struct{ *pulumi.OutputState }

func (DatabaseServerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseServerType)(nil)).Elem()
}

func (o DatabaseServerTypeOutput) ToDatabaseServerTypeOutput() DatabaseServerTypeOutput {
	return o
}

func (o DatabaseServerTypeOutput) ToDatabaseServerTypeOutputWithContext(ctx context.Context) DatabaseServerTypeOutput {
	return o
}

func (o DatabaseServerTypeOutput) ToDatabaseServerTypePtrOutput() DatabaseServerTypePtrOutput {
	return o.ToDatabaseServerTypePtrOutputWithContext(context.Background())
}

func (o DatabaseServerTypeOutput) ToDatabaseServerTypePtrOutputWithContext(ctx context.Context) DatabaseServerTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabaseServerType) *DatabaseServerType {
		return &v
	}).(DatabaseServerTypePtrOutput)
}

func (o DatabaseServerTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DatabaseServerTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseServerType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DatabaseServerTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseServerTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseServerType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DatabaseServerTypePtrOutput struct{ *pulumi.OutputState }

func (DatabaseServerTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseServerType)(nil)).Elem()
}

func (o DatabaseServerTypePtrOutput) ToDatabaseServerTypePtrOutput() DatabaseServerTypePtrOutput {
	return o
}

func (o DatabaseServerTypePtrOutput) ToDatabaseServerTypePtrOutputWithContext(ctx context.Context) DatabaseServerTypePtrOutput {
	return o
}

func (o DatabaseServerTypePtrOutput) Elem() DatabaseServerTypeOutput {
	return o.ApplyT(func(v *DatabaseServerType) DatabaseServerType {
		if v != nil {
			return *v
		}
		var ret DatabaseServerType
		return ret
	}).(DatabaseServerTypeOutput)
}

func (o DatabaseServerTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseServerTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DatabaseServerType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DatabaseServerTypeInput interface {
	pulumi.Input

	ToDatabaseServerTypeOutput() DatabaseServerTypeOutput
	ToDatabaseServerTypeOutputWithContext(context.Context) DatabaseServerTypeOutput
}

var databaseServerTypePtrType = reflect.TypeOf((**DatabaseServerType)(nil)).Elem()

type DatabaseServerTypePtrInput interface {
	pulumi.Input

	ToDatabaseServerTypePtrOutput() DatabaseServerTypePtrOutput
	ToDatabaseServerTypePtrOutputWithContext(context.Context) DatabaseServerTypePtrOutput
}

type databaseServerTypePtr string

func DatabaseServerTypePtr(v string) DatabaseServerTypePtrInput {
	return (*databaseServerTypePtr)(&v)
}

func (*databaseServerTypePtr) ElementType() reflect.Type {
	return databaseServerTypePtrType
}

func (in *databaseServerTypePtr) ToDatabaseServerTypePtrOutput() DatabaseServerTypePtrOutput {
	return pulumi.ToOutput(in).(DatabaseServerTypePtrOutput)
}

func (in *databaseServerTypePtr) ToDatabaseServerTypePtrOutputWithContext(ctx context.Context) DatabaseServerTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DatabaseServerTypePtrOutput)
}

type FrequencyUnit string

const (
	FrequencyUnitDay  = FrequencyUnit("Day")
	FrequencyUnitHour = FrequencyUnit("Hour")
)

func (FrequencyUnit) ElementType() reflect.Type {
	return reflect.TypeOf((*FrequencyUnit)(nil)).Elem()
}

func (e FrequencyUnit) ToFrequencyUnitOutput() FrequencyUnitOutput {
	return pulumi.ToOutput(e).(FrequencyUnitOutput)
}

func (e FrequencyUnit) ToFrequencyUnitOutputWithContext(ctx context.Context) FrequencyUnitOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FrequencyUnitOutput)
}

func (e FrequencyUnit) ToFrequencyUnitPtrOutput() FrequencyUnitPtrOutput {
	return e.ToFrequencyUnitPtrOutputWithContext(context.Background())
}

func (e FrequencyUnit) ToFrequencyUnitPtrOutputWithContext(ctx context.Context) FrequencyUnitPtrOutput {
	return FrequencyUnit(e).ToFrequencyUnitOutputWithContext(ctx).ToFrequencyUnitPtrOutputWithContext(ctx)
}

func (e FrequencyUnit) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FrequencyUnit) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FrequencyUnit) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FrequencyUnit) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FrequencyUnitOutput struct{ *pulumi.OutputState }

func (FrequencyUnitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrequencyUnit)(nil)).Elem()
}

func (o FrequencyUnitOutput) ToFrequencyUnitOutput() FrequencyUnitOutput {
	return o
}

func (o FrequencyUnitOutput) ToFrequencyUnitOutputWithContext(ctx context.Context) FrequencyUnitOutput {
	return o
}

func (o FrequencyUnitOutput) ToFrequencyUnitPtrOutput() FrequencyUnitPtrOutput {
	return o.ToFrequencyUnitPtrOutputWithContext(context.Background())
}

func (o FrequencyUnitOutput) ToFrequencyUnitPtrOutputWithContext(ctx context.Context) FrequencyUnitPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FrequencyUnit) *FrequencyUnit {
		return &v
	}).(FrequencyUnitPtrOutput)
}

func (o FrequencyUnitOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FrequencyUnitOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FrequencyUnit) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FrequencyUnitOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FrequencyUnitOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FrequencyUnit) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FrequencyUnitPtrOutput struct{ *pulumi.OutputState }

func (FrequencyUnitPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrequencyUnit)(nil)).Elem()
}

func (o FrequencyUnitPtrOutput) ToFrequencyUnitPtrOutput() FrequencyUnitPtrOutput {
	return o
}

func (o FrequencyUnitPtrOutput) ToFrequencyUnitPtrOutputWithContext(ctx context.Context) FrequencyUnitPtrOutput {
	return o
}

func (o FrequencyUnitPtrOutput) Elem() FrequencyUnitOutput {
	return o.ApplyT(func(v *FrequencyUnit) FrequencyUnit {
		if v != nil {
			return *v
		}
		var ret FrequencyUnit
		return ret
	}).(FrequencyUnitOutput)
}

func (o FrequencyUnitPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FrequencyUnitPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FrequencyUnit) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FrequencyUnitInput interface {
	pulumi.Input

	ToFrequencyUnitOutput() FrequencyUnitOutput
	ToFrequencyUnitOutputWithContext(context.Context) FrequencyUnitOutput
}

var frequencyUnitPtrType = reflect.TypeOf((**FrequencyUnit)(nil)).Elem()

type FrequencyUnitPtrInput interface {
	pulumi.Input

	ToFrequencyUnitPtrOutput() FrequencyUnitPtrOutput
	ToFrequencyUnitPtrOutputWithContext(context.Context) FrequencyUnitPtrOutput
}

type frequencyUnitPtr string

func FrequencyUnitPtr(v string) FrequencyUnitPtrInput {
	return (*frequencyUnitPtr)(&v)
}

func (*frequencyUnitPtr) ElementType() reflect.Type {
	return frequencyUnitPtrType
}

func (in *frequencyUnitPtr) ToFrequencyUnitPtrOutput() FrequencyUnitPtrOutput {
	return pulumi.ToOutput(in).(FrequencyUnitPtrOutput)
}

func (in *frequencyUnitPtr) ToFrequencyUnitPtrOutputWithContext(ctx context.Context) FrequencyUnitPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FrequencyUnitPtrOutput)
}

type HostNameType string

const (
	HostNameTypeVerified = HostNameType("Verified")
	HostNameTypeManaged  = HostNameType("Managed")
)

func (HostNameType) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameType)(nil)).Elem()
}

func (e HostNameType) ToHostNameTypeOutput() HostNameTypeOutput {
	return pulumi.ToOutput(e).(HostNameTypeOutput)
}

func (e HostNameType) ToHostNameTypeOutputWithContext(ctx context.Context) HostNameTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HostNameTypeOutput)
}

func (e HostNameType) ToHostNameTypePtrOutput() HostNameTypePtrOutput {
	return e.ToHostNameTypePtrOutputWithContext(context.Background())
}

func (e HostNameType) ToHostNameTypePtrOutputWithContext(ctx context.Context) HostNameTypePtrOutput {
	return HostNameType(e).ToHostNameTypeOutputWithContext(ctx).ToHostNameTypePtrOutputWithContext(ctx)
}

func (e HostNameType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostNameType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostNameType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HostNameType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HostNameTypeOutput struct{ *pulumi.OutputState }

func (HostNameTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameType)(nil)).Elem()
}

func (o HostNameTypeOutput) ToHostNameTypeOutput() HostNameTypeOutput {
	return o
}

func (o HostNameTypeOutput) ToHostNameTypeOutputWithContext(ctx context.Context) HostNameTypeOutput {
	return o
}

func (o HostNameTypeOutput) ToHostNameTypePtrOutput() HostNameTypePtrOutput {
	return o.ToHostNameTypePtrOutputWithContext(context.Background())
}

func (o HostNameTypeOutput) ToHostNameTypePtrOutputWithContext(ctx context.Context) HostNameTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostNameType) *HostNameType {
		return &v
	}).(HostNameTypePtrOutput)
}

func (o HostNameTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HostNameTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostNameType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HostNameTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostNameTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostNameType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HostNameTypePtrOutput struct{ *pulumi.OutputState }

func (HostNameTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostNameType)(nil)).Elem()
}

func (o HostNameTypePtrOutput) ToHostNameTypePtrOutput() HostNameTypePtrOutput {
	return o
}

func (o HostNameTypePtrOutput) ToHostNameTypePtrOutputWithContext(ctx context.Context) HostNameTypePtrOutput {
	return o
}

func (o HostNameTypePtrOutput) Elem() HostNameTypeOutput {
	return o.ApplyT(func(v *HostNameType) HostNameType {
		if v != nil {
			return *v
		}
		var ret HostNameType
		return ret
	}).(HostNameTypeOutput)
}

func (o HostNameTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostNameTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HostNameType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HostNameTypeInput interface {
	pulumi.Input

	ToHostNameTypeOutput() HostNameTypeOutput
	ToHostNameTypeOutputWithContext(context.Context) HostNameTypeOutput
}

var hostNameTypePtrType = reflect.TypeOf((**HostNameType)(nil)).Elem()

type HostNameTypePtrInput interface {
	pulumi.Input

	ToHostNameTypePtrOutput() HostNameTypePtrOutput
	ToHostNameTypePtrOutputWithContext(context.Context) HostNameTypePtrOutput
}

type hostNameTypePtr string

func HostNameTypePtr(v string) HostNameTypePtrInput {
	return (*hostNameTypePtr)(&v)
}

func (*hostNameTypePtr) ElementType() reflect.Type {
	return hostNameTypePtrType
}

func (in *hostNameTypePtr) ToHostNameTypePtrOutput() HostNameTypePtrOutput {
	return pulumi.ToOutput(in).(HostNameTypePtrOutput)
}

func (in *hostNameTypePtr) ToHostNameTypePtrOutputWithContext(ctx context.Context) HostNameTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HostNameTypePtrOutput)
}

type HostingEnvironmentStatus string

const (
	HostingEnvironmentStatusPreparing = HostingEnvironmentStatus("Preparing")
	HostingEnvironmentStatusReady     = HostingEnvironmentStatus("Ready")
	HostingEnvironmentStatusScaling   = HostingEnvironmentStatus("Scaling")
	HostingEnvironmentStatusDeleting  = HostingEnvironmentStatus("Deleting")
)

func (HostingEnvironmentStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentStatus)(nil)).Elem()
}

func (e HostingEnvironmentStatus) ToHostingEnvironmentStatusOutput() HostingEnvironmentStatusOutput {
	return pulumi.ToOutput(e).(HostingEnvironmentStatusOutput)
}

func (e HostingEnvironmentStatus) ToHostingEnvironmentStatusOutputWithContext(ctx context.Context) HostingEnvironmentStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HostingEnvironmentStatusOutput)
}

func (e HostingEnvironmentStatus) ToHostingEnvironmentStatusPtrOutput() HostingEnvironmentStatusPtrOutput {
	return e.ToHostingEnvironmentStatusPtrOutputWithContext(context.Background())
}

func (e HostingEnvironmentStatus) ToHostingEnvironmentStatusPtrOutputWithContext(ctx context.Context) HostingEnvironmentStatusPtrOutput {
	return HostingEnvironmentStatus(e).ToHostingEnvironmentStatusOutputWithContext(ctx).ToHostingEnvironmentStatusPtrOutputWithContext(ctx)
}

func (e HostingEnvironmentStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostingEnvironmentStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostingEnvironmentStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HostingEnvironmentStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HostingEnvironmentStatusOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentStatus)(nil)).Elem()
}

func (o HostingEnvironmentStatusOutput) ToHostingEnvironmentStatusOutput() HostingEnvironmentStatusOutput {
	return o
}

func (o HostingEnvironmentStatusOutput) ToHostingEnvironmentStatusOutputWithContext(ctx context.Context) HostingEnvironmentStatusOutput {
	return o
}

func (o HostingEnvironmentStatusOutput) ToHostingEnvironmentStatusPtrOutput() HostingEnvironmentStatusPtrOutput {
	return o.ToHostingEnvironmentStatusPtrOutputWithContext(context.Background())
}

func (o HostingEnvironmentStatusOutput) ToHostingEnvironmentStatusPtrOutputWithContext(ctx context.Context) HostingEnvironmentStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostingEnvironmentStatus) *HostingEnvironmentStatus {
		return &v
	}).(HostingEnvironmentStatusPtrOutput)
}

func (o HostingEnvironmentStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HostingEnvironmentStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostingEnvironmentStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HostingEnvironmentStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostingEnvironmentStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostingEnvironmentStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HostingEnvironmentStatusPtrOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingEnvironmentStatus)(nil)).Elem()
}

func (o HostingEnvironmentStatusPtrOutput) ToHostingEnvironmentStatusPtrOutput() HostingEnvironmentStatusPtrOutput {
	return o
}

func (o HostingEnvironmentStatusPtrOutput) ToHostingEnvironmentStatusPtrOutputWithContext(ctx context.Context) HostingEnvironmentStatusPtrOutput {
	return o
}

func (o HostingEnvironmentStatusPtrOutput) Elem() HostingEnvironmentStatusOutput {
	return o.ApplyT(func(v *HostingEnvironmentStatus) HostingEnvironmentStatus {
		if v != nil {
			return *v
		}
		var ret HostingEnvironmentStatus
		return ret
	}).(HostingEnvironmentStatusOutput)
}

func (o HostingEnvironmentStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostingEnvironmentStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HostingEnvironmentStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HostingEnvironmentStatusInput interface {
	pulumi.Input

	ToHostingEnvironmentStatusOutput() HostingEnvironmentStatusOutput
	ToHostingEnvironmentStatusOutputWithContext(context.Context) HostingEnvironmentStatusOutput
}

var hostingEnvironmentStatusPtrType = reflect.TypeOf((**HostingEnvironmentStatus)(nil)).Elem()

type HostingEnvironmentStatusPtrInput interface {
	pulumi.Input

	ToHostingEnvironmentStatusPtrOutput() HostingEnvironmentStatusPtrOutput
	ToHostingEnvironmentStatusPtrOutputWithContext(context.Context) HostingEnvironmentStatusPtrOutput
}

type hostingEnvironmentStatusPtr string

func HostingEnvironmentStatusPtr(v string) HostingEnvironmentStatusPtrInput {
	return (*hostingEnvironmentStatusPtr)(&v)
}

func (*hostingEnvironmentStatusPtr) ElementType() reflect.Type {
	return hostingEnvironmentStatusPtrType
}

func (in *hostingEnvironmentStatusPtr) ToHostingEnvironmentStatusPtrOutput() HostingEnvironmentStatusPtrOutput {
	return pulumi.ToOutput(in).(HostingEnvironmentStatusPtrOutput)
}

func (in *hostingEnvironmentStatusPtr) ToHostingEnvironmentStatusPtrOutputWithContext(ctx context.Context) HostingEnvironmentStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HostingEnvironmentStatusPtrOutput)
}

type InternalLoadBalancingMode string

const (
	InternalLoadBalancingModeNone       = InternalLoadBalancingMode("None")
	InternalLoadBalancingModeWeb        = InternalLoadBalancingMode("Web")
	InternalLoadBalancingModePublishing = InternalLoadBalancingMode("Publishing")
)

func (InternalLoadBalancingMode) ElementType() reflect.Type {
	return reflect.TypeOf((*InternalLoadBalancingMode)(nil)).Elem()
}

func (e InternalLoadBalancingMode) ToInternalLoadBalancingModeOutput() InternalLoadBalancingModeOutput {
	return pulumi.ToOutput(e).(InternalLoadBalancingModeOutput)
}

func (e InternalLoadBalancingMode) ToInternalLoadBalancingModeOutputWithContext(ctx context.Context) InternalLoadBalancingModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InternalLoadBalancingModeOutput)
}

func (e InternalLoadBalancingMode) ToInternalLoadBalancingModePtrOutput() InternalLoadBalancingModePtrOutput {
	return e.ToInternalLoadBalancingModePtrOutputWithContext(context.Background())
}

func (e InternalLoadBalancingMode) ToInternalLoadBalancingModePtrOutputWithContext(ctx context.Context) InternalLoadBalancingModePtrOutput {
	return InternalLoadBalancingMode(e).ToInternalLoadBalancingModeOutputWithContext(ctx).ToInternalLoadBalancingModePtrOutputWithContext(ctx)
}

func (e InternalLoadBalancingMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InternalLoadBalancingMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InternalLoadBalancingMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InternalLoadBalancingMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InternalLoadBalancingModeOutput struct{ *pulumi.OutputState }

func (InternalLoadBalancingModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InternalLoadBalancingMode)(nil)).Elem()
}

func (o InternalLoadBalancingModeOutput) ToInternalLoadBalancingModeOutput() InternalLoadBalancingModeOutput {
	return o
}

func (o InternalLoadBalancingModeOutput) ToInternalLoadBalancingModeOutputWithContext(ctx context.Context) InternalLoadBalancingModeOutput {
	return o
}

func (o InternalLoadBalancingModeOutput) ToInternalLoadBalancingModePtrOutput() InternalLoadBalancingModePtrOutput {
	return o.ToInternalLoadBalancingModePtrOutputWithContext(context.Background())
}

func (o InternalLoadBalancingModeOutput) ToInternalLoadBalancingModePtrOutputWithContext(ctx context.Context) InternalLoadBalancingModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InternalLoadBalancingMode) *InternalLoadBalancingMode {
		return &v
	}).(InternalLoadBalancingModePtrOutput)
}

func (o InternalLoadBalancingModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InternalLoadBalancingModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InternalLoadBalancingMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InternalLoadBalancingModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InternalLoadBalancingModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InternalLoadBalancingMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InternalLoadBalancingModePtrOutput struct{ *pulumi.OutputState }

func (InternalLoadBalancingModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InternalLoadBalancingMode)(nil)).Elem()
}

func (o InternalLoadBalancingModePtrOutput) ToInternalLoadBalancingModePtrOutput() InternalLoadBalancingModePtrOutput {
	return o
}

func (o InternalLoadBalancingModePtrOutput) ToInternalLoadBalancingModePtrOutputWithContext(ctx context.Context) InternalLoadBalancingModePtrOutput {
	return o
}

func (o InternalLoadBalancingModePtrOutput) Elem() InternalLoadBalancingModeOutput {
	return o.ApplyT(func(v *InternalLoadBalancingMode) InternalLoadBalancingMode {
		if v != nil {
			return *v
		}
		var ret InternalLoadBalancingMode
		return ret
	}).(InternalLoadBalancingModeOutput)
}

func (o InternalLoadBalancingModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InternalLoadBalancingModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InternalLoadBalancingMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type InternalLoadBalancingModeInput interface {
	pulumi.Input

	ToInternalLoadBalancingModeOutput() InternalLoadBalancingModeOutput
	ToInternalLoadBalancingModeOutputWithContext(context.Context) InternalLoadBalancingModeOutput
}

var internalLoadBalancingModePtrType = reflect.TypeOf((**InternalLoadBalancingMode)(nil)).Elem()

type InternalLoadBalancingModePtrInput interface {
	pulumi.Input

	ToInternalLoadBalancingModePtrOutput() InternalLoadBalancingModePtrOutput
	ToInternalLoadBalancingModePtrOutputWithContext(context.Context) InternalLoadBalancingModePtrOutput
}

type internalLoadBalancingModePtr string

func InternalLoadBalancingModePtr(v string) InternalLoadBalancingModePtrInput {
	return (*internalLoadBalancingModePtr)(&v)
}

func (*internalLoadBalancingModePtr) ElementType() reflect.Type {
	return internalLoadBalancingModePtrType
}

func (in *internalLoadBalancingModePtr) ToInternalLoadBalancingModePtrOutput() InternalLoadBalancingModePtrOutput {
	return pulumi.ToOutput(in).(InternalLoadBalancingModePtrOutput)
}

func (in *internalLoadBalancingModePtr) ToInternalLoadBalancingModePtrOutputWithContext(ctx context.Context) InternalLoadBalancingModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InternalLoadBalancingModePtrOutput)
}

type LogLevel string

const (
	LogLevelOff         = LogLevel("Off")
	LogLevelVerbose     = LogLevel("Verbose")
	LogLevelInformation = LogLevel("Information")
	LogLevelWarning     = LogLevel("Warning")
	LogLevelError       = LogLevel("Error")
)

func (LogLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*LogLevel)(nil)).Elem()
}

func (e LogLevel) ToLogLevelOutput() LogLevelOutput {
	return pulumi.ToOutput(e).(LogLevelOutput)
}

func (e LogLevel) ToLogLevelOutputWithContext(ctx context.Context) LogLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LogLevelOutput)
}

func (e LogLevel) ToLogLevelPtrOutput() LogLevelPtrOutput {
	return e.ToLogLevelPtrOutputWithContext(context.Background())
}

func (e LogLevel) ToLogLevelPtrOutputWithContext(ctx context.Context) LogLevelPtrOutput {
	return LogLevel(e).ToLogLevelOutputWithContext(ctx).ToLogLevelPtrOutputWithContext(ctx)
}

func (e LogLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LogLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LogLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LogLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LogLevelOutput struct{ *pulumi.OutputState }

func (LogLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogLevel)(nil)).Elem()
}

func (o LogLevelOutput) ToLogLevelOutput() LogLevelOutput {
	return o
}

func (o LogLevelOutput) ToLogLevelOutputWithContext(ctx context.Context) LogLevelOutput {
	return o
}

func (o LogLevelOutput) ToLogLevelPtrOutput() LogLevelPtrOutput {
	return o.ToLogLevelPtrOutputWithContext(context.Background())
}

func (o LogLevelOutput) ToLogLevelPtrOutputWithContext(ctx context.Context) LogLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogLevel) *LogLevel {
		return &v
	}).(LogLevelPtrOutput)
}

func (o LogLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LogLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LogLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LogLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LogLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LogLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LogLevelPtrOutput struct{ *pulumi.OutputState }

func (LogLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogLevel)(nil)).Elem()
}

func (o LogLevelPtrOutput) ToLogLevelPtrOutput() LogLevelPtrOutput {
	return o
}

func (o LogLevelPtrOutput) ToLogLevelPtrOutputWithContext(ctx context.Context) LogLevelPtrOutput {
	return o
}

func (o LogLevelPtrOutput) Elem() LogLevelOutput {
	return o.ApplyT(func(v *LogLevel) LogLevel {
		if v != nil {
			return *v
		}
		var ret LogLevel
		return ret
	}).(LogLevelOutput)
}

func (o LogLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LogLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LogLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LogLevelInput interface {
	pulumi.Input

	ToLogLevelOutput() LogLevelOutput
	ToLogLevelOutputWithContext(context.Context) LogLevelOutput
}

var logLevelPtrType = reflect.TypeOf((**LogLevel)(nil)).Elem()

type LogLevelPtrInput interface {
	pulumi.Input

	ToLogLevelPtrOutput() LogLevelPtrOutput
	ToLogLevelPtrOutputWithContext(context.Context) LogLevelPtrOutput
}

type logLevelPtr string

func LogLevelPtr(v string) LogLevelPtrInput {
	return (*logLevelPtr)(&v)
}

func (*logLevelPtr) ElementType() reflect.Type {
	return logLevelPtrType
}

func (in *logLevelPtr) ToLogLevelPtrOutput() LogLevelPtrOutput {
	return pulumi.ToOutput(in).(LogLevelPtrOutput)
}

func (in *logLevelPtr) ToLogLevelPtrOutputWithContext(ctx context.Context) LogLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LogLevelPtrOutput)
}

type ManagedPipelineMode string

const (
	ManagedPipelineModeIntegrated = ManagedPipelineMode("Integrated")
	ManagedPipelineModeClassic    = ManagedPipelineMode("Classic")
)

func (ManagedPipelineMode) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPipelineMode)(nil)).Elem()
}

func (e ManagedPipelineMode) ToManagedPipelineModeOutput() ManagedPipelineModeOutput {
	return pulumi.ToOutput(e).(ManagedPipelineModeOutput)
}

func (e ManagedPipelineMode) ToManagedPipelineModeOutputWithContext(ctx context.Context) ManagedPipelineModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedPipelineModeOutput)
}

func (e ManagedPipelineMode) ToManagedPipelineModePtrOutput() ManagedPipelineModePtrOutput {
	return e.ToManagedPipelineModePtrOutputWithContext(context.Background())
}

func (e ManagedPipelineMode) ToManagedPipelineModePtrOutputWithContext(ctx context.Context) ManagedPipelineModePtrOutput {
	return ManagedPipelineMode(e).ToManagedPipelineModeOutputWithContext(ctx).ToManagedPipelineModePtrOutputWithContext(ctx)
}

func (e ManagedPipelineMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedPipelineMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedPipelineMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedPipelineMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedPipelineModeOutput struct{ *pulumi.OutputState }

func (ManagedPipelineModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedPipelineMode)(nil)).Elem()
}

func (o ManagedPipelineModeOutput) ToManagedPipelineModeOutput() ManagedPipelineModeOutput {
	return o
}

func (o ManagedPipelineModeOutput) ToManagedPipelineModeOutputWithContext(ctx context.Context) ManagedPipelineModeOutput {
	return o
}

func (o ManagedPipelineModeOutput) ToManagedPipelineModePtrOutput() ManagedPipelineModePtrOutput {
	return o.ToManagedPipelineModePtrOutputWithContext(context.Background())
}

func (o ManagedPipelineModeOutput) ToManagedPipelineModePtrOutputWithContext(ctx context.Context) ManagedPipelineModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedPipelineMode) *ManagedPipelineMode {
		return &v
	}).(ManagedPipelineModePtrOutput)
}

func (o ManagedPipelineModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedPipelineModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedPipelineMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedPipelineModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedPipelineModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedPipelineMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedPipelineModePtrOutput struct{ *pulumi.OutputState }

func (ManagedPipelineModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPipelineMode)(nil)).Elem()
}

func (o ManagedPipelineModePtrOutput) ToManagedPipelineModePtrOutput() ManagedPipelineModePtrOutput {
	return o
}

func (o ManagedPipelineModePtrOutput) ToManagedPipelineModePtrOutputWithContext(ctx context.Context) ManagedPipelineModePtrOutput {
	return o
}

func (o ManagedPipelineModePtrOutput) Elem() ManagedPipelineModeOutput {
	return o.ApplyT(func(v *ManagedPipelineMode) ManagedPipelineMode {
		if v != nil {
			return *v
		}
		var ret ManagedPipelineMode
		return ret
	}).(ManagedPipelineModeOutput)
}

func (o ManagedPipelineModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedPipelineModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedPipelineMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ManagedPipelineModeInput interface {
	pulumi.Input

	ToManagedPipelineModeOutput() ManagedPipelineModeOutput
	ToManagedPipelineModeOutputWithContext(context.Context) ManagedPipelineModeOutput
}

var managedPipelineModePtrType = reflect.TypeOf((**ManagedPipelineMode)(nil)).Elem()

type ManagedPipelineModePtrInput interface {
	pulumi.Input

	ToManagedPipelineModePtrOutput() ManagedPipelineModePtrOutput
	ToManagedPipelineModePtrOutputWithContext(context.Context) ManagedPipelineModePtrOutput
}

type managedPipelineModePtr string

func ManagedPipelineModePtr(v string) ManagedPipelineModePtrInput {
	return (*managedPipelineModePtr)(&v)
}

func (*managedPipelineModePtr) ElementType() reflect.Type {
	return managedPipelineModePtrType
}

func (in *managedPipelineModePtr) ToManagedPipelineModePtrOutput() ManagedPipelineModePtrOutput {
	return pulumi.ToOutput(in).(ManagedPipelineModePtrOutput)
}

func (in *managedPipelineModePtr) ToManagedPipelineModePtrOutputWithContext(ctx context.Context) ManagedPipelineModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedPipelineModePtrOutput)
}

type ProvisioningState string

const (
	ProvisioningStateSucceeded  = ProvisioningState("Succeeded")
	ProvisioningStateFailed     = ProvisioningState("Failed")
	ProvisioningStateCanceled   = ProvisioningState("Canceled")
	ProvisioningStateInProgress = ProvisioningState("InProgress")
	ProvisioningStateDeleting   = ProvisioningState("Deleting")
)

func (ProvisioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningState)(nil)).Elem()
}

func (e ProvisioningState) ToProvisioningStateOutput() ProvisioningStateOutput {
	return pulumi.ToOutput(e).(ProvisioningStateOutput)
}

func (e ProvisioningState) ToProvisioningStateOutputWithContext(ctx context.Context) ProvisioningStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProvisioningStateOutput)
}

func (e ProvisioningState) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return e.ToProvisioningStatePtrOutputWithContext(context.Background())
}

func (e ProvisioningState) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return ProvisioningState(e).ToProvisioningStateOutputWithContext(ctx).ToProvisioningStatePtrOutputWithContext(ctx)
}

func (e ProvisioningState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProvisioningState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProvisioningStateOutput struct{ *pulumi.OutputState }

func (ProvisioningStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningState)(nil)).Elem()
}

func (o ProvisioningStateOutput) ToProvisioningStateOutput() ProvisioningStateOutput {
	return o
}

func (o ProvisioningStateOutput) ToProvisioningStateOutputWithContext(ctx context.Context) ProvisioningStateOutput {
	return o
}

func (o ProvisioningStateOutput) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return o.ToProvisioningStatePtrOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProvisioningState) *ProvisioningState {
		return &v
	}).(ProvisioningStatePtrOutput)
}

func (o ProvisioningStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProvisioningStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProvisioningStatePtrOutput struct{ *pulumi.OutputState }

func (ProvisioningStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisioningState)(nil)).Elem()
}

func (o ProvisioningStatePtrOutput) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return o
}

func (o ProvisioningStatePtrOutput) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return o
}

func (o ProvisioningStatePtrOutput) Elem() ProvisioningStateOutput {
	return o.ApplyT(func(v *ProvisioningState) ProvisioningState {
		if v != nil {
			return *v
		}
		var ret ProvisioningState
		return ret
	}).(ProvisioningStateOutput)
}

func (o ProvisioningStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProvisioningState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProvisioningStateInput interface {
	pulumi.Input

	ToProvisioningStateOutput() ProvisioningStateOutput
	ToProvisioningStateOutputWithContext(context.Context) ProvisioningStateOutput
}

var provisioningStatePtrType = reflect.TypeOf((**ProvisioningState)(nil)).Elem()

type ProvisioningStatePtrInput interface {
	pulumi.Input

	ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput
	ToProvisioningStatePtrOutputWithContext(context.Context) ProvisioningStatePtrOutput
}

type provisioningStatePtr string

func ProvisioningStatePtr(v string) ProvisioningStatePtrInput {
	return (*provisioningStatePtr)(&v)
}

func (*provisioningStatePtr) ElementType() reflect.Type {
	return provisioningStatePtrType
}

func (in *provisioningStatePtr) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return pulumi.ToOutput(in).(ProvisioningStatePtrOutput)
}

func (in *provisioningStatePtr) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProvisioningStatePtrOutput)
}

type SiteLoadBalancing string

const (
	SiteLoadBalancingWeightedRoundRobin   = SiteLoadBalancing("WeightedRoundRobin")
	SiteLoadBalancingLeastRequests        = SiteLoadBalancing("LeastRequests")
	SiteLoadBalancingLeastResponseTime    = SiteLoadBalancing("LeastResponseTime")
	SiteLoadBalancingWeightedTotalTraffic = SiteLoadBalancing("WeightedTotalTraffic")
	SiteLoadBalancingRequestHash          = SiteLoadBalancing("RequestHash")
)

func (SiteLoadBalancing) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteLoadBalancing)(nil)).Elem()
}

func (e SiteLoadBalancing) ToSiteLoadBalancingOutput() SiteLoadBalancingOutput {
	return pulumi.ToOutput(e).(SiteLoadBalancingOutput)
}

func (e SiteLoadBalancing) ToSiteLoadBalancingOutputWithContext(ctx context.Context) SiteLoadBalancingOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SiteLoadBalancingOutput)
}

func (e SiteLoadBalancing) ToSiteLoadBalancingPtrOutput() SiteLoadBalancingPtrOutput {
	return e.ToSiteLoadBalancingPtrOutputWithContext(context.Background())
}

func (e SiteLoadBalancing) ToSiteLoadBalancingPtrOutputWithContext(ctx context.Context) SiteLoadBalancingPtrOutput {
	return SiteLoadBalancing(e).ToSiteLoadBalancingOutputWithContext(ctx).ToSiteLoadBalancingPtrOutputWithContext(ctx)
}

func (e SiteLoadBalancing) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SiteLoadBalancing) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SiteLoadBalancing) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SiteLoadBalancing) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SiteLoadBalancingOutput struct{ *pulumi.OutputState }

func (SiteLoadBalancingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteLoadBalancing)(nil)).Elem()
}

func (o SiteLoadBalancingOutput) ToSiteLoadBalancingOutput() SiteLoadBalancingOutput {
	return o
}

func (o SiteLoadBalancingOutput) ToSiteLoadBalancingOutputWithContext(ctx context.Context) SiteLoadBalancingOutput {
	return o
}

func (o SiteLoadBalancingOutput) ToSiteLoadBalancingPtrOutput() SiteLoadBalancingPtrOutput {
	return o.ToSiteLoadBalancingPtrOutputWithContext(context.Background())
}

func (o SiteLoadBalancingOutput) ToSiteLoadBalancingPtrOutputWithContext(ctx context.Context) SiteLoadBalancingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteLoadBalancing) *SiteLoadBalancing {
		return &v
	}).(SiteLoadBalancingPtrOutput)
}

func (o SiteLoadBalancingOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SiteLoadBalancingOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SiteLoadBalancing) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SiteLoadBalancingOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SiteLoadBalancingOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SiteLoadBalancing) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SiteLoadBalancingPtrOutput struct{ *pulumi.OutputState }

func (SiteLoadBalancingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteLoadBalancing)(nil)).Elem()
}

func (o SiteLoadBalancingPtrOutput) ToSiteLoadBalancingPtrOutput() SiteLoadBalancingPtrOutput {
	return o
}

func (o SiteLoadBalancingPtrOutput) ToSiteLoadBalancingPtrOutputWithContext(ctx context.Context) SiteLoadBalancingPtrOutput {
	return o
}

func (o SiteLoadBalancingPtrOutput) Elem() SiteLoadBalancingOutput {
	return o.ApplyT(func(v *SiteLoadBalancing) SiteLoadBalancing {
		if v != nil {
			return *v
		}
		var ret SiteLoadBalancing
		return ret
	}).(SiteLoadBalancingOutput)
}

func (o SiteLoadBalancingPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SiteLoadBalancingPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SiteLoadBalancing) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SiteLoadBalancingInput interface {
	pulumi.Input

	ToSiteLoadBalancingOutput() SiteLoadBalancingOutput
	ToSiteLoadBalancingOutputWithContext(context.Context) SiteLoadBalancingOutput
}

var siteLoadBalancingPtrType = reflect.TypeOf((**SiteLoadBalancing)(nil)).Elem()

type SiteLoadBalancingPtrInput interface {
	pulumi.Input

	ToSiteLoadBalancingPtrOutput() SiteLoadBalancingPtrOutput
	ToSiteLoadBalancingPtrOutputWithContext(context.Context) SiteLoadBalancingPtrOutput
}

type siteLoadBalancingPtr string

func SiteLoadBalancingPtr(v string) SiteLoadBalancingPtrInput {
	return (*siteLoadBalancingPtr)(&v)
}

func (*siteLoadBalancingPtr) ElementType() reflect.Type {
	return siteLoadBalancingPtrType
}

func (in *siteLoadBalancingPtr) ToSiteLoadBalancingPtrOutput() SiteLoadBalancingPtrOutput {
	return pulumi.ToOutput(in).(SiteLoadBalancingPtrOutput)
}

func (in *siteLoadBalancingPtr) ToSiteLoadBalancingPtrOutputWithContext(ctx context.Context) SiteLoadBalancingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SiteLoadBalancingPtrOutput)
}

type SslState string

const (
	SslStateDisabled       = SslState("Disabled")
	SslStateSniEnabled     = SslState("SniEnabled")
	SslStateIpBasedEnabled = SslState("IpBasedEnabled")
)

func (SslState) ElementType() reflect.Type {
	return reflect.TypeOf((*SslState)(nil)).Elem()
}

func (e SslState) ToSslStateOutput() SslStateOutput {
	return pulumi.ToOutput(e).(SslStateOutput)
}

func (e SslState) ToSslStateOutputWithContext(ctx context.Context) SslStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SslStateOutput)
}

func (e SslState) ToSslStatePtrOutput() SslStatePtrOutput {
	return e.ToSslStatePtrOutputWithContext(context.Background())
}

func (e SslState) ToSslStatePtrOutputWithContext(ctx context.Context) SslStatePtrOutput {
	return SslState(e).ToSslStateOutputWithContext(ctx).ToSslStatePtrOutputWithContext(ctx)
}

func (e SslState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SslState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SslStateOutput struct{ *pulumi.OutputState }

func (SslStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslState)(nil)).Elem()
}

func (o SslStateOutput) ToSslStateOutput() SslStateOutput {
	return o
}

func (o SslStateOutput) ToSslStateOutputWithContext(ctx context.Context) SslStateOutput {
	return o
}

func (o SslStateOutput) ToSslStatePtrOutput() SslStatePtrOutput {
	return o.ToSslStatePtrOutputWithContext(context.Background())
}

func (o SslStateOutput) ToSslStatePtrOutputWithContext(ctx context.Context) SslStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SslState) *SslState {
		return &v
	}).(SslStatePtrOutput)
}

func (o SslStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SslStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SslStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SslStatePtrOutput struct{ *pulumi.OutputState }

func (SslStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslState)(nil)).Elem()
}

func (o SslStatePtrOutput) ToSslStatePtrOutput() SslStatePtrOutput {
	return o
}

func (o SslStatePtrOutput) ToSslStatePtrOutputWithContext(ctx context.Context) SslStatePtrOutput {
	return o
}

func (o SslStatePtrOutput) Elem() SslStateOutput {
	return o.ApplyT(func(v *SslState) SslState {
		if v != nil {
			return *v
		}
		var ret SslState
		return ret
	}).(SslStateOutput)
}

func (o SslStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SslState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SslStateInput interface {
	pulumi.Input

	ToSslStateOutput() SslStateOutput
	ToSslStateOutputWithContext(context.Context) SslStateOutput
}

var sslStatePtrType = reflect.TypeOf((**SslState)(nil)).Elem()

type SslStatePtrInput interface {
	pulumi.Input

	ToSslStatePtrOutput() SslStatePtrOutput
	ToSslStatePtrOutputWithContext(context.Context) SslStatePtrOutput
}

type sslStatePtr string

func SslStatePtr(v string) SslStatePtrInput {
	return (*sslStatePtr)(&v)
}

func (*sslStatePtr) ElementType() reflect.Type {
	return sslStatePtrType
}

func (in *sslStatePtr) ToSslStatePtrOutput() SslStatePtrOutput {
	return pulumi.ToOutput(in).(SslStatePtrOutput)
}

func (in *sslStatePtr) ToSslStatePtrOutputWithContext(ctx context.Context) SslStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SslStatePtrOutput)
}

type UnauthenticatedClientAction string

const (
	UnauthenticatedClientActionRedirectToLoginPage = UnauthenticatedClientAction("RedirectToLoginPage")
	UnauthenticatedClientActionAllowAnonymous      = UnauthenticatedClientAction("AllowAnonymous")
)

func (UnauthenticatedClientAction) ElementType() reflect.Type {
	return reflect.TypeOf((*UnauthenticatedClientAction)(nil)).Elem()
}

func (e UnauthenticatedClientAction) ToUnauthenticatedClientActionOutput() UnauthenticatedClientActionOutput {
	return pulumi.ToOutput(e).(UnauthenticatedClientActionOutput)
}

func (e UnauthenticatedClientAction) ToUnauthenticatedClientActionOutputWithContext(ctx context.Context) UnauthenticatedClientActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UnauthenticatedClientActionOutput)
}

func (e UnauthenticatedClientAction) ToUnauthenticatedClientActionPtrOutput() UnauthenticatedClientActionPtrOutput {
	return e.ToUnauthenticatedClientActionPtrOutputWithContext(context.Background())
}

func (e UnauthenticatedClientAction) ToUnauthenticatedClientActionPtrOutputWithContext(ctx context.Context) UnauthenticatedClientActionPtrOutput {
	return UnauthenticatedClientAction(e).ToUnauthenticatedClientActionOutputWithContext(ctx).ToUnauthenticatedClientActionPtrOutputWithContext(ctx)
}

func (e UnauthenticatedClientAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UnauthenticatedClientAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UnauthenticatedClientAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UnauthenticatedClientAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UnauthenticatedClientActionOutput struct{ *pulumi.OutputState }

func (UnauthenticatedClientActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UnauthenticatedClientAction)(nil)).Elem()
}

func (o UnauthenticatedClientActionOutput) ToUnauthenticatedClientActionOutput() UnauthenticatedClientActionOutput {
	return o
}

func (o UnauthenticatedClientActionOutput) ToUnauthenticatedClientActionOutputWithContext(ctx context.Context) UnauthenticatedClientActionOutput {
	return o
}

func (o UnauthenticatedClientActionOutput) ToUnauthenticatedClientActionPtrOutput() UnauthenticatedClientActionPtrOutput {
	return o.ToUnauthenticatedClientActionPtrOutputWithContext(context.Background())
}

func (o UnauthenticatedClientActionOutput) ToUnauthenticatedClientActionPtrOutputWithContext(ctx context.Context) UnauthenticatedClientActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UnauthenticatedClientAction) *UnauthenticatedClientAction {
		return &v
	}).(UnauthenticatedClientActionPtrOutput)
}

func (o UnauthenticatedClientActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UnauthenticatedClientActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UnauthenticatedClientAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UnauthenticatedClientActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UnauthenticatedClientActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UnauthenticatedClientAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UnauthenticatedClientActionPtrOutput struct{ *pulumi.OutputState }

func (UnauthenticatedClientActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UnauthenticatedClientAction)(nil)).Elem()
}

func (o UnauthenticatedClientActionPtrOutput) ToUnauthenticatedClientActionPtrOutput() UnauthenticatedClientActionPtrOutput {
	return o
}

func (o UnauthenticatedClientActionPtrOutput) ToUnauthenticatedClientActionPtrOutputWithContext(ctx context.Context) UnauthenticatedClientActionPtrOutput {
	return o
}

func (o UnauthenticatedClientActionPtrOutput) Elem() UnauthenticatedClientActionOutput {
	return o.ApplyT(func(v *UnauthenticatedClientAction) UnauthenticatedClientAction {
		if v != nil {
			return *v
		}
		var ret UnauthenticatedClientAction
		return ret
	}).(UnauthenticatedClientActionOutput)
}

func (o UnauthenticatedClientActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UnauthenticatedClientActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UnauthenticatedClientAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type UnauthenticatedClientActionInput interface {
	pulumi.Input

	ToUnauthenticatedClientActionOutput() UnauthenticatedClientActionOutput
	ToUnauthenticatedClientActionOutputWithContext(context.Context) UnauthenticatedClientActionOutput
}

var unauthenticatedClientActionPtrType = reflect.TypeOf((**UnauthenticatedClientAction)(nil)).Elem()

type UnauthenticatedClientActionPtrInput interface {
	pulumi.Input

	ToUnauthenticatedClientActionPtrOutput() UnauthenticatedClientActionPtrOutput
	ToUnauthenticatedClientActionPtrOutputWithContext(context.Context) UnauthenticatedClientActionPtrOutput
}

type unauthenticatedClientActionPtr string

func UnauthenticatedClientActionPtr(v string) UnauthenticatedClientActionPtrInput {
	return (*unauthenticatedClientActionPtr)(&v)
}

func (*unauthenticatedClientActionPtr) ElementType() reflect.Type {
	return unauthenticatedClientActionPtrType
}

func (in *unauthenticatedClientActionPtr) ToUnauthenticatedClientActionPtrOutput() UnauthenticatedClientActionPtrOutput {
	return pulumi.ToOutput(in).(UnauthenticatedClientActionPtrOutput)
}

func (in *unauthenticatedClientActionPtr) ToUnauthenticatedClientActionPtrOutputWithContext(ctx context.Context) UnauthenticatedClientActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UnauthenticatedClientActionPtrOutput)
}

type WorkerSizeOptions string

const (
	WorkerSizeOptionsDefault = WorkerSizeOptions("Default")
	WorkerSizeOptionsSmall   = WorkerSizeOptions("Small")
	WorkerSizeOptionsMedium  = WorkerSizeOptions("Medium")
	WorkerSizeOptionsLarge   = WorkerSizeOptions("Large")
)

func (WorkerSizeOptions) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerSizeOptions)(nil)).Elem()
}

func (e WorkerSizeOptions) ToWorkerSizeOptionsOutput() WorkerSizeOptionsOutput {
	return pulumi.ToOutput(e).(WorkerSizeOptionsOutput)
}

func (e WorkerSizeOptions) ToWorkerSizeOptionsOutputWithContext(ctx context.Context) WorkerSizeOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WorkerSizeOptionsOutput)
}

func (e WorkerSizeOptions) ToWorkerSizeOptionsPtrOutput() WorkerSizeOptionsPtrOutput {
	return e.ToWorkerSizeOptionsPtrOutputWithContext(context.Background())
}

func (e WorkerSizeOptions) ToWorkerSizeOptionsPtrOutputWithContext(ctx context.Context) WorkerSizeOptionsPtrOutput {
	return WorkerSizeOptions(e).ToWorkerSizeOptionsOutputWithContext(ctx).ToWorkerSizeOptionsPtrOutputWithContext(ctx)
}

func (e WorkerSizeOptions) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkerSizeOptions) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkerSizeOptions) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WorkerSizeOptions) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WorkerSizeOptionsOutput struct{ *pulumi.OutputState }

func (WorkerSizeOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerSizeOptions)(nil)).Elem()
}

func (o WorkerSizeOptionsOutput) ToWorkerSizeOptionsOutput() WorkerSizeOptionsOutput {
	return o
}

func (o WorkerSizeOptionsOutput) ToWorkerSizeOptionsOutputWithContext(ctx context.Context) WorkerSizeOptionsOutput {
	return o
}

func (o WorkerSizeOptionsOutput) ToWorkerSizeOptionsPtrOutput() WorkerSizeOptionsPtrOutput {
	return o.ToWorkerSizeOptionsPtrOutputWithContext(context.Background())
}

func (o WorkerSizeOptionsOutput) ToWorkerSizeOptionsPtrOutputWithContext(ctx context.Context) WorkerSizeOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkerSizeOptions) *WorkerSizeOptions {
		return &v
	}).(WorkerSizeOptionsPtrOutput)
}

func (o WorkerSizeOptionsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WorkerSizeOptionsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkerSizeOptions) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WorkerSizeOptionsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkerSizeOptionsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkerSizeOptions) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkerSizeOptionsPtrOutput struct{ *pulumi.OutputState }

func (WorkerSizeOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkerSizeOptions)(nil)).Elem()
}

func (o WorkerSizeOptionsPtrOutput) ToWorkerSizeOptionsPtrOutput() WorkerSizeOptionsPtrOutput {
	return o
}

func (o WorkerSizeOptionsPtrOutput) ToWorkerSizeOptionsPtrOutputWithContext(ctx context.Context) WorkerSizeOptionsPtrOutput {
	return o
}

func (o WorkerSizeOptionsPtrOutput) Elem() WorkerSizeOptionsOutput {
	return o.ApplyT(func(v *WorkerSizeOptions) WorkerSizeOptions {
		if v != nil {
			return *v
		}
		var ret WorkerSizeOptions
		return ret
	}).(WorkerSizeOptionsOutput)
}

func (o WorkerSizeOptionsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkerSizeOptionsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WorkerSizeOptions) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WorkerSizeOptionsInput interface {
	pulumi.Input

	ToWorkerSizeOptionsOutput() WorkerSizeOptionsOutput
	ToWorkerSizeOptionsOutputWithContext(context.Context) WorkerSizeOptionsOutput
}

var workerSizeOptionsPtrType = reflect.TypeOf((**WorkerSizeOptions)(nil)).Elem()

type WorkerSizeOptionsPtrInput interface {
	pulumi.Input

	ToWorkerSizeOptionsPtrOutput() WorkerSizeOptionsPtrOutput
	ToWorkerSizeOptionsPtrOutputWithContext(context.Context) WorkerSizeOptionsPtrOutput
}

type workerSizeOptionsPtr string

func WorkerSizeOptionsPtr(v string) WorkerSizeOptionsPtrInput {
	return (*workerSizeOptionsPtr)(&v)
}

func (*workerSizeOptionsPtr) ElementType() reflect.Type {
	return workerSizeOptionsPtrType
}

func (in *workerSizeOptionsPtr) ToWorkerSizeOptionsPtrOutput() WorkerSizeOptionsPtrOutput {
	return pulumi.ToOutput(in).(WorkerSizeOptionsPtrOutput)
}

func (in *workerSizeOptionsPtr) ToWorkerSizeOptionsPtrOutputWithContext(ctx context.Context) WorkerSizeOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WorkerSizeOptionsPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessControlEntryActionOutput{})
	pulumi.RegisterOutputType(AccessControlEntryActionPtrOutput{})
	pulumi.RegisterOutputType(AutoHealActionTypeOutput{})
	pulumi.RegisterOutputType(AutoHealActionTypePtrOutput{})
	pulumi.RegisterOutputType(AzureResourceTypeOutput{})
	pulumi.RegisterOutputType(AzureResourceTypePtrOutput{})
	pulumi.RegisterOutputType(BuiltInAuthenticationProviderOutput{})
	pulumi.RegisterOutputType(BuiltInAuthenticationProviderPtrOutput{})
	pulumi.RegisterOutputType(ComputeModeOptionsOutput{})
	pulumi.RegisterOutputType(ComputeModeOptionsPtrOutput{})
	pulumi.RegisterOutputType(CustomHostNameDnsRecordTypeOutput{})
	pulumi.RegisterOutputType(CustomHostNameDnsRecordTypePtrOutput{})
	pulumi.RegisterOutputType(DatabaseServerTypeOutput{})
	pulumi.RegisterOutputType(DatabaseServerTypePtrOutput{})
	pulumi.RegisterOutputType(FrequencyUnitOutput{})
	pulumi.RegisterOutputType(FrequencyUnitPtrOutput{})
	pulumi.RegisterOutputType(HostNameTypeOutput{})
	pulumi.RegisterOutputType(HostNameTypePtrOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentStatusOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentStatusPtrOutput{})
	pulumi.RegisterOutputType(InternalLoadBalancingModeOutput{})
	pulumi.RegisterOutputType(InternalLoadBalancingModePtrOutput{})
	pulumi.RegisterOutputType(LogLevelOutput{})
	pulumi.RegisterOutputType(LogLevelPtrOutput{})
	pulumi.RegisterOutputType(ManagedPipelineModeOutput{})
	pulumi.RegisterOutputType(ManagedPipelineModePtrOutput{})
	pulumi.RegisterOutputType(ProvisioningStateOutput{})
	pulumi.RegisterOutputType(ProvisioningStatePtrOutput{})
	pulumi.RegisterOutputType(SiteLoadBalancingOutput{})
	pulumi.RegisterOutputType(SiteLoadBalancingPtrOutput{})
	pulumi.RegisterOutputType(SslStateOutput{})
	pulumi.RegisterOutputType(SslStatePtrOutput{})
	pulumi.RegisterOutputType(UnauthenticatedClientActionOutput{})
	pulumi.RegisterOutputType(UnauthenticatedClientActionPtrOutput{})
	pulumi.RegisterOutputType(WorkerSizeOptionsOutput{})
	pulumi.RegisterOutputType(WorkerSizeOptionsPtrOutput{})
}
