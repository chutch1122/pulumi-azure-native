


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Channel struct {
	pulumi.CustomResourceState

	ChannelFunctions pulumi.StringArrayOutput `pulumi:"channelFunctions"`
	ChannelType      pulumi.StringOutput      `pulumi:"channelType"`
	Credentials      pulumi.StringMapOutput   `pulumi:"credentials"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	Type             pulumi.StringOutput      `pulumi:"type"`
}


func NewChannel(ctx *pulumi.Context,
	name string, args *ChannelArgs, opts ...pulumi.ResourceOption) (*Channel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ChannelType == nil {
		return nil, errors.New("invalid value for required argument 'ChannelType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:engagementfabric:Channel"),
		},
	})
	opts = append(opts, aliases)
	var resource Channel
	err := ctx.RegisterResource("azure-native:engagementfabric/v20180901preview:Channel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelState, opts ...pulumi.ResourceOption) (*Channel, error) {
	var resource Channel
	err := ctx.ReadResource("azure-native:engagementfabric/v20180901preview:Channel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type channelState struct {
}

type ChannelState struct {
}

func (ChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelState)(nil)).Elem()
}

type channelArgs struct {
	AccountName       string            `pulumi:"accountName"`
	ChannelFunctions  []string          `pulumi:"channelFunctions"`
	ChannelName       *string           `pulumi:"channelName"`
	ChannelType       string            `pulumi:"channelType"`
	Credentials       map[string]string `pulumi:"credentials"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
}


type ChannelArgs struct {
	AccountName       pulumi.StringInput
	ChannelFunctions  pulumi.StringArrayInput
	ChannelName       pulumi.StringPtrInput
	ChannelType       pulumi.StringInput
	Credentials       pulumi.StringMapInput
	ResourceGroupName pulumi.StringInput
}

func (ChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelArgs)(nil)).Elem()
}

type ChannelInput interface {
	pulumi.Input

	ToChannelOutput() ChannelOutput
	ToChannelOutputWithContext(ctx context.Context) ChannelOutput
}

func (*Channel) ElementType() reflect.Type {
	return reflect.TypeOf((**Channel)(nil)).Elem()
}

func (i *Channel) ToChannelOutput() ChannelOutput {
	return i.ToChannelOutputWithContext(context.Background())
}

func (i *Channel) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelOutput)
}

type ChannelOutput struct{ *pulumi.OutputState }

func (ChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Channel)(nil)).Elem()
}

func (o ChannelOutput) ToChannelOutput() ChannelOutput {
	return o
}

func (o ChannelOutput) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return o
}

func (o ChannelOutput) ChannelFunctions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringArrayOutput { return v.ChannelFunctions }).(pulumi.StringArrayOutput)
}

func (o ChannelOutput) ChannelType() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.ChannelType }).(pulumi.StringOutput)
}

func (o ChannelOutput) Credentials() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringMapOutput { return v.Credentials }).(pulumi.StringMapOutput)
}

func (o ChannelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ChannelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ChannelOutput{})
}
