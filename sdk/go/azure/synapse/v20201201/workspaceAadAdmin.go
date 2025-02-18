


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type WorkspaceAadAdmin struct {
	pulumi.CustomResourceState

	AdministratorType pulumi.StringPtrOutput `pulumi:"administratorType"`
	Login             pulumi.StringPtrOutput `pulumi:"login"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	Sid               pulumi.StringPtrOutput `pulumi:"sid"`
	TenantId          pulumi.StringPtrOutput `pulumi:"tenantId"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewWorkspaceAadAdmin(ctx *pulumi.Context,
	name string, args *WorkspaceAadAdminArgs, opts ...pulumi.ResourceOption) (*WorkspaceAadAdmin, error) {
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
			Type: pulumi.String("azure-native:synapse:WorkspaceAadAdmin"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:WorkspaceAadAdmin"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:WorkspaceAadAdmin"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:WorkspaceAadAdmin"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:WorkspaceAadAdmin"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:WorkspaceAadAdmin"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:WorkspaceAadAdmin"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkspaceAadAdmin
	err := ctx.RegisterResource("azure-native:synapse/v20201201:WorkspaceAadAdmin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspaceAadAdmin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceAadAdminState, opts ...pulumi.ResourceOption) (*WorkspaceAadAdmin, error) {
	var resource WorkspaceAadAdmin
	err := ctx.ReadResource("azure-native:synapse/v20201201:WorkspaceAadAdmin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workspaceAadAdminState struct {
}

type WorkspaceAadAdminState struct {
}

func (WorkspaceAadAdminState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceAadAdminState)(nil)).Elem()
}

type workspaceAadAdminArgs struct {
	AdministratorType *string `pulumi:"administratorType"`
	Login             *string `pulumi:"login"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	Sid               *string `pulumi:"sid"`
	TenantId          *string `pulumi:"tenantId"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


type WorkspaceAadAdminArgs struct {
	AdministratorType pulumi.StringPtrInput
	Login             pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sid               pulumi.StringPtrInput
	TenantId          pulumi.StringPtrInput
	WorkspaceName     pulumi.StringInput
}

func (WorkspaceAadAdminArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceAadAdminArgs)(nil)).Elem()
}

type WorkspaceAadAdminInput interface {
	pulumi.Input

	ToWorkspaceAadAdminOutput() WorkspaceAadAdminOutput
	ToWorkspaceAadAdminOutputWithContext(ctx context.Context) WorkspaceAadAdminOutput
}

func (*WorkspaceAadAdmin) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceAadAdmin)(nil)).Elem()
}

func (i *WorkspaceAadAdmin) ToWorkspaceAadAdminOutput() WorkspaceAadAdminOutput {
	return i.ToWorkspaceAadAdminOutputWithContext(context.Background())
}

func (i *WorkspaceAadAdmin) ToWorkspaceAadAdminOutputWithContext(ctx context.Context) WorkspaceAadAdminOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceAadAdminOutput)
}

type WorkspaceAadAdminOutput struct{ *pulumi.OutputState }

func (WorkspaceAadAdminOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceAadAdmin)(nil)).Elem()
}

func (o WorkspaceAadAdminOutput) ToWorkspaceAadAdminOutput() WorkspaceAadAdminOutput {
	return o
}

func (o WorkspaceAadAdminOutput) ToWorkspaceAadAdminOutputWithContext(ctx context.Context) WorkspaceAadAdminOutput {
	return o
}

func (o WorkspaceAadAdminOutput) AdministratorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceAadAdmin) pulumi.StringPtrOutput { return v.AdministratorType }).(pulumi.StringPtrOutput)
}

func (o WorkspaceAadAdminOutput) Login() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceAadAdmin) pulumi.StringPtrOutput { return v.Login }).(pulumi.StringPtrOutput)
}

func (o WorkspaceAadAdminOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceAadAdmin) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkspaceAadAdminOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceAadAdmin) pulumi.StringPtrOutput { return v.Sid }).(pulumi.StringPtrOutput)
}

func (o WorkspaceAadAdminOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceAadAdmin) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o WorkspaceAadAdminOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceAadAdmin) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceAadAdminOutput{})
}
