


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListManagedClusterUserCredentials(ctx *pulumi.Context, args *ListManagedClusterUserCredentialsArgs, opts ...pulumi.InvokeOption) (*ListManagedClusterUserCredentialsResult, error) {
	var rv ListManagedClusterUserCredentialsResult
	err := ctx.Invoke("azure-native:containerservice/v20220301:listManagedClusterUserCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagedClusterUserCredentialsArgs struct {
	Format            *string `pulumi:"format"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceName      string  `pulumi:"resourceName"`
	ServerFqdn        *string `pulumi:"serverFqdn"`
}


type ListManagedClusterUserCredentialsResult struct {
	Kubeconfigs []CredentialResultResponse `pulumi:"kubeconfigs"`
}

func ListManagedClusterUserCredentialsOutput(ctx *pulumi.Context, args ListManagedClusterUserCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListManagedClusterUserCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListManagedClusterUserCredentialsResult, error) {
			args := v.(ListManagedClusterUserCredentialsArgs)
			r, err := ListManagedClusterUserCredentials(ctx, &args, opts...)
			var s ListManagedClusterUserCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListManagedClusterUserCredentialsResultOutput)
}

type ListManagedClusterUserCredentialsOutputArgs struct {
	Format            pulumi.StringPtrInput `pulumi:"format"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput    `pulumi:"resourceName"`
	ServerFqdn        pulumi.StringPtrInput `pulumi:"serverFqdn"`
}

func (ListManagedClusterUserCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterUserCredentialsArgs)(nil)).Elem()
}


type ListManagedClusterUserCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListManagedClusterUserCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterUserCredentialsResult)(nil)).Elem()
}

func (o ListManagedClusterUserCredentialsResultOutput) ToListManagedClusterUserCredentialsResultOutput() ListManagedClusterUserCredentialsResultOutput {
	return o
}

func (o ListManagedClusterUserCredentialsResultOutput) ToListManagedClusterUserCredentialsResultOutputWithContext(ctx context.Context) ListManagedClusterUserCredentialsResultOutput {
	return o
}

func (o ListManagedClusterUserCredentialsResultOutput) Kubeconfigs() CredentialResultResponseArrayOutput {
	return o.ApplyT(func(v ListManagedClusterUserCredentialsResult) []CredentialResultResponse { return v.Kubeconfigs }).(CredentialResultResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListManagedClusterUserCredentialsResultOutput{})
}
