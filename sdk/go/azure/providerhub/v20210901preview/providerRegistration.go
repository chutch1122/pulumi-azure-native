


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProviderRegistration struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                          `pulumi:"name"`
	Properties ProviderRegistrationResponsePropertiesOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                     `pulumi:"systemData"`
	Type       pulumi.StringOutput                          `pulumi:"type"`
}


func NewProviderRegistration(ctx *pulumi.Context,
	name string, args *ProviderRegistrationArgs, opts ...pulumi.ResourceOption) (*ProviderRegistration, error) {
	if args == nil {
		args = &ProviderRegistrationArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:providerhub:ProviderRegistration"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20201120:ProviderRegistration"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210501preview:ProviderRegistration"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210601preview:ProviderRegistration"),
		},
	})
	opts = append(opts, aliases)
	var resource ProviderRegistration
	err := ctx.RegisterResource("azure-native:providerhub/v20210901preview:ProviderRegistration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProviderRegistration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProviderRegistrationState, opts ...pulumi.ResourceOption) (*ProviderRegistration, error) {
	var resource ProviderRegistration
	err := ctx.ReadResource("azure-native:providerhub/v20210901preview:ProviderRegistration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type providerRegistrationState struct {
}

type ProviderRegistrationState struct {
}

func (ProviderRegistrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*providerRegistrationState)(nil)).Elem()
}

type providerRegistrationArgs struct {
	Properties        *ProviderRegistrationProperties `pulumi:"properties"`
	ProviderNamespace *string                         `pulumi:"providerNamespace"`
}


type ProviderRegistrationArgs struct {
	Properties        ProviderRegistrationPropertiesPtrInput
	ProviderNamespace pulumi.StringPtrInput
}

func (ProviderRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerRegistrationArgs)(nil)).Elem()
}

type ProviderRegistrationInput interface {
	pulumi.Input

	ToProviderRegistrationOutput() ProviderRegistrationOutput
	ToProviderRegistrationOutputWithContext(ctx context.Context) ProviderRegistrationOutput
}

func (*ProviderRegistration) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderRegistration)(nil)).Elem()
}

func (i *ProviderRegistration) ToProviderRegistrationOutput() ProviderRegistrationOutput {
	return i.ToProviderRegistrationOutputWithContext(context.Background())
}

func (i *ProviderRegistration) ToProviderRegistrationOutputWithContext(ctx context.Context) ProviderRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderRegistrationOutput)
}

type ProviderRegistrationOutput struct{ *pulumi.OutputState }

func (ProviderRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderRegistration)(nil)).Elem()
}

func (o ProviderRegistrationOutput) ToProviderRegistrationOutput() ProviderRegistrationOutput {
	return o
}

func (o ProviderRegistrationOutput) ToProviderRegistrationOutputWithContext(ctx context.Context) ProviderRegistrationOutput {
	return o
}

func (o ProviderRegistrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderRegistration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProviderRegistrationOutput) Properties() ProviderRegistrationResponsePropertiesOutput {
	return o.ApplyT(func(v *ProviderRegistration) ProviderRegistrationResponsePropertiesOutput { return v.Properties }).(ProviderRegistrationResponsePropertiesOutput)
}

func (o ProviderRegistrationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ProviderRegistration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ProviderRegistrationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderRegistration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProviderRegistrationOutput{})
}
