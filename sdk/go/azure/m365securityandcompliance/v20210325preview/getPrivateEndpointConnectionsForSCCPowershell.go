


package v20210325preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpointConnectionsForSCCPowershell(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionsForSCCPowershellArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionsForSCCPowershellResult, error) {
	var rv LookupPrivateEndpointConnectionsForSCCPowershellResult
	err := ctx.Invoke("azure-native:m365securityandcompliance/v20210325preview:getPrivateEndpointConnectionsForSCCPowershell", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionsForSCCPowershellArgs struct {
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	ResourceName                  string `pulumi:"resourceName"`
}


type LookupPrivateEndpointConnectionsForSCCPowershellResult struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
}

func LookupPrivateEndpointConnectionsForSCCPowershellOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionsForSCCPowershellOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionsForSCCPowershellResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionsForSCCPowershellResult, error) {
			args := v.(LookupPrivateEndpointConnectionsForSCCPowershellArgs)
			r, err := LookupPrivateEndpointConnectionsForSCCPowershell(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionsForSCCPowershellResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionsForSCCPowershellResultOutput)
}

type LookupPrivateEndpointConnectionsForSCCPowershellOutputArgs struct {
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName                  pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupPrivateEndpointConnectionsForSCCPowershellOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionsForSCCPowershellArgs)(nil)).Elem()
}


type LookupPrivateEndpointConnectionsForSCCPowershellResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionsForSCCPowershellResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) ToLookupPrivateEndpointConnectionsForSCCPowershellResultOutput() LookupPrivateEndpointConnectionsForSCCPowershellResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) ToLookupPrivateEndpointConnectionsForSCCPowershellResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionsForSCCPowershellResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForSCCPowershellResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForSCCPowershellResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForSCCPowershellResult) *PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForSCCPowershellResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForSCCPowershellResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForSCCPowershellResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPrivateEndpointConnectionsForSCCPowershellResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionsForSCCPowershellResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionsForSCCPowershellResultOutput{})
}
