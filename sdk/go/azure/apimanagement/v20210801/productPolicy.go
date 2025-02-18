


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProductPolicy struct {
	pulumi.CustomResourceState

	Format pulumi.StringPtrOutput `pulumi:"format"`
	Name   pulumi.StringOutput    `pulumi:"name"`
	Type   pulumi.StringOutput    `pulumi:"type"`
	Value  pulumi.StringOutput    `pulumi:"value"`
}


func NewProductPolicy(ctx *pulumi.Context,
	name string, args *ProductPolicyArgs, opts ...pulumi.ResourceOption) (*ProductPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	if isZero(args.Format) {
		args.Format = pulumi.StringPtr("xml")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ProductPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:ProductPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ProductPolicy
	err := ctx.RegisterResource("azure-native:apimanagement/v20210801:ProductPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProductPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductPolicyState, opts ...pulumi.ResourceOption) (*ProductPolicy, error) {
	var resource ProductPolicy
	err := ctx.ReadResource("azure-native:apimanagement/v20210801:ProductPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type productPolicyState struct {
}

type ProductPolicyState struct {
}

func (ProductPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*productPolicyState)(nil)).Elem()
}

type productPolicyArgs struct {
	Format            *string `pulumi:"format"`
	PolicyId          *string `pulumi:"policyId"`
	ProductId         string  `pulumi:"productId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
	Value             string  `pulumi:"value"`
}


type ProductPolicyArgs struct {
	Format            pulumi.StringPtrInput
	PolicyId          pulumi.StringPtrInput
	ProductId         pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	Value             pulumi.StringInput
}

func (ProductPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productPolicyArgs)(nil)).Elem()
}

type ProductPolicyInput interface {
	pulumi.Input

	ToProductPolicyOutput() ProductPolicyOutput
	ToProductPolicyOutputWithContext(ctx context.Context) ProductPolicyOutput
}

func (*ProductPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductPolicy)(nil)).Elem()
}

func (i *ProductPolicy) ToProductPolicyOutput() ProductPolicyOutput {
	return i.ToProductPolicyOutputWithContext(context.Background())
}

func (i *ProductPolicy) ToProductPolicyOutputWithContext(ctx context.Context) ProductPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductPolicyOutput)
}

type ProductPolicyOutput struct{ *pulumi.OutputState }

func (ProductPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductPolicy)(nil)).Elem()
}

func (o ProductPolicyOutput) ToProductPolicyOutput() ProductPolicyOutput {
	return o
}

func (o ProductPolicyOutput) ToProductPolicyOutputWithContext(ctx context.Context) ProductPolicyOutput {
	return o
}

func (o ProductPolicyOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductPolicy) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

func (o ProductPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProductPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ProductPolicyOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductPolicy) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProductPolicyOutput{})
}
