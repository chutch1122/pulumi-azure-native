


package migrate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerCollector(ctx *pulumi.Context, args *LookupServerCollectorArgs, opts ...pulumi.InvokeOption) (*LookupServerCollectorResult, error) {
	var rv LookupServerCollectorResult
	err := ctx.Invoke("azure-native:migrate:getServerCollector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerCollectorArgs struct {
	ProjectName         string `pulumi:"projectName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	ServerCollectorName string `pulumi:"serverCollectorName"`
}

type LookupServerCollectorResult struct {
	ETag       *string                     `pulumi:"eTag"`
	Id         string                      `pulumi:"id"`
	Name       string                      `pulumi:"name"`
	Properties CollectorPropertiesResponse `pulumi:"properties"`
	Type       string                      `pulumi:"type"`
}

func LookupServerCollectorOutput(ctx *pulumi.Context, args LookupServerCollectorOutputArgs, opts ...pulumi.InvokeOption) LookupServerCollectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerCollectorResult, error) {
			args := v.(LookupServerCollectorArgs)
			r, err := LookupServerCollector(ctx, &args, opts...)
			var s LookupServerCollectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerCollectorResultOutput)
}

type LookupServerCollectorOutputArgs struct {
	ProjectName         pulumi.StringInput `pulumi:"projectName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerCollectorName pulumi.StringInput `pulumi:"serverCollectorName"`
}

func (LookupServerCollectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerCollectorArgs)(nil)).Elem()
}

type LookupServerCollectorResultOutput struct{ *pulumi.OutputState }

func (LookupServerCollectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerCollectorResult)(nil)).Elem()
}

func (o LookupServerCollectorResultOutput) ToLookupServerCollectorResultOutput() LookupServerCollectorResultOutput {
	return o
}

func (o LookupServerCollectorResultOutput) ToLookupServerCollectorResultOutputWithContext(ctx context.Context) LookupServerCollectorResultOutput {
	return o
}

func (o LookupServerCollectorResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerCollectorResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupServerCollectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCollectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerCollectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCollectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerCollectorResultOutput) Properties() CollectorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupServerCollectorResult) CollectorPropertiesResponse { return v.Properties }).(CollectorPropertiesResponseOutput)
}

func (o LookupServerCollectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCollectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerCollectorResultOutput{})
}
