


package devtestlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Policy struct {
	pulumi.CustomResourceState

	CreatedDate       pulumi.StringOutput    `pulumi:"createdDate"`
	Description       pulumi.StringPtrOutput `pulumi:"description"`
	EvaluatorType     pulumi.StringPtrOutput `pulumi:"evaluatorType"`
	FactData          pulumi.StringPtrOutput `pulumi:"factData"`
	FactName          pulumi.StringPtrOutput `pulumi:"factName"`
	Location          pulumi.StringPtrOutput `pulumi:"location"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	Status            pulumi.StringPtrOutput `pulumi:"status"`
	Tags              pulumi.StringMapOutput `pulumi:"tags"`
	Threshold         pulumi.StringPtrOutput `pulumi:"threshold"`
	Type              pulumi.StringOutput    `pulumi:"type"`
	UniqueIdentifier  pulumi.StringOutput    `pulumi:"uniqueIdentifier"`
}


func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.PolicySetName == nil {
		return nil, errors.New("invalid value for required argument 'PolicySetName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:Policy"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:Policy"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:Policy"),
		},
	})
	opts = append(opts, aliases)
	var resource Policy
	err := ctx.RegisterResource("azure-native:devtestlab:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("azure-native:devtestlab:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type policyState struct {
}

type PolicyState struct {
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	Description       *string           `pulumi:"description"`
	EvaluatorType     *string           `pulumi:"evaluatorType"`
	FactData          *string           `pulumi:"factData"`
	FactName          *string           `pulumi:"factName"`
	LabName           string            `pulumi:"labName"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	PolicySetName     string            `pulumi:"policySetName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Status            *string           `pulumi:"status"`
	Tags              map[string]string `pulumi:"tags"`
	Threshold         *string           `pulumi:"threshold"`
}


type PolicyArgs struct {
	Description       pulumi.StringPtrInput
	EvaluatorType     pulumi.StringPtrInput
	FactData          pulumi.StringPtrInput
	FactName          pulumi.StringPtrInput
	LabName           pulumi.StringInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	PolicySetName     pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Status            pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	Threshold         pulumi.StringPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

type PolicyOutput struct{ *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

func (o PolicyOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o PolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PolicyOutput) EvaluatorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.EvaluatorType }).(pulumi.StringPtrOutput)
}

func (o PolicyOutput) FactData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.FactData }).(pulumi.StringPtrOutput)
}

func (o PolicyOutput) FactName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.FactName }).(pulumi.StringPtrOutput)
}

func (o PolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PolicyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o PolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PolicyOutput) Threshold() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.Threshold }).(pulumi.StringPtrOutput)
}

func (o PolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o PolicyOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyOutput{})
}
