


package v20170901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Factory struct {
	pulumi.CustomResourceState

	CreateTime        pulumi.StringOutput                       `pulumi:"createTime"`
	Identity          FactoryIdentityResponsePtrOutput          `pulumi:"identity"`
	Location          pulumi.StringPtrOutput                    `pulumi:"location"`
	Name              pulumi.StringOutput                       `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                       `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput                    `pulumi:"tags"`
	Type              pulumi.StringOutput                       `pulumi:"type"`
	Version           pulumi.StringOutput                       `pulumi:"version"`
	VstsConfiguration FactoryVSTSConfigurationResponsePtrOutput `pulumi:"vstsConfiguration"`
}


func NewFactory(ctx *pulumi.Context,
	name string, args *FactoryArgs, opts ...pulumi.ResourceOption) (*Factory, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datafactory:Factory"),
		},
		{
			Type: pulumi.String("azure-native:datafactory/v20180601:Factory"),
		},
	})
	opts = append(opts, aliases)
	var resource Factory
	err := ctx.RegisterResource("azure-native:datafactory/v20170901preview:Factory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFactory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FactoryState, opts ...pulumi.ResourceOption) (*Factory, error) {
	var resource Factory
	err := ctx.ReadResource("azure-native:datafactory/v20170901preview:Factory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type factoryState struct {
}

type FactoryState struct {
}

func (FactoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*factoryState)(nil)).Elem()
}

type factoryArgs struct {
	FactoryName       *string                   `pulumi:"factoryName"`
	Identity          *FactoryIdentity          `pulumi:"identity"`
	Location          *string                   `pulumi:"location"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	Tags              map[string]string         `pulumi:"tags"`
	VstsConfiguration *FactoryVSTSConfiguration `pulumi:"vstsConfiguration"`
}


type FactoryArgs struct {
	FactoryName       pulumi.StringPtrInput
	Identity          FactoryIdentityPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VstsConfiguration FactoryVSTSConfigurationPtrInput
}

func (FactoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*factoryArgs)(nil)).Elem()
}

type FactoryInput interface {
	pulumi.Input

	ToFactoryOutput() FactoryOutput
	ToFactoryOutputWithContext(ctx context.Context) FactoryOutput
}

func (*Factory) ElementType() reflect.Type {
	return reflect.TypeOf((**Factory)(nil)).Elem()
}

func (i *Factory) ToFactoryOutput() FactoryOutput {
	return i.ToFactoryOutputWithContext(context.Background())
}

func (i *Factory) ToFactoryOutputWithContext(ctx context.Context) FactoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FactoryOutput)
}

type FactoryOutput struct{ *pulumi.OutputState }

func (FactoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Factory)(nil)).Elem()
}

func (o FactoryOutput) ToFactoryOutput() FactoryOutput {
	return o
}

func (o FactoryOutput) ToFactoryOutputWithContext(ctx context.Context) FactoryOutput {
	return o
}

func (o FactoryOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Factory) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o FactoryOutput) Identity() FactoryIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Factory) FactoryIdentityResponsePtrOutput { return v.Identity }).(FactoryIdentityResponsePtrOutput)
}

func (o FactoryOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Factory) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o FactoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Factory) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FactoryOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Factory) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o FactoryOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Factory) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o FactoryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Factory) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o FactoryOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Factory) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func (o FactoryOutput) VstsConfiguration() FactoryVSTSConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Factory) FactoryVSTSConfigurationResponsePtrOutput { return v.VstsConfiguration }).(FactoryVSTSConfigurationResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(FactoryOutput{})
}
