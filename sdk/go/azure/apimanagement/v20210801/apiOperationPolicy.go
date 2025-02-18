


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiOperationPolicy struct {
	pulumi.CustomResourceState

	Format pulumi.StringPtrOutput `pulumi:"format"`
	Name   pulumi.StringOutput    `pulumi:"name"`
	Type   pulumi.StringOutput    `pulumi:"type"`
	Value  pulumi.StringOutput    `pulumi:"value"`
}


func NewApiOperationPolicy(ctx *pulumi.Context,
	name string, args *ApiOperationPolicyArgs, opts ...pulumi.ResourceOption) (*ApiOperationPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.OperationId == nil {
		return nil, errors.New("invalid value for required argument 'OperationId'")
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
			Type: pulumi.String("azure-native:apimanagement:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:ApiOperationPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiOperationPolicy
	err := ctx.RegisterResource("azure-native:apimanagement/v20210801:ApiOperationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiOperationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiOperationPolicyState, opts ...pulumi.ResourceOption) (*ApiOperationPolicy, error) {
	var resource ApiOperationPolicy
	err := ctx.ReadResource("azure-native:apimanagement/v20210801:ApiOperationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type apiOperationPolicyState struct {
}

type ApiOperationPolicyState struct {
}

func (ApiOperationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationPolicyState)(nil)).Elem()
}

type apiOperationPolicyArgs struct {
	ApiId             string  `pulumi:"apiId"`
	Format            *string `pulumi:"format"`
	OperationId       string  `pulumi:"operationId"`
	PolicyId          *string `pulumi:"policyId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
	Value             string  `pulumi:"value"`
}


type ApiOperationPolicyArgs struct {
	ApiId             pulumi.StringInput
	Format            pulumi.StringPtrInput
	OperationId       pulumi.StringInput
	PolicyId          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	Value             pulumi.StringInput
}

func (ApiOperationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationPolicyArgs)(nil)).Elem()
}

type ApiOperationPolicyInput interface {
	pulumi.Input

	ToApiOperationPolicyOutput() ApiOperationPolicyOutput
	ToApiOperationPolicyOutputWithContext(ctx context.Context) ApiOperationPolicyOutput
}

func (*ApiOperationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOperationPolicy)(nil)).Elem()
}

func (i *ApiOperationPolicy) ToApiOperationPolicyOutput() ApiOperationPolicyOutput {
	return i.ToApiOperationPolicyOutputWithContext(context.Background())
}

func (i *ApiOperationPolicy) ToApiOperationPolicyOutputWithContext(ctx context.Context) ApiOperationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOperationPolicyOutput)
}

type ApiOperationPolicyOutput struct{ *pulumi.OutputState }

func (ApiOperationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOperationPolicy)(nil)).Elem()
}

func (o ApiOperationPolicyOutput) ToApiOperationPolicyOutput() ApiOperationPolicyOutput {
	return o
}

func (o ApiOperationPolicyOutput) ToApiOperationPolicyOutputWithContext(ctx context.Context) ApiOperationPolicyOutput {
	return o
}

func (o ApiOperationPolicyOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiOperationPolicy) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

func (o ApiOperationPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiOperationPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApiOperationPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiOperationPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ApiOperationPolicyOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiOperationPolicy) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiOperationPolicyOutput{})
}
