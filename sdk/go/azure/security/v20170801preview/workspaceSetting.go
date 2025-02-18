


package v20170801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkspaceSetting struct {
	pulumi.CustomResourceState

	Name        pulumi.StringOutput `pulumi:"name"`
	Scope       pulumi.StringOutput `pulumi:"scope"`
	Type        pulumi.StringOutput `pulumi:"type"`
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}


func NewWorkspaceSetting(ctx *pulumi.Context,
	name string, args *WorkspaceSettingArgs, opts ...pulumi.ResourceOption) (*WorkspaceSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security:WorkspaceSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkspaceSetting
	err := ctx.RegisterResource("azure-native:security/v20170801preview:WorkspaceSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspaceSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceSettingState, opts ...pulumi.ResourceOption) (*WorkspaceSetting, error) {
	var resource WorkspaceSetting
	err := ctx.ReadResource("azure-native:security/v20170801preview:WorkspaceSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workspaceSettingState struct {
}

type WorkspaceSettingState struct {
}

func (WorkspaceSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceSettingState)(nil)).Elem()
}

type workspaceSettingArgs struct {
	Scope                string  `pulumi:"scope"`
	WorkspaceId          string  `pulumi:"workspaceId"`
	WorkspaceSettingName *string `pulumi:"workspaceSettingName"`
}


type WorkspaceSettingArgs struct {
	Scope                pulumi.StringInput
	WorkspaceId          pulumi.StringInput
	WorkspaceSettingName pulumi.StringPtrInput
}

func (WorkspaceSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceSettingArgs)(nil)).Elem()
}

type WorkspaceSettingInput interface {
	pulumi.Input

	ToWorkspaceSettingOutput() WorkspaceSettingOutput
	ToWorkspaceSettingOutputWithContext(ctx context.Context) WorkspaceSettingOutput
}

func (*WorkspaceSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceSetting)(nil)).Elem()
}

func (i *WorkspaceSetting) ToWorkspaceSettingOutput() WorkspaceSettingOutput {
	return i.ToWorkspaceSettingOutputWithContext(context.Background())
}

func (i *WorkspaceSetting) ToWorkspaceSettingOutputWithContext(ctx context.Context) WorkspaceSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSettingOutput)
}

type WorkspaceSettingOutput struct{ *pulumi.OutputState }

func (WorkspaceSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceSetting)(nil)).Elem()
}

func (o WorkspaceSettingOutput) ToWorkspaceSettingOutput() WorkspaceSettingOutput {
	return o
}

func (o WorkspaceSettingOutput) ToWorkspaceSettingOutputWithContext(ctx context.Context) WorkspaceSettingOutput {
	return o
}

func (o WorkspaceSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkspaceSettingOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceSetting) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

func (o WorkspaceSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o WorkspaceSettingOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceSetting) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceSettingOutput{})
}
