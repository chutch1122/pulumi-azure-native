


package v20190601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Identity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v Identity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type IdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type PolicyDefinitionReference struct {
	Parameters         interface{} `pulumi:"parameters"`
	PolicyDefinitionId *string     `pulumi:"policyDefinitionId"`
}





type PolicyDefinitionReferenceInput interface {
	pulumi.Input

	ToPolicyDefinitionReferenceOutput() PolicyDefinitionReferenceOutput
	ToPolicyDefinitionReferenceOutputWithContext(context.Context) PolicyDefinitionReferenceOutput
}

type PolicyDefinitionReferenceArgs struct {
	Parameters         pulumi.Input          `pulumi:"parameters"`
	PolicyDefinitionId pulumi.StringPtrInput `pulumi:"policyDefinitionId"`
}

func (PolicyDefinitionReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinitionReference)(nil)).Elem()
}

func (i PolicyDefinitionReferenceArgs) ToPolicyDefinitionReferenceOutput() PolicyDefinitionReferenceOutput {
	return i.ToPolicyDefinitionReferenceOutputWithContext(context.Background())
}

func (i PolicyDefinitionReferenceArgs) ToPolicyDefinitionReferenceOutputWithContext(ctx context.Context) PolicyDefinitionReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyDefinitionReferenceOutput)
}





type PolicyDefinitionReferenceArrayInput interface {
	pulumi.Input

	ToPolicyDefinitionReferenceArrayOutput() PolicyDefinitionReferenceArrayOutput
	ToPolicyDefinitionReferenceArrayOutputWithContext(context.Context) PolicyDefinitionReferenceArrayOutput
}

type PolicyDefinitionReferenceArray []PolicyDefinitionReferenceInput

func (PolicyDefinitionReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyDefinitionReference)(nil)).Elem()
}

func (i PolicyDefinitionReferenceArray) ToPolicyDefinitionReferenceArrayOutput() PolicyDefinitionReferenceArrayOutput {
	return i.ToPolicyDefinitionReferenceArrayOutputWithContext(context.Background())
}

func (i PolicyDefinitionReferenceArray) ToPolicyDefinitionReferenceArrayOutputWithContext(ctx context.Context) PolicyDefinitionReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyDefinitionReferenceArrayOutput)
}

type PolicyDefinitionReferenceOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinitionReference)(nil)).Elem()
}

func (o PolicyDefinitionReferenceOutput) ToPolicyDefinitionReferenceOutput() PolicyDefinitionReferenceOutput {
	return o
}

func (o PolicyDefinitionReferenceOutput) ToPolicyDefinitionReferenceOutputWithContext(ctx context.Context) PolicyDefinitionReferenceOutput {
	return o
}

func (o PolicyDefinitionReferenceOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v PolicyDefinitionReference) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o PolicyDefinitionReferenceOutput) PolicyDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionReference) *string { return v.PolicyDefinitionId }).(pulumi.StringPtrOutput)
}

type PolicyDefinitionReferenceArrayOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyDefinitionReference)(nil)).Elem()
}

func (o PolicyDefinitionReferenceArrayOutput) ToPolicyDefinitionReferenceArrayOutput() PolicyDefinitionReferenceArrayOutput {
	return o
}

func (o PolicyDefinitionReferenceArrayOutput) ToPolicyDefinitionReferenceArrayOutputWithContext(ctx context.Context) PolicyDefinitionReferenceArrayOutput {
	return o
}

func (o PolicyDefinitionReferenceArrayOutput) Index(i pulumi.IntInput) PolicyDefinitionReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyDefinitionReference {
		return vs[0].([]PolicyDefinitionReference)[vs[1].(int)]
	}).(PolicyDefinitionReferenceOutput)
}

type PolicyDefinitionReferenceResponse struct {
	Parameters         interface{} `pulumi:"parameters"`
	PolicyDefinitionId *string     `pulumi:"policyDefinitionId"`
}

type PolicyDefinitionReferenceResponseOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyDefinitionReferenceResponse)(nil)).Elem()
}

func (o PolicyDefinitionReferenceResponseOutput) ToPolicyDefinitionReferenceResponseOutput() PolicyDefinitionReferenceResponseOutput {
	return o
}

func (o PolicyDefinitionReferenceResponseOutput) ToPolicyDefinitionReferenceResponseOutputWithContext(ctx context.Context) PolicyDefinitionReferenceResponseOutput {
	return o
}

func (o PolicyDefinitionReferenceResponseOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v PolicyDefinitionReferenceResponse) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o PolicyDefinitionReferenceResponseOutput) PolicyDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyDefinitionReferenceResponse) *string { return v.PolicyDefinitionId }).(pulumi.StringPtrOutput)
}

type PolicyDefinitionReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (PolicyDefinitionReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyDefinitionReferenceResponse)(nil)).Elem()
}

func (o PolicyDefinitionReferenceResponseArrayOutput) ToPolicyDefinitionReferenceResponseArrayOutput() PolicyDefinitionReferenceResponseArrayOutput {
	return o
}

func (o PolicyDefinitionReferenceResponseArrayOutput) ToPolicyDefinitionReferenceResponseArrayOutputWithContext(ctx context.Context) PolicyDefinitionReferenceResponseArrayOutput {
	return o
}

func (o PolicyDefinitionReferenceResponseArrayOutput) Index(i pulumi.IntInput) PolicyDefinitionReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyDefinitionReferenceResponse {
		return vs[0].([]PolicyDefinitionReferenceResponse)[vs[1].(int)]
	}).(PolicyDefinitionReferenceResponseOutput)
}

type PolicySku struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type PolicySkuInput interface {
	pulumi.Input

	ToPolicySkuOutput() PolicySkuOutput
	ToPolicySkuOutputWithContext(context.Context) PolicySkuOutput
}

type PolicySkuArgs struct {
	Name pulumi.StringInput    `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (PolicySkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicySku)(nil)).Elem()
}

func (i PolicySkuArgs) ToPolicySkuOutput() PolicySkuOutput {
	return i.ToPolicySkuOutputWithContext(context.Background())
}

func (i PolicySkuArgs) ToPolicySkuOutputWithContext(ctx context.Context) PolicySkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySkuOutput)
}

func (i PolicySkuArgs) ToPolicySkuPtrOutput() PolicySkuPtrOutput {
	return i.ToPolicySkuPtrOutputWithContext(context.Background())
}

func (i PolicySkuArgs) ToPolicySkuPtrOutputWithContext(ctx context.Context) PolicySkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySkuOutput).ToPolicySkuPtrOutputWithContext(ctx)
}









type PolicySkuPtrInput interface {
	pulumi.Input

	ToPolicySkuPtrOutput() PolicySkuPtrOutput
	ToPolicySkuPtrOutputWithContext(context.Context) PolicySkuPtrOutput
}

type policySkuPtrType PolicySkuArgs

func PolicySkuPtr(v *PolicySkuArgs) PolicySkuPtrInput {
	return (*policySkuPtrType)(v)
}

func (*policySkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySku)(nil)).Elem()
}

func (i *policySkuPtrType) ToPolicySkuPtrOutput() PolicySkuPtrOutput {
	return i.ToPolicySkuPtrOutputWithContext(context.Background())
}

func (i *policySkuPtrType) ToPolicySkuPtrOutputWithContext(ctx context.Context) PolicySkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySkuPtrOutput)
}

type PolicySkuOutput struct{ *pulumi.OutputState }

func (PolicySkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicySku)(nil)).Elem()
}

func (o PolicySkuOutput) ToPolicySkuOutput() PolicySkuOutput {
	return o
}

func (o PolicySkuOutput) ToPolicySkuOutputWithContext(ctx context.Context) PolicySkuOutput {
	return o
}

func (o PolicySkuOutput) ToPolicySkuPtrOutput() PolicySkuPtrOutput {
	return o.ToPolicySkuPtrOutputWithContext(context.Background())
}

func (o PolicySkuOutput) ToPolicySkuPtrOutputWithContext(ctx context.Context) PolicySkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicySku) *PolicySku {
		return &v
	}).(PolicySkuPtrOutput)
}

func (o PolicySkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PolicySku) string { return v.Name }).(pulumi.StringOutput)
}

func (o PolicySkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicySku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type PolicySkuPtrOutput struct{ *pulumi.OutputState }

func (PolicySkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySku)(nil)).Elem()
}

func (o PolicySkuPtrOutput) ToPolicySkuPtrOutput() PolicySkuPtrOutput {
	return o
}

func (o PolicySkuPtrOutput) ToPolicySkuPtrOutputWithContext(ctx context.Context) PolicySkuPtrOutput {
	return o
}

func (o PolicySkuPtrOutput) Elem() PolicySkuOutput {
	return o.ApplyT(func(v *PolicySku) PolicySku {
		if v != nil {
			return *v
		}
		var ret PolicySku
		return ret
	}).(PolicySkuOutput)
}

func (o PolicySkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PolicySkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type PolicySkuResponse struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}

type PolicySkuResponseOutput struct{ *pulumi.OutputState }

func (PolicySkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicySkuResponse)(nil)).Elem()
}

func (o PolicySkuResponseOutput) ToPolicySkuResponseOutput() PolicySkuResponseOutput {
	return o
}

func (o PolicySkuResponseOutput) ToPolicySkuResponseOutputWithContext(ctx context.Context) PolicySkuResponseOutput {
	return o
}

func (o PolicySkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PolicySkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PolicySkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicySkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type PolicySkuResponsePtrOutput struct{ *pulumi.OutputState }

func (PolicySkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySkuResponse)(nil)).Elem()
}

func (o PolicySkuResponsePtrOutput) ToPolicySkuResponsePtrOutput() PolicySkuResponsePtrOutput {
	return o
}

func (o PolicySkuResponsePtrOutput) ToPolicySkuResponsePtrOutputWithContext(ctx context.Context) PolicySkuResponsePtrOutput {
	return o
}

func (o PolicySkuResponsePtrOutput) Elem() PolicySkuResponseOutput {
	return o.ApplyT(func(v *PolicySkuResponse) PolicySkuResponse {
		if v != nil {
			return *v
		}
		var ret PolicySkuResponse
		return ret
	}).(PolicySkuResponseOutput)
}

func (o PolicySkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PolicySkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionReferenceOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionReferenceArrayOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionReferenceResponseOutput{})
	pulumi.RegisterOutputType(PolicyDefinitionReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(PolicySkuOutput{})
	pulumi.RegisterOutputType(PolicySkuPtrOutput{})
	pulumi.RegisterOutputType(PolicySkuResponseOutput{})
	pulumi.RegisterOutputType(PolicySkuResponsePtrOutput{})
}
