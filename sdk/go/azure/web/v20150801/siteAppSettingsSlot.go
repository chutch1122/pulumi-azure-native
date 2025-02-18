


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SiteAppSettingsSlot struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringPtrOutput `pulumi:"kind"`
	Location   pulumi.StringOutput    `pulumi:"location"`
	Name       pulumi.StringPtrOutput `pulumi:"name"`
	Properties pulumi.StringMapOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput `pulumi:"tags"`
	Type       pulumi.StringPtrOutput `pulumi:"type"`
}


func NewSiteAppSettingsSlot(ctx *pulumi.Context,
	name string, args *SiteAppSettingsSlotArgs, opts ...pulumi.ResourceOption) (*SiteAppSettingsSlot, error) {
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
			Type: pulumi.String("azure-native:web:SiteAppSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteAppSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteAppSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteAppSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteAppSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteAppSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteAppSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteAppSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteAppSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteAppSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteAppSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteAppSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:SiteAppSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:SiteAppSettingsSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteAppSettingsSlot
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteAppSettingsSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteAppSettingsSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteAppSettingsSlotState, opts ...pulumi.ResourceOption) (*SiteAppSettingsSlot, error) {
	var resource SiteAppSettingsSlot
	err := ctx.ReadResource("azure-native:web/v20150801:SiteAppSettingsSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteAppSettingsSlotState struct {
}

type SiteAppSettingsSlotState struct {
}

func (SiteAppSettingsSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteAppSettingsSlotState)(nil)).Elem()
}

type siteAppSettingsSlotArgs struct {
	Id                *string           `pulumi:"id"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	Properties        map[string]string `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Slot              string            `pulumi:"slot"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
}


type SiteAppSettingsSlotArgs struct {
	Id                pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringInput
	Properties        pulumi.StringMapInput
	ResourceGroupName pulumi.StringInput
	Slot              pulumi.StringInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
}

func (SiteAppSettingsSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteAppSettingsSlotArgs)(nil)).Elem()
}

type SiteAppSettingsSlotInput interface {
	pulumi.Input

	ToSiteAppSettingsSlotOutput() SiteAppSettingsSlotOutput
	ToSiteAppSettingsSlotOutputWithContext(ctx context.Context) SiteAppSettingsSlotOutput
}

func (*SiteAppSettingsSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteAppSettingsSlot)(nil)).Elem()
}

func (i *SiteAppSettingsSlot) ToSiteAppSettingsSlotOutput() SiteAppSettingsSlotOutput {
	return i.ToSiteAppSettingsSlotOutputWithContext(context.Background())
}

func (i *SiteAppSettingsSlot) ToSiteAppSettingsSlotOutputWithContext(ctx context.Context) SiteAppSettingsSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteAppSettingsSlotOutput)
}

type SiteAppSettingsSlotOutput struct{ *pulumi.OutputState }

func (SiteAppSettingsSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteAppSettingsSlot)(nil)).Elem()
}

func (o SiteAppSettingsSlotOutput) ToSiteAppSettingsSlotOutput() SiteAppSettingsSlotOutput {
	return o
}

func (o SiteAppSettingsSlotOutput) ToSiteAppSettingsSlotOutputWithContext(ctx context.Context) SiteAppSettingsSlotOutput {
	return o
}

func (o SiteAppSettingsSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAppSettingsSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SiteAppSettingsSlotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteAppSettingsSlot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SiteAppSettingsSlotOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAppSettingsSlot) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SiteAppSettingsSlotOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SiteAppSettingsSlot) pulumi.StringMapOutput { return v.Properties }).(pulumi.StringMapOutput)
}

func (o SiteAppSettingsSlotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SiteAppSettingsSlot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SiteAppSettingsSlotOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAppSettingsSlot) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SiteAppSettingsSlotOutput{})
}
