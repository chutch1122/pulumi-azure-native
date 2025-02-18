


package v20201001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppSitePushSettings struct {
	pulumi.CustomResourceState

	DynamicTagsJson   pulumi.StringPtrOutput   `pulumi:"dynamicTagsJson"`
	IsPushEnabled     pulumi.BoolOutput        `pulumi:"isPushEnabled"`
	Kind              pulumi.StringPtrOutput   `pulumi:"kind"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	TagWhitelistJson  pulumi.StringPtrOutput   `pulumi:"tagWhitelistJson"`
	TagsRequiringAuth pulumi.StringPtrOutput   `pulumi:"tagsRequiringAuth"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewWebAppSitePushSettings(ctx *pulumi.Context,
	name string, args *WebAppSitePushSettingsArgs, opts ...pulumi.ResourceOption) (*WebAppSitePushSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IsPushEnabled == nil {
		return nil, errors.New("invalid value for required argument 'IsPushEnabled'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppSitePushSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppSitePushSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSitePushSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppSitePushSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppSitePushSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppSitePushSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppSitePushSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppSitePushSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppSitePushSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppSitePushSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSitePushSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppSitePushSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppSitePushSettings"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppSitePushSettings
	err := ctx.RegisterResource("azure-native:web/v20201001:WebAppSitePushSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppSitePushSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSitePushSettingsState, opts ...pulumi.ResourceOption) (*WebAppSitePushSettings, error) {
	var resource WebAppSitePushSettings
	err := ctx.ReadResource("azure-native:web/v20201001:WebAppSitePushSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppSitePushSettingsState struct {
}

type WebAppSitePushSettingsState struct {
}

func (WebAppSitePushSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSitePushSettingsState)(nil)).Elem()
}

type webAppSitePushSettingsArgs struct {
	DynamicTagsJson   *string `pulumi:"dynamicTagsJson"`
	IsPushEnabled     bool    `pulumi:"isPushEnabled"`
	Kind              *string `pulumi:"kind"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	TagWhitelistJson  *string `pulumi:"tagWhitelistJson"`
	TagsRequiringAuth *string `pulumi:"tagsRequiringAuth"`
}


type WebAppSitePushSettingsArgs struct {
	DynamicTagsJson   pulumi.StringPtrInput
	IsPushEnabled     pulumi.BoolInput
	Kind              pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	TagWhitelistJson  pulumi.StringPtrInput
	TagsRequiringAuth pulumi.StringPtrInput
}

func (WebAppSitePushSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSitePushSettingsArgs)(nil)).Elem()
}

type WebAppSitePushSettingsInput interface {
	pulumi.Input

	ToWebAppSitePushSettingsOutput() WebAppSitePushSettingsOutput
	ToWebAppSitePushSettingsOutputWithContext(ctx context.Context) WebAppSitePushSettingsOutput
}

func (*WebAppSitePushSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSitePushSettings)(nil)).Elem()
}

func (i *WebAppSitePushSettings) ToWebAppSitePushSettingsOutput() WebAppSitePushSettingsOutput {
	return i.ToWebAppSitePushSettingsOutputWithContext(context.Background())
}

func (i *WebAppSitePushSettings) ToWebAppSitePushSettingsOutputWithContext(ctx context.Context) WebAppSitePushSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSitePushSettingsOutput)
}

type WebAppSitePushSettingsOutput struct{ *pulumi.OutputState }

func (WebAppSitePushSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSitePushSettings)(nil)).Elem()
}

func (o WebAppSitePushSettingsOutput) ToWebAppSitePushSettingsOutput() WebAppSitePushSettingsOutput {
	return o
}

func (o WebAppSitePushSettingsOutput) ToWebAppSitePushSettingsOutputWithContext(ctx context.Context) WebAppSitePushSettingsOutput {
	return o
}

func (o WebAppSitePushSettingsOutput) DynamicTagsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSitePushSettings) pulumi.StringPtrOutput { return v.DynamicTagsJson }).(pulumi.StringPtrOutput)
}

func (o WebAppSitePushSettingsOutput) IsPushEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *WebAppSitePushSettings) pulumi.BoolOutput { return v.IsPushEnabled }).(pulumi.BoolOutput)
}

func (o WebAppSitePushSettingsOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSitePushSettings) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppSitePushSettingsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSitePushSettings) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppSitePushSettingsOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebAppSitePushSettings) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o WebAppSitePushSettingsOutput) TagWhitelistJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSitePushSettings) pulumi.StringPtrOutput { return v.TagWhitelistJson }).(pulumi.StringPtrOutput)
}

func (o WebAppSitePushSettingsOutput) TagsRequiringAuth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSitePushSettings) pulumi.StringPtrOutput { return v.TagsRequiringAuth }).(pulumi.StringPtrOutput)
}

func (o WebAppSitePushSettingsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSitePushSettings) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppSitePushSettingsOutput{})
}
