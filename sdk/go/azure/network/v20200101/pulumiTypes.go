


package v20200101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ARecord struct {
	Ipv4Address *string `pulumi:"ipv4Address"`
}





type ARecordInput interface {
	pulumi.Input

	ToARecordOutput() ARecordOutput
	ToARecordOutputWithContext(context.Context) ARecordOutput
}

type ARecordArgs struct {
	Ipv4Address pulumi.StringPtrInput `pulumi:"ipv4Address"`
}

func (ARecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ARecord)(nil)).Elem()
}

func (i ARecordArgs) ToARecordOutput() ARecordOutput {
	return i.ToARecordOutputWithContext(context.Background())
}

func (i ARecordArgs) ToARecordOutputWithContext(ctx context.Context) ARecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ARecordOutput)
}





type ARecordArrayInput interface {
	pulumi.Input

	ToARecordArrayOutput() ARecordArrayOutput
	ToARecordArrayOutputWithContext(context.Context) ARecordArrayOutput
}

type ARecordArray []ARecordInput

func (ARecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ARecord)(nil)).Elem()
}

func (i ARecordArray) ToARecordArrayOutput() ARecordArrayOutput {
	return i.ToARecordArrayOutputWithContext(context.Background())
}

func (i ARecordArray) ToARecordArrayOutputWithContext(ctx context.Context) ARecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ARecordArrayOutput)
}

type ARecordOutput struct{ *pulumi.OutputState }

func (ARecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ARecord)(nil)).Elem()
}

func (o ARecordOutput) ToARecordOutput() ARecordOutput {
	return o
}

func (o ARecordOutput) ToARecordOutputWithContext(ctx context.Context) ARecordOutput {
	return o
}

func (o ARecordOutput) Ipv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ARecord) *string { return v.Ipv4Address }).(pulumi.StringPtrOutput)
}

type ARecordArrayOutput struct{ *pulumi.OutputState }

func (ARecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ARecord)(nil)).Elem()
}

func (o ARecordArrayOutput) ToARecordArrayOutput() ARecordArrayOutput {
	return o
}

func (o ARecordArrayOutput) ToARecordArrayOutputWithContext(ctx context.Context) ARecordArrayOutput {
	return o
}

func (o ARecordArrayOutput) Index(i pulumi.IntInput) ARecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ARecord {
		return vs[0].([]ARecord)[vs[1].(int)]
	}).(ARecordOutput)
}

type ARecordResponse struct {
	Ipv4Address *string `pulumi:"ipv4Address"`
}

type ARecordResponseOutput struct{ *pulumi.OutputState }

func (ARecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ARecordResponse)(nil)).Elem()
}

func (o ARecordResponseOutput) ToARecordResponseOutput() ARecordResponseOutput {
	return o
}

func (o ARecordResponseOutput) ToARecordResponseOutputWithContext(ctx context.Context) ARecordResponseOutput {
	return o
}

func (o ARecordResponseOutput) Ipv4Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ARecordResponse) *string { return v.Ipv4Address }).(pulumi.StringPtrOutput)
}

type ARecordResponseArrayOutput struct{ *pulumi.OutputState }

func (ARecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ARecordResponse)(nil)).Elem()
}

func (o ARecordResponseArrayOutput) ToARecordResponseArrayOutput() ARecordResponseArrayOutput {
	return o
}

func (o ARecordResponseArrayOutput) ToARecordResponseArrayOutputWithContext(ctx context.Context) ARecordResponseArrayOutput {
	return o
}

func (o ARecordResponseArrayOutput) Index(i pulumi.IntInput) ARecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ARecordResponse {
		return vs[0].([]ARecordResponse)[vs[1].(int)]
	}).(ARecordResponseOutput)
}

type AaaaRecord struct {
	Ipv6Address *string `pulumi:"ipv6Address"`
}





type AaaaRecordInput interface {
	pulumi.Input

	ToAaaaRecordOutput() AaaaRecordOutput
	ToAaaaRecordOutputWithContext(context.Context) AaaaRecordOutput
}

type AaaaRecordArgs struct {
	Ipv6Address pulumi.StringPtrInput `pulumi:"ipv6Address"`
}

func (AaaaRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AaaaRecord)(nil)).Elem()
}

func (i AaaaRecordArgs) ToAaaaRecordOutput() AaaaRecordOutput {
	return i.ToAaaaRecordOutputWithContext(context.Background())
}

func (i AaaaRecordArgs) ToAaaaRecordOutputWithContext(ctx context.Context) AaaaRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AaaaRecordOutput)
}





type AaaaRecordArrayInput interface {
	pulumi.Input

	ToAaaaRecordArrayOutput() AaaaRecordArrayOutput
	ToAaaaRecordArrayOutputWithContext(context.Context) AaaaRecordArrayOutput
}

type AaaaRecordArray []AaaaRecordInput

func (AaaaRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AaaaRecord)(nil)).Elem()
}

func (i AaaaRecordArray) ToAaaaRecordArrayOutput() AaaaRecordArrayOutput {
	return i.ToAaaaRecordArrayOutputWithContext(context.Background())
}

func (i AaaaRecordArray) ToAaaaRecordArrayOutputWithContext(ctx context.Context) AaaaRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AaaaRecordArrayOutput)
}

type AaaaRecordOutput struct{ *pulumi.OutputState }

func (AaaaRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AaaaRecord)(nil)).Elem()
}

func (o AaaaRecordOutput) ToAaaaRecordOutput() AaaaRecordOutput {
	return o
}

func (o AaaaRecordOutput) ToAaaaRecordOutputWithContext(ctx context.Context) AaaaRecordOutput {
	return o
}

func (o AaaaRecordOutput) Ipv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AaaaRecord) *string { return v.Ipv6Address }).(pulumi.StringPtrOutput)
}

type AaaaRecordArrayOutput struct{ *pulumi.OutputState }

func (AaaaRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AaaaRecord)(nil)).Elem()
}

func (o AaaaRecordArrayOutput) ToAaaaRecordArrayOutput() AaaaRecordArrayOutput {
	return o
}

func (o AaaaRecordArrayOutput) ToAaaaRecordArrayOutputWithContext(ctx context.Context) AaaaRecordArrayOutput {
	return o
}

func (o AaaaRecordArrayOutput) Index(i pulumi.IntInput) AaaaRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AaaaRecord {
		return vs[0].([]AaaaRecord)[vs[1].(int)]
	}).(AaaaRecordOutput)
}

type AaaaRecordResponse struct {
	Ipv6Address *string `pulumi:"ipv6Address"`
}

type AaaaRecordResponseOutput struct{ *pulumi.OutputState }

func (AaaaRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AaaaRecordResponse)(nil)).Elem()
}

func (o AaaaRecordResponseOutput) ToAaaaRecordResponseOutput() AaaaRecordResponseOutput {
	return o
}

func (o AaaaRecordResponseOutput) ToAaaaRecordResponseOutputWithContext(ctx context.Context) AaaaRecordResponseOutput {
	return o
}

func (o AaaaRecordResponseOutput) Ipv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AaaaRecordResponse) *string { return v.Ipv6Address }).(pulumi.StringPtrOutput)
}

type AaaaRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (AaaaRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AaaaRecordResponse)(nil)).Elem()
}

func (o AaaaRecordResponseArrayOutput) ToAaaaRecordResponseArrayOutput() AaaaRecordResponseArrayOutput {
	return o
}

func (o AaaaRecordResponseArrayOutput) ToAaaaRecordResponseArrayOutputWithContext(ctx context.Context) AaaaRecordResponseArrayOutput {
	return o
}

func (o AaaaRecordResponseArrayOutput) Index(i pulumi.IntInput) AaaaRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AaaaRecordResponse {
		return vs[0].([]AaaaRecordResponse)[vs[1].(int)]
	}).(AaaaRecordResponseOutput)
}

type Backend struct {
	Address                    *string `pulumi:"address"`
	BackendHostHeader          *string `pulumi:"backendHostHeader"`
	EnabledState               *string `pulumi:"enabledState"`
	HttpPort                   *int    `pulumi:"httpPort"`
	HttpsPort                  *int    `pulumi:"httpsPort"`
	Priority                   *int    `pulumi:"priority"`
	PrivateLinkAlias           *string `pulumi:"privateLinkAlias"`
	PrivateLinkApprovalMessage *string `pulumi:"privateLinkApprovalMessage"`
	Weight                     *int    `pulumi:"weight"`
}





type BackendInput interface {
	pulumi.Input

	ToBackendOutput() BackendOutput
	ToBackendOutputWithContext(context.Context) BackendOutput
}

type BackendArgs struct {
	Address                    pulumi.StringPtrInput `pulumi:"address"`
	BackendHostHeader          pulumi.StringPtrInput `pulumi:"backendHostHeader"`
	EnabledState               pulumi.StringPtrInput `pulumi:"enabledState"`
	HttpPort                   pulumi.IntPtrInput    `pulumi:"httpPort"`
	HttpsPort                  pulumi.IntPtrInput    `pulumi:"httpsPort"`
	Priority                   pulumi.IntPtrInput    `pulumi:"priority"`
	PrivateLinkAlias           pulumi.StringPtrInput `pulumi:"privateLinkAlias"`
	PrivateLinkApprovalMessage pulumi.StringPtrInput `pulumi:"privateLinkApprovalMessage"`
	Weight                     pulumi.IntPtrInput    `pulumi:"weight"`
}

func (BackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Backend)(nil)).Elem()
}

func (i BackendArgs) ToBackendOutput() BackendOutput {
	return i.ToBackendOutputWithContext(context.Background())
}

func (i BackendArgs) ToBackendOutputWithContext(ctx context.Context) BackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendOutput)
}





type BackendArrayInput interface {
	pulumi.Input

	ToBackendArrayOutput() BackendArrayOutput
	ToBackendArrayOutputWithContext(context.Context) BackendArrayOutput
}

type BackendArray []BackendInput

func (BackendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Backend)(nil)).Elem()
}

func (i BackendArray) ToBackendArrayOutput() BackendArrayOutput {
	return i.ToBackendArrayOutputWithContext(context.Background())
}

func (i BackendArray) ToBackendArrayOutputWithContext(ctx context.Context) BackendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendArrayOutput)
}

type BackendOutput struct{ *pulumi.OutputState }

func (BackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Backend)(nil)).Elem()
}

func (o BackendOutput) ToBackendOutput() BackendOutput {
	return o
}

func (o BackendOutput) ToBackendOutputWithContext(ctx context.Context) BackendOutput {
	return o
}

func (o BackendOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Backend) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o BackendOutput) BackendHostHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Backend) *string { return v.BackendHostHeader }).(pulumi.StringPtrOutput)
}

func (o BackendOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Backend) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o BackendOutput) HttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Backend) *int { return v.HttpPort }).(pulumi.IntPtrOutput)
}

func (o BackendOutput) HttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Backend) *int { return v.HttpsPort }).(pulumi.IntPtrOutput)
}

func (o BackendOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Backend) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o BackendOutput) PrivateLinkAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Backend) *string { return v.PrivateLinkAlias }).(pulumi.StringPtrOutput)
}

func (o BackendOutput) PrivateLinkApprovalMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Backend) *string { return v.PrivateLinkApprovalMessage }).(pulumi.StringPtrOutput)
}

func (o BackendOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Backend) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type BackendArrayOutput struct{ *pulumi.OutputState }

func (BackendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Backend)(nil)).Elem()
}

func (o BackendArrayOutput) ToBackendArrayOutput() BackendArrayOutput {
	return o
}

func (o BackendArrayOutput) ToBackendArrayOutputWithContext(ctx context.Context) BackendArrayOutput {
	return o
}

func (o BackendArrayOutput) Index(i pulumi.IntInput) BackendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Backend {
		return vs[0].([]Backend)[vs[1].(int)]
	}).(BackendOutput)
}

type BackendPool struct {
	Backends              []Backend    `pulumi:"backends"`
	HealthProbeSettings   *SubResource `pulumi:"healthProbeSettings"`
	Id                    *string      `pulumi:"id"`
	LoadBalancingSettings *SubResource `pulumi:"loadBalancingSettings"`
	Name                  *string      `pulumi:"name"`
}





type BackendPoolInput interface {
	pulumi.Input

	ToBackendPoolOutput() BackendPoolOutput
	ToBackendPoolOutputWithContext(context.Context) BackendPoolOutput
}

type BackendPoolArgs struct {
	Backends              BackendArrayInput     `pulumi:"backends"`
	HealthProbeSettings   SubResourcePtrInput   `pulumi:"healthProbeSettings"`
	Id                    pulumi.StringPtrInput `pulumi:"id"`
	LoadBalancingSettings SubResourcePtrInput   `pulumi:"loadBalancingSettings"`
	Name                  pulumi.StringPtrInput `pulumi:"name"`
}

func (BackendPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendPool)(nil)).Elem()
}

func (i BackendPoolArgs) ToBackendPoolOutput() BackendPoolOutput {
	return i.ToBackendPoolOutputWithContext(context.Background())
}

func (i BackendPoolArgs) ToBackendPoolOutputWithContext(ctx context.Context) BackendPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendPoolOutput)
}





type BackendPoolArrayInput interface {
	pulumi.Input

	ToBackendPoolArrayOutput() BackendPoolArrayOutput
	ToBackendPoolArrayOutputWithContext(context.Context) BackendPoolArrayOutput
}

type BackendPoolArray []BackendPoolInput

func (BackendPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendPool)(nil)).Elem()
}

func (i BackendPoolArray) ToBackendPoolArrayOutput() BackendPoolArrayOutput {
	return i.ToBackendPoolArrayOutputWithContext(context.Background())
}

func (i BackendPoolArray) ToBackendPoolArrayOutputWithContext(ctx context.Context) BackendPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendPoolArrayOutput)
}

type BackendPoolOutput struct{ *pulumi.OutputState }

func (BackendPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendPool)(nil)).Elem()
}

func (o BackendPoolOutput) ToBackendPoolOutput() BackendPoolOutput {
	return o
}

func (o BackendPoolOutput) ToBackendPoolOutputWithContext(ctx context.Context) BackendPoolOutput {
	return o
}

func (o BackendPoolOutput) Backends() BackendArrayOutput {
	return o.ApplyT(func(v BackendPool) []Backend { return v.Backends }).(BackendArrayOutput)
}

func (o BackendPoolOutput) HealthProbeSettings() SubResourcePtrOutput {
	return o.ApplyT(func(v BackendPool) *SubResource { return v.HealthProbeSettings }).(SubResourcePtrOutput)
}

func (o BackendPoolOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendPool) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o BackendPoolOutput) LoadBalancingSettings() SubResourcePtrOutput {
	return o.ApplyT(func(v BackendPool) *SubResource { return v.LoadBalancingSettings }).(SubResourcePtrOutput)
}

func (o BackendPoolOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendPool) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type BackendPoolArrayOutput struct{ *pulumi.OutputState }

func (BackendPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendPool)(nil)).Elem()
}

func (o BackendPoolArrayOutput) ToBackendPoolArrayOutput() BackendPoolArrayOutput {
	return o
}

func (o BackendPoolArrayOutput) ToBackendPoolArrayOutputWithContext(ctx context.Context) BackendPoolArrayOutput {
	return o
}

func (o BackendPoolArrayOutput) Index(i pulumi.IntInput) BackendPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackendPool {
		return vs[0].([]BackendPool)[vs[1].(int)]
	}).(BackendPoolOutput)
}

type BackendPoolResponse struct {
	Backends              []BackendResponse    `pulumi:"backends"`
	HealthProbeSettings   *SubResourceResponse `pulumi:"healthProbeSettings"`
	Id                    *string              `pulumi:"id"`
	LoadBalancingSettings *SubResourceResponse `pulumi:"loadBalancingSettings"`
	Name                  *string              `pulumi:"name"`
	ResourceState         string               `pulumi:"resourceState"`
	Type                  string               `pulumi:"type"`
}

type BackendPoolResponseOutput struct{ *pulumi.OutputState }

func (BackendPoolResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendPoolResponse)(nil)).Elem()
}

func (o BackendPoolResponseOutput) ToBackendPoolResponseOutput() BackendPoolResponseOutput {
	return o
}

func (o BackendPoolResponseOutput) ToBackendPoolResponseOutputWithContext(ctx context.Context) BackendPoolResponseOutput {
	return o
}

func (o BackendPoolResponseOutput) Backends() BackendResponseArrayOutput {
	return o.ApplyT(func(v BackendPoolResponse) []BackendResponse { return v.Backends }).(BackendResponseArrayOutput)
}

func (o BackendPoolResponseOutput) HealthProbeSettings() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v BackendPoolResponse) *SubResourceResponse { return v.HealthProbeSettings }).(SubResourceResponsePtrOutput)
}

func (o BackendPoolResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendPoolResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o BackendPoolResponseOutput) LoadBalancingSettings() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v BackendPoolResponse) *SubResourceResponse { return v.LoadBalancingSettings }).(SubResourceResponsePtrOutput)
}

func (o BackendPoolResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendPoolResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o BackendPoolResponseOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v BackendPoolResponse) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o BackendPoolResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v BackendPoolResponse) string { return v.Type }).(pulumi.StringOutput)
}

type BackendPoolResponseArrayOutput struct{ *pulumi.OutputState }

func (BackendPoolResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendPoolResponse)(nil)).Elem()
}

func (o BackendPoolResponseArrayOutput) ToBackendPoolResponseArrayOutput() BackendPoolResponseArrayOutput {
	return o
}

func (o BackendPoolResponseArrayOutput) ToBackendPoolResponseArrayOutputWithContext(ctx context.Context) BackendPoolResponseArrayOutput {
	return o
}

func (o BackendPoolResponseArrayOutput) Index(i pulumi.IntInput) BackendPoolResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackendPoolResponse {
		return vs[0].([]BackendPoolResponse)[vs[1].(int)]
	}).(BackendPoolResponseOutput)
}

type BackendPoolsSettings struct {
	EnforceCertificateNameCheck *string `pulumi:"enforceCertificateNameCheck"`
	SendRecvTimeoutSeconds      *int    `pulumi:"sendRecvTimeoutSeconds"`
}


func (val *BackendPoolsSettings) Defaults() *BackendPoolsSettings {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnforceCertificateNameCheck) {
		enforceCertificateNameCheck_ := "Enabled"
		tmp.EnforceCertificateNameCheck = &enforceCertificateNameCheck_
	}
	return &tmp
}





type BackendPoolsSettingsInput interface {
	pulumi.Input

	ToBackendPoolsSettingsOutput() BackendPoolsSettingsOutput
	ToBackendPoolsSettingsOutputWithContext(context.Context) BackendPoolsSettingsOutput
}

type BackendPoolsSettingsArgs struct {
	EnforceCertificateNameCheck pulumi.StringPtrInput `pulumi:"enforceCertificateNameCheck"`
	SendRecvTimeoutSeconds      pulumi.IntPtrInput    `pulumi:"sendRecvTimeoutSeconds"`
}


func (val *BackendPoolsSettingsArgs) Defaults() *BackendPoolsSettingsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnforceCertificateNameCheck) {
		tmp.EnforceCertificateNameCheck = pulumi.StringPtr("Enabled")
	}
	return &tmp
}
func (BackendPoolsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendPoolsSettings)(nil)).Elem()
}

func (i BackendPoolsSettingsArgs) ToBackendPoolsSettingsOutput() BackendPoolsSettingsOutput {
	return i.ToBackendPoolsSettingsOutputWithContext(context.Background())
}

func (i BackendPoolsSettingsArgs) ToBackendPoolsSettingsOutputWithContext(ctx context.Context) BackendPoolsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendPoolsSettingsOutput)
}

func (i BackendPoolsSettingsArgs) ToBackendPoolsSettingsPtrOutput() BackendPoolsSettingsPtrOutput {
	return i.ToBackendPoolsSettingsPtrOutputWithContext(context.Background())
}

func (i BackendPoolsSettingsArgs) ToBackendPoolsSettingsPtrOutputWithContext(ctx context.Context) BackendPoolsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendPoolsSettingsOutput).ToBackendPoolsSettingsPtrOutputWithContext(ctx)
}









type BackendPoolsSettingsPtrInput interface {
	pulumi.Input

	ToBackendPoolsSettingsPtrOutput() BackendPoolsSettingsPtrOutput
	ToBackendPoolsSettingsPtrOutputWithContext(context.Context) BackendPoolsSettingsPtrOutput
}

type backendPoolsSettingsPtrType BackendPoolsSettingsArgs

func BackendPoolsSettingsPtr(v *BackendPoolsSettingsArgs) BackendPoolsSettingsPtrInput {
	return (*backendPoolsSettingsPtrType)(v)
}

func (*backendPoolsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendPoolsSettings)(nil)).Elem()
}

func (i *backendPoolsSettingsPtrType) ToBackendPoolsSettingsPtrOutput() BackendPoolsSettingsPtrOutput {
	return i.ToBackendPoolsSettingsPtrOutputWithContext(context.Background())
}

func (i *backendPoolsSettingsPtrType) ToBackendPoolsSettingsPtrOutputWithContext(ctx context.Context) BackendPoolsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendPoolsSettingsPtrOutput)
}

type BackendPoolsSettingsOutput struct{ *pulumi.OutputState }

func (BackendPoolsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendPoolsSettings)(nil)).Elem()
}

func (o BackendPoolsSettingsOutput) ToBackendPoolsSettingsOutput() BackendPoolsSettingsOutput {
	return o
}

func (o BackendPoolsSettingsOutput) ToBackendPoolsSettingsOutputWithContext(ctx context.Context) BackendPoolsSettingsOutput {
	return o
}

func (o BackendPoolsSettingsOutput) ToBackendPoolsSettingsPtrOutput() BackendPoolsSettingsPtrOutput {
	return o.ToBackendPoolsSettingsPtrOutputWithContext(context.Background())
}

func (o BackendPoolsSettingsOutput) ToBackendPoolsSettingsPtrOutputWithContext(ctx context.Context) BackendPoolsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackendPoolsSettings) *BackendPoolsSettings {
		return &v
	}).(BackendPoolsSettingsPtrOutput)
}

func (o BackendPoolsSettingsOutput) EnforceCertificateNameCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendPoolsSettings) *string { return v.EnforceCertificateNameCheck }).(pulumi.StringPtrOutput)
}

func (o BackendPoolsSettingsOutput) SendRecvTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackendPoolsSettings) *int { return v.SendRecvTimeoutSeconds }).(pulumi.IntPtrOutput)
}

type BackendPoolsSettingsPtrOutput struct{ *pulumi.OutputState }

func (BackendPoolsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendPoolsSettings)(nil)).Elem()
}

func (o BackendPoolsSettingsPtrOutput) ToBackendPoolsSettingsPtrOutput() BackendPoolsSettingsPtrOutput {
	return o
}

func (o BackendPoolsSettingsPtrOutput) ToBackendPoolsSettingsPtrOutputWithContext(ctx context.Context) BackendPoolsSettingsPtrOutput {
	return o
}

func (o BackendPoolsSettingsPtrOutput) Elem() BackendPoolsSettingsOutput {
	return o.ApplyT(func(v *BackendPoolsSettings) BackendPoolsSettings {
		if v != nil {
			return *v
		}
		var ret BackendPoolsSettings
		return ret
	}).(BackendPoolsSettingsOutput)
}

func (o BackendPoolsSettingsPtrOutput) EnforceCertificateNameCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendPoolsSettings) *string {
		if v == nil {
			return nil
		}
		return v.EnforceCertificateNameCheck
	}).(pulumi.StringPtrOutput)
}

func (o BackendPoolsSettingsPtrOutput) SendRecvTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackendPoolsSettings) *int {
		if v == nil {
			return nil
		}
		return v.SendRecvTimeoutSeconds
	}).(pulumi.IntPtrOutput)
}

type BackendPoolsSettingsResponse struct {
	EnforceCertificateNameCheck *string `pulumi:"enforceCertificateNameCheck"`
	SendRecvTimeoutSeconds      *int    `pulumi:"sendRecvTimeoutSeconds"`
}


func (val *BackendPoolsSettingsResponse) Defaults() *BackendPoolsSettingsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnforceCertificateNameCheck) {
		enforceCertificateNameCheck_ := "Enabled"
		tmp.EnforceCertificateNameCheck = &enforceCertificateNameCheck_
	}
	return &tmp
}

type BackendPoolsSettingsResponseOutput struct{ *pulumi.OutputState }

func (BackendPoolsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendPoolsSettingsResponse)(nil)).Elem()
}

func (o BackendPoolsSettingsResponseOutput) ToBackendPoolsSettingsResponseOutput() BackendPoolsSettingsResponseOutput {
	return o
}

func (o BackendPoolsSettingsResponseOutput) ToBackendPoolsSettingsResponseOutputWithContext(ctx context.Context) BackendPoolsSettingsResponseOutput {
	return o
}

func (o BackendPoolsSettingsResponseOutput) EnforceCertificateNameCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendPoolsSettingsResponse) *string { return v.EnforceCertificateNameCheck }).(pulumi.StringPtrOutput)
}

func (o BackendPoolsSettingsResponseOutput) SendRecvTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackendPoolsSettingsResponse) *int { return v.SendRecvTimeoutSeconds }).(pulumi.IntPtrOutput)
}

type BackendPoolsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (BackendPoolsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackendPoolsSettingsResponse)(nil)).Elem()
}

func (o BackendPoolsSettingsResponsePtrOutput) ToBackendPoolsSettingsResponsePtrOutput() BackendPoolsSettingsResponsePtrOutput {
	return o
}

func (o BackendPoolsSettingsResponsePtrOutput) ToBackendPoolsSettingsResponsePtrOutputWithContext(ctx context.Context) BackendPoolsSettingsResponsePtrOutput {
	return o
}

func (o BackendPoolsSettingsResponsePtrOutput) Elem() BackendPoolsSettingsResponseOutput {
	return o.ApplyT(func(v *BackendPoolsSettingsResponse) BackendPoolsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret BackendPoolsSettingsResponse
		return ret
	}).(BackendPoolsSettingsResponseOutput)
}

func (o BackendPoolsSettingsResponsePtrOutput) EnforceCertificateNameCheck() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackendPoolsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.EnforceCertificateNameCheck
	}).(pulumi.StringPtrOutput)
}

func (o BackendPoolsSettingsResponsePtrOutput) SendRecvTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackendPoolsSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.SendRecvTimeoutSeconds
	}).(pulumi.IntPtrOutput)
}

type BackendResponse struct {
	Address                    *string `pulumi:"address"`
	BackendHostHeader          *string `pulumi:"backendHostHeader"`
	EnabledState               *string `pulumi:"enabledState"`
	HttpPort                   *int    `pulumi:"httpPort"`
	HttpsPort                  *int    `pulumi:"httpsPort"`
	Priority                   *int    `pulumi:"priority"`
	PrivateEndpointStatus      string  `pulumi:"privateEndpointStatus"`
	PrivateLinkAlias           *string `pulumi:"privateLinkAlias"`
	PrivateLinkApprovalMessage *string `pulumi:"privateLinkApprovalMessage"`
	Weight                     *int    `pulumi:"weight"`
}

type BackendResponseOutput struct{ *pulumi.OutputState }

func (BackendResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendResponse)(nil)).Elem()
}

func (o BackendResponseOutput) ToBackendResponseOutput() BackendResponseOutput {
	return o
}

func (o BackendResponseOutput) ToBackendResponseOutputWithContext(ctx context.Context) BackendResponseOutput {
	return o
}

func (o BackendResponseOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendResponse) *string { return v.Address }).(pulumi.StringPtrOutput)
}

func (o BackendResponseOutput) BackendHostHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendResponse) *string { return v.BackendHostHeader }).(pulumi.StringPtrOutput)
}

func (o BackendResponseOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendResponse) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o BackendResponseOutput) HttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackendResponse) *int { return v.HttpPort }).(pulumi.IntPtrOutput)
}

func (o BackendResponseOutput) HttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackendResponse) *int { return v.HttpsPort }).(pulumi.IntPtrOutput)
}

func (o BackendResponseOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackendResponse) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o BackendResponseOutput) PrivateEndpointStatus() pulumi.StringOutput {
	return o.ApplyT(func(v BackendResponse) string { return v.PrivateEndpointStatus }).(pulumi.StringOutput)
}

func (o BackendResponseOutput) PrivateLinkAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendResponse) *string { return v.PrivateLinkAlias }).(pulumi.StringPtrOutput)
}

func (o BackendResponseOutput) PrivateLinkApprovalMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendResponse) *string { return v.PrivateLinkApprovalMessage }).(pulumi.StringPtrOutput)
}

func (o BackendResponseOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackendResponse) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type BackendResponseArrayOutput struct{ *pulumi.OutputState }

func (BackendResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendResponse)(nil)).Elem()
}

func (o BackendResponseArrayOutput) ToBackendResponseArrayOutput() BackendResponseArrayOutput {
	return o
}

func (o BackendResponseArrayOutput) ToBackendResponseArrayOutputWithContext(ctx context.Context) BackendResponseArrayOutput {
	return o
}

func (o BackendResponseArrayOutput) Index(i pulumi.IntInput) BackendResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackendResponse {
		return vs[0].([]BackendResponse)[vs[1].(int)]
	}).(BackendResponseOutput)
}

type CacheConfiguration struct {
	CacheDuration                *string `pulumi:"cacheDuration"`
	DynamicCompression           *string `pulumi:"dynamicCompression"`
	QueryParameterStripDirective *string `pulumi:"queryParameterStripDirective"`
	QueryParameters              *string `pulumi:"queryParameters"`
}

type CacheConfigurationResponse struct {
	CacheDuration                *string `pulumi:"cacheDuration"`
	DynamicCompression           *string `pulumi:"dynamicCompression"`
	QueryParameterStripDirective *string `pulumi:"queryParameterStripDirective"`
	QueryParameters              *string `pulumi:"queryParameters"`
}

type CnameRecord struct {
	Cname *string `pulumi:"cname"`
}





type CnameRecordInput interface {
	pulumi.Input

	ToCnameRecordOutput() CnameRecordOutput
	ToCnameRecordOutputWithContext(context.Context) CnameRecordOutput
}

type CnameRecordArgs struct {
	Cname pulumi.StringPtrInput `pulumi:"cname"`
}

func (CnameRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CnameRecord)(nil)).Elem()
}

func (i CnameRecordArgs) ToCnameRecordOutput() CnameRecordOutput {
	return i.ToCnameRecordOutputWithContext(context.Background())
}

func (i CnameRecordArgs) ToCnameRecordOutputWithContext(ctx context.Context) CnameRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CnameRecordOutput)
}

func (i CnameRecordArgs) ToCnameRecordPtrOutput() CnameRecordPtrOutput {
	return i.ToCnameRecordPtrOutputWithContext(context.Background())
}

func (i CnameRecordArgs) ToCnameRecordPtrOutputWithContext(ctx context.Context) CnameRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CnameRecordOutput).ToCnameRecordPtrOutputWithContext(ctx)
}









type CnameRecordPtrInput interface {
	pulumi.Input

	ToCnameRecordPtrOutput() CnameRecordPtrOutput
	ToCnameRecordPtrOutputWithContext(context.Context) CnameRecordPtrOutput
}

type cnameRecordPtrType CnameRecordArgs

func CnameRecordPtr(v *CnameRecordArgs) CnameRecordPtrInput {
	return (*cnameRecordPtrType)(v)
}

func (*cnameRecordPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CnameRecord)(nil)).Elem()
}

func (i *cnameRecordPtrType) ToCnameRecordPtrOutput() CnameRecordPtrOutput {
	return i.ToCnameRecordPtrOutputWithContext(context.Background())
}

func (i *cnameRecordPtrType) ToCnameRecordPtrOutputWithContext(ctx context.Context) CnameRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CnameRecordPtrOutput)
}

type CnameRecordOutput struct{ *pulumi.OutputState }

func (CnameRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CnameRecord)(nil)).Elem()
}

func (o CnameRecordOutput) ToCnameRecordOutput() CnameRecordOutput {
	return o
}

func (o CnameRecordOutput) ToCnameRecordOutputWithContext(ctx context.Context) CnameRecordOutput {
	return o
}

func (o CnameRecordOutput) ToCnameRecordPtrOutput() CnameRecordPtrOutput {
	return o.ToCnameRecordPtrOutputWithContext(context.Background())
}

func (o CnameRecordOutput) ToCnameRecordPtrOutputWithContext(ctx context.Context) CnameRecordPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CnameRecord) *CnameRecord {
		return &v
	}).(CnameRecordPtrOutput)
}

func (o CnameRecordOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CnameRecord) *string { return v.Cname }).(pulumi.StringPtrOutput)
}

type CnameRecordPtrOutput struct{ *pulumi.OutputState }

func (CnameRecordPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CnameRecord)(nil)).Elem()
}

func (o CnameRecordPtrOutput) ToCnameRecordPtrOutput() CnameRecordPtrOutput {
	return o
}

func (o CnameRecordPtrOutput) ToCnameRecordPtrOutputWithContext(ctx context.Context) CnameRecordPtrOutput {
	return o
}

func (o CnameRecordPtrOutput) Elem() CnameRecordOutput {
	return o.ApplyT(func(v *CnameRecord) CnameRecord {
		if v != nil {
			return *v
		}
		var ret CnameRecord
		return ret
	}).(CnameRecordOutput)
}

func (o CnameRecordPtrOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CnameRecord) *string {
		if v == nil {
			return nil
		}
		return v.Cname
	}).(pulumi.StringPtrOutput)
}

type CnameRecordResponse struct {
	Cname *string `pulumi:"cname"`
}

type CnameRecordResponseOutput struct{ *pulumi.OutputState }

func (CnameRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CnameRecordResponse)(nil)).Elem()
}

func (o CnameRecordResponseOutput) ToCnameRecordResponseOutput() CnameRecordResponseOutput {
	return o
}

func (o CnameRecordResponseOutput) ToCnameRecordResponseOutputWithContext(ctx context.Context) CnameRecordResponseOutput {
	return o
}

func (o CnameRecordResponseOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CnameRecordResponse) *string { return v.Cname }).(pulumi.StringPtrOutput)
}

type CnameRecordResponsePtrOutput struct{ *pulumi.OutputState }

func (CnameRecordResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CnameRecordResponse)(nil)).Elem()
}

func (o CnameRecordResponsePtrOutput) ToCnameRecordResponsePtrOutput() CnameRecordResponsePtrOutput {
	return o
}

func (o CnameRecordResponsePtrOutput) ToCnameRecordResponsePtrOutputWithContext(ctx context.Context) CnameRecordResponsePtrOutput {
	return o
}

func (o CnameRecordResponsePtrOutput) Elem() CnameRecordResponseOutput {
	return o.ApplyT(func(v *CnameRecordResponse) CnameRecordResponse {
		if v != nil {
			return *v
		}
		var ret CnameRecordResponse
		return ret
	}).(CnameRecordResponseOutput)
}

func (o CnameRecordResponsePtrOutput) Cname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CnameRecordResponse) *string {
		if v == nil {
			return nil
		}
		return v.Cname
	}).(pulumi.StringPtrOutput)
}

type CustomHttpsConfigurationResponse struct {
	CertificateSource string                                            `pulumi:"certificateSource"`
	CertificateType   *string                                           `pulumi:"certificateType"`
	MinimumTlsVersion string                                            `pulumi:"minimumTlsVersion"`
	ProtocolType      string                                            `pulumi:"protocolType"`
	SecretName        *string                                           `pulumi:"secretName"`
	SecretVersion     *string                                           `pulumi:"secretVersion"`
	Vault             *KeyVaultCertificateSourceParametersResponseVault `pulumi:"vault"`
}

type CustomHttpsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (CustomHttpsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomHttpsConfigurationResponse)(nil)).Elem()
}

func (o CustomHttpsConfigurationResponseOutput) ToCustomHttpsConfigurationResponseOutput() CustomHttpsConfigurationResponseOutput {
	return o
}

func (o CustomHttpsConfigurationResponseOutput) ToCustomHttpsConfigurationResponseOutputWithContext(ctx context.Context) CustomHttpsConfigurationResponseOutput {
	return o
}

func (o CustomHttpsConfigurationResponseOutput) CertificateSource() pulumi.StringOutput {
	return o.ApplyT(func(v CustomHttpsConfigurationResponse) string { return v.CertificateSource }).(pulumi.StringOutput)
}

func (o CustomHttpsConfigurationResponseOutput) CertificateType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomHttpsConfigurationResponse) *string { return v.CertificateType }).(pulumi.StringPtrOutput)
}

func (o CustomHttpsConfigurationResponseOutput) MinimumTlsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v CustomHttpsConfigurationResponse) string { return v.MinimumTlsVersion }).(pulumi.StringOutput)
}

func (o CustomHttpsConfigurationResponseOutput) ProtocolType() pulumi.StringOutput {
	return o.ApplyT(func(v CustomHttpsConfigurationResponse) string { return v.ProtocolType }).(pulumi.StringOutput)
}

func (o CustomHttpsConfigurationResponseOutput) SecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomHttpsConfigurationResponse) *string { return v.SecretName }).(pulumi.StringPtrOutput)
}

func (o CustomHttpsConfigurationResponseOutput) SecretVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomHttpsConfigurationResponse) *string { return v.SecretVersion }).(pulumi.StringPtrOutput)
}

func (o CustomHttpsConfigurationResponseOutput) Vault() KeyVaultCertificateSourceParametersResponseVaultPtrOutput {
	return o.ApplyT(func(v CustomHttpsConfigurationResponse) *KeyVaultCertificateSourceParametersResponseVault {
		return v.Vault
	}).(KeyVaultCertificateSourceParametersResponseVaultPtrOutput)
}

type ForwardingConfiguration struct {
	BackendPool          *SubResource        `pulumi:"backendPool"`
	CacheConfiguration   *CacheConfiguration `pulumi:"cacheConfiguration"`
	CustomForwardingPath *string             `pulumi:"customForwardingPath"`
	ForwardingProtocol   *string             `pulumi:"forwardingProtocol"`
	OdataType            string              `pulumi:"odataType"`
}

type ForwardingConfigurationResponse struct {
	BackendPool          *SubResourceResponse        `pulumi:"backendPool"`
	CacheConfiguration   *CacheConfigurationResponse `pulumi:"cacheConfiguration"`
	CustomForwardingPath *string                     `pulumi:"customForwardingPath"`
	ForwardingProtocol   *string                     `pulumi:"forwardingProtocol"`
	OdataType            string                      `pulumi:"odataType"`
}

type FrontendEndpoint struct {
	HostName                         *string                                                           `pulumi:"hostName"`
	Id                               *string                                                           `pulumi:"id"`
	Name                             *string                                                           `pulumi:"name"`
	SessionAffinityEnabledState      *string                                                           `pulumi:"sessionAffinityEnabledState"`
	SessionAffinityTtlSeconds        *int                                                              `pulumi:"sessionAffinityTtlSeconds"`
	WebApplicationFirewallPolicyLink *FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink `pulumi:"webApplicationFirewallPolicyLink"`
}





type FrontendEndpointInput interface {
	pulumi.Input

	ToFrontendEndpointOutput() FrontendEndpointOutput
	ToFrontendEndpointOutputWithContext(context.Context) FrontendEndpointOutput
}

type FrontendEndpointArgs struct {
	HostName                         pulumi.StringPtrInput                                                    `pulumi:"hostName"`
	Id                               pulumi.StringPtrInput                                                    `pulumi:"id"`
	Name                             pulumi.StringPtrInput                                                    `pulumi:"name"`
	SessionAffinityEnabledState      pulumi.StringPtrInput                                                    `pulumi:"sessionAffinityEnabledState"`
	SessionAffinityTtlSeconds        pulumi.IntPtrInput                                                       `pulumi:"sessionAffinityTtlSeconds"`
	WebApplicationFirewallPolicyLink FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrInput `pulumi:"webApplicationFirewallPolicyLink"`
}

func (FrontendEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendEndpoint)(nil)).Elem()
}

func (i FrontendEndpointArgs) ToFrontendEndpointOutput() FrontendEndpointOutput {
	return i.ToFrontendEndpointOutputWithContext(context.Background())
}

func (i FrontendEndpointArgs) ToFrontendEndpointOutputWithContext(ctx context.Context) FrontendEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendEndpointOutput)
}





type FrontendEndpointArrayInput interface {
	pulumi.Input

	ToFrontendEndpointArrayOutput() FrontendEndpointArrayOutput
	ToFrontendEndpointArrayOutputWithContext(context.Context) FrontendEndpointArrayOutput
}

type FrontendEndpointArray []FrontendEndpointInput

func (FrontendEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendEndpoint)(nil)).Elem()
}

func (i FrontendEndpointArray) ToFrontendEndpointArrayOutput() FrontendEndpointArrayOutput {
	return i.ToFrontendEndpointArrayOutputWithContext(context.Background())
}

func (i FrontendEndpointArray) ToFrontendEndpointArrayOutputWithContext(ctx context.Context) FrontendEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendEndpointArrayOutput)
}

type FrontendEndpointOutput struct{ *pulumi.OutputState }

func (FrontendEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendEndpoint)(nil)).Elem()
}

func (o FrontendEndpointOutput) ToFrontendEndpointOutput() FrontendEndpointOutput {
	return o
}

func (o FrontendEndpointOutput) ToFrontendEndpointOutputWithContext(ctx context.Context) FrontendEndpointOutput {
	return o
}

func (o FrontendEndpointOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpoint) *string { return v.HostName }).(pulumi.StringPtrOutput)
}

func (o FrontendEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o FrontendEndpointOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpoint) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FrontendEndpointOutput) SessionAffinityEnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpoint) *string { return v.SessionAffinityEnabledState }).(pulumi.StringPtrOutput)
}

func (o FrontendEndpointOutput) SessionAffinityTtlSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FrontendEndpoint) *int { return v.SessionAffinityTtlSeconds }).(pulumi.IntPtrOutput)
}

func (o FrontendEndpointOutput) WebApplicationFirewallPolicyLink() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ApplyT(func(v FrontendEndpoint) *FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink {
		return v.WebApplicationFirewallPolicyLink
	}).(FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput)
}

type FrontendEndpointArrayOutput struct{ *pulumi.OutputState }

func (FrontendEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendEndpoint)(nil)).Elem()
}

func (o FrontendEndpointArrayOutput) ToFrontendEndpointArrayOutput() FrontendEndpointArrayOutput {
	return o
}

func (o FrontendEndpointArrayOutput) ToFrontendEndpointArrayOutputWithContext(ctx context.Context) FrontendEndpointArrayOutput {
	return o
}

func (o FrontendEndpointArrayOutput) Index(i pulumi.IntInput) FrontendEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontendEndpoint {
		return vs[0].([]FrontendEndpoint)[vs[1].(int)]
	}).(FrontendEndpointOutput)
}

type FrontendEndpointResponse struct {
	CustomHttpsConfiguration         CustomHttpsConfigurationResponse                                          `pulumi:"customHttpsConfiguration"`
	CustomHttpsProvisioningState     string                                                                    `pulumi:"customHttpsProvisioningState"`
	CustomHttpsProvisioningSubstate  string                                                                    `pulumi:"customHttpsProvisioningSubstate"`
	HostName                         *string                                                                   `pulumi:"hostName"`
	Id                               *string                                                                   `pulumi:"id"`
	Name                             *string                                                                   `pulumi:"name"`
	ResourceState                    string                                                                    `pulumi:"resourceState"`
	SessionAffinityEnabledState      *string                                                                   `pulumi:"sessionAffinityEnabledState"`
	SessionAffinityTtlSeconds        *int                                                                      `pulumi:"sessionAffinityTtlSeconds"`
	Type                             string                                                                    `pulumi:"type"`
	WebApplicationFirewallPolicyLink *FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink `pulumi:"webApplicationFirewallPolicyLink"`
}

type FrontendEndpointResponseOutput struct{ *pulumi.OutputState }

func (FrontendEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendEndpointResponse)(nil)).Elem()
}

func (o FrontendEndpointResponseOutput) ToFrontendEndpointResponseOutput() FrontendEndpointResponseOutput {
	return o
}

func (o FrontendEndpointResponseOutput) ToFrontendEndpointResponseOutputWithContext(ctx context.Context) FrontendEndpointResponseOutput {
	return o
}

func (o FrontendEndpointResponseOutput) CustomHttpsConfiguration() CustomHttpsConfigurationResponseOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) CustomHttpsConfigurationResponse { return v.CustomHttpsConfiguration }).(CustomHttpsConfigurationResponseOutput)
}

func (o FrontendEndpointResponseOutput) CustomHttpsProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) string { return v.CustomHttpsProvisioningState }).(pulumi.StringOutput)
}

func (o FrontendEndpointResponseOutput) CustomHttpsProvisioningSubstate() pulumi.StringOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) string { return v.CustomHttpsProvisioningSubstate }).(pulumi.StringOutput)
}

func (o FrontendEndpointResponseOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) *string { return v.HostName }).(pulumi.StringPtrOutput)
}

func (o FrontendEndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o FrontendEndpointResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FrontendEndpointResponseOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o FrontendEndpointResponseOutput) SessionAffinityEnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) *string { return v.SessionAffinityEnabledState }).(pulumi.StringPtrOutput)
}

func (o FrontendEndpointResponseOutput) SessionAffinityTtlSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) *int { return v.SessionAffinityTtlSeconds }).(pulumi.IntPtrOutput)
}

func (o FrontendEndpointResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o FrontendEndpointResponseOutput) WebApplicationFirewallPolicyLink() FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ApplyT(func(v FrontendEndpointResponse) *FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink {
		return v.WebApplicationFirewallPolicyLink
	}).(FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput)
}

type FrontendEndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (FrontendEndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendEndpointResponse)(nil)).Elem()
}

func (o FrontendEndpointResponseArrayOutput) ToFrontendEndpointResponseArrayOutput() FrontendEndpointResponseArrayOutput {
	return o
}

func (o FrontendEndpointResponseArrayOutput) ToFrontendEndpointResponseArrayOutputWithContext(ctx context.Context) FrontendEndpointResponseArrayOutput {
	return o
}

func (o FrontendEndpointResponseArrayOutput) Index(i pulumi.IntInput) FrontendEndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontendEndpointResponse {
		return vs[0].([]FrontendEndpointResponse)[vs[1].(int)]
	}).(FrontendEndpointResponseOutput)
}

type FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink struct {
	Id *string `pulumi:"id"`
}

type FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput struct{ *pulumi.OutputState }

func (FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) ToFrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput() FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) ToFrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutputWithContext(ctx context.Context) FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput struct{ *pulumi.OutputState }

func (FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) ToFrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput() FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) ToFrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) Elem() FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput {
	return o.ApplyT(func(v *FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink) FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink {
		if v != nil {
			return *v
		}
		var ret FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink
		return ret
	}).(FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput)
}

func (o FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLink) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink struct {
	Id *string `pulumi:"id"`
}





type FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkInput interface {
	pulumi.Input

	ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput
	ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(context.Context) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput
}

type FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (i FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return i.ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(context.Background())
}

func (i FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(ctx context.Context) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput)
}

func (i FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return i.ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Background())
}

func (i FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkArgs) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput).ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx)
}









type FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrInput interface {
	pulumi.Input

	ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput
	ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Context) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput
}

type frontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrType FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkArgs

func FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtr(v *FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkArgs) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrInput {
	return (*frontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrType)(v)
}

func (*frontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (i *frontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrType) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return i.ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Background())
}

func (i *frontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrType) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput)
}

type FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput struct{ *pulumi.OutputState }

func (FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutputWithContext(ctx context.Context) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return o
}

func (o FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(context.Background())
}

func (o FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink) *FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink {
		return &v
	}).(FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput)
}

func (o FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput struct{ *pulumi.OutputState }

func (FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink)(nil)).Elem()
}

func (o FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) ToFrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutputWithContext(ctx context.Context) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput {
	return o
}

func (o FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) Elem() FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput {
	return o.ApplyT(func(v *FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink) FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink {
		if v != nil {
			return *v
		}
		var ret FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink
		return ret
	}).(FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput)
}

func (o FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type HeaderAction struct {
	HeaderActionType string  `pulumi:"headerActionType"`
	HeaderName       string  `pulumi:"headerName"`
	Value            *string `pulumi:"value"`
}





type HeaderActionInput interface {
	pulumi.Input

	ToHeaderActionOutput() HeaderActionOutput
	ToHeaderActionOutputWithContext(context.Context) HeaderActionOutput
}

type HeaderActionArgs struct {
	HeaderActionType pulumi.StringInput    `pulumi:"headerActionType"`
	HeaderName       pulumi.StringInput    `pulumi:"headerName"`
	Value            pulumi.StringPtrInput `pulumi:"value"`
}

func (HeaderActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HeaderAction)(nil)).Elem()
}

func (i HeaderActionArgs) ToHeaderActionOutput() HeaderActionOutput {
	return i.ToHeaderActionOutputWithContext(context.Background())
}

func (i HeaderActionArgs) ToHeaderActionOutputWithContext(ctx context.Context) HeaderActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HeaderActionOutput)
}





type HeaderActionArrayInput interface {
	pulumi.Input

	ToHeaderActionArrayOutput() HeaderActionArrayOutput
	ToHeaderActionArrayOutputWithContext(context.Context) HeaderActionArrayOutput
}

type HeaderActionArray []HeaderActionInput

func (HeaderActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HeaderAction)(nil)).Elem()
}

func (i HeaderActionArray) ToHeaderActionArrayOutput() HeaderActionArrayOutput {
	return i.ToHeaderActionArrayOutputWithContext(context.Background())
}

func (i HeaderActionArray) ToHeaderActionArrayOutputWithContext(ctx context.Context) HeaderActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HeaderActionArrayOutput)
}

type HeaderActionOutput struct{ *pulumi.OutputState }

func (HeaderActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HeaderAction)(nil)).Elem()
}

func (o HeaderActionOutput) ToHeaderActionOutput() HeaderActionOutput {
	return o
}

func (o HeaderActionOutput) ToHeaderActionOutputWithContext(ctx context.Context) HeaderActionOutput {
	return o
}

func (o HeaderActionOutput) HeaderActionType() pulumi.StringOutput {
	return o.ApplyT(func(v HeaderAction) string { return v.HeaderActionType }).(pulumi.StringOutput)
}

func (o HeaderActionOutput) HeaderName() pulumi.StringOutput {
	return o.ApplyT(func(v HeaderAction) string { return v.HeaderName }).(pulumi.StringOutput)
}

func (o HeaderActionOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HeaderAction) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type HeaderActionArrayOutput struct{ *pulumi.OutputState }

func (HeaderActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HeaderAction)(nil)).Elem()
}

func (o HeaderActionArrayOutput) ToHeaderActionArrayOutput() HeaderActionArrayOutput {
	return o
}

func (o HeaderActionArrayOutput) ToHeaderActionArrayOutputWithContext(ctx context.Context) HeaderActionArrayOutput {
	return o
}

func (o HeaderActionArrayOutput) Index(i pulumi.IntInput) HeaderActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HeaderAction {
		return vs[0].([]HeaderAction)[vs[1].(int)]
	}).(HeaderActionOutput)
}

type HeaderActionResponse struct {
	HeaderActionType string  `pulumi:"headerActionType"`
	HeaderName       string  `pulumi:"headerName"`
	Value            *string `pulumi:"value"`
}

type HeaderActionResponseOutput struct{ *pulumi.OutputState }

func (HeaderActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HeaderActionResponse)(nil)).Elem()
}

func (o HeaderActionResponseOutput) ToHeaderActionResponseOutput() HeaderActionResponseOutput {
	return o
}

func (o HeaderActionResponseOutput) ToHeaderActionResponseOutputWithContext(ctx context.Context) HeaderActionResponseOutput {
	return o
}

func (o HeaderActionResponseOutput) HeaderActionType() pulumi.StringOutput {
	return o.ApplyT(func(v HeaderActionResponse) string { return v.HeaderActionType }).(pulumi.StringOutput)
}

func (o HeaderActionResponseOutput) HeaderName() pulumi.StringOutput {
	return o.ApplyT(func(v HeaderActionResponse) string { return v.HeaderName }).(pulumi.StringOutput)
}

func (o HeaderActionResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HeaderActionResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type HeaderActionResponseArrayOutput struct{ *pulumi.OutputState }

func (HeaderActionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HeaderActionResponse)(nil)).Elem()
}

func (o HeaderActionResponseArrayOutput) ToHeaderActionResponseArrayOutput() HeaderActionResponseArrayOutput {
	return o
}

func (o HeaderActionResponseArrayOutput) ToHeaderActionResponseArrayOutputWithContext(ctx context.Context) HeaderActionResponseArrayOutput {
	return o
}

func (o HeaderActionResponseArrayOutput) Index(i pulumi.IntInput) HeaderActionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HeaderActionResponse {
		return vs[0].([]HeaderActionResponse)[vs[1].(int)]
	}).(HeaderActionResponseOutput)
}

type HealthProbeSettingsModel struct {
	EnabledState      *string `pulumi:"enabledState"`
	HealthProbeMethod *string `pulumi:"healthProbeMethod"`
	Id                *string `pulumi:"id"`
	IntervalInSeconds *int    `pulumi:"intervalInSeconds"`
	Name              *string `pulumi:"name"`
	Path              *string `pulumi:"path"`
	Protocol          *string `pulumi:"protocol"`
}


func (val *HealthProbeSettingsModel) Defaults() *HealthProbeSettingsModel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HealthProbeMethod) {
		healthProbeMethod_ := "HEAD"
		tmp.HealthProbeMethod = &healthProbeMethod_
	}
	return &tmp
}





type HealthProbeSettingsModelInput interface {
	pulumi.Input

	ToHealthProbeSettingsModelOutput() HealthProbeSettingsModelOutput
	ToHealthProbeSettingsModelOutputWithContext(context.Context) HealthProbeSettingsModelOutput
}

type HealthProbeSettingsModelArgs struct {
	EnabledState      pulumi.StringPtrInput `pulumi:"enabledState"`
	HealthProbeMethod pulumi.StringPtrInput `pulumi:"healthProbeMethod"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	IntervalInSeconds pulumi.IntPtrInput    `pulumi:"intervalInSeconds"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	Path              pulumi.StringPtrInput `pulumi:"path"`
	Protocol          pulumi.StringPtrInput `pulumi:"protocol"`
}


func (val *HealthProbeSettingsModelArgs) Defaults() *HealthProbeSettingsModelArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HealthProbeMethod) {
		tmp.HealthProbeMethod = pulumi.StringPtr("HEAD")
	}
	return &tmp
}
func (HealthProbeSettingsModelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthProbeSettingsModel)(nil)).Elem()
}

func (i HealthProbeSettingsModelArgs) ToHealthProbeSettingsModelOutput() HealthProbeSettingsModelOutput {
	return i.ToHealthProbeSettingsModelOutputWithContext(context.Background())
}

func (i HealthProbeSettingsModelArgs) ToHealthProbeSettingsModelOutputWithContext(ctx context.Context) HealthProbeSettingsModelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthProbeSettingsModelOutput)
}





type HealthProbeSettingsModelArrayInput interface {
	pulumi.Input

	ToHealthProbeSettingsModelArrayOutput() HealthProbeSettingsModelArrayOutput
	ToHealthProbeSettingsModelArrayOutputWithContext(context.Context) HealthProbeSettingsModelArrayOutput
}

type HealthProbeSettingsModelArray []HealthProbeSettingsModelInput

func (HealthProbeSettingsModelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthProbeSettingsModel)(nil)).Elem()
}

func (i HealthProbeSettingsModelArray) ToHealthProbeSettingsModelArrayOutput() HealthProbeSettingsModelArrayOutput {
	return i.ToHealthProbeSettingsModelArrayOutputWithContext(context.Background())
}

func (i HealthProbeSettingsModelArray) ToHealthProbeSettingsModelArrayOutputWithContext(ctx context.Context) HealthProbeSettingsModelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthProbeSettingsModelArrayOutput)
}

type HealthProbeSettingsModelOutput struct{ *pulumi.OutputState }

func (HealthProbeSettingsModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthProbeSettingsModel)(nil)).Elem()
}

func (o HealthProbeSettingsModelOutput) ToHealthProbeSettingsModelOutput() HealthProbeSettingsModelOutput {
	return o
}

func (o HealthProbeSettingsModelOutput) ToHealthProbeSettingsModelOutputWithContext(ctx context.Context) HealthProbeSettingsModelOutput {
	return o
}

func (o HealthProbeSettingsModelOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModel) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelOutput) HealthProbeMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModel) *string { return v.HealthProbeMethod }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModel) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelOutput) IntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModel) *int { return v.IntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o HealthProbeSettingsModelOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModel) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModel) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModel) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

type HealthProbeSettingsModelArrayOutput struct{ *pulumi.OutputState }

func (HealthProbeSettingsModelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthProbeSettingsModel)(nil)).Elem()
}

func (o HealthProbeSettingsModelArrayOutput) ToHealthProbeSettingsModelArrayOutput() HealthProbeSettingsModelArrayOutput {
	return o
}

func (o HealthProbeSettingsModelArrayOutput) ToHealthProbeSettingsModelArrayOutputWithContext(ctx context.Context) HealthProbeSettingsModelArrayOutput {
	return o
}

func (o HealthProbeSettingsModelArrayOutput) Index(i pulumi.IntInput) HealthProbeSettingsModelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthProbeSettingsModel {
		return vs[0].([]HealthProbeSettingsModel)[vs[1].(int)]
	}).(HealthProbeSettingsModelOutput)
}

type HealthProbeSettingsModelResponse struct {
	EnabledState      *string `pulumi:"enabledState"`
	HealthProbeMethod *string `pulumi:"healthProbeMethod"`
	Id                *string `pulumi:"id"`
	IntervalInSeconds *int    `pulumi:"intervalInSeconds"`
	Name              *string `pulumi:"name"`
	Path              *string `pulumi:"path"`
	Protocol          *string `pulumi:"protocol"`
	ResourceState     string  `pulumi:"resourceState"`
	Type              string  `pulumi:"type"`
}


func (val *HealthProbeSettingsModelResponse) Defaults() *HealthProbeSettingsModelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.HealthProbeMethod) {
		healthProbeMethod_ := "HEAD"
		tmp.HealthProbeMethod = &healthProbeMethod_
	}
	return &tmp
}

type HealthProbeSettingsModelResponseOutput struct{ *pulumi.OutputState }

func (HealthProbeSettingsModelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthProbeSettingsModelResponse)(nil)).Elem()
}

func (o HealthProbeSettingsModelResponseOutput) ToHealthProbeSettingsModelResponseOutput() HealthProbeSettingsModelResponseOutput {
	return o
}

func (o HealthProbeSettingsModelResponseOutput) ToHealthProbeSettingsModelResponseOutputWithContext(ctx context.Context) HealthProbeSettingsModelResponseOutput {
	return o
}

func (o HealthProbeSettingsModelResponseOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModelResponse) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelResponseOutput) HealthProbeMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModelResponse) *string { return v.HealthProbeMethod }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModelResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelResponseOutput) IntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModelResponse) *int { return v.IntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o HealthProbeSettingsModelResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModelResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModelResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthProbeSettingsModelResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o HealthProbeSettingsModelResponseOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v HealthProbeSettingsModelResponse) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o HealthProbeSettingsModelResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v HealthProbeSettingsModelResponse) string { return v.Type }).(pulumi.StringOutput)
}

type HealthProbeSettingsModelResponseArrayOutput struct{ *pulumi.OutputState }

func (HealthProbeSettingsModelResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthProbeSettingsModelResponse)(nil)).Elem()
}

func (o HealthProbeSettingsModelResponseArrayOutput) ToHealthProbeSettingsModelResponseArrayOutput() HealthProbeSettingsModelResponseArrayOutput {
	return o
}

func (o HealthProbeSettingsModelResponseArrayOutput) ToHealthProbeSettingsModelResponseArrayOutputWithContext(ctx context.Context) HealthProbeSettingsModelResponseArrayOutput {
	return o
}

func (o HealthProbeSettingsModelResponseArrayOutput) Index(i pulumi.IntInput) HealthProbeSettingsModelResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthProbeSettingsModelResponse {
		return vs[0].([]HealthProbeSettingsModelResponse)[vs[1].(int)]
	}).(HealthProbeSettingsModelResponseOutput)
}

type KeyVaultCertificateSourceParametersResponseVault struct {
	Id *string `pulumi:"id"`
}

type KeyVaultCertificateSourceParametersResponseVaultOutput struct{ *pulumi.OutputState }

func (KeyVaultCertificateSourceParametersResponseVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultCertificateSourceParametersResponseVault)(nil)).Elem()
}

func (o KeyVaultCertificateSourceParametersResponseVaultOutput) ToKeyVaultCertificateSourceParametersResponseVaultOutput() KeyVaultCertificateSourceParametersResponseVaultOutput {
	return o
}

func (o KeyVaultCertificateSourceParametersResponseVaultOutput) ToKeyVaultCertificateSourceParametersResponseVaultOutputWithContext(ctx context.Context) KeyVaultCertificateSourceParametersResponseVaultOutput {
	return o
}

func (o KeyVaultCertificateSourceParametersResponseVaultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultCertificateSourceParametersResponseVault) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type KeyVaultCertificateSourceParametersResponseVaultPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultCertificateSourceParametersResponseVaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultCertificateSourceParametersResponseVault)(nil)).Elem()
}

func (o KeyVaultCertificateSourceParametersResponseVaultPtrOutput) ToKeyVaultCertificateSourceParametersResponseVaultPtrOutput() KeyVaultCertificateSourceParametersResponseVaultPtrOutput {
	return o
}

func (o KeyVaultCertificateSourceParametersResponseVaultPtrOutput) ToKeyVaultCertificateSourceParametersResponseVaultPtrOutputWithContext(ctx context.Context) KeyVaultCertificateSourceParametersResponseVaultPtrOutput {
	return o
}

func (o KeyVaultCertificateSourceParametersResponseVaultPtrOutput) Elem() KeyVaultCertificateSourceParametersResponseVaultOutput {
	return o.ApplyT(func(v *KeyVaultCertificateSourceParametersResponseVault) KeyVaultCertificateSourceParametersResponseVault {
		if v != nil {
			return *v
		}
		var ret KeyVaultCertificateSourceParametersResponseVault
		return ret
	}).(KeyVaultCertificateSourceParametersResponseVaultOutput)
}

func (o KeyVaultCertificateSourceParametersResponseVaultPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultCertificateSourceParametersResponseVault) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type LoadBalancingSettingsModel struct {
	AdditionalLatencyMilliseconds *int    `pulumi:"additionalLatencyMilliseconds"`
	Id                            *string `pulumi:"id"`
	Name                          *string `pulumi:"name"`
	SampleSize                    *int    `pulumi:"sampleSize"`
	SuccessfulSamplesRequired     *int    `pulumi:"successfulSamplesRequired"`
}





type LoadBalancingSettingsModelInput interface {
	pulumi.Input

	ToLoadBalancingSettingsModelOutput() LoadBalancingSettingsModelOutput
	ToLoadBalancingSettingsModelOutputWithContext(context.Context) LoadBalancingSettingsModelOutput
}

type LoadBalancingSettingsModelArgs struct {
	AdditionalLatencyMilliseconds pulumi.IntPtrInput    `pulumi:"additionalLatencyMilliseconds"`
	Id                            pulumi.StringPtrInput `pulumi:"id"`
	Name                          pulumi.StringPtrInput `pulumi:"name"`
	SampleSize                    pulumi.IntPtrInput    `pulumi:"sampleSize"`
	SuccessfulSamplesRequired     pulumi.IntPtrInput    `pulumi:"successfulSamplesRequired"`
}

func (LoadBalancingSettingsModelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingSettingsModel)(nil)).Elem()
}

func (i LoadBalancingSettingsModelArgs) ToLoadBalancingSettingsModelOutput() LoadBalancingSettingsModelOutput {
	return i.ToLoadBalancingSettingsModelOutputWithContext(context.Background())
}

func (i LoadBalancingSettingsModelArgs) ToLoadBalancingSettingsModelOutputWithContext(ctx context.Context) LoadBalancingSettingsModelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingSettingsModelOutput)
}





type LoadBalancingSettingsModelArrayInput interface {
	pulumi.Input

	ToLoadBalancingSettingsModelArrayOutput() LoadBalancingSettingsModelArrayOutput
	ToLoadBalancingSettingsModelArrayOutputWithContext(context.Context) LoadBalancingSettingsModelArrayOutput
}

type LoadBalancingSettingsModelArray []LoadBalancingSettingsModelInput

func (LoadBalancingSettingsModelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingSettingsModel)(nil)).Elem()
}

func (i LoadBalancingSettingsModelArray) ToLoadBalancingSettingsModelArrayOutput() LoadBalancingSettingsModelArrayOutput {
	return i.ToLoadBalancingSettingsModelArrayOutputWithContext(context.Background())
}

func (i LoadBalancingSettingsModelArray) ToLoadBalancingSettingsModelArrayOutputWithContext(ctx context.Context) LoadBalancingSettingsModelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingSettingsModelArrayOutput)
}

type LoadBalancingSettingsModelOutput struct{ *pulumi.OutputState }

func (LoadBalancingSettingsModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingSettingsModel)(nil)).Elem()
}

func (o LoadBalancingSettingsModelOutput) ToLoadBalancingSettingsModelOutput() LoadBalancingSettingsModelOutput {
	return o
}

func (o LoadBalancingSettingsModelOutput) ToLoadBalancingSettingsModelOutputWithContext(ctx context.Context) LoadBalancingSettingsModelOutput {
	return o
}

func (o LoadBalancingSettingsModelOutput) AdditionalLatencyMilliseconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModel) *int { return v.AdditionalLatencyMilliseconds }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingSettingsModelOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModel) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingSettingsModelOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModel) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingSettingsModelOutput) SampleSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModel) *int { return v.SampleSize }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingSettingsModelOutput) SuccessfulSamplesRequired() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModel) *int { return v.SuccessfulSamplesRequired }).(pulumi.IntPtrOutput)
}

type LoadBalancingSettingsModelArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancingSettingsModelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingSettingsModel)(nil)).Elem()
}

func (o LoadBalancingSettingsModelArrayOutput) ToLoadBalancingSettingsModelArrayOutput() LoadBalancingSettingsModelArrayOutput {
	return o
}

func (o LoadBalancingSettingsModelArrayOutput) ToLoadBalancingSettingsModelArrayOutputWithContext(ctx context.Context) LoadBalancingSettingsModelArrayOutput {
	return o
}

func (o LoadBalancingSettingsModelArrayOutput) Index(i pulumi.IntInput) LoadBalancingSettingsModelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancingSettingsModel {
		return vs[0].([]LoadBalancingSettingsModel)[vs[1].(int)]
	}).(LoadBalancingSettingsModelOutput)
}

type LoadBalancingSettingsModelResponse struct {
	AdditionalLatencyMilliseconds *int    `pulumi:"additionalLatencyMilliseconds"`
	Id                            *string `pulumi:"id"`
	Name                          *string `pulumi:"name"`
	ResourceState                 string  `pulumi:"resourceState"`
	SampleSize                    *int    `pulumi:"sampleSize"`
	SuccessfulSamplesRequired     *int    `pulumi:"successfulSamplesRequired"`
	Type                          string  `pulumi:"type"`
}

type LoadBalancingSettingsModelResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancingSettingsModelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingSettingsModelResponse)(nil)).Elem()
}

func (o LoadBalancingSettingsModelResponseOutput) ToLoadBalancingSettingsModelResponseOutput() LoadBalancingSettingsModelResponseOutput {
	return o
}

func (o LoadBalancingSettingsModelResponseOutput) ToLoadBalancingSettingsModelResponseOutputWithContext(ctx context.Context) LoadBalancingSettingsModelResponseOutput {
	return o
}

func (o LoadBalancingSettingsModelResponseOutput) AdditionalLatencyMilliseconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModelResponse) *int { return v.AdditionalLatencyMilliseconds }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingSettingsModelResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModelResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingSettingsModelResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModelResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingSettingsModelResponseOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModelResponse) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LoadBalancingSettingsModelResponseOutput) SampleSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModelResponse) *int { return v.SampleSize }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingSettingsModelResponseOutput) SuccessfulSamplesRequired() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModelResponse) *int { return v.SuccessfulSamplesRequired }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingSettingsModelResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingSettingsModelResponse) string { return v.Type }).(pulumi.StringOutput)
}

type LoadBalancingSettingsModelResponseArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancingSettingsModelResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingSettingsModelResponse)(nil)).Elem()
}

func (o LoadBalancingSettingsModelResponseArrayOutput) ToLoadBalancingSettingsModelResponseArrayOutput() LoadBalancingSettingsModelResponseArrayOutput {
	return o
}

func (o LoadBalancingSettingsModelResponseArrayOutput) ToLoadBalancingSettingsModelResponseArrayOutputWithContext(ctx context.Context) LoadBalancingSettingsModelResponseArrayOutput {
	return o
}

func (o LoadBalancingSettingsModelResponseArrayOutput) Index(i pulumi.IntInput) LoadBalancingSettingsModelResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancingSettingsModelResponse {
		return vs[0].([]LoadBalancingSettingsModelResponse)[vs[1].(int)]
	}).(LoadBalancingSettingsModelResponseOutput)
}

type MxRecord struct {
	Exchange   *string `pulumi:"exchange"`
	Preference *int    `pulumi:"preference"`
}





type MxRecordInput interface {
	pulumi.Input

	ToMxRecordOutput() MxRecordOutput
	ToMxRecordOutputWithContext(context.Context) MxRecordOutput
}

type MxRecordArgs struct {
	Exchange   pulumi.StringPtrInput `pulumi:"exchange"`
	Preference pulumi.IntPtrInput    `pulumi:"preference"`
}

func (MxRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MxRecord)(nil)).Elem()
}

func (i MxRecordArgs) ToMxRecordOutput() MxRecordOutput {
	return i.ToMxRecordOutputWithContext(context.Background())
}

func (i MxRecordArgs) ToMxRecordOutputWithContext(ctx context.Context) MxRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordOutput)
}





type MxRecordArrayInput interface {
	pulumi.Input

	ToMxRecordArrayOutput() MxRecordArrayOutput
	ToMxRecordArrayOutputWithContext(context.Context) MxRecordArrayOutput
}

type MxRecordArray []MxRecordInput

func (MxRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MxRecord)(nil)).Elem()
}

func (i MxRecordArray) ToMxRecordArrayOutput() MxRecordArrayOutput {
	return i.ToMxRecordArrayOutputWithContext(context.Background())
}

func (i MxRecordArray) ToMxRecordArrayOutputWithContext(ctx context.Context) MxRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordArrayOutput)
}

type MxRecordOutput struct{ *pulumi.OutputState }

func (MxRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MxRecord)(nil)).Elem()
}

func (o MxRecordOutput) ToMxRecordOutput() MxRecordOutput {
	return o
}

func (o MxRecordOutput) ToMxRecordOutputWithContext(ctx context.Context) MxRecordOutput {
	return o
}

func (o MxRecordOutput) Exchange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MxRecord) *string { return v.Exchange }).(pulumi.StringPtrOutput)
}

func (o MxRecordOutput) Preference() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MxRecord) *int { return v.Preference }).(pulumi.IntPtrOutput)
}

type MxRecordArrayOutput struct{ *pulumi.OutputState }

func (MxRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MxRecord)(nil)).Elem()
}

func (o MxRecordArrayOutput) ToMxRecordArrayOutput() MxRecordArrayOutput {
	return o
}

func (o MxRecordArrayOutput) ToMxRecordArrayOutputWithContext(ctx context.Context) MxRecordArrayOutput {
	return o
}

func (o MxRecordArrayOutput) Index(i pulumi.IntInput) MxRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MxRecord {
		return vs[0].([]MxRecord)[vs[1].(int)]
	}).(MxRecordOutput)
}

type MxRecordResponse struct {
	Exchange   *string `pulumi:"exchange"`
	Preference *int    `pulumi:"preference"`
}

type MxRecordResponseOutput struct{ *pulumi.OutputState }

func (MxRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MxRecordResponse)(nil)).Elem()
}

func (o MxRecordResponseOutput) ToMxRecordResponseOutput() MxRecordResponseOutput {
	return o
}

func (o MxRecordResponseOutput) ToMxRecordResponseOutputWithContext(ctx context.Context) MxRecordResponseOutput {
	return o
}

func (o MxRecordResponseOutput) Exchange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MxRecordResponse) *string { return v.Exchange }).(pulumi.StringPtrOutput)
}

func (o MxRecordResponseOutput) Preference() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MxRecordResponse) *int { return v.Preference }).(pulumi.IntPtrOutput)
}

type MxRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (MxRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MxRecordResponse)(nil)).Elem()
}

func (o MxRecordResponseArrayOutput) ToMxRecordResponseArrayOutput() MxRecordResponseArrayOutput {
	return o
}

func (o MxRecordResponseArrayOutput) ToMxRecordResponseArrayOutputWithContext(ctx context.Context) MxRecordResponseArrayOutput {
	return o
}

func (o MxRecordResponseArrayOutput) Index(i pulumi.IntInput) MxRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MxRecordResponse {
		return vs[0].([]MxRecordResponse)[vs[1].(int)]
	}).(MxRecordResponseOutput)
}

type PtrRecord struct {
	Ptrdname *string `pulumi:"ptrdname"`
}





type PtrRecordInput interface {
	pulumi.Input

	ToPtrRecordOutput() PtrRecordOutput
	ToPtrRecordOutputWithContext(context.Context) PtrRecordOutput
}

type PtrRecordArgs struct {
	Ptrdname pulumi.StringPtrInput `pulumi:"ptrdname"`
}

func (PtrRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PtrRecord)(nil)).Elem()
}

func (i PtrRecordArgs) ToPtrRecordOutput() PtrRecordOutput {
	return i.ToPtrRecordOutputWithContext(context.Background())
}

func (i PtrRecordArgs) ToPtrRecordOutputWithContext(ctx context.Context) PtrRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PtrRecordOutput)
}





type PtrRecordArrayInput interface {
	pulumi.Input

	ToPtrRecordArrayOutput() PtrRecordArrayOutput
	ToPtrRecordArrayOutputWithContext(context.Context) PtrRecordArrayOutput
}

type PtrRecordArray []PtrRecordInput

func (PtrRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PtrRecord)(nil)).Elem()
}

func (i PtrRecordArray) ToPtrRecordArrayOutput() PtrRecordArrayOutput {
	return i.ToPtrRecordArrayOutputWithContext(context.Background())
}

func (i PtrRecordArray) ToPtrRecordArrayOutputWithContext(ctx context.Context) PtrRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PtrRecordArrayOutput)
}

type PtrRecordOutput struct{ *pulumi.OutputState }

func (PtrRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PtrRecord)(nil)).Elem()
}

func (o PtrRecordOutput) ToPtrRecordOutput() PtrRecordOutput {
	return o
}

func (o PtrRecordOutput) ToPtrRecordOutputWithContext(ctx context.Context) PtrRecordOutput {
	return o
}

func (o PtrRecordOutput) Ptrdname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PtrRecord) *string { return v.Ptrdname }).(pulumi.StringPtrOutput)
}

type PtrRecordArrayOutput struct{ *pulumi.OutputState }

func (PtrRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PtrRecord)(nil)).Elem()
}

func (o PtrRecordArrayOutput) ToPtrRecordArrayOutput() PtrRecordArrayOutput {
	return o
}

func (o PtrRecordArrayOutput) ToPtrRecordArrayOutputWithContext(ctx context.Context) PtrRecordArrayOutput {
	return o
}

func (o PtrRecordArrayOutput) Index(i pulumi.IntInput) PtrRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PtrRecord {
		return vs[0].([]PtrRecord)[vs[1].(int)]
	}).(PtrRecordOutput)
}

type PtrRecordResponse struct {
	Ptrdname *string `pulumi:"ptrdname"`
}

type PtrRecordResponseOutput struct{ *pulumi.OutputState }

func (PtrRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PtrRecordResponse)(nil)).Elem()
}

func (o PtrRecordResponseOutput) ToPtrRecordResponseOutput() PtrRecordResponseOutput {
	return o
}

func (o PtrRecordResponseOutput) ToPtrRecordResponseOutputWithContext(ctx context.Context) PtrRecordResponseOutput {
	return o
}

func (o PtrRecordResponseOutput) Ptrdname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PtrRecordResponse) *string { return v.Ptrdname }).(pulumi.StringPtrOutput)
}

type PtrRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (PtrRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PtrRecordResponse)(nil)).Elem()
}

func (o PtrRecordResponseArrayOutput) ToPtrRecordResponseArrayOutput() PtrRecordResponseArrayOutput {
	return o
}

func (o PtrRecordResponseArrayOutput) ToPtrRecordResponseArrayOutputWithContext(ctx context.Context) PtrRecordResponseArrayOutput {
	return o
}

func (o PtrRecordResponseArrayOutput) Index(i pulumi.IntInput) PtrRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PtrRecordResponse {
		return vs[0].([]PtrRecordResponse)[vs[1].(int)]
	}).(PtrRecordResponseOutput)
}

type RedirectConfiguration struct {
	CustomFragment    *string `pulumi:"customFragment"`
	CustomHost        *string `pulumi:"customHost"`
	CustomPath        *string `pulumi:"customPath"`
	CustomQueryString *string `pulumi:"customQueryString"`
	OdataType         string  `pulumi:"odataType"`
	RedirectProtocol  *string `pulumi:"redirectProtocol"`
	RedirectType      *string `pulumi:"redirectType"`
}

type RedirectConfigurationResponse struct {
	CustomFragment    *string `pulumi:"customFragment"`
	CustomHost        *string `pulumi:"customHost"`
	CustomPath        *string `pulumi:"customPath"`
	CustomQueryString *string `pulumi:"customQueryString"`
	OdataType         string  `pulumi:"odataType"`
	RedirectProtocol  *string `pulumi:"redirectProtocol"`
	RedirectType      *string `pulumi:"redirectType"`
}

type RoutingRule struct {
	AcceptedProtocols  []string      `pulumi:"acceptedProtocols"`
	EnabledState       *string       `pulumi:"enabledState"`
	FrontendEndpoints  []SubResource `pulumi:"frontendEndpoints"`
	Id                 *string       `pulumi:"id"`
	Name               *string       `pulumi:"name"`
	PatternsToMatch    []string      `pulumi:"patternsToMatch"`
	RouteConfiguration interface{}   `pulumi:"routeConfiguration"`
	RulesEngine        *SubResource  `pulumi:"rulesEngine"`
}





type RoutingRuleInput interface {
	pulumi.Input

	ToRoutingRuleOutput() RoutingRuleOutput
	ToRoutingRuleOutputWithContext(context.Context) RoutingRuleOutput
}

type RoutingRuleArgs struct {
	AcceptedProtocols  pulumi.StringArrayInput `pulumi:"acceptedProtocols"`
	EnabledState       pulumi.StringPtrInput   `pulumi:"enabledState"`
	FrontendEndpoints  SubResourceArrayInput   `pulumi:"frontendEndpoints"`
	Id                 pulumi.StringPtrInput   `pulumi:"id"`
	Name               pulumi.StringPtrInput   `pulumi:"name"`
	PatternsToMatch    pulumi.StringArrayInput `pulumi:"patternsToMatch"`
	RouteConfiguration pulumi.Input            `pulumi:"routeConfiguration"`
	RulesEngine        SubResourcePtrInput     `pulumi:"rulesEngine"`
}

func (RoutingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingRule)(nil)).Elem()
}

func (i RoutingRuleArgs) ToRoutingRuleOutput() RoutingRuleOutput {
	return i.ToRoutingRuleOutputWithContext(context.Background())
}

func (i RoutingRuleArgs) ToRoutingRuleOutputWithContext(ctx context.Context) RoutingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingRuleOutput)
}





type RoutingRuleArrayInput interface {
	pulumi.Input

	ToRoutingRuleArrayOutput() RoutingRuleArrayOutput
	ToRoutingRuleArrayOutputWithContext(context.Context) RoutingRuleArrayOutput
}

type RoutingRuleArray []RoutingRuleInput

func (RoutingRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingRule)(nil)).Elem()
}

func (i RoutingRuleArray) ToRoutingRuleArrayOutput() RoutingRuleArrayOutput {
	return i.ToRoutingRuleArrayOutputWithContext(context.Background())
}

func (i RoutingRuleArray) ToRoutingRuleArrayOutputWithContext(ctx context.Context) RoutingRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingRuleArrayOutput)
}

type RoutingRuleOutput struct{ *pulumi.OutputState }

func (RoutingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingRule)(nil)).Elem()
}

func (o RoutingRuleOutput) ToRoutingRuleOutput() RoutingRuleOutput {
	return o
}

func (o RoutingRuleOutput) ToRoutingRuleOutputWithContext(ctx context.Context) RoutingRuleOutput {
	return o
}

func (o RoutingRuleOutput) AcceptedProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RoutingRule) []string { return v.AcceptedProtocols }).(pulumi.StringArrayOutput)
}

func (o RoutingRuleOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingRule) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o RoutingRuleOutput) FrontendEndpoints() SubResourceArrayOutput {
	return o.ApplyT(func(v RoutingRule) []SubResource { return v.FrontendEndpoints }).(SubResourceArrayOutput)
}

func (o RoutingRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RoutingRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RoutingRuleOutput) PatternsToMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RoutingRule) []string { return v.PatternsToMatch }).(pulumi.StringArrayOutput)
}

func (o RoutingRuleOutput) RouteConfiguration() pulumi.AnyOutput {
	return o.ApplyT(func(v RoutingRule) interface{} { return v.RouteConfiguration }).(pulumi.AnyOutput)
}

func (o RoutingRuleOutput) RulesEngine() SubResourcePtrOutput {
	return o.ApplyT(func(v RoutingRule) *SubResource { return v.RulesEngine }).(SubResourcePtrOutput)
}

type RoutingRuleArrayOutput struct{ *pulumi.OutputState }

func (RoutingRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingRule)(nil)).Elem()
}

func (o RoutingRuleArrayOutput) ToRoutingRuleArrayOutput() RoutingRuleArrayOutput {
	return o
}

func (o RoutingRuleArrayOutput) ToRoutingRuleArrayOutputWithContext(ctx context.Context) RoutingRuleArrayOutput {
	return o
}

func (o RoutingRuleArrayOutput) Index(i pulumi.IntInput) RoutingRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingRule {
		return vs[0].([]RoutingRule)[vs[1].(int)]
	}).(RoutingRuleOutput)
}

type RoutingRuleResponse struct {
	AcceptedProtocols  []string              `pulumi:"acceptedProtocols"`
	EnabledState       *string               `pulumi:"enabledState"`
	FrontendEndpoints  []SubResourceResponse `pulumi:"frontendEndpoints"`
	Id                 *string               `pulumi:"id"`
	Name               *string               `pulumi:"name"`
	PatternsToMatch    []string              `pulumi:"patternsToMatch"`
	ResourceState      string                `pulumi:"resourceState"`
	RouteConfiguration interface{}           `pulumi:"routeConfiguration"`
	RulesEngine        *SubResourceResponse  `pulumi:"rulesEngine"`
	Type               string                `pulumi:"type"`
}

type RoutingRuleResponseOutput struct{ *pulumi.OutputState }

func (RoutingRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingRuleResponse)(nil)).Elem()
}

func (o RoutingRuleResponseOutput) ToRoutingRuleResponseOutput() RoutingRuleResponseOutput {
	return o
}

func (o RoutingRuleResponseOutput) ToRoutingRuleResponseOutputWithContext(ctx context.Context) RoutingRuleResponseOutput {
	return o
}

func (o RoutingRuleResponseOutput) AcceptedProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RoutingRuleResponse) []string { return v.AcceptedProtocols }).(pulumi.StringArrayOutput)
}

func (o RoutingRuleResponseOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingRuleResponse) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

func (o RoutingRuleResponseOutput) FrontendEndpoints() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v RoutingRuleResponse) []SubResourceResponse { return v.FrontendEndpoints }).(SubResourceResponseArrayOutput)
}

func (o RoutingRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RoutingRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RoutingRuleResponseOutput) PatternsToMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RoutingRuleResponse) []string { return v.PatternsToMatch }).(pulumi.StringArrayOutput)
}

func (o RoutingRuleResponseOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingRuleResponse) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o RoutingRuleResponseOutput) RouteConfiguration() pulumi.AnyOutput {
	return o.ApplyT(func(v RoutingRuleResponse) interface{} { return v.RouteConfiguration }).(pulumi.AnyOutput)
}

func (o RoutingRuleResponseOutput) RulesEngine() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v RoutingRuleResponse) *SubResourceResponse { return v.RulesEngine }).(SubResourceResponsePtrOutput)
}

func (o RoutingRuleResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingRuleResponse) string { return v.Type }).(pulumi.StringOutput)
}

type RoutingRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (RoutingRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingRuleResponse)(nil)).Elem()
}

func (o RoutingRuleResponseArrayOutput) ToRoutingRuleResponseArrayOutput() RoutingRuleResponseArrayOutput {
	return o
}

func (o RoutingRuleResponseArrayOutput) ToRoutingRuleResponseArrayOutputWithContext(ctx context.Context) RoutingRuleResponseArrayOutput {
	return o
}

func (o RoutingRuleResponseArrayOutput) Index(i pulumi.IntInput) RoutingRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingRuleResponse {
		return vs[0].([]RoutingRuleResponse)[vs[1].(int)]
	}).(RoutingRuleResponseOutput)
}

type RulesEngineAction struct {
	RequestHeaderActions       []HeaderAction `pulumi:"requestHeaderActions"`
	ResponseHeaderActions      []HeaderAction `pulumi:"responseHeaderActions"`
	RouteConfigurationOverride interface{}    `pulumi:"routeConfigurationOverride"`
}





type RulesEngineActionInput interface {
	pulumi.Input

	ToRulesEngineActionOutput() RulesEngineActionOutput
	ToRulesEngineActionOutputWithContext(context.Context) RulesEngineActionOutput
}

type RulesEngineActionArgs struct {
	RequestHeaderActions       HeaderActionArrayInput `pulumi:"requestHeaderActions"`
	ResponseHeaderActions      HeaderActionArrayInput `pulumi:"responseHeaderActions"`
	RouteConfigurationOverride pulumi.Input           `pulumi:"routeConfigurationOverride"`
}

func (RulesEngineActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineAction)(nil)).Elem()
}

func (i RulesEngineActionArgs) ToRulesEngineActionOutput() RulesEngineActionOutput {
	return i.ToRulesEngineActionOutputWithContext(context.Background())
}

func (i RulesEngineActionArgs) ToRulesEngineActionOutputWithContext(ctx context.Context) RulesEngineActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesEngineActionOutput)
}

type RulesEngineActionOutput struct{ *pulumi.OutputState }

func (RulesEngineActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineAction)(nil)).Elem()
}

func (o RulesEngineActionOutput) ToRulesEngineActionOutput() RulesEngineActionOutput {
	return o
}

func (o RulesEngineActionOutput) ToRulesEngineActionOutputWithContext(ctx context.Context) RulesEngineActionOutput {
	return o
}

func (o RulesEngineActionOutput) RequestHeaderActions() HeaderActionArrayOutput {
	return o.ApplyT(func(v RulesEngineAction) []HeaderAction { return v.RequestHeaderActions }).(HeaderActionArrayOutput)
}

func (o RulesEngineActionOutput) ResponseHeaderActions() HeaderActionArrayOutput {
	return o.ApplyT(func(v RulesEngineAction) []HeaderAction { return v.ResponseHeaderActions }).(HeaderActionArrayOutput)
}

func (o RulesEngineActionOutput) RouteConfigurationOverride() pulumi.AnyOutput {
	return o.ApplyT(func(v RulesEngineAction) interface{} { return v.RouteConfigurationOverride }).(pulumi.AnyOutput)
}

type RulesEngineActionResponse struct {
	RequestHeaderActions       []HeaderActionResponse `pulumi:"requestHeaderActions"`
	ResponseHeaderActions      []HeaderActionResponse `pulumi:"responseHeaderActions"`
	RouteConfigurationOverride interface{}            `pulumi:"routeConfigurationOverride"`
}

type RulesEngineActionResponseOutput struct{ *pulumi.OutputState }

func (RulesEngineActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineActionResponse)(nil)).Elem()
}

func (o RulesEngineActionResponseOutput) ToRulesEngineActionResponseOutput() RulesEngineActionResponseOutput {
	return o
}

func (o RulesEngineActionResponseOutput) ToRulesEngineActionResponseOutputWithContext(ctx context.Context) RulesEngineActionResponseOutput {
	return o
}

func (o RulesEngineActionResponseOutput) RequestHeaderActions() HeaderActionResponseArrayOutput {
	return o.ApplyT(func(v RulesEngineActionResponse) []HeaderActionResponse { return v.RequestHeaderActions }).(HeaderActionResponseArrayOutput)
}

func (o RulesEngineActionResponseOutput) ResponseHeaderActions() HeaderActionResponseArrayOutput {
	return o.ApplyT(func(v RulesEngineActionResponse) []HeaderActionResponse { return v.ResponseHeaderActions }).(HeaderActionResponseArrayOutput)
}

func (o RulesEngineActionResponseOutput) RouteConfigurationOverride() pulumi.AnyOutput {
	return o.ApplyT(func(v RulesEngineActionResponse) interface{} { return v.RouteConfigurationOverride }).(pulumi.AnyOutput)
}

type RulesEngineMatchCondition struct {
	NegateCondition          *bool    `pulumi:"negateCondition"`
	RulesEngineMatchValue    []string `pulumi:"rulesEngineMatchValue"`
	RulesEngineMatchVariable string   `pulumi:"rulesEngineMatchVariable"`
	RulesEngineOperator      string   `pulumi:"rulesEngineOperator"`
	Selector                 *string  `pulumi:"selector"`
	Transforms               []string `pulumi:"transforms"`
}





type RulesEngineMatchConditionInput interface {
	pulumi.Input

	ToRulesEngineMatchConditionOutput() RulesEngineMatchConditionOutput
	ToRulesEngineMatchConditionOutputWithContext(context.Context) RulesEngineMatchConditionOutput
}

type RulesEngineMatchConditionArgs struct {
	NegateCondition          pulumi.BoolPtrInput     `pulumi:"negateCondition"`
	RulesEngineMatchValue    pulumi.StringArrayInput `pulumi:"rulesEngineMatchValue"`
	RulesEngineMatchVariable pulumi.StringInput      `pulumi:"rulesEngineMatchVariable"`
	RulesEngineOperator      pulumi.StringInput      `pulumi:"rulesEngineOperator"`
	Selector                 pulumi.StringPtrInput   `pulumi:"selector"`
	Transforms               pulumi.StringArrayInput `pulumi:"transforms"`
}

func (RulesEngineMatchConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineMatchCondition)(nil)).Elem()
}

func (i RulesEngineMatchConditionArgs) ToRulesEngineMatchConditionOutput() RulesEngineMatchConditionOutput {
	return i.ToRulesEngineMatchConditionOutputWithContext(context.Background())
}

func (i RulesEngineMatchConditionArgs) ToRulesEngineMatchConditionOutputWithContext(ctx context.Context) RulesEngineMatchConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesEngineMatchConditionOutput)
}





type RulesEngineMatchConditionArrayInput interface {
	pulumi.Input

	ToRulesEngineMatchConditionArrayOutput() RulesEngineMatchConditionArrayOutput
	ToRulesEngineMatchConditionArrayOutputWithContext(context.Context) RulesEngineMatchConditionArrayOutput
}

type RulesEngineMatchConditionArray []RulesEngineMatchConditionInput

func (RulesEngineMatchConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineMatchCondition)(nil)).Elem()
}

func (i RulesEngineMatchConditionArray) ToRulesEngineMatchConditionArrayOutput() RulesEngineMatchConditionArrayOutput {
	return i.ToRulesEngineMatchConditionArrayOutputWithContext(context.Background())
}

func (i RulesEngineMatchConditionArray) ToRulesEngineMatchConditionArrayOutputWithContext(ctx context.Context) RulesEngineMatchConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesEngineMatchConditionArrayOutput)
}

type RulesEngineMatchConditionOutput struct{ *pulumi.OutputState }

func (RulesEngineMatchConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineMatchCondition)(nil)).Elem()
}

func (o RulesEngineMatchConditionOutput) ToRulesEngineMatchConditionOutput() RulesEngineMatchConditionOutput {
	return o
}

func (o RulesEngineMatchConditionOutput) ToRulesEngineMatchConditionOutputWithContext(ctx context.Context) RulesEngineMatchConditionOutput {
	return o
}

func (o RulesEngineMatchConditionOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RulesEngineMatchCondition) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o RulesEngineMatchConditionOutput) RulesEngineMatchValue() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RulesEngineMatchCondition) []string { return v.RulesEngineMatchValue }).(pulumi.StringArrayOutput)
}

func (o RulesEngineMatchConditionOutput) RulesEngineMatchVariable() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineMatchCondition) string { return v.RulesEngineMatchVariable }).(pulumi.StringOutput)
}

func (o RulesEngineMatchConditionOutput) RulesEngineOperator() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineMatchCondition) string { return v.RulesEngineOperator }).(pulumi.StringOutput)
}

func (o RulesEngineMatchConditionOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesEngineMatchCondition) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

func (o RulesEngineMatchConditionOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RulesEngineMatchCondition) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type RulesEngineMatchConditionArrayOutput struct{ *pulumi.OutputState }

func (RulesEngineMatchConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineMatchCondition)(nil)).Elem()
}

func (o RulesEngineMatchConditionArrayOutput) ToRulesEngineMatchConditionArrayOutput() RulesEngineMatchConditionArrayOutput {
	return o
}

func (o RulesEngineMatchConditionArrayOutput) ToRulesEngineMatchConditionArrayOutputWithContext(ctx context.Context) RulesEngineMatchConditionArrayOutput {
	return o
}

func (o RulesEngineMatchConditionArrayOutput) Index(i pulumi.IntInput) RulesEngineMatchConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RulesEngineMatchCondition {
		return vs[0].([]RulesEngineMatchCondition)[vs[1].(int)]
	}).(RulesEngineMatchConditionOutput)
}

type RulesEngineMatchConditionResponse struct {
	NegateCondition          *bool    `pulumi:"negateCondition"`
	RulesEngineMatchValue    []string `pulumi:"rulesEngineMatchValue"`
	RulesEngineMatchVariable string   `pulumi:"rulesEngineMatchVariable"`
	RulesEngineOperator      string   `pulumi:"rulesEngineOperator"`
	Selector                 *string  `pulumi:"selector"`
	Transforms               []string `pulumi:"transforms"`
}

type RulesEngineMatchConditionResponseOutput struct{ *pulumi.OutputState }

func (RulesEngineMatchConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineMatchConditionResponse)(nil)).Elem()
}

func (o RulesEngineMatchConditionResponseOutput) ToRulesEngineMatchConditionResponseOutput() RulesEngineMatchConditionResponseOutput {
	return o
}

func (o RulesEngineMatchConditionResponseOutput) ToRulesEngineMatchConditionResponseOutputWithContext(ctx context.Context) RulesEngineMatchConditionResponseOutput {
	return o
}

func (o RulesEngineMatchConditionResponseOutput) NegateCondition() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RulesEngineMatchConditionResponse) *bool { return v.NegateCondition }).(pulumi.BoolPtrOutput)
}

func (o RulesEngineMatchConditionResponseOutput) RulesEngineMatchValue() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RulesEngineMatchConditionResponse) []string { return v.RulesEngineMatchValue }).(pulumi.StringArrayOutput)
}

func (o RulesEngineMatchConditionResponseOutput) RulesEngineMatchVariable() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineMatchConditionResponse) string { return v.RulesEngineMatchVariable }).(pulumi.StringOutput)
}

func (o RulesEngineMatchConditionResponseOutput) RulesEngineOperator() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineMatchConditionResponse) string { return v.RulesEngineOperator }).(pulumi.StringOutput)
}

func (o RulesEngineMatchConditionResponseOutput) Selector() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesEngineMatchConditionResponse) *string { return v.Selector }).(pulumi.StringPtrOutput)
}

func (o RulesEngineMatchConditionResponseOutput) Transforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RulesEngineMatchConditionResponse) []string { return v.Transforms }).(pulumi.StringArrayOutput)
}

type RulesEngineMatchConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (RulesEngineMatchConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineMatchConditionResponse)(nil)).Elem()
}

func (o RulesEngineMatchConditionResponseArrayOutput) ToRulesEngineMatchConditionResponseArrayOutput() RulesEngineMatchConditionResponseArrayOutput {
	return o
}

func (o RulesEngineMatchConditionResponseArrayOutput) ToRulesEngineMatchConditionResponseArrayOutputWithContext(ctx context.Context) RulesEngineMatchConditionResponseArrayOutput {
	return o
}

func (o RulesEngineMatchConditionResponseArrayOutput) Index(i pulumi.IntInput) RulesEngineMatchConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RulesEngineMatchConditionResponse {
		return vs[0].([]RulesEngineMatchConditionResponse)[vs[1].(int)]
	}).(RulesEngineMatchConditionResponseOutput)
}

type RulesEngineResponse struct {
	Id            string                    `pulumi:"id"`
	Name          string                    `pulumi:"name"`
	ResourceState string                    `pulumi:"resourceState"`
	Rules         []RulesEngineRuleResponse `pulumi:"rules"`
	Type          string                    `pulumi:"type"`
}

type RulesEngineResponseOutput struct{ *pulumi.OutputState }

func (RulesEngineResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineResponse)(nil)).Elem()
}

func (o RulesEngineResponseOutput) ToRulesEngineResponseOutput() RulesEngineResponseOutput {
	return o
}

func (o RulesEngineResponseOutput) ToRulesEngineResponseOutputWithContext(ctx context.Context) RulesEngineResponseOutput {
	return o
}

func (o RulesEngineResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o RulesEngineResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RulesEngineResponseOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineResponse) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o RulesEngineResponseOutput) Rules() RulesEngineRuleResponseArrayOutput {
	return o.ApplyT(func(v RulesEngineResponse) []RulesEngineRuleResponse { return v.Rules }).(RulesEngineRuleResponseArrayOutput)
}

func (o RulesEngineResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineResponse) string { return v.Type }).(pulumi.StringOutput)
}

type RulesEngineResponseArrayOutput struct{ *pulumi.OutputState }

func (RulesEngineResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineResponse)(nil)).Elem()
}

func (o RulesEngineResponseArrayOutput) ToRulesEngineResponseArrayOutput() RulesEngineResponseArrayOutput {
	return o
}

func (o RulesEngineResponseArrayOutput) ToRulesEngineResponseArrayOutputWithContext(ctx context.Context) RulesEngineResponseArrayOutput {
	return o
}

func (o RulesEngineResponseArrayOutput) Index(i pulumi.IntInput) RulesEngineResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RulesEngineResponse {
		return vs[0].([]RulesEngineResponse)[vs[1].(int)]
	}).(RulesEngineResponseOutput)
}

type RulesEngineRule struct {
	Action                  RulesEngineAction           `pulumi:"action"`
	MatchConditions         []RulesEngineMatchCondition `pulumi:"matchConditions"`
	MatchProcessingBehavior *string                     `pulumi:"matchProcessingBehavior"`
	Name                    string                      `pulumi:"name"`
	Priority                int                         `pulumi:"priority"`
}





type RulesEngineRuleInput interface {
	pulumi.Input

	ToRulesEngineRuleOutput() RulesEngineRuleOutput
	ToRulesEngineRuleOutputWithContext(context.Context) RulesEngineRuleOutput
}

type RulesEngineRuleArgs struct {
	Action                  RulesEngineActionInput              `pulumi:"action"`
	MatchConditions         RulesEngineMatchConditionArrayInput `pulumi:"matchConditions"`
	MatchProcessingBehavior pulumi.StringPtrInput               `pulumi:"matchProcessingBehavior"`
	Name                    pulumi.StringInput                  `pulumi:"name"`
	Priority                pulumi.IntInput                     `pulumi:"priority"`
}

func (RulesEngineRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineRule)(nil)).Elem()
}

func (i RulesEngineRuleArgs) ToRulesEngineRuleOutput() RulesEngineRuleOutput {
	return i.ToRulesEngineRuleOutputWithContext(context.Background())
}

func (i RulesEngineRuleArgs) ToRulesEngineRuleOutputWithContext(ctx context.Context) RulesEngineRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesEngineRuleOutput)
}





type RulesEngineRuleArrayInput interface {
	pulumi.Input

	ToRulesEngineRuleArrayOutput() RulesEngineRuleArrayOutput
	ToRulesEngineRuleArrayOutputWithContext(context.Context) RulesEngineRuleArrayOutput
}

type RulesEngineRuleArray []RulesEngineRuleInput

func (RulesEngineRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineRule)(nil)).Elem()
}

func (i RulesEngineRuleArray) ToRulesEngineRuleArrayOutput() RulesEngineRuleArrayOutput {
	return i.ToRulesEngineRuleArrayOutputWithContext(context.Background())
}

func (i RulesEngineRuleArray) ToRulesEngineRuleArrayOutputWithContext(ctx context.Context) RulesEngineRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RulesEngineRuleArrayOutput)
}

type RulesEngineRuleOutput struct{ *pulumi.OutputState }

func (RulesEngineRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineRule)(nil)).Elem()
}

func (o RulesEngineRuleOutput) ToRulesEngineRuleOutput() RulesEngineRuleOutput {
	return o
}

func (o RulesEngineRuleOutput) ToRulesEngineRuleOutputWithContext(ctx context.Context) RulesEngineRuleOutput {
	return o
}

func (o RulesEngineRuleOutput) Action() RulesEngineActionOutput {
	return o.ApplyT(func(v RulesEngineRule) RulesEngineAction { return v.Action }).(RulesEngineActionOutput)
}

func (o RulesEngineRuleOutput) MatchConditions() RulesEngineMatchConditionArrayOutput {
	return o.ApplyT(func(v RulesEngineRule) []RulesEngineMatchCondition { return v.MatchConditions }).(RulesEngineMatchConditionArrayOutput)
}

func (o RulesEngineRuleOutput) MatchProcessingBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesEngineRule) *string { return v.MatchProcessingBehavior }).(pulumi.StringPtrOutput)
}

func (o RulesEngineRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineRule) string { return v.Name }).(pulumi.StringOutput)
}

func (o RulesEngineRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v RulesEngineRule) int { return v.Priority }).(pulumi.IntOutput)
}

type RulesEngineRuleArrayOutput struct{ *pulumi.OutputState }

func (RulesEngineRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineRule)(nil)).Elem()
}

func (o RulesEngineRuleArrayOutput) ToRulesEngineRuleArrayOutput() RulesEngineRuleArrayOutput {
	return o
}

func (o RulesEngineRuleArrayOutput) ToRulesEngineRuleArrayOutputWithContext(ctx context.Context) RulesEngineRuleArrayOutput {
	return o
}

func (o RulesEngineRuleArrayOutput) Index(i pulumi.IntInput) RulesEngineRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RulesEngineRule {
		return vs[0].([]RulesEngineRule)[vs[1].(int)]
	}).(RulesEngineRuleOutput)
}

type RulesEngineRuleResponse struct {
	Action                  RulesEngineActionResponse           `pulumi:"action"`
	MatchConditions         []RulesEngineMatchConditionResponse `pulumi:"matchConditions"`
	MatchProcessingBehavior *string                             `pulumi:"matchProcessingBehavior"`
	Name                    string                              `pulumi:"name"`
	Priority                int                                 `pulumi:"priority"`
}

type RulesEngineRuleResponseOutput struct{ *pulumi.OutputState }

func (RulesEngineRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesEngineRuleResponse)(nil)).Elem()
}

func (o RulesEngineRuleResponseOutput) ToRulesEngineRuleResponseOutput() RulesEngineRuleResponseOutput {
	return o
}

func (o RulesEngineRuleResponseOutput) ToRulesEngineRuleResponseOutputWithContext(ctx context.Context) RulesEngineRuleResponseOutput {
	return o
}

func (o RulesEngineRuleResponseOutput) Action() RulesEngineActionResponseOutput {
	return o.ApplyT(func(v RulesEngineRuleResponse) RulesEngineActionResponse { return v.Action }).(RulesEngineActionResponseOutput)
}

func (o RulesEngineRuleResponseOutput) MatchConditions() RulesEngineMatchConditionResponseArrayOutput {
	return o.ApplyT(func(v RulesEngineRuleResponse) []RulesEngineMatchConditionResponse { return v.MatchConditions }).(RulesEngineMatchConditionResponseArrayOutput)
}

func (o RulesEngineRuleResponseOutput) MatchProcessingBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesEngineRuleResponse) *string { return v.MatchProcessingBehavior }).(pulumi.StringPtrOutput)
}

func (o RulesEngineRuleResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RulesEngineRuleResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RulesEngineRuleResponseOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v RulesEngineRuleResponse) int { return v.Priority }).(pulumi.IntOutput)
}

type RulesEngineRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (RulesEngineRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RulesEngineRuleResponse)(nil)).Elem()
}

func (o RulesEngineRuleResponseArrayOutput) ToRulesEngineRuleResponseArrayOutput() RulesEngineRuleResponseArrayOutput {
	return o
}

func (o RulesEngineRuleResponseArrayOutput) ToRulesEngineRuleResponseArrayOutputWithContext(ctx context.Context) RulesEngineRuleResponseArrayOutput {
	return o
}

func (o RulesEngineRuleResponseArrayOutput) Index(i pulumi.IntInput) RulesEngineRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RulesEngineRuleResponse {
		return vs[0].([]RulesEngineRuleResponse)[vs[1].(int)]
	}).(RulesEngineRuleResponseOutput)
}

type SoaRecord struct {
	Email        *string  `pulumi:"email"`
	ExpireTime   *float64 `pulumi:"expireTime"`
	Host         *string  `pulumi:"host"`
	MinimumTtl   *float64 `pulumi:"minimumTtl"`
	RefreshTime  *float64 `pulumi:"refreshTime"`
	RetryTime    *float64 `pulumi:"retryTime"`
	SerialNumber *float64 `pulumi:"serialNumber"`
}





type SoaRecordInput interface {
	pulumi.Input

	ToSoaRecordOutput() SoaRecordOutput
	ToSoaRecordOutputWithContext(context.Context) SoaRecordOutput
}

type SoaRecordArgs struct {
	Email        pulumi.StringPtrInput  `pulumi:"email"`
	ExpireTime   pulumi.Float64PtrInput `pulumi:"expireTime"`
	Host         pulumi.StringPtrInput  `pulumi:"host"`
	MinimumTtl   pulumi.Float64PtrInput `pulumi:"minimumTtl"`
	RefreshTime  pulumi.Float64PtrInput `pulumi:"refreshTime"`
	RetryTime    pulumi.Float64PtrInput `pulumi:"retryTime"`
	SerialNumber pulumi.Float64PtrInput `pulumi:"serialNumber"`
}

func (SoaRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SoaRecord)(nil)).Elem()
}

func (i SoaRecordArgs) ToSoaRecordOutput() SoaRecordOutput {
	return i.ToSoaRecordOutputWithContext(context.Background())
}

func (i SoaRecordArgs) ToSoaRecordOutputWithContext(ctx context.Context) SoaRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoaRecordOutput)
}

func (i SoaRecordArgs) ToSoaRecordPtrOutput() SoaRecordPtrOutput {
	return i.ToSoaRecordPtrOutputWithContext(context.Background())
}

func (i SoaRecordArgs) ToSoaRecordPtrOutputWithContext(ctx context.Context) SoaRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoaRecordOutput).ToSoaRecordPtrOutputWithContext(ctx)
}









type SoaRecordPtrInput interface {
	pulumi.Input

	ToSoaRecordPtrOutput() SoaRecordPtrOutput
	ToSoaRecordPtrOutputWithContext(context.Context) SoaRecordPtrOutput
}

type soaRecordPtrType SoaRecordArgs

func SoaRecordPtr(v *SoaRecordArgs) SoaRecordPtrInput {
	return (*soaRecordPtrType)(v)
}

func (*soaRecordPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SoaRecord)(nil)).Elem()
}

func (i *soaRecordPtrType) ToSoaRecordPtrOutput() SoaRecordPtrOutput {
	return i.ToSoaRecordPtrOutputWithContext(context.Background())
}

func (i *soaRecordPtrType) ToSoaRecordPtrOutputWithContext(ctx context.Context) SoaRecordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoaRecordPtrOutput)
}

type SoaRecordOutput struct{ *pulumi.OutputState }

func (SoaRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SoaRecord)(nil)).Elem()
}

func (o SoaRecordOutput) ToSoaRecordOutput() SoaRecordOutput {
	return o
}

func (o SoaRecordOutput) ToSoaRecordOutputWithContext(ctx context.Context) SoaRecordOutput {
	return o
}

func (o SoaRecordOutput) ToSoaRecordPtrOutput() SoaRecordPtrOutput {
	return o.ToSoaRecordPtrOutputWithContext(context.Background())
}

func (o SoaRecordOutput) ToSoaRecordPtrOutputWithContext(ctx context.Context) SoaRecordPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SoaRecord) *SoaRecord {
		return &v
	}).(SoaRecordPtrOutput)
}

func (o SoaRecordOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoaRecord) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o SoaRecordOutput) ExpireTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.ExpireTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoaRecord) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o SoaRecordOutput) MinimumTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.MinimumTtl }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordOutput) RefreshTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.RefreshTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordOutput) RetryTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.RetryTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordOutput) SerialNumber() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecord) *float64 { return v.SerialNumber }).(pulumi.Float64PtrOutput)
}

type SoaRecordPtrOutput struct{ *pulumi.OutputState }

func (SoaRecordPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoaRecord)(nil)).Elem()
}

func (o SoaRecordPtrOutput) ToSoaRecordPtrOutput() SoaRecordPtrOutput {
	return o
}

func (o SoaRecordPtrOutput) ToSoaRecordPtrOutputWithContext(ctx context.Context) SoaRecordPtrOutput {
	return o
}

func (o SoaRecordPtrOutput) Elem() SoaRecordOutput {
	return o.ApplyT(func(v *SoaRecord) SoaRecord {
		if v != nil {
			return *v
		}
		var ret SoaRecord
		return ret
	}).(SoaRecordOutput)
}

func (o SoaRecordPtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoaRecord) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o SoaRecordPtrOutput) ExpireTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.ExpireTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordPtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoaRecord) *string {
		if v == nil {
			return nil
		}
		return v.Host
	}).(pulumi.StringPtrOutput)
}

func (o SoaRecordPtrOutput) MinimumTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.MinimumTtl
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordPtrOutput) RefreshTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.RefreshTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordPtrOutput) RetryTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.RetryTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordPtrOutput) SerialNumber() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecord) *float64 {
		if v == nil {
			return nil
		}
		return v.SerialNumber
	}).(pulumi.Float64PtrOutput)
}

type SoaRecordResponse struct {
	Email        *string  `pulumi:"email"`
	ExpireTime   *float64 `pulumi:"expireTime"`
	Host         *string  `pulumi:"host"`
	MinimumTtl   *float64 `pulumi:"minimumTtl"`
	RefreshTime  *float64 `pulumi:"refreshTime"`
	RetryTime    *float64 `pulumi:"retryTime"`
	SerialNumber *float64 `pulumi:"serialNumber"`
}

type SoaRecordResponseOutput struct{ *pulumi.OutputState }

func (SoaRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SoaRecordResponse)(nil)).Elem()
}

func (o SoaRecordResponseOutput) ToSoaRecordResponseOutput() SoaRecordResponseOutput {
	return o
}

func (o SoaRecordResponseOutput) ToSoaRecordResponseOutputWithContext(ctx context.Context) SoaRecordResponseOutput {
	return o
}

func (o SoaRecordResponseOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o SoaRecordResponseOutput) ExpireTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.ExpireTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponseOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o SoaRecordResponseOutput) MinimumTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.MinimumTtl }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponseOutput) RefreshTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.RefreshTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponseOutput) RetryTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.RetryTime }).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponseOutput) SerialNumber() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SoaRecordResponse) *float64 { return v.SerialNumber }).(pulumi.Float64PtrOutput)
}

type SoaRecordResponsePtrOutput struct{ *pulumi.OutputState }

func (SoaRecordResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoaRecordResponse)(nil)).Elem()
}

func (o SoaRecordResponsePtrOutput) ToSoaRecordResponsePtrOutput() SoaRecordResponsePtrOutput {
	return o
}

func (o SoaRecordResponsePtrOutput) ToSoaRecordResponsePtrOutputWithContext(ctx context.Context) SoaRecordResponsePtrOutput {
	return o
}

func (o SoaRecordResponsePtrOutput) Elem() SoaRecordResponseOutput {
	return o.ApplyT(func(v *SoaRecordResponse) SoaRecordResponse {
		if v != nil {
			return *v
		}
		var ret SoaRecordResponse
		return ret
	}).(SoaRecordResponseOutput)
}

func (o SoaRecordResponsePtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o SoaRecordResponsePtrOutput) ExpireTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.ExpireTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponsePtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *string {
		if v == nil {
			return nil
		}
		return v.Host
	}).(pulumi.StringPtrOutput)
}

func (o SoaRecordResponsePtrOutput) MinimumTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MinimumTtl
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponsePtrOutput) RefreshTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.RefreshTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponsePtrOutput) RetryTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.RetryTime
	}).(pulumi.Float64PtrOutput)
}

func (o SoaRecordResponsePtrOutput) SerialNumber() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SoaRecordResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.SerialNumber
	}).(pulumi.Float64PtrOutput)
}

type SrvRecord struct {
	Port     *int    `pulumi:"port"`
	Priority *int    `pulumi:"priority"`
	Target   *string `pulumi:"target"`
	Weight   *int    `pulumi:"weight"`
}





type SrvRecordInput interface {
	pulumi.Input

	ToSrvRecordOutput() SrvRecordOutput
	ToSrvRecordOutputWithContext(context.Context) SrvRecordOutput
}

type SrvRecordArgs struct {
	Port     pulumi.IntPtrInput    `pulumi:"port"`
	Priority pulumi.IntPtrInput    `pulumi:"priority"`
	Target   pulumi.StringPtrInput `pulumi:"target"`
	Weight   pulumi.IntPtrInput    `pulumi:"weight"`
}

func (SrvRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SrvRecord)(nil)).Elem()
}

func (i SrvRecordArgs) ToSrvRecordOutput() SrvRecordOutput {
	return i.ToSrvRecordOutputWithContext(context.Background())
}

func (i SrvRecordArgs) ToSrvRecordOutputWithContext(ctx context.Context) SrvRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SrvRecordOutput)
}





type SrvRecordArrayInput interface {
	pulumi.Input

	ToSrvRecordArrayOutput() SrvRecordArrayOutput
	ToSrvRecordArrayOutputWithContext(context.Context) SrvRecordArrayOutput
}

type SrvRecordArray []SrvRecordInput

func (SrvRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SrvRecord)(nil)).Elem()
}

func (i SrvRecordArray) ToSrvRecordArrayOutput() SrvRecordArrayOutput {
	return i.ToSrvRecordArrayOutputWithContext(context.Background())
}

func (i SrvRecordArray) ToSrvRecordArrayOutputWithContext(ctx context.Context) SrvRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SrvRecordArrayOutput)
}

type SrvRecordOutput struct{ *pulumi.OutputState }

func (SrvRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SrvRecord)(nil)).Elem()
}

func (o SrvRecordOutput) ToSrvRecordOutput() SrvRecordOutput {
	return o
}

func (o SrvRecordOutput) ToSrvRecordOutputWithContext(ctx context.Context) SrvRecordOutput {
	return o
}

func (o SrvRecordOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecord) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o SrvRecordOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecord) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o SrvRecordOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SrvRecord) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o SrvRecordOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecord) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type SrvRecordArrayOutput struct{ *pulumi.OutputState }

func (SrvRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SrvRecord)(nil)).Elem()
}

func (o SrvRecordArrayOutput) ToSrvRecordArrayOutput() SrvRecordArrayOutput {
	return o
}

func (o SrvRecordArrayOutput) ToSrvRecordArrayOutputWithContext(ctx context.Context) SrvRecordArrayOutput {
	return o
}

func (o SrvRecordArrayOutput) Index(i pulumi.IntInput) SrvRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SrvRecord {
		return vs[0].([]SrvRecord)[vs[1].(int)]
	}).(SrvRecordOutput)
}

type SrvRecordResponse struct {
	Port     *int    `pulumi:"port"`
	Priority *int    `pulumi:"priority"`
	Target   *string `pulumi:"target"`
	Weight   *int    `pulumi:"weight"`
}

type SrvRecordResponseOutput struct{ *pulumi.OutputState }

func (SrvRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SrvRecordResponse)(nil)).Elem()
}

func (o SrvRecordResponseOutput) ToSrvRecordResponseOutput() SrvRecordResponseOutput {
	return o
}

func (o SrvRecordResponseOutput) ToSrvRecordResponseOutputWithContext(ctx context.Context) SrvRecordResponseOutput {
	return o
}

func (o SrvRecordResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecordResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o SrvRecordResponseOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecordResponse) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o SrvRecordResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SrvRecordResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o SrvRecordResponseOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SrvRecordResponse) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type SrvRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (SrvRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SrvRecordResponse)(nil)).Elem()
}

func (o SrvRecordResponseArrayOutput) ToSrvRecordResponseArrayOutput() SrvRecordResponseArrayOutput {
	return o
}

func (o SrvRecordResponseArrayOutput) ToSrvRecordResponseArrayOutputWithContext(ctx context.Context) SrvRecordResponseArrayOutput {
	return o
}

func (o SrvRecordResponseArrayOutput) Index(i pulumi.IntInput) SrvRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SrvRecordResponse {
		return vs[0].([]SrvRecordResponse)[vs[1].(int)]
	}).(SrvRecordResponseOutput)
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

func (i SubResourceArgs) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput).ToSubResourcePtrOutputWithContext(ctx)
}









type SubResourcePtrInput interface {
	pulumi.Input

	ToSubResourcePtrOutput() SubResourcePtrOutput
	ToSubResourcePtrOutputWithContext(context.Context) SubResourcePtrOutput
}

type subResourcePtrType SubResourceArgs

func SubResourcePtr(v *SubResourceArgs) SubResourcePtrInput {
	return (*subResourcePtrType)(v)
}

func (*subResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (i *subResourcePtrType) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i *subResourcePtrType) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourcePtrOutput)
}





type SubResourceArrayInput interface {
	pulumi.Input

	ToSubResourceArrayOutput() SubResourceArrayOutput
	ToSubResourceArrayOutputWithContext(context.Context) SubResourceArrayOutput
}

type SubResourceArray []SubResourceInput

func (SubResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (i SubResourceArray) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return i.ToSubResourceArrayOutputWithContext(context.Background())
}

func (i SubResourceArray) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceArrayOutput)
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

func (o SubResourceOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o.ToSubResourcePtrOutputWithContext(context.Background())
}

func (o SubResourceOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubResource) *SubResource {
		return &v
	}).(SubResourcePtrOutput)
}

func (o SubResourceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResource) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourcePtrOutput struct{ *pulumi.OutputState }

func (SubResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) Elem() SubResourceOutput {
	return o.ApplyT(func(v *SubResource) SubResource {
		if v != nil {
			return *v
		}
		var ret SubResource
		return ret
	}).(SubResourceOutput)
}

func (o SubResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResource) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceArrayOutput struct{ *pulumi.OutputState }

func (SubResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) Index(i pulumi.IntInput) SubResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResource {
		return vs[0].([]SubResource)[vs[1].(int)]
	}).(SubResourceOutput)
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

type SubResourceResponsePtrOutput struct{ *pulumi.OutputState }

func (SubResourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) Elem() SubResourceResponseOutput {
	return o.ApplyT(func(v *SubResourceResponse) SubResourceResponse {
		if v != nil {
			return *v
		}
		var ret SubResourceResponse
		return ret
	}).(SubResourceResponseOutput)
}

func (o SubResourceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SubResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutput() SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutputWithContext(ctx context.Context) SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) Index(i pulumi.IntInput) SubResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResourceResponse {
		return vs[0].([]SubResourceResponse)[vs[1].(int)]
	}).(SubResourceResponseOutput)
}

type TxtRecord struct {
	Value []string `pulumi:"value"`
}





type TxtRecordInput interface {
	pulumi.Input

	ToTxtRecordOutput() TxtRecordOutput
	ToTxtRecordOutputWithContext(context.Context) TxtRecordOutput
}

type TxtRecordArgs struct {
	Value pulumi.StringArrayInput `pulumi:"value"`
}

func (TxtRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecord)(nil)).Elem()
}

func (i TxtRecordArgs) ToTxtRecordOutput() TxtRecordOutput {
	return i.ToTxtRecordOutputWithContext(context.Background())
}

func (i TxtRecordArgs) ToTxtRecordOutputWithContext(ctx context.Context) TxtRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordOutput)
}





type TxtRecordArrayInput interface {
	pulumi.Input

	ToTxtRecordArrayOutput() TxtRecordArrayOutput
	ToTxtRecordArrayOutputWithContext(context.Context) TxtRecordArrayOutput
}

type TxtRecordArray []TxtRecordInput

func (TxtRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TxtRecord)(nil)).Elem()
}

func (i TxtRecordArray) ToTxtRecordArrayOutput() TxtRecordArrayOutput {
	return i.ToTxtRecordArrayOutputWithContext(context.Background())
}

func (i TxtRecordArray) ToTxtRecordArrayOutputWithContext(ctx context.Context) TxtRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordArrayOutput)
}

type TxtRecordOutput struct{ *pulumi.OutputState }

func (TxtRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecord)(nil)).Elem()
}

func (o TxtRecordOutput) ToTxtRecordOutput() TxtRecordOutput {
	return o
}

func (o TxtRecordOutput) ToTxtRecordOutputWithContext(ctx context.Context) TxtRecordOutput {
	return o
}

func (o TxtRecordOutput) Value() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TxtRecord) []string { return v.Value }).(pulumi.StringArrayOutput)
}

type TxtRecordArrayOutput struct{ *pulumi.OutputState }

func (TxtRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TxtRecord)(nil)).Elem()
}

func (o TxtRecordArrayOutput) ToTxtRecordArrayOutput() TxtRecordArrayOutput {
	return o
}

func (o TxtRecordArrayOutput) ToTxtRecordArrayOutputWithContext(ctx context.Context) TxtRecordArrayOutput {
	return o
}

func (o TxtRecordArrayOutput) Index(i pulumi.IntInput) TxtRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TxtRecord {
		return vs[0].([]TxtRecord)[vs[1].(int)]
	}).(TxtRecordOutput)
}

type TxtRecordResponse struct {
	Value []string `pulumi:"value"`
}

type TxtRecordResponseOutput struct{ *pulumi.OutputState }

func (TxtRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecordResponse)(nil)).Elem()
}

func (o TxtRecordResponseOutput) ToTxtRecordResponseOutput() TxtRecordResponseOutput {
	return o
}

func (o TxtRecordResponseOutput) ToTxtRecordResponseOutputWithContext(ctx context.Context) TxtRecordResponseOutput {
	return o
}

func (o TxtRecordResponseOutput) Value() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TxtRecordResponse) []string { return v.Value }).(pulumi.StringArrayOutput)
}

type TxtRecordResponseArrayOutput struct{ *pulumi.OutputState }

func (TxtRecordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TxtRecordResponse)(nil)).Elem()
}

func (o TxtRecordResponseArrayOutput) ToTxtRecordResponseArrayOutput() TxtRecordResponseArrayOutput {
	return o
}

func (o TxtRecordResponseArrayOutput) ToTxtRecordResponseArrayOutputWithContext(ctx context.Context) TxtRecordResponseArrayOutput {
	return o
}

func (o TxtRecordResponseArrayOutput) Index(i pulumi.IntInput) TxtRecordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TxtRecordResponse {
		return vs[0].([]TxtRecordResponse)[vs[1].(int)]
	}).(TxtRecordResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ARecordOutput{})
	pulumi.RegisterOutputType(ARecordArrayOutput{})
	pulumi.RegisterOutputType(ARecordResponseOutput{})
	pulumi.RegisterOutputType(ARecordResponseArrayOutput{})
	pulumi.RegisterOutputType(AaaaRecordOutput{})
	pulumi.RegisterOutputType(AaaaRecordArrayOutput{})
	pulumi.RegisterOutputType(AaaaRecordResponseOutput{})
	pulumi.RegisterOutputType(AaaaRecordResponseArrayOutput{})
	pulumi.RegisterOutputType(BackendOutput{})
	pulumi.RegisterOutputType(BackendArrayOutput{})
	pulumi.RegisterOutputType(BackendPoolOutput{})
	pulumi.RegisterOutputType(BackendPoolArrayOutput{})
	pulumi.RegisterOutputType(BackendPoolResponseOutput{})
	pulumi.RegisterOutputType(BackendPoolResponseArrayOutput{})
	pulumi.RegisterOutputType(BackendPoolsSettingsOutput{})
	pulumi.RegisterOutputType(BackendPoolsSettingsPtrOutput{})
	pulumi.RegisterOutputType(BackendPoolsSettingsResponseOutput{})
	pulumi.RegisterOutputType(BackendPoolsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(BackendResponseOutput{})
	pulumi.RegisterOutputType(BackendResponseArrayOutput{})
	pulumi.RegisterOutputType(CnameRecordOutput{})
	pulumi.RegisterOutputType(CnameRecordPtrOutput{})
	pulumi.RegisterOutputType(CnameRecordResponseOutput{})
	pulumi.RegisterOutputType(CnameRecordResponsePtrOutput{})
	pulumi.RegisterOutputType(CustomHttpsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(FrontendEndpointOutput{})
	pulumi.RegisterOutputType(FrontendEndpointArrayOutput{})
	pulumi.RegisterOutputType(FrontendEndpointResponseOutput{})
	pulumi.RegisterOutputType(FrontendEndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkOutput{})
	pulumi.RegisterOutputType(FrontendEndpointUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput{})
	pulumi.RegisterOutputType(FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkOutput{})
	pulumi.RegisterOutputType(FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLinkPtrOutput{})
	pulumi.RegisterOutputType(HeaderActionOutput{})
	pulumi.RegisterOutputType(HeaderActionArrayOutput{})
	pulumi.RegisterOutputType(HeaderActionResponseOutput{})
	pulumi.RegisterOutputType(HeaderActionResponseArrayOutput{})
	pulumi.RegisterOutputType(HealthProbeSettingsModelOutput{})
	pulumi.RegisterOutputType(HealthProbeSettingsModelArrayOutput{})
	pulumi.RegisterOutputType(HealthProbeSettingsModelResponseOutput{})
	pulumi.RegisterOutputType(HealthProbeSettingsModelResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultCertificateSourceParametersResponseVaultOutput{})
	pulumi.RegisterOutputType(KeyVaultCertificateSourceParametersResponseVaultPtrOutput{})
	pulumi.RegisterOutputType(LoadBalancingSettingsModelOutput{})
	pulumi.RegisterOutputType(LoadBalancingSettingsModelArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancingSettingsModelResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancingSettingsModelResponseArrayOutput{})
	pulumi.RegisterOutputType(MxRecordOutput{})
	pulumi.RegisterOutputType(MxRecordArrayOutput{})
	pulumi.RegisterOutputType(MxRecordResponseOutput{})
	pulumi.RegisterOutputType(MxRecordResponseArrayOutput{})
	pulumi.RegisterOutputType(PtrRecordOutput{})
	pulumi.RegisterOutputType(PtrRecordArrayOutput{})
	pulumi.RegisterOutputType(PtrRecordResponseOutput{})
	pulumi.RegisterOutputType(PtrRecordResponseArrayOutput{})
	pulumi.RegisterOutputType(RoutingRuleOutput{})
	pulumi.RegisterOutputType(RoutingRuleArrayOutput{})
	pulumi.RegisterOutputType(RoutingRuleResponseOutput{})
	pulumi.RegisterOutputType(RoutingRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(RulesEngineActionOutput{})
	pulumi.RegisterOutputType(RulesEngineActionResponseOutput{})
	pulumi.RegisterOutputType(RulesEngineMatchConditionOutput{})
	pulumi.RegisterOutputType(RulesEngineMatchConditionArrayOutput{})
	pulumi.RegisterOutputType(RulesEngineMatchConditionResponseOutput{})
	pulumi.RegisterOutputType(RulesEngineMatchConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(RulesEngineResponseOutput{})
	pulumi.RegisterOutputType(RulesEngineResponseArrayOutput{})
	pulumi.RegisterOutputType(RulesEngineRuleOutput{})
	pulumi.RegisterOutputType(RulesEngineRuleArrayOutput{})
	pulumi.RegisterOutputType(RulesEngineRuleResponseOutput{})
	pulumi.RegisterOutputType(RulesEngineRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(SoaRecordOutput{})
	pulumi.RegisterOutputType(SoaRecordPtrOutput{})
	pulumi.RegisterOutputType(SoaRecordResponseOutput{})
	pulumi.RegisterOutputType(SoaRecordResponsePtrOutput{})
	pulumi.RegisterOutputType(SrvRecordOutput{})
	pulumi.RegisterOutputType(SrvRecordArrayOutput{})
	pulumi.RegisterOutputType(SrvRecordResponseOutput{})
	pulumi.RegisterOutputType(SrvRecordResponseArrayOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourcePtrOutput{})
	pulumi.RegisterOutputType(SubResourceArrayOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(TxtRecordOutput{})
	pulumi.RegisterOutputType(TxtRecordArrayOutput{})
	pulumi.RegisterOutputType(TxtRecordResponseOutput{})
	pulumi.RegisterOutputType(TxtRecordResponseArrayOutput{})
}
