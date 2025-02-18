


package v20220201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LinkedTemplateArtifact struct {
	Path     string      `pulumi:"path"`
	Template interface{} `pulumi:"template"`
}





type LinkedTemplateArtifactInput interface {
	pulumi.Input

	ToLinkedTemplateArtifactOutput() LinkedTemplateArtifactOutput
	ToLinkedTemplateArtifactOutputWithContext(context.Context) LinkedTemplateArtifactOutput
}

type LinkedTemplateArtifactArgs struct {
	Path     pulumi.StringInput `pulumi:"path"`
	Template pulumi.Input       `pulumi:"template"`
}

func (LinkedTemplateArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedTemplateArtifact)(nil)).Elem()
}

func (i LinkedTemplateArtifactArgs) ToLinkedTemplateArtifactOutput() LinkedTemplateArtifactOutput {
	return i.ToLinkedTemplateArtifactOutputWithContext(context.Background())
}

func (i LinkedTemplateArtifactArgs) ToLinkedTemplateArtifactOutputWithContext(ctx context.Context) LinkedTemplateArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedTemplateArtifactOutput)
}





type LinkedTemplateArtifactArrayInput interface {
	pulumi.Input

	ToLinkedTemplateArtifactArrayOutput() LinkedTemplateArtifactArrayOutput
	ToLinkedTemplateArtifactArrayOutputWithContext(context.Context) LinkedTemplateArtifactArrayOutput
}

type LinkedTemplateArtifactArray []LinkedTemplateArtifactInput

func (LinkedTemplateArtifactArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedTemplateArtifact)(nil)).Elem()
}

func (i LinkedTemplateArtifactArray) ToLinkedTemplateArtifactArrayOutput() LinkedTemplateArtifactArrayOutput {
	return i.ToLinkedTemplateArtifactArrayOutputWithContext(context.Background())
}

func (i LinkedTemplateArtifactArray) ToLinkedTemplateArtifactArrayOutputWithContext(ctx context.Context) LinkedTemplateArtifactArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedTemplateArtifactArrayOutput)
}

type LinkedTemplateArtifactOutput struct{ *pulumi.OutputState }

func (LinkedTemplateArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedTemplateArtifact)(nil)).Elem()
}

func (o LinkedTemplateArtifactOutput) ToLinkedTemplateArtifactOutput() LinkedTemplateArtifactOutput {
	return o
}

func (o LinkedTemplateArtifactOutput) ToLinkedTemplateArtifactOutputWithContext(ctx context.Context) LinkedTemplateArtifactOutput {
	return o
}

func (o LinkedTemplateArtifactOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LinkedTemplateArtifact) string { return v.Path }).(pulumi.StringOutput)
}

func (o LinkedTemplateArtifactOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v LinkedTemplateArtifact) interface{} { return v.Template }).(pulumi.AnyOutput)
}

type LinkedTemplateArtifactArrayOutput struct{ *pulumi.OutputState }

func (LinkedTemplateArtifactArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedTemplateArtifact)(nil)).Elem()
}

func (o LinkedTemplateArtifactArrayOutput) ToLinkedTemplateArtifactArrayOutput() LinkedTemplateArtifactArrayOutput {
	return o
}

func (o LinkedTemplateArtifactArrayOutput) ToLinkedTemplateArtifactArrayOutputWithContext(ctx context.Context) LinkedTemplateArtifactArrayOutput {
	return o
}

func (o LinkedTemplateArtifactArrayOutput) Index(i pulumi.IntInput) LinkedTemplateArtifactOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedTemplateArtifact {
		return vs[0].([]LinkedTemplateArtifact)[vs[1].(int)]
	}).(LinkedTemplateArtifactOutput)
}

type LinkedTemplateArtifactResponse struct {
	Path     string      `pulumi:"path"`
	Template interface{} `pulumi:"template"`
}

type LinkedTemplateArtifactResponseOutput struct{ *pulumi.OutputState }

func (LinkedTemplateArtifactResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedTemplateArtifactResponse)(nil)).Elem()
}

func (o LinkedTemplateArtifactResponseOutput) ToLinkedTemplateArtifactResponseOutput() LinkedTemplateArtifactResponseOutput {
	return o
}

func (o LinkedTemplateArtifactResponseOutput) ToLinkedTemplateArtifactResponseOutputWithContext(ctx context.Context) LinkedTemplateArtifactResponseOutput {
	return o
}

func (o LinkedTemplateArtifactResponseOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LinkedTemplateArtifactResponse) string { return v.Path }).(pulumi.StringOutput)
}

func (o LinkedTemplateArtifactResponseOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v LinkedTemplateArtifactResponse) interface{} { return v.Template }).(pulumi.AnyOutput)
}

type LinkedTemplateArtifactResponseArrayOutput struct{ *pulumi.OutputState }

func (LinkedTemplateArtifactResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedTemplateArtifactResponse)(nil)).Elem()
}

func (o LinkedTemplateArtifactResponseArrayOutput) ToLinkedTemplateArtifactResponseArrayOutput() LinkedTemplateArtifactResponseArrayOutput {
	return o
}

func (o LinkedTemplateArtifactResponseArrayOutput) ToLinkedTemplateArtifactResponseArrayOutputWithContext(ctx context.Context) LinkedTemplateArtifactResponseArrayOutput {
	return o
}

func (o LinkedTemplateArtifactResponseArrayOutput) Index(i pulumi.IntInput) LinkedTemplateArtifactResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedTemplateArtifactResponse {
		return vs[0].([]LinkedTemplateArtifactResponse)[vs[1].(int)]
	}).(LinkedTemplateArtifactResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type TemplateSpecVersionInfoResponse struct {
	Description  string `pulumi:"description"`
	TimeCreated  string `pulumi:"timeCreated"`
	TimeModified string `pulumi:"timeModified"`
}

type TemplateSpecVersionInfoResponseOutput struct{ *pulumi.OutputState }

func (TemplateSpecVersionInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateSpecVersionInfoResponse)(nil)).Elem()
}

func (o TemplateSpecVersionInfoResponseOutput) ToTemplateSpecVersionInfoResponseOutput() TemplateSpecVersionInfoResponseOutput {
	return o
}

func (o TemplateSpecVersionInfoResponseOutput) ToTemplateSpecVersionInfoResponseOutputWithContext(ctx context.Context) TemplateSpecVersionInfoResponseOutput {
	return o
}

func (o TemplateSpecVersionInfoResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecVersionInfoResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o TemplateSpecVersionInfoResponseOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecVersionInfoResponse) string { return v.TimeCreated }).(pulumi.StringOutput)
}

func (o TemplateSpecVersionInfoResponseOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v TemplateSpecVersionInfoResponse) string { return v.TimeModified }).(pulumi.StringOutput)
}

type TemplateSpecVersionInfoResponseMapOutput struct{ *pulumi.OutputState }

func (TemplateSpecVersionInfoResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TemplateSpecVersionInfoResponse)(nil)).Elem()
}

func (o TemplateSpecVersionInfoResponseMapOutput) ToTemplateSpecVersionInfoResponseMapOutput() TemplateSpecVersionInfoResponseMapOutput {
	return o
}

func (o TemplateSpecVersionInfoResponseMapOutput) ToTemplateSpecVersionInfoResponseMapOutputWithContext(ctx context.Context) TemplateSpecVersionInfoResponseMapOutput {
	return o
}

func (o TemplateSpecVersionInfoResponseMapOutput) MapIndex(k pulumi.StringInput) TemplateSpecVersionInfoResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TemplateSpecVersionInfoResponse {
		return vs[0].(map[string]TemplateSpecVersionInfoResponse)[vs[1].(string)]
	}).(TemplateSpecVersionInfoResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedTemplateArtifactOutput{})
	pulumi.RegisterOutputType(LinkedTemplateArtifactArrayOutput{})
	pulumi.RegisterOutputType(LinkedTemplateArtifactResponseOutput{})
	pulumi.RegisterOutputType(LinkedTemplateArtifactResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TemplateSpecVersionInfoResponseOutput{})
	pulumi.RegisterOutputType(TemplateSpecVersionInfoResponseMapOutput{})
}
