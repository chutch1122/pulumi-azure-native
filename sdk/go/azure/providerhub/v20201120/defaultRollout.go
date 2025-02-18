


package v20201120

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DefaultRollout struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                    `pulumi:"name"`
	Properties DefaultRolloutResponsePropertiesOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                    `pulumi:"type"`
}


func NewDefaultRollout(ctx *pulumi.Context,
	name string, args *DefaultRolloutArgs, opts ...pulumi.ResourceOption) (*DefaultRollout, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ProviderNamespace'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:providerhub:DefaultRollout"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210501preview:DefaultRollout"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210601preview:DefaultRollout"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210901preview:DefaultRollout"),
		},
	})
	opts = append(opts, aliases)
	var resource DefaultRollout
	err := ctx.RegisterResource("azure-native:providerhub/v20201120:DefaultRollout", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDefaultRollout(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultRolloutState, opts ...pulumi.ResourceOption) (*DefaultRollout, error) {
	var resource DefaultRollout
	err := ctx.ReadResource("azure-native:providerhub/v20201120:DefaultRollout", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type defaultRolloutState struct {
}

type DefaultRolloutState struct {
}

func (DefaultRolloutState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultRolloutState)(nil)).Elem()
}

type defaultRolloutArgs struct {
	Properties        *DefaultRolloutProperties `pulumi:"properties"`
	ProviderNamespace string                    `pulumi:"providerNamespace"`
	RolloutName       *string                   `pulumi:"rolloutName"`
}


type DefaultRolloutArgs struct {
	Properties        DefaultRolloutPropertiesPtrInput
	ProviderNamespace pulumi.StringInput
	RolloutName       pulumi.StringPtrInput
}

func (DefaultRolloutArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultRolloutArgs)(nil)).Elem()
}

type DefaultRolloutInput interface {
	pulumi.Input

	ToDefaultRolloutOutput() DefaultRolloutOutput
	ToDefaultRolloutOutputWithContext(ctx context.Context) DefaultRolloutOutput
}

func (*DefaultRollout) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultRollout)(nil)).Elem()
}

func (i *DefaultRollout) ToDefaultRolloutOutput() DefaultRolloutOutput {
	return i.ToDefaultRolloutOutputWithContext(context.Background())
}

func (i *DefaultRollout) ToDefaultRolloutOutputWithContext(ctx context.Context) DefaultRolloutOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultRolloutOutput)
}

type DefaultRolloutOutput struct{ *pulumi.OutputState }

func (DefaultRolloutOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultRollout)(nil)).Elem()
}

func (o DefaultRolloutOutput) ToDefaultRolloutOutput() DefaultRolloutOutput {
	return o
}

func (o DefaultRolloutOutput) ToDefaultRolloutOutputWithContext(ctx context.Context) DefaultRolloutOutput {
	return o
}

func (o DefaultRolloutOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultRollout) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DefaultRolloutOutput) Properties() DefaultRolloutResponsePropertiesOutput {
	return o.ApplyT(func(v *DefaultRollout) DefaultRolloutResponsePropertiesOutput { return v.Properties }).(DefaultRolloutResponsePropertiesOutput)
}

func (o DefaultRolloutOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultRollout) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DefaultRolloutOutput{})
}
