


package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DiagnosticSetting struct {
	pulumi.CustomResourceState

	EventHubAuthorizationRuleId pulumi.StringPtrOutput         `pulumi:"eventHubAuthorizationRuleId"`
	EventHubName                pulumi.StringPtrOutput         `pulumi:"eventHubName"`
	Logs                        LogSettingsResponseArrayOutput `pulumi:"logs"`
	Name                        pulumi.StringOutput            `pulumi:"name"`
	ServiceBusRuleId            pulumi.StringPtrOutput         `pulumi:"serviceBusRuleId"`
	StorageAccountId            pulumi.StringPtrOutput         `pulumi:"storageAccountId"`
	Type                        pulumi.StringOutput            `pulumi:"type"`
	WorkspaceId                 pulumi.StringPtrOutput         `pulumi:"workspaceId"`
}


func NewDiagnosticSetting(ctx *pulumi.Context,
	name string, args *DiagnosticSettingArgs, opts ...pulumi.ResourceOption) (*DiagnosticSetting, error) {
	if args == nil {
		args = &DiagnosticSettingArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:aadiam:DiagnosticSetting"),
		},
		{
			Type: pulumi.String("azure-native:aadiam/v20170401preview:DiagnosticSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource DiagnosticSetting
	err := ctx.RegisterResource("azure-native:aadiam/v20170401:DiagnosticSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDiagnosticSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiagnosticSettingState, opts ...pulumi.ResourceOption) (*DiagnosticSetting, error) {
	var resource DiagnosticSetting
	err := ctx.ReadResource("azure-native:aadiam/v20170401:DiagnosticSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type diagnosticSettingState struct {
}

type DiagnosticSettingState struct {
}

func (DiagnosticSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticSettingState)(nil)).Elem()
}

type diagnosticSettingArgs struct {
	EventHubAuthorizationRuleId *string       `pulumi:"eventHubAuthorizationRuleId"`
	EventHubName                *string       `pulumi:"eventHubName"`
	Logs                        []LogSettings `pulumi:"logs"`
	Name                        *string       `pulumi:"name"`
	ServiceBusRuleId            *string       `pulumi:"serviceBusRuleId"`
	StorageAccountId            *string       `pulumi:"storageAccountId"`
	WorkspaceId                 *string       `pulumi:"workspaceId"`
}


type DiagnosticSettingArgs struct {
	EventHubAuthorizationRuleId pulumi.StringPtrInput
	EventHubName                pulumi.StringPtrInput
	Logs                        LogSettingsArrayInput
	Name                        pulumi.StringPtrInput
	ServiceBusRuleId            pulumi.StringPtrInput
	StorageAccountId            pulumi.StringPtrInput
	WorkspaceId                 pulumi.StringPtrInput
}

func (DiagnosticSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticSettingArgs)(nil)).Elem()
}

type DiagnosticSettingInput interface {
	pulumi.Input

	ToDiagnosticSettingOutput() DiagnosticSettingOutput
	ToDiagnosticSettingOutputWithContext(ctx context.Context) DiagnosticSettingOutput
}

func (*DiagnosticSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticSetting)(nil)).Elem()
}

func (i *DiagnosticSetting) ToDiagnosticSettingOutput() DiagnosticSettingOutput {
	return i.ToDiagnosticSettingOutputWithContext(context.Background())
}

func (i *DiagnosticSetting) ToDiagnosticSettingOutputWithContext(ctx context.Context) DiagnosticSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticSettingOutput)
}

type DiagnosticSettingOutput struct{ *pulumi.OutputState }

func (DiagnosticSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticSetting)(nil)).Elem()
}

func (o DiagnosticSettingOutput) ToDiagnosticSettingOutput() DiagnosticSettingOutput {
	return o
}

func (o DiagnosticSettingOutput) ToDiagnosticSettingOutputWithContext(ctx context.Context) DiagnosticSettingOutput {
	return o
}

func (o DiagnosticSettingOutput) EventHubAuthorizationRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticSetting) pulumi.StringPtrOutput { return v.EventHubAuthorizationRuleId }).(pulumi.StringPtrOutput)
}

func (o DiagnosticSettingOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticSetting) pulumi.StringPtrOutput { return v.EventHubName }).(pulumi.StringPtrOutput)
}

func (o DiagnosticSettingOutput) Logs() LogSettingsResponseArrayOutput {
	return o.ApplyT(func(v *DiagnosticSetting) LogSettingsResponseArrayOutput { return v.Logs }).(LogSettingsResponseArrayOutput)
}

func (o DiagnosticSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DiagnosticSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DiagnosticSettingOutput) ServiceBusRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticSetting) pulumi.StringPtrOutput { return v.ServiceBusRuleId }).(pulumi.StringPtrOutput)
}

func (o DiagnosticSettingOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticSetting) pulumi.StringPtrOutput { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

func (o DiagnosticSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DiagnosticSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o DiagnosticSettingOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticSetting) pulumi.StringPtrOutput { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DiagnosticSettingOutput{})
}
