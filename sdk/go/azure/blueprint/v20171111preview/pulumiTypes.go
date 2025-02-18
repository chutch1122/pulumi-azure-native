


package v20171111preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssignmentLockSettings struct {
	Mode *string `pulumi:"mode"`
}





type AssignmentLockSettingsInput interface {
	pulumi.Input

	ToAssignmentLockSettingsOutput() AssignmentLockSettingsOutput
	ToAssignmentLockSettingsOutputWithContext(context.Context) AssignmentLockSettingsOutput
}

type AssignmentLockSettingsArgs struct {
	Mode pulumi.StringPtrInput `pulumi:"mode"`
}

func (AssignmentLockSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentLockSettings)(nil)).Elem()
}

func (i AssignmentLockSettingsArgs) ToAssignmentLockSettingsOutput() AssignmentLockSettingsOutput {
	return i.ToAssignmentLockSettingsOutputWithContext(context.Background())
}

func (i AssignmentLockSettingsArgs) ToAssignmentLockSettingsOutputWithContext(ctx context.Context) AssignmentLockSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentLockSettingsOutput)
}

func (i AssignmentLockSettingsArgs) ToAssignmentLockSettingsPtrOutput() AssignmentLockSettingsPtrOutput {
	return i.ToAssignmentLockSettingsPtrOutputWithContext(context.Background())
}

func (i AssignmentLockSettingsArgs) ToAssignmentLockSettingsPtrOutputWithContext(ctx context.Context) AssignmentLockSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentLockSettingsOutput).ToAssignmentLockSettingsPtrOutputWithContext(ctx)
}









type AssignmentLockSettingsPtrInput interface {
	pulumi.Input

	ToAssignmentLockSettingsPtrOutput() AssignmentLockSettingsPtrOutput
	ToAssignmentLockSettingsPtrOutputWithContext(context.Context) AssignmentLockSettingsPtrOutput
}

type assignmentLockSettingsPtrType AssignmentLockSettingsArgs

func AssignmentLockSettingsPtr(v *AssignmentLockSettingsArgs) AssignmentLockSettingsPtrInput {
	return (*assignmentLockSettingsPtrType)(v)
}

func (*assignmentLockSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentLockSettings)(nil)).Elem()
}

func (i *assignmentLockSettingsPtrType) ToAssignmentLockSettingsPtrOutput() AssignmentLockSettingsPtrOutput {
	return i.ToAssignmentLockSettingsPtrOutputWithContext(context.Background())
}

func (i *assignmentLockSettingsPtrType) ToAssignmentLockSettingsPtrOutputWithContext(ctx context.Context) AssignmentLockSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentLockSettingsPtrOutput)
}

type AssignmentLockSettingsOutput struct{ *pulumi.OutputState }

func (AssignmentLockSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentLockSettings)(nil)).Elem()
}

func (o AssignmentLockSettingsOutput) ToAssignmentLockSettingsOutput() AssignmentLockSettingsOutput {
	return o
}

func (o AssignmentLockSettingsOutput) ToAssignmentLockSettingsOutputWithContext(ctx context.Context) AssignmentLockSettingsOutput {
	return o
}

func (o AssignmentLockSettingsOutput) ToAssignmentLockSettingsPtrOutput() AssignmentLockSettingsPtrOutput {
	return o.ToAssignmentLockSettingsPtrOutputWithContext(context.Background())
}

func (o AssignmentLockSettingsOutput) ToAssignmentLockSettingsPtrOutputWithContext(ctx context.Context) AssignmentLockSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssignmentLockSettings) *AssignmentLockSettings {
		return &v
	}).(AssignmentLockSettingsPtrOutput)
}

func (o AssignmentLockSettingsOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssignmentLockSettings) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

type AssignmentLockSettingsPtrOutput struct{ *pulumi.OutputState }

func (AssignmentLockSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentLockSettings)(nil)).Elem()
}

func (o AssignmentLockSettingsPtrOutput) ToAssignmentLockSettingsPtrOutput() AssignmentLockSettingsPtrOutput {
	return o
}

func (o AssignmentLockSettingsPtrOutput) ToAssignmentLockSettingsPtrOutputWithContext(ctx context.Context) AssignmentLockSettingsPtrOutput {
	return o
}

func (o AssignmentLockSettingsPtrOutput) Elem() AssignmentLockSettingsOutput {
	return o.ApplyT(func(v *AssignmentLockSettings) AssignmentLockSettings {
		if v != nil {
			return *v
		}
		var ret AssignmentLockSettings
		return ret
	}).(AssignmentLockSettingsOutput)
}

func (o AssignmentLockSettingsPtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignmentLockSettings) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

type AssignmentLockSettingsResponse struct {
	Mode *string `pulumi:"mode"`
}

type AssignmentLockSettingsResponseOutput struct{ *pulumi.OutputState }

func (AssignmentLockSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentLockSettingsResponse)(nil)).Elem()
}

func (o AssignmentLockSettingsResponseOutput) ToAssignmentLockSettingsResponseOutput() AssignmentLockSettingsResponseOutput {
	return o
}

func (o AssignmentLockSettingsResponseOutput) ToAssignmentLockSettingsResponseOutputWithContext(ctx context.Context) AssignmentLockSettingsResponseOutput {
	return o
}

func (o AssignmentLockSettingsResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssignmentLockSettingsResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

type AssignmentLockSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AssignmentLockSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentLockSettingsResponse)(nil)).Elem()
}

func (o AssignmentLockSettingsResponsePtrOutput) ToAssignmentLockSettingsResponsePtrOutput() AssignmentLockSettingsResponsePtrOutput {
	return o
}

func (o AssignmentLockSettingsResponsePtrOutput) ToAssignmentLockSettingsResponsePtrOutputWithContext(ctx context.Context) AssignmentLockSettingsResponsePtrOutput {
	return o
}

func (o AssignmentLockSettingsResponsePtrOutput) Elem() AssignmentLockSettingsResponseOutput {
	return o.ApplyT(func(v *AssignmentLockSettingsResponse) AssignmentLockSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AssignmentLockSettingsResponse
		return ret
	}).(AssignmentLockSettingsResponseOutput)
}

func (o AssignmentLockSettingsResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignmentLockSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

type AssignmentStatusResponse struct {
	LastModified string `pulumi:"lastModified"`
	TimeCreated  string `pulumi:"timeCreated"`
}

type AssignmentStatusResponseOutput struct{ *pulumi.OutputState }

func (AssignmentStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentStatusResponse)(nil)).Elem()
}

func (o AssignmentStatusResponseOutput) ToAssignmentStatusResponseOutput() AssignmentStatusResponseOutput {
	return o
}

func (o AssignmentStatusResponseOutput) ToAssignmentStatusResponseOutputWithContext(ctx context.Context) AssignmentStatusResponseOutput {
	return o
}

func (o AssignmentStatusResponseOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentStatusResponse) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o AssignmentStatusResponseOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentStatusResponse) string { return v.TimeCreated }).(pulumi.StringOutput)
}

type BlueprintStatusResponse struct {
	LastModified string `pulumi:"lastModified"`
	TimeCreated  string `pulumi:"timeCreated"`
}

type BlueprintStatusResponseOutput struct{ *pulumi.OutputState }

func (BlueprintStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlueprintStatusResponse)(nil)).Elem()
}

func (o BlueprintStatusResponseOutput) ToBlueprintStatusResponseOutput() BlueprintStatusResponseOutput {
	return o
}

func (o BlueprintStatusResponseOutput) ToBlueprintStatusResponseOutputWithContext(ctx context.Context) BlueprintStatusResponseOutput {
	return o
}

func (o BlueprintStatusResponseOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v BlueprintStatusResponse) string { return v.LastModified }).(pulumi.StringOutput)
}

func (o BlueprintStatusResponseOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v BlueprintStatusResponse) string { return v.TimeCreated }).(pulumi.StringOutput)
}

type ManagedServiceIdentity struct {
	PrincipalId *string `pulumi:"principalId"`
	TenantId    *string `pulumi:"tenantId"`
	Type        string  `pulumi:"type"`
}





type ManagedServiceIdentityInput interface {
	pulumi.Input

	ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput
	ToManagedServiceIdentityOutputWithContext(context.Context) ManagedServiceIdentityOutput
}

type ManagedServiceIdentityArgs struct {
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
	TenantId    pulumi.StringPtrInput `pulumi:"tenantId"`
	Type        pulumi.StringInput    `pulumi:"type"`
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

func (o ManagedServiceIdentityOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type ManagedServiceIdentityResponse struct {
	PrincipalId *string `pulumi:"principalId"`
	TenantId    *string `pulumi:"tenantId"`
	Type        string  `pulumi:"type"`
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

func (o ManagedServiceIdentityResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ParameterDefinition struct {
	AllowedValues []interface{} `pulumi:"allowedValues"`
	DefaultValue  interface{}   `pulumi:"defaultValue"`
	Description   *string       `pulumi:"description"`
	DisplayName   *string       `pulumi:"displayName"`
	StrongType    *string       `pulumi:"strongType"`
	Type          string        `pulumi:"type"`
}





type ParameterDefinitionInput interface {
	pulumi.Input

	ToParameterDefinitionOutput() ParameterDefinitionOutput
	ToParameterDefinitionOutputWithContext(context.Context) ParameterDefinitionOutput
}

type ParameterDefinitionArgs struct {
	AllowedValues pulumi.ArrayInput     `pulumi:"allowedValues"`
	DefaultValue  pulumi.Input          `pulumi:"defaultValue"`
	Description   pulumi.StringPtrInput `pulumi:"description"`
	DisplayName   pulumi.StringPtrInput `pulumi:"displayName"`
	StrongType    pulumi.StringPtrInput `pulumi:"strongType"`
	Type          pulumi.StringInput    `pulumi:"type"`
}

func (ParameterDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinition)(nil)).Elem()
}

func (i ParameterDefinitionArgs) ToParameterDefinitionOutput() ParameterDefinitionOutput {
	return i.ToParameterDefinitionOutputWithContext(context.Background())
}

func (i ParameterDefinitionArgs) ToParameterDefinitionOutputWithContext(ctx context.Context) ParameterDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionOutput)
}





type ParameterDefinitionMapInput interface {
	pulumi.Input

	ToParameterDefinitionMapOutput() ParameterDefinitionMapOutput
	ToParameterDefinitionMapOutputWithContext(context.Context) ParameterDefinitionMapOutput
}

type ParameterDefinitionMap map[string]ParameterDefinitionInput

func (ParameterDefinitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterDefinition)(nil)).Elem()
}

func (i ParameterDefinitionMap) ToParameterDefinitionMapOutput() ParameterDefinitionMapOutput {
	return i.ToParameterDefinitionMapOutputWithContext(context.Background())
}

func (i ParameterDefinitionMap) ToParameterDefinitionMapOutputWithContext(ctx context.Context) ParameterDefinitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterDefinitionMapOutput)
}

type ParameterDefinitionOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinition)(nil)).Elem()
}

func (o ParameterDefinitionOutput) ToParameterDefinitionOutput() ParameterDefinitionOutput {
	return o
}

func (o ParameterDefinitionOutput) ToParameterDefinitionOutputWithContext(ctx context.Context) ParameterDefinitionOutput {
	return o
}

func (o ParameterDefinitionOutput) AllowedValues() pulumi.ArrayOutput {
	return o.ApplyT(func(v ParameterDefinition) []interface{} { return v.AllowedValues }).(pulumi.ArrayOutput)
}

func (o ParameterDefinitionOutput) DefaultValue() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterDefinition) interface{} { return v.DefaultValue }).(pulumi.AnyOutput)
}

func (o ParameterDefinitionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinition) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionOutput) StrongType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinition) *string { return v.StrongType }).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterDefinition) string { return v.Type }).(pulumi.StringOutput)
}

type ParameterDefinitionMapOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterDefinition)(nil)).Elem()
}

func (o ParameterDefinitionMapOutput) ToParameterDefinitionMapOutput() ParameterDefinitionMapOutput {
	return o
}

func (o ParameterDefinitionMapOutput) ToParameterDefinitionMapOutputWithContext(ctx context.Context) ParameterDefinitionMapOutput {
	return o
}

func (o ParameterDefinitionMapOutput) MapIndex(k pulumi.StringInput) ParameterDefinitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterDefinition {
		return vs[0].(map[string]ParameterDefinition)[vs[1].(string)]
	}).(ParameterDefinitionOutput)
}

type ParameterDefinitionResponse struct {
	AllowedValues []interface{} `pulumi:"allowedValues"`
	DefaultValue  interface{}   `pulumi:"defaultValue"`
	Description   *string       `pulumi:"description"`
	DisplayName   *string       `pulumi:"displayName"`
	StrongType    *string       `pulumi:"strongType"`
	Type          string        `pulumi:"type"`
}

type ParameterDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterDefinitionResponse)(nil)).Elem()
}

func (o ParameterDefinitionResponseOutput) ToParameterDefinitionResponseOutput() ParameterDefinitionResponseOutput {
	return o
}

func (o ParameterDefinitionResponseOutput) ToParameterDefinitionResponseOutputWithContext(ctx context.Context) ParameterDefinitionResponseOutput {
	return o
}

func (o ParameterDefinitionResponseOutput) AllowedValues() pulumi.ArrayOutput {
	return o.ApplyT(func(v ParameterDefinitionResponse) []interface{} { return v.AllowedValues }).(pulumi.ArrayOutput)
}

func (o ParameterDefinitionResponseOutput) DefaultValue() pulumi.AnyOutput {
	return o.ApplyT(func(v ParameterDefinitionResponse) interface{} { return v.DefaultValue }).(pulumi.AnyOutput)
}

func (o ParameterDefinitionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionResponseOutput) StrongType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterDefinitionResponse) *string { return v.StrongType }).(pulumi.StringPtrOutput)
}

func (o ParameterDefinitionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterDefinitionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ParameterDefinitionResponseMapOutput struct{ *pulumi.OutputState }

func (ParameterDefinitionResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterDefinitionResponse)(nil)).Elem()
}

func (o ParameterDefinitionResponseMapOutput) ToParameterDefinitionResponseMapOutput() ParameterDefinitionResponseMapOutput {
	return o
}

func (o ParameterDefinitionResponseMapOutput) ToParameterDefinitionResponseMapOutputWithContext(ctx context.Context) ParameterDefinitionResponseMapOutput {
	return o
}

func (o ParameterDefinitionResponseMapOutput) MapIndex(k pulumi.StringInput) ParameterDefinitionResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterDefinitionResponse {
		return vs[0].(map[string]ParameterDefinitionResponse)[vs[1].(string)]
	}).(ParameterDefinitionResponseOutput)
}

type ParameterValueBase struct {
	Description *string `pulumi:"description"`
}





type ParameterValueBaseInput interface {
	pulumi.Input

	ToParameterValueBaseOutput() ParameterValueBaseOutput
	ToParameterValueBaseOutputWithContext(context.Context) ParameterValueBaseOutput
}

type ParameterValueBaseArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
}

func (ParameterValueBaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterValueBase)(nil)).Elem()
}

func (i ParameterValueBaseArgs) ToParameterValueBaseOutput() ParameterValueBaseOutput {
	return i.ToParameterValueBaseOutputWithContext(context.Background())
}

func (i ParameterValueBaseArgs) ToParameterValueBaseOutputWithContext(ctx context.Context) ParameterValueBaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterValueBaseOutput)
}





type ParameterValueBaseMapInput interface {
	pulumi.Input

	ToParameterValueBaseMapOutput() ParameterValueBaseMapOutput
	ToParameterValueBaseMapOutputWithContext(context.Context) ParameterValueBaseMapOutput
}

type ParameterValueBaseMap map[string]ParameterValueBaseInput

func (ParameterValueBaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterValueBase)(nil)).Elem()
}

func (i ParameterValueBaseMap) ToParameterValueBaseMapOutput() ParameterValueBaseMapOutput {
	return i.ToParameterValueBaseMapOutputWithContext(context.Background())
}

func (i ParameterValueBaseMap) ToParameterValueBaseMapOutputWithContext(ctx context.Context) ParameterValueBaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterValueBaseMapOutput)
}

type ParameterValueBaseOutput struct{ *pulumi.OutputState }

func (ParameterValueBaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterValueBase)(nil)).Elem()
}

func (o ParameterValueBaseOutput) ToParameterValueBaseOutput() ParameterValueBaseOutput {
	return o
}

func (o ParameterValueBaseOutput) ToParameterValueBaseOutputWithContext(ctx context.Context) ParameterValueBaseOutput {
	return o
}

func (o ParameterValueBaseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterValueBase) *string { return v.Description }).(pulumi.StringPtrOutput)
}

type ParameterValueBaseMapOutput struct{ *pulumi.OutputState }

func (ParameterValueBaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterValueBase)(nil)).Elem()
}

func (o ParameterValueBaseMapOutput) ToParameterValueBaseMapOutput() ParameterValueBaseMapOutput {
	return o
}

func (o ParameterValueBaseMapOutput) ToParameterValueBaseMapOutputWithContext(ctx context.Context) ParameterValueBaseMapOutput {
	return o
}

func (o ParameterValueBaseMapOutput) MapIndex(k pulumi.StringInput) ParameterValueBaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterValueBase {
		return vs[0].(map[string]ParameterValueBase)[vs[1].(string)]
	}).(ParameterValueBaseOutput)
}

type ParameterValueBaseResponse struct {
	Description *string `pulumi:"description"`
}

type ParameterValueBaseResponseOutput struct{ *pulumi.OutputState }

func (ParameterValueBaseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterValueBaseResponse)(nil)).Elem()
}

func (o ParameterValueBaseResponseOutput) ToParameterValueBaseResponseOutput() ParameterValueBaseResponseOutput {
	return o
}

func (o ParameterValueBaseResponseOutput) ToParameterValueBaseResponseOutputWithContext(ctx context.Context) ParameterValueBaseResponseOutput {
	return o
}

func (o ParameterValueBaseResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterValueBaseResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

type ParameterValueBaseResponseMapOutput struct{ *pulumi.OutputState }

func (ParameterValueBaseResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ParameterValueBaseResponse)(nil)).Elem()
}

func (o ParameterValueBaseResponseMapOutput) ToParameterValueBaseResponseMapOutput() ParameterValueBaseResponseMapOutput {
	return o
}

func (o ParameterValueBaseResponseMapOutput) ToParameterValueBaseResponseMapOutputWithContext(ctx context.Context) ParameterValueBaseResponseMapOutput {
	return o
}

func (o ParameterValueBaseResponseMapOutput) MapIndex(k pulumi.StringInput) ParameterValueBaseResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ParameterValueBaseResponse {
		return vs[0].(map[string]ParameterValueBaseResponse)[vs[1].(string)]
	}).(ParameterValueBaseResponseOutput)
}

type ResourceGroupDefinition struct {
	DependsOn   []string `pulumi:"dependsOn"`
	Description *string  `pulumi:"description"`
	DisplayName *string  `pulumi:"displayName"`
	Location    *string  `pulumi:"location"`
	Name        *string  `pulumi:"name"`
	StrongType  *string  `pulumi:"strongType"`
}





type ResourceGroupDefinitionInput interface {
	pulumi.Input

	ToResourceGroupDefinitionOutput() ResourceGroupDefinitionOutput
	ToResourceGroupDefinitionOutputWithContext(context.Context) ResourceGroupDefinitionOutput
}

type ResourceGroupDefinitionArgs struct {
	DependsOn   pulumi.StringArrayInput `pulumi:"dependsOn"`
	Description pulumi.StringPtrInput   `pulumi:"description"`
	DisplayName pulumi.StringPtrInput   `pulumi:"displayName"`
	Location    pulumi.StringPtrInput   `pulumi:"location"`
	Name        pulumi.StringPtrInput   `pulumi:"name"`
	StrongType  pulumi.StringPtrInput   `pulumi:"strongType"`
}

func (ResourceGroupDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupDefinition)(nil)).Elem()
}

func (i ResourceGroupDefinitionArgs) ToResourceGroupDefinitionOutput() ResourceGroupDefinitionOutput {
	return i.ToResourceGroupDefinitionOutputWithContext(context.Background())
}

func (i ResourceGroupDefinitionArgs) ToResourceGroupDefinitionOutputWithContext(ctx context.Context) ResourceGroupDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupDefinitionOutput)
}





type ResourceGroupDefinitionMapInput interface {
	pulumi.Input

	ToResourceGroupDefinitionMapOutput() ResourceGroupDefinitionMapOutput
	ToResourceGroupDefinitionMapOutputWithContext(context.Context) ResourceGroupDefinitionMapOutput
}

type ResourceGroupDefinitionMap map[string]ResourceGroupDefinitionInput

func (ResourceGroupDefinitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceGroupDefinition)(nil)).Elem()
}

func (i ResourceGroupDefinitionMap) ToResourceGroupDefinitionMapOutput() ResourceGroupDefinitionMapOutput {
	return i.ToResourceGroupDefinitionMapOutputWithContext(context.Background())
}

func (i ResourceGroupDefinitionMap) ToResourceGroupDefinitionMapOutputWithContext(ctx context.Context) ResourceGroupDefinitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupDefinitionMapOutput)
}

type ResourceGroupDefinitionOutput struct{ *pulumi.OutputState }

func (ResourceGroupDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupDefinition)(nil)).Elem()
}

func (o ResourceGroupDefinitionOutput) ToResourceGroupDefinitionOutput() ResourceGroupDefinitionOutput {
	return o
}

func (o ResourceGroupDefinitionOutput) ToResourceGroupDefinitionOutputWithContext(ctx context.Context) ResourceGroupDefinitionOutput {
	return o
}

func (o ResourceGroupDefinitionOutput) DependsOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceGroupDefinition) []string { return v.DependsOn }).(pulumi.StringArrayOutput)
}

func (o ResourceGroupDefinitionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinition) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupDefinitionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinition) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupDefinitionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinition) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupDefinitionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinition) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupDefinitionOutput) StrongType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinition) *string { return v.StrongType }).(pulumi.StringPtrOutput)
}

type ResourceGroupDefinitionMapOutput struct{ *pulumi.OutputState }

func (ResourceGroupDefinitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceGroupDefinition)(nil)).Elem()
}

func (o ResourceGroupDefinitionMapOutput) ToResourceGroupDefinitionMapOutput() ResourceGroupDefinitionMapOutput {
	return o
}

func (o ResourceGroupDefinitionMapOutput) ToResourceGroupDefinitionMapOutputWithContext(ctx context.Context) ResourceGroupDefinitionMapOutput {
	return o
}

func (o ResourceGroupDefinitionMapOutput) MapIndex(k pulumi.StringInput) ResourceGroupDefinitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourceGroupDefinition {
		return vs[0].(map[string]ResourceGroupDefinition)[vs[1].(string)]
	}).(ResourceGroupDefinitionOutput)
}

type ResourceGroupDefinitionResponse struct {
	DependsOn   []string `pulumi:"dependsOn"`
	Description *string  `pulumi:"description"`
	DisplayName *string  `pulumi:"displayName"`
	Location    *string  `pulumi:"location"`
	Name        *string  `pulumi:"name"`
	StrongType  *string  `pulumi:"strongType"`
}

type ResourceGroupDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ResourceGroupDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupDefinitionResponse)(nil)).Elem()
}

func (o ResourceGroupDefinitionResponseOutput) ToResourceGroupDefinitionResponseOutput() ResourceGroupDefinitionResponseOutput {
	return o
}

func (o ResourceGroupDefinitionResponseOutput) ToResourceGroupDefinitionResponseOutputWithContext(ctx context.Context) ResourceGroupDefinitionResponseOutput {
	return o
}

func (o ResourceGroupDefinitionResponseOutput) DependsOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResourceGroupDefinitionResponse) []string { return v.DependsOn }).(pulumi.StringArrayOutput)
}

func (o ResourceGroupDefinitionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinitionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupDefinitionResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinitionResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupDefinitionResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinitionResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupDefinitionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinitionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupDefinitionResponseOutput) StrongType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupDefinitionResponse) *string { return v.StrongType }).(pulumi.StringPtrOutput)
}

type ResourceGroupDefinitionResponseMapOutput struct{ *pulumi.OutputState }

func (ResourceGroupDefinitionResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceGroupDefinitionResponse)(nil)).Elem()
}

func (o ResourceGroupDefinitionResponseMapOutput) ToResourceGroupDefinitionResponseMapOutput() ResourceGroupDefinitionResponseMapOutput {
	return o
}

func (o ResourceGroupDefinitionResponseMapOutput) ToResourceGroupDefinitionResponseMapOutputWithContext(ctx context.Context) ResourceGroupDefinitionResponseMapOutput {
	return o
}

func (o ResourceGroupDefinitionResponseMapOutput) MapIndex(k pulumi.StringInput) ResourceGroupDefinitionResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourceGroupDefinitionResponse {
		return vs[0].(map[string]ResourceGroupDefinitionResponse)[vs[1].(string)]
	}).(ResourceGroupDefinitionResponseOutput)
}

type ResourceGroupValue struct {
	Location *string `pulumi:"location"`
	Name     *string `pulumi:"name"`
}





type ResourceGroupValueInput interface {
	pulumi.Input

	ToResourceGroupValueOutput() ResourceGroupValueOutput
	ToResourceGroupValueOutputWithContext(context.Context) ResourceGroupValueOutput
}

type ResourceGroupValueArgs struct {
	Location pulumi.StringPtrInput `pulumi:"location"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
}

func (ResourceGroupValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupValue)(nil)).Elem()
}

func (i ResourceGroupValueArgs) ToResourceGroupValueOutput() ResourceGroupValueOutput {
	return i.ToResourceGroupValueOutputWithContext(context.Background())
}

func (i ResourceGroupValueArgs) ToResourceGroupValueOutputWithContext(ctx context.Context) ResourceGroupValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupValueOutput)
}





type ResourceGroupValueMapInput interface {
	pulumi.Input

	ToResourceGroupValueMapOutput() ResourceGroupValueMapOutput
	ToResourceGroupValueMapOutputWithContext(context.Context) ResourceGroupValueMapOutput
}

type ResourceGroupValueMap map[string]ResourceGroupValueInput

func (ResourceGroupValueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceGroupValue)(nil)).Elem()
}

func (i ResourceGroupValueMap) ToResourceGroupValueMapOutput() ResourceGroupValueMapOutput {
	return i.ToResourceGroupValueMapOutputWithContext(context.Background())
}

func (i ResourceGroupValueMap) ToResourceGroupValueMapOutputWithContext(ctx context.Context) ResourceGroupValueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupValueMapOutput)
}

type ResourceGroupValueOutput struct{ *pulumi.OutputState }

func (ResourceGroupValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupValue)(nil)).Elem()
}

func (o ResourceGroupValueOutput) ToResourceGroupValueOutput() ResourceGroupValueOutput {
	return o
}

func (o ResourceGroupValueOutput) ToResourceGroupValueOutputWithContext(ctx context.Context) ResourceGroupValueOutput {
	return o
}

func (o ResourceGroupValueOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupValue) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupValueOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupValue) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ResourceGroupValueMapOutput struct{ *pulumi.OutputState }

func (ResourceGroupValueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceGroupValue)(nil)).Elem()
}

func (o ResourceGroupValueMapOutput) ToResourceGroupValueMapOutput() ResourceGroupValueMapOutput {
	return o
}

func (o ResourceGroupValueMapOutput) ToResourceGroupValueMapOutputWithContext(ctx context.Context) ResourceGroupValueMapOutput {
	return o
}

func (o ResourceGroupValueMapOutput) MapIndex(k pulumi.StringInput) ResourceGroupValueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourceGroupValue {
		return vs[0].(map[string]ResourceGroupValue)[vs[1].(string)]
	}).(ResourceGroupValueOutput)
}

type ResourceGroupValueResponse struct {
	Location *string `pulumi:"location"`
	Name     *string `pulumi:"name"`
}

type ResourceGroupValueResponseOutput struct{ *pulumi.OutputState }

func (ResourceGroupValueResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupValueResponse)(nil)).Elem()
}

func (o ResourceGroupValueResponseOutput) ToResourceGroupValueResponseOutput() ResourceGroupValueResponseOutput {
	return o
}

func (o ResourceGroupValueResponseOutput) ToResourceGroupValueResponseOutputWithContext(ctx context.Context) ResourceGroupValueResponseOutput {
	return o
}

func (o ResourceGroupValueResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupValueResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ResourceGroupValueResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGroupValueResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ResourceGroupValueResponseMapOutput struct{ *pulumi.OutputState }

func (ResourceGroupValueResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceGroupValueResponse)(nil)).Elem()
}

func (o ResourceGroupValueResponseMapOutput) ToResourceGroupValueResponseMapOutput() ResourceGroupValueResponseMapOutput {
	return o
}

func (o ResourceGroupValueResponseMapOutput) ToResourceGroupValueResponseMapOutputWithContext(ctx context.Context) ResourceGroupValueResponseMapOutput {
	return o
}

func (o ResourceGroupValueResponseMapOutput) MapIndex(k pulumi.StringInput) ResourceGroupValueResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourceGroupValueResponse {
		return vs[0].(map[string]ResourceGroupValueResponse)[vs[1].(string)]
	}).(ResourceGroupValueResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AssignmentLockSettingsOutput{})
	pulumi.RegisterOutputType(AssignmentLockSettingsPtrOutput{})
	pulumi.RegisterOutputType(AssignmentLockSettingsResponseOutput{})
	pulumi.RegisterOutputType(AssignmentLockSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AssignmentStatusResponseOutput{})
	pulumi.RegisterOutputType(BlueprintStatusResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionMapOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ParameterDefinitionResponseMapOutput{})
	pulumi.RegisterOutputType(ParameterValueBaseOutput{})
	pulumi.RegisterOutputType(ParameterValueBaseMapOutput{})
	pulumi.RegisterOutputType(ParameterValueBaseResponseOutput{})
	pulumi.RegisterOutputType(ParameterValueBaseResponseMapOutput{})
	pulumi.RegisterOutputType(ResourceGroupDefinitionOutput{})
	pulumi.RegisterOutputType(ResourceGroupDefinitionMapOutput{})
	pulumi.RegisterOutputType(ResourceGroupDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ResourceGroupDefinitionResponseMapOutput{})
	pulumi.RegisterOutputType(ResourceGroupValueOutput{})
	pulumi.RegisterOutputType(ResourceGroupValueMapOutput{})
	pulumi.RegisterOutputType(ResourceGroupValueResponseOutput{})
	pulumi.RegisterOutputType(ResourceGroupValueResponseMapOutput{})
}
