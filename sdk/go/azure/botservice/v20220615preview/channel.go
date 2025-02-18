


package v20220615preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Channel struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput   `pulumi:"etag"`
	Kind       pulumi.StringPtrOutput   `pulumi:"kind"`
	Location   pulumi.StringPtrOutput   `pulumi:"location"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	Properties pulumi.AnyOutput         `pulumi:"properties"`
	Sku        SkuResponsePtrOutput     `pulumi:"sku"`
	Tags       pulumi.StringMapOutput   `pulumi:"tags"`
	Type       pulumi.StringOutput      `pulumi:"type"`
	Zones      pulumi.StringArrayOutput `pulumi:"zones"`
}


func NewChannel(ctx *pulumi.Context,
	name string, args *ChannelArgs, opts ...pulumi.ResourceOption) (*Channel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:botservice:Channel"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20171201:Channel"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20180712:Channel"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20200602:Channel"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20210301:Channel"),
		},
		{
			Type: pulumi.String("azure-native:botservice/v20210501preview:Channel"),
		},
	})
	opts = append(opts, aliases)
	var resource Channel
	err := ctx.RegisterResource("azure-native:botservice/v20220615preview:Channel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelState, opts ...pulumi.ResourceOption) (*Channel, error) {
	var resource Channel
	err := ctx.ReadResource("azure-native:botservice/v20220615preview:Channel", name, id, state, &resource, opts...)
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
	ChannelName       *string           `pulumi:"channelName"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	Properties        interface{}       `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      string            `pulumi:"resourceName"`
	Sku               *Sku              `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}


type ChannelArgs struct {
	ChannelName       pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Properties        pulumi.Input
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringInput
	Sku               SkuPtrInput
	Tags              pulumi.StringMapInput
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

func (o ChannelOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ChannelOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ChannelOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ChannelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ChannelOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Channel) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

func (o ChannelOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Channel) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o ChannelOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ChannelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ChannelOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ChannelOutput{})
}
