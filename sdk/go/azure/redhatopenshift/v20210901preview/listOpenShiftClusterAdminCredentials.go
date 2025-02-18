


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListOpenShiftClusterAdminCredentials(ctx *pulumi.Context, args *ListOpenShiftClusterAdminCredentialsArgs, opts ...pulumi.InvokeOption) (*ListOpenShiftClusterAdminCredentialsResult, error) {
	var rv ListOpenShiftClusterAdminCredentialsResult
	err := ctx.Invoke("azure-native:redhatopenshift/v20210901preview:listOpenShiftClusterAdminCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListOpenShiftClusterAdminCredentialsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type ListOpenShiftClusterAdminCredentialsResult struct {
	Kubeconfig *string `pulumi:"kubeconfig"`
}

func ListOpenShiftClusterAdminCredentialsOutput(ctx *pulumi.Context, args ListOpenShiftClusterAdminCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListOpenShiftClusterAdminCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListOpenShiftClusterAdminCredentialsResult, error) {
			args := v.(ListOpenShiftClusterAdminCredentialsArgs)
			r, err := ListOpenShiftClusterAdminCredentials(ctx, &args, opts...)
			var s ListOpenShiftClusterAdminCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListOpenShiftClusterAdminCredentialsResultOutput)
}

type ListOpenShiftClusterAdminCredentialsOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (ListOpenShiftClusterAdminCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOpenShiftClusterAdminCredentialsArgs)(nil)).Elem()
}


type ListOpenShiftClusterAdminCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListOpenShiftClusterAdminCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOpenShiftClusterAdminCredentialsResult)(nil)).Elem()
}

func (o ListOpenShiftClusterAdminCredentialsResultOutput) ToListOpenShiftClusterAdminCredentialsResultOutput() ListOpenShiftClusterAdminCredentialsResultOutput {
	return o
}

func (o ListOpenShiftClusterAdminCredentialsResultOutput) ToListOpenShiftClusterAdminCredentialsResultOutputWithContext(ctx context.Context) ListOpenShiftClusterAdminCredentialsResultOutput {
	return o
}

func (o ListOpenShiftClusterAdminCredentialsResultOutput) Kubeconfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListOpenShiftClusterAdminCredentialsResult) *string { return v.Kubeconfig }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListOpenShiftClusterAdminCredentialsResultOutput{})
}
