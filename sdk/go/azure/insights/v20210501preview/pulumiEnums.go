


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComparisonOperationType string

const (
	ComparisonOperationTypeEquals             = ComparisonOperationType("Equals")
	ComparisonOperationTypeNotEquals          = ComparisonOperationType("NotEquals")
	ComparisonOperationTypeGreaterThan        = ComparisonOperationType("GreaterThan")
	ComparisonOperationTypeGreaterThanOrEqual = ComparisonOperationType("GreaterThanOrEqual")
	ComparisonOperationTypeLessThan           = ComparisonOperationType("LessThan")
	ComparisonOperationTypeLessThanOrEqual    = ComparisonOperationType("LessThanOrEqual")
)

func (ComparisonOperationType) ElementType() reflect.Type {
	return reflect.TypeOf((*ComparisonOperationType)(nil)).Elem()
}

func (e ComparisonOperationType) ToComparisonOperationTypeOutput() ComparisonOperationTypeOutput {
	return pulumi.ToOutput(e).(ComparisonOperationTypeOutput)
}

func (e ComparisonOperationType) ToComparisonOperationTypeOutputWithContext(ctx context.Context) ComparisonOperationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ComparisonOperationTypeOutput)
}

func (e ComparisonOperationType) ToComparisonOperationTypePtrOutput() ComparisonOperationTypePtrOutput {
	return e.ToComparisonOperationTypePtrOutputWithContext(context.Background())
}

func (e ComparisonOperationType) ToComparisonOperationTypePtrOutputWithContext(ctx context.Context) ComparisonOperationTypePtrOutput {
	return ComparisonOperationType(e).ToComparisonOperationTypeOutputWithContext(ctx).ToComparisonOperationTypePtrOutputWithContext(ctx)
}

func (e ComparisonOperationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComparisonOperationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComparisonOperationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ComparisonOperationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ComparisonOperationTypeOutput struct{ *pulumi.OutputState }

func (ComparisonOperationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComparisonOperationType)(nil)).Elem()
}

func (o ComparisonOperationTypeOutput) ToComparisonOperationTypeOutput() ComparisonOperationTypeOutput {
	return o
}

func (o ComparisonOperationTypeOutput) ToComparisonOperationTypeOutputWithContext(ctx context.Context) ComparisonOperationTypeOutput {
	return o
}

func (o ComparisonOperationTypeOutput) ToComparisonOperationTypePtrOutput() ComparisonOperationTypePtrOutput {
	return o.ToComparisonOperationTypePtrOutputWithContext(context.Background())
}

func (o ComparisonOperationTypeOutput) ToComparisonOperationTypePtrOutputWithContext(ctx context.Context) ComparisonOperationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComparisonOperationType) *ComparisonOperationType {
		return &v
	}).(ComparisonOperationTypePtrOutput)
}

func (o ComparisonOperationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ComparisonOperationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComparisonOperationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ComparisonOperationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComparisonOperationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComparisonOperationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ComparisonOperationTypePtrOutput struct{ *pulumi.OutputState }

func (ComparisonOperationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComparisonOperationType)(nil)).Elem()
}

func (o ComparisonOperationTypePtrOutput) ToComparisonOperationTypePtrOutput() ComparisonOperationTypePtrOutput {
	return o
}

func (o ComparisonOperationTypePtrOutput) ToComparisonOperationTypePtrOutputWithContext(ctx context.Context) ComparisonOperationTypePtrOutput {
	return o
}

func (o ComparisonOperationTypePtrOutput) Elem() ComparisonOperationTypeOutput {
	return o.ApplyT(func(v *ComparisonOperationType) ComparisonOperationType {
		if v != nil {
			return *v
		}
		var ret ComparisonOperationType
		return ret
	}).(ComparisonOperationTypeOutput)
}

func (o ComparisonOperationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComparisonOperationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ComparisonOperationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ComparisonOperationTypeInput interface {
	pulumi.Input

	ToComparisonOperationTypeOutput() ComparisonOperationTypeOutput
	ToComparisonOperationTypeOutputWithContext(context.Context) ComparisonOperationTypeOutput
}

var comparisonOperationTypePtrType = reflect.TypeOf((**ComparisonOperationType)(nil)).Elem()

type ComparisonOperationTypePtrInput interface {
	pulumi.Input

	ToComparisonOperationTypePtrOutput() ComparisonOperationTypePtrOutput
	ToComparisonOperationTypePtrOutputWithContext(context.Context) ComparisonOperationTypePtrOutput
}

type comparisonOperationTypePtr string

func ComparisonOperationTypePtr(v string) ComparisonOperationTypePtrInput {
	return (*comparisonOperationTypePtr)(&v)
}

func (*comparisonOperationTypePtr) ElementType() reflect.Type {
	return comparisonOperationTypePtrType
}

func (in *comparisonOperationTypePtr) ToComparisonOperationTypePtrOutput() ComparisonOperationTypePtrOutput {
	return pulumi.ToOutput(in).(ComparisonOperationTypePtrOutput)
}

func (in *comparisonOperationTypePtr) ToComparisonOperationTypePtrOutputWithContext(ctx context.Context) ComparisonOperationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ComparisonOperationTypePtrOutput)
}

type MetricStatisticType string

const (
	MetricStatisticTypeAverage = MetricStatisticType("Average")
	MetricStatisticTypeMin     = MetricStatisticType("Min")
	MetricStatisticTypeMax     = MetricStatisticType("Max")
	MetricStatisticTypeSum     = MetricStatisticType("Sum")
	MetricStatisticTypeCount   = MetricStatisticType("Count")
)

func (MetricStatisticType) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricStatisticType)(nil)).Elem()
}

func (e MetricStatisticType) ToMetricStatisticTypeOutput() MetricStatisticTypeOutput {
	return pulumi.ToOutput(e).(MetricStatisticTypeOutput)
}

func (e MetricStatisticType) ToMetricStatisticTypeOutputWithContext(ctx context.Context) MetricStatisticTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MetricStatisticTypeOutput)
}

func (e MetricStatisticType) ToMetricStatisticTypePtrOutput() MetricStatisticTypePtrOutput {
	return e.ToMetricStatisticTypePtrOutputWithContext(context.Background())
}

func (e MetricStatisticType) ToMetricStatisticTypePtrOutputWithContext(ctx context.Context) MetricStatisticTypePtrOutput {
	return MetricStatisticType(e).ToMetricStatisticTypeOutputWithContext(ctx).ToMetricStatisticTypePtrOutputWithContext(ctx)
}

func (e MetricStatisticType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MetricStatisticType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MetricStatisticType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MetricStatisticType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MetricStatisticTypeOutput struct{ *pulumi.OutputState }

func (MetricStatisticTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricStatisticType)(nil)).Elem()
}

func (o MetricStatisticTypeOutput) ToMetricStatisticTypeOutput() MetricStatisticTypeOutput {
	return o
}

func (o MetricStatisticTypeOutput) ToMetricStatisticTypeOutputWithContext(ctx context.Context) MetricStatisticTypeOutput {
	return o
}

func (o MetricStatisticTypeOutput) ToMetricStatisticTypePtrOutput() MetricStatisticTypePtrOutput {
	return o.ToMetricStatisticTypePtrOutputWithContext(context.Background())
}

func (o MetricStatisticTypeOutput) ToMetricStatisticTypePtrOutputWithContext(ctx context.Context) MetricStatisticTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MetricStatisticType) *MetricStatisticType {
		return &v
	}).(MetricStatisticTypePtrOutput)
}

func (o MetricStatisticTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MetricStatisticTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MetricStatisticType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MetricStatisticTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MetricStatisticTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MetricStatisticType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MetricStatisticTypePtrOutput struct{ *pulumi.OutputState }

func (MetricStatisticTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricStatisticType)(nil)).Elem()
}

func (o MetricStatisticTypePtrOutput) ToMetricStatisticTypePtrOutput() MetricStatisticTypePtrOutput {
	return o
}

func (o MetricStatisticTypePtrOutput) ToMetricStatisticTypePtrOutputWithContext(ctx context.Context) MetricStatisticTypePtrOutput {
	return o
}

func (o MetricStatisticTypePtrOutput) Elem() MetricStatisticTypeOutput {
	return o.ApplyT(func(v *MetricStatisticType) MetricStatisticType {
		if v != nil {
			return *v
		}
		var ret MetricStatisticType
		return ret
	}).(MetricStatisticTypeOutput)
}

func (o MetricStatisticTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MetricStatisticTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MetricStatisticType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MetricStatisticTypeInput interface {
	pulumi.Input

	ToMetricStatisticTypeOutput() MetricStatisticTypeOutput
	ToMetricStatisticTypeOutputWithContext(context.Context) MetricStatisticTypeOutput
}

var metricStatisticTypePtrType = reflect.TypeOf((**MetricStatisticType)(nil)).Elem()

type MetricStatisticTypePtrInput interface {
	pulumi.Input

	ToMetricStatisticTypePtrOutput() MetricStatisticTypePtrOutput
	ToMetricStatisticTypePtrOutputWithContext(context.Context) MetricStatisticTypePtrOutput
}

type metricStatisticTypePtr string

func MetricStatisticTypePtr(v string) MetricStatisticTypePtrInput {
	return (*metricStatisticTypePtr)(&v)
}

func (*metricStatisticTypePtr) ElementType() reflect.Type {
	return metricStatisticTypePtrType
}

func (in *metricStatisticTypePtr) ToMetricStatisticTypePtrOutput() MetricStatisticTypePtrOutput {
	return pulumi.ToOutput(in).(MetricStatisticTypePtrOutput)
}

func (in *metricStatisticTypePtr) ToMetricStatisticTypePtrOutputWithContext(ctx context.Context) MetricStatisticTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MetricStatisticTypePtrOutput)
}

type OperationType string

const (
	OperationTypeScale = OperationType("Scale")
)

func (OperationType) ElementType() reflect.Type {
	return reflect.TypeOf((*OperationType)(nil)).Elem()
}

func (e OperationType) ToOperationTypeOutput() OperationTypeOutput {
	return pulumi.ToOutput(e).(OperationTypeOutput)
}

func (e OperationType) ToOperationTypeOutputWithContext(ctx context.Context) OperationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperationTypeOutput)
}

func (e OperationType) ToOperationTypePtrOutput() OperationTypePtrOutput {
	return e.ToOperationTypePtrOutputWithContext(context.Background())
}

func (e OperationType) ToOperationTypePtrOutputWithContext(ctx context.Context) OperationTypePtrOutput {
	return OperationType(e).ToOperationTypeOutputWithContext(ctx).ToOperationTypePtrOutputWithContext(ctx)
}

func (e OperationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OperationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperationTypeOutput struct{ *pulumi.OutputState }

func (OperationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperationType)(nil)).Elem()
}

func (o OperationTypeOutput) ToOperationTypeOutput() OperationTypeOutput {
	return o
}

func (o OperationTypeOutput) ToOperationTypeOutputWithContext(ctx context.Context) OperationTypeOutput {
	return o
}

func (o OperationTypeOutput) ToOperationTypePtrOutput() OperationTypePtrOutput {
	return o.ToOperationTypePtrOutputWithContext(context.Background())
}

func (o OperationTypeOutput) ToOperationTypePtrOutputWithContext(ctx context.Context) OperationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperationType) *OperationType {
		return &v
	}).(OperationTypePtrOutput)
}

func (o OperationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperationTypePtrOutput struct{ *pulumi.OutputState }

func (OperationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperationType)(nil)).Elem()
}

func (o OperationTypePtrOutput) ToOperationTypePtrOutput() OperationTypePtrOutput {
	return o
}

func (o OperationTypePtrOutput) ToOperationTypePtrOutputWithContext(ctx context.Context) OperationTypePtrOutput {
	return o
}

func (o OperationTypePtrOutput) Elem() OperationTypeOutput {
	return o.ApplyT(func(v *OperationType) OperationType {
		if v != nil {
			return *v
		}
		var ret OperationType
		return ret
	}).(OperationTypeOutput)
}

func (o OperationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OperationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OperationTypeInput interface {
	pulumi.Input

	ToOperationTypeOutput() OperationTypeOutput
	ToOperationTypeOutputWithContext(context.Context) OperationTypeOutput
}

var operationTypePtrType = reflect.TypeOf((**OperationType)(nil)).Elem()

type OperationTypePtrInput interface {
	pulumi.Input

	ToOperationTypePtrOutput() OperationTypePtrOutput
	ToOperationTypePtrOutputWithContext(context.Context) OperationTypePtrOutput
}

type operationTypePtr string

func OperationTypePtr(v string) OperationTypePtrInput {
	return (*operationTypePtr)(&v)
}

func (*operationTypePtr) ElementType() reflect.Type {
	return operationTypePtrType
}

func (in *operationTypePtr) ToOperationTypePtrOutput() OperationTypePtrOutput {
	return pulumi.ToOutput(in).(OperationTypePtrOutput)
}

func (in *operationTypePtr) ToOperationTypePtrOutputWithContext(ctx context.Context) OperationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperationTypePtrOutput)
}

type PredictiveAutoscalePolicyScaleMode string

const (
	PredictiveAutoscalePolicyScaleModeDisabled     = PredictiveAutoscalePolicyScaleMode("Disabled")
	PredictiveAutoscalePolicyScaleModeForecastOnly = PredictiveAutoscalePolicyScaleMode("ForecastOnly")
	PredictiveAutoscalePolicyScaleModeEnabled      = PredictiveAutoscalePolicyScaleMode("Enabled")
)

func (PredictiveAutoscalePolicyScaleMode) ElementType() reflect.Type {
	return reflect.TypeOf((*PredictiveAutoscalePolicyScaleMode)(nil)).Elem()
}

func (e PredictiveAutoscalePolicyScaleMode) ToPredictiveAutoscalePolicyScaleModeOutput() PredictiveAutoscalePolicyScaleModeOutput {
	return pulumi.ToOutput(e).(PredictiveAutoscalePolicyScaleModeOutput)
}

func (e PredictiveAutoscalePolicyScaleMode) ToPredictiveAutoscalePolicyScaleModeOutputWithContext(ctx context.Context) PredictiveAutoscalePolicyScaleModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PredictiveAutoscalePolicyScaleModeOutput)
}

func (e PredictiveAutoscalePolicyScaleMode) ToPredictiveAutoscalePolicyScaleModePtrOutput() PredictiveAutoscalePolicyScaleModePtrOutput {
	return e.ToPredictiveAutoscalePolicyScaleModePtrOutputWithContext(context.Background())
}

func (e PredictiveAutoscalePolicyScaleMode) ToPredictiveAutoscalePolicyScaleModePtrOutputWithContext(ctx context.Context) PredictiveAutoscalePolicyScaleModePtrOutput {
	return PredictiveAutoscalePolicyScaleMode(e).ToPredictiveAutoscalePolicyScaleModeOutputWithContext(ctx).ToPredictiveAutoscalePolicyScaleModePtrOutputWithContext(ctx)
}

func (e PredictiveAutoscalePolicyScaleMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PredictiveAutoscalePolicyScaleMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PredictiveAutoscalePolicyScaleMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PredictiveAutoscalePolicyScaleMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PredictiveAutoscalePolicyScaleModeOutput struct{ *pulumi.OutputState }

func (PredictiveAutoscalePolicyScaleModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PredictiveAutoscalePolicyScaleMode)(nil)).Elem()
}

func (o PredictiveAutoscalePolicyScaleModeOutput) ToPredictiveAutoscalePolicyScaleModeOutput() PredictiveAutoscalePolicyScaleModeOutput {
	return o
}

func (o PredictiveAutoscalePolicyScaleModeOutput) ToPredictiveAutoscalePolicyScaleModeOutputWithContext(ctx context.Context) PredictiveAutoscalePolicyScaleModeOutput {
	return o
}

func (o PredictiveAutoscalePolicyScaleModeOutput) ToPredictiveAutoscalePolicyScaleModePtrOutput() PredictiveAutoscalePolicyScaleModePtrOutput {
	return o.ToPredictiveAutoscalePolicyScaleModePtrOutputWithContext(context.Background())
}

func (o PredictiveAutoscalePolicyScaleModeOutput) ToPredictiveAutoscalePolicyScaleModePtrOutputWithContext(ctx context.Context) PredictiveAutoscalePolicyScaleModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PredictiveAutoscalePolicyScaleMode) *PredictiveAutoscalePolicyScaleMode {
		return &v
	}).(PredictiveAutoscalePolicyScaleModePtrOutput)
}

func (o PredictiveAutoscalePolicyScaleModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PredictiveAutoscalePolicyScaleModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PredictiveAutoscalePolicyScaleMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PredictiveAutoscalePolicyScaleModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PredictiveAutoscalePolicyScaleModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PredictiveAutoscalePolicyScaleMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PredictiveAutoscalePolicyScaleModePtrOutput struct{ *pulumi.OutputState }

func (PredictiveAutoscalePolicyScaleModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PredictiveAutoscalePolicyScaleMode)(nil)).Elem()
}

func (o PredictiveAutoscalePolicyScaleModePtrOutput) ToPredictiveAutoscalePolicyScaleModePtrOutput() PredictiveAutoscalePolicyScaleModePtrOutput {
	return o
}

func (o PredictiveAutoscalePolicyScaleModePtrOutput) ToPredictiveAutoscalePolicyScaleModePtrOutputWithContext(ctx context.Context) PredictiveAutoscalePolicyScaleModePtrOutput {
	return o
}

func (o PredictiveAutoscalePolicyScaleModePtrOutput) Elem() PredictiveAutoscalePolicyScaleModeOutput {
	return o.ApplyT(func(v *PredictiveAutoscalePolicyScaleMode) PredictiveAutoscalePolicyScaleMode {
		if v != nil {
			return *v
		}
		var ret PredictiveAutoscalePolicyScaleMode
		return ret
	}).(PredictiveAutoscalePolicyScaleModeOutput)
}

func (o PredictiveAutoscalePolicyScaleModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PredictiveAutoscalePolicyScaleModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PredictiveAutoscalePolicyScaleMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PredictiveAutoscalePolicyScaleModeInput interface {
	pulumi.Input

	ToPredictiveAutoscalePolicyScaleModeOutput() PredictiveAutoscalePolicyScaleModeOutput
	ToPredictiveAutoscalePolicyScaleModeOutputWithContext(context.Context) PredictiveAutoscalePolicyScaleModeOutput
}

var predictiveAutoscalePolicyScaleModePtrType = reflect.TypeOf((**PredictiveAutoscalePolicyScaleMode)(nil)).Elem()

type PredictiveAutoscalePolicyScaleModePtrInput interface {
	pulumi.Input

	ToPredictiveAutoscalePolicyScaleModePtrOutput() PredictiveAutoscalePolicyScaleModePtrOutput
	ToPredictiveAutoscalePolicyScaleModePtrOutputWithContext(context.Context) PredictiveAutoscalePolicyScaleModePtrOutput
}

type predictiveAutoscalePolicyScaleModePtr string

func PredictiveAutoscalePolicyScaleModePtr(v string) PredictiveAutoscalePolicyScaleModePtrInput {
	return (*predictiveAutoscalePolicyScaleModePtr)(&v)
}

func (*predictiveAutoscalePolicyScaleModePtr) ElementType() reflect.Type {
	return predictiveAutoscalePolicyScaleModePtrType
}

func (in *predictiveAutoscalePolicyScaleModePtr) ToPredictiveAutoscalePolicyScaleModePtrOutput() PredictiveAutoscalePolicyScaleModePtrOutput {
	return pulumi.ToOutput(in).(PredictiveAutoscalePolicyScaleModePtrOutput)
}

func (in *predictiveAutoscalePolicyScaleModePtr) ToPredictiveAutoscalePolicyScaleModePtrOutputWithContext(ctx context.Context) PredictiveAutoscalePolicyScaleModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PredictiveAutoscalePolicyScaleModePtrOutput)
}

type RecurrenceFrequency string

const (
	RecurrenceFrequencyNone   = RecurrenceFrequency("None")
	RecurrenceFrequencySecond = RecurrenceFrequency("Second")
	RecurrenceFrequencyMinute = RecurrenceFrequency("Minute")
	RecurrenceFrequencyHour   = RecurrenceFrequency("Hour")
	RecurrenceFrequencyDay    = RecurrenceFrequency("Day")
	RecurrenceFrequencyWeek   = RecurrenceFrequency("Week")
	RecurrenceFrequencyMonth  = RecurrenceFrequency("Month")
	RecurrenceFrequencyYear   = RecurrenceFrequency("Year")
)

func (RecurrenceFrequency) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceFrequency)(nil)).Elem()
}

func (e RecurrenceFrequency) ToRecurrenceFrequencyOutput() RecurrenceFrequencyOutput {
	return pulumi.ToOutput(e).(RecurrenceFrequencyOutput)
}

func (e RecurrenceFrequency) ToRecurrenceFrequencyOutputWithContext(ctx context.Context) RecurrenceFrequencyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RecurrenceFrequencyOutput)
}

func (e RecurrenceFrequency) ToRecurrenceFrequencyPtrOutput() RecurrenceFrequencyPtrOutput {
	return e.ToRecurrenceFrequencyPtrOutputWithContext(context.Background())
}

func (e RecurrenceFrequency) ToRecurrenceFrequencyPtrOutputWithContext(ctx context.Context) RecurrenceFrequencyPtrOutput {
	return RecurrenceFrequency(e).ToRecurrenceFrequencyOutputWithContext(ctx).ToRecurrenceFrequencyPtrOutputWithContext(ctx)
}

func (e RecurrenceFrequency) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecurrenceFrequency) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecurrenceFrequency) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RecurrenceFrequency) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RecurrenceFrequencyOutput struct{ *pulumi.OutputState }

func (RecurrenceFrequencyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceFrequency)(nil)).Elem()
}

func (o RecurrenceFrequencyOutput) ToRecurrenceFrequencyOutput() RecurrenceFrequencyOutput {
	return o
}

func (o RecurrenceFrequencyOutput) ToRecurrenceFrequencyOutputWithContext(ctx context.Context) RecurrenceFrequencyOutput {
	return o
}

func (o RecurrenceFrequencyOutput) ToRecurrenceFrequencyPtrOutput() RecurrenceFrequencyPtrOutput {
	return o.ToRecurrenceFrequencyPtrOutputWithContext(context.Background())
}

func (o RecurrenceFrequencyOutput) ToRecurrenceFrequencyPtrOutputWithContext(ctx context.Context) RecurrenceFrequencyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecurrenceFrequency) *RecurrenceFrequency {
		return &v
	}).(RecurrenceFrequencyPtrOutput)
}

func (o RecurrenceFrequencyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RecurrenceFrequencyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecurrenceFrequency) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RecurrenceFrequencyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecurrenceFrequencyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecurrenceFrequency) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RecurrenceFrequencyPtrOutput struct{ *pulumi.OutputState }

func (RecurrenceFrequencyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecurrenceFrequency)(nil)).Elem()
}

func (o RecurrenceFrequencyPtrOutput) ToRecurrenceFrequencyPtrOutput() RecurrenceFrequencyPtrOutput {
	return o
}

func (o RecurrenceFrequencyPtrOutput) ToRecurrenceFrequencyPtrOutputWithContext(ctx context.Context) RecurrenceFrequencyPtrOutput {
	return o
}

func (o RecurrenceFrequencyPtrOutput) Elem() RecurrenceFrequencyOutput {
	return o.ApplyT(func(v *RecurrenceFrequency) RecurrenceFrequency {
		if v != nil {
			return *v
		}
		var ret RecurrenceFrequency
		return ret
	}).(RecurrenceFrequencyOutput)
}

func (o RecurrenceFrequencyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecurrenceFrequencyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RecurrenceFrequency) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RecurrenceFrequencyInput interface {
	pulumi.Input

	ToRecurrenceFrequencyOutput() RecurrenceFrequencyOutput
	ToRecurrenceFrequencyOutputWithContext(context.Context) RecurrenceFrequencyOutput
}

var recurrenceFrequencyPtrType = reflect.TypeOf((**RecurrenceFrequency)(nil)).Elem()

type RecurrenceFrequencyPtrInput interface {
	pulumi.Input

	ToRecurrenceFrequencyPtrOutput() RecurrenceFrequencyPtrOutput
	ToRecurrenceFrequencyPtrOutputWithContext(context.Context) RecurrenceFrequencyPtrOutput
}

type recurrenceFrequencyPtr string

func RecurrenceFrequencyPtr(v string) RecurrenceFrequencyPtrInput {
	return (*recurrenceFrequencyPtr)(&v)
}

func (*recurrenceFrequencyPtr) ElementType() reflect.Type {
	return recurrenceFrequencyPtrType
}

func (in *recurrenceFrequencyPtr) ToRecurrenceFrequencyPtrOutput() RecurrenceFrequencyPtrOutput {
	return pulumi.ToOutput(in).(RecurrenceFrequencyPtrOutput)
}

func (in *recurrenceFrequencyPtr) ToRecurrenceFrequencyPtrOutputWithContext(ctx context.Context) RecurrenceFrequencyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RecurrenceFrequencyPtrOutput)
}

type ScaleDirection string

const (
	ScaleDirectionNone     = ScaleDirection("None")
	ScaleDirectionIncrease = ScaleDirection("Increase")
	ScaleDirectionDecrease = ScaleDirection("Decrease")
)

func (ScaleDirection) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleDirection)(nil)).Elem()
}

func (e ScaleDirection) ToScaleDirectionOutput() ScaleDirectionOutput {
	return pulumi.ToOutput(e).(ScaleDirectionOutput)
}

func (e ScaleDirection) ToScaleDirectionOutputWithContext(ctx context.Context) ScaleDirectionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScaleDirectionOutput)
}

func (e ScaleDirection) ToScaleDirectionPtrOutput() ScaleDirectionPtrOutput {
	return e.ToScaleDirectionPtrOutputWithContext(context.Background())
}

func (e ScaleDirection) ToScaleDirectionPtrOutputWithContext(ctx context.Context) ScaleDirectionPtrOutput {
	return ScaleDirection(e).ToScaleDirectionOutputWithContext(ctx).ToScaleDirectionPtrOutputWithContext(ctx)
}

func (e ScaleDirection) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScaleDirection) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScaleDirection) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScaleDirection) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScaleDirectionOutput struct{ *pulumi.OutputState }

func (ScaleDirectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleDirection)(nil)).Elem()
}

func (o ScaleDirectionOutput) ToScaleDirectionOutput() ScaleDirectionOutput {
	return o
}

func (o ScaleDirectionOutput) ToScaleDirectionOutputWithContext(ctx context.Context) ScaleDirectionOutput {
	return o
}

func (o ScaleDirectionOutput) ToScaleDirectionPtrOutput() ScaleDirectionPtrOutput {
	return o.ToScaleDirectionPtrOutputWithContext(context.Background())
}

func (o ScaleDirectionOutput) ToScaleDirectionPtrOutputWithContext(ctx context.Context) ScaleDirectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScaleDirection) *ScaleDirection {
		return &v
	}).(ScaleDirectionPtrOutput)
}

func (o ScaleDirectionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScaleDirectionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScaleDirection) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScaleDirectionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScaleDirectionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScaleDirection) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScaleDirectionPtrOutput struct{ *pulumi.OutputState }

func (ScaleDirectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleDirection)(nil)).Elem()
}

func (o ScaleDirectionPtrOutput) ToScaleDirectionPtrOutput() ScaleDirectionPtrOutput {
	return o
}

func (o ScaleDirectionPtrOutput) ToScaleDirectionPtrOutputWithContext(ctx context.Context) ScaleDirectionPtrOutput {
	return o
}

func (o ScaleDirectionPtrOutput) Elem() ScaleDirectionOutput {
	return o.ApplyT(func(v *ScaleDirection) ScaleDirection {
		if v != nil {
			return *v
		}
		var ret ScaleDirection
		return ret
	}).(ScaleDirectionOutput)
}

func (o ScaleDirectionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScaleDirectionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScaleDirection) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ScaleDirectionInput interface {
	pulumi.Input

	ToScaleDirectionOutput() ScaleDirectionOutput
	ToScaleDirectionOutputWithContext(context.Context) ScaleDirectionOutput
}

var scaleDirectionPtrType = reflect.TypeOf((**ScaleDirection)(nil)).Elem()

type ScaleDirectionPtrInput interface {
	pulumi.Input

	ToScaleDirectionPtrOutput() ScaleDirectionPtrOutput
	ToScaleDirectionPtrOutputWithContext(context.Context) ScaleDirectionPtrOutput
}

type scaleDirectionPtr string

func ScaleDirectionPtr(v string) ScaleDirectionPtrInput {
	return (*scaleDirectionPtr)(&v)
}

func (*scaleDirectionPtr) ElementType() reflect.Type {
	return scaleDirectionPtrType
}

func (in *scaleDirectionPtr) ToScaleDirectionPtrOutput() ScaleDirectionPtrOutput {
	return pulumi.ToOutput(in).(ScaleDirectionPtrOutput)
}

func (in *scaleDirectionPtr) ToScaleDirectionPtrOutputWithContext(ctx context.Context) ScaleDirectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScaleDirectionPtrOutput)
}

type ScaleRuleMetricDimensionOperationType string

const (
	ScaleRuleMetricDimensionOperationTypeEquals    = ScaleRuleMetricDimensionOperationType("Equals")
	ScaleRuleMetricDimensionOperationTypeNotEquals = ScaleRuleMetricDimensionOperationType("NotEquals")
)

type ScaleType string

const (
	ScaleTypeChangeCount             = ScaleType("ChangeCount")
	ScaleTypePercentChangeCount      = ScaleType("PercentChangeCount")
	ScaleTypeExactCount              = ScaleType("ExactCount")
	ScaleTypeServiceAllowedNextValue = ScaleType("ServiceAllowedNextValue")
)

func (ScaleType) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleType)(nil)).Elem()
}

func (e ScaleType) ToScaleTypeOutput() ScaleTypeOutput {
	return pulumi.ToOutput(e).(ScaleTypeOutput)
}

func (e ScaleType) ToScaleTypeOutputWithContext(ctx context.Context) ScaleTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScaleTypeOutput)
}

func (e ScaleType) ToScaleTypePtrOutput() ScaleTypePtrOutput {
	return e.ToScaleTypePtrOutputWithContext(context.Background())
}

func (e ScaleType) ToScaleTypePtrOutputWithContext(ctx context.Context) ScaleTypePtrOutput {
	return ScaleType(e).ToScaleTypeOutputWithContext(ctx).ToScaleTypePtrOutputWithContext(ctx)
}

func (e ScaleType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScaleType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScaleType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScaleType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScaleTypeOutput struct{ *pulumi.OutputState }

func (ScaleTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleType)(nil)).Elem()
}

func (o ScaleTypeOutput) ToScaleTypeOutput() ScaleTypeOutput {
	return o
}

func (o ScaleTypeOutput) ToScaleTypeOutputWithContext(ctx context.Context) ScaleTypeOutput {
	return o
}

func (o ScaleTypeOutput) ToScaleTypePtrOutput() ScaleTypePtrOutput {
	return o.ToScaleTypePtrOutputWithContext(context.Background())
}

func (o ScaleTypeOutput) ToScaleTypePtrOutputWithContext(ctx context.Context) ScaleTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScaleType) *ScaleType {
		return &v
	}).(ScaleTypePtrOutput)
}

func (o ScaleTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScaleTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScaleType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScaleTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScaleTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScaleType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScaleTypePtrOutput struct{ *pulumi.OutputState }

func (ScaleTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleType)(nil)).Elem()
}

func (o ScaleTypePtrOutput) ToScaleTypePtrOutput() ScaleTypePtrOutput {
	return o
}

func (o ScaleTypePtrOutput) ToScaleTypePtrOutputWithContext(ctx context.Context) ScaleTypePtrOutput {
	return o
}

func (o ScaleTypePtrOutput) Elem() ScaleTypeOutput {
	return o.ApplyT(func(v *ScaleType) ScaleType {
		if v != nil {
			return *v
		}
		var ret ScaleType
		return ret
	}).(ScaleTypeOutput)
}

func (o ScaleTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScaleTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScaleType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ScaleTypeInput interface {
	pulumi.Input

	ToScaleTypeOutput() ScaleTypeOutput
	ToScaleTypeOutputWithContext(context.Context) ScaleTypeOutput
}

var scaleTypePtrType = reflect.TypeOf((**ScaleType)(nil)).Elem()

type ScaleTypePtrInput interface {
	pulumi.Input

	ToScaleTypePtrOutput() ScaleTypePtrOutput
	ToScaleTypePtrOutputWithContext(context.Context) ScaleTypePtrOutput
}

type scaleTypePtr string

func ScaleTypePtr(v string) ScaleTypePtrInput {
	return (*scaleTypePtr)(&v)
}

func (*scaleTypePtr) ElementType() reflect.Type {
	return scaleTypePtrType
}

func (in *scaleTypePtr) ToScaleTypePtrOutput() ScaleTypePtrOutput {
	return pulumi.ToOutput(in).(ScaleTypePtrOutput)
}

func (in *scaleTypePtr) ToScaleTypePtrOutputWithContext(ctx context.Context) ScaleTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScaleTypePtrOutput)
}

type TimeAggregationType string

const (
	TimeAggregationTypeAverage = TimeAggregationType("Average")
	TimeAggregationTypeMinimum = TimeAggregationType("Minimum")
	TimeAggregationTypeMaximum = TimeAggregationType("Maximum")
	TimeAggregationTypeTotal   = TimeAggregationType("Total")
	TimeAggregationTypeCount   = TimeAggregationType("Count")
	TimeAggregationTypeLast    = TimeAggregationType("Last")
)

func (TimeAggregationType) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeAggregationType)(nil)).Elem()
}

func (e TimeAggregationType) ToTimeAggregationTypeOutput() TimeAggregationTypeOutput {
	return pulumi.ToOutput(e).(TimeAggregationTypeOutput)
}

func (e TimeAggregationType) ToTimeAggregationTypeOutputWithContext(ctx context.Context) TimeAggregationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TimeAggregationTypeOutput)
}

func (e TimeAggregationType) ToTimeAggregationTypePtrOutput() TimeAggregationTypePtrOutput {
	return e.ToTimeAggregationTypePtrOutputWithContext(context.Background())
}

func (e TimeAggregationType) ToTimeAggregationTypePtrOutputWithContext(ctx context.Context) TimeAggregationTypePtrOutput {
	return TimeAggregationType(e).ToTimeAggregationTypeOutputWithContext(ctx).ToTimeAggregationTypePtrOutputWithContext(ctx)
}

func (e TimeAggregationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TimeAggregationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TimeAggregationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TimeAggregationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TimeAggregationTypeOutput struct{ *pulumi.OutputState }

func (TimeAggregationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeAggregationType)(nil)).Elem()
}

func (o TimeAggregationTypeOutput) ToTimeAggregationTypeOutput() TimeAggregationTypeOutput {
	return o
}

func (o TimeAggregationTypeOutput) ToTimeAggregationTypeOutputWithContext(ctx context.Context) TimeAggregationTypeOutput {
	return o
}

func (o TimeAggregationTypeOutput) ToTimeAggregationTypePtrOutput() TimeAggregationTypePtrOutput {
	return o.ToTimeAggregationTypePtrOutputWithContext(context.Background())
}

func (o TimeAggregationTypeOutput) ToTimeAggregationTypePtrOutputWithContext(ctx context.Context) TimeAggregationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TimeAggregationType) *TimeAggregationType {
		return &v
	}).(TimeAggregationTypePtrOutput)
}

func (o TimeAggregationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TimeAggregationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TimeAggregationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TimeAggregationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TimeAggregationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TimeAggregationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TimeAggregationTypePtrOutput struct{ *pulumi.OutputState }

func (TimeAggregationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeAggregationType)(nil)).Elem()
}

func (o TimeAggregationTypePtrOutput) ToTimeAggregationTypePtrOutput() TimeAggregationTypePtrOutput {
	return o
}

func (o TimeAggregationTypePtrOutput) ToTimeAggregationTypePtrOutputWithContext(ctx context.Context) TimeAggregationTypePtrOutput {
	return o
}

func (o TimeAggregationTypePtrOutput) Elem() TimeAggregationTypeOutput {
	return o.ApplyT(func(v *TimeAggregationType) TimeAggregationType {
		if v != nil {
			return *v
		}
		var ret TimeAggregationType
		return ret
	}).(TimeAggregationTypeOutput)
}

func (o TimeAggregationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TimeAggregationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TimeAggregationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TimeAggregationTypeInput interface {
	pulumi.Input

	ToTimeAggregationTypeOutput() TimeAggregationTypeOutput
	ToTimeAggregationTypeOutputWithContext(context.Context) TimeAggregationTypeOutput
}

var timeAggregationTypePtrType = reflect.TypeOf((**TimeAggregationType)(nil)).Elem()

type TimeAggregationTypePtrInput interface {
	pulumi.Input

	ToTimeAggregationTypePtrOutput() TimeAggregationTypePtrOutput
	ToTimeAggregationTypePtrOutputWithContext(context.Context) TimeAggregationTypePtrOutput
}

type timeAggregationTypePtr string

func TimeAggregationTypePtr(v string) TimeAggregationTypePtrInput {
	return (*timeAggregationTypePtr)(&v)
}

func (*timeAggregationTypePtr) ElementType() reflect.Type {
	return timeAggregationTypePtrType
}

func (in *timeAggregationTypePtr) ToTimeAggregationTypePtrOutput() TimeAggregationTypePtrOutput {
	return pulumi.ToOutput(in).(TimeAggregationTypePtrOutput)
}

func (in *timeAggregationTypePtr) ToTimeAggregationTypePtrOutputWithContext(ctx context.Context) TimeAggregationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TimeAggregationTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ComparisonOperationTypeOutput{})
	pulumi.RegisterOutputType(ComparisonOperationTypePtrOutput{})
	pulumi.RegisterOutputType(MetricStatisticTypeOutput{})
	pulumi.RegisterOutputType(MetricStatisticTypePtrOutput{})
	pulumi.RegisterOutputType(OperationTypeOutput{})
	pulumi.RegisterOutputType(OperationTypePtrOutput{})
	pulumi.RegisterOutputType(PredictiveAutoscalePolicyScaleModeOutput{})
	pulumi.RegisterOutputType(PredictiveAutoscalePolicyScaleModePtrOutput{})
	pulumi.RegisterOutputType(RecurrenceFrequencyOutput{})
	pulumi.RegisterOutputType(RecurrenceFrequencyPtrOutput{})
	pulumi.RegisterOutputType(ScaleDirectionOutput{})
	pulumi.RegisterOutputType(ScaleDirectionPtrOutput{})
	pulumi.RegisterOutputType(ScaleTypeOutput{})
	pulumi.RegisterOutputType(ScaleTypePtrOutput{})
	pulumi.RegisterOutputType(TimeAggregationTypeOutput{})
	pulumi.RegisterOutputType(TimeAggregationTypePtrOutput{})
}
