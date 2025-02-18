


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkspaceNotebookAccessToken(ctx *pulumi.Context, args *ListWorkspaceNotebookAccessTokenArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceNotebookAccessTokenResult, error) {
	var rv ListWorkspaceNotebookAccessTokenResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:listWorkspaceNotebookAccessToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceNotebookAccessTokenArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

type ListWorkspaceNotebookAccessTokenResult struct {
	AccessToken        string `pulumi:"accessToken"`
	ExpiresIn          int    `pulumi:"expiresIn"`
	HostName           string `pulumi:"hostName"`
	NotebookResourceId string `pulumi:"notebookResourceId"`
	PublicDns          string `pulumi:"publicDns"`
	RefreshToken       string `pulumi:"refreshToken"`
	Scope              string `pulumi:"scope"`
	TokenType          string `pulumi:"tokenType"`
}

func ListWorkspaceNotebookAccessTokenOutput(ctx *pulumi.Context, args ListWorkspaceNotebookAccessTokenOutputArgs, opts ...pulumi.InvokeOption) ListWorkspaceNotebookAccessTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkspaceNotebookAccessTokenResult, error) {
			args := v.(ListWorkspaceNotebookAccessTokenArgs)
			r, err := ListWorkspaceNotebookAccessToken(ctx, &args, opts...)
			var s ListWorkspaceNotebookAccessTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkspaceNotebookAccessTokenResultOutput)
}

type ListWorkspaceNotebookAccessTokenOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListWorkspaceNotebookAccessTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceNotebookAccessTokenArgs)(nil)).Elem()
}

type ListWorkspaceNotebookAccessTokenResultOutput struct{ *pulumi.OutputState }

func (ListWorkspaceNotebookAccessTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceNotebookAccessTokenResult)(nil)).Elem()
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) ToListWorkspaceNotebookAccessTokenResultOutput() ListWorkspaceNotebookAccessTokenResultOutput {
	return o
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) ToListWorkspaceNotebookAccessTokenResultOutputWithContext(ctx context.Context) ListWorkspaceNotebookAccessTokenResultOutput {
	return o
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) AccessToken() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookAccessTokenResult) string { return v.AccessToken }).(pulumi.StringOutput)
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) ExpiresIn() pulumi.IntOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookAccessTokenResult) int { return v.ExpiresIn }).(pulumi.IntOutput)
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookAccessTokenResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) NotebookResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookAccessTokenResult) string { return v.NotebookResourceId }).(pulumi.StringOutput)
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) PublicDns() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookAccessTokenResult) string { return v.PublicDns }).(pulumi.StringOutput)
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) RefreshToken() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookAccessTokenResult) string { return v.RefreshToken }).(pulumi.StringOutput)
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookAccessTokenResult) string { return v.Scope }).(pulumi.StringOutput)
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) TokenType() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookAccessTokenResult) string { return v.TokenType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkspaceNotebookAccessTokenResultOutput{})
}
