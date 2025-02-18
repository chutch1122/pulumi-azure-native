


package v20160129

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureSku struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
}





type AzureSkuInput interface {
	pulumi.Input

	ToAzureSkuOutput() AzureSkuOutput
	ToAzureSkuOutputWithContext(context.Context) AzureSkuOutput
}

type AzureSkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (AzureSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSku)(nil)).Elem()
}

func (i AzureSkuArgs) ToAzureSkuOutput() AzureSkuOutput {
	return i.ToAzureSkuOutputWithContext(context.Background())
}

func (i AzureSkuArgs) ToAzureSkuOutputWithContext(ctx context.Context) AzureSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuOutput)
}

func (i AzureSkuArgs) ToAzureSkuPtrOutput() AzureSkuPtrOutput {
	return i.ToAzureSkuPtrOutputWithContext(context.Background())
}

func (i AzureSkuArgs) ToAzureSkuPtrOutputWithContext(ctx context.Context) AzureSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuOutput).ToAzureSkuPtrOutputWithContext(ctx)
}









type AzureSkuPtrInput interface {
	pulumi.Input

	ToAzureSkuPtrOutput() AzureSkuPtrOutput
	ToAzureSkuPtrOutputWithContext(context.Context) AzureSkuPtrOutput
}

type azureSkuPtrType AzureSkuArgs

func AzureSkuPtr(v *AzureSkuArgs) AzureSkuPtrInput {
	return (*azureSkuPtrType)(v)
}

func (*azureSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSku)(nil)).Elem()
}

func (i *azureSkuPtrType) ToAzureSkuPtrOutput() AzureSkuPtrOutput {
	return i.ToAzureSkuPtrOutputWithContext(context.Background())
}

func (i *azureSkuPtrType) ToAzureSkuPtrOutputWithContext(ctx context.Context) AzureSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuPtrOutput)
}

type AzureSkuOutput struct{ *pulumi.OutputState }

func (AzureSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSku)(nil)).Elem()
}

func (o AzureSkuOutput) ToAzureSkuOutput() AzureSkuOutput {
	return o
}

func (o AzureSkuOutput) ToAzureSkuOutputWithContext(ctx context.Context) AzureSkuOutput {
	return o
}

func (o AzureSkuOutput) ToAzureSkuPtrOutput() AzureSkuPtrOutput {
	return o.ToAzureSkuPtrOutputWithContext(context.Background())
}

func (o AzureSkuOutput) ToAzureSkuPtrOutputWithContext(ctx context.Context) AzureSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureSku) *AzureSku {
		return &v
	}).(AzureSkuPtrOutput)
}

func (o AzureSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o AzureSkuOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSku) string { return v.Tier }).(pulumi.StringOutput)
}

type AzureSkuPtrOutput struct{ *pulumi.OutputState }

func (AzureSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSku)(nil)).Elem()
}

func (o AzureSkuPtrOutput) ToAzureSkuPtrOutput() AzureSkuPtrOutput {
	return o
}

func (o AzureSkuPtrOutput) ToAzureSkuPtrOutputWithContext(ctx context.Context) AzureSkuPtrOutput {
	return o
}

func (o AzureSkuPtrOutput) Elem() AzureSkuOutput {
	return o.ApplyT(func(v *AzureSku) AzureSku {
		if v != nil {
			return *v
		}
		var ret AzureSku
		return ret
	}).(AzureSkuOutput)
}

func (o AzureSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o AzureSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSku) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type AzureSkuResponse struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
}

type AzureSkuResponseOutput struct{ *pulumi.OutputState }

func (AzureSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuResponse)(nil)).Elem()
}

func (o AzureSkuResponseOutput) ToAzureSkuResponseOutput() AzureSkuResponseOutput {
	return o
}

func (o AzureSkuResponseOutput) ToAzureSkuResponseOutputWithContext(ctx context.Context) AzureSkuResponseOutput {
	return o
}

func (o AzureSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o AzureSkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type AzureSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSkuResponse)(nil)).Elem()
}

func (o AzureSkuResponsePtrOutput) ToAzureSkuResponsePtrOutput() AzureSkuResponsePtrOutput {
	return o
}

func (o AzureSkuResponsePtrOutput) ToAzureSkuResponsePtrOutputWithContext(ctx context.Context) AzureSkuResponsePtrOutput {
	return o
}

func (o AzureSkuResponsePtrOutput) Elem() AzureSkuResponseOutput {
	return o.ApplyT(func(v *AzureSkuResponse) AzureSkuResponse {
		if v != nil {
			return *v
		}
		var ret AzureSkuResponse
		return ret
	}).(AzureSkuResponseOutput)
}

func (o AzureSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o AzureSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureSkuOutput{})
	pulumi.RegisterOutputType(AzureSkuPtrOutput{})
	pulumi.RegisterOutputType(AzureSkuResponseOutput{})
	pulumi.RegisterOutputType(AzureSkuResponsePtrOutput{})
}
