


package v20200901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListNotebookKeys(ctx *pulumi.Context, args *ListNotebookKeysArgs, opts ...pulumi.InvokeOption) (*ListNotebookKeysResult, error) {
	var rv ListNotebookKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200901preview:listNotebookKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNotebookKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

type ListNotebookKeysResult struct {
	PrimaryAccessKey   string `pulumi:"primaryAccessKey"`
	SecondaryAccessKey string `pulumi:"secondaryAccessKey"`
}

func ListNotebookKeysOutput(ctx *pulumi.Context, args ListNotebookKeysOutputArgs, opts ...pulumi.InvokeOption) ListNotebookKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListNotebookKeysResult, error) {
			args := v.(ListNotebookKeysArgs)
			r, err := ListNotebookKeys(ctx, &args, opts...)
			var s ListNotebookKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListNotebookKeysResultOutput)
}

type ListNotebookKeysOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListNotebookKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNotebookKeysArgs)(nil)).Elem()
}

type ListNotebookKeysResultOutput struct{ *pulumi.OutputState }

func (ListNotebookKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNotebookKeysResult)(nil)).Elem()
}

func (o ListNotebookKeysResultOutput) ToListNotebookKeysResultOutput() ListNotebookKeysResultOutput {
	return o
}

func (o ListNotebookKeysResultOutput) ToListNotebookKeysResultOutputWithContext(ctx context.Context) ListNotebookKeysResultOutput {
	return o
}

func (o ListNotebookKeysResultOutput) PrimaryAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListNotebookKeysResult) string { return v.PrimaryAccessKey }).(pulumi.StringOutput)
}

func (o ListNotebookKeysResultOutput) SecondaryAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListNotebookKeysResult) string { return v.SecondaryAccessKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListNotebookKeysResultOutput{})
}
