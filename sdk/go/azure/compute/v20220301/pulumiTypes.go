


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdditionalCapabilities struct {
	HibernationEnabled *bool `pulumi:"hibernationEnabled"`
	UltraSSDEnabled    *bool `pulumi:"ultraSSDEnabled"`
}





type AdditionalCapabilitiesInput interface {
	pulumi.Input

	ToAdditionalCapabilitiesOutput() AdditionalCapabilitiesOutput
	ToAdditionalCapabilitiesOutputWithContext(context.Context) AdditionalCapabilitiesOutput
}

type AdditionalCapabilitiesArgs struct {
	HibernationEnabled pulumi.BoolPtrInput `pulumi:"hibernationEnabled"`
	UltraSSDEnabled    pulumi.BoolPtrInput `pulumi:"ultraSSDEnabled"`
}

func (AdditionalCapabilitiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalCapabilities)(nil)).Elem()
}

func (i AdditionalCapabilitiesArgs) ToAdditionalCapabilitiesOutput() AdditionalCapabilitiesOutput {
	return i.ToAdditionalCapabilitiesOutputWithContext(context.Background())
}

func (i AdditionalCapabilitiesArgs) ToAdditionalCapabilitiesOutputWithContext(ctx context.Context) AdditionalCapabilitiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalCapabilitiesOutput)
}

func (i AdditionalCapabilitiesArgs) ToAdditionalCapabilitiesPtrOutput() AdditionalCapabilitiesPtrOutput {
	return i.ToAdditionalCapabilitiesPtrOutputWithContext(context.Background())
}

func (i AdditionalCapabilitiesArgs) ToAdditionalCapabilitiesPtrOutputWithContext(ctx context.Context) AdditionalCapabilitiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalCapabilitiesOutput).ToAdditionalCapabilitiesPtrOutputWithContext(ctx)
}









type AdditionalCapabilitiesPtrInput interface {
	pulumi.Input

	ToAdditionalCapabilitiesPtrOutput() AdditionalCapabilitiesPtrOutput
	ToAdditionalCapabilitiesPtrOutputWithContext(context.Context) AdditionalCapabilitiesPtrOutput
}

type additionalCapabilitiesPtrType AdditionalCapabilitiesArgs

func AdditionalCapabilitiesPtr(v *AdditionalCapabilitiesArgs) AdditionalCapabilitiesPtrInput {
	return (*additionalCapabilitiesPtrType)(v)
}

func (*additionalCapabilitiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AdditionalCapabilities)(nil)).Elem()
}

func (i *additionalCapabilitiesPtrType) ToAdditionalCapabilitiesPtrOutput() AdditionalCapabilitiesPtrOutput {
	return i.ToAdditionalCapabilitiesPtrOutputWithContext(context.Background())
}

func (i *additionalCapabilitiesPtrType) ToAdditionalCapabilitiesPtrOutputWithContext(ctx context.Context) AdditionalCapabilitiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalCapabilitiesPtrOutput)
}

type AdditionalCapabilitiesOutput struct{ *pulumi.OutputState }

func (AdditionalCapabilitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalCapabilities)(nil)).Elem()
}

func (o AdditionalCapabilitiesOutput) ToAdditionalCapabilitiesOutput() AdditionalCapabilitiesOutput {
	return o
}

func (o AdditionalCapabilitiesOutput) ToAdditionalCapabilitiesOutputWithContext(ctx context.Context) AdditionalCapabilitiesOutput {
	return o
}

func (o AdditionalCapabilitiesOutput) ToAdditionalCapabilitiesPtrOutput() AdditionalCapabilitiesPtrOutput {
	return o.ToAdditionalCapabilitiesPtrOutputWithContext(context.Background())
}

func (o AdditionalCapabilitiesOutput) ToAdditionalCapabilitiesPtrOutputWithContext(ctx context.Context) AdditionalCapabilitiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AdditionalCapabilities) *AdditionalCapabilities {
		return &v
	}).(AdditionalCapabilitiesPtrOutput)
}

func (o AdditionalCapabilitiesOutput) HibernationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AdditionalCapabilities) *bool { return v.HibernationEnabled }).(pulumi.BoolPtrOutput)
}

func (o AdditionalCapabilitiesOutput) UltraSSDEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AdditionalCapabilities) *bool { return v.UltraSSDEnabled }).(pulumi.BoolPtrOutput)
}

type AdditionalCapabilitiesPtrOutput struct{ *pulumi.OutputState }

func (AdditionalCapabilitiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdditionalCapabilities)(nil)).Elem()
}

func (o AdditionalCapabilitiesPtrOutput) ToAdditionalCapabilitiesPtrOutput() AdditionalCapabilitiesPtrOutput {
	return o
}

func (o AdditionalCapabilitiesPtrOutput) ToAdditionalCapabilitiesPtrOutputWithContext(ctx context.Context) AdditionalCapabilitiesPtrOutput {
	return o
}

func (o AdditionalCapabilitiesPtrOutput) Elem() AdditionalCapabilitiesOutput {
	return o.ApplyT(func(v *AdditionalCapabilities) AdditionalCapabilities {
		if v != nil {
			return *v
		}
		var ret AdditionalCapabilities
		return ret
	}).(AdditionalCapabilitiesOutput)
}

func (o AdditionalCapabilitiesPtrOutput) HibernationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AdditionalCapabilities) *bool {
		if v == nil {
			return nil
		}
		return v.HibernationEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o AdditionalCapabilitiesPtrOutput) UltraSSDEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AdditionalCapabilities) *bool {
		if v == nil {
			return nil
		}
		return v.UltraSSDEnabled
	}).(pulumi.BoolPtrOutput)
}

type AdditionalCapabilitiesResponse struct {
	HibernationEnabled *bool `pulumi:"hibernationEnabled"`
	UltraSSDEnabled    *bool `pulumi:"ultraSSDEnabled"`
}

type AdditionalCapabilitiesResponseOutput struct{ *pulumi.OutputState }

func (AdditionalCapabilitiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalCapabilitiesResponse)(nil)).Elem()
}

func (o AdditionalCapabilitiesResponseOutput) ToAdditionalCapabilitiesResponseOutput() AdditionalCapabilitiesResponseOutput {
	return o
}

func (o AdditionalCapabilitiesResponseOutput) ToAdditionalCapabilitiesResponseOutputWithContext(ctx context.Context) AdditionalCapabilitiesResponseOutput {
	return o
}

func (o AdditionalCapabilitiesResponseOutput) HibernationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AdditionalCapabilitiesResponse) *bool { return v.HibernationEnabled }).(pulumi.BoolPtrOutput)
}

func (o AdditionalCapabilitiesResponseOutput) UltraSSDEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AdditionalCapabilitiesResponse) *bool { return v.UltraSSDEnabled }).(pulumi.BoolPtrOutput)
}

type AdditionalCapabilitiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AdditionalCapabilitiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdditionalCapabilitiesResponse)(nil)).Elem()
}

func (o AdditionalCapabilitiesResponsePtrOutput) ToAdditionalCapabilitiesResponsePtrOutput() AdditionalCapabilitiesResponsePtrOutput {
	return o
}

func (o AdditionalCapabilitiesResponsePtrOutput) ToAdditionalCapabilitiesResponsePtrOutputWithContext(ctx context.Context) AdditionalCapabilitiesResponsePtrOutput {
	return o
}

func (o AdditionalCapabilitiesResponsePtrOutput) Elem() AdditionalCapabilitiesResponseOutput {
	return o.ApplyT(func(v *AdditionalCapabilitiesResponse) AdditionalCapabilitiesResponse {
		if v != nil {
			return *v
		}
		var ret AdditionalCapabilitiesResponse
		return ret
	}).(AdditionalCapabilitiesResponseOutput)
}

func (o AdditionalCapabilitiesResponsePtrOutput) HibernationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AdditionalCapabilitiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.HibernationEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o AdditionalCapabilitiesResponsePtrOutput) UltraSSDEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AdditionalCapabilitiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.UltraSSDEnabled
	}).(pulumi.BoolPtrOutput)
}

type AdditionalUnattendContent struct {
	ComponentName *ComponentNames `pulumi:"componentName"`
	Content       *string         `pulumi:"content"`
	PassName      *PassNames      `pulumi:"passName"`
	SettingName   *SettingNames   `pulumi:"settingName"`
}





type AdditionalUnattendContentInput interface {
	pulumi.Input

	ToAdditionalUnattendContentOutput() AdditionalUnattendContentOutput
	ToAdditionalUnattendContentOutputWithContext(context.Context) AdditionalUnattendContentOutput
}

type AdditionalUnattendContentArgs struct {
	ComponentName ComponentNamesPtrInput `pulumi:"componentName"`
	Content       pulumi.StringPtrInput  `pulumi:"content"`
	PassName      PassNamesPtrInput      `pulumi:"passName"`
	SettingName   SettingNamesPtrInput   `pulumi:"settingName"`
}

func (AdditionalUnattendContentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalUnattendContent)(nil)).Elem()
}

func (i AdditionalUnattendContentArgs) ToAdditionalUnattendContentOutput() AdditionalUnattendContentOutput {
	return i.ToAdditionalUnattendContentOutputWithContext(context.Background())
}

func (i AdditionalUnattendContentArgs) ToAdditionalUnattendContentOutputWithContext(ctx context.Context) AdditionalUnattendContentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalUnattendContentOutput)
}





type AdditionalUnattendContentArrayInput interface {
	pulumi.Input

	ToAdditionalUnattendContentArrayOutput() AdditionalUnattendContentArrayOutput
	ToAdditionalUnattendContentArrayOutputWithContext(context.Context) AdditionalUnattendContentArrayOutput
}

type AdditionalUnattendContentArray []AdditionalUnattendContentInput

func (AdditionalUnattendContentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalUnattendContent)(nil)).Elem()
}

func (i AdditionalUnattendContentArray) ToAdditionalUnattendContentArrayOutput() AdditionalUnattendContentArrayOutput {
	return i.ToAdditionalUnattendContentArrayOutputWithContext(context.Background())
}

func (i AdditionalUnattendContentArray) ToAdditionalUnattendContentArrayOutputWithContext(ctx context.Context) AdditionalUnattendContentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalUnattendContentArrayOutput)
}

type AdditionalUnattendContentOutput struct{ *pulumi.OutputState }

func (AdditionalUnattendContentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalUnattendContent)(nil)).Elem()
}

func (o AdditionalUnattendContentOutput) ToAdditionalUnattendContentOutput() AdditionalUnattendContentOutput {
	return o
}

func (o AdditionalUnattendContentOutput) ToAdditionalUnattendContentOutputWithContext(ctx context.Context) AdditionalUnattendContentOutput {
	return o
}

func (o AdditionalUnattendContentOutput) ComponentName() ComponentNamesPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContent) *ComponentNames { return v.ComponentName }).(ComponentNamesPtrOutput)
}

func (o AdditionalUnattendContentOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContent) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o AdditionalUnattendContentOutput) PassName() PassNamesPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContent) *PassNames { return v.PassName }).(PassNamesPtrOutput)
}

func (o AdditionalUnattendContentOutput) SettingName() SettingNamesPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContent) *SettingNames { return v.SettingName }).(SettingNamesPtrOutput)
}

type AdditionalUnattendContentArrayOutput struct{ *pulumi.OutputState }

func (AdditionalUnattendContentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalUnattendContent)(nil)).Elem()
}

func (o AdditionalUnattendContentArrayOutput) ToAdditionalUnattendContentArrayOutput() AdditionalUnattendContentArrayOutput {
	return o
}

func (o AdditionalUnattendContentArrayOutput) ToAdditionalUnattendContentArrayOutputWithContext(ctx context.Context) AdditionalUnattendContentArrayOutput {
	return o
}

func (o AdditionalUnattendContentArrayOutput) Index(i pulumi.IntInput) AdditionalUnattendContentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdditionalUnattendContent {
		return vs[0].([]AdditionalUnattendContent)[vs[1].(int)]
	}).(AdditionalUnattendContentOutput)
}

type AdditionalUnattendContentResponse struct {
	ComponentName *string `pulumi:"componentName"`
	Content       *string `pulumi:"content"`
	PassName      *string `pulumi:"passName"`
	SettingName   *string `pulumi:"settingName"`
}

type AdditionalUnattendContentResponseOutput struct{ *pulumi.OutputState }

func (AdditionalUnattendContentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalUnattendContentResponse)(nil)).Elem()
}

func (o AdditionalUnattendContentResponseOutput) ToAdditionalUnattendContentResponseOutput() AdditionalUnattendContentResponseOutput {
	return o
}

func (o AdditionalUnattendContentResponseOutput) ToAdditionalUnattendContentResponseOutputWithContext(ctx context.Context) AdditionalUnattendContentResponseOutput {
	return o
}

func (o AdditionalUnattendContentResponseOutput) ComponentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContentResponse) *string { return v.ComponentName }).(pulumi.StringPtrOutput)
}

func (o AdditionalUnattendContentResponseOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContentResponse) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o AdditionalUnattendContentResponseOutput) PassName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContentResponse) *string { return v.PassName }).(pulumi.StringPtrOutput)
}

func (o AdditionalUnattendContentResponseOutput) SettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdditionalUnattendContentResponse) *string { return v.SettingName }).(pulumi.StringPtrOutput)
}

type AdditionalUnattendContentResponseArrayOutput struct{ *pulumi.OutputState }

func (AdditionalUnattendContentResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdditionalUnattendContentResponse)(nil)).Elem()
}

func (o AdditionalUnattendContentResponseArrayOutput) ToAdditionalUnattendContentResponseArrayOutput() AdditionalUnattendContentResponseArrayOutput {
	return o
}

func (o AdditionalUnattendContentResponseArrayOutput) ToAdditionalUnattendContentResponseArrayOutputWithContext(ctx context.Context) AdditionalUnattendContentResponseArrayOutput {
	return o
}

func (o AdditionalUnattendContentResponseArrayOutput) Index(i pulumi.IntInput) AdditionalUnattendContentResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdditionalUnattendContentResponse {
		return vs[0].([]AdditionalUnattendContentResponse)[vs[1].(int)]
	}).(AdditionalUnattendContentResponseOutput)
}

type ApiEntityReference struct {
	Id *string `pulumi:"id"`
}





type ApiEntityReferenceInput interface {
	pulumi.Input

	ToApiEntityReferenceOutput() ApiEntityReferenceOutput
	ToApiEntityReferenceOutputWithContext(context.Context) ApiEntityReferenceOutput
}

type ApiEntityReferenceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ApiEntityReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEntityReference)(nil)).Elem()
}

func (i ApiEntityReferenceArgs) ToApiEntityReferenceOutput() ApiEntityReferenceOutput {
	return i.ToApiEntityReferenceOutputWithContext(context.Background())
}

func (i ApiEntityReferenceArgs) ToApiEntityReferenceOutputWithContext(ctx context.Context) ApiEntityReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityReferenceOutput)
}

func (i ApiEntityReferenceArgs) ToApiEntityReferencePtrOutput() ApiEntityReferencePtrOutput {
	return i.ToApiEntityReferencePtrOutputWithContext(context.Background())
}

func (i ApiEntityReferenceArgs) ToApiEntityReferencePtrOutputWithContext(ctx context.Context) ApiEntityReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityReferenceOutput).ToApiEntityReferencePtrOutputWithContext(ctx)
}









type ApiEntityReferencePtrInput interface {
	pulumi.Input

	ToApiEntityReferencePtrOutput() ApiEntityReferencePtrOutput
	ToApiEntityReferencePtrOutputWithContext(context.Context) ApiEntityReferencePtrOutput
}

type apiEntityReferencePtrType ApiEntityReferenceArgs

func ApiEntityReferencePtr(v *ApiEntityReferenceArgs) ApiEntityReferencePtrInput {
	return (*apiEntityReferencePtrType)(v)
}

func (*apiEntityReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiEntityReference)(nil)).Elem()
}

func (i *apiEntityReferencePtrType) ToApiEntityReferencePtrOutput() ApiEntityReferencePtrOutput {
	return i.ToApiEntityReferencePtrOutputWithContext(context.Background())
}

func (i *apiEntityReferencePtrType) ToApiEntityReferencePtrOutputWithContext(ctx context.Context) ApiEntityReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityReferencePtrOutput)
}





type ApiEntityReferenceArrayInput interface {
	pulumi.Input

	ToApiEntityReferenceArrayOutput() ApiEntityReferenceArrayOutput
	ToApiEntityReferenceArrayOutputWithContext(context.Context) ApiEntityReferenceArrayOutput
}

type ApiEntityReferenceArray []ApiEntityReferenceInput

func (ApiEntityReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiEntityReference)(nil)).Elem()
}

func (i ApiEntityReferenceArray) ToApiEntityReferenceArrayOutput() ApiEntityReferenceArrayOutput {
	return i.ToApiEntityReferenceArrayOutputWithContext(context.Background())
}

func (i ApiEntityReferenceArray) ToApiEntityReferenceArrayOutputWithContext(ctx context.Context) ApiEntityReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityReferenceArrayOutput)
}

type ApiEntityReferenceOutput struct{ *pulumi.OutputState }

func (ApiEntityReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEntityReference)(nil)).Elem()
}

func (o ApiEntityReferenceOutput) ToApiEntityReferenceOutput() ApiEntityReferenceOutput {
	return o
}

func (o ApiEntityReferenceOutput) ToApiEntityReferenceOutputWithContext(ctx context.Context) ApiEntityReferenceOutput {
	return o
}

func (o ApiEntityReferenceOutput) ToApiEntityReferencePtrOutput() ApiEntityReferencePtrOutput {
	return o.ToApiEntityReferencePtrOutputWithContext(context.Background())
}

func (o ApiEntityReferenceOutput) ToApiEntityReferencePtrOutputWithContext(ctx context.Context) ApiEntityReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiEntityReference) *ApiEntityReference {
		return &v
	}).(ApiEntityReferencePtrOutput)
}

func (o ApiEntityReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntityReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ApiEntityReferencePtrOutput struct{ *pulumi.OutputState }

func (ApiEntityReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiEntityReference)(nil)).Elem()
}

func (o ApiEntityReferencePtrOutput) ToApiEntityReferencePtrOutput() ApiEntityReferencePtrOutput {
	return o
}

func (o ApiEntityReferencePtrOutput) ToApiEntityReferencePtrOutputWithContext(ctx context.Context) ApiEntityReferencePtrOutput {
	return o
}

func (o ApiEntityReferencePtrOutput) Elem() ApiEntityReferenceOutput {
	return o.ApplyT(func(v *ApiEntityReference) ApiEntityReference {
		if v != nil {
			return *v
		}
		var ret ApiEntityReference
		return ret
	}).(ApiEntityReferenceOutput)
}

func (o ApiEntityReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntityReference) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ApiEntityReferenceArrayOutput struct{ *pulumi.OutputState }

func (ApiEntityReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiEntityReference)(nil)).Elem()
}

func (o ApiEntityReferenceArrayOutput) ToApiEntityReferenceArrayOutput() ApiEntityReferenceArrayOutput {
	return o
}

func (o ApiEntityReferenceArrayOutput) ToApiEntityReferenceArrayOutputWithContext(ctx context.Context) ApiEntityReferenceArrayOutput {
	return o
}

func (o ApiEntityReferenceArrayOutput) Index(i pulumi.IntInput) ApiEntityReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiEntityReference {
		return vs[0].([]ApiEntityReference)[vs[1].(int)]
	}).(ApiEntityReferenceOutput)
}

type ApiEntityReferenceResponse struct {
	Id *string `pulumi:"id"`
}

type ApiEntityReferenceResponseOutput struct{ *pulumi.OutputState }

func (ApiEntityReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEntityReferenceResponse)(nil)).Elem()
}

func (o ApiEntityReferenceResponseOutput) ToApiEntityReferenceResponseOutput() ApiEntityReferenceResponseOutput {
	return o
}

func (o ApiEntityReferenceResponseOutput) ToApiEntityReferenceResponseOutputWithContext(ctx context.Context) ApiEntityReferenceResponseOutput {
	return o
}

func (o ApiEntityReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntityReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ApiEntityReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiEntityReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiEntityReferenceResponse)(nil)).Elem()
}

func (o ApiEntityReferenceResponsePtrOutput) ToApiEntityReferenceResponsePtrOutput() ApiEntityReferenceResponsePtrOutput {
	return o
}

func (o ApiEntityReferenceResponsePtrOutput) ToApiEntityReferenceResponsePtrOutputWithContext(ctx context.Context) ApiEntityReferenceResponsePtrOutput {
	return o
}

func (o ApiEntityReferenceResponsePtrOutput) Elem() ApiEntityReferenceResponseOutput {
	return o.ApplyT(func(v *ApiEntityReferenceResponse) ApiEntityReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ApiEntityReferenceResponse
		return ret
	}).(ApiEntityReferenceResponseOutput)
}

func (o ApiEntityReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntityReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ApiEntityReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (ApiEntityReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiEntityReferenceResponse)(nil)).Elem()
}

func (o ApiEntityReferenceResponseArrayOutput) ToApiEntityReferenceResponseArrayOutput() ApiEntityReferenceResponseArrayOutput {
	return o
}

func (o ApiEntityReferenceResponseArrayOutput) ToApiEntityReferenceResponseArrayOutputWithContext(ctx context.Context) ApiEntityReferenceResponseArrayOutput {
	return o
}

func (o ApiEntityReferenceResponseArrayOutput) Index(i pulumi.IntInput) ApiEntityReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiEntityReferenceResponse {
		return vs[0].([]ApiEntityReferenceResponse)[vs[1].(int)]
	}).(ApiEntityReferenceResponseOutput)
}

type ApiErrorBaseResponse struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
	Target  *string `pulumi:"target"`
}

type ApiErrorBaseResponseOutput struct{ *pulumi.OutputState }

func (ApiErrorBaseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiErrorBaseResponse)(nil)).Elem()
}

func (o ApiErrorBaseResponseOutput) ToApiErrorBaseResponseOutput() ApiErrorBaseResponseOutput {
	return o
}

func (o ApiErrorBaseResponseOutput) ToApiErrorBaseResponseOutputWithContext(ctx context.Context) ApiErrorBaseResponseOutput {
	return o
}

func (o ApiErrorBaseResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiErrorBaseResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ApiErrorBaseResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiErrorBaseResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ApiErrorBaseResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiErrorBaseResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type ApiErrorBaseResponseArrayOutput struct{ *pulumi.OutputState }

func (ApiErrorBaseResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiErrorBaseResponse)(nil)).Elem()
}

func (o ApiErrorBaseResponseArrayOutput) ToApiErrorBaseResponseArrayOutput() ApiErrorBaseResponseArrayOutput {
	return o
}

func (o ApiErrorBaseResponseArrayOutput) ToApiErrorBaseResponseArrayOutputWithContext(ctx context.Context) ApiErrorBaseResponseArrayOutput {
	return o
}

func (o ApiErrorBaseResponseArrayOutput) Index(i pulumi.IntInput) ApiErrorBaseResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiErrorBaseResponse {
		return vs[0].([]ApiErrorBaseResponse)[vs[1].(int)]
	}).(ApiErrorBaseResponseOutput)
}

type ApiErrorResponse struct {
	Code       *string                `pulumi:"code"`
	Details    []ApiErrorBaseResponse `pulumi:"details"`
	Innererror *InnerErrorResponse    `pulumi:"innererror"`
	Message    *string                `pulumi:"message"`
	Target     *string                `pulumi:"target"`
}

type ApiErrorResponseOutput struct{ *pulumi.OutputState }

func (ApiErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiErrorResponse)(nil)).Elem()
}

func (o ApiErrorResponseOutput) ToApiErrorResponseOutput() ApiErrorResponseOutput {
	return o
}

func (o ApiErrorResponseOutput) ToApiErrorResponseOutputWithContext(ctx context.Context) ApiErrorResponseOutput {
	return o
}

func (o ApiErrorResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiErrorResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ApiErrorResponseOutput) Details() ApiErrorBaseResponseArrayOutput {
	return o.ApplyT(func(v ApiErrorResponse) []ApiErrorBaseResponse { return v.Details }).(ApiErrorBaseResponseArrayOutput)
}

func (o ApiErrorResponseOutput) Innererror() InnerErrorResponsePtrOutput {
	return o.ApplyT(func(v ApiErrorResponse) *InnerErrorResponse { return v.Innererror }).(InnerErrorResponsePtrOutput)
}

func (o ApiErrorResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiErrorResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ApiErrorResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiErrorResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type ApiErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiErrorResponse)(nil)).Elem()
}

func (o ApiErrorResponsePtrOutput) ToApiErrorResponsePtrOutput() ApiErrorResponsePtrOutput {
	return o
}

func (o ApiErrorResponsePtrOutput) ToApiErrorResponsePtrOutputWithContext(ctx context.Context) ApiErrorResponsePtrOutput {
	return o
}

func (o ApiErrorResponsePtrOutput) Elem() ApiErrorResponseOutput {
	return o.ApplyT(func(v *ApiErrorResponse) ApiErrorResponse {
		if v != nil {
			return *v
		}
		var ret ApiErrorResponse
		return ret
	}).(ApiErrorResponseOutput)
}

func (o ApiErrorResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ApiErrorResponsePtrOutput) Details() ApiErrorBaseResponseArrayOutput {
	return o.ApplyT(func(v *ApiErrorResponse) []ApiErrorBaseResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(ApiErrorBaseResponseArrayOutput)
}

func (o ApiErrorResponsePtrOutput) Innererror() InnerErrorResponsePtrOutput {
	return o.ApplyT(func(v *ApiErrorResponse) *InnerErrorResponse {
		if v == nil {
			return nil
		}
		return v.Innererror
	}).(InnerErrorResponsePtrOutput)
}

func (o ApiErrorResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ApiErrorResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

type ApplicationProfile struct {
	GalleryApplications []VMGalleryApplication `pulumi:"galleryApplications"`
}





type ApplicationProfileInput interface {
	pulumi.Input

	ToApplicationProfileOutput() ApplicationProfileOutput
	ToApplicationProfileOutputWithContext(context.Context) ApplicationProfileOutput
}

type ApplicationProfileArgs struct {
	GalleryApplications VMGalleryApplicationArrayInput `pulumi:"galleryApplications"`
}

func (ApplicationProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationProfile)(nil)).Elem()
}

func (i ApplicationProfileArgs) ToApplicationProfileOutput() ApplicationProfileOutput {
	return i.ToApplicationProfileOutputWithContext(context.Background())
}

func (i ApplicationProfileArgs) ToApplicationProfileOutputWithContext(ctx context.Context) ApplicationProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationProfileOutput)
}

func (i ApplicationProfileArgs) ToApplicationProfilePtrOutput() ApplicationProfilePtrOutput {
	return i.ToApplicationProfilePtrOutputWithContext(context.Background())
}

func (i ApplicationProfileArgs) ToApplicationProfilePtrOutputWithContext(ctx context.Context) ApplicationProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationProfileOutput).ToApplicationProfilePtrOutputWithContext(ctx)
}









type ApplicationProfilePtrInput interface {
	pulumi.Input

	ToApplicationProfilePtrOutput() ApplicationProfilePtrOutput
	ToApplicationProfilePtrOutputWithContext(context.Context) ApplicationProfilePtrOutput
}

type applicationProfilePtrType ApplicationProfileArgs

func ApplicationProfilePtr(v *ApplicationProfileArgs) ApplicationProfilePtrInput {
	return (*applicationProfilePtrType)(v)
}

func (*applicationProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationProfile)(nil)).Elem()
}

func (i *applicationProfilePtrType) ToApplicationProfilePtrOutput() ApplicationProfilePtrOutput {
	return i.ToApplicationProfilePtrOutputWithContext(context.Background())
}

func (i *applicationProfilePtrType) ToApplicationProfilePtrOutputWithContext(ctx context.Context) ApplicationProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationProfilePtrOutput)
}

type ApplicationProfileOutput struct{ *pulumi.OutputState }

func (ApplicationProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationProfile)(nil)).Elem()
}

func (o ApplicationProfileOutput) ToApplicationProfileOutput() ApplicationProfileOutput {
	return o
}

func (o ApplicationProfileOutput) ToApplicationProfileOutputWithContext(ctx context.Context) ApplicationProfileOutput {
	return o
}

func (o ApplicationProfileOutput) ToApplicationProfilePtrOutput() ApplicationProfilePtrOutput {
	return o.ToApplicationProfilePtrOutputWithContext(context.Background())
}

func (o ApplicationProfileOutput) ToApplicationProfilePtrOutputWithContext(ctx context.Context) ApplicationProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationProfile) *ApplicationProfile {
		return &v
	}).(ApplicationProfilePtrOutput)
}

func (o ApplicationProfileOutput) GalleryApplications() VMGalleryApplicationArrayOutput {
	return o.ApplyT(func(v ApplicationProfile) []VMGalleryApplication { return v.GalleryApplications }).(VMGalleryApplicationArrayOutput)
}

type ApplicationProfilePtrOutput struct{ *pulumi.OutputState }

func (ApplicationProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationProfile)(nil)).Elem()
}

func (o ApplicationProfilePtrOutput) ToApplicationProfilePtrOutput() ApplicationProfilePtrOutput {
	return o
}

func (o ApplicationProfilePtrOutput) ToApplicationProfilePtrOutputWithContext(ctx context.Context) ApplicationProfilePtrOutput {
	return o
}

func (o ApplicationProfilePtrOutput) Elem() ApplicationProfileOutput {
	return o.ApplyT(func(v *ApplicationProfile) ApplicationProfile {
		if v != nil {
			return *v
		}
		var ret ApplicationProfile
		return ret
	}).(ApplicationProfileOutput)
}

func (o ApplicationProfilePtrOutput) GalleryApplications() VMGalleryApplicationArrayOutput {
	return o.ApplyT(func(v *ApplicationProfile) []VMGalleryApplication {
		if v == nil {
			return nil
		}
		return v.GalleryApplications
	}).(VMGalleryApplicationArrayOutput)
}

type ApplicationProfileResponse struct {
	GalleryApplications []VMGalleryApplicationResponse `pulumi:"galleryApplications"`
}

type ApplicationProfileResponseOutput struct{ *pulumi.OutputState }

func (ApplicationProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationProfileResponse)(nil)).Elem()
}

func (o ApplicationProfileResponseOutput) ToApplicationProfileResponseOutput() ApplicationProfileResponseOutput {
	return o
}

func (o ApplicationProfileResponseOutput) ToApplicationProfileResponseOutputWithContext(ctx context.Context) ApplicationProfileResponseOutput {
	return o
}

func (o ApplicationProfileResponseOutput) GalleryApplications() VMGalleryApplicationResponseArrayOutput {
	return o.ApplyT(func(v ApplicationProfileResponse) []VMGalleryApplicationResponse { return v.GalleryApplications }).(VMGalleryApplicationResponseArrayOutput)
}

type ApplicationProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationProfileResponse)(nil)).Elem()
}

func (o ApplicationProfileResponsePtrOutput) ToApplicationProfileResponsePtrOutput() ApplicationProfileResponsePtrOutput {
	return o
}

func (o ApplicationProfileResponsePtrOutput) ToApplicationProfileResponsePtrOutputWithContext(ctx context.Context) ApplicationProfileResponsePtrOutput {
	return o
}

func (o ApplicationProfileResponsePtrOutput) Elem() ApplicationProfileResponseOutput {
	return o.ApplyT(func(v *ApplicationProfileResponse) ApplicationProfileResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationProfileResponse
		return ret
	}).(ApplicationProfileResponseOutput)
}

func (o ApplicationProfileResponsePtrOutput) GalleryApplications() VMGalleryApplicationResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationProfileResponse) []VMGalleryApplicationResponse {
		if v == nil {
			return nil
		}
		return v.GalleryApplications
	}).(VMGalleryApplicationResponseArrayOutput)
}

type AutomaticOSUpgradePolicy struct {
	DisableAutomaticRollback *bool `pulumi:"disableAutomaticRollback"`
	EnableAutomaticOSUpgrade *bool `pulumi:"enableAutomaticOSUpgrade"`
	UseRollingUpgradePolicy  *bool `pulumi:"useRollingUpgradePolicy"`
}





type AutomaticOSUpgradePolicyInput interface {
	pulumi.Input

	ToAutomaticOSUpgradePolicyOutput() AutomaticOSUpgradePolicyOutput
	ToAutomaticOSUpgradePolicyOutputWithContext(context.Context) AutomaticOSUpgradePolicyOutput
}

type AutomaticOSUpgradePolicyArgs struct {
	DisableAutomaticRollback pulumi.BoolPtrInput `pulumi:"disableAutomaticRollback"`
	EnableAutomaticOSUpgrade pulumi.BoolPtrInput `pulumi:"enableAutomaticOSUpgrade"`
	UseRollingUpgradePolicy  pulumi.BoolPtrInput `pulumi:"useRollingUpgradePolicy"`
}

func (AutomaticOSUpgradePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomaticOSUpgradePolicy)(nil)).Elem()
}

func (i AutomaticOSUpgradePolicyArgs) ToAutomaticOSUpgradePolicyOutput() AutomaticOSUpgradePolicyOutput {
	return i.ToAutomaticOSUpgradePolicyOutputWithContext(context.Background())
}

func (i AutomaticOSUpgradePolicyArgs) ToAutomaticOSUpgradePolicyOutputWithContext(ctx context.Context) AutomaticOSUpgradePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomaticOSUpgradePolicyOutput)
}

func (i AutomaticOSUpgradePolicyArgs) ToAutomaticOSUpgradePolicyPtrOutput() AutomaticOSUpgradePolicyPtrOutput {
	return i.ToAutomaticOSUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i AutomaticOSUpgradePolicyArgs) ToAutomaticOSUpgradePolicyPtrOutputWithContext(ctx context.Context) AutomaticOSUpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomaticOSUpgradePolicyOutput).ToAutomaticOSUpgradePolicyPtrOutputWithContext(ctx)
}









type AutomaticOSUpgradePolicyPtrInput interface {
	pulumi.Input

	ToAutomaticOSUpgradePolicyPtrOutput() AutomaticOSUpgradePolicyPtrOutput
	ToAutomaticOSUpgradePolicyPtrOutputWithContext(context.Context) AutomaticOSUpgradePolicyPtrOutput
}

type automaticOSUpgradePolicyPtrType AutomaticOSUpgradePolicyArgs

func AutomaticOSUpgradePolicyPtr(v *AutomaticOSUpgradePolicyArgs) AutomaticOSUpgradePolicyPtrInput {
	return (*automaticOSUpgradePolicyPtrType)(v)
}

func (*automaticOSUpgradePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomaticOSUpgradePolicy)(nil)).Elem()
}

func (i *automaticOSUpgradePolicyPtrType) ToAutomaticOSUpgradePolicyPtrOutput() AutomaticOSUpgradePolicyPtrOutput {
	return i.ToAutomaticOSUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i *automaticOSUpgradePolicyPtrType) ToAutomaticOSUpgradePolicyPtrOutputWithContext(ctx context.Context) AutomaticOSUpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomaticOSUpgradePolicyPtrOutput)
}

type AutomaticOSUpgradePolicyOutput struct{ *pulumi.OutputState }

func (AutomaticOSUpgradePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomaticOSUpgradePolicy)(nil)).Elem()
}

func (o AutomaticOSUpgradePolicyOutput) ToAutomaticOSUpgradePolicyOutput() AutomaticOSUpgradePolicyOutput {
	return o
}

func (o AutomaticOSUpgradePolicyOutput) ToAutomaticOSUpgradePolicyOutputWithContext(ctx context.Context) AutomaticOSUpgradePolicyOutput {
	return o
}

func (o AutomaticOSUpgradePolicyOutput) ToAutomaticOSUpgradePolicyPtrOutput() AutomaticOSUpgradePolicyPtrOutput {
	return o.ToAutomaticOSUpgradePolicyPtrOutputWithContext(context.Background())
}

func (o AutomaticOSUpgradePolicyOutput) ToAutomaticOSUpgradePolicyPtrOutputWithContext(ctx context.Context) AutomaticOSUpgradePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutomaticOSUpgradePolicy) *AutomaticOSUpgradePolicy {
		return &v
	}).(AutomaticOSUpgradePolicyPtrOutput)
}

func (o AutomaticOSUpgradePolicyOutput) DisableAutomaticRollback() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutomaticOSUpgradePolicy) *bool { return v.DisableAutomaticRollback }).(pulumi.BoolPtrOutput)
}

func (o AutomaticOSUpgradePolicyOutput) EnableAutomaticOSUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutomaticOSUpgradePolicy) *bool { return v.EnableAutomaticOSUpgrade }).(pulumi.BoolPtrOutput)
}

func (o AutomaticOSUpgradePolicyOutput) UseRollingUpgradePolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutomaticOSUpgradePolicy) *bool { return v.UseRollingUpgradePolicy }).(pulumi.BoolPtrOutput)
}

type AutomaticOSUpgradePolicyPtrOutput struct{ *pulumi.OutputState }

func (AutomaticOSUpgradePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomaticOSUpgradePolicy)(nil)).Elem()
}

func (o AutomaticOSUpgradePolicyPtrOutput) ToAutomaticOSUpgradePolicyPtrOutput() AutomaticOSUpgradePolicyPtrOutput {
	return o
}

func (o AutomaticOSUpgradePolicyPtrOutput) ToAutomaticOSUpgradePolicyPtrOutputWithContext(ctx context.Context) AutomaticOSUpgradePolicyPtrOutput {
	return o
}

func (o AutomaticOSUpgradePolicyPtrOutput) Elem() AutomaticOSUpgradePolicyOutput {
	return o.ApplyT(func(v *AutomaticOSUpgradePolicy) AutomaticOSUpgradePolicy {
		if v != nil {
			return *v
		}
		var ret AutomaticOSUpgradePolicy
		return ret
	}).(AutomaticOSUpgradePolicyOutput)
}

func (o AutomaticOSUpgradePolicyPtrOutput) DisableAutomaticRollback() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutomaticOSUpgradePolicy) *bool {
		if v == nil {
			return nil
		}
		return v.DisableAutomaticRollback
	}).(pulumi.BoolPtrOutput)
}

func (o AutomaticOSUpgradePolicyPtrOutput) EnableAutomaticOSUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutomaticOSUpgradePolicy) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAutomaticOSUpgrade
	}).(pulumi.BoolPtrOutput)
}

func (o AutomaticOSUpgradePolicyPtrOutput) UseRollingUpgradePolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutomaticOSUpgradePolicy) *bool {
		if v == nil {
			return nil
		}
		return v.UseRollingUpgradePolicy
	}).(pulumi.BoolPtrOutput)
}

type AutomaticOSUpgradePolicyResponse struct {
	DisableAutomaticRollback *bool `pulumi:"disableAutomaticRollback"`
	EnableAutomaticOSUpgrade *bool `pulumi:"enableAutomaticOSUpgrade"`
	UseRollingUpgradePolicy  *bool `pulumi:"useRollingUpgradePolicy"`
}

type AutomaticOSUpgradePolicyResponseOutput struct{ *pulumi.OutputState }

func (AutomaticOSUpgradePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomaticOSUpgradePolicyResponse)(nil)).Elem()
}

func (o AutomaticOSUpgradePolicyResponseOutput) ToAutomaticOSUpgradePolicyResponseOutput() AutomaticOSUpgradePolicyResponseOutput {
	return o
}

func (o AutomaticOSUpgradePolicyResponseOutput) ToAutomaticOSUpgradePolicyResponseOutputWithContext(ctx context.Context) AutomaticOSUpgradePolicyResponseOutput {
	return o
}

func (o AutomaticOSUpgradePolicyResponseOutput) DisableAutomaticRollback() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutomaticOSUpgradePolicyResponse) *bool { return v.DisableAutomaticRollback }).(pulumi.BoolPtrOutput)
}

func (o AutomaticOSUpgradePolicyResponseOutput) EnableAutomaticOSUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutomaticOSUpgradePolicyResponse) *bool { return v.EnableAutomaticOSUpgrade }).(pulumi.BoolPtrOutput)
}

func (o AutomaticOSUpgradePolicyResponseOutput) UseRollingUpgradePolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutomaticOSUpgradePolicyResponse) *bool { return v.UseRollingUpgradePolicy }).(pulumi.BoolPtrOutput)
}

type AutomaticOSUpgradePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (AutomaticOSUpgradePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomaticOSUpgradePolicyResponse)(nil)).Elem()
}

func (o AutomaticOSUpgradePolicyResponsePtrOutput) ToAutomaticOSUpgradePolicyResponsePtrOutput() AutomaticOSUpgradePolicyResponsePtrOutput {
	return o
}

func (o AutomaticOSUpgradePolicyResponsePtrOutput) ToAutomaticOSUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) AutomaticOSUpgradePolicyResponsePtrOutput {
	return o
}

func (o AutomaticOSUpgradePolicyResponsePtrOutput) Elem() AutomaticOSUpgradePolicyResponseOutput {
	return o.ApplyT(func(v *AutomaticOSUpgradePolicyResponse) AutomaticOSUpgradePolicyResponse {
		if v != nil {
			return *v
		}
		var ret AutomaticOSUpgradePolicyResponse
		return ret
	}).(AutomaticOSUpgradePolicyResponseOutput)
}

func (o AutomaticOSUpgradePolicyResponsePtrOutput) DisableAutomaticRollback() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutomaticOSUpgradePolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisableAutomaticRollback
	}).(pulumi.BoolPtrOutput)
}

func (o AutomaticOSUpgradePolicyResponsePtrOutput) EnableAutomaticOSUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutomaticOSUpgradePolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAutomaticOSUpgrade
	}).(pulumi.BoolPtrOutput)
}

func (o AutomaticOSUpgradePolicyResponsePtrOutput) UseRollingUpgradePolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutomaticOSUpgradePolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.UseRollingUpgradePolicy
	}).(pulumi.BoolPtrOutput)
}

type AutomaticRepairsPolicy struct {
	Enabled      *bool   `pulumi:"enabled"`
	GracePeriod  *string `pulumi:"gracePeriod"`
	RepairAction *string `pulumi:"repairAction"`
}





type AutomaticRepairsPolicyInput interface {
	pulumi.Input

	ToAutomaticRepairsPolicyOutput() AutomaticRepairsPolicyOutput
	ToAutomaticRepairsPolicyOutputWithContext(context.Context) AutomaticRepairsPolicyOutput
}

type AutomaticRepairsPolicyArgs struct {
	Enabled      pulumi.BoolPtrInput   `pulumi:"enabled"`
	GracePeriod  pulumi.StringPtrInput `pulumi:"gracePeriod"`
	RepairAction pulumi.StringPtrInput `pulumi:"repairAction"`
}

func (AutomaticRepairsPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomaticRepairsPolicy)(nil)).Elem()
}

func (i AutomaticRepairsPolicyArgs) ToAutomaticRepairsPolicyOutput() AutomaticRepairsPolicyOutput {
	return i.ToAutomaticRepairsPolicyOutputWithContext(context.Background())
}

func (i AutomaticRepairsPolicyArgs) ToAutomaticRepairsPolicyOutputWithContext(ctx context.Context) AutomaticRepairsPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomaticRepairsPolicyOutput)
}

func (i AutomaticRepairsPolicyArgs) ToAutomaticRepairsPolicyPtrOutput() AutomaticRepairsPolicyPtrOutput {
	return i.ToAutomaticRepairsPolicyPtrOutputWithContext(context.Background())
}

func (i AutomaticRepairsPolicyArgs) ToAutomaticRepairsPolicyPtrOutputWithContext(ctx context.Context) AutomaticRepairsPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomaticRepairsPolicyOutput).ToAutomaticRepairsPolicyPtrOutputWithContext(ctx)
}









type AutomaticRepairsPolicyPtrInput interface {
	pulumi.Input

	ToAutomaticRepairsPolicyPtrOutput() AutomaticRepairsPolicyPtrOutput
	ToAutomaticRepairsPolicyPtrOutputWithContext(context.Context) AutomaticRepairsPolicyPtrOutput
}

type automaticRepairsPolicyPtrType AutomaticRepairsPolicyArgs

func AutomaticRepairsPolicyPtr(v *AutomaticRepairsPolicyArgs) AutomaticRepairsPolicyPtrInput {
	return (*automaticRepairsPolicyPtrType)(v)
}

func (*automaticRepairsPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomaticRepairsPolicy)(nil)).Elem()
}

func (i *automaticRepairsPolicyPtrType) ToAutomaticRepairsPolicyPtrOutput() AutomaticRepairsPolicyPtrOutput {
	return i.ToAutomaticRepairsPolicyPtrOutputWithContext(context.Background())
}

func (i *automaticRepairsPolicyPtrType) ToAutomaticRepairsPolicyPtrOutputWithContext(ctx context.Context) AutomaticRepairsPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomaticRepairsPolicyPtrOutput)
}

type AutomaticRepairsPolicyOutput struct{ *pulumi.OutputState }

func (AutomaticRepairsPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomaticRepairsPolicy)(nil)).Elem()
}

func (o AutomaticRepairsPolicyOutput) ToAutomaticRepairsPolicyOutput() AutomaticRepairsPolicyOutput {
	return o
}

func (o AutomaticRepairsPolicyOutput) ToAutomaticRepairsPolicyOutputWithContext(ctx context.Context) AutomaticRepairsPolicyOutput {
	return o
}

func (o AutomaticRepairsPolicyOutput) ToAutomaticRepairsPolicyPtrOutput() AutomaticRepairsPolicyPtrOutput {
	return o.ToAutomaticRepairsPolicyPtrOutputWithContext(context.Background())
}

func (o AutomaticRepairsPolicyOutput) ToAutomaticRepairsPolicyPtrOutputWithContext(ctx context.Context) AutomaticRepairsPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutomaticRepairsPolicy) *AutomaticRepairsPolicy {
		return &v
	}).(AutomaticRepairsPolicyPtrOutput)
}

func (o AutomaticRepairsPolicyOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutomaticRepairsPolicy) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AutomaticRepairsPolicyOutput) GracePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomaticRepairsPolicy) *string { return v.GracePeriod }).(pulumi.StringPtrOutput)
}

func (o AutomaticRepairsPolicyOutput) RepairAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomaticRepairsPolicy) *string { return v.RepairAction }).(pulumi.StringPtrOutput)
}

type AutomaticRepairsPolicyPtrOutput struct{ *pulumi.OutputState }

func (AutomaticRepairsPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomaticRepairsPolicy)(nil)).Elem()
}

func (o AutomaticRepairsPolicyPtrOutput) ToAutomaticRepairsPolicyPtrOutput() AutomaticRepairsPolicyPtrOutput {
	return o
}

func (o AutomaticRepairsPolicyPtrOutput) ToAutomaticRepairsPolicyPtrOutputWithContext(ctx context.Context) AutomaticRepairsPolicyPtrOutput {
	return o
}

func (o AutomaticRepairsPolicyPtrOutput) Elem() AutomaticRepairsPolicyOutput {
	return o.ApplyT(func(v *AutomaticRepairsPolicy) AutomaticRepairsPolicy {
		if v != nil {
			return *v
		}
		var ret AutomaticRepairsPolicy
		return ret
	}).(AutomaticRepairsPolicyOutput)
}

func (o AutomaticRepairsPolicyPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutomaticRepairsPolicy) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AutomaticRepairsPolicyPtrOutput) GracePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomaticRepairsPolicy) *string {
		if v == nil {
			return nil
		}
		return v.GracePeriod
	}).(pulumi.StringPtrOutput)
}

func (o AutomaticRepairsPolicyPtrOutput) RepairAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomaticRepairsPolicy) *string {
		if v == nil {
			return nil
		}
		return v.RepairAction
	}).(pulumi.StringPtrOutput)
}

type AutomaticRepairsPolicyResponse struct {
	Enabled      *bool   `pulumi:"enabled"`
	GracePeriod  *string `pulumi:"gracePeriod"`
	RepairAction *string `pulumi:"repairAction"`
}

type AutomaticRepairsPolicyResponseOutput struct{ *pulumi.OutputState }

func (AutomaticRepairsPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomaticRepairsPolicyResponse)(nil)).Elem()
}

func (o AutomaticRepairsPolicyResponseOutput) ToAutomaticRepairsPolicyResponseOutput() AutomaticRepairsPolicyResponseOutput {
	return o
}

func (o AutomaticRepairsPolicyResponseOutput) ToAutomaticRepairsPolicyResponseOutputWithContext(ctx context.Context) AutomaticRepairsPolicyResponseOutput {
	return o
}

func (o AutomaticRepairsPolicyResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutomaticRepairsPolicyResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AutomaticRepairsPolicyResponseOutput) GracePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomaticRepairsPolicyResponse) *string { return v.GracePeriod }).(pulumi.StringPtrOutput)
}

func (o AutomaticRepairsPolicyResponseOutput) RepairAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomaticRepairsPolicyResponse) *string { return v.RepairAction }).(pulumi.StringPtrOutput)
}

type AutomaticRepairsPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (AutomaticRepairsPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomaticRepairsPolicyResponse)(nil)).Elem()
}

func (o AutomaticRepairsPolicyResponsePtrOutput) ToAutomaticRepairsPolicyResponsePtrOutput() AutomaticRepairsPolicyResponsePtrOutput {
	return o
}

func (o AutomaticRepairsPolicyResponsePtrOutput) ToAutomaticRepairsPolicyResponsePtrOutputWithContext(ctx context.Context) AutomaticRepairsPolicyResponsePtrOutput {
	return o
}

func (o AutomaticRepairsPolicyResponsePtrOutput) Elem() AutomaticRepairsPolicyResponseOutput {
	return o.ApplyT(func(v *AutomaticRepairsPolicyResponse) AutomaticRepairsPolicyResponse {
		if v != nil {
			return *v
		}
		var ret AutomaticRepairsPolicyResponse
		return ret
	}).(AutomaticRepairsPolicyResponseOutput)
}

func (o AutomaticRepairsPolicyResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutomaticRepairsPolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AutomaticRepairsPolicyResponsePtrOutput) GracePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomaticRepairsPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.GracePeriod
	}).(pulumi.StringPtrOutput)
}

func (o AutomaticRepairsPolicyResponsePtrOutput) RepairAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomaticRepairsPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.RepairAction
	}).(pulumi.StringPtrOutput)
}

type AvailablePatchSummaryResponse struct {
	AssessmentActivityId          string           `pulumi:"assessmentActivityId"`
	CriticalAndSecurityPatchCount int              `pulumi:"criticalAndSecurityPatchCount"`
	Error                         ApiErrorResponse `pulumi:"error"`
	LastModifiedTime              string           `pulumi:"lastModifiedTime"`
	OtherPatchCount               int              `pulumi:"otherPatchCount"`
	RebootPending                 bool             `pulumi:"rebootPending"`
	StartTime                     string           `pulumi:"startTime"`
	Status                        string           `pulumi:"status"`
}

type AvailablePatchSummaryResponseOutput struct{ *pulumi.OutputState }

func (AvailablePatchSummaryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailablePatchSummaryResponse)(nil)).Elem()
}

func (o AvailablePatchSummaryResponseOutput) ToAvailablePatchSummaryResponseOutput() AvailablePatchSummaryResponseOutput {
	return o
}

func (o AvailablePatchSummaryResponseOutput) ToAvailablePatchSummaryResponseOutputWithContext(ctx context.Context) AvailablePatchSummaryResponseOutput {
	return o
}

func (o AvailablePatchSummaryResponseOutput) AssessmentActivityId() pulumi.StringOutput {
	return o.ApplyT(func(v AvailablePatchSummaryResponse) string { return v.AssessmentActivityId }).(pulumi.StringOutput)
}

func (o AvailablePatchSummaryResponseOutput) CriticalAndSecurityPatchCount() pulumi.IntOutput {
	return o.ApplyT(func(v AvailablePatchSummaryResponse) int { return v.CriticalAndSecurityPatchCount }).(pulumi.IntOutput)
}

func (o AvailablePatchSummaryResponseOutput) Error() ApiErrorResponseOutput {
	return o.ApplyT(func(v AvailablePatchSummaryResponse) ApiErrorResponse { return v.Error }).(ApiErrorResponseOutput)
}

func (o AvailablePatchSummaryResponseOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v AvailablePatchSummaryResponse) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o AvailablePatchSummaryResponseOutput) OtherPatchCount() pulumi.IntOutput {
	return o.ApplyT(func(v AvailablePatchSummaryResponse) int { return v.OtherPatchCount }).(pulumi.IntOutput)
}

func (o AvailablePatchSummaryResponseOutput) RebootPending() pulumi.BoolOutput {
	return o.ApplyT(func(v AvailablePatchSummaryResponse) bool { return v.RebootPending }).(pulumi.BoolOutput)
}

func (o AvailablePatchSummaryResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v AvailablePatchSummaryResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o AvailablePatchSummaryResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v AvailablePatchSummaryResponse) string { return v.Status }).(pulumi.StringOutput)
}

type AvailablePatchSummaryResponsePtrOutput struct{ *pulumi.OutputState }

func (AvailablePatchSummaryResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailablePatchSummaryResponse)(nil)).Elem()
}

func (o AvailablePatchSummaryResponsePtrOutput) ToAvailablePatchSummaryResponsePtrOutput() AvailablePatchSummaryResponsePtrOutput {
	return o
}

func (o AvailablePatchSummaryResponsePtrOutput) ToAvailablePatchSummaryResponsePtrOutputWithContext(ctx context.Context) AvailablePatchSummaryResponsePtrOutput {
	return o
}

func (o AvailablePatchSummaryResponsePtrOutput) Elem() AvailablePatchSummaryResponseOutput {
	return o.ApplyT(func(v *AvailablePatchSummaryResponse) AvailablePatchSummaryResponse {
		if v != nil {
			return *v
		}
		var ret AvailablePatchSummaryResponse
		return ret
	}).(AvailablePatchSummaryResponseOutput)
}

func (o AvailablePatchSummaryResponsePtrOutput) AssessmentActivityId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvailablePatchSummaryResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AssessmentActivityId
	}).(pulumi.StringPtrOutput)
}

func (o AvailablePatchSummaryResponsePtrOutput) CriticalAndSecurityPatchCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AvailablePatchSummaryResponse) *int {
		if v == nil {
			return nil
		}
		return &v.CriticalAndSecurityPatchCount
	}).(pulumi.IntPtrOutput)
}

func (o AvailablePatchSummaryResponsePtrOutput) Error() ApiErrorResponsePtrOutput {
	return o.ApplyT(func(v *AvailablePatchSummaryResponse) *ApiErrorResponse {
		if v == nil {
			return nil
		}
		return &v.Error
	}).(ApiErrorResponsePtrOutput)
}

func (o AvailablePatchSummaryResponsePtrOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvailablePatchSummaryResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastModifiedTime
	}).(pulumi.StringPtrOutput)
}

func (o AvailablePatchSummaryResponsePtrOutput) OtherPatchCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AvailablePatchSummaryResponse) *int {
		if v == nil {
			return nil
		}
		return &v.OtherPatchCount
	}).(pulumi.IntPtrOutput)
}

func (o AvailablePatchSummaryResponsePtrOutput) RebootPending() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AvailablePatchSummaryResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.RebootPending
	}).(pulumi.BoolPtrOutput)
}

func (o AvailablePatchSummaryResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvailablePatchSummaryResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o AvailablePatchSummaryResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvailablePatchSummaryResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type BillingProfile struct {
	MaxPrice *float64 `pulumi:"maxPrice"`
}





type BillingProfileInput interface {
	pulumi.Input

	ToBillingProfileOutput() BillingProfileOutput
	ToBillingProfileOutputWithContext(context.Context) BillingProfileOutput
}

type BillingProfileArgs struct {
	MaxPrice pulumi.Float64PtrInput `pulumi:"maxPrice"`
}

func (BillingProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BillingProfile)(nil)).Elem()
}

func (i BillingProfileArgs) ToBillingProfileOutput() BillingProfileOutput {
	return i.ToBillingProfileOutputWithContext(context.Background())
}

func (i BillingProfileArgs) ToBillingProfileOutputWithContext(ctx context.Context) BillingProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingProfileOutput)
}

func (i BillingProfileArgs) ToBillingProfilePtrOutput() BillingProfilePtrOutput {
	return i.ToBillingProfilePtrOutputWithContext(context.Background())
}

func (i BillingProfileArgs) ToBillingProfilePtrOutputWithContext(ctx context.Context) BillingProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingProfileOutput).ToBillingProfilePtrOutputWithContext(ctx)
}









type BillingProfilePtrInput interface {
	pulumi.Input

	ToBillingProfilePtrOutput() BillingProfilePtrOutput
	ToBillingProfilePtrOutputWithContext(context.Context) BillingProfilePtrOutput
}

type billingProfilePtrType BillingProfileArgs

func BillingProfilePtr(v *BillingProfileArgs) BillingProfilePtrInput {
	return (*billingProfilePtrType)(v)
}

func (*billingProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingProfile)(nil)).Elem()
}

func (i *billingProfilePtrType) ToBillingProfilePtrOutput() BillingProfilePtrOutput {
	return i.ToBillingProfilePtrOutputWithContext(context.Background())
}

func (i *billingProfilePtrType) ToBillingProfilePtrOutputWithContext(ctx context.Context) BillingProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingProfilePtrOutput)
}

type BillingProfileOutput struct{ *pulumi.OutputState }

func (BillingProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BillingProfile)(nil)).Elem()
}

func (o BillingProfileOutput) ToBillingProfileOutput() BillingProfileOutput {
	return o
}

func (o BillingProfileOutput) ToBillingProfileOutputWithContext(ctx context.Context) BillingProfileOutput {
	return o
}

func (o BillingProfileOutput) ToBillingProfilePtrOutput() BillingProfilePtrOutput {
	return o.ToBillingProfilePtrOutputWithContext(context.Background())
}

func (o BillingProfileOutput) ToBillingProfilePtrOutputWithContext(ctx context.Context) BillingProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BillingProfile) *BillingProfile {
		return &v
	}).(BillingProfilePtrOutput)
}

func (o BillingProfileOutput) MaxPrice() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v BillingProfile) *float64 { return v.MaxPrice }).(pulumi.Float64PtrOutput)
}

type BillingProfilePtrOutput struct{ *pulumi.OutputState }

func (BillingProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingProfile)(nil)).Elem()
}

func (o BillingProfilePtrOutput) ToBillingProfilePtrOutput() BillingProfilePtrOutput {
	return o
}

func (o BillingProfilePtrOutput) ToBillingProfilePtrOutputWithContext(ctx context.Context) BillingProfilePtrOutput {
	return o
}

func (o BillingProfilePtrOutput) Elem() BillingProfileOutput {
	return o.ApplyT(func(v *BillingProfile) BillingProfile {
		if v != nil {
			return *v
		}
		var ret BillingProfile
		return ret
	}).(BillingProfileOutput)
}

func (o BillingProfilePtrOutput) MaxPrice() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *BillingProfile) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxPrice
	}).(pulumi.Float64PtrOutput)
}

type BillingProfileResponse struct {
	MaxPrice *float64 `pulumi:"maxPrice"`
}

type BillingProfileResponseOutput struct{ *pulumi.OutputState }

func (BillingProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BillingProfileResponse)(nil)).Elem()
}

func (o BillingProfileResponseOutput) ToBillingProfileResponseOutput() BillingProfileResponseOutput {
	return o
}

func (o BillingProfileResponseOutput) ToBillingProfileResponseOutputWithContext(ctx context.Context) BillingProfileResponseOutput {
	return o
}

func (o BillingProfileResponseOutput) MaxPrice() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v BillingProfileResponse) *float64 { return v.MaxPrice }).(pulumi.Float64PtrOutput)
}

type BillingProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (BillingProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingProfileResponse)(nil)).Elem()
}

func (o BillingProfileResponsePtrOutput) ToBillingProfileResponsePtrOutput() BillingProfileResponsePtrOutput {
	return o
}

func (o BillingProfileResponsePtrOutput) ToBillingProfileResponsePtrOutputWithContext(ctx context.Context) BillingProfileResponsePtrOutput {
	return o
}

func (o BillingProfileResponsePtrOutput) Elem() BillingProfileResponseOutput {
	return o.ApplyT(func(v *BillingProfileResponse) BillingProfileResponse {
		if v != nil {
			return *v
		}
		var ret BillingProfileResponse
		return ret
	}).(BillingProfileResponseOutput)
}

func (o BillingProfileResponsePtrOutput) MaxPrice() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *BillingProfileResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxPrice
	}).(pulumi.Float64PtrOutput)
}

type BootDiagnostics struct {
	Enabled    *bool   `pulumi:"enabled"`
	StorageUri *string `pulumi:"storageUri"`
}





type BootDiagnosticsInput interface {
	pulumi.Input

	ToBootDiagnosticsOutput() BootDiagnosticsOutput
	ToBootDiagnosticsOutputWithContext(context.Context) BootDiagnosticsOutput
}

type BootDiagnosticsArgs struct {
	Enabled    pulumi.BoolPtrInput   `pulumi:"enabled"`
	StorageUri pulumi.StringPtrInput `pulumi:"storageUri"`
}

func (BootDiagnosticsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BootDiagnostics)(nil)).Elem()
}

func (i BootDiagnosticsArgs) ToBootDiagnosticsOutput() BootDiagnosticsOutput {
	return i.ToBootDiagnosticsOutputWithContext(context.Background())
}

func (i BootDiagnosticsArgs) ToBootDiagnosticsOutputWithContext(ctx context.Context) BootDiagnosticsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootDiagnosticsOutput)
}

func (i BootDiagnosticsArgs) ToBootDiagnosticsPtrOutput() BootDiagnosticsPtrOutput {
	return i.ToBootDiagnosticsPtrOutputWithContext(context.Background())
}

func (i BootDiagnosticsArgs) ToBootDiagnosticsPtrOutputWithContext(ctx context.Context) BootDiagnosticsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootDiagnosticsOutput).ToBootDiagnosticsPtrOutputWithContext(ctx)
}









type BootDiagnosticsPtrInput interface {
	pulumi.Input

	ToBootDiagnosticsPtrOutput() BootDiagnosticsPtrOutput
	ToBootDiagnosticsPtrOutputWithContext(context.Context) BootDiagnosticsPtrOutput
}

type bootDiagnosticsPtrType BootDiagnosticsArgs

func BootDiagnosticsPtr(v *BootDiagnosticsArgs) BootDiagnosticsPtrInput {
	return (*bootDiagnosticsPtrType)(v)
}

func (*bootDiagnosticsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BootDiagnostics)(nil)).Elem()
}

func (i *bootDiagnosticsPtrType) ToBootDiagnosticsPtrOutput() BootDiagnosticsPtrOutput {
	return i.ToBootDiagnosticsPtrOutputWithContext(context.Background())
}

func (i *bootDiagnosticsPtrType) ToBootDiagnosticsPtrOutputWithContext(ctx context.Context) BootDiagnosticsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootDiagnosticsPtrOutput)
}

type BootDiagnosticsOutput struct{ *pulumi.OutputState }

func (BootDiagnosticsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BootDiagnostics)(nil)).Elem()
}

func (o BootDiagnosticsOutput) ToBootDiagnosticsOutput() BootDiagnosticsOutput {
	return o
}

func (o BootDiagnosticsOutput) ToBootDiagnosticsOutputWithContext(ctx context.Context) BootDiagnosticsOutput {
	return o
}

func (o BootDiagnosticsOutput) ToBootDiagnosticsPtrOutput() BootDiagnosticsPtrOutput {
	return o.ToBootDiagnosticsPtrOutputWithContext(context.Background())
}

func (o BootDiagnosticsOutput) ToBootDiagnosticsPtrOutputWithContext(ctx context.Context) BootDiagnosticsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BootDiagnostics) *BootDiagnostics {
		return &v
	}).(BootDiagnosticsPtrOutput)
}

func (o BootDiagnosticsOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BootDiagnostics) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o BootDiagnosticsOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BootDiagnostics) *string { return v.StorageUri }).(pulumi.StringPtrOutput)
}

type BootDiagnosticsPtrOutput struct{ *pulumi.OutputState }

func (BootDiagnosticsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BootDiagnostics)(nil)).Elem()
}

func (o BootDiagnosticsPtrOutput) ToBootDiagnosticsPtrOutput() BootDiagnosticsPtrOutput {
	return o
}

func (o BootDiagnosticsPtrOutput) ToBootDiagnosticsPtrOutputWithContext(ctx context.Context) BootDiagnosticsPtrOutput {
	return o
}

func (o BootDiagnosticsPtrOutput) Elem() BootDiagnosticsOutput {
	return o.ApplyT(func(v *BootDiagnostics) BootDiagnostics {
		if v != nil {
			return *v
		}
		var ret BootDiagnostics
		return ret
	}).(BootDiagnosticsOutput)
}

func (o BootDiagnosticsPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BootDiagnostics) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o BootDiagnosticsPtrOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BootDiagnostics) *string {
		if v == nil {
			return nil
		}
		return v.StorageUri
	}).(pulumi.StringPtrOutput)
}

type BootDiagnosticsInstanceViewResponse struct {
	ConsoleScreenshotBlobUri string                     `pulumi:"consoleScreenshotBlobUri"`
	SerialConsoleLogBlobUri  string                     `pulumi:"serialConsoleLogBlobUri"`
	Status                   InstanceViewStatusResponse `pulumi:"status"`
}

type BootDiagnosticsInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (BootDiagnosticsInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BootDiagnosticsInstanceViewResponse)(nil)).Elem()
}

func (o BootDiagnosticsInstanceViewResponseOutput) ToBootDiagnosticsInstanceViewResponseOutput() BootDiagnosticsInstanceViewResponseOutput {
	return o
}

func (o BootDiagnosticsInstanceViewResponseOutput) ToBootDiagnosticsInstanceViewResponseOutputWithContext(ctx context.Context) BootDiagnosticsInstanceViewResponseOutput {
	return o
}

func (o BootDiagnosticsInstanceViewResponseOutput) ConsoleScreenshotBlobUri() pulumi.StringOutput {
	return o.ApplyT(func(v BootDiagnosticsInstanceViewResponse) string { return v.ConsoleScreenshotBlobUri }).(pulumi.StringOutput)
}

func (o BootDiagnosticsInstanceViewResponseOutput) SerialConsoleLogBlobUri() pulumi.StringOutput {
	return o.ApplyT(func(v BootDiagnosticsInstanceViewResponse) string { return v.SerialConsoleLogBlobUri }).(pulumi.StringOutput)
}

func (o BootDiagnosticsInstanceViewResponseOutput) Status() InstanceViewStatusResponseOutput {
	return o.ApplyT(func(v BootDiagnosticsInstanceViewResponse) InstanceViewStatusResponse { return v.Status }).(InstanceViewStatusResponseOutput)
}

type BootDiagnosticsInstanceViewResponsePtrOutput struct{ *pulumi.OutputState }

func (BootDiagnosticsInstanceViewResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BootDiagnosticsInstanceViewResponse)(nil)).Elem()
}

func (o BootDiagnosticsInstanceViewResponsePtrOutput) ToBootDiagnosticsInstanceViewResponsePtrOutput() BootDiagnosticsInstanceViewResponsePtrOutput {
	return o
}

func (o BootDiagnosticsInstanceViewResponsePtrOutput) ToBootDiagnosticsInstanceViewResponsePtrOutputWithContext(ctx context.Context) BootDiagnosticsInstanceViewResponsePtrOutput {
	return o
}

func (o BootDiagnosticsInstanceViewResponsePtrOutput) Elem() BootDiagnosticsInstanceViewResponseOutput {
	return o.ApplyT(func(v *BootDiagnosticsInstanceViewResponse) BootDiagnosticsInstanceViewResponse {
		if v != nil {
			return *v
		}
		var ret BootDiagnosticsInstanceViewResponse
		return ret
	}).(BootDiagnosticsInstanceViewResponseOutput)
}

func (o BootDiagnosticsInstanceViewResponsePtrOutput) ConsoleScreenshotBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BootDiagnosticsInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ConsoleScreenshotBlobUri
	}).(pulumi.StringPtrOutput)
}

func (o BootDiagnosticsInstanceViewResponsePtrOutput) SerialConsoleLogBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BootDiagnosticsInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SerialConsoleLogBlobUri
	}).(pulumi.StringPtrOutput)
}

func (o BootDiagnosticsInstanceViewResponsePtrOutput) Status() InstanceViewStatusResponsePtrOutput {
	return o.ApplyT(func(v *BootDiagnosticsInstanceViewResponse) *InstanceViewStatusResponse {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(InstanceViewStatusResponsePtrOutput)
}

type BootDiagnosticsResponse struct {
	Enabled    *bool   `pulumi:"enabled"`
	StorageUri *string `pulumi:"storageUri"`
}

type BootDiagnosticsResponseOutput struct{ *pulumi.OutputState }

func (BootDiagnosticsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BootDiagnosticsResponse)(nil)).Elem()
}

func (o BootDiagnosticsResponseOutput) ToBootDiagnosticsResponseOutput() BootDiagnosticsResponseOutput {
	return o
}

func (o BootDiagnosticsResponseOutput) ToBootDiagnosticsResponseOutputWithContext(ctx context.Context) BootDiagnosticsResponseOutput {
	return o
}

func (o BootDiagnosticsResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BootDiagnosticsResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o BootDiagnosticsResponseOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BootDiagnosticsResponse) *string { return v.StorageUri }).(pulumi.StringPtrOutput)
}

type BootDiagnosticsResponsePtrOutput struct{ *pulumi.OutputState }

func (BootDiagnosticsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BootDiagnosticsResponse)(nil)).Elem()
}

func (o BootDiagnosticsResponsePtrOutput) ToBootDiagnosticsResponsePtrOutput() BootDiagnosticsResponsePtrOutput {
	return o
}

func (o BootDiagnosticsResponsePtrOutput) ToBootDiagnosticsResponsePtrOutputWithContext(ctx context.Context) BootDiagnosticsResponsePtrOutput {
	return o
}

func (o BootDiagnosticsResponsePtrOutput) Elem() BootDiagnosticsResponseOutput {
	return o.ApplyT(func(v *BootDiagnosticsResponse) BootDiagnosticsResponse {
		if v != nil {
			return *v
		}
		var ret BootDiagnosticsResponse
		return ret
	}).(BootDiagnosticsResponseOutput)
}

func (o BootDiagnosticsResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BootDiagnosticsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o BootDiagnosticsResponsePtrOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BootDiagnosticsResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageUri
	}).(pulumi.StringPtrOutput)
}

type CapacityReservationGroupInstanceViewResponse struct {
	CapacityReservations []CapacityReservationInstanceViewWithNameResponse `pulumi:"capacityReservations"`
}

type CapacityReservationGroupInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (CapacityReservationGroupInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityReservationGroupInstanceViewResponse)(nil)).Elem()
}

func (o CapacityReservationGroupInstanceViewResponseOutput) ToCapacityReservationGroupInstanceViewResponseOutput() CapacityReservationGroupInstanceViewResponseOutput {
	return o
}

func (o CapacityReservationGroupInstanceViewResponseOutput) ToCapacityReservationGroupInstanceViewResponseOutputWithContext(ctx context.Context) CapacityReservationGroupInstanceViewResponseOutput {
	return o
}

func (o CapacityReservationGroupInstanceViewResponseOutput) CapacityReservations() CapacityReservationInstanceViewWithNameResponseArrayOutput {
	return o.ApplyT(func(v CapacityReservationGroupInstanceViewResponse) []CapacityReservationInstanceViewWithNameResponse {
		return v.CapacityReservations
	}).(CapacityReservationInstanceViewWithNameResponseArrayOutput)
}

type CapacityReservationInstanceViewResponse struct {
	Statuses        []InstanceViewStatusResponse            `pulumi:"statuses"`
	UtilizationInfo *CapacityReservationUtilizationResponse `pulumi:"utilizationInfo"`
}

type CapacityReservationInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (CapacityReservationInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityReservationInstanceViewResponse)(nil)).Elem()
}

func (o CapacityReservationInstanceViewResponseOutput) ToCapacityReservationInstanceViewResponseOutput() CapacityReservationInstanceViewResponseOutput {
	return o
}

func (o CapacityReservationInstanceViewResponseOutput) ToCapacityReservationInstanceViewResponseOutputWithContext(ctx context.Context) CapacityReservationInstanceViewResponseOutput {
	return o
}

func (o CapacityReservationInstanceViewResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v CapacityReservationInstanceViewResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

func (o CapacityReservationInstanceViewResponseOutput) UtilizationInfo() CapacityReservationUtilizationResponsePtrOutput {
	return o.ApplyT(func(v CapacityReservationInstanceViewResponse) *CapacityReservationUtilizationResponse {
		return v.UtilizationInfo
	}).(CapacityReservationUtilizationResponsePtrOutput)
}

type CapacityReservationInstanceViewWithNameResponse struct {
	Name            string                                  `pulumi:"name"`
	Statuses        []InstanceViewStatusResponse            `pulumi:"statuses"`
	UtilizationInfo *CapacityReservationUtilizationResponse `pulumi:"utilizationInfo"`
}

type CapacityReservationInstanceViewWithNameResponseOutput struct{ *pulumi.OutputState }

func (CapacityReservationInstanceViewWithNameResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityReservationInstanceViewWithNameResponse)(nil)).Elem()
}

func (o CapacityReservationInstanceViewWithNameResponseOutput) ToCapacityReservationInstanceViewWithNameResponseOutput() CapacityReservationInstanceViewWithNameResponseOutput {
	return o
}

func (o CapacityReservationInstanceViewWithNameResponseOutput) ToCapacityReservationInstanceViewWithNameResponseOutputWithContext(ctx context.Context) CapacityReservationInstanceViewWithNameResponseOutput {
	return o
}

func (o CapacityReservationInstanceViewWithNameResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CapacityReservationInstanceViewWithNameResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CapacityReservationInstanceViewWithNameResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v CapacityReservationInstanceViewWithNameResponse) []InstanceViewStatusResponse {
		return v.Statuses
	}).(InstanceViewStatusResponseArrayOutput)
}

func (o CapacityReservationInstanceViewWithNameResponseOutput) UtilizationInfo() CapacityReservationUtilizationResponsePtrOutput {
	return o.ApplyT(func(v CapacityReservationInstanceViewWithNameResponse) *CapacityReservationUtilizationResponse {
		return v.UtilizationInfo
	}).(CapacityReservationUtilizationResponsePtrOutput)
}

type CapacityReservationInstanceViewWithNameResponseArrayOutput struct{ *pulumi.OutputState }

func (CapacityReservationInstanceViewWithNameResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CapacityReservationInstanceViewWithNameResponse)(nil)).Elem()
}

func (o CapacityReservationInstanceViewWithNameResponseArrayOutput) ToCapacityReservationInstanceViewWithNameResponseArrayOutput() CapacityReservationInstanceViewWithNameResponseArrayOutput {
	return o
}

func (o CapacityReservationInstanceViewWithNameResponseArrayOutput) ToCapacityReservationInstanceViewWithNameResponseArrayOutputWithContext(ctx context.Context) CapacityReservationInstanceViewWithNameResponseArrayOutput {
	return o
}

func (o CapacityReservationInstanceViewWithNameResponseArrayOutput) Index(i pulumi.IntInput) CapacityReservationInstanceViewWithNameResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CapacityReservationInstanceViewWithNameResponse {
		return vs[0].([]CapacityReservationInstanceViewWithNameResponse)[vs[1].(int)]
	}).(CapacityReservationInstanceViewWithNameResponseOutput)
}

type CapacityReservationProfile struct {
	CapacityReservationGroup *SubResource `pulumi:"capacityReservationGroup"`
}





type CapacityReservationProfileInput interface {
	pulumi.Input

	ToCapacityReservationProfileOutput() CapacityReservationProfileOutput
	ToCapacityReservationProfileOutputWithContext(context.Context) CapacityReservationProfileOutput
}

type CapacityReservationProfileArgs struct {
	CapacityReservationGroup SubResourcePtrInput `pulumi:"capacityReservationGroup"`
}

func (CapacityReservationProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityReservationProfile)(nil)).Elem()
}

func (i CapacityReservationProfileArgs) ToCapacityReservationProfileOutput() CapacityReservationProfileOutput {
	return i.ToCapacityReservationProfileOutputWithContext(context.Background())
}

func (i CapacityReservationProfileArgs) ToCapacityReservationProfileOutputWithContext(ctx context.Context) CapacityReservationProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityReservationProfileOutput)
}

func (i CapacityReservationProfileArgs) ToCapacityReservationProfilePtrOutput() CapacityReservationProfilePtrOutput {
	return i.ToCapacityReservationProfilePtrOutputWithContext(context.Background())
}

func (i CapacityReservationProfileArgs) ToCapacityReservationProfilePtrOutputWithContext(ctx context.Context) CapacityReservationProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityReservationProfileOutput).ToCapacityReservationProfilePtrOutputWithContext(ctx)
}









type CapacityReservationProfilePtrInput interface {
	pulumi.Input

	ToCapacityReservationProfilePtrOutput() CapacityReservationProfilePtrOutput
	ToCapacityReservationProfilePtrOutputWithContext(context.Context) CapacityReservationProfilePtrOutput
}

type capacityReservationProfilePtrType CapacityReservationProfileArgs

func CapacityReservationProfilePtr(v *CapacityReservationProfileArgs) CapacityReservationProfilePtrInput {
	return (*capacityReservationProfilePtrType)(v)
}

func (*capacityReservationProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservationProfile)(nil)).Elem()
}

func (i *capacityReservationProfilePtrType) ToCapacityReservationProfilePtrOutput() CapacityReservationProfilePtrOutput {
	return i.ToCapacityReservationProfilePtrOutputWithContext(context.Background())
}

func (i *capacityReservationProfilePtrType) ToCapacityReservationProfilePtrOutputWithContext(ctx context.Context) CapacityReservationProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityReservationProfilePtrOutput)
}

type CapacityReservationProfileOutput struct{ *pulumi.OutputState }

func (CapacityReservationProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityReservationProfile)(nil)).Elem()
}

func (o CapacityReservationProfileOutput) ToCapacityReservationProfileOutput() CapacityReservationProfileOutput {
	return o
}

func (o CapacityReservationProfileOutput) ToCapacityReservationProfileOutputWithContext(ctx context.Context) CapacityReservationProfileOutput {
	return o
}

func (o CapacityReservationProfileOutput) ToCapacityReservationProfilePtrOutput() CapacityReservationProfilePtrOutput {
	return o.ToCapacityReservationProfilePtrOutputWithContext(context.Background())
}

func (o CapacityReservationProfileOutput) ToCapacityReservationProfilePtrOutputWithContext(ctx context.Context) CapacityReservationProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CapacityReservationProfile) *CapacityReservationProfile {
		return &v
	}).(CapacityReservationProfilePtrOutput)
}

func (o CapacityReservationProfileOutput) CapacityReservationGroup() SubResourcePtrOutput {
	return o.ApplyT(func(v CapacityReservationProfile) *SubResource { return v.CapacityReservationGroup }).(SubResourcePtrOutput)
}

type CapacityReservationProfilePtrOutput struct{ *pulumi.OutputState }

func (CapacityReservationProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservationProfile)(nil)).Elem()
}

func (o CapacityReservationProfilePtrOutput) ToCapacityReservationProfilePtrOutput() CapacityReservationProfilePtrOutput {
	return o
}

func (o CapacityReservationProfilePtrOutput) ToCapacityReservationProfilePtrOutputWithContext(ctx context.Context) CapacityReservationProfilePtrOutput {
	return o
}

func (o CapacityReservationProfilePtrOutput) Elem() CapacityReservationProfileOutput {
	return o.ApplyT(func(v *CapacityReservationProfile) CapacityReservationProfile {
		if v != nil {
			return *v
		}
		var ret CapacityReservationProfile
		return ret
	}).(CapacityReservationProfileOutput)
}

func (o CapacityReservationProfilePtrOutput) CapacityReservationGroup() SubResourcePtrOutput {
	return o.ApplyT(func(v *CapacityReservationProfile) *SubResource {
		if v == nil {
			return nil
		}
		return v.CapacityReservationGroup
	}).(SubResourcePtrOutput)
}

type CapacityReservationProfileResponse struct {
	CapacityReservationGroup *SubResourceResponse `pulumi:"capacityReservationGroup"`
}

type CapacityReservationProfileResponseOutput struct{ *pulumi.OutputState }

func (CapacityReservationProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityReservationProfileResponse)(nil)).Elem()
}

func (o CapacityReservationProfileResponseOutput) ToCapacityReservationProfileResponseOutput() CapacityReservationProfileResponseOutput {
	return o
}

func (o CapacityReservationProfileResponseOutput) ToCapacityReservationProfileResponseOutputWithContext(ctx context.Context) CapacityReservationProfileResponseOutput {
	return o
}

func (o CapacityReservationProfileResponseOutput) CapacityReservationGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v CapacityReservationProfileResponse) *SubResourceResponse { return v.CapacityReservationGroup }).(SubResourceResponsePtrOutput)
}

type CapacityReservationProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (CapacityReservationProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservationProfileResponse)(nil)).Elem()
}

func (o CapacityReservationProfileResponsePtrOutput) ToCapacityReservationProfileResponsePtrOutput() CapacityReservationProfileResponsePtrOutput {
	return o
}

func (o CapacityReservationProfileResponsePtrOutput) ToCapacityReservationProfileResponsePtrOutputWithContext(ctx context.Context) CapacityReservationProfileResponsePtrOutput {
	return o
}

func (o CapacityReservationProfileResponsePtrOutput) Elem() CapacityReservationProfileResponseOutput {
	return o.ApplyT(func(v *CapacityReservationProfileResponse) CapacityReservationProfileResponse {
		if v != nil {
			return *v
		}
		var ret CapacityReservationProfileResponse
		return ret
	}).(CapacityReservationProfileResponseOutput)
}

func (o CapacityReservationProfileResponsePtrOutput) CapacityReservationGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *CapacityReservationProfileResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return v.CapacityReservationGroup
	}).(SubResourceResponsePtrOutput)
}

type CapacityReservationUtilizationResponse struct {
	VirtualMachinesAllocated []SubResourceReadOnlyResponse `pulumi:"virtualMachinesAllocated"`
}

type CapacityReservationUtilizationResponseOutput struct{ *pulumi.OutputState }

func (CapacityReservationUtilizationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityReservationUtilizationResponse)(nil)).Elem()
}

func (o CapacityReservationUtilizationResponseOutput) ToCapacityReservationUtilizationResponseOutput() CapacityReservationUtilizationResponseOutput {
	return o
}

func (o CapacityReservationUtilizationResponseOutput) ToCapacityReservationUtilizationResponseOutputWithContext(ctx context.Context) CapacityReservationUtilizationResponseOutput {
	return o
}

func (o CapacityReservationUtilizationResponseOutput) VirtualMachinesAllocated() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v CapacityReservationUtilizationResponse) []SubResourceReadOnlyResponse {
		return v.VirtualMachinesAllocated
	}).(SubResourceReadOnlyResponseArrayOutput)
}

type CapacityReservationUtilizationResponsePtrOutput struct{ *pulumi.OutputState }

func (CapacityReservationUtilizationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservationUtilizationResponse)(nil)).Elem()
}

func (o CapacityReservationUtilizationResponsePtrOutput) ToCapacityReservationUtilizationResponsePtrOutput() CapacityReservationUtilizationResponsePtrOutput {
	return o
}

func (o CapacityReservationUtilizationResponsePtrOutput) ToCapacityReservationUtilizationResponsePtrOutputWithContext(ctx context.Context) CapacityReservationUtilizationResponsePtrOutput {
	return o
}

func (o CapacityReservationUtilizationResponsePtrOutput) Elem() CapacityReservationUtilizationResponseOutput {
	return o.ApplyT(func(v *CapacityReservationUtilizationResponse) CapacityReservationUtilizationResponse {
		if v != nil {
			return *v
		}
		var ret CapacityReservationUtilizationResponse
		return ret
	}).(CapacityReservationUtilizationResponseOutput)
}

func (o CapacityReservationUtilizationResponsePtrOutput) VirtualMachinesAllocated() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v *CapacityReservationUtilizationResponse) []SubResourceReadOnlyResponse {
		if v == nil {
			return nil
		}
		return v.VirtualMachinesAllocated
	}).(SubResourceReadOnlyResponseArrayOutput)
}

type DataDisk struct {
	Caching                 *CachingTypes          `pulumi:"caching"`
	CreateOption            string                 `pulumi:"createOption"`
	DeleteOption            *string                `pulumi:"deleteOption"`
	DetachOption            *string                `pulumi:"detachOption"`
	DiskSizeGB              *int                   `pulumi:"diskSizeGB"`
	Image                   *VirtualHardDisk       `pulumi:"image"`
	Lun                     int                    `pulumi:"lun"`
	ManagedDisk             *ManagedDiskParameters `pulumi:"managedDisk"`
	Name                    *string                `pulumi:"name"`
	ToBeDetached            *bool                  `pulumi:"toBeDetached"`
	Vhd                     *VirtualHardDisk       `pulumi:"vhd"`
	WriteAcceleratorEnabled *bool                  `pulumi:"writeAcceleratorEnabled"`
}





type DataDiskInput interface {
	pulumi.Input

	ToDataDiskOutput() DataDiskOutput
	ToDataDiskOutputWithContext(context.Context) DataDiskOutput
}

type DataDiskArgs struct {
	Caching                 CachingTypesPtrInput          `pulumi:"caching"`
	CreateOption            pulumi.StringInput            `pulumi:"createOption"`
	DeleteOption            pulumi.StringPtrInput         `pulumi:"deleteOption"`
	DetachOption            pulumi.StringPtrInput         `pulumi:"detachOption"`
	DiskSizeGB              pulumi.IntPtrInput            `pulumi:"diskSizeGB"`
	Image                   VirtualHardDiskPtrInput       `pulumi:"image"`
	Lun                     pulumi.IntInput               `pulumi:"lun"`
	ManagedDisk             ManagedDiskParametersPtrInput `pulumi:"managedDisk"`
	Name                    pulumi.StringPtrInput         `pulumi:"name"`
	ToBeDetached            pulumi.BoolPtrInput           `pulumi:"toBeDetached"`
	Vhd                     VirtualHardDiskPtrInput       `pulumi:"vhd"`
	WriteAcceleratorEnabled pulumi.BoolPtrInput           `pulumi:"writeAcceleratorEnabled"`
}

func (DataDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDisk)(nil)).Elem()
}

func (i DataDiskArgs) ToDataDiskOutput() DataDiskOutput {
	return i.ToDataDiskOutputWithContext(context.Background())
}

func (i DataDiskArgs) ToDataDiskOutputWithContext(ctx context.Context) DataDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskOutput)
}





type DataDiskArrayInput interface {
	pulumi.Input

	ToDataDiskArrayOutput() DataDiskArrayOutput
	ToDataDiskArrayOutputWithContext(context.Context) DataDiskArrayOutput
}

type DataDiskArray []DataDiskInput

func (DataDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDisk)(nil)).Elem()
}

func (i DataDiskArray) ToDataDiskArrayOutput() DataDiskArrayOutput {
	return i.ToDataDiskArrayOutputWithContext(context.Background())
}

func (i DataDiskArray) ToDataDiskArrayOutputWithContext(ctx context.Context) DataDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskArrayOutput)
}

type DataDiskOutput struct{ *pulumi.OutputState }

func (DataDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDisk)(nil)).Elem()
}

func (o DataDiskOutput) ToDataDiskOutput() DataDiskOutput {
	return o
}

func (o DataDiskOutput) ToDataDiskOutputWithContext(ctx context.Context) DataDiskOutput {
	return o
}

func (o DataDiskOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v DataDisk) *CachingTypes { return v.Caching }).(CachingTypesPtrOutput)
}

func (o DataDiskOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v DataDisk) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o DataDiskOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDisk) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o DataDiskOutput) DetachOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDisk) *string { return v.DetachOption }).(pulumi.StringPtrOutput)
}

func (o DataDiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataDisk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o DataDiskOutput) Image() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v DataDisk) *VirtualHardDisk { return v.Image }).(VirtualHardDiskPtrOutput)
}

func (o DataDiskOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v DataDisk) int { return v.Lun }).(pulumi.IntOutput)
}

func (o DataDiskOutput) ManagedDisk() ManagedDiskParametersPtrOutput {
	return o.ApplyT(func(v DataDisk) *ManagedDiskParameters { return v.ManagedDisk }).(ManagedDiskParametersPtrOutput)
}

func (o DataDiskOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDisk) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DataDiskOutput) ToBeDetached() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataDisk) *bool { return v.ToBeDetached }).(pulumi.BoolPtrOutput)
}

func (o DataDiskOutput) Vhd() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v DataDisk) *VirtualHardDisk { return v.Vhd }).(VirtualHardDiskPtrOutput)
}

func (o DataDiskOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataDisk) *bool { return v.WriteAcceleratorEnabled }).(pulumi.BoolPtrOutput)
}

type DataDiskArrayOutput struct{ *pulumi.OutputState }

func (DataDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDisk)(nil)).Elem()
}

func (o DataDiskArrayOutput) ToDataDiskArrayOutput() DataDiskArrayOutput {
	return o
}

func (o DataDiskArrayOutput) ToDataDiskArrayOutputWithContext(ctx context.Context) DataDiskArrayOutput {
	return o
}

func (o DataDiskArrayOutput) Index(i pulumi.IntInput) DataDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDisk {
		return vs[0].([]DataDisk)[vs[1].(int)]
	}).(DataDiskOutput)
}

type DataDiskResponse struct {
	Caching                 *string                        `pulumi:"caching"`
	CreateOption            string                         `pulumi:"createOption"`
	DeleteOption            *string                        `pulumi:"deleteOption"`
	DetachOption            *string                        `pulumi:"detachOption"`
	DiskIOPSReadWrite       float64                        `pulumi:"diskIOPSReadWrite"`
	DiskMBpsReadWrite       float64                        `pulumi:"diskMBpsReadWrite"`
	DiskSizeGB              *int                           `pulumi:"diskSizeGB"`
	Image                   *VirtualHardDiskResponse       `pulumi:"image"`
	Lun                     int                            `pulumi:"lun"`
	ManagedDisk             *ManagedDiskParametersResponse `pulumi:"managedDisk"`
	Name                    *string                        `pulumi:"name"`
	ToBeDetached            *bool                          `pulumi:"toBeDetached"`
	Vhd                     *VirtualHardDiskResponse       `pulumi:"vhd"`
	WriteAcceleratorEnabled *bool                          `pulumi:"writeAcceleratorEnabled"`
}

type DataDiskResponseOutput struct{ *pulumi.OutputState }

func (DataDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDiskResponse)(nil)).Elem()
}

func (o DataDiskResponseOutput) ToDataDiskResponseOutput() DataDiskResponseOutput {
	return o
}

func (o DataDiskResponseOutput) ToDataDiskResponseOutputWithContext(ctx context.Context) DataDiskResponseOutput {
	return o
}

func (o DataDiskResponseOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *string { return v.Caching }).(pulumi.StringPtrOutput)
}

func (o DataDiskResponseOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v DataDiskResponse) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o DataDiskResponseOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o DataDiskResponseOutput) DetachOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *string { return v.DetachOption }).(pulumi.StringPtrOutput)
}

func (o DataDiskResponseOutput) DiskIOPSReadWrite() pulumi.Float64Output {
	return o.ApplyT(func(v DataDiskResponse) float64 { return v.DiskIOPSReadWrite }).(pulumi.Float64Output)
}

func (o DataDiskResponseOutput) DiskMBpsReadWrite() pulumi.Float64Output {
	return o.ApplyT(func(v DataDiskResponse) float64 { return v.DiskMBpsReadWrite }).(pulumi.Float64Output)
}

func (o DataDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o DataDiskResponseOutput) Image() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *VirtualHardDiskResponse { return v.Image }).(VirtualHardDiskResponsePtrOutput)
}

func (o DataDiskResponseOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v DataDiskResponse) int { return v.Lun }).(pulumi.IntOutput)
}

func (o DataDiskResponseOutput) ManagedDisk() ManagedDiskParametersResponsePtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *ManagedDiskParametersResponse { return v.ManagedDisk }).(ManagedDiskParametersResponsePtrOutput)
}

func (o DataDiskResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DataDiskResponseOutput) ToBeDetached() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *bool { return v.ToBeDetached }).(pulumi.BoolPtrOutput)
}

func (o DataDiskResponseOutput) Vhd() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *VirtualHardDiskResponse { return v.Vhd }).(VirtualHardDiskResponsePtrOutput)
}

func (o DataDiskResponseOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *bool { return v.WriteAcceleratorEnabled }).(pulumi.BoolPtrOutput)
}

type DataDiskResponseArrayOutput struct{ *pulumi.OutputState }

func (DataDiskResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDiskResponse)(nil)).Elem()
}

func (o DataDiskResponseArrayOutput) ToDataDiskResponseArrayOutput() DataDiskResponseArrayOutput {
	return o
}

func (o DataDiskResponseArrayOutput) ToDataDiskResponseArrayOutputWithContext(ctx context.Context) DataDiskResponseArrayOutput {
	return o
}

func (o DataDiskResponseArrayOutput) Index(i pulumi.IntInput) DataDiskResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDiskResponse {
		return vs[0].([]DataDiskResponse)[vs[1].(int)]
	}).(DataDiskResponseOutput)
}

type DedicatedHostAllocatableVMResponse struct {
	Count  *float64 `pulumi:"count"`
	VmSize *string  `pulumi:"vmSize"`
}

type DedicatedHostAllocatableVMResponseOutput struct{ *pulumi.OutputState }

func (DedicatedHostAllocatableVMResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedHostAllocatableVMResponse)(nil)).Elem()
}

func (o DedicatedHostAllocatableVMResponseOutput) ToDedicatedHostAllocatableVMResponseOutput() DedicatedHostAllocatableVMResponseOutput {
	return o
}

func (o DedicatedHostAllocatableVMResponseOutput) ToDedicatedHostAllocatableVMResponseOutputWithContext(ctx context.Context) DedicatedHostAllocatableVMResponseOutput {
	return o
}

func (o DedicatedHostAllocatableVMResponseOutput) Count() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DedicatedHostAllocatableVMResponse) *float64 { return v.Count }).(pulumi.Float64PtrOutput)
}

func (o DedicatedHostAllocatableVMResponseOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DedicatedHostAllocatableVMResponse) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type DedicatedHostAllocatableVMResponseArrayOutput struct{ *pulumi.OutputState }

func (DedicatedHostAllocatableVMResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DedicatedHostAllocatableVMResponse)(nil)).Elem()
}

func (o DedicatedHostAllocatableVMResponseArrayOutput) ToDedicatedHostAllocatableVMResponseArrayOutput() DedicatedHostAllocatableVMResponseArrayOutput {
	return o
}

func (o DedicatedHostAllocatableVMResponseArrayOutput) ToDedicatedHostAllocatableVMResponseArrayOutputWithContext(ctx context.Context) DedicatedHostAllocatableVMResponseArrayOutput {
	return o
}

func (o DedicatedHostAllocatableVMResponseArrayOutput) Index(i pulumi.IntInput) DedicatedHostAllocatableVMResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DedicatedHostAllocatableVMResponse {
		return vs[0].([]DedicatedHostAllocatableVMResponse)[vs[1].(int)]
	}).(DedicatedHostAllocatableVMResponseOutput)
}

type DedicatedHostAvailableCapacityResponse struct {
	AllocatableVMs []DedicatedHostAllocatableVMResponse `pulumi:"allocatableVMs"`
}

type DedicatedHostAvailableCapacityResponseOutput struct{ *pulumi.OutputState }

func (DedicatedHostAvailableCapacityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedHostAvailableCapacityResponse)(nil)).Elem()
}

func (o DedicatedHostAvailableCapacityResponseOutput) ToDedicatedHostAvailableCapacityResponseOutput() DedicatedHostAvailableCapacityResponseOutput {
	return o
}

func (o DedicatedHostAvailableCapacityResponseOutput) ToDedicatedHostAvailableCapacityResponseOutputWithContext(ctx context.Context) DedicatedHostAvailableCapacityResponseOutput {
	return o
}

func (o DedicatedHostAvailableCapacityResponseOutput) AllocatableVMs() DedicatedHostAllocatableVMResponseArrayOutput {
	return o.ApplyT(func(v DedicatedHostAvailableCapacityResponse) []DedicatedHostAllocatableVMResponse {
		return v.AllocatableVMs
	}).(DedicatedHostAllocatableVMResponseArrayOutput)
}

type DedicatedHostAvailableCapacityResponsePtrOutput struct{ *pulumi.OutputState }

func (DedicatedHostAvailableCapacityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHostAvailableCapacityResponse)(nil)).Elem()
}

func (o DedicatedHostAvailableCapacityResponsePtrOutput) ToDedicatedHostAvailableCapacityResponsePtrOutput() DedicatedHostAvailableCapacityResponsePtrOutput {
	return o
}

func (o DedicatedHostAvailableCapacityResponsePtrOutput) ToDedicatedHostAvailableCapacityResponsePtrOutputWithContext(ctx context.Context) DedicatedHostAvailableCapacityResponsePtrOutput {
	return o
}

func (o DedicatedHostAvailableCapacityResponsePtrOutput) Elem() DedicatedHostAvailableCapacityResponseOutput {
	return o.ApplyT(func(v *DedicatedHostAvailableCapacityResponse) DedicatedHostAvailableCapacityResponse {
		if v != nil {
			return *v
		}
		var ret DedicatedHostAvailableCapacityResponse
		return ret
	}).(DedicatedHostAvailableCapacityResponseOutput)
}

func (o DedicatedHostAvailableCapacityResponsePtrOutput) AllocatableVMs() DedicatedHostAllocatableVMResponseArrayOutput {
	return o.ApplyT(func(v *DedicatedHostAvailableCapacityResponse) []DedicatedHostAllocatableVMResponse {
		if v == nil {
			return nil
		}
		return v.AllocatableVMs
	}).(DedicatedHostAllocatableVMResponseArrayOutput)
}

type DedicatedHostGroupInstanceViewResponse struct {
	Hosts []DedicatedHostInstanceViewWithNameResponse `pulumi:"hosts"`
}

type DedicatedHostGroupInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (DedicatedHostGroupInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedHostGroupInstanceViewResponse)(nil)).Elem()
}

func (o DedicatedHostGroupInstanceViewResponseOutput) ToDedicatedHostGroupInstanceViewResponseOutput() DedicatedHostGroupInstanceViewResponseOutput {
	return o
}

func (o DedicatedHostGroupInstanceViewResponseOutput) ToDedicatedHostGroupInstanceViewResponseOutputWithContext(ctx context.Context) DedicatedHostGroupInstanceViewResponseOutput {
	return o
}

func (o DedicatedHostGroupInstanceViewResponseOutput) Hosts() DedicatedHostInstanceViewWithNameResponseArrayOutput {
	return o.ApplyT(func(v DedicatedHostGroupInstanceViewResponse) []DedicatedHostInstanceViewWithNameResponse {
		return v.Hosts
	}).(DedicatedHostInstanceViewWithNameResponseArrayOutput)
}

type DedicatedHostGroupPropertiesAdditionalCapabilities struct {
	UltraSSDEnabled *bool `pulumi:"ultraSSDEnabled"`
}





type DedicatedHostGroupPropertiesAdditionalCapabilitiesInput interface {
	pulumi.Input

	ToDedicatedHostGroupPropertiesAdditionalCapabilitiesOutput() DedicatedHostGroupPropertiesAdditionalCapabilitiesOutput
	ToDedicatedHostGroupPropertiesAdditionalCapabilitiesOutputWithContext(context.Context) DedicatedHostGroupPropertiesAdditionalCapabilitiesOutput
}

type DedicatedHostGroupPropertiesAdditionalCapabilitiesArgs struct {
	UltraSSDEnabled pulumi.BoolPtrInput `pulumi:"ultraSSDEnabled"`
}

func (DedicatedHostGroupPropertiesAdditionalCapabilitiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedHostGroupPropertiesAdditionalCapabilities)(nil)).Elem()
}

func (i DedicatedHostGroupPropertiesAdditionalCapabilitiesArgs) ToDedicatedHostGroupPropertiesAdditionalCapabilitiesOutput() DedicatedHostGroupPropertiesAdditionalCapabilitiesOutput {
	return i.ToDedicatedHostGroupPropertiesAdditionalCapabilitiesOutputWithContext(context.Background())
}

func (i DedicatedHostGroupPropertiesAdditionalCapabilitiesArgs) ToDedicatedHostGroupPropertiesAdditionalCapabilitiesOutputWithContext(ctx context.Context) DedicatedHostGroupPropertiesAdditionalCapabilitiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostGroupPropertiesAdditionalCapabilitiesOutput)
}

func (i DedicatedHostGroupPropertiesAdditionalCapabilitiesArgs) ToDedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput() DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput {
	return i.ToDedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutputWithContext(context.Background())
}

func (i DedicatedHostGroupPropertiesAdditionalCapabilitiesArgs) ToDedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutputWithContext(ctx context.Context) DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostGroupPropertiesAdditionalCapabilitiesOutput).ToDedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutputWithContext(ctx)
}









type DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrInput interface {
	pulumi.Input

	ToDedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput() DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput
	ToDedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutputWithContext(context.Context) DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput
}

type dedicatedHostGroupPropertiesAdditionalCapabilitiesPtrType DedicatedHostGroupPropertiesAdditionalCapabilitiesArgs

func DedicatedHostGroupPropertiesAdditionalCapabilitiesPtr(v *DedicatedHostGroupPropertiesAdditionalCapabilitiesArgs) DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrInput {
	return (*dedicatedHostGroupPropertiesAdditionalCapabilitiesPtrType)(v)
}

func (*dedicatedHostGroupPropertiesAdditionalCapabilitiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHostGroupPropertiesAdditionalCapabilities)(nil)).Elem()
}

func (i *dedicatedHostGroupPropertiesAdditionalCapabilitiesPtrType) ToDedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput() DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput {
	return i.ToDedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutputWithContext(context.Background())
}

func (i *dedicatedHostGroupPropertiesAdditionalCapabilitiesPtrType) ToDedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutputWithContext(ctx context.Context) DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput)
}

type DedicatedHostGroupPropertiesAdditionalCapabilitiesOutput struct{ *pulumi.OutputState }

func (DedicatedHostGroupPropertiesAdditionalCapabilitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedHostGroupPropertiesAdditionalCapabilities)(nil)).Elem()
}

func (o DedicatedHostGroupPropertiesAdditionalCapabilitiesOutput) ToDedicatedHostGroupPropertiesAdditionalCapabilitiesOutput() DedicatedHostGroupPropertiesAdditionalCapabilitiesOutput {
	return o
}

func (o DedicatedHostGroupPropertiesAdditionalCapabilitiesOutput) ToDedicatedHostGroupPropertiesAdditionalCapabilitiesOutputWithContext(ctx context.Context) DedicatedHostGroupPropertiesAdditionalCapabilitiesOutput {
	return o
}

func (o DedicatedHostGroupPropertiesAdditionalCapabilitiesOutput) ToDedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput() DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput {
	return o.ToDedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutputWithContext(context.Background())
}

func (o DedicatedHostGroupPropertiesAdditionalCapabilitiesOutput) ToDedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutputWithContext(ctx context.Context) DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DedicatedHostGroupPropertiesAdditionalCapabilities) *DedicatedHostGroupPropertiesAdditionalCapabilities {
		return &v
	}).(DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput)
}

func (o DedicatedHostGroupPropertiesAdditionalCapabilitiesOutput) UltraSSDEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DedicatedHostGroupPropertiesAdditionalCapabilities) *bool { return v.UltraSSDEnabled }).(pulumi.BoolPtrOutput)
}

type DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput struct{ *pulumi.OutputState }

func (DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHostGroupPropertiesAdditionalCapabilities)(nil)).Elem()
}

func (o DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput) ToDedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput() DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput {
	return o
}

func (o DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput) ToDedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutputWithContext(ctx context.Context) DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput {
	return o
}

func (o DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput) Elem() DedicatedHostGroupPropertiesAdditionalCapabilitiesOutput {
	return o.ApplyT(func(v *DedicatedHostGroupPropertiesAdditionalCapabilities) DedicatedHostGroupPropertiesAdditionalCapabilities {
		if v != nil {
			return *v
		}
		var ret DedicatedHostGroupPropertiesAdditionalCapabilities
		return ret
	}).(DedicatedHostGroupPropertiesAdditionalCapabilitiesOutput)
}

func (o DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput) UltraSSDEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DedicatedHostGroupPropertiesAdditionalCapabilities) *bool {
		if v == nil {
			return nil
		}
		return v.UltraSSDEnabled
	}).(pulumi.BoolPtrOutput)
}

type DedicatedHostGroupPropertiesResponseAdditionalCapabilities struct {
	UltraSSDEnabled *bool `pulumi:"ultraSSDEnabled"`
}

type DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesOutput struct{ *pulumi.OutputState }

func (DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedHostGroupPropertiesResponseAdditionalCapabilities)(nil)).Elem()
}

func (o DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesOutput) ToDedicatedHostGroupPropertiesResponseAdditionalCapabilitiesOutput() DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesOutput {
	return o
}

func (o DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesOutput) ToDedicatedHostGroupPropertiesResponseAdditionalCapabilitiesOutputWithContext(ctx context.Context) DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesOutput {
	return o
}

func (o DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesOutput) UltraSSDEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DedicatedHostGroupPropertiesResponseAdditionalCapabilities) *bool { return v.UltraSSDEnabled }).(pulumi.BoolPtrOutput)
}

type DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesPtrOutput struct{ *pulumi.OutputState }

func (DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHostGroupPropertiesResponseAdditionalCapabilities)(nil)).Elem()
}

func (o DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesPtrOutput) ToDedicatedHostGroupPropertiesResponseAdditionalCapabilitiesPtrOutput() DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesPtrOutput {
	return o
}

func (o DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesPtrOutput) ToDedicatedHostGroupPropertiesResponseAdditionalCapabilitiesPtrOutputWithContext(ctx context.Context) DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesPtrOutput {
	return o
}

func (o DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesPtrOutput) Elem() DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesOutput {
	return o.ApplyT(func(v *DedicatedHostGroupPropertiesResponseAdditionalCapabilities) DedicatedHostGroupPropertiesResponseAdditionalCapabilities {
		if v != nil {
			return *v
		}
		var ret DedicatedHostGroupPropertiesResponseAdditionalCapabilities
		return ret
	}).(DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesOutput)
}

func (o DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesPtrOutput) UltraSSDEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DedicatedHostGroupPropertiesResponseAdditionalCapabilities) *bool {
		if v == nil {
			return nil
		}
		return v.UltraSSDEnabled
	}).(pulumi.BoolPtrOutput)
}

type DedicatedHostInstanceViewResponse struct {
	AssetId           string                                  `pulumi:"assetId"`
	AvailableCapacity *DedicatedHostAvailableCapacityResponse `pulumi:"availableCapacity"`
	Statuses          []InstanceViewStatusResponse            `pulumi:"statuses"`
}

type DedicatedHostInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (DedicatedHostInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedHostInstanceViewResponse)(nil)).Elem()
}

func (o DedicatedHostInstanceViewResponseOutput) ToDedicatedHostInstanceViewResponseOutput() DedicatedHostInstanceViewResponseOutput {
	return o
}

func (o DedicatedHostInstanceViewResponseOutput) ToDedicatedHostInstanceViewResponseOutputWithContext(ctx context.Context) DedicatedHostInstanceViewResponseOutput {
	return o
}

func (o DedicatedHostInstanceViewResponseOutput) AssetId() pulumi.StringOutput {
	return o.ApplyT(func(v DedicatedHostInstanceViewResponse) string { return v.AssetId }).(pulumi.StringOutput)
}

func (o DedicatedHostInstanceViewResponseOutput) AvailableCapacity() DedicatedHostAvailableCapacityResponsePtrOutput {
	return o.ApplyT(func(v DedicatedHostInstanceViewResponse) *DedicatedHostAvailableCapacityResponse {
		return v.AvailableCapacity
	}).(DedicatedHostAvailableCapacityResponsePtrOutput)
}

func (o DedicatedHostInstanceViewResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v DedicatedHostInstanceViewResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

type DedicatedHostInstanceViewWithNameResponse struct {
	AssetId           string                                  `pulumi:"assetId"`
	AvailableCapacity *DedicatedHostAvailableCapacityResponse `pulumi:"availableCapacity"`
	Name              string                                  `pulumi:"name"`
	Statuses          []InstanceViewStatusResponse            `pulumi:"statuses"`
}

type DedicatedHostInstanceViewWithNameResponseOutput struct{ *pulumi.OutputState }

func (DedicatedHostInstanceViewWithNameResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DedicatedHostInstanceViewWithNameResponse)(nil)).Elem()
}

func (o DedicatedHostInstanceViewWithNameResponseOutput) ToDedicatedHostInstanceViewWithNameResponseOutput() DedicatedHostInstanceViewWithNameResponseOutput {
	return o
}

func (o DedicatedHostInstanceViewWithNameResponseOutput) ToDedicatedHostInstanceViewWithNameResponseOutputWithContext(ctx context.Context) DedicatedHostInstanceViewWithNameResponseOutput {
	return o
}

func (o DedicatedHostInstanceViewWithNameResponseOutput) AssetId() pulumi.StringOutput {
	return o.ApplyT(func(v DedicatedHostInstanceViewWithNameResponse) string { return v.AssetId }).(pulumi.StringOutput)
}

func (o DedicatedHostInstanceViewWithNameResponseOutput) AvailableCapacity() DedicatedHostAvailableCapacityResponsePtrOutput {
	return o.ApplyT(func(v DedicatedHostInstanceViewWithNameResponse) *DedicatedHostAvailableCapacityResponse {
		return v.AvailableCapacity
	}).(DedicatedHostAvailableCapacityResponsePtrOutput)
}

func (o DedicatedHostInstanceViewWithNameResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DedicatedHostInstanceViewWithNameResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o DedicatedHostInstanceViewWithNameResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v DedicatedHostInstanceViewWithNameResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

type DedicatedHostInstanceViewWithNameResponseArrayOutput struct{ *pulumi.OutputState }

func (DedicatedHostInstanceViewWithNameResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DedicatedHostInstanceViewWithNameResponse)(nil)).Elem()
}

func (o DedicatedHostInstanceViewWithNameResponseArrayOutput) ToDedicatedHostInstanceViewWithNameResponseArrayOutput() DedicatedHostInstanceViewWithNameResponseArrayOutput {
	return o
}

func (o DedicatedHostInstanceViewWithNameResponseArrayOutput) ToDedicatedHostInstanceViewWithNameResponseArrayOutputWithContext(ctx context.Context) DedicatedHostInstanceViewWithNameResponseArrayOutput {
	return o
}

func (o DedicatedHostInstanceViewWithNameResponseArrayOutput) Index(i pulumi.IntInput) DedicatedHostInstanceViewWithNameResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DedicatedHostInstanceViewWithNameResponse {
		return vs[0].([]DedicatedHostInstanceViewWithNameResponse)[vs[1].(int)]
	}).(DedicatedHostInstanceViewWithNameResponseOutput)
}

type DiagnosticsProfile struct {
	BootDiagnostics *BootDiagnostics `pulumi:"bootDiagnostics"`
}





type DiagnosticsProfileInput interface {
	pulumi.Input

	ToDiagnosticsProfileOutput() DiagnosticsProfileOutput
	ToDiagnosticsProfileOutputWithContext(context.Context) DiagnosticsProfileOutput
}

type DiagnosticsProfileArgs struct {
	BootDiagnostics BootDiagnosticsPtrInput `pulumi:"bootDiagnostics"`
}

func (DiagnosticsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsProfile)(nil)).Elem()
}

func (i DiagnosticsProfileArgs) ToDiagnosticsProfileOutput() DiagnosticsProfileOutput {
	return i.ToDiagnosticsProfileOutputWithContext(context.Background())
}

func (i DiagnosticsProfileArgs) ToDiagnosticsProfileOutputWithContext(ctx context.Context) DiagnosticsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsProfileOutput)
}

func (i DiagnosticsProfileArgs) ToDiagnosticsProfilePtrOutput() DiagnosticsProfilePtrOutput {
	return i.ToDiagnosticsProfilePtrOutputWithContext(context.Background())
}

func (i DiagnosticsProfileArgs) ToDiagnosticsProfilePtrOutputWithContext(ctx context.Context) DiagnosticsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsProfileOutput).ToDiagnosticsProfilePtrOutputWithContext(ctx)
}









type DiagnosticsProfilePtrInput interface {
	pulumi.Input

	ToDiagnosticsProfilePtrOutput() DiagnosticsProfilePtrOutput
	ToDiagnosticsProfilePtrOutputWithContext(context.Context) DiagnosticsProfilePtrOutput
}

type diagnosticsProfilePtrType DiagnosticsProfileArgs

func DiagnosticsProfilePtr(v *DiagnosticsProfileArgs) DiagnosticsProfilePtrInput {
	return (*diagnosticsProfilePtrType)(v)
}

func (*diagnosticsProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsProfile)(nil)).Elem()
}

func (i *diagnosticsProfilePtrType) ToDiagnosticsProfilePtrOutput() DiagnosticsProfilePtrOutput {
	return i.ToDiagnosticsProfilePtrOutputWithContext(context.Background())
}

func (i *diagnosticsProfilePtrType) ToDiagnosticsProfilePtrOutputWithContext(ctx context.Context) DiagnosticsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsProfilePtrOutput)
}

type DiagnosticsProfileOutput struct{ *pulumi.OutputState }

func (DiagnosticsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsProfile)(nil)).Elem()
}

func (o DiagnosticsProfileOutput) ToDiagnosticsProfileOutput() DiagnosticsProfileOutput {
	return o
}

func (o DiagnosticsProfileOutput) ToDiagnosticsProfileOutputWithContext(ctx context.Context) DiagnosticsProfileOutput {
	return o
}

func (o DiagnosticsProfileOutput) ToDiagnosticsProfilePtrOutput() DiagnosticsProfilePtrOutput {
	return o.ToDiagnosticsProfilePtrOutputWithContext(context.Background())
}

func (o DiagnosticsProfileOutput) ToDiagnosticsProfilePtrOutputWithContext(ctx context.Context) DiagnosticsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiagnosticsProfile) *DiagnosticsProfile {
		return &v
	}).(DiagnosticsProfilePtrOutput)
}

func (o DiagnosticsProfileOutput) BootDiagnostics() BootDiagnosticsPtrOutput {
	return o.ApplyT(func(v DiagnosticsProfile) *BootDiagnostics { return v.BootDiagnostics }).(BootDiagnosticsPtrOutput)
}

type DiagnosticsProfilePtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsProfile)(nil)).Elem()
}

func (o DiagnosticsProfilePtrOutput) ToDiagnosticsProfilePtrOutput() DiagnosticsProfilePtrOutput {
	return o
}

func (o DiagnosticsProfilePtrOutput) ToDiagnosticsProfilePtrOutputWithContext(ctx context.Context) DiagnosticsProfilePtrOutput {
	return o
}

func (o DiagnosticsProfilePtrOutput) Elem() DiagnosticsProfileOutput {
	return o.ApplyT(func(v *DiagnosticsProfile) DiagnosticsProfile {
		if v != nil {
			return *v
		}
		var ret DiagnosticsProfile
		return ret
	}).(DiagnosticsProfileOutput)
}

func (o DiagnosticsProfilePtrOutput) BootDiagnostics() BootDiagnosticsPtrOutput {
	return o.ApplyT(func(v *DiagnosticsProfile) *BootDiagnostics {
		if v == nil {
			return nil
		}
		return v.BootDiagnostics
	}).(BootDiagnosticsPtrOutput)
}

type DiagnosticsProfileResponse struct {
	BootDiagnostics *BootDiagnosticsResponse `pulumi:"bootDiagnostics"`
}

type DiagnosticsProfileResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticsProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsProfileResponse)(nil)).Elem()
}

func (o DiagnosticsProfileResponseOutput) ToDiagnosticsProfileResponseOutput() DiagnosticsProfileResponseOutput {
	return o
}

func (o DiagnosticsProfileResponseOutput) ToDiagnosticsProfileResponseOutputWithContext(ctx context.Context) DiagnosticsProfileResponseOutput {
	return o
}

func (o DiagnosticsProfileResponseOutput) BootDiagnostics() BootDiagnosticsResponsePtrOutput {
	return o.ApplyT(func(v DiagnosticsProfileResponse) *BootDiagnosticsResponse { return v.BootDiagnostics }).(BootDiagnosticsResponsePtrOutput)
}

type DiagnosticsProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsProfileResponse)(nil)).Elem()
}

func (o DiagnosticsProfileResponsePtrOutput) ToDiagnosticsProfileResponsePtrOutput() DiagnosticsProfileResponsePtrOutput {
	return o
}

func (o DiagnosticsProfileResponsePtrOutput) ToDiagnosticsProfileResponsePtrOutputWithContext(ctx context.Context) DiagnosticsProfileResponsePtrOutput {
	return o
}

func (o DiagnosticsProfileResponsePtrOutput) Elem() DiagnosticsProfileResponseOutput {
	return o.ApplyT(func(v *DiagnosticsProfileResponse) DiagnosticsProfileResponse {
		if v != nil {
			return *v
		}
		var ret DiagnosticsProfileResponse
		return ret
	}).(DiagnosticsProfileResponseOutput)
}

func (o DiagnosticsProfileResponsePtrOutput) BootDiagnostics() BootDiagnosticsResponsePtrOutput {
	return o.ApplyT(func(v *DiagnosticsProfileResponse) *BootDiagnosticsResponse {
		if v == nil {
			return nil
		}
		return v.BootDiagnostics
	}).(BootDiagnosticsResponsePtrOutput)
}

type DiffDiskSettings struct {
	Option    *string `pulumi:"option"`
	Placement *string `pulumi:"placement"`
}





type DiffDiskSettingsInput interface {
	pulumi.Input

	ToDiffDiskSettingsOutput() DiffDiskSettingsOutput
	ToDiffDiskSettingsOutputWithContext(context.Context) DiffDiskSettingsOutput
}

type DiffDiskSettingsArgs struct {
	Option    pulumi.StringPtrInput `pulumi:"option"`
	Placement pulumi.StringPtrInput `pulumi:"placement"`
}

func (DiffDiskSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiffDiskSettings)(nil)).Elem()
}

func (i DiffDiskSettingsArgs) ToDiffDiskSettingsOutput() DiffDiskSettingsOutput {
	return i.ToDiffDiskSettingsOutputWithContext(context.Background())
}

func (i DiffDiskSettingsArgs) ToDiffDiskSettingsOutputWithContext(ctx context.Context) DiffDiskSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiffDiskSettingsOutput)
}

func (i DiffDiskSettingsArgs) ToDiffDiskSettingsPtrOutput() DiffDiskSettingsPtrOutput {
	return i.ToDiffDiskSettingsPtrOutputWithContext(context.Background())
}

func (i DiffDiskSettingsArgs) ToDiffDiskSettingsPtrOutputWithContext(ctx context.Context) DiffDiskSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiffDiskSettingsOutput).ToDiffDiskSettingsPtrOutputWithContext(ctx)
}









type DiffDiskSettingsPtrInput interface {
	pulumi.Input

	ToDiffDiskSettingsPtrOutput() DiffDiskSettingsPtrOutput
	ToDiffDiskSettingsPtrOutputWithContext(context.Context) DiffDiskSettingsPtrOutput
}

type diffDiskSettingsPtrType DiffDiskSettingsArgs

func DiffDiskSettingsPtr(v *DiffDiskSettingsArgs) DiffDiskSettingsPtrInput {
	return (*diffDiskSettingsPtrType)(v)
}

func (*diffDiskSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiffDiskSettings)(nil)).Elem()
}

func (i *diffDiskSettingsPtrType) ToDiffDiskSettingsPtrOutput() DiffDiskSettingsPtrOutput {
	return i.ToDiffDiskSettingsPtrOutputWithContext(context.Background())
}

func (i *diffDiskSettingsPtrType) ToDiffDiskSettingsPtrOutputWithContext(ctx context.Context) DiffDiskSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiffDiskSettingsPtrOutput)
}

type DiffDiskSettingsOutput struct{ *pulumi.OutputState }

func (DiffDiskSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiffDiskSettings)(nil)).Elem()
}

func (o DiffDiskSettingsOutput) ToDiffDiskSettingsOutput() DiffDiskSettingsOutput {
	return o
}

func (o DiffDiskSettingsOutput) ToDiffDiskSettingsOutputWithContext(ctx context.Context) DiffDiskSettingsOutput {
	return o
}

func (o DiffDiskSettingsOutput) ToDiffDiskSettingsPtrOutput() DiffDiskSettingsPtrOutput {
	return o.ToDiffDiskSettingsPtrOutputWithContext(context.Background())
}

func (o DiffDiskSettingsOutput) ToDiffDiskSettingsPtrOutputWithContext(ctx context.Context) DiffDiskSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiffDiskSettings) *DiffDiskSettings {
		return &v
	}).(DiffDiskSettingsPtrOutput)
}

func (o DiffDiskSettingsOutput) Option() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiffDiskSettings) *string { return v.Option }).(pulumi.StringPtrOutput)
}

func (o DiffDiskSettingsOutput) Placement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiffDiskSettings) *string { return v.Placement }).(pulumi.StringPtrOutput)
}

type DiffDiskSettingsPtrOutput struct{ *pulumi.OutputState }

func (DiffDiskSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiffDiskSettings)(nil)).Elem()
}

func (o DiffDiskSettingsPtrOutput) ToDiffDiskSettingsPtrOutput() DiffDiskSettingsPtrOutput {
	return o
}

func (o DiffDiskSettingsPtrOutput) ToDiffDiskSettingsPtrOutputWithContext(ctx context.Context) DiffDiskSettingsPtrOutput {
	return o
}

func (o DiffDiskSettingsPtrOutput) Elem() DiffDiskSettingsOutput {
	return o.ApplyT(func(v *DiffDiskSettings) DiffDiskSettings {
		if v != nil {
			return *v
		}
		var ret DiffDiskSettings
		return ret
	}).(DiffDiskSettingsOutput)
}

func (o DiffDiskSettingsPtrOutput) Option() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiffDiskSettings) *string {
		if v == nil {
			return nil
		}
		return v.Option
	}).(pulumi.StringPtrOutput)
}

func (o DiffDiskSettingsPtrOutput) Placement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiffDiskSettings) *string {
		if v == nil {
			return nil
		}
		return v.Placement
	}).(pulumi.StringPtrOutput)
}

type DiffDiskSettingsResponse struct {
	Option    *string `pulumi:"option"`
	Placement *string `pulumi:"placement"`
}

type DiffDiskSettingsResponseOutput struct{ *pulumi.OutputState }

func (DiffDiskSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiffDiskSettingsResponse)(nil)).Elem()
}

func (o DiffDiskSettingsResponseOutput) ToDiffDiskSettingsResponseOutput() DiffDiskSettingsResponseOutput {
	return o
}

func (o DiffDiskSettingsResponseOutput) ToDiffDiskSettingsResponseOutputWithContext(ctx context.Context) DiffDiskSettingsResponseOutput {
	return o
}

func (o DiffDiskSettingsResponseOutput) Option() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiffDiskSettingsResponse) *string { return v.Option }).(pulumi.StringPtrOutput)
}

func (o DiffDiskSettingsResponseOutput) Placement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiffDiskSettingsResponse) *string { return v.Placement }).(pulumi.StringPtrOutput)
}

type DiffDiskSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (DiffDiskSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiffDiskSettingsResponse)(nil)).Elem()
}

func (o DiffDiskSettingsResponsePtrOutput) ToDiffDiskSettingsResponsePtrOutput() DiffDiskSettingsResponsePtrOutput {
	return o
}

func (o DiffDiskSettingsResponsePtrOutput) ToDiffDiskSettingsResponsePtrOutputWithContext(ctx context.Context) DiffDiskSettingsResponsePtrOutput {
	return o
}

func (o DiffDiskSettingsResponsePtrOutput) Elem() DiffDiskSettingsResponseOutput {
	return o.ApplyT(func(v *DiffDiskSettingsResponse) DiffDiskSettingsResponse {
		if v != nil {
			return *v
		}
		var ret DiffDiskSettingsResponse
		return ret
	}).(DiffDiskSettingsResponseOutput)
}

func (o DiffDiskSettingsResponsePtrOutput) Option() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiffDiskSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Option
	}).(pulumi.StringPtrOutput)
}

func (o DiffDiskSettingsResponsePtrOutput) Placement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiffDiskSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Placement
	}).(pulumi.StringPtrOutput)
}

type DiskEncryptionSetParameters struct {
	Id *string `pulumi:"id"`
}





type DiskEncryptionSetParametersInput interface {
	pulumi.Input

	ToDiskEncryptionSetParametersOutput() DiskEncryptionSetParametersOutput
	ToDiskEncryptionSetParametersOutputWithContext(context.Context) DiskEncryptionSetParametersOutput
}

type DiskEncryptionSetParametersArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (DiskEncryptionSetParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSetParameters)(nil)).Elem()
}

func (i DiskEncryptionSetParametersArgs) ToDiskEncryptionSetParametersOutput() DiskEncryptionSetParametersOutput {
	return i.ToDiskEncryptionSetParametersOutputWithContext(context.Background())
}

func (i DiskEncryptionSetParametersArgs) ToDiskEncryptionSetParametersOutputWithContext(ctx context.Context) DiskEncryptionSetParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSetParametersOutput)
}

func (i DiskEncryptionSetParametersArgs) ToDiskEncryptionSetParametersPtrOutput() DiskEncryptionSetParametersPtrOutput {
	return i.ToDiskEncryptionSetParametersPtrOutputWithContext(context.Background())
}

func (i DiskEncryptionSetParametersArgs) ToDiskEncryptionSetParametersPtrOutputWithContext(ctx context.Context) DiskEncryptionSetParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSetParametersOutput).ToDiskEncryptionSetParametersPtrOutputWithContext(ctx)
}









type DiskEncryptionSetParametersPtrInput interface {
	pulumi.Input

	ToDiskEncryptionSetParametersPtrOutput() DiskEncryptionSetParametersPtrOutput
	ToDiskEncryptionSetParametersPtrOutputWithContext(context.Context) DiskEncryptionSetParametersPtrOutput
}

type diskEncryptionSetParametersPtrType DiskEncryptionSetParametersArgs

func DiskEncryptionSetParametersPtr(v *DiskEncryptionSetParametersArgs) DiskEncryptionSetParametersPtrInput {
	return (*diskEncryptionSetParametersPtrType)(v)
}

func (*diskEncryptionSetParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSetParameters)(nil)).Elem()
}

func (i *diskEncryptionSetParametersPtrType) ToDiskEncryptionSetParametersPtrOutput() DiskEncryptionSetParametersPtrOutput {
	return i.ToDiskEncryptionSetParametersPtrOutputWithContext(context.Background())
}

func (i *diskEncryptionSetParametersPtrType) ToDiskEncryptionSetParametersPtrOutputWithContext(ctx context.Context) DiskEncryptionSetParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSetParametersPtrOutput)
}

type DiskEncryptionSetParametersOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSetParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSetParameters)(nil)).Elem()
}

func (o DiskEncryptionSetParametersOutput) ToDiskEncryptionSetParametersOutput() DiskEncryptionSetParametersOutput {
	return o
}

func (o DiskEncryptionSetParametersOutput) ToDiskEncryptionSetParametersOutputWithContext(ctx context.Context) DiskEncryptionSetParametersOutput {
	return o
}

func (o DiskEncryptionSetParametersOutput) ToDiskEncryptionSetParametersPtrOutput() DiskEncryptionSetParametersPtrOutput {
	return o.ToDiskEncryptionSetParametersPtrOutputWithContext(context.Background())
}

func (o DiskEncryptionSetParametersOutput) ToDiskEncryptionSetParametersPtrOutputWithContext(ctx context.Context) DiskEncryptionSetParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskEncryptionSetParameters) *DiskEncryptionSetParameters {
		return &v
	}).(DiskEncryptionSetParametersPtrOutput)
}

func (o DiskEncryptionSetParametersOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskEncryptionSetParameters) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type DiskEncryptionSetParametersPtrOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSetParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSetParameters)(nil)).Elem()
}

func (o DiskEncryptionSetParametersPtrOutput) ToDiskEncryptionSetParametersPtrOutput() DiskEncryptionSetParametersPtrOutput {
	return o
}

func (o DiskEncryptionSetParametersPtrOutput) ToDiskEncryptionSetParametersPtrOutputWithContext(ctx context.Context) DiskEncryptionSetParametersPtrOutput {
	return o
}

func (o DiskEncryptionSetParametersPtrOutput) Elem() DiskEncryptionSetParametersOutput {
	return o.ApplyT(func(v *DiskEncryptionSetParameters) DiskEncryptionSetParameters {
		if v != nil {
			return *v
		}
		var ret DiskEncryptionSetParameters
		return ret
	}).(DiskEncryptionSetParametersOutput)
}

func (o DiskEncryptionSetParametersPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSetParameters) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type DiskEncryptionSetParametersResponse struct {
	Id *string `pulumi:"id"`
}

type DiskEncryptionSetParametersResponseOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSetParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSetParametersResponse)(nil)).Elem()
}

func (o DiskEncryptionSetParametersResponseOutput) ToDiskEncryptionSetParametersResponseOutput() DiskEncryptionSetParametersResponseOutput {
	return o
}

func (o DiskEncryptionSetParametersResponseOutput) ToDiskEncryptionSetParametersResponseOutputWithContext(ctx context.Context) DiskEncryptionSetParametersResponseOutput {
	return o
}

func (o DiskEncryptionSetParametersResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskEncryptionSetParametersResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type DiskEncryptionSetParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSetParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSetParametersResponse)(nil)).Elem()
}

func (o DiskEncryptionSetParametersResponsePtrOutput) ToDiskEncryptionSetParametersResponsePtrOutput() DiskEncryptionSetParametersResponsePtrOutput {
	return o
}

func (o DiskEncryptionSetParametersResponsePtrOutput) ToDiskEncryptionSetParametersResponsePtrOutputWithContext(ctx context.Context) DiskEncryptionSetParametersResponsePtrOutput {
	return o
}

func (o DiskEncryptionSetParametersResponsePtrOutput) Elem() DiskEncryptionSetParametersResponseOutput {
	return o.ApplyT(func(v *DiskEncryptionSetParametersResponse) DiskEncryptionSetParametersResponse {
		if v != nil {
			return *v
		}
		var ret DiskEncryptionSetParametersResponse
		return ret
	}).(DiskEncryptionSetParametersResponseOutput)
}

func (o DiskEncryptionSetParametersResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSetParametersResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type DiskEncryptionSettings struct {
	DiskEncryptionKey *KeyVaultSecretReference `pulumi:"diskEncryptionKey"`
	Enabled           *bool                    `pulumi:"enabled"`
	KeyEncryptionKey  *KeyVaultKeyReference    `pulumi:"keyEncryptionKey"`
}





type DiskEncryptionSettingsInput interface {
	pulumi.Input

	ToDiskEncryptionSettingsOutput() DiskEncryptionSettingsOutput
	ToDiskEncryptionSettingsOutputWithContext(context.Context) DiskEncryptionSettingsOutput
}

type DiskEncryptionSettingsArgs struct {
	DiskEncryptionKey KeyVaultSecretReferencePtrInput `pulumi:"diskEncryptionKey"`
	Enabled           pulumi.BoolPtrInput             `pulumi:"enabled"`
	KeyEncryptionKey  KeyVaultKeyReferencePtrInput    `pulumi:"keyEncryptionKey"`
}

func (DiskEncryptionSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSettings)(nil)).Elem()
}

func (i DiskEncryptionSettingsArgs) ToDiskEncryptionSettingsOutput() DiskEncryptionSettingsOutput {
	return i.ToDiskEncryptionSettingsOutputWithContext(context.Background())
}

func (i DiskEncryptionSettingsArgs) ToDiskEncryptionSettingsOutputWithContext(ctx context.Context) DiskEncryptionSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSettingsOutput)
}

func (i DiskEncryptionSettingsArgs) ToDiskEncryptionSettingsPtrOutput() DiskEncryptionSettingsPtrOutput {
	return i.ToDiskEncryptionSettingsPtrOutputWithContext(context.Background())
}

func (i DiskEncryptionSettingsArgs) ToDiskEncryptionSettingsPtrOutputWithContext(ctx context.Context) DiskEncryptionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSettingsOutput).ToDiskEncryptionSettingsPtrOutputWithContext(ctx)
}









type DiskEncryptionSettingsPtrInput interface {
	pulumi.Input

	ToDiskEncryptionSettingsPtrOutput() DiskEncryptionSettingsPtrOutput
	ToDiskEncryptionSettingsPtrOutputWithContext(context.Context) DiskEncryptionSettingsPtrOutput
}

type diskEncryptionSettingsPtrType DiskEncryptionSettingsArgs

func DiskEncryptionSettingsPtr(v *DiskEncryptionSettingsArgs) DiskEncryptionSettingsPtrInput {
	return (*diskEncryptionSettingsPtrType)(v)
}

func (*diskEncryptionSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSettings)(nil)).Elem()
}

func (i *diskEncryptionSettingsPtrType) ToDiskEncryptionSettingsPtrOutput() DiskEncryptionSettingsPtrOutput {
	return i.ToDiskEncryptionSettingsPtrOutputWithContext(context.Background())
}

func (i *diskEncryptionSettingsPtrType) ToDiskEncryptionSettingsPtrOutputWithContext(ctx context.Context) DiskEncryptionSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSettingsPtrOutput)
}

type DiskEncryptionSettingsOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSettings)(nil)).Elem()
}

func (o DiskEncryptionSettingsOutput) ToDiskEncryptionSettingsOutput() DiskEncryptionSettingsOutput {
	return o
}

func (o DiskEncryptionSettingsOutput) ToDiskEncryptionSettingsOutputWithContext(ctx context.Context) DiskEncryptionSettingsOutput {
	return o
}

func (o DiskEncryptionSettingsOutput) ToDiskEncryptionSettingsPtrOutput() DiskEncryptionSettingsPtrOutput {
	return o.ToDiskEncryptionSettingsPtrOutputWithContext(context.Background())
}

func (o DiskEncryptionSettingsOutput) ToDiskEncryptionSettingsPtrOutputWithContext(ctx context.Context) DiskEncryptionSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskEncryptionSettings) *DiskEncryptionSettings {
		return &v
	}).(DiskEncryptionSettingsPtrOutput)
}

func (o DiskEncryptionSettingsOutput) DiskEncryptionKey() KeyVaultSecretReferencePtrOutput {
	return o.ApplyT(func(v DiskEncryptionSettings) *KeyVaultSecretReference { return v.DiskEncryptionKey }).(KeyVaultSecretReferencePtrOutput)
}

func (o DiskEncryptionSettingsOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DiskEncryptionSettings) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o DiskEncryptionSettingsOutput) KeyEncryptionKey() KeyVaultKeyReferencePtrOutput {
	return o.ApplyT(func(v DiskEncryptionSettings) *KeyVaultKeyReference { return v.KeyEncryptionKey }).(KeyVaultKeyReferencePtrOutput)
}

type DiskEncryptionSettingsPtrOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSettings)(nil)).Elem()
}

func (o DiskEncryptionSettingsPtrOutput) ToDiskEncryptionSettingsPtrOutput() DiskEncryptionSettingsPtrOutput {
	return o
}

func (o DiskEncryptionSettingsPtrOutput) ToDiskEncryptionSettingsPtrOutputWithContext(ctx context.Context) DiskEncryptionSettingsPtrOutput {
	return o
}

func (o DiskEncryptionSettingsPtrOutput) Elem() DiskEncryptionSettingsOutput {
	return o.ApplyT(func(v *DiskEncryptionSettings) DiskEncryptionSettings {
		if v != nil {
			return *v
		}
		var ret DiskEncryptionSettings
		return ret
	}).(DiskEncryptionSettingsOutput)
}

func (o DiskEncryptionSettingsPtrOutput) DiskEncryptionKey() KeyVaultSecretReferencePtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSettings) *KeyVaultSecretReference {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionKey
	}).(KeyVaultSecretReferencePtrOutput)
}

func (o DiskEncryptionSettingsPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSettings) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o DiskEncryptionSettingsPtrOutput) KeyEncryptionKey() KeyVaultKeyReferencePtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSettings) *KeyVaultKeyReference {
		if v == nil {
			return nil
		}
		return v.KeyEncryptionKey
	}).(KeyVaultKeyReferencePtrOutput)
}

type DiskEncryptionSettingsResponse struct {
	DiskEncryptionKey *KeyVaultSecretReferenceResponse `pulumi:"diskEncryptionKey"`
	Enabled           *bool                            `pulumi:"enabled"`
	KeyEncryptionKey  *KeyVaultKeyReferenceResponse    `pulumi:"keyEncryptionKey"`
}

type DiskEncryptionSettingsResponseOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSettingsResponse)(nil)).Elem()
}

func (o DiskEncryptionSettingsResponseOutput) ToDiskEncryptionSettingsResponseOutput() DiskEncryptionSettingsResponseOutput {
	return o
}

func (o DiskEncryptionSettingsResponseOutput) ToDiskEncryptionSettingsResponseOutputWithContext(ctx context.Context) DiskEncryptionSettingsResponseOutput {
	return o
}

func (o DiskEncryptionSettingsResponseOutput) DiskEncryptionKey() KeyVaultSecretReferenceResponsePtrOutput {
	return o.ApplyT(func(v DiskEncryptionSettingsResponse) *KeyVaultSecretReferenceResponse { return v.DiskEncryptionKey }).(KeyVaultSecretReferenceResponsePtrOutput)
}

func (o DiskEncryptionSettingsResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DiskEncryptionSettingsResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o DiskEncryptionSettingsResponseOutput) KeyEncryptionKey() KeyVaultKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v DiskEncryptionSettingsResponse) *KeyVaultKeyReferenceResponse { return v.KeyEncryptionKey }).(KeyVaultKeyReferenceResponsePtrOutput)
}

type DiskEncryptionSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionSettingsResponse)(nil)).Elem()
}

func (o DiskEncryptionSettingsResponsePtrOutput) ToDiskEncryptionSettingsResponsePtrOutput() DiskEncryptionSettingsResponsePtrOutput {
	return o
}

func (o DiskEncryptionSettingsResponsePtrOutput) ToDiskEncryptionSettingsResponsePtrOutputWithContext(ctx context.Context) DiskEncryptionSettingsResponsePtrOutput {
	return o
}

func (o DiskEncryptionSettingsResponsePtrOutput) Elem() DiskEncryptionSettingsResponseOutput {
	return o.ApplyT(func(v *DiskEncryptionSettingsResponse) DiskEncryptionSettingsResponse {
		if v != nil {
			return *v
		}
		var ret DiskEncryptionSettingsResponse
		return ret
	}).(DiskEncryptionSettingsResponseOutput)
}

func (o DiskEncryptionSettingsResponsePtrOutput) DiskEncryptionKey() KeyVaultSecretReferenceResponsePtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSettingsResponse) *KeyVaultSecretReferenceResponse {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionKey
	}).(KeyVaultSecretReferenceResponsePtrOutput)
}

func (o DiskEncryptionSettingsResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o DiskEncryptionSettingsResponsePtrOutput) KeyEncryptionKey() KeyVaultKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v *DiskEncryptionSettingsResponse) *KeyVaultKeyReferenceResponse {
		if v == nil {
			return nil
		}
		return v.KeyEncryptionKey
	}).(KeyVaultKeyReferenceResponsePtrOutput)
}

type DiskEncryptionSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskEncryptionSettingsResponse)(nil)).Elem()
}

func (o DiskEncryptionSettingsResponseArrayOutput) ToDiskEncryptionSettingsResponseArrayOutput() DiskEncryptionSettingsResponseArrayOutput {
	return o
}

func (o DiskEncryptionSettingsResponseArrayOutput) ToDiskEncryptionSettingsResponseArrayOutputWithContext(ctx context.Context) DiskEncryptionSettingsResponseArrayOutput {
	return o
}

func (o DiskEncryptionSettingsResponseArrayOutput) Index(i pulumi.IntInput) DiskEncryptionSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiskEncryptionSettingsResponse {
		return vs[0].([]DiskEncryptionSettingsResponse)[vs[1].(int)]
	}).(DiskEncryptionSettingsResponseOutput)
}

type DiskInstanceViewResponse struct {
	EncryptionSettings []DiskEncryptionSettingsResponse `pulumi:"encryptionSettings"`
	Name               *string                          `pulumi:"name"`
	Statuses           []InstanceViewStatusResponse     `pulumi:"statuses"`
}

type DiskInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (DiskInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskInstanceViewResponse)(nil)).Elem()
}

func (o DiskInstanceViewResponseOutput) ToDiskInstanceViewResponseOutput() DiskInstanceViewResponseOutput {
	return o
}

func (o DiskInstanceViewResponseOutput) ToDiskInstanceViewResponseOutputWithContext(ctx context.Context) DiskInstanceViewResponseOutput {
	return o
}

func (o DiskInstanceViewResponseOutput) EncryptionSettings() DiskEncryptionSettingsResponseArrayOutput {
	return o.ApplyT(func(v DiskInstanceViewResponse) []DiskEncryptionSettingsResponse { return v.EncryptionSettings }).(DiskEncryptionSettingsResponseArrayOutput)
}

func (o DiskInstanceViewResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskInstanceViewResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DiskInstanceViewResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v DiskInstanceViewResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

type DiskInstanceViewResponseArrayOutput struct{ *pulumi.OutputState }

func (DiskInstanceViewResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskInstanceViewResponse)(nil)).Elem()
}

func (o DiskInstanceViewResponseArrayOutput) ToDiskInstanceViewResponseArrayOutput() DiskInstanceViewResponseArrayOutput {
	return o
}

func (o DiskInstanceViewResponseArrayOutput) ToDiskInstanceViewResponseArrayOutputWithContext(ctx context.Context) DiskInstanceViewResponseArrayOutput {
	return o
}

func (o DiskInstanceViewResponseArrayOutput) Index(i pulumi.IntInput) DiskInstanceViewResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiskInstanceViewResponse {
		return vs[0].([]DiskInstanceViewResponse)[vs[1].(int)]
	}).(DiskInstanceViewResponseOutput)
}

type DiskRestorePointInstanceViewResponse struct {
	Id                *string                                    `pulumi:"id"`
	ReplicationStatus *DiskRestorePointReplicationStatusResponse `pulumi:"replicationStatus"`
}

type DiskRestorePointInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (DiskRestorePointInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskRestorePointInstanceViewResponse)(nil)).Elem()
}

func (o DiskRestorePointInstanceViewResponseOutput) ToDiskRestorePointInstanceViewResponseOutput() DiskRestorePointInstanceViewResponseOutput {
	return o
}

func (o DiskRestorePointInstanceViewResponseOutput) ToDiskRestorePointInstanceViewResponseOutputWithContext(ctx context.Context) DiskRestorePointInstanceViewResponseOutput {
	return o
}

func (o DiskRestorePointInstanceViewResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskRestorePointInstanceViewResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o DiskRestorePointInstanceViewResponseOutput) ReplicationStatus() DiskRestorePointReplicationStatusResponsePtrOutput {
	return o.ApplyT(func(v DiskRestorePointInstanceViewResponse) *DiskRestorePointReplicationStatusResponse {
		return v.ReplicationStatus
	}).(DiskRestorePointReplicationStatusResponsePtrOutput)
}

type DiskRestorePointInstanceViewResponseArrayOutput struct{ *pulumi.OutputState }

func (DiskRestorePointInstanceViewResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskRestorePointInstanceViewResponse)(nil)).Elem()
}

func (o DiskRestorePointInstanceViewResponseArrayOutput) ToDiskRestorePointInstanceViewResponseArrayOutput() DiskRestorePointInstanceViewResponseArrayOutput {
	return o
}

func (o DiskRestorePointInstanceViewResponseArrayOutput) ToDiskRestorePointInstanceViewResponseArrayOutputWithContext(ctx context.Context) DiskRestorePointInstanceViewResponseArrayOutput {
	return o
}

func (o DiskRestorePointInstanceViewResponseArrayOutput) Index(i pulumi.IntInput) DiskRestorePointInstanceViewResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiskRestorePointInstanceViewResponse {
		return vs[0].([]DiskRestorePointInstanceViewResponse)[vs[1].(int)]
	}).(DiskRestorePointInstanceViewResponseOutput)
}

type DiskRestorePointReplicationStatusResponse struct {
	CompletionPercent *int                        `pulumi:"completionPercent"`
	Status            *InstanceViewStatusResponse `pulumi:"status"`
}

type DiskRestorePointReplicationStatusResponseOutput struct{ *pulumi.OutputState }

func (DiskRestorePointReplicationStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskRestorePointReplicationStatusResponse)(nil)).Elem()
}

func (o DiskRestorePointReplicationStatusResponseOutput) ToDiskRestorePointReplicationStatusResponseOutput() DiskRestorePointReplicationStatusResponseOutput {
	return o
}

func (o DiskRestorePointReplicationStatusResponseOutput) ToDiskRestorePointReplicationStatusResponseOutputWithContext(ctx context.Context) DiskRestorePointReplicationStatusResponseOutput {
	return o
}

func (o DiskRestorePointReplicationStatusResponseOutput) CompletionPercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DiskRestorePointReplicationStatusResponse) *int { return v.CompletionPercent }).(pulumi.IntPtrOutput)
}

func (o DiskRestorePointReplicationStatusResponseOutput) Status() InstanceViewStatusResponsePtrOutput {
	return o.ApplyT(func(v DiskRestorePointReplicationStatusResponse) *InstanceViewStatusResponse { return v.Status }).(InstanceViewStatusResponsePtrOutput)
}

type DiskRestorePointReplicationStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (DiskRestorePointReplicationStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskRestorePointReplicationStatusResponse)(nil)).Elem()
}

func (o DiskRestorePointReplicationStatusResponsePtrOutput) ToDiskRestorePointReplicationStatusResponsePtrOutput() DiskRestorePointReplicationStatusResponsePtrOutput {
	return o
}

func (o DiskRestorePointReplicationStatusResponsePtrOutput) ToDiskRestorePointReplicationStatusResponsePtrOutputWithContext(ctx context.Context) DiskRestorePointReplicationStatusResponsePtrOutput {
	return o
}

func (o DiskRestorePointReplicationStatusResponsePtrOutput) Elem() DiskRestorePointReplicationStatusResponseOutput {
	return o.ApplyT(func(v *DiskRestorePointReplicationStatusResponse) DiskRestorePointReplicationStatusResponse {
		if v != nil {
			return *v
		}
		var ret DiskRestorePointReplicationStatusResponse
		return ret
	}).(DiskRestorePointReplicationStatusResponseOutput)
}

func (o DiskRestorePointReplicationStatusResponsePtrOutput) CompletionPercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DiskRestorePointReplicationStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.CompletionPercent
	}).(pulumi.IntPtrOutput)
}

func (o DiskRestorePointReplicationStatusResponsePtrOutput) Status() InstanceViewStatusResponsePtrOutput {
	return o.ApplyT(func(v *DiskRestorePointReplicationStatusResponse) *InstanceViewStatusResponse {
		if v == nil {
			return nil
		}
		return v.Status
	}).(InstanceViewStatusResponsePtrOutput)
}

type ExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ExtendedLocationInput interface {
	pulumi.Input

	ToExtendedLocationOutput() ExtendedLocationOutput
	ToExtendedLocationOutputWithContext(context.Context) ExtendedLocationOutput
}

type ExtendedLocationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
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

func (o ExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
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

func (o ExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
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

func (o ExtendedLocationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
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
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type HardwareProfile struct {
	VmSize           *string           `pulumi:"vmSize"`
	VmSizeProperties *VMSizeProperties `pulumi:"vmSizeProperties"`
}





type HardwareProfileInput interface {
	pulumi.Input

	ToHardwareProfileOutput() HardwareProfileOutput
	ToHardwareProfileOutputWithContext(context.Context) HardwareProfileOutput
}

type HardwareProfileArgs struct {
	VmSize           pulumi.StringPtrInput    `pulumi:"vmSize"`
	VmSizeProperties VMSizePropertiesPtrInput `pulumi:"vmSizeProperties"`
}

func (HardwareProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfile)(nil)).Elem()
}

func (i HardwareProfileArgs) ToHardwareProfileOutput() HardwareProfileOutput {
	return i.ToHardwareProfileOutputWithContext(context.Background())
}

func (i HardwareProfileArgs) ToHardwareProfileOutputWithContext(ctx context.Context) HardwareProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardwareProfileOutput)
}

func (i HardwareProfileArgs) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return i.ToHardwareProfilePtrOutputWithContext(context.Background())
}

func (i HardwareProfileArgs) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardwareProfileOutput).ToHardwareProfilePtrOutputWithContext(ctx)
}









type HardwareProfilePtrInput interface {
	pulumi.Input

	ToHardwareProfilePtrOutput() HardwareProfilePtrOutput
	ToHardwareProfilePtrOutputWithContext(context.Context) HardwareProfilePtrOutput
}

type hardwareProfilePtrType HardwareProfileArgs

func HardwareProfilePtr(v *HardwareProfileArgs) HardwareProfilePtrInput {
	return (*hardwareProfilePtrType)(v)
}

func (*hardwareProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfile)(nil)).Elem()
}

func (i *hardwareProfilePtrType) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return i.ToHardwareProfilePtrOutputWithContext(context.Background())
}

func (i *hardwareProfilePtrType) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardwareProfilePtrOutput)
}

type HardwareProfileOutput struct{ *pulumi.OutputState }

func (HardwareProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfile)(nil)).Elem()
}

func (o HardwareProfileOutput) ToHardwareProfileOutput() HardwareProfileOutput {
	return o
}

func (o HardwareProfileOutput) ToHardwareProfileOutputWithContext(ctx context.Context) HardwareProfileOutput {
	return o
}

func (o HardwareProfileOutput) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return o.ToHardwareProfilePtrOutputWithContext(context.Background())
}

func (o HardwareProfileOutput) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HardwareProfile) *HardwareProfile {
		return &v
	}).(HardwareProfilePtrOutput)
}

func (o HardwareProfileOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HardwareProfile) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

func (o HardwareProfileOutput) VmSizeProperties() VMSizePropertiesPtrOutput {
	return o.ApplyT(func(v HardwareProfile) *VMSizeProperties { return v.VmSizeProperties }).(VMSizePropertiesPtrOutput)
}

type HardwareProfilePtrOutput struct{ *pulumi.OutputState }

func (HardwareProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfile)(nil)).Elem()
}

func (o HardwareProfilePtrOutput) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return o
}

func (o HardwareProfilePtrOutput) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return o
}

func (o HardwareProfilePtrOutput) Elem() HardwareProfileOutput {
	return o.ApplyT(func(v *HardwareProfile) HardwareProfile {
		if v != nil {
			return *v
		}
		var ret HardwareProfile
		return ret
	}).(HardwareProfileOutput)
}

func (o HardwareProfilePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardwareProfile) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

func (o HardwareProfilePtrOutput) VmSizeProperties() VMSizePropertiesPtrOutput {
	return o.ApplyT(func(v *HardwareProfile) *VMSizeProperties {
		if v == nil {
			return nil
		}
		return v.VmSizeProperties
	}).(VMSizePropertiesPtrOutput)
}

type HardwareProfileResponse struct {
	VmSize           *string                   `pulumi:"vmSize"`
	VmSizeProperties *VMSizePropertiesResponse `pulumi:"vmSizeProperties"`
}

type HardwareProfileResponseOutput struct{ *pulumi.OutputState }

func (HardwareProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfileResponse)(nil)).Elem()
}

func (o HardwareProfileResponseOutput) ToHardwareProfileResponseOutput() HardwareProfileResponseOutput {
	return o
}

func (o HardwareProfileResponseOutput) ToHardwareProfileResponseOutputWithContext(ctx context.Context) HardwareProfileResponseOutput {
	return o
}

func (o HardwareProfileResponseOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HardwareProfileResponse) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

func (o HardwareProfileResponseOutput) VmSizeProperties() VMSizePropertiesResponsePtrOutput {
	return o.ApplyT(func(v HardwareProfileResponse) *VMSizePropertiesResponse { return v.VmSizeProperties }).(VMSizePropertiesResponsePtrOutput)
}

type HardwareProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (HardwareProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfileResponse)(nil)).Elem()
}

func (o HardwareProfileResponsePtrOutput) ToHardwareProfileResponsePtrOutput() HardwareProfileResponsePtrOutput {
	return o
}

func (o HardwareProfileResponsePtrOutput) ToHardwareProfileResponsePtrOutputWithContext(ctx context.Context) HardwareProfileResponsePtrOutput {
	return o
}

func (o HardwareProfileResponsePtrOutput) Elem() HardwareProfileResponseOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) HardwareProfileResponse {
		if v != nil {
			return *v
		}
		var ret HardwareProfileResponse
		return ret
	}).(HardwareProfileResponseOutput)
}

func (o HardwareProfileResponsePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

func (o HardwareProfileResponsePtrOutput) VmSizeProperties() VMSizePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *VMSizePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.VmSizeProperties
	}).(VMSizePropertiesResponsePtrOutput)
}

type ImageDataDisk struct {
	BlobUri            *string                      `pulumi:"blobUri"`
	Caching            *CachingTypes                `pulumi:"caching"`
	DiskEncryptionSet  *DiskEncryptionSetParameters `pulumi:"diskEncryptionSet"`
	DiskSizeGB         *int                         `pulumi:"diskSizeGB"`
	Lun                int                          `pulumi:"lun"`
	ManagedDisk        *SubResource                 `pulumi:"managedDisk"`
	Snapshot           *SubResource                 `pulumi:"snapshot"`
	StorageAccountType *string                      `pulumi:"storageAccountType"`
}





type ImageDataDiskInput interface {
	pulumi.Input

	ToImageDataDiskOutput() ImageDataDiskOutput
	ToImageDataDiskOutputWithContext(context.Context) ImageDataDiskOutput
}

type ImageDataDiskArgs struct {
	BlobUri            pulumi.StringPtrInput               `pulumi:"blobUri"`
	Caching            CachingTypesPtrInput                `pulumi:"caching"`
	DiskEncryptionSet  DiskEncryptionSetParametersPtrInput `pulumi:"diskEncryptionSet"`
	DiskSizeGB         pulumi.IntPtrInput                  `pulumi:"diskSizeGB"`
	Lun                pulumi.IntInput                     `pulumi:"lun"`
	ManagedDisk        SubResourcePtrInput                 `pulumi:"managedDisk"`
	Snapshot           SubResourcePtrInput                 `pulumi:"snapshot"`
	StorageAccountType pulumi.StringPtrInput               `pulumi:"storageAccountType"`
}

func (ImageDataDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageDataDisk)(nil)).Elem()
}

func (i ImageDataDiskArgs) ToImageDataDiskOutput() ImageDataDiskOutput {
	return i.ToImageDataDiskOutputWithContext(context.Background())
}

func (i ImageDataDiskArgs) ToImageDataDiskOutputWithContext(ctx context.Context) ImageDataDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageDataDiskOutput)
}





type ImageDataDiskArrayInput interface {
	pulumi.Input

	ToImageDataDiskArrayOutput() ImageDataDiskArrayOutput
	ToImageDataDiskArrayOutputWithContext(context.Context) ImageDataDiskArrayOutput
}

type ImageDataDiskArray []ImageDataDiskInput

func (ImageDataDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageDataDisk)(nil)).Elem()
}

func (i ImageDataDiskArray) ToImageDataDiskArrayOutput() ImageDataDiskArrayOutput {
	return i.ToImageDataDiskArrayOutputWithContext(context.Background())
}

func (i ImageDataDiskArray) ToImageDataDiskArrayOutputWithContext(ctx context.Context) ImageDataDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageDataDiskArrayOutput)
}

type ImageDataDiskOutput struct{ *pulumi.OutputState }

func (ImageDataDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageDataDisk)(nil)).Elem()
}

func (o ImageDataDiskOutput) ToImageDataDiskOutput() ImageDataDiskOutput {
	return o
}

func (o ImageDataDiskOutput) ToImageDataDiskOutputWithContext(ctx context.Context) ImageDataDiskOutput {
	return o
}

func (o ImageDataDiskOutput) BlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageDataDisk) *string { return v.BlobUri }).(pulumi.StringPtrOutput)
}

func (o ImageDataDiskOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v ImageDataDisk) *CachingTypes { return v.Caching }).(CachingTypesPtrOutput)
}

func (o ImageDataDiskOutput) DiskEncryptionSet() DiskEncryptionSetParametersPtrOutput {
	return o.ApplyT(func(v ImageDataDisk) *DiskEncryptionSetParameters { return v.DiskEncryptionSet }).(DiskEncryptionSetParametersPtrOutput)
}

func (o ImageDataDiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImageDataDisk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ImageDataDiskOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v ImageDataDisk) int { return v.Lun }).(pulumi.IntOutput)
}

func (o ImageDataDiskOutput) ManagedDisk() SubResourcePtrOutput {
	return o.ApplyT(func(v ImageDataDisk) *SubResource { return v.ManagedDisk }).(SubResourcePtrOutput)
}

func (o ImageDataDiskOutput) Snapshot() SubResourcePtrOutput {
	return o.ApplyT(func(v ImageDataDisk) *SubResource { return v.Snapshot }).(SubResourcePtrOutput)
}

func (o ImageDataDiskOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageDataDisk) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type ImageDataDiskArrayOutput struct{ *pulumi.OutputState }

func (ImageDataDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageDataDisk)(nil)).Elem()
}

func (o ImageDataDiskArrayOutput) ToImageDataDiskArrayOutput() ImageDataDiskArrayOutput {
	return o
}

func (o ImageDataDiskArrayOutput) ToImageDataDiskArrayOutputWithContext(ctx context.Context) ImageDataDiskArrayOutput {
	return o
}

func (o ImageDataDiskArrayOutput) Index(i pulumi.IntInput) ImageDataDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ImageDataDisk {
		return vs[0].([]ImageDataDisk)[vs[1].(int)]
	}).(ImageDataDiskOutput)
}

type ImageDataDiskResponse struct {
	BlobUri            *string                              `pulumi:"blobUri"`
	Caching            *string                              `pulumi:"caching"`
	DiskEncryptionSet  *DiskEncryptionSetParametersResponse `pulumi:"diskEncryptionSet"`
	DiskSizeGB         *int                                 `pulumi:"diskSizeGB"`
	Lun                int                                  `pulumi:"lun"`
	ManagedDisk        *SubResourceResponse                 `pulumi:"managedDisk"`
	Snapshot           *SubResourceResponse                 `pulumi:"snapshot"`
	StorageAccountType *string                              `pulumi:"storageAccountType"`
}

type ImageDataDiskResponseOutput struct{ *pulumi.OutputState }

func (ImageDataDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageDataDiskResponse)(nil)).Elem()
}

func (o ImageDataDiskResponseOutput) ToImageDataDiskResponseOutput() ImageDataDiskResponseOutput {
	return o
}

func (o ImageDataDiskResponseOutput) ToImageDataDiskResponseOutputWithContext(ctx context.Context) ImageDataDiskResponseOutput {
	return o
}

func (o ImageDataDiskResponseOutput) BlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageDataDiskResponse) *string { return v.BlobUri }).(pulumi.StringPtrOutput)
}

func (o ImageDataDiskResponseOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageDataDiskResponse) *string { return v.Caching }).(pulumi.StringPtrOutput)
}

func (o ImageDataDiskResponseOutput) DiskEncryptionSet() DiskEncryptionSetParametersResponsePtrOutput {
	return o.ApplyT(func(v ImageDataDiskResponse) *DiskEncryptionSetParametersResponse { return v.DiskEncryptionSet }).(DiskEncryptionSetParametersResponsePtrOutput)
}

func (o ImageDataDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImageDataDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ImageDataDiskResponseOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v ImageDataDiskResponse) int { return v.Lun }).(pulumi.IntOutput)
}

func (o ImageDataDiskResponseOutput) ManagedDisk() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ImageDataDiskResponse) *SubResourceResponse { return v.ManagedDisk }).(SubResourceResponsePtrOutput)
}

func (o ImageDataDiskResponseOutput) Snapshot() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ImageDataDiskResponse) *SubResourceResponse { return v.Snapshot }).(SubResourceResponsePtrOutput)
}

func (o ImageDataDiskResponseOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageDataDiskResponse) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type ImageDataDiskResponseArrayOutput struct{ *pulumi.OutputState }

func (ImageDataDiskResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageDataDiskResponse)(nil)).Elem()
}

func (o ImageDataDiskResponseArrayOutput) ToImageDataDiskResponseArrayOutput() ImageDataDiskResponseArrayOutput {
	return o
}

func (o ImageDataDiskResponseArrayOutput) ToImageDataDiskResponseArrayOutputWithContext(ctx context.Context) ImageDataDiskResponseArrayOutput {
	return o
}

func (o ImageDataDiskResponseArrayOutput) Index(i pulumi.IntInput) ImageDataDiskResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ImageDataDiskResponse {
		return vs[0].([]ImageDataDiskResponse)[vs[1].(int)]
	}).(ImageDataDiskResponseOutput)
}

type ImageOSDisk struct {
	BlobUri            *string                      `pulumi:"blobUri"`
	Caching            *CachingTypes                `pulumi:"caching"`
	DiskEncryptionSet  *DiskEncryptionSetParameters `pulumi:"diskEncryptionSet"`
	DiskSizeGB         *int                         `pulumi:"diskSizeGB"`
	ManagedDisk        *SubResource                 `pulumi:"managedDisk"`
	OsState            OperatingSystemStateTypes    `pulumi:"osState"`
	OsType             OperatingSystemTypes         `pulumi:"osType"`
	Snapshot           *SubResource                 `pulumi:"snapshot"`
	StorageAccountType *string                      `pulumi:"storageAccountType"`
}





type ImageOSDiskInput interface {
	pulumi.Input

	ToImageOSDiskOutput() ImageOSDiskOutput
	ToImageOSDiskOutputWithContext(context.Context) ImageOSDiskOutput
}

type ImageOSDiskArgs struct {
	BlobUri            pulumi.StringPtrInput               `pulumi:"blobUri"`
	Caching            CachingTypesPtrInput                `pulumi:"caching"`
	DiskEncryptionSet  DiskEncryptionSetParametersPtrInput `pulumi:"diskEncryptionSet"`
	DiskSizeGB         pulumi.IntPtrInput                  `pulumi:"diskSizeGB"`
	ManagedDisk        SubResourcePtrInput                 `pulumi:"managedDisk"`
	OsState            OperatingSystemStateTypesInput      `pulumi:"osState"`
	OsType             OperatingSystemTypesInput           `pulumi:"osType"`
	Snapshot           SubResourcePtrInput                 `pulumi:"snapshot"`
	StorageAccountType pulumi.StringPtrInput               `pulumi:"storageAccountType"`
}

func (ImageOSDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageOSDisk)(nil)).Elem()
}

func (i ImageOSDiskArgs) ToImageOSDiskOutput() ImageOSDiskOutput {
	return i.ToImageOSDiskOutputWithContext(context.Background())
}

func (i ImageOSDiskArgs) ToImageOSDiskOutputWithContext(ctx context.Context) ImageOSDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOSDiskOutput)
}

func (i ImageOSDiskArgs) ToImageOSDiskPtrOutput() ImageOSDiskPtrOutput {
	return i.ToImageOSDiskPtrOutputWithContext(context.Background())
}

func (i ImageOSDiskArgs) ToImageOSDiskPtrOutputWithContext(ctx context.Context) ImageOSDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOSDiskOutput).ToImageOSDiskPtrOutputWithContext(ctx)
}









type ImageOSDiskPtrInput interface {
	pulumi.Input

	ToImageOSDiskPtrOutput() ImageOSDiskPtrOutput
	ToImageOSDiskPtrOutputWithContext(context.Context) ImageOSDiskPtrOutput
}

type imageOSDiskPtrType ImageOSDiskArgs

func ImageOSDiskPtr(v *ImageOSDiskArgs) ImageOSDiskPtrInput {
	return (*imageOSDiskPtrType)(v)
}

func (*imageOSDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageOSDisk)(nil)).Elem()
}

func (i *imageOSDiskPtrType) ToImageOSDiskPtrOutput() ImageOSDiskPtrOutput {
	return i.ToImageOSDiskPtrOutputWithContext(context.Background())
}

func (i *imageOSDiskPtrType) ToImageOSDiskPtrOutputWithContext(ctx context.Context) ImageOSDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOSDiskPtrOutput)
}

type ImageOSDiskOutput struct{ *pulumi.OutputState }

func (ImageOSDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageOSDisk)(nil)).Elem()
}

func (o ImageOSDiskOutput) ToImageOSDiskOutput() ImageOSDiskOutput {
	return o
}

func (o ImageOSDiskOutput) ToImageOSDiskOutputWithContext(ctx context.Context) ImageOSDiskOutput {
	return o
}

func (o ImageOSDiskOutput) ToImageOSDiskPtrOutput() ImageOSDiskPtrOutput {
	return o.ToImageOSDiskPtrOutputWithContext(context.Background())
}

func (o ImageOSDiskOutput) ToImageOSDiskPtrOutputWithContext(ctx context.Context) ImageOSDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageOSDisk) *ImageOSDisk {
		return &v
	}).(ImageOSDiskPtrOutput)
}

func (o ImageOSDiskOutput) BlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageOSDisk) *string { return v.BlobUri }).(pulumi.StringPtrOutput)
}

func (o ImageOSDiskOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v ImageOSDisk) *CachingTypes { return v.Caching }).(CachingTypesPtrOutput)
}

func (o ImageOSDiskOutput) DiskEncryptionSet() DiskEncryptionSetParametersPtrOutput {
	return o.ApplyT(func(v ImageOSDisk) *DiskEncryptionSetParameters { return v.DiskEncryptionSet }).(DiskEncryptionSetParametersPtrOutput)
}

func (o ImageOSDiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImageOSDisk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ImageOSDiskOutput) ManagedDisk() SubResourcePtrOutput {
	return o.ApplyT(func(v ImageOSDisk) *SubResource { return v.ManagedDisk }).(SubResourcePtrOutput)
}

func (o ImageOSDiskOutput) OsState() OperatingSystemStateTypesOutput {
	return o.ApplyT(func(v ImageOSDisk) OperatingSystemStateTypes { return v.OsState }).(OperatingSystemStateTypesOutput)
}

func (o ImageOSDiskOutput) OsType() OperatingSystemTypesOutput {
	return o.ApplyT(func(v ImageOSDisk) OperatingSystemTypes { return v.OsType }).(OperatingSystemTypesOutput)
}

func (o ImageOSDiskOutput) Snapshot() SubResourcePtrOutput {
	return o.ApplyT(func(v ImageOSDisk) *SubResource { return v.Snapshot }).(SubResourcePtrOutput)
}

func (o ImageOSDiskOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageOSDisk) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type ImageOSDiskPtrOutput struct{ *pulumi.OutputState }

func (ImageOSDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageOSDisk)(nil)).Elem()
}

func (o ImageOSDiskPtrOutput) ToImageOSDiskPtrOutput() ImageOSDiskPtrOutput {
	return o
}

func (o ImageOSDiskPtrOutput) ToImageOSDiskPtrOutputWithContext(ctx context.Context) ImageOSDiskPtrOutput {
	return o
}

func (o ImageOSDiskPtrOutput) Elem() ImageOSDiskOutput {
	return o.ApplyT(func(v *ImageOSDisk) ImageOSDisk {
		if v != nil {
			return *v
		}
		var ret ImageOSDisk
		return ret
	}).(ImageOSDiskOutput)
}

func (o ImageOSDiskPtrOutput) BlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageOSDisk) *string {
		if v == nil {
			return nil
		}
		return v.BlobUri
	}).(pulumi.StringPtrOutput)
}

func (o ImageOSDiskPtrOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v *ImageOSDisk) *CachingTypes {
		if v == nil {
			return nil
		}
		return v.Caching
	}).(CachingTypesPtrOutput)
}

func (o ImageOSDiskPtrOutput) DiskEncryptionSet() DiskEncryptionSetParametersPtrOutput {
	return o.ApplyT(func(v *ImageOSDisk) *DiskEncryptionSetParameters {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionSet
	}).(DiskEncryptionSetParametersPtrOutput)
}

func (o ImageOSDiskPtrOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImageOSDisk) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o ImageOSDiskPtrOutput) ManagedDisk() SubResourcePtrOutput {
	return o.ApplyT(func(v *ImageOSDisk) *SubResource {
		if v == nil {
			return nil
		}
		return v.ManagedDisk
	}).(SubResourcePtrOutput)
}

func (o ImageOSDiskPtrOutput) OsState() OperatingSystemStateTypesPtrOutput {
	return o.ApplyT(func(v *ImageOSDisk) *OperatingSystemStateTypes {
		if v == nil {
			return nil
		}
		return &v.OsState
	}).(OperatingSystemStateTypesPtrOutput)
}

func (o ImageOSDiskPtrOutput) OsType() OperatingSystemTypesPtrOutput {
	return o.ApplyT(func(v *ImageOSDisk) *OperatingSystemTypes {
		if v == nil {
			return nil
		}
		return &v.OsType
	}).(OperatingSystemTypesPtrOutput)
}

func (o ImageOSDiskPtrOutput) Snapshot() SubResourcePtrOutput {
	return o.ApplyT(func(v *ImageOSDisk) *SubResource {
		if v == nil {
			return nil
		}
		return v.Snapshot
	}).(SubResourcePtrOutput)
}

func (o ImageOSDiskPtrOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageOSDisk) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountType
	}).(pulumi.StringPtrOutput)
}

type ImageOSDiskResponse struct {
	BlobUri            *string                              `pulumi:"blobUri"`
	Caching            *string                              `pulumi:"caching"`
	DiskEncryptionSet  *DiskEncryptionSetParametersResponse `pulumi:"diskEncryptionSet"`
	DiskSizeGB         *int                                 `pulumi:"diskSizeGB"`
	ManagedDisk        *SubResourceResponse                 `pulumi:"managedDisk"`
	OsState            string                               `pulumi:"osState"`
	OsType             string                               `pulumi:"osType"`
	Snapshot           *SubResourceResponse                 `pulumi:"snapshot"`
	StorageAccountType *string                              `pulumi:"storageAccountType"`
}

type ImageOSDiskResponseOutput struct{ *pulumi.OutputState }

func (ImageOSDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageOSDiskResponse)(nil)).Elem()
}

func (o ImageOSDiskResponseOutput) ToImageOSDiskResponseOutput() ImageOSDiskResponseOutput {
	return o
}

func (o ImageOSDiskResponseOutput) ToImageOSDiskResponseOutputWithContext(ctx context.Context) ImageOSDiskResponseOutput {
	return o
}

func (o ImageOSDiskResponseOutput) BlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageOSDiskResponse) *string { return v.BlobUri }).(pulumi.StringPtrOutput)
}

func (o ImageOSDiskResponseOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageOSDiskResponse) *string { return v.Caching }).(pulumi.StringPtrOutput)
}

func (o ImageOSDiskResponseOutput) DiskEncryptionSet() DiskEncryptionSetParametersResponsePtrOutput {
	return o.ApplyT(func(v ImageOSDiskResponse) *DiskEncryptionSetParametersResponse { return v.DiskEncryptionSet }).(DiskEncryptionSetParametersResponsePtrOutput)
}

func (o ImageOSDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImageOSDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ImageOSDiskResponseOutput) ManagedDisk() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ImageOSDiskResponse) *SubResourceResponse { return v.ManagedDisk }).(SubResourceResponsePtrOutput)
}

func (o ImageOSDiskResponseOutput) OsState() pulumi.StringOutput {
	return o.ApplyT(func(v ImageOSDiskResponse) string { return v.OsState }).(pulumi.StringOutput)
}

func (o ImageOSDiskResponseOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v ImageOSDiskResponse) string { return v.OsType }).(pulumi.StringOutput)
}

func (o ImageOSDiskResponseOutput) Snapshot() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ImageOSDiskResponse) *SubResourceResponse { return v.Snapshot }).(SubResourceResponsePtrOutput)
}

func (o ImageOSDiskResponseOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageOSDiskResponse) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type ImageOSDiskResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageOSDiskResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageOSDiskResponse)(nil)).Elem()
}

func (o ImageOSDiskResponsePtrOutput) ToImageOSDiskResponsePtrOutput() ImageOSDiskResponsePtrOutput {
	return o
}

func (o ImageOSDiskResponsePtrOutput) ToImageOSDiskResponsePtrOutputWithContext(ctx context.Context) ImageOSDiskResponsePtrOutput {
	return o
}

func (o ImageOSDiskResponsePtrOutput) Elem() ImageOSDiskResponseOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) ImageOSDiskResponse {
		if v != nil {
			return *v
		}
		var ret ImageOSDiskResponse
		return ret
	}).(ImageOSDiskResponseOutput)
}

func (o ImageOSDiskResponsePtrOutput) BlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.BlobUri
	}).(pulumi.StringPtrOutput)
}

func (o ImageOSDiskResponsePtrOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.Caching
	}).(pulumi.StringPtrOutput)
}

func (o ImageOSDiskResponsePtrOutput) DiskEncryptionSet() DiskEncryptionSetParametersResponsePtrOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) *DiskEncryptionSetParametersResponse {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionSet
	}).(DiskEncryptionSetParametersResponsePtrOutput)
}

func (o ImageOSDiskResponsePtrOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o ImageOSDiskResponsePtrOutput) ManagedDisk() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return v.ManagedDisk
	}).(SubResourceResponsePtrOutput)
}

func (o ImageOSDiskResponsePtrOutput) OsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsState
	}).(pulumi.StringPtrOutput)
}

func (o ImageOSDiskResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o ImageOSDiskResponsePtrOutput) Snapshot() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return v.Snapshot
	}).(SubResourceResponsePtrOutput)
}

func (o ImageOSDiskResponsePtrOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountType
	}).(pulumi.StringPtrOutput)
}

type ImageReference struct {
	CommunityGalleryImageId *string `pulumi:"communityGalleryImageId"`
	Id                      *string `pulumi:"id"`
	Offer                   *string `pulumi:"offer"`
	Publisher               *string `pulumi:"publisher"`
	SharedGalleryImageId    *string `pulumi:"sharedGalleryImageId"`
	Sku                     *string `pulumi:"sku"`
	Version                 *string `pulumi:"version"`
}





type ImageReferenceInput interface {
	pulumi.Input

	ToImageReferenceOutput() ImageReferenceOutput
	ToImageReferenceOutputWithContext(context.Context) ImageReferenceOutput
}

type ImageReferenceArgs struct {
	CommunityGalleryImageId pulumi.StringPtrInput `pulumi:"communityGalleryImageId"`
	Id                      pulumi.StringPtrInput `pulumi:"id"`
	Offer                   pulumi.StringPtrInput `pulumi:"offer"`
	Publisher               pulumi.StringPtrInput `pulumi:"publisher"`
	SharedGalleryImageId    pulumi.StringPtrInput `pulumi:"sharedGalleryImageId"`
	Sku                     pulumi.StringPtrInput `pulumi:"sku"`
	Version                 pulumi.StringPtrInput `pulumi:"version"`
}

func (ImageReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReference)(nil)).Elem()
}

func (i ImageReferenceArgs) ToImageReferenceOutput() ImageReferenceOutput {
	return i.ToImageReferenceOutputWithContext(context.Background())
}

func (i ImageReferenceArgs) ToImageReferenceOutputWithContext(ctx context.Context) ImageReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferenceOutput)
}

func (i ImageReferenceArgs) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return i.ToImageReferencePtrOutputWithContext(context.Background())
}

func (i ImageReferenceArgs) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferenceOutput).ToImageReferencePtrOutputWithContext(ctx)
}









type ImageReferencePtrInput interface {
	pulumi.Input

	ToImageReferencePtrOutput() ImageReferencePtrOutput
	ToImageReferencePtrOutputWithContext(context.Context) ImageReferencePtrOutput
}

type imageReferencePtrType ImageReferenceArgs

func ImageReferencePtr(v *ImageReferenceArgs) ImageReferencePtrInput {
	return (*imageReferencePtrType)(v)
}

func (*imageReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageReference)(nil)).Elem()
}

func (i *imageReferencePtrType) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return i.ToImageReferencePtrOutputWithContext(context.Background())
}

func (i *imageReferencePtrType) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferencePtrOutput)
}

type ImageReferenceOutput struct{ *pulumi.OutputState }

func (ImageReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReference)(nil)).Elem()
}

func (o ImageReferenceOutput) ToImageReferenceOutput() ImageReferenceOutput {
	return o
}

func (o ImageReferenceOutput) ToImageReferenceOutputWithContext(ctx context.Context) ImageReferenceOutput {
	return o
}

func (o ImageReferenceOutput) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return o.ToImageReferencePtrOutputWithContext(context.Background())
}

func (o ImageReferenceOutput) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageReference) *ImageReference {
		return &v
	}).(ImageReferencePtrOutput)
}

func (o ImageReferenceOutput) CommunityGalleryImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.CommunityGalleryImageId }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) SharedGalleryImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.SharedGalleryImageId }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ImageReferencePtrOutput struct{ *pulumi.OutputState }

func (ImageReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageReference)(nil)).Elem()
}

func (o ImageReferencePtrOutput) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return o
}

func (o ImageReferencePtrOutput) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return o
}

func (o ImageReferencePtrOutput) Elem() ImageReferenceOutput {
	return o.ApplyT(func(v *ImageReference) ImageReference {
		if v != nil {
			return *v
		}
		var ret ImageReference
		return ret
	}).(ImageReferenceOutput)
}

func (o ImageReferencePtrOutput) CommunityGalleryImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.CommunityGalleryImageId
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) SharedGalleryImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.SharedGalleryImageId
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ImageReferenceResponse struct {
	CommunityGalleryImageId *string `pulumi:"communityGalleryImageId"`
	ExactVersion            string  `pulumi:"exactVersion"`
	Id                      *string `pulumi:"id"`
	Offer                   *string `pulumi:"offer"`
	Publisher               *string `pulumi:"publisher"`
	SharedGalleryImageId    *string `pulumi:"sharedGalleryImageId"`
	Sku                     *string `pulumi:"sku"`
	Version                 *string `pulumi:"version"`
}

type ImageReferenceResponseOutput struct{ *pulumi.OutputState }

func (ImageReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReferenceResponse)(nil)).Elem()
}

func (o ImageReferenceResponseOutput) ToImageReferenceResponseOutput() ImageReferenceResponseOutput {
	return o
}

func (o ImageReferenceResponseOutput) ToImageReferenceResponseOutputWithContext(ctx context.Context) ImageReferenceResponseOutput {
	return o
}

func (o ImageReferenceResponseOutput) CommunityGalleryImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.CommunityGalleryImageId }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) ExactVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ImageReferenceResponse) string { return v.ExactVersion }).(pulumi.StringOutput)
}

func (o ImageReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) SharedGalleryImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.SharedGalleryImageId }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ImageReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageReferenceResponse)(nil)).Elem()
}

func (o ImageReferenceResponsePtrOutput) ToImageReferenceResponsePtrOutput() ImageReferenceResponsePtrOutput {
	return o
}

func (o ImageReferenceResponsePtrOutput) ToImageReferenceResponsePtrOutputWithContext(ctx context.Context) ImageReferenceResponsePtrOutput {
	return o
}

func (o ImageReferenceResponsePtrOutput) Elem() ImageReferenceResponseOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) ImageReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ImageReferenceResponse
		return ret
	}).(ImageReferenceResponseOutput)
}

func (o ImageReferenceResponsePtrOutput) CommunityGalleryImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.CommunityGalleryImageId
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) ExactVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ExactVersion
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) SharedGalleryImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.SharedGalleryImageId
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ImageStorageProfile struct {
	DataDisks     []ImageDataDisk `pulumi:"dataDisks"`
	OsDisk        *ImageOSDisk    `pulumi:"osDisk"`
	ZoneResilient *bool           `pulumi:"zoneResilient"`
}





type ImageStorageProfileInput interface {
	pulumi.Input

	ToImageStorageProfileOutput() ImageStorageProfileOutput
	ToImageStorageProfileOutputWithContext(context.Context) ImageStorageProfileOutput
}

type ImageStorageProfileArgs struct {
	DataDisks     ImageDataDiskArrayInput `pulumi:"dataDisks"`
	OsDisk        ImageOSDiskPtrInput     `pulumi:"osDisk"`
	ZoneResilient pulumi.BoolPtrInput     `pulumi:"zoneResilient"`
}

func (ImageStorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageStorageProfile)(nil)).Elem()
}

func (i ImageStorageProfileArgs) ToImageStorageProfileOutput() ImageStorageProfileOutput {
	return i.ToImageStorageProfileOutputWithContext(context.Background())
}

func (i ImageStorageProfileArgs) ToImageStorageProfileOutputWithContext(ctx context.Context) ImageStorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageStorageProfileOutput)
}

func (i ImageStorageProfileArgs) ToImageStorageProfilePtrOutput() ImageStorageProfilePtrOutput {
	return i.ToImageStorageProfilePtrOutputWithContext(context.Background())
}

func (i ImageStorageProfileArgs) ToImageStorageProfilePtrOutputWithContext(ctx context.Context) ImageStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageStorageProfileOutput).ToImageStorageProfilePtrOutputWithContext(ctx)
}









type ImageStorageProfilePtrInput interface {
	pulumi.Input

	ToImageStorageProfilePtrOutput() ImageStorageProfilePtrOutput
	ToImageStorageProfilePtrOutputWithContext(context.Context) ImageStorageProfilePtrOutput
}

type imageStorageProfilePtrType ImageStorageProfileArgs

func ImageStorageProfilePtr(v *ImageStorageProfileArgs) ImageStorageProfilePtrInput {
	return (*imageStorageProfilePtrType)(v)
}

func (*imageStorageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageStorageProfile)(nil)).Elem()
}

func (i *imageStorageProfilePtrType) ToImageStorageProfilePtrOutput() ImageStorageProfilePtrOutput {
	return i.ToImageStorageProfilePtrOutputWithContext(context.Background())
}

func (i *imageStorageProfilePtrType) ToImageStorageProfilePtrOutputWithContext(ctx context.Context) ImageStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageStorageProfilePtrOutput)
}

type ImageStorageProfileOutput struct{ *pulumi.OutputState }

func (ImageStorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageStorageProfile)(nil)).Elem()
}

func (o ImageStorageProfileOutput) ToImageStorageProfileOutput() ImageStorageProfileOutput {
	return o
}

func (o ImageStorageProfileOutput) ToImageStorageProfileOutputWithContext(ctx context.Context) ImageStorageProfileOutput {
	return o
}

func (o ImageStorageProfileOutput) ToImageStorageProfilePtrOutput() ImageStorageProfilePtrOutput {
	return o.ToImageStorageProfilePtrOutputWithContext(context.Background())
}

func (o ImageStorageProfileOutput) ToImageStorageProfilePtrOutputWithContext(ctx context.Context) ImageStorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageStorageProfile) *ImageStorageProfile {
		return &v
	}).(ImageStorageProfilePtrOutput)
}

func (o ImageStorageProfileOutput) DataDisks() ImageDataDiskArrayOutput {
	return o.ApplyT(func(v ImageStorageProfile) []ImageDataDisk { return v.DataDisks }).(ImageDataDiskArrayOutput)
}

func (o ImageStorageProfileOutput) OsDisk() ImageOSDiskPtrOutput {
	return o.ApplyT(func(v ImageStorageProfile) *ImageOSDisk { return v.OsDisk }).(ImageOSDiskPtrOutput)
}

func (o ImageStorageProfileOutput) ZoneResilient() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ImageStorageProfile) *bool { return v.ZoneResilient }).(pulumi.BoolPtrOutput)
}

type ImageStorageProfilePtrOutput struct{ *pulumi.OutputState }

func (ImageStorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageStorageProfile)(nil)).Elem()
}

func (o ImageStorageProfilePtrOutput) ToImageStorageProfilePtrOutput() ImageStorageProfilePtrOutput {
	return o
}

func (o ImageStorageProfilePtrOutput) ToImageStorageProfilePtrOutputWithContext(ctx context.Context) ImageStorageProfilePtrOutput {
	return o
}

func (o ImageStorageProfilePtrOutput) Elem() ImageStorageProfileOutput {
	return o.ApplyT(func(v *ImageStorageProfile) ImageStorageProfile {
		if v != nil {
			return *v
		}
		var ret ImageStorageProfile
		return ret
	}).(ImageStorageProfileOutput)
}

func (o ImageStorageProfilePtrOutput) DataDisks() ImageDataDiskArrayOutput {
	return o.ApplyT(func(v *ImageStorageProfile) []ImageDataDisk {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(ImageDataDiskArrayOutput)
}

func (o ImageStorageProfilePtrOutput) OsDisk() ImageOSDiskPtrOutput {
	return o.ApplyT(func(v *ImageStorageProfile) *ImageOSDisk {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(ImageOSDiskPtrOutput)
}

func (o ImageStorageProfilePtrOutput) ZoneResilient() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ImageStorageProfile) *bool {
		if v == nil {
			return nil
		}
		return v.ZoneResilient
	}).(pulumi.BoolPtrOutput)
}

type ImageStorageProfileResponse struct {
	DataDisks     []ImageDataDiskResponse `pulumi:"dataDisks"`
	OsDisk        *ImageOSDiskResponse    `pulumi:"osDisk"`
	ZoneResilient *bool                   `pulumi:"zoneResilient"`
}

type ImageStorageProfileResponseOutput struct{ *pulumi.OutputState }

func (ImageStorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageStorageProfileResponse)(nil)).Elem()
}

func (o ImageStorageProfileResponseOutput) ToImageStorageProfileResponseOutput() ImageStorageProfileResponseOutput {
	return o
}

func (o ImageStorageProfileResponseOutput) ToImageStorageProfileResponseOutputWithContext(ctx context.Context) ImageStorageProfileResponseOutput {
	return o
}

func (o ImageStorageProfileResponseOutput) DataDisks() ImageDataDiskResponseArrayOutput {
	return o.ApplyT(func(v ImageStorageProfileResponse) []ImageDataDiskResponse { return v.DataDisks }).(ImageDataDiskResponseArrayOutput)
}

func (o ImageStorageProfileResponseOutput) OsDisk() ImageOSDiskResponsePtrOutput {
	return o.ApplyT(func(v ImageStorageProfileResponse) *ImageOSDiskResponse { return v.OsDisk }).(ImageOSDiskResponsePtrOutput)
}

func (o ImageStorageProfileResponseOutput) ZoneResilient() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ImageStorageProfileResponse) *bool { return v.ZoneResilient }).(pulumi.BoolPtrOutput)
}

type ImageStorageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageStorageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageStorageProfileResponse)(nil)).Elem()
}

func (o ImageStorageProfileResponsePtrOutput) ToImageStorageProfileResponsePtrOutput() ImageStorageProfileResponsePtrOutput {
	return o
}

func (o ImageStorageProfileResponsePtrOutput) ToImageStorageProfileResponsePtrOutputWithContext(ctx context.Context) ImageStorageProfileResponsePtrOutput {
	return o
}

func (o ImageStorageProfileResponsePtrOutput) Elem() ImageStorageProfileResponseOutput {
	return o.ApplyT(func(v *ImageStorageProfileResponse) ImageStorageProfileResponse {
		if v != nil {
			return *v
		}
		var ret ImageStorageProfileResponse
		return ret
	}).(ImageStorageProfileResponseOutput)
}

func (o ImageStorageProfileResponsePtrOutput) DataDisks() ImageDataDiskResponseArrayOutput {
	return o.ApplyT(func(v *ImageStorageProfileResponse) []ImageDataDiskResponse {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(ImageDataDiskResponseArrayOutput)
}

func (o ImageStorageProfileResponsePtrOutput) OsDisk() ImageOSDiskResponsePtrOutput {
	return o.ApplyT(func(v *ImageStorageProfileResponse) *ImageOSDiskResponse {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(ImageOSDiskResponsePtrOutput)
}

func (o ImageStorageProfileResponsePtrOutput) ZoneResilient() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ImageStorageProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ZoneResilient
	}).(pulumi.BoolPtrOutput)
}

type InnerErrorResponse struct {
	Errordetail   *string `pulumi:"errordetail"`
	Exceptiontype *string `pulumi:"exceptiontype"`
}

type InnerErrorResponseOutput struct{ *pulumi.OutputState }

func (InnerErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InnerErrorResponse)(nil)).Elem()
}

func (o InnerErrorResponseOutput) ToInnerErrorResponseOutput() InnerErrorResponseOutput {
	return o
}

func (o InnerErrorResponseOutput) ToInnerErrorResponseOutputWithContext(ctx context.Context) InnerErrorResponseOutput {
	return o
}

func (o InnerErrorResponseOutput) Errordetail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerErrorResponse) *string { return v.Errordetail }).(pulumi.StringPtrOutput)
}

func (o InnerErrorResponseOutput) Exceptiontype() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerErrorResponse) *string { return v.Exceptiontype }).(pulumi.StringPtrOutput)
}

type InnerErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (InnerErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InnerErrorResponse)(nil)).Elem()
}

func (o InnerErrorResponsePtrOutput) ToInnerErrorResponsePtrOutput() InnerErrorResponsePtrOutput {
	return o
}

func (o InnerErrorResponsePtrOutput) ToInnerErrorResponsePtrOutputWithContext(ctx context.Context) InnerErrorResponsePtrOutput {
	return o
}

func (o InnerErrorResponsePtrOutput) Elem() InnerErrorResponseOutput {
	return o.ApplyT(func(v *InnerErrorResponse) InnerErrorResponse {
		if v != nil {
			return *v
		}
		var ret InnerErrorResponse
		return ret
	}).(InnerErrorResponseOutput)
}

func (o InnerErrorResponsePtrOutput) Errordetail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InnerErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Errordetail
	}).(pulumi.StringPtrOutput)
}

func (o InnerErrorResponsePtrOutput) Exceptiontype() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InnerErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Exceptiontype
	}).(pulumi.StringPtrOutput)
}

type InstanceViewStatus struct {
	Code          *string           `pulumi:"code"`
	DisplayStatus *string           `pulumi:"displayStatus"`
	Level         *StatusLevelTypes `pulumi:"level"`
	Message       *string           `pulumi:"message"`
	Time          *string           `pulumi:"time"`
}





type InstanceViewStatusInput interface {
	pulumi.Input

	ToInstanceViewStatusOutput() InstanceViewStatusOutput
	ToInstanceViewStatusOutputWithContext(context.Context) InstanceViewStatusOutput
}

type InstanceViewStatusArgs struct {
	Code          pulumi.StringPtrInput    `pulumi:"code"`
	DisplayStatus pulumi.StringPtrInput    `pulumi:"displayStatus"`
	Level         StatusLevelTypesPtrInput `pulumi:"level"`
	Message       pulumi.StringPtrInput    `pulumi:"message"`
	Time          pulumi.StringPtrInput    `pulumi:"time"`
}

func (InstanceViewStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceViewStatus)(nil)).Elem()
}

func (i InstanceViewStatusArgs) ToInstanceViewStatusOutput() InstanceViewStatusOutput {
	return i.ToInstanceViewStatusOutputWithContext(context.Background())
}

func (i InstanceViewStatusArgs) ToInstanceViewStatusOutputWithContext(ctx context.Context) InstanceViewStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceViewStatusOutput)
}

func (i InstanceViewStatusArgs) ToInstanceViewStatusPtrOutput() InstanceViewStatusPtrOutput {
	return i.ToInstanceViewStatusPtrOutputWithContext(context.Background())
}

func (i InstanceViewStatusArgs) ToInstanceViewStatusPtrOutputWithContext(ctx context.Context) InstanceViewStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceViewStatusOutput).ToInstanceViewStatusPtrOutputWithContext(ctx)
}









type InstanceViewStatusPtrInput interface {
	pulumi.Input

	ToInstanceViewStatusPtrOutput() InstanceViewStatusPtrOutput
	ToInstanceViewStatusPtrOutputWithContext(context.Context) InstanceViewStatusPtrOutput
}

type instanceViewStatusPtrType InstanceViewStatusArgs

func InstanceViewStatusPtr(v *InstanceViewStatusArgs) InstanceViewStatusPtrInput {
	return (*instanceViewStatusPtrType)(v)
}

func (*instanceViewStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceViewStatus)(nil)).Elem()
}

func (i *instanceViewStatusPtrType) ToInstanceViewStatusPtrOutput() InstanceViewStatusPtrOutput {
	return i.ToInstanceViewStatusPtrOutputWithContext(context.Background())
}

func (i *instanceViewStatusPtrType) ToInstanceViewStatusPtrOutputWithContext(ctx context.Context) InstanceViewStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceViewStatusPtrOutput)
}





type InstanceViewStatusArrayInput interface {
	pulumi.Input

	ToInstanceViewStatusArrayOutput() InstanceViewStatusArrayOutput
	ToInstanceViewStatusArrayOutputWithContext(context.Context) InstanceViewStatusArrayOutput
}

type InstanceViewStatusArray []InstanceViewStatusInput

func (InstanceViewStatusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceViewStatus)(nil)).Elem()
}

func (i InstanceViewStatusArray) ToInstanceViewStatusArrayOutput() InstanceViewStatusArrayOutput {
	return i.ToInstanceViewStatusArrayOutputWithContext(context.Background())
}

func (i InstanceViewStatusArray) ToInstanceViewStatusArrayOutputWithContext(ctx context.Context) InstanceViewStatusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceViewStatusArrayOutput)
}

type InstanceViewStatusOutput struct{ *pulumi.OutputState }

func (InstanceViewStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceViewStatus)(nil)).Elem()
}

func (o InstanceViewStatusOutput) ToInstanceViewStatusOutput() InstanceViewStatusOutput {
	return o
}

func (o InstanceViewStatusOutput) ToInstanceViewStatusOutputWithContext(ctx context.Context) InstanceViewStatusOutput {
	return o
}

func (o InstanceViewStatusOutput) ToInstanceViewStatusPtrOutput() InstanceViewStatusPtrOutput {
	return o.ToInstanceViewStatusPtrOutputWithContext(context.Background())
}

func (o InstanceViewStatusOutput) ToInstanceViewStatusPtrOutputWithContext(ctx context.Context) InstanceViewStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceViewStatus) *InstanceViewStatus {
		return &v
	}).(InstanceViewStatusPtrOutput)
}

func (o InstanceViewStatusOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatus) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatus) *string { return v.DisplayStatus }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusOutput) Level() StatusLevelTypesPtrOutput {
	return o.ApplyT(func(v InstanceViewStatus) *StatusLevelTypes { return v.Level }).(StatusLevelTypesPtrOutput)
}

func (o InstanceViewStatusOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatus) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatus) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type InstanceViewStatusPtrOutput struct{ *pulumi.OutputState }

func (InstanceViewStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceViewStatus)(nil)).Elem()
}

func (o InstanceViewStatusPtrOutput) ToInstanceViewStatusPtrOutput() InstanceViewStatusPtrOutput {
	return o
}

func (o InstanceViewStatusPtrOutput) ToInstanceViewStatusPtrOutputWithContext(ctx context.Context) InstanceViewStatusPtrOutput {
	return o
}

func (o InstanceViewStatusPtrOutput) Elem() InstanceViewStatusOutput {
	return o.ApplyT(func(v *InstanceViewStatus) InstanceViewStatus {
		if v != nil {
			return *v
		}
		var ret InstanceViewStatus
		return ret
	}).(InstanceViewStatusOutput)
}

func (o InstanceViewStatusPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatus) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusPtrOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatus) *string {
		if v == nil {
			return nil
		}
		return v.DisplayStatus
	}).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusPtrOutput) Level() StatusLevelTypesPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatus) *StatusLevelTypes {
		if v == nil {
			return nil
		}
		return v.Level
	}).(StatusLevelTypesPtrOutput)
}

func (o InstanceViewStatusPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatus) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusPtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatus) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

type InstanceViewStatusArrayOutput struct{ *pulumi.OutputState }

func (InstanceViewStatusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceViewStatus)(nil)).Elem()
}

func (o InstanceViewStatusArrayOutput) ToInstanceViewStatusArrayOutput() InstanceViewStatusArrayOutput {
	return o
}

func (o InstanceViewStatusArrayOutput) ToInstanceViewStatusArrayOutputWithContext(ctx context.Context) InstanceViewStatusArrayOutput {
	return o
}

func (o InstanceViewStatusArrayOutput) Index(i pulumi.IntInput) InstanceViewStatusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceViewStatus {
		return vs[0].([]InstanceViewStatus)[vs[1].(int)]
	}).(InstanceViewStatusOutput)
}

type InstanceViewStatusResponse struct {
	Code          *string `pulumi:"code"`
	DisplayStatus *string `pulumi:"displayStatus"`
	Level         *string `pulumi:"level"`
	Message       *string `pulumi:"message"`
	Time          *string `pulumi:"time"`
}

type InstanceViewStatusResponseOutput struct{ *pulumi.OutputState }

func (InstanceViewStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceViewStatusResponse)(nil)).Elem()
}

func (o InstanceViewStatusResponseOutput) ToInstanceViewStatusResponseOutput() InstanceViewStatusResponseOutput {
	return o
}

func (o InstanceViewStatusResponseOutput) ToInstanceViewStatusResponseOutputWithContext(ctx context.Context) InstanceViewStatusResponseOutput {
	return o
}

func (o InstanceViewStatusResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatusResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponseOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatusResponse) *string { return v.DisplayStatus }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatusResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatusResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponseOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceViewStatusResponse) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type InstanceViewStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (InstanceViewStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceViewStatusResponse)(nil)).Elem()
}

func (o InstanceViewStatusResponsePtrOutput) ToInstanceViewStatusResponsePtrOutput() InstanceViewStatusResponsePtrOutput {
	return o
}

func (o InstanceViewStatusResponsePtrOutput) ToInstanceViewStatusResponsePtrOutputWithContext(ctx context.Context) InstanceViewStatusResponsePtrOutput {
	return o
}

func (o InstanceViewStatusResponsePtrOutput) Elem() InstanceViewStatusResponseOutput {
	return o.ApplyT(func(v *InstanceViewStatusResponse) InstanceViewStatusResponse {
		if v != nil {
			return *v
		}
		var ret InstanceViewStatusResponse
		return ret
	}).(InstanceViewStatusResponseOutput)
}

func (o InstanceViewStatusResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponsePtrOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisplayStatus
	}).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponsePtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Level
	}).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o InstanceViewStatusResponsePtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceViewStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

type InstanceViewStatusResponseArrayOutput struct{ *pulumi.OutputState }

func (InstanceViewStatusResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceViewStatusResponse)(nil)).Elem()
}

func (o InstanceViewStatusResponseArrayOutput) ToInstanceViewStatusResponseArrayOutput() InstanceViewStatusResponseArrayOutput {
	return o
}

func (o InstanceViewStatusResponseArrayOutput) ToInstanceViewStatusResponseArrayOutputWithContext(ctx context.Context) InstanceViewStatusResponseArrayOutput {
	return o
}

func (o InstanceViewStatusResponseArrayOutput) Index(i pulumi.IntInput) InstanceViewStatusResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceViewStatusResponse {
		return vs[0].([]InstanceViewStatusResponse)[vs[1].(int)]
	}).(InstanceViewStatusResponseOutput)
}

type KeyVaultKeyReference struct {
	KeyUrl      string      `pulumi:"keyUrl"`
	SourceVault SubResource `pulumi:"sourceVault"`
}





type KeyVaultKeyReferenceInput interface {
	pulumi.Input

	ToKeyVaultKeyReferenceOutput() KeyVaultKeyReferenceOutput
	ToKeyVaultKeyReferenceOutputWithContext(context.Context) KeyVaultKeyReferenceOutput
}

type KeyVaultKeyReferenceArgs struct {
	KeyUrl      pulumi.StringInput `pulumi:"keyUrl"`
	SourceVault SubResourceInput   `pulumi:"sourceVault"`
}

func (KeyVaultKeyReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReference)(nil)).Elem()
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferenceOutput() KeyVaultKeyReferenceOutput {
	return i.ToKeyVaultKeyReferenceOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferenceOutputWithContext(ctx context.Context) KeyVaultKeyReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceOutput)
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return i.ToKeyVaultKeyReferencePtrOutputWithContext(context.Background())
}

func (i KeyVaultKeyReferenceArgs) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferenceOutput).ToKeyVaultKeyReferencePtrOutputWithContext(ctx)
}









type KeyVaultKeyReferencePtrInput interface {
	pulumi.Input

	ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput
	ToKeyVaultKeyReferencePtrOutputWithContext(context.Context) KeyVaultKeyReferencePtrOutput
}

type keyVaultKeyReferencePtrType KeyVaultKeyReferenceArgs

func KeyVaultKeyReferencePtr(v *KeyVaultKeyReferenceArgs) KeyVaultKeyReferencePtrInput {
	return (*keyVaultKeyReferencePtrType)(v)
}

func (*keyVaultKeyReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReference)(nil)).Elem()
}

func (i *keyVaultKeyReferencePtrType) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return i.ToKeyVaultKeyReferencePtrOutputWithContext(context.Background())
}

func (i *keyVaultKeyReferencePtrType) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyReferencePtrOutput)
}

type KeyVaultKeyReferenceOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReference)(nil)).Elem()
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferenceOutput() KeyVaultKeyReferenceOutput {
	return o
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferenceOutputWithContext(ctx context.Context) KeyVaultKeyReferenceOutput {
	return o
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return o.ToKeyVaultKeyReferencePtrOutputWithContext(context.Background())
}

func (o KeyVaultKeyReferenceOutput) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultKeyReference) *KeyVaultKeyReference {
		return &v
	}).(KeyVaultKeyReferencePtrOutput)
}

func (o KeyVaultKeyReferenceOutput) KeyUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultKeyReference) string { return v.KeyUrl }).(pulumi.StringOutput)
}

func (o KeyVaultKeyReferenceOutput) SourceVault() SubResourceOutput {
	return o.ApplyT(func(v KeyVaultKeyReference) SubResource { return v.SourceVault }).(SubResourceOutput)
}

type KeyVaultKeyReferencePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReference)(nil)).Elem()
}

func (o KeyVaultKeyReferencePtrOutput) ToKeyVaultKeyReferencePtrOutput() KeyVaultKeyReferencePtrOutput {
	return o
}

func (o KeyVaultKeyReferencePtrOutput) ToKeyVaultKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferencePtrOutput {
	return o
}

func (o KeyVaultKeyReferencePtrOutput) Elem() KeyVaultKeyReferenceOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) KeyVaultKeyReference {
		if v != nil {
			return *v
		}
		var ret KeyVaultKeyReference
		return ret
	}).(KeyVaultKeyReferenceOutput)
}

func (o KeyVaultKeyReferencePtrOutput) KeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) *string {
		if v == nil {
			return nil
		}
		return &v.KeyUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferencePtrOutput) SourceVault() SubResourcePtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReference) *SubResource {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SubResourcePtrOutput)
}

type KeyVaultKeyReferenceResponse struct {
	KeyUrl      string              `pulumi:"keyUrl"`
	SourceVault SubResourceResponse `pulumi:"sourceVault"`
}

type KeyVaultKeyReferenceResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyReferenceResponse)(nil)).Elem()
}

func (o KeyVaultKeyReferenceResponseOutput) ToKeyVaultKeyReferenceResponseOutput() KeyVaultKeyReferenceResponseOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseOutput) ToKeyVaultKeyReferenceResponseOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponseOutput {
	return o
}

func (o KeyVaultKeyReferenceResponseOutput) KeyUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponse) string { return v.KeyUrl }).(pulumi.StringOutput)
}

func (o KeyVaultKeyReferenceResponseOutput) SourceVault() SubResourceResponseOutput {
	return o.ApplyT(func(v KeyVaultKeyReferenceResponse) SubResourceResponse { return v.SourceVault }).(SubResourceResponseOutput)
}

type KeyVaultKeyReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultKeyReferenceResponse)(nil)).Elem()
}

func (o KeyVaultKeyReferenceResponsePtrOutput) ToKeyVaultKeyReferenceResponsePtrOutput() KeyVaultKeyReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultKeyReferenceResponsePtrOutput) ToKeyVaultKeyReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultKeyReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultKeyReferenceResponsePtrOutput) Elem() KeyVaultKeyReferenceResponseOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) KeyVaultKeyReferenceResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultKeyReferenceResponse
		return ret
	}).(KeyVaultKeyReferenceResponseOutput)
}

func (o KeyVaultKeyReferenceResponsePtrOutput) KeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultKeyReferenceResponsePtrOutput) SourceVault() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *KeyVaultKeyReferenceResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SubResourceResponsePtrOutput)
}

type KeyVaultSecretReference struct {
	SecretUrl   string      `pulumi:"secretUrl"`
	SourceVault SubResource `pulumi:"sourceVault"`
}





type KeyVaultSecretReferenceInput interface {
	pulumi.Input

	ToKeyVaultSecretReferenceOutput() KeyVaultSecretReferenceOutput
	ToKeyVaultSecretReferenceOutputWithContext(context.Context) KeyVaultSecretReferenceOutput
}

type KeyVaultSecretReferenceArgs struct {
	SecretUrl   pulumi.StringInput `pulumi:"secretUrl"`
	SourceVault SubResourceInput   `pulumi:"sourceVault"`
}

func (KeyVaultSecretReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultSecretReference)(nil)).Elem()
}

func (i KeyVaultSecretReferenceArgs) ToKeyVaultSecretReferenceOutput() KeyVaultSecretReferenceOutput {
	return i.ToKeyVaultSecretReferenceOutputWithContext(context.Background())
}

func (i KeyVaultSecretReferenceArgs) ToKeyVaultSecretReferenceOutputWithContext(ctx context.Context) KeyVaultSecretReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultSecretReferenceOutput)
}

func (i KeyVaultSecretReferenceArgs) ToKeyVaultSecretReferencePtrOutput() KeyVaultSecretReferencePtrOutput {
	return i.ToKeyVaultSecretReferencePtrOutputWithContext(context.Background())
}

func (i KeyVaultSecretReferenceArgs) ToKeyVaultSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultSecretReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultSecretReferenceOutput).ToKeyVaultSecretReferencePtrOutputWithContext(ctx)
}









type KeyVaultSecretReferencePtrInput interface {
	pulumi.Input

	ToKeyVaultSecretReferencePtrOutput() KeyVaultSecretReferencePtrOutput
	ToKeyVaultSecretReferencePtrOutputWithContext(context.Context) KeyVaultSecretReferencePtrOutput
}

type keyVaultSecretReferencePtrType KeyVaultSecretReferenceArgs

func KeyVaultSecretReferencePtr(v *KeyVaultSecretReferenceArgs) KeyVaultSecretReferencePtrInput {
	return (*keyVaultSecretReferencePtrType)(v)
}

func (*keyVaultSecretReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultSecretReference)(nil)).Elem()
}

func (i *keyVaultSecretReferencePtrType) ToKeyVaultSecretReferencePtrOutput() KeyVaultSecretReferencePtrOutput {
	return i.ToKeyVaultSecretReferencePtrOutputWithContext(context.Background())
}

func (i *keyVaultSecretReferencePtrType) ToKeyVaultSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultSecretReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultSecretReferencePtrOutput)
}

type KeyVaultSecretReferenceOutput struct{ *pulumi.OutputState }

func (KeyVaultSecretReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultSecretReference)(nil)).Elem()
}

func (o KeyVaultSecretReferenceOutput) ToKeyVaultSecretReferenceOutput() KeyVaultSecretReferenceOutput {
	return o
}

func (o KeyVaultSecretReferenceOutput) ToKeyVaultSecretReferenceOutputWithContext(ctx context.Context) KeyVaultSecretReferenceOutput {
	return o
}

func (o KeyVaultSecretReferenceOutput) ToKeyVaultSecretReferencePtrOutput() KeyVaultSecretReferencePtrOutput {
	return o.ToKeyVaultSecretReferencePtrOutputWithContext(context.Background())
}

func (o KeyVaultSecretReferenceOutput) ToKeyVaultSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultSecretReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultSecretReference) *KeyVaultSecretReference {
		return &v
	}).(KeyVaultSecretReferencePtrOutput)
}

func (o KeyVaultSecretReferenceOutput) SecretUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSecretReference) string { return v.SecretUrl }).(pulumi.StringOutput)
}

func (o KeyVaultSecretReferenceOutput) SourceVault() SubResourceOutput {
	return o.ApplyT(func(v KeyVaultSecretReference) SubResource { return v.SourceVault }).(SubResourceOutput)
}

type KeyVaultSecretReferencePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultSecretReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultSecretReference)(nil)).Elem()
}

func (o KeyVaultSecretReferencePtrOutput) ToKeyVaultSecretReferencePtrOutput() KeyVaultSecretReferencePtrOutput {
	return o
}

func (o KeyVaultSecretReferencePtrOutput) ToKeyVaultSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultSecretReferencePtrOutput {
	return o
}

func (o KeyVaultSecretReferencePtrOutput) Elem() KeyVaultSecretReferenceOutput {
	return o.ApplyT(func(v *KeyVaultSecretReference) KeyVaultSecretReference {
		if v != nil {
			return *v
		}
		var ret KeyVaultSecretReference
		return ret
	}).(KeyVaultSecretReferenceOutput)
}

func (o KeyVaultSecretReferencePtrOutput) SecretUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultSecretReference) *string {
		if v == nil {
			return nil
		}
		return &v.SecretUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultSecretReferencePtrOutput) SourceVault() SubResourcePtrOutput {
	return o.ApplyT(func(v *KeyVaultSecretReference) *SubResource {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SubResourcePtrOutput)
}

type KeyVaultSecretReferenceResponse struct {
	SecretUrl   string              `pulumi:"secretUrl"`
	SourceVault SubResourceResponse `pulumi:"sourceVault"`
}

type KeyVaultSecretReferenceResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultSecretReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultSecretReferenceResponse)(nil)).Elem()
}

func (o KeyVaultSecretReferenceResponseOutput) ToKeyVaultSecretReferenceResponseOutput() KeyVaultSecretReferenceResponseOutput {
	return o
}

func (o KeyVaultSecretReferenceResponseOutput) ToKeyVaultSecretReferenceResponseOutputWithContext(ctx context.Context) KeyVaultSecretReferenceResponseOutput {
	return o
}

func (o KeyVaultSecretReferenceResponseOutput) SecretUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultSecretReferenceResponse) string { return v.SecretUrl }).(pulumi.StringOutput)
}

func (o KeyVaultSecretReferenceResponseOutput) SourceVault() SubResourceResponseOutput {
	return o.ApplyT(func(v KeyVaultSecretReferenceResponse) SubResourceResponse { return v.SourceVault }).(SubResourceResponseOutput)
}

type KeyVaultSecretReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultSecretReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultSecretReferenceResponse)(nil)).Elem()
}

func (o KeyVaultSecretReferenceResponsePtrOutput) ToKeyVaultSecretReferenceResponsePtrOutput() KeyVaultSecretReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultSecretReferenceResponsePtrOutput) ToKeyVaultSecretReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultSecretReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultSecretReferenceResponsePtrOutput) Elem() KeyVaultSecretReferenceResponseOutput {
	return o.ApplyT(func(v *KeyVaultSecretReferenceResponse) KeyVaultSecretReferenceResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultSecretReferenceResponse
		return ret
	}).(KeyVaultSecretReferenceResponseOutput)
}

func (o KeyVaultSecretReferenceResponsePtrOutput) SecretUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultSecretReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecretUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultSecretReferenceResponsePtrOutput) SourceVault() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *KeyVaultSecretReferenceResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SubResourceResponsePtrOutput)
}

type LastPatchInstallationSummaryResponse struct {
	Error                     ApiErrorResponse `pulumi:"error"`
	ExcludedPatchCount        int              `pulumi:"excludedPatchCount"`
	FailedPatchCount          int              `pulumi:"failedPatchCount"`
	InstallationActivityId    string           `pulumi:"installationActivityId"`
	InstalledPatchCount       int              `pulumi:"installedPatchCount"`
	LastModifiedTime          string           `pulumi:"lastModifiedTime"`
	MaintenanceWindowExceeded bool             `pulumi:"maintenanceWindowExceeded"`
	NotSelectedPatchCount     int              `pulumi:"notSelectedPatchCount"`
	PendingPatchCount         int              `pulumi:"pendingPatchCount"`
	StartTime                 string           `pulumi:"startTime"`
	Status                    string           `pulumi:"status"`
}

type LastPatchInstallationSummaryResponseOutput struct{ *pulumi.OutputState }

func (LastPatchInstallationSummaryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LastPatchInstallationSummaryResponse)(nil)).Elem()
}

func (o LastPatchInstallationSummaryResponseOutput) ToLastPatchInstallationSummaryResponseOutput() LastPatchInstallationSummaryResponseOutput {
	return o
}

func (o LastPatchInstallationSummaryResponseOutput) ToLastPatchInstallationSummaryResponseOutputWithContext(ctx context.Context) LastPatchInstallationSummaryResponseOutput {
	return o
}

func (o LastPatchInstallationSummaryResponseOutput) Error() ApiErrorResponseOutput {
	return o.ApplyT(func(v LastPatchInstallationSummaryResponse) ApiErrorResponse { return v.Error }).(ApiErrorResponseOutput)
}

func (o LastPatchInstallationSummaryResponseOutput) ExcludedPatchCount() pulumi.IntOutput {
	return o.ApplyT(func(v LastPatchInstallationSummaryResponse) int { return v.ExcludedPatchCount }).(pulumi.IntOutput)
}

func (o LastPatchInstallationSummaryResponseOutput) FailedPatchCount() pulumi.IntOutput {
	return o.ApplyT(func(v LastPatchInstallationSummaryResponse) int { return v.FailedPatchCount }).(pulumi.IntOutput)
}

func (o LastPatchInstallationSummaryResponseOutput) InstallationActivityId() pulumi.StringOutput {
	return o.ApplyT(func(v LastPatchInstallationSummaryResponse) string { return v.InstallationActivityId }).(pulumi.StringOutput)
}

func (o LastPatchInstallationSummaryResponseOutput) InstalledPatchCount() pulumi.IntOutput {
	return o.ApplyT(func(v LastPatchInstallationSummaryResponse) int { return v.InstalledPatchCount }).(pulumi.IntOutput)
}

func (o LastPatchInstallationSummaryResponseOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LastPatchInstallationSummaryResponse) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LastPatchInstallationSummaryResponseOutput) MaintenanceWindowExceeded() pulumi.BoolOutput {
	return o.ApplyT(func(v LastPatchInstallationSummaryResponse) bool { return v.MaintenanceWindowExceeded }).(pulumi.BoolOutput)
}

func (o LastPatchInstallationSummaryResponseOutput) NotSelectedPatchCount() pulumi.IntOutput {
	return o.ApplyT(func(v LastPatchInstallationSummaryResponse) int { return v.NotSelectedPatchCount }).(pulumi.IntOutput)
}

func (o LastPatchInstallationSummaryResponseOutput) PendingPatchCount() pulumi.IntOutput {
	return o.ApplyT(func(v LastPatchInstallationSummaryResponse) int { return v.PendingPatchCount }).(pulumi.IntOutput)
}

func (o LastPatchInstallationSummaryResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LastPatchInstallationSummaryResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o LastPatchInstallationSummaryResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LastPatchInstallationSummaryResponse) string { return v.Status }).(pulumi.StringOutput)
}

type LastPatchInstallationSummaryResponsePtrOutput struct{ *pulumi.OutputState }

func (LastPatchInstallationSummaryResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LastPatchInstallationSummaryResponse)(nil)).Elem()
}

func (o LastPatchInstallationSummaryResponsePtrOutput) ToLastPatchInstallationSummaryResponsePtrOutput() LastPatchInstallationSummaryResponsePtrOutput {
	return o
}

func (o LastPatchInstallationSummaryResponsePtrOutput) ToLastPatchInstallationSummaryResponsePtrOutputWithContext(ctx context.Context) LastPatchInstallationSummaryResponsePtrOutput {
	return o
}

func (o LastPatchInstallationSummaryResponsePtrOutput) Elem() LastPatchInstallationSummaryResponseOutput {
	return o.ApplyT(func(v *LastPatchInstallationSummaryResponse) LastPatchInstallationSummaryResponse {
		if v != nil {
			return *v
		}
		var ret LastPatchInstallationSummaryResponse
		return ret
	}).(LastPatchInstallationSummaryResponseOutput)
}

func (o LastPatchInstallationSummaryResponsePtrOutput) Error() ApiErrorResponsePtrOutput {
	return o.ApplyT(func(v *LastPatchInstallationSummaryResponse) *ApiErrorResponse {
		if v == nil {
			return nil
		}
		return &v.Error
	}).(ApiErrorResponsePtrOutput)
}

func (o LastPatchInstallationSummaryResponsePtrOutput) ExcludedPatchCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LastPatchInstallationSummaryResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ExcludedPatchCount
	}).(pulumi.IntPtrOutput)
}

func (o LastPatchInstallationSummaryResponsePtrOutput) FailedPatchCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LastPatchInstallationSummaryResponse) *int {
		if v == nil {
			return nil
		}
		return &v.FailedPatchCount
	}).(pulumi.IntPtrOutput)
}

func (o LastPatchInstallationSummaryResponsePtrOutput) InstallationActivityId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LastPatchInstallationSummaryResponse) *string {
		if v == nil {
			return nil
		}
		return &v.InstallationActivityId
	}).(pulumi.StringPtrOutput)
}

func (o LastPatchInstallationSummaryResponsePtrOutput) InstalledPatchCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LastPatchInstallationSummaryResponse) *int {
		if v == nil {
			return nil
		}
		return &v.InstalledPatchCount
	}).(pulumi.IntPtrOutput)
}

func (o LastPatchInstallationSummaryResponsePtrOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LastPatchInstallationSummaryResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastModifiedTime
	}).(pulumi.StringPtrOutput)
}

func (o LastPatchInstallationSummaryResponsePtrOutput) MaintenanceWindowExceeded() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LastPatchInstallationSummaryResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.MaintenanceWindowExceeded
	}).(pulumi.BoolPtrOutput)
}

func (o LastPatchInstallationSummaryResponsePtrOutput) NotSelectedPatchCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LastPatchInstallationSummaryResponse) *int {
		if v == nil {
			return nil
		}
		return &v.NotSelectedPatchCount
	}).(pulumi.IntPtrOutput)
}

func (o LastPatchInstallationSummaryResponsePtrOutput) PendingPatchCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LastPatchInstallationSummaryResponse) *int {
		if v == nil {
			return nil
		}
		return &v.PendingPatchCount
	}).(pulumi.IntPtrOutput)
}

func (o LastPatchInstallationSummaryResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LastPatchInstallationSummaryResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o LastPatchInstallationSummaryResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LastPatchInstallationSummaryResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type LinuxConfiguration struct {
	DisablePasswordAuthentication *bool               `pulumi:"disablePasswordAuthentication"`
	PatchSettings                 *LinuxPatchSettings `pulumi:"patchSettings"`
	ProvisionVMAgent              *bool               `pulumi:"provisionVMAgent"`
	Ssh                           *SshConfiguration   `pulumi:"ssh"`
}





type LinuxConfigurationInput interface {
	pulumi.Input

	ToLinuxConfigurationOutput() LinuxConfigurationOutput
	ToLinuxConfigurationOutputWithContext(context.Context) LinuxConfigurationOutput
}

type LinuxConfigurationArgs struct {
	DisablePasswordAuthentication pulumi.BoolPtrInput        `pulumi:"disablePasswordAuthentication"`
	PatchSettings                 LinuxPatchSettingsPtrInput `pulumi:"patchSettings"`
	ProvisionVMAgent              pulumi.BoolPtrInput        `pulumi:"provisionVMAgent"`
	Ssh                           SshConfigurationPtrInput   `pulumi:"ssh"`
}

func (LinuxConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxConfiguration)(nil)).Elem()
}

func (i LinuxConfigurationArgs) ToLinuxConfigurationOutput() LinuxConfigurationOutput {
	return i.ToLinuxConfigurationOutputWithContext(context.Background())
}

func (i LinuxConfigurationArgs) ToLinuxConfigurationOutputWithContext(ctx context.Context) LinuxConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxConfigurationOutput)
}

func (i LinuxConfigurationArgs) ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput {
	return i.ToLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (i LinuxConfigurationArgs) ToLinuxConfigurationPtrOutputWithContext(ctx context.Context) LinuxConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxConfigurationOutput).ToLinuxConfigurationPtrOutputWithContext(ctx)
}









type LinuxConfigurationPtrInput interface {
	pulumi.Input

	ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput
	ToLinuxConfigurationPtrOutputWithContext(context.Context) LinuxConfigurationPtrOutput
}

type linuxConfigurationPtrType LinuxConfigurationArgs

func LinuxConfigurationPtr(v *LinuxConfigurationArgs) LinuxConfigurationPtrInput {
	return (*linuxConfigurationPtrType)(v)
}

func (*linuxConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxConfiguration)(nil)).Elem()
}

func (i *linuxConfigurationPtrType) ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput {
	return i.ToLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (i *linuxConfigurationPtrType) ToLinuxConfigurationPtrOutputWithContext(ctx context.Context) LinuxConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxConfigurationPtrOutput)
}

type LinuxConfigurationOutput struct{ *pulumi.OutputState }

func (LinuxConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxConfiguration)(nil)).Elem()
}

func (o LinuxConfigurationOutput) ToLinuxConfigurationOutput() LinuxConfigurationOutput {
	return o
}

func (o LinuxConfigurationOutput) ToLinuxConfigurationOutputWithContext(ctx context.Context) LinuxConfigurationOutput {
	return o
}

func (o LinuxConfigurationOutput) ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput {
	return o.ToLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (o LinuxConfigurationOutput) ToLinuxConfigurationPtrOutputWithContext(ctx context.Context) LinuxConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinuxConfiguration) *LinuxConfiguration {
		return &v
	}).(LinuxConfigurationPtrOutput)
}

func (o LinuxConfigurationOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LinuxConfiguration) *bool { return v.DisablePasswordAuthentication }).(pulumi.BoolPtrOutput)
}

func (o LinuxConfigurationOutput) PatchSettings() LinuxPatchSettingsPtrOutput {
	return o.ApplyT(func(v LinuxConfiguration) *LinuxPatchSettings { return v.PatchSettings }).(LinuxPatchSettingsPtrOutput)
}

func (o LinuxConfigurationOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LinuxConfiguration) *bool { return v.ProvisionVMAgent }).(pulumi.BoolPtrOutput)
}

func (o LinuxConfigurationOutput) Ssh() SshConfigurationPtrOutput {
	return o.ApplyT(func(v LinuxConfiguration) *SshConfiguration { return v.Ssh }).(SshConfigurationPtrOutput)
}

type LinuxConfigurationPtrOutput struct{ *pulumi.OutputState }

func (LinuxConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxConfiguration)(nil)).Elem()
}

func (o LinuxConfigurationPtrOutput) ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput {
	return o
}

func (o LinuxConfigurationPtrOutput) ToLinuxConfigurationPtrOutputWithContext(ctx context.Context) LinuxConfigurationPtrOutput {
	return o
}

func (o LinuxConfigurationPtrOutput) Elem() LinuxConfigurationOutput {
	return o.ApplyT(func(v *LinuxConfiguration) LinuxConfiguration {
		if v != nil {
			return *v
		}
		var ret LinuxConfiguration
		return ret
	}).(LinuxConfigurationOutput)
}

func (o LinuxConfigurationPtrOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LinuxConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.DisablePasswordAuthentication
	}).(pulumi.BoolPtrOutput)
}

func (o LinuxConfigurationPtrOutput) PatchSettings() LinuxPatchSettingsPtrOutput {
	return o.ApplyT(func(v *LinuxConfiguration) *LinuxPatchSettings {
		if v == nil {
			return nil
		}
		return v.PatchSettings
	}).(LinuxPatchSettingsPtrOutput)
}

func (o LinuxConfigurationPtrOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LinuxConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.ProvisionVMAgent
	}).(pulumi.BoolPtrOutput)
}

func (o LinuxConfigurationPtrOutput) Ssh() SshConfigurationPtrOutput {
	return o.ApplyT(func(v *LinuxConfiguration) *SshConfiguration {
		if v == nil {
			return nil
		}
		return v.Ssh
	}).(SshConfigurationPtrOutput)
}

type LinuxConfigurationResponse struct {
	DisablePasswordAuthentication *bool                       `pulumi:"disablePasswordAuthentication"`
	PatchSettings                 *LinuxPatchSettingsResponse `pulumi:"patchSettings"`
	ProvisionVMAgent              *bool                       `pulumi:"provisionVMAgent"`
	Ssh                           *SshConfigurationResponse   `pulumi:"ssh"`
}

type LinuxConfigurationResponseOutput struct{ *pulumi.OutputState }

func (LinuxConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxConfigurationResponse)(nil)).Elem()
}

func (o LinuxConfigurationResponseOutput) ToLinuxConfigurationResponseOutput() LinuxConfigurationResponseOutput {
	return o
}

func (o LinuxConfigurationResponseOutput) ToLinuxConfigurationResponseOutputWithContext(ctx context.Context) LinuxConfigurationResponseOutput {
	return o
}

func (o LinuxConfigurationResponseOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LinuxConfigurationResponse) *bool { return v.DisablePasswordAuthentication }).(pulumi.BoolPtrOutput)
}

func (o LinuxConfigurationResponseOutput) PatchSettings() LinuxPatchSettingsResponsePtrOutput {
	return o.ApplyT(func(v LinuxConfigurationResponse) *LinuxPatchSettingsResponse { return v.PatchSettings }).(LinuxPatchSettingsResponsePtrOutput)
}

func (o LinuxConfigurationResponseOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LinuxConfigurationResponse) *bool { return v.ProvisionVMAgent }).(pulumi.BoolPtrOutput)
}

func (o LinuxConfigurationResponseOutput) Ssh() SshConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LinuxConfigurationResponse) *SshConfigurationResponse { return v.Ssh }).(SshConfigurationResponsePtrOutput)
}

type LinuxConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (LinuxConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxConfigurationResponse)(nil)).Elem()
}

func (o LinuxConfigurationResponsePtrOutput) ToLinuxConfigurationResponsePtrOutput() LinuxConfigurationResponsePtrOutput {
	return o
}

func (o LinuxConfigurationResponsePtrOutput) ToLinuxConfigurationResponsePtrOutputWithContext(ctx context.Context) LinuxConfigurationResponsePtrOutput {
	return o
}

func (o LinuxConfigurationResponsePtrOutput) Elem() LinuxConfigurationResponseOutput {
	return o.ApplyT(func(v *LinuxConfigurationResponse) LinuxConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret LinuxConfigurationResponse
		return ret
	}).(LinuxConfigurationResponseOutput)
}

func (o LinuxConfigurationResponsePtrOutput) DisablePasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LinuxConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisablePasswordAuthentication
	}).(pulumi.BoolPtrOutput)
}

func (o LinuxConfigurationResponsePtrOutput) PatchSettings() LinuxPatchSettingsResponsePtrOutput {
	return o.ApplyT(func(v *LinuxConfigurationResponse) *LinuxPatchSettingsResponse {
		if v == nil {
			return nil
		}
		return v.PatchSettings
	}).(LinuxPatchSettingsResponsePtrOutput)
}

func (o LinuxConfigurationResponsePtrOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LinuxConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ProvisionVMAgent
	}).(pulumi.BoolPtrOutput)
}

func (o LinuxConfigurationResponsePtrOutput) Ssh() SshConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *LinuxConfigurationResponse) *SshConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Ssh
	}).(SshConfigurationResponsePtrOutput)
}

type LinuxPatchSettings struct {
	AssessmentMode              *string                                       `pulumi:"assessmentMode"`
	AutomaticByPlatformSettings *LinuxVMGuestPatchAutomaticByPlatformSettings `pulumi:"automaticByPlatformSettings"`
	PatchMode                   *string                                       `pulumi:"patchMode"`
}





type LinuxPatchSettingsInput interface {
	pulumi.Input

	ToLinuxPatchSettingsOutput() LinuxPatchSettingsOutput
	ToLinuxPatchSettingsOutputWithContext(context.Context) LinuxPatchSettingsOutput
}

type LinuxPatchSettingsArgs struct {
	AssessmentMode              pulumi.StringPtrInput                                `pulumi:"assessmentMode"`
	AutomaticByPlatformSettings LinuxVMGuestPatchAutomaticByPlatformSettingsPtrInput `pulumi:"automaticByPlatformSettings"`
	PatchMode                   pulumi.StringPtrInput                                `pulumi:"patchMode"`
}

func (LinuxPatchSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxPatchSettings)(nil)).Elem()
}

func (i LinuxPatchSettingsArgs) ToLinuxPatchSettingsOutput() LinuxPatchSettingsOutput {
	return i.ToLinuxPatchSettingsOutputWithContext(context.Background())
}

func (i LinuxPatchSettingsArgs) ToLinuxPatchSettingsOutputWithContext(ctx context.Context) LinuxPatchSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxPatchSettingsOutput)
}

func (i LinuxPatchSettingsArgs) ToLinuxPatchSettingsPtrOutput() LinuxPatchSettingsPtrOutput {
	return i.ToLinuxPatchSettingsPtrOutputWithContext(context.Background())
}

func (i LinuxPatchSettingsArgs) ToLinuxPatchSettingsPtrOutputWithContext(ctx context.Context) LinuxPatchSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxPatchSettingsOutput).ToLinuxPatchSettingsPtrOutputWithContext(ctx)
}









type LinuxPatchSettingsPtrInput interface {
	pulumi.Input

	ToLinuxPatchSettingsPtrOutput() LinuxPatchSettingsPtrOutput
	ToLinuxPatchSettingsPtrOutputWithContext(context.Context) LinuxPatchSettingsPtrOutput
}

type linuxPatchSettingsPtrType LinuxPatchSettingsArgs

func LinuxPatchSettingsPtr(v *LinuxPatchSettingsArgs) LinuxPatchSettingsPtrInput {
	return (*linuxPatchSettingsPtrType)(v)
}

func (*linuxPatchSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxPatchSettings)(nil)).Elem()
}

func (i *linuxPatchSettingsPtrType) ToLinuxPatchSettingsPtrOutput() LinuxPatchSettingsPtrOutput {
	return i.ToLinuxPatchSettingsPtrOutputWithContext(context.Background())
}

func (i *linuxPatchSettingsPtrType) ToLinuxPatchSettingsPtrOutputWithContext(ctx context.Context) LinuxPatchSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxPatchSettingsPtrOutput)
}

type LinuxPatchSettingsOutput struct{ *pulumi.OutputState }

func (LinuxPatchSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxPatchSettings)(nil)).Elem()
}

func (o LinuxPatchSettingsOutput) ToLinuxPatchSettingsOutput() LinuxPatchSettingsOutput {
	return o
}

func (o LinuxPatchSettingsOutput) ToLinuxPatchSettingsOutputWithContext(ctx context.Context) LinuxPatchSettingsOutput {
	return o
}

func (o LinuxPatchSettingsOutput) ToLinuxPatchSettingsPtrOutput() LinuxPatchSettingsPtrOutput {
	return o.ToLinuxPatchSettingsPtrOutputWithContext(context.Background())
}

func (o LinuxPatchSettingsOutput) ToLinuxPatchSettingsPtrOutputWithContext(ctx context.Context) LinuxPatchSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinuxPatchSettings) *LinuxPatchSettings {
		return &v
	}).(LinuxPatchSettingsPtrOutput)
}

func (o LinuxPatchSettingsOutput) AssessmentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxPatchSettings) *string { return v.AssessmentMode }).(pulumi.StringPtrOutput)
}

func (o LinuxPatchSettingsOutput) AutomaticByPlatformSettings() LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return o.ApplyT(func(v LinuxPatchSettings) *LinuxVMGuestPatchAutomaticByPlatformSettings {
		return v.AutomaticByPlatformSettings
	}).(LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput)
}

func (o LinuxPatchSettingsOutput) PatchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxPatchSettings) *string { return v.PatchMode }).(pulumi.StringPtrOutput)
}

type LinuxPatchSettingsPtrOutput struct{ *pulumi.OutputState }

func (LinuxPatchSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxPatchSettings)(nil)).Elem()
}

func (o LinuxPatchSettingsPtrOutput) ToLinuxPatchSettingsPtrOutput() LinuxPatchSettingsPtrOutput {
	return o
}

func (o LinuxPatchSettingsPtrOutput) ToLinuxPatchSettingsPtrOutputWithContext(ctx context.Context) LinuxPatchSettingsPtrOutput {
	return o
}

func (o LinuxPatchSettingsPtrOutput) Elem() LinuxPatchSettingsOutput {
	return o.ApplyT(func(v *LinuxPatchSettings) LinuxPatchSettings {
		if v != nil {
			return *v
		}
		var ret LinuxPatchSettings
		return ret
	}).(LinuxPatchSettingsOutput)
}

func (o LinuxPatchSettingsPtrOutput) AssessmentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxPatchSettings) *string {
		if v == nil {
			return nil
		}
		return v.AssessmentMode
	}).(pulumi.StringPtrOutput)
}

func (o LinuxPatchSettingsPtrOutput) AutomaticByPlatformSettings() LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return o.ApplyT(func(v *LinuxPatchSettings) *LinuxVMGuestPatchAutomaticByPlatformSettings {
		if v == nil {
			return nil
		}
		return v.AutomaticByPlatformSettings
	}).(LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput)
}

func (o LinuxPatchSettingsPtrOutput) PatchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxPatchSettings) *string {
		if v == nil {
			return nil
		}
		return v.PatchMode
	}).(pulumi.StringPtrOutput)
}

type LinuxPatchSettingsResponse struct {
	AssessmentMode              *string                                               `pulumi:"assessmentMode"`
	AutomaticByPlatformSettings *LinuxVMGuestPatchAutomaticByPlatformSettingsResponse `pulumi:"automaticByPlatformSettings"`
	PatchMode                   *string                                               `pulumi:"patchMode"`
}

type LinuxPatchSettingsResponseOutput struct{ *pulumi.OutputState }

func (LinuxPatchSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxPatchSettingsResponse)(nil)).Elem()
}

func (o LinuxPatchSettingsResponseOutput) ToLinuxPatchSettingsResponseOutput() LinuxPatchSettingsResponseOutput {
	return o
}

func (o LinuxPatchSettingsResponseOutput) ToLinuxPatchSettingsResponseOutputWithContext(ctx context.Context) LinuxPatchSettingsResponseOutput {
	return o
}

func (o LinuxPatchSettingsResponseOutput) AssessmentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxPatchSettingsResponse) *string { return v.AssessmentMode }).(pulumi.StringPtrOutput)
}

func (o LinuxPatchSettingsResponseOutput) AutomaticByPlatformSettings() LinuxVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput {
	return o.ApplyT(func(v LinuxPatchSettingsResponse) *LinuxVMGuestPatchAutomaticByPlatformSettingsResponse {
		return v.AutomaticByPlatformSettings
	}).(LinuxVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput)
}

func (o LinuxPatchSettingsResponseOutput) PatchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxPatchSettingsResponse) *string { return v.PatchMode }).(pulumi.StringPtrOutput)
}

type LinuxPatchSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (LinuxPatchSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxPatchSettingsResponse)(nil)).Elem()
}

func (o LinuxPatchSettingsResponsePtrOutput) ToLinuxPatchSettingsResponsePtrOutput() LinuxPatchSettingsResponsePtrOutput {
	return o
}

func (o LinuxPatchSettingsResponsePtrOutput) ToLinuxPatchSettingsResponsePtrOutputWithContext(ctx context.Context) LinuxPatchSettingsResponsePtrOutput {
	return o
}

func (o LinuxPatchSettingsResponsePtrOutput) Elem() LinuxPatchSettingsResponseOutput {
	return o.ApplyT(func(v *LinuxPatchSettingsResponse) LinuxPatchSettingsResponse {
		if v != nil {
			return *v
		}
		var ret LinuxPatchSettingsResponse
		return ret
	}).(LinuxPatchSettingsResponseOutput)
}

func (o LinuxPatchSettingsResponsePtrOutput) AssessmentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxPatchSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AssessmentMode
	}).(pulumi.StringPtrOutput)
}

func (o LinuxPatchSettingsResponsePtrOutput) AutomaticByPlatformSettings() LinuxVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput {
	return o.ApplyT(func(v *LinuxPatchSettingsResponse) *LinuxVMGuestPatchAutomaticByPlatformSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutomaticByPlatformSettings
	}).(LinuxVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput)
}

func (o LinuxPatchSettingsResponsePtrOutput) PatchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxPatchSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.PatchMode
	}).(pulumi.StringPtrOutput)
}

type LinuxVMGuestPatchAutomaticByPlatformSettings struct {
	RebootSetting *string `pulumi:"rebootSetting"`
}





type LinuxVMGuestPatchAutomaticByPlatformSettingsInput interface {
	pulumi.Input

	ToLinuxVMGuestPatchAutomaticByPlatformSettingsOutput() LinuxVMGuestPatchAutomaticByPlatformSettingsOutput
	ToLinuxVMGuestPatchAutomaticByPlatformSettingsOutputWithContext(context.Context) LinuxVMGuestPatchAutomaticByPlatformSettingsOutput
}

type LinuxVMGuestPatchAutomaticByPlatformSettingsArgs struct {
	RebootSetting pulumi.StringPtrInput `pulumi:"rebootSetting"`
}

func (LinuxVMGuestPatchAutomaticByPlatformSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxVMGuestPatchAutomaticByPlatformSettings)(nil)).Elem()
}

func (i LinuxVMGuestPatchAutomaticByPlatformSettingsArgs) ToLinuxVMGuestPatchAutomaticByPlatformSettingsOutput() LinuxVMGuestPatchAutomaticByPlatformSettingsOutput {
	return i.ToLinuxVMGuestPatchAutomaticByPlatformSettingsOutputWithContext(context.Background())
}

func (i LinuxVMGuestPatchAutomaticByPlatformSettingsArgs) ToLinuxVMGuestPatchAutomaticByPlatformSettingsOutputWithContext(ctx context.Context) LinuxVMGuestPatchAutomaticByPlatformSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxVMGuestPatchAutomaticByPlatformSettingsOutput)
}

func (i LinuxVMGuestPatchAutomaticByPlatformSettingsArgs) ToLinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput() LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return i.ToLinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutputWithContext(context.Background())
}

func (i LinuxVMGuestPatchAutomaticByPlatformSettingsArgs) ToLinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutputWithContext(ctx context.Context) LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxVMGuestPatchAutomaticByPlatformSettingsOutput).ToLinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutputWithContext(ctx)
}









type LinuxVMGuestPatchAutomaticByPlatformSettingsPtrInput interface {
	pulumi.Input

	ToLinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput() LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput
	ToLinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutputWithContext(context.Context) LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput
}

type linuxVMGuestPatchAutomaticByPlatformSettingsPtrType LinuxVMGuestPatchAutomaticByPlatformSettingsArgs

func LinuxVMGuestPatchAutomaticByPlatformSettingsPtr(v *LinuxVMGuestPatchAutomaticByPlatformSettingsArgs) LinuxVMGuestPatchAutomaticByPlatformSettingsPtrInput {
	return (*linuxVMGuestPatchAutomaticByPlatformSettingsPtrType)(v)
}

func (*linuxVMGuestPatchAutomaticByPlatformSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxVMGuestPatchAutomaticByPlatformSettings)(nil)).Elem()
}

func (i *linuxVMGuestPatchAutomaticByPlatformSettingsPtrType) ToLinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput() LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return i.ToLinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutputWithContext(context.Background())
}

func (i *linuxVMGuestPatchAutomaticByPlatformSettingsPtrType) ToLinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutputWithContext(ctx context.Context) LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput)
}

type LinuxVMGuestPatchAutomaticByPlatformSettingsOutput struct{ *pulumi.OutputState }

func (LinuxVMGuestPatchAutomaticByPlatformSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxVMGuestPatchAutomaticByPlatformSettings)(nil)).Elem()
}

func (o LinuxVMGuestPatchAutomaticByPlatformSettingsOutput) ToLinuxVMGuestPatchAutomaticByPlatformSettingsOutput() LinuxVMGuestPatchAutomaticByPlatformSettingsOutput {
	return o
}

func (o LinuxVMGuestPatchAutomaticByPlatformSettingsOutput) ToLinuxVMGuestPatchAutomaticByPlatformSettingsOutputWithContext(ctx context.Context) LinuxVMGuestPatchAutomaticByPlatformSettingsOutput {
	return o
}

func (o LinuxVMGuestPatchAutomaticByPlatformSettingsOutput) ToLinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput() LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return o.ToLinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutputWithContext(context.Background())
}

func (o LinuxVMGuestPatchAutomaticByPlatformSettingsOutput) ToLinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutputWithContext(ctx context.Context) LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinuxVMGuestPatchAutomaticByPlatformSettings) *LinuxVMGuestPatchAutomaticByPlatformSettings {
		return &v
	}).(LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput)
}

func (o LinuxVMGuestPatchAutomaticByPlatformSettingsOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxVMGuestPatchAutomaticByPlatformSettings) *string { return v.RebootSetting }).(pulumi.StringPtrOutput)
}

type LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput struct{ *pulumi.OutputState }

func (LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxVMGuestPatchAutomaticByPlatformSettings)(nil)).Elem()
}

func (o LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput) ToLinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput() LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return o
}

func (o LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput) ToLinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutputWithContext(ctx context.Context) LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return o
}

func (o LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput) Elem() LinuxVMGuestPatchAutomaticByPlatformSettingsOutput {
	return o.ApplyT(func(v *LinuxVMGuestPatchAutomaticByPlatformSettings) LinuxVMGuestPatchAutomaticByPlatformSettings {
		if v != nil {
			return *v
		}
		var ret LinuxVMGuestPatchAutomaticByPlatformSettings
		return ret
	}).(LinuxVMGuestPatchAutomaticByPlatformSettingsOutput)
}

func (o LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxVMGuestPatchAutomaticByPlatformSettings) *string {
		if v == nil {
			return nil
		}
		return v.RebootSetting
	}).(pulumi.StringPtrOutput)
}

type LinuxVMGuestPatchAutomaticByPlatformSettingsResponse struct {
	RebootSetting *string `pulumi:"rebootSetting"`
}

type LinuxVMGuestPatchAutomaticByPlatformSettingsResponseOutput struct{ *pulumi.OutputState }

func (LinuxVMGuestPatchAutomaticByPlatformSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxVMGuestPatchAutomaticByPlatformSettingsResponse)(nil)).Elem()
}

func (o LinuxVMGuestPatchAutomaticByPlatformSettingsResponseOutput) ToLinuxVMGuestPatchAutomaticByPlatformSettingsResponseOutput() LinuxVMGuestPatchAutomaticByPlatformSettingsResponseOutput {
	return o
}

func (o LinuxVMGuestPatchAutomaticByPlatformSettingsResponseOutput) ToLinuxVMGuestPatchAutomaticByPlatformSettingsResponseOutputWithContext(ctx context.Context) LinuxVMGuestPatchAutomaticByPlatformSettingsResponseOutput {
	return o
}

func (o LinuxVMGuestPatchAutomaticByPlatformSettingsResponseOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxVMGuestPatchAutomaticByPlatformSettingsResponse) *string { return v.RebootSetting }).(pulumi.StringPtrOutput)
}

type LinuxVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (LinuxVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxVMGuestPatchAutomaticByPlatformSettingsResponse)(nil)).Elem()
}

func (o LinuxVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput) ToLinuxVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput() LinuxVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput {
	return o
}

func (o LinuxVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput) ToLinuxVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutputWithContext(ctx context.Context) LinuxVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput {
	return o
}

func (o LinuxVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput) Elem() LinuxVMGuestPatchAutomaticByPlatformSettingsResponseOutput {
	return o.ApplyT(func(v *LinuxVMGuestPatchAutomaticByPlatformSettingsResponse) LinuxVMGuestPatchAutomaticByPlatformSettingsResponse {
		if v != nil {
			return *v
		}
		var ret LinuxVMGuestPatchAutomaticByPlatformSettingsResponse
		return ret
	}).(LinuxVMGuestPatchAutomaticByPlatformSettingsResponseOutput)
}

func (o LinuxVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxVMGuestPatchAutomaticByPlatformSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.RebootSetting
	}).(pulumi.StringPtrOutput)
}

type LogAnalyticsOutputResponse struct {
	Output string `pulumi:"output"`
}

type LogAnalyticsOutputResponseOutput struct{ *pulumi.OutputState }

func (LogAnalyticsOutputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsOutputResponse)(nil)).Elem()
}

func (o LogAnalyticsOutputResponseOutput) ToLogAnalyticsOutputResponseOutput() LogAnalyticsOutputResponseOutput {
	return o
}

func (o LogAnalyticsOutputResponseOutput) ToLogAnalyticsOutputResponseOutputWithContext(ctx context.Context) LogAnalyticsOutputResponseOutput {
	return o
}

func (o LogAnalyticsOutputResponseOutput) Output() pulumi.StringOutput {
	return o.ApplyT(func(v LogAnalyticsOutputResponse) string { return v.Output }).(pulumi.StringOutput)
}

type MaintenanceRedeployStatusResponse struct {
	IsCustomerInitiatedMaintenanceAllowed *bool   `pulumi:"isCustomerInitiatedMaintenanceAllowed"`
	LastOperationMessage                  *string `pulumi:"lastOperationMessage"`
	LastOperationResultCode               *string `pulumi:"lastOperationResultCode"`
	MaintenanceWindowEndTime              *string `pulumi:"maintenanceWindowEndTime"`
	MaintenanceWindowStartTime            *string `pulumi:"maintenanceWindowStartTime"`
	PreMaintenanceWindowEndTime           *string `pulumi:"preMaintenanceWindowEndTime"`
	PreMaintenanceWindowStartTime         *string `pulumi:"preMaintenanceWindowStartTime"`
}

type MaintenanceRedeployStatusResponseOutput struct{ *pulumi.OutputState }

func (MaintenanceRedeployStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceRedeployStatusResponse)(nil)).Elem()
}

func (o MaintenanceRedeployStatusResponseOutput) ToMaintenanceRedeployStatusResponseOutput() MaintenanceRedeployStatusResponseOutput {
	return o
}

func (o MaintenanceRedeployStatusResponseOutput) ToMaintenanceRedeployStatusResponseOutputWithContext(ctx context.Context) MaintenanceRedeployStatusResponseOutput {
	return o
}

func (o MaintenanceRedeployStatusResponseOutput) IsCustomerInitiatedMaintenanceAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MaintenanceRedeployStatusResponse) *bool { return v.IsCustomerInitiatedMaintenanceAllowed }).(pulumi.BoolPtrOutput)
}

func (o MaintenanceRedeployStatusResponseOutput) LastOperationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceRedeployStatusResponse) *string { return v.LastOperationMessage }).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponseOutput) LastOperationResultCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceRedeployStatusResponse) *string { return v.LastOperationResultCode }).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponseOutput) MaintenanceWindowEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceRedeployStatusResponse) *string { return v.MaintenanceWindowEndTime }).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponseOutput) MaintenanceWindowStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceRedeployStatusResponse) *string { return v.MaintenanceWindowStartTime }).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponseOutput) PreMaintenanceWindowEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceRedeployStatusResponse) *string { return v.PreMaintenanceWindowEndTime }).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponseOutput) PreMaintenanceWindowStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MaintenanceRedeployStatusResponse) *string { return v.PreMaintenanceWindowStartTime }).(pulumi.StringPtrOutput)
}

type MaintenanceRedeployStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (MaintenanceRedeployStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceRedeployStatusResponse)(nil)).Elem()
}

func (o MaintenanceRedeployStatusResponsePtrOutput) ToMaintenanceRedeployStatusResponsePtrOutput() MaintenanceRedeployStatusResponsePtrOutput {
	return o
}

func (o MaintenanceRedeployStatusResponsePtrOutput) ToMaintenanceRedeployStatusResponsePtrOutputWithContext(ctx context.Context) MaintenanceRedeployStatusResponsePtrOutput {
	return o
}

func (o MaintenanceRedeployStatusResponsePtrOutput) Elem() MaintenanceRedeployStatusResponseOutput {
	return o.ApplyT(func(v *MaintenanceRedeployStatusResponse) MaintenanceRedeployStatusResponse {
		if v != nil {
			return *v
		}
		var ret MaintenanceRedeployStatusResponse
		return ret
	}).(MaintenanceRedeployStatusResponseOutput)
}

func (o MaintenanceRedeployStatusResponsePtrOutput) IsCustomerInitiatedMaintenanceAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MaintenanceRedeployStatusResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsCustomerInitiatedMaintenanceAllowed
	}).(pulumi.BoolPtrOutput)
}

func (o MaintenanceRedeployStatusResponsePtrOutput) LastOperationMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceRedeployStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastOperationMessage
	}).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponsePtrOutput) LastOperationResultCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceRedeployStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastOperationResultCode
	}).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponsePtrOutput) MaintenanceWindowEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceRedeployStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.MaintenanceWindowEndTime
	}).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponsePtrOutput) MaintenanceWindowStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceRedeployStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.MaintenanceWindowStartTime
	}).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponsePtrOutput) PreMaintenanceWindowEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceRedeployStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.PreMaintenanceWindowEndTime
	}).(pulumi.StringPtrOutput)
}

func (o MaintenanceRedeployStatusResponsePtrOutput) PreMaintenanceWindowStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceRedeployStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.PreMaintenanceWindowStartTime
	}).(pulumi.StringPtrOutput)
}

type ManagedDiskParameters struct {
	DiskEncryptionSet  *DiskEncryptionSetParameters `pulumi:"diskEncryptionSet"`
	Id                 *string                      `pulumi:"id"`
	SecurityProfile    *VMDiskSecurityProfile       `pulumi:"securityProfile"`
	StorageAccountType *string                      `pulumi:"storageAccountType"`
}





type ManagedDiskParametersInput interface {
	pulumi.Input

	ToManagedDiskParametersOutput() ManagedDiskParametersOutput
	ToManagedDiskParametersOutputWithContext(context.Context) ManagedDiskParametersOutput
}

type ManagedDiskParametersArgs struct {
	DiskEncryptionSet  DiskEncryptionSetParametersPtrInput `pulumi:"diskEncryptionSet"`
	Id                 pulumi.StringPtrInput               `pulumi:"id"`
	SecurityProfile    VMDiskSecurityProfilePtrInput       `pulumi:"securityProfile"`
	StorageAccountType pulumi.StringPtrInput               `pulumi:"storageAccountType"`
}

func (ManagedDiskParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedDiskParameters)(nil)).Elem()
}

func (i ManagedDiskParametersArgs) ToManagedDiskParametersOutput() ManagedDiskParametersOutput {
	return i.ToManagedDiskParametersOutputWithContext(context.Background())
}

func (i ManagedDiskParametersArgs) ToManagedDiskParametersOutputWithContext(ctx context.Context) ManagedDiskParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedDiskParametersOutput)
}

func (i ManagedDiskParametersArgs) ToManagedDiskParametersPtrOutput() ManagedDiskParametersPtrOutput {
	return i.ToManagedDiskParametersPtrOutputWithContext(context.Background())
}

func (i ManagedDiskParametersArgs) ToManagedDiskParametersPtrOutputWithContext(ctx context.Context) ManagedDiskParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedDiskParametersOutput).ToManagedDiskParametersPtrOutputWithContext(ctx)
}









type ManagedDiskParametersPtrInput interface {
	pulumi.Input

	ToManagedDiskParametersPtrOutput() ManagedDiskParametersPtrOutput
	ToManagedDiskParametersPtrOutputWithContext(context.Context) ManagedDiskParametersPtrOutput
}

type managedDiskParametersPtrType ManagedDiskParametersArgs

func ManagedDiskParametersPtr(v *ManagedDiskParametersArgs) ManagedDiskParametersPtrInput {
	return (*managedDiskParametersPtrType)(v)
}

func (*managedDiskParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedDiskParameters)(nil)).Elem()
}

func (i *managedDiskParametersPtrType) ToManagedDiskParametersPtrOutput() ManagedDiskParametersPtrOutput {
	return i.ToManagedDiskParametersPtrOutputWithContext(context.Background())
}

func (i *managedDiskParametersPtrType) ToManagedDiskParametersPtrOutputWithContext(ctx context.Context) ManagedDiskParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedDiskParametersPtrOutput)
}

type ManagedDiskParametersOutput struct{ *pulumi.OutputState }

func (ManagedDiskParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedDiskParameters)(nil)).Elem()
}

func (o ManagedDiskParametersOutput) ToManagedDiskParametersOutput() ManagedDiskParametersOutput {
	return o
}

func (o ManagedDiskParametersOutput) ToManagedDiskParametersOutputWithContext(ctx context.Context) ManagedDiskParametersOutput {
	return o
}

func (o ManagedDiskParametersOutput) ToManagedDiskParametersPtrOutput() ManagedDiskParametersPtrOutput {
	return o.ToManagedDiskParametersPtrOutputWithContext(context.Background())
}

func (o ManagedDiskParametersOutput) ToManagedDiskParametersPtrOutputWithContext(ctx context.Context) ManagedDiskParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedDiskParameters) *ManagedDiskParameters {
		return &v
	}).(ManagedDiskParametersPtrOutput)
}

func (o ManagedDiskParametersOutput) DiskEncryptionSet() DiskEncryptionSetParametersPtrOutput {
	return o.ApplyT(func(v ManagedDiskParameters) *DiskEncryptionSetParameters { return v.DiskEncryptionSet }).(DiskEncryptionSetParametersPtrOutput)
}

func (o ManagedDiskParametersOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedDiskParameters) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ManagedDiskParametersOutput) SecurityProfile() VMDiskSecurityProfilePtrOutput {
	return o.ApplyT(func(v ManagedDiskParameters) *VMDiskSecurityProfile { return v.SecurityProfile }).(VMDiskSecurityProfilePtrOutput)
}

func (o ManagedDiskParametersOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedDiskParameters) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type ManagedDiskParametersPtrOutput struct{ *pulumi.OutputState }

func (ManagedDiskParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedDiskParameters)(nil)).Elem()
}

func (o ManagedDiskParametersPtrOutput) ToManagedDiskParametersPtrOutput() ManagedDiskParametersPtrOutput {
	return o
}

func (o ManagedDiskParametersPtrOutput) ToManagedDiskParametersPtrOutputWithContext(ctx context.Context) ManagedDiskParametersPtrOutput {
	return o
}

func (o ManagedDiskParametersPtrOutput) Elem() ManagedDiskParametersOutput {
	return o.ApplyT(func(v *ManagedDiskParameters) ManagedDiskParameters {
		if v != nil {
			return *v
		}
		var ret ManagedDiskParameters
		return ret
	}).(ManagedDiskParametersOutput)
}

func (o ManagedDiskParametersPtrOutput) DiskEncryptionSet() DiskEncryptionSetParametersPtrOutput {
	return o.ApplyT(func(v *ManagedDiskParameters) *DiskEncryptionSetParameters {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionSet
	}).(DiskEncryptionSetParametersPtrOutput)
}

func (o ManagedDiskParametersPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDiskParameters) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ManagedDiskParametersPtrOutput) SecurityProfile() VMDiskSecurityProfilePtrOutput {
	return o.ApplyT(func(v *ManagedDiskParameters) *VMDiskSecurityProfile {
		if v == nil {
			return nil
		}
		return v.SecurityProfile
	}).(VMDiskSecurityProfilePtrOutput)
}

func (o ManagedDiskParametersPtrOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDiskParameters) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountType
	}).(pulumi.StringPtrOutput)
}

type ManagedDiskParametersResponse struct {
	DiskEncryptionSet  *DiskEncryptionSetParametersResponse `pulumi:"diskEncryptionSet"`
	Id                 *string                              `pulumi:"id"`
	SecurityProfile    *VMDiskSecurityProfileResponse       `pulumi:"securityProfile"`
	StorageAccountType *string                              `pulumi:"storageAccountType"`
}

type ManagedDiskParametersResponseOutput struct{ *pulumi.OutputState }

func (ManagedDiskParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedDiskParametersResponse)(nil)).Elem()
}

func (o ManagedDiskParametersResponseOutput) ToManagedDiskParametersResponseOutput() ManagedDiskParametersResponseOutput {
	return o
}

func (o ManagedDiskParametersResponseOutput) ToManagedDiskParametersResponseOutputWithContext(ctx context.Context) ManagedDiskParametersResponseOutput {
	return o
}

func (o ManagedDiskParametersResponseOutput) DiskEncryptionSet() DiskEncryptionSetParametersResponsePtrOutput {
	return o.ApplyT(func(v ManagedDiskParametersResponse) *DiskEncryptionSetParametersResponse { return v.DiskEncryptionSet }).(DiskEncryptionSetParametersResponsePtrOutput)
}

func (o ManagedDiskParametersResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedDiskParametersResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ManagedDiskParametersResponseOutput) SecurityProfile() VMDiskSecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v ManagedDiskParametersResponse) *VMDiskSecurityProfileResponse { return v.SecurityProfile }).(VMDiskSecurityProfileResponsePtrOutput)
}

func (o ManagedDiskParametersResponseOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedDiskParametersResponse) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type ManagedDiskParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedDiskParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedDiskParametersResponse)(nil)).Elem()
}

func (o ManagedDiskParametersResponsePtrOutput) ToManagedDiskParametersResponsePtrOutput() ManagedDiskParametersResponsePtrOutput {
	return o
}

func (o ManagedDiskParametersResponsePtrOutput) ToManagedDiskParametersResponsePtrOutputWithContext(ctx context.Context) ManagedDiskParametersResponsePtrOutput {
	return o
}

func (o ManagedDiskParametersResponsePtrOutput) Elem() ManagedDiskParametersResponseOutput {
	return o.ApplyT(func(v *ManagedDiskParametersResponse) ManagedDiskParametersResponse {
		if v != nil {
			return *v
		}
		var ret ManagedDiskParametersResponse
		return ret
	}).(ManagedDiskParametersResponseOutput)
}

func (o ManagedDiskParametersResponsePtrOutput) DiskEncryptionSet() DiskEncryptionSetParametersResponsePtrOutput {
	return o.ApplyT(func(v *ManagedDiskParametersResponse) *DiskEncryptionSetParametersResponse {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionSet
	}).(DiskEncryptionSetParametersResponsePtrOutput)
}

func (o ManagedDiskParametersResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDiskParametersResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ManagedDiskParametersResponsePtrOutput) SecurityProfile() VMDiskSecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v *ManagedDiskParametersResponse) *VMDiskSecurityProfileResponse {
		if v == nil {
			return nil
		}
		return v.SecurityProfile
	}).(VMDiskSecurityProfileResponsePtrOutput)
}

func (o ManagedDiskParametersResponsePtrOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedDiskParametersResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountType
	}).(pulumi.StringPtrOutput)
}

type NetworkInterfaceReference struct {
	DeleteOption *string `pulumi:"deleteOption"`
	Id           *string `pulumi:"id"`
	Primary      *bool   `pulumi:"primary"`
}





type NetworkInterfaceReferenceInput interface {
	pulumi.Input

	ToNetworkInterfaceReferenceOutput() NetworkInterfaceReferenceOutput
	ToNetworkInterfaceReferenceOutputWithContext(context.Context) NetworkInterfaceReferenceOutput
}

type NetworkInterfaceReferenceArgs struct {
	DeleteOption pulumi.StringPtrInput `pulumi:"deleteOption"`
	Id           pulumi.StringPtrInput `pulumi:"id"`
	Primary      pulumi.BoolPtrInput   `pulumi:"primary"`
}

func (NetworkInterfaceReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceReference)(nil)).Elem()
}

func (i NetworkInterfaceReferenceArgs) ToNetworkInterfaceReferenceOutput() NetworkInterfaceReferenceOutput {
	return i.ToNetworkInterfaceReferenceOutputWithContext(context.Background())
}

func (i NetworkInterfaceReferenceArgs) ToNetworkInterfaceReferenceOutputWithContext(ctx context.Context) NetworkInterfaceReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceReferenceOutput)
}





type NetworkInterfaceReferenceArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceReferenceArrayOutput() NetworkInterfaceReferenceArrayOutput
	ToNetworkInterfaceReferenceArrayOutputWithContext(context.Context) NetworkInterfaceReferenceArrayOutput
}

type NetworkInterfaceReferenceArray []NetworkInterfaceReferenceInput

func (NetworkInterfaceReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceReference)(nil)).Elem()
}

func (i NetworkInterfaceReferenceArray) ToNetworkInterfaceReferenceArrayOutput() NetworkInterfaceReferenceArrayOutput {
	return i.ToNetworkInterfaceReferenceArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceReferenceArray) ToNetworkInterfaceReferenceArrayOutputWithContext(ctx context.Context) NetworkInterfaceReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceReferenceArrayOutput)
}

type NetworkInterfaceReferenceOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceReference)(nil)).Elem()
}

func (o NetworkInterfaceReferenceOutput) ToNetworkInterfaceReferenceOutput() NetworkInterfaceReferenceOutput {
	return o
}

func (o NetworkInterfaceReferenceOutput) ToNetworkInterfaceReferenceOutputWithContext(ctx context.Context) NetworkInterfaceReferenceOutput {
	return o
}

func (o NetworkInterfaceReferenceOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceReference) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceReferenceOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceReference) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

type NetworkInterfaceReferenceArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceReference)(nil)).Elem()
}

func (o NetworkInterfaceReferenceArrayOutput) ToNetworkInterfaceReferenceArrayOutput() NetworkInterfaceReferenceArrayOutput {
	return o
}

func (o NetworkInterfaceReferenceArrayOutput) ToNetworkInterfaceReferenceArrayOutputWithContext(ctx context.Context) NetworkInterfaceReferenceArrayOutput {
	return o
}

func (o NetworkInterfaceReferenceArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceReference {
		return vs[0].([]NetworkInterfaceReference)[vs[1].(int)]
	}).(NetworkInterfaceReferenceOutput)
}

type NetworkInterfaceReferenceResponse struct {
	DeleteOption *string `pulumi:"deleteOption"`
	Id           *string `pulumi:"id"`
	Primary      *bool   `pulumi:"primary"`
}

type NetworkInterfaceReferenceResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceReferenceResponse)(nil)).Elem()
}

func (o NetworkInterfaceReferenceResponseOutput) ToNetworkInterfaceReferenceResponseOutput() NetworkInterfaceReferenceResponseOutput {
	return o
}

func (o NetworkInterfaceReferenceResponseOutput) ToNetworkInterfaceReferenceResponseOutputWithContext(ctx context.Context) NetworkInterfaceReferenceResponseOutput {
	return o
}

func (o NetworkInterfaceReferenceResponseOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceReferenceResponse) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceReferenceResponseOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceReferenceResponse) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

type NetworkInterfaceReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceReferenceResponse)(nil)).Elem()
}

func (o NetworkInterfaceReferenceResponseArrayOutput) ToNetworkInterfaceReferenceResponseArrayOutput() NetworkInterfaceReferenceResponseArrayOutput {
	return o
}

func (o NetworkInterfaceReferenceResponseArrayOutput) ToNetworkInterfaceReferenceResponseArrayOutputWithContext(ctx context.Context) NetworkInterfaceReferenceResponseArrayOutput {
	return o
}

func (o NetworkInterfaceReferenceResponseArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceReferenceResponse {
		return vs[0].([]NetworkInterfaceReferenceResponse)[vs[1].(int)]
	}).(NetworkInterfaceReferenceResponseOutput)
}

type NetworkProfile struct {
	NetworkApiVersion              *string                                       `pulumi:"networkApiVersion"`
	NetworkInterfaceConfigurations []VirtualMachineNetworkInterfaceConfiguration `pulumi:"networkInterfaceConfigurations"`
	NetworkInterfaces              []NetworkInterfaceReference                   `pulumi:"networkInterfaces"`
}





type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(context.Context) NetworkProfileOutput
}

type NetworkProfileArgs struct {
	NetworkApiVersion              pulumi.StringPtrInput                                 `pulumi:"networkApiVersion"`
	NetworkInterfaceConfigurations VirtualMachineNetworkInterfaceConfigurationArrayInput `pulumi:"networkInterfaceConfigurations"`
	NetworkInterfaces              NetworkInterfaceReferenceArrayInput                   `pulumi:"networkInterfaces"`
}

func (NetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (i NetworkProfileArgs) ToNetworkProfileOutput() NetworkProfileOutput {
	return i.ToNetworkProfileOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput)
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput).ToNetworkProfilePtrOutputWithContext(ctx)
}









type NetworkProfilePtrInput interface {
	pulumi.Input

	ToNetworkProfilePtrOutput() NetworkProfilePtrOutput
	ToNetworkProfilePtrOutputWithContext(context.Context) NetworkProfilePtrOutput
}

type networkProfilePtrType NetworkProfileArgs

func NetworkProfilePtr(v *NetworkProfileArgs) NetworkProfilePtrInput {
	return (*networkProfilePtrType)(v)
}

func (*networkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfilePtrOutput)
}

type NetworkProfileOutput struct{ *pulumi.OutputState }

func (NetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileOutput) ToNetworkProfileOutput() NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkProfile) *NetworkProfile {
		return &v
	}).(NetworkProfilePtrOutput)
}

func (o NetworkProfileOutput) NetworkApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfile) *string { return v.NetworkApiVersion }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileOutput) NetworkInterfaceConfigurations() VirtualMachineNetworkInterfaceConfigurationArrayOutput {
	return o.ApplyT(func(v NetworkProfile) []VirtualMachineNetworkInterfaceConfiguration {
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineNetworkInterfaceConfigurationArrayOutput)
}

func (o NetworkProfileOutput) NetworkInterfaces() NetworkInterfaceReferenceArrayOutput {
	return o.ApplyT(func(v NetworkProfile) []NetworkInterfaceReference { return v.NetworkInterfaces }).(NetworkInterfaceReferenceArrayOutput)
}

type NetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) Elem() NetworkProfileOutput {
	return o.ApplyT(func(v *NetworkProfile) NetworkProfile {
		if v != nil {
			return *v
		}
		var ret NetworkProfile
		return ret
	}).(NetworkProfileOutput)
}

func (o NetworkProfilePtrOutput) NetworkApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.NetworkApiVersion
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfilePtrOutput) NetworkInterfaceConfigurations() VirtualMachineNetworkInterfaceConfigurationArrayOutput {
	return o.ApplyT(func(v *NetworkProfile) []VirtualMachineNetworkInterfaceConfiguration {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineNetworkInterfaceConfigurationArrayOutput)
}

func (o NetworkProfilePtrOutput) NetworkInterfaces() NetworkInterfaceReferenceArrayOutput {
	return o.ApplyT(func(v *NetworkProfile) []NetworkInterfaceReference {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(NetworkInterfaceReferenceArrayOutput)
}

type NetworkProfileResponse struct {
	NetworkApiVersion              *string                                               `pulumi:"networkApiVersion"`
	NetworkInterfaceConfigurations []VirtualMachineNetworkInterfaceConfigurationResponse `pulumi:"networkInterfaceConfigurations"`
	NetworkInterfaces              []NetworkInterfaceReferenceResponse                   `pulumi:"networkInterfaces"`
}

type NetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutput() NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutputWithContext(ctx context.Context) NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) NetworkApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *string { return v.NetworkApiVersion }).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponseOutput) NetworkInterfaceConfigurations() VirtualMachineNetworkInterfaceConfigurationResponseArrayOutput {
	return o.ApplyT(func(v NetworkProfileResponse) []VirtualMachineNetworkInterfaceConfigurationResponse {
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineNetworkInterfaceConfigurationResponseArrayOutput)
}

func (o NetworkProfileResponseOutput) NetworkInterfaces() NetworkInterfaceReferenceResponseArrayOutput {
	return o.ApplyT(func(v NetworkProfileResponse) []NetworkInterfaceReferenceResponse { return v.NetworkInterfaces }).(NetworkInterfaceReferenceResponseArrayOutput)
}

type NetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) Elem() NetworkProfileResponseOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) NetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret NetworkProfileResponse
		return ret
	}).(NetworkProfileResponseOutput)
}

func (o NetworkProfileResponsePtrOutput) NetworkApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.NetworkApiVersion
	}).(pulumi.StringPtrOutput)
}

func (o NetworkProfileResponsePtrOutput) NetworkInterfaceConfigurations() VirtualMachineNetworkInterfaceConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) []VirtualMachineNetworkInterfaceConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineNetworkInterfaceConfigurationResponseArrayOutput)
}

func (o NetworkProfileResponsePtrOutput) NetworkInterfaces() NetworkInterfaceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) []NetworkInterfaceReferenceResponse {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(NetworkInterfaceReferenceResponseArrayOutput)
}

type OSDisk struct {
	Caching                 *CachingTypes           `pulumi:"caching"`
	CreateOption            string                  `pulumi:"createOption"`
	DeleteOption            *string                 `pulumi:"deleteOption"`
	DiffDiskSettings        *DiffDiskSettings       `pulumi:"diffDiskSettings"`
	DiskSizeGB              *int                    `pulumi:"diskSizeGB"`
	EncryptionSettings      *DiskEncryptionSettings `pulumi:"encryptionSettings"`
	Image                   *VirtualHardDisk        `pulumi:"image"`
	ManagedDisk             *ManagedDiskParameters  `pulumi:"managedDisk"`
	Name                    *string                 `pulumi:"name"`
	OsType                  *OperatingSystemTypes   `pulumi:"osType"`
	Vhd                     *VirtualHardDisk        `pulumi:"vhd"`
	WriteAcceleratorEnabled *bool                   `pulumi:"writeAcceleratorEnabled"`
}





type OSDiskInput interface {
	pulumi.Input

	ToOSDiskOutput() OSDiskOutput
	ToOSDiskOutputWithContext(context.Context) OSDiskOutput
}

type OSDiskArgs struct {
	Caching                 CachingTypesPtrInput           `pulumi:"caching"`
	CreateOption            pulumi.StringInput             `pulumi:"createOption"`
	DeleteOption            pulumi.StringPtrInput          `pulumi:"deleteOption"`
	DiffDiskSettings        DiffDiskSettingsPtrInput       `pulumi:"diffDiskSettings"`
	DiskSizeGB              pulumi.IntPtrInput             `pulumi:"diskSizeGB"`
	EncryptionSettings      DiskEncryptionSettingsPtrInput `pulumi:"encryptionSettings"`
	Image                   VirtualHardDiskPtrInput        `pulumi:"image"`
	ManagedDisk             ManagedDiskParametersPtrInput  `pulumi:"managedDisk"`
	Name                    pulumi.StringPtrInput          `pulumi:"name"`
	OsType                  OperatingSystemTypesPtrInput   `pulumi:"osType"`
	Vhd                     VirtualHardDiskPtrInput        `pulumi:"vhd"`
	WriteAcceleratorEnabled pulumi.BoolPtrInput            `pulumi:"writeAcceleratorEnabled"`
}

func (OSDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDisk)(nil)).Elem()
}

func (i OSDiskArgs) ToOSDiskOutput() OSDiskOutput {
	return i.ToOSDiskOutputWithContext(context.Background())
}

func (i OSDiskArgs) ToOSDiskOutputWithContext(ctx context.Context) OSDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDiskOutput)
}

func (i OSDiskArgs) ToOSDiskPtrOutput() OSDiskPtrOutput {
	return i.ToOSDiskPtrOutputWithContext(context.Background())
}

func (i OSDiskArgs) ToOSDiskPtrOutputWithContext(ctx context.Context) OSDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDiskOutput).ToOSDiskPtrOutputWithContext(ctx)
}









type OSDiskPtrInput interface {
	pulumi.Input

	ToOSDiskPtrOutput() OSDiskPtrOutput
	ToOSDiskPtrOutputWithContext(context.Context) OSDiskPtrOutput
}

type osdiskPtrType OSDiskArgs

func OSDiskPtr(v *OSDiskArgs) OSDiskPtrInput {
	return (*osdiskPtrType)(v)
}

func (*osdiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDisk)(nil)).Elem()
}

func (i *osdiskPtrType) ToOSDiskPtrOutput() OSDiskPtrOutput {
	return i.ToOSDiskPtrOutputWithContext(context.Background())
}

func (i *osdiskPtrType) ToOSDiskPtrOutputWithContext(ctx context.Context) OSDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDiskPtrOutput)
}

type OSDiskOutput struct{ *pulumi.OutputState }

func (OSDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDisk)(nil)).Elem()
}

func (o OSDiskOutput) ToOSDiskOutput() OSDiskOutput {
	return o
}

func (o OSDiskOutput) ToOSDiskOutputWithContext(ctx context.Context) OSDiskOutput {
	return o
}

func (o OSDiskOutput) ToOSDiskPtrOutput() OSDiskPtrOutput {
	return o.ToOSDiskPtrOutputWithContext(context.Background())
}

func (o OSDiskOutput) ToOSDiskPtrOutputWithContext(ctx context.Context) OSDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSDisk) *OSDisk {
		return &v
	}).(OSDiskPtrOutput)
}

func (o OSDiskOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v OSDisk) *CachingTypes { return v.Caching }).(CachingTypesPtrOutput)
}

func (o OSDiskOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v OSDisk) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o OSDiskOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDisk) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o OSDiskOutput) DiffDiskSettings() DiffDiskSettingsPtrOutput {
	return o.ApplyT(func(v OSDisk) *DiffDiskSettings { return v.DiffDiskSettings }).(DiffDiskSettingsPtrOutput)
}

func (o OSDiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OSDisk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o OSDiskOutput) EncryptionSettings() DiskEncryptionSettingsPtrOutput {
	return o.ApplyT(func(v OSDisk) *DiskEncryptionSettings { return v.EncryptionSettings }).(DiskEncryptionSettingsPtrOutput)
}

func (o OSDiskOutput) Image() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v OSDisk) *VirtualHardDisk { return v.Image }).(VirtualHardDiskPtrOutput)
}

func (o OSDiskOutput) ManagedDisk() ManagedDiskParametersPtrOutput {
	return o.ApplyT(func(v OSDisk) *ManagedDiskParameters { return v.ManagedDisk }).(ManagedDiskParametersPtrOutput)
}

func (o OSDiskOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDisk) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OSDiskOutput) OsType() OperatingSystemTypesPtrOutput {
	return o.ApplyT(func(v OSDisk) *OperatingSystemTypes { return v.OsType }).(OperatingSystemTypesPtrOutput)
}

func (o OSDiskOutput) Vhd() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v OSDisk) *VirtualHardDisk { return v.Vhd }).(VirtualHardDiskPtrOutput)
}

func (o OSDiskOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v OSDisk) *bool { return v.WriteAcceleratorEnabled }).(pulumi.BoolPtrOutput)
}

type OSDiskPtrOutput struct{ *pulumi.OutputState }

func (OSDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDisk)(nil)).Elem()
}

func (o OSDiskPtrOutput) ToOSDiskPtrOutput() OSDiskPtrOutput {
	return o
}

func (o OSDiskPtrOutput) ToOSDiskPtrOutputWithContext(ctx context.Context) OSDiskPtrOutput {
	return o
}

func (o OSDiskPtrOutput) Elem() OSDiskOutput {
	return o.ApplyT(func(v *OSDisk) OSDisk {
		if v != nil {
			return *v
		}
		var ret OSDisk
		return ret
	}).(OSDiskOutput)
}

func (o OSDiskPtrOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v *OSDisk) *CachingTypes {
		if v == nil {
			return nil
		}
		return v.Caching
	}).(CachingTypesPtrOutput)
}

func (o OSDiskPtrOutput) CreateOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDisk) *string {
		if v == nil {
			return nil
		}
		return &v.CreateOption
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskPtrOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDisk) *string {
		if v == nil {
			return nil
		}
		return v.DeleteOption
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskPtrOutput) DiffDiskSettings() DiffDiskSettingsPtrOutput {
	return o.ApplyT(func(v *OSDisk) *DiffDiskSettings {
		if v == nil {
			return nil
		}
		return v.DiffDiskSettings
	}).(DiffDiskSettingsPtrOutput)
}

func (o OSDiskPtrOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OSDisk) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o OSDiskPtrOutput) EncryptionSettings() DiskEncryptionSettingsPtrOutput {
	return o.ApplyT(func(v *OSDisk) *DiskEncryptionSettings {
		if v == nil {
			return nil
		}
		return v.EncryptionSettings
	}).(DiskEncryptionSettingsPtrOutput)
}

func (o OSDiskPtrOutput) Image() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v *OSDisk) *VirtualHardDisk {
		if v == nil {
			return nil
		}
		return v.Image
	}).(VirtualHardDiskPtrOutput)
}

func (o OSDiskPtrOutput) ManagedDisk() ManagedDiskParametersPtrOutput {
	return o.ApplyT(func(v *OSDisk) *ManagedDiskParameters {
		if v == nil {
			return nil
		}
		return v.ManagedDisk
	}).(ManagedDiskParametersPtrOutput)
}

func (o OSDiskPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDisk) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskPtrOutput) OsType() OperatingSystemTypesPtrOutput {
	return o.ApplyT(func(v *OSDisk) *OperatingSystemTypes {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(OperatingSystemTypesPtrOutput)
}

func (o OSDiskPtrOutput) Vhd() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v *OSDisk) *VirtualHardDisk {
		if v == nil {
			return nil
		}
		return v.Vhd
	}).(VirtualHardDiskPtrOutput)
}

func (o OSDiskPtrOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OSDisk) *bool {
		if v == nil {
			return nil
		}
		return v.WriteAcceleratorEnabled
	}).(pulumi.BoolPtrOutput)
}

type OSDiskResponse struct {
	Caching                 *string                         `pulumi:"caching"`
	CreateOption            string                          `pulumi:"createOption"`
	DeleteOption            *string                         `pulumi:"deleteOption"`
	DiffDiskSettings        *DiffDiskSettingsResponse       `pulumi:"diffDiskSettings"`
	DiskSizeGB              *int                            `pulumi:"diskSizeGB"`
	EncryptionSettings      *DiskEncryptionSettingsResponse `pulumi:"encryptionSettings"`
	Image                   *VirtualHardDiskResponse        `pulumi:"image"`
	ManagedDisk             *ManagedDiskParametersResponse  `pulumi:"managedDisk"`
	Name                    *string                         `pulumi:"name"`
	OsType                  *string                         `pulumi:"osType"`
	Vhd                     *VirtualHardDiskResponse        `pulumi:"vhd"`
	WriteAcceleratorEnabled *bool                           `pulumi:"writeAcceleratorEnabled"`
}

type OSDiskResponseOutput struct{ *pulumi.OutputState }

func (OSDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDiskResponse)(nil)).Elem()
}

func (o OSDiskResponseOutput) ToOSDiskResponseOutput() OSDiskResponseOutput {
	return o
}

func (o OSDiskResponseOutput) ToOSDiskResponseOutputWithContext(ctx context.Context) OSDiskResponseOutput {
	return o
}

func (o OSDiskResponseOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *string { return v.Caching }).(pulumi.StringPtrOutput)
}

func (o OSDiskResponseOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v OSDiskResponse) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o OSDiskResponseOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o OSDiskResponseOutput) DiffDiskSettings() DiffDiskSettingsResponsePtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *DiffDiskSettingsResponse { return v.DiffDiskSettings }).(DiffDiskSettingsResponsePtrOutput)
}

func (o OSDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o OSDiskResponseOutput) EncryptionSettings() DiskEncryptionSettingsResponsePtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *DiskEncryptionSettingsResponse { return v.EncryptionSettings }).(DiskEncryptionSettingsResponsePtrOutput)
}

func (o OSDiskResponseOutput) Image() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *VirtualHardDiskResponse { return v.Image }).(VirtualHardDiskResponsePtrOutput)
}

func (o OSDiskResponseOutput) ManagedDisk() ManagedDiskParametersResponsePtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *ManagedDiskParametersResponse { return v.ManagedDisk }).(ManagedDiskParametersResponsePtrOutput)
}

func (o OSDiskResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OSDiskResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o OSDiskResponseOutput) Vhd() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *VirtualHardDiskResponse { return v.Vhd }).(VirtualHardDiskResponsePtrOutput)
}

func (o OSDiskResponseOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v OSDiskResponse) *bool { return v.WriteAcceleratorEnabled }).(pulumi.BoolPtrOutput)
}

type OSDiskResponsePtrOutput struct{ *pulumi.OutputState }

func (OSDiskResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDiskResponse)(nil)).Elem()
}

func (o OSDiskResponsePtrOutput) ToOSDiskResponsePtrOutput() OSDiskResponsePtrOutput {
	return o
}

func (o OSDiskResponsePtrOutput) ToOSDiskResponsePtrOutputWithContext(ctx context.Context) OSDiskResponsePtrOutput {
	return o
}

func (o OSDiskResponsePtrOutput) Elem() OSDiskResponseOutput {
	return o.ApplyT(func(v *OSDiskResponse) OSDiskResponse {
		if v != nil {
			return *v
		}
		var ret OSDiskResponse
		return ret
	}).(OSDiskResponseOutput)
}

func (o OSDiskResponsePtrOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.Caching
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskResponsePtrOutput) CreateOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreateOption
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskResponsePtrOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeleteOption
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskResponsePtrOutput) DiffDiskSettings() DiffDiskSettingsResponsePtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *DiffDiskSettingsResponse {
		if v == nil {
			return nil
		}
		return v.DiffDiskSettings
	}).(DiffDiskSettingsResponsePtrOutput)
}

func (o OSDiskResponsePtrOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o OSDiskResponsePtrOutput) EncryptionSettings() DiskEncryptionSettingsResponsePtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *DiskEncryptionSettingsResponse {
		if v == nil {
			return nil
		}
		return v.EncryptionSettings
	}).(DiskEncryptionSettingsResponsePtrOutput)
}

func (o OSDiskResponsePtrOutput) Image() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *VirtualHardDiskResponse {
		if v == nil {
			return nil
		}
		return v.Image
	}).(VirtualHardDiskResponsePtrOutput)
}

func (o OSDiskResponsePtrOutput) ManagedDisk() ManagedDiskParametersResponsePtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *ManagedDiskParametersResponse {
		if v == nil {
			return nil
		}
		return v.ManagedDisk
	}).(ManagedDiskParametersResponsePtrOutput)
}

func (o OSDiskResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskResponsePtrOutput) Vhd() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *VirtualHardDiskResponse {
		if v == nil {
			return nil
		}
		return v.Vhd
	}).(VirtualHardDiskResponsePtrOutput)
}

func (o OSDiskResponsePtrOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OSDiskResponse) *bool {
		if v == nil {
			return nil
		}
		return v.WriteAcceleratorEnabled
	}).(pulumi.BoolPtrOutput)
}

type OSProfile struct {
	AdminPassword               *string               `pulumi:"adminPassword"`
	AdminUsername               *string               `pulumi:"adminUsername"`
	AllowExtensionOperations    *bool                 `pulumi:"allowExtensionOperations"`
	ComputerName                *string               `pulumi:"computerName"`
	CustomData                  *string               `pulumi:"customData"`
	LinuxConfiguration          *LinuxConfiguration   `pulumi:"linuxConfiguration"`
	RequireGuestProvisionSignal *bool                 `pulumi:"requireGuestProvisionSignal"`
	Secrets                     []VaultSecretGroup    `pulumi:"secrets"`
	WindowsConfiguration        *WindowsConfiguration `pulumi:"windowsConfiguration"`
}





type OSProfileInput interface {
	pulumi.Input

	ToOSProfileOutput() OSProfileOutput
	ToOSProfileOutputWithContext(context.Context) OSProfileOutput
}

type OSProfileArgs struct {
	AdminPassword               pulumi.StringPtrInput        `pulumi:"adminPassword"`
	AdminUsername               pulumi.StringPtrInput        `pulumi:"adminUsername"`
	AllowExtensionOperations    pulumi.BoolPtrInput          `pulumi:"allowExtensionOperations"`
	ComputerName                pulumi.StringPtrInput        `pulumi:"computerName"`
	CustomData                  pulumi.StringPtrInput        `pulumi:"customData"`
	LinuxConfiguration          LinuxConfigurationPtrInput   `pulumi:"linuxConfiguration"`
	RequireGuestProvisionSignal pulumi.BoolPtrInput          `pulumi:"requireGuestProvisionSignal"`
	Secrets                     VaultSecretGroupArrayInput   `pulumi:"secrets"`
	WindowsConfiguration        WindowsConfigurationPtrInput `pulumi:"windowsConfiguration"`
}

func (OSProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfile)(nil)).Elem()
}

func (i OSProfileArgs) ToOSProfileOutput() OSProfileOutput {
	return i.ToOSProfileOutputWithContext(context.Background())
}

func (i OSProfileArgs) ToOSProfileOutputWithContext(ctx context.Context) OSProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileOutput)
}

func (i OSProfileArgs) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return i.ToOSProfilePtrOutputWithContext(context.Background())
}

func (i OSProfileArgs) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfileOutput).ToOSProfilePtrOutputWithContext(ctx)
}









type OSProfilePtrInput interface {
	pulumi.Input

	ToOSProfilePtrOutput() OSProfilePtrOutput
	ToOSProfilePtrOutputWithContext(context.Context) OSProfilePtrOutput
}

type osprofilePtrType OSProfileArgs

func OSProfilePtr(v *OSProfileArgs) OSProfilePtrInput {
	return (*osprofilePtrType)(v)
}

func (*osprofilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfile)(nil)).Elem()
}

func (i *osprofilePtrType) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return i.ToOSProfilePtrOutputWithContext(context.Background())
}

func (i *osprofilePtrType) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSProfilePtrOutput)
}

type OSProfileOutput struct{ *pulumi.OutputState }

func (OSProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfile)(nil)).Elem()
}

func (o OSProfileOutput) ToOSProfileOutput() OSProfileOutput {
	return o
}

func (o OSProfileOutput) ToOSProfileOutputWithContext(ctx context.Context) OSProfileOutput {
	return o
}

func (o OSProfileOutput) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return o.ToOSProfilePtrOutputWithContext(context.Background())
}

func (o OSProfileOutput) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSProfile) *OSProfile {
		return &v
	}).(OSProfilePtrOutput)
}

func (o OSProfileOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfile) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o OSProfileOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfile) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o OSProfileOutput) AllowExtensionOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v OSProfile) *bool { return v.AllowExtensionOperations }).(pulumi.BoolPtrOutput)
}

func (o OSProfileOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfile) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o OSProfileOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfile) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

func (o OSProfileOutput) LinuxConfiguration() LinuxConfigurationPtrOutput {
	return o.ApplyT(func(v OSProfile) *LinuxConfiguration { return v.LinuxConfiguration }).(LinuxConfigurationPtrOutput)
}

func (o OSProfileOutput) RequireGuestProvisionSignal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v OSProfile) *bool { return v.RequireGuestProvisionSignal }).(pulumi.BoolPtrOutput)
}

func (o OSProfileOutput) Secrets() VaultSecretGroupArrayOutput {
	return o.ApplyT(func(v OSProfile) []VaultSecretGroup { return v.Secrets }).(VaultSecretGroupArrayOutput)
}

func (o OSProfileOutput) WindowsConfiguration() WindowsConfigurationPtrOutput {
	return o.ApplyT(func(v OSProfile) *WindowsConfiguration { return v.WindowsConfiguration }).(WindowsConfigurationPtrOutput)
}

type OSProfilePtrOutput struct{ *pulumi.OutputState }

func (OSProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfile)(nil)).Elem()
}

func (o OSProfilePtrOutput) ToOSProfilePtrOutput() OSProfilePtrOutput {
	return o
}

func (o OSProfilePtrOutput) ToOSProfilePtrOutputWithContext(ctx context.Context) OSProfilePtrOutput {
	return o
}

func (o OSProfilePtrOutput) Elem() OSProfileOutput {
	return o.ApplyT(func(v *OSProfile) OSProfile {
		if v != nil {
			return *v
		}
		var ret OSProfile
		return ret
	}).(OSProfileOutput)
}

func (o OSProfilePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o OSProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o OSProfilePtrOutput) AllowExtensionOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OSProfile) *bool {
		if v == nil {
			return nil
		}
		return v.AllowExtensionOperations
	}).(pulumi.BoolPtrOutput)
}

func (o OSProfilePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfile) *string {
		if v == nil {
			return nil
		}
		return v.ComputerName
	}).(pulumi.StringPtrOutput)
}

func (o OSProfilePtrOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfile) *string {
		if v == nil {
			return nil
		}
		return v.CustomData
	}).(pulumi.StringPtrOutput)
}

func (o OSProfilePtrOutput) LinuxConfiguration() LinuxConfigurationPtrOutput {
	return o.ApplyT(func(v *OSProfile) *LinuxConfiguration {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(LinuxConfigurationPtrOutput)
}

func (o OSProfilePtrOutput) RequireGuestProvisionSignal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OSProfile) *bool {
		if v == nil {
			return nil
		}
		return v.RequireGuestProvisionSignal
	}).(pulumi.BoolPtrOutput)
}

func (o OSProfilePtrOutput) Secrets() VaultSecretGroupArrayOutput {
	return o.ApplyT(func(v *OSProfile) []VaultSecretGroup {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(VaultSecretGroupArrayOutput)
}

func (o OSProfilePtrOutput) WindowsConfiguration() WindowsConfigurationPtrOutput {
	return o.ApplyT(func(v *OSProfile) *WindowsConfiguration {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(WindowsConfigurationPtrOutput)
}

type OSProfileResponse struct {
	AdminPassword               *string                       `pulumi:"adminPassword"`
	AdminUsername               *string                       `pulumi:"adminUsername"`
	AllowExtensionOperations    *bool                         `pulumi:"allowExtensionOperations"`
	ComputerName                *string                       `pulumi:"computerName"`
	CustomData                  *string                       `pulumi:"customData"`
	LinuxConfiguration          *LinuxConfigurationResponse   `pulumi:"linuxConfiguration"`
	RequireGuestProvisionSignal *bool                         `pulumi:"requireGuestProvisionSignal"`
	Secrets                     []VaultSecretGroupResponse    `pulumi:"secrets"`
	WindowsConfiguration        *WindowsConfigurationResponse `pulumi:"windowsConfiguration"`
}

type OSProfileResponseOutput struct{ *pulumi.OutputState }

func (OSProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSProfileResponse)(nil)).Elem()
}

func (o OSProfileResponseOutput) ToOSProfileResponseOutput() OSProfileResponseOutput {
	return o
}

func (o OSProfileResponseOutput) ToOSProfileResponseOutputWithContext(ctx context.Context) OSProfileResponseOutput {
	return o
}

func (o OSProfileResponseOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o OSProfileResponseOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o OSProfileResponseOutput) AllowExtensionOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *bool { return v.AllowExtensionOperations }).(pulumi.BoolPtrOutput)
}

func (o OSProfileResponseOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o OSProfileResponseOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

func (o OSProfileResponseOutput) LinuxConfiguration() LinuxConfigurationResponsePtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *LinuxConfigurationResponse { return v.LinuxConfiguration }).(LinuxConfigurationResponsePtrOutput)
}

func (o OSProfileResponseOutput) RequireGuestProvisionSignal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *bool { return v.RequireGuestProvisionSignal }).(pulumi.BoolPtrOutput)
}

func (o OSProfileResponseOutput) Secrets() VaultSecretGroupResponseArrayOutput {
	return o.ApplyT(func(v OSProfileResponse) []VaultSecretGroupResponse { return v.Secrets }).(VaultSecretGroupResponseArrayOutput)
}

func (o OSProfileResponseOutput) WindowsConfiguration() WindowsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v OSProfileResponse) *WindowsConfigurationResponse { return v.WindowsConfiguration }).(WindowsConfigurationResponsePtrOutput)
}

type OSProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (OSProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSProfileResponse)(nil)).Elem()
}

func (o OSProfileResponsePtrOutput) ToOSProfileResponsePtrOutput() OSProfileResponsePtrOutput {
	return o
}

func (o OSProfileResponsePtrOutput) ToOSProfileResponsePtrOutputWithContext(ctx context.Context) OSProfileResponsePtrOutput {
	return o
}

func (o OSProfileResponsePtrOutput) Elem() OSProfileResponseOutput {
	return o.ApplyT(func(v *OSProfileResponse) OSProfileResponse {
		if v != nil {
			return *v
		}
		var ret OSProfileResponse
		return ret
	}).(OSProfileResponseOutput)
}

func (o OSProfileResponsePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileResponsePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileResponsePtrOutput) AllowExtensionOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AllowExtensionOperations
	}).(pulumi.BoolPtrOutput)
}

func (o OSProfileResponsePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ComputerName
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileResponsePtrOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomData
	}).(pulumi.StringPtrOutput)
}

func (o OSProfileResponsePtrOutput) LinuxConfiguration() LinuxConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *LinuxConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(LinuxConfigurationResponsePtrOutput)
}

func (o OSProfileResponsePtrOutput) RequireGuestProvisionSignal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequireGuestProvisionSignal
	}).(pulumi.BoolPtrOutput)
}

func (o OSProfileResponsePtrOutput) Secrets() VaultSecretGroupResponseArrayOutput {
	return o.ApplyT(func(v *OSProfileResponse) []VaultSecretGroupResponse {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(VaultSecretGroupResponseArrayOutput)
}

func (o OSProfileResponsePtrOutput) WindowsConfiguration() WindowsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *OSProfileResponse) *WindowsConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(WindowsConfigurationResponsePtrOutput)
}

type PatchSettings struct {
	AssessmentMode              *string                                         `pulumi:"assessmentMode"`
	AutomaticByPlatformSettings *WindowsVMGuestPatchAutomaticByPlatformSettings `pulumi:"automaticByPlatformSettings"`
	EnableHotpatching           *bool                                           `pulumi:"enableHotpatching"`
	PatchMode                   *string                                         `pulumi:"patchMode"`
}





type PatchSettingsInput interface {
	pulumi.Input

	ToPatchSettingsOutput() PatchSettingsOutput
	ToPatchSettingsOutputWithContext(context.Context) PatchSettingsOutput
}

type PatchSettingsArgs struct {
	AssessmentMode              pulumi.StringPtrInput                                  `pulumi:"assessmentMode"`
	AutomaticByPlatformSettings WindowsVMGuestPatchAutomaticByPlatformSettingsPtrInput `pulumi:"automaticByPlatformSettings"`
	EnableHotpatching           pulumi.BoolPtrInput                                    `pulumi:"enableHotpatching"`
	PatchMode                   pulumi.StringPtrInput                                  `pulumi:"patchMode"`
}

func (PatchSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PatchSettings)(nil)).Elem()
}

func (i PatchSettingsArgs) ToPatchSettingsOutput() PatchSettingsOutput {
	return i.ToPatchSettingsOutputWithContext(context.Background())
}

func (i PatchSettingsArgs) ToPatchSettingsOutputWithContext(ctx context.Context) PatchSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchSettingsOutput)
}

func (i PatchSettingsArgs) ToPatchSettingsPtrOutput() PatchSettingsPtrOutput {
	return i.ToPatchSettingsPtrOutputWithContext(context.Background())
}

func (i PatchSettingsArgs) ToPatchSettingsPtrOutputWithContext(ctx context.Context) PatchSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchSettingsOutput).ToPatchSettingsPtrOutputWithContext(ctx)
}









type PatchSettingsPtrInput interface {
	pulumi.Input

	ToPatchSettingsPtrOutput() PatchSettingsPtrOutput
	ToPatchSettingsPtrOutputWithContext(context.Context) PatchSettingsPtrOutput
}

type patchSettingsPtrType PatchSettingsArgs

func PatchSettingsPtr(v *PatchSettingsArgs) PatchSettingsPtrInput {
	return (*patchSettingsPtrType)(v)
}

func (*patchSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PatchSettings)(nil)).Elem()
}

func (i *patchSettingsPtrType) ToPatchSettingsPtrOutput() PatchSettingsPtrOutput {
	return i.ToPatchSettingsPtrOutputWithContext(context.Background())
}

func (i *patchSettingsPtrType) ToPatchSettingsPtrOutputWithContext(ctx context.Context) PatchSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PatchSettingsPtrOutput)
}

type PatchSettingsOutput struct{ *pulumi.OutputState }

func (PatchSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PatchSettings)(nil)).Elem()
}

func (o PatchSettingsOutput) ToPatchSettingsOutput() PatchSettingsOutput {
	return o
}

func (o PatchSettingsOutput) ToPatchSettingsOutputWithContext(ctx context.Context) PatchSettingsOutput {
	return o
}

func (o PatchSettingsOutput) ToPatchSettingsPtrOutput() PatchSettingsPtrOutput {
	return o.ToPatchSettingsPtrOutputWithContext(context.Background())
}

func (o PatchSettingsOutput) ToPatchSettingsPtrOutputWithContext(ctx context.Context) PatchSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PatchSettings) *PatchSettings {
		return &v
	}).(PatchSettingsPtrOutput)
}

func (o PatchSettingsOutput) AssessmentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PatchSettings) *string { return v.AssessmentMode }).(pulumi.StringPtrOutput)
}

func (o PatchSettingsOutput) AutomaticByPlatformSettings() WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return o.ApplyT(func(v PatchSettings) *WindowsVMGuestPatchAutomaticByPlatformSettings {
		return v.AutomaticByPlatformSettings
	}).(WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput)
}

func (o PatchSettingsOutput) EnableHotpatching() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PatchSettings) *bool { return v.EnableHotpatching }).(pulumi.BoolPtrOutput)
}

func (o PatchSettingsOutput) PatchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PatchSettings) *string { return v.PatchMode }).(pulumi.StringPtrOutput)
}

type PatchSettingsPtrOutput struct{ *pulumi.OutputState }

func (PatchSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PatchSettings)(nil)).Elem()
}

func (o PatchSettingsPtrOutput) ToPatchSettingsPtrOutput() PatchSettingsPtrOutput {
	return o
}

func (o PatchSettingsPtrOutput) ToPatchSettingsPtrOutputWithContext(ctx context.Context) PatchSettingsPtrOutput {
	return o
}

func (o PatchSettingsPtrOutput) Elem() PatchSettingsOutput {
	return o.ApplyT(func(v *PatchSettings) PatchSettings {
		if v != nil {
			return *v
		}
		var ret PatchSettings
		return ret
	}).(PatchSettingsOutput)
}

func (o PatchSettingsPtrOutput) AssessmentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PatchSettings) *string {
		if v == nil {
			return nil
		}
		return v.AssessmentMode
	}).(pulumi.StringPtrOutput)
}

func (o PatchSettingsPtrOutput) AutomaticByPlatformSettings() WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return o.ApplyT(func(v *PatchSettings) *WindowsVMGuestPatchAutomaticByPlatformSettings {
		if v == nil {
			return nil
		}
		return v.AutomaticByPlatformSettings
	}).(WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput)
}

func (o PatchSettingsPtrOutput) EnableHotpatching() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PatchSettings) *bool {
		if v == nil {
			return nil
		}
		return v.EnableHotpatching
	}).(pulumi.BoolPtrOutput)
}

func (o PatchSettingsPtrOutput) PatchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PatchSettings) *string {
		if v == nil {
			return nil
		}
		return v.PatchMode
	}).(pulumi.StringPtrOutput)
}

type PatchSettingsResponse struct {
	AssessmentMode              *string                                                 `pulumi:"assessmentMode"`
	AutomaticByPlatformSettings *WindowsVMGuestPatchAutomaticByPlatformSettingsResponse `pulumi:"automaticByPlatformSettings"`
	EnableHotpatching           *bool                                                   `pulumi:"enableHotpatching"`
	PatchMode                   *string                                                 `pulumi:"patchMode"`
}

type PatchSettingsResponseOutput struct{ *pulumi.OutputState }

func (PatchSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PatchSettingsResponse)(nil)).Elem()
}

func (o PatchSettingsResponseOutput) ToPatchSettingsResponseOutput() PatchSettingsResponseOutput {
	return o
}

func (o PatchSettingsResponseOutput) ToPatchSettingsResponseOutputWithContext(ctx context.Context) PatchSettingsResponseOutput {
	return o
}

func (o PatchSettingsResponseOutput) AssessmentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PatchSettingsResponse) *string { return v.AssessmentMode }).(pulumi.StringPtrOutput)
}

func (o PatchSettingsResponseOutput) AutomaticByPlatformSettings() WindowsVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput {
	return o.ApplyT(func(v PatchSettingsResponse) *WindowsVMGuestPatchAutomaticByPlatformSettingsResponse {
		return v.AutomaticByPlatformSettings
	}).(WindowsVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput)
}

func (o PatchSettingsResponseOutput) EnableHotpatching() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PatchSettingsResponse) *bool { return v.EnableHotpatching }).(pulumi.BoolPtrOutput)
}

func (o PatchSettingsResponseOutput) PatchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PatchSettingsResponse) *string { return v.PatchMode }).(pulumi.StringPtrOutput)
}

type PatchSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (PatchSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PatchSettingsResponse)(nil)).Elem()
}

func (o PatchSettingsResponsePtrOutput) ToPatchSettingsResponsePtrOutput() PatchSettingsResponsePtrOutput {
	return o
}

func (o PatchSettingsResponsePtrOutput) ToPatchSettingsResponsePtrOutputWithContext(ctx context.Context) PatchSettingsResponsePtrOutput {
	return o
}

func (o PatchSettingsResponsePtrOutput) Elem() PatchSettingsResponseOutput {
	return o.ApplyT(func(v *PatchSettingsResponse) PatchSettingsResponse {
		if v != nil {
			return *v
		}
		var ret PatchSettingsResponse
		return ret
	}).(PatchSettingsResponseOutput)
}

func (o PatchSettingsResponsePtrOutput) AssessmentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PatchSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AssessmentMode
	}).(pulumi.StringPtrOutput)
}

func (o PatchSettingsResponsePtrOutput) AutomaticByPlatformSettings() WindowsVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput {
	return o.ApplyT(func(v *PatchSettingsResponse) *WindowsVMGuestPatchAutomaticByPlatformSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutomaticByPlatformSettings
	}).(WindowsVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput)
}

func (o PatchSettingsResponsePtrOutput) EnableHotpatching() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PatchSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableHotpatching
	}).(pulumi.BoolPtrOutput)
}

func (o PatchSettingsResponsePtrOutput) PatchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PatchSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.PatchMode
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

type ProximityPlacementGroupPropertiesIntent struct {
	VmSizes []string `pulumi:"vmSizes"`
}





type ProximityPlacementGroupPropertiesIntentInput interface {
	pulumi.Input

	ToProximityPlacementGroupPropertiesIntentOutput() ProximityPlacementGroupPropertiesIntentOutput
	ToProximityPlacementGroupPropertiesIntentOutputWithContext(context.Context) ProximityPlacementGroupPropertiesIntentOutput
}

type ProximityPlacementGroupPropertiesIntentArgs struct {
	VmSizes pulumi.StringArrayInput `pulumi:"vmSizes"`
}

func (ProximityPlacementGroupPropertiesIntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProximityPlacementGroupPropertiesIntent)(nil)).Elem()
}

func (i ProximityPlacementGroupPropertiesIntentArgs) ToProximityPlacementGroupPropertiesIntentOutput() ProximityPlacementGroupPropertiesIntentOutput {
	return i.ToProximityPlacementGroupPropertiesIntentOutputWithContext(context.Background())
}

func (i ProximityPlacementGroupPropertiesIntentArgs) ToProximityPlacementGroupPropertiesIntentOutputWithContext(ctx context.Context) ProximityPlacementGroupPropertiesIntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProximityPlacementGroupPropertiesIntentOutput)
}

func (i ProximityPlacementGroupPropertiesIntentArgs) ToProximityPlacementGroupPropertiesIntentPtrOutput() ProximityPlacementGroupPropertiesIntentPtrOutput {
	return i.ToProximityPlacementGroupPropertiesIntentPtrOutputWithContext(context.Background())
}

func (i ProximityPlacementGroupPropertiesIntentArgs) ToProximityPlacementGroupPropertiesIntentPtrOutputWithContext(ctx context.Context) ProximityPlacementGroupPropertiesIntentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProximityPlacementGroupPropertiesIntentOutput).ToProximityPlacementGroupPropertiesIntentPtrOutputWithContext(ctx)
}









type ProximityPlacementGroupPropertiesIntentPtrInput interface {
	pulumi.Input

	ToProximityPlacementGroupPropertiesIntentPtrOutput() ProximityPlacementGroupPropertiesIntentPtrOutput
	ToProximityPlacementGroupPropertiesIntentPtrOutputWithContext(context.Context) ProximityPlacementGroupPropertiesIntentPtrOutput
}

type proximityPlacementGroupPropertiesIntentPtrType ProximityPlacementGroupPropertiesIntentArgs

func ProximityPlacementGroupPropertiesIntentPtr(v *ProximityPlacementGroupPropertiesIntentArgs) ProximityPlacementGroupPropertiesIntentPtrInput {
	return (*proximityPlacementGroupPropertiesIntentPtrType)(v)
}

func (*proximityPlacementGroupPropertiesIntentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProximityPlacementGroupPropertiesIntent)(nil)).Elem()
}

func (i *proximityPlacementGroupPropertiesIntentPtrType) ToProximityPlacementGroupPropertiesIntentPtrOutput() ProximityPlacementGroupPropertiesIntentPtrOutput {
	return i.ToProximityPlacementGroupPropertiesIntentPtrOutputWithContext(context.Background())
}

func (i *proximityPlacementGroupPropertiesIntentPtrType) ToProximityPlacementGroupPropertiesIntentPtrOutputWithContext(ctx context.Context) ProximityPlacementGroupPropertiesIntentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProximityPlacementGroupPropertiesIntentPtrOutput)
}

type ProximityPlacementGroupPropertiesIntentOutput struct{ *pulumi.OutputState }

func (ProximityPlacementGroupPropertiesIntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProximityPlacementGroupPropertiesIntent)(nil)).Elem()
}

func (o ProximityPlacementGroupPropertiesIntentOutput) ToProximityPlacementGroupPropertiesIntentOutput() ProximityPlacementGroupPropertiesIntentOutput {
	return o
}

func (o ProximityPlacementGroupPropertiesIntentOutput) ToProximityPlacementGroupPropertiesIntentOutputWithContext(ctx context.Context) ProximityPlacementGroupPropertiesIntentOutput {
	return o
}

func (o ProximityPlacementGroupPropertiesIntentOutput) ToProximityPlacementGroupPropertiesIntentPtrOutput() ProximityPlacementGroupPropertiesIntentPtrOutput {
	return o.ToProximityPlacementGroupPropertiesIntentPtrOutputWithContext(context.Background())
}

func (o ProximityPlacementGroupPropertiesIntentOutput) ToProximityPlacementGroupPropertiesIntentPtrOutputWithContext(ctx context.Context) ProximityPlacementGroupPropertiesIntentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProximityPlacementGroupPropertiesIntent) *ProximityPlacementGroupPropertiesIntent {
		return &v
	}).(ProximityPlacementGroupPropertiesIntentPtrOutput)
}

func (o ProximityPlacementGroupPropertiesIntentOutput) VmSizes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ProximityPlacementGroupPropertiesIntent) []string { return v.VmSizes }).(pulumi.StringArrayOutput)
}

type ProximityPlacementGroupPropertiesIntentPtrOutput struct{ *pulumi.OutputState }

func (ProximityPlacementGroupPropertiesIntentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProximityPlacementGroupPropertiesIntent)(nil)).Elem()
}

func (o ProximityPlacementGroupPropertiesIntentPtrOutput) ToProximityPlacementGroupPropertiesIntentPtrOutput() ProximityPlacementGroupPropertiesIntentPtrOutput {
	return o
}

func (o ProximityPlacementGroupPropertiesIntentPtrOutput) ToProximityPlacementGroupPropertiesIntentPtrOutputWithContext(ctx context.Context) ProximityPlacementGroupPropertiesIntentPtrOutput {
	return o
}

func (o ProximityPlacementGroupPropertiesIntentPtrOutput) Elem() ProximityPlacementGroupPropertiesIntentOutput {
	return o.ApplyT(func(v *ProximityPlacementGroupPropertiesIntent) ProximityPlacementGroupPropertiesIntent {
		if v != nil {
			return *v
		}
		var ret ProximityPlacementGroupPropertiesIntent
		return ret
	}).(ProximityPlacementGroupPropertiesIntentOutput)
}

func (o ProximityPlacementGroupPropertiesIntentPtrOutput) VmSizes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProximityPlacementGroupPropertiesIntent) []string {
		if v == nil {
			return nil
		}
		return v.VmSizes
	}).(pulumi.StringArrayOutput)
}

type ProximityPlacementGroupPropertiesResponseIntent struct {
	VmSizes []string `pulumi:"vmSizes"`
}

type ProximityPlacementGroupPropertiesResponseIntentOutput struct{ *pulumi.OutputState }

func (ProximityPlacementGroupPropertiesResponseIntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProximityPlacementGroupPropertiesResponseIntent)(nil)).Elem()
}

func (o ProximityPlacementGroupPropertiesResponseIntentOutput) ToProximityPlacementGroupPropertiesResponseIntentOutput() ProximityPlacementGroupPropertiesResponseIntentOutput {
	return o
}

func (o ProximityPlacementGroupPropertiesResponseIntentOutput) ToProximityPlacementGroupPropertiesResponseIntentOutputWithContext(ctx context.Context) ProximityPlacementGroupPropertiesResponseIntentOutput {
	return o
}

func (o ProximityPlacementGroupPropertiesResponseIntentOutput) VmSizes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ProximityPlacementGroupPropertiesResponseIntent) []string { return v.VmSizes }).(pulumi.StringArrayOutput)
}

type ProximityPlacementGroupPropertiesResponseIntentPtrOutput struct{ *pulumi.OutputState }

func (ProximityPlacementGroupPropertiesResponseIntentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProximityPlacementGroupPropertiesResponseIntent)(nil)).Elem()
}

func (o ProximityPlacementGroupPropertiesResponseIntentPtrOutput) ToProximityPlacementGroupPropertiesResponseIntentPtrOutput() ProximityPlacementGroupPropertiesResponseIntentPtrOutput {
	return o
}

func (o ProximityPlacementGroupPropertiesResponseIntentPtrOutput) ToProximityPlacementGroupPropertiesResponseIntentPtrOutputWithContext(ctx context.Context) ProximityPlacementGroupPropertiesResponseIntentPtrOutput {
	return o
}

func (o ProximityPlacementGroupPropertiesResponseIntentPtrOutput) Elem() ProximityPlacementGroupPropertiesResponseIntentOutput {
	return o.ApplyT(func(v *ProximityPlacementGroupPropertiesResponseIntent) ProximityPlacementGroupPropertiesResponseIntent {
		if v != nil {
			return *v
		}
		var ret ProximityPlacementGroupPropertiesResponseIntent
		return ret
	}).(ProximityPlacementGroupPropertiesResponseIntentOutput)
}

func (o ProximityPlacementGroupPropertiesResponseIntentPtrOutput) VmSizes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProximityPlacementGroupPropertiesResponseIntent) []string {
		if v == nil {
			return nil
		}
		return v.VmSizes
	}).(pulumi.StringArrayOutput)
}

type PublicIPAddressSku struct {
	Name *string `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type PublicIPAddressSkuInput interface {
	pulumi.Input

	ToPublicIPAddressSkuOutput() PublicIPAddressSkuOutput
	ToPublicIPAddressSkuOutputWithContext(context.Context) PublicIPAddressSkuOutput
}

type PublicIPAddressSkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (PublicIPAddressSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressSku)(nil)).Elem()
}

func (i PublicIPAddressSkuArgs) ToPublicIPAddressSkuOutput() PublicIPAddressSkuOutput {
	return i.ToPublicIPAddressSkuOutputWithContext(context.Background())
}

func (i PublicIPAddressSkuArgs) ToPublicIPAddressSkuOutputWithContext(ctx context.Context) PublicIPAddressSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressSkuOutput)
}

func (i PublicIPAddressSkuArgs) ToPublicIPAddressSkuPtrOutput() PublicIPAddressSkuPtrOutput {
	return i.ToPublicIPAddressSkuPtrOutputWithContext(context.Background())
}

func (i PublicIPAddressSkuArgs) ToPublicIPAddressSkuPtrOutputWithContext(ctx context.Context) PublicIPAddressSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressSkuOutput).ToPublicIPAddressSkuPtrOutputWithContext(ctx)
}









type PublicIPAddressSkuPtrInput interface {
	pulumi.Input

	ToPublicIPAddressSkuPtrOutput() PublicIPAddressSkuPtrOutput
	ToPublicIPAddressSkuPtrOutputWithContext(context.Context) PublicIPAddressSkuPtrOutput
}

type publicIPAddressSkuPtrType PublicIPAddressSkuArgs

func PublicIPAddressSkuPtr(v *PublicIPAddressSkuArgs) PublicIPAddressSkuPtrInput {
	return (*publicIPAddressSkuPtrType)(v)
}

func (*publicIPAddressSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressSku)(nil)).Elem()
}

func (i *publicIPAddressSkuPtrType) ToPublicIPAddressSkuPtrOutput() PublicIPAddressSkuPtrOutput {
	return i.ToPublicIPAddressSkuPtrOutputWithContext(context.Background())
}

func (i *publicIPAddressSkuPtrType) ToPublicIPAddressSkuPtrOutputWithContext(ctx context.Context) PublicIPAddressSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressSkuPtrOutput)
}

type PublicIPAddressSkuOutput struct{ *pulumi.OutputState }

func (PublicIPAddressSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressSku)(nil)).Elem()
}

func (o PublicIPAddressSkuOutput) ToPublicIPAddressSkuOutput() PublicIPAddressSkuOutput {
	return o
}

func (o PublicIPAddressSkuOutput) ToPublicIPAddressSkuOutputWithContext(ctx context.Context) PublicIPAddressSkuOutput {
	return o
}

func (o PublicIPAddressSkuOutput) ToPublicIPAddressSkuPtrOutput() PublicIPAddressSkuPtrOutput {
	return o.ToPublicIPAddressSkuPtrOutputWithContext(context.Background())
}

func (o PublicIPAddressSkuOutput) ToPublicIPAddressSkuPtrOutputWithContext(ctx context.Context) PublicIPAddressSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIPAddressSku) *PublicIPAddressSku {
		return &v
	}).(PublicIPAddressSkuPtrOutput)
}

func (o PublicIPAddressSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type PublicIPAddressSkuPtrOutput struct{ *pulumi.OutputState }

func (PublicIPAddressSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressSku)(nil)).Elem()
}

func (o PublicIPAddressSkuPtrOutput) ToPublicIPAddressSkuPtrOutput() PublicIPAddressSkuPtrOutput {
	return o
}

func (o PublicIPAddressSkuPtrOutput) ToPublicIPAddressSkuPtrOutputWithContext(ctx context.Context) PublicIPAddressSkuPtrOutput {
	return o
}

func (o PublicIPAddressSkuPtrOutput) Elem() PublicIPAddressSkuOutput {
	return o.ApplyT(func(v *PublicIPAddressSku) PublicIPAddressSku {
		if v != nil {
			return *v
		}
		var ret PublicIPAddressSku
		return ret
	}).(PublicIPAddressSkuOutput)
}

func (o PublicIPAddressSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressSku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type PublicIPAddressSkuResponse struct {
	Name *string `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}

type PublicIPAddressSkuResponseOutput struct{ *pulumi.OutputState }

func (PublicIPAddressSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressSkuResponse)(nil)).Elem()
}

func (o PublicIPAddressSkuResponseOutput) ToPublicIPAddressSkuResponseOutput() PublicIPAddressSkuResponseOutput {
	return o
}

func (o PublicIPAddressSkuResponseOutput) ToPublicIPAddressSkuResponseOutputWithContext(ctx context.Context) PublicIPAddressSkuResponseOutput {
	return o
}

func (o PublicIPAddressSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type PublicIPAddressSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (PublicIPAddressSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressSkuResponse)(nil)).Elem()
}

func (o PublicIPAddressSkuResponsePtrOutput) ToPublicIPAddressSkuResponsePtrOutput() PublicIPAddressSkuResponsePtrOutput {
	return o
}

func (o PublicIPAddressSkuResponsePtrOutput) ToPublicIPAddressSkuResponsePtrOutputWithContext(ctx context.Context) PublicIPAddressSkuResponsePtrOutput {
	return o
}

func (o PublicIPAddressSkuResponsePtrOutput) Elem() PublicIPAddressSkuResponseOutput {
	return o.ApplyT(func(v *PublicIPAddressSkuResponse) PublicIPAddressSkuResponse {
		if v != nil {
			return *v
		}
		var ret PublicIPAddressSkuResponse
		return ret
	}).(PublicIPAddressSkuResponseOutput)
}

func (o PublicIPAddressSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type RestorePointCollectionSourceProperties struct {
	Id *string `pulumi:"id"`
}





type RestorePointCollectionSourcePropertiesInput interface {
	pulumi.Input

	ToRestorePointCollectionSourcePropertiesOutput() RestorePointCollectionSourcePropertiesOutput
	ToRestorePointCollectionSourcePropertiesOutputWithContext(context.Context) RestorePointCollectionSourcePropertiesOutput
}

type RestorePointCollectionSourcePropertiesArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (RestorePointCollectionSourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RestorePointCollectionSourceProperties)(nil)).Elem()
}

func (i RestorePointCollectionSourcePropertiesArgs) ToRestorePointCollectionSourcePropertiesOutput() RestorePointCollectionSourcePropertiesOutput {
	return i.ToRestorePointCollectionSourcePropertiesOutputWithContext(context.Background())
}

func (i RestorePointCollectionSourcePropertiesArgs) ToRestorePointCollectionSourcePropertiesOutputWithContext(ctx context.Context) RestorePointCollectionSourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestorePointCollectionSourcePropertiesOutput)
}

func (i RestorePointCollectionSourcePropertiesArgs) ToRestorePointCollectionSourcePropertiesPtrOutput() RestorePointCollectionSourcePropertiesPtrOutput {
	return i.ToRestorePointCollectionSourcePropertiesPtrOutputWithContext(context.Background())
}

func (i RestorePointCollectionSourcePropertiesArgs) ToRestorePointCollectionSourcePropertiesPtrOutputWithContext(ctx context.Context) RestorePointCollectionSourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestorePointCollectionSourcePropertiesOutput).ToRestorePointCollectionSourcePropertiesPtrOutputWithContext(ctx)
}









type RestorePointCollectionSourcePropertiesPtrInput interface {
	pulumi.Input

	ToRestorePointCollectionSourcePropertiesPtrOutput() RestorePointCollectionSourcePropertiesPtrOutput
	ToRestorePointCollectionSourcePropertiesPtrOutputWithContext(context.Context) RestorePointCollectionSourcePropertiesPtrOutput
}

type restorePointCollectionSourcePropertiesPtrType RestorePointCollectionSourcePropertiesArgs

func RestorePointCollectionSourcePropertiesPtr(v *RestorePointCollectionSourcePropertiesArgs) RestorePointCollectionSourcePropertiesPtrInput {
	return (*restorePointCollectionSourcePropertiesPtrType)(v)
}

func (*restorePointCollectionSourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePointCollectionSourceProperties)(nil)).Elem()
}

func (i *restorePointCollectionSourcePropertiesPtrType) ToRestorePointCollectionSourcePropertiesPtrOutput() RestorePointCollectionSourcePropertiesPtrOutput {
	return i.ToRestorePointCollectionSourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *restorePointCollectionSourcePropertiesPtrType) ToRestorePointCollectionSourcePropertiesPtrOutputWithContext(ctx context.Context) RestorePointCollectionSourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestorePointCollectionSourcePropertiesPtrOutput)
}

type RestorePointCollectionSourcePropertiesOutput struct{ *pulumi.OutputState }

func (RestorePointCollectionSourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestorePointCollectionSourceProperties)(nil)).Elem()
}

func (o RestorePointCollectionSourcePropertiesOutput) ToRestorePointCollectionSourcePropertiesOutput() RestorePointCollectionSourcePropertiesOutput {
	return o
}

func (o RestorePointCollectionSourcePropertiesOutput) ToRestorePointCollectionSourcePropertiesOutputWithContext(ctx context.Context) RestorePointCollectionSourcePropertiesOutput {
	return o
}

func (o RestorePointCollectionSourcePropertiesOutput) ToRestorePointCollectionSourcePropertiesPtrOutput() RestorePointCollectionSourcePropertiesPtrOutput {
	return o.ToRestorePointCollectionSourcePropertiesPtrOutputWithContext(context.Background())
}

func (o RestorePointCollectionSourcePropertiesOutput) ToRestorePointCollectionSourcePropertiesPtrOutputWithContext(ctx context.Context) RestorePointCollectionSourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RestorePointCollectionSourceProperties) *RestorePointCollectionSourceProperties {
		return &v
	}).(RestorePointCollectionSourcePropertiesPtrOutput)
}

func (o RestorePointCollectionSourcePropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestorePointCollectionSourceProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type RestorePointCollectionSourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (RestorePointCollectionSourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePointCollectionSourceProperties)(nil)).Elem()
}

func (o RestorePointCollectionSourcePropertiesPtrOutput) ToRestorePointCollectionSourcePropertiesPtrOutput() RestorePointCollectionSourcePropertiesPtrOutput {
	return o
}

func (o RestorePointCollectionSourcePropertiesPtrOutput) ToRestorePointCollectionSourcePropertiesPtrOutputWithContext(ctx context.Context) RestorePointCollectionSourcePropertiesPtrOutput {
	return o
}

func (o RestorePointCollectionSourcePropertiesPtrOutput) Elem() RestorePointCollectionSourcePropertiesOutput {
	return o.ApplyT(func(v *RestorePointCollectionSourceProperties) RestorePointCollectionSourceProperties {
		if v != nil {
			return *v
		}
		var ret RestorePointCollectionSourceProperties
		return ret
	}).(RestorePointCollectionSourcePropertiesOutput)
}

func (o RestorePointCollectionSourcePropertiesPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestorePointCollectionSourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type RestorePointCollectionSourcePropertiesResponse struct {
	Id       *string `pulumi:"id"`
	Location string  `pulumi:"location"`
}

type RestorePointCollectionSourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (RestorePointCollectionSourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestorePointCollectionSourcePropertiesResponse)(nil)).Elem()
}

func (o RestorePointCollectionSourcePropertiesResponseOutput) ToRestorePointCollectionSourcePropertiesResponseOutput() RestorePointCollectionSourcePropertiesResponseOutput {
	return o
}

func (o RestorePointCollectionSourcePropertiesResponseOutput) ToRestorePointCollectionSourcePropertiesResponseOutputWithContext(ctx context.Context) RestorePointCollectionSourcePropertiesResponseOutput {
	return o
}

func (o RestorePointCollectionSourcePropertiesResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestorePointCollectionSourcePropertiesResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RestorePointCollectionSourcePropertiesResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v RestorePointCollectionSourcePropertiesResponse) string { return v.Location }).(pulumi.StringOutput)
}

type RestorePointCollectionSourcePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (RestorePointCollectionSourcePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePointCollectionSourcePropertiesResponse)(nil)).Elem()
}

func (o RestorePointCollectionSourcePropertiesResponsePtrOutput) ToRestorePointCollectionSourcePropertiesResponsePtrOutput() RestorePointCollectionSourcePropertiesResponsePtrOutput {
	return o
}

func (o RestorePointCollectionSourcePropertiesResponsePtrOutput) ToRestorePointCollectionSourcePropertiesResponsePtrOutputWithContext(ctx context.Context) RestorePointCollectionSourcePropertiesResponsePtrOutput {
	return o
}

func (o RestorePointCollectionSourcePropertiesResponsePtrOutput) Elem() RestorePointCollectionSourcePropertiesResponseOutput {
	return o.ApplyT(func(v *RestorePointCollectionSourcePropertiesResponse) RestorePointCollectionSourcePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret RestorePointCollectionSourcePropertiesResponse
		return ret
	}).(RestorePointCollectionSourcePropertiesResponseOutput)
}

func (o RestorePointCollectionSourcePropertiesResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestorePointCollectionSourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o RestorePointCollectionSourcePropertiesResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestorePointCollectionSourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

type RestorePointInstanceViewResponse struct {
	DiskRestorePoints []DiskRestorePointInstanceViewResponse `pulumi:"diskRestorePoints"`
	Statuses          []InstanceViewStatusResponse           `pulumi:"statuses"`
}

type RestorePointInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (RestorePointInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestorePointInstanceViewResponse)(nil)).Elem()
}

func (o RestorePointInstanceViewResponseOutput) ToRestorePointInstanceViewResponseOutput() RestorePointInstanceViewResponseOutput {
	return o
}

func (o RestorePointInstanceViewResponseOutput) ToRestorePointInstanceViewResponseOutputWithContext(ctx context.Context) RestorePointInstanceViewResponseOutput {
	return o
}

func (o RestorePointInstanceViewResponseOutput) DiskRestorePoints() DiskRestorePointInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v RestorePointInstanceViewResponse) []DiskRestorePointInstanceViewResponse {
		return v.DiskRestorePoints
	}).(DiskRestorePointInstanceViewResponseArrayOutput)
}

func (o RestorePointInstanceViewResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v RestorePointInstanceViewResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

type RestorePointResponse struct {
	ConsistencyMode    *string                            `pulumi:"consistencyMode"`
	ExcludeDisks       []ApiEntityReferenceResponse       `pulumi:"excludeDisks"`
	Id                 string                             `pulumi:"id"`
	InstanceView       RestorePointInstanceViewResponse   `pulumi:"instanceView"`
	Name               string                             `pulumi:"name"`
	ProvisioningState  string                             `pulumi:"provisioningState"`
	SourceMetadata     RestorePointSourceMetadataResponse `pulumi:"sourceMetadata"`
	SourceRestorePoint *ApiEntityReferenceResponse        `pulumi:"sourceRestorePoint"`
	TimeCreated        *string                            `pulumi:"timeCreated"`
	Type               string                             `pulumi:"type"`
}

type RestorePointResponseOutput struct{ *pulumi.OutputState }

func (RestorePointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestorePointResponse)(nil)).Elem()
}

func (o RestorePointResponseOutput) ToRestorePointResponseOutput() RestorePointResponseOutput {
	return o
}

func (o RestorePointResponseOutput) ToRestorePointResponseOutputWithContext(ctx context.Context) RestorePointResponseOutput {
	return o
}

func (o RestorePointResponseOutput) ConsistencyMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestorePointResponse) *string { return v.ConsistencyMode }).(pulumi.StringPtrOutput)
}

func (o RestorePointResponseOutput) ExcludeDisks() ApiEntityReferenceResponseArrayOutput {
	return o.ApplyT(func(v RestorePointResponse) []ApiEntityReferenceResponse { return v.ExcludeDisks }).(ApiEntityReferenceResponseArrayOutput)
}

func (o RestorePointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RestorePointResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o RestorePointResponseOutput) InstanceView() RestorePointInstanceViewResponseOutput {
	return o.ApplyT(func(v RestorePointResponse) RestorePointInstanceViewResponse { return v.InstanceView }).(RestorePointInstanceViewResponseOutput)
}

func (o RestorePointResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RestorePointResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RestorePointResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v RestorePointResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RestorePointResponseOutput) SourceMetadata() RestorePointSourceMetadataResponseOutput {
	return o.ApplyT(func(v RestorePointResponse) RestorePointSourceMetadataResponse { return v.SourceMetadata }).(RestorePointSourceMetadataResponseOutput)
}

func (o RestorePointResponseOutput) SourceRestorePoint() ApiEntityReferenceResponsePtrOutput {
	return o.ApplyT(func(v RestorePointResponse) *ApiEntityReferenceResponse { return v.SourceRestorePoint }).(ApiEntityReferenceResponsePtrOutput)
}

func (o RestorePointResponseOutput) TimeCreated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestorePointResponse) *string { return v.TimeCreated }).(pulumi.StringPtrOutput)
}

func (o RestorePointResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RestorePointResponse) string { return v.Type }).(pulumi.StringOutput)
}

type RestorePointResponseArrayOutput struct{ *pulumi.OutputState }

func (RestorePointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RestorePointResponse)(nil)).Elem()
}

func (o RestorePointResponseArrayOutput) ToRestorePointResponseArrayOutput() RestorePointResponseArrayOutput {
	return o
}

func (o RestorePointResponseArrayOutput) ToRestorePointResponseArrayOutputWithContext(ctx context.Context) RestorePointResponseArrayOutput {
	return o
}

func (o RestorePointResponseArrayOutput) Index(i pulumi.IntInput) RestorePointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RestorePointResponse {
		return vs[0].([]RestorePointResponse)[vs[1].(int)]
	}).(RestorePointResponseOutput)
}

type RestorePointSourceMetadataResponse struct {
	DiagnosticsProfile *DiagnosticsProfileResponse                 `pulumi:"diagnosticsProfile"`
	HardwareProfile    *HardwareProfileResponse                    `pulumi:"hardwareProfile"`
	LicenseType        *string                                     `pulumi:"licenseType"`
	Location           *string                                     `pulumi:"location"`
	OsProfile          *OSProfileResponse                          `pulumi:"osProfile"`
	SecurityProfile    *SecurityProfileResponse                    `pulumi:"securityProfile"`
	StorageProfile     *RestorePointSourceVMStorageProfileResponse `pulumi:"storageProfile"`
	VmId               *string                                     `pulumi:"vmId"`
}

type RestorePointSourceMetadataResponseOutput struct{ *pulumi.OutputState }

func (RestorePointSourceMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestorePointSourceMetadataResponse)(nil)).Elem()
}

func (o RestorePointSourceMetadataResponseOutput) ToRestorePointSourceMetadataResponseOutput() RestorePointSourceMetadataResponseOutput {
	return o
}

func (o RestorePointSourceMetadataResponseOutput) ToRestorePointSourceMetadataResponseOutputWithContext(ctx context.Context) RestorePointSourceMetadataResponseOutput {
	return o
}

func (o RestorePointSourceMetadataResponseOutput) DiagnosticsProfile() DiagnosticsProfileResponsePtrOutput {
	return o.ApplyT(func(v RestorePointSourceMetadataResponse) *DiagnosticsProfileResponse { return v.DiagnosticsProfile }).(DiagnosticsProfileResponsePtrOutput)
}

func (o RestorePointSourceMetadataResponseOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v RestorePointSourceMetadataResponse) *HardwareProfileResponse { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

func (o RestorePointSourceMetadataResponseOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestorePointSourceMetadataResponse) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o RestorePointSourceMetadataResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestorePointSourceMetadataResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o RestorePointSourceMetadataResponseOutput) OsProfile() OSProfileResponsePtrOutput {
	return o.ApplyT(func(v RestorePointSourceMetadataResponse) *OSProfileResponse { return v.OsProfile }).(OSProfileResponsePtrOutput)
}

func (o RestorePointSourceMetadataResponseOutput) SecurityProfile() SecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v RestorePointSourceMetadataResponse) *SecurityProfileResponse { return v.SecurityProfile }).(SecurityProfileResponsePtrOutput)
}

func (o RestorePointSourceMetadataResponseOutput) StorageProfile() RestorePointSourceVMStorageProfileResponsePtrOutput {
	return o.ApplyT(func(v RestorePointSourceMetadataResponse) *RestorePointSourceVMStorageProfileResponse {
		return v.StorageProfile
	}).(RestorePointSourceVMStorageProfileResponsePtrOutput)
}

func (o RestorePointSourceMetadataResponseOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestorePointSourceMetadataResponse) *string { return v.VmId }).(pulumi.StringPtrOutput)
}

type RestorePointSourceVMDataDiskResponse struct {
	Caching          *string                        `pulumi:"caching"`
	DiskRestorePoint *ApiEntityReferenceResponse    `pulumi:"diskRestorePoint"`
	DiskSizeGB       *int                           `pulumi:"diskSizeGB"`
	Lun              *int                           `pulumi:"lun"`
	ManagedDisk      *ManagedDiskParametersResponse `pulumi:"managedDisk"`
	Name             *string                        `pulumi:"name"`
}

type RestorePointSourceVMDataDiskResponseOutput struct{ *pulumi.OutputState }

func (RestorePointSourceVMDataDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestorePointSourceVMDataDiskResponse)(nil)).Elem()
}

func (o RestorePointSourceVMDataDiskResponseOutput) ToRestorePointSourceVMDataDiskResponseOutput() RestorePointSourceVMDataDiskResponseOutput {
	return o
}

func (o RestorePointSourceVMDataDiskResponseOutput) ToRestorePointSourceVMDataDiskResponseOutputWithContext(ctx context.Context) RestorePointSourceVMDataDiskResponseOutput {
	return o
}

func (o RestorePointSourceVMDataDiskResponseOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestorePointSourceVMDataDiskResponse) *string { return v.Caching }).(pulumi.StringPtrOutput)
}

func (o RestorePointSourceVMDataDiskResponseOutput) DiskRestorePoint() ApiEntityReferenceResponsePtrOutput {
	return o.ApplyT(func(v RestorePointSourceVMDataDiskResponse) *ApiEntityReferenceResponse { return v.DiskRestorePoint }).(ApiEntityReferenceResponsePtrOutput)
}

func (o RestorePointSourceVMDataDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RestorePointSourceVMDataDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o RestorePointSourceVMDataDiskResponseOutput) Lun() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RestorePointSourceVMDataDiskResponse) *int { return v.Lun }).(pulumi.IntPtrOutput)
}

func (o RestorePointSourceVMDataDiskResponseOutput) ManagedDisk() ManagedDiskParametersResponsePtrOutput {
	return o.ApplyT(func(v RestorePointSourceVMDataDiskResponse) *ManagedDiskParametersResponse { return v.ManagedDisk }).(ManagedDiskParametersResponsePtrOutput)
}

func (o RestorePointSourceVMDataDiskResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestorePointSourceVMDataDiskResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type RestorePointSourceVMDataDiskResponseArrayOutput struct{ *pulumi.OutputState }

func (RestorePointSourceVMDataDiskResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RestorePointSourceVMDataDiskResponse)(nil)).Elem()
}

func (o RestorePointSourceVMDataDiskResponseArrayOutput) ToRestorePointSourceVMDataDiskResponseArrayOutput() RestorePointSourceVMDataDiskResponseArrayOutput {
	return o
}

func (o RestorePointSourceVMDataDiskResponseArrayOutput) ToRestorePointSourceVMDataDiskResponseArrayOutputWithContext(ctx context.Context) RestorePointSourceVMDataDiskResponseArrayOutput {
	return o
}

func (o RestorePointSourceVMDataDiskResponseArrayOutput) Index(i pulumi.IntInput) RestorePointSourceVMDataDiskResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RestorePointSourceVMDataDiskResponse {
		return vs[0].([]RestorePointSourceVMDataDiskResponse)[vs[1].(int)]
	}).(RestorePointSourceVMDataDiskResponseOutput)
}

type RestorePointSourceVMOSDiskResponse struct {
	Caching            *string                         `pulumi:"caching"`
	DiskRestorePoint   *ApiEntityReferenceResponse     `pulumi:"diskRestorePoint"`
	DiskSizeGB         *int                            `pulumi:"diskSizeGB"`
	EncryptionSettings *DiskEncryptionSettingsResponse `pulumi:"encryptionSettings"`
	ManagedDisk        *ManagedDiskParametersResponse  `pulumi:"managedDisk"`
	Name               *string                         `pulumi:"name"`
	OsType             *string                         `pulumi:"osType"`
}

type RestorePointSourceVMOSDiskResponseOutput struct{ *pulumi.OutputState }

func (RestorePointSourceVMOSDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestorePointSourceVMOSDiskResponse)(nil)).Elem()
}

func (o RestorePointSourceVMOSDiskResponseOutput) ToRestorePointSourceVMOSDiskResponseOutput() RestorePointSourceVMOSDiskResponseOutput {
	return o
}

func (o RestorePointSourceVMOSDiskResponseOutput) ToRestorePointSourceVMOSDiskResponseOutputWithContext(ctx context.Context) RestorePointSourceVMOSDiskResponseOutput {
	return o
}

func (o RestorePointSourceVMOSDiskResponseOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestorePointSourceVMOSDiskResponse) *string { return v.Caching }).(pulumi.StringPtrOutput)
}

func (o RestorePointSourceVMOSDiskResponseOutput) DiskRestorePoint() ApiEntityReferenceResponsePtrOutput {
	return o.ApplyT(func(v RestorePointSourceVMOSDiskResponse) *ApiEntityReferenceResponse { return v.DiskRestorePoint }).(ApiEntityReferenceResponsePtrOutput)
}

func (o RestorePointSourceVMOSDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RestorePointSourceVMOSDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o RestorePointSourceVMOSDiskResponseOutput) EncryptionSettings() DiskEncryptionSettingsResponsePtrOutput {
	return o.ApplyT(func(v RestorePointSourceVMOSDiskResponse) *DiskEncryptionSettingsResponse {
		return v.EncryptionSettings
	}).(DiskEncryptionSettingsResponsePtrOutput)
}

func (o RestorePointSourceVMOSDiskResponseOutput) ManagedDisk() ManagedDiskParametersResponsePtrOutput {
	return o.ApplyT(func(v RestorePointSourceVMOSDiskResponse) *ManagedDiskParametersResponse { return v.ManagedDisk }).(ManagedDiskParametersResponsePtrOutput)
}

func (o RestorePointSourceVMOSDiskResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestorePointSourceVMOSDiskResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RestorePointSourceVMOSDiskResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestorePointSourceVMOSDiskResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

type RestorePointSourceVMOSDiskResponsePtrOutput struct{ *pulumi.OutputState }

func (RestorePointSourceVMOSDiskResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePointSourceVMOSDiskResponse)(nil)).Elem()
}

func (o RestorePointSourceVMOSDiskResponsePtrOutput) ToRestorePointSourceVMOSDiskResponsePtrOutput() RestorePointSourceVMOSDiskResponsePtrOutput {
	return o
}

func (o RestorePointSourceVMOSDiskResponsePtrOutput) ToRestorePointSourceVMOSDiskResponsePtrOutputWithContext(ctx context.Context) RestorePointSourceVMOSDiskResponsePtrOutput {
	return o
}

func (o RestorePointSourceVMOSDiskResponsePtrOutput) Elem() RestorePointSourceVMOSDiskResponseOutput {
	return o.ApplyT(func(v *RestorePointSourceVMOSDiskResponse) RestorePointSourceVMOSDiskResponse {
		if v != nil {
			return *v
		}
		var ret RestorePointSourceVMOSDiskResponse
		return ret
	}).(RestorePointSourceVMOSDiskResponseOutput)
}

func (o RestorePointSourceVMOSDiskResponsePtrOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestorePointSourceVMOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.Caching
	}).(pulumi.StringPtrOutput)
}

func (o RestorePointSourceVMOSDiskResponsePtrOutput) DiskRestorePoint() ApiEntityReferenceResponsePtrOutput {
	return o.ApplyT(func(v *RestorePointSourceVMOSDiskResponse) *ApiEntityReferenceResponse {
		if v == nil {
			return nil
		}
		return v.DiskRestorePoint
	}).(ApiEntityReferenceResponsePtrOutput)
}

func (o RestorePointSourceVMOSDiskResponsePtrOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RestorePointSourceVMOSDiskResponse) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o RestorePointSourceVMOSDiskResponsePtrOutput) EncryptionSettings() DiskEncryptionSettingsResponsePtrOutput {
	return o.ApplyT(func(v *RestorePointSourceVMOSDiskResponse) *DiskEncryptionSettingsResponse {
		if v == nil {
			return nil
		}
		return v.EncryptionSettings
	}).(DiskEncryptionSettingsResponsePtrOutput)
}

func (o RestorePointSourceVMOSDiskResponsePtrOutput) ManagedDisk() ManagedDiskParametersResponsePtrOutput {
	return o.ApplyT(func(v *RestorePointSourceVMOSDiskResponse) *ManagedDiskParametersResponse {
		if v == nil {
			return nil
		}
		return v.ManagedDisk
	}).(ManagedDiskParametersResponsePtrOutput)
}

func (o RestorePointSourceVMOSDiskResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestorePointSourceVMOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o RestorePointSourceVMOSDiskResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestorePointSourceVMOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

type RestorePointSourceVMStorageProfileResponse struct {
	DataDisks []RestorePointSourceVMDataDiskResponse `pulumi:"dataDisks"`
	OsDisk    *RestorePointSourceVMOSDiskResponse    `pulumi:"osDisk"`
}

type RestorePointSourceVMStorageProfileResponseOutput struct{ *pulumi.OutputState }

func (RestorePointSourceVMStorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestorePointSourceVMStorageProfileResponse)(nil)).Elem()
}

func (o RestorePointSourceVMStorageProfileResponseOutput) ToRestorePointSourceVMStorageProfileResponseOutput() RestorePointSourceVMStorageProfileResponseOutput {
	return o
}

func (o RestorePointSourceVMStorageProfileResponseOutput) ToRestorePointSourceVMStorageProfileResponseOutputWithContext(ctx context.Context) RestorePointSourceVMStorageProfileResponseOutput {
	return o
}

func (o RestorePointSourceVMStorageProfileResponseOutput) DataDisks() RestorePointSourceVMDataDiskResponseArrayOutput {
	return o.ApplyT(func(v RestorePointSourceVMStorageProfileResponse) []RestorePointSourceVMDataDiskResponse {
		return v.DataDisks
	}).(RestorePointSourceVMDataDiskResponseArrayOutput)
}

func (o RestorePointSourceVMStorageProfileResponseOutput) OsDisk() RestorePointSourceVMOSDiskResponsePtrOutput {
	return o.ApplyT(func(v RestorePointSourceVMStorageProfileResponse) *RestorePointSourceVMOSDiskResponse {
		return v.OsDisk
	}).(RestorePointSourceVMOSDiskResponsePtrOutput)
}

type RestorePointSourceVMStorageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (RestorePointSourceVMStorageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePointSourceVMStorageProfileResponse)(nil)).Elem()
}

func (o RestorePointSourceVMStorageProfileResponsePtrOutput) ToRestorePointSourceVMStorageProfileResponsePtrOutput() RestorePointSourceVMStorageProfileResponsePtrOutput {
	return o
}

func (o RestorePointSourceVMStorageProfileResponsePtrOutput) ToRestorePointSourceVMStorageProfileResponsePtrOutputWithContext(ctx context.Context) RestorePointSourceVMStorageProfileResponsePtrOutput {
	return o
}

func (o RestorePointSourceVMStorageProfileResponsePtrOutput) Elem() RestorePointSourceVMStorageProfileResponseOutput {
	return o.ApplyT(func(v *RestorePointSourceVMStorageProfileResponse) RestorePointSourceVMStorageProfileResponse {
		if v != nil {
			return *v
		}
		var ret RestorePointSourceVMStorageProfileResponse
		return ret
	}).(RestorePointSourceVMStorageProfileResponseOutput)
}

func (o RestorePointSourceVMStorageProfileResponsePtrOutput) DataDisks() RestorePointSourceVMDataDiskResponseArrayOutput {
	return o.ApplyT(func(v *RestorePointSourceVMStorageProfileResponse) []RestorePointSourceVMDataDiskResponse {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(RestorePointSourceVMDataDiskResponseArrayOutput)
}

func (o RestorePointSourceVMStorageProfileResponsePtrOutput) OsDisk() RestorePointSourceVMOSDiskResponsePtrOutput {
	return o.ApplyT(func(v *RestorePointSourceVMStorageProfileResponse) *RestorePointSourceVMOSDiskResponse {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(RestorePointSourceVMOSDiskResponsePtrOutput)
}

type RollingUpgradePolicy struct {
	EnableCrossZoneUpgrade              *bool   `pulumi:"enableCrossZoneUpgrade"`
	MaxBatchInstancePercent             *int    `pulumi:"maxBatchInstancePercent"`
	MaxUnhealthyInstancePercent         *int    `pulumi:"maxUnhealthyInstancePercent"`
	MaxUnhealthyUpgradedInstancePercent *int    `pulumi:"maxUnhealthyUpgradedInstancePercent"`
	PauseTimeBetweenBatches             *string `pulumi:"pauseTimeBetweenBatches"`
	PrioritizeUnhealthyInstances        *bool   `pulumi:"prioritizeUnhealthyInstances"`
}





type RollingUpgradePolicyInput interface {
	pulumi.Input

	ToRollingUpgradePolicyOutput() RollingUpgradePolicyOutput
	ToRollingUpgradePolicyOutputWithContext(context.Context) RollingUpgradePolicyOutput
}

type RollingUpgradePolicyArgs struct {
	EnableCrossZoneUpgrade              pulumi.BoolPtrInput   `pulumi:"enableCrossZoneUpgrade"`
	MaxBatchInstancePercent             pulumi.IntPtrInput    `pulumi:"maxBatchInstancePercent"`
	MaxUnhealthyInstancePercent         pulumi.IntPtrInput    `pulumi:"maxUnhealthyInstancePercent"`
	MaxUnhealthyUpgradedInstancePercent pulumi.IntPtrInput    `pulumi:"maxUnhealthyUpgradedInstancePercent"`
	PauseTimeBetweenBatches             pulumi.StringPtrInput `pulumi:"pauseTimeBetweenBatches"`
	PrioritizeUnhealthyInstances        pulumi.BoolPtrInput   `pulumi:"prioritizeUnhealthyInstances"`
}

func (RollingUpgradePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RollingUpgradePolicy)(nil)).Elem()
}

func (i RollingUpgradePolicyArgs) ToRollingUpgradePolicyOutput() RollingUpgradePolicyOutput {
	return i.ToRollingUpgradePolicyOutputWithContext(context.Background())
}

func (i RollingUpgradePolicyArgs) ToRollingUpgradePolicyOutputWithContext(ctx context.Context) RollingUpgradePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RollingUpgradePolicyOutput)
}

func (i RollingUpgradePolicyArgs) ToRollingUpgradePolicyPtrOutput() RollingUpgradePolicyPtrOutput {
	return i.ToRollingUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i RollingUpgradePolicyArgs) ToRollingUpgradePolicyPtrOutputWithContext(ctx context.Context) RollingUpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RollingUpgradePolicyOutput).ToRollingUpgradePolicyPtrOutputWithContext(ctx)
}









type RollingUpgradePolicyPtrInput interface {
	pulumi.Input

	ToRollingUpgradePolicyPtrOutput() RollingUpgradePolicyPtrOutput
	ToRollingUpgradePolicyPtrOutputWithContext(context.Context) RollingUpgradePolicyPtrOutput
}

type rollingUpgradePolicyPtrType RollingUpgradePolicyArgs

func RollingUpgradePolicyPtr(v *RollingUpgradePolicyArgs) RollingUpgradePolicyPtrInput {
	return (*rollingUpgradePolicyPtrType)(v)
}

func (*rollingUpgradePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RollingUpgradePolicy)(nil)).Elem()
}

func (i *rollingUpgradePolicyPtrType) ToRollingUpgradePolicyPtrOutput() RollingUpgradePolicyPtrOutput {
	return i.ToRollingUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i *rollingUpgradePolicyPtrType) ToRollingUpgradePolicyPtrOutputWithContext(ctx context.Context) RollingUpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RollingUpgradePolicyPtrOutput)
}

type RollingUpgradePolicyOutput struct{ *pulumi.OutputState }

func (RollingUpgradePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RollingUpgradePolicy)(nil)).Elem()
}

func (o RollingUpgradePolicyOutput) ToRollingUpgradePolicyOutput() RollingUpgradePolicyOutput {
	return o
}

func (o RollingUpgradePolicyOutput) ToRollingUpgradePolicyOutputWithContext(ctx context.Context) RollingUpgradePolicyOutput {
	return o
}

func (o RollingUpgradePolicyOutput) ToRollingUpgradePolicyPtrOutput() RollingUpgradePolicyPtrOutput {
	return o.ToRollingUpgradePolicyPtrOutputWithContext(context.Background())
}

func (o RollingUpgradePolicyOutput) ToRollingUpgradePolicyPtrOutputWithContext(ctx context.Context) RollingUpgradePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RollingUpgradePolicy) *RollingUpgradePolicy {
		return &v
	}).(RollingUpgradePolicyPtrOutput)
}

func (o RollingUpgradePolicyOutput) EnableCrossZoneUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicy) *bool { return v.EnableCrossZoneUpgrade }).(pulumi.BoolPtrOutput)
}

func (o RollingUpgradePolicyOutput) MaxBatchInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicy) *int { return v.MaxBatchInstancePercent }).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyOutput) MaxUnhealthyInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicy) *int { return v.MaxUnhealthyInstancePercent }).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyOutput) MaxUnhealthyUpgradedInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicy) *int { return v.MaxUnhealthyUpgradedInstancePercent }).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyOutput) PauseTimeBetweenBatches() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicy) *string { return v.PauseTimeBetweenBatches }).(pulumi.StringPtrOutput)
}

func (o RollingUpgradePolicyOutput) PrioritizeUnhealthyInstances() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicy) *bool { return v.PrioritizeUnhealthyInstances }).(pulumi.BoolPtrOutput)
}

type RollingUpgradePolicyPtrOutput struct{ *pulumi.OutputState }

func (RollingUpgradePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RollingUpgradePolicy)(nil)).Elem()
}

func (o RollingUpgradePolicyPtrOutput) ToRollingUpgradePolicyPtrOutput() RollingUpgradePolicyPtrOutput {
	return o
}

func (o RollingUpgradePolicyPtrOutput) ToRollingUpgradePolicyPtrOutputWithContext(ctx context.Context) RollingUpgradePolicyPtrOutput {
	return o
}

func (o RollingUpgradePolicyPtrOutput) Elem() RollingUpgradePolicyOutput {
	return o.ApplyT(func(v *RollingUpgradePolicy) RollingUpgradePolicy {
		if v != nil {
			return *v
		}
		var ret RollingUpgradePolicy
		return ret
	}).(RollingUpgradePolicyOutput)
}

func (o RollingUpgradePolicyPtrOutput) EnableCrossZoneUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicy) *bool {
		if v == nil {
			return nil
		}
		return v.EnableCrossZoneUpgrade
	}).(pulumi.BoolPtrOutput)
}

func (o RollingUpgradePolicyPtrOutput) MaxBatchInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxBatchInstancePercent
	}).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyPtrOutput) MaxUnhealthyInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxUnhealthyInstancePercent
	}).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyPtrOutput) MaxUnhealthyUpgradedInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxUnhealthyUpgradedInstancePercent
	}).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyPtrOutput) PauseTimeBetweenBatches() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicy) *string {
		if v == nil {
			return nil
		}
		return v.PauseTimeBetweenBatches
	}).(pulumi.StringPtrOutput)
}

func (o RollingUpgradePolicyPtrOutput) PrioritizeUnhealthyInstances() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicy) *bool {
		if v == nil {
			return nil
		}
		return v.PrioritizeUnhealthyInstances
	}).(pulumi.BoolPtrOutput)
}

type RollingUpgradePolicyResponse struct {
	EnableCrossZoneUpgrade              *bool   `pulumi:"enableCrossZoneUpgrade"`
	MaxBatchInstancePercent             *int    `pulumi:"maxBatchInstancePercent"`
	MaxUnhealthyInstancePercent         *int    `pulumi:"maxUnhealthyInstancePercent"`
	MaxUnhealthyUpgradedInstancePercent *int    `pulumi:"maxUnhealthyUpgradedInstancePercent"`
	PauseTimeBetweenBatches             *string `pulumi:"pauseTimeBetweenBatches"`
	PrioritizeUnhealthyInstances        *bool   `pulumi:"prioritizeUnhealthyInstances"`
}

type RollingUpgradePolicyResponseOutput struct{ *pulumi.OutputState }

func (RollingUpgradePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RollingUpgradePolicyResponse)(nil)).Elem()
}

func (o RollingUpgradePolicyResponseOutput) ToRollingUpgradePolicyResponseOutput() RollingUpgradePolicyResponseOutput {
	return o
}

func (o RollingUpgradePolicyResponseOutput) ToRollingUpgradePolicyResponseOutputWithContext(ctx context.Context) RollingUpgradePolicyResponseOutput {
	return o
}

func (o RollingUpgradePolicyResponseOutput) EnableCrossZoneUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicyResponse) *bool { return v.EnableCrossZoneUpgrade }).(pulumi.BoolPtrOutput)
}

func (o RollingUpgradePolicyResponseOutput) MaxBatchInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicyResponse) *int { return v.MaxBatchInstancePercent }).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyResponseOutput) MaxUnhealthyInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicyResponse) *int { return v.MaxUnhealthyInstancePercent }).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyResponseOutput) MaxUnhealthyUpgradedInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicyResponse) *int { return v.MaxUnhealthyUpgradedInstancePercent }).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyResponseOutput) PauseTimeBetweenBatches() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicyResponse) *string { return v.PauseTimeBetweenBatches }).(pulumi.StringPtrOutput)
}

func (o RollingUpgradePolicyResponseOutput) PrioritizeUnhealthyInstances() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RollingUpgradePolicyResponse) *bool { return v.PrioritizeUnhealthyInstances }).(pulumi.BoolPtrOutput)
}

type RollingUpgradePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (RollingUpgradePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RollingUpgradePolicyResponse)(nil)).Elem()
}

func (o RollingUpgradePolicyResponsePtrOutput) ToRollingUpgradePolicyResponsePtrOutput() RollingUpgradePolicyResponsePtrOutput {
	return o
}

func (o RollingUpgradePolicyResponsePtrOutput) ToRollingUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) RollingUpgradePolicyResponsePtrOutput {
	return o
}

func (o RollingUpgradePolicyResponsePtrOutput) Elem() RollingUpgradePolicyResponseOutput {
	return o.ApplyT(func(v *RollingUpgradePolicyResponse) RollingUpgradePolicyResponse {
		if v != nil {
			return *v
		}
		var ret RollingUpgradePolicyResponse
		return ret
	}).(RollingUpgradePolicyResponseOutput)
}

func (o RollingUpgradePolicyResponsePtrOutput) EnableCrossZoneUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableCrossZoneUpgrade
	}).(pulumi.BoolPtrOutput)
}

func (o RollingUpgradePolicyResponsePtrOutput) MaxBatchInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxBatchInstancePercent
	}).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyResponsePtrOutput) MaxUnhealthyInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxUnhealthyInstancePercent
	}).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyResponsePtrOutput) MaxUnhealthyUpgradedInstancePercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxUnhealthyUpgradedInstancePercent
	}).(pulumi.IntPtrOutput)
}

func (o RollingUpgradePolicyResponsePtrOutput) PauseTimeBetweenBatches() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.PauseTimeBetweenBatches
	}).(pulumi.StringPtrOutput)
}

func (o RollingUpgradePolicyResponsePtrOutput) PrioritizeUnhealthyInstances() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RollingUpgradePolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.PrioritizeUnhealthyInstances
	}).(pulumi.BoolPtrOutput)
}

type RunCommandInputParameter struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type RunCommandInputParameterInput interface {
	pulumi.Input

	ToRunCommandInputParameterOutput() RunCommandInputParameterOutput
	ToRunCommandInputParameterOutputWithContext(context.Context) RunCommandInputParameterOutput
}

type RunCommandInputParameterArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (RunCommandInputParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunCommandInputParameter)(nil)).Elem()
}

func (i RunCommandInputParameterArgs) ToRunCommandInputParameterOutput() RunCommandInputParameterOutput {
	return i.ToRunCommandInputParameterOutputWithContext(context.Background())
}

func (i RunCommandInputParameterArgs) ToRunCommandInputParameterOutputWithContext(ctx context.Context) RunCommandInputParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunCommandInputParameterOutput)
}





type RunCommandInputParameterArrayInput interface {
	pulumi.Input

	ToRunCommandInputParameterArrayOutput() RunCommandInputParameterArrayOutput
	ToRunCommandInputParameterArrayOutputWithContext(context.Context) RunCommandInputParameterArrayOutput
}

type RunCommandInputParameterArray []RunCommandInputParameterInput

func (RunCommandInputParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RunCommandInputParameter)(nil)).Elem()
}

func (i RunCommandInputParameterArray) ToRunCommandInputParameterArrayOutput() RunCommandInputParameterArrayOutput {
	return i.ToRunCommandInputParameterArrayOutputWithContext(context.Background())
}

func (i RunCommandInputParameterArray) ToRunCommandInputParameterArrayOutputWithContext(ctx context.Context) RunCommandInputParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunCommandInputParameterArrayOutput)
}

type RunCommandInputParameterOutput struct{ *pulumi.OutputState }

func (RunCommandInputParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunCommandInputParameter)(nil)).Elem()
}

func (o RunCommandInputParameterOutput) ToRunCommandInputParameterOutput() RunCommandInputParameterOutput {
	return o
}

func (o RunCommandInputParameterOutput) ToRunCommandInputParameterOutputWithContext(ctx context.Context) RunCommandInputParameterOutput {
	return o
}

func (o RunCommandInputParameterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RunCommandInputParameter) string { return v.Name }).(pulumi.StringOutput)
}

func (o RunCommandInputParameterOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v RunCommandInputParameter) string { return v.Value }).(pulumi.StringOutput)
}

type RunCommandInputParameterArrayOutput struct{ *pulumi.OutputState }

func (RunCommandInputParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RunCommandInputParameter)(nil)).Elem()
}

func (o RunCommandInputParameterArrayOutput) ToRunCommandInputParameterArrayOutput() RunCommandInputParameterArrayOutput {
	return o
}

func (o RunCommandInputParameterArrayOutput) ToRunCommandInputParameterArrayOutputWithContext(ctx context.Context) RunCommandInputParameterArrayOutput {
	return o
}

func (o RunCommandInputParameterArrayOutput) Index(i pulumi.IntInput) RunCommandInputParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RunCommandInputParameter {
		return vs[0].([]RunCommandInputParameter)[vs[1].(int)]
	}).(RunCommandInputParameterOutput)
}

type RunCommandInputParameterResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type RunCommandInputParameterResponseOutput struct{ *pulumi.OutputState }

func (RunCommandInputParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunCommandInputParameterResponse)(nil)).Elem()
}

func (o RunCommandInputParameterResponseOutput) ToRunCommandInputParameterResponseOutput() RunCommandInputParameterResponseOutput {
	return o
}

func (o RunCommandInputParameterResponseOutput) ToRunCommandInputParameterResponseOutputWithContext(ctx context.Context) RunCommandInputParameterResponseOutput {
	return o
}

func (o RunCommandInputParameterResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RunCommandInputParameterResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RunCommandInputParameterResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v RunCommandInputParameterResponse) string { return v.Value }).(pulumi.StringOutput)
}

type RunCommandInputParameterResponseArrayOutput struct{ *pulumi.OutputState }

func (RunCommandInputParameterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RunCommandInputParameterResponse)(nil)).Elem()
}

func (o RunCommandInputParameterResponseArrayOutput) ToRunCommandInputParameterResponseArrayOutput() RunCommandInputParameterResponseArrayOutput {
	return o
}

func (o RunCommandInputParameterResponseArrayOutput) ToRunCommandInputParameterResponseArrayOutputWithContext(ctx context.Context) RunCommandInputParameterResponseArrayOutput {
	return o
}

func (o RunCommandInputParameterResponseArrayOutput) Index(i pulumi.IntInput) RunCommandInputParameterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RunCommandInputParameterResponse {
		return vs[0].([]RunCommandInputParameterResponse)[vs[1].(int)]
	}).(RunCommandInputParameterResponseOutput)
}

type ScaleInPolicy struct {
	ForceDeletion *bool    `pulumi:"forceDeletion"`
	Rules         []string `pulumi:"rules"`
}





type ScaleInPolicyInput interface {
	pulumi.Input

	ToScaleInPolicyOutput() ScaleInPolicyOutput
	ToScaleInPolicyOutputWithContext(context.Context) ScaleInPolicyOutput
}

type ScaleInPolicyArgs struct {
	ForceDeletion pulumi.BoolPtrInput     `pulumi:"forceDeletion"`
	Rules         pulumi.StringArrayInput `pulumi:"rules"`
}

func (ScaleInPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleInPolicy)(nil)).Elem()
}

func (i ScaleInPolicyArgs) ToScaleInPolicyOutput() ScaleInPolicyOutput {
	return i.ToScaleInPolicyOutputWithContext(context.Background())
}

func (i ScaleInPolicyArgs) ToScaleInPolicyOutputWithContext(ctx context.Context) ScaleInPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleInPolicyOutput)
}

func (i ScaleInPolicyArgs) ToScaleInPolicyPtrOutput() ScaleInPolicyPtrOutput {
	return i.ToScaleInPolicyPtrOutputWithContext(context.Background())
}

func (i ScaleInPolicyArgs) ToScaleInPolicyPtrOutputWithContext(ctx context.Context) ScaleInPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleInPolicyOutput).ToScaleInPolicyPtrOutputWithContext(ctx)
}









type ScaleInPolicyPtrInput interface {
	pulumi.Input

	ToScaleInPolicyPtrOutput() ScaleInPolicyPtrOutput
	ToScaleInPolicyPtrOutputWithContext(context.Context) ScaleInPolicyPtrOutput
}

type scaleInPolicyPtrType ScaleInPolicyArgs

func ScaleInPolicyPtr(v *ScaleInPolicyArgs) ScaleInPolicyPtrInput {
	return (*scaleInPolicyPtrType)(v)
}

func (*scaleInPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleInPolicy)(nil)).Elem()
}

func (i *scaleInPolicyPtrType) ToScaleInPolicyPtrOutput() ScaleInPolicyPtrOutput {
	return i.ToScaleInPolicyPtrOutputWithContext(context.Background())
}

func (i *scaleInPolicyPtrType) ToScaleInPolicyPtrOutputWithContext(ctx context.Context) ScaleInPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleInPolicyPtrOutput)
}

type ScaleInPolicyOutput struct{ *pulumi.OutputState }

func (ScaleInPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleInPolicy)(nil)).Elem()
}

func (o ScaleInPolicyOutput) ToScaleInPolicyOutput() ScaleInPolicyOutput {
	return o
}

func (o ScaleInPolicyOutput) ToScaleInPolicyOutputWithContext(ctx context.Context) ScaleInPolicyOutput {
	return o
}

func (o ScaleInPolicyOutput) ToScaleInPolicyPtrOutput() ScaleInPolicyPtrOutput {
	return o.ToScaleInPolicyPtrOutputWithContext(context.Background())
}

func (o ScaleInPolicyOutput) ToScaleInPolicyPtrOutputWithContext(ctx context.Context) ScaleInPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScaleInPolicy) *ScaleInPolicy {
		return &v
	}).(ScaleInPolicyPtrOutput)
}

func (o ScaleInPolicyOutput) ForceDeletion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ScaleInPolicy) *bool { return v.ForceDeletion }).(pulumi.BoolPtrOutput)
}

func (o ScaleInPolicyOutput) Rules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScaleInPolicy) []string { return v.Rules }).(pulumi.StringArrayOutput)
}

type ScaleInPolicyPtrOutput struct{ *pulumi.OutputState }

func (ScaleInPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleInPolicy)(nil)).Elem()
}

func (o ScaleInPolicyPtrOutput) ToScaleInPolicyPtrOutput() ScaleInPolicyPtrOutput {
	return o
}

func (o ScaleInPolicyPtrOutput) ToScaleInPolicyPtrOutputWithContext(ctx context.Context) ScaleInPolicyPtrOutput {
	return o
}

func (o ScaleInPolicyPtrOutput) Elem() ScaleInPolicyOutput {
	return o.ApplyT(func(v *ScaleInPolicy) ScaleInPolicy {
		if v != nil {
			return *v
		}
		var ret ScaleInPolicy
		return ret
	}).(ScaleInPolicyOutput)
}

func (o ScaleInPolicyPtrOutput) ForceDeletion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScaleInPolicy) *bool {
		if v == nil {
			return nil
		}
		return v.ForceDeletion
	}).(pulumi.BoolPtrOutput)
}

func (o ScaleInPolicyPtrOutput) Rules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScaleInPolicy) []string {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(pulumi.StringArrayOutput)
}

type ScaleInPolicyResponse struct {
	ForceDeletion *bool    `pulumi:"forceDeletion"`
	Rules         []string `pulumi:"rules"`
}

type ScaleInPolicyResponseOutput struct{ *pulumi.OutputState }

func (ScaleInPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleInPolicyResponse)(nil)).Elem()
}

func (o ScaleInPolicyResponseOutput) ToScaleInPolicyResponseOutput() ScaleInPolicyResponseOutput {
	return o
}

func (o ScaleInPolicyResponseOutput) ToScaleInPolicyResponseOutputWithContext(ctx context.Context) ScaleInPolicyResponseOutput {
	return o
}

func (o ScaleInPolicyResponseOutput) ForceDeletion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ScaleInPolicyResponse) *bool { return v.ForceDeletion }).(pulumi.BoolPtrOutput)
}

func (o ScaleInPolicyResponseOutput) Rules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScaleInPolicyResponse) []string { return v.Rules }).(pulumi.StringArrayOutput)
}

type ScaleInPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ScaleInPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleInPolicyResponse)(nil)).Elem()
}

func (o ScaleInPolicyResponsePtrOutput) ToScaleInPolicyResponsePtrOutput() ScaleInPolicyResponsePtrOutput {
	return o
}

func (o ScaleInPolicyResponsePtrOutput) ToScaleInPolicyResponsePtrOutputWithContext(ctx context.Context) ScaleInPolicyResponsePtrOutput {
	return o
}

func (o ScaleInPolicyResponsePtrOutput) Elem() ScaleInPolicyResponseOutput {
	return o.ApplyT(func(v *ScaleInPolicyResponse) ScaleInPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ScaleInPolicyResponse
		return ret
	}).(ScaleInPolicyResponseOutput)
}

func (o ScaleInPolicyResponsePtrOutput) ForceDeletion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScaleInPolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ForceDeletion
	}).(pulumi.BoolPtrOutput)
}

func (o ScaleInPolicyResponsePtrOutput) Rules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScaleInPolicyResponse) []string {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(pulumi.StringArrayOutput)
}

type ScheduledEventsProfile struct {
	TerminateNotificationProfile *TerminateNotificationProfile `pulumi:"terminateNotificationProfile"`
}





type ScheduledEventsProfileInput interface {
	pulumi.Input

	ToScheduledEventsProfileOutput() ScheduledEventsProfileOutput
	ToScheduledEventsProfileOutputWithContext(context.Context) ScheduledEventsProfileOutput
}

type ScheduledEventsProfileArgs struct {
	TerminateNotificationProfile TerminateNotificationProfilePtrInput `pulumi:"terminateNotificationProfile"`
}

func (ScheduledEventsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledEventsProfile)(nil)).Elem()
}

func (i ScheduledEventsProfileArgs) ToScheduledEventsProfileOutput() ScheduledEventsProfileOutput {
	return i.ToScheduledEventsProfileOutputWithContext(context.Background())
}

func (i ScheduledEventsProfileArgs) ToScheduledEventsProfileOutputWithContext(ctx context.Context) ScheduledEventsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledEventsProfileOutput)
}

func (i ScheduledEventsProfileArgs) ToScheduledEventsProfilePtrOutput() ScheduledEventsProfilePtrOutput {
	return i.ToScheduledEventsProfilePtrOutputWithContext(context.Background())
}

func (i ScheduledEventsProfileArgs) ToScheduledEventsProfilePtrOutputWithContext(ctx context.Context) ScheduledEventsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledEventsProfileOutput).ToScheduledEventsProfilePtrOutputWithContext(ctx)
}









type ScheduledEventsProfilePtrInput interface {
	pulumi.Input

	ToScheduledEventsProfilePtrOutput() ScheduledEventsProfilePtrOutput
	ToScheduledEventsProfilePtrOutputWithContext(context.Context) ScheduledEventsProfilePtrOutput
}

type scheduledEventsProfilePtrType ScheduledEventsProfileArgs

func ScheduledEventsProfilePtr(v *ScheduledEventsProfileArgs) ScheduledEventsProfilePtrInput {
	return (*scheduledEventsProfilePtrType)(v)
}

func (*scheduledEventsProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledEventsProfile)(nil)).Elem()
}

func (i *scheduledEventsProfilePtrType) ToScheduledEventsProfilePtrOutput() ScheduledEventsProfilePtrOutput {
	return i.ToScheduledEventsProfilePtrOutputWithContext(context.Background())
}

func (i *scheduledEventsProfilePtrType) ToScheduledEventsProfilePtrOutputWithContext(ctx context.Context) ScheduledEventsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledEventsProfilePtrOutput)
}

type ScheduledEventsProfileOutput struct{ *pulumi.OutputState }

func (ScheduledEventsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledEventsProfile)(nil)).Elem()
}

func (o ScheduledEventsProfileOutput) ToScheduledEventsProfileOutput() ScheduledEventsProfileOutput {
	return o
}

func (o ScheduledEventsProfileOutput) ToScheduledEventsProfileOutputWithContext(ctx context.Context) ScheduledEventsProfileOutput {
	return o
}

func (o ScheduledEventsProfileOutput) ToScheduledEventsProfilePtrOutput() ScheduledEventsProfilePtrOutput {
	return o.ToScheduledEventsProfilePtrOutputWithContext(context.Background())
}

func (o ScheduledEventsProfileOutput) ToScheduledEventsProfilePtrOutputWithContext(ctx context.Context) ScheduledEventsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduledEventsProfile) *ScheduledEventsProfile {
		return &v
	}).(ScheduledEventsProfilePtrOutput)
}

func (o ScheduledEventsProfileOutput) TerminateNotificationProfile() TerminateNotificationProfilePtrOutput {
	return o.ApplyT(func(v ScheduledEventsProfile) *TerminateNotificationProfile { return v.TerminateNotificationProfile }).(TerminateNotificationProfilePtrOutput)
}

type ScheduledEventsProfilePtrOutput struct{ *pulumi.OutputState }

func (ScheduledEventsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledEventsProfile)(nil)).Elem()
}

func (o ScheduledEventsProfilePtrOutput) ToScheduledEventsProfilePtrOutput() ScheduledEventsProfilePtrOutput {
	return o
}

func (o ScheduledEventsProfilePtrOutput) ToScheduledEventsProfilePtrOutputWithContext(ctx context.Context) ScheduledEventsProfilePtrOutput {
	return o
}

func (o ScheduledEventsProfilePtrOutput) Elem() ScheduledEventsProfileOutput {
	return o.ApplyT(func(v *ScheduledEventsProfile) ScheduledEventsProfile {
		if v != nil {
			return *v
		}
		var ret ScheduledEventsProfile
		return ret
	}).(ScheduledEventsProfileOutput)
}

func (o ScheduledEventsProfilePtrOutput) TerminateNotificationProfile() TerminateNotificationProfilePtrOutput {
	return o.ApplyT(func(v *ScheduledEventsProfile) *TerminateNotificationProfile {
		if v == nil {
			return nil
		}
		return v.TerminateNotificationProfile
	}).(TerminateNotificationProfilePtrOutput)
}

type ScheduledEventsProfileResponse struct {
	TerminateNotificationProfile *TerminateNotificationProfileResponse `pulumi:"terminateNotificationProfile"`
}

type ScheduledEventsProfileResponseOutput struct{ *pulumi.OutputState }

func (ScheduledEventsProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledEventsProfileResponse)(nil)).Elem()
}

func (o ScheduledEventsProfileResponseOutput) ToScheduledEventsProfileResponseOutput() ScheduledEventsProfileResponseOutput {
	return o
}

func (o ScheduledEventsProfileResponseOutput) ToScheduledEventsProfileResponseOutputWithContext(ctx context.Context) ScheduledEventsProfileResponseOutput {
	return o
}

func (o ScheduledEventsProfileResponseOutput) TerminateNotificationProfile() TerminateNotificationProfileResponsePtrOutput {
	return o.ApplyT(func(v ScheduledEventsProfileResponse) *TerminateNotificationProfileResponse {
		return v.TerminateNotificationProfile
	}).(TerminateNotificationProfileResponsePtrOutput)
}

type ScheduledEventsProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ScheduledEventsProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledEventsProfileResponse)(nil)).Elem()
}

func (o ScheduledEventsProfileResponsePtrOutput) ToScheduledEventsProfileResponsePtrOutput() ScheduledEventsProfileResponsePtrOutput {
	return o
}

func (o ScheduledEventsProfileResponsePtrOutput) ToScheduledEventsProfileResponsePtrOutputWithContext(ctx context.Context) ScheduledEventsProfileResponsePtrOutput {
	return o
}

func (o ScheduledEventsProfileResponsePtrOutput) Elem() ScheduledEventsProfileResponseOutput {
	return o.ApplyT(func(v *ScheduledEventsProfileResponse) ScheduledEventsProfileResponse {
		if v != nil {
			return *v
		}
		var ret ScheduledEventsProfileResponse
		return ret
	}).(ScheduledEventsProfileResponseOutput)
}

func (o ScheduledEventsProfileResponsePtrOutput) TerminateNotificationProfile() TerminateNotificationProfileResponsePtrOutput {
	return o.ApplyT(func(v *ScheduledEventsProfileResponse) *TerminateNotificationProfileResponse {
		if v == nil {
			return nil
		}
		return v.TerminateNotificationProfile
	}).(TerminateNotificationProfileResponsePtrOutput)
}

type SecurityProfile struct {
	EncryptionAtHost *bool         `pulumi:"encryptionAtHost"`
	SecurityType     *string       `pulumi:"securityType"`
	UefiSettings     *UefiSettings `pulumi:"uefiSettings"`
}





type SecurityProfileInput interface {
	pulumi.Input

	ToSecurityProfileOutput() SecurityProfileOutput
	ToSecurityProfileOutputWithContext(context.Context) SecurityProfileOutput
}

type SecurityProfileArgs struct {
	EncryptionAtHost pulumi.BoolPtrInput   `pulumi:"encryptionAtHost"`
	SecurityType     pulumi.StringPtrInput `pulumi:"securityType"`
	UefiSettings     UefiSettingsPtrInput  `pulumi:"uefiSettings"`
}

func (SecurityProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityProfile)(nil)).Elem()
}

func (i SecurityProfileArgs) ToSecurityProfileOutput() SecurityProfileOutput {
	return i.ToSecurityProfileOutputWithContext(context.Background())
}

func (i SecurityProfileArgs) ToSecurityProfileOutputWithContext(ctx context.Context) SecurityProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityProfileOutput)
}

func (i SecurityProfileArgs) ToSecurityProfilePtrOutput() SecurityProfilePtrOutput {
	return i.ToSecurityProfilePtrOutputWithContext(context.Background())
}

func (i SecurityProfileArgs) ToSecurityProfilePtrOutputWithContext(ctx context.Context) SecurityProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityProfileOutput).ToSecurityProfilePtrOutputWithContext(ctx)
}









type SecurityProfilePtrInput interface {
	pulumi.Input

	ToSecurityProfilePtrOutput() SecurityProfilePtrOutput
	ToSecurityProfilePtrOutputWithContext(context.Context) SecurityProfilePtrOutput
}

type securityProfilePtrType SecurityProfileArgs

func SecurityProfilePtr(v *SecurityProfileArgs) SecurityProfilePtrInput {
	return (*securityProfilePtrType)(v)
}

func (*securityProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityProfile)(nil)).Elem()
}

func (i *securityProfilePtrType) ToSecurityProfilePtrOutput() SecurityProfilePtrOutput {
	return i.ToSecurityProfilePtrOutputWithContext(context.Background())
}

func (i *securityProfilePtrType) ToSecurityProfilePtrOutputWithContext(ctx context.Context) SecurityProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityProfilePtrOutput)
}

type SecurityProfileOutput struct{ *pulumi.OutputState }

func (SecurityProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityProfile)(nil)).Elem()
}

func (o SecurityProfileOutput) ToSecurityProfileOutput() SecurityProfileOutput {
	return o
}

func (o SecurityProfileOutput) ToSecurityProfileOutputWithContext(ctx context.Context) SecurityProfileOutput {
	return o
}

func (o SecurityProfileOutput) ToSecurityProfilePtrOutput() SecurityProfilePtrOutput {
	return o.ToSecurityProfilePtrOutputWithContext(context.Background())
}

func (o SecurityProfileOutput) ToSecurityProfilePtrOutputWithContext(ctx context.Context) SecurityProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityProfile) *SecurityProfile {
		return &v
	}).(SecurityProfilePtrOutput)
}

func (o SecurityProfileOutput) EncryptionAtHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SecurityProfile) *bool { return v.EncryptionAtHost }).(pulumi.BoolPtrOutput)
}

func (o SecurityProfileOutput) SecurityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityProfile) *string { return v.SecurityType }).(pulumi.StringPtrOutput)
}

func (o SecurityProfileOutput) UefiSettings() UefiSettingsPtrOutput {
	return o.ApplyT(func(v SecurityProfile) *UefiSettings { return v.UefiSettings }).(UefiSettingsPtrOutput)
}

type SecurityProfilePtrOutput struct{ *pulumi.OutputState }

func (SecurityProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityProfile)(nil)).Elem()
}

func (o SecurityProfilePtrOutput) ToSecurityProfilePtrOutput() SecurityProfilePtrOutput {
	return o
}

func (o SecurityProfilePtrOutput) ToSecurityProfilePtrOutputWithContext(ctx context.Context) SecurityProfilePtrOutput {
	return o
}

func (o SecurityProfilePtrOutput) Elem() SecurityProfileOutput {
	return o.ApplyT(func(v *SecurityProfile) SecurityProfile {
		if v != nil {
			return *v
		}
		var ret SecurityProfile
		return ret
	}).(SecurityProfileOutput)
}

func (o SecurityProfilePtrOutput) EncryptionAtHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityProfile) *bool {
		if v == nil {
			return nil
		}
		return v.EncryptionAtHost
	}).(pulumi.BoolPtrOutput)
}

func (o SecurityProfilePtrOutput) SecurityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityProfile) *string {
		if v == nil {
			return nil
		}
		return v.SecurityType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityProfilePtrOutput) UefiSettings() UefiSettingsPtrOutput {
	return o.ApplyT(func(v *SecurityProfile) *UefiSettings {
		if v == nil {
			return nil
		}
		return v.UefiSettings
	}).(UefiSettingsPtrOutput)
}

type SecurityProfileResponse struct {
	EncryptionAtHost *bool                 `pulumi:"encryptionAtHost"`
	SecurityType     *string               `pulumi:"securityType"`
	UefiSettings     *UefiSettingsResponse `pulumi:"uefiSettings"`
}

type SecurityProfileResponseOutput struct{ *pulumi.OutputState }

func (SecurityProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityProfileResponse)(nil)).Elem()
}

func (o SecurityProfileResponseOutput) ToSecurityProfileResponseOutput() SecurityProfileResponseOutput {
	return o
}

func (o SecurityProfileResponseOutput) ToSecurityProfileResponseOutputWithContext(ctx context.Context) SecurityProfileResponseOutput {
	return o
}

func (o SecurityProfileResponseOutput) EncryptionAtHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SecurityProfileResponse) *bool { return v.EncryptionAtHost }).(pulumi.BoolPtrOutput)
}

func (o SecurityProfileResponseOutput) SecurityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityProfileResponse) *string { return v.SecurityType }).(pulumi.StringPtrOutput)
}

func (o SecurityProfileResponseOutput) UefiSettings() UefiSettingsResponsePtrOutput {
	return o.ApplyT(func(v SecurityProfileResponse) *UefiSettingsResponse { return v.UefiSettings }).(UefiSettingsResponsePtrOutput)
}

type SecurityProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (SecurityProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityProfileResponse)(nil)).Elem()
}

func (o SecurityProfileResponsePtrOutput) ToSecurityProfileResponsePtrOutput() SecurityProfileResponsePtrOutput {
	return o
}

func (o SecurityProfileResponsePtrOutput) ToSecurityProfileResponsePtrOutputWithContext(ctx context.Context) SecurityProfileResponsePtrOutput {
	return o
}

func (o SecurityProfileResponsePtrOutput) Elem() SecurityProfileResponseOutput {
	return o.ApplyT(func(v *SecurityProfileResponse) SecurityProfileResponse {
		if v != nil {
			return *v
		}
		var ret SecurityProfileResponse
		return ret
	}).(SecurityProfileResponseOutput)
}

func (o SecurityProfileResponsePtrOutput) EncryptionAtHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EncryptionAtHost
	}).(pulumi.BoolPtrOutput)
}

func (o SecurityProfileResponsePtrOutput) SecurityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.SecurityType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityProfileResponsePtrOutput) UefiSettings() UefiSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SecurityProfileResponse) *UefiSettingsResponse {
		if v == nil {
			return nil
		}
		return v.UefiSettings
	}).(UefiSettingsResponsePtrOutput)
}

type Sku struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     *string  `pulumi:"name"`
	Tier     *string  `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.Float64PtrInput `pulumi:"capacity"`
	Name     pulumi.StringPtrInput  `pulumi:"name"`
	Tier     pulumi.StringPtrInput  `pulumi:"tier"`
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

func (o SkuOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v Sku) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
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

func (o SkuPtrOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Sku) *float64 {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.Float64PtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     *string  `pulumi:"name"`
	Tier     *string  `pulumi:"tier"`
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

func (o SkuResponseOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SkuResponse) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
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

func (o SkuResponsePtrOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SkuResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.Float64PtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SpotRestorePolicy struct {
	Enabled        *bool   `pulumi:"enabled"`
	RestoreTimeout *string `pulumi:"restoreTimeout"`
}





type SpotRestorePolicyInput interface {
	pulumi.Input

	ToSpotRestorePolicyOutput() SpotRestorePolicyOutput
	ToSpotRestorePolicyOutputWithContext(context.Context) SpotRestorePolicyOutput
}

type SpotRestorePolicyArgs struct {
	Enabled        pulumi.BoolPtrInput   `pulumi:"enabled"`
	RestoreTimeout pulumi.StringPtrInput `pulumi:"restoreTimeout"`
}

func (SpotRestorePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SpotRestorePolicy)(nil)).Elem()
}

func (i SpotRestorePolicyArgs) ToSpotRestorePolicyOutput() SpotRestorePolicyOutput {
	return i.ToSpotRestorePolicyOutputWithContext(context.Background())
}

func (i SpotRestorePolicyArgs) ToSpotRestorePolicyOutputWithContext(ctx context.Context) SpotRestorePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpotRestorePolicyOutput)
}

func (i SpotRestorePolicyArgs) ToSpotRestorePolicyPtrOutput() SpotRestorePolicyPtrOutput {
	return i.ToSpotRestorePolicyPtrOutputWithContext(context.Background())
}

func (i SpotRestorePolicyArgs) ToSpotRestorePolicyPtrOutputWithContext(ctx context.Context) SpotRestorePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpotRestorePolicyOutput).ToSpotRestorePolicyPtrOutputWithContext(ctx)
}









type SpotRestorePolicyPtrInput interface {
	pulumi.Input

	ToSpotRestorePolicyPtrOutput() SpotRestorePolicyPtrOutput
	ToSpotRestorePolicyPtrOutputWithContext(context.Context) SpotRestorePolicyPtrOutput
}

type spotRestorePolicyPtrType SpotRestorePolicyArgs

func SpotRestorePolicyPtr(v *SpotRestorePolicyArgs) SpotRestorePolicyPtrInput {
	return (*spotRestorePolicyPtrType)(v)
}

func (*spotRestorePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SpotRestorePolicy)(nil)).Elem()
}

func (i *spotRestorePolicyPtrType) ToSpotRestorePolicyPtrOutput() SpotRestorePolicyPtrOutput {
	return i.ToSpotRestorePolicyPtrOutputWithContext(context.Background())
}

func (i *spotRestorePolicyPtrType) ToSpotRestorePolicyPtrOutputWithContext(ctx context.Context) SpotRestorePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpotRestorePolicyPtrOutput)
}

type SpotRestorePolicyOutput struct{ *pulumi.OutputState }

func (SpotRestorePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SpotRestorePolicy)(nil)).Elem()
}

func (o SpotRestorePolicyOutput) ToSpotRestorePolicyOutput() SpotRestorePolicyOutput {
	return o
}

func (o SpotRestorePolicyOutput) ToSpotRestorePolicyOutputWithContext(ctx context.Context) SpotRestorePolicyOutput {
	return o
}

func (o SpotRestorePolicyOutput) ToSpotRestorePolicyPtrOutput() SpotRestorePolicyPtrOutput {
	return o.ToSpotRestorePolicyPtrOutputWithContext(context.Background())
}

func (o SpotRestorePolicyOutput) ToSpotRestorePolicyPtrOutputWithContext(ctx context.Context) SpotRestorePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SpotRestorePolicy) *SpotRestorePolicy {
		return &v
	}).(SpotRestorePolicyPtrOutput)
}

func (o SpotRestorePolicyOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SpotRestorePolicy) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o SpotRestorePolicyOutput) RestoreTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SpotRestorePolicy) *string { return v.RestoreTimeout }).(pulumi.StringPtrOutput)
}

type SpotRestorePolicyPtrOutput struct{ *pulumi.OutputState }

func (SpotRestorePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpotRestorePolicy)(nil)).Elem()
}

func (o SpotRestorePolicyPtrOutput) ToSpotRestorePolicyPtrOutput() SpotRestorePolicyPtrOutput {
	return o
}

func (o SpotRestorePolicyPtrOutput) ToSpotRestorePolicyPtrOutputWithContext(ctx context.Context) SpotRestorePolicyPtrOutput {
	return o
}

func (o SpotRestorePolicyPtrOutput) Elem() SpotRestorePolicyOutput {
	return o.ApplyT(func(v *SpotRestorePolicy) SpotRestorePolicy {
		if v != nil {
			return *v
		}
		var ret SpotRestorePolicy
		return ret
	}).(SpotRestorePolicyOutput)
}

func (o SpotRestorePolicyPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SpotRestorePolicy) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o SpotRestorePolicyPtrOutput) RestoreTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SpotRestorePolicy) *string {
		if v == nil {
			return nil
		}
		return v.RestoreTimeout
	}).(pulumi.StringPtrOutput)
}

type SpotRestorePolicyResponse struct {
	Enabled        *bool   `pulumi:"enabled"`
	RestoreTimeout *string `pulumi:"restoreTimeout"`
}

type SpotRestorePolicyResponseOutput struct{ *pulumi.OutputState }

func (SpotRestorePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SpotRestorePolicyResponse)(nil)).Elem()
}

func (o SpotRestorePolicyResponseOutput) ToSpotRestorePolicyResponseOutput() SpotRestorePolicyResponseOutput {
	return o
}

func (o SpotRestorePolicyResponseOutput) ToSpotRestorePolicyResponseOutputWithContext(ctx context.Context) SpotRestorePolicyResponseOutput {
	return o
}

func (o SpotRestorePolicyResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SpotRestorePolicyResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o SpotRestorePolicyResponseOutput) RestoreTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SpotRestorePolicyResponse) *string { return v.RestoreTimeout }).(pulumi.StringPtrOutput)
}

type SpotRestorePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (SpotRestorePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpotRestorePolicyResponse)(nil)).Elem()
}

func (o SpotRestorePolicyResponsePtrOutput) ToSpotRestorePolicyResponsePtrOutput() SpotRestorePolicyResponsePtrOutput {
	return o
}

func (o SpotRestorePolicyResponsePtrOutput) ToSpotRestorePolicyResponsePtrOutputWithContext(ctx context.Context) SpotRestorePolicyResponsePtrOutput {
	return o
}

func (o SpotRestorePolicyResponsePtrOutput) Elem() SpotRestorePolicyResponseOutput {
	return o.ApplyT(func(v *SpotRestorePolicyResponse) SpotRestorePolicyResponse {
		if v != nil {
			return *v
		}
		var ret SpotRestorePolicyResponse
		return ret
	}).(SpotRestorePolicyResponseOutput)
}

func (o SpotRestorePolicyResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SpotRestorePolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o SpotRestorePolicyResponsePtrOutput) RestoreTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SpotRestorePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.RestoreTimeout
	}).(pulumi.StringPtrOutput)
}

type SshConfiguration struct {
	PublicKeys []SshPublicKeyType `pulumi:"publicKeys"`
}





type SshConfigurationInput interface {
	pulumi.Input

	ToSshConfigurationOutput() SshConfigurationOutput
	ToSshConfigurationOutputWithContext(context.Context) SshConfigurationOutput
}

type SshConfigurationArgs struct {
	PublicKeys SshPublicKeyTypeArrayInput `pulumi:"publicKeys"`
}

func (SshConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SshConfiguration)(nil)).Elem()
}

func (i SshConfigurationArgs) ToSshConfigurationOutput() SshConfigurationOutput {
	return i.ToSshConfigurationOutputWithContext(context.Background())
}

func (i SshConfigurationArgs) ToSshConfigurationOutputWithContext(ctx context.Context) SshConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshConfigurationOutput)
}

func (i SshConfigurationArgs) ToSshConfigurationPtrOutput() SshConfigurationPtrOutput {
	return i.ToSshConfigurationPtrOutputWithContext(context.Background())
}

func (i SshConfigurationArgs) ToSshConfigurationPtrOutputWithContext(ctx context.Context) SshConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshConfigurationOutput).ToSshConfigurationPtrOutputWithContext(ctx)
}









type SshConfigurationPtrInput interface {
	pulumi.Input

	ToSshConfigurationPtrOutput() SshConfigurationPtrOutput
	ToSshConfigurationPtrOutputWithContext(context.Context) SshConfigurationPtrOutput
}

type sshConfigurationPtrType SshConfigurationArgs

func SshConfigurationPtr(v *SshConfigurationArgs) SshConfigurationPtrInput {
	return (*sshConfigurationPtrType)(v)
}

func (*sshConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SshConfiguration)(nil)).Elem()
}

func (i *sshConfigurationPtrType) ToSshConfigurationPtrOutput() SshConfigurationPtrOutput {
	return i.ToSshConfigurationPtrOutputWithContext(context.Background())
}

func (i *sshConfigurationPtrType) ToSshConfigurationPtrOutputWithContext(ctx context.Context) SshConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshConfigurationPtrOutput)
}

type SshConfigurationOutput struct{ *pulumi.OutputState }

func (SshConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshConfiguration)(nil)).Elem()
}

func (o SshConfigurationOutput) ToSshConfigurationOutput() SshConfigurationOutput {
	return o
}

func (o SshConfigurationOutput) ToSshConfigurationOutputWithContext(ctx context.Context) SshConfigurationOutput {
	return o
}

func (o SshConfigurationOutput) ToSshConfigurationPtrOutput() SshConfigurationPtrOutput {
	return o.ToSshConfigurationPtrOutputWithContext(context.Background())
}

func (o SshConfigurationOutput) ToSshConfigurationPtrOutputWithContext(ctx context.Context) SshConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SshConfiguration) *SshConfiguration {
		return &v
	}).(SshConfigurationPtrOutput)
}

func (o SshConfigurationOutput) PublicKeys() SshPublicKeyTypeArrayOutput {
	return o.ApplyT(func(v SshConfiguration) []SshPublicKeyType { return v.PublicKeys }).(SshPublicKeyTypeArrayOutput)
}

type SshConfigurationPtrOutput struct{ *pulumi.OutputState }

func (SshConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshConfiguration)(nil)).Elem()
}

func (o SshConfigurationPtrOutput) ToSshConfigurationPtrOutput() SshConfigurationPtrOutput {
	return o
}

func (o SshConfigurationPtrOutput) ToSshConfigurationPtrOutputWithContext(ctx context.Context) SshConfigurationPtrOutput {
	return o
}

func (o SshConfigurationPtrOutput) Elem() SshConfigurationOutput {
	return o.ApplyT(func(v *SshConfiguration) SshConfiguration {
		if v != nil {
			return *v
		}
		var ret SshConfiguration
		return ret
	}).(SshConfigurationOutput)
}

func (o SshConfigurationPtrOutput) PublicKeys() SshPublicKeyTypeArrayOutput {
	return o.ApplyT(func(v *SshConfiguration) []SshPublicKeyType {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(SshPublicKeyTypeArrayOutput)
}

type SshConfigurationResponse struct {
	PublicKeys []SshPublicKeyResponse `pulumi:"publicKeys"`
}

type SshConfigurationResponseOutput struct{ *pulumi.OutputState }

func (SshConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshConfigurationResponse)(nil)).Elem()
}

func (o SshConfigurationResponseOutput) ToSshConfigurationResponseOutput() SshConfigurationResponseOutput {
	return o
}

func (o SshConfigurationResponseOutput) ToSshConfigurationResponseOutputWithContext(ctx context.Context) SshConfigurationResponseOutput {
	return o
}

func (o SshConfigurationResponseOutput) PublicKeys() SshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v SshConfigurationResponse) []SshPublicKeyResponse { return v.PublicKeys }).(SshPublicKeyResponseArrayOutput)
}

type SshConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (SshConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshConfigurationResponse)(nil)).Elem()
}

func (o SshConfigurationResponsePtrOutput) ToSshConfigurationResponsePtrOutput() SshConfigurationResponsePtrOutput {
	return o
}

func (o SshConfigurationResponsePtrOutput) ToSshConfigurationResponsePtrOutputWithContext(ctx context.Context) SshConfigurationResponsePtrOutput {
	return o
}

func (o SshConfigurationResponsePtrOutput) Elem() SshConfigurationResponseOutput {
	return o.ApplyT(func(v *SshConfigurationResponse) SshConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret SshConfigurationResponse
		return ret
	}).(SshConfigurationResponseOutput)
}

func (o SshConfigurationResponsePtrOutput) PublicKeys() SshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v *SshConfigurationResponse) []SshPublicKeyResponse {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(SshPublicKeyResponseArrayOutput)
}

type SshPublicKeyType struct {
	KeyData *string `pulumi:"keyData"`
	Path    *string `pulumi:"path"`
}





type SshPublicKeyTypeInput interface {
	pulumi.Input

	ToSshPublicKeyTypeOutput() SshPublicKeyTypeOutput
	ToSshPublicKeyTypeOutputWithContext(context.Context) SshPublicKeyTypeOutput
}

type SshPublicKeyTypeArgs struct {
	KeyData pulumi.StringPtrInput `pulumi:"keyData"`
	Path    pulumi.StringPtrInput `pulumi:"path"`
}

func (SshPublicKeyTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicKeyType)(nil)).Elem()
}

func (i SshPublicKeyTypeArgs) ToSshPublicKeyTypeOutput() SshPublicKeyTypeOutput {
	return i.ToSshPublicKeyTypeOutputWithContext(context.Background())
}

func (i SshPublicKeyTypeArgs) ToSshPublicKeyTypeOutputWithContext(ctx context.Context) SshPublicKeyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyTypeOutput)
}





type SshPublicKeyTypeArrayInput interface {
	pulumi.Input

	ToSshPublicKeyTypeArrayOutput() SshPublicKeyTypeArrayOutput
	ToSshPublicKeyTypeArrayOutputWithContext(context.Context) SshPublicKeyTypeArrayOutput
}

type SshPublicKeyTypeArray []SshPublicKeyTypeInput

func (SshPublicKeyTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshPublicKeyType)(nil)).Elem()
}

func (i SshPublicKeyTypeArray) ToSshPublicKeyTypeArrayOutput() SshPublicKeyTypeArrayOutput {
	return i.ToSshPublicKeyTypeArrayOutputWithContext(context.Background())
}

func (i SshPublicKeyTypeArray) ToSshPublicKeyTypeArrayOutputWithContext(ctx context.Context) SshPublicKeyTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyTypeArrayOutput)
}

type SshPublicKeyTypeOutput struct{ *pulumi.OutputState }

func (SshPublicKeyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicKeyType)(nil)).Elem()
}

func (o SshPublicKeyTypeOutput) ToSshPublicKeyTypeOutput() SshPublicKeyTypeOutput {
	return o
}

func (o SshPublicKeyTypeOutput) ToSshPublicKeyTypeOutputWithContext(ctx context.Context) SshPublicKeyTypeOutput {
	return o
}

func (o SshPublicKeyTypeOutput) KeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SshPublicKeyType) *string { return v.KeyData }).(pulumi.StringPtrOutput)
}

func (o SshPublicKeyTypeOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SshPublicKeyType) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type SshPublicKeyTypeArrayOutput struct{ *pulumi.OutputState }

func (SshPublicKeyTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshPublicKeyType)(nil)).Elem()
}

func (o SshPublicKeyTypeArrayOutput) ToSshPublicKeyTypeArrayOutput() SshPublicKeyTypeArrayOutput {
	return o
}

func (o SshPublicKeyTypeArrayOutput) ToSshPublicKeyTypeArrayOutputWithContext(ctx context.Context) SshPublicKeyTypeArrayOutput {
	return o
}

func (o SshPublicKeyTypeArrayOutput) Index(i pulumi.IntInput) SshPublicKeyTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SshPublicKeyType {
		return vs[0].([]SshPublicKeyType)[vs[1].(int)]
	}).(SshPublicKeyTypeOutput)
}

type SshPublicKeyResponse struct {
	KeyData *string `pulumi:"keyData"`
	Path    *string `pulumi:"path"`
}

type SshPublicKeyResponseOutput struct{ *pulumi.OutputState }

func (SshPublicKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicKeyResponse)(nil)).Elem()
}

func (o SshPublicKeyResponseOutput) ToSshPublicKeyResponseOutput() SshPublicKeyResponseOutput {
	return o
}

func (o SshPublicKeyResponseOutput) ToSshPublicKeyResponseOutputWithContext(ctx context.Context) SshPublicKeyResponseOutput {
	return o
}

func (o SshPublicKeyResponseOutput) KeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SshPublicKeyResponse) *string { return v.KeyData }).(pulumi.StringPtrOutput)
}

func (o SshPublicKeyResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SshPublicKeyResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type SshPublicKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (SshPublicKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshPublicKeyResponse)(nil)).Elem()
}

func (o SshPublicKeyResponseArrayOutput) ToSshPublicKeyResponseArrayOutput() SshPublicKeyResponseArrayOutput {
	return o
}

func (o SshPublicKeyResponseArrayOutput) ToSshPublicKeyResponseArrayOutputWithContext(ctx context.Context) SshPublicKeyResponseArrayOutput {
	return o
}

func (o SshPublicKeyResponseArrayOutput) Index(i pulumi.IntInput) SshPublicKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SshPublicKeyResponse {
		return vs[0].([]SshPublicKeyResponse)[vs[1].(int)]
	}).(SshPublicKeyResponseOutput)
}

type StorageProfile struct {
	DataDisks      []DataDisk      `pulumi:"dataDisks"`
	ImageReference *ImageReference `pulumi:"imageReference"`
	OsDisk         *OSDisk         `pulumi:"osDisk"`
}





type StorageProfileInput interface {
	pulumi.Input

	ToStorageProfileOutput() StorageProfileOutput
	ToStorageProfileOutputWithContext(context.Context) StorageProfileOutput
}

type StorageProfileArgs struct {
	DataDisks      DataDiskArrayInput     `pulumi:"dataDisks"`
	ImageReference ImageReferencePtrInput `pulumi:"imageReference"`
	OsDisk         OSDiskPtrInput         `pulumi:"osDisk"`
}

func (StorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (i StorageProfileArgs) ToStorageProfileOutput() StorageProfileOutput {
	return i.ToStorageProfileOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput)
}

func (i StorageProfileArgs) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput).ToStorageProfilePtrOutputWithContext(ctx)
}









type StorageProfilePtrInput interface {
	pulumi.Input

	ToStorageProfilePtrOutput() StorageProfilePtrOutput
	ToStorageProfilePtrOutputWithContext(context.Context) StorageProfilePtrOutput
}

type storageProfilePtrType StorageProfileArgs

func StorageProfilePtr(v *StorageProfileArgs) StorageProfilePtrInput {
	return (*storageProfilePtrType)(v)
}

func (*storageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfilePtrOutput)
}

type StorageProfileOutput struct{ *pulumi.OutputState }

func (StorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (o StorageProfileOutput) ToStorageProfileOutput() StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (o StorageProfileOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageProfile) *StorageProfile {
		return &v
	}).(StorageProfilePtrOutput)
}

func (o StorageProfileOutput) DataDisks() DataDiskArrayOutput {
	return o.ApplyT(func(v StorageProfile) []DataDisk { return v.DataDisks }).(DataDiskArrayOutput)
}

func (o StorageProfileOutput) ImageReference() ImageReferencePtrOutput {
	return o.ApplyT(func(v StorageProfile) *ImageReference { return v.ImageReference }).(ImageReferencePtrOutput)
}

func (o StorageProfileOutput) OsDisk() OSDiskPtrOutput {
	return o.ApplyT(func(v StorageProfile) *OSDisk { return v.OsDisk }).(OSDiskPtrOutput)
}

type StorageProfilePtrOutput struct{ *pulumi.OutputState }

func (StorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) Elem() StorageProfileOutput {
	return o.ApplyT(func(v *StorageProfile) StorageProfile {
		if v != nil {
			return *v
		}
		var ret StorageProfile
		return ret
	}).(StorageProfileOutput)
}

func (o StorageProfilePtrOutput) DataDisks() DataDiskArrayOutput {
	return o.ApplyT(func(v *StorageProfile) []DataDisk {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(DataDiskArrayOutput)
}

func (o StorageProfilePtrOutput) ImageReference() ImageReferencePtrOutput {
	return o.ApplyT(func(v *StorageProfile) *ImageReference {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(ImageReferencePtrOutput)
}

func (o StorageProfilePtrOutput) OsDisk() OSDiskPtrOutput {
	return o.ApplyT(func(v *StorageProfile) *OSDisk {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(OSDiskPtrOutput)
}

type StorageProfileResponse struct {
	DataDisks      []DataDiskResponse      `pulumi:"dataDisks"`
	ImageReference *ImageReferenceResponse `pulumi:"imageReference"`
	OsDisk         *OSDiskResponse         `pulumi:"osDisk"`
}

type StorageProfileResponseOutput struct{ *pulumi.OutputState }

func (StorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutput() StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutputWithContext(ctx context.Context) StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) DataDisks() DataDiskResponseArrayOutput {
	return o.ApplyT(func(v StorageProfileResponse) []DataDiskResponse { return v.DataDisks }).(DataDiskResponseArrayOutput)
}

func (o StorageProfileResponseOutput) ImageReference() ImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *ImageReferenceResponse { return v.ImageReference }).(ImageReferenceResponsePtrOutput)
}

func (o StorageProfileResponseOutput) OsDisk() OSDiskResponsePtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *OSDiskResponse { return v.OsDisk }).(OSDiskResponsePtrOutput)
}

type StorageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) Elem() StorageProfileResponseOutput {
	return o.ApplyT(func(v *StorageProfileResponse) StorageProfileResponse {
		if v != nil {
			return *v
		}
		var ret StorageProfileResponse
		return ret
	}).(StorageProfileResponseOutput)
}

func (o StorageProfileResponsePtrOutput) DataDisks() DataDiskResponseArrayOutput {
	return o.ApplyT(func(v *StorageProfileResponse) []DataDiskResponse {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(DataDiskResponseArrayOutput)
}

func (o StorageProfileResponsePtrOutput) ImageReference() ImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *ImageReferenceResponse {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(ImageReferenceResponsePtrOutput)
}

func (o StorageProfileResponsePtrOutput) OsDisk() OSDiskResponsePtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *OSDiskResponse {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(OSDiskResponsePtrOutput)
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

type SubResourceReadOnlyResponse struct {
	Id string `pulumi:"id"`
}

type SubResourceReadOnlyResponseOutput struct{ *pulumi.OutputState }

func (SubResourceReadOnlyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceReadOnlyResponse)(nil)).Elem()
}

func (o SubResourceReadOnlyResponseOutput) ToSubResourceReadOnlyResponseOutput() SubResourceReadOnlyResponseOutput {
	return o
}

func (o SubResourceReadOnlyResponseOutput) ToSubResourceReadOnlyResponseOutputWithContext(ctx context.Context) SubResourceReadOnlyResponseOutput {
	return o
}

func (o SubResourceReadOnlyResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SubResourceReadOnlyResponse) string { return v.Id }).(pulumi.StringOutput)
}

type SubResourceReadOnlyResponseArrayOutput struct{ *pulumi.OutputState }

func (SubResourceReadOnlyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResourceReadOnlyResponse)(nil)).Elem()
}

func (o SubResourceReadOnlyResponseArrayOutput) ToSubResourceReadOnlyResponseArrayOutput() SubResourceReadOnlyResponseArrayOutput {
	return o
}

func (o SubResourceReadOnlyResponseArrayOutput) ToSubResourceReadOnlyResponseArrayOutputWithContext(ctx context.Context) SubResourceReadOnlyResponseArrayOutput {
	return o
}

func (o SubResourceReadOnlyResponseArrayOutput) Index(i pulumi.IntInput) SubResourceReadOnlyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResourceReadOnlyResponse {
		return vs[0].([]SubResourceReadOnlyResponse)[vs[1].(int)]
	}).(SubResourceReadOnlyResponseOutput)
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

type SubResourceWithColocationStatusResponse struct {
	ColocationStatus *InstanceViewStatusResponse `pulumi:"colocationStatus"`
	Id               *string                     `pulumi:"id"`
}

type SubResourceWithColocationStatusResponseOutput struct{ *pulumi.OutputState }

func (SubResourceWithColocationStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceWithColocationStatusResponse)(nil)).Elem()
}

func (o SubResourceWithColocationStatusResponseOutput) ToSubResourceWithColocationStatusResponseOutput() SubResourceWithColocationStatusResponseOutput {
	return o
}

func (o SubResourceWithColocationStatusResponseOutput) ToSubResourceWithColocationStatusResponseOutputWithContext(ctx context.Context) SubResourceWithColocationStatusResponseOutput {
	return o
}

func (o SubResourceWithColocationStatusResponseOutput) ColocationStatus() InstanceViewStatusResponsePtrOutput {
	return o.ApplyT(func(v SubResourceWithColocationStatusResponse) *InstanceViewStatusResponse { return v.ColocationStatus }).(InstanceViewStatusResponsePtrOutput)
}

func (o SubResourceWithColocationStatusResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResourceWithColocationStatusResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourceWithColocationStatusResponseArrayOutput struct{ *pulumi.OutputState }

func (SubResourceWithColocationStatusResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResourceWithColocationStatusResponse)(nil)).Elem()
}

func (o SubResourceWithColocationStatusResponseArrayOutput) ToSubResourceWithColocationStatusResponseArrayOutput() SubResourceWithColocationStatusResponseArrayOutput {
	return o
}

func (o SubResourceWithColocationStatusResponseArrayOutput) ToSubResourceWithColocationStatusResponseArrayOutputWithContext(ctx context.Context) SubResourceWithColocationStatusResponseArrayOutput {
	return o
}

func (o SubResourceWithColocationStatusResponseArrayOutput) Index(i pulumi.IntInput) SubResourceWithColocationStatusResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResourceWithColocationStatusResponse {
		return vs[0].([]SubResourceWithColocationStatusResponse)[vs[1].(int)]
	}).(SubResourceWithColocationStatusResponseOutput)
}

type TerminateNotificationProfile struct {
	Enable           *bool   `pulumi:"enable"`
	NotBeforeTimeout *string `pulumi:"notBeforeTimeout"`
}





type TerminateNotificationProfileInput interface {
	pulumi.Input

	ToTerminateNotificationProfileOutput() TerminateNotificationProfileOutput
	ToTerminateNotificationProfileOutputWithContext(context.Context) TerminateNotificationProfileOutput
}

type TerminateNotificationProfileArgs struct {
	Enable           pulumi.BoolPtrInput   `pulumi:"enable"`
	NotBeforeTimeout pulumi.StringPtrInput `pulumi:"notBeforeTimeout"`
}

func (TerminateNotificationProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TerminateNotificationProfile)(nil)).Elem()
}

func (i TerminateNotificationProfileArgs) ToTerminateNotificationProfileOutput() TerminateNotificationProfileOutput {
	return i.ToTerminateNotificationProfileOutputWithContext(context.Background())
}

func (i TerminateNotificationProfileArgs) ToTerminateNotificationProfileOutputWithContext(ctx context.Context) TerminateNotificationProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerminateNotificationProfileOutput)
}

func (i TerminateNotificationProfileArgs) ToTerminateNotificationProfilePtrOutput() TerminateNotificationProfilePtrOutput {
	return i.ToTerminateNotificationProfilePtrOutputWithContext(context.Background())
}

func (i TerminateNotificationProfileArgs) ToTerminateNotificationProfilePtrOutputWithContext(ctx context.Context) TerminateNotificationProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerminateNotificationProfileOutput).ToTerminateNotificationProfilePtrOutputWithContext(ctx)
}









type TerminateNotificationProfilePtrInput interface {
	pulumi.Input

	ToTerminateNotificationProfilePtrOutput() TerminateNotificationProfilePtrOutput
	ToTerminateNotificationProfilePtrOutputWithContext(context.Context) TerminateNotificationProfilePtrOutput
}

type terminateNotificationProfilePtrType TerminateNotificationProfileArgs

func TerminateNotificationProfilePtr(v *TerminateNotificationProfileArgs) TerminateNotificationProfilePtrInput {
	return (*terminateNotificationProfilePtrType)(v)
}

func (*terminateNotificationProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TerminateNotificationProfile)(nil)).Elem()
}

func (i *terminateNotificationProfilePtrType) ToTerminateNotificationProfilePtrOutput() TerminateNotificationProfilePtrOutput {
	return i.ToTerminateNotificationProfilePtrOutputWithContext(context.Background())
}

func (i *terminateNotificationProfilePtrType) ToTerminateNotificationProfilePtrOutputWithContext(ctx context.Context) TerminateNotificationProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TerminateNotificationProfilePtrOutput)
}

type TerminateNotificationProfileOutput struct{ *pulumi.OutputState }

func (TerminateNotificationProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TerminateNotificationProfile)(nil)).Elem()
}

func (o TerminateNotificationProfileOutput) ToTerminateNotificationProfileOutput() TerminateNotificationProfileOutput {
	return o
}

func (o TerminateNotificationProfileOutput) ToTerminateNotificationProfileOutputWithContext(ctx context.Context) TerminateNotificationProfileOutput {
	return o
}

func (o TerminateNotificationProfileOutput) ToTerminateNotificationProfilePtrOutput() TerminateNotificationProfilePtrOutput {
	return o.ToTerminateNotificationProfilePtrOutputWithContext(context.Background())
}

func (o TerminateNotificationProfileOutput) ToTerminateNotificationProfilePtrOutputWithContext(ctx context.Context) TerminateNotificationProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TerminateNotificationProfile) *TerminateNotificationProfile {
		return &v
	}).(TerminateNotificationProfilePtrOutput)
}

func (o TerminateNotificationProfileOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TerminateNotificationProfile) *bool { return v.Enable }).(pulumi.BoolPtrOutput)
}

func (o TerminateNotificationProfileOutput) NotBeforeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TerminateNotificationProfile) *string { return v.NotBeforeTimeout }).(pulumi.StringPtrOutput)
}

type TerminateNotificationProfilePtrOutput struct{ *pulumi.OutputState }

func (TerminateNotificationProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TerminateNotificationProfile)(nil)).Elem()
}

func (o TerminateNotificationProfilePtrOutput) ToTerminateNotificationProfilePtrOutput() TerminateNotificationProfilePtrOutput {
	return o
}

func (o TerminateNotificationProfilePtrOutput) ToTerminateNotificationProfilePtrOutputWithContext(ctx context.Context) TerminateNotificationProfilePtrOutput {
	return o
}

func (o TerminateNotificationProfilePtrOutput) Elem() TerminateNotificationProfileOutput {
	return o.ApplyT(func(v *TerminateNotificationProfile) TerminateNotificationProfile {
		if v != nil {
			return *v
		}
		var ret TerminateNotificationProfile
		return ret
	}).(TerminateNotificationProfileOutput)
}

func (o TerminateNotificationProfilePtrOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TerminateNotificationProfile) *bool {
		if v == nil {
			return nil
		}
		return v.Enable
	}).(pulumi.BoolPtrOutput)
}

func (o TerminateNotificationProfilePtrOutput) NotBeforeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TerminateNotificationProfile) *string {
		if v == nil {
			return nil
		}
		return v.NotBeforeTimeout
	}).(pulumi.StringPtrOutput)
}

type TerminateNotificationProfileResponse struct {
	Enable           *bool   `pulumi:"enable"`
	NotBeforeTimeout *string `pulumi:"notBeforeTimeout"`
}

type TerminateNotificationProfileResponseOutput struct{ *pulumi.OutputState }

func (TerminateNotificationProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TerminateNotificationProfileResponse)(nil)).Elem()
}

func (o TerminateNotificationProfileResponseOutput) ToTerminateNotificationProfileResponseOutput() TerminateNotificationProfileResponseOutput {
	return o
}

func (o TerminateNotificationProfileResponseOutput) ToTerminateNotificationProfileResponseOutputWithContext(ctx context.Context) TerminateNotificationProfileResponseOutput {
	return o
}

func (o TerminateNotificationProfileResponseOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TerminateNotificationProfileResponse) *bool { return v.Enable }).(pulumi.BoolPtrOutput)
}

func (o TerminateNotificationProfileResponseOutput) NotBeforeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TerminateNotificationProfileResponse) *string { return v.NotBeforeTimeout }).(pulumi.StringPtrOutput)
}

type TerminateNotificationProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (TerminateNotificationProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TerminateNotificationProfileResponse)(nil)).Elem()
}

func (o TerminateNotificationProfileResponsePtrOutput) ToTerminateNotificationProfileResponsePtrOutput() TerminateNotificationProfileResponsePtrOutput {
	return o
}

func (o TerminateNotificationProfileResponsePtrOutput) ToTerminateNotificationProfileResponsePtrOutputWithContext(ctx context.Context) TerminateNotificationProfileResponsePtrOutput {
	return o
}

func (o TerminateNotificationProfileResponsePtrOutput) Elem() TerminateNotificationProfileResponseOutput {
	return o.ApplyT(func(v *TerminateNotificationProfileResponse) TerminateNotificationProfileResponse {
		if v != nil {
			return *v
		}
		var ret TerminateNotificationProfileResponse
		return ret
	}).(TerminateNotificationProfileResponseOutput)
}

func (o TerminateNotificationProfileResponsePtrOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TerminateNotificationProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enable
	}).(pulumi.BoolPtrOutput)
}

func (o TerminateNotificationProfileResponsePtrOutput) NotBeforeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TerminateNotificationProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.NotBeforeTimeout
	}).(pulumi.StringPtrOutput)
}

type UefiSettings struct {
	SecureBootEnabled *bool `pulumi:"secureBootEnabled"`
	VTpmEnabled       *bool `pulumi:"vTpmEnabled"`
}





type UefiSettingsInput interface {
	pulumi.Input

	ToUefiSettingsOutput() UefiSettingsOutput
	ToUefiSettingsOutputWithContext(context.Context) UefiSettingsOutput
}

type UefiSettingsArgs struct {
	SecureBootEnabled pulumi.BoolPtrInput `pulumi:"secureBootEnabled"`
	VTpmEnabled       pulumi.BoolPtrInput `pulumi:"vTpmEnabled"`
}

func (UefiSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UefiSettings)(nil)).Elem()
}

func (i UefiSettingsArgs) ToUefiSettingsOutput() UefiSettingsOutput {
	return i.ToUefiSettingsOutputWithContext(context.Background())
}

func (i UefiSettingsArgs) ToUefiSettingsOutputWithContext(ctx context.Context) UefiSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UefiSettingsOutput)
}

func (i UefiSettingsArgs) ToUefiSettingsPtrOutput() UefiSettingsPtrOutput {
	return i.ToUefiSettingsPtrOutputWithContext(context.Background())
}

func (i UefiSettingsArgs) ToUefiSettingsPtrOutputWithContext(ctx context.Context) UefiSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UefiSettingsOutput).ToUefiSettingsPtrOutputWithContext(ctx)
}









type UefiSettingsPtrInput interface {
	pulumi.Input

	ToUefiSettingsPtrOutput() UefiSettingsPtrOutput
	ToUefiSettingsPtrOutputWithContext(context.Context) UefiSettingsPtrOutput
}

type uefiSettingsPtrType UefiSettingsArgs

func UefiSettingsPtr(v *UefiSettingsArgs) UefiSettingsPtrInput {
	return (*uefiSettingsPtrType)(v)
}

func (*uefiSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UefiSettings)(nil)).Elem()
}

func (i *uefiSettingsPtrType) ToUefiSettingsPtrOutput() UefiSettingsPtrOutput {
	return i.ToUefiSettingsPtrOutputWithContext(context.Background())
}

func (i *uefiSettingsPtrType) ToUefiSettingsPtrOutputWithContext(ctx context.Context) UefiSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UefiSettingsPtrOutput)
}

type UefiSettingsOutput struct{ *pulumi.OutputState }

func (UefiSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UefiSettings)(nil)).Elem()
}

func (o UefiSettingsOutput) ToUefiSettingsOutput() UefiSettingsOutput {
	return o
}

func (o UefiSettingsOutput) ToUefiSettingsOutputWithContext(ctx context.Context) UefiSettingsOutput {
	return o
}

func (o UefiSettingsOutput) ToUefiSettingsPtrOutput() UefiSettingsPtrOutput {
	return o.ToUefiSettingsPtrOutputWithContext(context.Background())
}

func (o UefiSettingsOutput) ToUefiSettingsPtrOutputWithContext(ctx context.Context) UefiSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UefiSettings) *UefiSettings {
		return &v
	}).(UefiSettingsPtrOutput)
}

func (o UefiSettingsOutput) SecureBootEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UefiSettings) *bool { return v.SecureBootEnabled }).(pulumi.BoolPtrOutput)
}

func (o UefiSettingsOutput) VTpmEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UefiSettings) *bool { return v.VTpmEnabled }).(pulumi.BoolPtrOutput)
}

type UefiSettingsPtrOutput struct{ *pulumi.OutputState }

func (UefiSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UefiSettings)(nil)).Elem()
}

func (o UefiSettingsPtrOutput) ToUefiSettingsPtrOutput() UefiSettingsPtrOutput {
	return o
}

func (o UefiSettingsPtrOutput) ToUefiSettingsPtrOutputWithContext(ctx context.Context) UefiSettingsPtrOutput {
	return o
}

func (o UefiSettingsPtrOutput) Elem() UefiSettingsOutput {
	return o.ApplyT(func(v *UefiSettings) UefiSettings {
		if v != nil {
			return *v
		}
		var ret UefiSettings
		return ret
	}).(UefiSettingsOutput)
}

func (o UefiSettingsPtrOutput) SecureBootEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UefiSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SecureBootEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o UefiSettingsPtrOutput) VTpmEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UefiSettings) *bool {
		if v == nil {
			return nil
		}
		return v.VTpmEnabled
	}).(pulumi.BoolPtrOutput)
}

type UefiSettingsResponse struct {
	SecureBootEnabled *bool `pulumi:"secureBootEnabled"`
	VTpmEnabled       *bool `pulumi:"vTpmEnabled"`
}

type UefiSettingsResponseOutput struct{ *pulumi.OutputState }

func (UefiSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UefiSettingsResponse)(nil)).Elem()
}

func (o UefiSettingsResponseOutput) ToUefiSettingsResponseOutput() UefiSettingsResponseOutput {
	return o
}

func (o UefiSettingsResponseOutput) ToUefiSettingsResponseOutputWithContext(ctx context.Context) UefiSettingsResponseOutput {
	return o
}

func (o UefiSettingsResponseOutput) SecureBootEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UefiSettingsResponse) *bool { return v.SecureBootEnabled }).(pulumi.BoolPtrOutput)
}

func (o UefiSettingsResponseOutput) VTpmEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UefiSettingsResponse) *bool { return v.VTpmEnabled }).(pulumi.BoolPtrOutput)
}

type UefiSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (UefiSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UefiSettingsResponse)(nil)).Elem()
}

func (o UefiSettingsResponsePtrOutput) ToUefiSettingsResponsePtrOutput() UefiSettingsResponsePtrOutput {
	return o
}

func (o UefiSettingsResponsePtrOutput) ToUefiSettingsResponsePtrOutputWithContext(ctx context.Context) UefiSettingsResponsePtrOutput {
	return o
}

func (o UefiSettingsResponsePtrOutput) Elem() UefiSettingsResponseOutput {
	return o.ApplyT(func(v *UefiSettingsResponse) UefiSettingsResponse {
		if v != nil {
			return *v
		}
		var ret UefiSettingsResponse
		return ret
	}).(UefiSettingsResponseOutput)
}

func (o UefiSettingsResponsePtrOutput) SecureBootEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UefiSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SecureBootEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o UefiSettingsResponsePtrOutput) VTpmEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UefiSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.VTpmEnabled
	}).(pulumi.BoolPtrOutput)
}

type UpgradePolicy struct {
	AutomaticOSUpgradePolicy *AutomaticOSUpgradePolicy `pulumi:"automaticOSUpgradePolicy"`
	Mode                     *UpgradeMode              `pulumi:"mode"`
	RollingUpgradePolicy     *RollingUpgradePolicy     `pulumi:"rollingUpgradePolicy"`
}





type UpgradePolicyInput interface {
	pulumi.Input

	ToUpgradePolicyOutput() UpgradePolicyOutput
	ToUpgradePolicyOutputWithContext(context.Context) UpgradePolicyOutput
}

type UpgradePolicyArgs struct {
	AutomaticOSUpgradePolicy AutomaticOSUpgradePolicyPtrInput `pulumi:"automaticOSUpgradePolicy"`
	Mode                     UpgradeModePtrInput              `pulumi:"mode"`
	RollingUpgradePolicy     RollingUpgradePolicyPtrInput     `pulumi:"rollingUpgradePolicy"`
}

func (UpgradePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UpgradePolicy)(nil)).Elem()
}

func (i UpgradePolicyArgs) ToUpgradePolicyOutput() UpgradePolicyOutput {
	return i.ToUpgradePolicyOutputWithContext(context.Background())
}

func (i UpgradePolicyArgs) ToUpgradePolicyOutputWithContext(ctx context.Context) UpgradePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradePolicyOutput)
}

func (i UpgradePolicyArgs) ToUpgradePolicyPtrOutput() UpgradePolicyPtrOutput {
	return i.ToUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i UpgradePolicyArgs) ToUpgradePolicyPtrOutputWithContext(ctx context.Context) UpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradePolicyOutput).ToUpgradePolicyPtrOutputWithContext(ctx)
}









type UpgradePolicyPtrInput interface {
	pulumi.Input

	ToUpgradePolicyPtrOutput() UpgradePolicyPtrOutput
	ToUpgradePolicyPtrOutputWithContext(context.Context) UpgradePolicyPtrOutput
}

type upgradePolicyPtrType UpgradePolicyArgs

func UpgradePolicyPtr(v *UpgradePolicyArgs) UpgradePolicyPtrInput {
	return (*upgradePolicyPtrType)(v)
}

func (*upgradePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradePolicy)(nil)).Elem()
}

func (i *upgradePolicyPtrType) ToUpgradePolicyPtrOutput() UpgradePolicyPtrOutput {
	return i.ToUpgradePolicyPtrOutputWithContext(context.Background())
}

func (i *upgradePolicyPtrType) ToUpgradePolicyPtrOutputWithContext(ctx context.Context) UpgradePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradePolicyPtrOutput)
}

type UpgradePolicyOutput struct{ *pulumi.OutputState }

func (UpgradePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpgradePolicy)(nil)).Elem()
}

func (o UpgradePolicyOutput) ToUpgradePolicyOutput() UpgradePolicyOutput {
	return o
}

func (o UpgradePolicyOutput) ToUpgradePolicyOutputWithContext(ctx context.Context) UpgradePolicyOutput {
	return o
}

func (o UpgradePolicyOutput) ToUpgradePolicyPtrOutput() UpgradePolicyPtrOutput {
	return o.ToUpgradePolicyPtrOutputWithContext(context.Background())
}

func (o UpgradePolicyOutput) ToUpgradePolicyPtrOutputWithContext(ctx context.Context) UpgradePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UpgradePolicy) *UpgradePolicy {
		return &v
	}).(UpgradePolicyPtrOutput)
}

func (o UpgradePolicyOutput) AutomaticOSUpgradePolicy() AutomaticOSUpgradePolicyPtrOutput {
	return o.ApplyT(func(v UpgradePolicy) *AutomaticOSUpgradePolicy { return v.AutomaticOSUpgradePolicy }).(AutomaticOSUpgradePolicyPtrOutput)
}

func (o UpgradePolicyOutput) Mode() UpgradeModePtrOutput {
	return o.ApplyT(func(v UpgradePolicy) *UpgradeMode { return v.Mode }).(UpgradeModePtrOutput)
}

func (o UpgradePolicyOutput) RollingUpgradePolicy() RollingUpgradePolicyPtrOutput {
	return o.ApplyT(func(v UpgradePolicy) *RollingUpgradePolicy { return v.RollingUpgradePolicy }).(RollingUpgradePolicyPtrOutput)
}

type UpgradePolicyPtrOutput struct{ *pulumi.OutputState }

func (UpgradePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradePolicy)(nil)).Elem()
}

func (o UpgradePolicyPtrOutput) ToUpgradePolicyPtrOutput() UpgradePolicyPtrOutput {
	return o
}

func (o UpgradePolicyPtrOutput) ToUpgradePolicyPtrOutputWithContext(ctx context.Context) UpgradePolicyPtrOutput {
	return o
}

func (o UpgradePolicyPtrOutput) Elem() UpgradePolicyOutput {
	return o.ApplyT(func(v *UpgradePolicy) UpgradePolicy {
		if v != nil {
			return *v
		}
		var ret UpgradePolicy
		return ret
	}).(UpgradePolicyOutput)
}

func (o UpgradePolicyPtrOutput) AutomaticOSUpgradePolicy() AutomaticOSUpgradePolicyPtrOutput {
	return o.ApplyT(func(v *UpgradePolicy) *AutomaticOSUpgradePolicy {
		if v == nil {
			return nil
		}
		return v.AutomaticOSUpgradePolicy
	}).(AutomaticOSUpgradePolicyPtrOutput)
}

func (o UpgradePolicyPtrOutput) Mode() UpgradeModePtrOutput {
	return o.ApplyT(func(v *UpgradePolicy) *UpgradeMode {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(UpgradeModePtrOutput)
}

func (o UpgradePolicyPtrOutput) RollingUpgradePolicy() RollingUpgradePolicyPtrOutput {
	return o.ApplyT(func(v *UpgradePolicy) *RollingUpgradePolicy {
		if v == nil {
			return nil
		}
		return v.RollingUpgradePolicy
	}).(RollingUpgradePolicyPtrOutput)
}

type UpgradePolicyResponse struct {
	AutomaticOSUpgradePolicy *AutomaticOSUpgradePolicyResponse `pulumi:"automaticOSUpgradePolicy"`
	Mode                     *string                           `pulumi:"mode"`
	RollingUpgradePolicy     *RollingUpgradePolicyResponse     `pulumi:"rollingUpgradePolicy"`
}

type UpgradePolicyResponseOutput struct{ *pulumi.OutputState }

func (UpgradePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpgradePolicyResponse)(nil)).Elem()
}

func (o UpgradePolicyResponseOutput) ToUpgradePolicyResponseOutput() UpgradePolicyResponseOutput {
	return o
}

func (o UpgradePolicyResponseOutput) ToUpgradePolicyResponseOutputWithContext(ctx context.Context) UpgradePolicyResponseOutput {
	return o
}

func (o UpgradePolicyResponseOutput) AutomaticOSUpgradePolicy() AutomaticOSUpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v UpgradePolicyResponse) *AutomaticOSUpgradePolicyResponse { return v.AutomaticOSUpgradePolicy }).(AutomaticOSUpgradePolicyResponsePtrOutput)
}

func (o UpgradePolicyResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpgradePolicyResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o UpgradePolicyResponseOutput) RollingUpgradePolicy() RollingUpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v UpgradePolicyResponse) *RollingUpgradePolicyResponse { return v.RollingUpgradePolicy }).(RollingUpgradePolicyResponsePtrOutput)
}

type UpgradePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (UpgradePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradePolicyResponse)(nil)).Elem()
}

func (o UpgradePolicyResponsePtrOutput) ToUpgradePolicyResponsePtrOutput() UpgradePolicyResponsePtrOutput {
	return o
}

func (o UpgradePolicyResponsePtrOutput) ToUpgradePolicyResponsePtrOutputWithContext(ctx context.Context) UpgradePolicyResponsePtrOutput {
	return o
}

func (o UpgradePolicyResponsePtrOutput) Elem() UpgradePolicyResponseOutput {
	return o.ApplyT(func(v *UpgradePolicyResponse) UpgradePolicyResponse {
		if v != nil {
			return *v
		}
		var ret UpgradePolicyResponse
		return ret
	}).(UpgradePolicyResponseOutput)
}

func (o UpgradePolicyResponsePtrOutput) AutomaticOSUpgradePolicy() AutomaticOSUpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v *UpgradePolicyResponse) *AutomaticOSUpgradePolicyResponse {
		if v == nil {
			return nil
		}
		return v.AutomaticOSUpgradePolicy
	}).(AutomaticOSUpgradePolicyResponsePtrOutput)
}

func (o UpgradePolicyResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

func (o UpgradePolicyResponsePtrOutput) RollingUpgradePolicy() RollingUpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v *UpgradePolicyResponse) *RollingUpgradePolicyResponse {
		if v == nil {
			return nil
		}
		return v.RollingUpgradePolicy
	}).(RollingUpgradePolicyResponsePtrOutput)
}

type UserAssignedIdentitiesResponseUserAssignedIdentities struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserAssignedIdentitiesResponseUserAssignedIdentitiesOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentitiesResponseUserAssignedIdentitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentitiesResponseUserAssignedIdentities)(nil)).Elem()
}

func (o UserAssignedIdentitiesResponseUserAssignedIdentitiesOutput) ToUserAssignedIdentitiesResponseUserAssignedIdentitiesOutput() UserAssignedIdentitiesResponseUserAssignedIdentitiesOutput {
	return o
}

func (o UserAssignedIdentitiesResponseUserAssignedIdentitiesOutput) ToUserAssignedIdentitiesResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) UserAssignedIdentitiesResponseUserAssignedIdentitiesOutput {
	return o
}

func (o UserAssignedIdentitiesResponseUserAssignedIdentitiesOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentitiesResponseUserAssignedIdentities) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentitiesResponseUserAssignedIdentitiesOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentitiesResponseUserAssignedIdentities) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentitiesResponseUserAssignedIdentitiesMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentitiesResponseUserAssignedIdentitiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentitiesResponseUserAssignedIdentities)(nil)).Elem()
}

func (o UserAssignedIdentitiesResponseUserAssignedIdentitiesMapOutput) ToUserAssignedIdentitiesResponseUserAssignedIdentitiesMapOutput() UserAssignedIdentitiesResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o UserAssignedIdentitiesResponseUserAssignedIdentitiesMapOutput) ToUserAssignedIdentitiesResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) UserAssignedIdentitiesResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o UserAssignedIdentitiesResponseUserAssignedIdentitiesMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentitiesResponseUserAssignedIdentitiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentitiesResponseUserAssignedIdentities {
		return vs[0].(map[string]UserAssignedIdentitiesResponseUserAssignedIdentities)[vs[1].(string)]
	}).(UserAssignedIdentitiesResponseUserAssignedIdentitiesOutput)
}

type VMDiskSecurityProfile struct {
	DiskEncryptionSet      *DiskEncryptionSetParameters `pulumi:"diskEncryptionSet"`
	SecurityEncryptionType *string                      `pulumi:"securityEncryptionType"`
}





type VMDiskSecurityProfileInput interface {
	pulumi.Input

	ToVMDiskSecurityProfileOutput() VMDiskSecurityProfileOutput
	ToVMDiskSecurityProfileOutputWithContext(context.Context) VMDiskSecurityProfileOutput
}

type VMDiskSecurityProfileArgs struct {
	DiskEncryptionSet      DiskEncryptionSetParametersPtrInput `pulumi:"diskEncryptionSet"`
	SecurityEncryptionType pulumi.StringPtrInput               `pulumi:"securityEncryptionType"`
}

func (VMDiskSecurityProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMDiskSecurityProfile)(nil)).Elem()
}

func (i VMDiskSecurityProfileArgs) ToVMDiskSecurityProfileOutput() VMDiskSecurityProfileOutput {
	return i.ToVMDiskSecurityProfileOutputWithContext(context.Background())
}

func (i VMDiskSecurityProfileArgs) ToVMDiskSecurityProfileOutputWithContext(ctx context.Context) VMDiskSecurityProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMDiskSecurityProfileOutput)
}

func (i VMDiskSecurityProfileArgs) ToVMDiskSecurityProfilePtrOutput() VMDiskSecurityProfilePtrOutput {
	return i.ToVMDiskSecurityProfilePtrOutputWithContext(context.Background())
}

func (i VMDiskSecurityProfileArgs) ToVMDiskSecurityProfilePtrOutputWithContext(ctx context.Context) VMDiskSecurityProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMDiskSecurityProfileOutput).ToVMDiskSecurityProfilePtrOutputWithContext(ctx)
}









type VMDiskSecurityProfilePtrInput interface {
	pulumi.Input

	ToVMDiskSecurityProfilePtrOutput() VMDiskSecurityProfilePtrOutput
	ToVMDiskSecurityProfilePtrOutputWithContext(context.Context) VMDiskSecurityProfilePtrOutput
}

type vmdiskSecurityProfilePtrType VMDiskSecurityProfileArgs

func VMDiskSecurityProfilePtr(v *VMDiskSecurityProfileArgs) VMDiskSecurityProfilePtrInput {
	return (*vmdiskSecurityProfilePtrType)(v)
}

func (*vmdiskSecurityProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VMDiskSecurityProfile)(nil)).Elem()
}

func (i *vmdiskSecurityProfilePtrType) ToVMDiskSecurityProfilePtrOutput() VMDiskSecurityProfilePtrOutput {
	return i.ToVMDiskSecurityProfilePtrOutputWithContext(context.Background())
}

func (i *vmdiskSecurityProfilePtrType) ToVMDiskSecurityProfilePtrOutputWithContext(ctx context.Context) VMDiskSecurityProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMDiskSecurityProfilePtrOutput)
}

type VMDiskSecurityProfileOutput struct{ *pulumi.OutputState }

func (VMDiskSecurityProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMDiskSecurityProfile)(nil)).Elem()
}

func (o VMDiskSecurityProfileOutput) ToVMDiskSecurityProfileOutput() VMDiskSecurityProfileOutput {
	return o
}

func (o VMDiskSecurityProfileOutput) ToVMDiskSecurityProfileOutputWithContext(ctx context.Context) VMDiskSecurityProfileOutput {
	return o
}

func (o VMDiskSecurityProfileOutput) ToVMDiskSecurityProfilePtrOutput() VMDiskSecurityProfilePtrOutput {
	return o.ToVMDiskSecurityProfilePtrOutputWithContext(context.Background())
}

func (o VMDiskSecurityProfileOutput) ToVMDiskSecurityProfilePtrOutputWithContext(ctx context.Context) VMDiskSecurityProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VMDiskSecurityProfile) *VMDiskSecurityProfile {
		return &v
	}).(VMDiskSecurityProfilePtrOutput)
}

func (o VMDiskSecurityProfileOutput) DiskEncryptionSet() DiskEncryptionSetParametersPtrOutput {
	return o.ApplyT(func(v VMDiskSecurityProfile) *DiskEncryptionSetParameters { return v.DiskEncryptionSet }).(DiskEncryptionSetParametersPtrOutput)
}

func (o VMDiskSecurityProfileOutput) SecurityEncryptionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMDiskSecurityProfile) *string { return v.SecurityEncryptionType }).(pulumi.StringPtrOutput)
}

type VMDiskSecurityProfilePtrOutput struct{ *pulumi.OutputState }

func (VMDiskSecurityProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMDiskSecurityProfile)(nil)).Elem()
}

func (o VMDiskSecurityProfilePtrOutput) ToVMDiskSecurityProfilePtrOutput() VMDiskSecurityProfilePtrOutput {
	return o
}

func (o VMDiskSecurityProfilePtrOutput) ToVMDiskSecurityProfilePtrOutputWithContext(ctx context.Context) VMDiskSecurityProfilePtrOutput {
	return o
}

func (o VMDiskSecurityProfilePtrOutput) Elem() VMDiskSecurityProfileOutput {
	return o.ApplyT(func(v *VMDiskSecurityProfile) VMDiskSecurityProfile {
		if v != nil {
			return *v
		}
		var ret VMDiskSecurityProfile
		return ret
	}).(VMDiskSecurityProfileOutput)
}

func (o VMDiskSecurityProfilePtrOutput) DiskEncryptionSet() DiskEncryptionSetParametersPtrOutput {
	return o.ApplyT(func(v *VMDiskSecurityProfile) *DiskEncryptionSetParameters {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionSet
	}).(DiskEncryptionSetParametersPtrOutput)
}

func (o VMDiskSecurityProfilePtrOutput) SecurityEncryptionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMDiskSecurityProfile) *string {
		if v == nil {
			return nil
		}
		return v.SecurityEncryptionType
	}).(pulumi.StringPtrOutput)
}

type VMDiskSecurityProfileResponse struct {
	DiskEncryptionSet      *DiskEncryptionSetParametersResponse `pulumi:"diskEncryptionSet"`
	SecurityEncryptionType *string                              `pulumi:"securityEncryptionType"`
}

type VMDiskSecurityProfileResponseOutput struct{ *pulumi.OutputState }

func (VMDiskSecurityProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMDiskSecurityProfileResponse)(nil)).Elem()
}

func (o VMDiskSecurityProfileResponseOutput) ToVMDiskSecurityProfileResponseOutput() VMDiskSecurityProfileResponseOutput {
	return o
}

func (o VMDiskSecurityProfileResponseOutput) ToVMDiskSecurityProfileResponseOutputWithContext(ctx context.Context) VMDiskSecurityProfileResponseOutput {
	return o
}

func (o VMDiskSecurityProfileResponseOutput) DiskEncryptionSet() DiskEncryptionSetParametersResponsePtrOutput {
	return o.ApplyT(func(v VMDiskSecurityProfileResponse) *DiskEncryptionSetParametersResponse { return v.DiskEncryptionSet }).(DiskEncryptionSetParametersResponsePtrOutput)
}

func (o VMDiskSecurityProfileResponseOutput) SecurityEncryptionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMDiskSecurityProfileResponse) *string { return v.SecurityEncryptionType }).(pulumi.StringPtrOutput)
}

type VMDiskSecurityProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VMDiskSecurityProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMDiskSecurityProfileResponse)(nil)).Elem()
}

func (o VMDiskSecurityProfileResponsePtrOutput) ToVMDiskSecurityProfileResponsePtrOutput() VMDiskSecurityProfileResponsePtrOutput {
	return o
}

func (o VMDiskSecurityProfileResponsePtrOutput) ToVMDiskSecurityProfileResponsePtrOutputWithContext(ctx context.Context) VMDiskSecurityProfileResponsePtrOutput {
	return o
}

func (o VMDiskSecurityProfileResponsePtrOutput) Elem() VMDiskSecurityProfileResponseOutput {
	return o.ApplyT(func(v *VMDiskSecurityProfileResponse) VMDiskSecurityProfileResponse {
		if v != nil {
			return *v
		}
		var ret VMDiskSecurityProfileResponse
		return ret
	}).(VMDiskSecurityProfileResponseOutput)
}

func (o VMDiskSecurityProfileResponsePtrOutput) DiskEncryptionSet() DiskEncryptionSetParametersResponsePtrOutput {
	return o.ApplyT(func(v *VMDiskSecurityProfileResponse) *DiskEncryptionSetParametersResponse {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionSet
	}).(DiskEncryptionSetParametersResponsePtrOutput)
}

func (o VMDiskSecurityProfileResponsePtrOutput) SecurityEncryptionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMDiskSecurityProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.SecurityEncryptionType
	}).(pulumi.StringPtrOutput)
}

type VMGalleryApplication struct {
	ConfigurationReference          *string `pulumi:"configurationReference"`
	EnableAutomaticUpgrade          *bool   `pulumi:"enableAutomaticUpgrade"`
	Order                           *int    `pulumi:"order"`
	PackageReferenceId              string  `pulumi:"packageReferenceId"`
	Tags                            *string `pulumi:"tags"`
	TreatFailureAsDeploymentFailure *bool   `pulumi:"treatFailureAsDeploymentFailure"`
}





type VMGalleryApplicationInput interface {
	pulumi.Input

	ToVMGalleryApplicationOutput() VMGalleryApplicationOutput
	ToVMGalleryApplicationOutputWithContext(context.Context) VMGalleryApplicationOutput
}

type VMGalleryApplicationArgs struct {
	ConfigurationReference          pulumi.StringPtrInput `pulumi:"configurationReference"`
	EnableAutomaticUpgrade          pulumi.BoolPtrInput   `pulumi:"enableAutomaticUpgrade"`
	Order                           pulumi.IntPtrInput    `pulumi:"order"`
	PackageReferenceId              pulumi.StringInput    `pulumi:"packageReferenceId"`
	Tags                            pulumi.StringPtrInput `pulumi:"tags"`
	TreatFailureAsDeploymentFailure pulumi.BoolPtrInput   `pulumi:"treatFailureAsDeploymentFailure"`
}

func (VMGalleryApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMGalleryApplication)(nil)).Elem()
}

func (i VMGalleryApplicationArgs) ToVMGalleryApplicationOutput() VMGalleryApplicationOutput {
	return i.ToVMGalleryApplicationOutputWithContext(context.Background())
}

func (i VMGalleryApplicationArgs) ToVMGalleryApplicationOutputWithContext(ctx context.Context) VMGalleryApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMGalleryApplicationOutput)
}





type VMGalleryApplicationArrayInput interface {
	pulumi.Input

	ToVMGalleryApplicationArrayOutput() VMGalleryApplicationArrayOutput
	ToVMGalleryApplicationArrayOutputWithContext(context.Context) VMGalleryApplicationArrayOutput
}

type VMGalleryApplicationArray []VMGalleryApplicationInput

func (VMGalleryApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMGalleryApplication)(nil)).Elem()
}

func (i VMGalleryApplicationArray) ToVMGalleryApplicationArrayOutput() VMGalleryApplicationArrayOutput {
	return i.ToVMGalleryApplicationArrayOutputWithContext(context.Background())
}

func (i VMGalleryApplicationArray) ToVMGalleryApplicationArrayOutputWithContext(ctx context.Context) VMGalleryApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMGalleryApplicationArrayOutput)
}

type VMGalleryApplicationOutput struct{ *pulumi.OutputState }

func (VMGalleryApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMGalleryApplication)(nil)).Elem()
}

func (o VMGalleryApplicationOutput) ToVMGalleryApplicationOutput() VMGalleryApplicationOutput {
	return o
}

func (o VMGalleryApplicationOutput) ToVMGalleryApplicationOutputWithContext(ctx context.Context) VMGalleryApplicationOutput {
	return o
}

func (o VMGalleryApplicationOutput) ConfigurationReference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMGalleryApplication) *string { return v.ConfigurationReference }).(pulumi.StringPtrOutput)
}

func (o VMGalleryApplicationOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VMGalleryApplication) *bool { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
}

func (o VMGalleryApplicationOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VMGalleryApplication) *int { return v.Order }).(pulumi.IntPtrOutput)
}

func (o VMGalleryApplicationOutput) PackageReferenceId() pulumi.StringOutput {
	return o.ApplyT(func(v VMGalleryApplication) string { return v.PackageReferenceId }).(pulumi.StringOutput)
}

func (o VMGalleryApplicationOutput) Tags() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMGalleryApplication) *string { return v.Tags }).(pulumi.StringPtrOutput)
}

func (o VMGalleryApplicationOutput) TreatFailureAsDeploymentFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VMGalleryApplication) *bool { return v.TreatFailureAsDeploymentFailure }).(pulumi.BoolPtrOutput)
}

type VMGalleryApplicationArrayOutput struct{ *pulumi.OutputState }

func (VMGalleryApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMGalleryApplication)(nil)).Elem()
}

func (o VMGalleryApplicationArrayOutput) ToVMGalleryApplicationArrayOutput() VMGalleryApplicationArrayOutput {
	return o
}

func (o VMGalleryApplicationArrayOutput) ToVMGalleryApplicationArrayOutputWithContext(ctx context.Context) VMGalleryApplicationArrayOutput {
	return o
}

func (o VMGalleryApplicationArrayOutput) Index(i pulumi.IntInput) VMGalleryApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VMGalleryApplication {
		return vs[0].([]VMGalleryApplication)[vs[1].(int)]
	}).(VMGalleryApplicationOutput)
}

type VMGalleryApplicationResponse struct {
	ConfigurationReference          *string `pulumi:"configurationReference"`
	EnableAutomaticUpgrade          *bool   `pulumi:"enableAutomaticUpgrade"`
	Order                           *int    `pulumi:"order"`
	PackageReferenceId              string  `pulumi:"packageReferenceId"`
	Tags                            *string `pulumi:"tags"`
	TreatFailureAsDeploymentFailure *bool   `pulumi:"treatFailureAsDeploymentFailure"`
}

type VMGalleryApplicationResponseOutput struct{ *pulumi.OutputState }

func (VMGalleryApplicationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMGalleryApplicationResponse)(nil)).Elem()
}

func (o VMGalleryApplicationResponseOutput) ToVMGalleryApplicationResponseOutput() VMGalleryApplicationResponseOutput {
	return o
}

func (o VMGalleryApplicationResponseOutput) ToVMGalleryApplicationResponseOutputWithContext(ctx context.Context) VMGalleryApplicationResponseOutput {
	return o
}

func (o VMGalleryApplicationResponseOutput) ConfigurationReference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMGalleryApplicationResponse) *string { return v.ConfigurationReference }).(pulumi.StringPtrOutput)
}

func (o VMGalleryApplicationResponseOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VMGalleryApplicationResponse) *bool { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
}

func (o VMGalleryApplicationResponseOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VMGalleryApplicationResponse) *int { return v.Order }).(pulumi.IntPtrOutput)
}

func (o VMGalleryApplicationResponseOutput) PackageReferenceId() pulumi.StringOutput {
	return o.ApplyT(func(v VMGalleryApplicationResponse) string { return v.PackageReferenceId }).(pulumi.StringOutput)
}

func (o VMGalleryApplicationResponseOutput) Tags() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMGalleryApplicationResponse) *string { return v.Tags }).(pulumi.StringPtrOutput)
}

func (o VMGalleryApplicationResponseOutput) TreatFailureAsDeploymentFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VMGalleryApplicationResponse) *bool { return v.TreatFailureAsDeploymentFailure }).(pulumi.BoolPtrOutput)
}

type VMGalleryApplicationResponseArrayOutput struct{ *pulumi.OutputState }

func (VMGalleryApplicationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMGalleryApplicationResponse)(nil)).Elem()
}

func (o VMGalleryApplicationResponseArrayOutput) ToVMGalleryApplicationResponseArrayOutput() VMGalleryApplicationResponseArrayOutput {
	return o
}

func (o VMGalleryApplicationResponseArrayOutput) ToVMGalleryApplicationResponseArrayOutputWithContext(ctx context.Context) VMGalleryApplicationResponseArrayOutput {
	return o
}

func (o VMGalleryApplicationResponseArrayOutput) Index(i pulumi.IntInput) VMGalleryApplicationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VMGalleryApplicationResponse {
		return vs[0].([]VMGalleryApplicationResponse)[vs[1].(int)]
	}).(VMGalleryApplicationResponseOutput)
}

type VMSizeProperties struct {
	VCPUsAvailable *int `pulumi:"vCPUsAvailable"`
	VCPUsPerCore   *int `pulumi:"vCPUsPerCore"`
}





type VMSizePropertiesInput interface {
	pulumi.Input

	ToVMSizePropertiesOutput() VMSizePropertiesOutput
	ToVMSizePropertiesOutputWithContext(context.Context) VMSizePropertiesOutput
}

type VMSizePropertiesArgs struct {
	VCPUsAvailable pulumi.IntPtrInput `pulumi:"vCPUsAvailable"`
	VCPUsPerCore   pulumi.IntPtrInput `pulumi:"vCPUsPerCore"`
}

func (VMSizePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSizeProperties)(nil)).Elem()
}

func (i VMSizePropertiesArgs) ToVMSizePropertiesOutput() VMSizePropertiesOutput {
	return i.ToVMSizePropertiesOutputWithContext(context.Background())
}

func (i VMSizePropertiesArgs) ToVMSizePropertiesOutputWithContext(ctx context.Context) VMSizePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSizePropertiesOutput)
}

func (i VMSizePropertiesArgs) ToVMSizePropertiesPtrOutput() VMSizePropertiesPtrOutput {
	return i.ToVMSizePropertiesPtrOutputWithContext(context.Background())
}

func (i VMSizePropertiesArgs) ToVMSizePropertiesPtrOutputWithContext(ctx context.Context) VMSizePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSizePropertiesOutput).ToVMSizePropertiesPtrOutputWithContext(ctx)
}









type VMSizePropertiesPtrInput interface {
	pulumi.Input

	ToVMSizePropertiesPtrOutput() VMSizePropertiesPtrOutput
	ToVMSizePropertiesPtrOutputWithContext(context.Context) VMSizePropertiesPtrOutput
}

type vmsizePropertiesPtrType VMSizePropertiesArgs

func VMSizePropertiesPtr(v *VMSizePropertiesArgs) VMSizePropertiesPtrInput {
	return (*vmsizePropertiesPtrType)(v)
}

func (*vmsizePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VMSizeProperties)(nil)).Elem()
}

func (i *vmsizePropertiesPtrType) ToVMSizePropertiesPtrOutput() VMSizePropertiesPtrOutput {
	return i.ToVMSizePropertiesPtrOutputWithContext(context.Background())
}

func (i *vmsizePropertiesPtrType) ToVMSizePropertiesPtrOutputWithContext(ctx context.Context) VMSizePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSizePropertiesPtrOutput)
}

type VMSizePropertiesOutput struct{ *pulumi.OutputState }

func (VMSizePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSizeProperties)(nil)).Elem()
}

func (o VMSizePropertiesOutput) ToVMSizePropertiesOutput() VMSizePropertiesOutput {
	return o
}

func (o VMSizePropertiesOutput) ToVMSizePropertiesOutputWithContext(ctx context.Context) VMSizePropertiesOutput {
	return o
}

func (o VMSizePropertiesOutput) ToVMSizePropertiesPtrOutput() VMSizePropertiesPtrOutput {
	return o.ToVMSizePropertiesPtrOutputWithContext(context.Background())
}

func (o VMSizePropertiesOutput) ToVMSizePropertiesPtrOutputWithContext(ctx context.Context) VMSizePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VMSizeProperties) *VMSizeProperties {
		return &v
	}).(VMSizePropertiesPtrOutput)
}

func (o VMSizePropertiesOutput) VCPUsAvailable() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VMSizeProperties) *int { return v.VCPUsAvailable }).(pulumi.IntPtrOutput)
}

func (o VMSizePropertiesOutput) VCPUsPerCore() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VMSizeProperties) *int { return v.VCPUsPerCore }).(pulumi.IntPtrOutput)
}

type VMSizePropertiesPtrOutput struct{ *pulumi.OutputState }

func (VMSizePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMSizeProperties)(nil)).Elem()
}

func (o VMSizePropertiesPtrOutput) ToVMSizePropertiesPtrOutput() VMSizePropertiesPtrOutput {
	return o
}

func (o VMSizePropertiesPtrOutput) ToVMSizePropertiesPtrOutputWithContext(ctx context.Context) VMSizePropertiesPtrOutput {
	return o
}

func (o VMSizePropertiesPtrOutput) Elem() VMSizePropertiesOutput {
	return o.ApplyT(func(v *VMSizeProperties) VMSizeProperties {
		if v != nil {
			return *v
		}
		var ret VMSizeProperties
		return ret
	}).(VMSizePropertiesOutput)
}

func (o VMSizePropertiesPtrOutput) VCPUsAvailable() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VMSizeProperties) *int {
		if v == nil {
			return nil
		}
		return v.VCPUsAvailable
	}).(pulumi.IntPtrOutput)
}

func (o VMSizePropertiesPtrOutput) VCPUsPerCore() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VMSizeProperties) *int {
		if v == nil {
			return nil
		}
		return v.VCPUsPerCore
	}).(pulumi.IntPtrOutput)
}

type VMSizePropertiesResponse struct {
	VCPUsAvailable *int `pulumi:"vCPUsAvailable"`
	VCPUsPerCore   *int `pulumi:"vCPUsPerCore"`
}

type VMSizePropertiesResponseOutput struct{ *pulumi.OutputState }

func (VMSizePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSizePropertiesResponse)(nil)).Elem()
}

func (o VMSizePropertiesResponseOutput) ToVMSizePropertiesResponseOutput() VMSizePropertiesResponseOutput {
	return o
}

func (o VMSizePropertiesResponseOutput) ToVMSizePropertiesResponseOutputWithContext(ctx context.Context) VMSizePropertiesResponseOutput {
	return o
}

func (o VMSizePropertiesResponseOutput) VCPUsAvailable() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VMSizePropertiesResponse) *int { return v.VCPUsAvailable }).(pulumi.IntPtrOutput)
}

func (o VMSizePropertiesResponseOutput) VCPUsPerCore() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VMSizePropertiesResponse) *int { return v.VCPUsPerCore }).(pulumi.IntPtrOutput)
}

type VMSizePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VMSizePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMSizePropertiesResponse)(nil)).Elem()
}

func (o VMSizePropertiesResponsePtrOutput) ToVMSizePropertiesResponsePtrOutput() VMSizePropertiesResponsePtrOutput {
	return o
}

func (o VMSizePropertiesResponsePtrOutput) ToVMSizePropertiesResponsePtrOutputWithContext(ctx context.Context) VMSizePropertiesResponsePtrOutput {
	return o
}

func (o VMSizePropertiesResponsePtrOutput) Elem() VMSizePropertiesResponseOutput {
	return o.ApplyT(func(v *VMSizePropertiesResponse) VMSizePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret VMSizePropertiesResponse
		return ret
	}).(VMSizePropertiesResponseOutput)
}

func (o VMSizePropertiesResponsePtrOutput) VCPUsAvailable() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VMSizePropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.VCPUsAvailable
	}).(pulumi.IntPtrOutput)
}

func (o VMSizePropertiesResponsePtrOutput) VCPUsPerCore() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VMSizePropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.VCPUsPerCore
	}).(pulumi.IntPtrOutput)
}

type VaultCertificate struct {
	CertificateStore *string `pulumi:"certificateStore"`
	CertificateUrl   *string `pulumi:"certificateUrl"`
}





type VaultCertificateInput interface {
	pulumi.Input

	ToVaultCertificateOutput() VaultCertificateOutput
	ToVaultCertificateOutputWithContext(context.Context) VaultCertificateOutput
}

type VaultCertificateArgs struct {
	CertificateStore pulumi.StringPtrInput `pulumi:"certificateStore"`
	CertificateUrl   pulumi.StringPtrInput `pulumi:"certificateUrl"`
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

func (o VaultCertificateOutput) CertificateStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultCertificate) *string { return v.CertificateStore }).(pulumi.StringPtrOutput)
}

func (o VaultCertificateOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultCertificate) *string { return v.CertificateUrl }).(pulumi.StringPtrOutput)
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
	CertificateStore *string `pulumi:"certificateStore"`
	CertificateUrl   *string `pulumi:"certificateUrl"`
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

func (o VaultCertificateResponseOutput) CertificateStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultCertificateResponse) *string { return v.CertificateStore }).(pulumi.StringPtrOutput)
}

func (o VaultCertificateResponseOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultCertificateResponse) *string { return v.CertificateUrl }).(pulumi.StringPtrOutput)
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
	SourceVault       *SubResource       `pulumi:"sourceVault"`
	VaultCertificates []VaultCertificate `pulumi:"vaultCertificates"`
}





type VaultSecretGroupInput interface {
	pulumi.Input

	ToVaultSecretGroupOutput() VaultSecretGroupOutput
	ToVaultSecretGroupOutputWithContext(context.Context) VaultSecretGroupOutput
}

type VaultSecretGroupArgs struct {
	SourceVault       SubResourcePtrInput        `pulumi:"sourceVault"`
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

func (o VaultSecretGroupOutput) SourceVault() SubResourcePtrOutput {
	return o.ApplyT(func(v VaultSecretGroup) *SubResource { return v.SourceVault }).(SubResourcePtrOutput)
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
	SourceVault       *SubResourceResponse       `pulumi:"sourceVault"`
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

func (o VaultSecretGroupResponseOutput) SourceVault() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VaultSecretGroupResponse) *SubResourceResponse { return v.SourceVault }).(SubResourceResponsePtrOutput)
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

type VirtualHardDisk struct {
	Uri *string `pulumi:"uri"`
}





type VirtualHardDiskInput interface {
	pulumi.Input

	ToVirtualHardDiskOutput() VirtualHardDiskOutput
	ToVirtualHardDiskOutputWithContext(context.Context) VirtualHardDiskOutput
}

type VirtualHardDiskArgs struct {
	Uri pulumi.StringPtrInput `pulumi:"uri"`
}

func (VirtualHardDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHardDisk)(nil)).Elem()
}

func (i VirtualHardDiskArgs) ToVirtualHardDiskOutput() VirtualHardDiskOutput {
	return i.ToVirtualHardDiskOutputWithContext(context.Background())
}

func (i VirtualHardDiskArgs) ToVirtualHardDiskOutputWithContext(ctx context.Context) VirtualHardDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHardDiskOutput)
}

func (i VirtualHardDiskArgs) ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput {
	return i.ToVirtualHardDiskPtrOutputWithContext(context.Background())
}

func (i VirtualHardDiskArgs) ToVirtualHardDiskPtrOutputWithContext(ctx context.Context) VirtualHardDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHardDiskOutput).ToVirtualHardDiskPtrOutputWithContext(ctx)
}









type VirtualHardDiskPtrInput interface {
	pulumi.Input

	ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput
	ToVirtualHardDiskPtrOutputWithContext(context.Context) VirtualHardDiskPtrOutput
}

type virtualHardDiskPtrType VirtualHardDiskArgs

func VirtualHardDiskPtr(v *VirtualHardDiskArgs) VirtualHardDiskPtrInput {
	return (*virtualHardDiskPtrType)(v)
}

func (*virtualHardDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHardDisk)(nil)).Elem()
}

func (i *virtualHardDiskPtrType) ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput {
	return i.ToVirtualHardDiskPtrOutputWithContext(context.Background())
}

func (i *virtualHardDiskPtrType) ToVirtualHardDiskPtrOutputWithContext(ctx context.Context) VirtualHardDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHardDiskPtrOutput)
}

type VirtualHardDiskOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHardDisk)(nil)).Elem()
}

func (o VirtualHardDiskOutput) ToVirtualHardDiskOutput() VirtualHardDiskOutput {
	return o
}

func (o VirtualHardDiskOutput) ToVirtualHardDiskOutputWithContext(ctx context.Context) VirtualHardDiskOutput {
	return o
}

func (o VirtualHardDiskOutput) ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput {
	return o.ToVirtualHardDiskPtrOutputWithContext(context.Background())
}

func (o VirtualHardDiskOutput) ToVirtualHardDiskPtrOutputWithContext(ctx context.Context) VirtualHardDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualHardDisk) *VirtualHardDisk {
		return &v
	}).(VirtualHardDiskPtrOutput)
}

func (o VirtualHardDiskOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHardDisk) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type VirtualHardDiskPtrOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHardDisk)(nil)).Elem()
}

func (o VirtualHardDiskPtrOutput) ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput {
	return o
}

func (o VirtualHardDiskPtrOutput) ToVirtualHardDiskPtrOutputWithContext(ctx context.Context) VirtualHardDiskPtrOutput {
	return o
}

func (o VirtualHardDiskPtrOutput) Elem() VirtualHardDiskOutput {
	return o.ApplyT(func(v *VirtualHardDisk) VirtualHardDisk {
		if v != nil {
			return *v
		}
		var ret VirtualHardDisk
		return ret
	}).(VirtualHardDiskOutput)
}

func (o VirtualHardDiskPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHardDisk) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type VirtualHardDiskResponse struct {
	Uri *string `pulumi:"uri"`
}

type VirtualHardDiskResponseOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHardDiskResponse)(nil)).Elem()
}

func (o VirtualHardDiskResponseOutput) ToVirtualHardDiskResponseOutput() VirtualHardDiskResponseOutput {
	return o
}

func (o VirtualHardDiskResponseOutput) ToVirtualHardDiskResponseOutputWithContext(ctx context.Context) VirtualHardDiskResponseOutput {
	return o
}

func (o VirtualHardDiskResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHardDiskResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type VirtualHardDiskResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHardDiskResponse)(nil)).Elem()
}

func (o VirtualHardDiskResponsePtrOutput) ToVirtualHardDiskResponsePtrOutput() VirtualHardDiskResponsePtrOutput {
	return o
}

func (o VirtualHardDiskResponsePtrOutput) ToVirtualHardDiskResponsePtrOutputWithContext(ctx context.Context) VirtualHardDiskResponsePtrOutput {
	return o
}

func (o VirtualHardDiskResponsePtrOutput) Elem() VirtualHardDiskResponseOutput {
	return o.ApplyT(func(v *VirtualHardDiskResponse) VirtualHardDiskResponse {
		if v != nil {
			return *v
		}
		var ret VirtualHardDiskResponse
		return ret
	}).(VirtualHardDiskResponseOutput)
}

func (o VirtualHardDiskResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHardDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineAgentInstanceViewResponse struct {
	ExtensionHandlers []VirtualMachineExtensionHandlerInstanceViewResponse `pulumi:"extensionHandlers"`
	Statuses          []InstanceViewStatusResponse                         `pulumi:"statuses"`
	VmAgentVersion    *string                                              `pulumi:"vmAgentVersion"`
}

type VirtualMachineAgentInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineAgentInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineAgentInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineAgentInstanceViewResponseOutput) ToVirtualMachineAgentInstanceViewResponseOutput() VirtualMachineAgentInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineAgentInstanceViewResponseOutput) ToVirtualMachineAgentInstanceViewResponseOutputWithContext(ctx context.Context) VirtualMachineAgentInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineAgentInstanceViewResponseOutput) ExtensionHandlers() VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineAgentInstanceViewResponse) []VirtualMachineExtensionHandlerInstanceViewResponse {
		return v.ExtensionHandlers
	}).(VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput)
}

func (o VirtualMachineAgentInstanceViewResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineAgentInstanceViewResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineAgentInstanceViewResponseOutput) VmAgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineAgentInstanceViewResponse) *string { return v.VmAgentVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineAgentInstanceViewResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineAgentInstanceViewResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineAgentInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineAgentInstanceViewResponsePtrOutput) ToVirtualMachineAgentInstanceViewResponsePtrOutput() VirtualMachineAgentInstanceViewResponsePtrOutput {
	return o
}

func (o VirtualMachineAgentInstanceViewResponsePtrOutput) ToVirtualMachineAgentInstanceViewResponsePtrOutputWithContext(ctx context.Context) VirtualMachineAgentInstanceViewResponsePtrOutput {
	return o
}

func (o VirtualMachineAgentInstanceViewResponsePtrOutput) Elem() VirtualMachineAgentInstanceViewResponseOutput {
	return o.ApplyT(func(v *VirtualMachineAgentInstanceViewResponse) VirtualMachineAgentInstanceViewResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineAgentInstanceViewResponse
		return ret
	}).(VirtualMachineAgentInstanceViewResponseOutput)
}

func (o VirtualMachineAgentInstanceViewResponsePtrOutput) ExtensionHandlers() VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineAgentInstanceViewResponse) []VirtualMachineExtensionHandlerInstanceViewResponse {
		if v == nil {
			return nil
		}
		return v.ExtensionHandlers
	}).(VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput)
}

func (o VirtualMachineAgentInstanceViewResponsePtrOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineAgentInstanceViewResponse) []InstanceViewStatusResponse {
		if v == nil {
			return nil
		}
		return v.Statuses
	}).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineAgentInstanceViewResponsePtrOutput) VmAgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineAgentInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmAgentVersion
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionHandlerInstanceViewResponse struct {
	Status             *InstanceViewStatusResponse `pulumi:"status"`
	Type               *string                     `pulumi:"type"`
	TypeHandlerVersion *string                     `pulumi:"typeHandlerVersion"`
}

type VirtualMachineExtensionHandlerInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionHandlerInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineExtensionHandlerInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseOutput) ToVirtualMachineExtensionHandlerInstanceViewResponseOutput() VirtualMachineExtensionHandlerInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseOutput) ToVirtualMachineExtensionHandlerInstanceViewResponseOutputWithContext(ctx context.Context) VirtualMachineExtensionHandlerInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseOutput) Status() InstanceViewStatusResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionHandlerInstanceViewResponse) *InstanceViewStatusResponse {
		return v.Status
	}).(InstanceViewStatusResponsePtrOutput)
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionHandlerInstanceViewResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionHandlerInstanceViewResponse) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineExtensionHandlerInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput) ToVirtualMachineExtensionHandlerInstanceViewResponseArrayOutput() VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput {
	return o
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput) ToVirtualMachineExtensionHandlerInstanceViewResponseArrayOutputWithContext(ctx context.Context) VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput {
	return o
}

func (o VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineExtensionHandlerInstanceViewResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineExtensionHandlerInstanceViewResponse {
		return vs[0].([]VirtualMachineExtensionHandlerInstanceViewResponse)[vs[1].(int)]
	}).(VirtualMachineExtensionHandlerInstanceViewResponseOutput)
}

type VirtualMachineExtensionInstanceView struct {
	Name               *string              `pulumi:"name"`
	Statuses           []InstanceViewStatus `pulumi:"statuses"`
	Substatuses        []InstanceViewStatus `pulumi:"substatuses"`
	Type               *string              `pulumi:"type"`
	TypeHandlerVersion *string              `pulumi:"typeHandlerVersion"`
}





type VirtualMachineExtensionInstanceViewInput interface {
	pulumi.Input

	ToVirtualMachineExtensionInstanceViewOutput() VirtualMachineExtensionInstanceViewOutput
	ToVirtualMachineExtensionInstanceViewOutputWithContext(context.Context) VirtualMachineExtensionInstanceViewOutput
}

type VirtualMachineExtensionInstanceViewArgs struct {
	Name               pulumi.StringPtrInput        `pulumi:"name"`
	Statuses           InstanceViewStatusArrayInput `pulumi:"statuses"`
	Substatuses        InstanceViewStatusArrayInput `pulumi:"substatuses"`
	Type               pulumi.StringPtrInput        `pulumi:"type"`
	TypeHandlerVersion pulumi.StringPtrInput        `pulumi:"typeHandlerVersion"`
}

func (VirtualMachineExtensionInstanceViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineExtensionInstanceView)(nil)).Elem()
}

func (i VirtualMachineExtensionInstanceViewArgs) ToVirtualMachineExtensionInstanceViewOutput() VirtualMachineExtensionInstanceViewOutput {
	return i.ToVirtualMachineExtensionInstanceViewOutputWithContext(context.Background())
}

func (i VirtualMachineExtensionInstanceViewArgs) ToVirtualMachineExtensionInstanceViewOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineExtensionInstanceViewOutput)
}

func (i VirtualMachineExtensionInstanceViewArgs) ToVirtualMachineExtensionInstanceViewPtrOutput() VirtualMachineExtensionInstanceViewPtrOutput {
	return i.ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(context.Background())
}

func (i VirtualMachineExtensionInstanceViewArgs) ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineExtensionInstanceViewOutput).ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(ctx)
}









type VirtualMachineExtensionInstanceViewPtrInput interface {
	pulumi.Input

	ToVirtualMachineExtensionInstanceViewPtrOutput() VirtualMachineExtensionInstanceViewPtrOutput
	ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(context.Context) VirtualMachineExtensionInstanceViewPtrOutput
}

type virtualMachineExtensionInstanceViewPtrType VirtualMachineExtensionInstanceViewArgs

func VirtualMachineExtensionInstanceViewPtr(v *VirtualMachineExtensionInstanceViewArgs) VirtualMachineExtensionInstanceViewPtrInput {
	return (*virtualMachineExtensionInstanceViewPtrType)(v)
}

func (*virtualMachineExtensionInstanceViewPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineExtensionInstanceView)(nil)).Elem()
}

func (i *virtualMachineExtensionInstanceViewPtrType) ToVirtualMachineExtensionInstanceViewPtrOutput() VirtualMachineExtensionInstanceViewPtrOutput {
	return i.ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(context.Background())
}

func (i *virtualMachineExtensionInstanceViewPtrType) ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineExtensionInstanceViewPtrOutput)
}

type VirtualMachineExtensionInstanceViewOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionInstanceViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineExtensionInstanceView)(nil)).Elem()
}

func (o VirtualMachineExtensionInstanceViewOutput) ToVirtualMachineExtensionInstanceViewOutput() VirtualMachineExtensionInstanceViewOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewOutput) ToVirtualMachineExtensionInstanceViewOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewOutput) ToVirtualMachineExtensionInstanceViewPtrOutput() VirtualMachineExtensionInstanceViewPtrOutput {
	return o.ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(context.Background())
}

func (o VirtualMachineExtensionInstanceViewOutput) ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineExtensionInstanceView) *VirtualMachineExtensionInstanceView {
		return &v
	}).(VirtualMachineExtensionInstanceViewPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceView) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewOutput) Statuses() InstanceViewStatusArrayOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceView) []InstanceViewStatus { return v.Statuses }).(InstanceViewStatusArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewOutput) Substatuses() InstanceViewStatusArrayOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceView) []InstanceViewStatus { return v.Substatuses }).(InstanceViewStatusArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceView) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceView) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionInstanceViewPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionInstanceViewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineExtensionInstanceView)(nil)).Elem()
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) ToVirtualMachineExtensionInstanceViewPtrOutput() VirtualMachineExtensionInstanceViewPtrOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) ToVirtualMachineExtensionInstanceViewPtrOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewPtrOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) Elem() VirtualMachineExtensionInstanceViewOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceView) VirtualMachineExtensionInstanceView {
		if v != nil {
			return *v
		}
		var ret VirtualMachineExtensionInstanceView
		return ret
	}).(VirtualMachineExtensionInstanceViewOutput)
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceView) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) Statuses() InstanceViewStatusArrayOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceView) []InstanceViewStatus {
		if v == nil {
			return nil
		}
		return v.Statuses
	}).(InstanceViewStatusArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) Substatuses() InstanceViewStatusArrayOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceView) []InstanceViewStatus {
		if v == nil {
			return nil
		}
		return v.Substatuses
	}).(InstanceViewStatusArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceView) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewPtrOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceView) *string {
		if v == nil {
			return nil
		}
		return v.TypeHandlerVersion
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionInstanceViewResponse struct {
	Name               *string                      `pulumi:"name"`
	Statuses           []InstanceViewStatusResponse `pulumi:"statuses"`
	Substatuses        []InstanceViewStatusResponse `pulumi:"substatuses"`
	Type               *string                      `pulumi:"type"`
	TypeHandlerVersion *string                      `pulumi:"typeHandlerVersion"`
}

type VirtualMachineExtensionInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) ToVirtualMachineExtensionInstanceViewResponseOutput() VirtualMachineExtensionInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) ToVirtualMachineExtensionInstanceViewResponseOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceViewResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceViewResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) Substatuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceViewResponse) []InstanceViewStatusResponse { return v.Substatuses }).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceViewResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewResponseOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionInstanceViewResponse) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionInstanceViewResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionInstanceViewResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) ToVirtualMachineExtensionInstanceViewResponsePtrOutput() VirtualMachineExtensionInstanceViewResponsePtrOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) ToVirtualMachineExtensionInstanceViewResponsePtrOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewResponsePtrOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) Elem() VirtualMachineExtensionInstanceViewResponseOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceViewResponse) VirtualMachineExtensionInstanceViewResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineExtensionInstanceViewResponse
		return ret
	}).(VirtualMachineExtensionInstanceViewResponseOutput)
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceViewResponse) []InstanceViewStatusResponse {
		if v == nil {
			return nil
		}
		return v.Statuses
	}).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) Substatuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceViewResponse) []InstanceViewStatusResponse {
		if v == nil {
			return nil
		}
		return v.Substatuses
	}).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionInstanceViewResponsePtrOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtensionInstanceViewResponse) *string {
		if v == nil {
			return nil
		}
		return v.TypeHandlerVersion
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionInstanceViewResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionInstanceViewResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineExtensionInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionInstanceViewResponseArrayOutput) ToVirtualMachineExtensionInstanceViewResponseArrayOutput() VirtualMachineExtensionInstanceViewResponseArrayOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewResponseArrayOutput) ToVirtualMachineExtensionInstanceViewResponseArrayOutputWithContext(ctx context.Context) VirtualMachineExtensionInstanceViewResponseArrayOutput {
	return o
}

func (o VirtualMachineExtensionInstanceViewResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineExtensionInstanceViewResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineExtensionInstanceViewResponse {
		return vs[0].([]VirtualMachineExtensionInstanceViewResponse)[vs[1].(int)]
	}).(VirtualMachineExtensionInstanceViewResponseOutput)
}

type VirtualMachineExtensionResponse struct {
	AutoUpgradeMinorVersion       *bool                                        `pulumi:"autoUpgradeMinorVersion"`
	EnableAutomaticUpgrade        *bool                                        `pulumi:"enableAutomaticUpgrade"`
	ForceUpdateTag                *string                                      `pulumi:"forceUpdateTag"`
	Id                            string                                       `pulumi:"id"`
	InstanceView                  *VirtualMachineExtensionInstanceViewResponse `pulumi:"instanceView"`
	Location                      *string                                      `pulumi:"location"`
	Name                          string                                       `pulumi:"name"`
	ProtectedSettings             interface{}                                  `pulumi:"protectedSettings"`
	ProtectedSettingsFromKeyVault *KeyVaultSecretReferenceResponse             `pulumi:"protectedSettingsFromKeyVault"`
	ProvisioningState             string                                       `pulumi:"provisioningState"`
	Publisher                     *string                                      `pulumi:"publisher"`
	Settings                      interface{}                                  `pulumi:"settings"`
	SuppressFailures              *bool                                        `pulumi:"suppressFailures"`
	Tags                          map[string]string                            `pulumi:"tags"`
	Type                          string                                       `pulumi:"type"`
	TypeHandlerVersion            *string                                      `pulumi:"typeHandlerVersion"`
}

type VirtualMachineExtensionResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineExtensionResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionResponseOutput) ToVirtualMachineExtensionResponseOutput() VirtualMachineExtensionResponseOutput {
	return o
}

func (o VirtualMachineExtensionResponseOutput) ToVirtualMachineExtensionResponseOutputWithContext(ctx context.Context) VirtualMachineExtensionResponseOutput {
	return o
}

func (o VirtualMachineExtensionResponseOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineExtensionResponseOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *bool { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineExtensionResponseOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualMachineExtensionResponseOutput) InstanceView() VirtualMachineExtensionInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *VirtualMachineExtensionInstanceViewResponse {
		return v.InstanceView
	}).(VirtualMachineExtensionInstanceViewResponsePtrOutput)
}

func (o VirtualMachineExtensionResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineExtensionResponseOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o VirtualMachineExtensionResponseOutput) ProtectedSettingsFromKeyVault() KeyVaultSecretReferenceResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *KeyVaultSecretReferenceResponse {
		return v.ProtectedSettingsFromKeyVault
	}).(KeyVaultSecretReferenceResponsePtrOutput)
}

func (o VirtualMachineExtensionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineExtensionResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionResponseOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o VirtualMachineExtensionResponseOutput) SuppressFailures() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *bool { return v.SuppressFailures }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineExtensionResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualMachineExtensionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualMachineExtensionResponseOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineExtensionResponse) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineExtensionResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineExtensionResponse)(nil)).Elem()
}

func (o VirtualMachineExtensionResponseArrayOutput) ToVirtualMachineExtensionResponseArrayOutput() VirtualMachineExtensionResponseArrayOutput {
	return o
}

func (o VirtualMachineExtensionResponseArrayOutput) ToVirtualMachineExtensionResponseArrayOutputWithContext(ctx context.Context) VirtualMachineExtensionResponseArrayOutput {
	return o
}

func (o VirtualMachineExtensionResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineExtensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineExtensionResponse {
		return vs[0].([]VirtualMachineExtensionResponse)[vs[1].(int)]
	}).(VirtualMachineExtensionResponseOutput)
}

type VirtualMachineHealthStatusResponse struct {
	Status InstanceViewStatusResponse `pulumi:"status"`
}

type VirtualMachineHealthStatusResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineHealthStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineHealthStatusResponse)(nil)).Elem()
}

func (o VirtualMachineHealthStatusResponseOutput) ToVirtualMachineHealthStatusResponseOutput() VirtualMachineHealthStatusResponseOutput {
	return o
}

func (o VirtualMachineHealthStatusResponseOutput) ToVirtualMachineHealthStatusResponseOutputWithContext(ctx context.Context) VirtualMachineHealthStatusResponseOutput {
	return o
}

func (o VirtualMachineHealthStatusResponseOutput) Status() InstanceViewStatusResponseOutput {
	return o.ApplyT(func(v VirtualMachineHealthStatusResponse) InstanceViewStatusResponse { return v.Status }).(InstanceViewStatusResponseOutput)
}

type VirtualMachineIdentity struct {
	Type                   *ResourceIdentityType  `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type VirtualMachineIdentityInput interface {
	pulumi.Input

	ToVirtualMachineIdentityOutput() VirtualMachineIdentityOutput
	ToVirtualMachineIdentityOutputWithContext(context.Context) VirtualMachineIdentityOutput
}

type VirtualMachineIdentityArgs struct {
	Type                   ResourceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput              `pulumi:"userAssignedIdentities"`
}

func (VirtualMachineIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineIdentity)(nil)).Elem()
}

func (i VirtualMachineIdentityArgs) ToVirtualMachineIdentityOutput() VirtualMachineIdentityOutput {
	return i.ToVirtualMachineIdentityOutputWithContext(context.Background())
}

func (i VirtualMachineIdentityArgs) ToVirtualMachineIdentityOutputWithContext(ctx context.Context) VirtualMachineIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineIdentityOutput)
}

func (i VirtualMachineIdentityArgs) ToVirtualMachineIdentityPtrOutput() VirtualMachineIdentityPtrOutput {
	return i.ToVirtualMachineIdentityPtrOutputWithContext(context.Background())
}

func (i VirtualMachineIdentityArgs) ToVirtualMachineIdentityPtrOutputWithContext(ctx context.Context) VirtualMachineIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineIdentityOutput).ToVirtualMachineIdentityPtrOutputWithContext(ctx)
}









type VirtualMachineIdentityPtrInput interface {
	pulumi.Input

	ToVirtualMachineIdentityPtrOutput() VirtualMachineIdentityPtrOutput
	ToVirtualMachineIdentityPtrOutputWithContext(context.Context) VirtualMachineIdentityPtrOutput
}

type virtualMachineIdentityPtrType VirtualMachineIdentityArgs

func VirtualMachineIdentityPtr(v *VirtualMachineIdentityArgs) VirtualMachineIdentityPtrInput {
	return (*virtualMachineIdentityPtrType)(v)
}

func (*virtualMachineIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineIdentity)(nil)).Elem()
}

func (i *virtualMachineIdentityPtrType) ToVirtualMachineIdentityPtrOutput() VirtualMachineIdentityPtrOutput {
	return i.ToVirtualMachineIdentityPtrOutputWithContext(context.Background())
}

func (i *virtualMachineIdentityPtrType) ToVirtualMachineIdentityPtrOutputWithContext(ctx context.Context) VirtualMachineIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineIdentityPtrOutput)
}

type VirtualMachineIdentityOutput struct{ *pulumi.OutputState }

func (VirtualMachineIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineIdentity)(nil)).Elem()
}

func (o VirtualMachineIdentityOutput) ToVirtualMachineIdentityOutput() VirtualMachineIdentityOutput {
	return o
}

func (o VirtualMachineIdentityOutput) ToVirtualMachineIdentityOutputWithContext(ctx context.Context) VirtualMachineIdentityOutput {
	return o
}

func (o VirtualMachineIdentityOutput) ToVirtualMachineIdentityPtrOutput() VirtualMachineIdentityPtrOutput {
	return o.ToVirtualMachineIdentityPtrOutputWithContext(context.Background())
}

func (o VirtualMachineIdentityOutput) ToVirtualMachineIdentityPtrOutputWithContext(ctx context.Context) VirtualMachineIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineIdentity) *VirtualMachineIdentity {
		return &v
	}).(VirtualMachineIdentityPtrOutput)
}

func (o VirtualMachineIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v VirtualMachineIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

func (o VirtualMachineIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v VirtualMachineIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type VirtualMachineIdentityPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineIdentity)(nil)).Elem()
}

func (o VirtualMachineIdentityPtrOutput) ToVirtualMachineIdentityPtrOutput() VirtualMachineIdentityPtrOutput {
	return o
}

func (o VirtualMachineIdentityPtrOutput) ToVirtualMachineIdentityPtrOutputWithContext(ctx context.Context) VirtualMachineIdentityPtrOutput {
	return o
}

func (o VirtualMachineIdentityPtrOutput) Elem() VirtualMachineIdentityOutput {
	return o.ApplyT(func(v *VirtualMachineIdentity) VirtualMachineIdentity {
		if v != nil {
			return *v
		}
		var ret VirtualMachineIdentity
		return ret
	}).(VirtualMachineIdentityOutput)
}

func (o VirtualMachineIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *VirtualMachineIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

func (o VirtualMachineIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *VirtualMachineIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type VirtualMachineIdentityResponse struct {
	PrincipalId            string                                                          `pulumi:"principalId"`
	TenantId               string                                                          `pulumi:"tenantId"`
	Type                   *string                                                         `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentitiesResponseUserAssignedIdentities `pulumi:"userAssignedIdentities"`
}

type VirtualMachineIdentityResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineIdentityResponse)(nil)).Elem()
}

func (o VirtualMachineIdentityResponseOutput) ToVirtualMachineIdentityResponseOutput() VirtualMachineIdentityResponseOutput {
	return o
}

func (o VirtualMachineIdentityResponseOutput) ToVirtualMachineIdentityResponseOutputWithContext(ctx context.Context) VirtualMachineIdentityResponseOutput {
	return o
}

func (o VirtualMachineIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o VirtualMachineIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o VirtualMachineIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentitiesResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v VirtualMachineIdentityResponse) map[string]UserAssignedIdentitiesResponseUserAssignedIdentities {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentitiesResponseUserAssignedIdentitiesMapOutput)
}

type VirtualMachineIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineIdentityResponse)(nil)).Elem()
}

func (o VirtualMachineIdentityResponsePtrOutput) ToVirtualMachineIdentityResponsePtrOutput() VirtualMachineIdentityResponsePtrOutput {
	return o
}

func (o VirtualMachineIdentityResponsePtrOutput) ToVirtualMachineIdentityResponsePtrOutputWithContext(ctx context.Context) VirtualMachineIdentityResponsePtrOutput {
	return o
}

func (o VirtualMachineIdentityResponsePtrOutput) Elem() VirtualMachineIdentityResponseOutput {
	return o.ApplyT(func(v *VirtualMachineIdentityResponse) VirtualMachineIdentityResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineIdentityResponse
		return ret
	}).(VirtualMachineIdentityResponseOutput)
}

func (o VirtualMachineIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentitiesResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v *VirtualMachineIdentityResponse) map[string]UserAssignedIdentitiesResponseUserAssignedIdentities {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentitiesResponseUserAssignedIdentitiesMapOutput)
}

type VirtualMachineInstanceViewResponse struct {
	AssignedHost              string                                        `pulumi:"assignedHost"`
	BootDiagnostics           *BootDiagnosticsInstanceViewResponse          `pulumi:"bootDiagnostics"`
	ComputerName              *string                                       `pulumi:"computerName"`
	Disks                     []DiskInstanceViewResponse                    `pulumi:"disks"`
	Extensions                []VirtualMachineExtensionInstanceViewResponse `pulumi:"extensions"`
	HyperVGeneration          *string                                       `pulumi:"hyperVGeneration"`
	MaintenanceRedeployStatus *MaintenanceRedeployStatusResponse            `pulumi:"maintenanceRedeployStatus"`
	OsName                    *string                                       `pulumi:"osName"`
	OsVersion                 *string                                       `pulumi:"osVersion"`
	PatchStatus               *VirtualMachinePatchStatusResponse            `pulumi:"patchStatus"`
	PlatformFaultDomain       *int                                          `pulumi:"platformFaultDomain"`
	PlatformUpdateDomain      *int                                          `pulumi:"platformUpdateDomain"`
	RdpThumbPrint             *string                                       `pulumi:"rdpThumbPrint"`
	Statuses                  []InstanceViewStatusResponse                  `pulumi:"statuses"`
	VmAgent                   *VirtualMachineAgentInstanceViewResponse      `pulumi:"vmAgent"`
	VmHealth                  VirtualMachineHealthStatusResponse            `pulumi:"vmHealth"`
}

type VirtualMachineInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineInstanceViewResponseOutput) ToVirtualMachineInstanceViewResponseOutput() VirtualMachineInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineInstanceViewResponseOutput) ToVirtualMachineInstanceViewResponseOutputWithContext(ctx context.Context) VirtualMachineInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineInstanceViewResponseOutput) AssignedHost() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) string { return v.AssignedHost }).(pulumi.StringOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) BootDiagnostics() BootDiagnosticsInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *BootDiagnosticsInstanceViewResponse {
		return v.BootDiagnostics
	}).(BootDiagnosticsInstanceViewResponsePtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) Disks() DiskInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) []DiskInstanceViewResponse { return v.Disks }).(DiskInstanceViewResponseArrayOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) Extensions() VirtualMachineExtensionInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) []VirtualMachineExtensionInstanceViewResponse {
		return v.Extensions
	}).(VirtualMachineExtensionInstanceViewResponseArrayOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *string { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) MaintenanceRedeployStatus() MaintenanceRedeployStatusResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *MaintenanceRedeployStatusResponse {
		return v.MaintenanceRedeployStatus
	}).(MaintenanceRedeployStatusResponsePtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) OsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *string { return v.OsName }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) OsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *string { return v.OsVersion }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) PatchStatus() VirtualMachinePatchStatusResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *VirtualMachinePatchStatusResponse { return v.PatchStatus }).(VirtualMachinePatchStatusResponsePtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) PlatformFaultDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *int { return v.PlatformFaultDomain }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) PlatformUpdateDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *int { return v.PlatformUpdateDomain }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) RdpThumbPrint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *string { return v.RdpThumbPrint }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) VmAgent() VirtualMachineAgentInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) *VirtualMachineAgentInstanceViewResponse { return v.VmAgent }).(VirtualMachineAgentInstanceViewResponsePtrOutput)
}

func (o VirtualMachineInstanceViewResponseOutput) VmHealth() VirtualMachineHealthStatusResponseOutput {
	return o.ApplyT(func(v VirtualMachineInstanceViewResponse) VirtualMachineHealthStatusResponse { return v.VmHealth }).(VirtualMachineHealthStatusResponseOutput)
}

type VirtualMachineIpTag struct {
	IpTagType *string `pulumi:"ipTagType"`
	Tag       *string `pulumi:"tag"`
}





type VirtualMachineIpTagInput interface {
	pulumi.Input

	ToVirtualMachineIpTagOutput() VirtualMachineIpTagOutput
	ToVirtualMachineIpTagOutputWithContext(context.Context) VirtualMachineIpTagOutput
}

type VirtualMachineIpTagArgs struct {
	IpTagType pulumi.StringPtrInput `pulumi:"ipTagType"`
	Tag       pulumi.StringPtrInput `pulumi:"tag"`
}

func (VirtualMachineIpTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineIpTag)(nil)).Elem()
}

func (i VirtualMachineIpTagArgs) ToVirtualMachineIpTagOutput() VirtualMachineIpTagOutput {
	return i.ToVirtualMachineIpTagOutputWithContext(context.Background())
}

func (i VirtualMachineIpTagArgs) ToVirtualMachineIpTagOutputWithContext(ctx context.Context) VirtualMachineIpTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineIpTagOutput)
}





type VirtualMachineIpTagArrayInput interface {
	pulumi.Input

	ToVirtualMachineIpTagArrayOutput() VirtualMachineIpTagArrayOutput
	ToVirtualMachineIpTagArrayOutputWithContext(context.Context) VirtualMachineIpTagArrayOutput
}

type VirtualMachineIpTagArray []VirtualMachineIpTagInput

func (VirtualMachineIpTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineIpTag)(nil)).Elem()
}

func (i VirtualMachineIpTagArray) ToVirtualMachineIpTagArrayOutput() VirtualMachineIpTagArrayOutput {
	return i.ToVirtualMachineIpTagArrayOutputWithContext(context.Background())
}

func (i VirtualMachineIpTagArray) ToVirtualMachineIpTagArrayOutputWithContext(ctx context.Context) VirtualMachineIpTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineIpTagArrayOutput)
}

type VirtualMachineIpTagOutput struct{ *pulumi.OutputState }

func (VirtualMachineIpTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineIpTag)(nil)).Elem()
}

func (o VirtualMachineIpTagOutput) ToVirtualMachineIpTagOutput() VirtualMachineIpTagOutput {
	return o
}

func (o VirtualMachineIpTagOutput) ToVirtualMachineIpTagOutputWithContext(ctx context.Context) VirtualMachineIpTagOutput {
	return o
}

func (o VirtualMachineIpTagOutput) IpTagType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineIpTag) *string { return v.IpTagType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineIpTagOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineIpTag) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

type VirtualMachineIpTagArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineIpTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineIpTag)(nil)).Elem()
}

func (o VirtualMachineIpTagArrayOutput) ToVirtualMachineIpTagArrayOutput() VirtualMachineIpTagArrayOutput {
	return o
}

func (o VirtualMachineIpTagArrayOutput) ToVirtualMachineIpTagArrayOutputWithContext(ctx context.Context) VirtualMachineIpTagArrayOutput {
	return o
}

func (o VirtualMachineIpTagArrayOutput) Index(i pulumi.IntInput) VirtualMachineIpTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineIpTag {
		return vs[0].([]VirtualMachineIpTag)[vs[1].(int)]
	}).(VirtualMachineIpTagOutput)
}

type VirtualMachineIpTagResponse struct {
	IpTagType *string `pulumi:"ipTagType"`
	Tag       *string `pulumi:"tag"`
}

type VirtualMachineIpTagResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineIpTagResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineIpTagResponse)(nil)).Elem()
}

func (o VirtualMachineIpTagResponseOutput) ToVirtualMachineIpTagResponseOutput() VirtualMachineIpTagResponseOutput {
	return o
}

func (o VirtualMachineIpTagResponseOutput) ToVirtualMachineIpTagResponseOutputWithContext(ctx context.Context) VirtualMachineIpTagResponseOutput {
	return o
}

func (o VirtualMachineIpTagResponseOutput) IpTagType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineIpTagResponse) *string { return v.IpTagType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineIpTagResponseOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineIpTagResponse) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

type VirtualMachineIpTagResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineIpTagResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineIpTagResponse)(nil)).Elem()
}

func (o VirtualMachineIpTagResponseArrayOutput) ToVirtualMachineIpTagResponseArrayOutput() VirtualMachineIpTagResponseArrayOutput {
	return o
}

func (o VirtualMachineIpTagResponseArrayOutput) ToVirtualMachineIpTagResponseArrayOutputWithContext(ctx context.Context) VirtualMachineIpTagResponseArrayOutput {
	return o
}

func (o VirtualMachineIpTagResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineIpTagResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineIpTagResponse {
		return vs[0].([]VirtualMachineIpTagResponse)[vs[1].(int)]
	}).(VirtualMachineIpTagResponseOutput)
}

type VirtualMachineNetworkInterfaceConfiguration struct {
	DeleteOption                *string                                                 `pulumi:"deleteOption"`
	DnsSettings                 *VirtualMachineNetworkInterfaceDnsSettingsConfiguration `pulumi:"dnsSettings"`
	DscpConfiguration           *SubResource                                            `pulumi:"dscpConfiguration"`
	EnableAcceleratedNetworking *bool                                                   `pulumi:"enableAcceleratedNetworking"`
	EnableFpga                  *bool                                                   `pulumi:"enableFpga"`
	EnableIPForwarding          *bool                                                   `pulumi:"enableIPForwarding"`
	IpConfigurations            []VirtualMachineNetworkInterfaceIPConfiguration         `pulumi:"ipConfigurations"`
	Name                        string                                                  `pulumi:"name"`
	NetworkSecurityGroup        *SubResource                                            `pulumi:"networkSecurityGroup"`
	Primary                     *bool                                                   `pulumi:"primary"`
}





type VirtualMachineNetworkInterfaceConfigurationInput interface {
	pulumi.Input

	ToVirtualMachineNetworkInterfaceConfigurationOutput() VirtualMachineNetworkInterfaceConfigurationOutput
	ToVirtualMachineNetworkInterfaceConfigurationOutputWithContext(context.Context) VirtualMachineNetworkInterfaceConfigurationOutput
}

type VirtualMachineNetworkInterfaceConfigurationArgs struct {
	DeleteOption                pulumi.StringPtrInput                                          `pulumi:"deleteOption"`
	DnsSettings                 VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrInput `pulumi:"dnsSettings"`
	DscpConfiguration           SubResourcePtrInput                                            `pulumi:"dscpConfiguration"`
	EnableAcceleratedNetworking pulumi.BoolPtrInput                                            `pulumi:"enableAcceleratedNetworking"`
	EnableFpga                  pulumi.BoolPtrInput                                            `pulumi:"enableFpga"`
	EnableIPForwarding          pulumi.BoolPtrInput                                            `pulumi:"enableIPForwarding"`
	IpConfigurations            VirtualMachineNetworkInterfaceIPConfigurationArrayInput        `pulumi:"ipConfigurations"`
	Name                        pulumi.StringInput                                             `pulumi:"name"`
	NetworkSecurityGroup        SubResourcePtrInput                                            `pulumi:"networkSecurityGroup"`
	Primary                     pulumi.BoolPtrInput                                            `pulumi:"primary"`
}

func (VirtualMachineNetworkInterfaceConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineNetworkInterfaceConfiguration)(nil)).Elem()
}

func (i VirtualMachineNetworkInterfaceConfigurationArgs) ToVirtualMachineNetworkInterfaceConfigurationOutput() VirtualMachineNetworkInterfaceConfigurationOutput {
	return i.ToVirtualMachineNetworkInterfaceConfigurationOutputWithContext(context.Background())
}

func (i VirtualMachineNetworkInterfaceConfigurationArgs) ToVirtualMachineNetworkInterfaceConfigurationOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineNetworkInterfaceConfigurationOutput)
}





type VirtualMachineNetworkInterfaceConfigurationArrayInput interface {
	pulumi.Input

	ToVirtualMachineNetworkInterfaceConfigurationArrayOutput() VirtualMachineNetworkInterfaceConfigurationArrayOutput
	ToVirtualMachineNetworkInterfaceConfigurationArrayOutputWithContext(context.Context) VirtualMachineNetworkInterfaceConfigurationArrayOutput
}

type VirtualMachineNetworkInterfaceConfigurationArray []VirtualMachineNetworkInterfaceConfigurationInput

func (VirtualMachineNetworkInterfaceConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineNetworkInterfaceConfiguration)(nil)).Elem()
}

func (i VirtualMachineNetworkInterfaceConfigurationArray) ToVirtualMachineNetworkInterfaceConfigurationArrayOutput() VirtualMachineNetworkInterfaceConfigurationArrayOutput {
	return i.ToVirtualMachineNetworkInterfaceConfigurationArrayOutputWithContext(context.Background())
}

func (i VirtualMachineNetworkInterfaceConfigurationArray) ToVirtualMachineNetworkInterfaceConfigurationArrayOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineNetworkInterfaceConfigurationArrayOutput)
}

type VirtualMachineNetworkInterfaceConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualMachineNetworkInterfaceConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineNetworkInterfaceConfiguration)(nil)).Elem()
}

func (o VirtualMachineNetworkInterfaceConfigurationOutput) ToVirtualMachineNetworkInterfaceConfigurationOutput() VirtualMachineNetworkInterfaceConfigurationOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceConfigurationOutput) ToVirtualMachineNetworkInterfaceConfigurationOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceConfigurationOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceConfigurationOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfiguration) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineNetworkInterfaceConfigurationOutput) DnsSettings() VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfiguration) *VirtualMachineNetworkInterfaceDnsSettingsConfiguration {
		return v.DnsSettings
	}).(VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput)
}

func (o VirtualMachineNetworkInterfaceConfigurationOutput) DscpConfiguration() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfiguration) *SubResource { return v.DscpConfiguration }).(SubResourcePtrOutput)
}

func (o VirtualMachineNetworkInterfaceConfigurationOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfiguration) *bool { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineNetworkInterfaceConfigurationOutput) EnableFpga() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfiguration) *bool { return v.EnableFpga }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineNetworkInterfaceConfigurationOutput) EnableIPForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfiguration) *bool { return v.EnableIPForwarding }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineNetworkInterfaceConfigurationOutput) IpConfigurations() VirtualMachineNetworkInterfaceIPConfigurationArrayOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfiguration) []VirtualMachineNetworkInterfaceIPConfiguration {
		return v.IpConfigurations
	}).(VirtualMachineNetworkInterfaceIPConfigurationArrayOutput)
}

func (o VirtualMachineNetworkInterfaceConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineNetworkInterfaceConfigurationOutput) NetworkSecurityGroup() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfiguration) *SubResource { return v.NetworkSecurityGroup }).(SubResourcePtrOutput)
}

func (o VirtualMachineNetworkInterfaceConfigurationOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfiguration) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

type VirtualMachineNetworkInterfaceConfigurationArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineNetworkInterfaceConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineNetworkInterfaceConfiguration)(nil)).Elem()
}

func (o VirtualMachineNetworkInterfaceConfigurationArrayOutput) ToVirtualMachineNetworkInterfaceConfigurationArrayOutput() VirtualMachineNetworkInterfaceConfigurationArrayOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceConfigurationArrayOutput) ToVirtualMachineNetworkInterfaceConfigurationArrayOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceConfigurationArrayOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceConfigurationArrayOutput) Index(i pulumi.IntInput) VirtualMachineNetworkInterfaceConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineNetworkInterfaceConfiguration {
		return vs[0].([]VirtualMachineNetworkInterfaceConfiguration)[vs[1].(int)]
	}).(VirtualMachineNetworkInterfaceConfigurationOutput)
}

type VirtualMachineNetworkInterfaceConfigurationResponse struct {
	DeleteOption                *string                                                         `pulumi:"deleteOption"`
	DnsSettings                 *VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponse `pulumi:"dnsSettings"`
	DscpConfiguration           *SubResourceResponse                                            `pulumi:"dscpConfiguration"`
	EnableAcceleratedNetworking *bool                                                           `pulumi:"enableAcceleratedNetworking"`
	EnableFpga                  *bool                                                           `pulumi:"enableFpga"`
	EnableIPForwarding          *bool                                                           `pulumi:"enableIPForwarding"`
	IpConfigurations            []VirtualMachineNetworkInterfaceIPConfigurationResponse         `pulumi:"ipConfigurations"`
	Name                        string                                                          `pulumi:"name"`
	NetworkSecurityGroup        *SubResourceResponse                                            `pulumi:"networkSecurityGroup"`
	Primary                     *bool                                                           `pulumi:"primary"`
}

type VirtualMachineNetworkInterfaceConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineNetworkInterfaceConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineNetworkInterfaceConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineNetworkInterfaceConfigurationResponseOutput) ToVirtualMachineNetworkInterfaceConfigurationResponseOutput() VirtualMachineNetworkInterfaceConfigurationResponseOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceConfigurationResponseOutput) ToVirtualMachineNetworkInterfaceConfigurationResponseOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceConfigurationResponseOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceConfigurationResponseOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfigurationResponse) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineNetworkInterfaceConfigurationResponseOutput) DnsSettings() VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfigurationResponse) *VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponse {
		return v.DnsSettings
	}).(VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponsePtrOutput)
}

func (o VirtualMachineNetworkInterfaceConfigurationResponseOutput) DscpConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfigurationResponse) *SubResourceResponse {
		return v.DscpConfiguration
	}).(SubResourceResponsePtrOutput)
}

func (o VirtualMachineNetworkInterfaceConfigurationResponseOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfigurationResponse) *bool {
		return v.EnableAcceleratedNetworking
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineNetworkInterfaceConfigurationResponseOutput) EnableFpga() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfigurationResponse) *bool { return v.EnableFpga }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineNetworkInterfaceConfigurationResponseOutput) EnableIPForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfigurationResponse) *bool { return v.EnableIPForwarding }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineNetworkInterfaceConfigurationResponseOutput) IpConfigurations() VirtualMachineNetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfigurationResponse) []VirtualMachineNetworkInterfaceIPConfigurationResponse {
		return v.IpConfigurations
	}).(VirtualMachineNetworkInterfaceIPConfigurationResponseArrayOutput)
}

func (o VirtualMachineNetworkInterfaceConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineNetworkInterfaceConfigurationResponseOutput) NetworkSecurityGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfigurationResponse) *SubResourceResponse {
		return v.NetworkSecurityGroup
	}).(SubResourceResponsePtrOutput)
}

func (o VirtualMachineNetworkInterfaceConfigurationResponseOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceConfigurationResponse) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

type VirtualMachineNetworkInterfaceConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineNetworkInterfaceConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineNetworkInterfaceConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineNetworkInterfaceConfigurationResponseArrayOutput) ToVirtualMachineNetworkInterfaceConfigurationResponseArrayOutput() VirtualMachineNetworkInterfaceConfigurationResponseArrayOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceConfigurationResponseArrayOutput) ToVirtualMachineNetworkInterfaceConfigurationResponseArrayOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceConfigurationResponseArrayOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceConfigurationResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineNetworkInterfaceConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineNetworkInterfaceConfigurationResponse {
		return vs[0].([]VirtualMachineNetworkInterfaceConfigurationResponse)[vs[1].(int)]
	}).(VirtualMachineNetworkInterfaceConfigurationResponseOutput)
}

type VirtualMachineNetworkInterfaceDnsSettingsConfiguration struct {
	DnsServers []string `pulumi:"dnsServers"`
}





type VirtualMachineNetworkInterfaceDnsSettingsConfigurationInput interface {
	pulumi.Input

	ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput() VirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput
	ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationOutputWithContext(context.Context) VirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput
}

type VirtualMachineNetworkInterfaceDnsSettingsConfigurationArgs struct {
	DnsServers pulumi.StringArrayInput `pulumi:"dnsServers"`
}

func (VirtualMachineNetworkInterfaceDnsSettingsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineNetworkInterfaceDnsSettingsConfiguration)(nil)).Elem()
}

func (i VirtualMachineNetworkInterfaceDnsSettingsConfigurationArgs) ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput() VirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput {
	return i.ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationOutputWithContext(context.Background())
}

func (i VirtualMachineNetworkInterfaceDnsSettingsConfigurationArgs) ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput)
}

func (i VirtualMachineNetworkInterfaceDnsSettingsConfigurationArgs) ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput() VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput {
	return i.ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutputWithContext(context.Background())
}

func (i VirtualMachineNetworkInterfaceDnsSettingsConfigurationArgs) ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput).ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutputWithContext(ctx)
}









type VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrInput interface {
	pulumi.Input

	ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput() VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput
	ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutputWithContext(context.Context) VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput
}

type virtualMachineNetworkInterfaceDnsSettingsConfigurationPtrType VirtualMachineNetworkInterfaceDnsSettingsConfigurationArgs

func VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtr(v *VirtualMachineNetworkInterfaceDnsSettingsConfigurationArgs) VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrInput {
	return (*virtualMachineNetworkInterfaceDnsSettingsConfigurationPtrType)(v)
}

func (*virtualMachineNetworkInterfaceDnsSettingsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineNetworkInterfaceDnsSettingsConfiguration)(nil)).Elem()
}

func (i *virtualMachineNetworkInterfaceDnsSettingsConfigurationPtrType) ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput() VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput {
	return i.ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutputWithContext(context.Background())
}

func (i *virtualMachineNetworkInterfaceDnsSettingsConfigurationPtrType) ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput)
}

type VirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineNetworkInterfaceDnsSettingsConfiguration)(nil)).Elem()
}

func (o VirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput) ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput() VirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput) ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput) ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput() VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput {
	return o.ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutputWithContext(context.Background())
}

func (o VirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput) ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineNetworkInterfaceDnsSettingsConfiguration) *VirtualMachineNetworkInterfaceDnsSettingsConfiguration {
		return &v
	}).(VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput)
}

func (o VirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceDnsSettingsConfiguration) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

type VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineNetworkInterfaceDnsSettingsConfiguration)(nil)).Elem()
}

func (o VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput) ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput() VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput) ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput) Elem() VirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput {
	return o.ApplyT(func(v *VirtualMachineNetworkInterfaceDnsSettingsConfiguration) VirtualMachineNetworkInterfaceDnsSettingsConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualMachineNetworkInterfaceDnsSettingsConfiguration
		return ret
	}).(VirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput)
}

func (o VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineNetworkInterfaceDnsSettingsConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

type VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponse struct {
	DnsServers []string `pulumi:"dnsServers"`
}

type VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponseOutput) ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationResponseOutput() VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponseOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponseOutput) ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationResponseOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponseOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponseOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponse) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

type VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponsePtrOutput) ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationResponsePtrOutput() VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponsePtrOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponsePtrOutput) ToVirtualMachineNetworkInterfaceDnsSettingsConfigurationResponsePtrOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponsePtrOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponsePtrOutput) Elem() VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponseOutput {
	return o.ApplyT(func(v *VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponse) VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponse
		return ret
	}).(VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponseOutput)
}

func (o VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponsePtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

type VirtualMachineNetworkInterfaceIPConfiguration struct {
	ApplicationGatewayBackendAddressPools []SubResource                               `pulumi:"applicationGatewayBackendAddressPools"`
	ApplicationSecurityGroups             []SubResource                               `pulumi:"applicationSecurityGroups"`
	LoadBalancerBackendAddressPools       []SubResource                               `pulumi:"loadBalancerBackendAddressPools"`
	Name                                  string                                      `pulumi:"name"`
	Primary                               *bool                                       `pulumi:"primary"`
	PrivateIPAddressVersion               *string                                     `pulumi:"privateIPAddressVersion"`
	PublicIPAddressConfiguration          *VirtualMachinePublicIPAddressConfiguration `pulumi:"publicIPAddressConfiguration"`
	Subnet                                *SubResource                                `pulumi:"subnet"`
}





type VirtualMachineNetworkInterfaceIPConfigurationInput interface {
	pulumi.Input

	ToVirtualMachineNetworkInterfaceIPConfigurationOutput() VirtualMachineNetworkInterfaceIPConfigurationOutput
	ToVirtualMachineNetworkInterfaceIPConfigurationOutputWithContext(context.Context) VirtualMachineNetworkInterfaceIPConfigurationOutput
}

type VirtualMachineNetworkInterfaceIPConfigurationArgs struct {
	ApplicationGatewayBackendAddressPools SubResourceArrayInput                              `pulumi:"applicationGatewayBackendAddressPools"`
	ApplicationSecurityGroups             SubResourceArrayInput                              `pulumi:"applicationSecurityGroups"`
	LoadBalancerBackendAddressPools       SubResourceArrayInput                              `pulumi:"loadBalancerBackendAddressPools"`
	Name                                  pulumi.StringInput                                 `pulumi:"name"`
	Primary                               pulumi.BoolPtrInput                                `pulumi:"primary"`
	PrivateIPAddressVersion               pulumi.StringPtrInput                              `pulumi:"privateIPAddressVersion"`
	PublicIPAddressConfiguration          VirtualMachinePublicIPAddressConfigurationPtrInput `pulumi:"publicIPAddressConfiguration"`
	Subnet                                SubResourcePtrInput                                `pulumi:"subnet"`
}

func (VirtualMachineNetworkInterfaceIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineNetworkInterfaceIPConfiguration)(nil)).Elem()
}

func (i VirtualMachineNetworkInterfaceIPConfigurationArgs) ToVirtualMachineNetworkInterfaceIPConfigurationOutput() VirtualMachineNetworkInterfaceIPConfigurationOutput {
	return i.ToVirtualMachineNetworkInterfaceIPConfigurationOutputWithContext(context.Background())
}

func (i VirtualMachineNetworkInterfaceIPConfigurationArgs) ToVirtualMachineNetworkInterfaceIPConfigurationOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineNetworkInterfaceIPConfigurationOutput)
}





type VirtualMachineNetworkInterfaceIPConfigurationArrayInput interface {
	pulumi.Input

	ToVirtualMachineNetworkInterfaceIPConfigurationArrayOutput() VirtualMachineNetworkInterfaceIPConfigurationArrayOutput
	ToVirtualMachineNetworkInterfaceIPConfigurationArrayOutputWithContext(context.Context) VirtualMachineNetworkInterfaceIPConfigurationArrayOutput
}

type VirtualMachineNetworkInterfaceIPConfigurationArray []VirtualMachineNetworkInterfaceIPConfigurationInput

func (VirtualMachineNetworkInterfaceIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineNetworkInterfaceIPConfiguration)(nil)).Elem()
}

func (i VirtualMachineNetworkInterfaceIPConfigurationArray) ToVirtualMachineNetworkInterfaceIPConfigurationArrayOutput() VirtualMachineNetworkInterfaceIPConfigurationArrayOutput {
	return i.ToVirtualMachineNetworkInterfaceIPConfigurationArrayOutputWithContext(context.Background())
}

func (i VirtualMachineNetworkInterfaceIPConfigurationArray) ToVirtualMachineNetworkInterfaceIPConfigurationArrayOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineNetworkInterfaceIPConfigurationArrayOutput)
}

type VirtualMachineNetworkInterfaceIPConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualMachineNetworkInterfaceIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineNetworkInterfaceIPConfiguration)(nil)).Elem()
}

func (o VirtualMachineNetworkInterfaceIPConfigurationOutput) ToVirtualMachineNetworkInterfaceIPConfigurationOutput() VirtualMachineNetworkInterfaceIPConfigurationOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceIPConfigurationOutput) ToVirtualMachineNetworkInterfaceIPConfigurationOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceIPConfigurationOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceIPConfigurationOutput) ApplicationGatewayBackendAddressPools() SubResourceArrayOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceIPConfiguration) []SubResource {
		return v.ApplicationGatewayBackendAddressPools
	}).(SubResourceArrayOutput)
}

func (o VirtualMachineNetworkInterfaceIPConfigurationOutput) ApplicationSecurityGroups() SubResourceArrayOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceIPConfiguration) []SubResource {
		return v.ApplicationSecurityGroups
	}).(SubResourceArrayOutput)
}

func (o VirtualMachineNetworkInterfaceIPConfigurationOutput) LoadBalancerBackendAddressPools() SubResourceArrayOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceIPConfiguration) []SubResource {
		return v.LoadBalancerBackendAddressPools
	}).(SubResourceArrayOutput)
}

func (o VirtualMachineNetworkInterfaceIPConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceIPConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineNetworkInterfaceIPConfigurationOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceIPConfiguration) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineNetworkInterfaceIPConfigurationOutput) PrivateIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceIPConfiguration) *string { return v.PrivateIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineNetworkInterfaceIPConfigurationOutput) PublicIPAddressConfiguration() VirtualMachinePublicIPAddressConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceIPConfiguration) *VirtualMachinePublicIPAddressConfiguration {
		return v.PublicIPAddressConfiguration
	}).(VirtualMachinePublicIPAddressConfigurationPtrOutput)
}

func (o VirtualMachineNetworkInterfaceIPConfigurationOutput) Subnet() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceIPConfiguration) *SubResource { return v.Subnet }).(SubResourcePtrOutput)
}

type VirtualMachineNetworkInterfaceIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineNetworkInterfaceIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineNetworkInterfaceIPConfiguration)(nil)).Elem()
}

func (o VirtualMachineNetworkInterfaceIPConfigurationArrayOutput) ToVirtualMachineNetworkInterfaceIPConfigurationArrayOutput() VirtualMachineNetworkInterfaceIPConfigurationArrayOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceIPConfigurationArrayOutput) ToVirtualMachineNetworkInterfaceIPConfigurationArrayOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceIPConfigurationArrayOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceIPConfigurationArrayOutput) Index(i pulumi.IntInput) VirtualMachineNetworkInterfaceIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineNetworkInterfaceIPConfiguration {
		return vs[0].([]VirtualMachineNetworkInterfaceIPConfiguration)[vs[1].(int)]
	}).(VirtualMachineNetworkInterfaceIPConfigurationOutput)
}

type VirtualMachineNetworkInterfaceIPConfigurationResponse struct {
	ApplicationGatewayBackendAddressPools []SubResourceResponse                               `pulumi:"applicationGatewayBackendAddressPools"`
	ApplicationSecurityGroups             []SubResourceResponse                               `pulumi:"applicationSecurityGroups"`
	LoadBalancerBackendAddressPools       []SubResourceResponse                               `pulumi:"loadBalancerBackendAddressPools"`
	Name                                  string                                              `pulumi:"name"`
	Primary                               *bool                                               `pulumi:"primary"`
	PrivateIPAddressVersion               *string                                             `pulumi:"privateIPAddressVersion"`
	PublicIPAddressConfiguration          *VirtualMachinePublicIPAddressConfigurationResponse `pulumi:"publicIPAddressConfiguration"`
	Subnet                                *SubResourceResponse                                `pulumi:"subnet"`
}

type VirtualMachineNetworkInterfaceIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineNetworkInterfaceIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineNetworkInterfaceIPConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineNetworkInterfaceIPConfigurationResponseOutput) ToVirtualMachineNetworkInterfaceIPConfigurationResponseOutput() VirtualMachineNetworkInterfaceIPConfigurationResponseOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceIPConfigurationResponseOutput) ToVirtualMachineNetworkInterfaceIPConfigurationResponseOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceIPConfigurationResponseOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceIPConfigurationResponseOutput) ApplicationGatewayBackendAddressPools() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceIPConfigurationResponse) []SubResourceResponse {
		return v.ApplicationGatewayBackendAddressPools
	}).(SubResourceResponseArrayOutput)
}

func (o VirtualMachineNetworkInterfaceIPConfigurationResponseOutput) ApplicationSecurityGroups() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceIPConfigurationResponse) []SubResourceResponse {
		return v.ApplicationSecurityGroups
	}).(SubResourceResponseArrayOutput)
}

func (o VirtualMachineNetworkInterfaceIPConfigurationResponseOutput) LoadBalancerBackendAddressPools() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceIPConfigurationResponse) []SubResourceResponse {
		return v.LoadBalancerBackendAddressPools
	}).(SubResourceResponseArrayOutput)
}

func (o VirtualMachineNetworkInterfaceIPConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceIPConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineNetworkInterfaceIPConfigurationResponseOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceIPConfigurationResponse) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineNetworkInterfaceIPConfigurationResponseOutput) PrivateIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceIPConfigurationResponse) *string {
		return v.PrivateIPAddressVersion
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineNetworkInterfaceIPConfigurationResponseOutput) PublicIPAddressConfiguration() VirtualMachinePublicIPAddressConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceIPConfigurationResponse) *VirtualMachinePublicIPAddressConfigurationResponse {
		return v.PublicIPAddressConfiguration
	}).(VirtualMachinePublicIPAddressConfigurationResponsePtrOutput)
}

func (o VirtualMachineNetworkInterfaceIPConfigurationResponseOutput) Subnet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineNetworkInterfaceIPConfigurationResponse) *SubResourceResponse { return v.Subnet }).(SubResourceResponsePtrOutput)
}

type VirtualMachineNetworkInterfaceIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineNetworkInterfaceIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineNetworkInterfaceIPConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineNetworkInterfaceIPConfigurationResponseArrayOutput) ToVirtualMachineNetworkInterfaceIPConfigurationResponseArrayOutput() VirtualMachineNetworkInterfaceIPConfigurationResponseArrayOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceIPConfigurationResponseArrayOutput) ToVirtualMachineNetworkInterfaceIPConfigurationResponseArrayOutputWithContext(ctx context.Context) VirtualMachineNetworkInterfaceIPConfigurationResponseArrayOutput {
	return o
}

func (o VirtualMachineNetworkInterfaceIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineNetworkInterfaceIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineNetworkInterfaceIPConfigurationResponse {
		return vs[0].([]VirtualMachineNetworkInterfaceIPConfigurationResponse)[vs[1].(int)]
	}).(VirtualMachineNetworkInterfaceIPConfigurationResponseOutput)
}

type VirtualMachinePatchStatusResponse struct {
	AvailablePatchSummary        *AvailablePatchSummaryResponse        `pulumi:"availablePatchSummary"`
	ConfigurationStatuses        []InstanceViewStatusResponse          `pulumi:"configurationStatuses"`
	LastPatchInstallationSummary *LastPatchInstallationSummaryResponse `pulumi:"lastPatchInstallationSummary"`
}

type VirtualMachinePatchStatusResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachinePatchStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachinePatchStatusResponse)(nil)).Elem()
}

func (o VirtualMachinePatchStatusResponseOutput) ToVirtualMachinePatchStatusResponseOutput() VirtualMachinePatchStatusResponseOutput {
	return o
}

func (o VirtualMachinePatchStatusResponseOutput) ToVirtualMachinePatchStatusResponseOutputWithContext(ctx context.Context) VirtualMachinePatchStatusResponseOutput {
	return o
}

func (o VirtualMachinePatchStatusResponseOutput) AvailablePatchSummary() AvailablePatchSummaryResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachinePatchStatusResponse) *AvailablePatchSummaryResponse {
		return v.AvailablePatchSummary
	}).(AvailablePatchSummaryResponsePtrOutput)
}

func (o VirtualMachinePatchStatusResponseOutput) ConfigurationStatuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachinePatchStatusResponse) []InstanceViewStatusResponse { return v.ConfigurationStatuses }).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachinePatchStatusResponseOutput) LastPatchInstallationSummary() LastPatchInstallationSummaryResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachinePatchStatusResponse) *LastPatchInstallationSummaryResponse {
		return v.LastPatchInstallationSummary
	}).(LastPatchInstallationSummaryResponsePtrOutput)
}

type VirtualMachinePatchStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachinePatchStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachinePatchStatusResponse)(nil)).Elem()
}

func (o VirtualMachinePatchStatusResponsePtrOutput) ToVirtualMachinePatchStatusResponsePtrOutput() VirtualMachinePatchStatusResponsePtrOutput {
	return o
}

func (o VirtualMachinePatchStatusResponsePtrOutput) ToVirtualMachinePatchStatusResponsePtrOutputWithContext(ctx context.Context) VirtualMachinePatchStatusResponsePtrOutput {
	return o
}

func (o VirtualMachinePatchStatusResponsePtrOutput) Elem() VirtualMachinePatchStatusResponseOutput {
	return o.ApplyT(func(v *VirtualMachinePatchStatusResponse) VirtualMachinePatchStatusResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachinePatchStatusResponse
		return ret
	}).(VirtualMachinePatchStatusResponseOutput)
}

func (o VirtualMachinePatchStatusResponsePtrOutput) AvailablePatchSummary() AvailablePatchSummaryResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachinePatchStatusResponse) *AvailablePatchSummaryResponse {
		if v == nil {
			return nil
		}
		return v.AvailablePatchSummary
	}).(AvailablePatchSummaryResponsePtrOutput)
}

func (o VirtualMachinePatchStatusResponsePtrOutput) ConfigurationStatuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachinePatchStatusResponse) []InstanceViewStatusResponse {
		if v == nil {
			return nil
		}
		return v.ConfigurationStatuses
	}).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachinePatchStatusResponsePtrOutput) LastPatchInstallationSummary() LastPatchInstallationSummaryResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachinePatchStatusResponse) *LastPatchInstallationSummaryResponse {
		if v == nil {
			return nil
		}
		return v.LastPatchInstallationSummary
	}).(LastPatchInstallationSummaryResponsePtrOutput)
}

type VirtualMachinePublicIPAddressConfiguration struct {
	DeleteOption             *string                                                `pulumi:"deleteOption"`
	DnsSettings              *VirtualMachinePublicIPAddressDnsSettingsConfiguration `pulumi:"dnsSettings"`
	IdleTimeoutInMinutes     *int                                                   `pulumi:"idleTimeoutInMinutes"`
	IpTags                   []VirtualMachineIpTag                                  `pulumi:"ipTags"`
	Name                     string                                                 `pulumi:"name"`
	PublicIPAddressVersion   *string                                                `pulumi:"publicIPAddressVersion"`
	PublicIPAllocationMethod *string                                                `pulumi:"publicIPAllocationMethod"`
	PublicIPPrefix           *SubResource                                           `pulumi:"publicIPPrefix"`
	Sku                      *PublicIPAddressSku                                    `pulumi:"sku"`
}





type VirtualMachinePublicIPAddressConfigurationInput interface {
	pulumi.Input

	ToVirtualMachinePublicIPAddressConfigurationOutput() VirtualMachinePublicIPAddressConfigurationOutput
	ToVirtualMachinePublicIPAddressConfigurationOutputWithContext(context.Context) VirtualMachinePublicIPAddressConfigurationOutput
}

type VirtualMachinePublicIPAddressConfigurationArgs struct {
	DeleteOption             pulumi.StringPtrInput                                         `pulumi:"deleteOption"`
	DnsSettings              VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrInput `pulumi:"dnsSettings"`
	IdleTimeoutInMinutes     pulumi.IntPtrInput                                            `pulumi:"idleTimeoutInMinutes"`
	IpTags                   VirtualMachineIpTagArrayInput                                 `pulumi:"ipTags"`
	Name                     pulumi.StringInput                                            `pulumi:"name"`
	PublicIPAddressVersion   pulumi.StringPtrInput                                         `pulumi:"publicIPAddressVersion"`
	PublicIPAllocationMethod pulumi.StringPtrInput                                         `pulumi:"publicIPAllocationMethod"`
	PublicIPPrefix           SubResourcePtrInput                                           `pulumi:"publicIPPrefix"`
	Sku                      PublicIPAddressSkuPtrInput                                    `pulumi:"sku"`
}

func (VirtualMachinePublicIPAddressConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachinePublicIPAddressConfiguration)(nil)).Elem()
}

func (i VirtualMachinePublicIPAddressConfigurationArgs) ToVirtualMachinePublicIPAddressConfigurationOutput() VirtualMachinePublicIPAddressConfigurationOutput {
	return i.ToVirtualMachinePublicIPAddressConfigurationOutputWithContext(context.Background())
}

func (i VirtualMachinePublicIPAddressConfigurationArgs) ToVirtualMachinePublicIPAddressConfigurationOutputWithContext(ctx context.Context) VirtualMachinePublicIPAddressConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachinePublicIPAddressConfigurationOutput)
}

func (i VirtualMachinePublicIPAddressConfigurationArgs) ToVirtualMachinePublicIPAddressConfigurationPtrOutput() VirtualMachinePublicIPAddressConfigurationPtrOutput {
	return i.ToVirtualMachinePublicIPAddressConfigurationPtrOutputWithContext(context.Background())
}

func (i VirtualMachinePublicIPAddressConfigurationArgs) ToVirtualMachinePublicIPAddressConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachinePublicIPAddressConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachinePublicIPAddressConfigurationOutput).ToVirtualMachinePublicIPAddressConfigurationPtrOutputWithContext(ctx)
}









type VirtualMachinePublicIPAddressConfigurationPtrInput interface {
	pulumi.Input

	ToVirtualMachinePublicIPAddressConfigurationPtrOutput() VirtualMachinePublicIPAddressConfigurationPtrOutput
	ToVirtualMachinePublicIPAddressConfigurationPtrOutputWithContext(context.Context) VirtualMachinePublicIPAddressConfigurationPtrOutput
}

type virtualMachinePublicIPAddressConfigurationPtrType VirtualMachinePublicIPAddressConfigurationArgs

func VirtualMachinePublicIPAddressConfigurationPtr(v *VirtualMachinePublicIPAddressConfigurationArgs) VirtualMachinePublicIPAddressConfigurationPtrInput {
	return (*virtualMachinePublicIPAddressConfigurationPtrType)(v)
}

func (*virtualMachinePublicIPAddressConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachinePublicIPAddressConfiguration)(nil)).Elem()
}

func (i *virtualMachinePublicIPAddressConfigurationPtrType) ToVirtualMachinePublicIPAddressConfigurationPtrOutput() VirtualMachinePublicIPAddressConfigurationPtrOutput {
	return i.ToVirtualMachinePublicIPAddressConfigurationPtrOutputWithContext(context.Background())
}

func (i *virtualMachinePublicIPAddressConfigurationPtrType) ToVirtualMachinePublicIPAddressConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachinePublicIPAddressConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachinePublicIPAddressConfigurationPtrOutput)
}

type VirtualMachinePublicIPAddressConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualMachinePublicIPAddressConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachinePublicIPAddressConfiguration)(nil)).Elem()
}

func (o VirtualMachinePublicIPAddressConfigurationOutput) ToVirtualMachinePublicIPAddressConfigurationOutput() VirtualMachinePublicIPAddressConfigurationOutput {
	return o
}

func (o VirtualMachinePublicIPAddressConfigurationOutput) ToVirtualMachinePublicIPAddressConfigurationOutputWithContext(ctx context.Context) VirtualMachinePublicIPAddressConfigurationOutput {
	return o
}

func (o VirtualMachinePublicIPAddressConfigurationOutput) ToVirtualMachinePublicIPAddressConfigurationPtrOutput() VirtualMachinePublicIPAddressConfigurationPtrOutput {
	return o.ToVirtualMachinePublicIPAddressConfigurationPtrOutputWithContext(context.Background())
}

func (o VirtualMachinePublicIPAddressConfigurationOutput) ToVirtualMachinePublicIPAddressConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachinePublicIPAddressConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachinePublicIPAddressConfiguration) *VirtualMachinePublicIPAddressConfiguration {
		return &v
	}).(VirtualMachinePublicIPAddressConfigurationPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressConfiguration) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationOutput) DnsSettings() VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressConfiguration) *VirtualMachinePublicIPAddressDnsSettingsConfiguration {
		return v.DnsSettings
	}).(VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressConfiguration) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationOutput) IpTags() VirtualMachineIpTagArrayOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressConfiguration) []VirtualMachineIpTag { return v.IpTags }).(VirtualMachineIpTagArrayOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressConfiguration) *string { return v.PublicIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationOutput) PublicIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressConfiguration) *string { return v.PublicIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationOutput) PublicIPPrefix() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressConfiguration) *SubResource { return v.PublicIPPrefix }).(SubResourcePtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationOutput) Sku() PublicIPAddressSkuPtrOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressConfiguration) *PublicIPAddressSku { return v.Sku }).(PublicIPAddressSkuPtrOutput)
}

type VirtualMachinePublicIPAddressConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachinePublicIPAddressConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachinePublicIPAddressConfiguration)(nil)).Elem()
}

func (o VirtualMachinePublicIPAddressConfigurationPtrOutput) ToVirtualMachinePublicIPAddressConfigurationPtrOutput() VirtualMachinePublicIPAddressConfigurationPtrOutput {
	return o
}

func (o VirtualMachinePublicIPAddressConfigurationPtrOutput) ToVirtualMachinePublicIPAddressConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachinePublicIPAddressConfigurationPtrOutput {
	return o
}

func (o VirtualMachinePublicIPAddressConfigurationPtrOutput) Elem() VirtualMachinePublicIPAddressConfigurationOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfiguration) VirtualMachinePublicIPAddressConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualMachinePublicIPAddressConfiguration
		return ret
	}).(VirtualMachinePublicIPAddressConfigurationOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationPtrOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.DeleteOption
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationPtrOutput) DnsSettings() VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfiguration) *VirtualMachinePublicIPAddressDnsSettingsConfiguration {
		if v == nil {
			return nil
		}
		return v.DnsSettings
	}).(VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationPtrOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfiguration) *int {
		if v == nil {
			return nil
		}
		return v.IdleTimeoutInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationPtrOutput) IpTags() VirtualMachineIpTagArrayOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfiguration) []VirtualMachineIpTag {
		if v == nil {
			return nil
		}
		return v.IpTags
	}).(VirtualMachineIpTagArrayOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationPtrOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.PublicIPAddressVersion
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationPtrOutput) PublicIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.PublicIPAllocationMethod
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationPtrOutput) PublicIPPrefix() SubResourcePtrOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfiguration) *SubResource {
		if v == nil {
			return nil
		}
		return v.PublicIPPrefix
	}).(SubResourcePtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationPtrOutput) Sku() PublicIPAddressSkuPtrOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfiguration) *PublicIPAddressSku {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(PublicIPAddressSkuPtrOutput)
}

type VirtualMachinePublicIPAddressConfigurationResponse struct {
	DeleteOption             *string                                                        `pulumi:"deleteOption"`
	DnsSettings              *VirtualMachinePublicIPAddressDnsSettingsConfigurationResponse `pulumi:"dnsSettings"`
	IdleTimeoutInMinutes     *int                                                           `pulumi:"idleTimeoutInMinutes"`
	IpTags                   []VirtualMachineIpTagResponse                                  `pulumi:"ipTags"`
	Name                     string                                                         `pulumi:"name"`
	PublicIPAddressVersion   *string                                                        `pulumi:"publicIPAddressVersion"`
	PublicIPAllocationMethod *string                                                        `pulumi:"publicIPAllocationMethod"`
	PublicIPPrefix           *SubResourceResponse                                           `pulumi:"publicIPPrefix"`
	Sku                      *PublicIPAddressSkuResponse                                    `pulumi:"sku"`
}

type VirtualMachinePublicIPAddressConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachinePublicIPAddressConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachinePublicIPAddressConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachinePublicIPAddressConfigurationResponseOutput) ToVirtualMachinePublicIPAddressConfigurationResponseOutput() VirtualMachinePublicIPAddressConfigurationResponseOutput {
	return o
}

func (o VirtualMachinePublicIPAddressConfigurationResponseOutput) ToVirtualMachinePublicIPAddressConfigurationResponseOutputWithContext(ctx context.Context) VirtualMachinePublicIPAddressConfigurationResponseOutput {
	return o
}

func (o VirtualMachinePublicIPAddressConfigurationResponseOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressConfigurationResponse) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationResponseOutput) DnsSettings() VirtualMachinePublicIPAddressDnsSettingsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressConfigurationResponse) *VirtualMachinePublicIPAddressDnsSettingsConfigurationResponse {
		return v.DnsSettings
	}).(VirtualMachinePublicIPAddressDnsSettingsConfigurationResponsePtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationResponseOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressConfigurationResponse) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationResponseOutput) IpTags() VirtualMachineIpTagResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressConfigurationResponse) []VirtualMachineIpTagResponse {
		return v.IpTags
	}).(VirtualMachineIpTagResponseArrayOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationResponseOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressConfigurationResponse) *string { return v.PublicIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationResponseOutput) PublicIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressConfigurationResponse) *string { return v.PublicIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationResponseOutput) PublicIPPrefix() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressConfigurationResponse) *SubResourceResponse {
		return v.PublicIPPrefix
	}).(SubResourceResponsePtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationResponseOutput) Sku() PublicIPAddressSkuResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressConfigurationResponse) *PublicIPAddressSkuResponse { return v.Sku }).(PublicIPAddressSkuResponsePtrOutput)
}

type VirtualMachinePublicIPAddressConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachinePublicIPAddressConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachinePublicIPAddressConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachinePublicIPAddressConfigurationResponsePtrOutput) ToVirtualMachinePublicIPAddressConfigurationResponsePtrOutput() VirtualMachinePublicIPAddressConfigurationResponsePtrOutput {
	return o
}

func (o VirtualMachinePublicIPAddressConfigurationResponsePtrOutput) ToVirtualMachinePublicIPAddressConfigurationResponsePtrOutputWithContext(ctx context.Context) VirtualMachinePublicIPAddressConfigurationResponsePtrOutput {
	return o
}

func (o VirtualMachinePublicIPAddressConfigurationResponsePtrOutput) Elem() VirtualMachinePublicIPAddressConfigurationResponseOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfigurationResponse) VirtualMachinePublicIPAddressConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachinePublicIPAddressConfigurationResponse
		return ret
	}).(VirtualMachinePublicIPAddressConfigurationResponseOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationResponsePtrOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeleteOption
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationResponsePtrOutput) DnsSettings() VirtualMachinePublicIPAddressDnsSettingsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfigurationResponse) *VirtualMachinePublicIPAddressDnsSettingsConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.DnsSettings
	}).(VirtualMachinePublicIPAddressDnsSettingsConfigurationResponsePtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationResponsePtrOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfigurationResponse) *int {
		if v == nil {
			return nil
		}
		return v.IdleTimeoutInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationResponsePtrOutput) IpTags() VirtualMachineIpTagResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfigurationResponse) []VirtualMachineIpTagResponse {
		if v == nil {
			return nil
		}
		return v.IpTags
	}).(VirtualMachineIpTagResponseArrayOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationResponsePtrOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicIPAddressVersion
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationResponsePtrOutput) PublicIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicIPAllocationMethod
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationResponsePtrOutput) PublicIPPrefix() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfigurationResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return v.PublicIPPrefix
	}).(SubResourceResponsePtrOutput)
}

func (o VirtualMachinePublicIPAddressConfigurationResponsePtrOutput) Sku() PublicIPAddressSkuResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressConfigurationResponse) *PublicIPAddressSkuResponse {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(PublicIPAddressSkuResponsePtrOutput)
}

type VirtualMachinePublicIPAddressDnsSettingsConfiguration struct {
	DomainNameLabel string `pulumi:"domainNameLabel"`
}





type VirtualMachinePublicIPAddressDnsSettingsConfigurationInput interface {
	pulumi.Input

	ToVirtualMachinePublicIPAddressDnsSettingsConfigurationOutput() VirtualMachinePublicIPAddressDnsSettingsConfigurationOutput
	ToVirtualMachinePublicIPAddressDnsSettingsConfigurationOutputWithContext(context.Context) VirtualMachinePublicIPAddressDnsSettingsConfigurationOutput
}

type VirtualMachinePublicIPAddressDnsSettingsConfigurationArgs struct {
	DomainNameLabel pulumi.StringInput `pulumi:"domainNameLabel"`
}

func (VirtualMachinePublicIPAddressDnsSettingsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachinePublicIPAddressDnsSettingsConfiguration)(nil)).Elem()
}

func (i VirtualMachinePublicIPAddressDnsSettingsConfigurationArgs) ToVirtualMachinePublicIPAddressDnsSettingsConfigurationOutput() VirtualMachinePublicIPAddressDnsSettingsConfigurationOutput {
	return i.ToVirtualMachinePublicIPAddressDnsSettingsConfigurationOutputWithContext(context.Background())
}

func (i VirtualMachinePublicIPAddressDnsSettingsConfigurationArgs) ToVirtualMachinePublicIPAddressDnsSettingsConfigurationOutputWithContext(ctx context.Context) VirtualMachinePublicIPAddressDnsSettingsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachinePublicIPAddressDnsSettingsConfigurationOutput)
}

func (i VirtualMachinePublicIPAddressDnsSettingsConfigurationArgs) ToVirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput() VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput {
	return i.ToVirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutputWithContext(context.Background())
}

func (i VirtualMachinePublicIPAddressDnsSettingsConfigurationArgs) ToVirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachinePublicIPAddressDnsSettingsConfigurationOutput).ToVirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutputWithContext(ctx)
}









type VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrInput interface {
	pulumi.Input

	ToVirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput() VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput
	ToVirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutputWithContext(context.Context) VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput
}

type virtualMachinePublicIPAddressDnsSettingsConfigurationPtrType VirtualMachinePublicIPAddressDnsSettingsConfigurationArgs

func VirtualMachinePublicIPAddressDnsSettingsConfigurationPtr(v *VirtualMachinePublicIPAddressDnsSettingsConfigurationArgs) VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrInput {
	return (*virtualMachinePublicIPAddressDnsSettingsConfigurationPtrType)(v)
}

func (*virtualMachinePublicIPAddressDnsSettingsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachinePublicIPAddressDnsSettingsConfiguration)(nil)).Elem()
}

func (i *virtualMachinePublicIPAddressDnsSettingsConfigurationPtrType) ToVirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput() VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput {
	return i.ToVirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutputWithContext(context.Background())
}

func (i *virtualMachinePublicIPAddressDnsSettingsConfigurationPtrType) ToVirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput)
}

type VirtualMachinePublicIPAddressDnsSettingsConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualMachinePublicIPAddressDnsSettingsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachinePublicIPAddressDnsSettingsConfiguration)(nil)).Elem()
}

func (o VirtualMachinePublicIPAddressDnsSettingsConfigurationOutput) ToVirtualMachinePublicIPAddressDnsSettingsConfigurationOutput() VirtualMachinePublicIPAddressDnsSettingsConfigurationOutput {
	return o
}

func (o VirtualMachinePublicIPAddressDnsSettingsConfigurationOutput) ToVirtualMachinePublicIPAddressDnsSettingsConfigurationOutputWithContext(ctx context.Context) VirtualMachinePublicIPAddressDnsSettingsConfigurationOutput {
	return o
}

func (o VirtualMachinePublicIPAddressDnsSettingsConfigurationOutput) ToVirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput() VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput {
	return o.ToVirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutputWithContext(context.Background())
}

func (o VirtualMachinePublicIPAddressDnsSettingsConfigurationOutput) ToVirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachinePublicIPAddressDnsSettingsConfiguration) *VirtualMachinePublicIPAddressDnsSettingsConfiguration {
		return &v
	}).(VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput)
}

func (o VirtualMachinePublicIPAddressDnsSettingsConfigurationOutput) DomainNameLabel() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressDnsSettingsConfiguration) string { return v.DomainNameLabel }).(pulumi.StringOutput)
}

type VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachinePublicIPAddressDnsSettingsConfiguration)(nil)).Elem()
}

func (o VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput) ToVirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput() VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput {
	return o
}

func (o VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput) ToVirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput {
	return o
}

func (o VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput) Elem() VirtualMachinePublicIPAddressDnsSettingsConfigurationOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressDnsSettingsConfiguration) VirtualMachinePublicIPAddressDnsSettingsConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualMachinePublicIPAddressDnsSettingsConfiguration
		return ret
	}).(VirtualMachinePublicIPAddressDnsSettingsConfigurationOutput)
}

func (o VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressDnsSettingsConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.DomainNameLabel
	}).(pulumi.StringPtrOutput)
}

type VirtualMachinePublicIPAddressDnsSettingsConfigurationResponse struct {
	DomainNameLabel string `pulumi:"domainNameLabel"`
}

type VirtualMachinePublicIPAddressDnsSettingsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachinePublicIPAddressDnsSettingsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachinePublicIPAddressDnsSettingsConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachinePublicIPAddressDnsSettingsConfigurationResponseOutput) ToVirtualMachinePublicIPAddressDnsSettingsConfigurationResponseOutput() VirtualMachinePublicIPAddressDnsSettingsConfigurationResponseOutput {
	return o
}

func (o VirtualMachinePublicIPAddressDnsSettingsConfigurationResponseOutput) ToVirtualMachinePublicIPAddressDnsSettingsConfigurationResponseOutputWithContext(ctx context.Context) VirtualMachinePublicIPAddressDnsSettingsConfigurationResponseOutput {
	return o
}

func (o VirtualMachinePublicIPAddressDnsSettingsConfigurationResponseOutput) DomainNameLabel() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachinePublicIPAddressDnsSettingsConfigurationResponse) string { return v.DomainNameLabel }).(pulumi.StringOutput)
}

type VirtualMachinePublicIPAddressDnsSettingsConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachinePublicIPAddressDnsSettingsConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachinePublicIPAddressDnsSettingsConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachinePublicIPAddressDnsSettingsConfigurationResponsePtrOutput) ToVirtualMachinePublicIPAddressDnsSettingsConfigurationResponsePtrOutput() VirtualMachinePublicIPAddressDnsSettingsConfigurationResponsePtrOutput {
	return o
}

func (o VirtualMachinePublicIPAddressDnsSettingsConfigurationResponsePtrOutput) ToVirtualMachinePublicIPAddressDnsSettingsConfigurationResponsePtrOutputWithContext(ctx context.Context) VirtualMachinePublicIPAddressDnsSettingsConfigurationResponsePtrOutput {
	return o
}

func (o VirtualMachinePublicIPAddressDnsSettingsConfigurationResponsePtrOutput) Elem() VirtualMachinePublicIPAddressDnsSettingsConfigurationResponseOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressDnsSettingsConfigurationResponse) VirtualMachinePublicIPAddressDnsSettingsConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachinePublicIPAddressDnsSettingsConfigurationResponse
		return ret
	}).(VirtualMachinePublicIPAddressDnsSettingsConfigurationResponseOutput)
}

func (o VirtualMachinePublicIPAddressDnsSettingsConfigurationResponsePtrOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachinePublicIPAddressDnsSettingsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DomainNameLabel
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineRunCommandInstanceViewResponse struct {
	EndTime          *string                      `pulumi:"endTime"`
	Error            *string                      `pulumi:"error"`
	ExecutionMessage *string                      `pulumi:"executionMessage"`
	ExecutionState   *string                      `pulumi:"executionState"`
	ExitCode         *int                         `pulumi:"exitCode"`
	Output           *string                      `pulumi:"output"`
	StartTime        *string                      `pulumi:"startTime"`
	Statuses         []InstanceViewStatusResponse `pulumi:"statuses"`
}

type VirtualMachineRunCommandInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineRunCommandInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineRunCommandInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineRunCommandInstanceViewResponseOutput) ToVirtualMachineRunCommandInstanceViewResponseOutput() VirtualMachineRunCommandInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineRunCommandInstanceViewResponseOutput) ToVirtualMachineRunCommandInstanceViewResponseOutputWithContext(ctx context.Context) VirtualMachineRunCommandInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineRunCommandInstanceViewResponseOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineRunCommandInstanceViewResponse) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineRunCommandInstanceViewResponseOutput) Error() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineRunCommandInstanceViewResponse) *string { return v.Error }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineRunCommandInstanceViewResponseOutput) ExecutionMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineRunCommandInstanceViewResponse) *string { return v.ExecutionMessage }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineRunCommandInstanceViewResponseOutput) ExecutionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineRunCommandInstanceViewResponse) *string { return v.ExecutionState }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineRunCommandInstanceViewResponseOutput) ExitCode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineRunCommandInstanceViewResponse) *int { return v.ExitCode }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineRunCommandInstanceViewResponseOutput) Output() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineRunCommandInstanceViewResponse) *string { return v.Output }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineRunCommandInstanceViewResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineRunCommandInstanceViewResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineRunCommandInstanceViewResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineRunCommandInstanceViewResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

type VirtualMachineRunCommandScriptSource struct {
	CommandId *string `pulumi:"commandId"`
	Script    *string `pulumi:"script"`
	ScriptUri *string `pulumi:"scriptUri"`
}





type VirtualMachineRunCommandScriptSourceInput interface {
	pulumi.Input

	ToVirtualMachineRunCommandScriptSourceOutput() VirtualMachineRunCommandScriptSourceOutput
	ToVirtualMachineRunCommandScriptSourceOutputWithContext(context.Context) VirtualMachineRunCommandScriptSourceOutput
}

type VirtualMachineRunCommandScriptSourceArgs struct {
	CommandId pulumi.StringPtrInput `pulumi:"commandId"`
	Script    pulumi.StringPtrInput `pulumi:"script"`
	ScriptUri pulumi.StringPtrInput `pulumi:"scriptUri"`
}

func (VirtualMachineRunCommandScriptSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineRunCommandScriptSource)(nil)).Elem()
}

func (i VirtualMachineRunCommandScriptSourceArgs) ToVirtualMachineRunCommandScriptSourceOutput() VirtualMachineRunCommandScriptSourceOutput {
	return i.ToVirtualMachineRunCommandScriptSourceOutputWithContext(context.Background())
}

func (i VirtualMachineRunCommandScriptSourceArgs) ToVirtualMachineRunCommandScriptSourceOutputWithContext(ctx context.Context) VirtualMachineRunCommandScriptSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineRunCommandScriptSourceOutput)
}

func (i VirtualMachineRunCommandScriptSourceArgs) ToVirtualMachineRunCommandScriptSourcePtrOutput() VirtualMachineRunCommandScriptSourcePtrOutput {
	return i.ToVirtualMachineRunCommandScriptSourcePtrOutputWithContext(context.Background())
}

func (i VirtualMachineRunCommandScriptSourceArgs) ToVirtualMachineRunCommandScriptSourcePtrOutputWithContext(ctx context.Context) VirtualMachineRunCommandScriptSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineRunCommandScriptSourceOutput).ToVirtualMachineRunCommandScriptSourcePtrOutputWithContext(ctx)
}









type VirtualMachineRunCommandScriptSourcePtrInput interface {
	pulumi.Input

	ToVirtualMachineRunCommandScriptSourcePtrOutput() VirtualMachineRunCommandScriptSourcePtrOutput
	ToVirtualMachineRunCommandScriptSourcePtrOutputWithContext(context.Context) VirtualMachineRunCommandScriptSourcePtrOutput
}

type virtualMachineRunCommandScriptSourcePtrType VirtualMachineRunCommandScriptSourceArgs

func VirtualMachineRunCommandScriptSourcePtr(v *VirtualMachineRunCommandScriptSourceArgs) VirtualMachineRunCommandScriptSourcePtrInput {
	return (*virtualMachineRunCommandScriptSourcePtrType)(v)
}

func (*virtualMachineRunCommandScriptSourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineRunCommandScriptSource)(nil)).Elem()
}

func (i *virtualMachineRunCommandScriptSourcePtrType) ToVirtualMachineRunCommandScriptSourcePtrOutput() VirtualMachineRunCommandScriptSourcePtrOutput {
	return i.ToVirtualMachineRunCommandScriptSourcePtrOutputWithContext(context.Background())
}

func (i *virtualMachineRunCommandScriptSourcePtrType) ToVirtualMachineRunCommandScriptSourcePtrOutputWithContext(ctx context.Context) VirtualMachineRunCommandScriptSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineRunCommandScriptSourcePtrOutput)
}

type VirtualMachineRunCommandScriptSourceOutput struct{ *pulumi.OutputState }

func (VirtualMachineRunCommandScriptSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineRunCommandScriptSource)(nil)).Elem()
}

func (o VirtualMachineRunCommandScriptSourceOutput) ToVirtualMachineRunCommandScriptSourceOutput() VirtualMachineRunCommandScriptSourceOutput {
	return o
}

func (o VirtualMachineRunCommandScriptSourceOutput) ToVirtualMachineRunCommandScriptSourceOutputWithContext(ctx context.Context) VirtualMachineRunCommandScriptSourceOutput {
	return o
}

func (o VirtualMachineRunCommandScriptSourceOutput) ToVirtualMachineRunCommandScriptSourcePtrOutput() VirtualMachineRunCommandScriptSourcePtrOutput {
	return o.ToVirtualMachineRunCommandScriptSourcePtrOutputWithContext(context.Background())
}

func (o VirtualMachineRunCommandScriptSourceOutput) ToVirtualMachineRunCommandScriptSourcePtrOutputWithContext(ctx context.Context) VirtualMachineRunCommandScriptSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineRunCommandScriptSource) *VirtualMachineRunCommandScriptSource {
		return &v
	}).(VirtualMachineRunCommandScriptSourcePtrOutput)
}

func (o VirtualMachineRunCommandScriptSourceOutput) CommandId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineRunCommandScriptSource) *string { return v.CommandId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineRunCommandScriptSourceOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineRunCommandScriptSource) *string { return v.Script }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineRunCommandScriptSourceOutput) ScriptUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineRunCommandScriptSource) *string { return v.ScriptUri }).(pulumi.StringPtrOutput)
}

type VirtualMachineRunCommandScriptSourcePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineRunCommandScriptSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineRunCommandScriptSource)(nil)).Elem()
}

func (o VirtualMachineRunCommandScriptSourcePtrOutput) ToVirtualMachineRunCommandScriptSourcePtrOutput() VirtualMachineRunCommandScriptSourcePtrOutput {
	return o
}

func (o VirtualMachineRunCommandScriptSourcePtrOutput) ToVirtualMachineRunCommandScriptSourcePtrOutputWithContext(ctx context.Context) VirtualMachineRunCommandScriptSourcePtrOutput {
	return o
}

func (o VirtualMachineRunCommandScriptSourcePtrOutput) Elem() VirtualMachineRunCommandScriptSourceOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandScriptSource) VirtualMachineRunCommandScriptSource {
		if v != nil {
			return *v
		}
		var ret VirtualMachineRunCommandScriptSource
		return ret
	}).(VirtualMachineRunCommandScriptSourceOutput)
}

func (o VirtualMachineRunCommandScriptSourcePtrOutput) CommandId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandScriptSource) *string {
		if v == nil {
			return nil
		}
		return v.CommandId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineRunCommandScriptSourcePtrOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandScriptSource) *string {
		if v == nil {
			return nil
		}
		return v.Script
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineRunCommandScriptSourcePtrOutput) ScriptUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandScriptSource) *string {
		if v == nil {
			return nil
		}
		return v.ScriptUri
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineRunCommandScriptSourceResponse struct {
	CommandId *string `pulumi:"commandId"`
	Script    *string `pulumi:"script"`
	ScriptUri *string `pulumi:"scriptUri"`
}

type VirtualMachineRunCommandScriptSourceResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineRunCommandScriptSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineRunCommandScriptSourceResponse)(nil)).Elem()
}

func (o VirtualMachineRunCommandScriptSourceResponseOutput) ToVirtualMachineRunCommandScriptSourceResponseOutput() VirtualMachineRunCommandScriptSourceResponseOutput {
	return o
}

func (o VirtualMachineRunCommandScriptSourceResponseOutput) ToVirtualMachineRunCommandScriptSourceResponseOutputWithContext(ctx context.Context) VirtualMachineRunCommandScriptSourceResponseOutput {
	return o
}

func (o VirtualMachineRunCommandScriptSourceResponseOutput) CommandId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineRunCommandScriptSourceResponse) *string { return v.CommandId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineRunCommandScriptSourceResponseOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineRunCommandScriptSourceResponse) *string { return v.Script }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineRunCommandScriptSourceResponseOutput) ScriptUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineRunCommandScriptSourceResponse) *string { return v.ScriptUri }).(pulumi.StringPtrOutput)
}

type VirtualMachineRunCommandScriptSourceResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineRunCommandScriptSourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineRunCommandScriptSourceResponse)(nil)).Elem()
}

func (o VirtualMachineRunCommandScriptSourceResponsePtrOutput) ToVirtualMachineRunCommandScriptSourceResponsePtrOutput() VirtualMachineRunCommandScriptSourceResponsePtrOutput {
	return o
}

func (o VirtualMachineRunCommandScriptSourceResponsePtrOutput) ToVirtualMachineRunCommandScriptSourceResponsePtrOutputWithContext(ctx context.Context) VirtualMachineRunCommandScriptSourceResponsePtrOutput {
	return o
}

func (o VirtualMachineRunCommandScriptSourceResponsePtrOutput) Elem() VirtualMachineRunCommandScriptSourceResponseOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandScriptSourceResponse) VirtualMachineRunCommandScriptSourceResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineRunCommandScriptSourceResponse
		return ret
	}).(VirtualMachineRunCommandScriptSourceResponseOutput)
}

func (o VirtualMachineRunCommandScriptSourceResponsePtrOutput) CommandId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandScriptSourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.CommandId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineRunCommandScriptSourceResponsePtrOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandScriptSourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Script
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineRunCommandScriptSourceResponsePtrOutput) ScriptUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandScriptSourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScriptUri
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetDataDisk struct {
	Caching                 *CachingTypes                                `pulumi:"caching"`
	CreateOption            string                                       `pulumi:"createOption"`
	DeleteOption            *string                                      `pulumi:"deleteOption"`
	DiskIOPSReadWrite       *float64                                     `pulumi:"diskIOPSReadWrite"`
	DiskMBpsReadWrite       *float64                                     `pulumi:"diskMBpsReadWrite"`
	DiskSizeGB              *int                                         `pulumi:"diskSizeGB"`
	Lun                     int                                          `pulumi:"lun"`
	ManagedDisk             *VirtualMachineScaleSetManagedDiskParameters `pulumi:"managedDisk"`
	Name                    *string                                      `pulumi:"name"`
	WriteAcceleratorEnabled *bool                                        `pulumi:"writeAcceleratorEnabled"`
}





type VirtualMachineScaleSetDataDiskInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetDataDiskOutput() VirtualMachineScaleSetDataDiskOutput
	ToVirtualMachineScaleSetDataDiskOutputWithContext(context.Context) VirtualMachineScaleSetDataDiskOutput
}

type VirtualMachineScaleSetDataDiskArgs struct {
	Caching                 CachingTypesPtrInput                                `pulumi:"caching"`
	CreateOption            pulumi.StringInput                                  `pulumi:"createOption"`
	DeleteOption            pulumi.StringPtrInput                               `pulumi:"deleteOption"`
	DiskIOPSReadWrite       pulumi.Float64PtrInput                              `pulumi:"diskIOPSReadWrite"`
	DiskMBpsReadWrite       pulumi.Float64PtrInput                              `pulumi:"diskMBpsReadWrite"`
	DiskSizeGB              pulumi.IntPtrInput                                  `pulumi:"diskSizeGB"`
	Lun                     pulumi.IntInput                                     `pulumi:"lun"`
	ManagedDisk             VirtualMachineScaleSetManagedDiskParametersPtrInput `pulumi:"managedDisk"`
	Name                    pulumi.StringPtrInput                               `pulumi:"name"`
	WriteAcceleratorEnabled pulumi.BoolPtrInput                                 `pulumi:"writeAcceleratorEnabled"`
}

func (VirtualMachineScaleSetDataDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetDataDisk)(nil)).Elem()
}

func (i VirtualMachineScaleSetDataDiskArgs) ToVirtualMachineScaleSetDataDiskOutput() VirtualMachineScaleSetDataDiskOutput {
	return i.ToVirtualMachineScaleSetDataDiskOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetDataDiskArgs) ToVirtualMachineScaleSetDataDiskOutputWithContext(ctx context.Context) VirtualMachineScaleSetDataDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetDataDiskOutput)
}





type VirtualMachineScaleSetDataDiskArrayInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetDataDiskArrayOutput() VirtualMachineScaleSetDataDiskArrayOutput
	ToVirtualMachineScaleSetDataDiskArrayOutputWithContext(context.Context) VirtualMachineScaleSetDataDiskArrayOutput
}

type VirtualMachineScaleSetDataDiskArray []VirtualMachineScaleSetDataDiskInput

func (VirtualMachineScaleSetDataDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetDataDisk)(nil)).Elem()
}

func (i VirtualMachineScaleSetDataDiskArray) ToVirtualMachineScaleSetDataDiskArrayOutput() VirtualMachineScaleSetDataDiskArrayOutput {
	return i.ToVirtualMachineScaleSetDataDiskArrayOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetDataDiskArray) ToVirtualMachineScaleSetDataDiskArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetDataDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetDataDiskArrayOutput)
}

type VirtualMachineScaleSetDataDiskOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetDataDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetDataDisk)(nil)).Elem()
}

func (o VirtualMachineScaleSetDataDiskOutput) ToVirtualMachineScaleSetDataDiskOutput() VirtualMachineScaleSetDataDiskOutput {
	return o
}

func (o VirtualMachineScaleSetDataDiskOutput) ToVirtualMachineScaleSetDataDiskOutputWithContext(ctx context.Context) VirtualMachineScaleSetDataDiskOutput {
	return o
}

func (o VirtualMachineScaleSetDataDiskOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDisk) *CachingTypes { return v.Caching }).(CachingTypesPtrOutput)
}

func (o VirtualMachineScaleSetDataDiskOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDisk) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetDataDiskOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDisk) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetDataDiskOutput) DiskIOPSReadWrite() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDisk) *float64 { return v.DiskIOPSReadWrite }).(pulumi.Float64PtrOutput)
}

func (o VirtualMachineScaleSetDataDiskOutput) DiskMBpsReadWrite() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDisk) *float64 { return v.DiskMBpsReadWrite }).(pulumi.Float64PtrOutput)
}

func (o VirtualMachineScaleSetDataDiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDisk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetDataDiskOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDisk) int { return v.Lun }).(pulumi.IntOutput)
}

func (o VirtualMachineScaleSetDataDiskOutput) ManagedDisk() VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDisk) *VirtualMachineScaleSetManagedDiskParameters {
		return v.ManagedDisk
	}).(VirtualMachineScaleSetManagedDiskParametersPtrOutput)
}

func (o VirtualMachineScaleSetDataDiskOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDisk) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetDataDiskOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDisk) *bool { return v.WriteAcceleratorEnabled }).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetDataDiskArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetDataDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetDataDisk)(nil)).Elem()
}

func (o VirtualMachineScaleSetDataDiskArrayOutput) ToVirtualMachineScaleSetDataDiskArrayOutput() VirtualMachineScaleSetDataDiskArrayOutput {
	return o
}

func (o VirtualMachineScaleSetDataDiskArrayOutput) ToVirtualMachineScaleSetDataDiskArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetDataDiskArrayOutput {
	return o
}

func (o VirtualMachineScaleSetDataDiskArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetDataDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetDataDisk {
		return vs[0].([]VirtualMachineScaleSetDataDisk)[vs[1].(int)]
	}).(VirtualMachineScaleSetDataDiskOutput)
}

type VirtualMachineScaleSetDataDiskResponse struct {
	Caching                 *string                                              `pulumi:"caching"`
	CreateOption            string                                               `pulumi:"createOption"`
	DeleteOption            *string                                              `pulumi:"deleteOption"`
	DiskIOPSReadWrite       *float64                                             `pulumi:"diskIOPSReadWrite"`
	DiskMBpsReadWrite       *float64                                             `pulumi:"diskMBpsReadWrite"`
	DiskSizeGB              *int                                                 `pulumi:"diskSizeGB"`
	Lun                     int                                                  `pulumi:"lun"`
	ManagedDisk             *VirtualMachineScaleSetManagedDiskParametersResponse `pulumi:"managedDisk"`
	Name                    *string                                              `pulumi:"name"`
	WriteAcceleratorEnabled *bool                                                `pulumi:"writeAcceleratorEnabled"`
}

type VirtualMachineScaleSetDataDiskResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetDataDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetDataDiskResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) ToVirtualMachineScaleSetDataDiskResponseOutput() VirtualMachineScaleSetDataDiskResponseOutput {
	return o
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) ToVirtualMachineScaleSetDataDiskResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetDataDiskResponseOutput {
	return o
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDiskResponse) *string { return v.Caching }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDiskResponse) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDiskResponse) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) DiskIOPSReadWrite() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDiskResponse) *float64 { return v.DiskIOPSReadWrite }).(pulumi.Float64PtrOutput)
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) DiskMBpsReadWrite() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDiskResponse) *float64 { return v.DiskMBpsReadWrite }).(pulumi.Float64PtrOutput)
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDiskResponse) int { return v.Lun }).(pulumi.IntOutput)
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) ManagedDisk() VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDiskResponse) *VirtualMachineScaleSetManagedDiskParametersResponse {
		return v.ManagedDisk
	}).(VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput)
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDiskResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetDataDiskResponseOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetDataDiskResponse) *bool { return v.WriteAcceleratorEnabled }).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetDataDiskResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetDataDiskResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetDataDiskResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetDataDiskResponseArrayOutput) ToVirtualMachineScaleSetDataDiskResponseArrayOutput() VirtualMachineScaleSetDataDiskResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetDataDiskResponseArrayOutput) ToVirtualMachineScaleSetDataDiskResponseArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetDataDiskResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetDataDiskResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetDataDiskResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetDataDiskResponse {
		return vs[0].([]VirtualMachineScaleSetDataDiskResponse)[vs[1].(int)]
	}).(VirtualMachineScaleSetDataDiskResponseOutput)
}

type VirtualMachineScaleSetExtensionType struct {
	AutoUpgradeMinorVersion       *bool                    `pulumi:"autoUpgradeMinorVersion"`
	EnableAutomaticUpgrade        *bool                    `pulumi:"enableAutomaticUpgrade"`
	ForceUpdateTag                *string                  `pulumi:"forceUpdateTag"`
	Name                          *string                  `pulumi:"name"`
	ProtectedSettings             interface{}              `pulumi:"protectedSettings"`
	ProtectedSettingsFromKeyVault *KeyVaultSecretReference `pulumi:"protectedSettingsFromKeyVault"`
	ProvisionAfterExtensions      []string                 `pulumi:"provisionAfterExtensions"`
	Publisher                     *string                  `pulumi:"publisher"`
	Settings                      interface{}              `pulumi:"settings"`
	SuppressFailures              *bool                    `pulumi:"suppressFailures"`
	Type                          *string                  `pulumi:"type"`
	TypeHandlerVersion            *string                  `pulumi:"typeHandlerVersion"`
}





type VirtualMachineScaleSetExtensionTypeInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetExtensionTypeOutput() VirtualMachineScaleSetExtensionTypeOutput
	ToVirtualMachineScaleSetExtensionTypeOutputWithContext(context.Context) VirtualMachineScaleSetExtensionTypeOutput
}

type VirtualMachineScaleSetExtensionTypeArgs struct {
	AutoUpgradeMinorVersion       pulumi.BoolPtrInput             `pulumi:"autoUpgradeMinorVersion"`
	EnableAutomaticUpgrade        pulumi.BoolPtrInput             `pulumi:"enableAutomaticUpgrade"`
	ForceUpdateTag                pulumi.StringPtrInput           `pulumi:"forceUpdateTag"`
	Name                          pulumi.StringPtrInput           `pulumi:"name"`
	ProtectedSettings             pulumi.Input                    `pulumi:"protectedSettings"`
	ProtectedSettingsFromKeyVault KeyVaultSecretReferencePtrInput `pulumi:"protectedSettingsFromKeyVault"`
	ProvisionAfterExtensions      pulumi.StringArrayInput         `pulumi:"provisionAfterExtensions"`
	Publisher                     pulumi.StringPtrInput           `pulumi:"publisher"`
	Settings                      pulumi.Input                    `pulumi:"settings"`
	SuppressFailures              pulumi.BoolPtrInput             `pulumi:"suppressFailures"`
	Type                          pulumi.StringPtrInput           `pulumi:"type"`
	TypeHandlerVersion            pulumi.StringPtrInput           `pulumi:"typeHandlerVersion"`
}

func (VirtualMachineScaleSetExtensionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtensionType)(nil)).Elem()
}

func (i VirtualMachineScaleSetExtensionTypeArgs) ToVirtualMachineScaleSetExtensionTypeOutput() VirtualMachineScaleSetExtensionTypeOutput {
	return i.ToVirtualMachineScaleSetExtensionTypeOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetExtensionTypeArgs) ToVirtualMachineScaleSetExtensionTypeOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionTypeOutput)
}





type VirtualMachineScaleSetExtensionTypeArrayInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetExtensionTypeArrayOutput() VirtualMachineScaleSetExtensionTypeArrayOutput
	ToVirtualMachineScaleSetExtensionTypeArrayOutputWithContext(context.Context) VirtualMachineScaleSetExtensionTypeArrayOutput
}

type VirtualMachineScaleSetExtensionTypeArray []VirtualMachineScaleSetExtensionTypeInput

func (VirtualMachineScaleSetExtensionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetExtensionType)(nil)).Elem()
}

func (i VirtualMachineScaleSetExtensionTypeArray) ToVirtualMachineScaleSetExtensionTypeArrayOutput() VirtualMachineScaleSetExtensionTypeArrayOutput {
	return i.ToVirtualMachineScaleSetExtensionTypeArrayOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetExtensionTypeArray) ToVirtualMachineScaleSetExtensionTypeArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionTypeArrayOutput)
}

type VirtualMachineScaleSetExtensionTypeOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtensionType)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionTypeOutput) ToVirtualMachineScaleSetExtensionTypeOutput() VirtualMachineScaleSetExtensionTypeOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionTypeOutput) ToVirtualMachineScaleSetExtensionTypeOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionTypeOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionTypeOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetExtensionTypeOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) *bool { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetExtensionTypeOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionTypeOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o VirtualMachineScaleSetExtensionTypeOutput) ProtectedSettingsFromKeyVault() KeyVaultSecretReferencePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) *KeyVaultSecretReference {
		return v.ProtectedSettingsFromKeyVault
	}).(KeyVaultSecretReferencePtrOutput)
}

func (o VirtualMachineScaleSetExtensionTypeOutput) ProvisionAfterExtensions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) []string { return v.ProvisionAfterExtensions }).(pulumi.StringArrayOutput)
}

func (o VirtualMachineScaleSetExtensionTypeOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionTypeOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o VirtualMachineScaleSetExtensionTypeOutput) SuppressFailures() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) *bool { return v.SuppressFailures }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetExtensionTypeOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionTypeOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionType) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetExtensionTypeArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetExtensionType)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionTypeArrayOutput) ToVirtualMachineScaleSetExtensionTypeArrayOutput() VirtualMachineScaleSetExtensionTypeArrayOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionTypeArrayOutput) ToVirtualMachineScaleSetExtensionTypeArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionTypeArrayOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionTypeArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetExtensionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetExtensionType {
		return vs[0].([]VirtualMachineScaleSetExtensionType)[vs[1].(int)]
	}).(VirtualMachineScaleSetExtensionTypeOutput)
}

type VirtualMachineScaleSetExtensionProfile struct {
	Extensions           []VirtualMachineScaleSetExtensionType `pulumi:"extensions"`
	ExtensionsTimeBudget *string                               `pulumi:"extensionsTimeBudget"`
}





type VirtualMachineScaleSetExtensionProfileInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetExtensionProfileOutput() VirtualMachineScaleSetExtensionProfileOutput
	ToVirtualMachineScaleSetExtensionProfileOutputWithContext(context.Context) VirtualMachineScaleSetExtensionProfileOutput
}

type VirtualMachineScaleSetExtensionProfileArgs struct {
	Extensions           VirtualMachineScaleSetExtensionTypeArrayInput `pulumi:"extensions"`
	ExtensionsTimeBudget pulumi.StringPtrInput                         `pulumi:"extensionsTimeBudget"`
}

func (VirtualMachineScaleSetExtensionProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtensionProfile)(nil)).Elem()
}

func (i VirtualMachineScaleSetExtensionProfileArgs) ToVirtualMachineScaleSetExtensionProfileOutput() VirtualMachineScaleSetExtensionProfileOutput {
	return i.ToVirtualMachineScaleSetExtensionProfileOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetExtensionProfileArgs) ToVirtualMachineScaleSetExtensionProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionProfileOutput)
}

func (i VirtualMachineScaleSetExtensionProfileArgs) ToVirtualMachineScaleSetExtensionProfilePtrOutput() VirtualMachineScaleSetExtensionProfilePtrOutput {
	return i.ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetExtensionProfileArgs) ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionProfileOutput).ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetExtensionProfilePtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetExtensionProfilePtrOutput() VirtualMachineScaleSetExtensionProfilePtrOutput
	ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(context.Context) VirtualMachineScaleSetExtensionProfilePtrOutput
}

type virtualMachineScaleSetExtensionProfilePtrType VirtualMachineScaleSetExtensionProfileArgs

func VirtualMachineScaleSetExtensionProfilePtr(v *VirtualMachineScaleSetExtensionProfileArgs) VirtualMachineScaleSetExtensionProfilePtrInput {
	return (*virtualMachineScaleSetExtensionProfilePtrType)(v)
}

func (*virtualMachineScaleSetExtensionProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetExtensionProfile)(nil)).Elem()
}

func (i *virtualMachineScaleSetExtensionProfilePtrType) ToVirtualMachineScaleSetExtensionProfilePtrOutput() VirtualMachineScaleSetExtensionProfilePtrOutput {
	return i.ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetExtensionProfilePtrType) ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionProfilePtrOutput)
}

type VirtualMachineScaleSetExtensionProfileOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtensionProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionProfileOutput) ToVirtualMachineScaleSetExtensionProfileOutput() VirtualMachineScaleSetExtensionProfileOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfileOutput) ToVirtualMachineScaleSetExtensionProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfileOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfileOutput) ToVirtualMachineScaleSetExtensionProfilePtrOutput() VirtualMachineScaleSetExtensionProfilePtrOutput {
	return o.ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetExtensionProfileOutput) ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetExtensionProfile) *VirtualMachineScaleSetExtensionProfile {
		return &v
	}).(VirtualMachineScaleSetExtensionProfilePtrOutput)
}

func (o VirtualMachineScaleSetExtensionProfileOutput) Extensions() VirtualMachineScaleSetExtensionTypeArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionProfile) []VirtualMachineScaleSetExtensionType {
		return v.Extensions
	}).(VirtualMachineScaleSetExtensionTypeArrayOutput)
}

func (o VirtualMachineScaleSetExtensionProfileOutput) ExtensionsTimeBudget() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionProfile) *string { return v.ExtensionsTimeBudget }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetExtensionProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetExtensionProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionProfilePtrOutput) ToVirtualMachineScaleSetExtensionProfilePtrOutput() VirtualMachineScaleSetExtensionProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfilePtrOutput) ToVirtualMachineScaleSetExtensionProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfilePtrOutput) Elem() VirtualMachineScaleSetExtensionProfileOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtensionProfile) VirtualMachineScaleSetExtensionProfile {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetExtensionProfile
		return ret
	}).(VirtualMachineScaleSetExtensionProfileOutput)
}

func (o VirtualMachineScaleSetExtensionProfilePtrOutput) Extensions() VirtualMachineScaleSetExtensionTypeArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtensionProfile) []VirtualMachineScaleSetExtensionType {
		if v == nil {
			return nil
		}
		return v.Extensions
	}).(VirtualMachineScaleSetExtensionTypeArrayOutput)
}

func (o VirtualMachineScaleSetExtensionProfilePtrOutput) ExtensionsTimeBudget() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtensionProfile) *string {
		if v == nil {
			return nil
		}
		return v.ExtensionsTimeBudget
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetExtensionProfileResponse struct {
	Extensions           []VirtualMachineScaleSetExtensionResponse `pulumi:"extensions"`
	ExtensionsTimeBudget *string                                   `pulumi:"extensionsTimeBudget"`
}

type VirtualMachineScaleSetExtensionProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtensionProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionProfileResponseOutput) ToVirtualMachineScaleSetExtensionProfileResponseOutput() VirtualMachineScaleSetExtensionProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfileResponseOutput) ToVirtualMachineScaleSetExtensionProfileResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfileResponseOutput) Extensions() VirtualMachineScaleSetExtensionResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionProfileResponse) []VirtualMachineScaleSetExtensionResponse {
		return v.Extensions
	}).(VirtualMachineScaleSetExtensionResponseArrayOutput)
}

func (o VirtualMachineScaleSetExtensionProfileResponseOutput) ExtensionsTimeBudget() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionProfileResponse) *string { return v.ExtensionsTimeBudget }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetExtensionProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetExtensionProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionProfileResponsePtrOutput) ToVirtualMachineScaleSetExtensionProfileResponsePtrOutput() VirtualMachineScaleSetExtensionProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfileResponsePtrOutput) ToVirtualMachineScaleSetExtensionProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionProfileResponsePtrOutput) Elem() VirtualMachineScaleSetExtensionProfileResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtensionProfileResponse) VirtualMachineScaleSetExtensionProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetExtensionProfileResponse
		return ret
	}).(VirtualMachineScaleSetExtensionProfileResponseOutput)
}

func (o VirtualMachineScaleSetExtensionProfileResponsePtrOutput) Extensions() VirtualMachineScaleSetExtensionResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtensionProfileResponse) []VirtualMachineScaleSetExtensionResponse {
		if v == nil {
			return nil
		}
		return v.Extensions
	}).(VirtualMachineScaleSetExtensionResponseArrayOutput)
}

func (o VirtualMachineScaleSetExtensionProfileResponsePtrOutput) ExtensionsTimeBudget() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtensionProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExtensionsTimeBudget
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetExtensionResponse struct {
	AutoUpgradeMinorVersion       *bool                            `pulumi:"autoUpgradeMinorVersion"`
	EnableAutomaticUpgrade        *bool                            `pulumi:"enableAutomaticUpgrade"`
	ForceUpdateTag                *string                          `pulumi:"forceUpdateTag"`
	Id                            string                           `pulumi:"id"`
	Name                          *string                          `pulumi:"name"`
	ProtectedSettings             interface{}                      `pulumi:"protectedSettings"`
	ProtectedSettingsFromKeyVault *KeyVaultSecretReferenceResponse `pulumi:"protectedSettingsFromKeyVault"`
	ProvisionAfterExtensions      []string                         `pulumi:"provisionAfterExtensions"`
	ProvisioningState             string                           `pulumi:"provisioningState"`
	Publisher                     *string                          `pulumi:"publisher"`
	Settings                      interface{}                      `pulumi:"settings"`
	SuppressFailures              *bool                            `pulumi:"suppressFailures"`
	Type                          string                           `pulumi:"type"`
	TypeHandlerVersion            *string                          `pulumi:"typeHandlerVersion"`
}

type VirtualMachineScaleSetExtensionResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetExtensionResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionResponseOutput) ToVirtualMachineScaleSetExtensionResponseOutput() VirtualMachineScaleSetExtensionResponseOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionResponseOutput) ToVirtualMachineScaleSetExtensionResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionResponseOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionResponseOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *bool { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) interface{} { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) ProtectedSettingsFromKeyVault() KeyVaultSecretReferenceResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *KeyVaultSecretReferenceResponse {
		return v.ProtectedSettingsFromKeyVault
	}).(KeyVaultSecretReferenceResponsePtrOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) ProvisionAfterExtensions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) []string { return v.ProvisionAfterExtensions }).(pulumi.StringArrayOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) interface{} { return v.Settings }).(pulumi.AnyOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) SuppressFailures() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *bool { return v.SuppressFailures }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetExtensionResponseOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetExtensionResponse) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetExtensionResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetExtensionResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionResponseArrayOutput) ToVirtualMachineScaleSetExtensionResponseArrayOutput() VirtualMachineScaleSetExtensionResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionResponseArrayOutput) ToVirtualMachineScaleSetExtensionResponseArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetExtensionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetExtensionResponse {
		return vs[0].([]VirtualMachineScaleSetExtensionResponse)[vs[1].(int)]
	}).(VirtualMachineScaleSetExtensionResponseOutput)
}

type VirtualMachineScaleSetHardwareProfile struct {
	VmSizeProperties *VMSizeProperties `pulumi:"vmSizeProperties"`
}





type VirtualMachineScaleSetHardwareProfileInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetHardwareProfileOutput() VirtualMachineScaleSetHardwareProfileOutput
	ToVirtualMachineScaleSetHardwareProfileOutputWithContext(context.Context) VirtualMachineScaleSetHardwareProfileOutput
}

type VirtualMachineScaleSetHardwareProfileArgs struct {
	VmSizeProperties VMSizePropertiesPtrInput `pulumi:"vmSizeProperties"`
}

func (VirtualMachineScaleSetHardwareProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetHardwareProfile)(nil)).Elem()
}

func (i VirtualMachineScaleSetHardwareProfileArgs) ToVirtualMachineScaleSetHardwareProfileOutput() VirtualMachineScaleSetHardwareProfileOutput {
	return i.ToVirtualMachineScaleSetHardwareProfileOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetHardwareProfileArgs) ToVirtualMachineScaleSetHardwareProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetHardwareProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetHardwareProfileOutput)
}

func (i VirtualMachineScaleSetHardwareProfileArgs) ToVirtualMachineScaleSetHardwareProfilePtrOutput() VirtualMachineScaleSetHardwareProfilePtrOutput {
	return i.ToVirtualMachineScaleSetHardwareProfilePtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetHardwareProfileArgs) ToVirtualMachineScaleSetHardwareProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetHardwareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetHardwareProfileOutput).ToVirtualMachineScaleSetHardwareProfilePtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetHardwareProfilePtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetHardwareProfilePtrOutput() VirtualMachineScaleSetHardwareProfilePtrOutput
	ToVirtualMachineScaleSetHardwareProfilePtrOutputWithContext(context.Context) VirtualMachineScaleSetHardwareProfilePtrOutput
}

type virtualMachineScaleSetHardwareProfilePtrType VirtualMachineScaleSetHardwareProfileArgs

func VirtualMachineScaleSetHardwareProfilePtr(v *VirtualMachineScaleSetHardwareProfileArgs) VirtualMachineScaleSetHardwareProfilePtrInput {
	return (*virtualMachineScaleSetHardwareProfilePtrType)(v)
}

func (*virtualMachineScaleSetHardwareProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetHardwareProfile)(nil)).Elem()
}

func (i *virtualMachineScaleSetHardwareProfilePtrType) ToVirtualMachineScaleSetHardwareProfilePtrOutput() VirtualMachineScaleSetHardwareProfilePtrOutput {
	return i.ToVirtualMachineScaleSetHardwareProfilePtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetHardwareProfilePtrType) ToVirtualMachineScaleSetHardwareProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetHardwareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetHardwareProfilePtrOutput)
}

type VirtualMachineScaleSetHardwareProfileOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetHardwareProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetHardwareProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetHardwareProfileOutput) ToVirtualMachineScaleSetHardwareProfileOutput() VirtualMachineScaleSetHardwareProfileOutput {
	return o
}

func (o VirtualMachineScaleSetHardwareProfileOutput) ToVirtualMachineScaleSetHardwareProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetHardwareProfileOutput {
	return o
}

func (o VirtualMachineScaleSetHardwareProfileOutput) ToVirtualMachineScaleSetHardwareProfilePtrOutput() VirtualMachineScaleSetHardwareProfilePtrOutput {
	return o.ToVirtualMachineScaleSetHardwareProfilePtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetHardwareProfileOutput) ToVirtualMachineScaleSetHardwareProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetHardwareProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetHardwareProfile) *VirtualMachineScaleSetHardwareProfile {
		return &v
	}).(VirtualMachineScaleSetHardwareProfilePtrOutput)
}

func (o VirtualMachineScaleSetHardwareProfileOutput) VmSizeProperties() VMSizePropertiesPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetHardwareProfile) *VMSizeProperties { return v.VmSizeProperties }).(VMSizePropertiesPtrOutput)
}

type VirtualMachineScaleSetHardwareProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetHardwareProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetHardwareProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetHardwareProfilePtrOutput) ToVirtualMachineScaleSetHardwareProfilePtrOutput() VirtualMachineScaleSetHardwareProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetHardwareProfilePtrOutput) ToVirtualMachineScaleSetHardwareProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetHardwareProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetHardwareProfilePtrOutput) Elem() VirtualMachineScaleSetHardwareProfileOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetHardwareProfile) VirtualMachineScaleSetHardwareProfile {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetHardwareProfile
		return ret
	}).(VirtualMachineScaleSetHardwareProfileOutput)
}

func (o VirtualMachineScaleSetHardwareProfilePtrOutput) VmSizeProperties() VMSizePropertiesPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetHardwareProfile) *VMSizeProperties {
		if v == nil {
			return nil
		}
		return v.VmSizeProperties
	}).(VMSizePropertiesPtrOutput)
}

type VirtualMachineScaleSetHardwareProfileResponse struct {
	VmSizeProperties *VMSizePropertiesResponse `pulumi:"vmSizeProperties"`
}

type VirtualMachineScaleSetHardwareProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetHardwareProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetHardwareProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetHardwareProfileResponseOutput) ToVirtualMachineScaleSetHardwareProfileResponseOutput() VirtualMachineScaleSetHardwareProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetHardwareProfileResponseOutput) ToVirtualMachineScaleSetHardwareProfileResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetHardwareProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetHardwareProfileResponseOutput) VmSizeProperties() VMSizePropertiesResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetHardwareProfileResponse) *VMSizePropertiesResponse {
		return v.VmSizeProperties
	}).(VMSizePropertiesResponsePtrOutput)
}

type VirtualMachineScaleSetHardwareProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetHardwareProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetHardwareProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetHardwareProfileResponsePtrOutput) ToVirtualMachineScaleSetHardwareProfileResponsePtrOutput() VirtualMachineScaleSetHardwareProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetHardwareProfileResponsePtrOutput) ToVirtualMachineScaleSetHardwareProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetHardwareProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetHardwareProfileResponsePtrOutput) Elem() VirtualMachineScaleSetHardwareProfileResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetHardwareProfileResponse) VirtualMachineScaleSetHardwareProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetHardwareProfileResponse
		return ret
	}).(VirtualMachineScaleSetHardwareProfileResponseOutput)
}

func (o VirtualMachineScaleSetHardwareProfileResponsePtrOutput) VmSizeProperties() VMSizePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetHardwareProfileResponse) *VMSizePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.VmSizeProperties
	}).(VMSizePropertiesResponsePtrOutput)
}

type VirtualMachineScaleSetIPConfiguration struct {
	ApplicationGatewayBackendAddressPools []SubResource                                       `pulumi:"applicationGatewayBackendAddressPools"`
	ApplicationSecurityGroups             []SubResource                                       `pulumi:"applicationSecurityGroups"`
	Id                                    *string                                             `pulumi:"id"`
	LoadBalancerBackendAddressPools       []SubResource                                       `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatPools           []SubResource                                       `pulumi:"loadBalancerInboundNatPools"`
	Name                                  string                                              `pulumi:"name"`
	Primary                               *bool                                               `pulumi:"primary"`
	PrivateIPAddressVersion               *string                                             `pulumi:"privateIPAddressVersion"`
	PublicIPAddressConfiguration          *VirtualMachineScaleSetPublicIPAddressConfiguration `pulumi:"publicIPAddressConfiguration"`
	Subnet                                *ApiEntityReference                                 `pulumi:"subnet"`
}





type VirtualMachineScaleSetIPConfigurationInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetIPConfigurationOutput() VirtualMachineScaleSetIPConfigurationOutput
	ToVirtualMachineScaleSetIPConfigurationOutputWithContext(context.Context) VirtualMachineScaleSetIPConfigurationOutput
}

type VirtualMachineScaleSetIPConfigurationArgs struct {
	ApplicationGatewayBackendAddressPools SubResourceArrayInput                                      `pulumi:"applicationGatewayBackendAddressPools"`
	ApplicationSecurityGroups             SubResourceArrayInput                                      `pulumi:"applicationSecurityGroups"`
	Id                                    pulumi.StringPtrInput                                      `pulumi:"id"`
	LoadBalancerBackendAddressPools       SubResourceArrayInput                                      `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatPools           SubResourceArrayInput                                      `pulumi:"loadBalancerInboundNatPools"`
	Name                                  pulumi.StringInput                                         `pulumi:"name"`
	Primary                               pulumi.BoolPtrInput                                        `pulumi:"primary"`
	PrivateIPAddressVersion               pulumi.StringPtrInput                                      `pulumi:"privateIPAddressVersion"`
	PublicIPAddressConfiguration          VirtualMachineScaleSetPublicIPAddressConfigurationPtrInput `pulumi:"publicIPAddressConfiguration"`
	Subnet                                ApiEntityReferencePtrInput                                 `pulumi:"subnet"`
}

func (VirtualMachineScaleSetIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIPConfiguration)(nil)).Elem()
}

func (i VirtualMachineScaleSetIPConfigurationArgs) ToVirtualMachineScaleSetIPConfigurationOutput() VirtualMachineScaleSetIPConfigurationOutput {
	return i.ToVirtualMachineScaleSetIPConfigurationOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetIPConfigurationArgs) ToVirtualMachineScaleSetIPConfigurationOutputWithContext(ctx context.Context) VirtualMachineScaleSetIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetIPConfigurationOutput)
}





type VirtualMachineScaleSetIPConfigurationArrayInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetIPConfigurationArrayOutput() VirtualMachineScaleSetIPConfigurationArrayOutput
	ToVirtualMachineScaleSetIPConfigurationArrayOutputWithContext(context.Context) VirtualMachineScaleSetIPConfigurationArrayOutput
}

type VirtualMachineScaleSetIPConfigurationArray []VirtualMachineScaleSetIPConfigurationInput

func (VirtualMachineScaleSetIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetIPConfiguration)(nil)).Elem()
}

func (i VirtualMachineScaleSetIPConfigurationArray) ToVirtualMachineScaleSetIPConfigurationArrayOutput() VirtualMachineScaleSetIPConfigurationArrayOutput {
	return i.ToVirtualMachineScaleSetIPConfigurationArrayOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetIPConfigurationArray) ToVirtualMachineScaleSetIPConfigurationArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetIPConfigurationArrayOutput)
}

type VirtualMachineScaleSetIPConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIPConfiguration)(nil)).Elem()
}

func (o VirtualMachineScaleSetIPConfigurationOutput) ToVirtualMachineScaleSetIPConfigurationOutput() VirtualMachineScaleSetIPConfigurationOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationOutput) ToVirtualMachineScaleSetIPConfigurationOutputWithContext(ctx context.Context) VirtualMachineScaleSetIPConfigurationOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationOutput) ApplicationGatewayBackendAddressPools() SubResourceArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) []SubResource {
		return v.ApplicationGatewayBackendAddressPools
	}).(SubResourceArrayOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) ApplicationSecurityGroups() SubResourceArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) []SubResource { return v.ApplicationSecurityGroups }).(SubResourceArrayOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) LoadBalancerBackendAddressPools() SubResourceArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) []SubResource { return v.LoadBalancerBackendAddressPools }).(SubResourceArrayOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) LoadBalancerInboundNatPools() SubResourceArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) []SubResource { return v.LoadBalancerInboundNatPools }).(SubResourceArrayOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) PrivateIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) *string { return v.PrivateIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) PublicIPAddressConfiguration() VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) *VirtualMachineScaleSetPublicIPAddressConfiguration {
		return v.PublicIPAddressConfiguration
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput)
}

func (o VirtualMachineScaleSetIPConfigurationOutput) Subnet() ApiEntityReferencePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfiguration) *ApiEntityReference { return v.Subnet }).(ApiEntityReferencePtrOutput)
}

type VirtualMachineScaleSetIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetIPConfiguration)(nil)).Elem()
}

func (o VirtualMachineScaleSetIPConfigurationArrayOutput) ToVirtualMachineScaleSetIPConfigurationArrayOutput() VirtualMachineScaleSetIPConfigurationArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationArrayOutput) ToVirtualMachineScaleSetIPConfigurationArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetIPConfigurationArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetIPConfiguration {
		return vs[0].([]VirtualMachineScaleSetIPConfiguration)[vs[1].(int)]
	}).(VirtualMachineScaleSetIPConfigurationOutput)
}

type VirtualMachineScaleSetIPConfigurationResponse struct {
	ApplicationGatewayBackendAddressPools []SubResourceResponse                                       `pulumi:"applicationGatewayBackendAddressPools"`
	ApplicationSecurityGroups             []SubResourceResponse                                       `pulumi:"applicationSecurityGroups"`
	Id                                    *string                                                     `pulumi:"id"`
	LoadBalancerBackendAddressPools       []SubResourceResponse                                       `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatPools           []SubResourceResponse                                       `pulumi:"loadBalancerInboundNatPools"`
	Name                                  string                                                      `pulumi:"name"`
	Primary                               *bool                                                       `pulumi:"primary"`
	PrivateIPAddressVersion               *string                                                     `pulumi:"privateIPAddressVersion"`
	PublicIPAddressConfiguration          *VirtualMachineScaleSetPublicIPAddressConfigurationResponse `pulumi:"publicIPAddressConfiguration"`
	Subnet                                *ApiEntityReferenceResponse                                 `pulumi:"subnet"`
}

type VirtualMachineScaleSetIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIPConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) ToVirtualMachineScaleSetIPConfigurationResponseOutput() VirtualMachineScaleSetIPConfigurationResponseOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) ToVirtualMachineScaleSetIPConfigurationResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetIPConfigurationResponseOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) ApplicationGatewayBackendAddressPools() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) []SubResourceResponse {
		return v.ApplicationGatewayBackendAddressPools
	}).(SubResourceResponseArrayOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) ApplicationSecurityGroups() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) []SubResourceResponse {
		return v.ApplicationSecurityGroups
	}).(SubResourceResponseArrayOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) LoadBalancerBackendAddressPools() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) []SubResourceResponse {
		return v.LoadBalancerBackendAddressPools
	}).(SubResourceResponseArrayOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) LoadBalancerInboundNatPools() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) []SubResourceResponse {
		return v.LoadBalancerInboundNatPools
	}).(SubResourceResponseArrayOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) PrivateIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) *string { return v.PrivateIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) PublicIPAddressConfiguration() VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) *VirtualMachineScaleSetPublicIPAddressConfigurationResponse {
		return v.PublicIPAddressConfiguration
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput)
}

func (o VirtualMachineScaleSetIPConfigurationResponseOutput) Subnet() ApiEntityReferenceResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIPConfigurationResponse) *ApiEntityReferenceResponse { return v.Subnet }).(ApiEntityReferenceResponsePtrOutput)
}

type VirtualMachineScaleSetIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetIPConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetIPConfigurationResponseArrayOutput) ToVirtualMachineScaleSetIPConfigurationResponseArrayOutput() VirtualMachineScaleSetIPConfigurationResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationResponseArrayOutput) ToVirtualMachineScaleSetIPConfigurationResponseArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetIPConfigurationResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetIPConfigurationResponse {
		return vs[0].([]VirtualMachineScaleSetIPConfigurationResponse)[vs[1].(int)]
	}).(VirtualMachineScaleSetIPConfigurationResponseOutput)
}

type VirtualMachineScaleSetIdentity struct {
	Type                   *ResourceIdentityType  `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type VirtualMachineScaleSetIdentityInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetIdentityOutput() VirtualMachineScaleSetIdentityOutput
	ToVirtualMachineScaleSetIdentityOutputWithContext(context.Context) VirtualMachineScaleSetIdentityOutput
}

type VirtualMachineScaleSetIdentityArgs struct {
	Type                   ResourceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput              `pulumi:"userAssignedIdentities"`
}

func (VirtualMachineScaleSetIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIdentity)(nil)).Elem()
}

func (i VirtualMachineScaleSetIdentityArgs) ToVirtualMachineScaleSetIdentityOutput() VirtualMachineScaleSetIdentityOutput {
	return i.ToVirtualMachineScaleSetIdentityOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetIdentityArgs) ToVirtualMachineScaleSetIdentityOutputWithContext(ctx context.Context) VirtualMachineScaleSetIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetIdentityOutput)
}

func (i VirtualMachineScaleSetIdentityArgs) ToVirtualMachineScaleSetIdentityPtrOutput() VirtualMachineScaleSetIdentityPtrOutput {
	return i.ToVirtualMachineScaleSetIdentityPtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetIdentityArgs) ToVirtualMachineScaleSetIdentityPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetIdentityOutput).ToVirtualMachineScaleSetIdentityPtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetIdentityPtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetIdentityPtrOutput() VirtualMachineScaleSetIdentityPtrOutput
	ToVirtualMachineScaleSetIdentityPtrOutputWithContext(context.Context) VirtualMachineScaleSetIdentityPtrOutput
}

type virtualMachineScaleSetIdentityPtrType VirtualMachineScaleSetIdentityArgs

func VirtualMachineScaleSetIdentityPtr(v *VirtualMachineScaleSetIdentityArgs) VirtualMachineScaleSetIdentityPtrInput {
	return (*virtualMachineScaleSetIdentityPtrType)(v)
}

func (*virtualMachineScaleSetIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetIdentity)(nil)).Elem()
}

func (i *virtualMachineScaleSetIdentityPtrType) ToVirtualMachineScaleSetIdentityPtrOutput() VirtualMachineScaleSetIdentityPtrOutput {
	return i.ToVirtualMachineScaleSetIdentityPtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetIdentityPtrType) ToVirtualMachineScaleSetIdentityPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetIdentityPtrOutput)
}

type VirtualMachineScaleSetIdentityOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIdentity)(nil)).Elem()
}

func (o VirtualMachineScaleSetIdentityOutput) ToVirtualMachineScaleSetIdentityOutput() VirtualMachineScaleSetIdentityOutput {
	return o
}

func (o VirtualMachineScaleSetIdentityOutput) ToVirtualMachineScaleSetIdentityOutputWithContext(ctx context.Context) VirtualMachineScaleSetIdentityOutput {
	return o
}

func (o VirtualMachineScaleSetIdentityOutput) ToVirtualMachineScaleSetIdentityPtrOutput() VirtualMachineScaleSetIdentityPtrOutput {
	return o.ToVirtualMachineScaleSetIdentityPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetIdentityOutput) ToVirtualMachineScaleSetIdentityPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetIdentity) *VirtualMachineScaleSetIdentity {
		return &v
	}).(VirtualMachineScaleSetIdentityPtrOutput)
}

func (o VirtualMachineScaleSetIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

func (o VirtualMachineScaleSetIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type VirtualMachineScaleSetIdentityPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetIdentity)(nil)).Elem()
}

func (o VirtualMachineScaleSetIdentityPtrOutput) ToVirtualMachineScaleSetIdentityPtrOutput() VirtualMachineScaleSetIdentityPtrOutput {
	return o
}

func (o VirtualMachineScaleSetIdentityPtrOutput) ToVirtualMachineScaleSetIdentityPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetIdentityPtrOutput {
	return o
}

func (o VirtualMachineScaleSetIdentityPtrOutput) Elem() VirtualMachineScaleSetIdentityOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetIdentity) VirtualMachineScaleSetIdentity {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetIdentity
		return ret
	}).(VirtualMachineScaleSetIdentityOutput)
}

func (o VirtualMachineScaleSetIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

func (o VirtualMachineScaleSetIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type VirtualMachineScaleSetIdentityResponse struct {
	PrincipalId            string                                                          `pulumi:"principalId"`
	TenantId               string                                                          `pulumi:"tenantId"`
	Type                   *string                                                         `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentitiesResponseUserAssignedIdentities `pulumi:"userAssignedIdentities"`
}

type VirtualMachineScaleSetIdentityResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIdentityResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetIdentityResponseOutput) ToVirtualMachineScaleSetIdentityResponseOutput() VirtualMachineScaleSetIdentityResponseOutput {
	return o
}

func (o VirtualMachineScaleSetIdentityResponseOutput) ToVirtualMachineScaleSetIdentityResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetIdentityResponseOutput {
	return o
}

func (o VirtualMachineScaleSetIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentitiesResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIdentityResponse) map[string]UserAssignedIdentitiesResponseUserAssignedIdentities {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentitiesResponseUserAssignedIdentitiesMapOutput)
}

type VirtualMachineScaleSetIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetIdentityResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetIdentityResponsePtrOutput) ToVirtualMachineScaleSetIdentityResponsePtrOutput() VirtualMachineScaleSetIdentityResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetIdentityResponsePtrOutput) ToVirtualMachineScaleSetIdentityResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetIdentityResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetIdentityResponsePtrOutput) Elem() VirtualMachineScaleSetIdentityResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetIdentityResponse) VirtualMachineScaleSetIdentityResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetIdentityResponse
		return ret
	}).(VirtualMachineScaleSetIdentityResponseOutput)
}

func (o VirtualMachineScaleSetIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentitiesResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetIdentityResponse) map[string]UserAssignedIdentitiesResponseUserAssignedIdentities {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentitiesResponseUserAssignedIdentitiesMapOutput)
}

type VirtualMachineScaleSetIpTag struct {
	IpTagType *string `pulumi:"ipTagType"`
	Tag       *string `pulumi:"tag"`
}





type VirtualMachineScaleSetIpTagInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetIpTagOutput() VirtualMachineScaleSetIpTagOutput
	ToVirtualMachineScaleSetIpTagOutputWithContext(context.Context) VirtualMachineScaleSetIpTagOutput
}

type VirtualMachineScaleSetIpTagArgs struct {
	IpTagType pulumi.StringPtrInput `pulumi:"ipTagType"`
	Tag       pulumi.StringPtrInput `pulumi:"tag"`
}

func (VirtualMachineScaleSetIpTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIpTag)(nil)).Elem()
}

func (i VirtualMachineScaleSetIpTagArgs) ToVirtualMachineScaleSetIpTagOutput() VirtualMachineScaleSetIpTagOutput {
	return i.ToVirtualMachineScaleSetIpTagOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetIpTagArgs) ToVirtualMachineScaleSetIpTagOutputWithContext(ctx context.Context) VirtualMachineScaleSetIpTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetIpTagOutput)
}





type VirtualMachineScaleSetIpTagArrayInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetIpTagArrayOutput() VirtualMachineScaleSetIpTagArrayOutput
	ToVirtualMachineScaleSetIpTagArrayOutputWithContext(context.Context) VirtualMachineScaleSetIpTagArrayOutput
}

type VirtualMachineScaleSetIpTagArray []VirtualMachineScaleSetIpTagInput

func (VirtualMachineScaleSetIpTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetIpTag)(nil)).Elem()
}

func (i VirtualMachineScaleSetIpTagArray) ToVirtualMachineScaleSetIpTagArrayOutput() VirtualMachineScaleSetIpTagArrayOutput {
	return i.ToVirtualMachineScaleSetIpTagArrayOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetIpTagArray) ToVirtualMachineScaleSetIpTagArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetIpTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetIpTagArrayOutput)
}

type VirtualMachineScaleSetIpTagOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIpTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIpTag)(nil)).Elem()
}

func (o VirtualMachineScaleSetIpTagOutput) ToVirtualMachineScaleSetIpTagOutput() VirtualMachineScaleSetIpTagOutput {
	return o
}

func (o VirtualMachineScaleSetIpTagOutput) ToVirtualMachineScaleSetIpTagOutputWithContext(ctx context.Context) VirtualMachineScaleSetIpTagOutput {
	return o
}

func (o VirtualMachineScaleSetIpTagOutput) IpTagType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIpTag) *string { return v.IpTagType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIpTagOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIpTag) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetIpTagArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIpTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetIpTag)(nil)).Elem()
}

func (o VirtualMachineScaleSetIpTagArrayOutput) ToVirtualMachineScaleSetIpTagArrayOutput() VirtualMachineScaleSetIpTagArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIpTagArrayOutput) ToVirtualMachineScaleSetIpTagArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetIpTagArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIpTagArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetIpTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetIpTag {
		return vs[0].([]VirtualMachineScaleSetIpTag)[vs[1].(int)]
	}).(VirtualMachineScaleSetIpTagOutput)
}

type VirtualMachineScaleSetIpTagResponse struct {
	IpTagType *string `pulumi:"ipTagType"`
	Tag       *string `pulumi:"tag"`
}

type VirtualMachineScaleSetIpTagResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIpTagResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetIpTagResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetIpTagResponseOutput) ToVirtualMachineScaleSetIpTagResponseOutput() VirtualMachineScaleSetIpTagResponseOutput {
	return o
}

func (o VirtualMachineScaleSetIpTagResponseOutput) ToVirtualMachineScaleSetIpTagResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetIpTagResponseOutput {
	return o
}

func (o VirtualMachineScaleSetIpTagResponseOutput) IpTagType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIpTagResponse) *string { return v.IpTagType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetIpTagResponseOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetIpTagResponse) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetIpTagResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetIpTagResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetIpTagResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetIpTagResponseArrayOutput) ToVirtualMachineScaleSetIpTagResponseArrayOutput() VirtualMachineScaleSetIpTagResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIpTagResponseArrayOutput) ToVirtualMachineScaleSetIpTagResponseArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetIpTagResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetIpTagResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetIpTagResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetIpTagResponse {
		return vs[0].([]VirtualMachineScaleSetIpTagResponse)[vs[1].(int)]
	}).(VirtualMachineScaleSetIpTagResponseOutput)
}

type VirtualMachineScaleSetManagedDiskParameters struct {
	DiskEncryptionSet  *DiskEncryptionSetParameters `pulumi:"diskEncryptionSet"`
	SecurityProfile    *VMDiskSecurityProfile       `pulumi:"securityProfile"`
	StorageAccountType *string                      `pulumi:"storageAccountType"`
}





type VirtualMachineScaleSetManagedDiskParametersInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetManagedDiskParametersOutput() VirtualMachineScaleSetManagedDiskParametersOutput
	ToVirtualMachineScaleSetManagedDiskParametersOutputWithContext(context.Context) VirtualMachineScaleSetManagedDiskParametersOutput
}

type VirtualMachineScaleSetManagedDiskParametersArgs struct {
	DiskEncryptionSet  DiskEncryptionSetParametersPtrInput `pulumi:"diskEncryptionSet"`
	SecurityProfile    VMDiskSecurityProfilePtrInput       `pulumi:"securityProfile"`
	StorageAccountType pulumi.StringPtrInput               `pulumi:"storageAccountType"`
}

func (VirtualMachineScaleSetManagedDiskParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetManagedDiskParameters)(nil)).Elem()
}

func (i VirtualMachineScaleSetManagedDiskParametersArgs) ToVirtualMachineScaleSetManagedDiskParametersOutput() VirtualMachineScaleSetManagedDiskParametersOutput {
	return i.ToVirtualMachineScaleSetManagedDiskParametersOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetManagedDiskParametersArgs) ToVirtualMachineScaleSetManagedDiskParametersOutputWithContext(ctx context.Context) VirtualMachineScaleSetManagedDiskParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetManagedDiskParametersOutput)
}

func (i VirtualMachineScaleSetManagedDiskParametersArgs) ToVirtualMachineScaleSetManagedDiskParametersPtrOutput() VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return i.ToVirtualMachineScaleSetManagedDiskParametersPtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetManagedDiskParametersArgs) ToVirtualMachineScaleSetManagedDiskParametersPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetManagedDiskParametersOutput).ToVirtualMachineScaleSetManagedDiskParametersPtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetManagedDiskParametersPtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetManagedDiskParametersPtrOutput() VirtualMachineScaleSetManagedDiskParametersPtrOutput
	ToVirtualMachineScaleSetManagedDiskParametersPtrOutputWithContext(context.Context) VirtualMachineScaleSetManagedDiskParametersPtrOutput
}

type virtualMachineScaleSetManagedDiskParametersPtrType VirtualMachineScaleSetManagedDiskParametersArgs

func VirtualMachineScaleSetManagedDiskParametersPtr(v *VirtualMachineScaleSetManagedDiskParametersArgs) VirtualMachineScaleSetManagedDiskParametersPtrInput {
	return (*virtualMachineScaleSetManagedDiskParametersPtrType)(v)
}

func (*virtualMachineScaleSetManagedDiskParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetManagedDiskParameters)(nil)).Elem()
}

func (i *virtualMachineScaleSetManagedDiskParametersPtrType) ToVirtualMachineScaleSetManagedDiskParametersPtrOutput() VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return i.ToVirtualMachineScaleSetManagedDiskParametersPtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetManagedDiskParametersPtrType) ToVirtualMachineScaleSetManagedDiskParametersPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetManagedDiskParametersPtrOutput)
}

type VirtualMachineScaleSetManagedDiskParametersOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetManagedDiskParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetManagedDiskParameters)(nil)).Elem()
}

func (o VirtualMachineScaleSetManagedDiskParametersOutput) ToVirtualMachineScaleSetManagedDiskParametersOutput() VirtualMachineScaleSetManagedDiskParametersOutput {
	return o
}

func (o VirtualMachineScaleSetManagedDiskParametersOutput) ToVirtualMachineScaleSetManagedDiskParametersOutputWithContext(ctx context.Context) VirtualMachineScaleSetManagedDiskParametersOutput {
	return o
}

func (o VirtualMachineScaleSetManagedDiskParametersOutput) ToVirtualMachineScaleSetManagedDiskParametersPtrOutput() VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return o.ToVirtualMachineScaleSetManagedDiskParametersPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetManagedDiskParametersOutput) ToVirtualMachineScaleSetManagedDiskParametersPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetManagedDiskParameters) *VirtualMachineScaleSetManagedDiskParameters {
		return &v
	}).(VirtualMachineScaleSetManagedDiskParametersPtrOutput)
}

func (o VirtualMachineScaleSetManagedDiskParametersOutput) DiskEncryptionSet() DiskEncryptionSetParametersPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetManagedDiskParameters) *DiskEncryptionSetParameters {
		return v.DiskEncryptionSet
	}).(DiskEncryptionSetParametersPtrOutput)
}

func (o VirtualMachineScaleSetManagedDiskParametersOutput) SecurityProfile() VMDiskSecurityProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetManagedDiskParameters) *VMDiskSecurityProfile { return v.SecurityProfile }).(VMDiskSecurityProfilePtrOutput)
}

func (o VirtualMachineScaleSetManagedDiskParametersOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetManagedDiskParameters) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetManagedDiskParametersPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetManagedDiskParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetManagedDiskParameters)(nil)).Elem()
}

func (o VirtualMachineScaleSetManagedDiskParametersPtrOutput) ToVirtualMachineScaleSetManagedDiskParametersPtrOutput() VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return o
}

func (o VirtualMachineScaleSetManagedDiskParametersPtrOutput) ToVirtualMachineScaleSetManagedDiskParametersPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return o
}

func (o VirtualMachineScaleSetManagedDiskParametersPtrOutput) Elem() VirtualMachineScaleSetManagedDiskParametersOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetManagedDiskParameters) VirtualMachineScaleSetManagedDiskParameters {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetManagedDiskParameters
		return ret
	}).(VirtualMachineScaleSetManagedDiskParametersOutput)
}

func (o VirtualMachineScaleSetManagedDiskParametersPtrOutput) DiskEncryptionSet() DiskEncryptionSetParametersPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetManagedDiskParameters) *DiskEncryptionSetParameters {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionSet
	}).(DiskEncryptionSetParametersPtrOutput)
}

func (o VirtualMachineScaleSetManagedDiskParametersPtrOutput) SecurityProfile() VMDiskSecurityProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetManagedDiskParameters) *VMDiskSecurityProfile {
		if v == nil {
			return nil
		}
		return v.SecurityProfile
	}).(VMDiskSecurityProfilePtrOutput)
}

func (o VirtualMachineScaleSetManagedDiskParametersPtrOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetManagedDiskParameters) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountType
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetManagedDiskParametersResponse struct {
	DiskEncryptionSet  *DiskEncryptionSetParametersResponse `pulumi:"diskEncryptionSet"`
	SecurityProfile    *VMDiskSecurityProfileResponse       `pulumi:"securityProfile"`
	StorageAccountType *string                              `pulumi:"storageAccountType"`
}

type VirtualMachineScaleSetManagedDiskParametersResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetManagedDiskParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetManagedDiskParametersResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetManagedDiskParametersResponseOutput) ToVirtualMachineScaleSetManagedDiskParametersResponseOutput() VirtualMachineScaleSetManagedDiskParametersResponseOutput {
	return o
}

func (o VirtualMachineScaleSetManagedDiskParametersResponseOutput) ToVirtualMachineScaleSetManagedDiskParametersResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetManagedDiskParametersResponseOutput {
	return o
}

func (o VirtualMachineScaleSetManagedDiskParametersResponseOutput) DiskEncryptionSet() DiskEncryptionSetParametersResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetManagedDiskParametersResponse) *DiskEncryptionSetParametersResponse {
		return v.DiskEncryptionSet
	}).(DiskEncryptionSetParametersResponsePtrOutput)
}

func (o VirtualMachineScaleSetManagedDiskParametersResponseOutput) SecurityProfile() VMDiskSecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetManagedDiskParametersResponse) *VMDiskSecurityProfileResponse {
		return v.SecurityProfile
	}).(VMDiskSecurityProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetManagedDiskParametersResponseOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetManagedDiskParametersResponse) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetManagedDiskParametersResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput) ToVirtualMachineScaleSetManagedDiskParametersResponsePtrOutput() VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput) ToVirtualMachineScaleSetManagedDiskParametersResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput) Elem() VirtualMachineScaleSetManagedDiskParametersResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetManagedDiskParametersResponse) VirtualMachineScaleSetManagedDiskParametersResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetManagedDiskParametersResponse
		return ret
	}).(VirtualMachineScaleSetManagedDiskParametersResponseOutput)
}

func (o VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput) DiskEncryptionSet() DiskEncryptionSetParametersResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetManagedDiskParametersResponse) *DiskEncryptionSetParametersResponse {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionSet
	}).(DiskEncryptionSetParametersResponsePtrOutput)
}

func (o VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput) SecurityProfile() VMDiskSecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetManagedDiskParametersResponse) *VMDiskSecurityProfileResponse {
		if v == nil {
			return nil
		}
		return v.SecurityProfile
	}).(VMDiskSecurityProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetManagedDiskParametersResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountType
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetNetworkConfiguration struct {
	DeleteOption                *string                                                `pulumi:"deleteOption"`
	DnsSettings                 *VirtualMachineScaleSetNetworkConfigurationDnsSettings `pulumi:"dnsSettings"`
	EnableAcceleratedNetworking *bool                                                  `pulumi:"enableAcceleratedNetworking"`
	EnableFpga                  *bool                                                  `pulumi:"enableFpga"`
	EnableIPForwarding          *bool                                                  `pulumi:"enableIPForwarding"`
	Id                          *string                                                `pulumi:"id"`
	IpConfigurations            []VirtualMachineScaleSetIPConfiguration                `pulumi:"ipConfigurations"`
	Name                        string                                                 `pulumi:"name"`
	NetworkSecurityGroup        *SubResource                                           `pulumi:"networkSecurityGroup"`
	Primary                     *bool                                                  `pulumi:"primary"`
}





type VirtualMachineScaleSetNetworkConfigurationInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetNetworkConfigurationOutput() VirtualMachineScaleSetNetworkConfigurationOutput
	ToVirtualMachineScaleSetNetworkConfigurationOutputWithContext(context.Context) VirtualMachineScaleSetNetworkConfigurationOutput
}

type VirtualMachineScaleSetNetworkConfigurationArgs struct {
	DeleteOption                pulumi.StringPtrInput                                         `pulumi:"deleteOption"`
	DnsSettings                 VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrInput `pulumi:"dnsSettings"`
	EnableAcceleratedNetworking pulumi.BoolPtrInput                                           `pulumi:"enableAcceleratedNetworking"`
	EnableFpga                  pulumi.BoolPtrInput                                           `pulumi:"enableFpga"`
	EnableIPForwarding          pulumi.BoolPtrInput                                           `pulumi:"enableIPForwarding"`
	Id                          pulumi.StringPtrInput                                         `pulumi:"id"`
	IpConfigurations            VirtualMachineScaleSetIPConfigurationArrayInput               `pulumi:"ipConfigurations"`
	Name                        pulumi.StringInput                                            `pulumi:"name"`
	NetworkSecurityGroup        SubResourcePtrInput                                           `pulumi:"networkSecurityGroup"`
	Primary                     pulumi.BoolPtrInput                                           `pulumi:"primary"`
}

func (VirtualMachineScaleSetNetworkConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkConfiguration)(nil)).Elem()
}

func (i VirtualMachineScaleSetNetworkConfigurationArgs) ToVirtualMachineScaleSetNetworkConfigurationOutput() VirtualMachineScaleSetNetworkConfigurationOutput {
	return i.ToVirtualMachineScaleSetNetworkConfigurationOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetNetworkConfigurationArgs) ToVirtualMachineScaleSetNetworkConfigurationOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkConfigurationOutput)
}





type VirtualMachineScaleSetNetworkConfigurationArrayInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetNetworkConfigurationArrayOutput() VirtualMachineScaleSetNetworkConfigurationArrayOutput
	ToVirtualMachineScaleSetNetworkConfigurationArrayOutputWithContext(context.Context) VirtualMachineScaleSetNetworkConfigurationArrayOutput
}

type VirtualMachineScaleSetNetworkConfigurationArray []VirtualMachineScaleSetNetworkConfigurationInput

func (VirtualMachineScaleSetNetworkConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetNetworkConfiguration)(nil)).Elem()
}

func (i VirtualMachineScaleSetNetworkConfigurationArray) ToVirtualMachineScaleSetNetworkConfigurationArrayOutput() VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return i.ToVirtualMachineScaleSetNetworkConfigurationArrayOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetNetworkConfigurationArray) ToVirtualMachineScaleSetNetworkConfigurationArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkConfigurationArrayOutput)
}

type VirtualMachineScaleSetNetworkConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkConfiguration)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) ToVirtualMachineScaleSetNetworkConfigurationOutput() VirtualMachineScaleSetNetworkConfigurationOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) ToVirtualMachineScaleSetNetworkConfigurationOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) DnsSettings() VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) *VirtualMachineScaleSetNetworkConfigurationDnsSettings {
		return v.DnsSettings
	}).(VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) *bool { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) EnableFpga() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) *bool { return v.EnableFpga }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) EnableIPForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) *bool { return v.EnableIPForwarding }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) IpConfigurations() VirtualMachineScaleSetIPConfigurationArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) []VirtualMachineScaleSetIPConfiguration {
		return v.IpConfigurations
	}).(VirtualMachineScaleSetIPConfigurationArrayOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) NetworkSecurityGroup() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) *SubResource { return v.NetworkSecurityGroup }).(SubResourcePtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfiguration) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetNetworkConfigurationArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetNetworkConfiguration)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationArrayOutput) ToVirtualMachineScaleSetNetworkConfigurationArrayOutput() VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationArrayOutput) ToVirtualMachineScaleSetNetworkConfigurationArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetNetworkConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetNetworkConfiguration {
		return vs[0].([]VirtualMachineScaleSetNetworkConfiguration)[vs[1].(int)]
	}).(VirtualMachineScaleSetNetworkConfigurationOutput)
}

type VirtualMachineScaleSetNetworkConfigurationDnsSettings struct {
	DnsServers []string `pulumi:"dnsServers"`
}





type VirtualMachineScaleSetNetworkConfigurationDnsSettingsInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput
	ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsOutputWithContext(context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput
}

type VirtualMachineScaleSetNetworkConfigurationDnsSettingsArgs struct {
	DnsServers pulumi.StringArrayInput `pulumi:"dnsServers"`
}

func (VirtualMachineScaleSetNetworkConfigurationDnsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkConfigurationDnsSettings)(nil)).Elem()
}

func (i VirtualMachineScaleSetNetworkConfigurationDnsSettingsArgs) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput {
	return i.ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetNetworkConfigurationDnsSettingsArgs) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput)
}

func (i VirtualMachineScaleSetNetworkConfigurationDnsSettingsArgs) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput {
	return i.ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetNetworkConfigurationDnsSettingsArgs) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput).ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput
	ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutputWithContext(context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput
}

type virtualMachineScaleSetNetworkConfigurationDnsSettingsPtrType VirtualMachineScaleSetNetworkConfigurationDnsSettingsArgs

func VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtr(v *VirtualMachineScaleSetNetworkConfigurationDnsSettingsArgs) VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrInput {
	return (*virtualMachineScaleSetNetworkConfigurationDnsSettingsPtrType)(v)
}

func (*virtualMachineScaleSetNetworkConfigurationDnsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetNetworkConfigurationDnsSettings)(nil)).Elem()
}

func (i *virtualMachineScaleSetNetworkConfigurationDnsSettingsPtrType) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput {
	return i.ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetNetworkConfigurationDnsSettingsPtrType) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput)
}

type VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkConfigurationDnsSettings)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput {
	return o.ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetNetworkConfigurationDnsSettings) *VirtualMachineScaleSetNetworkConfigurationDnsSettings {
		return &v
	}).(VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationDnsSettings) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

type VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetNetworkConfigurationDnsSettings)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput) Elem() VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkConfigurationDnsSettings) VirtualMachineScaleSetNetworkConfigurationDnsSettings {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetNetworkConfigurationDnsSettings
		return ret
	}).(VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkConfigurationDnsSettings) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

type VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse struct {
	DnsServers []string `pulumi:"dnsServers"`
}

type VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

type VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput() VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput) ToVirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput) Elem() VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse) VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse
		return ret
	}).(VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

type VirtualMachineScaleSetNetworkConfigurationResponse struct {
	DeleteOption                *string                                                        `pulumi:"deleteOption"`
	DnsSettings                 *VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse `pulumi:"dnsSettings"`
	EnableAcceleratedNetworking *bool                                                          `pulumi:"enableAcceleratedNetworking"`
	EnableFpga                  *bool                                                          `pulumi:"enableFpga"`
	EnableIPForwarding          *bool                                                          `pulumi:"enableIPForwarding"`
	Id                          *string                                                        `pulumi:"id"`
	IpConfigurations            []VirtualMachineScaleSetIPConfigurationResponse                `pulumi:"ipConfigurations"`
	Name                        string                                                         `pulumi:"name"`
	NetworkSecurityGroup        *SubResourceResponse                                           `pulumi:"networkSecurityGroup"`
	Primary                     *bool                                                          `pulumi:"primary"`
}

type VirtualMachineScaleSetNetworkConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) ToVirtualMachineScaleSetNetworkConfigurationResponseOutput() VirtualMachineScaleSetNetworkConfigurationResponseOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) ToVirtualMachineScaleSetNetworkConfigurationResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationResponseOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) DnsSettings() VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) *VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponse {
		return v.DnsSettings
	}).(VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) *bool { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) EnableFpga() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) *bool { return v.EnableFpga }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) EnableIPForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) *bool { return v.EnableIPForwarding }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) IpConfigurations() VirtualMachineScaleSetIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) []VirtualMachineScaleSetIPConfigurationResponse {
		return v.IpConfigurations
	}).(VirtualMachineScaleSetIPConfigurationResponseArrayOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) NetworkSecurityGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) *SubResourceResponse {
		return v.NetworkSecurityGroup
	}).(SubResourceResponsePtrOutput)
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkConfigurationResponse) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachineScaleSetNetworkConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput) ToVirtualMachineScaleSetNetworkConfigurationResponseArrayOutput() VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput) ToVirtualMachineScaleSetNetworkConfigurationResponseArrayOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput) Index(i pulumi.IntInput) VirtualMachineScaleSetNetworkConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachineScaleSetNetworkConfigurationResponse {
		return vs[0].([]VirtualMachineScaleSetNetworkConfigurationResponse)[vs[1].(int)]
	}).(VirtualMachineScaleSetNetworkConfigurationResponseOutput)
}

type VirtualMachineScaleSetNetworkProfile struct {
	HealthProbe                    *ApiEntityReference                          `pulumi:"healthProbe"`
	NetworkApiVersion              *string                                      `pulumi:"networkApiVersion"`
	NetworkInterfaceConfigurations []VirtualMachineScaleSetNetworkConfiguration `pulumi:"networkInterfaceConfigurations"`
}





type VirtualMachineScaleSetNetworkProfileInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetNetworkProfileOutput() VirtualMachineScaleSetNetworkProfileOutput
	ToVirtualMachineScaleSetNetworkProfileOutputWithContext(context.Context) VirtualMachineScaleSetNetworkProfileOutput
}

type VirtualMachineScaleSetNetworkProfileArgs struct {
	HealthProbe                    ApiEntityReferencePtrInput                           `pulumi:"healthProbe"`
	NetworkApiVersion              pulumi.StringPtrInput                                `pulumi:"networkApiVersion"`
	NetworkInterfaceConfigurations VirtualMachineScaleSetNetworkConfigurationArrayInput `pulumi:"networkInterfaceConfigurations"`
}

func (VirtualMachineScaleSetNetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkProfile)(nil)).Elem()
}

func (i VirtualMachineScaleSetNetworkProfileArgs) ToVirtualMachineScaleSetNetworkProfileOutput() VirtualMachineScaleSetNetworkProfileOutput {
	return i.ToVirtualMachineScaleSetNetworkProfileOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetNetworkProfileArgs) ToVirtualMachineScaleSetNetworkProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkProfileOutput)
}

func (i VirtualMachineScaleSetNetworkProfileArgs) ToVirtualMachineScaleSetNetworkProfilePtrOutput() VirtualMachineScaleSetNetworkProfilePtrOutput {
	return i.ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetNetworkProfileArgs) ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkProfileOutput).ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetNetworkProfilePtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetNetworkProfilePtrOutput() VirtualMachineScaleSetNetworkProfilePtrOutput
	ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(context.Context) VirtualMachineScaleSetNetworkProfilePtrOutput
}

type virtualMachineScaleSetNetworkProfilePtrType VirtualMachineScaleSetNetworkProfileArgs

func VirtualMachineScaleSetNetworkProfilePtr(v *VirtualMachineScaleSetNetworkProfileArgs) VirtualMachineScaleSetNetworkProfilePtrInput {
	return (*virtualMachineScaleSetNetworkProfilePtrType)(v)
}

func (*virtualMachineScaleSetNetworkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetNetworkProfile)(nil)).Elem()
}

func (i *virtualMachineScaleSetNetworkProfilePtrType) ToVirtualMachineScaleSetNetworkProfilePtrOutput() VirtualMachineScaleSetNetworkProfilePtrOutput {
	return i.ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetNetworkProfilePtrType) ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetNetworkProfilePtrOutput)
}

type VirtualMachineScaleSetNetworkProfileOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkProfileOutput) ToVirtualMachineScaleSetNetworkProfileOutput() VirtualMachineScaleSetNetworkProfileOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfileOutput) ToVirtualMachineScaleSetNetworkProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfileOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfileOutput) ToVirtualMachineScaleSetNetworkProfilePtrOutput() VirtualMachineScaleSetNetworkProfilePtrOutput {
	return o.ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetNetworkProfileOutput) ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetNetworkProfile) *VirtualMachineScaleSetNetworkProfile {
		return &v
	}).(VirtualMachineScaleSetNetworkProfilePtrOutput)
}

func (o VirtualMachineScaleSetNetworkProfileOutput) HealthProbe() ApiEntityReferencePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkProfile) *ApiEntityReference { return v.HealthProbe }).(ApiEntityReferencePtrOutput)
}

func (o VirtualMachineScaleSetNetworkProfileOutput) NetworkApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkProfile) *string { return v.NetworkApiVersion }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetNetworkProfileOutput) NetworkInterfaceConfigurations() VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkProfile) []VirtualMachineScaleSetNetworkConfiguration {
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineScaleSetNetworkConfigurationArrayOutput)
}

type VirtualMachineScaleSetNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetNetworkProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkProfilePtrOutput) ToVirtualMachineScaleSetNetworkProfilePtrOutput() VirtualMachineScaleSetNetworkProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfilePtrOutput) ToVirtualMachineScaleSetNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfilePtrOutput) Elem() VirtualMachineScaleSetNetworkProfileOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkProfile) VirtualMachineScaleSetNetworkProfile {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetNetworkProfile
		return ret
	}).(VirtualMachineScaleSetNetworkProfileOutput)
}

func (o VirtualMachineScaleSetNetworkProfilePtrOutput) HealthProbe() ApiEntityReferencePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkProfile) *ApiEntityReference {
		if v == nil {
			return nil
		}
		return v.HealthProbe
	}).(ApiEntityReferencePtrOutput)
}

func (o VirtualMachineScaleSetNetworkProfilePtrOutput) NetworkApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.NetworkApiVersion
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetNetworkProfilePtrOutput) NetworkInterfaceConfigurations() VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkProfile) []VirtualMachineScaleSetNetworkConfiguration {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineScaleSetNetworkConfigurationArrayOutput)
}

type VirtualMachineScaleSetNetworkProfileResponse struct {
	HealthProbe                    *ApiEntityReferenceResponse                          `pulumi:"healthProbe"`
	NetworkApiVersion              *string                                              `pulumi:"networkApiVersion"`
	NetworkInterfaceConfigurations []VirtualMachineScaleSetNetworkConfigurationResponse `pulumi:"networkInterfaceConfigurations"`
}

type VirtualMachineScaleSetNetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetNetworkProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkProfileResponseOutput) ToVirtualMachineScaleSetNetworkProfileResponseOutput() VirtualMachineScaleSetNetworkProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfileResponseOutput) ToVirtualMachineScaleSetNetworkProfileResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfileResponseOutput) HealthProbe() ApiEntityReferenceResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkProfileResponse) *ApiEntityReferenceResponse { return v.HealthProbe }).(ApiEntityReferenceResponsePtrOutput)
}

func (o VirtualMachineScaleSetNetworkProfileResponseOutput) NetworkApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkProfileResponse) *string { return v.NetworkApiVersion }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetNetworkProfileResponseOutput) NetworkInterfaceConfigurations() VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetNetworkProfileResponse) []VirtualMachineScaleSetNetworkConfigurationResponse {
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput)
}

type VirtualMachineScaleSetNetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetNetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetNetworkProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetNetworkProfileResponsePtrOutput) ToVirtualMachineScaleSetNetworkProfileResponsePtrOutput() VirtualMachineScaleSetNetworkProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfileResponsePtrOutput) ToVirtualMachineScaleSetNetworkProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetNetworkProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetNetworkProfileResponsePtrOutput) Elem() VirtualMachineScaleSetNetworkProfileResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkProfileResponse) VirtualMachineScaleSetNetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetNetworkProfileResponse
		return ret
	}).(VirtualMachineScaleSetNetworkProfileResponseOutput)
}

func (o VirtualMachineScaleSetNetworkProfileResponsePtrOutput) HealthProbe() ApiEntityReferenceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkProfileResponse) *ApiEntityReferenceResponse {
		if v == nil {
			return nil
		}
		return v.HealthProbe
	}).(ApiEntityReferenceResponsePtrOutput)
}

func (o VirtualMachineScaleSetNetworkProfileResponsePtrOutput) NetworkApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.NetworkApiVersion
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetNetworkProfileResponsePtrOutput) NetworkInterfaceConfigurations() VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetNetworkProfileResponse) []VirtualMachineScaleSetNetworkConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput)
}

type VirtualMachineScaleSetOSDisk struct {
	Caching                 *CachingTypes                                `pulumi:"caching"`
	CreateOption            string                                       `pulumi:"createOption"`
	DeleteOption            *string                                      `pulumi:"deleteOption"`
	DiffDiskSettings        *DiffDiskSettings                            `pulumi:"diffDiskSettings"`
	DiskSizeGB              *int                                         `pulumi:"diskSizeGB"`
	Image                   *VirtualHardDisk                             `pulumi:"image"`
	ManagedDisk             *VirtualMachineScaleSetManagedDiskParameters `pulumi:"managedDisk"`
	Name                    *string                                      `pulumi:"name"`
	OsType                  *OperatingSystemTypes                        `pulumi:"osType"`
	VhdContainers           []string                                     `pulumi:"vhdContainers"`
	WriteAcceleratorEnabled *bool                                        `pulumi:"writeAcceleratorEnabled"`
}





type VirtualMachineScaleSetOSDiskInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetOSDiskOutput() VirtualMachineScaleSetOSDiskOutput
	ToVirtualMachineScaleSetOSDiskOutputWithContext(context.Context) VirtualMachineScaleSetOSDiskOutput
}

type VirtualMachineScaleSetOSDiskArgs struct {
	Caching                 CachingTypesPtrInput                                `pulumi:"caching"`
	CreateOption            pulumi.StringInput                                  `pulumi:"createOption"`
	DeleteOption            pulumi.StringPtrInput                               `pulumi:"deleteOption"`
	DiffDiskSettings        DiffDiskSettingsPtrInput                            `pulumi:"diffDiskSettings"`
	DiskSizeGB              pulumi.IntPtrInput                                  `pulumi:"diskSizeGB"`
	Image                   VirtualHardDiskPtrInput                             `pulumi:"image"`
	ManagedDisk             VirtualMachineScaleSetManagedDiskParametersPtrInput `pulumi:"managedDisk"`
	Name                    pulumi.StringPtrInput                               `pulumi:"name"`
	OsType                  OperatingSystemTypesPtrInput                        `pulumi:"osType"`
	VhdContainers           pulumi.StringArrayInput                             `pulumi:"vhdContainers"`
	WriteAcceleratorEnabled pulumi.BoolPtrInput                                 `pulumi:"writeAcceleratorEnabled"`
}

func (VirtualMachineScaleSetOSDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetOSDisk)(nil)).Elem()
}

func (i VirtualMachineScaleSetOSDiskArgs) ToVirtualMachineScaleSetOSDiskOutput() VirtualMachineScaleSetOSDiskOutput {
	return i.ToVirtualMachineScaleSetOSDiskOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetOSDiskArgs) ToVirtualMachineScaleSetOSDiskOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOSDiskOutput)
}

func (i VirtualMachineScaleSetOSDiskArgs) ToVirtualMachineScaleSetOSDiskPtrOutput() VirtualMachineScaleSetOSDiskPtrOutput {
	return i.ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetOSDiskArgs) ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOSDiskOutput).ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetOSDiskPtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetOSDiskPtrOutput() VirtualMachineScaleSetOSDiskPtrOutput
	ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(context.Context) VirtualMachineScaleSetOSDiskPtrOutput
}

type virtualMachineScaleSetOSDiskPtrType VirtualMachineScaleSetOSDiskArgs

func VirtualMachineScaleSetOSDiskPtr(v *VirtualMachineScaleSetOSDiskArgs) VirtualMachineScaleSetOSDiskPtrInput {
	return (*virtualMachineScaleSetOSDiskPtrType)(v)
}

func (*virtualMachineScaleSetOSDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetOSDisk)(nil)).Elem()
}

func (i *virtualMachineScaleSetOSDiskPtrType) ToVirtualMachineScaleSetOSDiskPtrOutput() VirtualMachineScaleSetOSDiskPtrOutput {
	return i.ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetOSDiskPtrType) ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOSDiskPtrOutput)
}

type VirtualMachineScaleSetOSDiskOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetOSDisk)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSDiskOutput) ToVirtualMachineScaleSetOSDiskOutput() VirtualMachineScaleSetOSDiskOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskOutput) ToVirtualMachineScaleSetOSDiskOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskOutput) ToVirtualMachineScaleSetOSDiskPtrOutput() VirtualMachineScaleSetOSDiskPtrOutput {
	return o.ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetOSDiskOutput) ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetOSDisk) *VirtualMachineScaleSetOSDisk {
		return &v
	}).(VirtualMachineScaleSetOSDiskPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *CachingTypes { return v.Caching }).(CachingTypesPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) DiffDiskSettings() DiffDiskSettingsPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *DiffDiskSettings { return v.DiffDiskSettings }).(DiffDiskSettingsPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) Image() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *VirtualHardDisk { return v.Image }).(VirtualHardDiskPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) ManagedDisk() VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *VirtualMachineScaleSetManagedDiskParameters {
		return v.ManagedDisk
	}).(VirtualMachineScaleSetManagedDiskParametersPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) OsType() OperatingSystemTypesPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *OperatingSystemTypes { return v.OsType }).(OperatingSystemTypesPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) VhdContainers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) []string { return v.VhdContainers }).(pulumi.StringArrayOutput)
}

func (o VirtualMachineScaleSetOSDiskOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDisk) *bool { return v.WriteAcceleratorEnabled }).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetOSDiskPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetOSDisk)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) ToVirtualMachineScaleSetOSDiskPtrOutput() VirtualMachineScaleSetOSDiskPtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) ToVirtualMachineScaleSetOSDiskPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskPtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) Elem() VirtualMachineScaleSetOSDiskOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) VirtualMachineScaleSetOSDisk {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetOSDisk
		return ret
	}).(VirtualMachineScaleSetOSDiskOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) Caching() CachingTypesPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *CachingTypes {
		if v == nil {
			return nil
		}
		return v.Caching
	}).(CachingTypesPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) CreateOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *string {
		if v == nil {
			return nil
		}
		return &v.CreateOption
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *string {
		if v == nil {
			return nil
		}
		return v.DeleteOption
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) DiffDiskSettings() DiffDiskSettingsPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *DiffDiskSettings {
		if v == nil {
			return nil
		}
		return v.DiffDiskSettings
	}).(DiffDiskSettingsPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) Image() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *VirtualHardDisk {
		if v == nil {
			return nil
		}
		return v.Image
	}).(VirtualHardDiskPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) ManagedDisk() VirtualMachineScaleSetManagedDiskParametersPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *VirtualMachineScaleSetManagedDiskParameters {
		if v == nil {
			return nil
		}
		return v.ManagedDisk
	}).(VirtualMachineScaleSetManagedDiskParametersPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) OsType() OperatingSystemTypesPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *OperatingSystemTypes {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(OperatingSystemTypesPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) VhdContainers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) []string {
		if v == nil {
			return nil
		}
		return v.VhdContainers
	}).(pulumi.StringArrayOutput)
}

func (o VirtualMachineScaleSetOSDiskPtrOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDisk) *bool {
		if v == nil {
			return nil
		}
		return v.WriteAcceleratorEnabled
	}).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetOSDiskResponse struct {
	Caching                 *string                                              `pulumi:"caching"`
	CreateOption            string                                               `pulumi:"createOption"`
	DeleteOption            *string                                              `pulumi:"deleteOption"`
	DiffDiskSettings        *DiffDiskSettingsResponse                            `pulumi:"diffDiskSettings"`
	DiskSizeGB              *int                                                 `pulumi:"diskSizeGB"`
	Image                   *VirtualHardDiskResponse                             `pulumi:"image"`
	ManagedDisk             *VirtualMachineScaleSetManagedDiskParametersResponse `pulumi:"managedDisk"`
	Name                    *string                                              `pulumi:"name"`
	OsType                  *string                                              `pulumi:"osType"`
	VhdContainers           []string                                             `pulumi:"vhdContainers"`
	WriteAcceleratorEnabled *bool                                                `pulumi:"writeAcceleratorEnabled"`
}

type VirtualMachineScaleSetOSDiskResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetOSDiskResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) ToVirtualMachineScaleSetOSDiskResponseOutput() VirtualMachineScaleSetOSDiskResponseOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) ToVirtualMachineScaleSetOSDiskResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskResponseOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *string { return v.Caching }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) DiffDiskSettings() DiffDiskSettingsResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *DiffDiskSettingsResponse { return v.DiffDiskSettings }).(DiffDiskSettingsResponsePtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) Image() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *VirtualHardDiskResponse { return v.Image }).(VirtualHardDiskResponsePtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) ManagedDisk() VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *VirtualMachineScaleSetManagedDiskParametersResponse {
		return v.ManagedDisk
	}).(VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) VhdContainers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) []string { return v.VhdContainers }).(pulumi.StringArrayOutput)
}

func (o VirtualMachineScaleSetOSDiskResponseOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSDiskResponse) *bool { return v.WriteAcceleratorEnabled }).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetOSDiskResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSDiskResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetOSDiskResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) ToVirtualMachineScaleSetOSDiskResponsePtrOutput() VirtualMachineScaleSetOSDiskResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) ToVirtualMachineScaleSetOSDiskResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSDiskResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) Elem() VirtualMachineScaleSetOSDiskResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) VirtualMachineScaleSetOSDiskResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetOSDiskResponse
		return ret
	}).(VirtualMachineScaleSetOSDiskResponseOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) Caching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.Caching
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) CreateOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreateOption
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeleteOption
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) DiffDiskSettings() DiffDiskSettingsResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *DiffDiskSettingsResponse {
		if v == nil {
			return nil
		}
		return v.DiffDiskSettings
	}).(DiffDiskSettingsResponsePtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) Image() VirtualHardDiskResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *VirtualHardDiskResponse {
		if v == nil {
			return nil
		}
		return v.Image
	}).(VirtualHardDiskResponsePtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) ManagedDisk() VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *VirtualMachineScaleSetManagedDiskParametersResponse {
		if v == nil {
			return nil
		}
		return v.ManagedDisk
	}).(VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) VhdContainers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) []string {
		if v == nil {
			return nil
		}
		return v.VhdContainers
	}).(pulumi.StringArrayOutput)
}

func (o VirtualMachineScaleSetOSDiskResponsePtrOutput) WriteAcceleratorEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSDiskResponse) *bool {
		if v == nil {
			return nil
		}
		return v.WriteAcceleratorEnabled
	}).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetOSProfile struct {
	AdminPassword            *string               `pulumi:"adminPassword"`
	AdminUsername            *string               `pulumi:"adminUsername"`
	AllowExtensionOperations *bool                 `pulumi:"allowExtensionOperations"`
	ComputerNamePrefix       *string               `pulumi:"computerNamePrefix"`
	CustomData               *string               `pulumi:"customData"`
	LinuxConfiguration       *LinuxConfiguration   `pulumi:"linuxConfiguration"`
	Secrets                  []VaultSecretGroup    `pulumi:"secrets"`
	WindowsConfiguration     *WindowsConfiguration `pulumi:"windowsConfiguration"`
}





type VirtualMachineScaleSetOSProfileInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetOSProfileOutput() VirtualMachineScaleSetOSProfileOutput
	ToVirtualMachineScaleSetOSProfileOutputWithContext(context.Context) VirtualMachineScaleSetOSProfileOutput
}

type VirtualMachineScaleSetOSProfileArgs struct {
	AdminPassword            pulumi.StringPtrInput        `pulumi:"adminPassword"`
	AdminUsername            pulumi.StringPtrInput        `pulumi:"adminUsername"`
	AllowExtensionOperations pulumi.BoolPtrInput          `pulumi:"allowExtensionOperations"`
	ComputerNamePrefix       pulumi.StringPtrInput        `pulumi:"computerNamePrefix"`
	CustomData               pulumi.StringPtrInput        `pulumi:"customData"`
	LinuxConfiguration       LinuxConfigurationPtrInput   `pulumi:"linuxConfiguration"`
	Secrets                  VaultSecretGroupArrayInput   `pulumi:"secrets"`
	WindowsConfiguration     WindowsConfigurationPtrInput `pulumi:"windowsConfiguration"`
}

func (VirtualMachineScaleSetOSProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetOSProfile)(nil)).Elem()
}

func (i VirtualMachineScaleSetOSProfileArgs) ToVirtualMachineScaleSetOSProfileOutput() VirtualMachineScaleSetOSProfileOutput {
	return i.ToVirtualMachineScaleSetOSProfileOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetOSProfileArgs) ToVirtualMachineScaleSetOSProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOSProfileOutput)
}

func (i VirtualMachineScaleSetOSProfileArgs) ToVirtualMachineScaleSetOSProfilePtrOutput() VirtualMachineScaleSetOSProfilePtrOutput {
	return i.ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetOSProfileArgs) ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOSProfileOutput).ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetOSProfilePtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetOSProfilePtrOutput() VirtualMachineScaleSetOSProfilePtrOutput
	ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(context.Context) VirtualMachineScaleSetOSProfilePtrOutput
}

type virtualMachineScaleSetOSProfilePtrType VirtualMachineScaleSetOSProfileArgs

func VirtualMachineScaleSetOSProfilePtr(v *VirtualMachineScaleSetOSProfileArgs) VirtualMachineScaleSetOSProfilePtrInput {
	return (*virtualMachineScaleSetOSProfilePtrType)(v)
}

func (*virtualMachineScaleSetOSProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetOSProfile)(nil)).Elem()
}

func (i *virtualMachineScaleSetOSProfilePtrType) ToVirtualMachineScaleSetOSProfilePtrOutput() VirtualMachineScaleSetOSProfilePtrOutput {
	return i.ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetOSProfilePtrType) ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOSProfilePtrOutput)
}

type VirtualMachineScaleSetOSProfileOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetOSProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSProfileOutput) ToVirtualMachineScaleSetOSProfileOutput() VirtualMachineScaleSetOSProfileOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfileOutput) ToVirtualMachineScaleSetOSProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfileOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfileOutput) ToVirtualMachineScaleSetOSProfilePtrOutput() VirtualMachineScaleSetOSProfilePtrOutput {
	return o.ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetOSProfileOutput) ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetOSProfile) *VirtualMachineScaleSetOSProfile {
		return &v
	}).(VirtualMachineScaleSetOSProfilePtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) AllowExtensionOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *bool { return v.AllowExtensionOperations }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) ComputerNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *string { return v.ComputerNamePrefix }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) LinuxConfiguration() LinuxConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *LinuxConfiguration { return v.LinuxConfiguration }).(LinuxConfigurationPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) Secrets() VaultSecretGroupArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) []VaultSecretGroup { return v.Secrets }).(VaultSecretGroupArrayOutput)
}

func (o VirtualMachineScaleSetOSProfileOutput) WindowsConfiguration() WindowsConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfile) *WindowsConfiguration { return v.WindowsConfiguration }).(WindowsConfigurationPtrOutput)
}

type VirtualMachineScaleSetOSProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetOSProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) ToVirtualMachineScaleSetOSProfilePtrOutput() VirtualMachineScaleSetOSProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) ToVirtualMachineScaleSetOSProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) Elem() VirtualMachineScaleSetOSProfileOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) VirtualMachineScaleSetOSProfile {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetOSProfile
		return ret
	}).(VirtualMachineScaleSetOSProfileOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) AllowExtensionOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *bool {
		if v == nil {
			return nil
		}
		return v.AllowExtensionOperations
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) ComputerNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *string {
		if v == nil {
			return nil
		}
		return v.ComputerNamePrefix
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *string {
		if v == nil {
			return nil
		}
		return v.CustomData
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) LinuxConfiguration() LinuxConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *LinuxConfiguration {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(LinuxConfigurationPtrOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) Secrets() VaultSecretGroupArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) []VaultSecretGroup {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(VaultSecretGroupArrayOutput)
}

func (o VirtualMachineScaleSetOSProfilePtrOutput) WindowsConfiguration() WindowsConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfile) *WindowsConfiguration {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(WindowsConfigurationPtrOutput)
}

type VirtualMachineScaleSetOSProfileResponse struct {
	AdminPassword            *string                       `pulumi:"adminPassword"`
	AdminUsername            *string                       `pulumi:"adminUsername"`
	AllowExtensionOperations *bool                         `pulumi:"allowExtensionOperations"`
	ComputerNamePrefix       *string                       `pulumi:"computerNamePrefix"`
	CustomData               *string                       `pulumi:"customData"`
	LinuxConfiguration       *LinuxConfigurationResponse   `pulumi:"linuxConfiguration"`
	Secrets                  []VaultSecretGroupResponse    `pulumi:"secrets"`
	WindowsConfiguration     *WindowsConfigurationResponse `pulumi:"windowsConfiguration"`
}

type VirtualMachineScaleSetOSProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetOSProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) ToVirtualMachineScaleSetOSProfileResponseOutput() VirtualMachineScaleSetOSProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) ToVirtualMachineScaleSetOSProfileResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) AllowExtensionOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *bool { return v.AllowExtensionOperations }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) ComputerNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *string { return v.ComputerNamePrefix }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) LinuxConfiguration() LinuxConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *LinuxConfigurationResponse {
		return v.LinuxConfiguration
	}).(LinuxConfigurationResponsePtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) Secrets() VaultSecretGroupResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) []VaultSecretGroupResponse { return v.Secrets }).(VaultSecretGroupResponseArrayOutput)
}

func (o VirtualMachineScaleSetOSProfileResponseOutput) WindowsConfiguration() WindowsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetOSProfileResponse) *WindowsConfigurationResponse {
		return v.WindowsConfiguration
	}).(WindowsConfigurationResponsePtrOutput)
}

type VirtualMachineScaleSetOSProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOSProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetOSProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) ToVirtualMachineScaleSetOSProfileResponsePtrOutput() VirtualMachineScaleSetOSProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) ToVirtualMachineScaleSetOSProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetOSProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) Elem() VirtualMachineScaleSetOSProfileResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) VirtualMachineScaleSetOSProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetOSProfileResponse
		return ret
	}).(VirtualMachineScaleSetOSProfileResponseOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) AllowExtensionOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AllowExtensionOperations
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) ComputerNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ComputerNamePrefix
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomData
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) LinuxConfiguration() LinuxConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *LinuxConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(LinuxConfigurationResponsePtrOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) Secrets() VaultSecretGroupResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) []VaultSecretGroupResponse {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(VaultSecretGroupResponseArrayOutput)
}

func (o VirtualMachineScaleSetOSProfileResponsePtrOutput) WindowsConfiguration() WindowsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetOSProfileResponse) *WindowsConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.WindowsConfiguration
	}).(WindowsConfigurationResponsePtrOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfiguration struct {
	DeleteOption           *string                                                        `pulumi:"deleteOption"`
	DnsSettings            *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings `pulumi:"dnsSettings"`
	IdleTimeoutInMinutes   *int                                                           `pulumi:"idleTimeoutInMinutes"`
	IpTags                 []VirtualMachineScaleSetIpTag                                  `pulumi:"ipTags"`
	Name                   string                                                         `pulumi:"name"`
	PublicIPAddressVersion *string                                                        `pulumi:"publicIPAddressVersion"`
	PublicIPPrefix         *SubResource                                                   `pulumi:"publicIPPrefix"`
	Sku                    *PublicIPAddressSku                                            `pulumi:"sku"`
}





type VirtualMachineScaleSetPublicIPAddressConfigurationInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetPublicIPAddressConfigurationOutput() VirtualMachineScaleSetPublicIPAddressConfigurationOutput
	ToVirtualMachineScaleSetPublicIPAddressConfigurationOutputWithContext(context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationOutput
}

type VirtualMachineScaleSetPublicIPAddressConfigurationArgs struct {
	DeleteOption           pulumi.StringPtrInput                                                 `pulumi:"deleteOption"`
	DnsSettings            VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrInput `pulumi:"dnsSettings"`
	IdleTimeoutInMinutes   pulumi.IntPtrInput                                                    `pulumi:"idleTimeoutInMinutes"`
	IpTags                 VirtualMachineScaleSetIpTagArrayInput                                 `pulumi:"ipTags"`
	Name                   pulumi.StringInput                                                    `pulumi:"name"`
	PublicIPAddressVersion pulumi.StringPtrInput                                                 `pulumi:"publicIPAddressVersion"`
	PublicIPPrefix         SubResourcePtrInput                                                   `pulumi:"publicIPPrefix"`
	Sku                    PublicIPAddressSkuPtrInput                                            `pulumi:"sku"`
}

func (VirtualMachineScaleSetPublicIPAddressConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetPublicIPAddressConfiguration)(nil)).Elem()
}

func (i VirtualMachineScaleSetPublicIPAddressConfigurationArgs) ToVirtualMachineScaleSetPublicIPAddressConfigurationOutput() VirtualMachineScaleSetPublicIPAddressConfigurationOutput {
	return i.ToVirtualMachineScaleSetPublicIPAddressConfigurationOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetPublicIPAddressConfigurationArgs) ToVirtualMachineScaleSetPublicIPAddressConfigurationOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetPublicIPAddressConfigurationOutput)
}

func (i VirtualMachineScaleSetPublicIPAddressConfigurationArgs) ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput {
	return i.ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetPublicIPAddressConfigurationArgs) ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetPublicIPAddressConfigurationOutput).ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetPublicIPAddressConfigurationPtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput
	ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutputWithContext(context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput
}

type virtualMachineScaleSetPublicIPAddressConfigurationPtrType VirtualMachineScaleSetPublicIPAddressConfigurationArgs

func VirtualMachineScaleSetPublicIPAddressConfigurationPtr(v *VirtualMachineScaleSetPublicIPAddressConfigurationArgs) VirtualMachineScaleSetPublicIPAddressConfigurationPtrInput {
	return (*virtualMachineScaleSetPublicIPAddressConfigurationPtrType)(v)
}

func (*virtualMachineScaleSetPublicIPAddressConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetPublicIPAddressConfiguration)(nil)).Elem()
}

func (i *virtualMachineScaleSetPublicIPAddressConfigurationPtrType) ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput {
	return i.ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetPublicIPAddressConfigurationPtrType) ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetPublicIPAddressConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetPublicIPAddressConfiguration)(nil)).Elem()
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationOutput() VirtualMachineScaleSetPublicIPAddressConfigurationOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput {
	return o.ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetPublicIPAddressConfiguration) *VirtualMachineScaleSetPublicIPAddressConfiguration {
		return &v
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfiguration) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) DnsSettings() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfiguration) *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings {
		return v.DnsSettings
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfiguration) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) IpTags() VirtualMachineScaleSetIpTagArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfiguration) []VirtualMachineScaleSetIpTag {
		return v.IpTags
	}).(VirtualMachineScaleSetIpTagArrayOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfiguration) *string { return v.PublicIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) PublicIPPrefix() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfiguration) *SubResource { return v.PublicIPPrefix }).(SubResourcePtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationOutput) Sku() PublicIPAddressSkuPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfiguration) *PublicIPAddressSku { return v.Sku }).(PublicIPAddressSkuPtrOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetPublicIPAddressConfiguration)(nil)).Elem()
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) Elem() VirtualMachineScaleSetPublicIPAddressConfigurationOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfiguration) VirtualMachineScaleSetPublicIPAddressConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetPublicIPAddressConfiguration
		return ret
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.DeleteOption
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) DnsSettings() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfiguration) *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings {
		if v == nil {
			return nil
		}
		return v.DnsSettings
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfiguration) *int {
		if v == nil {
			return nil
		}
		return v.IdleTimeoutInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) IpTags() VirtualMachineScaleSetIpTagArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfiguration) []VirtualMachineScaleSetIpTag {
		if v == nil {
			return nil
		}
		return v.IpTags
	}).(VirtualMachineScaleSetIpTagArrayOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.PublicIPAddressVersion
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) PublicIPPrefix() SubResourcePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfiguration) *SubResource {
		if v == nil {
			return nil
		}
		return v.PublicIPPrefix
	}).(SubResourcePtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput) Sku() PublicIPAddressSkuPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfiguration) *PublicIPAddressSku {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(PublicIPAddressSkuPtrOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings struct {
	DomainNameLabel string `pulumi:"domainNameLabel"`
}





type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput
	ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutputWithContext(context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput
}

type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsArgs struct {
	DomainNameLabel pulumi.StringInput `pulumi:"domainNameLabel"`
}

func (VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings)(nil)).Elem()
}

func (i VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsArgs) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput {
	return i.ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsArgs) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput)
}

func (i VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsArgs) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return i.ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsArgs) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput).ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput
	ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutputWithContext(context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput
}

type virtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrType VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsArgs

func VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtr(v *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsArgs) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrInput {
	return (*virtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrType)(v)
}

func (*virtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings)(nil)).Elem()
}

func (i *virtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrType) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return i.ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrType) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings)(nil)).Elem()
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return o.ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings) *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings {
		return &v
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput) DomainNameLabel() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings) string { return v.DomainNameLabel }).(pulumi.StringOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings)(nil)).Elem()
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput) Elem() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings
		return ret
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings) *string {
		if v == nil {
			return nil
		}
		return &v.DomainNameLabel
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse struct {
	DomainNameLabel string `pulumi:"domainNameLabel"`
}

type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput) DomainNameLabel() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse) string {
		return v.DomainNameLabel
	}).(pulumi.StringOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput) Elem() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse) VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse
		return ret
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DomainNameLabel
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfigurationResponse struct {
	DeleteOption           *string                                                                `pulumi:"deleteOption"`
	DnsSettings            *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse `pulumi:"dnsSettings"`
	IdleTimeoutInMinutes   *int                                                                   `pulumi:"idleTimeoutInMinutes"`
	IpTags                 []VirtualMachineScaleSetIpTagResponse                                  `pulumi:"ipTags"`
	Name                   string                                                                 `pulumi:"name"`
	PublicIPAddressVersion *string                                                                `pulumi:"publicIPAddressVersion"`
	PublicIPPrefix         *SubResourceResponse                                                   `pulumi:"publicIPPrefix"`
	Sku                    *PublicIPAddressSkuResponse                                            `pulumi:"sku"`
}

type VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetPublicIPAddressConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput() VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfigurationResponse) *string { return v.DeleteOption }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput) DnsSettings() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfigurationResponse) *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse {
		return v.DnsSettings
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfigurationResponse) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput) IpTags() VirtualMachineScaleSetIpTagResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfigurationResponse) []VirtualMachineScaleSetIpTagResponse {
		return v.IpTags
	}).(VirtualMachineScaleSetIpTagResponseArrayOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfigurationResponse) *string {
		return v.PublicIPAddressVersion
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput) PublicIPPrefix() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfigurationResponse) *SubResourceResponse {
		return v.PublicIPPrefix
	}).(SubResourceResponsePtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput) Sku() PublicIPAddressSkuResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetPublicIPAddressConfigurationResponse) *PublicIPAddressSkuResponse {
		return v.Sku
	}).(PublicIPAddressSkuResponsePtrOutput)
}

type VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetPublicIPAddressConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput() VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) ToVirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) Elem() VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationResponse) VirtualMachineScaleSetPublicIPAddressConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetPublicIPAddressConfigurationResponse
		return ret
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) DeleteOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeleteOption
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) DnsSettings() VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationResponse) *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponse {
		if v == nil {
			return nil
		}
		return v.DnsSettings
	}).(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationResponse) *int {
		if v == nil {
			return nil
		}
		return v.IdleTimeoutInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) IpTags() VirtualMachineScaleSetIpTagResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationResponse) []VirtualMachineScaleSetIpTagResponse {
		if v == nil {
			return nil
		}
		return v.IpTags
	}).(VirtualMachineScaleSetIpTagResponseArrayOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicIPAddressVersion
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) PublicIPPrefix() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return v.PublicIPPrefix
	}).(SubResourceResponsePtrOutput)
}

func (o VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput) Sku() PublicIPAddressSkuResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetPublicIPAddressConfigurationResponse) *PublicIPAddressSkuResponse {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(PublicIPAddressSkuResponsePtrOutput)
}

type VirtualMachineScaleSetStorageProfile struct {
	DataDisks      []VirtualMachineScaleSetDataDisk `pulumi:"dataDisks"`
	ImageReference *ImageReference                  `pulumi:"imageReference"`
	OsDisk         *VirtualMachineScaleSetOSDisk    `pulumi:"osDisk"`
}





type VirtualMachineScaleSetStorageProfileInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetStorageProfileOutput() VirtualMachineScaleSetStorageProfileOutput
	ToVirtualMachineScaleSetStorageProfileOutputWithContext(context.Context) VirtualMachineScaleSetStorageProfileOutput
}

type VirtualMachineScaleSetStorageProfileArgs struct {
	DataDisks      VirtualMachineScaleSetDataDiskArrayInput `pulumi:"dataDisks"`
	ImageReference ImageReferencePtrInput                   `pulumi:"imageReference"`
	OsDisk         VirtualMachineScaleSetOSDiskPtrInput     `pulumi:"osDisk"`
}

func (VirtualMachineScaleSetStorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetStorageProfile)(nil)).Elem()
}

func (i VirtualMachineScaleSetStorageProfileArgs) ToVirtualMachineScaleSetStorageProfileOutput() VirtualMachineScaleSetStorageProfileOutput {
	return i.ToVirtualMachineScaleSetStorageProfileOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetStorageProfileArgs) ToVirtualMachineScaleSetStorageProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetStorageProfileOutput)
}

func (i VirtualMachineScaleSetStorageProfileArgs) ToVirtualMachineScaleSetStorageProfilePtrOutput() VirtualMachineScaleSetStorageProfilePtrOutput {
	return i.ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetStorageProfileArgs) ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetStorageProfileOutput).ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetStorageProfilePtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetStorageProfilePtrOutput() VirtualMachineScaleSetStorageProfilePtrOutput
	ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(context.Context) VirtualMachineScaleSetStorageProfilePtrOutput
}

type virtualMachineScaleSetStorageProfilePtrType VirtualMachineScaleSetStorageProfileArgs

func VirtualMachineScaleSetStorageProfilePtr(v *VirtualMachineScaleSetStorageProfileArgs) VirtualMachineScaleSetStorageProfilePtrInput {
	return (*virtualMachineScaleSetStorageProfilePtrType)(v)
}

func (*virtualMachineScaleSetStorageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetStorageProfile)(nil)).Elem()
}

func (i *virtualMachineScaleSetStorageProfilePtrType) ToVirtualMachineScaleSetStorageProfilePtrOutput() VirtualMachineScaleSetStorageProfilePtrOutput {
	return i.ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetStorageProfilePtrType) ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetStorageProfilePtrOutput)
}

type VirtualMachineScaleSetStorageProfileOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetStorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetStorageProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetStorageProfileOutput) ToVirtualMachineScaleSetStorageProfileOutput() VirtualMachineScaleSetStorageProfileOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfileOutput) ToVirtualMachineScaleSetStorageProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfileOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfileOutput) ToVirtualMachineScaleSetStorageProfilePtrOutput() VirtualMachineScaleSetStorageProfilePtrOutput {
	return o.ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetStorageProfileOutput) ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetStorageProfile) *VirtualMachineScaleSetStorageProfile {
		return &v
	}).(VirtualMachineScaleSetStorageProfilePtrOutput)
}

func (o VirtualMachineScaleSetStorageProfileOutput) DataDisks() VirtualMachineScaleSetDataDiskArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetStorageProfile) []VirtualMachineScaleSetDataDisk { return v.DataDisks }).(VirtualMachineScaleSetDataDiskArrayOutput)
}

func (o VirtualMachineScaleSetStorageProfileOutput) ImageReference() ImageReferencePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetStorageProfile) *ImageReference { return v.ImageReference }).(ImageReferencePtrOutput)
}

func (o VirtualMachineScaleSetStorageProfileOutput) OsDisk() VirtualMachineScaleSetOSDiskPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetStorageProfile) *VirtualMachineScaleSetOSDisk { return v.OsDisk }).(VirtualMachineScaleSetOSDiskPtrOutput)
}

type VirtualMachineScaleSetStorageProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetStorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetStorageProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetStorageProfilePtrOutput) ToVirtualMachineScaleSetStorageProfilePtrOutput() VirtualMachineScaleSetStorageProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfilePtrOutput) ToVirtualMachineScaleSetStorageProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfilePtrOutput) Elem() VirtualMachineScaleSetStorageProfileOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfile) VirtualMachineScaleSetStorageProfile {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetStorageProfile
		return ret
	}).(VirtualMachineScaleSetStorageProfileOutput)
}

func (o VirtualMachineScaleSetStorageProfilePtrOutput) DataDisks() VirtualMachineScaleSetDataDiskArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfile) []VirtualMachineScaleSetDataDisk {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(VirtualMachineScaleSetDataDiskArrayOutput)
}

func (o VirtualMachineScaleSetStorageProfilePtrOutput) ImageReference() ImageReferencePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfile) *ImageReference {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(ImageReferencePtrOutput)
}

func (o VirtualMachineScaleSetStorageProfilePtrOutput) OsDisk() VirtualMachineScaleSetOSDiskPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfile) *VirtualMachineScaleSetOSDisk {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(VirtualMachineScaleSetOSDiskPtrOutput)
}

type VirtualMachineScaleSetStorageProfileResponse struct {
	DataDisks      []VirtualMachineScaleSetDataDiskResponse `pulumi:"dataDisks"`
	ImageReference *ImageReferenceResponse                  `pulumi:"imageReference"`
	OsDisk         *VirtualMachineScaleSetOSDiskResponse    `pulumi:"osDisk"`
}

type VirtualMachineScaleSetStorageProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetStorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetStorageProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetStorageProfileResponseOutput) ToVirtualMachineScaleSetStorageProfileResponseOutput() VirtualMachineScaleSetStorageProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfileResponseOutput) ToVirtualMachineScaleSetStorageProfileResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfileResponseOutput) DataDisks() VirtualMachineScaleSetDataDiskResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetStorageProfileResponse) []VirtualMachineScaleSetDataDiskResponse {
		return v.DataDisks
	}).(VirtualMachineScaleSetDataDiskResponseArrayOutput)
}

func (o VirtualMachineScaleSetStorageProfileResponseOutput) ImageReference() ImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetStorageProfileResponse) *ImageReferenceResponse { return v.ImageReference }).(ImageReferenceResponsePtrOutput)
}

func (o VirtualMachineScaleSetStorageProfileResponseOutput) OsDisk() VirtualMachineScaleSetOSDiskResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetStorageProfileResponse) *VirtualMachineScaleSetOSDiskResponse {
		return v.OsDisk
	}).(VirtualMachineScaleSetOSDiskResponsePtrOutput)
}

type VirtualMachineScaleSetStorageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetStorageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetStorageProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetStorageProfileResponsePtrOutput) ToVirtualMachineScaleSetStorageProfileResponsePtrOutput() VirtualMachineScaleSetStorageProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfileResponsePtrOutput) ToVirtualMachineScaleSetStorageProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetStorageProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetStorageProfileResponsePtrOutput) Elem() VirtualMachineScaleSetStorageProfileResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfileResponse) VirtualMachineScaleSetStorageProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetStorageProfileResponse
		return ret
	}).(VirtualMachineScaleSetStorageProfileResponseOutput)
}

func (o VirtualMachineScaleSetStorageProfileResponsePtrOutput) DataDisks() VirtualMachineScaleSetDataDiskResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfileResponse) []VirtualMachineScaleSetDataDiskResponse {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(VirtualMachineScaleSetDataDiskResponseArrayOutput)
}

func (o VirtualMachineScaleSetStorageProfileResponsePtrOutput) ImageReference() ImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfileResponse) *ImageReferenceResponse {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(ImageReferenceResponsePtrOutput)
}

func (o VirtualMachineScaleSetStorageProfileResponsePtrOutput) OsDisk() VirtualMachineScaleSetOSDiskResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetStorageProfileResponse) *VirtualMachineScaleSetOSDiskResponse {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(VirtualMachineScaleSetOSDiskResponsePtrOutput)
}

type VirtualMachineScaleSetVMInstanceViewResponse struct {
	AssignedHost              string                                        `pulumi:"assignedHost"`
	BootDiagnostics           *BootDiagnosticsInstanceViewResponse          `pulumi:"bootDiagnostics"`
	Disks                     []DiskInstanceViewResponse                    `pulumi:"disks"`
	Extensions                []VirtualMachineExtensionInstanceViewResponse `pulumi:"extensions"`
	MaintenanceRedeployStatus *MaintenanceRedeployStatusResponse            `pulumi:"maintenanceRedeployStatus"`
	PlacementGroupId          *string                                       `pulumi:"placementGroupId"`
	PlatformFaultDomain       *int                                          `pulumi:"platformFaultDomain"`
	PlatformUpdateDomain      *int                                          `pulumi:"platformUpdateDomain"`
	RdpThumbPrint             *string                                       `pulumi:"rdpThumbPrint"`
	Statuses                  []InstanceViewStatusResponse                  `pulumi:"statuses"`
	VmAgent                   *VirtualMachineAgentInstanceViewResponse      `pulumi:"vmAgent"`
	VmHealth                  VirtualMachineHealthStatusResponse            `pulumi:"vmHealth"`
}

type VirtualMachineScaleSetVMInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetVMInstanceViewResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) ToVirtualMachineScaleSetVMInstanceViewResponseOutput() VirtualMachineScaleSetVMInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) ToVirtualMachineScaleSetVMInstanceViewResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMInstanceViewResponseOutput {
	return o
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) AssignedHost() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) string { return v.AssignedHost }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) BootDiagnostics() BootDiagnosticsInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) *BootDiagnosticsInstanceViewResponse {
		return v.BootDiagnostics
	}).(BootDiagnosticsInstanceViewResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) Disks() DiskInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) []DiskInstanceViewResponse { return v.Disks }).(DiskInstanceViewResponseArrayOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) Extensions() VirtualMachineExtensionInstanceViewResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) []VirtualMachineExtensionInstanceViewResponse {
		return v.Extensions
	}).(VirtualMachineExtensionInstanceViewResponseArrayOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) MaintenanceRedeployStatus() MaintenanceRedeployStatusResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) *MaintenanceRedeployStatusResponse {
		return v.MaintenanceRedeployStatus
	}).(MaintenanceRedeployStatusResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) PlacementGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) *string { return v.PlacementGroupId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) PlatformFaultDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) *int { return v.PlatformFaultDomain }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) PlatformUpdateDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) *int { return v.PlatformUpdateDomain }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) RdpThumbPrint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) *string { return v.RdpThumbPrint }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) []InstanceViewStatusResponse { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) VmAgent() VirtualMachineAgentInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) *VirtualMachineAgentInstanceViewResponse {
		return v.VmAgent
	}).(VirtualMachineAgentInstanceViewResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMInstanceViewResponseOutput) VmHealth() VirtualMachineHealthStatusResponseOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMInstanceViewResponse) VirtualMachineHealthStatusResponse {
		return v.VmHealth
	}).(VirtualMachineHealthStatusResponseOutput)
}

type VirtualMachineScaleSetVMNetworkProfileConfiguration struct {
	NetworkInterfaceConfigurations []VirtualMachineScaleSetNetworkConfiguration `pulumi:"networkInterfaceConfigurations"`
}





type VirtualMachineScaleSetVMNetworkProfileConfigurationInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetVMNetworkProfileConfigurationOutput() VirtualMachineScaleSetVMNetworkProfileConfigurationOutput
	ToVirtualMachineScaleSetVMNetworkProfileConfigurationOutputWithContext(context.Context) VirtualMachineScaleSetVMNetworkProfileConfigurationOutput
}

type VirtualMachineScaleSetVMNetworkProfileConfigurationArgs struct {
	NetworkInterfaceConfigurations VirtualMachineScaleSetNetworkConfigurationArrayInput `pulumi:"networkInterfaceConfigurations"`
}

func (VirtualMachineScaleSetVMNetworkProfileConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetVMNetworkProfileConfiguration)(nil)).Elem()
}

func (i VirtualMachineScaleSetVMNetworkProfileConfigurationArgs) ToVirtualMachineScaleSetVMNetworkProfileConfigurationOutput() VirtualMachineScaleSetVMNetworkProfileConfigurationOutput {
	return i.ToVirtualMachineScaleSetVMNetworkProfileConfigurationOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetVMNetworkProfileConfigurationArgs) ToVirtualMachineScaleSetVMNetworkProfileConfigurationOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMNetworkProfileConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMNetworkProfileConfigurationOutput)
}

func (i VirtualMachineScaleSetVMNetworkProfileConfigurationArgs) ToVirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput() VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput {
	return i.ToVirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetVMNetworkProfileConfigurationArgs) ToVirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMNetworkProfileConfigurationOutput).ToVirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetVMNetworkProfileConfigurationPtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput() VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput
	ToVirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutputWithContext(context.Context) VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput
}

type virtualMachineScaleSetVMNetworkProfileConfigurationPtrType VirtualMachineScaleSetVMNetworkProfileConfigurationArgs

func VirtualMachineScaleSetVMNetworkProfileConfigurationPtr(v *VirtualMachineScaleSetVMNetworkProfileConfigurationArgs) VirtualMachineScaleSetVMNetworkProfileConfigurationPtrInput {
	return (*virtualMachineScaleSetVMNetworkProfileConfigurationPtrType)(v)
}

func (*virtualMachineScaleSetVMNetworkProfileConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMNetworkProfileConfiguration)(nil)).Elem()
}

func (i *virtualMachineScaleSetVMNetworkProfileConfigurationPtrType) ToVirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput() VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput {
	return i.ToVirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetVMNetworkProfileConfigurationPtrType) ToVirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput)
}

type VirtualMachineScaleSetVMNetworkProfileConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMNetworkProfileConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetVMNetworkProfileConfiguration)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMNetworkProfileConfigurationOutput) ToVirtualMachineScaleSetVMNetworkProfileConfigurationOutput() VirtualMachineScaleSetVMNetworkProfileConfigurationOutput {
	return o
}

func (o VirtualMachineScaleSetVMNetworkProfileConfigurationOutput) ToVirtualMachineScaleSetVMNetworkProfileConfigurationOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMNetworkProfileConfigurationOutput {
	return o
}

func (o VirtualMachineScaleSetVMNetworkProfileConfigurationOutput) ToVirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput() VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput {
	return o.ToVirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetVMNetworkProfileConfigurationOutput) ToVirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetVMNetworkProfileConfiguration) *VirtualMachineScaleSetVMNetworkProfileConfiguration {
		return &v
	}).(VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput)
}

func (o VirtualMachineScaleSetVMNetworkProfileConfigurationOutput) NetworkInterfaceConfigurations() VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMNetworkProfileConfiguration) []VirtualMachineScaleSetNetworkConfiguration {
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineScaleSetNetworkConfigurationArrayOutput)
}

type VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMNetworkProfileConfiguration)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput) ToVirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput() VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput) ToVirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput) Elem() VirtualMachineScaleSetVMNetworkProfileConfigurationOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMNetworkProfileConfiguration) VirtualMachineScaleSetVMNetworkProfileConfiguration {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetVMNetworkProfileConfiguration
		return ret
	}).(VirtualMachineScaleSetVMNetworkProfileConfigurationOutput)
}

func (o VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput) NetworkInterfaceConfigurations() VirtualMachineScaleSetNetworkConfigurationArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMNetworkProfileConfiguration) []VirtualMachineScaleSetNetworkConfiguration {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineScaleSetNetworkConfigurationArrayOutput)
}

type VirtualMachineScaleSetVMNetworkProfileConfigurationResponse struct {
	NetworkInterfaceConfigurations []VirtualMachineScaleSetNetworkConfigurationResponse `pulumi:"networkInterfaceConfigurations"`
}

type VirtualMachineScaleSetVMNetworkProfileConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMNetworkProfileConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetVMNetworkProfileConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMNetworkProfileConfigurationResponseOutput) ToVirtualMachineScaleSetVMNetworkProfileConfigurationResponseOutput() VirtualMachineScaleSetVMNetworkProfileConfigurationResponseOutput {
	return o
}

func (o VirtualMachineScaleSetVMNetworkProfileConfigurationResponseOutput) ToVirtualMachineScaleSetVMNetworkProfileConfigurationResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMNetworkProfileConfigurationResponseOutput {
	return o
}

func (o VirtualMachineScaleSetVMNetworkProfileConfigurationResponseOutput) NetworkInterfaceConfigurations() VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMNetworkProfileConfigurationResponse) []VirtualMachineScaleSetNetworkConfigurationResponse {
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput)
}

type VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMNetworkProfileConfigurationResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput) ToVirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput() VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput) ToVirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput) Elem() VirtualMachineScaleSetVMNetworkProfileConfigurationResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMNetworkProfileConfigurationResponse) VirtualMachineScaleSetVMNetworkProfileConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetVMNetworkProfileConfigurationResponse
		return ret
	}).(VirtualMachineScaleSetVMNetworkProfileConfigurationResponseOutput)
}

func (o VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput) NetworkInterfaceConfigurations() VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMNetworkProfileConfigurationResponse) []VirtualMachineScaleSetNetworkConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaceConfigurations
	}).(VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput)
}

type VirtualMachineScaleSetVMProfile struct {
	ApplicationProfile     *ApplicationProfile                     `pulumi:"applicationProfile"`
	BillingProfile         *BillingProfile                         `pulumi:"billingProfile"`
	CapacityReservation    *CapacityReservationProfile             `pulumi:"capacityReservation"`
	DiagnosticsProfile     *DiagnosticsProfile                     `pulumi:"diagnosticsProfile"`
	EvictionPolicy         *string                                 `pulumi:"evictionPolicy"`
	ExtensionProfile       *VirtualMachineScaleSetExtensionProfile `pulumi:"extensionProfile"`
	HardwareProfile        *VirtualMachineScaleSetHardwareProfile  `pulumi:"hardwareProfile"`
	LicenseType            *string                                 `pulumi:"licenseType"`
	NetworkProfile         *VirtualMachineScaleSetNetworkProfile   `pulumi:"networkProfile"`
	OsProfile              *VirtualMachineScaleSetOSProfile        `pulumi:"osProfile"`
	Priority               *string                                 `pulumi:"priority"`
	ScheduledEventsProfile *ScheduledEventsProfile                 `pulumi:"scheduledEventsProfile"`
	SecurityProfile        *SecurityProfile                        `pulumi:"securityProfile"`
	StorageProfile         *VirtualMachineScaleSetStorageProfile   `pulumi:"storageProfile"`
	UserData               *string                                 `pulumi:"userData"`
}





type VirtualMachineScaleSetVMProfileInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetVMProfileOutput() VirtualMachineScaleSetVMProfileOutput
	ToVirtualMachineScaleSetVMProfileOutputWithContext(context.Context) VirtualMachineScaleSetVMProfileOutput
}

type VirtualMachineScaleSetVMProfileArgs struct {
	ApplicationProfile     ApplicationProfilePtrInput                     `pulumi:"applicationProfile"`
	BillingProfile         BillingProfilePtrInput                         `pulumi:"billingProfile"`
	CapacityReservation    CapacityReservationProfilePtrInput             `pulumi:"capacityReservation"`
	DiagnosticsProfile     DiagnosticsProfilePtrInput                     `pulumi:"diagnosticsProfile"`
	EvictionPolicy         pulumi.StringPtrInput                          `pulumi:"evictionPolicy"`
	ExtensionProfile       VirtualMachineScaleSetExtensionProfilePtrInput `pulumi:"extensionProfile"`
	HardwareProfile        VirtualMachineScaleSetHardwareProfilePtrInput  `pulumi:"hardwareProfile"`
	LicenseType            pulumi.StringPtrInput                          `pulumi:"licenseType"`
	NetworkProfile         VirtualMachineScaleSetNetworkProfilePtrInput   `pulumi:"networkProfile"`
	OsProfile              VirtualMachineScaleSetOSProfilePtrInput        `pulumi:"osProfile"`
	Priority               pulumi.StringPtrInput                          `pulumi:"priority"`
	ScheduledEventsProfile ScheduledEventsProfilePtrInput                 `pulumi:"scheduledEventsProfile"`
	SecurityProfile        SecurityProfilePtrInput                        `pulumi:"securityProfile"`
	StorageProfile         VirtualMachineScaleSetStorageProfilePtrInput   `pulumi:"storageProfile"`
	UserData               pulumi.StringPtrInput                          `pulumi:"userData"`
}

func (VirtualMachineScaleSetVMProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetVMProfile)(nil)).Elem()
}

func (i VirtualMachineScaleSetVMProfileArgs) ToVirtualMachineScaleSetVMProfileOutput() VirtualMachineScaleSetVMProfileOutput {
	return i.ToVirtualMachineScaleSetVMProfileOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetVMProfileArgs) ToVirtualMachineScaleSetVMProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMProfileOutput)
}

func (i VirtualMachineScaleSetVMProfileArgs) ToVirtualMachineScaleSetVMProfilePtrOutput() VirtualMachineScaleSetVMProfilePtrOutput {
	return i.ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetVMProfileArgs) ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMProfileOutput).ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetVMProfilePtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetVMProfilePtrOutput() VirtualMachineScaleSetVMProfilePtrOutput
	ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(context.Context) VirtualMachineScaleSetVMProfilePtrOutput
}

type virtualMachineScaleSetVMProfilePtrType VirtualMachineScaleSetVMProfileArgs

func VirtualMachineScaleSetVMProfilePtr(v *VirtualMachineScaleSetVMProfileArgs) VirtualMachineScaleSetVMProfilePtrInput {
	return (*virtualMachineScaleSetVMProfilePtrType)(v)
}

func (*virtualMachineScaleSetVMProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMProfile)(nil)).Elem()
}

func (i *virtualMachineScaleSetVMProfilePtrType) ToVirtualMachineScaleSetVMProfilePtrOutput() VirtualMachineScaleSetVMProfilePtrOutput {
	return i.ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetVMProfilePtrType) ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMProfilePtrOutput)
}

type VirtualMachineScaleSetVMProfileOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetVMProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMProfileOutput) ToVirtualMachineScaleSetVMProfileOutput() VirtualMachineScaleSetVMProfileOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfileOutput) ToVirtualMachineScaleSetVMProfileOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfileOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfileOutput) ToVirtualMachineScaleSetVMProfilePtrOutput() VirtualMachineScaleSetVMProfilePtrOutput {
	return o.ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetVMProfileOutput) ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetVMProfile {
		return &v
	}).(VirtualMachineScaleSetVMProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) ApplicationProfile() ApplicationProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *ApplicationProfile { return v.ApplicationProfile }).(ApplicationProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) BillingProfile() BillingProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *BillingProfile { return v.BillingProfile }).(BillingProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) CapacityReservation() CapacityReservationProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *CapacityReservationProfile { return v.CapacityReservation }).(CapacityReservationProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) DiagnosticsProfile() DiagnosticsProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *DiagnosticsProfile { return v.DiagnosticsProfile }).(DiagnosticsProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) EvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *string { return v.EvictionPolicy }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) ExtensionProfile() VirtualMachineScaleSetExtensionProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetExtensionProfile {
		return v.ExtensionProfile
	}).(VirtualMachineScaleSetExtensionProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) HardwareProfile() VirtualMachineScaleSetHardwareProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetHardwareProfile {
		return v.HardwareProfile
	}).(VirtualMachineScaleSetHardwareProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) NetworkProfile() VirtualMachineScaleSetNetworkProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetNetworkProfile { return v.NetworkProfile }).(VirtualMachineScaleSetNetworkProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) OsProfile() VirtualMachineScaleSetOSProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetOSProfile { return v.OsProfile }).(VirtualMachineScaleSetOSProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *string { return v.Priority }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) ScheduledEventsProfile() ScheduledEventsProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *ScheduledEventsProfile { return v.ScheduledEventsProfile }).(ScheduledEventsProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) SecurityProfile() SecurityProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *SecurityProfile { return v.SecurityProfile }).(SecurityProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) StorageProfile() VirtualMachineScaleSetStorageProfilePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetStorageProfile { return v.StorageProfile }).(VirtualMachineScaleSetStorageProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfile) *string { return v.UserData }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetVMProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMProfile)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) ToVirtualMachineScaleSetVMProfilePtrOutput() VirtualMachineScaleSetVMProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) ToVirtualMachineScaleSetVMProfilePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfilePtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) Elem() VirtualMachineScaleSetVMProfileOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) VirtualMachineScaleSetVMProfile {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetVMProfile
		return ret
	}).(VirtualMachineScaleSetVMProfileOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) ApplicationProfile() ApplicationProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *ApplicationProfile {
		if v == nil {
			return nil
		}
		return v.ApplicationProfile
	}).(ApplicationProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) BillingProfile() BillingProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *BillingProfile {
		if v == nil {
			return nil
		}
		return v.BillingProfile
	}).(BillingProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) CapacityReservation() CapacityReservationProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *CapacityReservationProfile {
		if v == nil {
			return nil
		}
		return v.CapacityReservation
	}).(CapacityReservationProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) DiagnosticsProfile() DiagnosticsProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *DiagnosticsProfile {
		if v == nil {
			return nil
		}
		return v.DiagnosticsProfile
	}).(DiagnosticsProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) EvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *string {
		if v == nil {
			return nil
		}
		return v.EvictionPolicy
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) ExtensionProfile() VirtualMachineScaleSetExtensionProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetExtensionProfile {
		if v == nil {
			return nil
		}
		return v.ExtensionProfile
	}).(VirtualMachineScaleSetExtensionProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) HardwareProfile() VirtualMachineScaleSetHardwareProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetHardwareProfile {
		if v == nil {
			return nil
		}
		return v.HardwareProfile
	}).(VirtualMachineScaleSetHardwareProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *string {
		if v == nil {
			return nil
		}
		return v.LicenseType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) NetworkProfile() VirtualMachineScaleSetNetworkProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetNetworkProfile {
		if v == nil {
			return nil
		}
		return v.NetworkProfile
	}).(VirtualMachineScaleSetNetworkProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) OsProfile() VirtualMachineScaleSetOSProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetOSProfile {
		if v == nil {
			return nil
		}
		return v.OsProfile
	}).(VirtualMachineScaleSetOSProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *string {
		if v == nil {
			return nil
		}
		return v.Priority
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) ScheduledEventsProfile() ScheduledEventsProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *ScheduledEventsProfile {
		if v == nil {
			return nil
		}
		return v.ScheduledEventsProfile
	}).(ScheduledEventsProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) SecurityProfile() SecurityProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *SecurityProfile {
		if v == nil {
			return nil
		}
		return v.SecurityProfile
	}).(SecurityProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) StorageProfile() VirtualMachineScaleSetStorageProfilePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *VirtualMachineScaleSetStorageProfile {
		if v == nil {
			return nil
		}
		return v.StorageProfile
	}).(VirtualMachineScaleSetStorageProfilePtrOutput)
}

func (o VirtualMachineScaleSetVMProfilePtrOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfile) *string {
		if v == nil {
			return nil
		}
		return v.UserData
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetVMProfileResponse struct {
	ApplicationProfile     *ApplicationProfileResponse                     `pulumi:"applicationProfile"`
	BillingProfile         *BillingProfileResponse                         `pulumi:"billingProfile"`
	CapacityReservation    *CapacityReservationProfileResponse             `pulumi:"capacityReservation"`
	DiagnosticsProfile     *DiagnosticsProfileResponse                     `pulumi:"diagnosticsProfile"`
	EvictionPolicy         *string                                         `pulumi:"evictionPolicy"`
	ExtensionProfile       *VirtualMachineScaleSetExtensionProfileResponse `pulumi:"extensionProfile"`
	HardwareProfile        *VirtualMachineScaleSetHardwareProfileResponse  `pulumi:"hardwareProfile"`
	LicenseType            *string                                         `pulumi:"licenseType"`
	NetworkProfile         *VirtualMachineScaleSetNetworkProfileResponse   `pulumi:"networkProfile"`
	OsProfile              *VirtualMachineScaleSetOSProfileResponse        `pulumi:"osProfile"`
	Priority               *string                                         `pulumi:"priority"`
	ScheduledEventsProfile *ScheduledEventsProfileResponse                 `pulumi:"scheduledEventsProfile"`
	SecurityProfile        *SecurityProfileResponse                        `pulumi:"securityProfile"`
	StorageProfile         *VirtualMachineScaleSetStorageProfileResponse   `pulumi:"storageProfile"`
	UserData               *string                                         `pulumi:"userData"`
}

type VirtualMachineScaleSetVMProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetVMProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) ToVirtualMachineScaleSetVMProfileResponseOutput() VirtualMachineScaleSetVMProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) ToVirtualMachineScaleSetVMProfileResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfileResponseOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) ApplicationProfile() ApplicationProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *ApplicationProfileResponse {
		return v.ApplicationProfile
	}).(ApplicationProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) BillingProfile() BillingProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *BillingProfileResponse { return v.BillingProfile }).(BillingProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) CapacityReservation() CapacityReservationProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *CapacityReservationProfileResponse {
		return v.CapacityReservation
	}).(CapacityReservationProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) DiagnosticsProfile() DiagnosticsProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *DiagnosticsProfileResponse {
		return v.DiagnosticsProfile
	}).(DiagnosticsProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) EvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *string { return v.EvictionPolicy }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) ExtensionProfile() VirtualMachineScaleSetExtensionProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetExtensionProfileResponse {
		return v.ExtensionProfile
	}).(VirtualMachineScaleSetExtensionProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) HardwareProfile() VirtualMachineScaleSetHardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetHardwareProfileResponse {
		return v.HardwareProfile
	}).(VirtualMachineScaleSetHardwareProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) NetworkProfile() VirtualMachineScaleSetNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetNetworkProfileResponse {
		return v.NetworkProfile
	}).(VirtualMachineScaleSetNetworkProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) OsProfile() VirtualMachineScaleSetOSProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetOSProfileResponse {
		return v.OsProfile
	}).(VirtualMachineScaleSetOSProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *string { return v.Priority }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) ScheduledEventsProfile() ScheduledEventsProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *ScheduledEventsProfileResponse {
		return v.ScheduledEventsProfile
	}).(ScheduledEventsProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) SecurityProfile() SecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *SecurityProfileResponse { return v.SecurityProfile }).(SecurityProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) StorageProfile() VirtualMachineScaleSetStorageProfileResponsePtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetStorageProfileResponse {
		return v.StorageProfile
	}).(VirtualMachineScaleSetStorageProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponseOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProfileResponse) *string { return v.UserData }).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetVMProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMProfileResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) ToVirtualMachineScaleSetVMProfileResponsePtrOutput() VirtualMachineScaleSetVMProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) ToVirtualMachineScaleSetVMProfileResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProfileResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) Elem() VirtualMachineScaleSetVMProfileResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) VirtualMachineScaleSetVMProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetVMProfileResponse
		return ret
	}).(VirtualMachineScaleSetVMProfileResponseOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) ApplicationProfile() ApplicationProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *ApplicationProfileResponse {
		if v == nil {
			return nil
		}
		return v.ApplicationProfile
	}).(ApplicationProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) BillingProfile() BillingProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *BillingProfileResponse {
		if v == nil {
			return nil
		}
		return v.BillingProfile
	}).(BillingProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) CapacityReservation() CapacityReservationProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *CapacityReservationProfileResponse {
		if v == nil {
			return nil
		}
		return v.CapacityReservation
	}).(CapacityReservationProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) DiagnosticsProfile() DiagnosticsProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *DiagnosticsProfileResponse {
		if v == nil {
			return nil
		}
		return v.DiagnosticsProfile
	}).(DiagnosticsProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) EvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.EvictionPolicy
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) ExtensionProfile() VirtualMachineScaleSetExtensionProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetExtensionProfileResponse {
		if v == nil {
			return nil
		}
		return v.ExtensionProfile
	}).(VirtualMachineScaleSetExtensionProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) HardwareProfile() VirtualMachineScaleSetHardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetHardwareProfileResponse {
		if v == nil {
			return nil
		}
		return v.HardwareProfile
	}).(VirtualMachineScaleSetHardwareProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.LicenseType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) NetworkProfile() VirtualMachineScaleSetNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetNetworkProfileResponse {
		if v == nil {
			return nil
		}
		return v.NetworkProfile
	}).(VirtualMachineScaleSetNetworkProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) OsProfile() VirtualMachineScaleSetOSProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetOSProfileResponse {
		if v == nil {
			return nil
		}
		return v.OsProfile
	}).(VirtualMachineScaleSetOSProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Priority
	}).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) ScheduledEventsProfile() ScheduledEventsProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *ScheduledEventsProfileResponse {
		if v == nil {
			return nil
		}
		return v.ScheduledEventsProfile
	}).(ScheduledEventsProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) SecurityProfile() SecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *SecurityProfileResponse {
		if v == nil {
			return nil
		}
		return v.SecurityProfile
	}).(SecurityProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) StorageProfile() VirtualMachineScaleSetStorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *VirtualMachineScaleSetStorageProfileResponse {
		if v == nil {
			return nil
		}
		return v.StorageProfile
	}).(VirtualMachineScaleSetStorageProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMProfileResponsePtrOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserData
	}).(pulumi.StringPtrOutput)
}

type VirtualMachineScaleSetVMProtectionPolicy struct {
	ProtectFromScaleIn         *bool `pulumi:"protectFromScaleIn"`
	ProtectFromScaleSetActions *bool `pulumi:"protectFromScaleSetActions"`
}





type VirtualMachineScaleSetVMProtectionPolicyInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetVMProtectionPolicyOutput() VirtualMachineScaleSetVMProtectionPolicyOutput
	ToVirtualMachineScaleSetVMProtectionPolicyOutputWithContext(context.Context) VirtualMachineScaleSetVMProtectionPolicyOutput
}

type VirtualMachineScaleSetVMProtectionPolicyArgs struct {
	ProtectFromScaleIn         pulumi.BoolPtrInput `pulumi:"protectFromScaleIn"`
	ProtectFromScaleSetActions pulumi.BoolPtrInput `pulumi:"protectFromScaleSetActions"`
}

func (VirtualMachineScaleSetVMProtectionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetVMProtectionPolicy)(nil)).Elem()
}

func (i VirtualMachineScaleSetVMProtectionPolicyArgs) ToVirtualMachineScaleSetVMProtectionPolicyOutput() VirtualMachineScaleSetVMProtectionPolicyOutput {
	return i.ToVirtualMachineScaleSetVMProtectionPolicyOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetVMProtectionPolicyArgs) ToVirtualMachineScaleSetVMProtectionPolicyOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProtectionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMProtectionPolicyOutput)
}

func (i VirtualMachineScaleSetVMProtectionPolicyArgs) ToVirtualMachineScaleSetVMProtectionPolicyPtrOutput() VirtualMachineScaleSetVMProtectionPolicyPtrOutput {
	return i.ToVirtualMachineScaleSetVMProtectionPolicyPtrOutputWithContext(context.Background())
}

func (i VirtualMachineScaleSetVMProtectionPolicyArgs) ToVirtualMachineScaleSetVMProtectionPolicyPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProtectionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMProtectionPolicyOutput).ToVirtualMachineScaleSetVMProtectionPolicyPtrOutputWithContext(ctx)
}









type VirtualMachineScaleSetVMProtectionPolicyPtrInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetVMProtectionPolicyPtrOutput() VirtualMachineScaleSetVMProtectionPolicyPtrOutput
	ToVirtualMachineScaleSetVMProtectionPolicyPtrOutputWithContext(context.Context) VirtualMachineScaleSetVMProtectionPolicyPtrOutput
}

type virtualMachineScaleSetVMProtectionPolicyPtrType VirtualMachineScaleSetVMProtectionPolicyArgs

func VirtualMachineScaleSetVMProtectionPolicyPtr(v *VirtualMachineScaleSetVMProtectionPolicyArgs) VirtualMachineScaleSetVMProtectionPolicyPtrInput {
	return (*virtualMachineScaleSetVMProtectionPolicyPtrType)(v)
}

func (*virtualMachineScaleSetVMProtectionPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMProtectionPolicy)(nil)).Elem()
}

func (i *virtualMachineScaleSetVMProtectionPolicyPtrType) ToVirtualMachineScaleSetVMProtectionPolicyPtrOutput() VirtualMachineScaleSetVMProtectionPolicyPtrOutput {
	return i.ToVirtualMachineScaleSetVMProtectionPolicyPtrOutputWithContext(context.Background())
}

func (i *virtualMachineScaleSetVMProtectionPolicyPtrType) ToVirtualMachineScaleSetVMProtectionPolicyPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProtectionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMProtectionPolicyPtrOutput)
}

type VirtualMachineScaleSetVMProtectionPolicyOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMProtectionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetVMProtectionPolicy)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMProtectionPolicyOutput) ToVirtualMachineScaleSetVMProtectionPolicyOutput() VirtualMachineScaleSetVMProtectionPolicyOutput {
	return o
}

func (o VirtualMachineScaleSetVMProtectionPolicyOutput) ToVirtualMachineScaleSetVMProtectionPolicyOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProtectionPolicyOutput {
	return o
}

func (o VirtualMachineScaleSetVMProtectionPolicyOutput) ToVirtualMachineScaleSetVMProtectionPolicyPtrOutput() VirtualMachineScaleSetVMProtectionPolicyPtrOutput {
	return o.ToVirtualMachineScaleSetVMProtectionPolicyPtrOutputWithContext(context.Background())
}

func (o VirtualMachineScaleSetVMProtectionPolicyOutput) ToVirtualMachineScaleSetVMProtectionPolicyPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProtectionPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualMachineScaleSetVMProtectionPolicy) *VirtualMachineScaleSetVMProtectionPolicy {
		return &v
	}).(VirtualMachineScaleSetVMProtectionPolicyPtrOutput)
}

func (o VirtualMachineScaleSetVMProtectionPolicyOutput) ProtectFromScaleIn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProtectionPolicy) *bool { return v.ProtectFromScaleIn }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetVMProtectionPolicyOutput) ProtectFromScaleSetActions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProtectionPolicy) *bool { return v.ProtectFromScaleSetActions }).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetVMProtectionPolicyPtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMProtectionPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMProtectionPolicy)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMProtectionPolicyPtrOutput) ToVirtualMachineScaleSetVMProtectionPolicyPtrOutput() VirtualMachineScaleSetVMProtectionPolicyPtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMProtectionPolicyPtrOutput) ToVirtualMachineScaleSetVMProtectionPolicyPtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProtectionPolicyPtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMProtectionPolicyPtrOutput) Elem() VirtualMachineScaleSetVMProtectionPolicyOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProtectionPolicy) VirtualMachineScaleSetVMProtectionPolicy {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetVMProtectionPolicy
		return ret
	}).(VirtualMachineScaleSetVMProtectionPolicyOutput)
}

func (o VirtualMachineScaleSetVMProtectionPolicyPtrOutput) ProtectFromScaleIn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProtectionPolicy) *bool {
		if v == nil {
			return nil
		}
		return v.ProtectFromScaleIn
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetVMProtectionPolicyPtrOutput) ProtectFromScaleSetActions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProtectionPolicy) *bool {
		if v == nil {
			return nil
		}
		return v.ProtectFromScaleSetActions
	}).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetVMProtectionPolicyResponse struct {
	ProtectFromScaleIn         *bool `pulumi:"protectFromScaleIn"`
	ProtectFromScaleSetActions *bool `pulumi:"protectFromScaleSetActions"`
}

type VirtualMachineScaleSetVMProtectionPolicyResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMProtectionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSetVMProtectionPolicyResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMProtectionPolicyResponseOutput) ToVirtualMachineScaleSetVMProtectionPolicyResponseOutput() VirtualMachineScaleSetVMProtectionPolicyResponseOutput {
	return o
}

func (o VirtualMachineScaleSetVMProtectionPolicyResponseOutput) ToVirtualMachineScaleSetVMProtectionPolicyResponseOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProtectionPolicyResponseOutput {
	return o
}

func (o VirtualMachineScaleSetVMProtectionPolicyResponseOutput) ProtectFromScaleIn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProtectionPolicyResponse) *bool { return v.ProtectFromScaleIn }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetVMProtectionPolicyResponseOutput) ProtectFromScaleSetActions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualMachineScaleSetVMProtectionPolicyResponse) *bool { return v.ProtectFromScaleSetActions }).(pulumi.BoolPtrOutput)
}

type VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMProtectionPolicyResponse)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput) ToVirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput() VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput) ToVirtualMachineScaleSetVMProtectionPolicyResponsePtrOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput {
	return o
}

func (o VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput) Elem() VirtualMachineScaleSetVMProtectionPolicyResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProtectionPolicyResponse) VirtualMachineScaleSetVMProtectionPolicyResponse {
		if v != nil {
			return *v
		}
		var ret VirtualMachineScaleSetVMProtectionPolicyResponse
		return ret
	}).(VirtualMachineScaleSetVMProtectionPolicyResponseOutput)
}

func (o VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput) ProtectFromScaleIn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProtectionPolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ProtectFromScaleIn
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput) ProtectFromScaleSetActions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMProtectionPolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ProtectFromScaleSetActions
	}).(pulumi.BoolPtrOutput)
}

type WinRMConfiguration struct {
	Listeners []WinRMListener `pulumi:"listeners"`
}





type WinRMConfigurationInput interface {
	pulumi.Input

	ToWinRMConfigurationOutput() WinRMConfigurationOutput
	ToWinRMConfigurationOutputWithContext(context.Context) WinRMConfigurationOutput
}

type WinRMConfigurationArgs struct {
	Listeners WinRMListenerArrayInput `pulumi:"listeners"`
}

func (WinRMConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WinRMConfiguration)(nil)).Elem()
}

func (i WinRMConfigurationArgs) ToWinRMConfigurationOutput() WinRMConfigurationOutput {
	return i.ToWinRMConfigurationOutputWithContext(context.Background())
}

func (i WinRMConfigurationArgs) ToWinRMConfigurationOutputWithContext(ctx context.Context) WinRMConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WinRMConfigurationOutput)
}

func (i WinRMConfigurationArgs) ToWinRMConfigurationPtrOutput() WinRMConfigurationPtrOutput {
	return i.ToWinRMConfigurationPtrOutputWithContext(context.Background())
}

func (i WinRMConfigurationArgs) ToWinRMConfigurationPtrOutputWithContext(ctx context.Context) WinRMConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WinRMConfigurationOutput).ToWinRMConfigurationPtrOutputWithContext(ctx)
}









type WinRMConfigurationPtrInput interface {
	pulumi.Input

	ToWinRMConfigurationPtrOutput() WinRMConfigurationPtrOutput
	ToWinRMConfigurationPtrOutputWithContext(context.Context) WinRMConfigurationPtrOutput
}

type winRMConfigurationPtrType WinRMConfigurationArgs

func WinRMConfigurationPtr(v *WinRMConfigurationArgs) WinRMConfigurationPtrInput {
	return (*winRMConfigurationPtrType)(v)
}

func (*winRMConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WinRMConfiguration)(nil)).Elem()
}

func (i *winRMConfigurationPtrType) ToWinRMConfigurationPtrOutput() WinRMConfigurationPtrOutput {
	return i.ToWinRMConfigurationPtrOutputWithContext(context.Background())
}

func (i *winRMConfigurationPtrType) ToWinRMConfigurationPtrOutputWithContext(ctx context.Context) WinRMConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WinRMConfigurationPtrOutput)
}

type WinRMConfigurationOutput struct{ *pulumi.OutputState }

func (WinRMConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WinRMConfiguration)(nil)).Elem()
}

func (o WinRMConfigurationOutput) ToWinRMConfigurationOutput() WinRMConfigurationOutput {
	return o
}

func (o WinRMConfigurationOutput) ToWinRMConfigurationOutputWithContext(ctx context.Context) WinRMConfigurationOutput {
	return o
}

func (o WinRMConfigurationOutput) ToWinRMConfigurationPtrOutput() WinRMConfigurationPtrOutput {
	return o.ToWinRMConfigurationPtrOutputWithContext(context.Background())
}

func (o WinRMConfigurationOutput) ToWinRMConfigurationPtrOutputWithContext(ctx context.Context) WinRMConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WinRMConfiguration) *WinRMConfiguration {
		return &v
	}).(WinRMConfigurationPtrOutput)
}

func (o WinRMConfigurationOutput) Listeners() WinRMListenerArrayOutput {
	return o.ApplyT(func(v WinRMConfiguration) []WinRMListener { return v.Listeners }).(WinRMListenerArrayOutput)
}

type WinRMConfigurationPtrOutput struct{ *pulumi.OutputState }

func (WinRMConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WinRMConfiguration)(nil)).Elem()
}

func (o WinRMConfigurationPtrOutput) ToWinRMConfigurationPtrOutput() WinRMConfigurationPtrOutput {
	return o
}

func (o WinRMConfigurationPtrOutput) ToWinRMConfigurationPtrOutputWithContext(ctx context.Context) WinRMConfigurationPtrOutput {
	return o
}

func (o WinRMConfigurationPtrOutput) Elem() WinRMConfigurationOutput {
	return o.ApplyT(func(v *WinRMConfiguration) WinRMConfiguration {
		if v != nil {
			return *v
		}
		var ret WinRMConfiguration
		return ret
	}).(WinRMConfigurationOutput)
}

func (o WinRMConfigurationPtrOutput) Listeners() WinRMListenerArrayOutput {
	return o.ApplyT(func(v *WinRMConfiguration) []WinRMListener {
		if v == nil {
			return nil
		}
		return v.Listeners
	}).(WinRMListenerArrayOutput)
}

type WinRMConfigurationResponse struct {
	Listeners []WinRMListenerResponse `pulumi:"listeners"`
}

type WinRMConfigurationResponseOutput struct{ *pulumi.OutputState }

func (WinRMConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WinRMConfigurationResponse)(nil)).Elem()
}

func (o WinRMConfigurationResponseOutput) ToWinRMConfigurationResponseOutput() WinRMConfigurationResponseOutput {
	return o
}

func (o WinRMConfigurationResponseOutput) ToWinRMConfigurationResponseOutputWithContext(ctx context.Context) WinRMConfigurationResponseOutput {
	return o
}

func (o WinRMConfigurationResponseOutput) Listeners() WinRMListenerResponseArrayOutput {
	return o.ApplyT(func(v WinRMConfigurationResponse) []WinRMListenerResponse { return v.Listeners }).(WinRMListenerResponseArrayOutput)
}

type WinRMConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (WinRMConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WinRMConfigurationResponse)(nil)).Elem()
}

func (o WinRMConfigurationResponsePtrOutput) ToWinRMConfigurationResponsePtrOutput() WinRMConfigurationResponsePtrOutput {
	return o
}

func (o WinRMConfigurationResponsePtrOutput) ToWinRMConfigurationResponsePtrOutputWithContext(ctx context.Context) WinRMConfigurationResponsePtrOutput {
	return o
}

func (o WinRMConfigurationResponsePtrOutput) Elem() WinRMConfigurationResponseOutput {
	return o.ApplyT(func(v *WinRMConfigurationResponse) WinRMConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret WinRMConfigurationResponse
		return ret
	}).(WinRMConfigurationResponseOutput)
}

func (o WinRMConfigurationResponsePtrOutput) Listeners() WinRMListenerResponseArrayOutput {
	return o.ApplyT(func(v *WinRMConfigurationResponse) []WinRMListenerResponse {
		if v == nil {
			return nil
		}
		return v.Listeners
	}).(WinRMListenerResponseArrayOutput)
}

type WinRMListener struct {
	CertificateUrl *string        `pulumi:"certificateUrl"`
	Protocol       *ProtocolTypes `pulumi:"protocol"`
}





type WinRMListenerInput interface {
	pulumi.Input

	ToWinRMListenerOutput() WinRMListenerOutput
	ToWinRMListenerOutputWithContext(context.Context) WinRMListenerOutput
}

type WinRMListenerArgs struct {
	CertificateUrl pulumi.StringPtrInput `pulumi:"certificateUrl"`
	Protocol       ProtocolTypesPtrInput `pulumi:"protocol"`
}

func (WinRMListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WinRMListener)(nil)).Elem()
}

func (i WinRMListenerArgs) ToWinRMListenerOutput() WinRMListenerOutput {
	return i.ToWinRMListenerOutputWithContext(context.Background())
}

func (i WinRMListenerArgs) ToWinRMListenerOutputWithContext(ctx context.Context) WinRMListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WinRMListenerOutput)
}





type WinRMListenerArrayInput interface {
	pulumi.Input

	ToWinRMListenerArrayOutput() WinRMListenerArrayOutput
	ToWinRMListenerArrayOutputWithContext(context.Context) WinRMListenerArrayOutput
}

type WinRMListenerArray []WinRMListenerInput

func (WinRMListenerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WinRMListener)(nil)).Elem()
}

func (i WinRMListenerArray) ToWinRMListenerArrayOutput() WinRMListenerArrayOutput {
	return i.ToWinRMListenerArrayOutputWithContext(context.Background())
}

func (i WinRMListenerArray) ToWinRMListenerArrayOutputWithContext(ctx context.Context) WinRMListenerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WinRMListenerArrayOutput)
}

type WinRMListenerOutput struct{ *pulumi.OutputState }

func (WinRMListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WinRMListener)(nil)).Elem()
}

func (o WinRMListenerOutput) ToWinRMListenerOutput() WinRMListenerOutput {
	return o
}

func (o WinRMListenerOutput) ToWinRMListenerOutputWithContext(ctx context.Context) WinRMListenerOutput {
	return o
}

func (o WinRMListenerOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WinRMListener) *string { return v.CertificateUrl }).(pulumi.StringPtrOutput)
}

func (o WinRMListenerOutput) Protocol() ProtocolTypesPtrOutput {
	return o.ApplyT(func(v WinRMListener) *ProtocolTypes { return v.Protocol }).(ProtocolTypesPtrOutput)
}

type WinRMListenerArrayOutput struct{ *pulumi.OutputState }

func (WinRMListenerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WinRMListener)(nil)).Elem()
}

func (o WinRMListenerArrayOutput) ToWinRMListenerArrayOutput() WinRMListenerArrayOutput {
	return o
}

func (o WinRMListenerArrayOutput) ToWinRMListenerArrayOutputWithContext(ctx context.Context) WinRMListenerArrayOutput {
	return o
}

func (o WinRMListenerArrayOutput) Index(i pulumi.IntInput) WinRMListenerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WinRMListener {
		return vs[0].([]WinRMListener)[vs[1].(int)]
	}).(WinRMListenerOutput)
}

type WinRMListenerResponse struct {
	CertificateUrl *string `pulumi:"certificateUrl"`
	Protocol       *string `pulumi:"protocol"`
}

type WinRMListenerResponseOutput struct{ *pulumi.OutputState }

func (WinRMListenerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WinRMListenerResponse)(nil)).Elem()
}

func (o WinRMListenerResponseOutput) ToWinRMListenerResponseOutput() WinRMListenerResponseOutput {
	return o
}

func (o WinRMListenerResponseOutput) ToWinRMListenerResponseOutputWithContext(ctx context.Context) WinRMListenerResponseOutput {
	return o
}

func (o WinRMListenerResponseOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WinRMListenerResponse) *string { return v.CertificateUrl }).(pulumi.StringPtrOutput)
}

func (o WinRMListenerResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WinRMListenerResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

type WinRMListenerResponseArrayOutput struct{ *pulumi.OutputState }

func (WinRMListenerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WinRMListenerResponse)(nil)).Elem()
}

func (o WinRMListenerResponseArrayOutput) ToWinRMListenerResponseArrayOutput() WinRMListenerResponseArrayOutput {
	return o
}

func (o WinRMListenerResponseArrayOutput) ToWinRMListenerResponseArrayOutputWithContext(ctx context.Context) WinRMListenerResponseArrayOutput {
	return o
}

func (o WinRMListenerResponseArrayOutput) Index(i pulumi.IntInput) WinRMListenerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WinRMListenerResponse {
		return vs[0].([]WinRMListenerResponse)[vs[1].(int)]
	}).(WinRMListenerResponseOutput)
}

type WindowsConfiguration struct {
	AdditionalUnattendContent []AdditionalUnattendContent `pulumi:"additionalUnattendContent"`
	EnableAutomaticUpdates    *bool                       `pulumi:"enableAutomaticUpdates"`
	PatchSettings             *PatchSettings              `pulumi:"patchSettings"`
	ProvisionVMAgent          *bool                       `pulumi:"provisionVMAgent"`
	TimeZone                  *string                     `pulumi:"timeZone"`
	WinRM                     *WinRMConfiguration         `pulumi:"winRM"`
}





type WindowsConfigurationInput interface {
	pulumi.Input

	ToWindowsConfigurationOutput() WindowsConfigurationOutput
	ToWindowsConfigurationOutputWithContext(context.Context) WindowsConfigurationOutput
}

type WindowsConfigurationArgs struct {
	AdditionalUnattendContent AdditionalUnattendContentArrayInput `pulumi:"additionalUnattendContent"`
	EnableAutomaticUpdates    pulumi.BoolPtrInput                 `pulumi:"enableAutomaticUpdates"`
	PatchSettings             PatchSettingsPtrInput               `pulumi:"patchSettings"`
	ProvisionVMAgent          pulumi.BoolPtrInput                 `pulumi:"provisionVMAgent"`
	TimeZone                  pulumi.StringPtrInput               `pulumi:"timeZone"`
	WinRM                     WinRMConfigurationPtrInput          `pulumi:"winRM"`
}

func (WindowsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsConfiguration)(nil)).Elem()
}

func (i WindowsConfigurationArgs) ToWindowsConfigurationOutput() WindowsConfigurationOutput {
	return i.ToWindowsConfigurationOutputWithContext(context.Background())
}

func (i WindowsConfigurationArgs) ToWindowsConfigurationOutputWithContext(ctx context.Context) WindowsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsConfigurationOutput)
}

func (i WindowsConfigurationArgs) ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput {
	return i.ToWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (i WindowsConfigurationArgs) ToWindowsConfigurationPtrOutputWithContext(ctx context.Context) WindowsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsConfigurationOutput).ToWindowsConfigurationPtrOutputWithContext(ctx)
}









type WindowsConfigurationPtrInput interface {
	pulumi.Input

	ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput
	ToWindowsConfigurationPtrOutputWithContext(context.Context) WindowsConfigurationPtrOutput
}

type windowsConfigurationPtrType WindowsConfigurationArgs

func WindowsConfigurationPtr(v *WindowsConfigurationArgs) WindowsConfigurationPtrInput {
	return (*windowsConfigurationPtrType)(v)
}

func (*windowsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsConfiguration)(nil)).Elem()
}

func (i *windowsConfigurationPtrType) ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput {
	return i.ToWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (i *windowsConfigurationPtrType) ToWindowsConfigurationPtrOutputWithContext(ctx context.Context) WindowsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsConfigurationPtrOutput)
}

type WindowsConfigurationOutput struct{ *pulumi.OutputState }

func (WindowsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsConfiguration)(nil)).Elem()
}

func (o WindowsConfigurationOutput) ToWindowsConfigurationOutput() WindowsConfigurationOutput {
	return o
}

func (o WindowsConfigurationOutput) ToWindowsConfigurationOutputWithContext(ctx context.Context) WindowsConfigurationOutput {
	return o
}

func (o WindowsConfigurationOutput) ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput {
	return o.ToWindowsConfigurationPtrOutputWithContext(context.Background())
}

func (o WindowsConfigurationOutput) ToWindowsConfigurationPtrOutputWithContext(ctx context.Context) WindowsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WindowsConfiguration) *WindowsConfiguration {
		return &v
	}).(WindowsConfigurationPtrOutput)
}

func (o WindowsConfigurationOutput) AdditionalUnattendContent() AdditionalUnattendContentArrayOutput {
	return o.ApplyT(func(v WindowsConfiguration) []AdditionalUnattendContent { return v.AdditionalUnattendContent }).(AdditionalUnattendContentArrayOutput)
}

func (o WindowsConfigurationOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WindowsConfiguration) *bool { return v.EnableAutomaticUpdates }).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationOutput) PatchSettings() PatchSettingsPtrOutput {
	return o.ApplyT(func(v WindowsConfiguration) *PatchSettings { return v.PatchSettings }).(PatchSettingsPtrOutput)
}

func (o WindowsConfigurationOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WindowsConfiguration) *bool { return v.ProvisionVMAgent }).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsConfiguration) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

func (o WindowsConfigurationOutput) WinRM() WinRMConfigurationPtrOutput {
	return o.ApplyT(func(v WindowsConfiguration) *WinRMConfiguration { return v.WinRM }).(WinRMConfigurationPtrOutput)
}

type WindowsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (WindowsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsConfiguration)(nil)).Elem()
}

func (o WindowsConfigurationPtrOutput) ToWindowsConfigurationPtrOutput() WindowsConfigurationPtrOutput {
	return o
}

func (o WindowsConfigurationPtrOutput) ToWindowsConfigurationPtrOutputWithContext(ctx context.Context) WindowsConfigurationPtrOutput {
	return o
}

func (o WindowsConfigurationPtrOutput) Elem() WindowsConfigurationOutput {
	return o.ApplyT(func(v *WindowsConfiguration) WindowsConfiguration {
		if v != nil {
			return *v
		}
		var ret WindowsConfiguration
		return ret
	}).(WindowsConfigurationOutput)
}

func (o WindowsConfigurationPtrOutput) AdditionalUnattendContent() AdditionalUnattendContentArrayOutput {
	return o.ApplyT(func(v *WindowsConfiguration) []AdditionalUnattendContent {
		if v == nil {
			return nil
		}
		return v.AdditionalUnattendContent
	}).(AdditionalUnattendContentArrayOutput)
}

func (o WindowsConfigurationPtrOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WindowsConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAutomaticUpdates
	}).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationPtrOutput) PatchSettings() PatchSettingsPtrOutput {
	return o.ApplyT(func(v *WindowsConfiguration) *PatchSettings {
		if v == nil {
			return nil
		}
		return v.PatchSettings
	}).(PatchSettingsPtrOutput)
}

func (o WindowsConfigurationPtrOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WindowsConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.ProvisionVMAgent
	}).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationPtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

func (o WindowsConfigurationPtrOutput) WinRM() WinRMConfigurationPtrOutput {
	return o.ApplyT(func(v *WindowsConfiguration) *WinRMConfiguration {
		if v == nil {
			return nil
		}
		return v.WinRM
	}).(WinRMConfigurationPtrOutput)
}

type WindowsConfigurationResponse struct {
	AdditionalUnattendContent []AdditionalUnattendContentResponse `pulumi:"additionalUnattendContent"`
	EnableAutomaticUpdates    *bool                               `pulumi:"enableAutomaticUpdates"`
	PatchSettings             *PatchSettingsResponse              `pulumi:"patchSettings"`
	ProvisionVMAgent          *bool                               `pulumi:"provisionVMAgent"`
	TimeZone                  *string                             `pulumi:"timeZone"`
	WinRM                     *WinRMConfigurationResponse         `pulumi:"winRM"`
}

type WindowsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (WindowsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsConfigurationResponse)(nil)).Elem()
}

func (o WindowsConfigurationResponseOutput) ToWindowsConfigurationResponseOutput() WindowsConfigurationResponseOutput {
	return o
}

func (o WindowsConfigurationResponseOutput) ToWindowsConfigurationResponseOutputWithContext(ctx context.Context) WindowsConfigurationResponseOutput {
	return o
}

func (o WindowsConfigurationResponseOutput) AdditionalUnattendContent() AdditionalUnattendContentResponseArrayOutput {
	return o.ApplyT(func(v WindowsConfigurationResponse) []AdditionalUnattendContentResponse {
		return v.AdditionalUnattendContent
	}).(AdditionalUnattendContentResponseArrayOutput)
}

func (o WindowsConfigurationResponseOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WindowsConfigurationResponse) *bool { return v.EnableAutomaticUpdates }).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationResponseOutput) PatchSettings() PatchSettingsResponsePtrOutput {
	return o.ApplyT(func(v WindowsConfigurationResponse) *PatchSettingsResponse { return v.PatchSettings }).(PatchSettingsResponsePtrOutput)
}

func (o WindowsConfigurationResponseOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WindowsConfigurationResponse) *bool { return v.ProvisionVMAgent }).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationResponseOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsConfigurationResponse) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

func (o WindowsConfigurationResponseOutput) WinRM() WinRMConfigurationResponsePtrOutput {
	return o.ApplyT(func(v WindowsConfigurationResponse) *WinRMConfigurationResponse { return v.WinRM }).(WinRMConfigurationResponsePtrOutput)
}

type WindowsConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (WindowsConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsConfigurationResponse)(nil)).Elem()
}

func (o WindowsConfigurationResponsePtrOutput) ToWindowsConfigurationResponsePtrOutput() WindowsConfigurationResponsePtrOutput {
	return o
}

func (o WindowsConfigurationResponsePtrOutput) ToWindowsConfigurationResponsePtrOutputWithContext(ctx context.Context) WindowsConfigurationResponsePtrOutput {
	return o
}

func (o WindowsConfigurationResponsePtrOutput) Elem() WindowsConfigurationResponseOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) WindowsConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret WindowsConfigurationResponse
		return ret
	}).(WindowsConfigurationResponseOutput)
}

func (o WindowsConfigurationResponsePtrOutput) AdditionalUnattendContent() AdditionalUnattendContentResponseArrayOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) []AdditionalUnattendContentResponse {
		if v == nil {
			return nil
		}
		return v.AdditionalUnattendContent
	}).(AdditionalUnattendContentResponseArrayOutput)
}

func (o WindowsConfigurationResponsePtrOutput) EnableAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableAutomaticUpdates
	}).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationResponsePtrOutput) PatchSettings() PatchSettingsResponsePtrOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) *PatchSettingsResponse {
		if v == nil {
			return nil
		}
		return v.PatchSettings
	}).(PatchSettingsResponsePtrOutput)
}

func (o WindowsConfigurationResponsePtrOutput) ProvisionVMAgent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ProvisionVMAgent
	}).(pulumi.BoolPtrOutput)
}

func (o WindowsConfigurationResponsePtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

func (o WindowsConfigurationResponsePtrOutput) WinRM() WinRMConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *WindowsConfigurationResponse) *WinRMConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.WinRM
	}).(WinRMConfigurationResponsePtrOutput)
}

type WindowsVMGuestPatchAutomaticByPlatformSettings struct {
	RebootSetting *string `pulumi:"rebootSetting"`
}





type WindowsVMGuestPatchAutomaticByPlatformSettingsInput interface {
	pulumi.Input

	ToWindowsVMGuestPatchAutomaticByPlatformSettingsOutput() WindowsVMGuestPatchAutomaticByPlatformSettingsOutput
	ToWindowsVMGuestPatchAutomaticByPlatformSettingsOutputWithContext(context.Context) WindowsVMGuestPatchAutomaticByPlatformSettingsOutput
}

type WindowsVMGuestPatchAutomaticByPlatformSettingsArgs struct {
	RebootSetting pulumi.StringPtrInput `pulumi:"rebootSetting"`
}

func (WindowsVMGuestPatchAutomaticByPlatformSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsVMGuestPatchAutomaticByPlatformSettings)(nil)).Elem()
}

func (i WindowsVMGuestPatchAutomaticByPlatformSettingsArgs) ToWindowsVMGuestPatchAutomaticByPlatformSettingsOutput() WindowsVMGuestPatchAutomaticByPlatformSettingsOutput {
	return i.ToWindowsVMGuestPatchAutomaticByPlatformSettingsOutputWithContext(context.Background())
}

func (i WindowsVMGuestPatchAutomaticByPlatformSettingsArgs) ToWindowsVMGuestPatchAutomaticByPlatformSettingsOutputWithContext(ctx context.Context) WindowsVMGuestPatchAutomaticByPlatformSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsVMGuestPatchAutomaticByPlatformSettingsOutput)
}

func (i WindowsVMGuestPatchAutomaticByPlatformSettingsArgs) ToWindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput() WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return i.ToWindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutputWithContext(context.Background())
}

func (i WindowsVMGuestPatchAutomaticByPlatformSettingsArgs) ToWindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutputWithContext(ctx context.Context) WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsVMGuestPatchAutomaticByPlatformSettingsOutput).ToWindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutputWithContext(ctx)
}









type WindowsVMGuestPatchAutomaticByPlatformSettingsPtrInput interface {
	pulumi.Input

	ToWindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput() WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput
	ToWindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutputWithContext(context.Context) WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput
}

type windowsVMGuestPatchAutomaticByPlatformSettingsPtrType WindowsVMGuestPatchAutomaticByPlatformSettingsArgs

func WindowsVMGuestPatchAutomaticByPlatformSettingsPtr(v *WindowsVMGuestPatchAutomaticByPlatformSettingsArgs) WindowsVMGuestPatchAutomaticByPlatformSettingsPtrInput {
	return (*windowsVMGuestPatchAutomaticByPlatformSettingsPtrType)(v)
}

func (*windowsVMGuestPatchAutomaticByPlatformSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsVMGuestPatchAutomaticByPlatformSettings)(nil)).Elem()
}

func (i *windowsVMGuestPatchAutomaticByPlatformSettingsPtrType) ToWindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput() WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return i.ToWindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutputWithContext(context.Background())
}

func (i *windowsVMGuestPatchAutomaticByPlatformSettingsPtrType) ToWindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutputWithContext(ctx context.Context) WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput)
}

type WindowsVMGuestPatchAutomaticByPlatformSettingsOutput struct{ *pulumi.OutputState }

func (WindowsVMGuestPatchAutomaticByPlatformSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsVMGuestPatchAutomaticByPlatformSettings)(nil)).Elem()
}

func (o WindowsVMGuestPatchAutomaticByPlatformSettingsOutput) ToWindowsVMGuestPatchAutomaticByPlatformSettingsOutput() WindowsVMGuestPatchAutomaticByPlatformSettingsOutput {
	return o
}

func (o WindowsVMGuestPatchAutomaticByPlatformSettingsOutput) ToWindowsVMGuestPatchAutomaticByPlatformSettingsOutputWithContext(ctx context.Context) WindowsVMGuestPatchAutomaticByPlatformSettingsOutput {
	return o
}

func (o WindowsVMGuestPatchAutomaticByPlatformSettingsOutput) ToWindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput() WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return o.ToWindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutputWithContext(context.Background())
}

func (o WindowsVMGuestPatchAutomaticByPlatformSettingsOutput) ToWindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutputWithContext(ctx context.Context) WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WindowsVMGuestPatchAutomaticByPlatformSettings) *WindowsVMGuestPatchAutomaticByPlatformSettings {
		return &v
	}).(WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput)
}

func (o WindowsVMGuestPatchAutomaticByPlatformSettingsOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsVMGuestPatchAutomaticByPlatformSettings) *string { return v.RebootSetting }).(pulumi.StringPtrOutput)
}

type WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput struct{ *pulumi.OutputState }

func (WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsVMGuestPatchAutomaticByPlatformSettings)(nil)).Elem()
}

func (o WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput) ToWindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput() WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return o
}

func (o WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput) ToWindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutputWithContext(ctx context.Context) WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput {
	return o
}

func (o WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput) Elem() WindowsVMGuestPatchAutomaticByPlatformSettingsOutput {
	return o.ApplyT(func(v *WindowsVMGuestPatchAutomaticByPlatformSettings) WindowsVMGuestPatchAutomaticByPlatformSettings {
		if v != nil {
			return *v
		}
		var ret WindowsVMGuestPatchAutomaticByPlatformSettings
		return ret
	}).(WindowsVMGuestPatchAutomaticByPlatformSettingsOutput)
}

func (o WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsVMGuestPatchAutomaticByPlatformSettings) *string {
		if v == nil {
			return nil
		}
		return v.RebootSetting
	}).(pulumi.StringPtrOutput)
}

type WindowsVMGuestPatchAutomaticByPlatformSettingsResponse struct {
	RebootSetting *string `pulumi:"rebootSetting"`
}

type WindowsVMGuestPatchAutomaticByPlatformSettingsResponseOutput struct{ *pulumi.OutputState }

func (WindowsVMGuestPatchAutomaticByPlatformSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsVMGuestPatchAutomaticByPlatformSettingsResponse)(nil)).Elem()
}

func (o WindowsVMGuestPatchAutomaticByPlatformSettingsResponseOutput) ToWindowsVMGuestPatchAutomaticByPlatformSettingsResponseOutput() WindowsVMGuestPatchAutomaticByPlatformSettingsResponseOutput {
	return o
}

func (o WindowsVMGuestPatchAutomaticByPlatformSettingsResponseOutput) ToWindowsVMGuestPatchAutomaticByPlatformSettingsResponseOutputWithContext(ctx context.Context) WindowsVMGuestPatchAutomaticByPlatformSettingsResponseOutput {
	return o
}

func (o WindowsVMGuestPatchAutomaticByPlatformSettingsResponseOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsVMGuestPatchAutomaticByPlatformSettingsResponse) *string { return v.RebootSetting }).(pulumi.StringPtrOutput)
}

type WindowsVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (WindowsVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsVMGuestPatchAutomaticByPlatformSettingsResponse)(nil)).Elem()
}

func (o WindowsVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput) ToWindowsVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput() WindowsVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput {
	return o
}

func (o WindowsVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput) ToWindowsVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutputWithContext(ctx context.Context) WindowsVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput {
	return o
}

func (o WindowsVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput) Elem() WindowsVMGuestPatchAutomaticByPlatformSettingsResponseOutput {
	return o.ApplyT(func(v *WindowsVMGuestPatchAutomaticByPlatformSettingsResponse) WindowsVMGuestPatchAutomaticByPlatformSettingsResponse {
		if v != nil {
			return *v
		}
		var ret WindowsVMGuestPatchAutomaticByPlatformSettingsResponse
		return ret
	}).(WindowsVMGuestPatchAutomaticByPlatformSettingsResponseOutput)
}

func (o WindowsVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsVMGuestPatchAutomaticByPlatformSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.RebootSetting
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AdditionalCapabilitiesOutput{})
	pulumi.RegisterOutputType(AdditionalCapabilitiesPtrOutput{})
	pulumi.RegisterOutputType(AdditionalCapabilitiesResponseOutput{})
	pulumi.RegisterOutputType(AdditionalCapabilitiesResponsePtrOutput{})
	pulumi.RegisterOutputType(AdditionalUnattendContentOutput{})
	pulumi.RegisterOutputType(AdditionalUnattendContentArrayOutput{})
	pulumi.RegisterOutputType(AdditionalUnattendContentResponseOutput{})
	pulumi.RegisterOutputType(AdditionalUnattendContentResponseArrayOutput{})
	pulumi.RegisterOutputType(ApiEntityReferenceOutput{})
	pulumi.RegisterOutputType(ApiEntityReferencePtrOutput{})
	pulumi.RegisterOutputType(ApiEntityReferenceArrayOutput{})
	pulumi.RegisterOutputType(ApiEntityReferenceResponseOutput{})
	pulumi.RegisterOutputType(ApiEntityReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiEntityReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(ApiErrorBaseResponseOutput{})
	pulumi.RegisterOutputType(ApiErrorBaseResponseArrayOutput{})
	pulumi.RegisterOutputType(ApiErrorResponseOutput{})
	pulumi.RegisterOutputType(ApiErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationProfileOutput{})
	pulumi.RegisterOutputType(ApplicationProfilePtrOutput{})
	pulumi.RegisterOutputType(ApplicationProfileResponseOutput{})
	pulumi.RegisterOutputType(ApplicationProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(AutomaticOSUpgradePolicyOutput{})
	pulumi.RegisterOutputType(AutomaticOSUpgradePolicyPtrOutput{})
	pulumi.RegisterOutputType(AutomaticOSUpgradePolicyResponseOutput{})
	pulumi.RegisterOutputType(AutomaticOSUpgradePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(AutomaticRepairsPolicyOutput{})
	pulumi.RegisterOutputType(AutomaticRepairsPolicyPtrOutput{})
	pulumi.RegisterOutputType(AutomaticRepairsPolicyResponseOutput{})
	pulumi.RegisterOutputType(AutomaticRepairsPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(AvailablePatchSummaryResponseOutput{})
	pulumi.RegisterOutputType(AvailablePatchSummaryResponsePtrOutput{})
	pulumi.RegisterOutputType(BillingProfileOutput{})
	pulumi.RegisterOutputType(BillingProfilePtrOutput{})
	pulumi.RegisterOutputType(BillingProfileResponseOutput{})
	pulumi.RegisterOutputType(BillingProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(BootDiagnosticsOutput{})
	pulumi.RegisterOutputType(BootDiagnosticsPtrOutput{})
	pulumi.RegisterOutputType(BootDiagnosticsInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(BootDiagnosticsInstanceViewResponsePtrOutput{})
	pulumi.RegisterOutputType(BootDiagnosticsResponseOutput{})
	pulumi.RegisterOutputType(BootDiagnosticsResponsePtrOutput{})
	pulumi.RegisterOutputType(CapacityReservationGroupInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(CapacityReservationInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(CapacityReservationInstanceViewWithNameResponseOutput{})
	pulumi.RegisterOutputType(CapacityReservationInstanceViewWithNameResponseArrayOutput{})
	pulumi.RegisterOutputType(CapacityReservationProfileOutput{})
	pulumi.RegisterOutputType(CapacityReservationProfilePtrOutput{})
	pulumi.RegisterOutputType(CapacityReservationProfileResponseOutput{})
	pulumi.RegisterOutputType(CapacityReservationProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(CapacityReservationUtilizationResponseOutput{})
	pulumi.RegisterOutputType(CapacityReservationUtilizationResponsePtrOutput{})
	pulumi.RegisterOutputType(DataDiskOutput{})
	pulumi.RegisterOutputType(DataDiskArrayOutput{})
	pulumi.RegisterOutputType(DataDiskResponseOutput{})
	pulumi.RegisterOutputType(DataDiskResponseArrayOutput{})
	pulumi.RegisterOutputType(DedicatedHostAllocatableVMResponseOutput{})
	pulumi.RegisterOutputType(DedicatedHostAllocatableVMResponseArrayOutput{})
	pulumi.RegisterOutputType(DedicatedHostAvailableCapacityResponseOutput{})
	pulumi.RegisterOutputType(DedicatedHostAvailableCapacityResponsePtrOutput{})
	pulumi.RegisterOutputType(DedicatedHostGroupInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(DedicatedHostGroupPropertiesAdditionalCapabilitiesOutput{})
	pulumi.RegisterOutputType(DedicatedHostGroupPropertiesAdditionalCapabilitiesPtrOutput{})
	pulumi.RegisterOutputType(DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesOutput{})
	pulumi.RegisterOutputType(DedicatedHostGroupPropertiesResponseAdditionalCapabilitiesPtrOutput{})
	pulumi.RegisterOutputType(DedicatedHostInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(DedicatedHostInstanceViewWithNameResponseOutput{})
	pulumi.RegisterOutputType(DedicatedHostInstanceViewWithNameResponseArrayOutput{})
	pulumi.RegisterOutputType(DiagnosticsProfileOutput{})
	pulumi.RegisterOutputType(DiagnosticsProfilePtrOutput{})
	pulumi.RegisterOutputType(DiagnosticsProfileResponseOutput{})
	pulumi.RegisterOutputType(DiagnosticsProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(DiffDiskSettingsOutput{})
	pulumi.RegisterOutputType(DiffDiskSettingsPtrOutput{})
	pulumi.RegisterOutputType(DiffDiskSettingsResponseOutput{})
	pulumi.RegisterOutputType(DiffDiskSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSetParametersOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSetParametersPtrOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSetParametersResponseOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSetParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSettingsOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSettingsPtrOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSettingsResponseOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(DiskInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(DiskInstanceViewResponseArrayOutput{})
	pulumi.RegisterOutputType(DiskRestorePointInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(DiskRestorePointInstanceViewResponseArrayOutput{})
	pulumi.RegisterOutputType(DiskRestorePointReplicationStatusResponseOutput{})
	pulumi.RegisterOutputType(DiskRestorePointReplicationStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationOutput{})
	pulumi.RegisterOutputType(ExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponsePtrOutput{})
	pulumi.RegisterOutputType(HardwareProfileOutput{})
	pulumi.RegisterOutputType(HardwareProfilePtrOutput{})
	pulumi.RegisterOutputType(HardwareProfileResponseOutput{})
	pulumi.RegisterOutputType(HardwareProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageDataDiskOutput{})
	pulumi.RegisterOutputType(ImageDataDiskArrayOutput{})
	pulumi.RegisterOutputType(ImageDataDiskResponseOutput{})
	pulumi.RegisterOutputType(ImageDataDiskResponseArrayOutput{})
	pulumi.RegisterOutputType(ImageOSDiskOutput{})
	pulumi.RegisterOutputType(ImageOSDiskPtrOutput{})
	pulumi.RegisterOutputType(ImageOSDiskResponseOutput{})
	pulumi.RegisterOutputType(ImageOSDiskResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageReferenceOutput{})
	pulumi.RegisterOutputType(ImageReferencePtrOutput{})
	pulumi.RegisterOutputType(ImageReferenceResponseOutput{})
	pulumi.RegisterOutputType(ImageReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageStorageProfileOutput{})
	pulumi.RegisterOutputType(ImageStorageProfilePtrOutput{})
	pulumi.RegisterOutputType(ImageStorageProfileResponseOutput{})
	pulumi.RegisterOutputType(ImageStorageProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(InnerErrorResponseOutput{})
	pulumi.RegisterOutputType(InnerErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(InstanceViewStatusOutput{})
	pulumi.RegisterOutputType(InstanceViewStatusPtrOutput{})
	pulumi.RegisterOutputType(InstanceViewStatusArrayOutput{})
	pulumi.RegisterOutputType(InstanceViewStatusResponseOutput{})
	pulumi.RegisterOutputType(InstanceViewStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(InstanceViewStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultSecretReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultSecretReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultSecretReferenceResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultSecretReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(LastPatchInstallationSummaryResponseOutput{})
	pulumi.RegisterOutputType(LastPatchInstallationSummaryResponsePtrOutput{})
	pulumi.RegisterOutputType(LinuxConfigurationOutput{})
	pulumi.RegisterOutputType(LinuxConfigurationPtrOutput{})
	pulumi.RegisterOutputType(LinuxConfigurationResponseOutput{})
	pulumi.RegisterOutputType(LinuxConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(LinuxPatchSettingsOutput{})
	pulumi.RegisterOutputType(LinuxPatchSettingsPtrOutput{})
	pulumi.RegisterOutputType(LinuxPatchSettingsResponseOutput{})
	pulumi.RegisterOutputType(LinuxPatchSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(LinuxVMGuestPatchAutomaticByPlatformSettingsOutput{})
	pulumi.RegisterOutputType(LinuxVMGuestPatchAutomaticByPlatformSettingsPtrOutput{})
	pulumi.RegisterOutputType(LinuxVMGuestPatchAutomaticByPlatformSettingsResponseOutput{})
	pulumi.RegisterOutputType(LinuxVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsOutputResponseOutput{})
	pulumi.RegisterOutputType(MaintenanceRedeployStatusResponseOutput{})
	pulumi.RegisterOutputType(MaintenanceRedeployStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedDiskParametersOutput{})
	pulumi.RegisterOutputType(ManagedDiskParametersPtrOutput{})
	pulumi.RegisterOutputType(ManagedDiskParametersResponseOutput{})
	pulumi.RegisterOutputType(ManagedDiskParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceReferenceOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceReferenceArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceReferenceResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkProfileOutput{})
	pulumi.RegisterOutputType(NetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(OSDiskOutput{})
	pulumi.RegisterOutputType(OSDiskPtrOutput{})
	pulumi.RegisterOutputType(OSDiskResponseOutput{})
	pulumi.RegisterOutputType(OSDiskResponsePtrOutput{})
	pulumi.RegisterOutputType(OSProfileOutput{})
	pulumi.RegisterOutputType(OSProfilePtrOutput{})
	pulumi.RegisterOutputType(OSProfileResponseOutput{})
	pulumi.RegisterOutputType(OSProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(PatchSettingsOutput{})
	pulumi.RegisterOutputType(PatchSettingsPtrOutput{})
	pulumi.RegisterOutputType(PatchSettingsResponseOutput{})
	pulumi.RegisterOutputType(PatchSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(PlanOutput{})
	pulumi.RegisterOutputType(PlanPtrOutput{})
	pulumi.RegisterOutputType(PlanResponseOutput{})
	pulumi.RegisterOutputType(PlanResponsePtrOutput{})
	pulumi.RegisterOutputType(ProximityPlacementGroupPropertiesIntentOutput{})
	pulumi.RegisterOutputType(ProximityPlacementGroupPropertiesIntentPtrOutput{})
	pulumi.RegisterOutputType(ProximityPlacementGroupPropertiesResponseIntentOutput{})
	pulumi.RegisterOutputType(ProximityPlacementGroupPropertiesResponseIntentPtrOutput{})
	pulumi.RegisterOutputType(PublicIPAddressSkuOutput{})
	pulumi.RegisterOutputType(PublicIPAddressSkuPtrOutput{})
	pulumi.RegisterOutputType(PublicIPAddressSkuResponseOutput{})
	pulumi.RegisterOutputType(PublicIPAddressSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(RestorePointCollectionSourcePropertiesOutput{})
	pulumi.RegisterOutputType(RestorePointCollectionSourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(RestorePointCollectionSourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(RestorePointCollectionSourcePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(RestorePointInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(RestorePointResponseOutput{})
	pulumi.RegisterOutputType(RestorePointResponseArrayOutput{})
	pulumi.RegisterOutputType(RestorePointSourceMetadataResponseOutput{})
	pulumi.RegisterOutputType(RestorePointSourceVMDataDiskResponseOutput{})
	pulumi.RegisterOutputType(RestorePointSourceVMDataDiskResponseArrayOutput{})
	pulumi.RegisterOutputType(RestorePointSourceVMOSDiskResponseOutput{})
	pulumi.RegisterOutputType(RestorePointSourceVMOSDiskResponsePtrOutput{})
	pulumi.RegisterOutputType(RestorePointSourceVMStorageProfileResponseOutput{})
	pulumi.RegisterOutputType(RestorePointSourceVMStorageProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(RollingUpgradePolicyOutput{})
	pulumi.RegisterOutputType(RollingUpgradePolicyPtrOutput{})
	pulumi.RegisterOutputType(RollingUpgradePolicyResponseOutput{})
	pulumi.RegisterOutputType(RollingUpgradePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(RunCommandInputParameterOutput{})
	pulumi.RegisterOutputType(RunCommandInputParameterArrayOutput{})
	pulumi.RegisterOutputType(RunCommandInputParameterResponseOutput{})
	pulumi.RegisterOutputType(RunCommandInputParameterResponseArrayOutput{})
	pulumi.RegisterOutputType(ScaleInPolicyOutput{})
	pulumi.RegisterOutputType(ScaleInPolicyPtrOutput{})
	pulumi.RegisterOutputType(ScaleInPolicyResponseOutput{})
	pulumi.RegisterOutputType(ScaleInPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ScheduledEventsProfileOutput{})
	pulumi.RegisterOutputType(ScheduledEventsProfilePtrOutput{})
	pulumi.RegisterOutputType(ScheduledEventsProfileResponseOutput{})
	pulumi.RegisterOutputType(ScheduledEventsProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SecurityProfileOutput{})
	pulumi.RegisterOutputType(SecurityProfilePtrOutput{})
	pulumi.RegisterOutputType(SecurityProfileResponseOutput{})
	pulumi.RegisterOutputType(SecurityProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SpotRestorePolicyOutput{})
	pulumi.RegisterOutputType(SpotRestorePolicyPtrOutput{})
	pulumi.RegisterOutputType(SpotRestorePolicyResponseOutput{})
	pulumi.RegisterOutputType(SpotRestorePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(SshConfigurationOutput{})
	pulumi.RegisterOutputType(SshConfigurationPtrOutput{})
	pulumi.RegisterOutputType(SshConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SshConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(SshPublicKeyTypeOutput{})
	pulumi.RegisterOutputType(SshPublicKeyTypeArrayOutput{})
	pulumi.RegisterOutputType(SshPublicKeyResponseOutput{})
	pulumi.RegisterOutputType(SshPublicKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(StorageProfileOutput{})
	pulumi.RegisterOutputType(StorageProfilePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileResponseOutput{})
	pulumi.RegisterOutputType(StorageProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourcePtrOutput{})
	pulumi.RegisterOutputType(SubResourceArrayOutput{})
	pulumi.RegisterOutputType(SubResourceReadOnlyResponseOutput{})
	pulumi.RegisterOutputType(SubResourceReadOnlyResponseArrayOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SubResourceWithColocationStatusResponseOutput{})
	pulumi.RegisterOutputType(SubResourceWithColocationStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(TerminateNotificationProfileOutput{})
	pulumi.RegisterOutputType(TerminateNotificationProfilePtrOutput{})
	pulumi.RegisterOutputType(TerminateNotificationProfileResponseOutput{})
	pulumi.RegisterOutputType(TerminateNotificationProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(UefiSettingsOutput{})
	pulumi.RegisterOutputType(UefiSettingsPtrOutput{})
	pulumi.RegisterOutputType(UefiSettingsResponseOutput{})
	pulumi.RegisterOutputType(UefiSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(UpgradePolicyOutput{})
	pulumi.RegisterOutputType(UpgradePolicyPtrOutput{})
	pulumi.RegisterOutputType(UpgradePolicyResponseOutput{})
	pulumi.RegisterOutputType(UpgradePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentitiesResponseUserAssignedIdentitiesOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentitiesResponseUserAssignedIdentitiesMapOutput{})
	pulumi.RegisterOutputType(VMDiskSecurityProfileOutput{})
	pulumi.RegisterOutputType(VMDiskSecurityProfilePtrOutput{})
	pulumi.RegisterOutputType(VMDiskSecurityProfileResponseOutput{})
	pulumi.RegisterOutputType(VMDiskSecurityProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(VMGalleryApplicationOutput{})
	pulumi.RegisterOutputType(VMGalleryApplicationArrayOutput{})
	pulumi.RegisterOutputType(VMGalleryApplicationResponseOutput{})
	pulumi.RegisterOutputType(VMGalleryApplicationResponseArrayOutput{})
	pulumi.RegisterOutputType(VMSizePropertiesOutput{})
	pulumi.RegisterOutputType(VMSizePropertiesPtrOutput{})
	pulumi.RegisterOutputType(VMSizePropertiesResponseOutput{})
	pulumi.RegisterOutputType(VMSizePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(VaultCertificateOutput{})
	pulumi.RegisterOutputType(VaultCertificateArrayOutput{})
	pulumi.RegisterOutputType(VaultCertificateResponseOutput{})
	pulumi.RegisterOutputType(VaultCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupArrayOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupResponseOutput{})
	pulumi.RegisterOutputType(VaultSecretGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualHardDiskOutput{})
	pulumi.RegisterOutputType(VirtualHardDiskPtrOutput{})
	pulumi.RegisterOutputType(VirtualHardDiskResponseOutput{})
	pulumi.RegisterOutputType(VirtualHardDiskResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineAgentInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineAgentInstanceViewResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionHandlerInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionHandlerInstanceViewResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionInstanceViewOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionInstanceViewPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionInstanceViewResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionInstanceViewResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineExtensionResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineHealthStatusResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineIdentityOutput{})
	pulumi.RegisterOutputType(VirtualMachineIdentityPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineIdentityResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineIpTagOutput{})
	pulumi.RegisterOutputType(VirtualMachineIpTagArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineIpTagResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineIpTagResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineNetworkInterfaceConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualMachineNetworkInterfaceConfigurationArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineNetworkInterfaceConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineNetworkInterfaceConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineNetworkInterfaceDnsSettingsConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualMachineNetworkInterfaceDnsSettingsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineNetworkInterfaceDnsSettingsConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineNetworkInterfaceIPConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualMachineNetworkInterfaceIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineNetworkInterfaceIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineNetworkInterfaceIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachinePatchStatusResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachinePatchStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachinePublicIPAddressConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualMachinePublicIPAddressConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachinePublicIPAddressConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachinePublicIPAddressConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachinePublicIPAddressDnsSettingsConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualMachinePublicIPAddressDnsSettingsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachinePublicIPAddressDnsSettingsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachinePublicIPAddressDnsSettingsConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineRunCommandInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineRunCommandScriptSourceOutput{})
	pulumi.RegisterOutputType(VirtualMachineRunCommandScriptSourcePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineRunCommandScriptSourceResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineRunCommandScriptSourceResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetDataDiskOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetDataDiskArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetDataDiskResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetDataDiskResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionTypeOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionTypeArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionProfileOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetHardwareProfileOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetHardwareProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetHardwareProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetHardwareProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIPConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIdentityOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIdentityPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIdentityResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIpTagOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIpTagArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIpTagResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetIpTagResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetManagedDiskParametersOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetManagedDiskParametersPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetManagedDiskParametersResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetManagedDiskParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationDnsSettingsOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationDnsSettingsPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationDnsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkProfileOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetNetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSDiskOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSDiskPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSDiskResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSDiskResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSProfileOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetOSProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetPublicIPAddressConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetPublicIPAddressConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetPublicIPAddressConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetPublicIPAddressConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetStorageProfileOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetStorageProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetStorageProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetStorageProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMNetworkProfileConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMNetworkProfileConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMNetworkProfileConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMProfileOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMProtectionPolicyOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMProtectionPolicyPtrOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMProtectionPolicyResponseOutput{})
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(WinRMConfigurationOutput{})
	pulumi.RegisterOutputType(WinRMConfigurationPtrOutput{})
	pulumi.RegisterOutputType(WinRMConfigurationResponseOutput{})
	pulumi.RegisterOutputType(WinRMConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(WinRMListenerOutput{})
	pulumi.RegisterOutputType(WinRMListenerArrayOutput{})
	pulumi.RegisterOutputType(WinRMListenerResponseOutput{})
	pulumi.RegisterOutputType(WinRMListenerResponseArrayOutput{})
	pulumi.RegisterOutputType(WindowsConfigurationOutput{})
	pulumi.RegisterOutputType(WindowsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(WindowsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(WindowsConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(WindowsVMGuestPatchAutomaticByPlatformSettingsOutput{})
	pulumi.RegisterOutputType(WindowsVMGuestPatchAutomaticByPlatformSettingsPtrOutput{})
	pulumi.RegisterOutputType(WindowsVMGuestPatchAutomaticByPlatformSettingsResponseOutput{})
	pulumi.RegisterOutputType(WindowsVMGuestPatchAutomaticByPlatformSettingsResponsePtrOutput{})
}
