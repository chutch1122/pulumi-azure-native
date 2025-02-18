


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAFDEndpoint(ctx *pulumi.Context, args *LookupAFDEndpointArgs, opts ...pulumi.InvokeOption) (*LookupAFDEndpointResult, error) {
	var rv LookupAFDEndpointResult
	err := ctx.Invoke("azure-native:cdn/v20210601:getAFDEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAFDEndpointArgs struct {
	EndpointName      string `pulumi:"endpointName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAFDEndpointResult struct {
	AutoGeneratedDomainNameLabelScope *string            `pulumi:"autoGeneratedDomainNameLabelScope"`
	DeploymentStatus                  string             `pulumi:"deploymentStatus"`
	EnabledState                      *string            `pulumi:"enabledState"`
	HostName                          string             `pulumi:"hostName"`
	Id                                string             `pulumi:"id"`
	Location                          string             `pulumi:"location"`
	Name                              string             `pulumi:"name"`
	ProfileName                       string             `pulumi:"profileName"`
	ProvisioningState                 string             `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse `pulumi:"systemData"`
	Tags                              map[string]string  `pulumi:"tags"`
	Type                              string             `pulumi:"type"`
}

func LookupAFDEndpointOutput(ctx *pulumi.Context, args LookupAFDEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupAFDEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAFDEndpointResult, error) {
			args := v.(LookupAFDEndpointArgs)
			r, err := LookupAFDEndpoint(ctx, &args, opts...)
			var s LookupAFDEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAFDEndpointResultOutput)
}

type LookupAFDEndpointOutputArgs struct {
	EndpointName      pulumi.StringInput `pulumi:"endpointName"`
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAFDEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAFDEndpointArgs)(nil)).Elem()
}


type LookupAFDEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupAFDEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAFDEndpointResult)(nil)).Elem()
}

func (o LookupAFDEndpointResultOutput) ToLookupAFDEndpointResultOutput() LookupAFDEndpointResultOutput {
	return o
}

func (o LookupAFDEndpointResultOutput) ToLookupAFDEndpointResultOutputWithContext(ctx context.Context) LookupAFDEndpointResultOutput {
	return o
}

func (o LookupAFDEndpointResultOutput) AutoGeneratedDomainNameLabelScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) *string { return v.AutoGeneratedDomainNameLabelScope }).(pulumi.StringPtrOutput)
}

func (o LookupAFDEndpointResultOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) string { return v.DeploymentStatus }).(pulumi.StringOutput)
}

func (o LookupAFDEndpointResultOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o LookupAFDEndpointResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o LookupAFDEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAFDEndpointResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAFDEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAFDEndpointResultOutput) ProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) string { return v.ProfileName }).(pulumi.StringOutput)
}

func (o LookupAFDEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAFDEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAFDEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAFDEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAFDEndpointResultOutput{})
}
