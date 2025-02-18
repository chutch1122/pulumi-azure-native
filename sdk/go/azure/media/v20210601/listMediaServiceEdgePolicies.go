


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMediaServiceEdgePolicies(ctx *pulumi.Context, args *ListMediaServiceEdgePoliciesArgs, opts ...pulumi.InvokeOption) (*ListMediaServiceEdgePoliciesResult, error) {
	var rv ListMediaServiceEdgePoliciesResult
	err := ctx.Invoke("azure-native:media/v20210601:listMediaServiceEdgePolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMediaServiceEdgePoliciesArgs struct {
	AccountName       string  `pulumi:"accountName"`
	DeviceId          *string `pulumi:"deviceId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}

type ListMediaServiceEdgePoliciesResult struct {
	UsageDataCollectionPolicy *EdgeUsageDataCollectionPolicyResponse `pulumi:"usageDataCollectionPolicy"`
}

func ListMediaServiceEdgePoliciesOutput(ctx *pulumi.Context, args ListMediaServiceEdgePoliciesOutputArgs, opts ...pulumi.InvokeOption) ListMediaServiceEdgePoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMediaServiceEdgePoliciesResult, error) {
			args := v.(ListMediaServiceEdgePoliciesArgs)
			r, err := ListMediaServiceEdgePolicies(ctx, &args, opts...)
			var s ListMediaServiceEdgePoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMediaServiceEdgePoliciesResultOutput)
}

type ListMediaServiceEdgePoliciesOutputArgs struct {
	AccountName       pulumi.StringInput    `pulumi:"accountName"`
	DeviceId          pulumi.StringPtrInput `pulumi:"deviceId"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (ListMediaServiceEdgePoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMediaServiceEdgePoliciesArgs)(nil)).Elem()
}

type ListMediaServiceEdgePoliciesResultOutput struct{ *pulumi.OutputState }

func (ListMediaServiceEdgePoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMediaServiceEdgePoliciesResult)(nil)).Elem()
}

func (o ListMediaServiceEdgePoliciesResultOutput) ToListMediaServiceEdgePoliciesResultOutput() ListMediaServiceEdgePoliciesResultOutput {
	return o
}

func (o ListMediaServiceEdgePoliciesResultOutput) ToListMediaServiceEdgePoliciesResultOutputWithContext(ctx context.Context) ListMediaServiceEdgePoliciesResultOutput {
	return o
}

func (o ListMediaServiceEdgePoliciesResultOutput) UsageDataCollectionPolicy() EdgeUsageDataCollectionPolicyResponsePtrOutput {
	return o.ApplyT(func(v ListMediaServiceEdgePoliciesResult) *EdgeUsageDataCollectionPolicyResponse {
		return v.UsageDataCollectionPolicy
	}).(EdgeUsageDataCollectionPolicyResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMediaServiceEdgePoliciesResultOutput{})
}
