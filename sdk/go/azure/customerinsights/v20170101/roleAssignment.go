


package v20170101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type RoleAssignment struct {
	pulumi.CustomResourceState

	AssignmentName     pulumi.StringOutput                     `pulumi:"assignmentName"`
	ConflationPolicies ResourceSetDescriptionResponsePtrOutput `pulumi:"conflationPolicies"`
	Connectors         ResourceSetDescriptionResponsePtrOutput `pulumi:"connectors"`
	Description        pulumi.StringMapOutput                  `pulumi:"description"`
	DisplayName        pulumi.StringMapOutput                  `pulumi:"displayName"`
	Interactions       ResourceSetDescriptionResponsePtrOutput `pulumi:"interactions"`
	Kpis               ResourceSetDescriptionResponsePtrOutput `pulumi:"kpis"`
	Links              ResourceSetDescriptionResponsePtrOutput `pulumi:"links"`
	Name               pulumi.StringOutput                     `pulumi:"name"`
	Principals         AssignmentPrincipalResponseArrayOutput  `pulumi:"principals"`
	Profiles           ResourceSetDescriptionResponsePtrOutput `pulumi:"profiles"`
	ProvisioningState  pulumi.StringOutput                     `pulumi:"provisioningState"`
	RelationshipLinks  ResourceSetDescriptionResponsePtrOutput `pulumi:"relationshipLinks"`
	Relationships      ResourceSetDescriptionResponsePtrOutput `pulumi:"relationships"`
	Role               pulumi.StringOutput                     `pulumi:"role"`
	RoleAssignments    ResourceSetDescriptionResponsePtrOutput `pulumi:"roleAssignments"`
	SasPolicies        ResourceSetDescriptionResponsePtrOutput `pulumi:"sasPolicies"`
	Segments           ResourceSetDescriptionResponsePtrOutput `pulumi:"segments"`
	TenantId           pulumi.StringOutput                     `pulumi:"tenantId"`
	Type               pulumi.StringOutput                     `pulumi:"type"`
	Views              ResourceSetDescriptionResponsePtrOutput `pulumi:"views"`
	WidgetTypes        ResourceSetDescriptionResponsePtrOutput `pulumi:"widgetTypes"`
}


func NewRoleAssignment(ctx *pulumi.Context,
	name string, args *RoleAssignmentArgs, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HubName == nil {
		return nil, errors.New("invalid value for required argument 'HubName'")
	}
	if args.Principals == nil {
		return nil, errors.New("invalid value for required argument 'Principals'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:customerinsights:RoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights/v20170426:RoleAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource RoleAssignment
	err := ctx.RegisterResource("azure-native:customerinsights/v20170101:RoleAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRoleAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleAssignmentState, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	var resource RoleAssignment
	err := ctx.ReadResource("azure-native:customerinsights/v20170101:RoleAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type roleAssignmentState struct {
}

type RoleAssignmentState struct {
}

func (RoleAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentState)(nil)).Elem()
}

type roleAssignmentArgs struct {
	AssignmentName     *string                 `pulumi:"assignmentName"`
	ConflationPolicies *ResourceSetDescription `pulumi:"conflationPolicies"`
	Connectors         *ResourceSetDescription `pulumi:"connectors"`
	Description        map[string]string       `pulumi:"description"`
	DisplayName        map[string]string       `pulumi:"displayName"`
	HubName            string                  `pulumi:"hubName"`
	Interactions       *ResourceSetDescription `pulumi:"interactions"`
	Kpis               *ResourceSetDescription `pulumi:"kpis"`
	Links              *ResourceSetDescription `pulumi:"links"`
	Principals         []AssignmentPrincipal   `pulumi:"principals"`
	Profiles           *ResourceSetDescription `pulumi:"profiles"`
	RelationshipLinks  *ResourceSetDescription `pulumi:"relationshipLinks"`
	Relationships      *ResourceSetDescription `pulumi:"relationships"`
	ResourceGroupName  string                  `pulumi:"resourceGroupName"`
	Role               RoleTypes               `pulumi:"role"`
	RoleAssignments    *ResourceSetDescription `pulumi:"roleAssignments"`
	SasPolicies        *ResourceSetDescription `pulumi:"sasPolicies"`
	Segments           *ResourceSetDescription `pulumi:"segments"`
	Views              *ResourceSetDescription `pulumi:"views"`
	WidgetTypes        *ResourceSetDescription `pulumi:"widgetTypes"`
}


type RoleAssignmentArgs struct {
	AssignmentName     pulumi.StringPtrInput
	ConflationPolicies ResourceSetDescriptionPtrInput
	Connectors         ResourceSetDescriptionPtrInput
	Description        pulumi.StringMapInput
	DisplayName        pulumi.StringMapInput
	HubName            pulumi.StringInput
	Interactions       ResourceSetDescriptionPtrInput
	Kpis               ResourceSetDescriptionPtrInput
	Links              ResourceSetDescriptionPtrInput
	Principals         AssignmentPrincipalArrayInput
	Profiles           ResourceSetDescriptionPtrInput
	RelationshipLinks  ResourceSetDescriptionPtrInput
	Relationships      ResourceSetDescriptionPtrInput
	ResourceGroupName  pulumi.StringInput
	Role               RoleTypesInput
	RoleAssignments    ResourceSetDescriptionPtrInput
	SasPolicies        ResourceSetDescriptionPtrInput
	Segments           ResourceSetDescriptionPtrInput
	Views              ResourceSetDescriptionPtrInput
	WidgetTypes        ResourceSetDescriptionPtrInput
}

func (RoleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentArgs)(nil)).Elem()
}

type RoleAssignmentInput interface {
	pulumi.Input

	ToRoleAssignmentOutput() RoleAssignmentOutput
	ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput
}

func (*RoleAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleAssignment)(nil)).Elem()
}

func (i *RoleAssignment) ToRoleAssignmentOutput() RoleAssignmentOutput {
	return i.ToRoleAssignmentOutputWithContext(context.Background())
}

func (i *RoleAssignment) ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentOutput)
}

type RoleAssignmentOutput struct{ *pulumi.OutputState }

func (RoleAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleAssignment)(nil)).Elem()
}

func (o RoleAssignmentOutput) ToRoleAssignmentOutput() RoleAssignmentOutput {
	return o
}

func (o RoleAssignmentOutput) ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput {
	return o
}

func (o RoleAssignmentOutput) AssignmentName() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.AssignmentName }).(pulumi.StringOutput)
}

func (o RoleAssignmentOutput) ConflationPolicies() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.ConflationPolicies }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o RoleAssignmentOutput) Connectors() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.Connectors }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o RoleAssignmentOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringMapOutput { return v.Description }).(pulumi.StringMapOutput)
}

func (o RoleAssignmentOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringMapOutput { return v.DisplayName }).(pulumi.StringMapOutput)
}

func (o RoleAssignmentOutput) Interactions() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.Interactions }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o RoleAssignmentOutput) Kpis() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.Kpis }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o RoleAssignmentOutput) Links() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.Links }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o RoleAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RoleAssignmentOutput) Principals() AssignmentPrincipalResponseArrayOutput {
	return o.ApplyT(func(v *RoleAssignment) AssignmentPrincipalResponseArrayOutput { return v.Principals }).(AssignmentPrincipalResponseArrayOutput)
}

func (o RoleAssignmentOutput) Profiles() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.Profiles }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o RoleAssignmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RoleAssignmentOutput) RelationshipLinks() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.RelationshipLinks }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o RoleAssignmentOutput) Relationships() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.Relationships }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o RoleAssignmentOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o RoleAssignmentOutput) RoleAssignments() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.RoleAssignments }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o RoleAssignmentOutput) SasPolicies() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.SasPolicies }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o RoleAssignmentOutput) Segments() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.Segments }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o RoleAssignmentOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o RoleAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o RoleAssignmentOutput) Views() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.Views }).(ResourceSetDescriptionResponsePtrOutput)
}

func (o RoleAssignmentOutput) WidgetTypes() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.WidgetTypes }).(ResourceSetDescriptionResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RoleAssignmentOutput{})
}
