


package logz

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMonitorVMHosts(ctx *pulumi.Context, args *ListMonitorVMHostsArgs, opts ...pulumi.InvokeOption) (*ListMonitorVMHostsResult, error) {
	var rv ListMonitorVMHostsResult
	err := ctx.Invoke("azure-native:logz:listMonitorVMHosts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMonitorVMHostsArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListMonitorVMHostsResult struct {
	NextLink *string               `pulumi:"nextLink"`
	Value    []VMResourcesResponse `pulumi:"value"`
}

func ListMonitorVMHostsOutput(ctx *pulumi.Context, args ListMonitorVMHostsOutputArgs, opts ...pulumi.InvokeOption) ListMonitorVMHostsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMonitorVMHostsResult, error) {
			args := v.(ListMonitorVMHostsArgs)
			r, err := ListMonitorVMHosts(ctx, &args, opts...)
			var s ListMonitorVMHostsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMonitorVMHostsResultOutput)
}

type ListMonitorVMHostsOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListMonitorVMHostsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitorVMHostsArgs)(nil)).Elem()
}


type ListMonitorVMHostsResultOutput struct{ *pulumi.OutputState }

func (ListMonitorVMHostsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMonitorVMHostsResult)(nil)).Elem()
}

func (o ListMonitorVMHostsResultOutput) ToListMonitorVMHostsResultOutput() ListMonitorVMHostsResultOutput {
	return o
}

func (o ListMonitorVMHostsResultOutput) ToListMonitorVMHostsResultOutputWithContext(ctx context.Context) ListMonitorVMHostsResultOutput {
	return o
}

func (o ListMonitorVMHostsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListMonitorVMHostsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListMonitorVMHostsResultOutput) Value() VMResourcesResponseArrayOutput {
	return o.ApplyT(func(v ListMonitorVMHostsResult) []VMResourcesResponse { return v.Value }).(VMResourcesResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMonitorVMHostsResultOutput{})
}
