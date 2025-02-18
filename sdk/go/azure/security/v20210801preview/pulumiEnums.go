


package v20210801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StandardSupportedClouds string

const (
	StandardSupportedCloudsAWS = StandardSupportedClouds("AWS")
	StandardSupportedCloudsGCP = StandardSupportedClouds("GCP")
)

func (StandardSupportedClouds) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardSupportedClouds)(nil)).Elem()
}

func (e StandardSupportedClouds) ToStandardSupportedCloudsOutput() StandardSupportedCloudsOutput {
	return pulumi.ToOutput(e).(StandardSupportedCloudsOutput)
}

func (e StandardSupportedClouds) ToStandardSupportedCloudsOutputWithContext(ctx context.Context) StandardSupportedCloudsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StandardSupportedCloudsOutput)
}

func (e StandardSupportedClouds) ToStandardSupportedCloudsPtrOutput() StandardSupportedCloudsPtrOutput {
	return e.ToStandardSupportedCloudsPtrOutputWithContext(context.Background())
}

func (e StandardSupportedClouds) ToStandardSupportedCloudsPtrOutputWithContext(ctx context.Context) StandardSupportedCloudsPtrOutput {
	return StandardSupportedClouds(e).ToStandardSupportedCloudsOutputWithContext(ctx).ToStandardSupportedCloudsPtrOutputWithContext(ctx)
}

func (e StandardSupportedClouds) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StandardSupportedClouds) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StandardSupportedClouds) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StandardSupportedClouds) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StandardSupportedCloudsOutput struct{ *pulumi.OutputState }

func (StandardSupportedCloudsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StandardSupportedClouds)(nil)).Elem()
}

func (o StandardSupportedCloudsOutput) ToStandardSupportedCloudsOutput() StandardSupportedCloudsOutput {
	return o
}

func (o StandardSupportedCloudsOutput) ToStandardSupportedCloudsOutputWithContext(ctx context.Context) StandardSupportedCloudsOutput {
	return o
}

func (o StandardSupportedCloudsOutput) ToStandardSupportedCloudsPtrOutput() StandardSupportedCloudsPtrOutput {
	return o.ToStandardSupportedCloudsPtrOutputWithContext(context.Background())
}

func (o StandardSupportedCloudsOutput) ToStandardSupportedCloudsPtrOutputWithContext(ctx context.Context) StandardSupportedCloudsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StandardSupportedClouds) *StandardSupportedClouds {
		return &v
	}).(StandardSupportedCloudsPtrOutput)
}

func (o StandardSupportedCloudsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StandardSupportedCloudsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StandardSupportedClouds) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StandardSupportedCloudsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StandardSupportedCloudsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StandardSupportedClouds) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StandardSupportedCloudsPtrOutput struct{ *pulumi.OutputState }

func (StandardSupportedCloudsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StandardSupportedClouds)(nil)).Elem()
}

func (o StandardSupportedCloudsPtrOutput) ToStandardSupportedCloudsPtrOutput() StandardSupportedCloudsPtrOutput {
	return o
}

func (o StandardSupportedCloudsPtrOutput) ToStandardSupportedCloudsPtrOutputWithContext(ctx context.Context) StandardSupportedCloudsPtrOutput {
	return o
}

func (o StandardSupportedCloudsPtrOutput) Elem() StandardSupportedCloudsOutput {
	return o.ApplyT(func(v *StandardSupportedClouds) StandardSupportedClouds {
		if v != nil {
			return *v
		}
		var ret StandardSupportedClouds
		return ret
	}).(StandardSupportedCloudsOutput)
}

func (o StandardSupportedCloudsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StandardSupportedCloudsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StandardSupportedClouds) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StandardSupportedCloudsInput interface {
	pulumi.Input

	ToStandardSupportedCloudsOutput() StandardSupportedCloudsOutput
	ToStandardSupportedCloudsOutputWithContext(context.Context) StandardSupportedCloudsOutput
}

var standardSupportedCloudsPtrType = reflect.TypeOf((**StandardSupportedClouds)(nil)).Elem()

type StandardSupportedCloudsPtrInput interface {
	pulumi.Input

	ToStandardSupportedCloudsPtrOutput() StandardSupportedCloudsPtrOutput
	ToStandardSupportedCloudsPtrOutputWithContext(context.Context) StandardSupportedCloudsPtrOutput
}

type standardSupportedCloudsPtr string

func StandardSupportedCloudsPtr(v string) StandardSupportedCloudsPtrInput {
	return (*standardSupportedCloudsPtr)(&v)
}

func (*standardSupportedCloudsPtr) ElementType() reflect.Type {
	return standardSupportedCloudsPtrType
}

func (in *standardSupportedCloudsPtr) ToStandardSupportedCloudsPtrOutput() StandardSupportedCloudsPtrOutput {
	return pulumi.ToOutput(in).(StandardSupportedCloudsPtrOutput)
}

func (in *standardSupportedCloudsPtr) ToStandardSupportedCloudsPtrOutputWithContext(ctx context.Context) StandardSupportedCloudsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StandardSupportedCloudsPtrOutput)
}





type StandardSupportedCloudsArrayInput interface {
	pulumi.Input

	ToStandardSupportedCloudsArrayOutput() StandardSupportedCloudsArrayOutput
	ToStandardSupportedCloudsArrayOutputWithContext(context.Context) StandardSupportedCloudsArrayOutput
}

type StandardSupportedCloudsArray []StandardSupportedClouds

func (StandardSupportedCloudsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StandardSupportedClouds)(nil)).Elem()
}

func (i StandardSupportedCloudsArray) ToStandardSupportedCloudsArrayOutput() StandardSupportedCloudsArrayOutput {
	return i.ToStandardSupportedCloudsArrayOutputWithContext(context.Background())
}

func (i StandardSupportedCloudsArray) ToStandardSupportedCloudsArrayOutputWithContext(ctx context.Context) StandardSupportedCloudsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StandardSupportedCloudsArrayOutput)
}

type StandardSupportedCloudsArrayOutput struct{ *pulumi.OutputState }

func (StandardSupportedCloudsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StandardSupportedClouds)(nil)).Elem()
}

func (o StandardSupportedCloudsArrayOutput) ToStandardSupportedCloudsArrayOutput() StandardSupportedCloudsArrayOutput {
	return o
}

func (o StandardSupportedCloudsArrayOutput) ToStandardSupportedCloudsArrayOutputWithContext(ctx context.Context) StandardSupportedCloudsArrayOutput {
	return o
}

func (o StandardSupportedCloudsArrayOutput) Index(i pulumi.IntInput) StandardSupportedCloudsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StandardSupportedClouds {
		return vs[0].([]StandardSupportedClouds)[vs[1].(int)]
	}).(StandardSupportedCloudsOutput)
}

func init() {
	pulumi.RegisterOutputType(StandardSupportedCloudsOutput{})
	pulumi.RegisterOutputType(StandardSupportedCloudsPtrOutput{})
	pulumi.RegisterOutputType(StandardSupportedCloudsArrayOutput{})
}
