


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkspaceKeys(ctx *pulumi.Context, args *ListWorkspaceKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceKeysResult, error) {
	var rv ListWorkspaceKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20220601preview:listWorkspaceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

type ListWorkspaceKeysResult struct {
	AppInsightsInstrumentationKey string                                `pulumi:"appInsightsInstrumentationKey"`
	ContainerRegistryCredentials  RegistryListCredentialsResultResponse `pulumi:"containerRegistryCredentials"`
	NotebookAccessKeys            ListNotebookKeysResultResponse        `pulumi:"notebookAccessKeys"`
	UserStorageKey                string                                `pulumi:"userStorageKey"`
	UserStorageResourceId         string                                `pulumi:"userStorageResourceId"`
}

func ListWorkspaceKeysOutput(ctx *pulumi.Context, args ListWorkspaceKeysOutputArgs, opts ...pulumi.InvokeOption) ListWorkspaceKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkspaceKeysResult, error) {
			args := v.(ListWorkspaceKeysArgs)
			r, err := ListWorkspaceKeys(ctx, &args, opts...)
			var s ListWorkspaceKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkspaceKeysResultOutput)
}

type ListWorkspaceKeysOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListWorkspaceKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceKeysArgs)(nil)).Elem()
}

type ListWorkspaceKeysResultOutput struct{ *pulumi.OutputState }

func (ListWorkspaceKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceKeysResult)(nil)).Elem()
}

func (o ListWorkspaceKeysResultOutput) ToListWorkspaceKeysResultOutput() ListWorkspaceKeysResultOutput {
	return o
}

func (o ListWorkspaceKeysResultOutput) ToListWorkspaceKeysResultOutputWithContext(ctx context.Context) ListWorkspaceKeysResultOutput {
	return o
}

func (o ListWorkspaceKeysResultOutput) AppInsightsInstrumentationKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceKeysResult) string { return v.AppInsightsInstrumentationKey }).(pulumi.StringOutput)
}

func (o ListWorkspaceKeysResultOutput) ContainerRegistryCredentials() RegistryListCredentialsResultResponseOutput {
	return o.ApplyT(func(v ListWorkspaceKeysResult) RegistryListCredentialsResultResponse {
		return v.ContainerRegistryCredentials
	}).(RegistryListCredentialsResultResponseOutput)
}

func (o ListWorkspaceKeysResultOutput) NotebookAccessKeys() ListNotebookKeysResultResponseOutput {
	return o.ApplyT(func(v ListWorkspaceKeysResult) ListNotebookKeysResultResponse { return v.NotebookAccessKeys }).(ListNotebookKeysResultResponseOutput)
}

func (o ListWorkspaceKeysResultOutput) UserStorageKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceKeysResult) string { return v.UserStorageKey }).(pulumi.StringOutput)
}

func (o ListWorkspaceKeysResultOutput) UserStorageResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceKeysResult) string { return v.UserStorageResourceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkspaceKeysResultOutput{})
}
