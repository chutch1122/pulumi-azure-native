


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SiteConnectionStringsSlot struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringPtrOutput                   `pulumi:"kind"`
	Location   pulumi.StringOutput                      `pulumi:"location"`
	Name       pulumi.StringPtrOutput                   `pulumi:"name"`
	Properties ConnStringValueTypePairResponseMapOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                   `pulumi:"tags"`
	Type       pulumi.StringPtrOutput                   `pulumi:"type"`
}


func NewSiteConnectionStringsSlot(ctx *pulumi.Context,
	name string, args *SiteConnectionStringsSlotArgs, opts ...pulumi.ResourceOption) (*SiteConnectionStringsSlot, error) {
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
			Type: pulumi.String("azure-native:web:SiteConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:SiteConnectionStringsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:SiteConnectionStringsSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteConnectionStringsSlot
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteConnectionStringsSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteConnectionStringsSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteConnectionStringsSlotState, opts ...pulumi.ResourceOption) (*SiteConnectionStringsSlot, error) {
	var resource SiteConnectionStringsSlot
	err := ctx.ReadResource("azure-native:web/v20150801:SiteConnectionStringsSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteConnectionStringsSlotState struct {
}

type SiteConnectionStringsSlotState struct {
}

func (SiteConnectionStringsSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteConnectionStringsSlotState)(nil)).Elem()
}

type siteConnectionStringsSlotArgs struct {
	Id                *string                            `pulumi:"id"`
	Kind              *string                            `pulumi:"kind"`
	Location          *string                            `pulumi:"location"`
	Name              string                             `pulumi:"name"`
	Properties        map[string]ConnStringValueTypePair `pulumi:"properties"`
	ResourceGroupName string                             `pulumi:"resourceGroupName"`
	Slot              string                             `pulumi:"slot"`
	Tags              map[string]string                  `pulumi:"tags"`
	Type              *string                            `pulumi:"type"`
}


type SiteConnectionStringsSlotArgs struct {
	Id                pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringInput
	Properties        ConnStringValueTypePairMapInput
	ResourceGroupName pulumi.StringInput
	Slot              pulumi.StringInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
}

func (SiteConnectionStringsSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteConnectionStringsSlotArgs)(nil)).Elem()
}

type SiteConnectionStringsSlotInput interface {
	pulumi.Input

	ToSiteConnectionStringsSlotOutput() SiteConnectionStringsSlotOutput
	ToSiteConnectionStringsSlotOutputWithContext(ctx context.Context) SiteConnectionStringsSlotOutput
}

func (*SiteConnectionStringsSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteConnectionStringsSlot)(nil)).Elem()
}

func (i *SiteConnectionStringsSlot) ToSiteConnectionStringsSlotOutput() SiteConnectionStringsSlotOutput {
	return i.ToSiteConnectionStringsSlotOutputWithContext(context.Background())
}

func (i *SiteConnectionStringsSlot) ToSiteConnectionStringsSlotOutputWithContext(ctx context.Context) SiteConnectionStringsSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConnectionStringsSlotOutput)
}

type SiteConnectionStringsSlotOutput struct{ *pulumi.OutputState }

func (SiteConnectionStringsSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteConnectionStringsSlot)(nil)).Elem()
}

func (o SiteConnectionStringsSlotOutput) ToSiteConnectionStringsSlotOutput() SiteConnectionStringsSlotOutput {
	return o
}

func (o SiteConnectionStringsSlotOutput) ToSiteConnectionStringsSlotOutputWithContext(ctx context.Context) SiteConnectionStringsSlotOutput {
	return o
}

func (o SiteConnectionStringsSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConnectionStringsSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SiteConnectionStringsSlotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteConnectionStringsSlot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SiteConnectionStringsSlotOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConnectionStringsSlot) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SiteConnectionStringsSlotOutput) Properties() ConnStringValueTypePairResponseMapOutput {
	return o.ApplyT(func(v *SiteConnectionStringsSlot) ConnStringValueTypePairResponseMapOutput { return v.Properties }).(ConnStringValueTypePairResponseMapOutput)
}

func (o SiteConnectionStringsSlotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SiteConnectionStringsSlot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SiteConnectionStringsSlotOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConnectionStringsSlot) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SiteConnectionStringsSlotOutput{})
}
