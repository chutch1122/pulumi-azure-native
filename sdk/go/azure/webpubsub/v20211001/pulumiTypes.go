


package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EventHandler struct {
	Auth             *UpstreamAuthSettings `pulumi:"auth"`
	SystemEvents     []string              `pulumi:"systemEvents"`
	UrlTemplate      string                `pulumi:"urlTemplate"`
	UserEventPattern *string               `pulumi:"userEventPattern"`
}





type EventHandlerInput interface {
	pulumi.Input

	ToEventHandlerOutput() EventHandlerOutput
	ToEventHandlerOutputWithContext(context.Context) EventHandlerOutput
}

type EventHandlerArgs struct {
	Auth             UpstreamAuthSettingsPtrInput `pulumi:"auth"`
	SystemEvents     pulumi.StringArrayInput      `pulumi:"systemEvents"`
	UrlTemplate      pulumi.StringInput           `pulumi:"urlTemplate"`
	UserEventPattern pulumi.StringPtrInput        `pulumi:"userEventPattern"`
}

func (EventHandlerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHandler)(nil)).Elem()
}

func (i EventHandlerArgs) ToEventHandlerOutput() EventHandlerOutput {
	return i.ToEventHandlerOutputWithContext(context.Background())
}

func (i EventHandlerArgs) ToEventHandlerOutputWithContext(ctx context.Context) EventHandlerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHandlerOutput)
}





type EventHandlerArrayInput interface {
	pulumi.Input

	ToEventHandlerArrayOutput() EventHandlerArrayOutput
	ToEventHandlerArrayOutputWithContext(context.Context) EventHandlerArrayOutput
}

type EventHandlerArray []EventHandlerInput

func (EventHandlerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventHandler)(nil)).Elem()
}

func (i EventHandlerArray) ToEventHandlerArrayOutput() EventHandlerArrayOutput {
	return i.ToEventHandlerArrayOutputWithContext(context.Background())
}

func (i EventHandlerArray) ToEventHandlerArrayOutputWithContext(ctx context.Context) EventHandlerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHandlerArrayOutput)
}

type EventHandlerOutput struct{ *pulumi.OutputState }

func (EventHandlerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHandler)(nil)).Elem()
}

func (o EventHandlerOutput) ToEventHandlerOutput() EventHandlerOutput {
	return o
}

func (o EventHandlerOutput) ToEventHandlerOutputWithContext(ctx context.Context) EventHandlerOutput {
	return o
}

func (o EventHandlerOutput) Auth() UpstreamAuthSettingsPtrOutput {
	return o.ApplyT(func(v EventHandler) *UpstreamAuthSettings { return v.Auth }).(UpstreamAuthSettingsPtrOutput)
}

func (o EventHandlerOutput) SystemEvents() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EventHandler) []string { return v.SystemEvents }).(pulumi.StringArrayOutput)
}

func (o EventHandlerOutput) UrlTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v EventHandler) string { return v.UrlTemplate }).(pulumi.StringOutput)
}

func (o EventHandlerOutput) UserEventPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHandler) *string { return v.UserEventPattern }).(pulumi.StringPtrOutput)
}

type EventHandlerArrayOutput struct{ *pulumi.OutputState }

func (EventHandlerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventHandler)(nil)).Elem()
}

func (o EventHandlerArrayOutput) ToEventHandlerArrayOutput() EventHandlerArrayOutput {
	return o
}

func (o EventHandlerArrayOutput) ToEventHandlerArrayOutputWithContext(ctx context.Context) EventHandlerArrayOutput {
	return o
}

func (o EventHandlerArrayOutput) Index(i pulumi.IntInput) EventHandlerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventHandler {
		return vs[0].([]EventHandler)[vs[1].(int)]
	}).(EventHandlerOutput)
}

type EventHandlerResponse struct {
	Auth             *UpstreamAuthSettingsResponse `pulumi:"auth"`
	SystemEvents     []string                      `pulumi:"systemEvents"`
	UrlTemplate      string                        `pulumi:"urlTemplate"`
	UserEventPattern *string                       `pulumi:"userEventPattern"`
}

type EventHandlerResponseOutput struct{ *pulumi.OutputState }

func (EventHandlerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHandlerResponse)(nil)).Elem()
}

func (o EventHandlerResponseOutput) ToEventHandlerResponseOutput() EventHandlerResponseOutput {
	return o
}

func (o EventHandlerResponseOutput) ToEventHandlerResponseOutputWithContext(ctx context.Context) EventHandlerResponseOutput {
	return o
}

func (o EventHandlerResponseOutput) Auth() UpstreamAuthSettingsResponsePtrOutput {
	return o.ApplyT(func(v EventHandlerResponse) *UpstreamAuthSettingsResponse { return v.Auth }).(UpstreamAuthSettingsResponsePtrOutput)
}

func (o EventHandlerResponseOutput) SystemEvents() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EventHandlerResponse) []string { return v.SystemEvents }).(pulumi.StringArrayOutput)
}

func (o EventHandlerResponseOutput) UrlTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v EventHandlerResponse) string { return v.UrlTemplate }).(pulumi.StringOutput)
}

func (o EventHandlerResponseOutput) UserEventPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHandlerResponse) *string { return v.UserEventPattern }).(pulumi.StringPtrOutput)
}

type EventHandlerResponseArrayOutput struct{ *pulumi.OutputState }

func (EventHandlerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventHandlerResponse)(nil)).Elem()
}

func (o EventHandlerResponseArrayOutput) ToEventHandlerResponseArrayOutput() EventHandlerResponseArrayOutput {
	return o
}

func (o EventHandlerResponseArrayOutput) ToEventHandlerResponseArrayOutputWithContext(ctx context.Context) EventHandlerResponseArrayOutput {
	return o
}

func (o EventHandlerResponseArrayOutput) Index(i pulumi.IntInput) EventHandlerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventHandlerResponse {
		return vs[0].([]EventHandlerResponse)[vs[1].(int)]
	}).(EventHandlerResponseOutput)
}

type LiveTraceCategory struct {
	Enabled *string `pulumi:"enabled"`
	Name    *string `pulumi:"name"`
}





type LiveTraceCategoryInput interface {
	pulumi.Input

	ToLiveTraceCategoryOutput() LiveTraceCategoryOutput
	ToLiveTraceCategoryOutputWithContext(context.Context) LiveTraceCategoryOutput
}

type LiveTraceCategoryArgs struct {
	Enabled pulumi.StringPtrInput `pulumi:"enabled"`
	Name    pulumi.StringPtrInput `pulumi:"name"`
}

func (LiveTraceCategoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveTraceCategory)(nil)).Elem()
}

func (i LiveTraceCategoryArgs) ToLiveTraceCategoryOutput() LiveTraceCategoryOutput {
	return i.ToLiveTraceCategoryOutputWithContext(context.Background())
}

func (i LiveTraceCategoryArgs) ToLiveTraceCategoryOutputWithContext(ctx context.Context) LiveTraceCategoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveTraceCategoryOutput)
}





type LiveTraceCategoryArrayInput interface {
	pulumi.Input

	ToLiveTraceCategoryArrayOutput() LiveTraceCategoryArrayOutput
	ToLiveTraceCategoryArrayOutputWithContext(context.Context) LiveTraceCategoryArrayOutput
}

type LiveTraceCategoryArray []LiveTraceCategoryInput

func (LiveTraceCategoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveTraceCategory)(nil)).Elem()
}

func (i LiveTraceCategoryArray) ToLiveTraceCategoryArrayOutput() LiveTraceCategoryArrayOutput {
	return i.ToLiveTraceCategoryArrayOutputWithContext(context.Background())
}

func (i LiveTraceCategoryArray) ToLiveTraceCategoryArrayOutputWithContext(ctx context.Context) LiveTraceCategoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveTraceCategoryArrayOutput)
}

type LiveTraceCategoryOutput struct{ *pulumi.OutputState }

func (LiveTraceCategoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveTraceCategory)(nil)).Elem()
}

func (o LiveTraceCategoryOutput) ToLiveTraceCategoryOutput() LiveTraceCategoryOutput {
	return o
}

func (o LiveTraceCategoryOutput) ToLiveTraceCategoryOutputWithContext(ctx context.Context) LiveTraceCategoryOutput {
	return o
}

func (o LiveTraceCategoryOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveTraceCategory) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

func (o LiveTraceCategoryOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveTraceCategory) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type LiveTraceCategoryArrayOutput struct{ *pulumi.OutputState }

func (LiveTraceCategoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveTraceCategory)(nil)).Elem()
}

func (o LiveTraceCategoryArrayOutput) ToLiveTraceCategoryArrayOutput() LiveTraceCategoryArrayOutput {
	return o
}

func (o LiveTraceCategoryArrayOutput) ToLiveTraceCategoryArrayOutputWithContext(ctx context.Context) LiveTraceCategoryArrayOutput {
	return o
}

func (o LiveTraceCategoryArrayOutput) Index(i pulumi.IntInput) LiveTraceCategoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiveTraceCategory {
		return vs[0].([]LiveTraceCategory)[vs[1].(int)]
	}).(LiveTraceCategoryOutput)
}

type LiveTraceCategoryResponse struct {
	Enabled *string `pulumi:"enabled"`
	Name    *string `pulumi:"name"`
}

type LiveTraceCategoryResponseOutput struct{ *pulumi.OutputState }

func (LiveTraceCategoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveTraceCategoryResponse)(nil)).Elem()
}

func (o LiveTraceCategoryResponseOutput) ToLiveTraceCategoryResponseOutput() LiveTraceCategoryResponseOutput {
	return o
}

func (o LiveTraceCategoryResponseOutput) ToLiveTraceCategoryResponseOutputWithContext(ctx context.Context) LiveTraceCategoryResponseOutput {
	return o
}

func (o LiveTraceCategoryResponseOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveTraceCategoryResponse) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

func (o LiveTraceCategoryResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveTraceCategoryResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type LiveTraceCategoryResponseArrayOutput struct{ *pulumi.OutputState }

func (LiveTraceCategoryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LiveTraceCategoryResponse)(nil)).Elem()
}

func (o LiveTraceCategoryResponseArrayOutput) ToLiveTraceCategoryResponseArrayOutput() LiveTraceCategoryResponseArrayOutput {
	return o
}

func (o LiveTraceCategoryResponseArrayOutput) ToLiveTraceCategoryResponseArrayOutputWithContext(ctx context.Context) LiveTraceCategoryResponseArrayOutput {
	return o
}

func (o LiveTraceCategoryResponseArrayOutput) Index(i pulumi.IntInput) LiveTraceCategoryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LiveTraceCategoryResponse {
		return vs[0].([]LiveTraceCategoryResponse)[vs[1].(int)]
	}).(LiveTraceCategoryResponseOutput)
}

type LiveTraceConfiguration struct {
	Categories []LiveTraceCategory `pulumi:"categories"`
	Enabled    *string             `pulumi:"enabled"`
}


func (val *LiveTraceConfiguration) Defaults() *LiveTraceConfiguration {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Enabled) {
		enabled_ := "false"
		tmp.Enabled = &enabled_
	}
	return &tmp
}





type LiveTraceConfigurationInput interface {
	pulumi.Input

	ToLiveTraceConfigurationOutput() LiveTraceConfigurationOutput
	ToLiveTraceConfigurationOutputWithContext(context.Context) LiveTraceConfigurationOutput
}

type LiveTraceConfigurationArgs struct {
	Categories LiveTraceCategoryArrayInput `pulumi:"categories"`
	Enabled    pulumi.StringPtrInput       `pulumi:"enabled"`
}


func (val *LiveTraceConfigurationArgs) Defaults() *LiveTraceConfigurationArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Enabled) {
		tmp.Enabled = pulumi.StringPtr("false")
	}
	return &tmp
}
func (LiveTraceConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveTraceConfiguration)(nil)).Elem()
}

func (i LiveTraceConfigurationArgs) ToLiveTraceConfigurationOutput() LiveTraceConfigurationOutput {
	return i.ToLiveTraceConfigurationOutputWithContext(context.Background())
}

func (i LiveTraceConfigurationArgs) ToLiveTraceConfigurationOutputWithContext(ctx context.Context) LiveTraceConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveTraceConfigurationOutput)
}

func (i LiveTraceConfigurationArgs) ToLiveTraceConfigurationPtrOutput() LiveTraceConfigurationPtrOutput {
	return i.ToLiveTraceConfigurationPtrOutputWithContext(context.Background())
}

func (i LiveTraceConfigurationArgs) ToLiveTraceConfigurationPtrOutputWithContext(ctx context.Context) LiveTraceConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveTraceConfigurationOutput).ToLiveTraceConfigurationPtrOutputWithContext(ctx)
}









type LiveTraceConfigurationPtrInput interface {
	pulumi.Input

	ToLiveTraceConfigurationPtrOutput() LiveTraceConfigurationPtrOutput
	ToLiveTraceConfigurationPtrOutputWithContext(context.Context) LiveTraceConfigurationPtrOutput
}

type liveTraceConfigurationPtrType LiveTraceConfigurationArgs

func LiveTraceConfigurationPtr(v *LiveTraceConfigurationArgs) LiveTraceConfigurationPtrInput {
	return (*liveTraceConfigurationPtrType)(v)
}

func (*liveTraceConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveTraceConfiguration)(nil)).Elem()
}

func (i *liveTraceConfigurationPtrType) ToLiveTraceConfigurationPtrOutput() LiveTraceConfigurationPtrOutput {
	return i.ToLiveTraceConfigurationPtrOutputWithContext(context.Background())
}

func (i *liveTraceConfigurationPtrType) ToLiveTraceConfigurationPtrOutputWithContext(ctx context.Context) LiveTraceConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveTraceConfigurationPtrOutput)
}

type LiveTraceConfigurationOutput struct{ *pulumi.OutputState }

func (LiveTraceConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveTraceConfiguration)(nil)).Elem()
}

func (o LiveTraceConfigurationOutput) ToLiveTraceConfigurationOutput() LiveTraceConfigurationOutput {
	return o
}

func (o LiveTraceConfigurationOutput) ToLiveTraceConfigurationOutputWithContext(ctx context.Context) LiveTraceConfigurationOutput {
	return o
}

func (o LiveTraceConfigurationOutput) ToLiveTraceConfigurationPtrOutput() LiveTraceConfigurationPtrOutput {
	return o.ToLiveTraceConfigurationPtrOutputWithContext(context.Background())
}

func (o LiveTraceConfigurationOutput) ToLiveTraceConfigurationPtrOutputWithContext(ctx context.Context) LiveTraceConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LiveTraceConfiguration) *LiveTraceConfiguration {
		return &v
	}).(LiveTraceConfigurationPtrOutput)
}

func (o LiveTraceConfigurationOutput) Categories() LiveTraceCategoryArrayOutput {
	return o.ApplyT(func(v LiveTraceConfiguration) []LiveTraceCategory { return v.Categories }).(LiveTraceCategoryArrayOutput)
}

func (o LiveTraceConfigurationOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveTraceConfiguration) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

type LiveTraceConfigurationPtrOutput struct{ *pulumi.OutputState }

func (LiveTraceConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveTraceConfiguration)(nil)).Elem()
}

func (o LiveTraceConfigurationPtrOutput) ToLiveTraceConfigurationPtrOutput() LiveTraceConfigurationPtrOutput {
	return o
}

func (o LiveTraceConfigurationPtrOutput) ToLiveTraceConfigurationPtrOutputWithContext(ctx context.Context) LiveTraceConfigurationPtrOutput {
	return o
}

func (o LiveTraceConfigurationPtrOutput) Elem() LiveTraceConfigurationOutput {
	return o.ApplyT(func(v *LiveTraceConfiguration) LiveTraceConfiguration {
		if v != nil {
			return *v
		}
		var ret LiveTraceConfiguration
		return ret
	}).(LiveTraceConfigurationOutput)
}

func (o LiveTraceConfigurationPtrOutput) Categories() LiveTraceCategoryArrayOutput {
	return o.ApplyT(func(v *LiveTraceConfiguration) []LiveTraceCategory {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(LiveTraceCategoryArrayOutput)
}

func (o LiveTraceConfigurationPtrOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveTraceConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.StringPtrOutput)
}

type LiveTraceConfigurationResponse struct {
	Categories []LiveTraceCategoryResponse `pulumi:"categories"`
	Enabled    *string                     `pulumi:"enabled"`
}


func (val *LiveTraceConfigurationResponse) Defaults() *LiveTraceConfigurationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Enabled) {
		enabled_ := "false"
		tmp.Enabled = &enabled_
	}
	return &tmp
}

type LiveTraceConfigurationResponseOutput struct{ *pulumi.OutputState }

func (LiveTraceConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LiveTraceConfigurationResponse)(nil)).Elem()
}

func (o LiveTraceConfigurationResponseOutput) ToLiveTraceConfigurationResponseOutput() LiveTraceConfigurationResponseOutput {
	return o
}

func (o LiveTraceConfigurationResponseOutput) ToLiveTraceConfigurationResponseOutputWithContext(ctx context.Context) LiveTraceConfigurationResponseOutput {
	return o
}

func (o LiveTraceConfigurationResponseOutput) Categories() LiveTraceCategoryResponseArrayOutput {
	return o.ApplyT(func(v LiveTraceConfigurationResponse) []LiveTraceCategoryResponse { return v.Categories }).(LiveTraceCategoryResponseArrayOutput)
}

func (o LiveTraceConfigurationResponseOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LiveTraceConfigurationResponse) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

type LiveTraceConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (LiveTraceConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveTraceConfigurationResponse)(nil)).Elem()
}

func (o LiveTraceConfigurationResponsePtrOutput) ToLiveTraceConfigurationResponsePtrOutput() LiveTraceConfigurationResponsePtrOutput {
	return o
}

func (o LiveTraceConfigurationResponsePtrOutput) ToLiveTraceConfigurationResponsePtrOutputWithContext(ctx context.Context) LiveTraceConfigurationResponsePtrOutput {
	return o
}

func (o LiveTraceConfigurationResponsePtrOutput) Elem() LiveTraceConfigurationResponseOutput {
	return o.ApplyT(func(v *LiveTraceConfigurationResponse) LiveTraceConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret LiveTraceConfigurationResponse
		return ret
	}).(LiveTraceConfigurationResponseOutput)
}

func (o LiveTraceConfigurationResponsePtrOutput) Categories() LiveTraceCategoryResponseArrayOutput {
	return o.ApplyT(func(v *LiveTraceConfigurationResponse) []LiveTraceCategoryResponse {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(LiveTraceCategoryResponseArrayOutput)
}

func (o LiveTraceConfigurationResponsePtrOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveTraceConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.StringPtrOutput)
}

type ManagedIdentity struct {
	Type                   *string                `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ManagedIdentityInput interface {
	pulumi.Input

	ToManagedIdentityOutput() ManagedIdentityOutput
	ToManagedIdentityOutputWithContext(context.Context) ManagedIdentityOutput
}

type ManagedIdentityArgs struct {
	Type                   pulumi.StringPtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput       `pulumi:"userAssignedIdentities"`
}

func (ManagedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentity)(nil)).Elem()
}

func (i ManagedIdentityArgs) ToManagedIdentityOutput() ManagedIdentityOutput {
	return i.ToManagedIdentityOutputWithContext(context.Background())
}

func (i ManagedIdentityArgs) ToManagedIdentityOutputWithContext(ctx context.Context) ManagedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityOutput)
}

func (i ManagedIdentityArgs) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return i.ToManagedIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedIdentityArgs) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityOutput).ToManagedIdentityPtrOutputWithContext(ctx)
}









type ManagedIdentityPtrInput interface {
	pulumi.Input

	ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput
	ToManagedIdentityPtrOutputWithContext(context.Context) ManagedIdentityPtrOutput
}

type managedIdentityPtrType ManagedIdentityArgs

func ManagedIdentityPtr(v *ManagedIdentityArgs) ManagedIdentityPtrInput {
	return (*managedIdentityPtrType)(v)
}

func (*managedIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentity)(nil)).Elem()
}

func (i *managedIdentityPtrType) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return i.ToManagedIdentityPtrOutputWithContext(context.Background())
}

func (i *managedIdentityPtrType) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentityPtrOutput)
}

type ManagedIdentityOutput struct{ *pulumi.OutputState }

func (ManagedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentity)(nil)).Elem()
}

func (o ManagedIdentityOutput) ToManagedIdentityOutput() ManagedIdentityOutput {
	return o
}

func (o ManagedIdentityOutput) ToManagedIdentityOutputWithContext(ctx context.Context) ManagedIdentityOutput {
	return o
}

func (o ManagedIdentityOutput) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return o.ToManagedIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityOutput) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentity) *ManagedIdentity {
		return &v
	}).(ManagedIdentityPtrOutput)
}

func (o ManagedIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentity)(nil)).Elem()
}

func (o ManagedIdentityPtrOutput) ToManagedIdentityPtrOutput() ManagedIdentityPtrOutput {
	return o
}

func (o ManagedIdentityPtrOutput) ToManagedIdentityPtrOutputWithContext(ctx context.Context) ManagedIdentityPtrOutput {
	return o
}

func (o ManagedIdentityPtrOutput) Elem() ManagedIdentityOutput {
	return o.ApplyT(func(v *ManagedIdentity) ManagedIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedIdentity
		return ret
	}).(ManagedIdentityOutput)
}

func (o ManagedIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ManagedIdentityResponse struct {
	PrincipalId            string                                          `pulumi:"principalId"`
	TenantId               string                                          `pulumi:"tenantId"`
	Type                   *string                                         `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityPropertyResponse `pulumi:"userAssignedIdentities"`
}

type ManagedIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityResponse)(nil)).Elem()
}

func (o ManagedIdentityResponseOutput) ToManagedIdentityResponseOutput() ManagedIdentityResponseOutput {
	return o
}

func (o ManagedIdentityResponseOutput) ToManagedIdentityResponseOutputWithContext(ctx context.Context) ManagedIdentityResponseOutput {
	return o
}

func (o ManagedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityPropertyResponseMapOutput {
	return o.ApplyT(func(v ManagedIdentityResponse) map[string]UserAssignedIdentityPropertyResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityPropertyResponseMapOutput)
}

type ManagedIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityResponse)(nil)).Elem()
}

func (o ManagedIdentityResponsePtrOutput) ToManagedIdentityResponsePtrOutput() ManagedIdentityResponsePtrOutput {
	return o
}

func (o ManagedIdentityResponsePtrOutput) ToManagedIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedIdentityResponsePtrOutput {
	return o
}

func (o ManagedIdentityResponsePtrOutput) Elem() ManagedIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) ManagedIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityResponse
		return ret
	}).(ManagedIdentityResponseOutput)
}

func (o ManagedIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityPropertyResponseMapOutput {
	return o.ApplyT(func(v *ManagedIdentityResponse) map[string]UserAssignedIdentityPropertyResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityPropertyResponseMapOutput)
}

type ManagedIdentitySettings struct {
	Resource *string `pulumi:"resource"`
}





type ManagedIdentitySettingsInput interface {
	pulumi.Input

	ToManagedIdentitySettingsOutput() ManagedIdentitySettingsOutput
	ToManagedIdentitySettingsOutputWithContext(context.Context) ManagedIdentitySettingsOutput
}

type ManagedIdentitySettingsArgs struct {
	Resource pulumi.StringPtrInput `pulumi:"resource"`
}

func (ManagedIdentitySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentitySettings)(nil)).Elem()
}

func (i ManagedIdentitySettingsArgs) ToManagedIdentitySettingsOutput() ManagedIdentitySettingsOutput {
	return i.ToManagedIdentitySettingsOutputWithContext(context.Background())
}

func (i ManagedIdentitySettingsArgs) ToManagedIdentitySettingsOutputWithContext(ctx context.Context) ManagedIdentitySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentitySettingsOutput)
}

func (i ManagedIdentitySettingsArgs) ToManagedIdentitySettingsPtrOutput() ManagedIdentitySettingsPtrOutput {
	return i.ToManagedIdentitySettingsPtrOutputWithContext(context.Background())
}

func (i ManagedIdentitySettingsArgs) ToManagedIdentitySettingsPtrOutputWithContext(ctx context.Context) ManagedIdentitySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentitySettingsOutput).ToManagedIdentitySettingsPtrOutputWithContext(ctx)
}









type ManagedIdentitySettingsPtrInput interface {
	pulumi.Input

	ToManagedIdentitySettingsPtrOutput() ManagedIdentitySettingsPtrOutput
	ToManagedIdentitySettingsPtrOutputWithContext(context.Context) ManagedIdentitySettingsPtrOutput
}

type managedIdentitySettingsPtrType ManagedIdentitySettingsArgs

func ManagedIdentitySettingsPtr(v *ManagedIdentitySettingsArgs) ManagedIdentitySettingsPtrInput {
	return (*managedIdentitySettingsPtrType)(v)
}

func (*managedIdentitySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentitySettings)(nil)).Elem()
}

func (i *managedIdentitySettingsPtrType) ToManagedIdentitySettingsPtrOutput() ManagedIdentitySettingsPtrOutput {
	return i.ToManagedIdentitySettingsPtrOutputWithContext(context.Background())
}

func (i *managedIdentitySettingsPtrType) ToManagedIdentitySettingsPtrOutputWithContext(ctx context.Context) ManagedIdentitySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedIdentitySettingsPtrOutput)
}

type ManagedIdentitySettingsOutput struct{ *pulumi.OutputState }

func (ManagedIdentitySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentitySettings)(nil)).Elem()
}

func (o ManagedIdentitySettingsOutput) ToManagedIdentitySettingsOutput() ManagedIdentitySettingsOutput {
	return o
}

func (o ManagedIdentitySettingsOutput) ToManagedIdentitySettingsOutputWithContext(ctx context.Context) ManagedIdentitySettingsOutput {
	return o
}

func (o ManagedIdentitySettingsOutput) ToManagedIdentitySettingsPtrOutput() ManagedIdentitySettingsPtrOutput {
	return o.ToManagedIdentitySettingsPtrOutputWithContext(context.Background())
}

func (o ManagedIdentitySettingsOutput) ToManagedIdentitySettingsPtrOutputWithContext(ctx context.Context) ManagedIdentitySettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentitySettings) *ManagedIdentitySettings {
		return &v
	}).(ManagedIdentitySettingsPtrOutput)
}

func (o ManagedIdentitySettingsOutput) Resource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentitySettings) *string { return v.Resource }).(pulumi.StringPtrOutput)
}

type ManagedIdentitySettingsPtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentitySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentitySettings)(nil)).Elem()
}

func (o ManagedIdentitySettingsPtrOutput) ToManagedIdentitySettingsPtrOutput() ManagedIdentitySettingsPtrOutput {
	return o
}

func (o ManagedIdentitySettingsPtrOutput) ToManagedIdentitySettingsPtrOutputWithContext(ctx context.Context) ManagedIdentitySettingsPtrOutput {
	return o
}

func (o ManagedIdentitySettingsPtrOutput) Elem() ManagedIdentitySettingsOutput {
	return o.ApplyT(func(v *ManagedIdentitySettings) ManagedIdentitySettings {
		if v != nil {
			return *v
		}
		var ret ManagedIdentitySettings
		return ret
	}).(ManagedIdentitySettingsOutput)
}

func (o ManagedIdentitySettingsPtrOutput) Resource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentitySettings) *string {
		if v == nil {
			return nil
		}
		return v.Resource
	}).(pulumi.StringPtrOutput)
}

type ManagedIdentitySettingsResponse struct {
	Resource *string `pulumi:"resource"`
}

type ManagedIdentitySettingsResponseOutput struct{ *pulumi.OutputState }

func (ManagedIdentitySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentitySettingsResponse)(nil)).Elem()
}

func (o ManagedIdentitySettingsResponseOutput) ToManagedIdentitySettingsResponseOutput() ManagedIdentitySettingsResponseOutput {
	return o
}

func (o ManagedIdentitySettingsResponseOutput) ToManagedIdentitySettingsResponseOutputWithContext(ctx context.Context) ManagedIdentitySettingsResponseOutput {
	return o
}

func (o ManagedIdentitySettingsResponseOutput) Resource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedIdentitySettingsResponse) *string { return v.Resource }).(pulumi.StringPtrOutput)
}

type ManagedIdentitySettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentitySettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentitySettingsResponse)(nil)).Elem()
}

func (o ManagedIdentitySettingsResponsePtrOutput) ToManagedIdentitySettingsResponsePtrOutput() ManagedIdentitySettingsResponsePtrOutput {
	return o
}

func (o ManagedIdentitySettingsResponsePtrOutput) ToManagedIdentitySettingsResponsePtrOutputWithContext(ctx context.Context) ManagedIdentitySettingsResponsePtrOutput {
	return o
}

func (o ManagedIdentitySettingsResponsePtrOutput) Elem() ManagedIdentitySettingsResponseOutput {
	return o.ApplyT(func(v *ManagedIdentitySettingsResponse) ManagedIdentitySettingsResponse {
		if v != nil {
			return *v
		}
		var ret ManagedIdentitySettingsResponse
		return ret
	}).(ManagedIdentitySettingsResponseOutput)
}

func (o ManagedIdentitySettingsResponsePtrOutput) Resource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedIdentitySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Resource
	}).(pulumi.StringPtrOutput)
}

type NetworkACL struct {
	Allow []string `pulumi:"allow"`
	Deny  []string `pulumi:"deny"`
}





type NetworkACLInput interface {
	pulumi.Input

	ToNetworkACLOutput() NetworkACLOutput
	ToNetworkACLOutputWithContext(context.Context) NetworkACLOutput
}

type NetworkACLArgs struct {
	Allow pulumi.StringArrayInput `pulumi:"allow"`
	Deny  pulumi.StringArrayInput `pulumi:"deny"`
}

func (NetworkACLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkACL)(nil)).Elem()
}

func (i NetworkACLArgs) ToNetworkACLOutput() NetworkACLOutput {
	return i.ToNetworkACLOutputWithContext(context.Background())
}

func (i NetworkACLArgs) ToNetworkACLOutputWithContext(ctx context.Context) NetworkACLOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkACLOutput)
}

func (i NetworkACLArgs) ToNetworkACLPtrOutput() NetworkACLPtrOutput {
	return i.ToNetworkACLPtrOutputWithContext(context.Background())
}

func (i NetworkACLArgs) ToNetworkACLPtrOutputWithContext(ctx context.Context) NetworkACLPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkACLOutput).ToNetworkACLPtrOutputWithContext(ctx)
}









type NetworkACLPtrInput interface {
	pulumi.Input

	ToNetworkACLPtrOutput() NetworkACLPtrOutput
	ToNetworkACLPtrOutputWithContext(context.Context) NetworkACLPtrOutput
}

type networkACLPtrType NetworkACLArgs

func NetworkACLPtr(v *NetworkACLArgs) NetworkACLPtrInput {
	return (*networkACLPtrType)(v)
}

func (*networkACLPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkACL)(nil)).Elem()
}

func (i *networkACLPtrType) ToNetworkACLPtrOutput() NetworkACLPtrOutput {
	return i.ToNetworkACLPtrOutputWithContext(context.Background())
}

func (i *networkACLPtrType) ToNetworkACLPtrOutputWithContext(ctx context.Context) NetworkACLPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkACLPtrOutput)
}

type NetworkACLOutput struct{ *pulumi.OutputState }

func (NetworkACLOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkACL)(nil)).Elem()
}

func (o NetworkACLOutput) ToNetworkACLOutput() NetworkACLOutput {
	return o
}

func (o NetworkACLOutput) ToNetworkACLOutputWithContext(ctx context.Context) NetworkACLOutput {
	return o
}

func (o NetworkACLOutput) ToNetworkACLPtrOutput() NetworkACLPtrOutput {
	return o.ToNetworkACLPtrOutputWithContext(context.Background())
}

func (o NetworkACLOutput) ToNetworkACLPtrOutputWithContext(ctx context.Context) NetworkACLPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkACL) *NetworkACL {
		return &v
	}).(NetworkACLPtrOutput)
}

func (o NetworkACLOutput) Allow() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkACL) []string { return v.Allow }).(pulumi.StringArrayOutput)
}

func (o NetworkACLOutput) Deny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkACL) []string { return v.Deny }).(pulumi.StringArrayOutput)
}

type NetworkACLPtrOutput struct{ *pulumi.OutputState }

func (NetworkACLPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkACL)(nil)).Elem()
}

func (o NetworkACLPtrOutput) ToNetworkACLPtrOutput() NetworkACLPtrOutput {
	return o
}

func (o NetworkACLPtrOutput) ToNetworkACLPtrOutputWithContext(ctx context.Context) NetworkACLPtrOutput {
	return o
}

func (o NetworkACLPtrOutput) Elem() NetworkACLOutput {
	return o.ApplyT(func(v *NetworkACL) NetworkACL {
		if v != nil {
			return *v
		}
		var ret NetworkACL
		return ret
	}).(NetworkACLOutput)
}

func (o NetworkACLPtrOutput) Allow() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkACL) []string {
		if v == nil {
			return nil
		}
		return v.Allow
	}).(pulumi.StringArrayOutput)
}

func (o NetworkACLPtrOutput) Deny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkACL) []string {
		if v == nil {
			return nil
		}
		return v.Deny
	}).(pulumi.StringArrayOutput)
}

type NetworkACLResponse struct {
	Allow []string `pulumi:"allow"`
	Deny  []string `pulumi:"deny"`
}

type NetworkACLResponseOutput struct{ *pulumi.OutputState }

func (NetworkACLResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkACLResponse)(nil)).Elem()
}

func (o NetworkACLResponseOutput) ToNetworkACLResponseOutput() NetworkACLResponseOutput {
	return o
}

func (o NetworkACLResponseOutput) ToNetworkACLResponseOutputWithContext(ctx context.Context) NetworkACLResponseOutput {
	return o
}

func (o NetworkACLResponseOutput) Allow() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkACLResponse) []string { return v.Allow }).(pulumi.StringArrayOutput)
}

func (o NetworkACLResponseOutput) Deny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkACLResponse) []string { return v.Deny }).(pulumi.StringArrayOutput)
}

type NetworkACLResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkACLResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkACLResponse)(nil)).Elem()
}

func (o NetworkACLResponsePtrOutput) ToNetworkACLResponsePtrOutput() NetworkACLResponsePtrOutput {
	return o
}

func (o NetworkACLResponsePtrOutput) ToNetworkACLResponsePtrOutputWithContext(ctx context.Context) NetworkACLResponsePtrOutput {
	return o
}

func (o NetworkACLResponsePtrOutput) Elem() NetworkACLResponseOutput {
	return o.ApplyT(func(v *NetworkACLResponse) NetworkACLResponse {
		if v != nil {
			return *v
		}
		var ret NetworkACLResponse
		return ret
	}).(NetworkACLResponseOutput)
}

func (o NetworkACLResponsePtrOutput) Allow() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkACLResponse) []string {
		if v == nil {
			return nil
		}
		return v.Allow
	}).(pulumi.StringArrayOutput)
}

func (o NetworkACLResponsePtrOutput) Deny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkACLResponse) []string {
		if v == nil {
			return nil
		}
		return v.Deny
	}).(pulumi.StringArrayOutput)
}

type PrivateEndpoint struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointOutput() PrivateEndpointOutput
	ToPrivateEndpointOutputWithContext(context.Context) PrivateEndpointOutput
}

type PrivateEndpointArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return i.ToPrivateEndpointOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput)
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput).ToPrivateEndpointPtrOutputWithContext(ctx)
}









type PrivateEndpointPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput
	ToPrivateEndpointPtrOutputWithContext(context.Context) PrivateEndpointPtrOutput
}

type privateEndpointPtrType PrivateEndpointArgs

func PrivateEndpointPtr(v *PrivateEndpointArgs) PrivateEndpointPtrInput {
	return (*privateEndpointPtrType)(v)
}

func (*privateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPtrOutput)
}

type PrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpoint) *PrivateEndpoint {
		return &v
	}).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) Elem() PrivateEndpointOutput {
	return o.ApplyT(func(v *PrivateEndpoint) PrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret PrivateEndpoint
		return ret
	}).(PrivateEndpointOutput)
}

func (o PrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointACL struct {
	Allow []string `pulumi:"allow"`
	Deny  []string `pulumi:"deny"`
	Name  string   `pulumi:"name"`
}





type PrivateEndpointACLInput interface {
	pulumi.Input

	ToPrivateEndpointACLOutput() PrivateEndpointACLOutput
	ToPrivateEndpointACLOutputWithContext(context.Context) PrivateEndpointACLOutput
}

type PrivateEndpointACLArgs struct {
	Allow pulumi.StringArrayInput `pulumi:"allow"`
	Deny  pulumi.StringArrayInput `pulumi:"deny"`
	Name  pulumi.StringInput      `pulumi:"name"`
}

func (PrivateEndpointACLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointACL)(nil)).Elem()
}

func (i PrivateEndpointACLArgs) ToPrivateEndpointACLOutput() PrivateEndpointACLOutput {
	return i.ToPrivateEndpointACLOutputWithContext(context.Background())
}

func (i PrivateEndpointACLArgs) ToPrivateEndpointACLOutputWithContext(ctx context.Context) PrivateEndpointACLOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointACLOutput)
}





type PrivateEndpointACLArrayInput interface {
	pulumi.Input

	ToPrivateEndpointACLArrayOutput() PrivateEndpointACLArrayOutput
	ToPrivateEndpointACLArrayOutputWithContext(context.Context) PrivateEndpointACLArrayOutput
}

type PrivateEndpointACLArray []PrivateEndpointACLInput

func (PrivateEndpointACLArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointACL)(nil)).Elem()
}

func (i PrivateEndpointACLArray) ToPrivateEndpointACLArrayOutput() PrivateEndpointACLArrayOutput {
	return i.ToPrivateEndpointACLArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointACLArray) ToPrivateEndpointACLArrayOutputWithContext(ctx context.Context) PrivateEndpointACLArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointACLArrayOutput)
}

type PrivateEndpointACLOutput struct{ *pulumi.OutputState }

func (PrivateEndpointACLOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointACL)(nil)).Elem()
}

func (o PrivateEndpointACLOutput) ToPrivateEndpointACLOutput() PrivateEndpointACLOutput {
	return o
}

func (o PrivateEndpointACLOutput) ToPrivateEndpointACLOutputWithContext(ctx context.Context) PrivateEndpointACLOutput {
	return o
}

func (o PrivateEndpointACLOutput) Allow() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointACL) []string { return v.Allow }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointACLOutput) Deny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointACL) []string { return v.Deny }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointACLOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointACL) string { return v.Name }).(pulumi.StringOutput)
}

type PrivateEndpointACLArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointACLArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointACL)(nil)).Elem()
}

func (o PrivateEndpointACLArrayOutput) ToPrivateEndpointACLArrayOutput() PrivateEndpointACLArrayOutput {
	return o
}

func (o PrivateEndpointACLArrayOutput) ToPrivateEndpointACLArrayOutputWithContext(ctx context.Context) PrivateEndpointACLArrayOutput {
	return o
}

func (o PrivateEndpointACLArrayOutput) Index(i pulumi.IntInput) PrivateEndpointACLOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointACL {
		return vs[0].([]PrivateEndpointACL)[vs[1].(int)]
	}).(PrivateEndpointACLOutput)
}

type PrivateEndpointACLResponse struct {
	Allow []string `pulumi:"allow"`
	Deny  []string `pulumi:"deny"`
	Name  string   `pulumi:"name"`
}

type PrivateEndpointACLResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointACLResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointACLResponse)(nil)).Elem()
}

func (o PrivateEndpointACLResponseOutput) ToPrivateEndpointACLResponseOutput() PrivateEndpointACLResponseOutput {
	return o
}

func (o PrivateEndpointACLResponseOutput) ToPrivateEndpointACLResponseOutputWithContext(ctx context.Context) PrivateEndpointACLResponseOutput {
	return o
}

func (o PrivateEndpointACLResponseOutput) Allow() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointACLResponse) []string { return v.Allow }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointACLResponseOutput) Deny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointACLResponse) []string { return v.Deny }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointACLResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointACLResponse) string { return v.Name }).(pulumi.StringOutput)
}

type PrivateEndpointACLResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointACLResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointACLResponse)(nil)).Elem()
}

func (o PrivateEndpointACLResponseArrayOutput) ToPrivateEndpointACLResponseArrayOutput() PrivateEndpointACLResponseArrayOutput {
	return o
}

func (o PrivateEndpointACLResponseArrayOutput) ToPrivateEndpointACLResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointACLResponseArrayOutput {
	return o
}

func (o PrivateEndpointACLResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointACLResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointACLResponse {
		return vs[0].([]PrivateEndpointACLResponse)[vs[1].(int)]
	}).(PrivateEndpointACLResponseOutput)
}

type PrivateEndpointConnectionResponse struct {
	GroupIds                          []string                                   `pulumi:"groupIds"`
	Id                                string                                     `pulumi:"id"`
	Name                              string                                     `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                         `pulumi:"systemData"`
	Type                              string                                     `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id *string `pulumi:"id"`
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ResourceLogCategory struct {
	Enabled *string `pulumi:"enabled"`
	Name    *string `pulumi:"name"`
}





type ResourceLogCategoryInput interface {
	pulumi.Input

	ToResourceLogCategoryOutput() ResourceLogCategoryOutput
	ToResourceLogCategoryOutputWithContext(context.Context) ResourceLogCategoryOutput
}

type ResourceLogCategoryArgs struct {
	Enabled pulumi.StringPtrInput `pulumi:"enabled"`
	Name    pulumi.StringPtrInput `pulumi:"name"`
}

func (ResourceLogCategoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLogCategory)(nil)).Elem()
}

func (i ResourceLogCategoryArgs) ToResourceLogCategoryOutput() ResourceLogCategoryOutput {
	return i.ToResourceLogCategoryOutputWithContext(context.Background())
}

func (i ResourceLogCategoryArgs) ToResourceLogCategoryOutputWithContext(ctx context.Context) ResourceLogCategoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLogCategoryOutput)
}





type ResourceLogCategoryArrayInput interface {
	pulumi.Input

	ToResourceLogCategoryArrayOutput() ResourceLogCategoryArrayOutput
	ToResourceLogCategoryArrayOutputWithContext(context.Context) ResourceLogCategoryArrayOutput
}

type ResourceLogCategoryArray []ResourceLogCategoryInput

func (ResourceLogCategoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceLogCategory)(nil)).Elem()
}

func (i ResourceLogCategoryArray) ToResourceLogCategoryArrayOutput() ResourceLogCategoryArrayOutput {
	return i.ToResourceLogCategoryArrayOutputWithContext(context.Background())
}

func (i ResourceLogCategoryArray) ToResourceLogCategoryArrayOutputWithContext(ctx context.Context) ResourceLogCategoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLogCategoryArrayOutput)
}

type ResourceLogCategoryOutput struct{ *pulumi.OutputState }

func (ResourceLogCategoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLogCategory)(nil)).Elem()
}

func (o ResourceLogCategoryOutput) ToResourceLogCategoryOutput() ResourceLogCategoryOutput {
	return o
}

func (o ResourceLogCategoryOutput) ToResourceLogCategoryOutputWithContext(ctx context.Context) ResourceLogCategoryOutput {
	return o
}

func (o ResourceLogCategoryOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceLogCategory) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

func (o ResourceLogCategoryOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceLogCategory) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ResourceLogCategoryArrayOutput struct{ *pulumi.OutputState }

func (ResourceLogCategoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceLogCategory)(nil)).Elem()
}

func (o ResourceLogCategoryArrayOutput) ToResourceLogCategoryArrayOutput() ResourceLogCategoryArrayOutput {
	return o
}

func (o ResourceLogCategoryArrayOutput) ToResourceLogCategoryArrayOutputWithContext(ctx context.Context) ResourceLogCategoryArrayOutput {
	return o
}

func (o ResourceLogCategoryArrayOutput) Index(i pulumi.IntInput) ResourceLogCategoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceLogCategory {
		return vs[0].([]ResourceLogCategory)[vs[1].(int)]
	}).(ResourceLogCategoryOutput)
}

type ResourceLogCategoryResponse struct {
	Enabled *string `pulumi:"enabled"`
	Name    *string `pulumi:"name"`
}

type ResourceLogCategoryResponseOutput struct{ *pulumi.OutputState }

func (ResourceLogCategoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLogCategoryResponse)(nil)).Elem()
}

func (o ResourceLogCategoryResponseOutput) ToResourceLogCategoryResponseOutput() ResourceLogCategoryResponseOutput {
	return o
}

func (o ResourceLogCategoryResponseOutput) ToResourceLogCategoryResponseOutputWithContext(ctx context.Context) ResourceLogCategoryResponseOutput {
	return o
}

func (o ResourceLogCategoryResponseOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceLogCategoryResponse) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

func (o ResourceLogCategoryResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceLogCategoryResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ResourceLogCategoryResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceLogCategoryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceLogCategoryResponse)(nil)).Elem()
}

func (o ResourceLogCategoryResponseArrayOutput) ToResourceLogCategoryResponseArrayOutput() ResourceLogCategoryResponseArrayOutput {
	return o
}

func (o ResourceLogCategoryResponseArrayOutput) ToResourceLogCategoryResponseArrayOutputWithContext(ctx context.Context) ResourceLogCategoryResponseArrayOutput {
	return o
}

func (o ResourceLogCategoryResponseArrayOutput) Index(i pulumi.IntInput) ResourceLogCategoryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceLogCategoryResponse {
		return vs[0].([]ResourceLogCategoryResponse)[vs[1].(int)]
	}).(ResourceLogCategoryResponseOutput)
}

type ResourceLogConfiguration struct {
	Categories []ResourceLogCategory `pulumi:"categories"`
}





type ResourceLogConfigurationInput interface {
	pulumi.Input

	ToResourceLogConfigurationOutput() ResourceLogConfigurationOutput
	ToResourceLogConfigurationOutputWithContext(context.Context) ResourceLogConfigurationOutput
}

type ResourceLogConfigurationArgs struct {
	Categories ResourceLogCategoryArrayInput `pulumi:"categories"`
}

func (ResourceLogConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLogConfiguration)(nil)).Elem()
}

func (i ResourceLogConfigurationArgs) ToResourceLogConfigurationOutput() ResourceLogConfigurationOutput {
	return i.ToResourceLogConfigurationOutputWithContext(context.Background())
}

func (i ResourceLogConfigurationArgs) ToResourceLogConfigurationOutputWithContext(ctx context.Context) ResourceLogConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLogConfigurationOutput)
}

func (i ResourceLogConfigurationArgs) ToResourceLogConfigurationPtrOutput() ResourceLogConfigurationPtrOutput {
	return i.ToResourceLogConfigurationPtrOutputWithContext(context.Background())
}

func (i ResourceLogConfigurationArgs) ToResourceLogConfigurationPtrOutputWithContext(ctx context.Context) ResourceLogConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLogConfigurationOutput).ToResourceLogConfigurationPtrOutputWithContext(ctx)
}









type ResourceLogConfigurationPtrInput interface {
	pulumi.Input

	ToResourceLogConfigurationPtrOutput() ResourceLogConfigurationPtrOutput
	ToResourceLogConfigurationPtrOutputWithContext(context.Context) ResourceLogConfigurationPtrOutput
}

type resourceLogConfigurationPtrType ResourceLogConfigurationArgs

func ResourceLogConfigurationPtr(v *ResourceLogConfigurationArgs) ResourceLogConfigurationPtrInput {
	return (*resourceLogConfigurationPtrType)(v)
}

func (*resourceLogConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceLogConfiguration)(nil)).Elem()
}

func (i *resourceLogConfigurationPtrType) ToResourceLogConfigurationPtrOutput() ResourceLogConfigurationPtrOutput {
	return i.ToResourceLogConfigurationPtrOutputWithContext(context.Background())
}

func (i *resourceLogConfigurationPtrType) ToResourceLogConfigurationPtrOutputWithContext(ctx context.Context) ResourceLogConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLogConfigurationPtrOutput)
}

type ResourceLogConfigurationOutput struct{ *pulumi.OutputState }

func (ResourceLogConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLogConfiguration)(nil)).Elem()
}

func (o ResourceLogConfigurationOutput) ToResourceLogConfigurationOutput() ResourceLogConfigurationOutput {
	return o
}

func (o ResourceLogConfigurationOutput) ToResourceLogConfigurationOutputWithContext(ctx context.Context) ResourceLogConfigurationOutput {
	return o
}

func (o ResourceLogConfigurationOutput) ToResourceLogConfigurationPtrOutput() ResourceLogConfigurationPtrOutput {
	return o.ToResourceLogConfigurationPtrOutputWithContext(context.Background())
}

func (o ResourceLogConfigurationOutput) ToResourceLogConfigurationPtrOutputWithContext(ctx context.Context) ResourceLogConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceLogConfiguration) *ResourceLogConfiguration {
		return &v
	}).(ResourceLogConfigurationPtrOutput)
}

func (o ResourceLogConfigurationOutput) Categories() ResourceLogCategoryArrayOutput {
	return o.ApplyT(func(v ResourceLogConfiguration) []ResourceLogCategory { return v.Categories }).(ResourceLogCategoryArrayOutput)
}

type ResourceLogConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ResourceLogConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceLogConfiguration)(nil)).Elem()
}

func (o ResourceLogConfigurationPtrOutput) ToResourceLogConfigurationPtrOutput() ResourceLogConfigurationPtrOutput {
	return o
}

func (o ResourceLogConfigurationPtrOutput) ToResourceLogConfigurationPtrOutputWithContext(ctx context.Context) ResourceLogConfigurationPtrOutput {
	return o
}

func (o ResourceLogConfigurationPtrOutput) Elem() ResourceLogConfigurationOutput {
	return o.ApplyT(func(v *ResourceLogConfiguration) ResourceLogConfiguration {
		if v != nil {
			return *v
		}
		var ret ResourceLogConfiguration
		return ret
	}).(ResourceLogConfigurationOutput)
}

func (o ResourceLogConfigurationPtrOutput) Categories() ResourceLogCategoryArrayOutput {
	return o.ApplyT(func(v *ResourceLogConfiguration) []ResourceLogCategory {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(ResourceLogCategoryArrayOutput)
}

type ResourceLogConfigurationResponse struct {
	Categories []ResourceLogCategoryResponse `pulumi:"categories"`
}

type ResourceLogConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ResourceLogConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLogConfigurationResponse)(nil)).Elem()
}

func (o ResourceLogConfigurationResponseOutput) ToResourceLogConfigurationResponseOutput() ResourceLogConfigurationResponseOutput {
	return o
}

func (o ResourceLogConfigurationResponseOutput) ToResourceLogConfigurationResponseOutputWithContext(ctx context.Context) ResourceLogConfigurationResponseOutput {
	return o
}

func (o ResourceLogConfigurationResponseOutput) Categories() ResourceLogCategoryResponseArrayOutput {
	return o.ApplyT(func(v ResourceLogConfigurationResponse) []ResourceLogCategoryResponse { return v.Categories }).(ResourceLogCategoryResponseArrayOutput)
}

type ResourceLogConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceLogConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceLogConfigurationResponse)(nil)).Elem()
}

func (o ResourceLogConfigurationResponsePtrOutput) ToResourceLogConfigurationResponsePtrOutput() ResourceLogConfigurationResponsePtrOutput {
	return o
}

func (o ResourceLogConfigurationResponsePtrOutput) ToResourceLogConfigurationResponsePtrOutputWithContext(ctx context.Context) ResourceLogConfigurationResponsePtrOutput {
	return o
}

func (o ResourceLogConfigurationResponsePtrOutput) Elem() ResourceLogConfigurationResponseOutput {
	return o.ApplyT(func(v *ResourceLogConfigurationResponse) ResourceLogConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ResourceLogConfigurationResponse
		return ret
	}).(ResourceLogConfigurationResponseOutput)
}

func (o ResourceLogConfigurationResponsePtrOutput) Categories() ResourceLogCategoryResponseArrayOutput {
	return o.ApplyT(func(v *ResourceLogConfigurationResponse) []ResourceLogCategoryResponse {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(ResourceLogCategoryResponseArrayOutput)
}

type ResourceSku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type ResourceSkuInput interface {
	pulumi.Input

	ToResourceSkuOutput() ResourceSkuOutput
	ToResourceSkuOutputWithContext(context.Context) ResourceSkuOutput
}

type ResourceSkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (ResourceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (i ResourceSkuArgs) ToResourceSkuOutput() ResourceSkuOutput {
	return i.ToResourceSkuOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput)
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput).ToResourceSkuPtrOutputWithContext(ctx)
}









type ResourceSkuPtrInput interface {
	pulumi.Input

	ToResourceSkuPtrOutput() ResourceSkuPtrOutput
	ToResourceSkuPtrOutputWithContext(context.Context) ResourceSkuPtrOutput
}

type resourceSkuPtrType ResourceSkuArgs

func ResourceSkuPtr(v *ResourceSkuArgs) ResourceSkuPtrInput {
	return (*resourceSkuPtrType)(v)
}

func (*resourceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuPtrOutput)
}

type ResourceSkuOutput struct{ *pulumi.OutputState }

func (ResourceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (o ResourceSkuOutput) ToResourceSkuOutput() ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSku) *ResourceSku {
		return &v
	}).(ResourceSkuPtrOutput)
}

func (o ResourceSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ResourceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ResourceSkuPtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) Elem() ResourceSkuOutput {
	return o.ApplyT(func(v *ResourceSku) ResourceSku {
		if v != nil {
			return *v
		}
		var ret ResourceSku
		return ret
	}).(ResourceSkuOutput)
}

func (o ResourceSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ResourceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ResourceSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   string  `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     string  `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}

type ResourceSkuResponseOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutput() ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutputWithContext(ctx context.Context) ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ResourceSkuResponseOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSkuResponse) string { return v.Family }).(pulumi.StringOutput)
}

func (o ResourceSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceSkuResponseOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSkuResponse) string { return v.Size }).(pulumi.StringOutput)
}

func (o ResourceSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ResourceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) Elem() ResourceSkuResponseOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) ResourceSkuResponse {
		if v != nil {
			return *v
		}
		var ret ResourceSkuResponse
		return ret
	}).(ResourceSkuResponseOutput)
}

func (o ResourceSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Family
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Size
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SharedPrivateLinkResourceResponse struct {
	GroupId               string             `pulumi:"groupId"`
	Id                    string             `pulumi:"id"`
	Name                  string             `pulumi:"name"`
	PrivateLinkResourceId string             `pulumi:"privateLinkResourceId"`
	ProvisioningState     string             `pulumi:"provisioningState"`
	RequestMessage        *string            `pulumi:"requestMessage"`
	Status                string             `pulumi:"status"`
	SystemData            SystemDataResponse `pulumi:"systemData"`
	Type                  string             `pulumi:"type"`
}

type SharedPrivateLinkResourceResponseOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPrivateLinkResourceResponse)(nil)).Elem()
}

func (o SharedPrivateLinkResourceResponseOutput) ToSharedPrivateLinkResourceResponseOutput() SharedPrivateLinkResourceResponseOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseOutput) ToSharedPrivateLinkResourceResponseOutputWithContext(ctx context.Context) SharedPrivateLinkResourceResponseOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) string { return v.GroupId }).(pulumi.StringOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) PrivateLinkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) string { return v.PrivateLinkResourceId }).(pulumi.StringOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) *string { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SharedPrivateLinkResourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SharedPrivateLinkResourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type SharedPrivateLinkResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SharedPrivateLinkResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedPrivateLinkResourceResponse)(nil)).Elem()
}

func (o SharedPrivateLinkResourceResponseArrayOutput) ToSharedPrivateLinkResourceResponseArrayOutput() SharedPrivateLinkResourceResponseArrayOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseArrayOutput) ToSharedPrivateLinkResourceResponseArrayOutputWithContext(ctx context.Context) SharedPrivateLinkResourceResponseArrayOutput {
	return o
}

func (o SharedPrivateLinkResourceResponseArrayOutput) Index(i pulumi.IntInput) SharedPrivateLinkResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharedPrivateLinkResourceResponse {
		return vs[0].([]SharedPrivateLinkResourceResponse)[vs[1].(int)]
	}).(SharedPrivateLinkResourceResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type UpstreamAuthSettings struct {
	ManagedIdentity *ManagedIdentitySettings `pulumi:"managedIdentity"`
	Type            *string                  `pulumi:"type"`
}





type UpstreamAuthSettingsInput interface {
	pulumi.Input

	ToUpstreamAuthSettingsOutput() UpstreamAuthSettingsOutput
	ToUpstreamAuthSettingsOutputWithContext(context.Context) UpstreamAuthSettingsOutput
}

type UpstreamAuthSettingsArgs struct {
	ManagedIdentity ManagedIdentitySettingsPtrInput `pulumi:"managedIdentity"`
	Type            pulumi.StringPtrInput           `pulumi:"type"`
}

func (UpstreamAuthSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UpstreamAuthSettings)(nil)).Elem()
}

func (i UpstreamAuthSettingsArgs) ToUpstreamAuthSettingsOutput() UpstreamAuthSettingsOutput {
	return i.ToUpstreamAuthSettingsOutputWithContext(context.Background())
}

func (i UpstreamAuthSettingsArgs) ToUpstreamAuthSettingsOutputWithContext(ctx context.Context) UpstreamAuthSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpstreamAuthSettingsOutput)
}

func (i UpstreamAuthSettingsArgs) ToUpstreamAuthSettingsPtrOutput() UpstreamAuthSettingsPtrOutput {
	return i.ToUpstreamAuthSettingsPtrOutputWithContext(context.Background())
}

func (i UpstreamAuthSettingsArgs) ToUpstreamAuthSettingsPtrOutputWithContext(ctx context.Context) UpstreamAuthSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpstreamAuthSettingsOutput).ToUpstreamAuthSettingsPtrOutputWithContext(ctx)
}









type UpstreamAuthSettingsPtrInput interface {
	pulumi.Input

	ToUpstreamAuthSettingsPtrOutput() UpstreamAuthSettingsPtrOutput
	ToUpstreamAuthSettingsPtrOutputWithContext(context.Context) UpstreamAuthSettingsPtrOutput
}

type upstreamAuthSettingsPtrType UpstreamAuthSettingsArgs

func UpstreamAuthSettingsPtr(v *UpstreamAuthSettingsArgs) UpstreamAuthSettingsPtrInput {
	return (*upstreamAuthSettingsPtrType)(v)
}

func (*upstreamAuthSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UpstreamAuthSettings)(nil)).Elem()
}

func (i *upstreamAuthSettingsPtrType) ToUpstreamAuthSettingsPtrOutput() UpstreamAuthSettingsPtrOutput {
	return i.ToUpstreamAuthSettingsPtrOutputWithContext(context.Background())
}

func (i *upstreamAuthSettingsPtrType) ToUpstreamAuthSettingsPtrOutputWithContext(ctx context.Context) UpstreamAuthSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpstreamAuthSettingsPtrOutput)
}

type UpstreamAuthSettingsOutput struct{ *pulumi.OutputState }

func (UpstreamAuthSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpstreamAuthSettings)(nil)).Elem()
}

func (o UpstreamAuthSettingsOutput) ToUpstreamAuthSettingsOutput() UpstreamAuthSettingsOutput {
	return o
}

func (o UpstreamAuthSettingsOutput) ToUpstreamAuthSettingsOutputWithContext(ctx context.Context) UpstreamAuthSettingsOutput {
	return o
}

func (o UpstreamAuthSettingsOutput) ToUpstreamAuthSettingsPtrOutput() UpstreamAuthSettingsPtrOutput {
	return o.ToUpstreamAuthSettingsPtrOutputWithContext(context.Background())
}

func (o UpstreamAuthSettingsOutput) ToUpstreamAuthSettingsPtrOutputWithContext(ctx context.Context) UpstreamAuthSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UpstreamAuthSettings) *UpstreamAuthSettings {
		return &v
	}).(UpstreamAuthSettingsPtrOutput)
}

func (o UpstreamAuthSettingsOutput) ManagedIdentity() ManagedIdentitySettingsPtrOutput {
	return o.ApplyT(func(v UpstreamAuthSettings) *ManagedIdentitySettings { return v.ManagedIdentity }).(ManagedIdentitySettingsPtrOutput)
}

func (o UpstreamAuthSettingsOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpstreamAuthSettings) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type UpstreamAuthSettingsPtrOutput struct{ *pulumi.OutputState }

func (UpstreamAuthSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpstreamAuthSettings)(nil)).Elem()
}

func (o UpstreamAuthSettingsPtrOutput) ToUpstreamAuthSettingsPtrOutput() UpstreamAuthSettingsPtrOutput {
	return o
}

func (o UpstreamAuthSettingsPtrOutput) ToUpstreamAuthSettingsPtrOutputWithContext(ctx context.Context) UpstreamAuthSettingsPtrOutput {
	return o
}

func (o UpstreamAuthSettingsPtrOutput) Elem() UpstreamAuthSettingsOutput {
	return o.ApplyT(func(v *UpstreamAuthSettings) UpstreamAuthSettings {
		if v != nil {
			return *v
		}
		var ret UpstreamAuthSettings
		return ret
	}).(UpstreamAuthSettingsOutput)
}

func (o UpstreamAuthSettingsPtrOutput) ManagedIdentity() ManagedIdentitySettingsPtrOutput {
	return o.ApplyT(func(v *UpstreamAuthSettings) *ManagedIdentitySettings {
		if v == nil {
			return nil
		}
		return v.ManagedIdentity
	}).(ManagedIdentitySettingsPtrOutput)
}

func (o UpstreamAuthSettingsPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpstreamAuthSettings) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type UpstreamAuthSettingsResponse struct {
	ManagedIdentity *ManagedIdentitySettingsResponse `pulumi:"managedIdentity"`
	Type            *string                          `pulumi:"type"`
}

type UpstreamAuthSettingsResponseOutput struct{ *pulumi.OutputState }

func (UpstreamAuthSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpstreamAuthSettingsResponse)(nil)).Elem()
}

func (o UpstreamAuthSettingsResponseOutput) ToUpstreamAuthSettingsResponseOutput() UpstreamAuthSettingsResponseOutput {
	return o
}

func (o UpstreamAuthSettingsResponseOutput) ToUpstreamAuthSettingsResponseOutputWithContext(ctx context.Context) UpstreamAuthSettingsResponseOutput {
	return o
}

func (o UpstreamAuthSettingsResponseOutput) ManagedIdentity() ManagedIdentitySettingsResponsePtrOutput {
	return o.ApplyT(func(v UpstreamAuthSettingsResponse) *ManagedIdentitySettingsResponse { return v.ManagedIdentity }).(ManagedIdentitySettingsResponsePtrOutput)
}

func (o UpstreamAuthSettingsResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpstreamAuthSettingsResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type UpstreamAuthSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (UpstreamAuthSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpstreamAuthSettingsResponse)(nil)).Elem()
}

func (o UpstreamAuthSettingsResponsePtrOutput) ToUpstreamAuthSettingsResponsePtrOutput() UpstreamAuthSettingsResponsePtrOutput {
	return o
}

func (o UpstreamAuthSettingsResponsePtrOutput) ToUpstreamAuthSettingsResponsePtrOutputWithContext(ctx context.Context) UpstreamAuthSettingsResponsePtrOutput {
	return o
}

func (o UpstreamAuthSettingsResponsePtrOutput) Elem() UpstreamAuthSettingsResponseOutput {
	return o.ApplyT(func(v *UpstreamAuthSettingsResponse) UpstreamAuthSettingsResponse {
		if v != nil {
			return *v
		}
		var ret UpstreamAuthSettingsResponse
		return ret
	}).(UpstreamAuthSettingsResponseOutput)
}

func (o UpstreamAuthSettingsResponsePtrOutput) ManagedIdentity() ManagedIdentitySettingsResponsePtrOutput {
	return o.ApplyT(func(v *UpstreamAuthSettingsResponse) *ManagedIdentitySettingsResponse {
		if v == nil {
			return nil
		}
		return v.ManagedIdentity
	}).(ManagedIdentitySettingsResponsePtrOutput)
}

func (o UpstreamAuthSettingsResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpstreamAuthSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityPropertyResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserAssignedIdentityPropertyResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityPropertyResponse)(nil)).Elem()
}

func (o UserAssignedIdentityPropertyResponseOutput) ToUserAssignedIdentityPropertyResponseOutput() UserAssignedIdentityPropertyResponseOutput {
	return o
}

func (o UserAssignedIdentityPropertyResponseOutput) ToUserAssignedIdentityPropertyResponseOutputWithContext(ctx context.Context) UserAssignedIdentityPropertyResponseOutput {
	return o
}

func (o UserAssignedIdentityPropertyResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityPropertyResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityPropertyResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityPropertyResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentityPropertyResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityPropertyResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityPropertyResponse)(nil)).Elem()
}

func (o UserAssignedIdentityPropertyResponseMapOutput) ToUserAssignedIdentityPropertyResponseMapOutput() UserAssignedIdentityPropertyResponseMapOutput {
	return o
}

func (o UserAssignedIdentityPropertyResponseMapOutput) ToUserAssignedIdentityPropertyResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityPropertyResponseMapOutput {
	return o
}

func (o UserAssignedIdentityPropertyResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityPropertyResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityPropertyResponse {
		return vs[0].(map[string]UserAssignedIdentityPropertyResponse)[vs[1].(string)]
	}).(UserAssignedIdentityPropertyResponseOutput)
}

type WebPubSubHubProperties struct {
	AnonymousConnectPolicy *string        `pulumi:"anonymousConnectPolicy"`
	EventHandlers          []EventHandler `pulumi:"eventHandlers"`
}


func (val *WebPubSubHubProperties) Defaults() *WebPubSubHubProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AnonymousConnectPolicy) {
		anonymousConnectPolicy_ := "deny"
		tmp.AnonymousConnectPolicy = &anonymousConnectPolicy_
	}
	return &tmp
}





type WebPubSubHubPropertiesInput interface {
	pulumi.Input

	ToWebPubSubHubPropertiesOutput() WebPubSubHubPropertiesOutput
	ToWebPubSubHubPropertiesOutputWithContext(context.Context) WebPubSubHubPropertiesOutput
}

type WebPubSubHubPropertiesArgs struct {
	AnonymousConnectPolicy pulumi.StringPtrInput  `pulumi:"anonymousConnectPolicy"`
	EventHandlers          EventHandlerArrayInput `pulumi:"eventHandlers"`
}


func (val *WebPubSubHubPropertiesArgs) Defaults() *WebPubSubHubPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AnonymousConnectPolicy) {
		tmp.AnonymousConnectPolicy = pulumi.StringPtr("deny")
	}
	return &tmp
}
func (WebPubSubHubPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubHubProperties)(nil)).Elem()
}

func (i WebPubSubHubPropertiesArgs) ToWebPubSubHubPropertiesOutput() WebPubSubHubPropertiesOutput {
	return i.ToWebPubSubHubPropertiesOutputWithContext(context.Background())
}

func (i WebPubSubHubPropertiesArgs) ToWebPubSubHubPropertiesOutputWithContext(ctx context.Context) WebPubSubHubPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubHubPropertiesOutput)
}

type WebPubSubHubPropertiesOutput struct{ *pulumi.OutputState }

func (WebPubSubHubPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubHubProperties)(nil)).Elem()
}

func (o WebPubSubHubPropertiesOutput) ToWebPubSubHubPropertiesOutput() WebPubSubHubPropertiesOutput {
	return o
}

func (o WebPubSubHubPropertiesOutput) ToWebPubSubHubPropertiesOutputWithContext(ctx context.Context) WebPubSubHubPropertiesOutput {
	return o
}

func (o WebPubSubHubPropertiesOutput) AnonymousConnectPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebPubSubHubProperties) *string { return v.AnonymousConnectPolicy }).(pulumi.StringPtrOutput)
}

func (o WebPubSubHubPropertiesOutput) EventHandlers() EventHandlerArrayOutput {
	return o.ApplyT(func(v WebPubSubHubProperties) []EventHandler { return v.EventHandlers }).(EventHandlerArrayOutput)
}

type WebPubSubHubPropertiesResponse struct {
	AnonymousConnectPolicy *string                `pulumi:"anonymousConnectPolicy"`
	EventHandlers          []EventHandlerResponse `pulumi:"eventHandlers"`
}


func (val *WebPubSubHubPropertiesResponse) Defaults() *WebPubSubHubPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AnonymousConnectPolicy) {
		anonymousConnectPolicy_ := "deny"
		tmp.AnonymousConnectPolicy = &anonymousConnectPolicy_
	}
	return &tmp
}

type WebPubSubHubPropertiesResponseOutput struct{ *pulumi.OutputState }

func (WebPubSubHubPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubHubPropertiesResponse)(nil)).Elem()
}

func (o WebPubSubHubPropertiesResponseOutput) ToWebPubSubHubPropertiesResponseOutput() WebPubSubHubPropertiesResponseOutput {
	return o
}

func (o WebPubSubHubPropertiesResponseOutput) ToWebPubSubHubPropertiesResponseOutputWithContext(ctx context.Context) WebPubSubHubPropertiesResponseOutput {
	return o
}

func (o WebPubSubHubPropertiesResponseOutput) AnonymousConnectPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebPubSubHubPropertiesResponse) *string { return v.AnonymousConnectPolicy }).(pulumi.StringPtrOutput)
}

func (o WebPubSubHubPropertiesResponseOutput) EventHandlers() EventHandlerResponseArrayOutput {
	return o.ApplyT(func(v WebPubSubHubPropertiesResponse) []EventHandlerResponse { return v.EventHandlers }).(EventHandlerResponseArrayOutput)
}

type WebPubSubNetworkACLs struct {
	DefaultAction    *string              `pulumi:"defaultAction"`
	PrivateEndpoints []PrivateEndpointACL `pulumi:"privateEndpoints"`
	PublicNetwork    *NetworkACL          `pulumi:"publicNetwork"`
}





type WebPubSubNetworkACLsInput interface {
	pulumi.Input

	ToWebPubSubNetworkACLsOutput() WebPubSubNetworkACLsOutput
	ToWebPubSubNetworkACLsOutputWithContext(context.Context) WebPubSubNetworkACLsOutput
}

type WebPubSubNetworkACLsArgs struct {
	DefaultAction    pulumi.StringPtrInput        `pulumi:"defaultAction"`
	PrivateEndpoints PrivateEndpointACLArrayInput `pulumi:"privateEndpoints"`
	PublicNetwork    NetworkACLPtrInput           `pulumi:"publicNetwork"`
}

func (WebPubSubNetworkACLsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubNetworkACLs)(nil)).Elem()
}

func (i WebPubSubNetworkACLsArgs) ToWebPubSubNetworkACLsOutput() WebPubSubNetworkACLsOutput {
	return i.ToWebPubSubNetworkACLsOutputWithContext(context.Background())
}

func (i WebPubSubNetworkACLsArgs) ToWebPubSubNetworkACLsOutputWithContext(ctx context.Context) WebPubSubNetworkACLsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubNetworkACLsOutput)
}

func (i WebPubSubNetworkACLsArgs) ToWebPubSubNetworkACLsPtrOutput() WebPubSubNetworkACLsPtrOutput {
	return i.ToWebPubSubNetworkACLsPtrOutputWithContext(context.Background())
}

func (i WebPubSubNetworkACLsArgs) ToWebPubSubNetworkACLsPtrOutputWithContext(ctx context.Context) WebPubSubNetworkACLsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubNetworkACLsOutput).ToWebPubSubNetworkACLsPtrOutputWithContext(ctx)
}









type WebPubSubNetworkACLsPtrInput interface {
	pulumi.Input

	ToWebPubSubNetworkACLsPtrOutput() WebPubSubNetworkACLsPtrOutput
	ToWebPubSubNetworkACLsPtrOutputWithContext(context.Context) WebPubSubNetworkACLsPtrOutput
}

type webPubSubNetworkACLsPtrType WebPubSubNetworkACLsArgs

func WebPubSubNetworkACLsPtr(v *WebPubSubNetworkACLsArgs) WebPubSubNetworkACLsPtrInput {
	return (*webPubSubNetworkACLsPtrType)(v)
}

func (*webPubSubNetworkACLsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubNetworkACLs)(nil)).Elem()
}

func (i *webPubSubNetworkACLsPtrType) ToWebPubSubNetworkACLsPtrOutput() WebPubSubNetworkACLsPtrOutput {
	return i.ToWebPubSubNetworkACLsPtrOutputWithContext(context.Background())
}

func (i *webPubSubNetworkACLsPtrType) ToWebPubSubNetworkACLsPtrOutputWithContext(ctx context.Context) WebPubSubNetworkACLsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubNetworkACLsPtrOutput)
}

type WebPubSubNetworkACLsOutput struct{ *pulumi.OutputState }

func (WebPubSubNetworkACLsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubNetworkACLs)(nil)).Elem()
}

func (o WebPubSubNetworkACLsOutput) ToWebPubSubNetworkACLsOutput() WebPubSubNetworkACLsOutput {
	return o
}

func (o WebPubSubNetworkACLsOutput) ToWebPubSubNetworkACLsOutputWithContext(ctx context.Context) WebPubSubNetworkACLsOutput {
	return o
}

func (o WebPubSubNetworkACLsOutput) ToWebPubSubNetworkACLsPtrOutput() WebPubSubNetworkACLsPtrOutput {
	return o.ToWebPubSubNetworkACLsPtrOutputWithContext(context.Background())
}

func (o WebPubSubNetworkACLsOutput) ToWebPubSubNetworkACLsPtrOutputWithContext(ctx context.Context) WebPubSubNetworkACLsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebPubSubNetworkACLs) *WebPubSubNetworkACLs {
		return &v
	}).(WebPubSubNetworkACLsPtrOutput)
}

func (o WebPubSubNetworkACLsOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebPubSubNetworkACLs) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o WebPubSubNetworkACLsOutput) PrivateEndpoints() PrivateEndpointACLArrayOutput {
	return o.ApplyT(func(v WebPubSubNetworkACLs) []PrivateEndpointACL { return v.PrivateEndpoints }).(PrivateEndpointACLArrayOutput)
}

func (o WebPubSubNetworkACLsOutput) PublicNetwork() NetworkACLPtrOutput {
	return o.ApplyT(func(v WebPubSubNetworkACLs) *NetworkACL { return v.PublicNetwork }).(NetworkACLPtrOutput)
}

type WebPubSubNetworkACLsPtrOutput struct{ *pulumi.OutputState }

func (WebPubSubNetworkACLsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubNetworkACLs)(nil)).Elem()
}

func (o WebPubSubNetworkACLsPtrOutput) ToWebPubSubNetworkACLsPtrOutput() WebPubSubNetworkACLsPtrOutput {
	return o
}

func (o WebPubSubNetworkACLsPtrOutput) ToWebPubSubNetworkACLsPtrOutputWithContext(ctx context.Context) WebPubSubNetworkACLsPtrOutput {
	return o
}

func (o WebPubSubNetworkACLsPtrOutput) Elem() WebPubSubNetworkACLsOutput {
	return o.ApplyT(func(v *WebPubSubNetworkACLs) WebPubSubNetworkACLs {
		if v != nil {
			return *v
		}
		var ret WebPubSubNetworkACLs
		return ret
	}).(WebPubSubNetworkACLsOutput)
}

func (o WebPubSubNetworkACLsPtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebPubSubNetworkACLs) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o WebPubSubNetworkACLsPtrOutput) PrivateEndpoints() PrivateEndpointACLArrayOutput {
	return o.ApplyT(func(v *WebPubSubNetworkACLs) []PrivateEndpointACL {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoints
	}).(PrivateEndpointACLArrayOutput)
}

func (o WebPubSubNetworkACLsPtrOutput) PublicNetwork() NetworkACLPtrOutput {
	return o.ApplyT(func(v *WebPubSubNetworkACLs) *NetworkACL {
		if v == nil {
			return nil
		}
		return v.PublicNetwork
	}).(NetworkACLPtrOutput)
}

type WebPubSubNetworkACLsResponse struct {
	DefaultAction    *string                      `pulumi:"defaultAction"`
	PrivateEndpoints []PrivateEndpointACLResponse `pulumi:"privateEndpoints"`
	PublicNetwork    *NetworkACLResponse          `pulumi:"publicNetwork"`
}

type WebPubSubNetworkACLsResponseOutput struct{ *pulumi.OutputState }

func (WebPubSubNetworkACLsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubNetworkACLsResponse)(nil)).Elem()
}

func (o WebPubSubNetworkACLsResponseOutput) ToWebPubSubNetworkACLsResponseOutput() WebPubSubNetworkACLsResponseOutput {
	return o
}

func (o WebPubSubNetworkACLsResponseOutput) ToWebPubSubNetworkACLsResponseOutputWithContext(ctx context.Context) WebPubSubNetworkACLsResponseOutput {
	return o
}

func (o WebPubSubNetworkACLsResponseOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebPubSubNetworkACLsResponse) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o WebPubSubNetworkACLsResponseOutput) PrivateEndpoints() PrivateEndpointACLResponseArrayOutput {
	return o.ApplyT(func(v WebPubSubNetworkACLsResponse) []PrivateEndpointACLResponse { return v.PrivateEndpoints }).(PrivateEndpointACLResponseArrayOutput)
}

func (o WebPubSubNetworkACLsResponseOutput) PublicNetwork() NetworkACLResponsePtrOutput {
	return o.ApplyT(func(v WebPubSubNetworkACLsResponse) *NetworkACLResponse { return v.PublicNetwork }).(NetworkACLResponsePtrOutput)
}

type WebPubSubNetworkACLsResponsePtrOutput struct{ *pulumi.OutputState }

func (WebPubSubNetworkACLsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubNetworkACLsResponse)(nil)).Elem()
}

func (o WebPubSubNetworkACLsResponsePtrOutput) ToWebPubSubNetworkACLsResponsePtrOutput() WebPubSubNetworkACLsResponsePtrOutput {
	return o
}

func (o WebPubSubNetworkACLsResponsePtrOutput) ToWebPubSubNetworkACLsResponsePtrOutputWithContext(ctx context.Context) WebPubSubNetworkACLsResponsePtrOutput {
	return o
}

func (o WebPubSubNetworkACLsResponsePtrOutput) Elem() WebPubSubNetworkACLsResponseOutput {
	return o.ApplyT(func(v *WebPubSubNetworkACLsResponse) WebPubSubNetworkACLsResponse {
		if v != nil {
			return *v
		}
		var ret WebPubSubNetworkACLsResponse
		return ret
	}).(WebPubSubNetworkACLsResponseOutput)
}

func (o WebPubSubNetworkACLsResponsePtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebPubSubNetworkACLsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o WebPubSubNetworkACLsResponsePtrOutput) PrivateEndpoints() PrivateEndpointACLResponseArrayOutput {
	return o.ApplyT(func(v *WebPubSubNetworkACLsResponse) []PrivateEndpointACLResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoints
	}).(PrivateEndpointACLResponseArrayOutput)
}

func (o WebPubSubNetworkACLsResponsePtrOutput) PublicNetwork() NetworkACLResponsePtrOutput {
	return o.ApplyT(func(v *WebPubSubNetworkACLsResponse) *NetworkACLResponse {
		if v == nil {
			return nil
		}
		return v.PublicNetwork
	}).(NetworkACLResponsePtrOutput)
}

type WebPubSubTlsSettings struct {
	ClientCertEnabled *bool `pulumi:"clientCertEnabled"`
}


func (val *WebPubSubTlsSettings) Defaults() *WebPubSubTlsSettings {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ClientCertEnabled) {
		clientCertEnabled_ := true
		tmp.ClientCertEnabled = &clientCertEnabled_
	}
	return &tmp
}





type WebPubSubTlsSettingsInput interface {
	pulumi.Input

	ToWebPubSubTlsSettingsOutput() WebPubSubTlsSettingsOutput
	ToWebPubSubTlsSettingsOutputWithContext(context.Context) WebPubSubTlsSettingsOutput
}

type WebPubSubTlsSettingsArgs struct {
	ClientCertEnabled pulumi.BoolPtrInput `pulumi:"clientCertEnabled"`
}


func (val *WebPubSubTlsSettingsArgs) Defaults() *WebPubSubTlsSettingsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ClientCertEnabled) {
		tmp.ClientCertEnabled = pulumi.BoolPtr(true)
	}
	return &tmp
}
func (WebPubSubTlsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubTlsSettings)(nil)).Elem()
}

func (i WebPubSubTlsSettingsArgs) ToWebPubSubTlsSettingsOutput() WebPubSubTlsSettingsOutput {
	return i.ToWebPubSubTlsSettingsOutputWithContext(context.Background())
}

func (i WebPubSubTlsSettingsArgs) ToWebPubSubTlsSettingsOutputWithContext(ctx context.Context) WebPubSubTlsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubTlsSettingsOutput)
}

func (i WebPubSubTlsSettingsArgs) ToWebPubSubTlsSettingsPtrOutput() WebPubSubTlsSettingsPtrOutput {
	return i.ToWebPubSubTlsSettingsPtrOutputWithContext(context.Background())
}

func (i WebPubSubTlsSettingsArgs) ToWebPubSubTlsSettingsPtrOutputWithContext(ctx context.Context) WebPubSubTlsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubTlsSettingsOutput).ToWebPubSubTlsSettingsPtrOutputWithContext(ctx)
}









type WebPubSubTlsSettingsPtrInput interface {
	pulumi.Input

	ToWebPubSubTlsSettingsPtrOutput() WebPubSubTlsSettingsPtrOutput
	ToWebPubSubTlsSettingsPtrOutputWithContext(context.Context) WebPubSubTlsSettingsPtrOutput
}

type webPubSubTlsSettingsPtrType WebPubSubTlsSettingsArgs

func WebPubSubTlsSettingsPtr(v *WebPubSubTlsSettingsArgs) WebPubSubTlsSettingsPtrInput {
	return (*webPubSubTlsSettingsPtrType)(v)
}

func (*webPubSubTlsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubTlsSettings)(nil)).Elem()
}

func (i *webPubSubTlsSettingsPtrType) ToWebPubSubTlsSettingsPtrOutput() WebPubSubTlsSettingsPtrOutput {
	return i.ToWebPubSubTlsSettingsPtrOutputWithContext(context.Background())
}

func (i *webPubSubTlsSettingsPtrType) ToWebPubSubTlsSettingsPtrOutputWithContext(ctx context.Context) WebPubSubTlsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubTlsSettingsPtrOutput)
}

type WebPubSubTlsSettingsOutput struct{ *pulumi.OutputState }

func (WebPubSubTlsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubTlsSettings)(nil)).Elem()
}

func (o WebPubSubTlsSettingsOutput) ToWebPubSubTlsSettingsOutput() WebPubSubTlsSettingsOutput {
	return o
}

func (o WebPubSubTlsSettingsOutput) ToWebPubSubTlsSettingsOutputWithContext(ctx context.Context) WebPubSubTlsSettingsOutput {
	return o
}

func (o WebPubSubTlsSettingsOutput) ToWebPubSubTlsSettingsPtrOutput() WebPubSubTlsSettingsPtrOutput {
	return o.ToWebPubSubTlsSettingsPtrOutputWithContext(context.Background())
}

func (o WebPubSubTlsSettingsOutput) ToWebPubSubTlsSettingsPtrOutputWithContext(ctx context.Context) WebPubSubTlsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebPubSubTlsSettings) *WebPubSubTlsSettings {
		return &v
	}).(WebPubSubTlsSettingsPtrOutput)
}

func (o WebPubSubTlsSettingsOutput) ClientCertEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebPubSubTlsSettings) *bool { return v.ClientCertEnabled }).(pulumi.BoolPtrOutput)
}

type WebPubSubTlsSettingsPtrOutput struct{ *pulumi.OutputState }

func (WebPubSubTlsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubTlsSettings)(nil)).Elem()
}

func (o WebPubSubTlsSettingsPtrOutput) ToWebPubSubTlsSettingsPtrOutput() WebPubSubTlsSettingsPtrOutput {
	return o
}

func (o WebPubSubTlsSettingsPtrOutput) ToWebPubSubTlsSettingsPtrOutputWithContext(ctx context.Context) WebPubSubTlsSettingsPtrOutput {
	return o
}

func (o WebPubSubTlsSettingsPtrOutput) Elem() WebPubSubTlsSettingsOutput {
	return o.ApplyT(func(v *WebPubSubTlsSettings) WebPubSubTlsSettings {
		if v != nil {
			return *v
		}
		var ret WebPubSubTlsSettings
		return ret
	}).(WebPubSubTlsSettingsOutput)
}

func (o WebPubSubTlsSettingsPtrOutput) ClientCertEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebPubSubTlsSettings) *bool {
		if v == nil {
			return nil
		}
		return v.ClientCertEnabled
	}).(pulumi.BoolPtrOutput)
}

type WebPubSubTlsSettingsResponse struct {
	ClientCertEnabled *bool `pulumi:"clientCertEnabled"`
}


func (val *WebPubSubTlsSettingsResponse) Defaults() *WebPubSubTlsSettingsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ClientCertEnabled) {
		clientCertEnabled_ := true
		tmp.ClientCertEnabled = &clientCertEnabled_
	}
	return &tmp
}

type WebPubSubTlsSettingsResponseOutput struct{ *pulumi.OutputState }

func (WebPubSubTlsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebPubSubTlsSettingsResponse)(nil)).Elem()
}

func (o WebPubSubTlsSettingsResponseOutput) ToWebPubSubTlsSettingsResponseOutput() WebPubSubTlsSettingsResponseOutput {
	return o
}

func (o WebPubSubTlsSettingsResponseOutput) ToWebPubSubTlsSettingsResponseOutputWithContext(ctx context.Context) WebPubSubTlsSettingsResponseOutput {
	return o
}

func (o WebPubSubTlsSettingsResponseOutput) ClientCertEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebPubSubTlsSettingsResponse) *bool { return v.ClientCertEnabled }).(pulumi.BoolPtrOutput)
}

type WebPubSubTlsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (WebPubSubTlsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubTlsSettingsResponse)(nil)).Elem()
}

func (o WebPubSubTlsSettingsResponsePtrOutput) ToWebPubSubTlsSettingsResponsePtrOutput() WebPubSubTlsSettingsResponsePtrOutput {
	return o
}

func (o WebPubSubTlsSettingsResponsePtrOutput) ToWebPubSubTlsSettingsResponsePtrOutputWithContext(ctx context.Context) WebPubSubTlsSettingsResponsePtrOutput {
	return o
}

func (o WebPubSubTlsSettingsResponsePtrOutput) Elem() WebPubSubTlsSettingsResponseOutput {
	return o.ApplyT(func(v *WebPubSubTlsSettingsResponse) WebPubSubTlsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret WebPubSubTlsSettingsResponse
		return ret
	}).(WebPubSubTlsSettingsResponseOutput)
}

func (o WebPubSubTlsSettingsResponsePtrOutput) ClientCertEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebPubSubTlsSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ClientCertEnabled
	}).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EventHandlerOutput{})
	pulumi.RegisterOutputType(EventHandlerArrayOutput{})
	pulumi.RegisterOutputType(EventHandlerResponseOutput{})
	pulumi.RegisterOutputType(EventHandlerResponseArrayOutput{})
	pulumi.RegisterOutputType(LiveTraceCategoryOutput{})
	pulumi.RegisterOutputType(LiveTraceCategoryArrayOutput{})
	pulumi.RegisterOutputType(LiveTraceCategoryResponseOutput{})
	pulumi.RegisterOutputType(LiveTraceCategoryResponseArrayOutput{})
	pulumi.RegisterOutputType(LiveTraceConfigurationOutput{})
	pulumi.RegisterOutputType(LiveTraceConfigurationPtrOutput{})
	pulumi.RegisterOutputType(LiveTraceConfigurationResponseOutput{})
	pulumi.RegisterOutputType(LiveTraceConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityOutput{})
	pulumi.RegisterOutputType(ManagedIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentitySettingsOutput{})
	pulumi.RegisterOutputType(ManagedIdentitySettingsPtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentitySettingsResponseOutput{})
	pulumi.RegisterOutputType(ManagedIdentitySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkACLOutput{})
	pulumi.RegisterOutputType(NetworkACLPtrOutput{})
	pulumi.RegisterOutputType(NetworkACLResponseOutput{})
	pulumi.RegisterOutputType(NetworkACLResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointACLOutput{})
	pulumi.RegisterOutputType(PrivateEndpointACLArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointACLResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointACLResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceLogCategoryOutput{})
	pulumi.RegisterOutputType(ResourceLogCategoryArrayOutput{})
	pulumi.RegisterOutputType(ResourceLogCategoryResponseOutput{})
	pulumi.RegisterOutputType(ResourceLogCategoryResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceLogConfigurationOutput{})
	pulumi.RegisterOutputType(ResourceLogConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ResourceLogConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ResourceLogConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuOutput{})
	pulumi.RegisterOutputType(ResourceSkuPtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponseOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceResponseOutput{})
	pulumi.RegisterOutputType(SharedPrivateLinkResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UpstreamAuthSettingsOutput{})
	pulumi.RegisterOutputType(UpstreamAuthSettingsPtrOutput{})
	pulumi.RegisterOutputType(UpstreamAuthSettingsResponseOutput{})
	pulumi.RegisterOutputType(UpstreamAuthSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityPropertyResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityPropertyResponseMapOutput{})
	pulumi.RegisterOutputType(WebPubSubHubPropertiesOutput{})
	pulumi.RegisterOutputType(WebPubSubHubPropertiesResponseOutput{})
	pulumi.RegisterOutputType(WebPubSubNetworkACLsOutput{})
	pulumi.RegisterOutputType(WebPubSubNetworkACLsPtrOutput{})
	pulumi.RegisterOutputType(WebPubSubNetworkACLsResponseOutput{})
	pulumi.RegisterOutputType(WebPubSubNetworkACLsResponsePtrOutput{})
	pulumi.RegisterOutputType(WebPubSubTlsSettingsOutput{})
	pulumi.RegisterOutputType(WebPubSubTlsSettingsPtrOutput{})
	pulumi.RegisterOutputType(WebPubSubTlsSettingsResponseOutput{})
	pulumi.RegisterOutputType(WebPubSubTlsSettingsResponsePtrOutput{})
}
