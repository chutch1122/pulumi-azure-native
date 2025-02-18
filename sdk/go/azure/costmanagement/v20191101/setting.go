


package v20191101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Setting struct {
	pulumi.CustomResourceState

	Cache   SettingsPropertiesResponseCacheArrayOutput `pulumi:"cache"`
	Kind    pulumi.StringOutput                        `pulumi:"kind"`
	Name    pulumi.StringOutput                        `pulumi:"name"`
	Scope   pulumi.StringOutput                        `pulumi:"scope"`
	StartOn pulumi.StringPtrOutput                     `pulumi:"startOn"`
	Type    pulumi.StringOutput                        `pulumi:"type"`
}


func NewSetting(ctx *pulumi.Context,
	name string, args *SettingArgs, opts ...pulumi.ResourceOption) (*Setting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement:Setting"),
		},
	})
	opts = append(opts, aliases)
	var resource Setting
	err := ctx.RegisterResource("azure-native:costmanagement/v20191101:Setting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SettingState, opts ...pulumi.ResourceOption) (*Setting, error) {
	var resource Setting
	err := ctx.ReadResource("azure-native:costmanagement/v20191101:Setting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type settingState struct {
}

type SettingState struct {
}

func (SettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*settingState)(nil)).Elem()
}

type settingArgs struct {
	Cache       []SettingsPropertiesCache `pulumi:"cache"`
	Scope       string                    `pulumi:"scope"`
	SettingName *string                   `pulumi:"settingName"`
	StartOn     *string                   `pulumi:"startOn"`
}


type SettingArgs struct {
	Cache       SettingsPropertiesCacheArrayInput
	Scope       pulumi.StringInput
	SettingName pulumi.StringPtrInput
	StartOn     pulumi.StringPtrInput
}

func (SettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*settingArgs)(nil)).Elem()
}

type SettingInput interface {
	pulumi.Input

	ToSettingOutput() SettingOutput
	ToSettingOutputWithContext(ctx context.Context) SettingOutput
}

func (*Setting) ElementType() reflect.Type {
	return reflect.TypeOf((**Setting)(nil)).Elem()
}

func (i *Setting) ToSettingOutput() SettingOutput {
	return i.ToSettingOutputWithContext(context.Background())
}

func (i *Setting) ToSettingOutputWithContext(ctx context.Context) SettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingOutput)
}

type SettingOutput struct{ *pulumi.OutputState }

func (SettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Setting)(nil)).Elem()
}

func (o SettingOutput) ToSettingOutput() SettingOutput {
	return o
}

func (o SettingOutput) ToSettingOutputWithContext(ctx context.Context) SettingOutput {
	return o
}

func (o SettingOutput) Cache() SettingsPropertiesResponseCacheArrayOutput {
	return o.ApplyT(func(v *Setting) SettingsPropertiesResponseCacheArrayOutput { return v.Cache }).(SettingsPropertiesResponseCacheArrayOutput)
}

func (o SettingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o SettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SettingOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

func (o SettingOutput) StartOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringPtrOutput { return v.StartOn }).(pulumi.StringPtrOutput)
}

func (o SettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Setting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SettingOutput{})
}
