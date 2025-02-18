


package logic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkflowTriggerCallbackUrl(ctx *pulumi.Context, args *ListWorkflowTriggerCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListWorkflowTriggerCallbackUrlResult, error) {
	var rv ListWorkflowTriggerCallbackUrlResult
	err := ctx.Invoke("azure-native:logic:listWorkflowTriggerCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowTriggerCallbackUrlArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TriggerName       string `pulumi:"triggerName"`
	WorkflowName      string `pulumi:"workflowName"`
}


type ListWorkflowTriggerCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

func ListWorkflowTriggerCallbackUrlOutput(ctx *pulumi.Context, args ListWorkflowTriggerCallbackUrlOutputArgs, opts ...pulumi.InvokeOption) ListWorkflowTriggerCallbackUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkflowTriggerCallbackUrlResult, error) {
			args := v.(ListWorkflowTriggerCallbackUrlArgs)
			r, err := ListWorkflowTriggerCallbackUrl(ctx, &args, opts...)
			var s ListWorkflowTriggerCallbackUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkflowTriggerCallbackUrlResultOutput)
}

type ListWorkflowTriggerCallbackUrlOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TriggerName       pulumi.StringInput `pulumi:"triggerName"`
	WorkflowName      pulumi.StringInput `pulumi:"workflowName"`
}

func (ListWorkflowTriggerCallbackUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowTriggerCallbackUrlArgs)(nil)).Elem()
}


type ListWorkflowTriggerCallbackUrlResultOutput struct{ *pulumi.OutputState }

func (ListWorkflowTriggerCallbackUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowTriggerCallbackUrlResult)(nil)).Elem()
}

func (o ListWorkflowTriggerCallbackUrlResultOutput) ToListWorkflowTriggerCallbackUrlResultOutput() ListWorkflowTriggerCallbackUrlResultOutput {
	return o
}

func (o ListWorkflowTriggerCallbackUrlResultOutput) ToListWorkflowTriggerCallbackUrlResultOutputWithContext(ctx context.Context) ListWorkflowTriggerCallbackUrlResultOutput {
	return o
}

func (o ListWorkflowTriggerCallbackUrlResultOutput) BasePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowTriggerCallbackUrlResult) string { return v.BasePath }).(pulumi.StringOutput)
}

func (o ListWorkflowTriggerCallbackUrlResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowTriggerCallbackUrlResult) string { return v.Method }).(pulumi.StringOutput)
}

func (o ListWorkflowTriggerCallbackUrlResultOutput) Queries() WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o.ApplyT(func(v ListWorkflowTriggerCallbackUrlResult) *WorkflowTriggerListCallbackUrlQueriesResponse {
		return v.Queries
	}).(WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput)
}

func (o ListWorkflowTriggerCallbackUrlResultOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowTriggerCallbackUrlResult) string { return v.RelativePath }).(pulumi.StringOutput)
}

func (o ListWorkflowTriggerCallbackUrlResultOutput) RelativePathParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWorkflowTriggerCallbackUrlResult) []string { return v.RelativePathParameters }).(pulumi.StringArrayOutput)
}

func (o ListWorkflowTriggerCallbackUrlResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowTriggerCallbackUrlResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkflowTriggerCallbackUrlResultOutput{})
}
