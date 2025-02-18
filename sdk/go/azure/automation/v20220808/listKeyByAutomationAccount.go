


package v20220808

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListKeyByAutomationAccount(ctx *pulumi.Context, args *ListKeyByAutomationAccountArgs, opts ...pulumi.InvokeOption) (*ListKeyByAutomationAccountResult, error) {
	var rv ListKeyByAutomationAccountResult
	err := ctx.Invoke("azure-native:automation/v20220808:listKeyByAutomationAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListKeyByAutomationAccountArgs struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}

type ListKeyByAutomationAccountResult struct {
	Keys []KeyResponse `pulumi:"keys"`
}

func ListKeyByAutomationAccountOutput(ctx *pulumi.Context, args ListKeyByAutomationAccountOutputArgs, opts ...pulumi.InvokeOption) ListKeyByAutomationAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListKeyByAutomationAccountResult, error) {
			args := v.(ListKeyByAutomationAccountArgs)
			r, err := ListKeyByAutomationAccount(ctx, &args, opts...)
			var s ListKeyByAutomationAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListKeyByAutomationAccountResultOutput)
}

type ListKeyByAutomationAccountOutputArgs struct {
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListKeyByAutomationAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListKeyByAutomationAccountArgs)(nil)).Elem()
}

type ListKeyByAutomationAccountResultOutput struct{ *pulumi.OutputState }

func (ListKeyByAutomationAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListKeyByAutomationAccountResult)(nil)).Elem()
}

func (o ListKeyByAutomationAccountResultOutput) ToListKeyByAutomationAccountResultOutput() ListKeyByAutomationAccountResultOutput {
	return o
}

func (o ListKeyByAutomationAccountResultOutput) ToListKeyByAutomationAccountResultOutputWithContext(ctx context.Context) ListKeyByAutomationAccountResultOutput {
	return o
}

func (o ListKeyByAutomationAccountResultOutput) Keys() KeyResponseArrayOutput {
	return o.ApplyT(func(v ListKeyByAutomationAccountResult) []KeyResponse { return v.Keys }).(KeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListKeyByAutomationAccountResultOutput{})
}
