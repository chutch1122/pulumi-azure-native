


package v20180915

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicableScheduleResponse struct {
	Id             string            `pulumi:"id"`
	LabVmsShutdown *ScheduleResponse `pulumi:"labVmsShutdown"`
	LabVmsStartup  *ScheduleResponse `pulumi:"labVmsStartup"`
	Location       *string           `pulumi:"location"`
	Name           string            `pulumi:"name"`
	Tags           map[string]string `pulumi:"tags"`
	Type           string            `pulumi:"type"`
}


func (val *ApplicableScheduleResponse) Defaults() *ApplicableScheduleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.LabVmsShutdown = tmp.LabVmsShutdown.Defaults()

	tmp.LabVmsStartup = tmp.LabVmsStartup.Defaults()

	return &tmp
}

type ApplicableScheduleResponseOutput struct{ *pulumi.OutputState }

func (ApplicableScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicableScheduleResponse)(nil)).Elem()
}

func (o ApplicableScheduleResponseOutput) ToApplicableScheduleResponseOutput() ApplicableScheduleResponseOutput {
	return o
}

func (o ApplicableScheduleResponseOutput) ToApplicableScheduleResponseOutputWithContext(ctx context.Context) ApplicableScheduleResponseOutput {
	return o
}

func (o ApplicableScheduleResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicableScheduleResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ApplicableScheduleResponseOutput) LabVmsShutdown() ScheduleResponsePtrOutput {
	return o.ApplyT(func(v ApplicableScheduleResponse) *ScheduleResponse { return v.LabVmsShutdown }).(ScheduleResponsePtrOutput)
}

func (o ApplicableScheduleResponseOutput) LabVmsStartup() ScheduleResponsePtrOutput {
	return o.ApplyT(func(v ApplicableScheduleResponse) *ScheduleResponse { return v.LabVmsStartup }).(ScheduleResponsePtrOutput)
}

func (o ApplicableScheduleResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicableScheduleResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ApplicableScheduleResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicableScheduleResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicableScheduleResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ApplicableScheduleResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApplicableScheduleResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicableScheduleResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ArmTemplateParameterProperties struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type ArmTemplateParameterPropertiesInput interface {
	pulumi.Input

	ToArmTemplateParameterPropertiesOutput() ArmTemplateParameterPropertiesOutput
	ToArmTemplateParameterPropertiesOutputWithContext(context.Context) ArmTemplateParameterPropertiesOutput
}

type ArmTemplateParameterPropertiesArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ArmTemplateParameterPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmTemplateParameterProperties)(nil)).Elem()
}

func (i ArmTemplateParameterPropertiesArgs) ToArmTemplateParameterPropertiesOutput() ArmTemplateParameterPropertiesOutput {
	return i.ToArmTemplateParameterPropertiesOutputWithContext(context.Background())
}

func (i ArmTemplateParameterPropertiesArgs) ToArmTemplateParameterPropertiesOutputWithContext(ctx context.Context) ArmTemplateParameterPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmTemplateParameterPropertiesOutput)
}





type ArmTemplateParameterPropertiesArrayInput interface {
	pulumi.Input

	ToArmTemplateParameterPropertiesArrayOutput() ArmTemplateParameterPropertiesArrayOutput
	ToArmTemplateParameterPropertiesArrayOutputWithContext(context.Context) ArmTemplateParameterPropertiesArrayOutput
}

type ArmTemplateParameterPropertiesArray []ArmTemplateParameterPropertiesInput

func (ArmTemplateParameterPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArmTemplateParameterProperties)(nil)).Elem()
}

func (i ArmTemplateParameterPropertiesArray) ToArmTemplateParameterPropertiesArrayOutput() ArmTemplateParameterPropertiesArrayOutput {
	return i.ToArmTemplateParameterPropertiesArrayOutputWithContext(context.Background())
}

func (i ArmTemplateParameterPropertiesArray) ToArmTemplateParameterPropertiesArrayOutputWithContext(ctx context.Context) ArmTemplateParameterPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmTemplateParameterPropertiesArrayOutput)
}

type ArmTemplateParameterPropertiesOutput struct{ *pulumi.OutputState }

func (ArmTemplateParameterPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmTemplateParameterProperties)(nil)).Elem()
}

func (o ArmTemplateParameterPropertiesOutput) ToArmTemplateParameterPropertiesOutput() ArmTemplateParameterPropertiesOutput {
	return o
}

func (o ArmTemplateParameterPropertiesOutput) ToArmTemplateParameterPropertiesOutputWithContext(ctx context.Context) ArmTemplateParameterPropertiesOutput {
	return o
}

func (o ArmTemplateParameterPropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmTemplateParameterProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ArmTemplateParameterPropertiesOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmTemplateParameterProperties) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ArmTemplateParameterPropertiesArrayOutput struct{ *pulumi.OutputState }

func (ArmTemplateParameterPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArmTemplateParameterProperties)(nil)).Elem()
}

func (o ArmTemplateParameterPropertiesArrayOutput) ToArmTemplateParameterPropertiesArrayOutput() ArmTemplateParameterPropertiesArrayOutput {
	return o
}

func (o ArmTemplateParameterPropertiesArrayOutput) ToArmTemplateParameterPropertiesArrayOutputWithContext(ctx context.Context) ArmTemplateParameterPropertiesArrayOutput {
	return o
}

func (o ArmTemplateParameterPropertiesArrayOutput) Index(i pulumi.IntInput) ArmTemplateParameterPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArmTemplateParameterProperties {
		return vs[0].([]ArmTemplateParameterProperties)[vs[1].(int)]
	}).(ArmTemplateParameterPropertiesOutput)
}

type ArmTemplateParameterPropertiesResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type ArmTemplateParameterPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ArmTemplateParameterPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmTemplateParameterPropertiesResponse)(nil)).Elem()
}

func (o ArmTemplateParameterPropertiesResponseOutput) ToArmTemplateParameterPropertiesResponseOutput() ArmTemplateParameterPropertiesResponseOutput {
	return o
}

func (o ArmTemplateParameterPropertiesResponseOutput) ToArmTemplateParameterPropertiesResponseOutputWithContext(ctx context.Context) ArmTemplateParameterPropertiesResponseOutput {
	return o
}

func (o ArmTemplateParameterPropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmTemplateParameterPropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ArmTemplateParameterPropertiesResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmTemplateParameterPropertiesResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ArmTemplateParameterPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (ArmTemplateParameterPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArmTemplateParameterPropertiesResponse)(nil)).Elem()
}

func (o ArmTemplateParameterPropertiesResponseArrayOutput) ToArmTemplateParameterPropertiesResponseArrayOutput() ArmTemplateParameterPropertiesResponseArrayOutput {
	return o
}

func (o ArmTemplateParameterPropertiesResponseArrayOutput) ToArmTemplateParameterPropertiesResponseArrayOutputWithContext(ctx context.Context) ArmTemplateParameterPropertiesResponseArrayOutput {
	return o
}

func (o ArmTemplateParameterPropertiesResponseArrayOutput) Index(i pulumi.IntInput) ArmTemplateParameterPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArmTemplateParameterPropertiesResponse {
		return vs[0].([]ArmTemplateParameterPropertiesResponse)[vs[1].(int)]
	}).(ArmTemplateParameterPropertiesResponseOutput)
}

type ArtifactDeploymentStatusPropertiesResponse struct {
	ArtifactsApplied *int    `pulumi:"artifactsApplied"`
	DeploymentStatus *string `pulumi:"deploymentStatus"`
	TotalArtifacts   *int    `pulumi:"totalArtifacts"`
}

type ArtifactDeploymentStatusPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ArtifactDeploymentStatusPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactDeploymentStatusPropertiesResponse)(nil)).Elem()
}

func (o ArtifactDeploymentStatusPropertiesResponseOutput) ToArtifactDeploymentStatusPropertiesResponseOutput() ArtifactDeploymentStatusPropertiesResponseOutput {
	return o
}

func (o ArtifactDeploymentStatusPropertiesResponseOutput) ToArtifactDeploymentStatusPropertiesResponseOutputWithContext(ctx context.Context) ArtifactDeploymentStatusPropertiesResponseOutput {
	return o
}

func (o ArtifactDeploymentStatusPropertiesResponseOutput) ArtifactsApplied() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ArtifactDeploymentStatusPropertiesResponse) *int { return v.ArtifactsApplied }).(pulumi.IntPtrOutput)
}

func (o ArtifactDeploymentStatusPropertiesResponseOutput) DeploymentStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactDeploymentStatusPropertiesResponse) *string { return v.DeploymentStatus }).(pulumi.StringPtrOutput)
}

func (o ArtifactDeploymentStatusPropertiesResponseOutput) TotalArtifacts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ArtifactDeploymentStatusPropertiesResponse) *int { return v.TotalArtifacts }).(pulumi.IntPtrOutput)
}

type ArtifactInstallProperties struct {
	ArtifactId               *string                       `pulumi:"artifactId"`
	ArtifactTitle            *string                       `pulumi:"artifactTitle"`
	DeploymentStatusMessage  *string                       `pulumi:"deploymentStatusMessage"`
	InstallTime              *string                       `pulumi:"installTime"`
	Parameters               []ArtifactParameterProperties `pulumi:"parameters"`
	Status                   *string                       `pulumi:"status"`
	VmExtensionStatusMessage *string                       `pulumi:"vmExtensionStatusMessage"`
}





type ArtifactInstallPropertiesInput interface {
	pulumi.Input

	ToArtifactInstallPropertiesOutput() ArtifactInstallPropertiesOutput
	ToArtifactInstallPropertiesOutputWithContext(context.Context) ArtifactInstallPropertiesOutput
}

type ArtifactInstallPropertiesArgs struct {
	ArtifactId               pulumi.StringPtrInput                 `pulumi:"artifactId"`
	ArtifactTitle            pulumi.StringPtrInput                 `pulumi:"artifactTitle"`
	DeploymentStatusMessage  pulumi.StringPtrInput                 `pulumi:"deploymentStatusMessage"`
	InstallTime              pulumi.StringPtrInput                 `pulumi:"installTime"`
	Parameters               ArtifactParameterPropertiesArrayInput `pulumi:"parameters"`
	Status                   pulumi.StringPtrInput                 `pulumi:"status"`
	VmExtensionStatusMessage pulumi.StringPtrInput                 `pulumi:"vmExtensionStatusMessage"`
}

func (ArtifactInstallPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactInstallProperties)(nil)).Elem()
}

func (i ArtifactInstallPropertiesArgs) ToArtifactInstallPropertiesOutput() ArtifactInstallPropertiesOutput {
	return i.ToArtifactInstallPropertiesOutputWithContext(context.Background())
}

func (i ArtifactInstallPropertiesArgs) ToArtifactInstallPropertiesOutputWithContext(ctx context.Context) ArtifactInstallPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactInstallPropertiesOutput)
}





type ArtifactInstallPropertiesArrayInput interface {
	pulumi.Input

	ToArtifactInstallPropertiesArrayOutput() ArtifactInstallPropertiesArrayOutput
	ToArtifactInstallPropertiesArrayOutputWithContext(context.Context) ArtifactInstallPropertiesArrayOutput
}

type ArtifactInstallPropertiesArray []ArtifactInstallPropertiesInput

func (ArtifactInstallPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArtifactInstallProperties)(nil)).Elem()
}

func (i ArtifactInstallPropertiesArray) ToArtifactInstallPropertiesArrayOutput() ArtifactInstallPropertiesArrayOutput {
	return i.ToArtifactInstallPropertiesArrayOutputWithContext(context.Background())
}

func (i ArtifactInstallPropertiesArray) ToArtifactInstallPropertiesArrayOutputWithContext(ctx context.Context) ArtifactInstallPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactInstallPropertiesArrayOutput)
}

type ArtifactInstallPropertiesOutput struct{ *pulumi.OutputState }

func (ArtifactInstallPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactInstallProperties)(nil)).Elem()
}

func (o ArtifactInstallPropertiesOutput) ToArtifactInstallPropertiesOutput() ArtifactInstallPropertiesOutput {
	return o
}

func (o ArtifactInstallPropertiesOutput) ToArtifactInstallPropertiesOutputWithContext(ctx context.Context) ArtifactInstallPropertiesOutput {
	return o
}

func (o ArtifactInstallPropertiesOutput) ArtifactId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactInstallProperties) *string { return v.ArtifactId }).(pulumi.StringPtrOutput)
}

func (o ArtifactInstallPropertiesOutput) ArtifactTitle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactInstallProperties) *string { return v.ArtifactTitle }).(pulumi.StringPtrOutput)
}

func (o ArtifactInstallPropertiesOutput) DeploymentStatusMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactInstallProperties) *string { return v.DeploymentStatusMessage }).(pulumi.StringPtrOutput)
}

func (o ArtifactInstallPropertiesOutput) InstallTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactInstallProperties) *string { return v.InstallTime }).(pulumi.StringPtrOutput)
}

func (o ArtifactInstallPropertiesOutput) Parameters() ArtifactParameterPropertiesArrayOutput {
	return o.ApplyT(func(v ArtifactInstallProperties) []ArtifactParameterProperties { return v.Parameters }).(ArtifactParameterPropertiesArrayOutput)
}

func (o ArtifactInstallPropertiesOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactInstallProperties) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ArtifactInstallPropertiesOutput) VmExtensionStatusMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactInstallProperties) *string { return v.VmExtensionStatusMessage }).(pulumi.StringPtrOutput)
}

type ArtifactInstallPropertiesArrayOutput struct{ *pulumi.OutputState }

func (ArtifactInstallPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArtifactInstallProperties)(nil)).Elem()
}

func (o ArtifactInstallPropertiesArrayOutput) ToArtifactInstallPropertiesArrayOutput() ArtifactInstallPropertiesArrayOutput {
	return o
}

func (o ArtifactInstallPropertiesArrayOutput) ToArtifactInstallPropertiesArrayOutputWithContext(ctx context.Context) ArtifactInstallPropertiesArrayOutput {
	return o
}

func (o ArtifactInstallPropertiesArrayOutput) Index(i pulumi.IntInput) ArtifactInstallPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArtifactInstallProperties {
		return vs[0].([]ArtifactInstallProperties)[vs[1].(int)]
	}).(ArtifactInstallPropertiesOutput)
}

type ArtifactInstallPropertiesResponse struct {
	ArtifactId               *string                               `pulumi:"artifactId"`
	ArtifactTitle            *string                               `pulumi:"artifactTitle"`
	DeploymentStatusMessage  *string                               `pulumi:"deploymentStatusMessage"`
	InstallTime              *string                               `pulumi:"installTime"`
	Parameters               []ArtifactParameterPropertiesResponse `pulumi:"parameters"`
	Status                   *string                               `pulumi:"status"`
	VmExtensionStatusMessage *string                               `pulumi:"vmExtensionStatusMessage"`
}

type ArtifactInstallPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ArtifactInstallPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactInstallPropertiesResponse)(nil)).Elem()
}

func (o ArtifactInstallPropertiesResponseOutput) ToArtifactInstallPropertiesResponseOutput() ArtifactInstallPropertiesResponseOutput {
	return o
}

func (o ArtifactInstallPropertiesResponseOutput) ToArtifactInstallPropertiesResponseOutputWithContext(ctx context.Context) ArtifactInstallPropertiesResponseOutput {
	return o
}

func (o ArtifactInstallPropertiesResponseOutput) ArtifactId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactInstallPropertiesResponse) *string { return v.ArtifactId }).(pulumi.StringPtrOutput)
}

func (o ArtifactInstallPropertiesResponseOutput) ArtifactTitle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactInstallPropertiesResponse) *string { return v.ArtifactTitle }).(pulumi.StringPtrOutput)
}

func (o ArtifactInstallPropertiesResponseOutput) DeploymentStatusMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactInstallPropertiesResponse) *string { return v.DeploymentStatusMessage }).(pulumi.StringPtrOutput)
}

func (o ArtifactInstallPropertiesResponseOutput) InstallTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactInstallPropertiesResponse) *string { return v.InstallTime }).(pulumi.StringPtrOutput)
}

func (o ArtifactInstallPropertiesResponseOutput) Parameters() ArtifactParameterPropertiesResponseArrayOutput {
	return o.ApplyT(func(v ArtifactInstallPropertiesResponse) []ArtifactParameterPropertiesResponse { return v.Parameters }).(ArtifactParameterPropertiesResponseArrayOutput)
}

func (o ArtifactInstallPropertiesResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactInstallPropertiesResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ArtifactInstallPropertiesResponseOutput) VmExtensionStatusMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactInstallPropertiesResponse) *string { return v.VmExtensionStatusMessage }).(pulumi.StringPtrOutput)
}

type ArtifactInstallPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (ArtifactInstallPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArtifactInstallPropertiesResponse)(nil)).Elem()
}

func (o ArtifactInstallPropertiesResponseArrayOutput) ToArtifactInstallPropertiesResponseArrayOutput() ArtifactInstallPropertiesResponseArrayOutput {
	return o
}

func (o ArtifactInstallPropertiesResponseArrayOutput) ToArtifactInstallPropertiesResponseArrayOutputWithContext(ctx context.Context) ArtifactInstallPropertiesResponseArrayOutput {
	return o
}

func (o ArtifactInstallPropertiesResponseArrayOutput) Index(i pulumi.IntInput) ArtifactInstallPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArtifactInstallPropertiesResponse {
		return vs[0].([]ArtifactInstallPropertiesResponse)[vs[1].(int)]
	}).(ArtifactInstallPropertiesResponseOutput)
}

type ArtifactParameterProperties struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type ArtifactParameterPropertiesInput interface {
	pulumi.Input

	ToArtifactParameterPropertiesOutput() ArtifactParameterPropertiesOutput
	ToArtifactParameterPropertiesOutputWithContext(context.Context) ArtifactParameterPropertiesOutput
}

type ArtifactParameterPropertiesArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ArtifactParameterPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactParameterProperties)(nil)).Elem()
}

func (i ArtifactParameterPropertiesArgs) ToArtifactParameterPropertiesOutput() ArtifactParameterPropertiesOutput {
	return i.ToArtifactParameterPropertiesOutputWithContext(context.Background())
}

func (i ArtifactParameterPropertiesArgs) ToArtifactParameterPropertiesOutputWithContext(ctx context.Context) ArtifactParameterPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactParameterPropertiesOutput)
}





type ArtifactParameterPropertiesArrayInput interface {
	pulumi.Input

	ToArtifactParameterPropertiesArrayOutput() ArtifactParameterPropertiesArrayOutput
	ToArtifactParameterPropertiesArrayOutputWithContext(context.Context) ArtifactParameterPropertiesArrayOutput
}

type ArtifactParameterPropertiesArray []ArtifactParameterPropertiesInput

func (ArtifactParameterPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArtifactParameterProperties)(nil)).Elem()
}

func (i ArtifactParameterPropertiesArray) ToArtifactParameterPropertiesArrayOutput() ArtifactParameterPropertiesArrayOutput {
	return i.ToArtifactParameterPropertiesArrayOutputWithContext(context.Background())
}

func (i ArtifactParameterPropertiesArray) ToArtifactParameterPropertiesArrayOutputWithContext(ctx context.Context) ArtifactParameterPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactParameterPropertiesArrayOutput)
}

type ArtifactParameterPropertiesOutput struct{ *pulumi.OutputState }

func (ArtifactParameterPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactParameterProperties)(nil)).Elem()
}

func (o ArtifactParameterPropertiesOutput) ToArtifactParameterPropertiesOutput() ArtifactParameterPropertiesOutput {
	return o
}

func (o ArtifactParameterPropertiesOutput) ToArtifactParameterPropertiesOutputWithContext(ctx context.Context) ArtifactParameterPropertiesOutput {
	return o
}

func (o ArtifactParameterPropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactParameterProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ArtifactParameterPropertiesOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactParameterProperties) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ArtifactParameterPropertiesArrayOutput struct{ *pulumi.OutputState }

func (ArtifactParameterPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArtifactParameterProperties)(nil)).Elem()
}

func (o ArtifactParameterPropertiesArrayOutput) ToArtifactParameterPropertiesArrayOutput() ArtifactParameterPropertiesArrayOutput {
	return o
}

func (o ArtifactParameterPropertiesArrayOutput) ToArtifactParameterPropertiesArrayOutputWithContext(ctx context.Context) ArtifactParameterPropertiesArrayOutput {
	return o
}

func (o ArtifactParameterPropertiesArrayOutput) Index(i pulumi.IntInput) ArtifactParameterPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArtifactParameterProperties {
		return vs[0].([]ArtifactParameterProperties)[vs[1].(int)]
	}).(ArtifactParameterPropertiesOutput)
}

type ArtifactParameterPropertiesResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type ArtifactParameterPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ArtifactParameterPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactParameterPropertiesResponse)(nil)).Elem()
}

func (o ArtifactParameterPropertiesResponseOutput) ToArtifactParameterPropertiesResponseOutput() ArtifactParameterPropertiesResponseOutput {
	return o
}

func (o ArtifactParameterPropertiesResponseOutput) ToArtifactParameterPropertiesResponseOutputWithContext(ctx context.Context) ArtifactParameterPropertiesResponseOutput {
	return o
}

func (o ArtifactParameterPropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactParameterPropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ArtifactParameterPropertiesResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArtifactParameterPropertiesResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ArtifactParameterPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (ArtifactParameterPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArtifactParameterPropertiesResponse)(nil)).Elem()
}

func (o ArtifactParameterPropertiesResponseArrayOutput) ToArtifactParameterPropertiesResponseArrayOutput() ArtifactParameterPropertiesResponseArrayOutput {
	return o
}

func (o ArtifactParameterPropertiesResponseArrayOutput) ToArtifactParameterPropertiesResponseArrayOutputWithContext(ctx context.Context) ArtifactParameterPropertiesResponseArrayOutput {
	return o
}

func (o ArtifactParameterPropertiesResponseArrayOutput) Index(i pulumi.IntInput) ArtifactParameterPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArtifactParameterPropertiesResponse {
		return vs[0].([]ArtifactParameterPropertiesResponse)[vs[1].(int)]
	}).(ArtifactParameterPropertiesResponseOutput)
}

type AttachNewDataDiskOptions struct {
	DiskName    *string `pulumi:"diskName"`
	DiskSizeGiB *int    `pulumi:"diskSizeGiB"`
	DiskType    *string `pulumi:"diskType"`
}





type AttachNewDataDiskOptionsInput interface {
	pulumi.Input

	ToAttachNewDataDiskOptionsOutput() AttachNewDataDiskOptionsOutput
	ToAttachNewDataDiskOptionsOutputWithContext(context.Context) AttachNewDataDiskOptionsOutput
}

type AttachNewDataDiskOptionsArgs struct {
	DiskName    pulumi.StringPtrInput `pulumi:"diskName"`
	DiskSizeGiB pulumi.IntPtrInput    `pulumi:"diskSizeGiB"`
	DiskType    pulumi.StringPtrInput `pulumi:"diskType"`
}

func (AttachNewDataDiskOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AttachNewDataDiskOptions)(nil)).Elem()
}

func (i AttachNewDataDiskOptionsArgs) ToAttachNewDataDiskOptionsOutput() AttachNewDataDiskOptionsOutput {
	return i.ToAttachNewDataDiskOptionsOutputWithContext(context.Background())
}

func (i AttachNewDataDiskOptionsArgs) ToAttachNewDataDiskOptionsOutputWithContext(ctx context.Context) AttachNewDataDiskOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachNewDataDiskOptionsOutput)
}

func (i AttachNewDataDiskOptionsArgs) ToAttachNewDataDiskOptionsPtrOutput() AttachNewDataDiskOptionsPtrOutput {
	return i.ToAttachNewDataDiskOptionsPtrOutputWithContext(context.Background())
}

func (i AttachNewDataDiskOptionsArgs) ToAttachNewDataDiskOptionsPtrOutputWithContext(ctx context.Context) AttachNewDataDiskOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachNewDataDiskOptionsOutput).ToAttachNewDataDiskOptionsPtrOutputWithContext(ctx)
}









type AttachNewDataDiskOptionsPtrInput interface {
	pulumi.Input

	ToAttachNewDataDiskOptionsPtrOutput() AttachNewDataDiskOptionsPtrOutput
	ToAttachNewDataDiskOptionsPtrOutputWithContext(context.Context) AttachNewDataDiskOptionsPtrOutput
}

type attachNewDataDiskOptionsPtrType AttachNewDataDiskOptionsArgs

func AttachNewDataDiskOptionsPtr(v *AttachNewDataDiskOptionsArgs) AttachNewDataDiskOptionsPtrInput {
	return (*attachNewDataDiskOptionsPtrType)(v)
}

func (*attachNewDataDiskOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachNewDataDiskOptions)(nil)).Elem()
}

func (i *attachNewDataDiskOptionsPtrType) ToAttachNewDataDiskOptionsPtrOutput() AttachNewDataDiskOptionsPtrOutput {
	return i.ToAttachNewDataDiskOptionsPtrOutputWithContext(context.Background())
}

func (i *attachNewDataDiskOptionsPtrType) ToAttachNewDataDiskOptionsPtrOutputWithContext(ctx context.Context) AttachNewDataDiskOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachNewDataDiskOptionsPtrOutput)
}

type AttachNewDataDiskOptionsOutput struct{ *pulumi.OutputState }

func (AttachNewDataDiskOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttachNewDataDiskOptions)(nil)).Elem()
}

func (o AttachNewDataDiskOptionsOutput) ToAttachNewDataDiskOptionsOutput() AttachNewDataDiskOptionsOutput {
	return o
}

func (o AttachNewDataDiskOptionsOutput) ToAttachNewDataDiskOptionsOutputWithContext(ctx context.Context) AttachNewDataDiskOptionsOutput {
	return o
}

func (o AttachNewDataDiskOptionsOutput) ToAttachNewDataDiskOptionsPtrOutput() AttachNewDataDiskOptionsPtrOutput {
	return o.ToAttachNewDataDiskOptionsPtrOutputWithContext(context.Background())
}

func (o AttachNewDataDiskOptionsOutput) ToAttachNewDataDiskOptionsPtrOutputWithContext(ctx context.Context) AttachNewDataDiskOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AttachNewDataDiskOptions) *AttachNewDataDiskOptions {
		return &v
	}).(AttachNewDataDiskOptionsPtrOutput)
}

func (o AttachNewDataDiskOptionsOutput) DiskName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AttachNewDataDiskOptions) *string { return v.DiskName }).(pulumi.StringPtrOutput)
}

func (o AttachNewDataDiskOptionsOutput) DiskSizeGiB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AttachNewDataDiskOptions) *int { return v.DiskSizeGiB }).(pulumi.IntPtrOutput)
}

func (o AttachNewDataDiskOptionsOutput) DiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AttachNewDataDiskOptions) *string { return v.DiskType }).(pulumi.StringPtrOutput)
}

type AttachNewDataDiskOptionsPtrOutput struct{ *pulumi.OutputState }

func (AttachNewDataDiskOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachNewDataDiskOptions)(nil)).Elem()
}

func (o AttachNewDataDiskOptionsPtrOutput) ToAttachNewDataDiskOptionsPtrOutput() AttachNewDataDiskOptionsPtrOutput {
	return o
}

func (o AttachNewDataDiskOptionsPtrOutput) ToAttachNewDataDiskOptionsPtrOutputWithContext(ctx context.Context) AttachNewDataDiskOptionsPtrOutput {
	return o
}

func (o AttachNewDataDiskOptionsPtrOutput) Elem() AttachNewDataDiskOptionsOutput {
	return o.ApplyT(func(v *AttachNewDataDiskOptions) AttachNewDataDiskOptions {
		if v != nil {
			return *v
		}
		var ret AttachNewDataDiskOptions
		return ret
	}).(AttachNewDataDiskOptionsOutput)
}

func (o AttachNewDataDiskOptionsPtrOutput) DiskName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachNewDataDiskOptions) *string {
		if v == nil {
			return nil
		}
		return v.DiskName
	}).(pulumi.StringPtrOutput)
}

func (o AttachNewDataDiskOptionsPtrOutput) DiskSizeGiB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AttachNewDataDiskOptions) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGiB
	}).(pulumi.IntPtrOutput)
}

func (o AttachNewDataDiskOptionsPtrOutput) DiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachNewDataDiskOptions) *string {
		if v == nil {
			return nil
		}
		return v.DiskType
	}).(pulumi.StringPtrOutput)
}

type AttachNewDataDiskOptionsResponse struct {
	DiskName    *string `pulumi:"diskName"`
	DiskSizeGiB *int    `pulumi:"diskSizeGiB"`
	DiskType    *string `pulumi:"diskType"`
}

type AttachNewDataDiskOptionsResponseOutput struct{ *pulumi.OutputState }

func (AttachNewDataDiskOptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttachNewDataDiskOptionsResponse)(nil)).Elem()
}

func (o AttachNewDataDiskOptionsResponseOutput) ToAttachNewDataDiskOptionsResponseOutput() AttachNewDataDiskOptionsResponseOutput {
	return o
}

func (o AttachNewDataDiskOptionsResponseOutput) ToAttachNewDataDiskOptionsResponseOutputWithContext(ctx context.Context) AttachNewDataDiskOptionsResponseOutput {
	return o
}

func (o AttachNewDataDiskOptionsResponseOutput) DiskName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AttachNewDataDiskOptionsResponse) *string { return v.DiskName }).(pulumi.StringPtrOutput)
}

func (o AttachNewDataDiskOptionsResponseOutput) DiskSizeGiB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AttachNewDataDiskOptionsResponse) *int { return v.DiskSizeGiB }).(pulumi.IntPtrOutput)
}

func (o AttachNewDataDiskOptionsResponseOutput) DiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AttachNewDataDiskOptionsResponse) *string { return v.DiskType }).(pulumi.StringPtrOutput)
}

type AttachNewDataDiskOptionsResponsePtrOutput struct{ *pulumi.OutputState }

func (AttachNewDataDiskOptionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachNewDataDiskOptionsResponse)(nil)).Elem()
}

func (o AttachNewDataDiskOptionsResponsePtrOutput) ToAttachNewDataDiskOptionsResponsePtrOutput() AttachNewDataDiskOptionsResponsePtrOutput {
	return o
}

func (o AttachNewDataDiskOptionsResponsePtrOutput) ToAttachNewDataDiskOptionsResponsePtrOutputWithContext(ctx context.Context) AttachNewDataDiskOptionsResponsePtrOutput {
	return o
}

func (o AttachNewDataDiskOptionsResponsePtrOutput) Elem() AttachNewDataDiskOptionsResponseOutput {
	return o.ApplyT(func(v *AttachNewDataDiskOptionsResponse) AttachNewDataDiskOptionsResponse {
		if v != nil {
			return *v
		}
		var ret AttachNewDataDiskOptionsResponse
		return ret
	}).(AttachNewDataDiskOptionsResponseOutput)
}

func (o AttachNewDataDiskOptionsResponsePtrOutput) DiskName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachNewDataDiskOptionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DiskName
	}).(pulumi.StringPtrOutput)
}

func (o AttachNewDataDiskOptionsResponsePtrOutput) DiskSizeGiB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AttachNewDataDiskOptionsResponse) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGiB
	}).(pulumi.IntPtrOutput)
}

func (o AttachNewDataDiskOptionsResponsePtrOutput) DiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachNewDataDiskOptionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DiskType
	}).(pulumi.StringPtrOutput)
}

type BulkCreationParameters struct {
	InstanceCount *int `pulumi:"instanceCount"`
}





type BulkCreationParametersInput interface {
	pulumi.Input

	ToBulkCreationParametersOutput() BulkCreationParametersOutput
	ToBulkCreationParametersOutputWithContext(context.Context) BulkCreationParametersOutput
}

type BulkCreationParametersArgs struct {
	InstanceCount pulumi.IntPtrInput `pulumi:"instanceCount"`
}

func (BulkCreationParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BulkCreationParameters)(nil)).Elem()
}

func (i BulkCreationParametersArgs) ToBulkCreationParametersOutput() BulkCreationParametersOutput {
	return i.ToBulkCreationParametersOutputWithContext(context.Background())
}

func (i BulkCreationParametersArgs) ToBulkCreationParametersOutputWithContext(ctx context.Context) BulkCreationParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BulkCreationParametersOutput)
}

func (i BulkCreationParametersArgs) ToBulkCreationParametersPtrOutput() BulkCreationParametersPtrOutput {
	return i.ToBulkCreationParametersPtrOutputWithContext(context.Background())
}

func (i BulkCreationParametersArgs) ToBulkCreationParametersPtrOutputWithContext(ctx context.Context) BulkCreationParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BulkCreationParametersOutput).ToBulkCreationParametersPtrOutputWithContext(ctx)
}









type BulkCreationParametersPtrInput interface {
	pulumi.Input

	ToBulkCreationParametersPtrOutput() BulkCreationParametersPtrOutput
	ToBulkCreationParametersPtrOutputWithContext(context.Context) BulkCreationParametersPtrOutput
}

type bulkCreationParametersPtrType BulkCreationParametersArgs

func BulkCreationParametersPtr(v *BulkCreationParametersArgs) BulkCreationParametersPtrInput {
	return (*bulkCreationParametersPtrType)(v)
}

func (*bulkCreationParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BulkCreationParameters)(nil)).Elem()
}

func (i *bulkCreationParametersPtrType) ToBulkCreationParametersPtrOutput() BulkCreationParametersPtrOutput {
	return i.ToBulkCreationParametersPtrOutputWithContext(context.Background())
}

func (i *bulkCreationParametersPtrType) ToBulkCreationParametersPtrOutputWithContext(ctx context.Context) BulkCreationParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BulkCreationParametersPtrOutput)
}

type BulkCreationParametersOutput struct{ *pulumi.OutputState }

func (BulkCreationParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BulkCreationParameters)(nil)).Elem()
}

func (o BulkCreationParametersOutput) ToBulkCreationParametersOutput() BulkCreationParametersOutput {
	return o
}

func (o BulkCreationParametersOutput) ToBulkCreationParametersOutputWithContext(ctx context.Context) BulkCreationParametersOutput {
	return o
}

func (o BulkCreationParametersOutput) ToBulkCreationParametersPtrOutput() BulkCreationParametersPtrOutput {
	return o.ToBulkCreationParametersPtrOutputWithContext(context.Background())
}

func (o BulkCreationParametersOutput) ToBulkCreationParametersPtrOutputWithContext(ctx context.Context) BulkCreationParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BulkCreationParameters) *BulkCreationParameters {
		return &v
	}).(BulkCreationParametersPtrOutput)
}

func (o BulkCreationParametersOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BulkCreationParameters) *int { return v.InstanceCount }).(pulumi.IntPtrOutput)
}

type BulkCreationParametersPtrOutput struct{ *pulumi.OutputState }

func (BulkCreationParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BulkCreationParameters)(nil)).Elem()
}

func (o BulkCreationParametersPtrOutput) ToBulkCreationParametersPtrOutput() BulkCreationParametersPtrOutput {
	return o
}

func (o BulkCreationParametersPtrOutput) ToBulkCreationParametersPtrOutputWithContext(ctx context.Context) BulkCreationParametersPtrOutput {
	return o
}

func (o BulkCreationParametersPtrOutput) Elem() BulkCreationParametersOutput {
	return o.ApplyT(func(v *BulkCreationParameters) BulkCreationParameters {
		if v != nil {
			return *v
		}
		var ret BulkCreationParameters
		return ret
	}).(BulkCreationParametersOutput)
}

func (o BulkCreationParametersPtrOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BulkCreationParameters) *int {
		if v == nil {
			return nil
		}
		return v.InstanceCount
	}).(pulumi.IntPtrOutput)
}

type BulkCreationParametersResponse struct {
	InstanceCount *int `pulumi:"instanceCount"`
}

type BulkCreationParametersResponseOutput struct{ *pulumi.OutputState }

func (BulkCreationParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BulkCreationParametersResponse)(nil)).Elem()
}

func (o BulkCreationParametersResponseOutput) ToBulkCreationParametersResponseOutput() BulkCreationParametersResponseOutput {
	return o
}

func (o BulkCreationParametersResponseOutput) ToBulkCreationParametersResponseOutputWithContext(ctx context.Context) BulkCreationParametersResponseOutput {
	return o
}

func (o BulkCreationParametersResponseOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BulkCreationParametersResponse) *int { return v.InstanceCount }).(pulumi.IntPtrOutput)
}

type BulkCreationParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (BulkCreationParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BulkCreationParametersResponse)(nil)).Elem()
}

func (o BulkCreationParametersResponsePtrOutput) ToBulkCreationParametersResponsePtrOutput() BulkCreationParametersResponsePtrOutput {
	return o
}

func (o BulkCreationParametersResponsePtrOutput) ToBulkCreationParametersResponsePtrOutputWithContext(ctx context.Context) BulkCreationParametersResponsePtrOutput {
	return o
}

func (o BulkCreationParametersResponsePtrOutput) Elem() BulkCreationParametersResponseOutput {
	return o.ApplyT(func(v *BulkCreationParametersResponse) BulkCreationParametersResponse {
		if v != nil {
			return *v
		}
		var ret BulkCreationParametersResponse
		return ret
	}).(BulkCreationParametersResponseOutput)
}

func (o BulkCreationParametersResponsePtrOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BulkCreationParametersResponse) *int {
		if v == nil {
			return nil
		}
		return v.InstanceCount
	}).(pulumi.IntPtrOutput)
}

type ComputeDataDiskResponse struct {
	DiskSizeGiB   *int    `pulumi:"diskSizeGiB"`
	DiskUri       *string `pulumi:"diskUri"`
	ManagedDiskId *string `pulumi:"managedDiskId"`
	Name          *string `pulumi:"name"`
}

type ComputeDataDiskResponseOutput struct{ *pulumi.OutputState }

func (ComputeDataDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeDataDiskResponse)(nil)).Elem()
}

func (o ComputeDataDiskResponseOutput) ToComputeDataDiskResponseOutput() ComputeDataDiskResponseOutput {
	return o
}

func (o ComputeDataDiskResponseOutput) ToComputeDataDiskResponseOutputWithContext(ctx context.Context) ComputeDataDiskResponseOutput {
	return o
}

func (o ComputeDataDiskResponseOutput) DiskSizeGiB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ComputeDataDiskResponse) *int { return v.DiskSizeGiB }).(pulumi.IntPtrOutput)
}

func (o ComputeDataDiskResponseOutput) DiskUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeDataDiskResponse) *string { return v.DiskUri }).(pulumi.StringPtrOutput)
}

func (o ComputeDataDiskResponseOutput) ManagedDiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeDataDiskResponse) *string { return v.ManagedDiskId }).(pulumi.StringPtrOutput)
}

func (o ComputeDataDiskResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeDataDiskResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ComputeDataDiskResponseArrayOutput struct{ *pulumi.OutputState }

func (ComputeDataDiskResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ComputeDataDiskResponse)(nil)).Elem()
}

func (o ComputeDataDiskResponseArrayOutput) ToComputeDataDiskResponseArrayOutput() ComputeDataDiskResponseArrayOutput {
	return o
}

func (o ComputeDataDiskResponseArrayOutput) ToComputeDataDiskResponseArrayOutputWithContext(ctx context.Context) ComputeDataDiskResponseArrayOutput {
	return o
}

func (o ComputeDataDiskResponseArrayOutput) Index(i pulumi.IntInput) ComputeDataDiskResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ComputeDataDiskResponse {
		return vs[0].([]ComputeDataDiskResponse)[vs[1].(int)]
	}).(ComputeDataDiskResponseOutput)
}

type ComputeVmInstanceViewStatusResponse struct {
	Code          *string `pulumi:"code"`
	DisplayStatus *string `pulumi:"displayStatus"`
	Message       *string `pulumi:"message"`
}

type ComputeVmInstanceViewStatusResponseOutput struct{ *pulumi.OutputState }

func (ComputeVmInstanceViewStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeVmInstanceViewStatusResponse)(nil)).Elem()
}

func (o ComputeVmInstanceViewStatusResponseOutput) ToComputeVmInstanceViewStatusResponseOutput() ComputeVmInstanceViewStatusResponseOutput {
	return o
}

func (o ComputeVmInstanceViewStatusResponseOutput) ToComputeVmInstanceViewStatusResponseOutputWithContext(ctx context.Context) ComputeVmInstanceViewStatusResponseOutput {
	return o
}

func (o ComputeVmInstanceViewStatusResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeVmInstanceViewStatusResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ComputeVmInstanceViewStatusResponseOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeVmInstanceViewStatusResponse) *string { return v.DisplayStatus }).(pulumi.StringPtrOutput)
}

func (o ComputeVmInstanceViewStatusResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeVmInstanceViewStatusResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type ComputeVmInstanceViewStatusResponseArrayOutput struct{ *pulumi.OutputState }

func (ComputeVmInstanceViewStatusResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ComputeVmInstanceViewStatusResponse)(nil)).Elem()
}

func (o ComputeVmInstanceViewStatusResponseArrayOutput) ToComputeVmInstanceViewStatusResponseArrayOutput() ComputeVmInstanceViewStatusResponseArrayOutput {
	return o
}

func (o ComputeVmInstanceViewStatusResponseArrayOutput) ToComputeVmInstanceViewStatusResponseArrayOutputWithContext(ctx context.Context) ComputeVmInstanceViewStatusResponseArrayOutput {
	return o
}

func (o ComputeVmInstanceViewStatusResponseArrayOutput) Index(i pulumi.IntInput) ComputeVmInstanceViewStatusResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ComputeVmInstanceViewStatusResponse {
		return vs[0].([]ComputeVmInstanceViewStatusResponse)[vs[1].(int)]
	}).(ComputeVmInstanceViewStatusResponseOutput)
}

type ComputeVmPropertiesResponse struct {
	DataDiskIds        []string                              `pulumi:"dataDiskIds"`
	DataDisks          []ComputeDataDiskResponse             `pulumi:"dataDisks"`
	NetworkInterfaceId *string                               `pulumi:"networkInterfaceId"`
	OsDiskId           *string                               `pulumi:"osDiskId"`
	OsType             *string                               `pulumi:"osType"`
	Statuses           []ComputeVmInstanceViewStatusResponse `pulumi:"statuses"`
	VmSize             *string                               `pulumi:"vmSize"`
}

type ComputeVmPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ComputeVmPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeVmPropertiesResponse)(nil)).Elem()
}

func (o ComputeVmPropertiesResponseOutput) ToComputeVmPropertiesResponseOutput() ComputeVmPropertiesResponseOutput {
	return o
}

func (o ComputeVmPropertiesResponseOutput) ToComputeVmPropertiesResponseOutputWithContext(ctx context.Context) ComputeVmPropertiesResponseOutput {
	return o
}

func (o ComputeVmPropertiesResponseOutput) DataDiskIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ComputeVmPropertiesResponse) []string { return v.DataDiskIds }).(pulumi.StringArrayOutput)
}

func (o ComputeVmPropertiesResponseOutput) DataDisks() ComputeDataDiskResponseArrayOutput {
	return o.ApplyT(func(v ComputeVmPropertiesResponse) []ComputeDataDiskResponse { return v.DataDisks }).(ComputeDataDiskResponseArrayOutput)
}

func (o ComputeVmPropertiesResponseOutput) NetworkInterfaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeVmPropertiesResponse) *string { return v.NetworkInterfaceId }).(pulumi.StringPtrOutput)
}

func (o ComputeVmPropertiesResponseOutput) OsDiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeVmPropertiesResponse) *string { return v.OsDiskId }).(pulumi.StringPtrOutput)
}

func (o ComputeVmPropertiesResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeVmPropertiesResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o ComputeVmPropertiesResponseOutput) Statuses() ComputeVmInstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v ComputeVmPropertiesResponse) []ComputeVmInstanceViewStatusResponse { return v.Statuses }).(ComputeVmInstanceViewStatusResponseArrayOutput)
}

func (o ComputeVmPropertiesResponseOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ComputeVmPropertiesResponse) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type CustomImagePropertiesCustom struct {
	ImageName *string `pulumi:"imageName"`
	OsType    string  `pulumi:"osType"`
	SysPrep   *bool   `pulumi:"sysPrep"`
}





type CustomImagePropertiesCustomInput interface {
	pulumi.Input

	ToCustomImagePropertiesCustomOutput() CustomImagePropertiesCustomOutput
	ToCustomImagePropertiesCustomOutputWithContext(context.Context) CustomImagePropertiesCustomOutput
}

type CustomImagePropertiesCustomArgs struct {
	ImageName pulumi.StringPtrInput `pulumi:"imageName"`
	OsType    pulumi.StringInput    `pulumi:"osType"`
	SysPrep   pulumi.BoolPtrInput   `pulumi:"sysPrep"`
}

func (CustomImagePropertiesCustomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImagePropertiesCustom)(nil)).Elem()
}

func (i CustomImagePropertiesCustomArgs) ToCustomImagePropertiesCustomOutput() CustomImagePropertiesCustomOutput {
	return i.ToCustomImagePropertiesCustomOutputWithContext(context.Background())
}

func (i CustomImagePropertiesCustomArgs) ToCustomImagePropertiesCustomOutputWithContext(ctx context.Context) CustomImagePropertiesCustomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesCustomOutput)
}

func (i CustomImagePropertiesCustomArgs) ToCustomImagePropertiesCustomPtrOutput() CustomImagePropertiesCustomPtrOutput {
	return i.ToCustomImagePropertiesCustomPtrOutputWithContext(context.Background())
}

func (i CustomImagePropertiesCustomArgs) ToCustomImagePropertiesCustomPtrOutputWithContext(ctx context.Context) CustomImagePropertiesCustomPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesCustomOutput).ToCustomImagePropertiesCustomPtrOutputWithContext(ctx)
}









type CustomImagePropertiesCustomPtrInput interface {
	pulumi.Input

	ToCustomImagePropertiesCustomPtrOutput() CustomImagePropertiesCustomPtrOutput
	ToCustomImagePropertiesCustomPtrOutputWithContext(context.Context) CustomImagePropertiesCustomPtrOutput
}

type customImagePropertiesCustomPtrType CustomImagePropertiesCustomArgs

func CustomImagePropertiesCustomPtr(v *CustomImagePropertiesCustomArgs) CustomImagePropertiesCustomPtrInput {
	return (*customImagePropertiesCustomPtrType)(v)
}

func (*customImagePropertiesCustomPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImagePropertiesCustom)(nil)).Elem()
}

func (i *customImagePropertiesCustomPtrType) ToCustomImagePropertiesCustomPtrOutput() CustomImagePropertiesCustomPtrOutput {
	return i.ToCustomImagePropertiesCustomPtrOutputWithContext(context.Background())
}

func (i *customImagePropertiesCustomPtrType) ToCustomImagePropertiesCustomPtrOutputWithContext(ctx context.Context) CustomImagePropertiesCustomPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesCustomPtrOutput)
}

type CustomImagePropertiesCustomOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesCustomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImagePropertiesCustom)(nil)).Elem()
}

func (o CustomImagePropertiesCustomOutput) ToCustomImagePropertiesCustomOutput() CustomImagePropertiesCustomOutput {
	return o
}

func (o CustomImagePropertiesCustomOutput) ToCustomImagePropertiesCustomOutputWithContext(ctx context.Context) CustomImagePropertiesCustomOutput {
	return o
}

func (o CustomImagePropertiesCustomOutput) ToCustomImagePropertiesCustomPtrOutput() CustomImagePropertiesCustomPtrOutput {
	return o.ToCustomImagePropertiesCustomPtrOutputWithContext(context.Background())
}

func (o CustomImagePropertiesCustomOutput) ToCustomImagePropertiesCustomPtrOutputWithContext(ctx context.Context) CustomImagePropertiesCustomPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomImagePropertiesCustom) *CustomImagePropertiesCustom {
		return &v
	}).(CustomImagePropertiesCustomPtrOutput)
}

func (o CustomImagePropertiesCustomOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesCustom) *string { return v.ImageName }).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesCustomOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v CustomImagePropertiesCustom) string { return v.OsType }).(pulumi.StringOutput)
}

func (o CustomImagePropertiesCustomOutput) SysPrep() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesCustom) *bool { return v.SysPrep }).(pulumi.BoolPtrOutput)
}

type CustomImagePropertiesCustomPtrOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesCustomPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImagePropertiesCustom)(nil)).Elem()
}

func (o CustomImagePropertiesCustomPtrOutput) ToCustomImagePropertiesCustomPtrOutput() CustomImagePropertiesCustomPtrOutput {
	return o
}

func (o CustomImagePropertiesCustomPtrOutput) ToCustomImagePropertiesCustomPtrOutputWithContext(ctx context.Context) CustomImagePropertiesCustomPtrOutput {
	return o
}

func (o CustomImagePropertiesCustomPtrOutput) Elem() CustomImagePropertiesCustomOutput {
	return o.ApplyT(func(v *CustomImagePropertiesCustom) CustomImagePropertiesCustom {
		if v != nil {
			return *v
		}
		var ret CustomImagePropertiesCustom
		return ret
	}).(CustomImagePropertiesCustomOutput)
}

func (o CustomImagePropertiesCustomPtrOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesCustom) *string {
		if v == nil {
			return nil
		}
		return v.ImageName
	}).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesCustomPtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesCustom) *string {
		if v == nil {
			return nil
		}
		return &v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesCustomPtrOutput) SysPrep() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesCustom) *bool {
		if v == nil {
			return nil
		}
		return v.SysPrep
	}).(pulumi.BoolPtrOutput)
}

type CustomImagePropertiesCustomResponse struct {
	ImageName *string `pulumi:"imageName"`
	OsType    string  `pulumi:"osType"`
	SysPrep   *bool   `pulumi:"sysPrep"`
}

type CustomImagePropertiesCustomResponseOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesCustomResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImagePropertiesCustomResponse)(nil)).Elem()
}

func (o CustomImagePropertiesCustomResponseOutput) ToCustomImagePropertiesCustomResponseOutput() CustomImagePropertiesCustomResponseOutput {
	return o
}

func (o CustomImagePropertiesCustomResponseOutput) ToCustomImagePropertiesCustomResponseOutputWithContext(ctx context.Context) CustomImagePropertiesCustomResponseOutput {
	return o
}

func (o CustomImagePropertiesCustomResponseOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesCustomResponse) *string { return v.ImageName }).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesCustomResponseOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v CustomImagePropertiesCustomResponse) string { return v.OsType }).(pulumi.StringOutput)
}

func (o CustomImagePropertiesCustomResponseOutput) SysPrep() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesCustomResponse) *bool { return v.SysPrep }).(pulumi.BoolPtrOutput)
}

type CustomImagePropertiesCustomResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesCustomResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImagePropertiesCustomResponse)(nil)).Elem()
}

func (o CustomImagePropertiesCustomResponsePtrOutput) ToCustomImagePropertiesCustomResponsePtrOutput() CustomImagePropertiesCustomResponsePtrOutput {
	return o
}

func (o CustomImagePropertiesCustomResponsePtrOutput) ToCustomImagePropertiesCustomResponsePtrOutputWithContext(ctx context.Context) CustomImagePropertiesCustomResponsePtrOutput {
	return o
}

func (o CustomImagePropertiesCustomResponsePtrOutput) Elem() CustomImagePropertiesCustomResponseOutput {
	return o.ApplyT(func(v *CustomImagePropertiesCustomResponse) CustomImagePropertiesCustomResponse {
		if v != nil {
			return *v
		}
		var ret CustomImagePropertiesCustomResponse
		return ret
	}).(CustomImagePropertiesCustomResponseOutput)
}

func (o CustomImagePropertiesCustomResponsePtrOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesCustomResponse) *string {
		if v == nil {
			return nil
		}
		return v.ImageName
	}).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesCustomResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesCustomResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesCustomResponsePtrOutput) SysPrep() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesCustomResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SysPrep
	}).(pulumi.BoolPtrOutput)
}

type CustomImagePropertiesFromPlan struct {
	Id        *string `pulumi:"id"`
	Offer     *string `pulumi:"offer"`
	Publisher *string `pulumi:"publisher"`
}





type CustomImagePropertiesFromPlanInput interface {
	pulumi.Input

	ToCustomImagePropertiesFromPlanOutput() CustomImagePropertiesFromPlanOutput
	ToCustomImagePropertiesFromPlanOutputWithContext(context.Context) CustomImagePropertiesFromPlanOutput
}

type CustomImagePropertiesFromPlanArgs struct {
	Id        pulumi.StringPtrInput `pulumi:"id"`
	Offer     pulumi.StringPtrInput `pulumi:"offer"`
	Publisher pulumi.StringPtrInput `pulumi:"publisher"`
}

func (CustomImagePropertiesFromPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImagePropertiesFromPlan)(nil)).Elem()
}

func (i CustomImagePropertiesFromPlanArgs) ToCustomImagePropertiesFromPlanOutput() CustomImagePropertiesFromPlanOutput {
	return i.ToCustomImagePropertiesFromPlanOutputWithContext(context.Background())
}

func (i CustomImagePropertiesFromPlanArgs) ToCustomImagePropertiesFromPlanOutputWithContext(ctx context.Context) CustomImagePropertiesFromPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesFromPlanOutput)
}

func (i CustomImagePropertiesFromPlanArgs) ToCustomImagePropertiesFromPlanPtrOutput() CustomImagePropertiesFromPlanPtrOutput {
	return i.ToCustomImagePropertiesFromPlanPtrOutputWithContext(context.Background())
}

func (i CustomImagePropertiesFromPlanArgs) ToCustomImagePropertiesFromPlanPtrOutputWithContext(ctx context.Context) CustomImagePropertiesFromPlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesFromPlanOutput).ToCustomImagePropertiesFromPlanPtrOutputWithContext(ctx)
}









type CustomImagePropertiesFromPlanPtrInput interface {
	pulumi.Input

	ToCustomImagePropertiesFromPlanPtrOutput() CustomImagePropertiesFromPlanPtrOutput
	ToCustomImagePropertiesFromPlanPtrOutputWithContext(context.Context) CustomImagePropertiesFromPlanPtrOutput
}

type customImagePropertiesFromPlanPtrType CustomImagePropertiesFromPlanArgs

func CustomImagePropertiesFromPlanPtr(v *CustomImagePropertiesFromPlanArgs) CustomImagePropertiesFromPlanPtrInput {
	return (*customImagePropertiesFromPlanPtrType)(v)
}

func (*customImagePropertiesFromPlanPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImagePropertiesFromPlan)(nil)).Elem()
}

func (i *customImagePropertiesFromPlanPtrType) ToCustomImagePropertiesFromPlanPtrOutput() CustomImagePropertiesFromPlanPtrOutput {
	return i.ToCustomImagePropertiesFromPlanPtrOutputWithContext(context.Background())
}

func (i *customImagePropertiesFromPlanPtrType) ToCustomImagePropertiesFromPlanPtrOutputWithContext(ctx context.Context) CustomImagePropertiesFromPlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesFromPlanPtrOutput)
}

type CustomImagePropertiesFromPlanOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesFromPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImagePropertiesFromPlan)(nil)).Elem()
}

func (o CustomImagePropertiesFromPlanOutput) ToCustomImagePropertiesFromPlanOutput() CustomImagePropertiesFromPlanOutput {
	return o
}

func (o CustomImagePropertiesFromPlanOutput) ToCustomImagePropertiesFromPlanOutputWithContext(ctx context.Context) CustomImagePropertiesFromPlanOutput {
	return o
}

func (o CustomImagePropertiesFromPlanOutput) ToCustomImagePropertiesFromPlanPtrOutput() CustomImagePropertiesFromPlanPtrOutput {
	return o.ToCustomImagePropertiesFromPlanPtrOutputWithContext(context.Background())
}

func (o CustomImagePropertiesFromPlanOutput) ToCustomImagePropertiesFromPlanPtrOutputWithContext(ctx context.Context) CustomImagePropertiesFromPlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomImagePropertiesFromPlan) *CustomImagePropertiesFromPlan {
		return &v
	}).(CustomImagePropertiesFromPlanPtrOutput)
}

func (o CustomImagePropertiesFromPlanOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromPlan) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesFromPlanOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromPlan) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesFromPlanOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromPlan) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type CustomImagePropertiesFromPlanPtrOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesFromPlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImagePropertiesFromPlan)(nil)).Elem()
}

func (o CustomImagePropertiesFromPlanPtrOutput) ToCustomImagePropertiesFromPlanPtrOutput() CustomImagePropertiesFromPlanPtrOutput {
	return o
}

func (o CustomImagePropertiesFromPlanPtrOutput) ToCustomImagePropertiesFromPlanPtrOutputWithContext(ctx context.Context) CustomImagePropertiesFromPlanPtrOutput {
	return o
}

func (o CustomImagePropertiesFromPlanPtrOutput) Elem() CustomImagePropertiesFromPlanOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromPlan) CustomImagePropertiesFromPlan {
		if v != nil {
			return *v
		}
		var ret CustomImagePropertiesFromPlan
		return ret
	}).(CustomImagePropertiesFromPlanOutput)
}

func (o CustomImagePropertiesFromPlanPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromPlan) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesFromPlanPtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromPlan) *string {
		if v == nil {
			return nil
		}
		return v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesFromPlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromPlan) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

type CustomImagePropertiesFromPlanResponse struct {
	Id        *string `pulumi:"id"`
	Offer     *string `pulumi:"offer"`
	Publisher *string `pulumi:"publisher"`
}

type CustomImagePropertiesFromPlanResponseOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesFromPlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImagePropertiesFromPlanResponse)(nil)).Elem()
}

func (o CustomImagePropertiesFromPlanResponseOutput) ToCustomImagePropertiesFromPlanResponseOutput() CustomImagePropertiesFromPlanResponseOutput {
	return o
}

func (o CustomImagePropertiesFromPlanResponseOutput) ToCustomImagePropertiesFromPlanResponseOutputWithContext(ctx context.Context) CustomImagePropertiesFromPlanResponseOutput {
	return o
}

func (o CustomImagePropertiesFromPlanResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromPlanResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesFromPlanResponseOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromPlanResponse) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesFromPlanResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromPlanResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type CustomImagePropertiesFromPlanResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesFromPlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImagePropertiesFromPlanResponse)(nil)).Elem()
}

func (o CustomImagePropertiesFromPlanResponsePtrOutput) ToCustomImagePropertiesFromPlanResponsePtrOutput() CustomImagePropertiesFromPlanResponsePtrOutput {
	return o
}

func (o CustomImagePropertiesFromPlanResponsePtrOutput) ToCustomImagePropertiesFromPlanResponsePtrOutputWithContext(ctx context.Context) CustomImagePropertiesFromPlanResponsePtrOutput {
	return o
}

func (o CustomImagePropertiesFromPlanResponsePtrOutput) Elem() CustomImagePropertiesFromPlanResponseOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromPlanResponse) CustomImagePropertiesFromPlanResponse {
		if v != nil {
			return *v
		}
		var ret CustomImagePropertiesFromPlanResponse
		return ret
	}).(CustomImagePropertiesFromPlanResponseOutput)
}

func (o CustomImagePropertiesFromPlanResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesFromPlanResponsePtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesFromPlanResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

type CustomImagePropertiesFromVm struct {
	LinuxOsInfo   *LinuxOsInfo   `pulumi:"linuxOsInfo"`
	SourceVmId    *string        `pulumi:"sourceVmId"`
	WindowsOsInfo *WindowsOsInfo `pulumi:"windowsOsInfo"`
}





type CustomImagePropertiesFromVmInput interface {
	pulumi.Input

	ToCustomImagePropertiesFromVmOutput() CustomImagePropertiesFromVmOutput
	ToCustomImagePropertiesFromVmOutputWithContext(context.Context) CustomImagePropertiesFromVmOutput
}

type CustomImagePropertiesFromVmArgs struct {
	LinuxOsInfo   LinuxOsInfoPtrInput   `pulumi:"linuxOsInfo"`
	SourceVmId    pulumi.StringPtrInput `pulumi:"sourceVmId"`
	WindowsOsInfo WindowsOsInfoPtrInput `pulumi:"windowsOsInfo"`
}

func (CustomImagePropertiesFromVmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImagePropertiesFromVm)(nil)).Elem()
}

func (i CustomImagePropertiesFromVmArgs) ToCustomImagePropertiesFromVmOutput() CustomImagePropertiesFromVmOutput {
	return i.ToCustomImagePropertiesFromVmOutputWithContext(context.Background())
}

func (i CustomImagePropertiesFromVmArgs) ToCustomImagePropertiesFromVmOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesFromVmOutput)
}

func (i CustomImagePropertiesFromVmArgs) ToCustomImagePropertiesFromVmPtrOutput() CustomImagePropertiesFromVmPtrOutput {
	return i.ToCustomImagePropertiesFromVmPtrOutputWithContext(context.Background())
}

func (i CustomImagePropertiesFromVmArgs) ToCustomImagePropertiesFromVmPtrOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesFromVmOutput).ToCustomImagePropertiesFromVmPtrOutputWithContext(ctx)
}









type CustomImagePropertiesFromVmPtrInput interface {
	pulumi.Input

	ToCustomImagePropertiesFromVmPtrOutput() CustomImagePropertiesFromVmPtrOutput
	ToCustomImagePropertiesFromVmPtrOutputWithContext(context.Context) CustomImagePropertiesFromVmPtrOutput
}

type customImagePropertiesFromVmPtrType CustomImagePropertiesFromVmArgs

func CustomImagePropertiesFromVmPtr(v *CustomImagePropertiesFromVmArgs) CustomImagePropertiesFromVmPtrInput {
	return (*customImagePropertiesFromVmPtrType)(v)
}

func (*customImagePropertiesFromVmPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImagePropertiesFromVm)(nil)).Elem()
}

func (i *customImagePropertiesFromVmPtrType) ToCustomImagePropertiesFromVmPtrOutput() CustomImagePropertiesFromVmPtrOutput {
	return i.ToCustomImagePropertiesFromVmPtrOutputWithContext(context.Background())
}

func (i *customImagePropertiesFromVmPtrType) ToCustomImagePropertiesFromVmPtrOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomImagePropertiesFromVmPtrOutput)
}

type CustomImagePropertiesFromVmOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesFromVmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImagePropertiesFromVm)(nil)).Elem()
}

func (o CustomImagePropertiesFromVmOutput) ToCustomImagePropertiesFromVmOutput() CustomImagePropertiesFromVmOutput {
	return o
}

func (o CustomImagePropertiesFromVmOutput) ToCustomImagePropertiesFromVmOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmOutput {
	return o
}

func (o CustomImagePropertiesFromVmOutput) ToCustomImagePropertiesFromVmPtrOutput() CustomImagePropertiesFromVmPtrOutput {
	return o.ToCustomImagePropertiesFromVmPtrOutputWithContext(context.Background())
}

func (o CustomImagePropertiesFromVmOutput) ToCustomImagePropertiesFromVmPtrOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomImagePropertiesFromVm) *CustomImagePropertiesFromVm {
		return &v
	}).(CustomImagePropertiesFromVmPtrOutput)
}

func (o CustomImagePropertiesFromVmOutput) LinuxOsInfo() LinuxOsInfoPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromVm) *LinuxOsInfo { return v.LinuxOsInfo }).(LinuxOsInfoPtrOutput)
}

func (o CustomImagePropertiesFromVmOutput) SourceVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromVm) *string { return v.SourceVmId }).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesFromVmOutput) WindowsOsInfo() WindowsOsInfoPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromVm) *WindowsOsInfo { return v.WindowsOsInfo }).(WindowsOsInfoPtrOutput)
}

type CustomImagePropertiesFromVmPtrOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesFromVmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImagePropertiesFromVm)(nil)).Elem()
}

func (o CustomImagePropertiesFromVmPtrOutput) ToCustomImagePropertiesFromVmPtrOutput() CustomImagePropertiesFromVmPtrOutput {
	return o
}

func (o CustomImagePropertiesFromVmPtrOutput) ToCustomImagePropertiesFromVmPtrOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmPtrOutput {
	return o
}

func (o CustomImagePropertiesFromVmPtrOutput) Elem() CustomImagePropertiesFromVmOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromVm) CustomImagePropertiesFromVm {
		if v != nil {
			return *v
		}
		var ret CustomImagePropertiesFromVm
		return ret
	}).(CustomImagePropertiesFromVmOutput)
}

func (o CustomImagePropertiesFromVmPtrOutput) LinuxOsInfo() LinuxOsInfoPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromVm) *LinuxOsInfo {
		if v == nil {
			return nil
		}
		return v.LinuxOsInfo
	}).(LinuxOsInfoPtrOutput)
}

func (o CustomImagePropertiesFromVmPtrOutput) SourceVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromVm) *string {
		if v == nil {
			return nil
		}
		return v.SourceVmId
	}).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesFromVmPtrOutput) WindowsOsInfo() WindowsOsInfoPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromVm) *WindowsOsInfo {
		if v == nil {
			return nil
		}
		return v.WindowsOsInfo
	}).(WindowsOsInfoPtrOutput)
}

type CustomImagePropertiesFromVmResponse struct {
	LinuxOsInfo   *LinuxOsInfoResponse   `pulumi:"linuxOsInfo"`
	SourceVmId    *string                `pulumi:"sourceVmId"`
	WindowsOsInfo *WindowsOsInfoResponse `pulumi:"windowsOsInfo"`
}

type CustomImagePropertiesFromVmResponseOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesFromVmResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomImagePropertiesFromVmResponse)(nil)).Elem()
}

func (o CustomImagePropertiesFromVmResponseOutput) ToCustomImagePropertiesFromVmResponseOutput() CustomImagePropertiesFromVmResponseOutput {
	return o
}

func (o CustomImagePropertiesFromVmResponseOutput) ToCustomImagePropertiesFromVmResponseOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmResponseOutput {
	return o
}

func (o CustomImagePropertiesFromVmResponseOutput) LinuxOsInfo() LinuxOsInfoResponsePtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromVmResponse) *LinuxOsInfoResponse { return v.LinuxOsInfo }).(LinuxOsInfoResponsePtrOutput)
}

func (o CustomImagePropertiesFromVmResponseOutput) SourceVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromVmResponse) *string { return v.SourceVmId }).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesFromVmResponseOutput) WindowsOsInfo() WindowsOsInfoResponsePtrOutput {
	return o.ApplyT(func(v CustomImagePropertiesFromVmResponse) *WindowsOsInfoResponse { return v.WindowsOsInfo }).(WindowsOsInfoResponsePtrOutput)
}

type CustomImagePropertiesFromVmResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomImagePropertiesFromVmResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomImagePropertiesFromVmResponse)(nil)).Elem()
}

func (o CustomImagePropertiesFromVmResponsePtrOutput) ToCustomImagePropertiesFromVmResponsePtrOutput() CustomImagePropertiesFromVmResponsePtrOutput {
	return o
}

func (o CustomImagePropertiesFromVmResponsePtrOutput) ToCustomImagePropertiesFromVmResponsePtrOutputWithContext(ctx context.Context) CustomImagePropertiesFromVmResponsePtrOutput {
	return o
}

func (o CustomImagePropertiesFromVmResponsePtrOutput) Elem() CustomImagePropertiesFromVmResponseOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromVmResponse) CustomImagePropertiesFromVmResponse {
		if v != nil {
			return *v
		}
		var ret CustomImagePropertiesFromVmResponse
		return ret
	}).(CustomImagePropertiesFromVmResponseOutput)
}

func (o CustomImagePropertiesFromVmResponsePtrOutput) LinuxOsInfo() LinuxOsInfoResponsePtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromVmResponse) *LinuxOsInfoResponse {
		if v == nil {
			return nil
		}
		return v.LinuxOsInfo
	}).(LinuxOsInfoResponsePtrOutput)
}

func (o CustomImagePropertiesFromVmResponsePtrOutput) SourceVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromVmResponse) *string {
		if v == nil {
			return nil
		}
		return v.SourceVmId
	}).(pulumi.StringPtrOutput)
}

func (o CustomImagePropertiesFromVmResponsePtrOutput) WindowsOsInfo() WindowsOsInfoResponsePtrOutput {
	return o.ApplyT(func(v *CustomImagePropertiesFromVmResponse) *WindowsOsInfoResponse {
		if v == nil {
			return nil
		}
		return v.WindowsOsInfo
	}).(WindowsOsInfoResponsePtrOutput)
}

type DataDiskProperties struct {
	AttachNewDataDiskOptions *AttachNewDataDiskOptions `pulumi:"attachNewDataDiskOptions"`
	ExistingLabDiskId        *string                   `pulumi:"existingLabDiskId"`
	HostCaching              *string                   `pulumi:"hostCaching"`
}





type DataDiskPropertiesInput interface {
	pulumi.Input

	ToDataDiskPropertiesOutput() DataDiskPropertiesOutput
	ToDataDiskPropertiesOutputWithContext(context.Context) DataDiskPropertiesOutput
}

type DataDiskPropertiesArgs struct {
	AttachNewDataDiskOptions AttachNewDataDiskOptionsPtrInput `pulumi:"attachNewDataDiskOptions"`
	ExistingLabDiskId        pulumi.StringPtrInput            `pulumi:"existingLabDiskId"`
	HostCaching              pulumi.StringPtrInput            `pulumi:"hostCaching"`
}

func (DataDiskPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDiskProperties)(nil)).Elem()
}

func (i DataDiskPropertiesArgs) ToDataDiskPropertiesOutput() DataDiskPropertiesOutput {
	return i.ToDataDiskPropertiesOutputWithContext(context.Background())
}

func (i DataDiskPropertiesArgs) ToDataDiskPropertiesOutputWithContext(ctx context.Context) DataDiskPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskPropertiesOutput)
}





type DataDiskPropertiesArrayInput interface {
	pulumi.Input

	ToDataDiskPropertiesArrayOutput() DataDiskPropertiesArrayOutput
	ToDataDiskPropertiesArrayOutputWithContext(context.Context) DataDiskPropertiesArrayOutput
}

type DataDiskPropertiesArray []DataDiskPropertiesInput

func (DataDiskPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDiskProperties)(nil)).Elem()
}

func (i DataDiskPropertiesArray) ToDataDiskPropertiesArrayOutput() DataDiskPropertiesArrayOutput {
	return i.ToDataDiskPropertiesArrayOutputWithContext(context.Background())
}

func (i DataDiskPropertiesArray) ToDataDiskPropertiesArrayOutputWithContext(ctx context.Context) DataDiskPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskPropertiesArrayOutput)
}

type DataDiskPropertiesOutput struct{ *pulumi.OutputState }

func (DataDiskPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDiskProperties)(nil)).Elem()
}

func (o DataDiskPropertiesOutput) ToDataDiskPropertiesOutput() DataDiskPropertiesOutput {
	return o
}

func (o DataDiskPropertiesOutput) ToDataDiskPropertiesOutputWithContext(ctx context.Context) DataDiskPropertiesOutput {
	return o
}

func (o DataDiskPropertiesOutput) AttachNewDataDiskOptions() AttachNewDataDiskOptionsPtrOutput {
	return o.ApplyT(func(v DataDiskProperties) *AttachNewDataDiskOptions { return v.AttachNewDataDiskOptions }).(AttachNewDataDiskOptionsPtrOutput)
}

func (o DataDiskPropertiesOutput) ExistingLabDiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskProperties) *string { return v.ExistingLabDiskId }).(pulumi.StringPtrOutput)
}

func (o DataDiskPropertiesOutput) HostCaching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskProperties) *string { return v.HostCaching }).(pulumi.StringPtrOutput)
}

type DataDiskPropertiesArrayOutput struct{ *pulumi.OutputState }

func (DataDiskPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDiskProperties)(nil)).Elem()
}

func (o DataDiskPropertiesArrayOutput) ToDataDiskPropertiesArrayOutput() DataDiskPropertiesArrayOutput {
	return o
}

func (o DataDiskPropertiesArrayOutput) ToDataDiskPropertiesArrayOutputWithContext(ctx context.Context) DataDiskPropertiesArrayOutput {
	return o
}

func (o DataDiskPropertiesArrayOutput) Index(i pulumi.IntInput) DataDiskPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDiskProperties {
		return vs[0].([]DataDiskProperties)[vs[1].(int)]
	}).(DataDiskPropertiesOutput)
}

type DataDiskPropertiesResponse struct {
	AttachNewDataDiskOptions *AttachNewDataDiskOptionsResponse `pulumi:"attachNewDataDiskOptions"`
	ExistingLabDiskId        *string                           `pulumi:"existingLabDiskId"`
	HostCaching              *string                           `pulumi:"hostCaching"`
}

type DataDiskPropertiesResponseOutput struct{ *pulumi.OutputState }

func (DataDiskPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDiskPropertiesResponse)(nil)).Elem()
}

func (o DataDiskPropertiesResponseOutput) ToDataDiskPropertiesResponseOutput() DataDiskPropertiesResponseOutput {
	return o
}

func (o DataDiskPropertiesResponseOutput) ToDataDiskPropertiesResponseOutputWithContext(ctx context.Context) DataDiskPropertiesResponseOutput {
	return o
}

func (o DataDiskPropertiesResponseOutput) AttachNewDataDiskOptions() AttachNewDataDiskOptionsResponsePtrOutput {
	return o.ApplyT(func(v DataDiskPropertiesResponse) *AttachNewDataDiskOptionsResponse {
		return v.AttachNewDataDiskOptions
	}).(AttachNewDataDiskOptionsResponsePtrOutput)
}

func (o DataDiskPropertiesResponseOutput) ExistingLabDiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskPropertiesResponse) *string { return v.ExistingLabDiskId }).(pulumi.StringPtrOutput)
}

func (o DataDiskPropertiesResponseOutput) HostCaching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskPropertiesResponse) *string { return v.HostCaching }).(pulumi.StringPtrOutput)
}

type DataDiskPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (DataDiskPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDiskPropertiesResponse)(nil)).Elem()
}

func (o DataDiskPropertiesResponseArrayOutput) ToDataDiskPropertiesResponseArrayOutput() DataDiskPropertiesResponseArrayOutput {
	return o
}

func (o DataDiskPropertiesResponseArrayOutput) ToDataDiskPropertiesResponseArrayOutputWithContext(ctx context.Context) DataDiskPropertiesResponseArrayOutput {
	return o
}

func (o DataDiskPropertiesResponseArrayOutput) Index(i pulumi.IntInput) DataDiskPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDiskPropertiesResponse {
		return vs[0].([]DataDiskPropertiesResponse)[vs[1].(int)]
	}).(DataDiskPropertiesResponseOutput)
}

type DataDiskStorageTypeInfo struct {
	Lun         *string `pulumi:"lun"`
	StorageType *string `pulumi:"storageType"`
}





type DataDiskStorageTypeInfoInput interface {
	pulumi.Input

	ToDataDiskStorageTypeInfoOutput() DataDiskStorageTypeInfoOutput
	ToDataDiskStorageTypeInfoOutputWithContext(context.Context) DataDiskStorageTypeInfoOutput
}

type DataDiskStorageTypeInfoArgs struct {
	Lun         pulumi.StringPtrInput `pulumi:"lun"`
	StorageType pulumi.StringPtrInput `pulumi:"storageType"`
}

func (DataDiskStorageTypeInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDiskStorageTypeInfo)(nil)).Elem()
}

func (i DataDiskStorageTypeInfoArgs) ToDataDiskStorageTypeInfoOutput() DataDiskStorageTypeInfoOutput {
	return i.ToDataDiskStorageTypeInfoOutputWithContext(context.Background())
}

func (i DataDiskStorageTypeInfoArgs) ToDataDiskStorageTypeInfoOutputWithContext(ctx context.Context) DataDiskStorageTypeInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskStorageTypeInfoOutput)
}





type DataDiskStorageTypeInfoArrayInput interface {
	pulumi.Input

	ToDataDiskStorageTypeInfoArrayOutput() DataDiskStorageTypeInfoArrayOutput
	ToDataDiskStorageTypeInfoArrayOutputWithContext(context.Context) DataDiskStorageTypeInfoArrayOutput
}

type DataDiskStorageTypeInfoArray []DataDiskStorageTypeInfoInput

func (DataDiskStorageTypeInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDiskStorageTypeInfo)(nil)).Elem()
}

func (i DataDiskStorageTypeInfoArray) ToDataDiskStorageTypeInfoArrayOutput() DataDiskStorageTypeInfoArrayOutput {
	return i.ToDataDiskStorageTypeInfoArrayOutputWithContext(context.Background())
}

func (i DataDiskStorageTypeInfoArray) ToDataDiskStorageTypeInfoArrayOutputWithContext(ctx context.Context) DataDiskStorageTypeInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskStorageTypeInfoArrayOutput)
}

type DataDiskStorageTypeInfoOutput struct{ *pulumi.OutputState }

func (DataDiskStorageTypeInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDiskStorageTypeInfo)(nil)).Elem()
}

func (o DataDiskStorageTypeInfoOutput) ToDataDiskStorageTypeInfoOutput() DataDiskStorageTypeInfoOutput {
	return o
}

func (o DataDiskStorageTypeInfoOutput) ToDataDiskStorageTypeInfoOutputWithContext(ctx context.Context) DataDiskStorageTypeInfoOutput {
	return o
}

func (o DataDiskStorageTypeInfoOutput) Lun() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskStorageTypeInfo) *string { return v.Lun }).(pulumi.StringPtrOutput)
}

func (o DataDiskStorageTypeInfoOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskStorageTypeInfo) *string { return v.StorageType }).(pulumi.StringPtrOutput)
}

type DataDiskStorageTypeInfoArrayOutput struct{ *pulumi.OutputState }

func (DataDiskStorageTypeInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDiskStorageTypeInfo)(nil)).Elem()
}

func (o DataDiskStorageTypeInfoArrayOutput) ToDataDiskStorageTypeInfoArrayOutput() DataDiskStorageTypeInfoArrayOutput {
	return o
}

func (o DataDiskStorageTypeInfoArrayOutput) ToDataDiskStorageTypeInfoArrayOutputWithContext(ctx context.Context) DataDiskStorageTypeInfoArrayOutput {
	return o
}

func (o DataDiskStorageTypeInfoArrayOutput) Index(i pulumi.IntInput) DataDiskStorageTypeInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDiskStorageTypeInfo {
		return vs[0].([]DataDiskStorageTypeInfo)[vs[1].(int)]
	}).(DataDiskStorageTypeInfoOutput)
}

type DataDiskStorageTypeInfoResponse struct {
	Lun         *string `pulumi:"lun"`
	StorageType *string `pulumi:"storageType"`
}

type DataDiskStorageTypeInfoResponseOutput struct{ *pulumi.OutputState }

func (DataDiskStorageTypeInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDiskStorageTypeInfoResponse)(nil)).Elem()
}

func (o DataDiskStorageTypeInfoResponseOutput) ToDataDiskStorageTypeInfoResponseOutput() DataDiskStorageTypeInfoResponseOutput {
	return o
}

func (o DataDiskStorageTypeInfoResponseOutput) ToDataDiskStorageTypeInfoResponseOutputWithContext(ctx context.Context) DataDiskStorageTypeInfoResponseOutput {
	return o
}

func (o DataDiskStorageTypeInfoResponseOutput) Lun() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskStorageTypeInfoResponse) *string { return v.Lun }).(pulumi.StringPtrOutput)
}

func (o DataDiskStorageTypeInfoResponseOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskStorageTypeInfoResponse) *string { return v.StorageType }).(pulumi.StringPtrOutput)
}

type DataDiskStorageTypeInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (DataDiskStorageTypeInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDiskStorageTypeInfoResponse)(nil)).Elem()
}

func (o DataDiskStorageTypeInfoResponseArrayOutput) ToDataDiskStorageTypeInfoResponseArrayOutput() DataDiskStorageTypeInfoResponseArrayOutput {
	return o
}

func (o DataDiskStorageTypeInfoResponseArrayOutput) ToDataDiskStorageTypeInfoResponseArrayOutputWithContext(ctx context.Context) DataDiskStorageTypeInfoResponseArrayOutput {
	return o
}

func (o DataDiskStorageTypeInfoResponseArrayOutput) Index(i pulumi.IntInput) DataDiskStorageTypeInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDiskStorageTypeInfoResponse {
		return vs[0].([]DataDiskStorageTypeInfoResponse)[vs[1].(int)]
	}).(DataDiskStorageTypeInfoResponseOutput)
}

type DayDetails struct {
	Time *string `pulumi:"time"`
}





type DayDetailsInput interface {
	pulumi.Input

	ToDayDetailsOutput() DayDetailsOutput
	ToDayDetailsOutputWithContext(context.Context) DayDetailsOutput
}

type DayDetailsArgs struct {
	Time pulumi.StringPtrInput `pulumi:"time"`
}

func (DayDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DayDetails)(nil)).Elem()
}

func (i DayDetailsArgs) ToDayDetailsOutput() DayDetailsOutput {
	return i.ToDayDetailsOutputWithContext(context.Background())
}

func (i DayDetailsArgs) ToDayDetailsOutputWithContext(ctx context.Context) DayDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayDetailsOutput)
}

func (i DayDetailsArgs) ToDayDetailsPtrOutput() DayDetailsPtrOutput {
	return i.ToDayDetailsPtrOutputWithContext(context.Background())
}

func (i DayDetailsArgs) ToDayDetailsPtrOutputWithContext(ctx context.Context) DayDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayDetailsOutput).ToDayDetailsPtrOutputWithContext(ctx)
}









type DayDetailsPtrInput interface {
	pulumi.Input

	ToDayDetailsPtrOutput() DayDetailsPtrOutput
	ToDayDetailsPtrOutputWithContext(context.Context) DayDetailsPtrOutput
}

type dayDetailsPtrType DayDetailsArgs

func DayDetailsPtr(v *DayDetailsArgs) DayDetailsPtrInput {
	return (*dayDetailsPtrType)(v)
}

func (*dayDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DayDetails)(nil)).Elem()
}

func (i *dayDetailsPtrType) ToDayDetailsPtrOutput() DayDetailsPtrOutput {
	return i.ToDayDetailsPtrOutputWithContext(context.Background())
}

func (i *dayDetailsPtrType) ToDayDetailsPtrOutputWithContext(ctx context.Context) DayDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayDetailsPtrOutput)
}

type DayDetailsOutput struct{ *pulumi.OutputState }

func (DayDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DayDetails)(nil)).Elem()
}

func (o DayDetailsOutput) ToDayDetailsOutput() DayDetailsOutput {
	return o
}

func (o DayDetailsOutput) ToDayDetailsOutputWithContext(ctx context.Context) DayDetailsOutput {
	return o
}

func (o DayDetailsOutput) ToDayDetailsPtrOutput() DayDetailsPtrOutput {
	return o.ToDayDetailsPtrOutputWithContext(context.Background())
}

func (o DayDetailsOutput) ToDayDetailsPtrOutputWithContext(ctx context.Context) DayDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DayDetails) *DayDetails {
		return &v
	}).(DayDetailsPtrOutput)
}

func (o DayDetailsOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DayDetails) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type DayDetailsPtrOutput struct{ *pulumi.OutputState }

func (DayDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DayDetails)(nil)).Elem()
}

func (o DayDetailsPtrOutput) ToDayDetailsPtrOutput() DayDetailsPtrOutput {
	return o
}

func (o DayDetailsPtrOutput) ToDayDetailsPtrOutputWithContext(ctx context.Context) DayDetailsPtrOutput {
	return o
}

func (o DayDetailsPtrOutput) Elem() DayDetailsOutput {
	return o.ApplyT(func(v *DayDetails) DayDetails {
		if v != nil {
			return *v
		}
		var ret DayDetails
		return ret
	}).(DayDetailsOutput)
}

func (o DayDetailsPtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DayDetails) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

type DayDetailsResponse struct {
	Time *string `pulumi:"time"`
}

type DayDetailsResponseOutput struct{ *pulumi.OutputState }

func (DayDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DayDetailsResponse)(nil)).Elem()
}

func (o DayDetailsResponseOutput) ToDayDetailsResponseOutput() DayDetailsResponseOutput {
	return o
}

func (o DayDetailsResponseOutput) ToDayDetailsResponseOutputWithContext(ctx context.Context) DayDetailsResponseOutput {
	return o
}

func (o DayDetailsResponseOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DayDetailsResponse) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type DayDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (DayDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DayDetailsResponse)(nil)).Elem()
}

func (o DayDetailsResponsePtrOutput) ToDayDetailsResponsePtrOutput() DayDetailsResponsePtrOutput {
	return o
}

func (o DayDetailsResponsePtrOutput) ToDayDetailsResponsePtrOutputWithContext(ctx context.Context) DayDetailsResponsePtrOutput {
	return o
}

func (o DayDetailsResponsePtrOutput) Elem() DayDetailsResponseOutput {
	return o.ApplyT(func(v *DayDetailsResponse) DayDetailsResponse {
		if v != nil {
			return *v
		}
		var ret DayDetailsResponse
		return ret
	}).(DayDetailsResponseOutput)
}

func (o DayDetailsResponsePtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DayDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

type EnvironmentDeploymentProperties struct {
	ArmTemplateId *string                          `pulumi:"armTemplateId"`
	Parameters    []ArmTemplateParameterProperties `pulumi:"parameters"`
}





type EnvironmentDeploymentPropertiesInput interface {
	pulumi.Input

	ToEnvironmentDeploymentPropertiesOutput() EnvironmentDeploymentPropertiesOutput
	ToEnvironmentDeploymentPropertiesOutputWithContext(context.Context) EnvironmentDeploymentPropertiesOutput
}

type EnvironmentDeploymentPropertiesArgs struct {
	ArmTemplateId pulumi.StringPtrInput                    `pulumi:"armTemplateId"`
	Parameters    ArmTemplateParameterPropertiesArrayInput `pulumi:"parameters"`
}

func (EnvironmentDeploymentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentDeploymentProperties)(nil)).Elem()
}

func (i EnvironmentDeploymentPropertiesArgs) ToEnvironmentDeploymentPropertiesOutput() EnvironmentDeploymentPropertiesOutput {
	return i.ToEnvironmentDeploymentPropertiesOutputWithContext(context.Background())
}

func (i EnvironmentDeploymentPropertiesArgs) ToEnvironmentDeploymentPropertiesOutputWithContext(ctx context.Context) EnvironmentDeploymentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentDeploymentPropertiesOutput)
}

func (i EnvironmentDeploymentPropertiesArgs) ToEnvironmentDeploymentPropertiesPtrOutput() EnvironmentDeploymentPropertiesPtrOutput {
	return i.ToEnvironmentDeploymentPropertiesPtrOutputWithContext(context.Background())
}

func (i EnvironmentDeploymentPropertiesArgs) ToEnvironmentDeploymentPropertiesPtrOutputWithContext(ctx context.Context) EnvironmentDeploymentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentDeploymentPropertiesOutput).ToEnvironmentDeploymentPropertiesPtrOutputWithContext(ctx)
}









type EnvironmentDeploymentPropertiesPtrInput interface {
	pulumi.Input

	ToEnvironmentDeploymentPropertiesPtrOutput() EnvironmentDeploymentPropertiesPtrOutput
	ToEnvironmentDeploymentPropertiesPtrOutputWithContext(context.Context) EnvironmentDeploymentPropertiesPtrOutput
}

type environmentDeploymentPropertiesPtrType EnvironmentDeploymentPropertiesArgs

func EnvironmentDeploymentPropertiesPtr(v *EnvironmentDeploymentPropertiesArgs) EnvironmentDeploymentPropertiesPtrInput {
	return (*environmentDeploymentPropertiesPtrType)(v)
}

func (*environmentDeploymentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentDeploymentProperties)(nil)).Elem()
}

func (i *environmentDeploymentPropertiesPtrType) ToEnvironmentDeploymentPropertiesPtrOutput() EnvironmentDeploymentPropertiesPtrOutput {
	return i.ToEnvironmentDeploymentPropertiesPtrOutputWithContext(context.Background())
}

func (i *environmentDeploymentPropertiesPtrType) ToEnvironmentDeploymentPropertiesPtrOutputWithContext(ctx context.Context) EnvironmentDeploymentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentDeploymentPropertiesPtrOutput)
}

type EnvironmentDeploymentPropertiesOutput struct{ *pulumi.OutputState }

func (EnvironmentDeploymentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentDeploymentProperties)(nil)).Elem()
}

func (o EnvironmentDeploymentPropertiesOutput) ToEnvironmentDeploymentPropertiesOutput() EnvironmentDeploymentPropertiesOutput {
	return o
}

func (o EnvironmentDeploymentPropertiesOutput) ToEnvironmentDeploymentPropertiesOutputWithContext(ctx context.Context) EnvironmentDeploymentPropertiesOutput {
	return o
}

func (o EnvironmentDeploymentPropertiesOutput) ToEnvironmentDeploymentPropertiesPtrOutput() EnvironmentDeploymentPropertiesPtrOutput {
	return o.ToEnvironmentDeploymentPropertiesPtrOutputWithContext(context.Background())
}

func (o EnvironmentDeploymentPropertiesOutput) ToEnvironmentDeploymentPropertiesPtrOutputWithContext(ctx context.Context) EnvironmentDeploymentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentDeploymentProperties) *EnvironmentDeploymentProperties {
		return &v
	}).(EnvironmentDeploymentPropertiesPtrOutput)
}

func (o EnvironmentDeploymentPropertiesOutput) ArmTemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentDeploymentProperties) *string { return v.ArmTemplateId }).(pulumi.StringPtrOutput)
}

func (o EnvironmentDeploymentPropertiesOutput) Parameters() ArmTemplateParameterPropertiesArrayOutput {
	return o.ApplyT(func(v EnvironmentDeploymentProperties) []ArmTemplateParameterProperties { return v.Parameters }).(ArmTemplateParameterPropertiesArrayOutput)
}

type EnvironmentDeploymentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EnvironmentDeploymentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentDeploymentProperties)(nil)).Elem()
}

func (o EnvironmentDeploymentPropertiesPtrOutput) ToEnvironmentDeploymentPropertiesPtrOutput() EnvironmentDeploymentPropertiesPtrOutput {
	return o
}

func (o EnvironmentDeploymentPropertiesPtrOutput) ToEnvironmentDeploymentPropertiesPtrOutputWithContext(ctx context.Context) EnvironmentDeploymentPropertiesPtrOutput {
	return o
}

func (o EnvironmentDeploymentPropertiesPtrOutput) Elem() EnvironmentDeploymentPropertiesOutput {
	return o.ApplyT(func(v *EnvironmentDeploymentProperties) EnvironmentDeploymentProperties {
		if v != nil {
			return *v
		}
		var ret EnvironmentDeploymentProperties
		return ret
	}).(EnvironmentDeploymentPropertiesOutput)
}

func (o EnvironmentDeploymentPropertiesPtrOutput) ArmTemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentDeploymentProperties) *string {
		if v == nil {
			return nil
		}
		return v.ArmTemplateId
	}).(pulumi.StringPtrOutput)
}

func (o EnvironmentDeploymentPropertiesPtrOutput) Parameters() ArmTemplateParameterPropertiesArrayOutput {
	return o.ApplyT(func(v *EnvironmentDeploymentProperties) []ArmTemplateParameterProperties {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(ArmTemplateParameterPropertiesArrayOutput)
}

type EnvironmentDeploymentPropertiesResponse struct {
	ArmTemplateId *string                                  `pulumi:"armTemplateId"`
	Parameters    []ArmTemplateParameterPropertiesResponse `pulumi:"parameters"`
}

type EnvironmentDeploymentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EnvironmentDeploymentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentDeploymentPropertiesResponse)(nil)).Elem()
}

func (o EnvironmentDeploymentPropertiesResponseOutput) ToEnvironmentDeploymentPropertiesResponseOutput() EnvironmentDeploymentPropertiesResponseOutput {
	return o
}

func (o EnvironmentDeploymentPropertiesResponseOutput) ToEnvironmentDeploymentPropertiesResponseOutputWithContext(ctx context.Context) EnvironmentDeploymentPropertiesResponseOutput {
	return o
}

func (o EnvironmentDeploymentPropertiesResponseOutput) ArmTemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentDeploymentPropertiesResponse) *string { return v.ArmTemplateId }).(pulumi.StringPtrOutput)
}

func (o EnvironmentDeploymentPropertiesResponseOutput) Parameters() ArmTemplateParameterPropertiesResponseArrayOutput {
	return o.ApplyT(func(v EnvironmentDeploymentPropertiesResponse) []ArmTemplateParameterPropertiesResponse {
		return v.Parameters
	}).(ArmTemplateParameterPropertiesResponseArrayOutput)
}

type EnvironmentDeploymentPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (EnvironmentDeploymentPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentDeploymentPropertiesResponse)(nil)).Elem()
}

func (o EnvironmentDeploymentPropertiesResponsePtrOutput) ToEnvironmentDeploymentPropertiesResponsePtrOutput() EnvironmentDeploymentPropertiesResponsePtrOutput {
	return o
}

func (o EnvironmentDeploymentPropertiesResponsePtrOutput) ToEnvironmentDeploymentPropertiesResponsePtrOutputWithContext(ctx context.Context) EnvironmentDeploymentPropertiesResponsePtrOutput {
	return o
}

func (o EnvironmentDeploymentPropertiesResponsePtrOutput) Elem() EnvironmentDeploymentPropertiesResponseOutput {
	return o.ApplyT(func(v *EnvironmentDeploymentPropertiesResponse) EnvironmentDeploymentPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret EnvironmentDeploymentPropertiesResponse
		return ret
	}).(EnvironmentDeploymentPropertiesResponseOutput)
}

func (o EnvironmentDeploymentPropertiesResponsePtrOutput) ArmTemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentDeploymentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ArmTemplateId
	}).(pulumi.StringPtrOutput)
}

func (o EnvironmentDeploymentPropertiesResponsePtrOutput) Parameters() ArmTemplateParameterPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *EnvironmentDeploymentPropertiesResponse) []ArmTemplateParameterPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(ArmTemplateParameterPropertiesResponseArrayOutput)
}

type Event struct {
	EventName *string `pulumi:"eventName"`
}





type EventInput interface {
	pulumi.Input

	ToEventOutput() EventOutput
	ToEventOutputWithContext(context.Context) EventOutput
}

type EventArgs struct {
	EventName pulumi.StringPtrInput `pulumi:"eventName"`
}

func (EventArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Event)(nil)).Elem()
}

func (i EventArgs) ToEventOutput() EventOutput {
	return i.ToEventOutputWithContext(context.Background())
}

func (i EventArgs) ToEventOutputWithContext(ctx context.Context) EventOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventOutput)
}





type EventArrayInput interface {
	pulumi.Input

	ToEventArrayOutput() EventArrayOutput
	ToEventArrayOutputWithContext(context.Context) EventArrayOutput
}

type EventArray []EventInput

func (EventArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Event)(nil)).Elem()
}

func (i EventArray) ToEventArrayOutput() EventArrayOutput {
	return i.ToEventArrayOutputWithContext(context.Background())
}

func (i EventArray) ToEventArrayOutputWithContext(ctx context.Context) EventArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventArrayOutput)
}

type EventOutput struct{ *pulumi.OutputState }

func (EventOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Event)(nil)).Elem()
}

func (o EventOutput) ToEventOutput() EventOutput {
	return o
}

func (o EventOutput) ToEventOutputWithContext(ctx context.Context) EventOutput {
	return o
}

func (o EventOutput) EventName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Event) *string { return v.EventName }).(pulumi.StringPtrOutput)
}

type EventArrayOutput struct{ *pulumi.OutputState }

func (EventArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Event)(nil)).Elem()
}

func (o EventArrayOutput) ToEventArrayOutput() EventArrayOutput {
	return o
}

func (o EventArrayOutput) ToEventArrayOutputWithContext(ctx context.Context) EventArrayOutput {
	return o
}

func (o EventArrayOutput) Index(i pulumi.IntInput) EventOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Event {
		return vs[0].([]Event)[vs[1].(int)]
	}).(EventOutput)
}

type EventResponse struct {
	EventName *string `pulumi:"eventName"`
}

type EventResponseOutput struct{ *pulumi.OutputState }

func (EventResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventResponse)(nil)).Elem()
}

func (o EventResponseOutput) ToEventResponseOutput() EventResponseOutput {
	return o
}

func (o EventResponseOutput) ToEventResponseOutputWithContext(ctx context.Context) EventResponseOutput {
	return o
}

func (o EventResponseOutput) EventName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventResponse) *string { return v.EventName }).(pulumi.StringPtrOutput)
}

type EventResponseArrayOutput struct{ *pulumi.OutputState }

func (EventResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventResponse)(nil)).Elem()
}

func (o EventResponseArrayOutput) ToEventResponseArrayOutput() EventResponseArrayOutput {
	return o
}

func (o EventResponseArrayOutput) ToEventResponseArrayOutputWithContext(ctx context.Context) EventResponseArrayOutput {
	return o
}

func (o EventResponseArrayOutput) Index(i pulumi.IntInput) EventResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventResponse {
		return vs[0].([]EventResponse)[vs[1].(int)]
	}).(EventResponseOutput)
}

type ExternalSubnetResponse struct {
	Id   *string `pulumi:"id"`
	Name *string `pulumi:"name"`
}

type ExternalSubnetResponseOutput struct{ *pulumi.OutputState }

func (ExternalSubnetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExternalSubnetResponse)(nil)).Elem()
}

func (o ExternalSubnetResponseOutput) ToExternalSubnetResponseOutput() ExternalSubnetResponseOutput {
	return o
}

func (o ExternalSubnetResponseOutput) ToExternalSubnetResponseOutputWithContext(ctx context.Context) ExternalSubnetResponseOutput {
	return o
}

func (o ExternalSubnetResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExternalSubnetResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ExternalSubnetResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExternalSubnetResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ExternalSubnetResponseArrayOutput struct{ *pulumi.OutputState }

func (ExternalSubnetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExternalSubnetResponse)(nil)).Elem()
}

func (o ExternalSubnetResponseArrayOutput) ToExternalSubnetResponseArrayOutput() ExternalSubnetResponseArrayOutput {
	return o
}

func (o ExternalSubnetResponseArrayOutput) ToExternalSubnetResponseArrayOutputWithContext(ctx context.Context) ExternalSubnetResponseArrayOutput {
	return o
}

func (o ExternalSubnetResponseArrayOutput) Index(i pulumi.IntInput) ExternalSubnetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExternalSubnetResponse {
		return vs[0].([]ExternalSubnetResponse)[vs[1].(int)]
	}).(ExternalSubnetResponseOutput)
}

type FormulaPropertiesFromVm struct {
	LabVmId *string `pulumi:"labVmId"`
}





type FormulaPropertiesFromVmInput interface {
	pulumi.Input

	ToFormulaPropertiesFromVmOutput() FormulaPropertiesFromVmOutput
	ToFormulaPropertiesFromVmOutputWithContext(context.Context) FormulaPropertiesFromVmOutput
}

type FormulaPropertiesFromVmArgs struct {
	LabVmId pulumi.StringPtrInput `pulumi:"labVmId"`
}

func (FormulaPropertiesFromVmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FormulaPropertiesFromVm)(nil)).Elem()
}

func (i FormulaPropertiesFromVmArgs) ToFormulaPropertiesFromVmOutput() FormulaPropertiesFromVmOutput {
	return i.ToFormulaPropertiesFromVmOutputWithContext(context.Background())
}

func (i FormulaPropertiesFromVmArgs) ToFormulaPropertiesFromVmOutputWithContext(ctx context.Context) FormulaPropertiesFromVmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FormulaPropertiesFromVmOutput)
}

func (i FormulaPropertiesFromVmArgs) ToFormulaPropertiesFromVmPtrOutput() FormulaPropertiesFromVmPtrOutput {
	return i.ToFormulaPropertiesFromVmPtrOutputWithContext(context.Background())
}

func (i FormulaPropertiesFromVmArgs) ToFormulaPropertiesFromVmPtrOutputWithContext(ctx context.Context) FormulaPropertiesFromVmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FormulaPropertiesFromVmOutput).ToFormulaPropertiesFromVmPtrOutputWithContext(ctx)
}









type FormulaPropertiesFromVmPtrInput interface {
	pulumi.Input

	ToFormulaPropertiesFromVmPtrOutput() FormulaPropertiesFromVmPtrOutput
	ToFormulaPropertiesFromVmPtrOutputWithContext(context.Context) FormulaPropertiesFromVmPtrOutput
}

type formulaPropertiesFromVmPtrType FormulaPropertiesFromVmArgs

func FormulaPropertiesFromVmPtr(v *FormulaPropertiesFromVmArgs) FormulaPropertiesFromVmPtrInput {
	return (*formulaPropertiesFromVmPtrType)(v)
}

func (*formulaPropertiesFromVmPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FormulaPropertiesFromVm)(nil)).Elem()
}

func (i *formulaPropertiesFromVmPtrType) ToFormulaPropertiesFromVmPtrOutput() FormulaPropertiesFromVmPtrOutput {
	return i.ToFormulaPropertiesFromVmPtrOutputWithContext(context.Background())
}

func (i *formulaPropertiesFromVmPtrType) ToFormulaPropertiesFromVmPtrOutputWithContext(ctx context.Context) FormulaPropertiesFromVmPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FormulaPropertiesFromVmPtrOutput)
}

type FormulaPropertiesFromVmOutput struct{ *pulumi.OutputState }

func (FormulaPropertiesFromVmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FormulaPropertiesFromVm)(nil)).Elem()
}

func (o FormulaPropertiesFromVmOutput) ToFormulaPropertiesFromVmOutput() FormulaPropertiesFromVmOutput {
	return o
}

func (o FormulaPropertiesFromVmOutput) ToFormulaPropertiesFromVmOutputWithContext(ctx context.Context) FormulaPropertiesFromVmOutput {
	return o
}

func (o FormulaPropertiesFromVmOutput) ToFormulaPropertiesFromVmPtrOutput() FormulaPropertiesFromVmPtrOutput {
	return o.ToFormulaPropertiesFromVmPtrOutputWithContext(context.Background())
}

func (o FormulaPropertiesFromVmOutput) ToFormulaPropertiesFromVmPtrOutputWithContext(ctx context.Context) FormulaPropertiesFromVmPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FormulaPropertiesFromVm) *FormulaPropertiesFromVm {
		return &v
	}).(FormulaPropertiesFromVmPtrOutput)
}

func (o FormulaPropertiesFromVmOutput) LabVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FormulaPropertiesFromVm) *string { return v.LabVmId }).(pulumi.StringPtrOutput)
}

type FormulaPropertiesFromVmPtrOutput struct{ *pulumi.OutputState }

func (FormulaPropertiesFromVmPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FormulaPropertiesFromVm)(nil)).Elem()
}

func (o FormulaPropertiesFromVmPtrOutput) ToFormulaPropertiesFromVmPtrOutput() FormulaPropertiesFromVmPtrOutput {
	return o
}

func (o FormulaPropertiesFromVmPtrOutput) ToFormulaPropertiesFromVmPtrOutputWithContext(ctx context.Context) FormulaPropertiesFromVmPtrOutput {
	return o
}

func (o FormulaPropertiesFromVmPtrOutput) Elem() FormulaPropertiesFromVmOutput {
	return o.ApplyT(func(v *FormulaPropertiesFromVm) FormulaPropertiesFromVm {
		if v != nil {
			return *v
		}
		var ret FormulaPropertiesFromVm
		return ret
	}).(FormulaPropertiesFromVmOutput)
}

func (o FormulaPropertiesFromVmPtrOutput) LabVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FormulaPropertiesFromVm) *string {
		if v == nil {
			return nil
		}
		return v.LabVmId
	}).(pulumi.StringPtrOutput)
}

type FormulaPropertiesFromVmResponse struct {
	LabVmId *string `pulumi:"labVmId"`
}

type FormulaPropertiesFromVmResponseOutput struct{ *pulumi.OutputState }

func (FormulaPropertiesFromVmResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FormulaPropertiesFromVmResponse)(nil)).Elem()
}

func (o FormulaPropertiesFromVmResponseOutput) ToFormulaPropertiesFromVmResponseOutput() FormulaPropertiesFromVmResponseOutput {
	return o
}

func (o FormulaPropertiesFromVmResponseOutput) ToFormulaPropertiesFromVmResponseOutputWithContext(ctx context.Context) FormulaPropertiesFromVmResponseOutput {
	return o
}

func (o FormulaPropertiesFromVmResponseOutput) LabVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FormulaPropertiesFromVmResponse) *string { return v.LabVmId }).(pulumi.StringPtrOutput)
}

type FormulaPropertiesFromVmResponsePtrOutput struct{ *pulumi.OutputState }

func (FormulaPropertiesFromVmResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FormulaPropertiesFromVmResponse)(nil)).Elem()
}

func (o FormulaPropertiesFromVmResponsePtrOutput) ToFormulaPropertiesFromVmResponsePtrOutput() FormulaPropertiesFromVmResponsePtrOutput {
	return o
}

func (o FormulaPropertiesFromVmResponsePtrOutput) ToFormulaPropertiesFromVmResponsePtrOutputWithContext(ctx context.Context) FormulaPropertiesFromVmResponsePtrOutput {
	return o
}

func (o FormulaPropertiesFromVmResponsePtrOutput) Elem() FormulaPropertiesFromVmResponseOutput {
	return o.ApplyT(func(v *FormulaPropertiesFromVmResponse) FormulaPropertiesFromVmResponse {
		if v != nil {
			return *v
		}
		var ret FormulaPropertiesFromVmResponse
		return ret
	}).(FormulaPropertiesFromVmResponseOutput)
}

func (o FormulaPropertiesFromVmResponsePtrOutput) LabVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FormulaPropertiesFromVmResponse) *string {
		if v == nil {
			return nil
		}
		return v.LabVmId
	}).(pulumi.StringPtrOutput)
}

type GalleryImageReference struct {
	Offer     *string `pulumi:"offer"`
	OsType    *string `pulumi:"osType"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
	Version   *string `pulumi:"version"`
}





type GalleryImageReferenceInput interface {
	pulumi.Input

	ToGalleryImageReferenceOutput() GalleryImageReferenceOutput
	ToGalleryImageReferenceOutputWithContext(context.Context) GalleryImageReferenceOutput
}

type GalleryImageReferenceArgs struct {
	Offer     pulumi.StringPtrInput `pulumi:"offer"`
	OsType    pulumi.StringPtrInput `pulumi:"osType"`
	Publisher pulumi.StringPtrInput `pulumi:"publisher"`
	Sku       pulumi.StringPtrInput `pulumi:"sku"`
	Version   pulumi.StringPtrInput `pulumi:"version"`
}

func (GalleryImageReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageReference)(nil)).Elem()
}

func (i GalleryImageReferenceArgs) ToGalleryImageReferenceOutput() GalleryImageReferenceOutput {
	return i.ToGalleryImageReferenceOutputWithContext(context.Background())
}

func (i GalleryImageReferenceArgs) ToGalleryImageReferenceOutputWithContext(ctx context.Context) GalleryImageReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageReferenceOutput)
}

func (i GalleryImageReferenceArgs) ToGalleryImageReferencePtrOutput() GalleryImageReferencePtrOutput {
	return i.ToGalleryImageReferencePtrOutputWithContext(context.Background())
}

func (i GalleryImageReferenceArgs) ToGalleryImageReferencePtrOutputWithContext(ctx context.Context) GalleryImageReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageReferenceOutput).ToGalleryImageReferencePtrOutputWithContext(ctx)
}









type GalleryImageReferencePtrInput interface {
	pulumi.Input

	ToGalleryImageReferencePtrOutput() GalleryImageReferencePtrOutput
	ToGalleryImageReferencePtrOutputWithContext(context.Context) GalleryImageReferencePtrOutput
}

type galleryImageReferencePtrType GalleryImageReferenceArgs

func GalleryImageReferencePtr(v *GalleryImageReferenceArgs) GalleryImageReferencePtrInput {
	return (*galleryImageReferencePtrType)(v)
}

func (*galleryImageReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageReference)(nil)).Elem()
}

func (i *galleryImageReferencePtrType) ToGalleryImageReferencePtrOutput() GalleryImageReferencePtrOutput {
	return i.ToGalleryImageReferencePtrOutputWithContext(context.Background())
}

func (i *galleryImageReferencePtrType) ToGalleryImageReferencePtrOutputWithContext(ctx context.Context) GalleryImageReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageReferencePtrOutput)
}

type GalleryImageReferenceOutput struct{ *pulumi.OutputState }

func (GalleryImageReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageReference)(nil)).Elem()
}

func (o GalleryImageReferenceOutput) ToGalleryImageReferenceOutput() GalleryImageReferenceOutput {
	return o
}

func (o GalleryImageReferenceOutput) ToGalleryImageReferenceOutputWithContext(ctx context.Context) GalleryImageReferenceOutput {
	return o
}

func (o GalleryImageReferenceOutput) ToGalleryImageReferencePtrOutput() GalleryImageReferencePtrOutput {
	return o.ToGalleryImageReferencePtrOutputWithContext(context.Background())
}

func (o GalleryImageReferenceOutput) ToGalleryImageReferencePtrOutputWithContext(ctx context.Context) GalleryImageReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GalleryImageReference) *GalleryImageReference {
		return &v
	}).(GalleryImageReferencePtrOutput)
}

func (o GalleryImageReferenceOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReference) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReference) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReference) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReference) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReference) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type GalleryImageReferencePtrOutput struct{ *pulumi.OutputState }

func (GalleryImageReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageReference)(nil)).Elem()
}

func (o GalleryImageReferencePtrOutput) ToGalleryImageReferencePtrOutput() GalleryImageReferencePtrOutput {
	return o
}

func (o GalleryImageReferencePtrOutput) ToGalleryImageReferencePtrOutputWithContext(ctx context.Context) GalleryImageReferencePtrOutput {
	return o
}

func (o GalleryImageReferencePtrOutput) Elem() GalleryImageReferenceOutput {
	return o.ApplyT(func(v *GalleryImageReference) GalleryImageReference {
		if v != nil {
			return *v
		}
		var ret GalleryImageReference
		return ret
	}).(GalleryImageReferenceOutput)
}

func (o GalleryImageReferencePtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferencePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReference) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferencePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferencePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferencePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type GalleryImageReferenceResponse struct {
	Offer     *string `pulumi:"offer"`
	OsType    *string `pulumi:"osType"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
	Version   *string `pulumi:"version"`
}

type GalleryImageReferenceResponseOutput struct{ *pulumi.OutputState }

func (GalleryImageReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageReferenceResponse)(nil)).Elem()
}

func (o GalleryImageReferenceResponseOutput) ToGalleryImageReferenceResponseOutput() GalleryImageReferenceResponseOutput {
	return o
}

func (o GalleryImageReferenceResponseOutput) ToGalleryImageReferenceResponseOutputWithContext(ctx context.Context) GalleryImageReferenceResponseOutput {
	return o
}

func (o GalleryImageReferenceResponseOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponseOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type GalleryImageReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (GalleryImageReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageReferenceResponse)(nil)).Elem()
}

func (o GalleryImageReferenceResponsePtrOutput) ToGalleryImageReferenceResponsePtrOutput() GalleryImageReferenceResponsePtrOutput {
	return o
}

func (o GalleryImageReferenceResponsePtrOutput) ToGalleryImageReferenceResponsePtrOutputWithContext(ctx context.Context) GalleryImageReferenceResponsePtrOutput {
	return o
}

func (o GalleryImageReferenceResponsePtrOutput) Elem() GalleryImageReferenceResponseOutput {
	return o.ApplyT(func(v *GalleryImageReferenceResponse) GalleryImageReferenceResponse {
		if v != nil {
			return *v
		}
		var ret GalleryImageReferenceResponse
		return ret
	}).(GalleryImageReferenceResponseOutput)
}

func (o GalleryImageReferenceResponsePtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponsePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type HourDetails struct {
	Minute *int `pulumi:"minute"`
}





type HourDetailsInput interface {
	pulumi.Input

	ToHourDetailsOutput() HourDetailsOutput
	ToHourDetailsOutputWithContext(context.Context) HourDetailsOutput
}

type HourDetailsArgs struct {
	Minute pulumi.IntPtrInput `pulumi:"minute"`
}

func (HourDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HourDetails)(nil)).Elem()
}

func (i HourDetailsArgs) ToHourDetailsOutput() HourDetailsOutput {
	return i.ToHourDetailsOutputWithContext(context.Background())
}

func (i HourDetailsArgs) ToHourDetailsOutputWithContext(ctx context.Context) HourDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourDetailsOutput)
}

func (i HourDetailsArgs) ToHourDetailsPtrOutput() HourDetailsPtrOutput {
	return i.ToHourDetailsPtrOutputWithContext(context.Background())
}

func (i HourDetailsArgs) ToHourDetailsPtrOutputWithContext(ctx context.Context) HourDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourDetailsOutput).ToHourDetailsPtrOutputWithContext(ctx)
}









type HourDetailsPtrInput interface {
	pulumi.Input

	ToHourDetailsPtrOutput() HourDetailsPtrOutput
	ToHourDetailsPtrOutputWithContext(context.Context) HourDetailsPtrOutput
}

type hourDetailsPtrType HourDetailsArgs

func HourDetailsPtr(v *HourDetailsArgs) HourDetailsPtrInput {
	return (*hourDetailsPtrType)(v)
}

func (*hourDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HourDetails)(nil)).Elem()
}

func (i *hourDetailsPtrType) ToHourDetailsPtrOutput() HourDetailsPtrOutput {
	return i.ToHourDetailsPtrOutputWithContext(context.Background())
}

func (i *hourDetailsPtrType) ToHourDetailsPtrOutputWithContext(ctx context.Context) HourDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourDetailsPtrOutput)
}

type HourDetailsOutput struct{ *pulumi.OutputState }

func (HourDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HourDetails)(nil)).Elem()
}

func (o HourDetailsOutput) ToHourDetailsOutput() HourDetailsOutput {
	return o
}

func (o HourDetailsOutput) ToHourDetailsOutputWithContext(ctx context.Context) HourDetailsOutput {
	return o
}

func (o HourDetailsOutput) ToHourDetailsPtrOutput() HourDetailsPtrOutput {
	return o.ToHourDetailsPtrOutputWithContext(context.Background())
}

func (o HourDetailsOutput) ToHourDetailsPtrOutputWithContext(ctx context.Context) HourDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HourDetails) *HourDetails {
		return &v
	}).(HourDetailsPtrOutput)
}

func (o HourDetailsOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HourDetails) *int { return v.Minute }).(pulumi.IntPtrOutput)
}

type HourDetailsPtrOutput struct{ *pulumi.OutputState }

func (HourDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HourDetails)(nil)).Elem()
}

func (o HourDetailsPtrOutput) ToHourDetailsPtrOutput() HourDetailsPtrOutput {
	return o
}

func (o HourDetailsPtrOutput) ToHourDetailsPtrOutputWithContext(ctx context.Context) HourDetailsPtrOutput {
	return o
}

func (o HourDetailsPtrOutput) Elem() HourDetailsOutput {
	return o.ApplyT(func(v *HourDetails) HourDetails {
		if v != nil {
			return *v
		}
		var ret HourDetails
		return ret
	}).(HourDetailsOutput)
}

func (o HourDetailsPtrOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HourDetails) *int {
		if v == nil {
			return nil
		}
		return v.Minute
	}).(pulumi.IntPtrOutput)
}

type HourDetailsResponse struct {
	Minute *int `pulumi:"minute"`
}

type HourDetailsResponseOutput struct{ *pulumi.OutputState }

func (HourDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HourDetailsResponse)(nil)).Elem()
}

func (o HourDetailsResponseOutput) ToHourDetailsResponseOutput() HourDetailsResponseOutput {
	return o
}

func (o HourDetailsResponseOutput) ToHourDetailsResponseOutputWithContext(ctx context.Context) HourDetailsResponseOutput {
	return o
}

func (o HourDetailsResponseOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HourDetailsResponse) *int { return v.Minute }).(pulumi.IntPtrOutput)
}

type HourDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (HourDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HourDetailsResponse)(nil)).Elem()
}

func (o HourDetailsResponsePtrOutput) ToHourDetailsResponsePtrOutput() HourDetailsResponsePtrOutput {
	return o
}

func (o HourDetailsResponsePtrOutput) ToHourDetailsResponsePtrOutputWithContext(ctx context.Context) HourDetailsResponsePtrOutput {
	return o
}

func (o HourDetailsResponsePtrOutput) Elem() HourDetailsResponseOutput {
	return o.ApplyT(func(v *HourDetailsResponse) HourDetailsResponse {
		if v != nil {
			return *v
		}
		var ret HourDetailsResponse
		return ret
	}).(HourDetailsResponseOutput)
}

func (o HourDetailsResponsePtrOutput) Minute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HourDetailsResponse) *int {
		if v == nil {
			return nil
		}
		return v.Minute
	}).(pulumi.IntPtrOutput)
}

type IdentityProperties struct {
	ClientSecretUrl *string `pulumi:"clientSecretUrl"`
	PrincipalId     *string `pulumi:"principalId"`
	TenantId        *string `pulumi:"tenantId"`
	Type            *string `pulumi:"type"`
}





type IdentityPropertiesInput interface {
	pulumi.Input

	ToIdentityPropertiesOutput() IdentityPropertiesOutput
	ToIdentityPropertiesOutputWithContext(context.Context) IdentityPropertiesOutput
}

type IdentityPropertiesArgs struct {
	ClientSecretUrl pulumi.StringPtrInput `pulumi:"clientSecretUrl"`
	PrincipalId     pulumi.StringPtrInput `pulumi:"principalId"`
	TenantId        pulumi.StringPtrInput `pulumi:"tenantId"`
	Type            pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProperties)(nil)).Elem()
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesOutput() IdentityPropertiesOutput {
	return i.ToIdentityPropertiesOutputWithContext(context.Background())
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesOutputWithContext(ctx context.Context) IdentityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesOutput)
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return i.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesOutput).ToIdentityPropertiesPtrOutputWithContext(ctx)
}









type IdentityPropertiesPtrInput interface {
	pulumi.Input

	ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput
	ToIdentityPropertiesPtrOutputWithContext(context.Context) IdentityPropertiesPtrOutput
}

type identityPropertiesPtrType IdentityPropertiesArgs

func IdentityPropertiesPtr(v *IdentityPropertiesArgs) IdentityPropertiesPtrInput {
	return (*identityPropertiesPtrType)(v)
}

func (*identityPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProperties)(nil)).Elem()
}

func (i *identityPropertiesPtrType) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return i.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i *identityPropertiesPtrType) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesPtrOutput)
}

type IdentityPropertiesOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProperties)(nil)).Elem()
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesOutput() IdentityPropertiesOutput {
	return o
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesOutputWithContext(ctx context.Context) IdentityPropertiesOutput {
	return o
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return o.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityProperties) *IdentityProperties {
		return &v
	}).(IdentityPropertiesPtrOutput)
}

func (o IdentityPropertiesOutput) ClientSecretUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityProperties) *string { return v.ClientSecretUrl }).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityProperties) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityProperties) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityPropertiesPtrOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProperties)(nil)).Elem()
}

func (o IdentityPropertiesPtrOutput) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return o
}

func (o IdentityPropertiesPtrOutput) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return o
}

func (o IdentityPropertiesPtrOutput) Elem() IdentityPropertiesOutput {
	return o.ApplyT(func(v *IdentityProperties) IdentityProperties {
		if v != nil {
			return *v
		}
		var ret IdentityProperties
		return ret
	}).(IdentityPropertiesOutput)
}

func (o IdentityPropertiesPtrOutput) ClientSecretUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretUrl
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type IdentityPropertiesResponse struct {
	ClientSecretUrl *string `pulumi:"clientSecretUrl"`
	PrincipalId     *string `pulumi:"principalId"`
	TenantId        *string `pulumi:"tenantId"`
	Type            *string `pulumi:"type"`
}

type IdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityPropertiesResponse)(nil)).Elem()
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponseOutput() IdentityPropertiesResponseOutput {
	return o
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponseOutputWithContext(ctx context.Context) IdentityPropertiesResponseOutput {
	return o
}

func (o IdentityPropertiesResponseOutput) ClientSecretUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) *string { return v.ClientSecretUrl }).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityPropertiesResponse)(nil)).Elem()
}

func (o IdentityPropertiesResponsePtrOutput) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return o
}

func (o IdentityPropertiesResponsePtrOutput) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return o
}

func (o IdentityPropertiesResponsePtrOutput) Elem() IdentityPropertiesResponseOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) IdentityPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret IdentityPropertiesResponse
		return ret
	}).(IdentityPropertiesResponseOutput)
}

func (o IdentityPropertiesResponsePtrOutput) ClientSecretUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretUrl
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type InboundNatRule struct {
	BackendPort       *int    `pulumi:"backendPort"`
	FrontendPort      *int    `pulumi:"frontendPort"`
	TransportProtocol *string `pulumi:"transportProtocol"`
}





type InboundNatRuleInput interface {
	pulumi.Input

	ToInboundNatRuleOutput() InboundNatRuleOutput
	ToInboundNatRuleOutputWithContext(context.Context) InboundNatRuleOutput
}

type InboundNatRuleArgs struct {
	BackendPort       pulumi.IntPtrInput    `pulumi:"backendPort"`
	FrontendPort      pulumi.IntPtrInput    `pulumi:"frontendPort"`
	TransportProtocol pulumi.StringPtrInput `pulumi:"transportProtocol"`
}

func (InboundNatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatRule)(nil)).Elem()
}

func (i InboundNatRuleArgs) ToInboundNatRuleOutput() InboundNatRuleOutput {
	return i.ToInboundNatRuleOutputWithContext(context.Background())
}

func (i InboundNatRuleArgs) ToInboundNatRuleOutputWithContext(ctx context.Context) InboundNatRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatRuleOutput)
}





type InboundNatRuleArrayInput interface {
	pulumi.Input

	ToInboundNatRuleArrayOutput() InboundNatRuleArrayOutput
	ToInboundNatRuleArrayOutputWithContext(context.Context) InboundNatRuleArrayOutput
}

type InboundNatRuleArray []InboundNatRuleInput

func (InboundNatRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatRule)(nil)).Elem()
}

func (i InboundNatRuleArray) ToInboundNatRuleArrayOutput() InboundNatRuleArrayOutput {
	return i.ToInboundNatRuleArrayOutputWithContext(context.Background())
}

func (i InboundNatRuleArray) ToInboundNatRuleArrayOutputWithContext(ctx context.Context) InboundNatRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatRuleArrayOutput)
}

type InboundNatRuleOutput struct{ *pulumi.OutputState }

func (InboundNatRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatRule)(nil)).Elem()
}

func (o InboundNatRuleOutput) ToInboundNatRuleOutput() InboundNatRuleOutput {
	return o
}

func (o InboundNatRuleOutput) ToInboundNatRuleOutputWithContext(ctx context.Context) InboundNatRuleOutput {
	return o
}

func (o InboundNatRuleOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRule) *int { return v.BackendPort }).(pulumi.IntPtrOutput)
}

func (o InboundNatRuleOutput) FrontendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRule) *int { return v.FrontendPort }).(pulumi.IntPtrOutput)
}

func (o InboundNatRuleOutput) TransportProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRule) *string { return v.TransportProtocol }).(pulumi.StringPtrOutput)
}

type InboundNatRuleArrayOutput struct{ *pulumi.OutputState }

func (InboundNatRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatRule)(nil)).Elem()
}

func (o InboundNatRuleArrayOutput) ToInboundNatRuleArrayOutput() InboundNatRuleArrayOutput {
	return o
}

func (o InboundNatRuleArrayOutput) ToInboundNatRuleArrayOutputWithContext(ctx context.Context) InboundNatRuleArrayOutput {
	return o
}

func (o InboundNatRuleArrayOutput) Index(i pulumi.IntInput) InboundNatRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundNatRule {
		return vs[0].([]InboundNatRule)[vs[1].(int)]
	}).(InboundNatRuleOutput)
}

type InboundNatRuleResponse struct {
	BackendPort       *int    `pulumi:"backendPort"`
	FrontendPort      *int    `pulumi:"frontendPort"`
	TransportProtocol *string `pulumi:"transportProtocol"`
}

type InboundNatRuleResponseOutput struct{ *pulumi.OutputState }

func (InboundNatRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatRuleResponse)(nil)).Elem()
}

func (o InboundNatRuleResponseOutput) ToInboundNatRuleResponseOutput() InboundNatRuleResponseOutput {
	return o
}

func (o InboundNatRuleResponseOutput) ToInboundNatRuleResponseOutputWithContext(ctx context.Context) InboundNatRuleResponseOutput {
	return o
}

func (o InboundNatRuleResponseOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *int { return v.BackendPort }).(pulumi.IntPtrOutput)
}

func (o InboundNatRuleResponseOutput) FrontendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *int { return v.FrontendPort }).(pulumi.IntPtrOutput)
}

func (o InboundNatRuleResponseOutput) TransportProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *string { return v.TransportProtocol }).(pulumi.StringPtrOutput)
}

type InboundNatRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (InboundNatRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatRuleResponse)(nil)).Elem()
}

func (o InboundNatRuleResponseArrayOutput) ToInboundNatRuleResponseArrayOutput() InboundNatRuleResponseArrayOutput {
	return o
}

func (o InboundNatRuleResponseArrayOutput) ToInboundNatRuleResponseArrayOutputWithContext(ctx context.Context) InboundNatRuleResponseArrayOutput {
	return o
}

func (o InboundNatRuleResponseArrayOutput) Index(i pulumi.IntInput) InboundNatRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundNatRuleResponse {
		return vs[0].([]InboundNatRuleResponse)[vs[1].(int)]
	}).(InboundNatRuleResponseOutput)
}

type LabAnnouncementProperties struct {
	Enabled        *string `pulumi:"enabled"`
	ExpirationDate *string `pulumi:"expirationDate"`
	Expired        *bool   `pulumi:"expired"`
	Markdown       *string `pulumi:"markdown"`
	Title          *string `pulumi:"title"`
}





type LabAnnouncementPropertiesInput interface {
	pulumi.Input

	ToLabAnnouncementPropertiesOutput() LabAnnouncementPropertiesOutput
	ToLabAnnouncementPropertiesOutputWithContext(context.Context) LabAnnouncementPropertiesOutput
}

type LabAnnouncementPropertiesArgs struct {
	Enabled        pulumi.StringPtrInput `pulumi:"enabled"`
	ExpirationDate pulumi.StringPtrInput `pulumi:"expirationDate"`
	Expired        pulumi.BoolPtrInput   `pulumi:"expired"`
	Markdown       pulumi.StringPtrInput `pulumi:"markdown"`
	Title          pulumi.StringPtrInput `pulumi:"title"`
}

func (LabAnnouncementPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LabAnnouncementProperties)(nil)).Elem()
}

func (i LabAnnouncementPropertiesArgs) ToLabAnnouncementPropertiesOutput() LabAnnouncementPropertiesOutput {
	return i.ToLabAnnouncementPropertiesOutputWithContext(context.Background())
}

func (i LabAnnouncementPropertiesArgs) ToLabAnnouncementPropertiesOutputWithContext(ctx context.Context) LabAnnouncementPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabAnnouncementPropertiesOutput)
}

func (i LabAnnouncementPropertiesArgs) ToLabAnnouncementPropertiesPtrOutput() LabAnnouncementPropertiesPtrOutput {
	return i.ToLabAnnouncementPropertiesPtrOutputWithContext(context.Background())
}

func (i LabAnnouncementPropertiesArgs) ToLabAnnouncementPropertiesPtrOutputWithContext(ctx context.Context) LabAnnouncementPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabAnnouncementPropertiesOutput).ToLabAnnouncementPropertiesPtrOutputWithContext(ctx)
}









type LabAnnouncementPropertiesPtrInput interface {
	pulumi.Input

	ToLabAnnouncementPropertiesPtrOutput() LabAnnouncementPropertiesPtrOutput
	ToLabAnnouncementPropertiesPtrOutputWithContext(context.Context) LabAnnouncementPropertiesPtrOutput
}

type labAnnouncementPropertiesPtrType LabAnnouncementPropertiesArgs

func LabAnnouncementPropertiesPtr(v *LabAnnouncementPropertiesArgs) LabAnnouncementPropertiesPtrInput {
	return (*labAnnouncementPropertiesPtrType)(v)
}

func (*labAnnouncementPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LabAnnouncementProperties)(nil)).Elem()
}

func (i *labAnnouncementPropertiesPtrType) ToLabAnnouncementPropertiesPtrOutput() LabAnnouncementPropertiesPtrOutput {
	return i.ToLabAnnouncementPropertiesPtrOutputWithContext(context.Background())
}

func (i *labAnnouncementPropertiesPtrType) ToLabAnnouncementPropertiesPtrOutputWithContext(ctx context.Context) LabAnnouncementPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabAnnouncementPropertiesPtrOutput)
}

type LabAnnouncementPropertiesOutput struct{ *pulumi.OutputState }

func (LabAnnouncementPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabAnnouncementProperties)(nil)).Elem()
}

func (o LabAnnouncementPropertiesOutput) ToLabAnnouncementPropertiesOutput() LabAnnouncementPropertiesOutput {
	return o
}

func (o LabAnnouncementPropertiesOutput) ToLabAnnouncementPropertiesOutputWithContext(ctx context.Context) LabAnnouncementPropertiesOutput {
	return o
}

func (o LabAnnouncementPropertiesOutput) ToLabAnnouncementPropertiesPtrOutput() LabAnnouncementPropertiesPtrOutput {
	return o.ToLabAnnouncementPropertiesPtrOutputWithContext(context.Background())
}

func (o LabAnnouncementPropertiesOutput) ToLabAnnouncementPropertiesPtrOutputWithContext(ctx context.Context) LabAnnouncementPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LabAnnouncementProperties) *LabAnnouncementProperties {
		return &v
	}).(LabAnnouncementPropertiesPtrOutput)
}

func (o LabAnnouncementPropertiesOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabAnnouncementProperties) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

func (o LabAnnouncementPropertiesOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabAnnouncementProperties) *string { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

func (o LabAnnouncementPropertiesOutput) Expired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LabAnnouncementProperties) *bool { return v.Expired }).(pulumi.BoolPtrOutput)
}

func (o LabAnnouncementPropertiesOutput) Markdown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabAnnouncementProperties) *string { return v.Markdown }).(pulumi.StringPtrOutput)
}

func (o LabAnnouncementPropertiesOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabAnnouncementProperties) *string { return v.Title }).(pulumi.StringPtrOutput)
}

type LabAnnouncementPropertiesPtrOutput struct{ *pulumi.OutputState }

func (LabAnnouncementPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabAnnouncementProperties)(nil)).Elem()
}

func (o LabAnnouncementPropertiesPtrOutput) ToLabAnnouncementPropertiesPtrOutput() LabAnnouncementPropertiesPtrOutput {
	return o
}

func (o LabAnnouncementPropertiesPtrOutput) ToLabAnnouncementPropertiesPtrOutputWithContext(ctx context.Context) LabAnnouncementPropertiesPtrOutput {
	return o
}

func (o LabAnnouncementPropertiesPtrOutput) Elem() LabAnnouncementPropertiesOutput {
	return o.ApplyT(func(v *LabAnnouncementProperties) LabAnnouncementProperties {
		if v != nil {
			return *v
		}
		var ret LabAnnouncementProperties
		return ret
	}).(LabAnnouncementPropertiesOutput)
}

func (o LabAnnouncementPropertiesPtrOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabAnnouncementProperties) *string {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.StringPtrOutput)
}

func (o LabAnnouncementPropertiesPtrOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabAnnouncementProperties) *string {
		if v == nil {
			return nil
		}
		return v.ExpirationDate
	}).(pulumi.StringPtrOutput)
}

func (o LabAnnouncementPropertiesPtrOutput) Expired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LabAnnouncementProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Expired
	}).(pulumi.BoolPtrOutput)
}

func (o LabAnnouncementPropertiesPtrOutput) Markdown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabAnnouncementProperties) *string {
		if v == nil {
			return nil
		}
		return v.Markdown
	}).(pulumi.StringPtrOutput)
}

func (o LabAnnouncementPropertiesPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabAnnouncementProperties) *string {
		if v == nil {
			return nil
		}
		return v.Title
	}).(pulumi.StringPtrOutput)
}

type LabAnnouncementPropertiesResponse struct {
	Enabled           *string `pulumi:"enabled"`
	ExpirationDate    *string `pulumi:"expirationDate"`
	Expired           *bool   `pulumi:"expired"`
	Markdown          *string `pulumi:"markdown"`
	ProvisioningState string  `pulumi:"provisioningState"`
	Title             *string `pulumi:"title"`
	UniqueIdentifier  string  `pulumi:"uniqueIdentifier"`
}

type LabAnnouncementPropertiesResponseOutput struct{ *pulumi.OutputState }

func (LabAnnouncementPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabAnnouncementPropertiesResponse)(nil)).Elem()
}

func (o LabAnnouncementPropertiesResponseOutput) ToLabAnnouncementPropertiesResponseOutput() LabAnnouncementPropertiesResponseOutput {
	return o
}

func (o LabAnnouncementPropertiesResponseOutput) ToLabAnnouncementPropertiesResponseOutputWithContext(ctx context.Context) LabAnnouncementPropertiesResponseOutput {
	return o
}

func (o LabAnnouncementPropertiesResponseOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabAnnouncementPropertiesResponse) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

func (o LabAnnouncementPropertiesResponseOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabAnnouncementPropertiesResponse) *string { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

func (o LabAnnouncementPropertiesResponseOutput) Expired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LabAnnouncementPropertiesResponse) *bool { return v.Expired }).(pulumi.BoolPtrOutput)
}

func (o LabAnnouncementPropertiesResponseOutput) Markdown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabAnnouncementPropertiesResponse) *string { return v.Markdown }).(pulumi.StringPtrOutput)
}

func (o LabAnnouncementPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LabAnnouncementPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LabAnnouncementPropertiesResponseOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabAnnouncementPropertiesResponse) *string { return v.Title }).(pulumi.StringPtrOutput)
}

func (o LabAnnouncementPropertiesResponseOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LabAnnouncementPropertiesResponse) string { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

type LabAnnouncementPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (LabAnnouncementPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabAnnouncementPropertiesResponse)(nil)).Elem()
}

func (o LabAnnouncementPropertiesResponsePtrOutput) ToLabAnnouncementPropertiesResponsePtrOutput() LabAnnouncementPropertiesResponsePtrOutput {
	return o
}

func (o LabAnnouncementPropertiesResponsePtrOutput) ToLabAnnouncementPropertiesResponsePtrOutputWithContext(ctx context.Context) LabAnnouncementPropertiesResponsePtrOutput {
	return o
}

func (o LabAnnouncementPropertiesResponsePtrOutput) Elem() LabAnnouncementPropertiesResponseOutput {
	return o.ApplyT(func(v *LabAnnouncementPropertiesResponse) LabAnnouncementPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret LabAnnouncementPropertiesResponse
		return ret
	}).(LabAnnouncementPropertiesResponseOutput)
}

func (o LabAnnouncementPropertiesResponsePtrOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabAnnouncementPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.StringPtrOutput)
}

func (o LabAnnouncementPropertiesResponsePtrOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabAnnouncementPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExpirationDate
	}).(pulumi.StringPtrOutput)
}

func (o LabAnnouncementPropertiesResponsePtrOutput) Expired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LabAnnouncementPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Expired
	}).(pulumi.BoolPtrOutput)
}

func (o LabAnnouncementPropertiesResponsePtrOutput) Markdown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabAnnouncementPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Markdown
	}).(pulumi.StringPtrOutput)
}

func (o LabAnnouncementPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabAnnouncementPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o LabAnnouncementPropertiesResponsePtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabAnnouncementPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Title
	}).(pulumi.StringPtrOutput)
}

func (o LabAnnouncementPropertiesResponsePtrOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabAnnouncementPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UniqueIdentifier
	}).(pulumi.StringPtrOutput)
}

type LabSupportProperties struct {
	Enabled  *string `pulumi:"enabled"`
	Markdown *string `pulumi:"markdown"`
}





type LabSupportPropertiesInput interface {
	pulumi.Input

	ToLabSupportPropertiesOutput() LabSupportPropertiesOutput
	ToLabSupportPropertiesOutputWithContext(context.Context) LabSupportPropertiesOutput
}

type LabSupportPropertiesArgs struct {
	Enabled  pulumi.StringPtrInput `pulumi:"enabled"`
	Markdown pulumi.StringPtrInput `pulumi:"markdown"`
}

func (LabSupportPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LabSupportProperties)(nil)).Elem()
}

func (i LabSupportPropertiesArgs) ToLabSupportPropertiesOutput() LabSupportPropertiesOutput {
	return i.ToLabSupportPropertiesOutputWithContext(context.Background())
}

func (i LabSupportPropertiesArgs) ToLabSupportPropertiesOutputWithContext(ctx context.Context) LabSupportPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabSupportPropertiesOutput)
}

func (i LabSupportPropertiesArgs) ToLabSupportPropertiesPtrOutput() LabSupportPropertiesPtrOutput {
	return i.ToLabSupportPropertiesPtrOutputWithContext(context.Background())
}

func (i LabSupportPropertiesArgs) ToLabSupportPropertiesPtrOutputWithContext(ctx context.Context) LabSupportPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabSupportPropertiesOutput).ToLabSupportPropertiesPtrOutputWithContext(ctx)
}









type LabSupportPropertiesPtrInput interface {
	pulumi.Input

	ToLabSupportPropertiesPtrOutput() LabSupportPropertiesPtrOutput
	ToLabSupportPropertiesPtrOutputWithContext(context.Context) LabSupportPropertiesPtrOutput
}

type labSupportPropertiesPtrType LabSupportPropertiesArgs

func LabSupportPropertiesPtr(v *LabSupportPropertiesArgs) LabSupportPropertiesPtrInput {
	return (*labSupportPropertiesPtrType)(v)
}

func (*labSupportPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LabSupportProperties)(nil)).Elem()
}

func (i *labSupportPropertiesPtrType) ToLabSupportPropertiesPtrOutput() LabSupportPropertiesPtrOutput {
	return i.ToLabSupportPropertiesPtrOutputWithContext(context.Background())
}

func (i *labSupportPropertiesPtrType) ToLabSupportPropertiesPtrOutputWithContext(ctx context.Context) LabSupportPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabSupportPropertiesPtrOutput)
}

type LabSupportPropertiesOutput struct{ *pulumi.OutputState }

func (LabSupportPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabSupportProperties)(nil)).Elem()
}

func (o LabSupportPropertiesOutput) ToLabSupportPropertiesOutput() LabSupportPropertiesOutput {
	return o
}

func (o LabSupportPropertiesOutput) ToLabSupportPropertiesOutputWithContext(ctx context.Context) LabSupportPropertiesOutput {
	return o
}

func (o LabSupportPropertiesOutput) ToLabSupportPropertiesPtrOutput() LabSupportPropertiesPtrOutput {
	return o.ToLabSupportPropertiesPtrOutputWithContext(context.Background())
}

func (o LabSupportPropertiesOutput) ToLabSupportPropertiesPtrOutputWithContext(ctx context.Context) LabSupportPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LabSupportProperties) *LabSupportProperties {
		return &v
	}).(LabSupportPropertiesPtrOutput)
}

func (o LabSupportPropertiesOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabSupportProperties) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

func (o LabSupportPropertiesOutput) Markdown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabSupportProperties) *string { return v.Markdown }).(pulumi.StringPtrOutput)
}

type LabSupportPropertiesPtrOutput struct{ *pulumi.OutputState }

func (LabSupportPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabSupportProperties)(nil)).Elem()
}

func (o LabSupportPropertiesPtrOutput) ToLabSupportPropertiesPtrOutput() LabSupportPropertiesPtrOutput {
	return o
}

func (o LabSupportPropertiesPtrOutput) ToLabSupportPropertiesPtrOutputWithContext(ctx context.Context) LabSupportPropertiesPtrOutput {
	return o
}

func (o LabSupportPropertiesPtrOutput) Elem() LabSupportPropertiesOutput {
	return o.ApplyT(func(v *LabSupportProperties) LabSupportProperties {
		if v != nil {
			return *v
		}
		var ret LabSupportProperties
		return ret
	}).(LabSupportPropertiesOutput)
}

func (o LabSupportPropertiesPtrOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabSupportProperties) *string {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.StringPtrOutput)
}

func (o LabSupportPropertiesPtrOutput) Markdown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabSupportProperties) *string {
		if v == nil {
			return nil
		}
		return v.Markdown
	}).(pulumi.StringPtrOutput)
}

type LabSupportPropertiesResponse struct {
	Enabled  *string `pulumi:"enabled"`
	Markdown *string `pulumi:"markdown"`
}

type LabSupportPropertiesResponseOutput struct{ *pulumi.OutputState }

func (LabSupportPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabSupportPropertiesResponse)(nil)).Elem()
}

func (o LabSupportPropertiesResponseOutput) ToLabSupportPropertiesResponseOutput() LabSupportPropertiesResponseOutput {
	return o
}

func (o LabSupportPropertiesResponseOutput) ToLabSupportPropertiesResponseOutputWithContext(ctx context.Context) LabSupportPropertiesResponseOutput {
	return o
}

func (o LabSupportPropertiesResponseOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabSupportPropertiesResponse) *string { return v.Enabled }).(pulumi.StringPtrOutput)
}

func (o LabSupportPropertiesResponseOutput) Markdown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabSupportPropertiesResponse) *string { return v.Markdown }).(pulumi.StringPtrOutput)
}

type LabSupportPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (LabSupportPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabSupportPropertiesResponse)(nil)).Elem()
}

func (o LabSupportPropertiesResponsePtrOutput) ToLabSupportPropertiesResponsePtrOutput() LabSupportPropertiesResponsePtrOutput {
	return o
}

func (o LabSupportPropertiesResponsePtrOutput) ToLabSupportPropertiesResponsePtrOutputWithContext(ctx context.Context) LabSupportPropertiesResponsePtrOutput {
	return o
}

func (o LabSupportPropertiesResponsePtrOutput) Elem() LabSupportPropertiesResponseOutput {
	return o.ApplyT(func(v *LabSupportPropertiesResponse) LabSupportPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret LabSupportPropertiesResponse
		return ret
	}).(LabSupportPropertiesResponseOutput)
}

func (o LabSupportPropertiesResponsePtrOutput) Enabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabSupportPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.StringPtrOutput)
}

func (o LabSupportPropertiesResponsePtrOutput) Markdown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabSupportPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Markdown
	}).(pulumi.StringPtrOutput)
}

type LabVhdResponse struct {
	Id *string `pulumi:"id"`
}

type LabVhdResponseOutput struct{ *pulumi.OutputState }

func (LabVhdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabVhdResponse)(nil)).Elem()
}

func (o LabVhdResponseOutput) ToLabVhdResponseOutput() LabVhdResponseOutput {
	return o
}

func (o LabVhdResponseOutput) ToLabVhdResponseOutputWithContext(ctx context.Context) LabVhdResponseOutput {
	return o
}

func (o LabVhdResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVhdResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type LabVhdResponseArrayOutput struct{ *pulumi.OutputState }

func (LabVhdResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LabVhdResponse)(nil)).Elem()
}

func (o LabVhdResponseArrayOutput) ToLabVhdResponseArrayOutput() LabVhdResponseArrayOutput {
	return o
}

func (o LabVhdResponseArrayOutput) ToLabVhdResponseArrayOutputWithContext(ctx context.Context) LabVhdResponseArrayOutput {
	return o
}

func (o LabVhdResponseArrayOutput) Index(i pulumi.IntInput) LabVhdResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LabVhdResponse {
		return vs[0].([]LabVhdResponse)[vs[1].(int)]
	}).(LabVhdResponseOutput)
}

type LabVirtualMachineCreationParameter struct {
	AllowClaim                 *bool                       `pulumi:"allowClaim"`
	Artifacts                  []ArtifactInstallProperties `pulumi:"artifacts"`
	BulkCreationParameters     *BulkCreationParameters     `pulumi:"bulkCreationParameters"`
	CreatedDate                *string                     `pulumi:"createdDate"`
	CustomImageId              *string                     `pulumi:"customImageId"`
	DataDiskParameters         []DataDiskProperties        `pulumi:"dataDiskParameters"`
	DisallowPublicIpAddress    *bool                       `pulumi:"disallowPublicIpAddress"`
	EnvironmentId              *string                     `pulumi:"environmentId"`
	ExpirationDate             *string                     `pulumi:"expirationDate"`
	GalleryImageReference      *GalleryImageReference      `pulumi:"galleryImageReference"`
	IsAuthenticationWithSshKey *bool                       `pulumi:"isAuthenticationWithSshKey"`
	LabSubnetName              *string                     `pulumi:"labSubnetName"`
	LabVirtualNetworkId        *string                     `pulumi:"labVirtualNetworkId"`
	Location                   *string                     `pulumi:"location"`
	Name                       *string                     `pulumi:"name"`
	NetworkInterface           *NetworkInterfaceProperties `pulumi:"networkInterface"`
	Notes                      *string                     `pulumi:"notes"`
	OwnerObjectId              *string                     `pulumi:"ownerObjectId"`
	OwnerUserPrincipalName     *string                     `pulumi:"ownerUserPrincipalName"`
	Password                   *string                     `pulumi:"password"`
	PlanId                     *string                     `pulumi:"planId"`
	ScheduleParameters         []ScheduleCreationParameter `pulumi:"scheduleParameters"`
	Size                       *string                     `pulumi:"size"`
	SshKey                     *string                     `pulumi:"sshKey"`
	StorageType                *string                     `pulumi:"storageType"`
	Tags                       map[string]string           `pulumi:"tags"`
	UserName                   *string                     `pulumi:"userName"`
}


func (val *LabVirtualMachineCreationParameter) Defaults() *LabVirtualMachineCreationParameter {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllowClaim) {
		allowClaim_ := false
		tmp.AllowClaim = &allowClaim_
	}
	if isZero(tmp.DisallowPublicIpAddress) {
		disallowPublicIpAddress_ := false
		tmp.DisallowPublicIpAddress = &disallowPublicIpAddress_
	}
	if isZero(tmp.OwnerObjectId) {
		ownerObjectId_ := "dynamicValue"
		tmp.OwnerObjectId = &ownerObjectId_
	}
	if isZero(tmp.StorageType) {
		storageType_ := "labStorageType"
		tmp.StorageType = &storageType_
	}
	return &tmp
}





type LabVirtualMachineCreationParameterInput interface {
	pulumi.Input

	ToLabVirtualMachineCreationParameterOutput() LabVirtualMachineCreationParameterOutput
	ToLabVirtualMachineCreationParameterOutputWithContext(context.Context) LabVirtualMachineCreationParameterOutput
}

type LabVirtualMachineCreationParameterArgs struct {
	AllowClaim                 pulumi.BoolPtrInput                 `pulumi:"allowClaim"`
	Artifacts                  ArtifactInstallPropertiesArrayInput `pulumi:"artifacts"`
	BulkCreationParameters     BulkCreationParametersPtrInput      `pulumi:"bulkCreationParameters"`
	CreatedDate                pulumi.StringPtrInput               `pulumi:"createdDate"`
	CustomImageId              pulumi.StringPtrInput               `pulumi:"customImageId"`
	DataDiskParameters         DataDiskPropertiesArrayInput        `pulumi:"dataDiskParameters"`
	DisallowPublicIpAddress    pulumi.BoolPtrInput                 `pulumi:"disallowPublicIpAddress"`
	EnvironmentId              pulumi.StringPtrInput               `pulumi:"environmentId"`
	ExpirationDate             pulumi.StringPtrInput               `pulumi:"expirationDate"`
	GalleryImageReference      GalleryImageReferencePtrInput       `pulumi:"galleryImageReference"`
	IsAuthenticationWithSshKey pulumi.BoolPtrInput                 `pulumi:"isAuthenticationWithSshKey"`
	LabSubnetName              pulumi.StringPtrInput               `pulumi:"labSubnetName"`
	LabVirtualNetworkId        pulumi.StringPtrInput               `pulumi:"labVirtualNetworkId"`
	Location                   pulumi.StringPtrInput               `pulumi:"location"`
	Name                       pulumi.StringPtrInput               `pulumi:"name"`
	NetworkInterface           NetworkInterfacePropertiesPtrInput  `pulumi:"networkInterface"`
	Notes                      pulumi.StringPtrInput               `pulumi:"notes"`
	OwnerObjectId              pulumi.StringPtrInput               `pulumi:"ownerObjectId"`
	OwnerUserPrincipalName     pulumi.StringPtrInput               `pulumi:"ownerUserPrincipalName"`
	Password                   pulumi.StringPtrInput               `pulumi:"password"`
	PlanId                     pulumi.StringPtrInput               `pulumi:"planId"`
	ScheduleParameters         ScheduleCreationParameterArrayInput `pulumi:"scheduleParameters"`
	Size                       pulumi.StringPtrInput               `pulumi:"size"`
	SshKey                     pulumi.StringPtrInput               `pulumi:"sshKey"`
	StorageType                pulumi.StringPtrInput               `pulumi:"storageType"`
	Tags                       pulumi.StringMapInput               `pulumi:"tags"`
	UserName                   pulumi.StringPtrInput               `pulumi:"userName"`
}


func (val *LabVirtualMachineCreationParameterArgs) Defaults() *LabVirtualMachineCreationParameterArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllowClaim) {
		tmp.AllowClaim = pulumi.BoolPtr(false)
	}
	if isZero(tmp.DisallowPublicIpAddress) {
		tmp.DisallowPublicIpAddress = pulumi.BoolPtr(false)
	}
	if isZero(tmp.OwnerObjectId) {
		tmp.OwnerObjectId = pulumi.StringPtr("dynamicValue")
	}
	if isZero(tmp.StorageType) {
		tmp.StorageType = pulumi.StringPtr("labStorageType")
	}
	return &tmp
}
func (LabVirtualMachineCreationParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LabVirtualMachineCreationParameter)(nil)).Elem()
}

func (i LabVirtualMachineCreationParameterArgs) ToLabVirtualMachineCreationParameterOutput() LabVirtualMachineCreationParameterOutput {
	return i.ToLabVirtualMachineCreationParameterOutputWithContext(context.Background())
}

func (i LabVirtualMachineCreationParameterArgs) ToLabVirtualMachineCreationParameterOutputWithContext(ctx context.Context) LabVirtualMachineCreationParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabVirtualMachineCreationParameterOutput)
}

func (i LabVirtualMachineCreationParameterArgs) ToLabVirtualMachineCreationParameterPtrOutput() LabVirtualMachineCreationParameterPtrOutput {
	return i.ToLabVirtualMachineCreationParameterPtrOutputWithContext(context.Background())
}

func (i LabVirtualMachineCreationParameterArgs) ToLabVirtualMachineCreationParameterPtrOutputWithContext(ctx context.Context) LabVirtualMachineCreationParameterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabVirtualMachineCreationParameterOutput).ToLabVirtualMachineCreationParameterPtrOutputWithContext(ctx)
}









type LabVirtualMachineCreationParameterPtrInput interface {
	pulumi.Input

	ToLabVirtualMachineCreationParameterPtrOutput() LabVirtualMachineCreationParameterPtrOutput
	ToLabVirtualMachineCreationParameterPtrOutputWithContext(context.Context) LabVirtualMachineCreationParameterPtrOutput
}

type labVirtualMachineCreationParameterPtrType LabVirtualMachineCreationParameterArgs

func LabVirtualMachineCreationParameterPtr(v *LabVirtualMachineCreationParameterArgs) LabVirtualMachineCreationParameterPtrInput {
	return (*labVirtualMachineCreationParameterPtrType)(v)
}

func (*labVirtualMachineCreationParameterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LabVirtualMachineCreationParameter)(nil)).Elem()
}

func (i *labVirtualMachineCreationParameterPtrType) ToLabVirtualMachineCreationParameterPtrOutput() LabVirtualMachineCreationParameterPtrOutput {
	return i.ToLabVirtualMachineCreationParameterPtrOutputWithContext(context.Background())
}

func (i *labVirtualMachineCreationParameterPtrType) ToLabVirtualMachineCreationParameterPtrOutputWithContext(ctx context.Context) LabVirtualMachineCreationParameterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabVirtualMachineCreationParameterPtrOutput)
}

type LabVirtualMachineCreationParameterOutput struct{ *pulumi.OutputState }

func (LabVirtualMachineCreationParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabVirtualMachineCreationParameter)(nil)).Elem()
}

func (o LabVirtualMachineCreationParameterOutput) ToLabVirtualMachineCreationParameterOutput() LabVirtualMachineCreationParameterOutput {
	return o
}

func (o LabVirtualMachineCreationParameterOutput) ToLabVirtualMachineCreationParameterOutputWithContext(ctx context.Context) LabVirtualMachineCreationParameterOutput {
	return o
}

func (o LabVirtualMachineCreationParameterOutput) ToLabVirtualMachineCreationParameterPtrOutput() LabVirtualMachineCreationParameterPtrOutput {
	return o.ToLabVirtualMachineCreationParameterPtrOutputWithContext(context.Background())
}

func (o LabVirtualMachineCreationParameterOutput) ToLabVirtualMachineCreationParameterPtrOutputWithContext(ctx context.Context) LabVirtualMachineCreationParameterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LabVirtualMachineCreationParameter) *LabVirtualMachineCreationParameter {
		return &v
	}).(LabVirtualMachineCreationParameterPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) AllowClaim() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *bool { return v.AllowClaim }).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) Artifacts() ArtifactInstallPropertiesArrayOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) []ArtifactInstallProperties { return v.Artifacts }).(ArtifactInstallPropertiesArrayOutput)
}

func (o LabVirtualMachineCreationParameterOutput) BulkCreationParameters() BulkCreationParametersPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *BulkCreationParameters { return v.BulkCreationParameters }).(BulkCreationParametersPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *string { return v.CreatedDate }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) CustomImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *string { return v.CustomImageId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) DataDiskParameters() DataDiskPropertiesArrayOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) []DataDiskProperties { return v.DataDiskParameters }).(DataDiskPropertiesArrayOutput)
}

func (o LabVirtualMachineCreationParameterOutput) DisallowPublicIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *bool { return v.DisallowPublicIpAddress }).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *string { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *string { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) GalleryImageReference() GalleryImageReferencePtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *GalleryImageReference { return v.GalleryImageReference }).(GalleryImageReferencePtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) IsAuthenticationWithSshKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *bool { return v.IsAuthenticationWithSshKey }).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *string { return v.LabSubnetName }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) LabVirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *string { return v.LabVirtualNetworkId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) NetworkInterface() NetworkInterfacePropertiesPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *NetworkInterfaceProperties { return v.NetworkInterface }).(NetworkInterfacePropertiesPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) OwnerObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *string { return v.OwnerObjectId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) OwnerUserPrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *string { return v.OwnerUserPrincipalName }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) PlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *string { return v.PlanId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) ScheduleParameters() ScheduleCreationParameterArrayOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) []ScheduleCreationParameter { return v.ScheduleParameters }).(ScheduleCreationParameterArrayOutput)
}

func (o LabVirtualMachineCreationParameterOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) SshKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *string { return v.SshKey }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *string { return v.StorageType }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LabVirtualMachineCreationParameterOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameter) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

type LabVirtualMachineCreationParameterPtrOutput struct{ *pulumi.OutputState }

func (LabVirtualMachineCreationParameterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabVirtualMachineCreationParameter)(nil)).Elem()
}

func (o LabVirtualMachineCreationParameterPtrOutput) ToLabVirtualMachineCreationParameterPtrOutput() LabVirtualMachineCreationParameterPtrOutput {
	return o
}

func (o LabVirtualMachineCreationParameterPtrOutput) ToLabVirtualMachineCreationParameterPtrOutputWithContext(ctx context.Context) LabVirtualMachineCreationParameterPtrOutput {
	return o
}

func (o LabVirtualMachineCreationParameterPtrOutput) Elem() LabVirtualMachineCreationParameterOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) LabVirtualMachineCreationParameter {
		if v != nil {
			return *v
		}
		var ret LabVirtualMachineCreationParameter
		return ret
	}).(LabVirtualMachineCreationParameterOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) AllowClaim() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *bool {
		if v == nil {
			return nil
		}
		return v.AllowClaim
	}).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) Artifacts() ArtifactInstallPropertiesArrayOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) []ArtifactInstallProperties {
		if v == nil {
			return nil
		}
		return v.Artifacts
	}).(ArtifactInstallPropertiesArrayOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) BulkCreationParameters() BulkCreationParametersPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *BulkCreationParameters {
		if v == nil {
			return nil
		}
		return v.BulkCreationParameters
	}).(BulkCreationParametersPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *string {
		if v == nil {
			return nil
		}
		return v.CreatedDate
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) CustomImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *string {
		if v == nil {
			return nil
		}
		return v.CustomImageId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) DataDiskParameters() DataDiskPropertiesArrayOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) []DataDiskProperties {
		if v == nil {
			return nil
		}
		return v.DataDiskParameters
	}).(DataDiskPropertiesArrayOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) DisallowPublicIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *bool {
		if v == nil {
			return nil
		}
		return v.DisallowPublicIpAddress
	}).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *string {
		if v == nil {
			return nil
		}
		return v.EnvironmentId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *string {
		if v == nil {
			return nil
		}
		return v.ExpirationDate
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) GalleryImageReference() GalleryImageReferencePtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *GalleryImageReference {
		if v == nil {
			return nil
		}
		return v.GalleryImageReference
	}).(GalleryImageReferencePtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) IsAuthenticationWithSshKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *bool {
		if v == nil {
			return nil
		}
		return v.IsAuthenticationWithSshKey
	}).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *string {
		if v == nil {
			return nil
		}
		return v.LabSubnetName
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) LabVirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *string {
		if v == nil {
			return nil
		}
		return v.LabVirtualNetworkId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) NetworkInterface() NetworkInterfacePropertiesPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *NetworkInterfaceProperties {
		if v == nil {
			return nil
		}
		return v.NetworkInterface
	}).(NetworkInterfacePropertiesPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *string {
		if v == nil {
			return nil
		}
		return v.Notes
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) OwnerObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *string {
		if v == nil {
			return nil
		}
		return v.OwnerObjectId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) OwnerUserPrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *string {
		if v == nil {
			return nil
		}
		return v.OwnerUserPrincipalName
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) PlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *string {
		if v == nil {
			return nil
		}
		return v.PlanId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) ScheduleParameters() ScheduleCreationParameterArrayOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) []ScheduleCreationParameter {
		if v == nil {
			return nil
		}
		return v.ScheduleParameters
	}).(ScheduleCreationParameterArrayOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) SshKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *string {
		if v == nil {
			return nil
		}
		return v.SshKey
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *string {
		if v == nil {
			return nil
		}
		return v.StorageType
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o LabVirtualMachineCreationParameterPtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameter) *string {
		if v == nil {
			return nil
		}
		return v.UserName
	}).(pulumi.StringPtrOutput)
}

type LabVirtualMachineCreationParameterResponse struct {
	AllowClaim                 *bool                               `pulumi:"allowClaim"`
	Artifacts                  []ArtifactInstallPropertiesResponse `pulumi:"artifacts"`
	BulkCreationParameters     *BulkCreationParametersResponse     `pulumi:"bulkCreationParameters"`
	CreatedDate                *string                             `pulumi:"createdDate"`
	CustomImageId              *string                             `pulumi:"customImageId"`
	DataDiskParameters         []DataDiskPropertiesResponse        `pulumi:"dataDiskParameters"`
	DisallowPublicIpAddress    *bool                               `pulumi:"disallowPublicIpAddress"`
	EnvironmentId              *string                             `pulumi:"environmentId"`
	ExpirationDate             *string                             `pulumi:"expirationDate"`
	GalleryImageReference      *GalleryImageReferenceResponse      `pulumi:"galleryImageReference"`
	IsAuthenticationWithSshKey *bool                               `pulumi:"isAuthenticationWithSshKey"`
	LabSubnetName              *string                             `pulumi:"labSubnetName"`
	LabVirtualNetworkId        *string                             `pulumi:"labVirtualNetworkId"`
	Location                   *string                             `pulumi:"location"`
	Name                       *string                             `pulumi:"name"`
	NetworkInterface           *NetworkInterfacePropertiesResponse `pulumi:"networkInterface"`
	Notes                      *string                             `pulumi:"notes"`
	OwnerObjectId              *string                             `pulumi:"ownerObjectId"`
	OwnerUserPrincipalName     *string                             `pulumi:"ownerUserPrincipalName"`
	Password                   *string                             `pulumi:"password"`
	PlanId                     *string                             `pulumi:"planId"`
	ScheduleParameters         []ScheduleCreationParameterResponse `pulumi:"scheduleParameters"`
	Size                       *string                             `pulumi:"size"`
	SshKey                     *string                             `pulumi:"sshKey"`
	StorageType                *string                             `pulumi:"storageType"`
	Tags                       map[string]string                   `pulumi:"tags"`
	UserName                   *string                             `pulumi:"userName"`
}


func (val *LabVirtualMachineCreationParameterResponse) Defaults() *LabVirtualMachineCreationParameterResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllowClaim) {
		allowClaim_ := false
		tmp.AllowClaim = &allowClaim_
	}
	if isZero(tmp.DisallowPublicIpAddress) {
		disallowPublicIpAddress_ := false
		tmp.DisallowPublicIpAddress = &disallowPublicIpAddress_
	}
	if isZero(tmp.OwnerObjectId) {
		ownerObjectId_ := "dynamicValue"
		tmp.OwnerObjectId = &ownerObjectId_
	}
	if isZero(tmp.StorageType) {
		storageType_ := "labStorageType"
		tmp.StorageType = &storageType_
	}
	return &tmp
}

type LabVirtualMachineCreationParameterResponseOutput struct{ *pulumi.OutputState }

func (LabVirtualMachineCreationParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabVirtualMachineCreationParameterResponse)(nil)).Elem()
}

func (o LabVirtualMachineCreationParameterResponseOutput) ToLabVirtualMachineCreationParameterResponseOutput() LabVirtualMachineCreationParameterResponseOutput {
	return o
}

func (o LabVirtualMachineCreationParameterResponseOutput) ToLabVirtualMachineCreationParameterResponseOutputWithContext(ctx context.Context) LabVirtualMachineCreationParameterResponseOutput {
	return o
}

func (o LabVirtualMachineCreationParameterResponseOutput) AllowClaim() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *bool { return v.AllowClaim }).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) Artifacts() ArtifactInstallPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) []ArtifactInstallPropertiesResponse {
		return v.Artifacts
	}).(ArtifactInstallPropertiesResponseArrayOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) BulkCreationParameters() BulkCreationParametersResponsePtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *BulkCreationParametersResponse {
		return v.BulkCreationParameters
	}).(BulkCreationParametersResponsePtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *string { return v.CreatedDate }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) CustomImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *string { return v.CustomImageId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) DataDiskParameters() DataDiskPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) []DataDiskPropertiesResponse {
		return v.DataDiskParameters
	}).(DataDiskPropertiesResponseArrayOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) DisallowPublicIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *bool { return v.DisallowPublicIpAddress }).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *string { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *string { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) GalleryImageReference() GalleryImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *GalleryImageReferenceResponse {
		return v.GalleryImageReference
	}).(GalleryImageReferenceResponsePtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) IsAuthenticationWithSshKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *bool { return v.IsAuthenticationWithSshKey }).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *string { return v.LabSubnetName }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) LabVirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *string { return v.LabVirtualNetworkId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) NetworkInterface() NetworkInterfacePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *NetworkInterfacePropertiesResponse {
		return v.NetworkInterface
	}).(NetworkInterfacePropertiesResponsePtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) OwnerObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *string { return v.OwnerObjectId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) OwnerUserPrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *string { return v.OwnerUserPrincipalName }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) PlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *string { return v.PlanId }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) ScheduleParameters() ScheduleCreationParameterResponseArrayOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) []ScheduleCreationParameterResponse {
		return v.ScheduleParameters
	}).(ScheduleCreationParameterResponseArrayOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) SshKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *string { return v.SshKey }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *string { return v.StorageType }).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LabVirtualMachineCreationParameterResponseOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LabVirtualMachineCreationParameterResponse) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

type LabVirtualMachineCreationParameterResponsePtrOutput struct{ *pulumi.OutputState }

func (LabVirtualMachineCreationParameterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabVirtualMachineCreationParameterResponse)(nil)).Elem()
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) ToLabVirtualMachineCreationParameterResponsePtrOutput() LabVirtualMachineCreationParameterResponsePtrOutput {
	return o
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) ToLabVirtualMachineCreationParameterResponsePtrOutputWithContext(ctx context.Context) LabVirtualMachineCreationParameterResponsePtrOutput {
	return o
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) Elem() LabVirtualMachineCreationParameterResponseOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) LabVirtualMachineCreationParameterResponse {
		if v != nil {
			return *v
		}
		var ret LabVirtualMachineCreationParameterResponse
		return ret
	}).(LabVirtualMachineCreationParameterResponseOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) AllowClaim() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AllowClaim
	}).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) Artifacts() ArtifactInstallPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) []ArtifactInstallPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Artifacts
	}).(ArtifactInstallPropertiesResponseArrayOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) BulkCreationParameters() BulkCreationParametersResponsePtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *BulkCreationParametersResponse {
		if v == nil {
			return nil
		}
		return v.BulkCreationParameters
	}).(BulkCreationParametersResponsePtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedDate
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) CustomImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomImageId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) DataDiskParameters() DataDiskPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) []DataDiskPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.DataDiskParameters
	}).(DataDiskPropertiesResponseArrayOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) DisallowPublicIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisallowPublicIpAddress
	}).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *string {
		if v == nil {
			return nil
		}
		return v.EnvironmentId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExpirationDate
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) GalleryImageReference() GalleryImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *GalleryImageReferenceResponse {
		if v == nil {
			return nil
		}
		return v.GalleryImageReference
	}).(GalleryImageReferenceResponsePtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) IsAuthenticationWithSshKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsAuthenticationWithSshKey
	}).(pulumi.BoolPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *string {
		if v == nil {
			return nil
		}
		return v.LabSubnetName
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) LabVirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *string {
		if v == nil {
			return nil
		}
		return v.LabVirtualNetworkId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) NetworkInterface() NetworkInterfacePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *NetworkInterfacePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.NetworkInterface
	}).(NetworkInterfacePropertiesResponsePtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *string {
		if v == nil {
			return nil
		}
		return v.Notes
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) OwnerObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *string {
		if v == nil {
			return nil
		}
		return v.OwnerObjectId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) OwnerUserPrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *string {
		if v == nil {
			return nil
		}
		return v.OwnerUserPrincipalName
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) PlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *string {
		if v == nil {
			return nil
		}
		return v.PlanId
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) ScheduleParameters() ScheduleCreationParameterResponseArrayOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) []ScheduleCreationParameterResponse {
		if v == nil {
			return nil
		}
		return v.ScheduleParameters
	}).(ScheduleCreationParameterResponseArrayOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) SshKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *string {
		if v == nil {
			return nil
		}
		return v.SshKey
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageType
	}).(pulumi.StringPtrOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o LabVirtualMachineCreationParameterResponsePtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabVirtualMachineCreationParameterResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserName
	}).(pulumi.StringPtrOutput)
}

type LinuxOsInfo struct {
	LinuxOsState *string `pulumi:"linuxOsState"`
}





type LinuxOsInfoInput interface {
	pulumi.Input

	ToLinuxOsInfoOutput() LinuxOsInfoOutput
	ToLinuxOsInfoOutputWithContext(context.Context) LinuxOsInfoOutput
}

type LinuxOsInfoArgs struct {
	LinuxOsState pulumi.StringPtrInput `pulumi:"linuxOsState"`
}

func (LinuxOsInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxOsInfo)(nil)).Elem()
}

func (i LinuxOsInfoArgs) ToLinuxOsInfoOutput() LinuxOsInfoOutput {
	return i.ToLinuxOsInfoOutputWithContext(context.Background())
}

func (i LinuxOsInfoArgs) ToLinuxOsInfoOutputWithContext(ctx context.Context) LinuxOsInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxOsInfoOutput)
}

func (i LinuxOsInfoArgs) ToLinuxOsInfoPtrOutput() LinuxOsInfoPtrOutput {
	return i.ToLinuxOsInfoPtrOutputWithContext(context.Background())
}

func (i LinuxOsInfoArgs) ToLinuxOsInfoPtrOutputWithContext(ctx context.Context) LinuxOsInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxOsInfoOutput).ToLinuxOsInfoPtrOutputWithContext(ctx)
}









type LinuxOsInfoPtrInput interface {
	pulumi.Input

	ToLinuxOsInfoPtrOutput() LinuxOsInfoPtrOutput
	ToLinuxOsInfoPtrOutputWithContext(context.Context) LinuxOsInfoPtrOutput
}

type linuxOsInfoPtrType LinuxOsInfoArgs

func LinuxOsInfoPtr(v *LinuxOsInfoArgs) LinuxOsInfoPtrInput {
	return (*linuxOsInfoPtrType)(v)
}

func (*linuxOsInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxOsInfo)(nil)).Elem()
}

func (i *linuxOsInfoPtrType) ToLinuxOsInfoPtrOutput() LinuxOsInfoPtrOutput {
	return i.ToLinuxOsInfoPtrOutputWithContext(context.Background())
}

func (i *linuxOsInfoPtrType) ToLinuxOsInfoPtrOutputWithContext(ctx context.Context) LinuxOsInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxOsInfoPtrOutput)
}

type LinuxOsInfoOutput struct{ *pulumi.OutputState }

func (LinuxOsInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxOsInfo)(nil)).Elem()
}

func (o LinuxOsInfoOutput) ToLinuxOsInfoOutput() LinuxOsInfoOutput {
	return o
}

func (o LinuxOsInfoOutput) ToLinuxOsInfoOutputWithContext(ctx context.Context) LinuxOsInfoOutput {
	return o
}

func (o LinuxOsInfoOutput) ToLinuxOsInfoPtrOutput() LinuxOsInfoPtrOutput {
	return o.ToLinuxOsInfoPtrOutputWithContext(context.Background())
}

func (o LinuxOsInfoOutput) ToLinuxOsInfoPtrOutputWithContext(ctx context.Context) LinuxOsInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinuxOsInfo) *LinuxOsInfo {
		return &v
	}).(LinuxOsInfoPtrOutput)
}

func (o LinuxOsInfoOutput) LinuxOsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxOsInfo) *string { return v.LinuxOsState }).(pulumi.StringPtrOutput)
}

type LinuxOsInfoPtrOutput struct{ *pulumi.OutputState }

func (LinuxOsInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxOsInfo)(nil)).Elem()
}

func (o LinuxOsInfoPtrOutput) ToLinuxOsInfoPtrOutput() LinuxOsInfoPtrOutput {
	return o
}

func (o LinuxOsInfoPtrOutput) ToLinuxOsInfoPtrOutputWithContext(ctx context.Context) LinuxOsInfoPtrOutput {
	return o
}

func (o LinuxOsInfoPtrOutput) Elem() LinuxOsInfoOutput {
	return o.ApplyT(func(v *LinuxOsInfo) LinuxOsInfo {
		if v != nil {
			return *v
		}
		var ret LinuxOsInfo
		return ret
	}).(LinuxOsInfoOutput)
}

func (o LinuxOsInfoPtrOutput) LinuxOsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxOsInfo) *string {
		if v == nil {
			return nil
		}
		return v.LinuxOsState
	}).(pulumi.StringPtrOutput)
}

type LinuxOsInfoResponse struct {
	LinuxOsState *string `pulumi:"linuxOsState"`
}

type LinuxOsInfoResponseOutput struct{ *pulumi.OutputState }

func (LinuxOsInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxOsInfoResponse)(nil)).Elem()
}

func (o LinuxOsInfoResponseOutput) ToLinuxOsInfoResponseOutput() LinuxOsInfoResponseOutput {
	return o
}

func (o LinuxOsInfoResponseOutput) ToLinuxOsInfoResponseOutputWithContext(ctx context.Context) LinuxOsInfoResponseOutput {
	return o
}

func (o LinuxOsInfoResponseOutput) LinuxOsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxOsInfoResponse) *string { return v.LinuxOsState }).(pulumi.StringPtrOutput)
}

type LinuxOsInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (LinuxOsInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxOsInfoResponse)(nil)).Elem()
}

func (o LinuxOsInfoResponsePtrOutput) ToLinuxOsInfoResponsePtrOutput() LinuxOsInfoResponsePtrOutput {
	return o
}

func (o LinuxOsInfoResponsePtrOutput) ToLinuxOsInfoResponsePtrOutputWithContext(ctx context.Context) LinuxOsInfoResponsePtrOutput {
	return o
}

func (o LinuxOsInfoResponsePtrOutput) Elem() LinuxOsInfoResponseOutput {
	return o.ApplyT(func(v *LinuxOsInfoResponse) LinuxOsInfoResponse {
		if v != nil {
			return *v
		}
		var ret LinuxOsInfoResponse
		return ret
	}).(LinuxOsInfoResponseOutput)
}

func (o LinuxOsInfoResponsePtrOutput) LinuxOsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxOsInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.LinuxOsState
	}).(pulumi.StringPtrOutput)
}

type NetworkInterfaceProperties struct {
	DnsName                            *string                             `pulumi:"dnsName"`
	PrivateIpAddress                   *string                             `pulumi:"privateIpAddress"`
	PublicIpAddress                    *string                             `pulumi:"publicIpAddress"`
	PublicIpAddressId                  *string                             `pulumi:"publicIpAddressId"`
	RdpAuthority                       *string                             `pulumi:"rdpAuthority"`
	SharedPublicIpAddressConfiguration *SharedPublicIpAddressConfiguration `pulumi:"sharedPublicIpAddressConfiguration"`
	SshAuthority                       *string                             `pulumi:"sshAuthority"`
	SubnetId                           *string                             `pulumi:"subnetId"`
	VirtualNetworkId                   *string                             `pulumi:"virtualNetworkId"`
}





type NetworkInterfacePropertiesInput interface {
	pulumi.Input

	ToNetworkInterfacePropertiesOutput() NetworkInterfacePropertiesOutput
	ToNetworkInterfacePropertiesOutputWithContext(context.Context) NetworkInterfacePropertiesOutput
}

type NetworkInterfacePropertiesArgs struct {
	DnsName                            pulumi.StringPtrInput                      `pulumi:"dnsName"`
	PrivateIpAddress                   pulumi.StringPtrInput                      `pulumi:"privateIpAddress"`
	PublicIpAddress                    pulumi.StringPtrInput                      `pulumi:"publicIpAddress"`
	PublicIpAddressId                  pulumi.StringPtrInput                      `pulumi:"publicIpAddressId"`
	RdpAuthority                       pulumi.StringPtrInput                      `pulumi:"rdpAuthority"`
	SharedPublicIpAddressConfiguration SharedPublicIpAddressConfigurationPtrInput `pulumi:"sharedPublicIpAddressConfiguration"`
	SshAuthority                       pulumi.StringPtrInput                      `pulumi:"sshAuthority"`
	SubnetId                           pulumi.StringPtrInput                      `pulumi:"subnetId"`
	VirtualNetworkId                   pulumi.StringPtrInput                      `pulumi:"virtualNetworkId"`
}

func (NetworkInterfacePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceProperties)(nil)).Elem()
}

func (i NetworkInterfacePropertiesArgs) ToNetworkInterfacePropertiesOutput() NetworkInterfacePropertiesOutput {
	return i.ToNetworkInterfacePropertiesOutputWithContext(context.Background())
}

func (i NetworkInterfacePropertiesArgs) ToNetworkInterfacePropertiesOutputWithContext(ctx context.Context) NetworkInterfacePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfacePropertiesOutput)
}

func (i NetworkInterfacePropertiesArgs) ToNetworkInterfacePropertiesPtrOutput() NetworkInterfacePropertiesPtrOutput {
	return i.ToNetworkInterfacePropertiesPtrOutputWithContext(context.Background())
}

func (i NetworkInterfacePropertiesArgs) ToNetworkInterfacePropertiesPtrOutputWithContext(ctx context.Context) NetworkInterfacePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfacePropertiesOutput).ToNetworkInterfacePropertiesPtrOutputWithContext(ctx)
}









type NetworkInterfacePropertiesPtrInput interface {
	pulumi.Input

	ToNetworkInterfacePropertiesPtrOutput() NetworkInterfacePropertiesPtrOutput
	ToNetworkInterfacePropertiesPtrOutputWithContext(context.Context) NetworkInterfacePropertiesPtrOutput
}

type networkInterfacePropertiesPtrType NetworkInterfacePropertiesArgs

func NetworkInterfacePropertiesPtr(v *NetworkInterfacePropertiesArgs) NetworkInterfacePropertiesPtrInput {
	return (*networkInterfacePropertiesPtrType)(v)
}

func (*networkInterfacePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceProperties)(nil)).Elem()
}

func (i *networkInterfacePropertiesPtrType) ToNetworkInterfacePropertiesPtrOutput() NetworkInterfacePropertiesPtrOutput {
	return i.ToNetworkInterfacePropertiesPtrOutputWithContext(context.Background())
}

func (i *networkInterfacePropertiesPtrType) ToNetworkInterfacePropertiesPtrOutputWithContext(ctx context.Context) NetworkInterfacePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfacePropertiesPtrOutput)
}

type NetworkInterfacePropertiesOutput struct{ *pulumi.OutputState }

func (NetworkInterfacePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceProperties)(nil)).Elem()
}

func (o NetworkInterfacePropertiesOutput) ToNetworkInterfacePropertiesOutput() NetworkInterfacePropertiesOutput {
	return o
}

func (o NetworkInterfacePropertiesOutput) ToNetworkInterfacePropertiesOutputWithContext(ctx context.Context) NetworkInterfacePropertiesOutput {
	return o
}

func (o NetworkInterfacePropertiesOutput) ToNetworkInterfacePropertiesPtrOutput() NetworkInterfacePropertiesPtrOutput {
	return o.ToNetworkInterfacePropertiesPtrOutputWithContext(context.Background())
}

func (o NetworkInterfacePropertiesOutput) ToNetworkInterfacePropertiesPtrOutputWithContext(ctx context.Context) NetworkInterfacePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkInterfaceProperties) *NetworkInterfaceProperties {
		return &v
	}).(NetworkInterfacePropertiesPtrOutput)
}

func (o NetworkInterfacePropertiesOutput) DnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceProperties) *string { return v.DnsName }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceProperties) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesOutput) PublicIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceProperties) *string { return v.PublicIpAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesOutput) PublicIpAddressId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceProperties) *string { return v.PublicIpAddressId }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesOutput) RdpAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceProperties) *string { return v.RdpAuthority }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesOutput) SharedPublicIpAddressConfiguration() SharedPublicIpAddressConfigurationPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceProperties) *SharedPublicIpAddressConfiguration {
		return v.SharedPublicIpAddressConfiguration
	}).(SharedPublicIpAddressConfigurationPtrOutput)
}

func (o NetworkInterfacePropertiesOutput) SshAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceProperties) *string { return v.SshAuthority }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceProperties) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesOutput) VirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceProperties) *string { return v.VirtualNetworkId }).(pulumi.StringPtrOutput)
}

type NetworkInterfacePropertiesPtrOutput struct{ *pulumi.OutputState }

func (NetworkInterfacePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceProperties)(nil)).Elem()
}

func (o NetworkInterfacePropertiesPtrOutput) ToNetworkInterfacePropertiesPtrOutput() NetworkInterfacePropertiesPtrOutput {
	return o
}

func (o NetworkInterfacePropertiesPtrOutput) ToNetworkInterfacePropertiesPtrOutputWithContext(ctx context.Context) NetworkInterfacePropertiesPtrOutput {
	return o
}

func (o NetworkInterfacePropertiesPtrOutput) Elem() NetworkInterfacePropertiesOutput {
	return o.ApplyT(func(v *NetworkInterfaceProperties) NetworkInterfaceProperties {
		if v != nil {
			return *v
		}
		var ret NetworkInterfaceProperties
		return ret
	}).(NetworkInterfacePropertiesOutput)
}

func (o NetworkInterfacePropertiesPtrOutput) DnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.DnsName
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesPtrOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrivateIpAddress
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesPtrOutput) PublicIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicIpAddress
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesPtrOutput) PublicIpAddressId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicIpAddressId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesPtrOutput) RdpAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.RdpAuthority
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesPtrOutput) SharedPublicIpAddressConfiguration() SharedPublicIpAddressConfigurationPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceProperties) *SharedPublicIpAddressConfiguration {
		if v == nil {
			return nil
		}
		return v.SharedPublicIpAddressConfiguration
	}).(SharedPublicIpAddressConfigurationPtrOutput)
}

func (o NetworkInterfacePropertiesPtrOutput) SshAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.SshAuthority
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesPtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.SubnetId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesPtrOutput) VirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceProperties) *string {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkId
	}).(pulumi.StringPtrOutput)
}

type NetworkInterfacePropertiesResponse struct {
	DnsName                            *string                                     `pulumi:"dnsName"`
	PrivateIpAddress                   *string                                     `pulumi:"privateIpAddress"`
	PublicIpAddress                    *string                                     `pulumi:"publicIpAddress"`
	PublicIpAddressId                  *string                                     `pulumi:"publicIpAddressId"`
	RdpAuthority                       *string                                     `pulumi:"rdpAuthority"`
	SharedPublicIpAddressConfiguration *SharedPublicIpAddressConfigurationResponse `pulumi:"sharedPublicIpAddressConfiguration"`
	SshAuthority                       *string                                     `pulumi:"sshAuthority"`
	SubnetId                           *string                                     `pulumi:"subnetId"`
	VirtualNetworkId                   *string                                     `pulumi:"virtualNetworkId"`
}

type NetworkInterfacePropertiesResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfacePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfacePropertiesResponse)(nil)).Elem()
}

func (o NetworkInterfacePropertiesResponseOutput) ToNetworkInterfacePropertiesResponseOutput() NetworkInterfacePropertiesResponseOutput {
	return o
}

func (o NetworkInterfacePropertiesResponseOutput) ToNetworkInterfacePropertiesResponseOutputWithContext(ctx context.Context) NetworkInterfacePropertiesResponseOutput {
	return o
}

func (o NetworkInterfacePropertiesResponseOutput) DnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacePropertiesResponse) *string { return v.DnsName }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesResponseOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacePropertiesResponse) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesResponseOutput) PublicIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacePropertiesResponse) *string { return v.PublicIpAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesResponseOutput) PublicIpAddressId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacePropertiesResponse) *string { return v.PublicIpAddressId }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesResponseOutput) RdpAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacePropertiesResponse) *string { return v.RdpAuthority }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesResponseOutput) SharedPublicIpAddressConfiguration() SharedPublicIpAddressConfigurationResponsePtrOutput {
	return o.ApplyT(func(v NetworkInterfacePropertiesResponse) *SharedPublicIpAddressConfigurationResponse {
		return v.SharedPublicIpAddressConfiguration
	}).(SharedPublicIpAddressConfigurationResponsePtrOutput)
}

func (o NetworkInterfacePropertiesResponseOutput) SshAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacePropertiesResponse) *string { return v.SshAuthority }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesResponseOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacePropertiesResponse) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesResponseOutput) VirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacePropertiesResponse) *string { return v.VirtualNetworkId }).(pulumi.StringPtrOutput)
}

type NetworkInterfacePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkInterfacePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfacePropertiesResponse)(nil)).Elem()
}

func (o NetworkInterfacePropertiesResponsePtrOutput) ToNetworkInterfacePropertiesResponsePtrOutput() NetworkInterfacePropertiesResponsePtrOutput {
	return o
}

func (o NetworkInterfacePropertiesResponsePtrOutput) ToNetworkInterfacePropertiesResponsePtrOutputWithContext(ctx context.Context) NetworkInterfacePropertiesResponsePtrOutput {
	return o
}

func (o NetworkInterfacePropertiesResponsePtrOutput) Elem() NetworkInterfacePropertiesResponseOutput {
	return o.ApplyT(func(v *NetworkInterfacePropertiesResponse) NetworkInterfacePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret NetworkInterfacePropertiesResponse
		return ret
	}).(NetworkInterfacePropertiesResponseOutput)
}

func (o NetworkInterfacePropertiesResponsePtrOutput) DnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DnsName
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesResponsePtrOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateIpAddress
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesResponsePtrOutput) PublicIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicIpAddress
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesResponsePtrOutput) PublicIpAddressId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicIpAddressId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesResponsePtrOutput) RdpAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RdpAuthority
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesResponsePtrOutput) SharedPublicIpAddressConfiguration() SharedPublicIpAddressConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *NetworkInterfacePropertiesResponse) *SharedPublicIpAddressConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.SharedPublicIpAddressConfiguration
	}).(SharedPublicIpAddressConfigurationResponsePtrOutput)
}

func (o NetworkInterfacePropertiesResponsePtrOutput) SshAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SshAuthority
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesResponsePtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubnetId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacePropertiesResponsePtrOutput) VirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfacePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkId
	}).(pulumi.StringPtrOutput)
}

type NotificationSettings struct {
	EmailRecipient     *string `pulumi:"emailRecipient"`
	NotificationLocale *string `pulumi:"notificationLocale"`
	Status             *string `pulumi:"status"`
	TimeInMinutes      *int    `pulumi:"timeInMinutes"`
	WebhookUrl         *string `pulumi:"webhookUrl"`
}


func (val *NotificationSettings) Defaults() *NotificationSettings {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "Disabled"
		tmp.Status = &status_
	}
	return &tmp
}





type NotificationSettingsInput interface {
	pulumi.Input

	ToNotificationSettingsOutput() NotificationSettingsOutput
	ToNotificationSettingsOutputWithContext(context.Context) NotificationSettingsOutput
}

type NotificationSettingsArgs struct {
	EmailRecipient     pulumi.StringPtrInput `pulumi:"emailRecipient"`
	NotificationLocale pulumi.StringPtrInput `pulumi:"notificationLocale"`
	Status             pulumi.StringPtrInput `pulumi:"status"`
	TimeInMinutes      pulumi.IntPtrInput    `pulumi:"timeInMinutes"`
	WebhookUrl         pulumi.StringPtrInput `pulumi:"webhookUrl"`
}


func (val *NotificationSettingsArgs) Defaults() *NotificationSettingsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		tmp.Status = pulumi.StringPtr("Disabled")
	}
	return &tmp
}
func (NotificationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationSettings)(nil)).Elem()
}

func (i NotificationSettingsArgs) ToNotificationSettingsOutput() NotificationSettingsOutput {
	return i.ToNotificationSettingsOutputWithContext(context.Background())
}

func (i NotificationSettingsArgs) ToNotificationSettingsOutputWithContext(ctx context.Context) NotificationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationSettingsOutput)
}

func (i NotificationSettingsArgs) ToNotificationSettingsPtrOutput() NotificationSettingsPtrOutput {
	return i.ToNotificationSettingsPtrOutputWithContext(context.Background())
}

func (i NotificationSettingsArgs) ToNotificationSettingsPtrOutputWithContext(ctx context.Context) NotificationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationSettingsOutput).ToNotificationSettingsPtrOutputWithContext(ctx)
}









type NotificationSettingsPtrInput interface {
	pulumi.Input

	ToNotificationSettingsPtrOutput() NotificationSettingsPtrOutput
	ToNotificationSettingsPtrOutputWithContext(context.Context) NotificationSettingsPtrOutput
}

type notificationSettingsPtrType NotificationSettingsArgs

func NotificationSettingsPtr(v *NotificationSettingsArgs) NotificationSettingsPtrInput {
	return (*notificationSettingsPtrType)(v)
}

func (*notificationSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationSettings)(nil)).Elem()
}

func (i *notificationSettingsPtrType) ToNotificationSettingsPtrOutput() NotificationSettingsPtrOutput {
	return i.ToNotificationSettingsPtrOutputWithContext(context.Background())
}

func (i *notificationSettingsPtrType) ToNotificationSettingsPtrOutputWithContext(ctx context.Context) NotificationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationSettingsPtrOutput)
}

type NotificationSettingsOutput struct{ *pulumi.OutputState }

func (NotificationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationSettings)(nil)).Elem()
}

func (o NotificationSettingsOutput) ToNotificationSettingsOutput() NotificationSettingsOutput {
	return o
}

func (o NotificationSettingsOutput) ToNotificationSettingsOutputWithContext(ctx context.Context) NotificationSettingsOutput {
	return o
}

func (o NotificationSettingsOutput) ToNotificationSettingsPtrOutput() NotificationSettingsPtrOutput {
	return o.ToNotificationSettingsPtrOutputWithContext(context.Background())
}

func (o NotificationSettingsOutput) ToNotificationSettingsPtrOutputWithContext(ctx context.Context) NotificationSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotificationSettings) *NotificationSettings {
		return &v
	}).(NotificationSettingsPtrOutput)
}

func (o NotificationSettingsOutput) EmailRecipient() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationSettings) *string { return v.EmailRecipient }).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsOutput) NotificationLocale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationSettings) *string { return v.NotificationLocale }).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationSettings) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsOutput) TimeInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NotificationSettings) *int { return v.TimeInMinutes }).(pulumi.IntPtrOutput)
}

func (o NotificationSettingsOutput) WebhookUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationSettings) *string { return v.WebhookUrl }).(pulumi.StringPtrOutput)
}

type NotificationSettingsPtrOutput struct{ *pulumi.OutputState }

func (NotificationSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationSettings)(nil)).Elem()
}

func (o NotificationSettingsPtrOutput) ToNotificationSettingsPtrOutput() NotificationSettingsPtrOutput {
	return o
}

func (o NotificationSettingsPtrOutput) ToNotificationSettingsPtrOutputWithContext(ctx context.Context) NotificationSettingsPtrOutput {
	return o
}

func (o NotificationSettingsPtrOutput) Elem() NotificationSettingsOutput {
	return o.ApplyT(func(v *NotificationSettings) NotificationSettings {
		if v != nil {
			return *v
		}
		var ret NotificationSettings
		return ret
	}).(NotificationSettingsOutput)
}

func (o NotificationSettingsPtrOutput) EmailRecipient() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationSettings) *string {
		if v == nil {
			return nil
		}
		return v.EmailRecipient
	}).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsPtrOutput) NotificationLocale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationSettings) *string {
		if v == nil {
			return nil
		}
		return v.NotificationLocale
	}).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationSettings) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsPtrOutput) TimeInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NotificationSettings) *int {
		if v == nil {
			return nil
		}
		return v.TimeInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o NotificationSettingsPtrOutput) WebhookUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationSettings) *string {
		if v == nil {
			return nil
		}
		return v.WebhookUrl
	}).(pulumi.StringPtrOutput)
}

type NotificationSettingsResponse struct {
	EmailRecipient     *string `pulumi:"emailRecipient"`
	NotificationLocale *string `pulumi:"notificationLocale"`
	Status             *string `pulumi:"status"`
	TimeInMinutes      *int    `pulumi:"timeInMinutes"`
	WebhookUrl         *string `pulumi:"webhookUrl"`
}


func (val *NotificationSettingsResponse) Defaults() *NotificationSettingsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "Disabled"
		tmp.Status = &status_
	}
	return &tmp
}

type NotificationSettingsResponseOutput struct{ *pulumi.OutputState }

func (NotificationSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationSettingsResponse)(nil)).Elem()
}

func (o NotificationSettingsResponseOutput) ToNotificationSettingsResponseOutput() NotificationSettingsResponseOutput {
	return o
}

func (o NotificationSettingsResponseOutput) ToNotificationSettingsResponseOutputWithContext(ctx context.Context) NotificationSettingsResponseOutput {
	return o
}

func (o NotificationSettingsResponseOutput) EmailRecipient() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationSettingsResponse) *string { return v.EmailRecipient }).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsResponseOutput) NotificationLocale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationSettingsResponse) *string { return v.NotificationLocale }).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationSettingsResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsResponseOutput) TimeInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NotificationSettingsResponse) *int { return v.TimeInMinutes }).(pulumi.IntPtrOutput)
}

func (o NotificationSettingsResponseOutput) WebhookUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationSettingsResponse) *string { return v.WebhookUrl }).(pulumi.StringPtrOutput)
}

type NotificationSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (NotificationSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationSettingsResponse)(nil)).Elem()
}

func (o NotificationSettingsResponsePtrOutput) ToNotificationSettingsResponsePtrOutput() NotificationSettingsResponsePtrOutput {
	return o
}

func (o NotificationSettingsResponsePtrOutput) ToNotificationSettingsResponsePtrOutputWithContext(ctx context.Context) NotificationSettingsResponsePtrOutput {
	return o
}

func (o NotificationSettingsResponsePtrOutput) Elem() NotificationSettingsResponseOutput {
	return o.ApplyT(func(v *NotificationSettingsResponse) NotificationSettingsResponse {
		if v != nil {
			return *v
		}
		var ret NotificationSettingsResponse
		return ret
	}).(NotificationSettingsResponseOutput)
}

func (o NotificationSettingsResponsePtrOutput) EmailRecipient() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.EmailRecipient
	}).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsResponsePtrOutput) NotificationLocale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.NotificationLocale
	}).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o NotificationSettingsResponsePtrOutput) TimeInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NotificationSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.TimeInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o NotificationSettingsResponsePtrOutput) WebhookUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.WebhookUrl
	}).(pulumi.StringPtrOutput)
}

type Port struct {
	BackendPort       *int    `pulumi:"backendPort"`
	TransportProtocol *string `pulumi:"transportProtocol"`
}





type PortInput interface {
	pulumi.Input

	ToPortOutput() PortOutput
	ToPortOutputWithContext(context.Context) PortOutput
}

type PortArgs struct {
	BackendPort       pulumi.IntPtrInput    `pulumi:"backendPort"`
	TransportProtocol pulumi.StringPtrInput `pulumi:"transportProtocol"`
}

func (PortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Port)(nil)).Elem()
}

func (i PortArgs) ToPortOutput() PortOutput {
	return i.ToPortOutputWithContext(context.Background())
}

func (i PortArgs) ToPortOutputWithContext(ctx context.Context) PortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortOutput)
}





type PortArrayInput interface {
	pulumi.Input

	ToPortArrayOutput() PortArrayOutput
	ToPortArrayOutputWithContext(context.Context) PortArrayOutput
}

type PortArray []PortInput

func (PortArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Port)(nil)).Elem()
}

func (i PortArray) ToPortArrayOutput() PortArrayOutput {
	return i.ToPortArrayOutputWithContext(context.Background())
}

func (i PortArray) ToPortArrayOutputWithContext(ctx context.Context) PortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortArrayOutput)
}

type PortOutput struct{ *pulumi.OutputState }

func (PortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Port)(nil)).Elem()
}

func (o PortOutput) ToPortOutput() PortOutput {
	return o
}

func (o PortOutput) ToPortOutputWithContext(ctx context.Context) PortOutput {
	return o
}

func (o PortOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Port) *int { return v.BackendPort }).(pulumi.IntPtrOutput)
}

func (o PortOutput) TransportProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Port) *string { return v.TransportProtocol }).(pulumi.StringPtrOutput)
}

type PortArrayOutput struct{ *pulumi.OutputState }

func (PortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Port)(nil)).Elem()
}

func (o PortArrayOutput) ToPortArrayOutput() PortArrayOutput {
	return o
}

func (o PortArrayOutput) ToPortArrayOutputWithContext(ctx context.Context) PortArrayOutput {
	return o
}

func (o PortArrayOutput) Index(i pulumi.IntInput) PortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Port {
		return vs[0].([]Port)[vs[1].(int)]
	}).(PortOutput)
}

type PortResponse struct {
	BackendPort       *int    `pulumi:"backendPort"`
	TransportProtocol *string `pulumi:"transportProtocol"`
}

type PortResponseOutput struct{ *pulumi.OutputState }

func (PortResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PortResponse)(nil)).Elem()
}

func (o PortResponseOutput) ToPortResponseOutput() PortResponseOutput {
	return o
}

func (o PortResponseOutput) ToPortResponseOutputWithContext(ctx context.Context) PortResponseOutput {
	return o
}

func (o PortResponseOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PortResponse) *int { return v.BackendPort }).(pulumi.IntPtrOutput)
}

func (o PortResponseOutput) TransportProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PortResponse) *string { return v.TransportProtocol }).(pulumi.StringPtrOutput)
}

type PortResponseArrayOutput struct{ *pulumi.OutputState }

func (PortResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PortResponse)(nil)).Elem()
}

func (o PortResponseArrayOutput) ToPortResponseArrayOutput() PortResponseArrayOutput {
	return o
}

func (o PortResponseArrayOutput) ToPortResponseArrayOutputWithContext(ctx context.Context) PortResponseArrayOutput {
	return o
}

func (o PortResponseArrayOutput) Index(i pulumi.IntInput) PortResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PortResponse {
		return vs[0].([]PortResponse)[vs[1].(int)]
	}).(PortResponseOutput)
}

type ScheduleCreationParameter struct {
	DailyRecurrence      *DayDetails           `pulumi:"dailyRecurrence"`
	HourlyRecurrence     *HourDetails          `pulumi:"hourlyRecurrence"`
	Name                 *string               `pulumi:"name"`
	NotificationSettings *NotificationSettings `pulumi:"notificationSettings"`
	Status               *string               `pulumi:"status"`
	Tags                 map[string]string     `pulumi:"tags"`
	TargetResourceId     *string               `pulumi:"targetResourceId"`
	TaskType             *string               `pulumi:"taskType"`
	TimeZoneId           *string               `pulumi:"timeZoneId"`
	WeeklyRecurrence     *WeekDetails          `pulumi:"weeklyRecurrence"`
}


func (val *ScheduleCreationParameter) Defaults() *ScheduleCreationParameter {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NotificationSettings = tmp.NotificationSettings.Defaults()

	if isZero(tmp.Status) {
		status_ := "Disabled"
		tmp.Status = &status_
	}
	return &tmp
}





type ScheduleCreationParameterInput interface {
	pulumi.Input

	ToScheduleCreationParameterOutput() ScheduleCreationParameterOutput
	ToScheduleCreationParameterOutputWithContext(context.Context) ScheduleCreationParameterOutput
}

type ScheduleCreationParameterArgs struct {
	DailyRecurrence      DayDetailsPtrInput           `pulumi:"dailyRecurrence"`
	HourlyRecurrence     HourDetailsPtrInput          `pulumi:"hourlyRecurrence"`
	Name                 pulumi.StringPtrInput        `pulumi:"name"`
	NotificationSettings NotificationSettingsPtrInput `pulumi:"notificationSettings"`
	Status               pulumi.StringPtrInput        `pulumi:"status"`
	Tags                 pulumi.StringMapInput        `pulumi:"tags"`
	TargetResourceId     pulumi.StringPtrInput        `pulumi:"targetResourceId"`
	TaskType             pulumi.StringPtrInput        `pulumi:"taskType"`
	TimeZoneId           pulumi.StringPtrInput        `pulumi:"timeZoneId"`
	WeeklyRecurrence     WeekDetailsPtrInput          `pulumi:"weeklyRecurrence"`
}


func (val *ScheduleCreationParameterArgs) Defaults() *ScheduleCreationParameterArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	if isZero(tmp.Status) {
		tmp.Status = pulumi.StringPtr("Disabled")
	}
	return &tmp
}
func (ScheduleCreationParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleCreationParameter)(nil)).Elem()
}

func (i ScheduleCreationParameterArgs) ToScheduleCreationParameterOutput() ScheduleCreationParameterOutput {
	return i.ToScheduleCreationParameterOutputWithContext(context.Background())
}

func (i ScheduleCreationParameterArgs) ToScheduleCreationParameterOutputWithContext(ctx context.Context) ScheduleCreationParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleCreationParameterOutput)
}





type ScheduleCreationParameterArrayInput interface {
	pulumi.Input

	ToScheduleCreationParameterArrayOutput() ScheduleCreationParameterArrayOutput
	ToScheduleCreationParameterArrayOutputWithContext(context.Context) ScheduleCreationParameterArrayOutput
}

type ScheduleCreationParameterArray []ScheduleCreationParameterInput

func (ScheduleCreationParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleCreationParameter)(nil)).Elem()
}

func (i ScheduleCreationParameterArray) ToScheduleCreationParameterArrayOutput() ScheduleCreationParameterArrayOutput {
	return i.ToScheduleCreationParameterArrayOutputWithContext(context.Background())
}

func (i ScheduleCreationParameterArray) ToScheduleCreationParameterArrayOutputWithContext(ctx context.Context) ScheduleCreationParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleCreationParameterArrayOutput)
}

type ScheduleCreationParameterOutput struct{ *pulumi.OutputState }

func (ScheduleCreationParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleCreationParameter)(nil)).Elem()
}

func (o ScheduleCreationParameterOutput) ToScheduleCreationParameterOutput() ScheduleCreationParameterOutput {
	return o
}

func (o ScheduleCreationParameterOutput) ToScheduleCreationParameterOutputWithContext(ctx context.Context) ScheduleCreationParameterOutput {
	return o
}

func (o ScheduleCreationParameterOutput) DailyRecurrence() DayDetailsPtrOutput {
	return o.ApplyT(func(v ScheduleCreationParameter) *DayDetails { return v.DailyRecurrence }).(DayDetailsPtrOutput)
}

func (o ScheduleCreationParameterOutput) HourlyRecurrence() HourDetailsPtrOutput {
	return o.ApplyT(func(v ScheduleCreationParameter) *HourDetails { return v.HourlyRecurrence }).(HourDetailsPtrOutput)
}

func (o ScheduleCreationParameterOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleCreationParameter) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ScheduleCreationParameterOutput) NotificationSettings() NotificationSettingsPtrOutput {
	return o.ApplyT(func(v ScheduleCreationParameter) *NotificationSettings { return v.NotificationSettings }).(NotificationSettingsPtrOutput)
}

func (o ScheduleCreationParameterOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleCreationParameter) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ScheduleCreationParameterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ScheduleCreationParameter) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ScheduleCreationParameterOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleCreationParameter) *string { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o ScheduleCreationParameterOutput) TaskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleCreationParameter) *string { return v.TaskType }).(pulumi.StringPtrOutput)
}

func (o ScheduleCreationParameterOutput) TimeZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleCreationParameter) *string { return v.TimeZoneId }).(pulumi.StringPtrOutput)
}

func (o ScheduleCreationParameterOutput) WeeklyRecurrence() WeekDetailsPtrOutput {
	return o.ApplyT(func(v ScheduleCreationParameter) *WeekDetails { return v.WeeklyRecurrence }).(WeekDetailsPtrOutput)
}

type ScheduleCreationParameterArrayOutput struct{ *pulumi.OutputState }

func (ScheduleCreationParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleCreationParameter)(nil)).Elem()
}

func (o ScheduleCreationParameterArrayOutput) ToScheduleCreationParameterArrayOutput() ScheduleCreationParameterArrayOutput {
	return o
}

func (o ScheduleCreationParameterArrayOutput) ToScheduleCreationParameterArrayOutputWithContext(ctx context.Context) ScheduleCreationParameterArrayOutput {
	return o
}

func (o ScheduleCreationParameterArrayOutput) Index(i pulumi.IntInput) ScheduleCreationParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScheduleCreationParameter {
		return vs[0].([]ScheduleCreationParameter)[vs[1].(int)]
	}).(ScheduleCreationParameterOutput)
}

type ScheduleCreationParameterResponse struct {
	DailyRecurrence      *DayDetailsResponse           `pulumi:"dailyRecurrence"`
	HourlyRecurrence     *HourDetailsResponse          `pulumi:"hourlyRecurrence"`
	Location             string                        `pulumi:"location"`
	Name                 *string                       `pulumi:"name"`
	NotificationSettings *NotificationSettingsResponse `pulumi:"notificationSettings"`
	Status               *string                       `pulumi:"status"`
	Tags                 map[string]string             `pulumi:"tags"`
	TargetResourceId     *string                       `pulumi:"targetResourceId"`
	TaskType             *string                       `pulumi:"taskType"`
	TimeZoneId           *string                       `pulumi:"timeZoneId"`
	WeeklyRecurrence     *WeekDetailsResponse          `pulumi:"weeklyRecurrence"`
}


func (val *ScheduleCreationParameterResponse) Defaults() *ScheduleCreationParameterResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NotificationSettings = tmp.NotificationSettings.Defaults()

	if isZero(tmp.Status) {
		status_ := "Disabled"
		tmp.Status = &status_
	}
	return &tmp
}

type ScheduleCreationParameterResponseOutput struct{ *pulumi.OutputState }

func (ScheduleCreationParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleCreationParameterResponse)(nil)).Elem()
}

func (o ScheduleCreationParameterResponseOutput) ToScheduleCreationParameterResponseOutput() ScheduleCreationParameterResponseOutput {
	return o
}

func (o ScheduleCreationParameterResponseOutput) ToScheduleCreationParameterResponseOutputWithContext(ctx context.Context) ScheduleCreationParameterResponseOutput {
	return o
}

func (o ScheduleCreationParameterResponseOutput) DailyRecurrence() DayDetailsResponsePtrOutput {
	return o.ApplyT(func(v ScheduleCreationParameterResponse) *DayDetailsResponse { return v.DailyRecurrence }).(DayDetailsResponsePtrOutput)
}

func (o ScheduleCreationParameterResponseOutput) HourlyRecurrence() HourDetailsResponsePtrOutput {
	return o.ApplyT(func(v ScheduleCreationParameterResponse) *HourDetailsResponse { return v.HourlyRecurrence }).(HourDetailsResponsePtrOutput)
}

func (o ScheduleCreationParameterResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleCreationParameterResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o ScheduleCreationParameterResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleCreationParameterResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ScheduleCreationParameterResponseOutput) NotificationSettings() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v ScheduleCreationParameterResponse) *NotificationSettingsResponse { return v.NotificationSettings }).(NotificationSettingsResponsePtrOutput)
}

func (o ScheduleCreationParameterResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleCreationParameterResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ScheduleCreationParameterResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ScheduleCreationParameterResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ScheduleCreationParameterResponseOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleCreationParameterResponse) *string { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o ScheduleCreationParameterResponseOutput) TaskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleCreationParameterResponse) *string { return v.TaskType }).(pulumi.StringPtrOutput)
}

func (o ScheduleCreationParameterResponseOutput) TimeZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleCreationParameterResponse) *string { return v.TimeZoneId }).(pulumi.StringPtrOutput)
}

func (o ScheduleCreationParameterResponseOutput) WeeklyRecurrence() WeekDetailsResponsePtrOutput {
	return o.ApplyT(func(v ScheduleCreationParameterResponse) *WeekDetailsResponse { return v.WeeklyRecurrence }).(WeekDetailsResponsePtrOutput)
}

type ScheduleCreationParameterResponseArrayOutput struct{ *pulumi.OutputState }

func (ScheduleCreationParameterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleCreationParameterResponse)(nil)).Elem()
}

func (o ScheduleCreationParameterResponseArrayOutput) ToScheduleCreationParameterResponseArrayOutput() ScheduleCreationParameterResponseArrayOutput {
	return o
}

func (o ScheduleCreationParameterResponseArrayOutput) ToScheduleCreationParameterResponseArrayOutputWithContext(ctx context.Context) ScheduleCreationParameterResponseArrayOutput {
	return o
}

func (o ScheduleCreationParameterResponseArrayOutput) Index(i pulumi.IntInput) ScheduleCreationParameterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScheduleCreationParameterResponse {
		return vs[0].([]ScheduleCreationParameterResponse)[vs[1].(int)]
	}).(ScheduleCreationParameterResponseOutput)
}

type ScheduleResponse struct {
	CreatedDate          string                        `pulumi:"createdDate"`
	DailyRecurrence      *DayDetailsResponse           `pulumi:"dailyRecurrence"`
	HourlyRecurrence     *HourDetailsResponse          `pulumi:"hourlyRecurrence"`
	Id                   string                        `pulumi:"id"`
	Location             *string                       `pulumi:"location"`
	Name                 string                        `pulumi:"name"`
	NotificationSettings *NotificationSettingsResponse `pulumi:"notificationSettings"`
	ProvisioningState    string                        `pulumi:"provisioningState"`
	Status               *string                       `pulumi:"status"`
	Tags                 map[string]string             `pulumi:"tags"`
	TargetResourceId     *string                       `pulumi:"targetResourceId"`
	TaskType             *string                       `pulumi:"taskType"`
	TimeZoneId           *string                       `pulumi:"timeZoneId"`
	Type                 string                        `pulumi:"type"`
	UniqueIdentifier     string                        `pulumi:"uniqueIdentifier"`
	WeeklyRecurrence     *WeekDetailsResponse          `pulumi:"weeklyRecurrence"`
}


func (val *ScheduleResponse) Defaults() *ScheduleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NotificationSettings = tmp.NotificationSettings.Defaults()

	if isZero(tmp.Status) {
		status_ := "Disabled"
		tmp.Status = &status_
	}
	return &tmp
}

type ScheduleResponseOutput struct{ *pulumi.OutputState }

func (ScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleResponse)(nil)).Elem()
}

func (o ScheduleResponseOutput) ToScheduleResponseOutput() ScheduleResponseOutput {
	return o
}

func (o ScheduleResponseOutput) ToScheduleResponseOutputWithContext(ctx context.Context) ScheduleResponseOutput {
	return o
}

func (o ScheduleResponseOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleResponse) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o ScheduleResponseOutput) DailyRecurrence() DayDetailsResponsePtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *DayDetailsResponse { return v.DailyRecurrence }).(DayDetailsResponsePtrOutput)
}

func (o ScheduleResponseOutput) HourlyRecurrence() HourDetailsResponsePtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *HourDetailsResponse { return v.HourlyRecurrence }).(HourDetailsResponsePtrOutput)
}

func (o ScheduleResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ScheduleResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ScheduleResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ScheduleResponseOutput) NotificationSettings() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *NotificationSettingsResponse { return v.NotificationSettings }).(NotificationSettingsResponsePtrOutput)
}

func (o ScheduleResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ScheduleResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ScheduleResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ScheduleResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ScheduleResponseOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *string { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o ScheduleResponseOutput) TaskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *string { return v.TaskType }).(pulumi.StringPtrOutput)
}

func (o ScheduleResponseOutput) TimeZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *string { return v.TimeZoneId }).(pulumi.StringPtrOutput)
}

func (o ScheduleResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ScheduleResponseOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleResponse) string { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func (o ScheduleResponseOutput) WeeklyRecurrence() WeekDetailsResponsePtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *WeekDetailsResponse { return v.WeeklyRecurrence }).(WeekDetailsResponsePtrOutput)
}

type ScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (ScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleResponse)(nil)).Elem()
}

func (o ScheduleResponsePtrOutput) ToScheduleResponsePtrOutput() ScheduleResponsePtrOutput {
	return o
}

func (o ScheduleResponsePtrOutput) ToScheduleResponsePtrOutputWithContext(ctx context.Context) ScheduleResponsePtrOutput {
	return o
}

func (o ScheduleResponsePtrOutput) Elem() ScheduleResponseOutput {
	return o.ApplyT(func(v *ScheduleResponse) ScheduleResponse {
		if v != nil {
			return *v
		}
		var ret ScheduleResponse
		return ret
	}).(ScheduleResponseOutput)
}

func (o ScheduleResponsePtrOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedDate
	}).(pulumi.StringPtrOutput)
}

func (o ScheduleResponsePtrOutput) DailyRecurrence() DayDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *DayDetailsResponse {
		if v == nil {
			return nil
		}
		return v.DailyRecurrence
	}).(DayDetailsResponsePtrOutput)
}

func (o ScheduleResponsePtrOutput) HourlyRecurrence() HourDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *HourDetailsResponse {
		if v == nil {
			return nil
		}
		return v.HourlyRecurrence
	}).(HourDetailsResponsePtrOutput)
}

func (o ScheduleResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ScheduleResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o ScheduleResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ScheduleResponsePtrOutput) NotificationSettings() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *NotificationSettingsResponse {
		if v == nil {
			return nil
		}
		return v.NotificationSettings
	}).(NotificationSettingsResponsePtrOutput)
}

func (o ScheduleResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o ScheduleResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o ScheduleResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ScheduleResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o ScheduleResponsePtrOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetResourceId
	}).(pulumi.StringPtrOutput)
}

func (o ScheduleResponsePtrOutput) TaskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.TaskType
	}).(pulumi.StringPtrOutput)
}

func (o ScheduleResponsePtrOutput) TimeZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeZoneId
	}).(pulumi.StringPtrOutput)
}

func (o ScheduleResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ScheduleResponsePtrOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UniqueIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o ScheduleResponsePtrOutput) WeeklyRecurrence() WeekDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *WeekDetailsResponse {
		if v == nil {
			return nil
		}
		return v.WeeklyRecurrence
	}).(WeekDetailsResponsePtrOutput)
}

type ScheduleResponseArrayOutput struct{ *pulumi.OutputState }

func (ScheduleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleResponse)(nil)).Elem()
}

func (o ScheduleResponseArrayOutput) ToScheduleResponseArrayOutput() ScheduleResponseArrayOutput {
	return o
}

func (o ScheduleResponseArrayOutput) ToScheduleResponseArrayOutputWithContext(ctx context.Context) ScheduleResponseArrayOutput {
	return o
}

func (o ScheduleResponseArrayOutput) Index(i pulumi.IntInput) ScheduleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScheduleResponse {
		return vs[0].([]ScheduleResponse)[vs[1].(int)]
	}).(ScheduleResponseOutput)
}

type SharedPublicIpAddressConfiguration struct {
	InboundNatRules []InboundNatRule `pulumi:"inboundNatRules"`
}





type SharedPublicIpAddressConfigurationInput interface {
	pulumi.Input

	ToSharedPublicIpAddressConfigurationOutput() SharedPublicIpAddressConfigurationOutput
	ToSharedPublicIpAddressConfigurationOutputWithContext(context.Context) SharedPublicIpAddressConfigurationOutput
}

type SharedPublicIpAddressConfigurationArgs struct {
	InboundNatRules InboundNatRuleArrayInput `pulumi:"inboundNatRules"`
}

func (SharedPublicIpAddressConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPublicIpAddressConfiguration)(nil)).Elem()
}

func (i SharedPublicIpAddressConfigurationArgs) ToSharedPublicIpAddressConfigurationOutput() SharedPublicIpAddressConfigurationOutput {
	return i.ToSharedPublicIpAddressConfigurationOutputWithContext(context.Background())
}

func (i SharedPublicIpAddressConfigurationArgs) ToSharedPublicIpAddressConfigurationOutputWithContext(ctx context.Context) SharedPublicIpAddressConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPublicIpAddressConfigurationOutput)
}

func (i SharedPublicIpAddressConfigurationArgs) ToSharedPublicIpAddressConfigurationPtrOutput() SharedPublicIpAddressConfigurationPtrOutput {
	return i.ToSharedPublicIpAddressConfigurationPtrOutputWithContext(context.Background())
}

func (i SharedPublicIpAddressConfigurationArgs) ToSharedPublicIpAddressConfigurationPtrOutputWithContext(ctx context.Context) SharedPublicIpAddressConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPublicIpAddressConfigurationOutput).ToSharedPublicIpAddressConfigurationPtrOutputWithContext(ctx)
}









type SharedPublicIpAddressConfigurationPtrInput interface {
	pulumi.Input

	ToSharedPublicIpAddressConfigurationPtrOutput() SharedPublicIpAddressConfigurationPtrOutput
	ToSharedPublicIpAddressConfigurationPtrOutputWithContext(context.Context) SharedPublicIpAddressConfigurationPtrOutput
}

type sharedPublicIpAddressConfigurationPtrType SharedPublicIpAddressConfigurationArgs

func SharedPublicIpAddressConfigurationPtr(v *SharedPublicIpAddressConfigurationArgs) SharedPublicIpAddressConfigurationPtrInput {
	return (*sharedPublicIpAddressConfigurationPtrType)(v)
}

func (*sharedPublicIpAddressConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedPublicIpAddressConfiguration)(nil)).Elem()
}

func (i *sharedPublicIpAddressConfigurationPtrType) ToSharedPublicIpAddressConfigurationPtrOutput() SharedPublicIpAddressConfigurationPtrOutput {
	return i.ToSharedPublicIpAddressConfigurationPtrOutputWithContext(context.Background())
}

func (i *sharedPublicIpAddressConfigurationPtrType) ToSharedPublicIpAddressConfigurationPtrOutputWithContext(ctx context.Context) SharedPublicIpAddressConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedPublicIpAddressConfigurationPtrOutput)
}

type SharedPublicIpAddressConfigurationOutput struct{ *pulumi.OutputState }

func (SharedPublicIpAddressConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPublicIpAddressConfiguration)(nil)).Elem()
}

func (o SharedPublicIpAddressConfigurationOutput) ToSharedPublicIpAddressConfigurationOutput() SharedPublicIpAddressConfigurationOutput {
	return o
}

func (o SharedPublicIpAddressConfigurationOutput) ToSharedPublicIpAddressConfigurationOutputWithContext(ctx context.Context) SharedPublicIpAddressConfigurationOutput {
	return o
}

func (o SharedPublicIpAddressConfigurationOutput) ToSharedPublicIpAddressConfigurationPtrOutput() SharedPublicIpAddressConfigurationPtrOutput {
	return o.ToSharedPublicIpAddressConfigurationPtrOutputWithContext(context.Background())
}

func (o SharedPublicIpAddressConfigurationOutput) ToSharedPublicIpAddressConfigurationPtrOutputWithContext(ctx context.Context) SharedPublicIpAddressConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SharedPublicIpAddressConfiguration) *SharedPublicIpAddressConfiguration {
		return &v
	}).(SharedPublicIpAddressConfigurationPtrOutput)
}

func (o SharedPublicIpAddressConfigurationOutput) InboundNatRules() InboundNatRuleArrayOutput {
	return o.ApplyT(func(v SharedPublicIpAddressConfiguration) []InboundNatRule { return v.InboundNatRules }).(InboundNatRuleArrayOutput)
}

type SharedPublicIpAddressConfigurationPtrOutput struct{ *pulumi.OutputState }

func (SharedPublicIpAddressConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedPublicIpAddressConfiguration)(nil)).Elem()
}

func (o SharedPublicIpAddressConfigurationPtrOutput) ToSharedPublicIpAddressConfigurationPtrOutput() SharedPublicIpAddressConfigurationPtrOutput {
	return o
}

func (o SharedPublicIpAddressConfigurationPtrOutput) ToSharedPublicIpAddressConfigurationPtrOutputWithContext(ctx context.Context) SharedPublicIpAddressConfigurationPtrOutput {
	return o
}

func (o SharedPublicIpAddressConfigurationPtrOutput) Elem() SharedPublicIpAddressConfigurationOutput {
	return o.ApplyT(func(v *SharedPublicIpAddressConfiguration) SharedPublicIpAddressConfiguration {
		if v != nil {
			return *v
		}
		var ret SharedPublicIpAddressConfiguration
		return ret
	}).(SharedPublicIpAddressConfigurationOutput)
}

func (o SharedPublicIpAddressConfigurationPtrOutput) InboundNatRules() InboundNatRuleArrayOutput {
	return o.ApplyT(func(v *SharedPublicIpAddressConfiguration) []InboundNatRule {
		if v == nil {
			return nil
		}
		return v.InboundNatRules
	}).(InboundNatRuleArrayOutput)
}

type SharedPublicIpAddressConfigurationResponse struct {
	InboundNatRules []InboundNatRuleResponse `pulumi:"inboundNatRules"`
}

type SharedPublicIpAddressConfigurationResponseOutput struct{ *pulumi.OutputState }

func (SharedPublicIpAddressConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedPublicIpAddressConfigurationResponse)(nil)).Elem()
}

func (o SharedPublicIpAddressConfigurationResponseOutput) ToSharedPublicIpAddressConfigurationResponseOutput() SharedPublicIpAddressConfigurationResponseOutput {
	return o
}

func (o SharedPublicIpAddressConfigurationResponseOutput) ToSharedPublicIpAddressConfigurationResponseOutputWithContext(ctx context.Context) SharedPublicIpAddressConfigurationResponseOutput {
	return o
}

func (o SharedPublicIpAddressConfigurationResponseOutput) InboundNatRules() InboundNatRuleResponseArrayOutput {
	return o.ApplyT(func(v SharedPublicIpAddressConfigurationResponse) []InboundNatRuleResponse { return v.InboundNatRules }).(InboundNatRuleResponseArrayOutput)
}

type SharedPublicIpAddressConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (SharedPublicIpAddressConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SharedPublicIpAddressConfigurationResponse)(nil)).Elem()
}

func (o SharedPublicIpAddressConfigurationResponsePtrOutput) ToSharedPublicIpAddressConfigurationResponsePtrOutput() SharedPublicIpAddressConfigurationResponsePtrOutput {
	return o
}

func (o SharedPublicIpAddressConfigurationResponsePtrOutput) ToSharedPublicIpAddressConfigurationResponsePtrOutputWithContext(ctx context.Context) SharedPublicIpAddressConfigurationResponsePtrOutput {
	return o
}

func (o SharedPublicIpAddressConfigurationResponsePtrOutput) Elem() SharedPublicIpAddressConfigurationResponseOutput {
	return o.ApplyT(func(v *SharedPublicIpAddressConfigurationResponse) SharedPublicIpAddressConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret SharedPublicIpAddressConfigurationResponse
		return ret
	}).(SharedPublicIpAddressConfigurationResponseOutput)
}

func (o SharedPublicIpAddressConfigurationResponsePtrOutput) InboundNatRules() InboundNatRuleResponseArrayOutput {
	return o.ApplyT(func(v *SharedPublicIpAddressConfigurationResponse) []InboundNatRuleResponse {
		if v == nil {
			return nil
		}
		return v.InboundNatRules
	}).(InboundNatRuleResponseArrayOutput)
}

type Subnet struct {
	AllowPublicIp *string `pulumi:"allowPublicIp"`
	LabSubnetName *string `pulumi:"labSubnetName"`
	ResourceId    *string `pulumi:"resourceId"`
}





type SubnetInput interface {
	pulumi.Input

	ToSubnetOutput() SubnetOutput
	ToSubnetOutputWithContext(context.Context) SubnetOutput
}

type SubnetArgs struct {
	AllowPublicIp pulumi.StringPtrInput `pulumi:"allowPublicIp"`
	LabSubnetName pulumi.StringPtrInput `pulumi:"labSubnetName"`
	ResourceId    pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (SubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnet)(nil)).Elem()
}

func (i SubnetArgs) ToSubnetOutput() SubnetOutput {
	return i.ToSubnetOutputWithContext(context.Background())
}

func (i SubnetArgs) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOutput)
}





type SubnetArrayInput interface {
	pulumi.Input

	ToSubnetArrayOutput() SubnetArrayOutput
	ToSubnetArrayOutputWithContext(context.Context) SubnetArrayOutput
}

type SubnetArray []SubnetInput

func (SubnetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Subnet)(nil)).Elem()
}

func (i SubnetArray) ToSubnetArrayOutput() SubnetArrayOutput {
	return i.ToSubnetArrayOutputWithContext(context.Background())
}

func (i SubnetArray) ToSubnetArrayOutputWithContext(ctx context.Context) SubnetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetArrayOutput)
}

type SubnetOutput struct{ *pulumi.OutputState }

func (SubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnet)(nil)).Elem()
}

func (o SubnetOutput) ToSubnetOutput() SubnetOutput {
	return o
}

func (o SubnetOutput) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return o
}

func (o SubnetOutput) AllowPublicIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Subnet) *string { return v.AllowPublicIp }).(pulumi.StringPtrOutput)
}

func (o SubnetOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Subnet) *string { return v.LabSubnetName }).(pulumi.StringPtrOutput)
}

func (o SubnetOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Subnet) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type SubnetArrayOutput struct{ *pulumi.OutputState }

func (SubnetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Subnet)(nil)).Elem()
}

func (o SubnetArrayOutput) ToSubnetArrayOutput() SubnetArrayOutput {
	return o
}

func (o SubnetArrayOutput) ToSubnetArrayOutputWithContext(ctx context.Context) SubnetArrayOutput {
	return o
}

func (o SubnetArrayOutput) Index(i pulumi.IntInput) SubnetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Subnet {
		return vs[0].([]Subnet)[vs[1].(int)]
	}).(SubnetOutput)
}

type SubnetOverride struct {
	LabSubnetName                      *string                                   `pulumi:"labSubnetName"`
	ResourceId                         *string                                   `pulumi:"resourceId"`
	SharedPublicIpAddressConfiguration *SubnetSharedPublicIpAddressConfiguration `pulumi:"sharedPublicIpAddressConfiguration"`
	UseInVmCreationPermission          *string                                   `pulumi:"useInVmCreationPermission"`
	UsePublicIpAddressPermission       *string                                   `pulumi:"usePublicIpAddressPermission"`
	VirtualNetworkPoolName             *string                                   `pulumi:"virtualNetworkPoolName"`
}





type SubnetOverrideInput interface {
	pulumi.Input

	ToSubnetOverrideOutput() SubnetOverrideOutput
	ToSubnetOverrideOutputWithContext(context.Context) SubnetOverrideOutput
}

type SubnetOverrideArgs struct {
	LabSubnetName                      pulumi.StringPtrInput                            `pulumi:"labSubnetName"`
	ResourceId                         pulumi.StringPtrInput                            `pulumi:"resourceId"`
	SharedPublicIpAddressConfiguration SubnetSharedPublicIpAddressConfigurationPtrInput `pulumi:"sharedPublicIpAddressConfiguration"`
	UseInVmCreationPermission          pulumi.StringPtrInput                            `pulumi:"useInVmCreationPermission"`
	UsePublicIpAddressPermission       pulumi.StringPtrInput                            `pulumi:"usePublicIpAddressPermission"`
	VirtualNetworkPoolName             pulumi.StringPtrInput                            `pulumi:"virtualNetworkPoolName"`
}

func (SubnetOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetOverride)(nil)).Elem()
}

func (i SubnetOverrideArgs) ToSubnetOverrideOutput() SubnetOverrideOutput {
	return i.ToSubnetOverrideOutputWithContext(context.Background())
}

func (i SubnetOverrideArgs) ToSubnetOverrideOutputWithContext(ctx context.Context) SubnetOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOverrideOutput)
}





type SubnetOverrideArrayInput interface {
	pulumi.Input

	ToSubnetOverrideArrayOutput() SubnetOverrideArrayOutput
	ToSubnetOverrideArrayOutputWithContext(context.Context) SubnetOverrideArrayOutput
}

type SubnetOverrideArray []SubnetOverrideInput

func (SubnetOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetOverride)(nil)).Elem()
}

func (i SubnetOverrideArray) ToSubnetOverrideArrayOutput() SubnetOverrideArrayOutput {
	return i.ToSubnetOverrideArrayOutputWithContext(context.Background())
}

func (i SubnetOverrideArray) ToSubnetOverrideArrayOutputWithContext(ctx context.Context) SubnetOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOverrideArrayOutput)
}

type SubnetOverrideOutput struct{ *pulumi.OutputState }

func (SubnetOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetOverride)(nil)).Elem()
}

func (o SubnetOverrideOutput) ToSubnetOverrideOutput() SubnetOverrideOutput {
	return o
}

func (o SubnetOverrideOutput) ToSubnetOverrideOutputWithContext(ctx context.Context) SubnetOverrideOutput {
	return o
}

func (o SubnetOverrideOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetOverride) *string { return v.LabSubnetName }).(pulumi.StringPtrOutput)
}

func (o SubnetOverrideOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetOverride) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o SubnetOverrideOutput) SharedPublicIpAddressConfiguration() SubnetSharedPublicIpAddressConfigurationPtrOutput {
	return o.ApplyT(func(v SubnetOverride) *SubnetSharedPublicIpAddressConfiguration {
		return v.SharedPublicIpAddressConfiguration
	}).(SubnetSharedPublicIpAddressConfigurationPtrOutput)
}

func (o SubnetOverrideOutput) UseInVmCreationPermission() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetOverride) *string { return v.UseInVmCreationPermission }).(pulumi.StringPtrOutput)
}

func (o SubnetOverrideOutput) UsePublicIpAddressPermission() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetOverride) *string { return v.UsePublicIpAddressPermission }).(pulumi.StringPtrOutput)
}

func (o SubnetOverrideOutput) VirtualNetworkPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetOverride) *string { return v.VirtualNetworkPoolName }).(pulumi.StringPtrOutput)
}

type SubnetOverrideArrayOutput struct{ *pulumi.OutputState }

func (SubnetOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetOverride)(nil)).Elem()
}

func (o SubnetOverrideArrayOutput) ToSubnetOverrideArrayOutput() SubnetOverrideArrayOutput {
	return o
}

func (o SubnetOverrideArrayOutput) ToSubnetOverrideArrayOutputWithContext(ctx context.Context) SubnetOverrideArrayOutput {
	return o
}

func (o SubnetOverrideArrayOutput) Index(i pulumi.IntInput) SubnetOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetOverride {
		return vs[0].([]SubnetOverride)[vs[1].(int)]
	}).(SubnetOverrideOutput)
}

type SubnetOverrideResponse struct {
	LabSubnetName                      *string                                           `pulumi:"labSubnetName"`
	ResourceId                         *string                                           `pulumi:"resourceId"`
	SharedPublicIpAddressConfiguration *SubnetSharedPublicIpAddressConfigurationResponse `pulumi:"sharedPublicIpAddressConfiguration"`
	UseInVmCreationPermission          *string                                           `pulumi:"useInVmCreationPermission"`
	UsePublicIpAddressPermission       *string                                           `pulumi:"usePublicIpAddressPermission"`
	VirtualNetworkPoolName             *string                                           `pulumi:"virtualNetworkPoolName"`
}

type SubnetOverrideResponseOutput struct{ *pulumi.OutputState }

func (SubnetOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetOverrideResponse)(nil)).Elem()
}

func (o SubnetOverrideResponseOutput) ToSubnetOverrideResponseOutput() SubnetOverrideResponseOutput {
	return o
}

func (o SubnetOverrideResponseOutput) ToSubnetOverrideResponseOutputWithContext(ctx context.Context) SubnetOverrideResponseOutput {
	return o
}

func (o SubnetOverrideResponseOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetOverrideResponse) *string { return v.LabSubnetName }).(pulumi.StringPtrOutput)
}

func (o SubnetOverrideResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetOverrideResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o SubnetOverrideResponseOutput) SharedPublicIpAddressConfiguration() SubnetSharedPublicIpAddressConfigurationResponsePtrOutput {
	return o.ApplyT(func(v SubnetOverrideResponse) *SubnetSharedPublicIpAddressConfigurationResponse {
		return v.SharedPublicIpAddressConfiguration
	}).(SubnetSharedPublicIpAddressConfigurationResponsePtrOutput)
}

func (o SubnetOverrideResponseOutput) UseInVmCreationPermission() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetOverrideResponse) *string { return v.UseInVmCreationPermission }).(pulumi.StringPtrOutput)
}

func (o SubnetOverrideResponseOutput) UsePublicIpAddressPermission() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetOverrideResponse) *string { return v.UsePublicIpAddressPermission }).(pulumi.StringPtrOutput)
}

func (o SubnetOverrideResponseOutput) VirtualNetworkPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetOverrideResponse) *string { return v.VirtualNetworkPoolName }).(pulumi.StringPtrOutput)
}

type SubnetOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (SubnetOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetOverrideResponse)(nil)).Elem()
}

func (o SubnetOverrideResponseArrayOutput) ToSubnetOverrideResponseArrayOutput() SubnetOverrideResponseArrayOutput {
	return o
}

func (o SubnetOverrideResponseArrayOutput) ToSubnetOverrideResponseArrayOutputWithContext(ctx context.Context) SubnetOverrideResponseArrayOutput {
	return o
}

func (o SubnetOverrideResponseArrayOutput) Index(i pulumi.IntInput) SubnetOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetOverrideResponse {
		return vs[0].([]SubnetOverrideResponse)[vs[1].(int)]
	}).(SubnetOverrideResponseOutput)
}

type SubnetResponse struct {
	AllowPublicIp *string `pulumi:"allowPublicIp"`
	LabSubnetName *string `pulumi:"labSubnetName"`
	ResourceId    *string `pulumi:"resourceId"`
}

type SubnetResponseOutput struct{ *pulumi.OutputState }

func (SubnetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResponse)(nil)).Elem()
}

func (o SubnetResponseOutput) ToSubnetResponseOutput() SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) ToSubnetResponseOutputWithContext(ctx context.Context) SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) AllowPublicIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.AllowPublicIp }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.LabSubnetName }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type SubnetResponseArrayOutput struct{ *pulumi.OutputState }

func (SubnetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetResponse)(nil)).Elem()
}

func (o SubnetResponseArrayOutput) ToSubnetResponseArrayOutput() SubnetResponseArrayOutput {
	return o
}

func (o SubnetResponseArrayOutput) ToSubnetResponseArrayOutputWithContext(ctx context.Context) SubnetResponseArrayOutput {
	return o
}

func (o SubnetResponseArrayOutput) Index(i pulumi.IntInput) SubnetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetResponse {
		return vs[0].([]SubnetResponse)[vs[1].(int)]
	}).(SubnetResponseOutput)
}

type SubnetSharedPublicIpAddressConfiguration struct {
	AllowedPorts []Port `pulumi:"allowedPorts"`
}





type SubnetSharedPublicIpAddressConfigurationInput interface {
	pulumi.Input

	ToSubnetSharedPublicIpAddressConfigurationOutput() SubnetSharedPublicIpAddressConfigurationOutput
	ToSubnetSharedPublicIpAddressConfigurationOutputWithContext(context.Context) SubnetSharedPublicIpAddressConfigurationOutput
}

type SubnetSharedPublicIpAddressConfigurationArgs struct {
	AllowedPorts PortArrayInput `pulumi:"allowedPorts"`
}

func (SubnetSharedPublicIpAddressConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetSharedPublicIpAddressConfiguration)(nil)).Elem()
}

func (i SubnetSharedPublicIpAddressConfigurationArgs) ToSubnetSharedPublicIpAddressConfigurationOutput() SubnetSharedPublicIpAddressConfigurationOutput {
	return i.ToSubnetSharedPublicIpAddressConfigurationOutputWithContext(context.Background())
}

func (i SubnetSharedPublicIpAddressConfigurationArgs) ToSubnetSharedPublicIpAddressConfigurationOutputWithContext(ctx context.Context) SubnetSharedPublicIpAddressConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetSharedPublicIpAddressConfigurationOutput)
}

func (i SubnetSharedPublicIpAddressConfigurationArgs) ToSubnetSharedPublicIpAddressConfigurationPtrOutput() SubnetSharedPublicIpAddressConfigurationPtrOutput {
	return i.ToSubnetSharedPublicIpAddressConfigurationPtrOutputWithContext(context.Background())
}

func (i SubnetSharedPublicIpAddressConfigurationArgs) ToSubnetSharedPublicIpAddressConfigurationPtrOutputWithContext(ctx context.Context) SubnetSharedPublicIpAddressConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetSharedPublicIpAddressConfigurationOutput).ToSubnetSharedPublicIpAddressConfigurationPtrOutputWithContext(ctx)
}









type SubnetSharedPublicIpAddressConfigurationPtrInput interface {
	pulumi.Input

	ToSubnetSharedPublicIpAddressConfigurationPtrOutput() SubnetSharedPublicIpAddressConfigurationPtrOutput
	ToSubnetSharedPublicIpAddressConfigurationPtrOutputWithContext(context.Context) SubnetSharedPublicIpAddressConfigurationPtrOutput
}

type subnetSharedPublicIpAddressConfigurationPtrType SubnetSharedPublicIpAddressConfigurationArgs

func SubnetSharedPublicIpAddressConfigurationPtr(v *SubnetSharedPublicIpAddressConfigurationArgs) SubnetSharedPublicIpAddressConfigurationPtrInput {
	return (*subnetSharedPublicIpAddressConfigurationPtrType)(v)
}

func (*subnetSharedPublicIpAddressConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetSharedPublicIpAddressConfiguration)(nil)).Elem()
}

func (i *subnetSharedPublicIpAddressConfigurationPtrType) ToSubnetSharedPublicIpAddressConfigurationPtrOutput() SubnetSharedPublicIpAddressConfigurationPtrOutput {
	return i.ToSubnetSharedPublicIpAddressConfigurationPtrOutputWithContext(context.Background())
}

func (i *subnetSharedPublicIpAddressConfigurationPtrType) ToSubnetSharedPublicIpAddressConfigurationPtrOutputWithContext(ctx context.Context) SubnetSharedPublicIpAddressConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetSharedPublicIpAddressConfigurationPtrOutput)
}

type SubnetSharedPublicIpAddressConfigurationOutput struct{ *pulumi.OutputState }

func (SubnetSharedPublicIpAddressConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetSharedPublicIpAddressConfiguration)(nil)).Elem()
}

func (o SubnetSharedPublicIpAddressConfigurationOutput) ToSubnetSharedPublicIpAddressConfigurationOutput() SubnetSharedPublicIpAddressConfigurationOutput {
	return o
}

func (o SubnetSharedPublicIpAddressConfigurationOutput) ToSubnetSharedPublicIpAddressConfigurationOutputWithContext(ctx context.Context) SubnetSharedPublicIpAddressConfigurationOutput {
	return o
}

func (o SubnetSharedPublicIpAddressConfigurationOutput) ToSubnetSharedPublicIpAddressConfigurationPtrOutput() SubnetSharedPublicIpAddressConfigurationPtrOutput {
	return o.ToSubnetSharedPublicIpAddressConfigurationPtrOutputWithContext(context.Background())
}

func (o SubnetSharedPublicIpAddressConfigurationOutput) ToSubnetSharedPublicIpAddressConfigurationPtrOutputWithContext(ctx context.Context) SubnetSharedPublicIpAddressConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubnetSharedPublicIpAddressConfiguration) *SubnetSharedPublicIpAddressConfiguration {
		return &v
	}).(SubnetSharedPublicIpAddressConfigurationPtrOutput)
}

func (o SubnetSharedPublicIpAddressConfigurationOutput) AllowedPorts() PortArrayOutput {
	return o.ApplyT(func(v SubnetSharedPublicIpAddressConfiguration) []Port { return v.AllowedPorts }).(PortArrayOutput)
}

type SubnetSharedPublicIpAddressConfigurationPtrOutput struct{ *pulumi.OutputState }

func (SubnetSharedPublicIpAddressConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetSharedPublicIpAddressConfiguration)(nil)).Elem()
}

func (o SubnetSharedPublicIpAddressConfigurationPtrOutput) ToSubnetSharedPublicIpAddressConfigurationPtrOutput() SubnetSharedPublicIpAddressConfigurationPtrOutput {
	return o
}

func (o SubnetSharedPublicIpAddressConfigurationPtrOutput) ToSubnetSharedPublicIpAddressConfigurationPtrOutputWithContext(ctx context.Context) SubnetSharedPublicIpAddressConfigurationPtrOutput {
	return o
}

func (o SubnetSharedPublicIpAddressConfigurationPtrOutput) Elem() SubnetSharedPublicIpAddressConfigurationOutput {
	return o.ApplyT(func(v *SubnetSharedPublicIpAddressConfiguration) SubnetSharedPublicIpAddressConfiguration {
		if v != nil {
			return *v
		}
		var ret SubnetSharedPublicIpAddressConfiguration
		return ret
	}).(SubnetSharedPublicIpAddressConfigurationOutput)
}

func (o SubnetSharedPublicIpAddressConfigurationPtrOutput) AllowedPorts() PortArrayOutput {
	return o.ApplyT(func(v *SubnetSharedPublicIpAddressConfiguration) []Port {
		if v == nil {
			return nil
		}
		return v.AllowedPorts
	}).(PortArrayOutput)
}

type SubnetSharedPublicIpAddressConfigurationResponse struct {
	AllowedPorts []PortResponse `pulumi:"allowedPorts"`
}

type SubnetSharedPublicIpAddressConfigurationResponseOutput struct{ *pulumi.OutputState }

func (SubnetSharedPublicIpAddressConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetSharedPublicIpAddressConfigurationResponse)(nil)).Elem()
}

func (o SubnetSharedPublicIpAddressConfigurationResponseOutput) ToSubnetSharedPublicIpAddressConfigurationResponseOutput() SubnetSharedPublicIpAddressConfigurationResponseOutput {
	return o
}

func (o SubnetSharedPublicIpAddressConfigurationResponseOutput) ToSubnetSharedPublicIpAddressConfigurationResponseOutputWithContext(ctx context.Context) SubnetSharedPublicIpAddressConfigurationResponseOutput {
	return o
}

func (o SubnetSharedPublicIpAddressConfigurationResponseOutput) AllowedPorts() PortResponseArrayOutput {
	return o.ApplyT(func(v SubnetSharedPublicIpAddressConfigurationResponse) []PortResponse { return v.AllowedPorts }).(PortResponseArrayOutput)
}

type SubnetSharedPublicIpAddressConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (SubnetSharedPublicIpAddressConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetSharedPublicIpAddressConfigurationResponse)(nil)).Elem()
}

func (o SubnetSharedPublicIpAddressConfigurationResponsePtrOutput) ToSubnetSharedPublicIpAddressConfigurationResponsePtrOutput() SubnetSharedPublicIpAddressConfigurationResponsePtrOutput {
	return o
}

func (o SubnetSharedPublicIpAddressConfigurationResponsePtrOutput) ToSubnetSharedPublicIpAddressConfigurationResponsePtrOutputWithContext(ctx context.Context) SubnetSharedPublicIpAddressConfigurationResponsePtrOutput {
	return o
}

func (o SubnetSharedPublicIpAddressConfigurationResponsePtrOutput) Elem() SubnetSharedPublicIpAddressConfigurationResponseOutput {
	return o.ApplyT(func(v *SubnetSharedPublicIpAddressConfigurationResponse) SubnetSharedPublicIpAddressConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret SubnetSharedPublicIpAddressConfigurationResponse
		return ret
	}).(SubnetSharedPublicIpAddressConfigurationResponseOutput)
}

func (o SubnetSharedPublicIpAddressConfigurationResponsePtrOutput) AllowedPorts() PortResponseArrayOutput {
	return o.ApplyT(func(v *SubnetSharedPublicIpAddressConfigurationResponse) []PortResponse {
		if v == nil {
			return nil
		}
		return v.AllowedPorts
	}).(PortResponseArrayOutput)
}

type UserIdentity struct {
	AppId         *string `pulumi:"appId"`
	ObjectId      *string `pulumi:"objectId"`
	PrincipalId   *string `pulumi:"principalId"`
	PrincipalName *string `pulumi:"principalName"`
	TenantId      *string `pulumi:"tenantId"`
}





type UserIdentityInput interface {
	pulumi.Input

	ToUserIdentityOutput() UserIdentityOutput
	ToUserIdentityOutputWithContext(context.Context) UserIdentityOutput
}

type UserIdentityArgs struct {
	AppId         pulumi.StringPtrInput `pulumi:"appId"`
	ObjectId      pulumi.StringPtrInput `pulumi:"objectId"`
	PrincipalId   pulumi.StringPtrInput `pulumi:"principalId"`
	PrincipalName pulumi.StringPtrInput `pulumi:"principalName"`
	TenantId      pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (UserIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentity)(nil)).Elem()
}

func (i UserIdentityArgs) ToUserIdentityOutput() UserIdentityOutput {
	return i.ToUserIdentityOutputWithContext(context.Background())
}

func (i UserIdentityArgs) ToUserIdentityOutputWithContext(ctx context.Context) UserIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityOutput)
}

func (i UserIdentityArgs) ToUserIdentityPtrOutput() UserIdentityPtrOutput {
	return i.ToUserIdentityPtrOutputWithContext(context.Background())
}

func (i UserIdentityArgs) ToUserIdentityPtrOutputWithContext(ctx context.Context) UserIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityOutput).ToUserIdentityPtrOutputWithContext(ctx)
}









type UserIdentityPtrInput interface {
	pulumi.Input

	ToUserIdentityPtrOutput() UserIdentityPtrOutput
	ToUserIdentityPtrOutputWithContext(context.Context) UserIdentityPtrOutput
}

type userIdentityPtrType UserIdentityArgs

func UserIdentityPtr(v *UserIdentityArgs) UserIdentityPtrInput {
	return (*userIdentityPtrType)(v)
}

func (*userIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserIdentity)(nil)).Elem()
}

func (i *userIdentityPtrType) ToUserIdentityPtrOutput() UserIdentityPtrOutput {
	return i.ToUserIdentityPtrOutputWithContext(context.Background())
}

func (i *userIdentityPtrType) ToUserIdentityPtrOutputWithContext(ctx context.Context) UserIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityPtrOutput)
}

type UserIdentityOutput struct{ *pulumi.OutputState }

func (UserIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentity)(nil)).Elem()
}

func (o UserIdentityOutput) ToUserIdentityOutput() UserIdentityOutput {
	return o
}

func (o UserIdentityOutput) ToUserIdentityOutputWithContext(ctx context.Context) UserIdentityOutput {
	return o
}

func (o UserIdentityOutput) ToUserIdentityPtrOutput() UserIdentityPtrOutput {
	return o.ToUserIdentityPtrOutputWithContext(context.Background())
}

func (o UserIdentityOutput) ToUserIdentityPtrOutputWithContext(ctx context.Context) UserIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserIdentity) *UserIdentity {
		return &v
	}).(UserIdentityPtrOutput)
}

func (o UserIdentityOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentity) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o UserIdentityOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentity) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o UserIdentityOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentity) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o UserIdentityOutput) PrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentity) *string { return v.PrincipalName }).(pulumi.StringPtrOutput)
}

func (o UserIdentityOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentity) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type UserIdentityPtrOutput struct{ *pulumi.OutputState }

func (UserIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserIdentity)(nil)).Elem()
}

func (o UserIdentityPtrOutput) ToUserIdentityPtrOutput() UserIdentityPtrOutput {
	return o
}

func (o UserIdentityPtrOutput) ToUserIdentityPtrOutputWithContext(ctx context.Context) UserIdentityPtrOutput {
	return o
}

func (o UserIdentityPtrOutput) Elem() UserIdentityOutput {
	return o.ApplyT(func(v *UserIdentity) UserIdentity {
		if v != nil {
			return *v
		}
		var ret UserIdentity
		return ret
	}).(UserIdentityOutput)
}

func (o UserIdentityPtrOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserIdentity) *string {
		if v == nil {
			return nil
		}
		return v.AppId
	}).(pulumi.StringPtrOutput)
}

func (o UserIdentityPtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserIdentity) *string {
		if v == nil {
			return nil
		}
		return v.ObjectId
	}).(pulumi.StringPtrOutput)
}

func (o UserIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserIdentity) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o UserIdentityPtrOutput) PrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserIdentity) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalName
	}).(pulumi.StringPtrOutput)
}

func (o UserIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserIdentity) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type UserIdentityResponse struct {
	AppId         *string `pulumi:"appId"`
	ObjectId      *string `pulumi:"objectId"`
	PrincipalId   *string `pulumi:"principalId"`
	PrincipalName *string `pulumi:"principalName"`
	TenantId      *string `pulumi:"tenantId"`
}

type UserIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityResponse)(nil)).Elem()
}

func (o UserIdentityResponseOutput) ToUserIdentityResponseOutput() UserIdentityResponseOutput {
	return o
}

func (o UserIdentityResponseOutput) ToUserIdentityResponseOutputWithContext(ctx context.Context) UserIdentityResponseOutput {
	return o
}

func (o UserIdentityResponseOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityResponse) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o UserIdentityResponseOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityResponse) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o UserIdentityResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o UserIdentityResponseOutput) PrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityResponse) *string { return v.PrincipalName }).(pulumi.StringPtrOutput)
}

func (o UserIdentityResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type UserIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (UserIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserIdentityResponse)(nil)).Elem()
}

func (o UserIdentityResponsePtrOutput) ToUserIdentityResponsePtrOutput() UserIdentityResponsePtrOutput {
	return o
}

func (o UserIdentityResponsePtrOutput) ToUserIdentityResponsePtrOutputWithContext(ctx context.Context) UserIdentityResponsePtrOutput {
	return o
}

func (o UserIdentityResponsePtrOutput) Elem() UserIdentityResponseOutput {
	return o.ApplyT(func(v *UserIdentityResponse) UserIdentityResponse {
		if v != nil {
			return *v
		}
		var ret UserIdentityResponse
		return ret
	}).(UserIdentityResponseOutput)
}

func (o UserIdentityResponsePtrOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppId
	}).(pulumi.StringPtrOutput)
}

func (o UserIdentityResponsePtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.ObjectId
	}).(pulumi.StringPtrOutput)
}

func (o UserIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o UserIdentityResponsePtrOutput) PrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalName
	}).(pulumi.StringPtrOutput)
}

func (o UserIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type UserSecretStore struct {
	KeyVaultId  *string `pulumi:"keyVaultId"`
	KeyVaultUri *string `pulumi:"keyVaultUri"`
}





type UserSecretStoreInput interface {
	pulumi.Input

	ToUserSecretStoreOutput() UserSecretStoreOutput
	ToUserSecretStoreOutputWithContext(context.Context) UserSecretStoreOutput
}

type UserSecretStoreArgs struct {
	KeyVaultId  pulumi.StringPtrInput `pulumi:"keyVaultId"`
	KeyVaultUri pulumi.StringPtrInput `pulumi:"keyVaultUri"`
}

func (UserSecretStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserSecretStore)(nil)).Elem()
}

func (i UserSecretStoreArgs) ToUserSecretStoreOutput() UserSecretStoreOutput {
	return i.ToUserSecretStoreOutputWithContext(context.Background())
}

func (i UserSecretStoreArgs) ToUserSecretStoreOutputWithContext(ctx context.Context) UserSecretStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSecretStoreOutput)
}

func (i UserSecretStoreArgs) ToUserSecretStorePtrOutput() UserSecretStorePtrOutput {
	return i.ToUserSecretStorePtrOutputWithContext(context.Background())
}

func (i UserSecretStoreArgs) ToUserSecretStorePtrOutputWithContext(ctx context.Context) UserSecretStorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSecretStoreOutput).ToUserSecretStorePtrOutputWithContext(ctx)
}









type UserSecretStorePtrInput interface {
	pulumi.Input

	ToUserSecretStorePtrOutput() UserSecretStorePtrOutput
	ToUserSecretStorePtrOutputWithContext(context.Context) UserSecretStorePtrOutput
}

type userSecretStorePtrType UserSecretStoreArgs

func UserSecretStorePtr(v *UserSecretStoreArgs) UserSecretStorePtrInput {
	return (*userSecretStorePtrType)(v)
}

func (*userSecretStorePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSecretStore)(nil)).Elem()
}

func (i *userSecretStorePtrType) ToUserSecretStorePtrOutput() UserSecretStorePtrOutput {
	return i.ToUserSecretStorePtrOutputWithContext(context.Background())
}

func (i *userSecretStorePtrType) ToUserSecretStorePtrOutputWithContext(ctx context.Context) UserSecretStorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSecretStorePtrOutput)
}

type UserSecretStoreOutput struct{ *pulumi.OutputState }

func (UserSecretStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserSecretStore)(nil)).Elem()
}

func (o UserSecretStoreOutput) ToUserSecretStoreOutput() UserSecretStoreOutput {
	return o
}

func (o UserSecretStoreOutput) ToUserSecretStoreOutputWithContext(ctx context.Context) UserSecretStoreOutput {
	return o
}

func (o UserSecretStoreOutput) ToUserSecretStorePtrOutput() UserSecretStorePtrOutput {
	return o.ToUserSecretStorePtrOutputWithContext(context.Background())
}

func (o UserSecretStoreOutput) ToUserSecretStorePtrOutputWithContext(ctx context.Context) UserSecretStorePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserSecretStore) *UserSecretStore {
		return &v
	}).(UserSecretStorePtrOutput)
}

func (o UserSecretStoreOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSecretStore) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

func (o UserSecretStoreOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSecretStore) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

type UserSecretStorePtrOutput struct{ *pulumi.OutputState }

func (UserSecretStorePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSecretStore)(nil)).Elem()
}

func (o UserSecretStorePtrOutput) ToUserSecretStorePtrOutput() UserSecretStorePtrOutput {
	return o
}

func (o UserSecretStorePtrOutput) ToUserSecretStorePtrOutputWithContext(ctx context.Context) UserSecretStorePtrOutput {
	return o
}

func (o UserSecretStorePtrOutput) Elem() UserSecretStoreOutput {
	return o.ApplyT(func(v *UserSecretStore) UserSecretStore {
		if v != nil {
			return *v
		}
		var ret UserSecretStore
		return ret
	}).(UserSecretStoreOutput)
}

func (o UserSecretStorePtrOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSecretStore) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultId
	}).(pulumi.StringPtrOutput)
}

func (o UserSecretStorePtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSecretStore) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

type UserSecretStoreResponse struct {
	KeyVaultId  *string `pulumi:"keyVaultId"`
	KeyVaultUri *string `pulumi:"keyVaultUri"`
}

type UserSecretStoreResponseOutput struct{ *pulumi.OutputState }

func (UserSecretStoreResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserSecretStoreResponse)(nil)).Elem()
}

func (o UserSecretStoreResponseOutput) ToUserSecretStoreResponseOutput() UserSecretStoreResponseOutput {
	return o
}

func (o UserSecretStoreResponseOutput) ToUserSecretStoreResponseOutputWithContext(ctx context.Context) UserSecretStoreResponseOutput {
	return o
}

func (o UserSecretStoreResponseOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSecretStoreResponse) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

func (o UserSecretStoreResponseOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserSecretStoreResponse) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

type UserSecretStoreResponsePtrOutput struct{ *pulumi.OutputState }

func (UserSecretStoreResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSecretStoreResponse)(nil)).Elem()
}

func (o UserSecretStoreResponsePtrOutput) ToUserSecretStoreResponsePtrOutput() UserSecretStoreResponsePtrOutput {
	return o
}

func (o UserSecretStoreResponsePtrOutput) ToUserSecretStoreResponsePtrOutputWithContext(ctx context.Context) UserSecretStoreResponsePtrOutput {
	return o
}

func (o UserSecretStoreResponsePtrOutput) Elem() UserSecretStoreResponseOutput {
	return o.ApplyT(func(v *UserSecretStoreResponse) UserSecretStoreResponse {
		if v != nil {
			return *v
		}
		var ret UserSecretStoreResponse
		return ret
	}).(UserSecretStoreResponseOutput)
}

func (o UserSecretStoreResponsePtrOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSecretStoreResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultId
	}).(pulumi.StringPtrOutput)
}

func (o UserSecretStoreResponsePtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserSecretStoreResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

type WeekDetails struct {
	Time     *string  `pulumi:"time"`
	Weekdays []string `pulumi:"weekdays"`
}





type WeekDetailsInput interface {
	pulumi.Input

	ToWeekDetailsOutput() WeekDetailsOutput
	ToWeekDetailsOutputWithContext(context.Context) WeekDetailsOutput
}

type WeekDetailsArgs struct {
	Time     pulumi.StringPtrInput   `pulumi:"time"`
	Weekdays pulumi.StringArrayInput `pulumi:"weekdays"`
}

func (WeekDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WeekDetails)(nil)).Elem()
}

func (i WeekDetailsArgs) ToWeekDetailsOutput() WeekDetailsOutput {
	return i.ToWeekDetailsOutputWithContext(context.Background())
}

func (i WeekDetailsArgs) ToWeekDetailsOutputWithContext(ctx context.Context) WeekDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeekDetailsOutput)
}

func (i WeekDetailsArgs) ToWeekDetailsPtrOutput() WeekDetailsPtrOutput {
	return i.ToWeekDetailsPtrOutputWithContext(context.Background())
}

func (i WeekDetailsArgs) ToWeekDetailsPtrOutputWithContext(ctx context.Context) WeekDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeekDetailsOutput).ToWeekDetailsPtrOutputWithContext(ctx)
}









type WeekDetailsPtrInput interface {
	pulumi.Input

	ToWeekDetailsPtrOutput() WeekDetailsPtrOutput
	ToWeekDetailsPtrOutputWithContext(context.Context) WeekDetailsPtrOutput
}

type weekDetailsPtrType WeekDetailsArgs

func WeekDetailsPtr(v *WeekDetailsArgs) WeekDetailsPtrInput {
	return (*weekDetailsPtrType)(v)
}

func (*weekDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WeekDetails)(nil)).Elem()
}

func (i *weekDetailsPtrType) ToWeekDetailsPtrOutput() WeekDetailsPtrOutput {
	return i.ToWeekDetailsPtrOutputWithContext(context.Background())
}

func (i *weekDetailsPtrType) ToWeekDetailsPtrOutputWithContext(ctx context.Context) WeekDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeekDetailsPtrOutput)
}

type WeekDetailsOutput struct{ *pulumi.OutputState }

func (WeekDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeekDetails)(nil)).Elem()
}

func (o WeekDetailsOutput) ToWeekDetailsOutput() WeekDetailsOutput {
	return o
}

func (o WeekDetailsOutput) ToWeekDetailsOutputWithContext(ctx context.Context) WeekDetailsOutput {
	return o
}

func (o WeekDetailsOutput) ToWeekDetailsPtrOutput() WeekDetailsPtrOutput {
	return o.ToWeekDetailsPtrOutputWithContext(context.Background())
}

func (o WeekDetailsOutput) ToWeekDetailsPtrOutputWithContext(ctx context.Context) WeekDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WeekDetails) *WeekDetails {
		return &v
	}).(WeekDetailsPtrOutput)
}

func (o WeekDetailsOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WeekDetails) *string { return v.Time }).(pulumi.StringPtrOutput)
}

func (o WeekDetailsOutput) Weekdays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WeekDetails) []string { return v.Weekdays }).(pulumi.StringArrayOutput)
}

type WeekDetailsPtrOutput struct{ *pulumi.OutputState }

func (WeekDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WeekDetails)(nil)).Elem()
}

func (o WeekDetailsPtrOutput) ToWeekDetailsPtrOutput() WeekDetailsPtrOutput {
	return o
}

func (o WeekDetailsPtrOutput) ToWeekDetailsPtrOutputWithContext(ctx context.Context) WeekDetailsPtrOutput {
	return o
}

func (o WeekDetailsPtrOutput) Elem() WeekDetailsOutput {
	return o.ApplyT(func(v *WeekDetails) WeekDetails {
		if v != nil {
			return *v
		}
		var ret WeekDetails
		return ret
	}).(WeekDetailsOutput)
}

func (o WeekDetailsPtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WeekDetails) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

func (o WeekDetailsPtrOutput) Weekdays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WeekDetails) []string {
		if v == nil {
			return nil
		}
		return v.Weekdays
	}).(pulumi.StringArrayOutput)
}

type WeekDetailsResponse struct {
	Time     *string  `pulumi:"time"`
	Weekdays []string `pulumi:"weekdays"`
}

type WeekDetailsResponseOutput struct{ *pulumi.OutputState }

func (WeekDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeekDetailsResponse)(nil)).Elem()
}

func (o WeekDetailsResponseOutput) ToWeekDetailsResponseOutput() WeekDetailsResponseOutput {
	return o
}

func (o WeekDetailsResponseOutput) ToWeekDetailsResponseOutputWithContext(ctx context.Context) WeekDetailsResponseOutput {
	return o
}

func (o WeekDetailsResponseOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WeekDetailsResponse) *string { return v.Time }).(pulumi.StringPtrOutput)
}

func (o WeekDetailsResponseOutput) Weekdays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WeekDetailsResponse) []string { return v.Weekdays }).(pulumi.StringArrayOutput)
}

type WeekDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (WeekDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WeekDetailsResponse)(nil)).Elem()
}

func (o WeekDetailsResponsePtrOutput) ToWeekDetailsResponsePtrOutput() WeekDetailsResponsePtrOutput {
	return o
}

func (o WeekDetailsResponsePtrOutput) ToWeekDetailsResponsePtrOutputWithContext(ctx context.Context) WeekDetailsResponsePtrOutput {
	return o
}

func (o WeekDetailsResponsePtrOutput) Elem() WeekDetailsResponseOutput {
	return o.ApplyT(func(v *WeekDetailsResponse) WeekDetailsResponse {
		if v != nil {
			return *v
		}
		var ret WeekDetailsResponse
		return ret
	}).(WeekDetailsResponseOutput)
}

func (o WeekDetailsResponsePtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WeekDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

func (o WeekDetailsResponsePtrOutput) Weekdays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WeekDetailsResponse) []string {
		if v == nil {
			return nil
		}
		return v.Weekdays
	}).(pulumi.StringArrayOutput)
}

type WindowsOsInfo struct {
	WindowsOsState *string `pulumi:"windowsOsState"`
}





type WindowsOsInfoInput interface {
	pulumi.Input

	ToWindowsOsInfoOutput() WindowsOsInfoOutput
	ToWindowsOsInfoOutputWithContext(context.Context) WindowsOsInfoOutput
}

type WindowsOsInfoArgs struct {
	WindowsOsState pulumi.StringPtrInput `pulumi:"windowsOsState"`
}

func (WindowsOsInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsOsInfo)(nil)).Elem()
}

func (i WindowsOsInfoArgs) ToWindowsOsInfoOutput() WindowsOsInfoOutput {
	return i.ToWindowsOsInfoOutputWithContext(context.Background())
}

func (i WindowsOsInfoArgs) ToWindowsOsInfoOutputWithContext(ctx context.Context) WindowsOsInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsOsInfoOutput)
}

func (i WindowsOsInfoArgs) ToWindowsOsInfoPtrOutput() WindowsOsInfoPtrOutput {
	return i.ToWindowsOsInfoPtrOutputWithContext(context.Background())
}

func (i WindowsOsInfoArgs) ToWindowsOsInfoPtrOutputWithContext(ctx context.Context) WindowsOsInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsOsInfoOutput).ToWindowsOsInfoPtrOutputWithContext(ctx)
}









type WindowsOsInfoPtrInput interface {
	pulumi.Input

	ToWindowsOsInfoPtrOutput() WindowsOsInfoPtrOutput
	ToWindowsOsInfoPtrOutputWithContext(context.Context) WindowsOsInfoPtrOutput
}

type windowsOsInfoPtrType WindowsOsInfoArgs

func WindowsOsInfoPtr(v *WindowsOsInfoArgs) WindowsOsInfoPtrInput {
	return (*windowsOsInfoPtrType)(v)
}

func (*windowsOsInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsOsInfo)(nil)).Elem()
}

func (i *windowsOsInfoPtrType) ToWindowsOsInfoPtrOutput() WindowsOsInfoPtrOutput {
	return i.ToWindowsOsInfoPtrOutputWithContext(context.Background())
}

func (i *windowsOsInfoPtrType) ToWindowsOsInfoPtrOutputWithContext(ctx context.Context) WindowsOsInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsOsInfoPtrOutput)
}

type WindowsOsInfoOutput struct{ *pulumi.OutputState }

func (WindowsOsInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsOsInfo)(nil)).Elem()
}

func (o WindowsOsInfoOutput) ToWindowsOsInfoOutput() WindowsOsInfoOutput {
	return o
}

func (o WindowsOsInfoOutput) ToWindowsOsInfoOutputWithContext(ctx context.Context) WindowsOsInfoOutput {
	return o
}

func (o WindowsOsInfoOutput) ToWindowsOsInfoPtrOutput() WindowsOsInfoPtrOutput {
	return o.ToWindowsOsInfoPtrOutputWithContext(context.Background())
}

func (o WindowsOsInfoOutput) ToWindowsOsInfoPtrOutputWithContext(ctx context.Context) WindowsOsInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WindowsOsInfo) *WindowsOsInfo {
		return &v
	}).(WindowsOsInfoPtrOutput)
}

func (o WindowsOsInfoOutput) WindowsOsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsOsInfo) *string { return v.WindowsOsState }).(pulumi.StringPtrOutput)
}

type WindowsOsInfoPtrOutput struct{ *pulumi.OutputState }

func (WindowsOsInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsOsInfo)(nil)).Elem()
}

func (o WindowsOsInfoPtrOutput) ToWindowsOsInfoPtrOutput() WindowsOsInfoPtrOutput {
	return o
}

func (o WindowsOsInfoPtrOutput) ToWindowsOsInfoPtrOutputWithContext(ctx context.Context) WindowsOsInfoPtrOutput {
	return o
}

func (o WindowsOsInfoPtrOutput) Elem() WindowsOsInfoOutput {
	return o.ApplyT(func(v *WindowsOsInfo) WindowsOsInfo {
		if v != nil {
			return *v
		}
		var ret WindowsOsInfo
		return ret
	}).(WindowsOsInfoOutput)
}

func (o WindowsOsInfoPtrOutput) WindowsOsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsOsInfo) *string {
		if v == nil {
			return nil
		}
		return v.WindowsOsState
	}).(pulumi.StringPtrOutput)
}

type WindowsOsInfoResponse struct {
	WindowsOsState *string `pulumi:"windowsOsState"`
}

type WindowsOsInfoResponseOutput struct{ *pulumi.OutputState }

func (WindowsOsInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsOsInfoResponse)(nil)).Elem()
}

func (o WindowsOsInfoResponseOutput) ToWindowsOsInfoResponseOutput() WindowsOsInfoResponseOutput {
	return o
}

func (o WindowsOsInfoResponseOutput) ToWindowsOsInfoResponseOutputWithContext(ctx context.Context) WindowsOsInfoResponseOutput {
	return o
}

func (o WindowsOsInfoResponseOutput) WindowsOsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsOsInfoResponse) *string { return v.WindowsOsState }).(pulumi.StringPtrOutput)
}

type WindowsOsInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (WindowsOsInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsOsInfoResponse)(nil)).Elem()
}

func (o WindowsOsInfoResponsePtrOutput) ToWindowsOsInfoResponsePtrOutput() WindowsOsInfoResponsePtrOutput {
	return o
}

func (o WindowsOsInfoResponsePtrOutput) ToWindowsOsInfoResponsePtrOutputWithContext(ctx context.Context) WindowsOsInfoResponsePtrOutput {
	return o
}

func (o WindowsOsInfoResponsePtrOutput) Elem() WindowsOsInfoResponseOutput {
	return o.ApplyT(func(v *WindowsOsInfoResponse) WindowsOsInfoResponse {
		if v != nil {
			return *v
		}
		var ret WindowsOsInfoResponse
		return ret
	}).(WindowsOsInfoResponseOutput)
}

func (o WindowsOsInfoResponsePtrOutput) WindowsOsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsOsInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.WindowsOsState
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicableScheduleResponseOutput{})
	pulumi.RegisterOutputType(ArmTemplateParameterPropertiesOutput{})
	pulumi.RegisterOutputType(ArmTemplateParameterPropertiesArrayOutput{})
	pulumi.RegisterOutputType(ArmTemplateParameterPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ArmTemplateParameterPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(ArtifactDeploymentStatusPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ArtifactInstallPropertiesOutput{})
	pulumi.RegisterOutputType(ArtifactInstallPropertiesArrayOutput{})
	pulumi.RegisterOutputType(ArtifactInstallPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ArtifactInstallPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(ArtifactParameterPropertiesOutput{})
	pulumi.RegisterOutputType(ArtifactParameterPropertiesArrayOutput{})
	pulumi.RegisterOutputType(ArtifactParameterPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ArtifactParameterPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(AttachNewDataDiskOptionsOutput{})
	pulumi.RegisterOutputType(AttachNewDataDiskOptionsPtrOutput{})
	pulumi.RegisterOutputType(AttachNewDataDiskOptionsResponseOutput{})
	pulumi.RegisterOutputType(AttachNewDataDiskOptionsResponsePtrOutput{})
	pulumi.RegisterOutputType(BulkCreationParametersOutput{})
	pulumi.RegisterOutputType(BulkCreationParametersPtrOutput{})
	pulumi.RegisterOutputType(BulkCreationParametersResponseOutput{})
	pulumi.RegisterOutputType(BulkCreationParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(ComputeDataDiskResponseOutput{})
	pulumi.RegisterOutputType(ComputeDataDiskResponseArrayOutput{})
	pulumi.RegisterOutputType(ComputeVmInstanceViewStatusResponseOutput{})
	pulumi.RegisterOutputType(ComputeVmInstanceViewStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(ComputeVmPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesCustomOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesCustomPtrOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesCustomResponseOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesCustomResponsePtrOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesFromPlanOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesFromPlanPtrOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesFromPlanResponseOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesFromPlanResponsePtrOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesFromVmOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesFromVmPtrOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesFromVmResponseOutput{})
	pulumi.RegisterOutputType(CustomImagePropertiesFromVmResponsePtrOutput{})
	pulumi.RegisterOutputType(DataDiskPropertiesOutput{})
	pulumi.RegisterOutputType(DataDiskPropertiesArrayOutput{})
	pulumi.RegisterOutputType(DataDiskPropertiesResponseOutput{})
	pulumi.RegisterOutputType(DataDiskPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(DataDiskStorageTypeInfoOutput{})
	pulumi.RegisterOutputType(DataDiskStorageTypeInfoArrayOutput{})
	pulumi.RegisterOutputType(DataDiskStorageTypeInfoResponseOutput{})
	pulumi.RegisterOutputType(DataDiskStorageTypeInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(DayDetailsOutput{})
	pulumi.RegisterOutputType(DayDetailsPtrOutput{})
	pulumi.RegisterOutputType(DayDetailsResponseOutput{})
	pulumi.RegisterOutputType(DayDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(EnvironmentDeploymentPropertiesOutput{})
	pulumi.RegisterOutputType(EnvironmentDeploymentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EnvironmentDeploymentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentDeploymentPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(EventOutput{})
	pulumi.RegisterOutputType(EventArrayOutput{})
	pulumi.RegisterOutputType(EventResponseOutput{})
	pulumi.RegisterOutputType(EventResponseArrayOutput{})
	pulumi.RegisterOutputType(ExternalSubnetResponseOutput{})
	pulumi.RegisterOutputType(ExternalSubnetResponseArrayOutput{})
	pulumi.RegisterOutputType(FormulaPropertiesFromVmOutput{})
	pulumi.RegisterOutputType(FormulaPropertiesFromVmPtrOutput{})
	pulumi.RegisterOutputType(FormulaPropertiesFromVmResponseOutput{})
	pulumi.RegisterOutputType(FormulaPropertiesFromVmResponsePtrOutput{})
	pulumi.RegisterOutputType(GalleryImageReferenceOutput{})
	pulumi.RegisterOutputType(GalleryImageReferencePtrOutput{})
	pulumi.RegisterOutputType(GalleryImageReferenceResponseOutput{})
	pulumi.RegisterOutputType(GalleryImageReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(HourDetailsOutput{})
	pulumi.RegisterOutputType(HourDetailsPtrOutput{})
	pulumi.RegisterOutputType(HourDetailsResponseOutput{})
	pulumi.RegisterOutputType(HourDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesPtrOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(InboundNatRuleOutput{})
	pulumi.RegisterOutputType(InboundNatRuleArrayOutput{})
	pulumi.RegisterOutputType(InboundNatRuleResponseOutput{})
	pulumi.RegisterOutputType(InboundNatRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(LabAnnouncementPropertiesOutput{})
	pulumi.RegisterOutputType(LabAnnouncementPropertiesPtrOutput{})
	pulumi.RegisterOutputType(LabAnnouncementPropertiesResponseOutput{})
	pulumi.RegisterOutputType(LabAnnouncementPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LabSupportPropertiesOutput{})
	pulumi.RegisterOutputType(LabSupportPropertiesPtrOutput{})
	pulumi.RegisterOutputType(LabSupportPropertiesResponseOutput{})
	pulumi.RegisterOutputType(LabSupportPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LabVhdResponseOutput{})
	pulumi.RegisterOutputType(LabVhdResponseArrayOutput{})
	pulumi.RegisterOutputType(LabVirtualMachineCreationParameterOutput{})
	pulumi.RegisterOutputType(LabVirtualMachineCreationParameterPtrOutput{})
	pulumi.RegisterOutputType(LabVirtualMachineCreationParameterResponseOutput{})
	pulumi.RegisterOutputType(LabVirtualMachineCreationParameterResponsePtrOutput{})
	pulumi.RegisterOutputType(LinuxOsInfoOutput{})
	pulumi.RegisterOutputType(LinuxOsInfoPtrOutput{})
	pulumi.RegisterOutputType(LinuxOsInfoResponseOutput{})
	pulumi.RegisterOutputType(LinuxOsInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfacePropertiesOutput{})
	pulumi.RegisterOutputType(NetworkInterfacePropertiesPtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfacePropertiesResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfacePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(NotificationSettingsOutput{})
	pulumi.RegisterOutputType(NotificationSettingsPtrOutput{})
	pulumi.RegisterOutputType(NotificationSettingsResponseOutput{})
	pulumi.RegisterOutputType(NotificationSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(PortOutput{})
	pulumi.RegisterOutputType(PortArrayOutput{})
	pulumi.RegisterOutputType(PortResponseOutput{})
	pulumi.RegisterOutputType(PortResponseArrayOutput{})
	pulumi.RegisterOutputType(ScheduleCreationParameterOutput{})
	pulumi.RegisterOutputType(ScheduleCreationParameterArrayOutput{})
	pulumi.RegisterOutputType(ScheduleCreationParameterResponseOutput{})
	pulumi.RegisterOutputType(ScheduleCreationParameterResponseArrayOutput{})
	pulumi.RegisterOutputType(ScheduleResponseOutput{})
	pulumi.RegisterOutputType(ScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(ScheduleResponseArrayOutput{})
	pulumi.RegisterOutputType(SharedPublicIpAddressConfigurationOutput{})
	pulumi.RegisterOutputType(SharedPublicIpAddressConfigurationPtrOutput{})
	pulumi.RegisterOutputType(SharedPublicIpAddressConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SharedPublicIpAddressConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(SubnetOutput{})
	pulumi.RegisterOutputType(SubnetArrayOutput{})
	pulumi.RegisterOutputType(SubnetOverrideOutput{})
	pulumi.RegisterOutputType(SubnetOverrideArrayOutput{})
	pulumi.RegisterOutputType(SubnetOverrideResponseOutput{})
	pulumi.RegisterOutputType(SubnetOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(SubnetResponseOutput{})
	pulumi.RegisterOutputType(SubnetResponseArrayOutput{})
	pulumi.RegisterOutputType(SubnetSharedPublicIpAddressConfigurationOutput{})
	pulumi.RegisterOutputType(SubnetSharedPublicIpAddressConfigurationPtrOutput{})
	pulumi.RegisterOutputType(SubnetSharedPublicIpAddressConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SubnetSharedPublicIpAddressConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(UserIdentityOutput{})
	pulumi.RegisterOutputType(UserIdentityPtrOutput{})
	pulumi.RegisterOutputType(UserIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(UserSecretStoreOutput{})
	pulumi.RegisterOutputType(UserSecretStorePtrOutput{})
	pulumi.RegisterOutputType(UserSecretStoreResponseOutput{})
	pulumi.RegisterOutputType(UserSecretStoreResponsePtrOutput{})
	pulumi.RegisterOutputType(WeekDetailsOutput{})
	pulumi.RegisterOutputType(WeekDetailsPtrOutput{})
	pulumi.RegisterOutputType(WeekDetailsResponseOutput{})
	pulumi.RegisterOutputType(WeekDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(WindowsOsInfoOutput{})
	pulumi.RegisterOutputType(WindowsOsInfoPtrOutput{})
	pulumi.RegisterOutputType(WindowsOsInfoResponseOutput{})
	pulumi.RegisterOutputType(WindowsOsInfoResponsePtrOutput{})
}
