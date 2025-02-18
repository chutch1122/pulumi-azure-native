


package v20210201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AllowedAudiencesValidation struct {
	AllowedAudiences []string `pulumi:"allowedAudiences"`
}





type AllowedAudiencesValidationInput interface {
	pulumi.Input

	ToAllowedAudiencesValidationOutput() AllowedAudiencesValidationOutput
	ToAllowedAudiencesValidationOutputWithContext(context.Context) AllowedAudiencesValidationOutput
}

type AllowedAudiencesValidationArgs struct {
	AllowedAudiences pulumi.StringArrayInput `pulumi:"allowedAudiences"`
}

func (AllowedAudiencesValidationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowedAudiencesValidation)(nil)).Elem()
}

func (i AllowedAudiencesValidationArgs) ToAllowedAudiencesValidationOutput() AllowedAudiencesValidationOutput {
	return i.ToAllowedAudiencesValidationOutputWithContext(context.Background())
}

func (i AllowedAudiencesValidationArgs) ToAllowedAudiencesValidationOutputWithContext(ctx context.Context) AllowedAudiencesValidationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowedAudiencesValidationOutput)
}

func (i AllowedAudiencesValidationArgs) ToAllowedAudiencesValidationPtrOutput() AllowedAudiencesValidationPtrOutput {
	return i.ToAllowedAudiencesValidationPtrOutputWithContext(context.Background())
}

func (i AllowedAudiencesValidationArgs) ToAllowedAudiencesValidationPtrOutputWithContext(ctx context.Context) AllowedAudiencesValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowedAudiencesValidationOutput).ToAllowedAudiencesValidationPtrOutputWithContext(ctx)
}









type AllowedAudiencesValidationPtrInput interface {
	pulumi.Input

	ToAllowedAudiencesValidationPtrOutput() AllowedAudiencesValidationPtrOutput
	ToAllowedAudiencesValidationPtrOutputWithContext(context.Context) AllowedAudiencesValidationPtrOutput
}

type allowedAudiencesValidationPtrType AllowedAudiencesValidationArgs

func AllowedAudiencesValidationPtr(v *AllowedAudiencesValidationArgs) AllowedAudiencesValidationPtrInput {
	return (*allowedAudiencesValidationPtrType)(v)
}

func (*allowedAudiencesValidationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowedAudiencesValidation)(nil)).Elem()
}

func (i *allowedAudiencesValidationPtrType) ToAllowedAudiencesValidationPtrOutput() AllowedAudiencesValidationPtrOutput {
	return i.ToAllowedAudiencesValidationPtrOutputWithContext(context.Background())
}

func (i *allowedAudiencesValidationPtrType) ToAllowedAudiencesValidationPtrOutputWithContext(ctx context.Context) AllowedAudiencesValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowedAudiencesValidationPtrOutput)
}

type AllowedAudiencesValidationOutput struct{ *pulumi.OutputState }

func (AllowedAudiencesValidationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowedAudiencesValidation)(nil)).Elem()
}

func (o AllowedAudiencesValidationOutput) ToAllowedAudiencesValidationOutput() AllowedAudiencesValidationOutput {
	return o
}

func (o AllowedAudiencesValidationOutput) ToAllowedAudiencesValidationOutputWithContext(ctx context.Context) AllowedAudiencesValidationOutput {
	return o
}

func (o AllowedAudiencesValidationOutput) ToAllowedAudiencesValidationPtrOutput() AllowedAudiencesValidationPtrOutput {
	return o.ToAllowedAudiencesValidationPtrOutputWithContext(context.Background())
}

func (o AllowedAudiencesValidationOutput) ToAllowedAudiencesValidationPtrOutputWithContext(ctx context.Context) AllowedAudiencesValidationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AllowedAudiencesValidation) *AllowedAudiencesValidation {
		return &v
	}).(AllowedAudiencesValidationPtrOutput)
}

func (o AllowedAudiencesValidationOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AllowedAudiencesValidation) []string { return v.AllowedAudiences }).(pulumi.StringArrayOutput)
}

type AllowedAudiencesValidationPtrOutput struct{ *pulumi.OutputState }

func (AllowedAudiencesValidationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowedAudiencesValidation)(nil)).Elem()
}

func (o AllowedAudiencesValidationPtrOutput) ToAllowedAudiencesValidationPtrOutput() AllowedAudiencesValidationPtrOutput {
	return o
}

func (o AllowedAudiencesValidationPtrOutput) ToAllowedAudiencesValidationPtrOutputWithContext(ctx context.Context) AllowedAudiencesValidationPtrOutput {
	return o
}

func (o AllowedAudiencesValidationPtrOutput) Elem() AllowedAudiencesValidationOutput {
	return o.ApplyT(func(v *AllowedAudiencesValidation) AllowedAudiencesValidation {
		if v != nil {
			return *v
		}
		var ret AllowedAudiencesValidation
		return ret
	}).(AllowedAudiencesValidationOutput)
}

func (o AllowedAudiencesValidationPtrOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AllowedAudiencesValidation) []string {
		if v == nil {
			return nil
		}
		return v.AllowedAudiences
	}).(pulumi.StringArrayOutput)
}

type AllowedAudiencesValidationResponse struct {
	AllowedAudiences []string `pulumi:"allowedAudiences"`
}

type AllowedAudiencesValidationResponseOutput struct{ *pulumi.OutputState }

func (AllowedAudiencesValidationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowedAudiencesValidationResponse)(nil)).Elem()
}

func (o AllowedAudiencesValidationResponseOutput) ToAllowedAudiencesValidationResponseOutput() AllowedAudiencesValidationResponseOutput {
	return o
}

func (o AllowedAudiencesValidationResponseOutput) ToAllowedAudiencesValidationResponseOutputWithContext(ctx context.Context) AllowedAudiencesValidationResponseOutput {
	return o
}

func (o AllowedAudiencesValidationResponseOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AllowedAudiencesValidationResponse) []string { return v.AllowedAudiences }).(pulumi.StringArrayOutput)
}

type AllowedAudiencesValidationResponsePtrOutput struct{ *pulumi.OutputState }

func (AllowedAudiencesValidationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowedAudiencesValidationResponse)(nil)).Elem()
}

func (o AllowedAudiencesValidationResponsePtrOutput) ToAllowedAudiencesValidationResponsePtrOutput() AllowedAudiencesValidationResponsePtrOutput {
	return o
}

func (o AllowedAudiencesValidationResponsePtrOutput) ToAllowedAudiencesValidationResponsePtrOutputWithContext(ctx context.Context) AllowedAudiencesValidationResponsePtrOutput {
	return o
}

func (o AllowedAudiencesValidationResponsePtrOutput) Elem() AllowedAudiencesValidationResponseOutput {
	return o.ApplyT(func(v *AllowedAudiencesValidationResponse) AllowedAudiencesValidationResponse {
		if v != nil {
			return *v
		}
		var ret AllowedAudiencesValidationResponse
		return ret
	}).(AllowedAudiencesValidationResponseOutput)
}

func (o AllowedAudiencesValidationResponsePtrOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AllowedAudiencesValidationResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedAudiences
	}).(pulumi.StringArrayOutput)
}

type AllowedPrincipals struct {
	Groups     []string `pulumi:"groups"`
	Identities []string `pulumi:"identities"`
}





type AllowedPrincipalsInput interface {
	pulumi.Input

	ToAllowedPrincipalsOutput() AllowedPrincipalsOutput
	ToAllowedPrincipalsOutputWithContext(context.Context) AllowedPrincipalsOutput
}

type AllowedPrincipalsArgs struct {
	Groups     pulumi.StringArrayInput `pulumi:"groups"`
	Identities pulumi.StringArrayInput `pulumi:"identities"`
}

func (AllowedPrincipalsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowedPrincipals)(nil)).Elem()
}

func (i AllowedPrincipalsArgs) ToAllowedPrincipalsOutput() AllowedPrincipalsOutput {
	return i.ToAllowedPrincipalsOutputWithContext(context.Background())
}

func (i AllowedPrincipalsArgs) ToAllowedPrincipalsOutputWithContext(ctx context.Context) AllowedPrincipalsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowedPrincipalsOutput)
}

func (i AllowedPrincipalsArgs) ToAllowedPrincipalsPtrOutput() AllowedPrincipalsPtrOutput {
	return i.ToAllowedPrincipalsPtrOutputWithContext(context.Background())
}

func (i AllowedPrincipalsArgs) ToAllowedPrincipalsPtrOutputWithContext(ctx context.Context) AllowedPrincipalsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowedPrincipalsOutput).ToAllowedPrincipalsPtrOutputWithContext(ctx)
}









type AllowedPrincipalsPtrInput interface {
	pulumi.Input

	ToAllowedPrincipalsPtrOutput() AllowedPrincipalsPtrOutput
	ToAllowedPrincipalsPtrOutputWithContext(context.Context) AllowedPrincipalsPtrOutput
}

type allowedPrincipalsPtrType AllowedPrincipalsArgs

func AllowedPrincipalsPtr(v *AllowedPrincipalsArgs) AllowedPrincipalsPtrInput {
	return (*allowedPrincipalsPtrType)(v)
}

func (*allowedPrincipalsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowedPrincipals)(nil)).Elem()
}

func (i *allowedPrincipalsPtrType) ToAllowedPrincipalsPtrOutput() AllowedPrincipalsPtrOutput {
	return i.ToAllowedPrincipalsPtrOutputWithContext(context.Background())
}

func (i *allowedPrincipalsPtrType) ToAllowedPrincipalsPtrOutputWithContext(ctx context.Context) AllowedPrincipalsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowedPrincipalsPtrOutput)
}

type AllowedPrincipalsOutput struct{ *pulumi.OutputState }

func (AllowedPrincipalsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowedPrincipals)(nil)).Elem()
}

func (o AllowedPrincipalsOutput) ToAllowedPrincipalsOutput() AllowedPrincipalsOutput {
	return o
}

func (o AllowedPrincipalsOutput) ToAllowedPrincipalsOutputWithContext(ctx context.Context) AllowedPrincipalsOutput {
	return o
}

func (o AllowedPrincipalsOutput) ToAllowedPrincipalsPtrOutput() AllowedPrincipalsPtrOutput {
	return o.ToAllowedPrincipalsPtrOutputWithContext(context.Background())
}

func (o AllowedPrincipalsOutput) ToAllowedPrincipalsPtrOutputWithContext(ctx context.Context) AllowedPrincipalsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AllowedPrincipals) *AllowedPrincipals {
		return &v
	}).(AllowedPrincipalsPtrOutput)
}

func (o AllowedPrincipalsOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AllowedPrincipals) []string { return v.Groups }).(pulumi.StringArrayOutput)
}

func (o AllowedPrincipalsOutput) Identities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AllowedPrincipals) []string { return v.Identities }).(pulumi.StringArrayOutput)
}

type AllowedPrincipalsPtrOutput struct{ *pulumi.OutputState }

func (AllowedPrincipalsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowedPrincipals)(nil)).Elem()
}

func (o AllowedPrincipalsPtrOutput) ToAllowedPrincipalsPtrOutput() AllowedPrincipalsPtrOutput {
	return o
}

func (o AllowedPrincipalsPtrOutput) ToAllowedPrincipalsPtrOutputWithContext(ctx context.Context) AllowedPrincipalsPtrOutput {
	return o
}

func (o AllowedPrincipalsPtrOutput) Elem() AllowedPrincipalsOutput {
	return o.ApplyT(func(v *AllowedPrincipals) AllowedPrincipals {
		if v != nil {
			return *v
		}
		var ret AllowedPrincipals
		return ret
	}).(AllowedPrincipalsOutput)
}

func (o AllowedPrincipalsPtrOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AllowedPrincipals) []string {
		if v == nil {
			return nil
		}
		return v.Groups
	}).(pulumi.StringArrayOutput)
}

func (o AllowedPrincipalsPtrOutput) Identities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AllowedPrincipals) []string {
		if v == nil {
			return nil
		}
		return v.Identities
	}).(pulumi.StringArrayOutput)
}

type AllowedPrincipalsResponse struct {
	Groups     []string `pulumi:"groups"`
	Identities []string `pulumi:"identities"`
}

type AllowedPrincipalsResponseOutput struct{ *pulumi.OutputState }

func (AllowedPrincipalsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowedPrincipalsResponse)(nil)).Elem()
}

func (o AllowedPrincipalsResponseOutput) ToAllowedPrincipalsResponseOutput() AllowedPrincipalsResponseOutput {
	return o
}

func (o AllowedPrincipalsResponseOutput) ToAllowedPrincipalsResponseOutputWithContext(ctx context.Context) AllowedPrincipalsResponseOutput {
	return o
}

func (o AllowedPrincipalsResponseOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AllowedPrincipalsResponse) []string { return v.Groups }).(pulumi.StringArrayOutput)
}

func (o AllowedPrincipalsResponseOutput) Identities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AllowedPrincipalsResponse) []string { return v.Identities }).(pulumi.StringArrayOutput)
}

type AllowedPrincipalsResponsePtrOutput struct{ *pulumi.OutputState }

func (AllowedPrincipalsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowedPrincipalsResponse)(nil)).Elem()
}

func (o AllowedPrincipalsResponsePtrOutput) ToAllowedPrincipalsResponsePtrOutput() AllowedPrincipalsResponsePtrOutput {
	return o
}

func (o AllowedPrincipalsResponsePtrOutput) ToAllowedPrincipalsResponsePtrOutputWithContext(ctx context.Context) AllowedPrincipalsResponsePtrOutput {
	return o
}

func (o AllowedPrincipalsResponsePtrOutput) Elem() AllowedPrincipalsResponseOutput {
	return o.ApplyT(func(v *AllowedPrincipalsResponse) AllowedPrincipalsResponse {
		if v != nil {
			return *v
		}
		var ret AllowedPrincipalsResponse
		return ret
	}).(AllowedPrincipalsResponseOutput)
}

func (o AllowedPrincipalsResponsePtrOutput) Groups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AllowedPrincipalsResponse) []string {
		if v == nil {
			return nil
		}
		return v.Groups
	}).(pulumi.StringArrayOutput)
}

func (o AllowedPrincipalsResponsePtrOutput) Identities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AllowedPrincipalsResponse) []string {
		if v == nil {
			return nil
		}
		return v.Identities
	}).(pulumi.StringArrayOutput)
}

type ApiDefinitionInfo struct {
	Url *string `pulumi:"url"`
}





type ApiDefinitionInfoInput interface {
	pulumi.Input

	ToApiDefinitionInfoOutput() ApiDefinitionInfoOutput
	ToApiDefinitionInfoOutputWithContext(context.Context) ApiDefinitionInfoOutput
}

type ApiDefinitionInfoArgs struct {
	Url pulumi.StringPtrInput `pulumi:"url"`
}

func (ApiDefinitionInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiDefinitionInfo)(nil)).Elem()
}

func (i ApiDefinitionInfoArgs) ToApiDefinitionInfoOutput() ApiDefinitionInfoOutput {
	return i.ToApiDefinitionInfoOutputWithContext(context.Background())
}

func (i ApiDefinitionInfoArgs) ToApiDefinitionInfoOutputWithContext(ctx context.Context) ApiDefinitionInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDefinitionInfoOutput)
}

func (i ApiDefinitionInfoArgs) ToApiDefinitionInfoPtrOutput() ApiDefinitionInfoPtrOutput {
	return i.ToApiDefinitionInfoPtrOutputWithContext(context.Background())
}

func (i ApiDefinitionInfoArgs) ToApiDefinitionInfoPtrOutputWithContext(ctx context.Context) ApiDefinitionInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDefinitionInfoOutput).ToApiDefinitionInfoPtrOutputWithContext(ctx)
}









type ApiDefinitionInfoPtrInput interface {
	pulumi.Input

	ToApiDefinitionInfoPtrOutput() ApiDefinitionInfoPtrOutput
	ToApiDefinitionInfoPtrOutputWithContext(context.Context) ApiDefinitionInfoPtrOutput
}

type apiDefinitionInfoPtrType ApiDefinitionInfoArgs

func ApiDefinitionInfoPtr(v *ApiDefinitionInfoArgs) ApiDefinitionInfoPtrInput {
	return (*apiDefinitionInfoPtrType)(v)
}

func (*apiDefinitionInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDefinitionInfo)(nil)).Elem()
}

func (i *apiDefinitionInfoPtrType) ToApiDefinitionInfoPtrOutput() ApiDefinitionInfoPtrOutput {
	return i.ToApiDefinitionInfoPtrOutputWithContext(context.Background())
}

func (i *apiDefinitionInfoPtrType) ToApiDefinitionInfoPtrOutputWithContext(ctx context.Context) ApiDefinitionInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDefinitionInfoPtrOutput)
}

type ApiDefinitionInfoOutput struct{ *pulumi.OutputState }

func (ApiDefinitionInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiDefinitionInfo)(nil)).Elem()
}

func (o ApiDefinitionInfoOutput) ToApiDefinitionInfoOutput() ApiDefinitionInfoOutput {
	return o
}

func (o ApiDefinitionInfoOutput) ToApiDefinitionInfoOutputWithContext(ctx context.Context) ApiDefinitionInfoOutput {
	return o
}

func (o ApiDefinitionInfoOutput) ToApiDefinitionInfoPtrOutput() ApiDefinitionInfoPtrOutput {
	return o.ToApiDefinitionInfoPtrOutputWithContext(context.Background())
}

func (o ApiDefinitionInfoOutput) ToApiDefinitionInfoPtrOutputWithContext(ctx context.Context) ApiDefinitionInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiDefinitionInfo) *ApiDefinitionInfo {
		return &v
	}).(ApiDefinitionInfoPtrOutput)
}

func (o ApiDefinitionInfoOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiDefinitionInfo) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type ApiDefinitionInfoPtrOutput struct{ *pulumi.OutputState }

func (ApiDefinitionInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDefinitionInfo)(nil)).Elem()
}

func (o ApiDefinitionInfoPtrOutput) ToApiDefinitionInfoPtrOutput() ApiDefinitionInfoPtrOutput {
	return o
}

func (o ApiDefinitionInfoPtrOutput) ToApiDefinitionInfoPtrOutputWithContext(ctx context.Context) ApiDefinitionInfoPtrOutput {
	return o
}

func (o ApiDefinitionInfoPtrOutput) Elem() ApiDefinitionInfoOutput {
	return o.ApplyT(func(v *ApiDefinitionInfo) ApiDefinitionInfo {
		if v != nil {
			return *v
		}
		var ret ApiDefinitionInfo
		return ret
	}).(ApiDefinitionInfoOutput)
}

func (o ApiDefinitionInfoPtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiDefinitionInfo) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type ApiDefinitionInfoResponse struct {
	Url *string `pulumi:"url"`
}

type ApiDefinitionInfoResponseOutput struct{ *pulumi.OutputState }

func (ApiDefinitionInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiDefinitionInfoResponse)(nil)).Elem()
}

func (o ApiDefinitionInfoResponseOutput) ToApiDefinitionInfoResponseOutput() ApiDefinitionInfoResponseOutput {
	return o
}

func (o ApiDefinitionInfoResponseOutput) ToApiDefinitionInfoResponseOutputWithContext(ctx context.Context) ApiDefinitionInfoResponseOutput {
	return o
}

func (o ApiDefinitionInfoResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiDefinitionInfoResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type ApiDefinitionInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiDefinitionInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDefinitionInfoResponse)(nil)).Elem()
}

func (o ApiDefinitionInfoResponsePtrOutput) ToApiDefinitionInfoResponsePtrOutput() ApiDefinitionInfoResponsePtrOutput {
	return o
}

func (o ApiDefinitionInfoResponsePtrOutput) ToApiDefinitionInfoResponsePtrOutputWithContext(ctx context.Context) ApiDefinitionInfoResponsePtrOutput {
	return o
}

func (o ApiDefinitionInfoResponsePtrOutput) Elem() ApiDefinitionInfoResponseOutput {
	return o.ApplyT(func(v *ApiDefinitionInfoResponse) ApiDefinitionInfoResponse {
		if v != nil {
			return *v
		}
		var ret ApiDefinitionInfoResponse
		return ret
	}).(ApiDefinitionInfoResponseOutput)
}

func (o ApiDefinitionInfoResponsePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiDefinitionInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type ApiManagementConfig struct {
	Id *string `pulumi:"id"`
}





type ApiManagementConfigInput interface {
	pulumi.Input

	ToApiManagementConfigOutput() ApiManagementConfigOutput
	ToApiManagementConfigOutputWithContext(context.Context) ApiManagementConfigOutput
}

type ApiManagementConfigArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ApiManagementConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiManagementConfig)(nil)).Elem()
}

func (i ApiManagementConfigArgs) ToApiManagementConfigOutput() ApiManagementConfigOutput {
	return i.ToApiManagementConfigOutputWithContext(context.Background())
}

func (i ApiManagementConfigArgs) ToApiManagementConfigOutputWithContext(ctx context.Context) ApiManagementConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementConfigOutput)
}

func (i ApiManagementConfigArgs) ToApiManagementConfigPtrOutput() ApiManagementConfigPtrOutput {
	return i.ToApiManagementConfigPtrOutputWithContext(context.Background())
}

func (i ApiManagementConfigArgs) ToApiManagementConfigPtrOutputWithContext(ctx context.Context) ApiManagementConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementConfigOutput).ToApiManagementConfigPtrOutputWithContext(ctx)
}









type ApiManagementConfigPtrInput interface {
	pulumi.Input

	ToApiManagementConfigPtrOutput() ApiManagementConfigPtrOutput
	ToApiManagementConfigPtrOutputWithContext(context.Context) ApiManagementConfigPtrOutput
}

type apiManagementConfigPtrType ApiManagementConfigArgs

func ApiManagementConfigPtr(v *ApiManagementConfigArgs) ApiManagementConfigPtrInput {
	return (*apiManagementConfigPtrType)(v)
}

func (*apiManagementConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiManagementConfig)(nil)).Elem()
}

func (i *apiManagementConfigPtrType) ToApiManagementConfigPtrOutput() ApiManagementConfigPtrOutput {
	return i.ToApiManagementConfigPtrOutputWithContext(context.Background())
}

func (i *apiManagementConfigPtrType) ToApiManagementConfigPtrOutputWithContext(ctx context.Context) ApiManagementConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementConfigPtrOutput)
}

type ApiManagementConfigOutput struct{ *pulumi.OutputState }

func (ApiManagementConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiManagementConfig)(nil)).Elem()
}

func (o ApiManagementConfigOutput) ToApiManagementConfigOutput() ApiManagementConfigOutput {
	return o
}

func (o ApiManagementConfigOutput) ToApiManagementConfigOutputWithContext(ctx context.Context) ApiManagementConfigOutput {
	return o
}

func (o ApiManagementConfigOutput) ToApiManagementConfigPtrOutput() ApiManagementConfigPtrOutput {
	return o.ToApiManagementConfigPtrOutputWithContext(context.Background())
}

func (o ApiManagementConfigOutput) ToApiManagementConfigPtrOutputWithContext(ctx context.Context) ApiManagementConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiManagementConfig) *ApiManagementConfig {
		return &v
	}).(ApiManagementConfigPtrOutput)
}

func (o ApiManagementConfigOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiManagementConfig) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ApiManagementConfigPtrOutput struct{ *pulumi.OutputState }

func (ApiManagementConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiManagementConfig)(nil)).Elem()
}

func (o ApiManagementConfigPtrOutput) ToApiManagementConfigPtrOutput() ApiManagementConfigPtrOutput {
	return o
}

func (o ApiManagementConfigPtrOutput) ToApiManagementConfigPtrOutputWithContext(ctx context.Context) ApiManagementConfigPtrOutput {
	return o
}

func (o ApiManagementConfigPtrOutput) Elem() ApiManagementConfigOutput {
	return o.ApplyT(func(v *ApiManagementConfig) ApiManagementConfig {
		if v != nil {
			return *v
		}
		var ret ApiManagementConfig
		return ret
	}).(ApiManagementConfigOutput)
}

func (o ApiManagementConfigPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiManagementConfig) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ApiManagementConfigResponse struct {
	Id *string `pulumi:"id"`
}

type ApiManagementConfigResponseOutput struct{ *pulumi.OutputState }

func (ApiManagementConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiManagementConfigResponse)(nil)).Elem()
}

func (o ApiManagementConfigResponseOutput) ToApiManagementConfigResponseOutput() ApiManagementConfigResponseOutput {
	return o
}

func (o ApiManagementConfigResponseOutput) ToApiManagementConfigResponseOutputWithContext(ctx context.Context) ApiManagementConfigResponseOutput {
	return o
}

func (o ApiManagementConfigResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiManagementConfigResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ApiManagementConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiManagementConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiManagementConfigResponse)(nil)).Elem()
}

func (o ApiManagementConfigResponsePtrOutput) ToApiManagementConfigResponsePtrOutput() ApiManagementConfigResponsePtrOutput {
	return o
}

func (o ApiManagementConfigResponsePtrOutput) ToApiManagementConfigResponsePtrOutputWithContext(ctx context.Context) ApiManagementConfigResponsePtrOutput {
	return o
}

func (o ApiManagementConfigResponsePtrOutput) Elem() ApiManagementConfigResponseOutput {
	return o.ApplyT(func(v *ApiManagementConfigResponse) ApiManagementConfigResponse {
		if v != nil {
			return *v
		}
		var ret ApiManagementConfigResponse
		return ret
	}).(ApiManagementConfigResponseOutput)
}

func (o ApiManagementConfigResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiManagementConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type AppLogsConfiguration struct {
	Destination               *string                    `pulumi:"destination"`
	LogAnalyticsConfiguration *LogAnalyticsConfiguration `pulumi:"logAnalyticsConfiguration"`
}





type AppLogsConfigurationInput interface {
	pulumi.Input

	ToAppLogsConfigurationOutput() AppLogsConfigurationOutput
	ToAppLogsConfigurationOutputWithContext(context.Context) AppLogsConfigurationOutput
}

type AppLogsConfigurationArgs struct {
	Destination               pulumi.StringPtrInput             `pulumi:"destination"`
	LogAnalyticsConfiguration LogAnalyticsConfigurationPtrInput `pulumi:"logAnalyticsConfiguration"`
}

func (AppLogsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppLogsConfiguration)(nil)).Elem()
}

func (i AppLogsConfigurationArgs) ToAppLogsConfigurationOutput() AppLogsConfigurationOutput {
	return i.ToAppLogsConfigurationOutputWithContext(context.Background())
}

func (i AppLogsConfigurationArgs) ToAppLogsConfigurationOutputWithContext(ctx context.Context) AppLogsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppLogsConfigurationOutput)
}

func (i AppLogsConfigurationArgs) ToAppLogsConfigurationPtrOutput() AppLogsConfigurationPtrOutput {
	return i.ToAppLogsConfigurationPtrOutputWithContext(context.Background())
}

func (i AppLogsConfigurationArgs) ToAppLogsConfigurationPtrOutputWithContext(ctx context.Context) AppLogsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppLogsConfigurationOutput).ToAppLogsConfigurationPtrOutputWithContext(ctx)
}









type AppLogsConfigurationPtrInput interface {
	pulumi.Input

	ToAppLogsConfigurationPtrOutput() AppLogsConfigurationPtrOutput
	ToAppLogsConfigurationPtrOutputWithContext(context.Context) AppLogsConfigurationPtrOutput
}

type appLogsConfigurationPtrType AppLogsConfigurationArgs

func AppLogsConfigurationPtr(v *AppLogsConfigurationArgs) AppLogsConfigurationPtrInput {
	return (*appLogsConfigurationPtrType)(v)
}

func (*appLogsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppLogsConfiguration)(nil)).Elem()
}

func (i *appLogsConfigurationPtrType) ToAppLogsConfigurationPtrOutput() AppLogsConfigurationPtrOutput {
	return i.ToAppLogsConfigurationPtrOutputWithContext(context.Background())
}

func (i *appLogsConfigurationPtrType) ToAppLogsConfigurationPtrOutputWithContext(ctx context.Context) AppLogsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppLogsConfigurationPtrOutput)
}

type AppLogsConfigurationOutput struct{ *pulumi.OutputState }

func (AppLogsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppLogsConfiguration)(nil)).Elem()
}

func (o AppLogsConfigurationOutput) ToAppLogsConfigurationOutput() AppLogsConfigurationOutput {
	return o
}

func (o AppLogsConfigurationOutput) ToAppLogsConfigurationOutputWithContext(ctx context.Context) AppLogsConfigurationOutput {
	return o
}

func (o AppLogsConfigurationOutput) ToAppLogsConfigurationPtrOutput() AppLogsConfigurationPtrOutput {
	return o.ToAppLogsConfigurationPtrOutputWithContext(context.Background())
}

func (o AppLogsConfigurationOutput) ToAppLogsConfigurationPtrOutputWithContext(ctx context.Context) AppLogsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppLogsConfiguration) *AppLogsConfiguration {
		return &v
	}).(AppLogsConfigurationPtrOutput)
}

func (o AppLogsConfigurationOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppLogsConfiguration) *string { return v.Destination }).(pulumi.StringPtrOutput)
}

func (o AppLogsConfigurationOutput) LogAnalyticsConfiguration() LogAnalyticsConfigurationPtrOutput {
	return o.ApplyT(func(v AppLogsConfiguration) *LogAnalyticsConfiguration { return v.LogAnalyticsConfiguration }).(LogAnalyticsConfigurationPtrOutput)
}

type AppLogsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (AppLogsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppLogsConfiguration)(nil)).Elem()
}

func (o AppLogsConfigurationPtrOutput) ToAppLogsConfigurationPtrOutput() AppLogsConfigurationPtrOutput {
	return o
}

func (o AppLogsConfigurationPtrOutput) ToAppLogsConfigurationPtrOutputWithContext(ctx context.Context) AppLogsConfigurationPtrOutput {
	return o
}

func (o AppLogsConfigurationPtrOutput) Elem() AppLogsConfigurationOutput {
	return o.ApplyT(func(v *AppLogsConfiguration) AppLogsConfiguration {
		if v != nil {
			return *v
		}
		var ret AppLogsConfiguration
		return ret
	}).(AppLogsConfigurationOutput)
}

func (o AppLogsConfigurationPtrOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppLogsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Destination
	}).(pulumi.StringPtrOutput)
}

func (o AppLogsConfigurationPtrOutput) LogAnalyticsConfiguration() LogAnalyticsConfigurationPtrOutput {
	return o.ApplyT(func(v *AppLogsConfiguration) *LogAnalyticsConfiguration {
		if v == nil {
			return nil
		}
		return v.LogAnalyticsConfiguration
	}).(LogAnalyticsConfigurationPtrOutput)
}

type AppLogsConfigurationResponse struct {
	Destination               *string                            `pulumi:"destination"`
	LogAnalyticsConfiguration *LogAnalyticsConfigurationResponse `pulumi:"logAnalyticsConfiguration"`
}

type AppLogsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (AppLogsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppLogsConfigurationResponse)(nil)).Elem()
}

func (o AppLogsConfigurationResponseOutput) ToAppLogsConfigurationResponseOutput() AppLogsConfigurationResponseOutput {
	return o
}

func (o AppLogsConfigurationResponseOutput) ToAppLogsConfigurationResponseOutputWithContext(ctx context.Context) AppLogsConfigurationResponseOutput {
	return o
}

func (o AppLogsConfigurationResponseOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppLogsConfigurationResponse) *string { return v.Destination }).(pulumi.StringPtrOutput)
}

func (o AppLogsConfigurationResponseOutput) LogAnalyticsConfiguration() LogAnalyticsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v AppLogsConfigurationResponse) *LogAnalyticsConfigurationResponse {
		return v.LogAnalyticsConfiguration
	}).(LogAnalyticsConfigurationResponsePtrOutput)
}

type AppLogsConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (AppLogsConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppLogsConfigurationResponse)(nil)).Elem()
}

func (o AppLogsConfigurationResponsePtrOutput) ToAppLogsConfigurationResponsePtrOutput() AppLogsConfigurationResponsePtrOutput {
	return o
}

func (o AppLogsConfigurationResponsePtrOutput) ToAppLogsConfigurationResponsePtrOutputWithContext(ctx context.Context) AppLogsConfigurationResponsePtrOutput {
	return o
}

func (o AppLogsConfigurationResponsePtrOutput) Elem() AppLogsConfigurationResponseOutput {
	return o.ApplyT(func(v *AppLogsConfigurationResponse) AppLogsConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret AppLogsConfigurationResponse
		return ret
	}).(AppLogsConfigurationResponseOutput)
}

func (o AppLogsConfigurationResponsePtrOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppLogsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Destination
	}).(pulumi.StringPtrOutput)
}

func (o AppLogsConfigurationResponsePtrOutput) LogAnalyticsConfiguration() LogAnalyticsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AppLogsConfigurationResponse) *LogAnalyticsConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.LogAnalyticsConfiguration
	}).(LogAnalyticsConfigurationResponsePtrOutput)
}

type AppRegistration struct {
	AppId                *string `pulumi:"appId"`
	AppSecretSettingName *string `pulumi:"appSecretSettingName"`
}





type AppRegistrationInput interface {
	pulumi.Input

	ToAppRegistrationOutput() AppRegistrationOutput
	ToAppRegistrationOutputWithContext(context.Context) AppRegistrationOutput
}

type AppRegistrationArgs struct {
	AppId                pulumi.StringPtrInput `pulumi:"appId"`
	AppSecretSettingName pulumi.StringPtrInput `pulumi:"appSecretSettingName"`
}

func (AppRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppRegistration)(nil)).Elem()
}

func (i AppRegistrationArgs) ToAppRegistrationOutput() AppRegistrationOutput {
	return i.ToAppRegistrationOutputWithContext(context.Background())
}

func (i AppRegistrationArgs) ToAppRegistrationOutputWithContext(ctx context.Context) AppRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppRegistrationOutput)
}

func (i AppRegistrationArgs) ToAppRegistrationPtrOutput() AppRegistrationPtrOutput {
	return i.ToAppRegistrationPtrOutputWithContext(context.Background())
}

func (i AppRegistrationArgs) ToAppRegistrationPtrOutputWithContext(ctx context.Context) AppRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppRegistrationOutput).ToAppRegistrationPtrOutputWithContext(ctx)
}









type AppRegistrationPtrInput interface {
	pulumi.Input

	ToAppRegistrationPtrOutput() AppRegistrationPtrOutput
	ToAppRegistrationPtrOutputWithContext(context.Context) AppRegistrationPtrOutput
}

type appRegistrationPtrType AppRegistrationArgs

func AppRegistrationPtr(v *AppRegistrationArgs) AppRegistrationPtrInput {
	return (*appRegistrationPtrType)(v)
}

func (*appRegistrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppRegistration)(nil)).Elem()
}

func (i *appRegistrationPtrType) ToAppRegistrationPtrOutput() AppRegistrationPtrOutput {
	return i.ToAppRegistrationPtrOutputWithContext(context.Background())
}

func (i *appRegistrationPtrType) ToAppRegistrationPtrOutputWithContext(ctx context.Context) AppRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppRegistrationPtrOutput)
}

type AppRegistrationOutput struct{ *pulumi.OutputState }

func (AppRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppRegistration)(nil)).Elem()
}

func (o AppRegistrationOutput) ToAppRegistrationOutput() AppRegistrationOutput {
	return o
}

func (o AppRegistrationOutput) ToAppRegistrationOutputWithContext(ctx context.Context) AppRegistrationOutput {
	return o
}

func (o AppRegistrationOutput) ToAppRegistrationPtrOutput() AppRegistrationPtrOutput {
	return o.ToAppRegistrationPtrOutputWithContext(context.Background())
}

func (o AppRegistrationOutput) ToAppRegistrationPtrOutputWithContext(ctx context.Context) AppRegistrationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppRegistration) *AppRegistration {
		return &v
	}).(AppRegistrationPtrOutput)
}

func (o AppRegistrationOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppRegistration) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o AppRegistrationOutput) AppSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppRegistration) *string { return v.AppSecretSettingName }).(pulumi.StringPtrOutput)
}

type AppRegistrationPtrOutput struct{ *pulumi.OutputState }

func (AppRegistrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppRegistration)(nil)).Elem()
}

func (o AppRegistrationPtrOutput) ToAppRegistrationPtrOutput() AppRegistrationPtrOutput {
	return o
}

func (o AppRegistrationPtrOutput) ToAppRegistrationPtrOutputWithContext(ctx context.Context) AppRegistrationPtrOutput {
	return o
}

func (o AppRegistrationPtrOutput) Elem() AppRegistrationOutput {
	return o.ApplyT(func(v *AppRegistration) AppRegistration {
		if v != nil {
			return *v
		}
		var ret AppRegistration
		return ret
	}).(AppRegistrationOutput)
}

func (o AppRegistrationPtrOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppRegistration) *string {
		if v == nil {
			return nil
		}
		return v.AppId
	}).(pulumi.StringPtrOutput)
}

func (o AppRegistrationPtrOutput) AppSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppRegistration) *string {
		if v == nil {
			return nil
		}
		return v.AppSecretSettingName
	}).(pulumi.StringPtrOutput)
}

type AppRegistrationResponse struct {
	AppId                *string `pulumi:"appId"`
	AppSecretSettingName *string `pulumi:"appSecretSettingName"`
}

type AppRegistrationResponseOutput struct{ *pulumi.OutputState }

func (AppRegistrationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppRegistrationResponse)(nil)).Elem()
}

func (o AppRegistrationResponseOutput) ToAppRegistrationResponseOutput() AppRegistrationResponseOutput {
	return o
}

func (o AppRegistrationResponseOutput) ToAppRegistrationResponseOutputWithContext(ctx context.Context) AppRegistrationResponseOutput {
	return o
}

func (o AppRegistrationResponseOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppRegistrationResponse) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o AppRegistrationResponseOutput) AppSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppRegistrationResponse) *string { return v.AppSecretSettingName }).(pulumi.StringPtrOutput)
}

type AppRegistrationResponsePtrOutput struct{ *pulumi.OutputState }

func (AppRegistrationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppRegistrationResponse)(nil)).Elem()
}

func (o AppRegistrationResponsePtrOutput) ToAppRegistrationResponsePtrOutput() AppRegistrationResponsePtrOutput {
	return o
}

func (o AppRegistrationResponsePtrOutput) ToAppRegistrationResponsePtrOutputWithContext(ctx context.Context) AppRegistrationResponsePtrOutput {
	return o
}

func (o AppRegistrationResponsePtrOutput) Elem() AppRegistrationResponseOutput {
	return o.ApplyT(func(v *AppRegistrationResponse) AppRegistrationResponse {
		if v != nil {
			return *v
		}
		var ret AppRegistrationResponse
		return ret
	}).(AppRegistrationResponseOutput)
}

func (o AppRegistrationResponsePtrOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppId
	}).(pulumi.StringPtrOutput)
}

func (o AppRegistrationResponsePtrOutput) AppSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppSecretSettingName
	}).(pulumi.StringPtrOutput)
}

type Apple struct {
	Enabled      *bool              `pulumi:"enabled"`
	Login        *LoginScopes       `pulumi:"login"`
	Registration *AppleRegistration `pulumi:"registration"`
}





type AppleInput interface {
	pulumi.Input

	ToAppleOutput() AppleOutput
	ToAppleOutputWithContext(context.Context) AppleOutput
}

type AppleArgs struct {
	Enabled      pulumi.BoolPtrInput       `pulumi:"enabled"`
	Login        LoginScopesPtrInput       `pulumi:"login"`
	Registration AppleRegistrationPtrInput `pulumi:"registration"`
}

func (AppleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Apple)(nil)).Elem()
}

func (i AppleArgs) ToAppleOutput() AppleOutput {
	return i.ToAppleOutputWithContext(context.Background())
}

func (i AppleArgs) ToAppleOutputWithContext(ctx context.Context) AppleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppleOutput)
}

func (i AppleArgs) ToApplePtrOutput() ApplePtrOutput {
	return i.ToApplePtrOutputWithContext(context.Background())
}

func (i AppleArgs) ToApplePtrOutputWithContext(ctx context.Context) ApplePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppleOutput).ToApplePtrOutputWithContext(ctx)
}









type ApplePtrInput interface {
	pulumi.Input

	ToApplePtrOutput() ApplePtrOutput
	ToApplePtrOutputWithContext(context.Context) ApplePtrOutput
}

type applePtrType AppleArgs

func ApplePtr(v *AppleArgs) ApplePtrInput {
	return (*applePtrType)(v)
}

func (*applePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Apple)(nil)).Elem()
}

func (i *applePtrType) ToApplePtrOutput() ApplePtrOutput {
	return i.ToApplePtrOutputWithContext(context.Background())
}

func (i *applePtrType) ToApplePtrOutputWithContext(ctx context.Context) ApplePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplePtrOutput)
}

type AppleOutput struct{ *pulumi.OutputState }

func (AppleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Apple)(nil)).Elem()
}

func (o AppleOutput) ToAppleOutput() AppleOutput {
	return o
}

func (o AppleOutput) ToAppleOutputWithContext(ctx context.Context) AppleOutput {
	return o
}

func (o AppleOutput) ToApplePtrOutput() ApplePtrOutput {
	return o.ToApplePtrOutputWithContext(context.Background())
}

func (o AppleOutput) ToApplePtrOutputWithContext(ctx context.Context) ApplePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Apple) *Apple {
		return &v
	}).(ApplePtrOutput)
}

func (o AppleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Apple) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AppleOutput) Login() LoginScopesPtrOutput {
	return o.ApplyT(func(v Apple) *LoginScopes { return v.Login }).(LoginScopesPtrOutput)
}

func (o AppleOutput) Registration() AppleRegistrationPtrOutput {
	return o.ApplyT(func(v Apple) *AppleRegistration { return v.Registration }).(AppleRegistrationPtrOutput)
}

type ApplePtrOutput struct{ *pulumi.OutputState }

func (ApplePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Apple)(nil)).Elem()
}

func (o ApplePtrOutput) ToApplePtrOutput() ApplePtrOutput {
	return o
}

func (o ApplePtrOutput) ToApplePtrOutputWithContext(ctx context.Context) ApplePtrOutput {
	return o
}

func (o ApplePtrOutput) Elem() AppleOutput {
	return o.ApplyT(func(v *Apple) Apple {
		if v != nil {
			return *v
		}
		var ret Apple
		return ret
	}).(AppleOutput)
}

func (o ApplePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Apple) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o ApplePtrOutput) Login() LoginScopesPtrOutput {
	return o.ApplyT(func(v *Apple) *LoginScopes {
		if v == nil {
			return nil
		}
		return v.Login
	}).(LoginScopesPtrOutput)
}

func (o ApplePtrOutput) Registration() AppleRegistrationPtrOutput {
	return o.ApplyT(func(v *Apple) *AppleRegistration {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(AppleRegistrationPtrOutput)
}

type AppleRegistration struct {
	ClientId                *string `pulumi:"clientId"`
	ClientSecretSettingName *string `pulumi:"clientSecretSettingName"`
}





type AppleRegistrationInput interface {
	pulumi.Input

	ToAppleRegistrationOutput() AppleRegistrationOutput
	ToAppleRegistrationOutputWithContext(context.Context) AppleRegistrationOutput
}

type AppleRegistrationArgs struct {
	ClientId                pulumi.StringPtrInput `pulumi:"clientId"`
	ClientSecretSettingName pulumi.StringPtrInput `pulumi:"clientSecretSettingName"`
}

func (AppleRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppleRegistration)(nil)).Elem()
}

func (i AppleRegistrationArgs) ToAppleRegistrationOutput() AppleRegistrationOutput {
	return i.ToAppleRegistrationOutputWithContext(context.Background())
}

func (i AppleRegistrationArgs) ToAppleRegistrationOutputWithContext(ctx context.Context) AppleRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppleRegistrationOutput)
}

func (i AppleRegistrationArgs) ToAppleRegistrationPtrOutput() AppleRegistrationPtrOutput {
	return i.ToAppleRegistrationPtrOutputWithContext(context.Background())
}

func (i AppleRegistrationArgs) ToAppleRegistrationPtrOutputWithContext(ctx context.Context) AppleRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppleRegistrationOutput).ToAppleRegistrationPtrOutputWithContext(ctx)
}









type AppleRegistrationPtrInput interface {
	pulumi.Input

	ToAppleRegistrationPtrOutput() AppleRegistrationPtrOutput
	ToAppleRegistrationPtrOutputWithContext(context.Context) AppleRegistrationPtrOutput
}

type appleRegistrationPtrType AppleRegistrationArgs

func AppleRegistrationPtr(v *AppleRegistrationArgs) AppleRegistrationPtrInput {
	return (*appleRegistrationPtrType)(v)
}

func (*appleRegistrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppleRegistration)(nil)).Elem()
}

func (i *appleRegistrationPtrType) ToAppleRegistrationPtrOutput() AppleRegistrationPtrOutput {
	return i.ToAppleRegistrationPtrOutputWithContext(context.Background())
}

func (i *appleRegistrationPtrType) ToAppleRegistrationPtrOutputWithContext(ctx context.Context) AppleRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppleRegistrationPtrOutput)
}

type AppleRegistrationOutput struct{ *pulumi.OutputState }

func (AppleRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppleRegistration)(nil)).Elem()
}

func (o AppleRegistrationOutput) ToAppleRegistrationOutput() AppleRegistrationOutput {
	return o
}

func (o AppleRegistrationOutput) ToAppleRegistrationOutputWithContext(ctx context.Context) AppleRegistrationOutput {
	return o
}

func (o AppleRegistrationOutput) ToAppleRegistrationPtrOutput() AppleRegistrationPtrOutput {
	return o.ToAppleRegistrationPtrOutputWithContext(context.Background())
}

func (o AppleRegistrationOutput) ToAppleRegistrationPtrOutputWithContext(ctx context.Context) AppleRegistrationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppleRegistration) *AppleRegistration {
		return &v
	}).(AppleRegistrationPtrOutput)
}

func (o AppleRegistrationOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppleRegistration) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o AppleRegistrationOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppleRegistration) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

type AppleRegistrationPtrOutput struct{ *pulumi.OutputState }

func (AppleRegistrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppleRegistration)(nil)).Elem()
}

func (o AppleRegistrationPtrOutput) ToAppleRegistrationPtrOutput() AppleRegistrationPtrOutput {
	return o
}

func (o AppleRegistrationPtrOutput) ToAppleRegistrationPtrOutputWithContext(ctx context.Context) AppleRegistrationPtrOutput {
	return o
}

func (o AppleRegistrationPtrOutput) Elem() AppleRegistrationOutput {
	return o.ApplyT(func(v *AppleRegistration) AppleRegistration {
		if v != nil {
			return *v
		}
		var ret AppleRegistration
		return ret
	}).(AppleRegistrationOutput)
}

func (o AppleRegistrationPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppleRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AppleRegistrationPtrOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppleRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretSettingName
	}).(pulumi.StringPtrOutput)
}

type AppleRegistrationResponse struct {
	ClientId                *string `pulumi:"clientId"`
	ClientSecretSettingName *string `pulumi:"clientSecretSettingName"`
}

type AppleRegistrationResponseOutput struct{ *pulumi.OutputState }

func (AppleRegistrationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppleRegistrationResponse)(nil)).Elem()
}

func (o AppleRegistrationResponseOutput) ToAppleRegistrationResponseOutput() AppleRegistrationResponseOutput {
	return o
}

func (o AppleRegistrationResponseOutput) ToAppleRegistrationResponseOutputWithContext(ctx context.Context) AppleRegistrationResponseOutput {
	return o
}

func (o AppleRegistrationResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppleRegistrationResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o AppleRegistrationResponseOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppleRegistrationResponse) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

type AppleRegistrationResponsePtrOutput struct{ *pulumi.OutputState }

func (AppleRegistrationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppleRegistrationResponse)(nil)).Elem()
}

func (o AppleRegistrationResponsePtrOutput) ToAppleRegistrationResponsePtrOutput() AppleRegistrationResponsePtrOutput {
	return o
}

func (o AppleRegistrationResponsePtrOutput) ToAppleRegistrationResponsePtrOutputWithContext(ctx context.Context) AppleRegistrationResponsePtrOutput {
	return o
}

func (o AppleRegistrationResponsePtrOutput) Elem() AppleRegistrationResponseOutput {
	return o.ApplyT(func(v *AppleRegistrationResponse) AppleRegistrationResponse {
		if v != nil {
			return *v
		}
		var ret AppleRegistrationResponse
		return ret
	}).(AppleRegistrationResponseOutput)
}

func (o AppleRegistrationResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppleRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AppleRegistrationResponsePtrOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppleRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretSettingName
	}).(pulumi.StringPtrOutput)
}

type AppleResponse struct {
	Enabled      *bool                      `pulumi:"enabled"`
	Login        *LoginScopesResponse       `pulumi:"login"`
	Registration *AppleRegistrationResponse `pulumi:"registration"`
}

type AppleResponseOutput struct{ *pulumi.OutputState }

func (AppleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppleResponse)(nil)).Elem()
}

func (o AppleResponseOutput) ToAppleResponseOutput() AppleResponseOutput {
	return o
}

func (o AppleResponseOutput) ToAppleResponseOutputWithContext(ctx context.Context) AppleResponseOutput {
	return o
}

func (o AppleResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AppleResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AppleResponseOutput) Login() LoginScopesResponsePtrOutput {
	return o.ApplyT(func(v AppleResponse) *LoginScopesResponse { return v.Login }).(LoginScopesResponsePtrOutput)
}

func (o AppleResponseOutput) Registration() AppleRegistrationResponsePtrOutput {
	return o.ApplyT(func(v AppleResponse) *AppleRegistrationResponse { return v.Registration }).(AppleRegistrationResponsePtrOutput)
}

type AppleResponsePtrOutput struct{ *pulumi.OutputState }

func (AppleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppleResponse)(nil)).Elem()
}

func (o AppleResponsePtrOutput) ToAppleResponsePtrOutput() AppleResponsePtrOutput {
	return o
}

func (o AppleResponsePtrOutput) ToAppleResponsePtrOutputWithContext(ctx context.Context) AppleResponsePtrOutput {
	return o
}

func (o AppleResponsePtrOutput) Elem() AppleResponseOutput {
	return o.ApplyT(func(v *AppleResponse) AppleResponse {
		if v != nil {
			return *v
		}
		var ret AppleResponse
		return ret
	}).(AppleResponseOutput)
}

func (o AppleResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppleResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AppleResponsePtrOutput) Login() LoginScopesResponsePtrOutput {
	return o.ApplyT(func(v *AppleResponse) *LoginScopesResponse {
		if v == nil {
			return nil
		}
		return v.Login
	}).(LoginScopesResponsePtrOutput)
}

func (o AppleResponsePtrOutput) Registration() AppleRegistrationResponsePtrOutput {
	return o.ApplyT(func(v *AppleResponse) *AppleRegistrationResponse {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(AppleRegistrationResponsePtrOutput)
}

type ApplicationLogsConfig struct {
	AzureBlobStorage  *AzureBlobStorageApplicationLogsConfig  `pulumi:"azureBlobStorage"`
	AzureTableStorage *AzureTableStorageApplicationLogsConfig `pulumi:"azureTableStorage"`
	FileSystem        *FileSystemApplicationLogsConfig        `pulumi:"fileSystem"`
}


func (val *ApplicationLogsConfig) Defaults() *ApplicationLogsConfig {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.FileSystem = tmp.FileSystem.Defaults()

	return &tmp
}





type ApplicationLogsConfigInput interface {
	pulumi.Input

	ToApplicationLogsConfigOutput() ApplicationLogsConfigOutput
	ToApplicationLogsConfigOutputWithContext(context.Context) ApplicationLogsConfigOutput
}

type ApplicationLogsConfigArgs struct {
	AzureBlobStorage  AzureBlobStorageApplicationLogsConfigPtrInput  `pulumi:"azureBlobStorage"`
	AzureTableStorage AzureTableStorageApplicationLogsConfigPtrInput `pulumi:"azureTableStorage"`
	FileSystem        FileSystemApplicationLogsConfigPtrInput        `pulumi:"fileSystem"`
}


func (val *ApplicationLogsConfigArgs) Defaults() *ApplicationLogsConfigArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	return &tmp
}
func (ApplicationLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationLogsConfig)(nil)).Elem()
}

func (i ApplicationLogsConfigArgs) ToApplicationLogsConfigOutput() ApplicationLogsConfigOutput {
	return i.ToApplicationLogsConfigOutputWithContext(context.Background())
}

func (i ApplicationLogsConfigArgs) ToApplicationLogsConfigOutputWithContext(ctx context.Context) ApplicationLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLogsConfigOutput)
}

func (i ApplicationLogsConfigArgs) ToApplicationLogsConfigPtrOutput() ApplicationLogsConfigPtrOutput {
	return i.ToApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i ApplicationLogsConfigArgs) ToApplicationLogsConfigPtrOutputWithContext(ctx context.Context) ApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLogsConfigOutput).ToApplicationLogsConfigPtrOutputWithContext(ctx)
}









type ApplicationLogsConfigPtrInput interface {
	pulumi.Input

	ToApplicationLogsConfigPtrOutput() ApplicationLogsConfigPtrOutput
	ToApplicationLogsConfigPtrOutputWithContext(context.Context) ApplicationLogsConfigPtrOutput
}

type applicationLogsConfigPtrType ApplicationLogsConfigArgs

func ApplicationLogsConfigPtr(v *ApplicationLogsConfigArgs) ApplicationLogsConfigPtrInput {
	return (*applicationLogsConfigPtrType)(v)
}

func (*applicationLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLogsConfig)(nil)).Elem()
}

func (i *applicationLogsConfigPtrType) ToApplicationLogsConfigPtrOutput() ApplicationLogsConfigPtrOutput {
	return i.ToApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i *applicationLogsConfigPtrType) ToApplicationLogsConfigPtrOutputWithContext(ctx context.Context) ApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLogsConfigPtrOutput)
}

type ApplicationLogsConfigOutput struct{ *pulumi.OutputState }

func (ApplicationLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationLogsConfig)(nil)).Elem()
}

func (o ApplicationLogsConfigOutput) ToApplicationLogsConfigOutput() ApplicationLogsConfigOutput {
	return o
}

func (o ApplicationLogsConfigOutput) ToApplicationLogsConfigOutputWithContext(ctx context.Context) ApplicationLogsConfigOutput {
	return o
}

func (o ApplicationLogsConfigOutput) ToApplicationLogsConfigPtrOutput() ApplicationLogsConfigPtrOutput {
	return o.ToApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (o ApplicationLogsConfigOutput) ToApplicationLogsConfigPtrOutputWithContext(ctx context.Context) ApplicationLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationLogsConfig) *ApplicationLogsConfig {
		return &v
	}).(ApplicationLogsConfigPtrOutput)
}

func (o ApplicationLogsConfigOutput) AzureBlobStorage() AzureBlobStorageApplicationLogsConfigPtrOutput {
	return o.ApplyT(func(v ApplicationLogsConfig) *AzureBlobStorageApplicationLogsConfig { return v.AzureBlobStorage }).(AzureBlobStorageApplicationLogsConfigPtrOutput)
}

func (o ApplicationLogsConfigOutput) AzureTableStorage() AzureTableStorageApplicationLogsConfigPtrOutput {
	return o.ApplyT(func(v ApplicationLogsConfig) *AzureTableStorageApplicationLogsConfig { return v.AzureTableStorage }).(AzureTableStorageApplicationLogsConfigPtrOutput)
}

func (o ApplicationLogsConfigOutput) FileSystem() FileSystemApplicationLogsConfigPtrOutput {
	return o.ApplyT(func(v ApplicationLogsConfig) *FileSystemApplicationLogsConfig { return v.FileSystem }).(FileSystemApplicationLogsConfigPtrOutput)
}

type ApplicationLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (ApplicationLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLogsConfig)(nil)).Elem()
}

func (o ApplicationLogsConfigPtrOutput) ToApplicationLogsConfigPtrOutput() ApplicationLogsConfigPtrOutput {
	return o
}

func (o ApplicationLogsConfigPtrOutput) ToApplicationLogsConfigPtrOutputWithContext(ctx context.Context) ApplicationLogsConfigPtrOutput {
	return o
}

func (o ApplicationLogsConfigPtrOutput) Elem() ApplicationLogsConfigOutput {
	return o.ApplyT(func(v *ApplicationLogsConfig) ApplicationLogsConfig {
		if v != nil {
			return *v
		}
		var ret ApplicationLogsConfig
		return ret
	}).(ApplicationLogsConfigOutput)
}

func (o ApplicationLogsConfigPtrOutput) AzureBlobStorage() AzureBlobStorageApplicationLogsConfigPtrOutput {
	return o.ApplyT(func(v *ApplicationLogsConfig) *AzureBlobStorageApplicationLogsConfig {
		if v == nil {
			return nil
		}
		return v.AzureBlobStorage
	}).(AzureBlobStorageApplicationLogsConfigPtrOutput)
}

func (o ApplicationLogsConfigPtrOutput) AzureTableStorage() AzureTableStorageApplicationLogsConfigPtrOutput {
	return o.ApplyT(func(v *ApplicationLogsConfig) *AzureTableStorageApplicationLogsConfig {
		if v == nil {
			return nil
		}
		return v.AzureTableStorage
	}).(AzureTableStorageApplicationLogsConfigPtrOutput)
}

func (o ApplicationLogsConfigPtrOutput) FileSystem() FileSystemApplicationLogsConfigPtrOutput {
	return o.ApplyT(func(v *ApplicationLogsConfig) *FileSystemApplicationLogsConfig {
		if v == nil {
			return nil
		}
		return v.FileSystem
	}).(FileSystemApplicationLogsConfigPtrOutput)
}

type ApplicationLogsConfigResponse struct {
	AzureBlobStorage  *AzureBlobStorageApplicationLogsConfigResponse  `pulumi:"azureBlobStorage"`
	AzureTableStorage *AzureTableStorageApplicationLogsConfigResponse `pulumi:"azureTableStorage"`
	FileSystem        *FileSystemApplicationLogsConfigResponse        `pulumi:"fileSystem"`
}


func (val *ApplicationLogsConfigResponse) Defaults() *ApplicationLogsConfigResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.FileSystem = tmp.FileSystem.Defaults()

	return &tmp
}

type ApplicationLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (ApplicationLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationLogsConfigResponse)(nil)).Elem()
}

func (o ApplicationLogsConfigResponseOutput) ToApplicationLogsConfigResponseOutput() ApplicationLogsConfigResponseOutput {
	return o
}

func (o ApplicationLogsConfigResponseOutput) ToApplicationLogsConfigResponseOutputWithContext(ctx context.Context) ApplicationLogsConfigResponseOutput {
	return o
}

func (o ApplicationLogsConfigResponseOutput) AzureBlobStorage() AzureBlobStorageApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v ApplicationLogsConfigResponse) *AzureBlobStorageApplicationLogsConfigResponse {
		return v.AzureBlobStorage
	}).(AzureBlobStorageApplicationLogsConfigResponsePtrOutput)
}

func (o ApplicationLogsConfigResponseOutput) AzureTableStorage() AzureTableStorageApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v ApplicationLogsConfigResponse) *AzureTableStorageApplicationLogsConfigResponse {
		return v.AzureTableStorage
	}).(AzureTableStorageApplicationLogsConfigResponsePtrOutput)
}

func (o ApplicationLogsConfigResponseOutput) FileSystem() FileSystemApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v ApplicationLogsConfigResponse) *FileSystemApplicationLogsConfigResponse { return v.FileSystem }).(FileSystemApplicationLogsConfigResponsePtrOutput)
}

type ApplicationLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLogsConfigResponse)(nil)).Elem()
}

func (o ApplicationLogsConfigResponsePtrOutput) ToApplicationLogsConfigResponsePtrOutput() ApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o ApplicationLogsConfigResponsePtrOutput) ToApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) ApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o ApplicationLogsConfigResponsePtrOutput) Elem() ApplicationLogsConfigResponseOutput {
	return o.ApplyT(func(v *ApplicationLogsConfigResponse) ApplicationLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationLogsConfigResponse
		return ret
	}).(ApplicationLogsConfigResponseOutput)
}

func (o ApplicationLogsConfigResponsePtrOutput) AzureBlobStorage() AzureBlobStorageApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationLogsConfigResponse) *AzureBlobStorageApplicationLogsConfigResponse {
		if v == nil {
			return nil
		}
		return v.AzureBlobStorage
	}).(AzureBlobStorageApplicationLogsConfigResponsePtrOutput)
}

func (o ApplicationLogsConfigResponsePtrOutput) AzureTableStorage() AzureTableStorageApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationLogsConfigResponse) *AzureTableStorageApplicationLogsConfigResponse {
		if v == nil {
			return nil
		}
		return v.AzureTableStorage
	}).(AzureTableStorageApplicationLogsConfigResponsePtrOutput)
}

func (o ApplicationLogsConfigResponsePtrOutput) FileSystem() FileSystemApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationLogsConfigResponse) *FileSystemApplicationLogsConfigResponse {
		if v == nil {
			return nil
		}
		return v.FileSystem
	}).(FileSystemApplicationLogsConfigResponsePtrOutput)
}

type ArcConfiguration struct {
	ArtifactStorageAccessMode    *string                `pulumi:"artifactStorageAccessMode"`
	ArtifactStorageClassName     *string                `pulumi:"artifactStorageClassName"`
	ArtifactStorageMountPath     *string                `pulumi:"artifactStorageMountPath"`
	ArtifactStorageNodeName      *string                `pulumi:"artifactStorageNodeName"`
	ArtifactsStorageType         *StorageType           `pulumi:"artifactsStorageType"`
	FrontEndServiceConfiguration *FrontEndConfiguration `pulumi:"frontEndServiceConfiguration"`
	KubeConfig                   *string                `pulumi:"kubeConfig"`
}





type ArcConfigurationInput interface {
	pulumi.Input

	ToArcConfigurationOutput() ArcConfigurationOutput
	ToArcConfigurationOutputWithContext(context.Context) ArcConfigurationOutput
}

type ArcConfigurationArgs struct {
	ArtifactStorageAccessMode    pulumi.StringPtrInput         `pulumi:"artifactStorageAccessMode"`
	ArtifactStorageClassName     pulumi.StringPtrInput         `pulumi:"artifactStorageClassName"`
	ArtifactStorageMountPath     pulumi.StringPtrInput         `pulumi:"artifactStorageMountPath"`
	ArtifactStorageNodeName      pulumi.StringPtrInput         `pulumi:"artifactStorageNodeName"`
	ArtifactsStorageType         StorageTypePtrInput           `pulumi:"artifactsStorageType"`
	FrontEndServiceConfiguration FrontEndConfigurationPtrInput `pulumi:"frontEndServiceConfiguration"`
	KubeConfig                   pulumi.StringPtrInput         `pulumi:"kubeConfig"`
}

func (ArcConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcConfiguration)(nil)).Elem()
}

func (i ArcConfigurationArgs) ToArcConfigurationOutput() ArcConfigurationOutput {
	return i.ToArcConfigurationOutputWithContext(context.Background())
}

func (i ArcConfigurationArgs) ToArcConfigurationOutputWithContext(ctx context.Context) ArcConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArcConfigurationOutput)
}

func (i ArcConfigurationArgs) ToArcConfigurationPtrOutput() ArcConfigurationPtrOutput {
	return i.ToArcConfigurationPtrOutputWithContext(context.Background())
}

func (i ArcConfigurationArgs) ToArcConfigurationPtrOutputWithContext(ctx context.Context) ArcConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArcConfigurationOutput).ToArcConfigurationPtrOutputWithContext(ctx)
}









type ArcConfigurationPtrInput interface {
	pulumi.Input

	ToArcConfigurationPtrOutput() ArcConfigurationPtrOutput
	ToArcConfigurationPtrOutputWithContext(context.Context) ArcConfigurationPtrOutput
}

type arcConfigurationPtrType ArcConfigurationArgs

func ArcConfigurationPtr(v *ArcConfigurationArgs) ArcConfigurationPtrInput {
	return (*arcConfigurationPtrType)(v)
}

func (*arcConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcConfiguration)(nil)).Elem()
}

func (i *arcConfigurationPtrType) ToArcConfigurationPtrOutput() ArcConfigurationPtrOutput {
	return i.ToArcConfigurationPtrOutputWithContext(context.Background())
}

func (i *arcConfigurationPtrType) ToArcConfigurationPtrOutputWithContext(ctx context.Context) ArcConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArcConfigurationPtrOutput)
}

type ArcConfigurationOutput struct{ *pulumi.OutputState }

func (ArcConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcConfiguration)(nil)).Elem()
}

func (o ArcConfigurationOutput) ToArcConfigurationOutput() ArcConfigurationOutput {
	return o
}

func (o ArcConfigurationOutput) ToArcConfigurationOutputWithContext(ctx context.Context) ArcConfigurationOutput {
	return o
}

func (o ArcConfigurationOutput) ToArcConfigurationPtrOutput() ArcConfigurationPtrOutput {
	return o.ToArcConfigurationPtrOutputWithContext(context.Background())
}

func (o ArcConfigurationOutput) ToArcConfigurationPtrOutputWithContext(ctx context.Context) ArcConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArcConfiguration) *ArcConfiguration {
		return &v
	}).(ArcConfigurationPtrOutput)
}

func (o ArcConfigurationOutput) ArtifactStorageAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfiguration) *string { return v.ArtifactStorageAccessMode }).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationOutput) ArtifactStorageClassName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfiguration) *string { return v.ArtifactStorageClassName }).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationOutput) ArtifactStorageMountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfiguration) *string { return v.ArtifactStorageMountPath }).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationOutput) ArtifactStorageNodeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfiguration) *string { return v.ArtifactStorageNodeName }).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationOutput) ArtifactsStorageType() StorageTypePtrOutput {
	return o.ApplyT(func(v ArcConfiguration) *StorageType { return v.ArtifactsStorageType }).(StorageTypePtrOutput)
}

func (o ArcConfigurationOutput) FrontEndServiceConfiguration() FrontEndConfigurationPtrOutput {
	return o.ApplyT(func(v ArcConfiguration) *FrontEndConfiguration { return v.FrontEndServiceConfiguration }).(FrontEndConfigurationPtrOutput)
}

func (o ArcConfigurationOutput) KubeConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfiguration) *string { return v.KubeConfig }).(pulumi.StringPtrOutput)
}

type ArcConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ArcConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcConfiguration)(nil)).Elem()
}

func (o ArcConfigurationPtrOutput) ToArcConfigurationPtrOutput() ArcConfigurationPtrOutput {
	return o
}

func (o ArcConfigurationPtrOutput) ToArcConfigurationPtrOutputWithContext(ctx context.Context) ArcConfigurationPtrOutput {
	return o
}

func (o ArcConfigurationPtrOutput) Elem() ArcConfigurationOutput {
	return o.ApplyT(func(v *ArcConfiguration) ArcConfiguration {
		if v != nil {
			return *v
		}
		var ret ArcConfiguration
		return ret
	}).(ArcConfigurationOutput)
}

func (o ArcConfigurationPtrOutput) ArtifactStorageAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactStorageAccessMode
	}).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationPtrOutput) ArtifactStorageClassName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactStorageClassName
	}).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationPtrOutput) ArtifactStorageMountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactStorageMountPath
	}).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationPtrOutput) ArtifactStorageNodeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactStorageNodeName
	}).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationPtrOutput) ArtifactsStorageType() StorageTypePtrOutput {
	return o.ApplyT(func(v *ArcConfiguration) *StorageType {
		if v == nil {
			return nil
		}
		return v.ArtifactsStorageType
	}).(StorageTypePtrOutput)
}

func (o ArcConfigurationPtrOutput) FrontEndServiceConfiguration() FrontEndConfigurationPtrOutput {
	return o.ApplyT(func(v *ArcConfiguration) *FrontEndConfiguration {
		if v == nil {
			return nil
		}
		return v.FrontEndServiceConfiguration
	}).(FrontEndConfigurationPtrOutput)
}

func (o ArcConfigurationPtrOutput) KubeConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.KubeConfig
	}).(pulumi.StringPtrOutput)
}

type ArcConfigurationResponse struct {
	ArtifactStorageAccessMode    *string                        `pulumi:"artifactStorageAccessMode"`
	ArtifactStorageClassName     *string                        `pulumi:"artifactStorageClassName"`
	ArtifactStorageMountPath     *string                        `pulumi:"artifactStorageMountPath"`
	ArtifactStorageNodeName      *string                        `pulumi:"artifactStorageNodeName"`
	ArtifactsStorageType         *string                        `pulumi:"artifactsStorageType"`
	FrontEndServiceConfiguration *FrontEndConfigurationResponse `pulumi:"frontEndServiceConfiguration"`
}

type ArcConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ArcConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcConfigurationResponse)(nil)).Elem()
}

func (o ArcConfigurationResponseOutput) ToArcConfigurationResponseOutput() ArcConfigurationResponseOutput {
	return o
}

func (o ArcConfigurationResponseOutput) ToArcConfigurationResponseOutputWithContext(ctx context.Context) ArcConfigurationResponseOutput {
	return o
}

func (o ArcConfigurationResponseOutput) ArtifactStorageAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfigurationResponse) *string { return v.ArtifactStorageAccessMode }).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponseOutput) ArtifactStorageClassName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfigurationResponse) *string { return v.ArtifactStorageClassName }).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponseOutput) ArtifactStorageMountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfigurationResponse) *string { return v.ArtifactStorageMountPath }).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponseOutput) ArtifactStorageNodeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfigurationResponse) *string { return v.ArtifactStorageNodeName }).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponseOutput) ArtifactsStorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfigurationResponse) *string { return v.ArtifactsStorageType }).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponseOutput) FrontEndServiceConfiguration() FrontEndConfigurationResponsePtrOutput {
	return o.ApplyT(func(v ArcConfigurationResponse) *FrontEndConfigurationResponse { return v.FrontEndServiceConfiguration }).(FrontEndConfigurationResponsePtrOutput)
}

type ArcConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ArcConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcConfigurationResponse)(nil)).Elem()
}

func (o ArcConfigurationResponsePtrOutput) ToArcConfigurationResponsePtrOutput() ArcConfigurationResponsePtrOutput {
	return o
}

func (o ArcConfigurationResponsePtrOutput) ToArcConfigurationResponsePtrOutputWithContext(ctx context.Context) ArcConfigurationResponsePtrOutput {
	return o
}

func (o ArcConfigurationResponsePtrOutput) Elem() ArcConfigurationResponseOutput {
	return o.ApplyT(func(v *ArcConfigurationResponse) ArcConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ArcConfigurationResponse
		return ret
	}).(ArcConfigurationResponseOutput)
}

func (o ArcConfigurationResponsePtrOutput) ArtifactStorageAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactStorageAccessMode
	}).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponsePtrOutput) ArtifactStorageClassName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactStorageClassName
	}).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponsePtrOutput) ArtifactStorageMountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactStorageMountPath
	}).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponsePtrOutput) ArtifactStorageNodeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactStorageNodeName
	}).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponsePtrOutput) ArtifactsStorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactsStorageType
	}).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponsePtrOutput) FrontEndServiceConfiguration() FrontEndConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ArcConfigurationResponse) *FrontEndConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.FrontEndServiceConfiguration
	}).(FrontEndConfigurationResponsePtrOutput)
}

type ArmIdWrapperResponse struct {
	Id string `pulumi:"id"`
}

type ArmIdWrapperResponseOutput struct{ *pulumi.OutputState }

func (ArmIdWrapperResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmIdWrapperResponse)(nil)).Elem()
}

func (o ArmIdWrapperResponseOutput) ToArmIdWrapperResponseOutput() ArmIdWrapperResponseOutput {
	return o
}

func (o ArmIdWrapperResponseOutput) ToArmIdWrapperResponseOutputWithContext(ctx context.Context) ArmIdWrapperResponseOutput {
	return o
}

func (o ArmIdWrapperResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ArmIdWrapperResponse) string { return v.Id }).(pulumi.StringOutput)
}

type ArmIdWrapperResponsePtrOutput struct{ *pulumi.OutputState }

func (ArmIdWrapperResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmIdWrapperResponse)(nil)).Elem()
}

func (o ArmIdWrapperResponsePtrOutput) ToArmIdWrapperResponsePtrOutput() ArmIdWrapperResponsePtrOutput {
	return o
}

func (o ArmIdWrapperResponsePtrOutput) ToArmIdWrapperResponsePtrOutputWithContext(ctx context.Context) ArmIdWrapperResponsePtrOutput {
	return o
}

func (o ArmIdWrapperResponsePtrOutput) Elem() ArmIdWrapperResponseOutput {
	return o.ApplyT(func(v *ArmIdWrapperResponse) ArmIdWrapperResponse {
		if v != nil {
			return *v
		}
		var ret ArmIdWrapperResponse
		return ret
	}).(ArmIdWrapperResponseOutput)
}

func (o ArmIdWrapperResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmIdWrapperResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type ArmPlanResponse struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
	Version       *string `pulumi:"version"`
}

type ArmPlanResponseOutput struct{ *pulumi.OutputState }

func (ArmPlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmPlanResponse)(nil)).Elem()
}

func (o ArmPlanResponseOutput) ToArmPlanResponseOutput() ArmPlanResponseOutput {
	return o
}

func (o ArmPlanResponseOutput) ToArmPlanResponseOutputWithContext(ctx context.Context) ArmPlanResponseOutput {
	return o
}

func (o ArmPlanResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlanResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponseOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlanResponse) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponseOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlanResponse) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlanResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlanResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ArmPlanResponsePtrOutput struct{ *pulumi.OutputState }

func (ArmPlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmPlanResponse)(nil)).Elem()
}

func (o ArmPlanResponsePtrOutput) ToArmPlanResponsePtrOutput() ArmPlanResponsePtrOutput {
	return o
}

func (o ArmPlanResponsePtrOutput) ToArmPlanResponsePtrOutputWithContext(ctx context.Context) ArmPlanResponsePtrOutput {
	return o
}

func (o ArmPlanResponsePtrOutput) Elem() ArmPlanResponseOutput {
	return o.ApplyT(func(v *ArmPlanResponse) ArmPlanResponse {
		if v != nil {
			return *v
		}
		var ret ArmPlanResponse
		return ret
	}).(ArmPlanResponseOutput)
}

func (o ArmPlanResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponsePtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponsePtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type AuthPlatform struct {
	ConfigFilePath *string `pulumi:"configFilePath"`
	Enabled        *bool   `pulumi:"enabled"`
	RuntimeVersion *string `pulumi:"runtimeVersion"`
}





type AuthPlatformInput interface {
	pulumi.Input

	ToAuthPlatformOutput() AuthPlatformOutput
	ToAuthPlatformOutputWithContext(context.Context) AuthPlatformOutput
}

type AuthPlatformArgs struct {
	ConfigFilePath pulumi.StringPtrInput `pulumi:"configFilePath"`
	Enabled        pulumi.BoolPtrInput   `pulumi:"enabled"`
	RuntimeVersion pulumi.StringPtrInput `pulumi:"runtimeVersion"`
}

func (AuthPlatformArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthPlatform)(nil)).Elem()
}

func (i AuthPlatformArgs) ToAuthPlatformOutput() AuthPlatformOutput {
	return i.ToAuthPlatformOutputWithContext(context.Background())
}

func (i AuthPlatformArgs) ToAuthPlatformOutputWithContext(ctx context.Context) AuthPlatformOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthPlatformOutput)
}

func (i AuthPlatformArgs) ToAuthPlatformPtrOutput() AuthPlatformPtrOutput {
	return i.ToAuthPlatformPtrOutputWithContext(context.Background())
}

func (i AuthPlatformArgs) ToAuthPlatformPtrOutputWithContext(ctx context.Context) AuthPlatformPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthPlatformOutput).ToAuthPlatformPtrOutputWithContext(ctx)
}









type AuthPlatformPtrInput interface {
	pulumi.Input

	ToAuthPlatformPtrOutput() AuthPlatformPtrOutput
	ToAuthPlatformPtrOutputWithContext(context.Context) AuthPlatformPtrOutput
}

type authPlatformPtrType AuthPlatformArgs

func AuthPlatformPtr(v *AuthPlatformArgs) AuthPlatformPtrInput {
	return (*authPlatformPtrType)(v)
}

func (*authPlatformPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthPlatform)(nil)).Elem()
}

func (i *authPlatformPtrType) ToAuthPlatformPtrOutput() AuthPlatformPtrOutput {
	return i.ToAuthPlatformPtrOutputWithContext(context.Background())
}

func (i *authPlatformPtrType) ToAuthPlatformPtrOutputWithContext(ctx context.Context) AuthPlatformPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthPlatformPtrOutput)
}

type AuthPlatformOutput struct{ *pulumi.OutputState }

func (AuthPlatformOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthPlatform)(nil)).Elem()
}

func (o AuthPlatformOutput) ToAuthPlatformOutput() AuthPlatformOutput {
	return o
}

func (o AuthPlatformOutput) ToAuthPlatformOutputWithContext(ctx context.Context) AuthPlatformOutput {
	return o
}

func (o AuthPlatformOutput) ToAuthPlatformPtrOutput() AuthPlatformPtrOutput {
	return o.ToAuthPlatformPtrOutputWithContext(context.Background())
}

func (o AuthPlatformOutput) ToAuthPlatformPtrOutputWithContext(ctx context.Context) AuthPlatformPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthPlatform) *AuthPlatform {
		return &v
	}).(AuthPlatformPtrOutput)
}

func (o AuthPlatformOutput) ConfigFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthPlatform) *string { return v.ConfigFilePath }).(pulumi.StringPtrOutput)
}

func (o AuthPlatformOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AuthPlatform) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AuthPlatformOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthPlatform) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

type AuthPlatformPtrOutput struct{ *pulumi.OutputState }

func (AuthPlatformPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthPlatform)(nil)).Elem()
}

func (o AuthPlatformPtrOutput) ToAuthPlatformPtrOutput() AuthPlatformPtrOutput {
	return o
}

func (o AuthPlatformPtrOutput) ToAuthPlatformPtrOutputWithContext(ctx context.Context) AuthPlatformPtrOutput {
	return o
}

func (o AuthPlatformPtrOutput) Elem() AuthPlatformOutput {
	return o.ApplyT(func(v *AuthPlatform) AuthPlatform {
		if v != nil {
			return *v
		}
		var ret AuthPlatform
		return ret
	}).(AuthPlatformOutput)
}

func (o AuthPlatformPtrOutput) ConfigFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthPlatform) *string {
		if v == nil {
			return nil
		}
		return v.ConfigFilePath
	}).(pulumi.StringPtrOutput)
}

func (o AuthPlatformPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthPlatform) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AuthPlatformPtrOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthPlatform) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeVersion
	}).(pulumi.StringPtrOutput)
}

type AuthPlatformResponse struct {
	ConfigFilePath *string `pulumi:"configFilePath"`
	Enabled        *bool   `pulumi:"enabled"`
	RuntimeVersion *string `pulumi:"runtimeVersion"`
}

type AuthPlatformResponseOutput struct{ *pulumi.OutputState }

func (AuthPlatformResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthPlatformResponse)(nil)).Elem()
}

func (o AuthPlatformResponseOutput) ToAuthPlatformResponseOutput() AuthPlatformResponseOutput {
	return o
}

func (o AuthPlatformResponseOutput) ToAuthPlatformResponseOutputWithContext(ctx context.Context) AuthPlatformResponseOutput {
	return o
}

func (o AuthPlatformResponseOutput) ConfigFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthPlatformResponse) *string { return v.ConfigFilePath }).(pulumi.StringPtrOutput)
}

func (o AuthPlatformResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AuthPlatformResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AuthPlatformResponseOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthPlatformResponse) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

type AuthPlatformResponsePtrOutput struct{ *pulumi.OutputState }

func (AuthPlatformResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthPlatformResponse)(nil)).Elem()
}

func (o AuthPlatformResponsePtrOutput) ToAuthPlatformResponsePtrOutput() AuthPlatformResponsePtrOutput {
	return o
}

func (o AuthPlatformResponsePtrOutput) ToAuthPlatformResponsePtrOutputWithContext(ctx context.Context) AuthPlatformResponsePtrOutput {
	return o
}

func (o AuthPlatformResponsePtrOutput) Elem() AuthPlatformResponseOutput {
	return o.ApplyT(func(v *AuthPlatformResponse) AuthPlatformResponse {
		if v != nil {
			return *v
		}
		var ret AuthPlatformResponse
		return ret
	}).(AuthPlatformResponseOutput)
}

func (o AuthPlatformResponsePtrOutput) ConfigFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthPlatformResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConfigFilePath
	}).(pulumi.StringPtrOutput)
}

func (o AuthPlatformResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthPlatformResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AuthPlatformResponsePtrOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthPlatformResponse) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeVersion
	}).(pulumi.StringPtrOutput)
}

type AutoHealActions struct {
	ActionType              *AutoHealActionType   `pulumi:"actionType"`
	CustomAction            *AutoHealCustomAction `pulumi:"customAction"`
	MinProcessExecutionTime *string               `pulumi:"minProcessExecutionTime"`
}





type AutoHealActionsInput interface {
	pulumi.Input

	ToAutoHealActionsOutput() AutoHealActionsOutput
	ToAutoHealActionsOutputWithContext(context.Context) AutoHealActionsOutput
}

type AutoHealActionsArgs struct {
	ActionType              AutoHealActionTypePtrInput   `pulumi:"actionType"`
	CustomAction            AutoHealCustomActionPtrInput `pulumi:"customAction"`
	MinProcessExecutionTime pulumi.StringPtrInput        `pulumi:"minProcessExecutionTime"`
}

func (AutoHealActionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealActions)(nil)).Elem()
}

func (i AutoHealActionsArgs) ToAutoHealActionsOutput() AutoHealActionsOutput {
	return i.ToAutoHealActionsOutputWithContext(context.Background())
}

func (i AutoHealActionsArgs) ToAutoHealActionsOutputWithContext(ctx context.Context) AutoHealActionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealActionsOutput)
}

func (i AutoHealActionsArgs) ToAutoHealActionsPtrOutput() AutoHealActionsPtrOutput {
	return i.ToAutoHealActionsPtrOutputWithContext(context.Background())
}

func (i AutoHealActionsArgs) ToAutoHealActionsPtrOutputWithContext(ctx context.Context) AutoHealActionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealActionsOutput).ToAutoHealActionsPtrOutputWithContext(ctx)
}









type AutoHealActionsPtrInput interface {
	pulumi.Input

	ToAutoHealActionsPtrOutput() AutoHealActionsPtrOutput
	ToAutoHealActionsPtrOutputWithContext(context.Context) AutoHealActionsPtrOutput
}

type autoHealActionsPtrType AutoHealActionsArgs

func AutoHealActionsPtr(v *AutoHealActionsArgs) AutoHealActionsPtrInput {
	return (*autoHealActionsPtrType)(v)
}

func (*autoHealActionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealActions)(nil)).Elem()
}

func (i *autoHealActionsPtrType) ToAutoHealActionsPtrOutput() AutoHealActionsPtrOutput {
	return i.ToAutoHealActionsPtrOutputWithContext(context.Background())
}

func (i *autoHealActionsPtrType) ToAutoHealActionsPtrOutputWithContext(ctx context.Context) AutoHealActionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealActionsPtrOutput)
}

type AutoHealActionsOutput struct{ *pulumi.OutputState }

func (AutoHealActionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealActions)(nil)).Elem()
}

func (o AutoHealActionsOutput) ToAutoHealActionsOutput() AutoHealActionsOutput {
	return o
}

func (o AutoHealActionsOutput) ToAutoHealActionsOutputWithContext(ctx context.Context) AutoHealActionsOutput {
	return o
}

func (o AutoHealActionsOutput) ToAutoHealActionsPtrOutput() AutoHealActionsPtrOutput {
	return o.ToAutoHealActionsPtrOutputWithContext(context.Background())
}

func (o AutoHealActionsOutput) ToAutoHealActionsPtrOutputWithContext(ctx context.Context) AutoHealActionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealActions) *AutoHealActions {
		return &v
	}).(AutoHealActionsPtrOutput)
}

func (o AutoHealActionsOutput) ActionType() AutoHealActionTypePtrOutput {
	return o.ApplyT(func(v AutoHealActions) *AutoHealActionType { return v.ActionType }).(AutoHealActionTypePtrOutput)
}

func (o AutoHealActionsOutput) CustomAction() AutoHealCustomActionPtrOutput {
	return o.ApplyT(func(v AutoHealActions) *AutoHealCustomAction { return v.CustomAction }).(AutoHealCustomActionPtrOutput)
}

func (o AutoHealActionsOutput) MinProcessExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealActions) *string { return v.MinProcessExecutionTime }).(pulumi.StringPtrOutput)
}

type AutoHealActionsPtrOutput struct{ *pulumi.OutputState }

func (AutoHealActionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealActions)(nil)).Elem()
}

func (o AutoHealActionsPtrOutput) ToAutoHealActionsPtrOutput() AutoHealActionsPtrOutput {
	return o
}

func (o AutoHealActionsPtrOutput) ToAutoHealActionsPtrOutputWithContext(ctx context.Context) AutoHealActionsPtrOutput {
	return o
}

func (o AutoHealActionsPtrOutput) Elem() AutoHealActionsOutput {
	return o.ApplyT(func(v *AutoHealActions) AutoHealActions {
		if v != nil {
			return *v
		}
		var ret AutoHealActions
		return ret
	}).(AutoHealActionsOutput)
}

func (o AutoHealActionsPtrOutput) ActionType() AutoHealActionTypePtrOutput {
	return o.ApplyT(func(v *AutoHealActions) *AutoHealActionType {
		if v == nil {
			return nil
		}
		return v.ActionType
	}).(AutoHealActionTypePtrOutput)
}

func (o AutoHealActionsPtrOutput) CustomAction() AutoHealCustomActionPtrOutput {
	return o.ApplyT(func(v *AutoHealActions) *AutoHealCustomAction {
		if v == nil {
			return nil
		}
		return v.CustomAction
	}).(AutoHealCustomActionPtrOutput)
}

func (o AutoHealActionsPtrOutput) MinProcessExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealActions) *string {
		if v == nil {
			return nil
		}
		return v.MinProcessExecutionTime
	}).(pulumi.StringPtrOutput)
}

type AutoHealActionsResponse struct {
	ActionType              *string                       `pulumi:"actionType"`
	CustomAction            *AutoHealCustomActionResponse `pulumi:"customAction"`
	MinProcessExecutionTime *string                       `pulumi:"minProcessExecutionTime"`
}

type AutoHealActionsResponseOutput struct{ *pulumi.OutputState }

func (AutoHealActionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealActionsResponse)(nil)).Elem()
}

func (o AutoHealActionsResponseOutput) ToAutoHealActionsResponseOutput() AutoHealActionsResponseOutput {
	return o
}

func (o AutoHealActionsResponseOutput) ToAutoHealActionsResponseOutputWithContext(ctx context.Context) AutoHealActionsResponseOutput {
	return o
}

func (o AutoHealActionsResponseOutput) ActionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealActionsResponse) *string { return v.ActionType }).(pulumi.StringPtrOutput)
}

func (o AutoHealActionsResponseOutput) CustomAction() AutoHealCustomActionResponsePtrOutput {
	return o.ApplyT(func(v AutoHealActionsResponse) *AutoHealCustomActionResponse { return v.CustomAction }).(AutoHealCustomActionResponsePtrOutput)
}

func (o AutoHealActionsResponseOutput) MinProcessExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealActionsResponse) *string { return v.MinProcessExecutionTime }).(pulumi.StringPtrOutput)
}

type AutoHealActionsResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoHealActionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealActionsResponse)(nil)).Elem()
}

func (o AutoHealActionsResponsePtrOutput) ToAutoHealActionsResponsePtrOutput() AutoHealActionsResponsePtrOutput {
	return o
}

func (o AutoHealActionsResponsePtrOutput) ToAutoHealActionsResponsePtrOutputWithContext(ctx context.Context) AutoHealActionsResponsePtrOutput {
	return o
}

func (o AutoHealActionsResponsePtrOutput) Elem() AutoHealActionsResponseOutput {
	return o.ApplyT(func(v *AutoHealActionsResponse) AutoHealActionsResponse {
		if v != nil {
			return *v
		}
		var ret AutoHealActionsResponse
		return ret
	}).(AutoHealActionsResponseOutput)
}

func (o AutoHealActionsResponsePtrOutput) ActionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealActionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionType
	}).(pulumi.StringPtrOutput)
}

func (o AutoHealActionsResponsePtrOutput) CustomAction() AutoHealCustomActionResponsePtrOutput {
	return o.ApplyT(func(v *AutoHealActionsResponse) *AutoHealCustomActionResponse {
		if v == nil {
			return nil
		}
		return v.CustomAction
	}).(AutoHealCustomActionResponsePtrOutput)
}

func (o AutoHealActionsResponsePtrOutput) MinProcessExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealActionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.MinProcessExecutionTime
	}).(pulumi.StringPtrOutput)
}

type AutoHealCustomAction struct {
	Exe        *string `pulumi:"exe"`
	Parameters *string `pulumi:"parameters"`
}





type AutoHealCustomActionInput interface {
	pulumi.Input

	ToAutoHealCustomActionOutput() AutoHealCustomActionOutput
	ToAutoHealCustomActionOutputWithContext(context.Context) AutoHealCustomActionOutput
}

type AutoHealCustomActionArgs struct {
	Exe        pulumi.StringPtrInput `pulumi:"exe"`
	Parameters pulumi.StringPtrInput `pulumi:"parameters"`
}

func (AutoHealCustomActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealCustomAction)(nil)).Elem()
}

func (i AutoHealCustomActionArgs) ToAutoHealCustomActionOutput() AutoHealCustomActionOutput {
	return i.ToAutoHealCustomActionOutputWithContext(context.Background())
}

func (i AutoHealCustomActionArgs) ToAutoHealCustomActionOutputWithContext(ctx context.Context) AutoHealCustomActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealCustomActionOutput)
}

func (i AutoHealCustomActionArgs) ToAutoHealCustomActionPtrOutput() AutoHealCustomActionPtrOutput {
	return i.ToAutoHealCustomActionPtrOutputWithContext(context.Background())
}

func (i AutoHealCustomActionArgs) ToAutoHealCustomActionPtrOutputWithContext(ctx context.Context) AutoHealCustomActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealCustomActionOutput).ToAutoHealCustomActionPtrOutputWithContext(ctx)
}









type AutoHealCustomActionPtrInput interface {
	pulumi.Input

	ToAutoHealCustomActionPtrOutput() AutoHealCustomActionPtrOutput
	ToAutoHealCustomActionPtrOutputWithContext(context.Context) AutoHealCustomActionPtrOutput
}

type autoHealCustomActionPtrType AutoHealCustomActionArgs

func AutoHealCustomActionPtr(v *AutoHealCustomActionArgs) AutoHealCustomActionPtrInput {
	return (*autoHealCustomActionPtrType)(v)
}

func (*autoHealCustomActionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealCustomAction)(nil)).Elem()
}

func (i *autoHealCustomActionPtrType) ToAutoHealCustomActionPtrOutput() AutoHealCustomActionPtrOutput {
	return i.ToAutoHealCustomActionPtrOutputWithContext(context.Background())
}

func (i *autoHealCustomActionPtrType) ToAutoHealCustomActionPtrOutputWithContext(ctx context.Context) AutoHealCustomActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealCustomActionPtrOutput)
}

type AutoHealCustomActionOutput struct{ *pulumi.OutputState }

func (AutoHealCustomActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealCustomAction)(nil)).Elem()
}

func (o AutoHealCustomActionOutput) ToAutoHealCustomActionOutput() AutoHealCustomActionOutput {
	return o
}

func (o AutoHealCustomActionOutput) ToAutoHealCustomActionOutputWithContext(ctx context.Context) AutoHealCustomActionOutput {
	return o
}

func (o AutoHealCustomActionOutput) ToAutoHealCustomActionPtrOutput() AutoHealCustomActionPtrOutput {
	return o.ToAutoHealCustomActionPtrOutputWithContext(context.Background())
}

func (o AutoHealCustomActionOutput) ToAutoHealCustomActionPtrOutputWithContext(ctx context.Context) AutoHealCustomActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealCustomAction) *AutoHealCustomAction {
		return &v
	}).(AutoHealCustomActionPtrOutput)
}

func (o AutoHealCustomActionOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealCustomAction) *string { return v.Exe }).(pulumi.StringPtrOutput)
}

func (o AutoHealCustomActionOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealCustomAction) *string { return v.Parameters }).(pulumi.StringPtrOutput)
}

type AutoHealCustomActionPtrOutput struct{ *pulumi.OutputState }

func (AutoHealCustomActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealCustomAction)(nil)).Elem()
}

func (o AutoHealCustomActionPtrOutput) ToAutoHealCustomActionPtrOutput() AutoHealCustomActionPtrOutput {
	return o
}

func (o AutoHealCustomActionPtrOutput) ToAutoHealCustomActionPtrOutputWithContext(ctx context.Context) AutoHealCustomActionPtrOutput {
	return o
}

func (o AutoHealCustomActionPtrOutput) Elem() AutoHealCustomActionOutput {
	return o.ApplyT(func(v *AutoHealCustomAction) AutoHealCustomAction {
		if v != nil {
			return *v
		}
		var ret AutoHealCustomAction
		return ret
	}).(AutoHealCustomActionOutput)
}

func (o AutoHealCustomActionPtrOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealCustomAction) *string {
		if v == nil {
			return nil
		}
		return v.Exe
	}).(pulumi.StringPtrOutput)
}

func (o AutoHealCustomActionPtrOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealCustomAction) *string {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.StringPtrOutput)
}

type AutoHealCustomActionResponse struct {
	Exe        *string `pulumi:"exe"`
	Parameters *string `pulumi:"parameters"`
}

type AutoHealCustomActionResponseOutput struct{ *pulumi.OutputState }

func (AutoHealCustomActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealCustomActionResponse)(nil)).Elem()
}

func (o AutoHealCustomActionResponseOutput) ToAutoHealCustomActionResponseOutput() AutoHealCustomActionResponseOutput {
	return o
}

func (o AutoHealCustomActionResponseOutput) ToAutoHealCustomActionResponseOutputWithContext(ctx context.Context) AutoHealCustomActionResponseOutput {
	return o
}

func (o AutoHealCustomActionResponseOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealCustomActionResponse) *string { return v.Exe }).(pulumi.StringPtrOutput)
}

func (o AutoHealCustomActionResponseOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealCustomActionResponse) *string { return v.Parameters }).(pulumi.StringPtrOutput)
}

type AutoHealCustomActionResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoHealCustomActionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealCustomActionResponse)(nil)).Elem()
}

func (o AutoHealCustomActionResponsePtrOutput) ToAutoHealCustomActionResponsePtrOutput() AutoHealCustomActionResponsePtrOutput {
	return o
}

func (o AutoHealCustomActionResponsePtrOutput) ToAutoHealCustomActionResponsePtrOutputWithContext(ctx context.Context) AutoHealCustomActionResponsePtrOutput {
	return o
}

func (o AutoHealCustomActionResponsePtrOutput) Elem() AutoHealCustomActionResponseOutput {
	return o.ApplyT(func(v *AutoHealCustomActionResponse) AutoHealCustomActionResponse {
		if v != nil {
			return *v
		}
		var ret AutoHealCustomActionResponse
		return ret
	}).(AutoHealCustomActionResponseOutput)
}

func (o AutoHealCustomActionResponsePtrOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealCustomActionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Exe
	}).(pulumi.StringPtrOutput)
}

func (o AutoHealCustomActionResponsePtrOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealCustomActionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.StringPtrOutput)
}

type AutoHealRules struct {
	Actions  *AutoHealActions  `pulumi:"actions"`
	Triggers *AutoHealTriggers `pulumi:"triggers"`
}





type AutoHealRulesInput interface {
	pulumi.Input

	ToAutoHealRulesOutput() AutoHealRulesOutput
	ToAutoHealRulesOutputWithContext(context.Context) AutoHealRulesOutput
}

type AutoHealRulesArgs struct {
	Actions  AutoHealActionsPtrInput  `pulumi:"actions"`
	Triggers AutoHealTriggersPtrInput `pulumi:"triggers"`
}

func (AutoHealRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealRules)(nil)).Elem()
}

func (i AutoHealRulesArgs) ToAutoHealRulesOutput() AutoHealRulesOutput {
	return i.ToAutoHealRulesOutputWithContext(context.Background())
}

func (i AutoHealRulesArgs) ToAutoHealRulesOutputWithContext(ctx context.Context) AutoHealRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealRulesOutput)
}

func (i AutoHealRulesArgs) ToAutoHealRulesPtrOutput() AutoHealRulesPtrOutput {
	return i.ToAutoHealRulesPtrOutputWithContext(context.Background())
}

func (i AutoHealRulesArgs) ToAutoHealRulesPtrOutputWithContext(ctx context.Context) AutoHealRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealRulesOutput).ToAutoHealRulesPtrOutputWithContext(ctx)
}









type AutoHealRulesPtrInput interface {
	pulumi.Input

	ToAutoHealRulesPtrOutput() AutoHealRulesPtrOutput
	ToAutoHealRulesPtrOutputWithContext(context.Context) AutoHealRulesPtrOutput
}

type autoHealRulesPtrType AutoHealRulesArgs

func AutoHealRulesPtr(v *AutoHealRulesArgs) AutoHealRulesPtrInput {
	return (*autoHealRulesPtrType)(v)
}

func (*autoHealRulesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealRules)(nil)).Elem()
}

func (i *autoHealRulesPtrType) ToAutoHealRulesPtrOutput() AutoHealRulesPtrOutput {
	return i.ToAutoHealRulesPtrOutputWithContext(context.Background())
}

func (i *autoHealRulesPtrType) ToAutoHealRulesPtrOutputWithContext(ctx context.Context) AutoHealRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealRulesPtrOutput)
}

type AutoHealRulesOutput struct{ *pulumi.OutputState }

func (AutoHealRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealRules)(nil)).Elem()
}

func (o AutoHealRulesOutput) ToAutoHealRulesOutput() AutoHealRulesOutput {
	return o
}

func (o AutoHealRulesOutput) ToAutoHealRulesOutputWithContext(ctx context.Context) AutoHealRulesOutput {
	return o
}

func (o AutoHealRulesOutput) ToAutoHealRulesPtrOutput() AutoHealRulesPtrOutput {
	return o.ToAutoHealRulesPtrOutputWithContext(context.Background())
}

func (o AutoHealRulesOutput) ToAutoHealRulesPtrOutputWithContext(ctx context.Context) AutoHealRulesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealRules) *AutoHealRules {
		return &v
	}).(AutoHealRulesPtrOutput)
}

func (o AutoHealRulesOutput) Actions() AutoHealActionsPtrOutput {
	return o.ApplyT(func(v AutoHealRules) *AutoHealActions { return v.Actions }).(AutoHealActionsPtrOutput)
}

func (o AutoHealRulesOutput) Triggers() AutoHealTriggersPtrOutput {
	return o.ApplyT(func(v AutoHealRules) *AutoHealTriggers { return v.Triggers }).(AutoHealTriggersPtrOutput)
}

type AutoHealRulesPtrOutput struct{ *pulumi.OutputState }

func (AutoHealRulesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealRules)(nil)).Elem()
}

func (o AutoHealRulesPtrOutput) ToAutoHealRulesPtrOutput() AutoHealRulesPtrOutput {
	return o
}

func (o AutoHealRulesPtrOutput) ToAutoHealRulesPtrOutputWithContext(ctx context.Context) AutoHealRulesPtrOutput {
	return o
}

func (o AutoHealRulesPtrOutput) Elem() AutoHealRulesOutput {
	return o.ApplyT(func(v *AutoHealRules) AutoHealRules {
		if v != nil {
			return *v
		}
		var ret AutoHealRules
		return ret
	}).(AutoHealRulesOutput)
}

func (o AutoHealRulesPtrOutput) Actions() AutoHealActionsPtrOutput {
	return o.ApplyT(func(v *AutoHealRules) *AutoHealActions {
		if v == nil {
			return nil
		}
		return v.Actions
	}).(AutoHealActionsPtrOutput)
}

func (o AutoHealRulesPtrOutput) Triggers() AutoHealTriggersPtrOutput {
	return o.ApplyT(func(v *AutoHealRules) *AutoHealTriggers {
		if v == nil {
			return nil
		}
		return v.Triggers
	}).(AutoHealTriggersPtrOutput)
}

type AutoHealRulesResponse struct {
	Actions  *AutoHealActionsResponse  `pulumi:"actions"`
	Triggers *AutoHealTriggersResponse `pulumi:"triggers"`
}

type AutoHealRulesResponseOutput struct{ *pulumi.OutputState }

func (AutoHealRulesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealRulesResponse)(nil)).Elem()
}

func (o AutoHealRulesResponseOutput) ToAutoHealRulesResponseOutput() AutoHealRulesResponseOutput {
	return o
}

func (o AutoHealRulesResponseOutput) ToAutoHealRulesResponseOutputWithContext(ctx context.Context) AutoHealRulesResponseOutput {
	return o
}

func (o AutoHealRulesResponseOutput) Actions() AutoHealActionsResponsePtrOutput {
	return o.ApplyT(func(v AutoHealRulesResponse) *AutoHealActionsResponse { return v.Actions }).(AutoHealActionsResponsePtrOutput)
}

func (o AutoHealRulesResponseOutput) Triggers() AutoHealTriggersResponsePtrOutput {
	return o.ApplyT(func(v AutoHealRulesResponse) *AutoHealTriggersResponse { return v.Triggers }).(AutoHealTriggersResponsePtrOutput)
}

type AutoHealRulesResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoHealRulesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealRulesResponse)(nil)).Elem()
}

func (o AutoHealRulesResponsePtrOutput) ToAutoHealRulesResponsePtrOutput() AutoHealRulesResponsePtrOutput {
	return o
}

func (o AutoHealRulesResponsePtrOutput) ToAutoHealRulesResponsePtrOutputWithContext(ctx context.Context) AutoHealRulesResponsePtrOutput {
	return o
}

func (o AutoHealRulesResponsePtrOutput) Elem() AutoHealRulesResponseOutput {
	return o.ApplyT(func(v *AutoHealRulesResponse) AutoHealRulesResponse {
		if v != nil {
			return *v
		}
		var ret AutoHealRulesResponse
		return ret
	}).(AutoHealRulesResponseOutput)
}

func (o AutoHealRulesResponsePtrOutput) Actions() AutoHealActionsResponsePtrOutput {
	return o.ApplyT(func(v *AutoHealRulesResponse) *AutoHealActionsResponse {
		if v == nil {
			return nil
		}
		return v.Actions
	}).(AutoHealActionsResponsePtrOutput)
}

func (o AutoHealRulesResponsePtrOutput) Triggers() AutoHealTriggersResponsePtrOutput {
	return o.ApplyT(func(v *AutoHealRulesResponse) *AutoHealTriggersResponse {
		if v == nil {
			return nil
		}
		return v.Triggers
	}).(AutoHealTriggersResponsePtrOutput)
}

type AutoHealTriggers struct {
	PrivateBytesInKB     *int                           `pulumi:"privateBytesInKB"`
	Requests             *RequestsBasedTrigger          `pulumi:"requests"`
	SlowRequests         *SlowRequestsBasedTrigger      `pulumi:"slowRequests"`
	SlowRequestsWithPath []SlowRequestsBasedTrigger     `pulumi:"slowRequestsWithPath"`
	StatusCodes          []StatusCodesBasedTrigger      `pulumi:"statusCodes"`
	StatusCodesRange     []StatusCodesRangeBasedTrigger `pulumi:"statusCodesRange"`
}





type AutoHealTriggersInput interface {
	pulumi.Input

	ToAutoHealTriggersOutput() AutoHealTriggersOutput
	ToAutoHealTriggersOutputWithContext(context.Context) AutoHealTriggersOutput
}

type AutoHealTriggersArgs struct {
	PrivateBytesInKB     pulumi.IntPtrInput                     `pulumi:"privateBytesInKB"`
	Requests             RequestsBasedTriggerPtrInput           `pulumi:"requests"`
	SlowRequests         SlowRequestsBasedTriggerPtrInput       `pulumi:"slowRequests"`
	SlowRequestsWithPath SlowRequestsBasedTriggerArrayInput     `pulumi:"slowRequestsWithPath"`
	StatusCodes          StatusCodesBasedTriggerArrayInput      `pulumi:"statusCodes"`
	StatusCodesRange     StatusCodesRangeBasedTriggerArrayInput `pulumi:"statusCodesRange"`
}

func (AutoHealTriggersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealTriggers)(nil)).Elem()
}

func (i AutoHealTriggersArgs) ToAutoHealTriggersOutput() AutoHealTriggersOutput {
	return i.ToAutoHealTriggersOutputWithContext(context.Background())
}

func (i AutoHealTriggersArgs) ToAutoHealTriggersOutputWithContext(ctx context.Context) AutoHealTriggersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealTriggersOutput)
}

func (i AutoHealTriggersArgs) ToAutoHealTriggersPtrOutput() AutoHealTriggersPtrOutput {
	return i.ToAutoHealTriggersPtrOutputWithContext(context.Background())
}

func (i AutoHealTriggersArgs) ToAutoHealTriggersPtrOutputWithContext(ctx context.Context) AutoHealTriggersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealTriggersOutput).ToAutoHealTriggersPtrOutputWithContext(ctx)
}









type AutoHealTriggersPtrInput interface {
	pulumi.Input

	ToAutoHealTriggersPtrOutput() AutoHealTriggersPtrOutput
	ToAutoHealTriggersPtrOutputWithContext(context.Context) AutoHealTriggersPtrOutput
}

type autoHealTriggersPtrType AutoHealTriggersArgs

func AutoHealTriggersPtr(v *AutoHealTriggersArgs) AutoHealTriggersPtrInput {
	return (*autoHealTriggersPtrType)(v)
}

func (*autoHealTriggersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealTriggers)(nil)).Elem()
}

func (i *autoHealTriggersPtrType) ToAutoHealTriggersPtrOutput() AutoHealTriggersPtrOutput {
	return i.ToAutoHealTriggersPtrOutputWithContext(context.Background())
}

func (i *autoHealTriggersPtrType) ToAutoHealTriggersPtrOutputWithContext(ctx context.Context) AutoHealTriggersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealTriggersPtrOutput)
}

type AutoHealTriggersOutput struct{ *pulumi.OutputState }

func (AutoHealTriggersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealTriggers)(nil)).Elem()
}

func (o AutoHealTriggersOutput) ToAutoHealTriggersOutput() AutoHealTriggersOutput {
	return o
}

func (o AutoHealTriggersOutput) ToAutoHealTriggersOutputWithContext(ctx context.Context) AutoHealTriggersOutput {
	return o
}

func (o AutoHealTriggersOutput) ToAutoHealTriggersPtrOutput() AutoHealTriggersPtrOutput {
	return o.ToAutoHealTriggersPtrOutputWithContext(context.Background())
}

func (o AutoHealTriggersOutput) ToAutoHealTriggersPtrOutputWithContext(ctx context.Context) AutoHealTriggersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealTriggers) *AutoHealTriggers {
		return &v
	}).(AutoHealTriggersPtrOutput)
}

func (o AutoHealTriggersOutput) PrivateBytesInKB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoHealTriggers) *int { return v.PrivateBytesInKB }).(pulumi.IntPtrOutput)
}

func (o AutoHealTriggersOutput) Requests() RequestsBasedTriggerPtrOutput {
	return o.ApplyT(func(v AutoHealTriggers) *RequestsBasedTrigger { return v.Requests }).(RequestsBasedTriggerPtrOutput)
}

func (o AutoHealTriggersOutput) SlowRequests() SlowRequestsBasedTriggerPtrOutput {
	return o.ApplyT(func(v AutoHealTriggers) *SlowRequestsBasedTrigger { return v.SlowRequests }).(SlowRequestsBasedTriggerPtrOutput)
}

func (o AutoHealTriggersOutput) SlowRequestsWithPath() SlowRequestsBasedTriggerArrayOutput {
	return o.ApplyT(func(v AutoHealTriggers) []SlowRequestsBasedTrigger { return v.SlowRequestsWithPath }).(SlowRequestsBasedTriggerArrayOutput)
}

func (o AutoHealTriggersOutput) StatusCodes() StatusCodesBasedTriggerArrayOutput {
	return o.ApplyT(func(v AutoHealTriggers) []StatusCodesBasedTrigger { return v.StatusCodes }).(StatusCodesBasedTriggerArrayOutput)
}

func (o AutoHealTriggersOutput) StatusCodesRange() StatusCodesRangeBasedTriggerArrayOutput {
	return o.ApplyT(func(v AutoHealTriggers) []StatusCodesRangeBasedTrigger { return v.StatusCodesRange }).(StatusCodesRangeBasedTriggerArrayOutput)
}

type AutoHealTriggersPtrOutput struct{ *pulumi.OutputState }

func (AutoHealTriggersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealTriggers)(nil)).Elem()
}

func (o AutoHealTriggersPtrOutput) ToAutoHealTriggersPtrOutput() AutoHealTriggersPtrOutput {
	return o
}

func (o AutoHealTriggersPtrOutput) ToAutoHealTriggersPtrOutputWithContext(ctx context.Context) AutoHealTriggersPtrOutput {
	return o
}

func (o AutoHealTriggersPtrOutput) Elem() AutoHealTriggersOutput {
	return o.ApplyT(func(v *AutoHealTriggers) AutoHealTriggers {
		if v != nil {
			return *v
		}
		var ret AutoHealTriggers
		return ret
	}).(AutoHealTriggersOutput)
}

func (o AutoHealTriggersPtrOutput) PrivateBytesInKB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoHealTriggers) *int {
		if v == nil {
			return nil
		}
		return v.PrivateBytesInKB
	}).(pulumi.IntPtrOutput)
}

func (o AutoHealTriggersPtrOutput) Requests() RequestsBasedTriggerPtrOutput {
	return o.ApplyT(func(v *AutoHealTriggers) *RequestsBasedTrigger {
		if v == nil {
			return nil
		}
		return v.Requests
	}).(RequestsBasedTriggerPtrOutput)
}

func (o AutoHealTriggersPtrOutput) SlowRequests() SlowRequestsBasedTriggerPtrOutput {
	return o.ApplyT(func(v *AutoHealTriggers) *SlowRequestsBasedTrigger {
		if v == nil {
			return nil
		}
		return v.SlowRequests
	}).(SlowRequestsBasedTriggerPtrOutput)
}

func (o AutoHealTriggersPtrOutput) SlowRequestsWithPath() SlowRequestsBasedTriggerArrayOutput {
	return o.ApplyT(func(v *AutoHealTriggers) []SlowRequestsBasedTrigger {
		if v == nil {
			return nil
		}
		return v.SlowRequestsWithPath
	}).(SlowRequestsBasedTriggerArrayOutput)
}

func (o AutoHealTriggersPtrOutput) StatusCodes() StatusCodesBasedTriggerArrayOutput {
	return o.ApplyT(func(v *AutoHealTriggers) []StatusCodesBasedTrigger {
		if v == nil {
			return nil
		}
		return v.StatusCodes
	}).(StatusCodesBasedTriggerArrayOutput)
}

func (o AutoHealTriggersPtrOutput) StatusCodesRange() StatusCodesRangeBasedTriggerArrayOutput {
	return o.ApplyT(func(v *AutoHealTriggers) []StatusCodesRangeBasedTrigger {
		if v == nil {
			return nil
		}
		return v.StatusCodesRange
	}).(StatusCodesRangeBasedTriggerArrayOutput)
}

type AutoHealTriggersResponse struct {
	PrivateBytesInKB     *int                                   `pulumi:"privateBytesInKB"`
	Requests             *RequestsBasedTriggerResponse          `pulumi:"requests"`
	SlowRequests         *SlowRequestsBasedTriggerResponse      `pulumi:"slowRequests"`
	SlowRequestsWithPath []SlowRequestsBasedTriggerResponse     `pulumi:"slowRequestsWithPath"`
	StatusCodes          []StatusCodesBasedTriggerResponse      `pulumi:"statusCodes"`
	StatusCodesRange     []StatusCodesRangeBasedTriggerResponse `pulumi:"statusCodesRange"`
}

type AutoHealTriggersResponseOutput struct{ *pulumi.OutputState }

func (AutoHealTriggersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealTriggersResponse)(nil)).Elem()
}

func (o AutoHealTriggersResponseOutput) ToAutoHealTriggersResponseOutput() AutoHealTriggersResponseOutput {
	return o
}

func (o AutoHealTriggersResponseOutput) ToAutoHealTriggersResponseOutputWithContext(ctx context.Context) AutoHealTriggersResponseOutput {
	return o
}

func (o AutoHealTriggersResponseOutput) PrivateBytesInKB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoHealTriggersResponse) *int { return v.PrivateBytesInKB }).(pulumi.IntPtrOutput)
}

func (o AutoHealTriggersResponseOutput) Requests() RequestsBasedTriggerResponsePtrOutput {
	return o.ApplyT(func(v AutoHealTriggersResponse) *RequestsBasedTriggerResponse { return v.Requests }).(RequestsBasedTriggerResponsePtrOutput)
}

func (o AutoHealTriggersResponseOutput) SlowRequests() SlowRequestsBasedTriggerResponsePtrOutput {
	return o.ApplyT(func(v AutoHealTriggersResponse) *SlowRequestsBasedTriggerResponse { return v.SlowRequests }).(SlowRequestsBasedTriggerResponsePtrOutput)
}

func (o AutoHealTriggersResponseOutput) SlowRequestsWithPath() SlowRequestsBasedTriggerResponseArrayOutput {
	return o.ApplyT(func(v AutoHealTriggersResponse) []SlowRequestsBasedTriggerResponse { return v.SlowRequestsWithPath }).(SlowRequestsBasedTriggerResponseArrayOutput)
}

func (o AutoHealTriggersResponseOutput) StatusCodes() StatusCodesBasedTriggerResponseArrayOutput {
	return o.ApplyT(func(v AutoHealTriggersResponse) []StatusCodesBasedTriggerResponse { return v.StatusCodes }).(StatusCodesBasedTriggerResponseArrayOutput)
}

func (o AutoHealTriggersResponseOutput) StatusCodesRange() StatusCodesRangeBasedTriggerResponseArrayOutput {
	return o.ApplyT(func(v AutoHealTriggersResponse) []StatusCodesRangeBasedTriggerResponse { return v.StatusCodesRange }).(StatusCodesRangeBasedTriggerResponseArrayOutput)
}

type AutoHealTriggersResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoHealTriggersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealTriggersResponse)(nil)).Elem()
}

func (o AutoHealTriggersResponsePtrOutput) ToAutoHealTriggersResponsePtrOutput() AutoHealTriggersResponsePtrOutput {
	return o
}

func (o AutoHealTriggersResponsePtrOutput) ToAutoHealTriggersResponsePtrOutputWithContext(ctx context.Context) AutoHealTriggersResponsePtrOutput {
	return o
}

func (o AutoHealTriggersResponsePtrOutput) Elem() AutoHealTriggersResponseOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) AutoHealTriggersResponse {
		if v != nil {
			return *v
		}
		var ret AutoHealTriggersResponse
		return ret
	}).(AutoHealTriggersResponseOutput)
}

func (o AutoHealTriggersResponsePtrOutput) PrivateBytesInKB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) *int {
		if v == nil {
			return nil
		}
		return v.PrivateBytesInKB
	}).(pulumi.IntPtrOutput)
}

func (o AutoHealTriggersResponsePtrOutput) Requests() RequestsBasedTriggerResponsePtrOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) *RequestsBasedTriggerResponse {
		if v == nil {
			return nil
		}
		return v.Requests
	}).(RequestsBasedTriggerResponsePtrOutput)
}

func (o AutoHealTriggersResponsePtrOutput) SlowRequests() SlowRequestsBasedTriggerResponsePtrOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) *SlowRequestsBasedTriggerResponse {
		if v == nil {
			return nil
		}
		return v.SlowRequests
	}).(SlowRequestsBasedTriggerResponsePtrOutput)
}

func (o AutoHealTriggersResponsePtrOutput) SlowRequestsWithPath() SlowRequestsBasedTriggerResponseArrayOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) []SlowRequestsBasedTriggerResponse {
		if v == nil {
			return nil
		}
		return v.SlowRequestsWithPath
	}).(SlowRequestsBasedTriggerResponseArrayOutput)
}

func (o AutoHealTriggersResponsePtrOutput) StatusCodes() StatusCodesBasedTriggerResponseArrayOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) []StatusCodesBasedTriggerResponse {
		if v == nil {
			return nil
		}
		return v.StatusCodes
	}).(StatusCodesBasedTriggerResponseArrayOutput)
}

func (o AutoHealTriggersResponsePtrOutput) StatusCodesRange() StatusCodesRangeBasedTriggerResponseArrayOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) []StatusCodesRangeBasedTriggerResponse {
		if v == nil {
			return nil
		}
		return v.StatusCodesRange
	}).(StatusCodesRangeBasedTriggerResponseArrayOutput)
}

type AzureActiveDirectory struct {
	Enabled           *bool                             `pulumi:"enabled"`
	IsAutoProvisioned *bool                             `pulumi:"isAutoProvisioned"`
	Login             *AzureActiveDirectoryLogin        `pulumi:"login"`
	Registration      *AzureActiveDirectoryRegistration `pulumi:"registration"`
	Validation        *AzureActiveDirectoryValidation   `pulumi:"validation"`
}





type AzureActiveDirectoryInput interface {
	pulumi.Input

	ToAzureActiveDirectoryOutput() AzureActiveDirectoryOutput
	ToAzureActiveDirectoryOutputWithContext(context.Context) AzureActiveDirectoryOutput
}

type AzureActiveDirectoryArgs struct {
	Enabled           pulumi.BoolPtrInput                      `pulumi:"enabled"`
	IsAutoProvisioned pulumi.BoolPtrInput                      `pulumi:"isAutoProvisioned"`
	Login             AzureActiveDirectoryLoginPtrInput        `pulumi:"login"`
	Registration      AzureActiveDirectoryRegistrationPtrInput `pulumi:"registration"`
	Validation        AzureActiveDirectoryValidationPtrInput   `pulumi:"validation"`
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

func (o AzureActiveDirectoryOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryOutput) IsAutoProvisioned() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *bool { return v.IsAutoProvisioned }).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryOutput) Login() AzureActiveDirectoryLoginPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *AzureActiveDirectoryLogin { return v.Login }).(AzureActiveDirectoryLoginPtrOutput)
}

func (o AzureActiveDirectoryOutput) Registration() AzureActiveDirectoryRegistrationPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *AzureActiveDirectoryRegistration { return v.Registration }).(AzureActiveDirectoryRegistrationPtrOutput)
}

func (o AzureActiveDirectoryOutput) Validation() AzureActiveDirectoryValidationPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectory) *AzureActiveDirectoryValidation { return v.Validation }).(AzureActiveDirectoryValidationPtrOutput)
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

func (o AzureActiveDirectoryPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryPtrOutput) IsAutoProvisioned() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *bool {
		if v == nil {
			return nil
		}
		return v.IsAutoProvisioned
	}).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryPtrOutput) Login() AzureActiveDirectoryLoginPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *AzureActiveDirectoryLogin {
		if v == nil {
			return nil
		}
		return v.Login
	}).(AzureActiveDirectoryLoginPtrOutput)
}

func (o AzureActiveDirectoryPtrOutput) Registration() AzureActiveDirectoryRegistrationPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *AzureActiveDirectoryRegistration {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(AzureActiveDirectoryRegistrationPtrOutput)
}

func (o AzureActiveDirectoryPtrOutput) Validation() AzureActiveDirectoryValidationPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectory) *AzureActiveDirectoryValidation {
		if v == nil {
			return nil
		}
		return v.Validation
	}).(AzureActiveDirectoryValidationPtrOutput)
}

type AzureActiveDirectoryLogin struct {
	DisableWWWAuthenticate *bool    `pulumi:"disableWWWAuthenticate"`
	LoginParameters        []string `pulumi:"loginParameters"`
}





type AzureActiveDirectoryLoginInput interface {
	pulumi.Input

	ToAzureActiveDirectoryLoginOutput() AzureActiveDirectoryLoginOutput
	ToAzureActiveDirectoryLoginOutputWithContext(context.Context) AzureActiveDirectoryLoginOutput
}

type AzureActiveDirectoryLoginArgs struct {
	DisableWWWAuthenticate pulumi.BoolPtrInput     `pulumi:"disableWWWAuthenticate"`
	LoginParameters        pulumi.StringArrayInput `pulumi:"loginParameters"`
}

func (AzureActiveDirectoryLoginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryLogin)(nil)).Elem()
}

func (i AzureActiveDirectoryLoginArgs) ToAzureActiveDirectoryLoginOutput() AzureActiveDirectoryLoginOutput {
	return i.ToAzureActiveDirectoryLoginOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryLoginArgs) ToAzureActiveDirectoryLoginOutputWithContext(ctx context.Context) AzureActiveDirectoryLoginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryLoginOutput)
}

func (i AzureActiveDirectoryLoginArgs) ToAzureActiveDirectoryLoginPtrOutput() AzureActiveDirectoryLoginPtrOutput {
	return i.ToAzureActiveDirectoryLoginPtrOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryLoginArgs) ToAzureActiveDirectoryLoginPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryLoginPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryLoginOutput).ToAzureActiveDirectoryLoginPtrOutputWithContext(ctx)
}









type AzureActiveDirectoryLoginPtrInput interface {
	pulumi.Input

	ToAzureActiveDirectoryLoginPtrOutput() AzureActiveDirectoryLoginPtrOutput
	ToAzureActiveDirectoryLoginPtrOutputWithContext(context.Context) AzureActiveDirectoryLoginPtrOutput
}

type azureActiveDirectoryLoginPtrType AzureActiveDirectoryLoginArgs

func AzureActiveDirectoryLoginPtr(v *AzureActiveDirectoryLoginArgs) AzureActiveDirectoryLoginPtrInput {
	return (*azureActiveDirectoryLoginPtrType)(v)
}

func (*azureActiveDirectoryLoginPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryLogin)(nil)).Elem()
}

func (i *azureActiveDirectoryLoginPtrType) ToAzureActiveDirectoryLoginPtrOutput() AzureActiveDirectoryLoginPtrOutput {
	return i.ToAzureActiveDirectoryLoginPtrOutputWithContext(context.Background())
}

func (i *azureActiveDirectoryLoginPtrType) ToAzureActiveDirectoryLoginPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryLoginPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryLoginPtrOutput)
}

type AzureActiveDirectoryLoginOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryLoginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryLogin)(nil)).Elem()
}

func (o AzureActiveDirectoryLoginOutput) ToAzureActiveDirectoryLoginOutput() AzureActiveDirectoryLoginOutput {
	return o
}

func (o AzureActiveDirectoryLoginOutput) ToAzureActiveDirectoryLoginOutputWithContext(ctx context.Context) AzureActiveDirectoryLoginOutput {
	return o
}

func (o AzureActiveDirectoryLoginOutput) ToAzureActiveDirectoryLoginPtrOutput() AzureActiveDirectoryLoginPtrOutput {
	return o.ToAzureActiveDirectoryLoginPtrOutputWithContext(context.Background())
}

func (o AzureActiveDirectoryLoginOutput) ToAzureActiveDirectoryLoginPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryLoginPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureActiveDirectoryLogin) *AzureActiveDirectoryLogin {
		return &v
	}).(AzureActiveDirectoryLoginPtrOutput)
}

func (o AzureActiveDirectoryLoginOutput) DisableWWWAuthenticate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryLogin) *bool { return v.DisableWWWAuthenticate }).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryLoginOutput) LoginParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureActiveDirectoryLogin) []string { return v.LoginParameters }).(pulumi.StringArrayOutput)
}

type AzureActiveDirectoryLoginPtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryLoginPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryLogin)(nil)).Elem()
}

func (o AzureActiveDirectoryLoginPtrOutput) ToAzureActiveDirectoryLoginPtrOutput() AzureActiveDirectoryLoginPtrOutput {
	return o
}

func (o AzureActiveDirectoryLoginPtrOutput) ToAzureActiveDirectoryLoginPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryLoginPtrOutput {
	return o
}

func (o AzureActiveDirectoryLoginPtrOutput) Elem() AzureActiveDirectoryLoginOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryLogin) AzureActiveDirectoryLogin {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectoryLogin
		return ret
	}).(AzureActiveDirectoryLoginOutput)
}

func (o AzureActiveDirectoryLoginPtrOutput) DisableWWWAuthenticate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryLogin) *bool {
		if v == nil {
			return nil
		}
		return v.DisableWWWAuthenticate
	}).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryLoginPtrOutput) LoginParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryLogin) []string {
		if v == nil {
			return nil
		}
		return v.LoginParameters
	}).(pulumi.StringArrayOutput)
}

type AzureActiveDirectoryLoginResponse struct {
	DisableWWWAuthenticate *bool    `pulumi:"disableWWWAuthenticate"`
	LoginParameters        []string `pulumi:"loginParameters"`
}

type AzureActiveDirectoryLoginResponseOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryLoginResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryLoginResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryLoginResponseOutput) ToAzureActiveDirectoryLoginResponseOutput() AzureActiveDirectoryLoginResponseOutput {
	return o
}

func (o AzureActiveDirectoryLoginResponseOutput) ToAzureActiveDirectoryLoginResponseOutputWithContext(ctx context.Context) AzureActiveDirectoryLoginResponseOutput {
	return o
}

func (o AzureActiveDirectoryLoginResponseOutput) DisableWWWAuthenticate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryLoginResponse) *bool { return v.DisableWWWAuthenticate }).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryLoginResponseOutput) LoginParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureActiveDirectoryLoginResponse) []string { return v.LoginParameters }).(pulumi.StringArrayOutput)
}

type AzureActiveDirectoryLoginResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryLoginResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryLoginResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryLoginResponsePtrOutput) ToAzureActiveDirectoryLoginResponsePtrOutput() AzureActiveDirectoryLoginResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryLoginResponsePtrOutput) ToAzureActiveDirectoryLoginResponsePtrOutputWithContext(ctx context.Context) AzureActiveDirectoryLoginResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryLoginResponsePtrOutput) Elem() AzureActiveDirectoryLoginResponseOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryLoginResponse) AzureActiveDirectoryLoginResponse {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectoryLoginResponse
		return ret
	}).(AzureActiveDirectoryLoginResponseOutput)
}

func (o AzureActiveDirectoryLoginResponsePtrOutput) DisableWWWAuthenticate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryLoginResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisableWWWAuthenticate
	}).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryLoginResponsePtrOutput) LoginParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryLoginResponse) []string {
		if v == nil {
			return nil
		}
		return v.LoginParameters
	}).(pulumi.StringArrayOutput)
}

type AzureActiveDirectoryRegistration struct {
	ClientId                                      *string `pulumi:"clientId"`
	ClientSecretCertificateIssuer                 *string `pulumi:"clientSecretCertificateIssuer"`
	ClientSecretCertificateSubjectAlternativeName *string `pulumi:"clientSecretCertificateSubjectAlternativeName"`
	ClientSecretCertificateThumbprint             *string `pulumi:"clientSecretCertificateThumbprint"`
	ClientSecretSettingName                       *string `pulumi:"clientSecretSettingName"`
	OpenIdIssuer                                  *string `pulumi:"openIdIssuer"`
}





type AzureActiveDirectoryRegistrationInput interface {
	pulumi.Input

	ToAzureActiveDirectoryRegistrationOutput() AzureActiveDirectoryRegistrationOutput
	ToAzureActiveDirectoryRegistrationOutputWithContext(context.Context) AzureActiveDirectoryRegistrationOutput
}

type AzureActiveDirectoryRegistrationArgs struct {
	ClientId                                      pulumi.StringPtrInput `pulumi:"clientId"`
	ClientSecretCertificateIssuer                 pulumi.StringPtrInput `pulumi:"clientSecretCertificateIssuer"`
	ClientSecretCertificateSubjectAlternativeName pulumi.StringPtrInput `pulumi:"clientSecretCertificateSubjectAlternativeName"`
	ClientSecretCertificateThumbprint             pulumi.StringPtrInput `pulumi:"clientSecretCertificateThumbprint"`
	ClientSecretSettingName                       pulumi.StringPtrInput `pulumi:"clientSecretSettingName"`
	OpenIdIssuer                                  pulumi.StringPtrInput `pulumi:"openIdIssuer"`
}

func (AzureActiveDirectoryRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryRegistration)(nil)).Elem()
}

func (i AzureActiveDirectoryRegistrationArgs) ToAzureActiveDirectoryRegistrationOutput() AzureActiveDirectoryRegistrationOutput {
	return i.ToAzureActiveDirectoryRegistrationOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryRegistrationArgs) ToAzureActiveDirectoryRegistrationOutputWithContext(ctx context.Context) AzureActiveDirectoryRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryRegistrationOutput)
}

func (i AzureActiveDirectoryRegistrationArgs) ToAzureActiveDirectoryRegistrationPtrOutput() AzureActiveDirectoryRegistrationPtrOutput {
	return i.ToAzureActiveDirectoryRegistrationPtrOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryRegistrationArgs) ToAzureActiveDirectoryRegistrationPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryRegistrationOutput).ToAzureActiveDirectoryRegistrationPtrOutputWithContext(ctx)
}









type AzureActiveDirectoryRegistrationPtrInput interface {
	pulumi.Input

	ToAzureActiveDirectoryRegistrationPtrOutput() AzureActiveDirectoryRegistrationPtrOutput
	ToAzureActiveDirectoryRegistrationPtrOutputWithContext(context.Context) AzureActiveDirectoryRegistrationPtrOutput
}

type azureActiveDirectoryRegistrationPtrType AzureActiveDirectoryRegistrationArgs

func AzureActiveDirectoryRegistrationPtr(v *AzureActiveDirectoryRegistrationArgs) AzureActiveDirectoryRegistrationPtrInput {
	return (*azureActiveDirectoryRegistrationPtrType)(v)
}

func (*azureActiveDirectoryRegistrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryRegistration)(nil)).Elem()
}

func (i *azureActiveDirectoryRegistrationPtrType) ToAzureActiveDirectoryRegistrationPtrOutput() AzureActiveDirectoryRegistrationPtrOutput {
	return i.ToAzureActiveDirectoryRegistrationPtrOutputWithContext(context.Background())
}

func (i *azureActiveDirectoryRegistrationPtrType) ToAzureActiveDirectoryRegistrationPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryRegistrationPtrOutput)
}

type AzureActiveDirectoryRegistrationOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryRegistration)(nil)).Elem()
}

func (o AzureActiveDirectoryRegistrationOutput) ToAzureActiveDirectoryRegistrationOutput() AzureActiveDirectoryRegistrationOutput {
	return o
}

func (o AzureActiveDirectoryRegistrationOutput) ToAzureActiveDirectoryRegistrationOutputWithContext(ctx context.Context) AzureActiveDirectoryRegistrationOutput {
	return o
}

func (o AzureActiveDirectoryRegistrationOutput) ToAzureActiveDirectoryRegistrationPtrOutput() AzureActiveDirectoryRegistrationPtrOutput {
	return o.ToAzureActiveDirectoryRegistrationPtrOutputWithContext(context.Background())
}

func (o AzureActiveDirectoryRegistrationOutput) ToAzureActiveDirectoryRegistrationPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryRegistrationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureActiveDirectoryRegistration) *AzureActiveDirectoryRegistration {
		return &v
	}).(AzureActiveDirectoryRegistrationPtrOutput)
}

func (o AzureActiveDirectoryRegistrationOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistration) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationOutput) ClientSecretCertificateIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistration) *string { return v.ClientSecretCertificateIssuer }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationOutput) ClientSecretCertificateSubjectAlternativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistration) *string {
		return v.ClientSecretCertificateSubjectAlternativeName
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationOutput) ClientSecretCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistration) *string { return v.ClientSecretCertificateThumbprint }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistration) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationOutput) OpenIdIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistration) *string { return v.OpenIdIssuer }).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryRegistrationPtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryRegistrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryRegistration)(nil)).Elem()
}

func (o AzureActiveDirectoryRegistrationPtrOutput) ToAzureActiveDirectoryRegistrationPtrOutput() AzureActiveDirectoryRegistrationPtrOutput {
	return o
}

func (o AzureActiveDirectoryRegistrationPtrOutput) ToAzureActiveDirectoryRegistrationPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryRegistrationPtrOutput {
	return o
}

func (o AzureActiveDirectoryRegistrationPtrOutput) Elem() AzureActiveDirectoryRegistrationOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistration) AzureActiveDirectoryRegistration {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectoryRegistration
		return ret
	}).(AzureActiveDirectoryRegistrationOutput)
}

func (o AzureActiveDirectoryRegistrationPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationPtrOutput) ClientSecretCertificateIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretCertificateIssuer
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationPtrOutput) ClientSecretCertificateSubjectAlternativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretCertificateSubjectAlternativeName
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationPtrOutput) ClientSecretCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretCertificateThumbprint
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationPtrOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretSettingName
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationPtrOutput) OpenIdIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistration) *string {
		if v == nil {
			return nil
		}
		return v.OpenIdIssuer
	}).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryRegistrationResponse struct {
	ClientId                                      *string `pulumi:"clientId"`
	ClientSecretCertificateIssuer                 *string `pulumi:"clientSecretCertificateIssuer"`
	ClientSecretCertificateSubjectAlternativeName *string `pulumi:"clientSecretCertificateSubjectAlternativeName"`
	ClientSecretCertificateThumbprint             *string `pulumi:"clientSecretCertificateThumbprint"`
	ClientSecretSettingName                       *string `pulumi:"clientSecretSettingName"`
	OpenIdIssuer                                  *string `pulumi:"openIdIssuer"`
}

type AzureActiveDirectoryRegistrationResponseOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryRegistrationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryRegistrationResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryRegistrationResponseOutput) ToAzureActiveDirectoryRegistrationResponseOutput() AzureActiveDirectoryRegistrationResponseOutput {
	return o
}

func (o AzureActiveDirectoryRegistrationResponseOutput) ToAzureActiveDirectoryRegistrationResponseOutputWithContext(ctx context.Context) AzureActiveDirectoryRegistrationResponseOutput {
	return o
}

func (o AzureActiveDirectoryRegistrationResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistrationResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponseOutput) ClientSecretCertificateIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistrationResponse) *string { return v.ClientSecretCertificateIssuer }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponseOutput) ClientSecretCertificateSubjectAlternativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistrationResponse) *string {
		return v.ClientSecretCertificateSubjectAlternativeName
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponseOutput) ClientSecretCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistrationResponse) *string { return v.ClientSecretCertificateThumbprint }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponseOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistrationResponse) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponseOutput) OpenIdIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryRegistrationResponse) *string { return v.OpenIdIssuer }).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryRegistrationResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryRegistrationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryRegistrationResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryRegistrationResponsePtrOutput) ToAzureActiveDirectoryRegistrationResponsePtrOutput() AzureActiveDirectoryRegistrationResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryRegistrationResponsePtrOutput) ToAzureActiveDirectoryRegistrationResponsePtrOutputWithContext(ctx context.Context) AzureActiveDirectoryRegistrationResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryRegistrationResponsePtrOutput) Elem() AzureActiveDirectoryRegistrationResponseOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistrationResponse) AzureActiveDirectoryRegistrationResponse {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectoryRegistrationResponse
		return ret
	}).(AzureActiveDirectoryRegistrationResponseOutput)
}

func (o AzureActiveDirectoryRegistrationResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponsePtrOutput) ClientSecretCertificateIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretCertificateIssuer
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponsePtrOutput) ClientSecretCertificateSubjectAlternativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretCertificateSubjectAlternativeName
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponsePtrOutput) ClientSecretCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretCertificateThumbprint
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponsePtrOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretSettingName
	}).(pulumi.StringPtrOutput)
}

func (o AzureActiveDirectoryRegistrationResponsePtrOutput) OpenIdIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.OpenIdIssuer
	}).(pulumi.StringPtrOutput)
}

type AzureActiveDirectoryResponse struct {
	Enabled           *bool                                     `pulumi:"enabled"`
	IsAutoProvisioned *bool                                     `pulumi:"isAutoProvisioned"`
	Login             *AzureActiveDirectoryLoginResponse        `pulumi:"login"`
	Registration      *AzureActiveDirectoryRegistrationResponse `pulumi:"registration"`
	Validation        *AzureActiveDirectoryValidationResponse   `pulumi:"validation"`
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

func (o AzureActiveDirectoryResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryResponseOutput) IsAutoProvisioned() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *bool { return v.IsAutoProvisioned }).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryResponseOutput) Login() AzureActiveDirectoryLoginResponsePtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *AzureActiveDirectoryLoginResponse { return v.Login }).(AzureActiveDirectoryLoginResponsePtrOutput)
}

func (o AzureActiveDirectoryResponseOutput) Registration() AzureActiveDirectoryRegistrationResponsePtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *AzureActiveDirectoryRegistrationResponse { return v.Registration }).(AzureActiveDirectoryRegistrationResponsePtrOutput)
}

func (o AzureActiveDirectoryResponseOutput) Validation() AzureActiveDirectoryValidationResponsePtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryResponse) *AzureActiveDirectoryValidationResponse { return v.Validation }).(AzureActiveDirectoryValidationResponsePtrOutput)
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

func (o AzureActiveDirectoryResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) IsAutoProvisioned() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsAutoProvisioned
	}).(pulumi.BoolPtrOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) Login() AzureActiveDirectoryLoginResponsePtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *AzureActiveDirectoryLoginResponse {
		if v == nil {
			return nil
		}
		return v.Login
	}).(AzureActiveDirectoryLoginResponsePtrOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) Registration() AzureActiveDirectoryRegistrationResponsePtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *AzureActiveDirectoryRegistrationResponse {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(AzureActiveDirectoryRegistrationResponsePtrOutput)
}

func (o AzureActiveDirectoryResponsePtrOutput) Validation() AzureActiveDirectoryValidationResponsePtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryResponse) *AzureActiveDirectoryValidationResponse {
		if v == nil {
			return nil
		}
		return v.Validation
	}).(AzureActiveDirectoryValidationResponsePtrOutput)
}

type AzureActiveDirectoryValidation struct {
	AllowedAudiences           []string                    `pulumi:"allowedAudiences"`
	DefaultAuthorizationPolicy *DefaultAuthorizationPolicy `pulumi:"defaultAuthorizationPolicy"`
	JwtClaimChecks             *JwtClaimChecks             `pulumi:"jwtClaimChecks"`
}





type AzureActiveDirectoryValidationInput interface {
	pulumi.Input

	ToAzureActiveDirectoryValidationOutput() AzureActiveDirectoryValidationOutput
	ToAzureActiveDirectoryValidationOutputWithContext(context.Context) AzureActiveDirectoryValidationOutput
}

type AzureActiveDirectoryValidationArgs struct {
	AllowedAudiences           pulumi.StringArrayInput            `pulumi:"allowedAudiences"`
	DefaultAuthorizationPolicy DefaultAuthorizationPolicyPtrInput `pulumi:"defaultAuthorizationPolicy"`
	JwtClaimChecks             JwtClaimChecksPtrInput             `pulumi:"jwtClaimChecks"`
}

func (AzureActiveDirectoryValidationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryValidation)(nil)).Elem()
}

func (i AzureActiveDirectoryValidationArgs) ToAzureActiveDirectoryValidationOutput() AzureActiveDirectoryValidationOutput {
	return i.ToAzureActiveDirectoryValidationOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryValidationArgs) ToAzureActiveDirectoryValidationOutputWithContext(ctx context.Context) AzureActiveDirectoryValidationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryValidationOutput)
}

func (i AzureActiveDirectoryValidationArgs) ToAzureActiveDirectoryValidationPtrOutput() AzureActiveDirectoryValidationPtrOutput {
	return i.ToAzureActiveDirectoryValidationPtrOutputWithContext(context.Background())
}

func (i AzureActiveDirectoryValidationArgs) ToAzureActiveDirectoryValidationPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryValidationOutput).ToAzureActiveDirectoryValidationPtrOutputWithContext(ctx)
}









type AzureActiveDirectoryValidationPtrInput interface {
	pulumi.Input

	ToAzureActiveDirectoryValidationPtrOutput() AzureActiveDirectoryValidationPtrOutput
	ToAzureActiveDirectoryValidationPtrOutputWithContext(context.Context) AzureActiveDirectoryValidationPtrOutput
}

type azureActiveDirectoryValidationPtrType AzureActiveDirectoryValidationArgs

func AzureActiveDirectoryValidationPtr(v *AzureActiveDirectoryValidationArgs) AzureActiveDirectoryValidationPtrInput {
	return (*azureActiveDirectoryValidationPtrType)(v)
}

func (*azureActiveDirectoryValidationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryValidation)(nil)).Elem()
}

func (i *azureActiveDirectoryValidationPtrType) ToAzureActiveDirectoryValidationPtrOutput() AzureActiveDirectoryValidationPtrOutput {
	return i.ToAzureActiveDirectoryValidationPtrOutputWithContext(context.Background())
}

func (i *azureActiveDirectoryValidationPtrType) ToAzureActiveDirectoryValidationPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureActiveDirectoryValidationPtrOutput)
}

type AzureActiveDirectoryValidationOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryValidationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryValidation)(nil)).Elem()
}

func (o AzureActiveDirectoryValidationOutput) ToAzureActiveDirectoryValidationOutput() AzureActiveDirectoryValidationOutput {
	return o
}

func (o AzureActiveDirectoryValidationOutput) ToAzureActiveDirectoryValidationOutputWithContext(ctx context.Context) AzureActiveDirectoryValidationOutput {
	return o
}

func (o AzureActiveDirectoryValidationOutput) ToAzureActiveDirectoryValidationPtrOutput() AzureActiveDirectoryValidationPtrOutput {
	return o.ToAzureActiveDirectoryValidationPtrOutputWithContext(context.Background())
}

func (o AzureActiveDirectoryValidationOutput) ToAzureActiveDirectoryValidationPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryValidationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureActiveDirectoryValidation) *AzureActiveDirectoryValidation {
		return &v
	}).(AzureActiveDirectoryValidationPtrOutput)
}

func (o AzureActiveDirectoryValidationOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureActiveDirectoryValidation) []string { return v.AllowedAudiences }).(pulumi.StringArrayOutput)
}

func (o AzureActiveDirectoryValidationOutput) DefaultAuthorizationPolicy() DefaultAuthorizationPolicyPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryValidation) *DefaultAuthorizationPolicy {
		return v.DefaultAuthorizationPolicy
	}).(DefaultAuthorizationPolicyPtrOutput)
}

func (o AzureActiveDirectoryValidationOutput) JwtClaimChecks() JwtClaimChecksPtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryValidation) *JwtClaimChecks { return v.JwtClaimChecks }).(JwtClaimChecksPtrOutput)
}

type AzureActiveDirectoryValidationPtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryValidationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryValidation)(nil)).Elem()
}

func (o AzureActiveDirectoryValidationPtrOutput) ToAzureActiveDirectoryValidationPtrOutput() AzureActiveDirectoryValidationPtrOutput {
	return o
}

func (o AzureActiveDirectoryValidationPtrOutput) ToAzureActiveDirectoryValidationPtrOutputWithContext(ctx context.Context) AzureActiveDirectoryValidationPtrOutput {
	return o
}

func (o AzureActiveDirectoryValidationPtrOutput) Elem() AzureActiveDirectoryValidationOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryValidation) AzureActiveDirectoryValidation {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectoryValidation
		return ret
	}).(AzureActiveDirectoryValidationOutput)
}

func (o AzureActiveDirectoryValidationPtrOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryValidation) []string {
		if v == nil {
			return nil
		}
		return v.AllowedAudiences
	}).(pulumi.StringArrayOutput)
}

func (o AzureActiveDirectoryValidationPtrOutput) DefaultAuthorizationPolicy() DefaultAuthorizationPolicyPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryValidation) *DefaultAuthorizationPolicy {
		if v == nil {
			return nil
		}
		return v.DefaultAuthorizationPolicy
	}).(DefaultAuthorizationPolicyPtrOutput)
}

func (o AzureActiveDirectoryValidationPtrOutput) JwtClaimChecks() JwtClaimChecksPtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryValidation) *JwtClaimChecks {
		if v == nil {
			return nil
		}
		return v.JwtClaimChecks
	}).(JwtClaimChecksPtrOutput)
}

type AzureActiveDirectoryValidationResponse struct {
	AllowedAudiences           []string                            `pulumi:"allowedAudiences"`
	DefaultAuthorizationPolicy *DefaultAuthorizationPolicyResponse `pulumi:"defaultAuthorizationPolicy"`
	JwtClaimChecks             *JwtClaimChecksResponse             `pulumi:"jwtClaimChecks"`
}

type AzureActiveDirectoryValidationResponseOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryValidationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureActiveDirectoryValidationResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryValidationResponseOutput) ToAzureActiveDirectoryValidationResponseOutput() AzureActiveDirectoryValidationResponseOutput {
	return o
}

func (o AzureActiveDirectoryValidationResponseOutput) ToAzureActiveDirectoryValidationResponseOutputWithContext(ctx context.Context) AzureActiveDirectoryValidationResponseOutput {
	return o
}

func (o AzureActiveDirectoryValidationResponseOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureActiveDirectoryValidationResponse) []string { return v.AllowedAudiences }).(pulumi.StringArrayOutput)
}

func (o AzureActiveDirectoryValidationResponseOutput) DefaultAuthorizationPolicy() DefaultAuthorizationPolicyResponsePtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryValidationResponse) *DefaultAuthorizationPolicyResponse {
		return v.DefaultAuthorizationPolicy
	}).(DefaultAuthorizationPolicyResponsePtrOutput)
}

func (o AzureActiveDirectoryValidationResponseOutput) JwtClaimChecks() JwtClaimChecksResponsePtrOutput {
	return o.ApplyT(func(v AzureActiveDirectoryValidationResponse) *JwtClaimChecksResponse { return v.JwtClaimChecks }).(JwtClaimChecksResponsePtrOutput)
}

type AzureActiveDirectoryValidationResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureActiveDirectoryValidationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureActiveDirectoryValidationResponse)(nil)).Elem()
}

func (o AzureActiveDirectoryValidationResponsePtrOutput) ToAzureActiveDirectoryValidationResponsePtrOutput() AzureActiveDirectoryValidationResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryValidationResponsePtrOutput) ToAzureActiveDirectoryValidationResponsePtrOutputWithContext(ctx context.Context) AzureActiveDirectoryValidationResponsePtrOutput {
	return o
}

func (o AzureActiveDirectoryValidationResponsePtrOutput) Elem() AzureActiveDirectoryValidationResponseOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryValidationResponse) AzureActiveDirectoryValidationResponse {
		if v != nil {
			return *v
		}
		var ret AzureActiveDirectoryValidationResponse
		return ret
	}).(AzureActiveDirectoryValidationResponseOutput)
}

func (o AzureActiveDirectoryValidationResponsePtrOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryValidationResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedAudiences
	}).(pulumi.StringArrayOutput)
}

func (o AzureActiveDirectoryValidationResponsePtrOutput) DefaultAuthorizationPolicy() DefaultAuthorizationPolicyResponsePtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryValidationResponse) *DefaultAuthorizationPolicyResponse {
		if v == nil {
			return nil
		}
		return v.DefaultAuthorizationPolicy
	}).(DefaultAuthorizationPolicyResponsePtrOutput)
}

func (o AzureActiveDirectoryValidationResponsePtrOutput) JwtClaimChecks() JwtClaimChecksResponsePtrOutput {
	return o.ApplyT(func(v *AzureActiveDirectoryValidationResponse) *JwtClaimChecksResponse {
		if v == nil {
			return nil
		}
		return v.JwtClaimChecks
	}).(JwtClaimChecksResponsePtrOutput)
}

type AzureBlobStorageApplicationLogsConfig struct {
	Level           *LogLevel `pulumi:"level"`
	RetentionInDays *int      `pulumi:"retentionInDays"`
	SasUrl          *string   `pulumi:"sasUrl"`
}





type AzureBlobStorageApplicationLogsConfigInput interface {
	pulumi.Input

	ToAzureBlobStorageApplicationLogsConfigOutput() AzureBlobStorageApplicationLogsConfigOutput
	ToAzureBlobStorageApplicationLogsConfigOutputWithContext(context.Context) AzureBlobStorageApplicationLogsConfigOutput
}

type AzureBlobStorageApplicationLogsConfigArgs struct {
	Level           LogLevelPtrInput      `pulumi:"level"`
	RetentionInDays pulumi.IntPtrInput    `pulumi:"retentionInDays"`
	SasUrl          pulumi.StringPtrInput `pulumi:"sasUrl"`
}

func (AzureBlobStorageApplicationLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageApplicationLogsConfig)(nil)).Elem()
}

func (i AzureBlobStorageApplicationLogsConfigArgs) ToAzureBlobStorageApplicationLogsConfigOutput() AzureBlobStorageApplicationLogsConfigOutput {
	return i.ToAzureBlobStorageApplicationLogsConfigOutputWithContext(context.Background())
}

func (i AzureBlobStorageApplicationLogsConfigArgs) ToAzureBlobStorageApplicationLogsConfigOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageApplicationLogsConfigOutput)
}

func (i AzureBlobStorageApplicationLogsConfigArgs) ToAzureBlobStorageApplicationLogsConfigPtrOutput() AzureBlobStorageApplicationLogsConfigPtrOutput {
	return i.ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i AzureBlobStorageApplicationLogsConfigArgs) ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageApplicationLogsConfigOutput).ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(ctx)
}









type AzureBlobStorageApplicationLogsConfigPtrInput interface {
	pulumi.Input

	ToAzureBlobStorageApplicationLogsConfigPtrOutput() AzureBlobStorageApplicationLogsConfigPtrOutput
	ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(context.Context) AzureBlobStorageApplicationLogsConfigPtrOutput
}

type azureBlobStorageApplicationLogsConfigPtrType AzureBlobStorageApplicationLogsConfigArgs

func AzureBlobStorageApplicationLogsConfigPtr(v *AzureBlobStorageApplicationLogsConfigArgs) AzureBlobStorageApplicationLogsConfigPtrInput {
	return (*azureBlobStorageApplicationLogsConfigPtrType)(v)
}

func (*azureBlobStorageApplicationLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageApplicationLogsConfig)(nil)).Elem()
}

func (i *azureBlobStorageApplicationLogsConfigPtrType) ToAzureBlobStorageApplicationLogsConfigPtrOutput() AzureBlobStorageApplicationLogsConfigPtrOutput {
	return i.ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i *azureBlobStorageApplicationLogsConfigPtrType) ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageApplicationLogsConfigPtrOutput)
}

type AzureBlobStorageApplicationLogsConfigOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageApplicationLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageApplicationLogsConfig)(nil)).Elem()
}

func (o AzureBlobStorageApplicationLogsConfigOutput) ToAzureBlobStorageApplicationLogsConfigOutput() AzureBlobStorageApplicationLogsConfigOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigOutput) ToAzureBlobStorageApplicationLogsConfigOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigOutput) ToAzureBlobStorageApplicationLogsConfigPtrOutput() AzureBlobStorageApplicationLogsConfigPtrOutput {
	return o.ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (o AzureBlobStorageApplicationLogsConfigOutput) ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureBlobStorageApplicationLogsConfig) *AzureBlobStorageApplicationLogsConfig {
		return &v
	}).(AzureBlobStorageApplicationLogsConfigPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigOutput) Level() LogLevelPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageApplicationLogsConfig) *LogLevel { return v.Level }).(LogLevelPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageApplicationLogsConfig) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageApplicationLogsConfig) *string { return v.SasUrl }).(pulumi.StringPtrOutput)
}

type AzureBlobStorageApplicationLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageApplicationLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageApplicationLogsConfig)(nil)).Elem()
}

func (o AzureBlobStorageApplicationLogsConfigPtrOutput) ToAzureBlobStorageApplicationLogsConfigPtrOutput() AzureBlobStorageApplicationLogsConfigPtrOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigPtrOutput) ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigPtrOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigPtrOutput) Elem() AzureBlobStorageApplicationLogsConfigOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfig) AzureBlobStorageApplicationLogsConfig {
		if v != nil {
			return *v
		}
		var ret AzureBlobStorageApplicationLogsConfig
		return ret
	}).(AzureBlobStorageApplicationLogsConfigOutput)
}

func (o AzureBlobStorageApplicationLogsConfigPtrOutput) Level() LogLevelPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfig) *LogLevel {
		if v == nil {
			return nil
		}
		return v.Level
	}).(LogLevelPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigPtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfig) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigPtrOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfig) *string {
		if v == nil {
			return nil
		}
		return v.SasUrl
	}).(pulumi.StringPtrOutput)
}

type AzureBlobStorageApplicationLogsConfigResponse struct {
	Level           *string `pulumi:"level"`
	RetentionInDays *int    `pulumi:"retentionInDays"`
	SasUrl          *string `pulumi:"sasUrl"`
}

type AzureBlobStorageApplicationLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageApplicationLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageApplicationLogsConfigResponse)(nil)).Elem()
}

func (o AzureBlobStorageApplicationLogsConfigResponseOutput) ToAzureBlobStorageApplicationLogsConfigResponseOutput() AzureBlobStorageApplicationLogsConfigResponseOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigResponseOutput) ToAzureBlobStorageApplicationLogsConfigResponseOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigResponseOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageApplicationLogsConfigResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigResponseOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageApplicationLogsConfigResponse) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigResponseOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageApplicationLogsConfigResponse) *string { return v.SasUrl }).(pulumi.StringPtrOutput)
}

type AzureBlobStorageApplicationLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageApplicationLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageApplicationLogsConfigResponse)(nil)).Elem()
}

func (o AzureBlobStorageApplicationLogsConfigResponsePtrOutput) ToAzureBlobStorageApplicationLogsConfigResponsePtrOutput() AzureBlobStorageApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigResponsePtrOutput) ToAzureBlobStorageApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigResponsePtrOutput) Elem() AzureBlobStorageApplicationLogsConfigResponseOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfigResponse) AzureBlobStorageApplicationLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret AzureBlobStorageApplicationLogsConfigResponse
		return ret
	}).(AzureBlobStorageApplicationLogsConfigResponseOutput)
}

func (o AzureBlobStorageApplicationLogsConfigResponsePtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Level
	}).(pulumi.StringPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigResponsePtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigResponsePtrOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.SasUrl
	}).(pulumi.StringPtrOutput)
}

type AzureBlobStorageHttpLogsConfig struct {
	Enabled         *bool   `pulumi:"enabled"`
	RetentionInDays *int    `pulumi:"retentionInDays"`
	SasUrl          *string `pulumi:"sasUrl"`
}





type AzureBlobStorageHttpLogsConfigInput interface {
	pulumi.Input

	ToAzureBlobStorageHttpLogsConfigOutput() AzureBlobStorageHttpLogsConfigOutput
	ToAzureBlobStorageHttpLogsConfigOutputWithContext(context.Context) AzureBlobStorageHttpLogsConfigOutput
}

type AzureBlobStorageHttpLogsConfigArgs struct {
	Enabled         pulumi.BoolPtrInput   `pulumi:"enabled"`
	RetentionInDays pulumi.IntPtrInput    `pulumi:"retentionInDays"`
	SasUrl          pulumi.StringPtrInput `pulumi:"sasUrl"`
}

func (AzureBlobStorageHttpLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageHttpLogsConfig)(nil)).Elem()
}

func (i AzureBlobStorageHttpLogsConfigArgs) ToAzureBlobStorageHttpLogsConfigOutput() AzureBlobStorageHttpLogsConfigOutput {
	return i.ToAzureBlobStorageHttpLogsConfigOutputWithContext(context.Background())
}

func (i AzureBlobStorageHttpLogsConfigArgs) ToAzureBlobStorageHttpLogsConfigOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageHttpLogsConfigOutput)
}

func (i AzureBlobStorageHttpLogsConfigArgs) ToAzureBlobStorageHttpLogsConfigPtrOutput() AzureBlobStorageHttpLogsConfigPtrOutput {
	return i.ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (i AzureBlobStorageHttpLogsConfigArgs) ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageHttpLogsConfigOutput).ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(ctx)
}









type AzureBlobStorageHttpLogsConfigPtrInput interface {
	pulumi.Input

	ToAzureBlobStorageHttpLogsConfigPtrOutput() AzureBlobStorageHttpLogsConfigPtrOutput
	ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(context.Context) AzureBlobStorageHttpLogsConfigPtrOutput
}

type azureBlobStorageHttpLogsConfigPtrType AzureBlobStorageHttpLogsConfigArgs

func AzureBlobStorageHttpLogsConfigPtr(v *AzureBlobStorageHttpLogsConfigArgs) AzureBlobStorageHttpLogsConfigPtrInput {
	return (*azureBlobStorageHttpLogsConfigPtrType)(v)
}

func (*azureBlobStorageHttpLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageHttpLogsConfig)(nil)).Elem()
}

func (i *azureBlobStorageHttpLogsConfigPtrType) ToAzureBlobStorageHttpLogsConfigPtrOutput() AzureBlobStorageHttpLogsConfigPtrOutput {
	return i.ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (i *azureBlobStorageHttpLogsConfigPtrType) ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageHttpLogsConfigPtrOutput)
}

type AzureBlobStorageHttpLogsConfigOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageHttpLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageHttpLogsConfig)(nil)).Elem()
}

func (o AzureBlobStorageHttpLogsConfigOutput) ToAzureBlobStorageHttpLogsConfigOutput() AzureBlobStorageHttpLogsConfigOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigOutput) ToAzureBlobStorageHttpLogsConfigOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigOutput) ToAzureBlobStorageHttpLogsConfigPtrOutput() AzureBlobStorageHttpLogsConfigPtrOutput {
	return o.ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (o AzureBlobStorageHttpLogsConfigOutput) ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureBlobStorageHttpLogsConfig) *AzureBlobStorageHttpLogsConfig {
		return &v
	}).(AzureBlobStorageHttpLogsConfigPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageHttpLogsConfig) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageHttpLogsConfig) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageHttpLogsConfig) *string { return v.SasUrl }).(pulumi.StringPtrOutput)
}

type AzureBlobStorageHttpLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageHttpLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageHttpLogsConfig)(nil)).Elem()
}

func (o AzureBlobStorageHttpLogsConfigPtrOutput) ToAzureBlobStorageHttpLogsConfigPtrOutput() AzureBlobStorageHttpLogsConfigPtrOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigPtrOutput) ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigPtrOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigPtrOutput) Elem() AzureBlobStorageHttpLogsConfigOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfig) AzureBlobStorageHttpLogsConfig {
		if v != nil {
			return *v
		}
		var ret AzureBlobStorageHttpLogsConfig
		return ret
	}).(AzureBlobStorageHttpLogsConfigOutput)
}

func (o AzureBlobStorageHttpLogsConfigPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfig) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigPtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfig) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigPtrOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfig) *string {
		if v == nil {
			return nil
		}
		return v.SasUrl
	}).(pulumi.StringPtrOutput)
}

type AzureBlobStorageHttpLogsConfigResponse struct {
	Enabled         *bool   `pulumi:"enabled"`
	RetentionInDays *int    `pulumi:"retentionInDays"`
	SasUrl          *string `pulumi:"sasUrl"`
}

type AzureBlobStorageHttpLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageHttpLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageHttpLogsConfigResponse)(nil)).Elem()
}

func (o AzureBlobStorageHttpLogsConfigResponseOutput) ToAzureBlobStorageHttpLogsConfigResponseOutput() AzureBlobStorageHttpLogsConfigResponseOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigResponseOutput) ToAzureBlobStorageHttpLogsConfigResponseOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigResponseOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageHttpLogsConfigResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigResponseOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageHttpLogsConfigResponse) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigResponseOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageHttpLogsConfigResponse) *string { return v.SasUrl }).(pulumi.StringPtrOutput)
}

type AzureBlobStorageHttpLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageHttpLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageHttpLogsConfigResponse)(nil)).Elem()
}

func (o AzureBlobStorageHttpLogsConfigResponsePtrOutput) ToAzureBlobStorageHttpLogsConfigResponsePtrOutput() AzureBlobStorageHttpLogsConfigResponsePtrOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigResponsePtrOutput) ToAzureBlobStorageHttpLogsConfigResponsePtrOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigResponsePtrOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigResponsePtrOutput) Elem() AzureBlobStorageHttpLogsConfigResponseOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfigResponse) AzureBlobStorageHttpLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret AzureBlobStorageHttpLogsConfigResponse
		return ret
	}).(AzureBlobStorageHttpLogsConfigResponseOutput)
}

func (o AzureBlobStorageHttpLogsConfigResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigResponsePtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigResponsePtrOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.SasUrl
	}).(pulumi.StringPtrOutput)
}

type AzureStaticWebApps struct {
	Enabled      *bool                           `pulumi:"enabled"`
	Registration *AzureStaticWebAppsRegistration `pulumi:"registration"`
}





type AzureStaticWebAppsInput interface {
	pulumi.Input

	ToAzureStaticWebAppsOutput() AzureStaticWebAppsOutput
	ToAzureStaticWebAppsOutputWithContext(context.Context) AzureStaticWebAppsOutput
}

type AzureStaticWebAppsArgs struct {
	Enabled      pulumi.BoolPtrInput                    `pulumi:"enabled"`
	Registration AzureStaticWebAppsRegistrationPtrInput `pulumi:"registration"`
}

func (AzureStaticWebAppsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStaticWebApps)(nil)).Elem()
}

func (i AzureStaticWebAppsArgs) ToAzureStaticWebAppsOutput() AzureStaticWebAppsOutput {
	return i.ToAzureStaticWebAppsOutputWithContext(context.Background())
}

func (i AzureStaticWebAppsArgs) ToAzureStaticWebAppsOutputWithContext(ctx context.Context) AzureStaticWebAppsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStaticWebAppsOutput)
}

func (i AzureStaticWebAppsArgs) ToAzureStaticWebAppsPtrOutput() AzureStaticWebAppsPtrOutput {
	return i.ToAzureStaticWebAppsPtrOutputWithContext(context.Background())
}

func (i AzureStaticWebAppsArgs) ToAzureStaticWebAppsPtrOutputWithContext(ctx context.Context) AzureStaticWebAppsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStaticWebAppsOutput).ToAzureStaticWebAppsPtrOutputWithContext(ctx)
}









type AzureStaticWebAppsPtrInput interface {
	pulumi.Input

	ToAzureStaticWebAppsPtrOutput() AzureStaticWebAppsPtrOutput
	ToAzureStaticWebAppsPtrOutputWithContext(context.Context) AzureStaticWebAppsPtrOutput
}

type azureStaticWebAppsPtrType AzureStaticWebAppsArgs

func AzureStaticWebAppsPtr(v *AzureStaticWebAppsArgs) AzureStaticWebAppsPtrInput {
	return (*azureStaticWebAppsPtrType)(v)
}

func (*azureStaticWebAppsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureStaticWebApps)(nil)).Elem()
}

func (i *azureStaticWebAppsPtrType) ToAzureStaticWebAppsPtrOutput() AzureStaticWebAppsPtrOutput {
	return i.ToAzureStaticWebAppsPtrOutputWithContext(context.Background())
}

func (i *azureStaticWebAppsPtrType) ToAzureStaticWebAppsPtrOutputWithContext(ctx context.Context) AzureStaticWebAppsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStaticWebAppsPtrOutput)
}

type AzureStaticWebAppsOutput struct{ *pulumi.OutputState }

func (AzureStaticWebAppsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStaticWebApps)(nil)).Elem()
}

func (o AzureStaticWebAppsOutput) ToAzureStaticWebAppsOutput() AzureStaticWebAppsOutput {
	return o
}

func (o AzureStaticWebAppsOutput) ToAzureStaticWebAppsOutputWithContext(ctx context.Context) AzureStaticWebAppsOutput {
	return o
}

func (o AzureStaticWebAppsOutput) ToAzureStaticWebAppsPtrOutput() AzureStaticWebAppsPtrOutput {
	return o.ToAzureStaticWebAppsPtrOutputWithContext(context.Background())
}

func (o AzureStaticWebAppsOutput) ToAzureStaticWebAppsPtrOutputWithContext(ctx context.Context) AzureStaticWebAppsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureStaticWebApps) *AzureStaticWebApps {
		return &v
	}).(AzureStaticWebAppsPtrOutput)
}

func (o AzureStaticWebAppsOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureStaticWebApps) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AzureStaticWebAppsOutput) Registration() AzureStaticWebAppsRegistrationPtrOutput {
	return o.ApplyT(func(v AzureStaticWebApps) *AzureStaticWebAppsRegistration { return v.Registration }).(AzureStaticWebAppsRegistrationPtrOutput)
}

type AzureStaticWebAppsPtrOutput struct{ *pulumi.OutputState }

func (AzureStaticWebAppsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureStaticWebApps)(nil)).Elem()
}

func (o AzureStaticWebAppsPtrOutput) ToAzureStaticWebAppsPtrOutput() AzureStaticWebAppsPtrOutput {
	return o
}

func (o AzureStaticWebAppsPtrOutput) ToAzureStaticWebAppsPtrOutputWithContext(ctx context.Context) AzureStaticWebAppsPtrOutput {
	return o
}

func (o AzureStaticWebAppsPtrOutput) Elem() AzureStaticWebAppsOutput {
	return o.ApplyT(func(v *AzureStaticWebApps) AzureStaticWebApps {
		if v != nil {
			return *v
		}
		var ret AzureStaticWebApps
		return ret
	}).(AzureStaticWebAppsOutput)
}

func (o AzureStaticWebAppsPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureStaticWebApps) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AzureStaticWebAppsPtrOutput) Registration() AzureStaticWebAppsRegistrationPtrOutput {
	return o.ApplyT(func(v *AzureStaticWebApps) *AzureStaticWebAppsRegistration {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(AzureStaticWebAppsRegistrationPtrOutput)
}

type AzureStaticWebAppsRegistration struct {
	ClientId *string `pulumi:"clientId"`
}





type AzureStaticWebAppsRegistrationInput interface {
	pulumi.Input

	ToAzureStaticWebAppsRegistrationOutput() AzureStaticWebAppsRegistrationOutput
	ToAzureStaticWebAppsRegistrationOutputWithContext(context.Context) AzureStaticWebAppsRegistrationOutput
}

type AzureStaticWebAppsRegistrationArgs struct {
	ClientId pulumi.StringPtrInput `pulumi:"clientId"`
}

func (AzureStaticWebAppsRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStaticWebAppsRegistration)(nil)).Elem()
}

func (i AzureStaticWebAppsRegistrationArgs) ToAzureStaticWebAppsRegistrationOutput() AzureStaticWebAppsRegistrationOutput {
	return i.ToAzureStaticWebAppsRegistrationOutputWithContext(context.Background())
}

func (i AzureStaticWebAppsRegistrationArgs) ToAzureStaticWebAppsRegistrationOutputWithContext(ctx context.Context) AzureStaticWebAppsRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStaticWebAppsRegistrationOutput)
}

func (i AzureStaticWebAppsRegistrationArgs) ToAzureStaticWebAppsRegistrationPtrOutput() AzureStaticWebAppsRegistrationPtrOutput {
	return i.ToAzureStaticWebAppsRegistrationPtrOutputWithContext(context.Background())
}

func (i AzureStaticWebAppsRegistrationArgs) ToAzureStaticWebAppsRegistrationPtrOutputWithContext(ctx context.Context) AzureStaticWebAppsRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStaticWebAppsRegistrationOutput).ToAzureStaticWebAppsRegistrationPtrOutputWithContext(ctx)
}









type AzureStaticWebAppsRegistrationPtrInput interface {
	pulumi.Input

	ToAzureStaticWebAppsRegistrationPtrOutput() AzureStaticWebAppsRegistrationPtrOutput
	ToAzureStaticWebAppsRegistrationPtrOutputWithContext(context.Context) AzureStaticWebAppsRegistrationPtrOutput
}

type azureStaticWebAppsRegistrationPtrType AzureStaticWebAppsRegistrationArgs

func AzureStaticWebAppsRegistrationPtr(v *AzureStaticWebAppsRegistrationArgs) AzureStaticWebAppsRegistrationPtrInput {
	return (*azureStaticWebAppsRegistrationPtrType)(v)
}

func (*azureStaticWebAppsRegistrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureStaticWebAppsRegistration)(nil)).Elem()
}

func (i *azureStaticWebAppsRegistrationPtrType) ToAzureStaticWebAppsRegistrationPtrOutput() AzureStaticWebAppsRegistrationPtrOutput {
	return i.ToAzureStaticWebAppsRegistrationPtrOutputWithContext(context.Background())
}

func (i *azureStaticWebAppsRegistrationPtrType) ToAzureStaticWebAppsRegistrationPtrOutputWithContext(ctx context.Context) AzureStaticWebAppsRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStaticWebAppsRegistrationPtrOutput)
}

type AzureStaticWebAppsRegistrationOutput struct{ *pulumi.OutputState }

func (AzureStaticWebAppsRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStaticWebAppsRegistration)(nil)).Elem()
}

func (o AzureStaticWebAppsRegistrationOutput) ToAzureStaticWebAppsRegistrationOutput() AzureStaticWebAppsRegistrationOutput {
	return o
}

func (o AzureStaticWebAppsRegistrationOutput) ToAzureStaticWebAppsRegistrationOutputWithContext(ctx context.Context) AzureStaticWebAppsRegistrationOutput {
	return o
}

func (o AzureStaticWebAppsRegistrationOutput) ToAzureStaticWebAppsRegistrationPtrOutput() AzureStaticWebAppsRegistrationPtrOutput {
	return o.ToAzureStaticWebAppsRegistrationPtrOutputWithContext(context.Background())
}

func (o AzureStaticWebAppsRegistrationOutput) ToAzureStaticWebAppsRegistrationPtrOutputWithContext(ctx context.Context) AzureStaticWebAppsRegistrationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureStaticWebAppsRegistration) *AzureStaticWebAppsRegistration {
		return &v
	}).(AzureStaticWebAppsRegistrationPtrOutput)
}

func (o AzureStaticWebAppsRegistrationOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStaticWebAppsRegistration) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

type AzureStaticWebAppsRegistrationPtrOutput struct{ *pulumi.OutputState }

func (AzureStaticWebAppsRegistrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureStaticWebAppsRegistration)(nil)).Elem()
}

func (o AzureStaticWebAppsRegistrationPtrOutput) ToAzureStaticWebAppsRegistrationPtrOutput() AzureStaticWebAppsRegistrationPtrOutput {
	return o
}

func (o AzureStaticWebAppsRegistrationPtrOutput) ToAzureStaticWebAppsRegistrationPtrOutputWithContext(ctx context.Context) AzureStaticWebAppsRegistrationPtrOutput {
	return o
}

func (o AzureStaticWebAppsRegistrationPtrOutput) Elem() AzureStaticWebAppsRegistrationOutput {
	return o.ApplyT(func(v *AzureStaticWebAppsRegistration) AzureStaticWebAppsRegistration {
		if v != nil {
			return *v
		}
		var ret AzureStaticWebAppsRegistration
		return ret
	}).(AzureStaticWebAppsRegistrationOutput)
}

func (o AzureStaticWebAppsRegistrationPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureStaticWebAppsRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

type AzureStaticWebAppsRegistrationResponse struct {
	ClientId *string `pulumi:"clientId"`
}

type AzureStaticWebAppsRegistrationResponseOutput struct{ *pulumi.OutputState }

func (AzureStaticWebAppsRegistrationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStaticWebAppsRegistrationResponse)(nil)).Elem()
}

func (o AzureStaticWebAppsRegistrationResponseOutput) ToAzureStaticWebAppsRegistrationResponseOutput() AzureStaticWebAppsRegistrationResponseOutput {
	return o
}

func (o AzureStaticWebAppsRegistrationResponseOutput) ToAzureStaticWebAppsRegistrationResponseOutputWithContext(ctx context.Context) AzureStaticWebAppsRegistrationResponseOutput {
	return o
}

func (o AzureStaticWebAppsRegistrationResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStaticWebAppsRegistrationResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

type AzureStaticWebAppsRegistrationResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureStaticWebAppsRegistrationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureStaticWebAppsRegistrationResponse)(nil)).Elem()
}

func (o AzureStaticWebAppsRegistrationResponsePtrOutput) ToAzureStaticWebAppsRegistrationResponsePtrOutput() AzureStaticWebAppsRegistrationResponsePtrOutput {
	return o
}

func (o AzureStaticWebAppsRegistrationResponsePtrOutput) ToAzureStaticWebAppsRegistrationResponsePtrOutputWithContext(ctx context.Context) AzureStaticWebAppsRegistrationResponsePtrOutput {
	return o
}

func (o AzureStaticWebAppsRegistrationResponsePtrOutput) Elem() AzureStaticWebAppsRegistrationResponseOutput {
	return o.ApplyT(func(v *AzureStaticWebAppsRegistrationResponse) AzureStaticWebAppsRegistrationResponse {
		if v != nil {
			return *v
		}
		var ret AzureStaticWebAppsRegistrationResponse
		return ret
	}).(AzureStaticWebAppsRegistrationResponseOutput)
}

func (o AzureStaticWebAppsRegistrationResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureStaticWebAppsRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

type AzureStaticWebAppsResponse struct {
	Enabled      *bool                                   `pulumi:"enabled"`
	Registration *AzureStaticWebAppsRegistrationResponse `pulumi:"registration"`
}

type AzureStaticWebAppsResponseOutput struct{ *pulumi.OutputState }

func (AzureStaticWebAppsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStaticWebAppsResponse)(nil)).Elem()
}

func (o AzureStaticWebAppsResponseOutput) ToAzureStaticWebAppsResponseOutput() AzureStaticWebAppsResponseOutput {
	return o
}

func (o AzureStaticWebAppsResponseOutput) ToAzureStaticWebAppsResponseOutputWithContext(ctx context.Context) AzureStaticWebAppsResponseOutput {
	return o
}

func (o AzureStaticWebAppsResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureStaticWebAppsResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AzureStaticWebAppsResponseOutput) Registration() AzureStaticWebAppsRegistrationResponsePtrOutput {
	return o.ApplyT(func(v AzureStaticWebAppsResponse) *AzureStaticWebAppsRegistrationResponse { return v.Registration }).(AzureStaticWebAppsRegistrationResponsePtrOutput)
}

type AzureStaticWebAppsResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureStaticWebAppsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureStaticWebAppsResponse)(nil)).Elem()
}

func (o AzureStaticWebAppsResponsePtrOutput) ToAzureStaticWebAppsResponsePtrOutput() AzureStaticWebAppsResponsePtrOutput {
	return o
}

func (o AzureStaticWebAppsResponsePtrOutput) ToAzureStaticWebAppsResponsePtrOutputWithContext(ctx context.Context) AzureStaticWebAppsResponsePtrOutput {
	return o
}

func (o AzureStaticWebAppsResponsePtrOutput) Elem() AzureStaticWebAppsResponseOutput {
	return o.ApplyT(func(v *AzureStaticWebAppsResponse) AzureStaticWebAppsResponse {
		if v != nil {
			return *v
		}
		var ret AzureStaticWebAppsResponse
		return ret
	}).(AzureStaticWebAppsResponseOutput)
}

func (o AzureStaticWebAppsResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureStaticWebAppsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AzureStaticWebAppsResponsePtrOutput) Registration() AzureStaticWebAppsRegistrationResponsePtrOutput {
	return o.ApplyT(func(v *AzureStaticWebAppsResponse) *AzureStaticWebAppsRegistrationResponse {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(AzureStaticWebAppsRegistrationResponsePtrOutput)
}

type AzureStorageInfoValue struct {
	AccessKey   *string           `pulumi:"accessKey"`
	AccountName *string           `pulumi:"accountName"`
	MountPath   *string           `pulumi:"mountPath"`
	ShareName   *string           `pulumi:"shareName"`
	Type        *AzureStorageType `pulumi:"type"`
}





type AzureStorageInfoValueInput interface {
	pulumi.Input

	ToAzureStorageInfoValueOutput() AzureStorageInfoValueOutput
	ToAzureStorageInfoValueOutputWithContext(context.Context) AzureStorageInfoValueOutput
}

type AzureStorageInfoValueArgs struct {
	AccessKey   pulumi.StringPtrInput    `pulumi:"accessKey"`
	AccountName pulumi.StringPtrInput    `pulumi:"accountName"`
	MountPath   pulumi.StringPtrInput    `pulumi:"mountPath"`
	ShareName   pulumi.StringPtrInput    `pulumi:"shareName"`
	Type        AzureStorageTypePtrInput `pulumi:"type"`
}

func (AzureStorageInfoValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStorageInfoValue)(nil)).Elem()
}

func (i AzureStorageInfoValueArgs) ToAzureStorageInfoValueOutput() AzureStorageInfoValueOutput {
	return i.ToAzureStorageInfoValueOutputWithContext(context.Background())
}

func (i AzureStorageInfoValueArgs) ToAzureStorageInfoValueOutputWithContext(ctx context.Context) AzureStorageInfoValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStorageInfoValueOutput)
}





type AzureStorageInfoValueMapInput interface {
	pulumi.Input

	ToAzureStorageInfoValueMapOutput() AzureStorageInfoValueMapOutput
	ToAzureStorageInfoValueMapOutputWithContext(context.Context) AzureStorageInfoValueMapOutput
}

type AzureStorageInfoValueMap map[string]AzureStorageInfoValueInput

func (AzureStorageInfoValueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AzureStorageInfoValue)(nil)).Elem()
}

func (i AzureStorageInfoValueMap) ToAzureStorageInfoValueMapOutput() AzureStorageInfoValueMapOutput {
	return i.ToAzureStorageInfoValueMapOutputWithContext(context.Background())
}

func (i AzureStorageInfoValueMap) ToAzureStorageInfoValueMapOutputWithContext(ctx context.Context) AzureStorageInfoValueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStorageInfoValueMapOutput)
}

type AzureStorageInfoValueOutput struct{ *pulumi.OutputState }

func (AzureStorageInfoValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStorageInfoValue)(nil)).Elem()
}

func (o AzureStorageInfoValueOutput) ToAzureStorageInfoValueOutput() AzureStorageInfoValueOutput {
	return o
}

func (o AzureStorageInfoValueOutput) ToAzureStorageInfoValueOutputWithContext(ctx context.Context) AzureStorageInfoValueOutput {
	return o
}

func (o AzureStorageInfoValueOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValue) *string { return v.AccessKey }).(pulumi.StringPtrOutput)
}

func (o AzureStorageInfoValueOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValue) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o AzureStorageInfoValueOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValue) *string { return v.MountPath }).(pulumi.StringPtrOutput)
}

func (o AzureStorageInfoValueOutput) ShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValue) *string { return v.ShareName }).(pulumi.StringPtrOutput)
}

func (o AzureStorageInfoValueOutput) Type() AzureStorageTypePtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValue) *AzureStorageType { return v.Type }).(AzureStorageTypePtrOutput)
}

type AzureStorageInfoValueMapOutput struct{ *pulumi.OutputState }

func (AzureStorageInfoValueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AzureStorageInfoValue)(nil)).Elem()
}

func (o AzureStorageInfoValueMapOutput) ToAzureStorageInfoValueMapOutput() AzureStorageInfoValueMapOutput {
	return o
}

func (o AzureStorageInfoValueMapOutput) ToAzureStorageInfoValueMapOutputWithContext(ctx context.Context) AzureStorageInfoValueMapOutput {
	return o
}

func (o AzureStorageInfoValueMapOutput) MapIndex(k pulumi.StringInput) AzureStorageInfoValueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AzureStorageInfoValue {
		return vs[0].(map[string]AzureStorageInfoValue)[vs[1].(string)]
	}).(AzureStorageInfoValueOutput)
}

type AzureStorageInfoValueResponse struct {
	AccessKey   *string `pulumi:"accessKey"`
	AccountName *string `pulumi:"accountName"`
	MountPath   *string `pulumi:"mountPath"`
	ShareName   *string `pulumi:"shareName"`
	State       string  `pulumi:"state"`
	Type        *string `pulumi:"type"`
}

type AzureStorageInfoValueResponseOutput struct{ *pulumi.OutputState }

func (AzureStorageInfoValueResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStorageInfoValueResponse)(nil)).Elem()
}

func (o AzureStorageInfoValueResponseOutput) ToAzureStorageInfoValueResponseOutput() AzureStorageInfoValueResponseOutput {
	return o
}

func (o AzureStorageInfoValueResponseOutput) ToAzureStorageInfoValueResponseOutputWithContext(ctx context.Context) AzureStorageInfoValueResponseOutput {
	return o
}

func (o AzureStorageInfoValueResponseOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValueResponse) *string { return v.AccessKey }).(pulumi.StringPtrOutput)
}

func (o AzureStorageInfoValueResponseOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValueResponse) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o AzureStorageInfoValueResponseOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValueResponse) *string { return v.MountPath }).(pulumi.StringPtrOutput)
}

func (o AzureStorageInfoValueResponseOutput) ShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValueResponse) *string { return v.ShareName }).(pulumi.StringPtrOutput)
}

func (o AzureStorageInfoValueResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v AzureStorageInfoValueResponse) string { return v.State }).(pulumi.StringOutput)
}

func (o AzureStorageInfoValueResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValueResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type AzureStorageInfoValueResponseMapOutput struct{ *pulumi.OutputState }

func (AzureStorageInfoValueResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AzureStorageInfoValueResponse)(nil)).Elem()
}

func (o AzureStorageInfoValueResponseMapOutput) ToAzureStorageInfoValueResponseMapOutput() AzureStorageInfoValueResponseMapOutput {
	return o
}

func (o AzureStorageInfoValueResponseMapOutput) ToAzureStorageInfoValueResponseMapOutputWithContext(ctx context.Context) AzureStorageInfoValueResponseMapOutput {
	return o
}

func (o AzureStorageInfoValueResponseMapOutput) MapIndex(k pulumi.StringInput) AzureStorageInfoValueResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AzureStorageInfoValueResponse {
		return vs[0].(map[string]AzureStorageInfoValueResponse)[vs[1].(string)]
	}).(AzureStorageInfoValueResponseOutput)
}

type AzureTableStorageApplicationLogsConfig struct {
	Level  *LogLevel `pulumi:"level"`
	SasUrl string    `pulumi:"sasUrl"`
}





type AzureTableStorageApplicationLogsConfigInput interface {
	pulumi.Input

	ToAzureTableStorageApplicationLogsConfigOutput() AzureTableStorageApplicationLogsConfigOutput
	ToAzureTableStorageApplicationLogsConfigOutputWithContext(context.Context) AzureTableStorageApplicationLogsConfigOutput
}

type AzureTableStorageApplicationLogsConfigArgs struct {
	Level  LogLevelPtrInput   `pulumi:"level"`
	SasUrl pulumi.StringInput `pulumi:"sasUrl"`
}

func (AzureTableStorageApplicationLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureTableStorageApplicationLogsConfig)(nil)).Elem()
}

func (i AzureTableStorageApplicationLogsConfigArgs) ToAzureTableStorageApplicationLogsConfigOutput() AzureTableStorageApplicationLogsConfigOutput {
	return i.ToAzureTableStorageApplicationLogsConfigOutputWithContext(context.Background())
}

func (i AzureTableStorageApplicationLogsConfigArgs) ToAzureTableStorageApplicationLogsConfigOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureTableStorageApplicationLogsConfigOutput)
}

func (i AzureTableStorageApplicationLogsConfigArgs) ToAzureTableStorageApplicationLogsConfigPtrOutput() AzureTableStorageApplicationLogsConfigPtrOutput {
	return i.ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i AzureTableStorageApplicationLogsConfigArgs) ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureTableStorageApplicationLogsConfigOutput).ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(ctx)
}









type AzureTableStorageApplicationLogsConfigPtrInput interface {
	pulumi.Input

	ToAzureTableStorageApplicationLogsConfigPtrOutput() AzureTableStorageApplicationLogsConfigPtrOutput
	ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(context.Context) AzureTableStorageApplicationLogsConfigPtrOutput
}

type azureTableStorageApplicationLogsConfigPtrType AzureTableStorageApplicationLogsConfigArgs

func AzureTableStorageApplicationLogsConfigPtr(v *AzureTableStorageApplicationLogsConfigArgs) AzureTableStorageApplicationLogsConfigPtrInput {
	return (*azureTableStorageApplicationLogsConfigPtrType)(v)
}

func (*azureTableStorageApplicationLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureTableStorageApplicationLogsConfig)(nil)).Elem()
}

func (i *azureTableStorageApplicationLogsConfigPtrType) ToAzureTableStorageApplicationLogsConfigPtrOutput() AzureTableStorageApplicationLogsConfigPtrOutput {
	return i.ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i *azureTableStorageApplicationLogsConfigPtrType) ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureTableStorageApplicationLogsConfigPtrOutput)
}

type AzureTableStorageApplicationLogsConfigOutput struct{ *pulumi.OutputState }

func (AzureTableStorageApplicationLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureTableStorageApplicationLogsConfig)(nil)).Elem()
}

func (o AzureTableStorageApplicationLogsConfigOutput) ToAzureTableStorageApplicationLogsConfigOutput() AzureTableStorageApplicationLogsConfigOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigOutput) ToAzureTableStorageApplicationLogsConfigOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigOutput) ToAzureTableStorageApplicationLogsConfigPtrOutput() AzureTableStorageApplicationLogsConfigPtrOutput {
	return o.ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (o AzureTableStorageApplicationLogsConfigOutput) ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureTableStorageApplicationLogsConfig) *AzureTableStorageApplicationLogsConfig {
		return &v
	}).(AzureTableStorageApplicationLogsConfigPtrOutput)
}

func (o AzureTableStorageApplicationLogsConfigOutput) Level() LogLevelPtrOutput {
	return o.ApplyT(func(v AzureTableStorageApplicationLogsConfig) *LogLevel { return v.Level }).(LogLevelPtrOutput)
}

func (o AzureTableStorageApplicationLogsConfigOutput) SasUrl() pulumi.StringOutput {
	return o.ApplyT(func(v AzureTableStorageApplicationLogsConfig) string { return v.SasUrl }).(pulumi.StringOutput)
}

type AzureTableStorageApplicationLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (AzureTableStorageApplicationLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureTableStorageApplicationLogsConfig)(nil)).Elem()
}

func (o AzureTableStorageApplicationLogsConfigPtrOutput) ToAzureTableStorageApplicationLogsConfigPtrOutput() AzureTableStorageApplicationLogsConfigPtrOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigPtrOutput) ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigPtrOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigPtrOutput) Elem() AzureTableStorageApplicationLogsConfigOutput {
	return o.ApplyT(func(v *AzureTableStorageApplicationLogsConfig) AzureTableStorageApplicationLogsConfig {
		if v != nil {
			return *v
		}
		var ret AzureTableStorageApplicationLogsConfig
		return ret
	}).(AzureTableStorageApplicationLogsConfigOutput)
}

func (o AzureTableStorageApplicationLogsConfigPtrOutput) Level() LogLevelPtrOutput {
	return o.ApplyT(func(v *AzureTableStorageApplicationLogsConfig) *LogLevel {
		if v == nil {
			return nil
		}
		return v.Level
	}).(LogLevelPtrOutput)
}

func (o AzureTableStorageApplicationLogsConfigPtrOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureTableStorageApplicationLogsConfig) *string {
		if v == nil {
			return nil
		}
		return &v.SasUrl
	}).(pulumi.StringPtrOutput)
}

type AzureTableStorageApplicationLogsConfigResponse struct {
	Level  *string `pulumi:"level"`
	SasUrl string  `pulumi:"sasUrl"`
}

type AzureTableStorageApplicationLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (AzureTableStorageApplicationLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureTableStorageApplicationLogsConfigResponse)(nil)).Elem()
}

func (o AzureTableStorageApplicationLogsConfigResponseOutput) ToAzureTableStorageApplicationLogsConfigResponseOutput() AzureTableStorageApplicationLogsConfigResponseOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigResponseOutput) ToAzureTableStorageApplicationLogsConfigResponseOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigResponseOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureTableStorageApplicationLogsConfigResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o AzureTableStorageApplicationLogsConfigResponseOutput) SasUrl() pulumi.StringOutput {
	return o.ApplyT(func(v AzureTableStorageApplicationLogsConfigResponse) string { return v.SasUrl }).(pulumi.StringOutput)
}

type AzureTableStorageApplicationLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureTableStorageApplicationLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureTableStorageApplicationLogsConfigResponse)(nil)).Elem()
}

func (o AzureTableStorageApplicationLogsConfigResponsePtrOutput) ToAzureTableStorageApplicationLogsConfigResponsePtrOutput() AzureTableStorageApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigResponsePtrOutput) ToAzureTableStorageApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigResponsePtrOutput) Elem() AzureTableStorageApplicationLogsConfigResponseOutput {
	return o.ApplyT(func(v *AzureTableStorageApplicationLogsConfigResponse) AzureTableStorageApplicationLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret AzureTableStorageApplicationLogsConfigResponse
		return ret
	}).(AzureTableStorageApplicationLogsConfigResponseOutput)
}

func (o AzureTableStorageApplicationLogsConfigResponsePtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureTableStorageApplicationLogsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Level
	}).(pulumi.StringPtrOutput)
}

func (o AzureTableStorageApplicationLogsConfigResponsePtrOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureTableStorageApplicationLogsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SasUrl
	}).(pulumi.StringPtrOutput)
}

type BackupItemResponse struct {
	BackupId             int                             `pulumi:"backupId"`
	BlobName             string                          `pulumi:"blobName"`
	CorrelationId        string                          `pulumi:"correlationId"`
	Created              string                          `pulumi:"created"`
	Databases            []DatabaseBackupSettingResponse `pulumi:"databases"`
	FinishedTimeStamp    string                          `pulumi:"finishedTimeStamp"`
	Id                   string                          `pulumi:"id"`
	Kind                 *string                         `pulumi:"kind"`
	LastRestoreTimeStamp string                          `pulumi:"lastRestoreTimeStamp"`
	Log                  string                          `pulumi:"log"`
	Name                 string                          `pulumi:"name"`
	Scheduled            bool                            `pulumi:"scheduled"`
	SizeInBytes          float64                         `pulumi:"sizeInBytes"`
	Status               string                          `pulumi:"status"`
	StorageAccountUrl    string                          `pulumi:"storageAccountUrl"`
	Type                 string                          `pulumi:"type"`
	WebsiteSizeInBytes   float64                         `pulumi:"websiteSizeInBytes"`
}

type BackupItemResponseOutput struct{ *pulumi.OutputState }

func (BackupItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupItemResponse)(nil)).Elem()
}

func (o BackupItemResponseOutput) ToBackupItemResponseOutput() BackupItemResponseOutput {
	return o
}

func (o BackupItemResponseOutput) ToBackupItemResponseOutputWithContext(ctx context.Context) BackupItemResponseOutput {
	return o
}

func (o BackupItemResponseOutput) BackupId() pulumi.IntOutput {
	return o.ApplyT(func(v BackupItemResponse) int { return v.BackupId }).(pulumi.IntOutput)
}

func (o BackupItemResponseOutput) BlobName() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.BlobName }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) CorrelationId() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.CorrelationId }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.Created }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) Databases() DatabaseBackupSettingResponseArrayOutput {
	return o.ApplyT(func(v BackupItemResponse) []DatabaseBackupSettingResponse { return v.Databases }).(DatabaseBackupSettingResponseArrayOutput)
}

func (o BackupItemResponseOutput) FinishedTimeStamp() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.FinishedTimeStamp }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupItemResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o BackupItemResponseOutput) LastRestoreTimeStamp() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.LastRestoreTimeStamp }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) Log() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.Log }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) Scheduled() pulumi.BoolOutput {
	return o.ApplyT(func(v BackupItemResponse) bool { return v.Scheduled }).(pulumi.BoolOutput)
}

func (o BackupItemResponseOutput) SizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v BackupItemResponse) float64 { return v.SizeInBytes }).(pulumi.Float64Output)
}

func (o BackupItemResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) StorageAccountUrl() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.StorageAccountUrl }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) WebsiteSizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v BackupItemResponse) float64 { return v.WebsiteSizeInBytes }).(pulumi.Float64Output)
}

type BackupItemResponseArrayOutput struct{ *pulumi.OutputState }

func (BackupItemResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackupItemResponse)(nil)).Elem()
}

func (o BackupItemResponseArrayOutput) ToBackupItemResponseArrayOutput() BackupItemResponseArrayOutput {
	return o
}

func (o BackupItemResponseArrayOutput) ToBackupItemResponseArrayOutputWithContext(ctx context.Context) BackupItemResponseArrayOutput {
	return o
}

func (o BackupItemResponseArrayOutput) Index(i pulumi.IntInput) BackupItemResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackupItemResponse {
		return vs[0].([]BackupItemResponse)[vs[1].(int)]
	}).(BackupItemResponseOutput)
}

type BackupSchedule struct {
	FrequencyInterval     int           `pulumi:"frequencyInterval"`
	FrequencyUnit         FrequencyUnit `pulumi:"frequencyUnit"`
	KeepAtLeastOneBackup  bool          `pulumi:"keepAtLeastOneBackup"`
	RetentionPeriodInDays int           `pulumi:"retentionPeriodInDays"`
	StartTime             *string       `pulumi:"startTime"`
}


func (val *BackupSchedule) Defaults() *BackupSchedule {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.FrequencyInterval) {
		tmp.FrequencyInterval = 7
	}
	if isZero(tmp.FrequencyUnit) {
		tmp.FrequencyUnit = FrequencyUnit("Day")
	}
	if isZero(tmp.KeepAtLeastOneBackup) {
		tmp.KeepAtLeastOneBackup = true
	}
	if isZero(tmp.RetentionPeriodInDays) {
		tmp.RetentionPeriodInDays = 30
	}
	return &tmp
}





type BackupScheduleInput interface {
	pulumi.Input

	ToBackupScheduleOutput() BackupScheduleOutput
	ToBackupScheduleOutputWithContext(context.Context) BackupScheduleOutput
}

type BackupScheduleArgs struct {
	FrequencyInterval     pulumi.IntInput       `pulumi:"frequencyInterval"`
	FrequencyUnit         FrequencyUnitInput    `pulumi:"frequencyUnit"`
	KeepAtLeastOneBackup  pulumi.BoolInput      `pulumi:"keepAtLeastOneBackup"`
	RetentionPeriodInDays pulumi.IntInput       `pulumi:"retentionPeriodInDays"`
	StartTime             pulumi.StringPtrInput `pulumi:"startTime"`
}


func (val *BackupScheduleArgs) Defaults() *BackupScheduleArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.FrequencyInterval) {
		tmp.FrequencyInterval = pulumi.Int(7)
	}
	if isZero(tmp.FrequencyUnit) {
		tmp.FrequencyUnit = FrequencyUnit("Day")
	}
	if isZero(tmp.KeepAtLeastOneBackup) {
		tmp.KeepAtLeastOneBackup = pulumi.Bool(true)
	}
	if isZero(tmp.RetentionPeriodInDays) {
		tmp.RetentionPeriodInDays = pulumi.Int(30)
	}
	return &tmp
}
func (BackupScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupSchedule)(nil)).Elem()
}

func (i BackupScheduleArgs) ToBackupScheduleOutput() BackupScheduleOutput {
	return i.ToBackupScheduleOutputWithContext(context.Background())
}

func (i BackupScheduleArgs) ToBackupScheduleOutputWithContext(ctx context.Context) BackupScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupScheduleOutput)
}

func (i BackupScheduleArgs) ToBackupSchedulePtrOutput() BackupSchedulePtrOutput {
	return i.ToBackupSchedulePtrOutputWithContext(context.Background())
}

func (i BackupScheduleArgs) ToBackupSchedulePtrOutputWithContext(ctx context.Context) BackupSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupScheduleOutput).ToBackupSchedulePtrOutputWithContext(ctx)
}









type BackupSchedulePtrInput interface {
	pulumi.Input

	ToBackupSchedulePtrOutput() BackupSchedulePtrOutput
	ToBackupSchedulePtrOutputWithContext(context.Context) BackupSchedulePtrOutput
}

type backupSchedulePtrType BackupScheduleArgs

func BackupSchedulePtr(v *BackupScheduleArgs) BackupSchedulePtrInput {
	return (*backupSchedulePtrType)(v)
}

func (*backupSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupSchedule)(nil)).Elem()
}

func (i *backupSchedulePtrType) ToBackupSchedulePtrOutput() BackupSchedulePtrOutput {
	return i.ToBackupSchedulePtrOutputWithContext(context.Background())
}

func (i *backupSchedulePtrType) ToBackupSchedulePtrOutputWithContext(ctx context.Context) BackupSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupSchedulePtrOutput)
}

type BackupScheduleOutput struct{ *pulumi.OutputState }

func (BackupScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupSchedule)(nil)).Elem()
}

func (o BackupScheduleOutput) ToBackupScheduleOutput() BackupScheduleOutput {
	return o
}

func (o BackupScheduleOutput) ToBackupScheduleOutputWithContext(ctx context.Context) BackupScheduleOutput {
	return o
}

func (o BackupScheduleOutput) ToBackupSchedulePtrOutput() BackupSchedulePtrOutput {
	return o.ToBackupSchedulePtrOutputWithContext(context.Background())
}

func (o BackupScheduleOutput) ToBackupSchedulePtrOutputWithContext(ctx context.Context) BackupSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupSchedule) *BackupSchedule {
		return &v
	}).(BackupSchedulePtrOutput)
}

func (o BackupScheduleOutput) FrequencyInterval() pulumi.IntOutput {
	return o.ApplyT(func(v BackupSchedule) int { return v.FrequencyInterval }).(pulumi.IntOutput)
}

func (o BackupScheduleOutput) FrequencyUnit() FrequencyUnitOutput {
	return o.ApplyT(func(v BackupSchedule) FrequencyUnit { return v.FrequencyUnit }).(FrequencyUnitOutput)
}

func (o BackupScheduleOutput) KeepAtLeastOneBackup() pulumi.BoolOutput {
	return o.ApplyT(func(v BackupSchedule) bool { return v.KeepAtLeastOneBackup }).(pulumi.BoolOutput)
}

func (o BackupScheduleOutput) RetentionPeriodInDays() pulumi.IntOutput {
	return o.ApplyT(func(v BackupSchedule) int { return v.RetentionPeriodInDays }).(pulumi.IntOutput)
}

func (o BackupScheduleOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupSchedule) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type BackupSchedulePtrOutput struct{ *pulumi.OutputState }

func (BackupSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupSchedule)(nil)).Elem()
}

func (o BackupSchedulePtrOutput) ToBackupSchedulePtrOutput() BackupSchedulePtrOutput {
	return o
}

func (o BackupSchedulePtrOutput) ToBackupSchedulePtrOutputWithContext(ctx context.Context) BackupSchedulePtrOutput {
	return o
}

func (o BackupSchedulePtrOutput) Elem() BackupScheduleOutput {
	return o.ApplyT(func(v *BackupSchedule) BackupSchedule {
		if v != nil {
			return *v
		}
		var ret BackupSchedule
		return ret
	}).(BackupScheduleOutput)
}

func (o BackupSchedulePtrOutput) FrequencyInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupSchedule) *int {
		if v == nil {
			return nil
		}
		return &v.FrequencyInterval
	}).(pulumi.IntPtrOutput)
}

func (o BackupSchedulePtrOutput) FrequencyUnit() FrequencyUnitPtrOutput {
	return o.ApplyT(func(v *BackupSchedule) *FrequencyUnit {
		if v == nil {
			return nil
		}
		return &v.FrequencyUnit
	}).(FrequencyUnitPtrOutput)
}

func (o BackupSchedulePtrOutput) KeepAtLeastOneBackup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackupSchedule) *bool {
		if v == nil {
			return nil
		}
		return &v.KeepAtLeastOneBackup
	}).(pulumi.BoolPtrOutput)
}

func (o BackupSchedulePtrOutput) RetentionPeriodInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupSchedule) *int {
		if v == nil {
			return nil
		}
		return &v.RetentionPeriodInDays
	}).(pulumi.IntPtrOutput)
}

func (o BackupSchedulePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupSchedule) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

type BackupScheduleResponse struct {
	FrequencyInterval     int     `pulumi:"frequencyInterval"`
	FrequencyUnit         string  `pulumi:"frequencyUnit"`
	KeepAtLeastOneBackup  bool    `pulumi:"keepAtLeastOneBackup"`
	LastExecutionTime     string  `pulumi:"lastExecutionTime"`
	RetentionPeriodInDays int     `pulumi:"retentionPeriodInDays"`
	StartTime             *string `pulumi:"startTime"`
}


func (val *BackupScheduleResponse) Defaults() *BackupScheduleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.FrequencyInterval) {
		tmp.FrequencyInterval = 7
	}
	if isZero(tmp.FrequencyUnit) {
		tmp.FrequencyUnit = "Day"
	}
	if isZero(tmp.KeepAtLeastOneBackup) {
		tmp.KeepAtLeastOneBackup = true
	}
	if isZero(tmp.RetentionPeriodInDays) {
		tmp.RetentionPeriodInDays = 30
	}
	return &tmp
}

type BackupScheduleResponseOutput struct{ *pulumi.OutputState }

func (BackupScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupScheduleResponse)(nil)).Elem()
}

func (o BackupScheduleResponseOutput) ToBackupScheduleResponseOutput() BackupScheduleResponseOutput {
	return o
}

func (o BackupScheduleResponseOutput) ToBackupScheduleResponseOutputWithContext(ctx context.Context) BackupScheduleResponseOutput {
	return o
}

func (o BackupScheduleResponseOutput) FrequencyInterval() pulumi.IntOutput {
	return o.ApplyT(func(v BackupScheduleResponse) int { return v.FrequencyInterval }).(pulumi.IntOutput)
}

func (o BackupScheduleResponseOutput) FrequencyUnit() pulumi.StringOutput {
	return o.ApplyT(func(v BackupScheduleResponse) string { return v.FrequencyUnit }).(pulumi.StringOutput)
}

func (o BackupScheduleResponseOutput) KeepAtLeastOneBackup() pulumi.BoolOutput {
	return o.ApplyT(func(v BackupScheduleResponse) bool { return v.KeepAtLeastOneBackup }).(pulumi.BoolOutput)
}

func (o BackupScheduleResponseOutput) LastExecutionTime() pulumi.StringOutput {
	return o.ApplyT(func(v BackupScheduleResponse) string { return v.LastExecutionTime }).(pulumi.StringOutput)
}

func (o BackupScheduleResponseOutput) RetentionPeriodInDays() pulumi.IntOutput {
	return o.ApplyT(func(v BackupScheduleResponse) int { return v.RetentionPeriodInDays }).(pulumi.IntOutput)
}

func (o BackupScheduleResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupScheduleResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type BackupScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (BackupScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupScheduleResponse)(nil)).Elem()
}

func (o BackupScheduleResponsePtrOutput) ToBackupScheduleResponsePtrOutput() BackupScheduleResponsePtrOutput {
	return o
}

func (o BackupScheduleResponsePtrOutput) ToBackupScheduleResponsePtrOutputWithContext(ctx context.Context) BackupScheduleResponsePtrOutput {
	return o
}

func (o BackupScheduleResponsePtrOutput) Elem() BackupScheduleResponseOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) BackupScheduleResponse {
		if v != nil {
			return *v
		}
		var ret BackupScheduleResponse
		return ret
	}).(BackupScheduleResponseOutput)
}

func (o BackupScheduleResponsePtrOutput) FrequencyInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.FrequencyInterval
	}).(pulumi.IntPtrOutput)
}

func (o BackupScheduleResponsePtrOutput) FrequencyUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FrequencyUnit
	}).(pulumi.StringPtrOutput)
}

func (o BackupScheduleResponsePtrOutput) KeepAtLeastOneBackup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.KeepAtLeastOneBackup
	}).(pulumi.BoolPtrOutput)
}

func (o BackupScheduleResponsePtrOutput) LastExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastExecutionTime
	}).(pulumi.StringPtrOutput)
}

func (o BackupScheduleResponsePtrOutput) RetentionPeriodInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.RetentionPeriodInDays
	}).(pulumi.IntPtrOutput)
}

func (o BackupScheduleResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

type BlobStorageTokenStore struct {
	SasUrlSettingName *string `pulumi:"sasUrlSettingName"`
}





type BlobStorageTokenStoreInput interface {
	pulumi.Input

	ToBlobStorageTokenStoreOutput() BlobStorageTokenStoreOutput
	ToBlobStorageTokenStoreOutputWithContext(context.Context) BlobStorageTokenStoreOutput
}

type BlobStorageTokenStoreArgs struct {
	SasUrlSettingName pulumi.StringPtrInput `pulumi:"sasUrlSettingName"`
}

func (BlobStorageTokenStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobStorageTokenStore)(nil)).Elem()
}

func (i BlobStorageTokenStoreArgs) ToBlobStorageTokenStoreOutput() BlobStorageTokenStoreOutput {
	return i.ToBlobStorageTokenStoreOutputWithContext(context.Background())
}

func (i BlobStorageTokenStoreArgs) ToBlobStorageTokenStoreOutputWithContext(ctx context.Context) BlobStorageTokenStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobStorageTokenStoreOutput)
}

func (i BlobStorageTokenStoreArgs) ToBlobStorageTokenStorePtrOutput() BlobStorageTokenStorePtrOutput {
	return i.ToBlobStorageTokenStorePtrOutputWithContext(context.Background())
}

func (i BlobStorageTokenStoreArgs) ToBlobStorageTokenStorePtrOutputWithContext(ctx context.Context) BlobStorageTokenStorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobStorageTokenStoreOutput).ToBlobStorageTokenStorePtrOutputWithContext(ctx)
}









type BlobStorageTokenStorePtrInput interface {
	pulumi.Input

	ToBlobStorageTokenStorePtrOutput() BlobStorageTokenStorePtrOutput
	ToBlobStorageTokenStorePtrOutputWithContext(context.Context) BlobStorageTokenStorePtrOutput
}

type blobStorageTokenStorePtrType BlobStorageTokenStoreArgs

func BlobStorageTokenStorePtr(v *BlobStorageTokenStoreArgs) BlobStorageTokenStorePtrInput {
	return (*blobStorageTokenStorePtrType)(v)
}

func (*blobStorageTokenStorePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobStorageTokenStore)(nil)).Elem()
}

func (i *blobStorageTokenStorePtrType) ToBlobStorageTokenStorePtrOutput() BlobStorageTokenStorePtrOutput {
	return i.ToBlobStorageTokenStorePtrOutputWithContext(context.Background())
}

func (i *blobStorageTokenStorePtrType) ToBlobStorageTokenStorePtrOutputWithContext(ctx context.Context) BlobStorageTokenStorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobStorageTokenStorePtrOutput)
}

type BlobStorageTokenStoreOutput struct{ *pulumi.OutputState }

func (BlobStorageTokenStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobStorageTokenStore)(nil)).Elem()
}

func (o BlobStorageTokenStoreOutput) ToBlobStorageTokenStoreOutput() BlobStorageTokenStoreOutput {
	return o
}

func (o BlobStorageTokenStoreOutput) ToBlobStorageTokenStoreOutputWithContext(ctx context.Context) BlobStorageTokenStoreOutput {
	return o
}

func (o BlobStorageTokenStoreOutput) ToBlobStorageTokenStorePtrOutput() BlobStorageTokenStorePtrOutput {
	return o.ToBlobStorageTokenStorePtrOutputWithContext(context.Background())
}

func (o BlobStorageTokenStoreOutput) ToBlobStorageTokenStorePtrOutputWithContext(ctx context.Context) BlobStorageTokenStorePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlobStorageTokenStore) *BlobStorageTokenStore {
		return &v
	}).(BlobStorageTokenStorePtrOutput)
}

func (o BlobStorageTokenStoreOutput) SasUrlSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobStorageTokenStore) *string { return v.SasUrlSettingName }).(pulumi.StringPtrOutput)
}

type BlobStorageTokenStorePtrOutput struct{ *pulumi.OutputState }

func (BlobStorageTokenStorePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobStorageTokenStore)(nil)).Elem()
}

func (o BlobStorageTokenStorePtrOutput) ToBlobStorageTokenStorePtrOutput() BlobStorageTokenStorePtrOutput {
	return o
}

func (o BlobStorageTokenStorePtrOutput) ToBlobStorageTokenStorePtrOutputWithContext(ctx context.Context) BlobStorageTokenStorePtrOutput {
	return o
}

func (o BlobStorageTokenStorePtrOutput) Elem() BlobStorageTokenStoreOutput {
	return o.ApplyT(func(v *BlobStorageTokenStore) BlobStorageTokenStore {
		if v != nil {
			return *v
		}
		var ret BlobStorageTokenStore
		return ret
	}).(BlobStorageTokenStoreOutput)
}

func (o BlobStorageTokenStorePtrOutput) SasUrlSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobStorageTokenStore) *string {
		if v == nil {
			return nil
		}
		return v.SasUrlSettingName
	}).(pulumi.StringPtrOutput)
}

type BlobStorageTokenStoreResponse struct {
	SasUrlSettingName *string `pulumi:"sasUrlSettingName"`
}

type BlobStorageTokenStoreResponseOutput struct{ *pulumi.OutputState }

func (BlobStorageTokenStoreResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobStorageTokenStoreResponse)(nil)).Elem()
}

func (o BlobStorageTokenStoreResponseOutput) ToBlobStorageTokenStoreResponseOutput() BlobStorageTokenStoreResponseOutput {
	return o
}

func (o BlobStorageTokenStoreResponseOutput) ToBlobStorageTokenStoreResponseOutputWithContext(ctx context.Context) BlobStorageTokenStoreResponseOutput {
	return o
}

func (o BlobStorageTokenStoreResponseOutput) SasUrlSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobStorageTokenStoreResponse) *string { return v.SasUrlSettingName }).(pulumi.StringPtrOutput)
}

type BlobStorageTokenStoreResponsePtrOutput struct{ *pulumi.OutputState }

func (BlobStorageTokenStoreResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobStorageTokenStoreResponse)(nil)).Elem()
}

func (o BlobStorageTokenStoreResponsePtrOutput) ToBlobStorageTokenStoreResponsePtrOutput() BlobStorageTokenStoreResponsePtrOutput {
	return o
}

func (o BlobStorageTokenStoreResponsePtrOutput) ToBlobStorageTokenStoreResponsePtrOutputWithContext(ctx context.Context) BlobStorageTokenStoreResponsePtrOutput {
	return o
}

func (o BlobStorageTokenStoreResponsePtrOutput) Elem() BlobStorageTokenStoreResponseOutput {
	return o.ApplyT(func(v *BlobStorageTokenStoreResponse) BlobStorageTokenStoreResponse {
		if v != nil {
			return *v
		}
		var ret BlobStorageTokenStoreResponse
		return ret
	}).(BlobStorageTokenStoreResponseOutput)
}

func (o BlobStorageTokenStoreResponsePtrOutput) SasUrlSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobStorageTokenStoreResponse) *string {
		if v == nil {
			return nil
		}
		return v.SasUrlSettingName
	}).(pulumi.StringPtrOutput)
}

type Capability struct {
	Name   *string `pulumi:"name"`
	Reason *string `pulumi:"reason"`
	Value  *string `pulumi:"value"`
}





type CapabilityInput interface {
	pulumi.Input

	ToCapabilityOutput() CapabilityOutput
	ToCapabilityOutputWithContext(context.Context) CapabilityOutput
}

type CapabilityArgs struct {
	Name   pulumi.StringPtrInput `pulumi:"name"`
	Reason pulumi.StringPtrInput `pulumi:"reason"`
	Value  pulumi.StringPtrInput `pulumi:"value"`
}

func (CapabilityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Capability)(nil)).Elem()
}

func (i CapabilityArgs) ToCapabilityOutput() CapabilityOutput {
	return i.ToCapabilityOutputWithContext(context.Background())
}

func (i CapabilityArgs) ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapabilityOutput)
}





type CapabilityArrayInput interface {
	pulumi.Input

	ToCapabilityArrayOutput() CapabilityArrayOutput
	ToCapabilityArrayOutputWithContext(context.Context) CapabilityArrayOutput
}

type CapabilityArray []CapabilityInput

func (CapabilityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Capability)(nil)).Elem()
}

func (i CapabilityArray) ToCapabilityArrayOutput() CapabilityArrayOutput {
	return i.ToCapabilityArrayOutputWithContext(context.Background())
}

func (i CapabilityArray) ToCapabilityArrayOutputWithContext(ctx context.Context) CapabilityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapabilityArrayOutput)
}

type CapabilityOutput struct{ *pulumi.OutputState }

func (CapabilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Capability)(nil)).Elem()
}

func (o CapabilityOutput) ToCapabilityOutput() CapabilityOutput {
	return o
}

func (o CapabilityOutput) ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput {
	return o
}

func (o CapabilityOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Capability) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CapabilityOutput) Reason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Capability) *string { return v.Reason }).(pulumi.StringPtrOutput)
}

func (o CapabilityOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Capability) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type CapabilityArrayOutput struct{ *pulumi.OutputState }

func (CapabilityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Capability)(nil)).Elem()
}

func (o CapabilityArrayOutput) ToCapabilityArrayOutput() CapabilityArrayOutput {
	return o
}

func (o CapabilityArrayOutput) ToCapabilityArrayOutputWithContext(ctx context.Context) CapabilityArrayOutput {
	return o
}

func (o CapabilityArrayOutput) Index(i pulumi.IntInput) CapabilityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Capability {
		return vs[0].([]Capability)[vs[1].(int)]
	}).(CapabilityOutput)
}

type CapabilityResponse struct {
	Name   *string `pulumi:"name"`
	Reason *string `pulumi:"reason"`
	Value  *string `pulumi:"value"`
}

type CapabilityResponseOutput struct{ *pulumi.OutputState }

func (CapabilityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapabilityResponse)(nil)).Elem()
}

func (o CapabilityResponseOutput) ToCapabilityResponseOutput() CapabilityResponseOutput {
	return o
}

func (o CapabilityResponseOutput) ToCapabilityResponseOutputWithContext(ctx context.Context) CapabilityResponseOutput {
	return o
}

func (o CapabilityResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CapabilityResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CapabilityResponseOutput) Reason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CapabilityResponse) *string { return v.Reason }).(pulumi.StringPtrOutput)
}

func (o CapabilityResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CapabilityResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type CapabilityResponseArrayOutput struct{ *pulumi.OutputState }

func (CapabilityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CapabilityResponse)(nil)).Elem()
}

func (o CapabilityResponseArrayOutput) ToCapabilityResponseArrayOutput() CapabilityResponseArrayOutput {
	return o
}

func (o CapabilityResponseArrayOutput) ToCapabilityResponseArrayOutputWithContext(ctx context.Context) CapabilityResponseArrayOutput {
	return o
}

func (o CapabilityResponseArrayOutput) Index(i pulumi.IntInput) CapabilityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CapabilityResponse {
		return vs[0].([]CapabilityResponse)[vs[1].(int)]
	}).(CapabilityResponseOutput)
}

type ClientRegistration struct {
	ClientId                *string `pulumi:"clientId"`
	ClientSecretSettingName *string `pulumi:"clientSecretSettingName"`
}





type ClientRegistrationInput interface {
	pulumi.Input

	ToClientRegistrationOutput() ClientRegistrationOutput
	ToClientRegistrationOutputWithContext(context.Context) ClientRegistrationOutput
}

type ClientRegistrationArgs struct {
	ClientId                pulumi.StringPtrInput `pulumi:"clientId"`
	ClientSecretSettingName pulumi.StringPtrInput `pulumi:"clientSecretSettingName"`
}

func (ClientRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientRegistration)(nil)).Elem()
}

func (i ClientRegistrationArgs) ToClientRegistrationOutput() ClientRegistrationOutput {
	return i.ToClientRegistrationOutputWithContext(context.Background())
}

func (i ClientRegistrationArgs) ToClientRegistrationOutputWithContext(ctx context.Context) ClientRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientRegistrationOutput)
}

func (i ClientRegistrationArgs) ToClientRegistrationPtrOutput() ClientRegistrationPtrOutput {
	return i.ToClientRegistrationPtrOutputWithContext(context.Background())
}

func (i ClientRegistrationArgs) ToClientRegistrationPtrOutputWithContext(ctx context.Context) ClientRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientRegistrationOutput).ToClientRegistrationPtrOutputWithContext(ctx)
}









type ClientRegistrationPtrInput interface {
	pulumi.Input

	ToClientRegistrationPtrOutput() ClientRegistrationPtrOutput
	ToClientRegistrationPtrOutputWithContext(context.Context) ClientRegistrationPtrOutput
}

type clientRegistrationPtrType ClientRegistrationArgs

func ClientRegistrationPtr(v *ClientRegistrationArgs) ClientRegistrationPtrInput {
	return (*clientRegistrationPtrType)(v)
}

func (*clientRegistrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientRegistration)(nil)).Elem()
}

func (i *clientRegistrationPtrType) ToClientRegistrationPtrOutput() ClientRegistrationPtrOutput {
	return i.ToClientRegistrationPtrOutputWithContext(context.Background())
}

func (i *clientRegistrationPtrType) ToClientRegistrationPtrOutputWithContext(ctx context.Context) ClientRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientRegistrationPtrOutput)
}

type ClientRegistrationOutput struct{ *pulumi.OutputState }

func (ClientRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientRegistration)(nil)).Elem()
}

func (o ClientRegistrationOutput) ToClientRegistrationOutput() ClientRegistrationOutput {
	return o
}

func (o ClientRegistrationOutput) ToClientRegistrationOutputWithContext(ctx context.Context) ClientRegistrationOutput {
	return o
}

func (o ClientRegistrationOutput) ToClientRegistrationPtrOutput() ClientRegistrationPtrOutput {
	return o.ToClientRegistrationPtrOutputWithContext(context.Background())
}

func (o ClientRegistrationOutput) ToClientRegistrationPtrOutputWithContext(ctx context.Context) ClientRegistrationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClientRegistration) *ClientRegistration {
		return &v
	}).(ClientRegistrationPtrOutput)
}

func (o ClientRegistrationOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientRegistration) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ClientRegistrationOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientRegistration) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

type ClientRegistrationPtrOutput struct{ *pulumi.OutputState }

func (ClientRegistrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientRegistration)(nil)).Elem()
}

func (o ClientRegistrationPtrOutput) ToClientRegistrationPtrOutput() ClientRegistrationPtrOutput {
	return o
}

func (o ClientRegistrationPtrOutput) ToClientRegistrationPtrOutputWithContext(ctx context.Context) ClientRegistrationPtrOutput {
	return o
}

func (o ClientRegistrationPtrOutput) Elem() ClientRegistrationOutput {
	return o.ApplyT(func(v *ClientRegistration) ClientRegistration {
		if v != nil {
			return *v
		}
		var ret ClientRegistration
		return ret
	}).(ClientRegistrationOutput)
}

func (o ClientRegistrationPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ClientRegistrationPtrOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretSettingName
	}).(pulumi.StringPtrOutput)
}

type ClientRegistrationResponse struct {
	ClientId                *string `pulumi:"clientId"`
	ClientSecretSettingName *string `pulumi:"clientSecretSettingName"`
}

type ClientRegistrationResponseOutput struct{ *pulumi.OutputState }

func (ClientRegistrationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientRegistrationResponse)(nil)).Elem()
}

func (o ClientRegistrationResponseOutput) ToClientRegistrationResponseOutput() ClientRegistrationResponseOutput {
	return o
}

func (o ClientRegistrationResponseOutput) ToClientRegistrationResponseOutputWithContext(ctx context.Context) ClientRegistrationResponseOutput {
	return o
}

func (o ClientRegistrationResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientRegistrationResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ClientRegistrationResponseOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientRegistrationResponse) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

type ClientRegistrationResponsePtrOutput struct{ *pulumi.OutputState }

func (ClientRegistrationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientRegistrationResponse)(nil)).Elem()
}

func (o ClientRegistrationResponsePtrOutput) ToClientRegistrationResponsePtrOutput() ClientRegistrationResponsePtrOutput {
	return o
}

func (o ClientRegistrationResponsePtrOutput) ToClientRegistrationResponsePtrOutputWithContext(ctx context.Context) ClientRegistrationResponsePtrOutput {
	return o
}

func (o ClientRegistrationResponsePtrOutput) Elem() ClientRegistrationResponseOutput {
	return o.ApplyT(func(v *ClientRegistrationResponse) ClientRegistrationResponse {
		if v != nil {
			return *v
		}
		var ret ClientRegistrationResponse
		return ret
	}).(ClientRegistrationResponseOutput)
}

func (o ClientRegistrationResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ClientRegistrationResponsePtrOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretSettingName
	}).(pulumi.StringPtrOutput)
}

type CloningInfo struct {
	AppSettingsOverrides      map[string]string `pulumi:"appSettingsOverrides"`
	CloneCustomHostNames      *bool             `pulumi:"cloneCustomHostNames"`
	CloneSourceControl        *bool             `pulumi:"cloneSourceControl"`
	ConfigureLoadBalancing    *bool             `pulumi:"configureLoadBalancing"`
	CorrelationId             *string           `pulumi:"correlationId"`
	HostingEnvironment        *string           `pulumi:"hostingEnvironment"`
	Overwrite                 *bool             `pulumi:"overwrite"`
	SourceWebAppId            string            `pulumi:"sourceWebAppId"`
	SourceWebAppLocation      *string           `pulumi:"sourceWebAppLocation"`
	TrafficManagerProfileId   *string           `pulumi:"trafficManagerProfileId"`
	TrafficManagerProfileName *string           `pulumi:"trafficManagerProfileName"`
}





type CloningInfoInput interface {
	pulumi.Input

	ToCloningInfoOutput() CloningInfoOutput
	ToCloningInfoOutputWithContext(context.Context) CloningInfoOutput
}

type CloningInfoArgs struct {
	AppSettingsOverrides      pulumi.StringMapInput `pulumi:"appSettingsOverrides"`
	CloneCustomHostNames      pulumi.BoolPtrInput   `pulumi:"cloneCustomHostNames"`
	CloneSourceControl        pulumi.BoolPtrInput   `pulumi:"cloneSourceControl"`
	ConfigureLoadBalancing    pulumi.BoolPtrInput   `pulumi:"configureLoadBalancing"`
	CorrelationId             pulumi.StringPtrInput `pulumi:"correlationId"`
	HostingEnvironment        pulumi.StringPtrInput `pulumi:"hostingEnvironment"`
	Overwrite                 pulumi.BoolPtrInput   `pulumi:"overwrite"`
	SourceWebAppId            pulumi.StringInput    `pulumi:"sourceWebAppId"`
	SourceWebAppLocation      pulumi.StringPtrInput `pulumi:"sourceWebAppLocation"`
	TrafficManagerProfileId   pulumi.StringPtrInput `pulumi:"trafficManagerProfileId"`
	TrafficManagerProfileName pulumi.StringPtrInput `pulumi:"trafficManagerProfileName"`
}

func (CloningInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloningInfo)(nil)).Elem()
}

func (i CloningInfoArgs) ToCloningInfoOutput() CloningInfoOutput {
	return i.ToCloningInfoOutputWithContext(context.Background())
}

func (i CloningInfoArgs) ToCloningInfoOutputWithContext(ctx context.Context) CloningInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloningInfoOutput)
}

func (i CloningInfoArgs) ToCloningInfoPtrOutput() CloningInfoPtrOutput {
	return i.ToCloningInfoPtrOutputWithContext(context.Background())
}

func (i CloningInfoArgs) ToCloningInfoPtrOutputWithContext(ctx context.Context) CloningInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloningInfoOutput).ToCloningInfoPtrOutputWithContext(ctx)
}









type CloningInfoPtrInput interface {
	pulumi.Input

	ToCloningInfoPtrOutput() CloningInfoPtrOutput
	ToCloningInfoPtrOutputWithContext(context.Context) CloningInfoPtrOutput
}

type cloningInfoPtrType CloningInfoArgs

func CloningInfoPtr(v *CloningInfoArgs) CloningInfoPtrInput {
	return (*cloningInfoPtrType)(v)
}

func (*cloningInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloningInfo)(nil)).Elem()
}

func (i *cloningInfoPtrType) ToCloningInfoPtrOutput() CloningInfoPtrOutput {
	return i.ToCloningInfoPtrOutputWithContext(context.Background())
}

func (i *cloningInfoPtrType) ToCloningInfoPtrOutputWithContext(ctx context.Context) CloningInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloningInfoPtrOutput)
}

type CloningInfoOutput struct{ *pulumi.OutputState }

func (CloningInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloningInfo)(nil)).Elem()
}

func (o CloningInfoOutput) ToCloningInfoOutput() CloningInfoOutput {
	return o
}

func (o CloningInfoOutput) ToCloningInfoOutputWithContext(ctx context.Context) CloningInfoOutput {
	return o
}

func (o CloningInfoOutput) ToCloningInfoPtrOutput() CloningInfoPtrOutput {
	return o.ToCloningInfoPtrOutputWithContext(context.Background())
}

func (o CloningInfoOutput) ToCloningInfoPtrOutputWithContext(ctx context.Context) CloningInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloningInfo) *CloningInfo {
		return &v
	}).(CloningInfoPtrOutput)
}

func (o CloningInfoOutput) AppSettingsOverrides() pulumi.StringMapOutput {
	return o.ApplyT(func(v CloningInfo) map[string]string { return v.AppSettingsOverrides }).(pulumi.StringMapOutput)
}

func (o CloningInfoOutput) CloneCustomHostNames() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfo) *bool { return v.CloneCustomHostNames }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoOutput) CloneSourceControl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfo) *bool { return v.CloneSourceControl }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoOutput) ConfigureLoadBalancing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfo) *bool { return v.ConfigureLoadBalancing }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfo) *string { return v.CorrelationId }).(pulumi.StringPtrOutput)
}

func (o CloningInfoOutput) HostingEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfo) *string { return v.HostingEnvironment }).(pulumi.StringPtrOutput)
}

func (o CloningInfoOutput) Overwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfo) *bool { return v.Overwrite }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoOutput) SourceWebAppId() pulumi.StringOutput {
	return o.ApplyT(func(v CloningInfo) string { return v.SourceWebAppId }).(pulumi.StringOutput)
}

func (o CloningInfoOutput) SourceWebAppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfo) *string { return v.SourceWebAppLocation }).(pulumi.StringPtrOutput)
}

func (o CloningInfoOutput) TrafficManagerProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfo) *string { return v.TrafficManagerProfileId }).(pulumi.StringPtrOutput)
}

func (o CloningInfoOutput) TrafficManagerProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfo) *string { return v.TrafficManagerProfileName }).(pulumi.StringPtrOutput)
}

type CloningInfoPtrOutput struct{ *pulumi.OutputState }

func (CloningInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloningInfo)(nil)).Elem()
}

func (o CloningInfoPtrOutput) ToCloningInfoPtrOutput() CloningInfoPtrOutput {
	return o
}

func (o CloningInfoPtrOutput) ToCloningInfoPtrOutputWithContext(ctx context.Context) CloningInfoPtrOutput {
	return o
}

func (o CloningInfoPtrOutput) Elem() CloningInfoOutput {
	return o.ApplyT(func(v *CloningInfo) CloningInfo {
		if v != nil {
			return *v
		}
		var ret CloningInfo
		return ret
	}).(CloningInfoOutput)
}

func (o CloningInfoPtrOutput) AppSettingsOverrides() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CloningInfo) map[string]string {
		if v == nil {
			return nil
		}
		return v.AppSettingsOverrides
	}).(pulumi.StringMapOutput)
}

func (o CloningInfoPtrOutput) CloneCustomHostNames() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *bool {
		if v == nil {
			return nil
		}
		return v.CloneCustomHostNames
	}).(pulumi.BoolPtrOutput)
}

func (o CloningInfoPtrOutput) CloneSourceControl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *bool {
		if v == nil {
			return nil
		}
		return v.CloneSourceControl
	}).(pulumi.BoolPtrOutput)
}

func (o CloningInfoPtrOutput) ConfigureLoadBalancing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *bool {
		if v == nil {
			return nil
		}
		return v.ConfigureLoadBalancing
	}).(pulumi.BoolPtrOutput)
}

func (o CloningInfoPtrOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *string {
		if v == nil {
			return nil
		}
		return v.CorrelationId
	}).(pulumi.StringPtrOutput)
}

func (o CloningInfoPtrOutput) HostingEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *string {
		if v == nil {
			return nil
		}
		return v.HostingEnvironment
	}).(pulumi.StringPtrOutput)
}

func (o CloningInfoPtrOutput) Overwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *bool {
		if v == nil {
			return nil
		}
		return v.Overwrite
	}).(pulumi.BoolPtrOutput)
}

func (o CloningInfoPtrOutput) SourceWebAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *string {
		if v == nil {
			return nil
		}
		return &v.SourceWebAppId
	}).(pulumi.StringPtrOutput)
}

func (o CloningInfoPtrOutput) SourceWebAppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *string {
		if v == nil {
			return nil
		}
		return v.SourceWebAppLocation
	}).(pulumi.StringPtrOutput)
}

func (o CloningInfoPtrOutput) TrafficManagerProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *string {
		if v == nil {
			return nil
		}
		return v.TrafficManagerProfileId
	}).(pulumi.StringPtrOutput)
}

func (o CloningInfoPtrOutput) TrafficManagerProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *string {
		if v == nil {
			return nil
		}
		return v.TrafficManagerProfileName
	}).(pulumi.StringPtrOutput)
}

type ConnStringInfo struct {
	ConnectionString *string               `pulumi:"connectionString"`
	Name             *string               `pulumi:"name"`
	Type             *ConnectionStringType `pulumi:"type"`
}





type ConnStringInfoInput interface {
	pulumi.Input

	ToConnStringInfoOutput() ConnStringInfoOutput
	ToConnStringInfoOutputWithContext(context.Context) ConnStringInfoOutput
}

type ConnStringInfoArgs struct {
	ConnectionString pulumi.StringPtrInput        `pulumi:"connectionString"`
	Name             pulumi.StringPtrInput        `pulumi:"name"`
	Type             ConnectionStringTypePtrInput `pulumi:"type"`
}

func (ConnStringInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringInfo)(nil)).Elem()
}

func (i ConnStringInfoArgs) ToConnStringInfoOutput() ConnStringInfoOutput {
	return i.ToConnStringInfoOutputWithContext(context.Background())
}

func (i ConnStringInfoArgs) ToConnStringInfoOutputWithContext(ctx context.Context) ConnStringInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnStringInfoOutput)
}





type ConnStringInfoArrayInput interface {
	pulumi.Input

	ToConnStringInfoArrayOutput() ConnStringInfoArrayOutput
	ToConnStringInfoArrayOutputWithContext(context.Context) ConnStringInfoArrayOutput
}

type ConnStringInfoArray []ConnStringInfoInput

func (ConnStringInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnStringInfo)(nil)).Elem()
}

func (i ConnStringInfoArray) ToConnStringInfoArrayOutput() ConnStringInfoArrayOutput {
	return i.ToConnStringInfoArrayOutputWithContext(context.Background())
}

func (i ConnStringInfoArray) ToConnStringInfoArrayOutputWithContext(ctx context.Context) ConnStringInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnStringInfoArrayOutput)
}

type ConnStringInfoOutput struct{ *pulumi.OutputState }

func (ConnStringInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringInfo)(nil)).Elem()
}

func (o ConnStringInfoOutput) ToConnStringInfoOutput() ConnStringInfoOutput {
	return o
}

func (o ConnStringInfoOutput) ToConnStringInfoOutputWithContext(ctx context.Context) ConnStringInfoOutput {
	return o
}

func (o ConnStringInfoOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnStringInfo) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o ConnStringInfoOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnStringInfo) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConnStringInfoOutput) Type() ConnectionStringTypePtrOutput {
	return o.ApplyT(func(v ConnStringInfo) *ConnectionStringType { return v.Type }).(ConnectionStringTypePtrOutput)
}

type ConnStringInfoArrayOutput struct{ *pulumi.OutputState }

func (ConnStringInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnStringInfo)(nil)).Elem()
}

func (o ConnStringInfoArrayOutput) ToConnStringInfoArrayOutput() ConnStringInfoArrayOutput {
	return o
}

func (o ConnStringInfoArrayOutput) ToConnStringInfoArrayOutputWithContext(ctx context.Context) ConnStringInfoArrayOutput {
	return o
}

func (o ConnStringInfoArrayOutput) Index(i pulumi.IntInput) ConnStringInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnStringInfo {
		return vs[0].([]ConnStringInfo)[vs[1].(int)]
	}).(ConnStringInfoOutput)
}

type ConnStringInfoResponse struct {
	ConnectionString *string `pulumi:"connectionString"`
	Name             *string `pulumi:"name"`
	Type             *string `pulumi:"type"`
}

type ConnStringInfoResponseOutput struct{ *pulumi.OutputState }

func (ConnStringInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringInfoResponse)(nil)).Elem()
}

func (o ConnStringInfoResponseOutput) ToConnStringInfoResponseOutput() ConnStringInfoResponseOutput {
	return o
}

func (o ConnStringInfoResponseOutput) ToConnStringInfoResponseOutputWithContext(ctx context.Context) ConnStringInfoResponseOutput {
	return o
}

func (o ConnStringInfoResponseOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnStringInfoResponse) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o ConnStringInfoResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnStringInfoResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConnStringInfoResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnStringInfoResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ConnStringInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (ConnStringInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnStringInfoResponse)(nil)).Elem()
}

func (o ConnStringInfoResponseArrayOutput) ToConnStringInfoResponseArrayOutput() ConnStringInfoResponseArrayOutput {
	return o
}

func (o ConnStringInfoResponseArrayOutput) ToConnStringInfoResponseArrayOutputWithContext(ctx context.Context) ConnStringInfoResponseArrayOutput {
	return o
}

func (o ConnStringInfoResponseArrayOutput) Index(i pulumi.IntInput) ConnStringInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnStringInfoResponse {
		return vs[0].([]ConnStringInfoResponse)[vs[1].(int)]
	}).(ConnStringInfoResponseOutput)
}

type ConnStringValueTypePair struct {
	Type  ConnectionStringType `pulumi:"type"`
	Value string               `pulumi:"value"`
}





type ConnStringValueTypePairInput interface {
	pulumi.Input

	ToConnStringValueTypePairOutput() ConnStringValueTypePairOutput
	ToConnStringValueTypePairOutputWithContext(context.Context) ConnStringValueTypePairOutput
}

type ConnStringValueTypePairArgs struct {
	Type  ConnectionStringTypeInput `pulumi:"type"`
	Value pulumi.StringInput        `pulumi:"value"`
}

func (ConnStringValueTypePairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringValueTypePair)(nil)).Elem()
}

func (i ConnStringValueTypePairArgs) ToConnStringValueTypePairOutput() ConnStringValueTypePairOutput {
	return i.ToConnStringValueTypePairOutputWithContext(context.Background())
}

func (i ConnStringValueTypePairArgs) ToConnStringValueTypePairOutputWithContext(ctx context.Context) ConnStringValueTypePairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnStringValueTypePairOutput)
}





type ConnStringValueTypePairMapInput interface {
	pulumi.Input

	ToConnStringValueTypePairMapOutput() ConnStringValueTypePairMapOutput
	ToConnStringValueTypePairMapOutputWithContext(context.Context) ConnStringValueTypePairMapOutput
}

type ConnStringValueTypePairMap map[string]ConnStringValueTypePairInput

func (ConnStringValueTypePairMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnStringValueTypePair)(nil)).Elem()
}

func (i ConnStringValueTypePairMap) ToConnStringValueTypePairMapOutput() ConnStringValueTypePairMapOutput {
	return i.ToConnStringValueTypePairMapOutputWithContext(context.Background())
}

func (i ConnStringValueTypePairMap) ToConnStringValueTypePairMapOutputWithContext(ctx context.Context) ConnStringValueTypePairMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnStringValueTypePairMapOutput)
}

type ConnStringValueTypePairOutput struct{ *pulumi.OutputState }

func (ConnStringValueTypePairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringValueTypePair)(nil)).Elem()
}

func (o ConnStringValueTypePairOutput) ToConnStringValueTypePairOutput() ConnStringValueTypePairOutput {
	return o
}

func (o ConnStringValueTypePairOutput) ToConnStringValueTypePairOutputWithContext(ctx context.Context) ConnStringValueTypePairOutput {
	return o
}

func (o ConnStringValueTypePairOutput) Type() ConnectionStringTypeOutput {
	return o.ApplyT(func(v ConnStringValueTypePair) ConnectionStringType { return v.Type }).(ConnectionStringTypeOutput)
}

func (o ConnStringValueTypePairOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ConnStringValueTypePair) string { return v.Value }).(pulumi.StringOutput)
}

type ConnStringValueTypePairMapOutput struct{ *pulumi.OutputState }

func (ConnStringValueTypePairMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnStringValueTypePair)(nil)).Elem()
}

func (o ConnStringValueTypePairMapOutput) ToConnStringValueTypePairMapOutput() ConnStringValueTypePairMapOutput {
	return o
}

func (o ConnStringValueTypePairMapOutput) ToConnStringValueTypePairMapOutputWithContext(ctx context.Context) ConnStringValueTypePairMapOutput {
	return o
}

func (o ConnStringValueTypePairMapOutput) MapIndex(k pulumi.StringInput) ConnStringValueTypePairOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ConnStringValueTypePair {
		return vs[0].(map[string]ConnStringValueTypePair)[vs[1].(string)]
	}).(ConnStringValueTypePairOutput)
}

type ConnStringValueTypePairResponse struct {
	Type  string `pulumi:"type"`
	Value string `pulumi:"value"`
}

type ConnStringValueTypePairResponseOutput struct{ *pulumi.OutputState }

func (ConnStringValueTypePairResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringValueTypePairResponse)(nil)).Elem()
}

func (o ConnStringValueTypePairResponseOutput) ToConnStringValueTypePairResponseOutput() ConnStringValueTypePairResponseOutput {
	return o
}

func (o ConnStringValueTypePairResponseOutput) ToConnStringValueTypePairResponseOutputWithContext(ctx context.Context) ConnStringValueTypePairResponseOutput {
	return o
}

func (o ConnStringValueTypePairResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ConnStringValueTypePairResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ConnStringValueTypePairResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ConnStringValueTypePairResponse) string { return v.Value }).(pulumi.StringOutput)
}

type ConnStringValueTypePairResponseMapOutput struct{ *pulumi.OutputState }

func (ConnStringValueTypePairResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnStringValueTypePairResponse)(nil)).Elem()
}

func (o ConnStringValueTypePairResponseMapOutput) ToConnStringValueTypePairResponseMapOutput() ConnStringValueTypePairResponseMapOutput {
	return o
}

func (o ConnStringValueTypePairResponseMapOutput) ToConnStringValueTypePairResponseMapOutputWithContext(ctx context.Context) ConnStringValueTypePairResponseMapOutput {
	return o
}

func (o ConnStringValueTypePairResponseMapOutput) MapIndex(k pulumi.StringInput) ConnStringValueTypePairResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ConnStringValueTypePairResponse {
		return vs[0].(map[string]ConnStringValueTypePairResponse)[vs[1].(string)]
	}).(ConnStringValueTypePairResponseOutput)
}

type CookieExpiration struct {
	Convention       *CookieExpirationConvention `pulumi:"convention"`
	TimeToExpiration *string                     `pulumi:"timeToExpiration"`
}





type CookieExpirationInput interface {
	pulumi.Input

	ToCookieExpirationOutput() CookieExpirationOutput
	ToCookieExpirationOutputWithContext(context.Context) CookieExpirationOutput
}

type CookieExpirationArgs struct {
	Convention       CookieExpirationConventionPtrInput `pulumi:"convention"`
	TimeToExpiration pulumi.StringPtrInput              `pulumi:"timeToExpiration"`
}

func (CookieExpirationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CookieExpiration)(nil)).Elem()
}

func (i CookieExpirationArgs) ToCookieExpirationOutput() CookieExpirationOutput {
	return i.ToCookieExpirationOutputWithContext(context.Background())
}

func (i CookieExpirationArgs) ToCookieExpirationOutputWithContext(ctx context.Context) CookieExpirationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CookieExpirationOutput)
}

func (i CookieExpirationArgs) ToCookieExpirationPtrOutput() CookieExpirationPtrOutput {
	return i.ToCookieExpirationPtrOutputWithContext(context.Background())
}

func (i CookieExpirationArgs) ToCookieExpirationPtrOutputWithContext(ctx context.Context) CookieExpirationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CookieExpirationOutput).ToCookieExpirationPtrOutputWithContext(ctx)
}









type CookieExpirationPtrInput interface {
	pulumi.Input

	ToCookieExpirationPtrOutput() CookieExpirationPtrOutput
	ToCookieExpirationPtrOutputWithContext(context.Context) CookieExpirationPtrOutput
}

type cookieExpirationPtrType CookieExpirationArgs

func CookieExpirationPtr(v *CookieExpirationArgs) CookieExpirationPtrInput {
	return (*cookieExpirationPtrType)(v)
}

func (*cookieExpirationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CookieExpiration)(nil)).Elem()
}

func (i *cookieExpirationPtrType) ToCookieExpirationPtrOutput() CookieExpirationPtrOutput {
	return i.ToCookieExpirationPtrOutputWithContext(context.Background())
}

func (i *cookieExpirationPtrType) ToCookieExpirationPtrOutputWithContext(ctx context.Context) CookieExpirationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CookieExpirationPtrOutput)
}

type CookieExpirationOutput struct{ *pulumi.OutputState }

func (CookieExpirationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CookieExpiration)(nil)).Elem()
}

func (o CookieExpirationOutput) ToCookieExpirationOutput() CookieExpirationOutput {
	return o
}

func (o CookieExpirationOutput) ToCookieExpirationOutputWithContext(ctx context.Context) CookieExpirationOutput {
	return o
}

func (o CookieExpirationOutput) ToCookieExpirationPtrOutput() CookieExpirationPtrOutput {
	return o.ToCookieExpirationPtrOutputWithContext(context.Background())
}

func (o CookieExpirationOutput) ToCookieExpirationPtrOutputWithContext(ctx context.Context) CookieExpirationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CookieExpiration) *CookieExpiration {
		return &v
	}).(CookieExpirationPtrOutput)
}

func (o CookieExpirationOutput) Convention() CookieExpirationConventionPtrOutput {
	return o.ApplyT(func(v CookieExpiration) *CookieExpirationConvention { return v.Convention }).(CookieExpirationConventionPtrOutput)
}

func (o CookieExpirationOutput) TimeToExpiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CookieExpiration) *string { return v.TimeToExpiration }).(pulumi.StringPtrOutput)
}

type CookieExpirationPtrOutput struct{ *pulumi.OutputState }

func (CookieExpirationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CookieExpiration)(nil)).Elem()
}

func (o CookieExpirationPtrOutput) ToCookieExpirationPtrOutput() CookieExpirationPtrOutput {
	return o
}

func (o CookieExpirationPtrOutput) ToCookieExpirationPtrOutputWithContext(ctx context.Context) CookieExpirationPtrOutput {
	return o
}

func (o CookieExpirationPtrOutput) Elem() CookieExpirationOutput {
	return o.ApplyT(func(v *CookieExpiration) CookieExpiration {
		if v != nil {
			return *v
		}
		var ret CookieExpiration
		return ret
	}).(CookieExpirationOutput)
}

func (o CookieExpirationPtrOutput) Convention() CookieExpirationConventionPtrOutput {
	return o.ApplyT(func(v *CookieExpiration) *CookieExpirationConvention {
		if v == nil {
			return nil
		}
		return v.Convention
	}).(CookieExpirationConventionPtrOutput)
}

func (o CookieExpirationPtrOutput) TimeToExpiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CookieExpiration) *string {
		if v == nil {
			return nil
		}
		return v.TimeToExpiration
	}).(pulumi.StringPtrOutput)
}

type CookieExpirationResponse struct {
	Convention       *string `pulumi:"convention"`
	TimeToExpiration *string `pulumi:"timeToExpiration"`
}

type CookieExpirationResponseOutput struct{ *pulumi.OutputState }

func (CookieExpirationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CookieExpirationResponse)(nil)).Elem()
}

func (o CookieExpirationResponseOutput) ToCookieExpirationResponseOutput() CookieExpirationResponseOutput {
	return o
}

func (o CookieExpirationResponseOutput) ToCookieExpirationResponseOutputWithContext(ctx context.Context) CookieExpirationResponseOutput {
	return o
}

func (o CookieExpirationResponseOutput) Convention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CookieExpirationResponse) *string { return v.Convention }).(pulumi.StringPtrOutput)
}

func (o CookieExpirationResponseOutput) TimeToExpiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CookieExpirationResponse) *string { return v.TimeToExpiration }).(pulumi.StringPtrOutput)
}

type CookieExpirationResponsePtrOutput struct{ *pulumi.OutputState }

func (CookieExpirationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CookieExpirationResponse)(nil)).Elem()
}

func (o CookieExpirationResponsePtrOutput) ToCookieExpirationResponsePtrOutput() CookieExpirationResponsePtrOutput {
	return o
}

func (o CookieExpirationResponsePtrOutput) ToCookieExpirationResponsePtrOutputWithContext(ctx context.Context) CookieExpirationResponsePtrOutput {
	return o
}

func (o CookieExpirationResponsePtrOutput) Elem() CookieExpirationResponseOutput {
	return o.ApplyT(func(v *CookieExpirationResponse) CookieExpirationResponse {
		if v != nil {
			return *v
		}
		var ret CookieExpirationResponse
		return ret
	}).(CookieExpirationResponseOutput)
}

func (o CookieExpirationResponsePtrOutput) Convention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CookieExpirationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Convention
	}).(pulumi.StringPtrOutput)
}

func (o CookieExpirationResponsePtrOutput) TimeToExpiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CookieExpirationResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeToExpiration
	}).(pulumi.StringPtrOutput)
}

type CorsSettings struct {
	AllowedOrigins     []string `pulumi:"allowedOrigins"`
	SupportCredentials *bool    `pulumi:"supportCredentials"`
}





type CorsSettingsInput interface {
	pulumi.Input

	ToCorsSettingsOutput() CorsSettingsOutput
	ToCorsSettingsOutputWithContext(context.Context) CorsSettingsOutput
}

type CorsSettingsArgs struct {
	AllowedOrigins     pulumi.StringArrayInput `pulumi:"allowedOrigins"`
	SupportCredentials pulumi.BoolPtrInput     `pulumi:"supportCredentials"`
}

func (CorsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsSettings)(nil)).Elem()
}

func (i CorsSettingsArgs) ToCorsSettingsOutput() CorsSettingsOutput {
	return i.ToCorsSettingsOutputWithContext(context.Background())
}

func (i CorsSettingsArgs) ToCorsSettingsOutputWithContext(ctx context.Context) CorsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsSettingsOutput)
}

func (i CorsSettingsArgs) ToCorsSettingsPtrOutput() CorsSettingsPtrOutput {
	return i.ToCorsSettingsPtrOutputWithContext(context.Background())
}

func (i CorsSettingsArgs) ToCorsSettingsPtrOutputWithContext(ctx context.Context) CorsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsSettingsOutput).ToCorsSettingsPtrOutputWithContext(ctx)
}









type CorsSettingsPtrInput interface {
	pulumi.Input

	ToCorsSettingsPtrOutput() CorsSettingsPtrOutput
	ToCorsSettingsPtrOutputWithContext(context.Context) CorsSettingsPtrOutput
}

type corsSettingsPtrType CorsSettingsArgs

func CorsSettingsPtr(v *CorsSettingsArgs) CorsSettingsPtrInput {
	return (*corsSettingsPtrType)(v)
}

func (*corsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CorsSettings)(nil)).Elem()
}

func (i *corsSettingsPtrType) ToCorsSettingsPtrOutput() CorsSettingsPtrOutput {
	return i.ToCorsSettingsPtrOutputWithContext(context.Background())
}

func (i *corsSettingsPtrType) ToCorsSettingsPtrOutputWithContext(ctx context.Context) CorsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsSettingsPtrOutput)
}

type CorsSettingsOutput struct{ *pulumi.OutputState }

func (CorsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsSettings)(nil)).Elem()
}

func (o CorsSettingsOutput) ToCorsSettingsOutput() CorsSettingsOutput {
	return o
}

func (o CorsSettingsOutput) ToCorsSettingsOutputWithContext(ctx context.Context) CorsSettingsOutput {
	return o
}

func (o CorsSettingsOutput) ToCorsSettingsPtrOutput() CorsSettingsPtrOutput {
	return o.ToCorsSettingsPtrOutputWithContext(context.Background())
}

func (o CorsSettingsOutput) ToCorsSettingsPtrOutputWithContext(ctx context.Context) CorsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CorsSettings) *CorsSettings {
		return &v
	}).(CorsSettingsPtrOutput)
}

func (o CorsSettingsOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsSettings) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

func (o CorsSettingsOutput) SupportCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CorsSettings) *bool { return v.SupportCredentials }).(pulumi.BoolPtrOutput)
}

type CorsSettingsPtrOutput struct{ *pulumi.OutputState }

func (CorsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CorsSettings)(nil)).Elem()
}

func (o CorsSettingsPtrOutput) ToCorsSettingsPtrOutput() CorsSettingsPtrOutput {
	return o
}

func (o CorsSettingsPtrOutput) ToCorsSettingsPtrOutputWithContext(ctx context.Context) CorsSettingsPtrOutput {
	return o
}

func (o CorsSettingsPtrOutput) Elem() CorsSettingsOutput {
	return o.ApplyT(func(v *CorsSettings) CorsSettings {
		if v != nil {
			return *v
		}
		var ret CorsSettings
		return ret
	}).(CorsSettingsOutput)
}

func (o CorsSettingsPtrOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CorsSettings) []string {
		if v == nil {
			return nil
		}
		return v.AllowedOrigins
	}).(pulumi.StringArrayOutput)
}

func (o CorsSettingsPtrOutput) SupportCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CorsSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SupportCredentials
	}).(pulumi.BoolPtrOutput)
}

type CorsSettingsResponse struct {
	AllowedOrigins     []string `pulumi:"allowedOrigins"`
	SupportCredentials *bool    `pulumi:"supportCredentials"`
}

type CorsSettingsResponseOutput struct{ *pulumi.OutputState }

func (CorsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsSettingsResponse)(nil)).Elem()
}

func (o CorsSettingsResponseOutput) ToCorsSettingsResponseOutput() CorsSettingsResponseOutput {
	return o
}

func (o CorsSettingsResponseOutput) ToCorsSettingsResponseOutputWithContext(ctx context.Context) CorsSettingsResponseOutput {
	return o
}

func (o CorsSettingsResponseOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsSettingsResponse) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

func (o CorsSettingsResponseOutput) SupportCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CorsSettingsResponse) *bool { return v.SupportCredentials }).(pulumi.BoolPtrOutput)
}

type CorsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (CorsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CorsSettingsResponse)(nil)).Elem()
}

func (o CorsSettingsResponsePtrOutput) ToCorsSettingsResponsePtrOutput() CorsSettingsResponsePtrOutput {
	return o
}

func (o CorsSettingsResponsePtrOutput) ToCorsSettingsResponsePtrOutputWithContext(ctx context.Context) CorsSettingsResponsePtrOutput {
	return o
}

func (o CorsSettingsResponsePtrOutput) Elem() CorsSettingsResponseOutput {
	return o.ApplyT(func(v *CorsSettingsResponse) CorsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret CorsSettingsResponse
		return ret
	}).(CorsSettingsResponseOutput)
}

func (o CorsSettingsResponsePtrOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CorsSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedOrigins
	}).(pulumi.StringArrayOutput)
}

func (o CorsSettingsResponsePtrOutput) SupportCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CorsSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SupportCredentials
	}).(pulumi.BoolPtrOutput)
}

type CustomOpenIdConnectProvider struct {
	Enabled      *bool                      `pulumi:"enabled"`
	Login        *OpenIdConnectLogin        `pulumi:"login"`
	Registration *OpenIdConnectRegistration `pulumi:"registration"`
}





type CustomOpenIdConnectProviderInput interface {
	pulumi.Input

	ToCustomOpenIdConnectProviderOutput() CustomOpenIdConnectProviderOutput
	ToCustomOpenIdConnectProviderOutputWithContext(context.Context) CustomOpenIdConnectProviderOutput
}

type CustomOpenIdConnectProviderArgs struct {
	Enabled      pulumi.BoolPtrInput               `pulumi:"enabled"`
	Login        OpenIdConnectLoginPtrInput        `pulumi:"login"`
	Registration OpenIdConnectRegistrationPtrInput `pulumi:"registration"`
}

func (CustomOpenIdConnectProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomOpenIdConnectProvider)(nil)).Elem()
}

func (i CustomOpenIdConnectProviderArgs) ToCustomOpenIdConnectProviderOutput() CustomOpenIdConnectProviderOutput {
	return i.ToCustomOpenIdConnectProviderOutputWithContext(context.Background())
}

func (i CustomOpenIdConnectProviderArgs) ToCustomOpenIdConnectProviderOutputWithContext(ctx context.Context) CustomOpenIdConnectProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomOpenIdConnectProviderOutput)
}





type CustomOpenIdConnectProviderMapInput interface {
	pulumi.Input

	ToCustomOpenIdConnectProviderMapOutput() CustomOpenIdConnectProviderMapOutput
	ToCustomOpenIdConnectProviderMapOutputWithContext(context.Context) CustomOpenIdConnectProviderMapOutput
}

type CustomOpenIdConnectProviderMap map[string]CustomOpenIdConnectProviderInput

func (CustomOpenIdConnectProviderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomOpenIdConnectProvider)(nil)).Elem()
}

func (i CustomOpenIdConnectProviderMap) ToCustomOpenIdConnectProviderMapOutput() CustomOpenIdConnectProviderMapOutput {
	return i.ToCustomOpenIdConnectProviderMapOutputWithContext(context.Background())
}

func (i CustomOpenIdConnectProviderMap) ToCustomOpenIdConnectProviderMapOutputWithContext(ctx context.Context) CustomOpenIdConnectProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomOpenIdConnectProviderMapOutput)
}

type CustomOpenIdConnectProviderOutput struct{ *pulumi.OutputState }

func (CustomOpenIdConnectProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomOpenIdConnectProvider)(nil)).Elem()
}

func (o CustomOpenIdConnectProviderOutput) ToCustomOpenIdConnectProviderOutput() CustomOpenIdConnectProviderOutput {
	return o
}

func (o CustomOpenIdConnectProviderOutput) ToCustomOpenIdConnectProviderOutputWithContext(ctx context.Context) CustomOpenIdConnectProviderOutput {
	return o
}

func (o CustomOpenIdConnectProviderOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomOpenIdConnectProvider) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o CustomOpenIdConnectProviderOutput) Login() OpenIdConnectLoginPtrOutput {
	return o.ApplyT(func(v CustomOpenIdConnectProvider) *OpenIdConnectLogin { return v.Login }).(OpenIdConnectLoginPtrOutput)
}

func (o CustomOpenIdConnectProviderOutput) Registration() OpenIdConnectRegistrationPtrOutput {
	return o.ApplyT(func(v CustomOpenIdConnectProvider) *OpenIdConnectRegistration { return v.Registration }).(OpenIdConnectRegistrationPtrOutput)
}

type CustomOpenIdConnectProviderMapOutput struct{ *pulumi.OutputState }

func (CustomOpenIdConnectProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomOpenIdConnectProvider)(nil)).Elem()
}

func (o CustomOpenIdConnectProviderMapOutput) ToCustomOpenIdConnectProviderMapOutput() CustomOpenIdConnectProviderMapOutput {
	return o
}

func (o CustomOpenIdConnectProviderMapOutput) ToCustomOpenIdConnectProviderMapOutputWithContext(ctx context.Context) CustomOpenIdConnectProviderMapOutput {
	return o
}

func (o CustomOpenIdConnectProviderMapOutput) MapIndex(k pulumi.StringInput) CustomOpenIdConnectProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CustomOpenIdConnectProvider {
		return vs[0].(map[string]CustomOpenIdConnectProvider)[vs[1].(string)]
	}).(CustomOpenIdConnectProviderOutput)
}

type CustomOpenIdConnectProviderResponse struct {
	Enabled      *bool                              `pulumi:"enabled"`
	Login        *OpenIdConnectLoginResponse        `pulumi:"login"`
	Registration *OpenIdConnectRegistrationResponse `pulumi:"registration"`
}

type CustomOpenIdConnectProviderResponseOutput struct{ *pulumi.OutputState }

func (CustomOpenIdConnectProviderResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomOpenIdConnectProviderResponse)(nil)).Elem()
}

func (o CustomOpenIdConnectProviderResponseOutput) ToCustomOpenIdConnectProviderResponseOutput() CustomOpenIdConnectProviderResponseOutput {
	return o
}

func (o CustomOpenIdConnectProviderResponseOutput) ToCustomOpenIdConnectProviderResponseOutputWithContext(ctx context.Context) CustomOpenIdConnectProviderResponseOutput {
	return o
}

func (o CustomOpenIdConnectProviderResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomOpenIdConnectProviderResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o CustomOpenIdConnectProviderResponseOutput) Login() OpenIdConnectLoginResponsePtrOutput {
	return o.ApplyT(func(v CustomOpenIdConnectProviderResponse) *OpenIdConnectLoginResponse { return v.Login }).(OpenIdConnectLoginResponsePtrOutput)
}

func (o CustomOpenIdConnectProviderResponseOutput) Registration() OpenIdConnectRegistrationResponsePtrOutput {
	return o.ApplyT(func(v CustomOpenIdConnectProviderResponse) *OpenIdConnectRegistrationResponse { return v.Registration }).(OpenIdConnectRegistrationResponsePtrOutput)
}

type CustomOpenIdConnectProviderResponseMapOutput struct{ *pulumi.OutputState }

func (CustomOpenIdConnectProviderResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CustomOpenIdConnectProviderResponse)(nil)).Elem()
}

func (o CustomOpenIdConnectProviderResponseMapOutput) ToCustomOpenIdConnectProviderResponseMapOutput() CustomOpenIdConnectProviderResponseMapOutput {
	return o
}

func (o CustomOpenIdConnectProviderResponseMapOutput) ToCustomOpenIdConnectProviderResponseMapOutputWithContext(ctx context.Context) CustomOpenIdConnectProviderResponseMapOutput {
	return o
}

func (o CustomOpenIdConnectProviderResponseMapOutput) MapIndex(k pulumi.StringInput) CustomOpenIdConnectProviderResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CustomOpenIdConnectProviderResponse {
		return vs[0].(map[string]CustomOpenIdConnectProviderResponse)[vs[1].(string)]
	}).(CustomOpenIdConnectProviderResponseOutput)
}

type DatabaseBackupSetting struct {
	ConnectionString     *string `pulumi:"connectionString"`
	ConnectionStringName *string `pulumi:"connectionStringName"`
	DatabaseType         string  `pulumi:"databaseType"`
	Name                 *string `pulumi:"name"`
}





type DatabaseBackupSettingInput interface {
	pulumi.Input

	ToDatabaseBackupSettingOutput() DatabaseBackupSettingOutput
	ToDatabaseBackupSettingOutputWithContext(context.Context) DatabaseBackupSettingOutput
}

type DatabaseBackupSettingArgs struct {
	ConnectionString     pulumi.StringPtrInput `pulumi:"connectionString"`
	ConnectionStringName pulumi.StringPtrInput `pulumi:"connectionStringName"`
	DatabaseType         pulumi.StringInput    `pulumi:"databaseType"`
	Name                 pulumi.StringPtrInput `pulumi:"name"`
}

func (DatabaseBackupSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseBackupSetting)(nil)).Elem()
}

func (i DatabaseBackupSettingArgs) ToDatabaseBackupSettingOutput() DatabaseBackupSettingOutput {
	return i.ToDatabaseBackupSettingOutputWithContext(context.Background())
}

func (i DatabaseBackupSettingArgs) ToDatabaseBackupSettingOutputWithContext(ctx context.Context) DatabaseBackupSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseBackupSettingOutput)
}





type DatabaseBackupSettingArrayInput interface {
	pulumi.Input

	ToDatabaseBackupSettingArrayOutput() DatabaseBackupSettingArrayOutput
	ToDatabaseBackupSettingArrayOutputWithContext(context.Context) DatabaseBackupSettingArrayOutput
}

type DatabaseBackupSettingArray []DatabaseBackupSettingInput

func (DatabaseBackupSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseBackupSetting)(nil)).Elem()
}

func (i DatabaseBackupSettingArray) ToDatabaseBackupSettingArrayOutput() DatabaseBackupSettingArrayOutput {
	return i.ToDatabaseBackupSettingArrayOutputWithContext(context.Background())
}

func (i DatabaseBackupSettingArray) ToDatabaseBackupSettingArrayOutputWithContext(ctx context.Context) DatabaseBackupSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseBackupSettingArrayOutput)
}

type DatabaseBackupSettingOutput struct{ *pulumi.OutputState }

func (DatabaseBackupSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseBackupSetting)(nil)).Elem()
}

func (o DatabaseBackupSettingOutput) ToDatabaseBackupSettingOutput() DatabaseBackupSettingOutput {
	return o
}

func (o DatabaseBackupSettingOutput) ToDatabaseBackupSettingOutputWithContext(ctx context.Context) DatabaseBackupSettingOutput {
	return o
}

func (o DatabaseBackupSettingOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSetting) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o DatabaseBackupSettingOutput) ConnectionStringName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSetting) *string { return v.ConnectionStringName }).(pulumi.StringPtrOutput)
}

func (o DatabaseBackupSettingOutput) DatabaseType() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseBackupSetting) string { return v.DatabaseType }).(pulumi.StringOutput)
}

func (o DatabaseBackupSettingOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSetting) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DatabaseBackupSettingArrayOutput struct{ *pulumi.OutputState }

func (DatabaseBackupSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseBackupSetting)(nil)).Elem()
}

func (o DatabaseBackupSettingArrayOutput) ToDatabaseBackupSettingArrayOutput() DatabaseBackupSettingArrayOutput {
	return o
}

func (o DatabaseBackupSettingArrayOutput) ToDatabaseBackupSettingArrayOutputWithContext(ctx context.Context) DatabaseBackupSettingArrayOutput {
	return o
}

func (o DatabaseBackupSettingArrayOutput) Index(i pulumi.IntInput) DatabaseBackupSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseBackupSetting {
		return vs[0].([]DatabaseBackupSetting)[vs[1].(int)]
	}).(DatabaseBackupSettingOutput)
}

type DatabaseBackupSettingResponse struct {
	ConnectionString     *string `pulumi:"connectionString"`
	ConnectionStringName *string `pulumi:"connectionStringName"`
	DatabaseType         string  `pulumi:"databaseType"`
	Name                 *string `pulumi:"name"`
}

type DatabaseBackupSettingResponseOutput struct{ *pulumi.OutputState }

func (DatabaseBackupSettingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseBackupSettingResponse)(nil)).Elem()
}

func (o DatabaseBackupSettingResponseOutput) ToDatabaseBackupSettingResponseOutput() DatabaseBackupSettingResponseOutput {
	return o
}

func (o DatabaseBackupSettingResponseOutput) ToDatabaseBackupSettingResponseOutputWithContext(ctx context.Context) DatabaseBackupSettingResponseOutput {
	return o
}

func (o DatabaseBackupSettingResponseOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSettingResponse) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o DatabaseBackupSettingResponseOutput) ConnectionStringName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSettingResponse) *string { return v.ConnectionStringName }).(pulumi.StringPtrOutput)
}

func (o DatabaseBackupSettingResponseOutput) DatabaseType() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseBackupSettingResponse) string { return v.DatabaseType }).(pulumi.StringOutput)
}

func (o DatabaseBackupSettingResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSettingResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DatabaseBackupSettingResponseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseBackupSettingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseBackupSettingResponse)(nil)).Elem()
}

func (o DatabaseBackupSettingResponseArrayOutput) ToDatabaseBackupSettingResponseArrayOutput() DatabaseBackupSettingResponseArrayOutput {
	return o
}

func (o DatabaseBackupSettingResponseArrayOutput) ToDatabaseBackupSettingResponseArrayOutputWithContext(ctx context.Context) DatabaseBackupSettingResponseArrayOutput {
	return o
}

func (o DatabaseBackupSettingResponseArrayOutput) Index(i pulumi.IntInput) DatabaseBackupSettingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseBackupSettingResponse {
		return vs[0].([]DatabaseBackupSettingResponse)[vs[1].(int)]
	}).(DatabaseBackupSettingResponseOutput)
}

type DefaultAuthorizationPolicy struct {
	AllowedApplications []string           `pulumi:"allowedApplications"`
	AllowedPrincipals   *AllowedPrincipals `pulumi:"allowedPrincipals"`
}





type DefaultAuthorizationPolicyInput interface {
	pulumi.Input

	ToDefaultAuthorizationPolicyOutput() DefaultAuthorizationPolicyOutput
	ToDefaultAuthorizationPolicyOutputWithContext(context.Context) DefaultAuthorizationPolicyOutput
}

type DefaultAuthorizationPolicyArgs struct {
	AllowedApplications pulumi.StringArrayInput   `pulumi:"allowedApplications"`
	AllowedPrincipals   AllowedPrincipalsPtrInput `pulumi:"allowedPrincipals"`
}

func (DefaultAuthorizationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultAuthorizationPolicy)(nil)).Elem()
}

func (i DefaultAuthorizationPolicyArgs) ToDefaultAuthorizationPolicyOutput() DefaultAuthorizationPolicyOutput {
	return i.ToDefaultAuthorizationPolicyOutputWithContext(context.Background())
}

func (i DefaultAuthorizationPolicyArgs) ToDefaultAuthorizationPolicyOutputWithContext(ctx context.Context) DefaultAuthorizationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultAuthorizationPolicyOutput)
}

func (i DefaultAuthorizationPolicyArgs) ToDefaultAuthorizationPolicyPtrOutput() DefaultAuthorizationPolicyPtrOutput {
	return i.ToDefaultAuthorizationPolicyPtrOutputWithContext(context.Background())
}

func (i DefaultAuthorizationPolicyArgs) ToDefaultAuthorizationPolicyPtrOutputWithContext(ctx context.Context) DefaultAuthorizationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultAuthorizationPolicyOutput).ToDefaultAuthorizationPolicyPtrOutputWithContext(ctx)
}









type DefaultAuthorizationPolicyPtrInput interface {
	pulumi.Input

	ToDefaultAuthorizationPolicyPtrOutput() DefaultAuthorizationPolicyPtrOutput
	ToDefaultAuthorizationPolicyPtrOutputWithContext(context.Context) DefaultAuthorizationPolicyPtrOutput
}

type defaultAuthorizationPolicyPtrType DefaultAuthorizationPolicyArgs

func DefaultAuthorizationPolicyPtr(v *DefaultAuthorizationPolicyArgs) DefaultAuthorizationPolicyPtrInput {
	return (*defaultAuthorizationPolicyPtrType)(v)
}

func (*defaultAuthorizationPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultAuthorizationPolicy)(nil)).Elem()
}

func (i *defaultAuthorizationPolicyPtrType) ToDefaultAuthorizationPolicyPtrOutput() DefaultAuthorizationPolicyPtrOutput {
	return i.ToDefaultAuthorizationPolicyPtrOutputWithContext(context.Background())
}

func (i *defaultAuthorizationPolicyPtrType) ToDefaultAuthorizationPolicyPtrOutputWithContext(ctx context.Context) DefaultAuthorizationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultAuthorizationPolicyPtrOutput)
}

type DefaultAuthorizationPolicyOutput struct{ *pulumi.OutputState }

func (DefaultAuthorizationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultAuthorizationPolicy)(nil)).Elem()
}

func (o DefaultAuthorizationPolicyOutput) ToDefaultAuthorizationPolicyOutput() DefaultAuthorizationPolicyOutput {
	return o
}

func (o DefaultAuthorizationPolicyOutput) ToDefaultAuthorizationPolicyOutputWithContext(ctx context.Context) DefaultAuthorizationPolicyOutput {
	return o
}

func (o DefaultAuthorizationPolicyOutput) ToDefaultAuthorizationPolicyPtrOutput() DefaultAuthorizationPolicyPtrOutput {
	return o.ToDefaultAuthorizationPolicyPtrOutputWithContext(context.Background())
}

func (o DefaultAuthorizationPolicyOutput) ToDefaultAuthorizationPolicyPtrOutputWithContext(ctx context.Context) DefaultAuthorizationPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefaultAuthorizationPolicy) *DefaultAuthorizationPolicy {
		return &v
	}).(DefaultAuthorizationPolicyPtrOutput)
}

func (o DefaultAuthorizationPolicyOutput) AllowedApplications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DefaultAuthorizationPolicy) []string { return v.AllowedApplications }).(pulumi.StringArrayOutput)
}

func (o DefaultAuthorizationPolicyOutput) AllowedPrincipals() AllowedPrincipalsPtrOutput {
	return o.ApplyT(func(v DefaultAuthorizationPolicy) *AllowedPrincipals { return v.AllowedPrincipals }).(AllowedPrincipalsPtrOutput)
}

type DefaultAuthorizationPolicyPtrOutput struct{ *pulumi.OutputState }

func (DefaultAuthorizationPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultAuthorizationPolicy)(nil)).Elem()
}

func (o DefaultAuthorizationPolicyPtrOutput) ToDefaultAuthorizationPolicyPtrOutput() DefaultAuthorizationPolicyPtrOutput {
	return o
}

func (o DefaultAuthorizationPolicyPtrOutput) ToDefaultAuthorizationPolicyPtrOutputWithContext(ctx context.Context) DefaultAuthorizationPolicyPtrOutput {
	return o
}

func (o DefaultAuthorizationPolicyPtrOutput) Elem() DefaultAuthorizationPolicyOutput {
	return o.ApplyT(func(v *DefaultAuthorizationPolicy) DefaultAuthorizationPolicy {
		if v != nil {
			return *v
		}
		var ret DefaultAuthorizationPolicy
		return ret
	}).(DefaultAuthorizationPolicyOutput)
}

func (o DefaultAuthorizationPolicyPtrOutput) AllowedApplications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DefaultAuthorizationPolicy) []string {
		if v == nil {
			return nil
		}
		return v.AllowedApplications
	}).(pulumi.StringArrayOutput)
}

func (o DefaultAuthorizationPolicyPtrOutput) AllowedPrincipals() AllowedPrincipalsPtrOutput {
	return o.ApplyT(func(v *DefaultAuthorizationPolicy) *AllowedPrincipals {
		if v == nil {
			return nil
		}
		return v.AllowedPrincipals
	}).(AllowedPrincipalsPtrOutput)
}

type DefaultAuthorizationPolicyResponse struct {
	AllowedApplications []string                   `pulumi:"allowedApplications"`
	AllowedPrincipals   *AllowedPrincipalsResponse `pulumi:"allowedPrincipals"`
}

type DefaultAuthorizationPolicyResponseOutput struct{ *pulumi.OutputState }

func (DefaultAuthorizationPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultAuthorizationPolicyResponse)(nil)).Elem()
}

func (o DefaultAuthorizationPolicyResponseOutput) ToDefaultAuthorizationPolicyResponseOutput() DefaultAuthorizationPolicyResponseOutput {
	return o
}

func (o DefaultAuthorizationPolicyResponseOutput) ToDefaultAuthorizationPolicyResponseOutputWithContext(ctx context.Context) DefaultAuthorizationPolicyResponseOutput {
	return o
}

func (o DefaultAuthorizationPolicyResponseOutput) AllowedApplications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DefaultAuthorizationPolicyResponse) []string { return v.AllowedApplications }).(pulumi.StringArrayOutput)
}

func (o DefaultAuthorizationPolicyResponseOutput) AllowedPrincipals() AllowedPrincipalsResponsePtrOutput {
	return o.ApplyT(func(v DefaultAuthorizationPolicyResponse) *AllowedPrincipalsResponse { return v.AllowedPrincipals }).(AllowedPrincipalsResponsePtrOutput)
}

type DefaultAuthorizationPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (DefaultAuthorizationPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultAuthorizationPolicyResponse)(nil)).Elem()
}

func (o DefaultAuthorizationPolicyResponsePtrOutput) ToDefaultAuthorizationPolicyResponsePtrOutput() DefaultAuthorizationPolicyResponsePtrOutput {
	return o
}

func (o DefaultAuthorizationPolicyResponsePtrOutput) ToDefaultAuthorizationPolicyResponsePtrOutputWithContext(ctx context.Context) DefaultAuthorizationPolicyResponsePtrOutput {
	return o
}

func (o DefaultAuthorizationPolicyResponsePtrOutput) Elem() DefaultAuthorizationPolicyResponseOutput {
	return o.ApplyT(func(v *DefaultAuthorizationPolicyResponse) DefaultAuthorizationPolicyResponse {
		if v != nil {
			return *v
		}
		var ret DefaultAuthorizationPolicyResponse
		return ret
	}).(DefaultAuthorizationPolicyResponseOutput)
}

func (o DefaultAuthorizationPolicyResponsePtrOutput) AllowedApplications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DefaultAuthorizationPolicyResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedApplications
	}).(pulumi.StringArrayOutput)
}

func (o DefaultAuthorizationPolicyResponsePtrOutput) AllowedPrincipals() AllowedPrincipalsResponsePtrOutput {
	return o.ApplyT(func(v *DefaultAuthorizationPolicyResponse) *AllowedPrincipalsResponse {
		if v == nil {
			return nil
		}
		return v.AllowedPrincipals
	}).(AllowedPrincipalsResponsePtrOutput)
}

type EnabledConfig struct {
	Enabled *bool `pulumi:"enabled"`
}





type EnabledConfigInput interface {
	pulumi.Input

	ToEnabledConfigOutput() EnabledConfigOutput
	ToEnabledConfigOutputWithContext(context.Context) EnabledConfigOutput
}

type EnabledConfigArgs struct {
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (EnabledConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledConfig)(nil)).Elem()
}

func (i EnabledConfigArgs) ToEnabledConfigOutput() EnabledConfigOutput {
	return i.ToEnabledConfigOutputWithContext(context.Background())
}

func (i EnabledConfigArgs) ToEnabledConfigOutputWithContext(ctx context.Context) EnabledConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledConfigOutput)
}

func (i EnabledConfigArgs) ToEnabledConfigPtrOutput() EnabledConfigPtrOutput {
	return i.ToEnabledConfigPtrOutputWithContext(context.Background())
}

func (i EnabledConfigArgs) ToEnabledConfigPtrOutputWithContext(ctx context.Context) EnabledConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledConfigOutput).ToEnabledConfigPtrOutputWithContext(ctx)
}









type EnabledConfigPtrInput interface {
	pulumi.Input

	ToEnabledConfigPtrOutput() EnabledConfigPtrOutput
	ToEnabledConfigPtrOutputWithContext(context.Context) EnabledConfigPtrOutput
}

type enabledConfigPtrType EnabledConfigArgs

func EnabledConfigPtr(v *EnabledConfigArgs) EnabledConfigPtrInput {
	return (*enabledConfigPtrType)(v)
}

func (*enabledConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledConfig)(nil)).Elem()
}

func (i *enabledConfigPtrType) ToEnabledConfigPtrOutput() EnabledConfigPtrOutput {
	return i.ToEnabledConfigPtrOutputWithContext(context.Background())
}

func (i *enabledConfigPtrType) ToEnabledConfigPtrOutputWithContext(ctx context.Context) EnabledConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledConfigPtrOutput)
}

type EnabledConfigOutput struct{ *pulumi.OutputState }

func (EnabledConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledConfig)(nil)).Elem()
}

func (o EnabledConfigOutput) ToEnabledConfigOutput() EnabledConfigOutput {
	return o
}

func (o EnabledConfigOutput) ToEnabledConfigOutputWithContext(ctx context.Context) EnabledConfigOutput {
	return o
}

func (o EnabledConfigOutput) ToEnabledConfigPtrOutput() EnabledConfigPtrOutput {
	return o.ToEnabledConfigPtrOutputWithContext(context.Background())
}

func (o EnabledConfigOutput) ToEnabledConfigPtrOutputWithContext(ctx context.Context) EnabledConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnabledConfig) *EnabledConfig {
		return &v
	}).(EnabledConfigPtrOutput)
}

func (o EnabledConfigOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EnabledConfig) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type EnabledConfigPtrOutput struct{ *pulumi.OutputState }

func (EnabledConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledConfig)(nil)).Elem()
}

func (o EnabledConfigPtrOutput) ToEnabledConfigPtrOutput() EnabledConfigPtrOutput {
	return o
}

func (o EnabledConfigPtrOutput) ToEnabledConfigPtrOutputWithContext(ctx context.Context) EnabledConfigPtrOutput {
	return o
}

func (o EnabledConfigPtrOutput) Elem() EnabledConfigOutput {
	return o.ApplyT(func(v *EnabledConfig) EnabledConfig {
		if v != nil {
			return *v
		}
		var ret EnabledConfig
		return ret
	}).(EnabledConfigOutput)
}

func (o EnabledConfigPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledConfig) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type EnabledConfigResponse struct {
	Enabled *bool `pulumi:"enabled"`
}

type EnabledConfigResponseOutput struct{ *pulumi.OutputState }

func (EnabledConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledConfigResponse)(nil)).Elem()
}

func (o EnabledConfigResponseOutput) ToEnabledConfigResponseOutput() EnabledConfigResponseOutput {
	return o
}

func (o EnabledConfigResponseOutput) ToEnabledConfigResponseOutputWithContext(ctx context.Context) EnabledConfigResponseOutput {
	return o
}

func (o EnabledConfigResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EnabledConfigResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type EnabledConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (EnabledConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledConfigResponse)(nil)).Elem()
}

func (o EnabledConfigResponsePtrOutput) ToEnabledConfigResponsePtrOutput() EnabledConfigResponsePtrOutput {
	return o
}

func (o EnabledConfigResponsePtrOutput) ToEnabledConfigResponsePtrOutputWithContext(ctx context.Context) EnabledConfigResponsePtrOutput {
	return o
}

func (o EnabledConfigResponsePtrOutput) Elem() EnabledConfigResponseOutput {
	return o.ApplyT(func(v *EnabledConfigResponse) EnabledConfigResponse {
		if v != nil {
			return *v
		}
		var ret EnabledConfigResponse
		return ret
	}).(EnabledConfigResponseOutput)
}

func (o EnabledConfigResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type ErrorEntityResponse struct {
	Code            *string               `pulumi:"code"`
	Details         []ErrorEntityResponse `pulumi:"details"`
	ExtendedCode    *string               `pulumi:"extendedCode"`
	InnerErrors     []ErrorEntityResponse `pulumi:"innerErrors"`
	Message         *string               `pulumi:"message"`
	MessageTemplate *string               `pulumi:"messageTemplate"`
	Parameters      []string              `pulumi:"parameters"`
	Target          *string               `pulumi:"target"`
}

type ErrorEntityResponseOutput struct{ *pulumi.OutputState }

func (ErrorEntityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorEntityResponse)(nil)).Elem()
}

func (o ErrorEntityResponseOutput) ToErrorEntityResponseOutput() ErrorEntityResponseOutput {
	return o
}

func (o ErrorEntityResponseOutput) ToErrorEntityResponseOutputWithContext(ctx context.Context) ErrorEntityResponseOutput {
	return o
}

func (o ErrorEntityResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorEntityResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ErrorEntityResponseOutput) Details() ErrorEntityResponseArrayOutput {
	return o.ApplyT(func(v ErrorEntityResponse) []ErrorEntityResponse { return v.Details }).(ErrorEntityResponseArrayOutput)
}

func (o ErrorEntityResponseOutput) ExtendedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorEntityResponse) *string { return v.ExtendedCode }).(pulumi.StringPtrOutput)
}

func (o ErrorEntityResponseOutput) InnerErrors() ErrorEntityResponseArrayOutput {
	return o.ApplyT(func(v ErrorEntityResponse) []ErrorEntityResponse { return v.InnerErrors }).(ErrorEntityResponseArrayOutput)
}

func (o ErrorEntityResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorEntityResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ErrorEntityResponseOutput) MessageTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorEntityResponse) *string { return v.MessageTemplate }).(pulumi.StringPtrOutput)
}

func (o ErrorEntityResponseOutput) Parameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ErrorEntityResponse) []string { return v.Parameters }).(pulumi.StringArrayOutput)
}

func (o ErrorEntityResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorEntityResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type ErrorEntityResponsePtrOutput struct{ *pulumi.OutputState }

func (ErrorEntityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorEntityResponse)(nil)).Elem()
}

func (o ErrorEntityResponsePtrOutput) ToErrorEntityResponsePtrOutput() ErrorEntityResponsePtrOutput {
	return o
}

func (o ErrorEntityResponsePtrOutput) ToErrorEntityResponsePtrOutputWithContext(ctx context.Context) ErrorEntityResponsePtrOutput {
	return o
}

func (o ErrorEntityResponsePtrOutput) Elem() ErrorEntityResponseOutput {
	return o.ApplyT(func(v *ErrorEntityResponse) ErrorEntityResponse {
		if v != nil {
			return *v
		}
		var ret ErrorEntityResponse
		return ret
	}).(ErrorEntityResponseOutput)
}

func (o ErrorEntityResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ErrorEntityResponsePtrOutput) Details() ErrorEntityResponseArrayOutput {
	return o.ApplyT(func(v *ErrorEntityResponse) []ErrorEntityResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(ErrorEntityResponseArrayOutput)
}

func (o ErrorEntityResponsePtrOutput) ExtendedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExtendedCode
	}).(pulumi.StringPtrOutput)
}

func (o ErrorEntityResponsePtrOutput) InnerErrors() ErrorEntityResponseArrayOutput {
	return o.ApplyT(func(v *ErrorEntityResponse) []ErrorEntityResponse {
		if v == nil {
			return nil
		}
		return v.InnerErrors
	}).(ErrorEntityResponseArrayOutput)
}

func (o ErrorEntityResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ErrorEntityResponsePtrOutput) MessageTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.MessageTemplate
	}).(pulumi.StringPtrOutput)
}

func (o ErrorEntityResponsePtrOutput) Parameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ErrorEntityResponse) []string {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.StringArrayOutput)
}

func (o ErrorEntityResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

type ErrorEntityResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorEntityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorEntityResponse)(nil)).Elem()
}

func (o ErrorEntityResponseArrayOutput) ToErrorEntityResponseArrayOutput() ErrorEntityResponseArrayOutput {
	return o
}

func (o ErrorEntityResponseArrayOutput) ToErrorEntityResponseArrayOutputWithContext(ctx context.Context) ErrorEntityResponseArrayOutput {
	return o
}

func (o ErrorEntityResponseArrayOutput) Index(i pulumi.IntInput) ErrorEntityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorEntityResponse {
		return vs[0].([]ErrorEntityResponse)[vs[1].(int)]
	}).(ErrorEntityResponseOutput)
}

type Experiments struct {
	RampUpRules []RampUpRule `pulumi:"rampUpRules"`
}





type ExperimentsInput interface {
	pulumi.Input

	ToExperimentsOutput() ExperimentsOutput
	ToExperimentsOutputWithContext(context.Context) ExperimentsOutput
}

type ExperimentsArgs struct {
	RampUpRules RampUpRuleArrayInput `pulumi:"rampUpRules"`
}

func (ExperimentsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Experiments)(nil)).Elem()
}

func (i ExperimentsArgs) ToExperimentsOutput() ExperimentsOutput {
	return i.ToExperimentsOutputWithContext(context.Background())
}

func (i ExperimentsArgs) ToExperimentsOutputWithContext(ctx context.Context) ExperimentsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentsOutput)
}

func (i ExperimentsArgs) ToExperimentsPtrOutput() ExperimentsPtrOutput {
	return i.ToExperimentsPtrOutputWithContext(context.Background())
}

func (i ExperimentsArgs) ToExperimentsPtrOutputWithContext(ctx context.Context) ExperimentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentsOutput).ToExperimentsPtrOutputWithContext(ctx)
}









type ExperimentsPtrInput interface {
	pulumi.Input

	ToExperimentsPtrOutput() ExperimentsPtrOutput
	ToExperimentsPtrOutputWithContext(context.Context) ExperimentsPtrOutput
}

type experimentsPtrType ExperimentsArgs

func ExperimentsPtr(v *ExperimentsArgs) ExperimentsPtrInput {
	return (*experimentsPtrType)(v)
}

func (*experimentsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Experiments)(nil)).Elem()
}

func (i *experimentsPtrType) ToExperimentsPtrOutput() ExperimentsPtrOutput {
	return i.ToExperimentsPtrOutputWithContext(context.Background())
}

func (i *experimentsPtrType) ToExperimentsPtrOutputWithContext(ctx context.Context) ExperimentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentsPtrOutput)
}

type ExperimentsOutput struct{ *pulumi.OutputState }

func (ExperimentsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Experiments)(nil)).Elem()
}

func (o ExperimentsOutput) ToExperimentsOutput() ExperimentsOutput {
	return o
}

func (o ExperimentsOutput) ToExperimentsOutputWithContext(ctx context.Context) ExperimentsOutput {
	return o
}

func (o ExperimentsOutput) ToExperimentsPtrOutput() ExperimentsPtrOutput {
	return o.ToExperimentsPtrOutputWithContext(context.Background())
}

func (o ExperimentsOutput) ToExperimentsPtrOutputWithContext(ctx context.Context) ExperimentsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Experiments) *Experiments {
		return &v
	}).(ExperimentsPtrOutput)
}

func (o ExperimentsOutput) RampUpRules() RampUpRuleArrayOutput {
	return o.ApplyT(func(v Experiments) []RampUpRule { return v.RampUpRules }).(RampUpRuleArrayOutput)
}

type ExperimentsPtrOutput struct{ *pulumi.OutputState }

func (ExperimentsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Experiments)(nil)).Elem()
}

func (o ExperimentsPtrOutput) ToExperimentsPtrOutput() ExperimentsPtrOutput {
	return o
}

func (o ExperimentsPtrOutput) ToExperimentsPtrOutputWithContext(ctx context.Context) ExperimentsPtrOutput {
	return o
}

func (o ExperimentsPtrOutput) Elem() ExperimentsOutput {
	return o.ApplyT(func(v *Experiments) Experiments {
		if v != nil {
			return *v
		}
		var ret Experiments
		return ret
	}).(ExperimentsOutput)
}

func (o ExperimentsPtrOutput) RampUpRules() RampUpRuleArrayOutput {
	return o.ApplyT(func(v *Experiments) []RampUpRule {
		if v == nil {
			return nil
		}
		return v.RampUpRules
	}).(RampUpRuleArrayOutput)
}

type ExperimentsResponse struct {
	RampUpRules []RampUpRuleResponse `pulumi:"rampUpRules"`
}

type ExperimentsResponseOutput struct{ *pulumi.OutputState }

func (ExperimentsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExperimentsResponse)(nil)).Elem()
}

func (o ExperimentsResponseOutput) ToExperimentsResponseOutput() ExperimentsResponseOutput {
	return o
}

func (o ExperimentsResponseOutput) ToExperimentsResponseOutputWithContext(ctx context.Context) ExperimentsResponseOutput {
	return o
}

func (o ExperimentsResponseOutput) RampUpRules() RampUpRuleResponseArrayOutput {
	return o.ApplyT(func(v ExperimentsResponse) []RampUpRuleResponse { return v.RampUpRules }).(RampUpRuleResponseArrayOutput)
}

type ExperimentsResponsePtrOutput struct{ *pulumi.OutputState }

func (ExperimentsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExperimentsResponse)(nil)).Elem()
}

func (o ExperimentsResponsePtrOutput) ToExperimentsResponsePtrOutput() ExperimentsResponsePtrOutput {
	return o
}

func (o ExperimentsResponsePtrOutput) ToExperimentsResponsePtrOutputWithContext(ctx context.Context) ExperimentsResponsePtrOutput {
	return o
}

func (o ExperimentsResponsePtrOutput) Elem() ExperimentsResponseOutput {
	return o.ApplyT(func(v *ExperimentsResponse) ExperimentsResponse {
		if v != nil {
			return *v
		}
		var ret ExperimentsResponse
		return ret
	}).(ExperimentsResponseOutput)
}

func (o ExperimentsResponsePtrOutput) RampUpRules() RampUpRuleResponseArrayOutput {
	return o.ApplyT(func(v *ExperimentsResponse) []RampUpRuleResponse {
		if v == nil {
			return nil
		}
		return v.RampUpRules
	}).(RampUpRuleResponseArrayOutput)
}

type ExtendedLocation struct {
	Name *string `pulumi:"name"`
}





type ExtendedLocationInput interface {
	pulumi.Input

	ToExtendedLocationOutput() ExtendedLocationOutput
	ToExtendedLocationOutputWithContext(context.Context) ExtendedLocationOutput
}

type ExtendedLocationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ExtendedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (i ExtendedLocationArgs) ToExtendedLocationOutput() ExtendedLocationOutput {
	return i.ToExtendedLocationOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput)
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput).ToExtendedLocationPtrOutputWithContext(ctx)
}









type ExtendedLocationPtrInput interface {
	pulumi.Input

	ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput
	ToExtendedLocationPtrOutputWithContext(context.Context) ExtendedLocationPtrOutput
}

type extendedLocationPtrType ExtendedLocationArgs

func ExtendedLocationPtr(v *ExtendedLocationArgs) ExtendedLocationPtrInput {
	return (*extendedLocationPtrType)(v)
}

func (*extendedLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationPtrOutput)
}

type ExtendedLocationOutput struct{ *pulumi.OutputState }

func (ExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationOutput) ToExtendedLocationOutput() ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocation) *ExtendedLocation {
		return &v
	}).(ExtendedLocationPtrOutput)
}

func (o ExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) Elem() ExtendedLocationOutput {
	return o.ApplyT(func(v *ExtendedLocation) ExtendedLocation {
		if v != nil {
			return *v
		}
		var ret ExtendedLocation
		return ret
	}).(ExtendedLocationOutput)
}

func (o ExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponse struct {
	Name *string `pulumi:"name"`
	Type string  `pulumi:"type"`
}

type ExtendedLocationResponseOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutputWithContext(ctx context.Context) ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ExtendedLocationResponsePtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) Elem() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) ExtendedLocationResponse {
		if v != nil {
			return *v
		}
		var ret ExtendedLocationResponse
		return ret
	}).(ExtendedLocationResponseOutput)
}

func (o ExtendedLocationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type Facebook struct {
	Enabled         *bool            `pulumi:"enabled"`
	GraphApiVersion *string          `pulumi:"graphApiVersion"`
	Login           *LoginScopes     `pulumi:"login"`
	Registration    *AppRegistration `pulumi:"registration"`
}





type FacebookInput interface {
	pulumi.Input

	ToFacebookOutput() FacebookOutput
	ToFacebookOutputWithContext(context.Context) FacebookOutput
}

type FacebookArgs struct {
	Enabled         pulumi.BoolPtrInput     `pulumi:"enabled"`
	GraphApiVersion pulumi.StringPtrInput   `pulumi:"graphApiVersion"`
	Login           LoginScopesPtrInput     `pulumi:"login"`
	Registration    AppRegistrationPtrInput `pulumi:"registration"`
}

func (FacebookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Facebook)(nil)).Elem()
}

func (i FacebookArgs) ToFacebookOutput() FacebookOutput {
	return i.ToFacebookOutputWithContext(context.Background())
}

func (i FacebookArgs) ToFacebookOutputWithContext(ctx context.Context) FacebookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FacebookOutput)
}

func (i FacebookArgs) ToFacebookPtrOutput() FacebookPtrOutput {
	return i.ToFacebookPtrOutputWithContext(context.Background())
}

func (i FacebookArgs) ToFacebookPtrOutputWithContext(ctx context.Context) FacebookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FacebookOutput).ToFacebookPtrOutputWithContext(ctx)
}









type FacebookPtrInput interface {
	pulumi.Input

	ToFacebookPtrOutput() FacebookPtrOutput
	ToFacebookPtrOutputWithContext(context.Context) FacebookPtrOutput
}

type facebookPtrType FacebookArgs

func FacebookPtr(v *FacebookArgs) FacebookPtrInput {
	return (*facebookPtrType)(v)
}

func (*facebookPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Facebook)(nil)).Elem()
}

func (i *facebookPtrType) ToFacebookPtrOutput() FacebookPtrOutput {
	return i.ToFacebookPtrOutputWithContext(context.Background())
}

func (i *facebookPtrType) ToFacebookPtrOutputWithContext(ctx context.Context) FacebookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FacebookPtrOutput)
}

type FacebookOutput struct{ *pulumi.OutputState }

func (FacebookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Facebook)(nil)).Elem()
}

func (o FacebookOutput) ToFacebookOutput() FacebookOutput {
	return o
}

func (o FacebookOutput) ToFacebookOutputWithContext(ctx context.Context) FacebookOutput {
	return o
}

func (o FacebookOutput) ToFacebookPtrOutput() FacebookPtrOutput {
	return o.ToFacebookPtrOutputWithContext(context.Background())
}

func (o FacebookOutput) ToFacebookPtrOutputWithContext(ctx context.Context) FacebookPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Facebook) *Facebook {
		return &v
	}).(FacebookPtrOutput)
}

func (o FacebookOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Facebook) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o FacebookOutput) GraphApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Facebook) *string { return v.GraphApiVersion }).(pulumi.StringPtrOutput)
}

func (o FacebookOutput) Login() LoginScopesPtrOutput {
	return o.ApplyT(func(v Facebook) *LoginScopes { return v.Login }).(LoginScopesPtrOutput)
}

func (o FacebookOutput) Registration() AppRegistrationPtrOutput {
	return o.ApplyT(func(v Facebook) *AppRegistration { return v.Registration }).(AppRegistrationPtrOutput)
}

type FacebookPtrOutput struct{ *pulumi.OutputState }

func (FacebookPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Facebook)(nil)).Elem()
}

func (o FacebookPtrOutput) ToFacebookPtrOutput() FacebookPtrOutput {
	return o
}

func (o FacebookPtrOutput) ToFacebookPtrOutputWithContext(ctx context.Context) FacebookPtrOutput {
	return o
}

func (o FacebookPtrOutput) Elem() FacebookOutput {
	return o.ApplyT(func(v *Facebook) Facebook {
		if v != nil {
			return *v
		}
		var ret Facebook
		return ret
	}).(FacebookOutput)
}

func (o FacebookPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Facebook) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o FacebookPtrOutput) GraphApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Facebook) *string {
		if v == nil {
			return nil
		}
		return v.GraphApiVersion
	}).(pulumi.StringPtrOutput)
}

func (o FacebookPtrOutput) Login() LoginScopesPtrOutput {
	return o.ApplyT(func(v *Facebook) *LoginScopes {
		if v == nil {
			return nil
		}
		return v.Login
	}).(LoginScopesPtrOutput)
}

func (o FacebookPtrOutput) Registration() AppRegistrationPtrOutput {
	return o.ApplyT(func(v *Facebook) *AppRegistration {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(AppRegistrationPtrOutput)
}

type FacebookResponse struct {
	Enabled         *bool                    `pulumi:"enabled"`
	GraphApiVersion *string                  `pulumi:"graphApiVersion"`
	Login           *LoginScopesResponse     `pulumi:"login"`
	Registration    *AppRegistrationResponse `pulumi:"registration"`
}

type FacebookResponseOutput struct{ *pulumi.OutputState }

func (FacebookResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FacebookResponse)(nil)).Elem()
}

func (o FacebookResponseOutput) ToFacebookResponseOutput() FacebookResponseOutput {
	return o
}

func (o FacebookResponseOutput) ToFacebookResponseOutputWithContext(ctx context.Context) FacebookResponseOutput {
	return o
}

func (o FacebookResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FacebookResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o FacebookResponseOutput) GraphApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FacebookResponse) *string { return v.GraphApiVersion }).(pulumi.StringPtrOutput)
}

func (o FacebookResponseOutput) Login() LoginScopesResponsePtrOutput {
	return o.ApplyT(func(v FacebookResponse) *LoginScopesResponse { return v.Login }).(LoginScopesResponsePtrOutput)
}

func (o FacebookResponseOutput) Registration() AppRegistrationResponsePtrOutput {
	return o.ApplyT(func(v FacebookResponse) *AppRegistrationResponse { return v.Registration }).(AppRegistrationResponsePtrOutput)
}

type FacebookResponsePtrOutput struct{ *pulumi.OutputState }

func (FacebookResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FacebookResponse)(nil)).Elem()
}

func (o FacebookResponsePtrOutput) ToFacebookResponsePtrOutput() FacebookResponsePtrOutput {
	return o
}

func (o FacebookResponsePtrOutput) ToFacebookResponsePtrOutputWithContext(ctx context.Context) FacebookResponsePtrOutput {
	return o
}

func (o FacebookResponsePtrOutput) Elem() FacebookResponseOutput {
	return o.ApplyT(func(v *FacebookResponse) FacebookResponse {
		if v != nil {
			return *v
		}
		var ret FacebookResponse
		return ret
	}).(FacebookResponseOutput)
}

func (o FacebookResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FacebookResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o FacebookResponsePtrOutput) GraphApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FacebookResponse) *string {
		if v == nil {
			return nil
		}
		return v.GraphApiVersion
	}).(pulumi.StringPtrOutput)
}

func (o FacebookResponsePtrOutput) Login() LoginScopesResponsePtrOutput {
	return o.ApplyT(func(v *FacebookResponse) *LoginScopesResponse {
		if v == nil {
			return nil
		}
		return v.Login
	}).(LoginScopesResponsePtrOutput)
}

func (o FacebookResponsePtrOutput) Registration() AppRegistrationResponsePtrOutput {
	return o.ApplyT(func(v *FacebookResponse) *AppRegistrationResponse {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(AppRegistrationResponsePtrOutput)
}

type FileSystemApplicationLogsConfig struct {
	Level *LogLevel `pulumi:"level"`
}


func (val *FileSystemApplicationLogsConfig) Defaults() *FileSystemApplicationLogsConfig {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Level) {
		level_ := LogLevel("Off")
		tmp.Level = &level_
	}
	return &tmp
}





type FileSystemApplicationLogsConfigInput interface {
	pulumi.Input

	ToFileSystemApplicationLogsConfigOutput() FileSystemApplicationLogsConfigOutput
	ToFileSystemApplicationLogsConfigOutputWithContext(context.Context) FileSystemApplicationLogsConfigOutput
}

type FileSystemApplicationLogsConfigArgs struct {
	Level LogLevelPtrInput `pulumi:"level"`
}


func (val *FileSystemApplicationLogsConfigArgs) Defaults() *FileSystemApplicationLogsConfigArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Level) {
		tmp.Level = LogLevel("Off")
	}
	return &tmp
}
func (FileSystemApplicationLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemApplicationLogsConfig)(nil)).Elem()
}

func (i FileSystemApplicationLogsConfigArgs) ToFileSystemApplicationLogsConfigOutput() FileSystemApplicationLogsConfigOutput {
	return i.ToFileSystemApplicationLogsConfigOutputWithContext(context.Background())
}

func (i FileSystemApplicationLogsConfigArgs) ToFileSystemApplicationLogsConfigOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemApplicationLogsConfigOutput)
}

func (i FileSystemApplicationLogsConfigArgs) ToFileSystemApplicationLogsConfigPtrOutput() FileSystemApplicationLogsConfigPtrOutput {
	return i.ToFileSystemApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i FileSystemApplicationLogsConfigArgs) ToFileSystemApplicationLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemApplicationLogsConfigOutput).ToFileSystemApplicationLogsConfigPtrOutputWithContext(ctx)
}









type FileSystemApplicationLogsConfigPtrInput interface {
	pulumi.Input

	ToFileSystemApplicationLogsConfigPtrOutput() FileSystemApplicationLogsConfigPtrOutput
	ToFileSystemApplicationLogsConfigPtrOutputWithContext(context.Context) FileSystemApplicationLogsConfigPtrOutput
}

type fileSystemApplicationLogsConfigPtrType FileSystemApplicationLogsConfigArgs

func FileSystemApplicationLogsConfigPtr(v *FileSystemApplicationLogsConfigArgs) FileSystemApplicationLogsConfigPtrInput {
	return (*fileSystemApplicationLogsConfigPtrType)(v)
}

func (*fileSystemApplicationLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemApplicationLogsConfig)(nil)).Elem()
}

func (i *fileSystemApplicationLogsConfigPtrType) ToFileSystemApplicationLogsConfigPtrOutput() FileSystemApplicationLogsConfigPtrOutput {
	return i.ToFileSystemApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i *fileSystemApplicationLogsConfigPtrType) ToFileSystemApplicationLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemApplicationLogsConfigPtrOutput)
}

type FileSystemApplicationLogsConfigOutput struct{ *pulumi.OutputState }

func (FileSystemApplicationLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemApplicationLogsConfig)(nil)).Elem()
}

func (o FileSystemApplicationLogsConfigOutput) ToFileSystemApplicationLogsConfigOutput() FileSystemApplicationLogsConfigOutput {
	return o
}

func (o FileSystemApplicationLogsConfigOutput) ToFileSystemApplicationLogsConfigOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigOutput {
	return o
}

func (o FileSystemApplicationLogsConfigOutput) ToFileSystemApplicationLogsConfigPtrOutput() FileSystemApplicationLogsConfigPtrOutput {
	return o.ToFileSystemApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (o FileSystemApplicationLogsConfigOutput) ToFileSystemApplicationLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FileSystemApplicationLogsConfig) *FileSystemApplicationLogsConfig {
		return &v
	}).(FileSystemApplicationLogsConfigPtrOutput)
}

func (o FileSystemApplicationLogsConfigOutput) Level() LogLevelPtrOutput {
	return o.ApplyT(func(v FileSystemApplicationLogsConfig) *LogLevel { return v.Level }).(LogLevelPtrOutput)
}

type FileSystemApplicationLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (FileSystemApplicationLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemApplicationLogsConfig)(nil)).Elem()
}

func (o FileSystemApplicationLogsConfigPtrOutput) ToFileSystemApplicationLogsConfigPtrOutput() FileSystemApplicationLogsConfigPtrOutput {
	return o
}

func (o FileSystemApplicationLogsConfigPtrOutput) ToFileSystemApplicationLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigPtrOutput {
	return o
}

func (o FileSystemApplicationLogsConfigPtrOutput) Elem() FileSystemApplicationLogsConfigOutput {
	return o.ApplyT(func(v *FileSystemApplicationLogsConfig) FileSystemApplicationLogsConfig {
		if v != nil {
			return *v
		}
		var ret FileSystemApplicationLogsConfig
		return ret
	}).(FileSystemApplicationLogsConfigOutput)
}

func (o FileSystemApplicationLogsConfigPtrOutput) Level() LogLevelPtrOutput {
	return o.ApplyT(func(v *FileSystemApplicationLogsConfig) *LogLevel {
		if v == nil {
			return nil
		}
		return v.Level
	}).(LogLevelPtrOutput)
}

type FileSystemApplicationLogsConfigResponse struct {
	Level *string `pulumi:"level"`
}


func (val *FileSystemApplicationLogsConfigResponse) Defaults() *FileSystemApplicationLogsConfigResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Level) {
		level_ := "Off"
		tmp.Level = &level_
	}
	return &tmp
}

type FileSystemApplicationLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (FileSystemApplicationLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemApplicationLogsConfigResponse)(nil)).Elem()
}

func (o FileSystemApplicationLogsConfigResponseOutput) ToFileSystemApplicationLogsConfigResponseOutput() FileSystemApplicationLogsConfigResponseOutput {
	return o
}

func (o FileSystemApplicationLogsConfigResponseOutput) ToFileSystemApplicationLogsConfigResponseOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigResponseOutput {
	return o
}

func (o FileSystemApplicationLogsConfigResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileSystemApplicationLogsConfigResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

type FileSystemApplicationLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (FileSystemApplicationLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemApplicationLogsConfigResponse)(nil)).Elem()
}

func (o FileSystemApplicationLogsConfigResponsePtrOutput) ToFileSystemApplicationLogsConfigResponsePtrOutput() FileSystemApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o FileSystemApplicationLogsConfigResponsePtrOutput) ToFileSystemApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o FileSystemApplicationLogsConfigResponsePtrOutput) Elem() FileSystemApplicationLogsConfigResponseOutput {
	return o.ApplyT(func(v *FileSystemApplicationLogsConfigResponse) FileSystemApplicationLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret FileSystemApplicationLogsConfigResponse
		return ret
	}).(FileSystemApplicationLogsConfigResponseOutput)
}

func (o FileSystemApplicationLogsConfigResponsePtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileSystemApplicationLogsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Level
	}).(pulumi.StringPtrOutput)
}

type FileSystemHttpLogsConfig struct {
	Enabled         *bool `pulumi:"enabled"`
	RetentionInDays *int  `pulumi:"retentionInDays"`
	RetentionInMb   *int  `pulumi:"retentionInMb"`
}





type FileSystemHttpLogsConfigInput interface {
	pulumi.Input

	ToFileSystemHttpLogsConfigOutput() FileSystemHttpLogsConfigOutput
	ToFileSystemHttpLogsConfigOutputWithContext(context.Context) FileSystemHttpLogsConfigOutput
}

type FileSystemHttpLogsConfigArgs struct {
	Enabled         pulumi.BoolPtrInput `pulumi:"enabled"`
	RetentionInDays pulumi.IntPtrInput  `pulumi:"retentionInDays"`
	RetentionInMb   pulumi.IntPtrInput  `pulumi:"retentionInMb"`
}

func (FileSystemHttpLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemHttpLogsConfig)(nil)).Elem()
}

func (i FileSystemHttpLogsConfigArgs) ToFileSystemHttpLogsConfigOutput() FileSystemHttpLogsConfigOutput {
	return i.ToFileSystemHttpLogsConfigOutputWithContext(context.Background())
}

func (i FileSystemHttpLogsConfigArgs) ToFileSystemHttpLogsConfigOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemHttpLogsConfigOutput)
}

func (i FileSystemHttpLogsConfigArgs) ToFileSystemHttpLogsConfigPtrOutput() FileSystemHttpLogsConfigPtrOutput {
	return i.ToFileSystemHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (i FileSystemHttpLogsConfigArgs) ToFileSystemHttpLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemHttpLogsConfigOutput).ToFileSystemHttpLogsConfigPtrOutputWithContext(ctx)
}









type FileSystemHttpLogsConfigPtrInput interface {
	pulumi.Input

	ToFileSystemHttpLogsConfigPtrOutput() FileSystemHttpLogsConfigPtrOutput
	ToFileSystemHttpLogsConfigPtrOutputWithContext(context.Context) FileSystemHttpLogsConfigPtrOutput
}

type fileSystemHttpLogsConfigPtrType FileSystemHttpLogsConfigArgs

func FileSystemHttpLogsConfigPtr(v *FileSystemHttpLogsConfigArgs) FileSystemHttpLogsConfigPtrInput {
	return (*fileSystemHttpLogsConfigPtrType)(v)
}

func (*fileSystemHttpLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemHttpLogsConfig)(nil)).Elem()
}

func (i *fileSystemHttpLogsConfigPtrType) ToFileSystemHttpLogsConfigPtrOutput() FileSystemHttpLogsConfigPtrOutput {
	return i.ToFileSystemHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (i *fileSystemHttpLogsConfigPtrType) ToFileSystemHttpLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemHttpLogsConfigPtrOutput)
}

type FileSystemHttpLogsConfigOutput struct{ *pulumi.OutputState }

func (FileSystemHttpLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemHttpLogsConfig)(nil)).Elem()
}

func (o FileSystemHttpLogsConfigOutput) ToFileSystemHttpLogsConfigOutput() FileSystemHttpLogsConfigOutput {
	return o
}

func (o FileSystemHttpLogsConfigOutput) ToFileSystemHttpLogsConfigOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigOutput {
	return o
}

func (o FileSystemHttpLogsConfigOutput) ToFileSystemHttpLogsConfigPtrOutput() FileSystemHttpLogsConfigPtrOutput {
	return o.ToFileSystemHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (o FileSystemHttpLogsConfigOutput) ToFileSystemHttpLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FileSystemHttpLogsConfig) *FileSystemHttpLogsConfig {
		return &v
	}).(FileSystemHttpLogsConfigPtrOutput)
}

func (o FileSystemHttpLogsConfigOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FileSystemHttpLogsConfig) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o FileSystemHttpLogsConfigOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FileSystemHttpLogsConfig) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o FileSystemHttpLogsConfigOutput) RetentionInMb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FileSystemHttpLogsConfig) *int { return v.RetentionInMb }).(pulumi.IntPtrOutput)
}

type FileSystemHttpLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (FileSystemHttpLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemHttpLogsConfig)(nil)).Elem()
}

func (o FileSystemHttpLogsConfigPtrOutput) ToFileSystemHttpLogsConfigPtrOutput() FileSystemHttpLogsConfigPtrOutput {
	return o
}

func (o FileSystemHttpLogsConfigPtrOutput) ToFileSystemHttpLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigPtrOutput {
	return o
}

func (o FileSystemHttpLogsConfigPtrOutput) Elem() FileSystemHttpLogsConfigOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfig) FileSystemHttpLogsConfig {
		if v != nil {
			return *v
		}
		var ret FileSystemHttpLogsConfig
		return ret
	}).(FileSystemHttpLogsConfigOutput)
}

func (o FileSystemHttpLogsConfigPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfig) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o FileSystemHttpLogsConfigPtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfig) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o FileSystemHttpLogsConfigPtrOutput) RetentionInMb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfig) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInMb
	}).(pulumi.IntPtrOutput)
}

type FileSystemHttpLogsConfigResponse struct {
	Enabled         *bool `pulumi:"enabled"`
	RetentionInDays *int  `pulumi:"retentionInDays"`
	RetentionInMb   *int  `pulumi:"retentionInMb"`
}

type FileSystemHttpLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (FileSystemHttpLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemHttpLogsConfigResponse)(nil)).Elem()
}

func (o FileSystemHttpLogsConfigResponseOutput) ToFileSystemHttpLogsConfigResponseOutput() FileSystemHttpLogsConfigResponseOutput {
	return o
}

func (o FileSystemHttpLogsConfigResponseOutput) ToFileSystemHttpLogsConfigResponseOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigResponseOutput {
	return o
}

func (o FileSystemHttpLogsConfigResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FileSystemHttpLogsConfigResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o FileSystemHttpLogsConfigResponseOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FileSystemHttpLogsConfigResponse) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o FileSystemHttpLogsConfigResponseOutput) RetentionInMb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FileSystemHttpLogsConfigResponse) *int { return v.RetentionInMb }).(pulumi.IntPtrOutput)
}

type FileSystemHttpLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (FileSystemHttpLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemHttpLogsConfigResponse)(nil)).Elem()
}

func (o FileSystemHttpLogsConfigResponsePtrOutput) ToFileSystemHttpLogsConfigResponsePtrOutput() FileSystemHttpLogsConfigResponsePtrOutput {
	return o
}

func (o FileSystemHttpLogsConfigResponsePtrOutput) ToFileSystemHttpLogsConfigResponsePtrOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigResponsePtrOutput {
	return o
}

func (o FileSystemHttpLogsConfigResponsePtrOutput) Elem() FileSystemHttpLogsConfigResponseOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfigResponse) FileSystemHttpLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret FileSystemHttpLogsConfigResponse
		return ret
	}).(FileSystemHttpLogsConfigResponseOutput)
}

func (o FileSystemHttpLogsConfigResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o FileSystemHttpLogsConfigResponsePtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o FileSystemHttpLogsConfigResponsePtrOutput) RetentionInMb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInMb
	}).(pulumi.IntPtrOutput)
}

type FileSystemTokenStore struct {
	Directory *string `pulumi:"directory"`
}





type FileSystemTokenStoreInput interface {
	pulumi.Input

	ToFileSystemTokenStoreOutput() FileSystemTokenStoreOutput
	ToFileSystemTokenStoreOutputWithContext(context.Context) FileSystemTokenStoreOutput
}

type FileSystemTokenStoreArgs struct {
	Directory pulumi.StringPtrInput `pulumi:"directory"`
}

func (FileSystemTokenStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemTokenStore)(nil)).Elem()
}

func (i FileSystemTokenStoreArgs) ToFileSystemTokenStoreOutput() FileSystemTokenStoreOutput {
	return i.ToFileSystemTokenStoreOutputWithContext(context.Background())
}

func (i FileSystemTokenStoreArgs) ToFileSystemTokenStoreOutputWithContext(ctx context.Context) FileSystemTokenStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemTokenStoreOutput)
}

func (i FileSystemTokenStoreArgs) ToFileSystemTokenStorePtrOutput() FileSystemTokenStorePtrOutput {
	return i.ToFileSystemTokenStorePtrOutputWithContext(context.Background())
}

func (i FileSystemTokenStoreArgs) ToFileSystemTokenStorePtrOutputWithContext(ctx context.Context) FileSystemTokenStorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemTokenStoreOutput).ToFileSystemTokenStorePtrOutputWithContext(ctx)
}









type FileSystemTokenStorePtrInput interface {
	pulumi.Input

	ToFileSystemTokenStorePtrOutput() FileSystemTokenStorePtrOutput
	ToFileSystemTokenStorePtrOutputWithContext(context.Context) FileSystemTokenStorePtrOutput
}

type fileSystemTokenStorePtrType FileSystemTokenStoreArgs

func FileSystemTokenStorePtr(v *FileSystemTokenStoreArgs) FileSystemTokenStorePtrInput {
	return (*fileSystemTokenStorePtrType)(v)
}

func (*fileSystemTokenStorePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemTokenStore)(nil)).Elem()
}

func (i *fileSystemTokenStorePtrType) ToFileSystemTokenStorePtrOutput() FileSystemTokenStorePtrOutput {
	return i.ToFileSystemTokenStorePtrOutputWithContext(context.Background())
}

func (i *fileSystemTokenStorePtrType) ToFileSystemTokenStorePtrOutputWithContext(ctx context.Context) FileSystemTokenStorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemTokenStorePtrOutput)
}

type FileSystemTokenStoreOutput struct{ *pulumi.OutputState }

func (FileSystemTokenStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemTokenStore)(nil)).Elem()
}

func (o FileSystemTokenStoreOutput) ToFileSystemTokenStoreOutput() FileSystemTokenStoreOutput {
	return o
}

func (o FileSystemTokenStoreOutput) ToFileSystemTokenStoreOutputWithContext(ctx context.Context) FileSystemTokenStoreOutput {
	return o
}

func (o FileSystemTokenStoreOutput) ToFileSystemTokenStorePtrOutput() FileSystemTokenStorePtrOutput {
	return o.ToFileSystemTokenStorePtrOutputWithContext(context.Background())
}

func (o FileSystemTokenStoreOutput) ToFileSystemTokenStorePtrOutputWithContext(ctx context.Context) FileSystemTokenStorePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FileSystemTokenStore) *FileSystemTokenStore {
		return &v
	}).(FileSystemTokenStorePtrOutput)
}

func (o FileSystemTokenStoreOutput) Directory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileSystemTokenStore) *string { return v.Directory }).(pulumi.StringPtrOutput)
}

type FileSystemTokenStorePtrOutput struct{ *pulumi.OutputState }

func (FileSystemTokenStorePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemTokenStore)(nil)).Elem()
}

func (o FileSystemTokenStorePtrOutput) ToFileSystemTokenStorePtrOutput() FileSystemTokenStorePtrOutput {
	return o
}

func (o FileSystemTokenStorePtrOutput) ToFileSystemTokenStorePtrOutputWithContext(ctx context.Context) FileSystemTokenStorePtrOutput {
	return o
}

func (o FileSystemTokenStorePtrOutput) Elem() FileSystemTokenStoreOutput {
	return o.ApplyT(func(v *FileSystemTokenStore) FileSystemTokenStore {
		if v != nil {
			return *v
		}
		var ret FileSystemTokenStore
		return ret
	}).(FileSystemTokenStoreOutput)
}

func (o FileSystemTokenStorePtrOutput) Directory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileSystemTokenStore) *string {
		if v == nil {
			return nil
		}
		return v.Directory
	}).(pulumi.StringPtrOutput)
}

type FileSystemTokenStoreResponse struct {
	Directory *string `pulumi:"directory"`
}

type FileSystemTokenStoreResponseOutput struct{ *pulumi.OutputState }

func (FileSystemTokenStoreResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemTokenStoreResponse)(nil)).Elem()
}

func (o FileSystemTokenStoreResponseOutput) ToFileSystemTokenStoreResponseOutput() FileSystemTokenStoreResponseOutput {
	return o
}

func (o FileSystemTokenStoreResponseOutput) ToFileSystemTokenStoreResponseOutputWithContext(ctx context.Context) FileSystemTokenStoreResponseOutput {
	return o
}

func (o FileSystemTokenStoreResponseOutput) Directory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileSystemTokenStoreResponse) *string { return v.Directory }).(pulumi.StringPtrOutput)
}

type FileSystemTokenStoreResponsePtrOutput struct{ *pulumi.OutputState }

func (FileSystemTokenStoreResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemTokenStoreResponse)(nil)).Elem()
}

func (o FileSystemTokenStoreResponsePtrOutput) ToFileSystemTokenStoreResponsePtrOutput() FileSystemTokenStoreResponsePtrOutput {
	return o
}

func (o FileSystemTokenStoreResponsePtrOutput) ToFileSystemTokenStoreResponsePtrOutputWithContext(ctx context.Context) FileSystemTokenStoreResponsePtrOutput {
	return o
}

func (o FileSystemTokenStoreResponsePtrOutput) Elem() FileSystemTokenStoreResponseOutput {
	return o.ApplyT(func(v *FileSystemTokenStoreResponse) FileSystemTokenStoreResponse {
		if v != nil {
			return *v
		}
		var ret FileSystemTokenStoreResponse
		return ret
	}).(FileSystemTokenStoreResponseOutput)
}

func (o FileSystemTokenStoreResponsePtrOutput) Directory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileSystemTokenStoreResponse) *string {
		if v == nil {
			return nil
		}
		return v.Directory
	}).(pulumi.StringPtrOutput)
}

type ForwardProxy struct {
	Convention            *ForwardProxyConvention `pulumi:"convention"`
	CustomHostHeaderName  *string                 `pulumi:"customHostHeaderName"`
	CustomProtoHeaderName *string                 `pulumi:"customProtoHeaderName"`
}





type ForwardProxyInput interface {
	pulumi.Input

	ToForwardProxyOutput() ForwardProxyOutput
	ToForwardProxyOutputWithContext(context.Context) ForwardProxyOutput
}

type ForwardProxyArgs struct {
	Convention            ForwardProxyConventionPtrInput `pulumi:"convention"`
	CustomHostHeaderName  pulumi.StringPtrInput          `pulumi:"customHostHeaderName"`
	CustomProtoHeaderName pulumi.StringPtrInput          `pulumi:"customProtoHeaderName"`
}

func (ForwardProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ForwardProxy)(nil)).Elem()
}

func (i ForwardProxyArgs) ToForwardProxyOutput() ForwardProxyOutput {
	return i.ToForwardProxyOutputWithContext(context.Background())
}

func (i ForwardProxyArgs) ToForwardProxyOutputWithContext(ctx context.Context) ForwardProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardProxyOutput)
}

func (i ForwardProxyArgs) ToForwardProxyPtrOutput() ForwardProxyPtrOutput {
	return i.ToForwardProxyPtrOutputWithContext(context.Background())
}

func (i ForwardProxyArgs) ToForwardProxyPtrOutputWithContext(ctx context.Context) ForwardProxyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardProxyOutput).ToForwardProxyPtrOutputWithContext(ctx)
}









type ForwardProxyPtrInput interface {
	pulumi.Input

	ToForwardProxyPtrOutput() ForwardProxyPtrOutput
	ToForwardProxyPtrOutputWithContext(context.Context) ForwardProxyPtrOutput
}

type forwardProxyPtrType ForwardProxyArgs

func ForwardProxyPtr(v *ForwardProxyArgs) ForwardProxyPtrInput {
	return (*forwardProxyPtrType)(v)
}

func (*forwardProxyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ForwardProxy)(nil)).Elem()
}

func (i *forwardProxyPtrType) ToForwardProxyPtrOutput() ForwardProxyPtrOutput {
	return i.ToForwardProxyPtrOutputWithContext(context.Background())
}

func (i *forwardProxyPtrType) ToForwardProxyPtrOutputWithContext(ctx context.Context) ForwardProxyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ForwardProxyPtrOutput)
}

type ForwardProxyOutput struct{ *pulumi.OutputState }

func (ForwardProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ForwardProxy)(nil)).Elem()
}

func (o ForwardProxyOutput) ToForwardProxyOutput() ForwardProxyOutput {
	return o
}

func (o ForwardProxyOutput) ToForwardProxyOutputWithContext(ctx context.Context) ForwardProxyOutput {
	return o
}

func (o ForwardProxyOutput) ToForwardProxyPtrOutput() ForwardProxyPtrOutput {
	return o.ToForwardProxyPtrOutputWithContext(context.Background())
}

func (o ForwardProxyOutput) ToForwardProxyPtrOutputWithContext(ctx context.Context) ForwardProxyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ForwardProxy) *ForwardProxy {
		return &v
	}).(ForwardProxyPtrOutput)
}

func (o ForwardProxyOutput) Convention() ForwardProxyConventionPtrOutput {
	return o.ApplyT(func(v ForwardProxy) *ForwardProxyConvention { return v.Convention }).(ForwardProxyConventionPtrOutput)
}

func (o ForwardProxyOutput) CustomHostHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForwardProxy) *string { return v.CustomHostHeaderName }).(pulumi.StringPtrOutput)
}

func (o ForwardProxyOutput) CustomProtoHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForwardProxy) *string { return v.CustomProtoHeaderName }).(pulumi.StringPtrOutput)
}

type ForwardProxyPtrOutput struct{ *pulumi.OutputState }

func (ForwardProxyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ForwardProxy)(nil)).Elem()
}

func (o ForwardProxyPtrOutput) ToForwardProxyPtrOutput() ForwardProxyPtrOutput {
	return o
}

func (o ForwardProxyPtrOutput) ToForwardProxyPtrOutputWithContext(ctx context.Context) ForwardProxyPtrOutput {
	return o
}

func (o ForwardProxyPtrOutput) Elem() ForwardProxyOutput {
	return o.ApplyT(func(v *ForwardProxy) ForwardProxy {
		if v != nil {
			return *v
		}
		var ret ForwardProxy
		return ret
	}).(ForwardProxyOutput)
}

func (o ForwardProxyPtrOutput) Convention() ForwardProxyConventionPtrOutput {
	return o.ApplyT(func(v *ForwardProxy) *ForwardProxyConvention {
		if v == nil {
			return nil
		}
		return v.Convention
	}).(ForwardProxyConventionPtrOutput)
}

func (o ForwardProxyPtrOutput) CustomHostHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ForwardProxy) *string {
		if v == nil {
			return nil
		}
		return v.CustomHostHeaderName
	}).(pulumi.StringPtrOutput)
}

func (o ForwardProxyPtrOutput) CustomProtoHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ForwardProxy) *string {
		if v == nil {
			return nil
		}
		return v.CustomProtoHeaderName
	}).(pulumi.StringPtrOutput)
}

type ForwardProxyResponse struct {
	Convention            *string `pulumi:"convention"`
	CustomHostHeaderName  *string `pulumi:"customHostHeaderName"`
	CustomProtoHeaderName *string `pulumi:"customProtoHeaderName"`
}

type ForwardProxyResponseOutput struct{ *pulumi.OutputState }

func (ForwardProxyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ForwardProxyResponse)(nil)).Elem()
}

func (o ForwardProxyResponseOutput) ToForwardProxyResponseOutput() ForwardProxyResponseOutput {
	return o
}

func (o ForwardProxyResponseOutput) ToForwardProxyResponseOutputWithContext(ctx context.Context) ForwardProxyResponseOutput {
	return o
}

func (o ForwardProxyResponseOutput) Convention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForwardProxyResponse) *string { return v.Convention }).(pulumi.StringPtrOutput)
}

func (o ForwardProxyResponseOutput) CustomHostHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForwardProxyResponse) *string { return v.CustomHostHeaderName }).(pulumi.StringPtrOutput)
}

func (o ForwardProxyResponseOutput) CustomProtoHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ForwardProxyResponse) *string { return v.CustomProtoHeaderName }).(pulumi.StringPtrOutput)
}

type ForwardProxyResponsePtrOutput struct{ *pulumi.OutputState }

func (ForwardProxyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ForwardProxyResponse)(nil)).Elem()
}

func (o ForwardProxyResponsePtrOutput) ToForwardProxyResponsePtrOutput() ForwardProxyResponsePtrOutput {
	return o
}

func (o ForwardProxyResponsePtrOutput) ToForwardProxyResponsePtrOutputWithContext(ctx context.Context) ForwardProxyResponsePtrOutput {
	return o
}

func (o ForwardProxyResponsePtrOutput) Elem() ForwardProxyResponseOutput {
	return o.ApplyT(func(v *ForwardProxyResponse) ForwardProxyResponse {
		if v != nil {
			return *v
		}
		var ret ForwardProxyResponse
		return ret
	}).(ForwardProxyResponseOutput)
}

func (o ForwardProxyResponsePtrOutput) Convention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ForwardProxyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Convention
	}).(pulumi.StringPtrOutput)
}

func (o ForwardProxyResponsePtrOutput) CustomHostHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ForwardProxyResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomHostHeaderName
	}).(pulumi.StringPtrOutput)
}

func (o ForwardProxyResponsePtrOutput) CustomProtoHeaderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ForwardProxyResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomProtoHeaderName
	}).(pulumi.StringPtrOutput)
}

type FrontEndConfiguration struct {
	Kind *FrontEndServiceType `pulumi:"kind"`
}





type FrontEndConfigurationInput interface {
	pulumi.Input

	ToFrontEndConfigurationOutput() FrontEndConfigurationOutput
	ToFrontEndConfigurationOutputWithContext(context.Context) FrontEndConfigurationOutput
}

type FrontEndConfigurationArgs struct {
	Kind FrontEndServiceTypePtrInput `pulumi:"kind"`
}

func (FrontEndConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontEndConfiguration)(nil)).Elem()
}

func (i FrontEndConfigurationArgs) ToFrontEndConfigurationOutput() FrontEndConfigurationOutput {
	return i.ToFrontEndConfigurationOutputWithContext(context.Background())
}

func (i FrontEndConfigurationArgs) ToFrontEndConfigurationOutputWithContext(ctx context.Context) FrontEndConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontEndConfigurationOutput)
}

func (i FrontEndConfigurationArgs) ToFrontEndConfigurationPtrOutput() FrontEndConfigurationPtrOutput {
	return i.ToFrontEndConfigurationPtrOutputWithContext(context.Background())
}

func (i FrontEndConfigurationArgs) ToFrontEndConfigurationPtrOutputWithContext(ctx context.Context) FrontEndConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontEndConfigurationOutput).ToFrontEndConfigurationPtrOutputWithContext(ctx)
}









type FrontEndConfigurationPtrInput interface {
	pulumi.Input

	ToFrontEndConfigurationPtrOutput() FrontEndConfigurationPtrOutput
	ToFrontEndConfigurationPtrOutputWithContext(context.Context) FrontEndConfigurationPtrOutput
}

type frontEndConfigurationPtrType FrontEndConfigurationArgs

func FrontEndConfigurationPtr(v *FrontEndConfigurationArgs) FrontEndConfigurationPtrInput {
	return (*frontEndConfigurationPtrType)(v)
}

func (*frontEndConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontEndConfiguration)(nil)).Elem()
}

func (i *frontEndConfigurationPtrType) ToFrontEndConfigurationPtrOutput() FrontEndConfigurationPtrOutput {
	return i.ToFrontEndConfigurationPtrOutputWithContext(context.Background())
}

func (i *frontEndConfigurationPtrType) ToFrontEndConfigurationPtrOutputWithContext(ctx context.Context) FrontEndConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontEndConfigurationPtrOutput)
}

type FrontEndConfigurationOutput struct{ *pulumi.OutputState }

func (FrontEndConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontEndConfiguration)(nil)).Elem()
}

func (o FrontEndConfigurationOutput) ToFrontEndConfigurationOutput() FrontEndConfigurationOutput {
	return o
}

func (o FrontEndConfigurationOutput) ToFrontEndConfigurationOutputWithContext(ctx context.Context) FrontEndConfigurationOutput {
	return o
}

func (o FrontEndConfigurationOutput) ToFrontEndConfigurationPtrOutput() FrontEndConfigurationPtrOutput {
	return o.ToFrontEndConfigurationPtrOutputWithContext(context.Background())
}

func (o FrontEndConfigurationOutput) ToFrontEndConfigurationPtrOutputWithContext(ctx context.Context) FrontEndConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FrontEndConfiguration) *FrontEndConfiguration {
		return &v
	}).(FrontEndConfigurationPtrOutput)
}

func (o FrontEndConfigurationOutput) Kind() FrontEndServiceTypePtrOutput {
	return o.ApplyT(func(v FrontEndConfiguration) *FrontEndServiceType { return v.Kind }).(FrontEndServiceTypePtrOutput)
}

type FrontEndConfigurationPtrOutput struct{ *pulumi.OutputState }

func (FrontEndConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontEndConfiguration)(nil)).Elem()
}

func (o FrontEndConfigurationPtrOutput) ToFrontEndConfigurationPtrOutput() FrontEndConfigurationPtrOutput {
	return o
}

func (o FrontEndConfigurationPtrOutput) ToFrontEndConfigurationPtrOutputWithContext(ctx context.Context) FrontEndConfigurationPtrOutput {
	return o
}

func (o FrontEndConfigurationPtrOutput) Elem() FrontEndConfigurationOutput {
	return o.ApplyT(func(v *FrontEndConfiguration) FrontEndConfiguration {
		if v != nil {
			return *v
		}
		var ret FrontEndConfiguration
		return ret
	}).(FrontEndConfigurationOutput)
}

func (o FrontEndConfigurationPtrOutput) Kind() FrontEndServiceTypePtrOutput {
	return o.ApplyT(func(v *FrontEndConfiguration) *FrontEndServiceType {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(FrontEndServiceTypePtrOutput)
}

type FrontEndConfigurationResponse struct {
	Kind *string `pulumi:"kind"`
}

type FrontEndConfigurationResponseOutput struct{ *pulumi.OutputState }

func (FrontEndConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontEndConfigurationResponse)(nil)).Elem()
}

func (o FrontEndConfigurationResponseOutput) ToFrontEndConfigurationResponseOutput() FrontEndConfigurationResponseOutput {
	return o
}

func (o FrontEndConfigurationResponseOutput) ToFrontEndConfigurationResponseOutputWithContext(ctx context.Context) FrontEndConfigurationResponseOutput {
	return o
}

func (o FrontEndConfigurationResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontEndConfigurationResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

type FrontEndConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (FrontEndConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontEndConfigurationResponse)(nil)).Elem()
}

func (o FrontEndConfigurationResponsePtrOutput) ToFrontEndConfigurationResponsePtrOutput() FrontEndConfigurationResponsePtrOutput {
	return o
}

func (o FrontEndConfigurationResponsePtrOutput) ToFrontEndConfigurationResponsePtrOutputWithContext(ctx context.Context) FrontEndConfigurationResponsePtrOutput {
	return o
}

func (o FrontEndConfigurationResponsePtrOutput) Elem() FrontEndConfigurationResponseOutput {
	return o.ApplyT(func(v *FrontEndConfigurationResponse) FrontEndConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret FrontEndConfigurationResponse
		return ret
	}).(FrontEndConfigurationResponseOutput)
}

func (o FrontEndConfigurationResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontEndConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

type GitHub struct {
	Enabled      *bool               `pulumi:"enabled"`
	Login        *LoginScopes        `pulumi:"login"`
	Registration *ClientRegistration `pulumi:"registration"`
}





type GitHubInput interface {
	pulumi.Input

	ToGitHubOutput() GitHubOutput
	ToGitHubOutputWithContext(context.Context) GitHubOutput
}

type GitHubArgs struct {
	Enabled      pulumi.BoolPtrInput        `pulumi:"enabled"`
	Login        LoginScopesPtrInput        `pulumi:"login"`
	Registration ClientRegistrationPtrInput `pulumi:"registration"`
}

func (GitHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHub)(nil)).Elem()
}

func (i GitHubArgs) ToGitHubOutput() GitHubOutput {
	return i.ToGitHubOutputWithContext(context.Background())
}

func (i GitHubArgs) ToGitHubOutputWithContext(ctx context.Context) GitHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubOutput)
}

func (i GitHubArgs) ToGitHubPtrOutput() GitHubPtrOutput {
	return i.ToGitHubPtrOutputWithContext(context.Background())
}

func (i GitHubArgs) ToGitHubPtrOutputWithContext(ctx context.Context) GitHubPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubOutput).ToGitHubPtrOutputWithContext(ctx)
}









type GitHubPtrInput interface {
	pulumi.Input

	ToGitHubPtrOutput() GitHubPtrOutput
	ToGitHubPtrOutputWithContext(context.Context) GitHubPtrOutput
}

type gitHubPtrType GitHubArgs

func GitHubPtr(v *GitHubArgs) GitHubPtrInput {
	return (*gitHubPtrType)(v)
}

func (*gitHubPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHub)(nil)).Elem()
}

func (i *gitHubPtrType) ToGitHubPtrOutput() GitHubPtrOutput {
	return i.ToGitHubPtrOutputWithContext(context.Background())
}

func (i *gitHubPtrType) ToGitHubPtrOutputWithContext(ctx context.Context) GitHubPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubPtrOutput)
}

type GitHubOutput struct{ *pulumi.OutputState }

func (GitHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHub)(nil)).Elem()
}

func (o GitHubOutput) ToGitHubOutput() GitHubOutput {
	return o
}

func (o GitHubOutput) ToGitHubOutputWithContext(ctx context.Context) GitHubOutput {
	return o
}

func (o GitHubOutput) ToGitHubPtrOutput() GitHubPtrOutput {
	return o.ToGitHubPtrOutputWithContext(context.Background())
}

func (o GitHubOutput) ToGitHubPtrOutputWithContext(ctx context.Context) GitHubPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitHub) *GitHub {
		return &v
	}).(GitHubPtrOutput)
}

func (o GitHubOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GitHub) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o GitHubOutput) Login() LoginScopesPtrOutput {
	return o.ApplyT(func(v GitHub) *LoginScopes { return v.Login }).(LoginScopesPtrOutput)
}

func (o GitHubOutput) Registration() ClientRegistrationPtrOutput {
	return o.ApplyT(func(v GitHub) *ClientRegistration { return v.Registration }).(ClientRegistrationPtrOutput)
}

type GitHubPtrOutput struct{ *pulumi.OutputState }

func (GitHubPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHub)(nil)).Elem()
}

func (o GitHubPtrOutput) ToGitHubPtrOutput() GitHubPtrOutput {
	return o
}

func (o GitHubPtrOutput) ToGitHubPtrOutputWithContext(ctx context.Context) GitHubPtrOutput {
	return o
}

func (o GitHubPtrOutput) Elem() GitHubOutput {
	return o.ApplyT(func(v *GitHub) GitHub {
		if v != nil {
			return *v
		}
		var ret GitHub
		return ret
	}).(GitHubOutput)
}

func (o GitHubPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GitHub) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o GitHubPtrOutput) Login() LoginScopesPtrOutput {
	return o.ApplyT(func(v *GitHub) *LoginScopes {
		if v == nil {
			return nil
		}
		return v.Login
	}).(LoginScopesPtrOutput)
}

func (o GitHubPtrOutput) Registration() ClientRegistrationPtrOutput {
	return o.ApplyT(func(v *GitHub) *ClientRegistration {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(ClientRegistrationPtrOutput)
}

type GitHubActionCodeConfiguration struct {
	RuntimeStack   *string `pulumi:"runtimeStack"`
	RuntimeVersion *string `pulumi:"runtimeVersion"`
}





type GitHubActionCodeConfigurationInput interface {
	pulumi.Input

	ToGitHubActionCodeConfigurationOutput() GitHubActionCodeConfigurationOutput
	ToGitHubActionCodeConfigurationOutputWithContext(context.Context) GitHubActionCodeConfigurationOutput
}

type GitHubActionCodeConfigurationArgs struct {
	RuntimeStack   pulumi.StringPtrInput `pulumi:"runtimeStack"`
	RuntimeVersion pulumi.StringPtrInput `pulumi:"runtimeVersion"`
}

func (GitHubActionCodeConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubActionCodeConfiguration)(nil)).Elem()
}

func (i GitHubActionCodeConfigurationArgs) ToGitHubActionCodeConfigurationOutput() GitHubActionCodeConfigurationOutput {
	return i.ToGitHubActionCodeConfigurationOutputWithContext(context.Background())
}

func (i GitHubActionCodeConfigurationArgs) ToGitHubActionCodeConfigurationOutputWithContext(ctx context.Context) GitHubActionCodeConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubActionCodeConfigurationOutput)
}

func (i GitHubActionCodeConfigurationArgs) ToGitHubActionCodeConfigurationPtrOutput() GitHubActionCodeConfigurationPtrOutput {
	return i.ToGitHubActionCodeConfigurationPtrOutputWithContext(context.Background())
}

func (i GitHubActionCodeConfigurationArgs) ToGitHubActionCodeConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionCodeConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubActionCodeConfigurationOutput).ToGitHubActionCodeConfigurationPtrOutputWithContext(ctx)
}









type GitHubActionCodeConfigurationPtrInput interface {
	pulumi.Input

	ToGitHubActionCodeConfigurationPtrOutput() GitHubActionCodeConfigurationPtrOutput
	ToGitHubActionCodeConfigurationPtrOutputWithContext(context.Context) GitHubActionCodeConfigurationPtrOutput
}

type gitHubActionCodeConfigurationPtrType GitHubActionCodeConfigurationArgs

func GitHubActionCodeConfigurationPtr(v *GitHubActionCodeConfigurationArgs) GitHubActionCodeConfigurationPtrInput {
	return (*gitHubActionCodeConfigurationPtrType)(v)
}

func (*gitHubActionCodeConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubActionCodeConfiguration)(nil)).Elem()
}

func (i *gitHubActionCodeConfigurationPtrType) ToGitHubActionCodeConfigurationPtrOutput() GitHubActionCodeConfigurationPtrOutput {
	return i.ToGitHubActionCodeConfigurationPtrOutputWithContext(context.Background())
}

func (i *gitHubActionCodeConfigurationPtrType) ToGitHubActionCodeConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionCodeConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubActionCodeConfigurationPtrOutput)
}

type GitHubActionCodeConfigurationOutput struct{ *pulumi.OutputState }

func (GitHubActionCodeConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubActionCodeConfiguration)(nil)).Elem()
}

func (o GitHubActionCodeConfigurationOutput) ToGitHubActionCodeConfigurationOutput() GitHubActionCodeConfigurationOutput {
	return o
}

func (o GitHubActionCodeConfigurationOutput) ToGitHubActionCodeConfigurationOutputWithContext(ctx context.Context) GitHubActionCodeConfigurationOutput {
	return o
}

func (o GitHubActionCodeConfigurationOutput) ToGitHubActionCodeConfigurationPtrOutput() GitHubActionCodeConfigurationPtrOutput {
	return o.ToGitHubActionCodeConfigurationPtrOutputWithContext(context.Background())
}

func (o GitHubActionCodeConfigurationOutput) ToGitHubActionCodeConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionCodeConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitHubActionCodeConfiguration) *GitHubActionCodeConfiguration {
		return &v
	}).(GitHubActionCodeConfigurationPtrOutput)
}

func (o GitHubActionCodeConfigurationOutput) RuntimeStack() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionCodeConfiguration) *string { return v.RuntimeStack }).(pulumi.StringPtrOutput)
}

func (o GitHubActionCodeConfigurationOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionCodeConfiguration) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

type GitHubActionCodeConfigurationPtrOutput struct{ *pulumi.OutputState }

func (GitHubActionCodeConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubActionCodeConfiguration)(nil)).Elem()
}

func (o GitHubActionCodeConfigurationPtrOutput) ToGitHubActionCodeConfigurationPtrOutput() GitHubActionCodeConfigurationPtrOutput {
	return o
}

func (o GitHubActionCodeConfigurationPtrOutput) ToGitHubActionCodeConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionCodeConfigurationPtrOutput {
	return o
}

func (o GitHubActionCodeConfigurationPtrOutput) Elem() GitHubActionCodeConfigurationOutput {
	return o.ApplyT(func(v *GitHubActionCodeConfiguration) GitHubActionCodeConfiguration {
		if v != nil {
			return *v
		}
		var ret GitHubActionCodeConfiguration
		return ret
	}).(GitHubActionCodeConfigurationOutput)
}

func (o GitHubActionCodeConfigurationPtrOutput) RuntimeStack() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionCodeConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeStack
	}).(pulumi.StringPtrOutput)
}

func (o GitHubActionCodeConfigurationPtrOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionCodeConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeVersion
	}).(pulumi.StringPtrOutput)
}

type GitHubActionCodeConfigurationResponse struct {
	RuntimeStack   *string `pulumi:"runtimeStack"`
	RuntimeVersion *string `pulumi:"runtimeVersion"`
}

type GitHubActionCodeConfigurationResponseOutput struct{ *pulumi.OutputState }

func (GitHubActionCodeConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubActionCodeConfigurationResponse)(nil)).Elem()
}

func (o GitHubActionCodeConfigurationResponseOutput) ToGitHubActionCodeConfigurationResponseOutput() GitHubActionCodeConfigurationResponseOutput {
	return o
}

func (o GitHubActionCodeConfigurationResponseOutput) ToGitHubActionCodeConfigurationResponseOutputWithContext(ctx context.Context) GitHubActionCodeConfigurationResponseOutput {
	return o
}

func (o GitHubActionCodeConfigurationResponseOutput) RuntimeStack() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionCodeConfigurationResponse) *string { return v.RuntimeStack }).(pulumi.StringPtrOutput)
}

func (o GitHubActionCodeConfigurationResponseOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionCodeConfigurationResponse) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

type GitHubActionCodeConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (GitHubActionCodeConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubActionCodeConfigurationResponse)(nil)).Elem()
}

func (o GitHubActionCodeConfigurationResponsePtrOutput) ToGitHubActionCodeConfigurationResponsePtrOutput() GitHubActionCodeConfigurationResponsePtrOutput {
	return o
}

func (o GitHubActionCodeConfigurationResponsePtrOutput) ToGitHubActionCodeConfigurationResponsePtrOutputWithContext(ctx context.Context) GitHubActionCodeConfigurationResponsePtrOutput {
	return o
}

func (o GitHubActionCodeConfigurationResponsePtrOutput) Elem() GitHubActionCodeConfigurationResponseOutput {
	return o.ApplyT(func(v *GitHubActionCodeConfigurationResponse) GitHubActionCodeConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret GitHubActionCodeConfigurationResponse
		return ret
	}).(GitHubActionCodeConfigurationResponseOutput)
}

func (o GitHubActionCodeConfigurationResponsePtrOutput) RuntimeStack() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionCodeConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeStack
	}).(pulumi.StringPtrOutput)
}

func (o GitHubActionCodeConfigurationResponsePtrOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionCodeConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeVersion
	}).(pulumi.StringPtrOutput)
}

type GitHubActionConfiguration struct {
	CodeConfiguration      *GitHubActionCodeConfiguration      `pulumi:"codeConfiguration"`
	ContainerConfiguration *GitHubActionContainerConfiguration `pulumi:"containerConfiguration"`
	GenerateWorkflowFile   *bool                               `pulumi:"generateWorkflowFile"`
	IsLinux                *bool                               `pulumi:"isLinux"`
}





type GitHubActionConfigurationInput interface {
	pulumi.Input

	ToGitHubActionConfigurationOutput() GitHubActionConfigurationOutput
	ToGitHubActionConfigurationOutputWithContext(context.Context) GitHubActionConfigurationOutput
}

type GitHubActionConfigurationArgs struct {
	CodeConfiguration      GitHubActionCodeConfigurationPtrInput      `pulumi:"codeConfiguration"`
	ContainerConfiguration GitHubActionContainerConfigurationPtrInput `pulumi:"containerConfiguration"`
	GenerateWorkflowFile   pulumi.BoolPtrInput                        `pulumi:"generateWorkflowFile"`
	IsLinux                pulumi.BoolPtrInput                        `pulumi:"isLinux"`
}

func (GitHubActionConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubActionConfiguration)(nil)).Elem()
}

func (i GitHubActionConfigurationArgs) ToGitHubActionConfigurationOutput() GitHubActionConfigurationOutput {
	return i.ToGitHubActionConfigurationOutputWithContext(context.Background())
}

func (i GitHubActionConfigurationArgs) ToGitHubActionConfigurationOutputWithContext(ctx context.Context) GitHubActionConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubActionConfigurationOutput)
}

func (i GitHubActionConfigurationArgs) ToGitHubActionConfigurationPtrOutput() GitHubActionConfigurationPtrOutput {
	return i.ToGitHubActionConfigurationPtrOutputWithContext(context.Background())
}

func (i GitHubActionConfigurationArgs) ToGitHubActionConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubActionConfigurationOutput).ToGitHubActionConfigurationPtrOutputWithContext(ctx)
}









type GitHubActionConfigurationPtrInput interface {
	pulumi.Input

	ToGitHubActionConfigurationPtrOutput() GitHubActionConfigurationPtrOutput
	ToGitHubActionConfigurationPtrOutputWithContext(context.Context) GitHubActionConfigurationPtrOutput
}

type gitHubActionConfigurationPtrType GitHubActionConfigurationArgs

func GitHubActionConfigurationPtr(v *GitHubActionConfigurationArgs) GitHubActionConfigurationPtrInput {
	return (*gitHubActionConfigurationPtrType)(v)
}

func (*gitHubActionConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubActionConfiguration)(nil)).Elem()
}

func (i *gitHubActionConfigurationPtrType) ToGitHubActionConfigurationPtrOutput() GitHubActionConfigurationPtrOutput {
	return i.ToGitHubActionConfigurationPtrOutputWithContext(context.Background())
}

func (i *gitHubActionConfigurationPtrType) ToGitHubActionConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubActionConfigurationPtrOutput)
}

type GitHubActionConfigurationOutput struct{ *pulumi.OutputState }

func (GitHubActionConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubActionConfiguration)(nil)).Elem()
}

func (o GitHubActionConfigurationOutput) ToGitHubActionConfigurationOutput() GitHubActionConfigurationOutput {
	return o
}

func (o GitHubActionConfigurationOutput) ToGitHubActionConfigurationOutputWithContext(ctx context.Context) GitHubActionConfigurationOutput {
	return o
}

func (o GitHubActionConfigurationOutput) ToGitHubActionConfigurationPtrOutput() GitHubActionConfigurationPtrOutput {
	return o.ToGitHubActionConfigurationPtrOutputWithContext(context.Background())
}

func (o GitHubActionConfigurationOutput) ToGitHubActionConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitHubActionConfiguration) *GitHubActionConfiguration {
		return &v
	}).(GitHubActionConfigurationPtrOutput)
}

func (o GitHubActionConfigurationOutput) CodeConfiguration() GitHubActionCodeConfigurationPtrOutput {
	return o.ApplyT(func(v GitHubActionConfiguration) *GitHubActionCodeConfiguration { return v.CodeConfiguration }).(GitHubActionCodeConfigurationPtrOutput)
}

func (o GitHubActionConfigurationOutput) ContainerConfiguration() GitHubActionContainerConfigurationPtrOutput {
	return o.ApplyT(func(v GitHubActionConfiguration) *GitHubActionContainerConfiguration { return v.ContainerConfiguration }).(GitHubActionContainerConfigurationPtrOutput)
}

func (o GitHubActionConfigurationOutput) GenerateWorkflowFile() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GitHubActionConfiguration) *bool { return v.GenerateWorkflowFile }).(pulumi.BoolPtrOutput)
}

func (o GitHubActionConfigurationOutput) IsLinux() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GitHubActionConfiguration) *bool { return v.IsLinux }).(pulumi.BoolPtrOutput)
}

type GitHubActionConfigurationPtrOutput struct{ *pulumi.OutputState }

func (GitHubActionConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubActionConfiguration)(nil)).Elem()
}

func (o GitHubActionConfigurationPtrOutput) ToGitHubActionConfigurationPtrOutput() GitHubActionConfigurationPtrOutput {
	return o
}

func (o GitHubActionConfigurationPtrOutput) ToGitHubActionConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionConfigurationPtrOutput {
	return o
}

func (o GitHubActionConfigurationPtrOutput) Elem() GitHubActionConfigurationOutput {
	return o.ApplyT(func(v *GitHubActionConfiguration) GitHubActionConfiguration {
		if v != nil {
			return *v
		}
		var ret GitHubActionConfiguration
		return ret
	}).(GitHubActionConfigurationOutput)
}

func (o GitHubActionConfigurationPtrOutput) CodeConfiguration() GitHubActionCodeConfigurationPtrOutput {
	return o.ApplyT(func(v *GitHubActionConfiguration) *GitHubActionCodeConfiguration {
		if v == nil {
			return nil
		}
		return v.CodeConfiguration
	}).(GitHubActionCodeConfigurationPtrOutput)
}

func (o GitHubActionConfigurationPtrOutput) ContainerConfiguration() GitHubActionContainerConfigurationPtrOutput {
	return o.ApplyT(func(v *GitHubActionConfiguration) *GitHubActionContainerConfiguration {
		if v == nil {
			return nil
		}
		return v.ContainerConfiguration
	}).(GitHubActionContainerConfigurationPtrOutput)
}

func (o GitHubActionConfigurationPtrOutput) GenerateWorkflowFile() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GitHubActionConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.GenerateWorkflowFile
	}).(pulumi.BoolPtrOutput)
}

func (o GitHubActionConfigurationPtrOutput) IsLinux() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GitHubActionConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.IsLinux
	}).(pulumi.BoolPtrOutput)
}

type GitHubActionConfigurationResponse struct {
	CodeConfiguration      *GitHubActionCodeConfigurationResponse      `pulumi:"codeConfiguration"`
	ContainerConfiguration *GitHubActionContainerConfigurationResponse `pulumi:"containerConfiguration"`
	GenerateWorkflowFile   *bool                                       `pulumi:"generateWorkflowFile"`
	IsLinux                *bool                                       `pulumi:"isLinux"`
}

type GitHubActionConfigurationResponseOutput struct{ *pulumi.OutputState }

func (GitHubActionConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubActionConfigurationResponse)(nil)).Elem()
}

func (o GitHubActionConfigurationResponseOutput) ToGitHubActionConfigurationResponseOutput() GitHubActionConfigurationResponseOutput {
	return o
}

func (o GitHubActionConfigurationResponseOutput) ToGitHubActionConfigurationResponseOutputWithContext(ctx context.Context) GitHubActionConfigurationResponseOutput {
	return o
}

func (o GitHubActionConfigurationResponseOutput) CodeConfiguration() GitHubActionCodeConfigurationResponsePtrOutput {
	return o.ApplyT(func(v GitHubActionConfigurationResponse) *GitHubActionCodeConfigurationResponse {
		return v.CodeConfiguration
	}).(GitHubActionCodeConfigurationResponsePtrOutput)
}

func (o GitHubActionConfigurationResponseOutput) ContainerConfiguration() GitHubActionContainerConfigurationResponsePtrOutput {
	return o.ApplyT(func(v GitHubActionConfigurationResponse) *GitHubActionContainerConfigurationResponse {
		return v.ContainerConfiguration
	}).(GitHubActionContainerConfigurationResponsePtrOutput)
}

func (o GitHubActionConfigurationResponseOutput) GenerateWorkflowFile() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GitHubActionConfigurationResponse) *bool { return v.GenerateWorkflowFile }).(pulumi.BoolPtrOutput)
}

func (o GitHubActionConfigurationResponseOutput) IsLinux() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GitHubActionConfigurationResponse) *bool { return v.IsLinux }).(pulumi.BoolPtrOutput)
}

type GitHubActionConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (GitHubActionConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubActionConfigurationResponse)(nil)).Elem()
}

func (o GitHubActionConfigurationResponsePtrOutput) ToGitHubActionConfigurationResponsePtrOutput() GitHubActionConfigurationResponsePtrOutput {
	return o
}

func (o GitHubActionConfigurationResponsePtrOutput) ToGitHubActionConfigurationResponsePtrOutputWithContext(ctx context.Context) GitHubActionConfigurationResponsePtrOutput {
	return o
}

func (o GitHubActionConfigurationResponsePtrOutput) Elem() GitHubActionConfigurationResponseOutput {
	return o.ApplyT(func(v *GitHubActionConfigurationResponse) GitHubActionConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret GitHubActionConfigurationResponse
		return ret
	}).(GitHubActionConfigurationResponseOutput)
}

func (o GitHubActionConfigurationResponsePtrOutput) CodeConfiguration() GitHubActionCodeConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *GitHubActionConfigurationResponse) *GitHubActionCodeConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.CodeConfiguration
	}).(GitHubActionCodeConfigurationResponsePtrOutput)
}

func (o GitHubActionConfigurationResponsePtrOutput) ContainerConfiguration() GitHubActionContainerConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *GitHubActionConfigurationResponse) *GitHubActionContainerConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.ContainerConfiguration
	}).(GitHubActionContainerConfigurationResponsePtrOutput)
}

func (o GitHubActionConfigurationResponsePtrOutput) GenerateWorkflowFile() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GitHubActionConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.GenerateWorkflowFile
	}).(pulumi.BoolPtrOutput)
}

func (o GitHubActionConfigurationResponsePtrOutput) IsLinux() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GitHubActionConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsLinux
	}).(pulumi.BoolPtrOutput)
}

type GitHubActionContainerConfiguration struct {
	ImageName *string `pulumi:"imageName"`
	Password  *string `pulumi:"password"`
	ServerUrl *string `pulumi:"serverUrl"`
	Username  *string `pulumi:"username"`
}





type GitHubActionContainerConfigurationInput interface {
	pulumi.Input

	ToGitHubActionContainerConfigurationOutput() GitHubActionContainerConfigurationOutput
	ToGitHubActionContainerConfigurationOutputWithContext(context.Context) GitHubActionContainerConfigurationOutput
}

type GitHubActionContainerConfigurationArgs struct {
	ImageName pulumi.StringPtrInput `pulumi:"imageName"`
	Password  pulumi.StringPtrInput `pulumi:"password"`
	ServerUrl pulumi.StringPtrInput `pulumi:"serverUrl"`
	Username  pulumi.StringPtrInput `pulumi:"username"`
}

func (GitHubActionContainerConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubActionContainerConfiguration)(nil)).Elem()
}

func (i GitHubActionContainerConfigurationArgs) ToGitHubActionContainerConfigurationOutput() GitHubActionContainerConfigurationOutput {
	return i.ToGitHubActionContainerConfigurationOutputWithContext(context.Background())
}

func (i GitHubActionContainerConfigurationArgs) ToGitHubActionContainerConfigurationOutputWithContext(ctx context.Context) GitHubActionContainerConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubActionContainerConfigurationOutput)
}

func (i GitHubActionContainerConfigurationArgs) ToGitHubActionContainerConfigurationPtrOutput() GitHubActionContainerConfigurationPtrOutput {
	return i.ToGitHubActionContainerConfigurationPtrOutputWithContext(context.Background())
}

func (i GitHubActionContainerConfigurationArgs) ToGitHubActionContainerConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionContainerConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubActionContainerConfigurationOutput).ToGitHubActionContainerConfigurationPtrOutputWithContext(ctx)
}









type GitHubActionContainerConfigurationPtrInput interface {
	pulumi.Input

	ToGitHubActionContainerConfigurationPtrOutput() GitHubActionContainerConfigurationPtrOutput
	ToGitHubActionContainerConfigurationPtrOutputWithContext(context.Context) GitHubActionContainerConfigurationPtrOutput
}

type gitHubActionContainerConfigurationPtrType GitHubActionContainerConfigurationArgs

func GitHubActionContainerConfigurationPtr(v *GitHubActionContainerConfigurationArgs) GitHubActionContainerConfigurationPtrInput {
	return (*gitHubActionContainerConfigurationPtrType)(v)
}

func (*gitHubActionContainerConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubActionContainerConfiguration)(nil)).Elem()
}

func (i *gitHubActionContainerConfigurationPtrType) ToGitHubActionContainerConfigurationPtrOutput() GitHubActionContainerConfigurationPtrOutput {
	return i.ToGitHubActionContainerConfigurationPtrOutputWithContext(context.Background())
}

func (i *gitHubActionContainerConfigurationPtrType) ToGitHubActionContainerConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionContainerConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubActionContainerConfigurationPtrOutput)
}

type GitHubActionContainerConfigurationOutput struct{ *pulumi.OutputState }

func (GitHubActionContainerConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubActionContainerConfiguration)(nil)).Elem()
}

func (o GitHubActionContainerConfigurationOutput) ToGitHubActionContainerConfigurationOutput() GitHubActionContainerConfigurationOutput {
	return o
}

func (o GitHubActionContainerConfigurationOutput) ToGitHubActionContainerConfigurationOutputWithContext(ctx context.Context) GitHubActionContainerConfigurationOutput {
	return o
}

func (o GitHubActionContainerConfigurationOutput) ToGitHubActionContainerConfigurationPtrOutput() GitHubActionContainerConfigurationPtrOutput {
	return o.ToGitHubActionContainerConfigurationPtrOutputWithContext(context.Background())
}

func (o GitHubActionContainerConfigurationOutput) ToGitHubActionContainerConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionContainerConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitHubActionContainerConfiguration) *GitHubActionContainerConfiguration {
		return &v
	}).(GitHubActionContainerConfigurationPtrOutput)
}

func (o GitHubActionContainerConfigurationOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionContainerConfiguration) *string { return v.ImageName }).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionContainerConfiguration) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationOutput) ServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionContainerConfiguration) *string { return v.ServerUrl }).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionContainerConfiguration) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type GitHubActionContainerConfigurationPtrOutput struct{ *pulumi.OutputState }

func (GitHubActionContainerConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubActionContainerConfiguration)(nil)).Elem()
}

func (o GitHubActionContainerConfigurationPtrOutput) ToGitHubActionContainerConfigurationPtrOutput() GitHubActionContainerConfigurationPtrOutput {
	return o
}

func (o GitHubActionContainerConfigurationPtrOutput) ToGitHubActionContainerConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionContainerConfigurationPtrOutput {
	return o
}

func (o GitHubActionContainerConfigurationPtrOutput) Elem() GitHubActionContainerConfigurationOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfiguration) GitHubActionContainerConfiguration {
		if v != nil {
			return *v
		}
		var ret GitHubActionContainerConfiguration
		return ret
	}).(GitHubActionContainerConfigurationOutput)
}

func (o GitHubActionContainerConfigurationPtrOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ImageName
	}).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationPtrOutput) ServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ServerUrl
	}).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type GitHubActionContainerConfigurationResponse struct {
	ImageName *string `pulumi:"imageName"`
	Password  *string `pulumi:"password"`
	ServerUrl *string `pulumi:"serverUrl"`
	Username  *string `pulumi:"username"`
}

type GitHubActionContainerConfigurationResponseOutput struct{ *pulumi.OutputState }

func (GitHubActionContainerConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubActionContainerConfigurationResponse)(nil)).Elem()
}

func (o GitHubActionContainerConfigurationResponseOutput) ToGitHubActionContainerConfigurationResponseOutput() GitHubActionContainerConfigurationResponseOutput {
	return o
}

func (o GitHubActionContainerConfigurationResponseOutput) ToGitHubActionContainerConfigurationResponseOutputWithContext(ctx context.Context) GitHubActionContainerConfigurationResponseOutput {
	return o
}

func (o GitHubActionContainerConfigurationResponseOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionContainerConfigurationResponse) *string { return v.ImageName }).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionContainerConfigurationResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationResponseOutput) ServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionContainerConfigurationResponse) *string { return v.ServerUrl }).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionContainerConfigurationResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type GitHubActionContainerConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (GitHubActionContainerConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubActionContainerConfigurationResponse)(nil)).Elem()
}

func (o GitHubActionContainerConfigurationResponsePtrOutput) ToGitHubActionContainerConfigurationResponsePtrOutput() GitHubActionContainerConfigurationResponsePtrOutput {
	return o
}

func (o GitHubActionContainerConfigurationResponsePtrOutput) ToGitHubActionContainerConfigurationResponsePtrOutputWithContext(ctx context.Context) GitHubActionContainerConfigurationResponsePtrOutput {
	return o
}

func (o GitHubActionContainerConfigurationResponsePtrOutput) Elem() GitHubActionContainerConfigurationResponseOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfigurationResponse) GitHubActionContainerConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret GitHubActionContainerConfigurationResponse
		return ret
	}).(GitHubActionContainerConfigurationResponseOutput)
}

func (o GitHubActionContainerConfigurationResponsePtrOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ImageName
	}).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationResponsePtrOutput) ServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServerUrl
	}).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type GitHubResponse struct {
	Enabled      *bool                       `pulumi:"enabled"`
	Login        *LoginScopesResponse        `pulumi:"login"`
	Registration *ClientRegistrationResponse `pulumi:"registration"`
}

type GitHubResponseOutput struct{ *pulumi.OutputState }

func (GitHubResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubResponse)(nil)).Elem()
}

func (o GitHubResponseOutput) ToGitHubResponseOutput() GitHubResponseOutput {
	return o
}

func (o GitHubResponseOutput) ToGitHubResponseOutputWithContext(ctx context.Context) GitHubResponseOutput {
	return o
}

func (o GitHubResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GitHubResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o GitHubResponseOutput) Login() LoginScopesResponsePtrOutput {
	return o.ApplyT(func(v GitHubResponse) *LoginScopesResponse { return v.Login }).(LoginScopesResponsePtrOutput)
}

func (o GitHubResponseOutput) Registration() ClientRegistrationResponsePtrOutput {
	return o.ApplyT(func(v GitHubResponse) *ClientRegistrationResponse { return v.Registration }).(ClientRegistrationResponsePtrOutput)
}

type GitHubResponsePtrOutput struct{ *pulumi.OutputState }

func (GitHubResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubResponse)(nil)).Elem()
}

func (o GitHubResponsePtrOutput) ToGitHubResponsePtrOutput() GitHubResponsePtrOutput {
	return o
}

func (o GitHubResponsePtrOutput) ToGitHubResponsePtrOutputWithContext(ctx context.Context) GitHubResponsePtrOutput {
	return o
}

func (o GitHubResponsePtrOutput) Elem() GitHubResponseOutput {
	return o.ApplyT(func(v *GitHubResponse) GitHubResponse {
		if v != nil {
			return *v
		}
		var ret GitHubResponse
		return ret
	}).(GitHubResponseOutput)
}

func (o GitHubResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GitHubResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o GitHubResponsePtrOutput) Login() LoginScopesResponsePtrOutput {
	return o.ApplyT(func(v *GitHubResponse) *LoginScopesResponse {
		if v == nil {
			return nil
		}
		return v.Login
	}).(LoginScopesResponsePtrOutput)
}

func (o GitHubResponsePtrOutput) Registration() ClientRegistrationResponsePtrOutput {
	return o.ApplyT(func(v *GitHubResponse) *ClientRegistrationResponse {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(ClientRegistrationResponsePtrOutput)
}

type GlobalValidation struct {
	ExcludedPaths               []string                       `pulumi:"excludedPaths"`
	RedirectToProvider          *string                        `pulumi:"redirectToProvider"`
	RequireAuthentication       *bool                          `pulumi:"requireAuthentication"`
	UnauthenticatedClientAction *UnauthenticatedClientActionV2 `pulumi:"unauthenticatedClientAction"`
}





type GlobalValidationInput interface {
	pulumi.Input

	ToGlobalValidationOutput() GlobalValidationOutput
	ToGlobalValidationOutputWithContext(context.Context) GlobalValidationOutput
}

type GlobalValidationArgs struct {
	ExcludedPaths               pulumi.StringArrayInput               `pulumi:"excludedPaths"`
	RedirectToProvider          pulumi.StringPtrInput                 `pulumi:"redirectToProvider"`
	RequireAuthentication       pulumi.BoolPtrInput                   `pulumi:"requireAuthentication"`
	UnauthenticatedClientAction UnauthenticatedClientActionV2PtrInput `pulumi:"unauthenticatedClientAction"`
}

func (GlobalValidationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalValidation)(nil)).Elem()
}

func (i GlobalValidationArgs) ToGlobalValidationOutput() GlobalValidationOutput {
	return i.ToGlobalValidationOutputWithContext(context.Background())
}

func (i GlobalValidationArgs) ToGlobalValidationOutputWithContext(ctx context.Context) GlobalValidationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalValidationOutput)
}

func (i GlobalValidationArgs) ToGlobalValidationPtrOutput() GlobalValidationPtrOutput {
	return i.ToGlobalValidationPtrOutputWithContext(context.Background())
}

func (i GlobalValidationArgs) ToGlobalValidationPtrOutputWithContext(ctx context.Context) GlobalValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalValidationOutput).ToGlobalValidationPtrOutputWithContext(ctx)
}









type GlobalValidationPtrInput interface {
	pulumi.Input

	ToGlobalValidationPtrOutput() GlobalValidationPtrOutput
	ToGlobalValidationPtrOutputWithContext(context.Context) GlobalValidationPtrOutput
}

type globalValidationPtrType GlobalValidationArgs

func GlobalValidationPtr(v *GlobalValidationArgs) GlobalValidationPtrInput {
	return (*globalValidationPtrType)(v)
}

func (*globalValidationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalValidation)(nil)).Elem()
}

func (i *globalValidationPtrType) ToGlobalValidationPtrOutput() GlobalValidationPtrOutput {
	return i.ToGlobalValidationPtrOutputWithContext(context.Background())
}

func (i *globalValidationPtrType) ToGlobalValidationPtrOutputWithContext(ctx context.Context) GlobalValidationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalValidationPtrOutput)
}

type GlobalValidationOutput struct{ *pulumi.OutputState }

func (GlobalValidationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalValidation)(nil)).Elem()
}

func (o GlobalValidationOutput) ToGlobalValidationOutput() GlobalValidationOutput {
	return o
}

func (o GlobalValidationOutput) ToGlobalValidationOutputWithContext(ctx context.Context) GlobalValidationOutput {
	return o
}

func (o GlobalValidationOutput) ToGlobalValidationPtrOutput() GlobalValidationPtrOutput {
	return o.ToGlobalValidationPtrOutputWithContext(context.Background())
}

func (o GlobalValidationOutput) ToGlobalValidationPtrOutputWithContext(ctx context.Context) GlobalValidationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GlobalValidation) *GlobalValidation {
		return &v
	}).(GlobalValidationPtrOutput)
}

func (o GlobalValidationOutput) ExcludedPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GlobalValidation) []string { return v.ExcludedPaths }).(pulumi.StringArrayOutput)
}

func (o GlobalValidationOutput) RedirectToProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GlobalValidation) *string { return v.RedirectToProvider }).(pulumi.StringPtrOutput)
}

func (o GlobalValidationOutput) RequireAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GlobalValidation) *bool { return v.RequireAuthentication }).(pulumi.BoolPtrOutput)
}

func (o GlobalValidationOutput) UnauthenticatedClientAction() UnauthenticatedClientActionV2PtrOutput {
	return o.ApplyT(func(v GlobalValidation) *UnauthenticatedClientActionV2 { return v.UnauthenticatedClientAction }).(UnauthenticatedClientActionV2PtrOutput)
}

type GlobalValidationPtrOutput struct{ *pulumi.OutputState }

func (GlobalValidationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalValidation)(nil)).Elem()
}

func (o GlobalValidationPtrOutput) ToGlobalValidationPtrOutput() GlobalValidationPtrOutput {
	return o
}

func (o GlobalValidationPtrOutput) ToGlobalValidationPtrOutputWithContext(ctx context.Context) GlobalValidationPtrOutput {
	return o
}

func (o GlobalValidationPtrOutput) Elem() GlobalValidationOutput {
	return o.ApplyT(func(v *GlobalValidation) GlobalValidation {
		if v != nil {
			return *v
		}
		var ret GlobalValidation
		return ret
	}).(GlobalValidationOutput)
}

func (o GlobalValidationPtrOutput) ExcludedPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GlobalValidation) []string {
		if v == nil {
			return nil
		}
		return v.ExcludedPaths
	}).(pulumi.StringArrayOutput)
}

func (o GlobalValidationPtrOutput) RedirectToProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GlobalValidation) *string {
		if v == nil {
			return nil
		}
		return v.RedirectToProvider
	}).(pulumi.StringPtrOutput)
}

func (o GlobalValidationPtrOutput) RequireAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GlobalValidation) *bool {
		if v == nil {
			return nil
		}
		return v.RequireAuthentication
	}).(pulumi.BoolPtrOutput)
}

func (o GlobalValidationPtrOutput) UnauthenticatedClientAction() UnauthenticatedClientActionV2PtrOutput {
	return o.ApplyT(func(v *GlobalValidation) *UnauthenticatedClientActionV2 {
		if v == nil {
			return nil
		}
		return v.UnauthenticatedClientAction
	}).(UnauthenticatedClientActionV2PtrOutput)
}

type GlobalValidationResponse struct {
	ExcludedPaths               []string `pulumi:"excludedPaths"`
	RedirectToProvider          *string  `pulumi:"redirectToProvider"`
	RequireAuthentication       *bool    `pulumi:"requireAuthentication"`
	UnauthenticatedClientAction *string  `pulumi:"unauthenticatedClientAction"`
}

type GlobalValidationResponseOutput struct{ *pulumi.OutputState }

func (GlobalValidationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalValidationResponse)(nil)).Elem()
}

func (o GlobalValidationResponseOutput) ToGlobalValidationResponseOutput() GlobalValidationResponseOutput {
	return o
}

func (o GlobalValidationResponseOutput) ToGlobalValidationResponseOutputWithContext(ctx context.Context) GlobalValidationResponseOutput {
	return o
}

func (o GlobalValidationResponseOutput) ExcludedPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GlobalValidationResponse) []string { return v.ExcludedPaths }).(pulumi.StringArrayOutput)
}

func (o GlobalValidationResponseOutput) RedirectToProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GlobalValidationResponse) *string { return v.RedirectToProvider }).(pulumi.StringPtrOutput)
}

func (o GlobalValidationResponseOutput) RequireAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GlobalValidationResponse) *bool { return v.RequireAuthentication }).(pulumi.BoolPtrOutput)
}

func (o GlobalValidationResponseOutput) UnauthenticatedClientAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GlobalValidationResponse) *string { return v.UnauthenticatedClientAction }).(pulumi.StringPtrOutput)
}

type GlobalValidationResponsePtrOutput struct{ *pulumi.OutputState }

func (GlobalValidationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalValidationResponse)(nil)).Elem()
}

func (o GlobalValidationResponsePtrOutput) ToGlobalValidationResponsePtrOutput() GlobalValidationResponsePtrOutput {
	return o
}

func (o GlobalValidationResponsePtrOutput) ToGlobalValidationResponsePtrOutputWithContext(ctx context.Context) GlobalValidationResponsePtrOutput {
	return o
}

func (o GlobalValidationResponsePtrOutput) Elem() GlobalValidationResponseOutput {
	return o.ApplyT(func(v *GlobalValidationResponse) GlobalValidationResponse {
		if v != nil {
			return *v
		}
		var ret GlobalValidationResponse
		return ret
	}).(GlobalValidationResponseOutput)
}

func (o GlobalValidationResponsePtrOutput) ExcludedPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GlobalValidationResponse) []string {
		if v == nil {
			return nil
		}
		return v.ExcludedPaths
	}).(pulumi.StringArrayOutput)
}

func (o GlobalValidationResponsePtrOutput) RedirectToProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GlobalValidationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RedirectToProvider
	}).(pulumi.StringPtrOutput)
}

func (o GlobalValidationResponsePtrOutput) RequireAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GlobalValidationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequireAuthentication
	}).(pulumi.BoolPtrOutput)
}

func (o GlobalValidationResponsePtrOutput) UnauthenticatedClientAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GlobalValidationResponse) *string {
		if v == nil {
			return nil
		}
		return v.UnauthenticatedClientAction
	}).(pulumi.StringPtrOutput)
}

type Google struct {
	Enabled      *bool                       `pulumi:"enabled"`
	Login        *LoginScopes                `pulumi:"login"`
	Registration *ClientRegistration         `pulumi:"registration"`
	Validation   *AllowedAudiencesValidation `pulumi:"validation"`
}





type GoogleInput interface {
	pulumi.Input

	ToGoogleOutput() GoogleOutput
	ToGoogleOutputWithContext(context.Context) GoogleOutput
}

type GoogleArgs struct {
	Enabled      pulumi.BoolPtrInput                `pulumi:"enabled"`
	Login        LoginScopesPtrInput                `pulumi:"login"`
	Registration ClientRegistrationPtrInput         `pulumi:"registration"`
	Validation   AllowedAudiencesValidationPtrInput `pulumi:"validation"`
}

func (GoogleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Google)(nil)).Elem()
}

func (i GoogleArgs) ToGoogleOutput() GoogleOutput {
	return i.ToGoogleOutputWithContext(context.Background())
}

func (i GoogleArgs) ToGoogleOutputWithContext(ctx context.Context) GoogleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleOutput)
}

func (i GoogleArgs) ToGooglePtrOutput() GooglePtrOutput {
	return i.ToGooglePtrOutputWithContext(context.Background())
}

func (i GoogleArgs) ToGooglePtrOutputWithContext(ctx context.Context) GooglePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleOutput).ToGooglePtrOutputWithContext(ctx)
}









type GooglePtrInput interface {
	pulumi.Input

	ToGooglePtrOutput() GooglePtrOutput
	ToGooglePtrOutputWithContext(context.Context) GooglePtrOutput
}

type googlePtrType GoogleArgs

func GooglePtr(v *GoogleArgs) GooglePtrInput {
	return (*googlePtrType)(v)
}

func (*googlePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Google)(nil)).Elem()
}

func (i *googlePtrType) ToGooglePtrOutput() GooglePtrOutput {
	return i.ToGooglePtrOutputWithContext(context.Background())
}

func (i *googlePtrType) ToGooglePtrOutputWithContext(ctx context.Context) GooglePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GooglePtrOutput)
}

type GoogleOutput struct{ *pulumi.OutputState }

func (GoogleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Google)(nil)).Elem()
}

func (o GoogleOutput) ToGoogleOutput() GoogleOutput {
	return o
}

func (o GoogleOutput) ToGoogleOutputWithContext(ctx context.Context) GoogleOutput {
	return o
}

func (o GoogleOutput) ToGooglePtrOutput() GooglePtrOutput {
	return o.ToGooglePtrOutputWithContext(context.Background())
}

func (o GoogleOutput) ToGooglePtrOutputWithContext(ctx context.Context) GooglePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Google) *Google {
		return &v
	}).(GooglePtrOutput)
}

func (o GoogleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Google) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o GoogleOutput) Login() LoginScopesPtrOutput {
	return o.ApplyT(func(v Google) *LoginScopes { return v.Login }).(LoginScopesPtrOutput)
}

func (o GoogleOutput) Registration() ClientRegistrationPtrOutput {
	return o.ApplyT(func(v Google) *ClientRegistration { return v.Registration }).(ClientRegistrationPtrOutput)
}

func (o GoogleOutput) Validation() AllowedAudiencesValidationPtrOutput {
	return o.ApplyT(func(v Google) *AllowedAudiencesValidation { return v.Validation }).(AllowedAudiencesValidationPtrOutput)
}

type GooglePtrOutput struct{ *pulumi.OutputState }

func (GooglePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Google)(nil)).Elem()
}

func (o GooglePtrOutput) ToGooglePtrOutput() GooglePtrOutput {
	return o
}

func (o GooglePtrOutput) ToGooglePtrOutputWithContext(ctx context.Context) GooglePtrOutput {
	return o
}

func (o GooglePtrOutput) Elem() GoogleOutput {
	return o.ApplyT(func(v *Google) Google {
		if v != nil {
			return *v
		}
		var ret Google
		return ret
	}).(GoogleOutput)
}

func (o GooglePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Google) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o GooglePtrOutput) Login() LoginScopesPtrOutput {
	return o.ApplyT(func(v *Google) *LoginScopes {
		if v == nil {
			return nil
		}
		return v.Login
	}).(LoginScopesPtrOutput)
}

func (o GooglePtrOutput) Registration() ClientRegistrationPtrOutput {
	return o.ApplyT(func(v *Google) *ClientRegistration {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(ClientRegistrationPtrOutput)
}

func (o GooglePtrOutput) Validation() AllowedAudiencesValidationPtrOutput {
	return o.ApplyT(func(v *Google) *AllowedAudiencesValidation {
		if v == nil {
			return nil
		}
		return v.Validation
	}).(AllowedAudiencesValidationPtrOutput)
}

type GoogleResponse struct {
	Enabled      *bool                               `pulumi:"enabled"`
	Login        *LoginScopesResponse                `pulumi:"login"`
	Registration *ClientRegistrationResponse         `pulumi:"registration"`
	Validation   *AllowedAudiencesValidationResponse `pulumi:"validation"`
}

type GoogleResponseOutput struct{ *pulumi.OutputState }

func (GoogleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleResponse)(nil)).Elem()
}

func (o GoogleResponseOutput) ToGoogleResponseOutput() GoogleResponseOutput {
	return o
}

func (o GoogleResponseOutput) ToGoogleResponseOutputWithContext(ctx context.Context) GoogleResponseOutput {
	return o
}

func (o GoogleResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GoogleResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o GoogleResponseOutput) Login() LoginScopesResponsePtrOutput {
	return o.ApplyT(func(v GoogleResponse) *LoginScopesResponse { return v.Login }).(LoginScopesResponsePtrOutput)
}

func (o GoogleResponseOutput) Registration() ClientRegistrationResponsePtrOutput {
	return o.ApplyT(func(v GoogleResponse) *ClientRegistrationResponse { return v.Registration }).(ClientRegistrationResponsePtrOutput)
}

func (o GoogleResponseOutput) Validation() AllowedAudiencesValidationResponsePtrOutput {
	return o.ApplyT(func(v GoogleResponse) *AllowedAudiencesValidationResponse { return v.Validation }).(AllowedAudiencesValidationResponsePtrOutput)
}

type GoogleResponsePtrOutput struct{ *pulumi.OutputState }

func (GoogleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GoogleResponse)(nil)).Elem()
}

func (o GoogleResponsePtrOutput) ToGoogleResponsePtrOutput() GoogleResponsePtrOutput {
	return o
}

func (o GoogleResponsePtrOutput) ToGoogleResponsePtrOutputWithContext(ctx context.Context) GoogleResponsePtrOutput {
	return o
}

func (o GoogleResponsePtrOutput) Elem() GoogleResponseOutput {
	return o.ApplyT(func(v *GoogleResponse) GoogleResponse {
		if v != nil {
			return *v
		}
		var ret GoogleResponse
		return ret
	}).(GoogleResponseOutput)
}

func (o GoogleResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GoogleResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o GoogleResponsePtrOutput) Login() LoginScopesResponsePtrOutput {
	return o.ApplyT(func(v *GoogleResponse) *LoginScopesResponse {
		if v == nil {
			return nil
		}
		return v.Login
	}).(LoginScopesResponsePtrOutput)
}

func (o GoogleResponsePtrOutput) Registration() ClientRegistrationResponsePtrOutput {
	return o.ApplyT(func(v *GoogleResponse) *ClientRegistrationResponse {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(ClientRegistrationResponsePtrOutput)
}

func (o GoogleResponsePtrOutput) Validation() AllowedAudiencesValidationResponsePtrOutput {
	return o.ApplyT(func(v *GoogleResponse) *AllowedAudiencesValidationResponse {
		if v == nil {
			return nil
		}
		return v.Validation
	}).(AllowedAudiencesValidationResponsePtrOutput)
}

type HandlerMapping struct {
	Arguments       *string `pulumi:"arguments"`
	Extension       *string `pulumi:"extension"`
	ScriptProcessor *string `pulumi:"scriptProcessor"`
}





type HandlerMappingInput interface {
	pulumi.Input

	ToHandlerMappingOutput() HandlerMappingOutput
	ToHandlerMappingOutputWithContext(context.Context) HandlerMappingOutput
}

type HandlerMappingArgs struct {
	Arguments       pulumi.StringPtrInput `pulumi:"arguments"`
	Extension       pulumi.StringPtrInput `pulumi:"extension"`
	ScriptProcessor pulumi.StringPtrInput `pulumi:"scriptProcessor"`
}

func (HandlerMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HandlerMapping)(nil)).Elem()
}

func (i HandlerMappingArgs) ToHandlerMappingOutput() HandlerMappingOutput {
	return i.ToHandlerMappingOutputWithContext(context.Background())
}

func (i HandlerMappingArgs) ToHandlerMappingOutputWithContext(ctx context.Context) HandlerMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HandlerMappingOutput)
}





type HandlerMappingArrayInput interface {
	pulumi.Input

	ToHandlerMappingArrayOutput() HandlerMappingArrayOutput
	ToHandlerMappingArrayOutputWithContext(context.Context) HandlerMappingArrayOutput
}

type HandlerMappingArray []HandlerMappingInput

func (HandlerMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HandlerMapping)(nil)).Elem()
}

func (i HandlerMappingArray) ToHandlerMappingArrayOutput() HandlerMappingArrayOutput {
	return i.ToHandlerMappingArrayOutputWithContext(context.Background())
}

func (i HandlerMappingArray) ToHandlerMappingArrayOutputWithContext(ctx context.Context) HandlerMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HandlerMappingArrayOutput)
}

type HandlerMappingOutput struct{ *pulumi.OutputState }

func (HandlerMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HandlerMapping)(nil)).Elem()
}

func (o HandlerMappingOutput) ToHandlerMappingOutput() HandlerMappingOutput {
	return o
}

func (o HandlerMappingOutput) ToHandlerMappingOutputWithContext(ctx context.Context) HandlerMappingOutput {
	return o
}

func (o HandlerMappingOutput) Arguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HandlerMapping) *string { return v.Arguments }).(pulumi.StringPtrOutput)
}

func (o HandlerMappingOutput) Extension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HandlerMapping) *string { return v.Extension }).(pulumi.StringPtrOutput)
}

func (o HandlerMappingOutput) ScriptProcessor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HandlerMapping) *string { return v.ScriptProcessor }).(pulumi.StringPtrOutput)
}

type HandlerMappingArrayOutput struct{ *pulumi.OutputState }

func (HandlerMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HandlerMapping)(nil)).Elem()
}

func (o HandlerMappingArrayOutput) ToHandlerMappingArrayOutput() HandlerMappingArrayOutput {
	return o
}

func (o HandlerMappingArrayOutput) ToHandlerMappingArrayOutputWithContext(ctx context.Context) HandlerMappingArrayOutput {
	return o
}

func (o HandlerMappingArrayOutput) Index(i pulumi.IntInput) HandlerMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HandlerMapping {
		return vs[0].([]HandlerMapping)[vs[1].(int)]
	}).(HandlerMappingOutput)
}

type HandlerMappingResponse struct {
	Arguments       *string `pulumi:"arguments"`
	Extension       *string `pulumi:"extension"`
	ScriptProcessor *string `pulumi:"scriptProcessor"`
}

type HandlerMappingResponseOutput struct{ *pulumi.OutputState }

func (HandlerMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HandlerMappingResponse)(nil)).Elem()
}

func (o HandlerMappingResponseOutput) ToHandlerMappingResponseOutput() HandlerMappingResponseOutput {
	return o
}

func (o HandlerMappingResponseOutput) ToHandlerMappingResponseOutputWithContext(ctx context.Context) HandlerMappingResponseOutput {
	return o
}

func (o HandlerMappingResponseOutput) Arguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HandlerMappingResponse) *string { return v.Arguments }).(pulumi.StringPtrOutput)
}

func (o HandlerMappingResponseOutput) Extension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HandlerMappingResponse) *string { return v.Extension }).(pulumi.StringPtrOutput)
}

func (o HandlerMappingResponseOutput) ScriptProcessor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HandlerMappingResponse) *string { return v.ScriptProcessor }).(pulumi.StringPtrOutput)
}

type HandlerMappingResponseArrayOutput struct{ *pulumi.OutputState }

func (HandlerMappingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HandlerMappingResponse)(nil)).Elem()
}

func (o HandlerMappingResponseArrayOutput) ToHandlerMappingResponseArrayOutput() HandlerMappingResponseArrayOutput {
	return o
}

func (o HandlerMappingResponseArrayOutput) ToHandlerMappingResponseArrayOutputWithContext(ctx context.Context) HandlerMappingResponseArrayOutput {
	return o
}

func (o HandlerMappingResponseArrayOutput) Index(i pulumi.IntInput) HandlerMappingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HandlerMappingResponse {
		return vs[0].([]HandlerMappingResponse)[vs[1].(int)]
	}).(HandlerMappingResponseOutput)
}

type HostNameSslState struct {
	HostType   *HostType `pulumi:"hostType"`
	Name       *string   `pulumi:"name"`
	SslState   *SslState `pulumi:"sslState"`
	Thumbprint *string   `pulumi:"thumbprint"`
	ToUpdate   *bool     `pulumi:"toUpdate"`
	VirtualIP  *string   `pulumi:"virtualIP"`
}





type HostNameSslStateInput interface {
	pulumi.Input

	ToHostNameSslStateOutput() HostNameSslStateOutput
	ToHostNameSslStateOutputWithContext(context.Context) HostNameSslStateOutput
}

type HostNameSslStateArgs struct {
	HostType   HostTypePtrInput      `pulumi:"hostType"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	SslState   SslStatePtrInput      `pulumi:"sslState"`
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
	ToUpdate   pulumi.BoolPtrInput   `pulumi:"toUpdate"`
	VirtualIP  pulumi.StringPtrInput `pulumi:"virtualIP"`
}

func (HostNameSslStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameSslState)(nil)).Elem()
}

func (i HostNameSslStateArgs) ToHostNameSslStateOutput() HostNameSslStateOutput {
	return i.ToHostNameSslStateOutputWithContext(context.Background())
}

func (i HostNameSslStateArgs) ToHostNameSslStateOutputWithContext(ctx context.Context) HostNameSslStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostNameSslStateOutput)
}





type HostNameSslStateArrayInput interface {
	pulumi.Input

	ToHostNameSslStateArrayOutput() HostNameSslStateArrayOutput
	ToHostNameSslStateArrayOutputWithContext(context.Context) HostNameSslStateArrayOutput
}

type HostNameSslStateArray []HostNameSslStateInput

func (HostNameSslStateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostNameSslState)(nil)).Elem()
}

func (i HostNameSslStateArray) ToHostNameSslStateArrayOutput() HostNameSslStateArrayOutput {
	return i.ToHostNameSslStateArrayOutputWithContext(context.Background())
}

func (i HostNameSslStateArray) ToHostNameSslStateArrayOutputWithContext(ctx context.Context) HostNameSslStateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostNameSslStateArrayOutput)
}

type HostNameSslStateOutput struct{ *pulumi.OutputState }

func (HostNameSslStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameSslState)(nil)).Elem()
}

func (o HostNameSslStateOutput) ToHostNameSslStateOutput() HostNameSslStateOutput {
	return o
}

func (o HostNameSslStateOutput) ToHostNameSslStateOutputWithContext(ctx context.Context) HostNameSslStateOutput {
	return o
}

func (o HostNameSslStateOutput) HostType() HostTypePtrOutput {
	return o.ApplyT(func(v HostNameSslState) *HostType { return v.HostType }).(HostTypePtrOutput)
}

func (o HostNameSslStateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslState) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HostNameSslStateOutput) SslState() SslStatePtrOutput {
	return o.ApplyT(func(v HostNameSslState) *SslState { return v.SslState }).(SslStatePtrOutput)
}

func (o HostNameSslStateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslState) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func (o HostNameSslStateOutput) ToUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HostNameSslState) *bool { return v.ToUpdate }).(pulumi.BoolPtrOutput)
}

func (o HostNameSslStateOutput) VirtualIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslState) *string { return v.VirtualIP }).(pulumi.StringPtrOutput)
}

type HostNameSslStateArrayOutput struct{ *pulumi.OutputState }

func (HostNameSslStateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostNameSslState)(nil)).Elem()
}

func (o HostNameSslStateArrayOutput) ToHostNameSslStateArrayOutput() HostNameSslStateArrayOutput {
	return o
}

func (o HostNameSslStateArrayOutput) ToHostNameSslStateArrayOutputWithContext(ctx context.Context) HostNameSslStateArrayOutput {
	return o
}

func (o HostNameSslStateArrayOutput) Index(i pulumi.IntInput) HostNameSslStateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostNameSslState {
		return vs[0].([]HostNameSslState)[vs[1].(int)]
	}).(HostNameSslStateOutput)
}

type HostNameSslStateResponse struct {
	HostType   *string `pulumi:"hostType"`
	Name       *string `pulumi:"name"`
	SslState   *string `pulumi:"sslState"`
	Thumbprint *string `pulumi:"thumbprint"`
	ToUpdate   *bool   `pulumi:"toUpdate"`
	VirtualIP  *string `pulumi:"virtualIP"`
}

type HostNameSslStateResponseOutput struct{ *pulumi.OutputState }

func (HostNameSslStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameSslStateResponse)(nil)).Elem()
}

func (o HostNameSslStateResponseOutput) ToHostNameSslStateResponseOutput() HostNameSslStateResponseOutput {
	return o
}

func (o HostNameSslStateResponseOutput) ToHostNameSslStateResponseOutputWithContext(ctx context.Context) HostNameSslStateResponseOutput {
	return o
}

func (o HostNameSslStateResponseOutput) HostType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslStateResponse) *string { return v.HostType }).(pulumi.StringPtrOutput)
}

func (o HostNameSslStateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslStateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HostNameSslStateResponseOutput) SslState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslStateResponse) *string { return v.SslState }).(pulumi.StringPtrOutput)
}

func (o HostNameSslStateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslStateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func (o HostNameSslStateResponseOutput) ToUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HostNameSslStateResponse) *bool { return v.ToUpdate }).(pulumi.BoolPtrOutput)
}

func (o HostNameSslStateResponseOutput) VirtualIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslStateResponse) *string { return v.VirtualIP }).(pulumi.StringPtrOutput)
}

type HostNameSslStateResponseArrayOutput struct{ *pulumi.OutputState }

func (HostNameSslStateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostNameSslStateResponse)(nil)).Elem()
}

func (o HostNameSslStateResponseArrayOutput) ToHostNameSslStateResponseArrayOutput() HostNameSslStateResponseArrayOutput {
	return o
}

func (o HostNameSslStateResponseArrayOutput) ToHostNameSslStateResponseArrayOutputWithContext(ctx context.Context) HostNameSslStateResponseArrayOutput {
	return o
}

func (o HostNameSslStateResponseArrayOutput) Index(i pulumi.IntInput) HostNameSslStateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostNameSslStateResponse {
		return vs[0].([]HostNameSslStateResponse)[vs[1].(int)]
	}).(HostNameSslStateResponseOutput)
}

type HostingEnvironmentProfile struct {
	Id *string `pulumi:"id"`
}





type HostingEnvironmentProfileInput interface {
	pulumi.Input

	ToHostingEnvironmentProfileOutput() HostingEnvironmentProfileOutput
	ToHostingEnvironmentProfileOutputWithContext(context.Context) HostingEnvironmentProfileOutput
}

type HostingEnvironmentProfileArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (HostingEnvironmentProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentProfile)(nil)).Elem()
}

func (i HostingEnvironmentProfileArgs) ToHostingEnvironmentProfileOutput() HostingEnvironmentProfileOutput {
	return i.ToHostingEnvironmentProfileOutputWithContext(context.Background())
}

func (i HostingEnvironmentProfileArgs) ToHostingEnvironmentProfileOutputWithContext(ctx context.Context) HostingEnvironmentProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentProfileOutput)
}

func (i HostingEnvironmentProfileArgs) ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput {
	return i.ToHostingEnvironmentProfilePtrOutputWithContext(context.Background())
}

func (i HostingEnvironmentProfileArgs) ToHostingEnvironmentProfilePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentProfileOutput).ToHostingEnvironmentProfilePtrOutputWithContext(ctx)
}









type HostingEnvironmentProfilePtrInput interface {
	pulumi.Input

	ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput
	ToHostingEnvironmentProfilePtrOutputWithContext(context.Context) HostingEnvironmentProfilePtrOutput
}

type hostingEnvironmentProfilePtrType HostingEnvironmentProfileArgs

func HostingEnvironmentProfilePtr(v *HostingEnvironmentProfileArgs) HostingEnvironmentProfilePtrInput {
	return (*hostingEnvironmentProfilePtrType)(v)
}

func (*hostingEnvironmentProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingEnvironmentProfile)(nil)).Elem()
}

func (i *hostingEnvironmentProfilePtrType) ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput {
	return i.ToHostingEnvironmentProfilePtrOutputWithContext(context.Background())
}

func (i *hostingEnvironmentProfilePtrType) ToHostingEnvironmentProfilePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentProfilePtrOutput)
}

type HostingEnvironmentProfileOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentProfile)(nil)).Elem()
}

func (o HostingEnvironmentProfileOutput) ToHostingEnvironmentProfileOutput() HostingEnvironmentProfileOutput {
	return o
}

func (o HostingEnvironmentProfileOutput) ToHostingEnvironmentProfileOutputWithContext(ctx context.Context) HostingEnvironmentProfileOutput {
	return o
}

func (o HostingEnvironmentProfileOutput) ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput {
	return o.ToHostingEnvironmentProfilePtrOutputWithContext(context.Background())
}

func (o HostingEnvironmentProfileOutput) ToHostingEnvironmentProfilePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostingEnvironmentProfile) *HostingEnvironmentProfile {
		return &v
	}).(HostingEnvironmentProfilePtrOutput)
}

func (o HostingEnvironmentProfileOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentProfile) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type HostingEnvironmentProfilePtrOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingEnvironmentProfile)(nil)).Elem()
}

func (o HostingEnvironmentProfilePtrOutput) ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput {
	return o
}

func (o HostingEnvironmentProfilePtrOutput) ToHostingEnvironmentProfilePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfilePtrOutput {
	return o
}

func (o HostingEnvironmentProfilePtrOutput) Elem() HostingEnvironmentProfileOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfile) HostingEnvironmentProfile {
		if v != nil {
			return *v
		}
		var ret HostingEnvironmentProfile
		return ret
	}).(HostingEnvironmentProfileOutput)
}

func (o HostingEnvironmentProfilePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfile) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type HostingEnvironmentProfileResponse struct {
	Id   *string `pulumi:"id"`
	Name string  `pulumi:"name"`
	Type string  `pulumi:"type"`
}

type HostingEnvironmentProfileResponseOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentProfileResponse)(nil)).Elem()
}

func (o HostingEnvironmentProfileResponseOutput) ToHostingEnvironmentProfileResponseOutput() HostingEnvironmentProfileResponseOutput {
	return o
}

func (o HostingEnvironmentProfileResponseOutput) ToHostingEnvironmentProfileResponseOutputWithContext(ctx context.Context) HostingEnvironmentProfileResponseOutput {
	return o
}

func (o HostingEnvironmentProfileResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentProfileResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfileResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v HostingEnvironmentProfileResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o HostingEnvironmentProfileResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v HostingEnvironmentProfileResponse) string { return v.Type }).(pulumi.StringOutput)
}

type HostingEnvironmentProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingEnvironmentProfileResponse)(nil)).Elem()
}

func (o HostingEnvironmentProfileResponsePtrOutput) ToHostingEnvironmentProfileResponsePtrOutput() HostingEnvironmentProfileResponsePtrOutput {
	return o
}

func (o HostingEnvironmentProfileResponsePtrOutput) ToHostingEnvironmentProfileResponsePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfileResponsePtrOutput {
	return o
}

func (o HostingEnvironmentProfileResponsePtrOutput) Elem() HostingEnvironmentProfileResponseOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) HostingEnvironmentProfileResponse {
		if v != nil {
			return *v
		}
		var ret HostingEnvironmentProfileResponse
		return ret
	}).(HostingEnvironmentProfileResponseOutput)
}

func (o HostingEnvironmentProfileResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfileResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfileResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type HttpLogsConfig struct {
	AzureBlobStorage *AzureBlobStorageHttpLogsConfig `pulumi:"azureBlobStorage"`
	FileSystem       *FileSystemHttpLogsConfig       `pulumi:"fileSystem"`
}





type HttpLogsConfigInput interface {
	pulumi.Input

	ToHttpLogsConfigOutput() HttpLogsConfigOutput
	ToHttpLogsConfigOutputWithContext(context.Context) HttpLogsConfigOutput
}

type HttpLogsConfigArgs struct {
	AzureBlobStorage AzureBlobStorageHttpLogsConfigPtrInput `pulumi:"azureBlobStorage"`
	FileSystem       FileSystemHttpLogsConfigPtrInput       `pulumi:"fileSystem"`
}

func (HttpLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpLogsConfig)(nil)).Elem()
}

func (i HttpLogsConfigArgs) ToHttpLogsConfigOutput() HttpLogsConfigOutput {
	return i.ToHttpLogsConfigOutputWithContext(context.Background())
}

func (i HttpLogsConfigArgs) ToHttpLogsConfigOutputWithContext(ctx context.Context) HttpLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpLogsConfigOutput)
}

func (i HttpLogsConfigArgs) ToHttpLogsConfigPtrOutput() HttpLogsConfigPtrOutput {
	return i.ToHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (i HttpLogsConfigArgs) ToHttpLogsConfigPtrOutputWithContext(ctx context.Context) HttpLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpLogsConfigOutput).ToHttpLogsConfigPtrOutputWithContext(ctx)
}









type HttpLogsConfigPtrInput interface {
	pulumi.Input

	ToHttpLogsConfigPtrOutput() HttpLogsConfigPtrOutput
	ToHttpLogsConfigPtrOutputWithContext(context.Context) HttpLogsConfigPtrOutput
}

type httpLogsConfigPtrType HttpLogsConfigArgs

func HttpLogsConfigPtr(v *HttpLogsConfigArgs) HttpLogsConfigPtrInput {
	return (*httpLogsConfigPtrType)(v)
}

func (*httpLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpLogsConfig)(nil)).Elem()
}

func (i *httpLogsConfigPtrType) ToHttpLogsConfigPtrOutput() HttpLogsConfigPtrOutput {
	return i.ToHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (i *httpLogsConfigPtrType) ToHttpLogsConfigPtrOutputWithContext(ctx context.Context) HttpLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpLogsConfigPtrOutput)
}

type HttpLogsConfigOutput struct{ *pulumi.OutputState }

func (HttpLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpLogsConfig)(nil)).Elem()
}

func (o HttpLogsConfigOutput) ToHttpLogsConfigOutput() HttpLogsConfigOutput {
	return o
}

func (o HttpLogsConfigOutput) ToHttpLogsConfigOutputWithContext(ctx context.Context) HttpLogsConfigOutput {
	return o
}

func (o HttpLogsConfigOutput) ToHttpLogsConfigPtrOutput() HttpLogsConfigPtrOutput {
	return o.ToHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (o HttpLogsConfigOutput) ToHttpLogsConfigPtrOutputWithContext(ctx context.Context) HttpLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpLogsConfig) *HttpLogsConfig {
		return &v
	}).(HttpLogsConfigPtrOutput)
}

func (o HttpLogsConfigOutput) AzureBlobStorage() AzureBlobStorageHttpLogsConfigPtrOutput {
	return o.ApplyT(func(v HttpLogsConfig) *AzureBlobStorageHttpLogsConfig { return v.AzureBlobStorage }).(AzureBlobStorageHttpLogsConfigPtrOutput)
}

func (o HttpLogsConfigOutput) FileSystem() FileSystemHttpLogsConfigPtrOutput {
	return o.ApplyT(func(v HttpLogsConfig) *FileSystemHttpLogsConfig { return v.FileSystem }).(FileSystemHttpLogsConfigPtrOutput)
}

type HttpLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (HttpLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpLogsConfig)(nil)).Elem()
}

func (o HttpLogsConfigPtrOutput) ToHttpLogsConfigPtrOutput() HttpLogsConfigPtrOutput {
	return o
}

func (o HttpLogsConfigPtrOutput) ToHttpLogsConfigPtrOutputWithContext(ctx context.Context) HttpLogsConfigPtrOutput {
	return o
}

func (o HttpLogsConfigPtrOutput) Elem() HttpLogsConfigOutput {
	return o.ApplyT(func(v *HttpLogsConfig) HttpLogsConfig {
		if v != nil {
			return *v
		}
		var ret HttpLogsConfig
		return ret
	}).(HttpLogsConfigOutput)
}

func (o HttpLogsConfigPtrOutput) AzureBlobStorage() AzureBlobStorageHttpLogsConfigPtrOutput {
	return o.ApplyT(func(v *HttpLogsConfig) *AzureBlobStorageHttpLogsConfig {
		if v == nil {
			return nil
		}
		return v.AzureBlobStorage
	}).(AzureBlobStorageHttpLogsConfigPtrOutput)
}

func (o HttpLogsConfigPtrOutput) FileSystem() FileSystemHttpLogsConfigPtrOutput {
	return o.ApplyT(func(v *HttpLogsConfig) *FileSystemHttpLogsConfig {
		if v == nil {
			return nil
		}
		return v.FileSystem
	}).(FileSystemHttpLogsConfigPtrOutput)
}

type HttpLogsConfigResponse struct {
	AzureBlobStorage *AzureBlobStorageHttpLogsConfigResponse `pulumi:"azureBlobStorage"`
	FileSystem       *FileSystemHttpLogsConfigResponse       `pulumi:"fileSystem"`
}

type HttpLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (HttpLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpLogsConfigResponse)(nil)).Elem()
}

func (o HttpLogsConfigResponseOutput) ToHttpLogsConfigResponseOutput() HttpLogsConfigResponseOutput {
	return o
}

func (o HttpLogsConfigResponseOutput) ToHttpLogsConfigResponseOutputWithContext(ctx context.Context) HttpLogsConfigResponseOutput {
	return o
}

func (o HttpLogsConfigResponseOutput) AzureBlobStorage() AzureBlobStorageHttpLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v HttpLogsConfigResponse) *AzureBlobStorageHttpLogsConfigResponse { return v.AzureBlobStorage }).(AzureBlobStorageHttpLogsConfigResponsePtrOutput)
}

func (o HttpLogsConfigResponseOutput) FileSystem() FileSystemHttpLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v HttpLogsConfigResponse) *FileSystemHttpLogsConfigResponse { return v.FileSystem }).(FileSystemHttpLogsConfigResponsePtrOutput)
}

type HttpLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (HttpLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpLogsConfigResponse)(nil)).Elem()
}

func (o HttpLogsConfigResponsePtrOutput) ToHttpLogsConfigResponsePtrOutput() HttpLogsConfigResponsePtrOutput {
	return o
}

func (o HttpLogsConfigResponsePtrOutput) ToHttpLogsConfigResponsePtrOutputWithContext(ctx context.Context) HttpLogsConfigResponsePtrOutput {
	return o
}

func (o HttpLogsConfigResponsePtrOutput) Elem() HttpLogsConfigResponseOutput {
	return o.ApplyT(func(v *HttpLogsConfigResponse) HttpLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret HttpLogsConfigResponse
		return ret
	}).(HttpLogsConfigResponseOutput)
}

func (o HttpLogsConfigResponsePtrOutput) AzureBlobStorage() AzureBlobStorageHttpLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v *HttpLogsConfigResponse) *AzureBlobStorageHttpLogsConfigResponse {
		if v == nil {
			return nil
		}
		return v.AzureBlobStorage
	}).(AzureBlobStorageHttpLogsConfigResponsePtrOutput)
}

func (o HttpLogsConfigResponsePtrOutput) FileSystem() FileSystemHttpLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v *HttpLogsConfigResponse) *FileSystemHttpLogsConfigResponse {
		if v == nil {
			return nil
		}
		return v.FileSystem
	}).(FileSystemHttpLogsConfigResponsePtrOutput)
}

type HttpSettings struct {
	ForwardProxy *ForwardProxy       `pulumi:"forwardProxy"`
	RequireHttps *bool               `pulumi:"requireHttps"`
	Routes       *HttpSettingsRoutes `pulumi:"routes"`
}





type HttpSettingsInput interface {
	pulumi.Input

	ToHttpSettingsOutput() HttpSettingsOutput
	ToHttpSettingsOutputWithContext(context.Context) HttpSettingsOutput
}

type HttpSettingsArgs struct {
	ForwardProxy ForwardProxyPtrInput       `pulumi:"forwardProxy"`
	RequireHttps pulumi.BoolPtrInput        `pulumi:"requireHttps"`
	Routes       HttpSettingsRoutesPtrInput `pulumi:"routes"`
}

func (HttpSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpSettings)(nil)).Elem()
}

func (i HttpSettingsArgs) ToHttpSettingsOutput() HttpSettingsOutput {
	return i.ToHttpSettingsOutputWithContext(context.Background())
}

func (i HttpSettingsArgs) ToHttpSettingsOutputWithContext(ctx context.Context) HttpSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpSettingsOutput)
}

func (i HttpSettingsArgs) ToHttpSettingsPtrOutput() HttpSettingsPtrOutput {
	return i.ToHttpSettingsPtrOutputWithContext(context.Background())
}

func (i HttpSettingsArgs) ToHttpSettingsPtrOutputWithContext(ctx context.Context) HttpSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpSettingsOutput).ToHttpSettingsPtrOutputWithContext(ctx)
}









type HttpSettingsPtrInput interface {
	pulumi.Input

	ToHttpSettingsPtrOutput() HttpSettingsPtrOutput
	ToHttpSettingsPtrOutputWithContext(context.Context) HttpSettingsPtrOutput
}

type httpSettingsPtrType HttpSettingsArgs

func HttpSettingsPtr(v *HttpSettingsArgs) HttpSettingsPtrInput {
	return (*httpSettingsPtrType)(v)
}

func (*httpSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpSettings)(nil)).Elem()
}

func (i *httpSettingsPtrType) ToHttpSettingsPtrOutput() HttpSettingsPtrOutput {
	return i.ToHttpSettingsPtrOutputWithContext(context.Background())
}

func (i *httpSettingsPtrType) ToHttpSettingsPtrOutputWithContext(ctx context.Context) HttpSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpSettingsPtrOutput)
}

type HttpSettingsOutput struct{ *pulumi.OutputState }

func (HttpSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpSettings)(nil)).Elem()
}

func (o HttpSettingsOutput) ToHttpSettingsOutput() HttpSettingsOutput {
	return o
}

func (o HttpSettingsOutput) ToHttpSettingsOutputWithContext(ctx context.Context) HttpSettingsOutput {
	return o
}

func (o HttpSettingsOutput) ToHttpSettingsPtrOutput() HttpSettingsPtrOutput {
	return o.ToHttpSettingsPtrOutputWithContext(context.Background())
}

func (o HttpSettingsOutput) ToHttpSettingsPtrOutputWithContext(ctx context.Context) HttpSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpSettings) *HttpSettings {
		return &v
	}).(HttpSettingsPtrOutput)
}

func (o HttpSettingsOutput) ForwardProxy() ForwardProxyPtrOutput {
	return o.ApplyT(func(v HttpSettings) *ForwardProxy { return v.ForwardProxy }).(ForwardProxyPtrOutput)
}

func (o HttpSettingsOutput) RequireHttps() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HttpSettings) *bool { return v.RequireHttps }).(pulumi.BoolPtrOutput)
}

func (o HttpSettingsOutput) Routes() HttpSettingsRoutesPtrOutput {
	return o.ApplyT(func(v HttpSettings) *HttpSettingsRoutes { return v.Routes }).(HttpSettingsRoutesPtrOutput)
}

type HttpSettingsPtrOutput struct{ *pulumi.OutputState }

func (HttpSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpSettings)(nil)).Elem()
}

func (o HttpSettingsPtrOutput) ToHttpSettingsPtrOutput() HttpSettingsPtrOutput {
	return o
}

func (o HttpSettingsPtrOutput) ToHttpSettingsPtrOutputWithContext(ctx context.Context) HttpSettingsPtrOutput {
	return o
}

func (o HttpSettingsPtrOutput) Elem() HttpSettingsOutput {
	return o.ApplyT(func(v *HttpSettings) HttpSettings {
		if v != nil {
			return *v
		}
		var ret HttpSettings
		return ret
	}).(HttpSettingsOutput)
}

func (o HttpSettingsPtrOutput) ForwardProxy() ForwardProxyPtrOutput {
	return o.ApplyT(func(v *HttpSettings) *ForwardProxy {
		if v == nil {
			return nil
		}
		return v.ForwardProxy
	}).(ForwardProxyPtrOutput)
}

func (o HttpSettingsPtrOutput) RequireHttps() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HttpSettings) *bool {
		if v == nil {
			return nil
		}
		return v.RequireHttps
	}).(pulumi.BoolPtrOutput)
}

func (o HttpSettingsPtrOutput) Routes() HttpSettingsRoutesPtrOutput {
	return o.ApplyT(func(v *HttpSettings) *HttpSettingsRoutes {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(HttpSettingsRoutesPtrOutput)
}

type HttpSettingsResponse struct {
	ForwardProxy *ForwardProxyResponse       `pulumi:"forwardProxy"`
	RequireHttps *bool                       `pulumi:"requireHttps"`
	Routes       *HttpSettingsRoutesResponse `pulumi:"routes"`
}

type HttpSettingsResponseOutput struct{ *pulumi.OutputState }

func (HttpSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpSettingsResponse)(nil)).Elem()
}

func (o HttpSettingsResponseOutput) ToHttpSettingsResponseOutput() HttpSettingsResponseOutput {
	return o
}

func (o HttpSettingsResponseOutput) ToHttpSettingsResponseOutputWithContext(ctx context.Context) HttpSettingsResponseOutput {
	return o
}

func (o HttpSettingsResponseOutput) ForwardProxy() ForwardProxyResponsePtrOutput {
	return o.ApplyT(func(v HttpSettingsResponse) *ForwardProxyResponse { return v.ForwardProxy }).(ForwardProxyResponsePtrOutput)
}

func (o HttpSettingsResponseOutput) RequireHttps() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HttpSettingsResponse) *bool { return v.RequireHttps }).(pulumi.BoolPtrOutput)
}

func (o HttpSettingsResponseOutput) Routes() HttpSettingsRoutesResponsePtrOutput {
	return o.ApplyT(func(v HttpSettingsResponse) *HttpSettingsRoutesResponse { return v.Routes }).(HttpSettingsRoutesResponsePtrOutput)
}

type HttpSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (HttpSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpSettingsResponse)(nil)).Elem()
}

func (o HttpSettingsResponsePtrOutput) ToHttpSettingsResponsePtrOutput() HttpSettingsResponsePtrOutput {
	return o
}

func (o HttpSettingsResponsePtrOutput) ToHttpSettingsResponsePtrOutputWithContext(ctx context.Context) HttpSettingsResponsePtrOutput {
	return o
}

func (o HttpSettingsResponsePtrOutput) Elem() HttpSettingsResponseOutput {
	return o.ApplyT(func(v *HttpSettingsResponse) HttpSettingsResponse {
		if v != nil {
			return *v
		}
		var ret HttpSettingsResponse
		return ret
	}).(HttpSettingsResponseOutput)
}

func (o HttpSettingsResponsePtrOutput) ForwardProxy() ForwardProxyResponsePtrOutput {
	return o.ApplyT(func(v *HttpSettingsResponse) *ForwardProxyResponse {
		if v == nil {
			return nil
		}
		return v.ForwardProxy
	}).(ForwardProxyResponsePtrOutput)
}

func (o HttpSettingsResponsePtrOutput) RequireHttps() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HttpSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequireHttps
	}).(pulumi.BoolPtrOutput)
}

func (o HttpSettingsResponsePtrOutput) Routes() HttpSettingsRoutesResponsePtrOutput {
	return o.ApplyT(func(v *HttpSettingsResponse) *HttpSettingsRoutesResponse {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(HttpSettingsRoutesResponsePtrOutput)
}

type HttpSettingsRoutes struct {
	ApiPrefix *string `pulumi:"apiPrefix"`
}





type HttpSettingsRoutesInput interface {
	pulumi.Input

	ToHttpSettingsRoutesOutput() HttpSettingsRoutesOutput
	ToHttpSettingsRoutesOutputWithContext(context.Context) HttpSettingsRoutesOutput
}

type HttpSettingsRoutesArgs struct {
	ApiPrefix pulumi.StringPtrInput `pulumi:"apiPrefix"`
}

func (HttpSettingsRoutesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpSettingsRoutes)(nil)).Elem()
}

func (i HttpSettingsRoutesArgs) ToHttpSettingsRoutesOutput() HttpSettingsRoutesOutput {
	return i.ToHttpSettingsRoutesOutputWithContext(context.Background())
}

func (i HttpSettingsRoutesArgs) ToHttpSettingsRoutesOutputWithContext(ctx context.Context) HttpSettingsRoutesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpSettingsRoutesOutput)
}

func (i HttpSettingsRoutesArgs) ToHttpSettingsRoutesPtrOutput() HttpSettingsRoutesPtrOutput {
	return i.ToHttpSettingsRoutesPtrOutputWithContext(context.Background())
}

func (i HttpSettingsRoutesArgs) ToHttpSettingsRoutesPtrOutputWithContext(ctx context.Context) HttpSettingsRoutesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpSettingsRoutesOutput).ToHttpSettingsRoutesPtrOutputWithContext(ctx)
}









type HttpSettingsRoutesPtrInput interface {
	pulumi.Input

	ToHttpSettingsRoutesPtrOutput() HttpSettingsRoutesPtrOutput
	ToHttpSettingsRoutesPtrOutputWithContext(context.Context) HttpSettingsRoutesPtrOutput
}

type httpSettingsRoutesPtrType HttpSettingsRoutesArgs

func HttpSettingsRoutesPtr(v *HttpSettingsRoutesArgs) HttpSettingsRoutesPtrInput {
	return (*httpSettingsRoutesPtrType)(v)
}

func (*httpSettingsRoutesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpSettingsRoutes)(nil)).Elem()
}

func (i *httpSettingsRoutesPtrType) ToHttpSettingsRoutesPtrOutput() HttpSettingsRoutesPtrOutput {
	return i.ToHttpSettingsRoutesPtrOutputWithContext(context.Background())
}

func (i *httpSettingsRoutesPtrType) ToHttpSettingsRoutesPtrOutputWithContext(ctx context.Context) HttpSettingsRoutesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpSettingsRoutesPtrOutput)
}

type HttpSettingsRoutesOutput struct{ *pulumi.OutputState }

func (HttpSettingsRoutesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpSettingsRoutes)(nil)).Elem()
}

func (o HttpSettingsRoutesOutput) ToHttpSettingsRoutesOutput() HttpSettingsRoutesOutput {
	return o
}

func (o HttpSettingsRoutesOutput) ToHttpSettingsRoutesOutputWithContext(ctx context.Context) HttpSettingsRoutesOutput {
	return o
}

func (o HttpSettingsRoutesOutput) ToHttpSettingsRoutesPtrOutput() HttpSettingsRoutesPtrOutput {
	return o.ToHttpSettingsRoutesPtrOutputWithContext(context.Background())
}

func (o HttpSettingsRoutesOutput) ToHttpSettingsRoutesPtrOutputWithContext(ctx context.Context) HttpSettingsRoutesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpSettingsRoutes) *HttpSettingsRoutes {
		return &v
	}).(HttpSettingsRoutesPtrOutput)
}

func (o HttpSettingsRoutesOutput) ApiPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpSettingsRoutes) *string { return v.ApiPrefix }).(pulumi.StringPtrOutput)
}

type HttpSettingsRoutesPtrOutput struct{ *pulumi.OutputState }

func (HttpSettingsRoutesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpSettingsRoutes)(nil)).Elem()
}

func (o HttpSettingsRoutesPtrOutput) ToHttpSettingsRoutesPtrOutput() HttpSettingsRoutesPtrOutput {
	return o
}

func (o HttpSettingsRoutesPtrOutput) ToHttpSettingsRoutesPtrOutputWithContext(ctx context.Context) HttpSettingsRoutesPtrOutput {
	return o
}

func (o HttpSettingsRoutesPtrOutput) Elem() HttpSettingsRoutesOutput {
	return o.ApplyT(func(v *HttpSettingsRoutes) HttpSettingsRoutes {
		if v != nil {
			return *v
		}
		var ret HttpSettingsRoutes
		return ret
	}).(HttpSettingsRoutesOutput)
}

func (o HttpSettingsRoutesPtrOutput) ApiPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpSettingsRoutes) *string {
		if v == nil {
			return nil
		}
		return v.ApiPrefix
	}).(pulumi.StringPtrOutput)
}

type HttpSettingsRoutesResponse struct {
	ApiPrefix *string `pulumi:"apiPrefix"`
}

type HttpSettingsRoutesResponseOutput struct{ *pulumi.OutputState }

func (HttpSettingsRoutesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpSettingsRoutesResponse)(nil)).Elem()
}

func (o HttpSettingsRoutesResponseOutput) ToHttpSettingsRoutesResponseOutput() HttpSettingsRoutesResponseOutput {
	return o
}

func (o HttpSettingsRoutesResponseOutput) ToHttpSettingsRoutesResponseOutputWithContext(ctx context.Context) HttpSettingsRoutesResponseOutput {
	return o
}

func (o HttpSettingsRoutesResponseOutput) ApiPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HttpSettingsRoutesResponse) *string { return v.ApiPrefix }).(pulumi.StringPtrOutput)
}

type HttpSettingsRoutesResponsePtrOutput struct{ *pulumi.OutputState }

func (HttpSettingsRoutesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpSettingsRoutesResponse)(nil)).Elem()
}

func (o HttpSettingsRoutesResponsePtrOutput) ToHttpSettingsRoutesResponsePtrOutput() HttpSettingsRoutesResponsePtrOutput {
	return o
}

func (o HttpSettingsRoutesResponsePtrOutput) ToHttpSettingsRoutesResponsePtrOutputWithContext(ctx context.Context) HttpSettingsRoutesResponsePtrOutput {
	return o
}

func (o HttpSettingsRoutesResponsePtrOutput) Elem() HttpSettingsRoutesResponseOutput {
	return o.ApplyT(func(v *HttpSettingsRoutesResponse) HttpSettingsRoutesResponse {
		if v != nil {
			return *v
		}
		var ret HttpSettingsRoutesResponse
		return ret
	}).(HttpSettingsRoutesResponseOutput)
}

func (o HttpSettingsRoutesResponsePtrOutput) ApiPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpSettingsRoutesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApiPrefix
	}).(pulumi.StringPtrOutput)
}

type IdentifierResponse struct {
	Id    string  `pulumi:"id"`
	Kind  *string `pulumi:"kind"`
	Name  string  `pulumi:"name"`
	Type  string  `pulumi:"type"`
	Value *string `pulumi:"value"`
}

type IdentifierResponseOutput struct{ *pulumi.OutputState }

func (IdentifierResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentifierResponse)(nil)).Elem()
}

func (o IdentifierResponseOutput) ToIdentifierResponseOutput() IdentifierResponseOutput {
	return o
}

func (o IdentifierResponseOutput) ToIdentifierResponseOutputWithContext(ctx context.Context) IdentifierResponseOutput {
	return o
}

func (o IdentifierResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IdentifierResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o IdentifierResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentifierResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o IdentifierResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v IdentifierResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o IdentifierResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentifierResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o IdentifierResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentifierResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type IdentifierResponseArrayOutput struct{ *pulumi.OutputState }

func (IdentifierResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentifierResponse)(nil)).Elem()
}

func (o IdentifierResponseArrayOutput) ToIdentifierResponseArrayOutput() IdentifierResponseArrayOutput {
	return o
}

func (o IdentifierResponseArrayOutput) ToIdentifierResponseArrayOutputWithContext(ctx context.Context) IdentifierResponseArrayOutput {
	return o
}

func (o IdentifierResponseArrayOutput) Index(i pulumi.IntInput) IdentifierResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IdentifierResponse {
		return vs[0].([]IdentifierResponse)[vs[1].(int)]
	}).(IdentifierResponseOutput)
}

type IdentityProviders struct {
	Apple                        *Apple                                 `pulumi:"apple"`
	AzureActiveDirectory         *AzureActiveDirectory                  `pulumi:"azureActiveDirectory"`
	AzureStaticWebApps           *AzureStaticWebApps                    `pulumi:"azureStaticWebApps"`
	CustomOpenIdConnectProviders map[string]CustomOpenIdConnectProvider `pulumi:"customOpenIdConnectProviders"`
	Facebook                     *Facebook                              `pulumi:"facebook"`
	GitHub                       *GitHub                                `pulumi:"gitHub"`
	Google                       *Google                                `pulumi:"google"`
	LegacyMicrosoftAccount       *LegacyMicrosoftAccount                `pulumi:"legacyMicrosoftAccount"`
	Twitter                      *Twitter                               `pulumi:"twitter"`
}





type IdentityProvidersInput interface {
	pulumi.Input

	ToIdentityProvidersOutput() IdentityProvidersOutput
	ToIdentityProvidersOutputWithContext(context.Context) IdentityProvidersOutput
}

type IdentityProvidersArgs struct {
	Apple                        ApplePtrInput                       `pulumi:"apple"`
	AzureActiveDirectory         AzureActiveDirectoryPtrInput        `pulumi:"azureActiveDirectory"`
	AzureStaticWebApps           AzureStaticWebAppsPtrInput          `pulumi:"azureStaticWebApps"`
	CustomOpenIdConnectProviders CustomOpenIdConnectProviderMapInput `pulumi:"customOpenIdConnectProviders"`
	Facebook                     FacebookPtrInput                    `pulumi:"facebook"`
	GitHub                       GitHubPtrInput                      `pulumi:"gitHub"`
	Google                       GooglePtrInput                      `pulumi:"google"`
	LegacyMicrosoftAccount       LegacyMicrosoftAccountPtrInput      `pulumi:"legacyMicrosoftAccount"`
	Twitter                      TwitterPtrInput                     `pulumi:"twitter"`
}

func (IdentityProvidersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProviders)(nil)).Elem()
}

func (i IdentityProvidersArgs) ToIdentityProvidersOutput() IdentityProvidersOutput {
	return i.ToIdentityProvidersOutputWithContext(context.Background())
}

func (i IdentityProvidersArgs) ToIdentityProvidersOutputWithContext(ctx context.Context) IdentityProvidersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProvidersOutput)
}

func (i IdentityProvidersArgs) ToIdentityProvidersPtrOutput() IdentityProvidersPtrOutput {
	return i.ToIdentityProvidersPtrOutputWithContext(context.Background())
}

func (i IdentityProvidersArgs) ToIdentityProvidersPtrOutputWithContext(ctx context.Context) IdentityProvidersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProvidersOutput).ToIdentityProvidersPtrOutputWithContext(ctx)
}









type IdentityProvidersPtrInput interface {
	pulumi.Input

	ToIdentityProvidersPtrOutput() IdentityProvidersPtrOutput
	ToIdentityProvidersPtrOutputWithContext(context.Context) IdentityProvidersPtrOutput
}

type identityProvidersPtrType IdentityProvidersArgs

func IdentityProvidersPtr(v *IdentityProvidersArgs) IdentityProvidersPtrInput {
	return (*identityProvidersPtrType)(v)
}

func (*identityProvidersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProviders)(nil)).Elem()
}

func (i *identityProvidersPtrType) ToIdentityProvidersPtrOutput() IdentityProvidersPtrOutput {
	return i.ToIdentityProvidersPtrOutputWithContext(context.Background())
}

func (i *identityProvidersPtrType) ToIdentityProvidersPtrOutputWithContext(ctx context.Context) IdentityProvidersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProvidersPtrOutput)
}

type IdentityProvidersOutput struct{ *pulumi.OutputState }

func (IdentityProvidersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProviders)(nil)).Elem()
}

func (o IdentityProvidersOutput) ToIdentityProvidersOutput() IdentityProvidersOutput {
	return o
}

func (o IdentityProvidersOutput) ToIdentityProvidersOutputWithContext(ctx context.Context) IdentityProvidersOutput {
	return o
}

func (o IdentityProvidersOutput) ToIdentityProvidersPtrOutput() IdentityProvidersPtrOutput {
	return o.ToIdentityProvidersPtrOutputWithContext(context.Background())
}

func (o IdentityProvidersOutput) ToIdentityProvidersPtrOutputWithContext(ctx context.Context) IdentityProvidersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityProviders) *IdentityProviders {
		return &v
	}).(IdentityProvidersPtrOutput)
}

func (o IdentityProvidersOutput) Apple() ApplePtrOutput {
	return o.ApplyT(func(v IdentityProviders) *Apple { return v.Apple }).(ApplePtrOutput)
}

func (o IdentityProvidersOutput) AzureActiveDirectory() AzureActiveDirectoryPtrOutput {
	return o.ApplyT(func(v IdentityProviders) *AzureActiveDirectory { return v.AzureActiveDirectory }).(AzureActiveDirectoryPtrOutput)
}

func (o IdentityProvidersOutput) AzureStaticWebApps() AzureStaticWebAppsPtrOutput {
	return o.ApplyT(func(v IdentityProviders) *AzureStaticWebApps { return v.AzureStaticWebApps }).(AzureStaticWebAppsPtrOutput)
}

func (o IdentityProvidersOutput) CustomOpenIdConnectProviders() CustomOpenIdConnectProviderMapOutput {
	return o.ApplyT(func(v IdentityProviders) map[string]CustomOpenIdConnectProvider {
		return v.CustomOpenIdConnectProviders
	}).(CustomOpenIdConnectProviderMapOutput)
}

func (o IdentityProvidersOutput) Facebook() FacebookPtrOutput {
	return o.ApplyT(func(v IdentityProviders) *Facebook { return v.Facebook }).(FacebookPtrOutput)
}

func (o IdentityProvidersOutput) GitHub() GitHubPtrOutput {
	return o.ApplyT(func(v IdentityProviders) *GitHub { return v.GitHub }).(GitHubPtrOutput)
}

func (o IdentityProvidersOutput) Google() GooglePtrOutput {
	return o.ApplyT(func(v IdentityProviders) *Google { return v.Google }).(GooglePtrOutput)
}

func (o IdentityProvidersOutput) LegacyMicrosoftAccount() LegacyMicrosoftAccountPtrOutput {
	return o.ApplyT(func(v IdentityProviders) *LegacyMicrosoftAccount { return v.LegacyMicrosoftAccount }).(LegacyMicrosoftAccountPtrOutput)
}

func (o IdentityProvidersOutput) Twitter() TwitterPtrOutput {
	return o.ApplyT(func(v IdentityProviders) *Twitter { return v.Twitter }).(TwitterPtrOutput)
}

type IdentityProvidersPtrOutput struct{ *pulumi.OutputState }

func (IdentityProvidersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProviders)(nil)).Elem()
}

func (o IdentityProvidersPtrOutput) ToIdentityProvidersPtrOutput() IdentityProvidersPtrOutput {
	return o
}

func (o IdentityProvidersPtrOutput) ToIdentityProvidersPtrOutputWithContext(ctx context.Context) IdentityProvidersPtrOutput {
	return o
}

func (o IdentityProvidersPtrOutput) Elem() IdentityProvidersOutput {
	return o.ApplyT(func(v *IdentityProviders) IdentityProviders {
		if v != nil {
			return *v
		}
		var ret IdentityProviders
		return ret
	}).(IdentityProvidersOutput)
}

func (o IdentityProvidersPtrOutput) Apple() ApplePtrOutput {
	return o.ApplyT(func(v *IdentityProviders) *Apple {
		if v == nil {
			return nil
		}
		return v.Apple
	}).(ApplePtrOutput)
}

func (o IdentityProvidersPtrOutput) AzureActiveDirectory() AzureActiveDirectoryPtrOutput {
	return o.ApplyT(func(v *IdentityProviders) *AzureActiveDirectory {
		if v == nil {
			return nil
		}
		return v.AzureActiveDirectory
	}).(AzureActiveDirectoryPtrOutput)
}

func (o IdentityProvidersPtrOutput) AzureStaticWebApps() AzureStaticWebAppsPtrOutput {
	return o.ApplyT(func(v *IdentityProviders) *AzureStaticWebApps {
		if v == nil {
			return nil
		}
		return v.AzureStaticWebApps
	}).(AzureStaticWebAppsPtrOutput)
}

func (o IdentityProvidersPtrOutput) CustomOpenIdConnectProviders() CustomOpenIdConnectProviderMapOutput {
	return o.ApplyT(func(v *IdentityProviders) map[string]CustomOpenIdConnectProvider {
		if v == nil {
			return nil
		}
		return v.CustomOpenIdConnectProviders
	}).(CustomOpenIdConnectProviderMapOutput)
}

func (o IdentityProvidersPtrOutput) Facebook() FacebookPtrOutput {
	return o.ApplyT(func(v *IdentityProviders) *Facebook {
		if v == nil {
			return nil
		}
		return v.Facebook
	}).(FacebookPtrOutput)
}

func (o IdentityProvidersPtrOutput) GitHub() GitHubPtrOutput {
	return o.ApplyT(func(v *IdentityProviders) *GitHub {
		if v == nil {
			return nil
		}
		return v.GitHub
	}).(GitHubPtrOutput)
}

func (o IdentityProvidersPtrOutput) Google() GooglePtrOutput {
	return o.ApplyT(func(v *IdentityProviders) *Google {
		if v == nil {
			return nil
		}
		return v.Google
	}).(GooglePtrOutput)
}

func (o IdentityProvidersPtrOutput) LegacyMicrosoftAccount() LegacyMicrosoftAccountPtrOutput {
	return o.ApplyT(func(v *IdentityProviders) *LegacyMicrosoftAccount {
		if v == nil {
			return nil
		}
		return v.LegacyMicrosoftAccount
	}).(LegacyMicrosoftAccountPtrOutput)
}

func (o IdentityProvidersPtrOutput) Twitter() TwitterPtrOutput {
	return o.ApplyT(func(v *IdentityProviders) *Twitter {
		if v == nil {
			return nil
		}
		return v.Twitter
	}).(TwitterPtrOutput)
}

type IdentityProvidersResponse struct {
	Apple                        *AppleResponse                                 `pulumi:"apple"`
	AzureActiveDirectory         *AzureActiveDirectoryResponse                  `pulumi:"azureActiveDirectory"`
	AzureStaticWebApps           *AzureStaticWebAppsResponse                    `pulumi:"azureStaticWebApps"`
	CustomOpenIdConnectProviders map[string]CustomOpenIdConnectProviderResponse `pulumi:"customOpenIdConnectProviders"`
	Facebook                     *FacebookResponse                              `pulumi:"facebook"`
	GitHub                       *GitHubResponse                                `pulumi:"gitHub"`
	Google                       *GoogleResponse                                `pulumi:"google"`
	LegacyMicrosoftAccount       *LegacyMicrosoftAccountResponse                `pulumi:"legacyMicrosoftAccount"`
	Twitter                      *TwitterResponse                               `pulumi:"twitter"`
}

type IdentityProvidersResponseOutput struct{ *pulumi.OutputState }

func (IdentityProvidersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProvidersResponse)(nil)).Elem()
}

func (o IdentityProvidersResponseOutput) ToIdentityProvidersResponseOutput() IdentityProvidersResponseOutput {
	return o
}

func (o IdentityProvidersResponseOutput) ToIdentityProvidersResponseOutputWithContext(ctx context.Context) IdentityProvidersResponseOutput {
	return o
}

func (o IdentityProvidersResponseOutput) Apple() AppleResponsePtrOutput {
	return o.ApplyT(func(v IdentityProvidersResponse) *AppleResponse { return v.Apple }).(AppleResponsePtrOutput)
}

func (o IdentityProvidersResponseOutput) AzureActiveDirectory() AzureActiveDirectoryResponsePtrOutput {
	return o.ApplyT(func(v IdentityProvidersResponse) *AzureActiveDirectoryResponse { return v.AzureActiveDirectory }).(AzureActiveDirectoryResponsePtrOutput)
}

func (o IdentityProvidersResponseOutput) AzureStaticWebApps() AzureStaticWebAppsResponsePtrOutput {
	return o.ApplyT(func(v IdentityProvidersResponse) *AzureStaticWebAppsResponse { return v.AzureStaticWebApps }).(AzureStaticWebAppsResponsePtrOutput)
}

func (o IdentityProvidersResponseOutput) CustomOpenIdConnectProviders() CustomOpenIdConnectProviderResponseMapOutput {
	return o.ApplyT(func(v IdentityProvidersResponse) map[string]CustomOpenIdConnectProviderResponse {
		return v.CustomOpenIdConnectProviders
	}).(CustomOpenIdConnectProviderResponseMapOutput)
}

func (o IdentityProvidersResponseOutput) Facebook() FacebookResponsePtrOutput {
	return o.ApplyT(func(v IdentityProvidersResponse) *FacebookResponse { return v.Facebook }).(FacebookResponsePtrOutput)
}

func (o IdentityProvidersResponseOutput) GitHub() GitHubResponsePtrOutput {
	return o.ApplyT(func(v IdentityProvidersResponse) *GitHubResponse { return v.GitHub }).(GitHubResponsePtrOutput)
}

func (o IdentityProvidersResponseOutput) Google() GoogleResponsePtrOutput {
	return o.ApplyT(func(v IdentityProvidersResponse) *GoogleResponse { return v.Google }).(GoogleResponsePtrOutput)
}

func (o IdentityProvidersResponseOutput) LegacyMicrosoftAccount() LegacyMicrosoftAccountResponsePtrOutput {
	return o.ApplyT(func(v IdentityProvidersResponse) *LegacyMicrosoftAccountResponse { return v.LegacyMicrosoftAccount }).(LegacyMicrosoftAccountResponsePtrOutput)
}

func (o IdentityProvidersResponseOutput) Twitter() TwitterResponsePtrOutput {
	return o.ApplyT(func(v IdentityProvidersResponse) *TwitterResponse { return v.Twitter }).(TwitterResponsePtrOutput)
}

type IdentityProvidersResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityProvidersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProvidersResponse)(nil)).Elem()
}

func (o IdentityProvidersResponsePtrOutput) ToIdentityProvidersResponsePtrOutput() IdentityProvidersResponsePtrOutput {
	return o
}

func (o IdentityProvidersResponsePtrOutput) ToIdentityProvidersResponsePtrOutputWithContext(ctx context.Context) IdentityProvidersResponsePtrOutput {
	return o
}

func (o IdentityProvidersResponsePtrOutput) Elem() IdentityProvidersResponseOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) IdentityProvidersResponse {
		if v != nil {
			return *v
		}
		var ret IdentityProvidersResponse
		return ret
	}).(IdentityProvidersResponseOutput)
}

func (o IdentityProvidersResponsePtrOutput) Apple() AppleResponsePtrOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) *AppleResponse {
		if v == nil {
			return nil
		}
		return v.Apple
	}).(AppleResponsePtrOutput)
}

func (o IdentityProvidersResponsePtrOutput) AzureActiveDirectory() AzureActiveDirectoryResponsePtrOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) *AzureActiveDirectoryResponse {
		if v == nil {
			return nil
		}
		return v.AzureActiveDirectory
	}).(AzureActiveDirectoryResponsePtrOutput)
}

func (o IdentityProvidersResponsePtrOutput) AzureStaticWebApps() AzureStaticWebAppsResponsePtrOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) *AzureStaticWebAppsResponse {
		if v == nil {
			return nil
		}
		return v.AzureStaticWebApps
	}).(AzureStaticWebAppsResponsePtrOutput)
}

func (o IdentityProvidersResponsePtrOutput) CustomOpenIdConnectProviders() CustomOpenIdConnectProviderResponseMapOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) map[string]CustomOpenIdConnectProviderResponse {
		if v == nil {
			return nil
		}
		return v.CustomOpenIdConnectProviders
	}).(CustomOpenIdConnectProviderResponseMapOutput)
}

func (o IdentityProvidersResponsePtrOutput) Facebook() FacebookResponsePtrOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) *FacebookResponse {
		if v == nil {
			return nil
		}
		return v.Facebook
	}).(FacebookResponsePtrOutput)
}

func (o IdentityProvidersResponsePtrOutput) GitHub() GitHubResponsePtrOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) *GitHubResponse {
		if v == nil {
			return nil
		}
		return v.GitHub
	}).(GitHubResponsePtrOutput)
}

func (o IdentityProvidersResponsePtrOutput) Google() GoogleResponsePtrOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) *GoogleResponse {
		if v == nil {
			return nil
		}
		return v.Google
	}).(GoogleResponsePtrOutput)
}

func (o IdentityProvidersResponsePtrOutput) LegacyMicrosoftAccount() LegacyMicrosoftAccountResponsePtrOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) *LegacyMicrosoftAccountResponse {
		if v == nil {
			return nil
		}
		return v.LegacyMicrosoftAccount
	}).(LegacyMicrosoftAccountResponsePtrOutput)
}

func (o IdentityProvidersResponsePtrOutput) Twitter() TwitterResponsePtrOutput {
	return o.ApplyT(func(v *IdentityProvidersResponse) *TwitterResponse {
		if v == nil {
			return nil
		}
		return v.Twitter
	}).(TwitterResponsePtrOutput)
}

type IpSecurityRestriction struct {
	Action               *string             `pulumi:"action"`
	Description          *string             `pulumi:"description"`
	Headers              map[string][]string `pulumi:"headers"`
	IpAddress            *string             `pulumi:"ipAddress"`
	Name                 *string             `pulumi:"name"`
	Priority             *int                `pulumi:"priority"`
	SubnetMask           *string             `pulumi:"subnetMask"`
	SubnetTrafficTag     *int                `pulumi:"subnetTrafficTag"`
	Tag                  *string             `pulumi:"tag"`
	VnetSubnetResourceId *string             `pulumi:"vnetSubnetResourceId"`
	VnetTrafficTag       *int                `pulumi:"vnetTrafficTag"`
}





type IpSecurityRestrictionInput interface {
	pulumi.Input

	ToIpSecurityRestrictionOutput() IpSecurityRestrictionOutput
	ToIpSecurityRestrictionOutputWithContext(context.Context) IpSecurityRestrictionOutput
}

type IpSecurityRestrictionArgs struct {
	Action               pulumi.StringPtrInput      `pulumi:"action"`
	Description          pulumi.StringPtrInput      `pulumi:"description"`
	Headers              pulumi.StringArrayMapInput `pulumi:"headers"`
	IpAddress            pulumi.StringPtrInput      `pulumi:"ipAddress"`
	Name                 pulumi.StringPtrInput      `pulumi:"name"`
	Priority             pulumi.IntPtrInput         `pulumi:"priority"`
	SubnetMask           pulumi.StringPtrInput      `pulumi:"subnetMask"`
	SubnetTrafficTag     pulumi.IntPtrInput         `pulumi:"subnetTrafficTag"`
	Tag                  pulumi.StringPtrInput      `pulumi:"tag"`
	VnetSubnetResourceId pulumi.StringPtrInput      `pulumi:"vnetSubnetResourceId"`
	VnetTrafficTag       pulumi.IntPtrInput         `pulumi:"vnetTrafficTag"`
}

func (IpSecurityRestrictionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpSecurityRestriction)(nil)).Elem()
}

func (i IpSecurityRestrictionArgs) ToIpSecurityRestrictionOutput() IpSecurityRestrictionOutput {
	return i.ToIpSecurityRestrictionOutputWithContext(context.Background())
}

func (i IpSecurityRestrictionArgs) ToIpSecurityRestrictionOutputWithContext(ctx context.Context) IpSecurityRestrictionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSecurityRestrictionOutput)
}





type IpSecurityRestrictionArrayInput interface {
	pulumi.Input

	ToIpSecurityRestrictionArrayOutput() IpSecurityRestrictionArrayOutput
	ToIpSecurityRestrictionArrayOutputWithContext(context.Context) IpSecurityRestrictionArrayOutput
}

type IpSecurityRestrictionArray []IpSecurityRestrictionInput

func (IpSecurityRestrictionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpSecurityRestriction)(nil)).Elem()
}

func (i IpSecurityRestrictionArray) ToIpSecurityRestrictionArrayOutput() IpSecurityRestrictionArrayOutput {
	return i.ToIpSecurityRestrictionArrayOutputWithContext(context.Background())
}

func (i IpSecurityRestrictionArray) ToIpSecurityRestrictionArrayOutputWithContext(ctx context.Context) IpSecurityRestrictionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSecurityRestrictionArrayOutput)
}

type IpSecurityRestrictionOutput struct{ *pulumi.OutputState }

func (IpSecurityRestrictionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpSecurityRestriction)(nil)).Elem()
}

func (o IpSecurityRestrictionOutput) ToIpSecurityRestrictionOutput() IpSecurityRestrictionOutput {
	return o
}

func (o IpSecurityRestrictionOutput) ToIpSecurityRestrictionOutputWithContext(ctx context.Context) IpSecurityRestrictionOutput {
	return o
}

func (o IpSecurityRestrictionOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionOutput) Headers() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v IpSecurityRestriction) map[string][]string { return v.Headers }).(pulumi.StringArrayMapOutput)
}

func (o IpSecurityRestrictionOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o IpSecurityRestrictionOutput) SubnetMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *string { return v.SubnetMask }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionOutput) SubnetTrafficTag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *int { return v.SubnetTrafficTag }).(pulumi.IntPtrOutput)
}

func (o IpSecurityRestrictionOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionOutput) VnetSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *string { return v.VnetSubnetResourceId }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionOutput) VnetTrafficTag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *int { return v.VnetTrafficTag }).(pulumi.IntPtrOutput)
}

type IpSecurityRestrictionArrayOutput struct{ *pulumi.OutputState }

func (IpSecurityRestrictionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpSecurityRestriction)(nil)).Elem()
}

func (o IpSecurityRestrictionArrayOutput) ToIpSecurityRestrictionArrayOutput() IpSecurityRestrictionArrayOutput {
	return o
}

func (o IpSecurityRestrictionArrayOutput) ToIpSecurityRestrictionArrayOutputWithContext(ctx context.Context) IpSecurityRestrictionArrayOutput {
	return o
}

func (o IpSecurityRestrictionArrayOutput) Index(i pulumi.IntInput) IpSecurityRestrictionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpSecurityRestriction {
		return vs[0].([]IpSecurityRestriction)[vs[1].(int)]
	}).(IpSecurityRestrictionOutput)
}

type IpSecurityRestrictionResponse struct {
	Action               *string             `pulumi:"action"`
	Description          *string             `pulumi:"description"`
	Headers              map[string][]string `pulumi:"headers"`
	IpAddress            *string             `pulumi:"ipAddress"`
	Name                 *string             `pulumi:"name"`
	Priority             *int                `pulumi:"priority"`
	SubnetMask           *string             `pulumi:"subnetMask"`
	SubnetTrafficTag     *int                `pulumi:"subnetTrafficTag"`
	Tag                  *string             `pulumi:"tag"`
	VnetSubnetResourceId *string             `pulumi:"vnetSubnetResourceId"`
	VnetTrafficTag       *int                `pulumi:"vnetTrafficTag"`
}

type IpSecurityRestrictionResponseOutput struct{ *pulumi.OutputState }

func (IpSecurityRestrictionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpSecurityRestrictionResponse)(nil)).Elem()
}

func (o IpSecurityRestrictionResponseOutput) ToIpSecurityRestrictionResponseOutput() IpSecurityRestrictionResponseOutput {
	return o
}

func (o IpSecurityRestrictionResponseOutput) ToIpSecurityRestrictionResponseOutputWithContext(ctx context.Context) IpSecurityRestrictionResponseOutput {
	return o
}

func (o IpSecurityRestrictionResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) Headers() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) map[string][]string { return v.Headers }).(pulumi.StringArrayMapOutput)
}

func (o IpSecurityRestrictionResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) SubnetMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *string { return v.SubnetMask }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) SubnetTrafficTag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *int { return v.SubnetTrafficTag }).(pulumi.IntPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) VnetSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *string { return v.VnetSubnetResourceId }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) VnetTrafficTag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *int { return v.VnetTrafficTag }).(pulumi.IntPtrOutput)
}

type IpSecurityRestrictionResponseArrayOutput struct{ *pulumi.OutputState }

func (IpSecurityRestrictionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpSecurityRestrictionResponse)(nil)).Elem()
}

func (o IpSecurityRestrictionResponseArrayOutput) ToIpSecurityRestrictionResponseArrayOutput() IpSecurityRestrictionResponseArrayOutput {
	return o
}

func (o IpSecurityRestrictionResponseArrayOutput) ToIpSecurityRestrictionResponseArrayOutputWithContext(ctx context.Context) IpSecurityRestrictionResponseArrayOutput {
	return o
}

func (o IpSecurityRestrictionResponseArrayOutput) Index(i pulumi.IntInput) IpSecurityRestrictionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpSecurityRestrictionResponse {
		return vs[0].([]IpSecurityRestrictionResponse)[vs[1].(int)]
	}).(IpSecurityRestrictionResponseOutput)
}

type JwtClaimChecks struct {
	AllowedClientApplications []string `pulumi:"allowedClientApplications"`
	AllowedGroups             []string `pulumi:"allowedGroups"`
}





type JwtClaimChecksInput interface {
	pulumi.Input

	ToJwtClaimChecksOutput() JwtClaimChecksOutput
	ToJwtClaimChecksOutputWithContext(context.Context) JwtClaimChecksOutput
}

type JwtClaimChecksArgs struct {
	AllowedClientApplications pulumi.StringArrayInput `pulumi:"allowedClientApplications"`
	AllowedGroups             pulumi.StringArrayInput `pulumi:"allowedGroups"`
}

func (JwtClaimChecksArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JwtClaimChecks)(nil)).Elem()
}

func (i JwtClaimChecksArgs) ToJwtClaimChecksOutput() JwtClaimChecksOutput {
	return i.ToJwtClaimChecksOutputWithContext(context.Background())
}

func (i JwtClaimChecksArgs) ToJwtClaimChecksOutputWithContext(ctx context.Context) JwtClaimChecksOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JwtClaimChecksOutput)
}

func (i JwtClaimChecksArgs) ToJwtClaimChecksPtrOutput() JwtClaimChecksPtrOutput {
	return i.ToJwtClaimChecksPtrOutputWithContext(context.Background())
}

func (i JwtClaimChecksArgs) ToJwtClaimChecksPtrOutputWithContext(ctx context.Context) JwtClaimChecksPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JwtClaimChecksOutput).ToJwtClaimChecksPtrOutputWithContext(ctx)
}









type JwtClaimChecksPtrInput interface {
	pulumi.Input

	ToJwtClaimChecksPtrOutput() JwtClaimChecksPtrOutput
	ToJwtClaimChecksPtrOutputWithContext(context.Context) JwtClaimChecksPtrOutput
}

type jwtClaimChecksPtrType JwtClaimChecksArgs

func JwtClaimChecksPtr(v *JwtClaimChecksArgs) JwtClaimChecksPtrInput {
	return (*jwtClaimChecksPtrType)(v)
}

func (*jwtClaimChecksPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JwtClaimChecks)(nil)).Elem()
}

func (i *jwtClaimChecksPtrType) ToJwtClaimChecksPtrOutput() JwtClaimChecksPtrOutput {
	return i.ToJwtClaimChecksPtrOutputWithContext(context.Background())
}

func (i *jwtClaimChecksPtrType) ToJwtClaimChecksPtrOutputWithContext(ctx context.Context) JwtClaimChecksPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JwtClaimChecksPtrOutput)
}

type JwtClaimChecksOutput struct{ *pulumi.OutputState }

func (JwtClaimChecksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JwtClaimChecks)(nil)).Elem()
}

func (o JwtClaimChecksOutput) ToJwtClaimChecksOutput() JwtClaimChecksOutput {
	return o
}

func (o JwtClaimChecksOutput) ToJwtClaimChecksOutputWithContext(ctx context.Context) JwtClaimChecksOutput {
	return o
}

func (o JwtClaimChecksOutput) ToJwtClaimChecksPtrOutput() JwtClaimChecksPtrOutput {
	return o.ToJwtClaimChecksPtrOutputWithContext(context.Background())
}

func (o JwtClaimChecksOutput) ToJwtClaimChecksPtrOutputWithContext(ctx context.Context) JwtClaimChecksPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JwtClaimChecks) *JwtClaimChecks {
		return &v
	}).(JwtClaimChecksPtrOutput)
}

func (o JwtClaimChecksOutput) AllowedClientApplications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JwtClaimChecks) []string { return v.AllowedClientApplications }).(pulumi.StringArrayOutput)
}

func (o JwtClaimChecksOutput) AllowedGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JwtClaimChecks) []string { return v.AllowedGroups }).(pulumi.StringArrayOutput)
}

type JwtClaimChecksPtrOutput struct{ *pulumi.OutputState }

func (JwtClaimChecksPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JwtClaimChecks)(nil)).Elem()
}

func (o JwtClaimChecksPtrOutput) ToJwtClaimChecksPtrOutput() JwtClaimChecksPtrOutput {
	return o
}

func (o JwtClaimChecksPtrOutput) ToJwtClaimChecksPtrOutputWithContext(ctx context.Context) JwtClaimChecksPtrOutput {
	return o
}

func (o JwtClaimChecksPtrOutput) Elem() JwtClaimChecksOutput {
	return o.ApplyT(func(v *JwtClaimChecks) JwtClaimChecks {
		if v != nil {
			return *v
		}
		var ret JwtClaimChecks
		return ret
	}).(JwtClaimChecksOutput)
}

func (o JwtClaimChecksPtrOutput) AllowedClientApplications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JwtClaimChecks) []string {
		if v == nil {
			return nil
		}
		return v.AllowedClientApplications
	}).(pulumi.StringArrayOutput)
}

func (o JwtClaimChecksPtrOutput) AllowedGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JwtClaimChecks) []string {
		if v == nil {
			return nil
		}
		return v.AllowedGroups
	}).(pulumi.StringArrayOutput)
}

type JwtClaimChecksResponse struct {
	AllowedClientApplications []string `pulumi:"allowedClientApplications"`
	AllowedGroups             []string `pulumi:"allowedGroups"`
}

type JwtClaimChecksResponseOutput struct{ *pulumi.OutputState }

func (JwtClaimChecksResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JwtClaimChecksResponse)(nil)).Elem()
}

func (o JwtClaimChecksResponseOutput) ToJwtClaimChecksResponseOutput() JwtClaimChecksResponseOutput {
	return o
}

func (o JwtClaimChecksResponseOutput) ToJwtClaimChecksResponseOutputWithContext(ctx context.Context) JwtClaimChecksResponseOutput {
	return o
}

func (o JwtClaimChecksResponseOutput) AllowedClientApplications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JwtClaimChecksResponse) []string { return v.AllowedClientApplications }).(pulumi.StringArrayOutput)
}

func (o JwtClaimChecksResponseOutput) AllowedGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JwtClaimChecksResponse) []string { return v.AllowedGroups }).(pulumi.StringArrayOutput)
}

type JwtClaimChecksResponsePtrOutput struct{ *pulumi.OutputState }

func (JwtClaimChecksResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JwtClaimChecksResponse)(nil)).Elem()
}

func (o JwtClaimChecksResponsePtrOutput) ToJwtClaimChecksResponsePtrOutput() JwtClaimChecksResponsePtrOutput {
	return o
}

func (o JwtClaimChecksResponsePtrOutput) ToJwtClaimChecksResponsePtrOutputWithContext(ctx context.Context) JwtClaimChecksResponsePtrOutput {
	return o
}

func (o JwtClaimChecksResponsePtrOutput) Elem() JwtClaimChecksResponseOutput {
	return o.ApplyT(func(v *JwtClaimChecksResponse) JwtClaimChecksResponse {
		if v != nil {
			return *v
		}
		var ret JwtClaimChecksResponse
		return ret
	}).(JwtClaimChecksResponseOutput)
}

func (o JwtClaimChecksResponsePtrOutput) AllowedClientApplications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JwtClaimChecksResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedClientApplications
	}).(pulumi.StringArrayOutput)
}

func (o JwtClaimChecksResponsePtrOutput) AllowedGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JwtClaimChecksResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedGroups
	}).(pulumi.StringArrayOutput)
}

type KubeEnvironmentProfile struct {
	Id *string `pulumi:"id"`
}





type KubeEnvironmentProfileInput interface {
	pulumi.Input

	ToKubeEnvironmentProfileOutput() KubeEnvironmentProfileOutput
	ToKubeEnvironmentProfileOutputWithContext(context.Context) KubeEnvironmentProfileOutput
}

type KubeEnvironmentProfileArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (KubeEnvironmentProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeEnvironmentProfile)(nil)).Elem()
}

func (i KubeEnvironmentProfileArgs) ToKubeEnvironmentProfileOutput() KubeEnvironmentProfileOutput {
	return i.ToKubeEnvironmentProfileOutputWithContext(context.Background())
}

func (i KubeEnvironmentProfileArgs) ToKubeEnvironmentProfileOutputWithContext(ctx context.Context) KubeEnvironmentProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeEnvironmentProfileOutput)
}

func (i KubeEnvironmentProfileArgs) ToKubeEnvironmentProfilePtrOutput() KubeEnvironmentProfilePtrOutput {
	return i.ToKubeEnvironmentProfilePtrOutputWithContext(context.Background())
}

func (i KubeEnvironmentProfileArgs) ToKubeEnvironmentProfilePtrOutputWithContext(ctx context.Context) KubeEnvironmentProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeEnvironmentProfileOutput).ToKubeEnvironmentProfilePtrOutputWithContext(ctx)
}









type KubeEnvironmentProfilePtrInput interface {
	pulumi.Input

	ToKubeEnvironmentProfilePtrOutput() KubeEnvironmentProfilePtrOutput
	ToKubeEnvironmentProfilePtrOutputWithContext(context.Context) KubeEnvironmentProfilePtrOutput
}

type kubeEnvironmentProfilePtrType KubeEnvironmentProfileArgs

func KubeEnvironmentProfilePtr(v *KubeEnvironmentProfileArgs) KubeEnvironmentProfilePtrInput {
	return (*kubeEnvironmentProfilePtrType)(v)
}

func (*kubeEnvironmentProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeEnvironmentProfile)(nil)).Elem()
}

func (i *kubeEnvironmentProfilePtrType) ToKubeEnvironmentProfilePtrOutput() KubeEnvironmentProfilePtrOutput {
	return i.ToKubeEnvironmentProfilePtrOutputWithContext(context.Background())
}

func (i *kubeEnvironmentProfilePtrType) ToKubeEnvironmentProfilePtrOutputWithContext(ctx context.Context) KubeEnvironmentProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeEnvironmentProfilePtrOutput)
}

type KubeEnvironmentProfileOutput struct{ *pulumi.OutputState }

func (KubeEnvironmentProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeEnvironmentProfile)(nil)).Elem()
}

func (o KubeEnvironmentProfileOutput) ToKubeEnvironmentProfileOutput() KubeEnvironmentProfileOutput {
	return o
}

func (o KubeEnvironmentProfileOutput) ToKubeEnvironmentProfileOutputWithContext(ctx context.Context) KubeEnvironmentProfileOutput {
	return o
}

func (o KubeEnvironmentProfileOutput) ToKubeEnvironmentProfilePtrOutput() KubeEnvironmentProfilePtrOutput {
	return o.ToKubeEnvironmentProfilePtrOutputWithContext(context.Background())
}

func (o KubeEnvironmentProfileOutput) ToKubeEnvironmentProfilePtrOutputWithContext(ctx context.Context) KubeEnvironmentProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubeEnvironmentProfile) *KubeEnvironmentProfile {
		return &v
	}).(KubeEnvironmentProfilePtrOutput)
}

func (o KubeEnvironmentProfileOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubeEnvironmentProfile) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type KubeEnvironmentProfilePtrOutput struct{ *pulumi.OutputState }

func (KubeEnvironmentProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeEnvironmentProfile)(nil)).Elem()
}

func (o KubeEnvironmentProfilePtrOutput) ToKubeEnvironmentProfilePtrOutput() KubeEnvironmentProfilePtrOutput {
	return o
}

func (o KubeEnvironmentProfilePtrOutput) ToKubeEnvironmentProfilePtrOutputWithContext(ctx context.Context) KubeEnvironmentProfilePtrOutput {
	return o
}

func (o KubeEnvironmentProfilePtrOutput) Elem() KubeEnvironmentProfileOutput {
	return o.ApplyT(func(v *KubeEnvironmentProfile) KubeEnvironmentProfile {
		if v != nil {
			return *v
		}
		var ret KubeEnvironmentProfile
		return ret
	}).(KubeEnvironmentProfileOutput)
}

func (o KubeEnvironmentProfilePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeEnvironmentProfile) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type KubeEnvironmentProfileResponse struct {
	Id   *string `pulumi:"id"`
	Name string  `pulumi:"name"`
	Type string  `pulumi:"type"`
}

type KubeEnvironmentProfileResponseOutput struct{ *pulumi.OutputState }

func (KubeEnvironmentProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeEnvironmentProfileResponse)(nil)).Elem()
}

func (o KubeEnvironmentProfileResponseOutput) ToKubeEnvironmentProfileResponseOutput() KubeEnvironmentProfileResponseOutput {
	return o
}

func (o KubeEnvironmentProfileResponseOutput) ToKubeEnvironmentProfileResponseOutputWithContext(ctx context.Context) KubeEnvironmentProfileResponseOutput {
	return o
}

func (o KubeEnvironmentProfileResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubeEnvironmentProfileResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o KubeEnvironmentProfileResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v KubeEnvironmentProfileResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o KubeEnvironmentProfileResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v KubeEnvironmentProfileResponse) string { return v.Type }).(pulumi.StringOutput)
}

type KubeEnvironmentProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (KubeEnvironmentProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeEnvironmentProfileResponse)(nil)).Elem()
}

func (o KubeEnvironmentProfileResponsePtrOutput) ToKubeEnvironmentProfileResponsePtrOutput() KubeEnvironmentProfileResponsePtrOutput {
	return o
}

func (o KubeEnvironmentProfileResponsePtrOutput) ToKubeEnvironmentProfileResponsePtrOutputWithContext(ctx context.Context) KubeEnvironmentProfileResponsePtrOutput {
	return o
}

func (o KubeEnvironmentProfileResponsePtrOutput) Elem() KubeEnvironmentProfileResponseOutput {
	return o.ApplyT(func(v *KubeEnvironmentProfileResponse) KubeEnvironmentProfileResponse {
		if v != nil {
			return *v
		}
		var ret KubeEnvironmentProfileResponse
		return ret
	}).(KubeEnvironmentProfileResponseOutput)
}

func (o KubeEnvironmentProfileResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o KubeEnvironmentProfileResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o KubeEnvironmentProfileResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type LegacyMicrosoftAccount struct {
	Enabled      *bool                       `pulumi:"enabled"`
	Login        *LoginScopes                `pulumi:"login"`
	Registration *ClientRegistration         `pulumi:"registration"`
	Validation   *AllowedAudiencesValidation `pulumi:"validation"`
}





type LegacyMicrosoftAccountInput interface {
	pulumi.Input

	ToLegacyMicrosoftAccountOutput() LegacyMicrosoftAccountOutput
	ToLegacyMicrosoftAccountOutputWithContext(context.Context) LegacyMicrosoftAccountOutput
}

type LegacyMicrosoftAccountArgs struct {
	Enabled      pulumi.BoolPtrInput                `pulumi:"enabled"`
	Login        LoginScopesPtrInput                `pulumi:"login"`
	Registration ClientRegistrationPtrInput         `pulumi:"registration"`
	Validation   AllowedAudiencesValidationPtrInput `pulumi:"validation"`
}

func (LegacyMicrosoftAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LegacyMicrosoftAccount)(nil)).Elem()
}

func (i LegacyMicrosoftAccountArgs) ToLegacyMicrosoftAccountOutput() LegacyMicrosoftAccountOutput {
	return i.ToLegacyMicrosoftAccountOutputWithContext(context.Background())
}

func (i LegacyMicrosoftAccountArgs) ToLegacyMicrosoftAccountOutputWithContext(ctx context.Context) LegacyMicrosoftAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LegacyMicrosoftAccountOutput)
}

func (i LegacyMicrosoftAccountArgs) ToLegacyMicrosoftAccountPtrOutput() LegacyMicrosoftAccountPtrOutput {
	return i.ToLegacyMicrosoftAccountPtrOutputWithContext(context.Background())
}

func (i LegacyMicrosoftAccountArgs) ToLegacyMicrosoftAccountPtrOutputWithContext(ctx context.Context) LegacyMicrosoftAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LegacyMicrosoftAccountOutput).ToLegacyMicrosoftAccountPtrOutputWithContext(ctx)
}









type LegacyMicrosoftAccountPtrInput interface {
	pulumi.Input

	ToLegacyMicrosoftAccountPtrOutput() LegacyMicrosoftAccountPtrOutput
	ToLegacyMicrosoftAccountPtrOutputWithContext(context.Context) LegacyMicrosoftAccountPtrOutput
}

type legacyMicrosoftAccountPtrType LegacyMicrosoftAccountArgs

func LegacyMicrosoftAccountPtr(v *LegacyMicrosoftAccountArgs) LegacyMicrosoftAccountPtrInput {
	return (*legacyMicrosoftAccountPtrType)(v)
}

func (*legacyMicrosoftAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LegacyMicrosoftAccount)(nil)).Elem()
}

func (i *legacyMicrosoftAccountPtrType) ToLegacyMicrosoftAccountPtrOutput() LegacyMicrosoftAccountPtrOutput {
	return i.ToLegacyMicrosoftAccountPtrOutputWithContext(context.Background())
}

func (i *legacyMicrosoftAccountPtrType) ToLegacyMicrosoftAccountPtrOutputWithContext(ctx context.Context) LegacyMicrosoftAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LegacyMicrosoftAccountPtrOutput)
}

type LegacyMicrosoftAccountOutput struct{ *pulumi.OutputState }

func (LegacyMicrosoftAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LegacyMicrosoftAccount)(nil)).Elem()
}

func (o LegacyMicrosoftAccountOutput) ToLegacyMicrosoftAccountOutput() LegacyMicrosoftAccountOutput {
	return o
}

func (o LegacyMicrosoftAccountOutput) ToLegacyMicrosoftAccountOutputWithContext(ctx context.Context) LegacyMicrosoftAccountOutput {
	return o
}

func (o LegacyMicrosoftAccountOutput) ToLegacyMicrosoftAccountPtrOutput() LegacyMicrosoftAccountPtrOutput {
	return o.ToLegacyMicrosoftAccountPtrOutputWithContext(context.Background())
}

func (o LegacyMicrosoftAccountOutput) ToLegacyMicrosoftAccountPtrOutputWithContext(ctx context.Context) LegacyMicrosoftAccountPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LegacyMicrosoftAccount) *LegacyMicrosoftAccount {
		return &v
	}).(LegacyMicrosoftAccountPtrOutput)
}

func (o LegacyMicrosoftAccountOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LegacyMicrosoftAccount) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LegacyMicrosoftAccountOutput) Login() LoginScopesPtrOutput {
	return o.ApplyT(func(v LegacyMicrosoftAccount) *LoginScopes { return v.Login }).(LoginScopesPtrOutput)
}

func (o LegacyMicrosoftAccountOutput) Registration() ClientRegistrationPtrOutput {
	return o.ApplyT(func(v LegacyMicrosoftAccount) *ClientRegistration { return v.Registration }).(ClientRegistrationPtrOutput)
}

func (o LegacyMicrosoftAccountOutput) Validation() AllowedAudiencesValidationPtrOutput {
	return o.ApplyT(func(v LegacyMicrosoftAccount) *AllowedAudiencesValidation { return v.Validation }).(AllowedAudiencesValidationPtrOutput)
}

type LegacyMicrosoftAccountPtrOutput struct{ *pulumi.OutputState }

func (LegacyMicrosoftAccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LegacyMicrosoftAccount)(nil)).Elem()
}

func (o LegacyMicrosoftAccountPtrOutput) ToLegacyMicrosoftAccountPtrOutput() LegacyMicrosoftAccountPtrOutput {
	return o
}

func (o LegacyMicrosoftAccountPtrOutput) ToLegacyMicrosoftAccountPtrOutputWithContext(ctx context.Context) LegacyMicrosoftAccountPtrOutput {
	return o
}

func (o LegacyMicrosoftAccountPtrOutput) Elem() LegacyMicrosoftAccountOutput {
	return o.ApplyT(func(v *LegacyMicrosoftAccount) LegacyMicrosoftAccount {
		if v != nil {
			return *v
		}
		var ret LegacyMicrosoftAccount
		return ret
	}).(LegacyMicrosoftAccountOutput)
}

func (o LegacyMicrosoftAccountPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LegacyMicrosoftAccount) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o LegacyMicrosoftAccountPtrOutput) Login() LoginScopesPtrOutput {
	return o.ApplyT(func(v *LegacyMicrosoftAccount) *LoginScopes {
		if v == nil {
			return nil
		}
		return v.Login
	}).(LoginScopesPtrOutput)
}

func (o LegacyMicrosoftAccountPtrOutput) Registration() ClientRegistrationPtrOutput {
	return o.ApplyT(func(v *LegacyMicrosoftAccount) *ClientRegistration {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(ClientRegistrationPtrOutput)
}

func (o LegacyMicrosoftAccountPtrOutput) Validation() AllowedAudiencesValidationPtrOutput {
	return o.ApplyT(func(v *LegacyMicrosoftAccount) *AllowedAudiencesValidation {
		if v == nil {
			return nil
		}
		return v.Validation
	}).(AllowedAudiencesValidationPtrOutput)
}

type LegacyMicrosoftAccountResponse struct {
	Enabled      *bool                               `pulumi:"enabled"`
	Login        *LoginScopesResponse                `pulumi:"login"`
	Registration *ClientRegistrationResponse         `pulumi:"registration"`
	Validation   *AllowedAudiencesValidationResponse `pulumi:"validation"`
}

type LegacyMicrosoftAccountResponseOutput struct{ *pulumi.OutputState }

func (LegacyMicrosoftAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LegacyMicrosoftAccountResponse)(nil)).Elem()
}

func (o LegacyMicrosoftAccountResponseOutput) ToLegacyMicrosoftAccountResponseOutput() LegacyMicrosoftAccountResponseOutput {
	return o
}

func (o LegacyMicrosoftAccountResponseOutput) ToLegacyMicrosoftAccountResponseOutputWithContext(ctx context.Context) LegacyMicrosoftAccountResponseOutput {
	return o
}

func (o LegacyMicrosoftAccountResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LegacyMicrosoftAccountResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LegacyMicrosoftAccountResponseOutput) Login() LoginScopesResponsePtrOutput {
	return o.ApplyT(func(v LegacyMicrosoftAccountResponse) *LoginScopesResponse { return v.Login }).(LoginScopesResponsePtrOutput)
}

func (o LegacyMicrosoftAccountResponseOutput) Registration() ClientRegistrationResponsePtrOutput {
	return o.ApplyT(func(v LegacyMicrosoftAccountResponse) *ClientRegistrationResponse { return v.Registration }).(ClientRegistrationResponsePtrOutput)
}

func (o LegacyMicrosoftAccountResponseOutput) Validation() AllowedAudiencesValidationResponsePtrOutput {
	return o.ApplyT(func(v LegacyMicrosoftAccountResponse) *AllowedAudiencesValidationResponse { return v.Validation }).(AllowedAudiencesValidationResponsePtrOutput)
}

type LegacyMicrosoftAccountResponsePtrOutput struct{ *pulumi.OutputState }

func (LegacyMicrosoftAccountResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LegacyMicrosoftAccountResponse)(nil)).Elem()
}

func (o LegacyMicrosoftAccountResponsePtrOutput) ToLegacyMicrosoftAccountResponsePtrOutput() LegacyMicrosoftAccountResponsePtrOutput {
	return o
}

func (o LegacyMicrosoftAccountResponsePtrOutput) ToLegacyMicrosoftAccountResponsePtrOutputWithContext(ctx context.Context) LegacyMicrosoftAccountResponsePtrOutput {
	return o
}

func (o LegacyMicrosoftAccountResponsePtrOutput) Elem() LegacyMicrosoftAccountResponseOutput {
	return o.ApplyT(func(v *LegacyMicrosoftAccountResponse) LegacyMicrosoftAccountResponse {
		if v != nil {
			return *v
		}
		var ret LegacyMicrosoftAccountResponse
		return ret
	}).(LegacyMicrosoftAccountResponseOutput)
}

func (o LegacyMicrosoftAccountResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LegacyMicrosoftAccountResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o LegacyMicrosoftAccountResponsePtrOutput) Login() LoginScopesResponsePtrOutput {
	return o.ApplyT(func(v *LegacyMicrosoftAccountResponse) *LoginScopesResponse {
		if v == nil {
			return nil
		}
		return v.Login
	}).(LoginScopesResponsePtrOutput)
}

func (o LegacyMicrosoftAccountResponsePtrOutput) Registration() ClientRegistrationResponsePtrOutput {
	return o.ApplyT(func(v *LegacyMicrosoftAccountResponse) *ClientRegistrationResponse {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(ClientRegistrationResponsePtrOutput)
}

func (o LegacyMicrosoftAccountResponsePtrOutput) Validation() AllowedAudiencesValidationResponsePtrOutput {
	return o.ApplyT(func(v *LegacyMicrosoftAccountResponse) *AllowedAudiencesValidationResponse {
		if v == nil {
			return nil
		}
		return v.Validation
	}).(AllowedAudiencesValidationResponsePtrOutput)
}

type LogAnalyticsConfiguration struct {
	CustomerId *string `pulumi:"customerId"`
	SharedKey  *string `pulumi:"sharedKey"`
}





type LogAnalyticsConfigurationInput interface {
	pulumi.Input

	ToLogAnalyticsConfigurationOutput() LogAnalyticsConfigurationOutput
	ToLogAnalyticsConfigurationOutputWithContext(context.Context) LogAnalyticsConfigurationOutput
}

type LogAnalyticsConfigurationArgs struct {
	CustomerId pulumi.StringPtrInput `pulumi:"customerId"`
	SharedKey  pulumi.StringPtrInput `pulumi:"sharedKey"`
}

func (LogAnalyticsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsConfiguration)(nil)).Elem()
}

func (i LogAnalyticsConfigurationArgs) ToLogAnalyticsConfigurationOutput() LogAnalyticsConfigurationOutput {
	return i.ToLogAnalyticsConfigurationOutputWithContext(context.Background())
}

func (i LogAnalyticsConfigurationArgs) ToLogAnalyticsConfigurationOutputWithContext(ctx context.Context) LogAnalyticsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsConfigurationOutput)
}

func (i LogAnalyticsConfigurationArgs) ToLogAnalyticsConfigurationPtrOutput() LogAnalyticsConfigurationPtrOutput {
	return i.ToLogAnalyticsConfigurationPtrOutputWithContext(context.Background())
}

func (i LogAnalyticsConfigurationArgs) ToLogAnalyticsConfigurationPtrOutputWithContext(ctx context.Context) LogAnalyticsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsConfigurationOutput).ToLogAnalyticsConfigurationPtrOutputWithContext(ctx)
}









type LogAnalyticsConfigurationPtrInput interface {
	pulumi.Input

	ToLogAnalyticsConfigurationPtrOutput() LogAnalyticsConfigurationPtrOutput
	ToLogAnalyticsConfigurationPtrOutputWithContext(context.Context) LogAnalyticsConfigurationPtrOutput
}

type logAnalyticsConfigurationPtrType LogAnalyticsConfigurationArgs

func LogAnalyticsConfigurationPtr(v *LogAnalyticsConfigurationArgs) LogAnalyticsConfigurationPtrInput {
	return (*logAnalyticsConfigurationPtrType)(v)
}

func (*logAnalyticsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsConfiguration)(nil)).Elem()
}

func (i *logAnalyticsConfigurationPtrType) ToLogAnalyticsConfigurationPtrOutput() LogAnalyticsConfigurationPtrOutput {
	return i.ToLogAnalyticsConfigurationPtrOutputWithContext(context.Background())
}

func (i *logAnalyticsConfigurationPtrType) ToLogAnalyticsConfigurationPtrOutputWithContext(ctx context.Context) LogAnalyticsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsConfigurationPtrOutput)
}

type LogAnalyticsConfigurationOutput struct{ *pulumi.OutputState }

func (LogAnalyticsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsConfiguration)(nil)).Elem()
}

func (o LogAnalyticsConfigurationOutput) ToLogAnalyticsConfigurationOutput() LogAnalyticsConfigurationOutput {
	return o
}

func (o LogAnalyticsConfigurationOutput) ToLogAnalyticsConfigurationOutputWithContext(ctx context.Context) LogAnalyticsConfigurationOutput {
	return o
}

func (o LogAnalyticsConfigurationOutput) ToLogAnalyticsConfigurationPtrOutput() LogAnalyticsConfigurationPtrOutput {
	return o.ToLogAnalyticsConfigurationPtrOutputWithContext(context.Background())
}

func (o LogAnalyticsConfigurationOutput) ToLogAnalyticsConfigurationPtrOutputWithContext(ctx context.Context) LogAnalyticsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogAnalyticsConfiguration) *LogAnalyticsConfiguration {
		return &v
	}).(LogAnalyticsConfigurationPtrOutput)
}

func (o LogAnalyticsConfigurationOutput) CustomerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsConfiguration) *string { return v.CustomerId }).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsConfigurationOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsConfiguration) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

type LogAnalyticsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (LogAnalyticsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsConfiguration)(nil)).Elem()
}

func (o LogAnalyticsConfigurationPtrOutput) ToLogAnalyticsConfigurationPtrOutput() LogAnalyticsConfigurationPtrOutput {
	return o
}

func (o LogAnalyticsConfigurationPtrOutput) ToLogAnalyticsConfigurationPtrOutputWithContext(ctx context.Context) LogAnalyticsConfigurationPtrOutput {
	return o
}

func (o LogAnalyticsConfigurationPtrOutput) Elem() LogAnalyticsConfigurationOutput {
	return o.ApplyT(func(v *LogAnalyticsConfiguration) LogAnalyticsConfiguration {
		if v != nil {
			return *v
		}
		var ret LogAnalyticsConfiguration
		return ret
	}).(LogAnalyticsConfigurationOutput)
}

func (o LogAnalyticsConfigurationPtrOutput) CustomerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.CustomerId
	}).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsConfigurationPtrOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.SharedKey
	}).(pulumi.StringPtrOutput)
}

type LogAnalyticsConfigurationResponse struct {
	CustomerId *string `pulumi:"customerId"`
}

type LogAnalyticsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (LogAnalyticsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsConfigurationResponse)(nil)).Elem()
}

func (o LogAnalyticsConfigurationResponseOutput) ToLogAnalyticsConfigurationResponseOutput() LogAnalyticsConfigurationResponseOutput {
	return o
}

func (o LogAnalyticsConfigurationResponseOutput) ToLogAnalyticsConfigurationResponseOutputWithContext(ctx context.Context) LogAnalyticsConfigurationResponseOutput {
	return o
}

func (o LogAnalyticsConfigurationResponseOutput) CustomerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsConfigurationResponse) *string { return v.CustomerId }).(pulumi.StringPtrOutput)
}

type LogAnalyticsConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (LogAnalyticsConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsConfigurationResponse)(nil)).Elem()
}

func (o LogAnalyticsConfigurationResponsePtrOutput) ToLogAnalyticsConfigurationResponsePtrOutput() LogAnalyticsConfigurationResponsePtrOutput {
	return o
}

func (o LogAnalyticsConfigurationResponsePtrOutput) ToLogAnalyticsConfigurationResponsePtrOutputWithContext(ctx context.Context) LogAnalyticsConfigurationResponsePtrOutput {
	return o
}

func (o LogAnalyticsConfigurationResponsePtrOutput) Elem() LogAnalyticsConfigurationResponseOutput {
	return o.ApplyT(func(v *LogAnalyticsConfigurationResponse) LogAnalyticsConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret LogAnalyticsConfigurationResponse
		return ret
	}).(LogAnalyticsConfigurationResponseOutput)
}

func (o LogAnalyticsConfigurationResponsePtrOutput) CustomerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomerId
	}).(pulumi.StringPtrOutput)
}

type Login struct {
	AllowedExternalRedirectUrls   []string          `pulumi:"allowedExternalRedirectUrls"`
	CookieExpiration              *CookieExpiration `pulumi:"cookieExpiration"`
	Nonce                         *Nonce            `pulumi:"nonce"`
	PreserveUrlFragmentsForLogins *bool             `pulumi:"preserveUrlFragmentsForLogins"`
	Routes                        *LoginRoutes      `pulumi:"routes"`
	TokenStore                    *TokenStore       `pulumi:"tokenStore"`
}





type LoginInput interface {
	pulumi.Input

	ToLoginOutput() LoginOutput
	ToLoginOutputWithContext(context.Context) LoginOutput
}

type LoginArgs struct {
	AllowedExternalRedirectUrls   pulumi.StringArrayInput  `pulumi:"allowedExternalRedirectUrls"`
	CookieExpiration              CookieExpirationPtrInput `pulumi:"cookieExpiration"`
	Nonce                         NoncePtrInput            `pulumi:"nonce"`
	PreserveUrlFragmentsForLogins pulumi.BoolPtrInput      `pulumi:"preserveUrlFragmentsForLogins"`
	Routes                        LoginRoutesPtrInput      `pulumi:"routes"`
	TokenStore                    TokenStorePtrInput       `pulumi:"tokenStore"`
}

func (LoginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Login)(nil)).Elem()
}

func (i LoginArgs) ToLoginOutput() LoginOutput {
	return i.ToLoginOutputWithContext(context.Background())
}

func (i LoginArgs) ToLoginOutputWithContext(ctx context.Context) LoginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginOutput)
}

func (i LoginArgs) ToLoginPtrOutput() LoginPtrOutput {
	return i.ToLoginPtrOutputWithContext(context.Background())
}

func (i LoginArgs) ToLoginPtrOutputWithContext(ctx context.Context) LoginPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginOutput).ToLoginPtrOutputWithContext(ctx)
}









type LoginPtrInput interface {
	pulumi.Input

	ToLoginPtrOutput() LoginPtrOutput
	ToLoginPtrOutputWithContext(context.Context) LoginPtrOutput
}

type loginPtrType LoginArgs

func LoginPtr(v *LoginArgs) LoginPtrInput {
	return (*loginPtrType)(v)
}

func (*loginPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Login)(nil)).Elem()
}

func (i *loginPtrType) ToLoginPtrOutput() LoginPtrOutput {
	return i.ToLoginPtrOutputWithContext(context.Background())
}

func (i *loginPtrType) ToLoginPtrOutputWithContext(ctx context.Context) LoginPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginPtrOutput)
}

type LoginOutput struct{ *pulumi.OutputState }

func (LoginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Login)(nil)).Elem()
}

func (o LoginOutput) ToLoginOutput() LoginOutput {
	return o
}

func (o LoginOutput) ToLoginOutputWithContext(ctx context.Context) LoginOutput {
	return o
}

func (o LoginOutput) ToLoginPtrOutput() LoginPtrOutput {
	return o.ToLoginPtrOutputWithContext(context.Background())
}

func (o LoginOutput) ToLoginPtrOutputWithContext(ctx context.Context) LoginPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Login) *Login {
		return &v
	}).(LoginPtrOutput)
}

func (o LoginOutput) AllowedExternalRedirectUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Login) []string { return v.AllowedExternalRedirectUrls }).(pulumi.StringArrayOutput)
}

func (o LoginOutput) CookieExpiration() CookieExpirationPtrOutput {
	return o.ApplyT(func(v Login) *CookieExpiration { return v.CookieExpiration }).(CookieExpirationPtrOutput)
}

func (o LoginOutput) Nonce() NoncePtrOutput {
	return o.ApplyT(func(v Login) *Nonce { return v.Nonce }).(NoncePtrOutput)
}

func (o LoginOutput) PreserveUrlFragmentsForLogins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Login) *bool { return v.PreserveUrlFragmentsForLogins }).(pulumi.BoolPtrOutput)
}

func (o LoginOutput) Routes() LoginRoutesPtrOutput {
	return o.ApplyT(func(v Login) *LoginRoutes { return v.Routes }).(LoginRoutesPtrOutput)
}

func (o LoginOutput) TokenStore() TokenStorePtrOutput {
	return o.ApplyT(func(v Login) *TokenStore { return v.TokenStore }).(TokenStorePtrOutput)
}

type LoginPtrOutput struct{ *pulumi.OutputState }

func (LoginPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Login)(nil)).Elem()
}

func (o LoginPtrOutput) ToLoginPtrOutput() LoginPtrOutput {
	return o
}

func (o LoginPtrOutput) ToLoginPtrOutputWithContext(ctx context.Context) LoginPtrOutput {
	return o
}

func (o LoginPtrOutput) Elem() LoginOutput {
	return o.ApplyT(func(v *Login) Login {
		if v != nil {
			return *v
		}
		var ret Login
		return ret
	}).(LoginOutput)
}

func (o LoginPtrOutput) AllowedExternalRedirectUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Login) []string {
		if v == nil {
			return nil
		}
		return v.AllowedExternalRedirectUrls
	}).(pulumi.StringArrayOutput)
}

func (o LoginPtrOutput) CookieExpiration() CookieExpirationPtrOutput {
	return o.ApplyT(func(v *Login) *CookieExpiration {
		if v == nil {
			return nil
		}
		return v.CookieExpiration
	}).(CookieExpirationPtrOutput)
}

func (o LoginPtrOutput) Nonce() NoncePtrOutput {
	return o.ApplyT(func(v *Login) *Nonce {
		if v == nil {
			return nil
		}
		return v.Nonce
	}).(NoncePtrOutput)
}

func (o LoginPtrOutput) PreserveUrlFragmentsForLogins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Login) *bool {
		if v == nil {
			return nil
		}
		return v.PreserveUrlFragmentsForLogins
	}).(pulumi.BoolPtrOutput)
}

func (o LoginPtrOutput) Routes() LoginRoutesPtrOutput {
	return o.ApplyT(func(v *Login) *LoginRoutes {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(LoginRoutesPtrOutput)
}

func (o LoginPtrOutput) TokenStore() TokenStorePtrOutput {
	return o.ApplyT(func(v *Login) *TokenStore {
		if v == nil {
			return nil
		}
		return v.TokenStore
	}).(TokenStorePtrOutput)
}

type LoginResponse struct {
	AllowedExternalRedirectUrls   []string                  `pulumi:"allowedExternalRedirectUrls"`
	CookieExpiration              *CookieExpirationResponse `pulumi:"cookieExpiration"`
	Nonce                         *NonceResponse            `pulumi:"nonce"`
	PreserveUrlFragmentsForLogins *bool                     `pulumi:"preserveUrlFragmentsForLogins"`
	Routes                        *LoginRoutesResponse      `pulumi:"routes"`
	TokenStore                    *TokenStoreResponse       `pulumi:"tokenStore"`
}

type LoginResponseOutput struct{ *pulumi.OutputState }

func (LoginResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoginResponse)(nil)).Elem()
}

func (o LoginResponseOutput) ToLoginResponseOutput() LoginResponseOutput {
	return o
}

func (o LoginResponseOutput) ToLoginResponseOutputWithContext(ctx context.Context) LoginResponseOutput {
	return o
}

func (o LoginResponseOutput) AllowedExternalRedirectUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LoginResponse) []string { return v.AllowedExternalRedirectUrls }).(pulumi.StringArrayOutput)
}

func (o LoginResponseOutput) CookieExpiration() CookieExpirationResponsePtrOutput {
	return o.ApplyT(func(v LoginResponse) *CookieExpirationResponse { return v.CookieExpiration }).(CookieExpirationResponsePtrOutput)
}

func (o LoginResponseOutput) Nonce() NonceResponsePtrOutput {
	return o.ApplyT(func(v LoginResponse) *NonceResponse { return v.Nonce }).(NonceResponsePtrOutput)
}

func (o LoginResponseOutput) PreserveUrlFragmentsForLogins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LoginResponse) *bool { return v.PreserveUrlFragmentsForLogins }).(pulumi.BoolPtrOutput)
}

func (o LoginResponseOutput) Routes() LoginRoutesResponsePtrOutput {
	return o.ApplyT(func(v LoginResponse) *LoginRoutesResponse { return v.Routes }).(LoginRoutesResponsePtrOutput)
}

func (o LoginResponseOutput) TokenStore() TokenStoreResponsePtrOutput {
	return o.ApplyT(func(v LoginResponse) *TokenStoreResponse { return v.TokenStore }).(TokenStoreResponsePtrOutput)
}

type LoginResponsePtrOutput struct{ *pulumi.OutputState }

func (LoginResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginResponse)(nil)).Elem()
}

func (o LoginResponsePtrOutput) ToLoginResponsePtrOutput() LoginResponsePtrOutput {
	return o
}

func (o LoginResponsePtrOutput) ToLoginResponsePtrOutputWithContext(ctx context.Context) LoginResponsePtrOutput {
	return o
}

func (o LoginResponsePtrOutput) Elem() LoginResponseOutput {
	return o.ApplyT(func(v *LoginResponse) LoginResponse {
		if v != nil {
			return *v
		}
		var ret LoginResponse
		return ret
	}).(LoginResponseOutput)
}

func (o LoginResponsePtrOutput) AllowedExternalRedirectUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LoginResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedExternalRedirectUrls
	}).(pulumi.StringArrayOutput)
}

func (o LoginResponsePtrOutput) CookieExpiration() CookieExpirationResponsePtrOutput {
	return o.ApplyT(func(v *LoginResponse) *CookieExpirationResponse {
		if v == nil {
			return nil
		}
		return v.CookieExpiration
	}).(CookieExpirationResponsePtrOutput)
}

func (o LoginResponsePtrOutput) Nonce() NonceResponsePtrOutput {
	return o.ApplyT(func(v *LoginResponse) *NonceResponse {
		if v == nil {
			return nil
		}
		return v.Nonce
	}).(NonceResponsePtrOutput)
}

func (o LoginResponsePtrOutput) PreserveUrlFragmentsForLogins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LoginResponse) *bool {
		if v == nil {
			return nil
		}
		return v.PreserveUrlFragmentsForLogins
	}).(pulumi.BoolPtrOutput)
}

func (o LoginResponsePtrOutput) Routes() LoginRoutesResponsePtrOutput {
	return o.ApplyT(func(v *LoginResponse) *LoginRoutesResponse {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(LoginRoutesResponsePtrOutput)
}

func (o LoginResponsePtrOutput) TokenStore() TokenStoreResponsePtrOutput {
	return o.ApplyT(func(v *LoginResponse) *TokenStoreResponse {
		if v == nil {
			return nil
		}
		return v.TokenStore
	}).(TokenStoreResponsePtrOutput)
}

type LoginRoutes struct {
	LogoutEndpoint *string `pulumi:"logoutEndpoint"`
}





type LoginRoutesInput interface {
	pulumi.Input

	ToLoginRoutesOutput() LoginRoutesOutput
	ToLoginRoutesOutputWithContext(context.Context) LoginRoutesOutput
}

type LoginRoutesArgs struct {
	LogoutEndpoint pulumi.StringPtrInput `pulumi:"logoutEndpoint"`
}

func (LoginRoutesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoginRoutes)(nil)).Elem()
}

func (i LoginRoutesArgs) ToLoginRoutesOutput() LoginRoutesOutput {
	return i.ToLoginRoutesOutputWithContext(context.Background())
}

func (i LoginRoutesArgs) ToLoginRoutesOutputWithContext(ctx context.Context) LoginRoutesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginRoutesOutput)
}

func (i LoginRoutesArgs) ToLoginRoutesPtrOutput() LoginRoutesPtrOutput {
	return i.ToLoginRoutesPtrOutputWithContext(context.Background())
}

func (i LoginRoutesArgs) ToLoginRoutesPtrOutputWithContext(ctx context.Context) LoginRoutesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginRoutesOutput).ToLoginRoutesPtrOutputWithContext(ctx)
}









type LoginRoutesPtrInput interface {
	pulumi.Input

	ToLoginRoutesPtrOutput() LoginRoutesPtrOutput
	ToLoginRoutesPtrOutputWithContext(context.Context) LoginRoutesPtrOutput
}

type loginRoutesPtrType LoginRoutesArgs

func LoginRoutesPtr(v *LoginRoutesArgs) LoginRoutesPtrInput {
	return (*loginRoutesPtrType)(v)
}

func (*loginRoutesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginRoutes)(nil)).Elem()
}

func (i *loginRoutesPtrType) ToLoginRoutesPtrOutput() LoginRoutesPtrOutput {
	return i.ToLoginRoutesPtrOutputWithContext(context.Background())
}

func (i *loginRoutesPtrType) ToLoginRoutesPtrOutputWithContext(ctx context.Context) LoginRoutesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginRoutesPtrOutput)
}

type LoginRoutesOutput struct{ *pulumi.OutputState }

func (LoginRoutesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoginRoutes)(nil)).Elem()
}

func (o LoginRoutesOutput) ToLoginRoutesOutput() LoginRoutesOutput {
	return o
}

func (o LoginRoutesOutput) ToLoginRoutesOutputWithContext(ctx context.Context) LoginRoutesOutput {
	return o
}

func (o LoginRoutesOutput) ToLoginRoutesPtrOutput() LoginRoutesPtrOutput {
	return o.ToLoginRoutesPtrOutputWithContext(context.Background())
}

func (o LoginRoutesOutput) ToLoginRoutesPtrOutputWithContext(ctx context.Context) LoginRoutesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoginRoutes) *LoginRoutes {
		return &v
	}).(LoginRoutesPtrOutput)
}

func (o LoginRoutesOutput) LogoutEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoginRoutes) *string { return v.LogoutEndpoint }).(pulumi.StringPtrOutput)
}

type LoginRoutesPtrOutput struct{ *pulumi.OutputState }

func (LoginRoutesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginRoutes)(nil)).Elem()
}

func (o LoginRoutesPtrOutput) ToLoginRoutesPtrOutput() LoginRoutesPtrOutput {
	return o
}

func (o LoginRoutesPtrOutput) ToLoginRoutesPtrOutputWithContext(ctx context.Context) LoginRoutesPtrOutput {
	return o
}

func (o LoginRoutesPtrOutput) Elem() LoginRoutesOutput {
	return o.ApplyT(func(v *LoginRoutes) LoginRoutes {
		if v != nil {
			return *v
		}
		var ret LoginRoutes
		return ret
	}).(LoginRoutesOutput)
}

func (o LoginRoutesPtrOutput) LogoutEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoginRoutes) *string {
		if v == nil {
			return nil
		}
		return v.LogoutEndpoint
	}).(pulumi.StringPtrOutput)
}

type LoginRoutesResponse struct {
	LogoutEndpoint *string `pulumi:"logoutEndpoint"`
}

type LoginRoutesResponseOutput struct{ *pulumi.OutputState }

func (LoginRoutesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoginRoutesResponse)(nil)).Elem()
}

func (o LoginRoutesResponseOutput) ToLoginRoutesResponseOutput() LoginRoutesResponseOutput {
	return o
}

func (o LoginRoutesResponseOutput) ToLoginRoutesResponseOutputWithContext(ctx context.Context) LoginRoutesResponseOutput {
	return o
}

func (o LoginRoutesResponseOutput) LogoutEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoginRoutesResponse) *string { return v.LogoutEndpoint }).(pulumi.StringPtrOutput)
}

type LoginRoutesResponsePtrOutput struct{ *pulumi.OutputState }

func (LoginRoutesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginRoutesResponse)(nil)).Elem()
}

func (o LoginRoutesResponsePtrOutput) ToLoginRoutesResponsePtrOutput() LoginRoutesResponsePtrOutput {
	return o
}

func (o LoginRoutesResponsePtrOutput) ToLoginRoutesResponsePtrOutputWithContext(ctx context.Context) LoginRoutesResponsePtrOutput {
	return o
}

func (o LoginRoutesResponsePtrOutput) Elem() LoginRoutesResponseOutput {
	return o.ApplyT(func(v *LoginRoutesResponse) LoginRoutesResponse {
		if v != nil {
			return *v
		}
		var ret LoginRoutesResponse
		return ret
	}).(LoginRoutesResponseOutput)
}

func (o LoginRoutesResponsePtrOutput) LogoutEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoginRoutesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LogoutEndpoint
	}).(pulumi.StringPtrOutput)
}

type LoginScopes struct {
	Scopes []string `pulumi:"scopes"`
}





type LoginScopesInput interface {
	pulumi.Input

	ToLoginScopesOutput() LoginScopesOutput
	ToLoginScopesOutputWithContext(context.Context) LoginScopesOutput
}

type LoginScopesArgs struct {
	Scopes pulumi.StringArrayInput `pulumi:"scopes"`
}

func (LoginScopesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoginScopes)(nil)).Elem()
}

func (i LoginScopesArgs) ToLoginScopesOutput() LoginScopesOutput {
	return i.ToLoginScopesOutputWithContext(context.Background())
}

func (i LoginScopesArgs) ToLoginScopesOutputWithContext(ctx context.Context) LoginScopesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginScopesOutput)
}

func (i LoginScopesArgs) ToLoginScopesPtrOutput() LoginScopesPtrOutput {
	return i.ToLoginScopesPtrOutputWithContext(context.Background())
}

func (i LoginScopesArgs) ToLoginScopesPtrOutputWithContext(ctx context.Context) LoginScopesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginScopesOutput).ToLoginScopesPtrOutputWithContext(ctx)
}









type LoginScopesPtrInput interface {
	pulumi.Input

	ToLoginScopesPtrOutput() LoginScopesPtrOutput
	ToLoginScopesPtrOutputWithContext(context.Context) LoginScopesPtrOutput
}

type loginScopesPtrType LoginScopesArgs

func LoginScopesPtr(v *LoginScopesArgs) LoginScopesPtrInput {
	return (*loginScopesPtrType)(v)
}

func (*loginScopesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginScopes)(nil)).Elem()
}

func (i *loginScopesPtrType) ToLoginScopesPtrOutput() LoginScopesPtrOutput {
	return i.ToLoginScopesPtrOutputWithContext(context.Background())
}

func (i *loginScopesPtrType) ToLoginScopesPtrOutputWithContext(ctx context.Context) LoginScopesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginScopesPtrOutput)
}

type LoginScopesOutput struct{ *pulumi.OutputState }

func (LoginScopesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoginScopes)(nil)).Elem()
}

func (o LoginScopesOutput) ToLoginScopesOutput() LoginScopesOutput {
	return o
}

func (o LoginScopesOutput) ToLoginScopesOutputWithContext(ctx context.Context) LoginScopesOutput {
	return o
}

func (o LoginScopesOutput) ToLoginScopesPtrOutput() LoginScopesPtrOutput {
	return o.ToLoginScopesPtrOutputWithContext(context.Background())
}

func (o LoginScopesOutput) ToLoginScopesPtrOutputWithContext(ctx context.Context) LoginScopesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoginScopes) *LoginScopes {
		return &v
	}).(LoginScopesPtrOutput)
}

func (o LoginScopesOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LoginScopes) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type LoginScopesPtrOutput struct{ *pulumi.OutputState }

func (LoginScopesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginScopes)(nil)).Elem()
}

func (o LoginScopesPtrOutput) ToLoginScopesPtrOutput() LoginScopesPtrOutput {
	return o
}

func (o LoginScopesPtrOutput) ToLoginScopesPtrOutputWithContext(ctx context.Context) LoginScopesPtrOutput {
	return o
}

func (o LoginScopesPtrOutput) Elem() LoginScopesOutput {
	return o.ApplyT(func(v *LoginScopes) LoginScopes {
		if v != nil {
			return *v
		}
		var ret LoginScopes
		return ret
	}).(LoginScopesOutput)
}

func (o LoginScopesPtrOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LoginScopes) []string {
		if v == nil {
			return nil
		}
		return v.Scopes
	}).(pulumi.StringArrayOutput)
}

type LoginScopesResponse struct {
	Scopes []string `pulumi:"scopes"`
}

type LoginScopesResponseOutput struct{ *pulumi.OutputState }

func (LoginScopesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoginScopesResponse)(nil)).Elem()
}

func (o LoginScopesResponseOutput) ToLoginScopesResponseOutput() LoginScopesResponseOutput {
	return o
}

func (o LoginScopesResponseOutput) ToLoginScopesResponseOutputWithContext(ctx context.Context) LoginScopesResponseOutput {
	return o
}

func (o LoginScopesResponseOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LoginScopesResponse) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type LoginScopesResponsePtrOutput struct{ *pulumi.OutputState }

func (LoginScopesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginScopesResponse)(nil)).Elem()
}

func (o LoginScopesResponsePtrOutput) ToLoginScopesResponsePtrOutput() LoginScopesResponsePtrOutput {
	return o
}

func (o LoginScopesResponsePtrOutput) ToLoginScopesResponsePtrOutputWithContext(ctx context.Context) LoginScopesResponsePtrOutput {
	return o
}

func (o LoginScopesResponsePtrOutput) Elem() LoginScopesResponseOutput {
	return o.ApplyT(func(v *LoginScopesResponse) LoginScopesResponse {
		if v != nil {
			return *v
		}
		var ret LoginScopesResponse
		return ret
	}).(LoginScopesResponseOutput)
}

func (o LoginScopesResponsePtrOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LoginScopesResponse) []string {
		if v == nil {
			return nil
		}
		return v.Scopes
	}).(pulumi.StringArrayOutput)
}

type ManagedServiceIdentity struct {
	Type                   *ManagedServiceIdentityType `pulumi:"type"`
	UserAssignedIdentities map[string]interface{}      `pulumi:"userAssignedIdentities"`
}





type ManagedServiceIdentityInput interface {
	pulumi.Input

	ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput
	ToManagedServiceIdentityOutputWithContext(context.Context) ManagedServiceIdentityOutput
}

type ManagedServiceIdentityArgs struct {
	Type                   ManagedServiceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput                    `pulumi:"userAssignedIdentities"`
}

func (ManagedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return i.ToManagedServiceIdentityOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput)
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput).ToManagedServiceIdentityPtrOutputWithContext(ctx)
}









type ManagedServiceIdentityPtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput
	ToManagedServiceIdentityPtrOutputWithContext(context.Context) ManagedServiceIdentityPtrOutput
}

type managedServiceIdentityPtrType ManagedServiceIdentityArgs

func ManagedServiceIdentityPtr(v *ManagedServiceIdentityArgs) ManagedServiceIdentityPtrInput {
	return (*managedServiceIdentityPtrType)(v)
}

func (*managedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityPtrOutput)
}

type ManagedServiceIdentityOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentity) *ManagedServiceIdentity {
		return &v
	}).(ManagedServiceIdentityPtrOutput)
}

func (o ManagedServiceIdentityOutput) Type() ManagedServiceIdentityTypePtrOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) *ManagedServiceIdentityType { return v.Type }).(ManagedServiceIdentityTypePtrOutput)
}

func (o ManagedServiceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) Elem() ManagedServiceIdentityOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) ManagedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentity
		return ret
	}).(ManagedServiceIdentityOutput)
}

func (o ManagedServiceIdentityPtrOutput) Type() ManagedServiceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) *ManagedServiceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ManagedServiceIdentityTypePtrOutput)
}

func (o ManagedServiceIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ManagedServiceIdentityResponse struct {
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   *string                                 `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}

type ManagedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ManagedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) Elem() ManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) ManagedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityResponse
		return ret
	}).(ManagedServiceIdentityResponseOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type NameValuePair struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type NameValuePairInput interface {
	pulumi.Input

	ToNameValuePairOutput() NameValuePairOutput
	ToNameValuePairOutputWithContext(context.Context) NameValuePairOutput
}

type NameValuePairArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (NameValuePairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NameValuePair)(nil)).Elem()
}

func (i NameValuePairArgs) ToNameValuePairOutput() NameValuePairOutput {
	return i.ToNameValuePairOutputWithContext(context.Background())
}

func (i NameValuePairArgs) ToNameValuePairOutputWithContext(ctx context.Context) NameValuePairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NameValuePairOutput)
}





type NameValuePairArrayInput interface {
	pulumi.Input

	ToNameValuePairArrayOutput() NameValuePairArrayOutput
	ToNameValuePairArrayOutputWithContext(context.Context) NameValuePairArrayOutput
}

type NameValuePairArray []NameValuePairInput

func (NameValuePairArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NameValuePair)(nil)).Elem()
}

func (i NameValuePairArray) ToNameValuePairArrayOutput() NameValuePairArrayOutput {
	return i.ToNameValuePairArrayOutputWithContext(context.Background())
}

func (i NameValuePairArray) ToNameValuePairArrayOutputWithContext(ctx context.Context) NameValuePairArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NameValuePairArrayOutput)
}

type NameValuePairOutput struct{ *pulumi.OutputState }

func (NameValuePairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NameValuePair)(nil)).Elem()
}

func (o NameValuePairOutput) ToNameValuePairOutput() NameValuePairOutput {
	return o
}

func (o NameValuePairOutput) ToNameValuePairOutputWithContext(ctx context.Context) NameValuePairOutput {
	return o
}

func (o NameValuePairOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NameValuePair) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NameValuePairOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NameValuePair) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type NameValuePairArrayOutput struct{ *pulumi.OutputState }

func (NameValuePairArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NameValuePair)(nil)).Elem()
}

func (o NameValuePairArrayOutput) ToNameValuePairArrayOutput() NameValuePairArrayOutput {
	return o
}

func (o NameValuePairArrayOutput) ToNameValuePairArrayOutputWithContext(ctx context.Context) NameValuePairArrayOutput {
	return o
}

func (o NameValuePairArrayOutput) Index(i pulumi.IntInput) NameValuePairOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NameValuePair {
		return vs[0].([]NameValuePair)[vs[1].(int)]
	}).(NameValuePairOutput)
}

type NameValuePairResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type NameValuePairResponseOutput struct{ *pulumi.OutputState }

func (NameValuePairResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NameValuePairResponse)(nil)).Elem()
}

func (o NameValuePairResponseOutput) ToNameValuePairResponseOutput() NameValuePairResponseOutput {
	return o
}

func (o NameValuePairResponseOutput) ToNameValuePairResponseOutputWithContext(ctx context.Context) NameValuePairResponseOutput {
	return o
}

func (o NameValuePairResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NameValuePairResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NameValuePairResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NameValuePairResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type NameValuePairResponseArrayOutput struct{ *pulumi.OutputState }

func (NameValuePairResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NameValuePairResponse)(nil)).Elem()
}

func (o NameValuePairResponseArrayOutput) ToNameValuePairResponseArrayOutput() NameValuePairResponseArrayOutput {
	return o
}

func (o NameValuePairResponseArrayOutput) ToNameValuePairResponseArrayOutputWithContext(ctx context.Context) NameValuePairResponseArrayOutput {
	return o
}

func (o NameValuePairResponseArrayOutput) Index(i pulumi.IntInput) NameValuePairResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NameValuePairResponse {
		return vs[0].([]NameValuePairResponse)[vs[1].(int)]
	}).(NameValuePairResponseOutput)
}

type Nonce struct {
	NonceExpirationInterval *string `pulumi:"nonceExpirationInterval"`
	ValidateNonce           *bool   `pulumi:"validateNonce"`
}





type NonceInput interface {
	pulumi.Input

	ToNonceOutput() NonceOutput
	ToNonceOutputWithContext(context.Context) NonceOutput
}

type NonceArgs struct {
	NonceExpirationInterval pulumi.StringPtrInput `pulumi:"nonceExpirationInterval"`
	ValidateNonce           pulumi.BoolPtrInput   `pulumi:"validateNonce"`
}

func (NonceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Nonce)(nil)).Elem()
}

func (i NonceArgs) ToNonceOutput() NonceOutput {
	return i.ToNonceOutputWithContext(context.Background())
}

func (i NonceArgs) ToNonceOutputWithContext(ctx context.Context) NonceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NonceOutput)
}

func (i NonceArgs) ToNoncePtrOutput() NoncePtrOutput {
	return i.ToNoncePtrOutputWithContext(context.Background())
}

func (i NonceArgs) ToNoncePtrOutputWithContext(ctx context.Context) NoncePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NonceOutput).ToNoncePtrOutputWithContext(ctx)
}









type NoncePtrInput interface {
	pulumi.Input

	ToNoncePtrOutput() NoncePtrOutput
	ToNoncePtrOutputWithContext(context.Context) NoncePtrOutput
}

type noncePtrType NonceArgs

func NoncePtr(v *NonceArgs) NoncePtrInput {
	return (*noncePtrType)(v)
}

func (*noncePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Nonce)(nil)).Elem()
}

func (i *noncePtrType) ToNoncePtrOutput() NoncePtrOutput {
	return i.ToNoncePtrOutputWithContext(context.Background())
}

func (i *noncePtrType) ToNoncePtrOutputWithContext(ctx context.Context) NoncePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoncePtrOutput)
}

type NonceOutput struct{ *pulumi.OutputState }

func (NonceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Nonce)(nil)).Elem()
}

func (o NonceOutput) ToNonceOutput() NonceOutput {
	return o
}

func (o NonceOutput) ToNonceOutputWithContext(ctx context.Context) NonceOutput {
	return o
}

func (o NonceOutput) ToNoncePtrOutput() NoncePtrOutput {
	return o.ToNoncePtrOutputWithContext(context.Background())
}

func (o NonceOutput) ToNoncePtrOutputWithContext(ctx context.Context) NoncePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Nonce) *Nonce {
		return &v
	}).(NoncePtrOutput)
}

func (o NonceOutput) NonceExpirationInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Nonce) *string { return v.NonceExpirationInterval }).(pulumi.StringPtrOutput)
}

func (o NonceOutput) ValidateNonce() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Nonce) *bool { return v.ValidateNonce }).(pulumi.BoolPtrOutput)
}

type NoncePtrOutput struct{ *pulumi.OutputState }

func (NoncePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Nonce)(nil)).Elem()
}

func (o NoncePtrOutput) ToNoncePtrOutput() NoncePtrOutput {
	return o
}

func (o NoncePtrOutput) ToNoncePtrOutputWithContext(ctx context.Context) NoncePtrOutput {
	return o
}

func (o NoncePtrOutput) Elem() NonceOutput {
	return o.ApplyT(func(v *Nonce) Nonce {
		if v != nil {
			return *v
		}
		var ret Nonce
		return ret
	}).(NonceOutput)
}

func (o NoncePtrOutput) NonceExpirationInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Nonce) *string {
		if v == nil {
			return nil
		}
		return v.NonceExpirationInterval
	}).(pulumi.StringPtrOutput)
}

func (o NoncePtrOutput) ValidateNonce() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Nonce) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateNonce
	}).(pulumi.BoolPtrOutput)
}

type NonceResponse struct {
	NonceExpirationInterval *string `pulumi:"nonceExpirationInterval"`
	ValidateNonce           *bool   `pulumi:"validateNonce"`
}

type NonceResponseOutput struct{ *pulumi.OutputState }

func (NonceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NonceResponse)(nil)).Elem()
}

func (o NonceResponseOutput) ToNonceResponseOutput() NonceResponseOutput {
	return o
}

func (o NonceResponseOutput) ToNonceResponseOutputWithContext(ctx context.Context) NonceResponseOutput {
	return o
}

func (o NonceResponseOutput) NonceExpirationInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NonceResponse) *string { return v.NonceExpirationInterval }).(pulumi.StringPtrOutput)
}

func (o NonceResponseOutput) ValidateNonce() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NonceResponse) *bool { return v.ValidateNonce }).(pulumi.BoolPtrOutput)
}

type NonceResponsePtrOutput struct{ *pulumi.OutputState }

func (NonceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NonceResponse)(nil)).Elem()
}

func (o NonceResponsePtrOutput) ToNonceResponsePtrOutput() NonceResponsePtrOutput {
	return o
}

func (o NonceResponsePtrOutput) ToNonceResponsePtrOutputWithContext(ctx context.Context) NonceResponsePtrOutput {
	return o
}

func (o NonceResponsePtrOutput) Elem() NonceResponseOutput {
	return o.ApplyT(func(v *NonceResponse) NonceResponse {
		if v != nil {
			return *v
		}
		var ret NonceResponse
		return ret
	}).(NonceResponseOutput)
}

func (o NonceResponsePtrOutput) NonceExpirationInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NonceResponse) *string {
		if v == nil {
			return nil
		}
		return v.NonceExpirationInterval
	}).(pulumi.StringPtrOutput)
}

func (o NonceResponsePtrOutput) ValidateNonce() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NonceResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ValidateNonce
	}).(pulumi.BoolPtrOutput)
}

type OpenIdConnectClientCredential struct {
	ClientSecretSettingName *string                 `pulumi:"clientSecretSettingName"`
	Method                  *ClientCredentialMethod `pulumi:"method"`
}





type OpenIdConnectClientCredentialInput interface {
	pulumi.Input

	ToOpenIdConnectClientCredentialOutput() OpenIdConnectClientCredentialOutput
	ToOpenIdConnectClientCredentialOutputWithContext(context.Context) OpenIdConnectClientCredentialOutput
}

type OpenIdConnectClientCredentialArgs struct {
	ClientSecretSettingName pulumi.StringPtrInput          `pulumi:"clientSecretSettingName"`
	Method                  ClientCredentialMethodPtrInput `pulumi:"method"`
}

func (OpenIdConnectClientCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectClientCredential)(nil)).Elem()
}

func (i OpenIdConnectClientCredentialArgs) ToOpenIdConnectClientCredentialOutput() OpenIdConnectClientCredentialOutput {
	return i.ToOpenIdConnectClientCredentialOutputWithContext(context.Background())
}

func (i OpenIdConnectClientCredentialArgs) ToOpenIdConnectClientCredentialOutputWithContext(ctx context.Context) OpenIdConnectClientCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectClientCredentialOutput)
}

func (i OpenIdConnectClientCredentialArgs) ToOpenIdConnectClientCredentialPtrOutput() OpenIdConnectClientCredentialPtrOutput {
	return i.ToOpenIdConnectClientCredentialPtrOutputWithContext(context.Background())
}

func (i OpenIdConnectClientCredentialArgs) ToOpenIdConnectClientCredentialPtrOutputWithContext(ctx context.Context) OpenIdConnectClientCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectClientCredentialOutput).ToOpenIdConnectClientCredentialPtrOutputWithContext(ctx)
}









type OpenIdConnectClientCredentialPtrInput interface {
	pulumi.Input

	ToOpenIdConnectClientCredentialPtrOutput() OpenIdConnectClientCredentialPtrOutput
	ToOpenIdConnectClientCredentialPtrOutputWithContext(context.Context) OpenIdConnectClientCredentialPtrOutput
}

type openIdConnectClientCredentialPtrType OpenIdConnectClientCredentialArgs

func OpenIdConnectClientCredentialPtr(v *OpenIdConnectClientCredentialArgs) OpenIdConnectClientCredentialPtrInput {
	return (*openIdConnectClientCredentialPtrType)(v)
}

func (*openIdConnectClientCredentialPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectClientCredential)(nil)).Elem()
}

func (i *openIdConnectClientCredentialPtrType) ToOpenIdConnectClientCredentialPtrOutput() OpenIdConnectClientCredentialPtrOutput {
	return i.ToOpenIdConnectClientCredentialPtrOutputWithContext(context.Background())
}

func (i *openIdConnectClientCredentialPtrType) ToOpenIdConnectClientCredentialPtrOutputWithContext(ctx context.Context) OpenIdConnectClientCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectClientCredentialPtrOutput)
}

type OpenIdConnectClientCredentialOutput struct{ *pulumi.OutputState }

func (OpenIdConnectClientCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectClientCredential)(nil)).Elem()
}

func (o OpenIdConnectClientCredentialOutput) ToOpenIdConnectClientCredentialOutput() OpenIdConnectClientCredentialOutput {
	return o
}

func (o OpenIdConnectClientCredentialOutput) ToOpenIdConnectClientCredentialOutputWithContext(ctx context.Context) OpenIdConnectClientCredentialOutput {
	return o
}

func (o OpenIdConnectClientCredentialOutput) ToOpenIdConnectClientCredentialPtrOutput() OpenIdConnectClientCredentialPtrOutput {
	return o.ToOpenIdConnectClientCredentialPtrOutputWithContext(context.Background())
}

func (o OpenIdConnectClientCredentialOutput) ToOpenIdConnectClientCredentialPtrOutputWithContext(ctx context.Context) OpenIdConnectClientCredentialPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OpenIdConnectClientCredential) *OpenIdConnectClientCredential {
		return &v
	}).(OpenIdConnectClientCredentialPtrOutput)
}

func (o OpenIdConnectClientCredentialOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectClientCredential) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectClientCredentialOutput) Method() ClientCredentialMethodPtrOutput {
	return o.ApplyT(func(v OpenIdConnectClientCredential) *ClientCredentialMethod { return v.Method }).(ClientCredentialMethodPtrOutput)
}

type OpenIdConnectClientCredentialPtrOutput struct{ *pulumi.OutputState }

func (OpenIdConnectClientCredentialPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectClientCredential)(nil)).Elem()
}

func (o OpenIdConnectClientCredentialPtrOutput) ToOpenIdConnectClientCredentialPtrOutput() OpenIdConnectClientCredentialPtrOutput {
	return o
}

func (o OpenIdConnectClientCredentialPtrOutput) ToOpenIdConnectClientCredentialPtrOutputWithContext(ctx context.Context) OpenIdConnectClientCredentialPtrOutput {
	return o
}

func (o OpenIdConnectClientCredentialPtrOutput) Elem() OpenIdConnectClientCredentialOutput {
	return o.ApplyT(func(v *OpenIdConnectClientCredential) OpenIdConnectClientCredential {
		if v != nil {
			return *v
		}
		var ret OpenIdConnectClientCredential
		return ret
	}).(OpenIdConnectClientCredentialOutput)
}

func (o OpenIdConnectClientCredentialPtrOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectClientCredential) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretSettingName
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectClientCredentialPtrOutput) Method() ClientCredentialMethodPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectClientCredential) *ClientCredentialMethod {
		if v == nil {
			return nil
		}
		return v.Method
	}).(ClientCredentialMethodPtrOutput)
}

type OpenIdConnectClientCredentialResponse struct {
	ClientSecretSettingName *string `pulumi:"clientSecretSettingName"`
	Method                  *string `pulumi:"method"`
}

type OpenIdConnectClientCredentialResponseOutput struct{ *pulumi.OutputState }

func (OpenIdConnectClientCredentialResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectClientCredentialResponse)(nil)).Elem()
}

func (o OpenIdConnectClientCredentialResponseOutput) ToOpenIdConnectClientCredentialResponseOutput() OpenIdConnectClientCredentialResponseOutput {
	return o
}

func (o OpenIdConnectClientCredentialResponseOutput) ToOpenIdConnectClientCredentialResponseOutputWithContext(ctx context.Context) OpenIdConnectClientCredentialResponseOutput {
	return o
}

func (o OpenIdConnectClientCredentialResponseOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectClientCredentialResponse) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectClientCredentialResponseOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectClientCredentialResponse) *string { return v.Method }).(pulumi.StringPtrOutput)
}

type OpenIdConnectClientCredentialResponsePtrOutput struct{ *pulumi.OutputState }

func (OpenIdConnectClientCredentialResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectClientCredentialResponse)(nil)).Elem()
}

func (o OpenIdConnectClientCredentialResponsePtrOutput) ToOpenIdConnectClientCredentialResponsePtrOutput() OpenIdConnectClientCredentialResponsePtrOutput {
	return o
}

func (o OpenIdConnectClientCredentialResponsePtrOutput) ToOpenIdConnectClientCredentialResponsePtrOutputWithContext(ctx context.Context) OpenIdConnectClientCredentialResponsePtrOutput {
	return o
}

func (o OpenIdConnectClientCredentialResponsePtrOutput) Elem() OpenIdConnectClientCredentialResponseOutput {
	return o.ApplyT(func(v *OpenIdConnectClientCredentialResponse) OpenIdConnectClientCredentialResponse {
		if v != nil {
			return *v
		}
		var ret OpenIdConnectClientCredentialResponse
		return ret
	}).(OpenIdConnectClientCredentialResponseOutput)
}

func (o OpenIdConnectClientCredentialResponsePtrOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectClientCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecretSettingName
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectClientCredentialResponsePtrOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectClientCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.Method
	}).(pulumi.StringPtrOutput)
}

type OpenIdConnectConfig struct {
	AuthorizationEndpoint        *string `pulumi:"authorizationEndpoint"`
	CertificationUri             *string `pulumi:"certificationUri"`
	Issuer                       *string `pulumi:"issuer"`
	TokenEndpoint                *string `pulumi:"tokenEndpoint"`
	WellKnownOpenIdConfiguration *string `pulumi:"wellKnownOpenIdConfiguration"`
}





type OpenIdConnectConfigInput interface {
	pulumi.Input

	ToOpenIdConnectConfigOutput() OpenIdConnectConfigOutput
	ToOpenIdConnectConfigOutputWithContext(context.Context) OpenIdConnectConfigOutput
}

type OpenIdConnectConfigArgs struct {
	AuthorizationEndpoint        pulumi.StringPtrInput `pulumi:"authorizationEndpoint"`
	CertificationUri             pulumi.StringPtrInput `pulumi:"certificationUri"`
	Issuer                       pulumi.StringPtrInput `pulumi:"issuer"`
	TokenEndpoint                pulumi.StringPtrInput `pulumi:"tokenEndpoint"`
	WellKnownOpenIdConfiguration pulumi.StringPtrInput `pulumi:"wellKnownOpenIdConfiguration"`
}

func (OpenIdConnectConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectConfig)(nil)).Elem()
}

func (i OpenIdConnectConfigArgs) ToOpenIdConnectConfigOutput() OpenIdConnectConfigOutput {
	return i.ToOpenIdConnectConfigOutputWithContext(context.Background())
}

func (i OpenIdConnectConfigArgs) ToOpenIdConnectConfigOutputWithContext(ctx context.Context) OpenIdConnectConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectConfigOutput)
}

func (i OpenIdConnectConfigArgs) ToOpenIdConnectConfigPtrOutput() OpenIdConnectConfigPtrOutput {
	return i.ToOpenIdConnectConfigPtrOutputWithContext(context.Background())
}

func (i OpenIdConnectConfigArgs) ToOpenIdConnectConfigPtrOutputWithContext(ctx context.Context) OpenIdConnectConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectConfigOutput).ToOpenIdConnectConfigPtrOutputWithContext(ctx)
}









type OpenIdConnectConfigPtrInput interface {
	pulumi.Input

	ToOpenIdConnectConfigPtrOutput() OpenIdConnectConfigPtrOutput
	ToOpenIdConnectConfigPtrOutputWithContext(context.Context) OpenIdConnectConfigPtrOutput
}

type openIdConnectConfigPtrType OpenIdConnectConfigArgs

func OpenIdConnectConfigPtr(v *OpenIdConnectConfigArgs) OpenIdConnectConfigPtrInput {
	return (*openIdConnectConfigPtrType)(v)
}

func (*openIdConnectConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectConfig)(nil)).Elem()
}

func (i *openIdConnectConfigPtrType) ToOpenIdConnectConfigPtrOutput() OpenIdConnectConfigPtrOutput {
	return i.ToOpenIdConnectConfigPtrOutputWithContext(context.Background())
}

func (i *openIdConnectConfigPtrType) ToOpenIdConnectConfigPtrOutputWithContext(ctx context.Context) OpenIdConnectConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectConfigPtrOutput)
}

type OpenIdConnectConfigOutput struct{ *pulumi.OutputState }

func (OpenIdConnectConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectConfig)(nil)).Elem()
}

func (o OpenIdConnectConfigOutput) ToOpenIdConnectConfigOutput() OpenIdConnectConfigOutput {
	return o
}

func (o OpenIdConnectConfigOutput) ToOpenIdConnectConfigOutputWithContext(ctx context.Context) OpenIdConnectConfigOutput {
	return o
}

func (o OpenIdConnectConfigOutput) ToOpenIdConnectConfigPtrOutput() OpenIdConnectConfigPtrOutput {
	return o.ToOpenIdConnectConfigPtrOutputWithContext(context.Background())
}

func (o OpenIdConnectConfigOutput) ToOpenIdConnectConfigPtrOutputWithContext(ctx context.Context) OpenIdConnectConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OpenIdConnectConfig) *OpenIdConnectConfig {
		return &v
	}).(OpenIdConnectConfigPtrOutput)
}

func (o OpenIdConnectConfigOutput) AuthorizationEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfig) *string { return v.AuthorizationEndpoint }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigOutput) CertificationUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfig) *string { return v.CertificationUri }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfig) *string { return v.Issuer }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigOutput) TokenEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfig) *string { return v.TokenEndpoint }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigOutput) WellKnownOpenIdConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfig) *string { return v.WellKnownOpenIdConfiguration }).(pulumi.StringPtrOutput)
}

type OpenIdConnectConfigPtrOutput struct{ *pulumi.OutputState }

func (OpenIdConnectConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectConfig)(nil)).Elem()
}

func (o OpenIdConnectConfigPtrOutput) ToOpenIdConnectConfigPtrOutput() OpenIdConnectConfigPtrOutput {
	return o
}

func (o OpenIdConnectConfigPtrOutput) ToOpenIdConnectConfigPtrOutputWithContext(ctx context.Context) OpenIdConnectConfigPtrOutput {
	return o
}

func (o OpenIdConnectConfigPtrOutput) Elem() OpenIdConnectConfigOutput {
	return o.ApplyT(func(v *OpenIdConnectConfig) OpenIdConnectConfig {
		if v != nil {
			return *v
		}
		var ret OpenIdConnectConfig
		return ret
	}).(OpenIdConnectConfigOutput)
}

func (o OpenIdConnectConfigPtrOutput) AuthorizationEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfig) *string {
		if v == nil {
			return nil
		}
		return v.AuthorizationEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigPtrOutput) CertificationUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfig) *string {
		if v == nil {
			return nil
		}
		return v.CertificationUri
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigPtrOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfig) *string {
		if v == nil {
			return nil
		}
		return v.Issuer
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigPtrOutput) TokenEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfig) *string {
		if v == nil {
			return nil
		}
		return v.TokenEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigPtrOutput) WellKnownOpenIdConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfig) *string {
		if v == nil {
			return nil
		}
		return v.WellKnownOpenIdConfiguration
	}).(pulumi.StringPtrOutput)
}

type OpenIdConnectConfigResponse struct {
	AuthorizationEndpoint        *string `pulumi:"authorizationEndpoint"`
	CertificationUri             *string `pulumi:"certificationUri"`
	Issuer                       *string `pulumi:"issuer"`
	TokenEndpoint                *string `pulumi:"tokenEndpoint"`
	WellKnownOpenIdConfiguration *string `pulumi:"wellKnownOpenIdConfiguration"`
}

type OpenIdConnectConfigResponseOutput struct{ *pulumi.OutputState }

func (OpenIdConnectConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectConfigResponse)(nil)).Elem()
}

func (o OpenIdConnectConfigResponseOutput) ToOpenIdConnectConfigResponseOutput() OpenIdConnectConfigResponseOutput {
	return o
}

func (o OpenIdConnectConfigResponseOutput) ToOpenIdConnectConfigResponseOutputWithContext(ctx context.Context) OpenIdConnectConfigResponseOutput {
	return o
}

func (o OpenIdConnectConfigResponseOutput) AuthorizationEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfigResponse) *string { return v.AuthorizationEndpoint }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigResponseOutput) CertificationUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfigResponse) *string { return v.CertificationUri }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigResponseOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfigResponse) *string { return v.Issuer }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigResponseOutput) TokenEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfigResponse) *string { return v.TokenEndpoint }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigResponseOutput) WellKnownOpenIdConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectConfigResponse) *string { return v.WellKnownOpenIdConfiguration }).(pulumi.StringPtrOutput)
}

type OpenIdConnectConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (OpenIdConnectConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectConfigResponse)(nil)).Elem()
}

func (o OpenIdConnectConfigResponsePtrOutput) ToOpenIdConnectConfigResponsePtrOutput() OpenIdConnectConfigResponsePtrOutput {
	return o
}

func (o OpenIdConnectConfigResponsePtrOutput) ToOpenIdConnectConfigResponsePtrOutputWithContext(ctx context.Context) OpenIdConnectConfigResponsePtrOutput {
	return o
}

func (o OpenIdConnectConfigResponsePtrOutput) Elem() OpenIdConnectConfigResponseOutput {
	return o.ApplyT(func(v *OpenIdConnectConfigResponse) OpenIdConnectConfigResponse {
		if v != nil {
			return *v
		}
		var ret OpenIdConnectConfigResponse
		return ret
	}).(OpenIdConnectConfigResponseOutput)
}

func (o OpenIdConnectConfigResponsePtrOutput) AuthorizationEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuthorizationEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigResponsePtrOutput) CertificationUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.CertificationUri
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigResponsePtrOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Issuer
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigResponsePtrOutput) TokenEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.TokenEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectConfigResponsePtrOutput) WellKnownOpenIdConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.WellKnownOpenIdConfiguration
	}).(pulumi.StringPtrOutput)
}

type OpenIdConnectLogin struct {
	NameClaimType *string  `pulumi:"nameClaimType"`
	Scopes        []string `pulumi:"scopes"`
}





type OpenIdConnectLoginInput interface {
	pulumi.Input

	ToOpenIdConnectLoginOutput() OpenIdConnectLoginOutput
	ToOpenIdConnectLoginOutputWithContext(context.Context) OpenIdConnectLoginOutput
}

type OpenIdConnectLoginArgs struct {
	NameClaimType pulumi.StringPtrInput   `pulumi:"nameClaimType"`
	Scopes        pulumi.StringArrayInput `pulumi:"scopes"`
}

func (OpenIdConnectLoginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectLogin)(nil)).Elem()
}

func (i OpenIdConnectLoginArgs) ToOpenIdConnectLoginOutput() OpenIdConnectLoginOutput {
	return i.ToOpenIdConnectLoginOutputWithContext(context.Background())
}

func (i OpenIdConnectLoginArgs) ToOpenIdConnectLoginOutputWithContext(ctx context.Context) OpenIdConnectLoginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectLoginOutput)
}

func (i OpenIdConnectLoginArgs) ToOpenIdConnectLoginPtrOutput() OpenIdConnectLoginPtrOutput {
	return i.ToOpenIdConnectLoginPtrOutputWithContext(context.Background())
}

func (i OpenIdConnectLoginArgs) ToOpenIdConnectLoginPtrOutputWithContext(ctx context.Context) OpenIdConnectLoginPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectLoginOutput).ToOpenIdConnectLoginPtrOutputWithContext(ctx)
}









type OpenIdConnectLoginPtrInput interface {
	pulumi.Input

	ToOpenIdConnectLoginPtrOutput() OpenIdConnectLoginPtrOutput
	ToOpenIdConnectLoginPtrOutputWithContext(context.Context) OpenIdConnectLoginPtrOutput
}

type openIdConnectLoginPtrType OpenIdConnectLoginArgs

func OpenIdConnectLoginPtr(v *OpenIdConnectLoginArgs) OpenIdConnectLoginPtrInput {
	return (*openIdConnectLoginPtrType)(v)
}

func (*openIdConnectLoginPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectLogin)(nil)).Elem()
}

func (i *openIdConnectLoginPtrType) ToOpenIdConnectLoginPtrOutput() OpenIdConnectLoginPtrOutput {
	return i.ToOpenIdConnectLoginPtrOutputWithContext(context.Background())
}

func (i *openIdConnectLoginPtrType) ToOpenIdConnectLoginPtrOutputWithContext(ctx context.Context) OpenIdConnectLoginPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectLoginPtrOutput)
}

type OpenIdConnectLoginOutput struct{ *pulumi.OutputState }

func (OpenIdConnectLoginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectLogin)(nil)).Elem()
}

func (o OpenIdConnectLoginOutput) ToOpenIdConnectLoginOutput() OpenIdConnectLoginOutput {
	return o
}

func (o OpenIdConnectLoginOutput) ToOpenIdConnectLoginOutputWithContext(ctx context.Context) OpenIdConnectLoginOutput {
	return o
}

func (o OpenIdConnectLoginOutput) ToOpenIdConnectLoginPtrOutput() OpenIdConnectLoginPtrOutput {
	return o.ToOpenIdConnectLoginPtrOutputWithContext(context.Background())
}

func (o OpenIdConnectLoginOutput) ToOpenIdConnectLoginPtrOutputWithContext(ctx context.Context) OpenIdConnectLoginPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OpenIdConnectLogin) *OpenIdConnectLogin {
		return &v
	}).(OpenIdConnectLoginPtrOutput)
}

func (o OpenIdConnectLoginOutput) NameClaimType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectLogin) *string { return v.NameClaimType }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectLoginOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v OpenIdConnectLogin) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type OpenIdConnectLoginPtrOutput struct{ *pulumi.OutputState }

func (OpenIdConnectLoginPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectLogin)(nil)).Elem()
}

func (o OpenIdConnectLoginPtrOutput) ToOpenIdConnectLoginPtrOutput() OpenIdConnectLoginPtrOutput {
	return o
}

func (o OpenIdConnectLoginPtrOutput) ToOpenIdConnectLoginPtrOutputWithContext(ctx context.Context) OpenIdConnectLoginPtrOutput {
	return o
}

func (o OpenIdConnectLoginPtrOutput) Elem() OpenIdConnectLoginOutput {
	return o.ApplyT(func(v *OpenIdConnectLogin) OpenIdConnectLogin {
		if v != nil {
			return *v
		}
		var ret OpenIdConnectLogin
		return ret
	}).(OpenIdConnectLoginOutput)
}

func (o OpenIdConnectLoginPtrOutput) NameClaimType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectLogin) *string {
		if v == nil {
			return nil
		}
		return v.NameClaimType
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectLoginPtrOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OpenIdConnectLogin) []string {
		if v == nil {
			return nil
		}
		return v.Scopes
	}).(pulumi.StringArrayOutput)
}

type OpenIdConnectLoginResponse struct {
	NameClaimType *string  `pulumi:"nameClaimType"`
	Scopes        []string `pulumi:"scopes"`
}

type OpenIdConnectLoginResponseOutput struct{ *pulumi.OutputState }

func (OpenIdConnectLoginResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectLoginResponse)(nil)).Elem()
}

func (o OpenIdConnectLoginResponseOutput) ToOpenIdConnectLoginResponseOutput() OpenIdConnectLoginResponseOutput {
	return o
}

func (o OpenIdConnectLoginResponseOutput) ToOpenIdConnectLoginResponseOutputWithContext(ctx context.Context) OpenIdConnectLoginResponseOutput {
	return o
}

func (o OpenIdConnectLoginResponseOutput) NameClaimType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectLoginResponse) *string { return v.NameClaimType }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectLoginResponseOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v OpenIdConnectLoginResponse) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type OpenIdConnectLoginResponsePtrOutput struct{ *pulumi.OutputState }

func (OpenIdConnectLoginResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectLoginResponse)(nil)).Elem()
}

func (o OpenIdConnectLoginResponsePtrOutput) ToOpenIdConnectLoginResponsePtrOutput() OpenIdConnectLoginResponsePtrOutput {
	return o
}

func (o OpenIdConnectLoginResponsePtrOutput) ToOpenIdConnectLoginResponsePtrOutputWithContext(ctx context.Context) OpenIdConnectLoginResponsePtrOutput {
	return o
}

func (o OpenIdConnectLoginResponsePtrOutput) Elem() OpenIdConnectLoginResponseOutput {
	return o.ApplyT(func(v *OpenIdConnectLoginResponse) OpenIdConnectLoginResponse {
		if v != nil {
			return *v
		}
		var ret OpenIdConnectLoginResponse
		return ret
	}).(OpenIdConnectLoginResponseOutput)
}

func (o OpenIdConnectLoginResponsePtrOutput) NameClaimType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectLoginResponse) *string {
		if v == nil {
			return nil
		}
		return v.NameClaimType
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectLoginResponsePtrOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OpenIdConnectLoginResponse) []string {
		if v == nil {
			return nil
		}
		return v.Scopes
	}).(pulumi.StringArrayOutput)
}

type OpenIdConnectRegistration struct {
	ClientCredential           *OpenIdConnectClientCredential `pulumi:"clientCredential"`
	ClientId                   *string                        `pulumi:"clientId"`
	OpenIdConnectConfiguration *OpenIdConnectConfig           `pulumi:"openIdConnectConfiguration"`
}





type OpenIdConnectRegistrationInput interface {
	pulumi.Input

	ToOpenIdConnectRegistrationOutput() OpenIdConnectRegistrationOutput
	ToOpenIdConnectRegistrationOutputWithContext(context.Context) OpenIdConnectRegistrationOutput
}

type OpenIdConnectRegistrationArgs struct {
	ClientCredential           OpenIdConnectClientCredentialPtrInput `pulumi:"clientCredential"`
	ClientId                   pulumi.StringPtrInput                 `pulumi:"clientId"`
	OpenIdConnectConfiguration OpenIdConnectConfigPtrInput           `pulumi:"openIdConnectConfiguration"`
}

func (OpenIdConnectRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectRegistration)(nil)).Elem()
}

func (i OpenIdConnectRegistrationArgs) ToOpenIdConnectRegistrationOutput() OpenIdConnectRegistrationOutput {
	return i.ToOpenIdConnectRegistrationOutputWithContext(context.Background())
}

func (i OpenIdConnectRegistrationArgs) ToOpenIdConnectRegistrationOutputWithContext(ctx context.Context) OpenIdConnectRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectRegistrationOutput)
}

func (i OpenIdConnectRegistrationArgs) ToOpenIdConnectRegistrationPtrOutput() OpenIdConnectRegistrationPtrOutput {
	return i.ToOpenIdConnectRegistrationPtrOutputWithContext(context.Background())
}

func (i OpenIdConnectRegistrationArgs) ToOpenIdConnectRegistrationPtrOutputWithContext(ctx context.Context) OpenIdConnectRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectRegistrationOutput).ToOpenIdConnectRegistrationPtrOutputWithContext(ctx)
}









type OpenIdConnectRegistrationPtrInput interface {
	pulumi.Input

	ToOpenIdConnectRegistrationPtrOutput() OpenIdConnectRegistrationPtrOutput
	ToOpenIdConnectRegistrationPtrOutputWithContext(context.Context) OpenIdConnectRegistrationPtrOutput
}

type openIdConnectRegistrationPtrType OpenIdConnectRegistrationArgs

func OpenIdConnectRegistrationPtr(v *OpenIdConnectRegistrationArgs) OpenIdConnectRegistrationPtrInput {
	return (*openIdConnectRegistrationPtrType)(v)
}

func (*openIdConnectRegistrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectRegistration)(nil)).Elem()
}

func (i *openIdConnectRegistrationPtrType) ToOpenIdConnectRegistrationPtrOutput() OpenIdConnectRegistrationPtrOutput {
	return i.ToOpenIdConnectRegistrationPtrOutputWithContext(context.Background())
}

func (i *openIdConnectRegistrationPtrType) ToOpenIdConnectRegistrationPtrOutputWithContext(ctx context.Context) OpenIdConnectRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectRegistrationPtrOutput)
}

type OpenIdConnectRegistrationOutput struct{ *pulumi.OutputState }

func (OpenIdConnectRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectRegistration)(nil)).Elem()
}

func (o OpenIdConnectRegistrationOutput) ToOpenIdConnectRegistrationOutput() OpenIdConnectRegistrationOutput {
	return o
}

func (o OpenIdConnectRegistrationOutput) ToOpenIdConnectRegistrationOutputWithContext(ctx context.Context) OpenIdConnectRegistrationOutput {
	return o
}

func (o OpenIdConnectRegistrationOutput) ToOpenIdConnectRegistrationPtrOutput() OpenIdConnectRegistrationPtrOutput {
	return o.ToOpenIdConnectRegistrationPtrOutputWithContext(context.Background())
}

func (o OpenIdConnectRegistrationOutput) ToOpenIdConnectRegistrationPtrOutputWithContext(ctx context.Context) OpenIdConnectRegistrationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OpenIdConnectRegistration) *OpenIdConnectRegistration {
		return &v
	}).(OpenIdConnectRegistrationPtrOutput)
}

func (o OpenIdConnectRegistrationOutput) ClientCredential() OpenIdConnectClientCredentialPtrOutput {
	return o.ApplyT(func(v OpenIdConnectRegistration) *OpenIdConnectClientCredential { return v.ClientCredential }).(OpenIdConnectClientCredentialPtrOutput)
}

func (o OpenIdConnectRegistrationOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectRegistration) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectRegistrationOutput) OpenIdConnectConfiguration() OpenIdConnectConfigPtrOutput {
	return o.ApplyT(func(v OpenIdConnectRegistration) *OpenIdConnectConfig { return v.OpenIdConnectConfiguration }).(OpenIdConnectConfigPtrOutput)
}

type OpenIdConnectRegistrationPtrOutput struct{ *pulumi.OutputState }

func (OpenIdConnectRegistrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectRegistration)(nil)).Elem()
}

func (o OpenIdConnectRegistrationPtrOutput) ToOpenIdConnectRegistrationPtrOutput() OpenIdConnectRegistrationPtrOutput {
	return o
}

func (o OpenIdConnectRegistrationPtrOutput) ToOpenIdConnectRegistrationPtrOutputWithContext(ctx context.Context) OpenIdConnectRegistrationPtrOutput {
	return o
}

func (o OpenIdConnectRegistrationPtrOutput) Elem() OpenIdConnectRegistrationOutput {
	return o.ApplyT(func(v *OpenIdConnectRegistration) OpenIdConnectRegistration {
		if v != nil {
			return *v
		}
		var ret OpenIdConnectRegistration
		return ret
	}).(OpenIdConnectRegistrationOutput)
}

func (o OpenIdConnectRegistrationPtrOutput) ClientCredential() OpenIdConnectClientCredentialPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectRegistration) *OpenIdConnectClientCredential {
		if v == nil {
			return nil
		}
		return v.ClientCredential
	}).(OpenIdConnectClientCredentialPtrOutput)
}

func (o OpenIdConnectRegistrationPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectRegistrationPtrOutput) OpenIdConnectConfiguration() OpenIdConnectConfigPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectRegistration) *OpenIdConnectConfig {
		if v == nil {
			return nil
		}
		return v.OpenIdConnectConfiguration
	}).(OpenIdConnectConfigPtrOutput)
}

type OpenIdConnectRegistrationResponse struct {
	ClientCredential           *OpenIdConnectClientCredentialResponse `pulumi:"clientCredential"`
	ClientId                   *string                                `pulumi:"clientId"`
	OpenIdConnectConfiguration *OpenIdConnectConfigResponse           `pulumi:"openIdConnectConfiguration"`
}

type OpenIdConnectRegistrationResponseOutput struct{ *pulumi.OutputState }

func (OpenIdConnectRegistrationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OpenIdConnectRegistrationResponse)(nil)).Elem()
}

func (o OpenIdConnectRegistrationResponseOutput) ToOpenIdConnectRegistrationResponseOutput() OpenIdConnectRegistrationResponseOutput {
	return o
}

func (o OpenIdConnectRegistrationResponseOutput) ToOpenIdConnectRegistrationResponseOutputWithContext(ctx context.Context) OpenIdConnectRegistrationResponseOutput {
	return o
}

func (o OpenIdConnectRegistrationResponseOutput) ClientCredential() OpenIdConnectClientCredentialResponsePtrOutput {
	return o.ApplyT(func(v OpenIdConnectRegistrationResponse) *OpenIdConnectClientCredentialResponse {
		return v.ClientCredential
	}).(OpenIdConnectClientCredentialResponsePtrOutput)
}

func (o OpenIdConnectRegistrationResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OpenIdConnectRegistrationResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectRegistrationResponseOutput) OpenIdConnectConfiguration() OpenIdConnectConfigResponsePtrOutput {
	return o.ApplyT(func(v OpenIdConnectRegistrationResponse) *OpenIdConnectConfigResponse {
		return v.OpenIdConnectConfiguration
	}).(OpenIdConnectConfigResponsePtrOutput)
}

type OpenIdConnectRegistrationResponsePtrOutput struct{ *pulumi.OutputState }

func (OpenIdConnectRegistrationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectRegistrationResponse)(nil)).Elem()
}

func (o OpenIdConnectRegistrationResponsePtrOutput) ToOpenIdConnectRegistrationResponsePtrOutput() OpenIdConnectRegistrationResponsePtrOutput {
	return o
}

func (o OpenIdConnectRegistrationResponsePtrOutput) ToOpenIdConnectRegistrationResponsePtrOutputWithContext(ctx context.Context) OpenIdConnectRegistrationResponsePtrOutput {
	return o
}

func (o OpenIdConnectRegistrationResponsePtrOutput) Elem() OpenIdConnectRegistrationResponseOutput {
	return o.ApplyT(func(v *OpenIdConnectRegistrationResponse) OpenIdConnectRegistrationResponse {
		if v != nil {
			return *v
		}
		var ret OpenIdConnectRegistrationResponse
		return ret
	}).(OpenIdConnectRegistrationResponseOutput)
}

func (o OpenIdConnectRegistrationResponsePtrOutput) ClientCredential() OpenIdConnectClientCredentialResponsePtrOutput {
	return o.ApplyT(func(v *OpenIdConnectRegistrationResponse) *OpenIdConnectClientCredentialResponse {
		if v == nil {
			return nil
		}
		return v.ClientCredential
	}).(OpenIdConnectClientCredentialResponsePtrOutput)
}

func (o OpenIdConnectRegistrationResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o OpenIdConnectRegistrationResponsePtrOutput) OpenIdConnectConfiguration() OpenIdConnectConfigResponsePtrOutput {
	return o.ApplyT(func(v *OpenIdConnectRegistrationResponse) *OpenIdConnectConfigResponse {
		if v == nil {
			return nil
		}
		return v.OpenIdConnectConfiguration
	}).(OpenIdConnectConfigResponsePtrOutput)
}

type PrivateLinkConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkConnectionStateOutput() PrivateLinkConnectionStateOutput
	ToPrivateLinkConnectionStateOutputWithContext(context.Context) PrivateLinkConnectionStateOutput
}

type PrivateLinkConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConnectionState)(nil)).Elem()
}

func (i PrivateLinkConnectionStateArgs) ToPrivateLinkConnectionStateOutput() PrivateLinkConnectionStateOutput {
	return i.ToPrivateLinkConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkConnectionStateArgs) ToPrivateLinkConnectionStateOutputWithContext(ctx context.Context) PrivateLinkConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionStateOutput)
}

func (i PrivateLinkConnectionStateArgs) ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput {
	return i.ToPrivateLinkConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkConnectionStateArgs) ToPrivateLinkConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionStateOutput).ToPrivateLinkConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput
	ToPrivateLinkConnectionStatePtrOutputWithContext(context.Context) PrivateLinkConnectionStatePtrOutput
}

type privateLinkConnectionStatePtrType PrivateLinkConnectionStateArgs

func PrivateLinkConnectionStatePtr(v *PrivateLinkConnectionStateArgs) PrivateLinkConnectionStatePtrInput {
	return (*privateLinkConnectionStatePtrType)(v)
}

func (*privateLinkConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkConnectionState)(nil)).Elem()
}

func (i *privateLinkConnectionStatePtrType) ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput {
	return i.ToPrivateLinkConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkConnectionStatePtrType) ToPrivateLinkConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionStatePtrOutput)
}

type PrivateLinkConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConnectionState)(nil)).Elem()
}

func (o PrivateLinkConnectionStateOutput) ToPrivateLinkConnectionStateOutput() PrivateLinkConnectionStateOutput {
	return o
}

func (o PrivateLinkConnectionStateOutput) ToPrivateLinkConnectionStateOutputWithContext(ctx context.Context) PrivateLinkConnectionStateOutput {
	return o
}

func (o PrivateLinkConnectionStateOutput) ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput {
	return o.ToPrivateLinkConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkConnectionStateOutput) ToPrivateLinkConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkConnectionState) *PrivateLinkConnectionState {
		return &v
	}).(PrivateLinkConnectionStatePtrOutput)
}

func (o PrivateLinkConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkConnectionState)(nil)).Elem()
}

func (o PrivateLinkConnectionStatePtrOutput) ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkConnectionStatePtrOutput) ToPrivateLinkConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkConnectionStatePtrOutput) Elem() PrivateLinkConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionState) PrivateLinkConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkConnectionState
		return ret
	}).(PrivateLinkConnectionStateOutput)
}

func (o PrivateLinkConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}

type PrivateLinkConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkConnectionStateResponseOutput) ToPrivateLinkConnectionStateResponseOutput() PrivateLinkConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkConnectionStateResponseOutput) ToPrivateLinkConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkConnectionStateResponsePtrOutput) ToPrivateLinkConnectionStateResponsePtrOutput() PrivateLinkConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkConnectionStateResponsePtrOutput) ToPrivateLinkConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkConnectionStateResponsePtrOutput) Elem() PrivateLinkConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) PrivateLinkConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkConnectionStateResponse
		return ret
	}).(PrivateLinkConnectionStateResponseOutput)
}

func (o PrivateLinkConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PushSettings struct {
	DynamicTagsJson   *string `pulumi:"dynamicTagsJson"`
	IsPushEnabled     bool    `pulumi:"isPushEnabled"`
	Kind              *string `pulumi:"kind"`
	TagWhitelistJson  *string `pulumi:"tagWhitelistJson"`
	TagsRequiringAuth *string `pulumi:"tagsRequiringAuth"`
}





type PushSettingsInput interface {
	pulumi.Input

	ToPushSettingsOutput() PushSettingsOutput
	ToPushSettingsOutputWithContext(context.Context) PushSettingsOutput
}

type PushSettingsArgs struct {
	DynamicTagsJson   pulumi.StringPtrInput `pulumi:"dynamicTagsJson"`
	IsPushEnabled     pulumi.BoolInput      `pulumi:"isPushEnabled"`
	Kind              pulumi.StringPtrInput `pulumi:"kind"`
	TagWhitelistJson  pulumi.StringPtrInput `pulumi:"tagWhitelistJson"`
	TagsRequiringAuth pulumi.StringPtrInput `pulumi:"tagsRequiringAuth"`
}

func (PushSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PushSettings)(nil)).Elem()
}

func (i PushSettingsArgs) ToPushSettingsOutput() PushSettingsOutput {
	return i.ToPushSettingsOutputWithContext(context.Background())
}

func (i PushSettingsArgs) ToPushSettingsOutputWithContext(ctx context.Context) PushSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushSettingsOutput)
}

func (i PushSettingsArgs) ToPushSettingsPtrOutput() PushSettingsPtrOutput {
	return i.ToPushSettingsPtrOutputWithContext(context.Background())
}

func (i PushSettingsArgs) ToPushSettingsPtrOutputWithContext(ctx context.Context) PushSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushSettingsOutput).ToPushSettingsPtrOutputWithContext(ctx)
}









type PushSettingsPtrInput interface {
	pulumi.Input

	ToPushSettingsPtrOutput() PushSettingsPtrOutput
	ToPushSettingsPtrOutputWithContext(context.Context) PushSettingsPtrOutput
}

type pushSettingsPtrType PushSettingsArgs

func PushSettingsPtr(v *PushSettingsArgs) PushSettingsPtrInput {
	return (*pushSettingsPtrType)(v)
}

func (*pushSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PushSettings)(nil)).Elem()
}

func (i *pushSettingsPtrType) ToPushSettingsPtrOutput() PushSettingsPtrOutput {
	return i.ToPushSettingsPtrOutputWithContext(context.Background())
}

func (i *pushSettingsPtrType) ToPushSettingsPtrOutputWithContext(ctx context.Context) PushSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushSettingsPtrOutput)
}

type PushSettingsOutput struct{ *pulumi.OutputState }

func (PushSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PushSettings)(nil)).Elem()
}

func (o PushSettingsOutput) ToPushSettingsOutput() PushSettingsOutput {
	return o
}

func (o PushSettingsOutput) ToPushSettingsOutputWithContext(ctx context.Context) PushSettingsOutput {
	return o
}

func (o PushSettingsOutput) ToPushSettingsPtrOutput() PushSettingsPtrOutput {
	return o.ToPushSettingsPtrOutputWithContext(context.Background())
}

func (o PushSettingsOutput) ToPushSettingsPtrOutputWithContext(ctx context.Context) PushSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PushSettings) *PushSettings {
		return &v
	}).(PushSettingsPtrOutput)
}

func (o PushSettingsOutput) DynamicTagsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PushSettings) *string { return v.DynamicTagsJson }).(pulumi.StringPtrOutput)
}

func (o PushSettingsOutput) IsPushEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v PushSettings) bool { return v.IsPushEnabled }).(pulumi.BoolOutput)
}

func (o PushSettingsOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PushSettings) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o PushSettingsOutput) TagWhitelistJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PushSettings) *string { return v.TagWhitelistJson }).(pulumi.StringPtrOutput)
}

func (o PushSettingsOutput) TagsRequiringAuth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PushSettings) *string { return v.TagsRequiringAuth }).(pulumi.StringPtrOutput)
}

type PushSettingsPtrOutput struct{ *pulumi.OutputState }

func (PushSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PushSettings)(nil)).Elem()
}

func (o PushSettingsPtrOutput) ToPushSettingsPtrOutput() PushSettingsPtrOutput {
	return o
}

func (o PushSettingsPtrOutput) ToPushSettingsPtrOutputWithContext(ctx context.Context) PushSettingsPtrOutput {
	return o
}

func (o PushSettingsPtrOutput) Elem() PushSettingsOutput {
	return o.ApplyT(func(v *PushSettings) PushSettings {
		if v != nil {
			return *v
		}
		var ret PushSettings
		return ret
	}).(PushSettingsOutput)
}

func (o PushSettingsPtrOutput) DynamicTagsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettings) *string {
		if v == nil {
			return nil
		}
		return v.DynamicTagsJson
	}).(pulumi.StringPtrOutput)
}

func (o PushSettingsPtrOutput) IsPushEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PushSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.IsPushEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o PushSettingsPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettings) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o PushSettingsPtrOutput) TagWhitelistJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettings) *string {
		if v == nil {
			return nil
		}
		return v.TagWhitelistJson
	}).(pulumi.StringPtrOutput)
}

func (o PushSettingsPtrOutput) TagsRequiringAuth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettings) *string {
		if v == nil {
			return nil
		}
		return v.TagsRequiringAuth
	}).(pulumi.StringPtrOutput)
}

type PushSettingsResponse struct {
	DynamicTagsJson   *string `pulumi:"dynamicTagsJson"`
	Id                string  `pulumi:"id"`
	IsPushEnabled     bool    `pulumi:"isPushEnabled"`
	Kind              *string `pulumi:"kind"`
	Name              string  `pulumi:"name"`
	TagWhitelistJson  *string `pulumi:"tagWhitelistJson"`
	TagsRequiringAuth *string `pulumi:"tagsRequiringAuth"`
	Type              string  `pulumi:"type"`
}

type PushSettingsResponseOutput struct{ *pulumi.OutputState }

func (PushSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PushSettingsResponse)(nil)).Elem()
}

func (o PushSettingsResponseOutput) ToPushSettingsResponseOutput() PushSettingsResponseOutput {
	return o
}

func (o PushSettingsResponseOutput) ToPushSettingsResponseOutputWithContext(ctx context.Context) PushSettingsResponseOutput {
	return o
}

func (o PushSettingsResponseOutput) DynamicTagsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PushSettingsResponse) *string { return v.DynamicTagsJson }).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PushSettingsResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PushSettingsResponseOutput) IsPushEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v PushSettingsResponse) bool { return v.IsPushEnabled }).(pulumi.BoolOutput)
}

func (o PushSettingsResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PushSettingsResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PushSettingsResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PushSettingsResponseOutput) TagWhitelistJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PushSettingsResponse) *string { return v.TagWhitelistJson }).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponseOutput) TagsRequiringAuth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PushSettingsResponse) *string { return v.TagsRequiringAuth }).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PushSettingsResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PushSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (PushSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PushSettingsResponse)(nil)).Elem()
}

func (o PushSettingsResponsePtrOutput) ToPushSettingsResponsePtrOutput() PushSettingsResponsePtrOutput {
	return o
}

func (o PushSettingsResponsePtrOutput) ToPushSettingsResponsePtrOutputWithContext(ctx context.Context) PushSettingsResponsePtrOutput {
	return o
}

func (o PushSettingsResponsePtrOutput) Elem() PushSettingsResponseOutput {
	return o.ApplyT(func(v *PushSettingsResponse) PushSettingsResponse {
		if v != nil {
			return *v
		}
		var ret PushSettingsResponse
		return ret
	}).(PushSettingsResponseOutput)
}

func (o PushSettingsResponsePtrOutput) DynamicTagsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DynamicTagsJson
	}).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponsePtrOutput) IsPushEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PushSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsPushEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o PushSettingsResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponsePtrOutput) TagWhitelistJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TagWhitelistJson
	}).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponsePtrOutput) TagsRequiringAuth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TagsRequiringAuth
	}).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type RampUpRule struct {
	ActionHostName            *string  `pulumi:"actionHostName"`
	ChangeDecisionCallbackUrl *string  `pulumi:"changeDecisionCallbackUrl"`
	ChangeIntervalInMinutes   *int     `pulumi:"changeIntervalInMinutes"`
	ChangeStep                *float64 `pulumi:"changeStep"`
	MaxReroutePercentage      *float64 `pulumi:"maxReroutePercentage"`
	MinReroutePercentage      *float64 `pulumi:"minReroutePercentage"`
	Name                      *string  `pulumi:"name"`
	ReroutePercentage         *float64 `pulumi:"reroutePercentage"`
}





type RampUpRuleInput interface {
	pulumi.Input

	ToRampUpRuleOutput() RampUpRuleOutput
	ToRampUpRuleOutputWithContext(context.Context) RampUpRuleOutput
}

type RampUpRuleArgs struct {
	ActionHostName            pulumi.StringPtrInput  `pulumi:"actionHostName"`
	ChangeDecisionCallbackUrl pulumi.StringPtrInput  `pulumi:"changeDecisionCallbackUrl"`
	ChangeIntervalInMinutes   pulumi.IntPtrInput     `pulumi:"changeIntervalInMinutes"`
	ChangeStep                pulumi.Float64PtrInput `pulumi:"changeStep"`
	MaxReroutePercentage      pulumi.Float64PtrInput `pulumi:"maxReroutePercentage"`
	MinReroutePercentage      pulumi.Float64PtrInput `pulumi:"minReroutePercentage"`
	Name                      pulumi.StringPtrInput  `pulumi:"name"`
	ReroutePercentage         pulumi.Float64PtrInput `pulumi:"reroutePercentage"`
}

func (RampUpRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RampUpRule)(nil)).Elem()
}

func (i RampUpRuleArgs) ToRampUpRuleOutput() RampUpRuleOutput {
	return i.ToRampUpRuleOutputWithContext(context.Background())
}

func (i RampUpRuleArgs) ToRampUpRuleOutputWithContext(ctx context.Context) RampUpRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RampUpRuleOutput)
}





type RampUpRuleArrayInput interface {
	pulumi.Input

	ToRampUpRuleArrayOutput() RampUpRuleArrayOutput
	ToRampUpRuleArrayOutputWithContext(context.Context) RampUpRuleArrayOutput
}

type RampUpRuleArray []RampUpRuleInput

func (RampUpRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RampUpRule)(nil)).Elem()
}

func (i RampUpRuleArray) ToRampUpRuleArrayOutput() RampUpRuleArrayOutput {
	return i.ToRampUpRuleArrayOutputWithContext(context.Background())
}

func (i RampUpRuleArray) ToRampUpRuleArrayOutputWithContext(ctx context.Context) RampUpRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RampUpRuleArrayOutput)
}

type RampUpRuleOutput struct{ *pulumi.OutputState }

func (RampUpRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RampUpRule)(nil)).Elem()
}

func (o RampUpRuleOutput) ToRampUpRuleOutput() RampUpRuleOutput {
	return o
}

func (o RampUpRuleOutput) ToRampUpRuleOutputWithContext(ctx context.Context) RampUpRuleOutput {
	return o
}

func (o RampUpRuleOutput) ActionHostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RampUpRule) *string { return v.ActionHostName }).(pulumi.StringPtrOutput)
}

func (o RampUpRuleOutput) ChangeDecisionCallbackUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RampUpRule) *string { return v.ChangeDecisionCallbackUrl }).(pulumi.StringPtrOutput)
}

func (o RampUpRuleOutput) ChangeIntervalInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RampUpRule) *int { return v.ChangeIntervalInMinutes }).(pulumi.IntPtrOutput)
}

func (o RampUpRuleOutput) ChangeStep() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRule) *float64 { return v.ChangeStep }).(pulumi.Float64PtrOutput)
}

func (o RampUpRuleOutput) MaxReroutePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRule) *float64 { return v.MaxReroutePercentage }).(pulumi.Float64PtrOutput)
}

func (o RampUpRuleOutput) MinReroutePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRule) *float64 { return v.MinReroutePercentage }).(pulumi.Float64PtrOutput)
}

func (o RampUpRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RampUpRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RampUpRuleOutput) ReroutePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRule) *float64 { return v.ReroutePercentage }).(pulumi.Float64PtrOutput)
}

type RampUpRuleArrayOutput struct{ *pulumi.OutputState }

func (RampUpRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RampUpRule)(nil)).Elem()
}

func (o RampUpRuleArrayOutput) ToRampUpRuleArrayOutput() RampUpRuleArrayOutput {
	return o
}

func (o RampUpRuleArrayOutput) ToRampUpRuleArrayOutputWithContext(ctx context.Context) RampUpRuleArrayOutput {
	return o
}

func (o RampUpRuleArrayOutput) Index(i pulumi.IntInput) RampUpRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RampUpRule {
		return vs[0].([]RampUpRule)[vs[1].(int)]
	}).(RampUpRuleOutput)
}

type RampUpRuleResponse struct {
	ActionHostName            *string  `pulumi:"actionHostName"`
	ChangeDecisionCallbackUrl *string  `pulumi:"changeDecisionCallbackUrl"`
	ChangeIntervalInMinutes   *int     `pulumi:"changeIntervalInMinutes"`
	ChangeStep                *float64 `pulumi:"changeStep"`
	MaxReroutePercentage      *float64 `pulumi:"maxReroutePercentage"`
	MinReroutePercentage      *float64 `pulumi:"minReroutePercentage"`
	Name                      *string  `pulumi:"name"`
	ReroutePercentage         *float64 `pulumi:"reroutePercentage"`
}

type RampUpRuleResponseOutput struct{ *pulumi.OutputState }

func (RampUpRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RampUpRuleResponse)(nil)).Elem()
}

func (o RampUpRuleResponseOutput) ToRampUpRuleResponseOutput() RampUpRuleResponseOutput {
	return o
}

func (o RampUpRuleResponseOutput) ToRampUpRuleResponseOutputWithContext(ctx context.Context) RampUpRuleResponseOutput {
	return o
}

func (o RampUpRuleResponseOutput) ActionHostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *string { return v.ActionHostName }).(pulumi.StringPtrOutput)
}

func (o RampUpRuleResponseOutput) ChangeDecisionCallbackUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *string { return v.ChangeDecisionCallbackUrl }).(pulumi.StringPtrOutput)
}

func (o RampUpRuleResponseOutput) ChangeIntervalInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *int { return v.ChangeIntervalInMinutes }).(pulumi.IntPtrOutput)
}

func (o RampUpRuleResponseOutput) ChangeStep() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *float64 { return v.ChangeStep }).(pulumi.Float64PtrOutput)
}

func (o RampUpRuleResponseOutput) MaxReroutePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *float64 { return v.MaxReroutePercentage }).(pulumi.Float64PtrOutput)
}

func (o RampUpRuleResponseOutput) MinReroutePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *float64 { return v.MinReroutePercentage }).(pulumi.Float64PtrOutput)
}

func (o RampUpRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RampUpRuleResponseOutput) ReroutePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *float64 { return v.ReroutePercentage }).(pulumi.Float64PtrOutput)
}

type RampUpRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (RampUpRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RampUpRuleResponse)(nil)).Elem()
}

func (o RampUpRuleResponseArrayOutput) ToRampUpRuleResponseArrayOutput() RampUpRuleResponseArrayOutput {
	return o
}

func (o RampUpRuleResponseArrayOutput) ToRampUpRuleResponseArrayOutputWithContext(ctx context.Context) RampUpRuleResponseArrayOutput {
	return o
}

func (o RampUpRuleResponseArrayOutput) Index(i pulumi.IntInput) RampUpRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RampUpRuleResponse {
		return vs[0].([]RampUpRuleResponse)[vs[1].(int)]
	}).(RampUpRuleResponseOutput)
}

type RemotePrivateEndpointConnectionResponse struct {
	Id                                string                              `pulumi:"id"`
	IpAddresses                       []string                            `pulumi:"ipAddresses"`
	Kind                              *string                             `pulumi:"kind"`
	Name                              string                              `pulumi:"name"`
	PrivateEndpoint                   *ArmIdWrapperResponse               `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                              `pulumi:"provisioningState"`
	Type                              string                              `pulumi:"type"`
}

type RemotePrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (RemotePrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemotePrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o RemotePrivateEndpointConnectionResponseOutput) ToRemotePrivateEndpointConnectionResponseOutput() RemotePrivateEndpointConnectionResponseOutput {
	return o
}

func (o RemotePrivateEndpointConnectionResponseOutput) ToRemotePrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) RemotePrivateEndpointConnectionResponseOutput {
	return o
}

func (o RemotePrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o RemotePrivateEndpointConnectionResponseOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o RemotePrivateEndpointConnectionResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RemotePrivateEndpointConnectionResponseOutput) PrivateEndpoint() ArmIdWrapperResponsePtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) *ArmIdWrapperResponse { return v.PrivateEndpoint }).(ArmIdWrapperResponsePtrOutput)
}

func (o RemotePrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) *PrivateLinkConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStateResponsePtrOutput)
}

func (o RemotePrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RemotePrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type RemotePrivateEndpointConnectionResponsePtrOutput struct{ *pulumi.OutputState }

func (RemotePrivateEndpointConnectionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemotePrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) ToRemotePrivateEndpointConnectionResponsePtrOutput() RemotePrivateEndpointConnectionResponsePtrOutput {
	return o
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) ToRemotePrivateEndpointConnectionResponsePtrOutputWithContext(ctx context.Context) RemotePrivateEndpointConnectionResponsePtrOutput {
	return o
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) Elem() RemotePrivateEndpointConnectionResponseOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointConnectionResponse) RemotePrivateEndpointConnectionResponse {
		if v != nil {
			return *v
		}
		var ret RemotePrivateEndpointConnectionResponse
		return ret
	}).(RemotePrivateEndpointConnectionResponseOutput)
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointConnectionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointConnectionResponse) []string {
		if v == nil {
			return nil
		}
		return v.IpAddresses
	}).(pulumi.StringArrayOutput)
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointConnectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointConnectionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) PrivateEndpoint() ArmIdWrapperResponsePtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointConnectionResponse) *ArmIdWrapperResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(ArmIdWrapperResponsePtrOutput)
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointConnectionResponse) *PrivateLinkConnectionStateResponse {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStateResponsePtrOutput)
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointConnectionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointConnectionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type RequestsBasedTrigger struct {
	Count        *int    `pulumi:"count"`
	TimeInterval *string `pulumi:"timeInterval"`
}





type RequestsBasedTriggerInput interface {
	pulumi.Input

	ToRequestsBasedTriggerOutput() RequestsBasedTriggerOutput
	ToRequestsBasedTriggerOutputWithContext(context.Context) RequestsBasedTriggerOutput
}

type RequestsBasedTriggerArgs struct {
	Count        pulumi.IntPtrInput    `pulumi:"count"`
	TimeInterval pulumi.StringPtrInput `pulumi:"timeInterval"`
}

func (RequestsBasedTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestsBasedTrigger)(nil)).Elem()
}

func (i RequestsBasedTriggerArgs) ToRequestsBasedTriggerOutput() RequestsBasedTriggerOutput {
	return i.ToRequestsBasedTriggerOutputWithContext(context.Background())
}

func (i RequestsBasedTriggerArgs) ToRequestsBasedTriggerOutputWithContext(ctx context.Context) RequestsBasedTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestsBasedTriggerOutput)
}

func (i RequestsBasedTriggerArgs) ToRequestsBasedTriggerPtrOutput() RequestsBasedTriggerPtrOutput {
	return i.ToRequestsBasedTriggerPtrOutputWithContext(context.Background())
}

func (i RequestsBasedTriggerArgs) ToRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) RequestsBasedTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestsBasedTriggerOutput).ToRequestsBasedTriggerPtrOutputWithContext(ctx)
}









type RequestsBasedTriggerPtrInput interface {
	pulumi.Input

	ToRequestsBasedTriggerPtrOutput() RequestsBasedTriggerPtrOutput
	ToRequestsBasedTriggerPtrOutputWithContext(context.Context) RequestsBasedTriggerPtrOutput
}

type requestsBasedTriggerPtrType RequestsBasedTriggerArgs

func RequestsBasedTriggerPtr(v *RequestsBasedTriggerArgs) RequestsBasedTriggerPtrInput {
	return (*requestsBasedTriggerPtrType)(v)
}

func (*requestsBasedTriggerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestsBasedTrigger)(nil)).Elem()
}

func (i *requestsBasedTriggerPtrType) ToRequestsBasedTriggerPtrOutput() RequestsBasedTriggerPtrOutput {
	return i.ToRequestsBasedTriggerPtrOutputWithContext(context.Background())
}

func (i *requestsBasedTriggerPtrType) ToRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) RequestsBasedTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestsBasedTriggerPtrOutput)
}

type RequestsBasedTriggerOutput struct{ *pulumi.OutputState }

func (RequestsBasedTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestsBasedTrigger)(nil)).Elem()
}

func (o RequestsBasedTriggerOutput) ToRequestsBasedTriggerOutput() RequestsBasedTriggerOutput {
	return o
}

func (o RequestsBasedTriggerOutput) ToRequestsBasedTriggerOutputWithContext(ctx context.Context) RequestsBasedTriggerOutput {
	return o
}

func (o RequestsBasedTriggerOutput) ToRequestsBasedTriggerPtrOutput() RequestsBasedTriggerPtrOutput {
	return o.ToRequestsBasedTriggerPtrOutputWithContext(context.Background())
}

func (o RequestsBasedTriggerOutput) ToRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) RequestsBasedTriggerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RequestsBasedTrigger) *RequestsBasedTrigger {
		return &v
	}).(RequestsBasedTriggerPtrOutput)
}

func (o RequestsBasedTriggerOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RequestsBasedTrigger) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o RequestsBasedTriggerOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RequestsBasedTrigger) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

type RequestsBasedTriggerPtrOutput struct{ *pulumi.OutputState }

func (RequestsBasedTriggerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestsBasedTrigger)(nil)).Elem()
}

func (o RequestsBasedTriggerPtrOutput) ToRequestsBasedTriggerPtrOutput() RequestsBasedTriggerPtrOutput {
	return o
}

func (o RequestsBasedTriggerPtrOutput) ToRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) RequestsBasedTriggerPtrOutput {
	return o
}

func (o RequestsBasedTriggerPtrOutput) Elem() RequestsBasedTriggerOutput {
	return o.ApplyT(func(v *RequestsBasedTrigger) RequestsBasedTrigger {
		if v != nil {
			return *v
		}
		var ret RequestsBasedTrigger
		return ret
	}).(RequestsBasedTriggerOutput)
}

func (o RequestsBasedTriggerPtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RequestsBasedTrigger) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o RequestsBasedTriggerPtrOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestsBasedTrigger) *string {
		if v == nil {
			return nil
		}
		return v.TimeInterval
	}).(pulumi.StringPtrOutput)
}

type RequestsBasedTriggerResponse struct {
	Count        *int    `pulumi:"count"`
	TimeInterval *string `pulumi:"timeInterval"`
}

type RequestsBasedTriggerResponseOutput struct{ *pulumi.OutputState }

func (RequestsBasedTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestsBasedTriggerResponse)(nil)).Elem()
}

func (o RequestsBasedTriggerResponseOutput) ToRequestsBasedTriggerResponseOutput() RequestsBasedTriggerResponseOutput {
	return o
}

func (o RequestsBasedTriggerResponseOutput) ToRequestsBasedTriggerResponseOutputWithContext(ctx context.Context) RequestsBasedTriggerResponseOutput {
	return o
}

func (o RequestsBasedTriggerResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RequestsBasedTriggerResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o RequestsBasedTriggerResponseOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RequestsBasedTriggerResponse) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

type RequestsBasedTriggerResponsePtrOutput struct{ *pulumi.OutputState }

func (RequestsBasedTriggerResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestsBasedTriggerResponse)(nil)).Elem()
}

func (o RequestsBasedTriggerResponsePtrOutput) ToRequestsBasedTriggerResponsePtrOutput() RequestsBasedTriggerResponsePtrOutput {
	return o
}

func (o RequestsBasedTriggerResponsePtrOutput) ToRequestsBasedTriggerResponsePtrOutputWithContext(ctx context.Context) RequestsBasedTriggerResponsePtrOutput {
	return o
}

func (o RequestsBasedTriggerResponsePtrOutput) Elem() RequestsBasedTriggerResponseOutput {
	return o.ApplyT(func(v *RequestsBasedTriggerResponse) RequestsBasedTriggerResponse {
		if v != nil {
			return *v
		}
		var ret RequestsBasedTriggerResponse
		return ret
	}).(RequestsBasedTriggerResponseOutput)
}

func (o RequestsBasedTriggerResponsePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RequestsBasedTriggerResponse) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o RequestsBasedTriggerResponsePtrOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestsBasedTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeInterval
	}).(pulumi.StringPtrOutput)
}

type ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse struct {
	Error      *ErrorEntityResponse                     `pulumi:"error"`
	Id         *string                                  `pulumi:"id"`
	Identity   *ManagedServiceIdentityResponse          `pulumi:"identity"`
	Location   *string                                  `pulumi:"location"`
	Name       *string                                  `pulumi:"name"`
	Plan       *ArmPlanResponse                         `pulumi:"plan"`
	Properties *RemotePrivateEndpointConnectionResponse `pulumi:"properties"`
	Sku        *SkuDescriptionResponse                  `pulumi:"sku"`
	Status     *string                                  `pulumi:"status"`
	Tags       map[string]string                        `pulumi:"tags"`
	Type       *string                                  `pulumi:"type"`
	Zones      []string                                 `pulumi:"zones"`
}

type ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) ToResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput() ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput {
	return o
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) ToResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput {
	return o
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Error() ErrorEntityResponsePtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *ErrorEntityResponse {
		return v.Error
	}).(ErrorEntityResponsePtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *ManagedServiceIdentityResponse {
		return v.Identity
	}).(ManagedServiceIdentityResponsePtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Plan() ArmPlanResponsePtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *ArmPlanResponse { return v.Plan }).(ArmPlanResponsePtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Properties() RemotePrivateEndpointConnectionResponsePtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *RemotePrivateEndpointConnectionResponse {
		return v.Properties
	}).(RemotePrivateEndpointConnectionResponsePtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Sku() SkuDescriptionResponsePtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *SkuDescriptionResponse {
		return v.Sku
	}).(SkuDescriptionResponsePtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) map[string]string {
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

type ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput) ToResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput() ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput) ToResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse {
		return vs[0].([]ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput)
}

type SiteConfig struct {
	AcrUseManagedIdentityCreds             *bool                            `pulumi:"acrUseManagedIdentityCreds"`
	AcrUserManagedIdentityID               *string                          `pulumi:"acrUserManagedIdentityID"`
	AlwaysOn                               *bool                            `pulumi:"alwaysOn"`
	ApiDefinition                          *ApiDefinitionInfo               `pulumi:"apiDefinition"`
	ApiManagementConfig                    *ApiManagementConfig             `pulumi:"apiManagementConfig"`
	AppCommandLine                         *string                          `pulumi:"appCommandLine"`
	AppSettings                            []NameValuePair                  `pulumi:"appSettings"`
	AutoHealEnabled                        *bool                            `pulumi:"autoHealEnabled"`
	AutoHealRules                          *AutoHealRules                   `pulumi:"autoHealRules"`
	AutoSwapSlotName                       *string                          `pulumi:"autoSwapSlotName"`
	AzureStorageAccounts                   map[string]AzureStorageInfoValue `pulumi:"azureStorageAccounts"`
	ConnectionStrings                      []ConnStringInfo                 `pulumi:"connectionStrings"`
	Cors                                   *CorsSettings                    `pulumi:"cors"`
	DefaultDocuments                       []string                         `pulumi:"defaultDocuments"`
	DetailedErrorLoggingEnabled            *bool                            `pulumi:"detailedErrorLoggingEnabled"`
	DocumentRoot                           *string                          `pulumi:"documentRoot"`
	Experiments                            *Experiments                     `pulumi:"experiments"`
	FtpsState                              *string                          `pulumi:"ftpsState"`
	FunctionAppScaleLimit                  *int                             `pulumi:"functionAppScaleLimit"`
	FunctionsRuntimeScaleMonitoringEnabled *bool                            `pulumi:"functionsRuntimeScaleMonitoringEnabled"`
	HandlerMappings                        []HandlerMapping                 `pulumi:"handlerMappings"`
	HealthCheckPath                        *string                          `pulumi:"healthCheckPath"`
	Http20Enabled                          *bool                            `pulumi:"http20Enabled"`
	HttpLoggingEnabled                     *bool                            `pulumi:"httpLoggingEnabled"`
	IpSecurityRestrictions                 []IpSecurityRestriction          `pulumi:"ipSecurityRestrictions"`
	JavaContainer                          *string                          `pulumi:"javaContainer"`
	JavaContainerVersion                   *string                          `pulumi:"javaContainerVersion"`
	JavaVersion                            *string                          `pulumi:"javaVersion"`
	KeyVaultReferenceIdentity              *string                          `pulumi:"keyVaultReferenceIdentity"`
	Limits                                 *SiteLimits                      `pulumi:"limits"`
	LinuxFxVersion                         *string                          `pulumi:"linuxFxVersion"`
	LoadBalancing                          *SiteLoadBalancing               `pulumi:"loadBalancing"`
	LocalMySqlEnabled                      *bool                            `pulumi:"localMySqlEnabled"`
	LogsDirectorySizeLimit                 *int                             `pulumi:"logsDirectorySizeLimit"`
	ManagedPipelineMode                    *ManagedPipelineMode             `pulumi:"managedPipelineMode"`
	ManagedServiceIdentityId               *int                             `pulumi:"managedServiceIdentityId"`
	MinTlsVersion                          *string                          `pulumi:"minTlsVersion"`
	MinimumElasticInstanceCount            *int                             `pulumi:"minimumElasticInstanceCount"`
	NetFrameworkVersion                    *string                          `pulumi:"netFrameworkVersion"`
	NodeVersion                            *string                          `pulumi:"nodeVersion"`
	NumberOfWorkers                        *int                             `pulumi:"numberOfWorkers"`
	PhpVersion                             *string                          `pulumi:"phpVersion"`
	PowerShellVersion                      *string                          `pulumi:"powerShellVersion"`
	PreWarmedInstanceCount                 *int                             `pulumi:"preWarmedInstanceCount"`
	PublicNetworkAccess                    *string                          `pulumi:"publicNetworkAccess"`
	PublishingUsername                     *string                          `pulumi:"publishingUsername"`
	Push                                   *PushSettings                    `pulumi:"push"`
	PythonVersion                          *string                          `pulumi:"pythonVersion"`
	RemoteDebuggingEnabled                 *bool                            `pulumi:"remoteDebuggingEnabled"`
	RemoteDebuggingVersion                 *string                          `pulumi:"remoteDebuggingVersion"`
	RequestTracingEnabled                  *bool                            `pulumi:"requestTracingEnabled"`
	RequestTracingExpirationTime           *string                          `pulumi:"requestTracingExpirationTime"`
	ScmIpSecurityRestrictions              []IpSecurityRestriction          `pulumi:"scmIpSecurityRestrictions"`
	ScmIpSecurityRestrictionsUseMain       *bool                            `pulumi:"scmIpSecurityRestrictionsUseMain"`
	ScmMinTlsVersion                       *string                          `pulumi:"scmMinTlsVersion"`
	ScmType                                *string                          `pulumi:"scmType"`
	TracingOptions                         *string                          `pulumi:"tracingOptions"`
	Use32BitWorkerProcess                  *bool                            `pulumi:"use32BitWorkerProcess"`
	VirtualApplications                    []VirtualApplication             `pulumi:"virtualApplications"`
	VnetName                               *string                          `pulumi:"vnetName"`
	VnetPrivatePortsCount                  *int                             `pulumi:"vnetPrivatePortsCount"`
	VnetRouteAllEnabled                    *bool                            `pulumi:"vnetRouteAllEnabled"`
	WebSocketsEnabled                      *bool                            `pulumi:"webSocketsEnabled"`
	WebsiteTimeZone                        *string                          `pulumi:"websiteTimeZone"`
	WindowsFxVersion                       *string                          `pulumi:"windowsFxVersion"`
	XManagedServiceIdentityId              *int                             `pulumi:"xManagedServiceIdentityId"`
}


func (val *SiteConfig) Defaults() *SiteConfig {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Http20Enabled) {
		http20Enabled_ := true
		tmp.Http20Enabled = &http20Enabled_
	}
	if isZero(tmp.LocalMySqlEnabled) {
		localMySqlEnabled_ := false
		tmp.LocalMySqlEnabled = &localMySqlEnabled_
	}
	if isZero(tmp.NetFrameworkVersion) {
		netFrameworkVersion_ := "v4.6"
		tmp.NetFrameworkVersion = &netFrameworkVersion_
	}
	return &tmp
}





type SiteConfigInput interface {
	pulumi.Input

	ToSiteConfigOutput() SiteConfigOutput
	ToSiteConfigOutputWithContext(context.Context) SiteConfigOutput
}

type SiteConfigArgs struct {
	AcrUseManagedIdentityCreds             pulumi.BoolPtrInput             `pulumi:"acrUseManagedIdentityCreds"`
	AcrUserManagedIdentityID               pulumi.StringPtrInput           `pulumi:"acrUserManagedIdentityID"`
	AlwaysOn                               pulumi.BoolPtrInput             `pulumi:"alwaysOn"`
	ApiDefinition                          ApiDefinitionInfoPtrInput       `pulumi:"apiDefinition"`
	ApiManagementConfig                    ApiManagementConfigPtrInput     `pulumi:"apiManagementConfig"`
	AppCommandLine                         pulumi.StringPtrInput           `pulumi:"appCommandLine"`
	AppSettings                            NameValuePairArrayInput         `pulumi:"appSettings"`
	AutoHealEnabled                        pulumi.BoolPtrInput             `pulumi:"autoHealEnabled"`
	AutoHealRules                          AutoHealRulesPtrInput           `pulumi:"autoHealRules"`
	AutoSwapSlotName                       pulumi.StringPtrInput           `pulumi:"autoSwapSlotName"`
	AzureStorageAccounts                   AzureStorageInfoValueMapInput   `pulumi:"azureStorageAccounts"`
	ConnectionStrings                      ConnStringInfoArrayInput        `pulumi:"connectionStrings"`
	Cors                                   CorsSettingsPtrInput            `pulumi:"cors"`
	DefaultDocuments                       pulumi.StringArrayInput         `pulumi:"defaultDocuments"`
	DetailedErrorLoggingEnabled            pulumi.BoolPtrInput             `pulumi:"detailedErrorLoggingEnabled"`
	DocumentRoot                           pulumi.StringPtrInput           `pulumi:"documentRoot"`
	Experiments                            ExperimentsPtrInput             `pulumi:"experiments"`
	FtpsState                              pulumi.StringPtrInput           `pulumi:"ftpsState"`
	FunctionAppScaleLimit                  pulumi.IntPtrInput              `pulumi:"functionAppScaleLimit"`
	FunctionsRuntimeScaleMonitoringEnabled pulumi.BoolPtrInput             `pulumi:"functionsRuntimeScaleMonitoringEnabled"`
	HandlerMappings                        HandlerMappingArrayInput        `pulumi:"handlerMappings"`
	HealthCheckPath                        pulumi.StringPtrInput           `pulumi:"healthCheckPath"`
	Http20Enabled                          pulumi.BoolPtrInput             `pulumi:"http20Enabled"`
	HttpLoggingEnabled                     pulumi.BoolPtrInput             `pulumi:"httpLoggingEnabled"`
	IpSecurityRestrictions                 IpSecurityRestrictionArrayInput `pulumi:"ipSecurityRestrictions"`
	JavaContainer                          pulumi.StringPtrInput           `pulumi:"javaContainer"`
	JavaContainerVersion                   pulumi.StringPtrInput           `pulumi:"javaContainerVersion"`
	JavaVersion                            pulumi.StringPtrInput           `pulumi:"javaVersion"`
	KeyVaultReferenceIdentity              pulumi.StringPtrInput           `pulumi:"keyVaultReferenceIdentity"`
	Limits                                 SiteLimitsPtrInput              `pulumi:"limits"`
	LinuxFxVersion                         pulumi.StringPtrInput           `pulumi:"linuxFxVersion"`
	LoadBalancing                          SiteLoadBalancingPtrInput       `pulumi:"loadBalancing"`
	LocalMySqlEnabled                      pulumi.BoolPtrInput             `pulumi:"localMySqlEnabled"`
	LogsDirectorySizeLimit                 pulumi.IntPtrInput              `pulumi:"logsDirectorySizeLimit"`
	ManagedPipelineMode                    ManagedPipelineModePtrInput     `pulumi:"managedPipelineMode"`
	ManagedServiceIdentityId               pulumi.IntPtrInput              `pulumi:"managedServiceIdentityId"`
	MinTlsVersion                          pulumi.StringPtrInput           `pulumi:"minTlsVersion"`
	MinimumElasticInstanceCount            pulumi.IntPtrInput              `pulumi:"minimumElasticInstanceCount"`
	NetFrameworkVersion                    pulumi.StringPtrInput           `pulumi:"netFrameworkVersion"`
	NodeVersion                            pulumi.StringPtrInput           `pulumi:"nodeVersion"`
	NumberOfWorkers                        pulumi.IntPtrInput              `pulumi:"numberOfWorkers"`
	PhpVersion                             pulumi.StringPtrInput           `pulumi:"phpVersion"`
	PowerShellVersion                      pulumi.StringPtrInput           `pulumi:"powerShellVersion"`
	PreWarmedInstanceCount                 pulumi.IntPtrInput              `pulumi:"preWarmedInstanceCount"`
	PublicNetworkAccess                    pulumi.StringPtrInput           `pulumi:"publicNetworkAccess"`
	PublishingUsername                     pulumi.StringPtrInput           `pulumi:"publishingUsername"`
	Push                                   PushSettingsPtrInput            `pulumi:"push"`
	PythonVersion                          pulumi.StringPtrInput           `pulumi:"pythonVersion"`
	RemoteDebuggingEnabled                 pulumi.BoolPtrInput             `pulumi:"remoteDebuggingEnabled"`
	RemoteDebuggingVersion                 pulumi.StringPtrInput           `pulumi:"remoteDebuggingVersion"`
	RequestTracingEnabled                  pulumi.BoolPtrInput             `pulumi:"requestTracingEnabled"`
	RequestTracingExpirationTime           pulumi.StringPtrInput           `pulumi:"requestTracingExpirationTime"`
	ScmIpSecurityRestrictions              IpSecurityRestrictionArrayInput `pulumi:"scmIpSecurityRestrictions"`
	ScmIpSecurityRestrictionsUseMain       pulumi.BoolPtrInput             `pulumi:"scmIpSecurityRestrictionsUseMain"`
	ScmMinTlsVersion                       pulumi.StringPtrInput           `pulumi:"scmMinTlsVersion"`
	ScmType                                pulumi.StringPtrInput           `pulumi:"scmType"`
	TracingOptions                         pulumi.StringPtrInput           `pulumi:"tracingOptions"`
	Use32BitWorkerProcess                  pulumi.BoolPtrInput             `pulumi:"use32BitWorkerProcess"`
	VirtualApplications                    VirtualApplicationArrayInput    `pulumi:"virtualApplications"`
	VnetName                               pulumi.StringPtrInput           `pulumi:"vnetName"`
	VnetPrivatePortsCount                  pulumi.IntPtrInput              `pulumi:"vnetPrivatePortsCount"`
	VnetRouteAllEnabled                    pulumi.BoolPtrInput             `pulumi:"vnetRouteAllEnabled"`
	WebSocketsEnabled                      pulumi.BoolPtrInput             `pulumi:"webSocketsEnabled"`
	WebsiteTimeZone                        pulumi.StringPtrInput           `pulumi:"websiteTimeZone"`
	WindowsFxVersion                       pulumi.StringPtrInput           `pulumi:"windowsFxVersion"`
	XManagedServiceIdentityId              pulumi.IntPtrInput              `pulumi:"xManagedServiceIdentityId"`
}


func (val *SiteConfigArgs) Defaults() *SiteConfigArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Http20Enabled) {
		tmp.Http20Enabled = pulumi.BoolPtr(true)
	}
	if isZero(tmp.LocalMySqlEnabled) {
		tmp.LocalMySqlEnabled = pulumi.BoolPtr(false)
	}
	if isZero(tmp.NetFrameworkVersion) {
		tmp.NetFrameworkVersion = pulumi.StringPtr("v4.6")
	}
	return &tmp
}
func (SiteConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteConfig)(nil)).Elem()
}

func (i SiteConfigArgs) ToSiteConfigOutput() SiteConfigOutput {
	return i.ToSiteConfigOutputWithContext(context.Background())
}

func (i SiteConfigArgs) ToSiteConfigOutputWithContext(ctx context.Context) SiteConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConfigOutput)
}

func (i SiteConfigArgs) ToSiteConfigPtrOutput() SiteConfigPtrOutput {
	return i.ToSiteConfigPtrOutputWithContext(context.Background())
}

func (i SiteConfigArgs) ToSiteConfigPtrOutputWithContext(ctx context.Context) SiteConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConfigOutput).ToSiteConfigPtrOutputWithContext(ctx)
}









type SiteConfigPtrInput interface {
	pulumi.Input

	ToSiteConfigPtrOutput() SiteConfigPtrOutput
	ToSiteConfigPtrOutputWithContext(context.Context) SiteConfigPtrOutput
}

type siteConfigPtrType SiteConfigArgs

func SiteConfigPtr(v *SiteConfigArgs) SiteConfigPtrInput {
	return (*siteConfigPtrType)(v)
}

func (*siteConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteConfig)(nil)).Elem()
}

func (i *siteConfigPtrType) ToSiteConfigPtrOutput() SiteConfigPtrOutput {
	return i.ToSiteConfigPtrOutputWithContext(context.Background())
}

func (i *siteConfigPtrType) ToSiteConfigPtrOutputWithContext(ctx context.Context) SiteConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConfigPtrOutput)
}

type SiteConfigOutput struct{ *pulumi.OutputState }

func (SiteConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteConfig)(nil)).Elem()
}

func (o SiteConfigOutput) ToSiteConfigOutput() SiteConfigOutput {
	return o
}

func (o SiteConfigOutput) ToSiteConfigOutputWithContext(ctx context.Context) SiteConfigOutput {
	return o
}

func (o SiteConfigOutput) ToSiteConfigPtrOutput() SiteConfigPtrOutput {
	return o.ToSiteConfigPtrOutputWithContext(context.Background())
}

func (o SiteConfigOutput) ToSiteConfigPtrOutputWithContext(ctx context.Context) SiteConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteConfig) *SiteConfig {
		return &v
	}).(SiteConfigPtrOutput)
}

func (o SiteConfigOutput) AcrUseManagedIdentityCreds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.AcrUseManagedIdentityCreds }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) AcrUserManagedIdentityID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.AcrUserManagedIdentityID }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) AlwaysOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.AlwaysOn }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) ApiDefinition() ApiDefinitionInfoPtrOutput {
	return o.ApplyT(func(v SiteConfig) *ApiDefinitionInfo { return v.ApiDefinition }).(ApiDefinitionInfoPtrOutput)
}

func (o SiteConfigOutput) ApiManagementConfig() ApiManagementConfigPtrOutput {
	return o.ApplyT(func(v SiteConfig) *ApiManagementConfig { return v.ApiManagementConfig }).(ApiManagementConfigPtrOutput)
}

func (o SiteConfigOutput) AppCommandLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.AppCommandLine }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) AppSettings() NameValuePairArrayOutput {
	return o.ApplyT(func(v SiteConfig) []NameValuePair { return v.AppSettings }).(NameValuePairArrayOutput)
}

func (o SiteConfigOutput) AutoHealEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.AutoHealEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) AutoHealRules() AutoHealRulesPtrOutput {
	return o.ApplyT(func(v SiteConfig) *AutoHealRules { return v.AutoHealRules }).(AutoHealRulesPtrOutput)
}

func (o SiteConfigOutput) AutoSwapSlotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.AutoSwapSlotName }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) AzureStorageAccounts() AzureStorageInfoValueMapOutput {
	return o.ApplyT(func(v SiteConfig) map[string]AzureStorageInfoValue { return v.AzureStorageAccounts }).(AzureStorageInfoValueMapOutput)
}

func (o SiteConfigOutput) ConnectionStrings() ConnStringInfoArrayOutput {
	return o.ApplyT(func(v SiteConfig) []ConnStringInfo { return v.ConnectionStrings }).(ConnStringInfoArrayOutput)
}

func (o SiteConfigOutput) Cors() CorsSettingsPtrOutput {
	return o.ApplyT(func(v SiteConfig) *CorsSettings { return v.Cors }).(CorsSettingsPtrOutput)
}

func (o SiteConfigOutput) DefaultDocuments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SiteConfig) []string { return v.DefaultDocuments }).(pulumi.StringArrayOutput)
}

func (o SiteConfigOutput) DetailedErrorLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.DetailedErrorLoggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) DocumentRoot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.DocumentRoot }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) Experiments() ExperimentsPtrOutput {
	return o.ApplyT(func(v SiteConfig) *Experiments { return v.Experiments }).(ExperimentsPtrOutput)
}

func (o SiteConfigOutput) FtpsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.FtpsState }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) FunctionAppScaleLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfig) *int { return v.FunctionAppScaleLimit }).(pulumi.IntPtrOutput)
}

func (o SiteConfigOutput) FunctionsRuntimeScaleMonitoringEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.FunctionsRuntimeScaleMonitoringEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) HandlerMappings() HandlerMappingArrayOutput {
	return o.ApplyT(func(v SiteConfig) []HandlerMapping { return v.HandlerMappings }).(HandlerMappingArrayOutput)
}

func (o SiteConfigOutput) HealthCheckPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.HealthCheckPath }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) Http20Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.Http20Enabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) HttpLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.HttpLoggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) IpSecurityRestrictions() IpSecurityRestrictionArrayOutput {
	return o.ApplyT(func(v SiteConfig) []IpSecurityRestriction { return v.IpSecurityRestrictions }).(IpSecurityRestrictionArrayOutput)
}

func (o SiteConfigOutput) JavaContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.JavaContainer }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) JavaContainerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.JavaContainerVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) JavaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.JavaVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) KeyVaultReferenceIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.KeyVaultReferenceIdentity }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) Limits() SiteLimitsPtrOutput {
	return o.ApplyT(func(v SiteConfig) *SiteLimits { return v.Limits }).(SiteLimitsPtrOutput)
}

func (o SiteConfigOutput) LinuxFxVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.LinuxFxVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) LoadBalancing() SiteLoadBalancingPtrOutput {
	return o.ApplyT(func(v SiteConfig) *SiteLoadBalancing { return v.LoadBalancing }).(SiteLoadBalancingPtrOutput)
}

func (o SiteConfigOutput) LocalMySqlEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.LocalMySqlEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) LogsDirectorySizeLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfig) *int { return v.LogsDirectorySizeLimit }).(pulumi.IntPtrOutput)
}

func (o SiteConfigOutput) ManagedPipelineMode() ManagedPipelineModePtrOutput {
	return o.ApplyT(func(v SiteConfig) *ManagedPipelineMode { return v.ManagedPipelineMode }).(ManagedPipelineModePtrOutput)
}

func (o SiteConfigOutput) ManagedServiceIdentityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfig) *int { return v.ManagedServiceIdentityId }).(pulumi.IntPtrOutput)
}

func (o SiteConfigOutput) MinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.MinTlsVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) MinimumElasticInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfig) *int { return v.MinimumElasticInstanceCount }).(pulumi.IntPtrOutput)
}

func (o SiteConfigOutput) NetFrameworkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.NetFrameworkVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) NodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.NodeVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfig) *int { return v.NumberOfWorkers }).(pulumi.IntPtrOutput)
}

func (o SiteConfigOutput) PhpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.PhpVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) PowerShellVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.PowerShellVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) PreWarmedInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfig) *int { return v.PreWarmedInstanceCount }).(pulumi.IntPtrOutput)
}

func (o SiteConfigOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) PublishingUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.PublishingUsername }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) Push() PushSettingsPtrOutput {
	return o.ApplyT(func(v SiteConfig) *PushSettings { return v.Push }).(PushSettingsPtrOutput)
}

func (o SiteConfigOutput) PythonVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.PythonVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) RemoteDebuggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.RemoteDebuggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) RemoteDebuggingVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.RemoteDebuggingVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) RequestTracingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.RequestTracingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) RequestTracingExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.RequestTracingExpirationTime }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) ScmIpSecurityRestrictions() IpSecurityRestrictionArrayOutput {
	return o.ApplyT(func(v SiteConfig) []IpSecurityRestriction { return v.ScmIpSecurityRestrictions }).(IpSecurityRestrictionArrayOutput)
}

func (o SiteConfigOutput) ScmIpSecurityRestrictionsUseMain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.ScmIpSecurityRestrictionsUseMain }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) ScmMinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.ScmMinTlsVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) ScmType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.ScmType }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) TracingOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.TracingOptions }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) Use32BitWorkerProcess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.Use32BitWorkerProcess }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) VirtualApplications() VirtualApplicationArrayOutput {
	return o.ApplyT(func(v SiteConfig) []VirtualApplication { return v.VirtualApplications }).(VirtualApplicationArrayOutput)
}

func (o SiteConfigOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.VnetName }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) VnetPrivatePortsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfig) *int { return v.VnetPrivatePortsCount }).(pulumi.IntPtrOutput)
}

func (o SiteConfigOutput) VnetRouteAllEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.VnetRouteAllEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) WebSocketsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.WebSocketsEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) WebsiteTimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.WebsiteTimeZone }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) WindowsFxVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.WindowsFxVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) XManagedServiceIdentityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfig) *int { return v.XManagedServiceIdentityId }).(pulumi.IntPtrOutput)
}

type SiteConfigPtrOutput struct{ *pulumi.OutputState }

func (SiteConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteConfig)(nil)).Elem()
}

func (o SiteConfigPtrOutput) ToSiteConfigPtrOutput() SiteConfigPtrOutput {
	return o
}

func (o SiteConfigPtrOutput) ToSiteConfigPtrOutputWithContext(ctx context.Context) SiteConfigPtrOutput {
	return o
}

func (o SiteConfigPtrOutput) Elem() SiteConfigOutput {
	return o.ApplyT(func(v *SiteConfig) SiteConfig {
		if v != nil {
			return *v
		}
		var ret SiteConfig
		return ret
	}).(SiteConfigOutput)
}

func (o SiteConfigPtrOutput) AcrUseManagedIdentityCreds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.AcrUseManagedIdentityCreds
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) AcrUserManagedIdentityID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.AcrUserManagedIdentityID
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) AlwaysOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.AlwaysOn
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) ApiDefinition() ApiDefinitionInfoPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *ApiDefinitionInfo {
		if v == nil {
			return nil
		}
		return v.ApiDefinition
	}).(ApiDefinitionInfoPtrOutput)
}

func (o SiteConfigPtrOutput) ApiManagementConfig() ApiManagementConfigPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *ApiManagementConfig {
		if v == nil {
			return nil
		}
		return v.ApiManagementConfig
	}).(ApiManagementConfigPtrOutput)
}

func (o SiteConfigPtrOutput) AppCommandLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.AppCommandLine
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) AppSettings() NameValuePairArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []NameValuePair {
		if v == nil {
			return nil
		}
		return v.AppSettings
	}).(NameValuePairArrayOutput)
}

func (o SiteConfigPtrOutput) AutoHealEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.AutoHealEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) AutoHealRules() AutoHealRulesPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *AutoHealRules {
		if v == nil {
			return nil
		}
		return v.AutoHealRules
	}).(AutoHealRulesPtrOutput)
}

func (o SiteConfigPtrOutput) AutoSwapSlotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.AutoSwapSlotName
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) AzureStorageAccounts() AzureStorageInfoValueMapOutput {
	return o.ApplyT(func(v *SiteConfig) map[string]AzureStorageInfoValue {
		if v == nil {
			return nil
		}
		return v.AzureStorageAccounts
	}).(AzureStorageInfoValueMapOutput)
}

func (o SiteConfigPtrOutput) ConnectionStrings() ConnStringInfoArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []ConnStringInfo {
		if v == nil {
			return nil
		}
		return v.ConnectionStrings
	}).(ConnStringInfoArrayOutput)
}

func (o SiteConfigPtrOutput) Cors() CorsSettingsPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *CorsSettings {
		if v == nil {
			return nil
		}
		return v.Cors
	}).(CorsSettingsPtrOutput)
}

func (o SiteConfigPtrOutput) DefaultDocuments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []string {
		if v == nil {
			return nil
		}
		return v.DefaultDocuments
	}).(pulumi.StringArrayOutput)
}

func (o SiteConfigPtrOutput) DetailedErrorLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.DetailedErrorLoggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) DocumentRoot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.DocumentRoot
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) Experiments() ExperimentsPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *Experiments {
		if v == nil {
			return nil
		}
		return v.Experiments
	}).(ExperimentsPtrOutput)
}

func (o SiteConfigPtrOutput) FtpsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.FtpsState
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) FunctionAppScaleLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *int {
		if v == nil {
			return nil
		}
		return v.FunctionAppScaleLimit
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigPtrOutput) FunctionsRuntimeScaleMonitoringEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.FunctionsRuntimeScaleMonitoringEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) HandlerMappings() HandlerMappingArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []HandlerMapping {
		if v == nil {
			return nil
		}
		return v.HandlerMappings
	}).(HandlerMappingArrayOutput)
}

func (o SiteConfigPtrOutput) HealthCheckPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.HealthCheckPath
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) Http20Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.Http20Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) HttpLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.HttpLoggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) IpSecurityRestrictions() IpSecurityRestrictionArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []IpSecurityRestriction {
		if v == nil {
			return nil
		}
		return v.IpSecurityRestrictions
	}).(IpSecurityRestrictionArrayOutput)
}

func (o SiteConfigPtrOutput) JavaContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.JavaContainer
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) JavaContainerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.JavaContainerVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) JavaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.JavaVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) KeyVaultReferenceIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultReferenceIdentity
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) Limits() SiteLimitsPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *SiteLimits {
		if v == nil {
			return nil
		}
		return v.Limits
	}).(SiteLimitsPtrOutput)
}

func (o SiteConfigPtrOutput) LinuxFxVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.LinuxFxVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) LoadBalancing() SiteLoadBalancingPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *SiteLoadBalancing {
		if v == nil {
			return nil
		}
		return v.LoadBalancing
	}).(SiteLoadBalancingPtrOutput)
}

func (o SiteConfigPtrOutput) LocalMySqlEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.LocalMySqlEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) LogsDirectorySizeLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *int {
		if v == nil {
			return nil
		}
		return v.LogsDirectorySizeLimit
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigPtrOutput) ManagedPipelineMode() ManagedPipelineModePtrOutput {
	return o.ApplyT(func(v *SiteConfig) *ManagedPipelineMode {
		if v == nil {
			return nil
		}
		return v.ManagedPipelineMode
	}).(ManagedPipelineModePtrOutput)
}

func (o SiteConfigPtrOutput) ManagedServiceIdentityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *int {
		if v == nil {
			return nil
		}
		return v.ManagedServiceIdentityId
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigPtrOutput) MinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.MinTlsVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) MinimumElasticInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *int {
		if v == nil {
			return nil
		}
		return v.MinimumElasticInstanceCount
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigPtrOutput) NetFrameworkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.NetFrameworkVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) NodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.NodeVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *int {
		if v == nil {
			return nil
		}
		return v.NumberOfWorkers
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigPtrOutput) PhpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.PhpVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) PowerShellVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.PowerShellVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) PreWarmedInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *int {
		if v == nil {
			return nil
		}
		return v.PreWarmedInstanceCount
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) PublishingUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.PublishingUsername
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) Push() PushSettingsPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *PushSettings {
		if v == nil {
			return nil
		}
		return v.Push
	}).(PushSettingsPtrOutput)
}

func (o SiteConfigPtrOutput) PythonVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.PythonVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) RemoteDebuggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.RemoteDebuggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) RemoteDebuggingVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.RemoteDebuggingVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) RequestTracingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.RequestTracingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) RequestTracingExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.RequestTracingExpirationTime
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) ScmIpSecurityRestrictions() IpSecurityRestrictionArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []IpSecurityRestriction {
		if v == nil {
			return nil
		}
		return v.ScmIpSecurityRestrictions
	}).(IpSecurityRestrictionArrayOutput)
}

func (o SiteConfigPtrOutput) ScmIpSecurityRestrictionsUseMain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.ScmIpSecurityRestrictionsUseMain
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) ScmMinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.ScmMinTlsVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) ScmType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.ScmType
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) TracingOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.TracingOptions
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) Use32BitWorkerProcess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.Use32BitWorkerProcess
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) VirtualApplications() VirtualApplicationArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []VirtualApplication {
		if v == nil {
			return nil
		}
		return v.VirtualApplications
	}).(VirtualApplicationArrayOutput)
}

func (o SiteConfigPtrOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.VnetName
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) VnetPrivatePortsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *int {
		if v == nil {
			return nil
		}
		return v.VnetPrivatePortsCount
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigPtrOutput) VnetRouteAllEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.VnetRouteAllEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) WebSocketsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.WebSocketsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) WebsiteTimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.WebsiteTimeZone
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) WindowsFxVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.WindowsFxVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) XManagedServiceIdentityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *int {
		if v == nil {
			return nil
		}
		return v.XManagedServiceIdentityId
	}).(pulumi.IntPtrOutput)
}

type SiteConfigResponse struct {
	AcrUseManagedIdentityCreds             *bool                                    `pulumi:"acrUseManagedIdentityCreds"`
	AcrUserManagedIdentityID               *string                                  `pulumi:"acrUserManagedIdentityID"`
	AlwaysOn                               *bool                                    `pulumi:"alwaysOn"`
	ApiDefinition                          *ApiDefinitionInfoResponse               `pulumi:"apiDefinition"`
	ApiManagementConfig                    *ApiManagementConfigResponse             `pulumi:"apiManagementConfig"`
	AppCommandLine                         *string                                  `pulumi:"appCommandLine"`
	AppSettings                            []NameValuePairResponse                  `pulumi:"appSettings"`
	AutoHealEnabled                        *bool                                    `pulumi:"autoHealEnabled"`
	AutoHealRules                          *AutoHealRulesResponse                   `pulumi:"autoHealRules"`
	AutoSwapSlotName                       *string                                  `pulumi:"autoSwapSlotName"`
	AzureStorageAccounts                   map[string]AzureStorageInfoValueResponse `pulumi:"azureStorageAccounts"`
	ConnectionStrings                      []ConnStringInfoResponse                 `pulumi:"connectionStrings"`
	Cors                                   *CorsSettingsResponse                    `pulumi:"cors"`
	DefaultDocuments                       []string                                 `pulumi:"defaultDocuments"`
	DetailedErrorLoggingEnabled            *bool                                    `pulumi:"detailedErrorLoggingEnabled"`
	DocumentRoot                           *string                                  `pulumi:"documentRoot"`
	Experiments                            *ExperimentsResponse                     `pulumi:"experiments"`
	FtpsState                              *string                                  `pulumi:"ftpsState"`
	FunctionAppScaleLimit                  *int                                     `pulumi:"functionAppScaleLimit"`
	FunctionsRuntimeScaleMonitoringEnabled *bool                                    `pulumi:"functionsRuntimeScaleMonitoringEnabled"`
	HandlerMappings                        []HandlerMappingResponse                 `pulumi:"handlerMappings"`
	HealthCheckPath                        *string                                  `pulumi:"healthCheckPath"`
	Http20Enabled                          *bool                                    `pulumi:"http20Enabled"`
	HttpLoggingEnabled                     *bool                                    `pulumi:"httpLoggingEnabled"`
	IpSecurityRestrictions                 []IpSecurityRestrictionResponse          `pulumi:"ipSecurityRestrictions"`
	JavaContainer                          *string                                  `pulumi:"javaContainer"`
	JavaContainerVersion                   *string                                  `pulumi:"javaContainerVersion"`
	JavaVersion                            *string                                  `pulumi:"javaVersion"`
	KeyVaultReferenceIdentity              *string                                  `pulumi:"keyVaultReferenceIdentity"`
	Limits                                 *SiteLimitsResponse                      `pulumi:"limits"`
	LinuxFxVersion                         *string                                  `pulumi:"linuxFxVersion"`
	LoadBalancing                          *string                                  `pulumi:"loadBalancing"`
	LocalMySqlEnabled                      *bool                                    `pulumi:"localMySqlEnabled"`
	LogsDirectorySizeLimit                 *int                                     `pulumi:"logsDirectorySizeLimit"`
	MachineKey                             SiteMachineKeyResponse                   `pulumi:"machineKey"`
	ManagedPipelineMode                    *string                                  `pulumi:"managedPipelineMode"`
	ManagedServiceIdentityId               *int                                     `pulumi:"managedServiceIdentityId"`
	MinTlsVersion                          *string                                  `pulumi:"minTlsVersion"`
	MinimumElasticInstanceCount            *int                                     `pulumi:"minimumElasticInstanceCount"`
	NetFrameworkVersion                    *string                                  `pulumi:"netFrameworkVersion"`
	NodeVersion                            *string                                  `pulumi:"nodeVersion"`
	NumberOfWorkers                        *int                                     `pulumi:"numberOfWorkers"`
	PhpVersion                             *string                                  `pulumi:"phpVersion"`
	PowerShellVersion                      *string                                  `pulumi:"powerShellVersion"`
	PreWarmedInstanceCount                 *int                                     `pulumi:"preWarmedInstanceCount"`
	PublicNetworkAccess                    *string                                  `pulumi:"publicNetworkAccess"`
	PublishingUsername                     *string                                  `pulumi:"publishingUsername"`
	Push                                   *PushSettingsResponse                    `pulumi:"push"`
	PythonVersion                          *string                                  `pulumi:"pythonVersion"`
	RemoteDebuggingEnabled                 *bool                                    `pulumi:"remoteDebuggingEnabled"`
	RemoteDebuggingVersion                 *string                                  `pulumi:"remoteDebuggingVersion"`
	RequestTracingEnabled                  *bool                                    `pulumi:"requestTracingEnabled"`
	RequestTracingExpirationTime           *string                                  `pulumi:"requestTracingExpirationTime"`
	ScmIpSecurityRestrictions              []IpSecurityRestrictionResponse          `pulumi:"scmIpSecurityRestrictions"`
	ScmIpSecurityRestrictionsUseMain       *bool                                    `pulumi:"scmIpSecurityRestrictionsUseMain"`
	ScmMinTlsVersion                       *string                                  `pulumi:"scmMinTlsVersion"`
	ScmType                                *string                                  `pulumi:"scmType"`
	TracingOptions                         *string                                  `pulumi:"tracingOptions"`
	Use32BitWorkerProcess                  *bool                                    `pulumi:"use32BitWorkerProcess"`
	VirtualApplications                    []VirtualApplicationResponse             `pulumi:"virtualApplications"`
	VnetName                               *string                                  `pulumi:"vnetName"`
	VnetPrivatePortsCount                  *int                                     `pulumi:"vnetPrivatePortsCount"`
	VnetRouteAllEnabled                    *bool                                    `pulumi:"vnetRouteAllEnabled"`
	WebSocketsEnabled                      *bool                                    `pulumi:"webSocketsEnabled"`
	WebsiteTimeZone                        *string                                  `pulumi:"websiteTimeZone"`
	WindowsFxVersion                       *string                                  `pulumi:"windowsFxVersion"`
	XManagedServiceIdentityId              *int                                     `pulumi:"xManagedServiceIdentityId"`
}


func (val *SiteConfigResponse) Defaults() *SiteConfigResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Http20Enabled) {
		http20Enabled_ := true
		tmp.Http20Enabled = &http20Enabled_
	}
	if isZero(tmp.LocalMySqlEnabled) {
		localMySqlEnabled_ := false
		tmp.LocalMySqlEnabled = &localMySqlEnabled_
	}
	if isZero(tmp.NetFrameworkVersion) {
		netFrameworkVersion_ := "v4.6"
		tmp.NetFrameworkVersion = &netFrameworkVersion_
	}
	return &tmp
}

type SiteConfigResponseOutput struct{ *pulumi.OutputState }

func (SiteConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteConfigResponse)(nil)).Elem()
}

func (o SiteConfigResponseOutput) ToSiteConfigResponseOutput() SiteConfigResponseOutput {
	return o
}

func (o SiteConfigResponseOutput) ToSiteConfigResponseOutputWithContext(ctx context.Context) SiteConfigResponseOutput {
	return o
}

func (o SiteConfigResponseOutput) AcrUseManagedIdentityCreds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.AcrUseManagedIdentityCreds }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) AcrUserManagedIdentityID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.AcrUserManagedIdentityID }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) AlwaysOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.AlwaysOn }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) ApiDefinition() ApiDefinitionInfoResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *ApiDefinitionInfoResponse { return v.ApiDefinition }).(ApiDefinitionInfoResponsePtrOutput)
}

func (o SiteConfigResponseOutput) ApiManagementConfig() ApiManagementConfigResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *ApiManagementConfigResponse { return v.ApiManagementConfig }).(ApiManagementConfigResponsePtrOutput)
}

func (o SiteConfigResponseOutput) AppCommandLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.AppCommandLine }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) AppSettings() NameValuePairResponseArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []NameValuePairResponse { return v.AppSettings }).(NameValuePairResponseArrayOutput)
}

func (o SiteConfigResponseOutput) AutoHealEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.AutoHealEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) AutoHealRules() AutoHealRulesResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *AutoHealRulesResponse { return v.AutoHealRules }).(AutoHealRulesResponsePtrOutput)
}

func (o SiteConfigResponseOutput) AutoSwapSlotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.AutoSwapSlotName }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) AzureStorageAccounts() AzureStorageInfoValueResponseMapOutput {
	return o.ApplyT(func(v SiteConfigResponse) map[string]AzureStorageInfoValueResponse { return v.AzureStorageAccounts }).(AzureStorageInfoValueResponseMapOutput)
}

func (o SiteConfigResponseOutput) ConnectionStrings() ConnStringInfoResponseArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []ConnStringInfoResponse { return v.ConnectionStrings }).(ConnStringInfoResponseArrayOutput)
}

func (o SiteConfigResponseOutput) Cors() CorsSettingsResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *CorsSettingsResponse { return v.Cors }).(CorsSettingsResponsePtrOutput)
}

func (o SiteConfigResponseOutput) DefaultDocuments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []string { return v.DefaultDocuments }).(pulumi.StringArrayOutput)
}

func (o SiteConfigResponseOutput) DetailedErrorLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.DetailedErrorLoggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) DocumentRoot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.DocumentRoot }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) Experiments() ExperimentsResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *ExperimentsResponse { return v.Experiments }).(ExperimentsResponsePtrOutput)
}

func (o SiteConfigResponseOutput) FtpsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.FtpsState }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) FunctionAppScaleLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *int { return v.FunctionAppScaleLimit }).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponseOutput) FunctionsRuntimeScaleMonitoringEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.FunctionsRuntimeScaleMonitoringEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) HandlerMappings() HandlerMappingResponseArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []HandlerMappingResponse { return v.HandlerMappings }).(HandlerMappingResponseArrayOutput)
}

func (o SiteConfigResponseOutput) HealthCheckPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.HealthCheckPath }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) Http20Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.Http20Enabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) HttpLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.HttpLoggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) IpSecurityRestrictions() IpSecurityRestrictionResponseArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []IpSecurityRestrictionResponse { return v.IpSecurityRestrictions }).(IpSecurityRestrictionResponseArrayOutput)
}

func (o SiteConfigResponseOutput) JavaContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.JavaContainer }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) JavaContainerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.JavaContainerVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) JavaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.JavaVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) KeyVaultReferenceIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.KeyVaultReferenceIdentity }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) Limits() SiteLimitsResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *SiteLimitsResponse { return v.Limits }).(SiteLimitsResponsePtrOutput)
}

func (o SiteConfigResponseOutput) LinuxFxVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.LinuxFxVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) LoadBalancing() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.LoadBalancing }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) LocalMySqlEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.LocalMySqlEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) LogsDirectorySizeLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *int { return v.LogsDirectorySizeLimit }).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponseOutput) MachineKey() SiteMachineKeyResponseOutput {
	return o.ApplyT(func(v SiteConfigResponse) SiteMachineKeyResponse { return v.MachineKey }).(SiteMachineKeyResponseOutput)
}

func (o SiteConfigResponseOutput) ManagedPipelineMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.ManagedPipelineMode }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) ManagedServiceIdentityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *int { return v.ManagedServiceIdentityId }).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponseOutput) MinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.MinTlsVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) MinimumElasticInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *int { return v.MinimumElasticInstanceCount }).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponseOutput) NetFrameworkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.NetFrameworkVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) NodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.NodeVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *int { return v.NumberOfWorkers }).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponseOutput) PhpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.PhpVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) PowerShellVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.PowerShellVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) PreWarmedInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *int { return v.PreWarmedInstanceCount }).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) PublishingUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.PublishingUsername }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) Push() PushSettingsResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *PushSettingsResponse { return v.Push }).(PushSettingsResponsePtrOutput)
}

func (o SiteConfigResponseOutput) PythonVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.PythonVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) RemoteDebuggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.RemoteDebuggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) RemoteDebuggingVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.RemoteDebuggingVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) RequestTracingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.RequestTracingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) RequestTracingExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.RequestTracingExpirationTime }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) ScmIpSecurityRestrictions() IpSecurityRestrictionResponseArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []IpSecurityRestrictionResponse { return v.ScmIpSecurityRestrictions }).(IpSecurityRestrictionResponseArrayOutput)
}

func (o SiteConfigResponseOutput) ScmIpSecurityRestrictionsUseMain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.ScmIpSecurityRestrictionsUseMain }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) ScmMinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.ScmMinTlsVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) ScmType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.ScmType }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) TracingOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.TracingOptions }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) Use32BitWorkerProcess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.Use32BitWorkerProcess }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) VirtualApplications() VirtualApplicationResponseArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []VirtualApplicationResponse { return v.VirtualApplications }).(VirtualApplicationResponseArrayOutput)
}

func (o SiteConfigResponseOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.VnetName }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) VnetPrivatePortsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *int { return v.VnetPrivatePortsCount }).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponseOutput) VnetRouteAllEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.VnetRouteAllEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) WebSocketsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.WebSocketsEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) WebsiteTimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.WebsiteTimeZone }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) WindowsFxVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.WindowsFxVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) XManagedServiceIdentityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *int { return v.XManagedServiceIdentityId }).(pulumi.IntPtrOutput)
}

type SiteConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (SiteConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteConfigResponse)(nil)).Elem()
}

func (o SiteConfigResponsePtrOutput) ToSiteConfigResponsePtrOutput() SiteConfigResponsePtrOutput {
	return o
}

func (o SiteConfigResponsePtrOutput) ToSiteConfigResponsePtrOutputWithContext(ctx context.Context) SiteConfigResponsePtrOutput {
	return o
}

func (o SiteConfigResponsePtrOutput) Elem() SiteConfigResponseOutput {
	return o.ApplyT(func(v *SiteConfigResponse) SiteConfigResponse {
		if v != nil {
			return *v
		}
		var ret SiteConfigResponse
		return ret
	}).(SiteConfigResponseOutput)
}

func (o SiteConfigResponsePtrOutput) AcrUseManagedIdentityCreds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AcrUseManagedIdentityCreds
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) AcrUserManagedIdentityID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.AcrUserManagedIdentityID
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) AlwaysOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AlwaysOn
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) ApiDefinition() ApiDefinitionInfoResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *ApiDefinitionInfoResponse {
		if v == nil {
			return nil
		}
		return v.ApiDefinition
	}).(ApiDefinitionInfoResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) ApiManagementConfig() ApiManagementConfigResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *ApiManagementConfigResponse {
		if v == nil {
			return nil
		}
		return v.ApiManagementConfig
	}).(ApiManagementConfigResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) AppCommandLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppCommandLine
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) AppSettings() NameValuePairResponseArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []NameValuePairResponse {
		if v == nil {
			return nil
		}
		return v.AppSettings
	}).(NameValuePairResponseArrayOutput)
}

func (o SiteConfigResponsePtrOutput) AutoHealEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AutoHealEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) AutoHealRules() AutoHealRulesResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *AutoHealRulesResponse {
		if v == nil {
			return nil
		}
		return v.AutoHealRules
	}).(AutoHealRulesResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) AutoSwapSlotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.AutoSwapSlotName
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) AzureStorageAccounts() AzureStorageInfoValueResponseMapOutput {
	return o.ApplyT(func(v *SiteConfigResponse) map[string]AzureStorageInfoValueResponse {
		if v == nil {
			return nil
		}
		return v.AzureStorageAccounts
	}).(AzureStorageInfoValueResponseMapOutput)
}

func (o SiteConfigResponsePtrOutput) ConnectionStrings() ConnStringInfoResponseArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []ConnStringInfoResponse {
		if v == nil {
			return nil
		}
		return v.ConnectionStrings
	}).(ConnStringInfoResponseArrayOutput)
}

func (o SiteConfigResponsePtrOutput) Cors() CorsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *CorsSettingsResponse {
		if v == nil {
			return nil
		}
		return v.Cors
	}).(CorsSettingsResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) DefaultDocuments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []string {
		if v == nil {
			return nil
		}
		return v.DefaultDocuments
	}).(pulumi.StringArrayOutput)
}

func (o SiteConfigResponsePtrOutput) DetailedErrorLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DetailedErrorLoggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) DocumentRoot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.DocumentRoot
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Experiments() ExperimentsResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *ExperimentsResponse {
		if v == nil {
			return nil
		}
		return v.Experiments
	}).(ExperimentsResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) FtpsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.FtpsState
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) FunctionAppScaleLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.FunctionAppScaleLimit
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponsePtrOutput) FunctionsRuntimeScaleMonitoringEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.FunctionsRuntimeScaleMonitoringEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) HandlerMappings() HandlerMappingResponseArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []HandlerMappingResponse {
		if v == nil {
			return nil
		}
		return v.HandlerMappings
	}).(HandlerMappingResponseArrayOutput)
}

func (o SiteConfigResponsePtrOutput) HealthCheckPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.HealthCheckPath
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Http20Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Http20Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) HttpLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.HttpLoggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) IpSecurityRestrictions() IpSecurityRestrictionResponseArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []IpSecurityRestrictionResponse {
		if v == nil {
			return nil
		}
		return v.IpSecurityRestrictions
	}).(IpSecurityRestrictionResponseArrayOutput)
}

func (o SiteConfigResponsePtrOutput) JavaContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.JavaContainer
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) JavaContainerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.JavaContainerVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) JavaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.JavaVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) KeyVaultReferenceIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultReferenceIdentity
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Limits() SiteLimitsResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *SiteLimitsResponse {
		if v == nil {
			return nil
		}
		return v.Limits
	}).(SiteLimitsResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) LinuxFxVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.LinuxFxVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) LoadBalancing() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancing
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) LocalMySqlEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.LocalMySqlEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) LogsDirectorySizeLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.LogsDirectorySizeLimit
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponsePtrOutput) MachineKey() SiteMachineKeyResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *SiteMachineKeyResponse {
		if v == nil {
			return nil
		}
		return &v.MachineKey
	}).(SiteMachineKeyResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) ManagedPipelineMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.ManagedPipelineMode
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) ManagedServiceIdentityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.ManagedServiceIdentityId
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponsePtrOutput) MinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.MinTlsVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) MinimumElasticInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.MinimumElasticInstanceCount
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponsePtrOutput) NetFrameworkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.NetFrameworkVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) NodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.NodeVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NumberOfWorkers
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponsePtrOutput) PhpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.PhpVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) PowerShellVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.PowerShellVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) PreWarmedInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.PreWarmedInstanceCount
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponsePtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) PublishingUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublishingUsername
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Push() PushSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *PushSettingsResponse {
		if v == nil {
			return nil
		}
		return v.Push
	}).(PushSettingsResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) PythonVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.PythonVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) RemoteDebuggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RemoteDebuggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) RemoteDebuggingVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.RemoteDebuggingVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) RequestTracingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequestTracingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) RequestTracingExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.RequestTracingExpirationTime
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) ScmIpSecurityRestrictions() IpSecurityRestrictionResponseArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []IpSecurityRestrictionResponse {
		if v == nil {
			return nil
		}
		return v.ScmIpSecurityRestrictions
	}).(IpSecurityRestrictionResponseArrayOutput)
}

func (o SiteConfigResponsePtrOutput) ScmIpSecurityRestrictionsUseMain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ScmIpSecurityRestrictionsUseMain
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) ScmMinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScmMinTlsVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) ScmType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScmType
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) TracingOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.TracingOptions
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Use32BitWorkerProcess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Use32BitWorkerProcess
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) VirtualApplications() VirtualApplicationResponseArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []VirtualApplicationResponse {
		if v == nil {
			return nil
		}
		return v.VirtualApplications
	}).(VirtualApplicationResponseArrayOutput)
}

func (o SiteConfigResponsePtrOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.VnetName
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) VnetPrivatePortsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.VnetPrivatePortsCount
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponsePtrOutput) VnetRouteAllEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.VnetRouteAllEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) WebSocketsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.WebSocketsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) WebsiteTimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.WebsiteTimeZone
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) WindowsFxVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.WindowsFxVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) XManagedServiceIdentityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.XManagedServiceIdentityId
	}).(pulumi.IntPtrOutput)
}

type SiteLimits struct {
	MaxDiskSizeInMb  *float64 `pulumi:"maxDiskSizeInMb"`
	MaxMemoryInMb    *float64 `pulumi:"maxMemoryInMb"`
	MaxPercentageCpu *float64 `pulumi:"maxPercentageCpu"`
}





type SiteLimitsInput interface {
	pulumi.Input

	ToSiteLimitsOutput() SiteLimitsOutput
	ToSiteLimitsOutputWithContext(context.Context) SiteLimitsOutput
}

type SiteLimitsArgs struct {
	MaxDiskSizeInMb  pulumi.Float64PtrInput `pulumi:"maxDiskSizeInMb"`
	MaxMemoryInMb    pulumi.Float64PtrInput `pulumi:"maxMemoryInMb"`
	MaxPercentageCpu pulumi.Float64PtrInput `pulumi:"maxPercentageCpu"`
}

func (SiteLimitsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteLimits)(nil)).Elem()
}

func (i SiteLimitsArgs) ToSiteLimitsOutput() SiteLimitsOutput {
	return i.ToSiteLimitsOutputWithContext(context.Background())
}

func (i SiteLimitsArgs) ToSiteLimitsOutputWithContext(ctx context.Context) SiteLimitsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteLimitsOutput)
}

func (i SiteLimitsArgs) ToSiteLimitsPtrOutput() SiteLimitsPtrOutput {
	return i.ToSiteLimitsPtrOutputWithContext(context.Background())
}

func (i SiteLimitsArgs) ToSiteLimitsPtrOutputWithContext(ctx context.Context) SiteLimitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteLimitsOutput).ToSiteLimitsPtrOutputWithContext(ctx)
}









type SiteLimitsPtrInput interface {
	pulumi.Input

	ToSiteLimitsPtrOutput() SiteLimitsPtrOutput
	ToSiteLimitsPtrOutputWithContext(context.Context) SiteLimitsPtrOutput
}

type siteLimitsPtrType SiteLimitsArgs

func SiteLimitsPtr(v *SiteLimitsArgs) SiteLimitsPtrInput {
	return (*siteLimitsPtrType)(v)
}

func (*siteLimitsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteLimits)(nil)).Elem()
}

func (i *siteLimitsPtrType) ToSiteLimitsPtrOutput() SiteLimitsPtrOutput {
	return i.ToSiteLimitsPtrOutputWithContext(context.Background())
}

func (i *siteLimitsPtrType) ToSiteLimitsPtrOutputWithContext(ctx context.Context) SiteLimitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteLimitsPtrOutput)
}

type SiteLimitsOutput struct{ *pulumi.OutputState }

func (SiteLimitsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteLimits)(nil)).Elem()
}

func (o SiteLimitsOutput) ToSiteLimitsOutput() SiteLimitsOutput {
	return o
}

func (o SiteLimitsOutput) ToSiteLimitsOutputWithContext(ctx context.Context) SiteLimitsOutput {
	return o
}

func (o SiteLimitsOutput) ToSiteLimitsPtrOutput() SiteLimitsPtrOutput {
	return o.ToSiteLimitsPtrOutputWithContext(context.Background())
}

func (o SiteLimitsOutput) ToSiteLimitsPtrOutputWithContext(ctx context.Context) SiteLimitsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteLimits) *SiteLimits {
		return &v
	}).(SiteLimitsPtrOutput)
}

func (o SiteLimitsOutput) MaxDiskSizeInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SiteLimits) *float64 { return v.MaxDiskSizeInMb }).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsOutput) MaxMemoryInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SiteLimits) *float64 { return v.MaxMemoryInMb }).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsOutput) MaxPercentageCpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SiteLimits) *float64 { return v.MaxPercentageCpu }).(pulumi.Float64PtrOutput)
}

type SiteLimitsPtrOutput struct{ *pulumi.OutputState }

func (SiteLimitsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteLimits)(nil)).Elem()
}

func (o SiteLimitsPtrOutput) ToSiteLimitsPtrOutput() SiteLimitsPtrOutput {
	return o
}

func (o SiteLimitsPtrOutput) ToSiteLimitsPtrOutputWithContext(ctx context.Context) SiteLimitsPtrOutput {
	return o
}

func (o SiteLimitsPtrOutput) Elem() SiteLimitsOutput {
	return o.ApplyT(func(v *SiteLimits) SiteLimits {
		if v != nil {
			return *v
		}
		var ret SiteLimits
		return ret
	}).(SiteLimitsOutput)
}

func (o SiteLimitsPtrOutput) MaxDiskSizeInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteLimits) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxDiskSizeInMb
	}).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsPtrOutput) MaxMemoryInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteLimits) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxMemoryInMb
	}).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsPtrOutput) MaxPercentageCpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteLimits) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxPercentageCpu
	}).(pulumi.Float64PtrOutput)
}

type SiteLimitsResponse struct {
	MaxDiskSizeInMb  *float64 `pulumi:"maxDiskSizeInMb"`
	MaxMemoryInMb    *float64 `pulumi:"maxMemoryInMb"`
	MaxPercentageCpu *float64 `pulumi:"maxPercentageCpu"`
}

type SiteLimitsResponseOutput struct{ *pulumi.OutputState }

func (SiteLimitsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteLimitsResponse)(nil)).Elem()
}

func (o SiteLimitsResponseOutput) ToSiteLimitsResponseOutput() SiteLimitsResponseOutput {
	return o
}

func (o SiteLimitsResponseOutput) ToSiteLimitsResponseOutputWithContext(ctx context.Context) SiteLimitsResponseOutput {
	return o
}

func (o SiteLimitsResponseOutput) MaxDiskSizeInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SiteLimitsResponse) *float64 { return v.MaxDiskSizeInMb }).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsResponseOutput) MaxMemoryInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SiteLimitsResponse) *float64 { return v.MaxMemoryInMb }).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsResponseOutput) MaxPercentageCpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SiteLimitsResponse) *float64 { return v.MaxPercentageCpu }).(pulumi.Float64PtrOutput)
}

type SiteLimitsResponsePtrOutput struct{ *pulumi.OutputState }

func (SiteLimitsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteLimitsResponse)(nil)).Elem()
}

func (o SiteLimitsResponsePtrOutput) ToSiteLimitsResponsePtrOutput() SiteLimitsResponsePtrOutput {
	return o
}

func (o SiteLimitsResponsePtrOutput) ToSiteLimitsResponsePtrOutputWithContext(ctx context.Context) SiteLimitsResponsePtrOutput {
	return o
}

func (o SiteLimitsResponsePtrOutput) Elem() SiteLimitsResponseOutput {
	return o.ApplyT(func(v *SiteLimitsResponse) SiteLimitsResponse {
		if v != nil {
			return *v
		}
		var ret SiteLimitsResponse
		return ret
	}).(SiteLimitsResponseOutput)
}

func (o SiteLimitsResponsePtrOutput) MaxDiskSizeInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteLimitsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxDiskSizeInMb
	}).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsResponsePtrOutput) MaxMemoryInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteLimitsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxMemoryInMb
	}).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsResponsePtrOutput) MaxPercentageCpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteLimitsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxPercentageCpu
	}).(pulumi.Float64PtrOutput)
}

type SiteMachineKeyResponse struct {
	Decryption    *string `pulumi:"decryption"`
	DecryptionKey *string `pulumi:"decryptionKey"`
	Validation    *string `pulumi:"validation"`
	ValidationKey *string `pulumi:"validationKey"`
}

type SiteMachineKeyResponseOutput struct{ *pulumi.OutputState }

func (SiteMachineKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteMachineKeyResponse)(nil)).Elem()
}

func (o SiteMachineKeyResponseOutput) ToSiteMachineKeyResponseOutput() SiteMachineKeyResponseOutput {
	return o
}

func (o SiteMachineKeyResponseOutput) ToSiteMachineKeyResponseOutputWithContext(ctx context.Context) SiteMachineKeyResponseOutput {
	return o
}

func (o SiteMachineKeyResponseOutput) Decryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteMachineKeyResponse) *string { return v.Decryption }).(pulumi.StringPtrOutput)
}

func (o SiteMachineKeyResponseOutput) DecryptionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteMachineKeyResponse) *string { return v.DecryptionKey }).(pulumi.StringPtrOutput)
}

func (o SiteMachineKeyResponseOutput) Validation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteMachineKeyResponse) *string { return v.Validation }).(pulumi.StringPtrOutput)
}

func (o SiteMachineKeyResponseOutput) ValidationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteMachineKeyResponse) *string { return v.ValidationKey }).(pulumi.StringPtrOutput)
}

type SiteMachineKeyResponsePtrOutput struct{ *pulumi.OutputState }

func (SiteMachineKeyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteMachineKeyResponse)(nil)).Elem()
}

func (o SiteMachineKeyResponsePtrOutput) ToSiteMachineKeyResponsePtrOutput() SiteMachineKeyResponsePtrOutput {
	return o
}

func (o SiteMachineKeyResponsePtrOutput) ToSiteMachineKeyResponsePtrOutputWithContext(ctx context.Context) SiteMachineKeyResponsePtrOutput {
	return o
}

func (o SiteMachineKeyResponsePtrOutput) Elem() SiteMachineKeyResponseOutput {
	return o.ApplyT(func(v *SiteMachineKeyResponse) SiteMachineKeyResponse {
		if v != nil {
			return *v
		}
		var ret SiteMachineKeyResponse
		return ret
	}).(SiteMachineKeyResponseOutput)
}

func (o SiteMachineKeyResponsePtrOutput) Decryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteMachineKeyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Decryption
	}).(pulumi.StringPtrOutput)
}

func (o SiteMachineKeyResponsePtrOutput) DecryptionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteMachineKeyResponse) *string {
		if v == nil {
			return nil
		}
		return v.DecryptionKey
	}).(pulumi.StringPtrOutput)
}

func (o SiteMachineKeyResponsePtrOutput) Validation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteMachineKeyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Validation
	}).(pulumi.StringPtrOutput)
}

func (o SiteMachineKeyResponsePtrOutput) ValidationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteMachineKeyResponse) *string {
		if v == nil {
			return nil
		}
		return v.ValidationKey
	}).(pulumi.StringPtrOutput)
}

type SkuCapacity struct {
	Default        *int    `pulumi:"default"`
	ElasticMaximum *int    `pulumi:"elasticMaximum"`
	Maximum        *int    `pulumi:"maximum"`
	Minimum        *int    `pulumi:"minimum"`
	ScaleType      *string `pulumi:"scaleType"`
}





type SkuCapacityInput interface {
	pulumi.Input

	ToSkuCapacityOutput() SkuCapacityOutput
	ToSkuCapacityOutputWithContext(context.Context) SkuCapacityOutput
}

type SkuCapacityArgs struct {
	Default        pulumi.IntPtrInput    `pulumi:"default"`
	ElasticMaximum pulumi.IntPtrInput    `pulumi:"elasticMaximum"`
	Maximum        pulumi.IntPtrInput    `pulumi:"maximum"`
	Minimum        pulumi.IntPtrInput    `pulumi:"minimum"`
	ScaleType      pulumi.StringPtrInput `pulumi:"scaleType"`
}

func (SkuCapacityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuCapacity)(nil)).Elem()
}

func (i SkuCapacityArgs) ToSkuCapacityOutput() SkuCapacityOutput {
	return i.ToSkuCapacityOutputWithContext(context.Background())
}

func (i SkuCapacityArgs) ToSkuCapacityOutputWithContext(ctx context.Context) SkuCapacityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuCapacityOutput)
}

func (i SkuCapacityArgs) ToSkuCapacityPtrOutput() SkuCapacityPtrOutput {
	return i.ToSkuCapacityPtrOutputWithContext(context.Background())
}

func (i SkuCapacityArgs) ToSkuCapacityPtrOutputWithContext(ctx context.Context) SkuCapacityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuCapacityOutput).ToSkuCapacityPtrOutputWithContext(ctx)
}









type SkuCapacityPtrInput interface {
	pulumi.Input

	ToSkuCapacityPtrOutput() SkuCapacityPtrOutput
	ToSkuCapacityPtrOutputWithContext(context.Context) SkuCapacityPtrOutput
}

type skuCapacityPtrType SkuCapacityArgs

func SkuCapacityPtr(v *SkuCapacityArgs) SkuCapacityPtrInput {
	return (*skuCapacityPtrType)(v)
}

func (*skuCapacityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuCapacity)(nil)).Elem()
}

func (i *skuCapacityPtrType) ToSkuCapacityPtrOutput() SkuCapacityPtrOutput {
	return i.ToSkuCapacityPtrOutputWithContext(context.Background())
}

func (i *skuCapacityPtrType) ToSkuCapacityPtrOutputWithContext(ctx context.Context) SkuCapacityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuCapacityPtrOutput)
}

type SkuCapacityOutput struct{ *pulumi.OutputState }

func (SkuCapacityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuCapacity)(nil)).Elem()
}

func (o SkuCapacityOutput) ToSkuCapacityOutput() SkuCapacityOutput {
	return o
}

func (o SkuCapacityOutput) ToSkuCapacityOutputWithContext(ctx context.Context) SkuCapacityOutput {
	return o
}

func (o SkuCapacityOutput) ToSkuCapacityPtrOutput() SkuCapacityPtrOutput {
	return o.ToSkuCapacityPtrOutputWithContext(context.Background())
}

func (o SkuCapacityOutput) ToSkuCapacityPtrOutputWithContext(ctx context.Context) SkuCapacityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuCapacity) *SkuCapacity {
		return &v
	}).(SkuCapacityPtrOutput)
}

func (o SkuCapacityOutput) Default() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacity) *int { return v.Default }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityOutput) ElasticMaximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacity) *int { return v.ElasticMaximum }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityOutput) Maximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacity) *int { return v.Maximum }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityOutput) Minimum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacity) *int { return v.Minimum }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityOutput) ScaleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuCapacity) *string { return v.ScaleType }).(pulumi.StringPtrOutput)
}

type SkuCapacityPtrOutput struct{ *pulumi.OutputState }

func (SkuCapacityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuCapacity)(nil)).Elem()
}

func (o SkuCapacityPtrOutput) ToSkuCapacityPtrOutput() SkuCapacityPtrOutput {
	return o
}

func (o SkuCapacityPtrOutput) ToSkuCapacityPtrOutputWithContext(ctx context.Context) SkuCapacityPtrOutput {
	return o
}

func (o SkuCapacityPtrOutput) Elem() SkuCapacityOutput {
	return o.ApplyT(func(v *SkuCapacity) SkuCapacity {
		if v != nil {
			return *v
		}
		var ret SkuCapacity
		return ret
	}).(SkuCapacityOutput)
}

func (o SkuCapacityPtrOutput) Default() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacity) *int {
		if v == nil {
			return nil
		}
		return v.Default
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityPtrOutput) ElasticMaximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacity) *int {
		if v == nil {
			return nil
		}
		return v.ElasticMaximum
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityPtrOutput) Maximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacity) *int {
		if v == nil {
			return nil
		}
		return v.Maximum
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityPtrOutput) Minimum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacity) *int {
		if v == nil {
			return nil
		}
		return v.Minimum
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityPtrOutput) ScaleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuCapacity) *string {
		if v == nil {
			return nil
		}
		return v.ScaleType
	}).(pulumi.StringPtrOutput)
}

type SkuCapacityResponse struct {
	Default        *int    `pulumi:"default"`
	ElasticMaximum *int    `pulumi:"elasticMaximum"`
	Maximum        *int    `pulumi:"maximum"`
	Minimum        *int    `pulumi:"minimum"`
	ScaleType      *string `pulumi:"scaleType"`
}

type SkuCapacityResponseOutput struct{ *pulumi.OutputState }

func (SkuCapacityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuCapacityResponse)(nil)).Elem()
}

func (o SkuCapacityResponseOutput) ToSkuCapacityResponseOutput() SkuCapacityResponseOutput {
	return o
}

func (o SkuCapacityResponseOutput) ToSkuCapacityResponseOutputWithContext(ctx context.Context) SkuCapacityResponseOutput {
	return o
}

func (o SkuCapacityResponseOutput) Default() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacityResponse) *int { return v.Default }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponseOutput) ElasticMaximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacityResponse) *int { return v.ElasticMaximum }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponseOutput) Maximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacityResponse) *int { return v.Maximum }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponseOutput) Minimum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacityResponse) *int { return v.Minimum }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponseOutput) ScaleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuCapacityResponse) *string { return v.ScaleType }).(pulumi.StringPtrOutput)
}

type SkuCapacityResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuCapacityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuCapacityResponse)(nil)).Elem()
}

func (o SkuCapacityResponsePtrOutput) ToSkuCapacityResponsePtrOutput() SkuCapacityResponsePtrOutput {
	return o
}

func (o SkuCapacityResponsePtrOutput) ToSkuCapacityResponsePtrOutputWithContext(ctx context.Context) SkuCapacityResponsePtrOutput {
	return o
}

func (o SkuCapacityResponsePtrOutput) Elem() SkuCapacityResponseOutput {
	return o.ApplyT(func(v *SkuCapacityResponse) SkuCapacityResponse {
		if v != nil {
			return *v
		}
		var ret SkuCapacityResponse
		return ret
	}).(SkuCapacityResponseOutput)
}

func (o SkuCapacityResponsePtrOutput) Default() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacityResponse) *int {
		if v == nil {
			return nil
		}
		return v.Default
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponsePtrOutput) ElasticMaximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacityResponse) *int {
		if v == nil {
			return nil
		}
		return v.ElasticMaximum
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponsePtrOutput) Maximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacityResponse) *int {
		if v == nil {
			return nil
		}
		return v.Maximum
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponsePtrOutput) Minimum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacityResponse) *int {
		if v == nil {
			return nil
		}
		return v.Minimum
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponsePtrOutput) ScaleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuCapacityResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScaleType
	}).(pulumi.StringPtrOutput)
}

type SkuDescription struct {
	Capabilities []Capability `pulumi:"capabilities"`
	Capacity     *int         `pulumi:"capacity"`
	Family       *string      `pulumi:"family"`
	Locations    []string     `pulumi:"locations"`
	Name         *string      `pulumi:"name"`
	Size         *string      `pulumi:"size"`
	SkuCapacity  *SkuCapacity `pulumi:"skuCapacity"`
	Tier         *string      `pulumi:"tier"`
}





type SkuDescriptionInput interface {
	pulumi.Input

	ToSkuDescriptionOutput() SkuDescriptionOutput
	ToSkuDescriptionOutputWithContext(context.Context) SkuDescriptionOutput
}

type SkuDescriptionArgs struct {
	Capabilities CapabilityArrayInput    `pulumi:"capabilities"`
	Capacity     pulumi.IntPtrInput      `pulumi:"capacity"`
	Family       pulumi.StringPtrInput   `pulumi:"family"`
	Locations    pulumi.StringArrayInput `pulumi:"locations"`
	Name         pulumi.StringPtrInput   `pulumi:"name"`
	Size         pulumi.StringPtrInput   `pulumi:"size"`
	SkuCapacity  SkuCapacityPtrInput     `pulumi:"skuCapacity"`
	Tier         pulumi.StringPtrInput   `pulumi:"tier"`
}

func (SkuDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDescription)(nil)).Elem()
}

func (i SkuDescriptionArgs) ToSkuDescriptionOutput() SkuDescriptionOutput {
	return i.ToSkuDescriptionOutputWithContext(context.Background())
}

func (i SkuDescriptionArgs) ToSkuDescriptionOutputWithContext(ctx context.Context) SkuDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionOutput)
}

func (i SkuDescriptionArgs) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return i.ToSkuDescriptionPtrOutputWithContext(context.Background())
}

func (i SkuDescriptionArgs) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionOutput).ToSkuDescriptionPtrOutputWithContext(ctx)
}









type SkuDescriptionPtrInput interface {
	pulumi.Input

	ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput
	ToSkuDescriptionPtrOutputWithContext(context.Context) SkuDescriptionPtrOutput
}

type skuDescriptionPtrType SkuDescriptionArgs

func SkuDescriptionPtr(v *SkuDescriptionArgs) SkuDescriptionPtrInput {
	return (*skuDescriptionPtrType)(v)
}

func (*skuDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuDescription)(nil)).Elem()
}

func (i *skuDescriptionPtrType) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return i.ToSkuDescriptionPtrOutputWithContext(context.Background())
}

func (i *skuDescriptionPtrType) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionPtrOutput)
}

type SkuDescriptionOutput struct{ *pulumi.OutputState }

func (SkuDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDescription)(nil)).Elem()
}

func (o SkuDescriptionOutput) ToSkuDescriptionOutput() SkuDescriptionOutput {
	return o
}

func (o SkuDescriptionOutput) ToSkuDescriptionOutputWithContext(ctx context.Context) SkuDescriptionOutput {
	return o
}

func (o SkuDescriptionOutput) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return o.ToSkuDescriptionPtrOutputWithContext(context.Background())
}

func (o SkuDescriptionOutput) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuDescription) *SkuDescription {
		return &v
	}).(SkuDescriptionPtrOutput)
}

func (o SkuDescriptionOutput) Capabilities() CapabilityArrayOutput {
	return o.ApplyT(func(v SkuDescription) []Capability { return v.Capabilities }).(CapabilityArrayOutput)
}

func (o SkuDescriptionOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuDescription) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SkuDescription) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o SkuDescriptionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionOutput) SkuCapacity() SkuCapacityPtrOutput {
	return o.ApplyT(func(v SkuDescription) *SkuCapacity { return v.SkuCapacity }).(SkuCapacityPtrOutput)
}

func (o SkuDescriptionOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuDescriptionPtrOutput struct{ *pulumi.OutputState }

func (SkuDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuDescription)(nil)).Elem()
}

func (o SkuDescriptionPtrOutput) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return o
}

func (o SkuDescriptionPtrOutput) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return o
}

func (o SkuDescriptionPtrOutput) Elem() SkuDescriptionOutput {
	return o.ApplyT(func(v *SkuDescription) SkuDescription {
		if v != nil {
			return *v
		}
		var ret SkuDescription
		return ret
	}).(SkuDescriptionOutput)
}

func (o SkuDescriptionPtrOutput) Capabilities() CapabilityArrayOutput {
	return o.ApplyT(func(v *SkuDescription) []Capability {
		if v == nil {
			return nil
		}
		return v.Capabilities
	}).(CapabilityArrayOutput)
}

func (o SkuDescriptionPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionPtrOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SkuDescription) []string {
		if v == nil {
			return nil
		}
		return v.Locations
	}).(pulumi.StringArrayOutput)
}

func (o SkuDescriptionPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionPtrOutput) SkuCapacity() SkuCapacityPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *SkuCapacity {
		if v == nil {
			return nil
		}
		return v.SkuCapacity
	}).(SkuCapacityPtrOutput)
}

func (o SkuDescriptionPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuDescriptionResponse struct {
	Capabilities []CapabilityResponse `pulumi:"capabilities"`
	Capacity     *int                 `pulumi:"capacity"`
	Family       *string              `pulumi:"family"`
	Locations    []string             `pulumi:"locations"`
	Name         *string              `pulumi:"name"`
	Size         *string              `pulumi:"size"`
	SkuCapacity  *SkuCapacityResponse `pulumi:"skuCapacity"`
	Tier         *string              `pulumi:"tier"`
}

type SkuDescriptionResponseOutput struct{ *pulumi.OutputState }

func (SkuDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDescriptionResponse)(nil)).Elem()
}

func (o SkuDescriptionResponseOutput) ToSkuDescriptionResponseOutput() SkuDescriptionResponseOutput {
	return o
}

func (o SkuDescriptionResponseOutput) ToSkuDescriptionResponseOutputWithContext(ctx context.Context) SkuDescriptionResponseOutput {
	return o
}

func (o SkuDescriptionResponseOutput) Capabilities() CapabilityResponseArrayOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) []CapabilityResponse { return v.Capabilities }).(CapabilityResponseArrayOutput)
}

func (o SkuDescriptionResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponseOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o SkuDescriptionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponseOutput) SkuCapacity() SkuCapacityResponsePtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *SkuCapacityResponse { return v.SkuCapacity }).(SkuCapacityResponsePtrOutput)
}

func (o SkuDescriptionResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuDescriptionResponse)(nil)).Elem()
}

func (o SkuDescriptionResponsePtrOutput) ToSkuDescriptionResponsePtrOutput() SkuDescriptionResponsePtrOutput {
	return o
}

func (o SkuDescriptionResponsePtrOutput) ToSkuDescriptionResponsePtrOutputWithContext(ctx context.Context) SkuDescriptionResponsePtrOutput {
	return o
}

func (o SkuDescriptionResponsePtrOutput) Elem() SkuDescriptionResponseOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) SkuDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret SkuDescriptionResponse
		return ret
	}).(SkuDescriptionResponseOutput)
}

func (o SkuDescriptionResponsePtrOutput) Capabilities() CapabilityResponseArrayOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) []CapabilityResponse {
		if v == nil {
			return nil
		}
		return v.Capabilities
	}).(CapabilityResponseArrayOutput)
}

func (o SkuDescriptionResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) []string {
		if v == nil {
			return nil
		}
		return v.Locations
	}).(pulumi.StringArrayOutput)
}

func (o SkuDescriptionResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) SkuCapacity() SkuCapacityResponsePtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *SkuCapacityResponse {
		if v == nil {
			return nil
		}
		return v.SkuCapacity
	}).(SkuCapacityResponsePtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SlotSwapStatusResponse struct {
	DestinationSlotName string `pulumi:"destinationSlotName"`
	SourceSlotName      string `pulumi:"sourceSlotName"`
	TimestampUtc        string `pulumi:"timestampUtc"`
}

type SlotSwapStatusResponseOutput struct{ *pulumi.OutputState }

func (SlotSwapStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SlotSwapStatusResponse)(nil)).Elem()
}

func (o SlotSwapStatusResponseOutput) ToSlotSwapStatusResponseOutput() SlotSwapStatusResponseOutput {
	return o
}

func (o SlotSwapStatusResponseOutput) ToSlotSwapStatusResponseOutputWithContext(ctx context.Context) SlotSwapStatusResponseOutput {
	return o
}

func (o SlotSwapStatusResponseOutput) DestinationSlotName() pulumi.StringOutput {
	return o.ApplyT(func(v SlotSwapStatusResponse) string { return v.DestinationSlotName }).(pulumi.StringOutput)
}

func (o SlotSwapStatusResponseOutput) SourceSlotName() pulumi.StringOutput {
	return o.ApplyT(func(v SlotSwapStatusResponse) string { return v.SourceSlotName }).(pulumi.StringOutput)
}

func (o SlotSwapStatusResponseOutput) TimestampUtc() pulumi.StringOutput {
	return o.ApplyT(func(v SlotSwapStatusResponse) string { return v.TimestampUtc }).(pulumi.StringOutput)
}

type SlowRequestsBasedTrigger struct {
	Count        *int    `pulumi:"count"`
	Path         *string `pulumi:"path"`
	TimeInterval *string `pulumi:"timeInterval"`
	TimeTaken    *string `pulumi:"timeTaken"`
}





type SlowRequestsBasedTriggerInput interface {
	pulumi.Input

	ToSlowRequestsBasedTriggerOutput() SlowRequestsBasedTriggerOutput
	ToSlowRequestsBasedTriggerOutputWithContext(context.Context) SlowRequestsBasedTriggerOutput
}

type SlowRequestsBasedTriggerArgs struct {
	Count        pulumi.IntPtrInput    `pulumi:"count"`
	Path         pulumi.StringPtrInput `pulumi:"path"`
	TimeInterval pulumi.StringPtrInput `pulumi:"timeInterval"`
	TimeTaken    pulumi.StringPtrInput `pulumi:"timeTaken"`
}

func (SlowRequestsBasedTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SlowRequestsBasedTrigger)(nil)).Elem()
}

func (i SlowRequestsBasedTriggerArgs) ToSlowRequestsBasedTriggerOutput() SlowRequestsBasedTriggerOutput {
	return i.ToSlowRequestsBasedTriggerOutputWithContext(context.Background())
}

func (i SlowRequestsBasedTriggerArgs) ToSlowRequestsBasedTriggerOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlowRequestsBasedTriggerOutput)
}

func (i SlowRequestsBasedTriggerArgs) ToSlowRequestsBasedTriggerPtrOutput() SlowRequestsBasedTriggerPtrOutput {
	return i.ToSlowRequestsBasedTriggerPtrOutputWithContext(context.Background())
}

func (i SlowRequestsBasedTriggerArgs) ToSlowRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlowRequestsBasedTriggerOutput).ToSlowRequestsBasedTriggerPtrOutputWithContext(ctx)
}









type SlowRequestsBasedTriggerPtrInput interface {
	pulumi.Input

	ToSlowRequestsBasedTriggerPtrOutput() SlowRequestsBasedTriggerPtrOutput
	ToSlowRequestsBasedTriggerPtrOutputWithContext(context.Context) SlowRequestsBasedTriggerPtrOutput
}

type slowRequestsBasedTriggerPtrType SlowRequestsBasedTriggerArgs

func SlowRequestsBasedTriggerPtr(v *SlowRequestsBasedTriggerArgs) SlowRequestsBasedTriggerPtrInput {
	return (*slowRequestsBasedTriggerPtrType)(v)
}

func (*slowRequestsBasedTriggerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SlowRequestsBasedTrigger)(nil)).Elem()
}

func (i *slowRequestsBasedTriggerPtrType) ToSlowRequestsBasedTriggerPtrOutput() SlowRequestsBasedTriggerPtrOutput {
	return i.ToSlowRequestsBasedTriggerPtrOutputWithContext(context.Background())
}

func (i *slowRequestsBasedTriggerPtrType) ToSlowRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlowRequestsBasedTriggerPtrOutput)
}





type SlowRequestsBasedTriggerArrayInput interface {
	pulumi.Input

	ToSlowRequestsBasedTriggerArrayOutput() SlowRequestsBasedTriggerArrayOutput
	ToSlowRequestsBasedTriggerArrayOutputWithContext(context.Context) SlowRequestsBasedTriggerArrayOutput
}

type SlowRequestsBasedTriggerArray []SlowRequestsBasedTriggerInput

func (SlowRequestsBasedTriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SlowRequestsBasedTrigger)(nil)).Elem()
}

func (i SlowRequestsBasedTriggerArray) ToSlowRequestsBasedTriggerArrayOutput() SlowRequestsBasedTriggerArrayOutput {
	return i.ToSlowRequestsBasedTriggerArrayOutputWithContext(context.Background())
}

func (i SlowRequestsBasedTriggerArray) ToSlowRequestsBasedTriggerArrayOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlowRequestsBasedTriggerArrayOutput)
}

type SlowRequestsBasedTriggerOutput struct{ *pulumi.OutputState }

func (SlowRequestsBasedTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SlowRequestsBasedTrigger)(nil)).Elem()
}

func (o SlowRequestsBasedTriggerOutput) ToSlowRequestsBasedTriggerOutput() SlowRequestsBasedTriggerOutput {
	return o
}

func (o SlowRequestsBasedTriggerOutput) ToSlowRequestsBasedTriggerOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerOutput {
	return o
}

func (o SlowRequestsBasedTriggerOutput) ToSlowRequestsBasedTriggerPtrOutput() SlowRequestsBasedTriggerPtrOutput {
	return o.ToSlowRequestsBasedTriggerPtrOutputWithContext(context.Background())
}

func (o SlowRequestsBasedTriggerOutput) ToSlowRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SlowRequestsBasedTrigger) *SlowRequestsBasedTrigger {
		return &v
	}).(SlowRequestsBasedTriggerPtrOutput)
}

func (o SlowRequestsBasedTriggerOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTrigger) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o SlowRequestsBasedTriggerOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTrigger) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTrigger) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerOutput) TimeTaken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTrigger) *string { return v.TimeTaken }).(pulumi.StringPtrOutput)
}

type SlowRequestsBasedTriggerPtrOutput struct{ *pulumi.OutputState }

func (SlowRequestsBasedTriggerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SlowRequestsBasedTrigger)(nil)).Elem()
}

func (o SlowRequestsBasedTriggerPtrOutput) ToSlowRequestsBasedTriggerPtrOutput() SlowRequestsBasedTriggerPtrOutput {
	return o
}

func (o SlowRequestsBasedTriggerPtrOutput) ToSlowRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerPtrOutput {
	return o
}

func (o SlowRequestsBasedTriggerPtrOutput) Elem() SlowRequestsBasedTriggerOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTrigger) SlowRequestsBasedTrigger {
		if v != nil {
			return *v
		}
		var ret SlowRequestsBasedTrigger
		return ret
	}).(SlowRequestsBasedTriggerOutput)
}

func (o SlowRequestsBasedTriggerPtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTrigger) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o SlowRequestsBasedTriggerPtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTrigger) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerPtrOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTrigger) *string {
		if v == nil {
			return nil
		}
		return v.TimeInterval
	}).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerPtrOutput) TimeTaken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTrigger) *string {
		if v == nil {
			return nil
		}
		return v.TimeTaken
	}).(pulumi.StringPtrOutput)
}

type SlowRequestsBasedTriggerArrayOutput struct{ *pulumi.OutputState }

func (SlowRequestsBasedTriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SlowRequestsBasedTrigger)(nil)).Elem()
}

func (o SlowRequestsBasedTriggerArrayOutput) ToSlowRequestsBasedTriggerArrayOutput() SlowRequestsBasedTriggerArrayOutput {
	return o
}

func (o SlowRequestsBasedTriggerArrayOutput) ToSlowRequestsBasedTriggerArrayOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerArrayOutput {
	return o
}

func (o SlowRequestsBasedTriggerArrayOutput) Index(i pulumi.IntInput) SlowRequestsBasedTriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SlowRequestsBasedTrigger {
		return vs[0].([]SlowRequestsBasedTrigger)[vs[1].(int)]
	}).(SlowRequestsBasedTriggerOutput)
}

type SlowRequestsBasedTriggerResponse struct {
	Count        *int    `pulumi:"count"`
	Path         *string `pulumi:"path"`
	TimeInterval *string `pulumi:"timeInterval"`
	TimeTaken    *string `pulumi:"timeTaken"`
}

type SlowRequestsBasedTriggerResponseOutput struct{ *pulumi.OutputState }

func (SlowRequestsBasedTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SlowRequestsBasedTriggerResponse)(nil)).Elem()
}

func (o SlowRequestsBasedTriggerResponseOutput) ToSlowRequestsBasedTriggerResponseOutput() SlowRequestsBasedTriggerResponseOutput {
	return o
}

func (o SlowRequestsBasedTriggerResponseOutput) ToSlowRequestsBasedTriggerResponseOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerResponseOutput {
	return o
}

func (o SlowRequestsBasedTriggerResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTriggerResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o SlowRequestsBasedTriggerResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTriggerResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerResponseOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTriggerResponse) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerResponseOutput) TimeTaken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTriggerResponse) *string { return v.TimeTaken }).(pulumi.StringPtrOutput)
}

type SlowRequestsBasedTriggerResponsePtrOutput struct{ *pulumi.OutputState }

func (SlowRequestsBasedTriggerResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SlowRequestsBasedTriggerResponse)(nil)).Elem()
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) ToSlowRequestsBasedTriggerResponsePtrOutput() SlowRequestsBasedTriggerResponsePtrOutput {
	return o
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) ToSlowRequestsBasedTriggerResponsePtrOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerResponsePtrOutput {
	return o
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) Elem() SlowRequestsBasedTriggerResponseOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTriggerResponse) SlowRequestsBasedTriggerResponse {
		if v != nil {
			return *v
		}
		var ret SlowRequestsBasedTriggerResponse
		return ret
	}).(SlowRequestsBasedTriggerResponseOutput)
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTriggerResponse) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeInterval
	}).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) TimeTaken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeTaken
	}).(pulumi.StringPtrOutput)
}

type SlowRequestsBasedTriggerResponseArrayOutput struct{ *pulumi.OutputState }

func (SlowRequestsBasedTriggerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SlowRequestsBasedTriggerResponse)(nil)).Elem()
}

func (o SlowRequestsBasedTriggerResponseArrayOutput) ToSlowRequestsBasedTriggerResponseArrayOutput() SlowRequestsBasedTriggerResponseArrayOutput {
	return o
}

func (o SlowRequestsBasedTriggerResponseArrayOutput) ToSlowRequestsBasedTriggerResponseArrayOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerResponseArrayOutput {
	return o
}

func (o SlowRequestsBasedTriggerResponseArrayOutput) Index(i pulumi.IntInput) SlowRequestsBasedTriggerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SlowRequestsBasedTriggerResponse {
		return vs[0].([]SlowRequestsBasedTriggerResponse)[vs[1].(int)]
	}).(SlowRequestsBasedTriggerResponseOutput)
}

type StaticSiteBuildProperties struct {
	ApiBuildCommand                    *string `pulumi:"apiBuildCommand"`
	ApiLocation                        *string `pulumi:"apiLocation"`
	AppArtifactLocation                *string `pulumi:"appArtifactLocation"`
	AppBuildCommand                    *string `pulumi:"appBuildCommand"`
	AppLocation                        *string `pulumi:"appLocation"`
	GithubActionSecretNameOverride     *string `pulumi:"githubActionSecretNameOverride"`
	OutputLocation                     *string `pulumi:"outputLocation"`
	SkipGithubActionWorkflowGeneration *bool   `pulumi:"skipGithubActionWorkflowGeneration"`
}





type StaticSiteBuildPropertiesInput interface {
	pulumi.Input

	ToStaticSiteBuildPropertiesOutput() StaticSiteBuildPropertiesOutput
	ToStaticSiteBuildPropertiesOutputWithContext(context.Context) StaticSiteBuildPropertiesOutput
}

type StaticSiteBuildPropertiesArgs struct {
	ApiBuildCommand                    pulumi.StringPtrInput `pulumi:"apiBuildCommand"`
	ApiLocation                        pulumi.StringPtrInput `pulumi:"apiLocation"`
	AppArtifactLocation                pulumi.StringPtrInput `pulumi:"appArtifactLocation"`
	AppBuildCommand                    pulumi.StringPtrInput `pulumi:"appBuildCommand"`
	AppLocation                        pulumi.StringPtrInput `pulumi:"appLocation"`
	GithubActionSecretNameOverride     pulumi.StringPtrInput `pulumi:"githubActionSecretNameOverride"`
	OutputLocation                     pulumi.StringPtrInput `pulumi:"outputLocation"`
	SkipGithubActionWorkflowGeneration pulumi.BoolPtrInput   `pulumi:"skipGithubActionWorkflowGeneration"`
}

func (StaticSiteBuildPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteBuildProperties)(nil)).Elem()
}

func (i StaticSiteBuildPropertiesArgs) ToStaticSiteBuildPropertiesOutput() StaticSiteBuildPropertiesOutput {
	return i.ToStaticSiteBuildPropertiesOutputWithContext(context.Background())
}

func (i StaticSiteBuildPropertiesArgs) ToStaticSiteBuildPropertiesOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteBuildPropertiesOutput)
}

func (i StaticSiteBuildPropertiesArgs) ToStaticSiteBuildPropertiesPtrOutput() StaticSiteBuildPropertiesPtrOutput {
	return i.ToStaticSiteBuildPropertiesPtrOutputWithContext(context.Background())
}

func (i StaticSiteBuildPropertiesArgs) ToStaticSiteBuildPropertiesPtrOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteBuildPropertiesOutput).ToStaticSiteBuildPropertiesPtrOutputWithContext(ctx)
}









type StaticSiteBuildPropertiesPtrInput interface {
	pulumi.Input

	ToStaticSiteBuildPropertiesPtrOutput() StaticSiteBuildPropertiesPtrOutput
	ToStaticSiteBuildPropertiesPtrOutputWithContext(context.Context) StaticSiteBuildPropertiesPtrOutput
}

type staticSiteBuildPropertiesPtrType StaticSiteBuildPropertiesArgs

func StaticSiteBuildPropertiesPtr(v *StaticSiteBuildPropertiesArgs) StaticSiteBuildPropertiesPtrInput {
	return (*staticSiteBuildPropertiesPtrType)(v)
}

func (*staticSiteBuildPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteBuildProperties)(nil)).Elem()
}

func (i *staticSiteBuildPropertiesPtrType) ToStaticSiteBuildPropertiesPtrOutput() StaticSiteBuildPropertiesPtrOutput {
	return i.ToStaticSiteBuildPropertiesPtrOutputWithContext(context.Background())
}

func (i *staticSiteBuildPropertiesPtrType) ToStaticSiteBuildPropertiesPtrOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteBuildPropertiesPtrOutput)
}

type StaticSiteBuildPropertiesOutput struct{ *pulumi.OutputState }

func (StaticSiteBuildPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteBuildProperties)(nil)).Elem()
}

func (o StaticSiteBuildPropertiesOutput) ToStaticSiteBuildPropertiesOutput() StaticSiteBuildPropertiesOutput {
	return o
}

func (o StaticSiteBuildPropertiesOutput) ToStaticSiteBuildPropertiesOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesOutput {
	return o
}

func (o StaticSiteBuildPropertiesOutput) ToStaticSiteBuildPropertiesPtrOutput() StaticSiteBuildPropertiesPtrOutput {
	return o.ToStaticSiteBuildPropertiesPtrOutputWithContext(context.Background())
}

func (o StaticSiteBuildPropertiesOutput) ToStaticSiteBuildPropertiesPtrOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StaticSiteBuildProperties) *StaticSiteBuildProperties {
		return &v
	}).(StaticSiteBuildPropertiesPtrOutput)
}

func (o StaticSiteBuildPropertiesOutput) ApiBuildCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *string { return v.ApiBuildCommand }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesOutput) ApiLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *string { return v.ApiLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesOutput) AppArtifactLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *string { return v.AppArtifactLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesOutput) AppBuildCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *string { return v.AppBuildCommand }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesOutput) AppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *string { return v.AppLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesOutput) GithubActionSecretNameOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *string { return v.GithubActionSecretNameOverride }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesOutput) OutputLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *string { return v.OutputLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesOutput) SkipGithubActionWorkflowGeneration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *bool { return v.SkipGithubActionWorkflowGeneration }).(pulumi.BoolPtrOutput)
}

type StaticSiteBuildPropertiesPtrOutput struct{ *pulumi.OutputState }

func (StaticSiteBuildPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteBuildProperties)(nil)).Elem()
}

func (o StaticSiteBuildPropertiesPtrOutput) ToStaticSiteBuildPropertiesPtrOutput() StaticSiteBuildPropertiesPtrOutput {
	return o
}

func (o StaticSiteBuildPropertiesPtrOutput) ToStaticSiteBuildPropertiesPtrOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesPtrOutput {
	return o
}

func (o StaticSiteBuildPropertiesPtrOutput) Elem() StaticSiteBuildPropertiesOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) StaticSiteBuildProperties {
		if v != nil {
			return *v
		}
		var ret StaticSiteBuildProperties
		return ret
	}).(StaticSiteBuildPropertiesOutput)
}

func (o StaticSiteBuildPropertiesPtrOutput) ApiBuildCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApiBuildCommand
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesPtrOutput) ApiLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApiLocation
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesPtrOutput) AppArtifactLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) *string {
		if v == nil {
			return nil
		}
		return v.AppArtifactLocation
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesPtrOutput) AppBuildCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) *string {
		if v == nil {
			return nil
		}
		return v.AppBuildCommand
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesPtrOutput) AppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) *string {
		if v == nil {
			return nil
		}
		return v.AppLocation
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesPtrOutput) GithubActionSecretNameOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) *string {
		if v == nil {
			return nil
		}
		return v.GithubActionSecretNameOverride
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesPtrOutput) OutputLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) *string {
		if v == nil {
			return nil
		}
		return v.OutputLocation
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesPtrOutput) SkipGithubActionWorkflowGeneration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) *bool {
		if v == nil {
			return nil
		}
		return v.SkipGithubActionWorkflowGeneration
	}).(pulumi.BoolPtrOutput)
}

type StaticSiteBuildPropertiesResponse struct {
	ApiBuildCommand                    *string `pulumi:"apiBuildCommand"`
	ApiLocation                        *string `pulumi:"apiLocation"`
	AppArtifactLocation                *string `pulumi:"appArtifactLocation"`
	AppBuildCommand                    *string `pulumi:"appBuildCommand"`
	AppLocation                        *string `pulumi:"appLocation"`
	GithubActionSecretNameOverride     *string `pulumi:"githubActionSecretNameOverride"`
	OutputLocation                     *string `pulumi:"outputLocation"`
	SkipGithubActionWorkflowGeneration *bool   `pulumi:"skipGithubActionWorkflowGeneration"`
}

type StaticSiteBuildPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StaticSiteBuildPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteBuildPropertiesResponse)(nil)).Elem()
}

func (o StaticSiteBuildPropertiesResponseOutput) ToStaticSiteBuildPropertiesResponseOutput() StaticSiteBuildPropertiesResponseOutput {
	return o
}

func (o StaticSiteBuildPropertiesResponseOutput) ToStaticSiteBuildPropertiesResponseOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesResponseOutput {
	return o
}

func (o StaticSiteBuildPropertiesResponseOutput) ApiBuildCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *string { return v.ApiBuildCommand }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponseOutput) ApiLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *string { return v.ApiLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponseOutput) AppArtifactLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *string { return v.AppArtifactLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponseOutput) AppBuildCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *string { return v.AppBuildCommand }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponseOutput) AppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *string { return v.AppLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponseOutput) GithubActionSecretNameOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *string { return v.GithubActionSecretNameOverride }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponseOutput) OutputLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *string { return v.OutputLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponseOutput) SkipGithubActionWorkflowGeneration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *bool { return v.SkipGithubActionWorkflowGeneration }).(pulumi.BoolPtrOutput)
}

type StaticSiteBuildPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (StaticSiteBuildPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteBuildPropertiesResponse)(nil)).Elem()
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) ToStaticSiteBuildPropertiesResponsePtrOutput() StaticSiteBuildPropertiesResponsePtrOutput {
	return o
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) ToStaticSiteBuildPropertiesResponsePtrOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesResponsePtrOutput {
	return o
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) Elem() StaticSiteBuildPropertiesResponseOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) StaticSiteBuildPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret StaticSiteBuildPropertiesResponse
		return ret
	}).(StaticSiteBuildPropertiesResponseOutput)
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) ApiBuildCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApiBuildCommand
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) ApiLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApiLocation
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) AppArtifactLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppArtifactLocation
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) AppBuildCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppBuildCommand
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) AppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppLocation
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) GithubActionSecretNameOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.GithubActionSecretNameOverride
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) OutputLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.OutputLocation
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) SkipGithubActionWorkflowGeneration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SkipGithubActionWorkflowGeneration
	}).(pulumi.BoolPtrOutput)
}

type StaticSiteTemplateOptions struct {
	Description           *string `pulumi:"description"`
	IsPrivate             *bool   `pulumi:"isPrivate"`
	Owner                 *string `pulumi:"owner"`
	RepositoryName        *string `pulumi:"repositoryName"`
	TemplateRepositoryUrl *string `pulumi:"templateRepositoryUrl"`
}





type StaticSiteTemplateOptionsInput interface {
	pulumi.Input

	ToStaticSiteTemplateOptionsOutput() StaticSiteTemplateOptionsOutput
	ToStaticSiteTemplateOptionsOutputWithContext(context.Context) StaticSiteTemplateOptionsOutput
}

type StaticSiteTemplateOptionsArgs struct {
	Description           pulumi.StringPtrInput `pulumi:"description"`
	IsPrivate             pulumi.BoolPtrInput   `pulumi:"isPrivate"`
	Owner                 pulumi.StringPtrInput `pulumi:"owner"`
	RepositoryName        pulumi.StringPtrInput `pulumi:"repositoryName"`
	TemplateRepositoryUrl pulumi.StringPtrInput `pulumi:"templateRepositoryUrl"`
}

func (StaticSiteTemplateOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteTemplateOptions)(nil)).Elem()
}

func (i StaticSiteTemplateOptionsArgs) ToStaticSiteTemplateOptionsOutput() StaticSiteTemplateOptionsOutput {
	return i.ToStaticSiteTemplateOptionsOutputWithContext(context.Background())
}

func (i StaticSiteTemplateOptionsArgs) ToStaticSiteTemplateOptionsOutputWithContext(ctx context.Context) StaticSiteTemplateOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteTemplateOptionsOutput)
}

func (i StaticSiteTemplateOptionsArgs) ToStaticSiteTemplateOptionsPtrOutput() StaticSiteTemplateOptionsPtrOutput {
	return i.ToStaticSiteTemplateOptionsPtrOutputWithContext(context.Background())
}

func (i StaticSiteTemplateOptionsArgs) ToStaticSiteTemplateOptionsPtrOutputWithContext(ctx context.Context) StaticSiteTemplateOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteTemplateOptionsOutput).ToStaticSiteTemplateOptionsPtrOutputWithContext(ctx)
}









type StaticSiteTemplateOptionsPtrInput interface {
	pulumi.Input

	ToStaticSiteTemplateOptionsPtrOutput() StaticSiteTemplateOptionsPtrOutput
	ToStaticSiteTemplateOptionsPtrOutputWithContext(context.Context) StaticSiteTemplateOptionsPtrOutput
}

type staticSiteTemplateOptionsPtrType StaticSiteTemplateOptionsArgs

func StaticSiteTemplateOptionsPtr(v *StaticSiteTemplateOptionsArgs) StaticSiteTemplateOptionsPtrInput {
	return (*staticSiteTemplateOptionsPtrType)(v)
}

func (*staticSiteTemplateOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteTemplateOptions)(nil)).Elem()
}

func (i *staticSiteTemplateOptionsPtrType) ToStaticSiteTemplateOptionsPtrOutput() StaticSiteTemplateOptionsPtrOutput {
	return i.ToStaticSiteTemplateOptionsPtrOutputWithContext(context.Background())
}

func (i *staticSiteTemplateOptionsPtrType) ToStaticSiteTemplateOptionsPtrOutputWithContext(ctx context.Context) StaticSiteTemplateOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteTemplateOptionsPtrOutput)
}

type StaticSiteTemplateOptionsOutput struct{ *pulumi.OutputState }

func (StaticSiteTemplateOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteTemplateOptions)(nil)).Elem()
}

func (o StaticSiteTemplateOptionsOutput) ToStaticSiteTemplateOptionsOutput() StaticSiteTemplateOptionsOutput {
	return o
}

func (o StaticSiteTemplateOptionsOutput) ToStaticSiteTemplateOptionsOutputWithContext(ctx context.Context) StaticSiteTemplateOptionsOutput {
	return o
}

func (o StaticSiteTemplateOptionsOutput) ToStaticSiteTemplateOptionsPtrOutput() StaticSiteTemplateOptionsPtrOutput {
	return o.ToStaticSiteTemplateOptionsPtrOutputWithContext(context.Background())
}

func (o StaticSiteTemplateOptionsOutput) ToStaticSiteTemplateOptionsPtrOutputWithContext(ctx context.Context) StaticSiteTemplateOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StaticSiteTemplateOptions) *StaticSiteTemplateOptions {
		return &v
	}).(StaticSiteTemplateOptionsPtrOutput)
}

func (o StaticSiteTemplateOptionsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptions) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsOutput) IsPrivate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptions) *bool { return v.IsPrivate }).(pulumi.BoolPtrOutput)
}

func (o StaticSiteTemplateOptionsOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptions) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptions) *string { return v.RepositoryName }).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsOutput) TemplateRepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptions) *string { return v.TemplateRepositoryUrl }).(pulumi.StringPtrOutput)
}

type StaticSiteTemplateOptionsPtrOutput struct{ *pulumi.OutputState }

func (StaticSiteTemplateOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteTemplateOptions)(nil)).Elem()
}

func (o StaticSiteTemplateOptionsPtrOutput) ToStaticSiteTemplateOptionsPtrOutput() StaticSiteTemplateOptionsPtrOutput {
	return o
}

func (o StaticSiteTemplateOptionsPtrOutput) ToStaticSiteTemplateOptionsPtrOutputWithContext(ctx context.Context) StaticSiteTemplateOptionsPtrOutput {
	return o
}

func (o StaticSiteTemplateOptionsPtrOutput) Elem() StaticSiteTemplateOptionsOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptions) StaticSiteTemplateOptions {
		if v != nil {
			return *v
		}
		var ret StaticSiteTemplateOptions
		return ret
	}).(StaticSiteTemplateOptionsOutput)
}

func (o StaticSiteTemplateOptionsPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptions) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsPtrOutput) IsPrivate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptions) *bool {
		if v == nil {
			return nil
		}
		return v.IsPrivate
	}).(pulumi.BoolPtrOutput)
}

func (o StaticSiteTemplateOptionsPtrOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptions) *string {
		if v == nil {
			return nil
		}
		return v.Owner
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsPtrOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptions) *string {
		if v == nil {
			return nil
		}
		return v.RepositoryName
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsPtrOutput) TemplateRepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptions) *string {
		if v == nil {
			return nil
		}
		return v.TemplateRepositoryUrl
	}).(pulumi.StringPtrOutput)
}

type StaticSiteTemplateOptionsResponse struct {
	Description           *string `pulumi:"description"`
	IsPrivate             *bool   `pulumi:"isPrivate"`
	Owner                 *string `pulumi:"owner"`
	RepositoryName        *string `pulumi:"repositoryName"`
	TemplateRepositoryUrl *string `pulumi:"templateRepositoryUrl"`
}

type StaticSiteTemplateOptionsResponseOutput struct{ *pulumi.OutputState }

func (StaticSiteTemplateOptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteTemplateOptionsResponse)(nil)).Elem()
}

func (o StaticSiteTemplateOptionsResponseOutput) ToStaticSiteTemplateOptionsResponseOutput() StaticSiteTemplateOptionsResponseOutput {
	return o
}

func (o StaticSiteTemplateOptionsResponseOutput) ToStaticSiteTemplateOptionsResponseOutputWithContext(ctx context.Context) StaticSiteTemplateOptionsResponseOutput {
	return o
}

func (o StaticSiteTemplateOptionsResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptionsResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsResponseOutput) IsPrivate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptionsResponse) *bool { return v.IsPrivate }).(pulumi.BoolPtrOutput)
}

func (o StaticSiteTemplateOptionsResponseOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptionsResponse) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsResponseOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptionsResponse) *string { return v.RepositoryName }).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsResponseOutput) TemplateRepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptionsResponse) *string { return v.TemplateRepositoryUrl }).(pulumi.StringPtrOutput)
}

type StaticSiteTemplateOptionsResponsePtrOutput struct{ *pulumi.OutputState }

func (StaticSiteTemplateOptionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteTemplateOptionsResponse)(nil)).Elem()
}

func (o StaticSiteTemplateOptionsResponsePtrOutput) ToStaticSiteTemplateOptionsResponsePtrOutput() StaticSiteTemplateOptionsResponsePtrOutput {
	return o
}

func (o StaticSiteTemplateOptionsResponsePtrOutput) ToStaticSiteTemplateOptionsResponsePtrOutputWithContext(ctx context.Context) StaticSiteTemplateOptionsResponsePtrOutput {
	return o
}

func (o StaticSiteTemplateOptionsResponsePtrOutput) Elem() StaticSiteTemplateOptionsResponseOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptionsResponse) StaticSiteTemplateOptionsResponse {
		if v != nil {
			return *v
		}
		var ret StaticSiteTemplateOptionsResponse
		return ret
	}).(StaticSiteTemplateOptionsResponseOutput)
}

func (o StaticSiteTemplateOptionsResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsResponsePtrOutput) IsPrivate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptionsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsPrivate
	}).(pulumi.BoolPtrOutput)
}

func (o StaticSiteTemplateOptionsResponsePtrOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Owner
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsResponsePtrOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.RepositoryName
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsResponsePtrOutput) TemplateRepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TemplateRepositoryUrl
	}).(pulumi.StringPtrOutput)
}

type StaticSiteUserARMResourceResponse struct {
	DisplayName string  `pulumi:"displayName"`
	Id          string  `pulumi:"id"`
	Kind        *string `pulumi:"kind"`
	Name        string  `pulumi:"name"`
	Provider    string  `pulumi:"provider"`
	Roles       *string `pulumi:"roles"`
	Type        string  `pulumi:"type"`
	UserId      string  `pulumi:"userId"`
}

type StaticSiteUserARMResourceResponseOutput struct{ *pulumi.OutputState }

func (StaticSiteUserARMResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteUserARMResourceResponse)(nil)).Elem()
}

func (o StaticSiteUserARMResourceResponseOutput) ToStaticSiteUserARMResourceResponseOutput() StaticSiteUserARMResourceResponseOutput {
	return o
}

func (o StaticSiteUserARMResourceResponseOutput) ToStaticSiteUserARMResourceResponseOutputWithContext(ctx context.Context) StaticSiteUserARMResourceResponseOutput {
	return o
}

func (o StaticSiteUserARMResourceResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserARMResourceResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o StaticSiteUserARMResourceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserARMResourceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o StaticSiteUserARMResourceResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteUserARMResourceResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o StaticSiteUserARMResourceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserARMResourceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o StaticSiteUserARMResourceResponseOutput) Provider() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserARMResourceResponse) string { return v.Provider }).(pulumi.StringOutput)
}

func (o StaticSiteUserARMResourceResponseOutput) Roles() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteUserARMResourceResponse) *string { return v.Roles }).(pulumi.StringPtrOutput)
}

func (o StaticSiteUserARMResourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserARMResourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o StaticSiteUserARMResourceResponseOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserARMResourceResponse) string { return v.UserId }).(pulumi.StringOutput)
}

type StaticSiteUserARMResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (StaticSiteUserARMResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StaticSiteUserARMResourceResponse)(nil)).Elem()
}

func (o StaticSiteUserARMResourceResponseArrayOutput) ToStaticSiteUserARMResourceResponseArrayOutput() StaticSiteUserARMResourceResponseArrayOutput {
	return o
}

func (o StaticSiteUserARMResourceResponseArrayOutput) ToStaticSiteUserARMResourceResponseArrayOutputWithContext(ctx context.Context) StaticSiteUserARMResourceResponseArrayOutput {
	return o
}

func (o StaticSiteUserARMResourceResponseArrayOutput) Index(i pulumi.IntInput) StaticSiteUserARMResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StaticSiteUserARMResourceResponse {
		return vs[0].([]StaticSiteUserARMResourceResponse)[vs[1].(int)]
	}).(StaticSiteUserARMResourceResponseOutput)
}

type StaticSiteUserProvidedFunctionAppResponse struct {
	CreatedOn             string  `pulumi:"createdOn"`
	FunctionAppRegion     *string `pulumi:"functionAppRegion"`
	FunctionAppResourceId *string `pulumi:"functionAppResourceId"`
	Id                    string  `pulumi:"id"`
	Kind                  *string `pulumi:"kind"`
	Name                  string  `pulumi:"name"`
	Type                  string  `pulumi:"type"`
}

type StaticSiteUserProvidedFunctionAppResponseOutput struct{ *pulumi.OutputState }

func (StaticSiteUserProvidedFunctionAppResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteUserProvidedFunctionAppResponse)(nil)).Elem()
}

func (o StaticSiteUserProvidedFunctionAppResponseOutput) ToStaticSiteUserProvidedFunctionAppResponseOutput() StaticSiteUserProvidedFunctionAppResponseOutput {
	return o
}

func (o StaticSiteUserProvidedFunctionAppResponseOutput) ToStaticSiteUserProvidedFunctionAppResponseOutputWithContext(ctx context.Context) StaticSiteUserProvidedFunctionAppResponseOutput {
	return o
}

func (o StaticSiteUserProvidedFunctionAppResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserProvidedFunctionAppResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o StaticSiteUserProvidedFunctionAppResponseOutput) FunctionAppRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteUserProvidedFunctionAppResponse) *string { return v.FunctionAppRegion }).(pulumi.StringPtrOutput)
}

func (o StaticSiteUserProvidedFunctionAppResponseOutput) FunctionAppResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteUserProvidedFunctionAppResponse) *string { return v.FunctionAppResourceId }).(pulumi.StringPtrOutput)
}

func (o StaticSiteUserProvidedFunctionAppResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserProvidedFunctionAppResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o StaticSiteUserProvidedFunctionAppResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteUserProvidedFunctionAppResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o StaticSiteUserProvidedFunctionAppResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserProvidedFunctionAppResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o StaticSiteUserProvidedFunctionAppResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserProvidedFunctionAppResponse) string { return v.Type }).(pulumi.StringOutput)
}

type StaticSiteUserProvidedFunctionAppResponseArrayOutput struct{ *pulumi.OutputState }

func (StaticSiteUserProvidedFunctionAppResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StaticSiteUserProvidedFunctionAppResponse)(nil)).Elem()
}

func (o StaticSiteUserProvidedFunctionAppResponseArrayOutput) ToStaticSiteUserProvidedFunctionAppResponseArrayOutput() StaticSiteUserProvidedFunctionAppResponseArrayOutput {
	return o
}

func (o StaticSiteUserProvidedFunctionAppResponseArrayOutput) ToStaticSiteUserProvidedFunctionAppResponseArrayOutputWithContext(ctx context.Context) StaticSiteUserProvidedFunctionAppResponseArrayOutput {
	return o
}

func (o StaticSiteUserProvidedFunctionAppResponseArrayOutput) Index(i pulumi.IntInput) StaticSiteUserProvidedFunctionAppResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StaticSiteUserProvidedFunctionAppResponse {
		return vs[0].([]StaticSiteUserProvidedFunctionAppResponse)[vs[1].(int)]
	}).(StaticSiteUserProvidedFunctionAppResponseOutput)
}

type StatusCodesBasedTrigger struct {
	Count        *int    `pulumi:"count"`
	Path         *string `pulumi:"path"`
	Status       *int    `pulumi:"status"`
	SubStatus    *int    `pulumi:"subStatus"`
	TimeInterval *string `pulumi:"timeInterval"`
	Win32Status  *int    `pulumi:"win32Status"`
}





type StatusCodesBasedTriggerInput interface {
	pulumi.Input

	ToStatusCodesBasedTriggerOutput() StatusCodesBasedTriggerOutput
	ToStatusCodesBasedTriggerOutputWithContext(context.Context) StatusCodesBasedTriggerOutput
}

type StatusCodesBasedTriggerArgs struct {
	Count        pulumi.IntPtrInput    `pulumi:"count"`
	Path         pulumi.StringPtrInput `pulumi:"path"`
	Status       pulumi.IntPtrInput    `pulumi:"status"`
	SubStatus    pulumi.IntPtrInput    `pulumi:"subStatus"`
	TimeInterval pulumi.StringPtrInput `pulumi:"timeInterval"`
	Win32Status  pulumi.IntPtrInput    `pulumi:"win32Status"`
}

func (StatusCodesBasedTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusCodesBasedTrigger)(nil)).Elem()
}

func (i StatusCodesBasedTriggerArgs) ToStatusCodesBasedTriggerOutput() StatusCodesBasedTriggerOutput {
	return i.ToStatusCodesBasedTriggerOutputWithContext(context.Background())
}

func (i StatusCodesBasedTriggerArgs) ToStatusCodesBasedTriggerOutputWithContext(ctx context.Context) StatusCodesBasedTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusCodesBasedTriggerOutput)
}





type StatusCodesBasedTriggerArrayInput interface {
	pulumi.Input

	ToStatusCodesBasedTriggerArrayOutput() StatusCodesBasedTriggerArrayOutput
	ToStatusCodesBasedTriggerArrayOutputWithContext(context.Context) StatusCodesBasedTriggerArrayOutput
}

type StatusCodesBasedTriggerArray []StatusCodesBasedTriggerInput

func (StatusCodesBasedTriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusCodesBasedTrigger)(nil)).Elem()
}

func (i StatusCodesBasedTriggerArray) ToStatusCodesBasedTriggerArrayOutput() StatusCodesBasedTriggerArrayOutput {
	return i.ToStatusCodesBasedTriggerArrayOutputWithContext(context.Background())
}

func (i StatusCodesBasedTriggerArray) ToStatusCodesBasedTriggerArrayOutputWithContext(ctx context.Context) StatusCodesBasedTriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusCodesBasedTriggerArrayOutput)
}

type StatusCodesBasedTriggerOutput struct{ *pulumi.OutputState }

func (StatusCodesBasedTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusCodesBasedTrigger)(nil)).Elem()
}

func (o StatusCodesBasedTriggerOutput) ToStatusCodesBasedTriggerOutput() StatusCodesBasedTriggerOutput {
	return o
}

func (o StatusCodesBasedTriggerOutput) ToStatusCodesBasedTriggerOutputWithContext(ctx context.Context) StatusCodesBasedTriggerOutput {
	return o
}

func (o StatusCodesBasedTriggerOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTrigger) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o StatusCodesBasedTriggerOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTrigger) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o StatusCodesBasedTriggerOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTrigger) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func (o StatusCodesBasedTriggerOutput) SubStatus() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTrigger) *int { return v.SubStatus }).(pulumi.IntPtrOutput)
}

func (o StatusCodesBasedTriggerOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTrigger) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

func (o StatusCodesBasedTriggerOutput) Win32Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTrigger) *int { return v.Win32Status }).(pulumi.IntPtrOutput)
}

type StatusCodesBasedTriggerArrayOutput struct{ *pulumi.OutputState }

func (StatusCodesBasedTriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusCodesBasedTrigger)(nil)).Elem()
}

func (o StatusCodesBasedTriggerArrayOutput) ToStatusCodesBasedTriggerArrayOutput() StatusCodesBasedTriggerArrayOutput {
	return o
}

func (o StatusCodesBasedTriggerArrayOutput) ToStatusCodesBasedTriggerArrayOutputWithContext(ctx context.Context) StatusCodesBasedTriggerArrayOutput {
	return o
}

func (o StatusCodesBasedTriggerArrayOutput) Index(i pulumi.IntInput) StatusCodesBasedTriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StatusCodesBasedTrigger {
		return vs[0].([]StatusCodesBasedTrigger)[vs[1].(int)]
	}).(StatusCodesBasedTriggerOutput)
}

type StatusCodesBasedTriggerResponse struct {
	Count        *int    `pulumi:"count"`
	Path         *string `pulumi:"path"`
	Status       *int    `pulumi:"status"`
	SubStatus    *int    `pulumi:"subStatus"`
	TimeInterval *string `pulumi:"timeInterval"`
	Win32Status  *int    `pulumi:"win32Status"`
}

type StatusCodesBasedTriggerResponseOutput struct{ *pulumi.OutputState }

func (StatusCodesBasedTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusCodesBasedTriggerResponse)(nil)).Elem()
}

func (o StatusCodesBasedTriggerResponseOutput) ToStatusCodesBasedTriggerResponseOutput() StatusCodesBasedTriggerResponseOutput {
	return o
}

func (o StatusCodesBasedTriggerResponseOutput) ToStatusCodesBasedTriggerResponseOutputWithContext(ctx context.Context) StatusCodesBasedTriggerResponseOutput {
	return o
}

func (o StatusCodesBasedTriggerResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTriggerResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o StatusCodesBasedTriggerResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTriggerResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o StatusCodesBasedTriggerResponseOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTriggerResponse) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func (o StatusCodesBasedTriggerResponseOutput) SubStatus() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTriggerResponse) *int { return v.SubStatus }).(pulumi.IntPtrOutput)
}

func (o StatusCodesBasedTriggerResponseOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTriggerResponse) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

func (o StatusCodesBasedTriggerResponseOutput) Win32Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTriggerResponse) *int { return v.Win32Status }).(pulumi.IntPtrOutput)
}

type StatusCodesBasedTriggerResponseArrayOutput struct{ *pulumi.OutputState }

func (StatusCodesBasedTriggerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusCodesBasedTriggerResponse)(nil)).Elem()
}

func (o StatusCodesBasedTriggerResponseArrayOutput) ToStatusCodesBasedTriggerResponseArrayOutput() StatusCodesBasedTriggerResponseArrayOutput {
	return o
}

func (o StatusCodesBasedTriggerResponseArrayOutput) ToStatusCodesBasedTriggerResponseArrayOutputWithContext(ctx context.Context) StatusCodesBasedTriggerResponseArrayOutput {
	return o
}

func (o StatusCodesBasedTriggerResponseArrayOutput) Index(i pulumi.IntInput) StatusCodesBasedTriggerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StatusCodesBasedTriggerResponse {
		return vs[0].([]StatusCodesBasedTriggerResponse)[vs[1].(int)]
	}).(StatusCodesBasedTriggerResponseOutput)
}

type StatusCodesRangeBasedTrigger struct {
	Count        *int    `pulumi:"count"`
	Path         *string `pulumi:"path"`
	StatusCodes  *string `pulumi:"statusCodes"`
	TimeInterval *string `pulumi:"timeInterval"`
}





type StatusCodesRangeBasedTriggerInput interface {
	pulumi.Input

	ToStatusCodesRangeBasedTriggerOutput() StatusCodesRangeBasedTriggerOutput
	ToStatusCodesRangeBasedTriggerOutputWithContext(context.Context) StatusCodesRangeBasedTriggerOutput
}

type StatusCodesRangeBasedTriggerArgs struct {
	Count        pulumi.IntPtrInput    `pulumi:"count"`
	Path         pulumi.StringPtrInput `pulumi:"path"`
	StatusCodes  pulumi.StringPtrInput `pulumi:"statusCodes"`
	TimeInterval pulumi.StringPtrInput `pulumi:"timeInterval"`
}

func (StatusCodesRangeBasedTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusCodesRangeBasedTrigger)(nil)).Elem()
}

func (i StatusCodesRangeBasedTriggerArgs) ToStatusCodesRangeBasedTriggerOutput() StatusCodesRangeBasedTriggerOutput {
	return i.ToStatusCodesRangeBasedTriggerOutputWithContext(context.Background())
}

func (i StatusCodesRangeBasedTriggerArgs) ToStatusCodesRangeBasedTriggerOutputWithContext(ctx context.Context) StatusCodesRangeBasedTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusCodesRangeBasedTriggerOutput)
}





type StatusCodesRangeBasedTriggerArrayInput interface {
	pulumi.Input

	ToStatusCodesRangeBasedTriggerArrayOutput() StatusCodesRangeBasedTriggerArrayOutput
	ToStatusCodesRangeBasedTriggerArrayOutputWithContext(context.Context) StatusCodesRangeBasedTriggerArrayOutput
}

type StatusCodesRangeBasedTriggerArray []StatusCodesRangeBasedTriggerInput

func (StatusCodesRangeBasedTriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusCodesRangeBasedTrigger)(nil)).Elem()
}

func (i StatusCodesRangeBasedTriggerArray) ToStatusCodesRangeBasedTriggerArrayOutput() StatusCodesRangeBasedTriggerArrayOutput {
	return i.ToStatusCodesRangeBasedTriggerArrayOutputWithContext(context.Background())
}

func (i StatusCodesRangeBasedTriggerArray) ToStatusCodesRangeBasedTriggerArrayOutputWithContext(ctx context.Context) StatusCodesRangeBasedTriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusCodesRangeBasedTriggerArrayOutput)
}

type StatusCodesRangeBasedTriggerOutput struct{ *pulumi.OutputState }

func (StatusCodesRangeBasedTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusCodesRangeBasedTrigger)(nil)).Elem()
}

func (o StatusCodesRangeBasedTriggerOutput) ToStatusCodesRangeBasedTriggerOutput() StatusCodesRangeBasedTriggerOutput {
	return o
}

func (o StatusCodesRangeBasedTriggerOutput) ToStatusCodesRangeBasedTriggerOutputWithContext(ctx context.Context) StatusCodesRangeBasedTriggerOutput {
	return o
}

func (o StatusCodesRangeBasedTriggerOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesRangeBasedTrigger) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o StatusCodesRangeBasedTriggerOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesRangeBasedTrigger) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o StatusCodesRangeBasedTriggerOutput) StatusCodes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesRangeBasedTrigger) *string { return v.StatusCodes }).(pulumi.StringPtrOutput)
}

func (o StatusCodesRangeBasedTriggerOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesRangeBasedTrigger) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

type StatusCodesRangeBasedTriggerArrayOutput struct{ *pulumi.OutputState }

func (StatusCodesRangeBasedTriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusCodesRangeBasedTrigger)(nil)).Elem()
}

func (o StatusCodesRangeBasedTriggerArrayOutput) ToStatusCodesRangeBasedTriggerArrayOutput() StatusCodesRangeBasedTriggerArrayOutput {
	return o
}

func (o StatusCodesRangeBasedTriggerArrayOutput) ToStatusCodesRangeBasedTriggerArrayOutputWithContext(ctx context.Context) StatusCodesRangeBasedTriggerArrayOutput {
	return o
}

func (o StatusCodesRangeBasedTriggerArrayOutput) Index(i pulumi.IntInput) StatusCodesRangeBasedTriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StatusCodesRangeBasedTrigger {
		return vs[0].([]StatusCodesRangeBasedTrigger)[vs[1].(int)]
	}).(StatusCodesRangeBasedTriggerOutput)
}

type StatusCodesRangeBasedTriggerResponse struct {
	Count        *int    `pulumi:"count"`
	Path         *string `pulumi:"path"`
	StatusCodes  *string `pulumi:"statusCodes"`
	TimeInterval *string `pulumi:"timeInterval"`
}

type StatusCodesRangeBasedTriggerResponseOutput struct{ *pulumi.OutputState }

func (StatusCodesRangeBasedTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusCodesRangeBasedTriggerResponse)(nil)).Elem()
}

func (o StatusCodesRangeBasedTriggerResponseOutput) ToStatusCodesRangeBasedTriggerResponseOutput() StatusCodesRangeBasedTriggerResponseOutput {
	return o
}

func (o StatusCodesRangeBasedTriggerResponseOutput) ToStatusCodesRangeBasedTriggerResponseOutputWithContext(ctx context.Context) StatusCodesRangeBasedTriggerResponseOutput {
	return o
}

func (o StatusCodesRangeBasedTriggerResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesRangeBasedTriggerResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o StatusCodesRangeBasedTriggerResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesRangeBasedTriggerResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o StatusCodesRangeBasedTriggerResponseOutput) StatusCodes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesRangeBasedTriggerResponse) *string { return v.StatusCodes }).(pulumi.StringPtrOutput)
}

func (o StatusCodesRangeBasedTriggerResponseOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesRangeBasedTriggerResponse) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

type StatusCodesRangeBasedTriggerResponseArrayOutput struct{ *pulumi.OutputState }

func (StatusCodesRangeBasedTriggerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusCodesRangeBasedTriggerResponse)(nil)).Elem()
}

func (o StatusCodesRangeBasedTriggerResponseArrayOutput) ToStatusCodesRangeBasedTriggerResponseArrayOutput() StatusCodesRangeBasedTriggerResponseArrayOutput {
	return o
}

func (o StatusCodesRangeBasedTriggerResponseArrayOutput) ToStatusCodesRangeBasedTriggerResponseArrayOutputWithContext(ctx context.Context) StatusCodesRangeBasedTriggerResponseArrayOutput {
	return o
}

func (o StatusCodesRangeBasedTriggerResponseArrayOutput) Index(i pulumi.IntInput) StatusCodesRangeBasedTriggerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StatusCodesRangeBasedTriggerResponse {
		return vs[0].([]StatusCodesRangeBasedTriggerResponse)[vs[1].(int)]
	}).(StatusCodesRangeBasedTriggerResponseOutput)
}

type TokenStore struct {
	AzureBlobStorage           *BlobStorageTokenStore `pulumi:"azureBlobStorage"`
	Enabled                    *bool                  `pulumi:"enabled"`
	FileSystem                 *FileSystemTokenStore  `pulumi:"fileSystem"`
	TokenRefreshExtensionHours *float64               `pulumi:"tokenRefreshExtensionHours"`
}





type TokenStoreInput interface {
	pulumi.Input

	ToTokenStoreOutput() TokenStoreOutput
	ToTokenStoreOutputWithContext(context.Context) TokenStoreOutput
}

type TokenStoreArgs struct {
	AzureBlobStorage           BlobStorageTokenStorePtrInput `pulumi:"azureBlobStorage"`
	Enabled                    pulumi.BoolPtrInput           `pulumi:"enabled"`
	FileSystem                 FileSystemTokenStorePtrInput  `pulumi:"fileSystem"`
	TokenRefreshExtensionHours pulumi.Float64PtrInput        `pulumi:"tokenRefreshExtensionHours"`
}

func (TokenStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenStore)(nil)).Elem()
}

func (i TokenStoreArgs) ToTokenStoreOutput() TokenStoreOutput {
	return i.ToTokenStoreOutputWithContext(context.Background())
}

func (i TokenStoreArgs) ToTokenStoreOutputWithContext(ctx context.Context) TokenStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenStoreOutput)
}

func (i TokenStoreArgs) ToTokenStorePtrOutput() TokenStorePtrOutput {
	return i.ToTokenStorePtrOutputWithContext(context.Background())
}

func (i TokenStoreArgs) ToTokenStorePtrOutputWithContext(ctx context.Context) TokenStorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenStoreOutput).ToTokenStorePtrOutputWithContext(ctx)
}









type TokenStorePtrInput interface {
	pulumi.Input

	ToTokenStorePtrOutput() TokenStorePtrOutput
	ToTokenStorePtrOutputWithContext(context.Context) TokenStorePtrOutput
}

type tokenStorePtrType TokenStoreArgs

func TokenStorePtr(v *TokenStoreArgs) TokenStorePtrInput {
	return (*tokenStorePtrType)(v)
}

func (*tokenStorePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenStore)(nil)).Elem()
}

func (i *tokenStorePtrType) ToTokenStorePtrOutput() TokenStorePtrOutput {
	return i.ToTokenStorePtrOutputWithContext(context.Background())
}

func (i *tokenStorePtrType) ToTokenStorePtrOutputWithContext(ctx context.Context) TokenStorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenStorePtrOutput)
}

type TokenStoreOutput struct{ *pulumi.OutputState }

func (TokenStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenStore)(nil)).Elem()
}

func (o TokenStoreOutput) ToTokenStoreOutput() TokenStoreOutput {
	return o
}

func (o TokenStoreOutput) ToTokenStoreOutputWithContext(ctx context.Context) TokenStoreOutput {
	return o
}

func (o TokenStoreOutput) ToTokenStorePtrOutput() TokenStorePtrOutput {
	return o.ToTokenStorePtrOutputWithContext(context.Background())
}

func (o TokenStoreOutput) ToTokenStorePtrOutputWithContext(ctx context.Context) TokenStorePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TokenStore) *TokenStore {
		return &v
	}).(TokenStorePtrOutput)
}

func (o TokenStoreOutput) AzureBlobStorage() BlobStorageTokenStorePtrOutput {
	return o.ApplyT(func(v TokenStore) *BlobStorageTokenStore { return v.AzureBlobStorage }).(BlobStorageTokenStorePtrOutput)
}

func (o TokenStoreOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TokenStore) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o TokenStoreOutput) FileSystem() FileSystemTokenStorePtrOutput {
	return o.ApplyT(func(v TokenStore) *FileSystemTokenStore { return v.FileSystem }).(FileSystemTokenStorePtrOutput)
}

func (o TokenStoreOutput) TokenRefreshExtensionHours() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v TokenStore) *float64 { return v.TokenRefreshExtensionHours }).(pulumi.Float64PtrOutput)
}

type TokenStorePtrOutput struct{ *pulumi.OutputState }

func (TokenStorePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenStore)(nil)).Elem()
}

func (o TokenStorePtrOutput) ToTokenStorePtrOutput() TokenStorePtrOutput {
	return o
}

func (o TokenStorePtrOutput) ToTokenStorePtrOutputWithContext(ctx context.Context) TokenStorePtrOutput {
	return o
}

func (o TokenStorePtrOutput) Elem() TokenStoreOutput {
	return o.ApplyT(func(v *TokenStore) TokenStore {
		if v != nil {
			return *v
		}
		var ret TokenStore
		return ret
	}).(TokenStoreOutput)
}

func (o TokenStorePtrOutput) AzureBlobStorage() BlobStorageTokenStorePtrOutput {
	return o.ApplyT(func(v *TokenStore) *BlobStorageTokenStore {
		if v == nil {
			return nil
		}
		return v.AzureBlobStorage
	}).(BlobStorageTokenStorePtrOutput)
}

func (o TokenStorePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TokenStore) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o TokenStorePtrOutput) FileSystem() FileSystemTokenStorePtrOutput {
	return o.ApplyT(func(v *TokenStore) *FileSystemTokenStore {
		if v == nil {
			return nil
		}
		return v.FileSystem
	}).(FileSystemTokenStorePtrOutput)
}

func (o TokenStorePtrOutput) TokenRefreshExtensionHours() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *TokenStore) *float64 {
		if v == nil {
			return nil
		}
		return v.TokenRefreshExtensionHours
	}).(pulumi.Float64PtrOutput)
}

type TokenStoreResponse struct {
	AzureBlobStorage           *BlobStorageTokenStoreResponse `pulumi:"azureBlobStorage"`
	Enabled                    *bool                          `pulumi:"enabled"`
	FileSystem                 *FileSystemTokenStoreResponse  `pulumi:"fileSystem"`
	TokenRefreshExtensionHours *float64                       `pulumi:"tokenRefreshExtensionHours"`
}

type TokenStoreResponseOutput struct{ *pulumi.OutputState }

func (TokenStoreResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenStoreResponse)(nil)).Elem()
}

func (o TokenStoreResponseOutput) ToTokenStoreResponseOutput() TokenStoreResponseOutput {
	return o
}

func (o TokenStoreResponseOutput) ToTokenStoreResponseOutputWithContext(ctx context.Context) TokenStoreResponseOutput {
	return o
}

func (o TokenStoreResponseOutput) AzureBlobStorage() BlobStorageTokenStoreResponsePtrOutput {
	return o.ApplyT(func(v TokenStoreResponse) *BlobStorageTokenStoreResponse { return v.AzureBlobStorage }).(BlobStorageTokenStoreResponsePtrOutput)
}

func (o TokenStoreResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TokenStoreResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o TokenStoreResponseOutput) FileSystem() FileSystemTokenStoreResponsePtrOutput {
	return o.ApplyT(func(v TokenStoreResponse) *FileSystemTokenStoreResponse { return v.FileSystem }).(FileSystemTokenStoreResponsePtrOutput)
}

func (o TokenStoreResponseOutput) TokenRefreshExtensionHours() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v TokenStoreResponse) *float64 { return v.TokenRefreshExtensionHours }).(pulumi.Float64PtrOutput)
}

type TokenStoreResponsePtrOutput struct{ *pulumi.OutputState }

func (TokenStoreResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenStoreResponse)(nil)).Elem()
}

func (o TokenStoreResponsePtrOutput) ToTokenStoreResponsePtrOutput() TokenStoreResponsePtrOutput {
	return o
}

func (o TokenStoreResponsePtrOutput) ToTokenStoreResponsePtrOutputWithContext(ctx context.Context) TokenStoreResponsePtrOutput {
	return o
}

func (o TokenStoreResponsePtrOutput) Elem() TokenStoreResponseOutput {
	return o.ApplyT(func(v *TokenStoreResponse) TokenStoreResponse {
		if v != nil {
			return *v
		}
		var ret TokenStoreResponse
		return ret
	}).(TokenStoreResponseOutput)
}

func (o TokenStoreResponsePtrOutput) AzureBlobStorage() BlobStorageTokenStoreResponsePtrOutput {
	return o.ApplyT(func(v *TokenStoreResponse) *BlobStorageTokenStoreResponse {
		if v == nil {
			return nil
		}
		return v.AzureBlobStorage
	}).(BlobStorageTokenStoreResponsePtrOutput)
}

func (o TokenStoreResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TokenStoreResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o TokenStoreResponsePtrOutput) FileSystem() FileSystemTokenStoreResponsePtrOutput {
	return o.ApplyT(func(v *TokenStoreResponse) *FileSystemTokenStoreResponse {
		if v == nil {
			return nil
		}
		return v.FileSystem
	}).(FileSystemTokenStoreResponsePtrOutput)
}

func (o TokenStoreResponsePtrOutput) TokenRefreshExtensionHours() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *TokenStoreResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.TokenRefreshExtensionHours
	}).(pulumi.Float64PtrOutput)
}

type Twitter struct {
	Enabled      *bool                `pulumi:"enabled"`
	Registration *TwitterRegistration `pulumi:"registration"`
}





type TwitterInput interface {
	pulumi.Input

	ToTwitterOutput() TwitterOutput
	ToTwitterOutputWithContext(context.Context) TwitterOutput
}

type TwitterArgs struct {
	Enabled      pulumi.BoolPtrInput         `pulumi:"enabled"`
	Registration TwitterRegistrationPtrInput `pulumi:"registration"`
}

func (TwitterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Twitter)(nil)).Elem()
}

func (i TwitterArgs) ToTwitterOutput() TwitterOutput {
	return i.ToTwitterOutputWithContext(context.Background())
}

func (i TwitterArgs) ToTwitterOutputWithContext(ctx context.Context) TwitterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwitterOutput)
}

func (i TwitterArgs) ToTwitterPtrOutput() TwitterPtrOutput {
	return i.ToTwitterPtrOutputWithContext(context.Background())
}

func (i TwitterArgs) ToTwitterPtrOutputWithContext(ctx context.Context) TwitterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwitterOutput).ToTwitterPtrOutputWithContext(ctx)
}









type TwitterPtrInput interface {
	pulumi.Input

	ToTwitterPtrOutput() TwitterPtrOutput
	ToTwitterPtrOutputWithContext(context.Context) TwitterPtrOutput
}

type twitterPtrType TwitterArgs

func TwitterPtr(v *TwitterArgs) TwitterPtrInput {
	return (*twitterPtrType)(v)
}

func (*twitterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Twitter)(nil)).Elem()
}

func (i *twitterPtrType) ToTwitterPtrOutput() TwitterPtrOutput {
	return i.ToTwitterPtrOutputWithContext(context.Background())
}

func (i *twitterPtrType) ToTwitterPtrOutputWithContext(ctx context.Context) TwitterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwitterPtrOutput)
}

type TwitterOutput struct{ *pulumi.OutputState }

func (TwitterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Twitter)(nil)).Elem()
}

func (o TwitterOutput) ToTwitterOutput() TwitterOutput {
	return o
}

func (o TwitterOutput) ToTwitterOutputWithContext(ctx context.Context) TwitterOutput {
	return o
}

func (o TwitterOutput) ToTwitterPtrOutput() TwitterPtrOutput {
	return o.ToTwitterPtrOutputWithContext(context.Background())
}

func (o TwitterOutput) ToTwitterPtrOutputWithContext(ctx context.Context) TwitterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Twitter) *Twitter {
		return &v
	}).(TwitterPtrOutput)
}

func (o TwitterOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Twitter) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o TwitterOutput) Registration() TwitterRegistrationPtrOutput {
	return o.ApplyT(func(v Twitter) *TwitterRegistration { return v.Registration }).(TwitterRegistrationPtrOutput)
}

type TwitterPtrOutput struct{ *pulumi.OutputState }

func (TwitterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Twitter)(nil)).Elem()
}

func (o TwitterPtrOutput) ToTwitterPtrOutput() TwitterPtrOutput {
	return o
}

func (o TwitterPtrOutput) ToTwitterPtrOutputWithContext(ctx context.Context) TwitterPtrOutput {
	return o
}

func (o TwitterPtrOutput) Elem() TwitterOutput {
	return o.ApplyT(func(v *Twitter) Twitter {
		if v != nil {
			return *v
		}
		var ret Twitter
		return ret
	}).(TwitterOutput)
}

func (o TwitterPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Twitter) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o TwitterPtrOutput) Registration() TwitterRegistrationPtrOutput {
	return o.ApplyT(func(v *Twitter) *TwitterRegistration {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(TwitterRegistrationPtrOutput)
}

type TwitterRegistration struct {
	ConsumerKey               *string `pulumi:"consumerKey"`
	ConsumerSecretSettingName *string `pulumi:"consumerSecretSettingName"`
}





type TwitterRegistrationInput interface {
	pulumi.Input

	ToTwitterRegistrationOutput() TwitterRegistrationOutput
	ToTwitterRegistrationOutputWithContext(context.Context) TwitterRegistrationOutput
}

type TwitterRegistrationArgs struct {
	ConsumerKey               pulumi.StringPtrInput `pulumi:"consumerKey"`
	ConsumerSecretSettingName pulumi.StringPtrInput `pulumi:"consumerSecretSettingName"`
}

func (TwitterRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TwitterRegistration)(nil)).Elem()
}

func (i TwitterRegistrationArgs) ToTwitterRegistrationOutput() TwitterRegistrationOutput {
	return i.ToTwitterRegistrationOutputWithContext(context.Background())
}

func (i TwitterRegistrationArgs) ToTwitterRegistrationOutputWithContext(ctx context.Context) TwitterRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwitterRegistrationOutput)
}

func (i TwitterRegistrationArgs) ToTwitterRegistrationPtrOutput() TwitterRegistrationPtrOutput {
	return i.ToTwitterRegistrationPtrOutputWithContext(context.Background())
}

func (i TwitterRegistrationArgs) ToTwitterRegistrationPtrOutputWithContext(ctx context.Context) TwitterRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwitterRegistrationOutput).ToTwitterRegistrationPtrOutputWithContext(ctx)
}









type TwitterRegistrationPtrInput interface {
	pulumi.Input

	ToTwitterRegistrationPtrOutput() TwitterRegistrationPtrOutput
	ToTwitterRegistrationPtrOutputWithContext(context.Context) TwitterRegistrationPtrOutput
}

type twitterRegistrationPtrType TwitterRegistrationArgs

func TwitterRegistrationPtr(v *TwitterRegistrationArgs) TwitterRegistrationPtrInput {
	return (*twitterRegistrationPtrType)(v)
}

func (*twitterRegistrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TwitterRegistration)(nil)).Elem()
}

func (i *twitterRegistrationPtrType) ToTwitterRegistrationPtrOutput() TwitterRegistrationPtrOutput {
	return i.ToTwitterRegistrationPtrOutputWithContext(context.Background())
}

func (i *twitterRegistrationPtrType) ToTwitterRegistrationPtrOutputWithContext(ctx context.Context) TwitterRegistrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TwitterRegistrationPtrOutput)
}

type TwitterRegistrationOutput struct{ *pulumi.OutputState }

func (TwitterRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TwitterRegistration)(nil)).Elem()
}

func (o TwitterRegistrationOutput) ToTwitterRegistrationOutput() TwitterRegistrationOutput {
	return o
}

func (o TwitterRegistrationOutput) ToTwitterRegistrationOutputWithContext(ctx context.Context) TwitterRegistrationOutput {
	return o
}

func (o TwitterRegistrationOutput) ToTwitterRegistrationPtrOutput() TwitterRegistrationPtrOutput {
	return o.ToTwitterRegistrationPtrOutputWithContext(context.Background())
}

func (o TwitterRegistrationOutput) ToTwitterRegistrationPtrOutputWithContext(ctx context.Context) TwitterRegistrationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TwitterRegistration) *TwitterRegistration {
		return &v
	}).(TwitterRegistrationPtrOutput)
}

func (o TwitterRegistrationOutput) ConsumerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TwitterRegistration) *string { return v.ConsumerKey }).(pulumi.StringPtrOutput)
}

func (o TwitterRegistrationOutput) ConsumerSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TwitterRegistration) *string { return v.ConsumerSecretSettingName }).(pulumi.StringPtrOutput)
}

type TwitterRegistrationPtrOutput struct{ *pulumi.OutputState }

func (TwitterRegistrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TwitterRegistration)(nil)).Elem()
}

func (o TwitterRegistrationPtrOutput) ToTwitterRegistrationPtrOutput() TwitterRegistrationPtrOutput {
	return o
}

func (o TwitterRegistrationPtrOutput) ToTwitterRegistrationPtrOutputWithContext(ctx context.Context) TwitterRegistrationPtrOutput {
	return o
}

func (o TwitterRegistrationPtrOutput) Elem() TwitterRegistrationOutput {
	return o.ApplyT(func(v *TwitterRegistration) TwitterRegistration {
		if v != nil {
			return *v
		}
		var ret TwitterRegistration
		return ret
	}).(TwitterRegistrationOutput)
}

func (o TwitterRegistrationPtrOutput) ConsumerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TwitterRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ConsumerKey
	}).(pulumi.StringPtrOutput)
}

func (o TwitterRegistrationPtrOutput) ConsumerSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TwitterRegistration) *string {
		if v == nil {
			return nil
		}
		return v.ConsumerSecretSettingName
	}).(pulumi.StringPtrOutput)
}

type TwitterRegistrationResponse struct {
	ConsumerKey               *string `pulumi:"consumerKey"`
	ConsumerSecretSettingName *string `pulumi:"consumerSecretSettingName"`
}

type TwitterRegistrationResponseOutput struct{ *pulumi.OutputState }

func (TwitterRegistrationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TwitterRegistrationResponse)(nil)).Elem()
}

func (o TwitterRegistrationResponseOutput) ToTwitterRegistrationResponseOutput() TwitterRegistrationResponseOutput {
	return o
}

func (o TwitterRegistrationResponseOutput) ToTwitterRegistrationResponseOutputWithContext(ctx context.Context) TwitterRegistrationResponseOutput {
	return o
}

func (o TwitterRegistrationResponseOutput) ConsumerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TwitterRegistrationResponse) *string { return v.ConsumerKey }).(pulumi.StringPtrOutput)
}

func (o TwitterRegistrationResponseOutput) ConsumerSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TwitterRegistrationResponse) *string { return v.ConsumerSecretSettingName }).(pulumi.StringPtrOutput)
}

type TwitterRegistrationResponsePtrOutput struct{ *pulumi.OutputState }

func (TwitterRegistrationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TwitterRegistrationResponse)(nil)).Elem()
}

func (o TwitterRegistrationResponsePtrOutput) ToTwitterRegistrationResponsePtrOutput() TwitterRegistrationResponsePtrOutput {
	return o
}

func (o TwitterRegistrationResponsePtrOutput) ToTwitterRegistrationResponsePtrOutputWithContext(ctx context.Context) TwitterRegistrationResponsePtrOutput {
	return o
}

func (o TwitterRegistrationResponsePtrOutput) Elem() TwitterRegistrationResponseOutput {
	return o.ApplyT(func(v *TwitterRegistrationResponse) TwitterRegistrationResponse {
		if v != nil {
			return *v
		}
		var ret TwitterRegistrationResponse
		return ret
	}).(TwitterRegistrationResponseOutput)
}

func (o TwitterRegistrationResponsePtrOutput) ConsumerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TwitterRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConsumerKey
	}).(pulumi.StringPtrOutput)
}

func (o TwitterRegistrationResponsePtrOutput) ConsumerSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TwitterRegistrationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConsumerSecretSettingName
	}).(pulumi.StringPtrOutput)
}

type TwitterResponse struct {
	Enabled      *bool                        `pulumi:"enabled"`
	Registration *TwitterRegistrationResponse `pulumi:"registration"`
}

type TwitterResponseOutput struct{ *pulumi.OutputState }

func (TwitterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TwitterResponse)(nil)).Elem()
}

func (o TwitterResponseOutput) ToTwitterResponseOutput() TwitterResponseOutput {
	return o
}

func (o TwitterResponseOutput) ToTwitterResponseOutputWithContext(ctx context.Context) TwitterResponseOutput {
	return o
}

func (o TwitterResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TwitterResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o TwitterResponseOutput) Registration() TwitterRegistrationResponsePtrOutput {
	return o.ApplyT(func(v TwitterResponse) *TwitterRegistrationResponse { return v.Registration }).(TwitterRegistrationResponsePtrOutput)
}

type TwitterResponsePtrOutput struct{ *pulumi.OutputState }

func (TwitterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TwitterResponse)(nil)).Elem()
}

func (o TwitterResponsePtrOutput) ToTwitterResponsePtrOutput() TwitterResponsePtrOutput {
	return o
}

func (o TwitterResponsePtrOutput) ToTwitterResponsePtrOutputWithContext(ctx context.Context) TwitterResponsePtrOutput {
	return o
}

func (o TwitterResponsePtrOutput) Elem() TwitterResponseOutput {
	return o.ApplyT(func(v *TwitterResponse) TwitterResponse {
		if v != nil {
			return *v
		}
		var ret TwitterResponse
		return ret
	}).(TwitterResponseOutput)
}

func (o TwitterResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TwitterResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o TwitterResponsePtrOutput) Registration() TwitterRegistrationResponsePtrOutput {
	return o.ApplyT(func(v *TwitterResponse) *TwitterRegistrationResponse {
		if v == nil {
			return nil
		}
		return v.Registration
	}).(TwitterRegistrationResponsePtrOutput)
}

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

type VirtualApplication struct {
	PhysicalPath       *string            `pulumi:"physicalPath"`
	PreloadEnabled     *bool              `pulumi:"preloadEnabled"`
	VirtualDirectories []VirtualDirectory `pulumi:"virtualDirectories"`
	VirtualPath        *string            `pulumi:"virtualPath"`
}





type VirtualApplicationInput interface {
	pulumi.Input

	ToVirtualApplicationOutput() VirtualApplicationOutput
	ToVirtualApplicationOutputWithContext(context.Context) VirtualApplicationOutput
}

type VirtualApplicationArgs struct {
	PhysicalPath       pulumi.StringPtrInput      `pulumi:"physicalPath"`
	PreloadEnabled     pulumi.BoolPtrInput        `pulumi:"preloadEnabled"`
	VirtualDirectories VirtualDirectoryArrayInput `pulumi:"virtualDirectories"`
	VirtualPath        pulumi.StringPtrInput      `pulumi:"virtualPath"`
}

func (VirtualApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualApplication)(nil)).Elem()
}

func (i VirtualApplicationArgs) ToVirtualApplicationOutput() VirtualApplicationOutput {
	return i.ToVirtualApplicationOutputWithContext(context.Background())
}

func (i VirtualApplicationArgs) ToVirtualApplicationOutputWithContext(ctx context.Context) VirtualApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualApplicationOutput)
}





type VirtualApplicationArrayInput interface {
	pulumi.Input

	ToVirtualApplicationArrayOutput() VirtualApplicationArrayOutput
	ToVirtualApplicationArrayOutputWithContext(context.Context) VirtualApplicationArrayOutput
}

type VirtualApplicationArray []VirtualApplicationInput

func (VirtualApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualApplication)(nil)).Elem()
}

func (i VirtualApplicationArray) ToVirtualApplicationArrayOutput() VirtualApplicationArrayOutput {
	return i.ToVirtualApplicationArrayOutputWithContext(context.Background())
}

func (i VirtualApplicationArray) ToVirtualApplicationArrayOutputWithContext(ctx context.Context) VirtualApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualApplicationArrayOutput)
}

type VirtualApplicationOutput struct{ *pulumi.OutputState }

func (VirtualApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualApplication)(nil)).Elem()
}

func (o VirtualApplicationOutput) ToVirtualApplicationOutput() VirtualApplicationOutput {
	return o
}

func (o VirtualApplicationOutput) ToVirtualApplicationOutputWithContext(ctx context.Context) VirtualApplicationOutput {
	return o
}

func (o VirtualApplicationOutput) PhysicalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualApplication) *string { return v.PhysicalPath }).(pulumi.StringPtrOutput)
}

func (o VirtualApplicationOutput) PreloadEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualApplication) *bool { return v.PreloadEnabled }).(pulumi.BoolPtrOutput)
}

func (o VirtualApplicationOutput) VirtualDirectories() VirtualDirectoryArrayOutput {
	return o.ApplyT(func(v VirtualApplication) []VirtualDirectory { return v.VirtualDirectories }).(VirtualDirectoryArrayOutput)
}

func (o VirtualApplicationOutput) VirtualPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualApplication) *string { return v.VirtualPath }).(pulumi.StringPtrOutput)
}

type VirtualApplicationArrayOutput struct{ *pulumi.OutputState }

func (VirtualApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualApplication)(nil)).Elem()
}

func (o VirtualApplicationArrayOutput) ToVirtualApplicationArrayOutput() VirtualApplicationArrayOutput {
	return o
}

func (o VirtualApplicationArrayOutput) ToVirtualApplicationArrayOutputWithContext(ctx context.Context) VirtualApplicationArrayOutput {
	return o
}

func (o VirtualApplicationArrayOutput) Index(i pulumi.IntInput) VirtualApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualApplication {
		return vs[0].([]VirtualApplication)[vs[1].(int)]
	}).(VirtualApplicationOutput)
}

type VirtualApplicationResponse struct {
	PhysicalPath       *string                    `pulumi:"physicalPath"`
	PreloadEnabled     *bool                      `pulumi:"preloadEnabled"`
	VirtualDirectories []VirtualDirectoryResponse `pulumi:"virtualDirectories"`
	VirtualPath        *string                    `pulumi:"virtualPath"`
}

type VirtualApplicationResponseOutput struct{ *pulumi.OutputState }

func (VirtualApplicationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualApplicationResponse)(nil)).Elem()
}

func (o VirtualApplicationResponseOutput) ToVirtualApplicationResponseOutput() VirtualApplicationResponseOutput {
	return o
}

func (o VirtualApplicationResponseOutput) ToVirtualApplicationResponseOutputWithContext(ctx context.Context) VirtualApplicationResponseOutput {
	return o
}

func (o VirtualApplicationResponseOutput) PhysicalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualApplicationResponse) *string { return v.PhysicalPath }).(pulumi.StringPtrOutput)
}

func (o VirtualApplicationResponseOutput) PreloadEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualApplicationResponse) *bool { return v.PreloadEnabled }).(pulumi.BoolPtrOutput)
}

func (o VirtualApplicationResponseOutput) VirtualDirectories() VirtualDirectoryResponseArrayOutput {
	return o.ApplyT(func(v VirtualApplicationResponse) []VirtualDirectoryResponse { return v.VirtualDirectories }).(VirtualDirectoryResponseArrayOutput)
}

func (o VirtualApplicationResponseOutput) VirtualPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualApplicationResponse) *string { return v.VirtualPath }).(pulumi.StringPtrOutput)
}

type VirtualApplicationResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualApplicationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualApplicationResponse)(nil)).Elem()
}

func (o VirtualApplicationResponseArrayOutput) ToVirtualApplicationResponseArrayOutput() VirtualApplicationResponseArrayOutput {
	return o
}

func (o VirtualApplicationResponseArrayOutput) ToVirtualApplicationResponseArrayOutputWithContext(ctx context.Context) VirtualApplicationResponseArrayOutput {
	return o
}

func (o VirtualApplicationResponseArrayOutput) Index(i pulumi.IntInput) VirtualApplicationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualApplicationResponse {
		return vs[0].([]VirtualApplicationResponse)[vs[1].(int)]
	}).(VirtualApplicationResponseOutput)
}

type VirtualDirectory struct {
	PhysicalPath *string `pulumi:"physicalPath"`
	VirtualPath  *string `pulumi:"virtualPath"`
}





type VirtualDirectoryInput interface {
	pulumi.Input

	ToVirtualDirectoryOutput() VirtualDirectoryOutput
	ToVirtualDirectoryOutputWithContext(context.Context) VirtualDirectoryOutput
}

type VirtualDirectoryArgs struct {
	PhysicalPath pulumi.StringPtrInput `pulumi:"physicalPath"`
	VirtualPath  pulumi.StringPtrInput `pulumi:"virtualPath"`
}

func (VirtualDirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDirectory)(nil)).Elem()
}

func (i VirtualDirectoryArgs) ToVirtualDirectoryOutput() VirtualDirectoryOutput {
	return i.ToVirtualDirectoryOutputWithContext(context.Background())
}

func (i VirtualDirectoryArgs) ToVirtualDirectoryOutputWithContext(ctx context.Context) VirtualDirectoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDirectoryOutput)
}





type VirtualDirectoryArrayInput interface {
	pulumi.Input

	ToVirtualDirectoryArrayOutput() VirtualDirectoryArrayOutput
	ToVirtualDirectoryArrayOutputWithContext(context.Context) VirtualDirectoryArrayOutput
}

type VirtualDirectoryArray []VirtualDirectoryInput

func (VirtualDirectoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDirectory)(nil)).Elem()
}

func (i VirtualDirectoryArray) ToVirtualDirectoryArrayOutput() VirtualDirectoryArrayOutput {
	return i.ToVirtualDirectoryArrayOutputWithContext(context.Background())
}

func (i VirtualDirectoryArray) ToVirtualDirectoryArrayOutputWithContext(ctx context.Context) VirtualDirectoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDirectoryArrayOutput)
}

type VirtualDirectoryOutput struct{ *pulumi.OutputState }

func (VirtualDirectoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDirectory)(nil)).Elem()
}

func (o VirtualDirectoryOutput) ToVirtualDirectoryOutput() VirtualDirectoryOutput {
	return o
}

func (o VirtualDirectoryOutput) ToVirtualDirectoryOutputWithContext(ctx context.Context) VirtualDirectoryOutput {
	return o
}

func (o VirtualDirectoryOutput) PhysicalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDirectory) *string { return v.PhysicalPath }).(pulumi.StringPtrOutput)
}

func (o VirtualDirectoryOutput) VirtualPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDirectory) *string { return v.VirtualPath }).(pulumi.StringPtrOutput)
}

type VirtualDirectoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualDirectoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDirectory)(nil)).Elem()
}

func (o VirtualDirectoryArrayOutput) ToVirtualDirectoryArrayOutput() VirtualDirectoryArrayOutput {
	return o
}

func (o VirtualDirectoryArrayOutput) ToVirtualDirectoryArrayOutputWithContext(ctx context.Context) VirtualDirectoryArrayOutput {
	return o
}

func (o VirtualDirectoryArrayOutput) Index(i pulumi.IntInput) VirtualDirectoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualDirectory {
		return vs[0].([]VirtualDirectory)[vs[1].(int)]
	}).(VirtualDirectoryOutput)
}

type VirtualDirectoryResponse struct {
	PhysicalPath *string `pulumi:"physicalPath"`
	VirtualPath  *string `pulumi:"virtualPath"`
}

type VirtualDirectoryResponseOutput struct{ *pulumi.OutputState }

func (VirtualDirectoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDirectoryResponse)(nil)).Elem()
}

func (o VirtualDirectoryResponseOutput) ToVirtualDirectoryResponseOutput() VirtualDirectoryResponseOutput {
	return o
}

func (o VirtualDirectoryResponseOutput) ToVirtualDirectoryResponseOutputWithContext(ctx context.Context) VirtualDirectoryResponseOutput {
	return o
}

func (o VirtualDirectoryResponseOutput) PhysicalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDirectoryResponse) *string { return v.PhysicalPath }).(pulumi.StringPtrOutput)
}

func (o VirtualDirectoryResponseOutput) VirtualPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDirectoryResponse) *string { return v.VirtualPath }).(pulumi.StringPtrOutput)
}

type VirtualDirectoryResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualDirectoryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDirectoryResponse)(nil)).Elem()
}

func (o VirtualDirectoryResponseArrayOutput) ToVirtualDirectoryResponseArrayOutput() VirtualDirectoryResponseArrayOutput {
	return o
}

func (o VirtualDirectoryResponseArrayOutput) ToVirtualDirectoryResponseArrayOutputWithContext(ctx context.Context) VirtualDirectoryResponseArrayOutput {
	return o
}

func (o VirtualDirectoryResponseArrayOutput) Index(i pulumi.IntInput) VirtualDirectoryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualDirectoryResponse {
		return vs[0].([]VirtualDirectoryResponse)[vs[1].(int)]
	}).(VirtualDirectoryResponseOutput)
}

type VirtualNetworkProfile struct {
	Id     string  `pulumi:"id"`
	Subnet *string `pulumi:"subnet"`
}





type VirtualNetworkProfileInput interface {
	pulumi.Input

	ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput
	ToVirtualNetworkProfileOutputWithContext(context.Context) VirtualNetworkProfileOutput
}

type VirtualNetworkProfileArgs struct {
	Id     pulumi.StringInput    `pulumi:"id"`
	Subnet pulumi.StringPtrInput `pulumi:"subnet"`
}

func (VirtualNetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfile)(nil)).Elem()
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput {
	return i.ToVirtualNetworkProfileOutputWithContext(context.Background())
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfileOutputWithContext(ctx context.Context) VirtualNetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkProfileOutput)
}

type VirtualNetworkProfileOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfile)(nil)).Elem()
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput {
	return o
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfileOutputWithContext(ctx context.Context) VirtualNetworkProfileOutput {
	return o
}

func (o VirtualNetworkProfileOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkProfile) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualNetworkProfileOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfile) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

type VirtualNetworkProfileResponse struct {
	Id     string  `pulumi:"id"`
	Name   string  `pulumi:"name"`
	Subnet *string `pulumi:"subnet"`
	Type   string  `pulumi:"type"`
}

type VirtualNetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfileResponse)(nil)).Elem()
}

func (o VirtualNetworkProfileResponseOutput) ToVirtualNetworkProfileResponseOutput() VirtualNetworkProfileResponseOutput {
	return o
}

func (o VirtualNetworkProfileResponseOutput) ToVirtualNetworkProfileResponseOutputWithContext(ctx context.Context) VirtualNetworkProfileResponseOutput {
	return o
}

func (o VirtualNetworkProfileResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualNetworkProfileResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkProfileResponseOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VnetRouteResponse struct {
	EndAddress   *string `pulumi:"endAddress"`
	Id           string  `pulumi:"id"`
	Kind         *string `pulumi:"kind"`
	Name         string  `pulumi:"name"`
	RouteType    *string `pulumi:"routeType"`
	StartAddress *string `pulumi:"startAddress"`
	Type         string  `pulumi:"type"`
}

type VnetRouteResponseOutput struct{ *pulumi.OutputState }

func (VnetRouteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VnetRouteResponse)(nil)).Elem()
}

func (o VnetRouteResponseOutput) ToVnetRouteResponseOutput() VnetRouteResponseOutput {
	return o
}

func (o VnetRouteResponseOutput) ToVnetRouteResponseOutputWithContext(ctx context.Context) VnetRouteResponseOutput {
	return o
}

func (o VnetRouteResponseOutput) EndAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRouteResponse) *string { return v.EndAddress }).(pulumi.StringPtrOutput)
}

func (o VnetRouteResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VnetRouteResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o VnetRouteResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRouteResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o VnetRouteResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VnetRouteResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VnetRouteResponseOutput) RouteType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRouteResponse) *string { return v.RouteType }).(pulumi.StringPtrOutput)
}

func (o VnetRouteResponseOutput) StartAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRouteResponse) *string { return v.StartAddress }).(pulumi.StringPtrOutput)
}

func (o VnetRouteResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VnetRouteResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VnetRouteResponseArrayOutput struct{ *pulumi.OutputState }

func (VnetRouteResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VnetRouteResponse)(nil)).Elem()
}

func (o VnetRouteResponseArrayOutput) ToVnetRouteResponseArrayOutput() VnetRouteResponseArrayOutput {
	return o
}

func (o VnetRouteResponseArrayOutput) ToVnetRouteResponseArrayOutputWithContext(ctx context.Context) VnetRouteResponseArrayOutput {
	return o
}

func (o VnetRouteResponseArrayOutput) Index(i pulumi.IntInput) VnetRouteResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VnetRouteResponse {
		return vs[0].([]VnetRouteResponse)[vs[1].(int)]
	}).(VnetRouteResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AllowedAudiencesValidationOutput{})
	pulumi.RegisterOutputType(AllowedAudiencesValidationPtrOutput{})
	pulumi.RegisterOutputType(AllowedAudiencesValidationResponseOutput{})
	pulumi.RegisterOutputType(AllowedAudiencesValidationResponsePtrOutput{})
	pulumi.RegisterOutputType(AllowedPrincipalsOutput{})
	pulumi.RegisterOutputType(AllowedPrincipalsPtrOutput{})
	pulumi.RegisterOutputType(AllowedPrincipalsResponseOutput{})
	pulumi.RegisterOutputType(AllowedPrincipalsResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiDefinitionInfoOutput{})
	pulumi.RegisterOutputType(ApiDefinitionInfoPtrOutput{})
	pulumi.RegisterOutputType(ApiDefinitionInfoResponseOutput{})
	pulumi.RegisterOutputType(ApiDefinitionInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiManagementConfigOutput{})
	pulumi.RegisterOutputType(ApiManagementConfigPtrOutput{})
	pulumi.RegisterOutputType(ApiManagementConfigResponseOutput{})
	pulumi.RegisterOutputType(ApiManagementConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(AppLogsConfigurationOutput{})
	pulumi.RegisterOutputType(AppLogsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(AppLogsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(AppLogsConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(AppRegistrationOutput{})
	pulumi.RegisterOutputType(AppRegistrationPtrOutput{})
	pulumi.RegisterOutputType(AppRegistrationResponseOutput{})
	pulumi.RegisterOutputType(AppRegistrationResponsePtrOutput{})
	pulumi.RegisterOutputType(AppleOutput{})
	pulumi.RegisterOutputType(ApplePtrOutput{})
	pulumi.RegisterOutputType(AppleRegistrationOutput{})
	pulumi.RegisterOutputType(AppleRegistrationPtrOutput{})
	pulumi.RegisterOutputType(AppleRegistrationResponseOutput{})
	pulumi.RegisterOutputType(AppleRegistrationResponsePtrOutput{})
	pulumi.RegisterOutputType(AppleResponseOutput{})
	pulumi.RegisterOutputType(AppleResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationLogsConfigOutput{})
	pulumi.RegisterOutputType(ApplicationLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(ApplicationLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(ApplicationLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(ArcConfigurationOutput{})
	pulumi.RegisterOutputType(ArcConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ArcConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ArcConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ArmIdWrapperResponseOutput{})
	pulumi.RegisterOutputType(ArmIdWrapperResponsePtrOutput{})
	pulumi.RegisterOutputType(ArmPlanResponseOutput{})
	pulumi.RegisterOutputType(ArmPlanResponsePtrOutput{})
	pulumi.RegisterOutputType(AuthPlatformOutput{})
	pulumi.RegisterOutputType(AuthPlatformPtrOutput{})
	pulumi.RegisterOutputType(AuthPlatformResponseOutput{})
	pulumi.RegisterOutputType(AuthPlatformResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoHealActionsOutput{})
	pulumi.RegisterOutputType(AutoHealActionsPtrOutput{})
	pulumi.RegisterOutputType(AutoHealActionsResponseOutput{})
	pulumi.RegisterOutputType(AutoHealActionsResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoHealCustomActionOutput{})
	pulumi.RegisterOutputType(AutoHealCustomActionPtrOutput{})
	pulumi.RegisterOutputType(AutoHealCustomActionResponseOutput{})
	pulumi.RegisterOutputType(AutoHealCustomActionResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoHealRulesOutput{})
	pulumi.RegisterOutputType(AutoHealRulesPtrOutput{})
	pulumi.RegisterOutputType(AutoHealRulesResponseOutput{})
	pulumi.RegisterOutputType(AutoHealRulesResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoHealTriggersOutput{})
	pulumi.RegisterOutputType(AutoHealTriggersPtrOutput{})
	pulumi.RegisterOutputType(AutoHealTriggersResponseOutput{})
	pulumi.RegisterOutputType(AutoHealTriggersResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryPtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryLoginOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryLoginPtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryLoginResponseOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryLoginResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryRegistrationOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryRegistrationPtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryRegistrationResponseOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryRegistrationResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryResponseOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryValidationOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryValidationPtrOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryValidationResponseOutput{})
	pulumi.RegisterOutputType(AzureActiveDirectoryValidationResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageApplicationLogsConfigOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageApplicationLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageApplicationLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageApplicationLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageHttpLogsConfigOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageHttpLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageHttpLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageHttpLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureStaticWebAppsOutput{})
	pulumi.RegisterOutputType(AzureStaticWebAppsPtrOutput{})
	pulumi.RegisterOutputType(AzureStaticWebAppsRegistrationOutput{})
	pulumi.RegisterOutputType(AzureStaticWebAppsRegistrationPtrOutput{})
	pulumi.RegisterOutputType(AzureStaticWebAppsRegistrationResponseOutput{})
	pulumi.RegisterOutputType(AzureStaticWebAppsRegistrationResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureStaticWebAppsResponseOutput{})
	pulumi.RegisterOutputType(AzureStaticWebAppsResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureStorageInfoValueOutput{})
	pulumi.RegisterOutputType(AzureStorageInfoValueMapOutput{})
	pulumi.RegisterOutputType(AzureStorageInfoValueResponseOutput{})
	pulumi.RegisterOutputType(AzureStorageInfoValueResponseMapOutput{})
	pulumi.RegisterOutputType(AzureTableStorageApplicationLogsConfigOutput{})
	pulumi.RegisterOutputType(AzureTableStorageApplicationLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(AzureTableStorageApplicationLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(AzureTableStorageApplicationLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(BackupItemResponseOutput{})
	pulumi.RegisterOutputType(BackupItemResponseArrayOutput{})
	pulumi.RegisterOutputType(BackupScheduleOutput{})
	pulumi.RegisterOutputType(BackupSchedulePtrOutput{})
	pulumi.RegisterOutputType(BackupScheduleResponseOutput{})
	pulumi.RegisterOutputType(BackupScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(BlobStorageTokenStoreOutput{})
	pulumi.RegisterOutputType(BlobStorageTokenStorePtrOutput{})
	pulumi.RegisterOutputType(BlobStorageTokenStoreResponseOutput{})
	pulumi.RegisterOutputType(BlobStorageTokenStoreResponsePtrOutput{})
	pulumi.RegisterOutputType(CapabilityOutput{})
	pulumi.RegisterOutputType(CapabilityArrayOutput{})
	pulumi.RegisterOutputType(CapabilityResponseOutput{})
	pulumi.RegisterOutputType(CapabilityResponseArrayOutput{})
	pulumi.RegisterOutputType(ClientRegistrationOutput{})
	pulumi.RegisterOutputType(ClientRegistrationPtrOutput{})
	pulumi.RegisterOutputType(ClientRegistrationResponseOutput{})
	pulumi.RegisterOutputType(ClientRegistrationResponsePtrOutput{})
	pulumi.RegisterOutputType(CloningInfoOutput{})
	pulumi.RegisterOutputType(CloningInfoPtrOutput{})
	pulumi.RegisterOutputType(ConnStringInfoOutput{})
	pulumi.RegisterOutputType(ConnStringInfoArrayOutput{})
	pulumi.RegisterOutputType(ConnStringInfoResponseOutput{})
	pulumi.RegisterOutputType(ConnStringInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ConnStringValueTypePairOutput{})
	pulumi.RegisterOutputType(ConnStringValueTypePairMapOutput{})
	pulumi.RegisterOutputType(ConnStringValueTypePairResponseOutput{})
	pulumi.RegisterOutputType(ConnStringValueTypePairResponseMapOutput{})
	pulumi.RegisterOutputType(CookieExpirationOutput{})
	pulumi.RegisterOutputType(CookieExpirationPtrOutput{})
	pulumi.RegisterOutputType(CookieExpirationResponseOutput{})
	pulumi.RegisterOutputType(CookieExpirationResponsePtrOutput{})
	pulumi.RegisterOutputType(CorsSettingsOutput{})
	pulumi.RegisterOutputType(CorsSettingsPtrOutput{})
	pulumi.RegisterOutputType(CorsSettingsResponseOutput{})
	pulumi.RegisterOutputType(CorsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(CustomOpenIdConnectProviderOutput{})
	pulumi.RegisterOutputType(CustomOpenIdConnectProviderMapOutput{})
	pulumi.RegisterOutputType(CustomOpenIdConnectProviderResponseOutput{})
	pulumi.RegisterOutputType(CustomOpenIdConnectProviderResponseMapOutput{})
	pulumi.RegisterOutputType(DatabaseBackupSettingOutput{})
	pulumi.RegisterOutputType(DatabaseBackupSettingArrayOutput{})
	pulumi.RegisterOutputType(DatabaseBackupSettingResponseOutput{})
	pulumi.RegisterOutputType(DatabaseBackupSettingResponseArrayOutput{})
	pulumi.RegisterOutputType(DefaultAuthorizationPolicyOutput{})
	pulumi.RegisterOutputType(DefaultAuthorizationPolicyPtrOutput{})
	pulumi.RegisterOutputType(DefaultAuthorizationPolicyResponseOutput{})
	pulumi.RegisterOutputType(DefaultAuthorizationPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(EnabledConfigOutput{})
	pulumi.RegisterOutputType(EnabledConfigPtrOutput{})
	pulumi.RegisterOutputType(EnabledConfigResponseOutput{})
	pulumi.RegisterOutputType(EnabledConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorEntityResponseOutput{})
	pulumi.RegisterOutputType(ErrorEntityResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorEntityResponseArrayOutput{})
	pulumi.RegisterOutputType(ExperimentsOutput{})
	pulumi.RegisterOutputType(ExperimentsPtrOutput{})
	pulumi.RegisterOutputType(ExperimentsResponseOutput{})
	pulumi.RegisterOutputType(ExperimentsResponsePtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationOutput{})
	pulumi.RegisterOutputType(ExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponsePtrOutput{})
	pulumi.RegisterOutputType(FacebookOutput{})
	pulumi.RegisterOutputType(FacebookPtrOutput{})
	pulumi.RegisterOutputType(FacebookResponseOutput{})
	pulumi.RegisterOutputType(FacebookResponsePtrOutput{})
	pulumi.RegisterOutputType(FileSystemApplicationLogsConfigOutput{})
	pulumi.RegisterOutputType(FileSystemApplicationLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(FileSystemApplicationLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(FileSystemApplicationLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(FileSystemHttpLogsConfigOutput{})
	pulumi.RegisterOutputType(FileSystemHttpLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(FileSystemHttpLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(FileSystemHttpLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(FileSystemTokenStoreOutput{})
	pulumi.RegisterOutputType(FileSystemTokenStorePtrOutput{})
	pulumi.RegisterOutputType(FileSystemTokenStoreResponseOutput{})
	pulumi.RegisterOutputType(FileSystemTokenStoreResponsePtrOutput{})
	pulumi.RegisterOutputType(ForwardProxyOutput{})
	pulumi.RegisterOutputType(ForwardProxyPtrOutput{})
	pulumi.RegisterOutputType(ForwardProxyResponseOutput{})
	pulumi.RegisterOutputType(ForwardProxyResponsePtrOutput{})
	pulumi.RegisterOutputType(FrontEndConfigurationOutput{})
	pulumi.RegisterOutputType(FrontEndConfigurationPtrOutput{})
	pulumi.RegisterOutputType(FrontEndConfigurationResponseOutput{})
	pulumi.RegisterOutputType(FrontEndConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(GitHubOutput{})
	pulumi.RegisterOutputType(GitHubPtrOutput{})
	pulumi.RegisterOutputType(GitHubActionCodeConfigurationOutput{})
	pulumi.RegisterOutputType(GitHubActionCodeConfigurationPtrOutput{})
	pulumi.RegisterOutputType(GitHubActionCodeConfigurationResponseOutput{})
	pulumi.RegisterOutputType(GitHubActionCodeConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(GitHubActionConfigurationOutput{})
	pulumi.RegisterOutputType(GitHubActionConfigurationPtrOutput{})
	pulumi.RegisterOutputType(GitHubActionConfigurationResponseOutput{})
	pulumi.RegisterOutputType(GitHubActionConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(GitHubActionContainerConfigurationOutput{})
	pulumi.RegisterOutputType(GitHubActionContainerConfigurationPtrOutput{})
	pulumi.RegisterOutputType(GitHubActionContainerConfigurationResponseOutput{})
	pulumi.RegisterOutputType(GitHubActionContainerConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(GitHubResponseOutput{})
	pulumi.RegisterOutputType(GitHubResponsePtrOutput{})
	pulumi.RegisterOutputType(GlobalValidationOutput{})
	pulumi.RegisterOutputType(GlobalValidationPtrOutput{})
	pulumi.RegisterOutputType(GlobalValidationResponseOutput{})
	pulumi.RegisterOutputType(GlobalValidationResponsePtrOutput{})
	pulumi.RegisterOutputType(GoogleOutput{})
	pulumi.RegisterOutputType(GooglePtrOutput{})
	pulumi.RegisterOutputType(GoogleResponseOutput{})
	pulumi.RegisterOutputType(GoogleResponsePtrOutput{})
	pulumi.RegisterOutputType(HandlerMappingOutput{})
	pulumi.RegisterOutputType(HandlerMappingArrayOutput{})
	pulumi.RegisterOutputType(HandlerMappingResponseOutput{})
	pulumi.RegisterOutputType(HandlerMappingResponseArrayOutput{})
	pulumi.RegisterOutputType(HostNameSslStateOutput{})
	pulumi.RegisterOutputType(HostNameSslStateArrayOutput{})
	pulumi.RegisterOutputType(HostNameSslStateResponseOutput{})
	pulumi.RegisterOutputType(HostNameSslStateResponseArrayOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentProfileOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentProfilePtrOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentProfileResponseOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(HttpLogsConfigOutput{})
	pulumi.RegisterOutputType(HttpLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(HttpLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(HttpLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(HttpSettingsOutput{})
	pulumi.RegisterOutputType(HttpSettingsPtrOutput{})
	pulumi.RegisterOutputType(HttpSettingsResponseOutput{})
	pulumi.RegisterOutputType(HttpSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(HttpSettingsRoutesOutput{})
	pulumi.RegisterOutputType(HttpSettingsRoutesPtrOutput{})
	pulumi.RegisterOutputType(HttpSettingsRoutesResponseOutput{})
	pulumi.RegisterOutputType(HttpSettingsRoutesResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentifierResponseOutput{})
	pulumi.RegisterOutputType(IdentifierResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityProvidersOutput{})
	pulumi.RegisterOutputType(IdentityProvidersPtrOutput{})
	pulumi.RegisterOutputType(IdentityProvidersResponseOutput{})
	pulumi.RegisterOutputType(IdentityProvidersResponsePtrOutput{})
	pulumi.RegisterOutputType(IpSecurityRestrictionOutput{})
	pulumi.RegisterOutputType(IpSecurityRestrictionArrayOutput{})
	pulumi.RegisterOutputType(IpSecurityRestrictionResponseOutput{})
	pulumi.RegisterOutputType(IpSecurityRestrictionResponseArrayOutput{})
	pulumi.RegisterOutputType(JwtClaimChecksOutput{})
	pulumi.RegisterOutputType(JwtClaimChecksPtrOutput{})
	pulumi.RegisterOutputType(JwtClaimChecksResponseOutput{})
	pulumi.RegisterOutputType(JwtClaimChecksResponsePtrOutput{})
	pulumi.RegisterOutputType(KubeEnvironmentProfileOutput{})
	pulumi.RegisterOutputType(KubeEnvironmentProfilePtrOutput{})
	pulumi.RegisterOutputType(KubeEnvironmentProfileResponseOutput{})
	pulumi.RegisterOutputType(KubeEnvironmentProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(LegacyMicrosoftAccountOutput{})
	pulumi.RegisterOutputType(LegacyMicrosoftAccountPtrOutput{})
	pulumi.RegisterOutputType(LegacyMicrosoftAccountResponseOutput{})
	pulumi.RegisterOutputType(LegacyMicrosoftAccountResponsePtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsConfigurationOutput{})
	pulumi.RegisterOutputType(LogAnalyticsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(LogAnalyticsConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(LoginOutput{})
	pulumi.RegisterOutputType(LoginPtrOutput{})
	pulumi.RegisterOutputType(LoginResponseOutput{})
	pulumi.RegisterOutputType(LoginResponsePtrOutput{})
	pulumi.RegisterOutputType(LoginRoutesOutput{})
	pulumi.RegisterOutputType(LoginRoutesPtrOutput{})
	pulumi.RegisterOutputType(LoginRoutesResponseOutput{})
	pulumi.RegisterOutputType(LoginRoutesResponsePtrOutput{})
	pulumi.RegisterOutputType(LoginScopesOutput{})
	pulumi.RegisterOutputType(LoginScopesPtrOutput{})
	pulumi.RegisterOutputType(LoginScopesResponseOutput{})
	pulumi.RegisterOutputType(LoginScopesResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(NameValuePairOutput{})
	pulumi.RegisterOutputType(NameValuePairArrayOutput{})
	pulumi.RegisterOutputType(NameValuePairResponseOutput{})
	pulumi.RegisterOutputType(NameValuePairResponseArrayOutput{})
	pulumi.RegisterOutputType(NonceOutput{})
	pulumi.RegisterOutputType(NoncePtrOutput{})
	pulumi.RegisterOutputType(NonceResponseOutput{})
	pulumi.RegisterOutputType(NonceResponsePtrOutput{})
	pulumi.RegisterOutputType(OpenIdConnectClientCredentialOutput{})
	pulumi.RegisterOutputType(OpenIdConnectClientCredentialPtrOutput{})
	pulumi.RegisterOutputType(OpenIdConnectClientCredentialResponseOutput{})
	pulumi.RegisterOutputType(OpenIdConnectClientCredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(OpenIdConnectConfigOutput{})
	pulumi.RegisterOutputType(OpenIdConnectConfigPtrOutput{})
	pulumi.RegisterOutputType(OpenIdConnectConfigResponseOutput{})
	pulumi.RegisterOutputType(OpenIdConnectConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(OpenIdConnectLoginOutput{})
	pulumi.RegisterOutputType(OpenIdConnectLoginPtrOutput{})
	pulumi.RegisterOutputType(OpenIdConnectLoginResponseOutput{})
	pulumi.RegisterOutputType(OpenIdConnectLoginResponsePtrOutput{})
	pulumi.RegisterOutputType(OpenIdConnectRegistrationOutput{})
	pulumi.RegisterOutputType(OpenIdConnectRegistrationPtrOutput{})
	pulumi.RegisterOutputType(OpenIdConnectRegistrationResponseOutput{})
	pulumi.RegisterOutputType(OpenIdConnectRegistrationResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(PushSettingsOutput{})
	pulumi.RegisterOutputType(PushSettingsPtrOutput{})
	pulumi.RegisterOutputType(PushSettingsResponseOutput{})
	pulumi.RegisterOutputType(PushSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(RampUpRuleOutput{})
	pulumi.RegisterOutputType(RampUpRuleArrayOutput{})
	pulumi.RegisterOutputType(RampUpRuleResponseOutput{})
	pulumi.RegisterOutputType(RampUpRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(RemotePrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(RemotePrivateEndpointConnectionResponsePtrOutput{})
	pulumi.RegisterOutputType(RequestsBasedTriggerOutput{})
	pulumi.RegisterOutputType(RequestsBasedTriggerPtrOutput{})
	pulumi.RegisterOutputType(RequestsBasedTriggerResponseOutput{})
	pulumi.RegisterOutputType(RequestsBasedTriggerResponsePtrOutput{})
	pulumi.RegisterOutputType(ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(SiteConfigOutput{})
	pulumi.RegisterOutputType(SiteConfigPtrOutput{})
	pulumi.RegisterOutputType(SiteConfigResponseOutput{})
	pulumi.RegisterOutputType(SiteConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(SiteLimitsOutput{})
	pulumi.RegisterOutputType(SiteLimitsPtrOutput{})
	pulumi.RegisterOutputType(SiteLimitsResponseOutput{})
	pulumi.RegisterOutputType(SiteLimitsResponsePtrOutput{})
	pulumi.RegisterOutputType(SiteMachineKeyResponseOutput{})
	pulumi.RegisterOutputType(SiteMachineKeyResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuCapacityOutput{})
	pulumi.RegisterOutputType(SkuCapacityPtrOutput{})
	pulumi.RegisterOutputType(SkuCapacityResponseOutput{})
	pulumi.RegisterOutputType(SkuCapacityResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuDescriptionOutput{})
	pulumi.RegisterOutputType(SkuDescriptionPtrOutput{})
	pulumi.RegisterOutputType(SkuDescriptionResponseOutput{})
	pulumi.RegisterOutputType(SkuDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(SlotSwapStatusResponseOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerPtrOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerArrayOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerResponseOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerResponsePtrOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerResponseArrayOutput{})
	pulumi.RegisterOutputType(StaticSiteBuildPropertiesOutput{})
	pulumi.RegisterOutputType(StaticSiteBuildPropertiesPtrOutput{})
	pulumi.RegisterOutputType(StaticSiteBuildPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StaticSiteBuildPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(StaticSiteTemplateOptionsOutput{})
	pulumi.RegisterOutputType(StaticSiteTemplateOptionsPtrOutput{})
	pulumi.RegisterOutputType(StaticSiteTemplateOptionsResponseOutput{})
	pulumi.RegisterOutputType(StaticSiteTemplateOptionsResponsePtrOutput{})
	pulumi.RegisterOutputType(StaticSiteUserARMResourceResponseOutput{})
	pulumi.RegisterOutputType(StaticSiteUserARMResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(StaticSiteUserProvidedFunctionAppResponseOutput{})
	pulumi.RegisterOutputType(StaticSiteUserProvidedFunctionAppResponseArrayOutput{})
	pulumi.RegisterOutputType(StatusCodesBasedTriggerOutput{})
	pulumi.RegisterOutputType(StatusCodesBasedTriggerArrayOutput{})
	pulumi.RegisterOutputType(StatusCodesBasedTriggerResponseOutput{})
	pulumi.RegisterOutputType(StatusCodesBasedTriggerResponseArrayOutput{})
	pulumi.RegisterOutputType(StatusCodesRangeBasedTriggerOutput{})
	pulumi.RegisterOutputType(StatusCodesRangeBasedTriggerArrayOutput{})
	pulumi.RegisterOutputType(StatusCodesRangeBasedTriggerResponseOutput{})
	pulumi.RegisterOutputType(StatusCodesRangeBasedTriggerResponseArrayOutput{})
	pulumi.RegisterOutputType(TokenStoreOutput{})
	pulumi.RegisterOutputType(TokenStorePtrOutput{})
	pulumi.RegisterOutputType(TokenStoreResponseOutput{})
	pulumi.RegisterOutputType(TokenStoreResponsePtrOutput{})
	pulumi.RegisterOutputType(TwitterOutput{})
	pulumi.RegisterOutputType(TwitterPtrOutput{})
	pulumi.RegisterOutputType(TwitterRegistrationOutput{})
	pulumi.RegisterOutputType(TwitterRegistrationPtrOutput{})
	pulumi.RegisterOutputType(TwitterRegistrationResponseOutput{})
	pulumi.RegisterOutputType(TwitterRegistrationResponsePtrOutput{})
	pulumi.RegisterOutputType(TwitterResponseOutput{})
	pulumi.RegisterOutputType(TwitterResponsePtrOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(VirtualApplicationOutput{})
	pulumi.RegisterOutputType(VirtualApplicationArrayOutput{})
	pulumi.RegisterOutputType(VirtualApplicationResponseOutput{})
	pulumi.RegisterOutputType(VirtualApplicationResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualDirectoryOutput{})
	pulumi.RegisterOutputType(VirtualDirectoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualDirectoryResponseOutput{})
	pulumi.RegisterOutputType(VirtualDirectoryResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfileOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(VnetRouteResponseOutput{})
	pulumi.RegisterOutputType(VnetRouteResponseArrayOutput{})
}
