


package migrate

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Group struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringPtrOutput        `pulumi:"eTag"`
	Name       pulumi.StringOutput           `pulumi:"name"`
	Properties GroupPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput           `pulumi:"type"`
}


func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20191001:Group"),
		},
	})
	opts = append(opts, aliases)
	var resource Group
	err := ctx.RegisterResource("azure-native:migrate:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("azure-native:migrate:Group", name, id, state, &resource, opts...)
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
	ETag              *string         `pulumi:"eTag"`
	GroupName         *string         `pulumi:"groupName"`
	ProjectName       string          `pulumi:"projectName"`
	Properties        GroupProperties `pulumi:"properties"`
	ResourceGroupName string          `pulumi:"resourceGroupName"`
}


type GroupArgs struct {
	ETag              pulumi.StringPtrInput
	GroupName         pulumi.StringPtrInput
	ProjectName       pulumi.StringInput
	Properties        GroupPropertiesInput
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

func (o GroupOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o GroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GroupOutput) Properties() GroupPropertiesResponseOutput {
	return o.ApplyT(func(v *Group) GroupPropertiesResponseOutput { return v.Properties }).(GroupPropertiesResponseOutput)
}

func (o GroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GroupOutput{})
}
