


package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Action struct {
	CompatibilityLevel    *int    `pulumi:"compatibilityLevel"`
	RequiresPreprocessing *bool   `pulumi:"requiresPreprocessing"`
	SqlExpression         *string `pulumi:"sqlExpression"`
}


func (val *Action) Defaults() *Action {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.RequiresPreprocessing) {
		requiresPreprocessing_ := true
		tmp.RequiresPreprocessing = &requiresPreprocessing_
	}
	return &tmp
}





type ActionInput interface {
	pulumi.Input

	ToActionOutput() ActionOutput
	ToActionOutputWithContext(context.Context) ActionOutput
}

type ActionArgs struct {
	CompatibilityLevel    pulumi.IntPtrInput    `pulumi:"compatibilityLevel"`
	RequiresPreprocessing pulumi.BoolPtrInput   `pulumi:"requiresPreprocessing"`
	SqlExpression         pulumi.StringPtrInput `pulumi:"sqlExpression"`
}


func (val *ActionArgs) Defaults() *ActionArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.RequiresPreprocessing) {
		tmp.RequiresPreprocessing = pulumi.BoolPtr(true)
	}
	return &tmp
}
func (ActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (i ActionArgs) ToActionOutput() ActionOutput {
	return i.ToActionOutputWithContext(context.Background())
}

func (i ActionArgs) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionOutput)
}

func (i ActionArgs) ToActionPtrOutput() ActionPtrOutput {
	return i.ToActionPtrOutputWithContext(context.Background())
}

func (i ActionArgs) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionOutput).ToActionPtrOutputWithContext(ctx)
}









type ActionPtrInput interface {
	pulumi.Input

	ToActionPtrOutput() ActionPtrOutput
	ToActionPtrOutputWithContext(context.Context) ActionPtrOutput
}

type actionPtrType ActionArgs

func ActionPtr(v *ActionArgs) ActionPtrInput {
	return (*actionPtrType)(v)
}

func (*actionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Action)(nil)).Elem()
}

func (i *actionPtrType) ToActionPtrOutput() ActionPtrOutput {
	return i.ToActionPtrOutputWithContext(context.Background())
}

func (i *actionPtrType) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionPtrOutput)
}

type ActionOutput struct{ *pulumi.OutputState }

func (ActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Action)(nil)).Elem()
}

func (o ActionOutput) ToActionOutput() ActionOutput {
	return o
}

func (o ActionOutput) ToActionOutputWithContext(ctx context.Context) ActionOutput {
	return o
}

func (o ActionOutput) ToActionPtrOutput() ActionPtrOutput {
	return o.ToActionPtrOutputWithContext(context.Background())
}

func (o ActionOutput) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Action) *Action {
		return &v
	}).(ActionPtrOutput)
}

func (o ActionOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Action) *int { return v.CompatibilityLevel }).(pulumi.IntPtrOutput)
}

func (o ActionOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Action) *bool { return v.RequiresPreprocessing }).(pulumi.BoolPtrOutput)
}

func (o ActionOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Action) *string { return v.SqlExpression }).(pulumi.StringPtrOutput)
}

type ActionPtrOutput struct{ *pulumi.OutputState }

func (ActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Action)(nil)).Elem()
}

func (o ActionPtrOutput) ToActionPtrOutput() ActionPtrOutput {
	return o
}

func (o ActionPtrOutput) ToActionPtrOutputWithContext(ctx context.Context) ActionPtrOutput {
	return o
}

func (o ActionPtrOutput) Elem() ActionOutput {
	return o.ApplyT(func(v *Action) Action {
		if v != nil {
			return *v
		}
		var ret Action
		return ret
	}).(ActionOutput)
}

func (o ActionPtrOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Action) *int {
		if v == nil {
			return nil
		}
		return v.CompatibilityLevel
	}).(pulumi.IntPtrOutput)
}

func (o ActionPtrOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Action) *bool {
		if v == nil {
			return nil
		}
		return v.RequiresPreprocessing
	}).(pulumi.BoolPtrOutput)
}

func (o ActionPtrOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Action) *string {
		if v == nil {
			return nil
		}
		return v.SqlExpression
	}).(pulumi.StringPtrOutput)
}

type ActionResponse struct {
	CompatibilityLevel    *int    `pulumi:"compatibilityLevel"`
	RequiresPreprocessing *bool   `pulumi:"requiresPreprocessing"`
	SqlExpression         *string `pulumi:"sqlExpression"`
}


func (val *ActionResponse) Defaults() *ActionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.RequiresPreprocessing) {
		requiresPreprocessing_ := true
		tmp.RequiresPreprocessing = &requiresPreprocessing_
	}
	return &tmp
}

type ActionResponseOutput struct{ *pulumi.OutputState }

func (ActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionResponse)(nil)).Elem()
}

func (o ActionResponseOutput) ToActionResponseOutput() ActionResponseOutput {
	return o
}

func (o ActionResponseOutput) ToActionResponseOutputWithContext(ctx context.Context) ActionResponseOutput {
	return o
}

func (o ActionResponseOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ActionResponse) *int { return v.CompatibilityLevel }).(pulumi.IntPtrOutput)
}

func (o ActionResponseOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ActionResponse) *bool { return v.RequiresPreprocessing }).(pulumi.BoolPtrOutput)
}

func (o ActionResponseOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionResponse) *string { return v.SqlExpression }).(pulumi.StringPtrOutput)
}

type ActionResponsePtrOutput struct{ *pulumi.OutputState }

func (ActionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionResponse)(nil)).Elem()
}

func (o ActionResponsePtrOutput) ToActionResponsePtrOutput() ActionResponsePtrOutput {
	return o
}

func (o ActionResponsePtrOutput) ToActionResponsePtrOutputWithContext(ctx context.Context) ActionResponsePtrOutput {
	return o
}

func (o ActionResponsePtrOutput) Elem() ActionResponseOutput {
	return o.ApplyT(func(v *ActionResponse) ActionResponse {
		if v != nil {
			return *v
		}
		var ret ActionResponse
		return ret
	}).(ActionResponseOutput)
}

func (o ActionResponsePtrOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ActionResponse) *int {
		if v == nil {
			return nil
		}
		return v.CompatibilityLevel
	}).(pulumi.IntPtrOutput)
}

func (o ActionResponsePtrOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ActionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequiresPreprocessing
	}).(pulumi.BoolPtrOutput)
}

func (o ActionResponsePtrOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionResponse) *string {
		if v == nil {
			return nil
		}
		return v.SqlExpression
	}).(pulumi.StringPtrOutput)
}

type CorrelationFilter struct {
	ContentType           *string           `pulumi:"contentType"`
	CorrelationId         *string           `pulumi:"correlationId"`
	Label                 *string           `pulumi:"label"`
	MessageId             *string           `pulumi:"messageId"`
	Properties            map[string]string `pulumi:"properties"`
	ReplyTo               *string           `pulumi:"replyTo"`
	ReplyToSessionId      *string           `pulumi:"replyToSessionId"`
	RequiresPreprocessing *bool             `pulumi:"requiresPreprocessing"`
	SessionId             *string           `pulumi:"sessionId"`
	To                    *string           `pulumi:"to"`
}


func (val *CorrelationFilter) Defaults() *CorrelationFilter {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.RequiresPreprocessing) {
		requiresPreprocessing_ := true
		tmp.RequiresPreprocessing = &requiresPreprocessing_
	}
	return &tmp
}





type CorrelationFilterInput interface {
	pulumi.Input

	ToCorrelationFilterOutput() CorrelationFilterOutput
	ToCorrelationFilterOutputWithContext(context.Context) CorrelationFilterOutput
}

type CorrelationFilterArgs struct {
	ContentType           pulumi.StringPtrInput `pulumi:"contentType"`
	CorrelationId         pulumi.StringPtrInput `pulumi:"correlationId"`
	Label                 pulumi.StringPtrInput `pulumi:"label"`
	MessageId             pulumi.StringPtrInput `pulumi:"messageId"`
	Properties            pulumi.StringMapInput `pulumi:"properties"`
	ReplyTo               pulumi.StringPtrInput `pulumi:"replyTo"`
	ReplyToSessionId      pulumi.StringPtrInput `pulumi:"replyToSessionId"`
	RequiresPreprocessing pulumi.BoolPtrInput   `pulumi:"requiresPreprocessing"`
	SessionId             pulumi.StringPtrInput `pulumi:"sessionId"`
	To                    pulumi.StringPtrInput `pulumi:"to"`
}


func (val *CorrelationFilterArgs) Defaults() *CorrelationFilterArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.RequiresPreprocessing) {
		tmp.RequiresPreprocessing = pulumi.BoolPtr(true)
	}
	return &tmp
}
func (CorrelationFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorrelationFilter)(nil)).Elem()
}

func (i CorrelationFilterArgs) ToCorrelationFilterOutput() CorrelationFilterOutput {
	return i.ToCorrelationFilterOutputWithContext(context.Background())
}

func (i CorrelationFilterArgs) ToCorrelationFilterOutputWithContext(ctx context.Context) CorrelationFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorrelationFilterOutput)
}

func (i CorrelationFilterArgs) ToCorrelationFilterPtrOutput() CorrelationFilterPtrOutput {
	return i.ToCorrelationFilterPtrOutputWithContext(context.Background())
}

func (i CorrelationFilterArgs) ToCorrelationFilterPtrOutputWithContext(ctx context.Context) CorrelationFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorrelationFilterOutput).ToCorrelationFilterPtrOutputWithContext(ctx)
}









type CorrelationFilterPtrInput interface {
	pulumi.Input

	ToCorrelationFilterPtrOutput() CorrelationFilterPtrOutput
	ToCorrelationFilterPtrOutputWithContext(context.Context) CorrelationFilterPtrOutput
}

type correlationFilterPtrType CorrelationFilterArgs

func CorrelationFilterPtr(v *CorrelationFilterArgs) CorrelationFilterPtrInput {
	return (*correlationFilterPtrType)(v)
}

func (*correlationFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CorrelationFilter)(nil)).Elem()
}

func (i *correlationFilterPtrType) ToCorrelationFilterPtrOutput() CorrelationFilterPtrOutput {
	return i.ToCorrelationFilterPtrOutputWithContext(context.Background())
}

func (i *correlationFilterPtrType) ToCorrelationFilterPtrOutputWithContext(ctx context.Context) CorrelationFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorrelationFilterPtrOutput)
}

type CorrelationFilterOutput struct{ *pulumi.OutputState }

func (CorrelationFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorrelationFilter)(nil)).Elem()
}

func (o CorrelationFilterOutput) ToCorrelationFilterOutput() CorrelationFilterOutput {
	return o
}

func (o CorrelationFilterOutput) ToCorrelationFilterOutputWithContext(ctx context.Context) CorrelationFilterOutput {
	return o
}

func (o CorrelationFilterOutput) ToCorrelationFilterPtrOutput() CorrelationFilterPtrOutput {
	return o.ToCorrelationFilterPtrOutputWithContext(context.Background())
}

func (o CorrelationFilterOutput) ToCorrelationFilterPtrOutputWithContext(ctx context.Context) CorrelationFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CorrelationFilter) *CorrelationFilter {
		return &v
	}).(CorrelationFilterPtrOutput)
}

func (o CorrelationFilterOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilter) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilter) *string { return v.CorrelationId }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilter) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilter) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v CorrelationFilter) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o CorrelationFilterOutput) ReplyTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilter) *string { return v.ReplyTo }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterOutput) ReplyToSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilter) *string { return v.ReplyToSessionId }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CorrelationFilter) *bool { return v.RequiresPreprocessing }).(pulumi.BoolPtrOutput)
}

func (o CorrelationFilterOutput) SessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilter) *string { return v.SessionId }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilter) *string { return v.To }).(pulumi.StringPtrOutput)
}

type CorrelationFilterPtrOutput struct{ *pulumi.OutputState }

func (CorrelationFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CorrelationFilter)(nil)).Elem()
}

func (o CorrelationFilterPtrOutput) ToCorrelationFilterPtrOutput() CorrelationFilterPtrOutput {
	return o
}

func (o CorrelationFilterPtrOutput) ToCorrelationFilterPtrOutputWithContext(ctx context.Context) CorrelationFilterPtrOutput {
	return o
}

func (o CorrelationFilterPtrOutput) Elem() CorrelationFilterOutput {
	return o.ApplyT(func(v *CorrelationFilter) CorrelationFilter {
		if v != nil {
			return *v
		}
		var ret CorrelationFilter
		return ret
	}).(CorrelationFilterOutput)
}

func (o CorrelationFilterPtrOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilter) *string {
		if v == nil {
			return nil
		}
		return v.ContentType
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterPtrOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilter) *string {
		if v == nil {
			return nil
		}
		return v.CorrelationId
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterPtrOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilter) *string {
		if v == nil {
			return nil
		}
		return v.Label
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterPtrOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilter) *string {
		if v == nil {
			return nil
		}
		return v.MessageId
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterPtrOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CorrelationFilter) map[string]string {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.StringMapOutput)
}

func (o CorrelationFilterPtrOutput) ReplyTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilter) *string {
		if v == nil {
			return nil
		}
		return v.ReplyTo
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterPtrOutput) ReplyToSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilter) *string {
		if v == nil {
			return nil
		}
		return v.ReplyToSessionId
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterPtrOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CorrelationFilter) *bool {
		if v == nil {
			return nil
		}
		return v.RequiresPreprocessing
	}).(pulumi.BoolPtrOutput)
}

func (o CorrelationFilterPtrOutput) SessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilter) *string {
		if v == nil {
			return nil
		}
		return v.SessionId
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterPtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilter) *string {
		if v == nil {
			return nil
		}
		return v.To
	}).(pulumi.StringPtrOutput)
}

type CorrelationFilterResponse struct {
	ContentType           *string           `pulumi:"contentType"`
	CorrelationId         *string           `pulumi:"correlationId"`
	Label                 *string           `pulumi:"label"`
	MessageId             *string           `pulumi:"messageId"`
	Properties            map[string]string `pulumi:"properties"`
	ReplyTo               *string           `pulumi:"replyTo"`
	ReplyToSessionId      *string           `pulumi:"replyToSessionId"`
	RequiresPreprocessing *bool             `pulumi:"requiresPreprocessing"`
	SessionId             *string           `pulumi:"sessionId"`
	To                    *string           `pulumi:"to"`
}


func (val *CorrelationFilterResponse) Defaults() *CorrelationFilterResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.RequiresPreprocessing) {
		requiresPreprocessing_ := true
		tmp.RequiresPreprocessing = &requiresPreprocessing_
	}
	return &tmp
}

type CorrelationFilterResponseOutput struct{ *pulumi.OutputState }

func (CorrelationFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorrelationFilterResponse)(nil)).Elem()
}

func (o CorrelationFilterResponseOutput) ToCorrelationFilterResponseOutput() CorrelationFilterResponseOutput {
	return o
}

func (o CorrelationFilterResponseOutput) ToCorrelationFilterResponseOutputWithContext(ctx context.Context) CorrelationFilterResponseOutput {
	return o
}

func (o CorrelationFilterResponseOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponseOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.CorrelationId }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponseOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.Label }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponseOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o CorrelationFilterResponseOutput) ReplyTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.ReplyTo }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponseOutput) ReplyToSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.ReplyToSessionId }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponseOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *bool { return v.RequiresPreprocessing }).(pulumi.BoolPtrOutput)
}

func (o CorrelationFilterResponseOutput) SessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.SessionId }).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponseOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorrelationFilterResponse) *string { return v.To }).(pulumi.StringPtrOutput)
}

type CorrelationFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (CorrelationFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CorrelationFilterResponse)(nil)).Elem()
}

func (o CorrelationFilterResponsePtrOutput) ToCorrelationFilterResponsePtrOutput() CorrelationFilterResponsePtrOutput {
	return o
}

func (o CorrelationFilterResponsePtrOutput) ToCorrelationFilterResponsePtrOutputWithContext(ctx context.Context) CorrelationFilterResponsePtrOutput {
	return o
}

func (o CorrelationFilterResponsePtrOutput) Elem() CorrelationFilterResponseOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) CorrelationFilterResponse {
		if v != nil {
			return *v
		}
		var ret CorrelationFilterResponse
		return ret
	}).(CorrelationFilterResponseOutput)
}

func (o CorrelationFilterResponsePtrOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContentType
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponsePtrOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.CorrelationId
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponsePtrOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.Label
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponsePtrOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.MessageId
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponsePtrOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(pulumi.StringMapOutput)
}

func (o CorrelationFilterResponsePtrOutput) ReplyTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReplyTo
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponsePtrOutput) ReplyToSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReplyToSessionId
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponsePtrOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequiresPreprocessing
	}).(pulumi.BoolPtrOutput)
}

func (o CorrelationFilterResponsePtrOutput) SessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.SessionId
	}).(pulumi.StringPtrOutput)
}

func (o CorrelationFilterResponsePtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CorrelationFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.To
	}).(pulumi.StringPtrOutput)
}

type MessageCountDetailsResponse struct {
	ActiveMessageCount             float64 `pulumi:"activeMessageCount"`
	DeadLetterMessageCount         float64 `pulumi:"deadLetterMessageCount"`
	ScheduledMessageCount          float64 `pulumi:"scheduledMessageCount"`
	TransferDeadLetterMessageCount float64 `pulumi:"transferDeadLetterMessageCount"`
	TransferMessageCount           float64 `pulumi:"transferMessageCount"`
}

type MessageCountDetailsResponseOutput struct{ *pulumi.OutputState }

func (MessageCountDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MessageCountDetailsResponse)(nil)).Elem()
}

func (o MessageCountDetailsResponseOutput) ToMessageCountDetailsResponseOutput() MessageCountDetailsResponseOutput {
	return o
}

func (o MessageCountDetailsResponseOutput) ToMessageCountDetailsResponseOutputWithContext(ctx context.Context) MessageCountDetailsResponseOutput {
	return o
}

func (o MessageCountDetailsResponseOutput) ActiveMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.ActiveMessageCount }).(pulumi.Float64Output)
}

func (o MessageCountDetailsResponseOutput) DeadLetterMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.DeadLetterMessageCount }).(pulumi.Float64Output)
}

func (o MessageCountDetailsResponseOutput) ScheduledMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.ScheduledMessageCount }).(pulumi.Float64Output)
}

func (o MessageCountDetailsResponseOutput) TransferDeadLetterMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.TransferDeadLetterMessageCount }).(pulumi.Float64Output)
}

func (o MessageCountDetailsResponseOutput) TransferMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.TransferMessageCount }).(pulumi.Float64Output)
}

type NWRuleSetIpRules struct {
	Action *string `pulumi:"action"`
	IpMask *string `pulumi:"ipMask"`
}


func (val *NWRuleSetIpRules) Defaults() *NWRuleSetIpRules {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}





type NWRuleSetIpRulesInput interface {
	pulumi.Input

	ToNWRuleSetIpRulesOutput() NWRuleSetIpRulesOutput
	ToNWRuleSetIpRulesOutputWithContext(context.Context) NWRuleSetIpRulesOutput
}

type NWRuleSetIpRulesArgs struct {
	Action pulumi.StringPtrInput `pulumi:"action"`
	IpMask pulumi.StringPtrInput `pulumi:"ipMask"`
}


func (val *NWRuleSetIpRulesArgs) Defaults() *NWRuleSetIpRulesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		tmp.Action = pulumi.StringPtr("Allow")
	}
	return &tmp
}
func (NWRuleSetIpRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetIpRules)(nil)).Elem()
}

func (i NWRuleSetIpRulesArgs) ToNWRuleSetIpRulesOutput() NWRuleSetIpRulesOutput {
	return i.ToNWRuleSetIpRulesOutputWithContext(context.Background())
}

func (i NWRuleSetIpRulesArgs) ToNWRuleSetIpRulesOutputWithContext(ctx context.Context) NWRuleSetIpRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NWRuleSetIpRulesOutput)
}





type NWRuleSetIpRulesArrayInput interface {
	pulumi.Input

	ToNWRuleSetIpRulesArrayOutput() NWRuleSetIpRulesArrayOutput
	ToNWRuleSetIpRulesArrayOutputWithContext(context.Context) NWRuleSetIpRulesArrayOutput
}

type NWRuleSetIpRulesArray []NWRuleSetIpRulesInput

func (NWRuleSetIpRulesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetIpRules)(nil)).Elem()
}

func (i NWRuleSetIpRulesArray) ToNWRuleSetIpRulesArrayOutput() NWRuleSetIpRulesArrayOutput {
	return i.ToNWRuleSetIpRulesArrayOutputWithContext(context.Background())
}

func (i NWRuleSetIpRulesArray) ToNWRuleSetIpRulesArrayOutputWithContext(ctx context.Context) NWRuleSetIpRulesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NWRuleSetIpRulesArrayOutput)
}

type NWRuleSetIpRulesOutput struct{ *pulumi.OutputState }

func (NWRuleSetIpRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetIpRules)(nil)).Elem()
}

func (o NWRuleSetIpRulesOutput) ToNWRuleSetIpRulesOutput() NWRuleSetIpRulesOutput {
	return o
}

func (o NWRuleSetIpRulesOutput) ToNWRuleSetIpRulesOutputWithContext(ctx context.Context) NWRuleSetIpRulesOutput {
	return o
}

func (o NWRuleSetIpRulesOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NWRuleSetIpRules) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o NWRuleSetIpRulesOutput) IpMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NWRuleSetIpRules) *string { return v.IpMask }).(pulumi.StringPtrOutput)
}

type NWRuleSetIpRulesArrayOutput struct{ *pulumi.OutputState }

func (NWRuleSetIpRulesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetIpRules)(nil)).Elem()
}

func (o NWRuleSetIpRulesArrayOutput) ToNWRuleSetIpRulesArrayOutput() NWRuleSetIpRulesArrayOutput {
	return o
}

func (o NWRuleSetIpRulesArrayOutput) ToNWRuleSetIpRulesArrayOutputWithContext(ctx context.Context) NWRuleSetIpRulesArrayOutput {
	return o
}

func (o NWRuleSetIpRulesArrayOutput) Index(i pulumi.IntInput) NWRuleSetIpRulesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NWRuleSetIpRules {
		return vs[0].([]NWRuleSetIpRules)[vs[1].(int)]
	}).(NWRuleSetIpRulesOutput)
}

type NWRuleSetIpRulesResponse struct {
	Action *string `pulumi:"action"`
	IpMask *string `pulumi:"ipMask"`
}


func (val *NWRuleSetIpRulesResponse) Defaults() *NWRuleSetIpRulesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}

type NWRuleSetIpRulesResponseOutput struct{ *pulumi.OutputState }

func (NWRuleSetIpRulesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetIpRulesResponse)(nil)).Elem()
}

func (o NWRuleSetIpRulesResponseOutput) ToNWRuleSetIpRulesResponseOutput() NWRuleSetIpRulesResponseOutput {
	return o
}

func (o NWRuleSetIpRulesResponseOutput) ToNWRuleSetIpRulesResponseOutputWithContext(ctx context.Context) NWRuleSetIpRulesResponseOutput {
	return o
}

func (o NWRuleSetIpRulesResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NWRuleSetIpRulesResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o NWRuleSetIpRulesResponseOutput) IpMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NWRuleSetIpRulesResponse) *string { return v.IpMask }).(pulumi.StringPtrOutput)
}

type NWRuleSetIpRulesResponseArrayOutput struct{ *pulumi.OutputState }

func (NWRuleSetIpRulesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetIpRulesResponse)(nil)).Elem()
}

func (o NWRuleSetIpRulesResponseArrayOutput) ToNWRuleSetIpRulesResponseArrayOutput() NWRuleSetIpRulesResponseArrayOutput {
	return o
}

func (o NWRuleSetIpRulesResponseArrayOutput) ToNWRuleSetIpRulesResponseArrayOutputWithContext(ctx context.Context) NWRuleSetIpRulesResponseArrayOutput {
	return o
}

func (o NWRuleSetIpRulesResponseArrayOutput) Index(i pulumi.IntInput) NWRuleSetIpRulesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NWRuleSetIpRulesResponse {
		return vs[0].([]NWRuleSetIpRulesResponse)[vs[1].(int)]
	}).(NWRuleSetIpRulesResponseOutput)
}

type NWRuleSetVirtualNetworkRules struct {
	IgnoreMissingVnetServiceEndpoint *bool   `pulumi:"ignoreMissingVnetServiceEndpoint"`
	Subnet                           *Subnet `pulumi:"subnet"`
}





type NWRuleSetVirtualNetworkRulesInput interface {
	pulumi.Input

	ToNWRuleSetVirtualNetworkRulesOutput() NWRuleSetVirtualNetworkRulesOutput
	ToNWRuleSetVirtualNetworkRulesOutputWithContext(context.Context) NWRuleSetVirtualNetworkRulesOutput
}

type NWRuleSetVirtualNetworkRulesArgs struct {
	IgnoreMissingVnetServiceEndpoint pulumi.BoolPtrInput `pulumi:"ignoreMissingVnetServiceEndpoint"`
	Subnet                           SubnetPtrInput      `pulumi:"subnet"`
}

func (NWRuleSetVirtualNetworkRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetVirtualNetworkRules)(nil)).Elem()
}

func (i NWRuleSetVirtualNetworkRulesArgs) ToNWRuleSetVirtualNetworkRulesOutput() NWRuleSetVirtualNetworkRulesOutput {
	return i.ToNWRuleSetVirtualNetworkRulesOutputWithContext(context.Background())
}

func (i NWRuleSetVirtualNetworkRulesArgs) ToNWRuleSetVirtualNetworkRulesOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NWRuleSetVirtualNetworkRulesOutput)
}





type NWRuleSetVirtualNetworkRulesArrayInput interface {
	pulumi.Input

	ToNWRuleSetVirtualNetworkRulesArrayOutput() NWRuleSetVirtualNetworkRulesArrayOutput
	ToNWRuleSetVirtualNetworkRulesArrayOutputWithContext(context.Context) NWRuleSetVirtualNetworkRulesArrayOutput
}

type NWRuleSetVirtualNetworkRulesArray []NWRuleSetVirtualNetworkRulesInput

func (NWRuleSetVirtualNetworkRulesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetVirtualNetworkRules)(nil)).Elem()
}

func (i NWRuleSetVirtualNetworkRulesArray) ToNWRuleSetVirtualNetworkRulesArrayOutput() NWRuleSetVirtualNetworkRulesArrayOutput {
	return i.ToNWRuleSetVirtualNetworkRulesArrayOutputWithContext(context.Background())
}

func (i NWRuleSetVirtualNetworkRulesArray) ToNWRuleSetVirtualNetworkRulesArrayOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NWRuleSetVirtualNetworkRulesArrayOutput)
}

type NWRuleSetVirtualNetworkRulesOutput struct{ *pulumi.OutputState }

func (NWRuleSetVirtualNetworkRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetVirtualNetworkRules)(nil)).Elem()
}

func (o NWRuleSetVirtualNetworkRulesOutput) ToNWRuleSetVirtualNetworkRulesOutput() NWRuleSetVirtualNetworkRulesOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesOutput) ToNWRuleSetVirtualNetworkRulesOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NWRuleSetVirtualNetworkRules) *bool { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

func (o NWRuleSetVirtualNetworkRulesOutput) Subnet() SubnetPtrOutput {
	return o.ApplyT(func(v NWRuleSetVirtualNetworkRules) *Subnet { return v.Subnet }).(SubnetPtrOutput)
}

type NWRuleSetVirtualNetworkRulesArrayOutput struct{ *pulumi.OutputState }

func (NWRuleSetVirtualNetworkRulesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetVirtualNetworkRules)(nil)).Elem()
}

func (o NWRuleSetVirtualNetworkRulesArrayOutput) ToNWRuleSetVirtualNetworkRulesArrayOutput() NWRuleSetVirtualNetworkRulesArrayOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesArrayOutput) ToNWRuleSetVirtualNetworkRulesArrayOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesArrayOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesArrayOutput) Index(i pulumi.IntInput) NWRuleSetVirtualNetworkRulesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NWRuleSetVirtualNetworkRules {
		return vs[0].([]NWRuleSetVirtualNetworkRules)[vs[1].(int)]
	}).(NWRuleSetVirtualNetworkRulesOutput)
}

type NWRuleSetVirtualNetworkRulesResponse struct {
	IgnoreMissingVnetServiceEndpoint *bool           `pulumi:"ignoreMissingVnetServiceEndpoint"`
	Subnet                           *SubnetResponse `pulumi:"subnet"`
}

type NWRuleSetVirtualNetworkRulesResponseOutput struct{ *pulumi.OutputState }

func (NWRuleSetVirtualNetworkRulesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NWRuleSetVirtualNetworkRulesResponse)(nil)).Elem()
}

func (o NWRuleSetVirtualNetworkRulesResponseOutput) ToNWRuleSetVirtualNetworkRulesResponseOutput() NWRuleSetVirtualNetworkRulesResponseOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesResponseOutput) ToNWRuleSetVirtualNetworkRulesResponseOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesResponseOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesResponseOutput) IgnoreMissingVnetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NWRuleSetVirtualNetworkRulesResponse) *bool { return v.IgnoreMissingVnetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

func (o NWRuleSetVirtualNetworkRulesResponseOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v NWRuleSetVirtualNetworkRulesResponse) *SubnetResponse { return v.Subnet }).(SubnetResponsePtrOutput)
}

type NWRuleSetVirtualNetworkRulesResponseArrayOutput struct{ *pulumi.OutputState }

func (NWRuleSetVirtualNetworkRulesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NWRuleSetVirtualNetworkRulesResponse)(nil)).Elem()
}

func (o NWRuleSetVirtualNetworkRulesResponseArrayOutput) ToNWRuleSetVirtualNetworkRulesResponseArrayOutput() NWRuleSetVirtualNetworkRulesResponseArrayOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesResponseArrayOutput) ToNWRuleSetVirtualNetworkRulesResponseArrayOutputWithContext(ctx context.Context) NWRuleSetVirtualNetworkRulesResponseArrayOutput {
	return o
}

func (o NWRuleSetVirtualNetworkRulesResponseArrayOutput) Index(i pulumi.IntInput) NWRuleSetVirtualNetworkRulesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NWRuleSetVirtualNetworkRulesResponse {
		return vs[0].([]NWRuleSetVirtualNetworkRulesResponse)[vs[1].(int)]
	}).(NWRuleSetVirtualNetworkRulesResponseOutput)
}

type SBSku struct {
	Capacity *int     `pulumi:"capacity"`
	Name     SkuName  `pulumi:"name"`
	Tier     *SkuTier `pulumi:"tier"`
}





type SBSkuInput interface {
	pulumi.Input

	ToSBSkuOutput() SBSkuOutput
	ToSBSkuOutputWithContext(context.Context) SBSkuOutput
}

type SBSkuArgs struct {
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	Name     SkuNameInput       `pulumi:"name"`
	Tier     SkuTierPtrInput    `pulumi:"tier"`
}

func (SBSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SBSku)(nil)).Elem()
}

func (i SBSkuArgs) ToSBSkuOutput() SBSkuOutput {
	return i.ToSBSkuOutputWithContext(context.Background())
}

func (i SBSkuArgs) ToSBSkuOutputWithContext(ctx context.Context) SBSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SBSkuOutput)
}

func (i SBSkuArgs) ToSBSkuPtrOutput() SBSkuPtrOutput {
	return i.ToSBSkuPtrOutputWithContext(context.Background())
}

func (i SBSkuArgs) ToSBSkuPtrOutputWithContext(ctx context.Context) SBSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SBSkuOutput).ToSBSkuPtrOutputWithContext(ctx)
}









type SBSkuPtrInput interface {
	pulumi.Input

	ToSBSkuPtrOutput() SBSkuPtrOutput
	ToSBSkuPtrOutputWithContext(context.Context) SBSkuPtrOutput
}

type sbskuPtrType SBSkuArgs

func SBSkuPtr(v *SBSkuArgs) SBSkuPtrInput {
	return (*sbskuPtrType)(v)
}

func (*sbskuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SBSku)(nil)).Elem()
}

func (i *sbskuPtrType) ToSBSkuPtrOutput() SBSkuPtrOutput {
	return i.ToSBSkuPtrOutputWithContext(context.Background())
}

func (i *sbskuPtrType) ToSBSkuPtrOutputWithContext(ctx context.Context) SBSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SBSkuPtrOutput)
}

type SBSkuOutput struct{ *pulumi.OutputState }

func (SBSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SBSku)(nil)).Elem()
}

func (o SBSkuOutput) ToSBSkuOutput() SBSkuOutput {
	return o
}

func (o SBSkuOutput) ToSBSkuOutputWithContext(ctx context.Context) SBSkuOutput {
	return o
}

func (o SBSkuOutput) ToSBSkuPtrOutput() SBSkuPtrOutput {
	return o.ToSBSkuPtrOutputWithContext(context.Background())
}

func (o SBSkuOutput) ToSBSkuPtrOutputWithContext(ctx context.Context) SBSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SBSku) *SBSku {
		return &v
	}).(SBSkuPtrOutput)
}

func (o SBSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SBSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SBSkuOutput) Name() SkuNameOutput {
	return o.ApplyT(func(v SBSku) SkuName { return v.Name }).(SkuNameOutput)
}

func (o SBSkuOutput) Tier() SkuTierPtrOutput {
	return o.ApplyT(func(v SBSku) *SkuTier { return v.Tier }).(SkuTierPtrOutput)
}

type SBSkuPtrOutput struct{ *pulumi.OutputState }

func (SBSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SBSku)(nil)).Elem()
}

func (o SBSkuPtrOutput) ToSBSkuPtrOutput() SBSkuPtrOutput {
	return o
}

func (o SBSkuPtrOutput) ToSBSkuPtrOutputWithContext(ctx context.Context) SBSkuPtrOutput {
	return o
}

func (o SBSkuPtrOutput) Elem() SBSkuOutput {
	return o.ApplyT(func(v *SBSku) SBSku {
		if v != nil {
			return *v
		}
		var ret SBSku
		return ret
	}).(SBSkuOutput)
}

func (o SBSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SBSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SBSkuPtrOutput) Name() SkuNamePtrOutput {
	return o.ApplyT(func(v *SBSku) *SkuName {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(SkuNamePtrOutput)
}

func (o SBSkuPtrOutput) Tier() SkuTierPtrOutput {
	return o.ApplyT(func(v *SBSku) *SkuTier {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(SkuTierPtrOutput)
}

type SBSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}

type SBSkuResponseOutput struct{ *pulumi.OutputState }

func (SBSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SBSkuResponse)(nil)).Elem()
}

func (o SBSkuResponseOutput) ToSBSkuResponseOutput() SBSkuResponseOutput {
	return o
}

func (o SBSkuResponseOutput) ToSBSkuResponseOutputWithContext(ctx context.Context) SBSkuResponseOutput {
	return o
}

func (o SBSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SBSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SBSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SBSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SBSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SBSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SBSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SBSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SBSkuResponse)(nil)).Elem()
}

func (o SBSkuResponsePtrOutput) ToSBSkuResponsePtrOutput() SBSkuResponsePtrOutput {
	return o
}

func (o SBSkuResponsePtrOutput) ToSBSkuResponsePtrOutputWithContext(ctx context.Context) SBSkuResponsePtrOutput {
	return o
}

func (o SBSkuResponsePtrOutput) Elem() SBSkuResponseOutput {
	return o.ApplyT(func(v *SBSkuResponse) SBSkuResponse {
		if v != nil {
			return *v
		}
		var ret SBSkuResponse
		return ret
	}).(SBSkuResponseOutput)
}

func (o SBSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SBSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SBSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SBSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SBSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SBSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SqlFilter struct {
	CompatibilityLevel    *int    `pulumi:"compatibilityLevel"`
	RequiresPreprocessing *bool   `pulumi:"requiresPreprocessing"`
	SqlExpression         *string `pulumi:"sqlExpression"`
}


func (val *SqlFilter) Defaults() *SqlFilter {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.CompatibilityLevel) {
		compatibilityLevel_ := 20
		tmp.CompatibilityLevel = &compatibilityLevel_
	}
	if isZero(tmp.RequiresPreprocessing) {
		requiresPreprocessing_ := true
		tmp.RequiresPreprocessing = &requiresPreprocessing_
	}
	return &tmp
}





type SqlFilterInput interface {
	pulumi.Input

	ToSqlFilterOutput() SqlFilterOutput
	ToSqlFilterOutputWithContext(context.Context) SqlFilterOutput
}

type SqlFilterArgs struct {
	CompatibilityLevel    pulumi.IntPtrInput    `pulumi:"compatibilityLevel"`
	RequiresPreprocessing pulumi.BoolPtrInput   `pulumi:"requiresPreprocessing"`
	SqlExpression         pulumi.StringPtrInput `pulumi:"sqlExpression"`
}


func (val *SqlFilterArgs) Defaults() *SqlFilterArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.CompatibilityLevel) {
		tmp.CompatibilityLevel = pulumi.IntPtr(20)
	}
	if isZero(tmp.RequiresPreprocessing) {
		tmp.RequiresPreprocessing = pulumi.BoolPtr(true)
	}
	return &tmp
}
func (SqlFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlFilter)(nil)).Elem()
}

func (i SqlFilterArgs) ToSqlFilterOutput() SqlFilterOutput {
	return i.ToSqlFilterOutputWithContext(context.Background())
}

func (i SqlFilterArgs) ToSqlFilterOutputWithContext(ctx context.Context) SqlFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlFilterOutput)
}

func (i SqlFilterArgs) ToSqlFilterPtrOutput() SqlFilterPtrOutput {
	return i.ToSqlFilterPtrOutputWithContext(context.Background())
}

func (i SqlFilterArgs) ToSqlFilterPtrOutputWithContext(ctx context.Context) SqlFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlFilterOutput).ToSqlFilterPtrOutputWithContext(ctx)
}









type SqlFilterPtrInput interface {
	pulumi.Input

	ToSqlFilterPtrOutput() SqlFilterPtrOutput
	ToSqlFilterPtrOutputWithContext(context.Context) SqlFilterPtrOutput
}

type sqlFilterPtrType SqlFilterArgs

func SqlFilterPtr(v *SqlFilterArgs) SqlFilterPtrInput {
	return (*sqlFilterPtrType)(v)
}

func (*sqlFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlFilter)(nil)).Elem()
}

func (i *sqlFilterPtrType) ToSqlFilterPtrOutput() SqlFilterPtrOutput {
	return i.ToSqlFilterPtrOutputWithContext(context.Background())
}

func (i *sqlFilterPtrType) ToSqlFilterPtrOutputWithContext(ctx context.Context) SqlFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlFilterPtrOutput)
}

type SqlFilterOutput struct{ *pulumi.OutputState }

func (SqlFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlFilter)(nil)).Elem()
}

func (o SqlFilterOutput) ToSqlFilterOutput() SqlFilterOutput {
	return o
}

func (o SqlFilterOutput) ToSqlFilterOutputWithContext(ctx context.Context) SqlFilterOutput {
	return o
}

func (o SqlFilterOutput) ToSqlFilterPtrOutput() SqlFilterPtrOutput {
	return o.ToSqlFilterPtrOutputWithContext(context.Background())
}

func (o SqlFilterOutput) ToSqlFilterPtrOutputWithContext(ctx context.Context) SqlFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlFilter) *SqlFilter {
		return &v
	}).(SqlFilterPtrOutput)
}

func (o SqlFilterOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlFilter) *int { return v.CompatibilityLevel }).(pulumi.IntPtrOutput)
}

func (o SqlFilterOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SqlFilter) *bool { return v.RequiresPreprocessing }).(pulumi.BoolPtrOutput)
}

func (o SqlFilterOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlFilter) *string { return v.SqlExpression }).(pulumi.StringPtrOutput)
}

type SqlFilterPtrOutput struct{ *pulumi.OutputState }

func (SqlFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlFilter)(nil)).Elem()
}

func (o SqlFilterPtrOutput) ToSqlFilterPtrOutput() SqlFilterPtrOutput {
	return o
}

func (o SqlFilterPtrOutput) ToSqlFilterPtrOutputWithContext(ctx context.Context) SqlFilterPtrOutput {
	return o
}

func (o SqlFilterPtrOutput) Elem() SqlFilterOutput {
	return o.ApplyT(func(v *SqlFilter) SqlFilter {
		if v != nil {
			return *v
		}
		var ret SqlFilter
		return ret
	}).(SqlFilterOutput)
}

func (o SqlFilterPtrOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlFilter) *int {
		if v == nil {
			return nil
		}
		return v.CompatibilityLevel
	}).(pulumi.IntPtrOutput)
}

func (o SqlFilterPtrOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlFilter) *bool {
		if v == nil {
			return nil
		}
		return v.RequiresPreprocessing
	}).(pulumi.BoolPtrOutput)
}

func (o SqlFilterPtrOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlFilter) *string {
		if v == nil {
			return nil
		}
		return v.SqlExpression
	}).(pulumi.StringPtrOutput)
}

type SqlFilterResponse struct {
	CompatibilityLevel    *int    `pulumi:"compatibilityLevel"`
	RequiresPreprocessing *bool   `pulumi:"requiresPreprocessing"`
	SqlExpression         *string `pulumi:"sqlExpression"`
}


func (val *SqlFilterResponse) Defaults() *SqlFilterResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.CompatibilityLevel) {
		compatibilityLevel_ := 20
		tmp.CompatibilityLevel = &compatibilityLevel_
	}
	if isZero(tmp.RequiresPreprocessing) {
		requiresPreprocessing_ := true
		tmp.RequiresPreprocessing = &requiresPreprocessing_
	}
	return &tmp
}

type SqlFilterResponseOutput struct{ *pulumi.OutputState }

func (SqlFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlFilterResponse)(nil)).Elem()
}

func (o SqlFilterResponseOutput) ToSqlFilterResponseOutput() SqlFilterResponseOutput {
	return o
}

func (o SqlFilterResponseOutput) ToSqlFilterResponseOutputWithContext(ctx context.Context) SqlFilterResponseOutput {
	return o
}

func (o SqlFilterResponseOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlFilterResponse) *int { return v.CompatibilityLevel }).(pulumi.IntPtrOutput)
}

func (o SqlFilterResponseOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SqlFilterResponse) *bool { return v.RequiresPreprocessing }).(pulumi.BoolPtrOutput)
}

func (o SqlFilterResponseOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlFilterResponse) *string { return v.SqlExpression }).(pulumi.StringPtrOutput)
}

type SqlFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (SqlFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlFilterResponse)(nil)).Elem()
}

func (o SqlFilterResponsePtrOutput) ToSqlFilterResponsePtrOutput() SqlFilterResponsePtrOutput {
	return o
}

func (o SqlFilterResponsePtrOutput) ToSqlFilterResponsePtrOutputWithContext(ctx context.Context) SqlFilterResponsePtrOutput {
	return o
}

func (o SqlFilterResponsePtrOutput) Elem() SqlFilterResponseOutput {
	return o.ApplyT(func(v *SqlFilterResponse) SqlFilterResponse {
		if v != nil {
			return *v
		}
		var ret SqlFilterResponse
		return ret
	}).(SqlFilterResponseOutput)
}

func (o SqlFilterResponsePtrOutput) CompatibilityLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlFilterResponse) *int {
		if v == nil {
			return nil
		}
		return v.CompatibilityLevel
	}).(pulumi.IntPtrOutput)
}

func (o SqlFilterResponsePtrOutput) RequiresPreprocessing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SqlFilterResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequiresPreprocessing
	}).(pulumi.BoolPtrOutput)
}

func (o SqlFilterResponsePtrOutput) SqlExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.SqlExpression
	}).(pulumi.StringPtrOutput)
}

type Subnet struct {
	Id string `pulumi:"id"`
}





type SubnetInput interface {
	pulumi.Input

	ToSubnetOutput() SubnetOutput
	ToSubnetOutputWithContext(context.Context) SubnetOutput
}

type SubnetArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (SubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnet)(nil)).Elem()
}

func (i SubnetArgs) ToSubnetOutput() SubnetOutput {
	return i.ToSubnetOutputWithContext(context.Background())
}

func (i SubnetArgs) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOutput)
}

func (i SubnetArgs) ToSubnetPtrOutput() SubnetPtrOutput {
	return i.ToSubnetPtrOutputWithContext(context.Background())
}

func (i SubnetArgs) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOutput).ToSubnetPtrOutputWithContext(ctx)
}









type SubnetPtrInput interface {
	pulumi.Input

	ToSubnetPtrOutput() SubnetPtrOutput
	ToSubnetPtrOutputWithContext(context.Context) SubnetPtrOutput
}

type subnetPtrType SubnetArgs

func SubnetPtr(v *SubnetArgs) SubnetPtrInput {
	return (*subnetPtrType)(v)
}

func (*subnetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Subnet)(nil)).Elem()
}

func (i *subnetPtrType) ToSubnetPtrOutput() SubnetPtrOutput {
	return i.ToSubnetPtrOutputWithContext(context.Background())
}

func (i *subnetPtrType) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetPtrOutput)
}

type SubnetOutput struct{ *pulumi.OutputState }

func (SubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnet)(nil)).Elem()
}

func (o SubnetOutput) ToSubnetOutput() SubnetOutput {
	return o
}

func (o SubnetOutput) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return o
}

func (o SubnetOutput) ToSubnetPtrOutput() SubnetPtrOutput {
	return o.ToSubnetPtrOutputWithContext(context.Background())
}

func (o SubnetOutput) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Subnet) *Subnet {
		return &v
	}).(SubnetPtrOutput)
}

func (o SubnetOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v Subnet) string { return v.Id }).(pulumi.StringOutput)
}

type SubnetPtrOutput struct{ *pulumi.OutputState }

func (SubnetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subnet)(nil)).Elem()
}

func (o SubnetPtrOutput) ToSubnetPtrOutput() SubnetPtrOutput {
	return o
}

func (o SubnetPtrOutput) ToSubnetPtrOutputWithContext(ctx context.Context) SubnetPtrOutput {
	return o
}

func (o SubnetPtrOutput) Elem() SubnetOutput {
	return o.ApplyT(func(v *Subnet) Subnet {
		if v != nil {
			return *v
		}
		var ret Subnet
		return ret
	}).(SubnetOutput)
}

func (o SubnetPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subnet) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type SubnetResponse struct {
	Id string `pulumi:"id"`
}

type SubnetResponseOutput struct{ *pulumi.OutputState }

func (SubnetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResponse)(nil)).Elem()
}

func (o SubnetResponseOutput) ToSubnetResponseOutput() SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) ToSubnetResponseOutputWithContext(ctx context.Context) SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SubnetResponse) string { return v.Id }).(pulumi.StringOutput)
}

type SubnetResponsePtrOutput struct{ *pulumi.OutputState }

func (SubnetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetResponse)(nil)).Elem()
}

func (o SubnetResponsePtrOutput) ToSubnetResponsePtrOutput() SubnetResponsePtrOutput {
	return o
}

func (o SubnetResponsePtrOutput) ToSubnetResponsePtrOutputWithContext(ctx context.Context) SubnetResponsePtrOutput {
	return o
}

func (o SubnetResponsePtrOutput) Elem() SubnetResponseOutput {
	return o.ApplyT(func(v *SubnetResponse) SubnetResponse {
		if v != nil {
			return *v
		}
		var ret SubnetResponse
		return ret
	}).(SubnetResponseOutput)
}

func (o SubnetResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionOutput{})
	pulumi.RegisterOutputType(ActionPtrOutput{})
	pulumi.RegisterOutputType(ActionResponseOutput{})
	pulumi.RegisterOutputType(ActionResponsePtrOutput{})
	pulumi.RegisterOutputType(CorrelationFilterOutput{})
	pulumi.RegisterOutputType(CorrelationFilterPtrOutput{})
	pulumi.RegisterOutputType(CorrelationFilterResponseOutput{})
	pulumi.RegisterOutputType(CorrelationFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(MessageCountDetailsResponseOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesArrayOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesResponseOutput{})
	pulumi.RegisterOutputType(NWRuleSetIpRulesResponseArrayOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesArrayOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesResponseOutput{})
	pulumi.RegisterOutputType(NWRuleSetVirtualNetworkRulesResponseArrayOutput{})
	pulumi.RegisterOutputType(SBSkuOutput{})
	pulumi.RegisterOutputType(SBSkuPtrOutput{})
	pulumi.RegisterOutputType(SBSkuResponseOutput{})
	pulumi.RegisterOutputType(SBSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlFilterOutput{})
	pulumi.RegisterOutputType(SqlFilterPtrOutput{})
	pulumi.RegisterOutputType(SqlFilterResponseOutput{})
	pulumi.RegisterOutputType(SqlFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(SubnetOutput{})
	pulumi.RegisterOutputType(SubnetPtrOutput{})
	pulumi.RegisterOutputType(SubnetResponseOutput{})
	pulumi.RegisterOutputType(SubnetResponsePtrOutput{})
}
