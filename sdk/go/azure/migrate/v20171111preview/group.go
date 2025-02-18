


package v20171111preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Group struct {
	pulumi.CustomResourceState

	Assessments      pulumi.StringArrayOutput `pulumi:"assessments"`
	CreatedTimestamp pulumi.StringOutput      `pulumi:"createdTimestamp"`
	ETag             pulumi.StringPtrOutput   `pulumi:"eTag"`
	Machines         pulumi.StringArrayOutput `pulumi:"machines"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	Type             pulumi.StringOutput      `pulumi:"type"`
	UpdatedTimestamp pulumi.StringOutput      `pulumi:"updatedTimestamp"`
}


func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Machines == nil {
		return nil, errors.New("invalid value for required argument 'Machines'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20180202:Group"),
		},
	})
	opts = append(opts, aliases)
	var resource Group
	err := ctx.RegisterResource("azure-native:migrate/v20171111preview:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("azure-native:migrate/v20171111preview:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type groupState struct {
}

type GroupState struct {
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	ETag              *string  `pulumi:"eTag"`
	GroupName         *string  `pulumi:"groupName"`
	Machines          []string `pulumi:"machines"`
	ProjectName       string   `pulumi:"projectName"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
}


type GroupArgs struct {
	ETag              pulumi.StringPtrInput
	GroupName         pulumi.StringPtrInput
	Machines          pulumi.StringArrayInput
	ProjectName       pulumi.StringInput
	ResourceGroupName pulumi.StringInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func (o GroupOutput) Assessments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Group) pulumi.StringArrayOutput { return v.Assessments }).(pulumi.StringArrayOutput)
}

func (o GroupOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

func (o GroupOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o GroupOutput) Machines() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Group) pulumi.StringArrayOutput { return v.Machines }).(pulumi.StringArrayOutput)
}

func (o GroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o GroupOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GroupOutput{})
}
