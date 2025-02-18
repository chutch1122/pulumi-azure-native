


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SiteAppSettings struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringPtrOutput `pulumi:"kind"`
	Location   pulumi.StringOutput    `pulumi:"location"`
	Name       pulumi.StringPtrOutput `pulumi:"name"`
	Properties pulumi.StringMapOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput `pulumi:"tags"`
	Type       pulumi.StringPtrOutput `pulumi:"type"`
}


func NewSiteAppSettings(ctx *pulumi.Context,
	name string, args *SiteAppSettingsArgs, opts ...pulumi.ResourceOption) (*SiteAppSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:SiteAppSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteAppSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteAppSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteAppSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteAppSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteAppSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteAppSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteAppSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteAppSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteAppSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteAppSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteAppSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:SiteAppSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:SiteAppSettings"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteAppSettings
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteAppSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteAppSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteAppSettingsState, opts ...pulumi.ResourceOption) (*SiteAppSettings, error) {
	var resource SiteAppSettings
	err := ctx.ReadResource("azure-native:web/v20150801:SiteAppSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteAppSettingsState struct {
}

type SiteAppSettingsState struct {
}

func (SiteAppSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteAppSettingsState)(nil)).Elem()
}

type siteAppSettingsArgs struct {
	Id                *string           `pulumi:"id"`
	Kind              *string           `pulumi:"kind"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	Properties        map[string]string `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
}


type SiteAppSettingsArgs struct {
	Id                pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringInput
	Properties        pulumi.StringMapInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
}

func (SiteAppSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteAppSettingsArgs)(nil)).Elem()
}

type SiteAppSettingsInput interface {
	pulumi.Input

	ToSiteAppSettingsOutput() SiteAppSettingsOutput
	ToSiteAppSettingsOutputWithContext(ctx context.Context) SiteAppSettingsOutput
}

func (*SiteAppSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteAppSettings)(nil)).Elem()
}

func (i *SiteAppSettings) ToSiteAppSettingsOutput() SiteAppSettingsOutput {
	return i.ToSiteAppSettingsOutputWithContext(context.Background())
}

func (i *SiteAppSettings) ToSiteAppSettingsOutputWithContext(ctx context.Context) SiteAppSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteAppSettingsOutput)
}

type SiteAppSettingsOutput struct{ *pulumi.OutputState }

func (SiteAppSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteAppSettings)(nil)).Elem()
}

func (o SiteAppSettingsOutput) ToSiteAppSettingsOutput() SiteAppSettingsOutput {
	return o
}

func (o SiteAppSettingsOutput) ToSiteAppSettingsOutputWithContext(ctx context.Context) SiteAppSettingsOutput {
	return o
}

func (o SiteAppSettingsOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAppSettings) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SiteAppSettingsOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteAppSettings) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SiteAppSettingsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAppSettings) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SiteAppSettingsOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SiteAppSettings) pulumi.StringMapOutput { return v.Properties }).(pulumi.StringMapOutput)
}

func (o SiteAppSettingsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SiteAppSettings) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SiteAppSettingsOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAppSettings) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SiteAppSettingsOutput{})
}
