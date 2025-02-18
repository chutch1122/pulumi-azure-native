


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRunLogSasUrl(ctx *pulumi.Context, args *ListRunLogSasUrlArgs, opts ...pulumi.InvokeOption) (*ListRunLogSasUrlResult, error) {
	var rv ListRunLogSasUrlResult
	err := ctx.Invoke("azure-native:containerregistry/v20190601preview:listRunLogSasUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRunLogSasUrlArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RunId             string `pulumi:"runId"`
}


type ListRunLogSasUrlResult struct {
	LogArtifactLink *string `pulumi:"logArtifactLink"`
	LogLink         *string `pulumi:"logLink"`
}

func ListRunLogSasUrlOutput(ctx *pulumi.Context, args ListRunLogSasUrlOutputArgs, opts ...pulumi.InvokeOption) ListRunLogSasUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListRunLogSasUrlResult, error) {
			args := v.(ListRunLogSasUrlArgs)
			r, err := ListRunLogSasUrl(ctx, &args, opts...)
			var s ListRunLogSasUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListRunLogSasUrlResultOutput)
}

type ListRunLogSasUrlOutputArgs struct {
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RunId             pulumi.StringInput `pulumi:"runId"`
}

func (ListRunLogSasUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRunLogSasUrlArgs)(nil)).Elem()
}


type ListRunLogSasUrlResultOutput struct{ *pulumi.OutputState }

func (ListRunLogSasUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRunLogSasUrlResult)(nil)).Elem()
}

func (o ListRunLogSasUrlResultOutput) ToListRunLogSasUrlResultOutput() ListRunLogSasUrlResultOutput {
	return o
}

func (o ListRunLogSasUrlResultOutput) ToListRunLogSasUrlResultOutputWithContext(ctx context.Context) ListRunLogSasUrlResultOutput {
	return o
}

func (o ListRunLogSasUrlResultOutput) LogArtifactLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListRunLogSasUrlResult) *string { return v.LogArtifactLink }).(pulumi.StringPtrOutput)
}

func (o ListRunLogSasUrlResultOutput) LogLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListRunLogSasUrlResult) *string { return v.LogLink }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRunLogSasUrlResultOutput{})
}
