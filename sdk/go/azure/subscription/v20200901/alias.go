


package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Alias struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                      `pulumi:"name"`
	Properties PutAliasResponsePropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                      `pulumi:"type"`
}


func NewAlias(ctx *pulumi.Context,
	name string, args *AliasArgs, opts ...pulumi.ResourceOption) (*Alias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:subscription:Alias"),
		},
		{
			Type: pulumi.String("azure-native:subscription/v20191001preview:Alias"),
		},
		{
			Type: pulumi.String("azure-native:subscription/v20211001:Alias"),
		},
	})
	opts = append(opts, aliases)
	var resource Alias
	err := ctx.RegisterResource("azure-native:subscription/v20200901:Alias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AliasState, opts ...pulumi.ResourceOption) (*Alias, error) {
	var resource Alias
	err := ctx.ReadResource("azure-native:subscription/v20200901:Alias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type aliasState struct {
}

type AliasState struct {
}

func (AliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasState)(nil)).Elem()
}

type aliasArgs struct {
	AliasName  *string                   `pulumi:"aliasName"`
	Properties PutAliasRequestProperties `pulumi:"properties"`
}


type AliasArgs struct {
	AliasName  pulumi.StringPtrInput
	Properties PutAliasRequestPropertiesInput
}

func (AliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasArgs)(nil)).Elem()
}

type AliasInput interface {
	pulumi.Input

	ToAliasOutput() AliasOutput
	ToAliasOutputWithContext(ctx context.Context) AliasOutput
}

func (*Alias) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil)).Elem()
}

func (i *Alias) ToAliasOutput() AliasOutput {
	return i.ToAliasOutputWithContext(context.Background())
}

func (i *Alias) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasOutput)
}

type AliasOutput struct{ *pulumi.OutputState }

func (AliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil)).Elem()
}

func (o AliasOutput) ToAliasOutput() AliasOutput {
	return o
}

func (o AliasOutput) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return o
}

func (o AliasOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AliasOutput) Properties() PutAliasResponsePropertiesResponseOutput {
	return o.ApplyT(func(v *Alias) PutAliasResponsePropertiesResponseOutput { return v.Properties }).(PutAliasResponsePropertiesResponseOutput)
}

func (o AliasOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AliasOutput{})
}
