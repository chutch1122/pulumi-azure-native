


package v20190601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomerSecret struct {
	Algorithm     SupportedAlgorithm `pulumi:"algorithm"`
	KeyIdentifier string             `pulumi:"keyIdentifier"`
	KeyValue      string             `pulumi:"keyValue"`
}





type CustomerSecretInput interface {
	pulumi.Input

	ToCustomerSecretOutput() CustomerSecretOutput
	ToCustomerSecretOutputWithContext(context.Context) CustomerSecretOutput
}

type CustomerSecretArgs struct {
	Algorithm     SupportedAlgorithmInput `pulumi:"algorithm"`
	KeyIdentifier pulumi.StringInput      `pulumi:"keyIdentifier"`
	KeyValue      pulumi.StringInput      `pulumi:"keyValue"`
}

func (CustomerSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerSecret)(nil)).Elem()
}

func (i CustomerSecretArgs) ToCustomerSecretOutput() CustomerSecretOutput {
	return i.ToCustomerSecretOutputWithContext(context.Background())
}

func (i CustomerSecretArgs) ToCustomerSecretOutputWithContext(ctx context.Context) CustomerSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSecretOutput)
}





type CustomerSecretArrayInput interface {
	pulumi.Input

	ToCustomerSecretArrayOutput() CustomerSecretArrayOutput
	ToCustomerSecretArrayOutputWithContext(context.Context) CustomerSecretArrayOutput
}

type CustomerSecretArray []CustomerSecretInput

func (CustomerSecretArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomerSecret)(nil)).Elem()
}

func (i CustomerSecretArray) ToCustomerSecretArrayOutput() CustomerSecretArrayOutput {
	return i.ToCustomerSecretArrayOutputWithContext(context.Background())
}

func (i CustomerSecretArray) ToCustomerSecretArrayOutputWithContext(ctx context.Context) CustomerSecretArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSecretArrayOutput)
}

type CustomerSecretOutput struct{ *pulumi.OutputState }

func (CustomerSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerSecret)(nil)).Elem()
}

func (o CustomerSecretOutput) ToCustomerSecretOutput() CustomerSecretOutput {
	return o
}

func (o CustomerSecretOutput) ToCustomerSecretOutputWithContext(ctx context.Context) CustomerSecretOutput {
	return o
}

func (o CustomerSecretOutput) Algorithm() SupportedAlgorithmOutput {
	return o.ApplyT(func(v CustomerSecret) SupportedAlgorithm { return v.Algorithm }).(SupportedAlgorithmOutput)
}

func (o CustomerSecretOutput) KeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v CustomerSecret) string { return v.KeyIdentifier }).(pulumi.StringOutput)
}

func (o CustomerSecretOutput) KeyValue() pulumi.StringOutput {
	return o.ApplyT(func(v CustomerSecret) string { return v.KeyValue }).(pulumi.StringOutput)
}

type CustomerSecretArrayOutput struct{ *pulumi.OutputState }

func (CustomerSecretArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomerSecret)(nil)).Elem()
}

func (o CustomerSecretArrayOutput) ToCustomerSecretArrayOutput() CustomerSecretArrayOutput {
	return o
}

func (o CustomerSecretArrayOutput) ToCustomerSecretArrayOutputWithContext(ctx context.Context) CustomerSecretArrayOutput {
	return o
}

func (o CustomerSecretArrayOutput) Index(i pulumi.IntInput) CustomerSecretOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomerSecret {
		return vs[0].([]CustomerSecret)[vs[1].(int)]
	}).(CustomerSecretOutput)
}

type CustomerSecretResponse struct {
	Algorithm     string `pulumi:"algorithm"`
	KeyIdentifier string `pulumi:"keyIdentifier"`
	KeyValue      string `pulumi:"keyValue"`
}

type CustomerSecretResponseOutput struct{ *pulumi.OutputState }

func (CustomerSecretResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerSecretResponse)(nil)).Elem()
}

func (o CustomerSecretResponseOutput) ToCustomerSecretResponseOutput() CustomerSecretResponseOutput {
	return o
}

func (o CustomerSecretResponseOutput) ToCustomerSecretResponseOutputWithContext(ctx context.Context) CustomerSecretResponseOutput {
	return o
}

func (o CustomerSecretResponseOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v CustomerSecretResponse) string { return v.Algorithm }).(pulumi.StringOutput)
}

func (o CustomerSecretResponseOutput) KeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v CustomerSecretResponse) string { return v.KeyIdentifier }).(pulumi.StringOutput)
}

func (o CustomerSecretResponseOutput) KeyValue() pulumi.StringOutput {
	return o.ApplyT(func(v CustomerSecretResponse) string { return v.KeyValue }).(pulumi.StringOutput)
}

type CustomerSecretResponseArrayOutput struct{ *pulumi.OutputState }

func (CustomerSecretResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CustomerSecretResponse)(nil)).Elem()
}

func (o CustomerSecretResponseArrayOutput) ToCustomerSecretResponseArrayOutput() CustomerSecretResponseArrayOutput {
	return o
}

func (o CustomerSecretResponseArrayOutput) ToCustomerSecretResponseArrayOutputWithContext(ctx context.Context) CustomerSecretResponseArrayOutput {
	return o
}

func (o CustomerSecretResponseArrayOutput) Index(i pulumi.IntInput) CustomerSecretResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CustomerSecretResponse {
		return vs[0].([]CustomerSecretResponse)[vs[1].(int)]
	}).(CustomerSecretResponseOutput)
}

type Schedule struct {
	Name       *string  `pulumi:"name"`
	PolicyList []string `pulumi:"policyList"`
}





type ScheduleInput interface {
	pulumi.Input

	ToScheduleOutput() ScheduleOutput
	ToScheduleOutputWithContext(context.Context) ScheduleOutput
}

type ScheduleArgs struct {
	Name       pulumi.StringPtrInput   `pulumi:"name"`
	PolicyList pulumi.StringArrayInput `pulumi:"policyList"`
}

func (ScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Schedule)(nil)).Elem()
}

func (i ScheduleArgs) ToScheduleOutput() ScheduleOutput {
	return i.ToScheduleOutputWithContext(context.Background())
}

func (i ScheduleArgs) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput)
}





type ScheduleArrayInput interface {
	pulumi.Input

	ToScheduleArrayOutput() ScheduleArrayOutput
	ToScheduleArrayOutputWithContext(context.Context) ScheduleArrayOutput
}

type ScheduleArray []ScheduleInput

func (ScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Schedule)(nil)).Elem()
}

func (i ScheduleArray) ToScheduleArrayOutput() ScheduleArrayOutput {
	return i.ToScheduleArrayOutputWithContext(context.Background())
}

func (i ScheduleArray) ToScheduleArrayOutputWithContext(ctx context.Context) ScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleArrayOutput)
}

type ScheduleOutput struct{ *pulumi.OutputState }

func (ScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Schedule)(nil)).Elem()
}

func (o ScheduleOutput) ToScheduleOutput() ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return o
}

func (o ScheduleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Schedule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ScheduleOutput) PolicyList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Schedule) []string { return v.PolicyList }).(pulumi.StringArrayOutput)
}

type ScheduleArrayOutput struct{ *pulumi.OutputState }

func (ScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Schedule)(nil)).Elem()
}

func (o ScheduleArrayOutput) ToScheduleArrayOutput() ScheduleArrayOutput {
	return o
}

func (o ScheduleArrayOutput) ToScheduleArrayOutputWithContext(ctx context.Context) ScheduleArrayOutput {
	return o
}

func (o ScheduleArrayOutput) Index(i pulumi.IntInput) ScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Schedule {
		return vs[0].([]Schedule)[vs[1].(int)]
	}).(ScheduleOutput)
}

type ScheduleResponse struct {
	Name       *string  `pulumi:"name"`
	PolicyList []string `pulumi:"policyList"`
}

type ScheduleResponseOutput struct{ *pulumi.OutputState }

func (ScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleResponse)(nil)).Elem()
}

func (o ScheduleResponseOutput) ToScheduleResponseOutput() ScheduleResponseOutput {
	return o
}

func (o ScheduleResponseOutput) ToScheduleResponseOutputWithContext(ctx context.Context) ScheduleResponseOutput {
	return o
}

func (o ScheduleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ScheduleResponseOutput) PolicyList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScheduleResponse) []string { return v.PolicyList }).(pulumi.StringArrayOutput)
}

type ScheduleResponseArrayOutput struct{ *pulumi.OutputState }

func (ScheduleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleResponse)(nil)).Elem()
}

func (o ScheduleResponseArrayOutput) ToScheduleResponseArrayOutput() ScheduleResponseArrayOutput {
	return o
}

func (o ScheduleResponseArrayOutput) ToScheduleResponseArrayOutputWithContext(ctx context.Context) ScheduleResponseArrayOutput {
	return o
}

func (o ScheduleResponseArrayOutput) Index(i pulumi.IntInput) ScheduleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScheduleResponse {
		return vs[0].([]ScheduleResponse)[vs[1].(int)]
	}).(ScheduleResponseOutput)
}

type Sku struct {
	Name *string `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
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
	Name *string `pulumi:"name"`
	Tier *string `pulumi:"tier"`
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

func init() {
	pulumi.RegisterOutputType(CustomerSecretOutput{})
	pulumi.RegisterOutputType(CustomerSecretArrayOutput{})
	pulumi.RegisterOutputType(CustomerSecretResponseOutput{})
	pulumi.RegisterOutputType(CustomerSecretResponseArrayOutput{})
	pulumi.RegisterOutputType(ScheduleOutput{})
	pulumi.RegisterOutputType(ScheduleArrayOutput{})
	pulumi.RegisterOutputType(ScheduleResponseOutput{})
	pulumi.RegisterOutputType(ScheduleResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
