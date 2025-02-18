


package v20151101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BasicDependencyResponse struct {
	Id           *string `pulumi:"id"`
	ResourceName *string `pulumi:"resourceName"`
	ResourceType *string `pulumi:"resourceType"`
}

type BasicDependencyResponseOutput struct{ *pulumi.OutputState }

func (BasicDependencyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BasicDependencyResponse)(nil)).Elem()
}

func (o BasicDependencyResponseOutput) ToBasicDependencyResponseOutput() BasicDependencyResponseOutput {
	return o
}

func (o BasicDependencyResponseOutput) ToBasicDependencyResponseOutputWithContext(ctx context.Context) BasicDependencyResponseOutput {
	return o
}

func (o BasicDependencyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BasicDependencyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o BasicDependencyResponseOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BasicDependencyResponse) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o BasicDependencyResponseOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BasicDependencyResponse) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

type BasicDependencyResponseArrayOutput struct{ *pulumi.OutputState }

func (BasicDependencyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BasicDependencyResponse)(nil)).Elem()
}

func (o BasicDependencyResponseArrayOutput) ToBasicDependencyResponseArrayOutput() BasicDependencyResponseArrayOutput {
	return o
}

func (o BasicDependencyResponseArrayOutput) ToBasicDependencyResponseArrayOutputWithContext(ctx context.Context) BasicDependencyResponseArrayOutput {
	return o
}

func (o BasicDependencyResponseArrayOutput) Index(i pulumi.IntInput) BasicDependencyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BasicDependencyResponse {
		return vs[0].([]BasicDependencyResponse)[vs[1].(int)]
	}).(BasicDependencyResponseOutput)
}

type DependencyResponse struct {
	DependsOn    []BasicDependencyResponse `pulumi:"dependsOn"`
	Id           *string                   `pulumi:"id"`
	ResourceName *string                   `pulumi:"resourceName"`
	ResourceType *string                   `pulumi:"resourceType"`
}

type DependencyResponseOutput struct{ *pulumi.OutputState }

func (DependencyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DependencyResponse)(nil)).Elem()
}

func (o DependencyResponseOutput) ToDependencyResponseOutput() DependencyResponseOutput {
	return o
}

func (o DependencyResponseOutput) ToDependencyResponseOutputWithContext(ctx context.Context) DependencyResponseOutput {
	return o
}

func (o DependencyResponseOutput) DependsOn() BasicDependencyResponseArrayOutput {
	return o.ApplyT(func(v DependencyResponse) []BasicDependencyResponse { return v.DependsOn }).(BasicDependencyResponseArrayOutput)
}

func (o DependencyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DependencyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o DependencyResponseOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DependencyResponse) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o DependencyResponseOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DependencyResponse) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

type DependencyResponseArrayOutput struct{ *pulumi.OutputState }

func (DependencyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DependencyResponse)(nil)).Elem()
}

func (o DependencyResponseArrayOutput) ToDependencyResponseArrayOutput() DependencyResponseArrayOutput {
	return o
}

func (o DependencyResponseArrayOutput) ToDependencyResponseArrayOutputWithContext(ctx context.Context) DependencyResponseArrayOutput {
	return o
}

func (o DependencyResponseArrayOutput) Index(i pulumi.IntInput) DependencyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DependencyResponse {
		return vs[0].([]DependencyResponse)[vs[1].(int)]
	}).(DependencyResponseOutput)
}

type DeploymentProperties struct {
	Mode           *DeploymentMode `pulumi:"mode"`
	Parameters     interface{}     `pulumi:"parameters"`
	ParametersLink *ParametersLink `pulumi:"parametersLink"`
	Template       interface{}     `pulumi:"template"`
	TemplateLink   *TemplateLink   `pulumi:"templateLink"`
}





type DeploymentPropertiesInput interface {
	pulumi.Input

	ToDeploymentPropertiesOutput() DeploymentPropertiesOutput
	ToDeploymentPropertiesOutputWithContext(context.Context) DeploymentPropertiesOutput
}

type DeploymentPropertiesArgs struct {
	Mode           DeploymentModePtrInput `pulumi:"mode"`
	Parameters     pulumi.Input           `pulumi:"parameters"`
	ParametersLink ParametersLinkPtrInput `pulumi:"parametersLink"`
	Template       pulumi.Input           `pulumi:"template"`
	TemplateLink   TemplateLinkPtrInput   `pulumi:"templateLink"`
}

func (DeploymentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentProperties)(nil)).Elem()
}

func (i DeploymentPropertiesArgs) ToDeploymentPropertiesOutput() DeploymentPropertiesOutput {
	return i.ToDeploymentPropertiesOutputWithContext(context.Background())
}

func (i DeploymentPropertiesArgs) ToDeploymentPropertiesOutputWithContext(ctx context.Context) DeploymentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPropertiesOutput)
}

func (i DeploymentPropertiesArgs) ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput {
	return i.ToDeploymentPropertiesPtrOutputWithContext(context.Background())
}

func (i DeploymentPropertiesArgs) ToDeploymentPropertiesPtrOutputWithContext(ctx context.Context) DeploymentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPropertiesOutput).ToDeploymentPropertiesPtrOutputWithContext(ctx)
}









type DeploymentPropertiesPtrInput interface {
	pulumi.Input

	ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput
	ToDeploymentPropertiesPtrOutputWithContext(context.Context) DeploymentPropertiesPtrOutput
}

type deploymentPropertiesPtrType DeploymentPropertiesArgs

func DeploymentPropertiesPtr(v *DeploymentPropertiesArgs) DeploymentPropertiesPtrInput {
	return (*deploymentPropertiesPtrType)(v)
}

func (*deploymentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentProperties)(nil)).Elem()
}

func (i *deploymentPropertiesPtrType) ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput {
	return i.ToDeploymentPropertiesPtrOutputWithContext(context.Background())
}

func (i *deploymentPropertiesPtrType) ToDeploymentPropertiesPtrOutputWithContext(ctx context.Context) DeploymentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPropertiesPtrOutput)
}

type DeploymentPropertiesOutput struct{ *pulumi.OutputState }

func (DeploymentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentProperties)(nil)).Elem()
}

func (o DeploymentPropertiesOutput) ToDeploymentPropertiesOutput() DeploymentPropertiesOutput {
	return o
}

func (o DeploymentPropertiesOutput) ToDeploymentPropertiesOutputWithContext(ctx context.Context) DeploymentPropertiesOutput {
	return o
}

func (o DeploymentPropertiesOutput) ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput {
	return o.ToDeploymentPropertiesPtrOutputWithContext(context.Background())
}

func (o DeploymentPropertiesOutput) ToDeploymentPropertiesPtrOutputWithContext(ctx context.Context) DeploymentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentProperties) *DeploymentProperties {
		return &v
	}).(DeploymentPropertiesPtrOutput)
}

func (o DeploymentPropertiesOutput) Mode() DeploymentModePtrOutput {
	return o.ApplyT(func(v DeploymentProperties) *DeploymentMode { return v.Mode }).(DeploymentModePtrOutput)
}

func (o DeploymentPropertiesOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v DeploymentProperties) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesOutput) ParametersLink() ParametersLinkPtrOutput {
	return o.ApplyT(func(v DeploymentProperties) *ParametersLink { return v.ParametersLink }).(ParametersLinkPtrOutput)
}

func (o DeploymentPropertiesOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v DeploymentProperties) interface{} { return v.Template }).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesOutput) TemplateLink() TemplateLinkPtrOutput {
	return o.ApplyT(func(v DeploymentProperties) *TemplateLink { return v.TemplateLink }).(TemplateLinkPtrOutput)
}

type DeploymentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (DeploymentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentProperties)(nil)).Elem()
}

func (o DeploymentPropertiesPtrOutput) ToDeploymentPropertiesPtrOutput() DeploymentPropertiesPtrOutput {
	return o
}

func (o DeploymentPropertiesPtrOutput) ToDeploymentPropertiesPtrOutputWithContext(ctx context.Context) DeploymentPropertiesPtrOutput {
	return o
}

func (o DeploymentPropertiesPtrOutput) Elem() DeploymentPropertiesOutput {
	return o.ApplyT(func(v *DeploymentProperties) DeploymentProperties {
		if v != nil {
			return *v
		}
		var ret DeploymentProperties
		return ret
	}).(DeploymentPropertiesOutput)
}

func (o DeploymentPropertiesPtrOutput) Mode() DeploymentModePtrOutput {
	return o.ApplyT(func(v *DeploymentProperties) *DeploymentMode {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(DeploymentModePtrOutput)
}

func (o DeploymentPropertiesPtrOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *DeploymentProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesPtrOutput) ParametersLink() ParametersLinkPtrOutput {
	return o.ApplyT(func(v *DeploymentProperties) *ParametersLink {
		if v == nil {
			return nil
		}
		return v.ParametersLink
	}).(ParametersLinkPtrOutput)
}

func (o DeploymentPropertiesPtrOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v *DeploymentProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.Template
	}).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesPtrOutput) TemplateLink() TemplateLinkPtrOutput {
	return o.ApplyT(func(v *DeploymentProperties) *TemplateLink {
		if v == nil {
			return nil
		}
		return v.TemplateLink
	}).(TemplateLinkPtrOutput)
}

type DeploymentPropertiesExtendedResponse struct {
	CorrelationId     *string                 `pulumi:"correlationId"`
	Dependencies      []DependencyResponse    `pulumi:"dependencies"`
	Error             ErrorResponseResponse   `pulumi:"error"`
	Mode              *string                 `pulumi:"mode"`
	Outputs           interface{}             `pulumi:"outputs"`
	Parameters        interface{}             `pulumi:"parameters"`
	ParametersLink    *ParametersLinkResponse `pulumi:"parametersLink"`
	Providers         []ProviderResponse      `pulumi:"providers"`
	ProvisioningState *string                 `pulumi:"provisioningState"`
	Template          interface{}             `pulumi:"template"`
	TemplateLink      *TemplateLinkResponse   `pulumi:"templateLink"`
	Timestamp         *string                 `pulumi:"timestamp"`
}

type DeploymentPropertiesExtendedResponseOutput struct{ *pulumi.OutputState }

func (DeploymentPropertiesExtendedResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentPropertiesExtendedResponse)(nil)).Elem()
}

func (o DeploymentPropertiesExtendedResponseOutput) ToDeploymentPropertiesExtendedResponseOutput() DeploymentPropertiesExtendedResponseOutput {
	return o
}

func (o DeploymentPropertiesExtendedResponseOutput) ToDeploymentPropertiesExtendedResponseOutputWithContext(ctx context.Context) DeploymentPropertiesExtendedResponseOutput {
	return o
}

func (o DeploymentPropertiesExtendedResponseOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) *string { return v.CorrelationId }).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Dependencies() DependencyResponseArrayOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) []DependencyResponse { return v.Dependencies }).(DependencyResponseArrayOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Error() ErrorResponseResponseOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) ErrorResponseResponse { return v.Error }).(ErrorResponseResponseOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Outputs() pulumi.AnyOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) interface{} { return v.Outputs }).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) ParametersLink() ParametersLinkResponsePtrOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) *ParametersLinkResponse { return v.ParametersLink }).(ParametersLinkResponsePtrOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Providers() ProviderResponseArrayOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) []ProviderResponse { return v.Providers }).(ProviderResponseArrayOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) interface{} { return v.Template }).(pulumi.AnyOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) TemplateLink() TemplateLinkResponsePtrOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) *TemplateLinkResponse { return v.TemplateLink }).(TemplateLinkResponsePtrOutput)
}

func (o DeploymentPropertiesExtendedResponseOutput) Timestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentPropertiesExtendedResponse) *string { return v.Timestamp }).(pulumi.StringPtrOutput)
}

type ErrorAdditionalInfoResponse struct {
	Info interface{} `pulumi:"info"`
	Type string      `pulumi:"type"`
}

type ErrorAdditionalInfoResponseOutput struct{ *pulumi.OutputState }

func (ErrorAdditionalInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (o ErrorAdditionalInfoResponseOutput) ToErrorAdditionalInfoResponseOutput() ErrorAdditionalInfoResponseOutput {
	return o
}

func (o ErrorAdditionalInfoResponseOutput) ToErrorAdditionalInfoResponseOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseOutput {
	return o
}

func (o ErrorAdditionalInfoResponseOutput) Info() pulumi.AnyOutput {
	return o.ApplyT(func(v ErrorAdditionalInfoResponse) interface{} { return v.Info }).(pulumi.AnyOutput)
}

func (o ErrorAdditionalInfoResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorAdditionalInfoResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ErrorAdditionalInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorAdditionalInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorAdditionalInfoResponse)(nil)).Elem()
}

func (o ErrorAdditionalInfoResponseArrayOutput) ToErrorAdditionalInfoResponseArrayOutput() ErrorAdditionalInfoResponseArrayOutput {
	return o
}

func (o ErrorAdditionalInfoResponseArrayOutput) ToErrorAdditionalInfoResponseArrayOutputWithContext(ctx context.Context) ErrorAdditionalInfoResponseArrayOutput {
	return o
}

func (o ErrorAdditionalInfoResponseArrayOutput) Index(i pulumi.IntInput) ErrorAdditionalInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorAdditionalInfoResponse {
		return vs[0].([]ErrorAdditionalInfoResponse)[vs[1].(int)]
	}).(ErrorAdditionalInfoResponseOutput)
}

type ErrorResponseResponse struct {
	AdditionalInfo []ErrorAdditionalInfoResponse `pulumi:"additionalInfo"`
	Code           string                        `pulumi:"code"`
	Details        []ErrorResponseResponse       `pulumi:"details"`
	Message        string                        `pulumi:"message"`
	Target         string                        `pulumi:"target"`
}

type ErrorResponseResponseOutput struct{ *pulumi.OutputState }

func (ErrorResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorResponseResponse)(nil)).Elem()
}

func (o ErrorResponseResponseOutput) ToErrorResponseResponseOutput() ErrorResponseResponseOutput {
	return o
}

func (o ErrorResponseResponseOutput) ToErrorResponseResponseOutputWithContext(ctx context.Context) ErrorResponseResponseOutput {
	return o
}

func (o ErrorResponseResponseOutput) AdditionalInfo() ErrorAdditionalInfoResponseArrayOutput {
	return o.ApplyT(func(v ErrorResponseResponse) []ErrorAdditionalInfoResponse { return v.AdditionalInfo }).(ErrorAdditionalInfoResponseArrayOutput)
}

func (o ErrorResponseResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponseResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorResponseResponseOutput) Details() ErrorResponseResponseArrayOutput {
	return o.ApplyT(func(v ErrorResponseResponse) []ErrorResponseResponse { return v.Details }).(ErrorResponseResponseArrayOutput)
}

func (o ErrorResponseResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponseResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ErrorResponseResponseOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorResponseResponse) string { return v.Target }).(pulumi.StringOutput)
}

type ErrorResponseResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorResponseResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorResponseResponse)(nil)).Elem()
}

func (o ErrorResponseResponseArrayOutput) ToErrorResponseResponseArrayOutput() ErrorResponseResponseArrayOutput {
	return o
}

func (o ErrorResponseResponseArrayOutput) ToErrorResponseResponseArrayOutputWithContext(ctx context.Context) ErrorResponseResponseArrayOutput {
	return o
}

func (o ErrorResponseResponseArrayOutput) Index(i pulumi.IntInput) ErrorResponseResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorResponseResponse {
		return vs[0].([]ErrorResponseResponse)[vs[1].(int)]
	}).(ErrorResponseResponseOutput)
}

type ParametersLink struct {
	ContentVersion *string `pulumi:"contentVersion"`
	Uri            string  `pulumi:"uri"`
}





type ParametersLinkInput interface {
	pulumi.Input

	ToParametersLinkOutput() ParametersLinkOutput
	ToParametersLinkOutputWithContext(context.Context) ParametersLinkOutput
}

type ParametersLinkArgs struct {
	ContentVersion pulumi.StringPtrInput `pulumi:"contentVersion"`
	Uri            pulumi.StringInput    `pulumi:"uri"`
}

func (ParametersLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParametersLink)(nil)).Elem()
}

func (i ParametersLinkArgs) ToParametersLinkOutput() ParametersLinkOutput {
	return i.ToParametersLinkOutputWithContext(context.Background())
}

func (i ParametersLinkArgs) ToParametersLinkOutputWithContext(ctx context.Context) ParametersLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParametersLinkOutput)
}

func (i ParametersLinkArgs) ToParametersLinkPtrOutput() ParametersLinkPtrOutput {
	return i.ToParametersLinkPtrOutputWithContext(context.Background())
}

func (i ParametersLinkArgs) ToParametersLinkPtrOutputWithContext(ctx context.Context) ParametersLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParametersLinkOutput).ToParametersLinkPtrOutputWithContext(ctx)
}









type ParametersLinkPtrInput interface {
	pulumi.Input

	ToParametersLinkPtrOutput() ParametersLinkPtrOutput
	ToParametersLinkPtrOutputWithContext(context.Context) ParametersLinkPtrOutput
}

type parametersLinkPtrType ParametersLinkArgs

func ParametersLinkPtr(v *ParametersLinkArgs) ParametersLinkPtrInput {
	return (*parametersLinkPtrType)(v)
}

func (*parametersLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ParametersLink)(nil)).Elem()
}

func (i *parametersLinkPtrType) ToParametersLinkPtrOutput() ParametersLinkPtrOutput {
	return i.ToParametersLinkPtrOutputWithContext(context.Background())
}

func (i *parametersLinkPtrType) ToParametersLinkPtrOutputWithContext(ctx context.Context) ParametersLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParametersLinkPtrOutput)
}

type ParametersLinkOutput struct{ *pulumi.OutputState }

func (ParametersLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParametersLink)(nil)).Elem()
}

func (o ParametersLinkOutput) ToParametersLinkOutput() ParametersLinkOutput {
	return o
}

func (o ParametersLinkOutput) ToParametersLinkOutputWithContext(ctx context.Context) ParametersLinkOutput {
	return o
}

func (o ParametersLinkOutput) ToParametersLinkPtrOutput() ParametersLinkPtrOutput {
	return o.ToParametersLinkPtrOutputWithContext(context.Background())
}

func (o ParametersLinkOutput) ToParametersLinkPtrOutputWithContext(ctx context.Context) ParametersLinkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ParametersLink) *ParametersLink {
		return &v
	}).(ParametersLinkPtrOutput)
}

func (o ParametersLinkOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParametersLink) *string { return v.ContentVersion }).(pulumi.StringPtrOutput)
}

func (o ParametersLinkOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ParametersLink) string { return v.Uri }).(pulumi.StringOutput)
}

type ParametersLinkPtrOutput struct{ *pulumi.OutputState }

func (ParametersLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParametersLink)(nil)).Elem()
}

func (o ParametersLinkPtrOutput) ToParametersLinkPtrOutput() ParametersLinkPtrOutput {
	return o
}

func (o ParametersLinkPtrOutput) ToParametersLinkPtrOutputWithContext(ctx context.Context) ParametersLinkPtrOutput {
	return o
}

func (o ParametersLinkPtrOutput) Elem() ParametersLinkOutput {
	return o.ApplyT(func(v *ParametersLink) ParametersLink {
		if v != nil {
			return *v
		}
		var ret ParametersLink
		return ret
	}).(ParametersLinkOutput)
}

func (o ParametersLinkPtrOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParametersLink) *string {
		if v == nil {
			return nil
		}
		return v.ContentVersion
	}).(pulumi.StringPtrOutput)
}

func (o ParametersLinkPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParametersLink) *string {
		if v == nil {
			return nil
		}
		return &v.Uri
	}).(pulumi.StringPtrOutput)
}

type ParametersLinkResponse struct {
	ContentVersion *string `pulumi:"contentVersion"`
	Uri            string  `pulumi:"uri"`
}

type ParametersLinkResponseOutput struct{ *pulumi.OutputState }

func (ParametersLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParametersLinkResponse)(nil)).Elem()
}

func (o ParametersLinkResponseOutput) ToParametersLinkResponseOutput() ParametersLinkResponseOutput {
	return o
}

func (o ParametersLinkResponseOutput) ToParametersLinkResponseOutputWithContext(ctx context.Context) ParametersLinkResponseOutput {
	return o
}

func (o ParametersLinkResponseOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParametersLinkResponse) *string { return v.ContentVersion }).(pulumi.StringPtrOutput)
}

func (o ParametersLinkResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v ParametersLinkResponse) string { return v.Uri }).(pulumi.StringOutput)
}

type ParametersLinkResponsePtrOutput struct{ *pulumi.OutputState }

func (ParametersLinkResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParametersLinkResponse)(nil)).Elem()
}

func (o ParametersLinkResponsePtrOutput) ToParametersLinkResponsePtrOutput() ParametersLinkResponsePtrOutput {
	return o
}

func (o ParametersLinkResponsePtrOutput) ToParametersLinkResponsePtrOutputWithContext(ctx context.Context) ParametersLinkResponsePtrOutput {
	return o
}

func (o ParametersLinkResponsePtrOutput) Elem() ParametersLinkResponseOutput {
	return o.ApplyT(func(v *ParametersLinkResponse) ParametersLinkResponse {
		if v != nil {
			return *v
		}
		var ret ParametersLinkResponse
		return ret
	}).(ParametersLinkResponseOutput)
}

func (o ParametersLinkResponsePtrOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParametersLinkResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContentVersion
	}).(pulumi.StringPtrOutput)
}

func (o ParametersLinkResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParametersLinkResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Uri
	}).(pulumi.StringPtrOutput)
}

type Plan struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
}





type PlanInput interface {
	pulumi.Input

	ToPlanOutput() PlanOutput
	ToPlanOutputWithContext(context.Context) PlanOutput
}

type PlanArgs struct {
	Name          pulumi.StringPtrInput `pulumi:"name"`
	Product       pulumi.StringPtrInput `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringPtrInput `pulumi:"publisher"`
}

func (PlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (i PlanArgs) ToPlanOutput() PlanOutput {
	return i.ToPlanOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput)
}

func (i PlanArgs) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput).ToPlanPtrOutputWithContext(ctx)
}









type PlanPtrInput interface {
	pulumi.Input

	ToPlanPtrOutput() PlanPtrOutput
	ToPlanPtrOutputWithContext(context.Context) PlanPtrOutput
}

type planPtrType PlanArgs

func PlanPtr(v *PlanArgs) PlanPtrInput {
	return (*planPtrType)(v)
}

func (*planPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (i *planPtrType) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i *planPtrType) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanPtrOutput)
}

type PlanOutput struct{ *pulumi.OutputState }

func (PlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (o PlanOutput) ToPlanOutput() PlanOutput {
	return o
}

func (o PlanOutput) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return o
}

func (o PlanOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o.ToPlanPtrOutputWithContext(context.Background())
}

func (o PlanOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Plan) *Plan {
		return &v
	}).(PlanPtrOutput)
}

func (o PlanOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PlanOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o PlanOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PlanOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type PlanPtrOutput struct{ *pulumi.OutputState }

func (PlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (o PlanPtrOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) Elem() PlanOutput {
	return o.ApplyT(func(v *Plan) Plan {
		if v != nil {
			return *v
		}
		var ret Plan
		return ret
	}).(PlanOutput)
}

func (o PlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

type PlanResponse struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
}

type PlanResponseOutput struct{ *pulumi.OutputState }

func (PlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanResponse)(nil)).Elem()
}

func (o PlanResponseOutput) ToPlanResponseOutput() PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) ToPlanResponseOutputWithContext(ctx context.Context) PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type PlanResponsePtrOutput struct{ *pulumi.OutputState }

func (PlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanResponse)(nil)).Elem()
}

func (o PlanResponsePtrOutput) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return o
}

func (o PlanResponsePtrOutput) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return o
}

func (o PlanResponsePtrOutput) Elem() PlanResponseOutput {
	return o.ApplyT(func(v *PlanResponse) PlanResponse {
		if v != nil {
			return *v
		}
		var ret PlanResponse
		return ret
	}).(PlanResponseOutput)
}

func (o PlanResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

type ProviderResourceTypeResponse struct {
	ApiVersions  []string          `pulumi:"apiVersions"`
	Locations    []string          `pulumi:"locations"`
	Properties   map[string]string `pulumi:"properties"`
	ResourceType *string           `pulumi:"resourceType"`
}

type ProviderResourceTypeResponseOutput struct{ *pulumi.OutputState }

func (ProviderResourceTypeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderResourceTypeResponse)(nil)).Elem()
}

func (o ProviderResourceTypeResponseOutput) ToProviderResourceTypeResponseOutput() ProviderResourceTypeResponseOutput {
	return o
}

func (o ProviderResourceTypeResponseOutput) ToProviderResourceTypeResponseOutputWithContext(ctx context.Context) ProviderResourceTypeResponseOutput {
	return o
}

func (o ProviderResourceTypeResponseOutput) ApiVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ProviderResourceTypeResponse) []string { return v.ApiVersions }).(pulumi.StringArrayOutput)
}

func (o ProviderResourceTypeResponseOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ProviderResourceTypeResponse) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o ProviderResourceTypeResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ProviderResourceTypeResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ProviderResourceTypeResponseOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResourceTypeResponse) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

type ProviderResourceTypeResponseArrayOutput struct{ *pulumi.OutputState }

func (ProviderResourceTypeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderResourceTypeResponse)(nil)).Elem()
}

func (o ProviderResourceTypeResponseArrayOutput) ToProviderResourceTypeResponseArrayOutput() ProviderResourceTypeResponseArrayOutput {
	return o
}

func (o ProviderResourceTypeResponseArrayOutput) ToProviderResourceTypeResponseArrayOutputWithContext(ctx context.Context) ProviderResourceTypeResponseArrayOutput {
	return o
}

func (o ProviderResourceTypeResponseArrayOutput) Index(i pulumi.IntInput) ProviderResourceTypeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProviderResourceTypeResponse {
		return vs[0].([]ProviderResourceTypeResponse)[vs[1].(int)]
	}).(ProviderResourceTypeResponseOutput)
}

type ProviderResponse struct {
	Id                *string                        `pulumi:"id"`
	Namespace         *string                        `pulumi:"namespace"`
	RegistrationState *string                        `pulumi:"registrationState"`
	ResourceTypes     []ProviderResourceTypeResponse `pulumi:"resourceTypes"`
}

type ProviderResponseOutput struct{ *pulumi.OutputState }

func (ProviderResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderResponse)(nil)).Elem()
}

func (o ProviderResponseOutput) ToProviderResponseOutput() ProviderResponseOutput {
	return o
}

func (o ProviderResponseOutput) ToProviderResponseOutputWithContext(ctx context.Context) ProviderResponseOutput {
	return o
}

func (o ProviderResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ProviderResponseOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResponse) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o ProviderResponseOutput) RegistrationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResponse) *string { return v.RegistrationState }).(pulumi.StringPtrOutput)
}

func (o ProviderResponseOutput) ResourceTypes() ProviderResourceTypeResponseArrayOutput {
	return o.ApplyT(func(v ProviderResponse) []ProviderResourceTypeResponse { return v.ResourceTypes }).(ProviderResourceTypeResponseArrayOutput)
}

type ProviderResponseArrayOutput struct{ *pulumi.OutputState }

func (ProviderResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderResponse)(nil)).Elem()
}

func (o ProviderResponseArrayOutput) ToProviderResponseArrayOutput() ProviderResponseArrayOutput {
	return o
}

func (o ProviderResponseArrayOutput) ToProviderResponseArrayOutputWithContext(ctx context.Context) ProviderResponseArrayOutput {
	return o
}

func (o ProviderResponseArrayOutput) Index(i pulumi.IntInput) ProviderResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProviderResponse {
		return vs[0].([]ProviderResponse)[vs[1].(int)]
	}).(ProviderResponseOutput)
}

type ResourceGroupPropertiesResponse struct {
	ProvisioningState string `pulumi:"provisioningState"`
}

type ResourceGroupPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ResourceGroupPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupPropertiesResponse)(nil)).Elem()
}

func (o ResourceGroupPropertiesResponseOutput) ToResourceGroupPropertiesResponseOutput() ResourceGroupPropertiesResponseOutput {
	return o
}

func (o ResourceGroupPropertiesResponseOutput) ToResourceGroupPropertiesResponseOutputWithContext(ctx context.Context) ResourceGroupPropertiesResponseOutput {
	return o
}

func (o ResourceGroupPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceGroupPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type TemplateLink struct {
	ContentVersion *string `pulumi:"contentVersion"`
	Uri            string  `pulumi:"uri"`
}





type TemplateLinkInput interface {
	pulumi.Input

	ToTemplateLinkOutput() TemplateLinkOutput
	ToTemplateLinkOutputWithContext(context.Context) TemplateLinkOutput
}

type TemplateLinkArgs struct {
	ContentVersion pulumi.StringPtrInput `pulumi:"contentVersion"`
	Uri            pulumi.StringInput    `pulumi:"uri"`
}

func (TemplateLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateLink)(nil)).Elem()
}

func (i TemplateLinkArgs) ToTemplateLinkOutput() TemplateLinkOutput {
	return i.ToTemplateLinkOutputWithContext(context.Background())
}

func (i TemplateLinkArgs) ToTemplateLinkOutputWithContext(ctx context.Context) TemplateLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateLinkOutput)
}

func (i TemplateLinkArgs) ToTemplateLinkPtrOutput() TemplateLinkPtrOutput {
	return i.ToTemplateLinkPtrOutputWithContext(context.Background())
}

func (i TemplateLinkArgs) ToTemplateLinkPtrOutputWithContext(ctx context.Context) TemplateLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateLinkOutput).ToTemplateLinkPtrOutputWithContext(ctx)
}









type TemplateLinkPtrInput interface {
	pulumi.Input

	ToTemplateLinkPtrOutput() TemplateLinkPtrOutput
	ToTemplateLinkPtrOutputWithContext(context.Context) TemplateLinkPtrOutput
}

type templateLinkPtrType TemplateLinkArgs

func TemplateLinkPtr(v *TemplateLinkArgs) TemplateLinkPtrInput {
	return (*templateLinkPtrType)(v)
}

func (*templateLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateLink)(nil)).Elem()
}

func (i *templateLinkPtrType) ToTemplateLinkPtrOutput() TemplateLinkPtrOutput {
	return i.ToTemplateLinkPtrOutputWithContext(context.Background())
}

func (i *templateLinkPtrType) ToTemplateLinkPtrOutputWithContext(ctx context.Context) TemplateLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateLinkPtrOutput)
}

type TemplateLinkOutput struct{ *pulumi.OutputState }

func (TemplateLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateLink)(nil)).Elem()
}

func (o TemplateLinkOutput) ToTemplateLinkOutput() TemplateLinkOutput {
	return o
}

func (o TemplateLinkOutput) ToTemplateLinkOutputWithContext(ctx context.Context) TemplateLinkOutput {
	return o
}

func (o TemplateLinkOutput) ToTemplateLinkPtrOutput() TemplateLinkPtrOutput {
	return o.ToTemplateLinkPtrOutputWithContext(context.Background())
}

func (o TemplateLinkOutput) ToTemplateLinkPtrOutputWithContext(ctx context.Context) TemplateLinkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TemplateLink) *TemplateLink {
		return &v
	}).(TemplateLinkPtrOutput)
}

func (o TemplateLinkOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TemplateLink) *string { return v.ContentVersion }).(pulumi.StringPtrOutput)
}

func (o TemplateLinkOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateLink) string { return v.Uri }).(pulumi.StringOutput)
}

type TemplateLinkPtrOutput struct{ *pulumi.OutputState }

func (TemplateLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateLink)(nil)).Elem()
}

func (o TemplateLinkPtrOutput) ToTemplateLinkPtrOutput() TemplateLinkPtrOutput {
	return o
}

func (o TemplateLinkPtrOutput) ToTemplateLinkPtrOutputWithContext(ctx context.Context) TemplateLinkPtrOutput {
	return o
}

func (o TemplateLinkPtrOutput) Elem() TemplateLinkOutput {
	return o.ApplyT(func(v *TemplateLink) TemplateLink {
		if v != nil {
			return *v
		}
		var ret TemplateLink
		return ret
	}).(TemplateLinkOutput)
}

func (o TemplateLinkPtrOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateLink) *string {
		if v == nil {
			return nil
		}
		return v.ContentVersion
	}).(pulumi.StringPtrOutput)
}

func (o TemplateLinkPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateLink) *string {
		if v == nil {
			return nil
		}
		return &v.Uri
	}).(pulumi.StringPtrOutput)
}

type TemplateLinkResponse struct {
	ContentVersion *string `pulumi:"contentVersion"`
	Uri            string  `pulumi:"uri"`
}

type TemplateLinkResponseOutput struct{ *pulumi.OutputState }

func (TemplateLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateLinkResponse)(nil)).Elem()
}

func (o TemplateLinkResponseOutput) ToTemplateLinkResponseOutput() TemplateLinkResponseOutput {
	return o
}

func (o TemplateLinkResponseOutput) ToTemplateLinkResponseOutputWithContext(ctx context.Context) TemplateLinkResponseOutput {
	return o
}

func (o TemplateLinkResponseOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TemplateLinkResponse) *string { return v.ContentVersion }).(pulumi.StringPtrOutput)
}

func (o TemplateLinkResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateLinkResponse) string { return v.Uri }).(pulumi.StringOutput)
}

type TemplateLinkResponsePtrOutput struct{ *pulumi.OutputState }

func (TemplateLinkResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateLinkResponse)(nil)).Elem()
}

func (o TemplateLinkResponsePtrOutput) ToTemplateLinkResponsePtrOutput() TemplateLinkResponsePtrOutput {
	return o
}

func (o TemplateLinkResponsePtrOutput) ToTemplateLinkResponsePtrOutputWithContext(ctx context.Context) TemplateLinkResponsePtrOutput {
	return o
}

func (o TemplateLinkResponsePtrOutput) Elem() TemplateLinkResponseOutput {
	return o.ApplyT(func(v *TemplateLinkResponse) TemplateLinkResponse {
		if v != nil {
			return *v
		}
		var ret TemplateLinkResponse
		return ret
	}).(TemplateLinkResponseOutput)
}

func (o TemplateLinkResponsePtrOutput) ContentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateLinkResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContentVersion
	}).(pulumi.StringPtrOutput)
}

func (o TemplateLinkResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateLinkResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Uri
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BasicDependencyResponseOutput{})
	pulumi.RegisterOutputType(BasicDependencyResponseArrayOutput{})
	pulumi.RegisterOutputType(DependencyResponseOutput{})
	pulumi.RegisterOutputType(DependencyResponseArrayOutput{})
	pulumi.RegisterOutputType(DeploymentPropertiesOutput{})
	pulumi.RegisterOutputType(DeploymentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(DeploymentPropertiesExtendedResponseOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseOutput{})
	pulumi.RegisterOutputType(ErrorAdditionalInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponseOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponseArrayOutput{})
	pulumi.RegisterOutputType(ParametersLinkOutput{})
	pulumi.RegisterOutputType(ParametersLinkPtrOutput{})
	pulumi.RegisterOutputType(ParametersLinkResponseOutput{})
	pulumi.RegisterOutputType(ParametersLinkResponsePtrOutput{})
	pulumi.RegisterOutputType(PlanOutput{})
	pulumi.RegisterOutputType(PlanPtrOutput{})
	pulumi.RegisterOutputType(PlanResponseOutput{})
	pulumi.RegisterOutputType(PlanResponsePtrOutput{})
	pulumi.RegisterOutputType(ProviderResourceTypeResponseOutput{})
	pulumi.RegisterOutputType(ProviderResourceTypeResponseArrayOutput{})
	pulumi.RegisterOutputType(ProviderResponseOutput{})
	pulumi.RegisterOutputType(ProviderResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceGroupPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TemplateLinkOutput{})
	pulumi.RegisterOutputType(TemplateLinkPtrOutput{})
	pulumi.RegisterOutputType(TemplateLinkResponseOutput{})
	pulumi.RegisterOutputType(TemplateLinkResponsePtrOutput{})
}
