


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AddRemoveReplicaScalingMechanism struct {
	Kind           string `pulumi:"kind"`
	MaxCount       int    `pulumi:"maxCount"`
	MinCount       int    `pulumi:"minCount"`
	ScaleIncrement int    `pulumi:"scaleIncrement"`
}





type AddRemoveReplicaScalingMechanismInput interface {
	pulumi.Input

	ToAddRemoveReplicaScalingMechanismOutput() AddRemoveReplicaScalingMechanismOutput
	ToAddRemoveReplicaScalingMechanismOutputWithContext(context.Context) AddRemoveReplicaScalingMechanismOutput
}

type AddRemoveReplicaScalingMechanismArgs struct {
	Kind           pulumi.StringInput `pulumi:"kind"`
	MaxCount       pulumi.IntInput    `pulumi:"maxCount"`
	MinCount       pulumi.IntInput    `pulumi:"minCount"`
	ScaleIncrement pulumi.IntInput    `pulumi:"scaleIncrement"`
}

func (AddRemoveReplicaScalingMechanismArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddRemoveReplicaScalingMechanism)(nil)).Elem()
}

func (i AddRemoveReplicaScalingMechanismArgs) ToAddRemoveReplicaScalingMechanismOutput() AddRemoveReplicaScalingMechanismOutput {
	return i.ToAddRemoveReplicaScalingMechanismOutputWithContext(context.Background())
}

func (i AddRemoveReplicaScalingMechanismArgs) ToAddRemoveReplicaScalingMechanismOutputWithContext(ctx context.Context) AddRemoveReplicaScalingMechanismOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddRemoveReplicaScalingMechanismOutput)
}

type AddRemoveReplicaScalingMechanismOutput struct{ *pulumi.OutputState }

func (AddRemoveReplicaScalingMechanismOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddRemoveReplicaScalingMechanism)(nil)).Elem()
}

func (o AddRemoveReplicaScalingMechanismOutput) ToAddRemoveReplicaScalingMechanismOutput() AddRemoveReplicaScalingMechanismOutput {
	return o
}

func (o AddRemoveReplicaScalingMechanismOutput) ToAddRemoveReplicaScalingMechanismOutputWithContext(ctx context.Context) AddRemoveReplicaScalingMechanismOutput {
	return o
}

func (o AddRemoveReplicaScalingMechanismOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v AddRemoveReplicaScalingMechanism) string { return v.Kind }).(pulumi.StringOutput)
}

func (o AddRemoveReplicaScalingMechanismOutput) MaxCount() pulumi.IntOutput {
	return o.ApplyT(func(v AddRemoveReplicaScalingMechanism) int { return v.MaxCount }).(pulumi.IntOutput)
}

func (o AddRemoveReplicaScalingMechanismOutput) MinCount() pulumi.IntOutput {
	return o.ApplyT(func(v AddRemoveReplicaScalingMechanism) int { return v.MinCount }).(pulumi.IntOutput)
}

func (o AddRemoveReplicaScalingMechanismOutput) ScaleIncrement() pulumi.IntOutput {
	return o.ApplyT(func(v AddRemoveReplicaScalingMechanism) int { return v.ScaleIncrement }).(pulumi.IntOutput)
}

type AddRemoveReplicaScalingMechanismResponse struct {
	Kind           string `pulumi:"kind"`
	MaxCount       int    `pulumi:"maxCount"`
	MinCount       int    `pulumi:"minCount"`
	ScaleIncrement int    `pulumi:"scaleIncrement"`
}

type AddRemoveReplicaScalingMechanismResponseOutput struct{ *pulumi.OutputState }

func (AddRemoveReplicaScalingMechanismResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddRemoveReplicaScalingMechanismResponse)(nil)).Elem()
}

func (o AddRemoveReplicaScalingMechanismResponseOutput) ToAddRemoveReplicaScalingMechanismResponseOutput() AddRemoveReplicaScalingMechanismResponseOutput {
	return o
}

func (o AddRemoveReplicaScalingMechanismResponseOutput) ToAddRemoveReplicaScalingMechanismResponseOutputWithContext(ctx context.Context) AddRemoveReplicaScalingMechanismResponseOutput {
	return o
}

func (o AddRemoveReplicaScalingMechanismResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v AddRemoveReplicaScalingMechanismResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o AddRemoveReplicaScalingMechanismResponseOutput) MaxCount() pulumi.IntOutput {
	return o.ApplyT(func(v AddRemoveReplicaScalingMechanismResponse) int { return v.MaxCount }).(pulumi.IntOutput)
}

func (o AddRemoveReplicaScalingMechanismResponseOutput) MinCount() pulumi.IntOutput {
	return o.ApplyT(func(v AddRemoveReplicaScalingMechanismResponse) int { return v.MinCount }).(pulumi.IntOutput)
}

func (o AddRemoveReplicaScalingMechanismResponseOutput) ScaleIncrement() pulumi.IntOutput {
	return o.ApplyT(func(v AddRemoveReplicaScalingMechanismResponse) int { return v.ScaleIncrement }).(pulumi.IntOutput)
}

type ApplicationScopedVolume struct {
	CreationParameters ApplicationScopedVolumeCreationParametersServiceFabricVolumeDisk `pulumi:"creationParameters"`
	DestinationPath    string                                                           `pulumi:"destinationPath"`
	Name               string                                                           `pulumi:"name"`
	ReadOnly           *bool                                                            `pulumi:"readOnly"`
}





type ApplicationScopedVolumeInput interface {
	pulumi.Input

	ToApplicationScopedVolumeOutput() ApplicationScopedVolumeOutput
	ToApplicationScopedVolumeOutputWithContext(context.Context) ApplicationScopedVolumeOutput
}

type ApplicationScopedVolumeArgs struct {
	CreationParameters ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskInput `pulumi:"creationParameters"`
	DestinationPath    pulumi.StringInput                                                    `pulumi:"destinationPath"`
	Name               pulumi.StringInput                                                    `pulumi:"name"`
	ReadOnly           pulumi.BoolPtrInput                                                   `pulumi:"readOnly"`
}

func (ApplicationScopedVolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationScopedVolume)(nil)).Elem()
}

func (i ApplicationScopedVolumeArgs) ToApplicationScopedVolumeOutput() ApplicationScopedVolumeOutput {
	return i.ToApplicationScopedVolumeOutputWithContext(context.Background())
}

func (i ApplicationScopedVolumeArgs) ToApplicationScopedVolumeOutputWithContext(ctx context.Context) ApplicationScopedVolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationScopedVolumeOutput)
}





type ApplicationScopedVolumeArrayInput interface {
	pulumi.Input

	ToApplicationScopedVolumeArrayOutput() ApplicationScopedVolumeArrayOutput
	ToApplicationScopedVolumeArrayOutputWithContext(context.Context) ApplicationScopedVolumeArrayOutput
}

type ApplicationScopedVolumeArray []ApplicationScopedVolumeInput

func (ApplicationScopedVolumeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationScopedVolume)(nil)).Elem()
}

func (i ApplicationScopedVolumeArray) ToApplicationScopedVolumeArrayOutput() ApplicationScopedVolumeArrayOutput {
	return i.ToApplicationScopedVolumeArrayOutputWithContext(context.Background())
}

func (i ApplicationScopedVolumeArray) ToApplicationScopedVolumeArrayOutputWithContext(ctx context.Context) ApplicationScopedVolumeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationScopedVolumeArrayOutput)
}

type ApplicationScopedVolumeOutput struct{ *pulumi.OutputState }

func (ApplicationScopedVolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationScopedVolume)(nil)).Elem()
}

func (o ApplicationScopedVolumeOutput) ToApplicationScopedVolumeOutput() ApplicationScopedVolumeOutput {
	return o
}

func (o ApplicationScopedVolumeOutput) ToApplicationScopedVolumeOutputWithContext(ctx context.Context) ApplicationScopedVolumeOutput {
	return o
}

func (o ApplicationScopedVolumeOutput) CreationParameters() ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput {
	return o.ApplyT(func(v ApplicationScopedVolume) ApplicationScopedVolumeCreationParametersServiceFabricVolumeDisk {
		return v.CreationParameters
	}).(ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput)
}

func (o ApplicationScopedVolumeOutput) DestinationPath() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationScopedVolume) string { return v.DestinationPath }).(pulumi.StringOutput)
}

func (o ApplicationScopedVolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationScopedVolume) string { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationScopedVolumeOutput) ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationScopedVolume) *bool { return v.ReadOnly }).(pulumi.BoolPtrOutput)
}

type ApplicationScopedVolumeArrayOutput struct{ *pulumi.OutputState }

func (ApplicationScopedVolumeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationScopedVolume)(nil)).Elem()
}

func (o ApplicationScopedVolumeArrayOutput) ToApplicationScopedVolumeArrayOutput() ApplicationScopedVolumeArrayOutput {
	return o
}

func (o ApplicationScopedVolumeArrayOutput) ToApplicationScopedVolumeArrayOutputWithContext(ctx context.Context) ApplicationScopedVolumeArrayOutput {
	return o
}

func (o ApplicationScopedVolumeArrayOutput) Index(i pulumi.IntInput) ApplicationScopedVolumeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationScopedVolume {
		return vs[0].([]ApplicationScopedVolume)[vs[1].(int)]
	}).(ApplicationScopedVolumeOutput)
}

type ApplicationScopedVolumeCreationParametersServiceFabricVolumeDisk struct {
	Description *string `pulumi:"description"`
	Kind        string  `pulumi:"kind"`
	SizeDisk    string  `pulumi:"sizeDisk"`
}





type ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskInput interface {
	pulumi.Input

	ToApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput() ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput
	ToApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutputWithContext(context.Context) ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput
}

type ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Kind        pulumi.StringInput    `pulumi:"kind"`
	SizeDisk    pulumi.StringInput    `pulumi:"sizeDisk"`
}

func (ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationScopedVolumeCreationParametersServiceFabricVolumeDisk)(nil)).Elem()
}

func (i ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskArgs) ToApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput() ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput {
	return i.ToApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutputWithContext(context.Background())
}

func (i ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskArgs) ToApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutputWithContext(ctx context.Context) ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput)
}

type ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput struct{ *pulumi.OutputState }

func (ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationScopedVolumeCreationParametersServiceFabricVolumeDisk)(nil)).Elem()
}

func (o ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput) ToApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput() ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput {
	return o
}

func (o ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput) ToApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutputWithContext(ctx context.Context) ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput {
	return o
}

func (o ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationScopedVolumeCreationParametersServiceFabricVolumeDisk) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationScopedVolumeCreationParametersServiceFabricVolumeDisk) string { return v.Kind }).(pulumi.StringOutput)
}

func (o ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput) SizeDisk() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationScopedVolumeCreationParametersServiceFabricVolumeDisk) string { return v.SizeDisk }).(pulumi.StringOutput)
}

type ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponse struct {
	Description *string `pulumi:"description"`
	Kind        string  `pulumi:"kind"`
	SizeDisk    string  `pulumi:"sizeDisk"`
}

type ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponseOutput struct{ *pulumi.OutputState }

func (ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponse)(nil)).Elem()
}

func (o ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponseOutput) ToApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponseOutput() ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponseOutput {
	return o
}

func (o ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponseOutput) ToApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponseOutputWithContext(ctx context.Context) ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponseOutput {
	return o
}

func (o ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponse) *string {
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponseOutput) SizeDisk() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponse) string {
		return v.SizeDisk
	}).(pulumi.StringOutput)
}

type ApplicationScopedVolumeResponse struct {
	CreationParameters ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponse `pulumi:"creationParameters"`
	DestinationPath    string                                                                   `pulumi:"destinationPath"`
	Name               string                                                                   `pulumi:"name"`
	ReadOnly           *bool                                                                    `pulumi:"readOnly"`
}

type ApplicationScopedVolumeResponseOutput struct{ *pulumi.OutputState }

func (ApplicationScopedVolumeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationScopedVolumeResponse)(nil)).Elem()
}

func (o ApplicationScopedVolumeResponseOutput) ToApplicationScopedVolumeResponseOutput() ApplicationScopedVolumeResponseOutput {
	return o
}

func (o ApplicationScopedVolumeResponseOutput) ToApplicationScopedVolumeResponseOutputWithContext(ctx context.Context) ApplicationScopedVolumeResponseOutput {
	return o
}

func (o ApplicationScopedVolumeResponseOutput) CreationParameters() ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponseOutput {
	return o.ApplyT(func(v ApplicationScopedVolumeResponse) ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponse {
		return v.CreationParameters
	}).(ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponseOutput)
}

func (o ApplicationScopedVolumeResponseOutput) DestinationPath() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationScopedVolumeResponse) string { return v.DestinationPath }).(pulumi.StringOutput)
}

func (o ApplicationScopedVolumeResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationScopedVolumeResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationScopedVolumeResponseOutput) ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationScopedVolumeResponse) *bool { return v.ReadOnly }).(pulumi.BoolPtrOutput)
}

type ApplicationScopedVolumeResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationScopedVolumeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationScopedVolumeResponse)(nil)).Elem()
}

func (o ApplicationScopedVolumeResponseArrayOutput) ToApplicationScopedVolumeResponseArrayOutput() ApplicationScopedVolumeResponseArrayOutput {
	return o
}

func (o ApplicationScopedVolumeResponseArrayOutput) ToApplicationScopedVolumeResponseArrayOutputWithContext(ctx context.Context) ApplicationScopedVolumeResponseArrayOutput {
	return o
}

func (o ApplicationScopedVolumeResponseArrayOutput) Index(i pulumi.IntInput) ApplicationScopedVolumeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationScopedVolumeResponse {
		return vs[0].([]ApplicationScopedVolumeResponse)[vs[1].(int)]
	}).(ApplicationScopedVolumeResponseOutput)
}

type AutoScalingPolicy struct {
	Mechanism AddRemoveReplicaScalingMechanism `pulumi:"mechanism"`
	Name      string                           `pulumi:"name"`
	Trigger   AverageLoadScalingTrigger        `pulumi:"trigger"`
}





type AutoScalingPolicyInput interface {
	pulumi.Input

	ToAutoScalingPolicyOutput() AutoScalingPolicyOutput
	ToAutoScalingPolicyOutputWithContext(context.Context) AutoScalingPolicyOutput
}

type AutoScalingPolicyArgs struct {
	Mechanism AddRemoveReplicaScalingMechanismInput `pulumi:"mechanism"`
	Name      pulumi.StringInput                    `pulumi:"name"`
	Trigger   AverageLoadScalingTriggerInput        `pulumi:"trigger"`
}

func (AutoScalingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScalingPolicy)(nil)).Elem()
}

func (i AutoScalingPolicyArgs) ToAutoScalingPolicyOutput() AutoScalingPolicyOutput {
	return i.ToAutoScalingPolicyOutputWithContext(context.Background())
}

func (i AutoScalingPolicyArgs) ToAutoScalingPolicyOutputWithContext(ctx context.Context) AutoScalingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScalingPolicyOutput)
}





type AutoScalingPolicyArrayInput interface {
	pulumi.Input

	ToAutoScalingPolicyArrayOutput() AutoScalingPolicyArrayOutput
	ToAutoScalingPolicyArrayOutputWithContext(context.Context) AutoScalingPolicyArrayOutput
}

type AutoScalingPolicyArray []AutoScalingPolicyInput

func (AutoScalingPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutoScalingPolicy)(nil)).Elem()
}

func (i AutoScalingPolicyArray) ToAutoScalingPolicyArrayOutput() AutoScalingPolicyArrayOutput {
	return i.ToAutoScalingPolicyArrayOutputWithContext(context.Background())
}

func (i AutoScalingPolicyArray) ToAutoScalingPolicyArrayOutputWithContext(ctx context.Context) AutoScalingPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScalingPolicyArrayOutput)
}

type AutoScalingPolicyOutput struct{ *pulumi.OutputState }

func (AutoScalingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScalingPolicy)(nil)).Elem()
}

func (o AutoScalingPolicyOutput) ToAutoScalingPolicyOutput() AutoScalingPolicyOutput {
	return o
}

func (o AutoScalingPolicyOutput) ToAutoScalingPolicyOutputWithContext(ctx context.Context) AutoScalingPolicyOutput {
	return o
}

func (o AutoScalingPolicyOutput) Mechanism() AddRemoveReplicaScalingMechanismOutput {
	return o.ApplyT(func(v AutoScalingPolicy) AddRemoveReplicaScalingMechanism { return v.Mechanism }).(AddRemoveReplicaScalingMechanismOutput)
}

func (o AutoScalingPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AutoScalingPolicy) string { return v.Name }).(pulumi.StringOutput)
}

func (o AutoScalingPolicyOutput) Trigger() AverageLoadScalingTriggerOutput {
	return o.ApplyT(func(v AutoScalingPolicy) AverageLoadScalingTrigger { return v.Trigger }).(AverageLoadScalingTriggerOutput)
}

type AutoScalingPolicyArrayOutput struct{ *pulumi.OutputState }

func (AutoScalingPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutoScalingPolicy)(nil)).Elem()
}

func (o AutoScalingPolicyArrayOutput) ToAutoScalingPolicyArrayOutput() AutoScalingPolicyArrayOutput {
	return o
}

func (o AutoScalingPolicyArrayOutput) ToAutoScalingPolicyArrayOutputWithContext(ctx context.Context) AutoScalingPolicyArrayOutput {
	return o
}

func (o AutoScalingPolicyArrayOutput) Index(i pulumi.IntInput) AutoScalingPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutoScalingPolicy {
		return vs[0].([]AutoScalingPolicy)[vs[1].(int)]
	}).(AutoScalingPolicyOutput)
}

type AutoScalingPolicyResponse struct {
	Mechanism AddRemoveReplicaScalingMechanismResponse `pulumi:"mechanism"`
	Name      string                                   `pulumi:"name"`
	Trigger   AverageLoadScalingTriggerResponse        `pulumi:"trigger"`
}

type AutoScalingPolicyResponseOutput struct{ *pulumi.OutputState }

func (AutoScalingPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScalingPolicyResponse)(nil)).Elem()
}

func (o AutoScalingPolicyResponseOutput) ToAutoScalingPolicyResponseOutput() AutoScalingPolicyResponseOutput {
	return o
}

func (o AutoScalingPolicyResponseOutput) ToAutoScalingPolicyResponseOutputWithContext(ctx context.Context) AutoScalingPolicyResponseOutput {
	return o
}

func (o AutoScalingPolicyResponseOutput) Mechanism() AddRemoveReplicaScalingMechanismResponseOutput {
	return o.ApplyT(func(v AutoScalingPolicyResponse) AddRemoveReplicaScalingMechanismResponse { return v.Mechanism }).(AddRemoveReplicaScalingMechanismResponseOutput)
}

func (o AutoScalingPolicyResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AutoScalingPolicyResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o AutoScalingPolicyResponseOutput) Trigger() AverageLoadScalingTriggerResponseOutput {
	return o.ApplyT(func(v AutoScalingPolicyResponse) AverageLoadScalingTriggerResponse { return v.Trigger }).(AverageLoadScalingTriggerResponseOutput)
}

type AutoScalingPolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (AutoScalingPolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutoScalingPolicyResponse)(nil)).Elem()
}

func (o AutoScalingPolicyResponseArrayOutput) ToAutoScalingPolicyResponseArrayOutput() AutoScalingPolicyResponseArrayOutput {
	return o
}

func (o AutoScalingPolicyResponseArrayOutput) ToAutoScalingPolicyResponseArrayOutputWithContext(ctx context.Context) AutoScalingPolicyResponseArrayOutput {
	return o
}

func (o AutoScalingPolicyResponseArrayOutput) Index(i pulumi.IntInput) AutoScalingPolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutoScalingPolicyResponse {
		return vs[0].([]AutoScalingPolicyResponse)[vs[1].(int)]
	}).(AutoScalingPolicyResponseOutput)
}

type AutoScalingResourceMetric struct {
	Kind string `pulumi:"kind"`
	Name string `pulumi:"name"`
}





type AutoScalingResourceMetricInput interface {
	pulumi.Input

	ToAutoScalingResourceMetricOutput() AutoScalingResourceMetricOutput
	ToAutoScalingResourceMetricOutputWithContext(context.Context) AutoScalingResourceMetricOutput
}

type AutoScalingResourceMetricArgs struct {
	Kind pulumi.StringInput `pulumi:"kind"`
	Name pulumi.StringInput `pulumi:"name"`
}

func (AutoScalingResourceMetricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScalingResourceMetric)(nil)).Elem()
}

func (i AutoScalingResourceMetricArgs) ToAutoScalingResourceMetricOutput() AutoScalingResourceMetricOutput {
	return i.ToAutoScalingResourceMetricOutputWithContext(context.Background())
}

func (i AutoScalingResourceMetricArgs) ToAutoScalingResourceMetricOutputWithContext(ctx context.Context) AutoScalingResourceMetricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoScalingResourceMetricOutput)
}

type AutoScalingResourceMetricOutput struct{ *pulumi.OutputState }

func (AutoScalingResourceMetricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScalingResourceMetric)(nil)).Elem()
}

func (o AutoScalingResourceMetricOutput) ToAutoScalingResourceMetricOutput() AutoScalingResourceMetricOutput {
	return o
}

func (o AutoScalingResourceMetricOutput) ToAutoScalingResourceMetricOutputWithContext(ctx context.Context) AutoScalingResourceMetricOutput {
	return o
}

func (o AutoScalingResourceMetricOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v AutoScalingResourceMetric) string { return v.Kind }).(pulumi.StringOutput)
}

func (o AutoScalingResourceMetricOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AutoScalingResourceMetric) string { return v.Name }).(pulumi.StringOutput)
}

type AutoScalingResourceMetricResponse struct {
	Kind string `pulumi:"kind"`
	Name string `pulumi:"name"`
}

type AutoScalingResourceMetricResponseOutput struct{ *pulumi.OutputState }

func (AutoScalingResourceMetricResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoScalingResourceMetricResponse)(nil)).Elem()
}

func (o AutoScalingResourceMetricResponseOutput) ToAutoScalingResourceMetricResponseOutput() AutoScalingResourceMetricResponseOutput {
	return o
}

func (o AutoScalingResourceMetricResponseOutput) ToAutoScalingResourceMetricResponseOutputWithContext(ctx context.Context) AutoScalingResourceMetricResponseOutput {
	return o
}

func (o AutoScalingResourceMetricResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v AutoScalingResourceMetricResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o AutoScalingResourceMetricResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AutoScalingResourceMetricResponse) string { return v.Name }).(pulumi.StringOutput)
}

type AverageLoadScalingTrigger struct {
	Kind                   string                    `pulumi:"kind"`
	LowerLoadThreshold     float64                   `pulumi:"lowerLoadThreshold"`
	Metric                 AutoScalingResourceMetric `pulumi:"metric"`
	ScaleIntervalInSeconds int                       `pulumi:"scaleIntervalInSeconds"`
	UpperLoadThreshold     float64                   `pulumi:"upperLoadThreshold"`
}





type AverageLoadScalingTriggerInput interface {
	pulumi.Input

	ToAverageLoadScalingTriggerOutput() AverageLoadScalingTriggerOutput
	ToAverageLoadScalingTriggerOutputWithContext(context.Context) AverageLoadScalingTriggerOutput
}

type AverageLoadScalingTriggerArgs struct {
	Kind                   pulumi.StringInput             `pulumi:"kind"`
	LowerLoadThreshold     pulumi.Float64Input            `pulumi:"lowerLoadThreshold"`
	Metric                 AutoScalingResourceMetricInput `pulumi:"metric"`
	ScaleIntervalInSeconds pulumi.IntInput                `pulumi:"scaleIntervalInSeconds"`
	UpperLoadThreshold     pulumi.Float64Input            `pulumi:"upperLoadThreshold"`
}

func (AverageLoadScalingTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AverageLoadScalingTrigger)(nil)).Elem()
}

func (i AverageLoadScalingTriggerArgs) ToAverageLoadScalingTriggerOutput() AverageLoadScalingTriggerOutput {
	return i.ToAverageLoadScalingTriggerOutputWithContext(context.Background())
}

func (i AverageLoadScalingTriggerArgs) ToAverageLoadScalingTriggerOutputWithContext(ctx context.Context) AverageLoadScalingTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AverageLoadScalingTriggerOutput)
}

type AverageLoadScalingTriggerOutput struct{ *pulumi.OutputState }

func (AverageLoadScalingTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AverageLoadScalingTrigger)(nil)).Elem()
}

func (o AverageLoadScalingTriggerOutput) ToAverageLoadScalingTriggerOutput() AverageLoadScalingTriggerOutput {
	return o
}

func (o AverageLoadScalingTriggerOutput) ToAverageLoadScalingTriggerOutputWithContext(ctx context.Context) AverageLoadScalingTriggerOutput {
	return o
}

func (o AverageLoadScalingTriggerOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v AverageLoadScalingTrigger) string { return v.Kind }).(pulumi.StringOutput)
}

func (o AverageLoadScalingTriggerOutput) LowerLoadThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v AverageLoadScalingTrigger) float64 { return v.LowerLoadThreshold }).(pulumi.Float64Output)
}

func (o AverageLoadScalingTriggerOutput) Metric() AutoScalingResourceMetricOutput {
	return o.ApplyT(func(v AverageLoadScalingTrigger) AutoScalingResourceMetric { return v.Metric }).(AutoScalingResourceMetricOutput)
}

func (o AverageLoadScalingTriggerOutput) ScaleIntervalInSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v AverageLoadScalingTrigger) int { return v.ScaleIntervalInSeconds }).(pulumi.IntOutput)
}

func (o AverageLoadScalingTriggerOutput) UpperLoadThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v AverageLoadScalingTrigger) float64 { return v.UpperLoadThreshold }).(pulumi.Float64Output)
}

type AverageLoadScalingTriggerResponse struct {
	Kind                   string                            `pulumi:"kind"`
	LowerLoadThreshold     float64                           `pulumi:"lowerLoadThreshold"`
	Metric                 AutoScalingResourceMetricResponse `pulumi:"metric"`
	ScaleIntervalInSeconds int                               `pulumi:"scaleIntervalInSeconds"`
	UpperLoadThreshold     float64                           `pulumi:"upperLoadThreshold"`
}

type AverageLoadScalingTriggerResponseOutput struct{ *pulumi.OutputState }

func (AverageLoadScalingTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AverageLoadScalingTriggerResponse)(nil)).Elem()
}

func (o AverageLoadScalingTriggerResponseOutput) ToAverageLoadScalingTriggerResponseOutput() AverageLoadScalingTriggerResponseOutput {
	return o
}

func (o AverageLoadScalingTriggerResponseOutput) ToAverageLoadScalingTriggerResponseOutputWithContext(ctx context.Context) AverageLoadScalingTriggerResponseOutput {
	return o
}

func (o AverageLoadScalingTriggerResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v AverageLoadScalingTriggerResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o AverageLoadScalingTriggerResponseOutput) LowerLoadThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v AverageLoadScalingTriggerResponse) float64 { return v.LowerLoadThreshold }).(pulumi.Float64Output)
}

func (o AverageLoadScalingTriggerResponseOutput) Metric() AutoScalingResourceMetricResponseOutput {
	return o.ApplyT(func(v AverageLoadScalingTriggerResponse) AutoScalingResourceMetricResponse { return v.Metric }).(AutoScalingResourceMetricResponseOutput)
}

func (o AverageLoadScalingTriggerResponseOutput) ScaleIntervalInSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v AverageLoadScalingTriggerResponse) int { return v.ScaleIntervalInSeconds }).(pulumi.IntOutput)
}

func (o AverageLoadScalingTriggerResponseOutput) UpperLoadThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v AverageLoadScalingTriggerResponse) float64 { return v.UpperLoadThreshold }).(pulumi.Float64Output)
}

type AzureInternalMonitoringPipelineSinkDescription struct {
	AccountName      *string     `pulumi:"accountName"`
	AutoKeyConfigUrl *string     `pulumi:"autoKeyConfigUrl"`
	Description      *string     `pulumi:"description"`
	FluentdConfigUrl interface{} `pulumi:"fluentdConfigUrl"`
	Kind             string      `pulumi:"kind"`
	MaConfigUrl      *string     `pulumi:"maConfigUrl"`
	Name             *string     `pulumi:"name"`
	Namespace        *string     `pulumi:"namespace"`
}





type AzureInternalMonitoringPipelineSinkDescriptionInput interface {
	pulumi.Input

	ToAzureInternalMonitoringPipelineSinkDescriptionOutput() AzureInternalMonitoringPipelineSinkDescriptionOutput
	ToAzureInternalMonitoringPipelineSinkDescriptionOutputWithContext(context.Context) AzureInternalMonitoringPipelineSinkDescriptionOutput
}

type AzureInternalMonitoringPipelineSinkDescriptionArgs struct {
	AccountName      pulumi.StringPtrInput `pulumi:"accountName"`
	AutoKeyConfigUrl pulumi.StringPtrInput `pulumi:"autoKeyConfigUrl"`
	Description      pulumi.StringPtrInput `pulumi:"description"`
	FluentdConfigUrl pulumi.Input          `pulumi:"fluentdConfigUrl"`
	Kind             pulumi.StringInput    `pulumi:"kind"`
	MaConfigUrl      pulumi.StringPtrInput `pulumi:"maConfigUrl"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	Namespace        pulumi.StringPtrInput `pulumi:"namespace"`
}

func (AzureInternalMonitoringPipelineSinkDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureInternalMonitoringPipelineSinkDescription)(nil)).Elem()
}

func (i AzureInternalMonitoringPipelineSinkDescriptionArgs) ToAzureInternalMonitoringPipelineSinkDescriptionOutput() AzureInternalMonitoringPipelineSinkDescriptionOutput {
	return i.ToAzureInternalMonitoringPipelineSinkDescriptionOutputWithContext(context.Background())
}

func (i AzureInternalMonitoringPipelineSinkDescriptionArgs) ToAzureInternalMonitoringPipelineSinkDescriptionOutputWithContext(ctx context.Context) AzureInternalMonitoringPipelineSinkDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureInternalMonitoringPipelineSinkDescriptionOutput)
}





type AzureInternalMonitoringPipelineSinkDescriptionArrayInput interface {
	pulumi.Input

	ToAzureInternalMonitoringPipelineSinkDescriptionArrayOutput() AzureInternalMonitoringPipelineSinkDescriptionArrayOutput
	ToAzureInternalMonitoringPipelineSinkDescriptionArrayOutputWithContext(context.Context) AzureInternalMonitoringPipelineSinkDescriptionArrayOutput
}

type AzureInternalMonitoringPipelineSinkDescriptionArray []AzureInternalMonitoringPipelineSinkDescriptionInput

func (AzureInternalMonitoringPipelineSinkDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureInternalMonitoringPipelineSinkDescription)(nil)).Elem()
}

func (i AzureInternalMonitoringPipelineSinkDescriptionArray) ToAzureInternalMonitoringPipelineSinkDescriptionArrayOutput() AzureInternalMonitoringPipelineSinkDescriptionArrayOutput {
	return i.ToAzureInternalMonitoringPipelineSinkDescriptionArrayOutputWithContext(context.Background())
}

func (i AzureInternalMonitoringPipelineSinkDescriptionArray) ToAzureInternalMonitoringPipelineSinkDescriptionArrayOutputWithContext(ctx context.Context) AzureInternalMonitoringPipelineSinkDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureInternalMonitoringPipelineSinkDescriptionArrayOutput)
}

type AzureInternalMonitoringPipelineSinkDescriptionOutput struct{ *pulumi.OutputState }

func (AzureInternalMonitoringPipelineSinkDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureInternalMonitoringPipelineSinkDescription)(nil)).Elem()
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) ToAzureInternalMonitoringPipelineSinkDescriptionOutput() AzureInternalMonitoringPipelineSinkDescriptionOutput {
	return o
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) ToAzureInternalMonitoringPipelineSinkDescriptionOutputWithContext(ctx context.Context) AzureInternalMonitoringPipelineSinkDescriptionOutput {
	return o
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescription) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) AutoKeyConfigUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescription) *string { return v.AutoKeyConfigUrl }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescription) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) FluentdConfigUrl() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescription) interface{} { return v.FluentdConfigUrl }).(pulumi.AnyOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescription) string { return v.Kind }).(pulumi.StringOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) MaConfigUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescription) *string { return v.MaConfigUrl }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescription) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescription) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

type AzureInternalMonitoringPipelineSinkDescriptionArrayOutput struct{ *pulumi.OutputState }

func (AzureInternalMonitoringPipelineSinkDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureInternalMonitoringPipelineSinkDescription)(nil)).Elem()
}

func (o AzureInternalMonitoringPipelineSinkDescriptionArrayOutput) ToAzureInternalMonitoringPipelineSinkDescriptionArrayOutput() AzureInternalMonitoringPipelineSinkDescriptionArrayOutput {
	return o
}

func (o AzureInternalMonitoringPipelineSinkDescriptionArrayOutput) ToAzureInternalMonitoringPipelineSinkDescriptionArrayOutputWithContext(ctx context.Context) AzureInternalMonitoringPipelineSinkDescriptionArrayOutput {
	return o
}

func (o AzureInternalMonitoringPipelineSinkDescriptionArrayOutput) Index(i pulumi.IntInput) AzureInternalMonitoringPipelineSinkDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureInternalMonitoringPipelineSinkDescription {
		return vs[0].([]AzureInternalMonitoringPipelineSinkDescription)[vs[1].(int)]
	}).(AzureInternalMonitoringPipelineSinkDescriptionOutput)
}

type AzureInternalMonitoringPipelineSinkDescriptionResponse struct {
	AccountName      *string     `pulumi:"accountName"`
	AutoKeyConfigUrl *string     `pulumi:"autoKeyConfigUrl"`
	Description      *string     `pulumi:"description"`
	FluentdConfigUrl interface{} `pulumi:"fluentdConfigUrl"`
	Kind             string      `pulumi:"kind"`
	MaConfigUrl      *string     `pulumi:"maConfigUrl"`
	Name             *string     `pulumi:"name"`
	Namespace        *string     `pulumi:"namespace"`
}

type AzureInternalMonitoringPipelineSinkDescriptionResponseOutput struct{ *pulumi.OutputState }

func (AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureInternalMonitoringPipelineSinkDescriptionResponse)(nil)).Elem()
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) ToAzureInternalMonitoringPipelineSinkDescriptionResponseOutput() AzureInternalMonitoringPipelineSinkDescriptionResponseOutput {
	return o
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) ToAzureInternalMonitoringPipelineSinkDescriptionResponseOutputWithContext(ctx context.Context) AzureInternalMonitoringPipelineSinkDescriptionResponseOutput {
	return o
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescriptionResponse) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) AutoKeyConfigUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescriptionResponse) *string { return v.AutoKeyConfigUrl }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescriptionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) FluentdConfigUrl() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescriptionResponse) interface{} { return v.FluentdConfigUrl }).(pulumi.AnyOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescriptionResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) MaConfigUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescriptionResponse) *string { return v.MaConfigUrl }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescriptionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescriptionResponse) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

type AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureInternalMonitoringPipelineSinkDescriptionResponse)(nil)).Elem()
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput) ToAzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput() AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput {
	return o
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput) ToAzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutputWithContext(ctx context.Context) AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput {
	return o
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput) Index(i pulumi.IntInput) AzureInternalMonitoringPipelineSinkDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureInternalMonitoringPipelineSinkDescriptionResponse {
		return vs[0].([]AzureInternalMonitoringPipelineSinkDescriptionResponse)[vs[1].(int)]
	}).(AzureInternalMonitoringPipelineSinkDescriptionResponseOutput)
}

type ContainerCodePackageProperties struct {
	Commands                []string                  `pulumi:"commands"`
	Diagnostics             *DiagnosticsRef           `pulumi:"diagnostics"`
	Endpoints               []EndpointProperties      `pulumi:"endpoints"`
	Entrypoint              *string                   `pulumi:"entrypoint"`
	EnvironmentVariables    []EnvironmentVariable     `pulumi:"environmentVariables"`
	Image                   string                    `pulumi:"image"`
	ImageRegistryCredential *ImageRegistryCredential  `pulumi:"imageRegistryCredential"`
	Labels                  []ContainerLabel          `pulumi:"labels"`
	Name                    string                    `pulumi:"name"`
	ReliableCollectionsRefs []ReliableCollectionsRef  `pulumi:"reliableCollectionsRefs"`
	Resources               ResourceRequirements      `pulumi:"resources"`
	Settings                []Setting                 `pulumi:"settings"`
	VolumeRefs              []VolumeReference         `pulumi:"volumeRefs"`
	Volumes                 []ApplicationScopedVolume `pulumi:"volumes"`
}





type ContainerCodePackagePropertiesInput interface {
	pulumi.Input

	ToContainerCodePackagePropertiesOutput() ContainerCodePackagePropertiesOutput
	ToContainerCodePackagePropertiesOutputWithContext(context.Context) ContainerCodePackagePropertiesOutput
}

type ContainerCodePackagePropertiesArgs struct {
	Commands                pulumi.StringArrayInput           `pulumi:"commands"`
	Diagnostics             DiagnosticsRefPtrInput            `pulumi:"diagnostics"`
	Endpoints               EndpointPropertiesArrayInput      `pulumi:"endpoints"`
	Entrypoint              pulumi.StringPtrInput             `pulumi:"entrypoint"`
	EnvironmentVariables    EnvironmentVariableArrayInput     `pulumi:"environmentVariables"`
	Image                   pulumi.StringInput                `pulumi:"image"`
	ImageRegistryCredential ImageRegistryCredentialPtrInput   `pulumi:"imageRegistryCredential"`
	Labels                  ContainerLabelArrayInput          `pulumi:"labels"`
	Name                    pulumi.StringInput                `pulumi:"name"`
	ReliableCollectionsRefs ReliableCollectionsRefArrayInput  `pulumi:"reliableCollectionsRefs"`
	Resources               ResourceRequirementsInput         `pulumi:"resources"`
	Settings                SettingArrayInput                 `pulumi:"settings"`
	VolumeRefs              VolumeReferenceArrayInput         `pulumi:"volumeRefs"`
	Volumes                 ApplicationScopedVolumeArrayInput `pulumi:"volumes"`
}

func (ContainerCodePackagePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerCodePackageProperties)(nil)).Elem()
}

func (i ContainerCodePackagePropertiesArgs) ToContainerCodePackagePropertiesOutput() ContainerCodePackagePropertiesOutput {
	return i.ToContainerCodePackagePropertiesOutputWithContext(context.Background())
}

func (i ContainerCodePackagePropertiesArgs) ToContainerCodePackagePropertiesOutputWithContext(ctx context.Context) ContainerCodePackagePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerCodePackagePropertiesOutput)
}





type ContainerCodePackagePropertiesArrayInput interface {
	pulumi.Input

	ToContainerCodePackagePropertiesArrayOutput() ContainerCodePackagePropertiesArrayOutput
	ToContainerCodePackagePropertiesArrayOutputWithContext(context.Context) ContainerCodePackagePropertiesArrayOutput
}

type ContainerCodePackagePropertiesArray []ContainerCodePackagePropertiesInput

func (ContainerCodePackagePropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerCodePackageProperties)(nil)).Elem()
}

func (i ContainerCodePackagePropertiesArray) ToContainerCodePackagePropertiesArrayOutput() ContainerCodePackagePropertiesArrayOutput {
	return i.ToContainerCodePackagePropertiesArrayOutputWithContext(context.Background())
}

func (i ContainerCodePackagePropertiesArray) ToContainerCodePackagePropertiesArrayOutputWithContext(ctx context.Context) ContainerCodePackagePropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerCodePackagePropertiesArrayOutput)
}

type ContainerCodePackagePropertiesOutput struct{ *pulumi.OutputState }

func (ContainerCodePackagePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerCodePackageProperties)(nil)).Elem()
}

func (o ContainerCodePackagePropertiesOutput) ToContainerCodePackagePropertiesOutput() ContainerCodePackagePropertiesOutput {
	return o
}

func (o ContainerCodePackagePropertiesOutput) ToContainerCodePackagePropertiesOutputWithContext(ctx context.Context) ContainerCodePackagePropertiesOutput {
	return o
}

func (o ContainerCodePackagePropertiesOutput) Commands() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) []string { return v.Commands }).(pulumi.StringArrayOutput)
}

func (o ContainerCodePackagePropertiesOutput) Diagnostics() DiagnosticsRefPtrOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) *DiagnosticsRef { return v.Diagnostics }).(DiagnosticsRefPtrOutput)
}

func (o ContainerCodePackagePropertiesOutput) Endpoints() EndpointPropertiesArrayOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) []EndpointProperties { return v.Endpoints }).(EndpointPropertiesArrayOutput)
}

func (o ContainerCodePackagePropertiesOutput) Entrypoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) *string { return v.Entrypoint }).(pulumi.StringPtrOutput)
}

func (o ContainerCodePackagePropertiesOutput) EnvironmentVariables() EnvironmentVariableArrayOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) []EnvironmentVariable { return v.EnvironmentVariables }).(EnvironmentVariableArrayOutput)
}

func (o ContainerCodePackagePropertiesOutput) Image() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) string { return v.Image }).(pulumi.StringOutput)
}

func (o ContainerCodePackagePropertiesOutput) ImageRegistryCredential() ImageRegistryCredentialPtrOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) *ImageRegistryCredential { return v.ImageRegistryCredential }).(ImageRegistryCredentialPtrOutput)
}

func (o ContainerCodePackagePropertiesOutput) Labels() ContainerLabelArrayOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) []ContainerLabel { return v.Labels }).(ContainerLabelArrayOutput)
}

func (o ContainerCodePackagePropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerCodePackagePropertiesOutput) ReliableCollectionsRefs() ReliableCollectionsRefArrayOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) []ReliableCollectionsRef { return v.ReliableCollectionsRefs }).(ReliableCollectionsRefArrayOutput)
}

func (o ContainerCodePackagePropertiesOutput) Resources() ResourceRequirementsOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) ResourceRequirements { return v.Resources }).(ResourceRequirementsOutput)
}

func (o ContainerCodePackagePropertiesOutput) Settings() SettingArrayOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) []Setting { return v.Settings }).(SettingArrayOutput)
}

func (o ContainerCodePackagePropertiesOutput) VolumeRefs() VolumeReferenceArrayOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) []VolumeReference { return v.VolumeRefs }).(VolumeReferenceArrayOutput)
}

func (o ContainerCodePackagePropertiesOutput) Volumes() ApplicationScopedVolumeArrayOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) []ApplicationScopedVolume { return v.Volumes }).(ApplicationScopedVolumeArrayOutput)
}

type ContainerCodePackagePropertiesArrayOutput struct{ *pulumi.OutputState }

func (ContainerCodePackagePropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerCodePackageProperties)(nil)).Elem()
}

func (o ContainerCodePackagePropertiesArrayOutput) ToContainerCodePackagePropertiesArrayOutput() ContainerCodePackagePropertiesArrayOutput {
	return o
}

func (o ContainerCodePackagePropertiesArrayOutput) ToContainerCodePackagePropertiesArrayOutputWithContext(ctx context.Context) ContainerCodePackagePropertiesArrayOutput {
	return o
}

func (o ContainerCodePackagePropertiesArrayOutput) Index(i pulumi.IntInput) ContainerCodePackagePropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerCodePackageProperties {
		return vs[0].([]ContainerCodePackageProperties)[vs[1].(int)]
	}).(ContainerCodePackagePropertiesOutput)
}

type ContainerCodePackagePropertiesResponse struct {
	Commands                []string                          `pulumi:"commands"`
	Diagnostics             *DiagnosticsRefResponse           `pulumi:"diagnostics"`
	Endpoints               []EndpointPropertiesResponse      `pulumi:"endpoints"`
	Entrypoint              *string                           `pulumi:"entrypoint"`
	EnvironmentVariables    []EnvironmentVariableResponse     `pulumi:"environmentVariables"`
	Image                   string                            `pulumi:"image"`
	ImageRegistryCredential *ImageRegistryCredentialResponse  `pulumi:"imageRegistryCredential"`
	InstanceView            ContainerInstanceViewResponse     `pulumi:"instanceView"`
	Labels                  []ContainerLabelResponse          `pulumi:"labels"`
	Name                    string                            `pulumi:"name"`
	ReliableCollectionsRefs []ReliableCollectionsRefResponse  `pulumi:"reliableCollectionsRefs"`
	Resources               ResourceRequirementsResponse      `pulumi:"resources"`
	Settings                []SettingResponse                 `pulumi:"settings"`
	VolumeRefs              []VolumeReferenceResponse         `pulumi:"volumeRefs"`
	Volumes                 []ApplicationScopedVolumeResponse `pulumi:"volumes"`
}

type ContainerCodePackagePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ContainerCodePackagePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerCodePackagePropertiesResponse)(nil)).Elem()
}

func (o ContainerCodePackagePropertiesResponseOutput) ToContainerCodePackagePropertiesResponseOutput() ContainerCodePackagePropertiesResponseOutput {
	return o
}

func (o ContainerCodePackagePropertiesResponseOutput) ToContainerCodePackagePropertiesResponseOutputWithContext(ctx context.Context) ContainerCodePackagePropertiesResponseOutput {
	return o
}

func (o ContainerCodePackagePropertiesResponseOutput) Commands() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) []string { return v.Commands }).(pulumi.StringArrayOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) Diagnostics() DiagnosticsRefResponsePtrOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) *DiagnosticsRefResponse { return v.Diagnostics }).(DiagnosticsRefResponsePtrOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) Endpoints() EndpointPropertiesResponseArrayOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) []EndpointPropertiesResponse { return v.Endpoints }).(EndpointPropertiesResponseArrayOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) Entrypoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) *string { return v.Entrypoint }).(pulumi.StringPtrOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) EnvironmentVariables() EnvironmentVariableResponseArrayOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) []EnvironmentVariableResponse {
		return v.EnvironmentVariables
	}).(EnvironmentVariableResponseArrayOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) Image() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) string { return v.Image }).(pulumi.StringOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) ImageRegistryCredential() ImageRegistryCredentialResponsePtrOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) *ImageRegistryCredentialResponse {
		return v.ImageRegistryCredential
	}).(ImageRegistryCredentialResponsePtrOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) InstanceView() ContainerInstanceViewResponseOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) ContainerInstanceViewResponse { return v.InstanceView }).(ContainerInstanceViewResponseOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) Labels() ContainerLabelResponseArrayOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) []ContainerLabelResponse { return v.Labels }).(ContainerLabelResponseArrayOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) ReliableCollectionsRefs() ReliableCollectionsRefResponseArrayOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) []ReliableCollectionsRefResponse {
		return v.ReliableCollectionsRefs
	}).(ReliableCollectionsRefResponseArrayOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) Resources() ResourceRequirementsResponseOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) ResourceRequirementsResponse { return v.Resources }).(ResourceRequirementsResponseOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) Settings() SettingResponseArrayOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) []SettingResponse { return v.Settings }).(SettingResponseArrayOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) VolumeRefs() VolumeReferenceResponseArrayOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) []VolumeReferenceResponse { return v.VolumeRefs }).(VolumeReferenceResponseArrayOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) Volumes() ApplicationScopedVolumeResponseArrayOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) []ApplicationScopedVolumeResponse { return v.Volumes }).(ApplicationScopedVolumeResponseArrayOutput)
}

type ContainerCodePackagePropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerCodePackagePropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerCodePackagePropertiesResponse)(nil)).Elem()
}

func (o ContainerCodePackagePropertiesResponseArrayOutput) ToContainerCodePackagePropertiesResponseArrayOutput() ContainerCodePackagePropertiesResponseArrayOutput {
	return o
}

func (o ContainerCodePackagePropertiesResponseArrayOutput) ToContainerCodePackagePropertiesResponseArrayOutputWithContext(ctx context.Context) ContainerCodePackagePropertiesResponseArrayOutput {
	return o
}

func (o ContainerCodePackagePropertiesResponseArrayOutput) Index(i pulumi.IntInput) ContainerCodePackagePropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerCodePackagePropertiesResponse {
		return vs[0].([]ContainerCodePackagePropertiesResponse)[vs[1].(int)]
	}).(ContainerCodePackagePropertiesResponseOutput)
}

type ContainerEventResponse struct {
	Count          *int    `pulumi:"count"`
	FirstTimestamp *string `pulumi:"firstTimestamp"`
	LastTimestamp  *string `pulumi:"lastTimestamp"`
	Message        *string `pulumi:"message"`
	Name           *string `pulumi:"name"`
	Type           *string `pulumi:"type"`
}

type ContainerEventResponseOutput struct{ *pulumi.OutputState }

func (ContainerEventResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerEventResponse)(nil)).Elem()
}

func (o ContainerEventResponseOutput) ToContainerEventResponseOutput() ContainerEventResponseOutput {
	return o
}

func (o ContainerEventResponseOutput) ToContainerEventResponseOutputWithContext(ctx context.Context) ContainerEventResponseOutput {
	return o
}

func (o ContainerEventResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerEventResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o ContainerEventResponseOutput) FirstTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerEventResponse) *string { return v.FirstTimestamp }).(pulumi.StringPtrOutput)
}

func (o ContainerEventResponseOutput) LastTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerEventResponse) *string { return v.LastTimestamp }).(pulumi.StringPtrOutput)
}

func (o ContainerEventResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerEventResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ContainerEventResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerEventResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ContainerEventResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerEventResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ContainerEventResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerEventResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerEventResponse)(nil)).Elem()
}

func (o ContainerEventResponseArrayOutput) ToContainerEventResponseArrayOutput() ContainerEventResponseArrayOutput {
	return o
}

func (o ContainerEventResponseArrayOutput) ToContainerEventResponseArrayOutputWithContext(ctx context.Context) ContainerEventResponseArrayOutput {
	return o
}

func (o ContainerEventResponseArrayOutput) Index(i pulumi.IntInput) ContainerEventResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerEventResponse {
		return vs[0].([]ContainerEventResponse)[vs[1].(int)]
	}).(ContainerEventResponseOutput)
}

type ContainerInstanceViewResponse struct {
	CurrentState  *ContainerStateResponse  `pulumi:"currentState"`
	Events        []ContainerEventResponse `pulumi:"events"`
	PreviousState *ContainerStateResponse  `pulumi:"previousState"`
	RestartCount  *int                     `pulumi:"restartCount"`
}

type ContainerInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (ContainerInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerInstanceViewResponse)(nil)).Elem()
}

func (o ContainerInstanceViewResponseOutput) ToContainerInstanceViewResponseOutput() ContainerInstanceViewResponseOutput {
	return o
}

func (o ContainerInstanceViewResponseOutput) ToContainerInstanceViewResponseOutputWithContext(ctx context.Context) ContainerInstanceViewResponseOutput {
	return o
}

func (o ContainerInstanceViewResponseOutput) CurrentState() ContainerStateResponsePtrOutput {
	return o.ApplyT(func(v ContainerInstanceViewResponse) *ContainerStateResponse { return v.CurrentState }).(ContainerStateResponsePtrOutput)
}

func (o ContainerInstanceViewResponseOutput) Events() ContainerEventResponseArrayOutput {
	return o.ApplyT(func(v ContainerInstanceViewResponse) []ContainerEventResponse { return v.Events }).(ContainerEventResponseArrayOutput)
}

func (o ContainerInstanceViewResponseOutput) PreviousState() ContainerStateResponsePtrOutput {
	return o.ApplyT(func(v ContainerInstanceViewResponse) *ContainerStateResponse { return v.PreviousState }).(ContainerStateResponsePtrOutput)
}

func (o ContainerInstanceViewResponseOutput) RestartCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerInstanceViewResponse) *int { return v.RestartCount }).(pulumi.IntPtrOutput)
}

type ContainerLabel struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type ContainerLabelInput interface {
	pulumi.Input

	ToContainerLabelOutput() ContainerLabelOutput
	ToContainerLabelOutputWithContext(context.Context) ContainerLabelOutput
}

type ContainerLabelArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (ContainerLabelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerLabel)(nil)).Elem()
}

func (i ContainerLabelArgs) ToContainerLabelOutput() ContainerLabelOutput {
	return i.ToContainerLabelOutputWithContext(context.Background())
}

func (i ContainerLabelArgs) ToContainerLabelOutputWithContext(ctx context.Context) ContainerLabelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerLabelOutput)
}





type ContainerLabelArrayInput interface {
	pulumi.Input

	ToContainerLabelArrayOutput() ContainerLabelArrayOutput
	ToContainerLabelArrayOutputWithContext(context.Context) ContainerLabelArrayOutput
}

type ContainerLabelArray []ContainerLabelInput

func (ContainerLabelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerLabel)(nil)).Elem()
}

func (i ContainerLabelArray) ToContainerLabelArrayOutput() ContainerLabelArrayOutput {
	return i.ToContainerLabelArrayOutputWithContext(context.Background())
}

func (i ContainerLabelArray) ToContainerLabelArrayOutputWithContext(ctx context.Context) ContainerLabelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerLabelArrayOutput)
}

type ContainerLabelOutput struct{ *pulumi.OutputState }

func (ContainerLabelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerLabel)(nil)).Elem()
}

func (o ContainerLabelOutput) ToContainerLabelOutput() ContainerLabelOutput {
	return o
}

func (o ContainerLabelOutput) ToContainerLabelOutputWithContext(ctx context.Context) ContainerLabelOutput {
	return o
}

func (o ContainerLabelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerLabel) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerLabelOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerLabel) string { return v.Value }).(pulumi.StringOutput)
}

type ContainerLabelArrayOutput struct{ *pulumi.OutputState }

func (ContainerLabelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerLabel)(nil)).Elem()
}

func (o ContainerLabelArrayOutput) ToContainerLabelArrayOutput() ContainerLabelArrayOutput {
	return o
}

func (o ContainerLabelArrayOutput) ToContainerLabelArrayOutputWithContext(ctx context.Context) ContainerLabelArrayOutput {
	return o
}

func (o ContainerLabelArrayOutput) Index(i pulumi.IntInput) ContainerLabelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerLabel {
		return vs[0].([]ContainerLabel)[vs[1].(int)]
	}).(ContainerLabelOutput)
}

type ContainerLabelResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type ContainerLabelResponseOutput struct{ *pulumi.OutputState }

func (ContainerLabelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerLabelResponse)(nil)).Elem()
}

func (o ContainerLabelResponseOutput) ToContainerLabelResponseOutput() ContainerLabelResponseOutput {
	return o
}

func (o ContainerLabelResponseOutput) ToContainerLabelResponseOutputWithContext(ctx context.Context) ContainerLabelResponseOutput {
	return o
}

func (o ContainerLabelResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerLabelResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerLabelResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerLabelResponse) string { return v.Value }).(pulumi.StringOutput)
}

type ContainerLabelResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerLabelResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerLabelResponse)(nil)).Elem()
}

func (o ContainerLabelResponseArrayOutput) ToContainerLabelResponseArrayOutput() ContainerLabelResponseArrayOutput {
	return o
}

func (o ContainerLabelResponseArrayOutput) ToContainerLabelResponseArrayOutputWithContext(ctx context.Context) ContainerLabelResponseArrayOutput {
	return o
}

func (o ContainerLabelResponseArrayOutput) Index(i pulumi.IntInput) ContainerLabelResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerLabelResponse {
		return vs[0].([]ContainerLabelResponse)[vs[1].(int)]
	}).(ContainerLabelResponseOutput)
}

type ContainerStateResponse struct {
	DetailStatus *string `pulumi:"detailStatus"`
	ExitCode     *string `pulumi:"exitCode"`
	FinishTime   *string `pulumi:"finishTime"`
	StartTime    *string `pulumi:"startTime"`
	State        *string `pulumi:"state"`
}

type ContainerStateResponseOutput struct{ *pulumi.OutputState }

func (ContainerStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerStateResponse)(nil)).Elem()
}

func (o ContainerStateResponseOutput) ToContainerStateResponseOutput() ContainerStateResponseOutput {
	return o
}

func (o ContainerStateResponseOutput) ToContainerStateResponseOutputWithContext(ctx context.Context) ContainerStateResponseOutput {
	return o
}

func (o ContainerStateResponseOutput) DetailStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerStateResponse) *string { return v.DetailStatus }).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponseOutput) ExitCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerStateResponse) *string { return v.ExitCode }).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponseOutput) FinishTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerStateResponse) *string { return v.FinishTime }).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerStateResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerStateResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type ContainerStateResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerStateResponse)(nil)).Elem()
}

func (o ContainerStateResponsePtrOutput) ToContainerStateResponsePtrOutput() ContainerStateResponsePtrOutput {
	return o
}

func (o ContainerStateResponsePtrOutput) ToContainerStateResponsePtrOutputWithContext(ctx context.Context) ContainerStateResponsePtrOutput {
	return o
}

func (o ContainerStateResponsePtrOutput) Elem() ContainerStateResponseOutput {
	return o.ApplyT(func(v *ContainerStateResponse) ContainerStateResponse {
		if v != nil {
			return *v
		}
		var ret ContainerStateResponse
		return ret
	}).(ContainerStateResponseOutput)
}

func (o ContainerStateResponsePtrOutput) DetailStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.DetailStatus
	}).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponsePtrOutput) ExitCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExitCode
	}).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponsePtrOutput) FinishTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.FinishTime
	}).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type DiagnosticsDescription struct {
	DefaultSinkRefs []string                                         `pulumi:"defaultSinkRefs"`
	Enabled         *bool                                            `pulumi:"enabled"`
	Sinks           []AzureInternalMonitoringPipelineSinkDescription `pulumi:"sinks"`
}





type DiagnosticsDescriptionInput interface {
	pulumi.Input

	ToDiagnosticsDescriptionOutput() DiagnosticsDescriptionOutput
	ToDiagnosticsDescriptionOutputWithContext(context.Context) DiagnosticsDescriptionOutput
}

type DiagnosticsDescriptionArgs struct {
	DefaultSinkRefs pulumi.StringArrayInput                                  `pulumi:"defaultSinkRefs"`
	Enabled         pulumi.BoolPtrInput                                      `pulumi:"enabled"`
	Sinks           AzureInternalMonitoringPipelineSinkDescriptionArrayInput `pulumi:"sinks"`
}

func (DiagnosticsDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsDescription)(nil)).Elem()
}

func (i DiagnosticsDescriptionArgs) ToDiagnosticsDescriptionOutput() DiagnosticsDescriptionOutput {
	return i.ToDiagnosticsDescriptionOutputWithContext(context.Background())
}

func (i DiagnosticsDescriptionArgs) ToDiagnosticsDescriptionOutputWithContext(ctx context.Context) DiagnosticsDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsDescriptionOutput)
}

func (i DiagnosticsDescriptionArgs) ToDiagnosticsDescriptionPtrOutput() DiagnosticsDescriptionPtrOutput {
	return i.ToDiagnosticsDescriptionPtrOutputWithContext(context.Background())
}

func (i DiagnosticsDescriptionArgs) ToDiagnosticsDescriptionPtrOutputWithContext(ctx context.Context) DiagnosticsDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsDescriptionOutput).ToDiagnosticsDescriptionPtrOutputWithContext(ctx)
}









type DiagnosticsDescriptionPtrInput interface {
	pulumi.Input

	ToDiagnosticsDescriptionPtrOutput() DiagnosticsDescriptionPtrOutput
	ToDiagnosticsDescriptionPtrOutputWithContext(context.Context) DiagnosticsDescriptionPtrOutput
}

type diagnosticsDescriptionPtrType DiagnosticsDescriptionArgs

func DiagnosticsDescriptionPtr(v *DiagnosticsDescriptionArgs) DiagnosticsDescriptionPtrInput {
	return (*diagnosticsDescriptionPtrType)(v)
}

func (*diagnosticsDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsDescription)(nil)).Elem()
}

func (i *diagnosticsDescriptionPtrType) ToDiagnosticsDescriptionPtrOutput() DiagnosticsDescriptionPtrOutput {
	return i.ToDiagnosticsDescriptionPtrOutputWithContext(context.Background())
}

func (i *diagnosticsDescriptionPtrType) ToDiagnosticsDescriptionPtrOutputWithContext(ctx context.Context) DiagnosticsDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsDescriptionPtrOutput)
}

type DiagnosticsDescriptionOutput struct{ *pulumi.OutputState }

func (DiagnosticsDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsDescription)(nil)).Elem()
}

func (o DiagnosticsDescriptionOutput) ToDiagnosticsDescriptionOutput() DiagnosticsDescriptionOutput {
	return o
}

func (o DiagnosticsDescriptionOutput) ToDiagnosticsDescriptionOutputWithContext(ctx context.Context) DiagnosticsDescriptionOutput {
	return o
}

func (o DiagnosticsDescriptionOutput) ToDiagnosticsDescriptionPtrOutput() DiagnosticsDescriptionPtrOutput {
	return o.ToDiagnosticsDescriptionPtrOutputWithContext(context.Background())
}

func (o DiagnosticsDescriptionOutput) ToDiagnosticsDescriptionPtrOutputWithContext(ctx context.Context) DiagnosticsDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiagnosticsDescription) *DiagnosticsDescription {
		return &v
	}).(DiagnosticsDescriptionPtrOutput)
}

func (o DiagnosticsDescriptionOutput) DefaultSinkRefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DiagnosticsDescription) []string { return v.DefaultSinkRefs }).(pulumi.StringArrayOutput)
}

func (o DiagnosticsDescriptionOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DiagnosticsDescription) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o DiagnosticsDescriptionOutput) Sinks() AzureInternalMonitoringPipelineSinkDescriptionArrayOutput {
	return o.ApplyT(func(v DiagnosticsDescription) []AzureInternalMonitoringPipelineSinkDescription { return v.Sinks }).(AzureInternalMonitoringPipelineSinkDescriptionArrayOutput)
}

type DiagnosticsDescriptionPtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsDescription)(nil)).Elem()
}

func (o DiagnosticsDescriptionPtrOutput) ToDiagnosticsDescriptionPtrOutput() DiagnosticsDescriptionPtrOutput {
	return o
}

func (o DiagnosticsDescriptionPtrOutput) ToDiagnosticsDescriptionPtrOutputWithContext(ctx context.Context) DiagnosticsDescriptionPtrOutput {
	return o
}

func (o DiagnosticsDescriptionPtrOutput) Elem() DiagnosticsDescriptionOutput {
	return o.ApplyT(func(v *DiagnosticsDescription) DiagnosticsDescription {
		if v != nil {
			return *v
		}
		var ret DiagnosticsDescription
		return ret
	}).(DiagnosticsDescriptionOutput)
}

func (o DiagnosticsDescriptionPtrOutput) DefaultSinkRefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DiagnosticsDescription) []string {
		if v == nil {
			return nil
		}
		return v.DefaultSinkRefs
	}).(pulumi.StringArrayOutput)
}

func (o DiagnosticsDescriptionPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiagnosticsDescription) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o DiagnosticsDescriptionPtrOutput) Sinks() AzureInternalMonitoringPipelineSinkDescriptionArrayOutput {
	return o.ApplyT(func(v *DiagnosticsDescription) []AzureInternalMonitoringPipelineSinkDescription {
		if v == nil {
			return nil
		}
		return v.Sinks
	}).(AzureInternalMonitoringPipelineSinkDescriptionArrayOutput)
}

type DiagnosticsDescriptionResponse struct {
	DefaultSinkRefs []string                                                 `pulumi:"defaultSinkRefs"`
	Enabled         *bool                                                    `pulumi:"enabled"`
	Sinks           []AzureInternalMonitoringPipelineSinkDescriptionResponse `pulumi:"sinks"`
}

type DiagnosticsDescriptionResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticsDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsDescriptionResponse)(nil)).Elem()
}

func (o DiagnosticsDescriptionResponseOutput) ToDiagnosticsDescriptionResponseOutput() DiagnosticsDescriptionResponseOutput {
	return o
}

func (o DiagnosticsDescriptionResponseOutput) ToDiagnosticsDescriptionResponseOutputWithContext(ctx context.Context) DiagnosticsDescriptionResponseOutput {
	return o
}

func (o DiagnosticsDescriptionResponseOutput) DefaultSinkRefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DiagnosticsDescriptionResponse) []string { return v.DefaultSinkRefs }).(pulumi.StringArrayOutput)
}

func (o DiagnosticsDescriptionResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DiagnosticsDescriptionResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o DiagnosticsDescriptionResponseOutput) Sinks() AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput {
	return o.ApplyT(func(v DiagnosticsDescriptionResponse) []AzureInternalMonitoringPipelineSinkDescriptionResponse {
		return v.Sinks
	}).(AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput)
}

type DiagnosticsDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsDescriptionResponse)(nil)).Elem()
}

func (o DiagnosticsDescriptionResponsePtrOutput) ToDiagnosticsDescriptionResponsePtrOutput() DiagnosticsDescriptionResponsePtrOutput {
	return o
}

func (o DiagnosticsDescriptionResponsePtrOutput) ToDiagnosticsDescriptionResponsePtrOutputWithContext(ctx context.Context) DiagnosticsDescriptionResponsePtrOutput {
	return o
}

func (o DiagnosticsDescriptionResponsePtrOutput) Elem() DiagnosticsDescriptionResponseOutput {
	return o.ApplyT(func(v *DiagnosticsDescriptionResponse) DiagnosticsDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret DiagnosticsDescriptionResponse
		return ret
	}).(DiagnosticsDescriptionResponseOutput)
}

func (o DiagnosticsDescriptionResponsePtrOutput) DefaultSinkRefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DiagnosticsDescriptionResponse) []string {
		if v == nil {
			return nil
		}
		return v.DefaultSinkRefs
	}).(pulumi.StringArrayOutput)
}

func (o DiagnosticsDescriptionResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiagnosticsDescriptionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o DiagnosticsDescriptionResponsePtrOutput) Sinks() AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput {
	return o.ApplyT(func(v *DiagnosticsDescriptionResponse) []AzureInternalMonitoringPipelineSinkDescriptionResponse {
		if v == nil {
			return nil
		}
		return v.Sinks
	}).(AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput)
}

type DiagnosticsRef struct {
	Enabled  *bool    `pulumi:"enabled"`
	SinkRefs []string `pulumi:"sinkRefs"`
}





type DiagnosticsRefInput interface {
	pulumi.Input

	ToDiagnosticsRefOutput() DiagnosticsRefOutput
	ToDiagnosticsRefOutputWithContext(context.Context) DiagnosticsRefOutput
}

type DiagnosticsRefArgs struct {
	Enabled  pulumi.BoolPtrInput     `pulumi:"enabled"`
	SinkRefs pulumi.StringArrayInput `pulumi:"sinkRefs"`
}

func (DiagnosticsRefArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsRef)(nil)).Elem()
}

func (i DiagnosticsRefArgs) ToDiagnosticsRefOutput() DiagnosticsRefOutput {
	return i.ToDiagnosticsRefOutputWithContext(context.Background())
}

func (i DiagnosticsRefArgs) ToDiagnosticsRefOutputWithContext(ctx context.Context) DiagnosticsRefOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsRefOutput)
}

func (i DiagnosticsRefArgs) ToDiagnosticsRefPtrOutput() DiagnosticsRefPtrOutput {
	return i.ToDiagnosticsRefPtrOutputWithContext(context.Background())
}

func (i DiagnosticsRefArgs) ToDiagnosticsRefPtrOutputWithContext(ctx context.Context) DiagnosticsRefPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsRefOutput).ToDiagnosticsRefPtrOutputWithContext(ctx)
}









type DiagnosticsRefPtrInput interface {
	pulumi.Input

	ToDiagnosticsRefPtrOutput() DiagnosticsRefPtrOutput
	ToDiagnosticsRefPtrOutputWithContext(context.Context) DiagnosticsRefPtrOutput
}

type diagnosticsRefPtrType DiagnosticsRefArgs

func DiagnosticsRefPtr(v *DiagnosticsRefArgs) DiagnosticsRefPtrInput {
	return (*diagnosticsRefPtrType)(v)
}

func (*diagnosticsRefPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsRef)(nil)).Elem()
}

func (i *diagnosticsRefPtrType) ToDiagnosticsRefPtrOutput() DiagnosticsRefPtrOutput {
	return i.ToDiagnosticsRefPtrOutputWithContext(context.Background())
}

func (i *diagnosticsRefPtrType) ToDiagnosticsRefPtrOutputWithContext(ctx context.Context) DiagnosticsRefPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsRefPtrOutput)
}

type DiagnosticsRefOutput struct{ *pulumi.OutputState }

func (DiagnosticsRefOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsRef)(nil)).Elem()
}

func (o DiagnosticsRefOutput) ToDiagnosticsRefOutput() DiagnosticsRefOutput {
	return o
}

func (o DiagnosticsRefOutput) ToDiagnosticsRefOutputWithContext(ctx context.Context) DiagnosticsRefOutput {
	return o
}

func (o DiagnosticsRefOutput) ToDiagnosticsRefPtrOutput() DiagnosticsRefPtrOutput {
	return o.ToDiagnosticsRefPtrOutputWithContext(context.Background())
}

func (o DiagnosticsRefOutput) ToDiagnosticsRefPtrOutputWithContext(ctx context.Context) DiagnosticsRefPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiagnosticsRef) *DiagnosticsRef {
		return &v
	}).(DiagnosticsRefPtrOutput)
}

func (o DiagnosticsRefOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DiagnosticsRef) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o DiagnosticsRefOutput) SinkRefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DiagnosticsRef) []string { return v.SinkRefs }).(pulumi.StringArrayOutput)
}

type DiagnosticsRefPtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsRefPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsRef)(nil)).Elem()
}

func (o DiagnosticsRefPtrOutput) ToDiagnosticsRefPtrOutput() DiagnosticsRefPtrOutput {
	return o
}

func (o DiagnosticsRefPtrOutput) ToDiagnosticsRefPtrOutputWithContext(ctx context.Context) DiagnosticsRefPtrOutput {
	return o
}

func (o DiagnosticsRefPtrOutput) Elem() DiagnosticsRefOutput {
	return o.ApplyT(func(v *DiagnosticsRef) DiagnosticsRef {
		if v != nil {
			return *v
		}
		var ret DiagnosticsRef
		return ret
	}).(DiagnosticsRefOutput)
}

func (o DiagnosticsRefPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiagnosticsRef) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o DiagnosticsRefPtrOutput) SinkRefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DiagnosticsRef) []string {
		if v == nil {
			return nil
		}
		return v.SinkRefs
	}).(pulumi.StringArrayOutput)
}

type DiagnosticsRefResponse struct {
	Enabled  *bool    `pulumi:"enabled"`
	SinkRefs []string `pulumi:"sinkRefs"`
}

type DiagnosticsRefResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticsRefResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsRefResponse)(nil)).Elem()
}

func (o DiagnosticsRefResponseOutput) ToDiagnosticsRefResponseOutput() DiagnosticsRefResponseOutput {
	return o
}

func (o DiagnosticsRefResponseOutput) ToDiagnosticsRefResponseOutputWithContext(ctx context.Context) DiagnosticsRefResponseOutput {
	return o
}

func (o DiagnosticsRefResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DiagnosticsRefResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o DiagnosticsRefResponseOutput) SinkRefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DiagnosticsRefResponse) []string { return v.SinkRefs }).(pulumi.StringArrayOutput)
}

type DiagnosticsRefResponsePtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsRefResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsRefResponse)(nil)).Elem()
}

func (o DiagnosticsRefResponsePtrOutput) ToDiagnosticsRefResponsePtrOutput() DiagnosticsRefResponsePtrOutput {
	return o
}

func (o DiagnosticsRefResponsePtrOutput) ToDiagnosticsRefResponsePtrOutputWithContext(ctx context.Context) DiagnosticsRefResponsePtrOutput {
	return o
}

func (o DiagnosticsRefResponsePtrOutput) Elem() DiagnosticsRefResponseOutput {
	return o.ApplyT(func(v *DiagnosticsRefResponse) DiagnosticsRefResponse {
		if v != nil {
			return *v
		}
		var ret DiagnosticsRefResponse
		return ret
	}).(DiagnosticsRefResponseOutput)
}

func (o DiagnosticsRefResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiagnosticsRefResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o DiagnosticsRefResponsePtrOutput) SinkRefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DiagnosticsRefResponse) []string {
		if v == nil {
			return nil
		}
		return v.SinkRefs
	}).(pulumi.StringArrayOutput)
}

type EndpointProperties struct {
	Name string `pulumi:"name"`
	Port *int   `pulumi:"port"`
}





type EndpointPropertiesInput interface {
	pulumi.Input

	ToEndpointPropertiesOutput() EndpointPropertiesOutput
	ToEndpointPropertiesOutputWithContext(context.Context) EndpointPropertiesOutput
}

type EndpointPropertiesArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Port pulumi.IntPtrInput `pulumi:"port"`
}

func (EndpointPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointProperties)(nil)).Elem()
}

func (i EndpointPropertiesArgs) ToEndpointPropertiesOutput() EndpointPropertiesOutput {
	return i.ToEndpointPropertiesOutputWithContext(context.Background())
}

func (i EndpointPropertiesArgs) ToEndpointPropertiesOutputWithContext(ctx context.Context) EndpointPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesOutput)
}





type EndpointPropertiesArrayInput interface {
	pulumi.Input

	ToEndpointPropertiesArrayOutput() EndpointPropertiesArrayOutput
	ToEndpointPropertiesArrayOutputWithContext(context.Context) EndpointPropertiesArrayOutput
}

type EndpointPropertiesArray []EndpointPropertiesInput

func (EndpointPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointProperties)(nil)).Elem()
}

func (i EndpointPropertiesArray) ToEndpointPropertiesArrayOutput() EndpointPropertiesArrayOutput {
	return i.ToEndpointPropertiesArrayOutputWithContext(context.Background())
}

func (i EndpointPropertiesArray) ToEndpointPropertiesArrayOutputWithContext(ctx context.Context) EndpointPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesArrayOutput)
}

type EndpointPropertiesOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointProperties)(nil)).Elem()
}

func (o EndpointPropertiesOutput) ToEndpointPropertiesOutput() EndpointPropertiesOutput {
	return o
}

func (o EndpointPropertiesOutput) ToEndpointPropertiesOutputWithContext(ctx context.Context) EndpointPropertiesOutput {
	return o
}

func (o EndpointPropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointProperties) string { return v.Name }).(pulumi.StringOutput)
}

func (o EndpointPropertiesOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EndpointProperties) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type EndpointPropertiesArrayOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointProperties)(nil)).Elem()
}

func (o EndpointPropertiesArrayOutput) ToEndpointPropertiesArrayOutput() EndpointPropertiesArrayOutput {
	return o
}

func (o EndpointPropertiesArrayOutput) ToEndpointPropertiesArrayOutputWithContext(ctx context.Context) EndpointPropertiesArrayOutput {
	return o
}

func (o EndpointPropertiesArrayOutput) Index(i pulumi.IntInput) EndpointPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointProperties {
		return vs[0].([]EndpointProperties)[vs[1].(int)]
	}).(EndpointPropertiesOutput)
}

type EndpointPropertiesResponse struct {
	Name string `pulumi:"name"`
	Port *int   `pulumi:"port"`
}

type EndpointPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesResponse)(nil)).Elem()
}

func (o EndpointPropertiesResponseOutput) ToEndpointPropertiesResponseOutput() EndpointPropertiesResponseOutput {
	return o
}

func (o EndpointPropertiesResponseOutput) ToEndpointPropertiesResponseOutputWithContext(ctx context.Context) EndpointPropertiesResponseOutput {
	return o
}

func (o EndpointPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o EndpointPropertiesResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type EndpointPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointPropertiesResponse)(nil)).Elem()
}

func (o EndpointPropertiesResponseArrayOutput) ToEndpointPropertiesResponseArrayOutput() EndpointPropertiesResponseArrayOutput {
	return o
}

func (o EndpointPropertiesResponseArrayOutput) ToEndpointPropertiesResponseArrayOutputWithContext(ctx context.Context) EndpointPropertiesResponseArrayOutput {
	return o
}

func (o EndpointPropertiesResponseArrayOutput) Index(i pulumi.IntInput) EndpointPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointPropertiesResponse {
		return vs[0].([]EndpointPropertiesResponse)[vs[1].(int)]
	}).(EndpointPropertiesResponseOutput)
}

type EndpointRef struct {
	Name *string `pulumi:"name"`
}





type EndpointRefInput interface {
	pulumi.Input

	ToEndpointRefOutput() EndpointRefOutput
	ToEndpointRefOutputWithContext(context.Context) EndpointRefOutput
}

type EndpointRefArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (EndpointRefArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointRef)(nil)).Elem()
}

func (i EndpointRefArgs) ToEndpointRefOutput() EndpointRefOutput {
	return i.ToEndpointRefOutputWithContext(context.Background())
}

func (i EndpointRefArgs) ToEndpointRefOutputWithContext(ctx context.Context) EndpointRefOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRefOutput)
}





type EndpointRefArrayInput interface {
	pulumi.Input

	ToEndpointRefArrayOutput() EndpointRefArrayOutput
	ToEndpointRefArrayOutputWithContext(context.Context) EndpointRefArrayOutput
}

type EndpointRefArray []EndpointRefInput

func (EndpointRefArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointRef)(nil)).Elem()
}

func (i EndpointRefArray) ToEndpointRefArrayOutput() EndpointRefArrayOutput {
	return i.ToEndpointRefArrayOutputWithContext(context.Background())
}

func (i EndpointRefArray) ToEndpointRefArrayOutputWithContext(ctx context.Context) EndpointRefArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRefArrayOutput)
}

type EndpointRefOutput struct{ *pulumi.OutputState }

func (EndpointRefOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointRef)(nil)).Elem()
}

func (o EndpointRefOutput) ToEndpointRefOutput() EndpointRefOutput {
	return o
}

func (o EndpointRefOutput) ToEndpointRefOutputWithContext(ctx context.Context) EndpointRefOutput {
	return o
}

func (o EndpointRefOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointRef) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type EndpointRefArrayOutput struct{ *pulumi.OutputState }

func (EndpointRefArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointRef)(nil)).Elem()
}

func (o EndpointRefArrayOutput) ToEndpointRefArrayOutput() EndpointRefArrayOutput {
	return o
}

func (o EndpointRefArrayOutput) ToEndpointRefArrayOutputWithContext(ctx context.Context) EndpointRefArrayOutput {
	return o
}

func (o EndpointRefArrayOutput) Index(i pulumi.IntInput) EndpointRefOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointRef {
		return vs[0].([]EndpointRef)[vs[1].(int)]
	}).(EndpointRefOutput)
}

type EndpointRefResponse struct {
	Name *string `pulumi:"name"`
}

type EndpointRefResponseOutput struct{ *pulumi.OutputState }

func (EndpointRefResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointRefResponse)(nil)).Elem()
}

func (o EndpointRefResponseOutput) ToEndpointRefResponseOutput() EndpointRefResponseOutput {
	return o
}

func (o EndpointRefResponseOutput) ToEndpointRefResponseOutputWithContext(ctx context.Context) EndpointRefResponseOutput {
	return o
}

func (o EndpointRefResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointRefResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type EndpointRefResponseArrayOutput struct{ *pulumi.OutputState }

func (EndpointRefResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointRefResponse)(nil)).Elem()
}

func (o EndpointRefResponseArrayOutput) ToEndpointRefResponseArrayOutput() EndpointRefResponseArrayOutput {
	return o
}

func (o EndpointRefResponseArrayOutput) ToEndpointRefResponseArrayOutputWithContext(ctx context.Context) EndpointRefResponseArrayOutput {
	return o
}

func (o EndpointRefResponseArrayOutput) Index(i pulumi.IntInput) EndpointRefResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointRefResponse {
		return vs[0].([]EndpointRefResponse)[vs[1].(int)]
	}).(EndpointRefResponseOutput)
}

type EnvironmentVariable struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type EnvironmentVariableInput interface {
	pulumi.Input

	ToEnvironmentVariableOutput() EnvironmentVariableOutput
	ToEnvironmentVariableOutputWithContext(context.Context) EnvironmentVariableOutput
}

type EnvironmentVariableArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (EnvironmentVariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVariable)(nil)).Elem()
}

func (i EnvironmentVariableArgs) ToEnvironmentVariableOutput() EnvironmentVariableOutput {
	return i.ToEnvironmentVariableOutputWithContext(context.Background())
}

func (i EnvironmentVariableArgs) ToEnvironmentVariableOutputWithContext(ctx context.Context) EnvironmentVariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentVariableOutput)
}





type EnvironmentVariableArrayInput interface {
	pulumi.Input

	ToEnvironmentVariableArrayOutput() EnvironmentVariableArrayOutput
	ToEnvironmentVariableArrayOutputWithContext(context.Context) EnvironmentVariableArrayOutput
}

type EnvironmentVariableArray []EnvironmentVariableInput

func (EnvironmentVariableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentVariable)(nil)).Elem()
}

func (i EnvironmentVariableArray) ToEnvironmentVariableArrayOutput() EnvironmentVariableArrayOutput {
	return i.ToEnvironmentVariableArrayOutputWithContext(context.Background())
}

func (i EnvironmentVariableArray) ToEnvironmentVariableArrayOutputWithContext(ctx context.Context) EnvironmentVariableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentVariableArrayOutput)
}

type EnvironmentVariableOutput struct{ *pulumi.OutputState }

func (EnvironmentVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVariable)(nil)).Elem()
}

func (o EnvironmentVariableOutput) ToEnvironmentVariableOutput() EnvironmentVariableOutput {
	return o
}

func (o EnvironmentVariableOutput) ToEnvironmentVariableOutputWithContext(ctx context.Context) EnvironmentVariableOutput {
	return o
}

func (o EnvironmentVariableOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVariable) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVariableOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVariable) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type EnvironmentVariableArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentVariableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentVariable)(nil)).Elem()
}

func (o EnvironmentVariableArrayOutput) ToEnvironmentVariableArrayOutput() EnvironmentVariableArrayOutput {
	return o
}

func (o EnvironmentVariableArrayOutput) ToEnvironmentVariableArrayOutputWithContext(ctx context.Context) EnvironmentVariableArrayOutput {
	return o
}

func (o EnvironmentVariableArrayOutput) Index(i pulumi.IntInput) EnvironmentVariableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentVariable {
		return vs[0].([]EnvironmentVariable)[vs[1].(int)]
	}).(EnvironmentVariableOutput)
}

type EnvironmentVariableResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type EnvironmentVariableResponseOutput struct{ *pulumi.OutputState }

func (EnvironmentVariableResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVariableResponse)(nil)).Elem()
}

func (o EnvironmentVariableResponseOutput) ToEnvironmentVariableResponseOutput() EnvironmentVariableResponseOutput {
	return o
}

func (o EnvironmentVariableResponseOutput) ToEnvironmentVariableResponseOutputWithContext(ctx context.Context) EnvironmentVariableResponseOutput {
	return o
}

func (o EnvironmentVariableResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVariableResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVariableResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVariableResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type EnvironmentVariableResponseArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentVariableResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentVariableResponse)(nil)).Elem()
}

func (o EnvironmentVariableResponseArrayOutput) ToEnvironmentVariableResponseArrayOutput() EnvironmentVariableResponseArrayOutput {
	return o
}

func (o EnvironmentVariableResponseArrayOutput) ToEnvironmentVariableResponseArrayOutputWithContext(ctx context.Context) EnvironmentVariableResponseArrayOutput {
	return o
}

func (o EnvironmentVariableResponseArrayOutput) Index(i pulumi.IntInput) EnvironmentVariableResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentVariableResponse {
		return vs[0].([]EnvironmentVariableResponse)[vs[1].(int)]
	}).(EnvironmentVariableResponseOutput)
}

type GatewayDestination struct {
	ApplicationName string `pulumi:"applicationName"`
	EndpointName    string `pulumi:"endpointName"`
	ServiceName     string `pulumi:"serviceName"`
}





type GatewayDestinationInput interface {
	pulumi.Input

	ToGatewayDestinationOutput() GatewayDestinationOutput
	ToGatewayDestinationOutputWithContext(context.Context) GatewayDestinationOutput
}

type GatewayDestinationArgs struct {
	ApplicationName pulumi.StringInput `pulumi:"applicationName"`
	EndpointName    pulumi.StringInput `pulumi:"endpointName"`
	ServiceName     pulumi.StringInput `pulumi:"serviceName"`
}

func (GatewayDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayDestination)(nil)).Elem()
}

func (i GatewayDestinationArgs) ToGatewayDestinationOutput() GatewayDestinationOutput {
	return i.ToGatewayDestinationOutputWithContext(context.Background())
}

func (i GatewayDestinationArgs) ToGatewayDestinationOutputWithContext(ctx context.Context) GatewayDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayDestinationOutput)
}

type GatewayDestinationOutput struct{ *pulumi.OutputState }

func (GatewayDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayDestination)(nil)).Elem()
}

func (o GatewayDestinationOutput) ToGatewayDestinationOutput() GatewayDestinationOutput {
	return o
}

func (o GatewayDestinationOutput) ToGatewayDestinationOutputWithContext(ctx context.Context) GatewayDestinationOutput {
	return o
}

func (o GatewayDestinationOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayDestination) string { return v.ApplicationName }).(pulumi.StringOutput)
}

func (o GatewayDestinationOutput) EndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayDestination) string { return v.EndpointName }).(pulumi.StringOutput)
}

func (o GatewayDestinationOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayDestination) string { return v.ServiceName }).(pulumi.StringOutput)
}

type GatewayDestinationResponse struct {
	ApplicationName string `pulumi:"applicationName"`
	EndpointName    string `pulumi:"endpointName"`
	ServiceName     string `pulumi:"serviceName"`
}

type GatewayDestinationResponseOutput struct{ *pulumi.OutputState }

func (GatewayDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayDestinationResponse)(nil)).Elem()
}

func (o GatewayDestinationResponseOutput) ToGatewayDestinationResponseOutput() GatewayDestinationResponseOutput {
	return o
}

func (o GatewayDestinationResponseOutput) ToGatewayDestinationResponseOutputWithContext(ctx context.Context) GatewayDestinationResponseOutput {
	return o
}

func (o GatewayDestinationResponseOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayDestinationResponse) string { return v.ApplicationName }).(pulumi.StringOutput)
}

func (o GatewayDestinationResponseOutput) EndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayDestinationResponse) string { return v.EndpointName }).(pulumi.StringOutput)
}

func (o GatewayDestinationResponseOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayDestinationResponse) string { return v.ServiceName }).(pulumi.StringOutput)
}

type HttpConfig struct {
	Hosts []HttpHostConfig `pulumi:"hosts"`
	Name  string           `pulumi:"name"`
	Port  int              `pulumi:"port"`
}





type HttpConfigInput interface {
	pulumi.Input

	ToHttpConfigOutput() HttpConfigOutput
	ToHttpConfigOutputWithContext(context.Context) HttpConfigOutput
}

type HttpConfigArgs struct {
	Hosts HttpHostConfigArrayInput `pulumi:"hosts"`
	Name  pulumi.StringInput       `pulumi:"name"`
	Port  pulumi.IntInput          `pulumi:"port"`
}

func (HttpConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpConfig)(nil)).Elem()
}

func (i HttpConfigArgs) ToHttpConfigOutput() HttpConfigOutput {
	return i.ToHttpConfigOutputWithContext(context.Background())
}

func (i HttpConfigArgs) ToHttpConfigOutputWithContext(ctx context.Context) HttpConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpConfigOutput)
}





type HttpConfigArrayInput interface {
	pulumi.Input

	ToHttpConfigArrayOutput() HttpConfigArrayOutput
	ToHttpConfigArrayOutputWithContext(context.Context) HttpConfigArrayOutput
}

type HttpConfigArray []HttpConfigInput

func (HttpConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpConfig)(nil)).Elem()
}

func (i HttpConfigArray) ToHttpConfigArrayOutput() HttpConfigArrayOutput {
	return i.ToHttpConfigArrayOutputWithContext(context.Background())
}

func (i HttpConfigArray) ToHttpConfigArrayOutputWithContext(ctx context.Context) HttpConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpConfigArrayOutput)
}

type HttpConfigOutput struct{ *pulumi.OutputState }

func (HttpConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpConfig)(nil)).Elem()
}

func (o HttpConfigOutput) ToHttpConfigOutput() HttpConfigOutput {
	return o
}

func (o HttpConfigOutput) ToHttpConfigOutputWithContext(ctx context.Context) HttpConfigOutput {
	return o
}

func (o HttpConfigOutput) Hosts() HttpHostConfigArrayOutput {
	return o.ApplyT(func(v HttpConfig) []HttpHostConfig { return v.Hosts }).(HttpHostConfigArrayOutput)
}

func (o HttpConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v HttpConfig) string { return v.Name }).(pulumi.StringOutput)
}

func (o HttpConfigOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v HttpConfig) int { return v.Port }).(pulumi.IntOutput)
}

type HttpConfigArrayOutput struct{ *pulumi.OutputState }

func (HttpConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpConfig)(nil)).Elem()
}

func (o HttpConfigArrayOutput) ToHttpConfigArrayOutput() HttpConfigArrayOutput {
	return o
}

func (o HttpConfigArrayOutput) ToHttpConfigArrayOutputWithContext(ctx context.Context) HttpConfigArrayOutput {
	return o
}

func (o HttpConfigArrayOutput) Index(i pulumi.IntInput) HttpConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HttpConfig {
		return vs[0].([]HttpConfig)[vs[1].(int)]
	}).(HttpConfigOutput)
}

type HttpConfigResponse struct {
	Hosts []HttpHostConfigResponse `pulumi:"hosts"`
	Name  string                   `pulumi:"name"`
	Port  int                      `pulumi:"port"`
}

type HttpConfigResponseOutput struct{ *pulumi.OutputState }

func (HttpConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpConfigResponse)(nil)).Elem()
}

func (o HttpConfigResponseOutput) ToHttpConfigResponseOutput() HttpConfigResponseOutput {
	return o
}

func (o HttpConfigResponseOutput) ToHttpConfigResponseOutputWithContext(ctx context.Context) HttpConfigResponseOutput {
	return o
}

func (o HttpConfigResponseOutput) Hosts() HttpHostConfigResponseArrayOutput {
	return o.ApplyT(func(v HttpConfigResponse) []HttpHostConfigResponse { return v.Hosts }).(HttpHostConfigResponseArrayOutput)
}

func (o HttpConfigResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v HttpConfigResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o HttpConfigResponseOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v HttpConfigResponse) int { return v.Port }).(pulumi.IntOutput)
}

type HttpConfigResponseArrayOutput struct{ *pulumi.OutputState }

func (HttpConfigResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpConfigResponse)(nil)).Elem()
}

func (o HttpConfigResponseArrayOutput) ToHttpConfigResponseArrayOutput() HttpConfigResponseArrayOutput {
	return o
}

func (o HttpConfigResponseArrayOutput) ToHttpConfigResponseArrayOutputWithContext(ctx context.Context) HttpConfigResponseArrayOutput {
	return o
}

func (o HttpConfigResponseArrayOutput) Index(i pulumi.IntInput) HttpConfigResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HttpConfigResponse {
		return vs[0].([]HttpConfigResponse)[vs[1].(int)]
	}).(HttpConfigResponseOutput)
}

type HttpHostConfig struct {
	Name   string            `pulumi:"name"`
	Routes []HttpRouteConfig `pulumi:"routes"`
}





type HttpHostConfigInput interface {
	pulumi.Input

	ToHttpHostConfigOutput() HttpHostConfigOutput
	ToHttpHostConfigOutputWithContext(context.Context) HttpHostConfigOutput
}

type HttpHostConfigArgs struct {
	Name   pulumi.StringInput        `pulumi:"name"`
	Routes HttpRouteConfigArrayInput `pulumi:"routes"`
}

func (HttpHostConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpHostConfig)(nil)).Elem()
}

func (i HttpHostConfigArgs) ToHttpHostConfigOutput() HttpHostConfigOutput {
	return i.ToHttpHostConfigOutputWithContext(context.Background())
}

func (i HttpHostConfigArgs) ToHttpHostConfigOutputWithContext(ctx context.Context) HttpHostConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpHostConfigOutput)
}





type HttpHostConfigArrayInput interface {
	pulumi.Input

	ToHttpHostConfigArrayOutput() HttpHostConfigArrayOutput
	ToHttpHostConfigArrayOutputWithContext(context.Context) HttpHostConfigArrayOutput
}

type HttpHostConfigArray []HttpHostConfigInput

func (HttpHostConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpHostConfig)(nil)).Elem()
}

func (i HttpHostConfigArray) ToHttpHostConfigArrayOutput() HttpHostConfigArrayOutput {
	return i.ToHttpHostConfigArrayOutputWithContext(context.Background())
}

func (i HttpHostConfigArray) ToHttpHostConfigArrayOutputWithContext(ctx context.Context) HttpHostConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpHostConfigArrayOutput)
}

type HttpHostConfigOutput struct{ *pulumi.OutputState }

func (HttpHostConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpHostConfig)(nil)).Elem()
}

func (o HttpHostConfigOutput) ToHttpHostConfigOutput() HttpHostConfigOutput {
	return o
}

func (o HttpHostConfigOutput) ToHttpHostConfigOutputWithContext(ctx context.Context) HttpHostConfigOutput {
	return o
}

func (o HttpHostConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v HttpHostConfig) string { return v.Name }).(pulumi.StringOutput)
}

func (o HttpHostConfigOutput) Routes() HttpRouteConfigArrayOutput {
	return o.ApplyT(func(v HttpHostConfig) []HttpRouteConfig { return v.Routes }).(HttpRouteConfigArrayOutput)
}

type HttpHostConfigArrayOutput struct{ *pulumi.OutputState }

func (HttpHostConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpHostConfig)(nil)).Elem()
}

func (o HttpHostConfigArrayOutput) ToHttpHostConfigArrayOutput() HttpHostConfigArrayOutput {
	return o
}

func (o HttpHostConfigArrayOutput) ToHttpHostConfigArrayOutputWithContext(ctx context.Context) HttpHostConfigArrayOutput {
	return o
}

func (o HttpHostConfigArrayOutput) Index(i pulumi.IntInput) HttpHostConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HttpHostConfig {
		return vs[0].([]HttpHostConfig)[vs[1].(int)]
	}).(HttpHostConfigOutput)
}

type HttpHostConfigResponse struct {
	Name   string                    `pulumi:"name"`
	Routes []HttpRouteConfigResponse `pulumi:"routes"`
}

type HttpHostConfigResponseOutput struct{ *pulumi.OutputState }

func (HttpHostConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpHostConfigResponse)(nil)).Elem()
}

func (o HttpHostConfigResponseOutput) ToHttpHostConfigResponseOutput() HttpHostConfigResponseOutput {
	return o
}

func (o HttpHostConfigResponseOutput) ToHttpHostConfigResponseOutputWithContext(ctx context.Context) HttpHostConfigResponseOutput {
	return o
}

func (o HttpHostConfigResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v HttpHostConfigResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o HttpHostConfigResponseOutput) Routes() HttpRouteConfigResponseArrayOutput {
	return o.ApplyT(func(v HttpHostConfigResponse) []HttpRouteConfigResponse { return v.Routes }).(HttpRouteConfigResponseArrayOutput)
}

type HttpHostConfigResponseArrayOutput struct{ *pulumi.OutputState }

func (HttpHostConfigResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpHostConfigResponse)(nil)).Elem()
}

func (o HttpHostConfigResponseArrayOutput) ToHttpHostConfigResponseArrayOutput() HttpHostConfigResponseArrayOutput {
	return o
}

func (o HttpHostConfigResponseArrayOutput) ToHttpHostConfigResponseArrayOutputWithContext(ctx context.Context) HttpHostConfigResponseArrayOutput {
	return o
}

func (o HttpHostConfigResponseArrayOutput) Index(i pulumi.IntInput) HttpHostConfigResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HttpHostConfigResponse {
		return vs[0].([]HttpHostConfigResponse)[vs[1].(int)]
	}).(HttpHostConfigResponseOutput)
}

type HttpRouteConfig struct {
	Destination GatewayDestination `pulumi:"destination"`
	Match       HttpRouteMatchRule `pulumi:"match"`
	Name        string             `pulumi:"name"`
}





type HttpRouteConfigInput interface {
	pulumi.Input

	ToHttpRouteConfigOutput() HttpRouteConfigOutput
	ToHttpRouteConfigOutputWithContext(context.Context) HttpRouteConfigOutput
}

type HttpRouteConfigArgs struct {
	Destination GatewayDestinationInput `pulumi:"destination"`
	Match       HttpRouteMatchRuleInput `pulumi:"match"`
	Name        pulumi.StringInput      `pulumi:"name"`
}

func (HttpRouteConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRouteConfig)(nil)).Elem()
}

func (i HttpRouteConfigArgs) ToHttpRouteConfigOutput() HttpRouteConfigOutput {
	return i.ToHttpRouteConfigOutputWithContext(context.Background())
}

func (i HttpRouteConfigArgs) ToHttpRouteConfigOutputWithContext(ctx context.Context) HttpRouteConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpRouteConfigOutput)
}





type HttpRouteConfigArrayInput interface {
	pulumi.Input

	ToHttpRouteConfigArrayOutput() HttpRouteConfigArrayOutput
	ToHttpRouteConfigArrayOutputWithContext(context.Context) HttpRouteConfigArrayOutput
}

type HttpRouteConfigArray []HttpRouteConfigInput

func (HttpRouteConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpRouteConfig)(nil)).Elem()
}

func (i HttpRouteConfigArray) ToHttpRouteConfigArrayOutput() HttpRouteConfigArrayOutput {
	return i.ToHttpRouteConfigArrayOutputWithContext(context.Background())
}

func (i HttpRouteConfigArray) ToHttpRouteConfigArrayOutputWithContext(ctx context.Context) HttpRouteConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpRouteConfigArrayOutput)
}

type HttpRouteConfigOutput struct{ *pulumi.OutputState }

func (HttpRouteConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRouteConfig)(nil)).Elem()
}

func (o HttpRouteConfigOutput) ToHttpRouteConfigOutput() HttpRouteConfigOutput {
	return o
}

func (o HttpRouteConfigOutput) ToHttpRouteConfigOutputWithContext(ctx context.Context) HttpRouteConfigOutput {
	return o
}

func (o HttpRouteConfigOutput) Destination() GatewayDestinationOutput {
	return o.ApplyT(func(v HttpRouteConfig) GatewayDestination { return v.Destination }).(GatewayDestinationOutput)
}

func (o HttpRouteConfigOutput) Match() HttpRouteMatchRuleOutput {
	return o.ApplyT(func(v HttpRouteConfig) HttpRouteMatchRule { return v.Match }).(HttpRouteMatchRuleOutput)
}

func (o HttpRouteConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v HttpRouteConfig) string { return v.Name }).(pulumi.StringOutput)
}

type HttpRouteConfigArrayOutput struct{ *pulumi.OutputState }

func (HttpRouteConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpRouteConfig)(nil)).Elem()
}

func (o HttpRouteConfigArrayOutput) ToHttpRouteConfigArrayOutput() HttpRouteConfigArrayOutput {
	return o
}

func (o HttpRouteConfigArrayOutput) ToHttpRouteConfigArrayOutputWithContext(ctx context.Context) HttpRouteConfigArrayOutput {
	return o
}

func (o HttpRouteConfigArrayOutput) Index(i pulumi.IntInput) HttpRouteConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HttpRouteConfig {
		return vs[0].([]HttpRouteConfig)[vs[1].(int)]
	}).(HttpRouteConfigOutput)
}

type HttpRouteConfigResponse struct {
	Destination GatewayDestinationResponse `pulumi:"destination"`
	Match       HttpRouteMatchRuleResponse `pulumi:"match"`
	Name        string                     `pulumi:"name"`
}

type HttpRouteConfigResponseOutput struct{ *pulumi.OutputState }

func (HttpRouteConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRouteConfigResponse)(nil)).Elem()
}

func (o HttpRouteConfigResponseOutput) ToHttpRouteConfigResponseOutput() HttpRouteConfigResponseOutput {
	return o
}

func (o HttpRouteConfigResponseOutput) ToHttpRouteConfigResponseOutputWithContext(ctx context.Context) HttpRouteConfigResponseOutput {
	return o
}

func (o HttpRouteConfigResponseOutput) Destination() GatewayDestinationResponseOutput {
	return o.ApplyT(func(v HttpRouteConfigResponse) GatewayDestinationResponse { return v.Destination }).(GatewayDestinationResponseOutput)
}

func (o HttpRouteConfigResponseOutput) Match() HttpRouteMatchRuleResponseOutput {
	return o.ApplyT(func(v HttpRouteConfigResponse) HttpRouteMatchRuleResponse { return v.Match }).(HttpRouteMatchRuleResponseOutput)
}

func (o HttpRouteConfigResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v HttpRouteConfigResponse) string { return v.Name }).(pulumi.StringOutput)
}

type HttpRouteConfigResponseArrayOutput struct{ *pulumi.OutputState }

func (HttpRouteConfigResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpRouteConfigResponse)(nil)).Elem()
}

func (o HttpRouteConfigResponseArrayOutput) ToHttpRouteConfigResponseArrayOutput() HttpRouteConfigResponseArrayOutput {
	return o
}

func (o HttpRouteConfigResponseArrayOutput) ToHttpRouteConfigResponseArrayOutputWithContext(ctx context.Context) HttpRouteConfigResponseArrayOutput {
	return o
}

func (o HttpRouteConfigResponseArrayOutput) Index(i pulumi.IntInput) HttpRouteConfigResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HttpRouteConfigResponse {
		return vs[0].([]HttpRouteConfigResponse)[vs[1].(int)]
	}).(HttpRouteConfigResponseOutput)
}

type HttpRouteMatchHeader struct {
	Name  string  `pulumi:"name"`
	Type  *string `pulumi:"type"`
	Value *string `pulumi:"value"`
}





type HttpRouteMatchHeaderInput interface {
	pulumi.Input

	ToHttpRouteMatchHeaderOutput() HttpRouteMatchHeaderOutput
	ToHttpRouteMatchHeaderOutputWithContext(context.Context) HttpRouteMatchHeaderOutput
}

type HttpRouteMatchHeaderArgs struct {
	Name  pulumi.StringInput    `pulumi:"name"`
	Type  pulumi.StringPtrInput `pulumi:"type"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (HttpRouteMatchHeaderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRouteMatchHeader)(nil)).Elem()
}

func (i HttpRouteMatchHeaderArgs) ToHttpRouteMatchHeaderOutput() HttpRouteMatchHeaderOutput {
	return i.ToHttpRouteMatchHeaderOutputWithContext(context.Background())
}

func (i HttpRouteMatchHeaderArgs) ToHttpRouteMatchHeaderOutputWithContext(ctx context.Context) HttpRouteMatchHeaderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpRouteMatchHeaderOutput)
}





type HttpRouteMatchHeaderArrayInput interface {
	pulumi.Input

	ToHttpRouteMatchHeaderArrayOutput() HttpRouteMatchHeaderArrayOutput
	ToHttpRouteMatchHeaderArrayOutputWithContext(context.Context) HttpRouteMatchHeaderArrayOutput
}

type HttpRouteMatchHeaderArray []HttpRouteMatchHeaderInput

func (HttpRouteMatchHeaderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpRouteMatchHeader)(nil)).Elem()
}

func (i HttpRouteMatchHeaderArray) ToHttpRouteMatchHeaderArrayOutput() HttpRouteMatchHeaderArrayOutput {
	return i.ToHttpRouteMatchHeaderArrayOutputWithContext(context.Background())
}

func (i HttpRouteMatchHeaderArray) ToHttpRouteMatchHeaderArrayOutputWithContext(ctx context.Context) HttpRouteMatchHeaderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpRouteMatchHeaderArrayOutput)
}

type HttpRouteMatchHeaderOutput struct{ *pulumi.OutputState }

func (HttpRouteMatchHeaderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRouteMatchHeader)(nil)).Elem()
}

func (o HttpRouteMatchHeaderOutput) ToHttpRouteMatchHeaderOutput() HttpRouteMatchHeaderOutput {
	return o
}

func (o HttpRouteMatchHeaderOutput) ToHttpRouteMatchHeaderOutputWithContext(ctx context.Context) HttpRouteMatchHeaderOutput {
	return o
}

func (o HttpRouteMatchHeaderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v HttpRouteMatchHeader) string { return v.Name }).(pulumi.StringOutput)
}

func (o HttpRouteMatchHeaderOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpRouteMatchHeader) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o HttpRouteMatchHeaderOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpRouteMatchHeader) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type HttpRouteMatchHeaderArrayOutput struct{ *pulumi.OutputState }

func (HttpRouteMatchHeaderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpRouteMatchHeader)(nil)).Elem()
}

func (o HttpRouteMatchHeaderArrayOutput) ToHttpRouteMatchHeaderArrayOutput() HttpRouteMatchHeaderArrayOutput {
	return o
}

func (o HttpRouteMatchHeaderArrayOutput) ToHttpRouteMatchHeaderArrayOutputWithContext(ctx context.Context) HttpRouteMatchHeaderArrayOutput {
	return o
}

func (o HttpRouteMatchHeaderArrayOutput) Index(i pulumi.IntInput) HttpRouteMatchHeaderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HttpRouteMatchHeader {
		return vs[0].([]HttpRouteMatchHeader)[vs[1].(int)]
	}).(HttpRouteMatchHeaderOutput)
}

type HttpRouteMatchHeaderResponse struct {
	Name  string  `pulumi:"name"`
	Type  *string `pulumi:"type"`
	Value *string `pulumi:"value"`
}

type HttpRouteMatchHeaderResponseOutput struct{ *pulumi.OutputState }

func (HttpRouteMatchHeaderResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRouteMatchHeaderResponse)(nil)).Elem()
}

func (o HttpRouteMatchHeaderResponseOutput) ToHttpRouteMatchHeaderResponseOutput() HttpRouteMatchHeaderResponseOutput {
	return o
}

func (o HttpRouteMatchHeaderResponseOutput) ToHttpRouteMatchHeaderResponseOutputWithContext(ctx context.Context) HttpRouteMatchHeaderResponseOutput {
	return o
}

func (o HttpRouteMatchHeaderResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v HttpRouteMatchHeaderResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o HttpRouteMatchHeaderResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpRouteMatchHeaderResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o HttpRouteMatchHeaderResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpRouteMatchHeaderResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type HttpRouteMatchHeaderResponseArrayOutput struct{ *pulumi.OutputState }

func (HttpRouteMatchHeaderResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HttpRouteMatchHeaderResponse)(nil)).Elem()
}

func (o HttpRouteMatchHeaderResponseArrayOutput) ToHttpRouteMatchHeaderResponseArrayOutput() HttpRouteMatchHeaderResponseArrayOutput {
	return o
}

func (o HttpRouteMatchHeaderResponseArrayOutput) ToHttpRouteMatchHeaderResponseArrayOutputWithContext(ctx context.Context) HttpRouteMatchHeaderResponseArrayOutput {
	return o
}

func (o HttpRouteMatchHeaderResponseArrayOutput) Index(i pulumi.IntInput) HttpRouteMatchHeaderResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HttpRouteMatchHeaderResponse {
		return vs[0].([]HttpRouteMatchHeaderResponse)[vs[1].(int)]
	}).(HttpRouteMatchHeaderResponseOutput)
}

type HttpRouteMatchPath struct {
	Rewrite *string `pulumi:"rewrite"`
	Type    string  `pulumi:"type"`
	Value   string  `pulumi:"value"`
}





type HttpRouteMatchPathInput interface {
	pulumi.Input

	ToHttpRouteMatchPathOutput() HttpRouteMatchPathOutput
	ToHttpRouteMatchPathOutputWithContext(context.Context) HttpRouteMatchPathOutput
}

type HttpRouteMatchPathArgs struct {
	Rewrite pulumi.StringPtrInput `pulumi:"rewrite"`
	Type    pulumi.StringInput    `pulumi:"type"`
	Value   pulumi.StringInput    `pulumi:"value"`
}

func (HttpRouteMatchPathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRouteMatchPath)(nil)).Elem()
}

func (i HttpRouteMatchPathArgs) ToHttpRouteMatchPathOutput() HttpRouteMatchPathOutput {
	return i.ToHttpRouteMatchPathOutputWithContext(context.Background())
}

func (i HttpRouteMatchPathArgs) ToHttpRouteMatchPathOutputWithContext(ctx context.Context) HttpRouteMatchPathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpRouteMatchPathOutput)
}

type HttpRouteMatchPathOutput struct{ *pulumi.OutputState }

func (HttpRouteMatchPathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRouteMatchPath)(nil)).Elem()
}

func (o HttpRouteMatchPathOutput) ToHttpRouteMatchPathOutput() HttpRouteMatchPathOutput {
	return o
}

func (o HttpRouteMatchPathOutput) ToHttpRouteMatchPathOutputWithContext(ctx context.Context) HttpRouteMatchPathOutput {
	return o
}

func (o HttpRouteMatchPathOutput) Rewrite() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpRouteMatchPath) *string { return v.Rewrite }).(pulumi.StringPtrOutput)
}

func (o HttpRouteMatchPathOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v HttpRouteMatchPath) string { return v.Type }).(pulumi.StringOutput)
}

func (o HttpRouteMatchPathOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v HttpRouteMatchPath) string { return v.Value }).(pulumi.StringOutput)
}

type HttpRouteMatchPathResponse struct {
	Rewrite *string `pulumi:"rewrite"`
	Type    string  `pulumi:"type"`
	Value   string  `pulumi:"value"`
}

type HttpRouteMatchPathResponseOutput struct{ *pulumi.OutputState }

func (HttpRouteMatchPathResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRouteMatchPathResponse)(nil)).Elem()
}

func (o HttpRouteMatchPathResponseOutput) ToHttpRouteMatchPathResponseOutput() HttpRouteMatchPathResponseOutput {
	return o
}

func (o HttpRouteMatchPathResponseOutput) ToHttpRouteMatchPathResponseOutputWithContext(ctx context.Context) HttpRouteMatchPathResponseOutput {
	return o
}

func (o HttpRouteMatchPathResponseOutput) Rewrite() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpRouteMatchPathResponse) *string { return v.Rewrite }).(pulumi.StringPtrOutput)
}

func (o HttpRouteMatchPathResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v HttpRouteMatchPathResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o HttpRouteMatchPathResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v HttpRouteMatchPathResponse) string { return v.Value }).(pulumi.StringOutput)
}

type HttpRouteMatchRule struct {
	Headers []HttpRouteMatchHeader `pulumi:"headers"`
	Path    HttpRouteMatchPath     `pulumi:"path"`
}





type HttpRouteMatchRuleInput interface {
	pulumi.Input

	ToHttpRouteMatchRuleOutput() HttpRouteMatchRuleOutput
	ToHttpRouteMatchRuleOutputWithContext(context.Context) HttpRouteMatchRuleOutput
}

type HttpRouteMatchRuleArgs struct {
	Headers HttpRouteMatchHeaderArrayInput `pulumi:"headers"`
	Path    HttpRouteMatchPathInput        `pulumi:"path"`
}

func (HttpRouteMatchRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRouteMatchRule)(nil)).Elem()
}

func (i HttpRouteMatchRuleArgs) ToHttpRouteMatchRuleOutput() HttpRouteMatchRuleOutput {
	return i.ToHttpRouteMatchRuleOutputWithContext(context.Background())
}

func (i HttpRouteMatchRuleArgs) ToHttpRouteMatchRuleOutputWithContext(ctx context.Context) HttpRouteMatchRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpRouteMatchRuleOutput)
}

type HttpRouteMatchRuleOutput struct{ *pulumi.OutputState }

func (HttpRouteMatchRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRouteMatchRule)(nil)).Elem()
}

func (o HttpRouteMatchRuleOutput) ToHttpRouteMatchRuleOutput() HttpRouteMatchRuleOutput {
	return o
}

func (o HttpRouteMatchRuleOutput) ToHttpRouteMatchRuleOutputWithContext(ctx context.Context) HttpRouteMatchRuleOutput {
	return o
}

func (o HttpRouteMatchRuleOutput) Headers() HttpRouteMatchHeaderArrayOutput {
	return o.ApplyT(func(v HttpRouteMatchRule) []HttpRouteMatchHeader { return v.Headers }).(HttpRouteMatchHeaderArrayOutput)
}

func (o HttpRouteMatchRuleOutput) Path() HttpRouteMatchPathOutput {
	return o.ApplyT(func(v HttpRouteMatchRule) HttpRouteMatchPath { return v.Path }).(HttpRouteMatchPathOutput)
}

type HttpRouteMatchRuleResponse struct {
	Headers []HttpRouteMatchHeaderResponse `pulumi:"headers"`
	Path    HttpRouteMatchPathResponse     `pulumi:"path"`
}

type HttpRouteMatchRuleResponseOutput struct{ *pulumi.OutputState }

func (HttpRouteMatchRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpRouteMatchRuleResponse)(nil)).Elem()
}

func (o HttpRouteMatchRuleResponseOutput) ToHttpRouteMatchRuleResponseOutput() HttpRouteMatchRuleResponseOutput {
	return o
}

func (o HttpRouteMatchRuleResponseOutput) ToHttpRouteMatchRuleResponseOutputWithContext(ctx context.Context) HttpRouteMatchRuleResponseOutput {
	return o
}

func (o HttpRouteMatchRuleResponseOutput) Headers() HttpRouteMatchHeaderResponseArrayOutput {
	return o.ApplyT(func(v HttpRouteMatchRuleResponse) []HttpRouteMatchHeaderResponse { return v.Headers }).(HttpRouteMatchHeaderResponseArrayOutput)
}

func (o HttpRouteMatchRuleResponseOutput) Path() HttpRouteMatchPathResponseOutput {
	return o.ApplyT(func(v HttpRouteMatchRuleResponse) HttpRouteMatchPathResponse { return v.Path }).(HttpRouteMatchPathResponseOutput)
}

type ImageRegistryCredential struct {
	Password *string `pulumi:"password"`
	Server   string  `pulumi:"server"`
	Username string  `pulumi:"username"`
}





type ImageRegistryCredentialInput interface {
	pulumi.Input

	ToImageRegistryCredentialOutput() ImageRegistryCredentialOutput
	ToImageRegistryCredentialOutputWithContext(context.Context) ImageRegistryCredentialOutput
}

type ImageRegistryCredentialArgs struct {
	Password pulumi.StringPtrInput `pulumi:"password"`
	Server   pulumi.StringInput    `pulumi:"server"`
	Username pulumi.StringInput    `pulumi:"username"`
}

func (ImageRegistryCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRegistryCredential)(nil)).Elem()
}

func (i ImageRegistryCredentialArgs) ToImageRegistryCredentialOutput() ImageRegistryCredentialOutput {
	return i.ToImageRegistryCredentialOutputWithContext(context.Background())
}

func (i ImageRegistryCredentialArgs) ToImageRegistryCredentialOutputWithContext(ctx context.Context) ImageRegistryCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRegistryCredentialOutput)
}

func (i ImageRegistryCredentialArgs) ToImageRegistryCredentialPtrOutput() ImageRegistryCredentialPtrOutput {
	return i.ToImageRegistryCredentialPtrOutputWithContext(context.Background())
}

func (i ImageRegistryCredentialArgs) ToImageRegistryCredentialPtrOutputWithContext(ctx context.Context) ImageRegistryCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRegistryCredentialOutput).ToImageRegistryCredentialPtrOutputWithContext(ctx)
}









type ImageRegistryCredentialPtrInput interface {
	pulumi.Input

	ToImageRegistryCredentialPtrOutput() ImageRegistryCredentialPtrOutput
	ToImageRegistryCredentialPtrOutputWithContext(context.Context) ImageRegistryCredentialPtrOutput
}

type imageRegistryCredentialPtrType ImageRegistryCredentialArgs

func ImageRegistryCredentialPtr(v *ImageRegistryCredentialArgs) ImageRegistryCredentialPtrInput {
	return (*imageRegistryCredentialPtrType)(v)
}

func (*imageRegistryCredentialPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRegistryCredential)(nil)).Elem()
}

func (i *imageRegistryCredentialPtrType) ToImageRegistryCredentialPtrOutput() ImageRegistryCredentialPtrOutput {
	return i.ToImageRegistryCredentialPtrOutputWithContext(context.Background())
}

func (i *imageRegistryCredentialPtrType) ToImageRegistryCredentialPtrOutputWithContext(ctx context.Context) ImageRegistryCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRegistryCredentialPtrOutput)
}

type ImageRegistryCredentialOutput struct{ *pulumi.OutputState }

func (ImageRegistryCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRegistryCredential)(nil)).Elem()
}

func (o ImageRegistryCredentialOutput) ToImageRegistryCredentialOutput() ImageRegistryCredentialOutput {
	return o
}

func (o ImageRegistryCredentialOutput) ToImageRegistryCredentialOutputWithContext(ctx context.Context) ImageRegistryCredentialOutput {
	return o
}

func (o ImageRegistryCredentialOutput) ToImageRegistryCredentialPtrOutput() ImageRegistryCredentialPtrOutput {
	return o.ToImageRegistryCredentialPtrOutputWithContext(context.Background())
}

func (o ImageRegistryCredentialOutput) ToImageRegistryCredentialPtrOutputWithContext(ctx context.Context) ImageRegistryCredentialPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageRegistryCredential) *ImageRegistryCredential {
		return &v
	}).(ImageRegistryCredentialPtrOutput)
}

func (o ImageRegistryCredentialOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageRegistryCredential) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v ImageRegistryCredential) string { return v.Server }).(pulumi.StringOutput)
}

func (o ImageRegistryCredentialOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v ImageRegistryCredential) string { return v.Username }).(pulumi.StringOutput)
}

type ImageRegistryCredentialPtrOutput struct{ *pulumi.OutputState }

func (ImageRegistryCredentialPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRegistryCredential)(nil)).Elem()
}

func (o ImageRegistryCredentialPtrOutput) ToImageRegistryCredentialPtrOutput() ImageRegistryCredentialPtrOutput {
	return o
}

func (o ImageRegistryCredentialPtrOutput) ToImageRegistryCredentialPtrOutputWithContext(ctx context.Context) ImageRegistryCredentialPtrOutput {
	return o
}

func (o ImageRegistryCredentialPtrOutput) Elem() ImageRegistryCredentialOutput {
	return o.ApplyT(func(v *ImageRegistryCredential) ImageRegistryCredential {
		if v != nil {
			return *v
		}
		var ret ImageRegistryCredential
		return ret
	}).(ImageRegistryCredentialOutput)
}

func (o ImageRegistryCredentialPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRegistryCredential) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialPtrOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRegistryCredential) *string {
		if v == nil {
			return nil
		}
		return &v.Server
	}).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRegistryCredential) *string {
		if v == nil {
			return nil
		}
		return &v.Username
	}).(pulumi.StringPtrOutput)
}

type ImageRegistryCredentialResponse struct {
	Password *string `pulumi:"password"`
	Server   string  `pulumi:"server"`
	Username string  `pulumi:"username"`
}

type ImageRegistryCredentialResponseOutput struct{ *pulumi.OutputState }

func (ImageRegistryCredentialResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRegistryCredentialResponse)(nil)).Elem()
}

func (o ImageRegistryCredentialResponseOutput) ToImageRegistryCredentialResponseOutput() ImageRegistryCredentialResponseOutput {
	return o
}

func (o ImageRegistryCredentialResponseOutput) ToImageRegistryCredentialResponseOutputWithContext(ctx context.Context) ImageRegistryCredentialResponseOutput {
	return o
}

func (o ImageRegistryCredentialResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageRegistryCredentialResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialResponseOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v ImageRegistryCredentialResponse) string { return v.Server }).(pulumi.StringOutput)
}

func (o ImageRegistryCredentialResponseOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v ImageRegistryCredentialResponse) string { return v.Username }).(pulumi.StringOutput)
}

type ImageRegistryCredentialResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageRegistryCredentialResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRegistryCredentialResponse)(nil)).Elem()
}

func (o ImageRegistryCredentialResponsePtrOutput) ToImageRegistryCredentialResponsePtrOutput() ImageRegistryCredentialResponsePtrOutput {
	return o
}

func (o ImageRegistryCredentialResponsePtrOutput) ToImageRegistryCredentialResponsePtrOutputWithContext(ctx context.Context) ImageRegistryCredentialResponsePtrOutput {
	return o
}

func (o ImageRegistryCredentialResponsePtrOutput) Elem() ImageRegistryCredentialResponseOutput {
	return o.ApplyT(func(v *ImageRegistryCredentialResponse) ImageRegistryCredentialResponse {
		if v != nil {
			return *v
		}
		var ret ImageRegistryCredentialResponse
		return ret
	}).(ImageRegistryCredentialResponseOutput)
}

func (o ImageRegistryCredentialResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRegistryCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialResponsePtrOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRegistryCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Server
	}).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRegistryCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Username
	}).(pulumi.StringPtrOutput)
}

type NetworkRef struct {
	EndpointRefs []EndpointRef `pulumi:"endpointRefs"`
	Name         *string       `pulumi:"name"`
}





type NetworkRefInput interface {
	pulumi.Input

	ToNetworkRefOutput() NetworkRefOutput
	ToNetworkRefOutputWithContext(context.Context) NetworkRefOutput
}

type NetworkRefArgs struct {
	EndpointRefs EndpointRefArrayInput `pulumi:"endpointRefs"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
}

func (NetworkRefArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRef)(nil)).Elem()
}

func (i NetworkRefArgs) ToNetworkRefOutput() NetworkRefOutput {
	return i.ToNetworkRefOutputWithContext(context.Background())
}

func (i NetworkRefArgs) ToNetworkRefOutputWithContext(ctx context.Context) NetworkRefOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRefOutput)
}





type NetworkRefArrayInput interface {
	pulumi.Input

	ToNetworkRefArrayOutput() NetworkRefArrayOutput
	ToNetworkRefArrayOutputWithContext(context.Context) NetworkRefArrayOutput
}

type NetworkRefArray []NetworkRefInput

func (NetworkRefArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkRef)(nil)).Elem()
}

func (i NetworkRefArray) ToNetworkRefArrayOutput() NetworkRefArrayOutput {
	return i.ToNetworkRefArrayOutputWithContext(context.Background())
}

func (i NetworkRefArray) ToNetworkRefArrayOutputWithContext(ctx context.Context) NetworkRefArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRefArrayOutput)
}

type NetworkRefOutput struct{ *pulumi.OutputState }

func (NetworkRefOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRef)(nil)).Elem()
}

func (o NetworkRefOutput) ToNetworkRefOutput() NetworkRefOutput {
	return o
}

func (o NetworkRefOutput) ToNetworkRefOutputWithContext(ctx context.Context) NetworkRefOutput {
	return o
}

func (o NetworkRefOutput) EndpointRefs() EndpointRefArrayOutput {
	return o.ApplyT(func(v NetworkRef) []EndpointRef { return v.EndpointRefs }).(EndpointRefArrayOutput)
}

func (o NetworkRefOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRef) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type NetworkRefArrayOutput struct{ *pulumi.OutputState }

func (NetworkRefArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkRef)(nil)).Elem()
}

func (o NetworkRefArrayOutput) ToNetworkRefArrayOutput() NetworkRefArrayOutput {
	return o
}

func (o NetworkRefArrayOutput) ToNetworkRefArrayOutputWithContext(ctx context.Context) NetworkRefArrayOutput {
	return o
}

func (o NetworkRefArrayOutput) Index(i pulumi.IntInput) NetworkRefOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkRef {
		return vs[0].([]NetworkRef)[vs[1].(int)]
	}).(NetworkRefOutput)
}

type NetworkRefResponse struct {
	EndpointRefs []EndpointRefResponse `pulumi:"endpointRefs"`
	Name         *string               `pulumi:"name"`
}

type NetworkRefResponseOutput struct{ *pulumi.OutputState }

func (NetworkRefResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRefResponse)(nil)).Elem()
}

func (o NetworkRefResponseOutput) ToNetworkRefResponseOutput() NetworkRefResponseOutput {
	return o
}

func (o NetworkRefResponseOutput) ToNetworkRefResponseOutputWithContext(ctx context.Context) NetworkRefResponseOutput {
	return o
}

func (o NetworkRefResponseOutput) EndpointRefs() EndpointRefResponseArrayOutput {
	return o.ApplyT(func(v NetworkRefResponse) []EndpointRefResponse { return v.EndpointRefs }).(EndpointRefResponseArrayOutput)
}

func (o NetworkRefResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRefResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type NetworkRefResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkRefResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkRefResponse)(nil)).Elem()
}

func (o NetworkRefResponseArrayOutput) ToNetworkRefResponseArrayOutput() NetworkRefResponseArrayOutput {
	return o
}

func (o NetworkRefResponseArrayOutput) ToNetworkRefResponseArrayOutputWithContext(ctx context.Context) NetworkRefResponseArrayOutput {
	return o
}

func (o NetworkRefResponseArrayOutput) Index(i pulumi.IntInput) NetworkRefResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkRefResponse {
		return vs[0].([]NetworkRefResponse)[vs[1].(int)]
	}).(NetworkRefResponseOutput)
}

type NetworkResourceProperties struct {
	Description *string `pulumi:"description"`
	Kind        string  `pulumi:"kind"`
}





type NetworkResourcePropertiesInput interface {
	pulumi.Input

	ToNetworkResourcePropertiesOutput() NetworkResourcePropertiesOutput
	ToNetworkResourcePropertiesOutputWithContext(context.Context) NetworkResourcePropertiesOutput
}

type NetworkResourcePropertiesArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Kind        pulumi.StringInput    `pulumi:"kind"`
}

func (NetworkResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkResourceProperties)(nil)).Elem()
}

func (i NetworkResourcePropertiesArgs) ToNetworkResourcePropertiesOutput() NetworkResourcePropertiesOutput {
	return i.ToNetworkResourcePropertiesOutputWithContext(context.Background())
}

func (i NetworkResourcePropertiesArgs) ToNetworkResourcePropertiesOutputWithContext(ctx context.Context) NetworkResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkResourcePropertiesOutput)
}

type NetworkResourcePropertiesOutput struct{ *pulumi.OutputState }

func (NetworkResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkResourceProperties)(nil)).Elem()
}

func (o NetworkResourcePropertiesOutput) ToNetworkResourcePropertiesOutput() NetworkResourcePropertiesOutput {
	return o
}

func (o NetworkResourcePropertiesOutput) ToNetworkResourcePropertiesOutputWithContext(ctx context.Context) NetworkResourcePropertiesOutput {
	return o
}

func (o NetworkResourcePropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkResourceProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkResourcePropertiesOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkResourceProperties) string { return v.Kind }).(pulumi.StringOutput)
}

type NetworkResourcePropertiesResponse struct {
	Description       *string `pulumi:"description"`
	Kind              string  `pulumi:"kind"`
	ProvisioningState string  `pulumi:"provisioningState"`
	Status            string  `pulumi:"status"`
	StatusDetails     string  `pulumi:"statusDetails"`
}

type NetworkResourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (NetworkResourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkResourcePropertiesResponse)(nil)).Elem()
}

func (o NetworkResourcePropertiesResponseOutput) ToNetworkResourcePropertiesResponseOutput() NetworkResourcePropertiesResponseOutput {
	return o
}

func (o NetworkResourcePropertiesResponseOutput) ToNetworkResourcePropertiesResponseOutputWithContext(ctx context.Context) NetworkResourcePropertiesResponseOutput {
	return o
}

func (o NetworkResourcePropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkResourcePropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkResourcePropertiesResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkResourcePropertiesResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o NetworkResourcePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkResourcePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NetworkResourcePropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkResourcePropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o NetworkResourcePropertiesResponseOutput) StatusDetails() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkResourcePropertiesResponse) string { return v.StatusDetails }).(pulumi.StringOutput)
}

type ReliableCollectionsRef struct {
	DoNotPersistState *bool  `pulumi:"doNotPersistState"`
	Name              string `pulumi:"name"`
}





type ReliableCollectionsRefInput interface {
	pulumi.Input

	ToReliableCollectionsRefOutput() ReliableCollectionsRefOutput
	ToReliableCollectionsRefOutputWithContext(context.Context) ReliableCollectionsRefOutput
}

type ReliableCollectionsRefArgs struct {
	DoNotPersistState pulumi.BoolPtrInput `pulumi:"doNotPersistState"`
	Name              pulumi.StringInput  `pulumi:"name"`
}

func (ReliableCollectionsRefArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReliableCollectionsRef)(nil)).Elem()
}

func (i ReliableCollectionsRefArgs) ToReliableCollectionsRefOutput() ReliableCollectionsRefOutput {
	return i.ToReliableCollectionsRefOutputWithContext(context.Background())
}

func (i ReliableCollectionsRefArgs) ToReliableCollectionsRefOutputWithContext(ctx context.Context) ReliableCollectionsRefOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReliableCollectionsRefOutput)
}





type ReliableCollectionsRefArrayInput interface {
	pulumi.Input

	ToReliableCollectionsRefArrayOutput() ReliableCollectionsRefArrayOutput
	ToReliableCollectionsRefArrayOutputWithContext(context.Context) ReliableCollectionsRefArrayOutput
}

type ReliableCollectionsRefArray []ReliableCollectionsRefInput

func (ReliableCollectionsRefArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReliableCollectionsRef)(nil)).Elem()
}

func (i ReliableCollectionsRefArray) ToReliableCollectionsRefArrayOutput() ReliableCollectionsRefArrayOutput {
	return i.ToReliableCollectionsRefArrayOutputWithContext(context.Background())
}

func (i ReliableCollectionsRefArray) ToReliableCollectionsRefArrayOutputWithContext(ctx context.Context) ReliableCollectionsRefArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReliableCollectionsRefArrayOutput)
}

type ReliableCollectionsRefOutput struct{ *pulumi.OutputState }

func (ReliableCollectionsRefOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReliableCollectionsRef)(nil)).Elem()
}

func (o ReliableCollectionsRefOutput) ToReliableCollectionsRefOutput() ReliableCollectionsRefOutput {
	return o
}

func (o ReliableCollectionsRefOutput) ToReliableCollectionsRefOutputWithContext(ctx context.Context) ReliableCollectionsRefOutput {
	return o
}

func (o ReliableCollectionsRefOutput) DoNotPersistState() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ReliableCollectionsRef) *bool { return v.DoNotPersistState }).(pulumi.BoolPtrOutput)
}

func (o ReliableCollectionsRefOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReliableCollectionsRef) string { return v.Name }).(pulumi.StringOutput)
}

type ReliableCollectionsRefArrayOutput struct{ *pulumi.OutputState }

func (ReliableCollectionsRefArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReliableCollectionsRef)(nil)).Elem()
}

func (o ReliableCollectionsRefArrayOutput) ToReliableCollectionsRefArrayOutput() ReliableCollectionsRefArrayOutput {
	return o
}

func (o ReliableCollectionsRefArrayOutput) ToReliableCollectionsRefArrayOutputWithContext(ctx context.Context) ReliableCollectionsRefArrayOutput {
	return o
}

func (o ReliableCollectionsRefArrayOutput) Index(i pulumi.IntInput) ReliableCollectionsRefOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReliableCollectionsRef {
		return vs[0].([]ReliableCollectionsRef)[vs[1].(int)]
	}).(ReliableCollectionsRefOutput)
}

type ReliableCollectionsRefResponse struct {
	DoNotPersistState *bool  `pulumi:"doNotPersistState"`
	Name              string `pulumi:"name"`
}

type ReliableCollectionsRefResponseOutput struct{ *pulumi.OutputState }

func (ReliableCollectionsRefResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReliableCollectionsRefResponse)(nil)).Elem()
}

func (o ReliableCollectionsRefResponseOutput) ToReliableCollectionsRefResponseOutput() ReliableCollectionsRefResponseOutput {
	return o
}

func (o ReliableCollectionsRefResponseOutput) ToReliableCollectionsRefResponseOutputWithContext(ctx context.Context) ReliableCollectionsRefResponseOutput {
	return o
}

func (o ReliableCollectionsRefResponseOutput) DoNotPersistState() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ReliableCollectionsRefResponse) *bool { return v.DoNotPersistState }).(pulumi.BoolPtrOutput)
}

func (o ReliableCollectionsRefResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReliableCollectionsRefResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ReliableCollectionsRefResponseArrayOutput struct{ *pulumi.OutputState }

func (ReliableCollectionsRefResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReliableCollectionsRefResponse)(nil)).Elem()
}

func (o ReliableCollectionsRefResponseArrayOutput) ToReliableCollectionsRefResponseArrayOutput() ReliableCollectionsRefResponseArrayOutput {
	return o
}

func (o ReliableCollectionsRefResponseArrayOutput) ToReliableCollectionsRefResponseArrayOutputWithContext(ctx context.Context) ReliableCollectionsRefResponseArrayOutput {
	return o
}

func (o ReliableCollectionsRefResponseArrayOutput) Index(i pulumi.IntInput) ReliableCollectionsRefResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReliableCollectionsRefResponse {
		return vs[0].([]ReliableCollectionsRefResponse)[vs[1].(int)]
	}).(ReliableCollectionsRefResponseOutput)
}

type ResourceLimits struct {
	Cpu        *float64 `pulumi:"cpu"`
	MemoryInGB *float64 `pulumi:"memoryInGB"`
}





type ResourceLimitsInput interface {
	pulumi.Input

	ToResourceLimitsOutput() ResourceLimitsOutput
	ToResourceLimitsOutputWithContext(context.Context) ResourceLimitsOutput
}

type ResourceLimitsArgs struct {
	Cpu        pulumi.Float64PtrInput `pulumi:"cpu"`
	MemoryInGB pulumi.Float64PtrInput `pulumi:"memoryInGB"`
}

func (ResourceLimitsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLimits)(nil)).Elem()
}

func (i ResourceLimitsArgs) ToResourceLimitsOutput() ResourceLimitsOutput {
	return i.ToResourceLimitsOutputWithContext(context.Background())
}

func (i ResourceLimitsArgs) ToResourceLimitsOutputWithContext(ctx context.Context) ResourceLimitsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLimitsOutput)
}

func (i ResourceLimitsArgs) ToResourceLimitsPtrOutput() ResourceLimitsPtrOutput {
	return i.ToResourceLimitsPtrOutputWithContext(context.Background())
}

func (i ResourceLimitsArgs) ToResourceLimitsPtrOutputWithContext(ctx context.Context) ResourceLimitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLimitsOutput).ToResourceLimitsPtrOutputWithContext(ctx)
}









type ResourceLimitsPtrInput interface {
	pulumi.Input

	ToResourceLimitsPtrOutput() ResourceLimitsPtrOutput
	ToResourceLimitsPtrOutputWithContext(context.Context) ResourceLimitsPtrOutput
}

type resourceLimitsPtrType ResourceLimitsArgs

func ResourceLimitsPtr(v *ResourceLimitsArgs) ResourceLimitsPtrInput {
	return (*resourceLimitsPtrType)(v)
}

func (*resourceLimitsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceLimits)(nil)).Elem()
}

func (i *resourceLimitsPtrType) ToResourceLimitsPtrOutput() ResourceLimitsPtrOutput {
	return i.ToResourceLimitsPtrOutputWithContext(context.Background())
}

func (i *resourceLimitsPtrType) ToResourceLimitsPtrOutputWithContext(ctx context.Context) ResourceLimitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLimitsPtrOutput)
}

type ResourceLimitsOutput struct{ *pulumi.OutputState }

func (ResourceLimitsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLimits)(nil)).Elem()
}

func (o ResourceLimitsOutput) ToResourceLimitsOutput() ResourceLimitsOutput {
	return o
}

func (o ResourceLimitsOutput) ToResourceLimitsOutputWithContext(ctx context.Context) ResourceLimitsOutput {
	return o
}

func (o ResourceLimitsOutput) ToResourceLimitsPtrOutput() ResourceLimitsPtrOutput {
	return o.ToResourceLimitsPtrOutputWithContext(context.Background())
}

func (o ResourceLimitsOutput) ToResourceLimitsPtrOutputWithContext(ctx context.Context) ResourceLimitsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceLimits) *ResourceLimits {
		return &v
	}).(ResourceLimitsPtrOutput)
}

func (o ResourceLimitsOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ResourceLimits) *float64 { return v.Cpu }).(pulumi.Float64PtrOutput)
}

func (o ResourceLimitsOutput) MemoryInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ResourceLimits) *float64 { return v.MemoryInGB }).(pulumi.Float64PtrOutput)
}

type ResourceLimitsPtrOutput struct{ *pulumi.OutputState }

func (ResourceLimitsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceLimits)(nil)).Elem()
}

func (o ResourceLimitsPtrOutput) ToResourceLimitsPtrOutput() ResourceLimitsPtrOutput {
	return o
}

func (o ResourceLimitsPtrOutput) ToResourceLimitsPtrOutputWithContext(ctx context.Context) ResourceLimitsPtrOutput {
	return o
}

func (o ResourceLimitsPtrOutput) Elem() ResourceLimitsOutput {
	return o.ApplyT(func(v *ResourceLimits) ResourceLimits {
		if v != nil {
			return *v
		}
		var ret ResourceLimits
		return ret
	}).(ResourceLimitsOutput)
}

func (o ResourceLimitsPtrOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ResourceLimits) *float64 {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.Float64PtrOutput)
}

func (o ResourceLimitsPtrOutput) MemoryInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ResourceLimits) *float64 {
		if v == nil {
			return nil
		}
		return v.MemoryInGB
	}).(pulumi.Float64PtrOutput)
}

type ResourceLimitsResponse struct {
	Cpu        *float64 `pulumi:"cpu"`
	MemoryInGB *float64 `pulumi:"memoryInGB"`
}

type ResourceLimitsResponseOutput struct{ *pulumi.OutputState }

func (ResourceLimitsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLimitsResponse)(nil)).Elem()
}

func (o ResourceLimitsResponseOutput) ToResourceLimitsResponseOutput() ResourceLimitsResponseOutput {
	return o
}

func (o ResourceLimitsResponseOutput) ToResourceLimitsResponseOutputWithContext(ctx context.Context) ResourceLimitsResponseOutput {
	return o
}

func (o ResourceLimitsResponseOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ResourceLimitsResponse) *float64 { return v.Cpu }).(pulumi.Float64PtrOutput)
}

func (o ResourceLimitsResponseOutput) MemoryInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ResourceLimitsResponse) *float64 { return v.MemoryInGB }).(pulumi.Float64PtrOutput)
}

type ResourceLimitsResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceLimitsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceLimitsResponse)(nil)).Elem()
}

func (o ResourceLimitsResponsePtrOutput) ToResourceLimitsResponsePtrOutput() ResourceLimitsResponsePtrOutput {
	return o
}

func (o ResourceLimitsResponsePtrOutput) ToResourceLimitsResponsePtrOutputWithContext(ctx context.Context) ResourceLimitsResponsePtrOutput {
	return o
}

func (o ResourceLimitsResponsePtrOutput) Elem() ResourceLimitsResponseOutput {
	return o.ApplyT(func(v *ResourceLimitsResponse) ResourceLimitsResponse {
		if v != nil {
			return *v
		}
		var ret ResourceLimitsResponse
		return ret
	}).(ResourceLimitsResponseOutput)
}

func (o ResourceLimitsResponsePtrOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ResourceLimitsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.Float64PtrOutput)
}

func (o ResourceLimitsResponsePtrOutput) MemoryInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ResourceLimitsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MemoryInGB
	}).(pulumi.Float64PtrOutput)
}

type ResourceRequests struct {
	Cpu        float64 `pulumi:"cpu"`
	MemoryInGB float64 `pulumi:"memoryInGB"`
}





type ResourceRequestsInput interface {
	pulumi.Input

	ToResourceRequestsOutput() ResourceRequestsOutput
	ToResourceRequestsOutputWithContext(context.Context) ResourceRequestsOutput
}

type ResourceRequestsArgs struct {
	Cpu        pulumi.Float64Input `pulumi:"cpu"`
	MemoryInGB pulumi.Float64Input `pulumi:"memoryInGB"`
}

func (ResourceRequestsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequests)(nil)).Elem()
}

func (i ResourceRequestsArgs) ToResourceRequestsOutput() ResourceRequestsOutput {
	return i.ToResourceRequestsOutputWithContext(context.Background())
}

func (i ResourceRequestsArgs) ToResourceRequestsOutputWithContext(ctx context.Context) ResourceRequestsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRequestsOutput)
}

type ResourceRequestsOutput struct{ *pulumi.OutputState }

func (ResourceRequestsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequests)(nil)).Elem()
}

func (o ResourceRequestsOutput) ToResourceRequestsOutput() ResourceRequestsOutput {
	return o
}

func (o ResourceRequestsOutput) ToResourceRequestsOutputWithContext(ctx context.Context) ResourceRequestsOutput {
	return o
}

func (o ResourceRequestsOutput) Cpu() pulumi.Float64Output {
	return o.ApplyT(func(v ResourceRequests) float64 { return v.Cpu }).(pulumi.Float64Output)
}

func (o ResourceRequestsOutput) MemoryInGB() pulumi.Float64Output {
	return o.ApplyT(func(v ResourceRequests) float64 { return v.MemoryInGB }).(pulumi.Float64Output)
}

type ResourceRequestsResponse struct {
	Cpu        float64 `pulumi:"cpu"`
	MemoryInGB float64 `pulumi:"memoryInGB"`
}

type ResourceRequestsResponseOutput struct{ *pulumi.OutputState }

func (ResourceRequestsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequestsResponse)(nil)).Elem()
}

func (o ResourceRequestsResponseOutput) ToResourceRequestsResponseOutput() ResourceRequestsResponseOutput {
	return o
}

func (o ResourceRequestsResponseOutput) ToResourceRequestsResponseOutputWithContext(ctx context.Context) ResourceRequestsResponseOutput {
	return o
}

func (o ResourceRequestsResponseOutput) Cpu() pulumi.Float64Output {
	return o.ApplyT(func(v ResourceRequestsResponse) float64 { return v.Cpu }).(pulumi.Float64Output)
}

func (o ResourceRequestsResponseOutput) MemoryInGB() pulumi.Float64Output {
	return o.ApplyT(func(v ResourceRequestsResponse) float64 { return v.MemoryInGB }).(pulumi.Float64Output)
}

type ResourceRequirements struct {
	Limits   *ResourceLimits  `pulumi:"limits"`
	Requests ResourceRequests `pulumi:"requests"`
}





type ResourceRequirementsInput interface {
	pulumi.Input

	ToResourceRequirementsOutput() ResourceRequirementsOutput
	ToResourceRequirementsOutputWithContext(context.Context) ResourceRequirementsOutput
}

type ResourceRequirementsArgs struct {
	Limits   ResourceLimitsPtrInput `pulumi:"limits"`
	Requests ResourceRequestsInput  `pulumi:"requests"`
}

func (ResourceRequirementsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequirements)(nil)).Elem()
}

func (i ResourceRequirementsArgs) ToResourceRequirementsOutput() ResourceRequirementsOutput {
	return i.ToResourceRequirementsOutputWithContext(context.Background())
}

func (i ResourceRequirementsArgs) ToResourceRequirementsOutputWithContext(ctx context.Context) ResourceRequirementsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRequirementsOutput)
}

type ResourceRequirementsOutput struct{ *pulumi.OutputState }

func (ResourceRequirementsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequirements)(nil)).Elem()
}

func (o ResourceRequirementsOutput) ToResourceRequirementsOutput() ResourceRequirementsOutput {
	return o
}

func (o ResourceRequirementsOutput) ToResourceRequirementsOutputWithContext(ctx context.Context) ResourceRequirementsOutput {
	return o
}

func (o ResourceRequirementsOutput) Limits() ResourceLimitsPtrOutput {
	return o.ApplyT(func(v ResourceRequirements) *ResourceLimits { return v.Limits }).(ResourceLimitsPtrOutput)
}

func (o ResourceRequirementsOutput) Requests() ResourceRequestsOutput {
	return o.ApplyT(func(v ResourceRequirements) ResourceRequests { return v.Requests }).(ResourceRequestsOutput)
}

type ResourceRequirementsResponse struct {
	Limits   *ResourceLimitsResponse  `pulumi:"limits"`
	Requests ResourceRequestsResponse `pulumi:"requests"`
}

type ResourceRequirementsResponseOutput struct{ *pulumi.OutputState }

func (ResourceRequirementsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequirementsResponse)(nil)).Elem()
}

func (o ResourceRequirementsResponseOutput) ToResourceRequirementsResponseOutput() ResourceRequirementsResponseOutput {
	return o
}

func (o ResourceRequirementsResponseOutput) ToResourceRequirementsResponseOutputWithContext(ctx context.Context) ResourceRequirementsResponseOutput {
	return o
}

func (o ResourceRequirementsResponseOutput) Limits() ResourceLimitsResponsePtrOutput {
	return o.ApplyT(func(v ResourceRequirementsResponse) *ResourceLimitsResponse { return v.Limits }).(ResourceLimitsResponsePtrOutput)
}

func (o ResourceRequirementsResponseOutput) Requests() ResourceRequestsResponseOutput {
	return o.ApplyT(func(v ResourceRequirementsResponse) ResourceRequestsResponse { return v.Requests }).(ResourceRequestsResponseOutput)
}

type SecretResourceProperties struct {
	ContentType *string `pulumi:"contentType"`
	Description *string `pulumi:"description"`
	Kind        string  `pulumi:"kind"`
}





type SecretResourcePropertiesInput interface {
	pulumi.Input

	ToSecretResourcePropertiesOutput() SecretResourcePropertiesOutput
	ToSecretResourcePropertiesOutputWithContext(context.Context) SecretResourcePropertiesOutput
}

type SecretResourcePropertiesArgs struct {
	ContentType pulumi.StringPtrInput `pulumi:"contentType"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	Kind        pulumi.StringInput    `pulumi:"kind"`
}

func (SecretResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretResourceProperties)(nil)).Elem()
}

func (i SecretResourcePropertiesArgs) ToSecretResourcePropertiesOutput() SecretResourcePropertiesOutput {
	return i.ToSecretResourcePropertiesOutputWithContext(context.Background())
}

func (i SecretResourcePropertiesArgs) ToSecretResourcePropertiesOutputWithContext(ctx context.Context) SecretResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretResourcePropertiesOutput)
}

type SecretResourcePropertiesOutput struct{ *pulumi.OutputState }

func (SecretResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretResourceProperties)(nil)).Elem()
}

func (o SecretResourcePropertiesOutput) ToSecretResourcePropertiesOutput() SecretResourcePropertiesOutput {
	return o
}

func (o SecretResourcePropertiesOutput) ToSecretResourcePropertiesOutputWithContext(ctx context.Context) SecretResourcePropertiesOutput {
	return o
}

func (o SecretResourcePropertiesOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretResourceProperties) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o SecretResourcePropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretResourceProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecretResourcePropertiesOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v SecretResourceProperties) string { return v.Kind }).(pulumi.StringOutput)
}

type SecretResourcePropertiesResponse struct {
	ContentType       *string `pulumi:"contentType"`
	Description       *string `pulumi:"description"`
	Kind              string  `pulumi:"kind"`
	ProvisioningState string  `pulumi:"provisioningState"`
	Status            string  `pulumi:"status"`
	StatusDetails     string  `pulumi:"statusDetails"`
}

type SecretResourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (SecretResourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretResourcePropertiesResponse)(nil)).Elem()
}

func (o SecretResourcePropertiesResponseOutput) ToSecretResourcePropertiesResponseOutput() SecretResourcePropertiesResponseOutput {
	return o
}

func (o SecretResourcePropertiesResponseOutput) ToSecretResourcePropertiesResponseOutputWithContext(ctx context.Context) SecretResourcePropertiesResponseOutput {
	return o
}

func (o SecretResourcePropertiesResponseOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretResourcePropertiesResponse) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o SecretResourcePropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretResourcePropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecretResourcePropertiesResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v SecretResourcePropertiesResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o SecretResourcePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v SecretResourcePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SecretResourcePropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v SecretResourcePropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o SecretResourcePropertiesResponseOutput) StatusDetails() pulumi.StringOutput {
	return o.ApplyT(func(v SecretResourcePropertiesResponse) string { return v.StatusDetails }).(pulumi.StringOutput)
}

type ServiceResourceDescription struct {
	AutoScalingPolicies []AutoScalingPolicy              `pulumi:"autoScalingPolicies"`
	CodePackages        []ContainerCodePackageProperties `pulumi:"codePackages"`
	Description         *string                          `pulumi:"description"`
	Diagnostics         *DiagnosticsRef                  `pulumi:"diagnostics"`
	Name                *string                          `pulumi:"name"`
	NetworkRefs         []NetworkRef                     `pulumi:"networkRefs"`
	OsType              string                           `pulumi:"osType"`
	ReplicaCount        *int                             `pulumi:"replicaCount"`
}





type ServiceResourceDescriptionInput interface {
	pulumi.Input

	ToServiceResourceDescriptionOutput() ServiceResourceDescriptionOutput
	ToServiceResourceDescriptionOutputWithContext(context.Context) ServiceResourceDescriptionOutput
}

type ServiceResourceDescriptionArgs struct {
	AutoScalingPolicies AutoScalingPolicyArrayInput              `pulumi:"autoScalingPolicies"`
	CodePackages        ContainerCodePackagePropertiesArrayInput `pulumi:"codePackages"`
	Description         pulumi.StringPtrInput                    `pulumi:"description"`
	Diagnostics         DiagnosticsRefPtrInput                   `pulumi:"diagnostics"`
	Name                pulumi.StringPtrInput                    `pulumi:"name"`
	NetworkRefs         NetworkRefArrayInput                     `pulumi:"networkRefs"`
	OsType              pulumi.StringInput                       `pulumi:"osType"`
	ReplicaCount        pulumi.IntPtrInput                       `pulumi:"replicaCount"`
}

func (ServiceResourceDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceResourceDescription)(nil)).Elem()
}

func (i ServiceResourceDescriptionArgs) ToServiceResourceDescriptionOutput() ServiceResourceDescriptionOutput {
	return i.ToServiceResourceDescriptionOutputWithContext(context.Background())
}

func (i ServiceResourceDescriptionArgs) ToServiceResourceDescriptionOutputWithContext(ctx context.Context) ServiceResourceDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceResourceDescriptionOutput)
}





type ServiceResourceDescriptionArrayInput interface {
	pulumi.Input

	ToServiceResourceDescriptionArrayOutput() ServiceResourceDescriptionArrayOutput
	ToServiceResourceDescriptionArrayOutputWithContext(context.Context) ServiceResourceDescriptionArrayOutput
}

type ServiceResourceDescriptionArray []ServiceResourceDescriptionInput

func (ServiceResourceDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceResourceDescription)(nil)).Elem()
}

func (i ServiceResourceDescriptionArray) ToServiceResourceDescriptionArrayOutput() ServiceResourceDescriptionArrayOutput {
	return i.ToServiceResourceDescriptionArrayOutputWithContext(context.Background())
}

func (i ServiceResourceDescriptionArray) ToServiceResourceDescriptionArrayOutputWithContext(ctx context.Context) ServiceResourceDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceResourceDescriptionArrayOutput)
}

type ServiceResourceDescriptionOutput struct{ *pulumi.OutputState }

func (ServiceResourceDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceResourceDescription)(nil)).Elem()
}

func (o ServiceResourceDescriptionOutput) ToServiceResourceDescriptionOutput() ServiceResourceDescriptionOutput {
	return o
}

func (o ServiceResourceDescriptionOutput) ToServiceResourceDescriptionOutputWithContext(ctx context.Context) ServiceResourceDescriptionOutput {
	return o
}

func (o ServiceResourceDescriptionOutput) AutoScalingPolicies() AutoScalingPolicyArrayOutput {
	return o.ApplyT(func(v ServiceResourceDescription) []AutoScalingPolicy { return v.AutoScalingPolicies }).(AutoScalingPolicyArrayOutput)
}

func (o ServiceResourceDescriptionOutput) CodePackages() ContainerCodePackagePropertiesArrayOutput {
	return o.ApplyT(func(v ServiceResourceDescription) []ContainerCodePackageProperties { return v.CodePackages }).(ContainerCodePackagePropertiesArrayOutput)
}

func (o ServiceResourceDescriptionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceResourceDescription) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ServiceResourceDescriptionOutput) Diagnostics() DiagnosticsRefPtrOutput {
	return o.ApplyT(func(v ServiceResourceDescription) *DiagnosticsRef { return v.Diagnostics }).(DiagnosticsRefPtrOutput)
}

func (o ServiceResourceDescriptionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceResourceDescription) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServiceResourceDescriptionOutput) NetworkRefs() NetworkRefArrayOutput {
	return o.ApplyT(func(v ServiceResourceDescription) []NetworkRef { return v.NetworkRefs }).(NetworkRefArrayOutput)
}

func (o ServiceResourceDescriptionOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResourceDescription) string { return v.OsType }).(pulumi.StringOutput)
}

func (o ServiceResourceDescriptionOutput) ReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceResourceDescription) *int { return v.ReplicaCount }).(pulumi.IntPtrOutput)
}

type ServiceResourceDescriptionArrayOutput struct{ *pulumi.OutputState }

func (ServiceResourceDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceResourceDescription)(nil)).Elem()
}

func (o ServiceResourceDescriptionArrayOutput) ToServiceResourceDescriptionArrayOutput() ServiceResourceDescriptionArrayOutput {
	return o
}

func (o ServiceResourceDescriptionArrayOutput) ToServiceResourceDescriptionArrayOutputWithContext(ctx context.Context) ServiceResourceDescriptionArrayOutput {
	return o
}

func (o ServiceResourceDescriptionArrayOutput) Index(i pulumi.IntInput) ServiceResourceDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceResourceDescription {
		return vs[0].([]ServiceResourceDescription)[vs[1].(int)]
	}).(ServiceResourceDescriptionOutput)
}

type ServiceResourceDescriptionResponse struct {
	AutoScalingPolicies []AutoScalingPolicyResponse              `pulumi:"autoScalingPolicies"`
	CodePackages        []ContainerCodePackagePropertiesResponse `pulumi:"codePackages"`
	Description         *string                                  `pulumi:"description"`
	Diagnostics         *DiagnosticsRefResponse                  `pulumi:"diagnostics"`
	HealthState         string                                   `pulumi:"healthState"`
	Id                  string                                   `pulumi:"id"`
	Name                *string                                  `pulumi:"name"`
	NetworkRefs         []NetworkRefResponse                     `pulumi:"networkRefs"`
	OsType              string                                   `pulumi:"osType"`
	ProvisioningState   string                                   `pulumi:"provisioningState"`
	ReplicaCount        *int                                     `pulumi:"replicaCount"`
	Status              string                                   `pulumi:"status"`
	StatusDetails       string                                   `pulumi:"statusDetails"`
	Type                string                                   `pulumi:"type"`
	UnhealthyEvaluation string                                   `pulumi:"unhealthyEvaluation"`
}

type ServiceResourceDescriptionResponseOutput struct{ *pulumi.OutputState }

func (ServiceResourceDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceResourceDescriptionResponse)(nil)).Elem()
}

func (o ServiceResourceDescriptionResponseOutput) ToServiceResourceDescriptionResponseOutput() ServiceResourceDescriptionResponseOutput {
	return o
}

func (o ServiceResourceDescriptionResponseOutput) ToServiceResourceDescriptionResponseOutputWithContext(ctx context.Context) ServiceResourceDescriptionResponseOutput {
	return o
}

func (o ServiceResourceDescriptionResponseOutput) AutoScalingPolicies() AutoScalingPolicyResponseArrayOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) []AutoScalingPolicyResponse { return v.AutoScalingPolicies }).(AutoScalingPolicyResponseArrayOutput)
}

func (o ServiceResourceDescriptionResponseOutput) CodePackages() ContainerCodePackagePropertiesResponseArrayOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) []ContainerCodePackagePropertiesResponse {
		return v.CodePackages
	}).(ContainerCodePackagePropertiesResponseArrayOutput)
}

func (o ServiceResourceDescriptionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ServiceResourceDescriptionResponseOutput) Diagnostics() DiagnosticsRefResponsePtrOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) *DiagnosticsRefResponse { return v.Diagnostics }).(DiagnosticsRefResponsePtrOutput)
}

func (o ServiceResourceDescriptionResponseOutput) HealthState() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) string { return v.HealthState }).(pulumi.StringOutput)
}

func (o ServiceResourceDescriptionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ServiceResourceDescriptionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServiceResourceDescriptionResponseOutput) NetworkRefs() NetworkRefResponseArrayOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) []NetworkRefResponse { return v.NetworkRefs }).(NetworkRefResponseArrayOutput)
}

func (o ServiceResourceDescriptionResponseOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) string { return v.OsType }).(pulumi.StringOutput)
}

func (o ServiceResourceDescriptionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServiceResourceDescriptionResponseOutput) ReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) *int { return v.ReplicaCount }).(pulumi.IntPtrOutput)
}

func (o ServiceResourceDescriptionResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o ServiceResourceDescriptionResponseOutput) StatusDetails() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) string { return v.StatusDetails }).(pulumi.StringOutput)
}

func (o ServiceResourceDescriptionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ServiceResourceDescriptionResponseOutput) UnhealthyEvaluation() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) string { return v.UnhealthyEvaluation }).(pulumi.StringOutput)
}

type ServiceResourceDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceResourceDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceResourceDescriptionResponse)(nil)).Elem()
}

func (o ServiceResourceDescriptionResponseArrayOutput) ToServiceResourceDescriptionResponseArrayOutput() ServiceResourceDescriptionResponseArrayOutput {
	return o
}

func (o ServiceResourceDescriptionResponseArrayOutput) ToServiceResourceDescriptionResponseArrayOutputWithContext(ctx context.Context) ServiceResourceDescriptionResponseArrayOutput {
	return o
}

func (o ServiceResourceDescriptionResponseArrayOutput) Index(i pulumi.IntInput) ServiceResourceDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceResourceDescriptionResponse {
		return vs[0].([]ServiceResourceDescriptionResponse)[vs[1].(int)]
	}).(ServiceResourceDescriptionResponseOutput)
}

type Setting struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type SettingInput interface {
	pulumi.Input

	ToSettingOutput() SettingOutput
	ToSettingOutputWithContext(context.Context) SettingOutput
}

type SettingArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (SettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Setting)(nil)).Elem()
}

func (i SettingArgs) ToSettingOutput() SettingOutput {
	return i.ToSettingOutputWithContext(context.Background())
}

func (i SettingArgs) ToSettingOutputWithContext(ctx context.Context) SettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingOutput)
}





type SettingArrayInput interface {
	pulumi.Input

	ToSettingArrayOutput() SettingArrayOutput
	ToSettingArrayOutputWithContext(context.Context) SettingArrayOutput
}

type SettingArray []SettingInput

func (SettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Setting)(nil)).Elem()
}

func (i SettingArray) ToSettingArrayOutput() SettingArrayOutput {
	return i.ToSettingArrayOutputWithContext(context.Background())
}

func (i SettingArray) ToSettingArrayOutputWithContext(ctx context.Context) SettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingArrayOutput)
}

type SettingOutput struct{ *pulumi.OutputState }

func (SettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Setting)(nil)).Elem()
}

func (o SettingOutput) ToSettingOutput() SettingOutput {
	return o
}

func (o SettingOutput) ToSettingOutputWithContext(ctx context.Context) SettingOutput {
	return o
}

func (o SettingOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Setting) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SettingOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Setting) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type SettingArrayOutput struct{ *pulumi.OutputState }

func (SettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Setting)(nil)).Elem()
}

func (o SettingArrayOutput) ToSettingArrayOutput() SettingArrayOutput {
	return o
}

func (o SettingArrayOutput) ToSettingArrayOutputWithContext(ctx context.Context) SettingArrayOutput {
	return o
}

func (o SettingArrayOutput) Index(i pulumi.IntInput) SettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Setting {
		return vs[0].([]Setting)[vs[1].(int)]
	}).(SettingOutput)
}

type SettingResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type SettingResponseOutput struct{ *pulumi.OutputState }

func (SettingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingResponse)(nil)).Elem()
}

func (o SettingResponseOutput) ToSettingResponseOutput() SettingResponseOutput {
	return o
}

func (o SettingResponseOutput) ToSettingResponseOutputWithContext(ctx context.Context) SettingResponseOutput {
	return o
}

func (o SettingResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SettingResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SettingResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SettingResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type SettingResponseArrayOutput struct{ *pulumi.OutputState }

func (SettingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingResponse)(nil)).Elem()
}

func (o SettingResponseArrayOutput) ToSettingResponseArrayOutput() SettingResponseArrayOutput {
	return o
}

func (o SettingResponseArrayOutput) ToSettingResponseArrayOutputWithContext(ctx context.Context) SettingResponseArrayOutput {
	return o
}

func (o SettingResponseArrayOutput) Index(i pulumi.IntInput) SettingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingResponse {
		return vs[0].([]SettingResponse)[vs[1].(int)]
	}).(SettingResponseOutput)
}

type TcpConfig struct {
	Destination GatewayDestination `pulumi:"destination"`
	Name        string             `pulumi:"name"`
	Port        int                `pulumi:"port"`
}





type TcpConfigInput interface {
	pulumi.Input

	ToTcpConfigOutput() TcpConfigOutput
	ToTcpConfigOutputWithContext(context.Context) TcpConfigOutput
}

type TcpConfigArgs struct {
	Destination GatewayDestinationInput `pulumi:"destination"`
	Name        pulumi.StringInput      `pulumi:"name"`
	Port        pulumi.IntInput         `pulumi:"port"`
}

func (TcpConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TcpConfig)(nil)).Elem()
}

func (i TcpConfigArgs) ToTcpConfigOutput() TcpConfigOutput {
	return i.ToTcpConfigOutputWithContext(context.Background())
}

func (i TcpConfigArgs) ToTcpConfigOutputWithContext(ctx context.Context) TcpConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TcpConfigOutput)
}





type TcpConfigArrayInput interface {
	pulumi.Input

	ToTcpConfigArrayOutput() TcpConfigArrayOutput
	ToTcpConfigArrayOutputWithContext(context.Context) TcpConfigArrayOutput
}

type TcpConfigArray []TcpConfigInput

func (TcpConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TcpConfig)(nil)).Elem()
}

func (i TcpConfigArray) ToTcpConfigArrayOutput() TcpConfigArrayOutput {
	return i.ToTcpConfigArrayOutputWithContext(context.Background())
}

func (i TcpConfigArray) ToTcpConfigArrayOutputWithContext(ctx context.Context) TcpConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TcpConfigArrayOutput)
}

type TcpConfigOutput struct{ *pulumi.OutputState }

func (TcpConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TcpConfig)(nil)).Elem()
}

func (o TcpConfigOutput) ToTcpConfigOutput() TcpConfigOutput {
	return o
}

func (o TcpConfigOutput) ToTcpConfigOutputWithContext(ctx context.Context) TcpConfigOutput {
	return o
}

func (o TcpConfigOutput) Destination() GatewayDestinationOutput {
	return o.ApplyT(func(v TcpConfig) GatewayDestination { return v.Destination }).(GatewayDestinationOutput)
}

func (o TcpConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TcpConfig) string { return v.Name }).(pulumi.StringOutput)
}

func (o TcpConfigOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v TcpConfig) int { return v.Port }).(pulumi.IntOutput)
}

type TcpConfigArrayOutput struct{ *pulumi.OutputState }

func (TcpConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TcpConfig)(nil)).Elem()
}

func (o TcpConfigArrayOutput) ToTcpConfigArrayOutput() TcpConfigArrayOutput {
	return o
}

func (o TcpConfigArrayOutput) ToTcpConfigArrayOutputWithContext(ctx context.Context) TcpConfigArrayOutput {
	return o
}

func (o TcpConfigArrayOutput) Index(i pulumi.IntInput) TcpConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TcpConfig {
		return vs[0].([]TcpConfig)[vs[1].(int)]
	}).(TcpConfigOutput)
}

type TcpConfigResponse struct {
	Destination GatewayDestinationResponse `pulumi:"destination"`
	Name        string                     `pulumi:"name"`
	Port        int                        `pulumi:"port"`
}

type TcpConfigResponseOutput struct{ *pulumi.OutputState }

func (TcpConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TcpConfigResponse)(nil)).Elem()
}

func (o TcpConfigResponseOutput) ToTcpConfigResponseOutput() TcpConfigResponseOutput {
	return o
}

func (o TcpConfigResponseOutput) ToTcpConfigResponseOutputWithContext(ctx context.Context) TcpConfigResponseOutput {
	return o
}

func (o TcpConfigResponseOutput) Destination() GatewayDestinationResponseOutput {
	return o.ApplyT(func(v TcpConfigResponse) GatewayDestinationResponse { return v.Destination }).(GatewayDestinationResponseOutput)
}

func (o TcpConfigResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TcpConfigResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o TcpConfigResponseOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v TcpConfigResponse) int { return v.Port }).(pulumi.IntOutput)
}

type TcpConfigResponseArrayOutput struct{ *pulumi.OutputState }

func (TcpConfigResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TcpConfigResponse)(nil)).Elem()
}

func (o TcpConfigResponseArrayOutput) ToTcpConfigResponseArrayOutput() TcpConfigResponseArrayOutput {
	return o
}

func (o TcpConfigResponseArrayOutput) ToTcpConfigResponseArrayOutputWithContext(ctx context.Context) TcpConfigResponseArrayOutput {
	return o
}

func (o TcpConfigResponseArrayOutput) Index(i pulumi.IntInput) TcpConfigResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TcpConfigResponse {
		return vs[0].([]TcpConfigResponse)[vs[1].(int)]
	}).(TcpConfigResponseOutput)
}

type VolumeProviderParametersAzureFile struct {
	AccountKey  *string `pulumi:"accountKey"`
	AccountName string  `pulumi:"accountName"`
	ShareName   string  `pulumi:"shareName"`
}





type VolumeProviderParametersAzureFileInput interface {
	pulumi.Input

	ToVolumeProviderParametersAzureFileOutput() VolumeProviderParametersAzureFileOutput
	ToVolumeProviderParametersAzureFileOutputWithContext(context.Context) VolumeProviderParametersAzureFileOutput
}

type VolumeProviderParametersAzureFileArgs struct {
	AccountKey  pulumi.StringPtrInput `pulumi:"accountKey"`
	AccountName pulumi.StringInput    `pulumi:"accountName"`
	ShareName   pulumi.StringInput    `pulumi:"shareName"`
}

func (VolumeProviderParametersAzureFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeProviderParametersAzureFile)(nil)).Elem()
}

func (i VolumeProviderParametersAzureFileArgs) ToVolumeProviderParametersAzureFileOutput() VolumeProviderParametersAzureFileOutput {
	return i.ToVolumeProviderParametersAzureFileOutputWithContext(context.Background())
}

func (i VolumeProviderParametersAzureFileArgs) ToVolumeProviderParametersAzureFileOutputWithContext(ctx context.Context) VolumeProviderParametersAzureFileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeProviderParametersAzureFileOutput)
}

func (i VolumeProviderParametersAzureFileArgs) ToVolumeProviderParametersAzureFilePtrOutput() VolumeProviderParametersAzureFilePtrOutput {
	return i.ToVolumeProviderParametersAzureFilePtrOutputWithContext(context.Background())
}

func (i VolumeProviderParametersAzureFileArgs) ToVolumeProviderParametersAzureFilePtrOutputWithContext(ctx context.Context) VolumeProviderParametersAzureFilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeProviderParametersAzureFileOutput).ToVolumeProviderParametersAzureFilePtrOutputWithContext(ctx)
}









type VolumeProviderParametersAzureFilePtrInput interface {
	pulumi.Input

	ToVolumeProviderParametersAzureFilePtrOutput() VolumeProviderParametersAzureFilePtrOutput
	ToVolumeProviderParametersAzureFilePtrOutputWithContext(context.Context) VolumeProviderParametersAzureFilePtrOutput
}

type volumeProviderParametersAzureFilePtrType VolumeProviderParametersAzureFileArgs

func VolumeProviderParametersAzureFilePtr(v *VolumeProviderParametersAzureFileArgs) VolumeProviderParametersAzureFilePtrInput {
	return (*volumeProviderParametersAzureFilePtrType)(v)
}

func (*volumeProviderParametersAzureFilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeProviderParametersAzureFile)(nil)).Elem()
}

func (i *volumeProviderParametersAzureFilePtrType) ToVolumeProviderParametersAzureFilePtrOutput() VolumeProviderParametersAzureFilePtrOutput {
	return i.ToVolumeProviderParametersAzureFilePtrOutputWithContext(context.Background())
}

func (i *volumeProviderParametersAzureFilePtrType) ToVolumeProviderParametersAzureFilePtrOutputWithContext(ctx context.Context) VolumeProviderParametersAzureFilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeProviderParametersAzureFilePtrOutput)
}

type VolumeProviderParametersAzureFileOutput struct{ *pulumi.OutputState }

func (VolumeProviderParametersAzureFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeProviderParametersAzureFile)(nil)).Elem()
}

func (o VolumeProviderParametersAzureFileOutput) ToVolumeProviderParametersAzureFileOutput() VolumeProviderParametersAzureFileOutput {
	return o
}

func (o VolumeProviderParametersAzureFileOutput) ToVolumeProviderParametersAzureFileOutputWithContext(ctx context.Context) VolumeProviderParametersAzureFileOutput {
	return o
}

func (o VolumeProviderParametersAzureFileOutput) ToVolumeProviderParametersAzureFilePtrOutput() VolumeProviderParametersAzureFilePtrOutput {
	return o.ToVolumeProviderParametersAzureFilePtrOutputWithContext(context.Background())
}

func (o VolumeProviderParametersAzureFileOutput) ToVolumeProviderParametersAzureFilePtrOutputWithContext(ctx context.Context) VolumeProviderParametersAzureFilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumeProviderParametersAzureFile) *VolumeProviderParametersAzureFile {
		return &v
	}).(VolumeProviderParametersAzureFilePtrOutput)
}

func (o VolumeProviderParametersAzureFileOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeProviderParametersAzureFile) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o VolumeProviderParametersAzureFileOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeProviderParametersAzureFile) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o VolumeProviderParametersAzureFileOutput) ShareName() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeProviderParametersAzureFile) string { return v.ShareName }).(pulumi.StringOutput)
}

type VolumeProviderParametersAzureFilePtrOutput struct{ *pulumi.OutputState }

func (VolumeProviderParametersAzureFilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeProviderParametersAzureFile)(nil)).Elem()
}

func (o VolumeProviderParametersAzureFilePtrOutput) ToVolumeProviderParametersAzureFilePtrOutput() VolumeProviderParametersAzureFilePtrOutput {
	return o
}

func (o VolumeProviderParametersAzureFilePtrOutput) ToVolumeProviderParametersAzureFilePtrOutputWithContext(ctx context.Context) VolumeProviderParametersAzureFilePtrOutput {
	return o
}

func (o VolumeProviderParametersAzureFilePtrOutput) Elem() VolumeProviderParametersAzureFileOutput {
	return o.ApplyT(func(v *VolumeProviderParametersAzureFile) VolumeProviderParametersAzureFile {
		if v != nil {
			return *v
		}
		var ret VolumeProviderParametersAzureFile
		return ret
	}).(VolumeProviderParametersAzureFileOutput)
}

func (o VolumeProviderParametersAzureFilePtrOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeProviderParametersAzureFile) *string {
		if v == nil {
			return nil
		}
		return v.AccountKey
	}).(pulumi.StringPtrOutput)
}

func (o VolumeProviderParametersAzureFilePtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeProviderParametersAzureFile) *string {
		if v == nil {
			return nil
		}
		return &v.AccountName
	}).(pulumi.StringPtrOutput)
}

func (o VolumeProviderParametersAzureFilePtrOutput) ShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeProviderParametersAzureFile) *string {
		if v == nil {
			return nil
		}
		return &v.ShareName
	}).(pulumi.StringPtrOutput)
}

type VolumeProviderParametersAzureFileResponse struct {
	AccountKey  *string `pulumi:"accountKey"`
	AccountName string  `pulumi:"accountName"`
	ShareName   string  `pulumi:"shareName"`
}

type VolumeProviderParametersAzureFileResponseOutput struct{ *pulumi.OutputState }

func (VolumeProviderParametersAzureFileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeProviderParametersAzureFileResponse)(nil)).Elem()
}

func (o VolumeProviderParametersAzureFileResponseOutput) ToVolumeProviderParametersAzureFileResponseOutput() VolumeProviderParametersAzureFileResponseOutput {
	return o
}

func (o VolumeProviderParametersAzureFileResponseOutput) ToVolumeProviderParametersAzureFileResponseOutputWithContext(ctx context.Context) VolumeProviderParametersAzureFileResponseOutput {
	return o
}

func (o VolumeProviderParametersAzureFileResponseOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeProviderParametersAzureFileResponse) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o VolumeProviderParametersAzureFileResponseOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeProviderParametersAzureFileResponse) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o VolumeProviderParametersAzureFileResponseOutput) ShareName() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeProviderParametersAzureFileResponse) string { return v.ShareName }).(pulumi.StringOutput)
}

type VolumeProviderParametersAzureFileResponsePtrOutput struct{ *pulumi.OutputState }

func (VolumeProviderParametersAzureFileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeProviderParametersAzureFileResponse)(nil)).Elem()
}

func (o VolumeProviderParametersAzureFileResponsePtrOutput) ToVolumeProviderParametersAzureFileResponsePtrOutput() VolumeProviderParametersAzureFileResponsePtrOutput {
	return o
}

func (o VolumeProviderParametersAzureFileResponsePtrOutput) ToVolumeProviderParametersAzureFileResponsePtrOutputWithContext(ctx context.Context) VolumeProviderParametersAzureFileResponsePtrOutput {
	return o
}

func (o VolumeProviderParametersAzureFileResponsePtrOutput) Elem() VolumeProviderParametersAzureFileResponseOutput {
	return o.ApplyT(func(v *VolumeProviderParametersAzureFileResponse) VolumeProviderParametersAzureFileResponse {
		if v != nil {
			return *v
		}
		var ret VolumeProviderParametersAzureFileResponse
		return ret
	}).(VolumeProviderParametersAzureFileResponseOutput)
}

func (o VolumeProviderParametersAzureFileResponsePtrOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeProviderParametersAzureFileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountKey
	}).(pulumi.StringPtrOutput)
}

func (o VolumeProviderParametersAzureFileResponsePtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeProviderParametersAzureFileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AccountName
	}).(pulumi.StringPtrOutput)
}

func (o VolumeProviderParametersAzureFileResponsePtrOutput) ShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeProviderParametersAzureFileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ShareName
	}).(pulumi.StringPtrOutput)
}

type VolumeReference struct {
	DestinationPath string `pulumi:"destinationPath"`
	Name            string `pulumi:"name"`
	ReadOnly        *bool  `pulumi:"readOnly"`
}





type VolumeReferenceInput interface {
	pulumi.Input

	ToVolumeReferenceOutput() VolumeReferenceOutput
	ToVolumeReferenceOutputWithContext(context.Context) VolumeReferenceOutput
}

type VolumeReferenceArgs struct {
	DestinationPath pulumi.StringInput  `pulumi:"destinationPath"`
	Name            pulumi.StringInput  `pulumi:"name"`
	ReadOnly        pulumi.BoolPtrInput `pulumi:"readOnly"`
}

func (VolumeReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeReference)(nil)).Elem()
}

func (i VolumeReferenceArgs) ToVolumeReferenceOutput() VolumeReferenceOutput {
	return i.ToVolumeReferenceOutputWithContext(context.Background())
}

func (i VolumeReferenceArgs) ToVolumeReferenceOutputWithContext(ctx context.Context) VolumeReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeReferenceOutput)
}





type VolumeReferenceArrayInput interface {
	pulumi.Input

	ToVolumeReferenceArrayOutput() VolumeReferenceArrayOutput
	ToVolumeReferenceArrayOutputWithContext(context.Context) VolumeReferenceArrayOutput
}

type VolumeReferenceArray []VolumeReferenceInput

func (VolumeReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeReference)(nil)).Elem()
}

func (i VolumeReferenceArray) ToVolumeReferenceArrayOutput() VolumeReferenceArrayOutput {
	return i.ToVolumeReferenceArrayOutputWithContext(context.Background())
}

func (i VolumeReferenceArray) ToVolumeReferenceArrayOutputWithContext(ctx context.Context) VolumeReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeReferenceArrayOutput)
}

type VolumeReferenceOutput struct{ *pulumi.OutputState }

func (VolumeReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeReference)(nil)).Elem()
}

func (o VolumeReferenceOutput) ToVolumeReferenceOutput() VolumeReferenceOutput {
	return o
}

func (o VolumeReferenceOutput) ToVolumeReferenceOutputWithContext(ctx context.Context) VolumeReferenceOutput {
	return o
}

func (o VolumeReferenceOutput) DestinationPath() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeReference) string { return v.DestinationPath }).(pulumi.StringOutput)
}

func (o VolumeReferenceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeReference) string { return v.Name }).(pulumi.StringOutput)
}

func (o VolumeReferenceOutput) ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeReference) *bool { return v.ReadOnly }).(pulumi.BoolPtrOutput)
}

type VolumeReferenceArrayOutput struct{ *pulumi.OutputState }

func (VolumeReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeReference)(nil)).Elem()
}

func (o VolumeReferenceArrayOutput) ToVolumeReferenceArrayOutput() VolumeReferenceArrayOutput {
	return o
}

func (o VolumeReferenceArrayOutput) ToVolumeReferenceArrayOutputWithContext(ctx context.Context) VolumeReferenceArrayOutput {
	return o
}

func (o VolumeReferenceArrayOutput) Index(i pulumi.IntInput) VolumeReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeReference {
		return vs[0].([]VolumeReference)[vs[1].(int)]
	}).(VolumeReferenceOutput)
}

type VolumeReferenceResponse struct {
	DestinationPath string `pulumi:"destinationPath"`
	Name            string `pulumi:"name"`
	ReadOnly        *bool  `pulumi:"readOnly"`
}

type VolumeReferenceResponseOutput struct{ *pulumi.OutputState }

func (VolumeReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeReferenceResponse)(nil)).Elem()
}

func (o VolumeReferenceResponseOutput) ToVolumeReferenceResponseOutput() VolumeReferenceResponseOutput {
	return o
}

func (o VolumeReferenceResponseOutput) ToVolumeReferenceResponseOutputWithContext(ctx context.Context) VolumeReferenceResponseOutput {
	return o
}

func (o VolumeReferenceResponseOutput) DestinationPath() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeReferenceResponse) string { return v.DestinationPath }).(pulumi.StringOutput)
}

func (o VolumeReferenceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeReferenceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VolumeReferenceResponseOutput) ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeReferenceResponse) *bool { return v.ReadOnly }).(pulumi.BoolPtrOutput)
}

type VolumeReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (VolumeReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeReferenceResponse)(nil)).Elem()
}

func (o VolumeReferenceResponseArrayOutput) ToVolumeReferenceResponseArrayOutput() VolumeReferenceResponseArrayOutput {
	return o
}

func (o VolumeReferenceResponseArrayOutput) ToVolumeReferenceResponseArrayOutputWithContext(ctx context.Context) VolumeReferenceResponseArrayOutput {
	return o
}

func (o VolumeReferenceResponseArrayOutput) Index(i pulumi.IntInput) VolumeReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeReferenceResponse {
		return vs[0].([]VolumeReferenceResponse)[vs[1].(int)]
	}).(VolumeReferenceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AddRemoveReplicaScalingMechanismOutput{})
	pulumi.RegisterOutputType(AddRemoveReplicaScalingMechanismResponseOutput{})
	pulumi.RegisterOutputType(ApplicationScopedVolumeOutput{})
	pulumi.RegisterOutputType(ApplicationScopedVolumeArrayOutput{})
	pulumi.RegisterOutputType(ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskOutput{})
	pulumi.RegisterOutputType(ApplicationScopedVolumeCreationParametersServiceFabricVolumeDiskResponseOutput{})
	pulumi.RegisterOutputType(ApplicationScopedVolumeResponseOutput{})
	pulumi.RegisterOutputType(ApplicationScopedVolumeResponseArrayOutput{})
	pulumi.RegisterOutputType(AutoScalingPolicyOutput{})
	pulumi.RegisterOutputType(AutoScalingPolicyArrayOutput{})
	pulumi.RegisterOutputType(AutoScalingPolicyResponseOutput{})
	pulumi.RegisterOutputType(AutoScalingPolicyResponseArrayOutput{})
	pulumi.RegisterOutputType(AutoScalingResourceMetricOutput{})
	pulumi.RegisterOutputType(AutoScalingResourceMetricResponseOutput{})
	pulumi.RegisterOutputType(AverageLoadScalingTriggerOutput{})
	pulumi.RegisterOutputType(AverageLoadScalingTriggerResponseOutput{})
	pulumi.RegisterOutputType(AzureInternalMonitoringPipelineSinkDescriptionOutput{})
	pulumi.RegisterOutputType(AzureInternalMonitoringPipelineSinkDescriptionArrayOutput{})
	pulumi.RegisterOutputType(AzureInternalMonitoringPipelineSinkDescriptionResponseOutput{})
	pulumi.RegisterOutputType(AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerCodePackagePropertiesOutput{})
	pulumi.RegisterOutputType(ContainerCodePackagePropertiesArrayOutput{})
	pulumi.RegisterOutputType(ContainerCodePackagePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ContainerCodePackagePropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerEventResponseOutput{})
	pulumi.RegisterOutputType(ContainerEventResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(ContainerLabelOutput{})
	pulumi.RegisterOutputType(ContainerLabelArrayOutput{})
	pulumi.RegisterOutputType(ContainerLabelResponseOutput{})
	pulumi.RegisterOutputType(ContainerLabelResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerStateResponseOutput{})
	pulumi.RegisterOutputType(ContainerStateResponsePtrOutput{})
	pulumi.RegisterOutputType(DiagnosticsDescriptionOutput{})
	pulumi.RegisterOutputType(DiagnosticsDescriptionPtrOutput{})
	pulumi.RegisterOutputType(DiagnosticsDescriptionResponseOutput{})
	pulumi.RegisterOutputType(DiagnosticsDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(DiagnosticsRefOutput{})
	pulumi.RegisterOutputType(DiagnosticsRefPtrOutput{})
	pulumi.RegisterOutputType(DiagnosticsRefResponseOutput{})
	pulumi.RegisterOutputType(DiagnosticsRefResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesArrayOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(EndpointRefOutput{})
	pulumi.RegisterOutputType(EndpointRefArrayOutput{})
	pulumi.RegisterOutputType(EndpointRefResponseOutput{})
	pulumi.RegisterOutputType(EndpointRefResponseArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableResponseArrayOutput{})
	pulumi.RegisterOutputType(GatewayDestinationOutput{})
	pulumi.RegisterOutputType(GatewayDestinationResponseOutput{})
	pulumi.RegisterOutputType(HttpConfigOutput{})
	pulumi.RegisterOutputType(HttpConfigArrayOutput{})
	pulumi.RegisterOutputType(HttpConfigResponseOutput{})
	pulumi.RegisterOutputType(HttpConfigResponseArrayOutput{})
	pulumi.RegisterOutputType(HttpHostConfigOutput{})
	pulumi.RegisterOutputType(HttpHostConfigArrayOutput{})
	pulumi.RegisterOutputType(HttpHostConfigResponseOutput{})
	pulumi.RegisterOutputType(HttpHostConfigResponseArrayOutput{})
	pulumi.RegisterOutputType(HttpRouteConfigOutput{})
	pulumi.RegisterOutputType(HttpRouteConfigArrayOutput{})
	pulumi.RegisterOutputType(HttpRouteConfigResponseOutput{})
	pulumi.RegisterOutputType(HttpRouteConfigResponseArrayOutput{})
	pulumi.RegisterOutputType(HttpRouteMatchHeaderOutput{})
	pulumi.RegisterOutputType(HttpRouteMatchHeaderArrayOutput{})
	pulumi.RegisterOutputType(HttpRouteMatchHeaderResponseOutput{})
	pulumi.RegisterOutputType(HttpRouteMatchHeaderResponseArrayOutput{})
	pulumi.RegisterOutputType(HttpRouteMatchPathOutput{})
	pulumi.RegisterOutputType(HttpRouteMatchPathResponseOutput{})
	pulumi.RegisterOutputType(HttpRouteMatchRuleOutput{})
	pulumi.RegisterOutputType(HttpRouteMatchRuleResponseOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialPtrOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialResponseOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkRefOutput{})
	pulumi.RegisterOutputType(NetworkRefArrayOutput{})
	pulumi.RegisterOutputType(NetworkRefResponseOutput{})
	pulumi.RegisterOutputType(NetworkRefResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkResourcePropertiesOutput{})
	pulumi.RegisterOutputType(NetworkResourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ReliableCollectionsRefOutput{})
	pulumi.RegisterOutputType(ReliableCollectionsRefArrayOutput{})
	pulumi.RegisterOutputType(ReliableCollectionsRefResponseOutput{})
	pulumi.RegisterOutputType(ReliableCollectionsRefResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceLimitsOutput{})
	pulumi.RegisterOutputType(ResourceLimitsPtrOutput{})
	pulumi.RegisterOutputType(ResourceLimitsResponseOutput{})
	pulumi.RegisterOutputType(ResourceLimitsResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceRequestsOutput{})
	pulumi.RegisterOutputType(ResourceRequestsResponseOutput{})
	pulumi.RegisterOutputType(ResourceRequirementsOutput{})
	pulumi.RegisterOutputType(ResourceRequirementsResponseOutput{})
	pulumi.RegisterOutputType(SecretResourcePropertiesOutput{})
	pulumi.RegisterOutputType(SecretResourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ServiceResourceDescriptionOutput{})
	pulumi.RegisterOutputType(ServiceResourceDescriptionArrayOutput{})
	pulumi.RegisterOutputType(ServiceResourceDescriptionResponseOutput{})
	pulumi.RegisterOutputType(ServiceResourceDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(SettingOutput{})
	pulumi.RegisterOutputType(SettingArrayOutput{})
	pulumi.RegisterOutputType(SettingResponseOutput{})
	pulumi.RegisterOutputType(SettingResponseArrayOutput{})
	pulumi.RegisterOutputType(TcpConfigOutput{})
	pulumi.RegisterOutputType(TcpConfigArrayOutput{})
	pulumi.RegisterOutputType(TcpConfigResponseOutput{})
	pulumi.RegisterOutputType(TcpConfigResponseArrayOutput{})
	pulumi.RegisterOutputType(VolumeProviderParametersAzureFileOutput{})
	pulumi.RegisterOutputType(VolumeProviderParametersAzureFilePtrOutput{})
	pulumi.RegisterOutputType(VolumeProviderParametersAzureFileResponseOutput{})
	pulumi.RegisterOutputType(VolumeProviderParametersAzureFileResponsePtrOutput{})
	pulumi.RegisterOutputType(VolumeReferenceOutput{})
	pulumi.RegisterOutputType(VolumeReferenceArrayOutput{})
	pulumi.RegisterOutputType(VolumeReferenceResponseOutput{})
	pulumi.RegisterOutputType(VolumeReferenceResponseArrayOutput{})
}
