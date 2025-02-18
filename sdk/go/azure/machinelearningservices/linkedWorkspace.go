


package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LinkedWorkspace struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                `pulumi:"name"`
	Properties LinkedWorkspacePropsResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                `pulumi:"type"`
}


func NewLinkedWorkspace(ctx *pulumi.Context,
	name string, args *LinkedWorkspaceArgs, opts ...pulumi.ResourceOption) (*LinkedWorkspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200301:LinkedWorkspace"),
		},
	})
	opts = append(opts, aliases)
	var resource LinkedWorkspace
	err := ctx.RegisterResource("azure-native:machinelearningservices:LinkedWorkspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLinkedWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedWorkspaceState, opts ...pulumi.ResourceOption) (*LinkedWorkspace, error) {
	var resource LinkedWorkspace
	err := ctx.ReadResource("azure-native:machinelearningservices:LinkedWorkspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type linkedWorkspaceState struct {
}

type LinkedWorkspaceState struct {
}

func (LinkedWorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedWorkspaceState)(nil)).Elem()
}

type linkedWorkspaceArgs struct {
	LinkName          *string               `pulumi:"linkName"`
	Name              *string               `pulumi:"name"`
	Properties        *LinkedWorkspaceProps `pulumi:"properties"`
	ResourceGroupName string                `pulumi:"resourceGroupName"`
	WorkspaceName     string                `pulumi:"workspaceName"`
}


type LinkedWorkspaceArgs struct {
	LinkName          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	Properties        LinkedWorkspacePropsPtrInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (LinkedWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedWorkspaceArgs)(nil)).Elem()
}

type LinkedWorkspaceInput interface {
	pulumi.Input

	ToLinkedWorkspaceOutput() LinkedWorkspaceOutput
	ToLinkedWorkspaceOutputWithContext(ctx context.Context) LinkedWorkspaceOutput
}

func (*LinkedWorkspace) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedWorkspace)(nil)).Elem()
}

func (i *LinkedWorkspace) ToLinkedWorkspaceOutput() LinkedWorkspaceOutput {
	return i.ToLinkedWorkspaceOutputWithContext(context.Background())
}

func (i *LinkedWorkspace) ToLinkedWorkspaceOutputWithContext(ctx context.Context) LinkedWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedWorkspaceOutput)
}

type LinkedWorkspaceOutput struct{ *pulumi.OutputState }

func (LinkedWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedWorkspace)(nil)).Elem()
}

func (o LinkedWorkspaceOutput) ToLinkedWorkspaceOutput() LinkedWorkspaceOutput {
	return o
}

func (o LinkedWorkspaceOutput) ToLinkedWorkspaceOutputWithContext(ctx context.Context) LinkedWorkspaceOutput {
	return o
}

func (o LinkedWorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedWorkspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LinkedWorkspaceOutput) Properties() LinkedWorkspacePropsResponseOutput {
	return o.ApplyT(func(v *LinkedWorkspace) LinkedWorkspacePropsResponseOutput { return v.Properties }).(LinkedWorkspacePropsResponseOutput)
}

func (o LinkedWorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedWorkspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedWorkspaceOutput{})
}
