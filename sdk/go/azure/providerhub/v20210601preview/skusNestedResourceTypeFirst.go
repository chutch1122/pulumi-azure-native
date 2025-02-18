


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SkusNestedResourceTypeFirst struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties SkuResourceResponsePropertiesOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput            `pulumi:"systemData"`
	Type       pulumi.StringOutput                 `pulumi:"type"`
}


func NewSkusNestedResourceTypeFirst(ctx *pulumi.Context,
	name string, args *SkusNestedResourceTypeFirstArgs, opts ...pulumi.ResourceOption) (*SkusNestedResourceTypeFirst, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NestedResourceTypeFirst == nil {
		return nil, errors.New("invalid value for required argument 'NestedResourceTypeFirst'")
	}
	if args.ProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ProviderNamespace'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:providerhub:SkusNestedResourceTypeFirst"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20201120:SkusNestedResourceTypeFirst"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210501preview:SkusNestedResourceTypeFirst"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210901preview:SkusNestedResourceTypeFirst"),
		},
	})
	opts = append(opts, aliases)
	var resource SkusNestedResourceTypeFirst
	err := ctx.RegisterResource("azure-native:providerhub/v20210601preview:SkusNestedResourceTypeFirst", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSkusNestedResourceTypeFirst(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SkusNestedResourceTypeFirstState, opts ...pulumi.ResourceOption) (*SkusNestedResourceTypeFirst, error) {
	var resource SkusNestedResourceTypeFirst
	err := ctx.ReadResource("azure-native:providerhub/v20210601preview:SkusNestedResourceTypeFirst", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type skusNestedResourceTypeFirstState struct {
}

type SkusNestedResourceTypeFirstState struct {
}

func (SkusNestedResourceTypeFirstState) ElementType() reflect.Type {
	return reflect.TypeOf((*skusNestedResourceTypeFirstState)(nil)).Elem()
}

type skusNestedResourceTypeFirstArgs struct {
	NestedResourceTypeFirst string                 `pulumi:"nestedResourceTypeFirst"`
	Properties              *SkuResourceProperties `pulumi:"properties"`
	ProviderNamespace       string                 `pulumi:"providerNamespace"`
	ResourceType            string                 `pulumi:"resourceType"`
	Sku                     *string                `pulumi:"sku"`
}


type SkusNestedResourceTypeFirstArgs struct {
	NestedResourceTypeFirst pulumi.StringInput
	Properties              SkuResourcePropertiesPtrInput
	ProviderNamespace       pulumi.StringInput
	ResourceType            pulumi.StringInput
	Sku                     pulumi.StringPtrInput
}

func (SkusNestedResourceTypeFirstArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*skusNestedResourceTypeFirstArgs)(nil)).Elem()
}

type SkusNestedResourceTypeFirstInput interface {
	pulumi.Input

	ToSkusNestedResourceTypeFirstOutput() SkusNestedResourceTypeFirstOutput
	ToSkusNestedResourceTypeFirstOutputWithContext(ctx context.Context) SkusNestedResourceTypeFirstOutput
}

func (*SkusNestedResourceTypeFirst) ElementType() reflect.Type {
	return reflect.TypeOf((**SkusNestedResourceTypeFirst)(nil)).Elem()
}

func (i *SkusNestedResourceTypeFirst) ToSkusNestedResourceTypeFirstOutput() SkusNestedResourceTypeFirstOutput {
	return i.ToSkusNestedResourceTypeFirstOutputWithContext(context.Background())
}

func (i *SkusNestedResourceTypeFirst) ToSkusNestedResourceTypeFirstOutputWithContext(ctx context.Context) SkusNestedResourceTypeFirstOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkusNestedResourceTypeFirstOutput)
}

type SkusNestedResourceTypeFirstOutput struct{ *pulumi.OutputState }

func (SkusNestedResourceTypeFirstOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkusNestedResourceTypeFirst)(nil)).Elem()
}

func (o SkusNestedResourceTypeFirstOutput) ToSkusNestedResourceTypeFirstOutput() SkusNestedResourceTypeFirstOutput {
	return o
}

func (o SkusNestedResourceTypeFirstOutput) ToSkusNestedResourceTypeFirstOutputWithContext(ctx context.Context) SkusNestedResourceTypeFirstOutput {
	return o
}

func (o SkusNestedResourceTypeFirstOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SkusNestedResourceTypeFirst) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SkusNestedResourceTypeFirstOutput) Properties() SkuResourceResponsePropertiesOutput {
	return o.ApplyT(func(v *SkusNestedResourceTypeFirst) SkuResourceResponsePropertiesOutput { return v.Properties }).(SkuResourceResponsePropertiesOutput)
}

func (o SkusNestedResourceTypeFirstOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SkusNestedResourceTypeFirst) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SkusNestedResourceTypeFirstOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SkusNestedResourceTypeFirst) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SkusNestedResourceTypeFirstOutput{})
}
