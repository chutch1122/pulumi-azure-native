


package v20200101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SiteAgentProperties struct {
	KeyVaultId  *string `pulumi:"keyVaultId"`
	KeyVaultUri *string `pulumi:"keyVaultUri"`
}





type SiteAgentPropertiesInput interface {
	pulumi.Input

	ToSiteAgentPropertiesOutput() SiteAgentPropertiesOutput
	ToSiteAgentPropertiesOutputWithContext(context.Context) SiteAgentPropertiesOutput
}

type SiteAgentPropertiesArgs struct {
	KeyVaultId  pulumi.StringPtrInput `pulumi:"keyVaultId"`
	KeyVaultUri pulumi.StringPtrInput `pulumi:"keyVaultUri"`
}

func (SiteAgentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteAgentProperties)(nil)).Elem()
}

func (i SiteAgentPropertiesArgs) ToSiteAgentPropertiesOutput() SiteAgentPropertiesOutput {
	return i.ToSiteAgentPropertiesOutputWithContext(context.Background())
}

func (i SiteAgentPropertiesArgs) ToSiteAgentPropertiesOutputWithContext(ctx context.Context) SiteAgentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteAgentPropertiesOutput)
}

func (i SiteAgentPropertiesArgs) ToSiteAgentPropertiesPtrOutput() SiteAgentPropertiesPtrOutput {
	return i.ToSiteAgentPropertiesPtrOutputWithContext(context.Background())
}

func (i SiteAgentPropertiesArgs) ToSiteAgentPropertiesPtrOutputWithContext(ctx context.Context) SiteAgentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteAgentPropertiesOutput).ToSiteAgentPropertiesPtrOutputWithContext(ctx)
}









type SiteAgentPropertiesPtrInput interface {
	pulumi.Input

	ToSiteAgentPropertiesPtrOutput() SiteAgentPropertiesPtrOutput
	ToSiteAgentPropertiesPtrOutputWithContext(context.Context) SiteAgentPropertiesPtrOutput
}

type siteAgentPropertiesPtrType SiteAgentPropertiesArgs

func SiteAgentPropertiesPtr(v *SiteAgentPropertiesArgs) SiteAgentPropertiesPtrInput {
	return (*siteAgentPropertiesPtrType)(v)
}

func (*siteAgentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteAgentProperties)(nil)).Elem()
}

func (i *siteAgentPropertiesPtrType) ToSiteAgentPropertiesPtrOutput() SiteAgentPropertiesPtrOutput {
	return i.ToSiteAgentPropertiesPtrOutputWithContext(context.Background())
}

func (i *siteAgentPropertiesPtrType) ToSiteAgentPropertiesPtrOutputWithContext(ctx context.Context) SiteAgentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteAgentPropertiesPtrOutput)
}

type SiteAgentPropertiesOutput struct{ *pulumi.OutputState }

func (SiteAgentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteAgentProperties)(nil)).Elem()
}

func (o SiteAgentPropertiesOutput) ToSiteAgentPropertiesOutput() SiteAgentPropertiesOutput {
	return o
}

func (o SiteAgentPropertiesOutput) ToSiteAgentPropertiesOutputWithContext(ctx context.Context) SiteAgentPropertiesOutput {
	return o
}

func (o SiteAgentPropertiesOutput) ToSiteAgentPropertiesPtrOutput() SiteAgentPropertiesPtrOutput {
	return o.ToSiteAgentPropertiesPtrOutputWithContext(context.Background())
}

func (o SiteAgentPropertiesOutput) ToSiteAgentPropertiesPtrOutputWithContext(ctx context.Context) SiteAgentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteAgentProperties) *SiteAgentProperties {
		return &v
	}).(SiteAgentPropertiesPtrOutput)
}

func (o SiteAgentPropertiesOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteAgentProperties) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

func (o SiteAgentPropertiesOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteAgentProperties) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

type SiteAgentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SiteAgentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteAgentProperties)(nil)).Elem()
}

func (o SiteAgentPropertiesPtrOutput) ToSiteAgentPropertiesPtrOutput() SiteAgentPropertiesPtrOutput {
	return o
}

func (o SiteAgentPropertiesPtrOutput) ToSiteAgentPropertiesPtrOutputWithContext(ctx context.Context) SiteAgentPropertiesPtrOutput {
	return o
}

func (o SiteAgentPropertiesPtrOutput) Elem() SiteAgentPropertiesOutput {
	return o.ApplyT(func(v *SiteAgentProperties) SiteAgentProperties {
		if v != nil {
			return *v
		}
		var ret SiteAgentProperties
		return ret
	}).(SiteAgentPropertiesOutput)
}

func (o SiteAgentPropertiesPtrOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAgentProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultId
	}).(pulumi.StringPtrOutput)
}

func (o SiteAgentPropertiesPtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAgentProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

type SiteAgentPropertiesResponse struct {
	Id               string  `pulumi:"id"`
	KeyVaultId       *string `pulumi:"keyVaultId"`
	KeyVaultUri      *string `pulumi:"keyVaultUri"`
	LastHeartBeatUtc string  `pulumi:"lastHeartBeatUtc"`
	Version          string  `pulumi:"version"`
}

type SiteAgentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SiteAgentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteAgentPropertiesResponse)(nil)).Elem()
}

func (o SiteAgentPropertiesResponseOutput) ToSiteAgentPropertiesResponseOutput() SiteAgentPropertiesResponseOutput {
	return o
}

func (o SiteAgentPropertiesResponseOutput) ToSiteAgentPropertiesResponseOutputWithContext(ctx context.Context) SiteAgentPropertiesResponseOutput {
	return o
}

func (o SiteAgentPropertiesResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SiteAgentPropertiesResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o SiteAgentPropertiesResponseOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteAgentPropertiesResponse) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

func (o SiteAgentPropertiesResponseOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteAgentPropertiesResponse) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o SiteAgentPropertiesResponseOutput) LastHeartBeatUtc() pulumi.StringOutput {
	return o.ApplyT(func(v SiteAgentPropertiesResponse) string { return v.LastHeartBeatUtc }).(pulumi.StringOutput)
}

func (o SiteAgentPropertiesResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v SiteAgentPropertiesResponse) string { return v.Version }).(pulumi.StringOutput)
}

type SiteAgentPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SiteAgentPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteAgentPropertiesResponse)(nil)).Elem()
}

func (o SiteAgentPropertiesResponsePtrOutput) ToSiteAgentPropertiesResponsePtrOutput() SiteAgentPropertiesResponsePtrOutput {
	return o
}

func (o SiteAgentPropertiesResponsePtrOutput) ToSiteAgentPropertiesResponsePtrOutputWithContext(ctx context.Context) SiteAgentPropertiesResponsePtrOutput {
	return o
}

func (o SiteAgentPropertiesResponsePtrOutput) Elem() SiteAgentPropertiesResponseOutput {
	return o.ApplyT(func(v *SiteAgentPropertiesResponse) SiteAgentPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SiteAgentPropertiesResponse
		return ret
	}).(SiteAgentPropertiesResponseOutput)
}

func (o SiteAgentPropertiesResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAgentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SiteAgentPropertiesResponsePtrOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAgentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultId
	}).(pulumi.StringPtrOutput)
}

func (o SiteAgentPropertiesResponsePtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAgentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o SiteAgentPropertiesResponsePtrOutput) LastHeartBeatUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAgentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastHeartBeatUtc
	}).(pulumi.StringPtrOutput)
}

func (o SiteAgentPropertiesResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAgentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type SiteProperties struct {
	AgentDetails                    *SiteAgentProperties `pulumi:"agentDetails"`
	ApplianceName                   *string              `pulumi:"applianceName"`
	DiscoverySolutionId             *string              `pulumi:"discoverySolutionId"`
	ServicePrincipalIdentityDetails *SiteSpnProperties   `pulumi:"servicePrincipalIdentityDetails"`
}





type SitePropertiesInput interface {
	pulumi.Input

	ToSitePropertiesOutput() SitePropertiesOutput
	ToSitePropertiesOutputWithContext(context.Context) SitePropertiesOutput
}

type SitePropertiesArgs struct {
	AgentDetails                    SiteAgentPropertiesPtrInput `pulumi:"agentDetails"`
	ApplianceName                   pulumi.StringPtrInput       `pulumi:"applianceName"`
	DiscoverySolutionId             pulumi.StringPtrInput       `pulumi:"discoverySolutionId"`
	ServicePrincipalIdentityDetails SiteSpnPropertiesPtrInput   `pulumi:"servicePrincipalIdentityDetails"`
}

func (SitePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteProperties)(nil)).Elem()
}

func (i SitePropertiesArgs) ToSitePropertiesOutput() SitePropertiesOutput {
	return i.ToSitePropertiesOutputWithContext(context.Background())
}

func (i SitePropertiesArgs) ToSitePropertiesOutputWithContext(ctx context.Context) SitePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SitePropertiesOutput)
}

func (i SitePropertiesArgs) ToSitePropertiesPtrOutput() SitePropertiesPtrOutput {
	return i.ToSitePropertiesPtrOutputWithContext(context.Background())
}

func (i SitePropertiesArgs) ToSitePropertiesPtrOutputWithContext(ctx context.Context) SitePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SitePropertiesOutput).ToSitePropertiesPtrOutputWithContext(ctx)
}









type SitePropertiesPtrInput interface {
	pulumi.Input

	ToSitePropertiesPtrOutput() SitePropertiesPtrOutput
	ToSitePropertiesPtrOutputWithContext(context.Context) SitePropertiesPtrOutput
}

type sitePropertiesPtrType SitePropertiesArgs

func SitePropertiesPtr(v *SitePropertiesArgs) SitePropertiesPtrInput {
	return (*sitePropertiesPtrType)(v)
}

func (*sitePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteProperties)(nil)).Elem()
}

func (i *sitePropertiesPtrType) ToSitePropertiesPtrOutput() SitePropertiesPtrOutput {
	return i.ToSitePropertiesPtrOutputWithContext(context.Background())
}

func (i *sitePropertiesPtrType) ToSitePropertiesPtrOutputWithContext(ctx context.Context) SitePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SitePropertiesPtrOutput)
}

type SitePropertiesOutput struct{ *pulumi.OutputState }

func (SitePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteProperties)(nil)).Elem()
}

func (o SitePropertiesOutput) ToSitePropertiesOutput() SitePropertiesOutput {
	return o
}

func (o SitePropertiesOutput) ToSitePropertiesOutputWithContext(ctx context.Context) SitePropertiesOutput {
	return o
}

func (o SitePropertiesOutput) ToSitePropertiesPtrOutput() SitePropertiesPtrOutput {
	return o.ToSitePropertiesPtrOutputWithContext(context.Background())
}

func (o SitePropertiesOutput) ToSitePropertiesPtrOutputWithContext(ctx context.Context) SitePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteProperties) *SiteProperties {
		return &v
	}).(SitePropertiesPtrOutput)
}

func (o SitePropertiesOutput) AgentDetails() SiteAgentPropertiesPtrOutput {
	return o.ApplyT(func(v SiteProperties) *SiteAgentProperties { return v.AgentDetails }).(SiteAgentPropertiesPtrOutput)
}

func (o SitePropertiesOutput) ApplianceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteProperties) *string { return v.ApplianceName }).(pulumi.StringPtrOutput)
}

func (o SitePropertiesOutput) DiscoverySolutionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteProperties) *string { return v.DiscoverySolutionId }).(pulumi.StringPtrOutput)
}

func (o SitePropertiesOutput) ServicePrincipalIdentityDetails() SiteSpnPropertiesPtrOutput {
	return o.ApplyT(func(v SiteProperties) *SiteSpnProperties { return v.ServicePrincipalIdentityDetails }).(SiteSpnPropertiesPtrOutput)
}

type SitePropertiesPtrOutput struct{ *pulumi.OutputState }

func (SitePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteProperties)(nil)).Elem()
}

func (o SitePropertiesPtrOutput) ToSitePropertiesPtrOutput() SitePropertiesPtrOutput {
	return o
}

func (o SitePropertiesPtrOutput) ToSitePropertiesPtrOutputWithContext(ctx context.Context) SitePropertiesPtrOutput {
	return o
}

func (o SitePropertiesPtrOutput) Elem() SitePropertiesOutput {
	return o.ApplyT(func(v *SiteProperties) SiteProperties {
		if v != nil {
			return *v
		}
		var ret SiteProperties
		return ret
	}).(SitePropertiesOutput)
}

func (o SitePropertiesPtrOutput) AgentDetails() SiteAgentPropertiesPtrOutput {
	return o.ApplyT(func(v *SiteProperties) *SiteAgentProperties {
		if v == nil {
			return nil
		}
		return v.AgentDetails
	}).(SiteAgentPropertiesPtrOutput)
}

func (o SitePropertiesPtrOutput) ApplianceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApplianceName
	}).(pulumi.StringPtrOutput)
}

func (o SitePropertiesPtrOutput) DiscoverySolutionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteProperties) *string {
		if v == nil {
			return nil
		}
		return v.DiscoverySolutionId
	}).(pulumi.StringPtrOutput)
}

func (o SitePropertiesPtrOutput) ServicePrincipalIdentityDetails() SiteSpnPropertiesPtrOutput {
	return o.ApplyT(func(v *SiteProperties) *SiteSpnProperties {
		if v == nil {
			return nil
		}
		return v.ServicePrincipalIdentityDetails
	}).(SiteSpnPropertiesPtrOutput)
}

type SitePropertiesResponse struct {
	AgentDetails                    *SiteAgentPropertiesResponse `pulumi:"agentDetails"`
	ApplianceName                   *string                      `pulumi:"applianceName"`
	DiscoverySolutionId             *string                      `pulumi:"discoverySolutionId"`
	ServiceEndpoint                 string                       `pulumi:"serviceEndpoint"`
	ServicePrincipalIdentityDetails *SiteSpnPropertiesResponse   `pulumi:"servicePrincipalIdentityDetails"`
}

type SitePropertiesResponseOutput struct{ *pulumi.OutputState }

func (SitePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SitePropertiesResponse)(nil)).Elem()
}

func (o SitePropertiesResponseOutput) ToSitePropertiesResponseOutput() SitePropertiesResponseOutput {
	return o
}

func (o SitePropertiesResponseOutput) ToSitePropertiesResponseOutputWithContext(ctx context.Context) SitePropertiesResponseOutput {
	return o
}

func (o SitePropertiesResponseOutput) AgentDetails() SiteAgentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v SitePropertiesResponse) *SiteAgentPropertiesResponse { return v.AgentDetails }).(SiteAgentPropertiesResponsePtrOutput)
}

func (o SitePropertiesResponseOutput) ApplianceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SitePropertiesResponse) *string { return v.ApplianceName }).(pulumi.StringPtrOutput)
}

func (o SitePropertiesResponseOutput) DiscoverySolutionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SitePropertiesResponse) *string { return v.DiscoverySolutionId }).(pulumi.StringPtrOutput)
}

func (o SitePropertiesResponseOutput) ServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v SitePropertiesResponse) string { return v.ServiceEndpoint }).(pulumi.StringOutput)
}

func (o SitePropertiesResponseOutput) ServicePrincipalIdentityDetails() SiteSpnPropertiesResponsePtrOutput {
	return o.ApplyT(func(v SitePropertiesResponse) *SiteSpnPropertiesResponse { return v.ServicePrincipalIdentityDetails }).(SiteSpnPropertiesResponsePtrOutput)
}

type SiteSpnProperties struct {
	AadAuthority  *string `pulumi:"aadAuthority"`
	ApplicationId *string `pulumi:"applicationId"`
	Audience      *string `pulumi:"audience"`
	ObjectId      *string `pulumi:"objectId"`
	RawCertData   *string `pulumi:"rawCertData"`
	TenantId      *string `pulumi:"tenantId"`
}





type SiteSpnPropertiesInput interface {
	pulumi.Input

	ToSiteSpnPropertiesOutput() SiteSpnPropertiesOutput
	ToSiteSpnPropertiesOutputWithContext(context.Context) SiteSpnPropertiesOutput
}

type SiteSpnPropertiesArgs struct {
	AadAuthority  pulumi.StringPtrInput `pulumi:"aadAuthority"`
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
	Audience      pulumi.StringPtrInput `pulumi:"audience"`
	ObjectId      pulumi.StringPtrInput `pulumi:"objectId"`
	RawCertData   pulumi.StringPtrInput `pulumi:"rawCertData"`
	TenantId      pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (SiteSpnPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteSpnProperties)(nil)).Elem()
}

func (i SiteSpnPropertiesArgs) ToSiteSpnPropertiesOutput() SiteSpnPropertiesOutput {
	return i.ToSiteSpnPropertiesOutputWithContext(context.Background())
}

func (i SiteSpnPropertiesArgs) ToSiteSpnPropertiesOutputWithContext(ctx context.Context) SiteSpnPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteSpnPropertiesOutput)
}

func (i SiteSpnPropertiesArgs) ToSiteSpnPropertiesPtrOutput() SiteSpnPropertiesPtrOutput {
	return i.ToSiteSpnPropertiesPtrOutputWithContext(context.Background())
}

func (i SiteSpnPropertiesArgs) ToSiteSpnPropertiesPtrOutputWithContext(ctx context.Context) SiteSpnPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteSpnPropertiesOutput).ToSiteSpnPropertiesPtrOutputWithContext(ctx)
}









type SiteSpnPropertiesPtrInput interface {
	pulumi.Input

	ToSiteSpnPropertiesPtrOutput() SiteSpnPropertiesPtrOutput
	ToSiteSpnPropertiesPtrOutputWithContext(context.Context) SiteSpnPropertiesPtrOutput
}

type siteSpnPropertiesPtrType SiteSpnPropertiesArgs

func SiteSpnPropertiesPtr(v *SiteSpnPropertiesArgs) SiteSpnPropertiesPtrInput {
	return (*siteSpnPropertiesPtrType)(v)
}

func (*siteSpnPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteSpnProperties)(nil)).Elem()
}

func (i *siteSpnPropertiesPtrType) ToSiteSpnPropertiesPtrOutput() SiteSpnPropertiesPtrOutput {
	return i.ToSiteSpnPropertiesPtrOutputWithContext(context.Background())
}

func (i *siteSpnPropertiesPtrType) ToSiteSpnPropertiesPtrOutputWithContext(ctx context.Context) SiteSpnPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteSpnPropertiesPtrOutput)
}

type SiteSpnPropertiesOutput struct{ *pulumi.OutputState }

func (SiteSpnPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteSpnProperties)(nil)).Elem()
}

func (o SiteSpnPropertiesOutput) ToSiteSpnPropertiesOutput() SiteSpnPropertiesOutput {
	return o
}

func (o SiteSpnPropertiesOutput) ToSiteSpnPropertiesOutputWithContext(ctx context.Context) SiteSpnPropertiesOutput {
	return o
}

func (o SiteSpnPropertiesOutput) ToSiteSpnPropertiesPtrOutput() SiteSpnPropertiesPtrOutput {
	return o.ToSiteSpnPropertiesPtrOutputWithContext(context.Background())
}

func (o SiteSpnPropertiesOutput) ToSiteSpnPropertiesPtrOutputWithContext(ctx context.Context) SiteSpnPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteSpnProperties) *SiteSpnProperties {
		return &v
	}).(SiteSpnPropertiesPtrOutput)
}

func (o SiteSpnPropertiesOutput) AadAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteSpnProperties) *string { return v.AadAuthority }).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteSpnProperties) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteSpnProperties) *string { return v.Audience }).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteSpnProperties) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesOutput) RawCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteSpnProperties) *string { return v.RawCertData }).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteSpnProperties) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type SiteSpnPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SiteSpnPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteSpnProperties)(nil)).Elem()
}

func (o SiteSpnPropertiesPtrOutput) ToSiteSpnPropertiesPtrOutput() SiteSpnPropertiesPtrOutput {
	return o
}

func (o SiteSpnPropertiesPtrOutput) ToSiteSpnPropertiesPtrOutputWithContext(ctx context.Context) SiteSpnPropertiesPtrOutput {
	return o
}

func (o SiteSpnPropertiesPtrOutput) Elem() SiteSpnPropertiesOutput {
	return o.ApplyT(func(v *SiteSpnProperties) SiteSpnProperties {
		if v != nil {
			return *v
		}
		var ret SiteSpnProperties
		return ret
	}).(SiteSpnPropertiesOutput)
}

func (o SiteSpnPropertiesPtrOutput) AadAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteSpnProperties) *string {
		if v == nil {
			return nil
		}
		return v.AadAuthority
	}).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesPtrOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteSpnProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesPtrOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteSpnProperties) *string {
		if v == nil {
			return nil
		}
		return v.Audience
	}).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesPtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteSpnProperties) *string {
		if v == nil {
			return nil
		}
		return v.ObjectId
	}).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesPtrOutput) RawCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteSpnProperties) *string {
		if v == nil {
			return nil
		}
		return v.RawCertData
	}).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteSpnProperties) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type SiteSpnPropertiesResponse struct {
	AadAuthority  *string `pulumi:"aadAuthority"`
	ApplicationId *string `pulumi:"applicationId"`
	Audience      *string `pulumi:"audience"`
	ObjectId      *string `pulumi:"objectId"`
	RawCertData   *string `pulumi:"rawCertData"`
	TenantId      *string `pulumi:"tenantId"`
}

type SiteSpnPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SiteSpnPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteSpnPropertiesResponse)(nil)).Elem()
}

func (o SiteSpnPropertiesResponseOutput) ToSiteSpnPropertiesResponseOutput() SiteSpnPropertiesResponseOutput {
	return o
}

func (o SiteSpnPropertiesResponseOutput) ToSiteSpnPropertiesResponseOutputWithContext(ctx context.Context) SiteSpnPropertiesResponseOutput {
	return o
}

func (o SiteSpnPropertiesResponseOutput) AadAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteSpnPropertiesResponse) *string { return v.AadAuthority }).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesResponseOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteSpnPropertiesResponse) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesResponseOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteSpnPropertiesResponse) *string { return v.Audience }).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesResponseOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteSpnPropertiesResponse) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesResponseOutput) RawCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteSpnPropertiesResponse) *string { return v.RawCertData }).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteSpnPropertiesResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type SiteSpnPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SiteSpnPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteSpnPropertiesResponse)(nil)).Elem()
}

func (o SiteSpnPropertiesResponsePtrOutput) ToSiteSpnPropertiesResponsePtrOutput() SiteSpnPropertiesResponsePtrOutput {
	return o
}

func (o SiteSpnPropertiesResponsePtrOutput) ToSiteSpnPropertiesResponsePtrOutputWithContext(ctx context.Context) SiteSpnPropertiesResponsePtrOutput {
	return o
}

func (o SiteSpnPropertiesResponsePtrOutput) Elem() SiteSpnPropertiesResponseOutput {
	return o.ApplyT(func(v *SiteSpnPropertiesResponse) SiteSpnPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SiteSpnPropertiesResponse
		return ret
	}).(SiteSpnPropertiesResponseOutput)
}

func (o SiteSpnPropertiesResponsePtrOutput) AadAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteSpnPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AadAuthority
	}).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesResponsePtrOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteSpnPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesResponsePtrOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteSpnPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Audience
	}).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesResponsePtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteSpnPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ObjectId
	}).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesResponsePtrOutput) RawCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteSpnPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RawCertData
	}).(pulumi.StringPtrOutput)
}

func (o SiteSpnPropertiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteSpnPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SiteAgentPropertiesOutput{})
	pulumi.RegisterOutputType(SiteAgentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SiteAgentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SiteAgentPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SitePropertiesOutput{})
	pulumi.RegisterOutputType(SitePropertiesPtrOutput{})
	pulumi.RegisterOutputType(SitePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SiteSpnPropertiesOutput{})
	pulumi.RegisterOutputType(SiteSpnPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SiteSpnPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SiteSpnPropertiesResponsePtrOutput{})
}
