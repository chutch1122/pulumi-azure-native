


package v20180501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PolicySetDefinition struct {
	pulumi.CustomResourceState

	Description       pulumi.StringPtrOutput                       `pulumi:"description"`
	DisplayName       pulumi.StringPtrOutput                       `pulumi:"displayName"`
	Metadata          pulumi.AnyOutput                             `pulumi:"metadata"`
	Name              pulumi.StringOutput                          `pulumi:"name"`
	Parameters        pulumi.AnyOutput                             `pulumi:"parameters"`
	PolicyDefinitions PolicyDefinitionReferenceResponseArrayOutput `pulumi:"policyDefinitions"`
	PolicyType        pulumi.StringPtrOutput                       `pulumi:"policyType"`
	Type              pulumi.StringOutput                          `pulumi:"type"`
}


func NewPolicySetDefinition(ctx *pulumi.Context,
	name string, args *PolicySetDefinitionArgs, opts ...pulumi.ResourceOption) (*PolicySetDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyDefinitions == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDefinitions'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization:PolicySetDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20170601preview:PolicySetDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180301:PolicySetDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190101:PolicySetDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190601:PolicySetDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20190901:PolicySetDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200301:PolicySetDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200901:PolicySetDefinition"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20210601:PolicySetDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource PolicySetDefinition
	err := ctx.RegisterResource("azure-native:authorization/v20180501:PolicySetDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPolicySetDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicySetDefinitionState, opts ...pulumi.ResourceOption) (*PolicySetDefinition, error) {
	var resource PolicySetDefinition
	err := ctx.ReadResource("azure-native:authorization/v20180501:PolicySetDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type policySetDefinitionState struct {
}

type PolicySetDefinitionState struct {
}

func (PolicySetDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*policySetDefinitionState)(nil)).Elem()
}

type policySetDefinitionArgs struct {
	Description             *string                     `pulumi:"description"`
	DisplayName             *string                     `pulumi:"displayName"`
	Metadata                interface{}                 `pulumi:"metadata"`
	Parameters              interface{}                 `pulumi:"parameters"`
	PolicyDefinitions       []PolicyDefinitionReference `pulumi:"policyDefinitions"`
	PolicySetDefinitionName *string                     `pulumi:"policySetDefinitionName"`
	PolicyType              *string                     `pulumi:"policyType"`
}


type PolicySetDefinitionArgs struct {
	Description             pulumi.StringPtrInput
	DisplayName             pulumi.StringPtrInput
	Metadata                pulumi.Input
	Parameters              pulumi.Input
	PolicyDefinitions       PolicyDefinitionReferenceArrayInput
	PolicySetDefinitionName pulumi.StringPtrInput
	PolicyType              pulumi.StringPtrInput
}

func (PolicySetDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policySetDefinitionArgs)(nil)).Elem()
}

type PolicySetDefinitionInput interface {
	pulumi.Input

	ToPolicySetDefinitionOutput() PolicySetDefinitionOutput
	ToPolicySetDefinitionOutputWithContext(ctx context.Context) PolicySetDefinitionOutput
}

func (*PolicySetDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySetDefinition)(nil)).Elem()
}

func (i *PolicySetDefinition) ToPolicySetDefinitionOutput() PolicySetDefinitionOutput {
	return i.ToPolicySetDefinitionOutputWithContext(context.Background())
}

func (i *PolicySetDefinition) ToPolicySetDefinitionOutputWithContext(ctx context.Context) PolicySetDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySetDefinitionOutput)
}

type PolicySetDefinitionOutput struct{ *pulumi.OutputState }

func (PolicySetDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySetDefinition)(nil)).Elem()
}

func (o PolicySetDefinitionOutput) ToPolicySetDefinitionOutput() PolicySetDefinitionOutput {
	return o
}

func (o PolicySetDefinitionOutput) ToPolicySetDefinitionOutputWithContext(ctx context.Context) PolicySetDefinitionOutput {
	return o
}

func (o PolicySetDefinitionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySetDefinition) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PolicySetDefinitionOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySetDefinition) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PolicySetDefinitionOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *PolicySetDefinition) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

func (o PolicySetDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicySetDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PolicySetDefinitionOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *PolicySetDefinition) pulumi.AnyOutput { return v.Parameters }).(pulumi.AnyOutput)
}

func (o PolicySetDefinitionOutput) PolicyDefinitions() PolicyDefinitionReferenceResponseArrayOutput {
	return o.ApplyT(func(v *PolicySetDefinition) PolicyDefinitionReferenceResponseArrayOutput { return v.PolicyDefinitions }).(PolicyDefinitionReferenceResponseArrayOutput)
}

func (o PolicySetDefinitionOutput) PolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySetDefinition) pulumi.StringPtrOutput { return v.PolicyType }).(pulumi.StringPtrOutput)
}

func (o PolicySetDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicySetDefinition) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicySetDefinitionOutput{})
}
