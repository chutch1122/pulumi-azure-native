


package authorization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PolicyExemption struct {
	pulumi.CustomResourceState

	Description                  pulumi.StringPtrOutput   `pulumi:"description"`
	DisplayName                  pulumi.StringPtrOutput   `pulumi:"displayName"`
	ExemptionCategory            pulumi.StringOutput      `pulumi:"exemptionCategory"`
	ExpiresOn                    pulumi.StringPtrOutput   `pulumi:"expiresOn"`
	Metadata                     pulumi.AnyOutput         `pulumi:"metadata"`
	Name                         pulumi.StringOutput      `pulumi:"name"`
	PolicyAssignmentId           pulumi.StringOutput      `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceIds pulumi.StringArrayOutput `pulumi:"policyDefinitionReferenceIds"`
	SystemData                   SystemDataResponseOutput `pulumi:"systemData"`
	Type                         pulumi.StringOutput      `pulumi:"type"`
}


func NewPolicyExemption(ctx *pulumi.Context,
	name string, args *PolicyExemptionArgs, opts ...pulumi.ResourceOption) (*PolicyExemption, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExemptionCategory == nil {
		return nil, errors.New("invalid value for required argument 'ExemptionCategory'")
	}
	if args.PolicyAssignmentId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyAssignmentId'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization/v20200701preview:PolicyExemption"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20220701preview:PolicyExemption"),
		},
	})
	opts = append(opts, aliases)
	var resource PolicyExemption
	err := ctx.RegisterResource("azure-native:authorization:PolicyExemption", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPolicyExemption(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyExemptionState, opts ...pulumi.ResourceOption) (*PolicyExemption, error) {
	var resource PolicyExemption
	err := ctx.ReadResource("azure-native:authorization:PolicyExemption", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type policyExemptionState struct {
}

type PolicyExemptionState struct {
}

func (PolicyExemptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyExemptionState)(nil)).Elem()
}

type policyExemptionArgs struct {
	Description                  *string     `pulumi:"description"`
	DisplayName                  *string     `pulumi:"displayName"`
	ExemptionCategory            string      `pulumi:"exemptionCategory"`
	ExpiresOn                    *string     `pulumi:"expiresOn"`
	Metadata                     interface{} `pulumi:"metadata"`
	PolicyAssignmentId           string      `pulumi:"policyAssignmentId"`
	PolicyDefinitionReferenceIds []string    `pulumi:"policyDefinitionReferenceIds"`
	PolicyExemptionName          *string     `pulumi:"policyExemptionName"`
	Scope                        string      `pulumi:"scope"`
}


type PolicyExemptionArgs struct {
	Description                  pulumi.StringPtrInput
	DisplayName                  pulumi.StringPtrInput
	ExemptionCategory            pulumi.StringInput
	ExpiresOn                    pulumi.StringPtrInput
	Metadata                     pulumi.Input
	PolicyAssignmentId           pulumi.StringInput
	PolicyDefinitionReferenceIds pulumi.StringArrayInput
	PolicyExemptionName          pulumi.StringPtrInput
	Scope                        pulumi.StringInput
}

func (PolicyExemptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyExemptionArgs)(nil)).Elem()
}

type PolicyExemptionInput interface {
	pulumi.Input

	ToPolicyExemptionOutput() PolicyExemptionOutput
	ToPolicyExemptionOutputWithContext(ctx context.Context) PolicyExemptionOutput
}

func (*PolicyExemption) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyExemption)(nil)).Elem()
}

func (i *PolicyExemption) ToPolicyExemptionOutput() PolicyExemptionOutput {
	return i.ToPolicyExemptionOutputWithContext(context.Background())
}

func (i *PolicyExemption) ToPolicyExemptionOutputWithContext(ctx context.Context) PolicyExemptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyExemptionOutput)
}

type PolicyExemptionOutput struct{ *pulumi.OutputState }

func (PolicyExemptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyExemption)(nil)).Elem()
}

func (o PolicyExemptionOutput) ToPolicyExemptionOutput() PolicyExemptionOutput {
	return o
}

func (o PolicyExemptionOutput) ToPolicyExemptionOutputWithContext(ctx context.Context) PolicyExemptionOutput {
	return o
}

func (o PolicyExemptionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyExemption) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PolicyExemptionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyExemption) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PolicyExemptionOutput) ExemptionCategory() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyExemption) pulumi.StringOutput { return v.ExemptionCategory }).(pulumi.StringOutput)
}

func (o PolicyExemptionOutput) ExpiresOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyExemption) pulumi.StringPtrOutput { return v.ExpiresOn }).(pulumi.StringPtrOutput)
}

func (o PolicyExemptionOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *PolicyExemption) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

func (o PolicyExemptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyExemption) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PolicyExemptionOutput) PolicyAssignmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyExemption) pulumi.StringOutput { return v.PolicyAssignmentId }).(pulumi.StringOutput)
}

func (o PolicyExemptionOutput) PolicyDefinitionReferenceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PolicyExemption) pulumi.StringArrayOutput { return v.PolicyDefinitionReferenceIds }).(pulumi.StringArrayOutput)
}

func (o PolicyExemptionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PolicyExemption) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PolicyExemptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyExemption) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyExemptionOutput{})
}
