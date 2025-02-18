


package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppPrivateEndpointConnectionSlot(ctx *pulumi.Context, args *LookupWebAppPrivateEndpointConnectionSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppPrivateEndpointConnectionSlotResult, error) {
	var rv LookupWebAppPrivateEndpointConnectionSlotResult
	err := ctx.Invoke("azure-native:web:getWebAppPrivateEndpointConnectionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppPrivateEndpointConnectionSlotArgs struct {
	Name                          string `pulumi:"name"`
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	Slot                          string `pulumi:"slot"`
}


type LookupWebAppPrivateEndpointConnectionSlotResult struct {
	Id                                string                              `pulumi:"id"`
	IpAddresses                       []string                            `pulumi:"ipAddresses"`
	Kind                              *string                             `pulumi:"kind"`
	Name                              string                              `pulumi:"name"`
	PrivateEndpoint                   *ArmIdWrapperResponse               `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                              `pulumi:"provisioningState"`
	Type                              string                              `pulumi:"type"`
}

func LookupWebAppPrivateEndpointConnectionSlotOutput(ctx *pulumi.Context, args LookupWebAppPrivateEndpointConnectionSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppPrivateEndpointConnectionSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppPrivateEndpointConnectionSlotResult, error) {
			args := v.(LookupWebAppPrivateEndpointConnectionSlotArgs)
			r, err := LookupWebAppPrivateEndpointConnectionSlot(ctx, &args, opts...)
			var s LookupWebAppPrivateEndpointConnectionSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppPrivateEndpointConnectionSlotResultOutput)
}

type LookupWebAppPrivateEndpointConnectionSlotOutputArgs struct {
	Name                          pulumi.StringInput `pulumi:"name"`
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot                          pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppPrivateEndpointConnectionSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppPrivateEndpointConnectionSlotArgs)(nil)).Elem()
}


type LookupWebAppPrivateEndpointConnectionSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppPrivateEndpointConnectionSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppPrivateEndpointConnectionSlotResult)(nil)).Elem()
}

func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) ToLookupWebAppPrivateEndpointConnectionSlotResultOutput() LookupWebAppPrivateEndpointConnectionSlotResultOutput {
	return o
}

func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) ToLookupWebAppPrivateEndpointConnectionSlotResultOutputWithContext(ctx context.Context) LookupWebAppPrivateEndpointConnectionSlotResultOutput {
	return o
}

func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionSlotResult) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) PrivateEndpoint() ArmIdWrapperResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionSlotResult) *ArmIdWrapperResponse {
		return v.PrivateEndpoint
	}).(ArmIdWrapperResponsePtrOutput)
}

func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionSlotResult) *PrivateLinkConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStateResponsePtrOutput)
}

func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionSlotResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWebAppPrivateEndpointConnectionSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPrivateEndpointConnectionSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppPrivateEndpointConnectionSlotResultOutput{})
}
