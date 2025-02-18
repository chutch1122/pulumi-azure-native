


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListGatewayKeys(ctx *pulumi.Context, args *ListGatewayKeysArgs, opts ...pulumi.InvokeOption) (*ListGatewayKeysResult, error) {
	var rv ListGatewayKeysResult
	err := ctx.Invoke("azure-native:apimanagement/v20220401preview:listGatewayKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListGatewayKeysArgs struct {
	GatewayId         string `pulumi:"gatewayId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type ListGatewayKeysResult struct {
	Primary   *string `pulumi:"primary"`
	Secondary *string `pulumi:"secondary"`
}

func ListGatewayKeysOutput(ctx *pulumi.Context, args ListGatewayKeysOutputArgs, opts ...pulumi.InvokeOption) ListGatewayKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListGatewayKeysResult, error) {
			args := v.(ListGatewayKeysArgs)
			r, err := ListGatewayKeys(ctx, &args, opts...)
			var s ListGatewayKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListGatewayKeysResultOutput)
}

type ListGatewayKeysOutputArgs struct {
	GatewayId         pulumi.StringInput `pulumi:"gatewayId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (ListGatewayKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListGatewayKeysArgs)(nil)).Elem()
}


type ListGatewayKeysResultOutput struct{ *pulumi.OutputState }

func (ListGatewayKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListGatewayKeysResult)(nil)).Elem()
}

func (o ListGatewayKeysResultOutput) ToListGatewayKeysResultOutput() ListGatewayKeysResultOutput {
	return o
}

func (o ListGatewayKeysResultOutput) ToListGatewayKeysResultOutputWithContext(ctx context.Context) ListGatewayKeysResultOutput {
	return o
}

func (o ListGatewayKeysResultOutput) Primary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListGatewayKeysResult) *string { return v.Primary }).(pulumi.StringPtrOutput)
}

func (o ListGatewayKeysResultOutput) Secondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListGatewayKeysResult) *string { return v.Secondary }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListGatewayKeysResultOutput{})
}
