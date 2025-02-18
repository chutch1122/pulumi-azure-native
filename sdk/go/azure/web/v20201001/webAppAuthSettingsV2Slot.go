


package v20201001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppAuthSettingsV2Slot struct {
	pulumi.CustomResourceState

	GlobalValidation  GlobalValidationResponsePtrOutput  `pulumi:"globalValidation"`
	HttpSettings      HttpSettingsResponsePtrOutput      `pulumi:"httpSettings"`
	IdentityProviders IdentityProvidersResponsePtrOutput `pulumi:"identityProviders"`
	Kind              pulumi.StringPtrOutput             `pulumi:"kind"`
	Login             LoginResponsePtrOutput             `pulumi:"login"`
	Name              pulumi.StringOutput                `pulumi:"name"`
	Platform          AuthPlatformResponsePtrOutput      `pulumi:"platform"`
	SystemData        SystemDataResponseOutput           `pulumi:"systemData"`
	Type              pulumi.StringOutput                `pulumi:"type"`
}


func NewWebAppAuthSettingsV2Slot(ctx *pulumi.Context,
	name string, args *WebAppAuthSettingsV2SlotArgs, opts ...pulumi.ResourceOption) (*WebAppAuthSettingsV2Slot, error) {
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
			Type: pulumi.String("azure-native:web:WebAppAuthSettingsV2Slot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppAuthSettingsV2Slot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppAuthSettingsV2Slot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppAuthSettingsV2Slot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppAuthSettingsV2Slot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppAuthSettingsV2Slot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppAuthSettingsV2Slot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppAuthSettingsV2Slot
	err := ctx.RegisterResource("azure-native:web/v20201001:WebAppAuthSettingsV2Slot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppAuthSettingsV2Slot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppAuthSettingsV2SlotState, opts ...pulumi.ResourceOption) (*WebAppAuthSettingsV2Slot, error) {
	var resource WebAppAuthSettingsV2Slot
	err := ctx.ReadResource("azure-native:web/v20201001:WebAppAuthSettingsV2Slot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppAuthSettingsV2SlotState struct {
}

type WebAppAuthSettingsV2SlotState struct {
}

func (WebAppAuthSettingsV2SlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAuthSettingsV2SlotState)(nil)).Elem()
}

type webAppAuthSettingsV2SlotArgs struct {
	GlobalValidation  *GlobalValidation  `pulumi:"globalValidation"`
	HttpSettings      *HttpSettings      `pulumi:"httpSettings"`
	IdentityProviders *IdentityProviders `pulumi:"identityProviders"`
	Kind              *string            `pulumi:"kind"`
	Login             *Login             `pulumi:"login"`
	Name              string             `pulumi:"name"`
	Platform          *AuthPlatform      `pulumi:"platform"`
	ResourceGroupName string             `pulumi:"resourceGroupName"`
	Slot              string             `pulumi:"slot"`
}


type WebAppAuthSettingsV2SlotArgs struct {
	GlobalValidation  GlobalValidationPtrInput
	HttpSettings      HttpSettingsPtrInput
	IdentityProviders IdentityProvidersPtrInput
	Kind              pulumi.StringPtrInput
	Login             LoginPtrInput
	Name              pulumi.StringInput
	Platform          AuthPlatformPtrInput
	ResourceGroupName pulumi.StringInput
	Slot              pulumi.StringInput
}

func (WebAppAuthSettingsV2SlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAuthSettingsV2SlotArgs)(nil)).Elem()
}

type WebAppAuthSettingsV2SlotInput interface {
	pulumi.Input

	ToWebAppAuthSettingsV2SlotOutput() WebAppAuthSettingsV2SlotOutput
	ToWebAppAuthSettingsV2SlotOutputWithContext(ctx context.Context) WebAppAuthSettingsV2SlotOutput
}

func (*WebAppAuthSettingsV2Slot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppAuthSettingsV2Slot)(nil)).Elem()
}

func (i *WebAppAuthSettingsV2Slot) ToWebAppAuthSettingsV2SlotOutput() WebAppAuthSettingsV2SlotOutput {
	return i.ToWebAppAuthSettingsV2SlotOutputWithContext(context.Background())
}

func (i *WebAppAuthSettingsV2Slot) ToWebAppAuthSettingsV2SlotOutputWithContext(ctx context.Context) WebAppAuthSettingsV2SlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppAuthSettingsV2SlotOutput)
}

type WebAppAuthSettingsV2SlotOutput struct{ *pulumi.OutputState }

func (WebAppAuthSettingsV2SlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppAuthSettingsV2Slot)(nil)).Elem()
}

func (o WebAppAuthSettingsV2SlotOutput) ToWebAppAuthSettingsV2SlotOutput() WebAppAuthSettingsV2SlotOutput {
	return o
}

func (o WebAppAuthSettingsV2SlotOutput) ToWebAppAuthSettingsV2SlotOutputWithContext(ctx context.Context) WebAppAuthSettingsV2SlotOutput {
	return o
}

func (o WebAppAuthSettingsV2SlotOutput) GlobalValidation() GlobalValidationResponsePtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2Slot) GlobalValidationResponsePtrOutput { return v.GlobalValidation }).(GlobalValidationResponsePtrOutput)
}

func (o WebAppAuthSettingsV2SlotOutput) HttpSettings() HttpSettingsResponsePtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2Slot) HttpSettingsResponsePtrOutput { return v.HttpSettings }).(HttpSettingsResponsePtrOutput)
}

func (o WebAppAuthSettingsV2SlotOutput) IdentityProviders() IdentityProvidersResponsePtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2Slot) IdentityProvidersResponsePtrOutput { return v.IdentityProviders }).(IdentityProvidersResponsePtrOutput)
}

func (o WebAppAuthSettingsV2SlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2Slot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsV2SlotOutput) Login() LoginResponsePtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2Slot) LoginResponsePtrOutput { return v.Login }).(LoginResponsePtrOutput)
}

func (o WebAppAuthSettingsV2SlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2Slot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppAuthSettingsV2SlotOutput) Platform() AuthPlatformResponsePtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2Slot) AuthPlatformResponsePtrOutput { return v.Platform }).(AuthPlatformResponsePtrOutput)
}

func (o WebAppAuthSettingsV2SlotOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2Slot) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o WebAppAuthSettingsV2SlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2Slot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppAuthSettingsV2SlotOutput{})
}
