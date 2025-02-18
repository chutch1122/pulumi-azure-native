


package blueprint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PolicyAssignmentArtifact struct {
	pulumi.CustomResourceState

	DependsOn          pulumi.StringArrayOutput        `pulumi:"dependsOn"`
	Description        pulumi.StringPtrOutput          `pulumi:"description"`
	DisplayName        pulumi.StringPtrOutput          `pulumi:"displayName"`
	Kind               pulumi.StringOutput             `pulumi:"kind"`
	Name               pulumi.StringOutput             `pulumi:"name"`
	Parameters         ParameterValueResponseMapOutput `pulumi:"parameters"`
	PolicyDefinitionId pulumi.StringOutput             `pulumi:"policyDefinitionId"`
	ResourceGroup      pulumi.StringPtrOutput          `pulumi:"resourceGroup"`
	Type               pulumi.StringOutput             `pulumi:"type"`
}


func NewPolicyAssignmentArtifact(ctx *pulumi.Context,
	name string, args *PolicyAssignmentArtifactArgs, opts ...pulumi.ResourceOption) (*PolicyAssignmentArtifact, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BlueprintName == nil {
		return nil, errors.New("invalid value for required argument 'BlueprintName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.Parameters == nil {
		return nil, errors.New("invalid value for required argument 'Parameters'")
	}
	if args.PolicyDefinitionId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDefinitionId'")
	}
	if args.ResourceScope == nil {
		return nil, errors.New("invalid value for required argument 'ResourceScope'")
	}
	args.Kind = pulumi.String("policyAssignment")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:blueprint/v20181101preview:PolicyAssignmentArtifact"),
		},
	})
	opts = append(opts, aliases)
	var resource PolicyAssignmentArtifact
	err := ctx.RegisterResource("azure-native:blueprint:PolicyAssignmentArtifact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPolicyAssignmentArtifact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyAssignmentArtifactState, opts ...pulumi.ResourceOption) (*PolicyAssignmentArtifact, error) {
	var resource PolicyAssignmentArtifact
	err := ctx.ReadResource("azure-native:blueprint:PolicyAssignmentArtifact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type policyAssignmentArtifactState struct {
}

type PolicyAssignmentArtifactState struct {
}

func (PolicyAssignmentArtifactState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAssignmentArtifactState)(nil)).Elem()
}

type policyAssignmentArtifactArgs struct {
	ArtifactName       *string                   `pulumi:"artifactName"`
	BlueprintName      string                    `pulumi:"blueprintName"`
	DependsOn          []string                  `pulumi:"dependsOn"`
	Description        *string                   `pulumi:"description"`
	DisplayName        *string                   `pulumi:"displayName"`
	Kind               string                    `pulumi:"kind"`
	Parameters         map[string]ParameterValue `pulumi:"parameters"`
	PolicyDefinitionId string                    `pulumi:"policyDefinitionId"`
	ResourceGroup      *string                   `pulumi:"resourceGroup"`
	ResourceScope      string                    `pulumi:"resourceScope"`
}


type PolicyAssignmentArtifactArgs struct {
	ArtifactName       pulumi.StringPtrInput
	BlueprintName      pulumi.StringInput
	DependsOn          pulumi.StringArrayInput
	Description        pulumi.StringPtrInput
	DisplayName        pulumi.StringPtrInput
	Kind               pulumi.StringInput
	Parameters         ParameterValueMapInput
	PolicyDefinitionId pulumi.StringInput
	ResourceGroup      pulumi.StringPtrInput
	ResourceScope      pulumi.StringInput
}

func (PolicyAssignmentArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAssignmentArtifactArgs)(nil)).Elem()
}

type PolicyAssignmentArtifactInput interface {
	pulumi.Input

	ToPolicyAssignmentArtifactOutput() PolicyAssignmentArtifactOutput
	ToPolicyAssignmentArtifactOutputWithContext(ctx context.Context) PolicyAssignmentArtifactOutput
}

func (*PolicyAssignmentArtifact) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignmentArtifact)(nil)).Elem()
}

func (i *PolicyAssignmentArtifact) ToPolicyAssignmentArtifactOutput() PolicyAssignmentArtifactOutput {
	return i.ToPolicyAssignmentArtifactOutputWithContext(context.Background())
}

func (i *PolicyAssignmentArtifact) ToPolicyAssignmentArtifactOutputWithContext(ctx context.Context) PolicyAssignmentArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAssignmentArtifactOutput)
}

type PolicyAssignmentArtifactOutput struct{ *pulumi.OutputState }

func (PolicyAssignmentArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAssignmentArtifact)(nil)).Elem()
}

func (o PolicyAssignmentArtifactOutput) ToPolicyAssignmentArtifactOutput() PolicyAssignmentArtifactOutput {
	return o
}

func (o PolicyAssignmentArtifactOutput) ToPolicyAssignmentArtifactOutputWithContext(ctx context.Context) PolicyAssignmentArtifactOutput {
	return o
}

func (o PolicyAssignmentArtifactOutput) DependsOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PolicyAssignmentArtifact) pulumi.StringArrayOutput { return v.DependsOn }).(pulumi.StringArrayOutput)
}

func (o PolicyAssignmentArtifactOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentArtifact) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentArtifactOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentArtifact) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentArtifactOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAssignmentArtifact) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o PolicyAssignmentArtifactOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAssignmentArtifact) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PolicyAssignmentArtifactOutput) Parameters() ParameterValueResponseMapOutput {
	return o.ApplyT(func(v *PolicyAssignmentArtifact) ParameterValueResponseMapOutput { return v.Parameters }).(ParameterValueResponseMapOutput)
}

func (o PolicyAssignmentArtifactOutput) PolicyDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAssignmentArtifact) pulumi.StringOutput { return v.PolicyDefinitionId }).(pulumi.StringOutput)
}

func (o PolicyAssignmentArtifactOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyAssignmentArtifact) pulumi.StringPtrOutput { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o PolicyAssignmentArtifactOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyAssignmentArtifact) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyAssignmentArtifactOutput{})
}
