


package v20210201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppMetadataSlot struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringPtrOutput `pulumi:"kind"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	Properties pulumi.StringMapOutput `pulumi:"properties"`
	Type       pulumi.StringOutput    `pulumi:"type"`
}


func NewWebAppMetadataSlot(ctx *pulumi.Context,
	name string, args *WebAppMetadataSlotArgs, opts ...pulumi.ResourceOption) (*WebAppMetadataSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppMetadataSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppMetadataSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppMetadataSlot
	err := ctx.RegisterResource("azure-native:web/v20210201:WebAppMetadataSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppMetadataSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppMetadataSlotState, opts ...pulumi.ResourceOption) (*WebAppMetadataSlot, error) {
	var resource WebAppMetadataSlot
	err := ctx.ReadResource("azure-native:web/v20210201:WebAppMetadataSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppMetadataSlotState struct {
}

type WebAppMetadataSlotState struct {
}

func (WebAppMetadataSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppMetadataSlotState)(nil)).Elem()
}

type webAppMetadataSlotArgs struct {
	Kind              *string           `pulumi:"kind"`
	Name              string            `pulumi:"name"`
	Properties        map[string]string `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Slot              string            `pulumi:"slot"`
}


type WebAppMetadataSlotArgs struct {
	Kind              pulumi.StringPtrInput
	Name              pulumi.StringInput
	Properties        pulumi.StringMapInput
	ResourceGroupName pulumi.StringInput
	Slot              pulumi.StringInput
}

func (WebAppMetadataSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppMetadataSlotArgs)(nil)).Elem()
}

type WebAppMetadataSlotInput interface {
	pulumi.Input

	ToWebAppMetadataSlotOutput() WebAppMetadataSlotOutput
	ToWebAppMetadataSlotOutputWithContext(ctx context.Context) WebAppMetadataSlotOutput
}

func (*WebAppMetadataSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppMetadataSlot)(nil)).Elem()
}

func (i *WebAppMetadataSlot) ToWebAppMetadataSlotOutput() WebAppMetadataSlotOutput {
	return i.ToWebAppMetadataSlotOutputWithContext(context.Background())
}

func (i *WebAppMetadataSlot) ToWebAppMetadataSlotOutputWithContext(ctx context.Context) WebAppMetadataSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppMetadataSlotOutput)
}

type WebAppMetadataSlotOutput struct{ *pulumi.OutputState }

func (WebAppMetadataSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppMetadataSlot)(nil)).Elem()
}

func (o WebAppMetadataSlotOutput) ToWebAppMetadataSlotOutput() WebAppMetadataSlotOutput {
	return o
}

func (o WebAppMetadataSlotOutput) ToWebAppMetadataSlotOutputWithContext(ctx context.Context) WebAppMetadataSlotOutput {
	return o
}

func (o WebAppMetadataSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppMetadataSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppMetadataSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppMetadataSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppMetadataSlotOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebAppMetadataSlot) pulumi.StringMapOutput { return v.Properties }).(pulumi.StringMapOutput)
}

func (o WebAppMetadataSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppMetadataSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppMetadataSlotOutput{})
}
