


package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureActiveDirectory struct {
	ClientApplication  *string `pulumi:"clientApplication"`
	ClusterApplication *string `pulumi:"clusterApplication"`
	TenantId           *string `pulumi:"tenantId"`
}





type AzureActiveDirectoryInput interface {
	pulumi.Input

	ToAzureActiveDirectoryOutput() AzureActiveDirectoryOutput
	ToAzureActiveDirectoryOutputWithContext(context.Context) AzureActiveDirectoryOutput
}

type AzureActiveDirectoryArgs struct {
	ClientApplication  pulumi.StringPtrInput `pulumi:"clientApplication"`
	ClusterApplication pulumi.StringPtrInput `pulumi:"clusterApplication"`
	TenantId           pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (AzureActiveDirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectory)(nil)).Elem()
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryOutput() AzureActiveDirectoryOutput {
	return i.ToAzureActiveDirectoryOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryOutputWithContext(ctx context.Context) AzureActiveDirectoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryOutput)
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return i.ToAzureActiveDirectoryPtrOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryArgs) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryOutput).ToAzureActiveDirectoryPtrOutputWithContext(ctx)
}









type AzureActiveDirectoryPtrInput interface {
	pulumi.Input

	ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput
	ToAzureActiveDirectoryPtrOutputWithContext(context.Context) AzureActiveDirectoryPtrOutput
}

type azureActiveDirectoryPtrType AzureActiveDirectoryArgs

func AzureActiveDirectoryPtr(v *AzureActiveDirectoryArgs) AzureActiveDirectoryPtrInput {
	return (*azureActiveDirectoryPtrType)(v)
}

func (*azureActiveDirectoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectory)(nil)).Elem()
}

func (i *azureActiveDirectoryPtrType) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return i.ToAzureActiveDirectoryPtrOutputWithContext(context.Background())
}

func (i *azureActiveDirectoryPtrType) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryPtrOutput)
}

type AzureActiveDirectoryOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectory)(nil)).Elem()
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryOutput() AzureActiveDirectoryOutput {
	return o
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryOutputWithContext(ctx context.Context) AzureActiveDirectoryOutput {
	return o
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return o.ToAzureActiveDirectoryPtrOutputWithContext(context.Background())
}

func (o AzureActiveDirectoryOutput) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureActiveDirectory) *AzureActiveDirectory {
		return &v
	}).(AzureActiveDirectoryPtrOutput)
}

func (o AzureActiveDirectoryOutput) ClientApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *string { return v.ClientApplication }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryOutput) ClusterApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *string { return v.ClusterApplication }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryPtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectory)(nil)).Elem()
}

func (o AzureActiveDirectoryPtrOutput) ToAzureActiveDirectoryPtrOutput() AzureActiveDirectoryPtrOutput {
	return o
}

func (o AzureActiveDirectoryPtrOutput) ToAzureActiveDirectoryPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryPtrOutput {
	return o
}

func (o AzureActiveDirectoryPtrOutput) Elem() AzureActiveDirectoryOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) AzureActiveDirectory {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectory
		return ret
	}).(AzureActiveDirectoryOutput)
}

func (o AzureActiveDirectoryPtrOutput) ClientApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *string {
		if v == nil {
			return nil
		}
		return v.ClientApplication
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryPtrOutput) ClusterApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *string {
		if v == nil {
			return nil
		}
		return v.ClusterApplication
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryResponse struct {
	ClientApplication  *string `pulumi:"clientApplication"`
	ClusterApplication *string `pulumi:"clusterApplication"`
	TenantId           *string `pulumi:"tenantId"`
}

type AzureActiveDirectoryResponseOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryResponseOutput) ToAzureActiveDirectoryResponseOutput() AzureActiveDirectoryResponseOutput {
	return o
}

func (o AzureActiveDirectoryResponseOutput) ToAzureActiveDirectoryResponseOutputWithContext(ctx context.Context) AzureActiveDirectoryResponseOutput {
	return o
}

func (o AzureActiveDirectoryResponseOutput) ClientApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *string { return v.ClientApplication }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryResponseOutput) ClusterApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *string { return v.ClusterApplication }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryResponsePtrOutput) ToAzureActiveDirectoryResponsePtrOutput() AzureActiveDirectoryResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryResponsePtrOutput) ToAzureActiveDirectoryResponsePtrOutputWithContext(ctx context.Context) AzureActiveDirectoryResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryResponsePtrOutput) Elem() AzureActiveDirectoryResponseOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) AzureActiveDirectoryResponse {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectoryResponse
		return ret
	}).(AzureActiveDirectoryResponseOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) ClientApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientApplication
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) ClusterApplication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClusterApplication
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type ClientCertificate struct {
	CommonName       *string `pulumi:"commonName"`
	IsAdmin          bool    `pulumi:"isAdmin"`
	IssuerThumbprint *string `pulumi:"issuerThumbprint"`
	Thumbprint       *string `pulumi:"thumbprint"`
}





type ClientCertificateInput interface {
	pulumi.Input

	ToClientCertificateOutput() ClientCertificateOutput
	ToClientCertificateOutputWithContext(context.Context) ClientCertificateOutput
}

type ClientCertificateArgs struct {
	CommonName       pulumi.StringPtrInput `pulumi:"commonName"`
	IsAdmin          pulumi.BoolInput      `pulumi:"isAdmin"`
	IssuerThumbprint pulumi.StringPtrInput `pulumi:"issuerThumbprint"`
	Thumbprint       pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (ClientCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificate)(nil)).Elem()
}

func (i ClientCertificateArgs) ToClientCertificateOutput() ClientCertificateOutput {
	return i.ToClientCertificateOutputWithContext(context.Background())
}

func (i ClientCertificateArgs) ToClientCertificateOutputWithContext(ctx context.Context) ClientCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateOutput)
}





type ClientCertificateArrayInput interface {
	pulumi.Input

	ToClientCertificateArrayOutput() ClientCertificateArrayOutput
	ToClientCertificateArrayOutputWithContext(context.Context) ClientCertificateArrayOutput
}

type ClientCertificateArray []ClientCertificateInput

func (ClientCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificate)(nil)).Elem()
}

func (i ClientCertificateArray) ToClientCertificateArrayOutput() ClientCertificateArrayOutput {
	return i.ToClientCertificateArrayOutputWithContext(context.Background())
}

func (i ClientCertificateArray) ToClientCertificateArrayOutputWithContext(ctx context.Context) ClientCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientCertificateArrayOutput)
}

type ClientCertificateOutput struct{ *pulumi.OutputState }

func (ClientCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificate)(nil)).Elem()
}

func (o ClientCertificateOutput) ToClientCertificateOutput() ClientCertificateOutput {
	return o
}

func (o ClientCertificateOutput) ToClientCertificateOutputWithContext(ctx context.Context) ClientCertificateOutput {
	return o
}

func (o ClientCertificateOutput) CommonName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCertificate) *string { return v.CommonName }).(pulumi.StringPtrOutput)
}

func (o ClientCertificateOutput) IsAdmin() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientCertificate) bool { return v.IsAdmin }).(pulumi.BoolOutput)
}

func (o ClientCertificateOutput) IssuerThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCertificate) *string { return v.IssuerThumbprint }).(pulumi.StringPtrOutput)
}

func (o ClientCertificateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCertificate) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type ClientCertificateArrayOutput struct{ *pulumi.OutputState }

func (ClientCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificate)(nil)).Elem()
}

func (o ClientCertificateArrayOutput) ToClientCertificateArrayOutput() ClientCertificateArrayOutput {
	return o
}

func (o ClientCertificateArrayOutput) ToClientCertificateArrayOutputWithContext(ctx context.Context) ClientCertificateArrayOutput {
	return o
}

func (o ClientCertificateArrayOutput) Index(i pulumi.IntInput) ClientCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientCertificate {
		return vs[0].([]ClientCertificate)[vs[1].(int)]
	}).(ClientCertificateOutput)
}

type ClientCertificateResponse struct {
	CommonName       *string `pulumi:"commonName"`
	IsAdmin          bool    `pulumi:"isAdmin"`
	IssuerThumbprint *string `pulumi:"issuerThumbprint"`
	Thumbprint       *string `pulumi:"thumbprint"`
}

type ClientCertificateResponseOutput struct{ *pulumi.OutputState }

func (ClientCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientCertificateResponse)(nil)).Elem()
}

func (o ClientCertificateResponseOutput) ToClientCertificateResponseOutput() ClientCertificateResponseOutput {
	return o
}

func (o ClientCertificateResponseOutput) ToClientCertificateResponseOutputWithContext(ctx context.Context) ClientCertificateResponseOutput {
	return o
}

func (o ClientCertificateResponseOutput) CommonName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCertificateResponse) *string { return v.CommonName }).(pulumi.StringPtrOutput)
}

func (o ClientCertificateResponseOutput) IsAdmin() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientCertificateResponse) bool { return v.IsAdmin }).(pulumi.BoolOutput)
}

func (o ClientCertificateResponseOutput) IssuerThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCertificateResponse) *string { return v.IssuerThumbprint }).(pulumi.StringPtrOutput)
}

func (o ClientCertificateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientCertificateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type ClientCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (ClientCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientCertificateResponse)(nil)).Elem()
}

func (o ClientCertificateResponseArrayOutput) ToClientCertificateResponseArrayOutput() ClientCertificateResponseArrayOutput {
	return o
}

func (o ClientCertificateResponseArrayOutput) ToClientCertificateResponseArrayOutputWithContext(ctx context.Context) ClientCertificateResponseArrayOutput {
	return o
}

func (o ClientCertificateResponseArrayOutput) Index(i pulumi.IntInput) ClientCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientCertificateResponse {
		return vs[0].([]ClientCertificateResponse)[vs[1].(int)]
	}).(ClientCertificateResponseOutput)
}

type EndpointRangeDescription struct {
	EndPort   int `pulumi:"endPort"`
	StartPort int `pulumi:"startPort"`
}





type EndpointRangeDescriptionInput interface {
	pulumi.Input

	ToEndpointRangeDescriptionOutput() EndpointRangeDescriptionOutput
	ToEndpointRangeDescriptionOutputWithContext(context.Context) EndpointRangeDescriptionOutput
}

type EndpointRangeDescriptionArgs struct {
	EndPort   pulumi.IntInput `pulumi:"endPort"`
	StartPort pulumi.IntInput `pulumi:"startPort"`
}

func (EndpointRangeDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointRangeDescription)(nil)).Elem()
}

func (i EndpointRangeDescriptionArgs) ToEndpointRangeDescriptionOutput() EndpointRangeDescriptionOutput {
	return i.ToEndpointRangeDescriptionOutputWithContext(context.Background())
}

func (i EndpointRangeDescriptionArgs) ToEndpointRangeDescriptionOutputWithContext(ctx context.Context) EndpointRangeDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRangeDescriptionOutput)
}

func (i EndpointRangeDescriptionArgs) ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput {
	return i.ToEndpointRangeDescriptionPtrOutputWithContext(context.Background())
}

func (i EndpointRangeDescriptionArgs) ToEndpointRangeDescriptionPtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRangeDescriptionOutput).ToEndpointRangeDescriptionPtrOutputWithContext(ctx)
}









type EndpointRangeDescriptionPtrInput interface {
	pulumi.Input

	ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput
	ToEndpointRangeDescriptionPtrOutputWithContext(context.Context) EndpointRangeDescriptionPtrOutput
}

type endpointRangeDescriptionPtrType EndpointRangeDescriptionArgs

func EndpointRangeDescriptionPtr(v *EndpointRangeDescriptionArgs) EndpointRangeDescriptionPtrInput {
	return (*endpointRangeDescriptionPtrType)(v)
}

func (*endpointRangeDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointRangeDescription)(nil)).Elem()
}

func (i *endpointRangeDescriptionPtrType) ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput {
	return i.ToEndpointRangeDescriptionPtrOutputWithContext(context.Background())
}

func (i *endpointRangeDescriptionPtrType) ToEndpointRangeDescriptionPtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointRangeDescriptionPtrOutput)
}

type EndpointRangeDescriptionOutput struct{ *pulumi.OutputState }

func (EndpointRangeDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointRangeDescription)(nil)).Elem()
}

func (o EndpointRangeDescriptionOutput) ToEndpointRangeDescriptionOutput() EndpointRangeDescriptionOutput {
	return o
}

func (o EndpointRangeDescriptionOutput) ToEndpointRangeDescriptionOutputWithContext(ctx context.Context) EndpointRangeDescriptionOutput {
	return o
}

func (o EndpointRangeDescriptionOutput) ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput {
	return o.ToEndpointRangeDescriptionPtrOutputWithContext(context.Background())
}

func (o EndpointRangeDescriptionOutput) ToEndpointRangeDescriptionPtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointRangeDescription) *EndpointRangeDescription {
		return &v
	}).(EndpointRangeDescriptionPtrOutput)
}

func (o EndpointRangeDescriptionOutput) EndPort() pulumi.IntOutput {
	return o.ApplyT(func(v EndpointRangeDescription) int { return v.EndPort }).(pulumi.IntOutput)
}

func (o EndpointRangeDescriptionOutput) StartPort() pulumi.IntOutput {
	return o.ApplyT(func(v EndpointRangeDescription) int { return v.StartPort }).(pulumi.IntOutput)
}

type EndpointRangeDescriptionPtrOutput struct{ *pulumi.OutputState }

func (EndpointRangeDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointRangeDescription)(nil)).Elem()
}

func (o EndpointRangeDescriptionPtrOutput) ToEndpointRangeDescriptionPtrOutput() EndpointRangeDescriptionPtrOutput {
	return o
}

func (o EndpointRangeDescriptionPtrOutput) ToEndpointRangeDescriptionPtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionPtrOutput {
	return o
}

func (o EndpointRangeDescriptionPtrOutput) Elem() EndpointRangeDescriptionOutput {
	return o.ApplyT(func(v *EndpointRangeDescription) EndpointRangeDescription {
		if v != nil {
			return *v
		}
		var ret EndpointRangeDescription
		return ret
	}).(EndpointRangeDescriptionOutput)
}

func (o EndpointRangeDescriptionPtrOutput) EndPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointRangeDescription) *int {
		if v == nil {
			return nil
		}
		return &v.EndPort
	}).(pulumi.IntPtrOutput)
}

func (o EndpointRangeDescriptionPtrOutput) StartPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointRangeDescription) *int {
		if v == nil {
			return nil
		}
		return &v.StartPort
	}).(pulumi.IntPtrOutput)
}

type EndpointRangeDescriptionResponse struct {
	EndPort   int `pulumi:"endPort"`
	StartPort int `pulumi:"startPort"`
}

type EndpointRangeDescriptionResponseOutput struct{ *pulumi.OutputState }

func (EndpointRangeDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointRangeDescriptionResponse)(nil)).Elem()
}

func (o EndpointRangeDescriptionResponseOutput) ToEndpointRangeDescriptionResponseOutput() EndpointRangeDescriptionResponseOutput {
	return o
}

func (o EndpointRangeDescriptionResponseOutput) ToEndpointRangeDescriptionResponseOutputWithContext(ctx context.Context) EndpointRangeDescriptionResponseOutput {
	return o
}

func (o EndpointRangeDescriptionResponseOutput) EndPort() pulumi.IntOutput {
	return o.ApplyT(func(v EndpointRangeDescriptionResponse) int { return v.EndPort }).(pulumi.IntOutput)
}

func (o EndpointRangeDescriptionResponseOutput) StartPort() pulumi.IntOutput {
	return o.ApplyT(func(v EndpointRangeDescriptionResponse) int { return v.StartPort }).(pulumi.IntOutput)
}

type EndpointRangeDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EndpointRangeDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointRangeDescriptionResponse)(nil)).Elem()
}

func (o EndpointRangeDescriptionResponsePtrOutput) ToEndpointRangeDescriptionResponsePtrOutput() EndpointRangeDescriptionResponsePtrOutput {
	return o
}

func (o EndpointRangeDescriptionResponsePtrOutput) ToEndpointRangeDescriptionResponsePtrOutputWithContext(ctx context.Context) EndpointRangeDescriptionResponsePtrOutput {
	return o
}

func (o EndpointRangeDescriptionResponsePtrOutput) Elem() EndpointRangeDescriptionResponseOutput {
	return o.ApplyT(func(v *EndpointRangeDescriptionResponse) EndpointRangeDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret EndpointRangeDescriptionResponse
		return ret
	}).(EndpointRangeDescriptionResponseOutput)
}

func (o EndpointRangeDescriptionResponsePtrOutput) EndPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointRangeDescriptionResponse) *int {
		if v == nil {
			return nil
		}
		return &v.EndPort
	}).(pulumi.IntPtrOutput)
}

func (o EndpointRangeDescriptionResponsePtrOutput) StartPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EndpointRangeDescriptionResponse) *int {
		if v == nil {
			return nil
		}
		return &v.StartPort
	}).(pulumi.IntPtrOutput)
}

type LoadBalancingRule struct {
	BackendPort      int     `pulumi:"backendPort"`
	FrontendPort     int     `pulumi:"frontendPort"`
	ProbeProtocol    string  `pulumi:"probeProtocol"`
	ProbeRequestPath *string `pulumi:"probeRequestPath"`
	Protocol         string  `pulumi:"protocol"`
}





type LoadBalancingRuleInput interface {
	pulumi.Input

	ToLoadBalancingRuleOutput() LoadBalancingRuleOutput
	ToLoadBalancingRuleOutputWithContext(context.Context) LoadBalancingRuleOutput
}

type LoadBalancingRuleArgs struct {
	BackendPort      pulumi.IntInput       `pulumi:"backendPort"`
	FrontendPort     pulumi.IntInput       `pulumi:"frontendPort"`
	ProbeProtocol    pulumi.StringInput    `pulumi:"probeProtocol"`
	ProbeRequestPath pulumi.StringPtrInput `pulumi:"probeRequestPath"`
	Protocol         pulumi.StringInput    `pulumi:"protocol"`
}

func (LoadBalancingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingRule)(nil)).Elem()
}

func (i LoadBalancingRuleArgs) ToLoadBalancingRuleOutput() LoadBalancingRuleOutput {
	return i.ToLoadBalancingRuleOutputWithContext(context.Background())
}

func (i LoadBalancingRuleArgs) ToLoadBalancingRuleOutputWithContext(ctx context.Context) LoadBalancingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingRuleOutput)
}





type LoadBalancingRuleArrayInput interface {
	pulumi.Input

	ToLoadBalancingRuleArrayOutput() LoadBalancingRuleArrayOutput
	ToLoadBalancingRuleArrayOutputWithContext(context.Context) LoadBalancingRuleArrayOutput
}

type LoadBalancingRuleArray []LoadBalancingRuleInput

func (LoadBalancingRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingRule)(nil)).Elem()
}

func (i LoadBalancingRuleArray) ToLoadBalancingRuleArrayOutput() LoadBalancingRuleArrayOutput {
	return i.ToLoadBalancingRuleArrayOutputWithContext(context.Background())
}

func (i LoadBalancingRuleArray) ToLoadBalancingRuleArrayOutputWithContext(ctx context.Context) LoadBalancingRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingRuleArrayOutput)
}

type LoadBalancingRuleOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingRule)(nil)).Elem()
}

func (o LoadBalancingRuleOutput) ToLoadBalancingRuleOutput() LoadBalancingRuleOutput {
	return o
}

func (o LoadBalancingRuleOutput) ToLoadBalancingRuleOutputWithContext(ctx context.Context) LoadBalancingRuleOutput {
	return o
}

func (o LoadBalancingRuleOutput) BackendPort() pulumi.IntOutput {
	return o.ApplyT(func(v LoadBalancingRule) int { return v.BackendPort }).(pulumi.IntOutput)
}

func (o LoadBalancingRuleOutput) FrontendPort() pulumi.IntOutput {
	return o.ApplyT(func(v LoadBalancingRule) int { return v.FrontendPort }).(pulumi.IntOutput)
}

func (o LoadBalancingRuleOutput) ProbeProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingRule) string { return v.ProbeProtocol }).(pulumi.StringOutput)
}

func (o LoadBalancingRuleOutput) ProbeRequestPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *string { return v.ProbeRequestPath }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingRule) string { return v.Protocol }).(pulumi.StringOutput)
}

type LoadBalancingRuleArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingRule)(nil)).Elem()
}

func (o LoadBalancingRuleArrayOutput) ToLoadBalancingRuleArrayOutput() LoadBalancingRuleArrayOutput {
	return o
}

func (o LoadBalancingRuleArrayOutput) ToLoadBalancingRuleArrayOutputWithContext(ctx context.Context) LoadBalancingRuleArrayOutput {
	return o
}

func (o LoadBalancingRuleArrayOutput) Index(i pulumi.IntInput) LoadBalancingRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancingRule {
		return vs[0].([]LoadBalancingRule)[vs[1].(int)]
	}).(LoadBalancingRuleOutput)
}

type LoadBalancingRuleResponse struct {
	BackendPort      int     `pulumi:"backendPort"`
	FrontendPort     int     `pulumi:"frontendPort"`
	ProbeProtocol    string  `pulumi:"probeProtocol"`
	ProbeRequestPath *string `pulumi:"probeRequestPath"`
	Protocol         string  `pulumi:"protocol"`
}

type LoadBalancingRuleResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingRuleResponse)(nil)).Elem()
}

func (o LoadBalancingRuleResponseOutput) ToLoadBalancingRuleResponseOutput() LoadBalancingRuleResponseOutput {
	return o
}

func (o LoadBalancingRuleResponseOutput) ToLoadBalancingRuleResponseOutputWithContext(ctx context.Context) LoadBalancingRuleResponseOutput {
	return o
}

func (o LoadBalancingRuleResponseOutput) BackendPort() pulumi.IntOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) int { return v.BackendPort }).(pulumi.IntOutput)
}

func (o LoadBalancingRuleResponseOutput) FrontendPort() pulumi.IntOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) int { return v.FrontendPort }).(pulumi.IntOutput)
}

func (o LoadBalancingRuleResponseOutput) ProbeProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) string { return v.ProbeProtocol }).(pulumi.StringOutput)
}

func (o LoadBalancingRuleResponseOutput) ProbeRequestPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *string { return v.ProbeRequestPath }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

type LoadBalancingRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingRuleResponse)(nil)).Elem()
}

func (o LoadBalancingRuleResponseArrayOutput) ToLoadBalancingRuleResponseArrayOutput() LoadBalancingRuleResponseArrayOutput {
	return o
}

func (o LoadBalancingRuleResponseArrayOutput) ToLoadBalancingRuleResponseArrayOutputWithContext(ctx context.Context) LoadBalancingRuleResponseArrayOutput {
	return o
}

func (o LoadBalancingRuleResponseArrayOutput) Index(i pulumi.IntInput) LoadBalancingRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancingRuleResponse {
		return vs[0].([]LoadBalancingRuleResponse)[vs[1].(int)]
	}).(LoadBalancingRuleResponseOutput)
}

type SettingsParameterDescription struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type SettingsParameterDescriptionInput interface {
	pulumi.Input

	ToSettingsParameterDescriptionOutput() SettingsParameterDescriptionOutput
	ToSettingsParameterDescriptionOutputWithContext(context.Context) SettingsParameterDescriptionOutput
}

type SettingsParameterDescriptionArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (SettingsParameterDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsParameterDescription)(nil)).Elem()
}

func (i SettingsParameterDescriptionArgs) ToSettingsParameterDescriptionOutput() SettingsParameterDescriptionOutput {
	return i.ToSettingsParameterDescriptionOutputWithContext(context.Background())
}

func (i SettingsParameterDescriptionArgs) ToSettingsParameterDescriptionOutputWithContext(ctx context.Context) SettingsParameterDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsParameterDescriptionOutput)
}





type SettingsParameterDescriptionArrayInput interface {
	pulumi.Input

	ToSettingsParameterDescriptionArrayOutput() SettingsParameterDescriptionArrayOutput
	ToSettingsParameterDescriptionArrayOutputWithContext(context.Context) SettingsParameterDescriptionArrayOutput
}

type SettingsParameterDescriptionArray []SettingsParameterDescriptionInput

func (SettingsParameterDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsParameterDescription)(nil)).Elem()
}

func (i SettingsParameterDescriptionArray) ToSettingsParameterDescriptionArrayOutput() SettingsParameterDescriptionArrayOutput {
	return i.ToSettingsParameterDescriptionArrayOutputWithContext(context.Background())
}

func (i SettingsParameterDescriptionArray) ToSettingsParameterDescriptionArrayOutputWithContext(ctx context.Context) SettingsParameterDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsParameterDescriptionArrayOutput)
}

type SettingsParameterDescriptionOutput struct{ *pulumi.OutputState }

func (SettingsParameterDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsParameterDescription)(nil)).Elem()
}

func (o SettingsParameterDescriptionOutput) ToSettingsParameterDescriptionOutput() SettingsParameterDescriptionOutput {
	return o
}

func (o SettingsParameterDescriptionOutput) ToSettingsParameterDescriptionOutputWithContext(ctx context.Context) SettingsParameterDescriptionOutput {
	return o
}

func (o SettingsParameterDescriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsParameterDescription) string { return v.Name }).(pulumi.StringOutput)
}

func (o SettingsParameterDescriptionOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsParameterDescription) string { return v.Value }).(pulumi.StringOutput)
}

type SettingsParameterDescriptionArrayOutput struct{ *pulumi.OutputState }

func (SettingsParameterDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsParameterDescription)(nil)).Elem()
}

func (o SettingsParameterDescriptionArrayOutput) ToSettingsParameterDescriptionArrayOutput() SettingsParameterDescriptionArrayOutput {
	return o
}

func (o SettingsParameterDescriptionArrayOutput) ToSettingsParameterDescriptionArrayOutputWithContext(ctx context.Context) SettingsParameterDescriptionArrayOutput {
	return o
}

func (o SettingsParameterDescriptionArrayOutput) Index(i pulumi.IntInput) SettingsParameterDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingsParameterDescription {
		return vs[0].([]SettingsParameterDescription)[vs[1].(int)]
	}).(SettingsParameterDescriptionOutput)
}

type SettingsParameterDescriptionResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type SettingsParameterDescriptionResponseOutput struct{ *pulumi.OutputState }

func (SettingsParameterDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsParameterDescriptionResponse)(nil)).Elem()
}

func (o SettingsParameterDescriptionResponseOutput) ToSettingsParameterDescriptionResponseOutput() SettingsParameterDescriptionResponseOutput {
	return o
}

func (o SettingsParameterDescriptionResponseOutput) ToSettingsParameterDescriptionResponseOutputWithContext(ctx context.Context) SettingsParameterDescriptionResponseOutput {
	return o
}

func (o SettingsParameterDescriptionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsParameterDescriptionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SettingsParameterDescriptionResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsParameterDescriptionResponse) string { return v.Value }).(pulumi.StringOutput)
}

type SettingsParameterDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (SettingsParameterDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsParameterDescriptionResponse)(nil)).Elem()
}

func (o SettingsParameterDescriptionResponseArrayOutput) ToSettingsParameterDescriptionResponseArrayOutput() SettingsParameterDescriptionResponseArrayOutput {
	return o
}

func (o SettingsParameterDescriptionResponseArrayOutput) ToSettingsParameterDescriptionResponseArrayOutputWithContext(ctx context.Context) SettingsParameterDescriptionResponseArrayOutput {
	return o
}

func (o SettingsParameterDescriptionResponseArrayOutput) Index(i pulumi.IntInput) SettingsParameterDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingsParameterDescriptionResponse {
		return vs[0].([]SettingsParameterDescriptionResponse)[vs[1].(int)]
	}).(SettingsParameterDescriptionResponseOutput)
}

type SettingsSectionDescription struct {
	Name       string                         `pulumi:"name"`
	Parameters []SettingsParameterDescription `pulumi:"parameters"`
}





type SettingsSectionDescriptionInput interface {
	pulumi.Input

	ToSettingsSectionDescriptionOutput() SettingsSectionDescriptionOutput
	ToSettingsSectionDescriptionOutputWithContext(context.Context) SettingsSectionDescriptionOutput
}

type SettingsSectionDescriptionArgs struct {
	Name       pulumi.StringInput                     `pulumi:"name"`
	Parameters SettingsParameterDescriptionArrayInput `pulumi:"parameters"`
}

func (SettingsSectionDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsSectionDescription)(nil)).Elem()
}

func (i SettingsSectionDescriptionArgs) ToSettingsSectionDescriptionOutput() SettingsSectionDescriptionOutput {
	return i.ToSettingsSectionDescriptionOutputWithContext(context.Background())
}

func (i SettingsSectionDescriptionArgs) ToSettingsSectionDescriptionOutputWithContext(ctx context.Context) SettingsSectionDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsSectionDescriptionOutput)
}





type SettingsSectionDescriptionArrayInput interface {
	pulumi.Input

	ToSettingsSectionDescriptionArrayOutput() SettingsSectionDescriptionArrayOutput
	ToSettingsSectionDescriptionArrayOutputWithContext(context.Context) SettingsSectionDescriptionArrayOutput
}

type SettingsSectionDescriptionArray []SettingsSectionDescriptionInput

func (SettingsSectionDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsSectionDescription)(nil)).Elem()
}

func (i SettingsSectionDescriptionArray) ToSettingsSectionDescriptionArrayOutput() SettingsSectionDescriptionArrayOutput {
	return i.ToSettingsSectionDescriptionArrayOutputWithContext(context.Background())
}

func (i SettingsSectionDescriptionArray) ToSettingsSectionDescriptionArrayOutputWithContext(ctx context.Context) SettingsSectionDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsSectionDescriptionArrayOutput)
}

type SettingsSectionDescriptionOutput struct{ *pulumi.OutputState }

func (SettingsSectionDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsSectionDescription)(nil)).Elem()
}

func (o SettingsSectionDescriptionOutput) ToSettingsSectionDescriptionOutput() SettingsSectionDescriptionOutput {
	return o
}

func (o SettingsSectionDescriptionOutput) ToSettingsSectionDescriptionOutputWithContext(ctx context.Context) SettingsSectionDescriptionOutput {
	return o
}

func (o SettingsSectionDescriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsSectionDescription) string { return v.Name }).(pulumi.StringOutput)
}

func (o SettingsSectionDescriptionOutput) Parameters() SettingsParameterDescriptionArrayOutput {
	return o.ApplyT(func(v SettingsSectionDescription) []SettingsParameterDescription { return v.Parameters }).(SettingsParameterDescriptionArrayOutput)
}

type SettingsSectionDescriptionArrayOutput struct{ *pulumi.OutputState }

func (SettingsSectionDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsSectionDescription)(nil)).Elem()
}

func (o SettingsSectionDescriptionArrayOutput) ToSettingsSectionDescriptionArrayOutput() SettingsSectionDescriptionArrayOutput {
	return o
}

func (o SettingsSectionDescriptionArrayOutput) ToSettingsSectionDescriptionArrayOutputWithContext(ctx context.Context) SettingsSectionDescriptionArrayOutput {
	return o
}

func (o SettingsSectionDescriptionArrayOutput) Index(i pulumi.IntInput) SettingsSectionDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingsSectionDescription {
		return vs[0].([]SettingsSectionDescription)[vs[1].(int)]
	}).(SettingsSectionDescriptionOutput)
}

type SettingsSectionDescriptionResponse struct {
	Name       string                                 `pulumi:"name"`
	Parameters []SettingsParameterDescriptionResponse `pulumi:"parameters"`
}

type SettingsSectionDescriptionResponseOutput struct{ *pulumi.OutputState }

func (SettingsSectionDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsSectionDescriptionResponse)(nil)).Elem()
}

func (o SettingsSectionDescriptionResponseOutput) ToSettingsSectionDescriptionResponseOutput() SettingsSectionDescriptionResponseOutput {
	return o
}

func (o SettingsSectionDescriptionResponseOutput) ToSettingsSectionDescriptionResponseOutputWithContext(ctx context.Context) SettingsSectionDescriptionResponseOutput {
	return o
}

func (o SettingsSectionDescriptionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsSectionDescriptionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SettingsSectionDescriptionResponseOutput) Parameters() SettingsParameterDescriptionResponseArrayOutput {
	return o.ApplyT(func(v SettingsSectionDescriptionResponse) []SettingsParameterDescriptionResponse { return v.Parameters }).(SettingsParameterDescriptionResponseArrayOutput)
}

type SettingsSectionDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (SettingsSectionDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsSectionDescriptionResponse)(nil)).Elem()
}

func (o SettingsSectionDescriptionResponseArrayOutput) ToSettingsSectionDescriptionResponseArrayOutput() SettingsSectionDescriptionResponseArrayOutput {
	return o
}

func (o SettingsSectionDescriptionResponseArrayOutput) ToSettingsSectionDescriptionResponseArrayOutputWithContext(ctx context.Context) SettingsSectionDescriptionResponseArrayOutput {
	return o
}

func (o SettingsSectionDescriptionResponseArrayOutput) Index(i pulumi.IntInput) SettingsSectionDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingsSectionDescriptionResponse {
		return vs[0].([]SettingsSectionDescriptionResponse)[vs[1].(int)]
	}).(SettingsSectionDescriptionResponseOutput)
}

type Sku struct {
	Name string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SubResource struct {
	Id *string `pulumi:"id"`
}





type SubResourceInput interface {
	pulumi.Input

	ToSubResourceOutput() SubResourceOutput
	ToSubResourceOutputWithContext(context.Context) SubResourceOutput
}

type SubResourceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (i SubResourceArgs) ToSubResourceOutput() SubResourceOutput {
	return i.ToSubResourceOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput)
}

type SubResourceOutput struct{ *pulumi.OutputState }

func (SubResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (o SubResourceOutput) ToSubResourceOutput() SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return o
}

func (o SubResourceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResource) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourceResponse struct {
	Id *string `pulumi:"id"`
}

type SubResourceResponseOutput struct{ *pulumi.OutputState }

func (SubResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type VMSSExtension struct {
	AutoUpgradeMinorVersion  *bool       `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag           *string     `pulumi:"forceUpdateTag"`
	Name                     string      `pulumi:"name"`
	ProtectedSettings        interface{} `pulumi:"protectedSettings"`
	ProvisionAfterExtensions []string    `pulumi:"provisionAfterExtensions"`
	Publisher                string      `pulumi:"publisher"`
	Settings                 interface{} `pulumi:"settings"`
	Type                     string      `pulumi:"type"`
	TypeHandlerVersion       string      `pulumi:"typeHandlerVersion"`
}





type VMSSExtensionInput interface {
	pulumi.Input

	ToVMSSExtensionOutput() VMSSExtensionOutput
	ToVMSSExtensionOutputWithContext(context.Context) VMSSExtensionOutput
}

type VMSSExtensionArgs struct {
	AutoUpgradeMinorVersion  pulumi.BoolPtrInput     `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag           pulumi.StringPtrInput   `pulumi:"forceUpdateTag"`
	Name                     pulumi.StringInput      `pulumi:"name"`
	ProtectedSettings        pulumi.Input            `pulumi:"protectedSettings"`
	ProvisionAfterExtensions pulumi.StringArrayInput `pulumi:"provisionAfterExtensions"`
	Publisher                pulumi.StringInput      `pulumi:"publisher"`
	Settings                 pulumi.Input            `pulumi:"settings"`
	Type                     pulumi.StringInput      `pulumi:"type"`
	TypeHandlerVersion       pulumi.StringInput      `pulumi:"typeHandlerVersion"`
}

func (VMSSExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSSExtension)(nil)).Elem()
}

func (i VMSSExtensionArgs) ToVMSSExtensionOutput() VMSSExtensionOutput {
	return i.ToVMSSExtensionOutputWithContext(context.Background())
}

func (i VMSSExtensionArgs) ToVMSSExtensionOutputWithContext(ctx context.Context) VMSSExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSSExtensionOutput)
}





type VMSSExtensionArrayInput interface {
	pulumi.Input

	ToVMSSExtensionArrayOutput() VMSSExtensionArrayOutput
	ToVMSSExtensionArrayOutputWithContext(context.Context) VMSSExtensionArrayOutput
}

type VMSSExtensionArray []VMSSExtensionInput

func (VMSSExtensionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMSSExtension)(nil)).Elem()
}

func (i VMSSExtensionArray) ToVMSSExtensionArrayOutput() VMSSExtensionArrayOutput {
	return i.ToVMSSExtensionArrayOutputWithContext(context.Background())
}

func (i VMSSExtensionArray) ToVMSSExtensionArrayOutputWithContext(ctx context.Context) VMSSExtensionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSSExtensionArrayOutput)
}

type VMSSExtensionOutput struct{ *pulumi.OutputState }

func (VMSSExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSSExtension)(nil)).Elem()
}

func (o VMSSExtensionOutput) ToVMSSExtensionOutput() VMSSExtensionOutput {
	return o
}

func (o VMSSExtensionOutput) ToVMSSExtensionOutputWithContext(ctx context.Context) VMSSExtensionOutput {
	return o
}

func (o VMSSExtensionOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VMSSExtension) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o VMSSExtensionOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMSSExtension) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o VMSSExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtension) string { return v.Name }).(pulumi.StringOutput)
}

func (o VMSSExtensionOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v VMSSExtension) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o VMSSExtensionOutput) ProvisionAfterExtensions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VMSSExtension) []string { return v.ProvisionAfterExtensions }).(pulumi.StringArrayOutput)
}

func (o VMSSExtensionOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtension) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o VMSSExtensionOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v VMSSExtension) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o VMSSExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtension) string { return v.Type }).(pulumi.StringOutput)
}

func (o VMSSExtensionOutput) TypeHandlerVersion() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtension) string { return v.TypeHandlerVersion }).(pulumi.StringOutput)
}

type VMSSExtensionArrayOutput struct{ *pulumi.OutputState }

func (VMSSExtensionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMSSExtension)(nil)).Elem()
}

func (o VMSSExtensionArrayOutput) ToVMSSExtensionArrayOutput() VMSSExtensionArrayOutput {
	return o
}

func (o VMSSExtensionArrayOutput) ToVMSSExtensionArrayOutputWithContext(ctx context.Context) VMSSExtensionArrayOutput {
	return o
}

func (o VMSSExtensionArrayOutput) Index(i pulumi.IntInput) VMSSExtensionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VMSSExtension {
		return vs[0].([]VMSSExtension)[vs[1].(int)]
	}).(VMSSExtensionOutput)
}

type VMSSExtensionResponse struct {
	AutoUpgradeMinorVersion  *bool       `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag           *string     `pulumi:"forceUpdateTag"`
	Name                     string      `pulumi:"name"`
	ProtectedSettings        interface{} `pulumi:"protectedSettings"`
	ProvisionAfterExtensions []string    `pulumi:"provisionAfterExtensions"`
	ProvisioningState        string      `pulumi:"provisioningState"`
	Publisher                string      `pulumi:"publisher"`
	Settings                 interface{} `pulumi:"settings"`
	Type                     string      `pulumi:"type"`
	TypeHandlerVersion       string      `pulumi:"typeHandlerVersion"`
}

type VMSSExtensionResponseOutput struct{ *pulumi.OutputState }

func (VMSSExtensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSSExtensionResponse)(nil)).Elem()
}

func (o VMSSExtensionResponseOutput) ToVMSSExtensionResponseOutput() VMSSExtensionResponseOutput {
	return o
}

func (o VMSSExtensionResponseOutput) ToVMSSExtensionResponseOutputWithContext(ctx context.Context) VMSSExtensionResponseOutput {
	return o
}

func (o VMSSExtensionResponseOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o VMSSExtensionResponseOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o VMSSExtensionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VMSSExtensionResponseOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o VMSSExtensionResponseOutput) ProvisionAfterExtensions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) []string { return v.ProvisionAfterExtensions }).(pulumi.StringArrayOutput)
}

func (o VMSSExtensionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VMSSExtensionResponseOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o VMSSExtensionResponseOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o VMSSExtensionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VMSSExtensionResponseOutput) TypeHandlerVersion() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSExtensionResponse) string { return v.TypeHandlerVersion }).(pulumi.StringOutput)
}

type VMSSExtensionResponseArrayOutput struct{ *pulumi.OutputState }

func (VMSSExtensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMSSExtensionResponse)(nil)).Elem()
}

func (o VMSSExtensionResponseArrayOutput) ToVMSSExtensionResponseArrayOutput() VMSSExtensionResponseArrayOutput {
	return o
}

func (o VMSSExtensionResponseArrayOutput) ToVMSSExtensionResponseArrayOutputWithContext(ctx context.Context) VMSSExtensionResponseArrayOutput {
	return o
}

func (o VMSSExtensionResponseArrayOutput) Index(i pulumi.IntInput) VMSSExtensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VMSSExtensionResponse {
		return vs[0].([]VMSSExtensionResponse)[vs[1].(int)]
	}).(VMSSExtensionResponseOutput)
}

type VaultCertificate struct {
	CertificateStore string `pulumi:"certificateStore"`
	CertificateUrl   string `pulumi:"certificateUrl"`
}





type VaultCertificateInput interface {
	pulumi.Input

	ToVaultCertificateOutput() VaultCertificateOutput
	ToVaultCertificateOutputWithContext(context.Context) VaultCertificateOutput
}

type VaultCertificateArgs struct {
	CertificateStore pulumi.StringInput `pulumi:"certificateStore"`
	CertificateUrl   pulumi.StringInput `pulumi:"certificateUrl"`
}

func (VaultCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultCertificate)(nil)).Elem()
}

func (i VaultCertificateArgs) ToVaultCertificateOutput() VaultCertificateOutput {
	return i.ToVaultCertificateOutputWithContext(context.Background())
}

func (i VaultCertificateArgs) ToVaultCertificateOutputWithContext(ctx context.Context) VaultCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultCertificateOutput)
}





type VaultCertificateArrayInput interface {
	pulumi.Input

	ToVaultCertificateArrayOutput() VaultCertificateArrayOutput
	ToVaultCertificateArrayOutputWithContext(context.Context) VaultCertificateArrayOutput
}

type VaultCertificateArray []VaultCertificateInput

func (VaultCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultCertificate)(nil)).Elem()
}

func (i VaultCertificateArray) ToVaultCertificateArrayOutput() VaultCertificateArrayOutput {
	return i.ToVaultCertificateArrayOutputWithContext(context.Background())
}

func (i VaultCertificateArray) ToVaultCertificateArrayOutputWithContext(ctx context.Context) VaultCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultCertificateArrayOutput)
}

type VaultCertificateOutput struct{ *pulumi.OutputState }

func (VaultCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultCertificate)(nil)).Elem()
}

func (o VaultCertificateOutput) ToVaultCertificateOutput() VaultCertificateOutput {
	return o
}

func (o VaultCertificateOutput) ToVaultCertificateOutputWithContext(ctx context.Context) VaultCertificateOutput {
	return o
}

func (o VaultCertificateOutput) CertificateStore() pulumi.StringOutput {
	return o.ApplyT(func(v VaultCertificate) string { return v.CertificateStore }).(pulumi.StringOutput)
}

func (o VaultCertificateOutput) CertificateUrl() pulumi.StringOutput {
	return o.ApplyT(func(v VaultCertificate) string { return v.CertificateUrl }).(pulumi.StringOutput)
}

type VaultCertificateArrayOutput struct{ *pulumi.OutputState }

func (VaultCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultCertificate)(nil)).Elem()
}

func (o VaultCertificateArrayOutput) ToVaultCertificateArrayOutput() VaultCertificateArrayOutput {
	return o
}

func (o VaultCertificateArrayOutput) ToVaultCertificateArrayOutputWithContext(ctx context.Context) VaultCertificateArrayOutput {
	return o
}

func (o VaultCertificateArrayOutput) Index(i pulumi.IntInput) VaultCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultCertificate {
		return vs[0].([]VaultCertificate)[vs[1].(int)]
	}).(VaultCertificateOutput)
}

type VaultCertificateResponse struct {
	CertificateStore string `pulumi:"certificateStore"`
	CertificateUrl   string `pulumi:"certificateUrl"`
}

type VaultCertificateResponseOutput struct{ *pulumi.OutputState }

func (VaultCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultCertificateResponse)(nil)).Elem()
}

func (o VaultCertificateResponseOutput) ToVaultCertificateResponseOutput() VaultCertificateResponseOutput {
	return o
}

func (o VaultCertificateResponseOutput) ToVaultCertificateResponseOutputWithContext(ctx context.Context) VaultCertificateResponseOutput {
	return o
}

func (o VaultCertificateResponseOutput) CertificateStore() pulumi.StringOutput {
	return o.ApplyT(func(v VaultCertificateResponse) string { return v.CertificateStore }).(pulumi.StringOutput)
}

func (o VaultCertificateResponseOutput) CertificateUrl() pulumi.StringOutput {
	return o.ApplyT(func(v VaultCertificateResponse) string { return v.CertificateUrl }).(pulumi.StringOutput)
}

type VaultCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VaultCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultCertificateResponse)(nil)).Elem()
}

func (o VaultCertificateResponseArrayOutput) ToVaultCertificateResponseArrayOutput() VaultCertificateResponseArrayOutput {
	return o
}

func (o VaultCertificateResponseArrayOutput) ToVaultCertificateResponseArrayOutputWithContext(ctx context.Context) VaultCertificateResponseArrayOutput {
	return o
}

func (o VaultCertificateResponseArrayOutput) Index(i pulumi.IntInput) VaultCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultCertificateResponse {
		return vs[0].([]VaultCertificateResponse)[vs[1].(int)]
	}).(VaultCertificateResponseOutput)
}

type VaultSecretGroup struct {
	SourceVault       SubResource        `pulumi:"sourceVault"`
	VaultCertificates []VaultCertificate `pulumi:"vaultCertificates"`
}





type VaultSecretGroupInput interface {
	pulumi.Input

	ToVaultSecretGroupOutput() VaultSecretGroupOutput
	ToVaultSecretGroupOutputWithContext(context.Context) VaultSecretGroupOutput
}

type VaultSecretGroupArgs struct {
	SourceVault       SubResourceInput           `pulumi:"sourceVault"`
	VaultCertificates VaultCertificateArrayInput `pulumi:"vaultCertificates"`
}

func (VaultSecretGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultSecretGroup)(nil)).Elem()
}

func (i VaultSecretGroupArgs) ToVaultSecretGroupOutput() VaultSecretGroupOutput {
	return i.ToVaultSecretGroupOutputWithContext(context.Background())
}

func (i VaultSecretGroupArgs) ToVaultSecretGroupOutputWithContext(ctx context.Context) VaultSecretGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultSecretGroupOutput)
}





type VaultSecretGroupArrayInput interface {
	pulumi.Input

	ToVaultSecretGroupArrayOutput() VaultSecretGroupArrayOutput
	ToVaultSecretGroupArrayOutputWithContext(context.Context) VaultSecretGroupArrayOutput
}

type VaultSecretGroupArray []VaultSecretGroupInput

func (VaultSecretGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultSecretGroup)(nil)).Elem()
}

func (i VaultSecretGroupArray) ToVaultSecretGroupArrayOutput() VaultSecretGroupArrayOutput {
	return i.ToVaultSecretGroupArrayOutputWithContext(context.Background())
}

func (i VaultSecretGroupArray) ToVaultSecretGroupArrayOutputWithContext(ctx context.Context) VaultSecretGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultSecretGroupArrayOutput)
}

type VaultSecretGroupOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultSecretGroup)(nil)).Elem()
}

func (o VaultSecretGroupOutput) ToVaultSecretGroupOutput() VaultSecretGroupOutput {
	return o
}

func (o VaultSecretGroupOutput) ToVaultSecretGroupOutputWithContext(ctx context.Context) VaultSecretGroupOutput {
	return o
}

func (o VaultSecretGroupOutput) SourceVault() SubResourceOutput {
	return o.ApplyT(func(v VaultSecretGroup) SubResource { return v.SourceVault }).(SubResourceOutput)
}

func (o VaultSecretGroupOutput) VaultCertificates() VaultCertificateArrayOutput {
	return o.ApplyT(func(v VaultSecretGroup) []VaultCertificate { return v.VaultCertificates }).(VaultCertificateArrayOutput)
}

type VaultSecretGroupArrayOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultSecretGroup)(nil)).Elem()
}

func (o VaultSecretGroupArrayOutput) ToVaultSecretGroupArrayOutput() VaultSecretGroupArrayOutput {
	return o
}

func (o VaultSecretGroupArrayOutput) ToVaultSecretGroupArrayOutputWithContext(ctx context.Context) VaultSecretGroupArrayOutput {
	return o
}

func (o VaultSecretGroupArrayOutput) Index(i pulumi.IntInput) VaultSecretGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultSecretGroup {
		return vs[0].([]VaultSecretGroup)[vs[1].(int)]
	}).(VaultSecretGroupOutput)
}

type VaultSecretGroupResponse struct {
	SourceVault       SubResourceResponse        `pulumi:"sourceVault"`
	VaultCertificates []VaultCertificateResponse `pulumi:"vaultCertificates"`
}

type VaultSecretGroupResponseOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultSecretGroupResponse)(nil)).Elem()
}

func (o VaultSecretGroupResponseOutput) ToVaultSecretGroupResponseOutput() VaultSecretGroupResponseOutput {
	return o
}

func (o VaultSecretGroupResponseOutput) ToVaultSecretGroupResponseOutputWithContext(ctx context.Context) VaultSecretGroupResponseOutput {
	return o
}

func (o VaultSecretGroupResponseOutput) SourceVault() SubResourceResponseOutput {
	return o.ApplyT(func(v VaultSecretGroupResponse) SubResourceResponse { return v.SourceVault }).(SubResourceResponseOutput)
}

func (o VaultSecretGroupResponseOutput) VaultCertificates() VaultCertificateResponseArrayOutput {
	return o.ApplyT(func(v VaultSecretGroupResponse) []VaultCertificateResponse { return v.VaultCertificates }).(VaultCertificateResponseArrayOutput)
}

type VaultSecretGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (VaultSecretGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VaultSecretGroupResponse)(nil)).Elem()
}

func (o VaultSecretGroupResponseArrayOutput) ToVaultSecretGroupResponseArrayOutput() VaultSecretGroupResponseArrayOutput {
	return o
}

func (o VaultSecretGroupResponseArrayOutput) ToVaultSecretGroupResponseArrayOutputWithContext(ctx context.Context) VaultSecretGroupResponseArrayOutput {
	return o
}

func (o VaultSecretGroupResponseArrayOutput) Index(i pulumi.IntInput) VaultSecretGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VaultSecretGroupResponse {
		return vs[0].([]VaultSecretGroupResponse)[vs[1].(int)]
	}).(VaultSecretGroupResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureActiveDirectoryOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryPtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryResponseOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryResponsePtrOutput{})
	pulumi.RegisterOutputType(ClientCertificateOutput{})
	pulumi.RegisterOutputType(ClientCertificateArrayOutput{})
	pulumi.RegisterOutputType(ClientCertificateResponseOutput{})
	pulumi.RegisterOutputType(ClientCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(EndpointRangeDescriptionOutput{})
	pulumi.RegisterOutputType(EndpointRangeDescriptionPtrOutput{})
	pulumi.RegisterOutputType(EndpointRangeDescriptionResponseOutput{})
	pulumi.RegisterOutputType(EndpointRangeDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(SettingsParameterDescriptionOutput{})
	pulumi.RegisterOutputType(SettingsParameterDescriptionArrayOutput{})
	pulumi.RegisterOutputType(SettingsParameterDescriptionResponseOutput{})
	pulumi.RegisterOutputType(SettingsParameterDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(SettingsSectionDescriptionOutput{})
	pulumi.RegisterOutputType(SettingsSectionDescriptionArrayOutput{})
	pulumi.RegisterOutputType(SettingsSectionDescriptionResponseOutput{})
	pulumi.RegisterOutputType(SettingsSectionDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(VMSSExtensionOutput{})
	pulumi.RegisterOutputType(VMSSExtensionArrayOutput{})
	pulumi.RegisterOutputType(VMSSExtensionResponseOutput{})
	pulumi.RegisterOutputType(VMSSExtensionResponseArrayOutput{})
	pulumi.RegisterOutputType(VaultCertificateOutput{})
	pulumi.RegisterOutputType(VaultCertificateArrayOutput{})
	pulumi.RegisterOutputType(VaultCertificateResponseOutput{})
	pulumi.RegisterOutputType(VaultCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupArrayOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupResponseOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupResponseArrayOutput{})
}
