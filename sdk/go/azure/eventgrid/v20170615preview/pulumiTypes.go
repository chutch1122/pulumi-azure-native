


package v20170615preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EventSubscriptionDestination struct {
	EndpointType *string `pulumi:"endpointType"`
	EndpointUrl  *string `pulumi:"endpointUrl"`
}





type EventSubscriptionDestinationInput interface {
	pulumi.Input

	ToEventSubscriptionDestinationOutput() EventSubscriptionDestinationOutput
	ToEventSubscriptionDestinationOutputWithContext(context.Context) EventSubscriptionDestinationOutput
}

type EventSubscriptionDestinationArgs struct {
	EndpointType pulumi.StringPtrInput `pulumi:"endpointType"`
	EndpointUrl  pulumi.StringPtrInput `pulumi:"endpointUrl"`
}

func (EventSubscriptionDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionDestination)(nil)).Elem()
}

func (i EventSubscriptionDestinationArgs) ToEventSubscriptionDestinationOutput() EventSubscriptionDestinationOutput {
	return i.ToEventSubscriptionDestinationOutputWithContext(context.Background())
}

func (i EventSubscriptionDestinationArgs) ToEventSubscriptionDestinationOutputWithContext(ctx context.Context) EventSubscriptionDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionDestinationOutput)
}

func (i EventSubscriptionDestinationArgs) ToEventSubscriptionDestinationPtrOutput() EventSubscriptionDestinationPtrOutput {
	return i.ToEventSubscriptionDestinationPtrOutputWithContext(context.Background())
}

func (i EventSubscriptionDestinationArgs) ToEventSubscriptionDestinationPtrOutputWithContext(ctx context.Context) EventSubscriptionDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionDestinationOutput).ToEventSubscriptionDestinationPtrOutputWithContext(ctx)
}









type EventSubscriptionDestinationPtrInput interface {
	pulumi.Input

	ToEventSubscriptionDestinationPtrOutput() EventSubscriptionDestinationPtrOutput
	ToEventSubscriptionDestinationPtrOutputWithContext(context.Context) EventSubscriptionDestinationPtrOutput
}

type eventSubscriptionDestinationPtrType EventSubscriptionDestinationArgs

func EventSubscriptionDestinationPtr(v *EventSubscriptionDestinationArgs) EventSubscriptionDestinationPtrInput {
	return (*eventSubscriptionDestinationPtrType)(v)
}

func (*eventSubscriptionDestinationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionDestination)(nil)).Elem()
}

func (i *eventSubscriptionDestinationPtrType) ToEventSubscriptionDestinationPtrOutput() EventSubscriptionDestinationPtrOutput {
	return i.ToEventSubscriptionDestinationPtrOutputWithContext(context.Background())
}

func (i *eventSubscriptionDestinationPtrType) ToEventSubscriptionDestinationPtrOutputWithContext(ctx context.Context) EventSubscriptionDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionDestinationPtrOutput)
}

type EventSubscriptionDestinationOutput struct{ *pulumi.OutputState }

func (EventSubscriptionDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionDestination)(nil)).Elem()
}

func (o EventSubscriptionDestinationOutput) ToEventSubscriptionDestinationOutput() EventSubscriptionDestinationOutput {
	return o
}

func (o EventSubscriptionDestinationOutput) ToEventSubscriptionDestinationOutputWithContext(ctx context.Context) EventSubscriptionDestinationOutput {
	return o
}

func (o EventSubscriptionDestinationOutput) ToEventSubscriptionDestinationPtrOutput() EventSubscriptionDestinationPtrOutput {
	return o.ToEventSubscriptionDestinationPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionDestinationOutput) ToEventSubscriptionDestinationPtrOutputWithContext(ctx context.Context) EventSubscriptionDestinationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventSubscriptionDestination) *EventSubscriptionDestination {
		return &v
	}).(EventSubscriptionDestinationPtrOutput)
}

func (o EventSubscriptionDestinationOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionDestination) *string { return v.EndpointType }).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionDestinationOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionDestination) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

type EventSubscriptionDestinationPtrOutput struct{ *pulumi.OutputState }

func (EventSubscriptionDestinationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionDestination)(nil)).Elem()
}

func (o EventSubscriptionDestinationPtrOutput) ToEventSubscriptionDestinationPtrOutput() EventSubscriptionDestinationPtrOutput {
	return o
}

func (o EventSubscriptionDestinationPtrOutput) ToEventSubscriptionDestinationPtrOutputWithContext(ctx context.Context) EventSubscriptionDestinationPtrOutput {
	return o
}

func (o EventSubscriptionDestinationPtrOutput) Elem() EventSubscriptionDestinationOutput {
	return o.ApplyT(func(v *EventSubscriptionDestination) EventSubscriptionDestination {
		if v != nil {
			return *v
		}
		var ret EventSubscriptionDestination
		return ret
	}).(EventSubscriptionDestinationOutput)
}

func (o EventSubscriptionDestinationPtrOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionDestination) *string {
		if v == nil {
			return nil
		}
		return v.EndpointType
	}).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionDestinationPtrOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionDestination) *string {
		if v == nil {
			return nil
		}
		return v.EndpointUrl
	}).(pulumi.StringPtrOutput)
}

type EventSubscriptionDestinationResponse struct {
	EndpointBaseUrl string  `pulumi:"endpointBaseUrl"`
	EndpointType    *string `pulumi:"endpointType"`
	EndpointUrl     *string `pulumi:"endpointUrl"`
}

type EventSubscriptionDestinationResponseOutput struct{ *pulumi.OutputState }

func (EventSubscriptionDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionDestinationResponse)(nil)).Elem()
}

func (o EventSubscriptionDestinationResponseOutput) ToEventSubscriptionDestinationResponseOutput() EventSubscriptionDestinationResponseOutput {
	return o
}

func (o EventSubscriptionDestinationResponseOutput) ToEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) EventSubscriptionDestinationResponseOutput {
	return o
}

func (o EventSubscriptionDestinationResponseOutput) EndpointBaseUrl() pulumi.StringOutput {
	return o.ApplyT(func(v EventSubscriptionDestinationResponse) string { return v.EndpointBaseUrl }).(pulumi.StringOutput)
}

func (o EventSubscriptionDestinationResponseOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionDestinationResponse) *string { return v.EndpointType }).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionDestinationResponseOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionDestinationResponse) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

type EventSubscriptionDestinationResponsePtrOutput struct{ *pulumi.OutputState }

func (EventSubscriptionDestinationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionDestinationResponse)(nil)).Elem()
}

func (o EventSubscriptionDestinationResponsePtrOutput) ToEventSubscriptionDestinationResponsePtrOutput() EventSubscriptionDestinationResponsePtrOutput {
	return o
}

func (o EventSubscriptionDestinationResponsePtrOutput) ToEventSubscriptionDestinationResponsePtrOutputWithContext(ctx context.Context) EventSubscriptionDestinationResponsePtrOutput {
	return o
}

func (o EventSubscriptionDestinationResponsePtrOutput) Elem() EventSubscriptionDestinationResponseOutput {
	return o.ApplyT(func(v *EventSubscriptionDestinationResponse) EventSubscriptionDestinationResponse {
		if v != nil {
			return *v
		}
		var ret EventSubscriptionDestinationResponse
		return ret
	}).(EventSubscriptionDestinationResponseOutput)
}

func (o EventSubscriptionDestinationResponsePtrOutput) EndpointBaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EndpointBaseUrl
	}).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionDestinationResponsePtrOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndpointType
	}).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionDestinationResponsePtrOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndpointUrl
	}).(pulumi.StringPtrOutput)
}

type EventSubscriptionFilter struct {
	IncludedEventTypes     []string `pulumi:"includedEventTypes"`
	IsSubjectCaseSensitive *bool    `pulumi:"isSubjectCaseSensitive"`
	SubjectBeginsWith      *string  `pulumi:"subjectBeginsWith"`
	SubjectEndsWith        *string  `pulumi:"subjectEndsWith"`
}


func (val *EventSubscriptionFilter) Defaults() *EventSubscriptionFilter {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSubjectCaseSensitive) {
		isSubjectCaseSensitive_ := false
		tmp.IsSubjectCaseSensitive = &isSubjectCaseSensitive_
	}
	return &tmp
}





type EventSubscriptionFilterInput interface {
	pulumi.Input

	ToEventSubscriptionFilterOutput() EventSubscriptionFilterOutput
	ToEventSubscriptionFilterOutputWithContext(context.Context) EventSubscriptionFilterOutput
}

type EventSubscriptionFilterArgs struct {
	IncludedEventTypes     pulumi.StringArrayInput `pulumi:"includedEventTypes"`
	IsSubjectCaseSensitive pulumi.BoolPtrInput     `pulumi:"isSubjectCaseSensitive"`
	SubjectBeginsWith      pulumi.StringPtrInput   `pulumi:"subjectBeginsWith"`
	SubjectEndsWith        pulumi.StringPtrInput   `pulumi:"subjectEndsWith"`
}


func (val *EventSubscriptionFilterArgs) Defaults() *EventSubscriptionFilterArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSubjectCaseSensitive) {
		tmp.IsSubjectCaseSensitive = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (EventSubscriptionFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionFilter)(nil)).Elem()
}

func (i EventSubscriptionFilterArgs) ToEventSubscriptionFilterOutput() EventSubscriptionFilterOutput {
	return i.ToEventSubscriptionFilterOutputWithContext(context.Background())
}

func (i EventSubscriptionFilterArgs) ToEventSubscriptionFilterOutputWithContext(ctx context.Context) EventSubscriptionFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionFilterOutput)
}

func (i EventSubscriptionFilterArgs) ToEventSubscriptionFilterPtrOutput() EventSubscriptionFilterPtrOutput {
	return i.ToEventSubscriptionFilterPtrOutputWithContext(context.Background())
}

func (i EventSubscriptionFilterArgs) ToEventSubscriptionFilterPtrOutputWithContext(ctx context.Context) EventSubscriptionFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionFilterOutput).ToEventSubscriptionFilterPtrOutputWithContext(ctx)
}









type EventSubscriptionFilterPtrInput interface {
	pulumi.Input

	ToEventSubscriptionFilterPtrOutput() EventSubscriptionFilterPtrOutput
	ToEventSubscriptionFilterPtrOutputWithContext(context.Context) EventSubscriptionFilterPtrOutput
}

type eventSubscriptionFilterPtrType EventSubscriptionFilterArgs

func EventSubscriptionFilterPtr(v *EventSubscriptionFilterArgs) EventSubscriptionFilterPtrInput {
	return (*eventSubscriptionFilterPtrType)(v)
}

func (*eventSubscriptionFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionFilter)(nil)).Elem()
}

func (i *eventSubscriptionFilterPtrType) ToEventSubscriptionFilterPtrOutput() EventSubscriptionFilterPtrOutput {
	return i.ToEventSubscriptionFilterPtrOutputWithContext(context.Background())
}

func (i *eventSubscriptionFilterPtrType) ToEventSubscriptionFilterPtrOutputWithContext(ctx context.Context) EventSubscriptionFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionFilterPtrOutput)
}

type EventSubscriptionFilterOutput struct{ *pulumi.OutputState }

func (EventSubscriptionFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionFilter)(nil)).Elem()
}

func (o EventSubscriptionFilterOutput) ToEventSubscriptionFilterOutput() EventSubscriptionFilterOutput {
	return o
}

func (o EventSubscriptionFilterOutput) ToEventSubscriptionFilterOutputWithContext(ctx context.Context) EventSubscriptionFilterOutput {
	return o
}

func (o EventSubscriptionFilterOutput) ToEventSubscriptionFilterPtrOutput() EventSubscriptionFilterPtrOutput {
	return o.ToEventSubscriptionFilterPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionFilterOutput) ToEventSubscriptionFilterPtrOutputWithContext(ctx context.Context) EventSubscriptionFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventSubscriptionFilter) *EventSubscriptionFilter {
		return &v
	}).(EventSubscriptionFilterPtrOutput)
}

func (o EventSubscriptionFilterOutput) IncludedEventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EventSubscriptionFilter) []string { return v.IncludedEventTypes }).(pulumi.StringArrayOutput)
}

func (o EventSubscriptionFilterOutput) IsSubjectCaseSensitive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilter) *bool { return v.IsSubjectCaseSensitive }).(pulumi.BoolPtrOutput)
}

func (o EventSubscriptionFilterOutput) SubjectBeginsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilter) *string { return v.SubjectBeginsWith }).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionFilterOutput) SubjectEndsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilter) *string { return v.SubjectEndsWith }).(pulumi.StringPtrOutput)
}

type EventSubscriptionFilterPtrOutput struct{ *pulumi.OutputState }

func (EventSubscriptionFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionFilter)(nil)).Elem()
}

func (o EventSubscriptionFilterPtrOutput) ToEventSubscriptionFilterPtrOutput() EventSubscriptionFilterPtrOutput {
	return o
}

func (o EventSubscriptionFilterPtrOutput) ToEventSubscriptionFilterPtrOutputWithContext(ctx context.Context) EventSubscriptionFilterPtrOutput {
	return o
}

func (o EventSubscriptionFilterPtrOutput) Elem() EventSubscriptionFilterOutput {
	return o.ApplyT(func(v *EventSubscriptionFilter) EventSubscriptionFilter {
		if v != nil {
			return *v
		}
		var ret EventSubscriptionFilter
		return ret
	}).(EventSubscriptionFilterOutput)
}

func (o EventSubscriptionFilterPtrOutput) IncludedEventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventSubscriptionFilter) []string {
		if v == nil {
			return nil
		}
		return v.IncludedEventTypes
	}).(pulumi.StringArrayOutput)
}

func (o EventSubscriptionFilterPtrOutput) IsSubjectCaseSensitive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilter) *bool {
		if v == nil {
			return nil
		}
		return v.IsSubjectCaseSensitive
	}).(pulumi.BoolPtrOutput)
}

func (o EventSubscriptionFilterPtrOutput) SubjectBeginsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilter) *string {
		if v == nil {
			return nil
		}
		return v.SubjectBeginsWith
	}).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionFilterPtrOutput) SubjectEndsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilter) *string {
		if v == nil {
			return nil
		}
		return v.SubjectEndsWith
	}).(pulumi.StringPtrOutput)
}

type EventSubscriptionFilterResponse struct {
	IncludedEventTypes     []string `pulumi:"includedEventTypes"`
	IsSubjectCaseSensitive *bool    `pulumi:"isSubjectCaseSensitive"`
	SubjectBeginsWith      *string  `pulumi:"subjectBeginsWith"`
	SubjectEndsWith        *string  `pulumi:"subjectEndsWith"`
}


func (val *EventSubscriptionFilterResponse) Defaults() *EventSubscriptionFilterResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSubjectCaseSensitive) {
		isSubjectCaseSensitive_ := false
		tmp.IsSubjectCaseSensitive = &isSubjectCaseSensitive_
	}
	return &tmp
}

type EventSubscriptionFilterResponseOutput struct{ *pulumi.OutputState }

func (EventSubscriptionFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionFilterResponse)(nil)).Elem()
}

func (o EventSubscriptionFilterResponseOutput) ToEventSubscriptionFilterResponseOutput() EventSubscriptionFilterResponseOutput {
	return o
}

func (o EventSubscriptionFilterResponseOutput) ToEventSubscriptionFilterResponseOutputWithContext(ctx context.Context) EventSubscriptionFilterResponseOutput {
	return o
}

func (o EventSubscriptionFilterResponseOutput) IncludedEventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) []string { return v.IncludedEventTypes }).(pulumi.StringArrayOutput)
}

func (o EventSubscriptionFilterResponseOutput) IsSubjectCaseSensitive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) *bool { return v.IsSubjectCaseSensitive }).(pulumi.BoolPtrOutput)
}

func (o EventSubscriptionFilterResponseOutput) SubjectBeginsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) *string { return v.SubjectBeginsWith }).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionFilterResponseOutput) SubjectEndsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) *string { return v.SubjectEndsWith }).(pulumi.StringPtrOutput)
}

type EventSubscriptionFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (EventSubscriptionFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionFilterResponse)(nil)).Elem()
}

func (o EventSubscriptionFilterResponsePtrOutput) ToEventSubscriptionFilterResponsePtrOutput() EventSubscriptionFilterResponsePtrOutput {
	return o
}

func (o EventSubscriptionFilterResponsePtrOutput) ToEventSubscriptionFilterResponsePtrOutputWithContext(ctx context.Context) EventSubscriptionFilterResponsePtrOutput {
	return o
}

func (o EventSubscriptionFilterResponsePtrOutput) Elem() EventSubscriptionFilterResponseOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) EventSubscriptionFilterResponse {
		if v != nil {
			return *v
		}
		var ret EventSubscriptionFilterResponse
		return ret
	}).(EventSubscriptionFilterResponseOutput)
}

func (o EventSubscriptionFilterResponsePtrOutput) IncludedEventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) []string {
		if v == nil {
			return nil
		}
		return v.IncludedEventTypes
	}).(pulumi.StringArrayOutput)
}

func (o EventSubscriptionFilterResponsePtrOutput) IsSubjectCaseSensitive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsSubjectCaseSensitive
	}).(pulumi.BoolPtrOutput)
}

func (o EventSubscriptionFilterResponsePtrOutput) SubjectBeginsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubjectBeginsWith
	}).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionFilterResponsePtrOutput) SubjectEndsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubjectEndsWith
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EventSubscriptionDestinationOutput{})
	pulumi.RegisterOutputType(EventSubscriptionDestinationPtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionDestinationResponseOutput{})
	pulumi.RegisterOutputType(EventSubscriptionDestinationResponsePtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterPtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterResponseOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterResponsePtrOutput{})
}
