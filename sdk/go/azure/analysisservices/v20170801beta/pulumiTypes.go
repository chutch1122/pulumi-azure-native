


package v20170801beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GatewayDetails struct {
	GatewayResourceId *string `pulumi:"gatewayResourceId"`
}





type GatewayDetailsInput interface {
	pulumi.Input

	ToGatewayDetailsOutput() GatewayDetailsOutput
	ToGatewayDetailsOutputWithContext(context.Context) GatewayDetailsOutput
}

type GatewayDetailsArgs struct {
	GatewayResourceId pulumi.StringPtrInput `pulumi:"gatewayResourceId"`
}

func (GatewayDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayDetails)(nil)).Elem()
}

func (i GatewayDetailsArgs) ToGatewayDetailsOutput() GatewayDetailsOutput {
	return i.ToGatewayDetailsOutputWithContext(context.Background())
}

func (i GatewayDetailsArgs) ToGatewayDetailsOutputWithContext(ctx context.Context) GatewayDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayDetailsOutput)
}

func (i GatewayDetailsArgs) ToGatewayDetailsPtrOutput() GatewayDetailsPtrOutput {
	return i.ToGatewayDetailsPtrOutputWithContext(context.Background())
}

func (i GatewayDetailsArgs) ToGatewayDetailsPtrOutputWithContext(ctx context.Context) GatewayDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayDetailsOutput).ToGatewayDetailsPtrOutputWithContext(ctx)
}









type GatewayDetailsPtrInput interface {
	pulumi.Input

	ToGatewayDetailsPtrOutput() GatewayDetailsPtrOutput
	ToGatewayDetailsPtrOutputWithContext(context.Context) GatewayDetailsPtrOutput
}

type gatewayDetailsPtrType GatewayDetailsArgs

func GatewayDetailsPtr(v *GatewayDetailsArgs) GatewayDetailsPtrInput {
	return (*gatewayDetailsPtrType)(v)
}

func (*gatewayDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayDetails)(nil)).Elem()
}

func (i *gatewayDetailsPtrType) ToGatewayDetailsPtrOutput() GatewayDetailsPtrOutput {
	return i.ToGatewayDetailsPtrOutputWithContext(context.Background())
}

func (i *gatewayDetailsPtrType) ToGatewayDetailsPtrOutputWithContext(ctx context.Context) GatewayDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayDetailsPtrOutput)
}

type GatewayDetailsOutput struct{ *pulumi.OutputState }

func (GatewayDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayDetails)(nil)).Elem()
}

func (o GatewayDetailsOutput) ToGatewayDetailsOutput() GatewayDetailsOutput {
	return o
}

func (o GatewayDetailsOutput) ToGatewayDetailsOutputWithContext(ctx context.Context) GatewayDetailsOutput {
	return o
}

func (o GatewayDetailsOutput) ToGatewayDetailsPtrOutput() GatewayDetailsPtrOutput {
	return o.ToGatewayDetailsPtrOutputWithContext(context.Background())
}

func (o GatewayDetailsOutput) ToGatewayDetailsPtrOutputWithContext(ctx context.Context) GatewayDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GatewayDetails) *GatewayDetails {
		return &v
	}).(GatewayDetailsPtrOutput)
}

func (o GatewayDetailsOutput) GatewayResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayDetails) *string { return v.GatewayResourceId }).(pulumi.StringPtrOutput)
}

type GatewayDetailsPtrOutput struct{ *pulumi.OutputState }

func (GatewayDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayDetails)(nil)).Elem()
}

func (o GatewayDetailsPtrOutput) ToGatewayDetailsPtrOutput() GatewayDetailsPtrOutput {
	return o
}

func (o GatewayDetailsPtrOutput) ToGatewayDetailsPtrOutputWithContext(ctx context.Context) GatewayDetailsPtrOutput {
	return o
}

func (o GatewayDetailsPtrOutput) Elem() GatewayDetailsOutput {
	return o.ApplyT(func(v *GatewayDetails) GatewayDetails {
		if v != nil {
			return *v
		}
		var ret GatewayDetails
		return ret
	}).(GatewayDetailsOutput)
}

func (o GatewayDetailsPtrOutput) GatewayResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayDetails) *string {
		if v == nil {
			return nil
		}
		return v.GatewayResourceId
	}).(pulumi.StringPtrOutput)
}

type GatewayDetailsResponse struct {
	DmtsClusterUri    string  `pulumi:"dmtsClusterUri"`
	GatewayObjectId   string  `pulumi:"gatewayObjectId"`
	GatewayResourceId *string `pulumi:"gatewayResourceId"`
}

type GatewayDetailsResponseOutput struct{ *pulumi.OutputState }

func (GatewayDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayDetailsResponse)(nil)).Elem()
}

func (o GatewayDetailsResponseOutput) ToGatewayDetailsResponseOutput() GatewayDetailsResponseOutput {
	return o
}

func (o GatewayDetailsResponseOutput) ToGatewayDetailsResponseOutputWithContext(ctx context.Context) GatewayDetailsResponseOutput {
	return o
}

func (o GatewayDetailsResponseOutput) DmtsClusterUri() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayDetailsResponse) string { return v.DmtsClusterUri }).(pulumi.StringOutput)
}

func (o GatewayDetailsResponseOutput) GatewayObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayDetailsResponse) string { return v.GatewayObjectId }).(pulumi.StringOutput)
}

func (o GatewayDetailsResponseOutput) GatewayResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewayDetailsResponse) *string { return v.GatewayResourceId }).(pulumi.StringPtrOutput)
}

type GatewayDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (GatewayDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayDetailsResponse)(nil)).Elem()
}

func (o GatewayDetailsResponsePtrOutput) ToGatewayDetailsResponsePtrOutput() GatewayDetailsResponsePtrOutput {
	return o
}

func (o GatewayDetailsResponsePtrOutput) ToGatewayDetailsResponsePtrOutputWithContext(ctx context.Context) GatewayDetailsResponsePtrOutput {
	return o
}

func (o GatewayDetailsResponsePtrOutput) Elem() GatewayDetailsResponseOutput {
	return o.ApplyT(func(v *GatewayDetailsResponse) GatewayDetailsResponse {
		if v != nil {
			return *v
		}
		var ret GatewayDetailsResponse
		return ret
	}).(GatewayDetailsResponseOutput)
}

func (o GatewayDetailsResponsePtrOutput) DmtsClusterUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DmtsClusterUri
	}).(pulumi.StringPtrOutput)
}

func (o GatewayDetailsResponsePtrOutput) GatewayObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.GatewayObjectId
	}).(pulumi.StringPtrOutput)
}

func (o GatewayDetailsResponsePtrOutput) GatewayResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.GatewayResourceId
	}).(pulumi.StringPtrOutput)
}

type IPv4FirewallRule struct {
	FirewallRuleName *string `pulumi:"firewallRuleName"`
	RangeEnd         *string `pulumi:"rangeEnd"`
	RangeStart       *string `pulumi:"rangeStart"`
}





type IPv4FirewallRuleInput interface {
	pulumi.Input

	ToIPv4FirewallRuleOutput() IPv4FirewallRuleOutput
	ToIPv4FirewallRuleOutputWithContext(context.Context) IPv4FirewallRuleOutput
}

type IPv4FirewallRuleArgs struct {
	FirewallRuleName pulumi.StringPtrInput `pulumi:"firewallRuleName"`
	RangeEnd         pulumi.StringPtrInput `pulumi:"rangeEnd"`
	RangeStart       pulumi.StringPtrInput `pulumi:"rangeStart"`
}

func (IPv4FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPv4FirewallRule)(nil)).Elem()
}

func (i IPv4FirewallRuleArgs) ToIPv4FirewallRuleOutput() IPv4FirewallRuleOutput {
	return i.ToIPv4FirewallRuleOutputWithContext(context.Background())
}

func (i IPv4FirewallRuleArgs) ToIPv4FirewallRuleOutputWithContext(ctx context.Context) IPv4FirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPv4FirewallRuleOutput)
}





type IPv4FirewallRuleArrayInput interface {
	pulumi.Input

	ToIPv4FirewallRuleArrayOutput() IPv4FirewallRuleArrayOutput
	ToIPv4FirewallRuleArrayOutputWithContext(context.Context) IPv4FirewallRuleArrayOutput
}

type IPv4FirewallRuleArray []IPv4FirewallRuleInput

func (IPv4FirewallRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPv4FirewallRule)(nil)).Elem()
}

func (i IPv4FirewallRuleArray) ToIPv4FirewallRuleArrayOutput() IPv4FirewallRuleArrayOutput {
	return i.ToIPv4FirewallRuleArrayOutputWithContext(context.Background())
}

func (i IPv4FirewallRuleArray) ToIPv4FirewallRuleArrayOutputWithContext(ctx context.Context) IPv4FirewallRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPv4FirewallRuleArrayOutput)
}

type IPv4FirewallRuleOutput struct{ *pulumi.OutputState }

func (IPv4FirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPv4FirewallRule)(nil)).Elem()
}

func (o IPv4FirewallRuleOutput) ToIPv4FirewallRuleOutput() IPv4FirewallRuleOutput {
	return o
}

func (o IPv4FirewallRuleOutput) ToIPv4FirewallRuleOutputWithContext(ctx context.Context) IPv4FirewallRuleOutput {
	return o
}

func (o IPv4FirewallRuleOutput) FirewallRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPv4FirewallRule) *string { return v.FirewallRuleName }).(pulumi.StringPtrOutput)
}

func (o IPv4FirewallRuleOutput) RangeEnd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPv4FirewallRule) *string { return v.RangeEnd }).(pulumi.StringPtrOutput)
}

func (o IPv4FirewallRuleOutput) RangeStart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPv4FirewallRule) *string { return v.RangeStart }).(pulumi.StringPtrOutput)
}

type IPv4FirewallRuleArrayOutput struct{ *pulumi.OutputState }

func (IPv4FirewallRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPv4FirewallRule)(nil)).Elem()
}

func (o IPv4FirewallRuleArrayOutput) ToIPv4FirewallRuleArrayOutput() IPv4FirewallRuleArrayOutput {
	return o
}

func (o IPv4FirewallRuleArrayOutput) ToIPv4FirewallRuleArrayOutputWithContext(ctx context.Context) IPv4FirewallRuleArrayOutput {
	return o
}

func (o IPv4FirewallRuleArrayOutput) Index(i pulumi.IntInput) IPv4FirewallRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPv4FirewallRule {
		return vs[0].([]IPv4FirewallRule)[vs[1].(int)]
	}).(IPv4FirewallRuleOutput)
}

type IPv4FirewallRuleResponse struct {
	FirewallRuleName *string `pulumi:"firewallRuleName"`
	RangeEnd         *string `pulumi:"rangeEnd"`
	RangeStart       *string `pulumi:"rangeStart"`
}

type IPv4FirewallRuleResponseOutput struct{ *pulumi.OutputState }

func (IPv4FirewallRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPv4FirewallRuleResponse)(nil)).Elem()
}

func (o IPv4FirewallRuleResponseOutput) ToIPv4FirewallRuleResponseOutput() IPv4FirewallRuleResponseOutput {
	return o
}

func (o IPv4FirewallRuleResponseOutput) ToIPv4FirewallRuleResponseOutputWithContext(ctx context.Context) IPv4FirewallRuleResponseOutput {
	return o
}

func (o IPv4FirewallRuleResponseOutput) FirewallRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPv4FirewallRuleResponse) *string { return v.FirewallRuleName }).(pulumi.StringPtrOutput)
}

func (o IPv4FirewallRuleResponseOutput) RangeEnd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPv4FirewallRuleResponse) *string { return v.RangeEnd }).(pulumi.StringPtrOutput)
}

func (o IPv4FirewallRuleResponseOutput) RangeStart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPv4FirewallRuleResponse) *string { return v.RangeStart }).(pulumi.StringPtrOutput)
}

type IPv4FirewallRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (IPv4FirewallRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPv4FirewallRuleResponse)(nil)).Elem()
}

func (o IPv4FirewallRuleResponseArrayOutput) ToIPv4FirewallRuleResponseArrayOutput() IPv4FirewallRuleResponseArrayOutput {
	return o
}

func (o IPv4FirewallRuleResponseArrayOutput) ToIPv4FirewallRuleResponseArrayOutputWithContext(ctx context.Context) IPv4FirewallRuleResponseArrayOutput {
	return o
}

func (o IPv4FirewallRuleResponseArrayOutput) Index(i pulumi.IntInput) IPv4FirewallRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPv4FirewallRuleResponse {
		return vs[0].([]IPv4FirewallRuleResponse)[vs[1].(int)]
	}).(IPv4FirewallRuleResponseOutput)
}

type IPv4FirewallSettings struct {
	EnablePowerBIService *bool              `pulumi:"enablePowerBIService"`
	FirewallRules        []IPv4FirewallRule `pulumi:"firewallRules"`
}





type IPv4FirewallSettingsInput interface {
	pulumi.Input

	ToIPv4FirewallSettingsOutput() IPv4FirewallSettingsOutput
	ToIPv4FirewallSettingsOutputWithContext(context.Context) IPv4FirewallSettingsOutput
}

type IPv4FirewallSettingsArgs struct {
	EnablePowerBIService pulumi.BoolPtrInput        `pulumi:"enablePowerBIService"`
	FirewallRules        IPv4FirewallRuleArrayInput `pulumi:"firewallRules"`
}

func (IPv4FirewallSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPv4FirewallSettings)(nil)).Elem()
}

func (i IPv4FirewallSettingsArgs) ToIPv4FirewallSettingsOutput() IPv4FirewallSettingsOutput {
	return i.ToIPv4FirewallSettingsOutputWithContext(context.Background())
}

func (i IPv4FirewallSettingsArgs) ToIPv4FirewallSettingsOutputWithContext(ctx context.Context) IPv4FirewallSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPv4FirewallSettingsOutput)
}

func (i IPv4FirewallSettingsArgs) ToIPv4FirewallSettingsPtrOutput() IPv4FirewallSettingsPtrOutput {
	return i.ToIPv4FirewallSettingsPtrOutputWithContext(context.Background())
}

func (i IPv4FirewallSettingsArgs) ToIPv4FirewallSettingsPtrOutputWithContext(ctx context.Context) IPv4FirewallSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPv4FirewallSettingsOutput).ToIPv4FirewallSettingsPtrOutputWithContext(ctx)
}









type IPv4FirewallSettingsPtrInput interface {
	pulumi.Input

	ToIPv4FirewallSettingsPtrOutput() IPv4FirewallSettingsPtrOutput
	ToIPv4FirewallSettingsPtrOutputWithContext(context.Context) IPv4FirewallSettingsPtrOutput
}

type ipv4FirewallSettingsPtrType IPv4FirewallSettingsArgs

func IPv4FirewallSettingsPtr(v *IPv4FirewallSettingsArgs) IPv4FirewallSettingsPtrInput {
	return (*ipv4FirewallSettingsPtrType)(v)
}

func (*ipv4FirewallSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IPv4FirewallSettings)(nil)).Elem()
}

func (i *ipv4FirewallSettingsPtrType) ToIPv4FirewallSettingsPtrOutput() IPv4FirewallSettingsPtrOutput {
	return i.ToIPv4FirewallSettingsPtrOutputWithContext(context.Background())
}

func (i *ipv4FirewallSettingsPtrType) ToIPv4FirewallSettingsPtrOutputWithContext(ctx context.Context) IPv4FirewallSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPv4FirewallSettingsPtrOutput)
}

type IPv4FirewallSettingsOutput struct{ *pulumi.OutputState }

func (IPv4FirewallSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPv4FirewallSettings)(nil)).Elem()
}

func (o IPv4FirewallSettingsOutput) ToIPv4FirewallSettingsOutput() IPv4FirewallSettingsOutput {
	return o
}

func (o IPv4FirewallSettingsOutput) ToIPv4FirewallSettingsOutputWithContext(ctx context.Context) IPv4FirewallSettingsOutput {
	return o
}

func (o IPv4FirewallSettingsOutput) ToIPv4FirewallSettingsPtrOutput() IPv4FirewallSettingsPtrOutput {
	return o.ToIPv4FirewallSettingsPtrOutputWithContext(context.Background())
}

func (o IPv4FirewallSettingsOutput) ToIPv4FirewallSettingsPtrOutputWithContext(ctx context.Context) IPv4FirewallSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IPv4FirewallSettings) *IPv4FirewallSettings {
		return &v
	}).(IPv4FirewallSettingsPtrOutput)
}

func (o IPv4FirewallSettingsOutput) EnablePowerBIService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IPv4FirewallSettings) *bool { return v.EnablePowerBIService }).(pulumi.BoolPtrOutput)
}

func (o IPv4FirewallSettingsOutput) FirewallRules() IPv4FirewallRuleArrayOutput {
	return o.ApplyT(func(v IPv4FirewallSettings) []IPv4FirewallRule { return v.FirewallRules }).(IPv4FirewallRuleArrayOutput)
}

type IPv4FirewallSettingsPtrOutput struct{ *pulumi.OutputState }

func (IPv4FirewallSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPv4FirewallSettings)(nil)).Elem()
}

func (o IPv4FirewallSettingsPtrOutput) ToIPv4FirewallSettingsPtrOutput() IPv4FirewallSettingsPtrOutput {
	return o
}

func (o IPv4FirewallSettingsPtrOutput) ToIPv4FirewallSettingsPtrOutputWithContext(ctx context.Context) IPv4FirewallSettingsPtrOutput {
	return o
}

func (o IPv4FirewallSettingsPtrOutput) Elem() IPv4FirewallSettingsOutput {
	return o.ApplyT(func(v *IPv4FirewallSettings) IPv4FirewallSettings {
		if v != nil {
			return *v
		}
		var ret IPv4FirewallSettings
		return ret
	}).(IPv4FirewallSettingsOutput)
}

func (o IPv4FirewallSettingsPtrOutput) EnablePowerBIService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IPv4FirewallSettings) *bool {
		if v == nil {
			return nil
		}
		return v.EnablePowerBIService
	}).(pulumi.BoolPtrOutput)
}

func (o IPv4FirewallSettingsPtrOutput) FirewallRules() IPv4FirewallRuleArrayOutput {
	return o.ApplyT(func(v *IPv4FirewallSettings) []IPv4FirewallRule {
		if v == nil {
			return nil
		}
		return v.FirewallRules
	}).(IPv4FirewallRuleArrayOutput)
}

type IPv4FirewallSettingsResponse struct {
	EnablePowerBIService *bool                      `pulumi:"enablePowerBIService"`
	FirewallRules        []IPv4FirewallRuleResponse `pulumi:"firewallRules"`
}

type IPv4FirewallSettingsResponseOutput struct{ *pulumi.OutputState }

func (IPv4FirewallSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPv4FirewallSettingsResponse)(nil)).Elem()
}

func (o IPv4FirewallSettingsResponseOutput) ToIPv4FirewallSettingsResponseOutput() IPv4FirewallSettingsResponseOutput {
	return o
}

func (o IPv4FirewallSettingsResponseOutput) ToIPv4FirewallSettingsResponseOutputWithContext(ctx context.Context) IPv4FirewallSettingsResponseOutput {
	return o
}

func (o IPv4FirewallSettingsResponseOutput) EnablePowerBIService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IPv4FirewallSettingsResponse) *bool { return v.EnablePowerBIService }).(pulumi.BoolPtrOutput)
}

func (o IPv4FirewallSettingsResponseOutput) FirewallRules() IPv4FirewallRuleResponseArrayOutput {
	return o.ApplyT(func(v IPv4FirewallSettingsResponse) []IPv4FirewallRuleResponse { return v.FirewallRules }).(IPv4FirewallRuleResponseArrayOutput)
}

type IPv4FirewallSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (IPv4FirewallSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPv4FirewallSettingsResponse)(nil)).Elem()
}

func (o IPv4FirewallSettingsResponsePtrOutput) ToIPv4FirewallSettingsResponsePtrOutput() IPv4FirewallSettingsResponsePtrOutput {
	return o
}

func (o IPv4FirewallSettingsResponsePtrOutput) ToIPv4FirewallSettingsResponsePtrOutputWithContext(ctx context.Context) IPv4FirewallSettingsResponsePtrOutput {
	return o
}

func (o IPv4FirewallSettingsResponsePtrOutput) Elem() IPv4FirewallSettingsResponseOutput {
	return o.ApplyT(func(v *IPv4FirewallSettingsResponse) IPv4FirewallSettingsResponse {
		if v != nil {
			return *v
		}
		var ret IPv4FirewallSettingsResponse
		return ret
	}).(IPv4FirewallSettingsResponseOutput)
}

func (o IPv4FirewallSettingsResponsePtrOutput) EnablePowerBIService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IPv4FirewallSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnablePowerBIService
	}).(pulumi.BoolPtrOutput)
}

func (o IPv4FirewallSettingsResponsePtrOutput) FirewallRules() IPv4FirewallRuleResponseArrayOutput {
	return o.ApplyT(func(v *IPv4FirewallSettingsResponse) []IPv4FirewallRuleResponse {
		if v == nil {
			return nil
		}
		return v.FirewallRules
	}).(IPv4FirewallRuleResponseArrayOutput)
}

type ResourceSku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}


func (val *ResourceSku) Defaults() *ResourceSku {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Capacity) {
		capacity_ := 1
		tmp.Capacity = &capacity_
	}
	return &tmp
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


func (val *ResourceSkuArgs) Defaults() *ResourceSkuArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Capacity) {
		tmp.Capacity = pulumi.IntPtr(1)
	}
	return &tmp
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

func (o ResourceSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ResourceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ResourceSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}


func (val *ResourceSkuResponse) Defaults() *ResourceSkuResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Capacity) {
		capacity_ := 1
		tmp.Capacity = &capacity_
	}
	return &tmp
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

func (o ResourceSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ServerAdministrators struct {
	Members []string `pulumi:"members"`
}





type ServerAdministratorsInput interface {
	pulumi.Input

	ToServerAdministratorsOutput() ServerAdministratorsOutput
	ToServerAdministratorsOutputWithContext(context.Context) ServerAdministratorsOutput
}

type ServerAdministratorsArgs struct {
	Members pulumi.StringArrayInput `pulumi:"members"`
}

func (ServerAdministratorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAdministrators)(nil)).Elem()
}

func (i ServerAdministratorsArgs) ToServerAdministratorsOutput() ServerAdministratorsOutput {
	return i.ToServerAdministratorsOutputWithContext(context.Background())
}

func (i ServerAdministratorsArgs) ToServerAdministratorsOutputWithContext(ctx context.Context) ServerAdministratorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAdministratorsOutput)
}

func (i ServerAdministratorsArgs) ToServerAdministratorsPtrOutput() ServerAdministratorsPtrOutput {
	return i.ToServerAdministratorsPtrOutputWithContext(context.Background())
}

func (i ServerAdministratorsArgs) ToServerAdministratorsPtrOutputWithContext(ctx context.Context) ServerAdministratorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAdministratorsOutput).ToServerAdministratorsPtrOutputWithContext(ctx)
}









type ServerAdministratorsPtrInput interface {
	pulumi.Input

	ToServerAdministratorsPtrOutput() ServerAdministratorsPtrOutput
	ToServerAdministratorsPtrOutputWithContext(context.Context) ServerAdministratorsPtrOutput
}

type serverAdministratorsPtrType ServerAdministratorsArgs

func ServerAdministratorsPtr(v *ServerAdministratorsArgs) ServerAdministratorsPtrInput {
	return (*serverAdministratorsPtrType)(v)
}

func (*serverAdministratorsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAdministrators)(nil)).Elem()
}

func (i *serverAdministratorsPtrType) ToServerAdministratorsPtrOutput() ServerAdministratorsPtrOutput {
	return i.ToServerAdministratorsPtrOutputWithContext(context.Background())
}

func (i *serverAdministratorsPtrType) ToServerAdministratorsPtrOutputWithContext(ctx context.Context) ServerAdministratorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerAdministratorsPtrOutput)
}

type ServerAdministratorsOutput struct{ *pulumi.OutputState }

func (ServerAdministratorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAdministrators)(nil)).Elem()
}

func (o ServerAdministratorsOutput) ToServerAdministratorsOutput() ServerAdministratorsOutput {
	return o
}

func (o ServerAdministratorsOutput) ToServerAdministratorsOutputWithContext(ctx context.Context) ServerAdministratorsOutput {
	return o
}

func (o ServerAdministratorsOutput) ToServerAdministratorsPtrOutput() ServerAdministratorsPtrOutput {
	return o.ToServerAdministratorsPtrOutputWithContext(context.Background())
}

func (o ServerAdministratorsOutput) ToServerAdministratorsPtrOutputWithContext(ctx context.Context) ServerAdministratorsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerAdministrators) *ServerAdministrators {
		return &v
	}).(ServerAdministratorsPtrOutput)
}

func (o ServerAdministratorsOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServerAdministrators) []string { return v.Members }).(pulumi.StringArrayOutput)
}

type ServerAdministratorsPtrOutput struct{ *pulumi.OutputState }

func (ServerAdministratorsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAdministrators)(nil)).Elem()
}

func (o ServerAdministratorsPtrOutput) ToServerAdministratorsPtrOutput() ServerAdministratorsPtrOutput {
	return o
}

func (o ServerAdministratorsPtrOutput) ToServerAdministratorsPtrOutputWithContext(ctx context.Context) ServerAdministratorsPtrOutput {
	return o
}

func (o ServerAdministratorsPtrOutput) Elem() ServerAdministratorsOutput {
	return o.ApplyT(func(v *ServerAdministrators) ServerAdministrators {
		if v != nil {
			return *v
		}
		var ret ServerAdministrators
		return ret
	}).(ServerAdministratorsOutput)
}

func (o ServerAdministratorsPtrOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerAdministrators) []string {
		if v == nil {
			return nil
		}
		return v.Members
	}).(pulumi.StringArrayOutput)
}

type ServerAdministratorsResponse struct {
	Members []string `pulumi:"members"`
}

type ServerAdministratorsResponseOutput struct{ *pulumi.OutputState }

func (ServerAdministratorsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerAdministratorsResponse)(nil)).Elem()
}

func (o ServerAdministratorsResponseOutput) ToServerAdministratorsResponseOutput() ServerAdministratorsResponseOutput {
	return o
}

func (o ServerAdministratorsResponseOutput) ToServerAdministratorsResponseOutputWithContext(ctx context.Context) ServerAdministratorsResponseOutput {
	return o
}

func (o ServerAdministratorsResponseOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServerAdministratorsResponse) []string { return v.Members }).(pulumi.StringArrayOutput)
}

type ServerAdministratorsResponsePtrOutput struct{ *pulumi.OutputState }

func (ServerAdministratorsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerAdministratorsResponse)(nil)).Elem()
}

func (o ServerAdministratorsResponsePtrOutput) ToServerAdministratorsResponsePtrOutput() ServerAdministratorsResponsePtrOutput {
	return o
}

func (o ServerAdministratorsResponsePtrOutput) ToServerAdministratorsResponsePtrOutputWithContext(ctx context.Context) ServerAdministratorsResponsePtrOutput {
	return o
}

func (o ServerAdministratorsResponsePtrOutput) Elem() ServerAdministratorsResponseOutput {
	return o.ApplyT(func(v *ServerAdministratorsResponse) ServerAdministratorsResponse {
		if v != nil {
			return *v
		}
		var ret ServerAdministratorsResponse
		return ret
	}).(ServerAdministratorsResponseOutput)
}

func (o ServerAdministratorsResponsePtrOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerAdministratorsResponse) []string {
		if v == nil {
			return nil
		}
		return v.Members
	}).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewayDetailsOutput{})
	pulumi.RegisterOutputType(GatewayDetailsPtrOutput{})
	pulumi.RegisterOutputType(GatewayDetailsResponseOutput{})
	pulumi.RegisterOutputType(GatewayDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(IPv4FirewallRuleOutput{})
	pulumi.RegisterOutputType(IPv4FirewallRuleArrayOutput{})
	pulumi.RegisterOutputType(IPv4FirewallRuleResponseOutput{})
	pulumi.RegisterOutputType(IPv4FirewallRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(IPv4FirewallSettingsOutput{})
	pulumi.RegisterOutputType(IPv4FirewallSettingsPtrOutput{})
	pulumi.RegisterOutputType(IPv4FirewallSettingsResponseOutput{})
	pulumi.RegisterOutputType(IPv4FirewallSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponseOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsPtrOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsResponseOutput{})
	pulumi.RegisterOutputType(ServerAdministratorsResponsePtrOutput{})
}
