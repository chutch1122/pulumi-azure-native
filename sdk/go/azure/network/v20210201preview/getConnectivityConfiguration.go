


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupConnectivityConfiguration(ctx *pulumi.Context, args *LookupConnectivityConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupConnectivityConfigurationResult, error) {
	var rv LookupConnectivityConfigurationResult
	err := ctx.Invoke("azure-native:network/v20210201preview:getConnectivityConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectivityConfigurationArgs struct {
	ConfigurationName  string `pulumi:"configurationName"`
	NetworkManagerName string `pulumi:"networkManagerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupConnectivityConfigurationResult struct {
	AppliesToGroups       []ConnectivityGroupItemResponse `pulumi:"appliesToGroups"`
	ConnectivityTopology  string                          `pulumi:"connectivityTopology"`
	DeleteExistingPeering *string                         `pulumi:"deleteExistingPeering"`
	Description           *string                         `pulumi:"description"`
	DisplayName           *string                         `pulumi:"displayName"`
	Etag                  string                          `pulumi:"etag"`
	Hubs                  []HubResponse                   `pulumi:"hubs"`
	Id                    string                          `pulumi:"id"`
	IsGlobal              *string                         `pulumi:"isGlobal"`
	Name                  string                          `pulumi:"name"`
	ProvisioningState     string                          `pulumi:"provisioningState"`
	SystemData            SystemDataResponse              `pulumi:"systemData"`
	Type                  string                          `pulumi:"type"`
}

func LookupConnectivityConfigurationOutput(ctx *pulumi.Context, args LookupConnectivityConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupConnectivityConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectivityConfigurationResult, error) {
			args := v.(LookupConnectivityConfigurationArgs)
			r, err := LookupConnectivityConfiguration(ctx, &args, opts...)
			var s LookupConnectivityConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectivityConfigurationResultOutput)
}

type LookupConnectivityConfigurationOutputArgs struct {
	ConfigurationName  pulumi.StringInput `pulumi:"configurationName"`
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectivityConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectivityConfigurationArgs)(nil)).Elem()
}


type LookupConnectivityConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupConnectivityConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectivityConfigurationResult)(nil)).Elem()
}

func (o LookupConnectivityConfigurationResultOutput) ToLookupConnectivityConfigurationResultOutput() LookupConnectivityConfigurationResultOutput {
	return o
}

func (o LookupConnectivityConfigurationResultOutput) ToLookupConnectivityConfigurationResultOutputWithContext(ctx context.Context) LookupConnectivityConfigurationResultOutput {
	return o
}

func (o LookupConnectivityConfigurationResultOutput) AppliesToGroups() ConnectivityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) []ConnectivityGroupItemResponse {
		return v.AppliesToGroups
	}).(ConnectivityGroupItemResponseArrayOutput)
}

func (o LookupConnectivityConfigurationResultOutput) ConnectivityTopology() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) string { return v.ConnectivityTopology }).(pulumi.StringOutput)
}

func (o LookupConnectivityConfigurationResultOutput) DeleteExistingPeering() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) *string { return v.DeleteExistingPeering }).(pulumi.StringPtrOutput)
}

func (o LookupConnectivityConfigurationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupConnectivityConfigurationResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupConnectivityConfigurationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupConnectivityConfigurationResultOutput) Hubs() HubResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) []HubResponse { return v.Hubs }).(HubResponseArrayOutput)
}

func (o LookupConnectivityConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConnectivityConfigurationResultOutput) IsGlobal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) *string { return v.IsGlobal }).(pulumi.StringPtrOutput)
}

func (o LookupConnectivityConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupConnectivityConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupConnectivityConfigurationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupConnectivityConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectivityConfigurationResultOutput{})
}
