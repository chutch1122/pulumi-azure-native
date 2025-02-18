


package v20180915

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Environment struct {
	pulumi.CustomResourceState

	ArmTemplateDisplayName pulumi.StringPtrOutput                           `pulumi:"armTemplateDisplayName"`
	CreatedByUser          pulumi.StringOutput                              `pulumi:"createdByUser"`
	DeploymentProperties   EnvironmentDeploymentPropertiesResponsePtrOutput `pulumi:"deploymentProperties"`
	Location               pulumi.StringPtrOutput                           `pulumi:"location"`
	Name                   pulumi.StringOutput                              `pulumi:"name"`
	ProvisioningState      pulumi.StringOutput                              `pulumi:"provisioningState"`
	ResourceGroupId        pulumi.StringOutput                              `pulumi:"resourceGroupId"`
	Tags                   pulumi.StringMapOutput                           `pulumi:"tags"`
	Type                   pulumi.StringOutput                              `pulumi:"type"`
	UniqueIdentifier       pulumi.StringOutput                              `pulumi:"uniqueIdentifier"`
}


func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:Environment"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:Environment"),
		},
	})
	opts = append(opts, aliases)
	var resource Environment
	err := ctx.RegisterResource("azure-native:devtestlab/v20180915:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("azure-native:devtestlab/v20180915:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type environmentState struct {
}

type EnvironmentState struct {
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	ArmTemplateDisplayName *string                          `pulumi:"armTemplateDisplayName"`
	DeploymentProperties   *EnvironmentDeploymentProperties `pulumi:"deploymentProperties"`
	LabName                string                           `pulumi:"labName"`
	Location               *string                          `pulumi:"location"`
	Name                   *string                          `pulumi:"name"`
	ResourceGroupName      string                           `pulumi:"resourceGroupName"`
	Tags                   map[string]string                `pulumi:"tags"`
	UserName               string                           `pulumi:"userName"`
}


type EnvironmentArgs struct {
	ArmTemplateDisplayName pulumi.StringPtrInput
	DeploymentProperties   EnvironmentDeploymentPropertiesPtrInput
	LabName                pulumi.StringInput
	Location               pulumi.StringPtrInput
	Name                   pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
	UserName               pulumi.StringInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}

type EnvironmentInput interface {
	pulumi.Input

	ToEnvironmentOutput() EnvironmentOutput
	ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput
}

func (*Environment) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (i *Environment) ToEnvironmentOutput() EnvironmentOutput {
	return i.ToEnvironmentOutputWithContext(context.Background())
}

func (i *Environment) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentOutput)
}

type EnvironmentOutput struct{ *pulumi.OutputState }

func (EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (o EnvironmentOutput) ToEnvironmentOutput() EnvironmentOutput {
	return o
}

func (o EnvironmentOutput) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return o
}

func (o EnvironmentOutput) ArmTemplateDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.ArmTemplateDisplayName }).(pulumi.StringPtrOutput)
}

func (o EnvironmentOutput) CreatedByUser() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.CreatedByUser }).(pulumi.StringOutput)
}

func (o EnvironmentOutput) DeploymentProperties() EnvironmentDeploymentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Environment) EnvironmentDeploymentPropertiesResponsePtrOutput { return v.DeploymentProperties }).(EnvironmentDeploymentPropertiesResponsePtrOutput)
}

func (o EnvironmentOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o EnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EnvironmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o EnvironmentOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

func (o EnvironmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o EnvironmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o EnvironmentOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EnvironmentOutput{})
}
