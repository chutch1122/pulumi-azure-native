


package v20171115preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	var rv LookupProjectResult
	err := ctx.Invoke("azure-native:datamigration/v20171115preview:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupProjectArgs struct {
	GroupName   string `pulumi:"groupName"`
	ProjectName string `pulumi:"projectName"`
	ServiceName string `pulumi:"serviceName"`
}


type LookupProjectResult struct {
	CreationTime         string                     `pulumi:"creationTime"`
	DatabasesInfo        []DatabaseInfoResponse     `pulumi:"databasesInfo"`
	Id                   string                     `pulumi:"id"`
	Location             string                     `pulumi:"location"`
	Name                 string                     `pulumi:"name"`
	ProvisioningState    string                     `pulumi:"provisioningState"`
	SourceConnectionInfo *SqlConnectionInfoResponse `pulumi:"sourceConnectionInfo"`
	SourcePlatform       string                     `pulumi:"sourcePlatform"`
	Tags                 map[string]string          `pulumi:"tags"`
	TargetConnectionInfo *SqlConnectionInfoResponse `pulumi:"targetConnectionInfo"`
	TargetPlatform       string                     `pulumi:"targetPlatform"`
	Type                 string                     `pulumi:"type"`
}


func (val *LookupProjectResult) Defaults() *LookupProjectResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceConnectionInfo = tmp.SourceConnectionInfo.Defaults()

	tmp.TargetConnectionInfo = tmp.TargetConnectionInfo.Defaults()

	return &tmp
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			var s LookupProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectResultOutput)
}

type LookupProjectOutputArgs struct {
	GroupName   pulumi.StringInput `pulumi:"groupName"`
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}


type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) DatabasesInfo() DatabaseInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []DatabaseInfoResponse { return v.DatabasesInfo }).(DatabaseInfoResponseArrayOutput)
}

func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) SourceConnectionInfo() SqlConnectionInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *SqlConnectionInfoResponse { return v.SourceConnectionInfo }).(SqlConnectionInfoResponsePtrOutput)
}

func (o LookupProjectResultOutput) SourcePlatform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.SourcePlatform }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProjectResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupProjectResultOutput) TargetConnectionInfo() SqlConnectionInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *SqlConnectionInfoResponse { return v.TargetConnectionInfo }).(SqlConnectionInfoResponsePtrOutput)
}

func (o LookupProjectResultOutput) TargetPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.TargetPlatform }).(pulumi.StringOutput)
}

func (o LookupProjectResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
