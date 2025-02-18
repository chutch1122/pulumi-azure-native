


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessReviewInstance struct {
	EndDateTime   *string `pulumi:"endDateTime"`
	StartDateTime *string `pulumi:"startDateTime"`
}





type AccessReviewInstanceInput interface {
	pulumi.Input

	ToAccessReviewInstanceOutput() AccessReviewInstanceOutput
	ToAccessReviewInstanceOutputWithContext(context.Context) AccessReviewInstanceOutput
}

type AccessReviewInstanceArgs struct {
	EndDateTime   pulumi.StringPtrInput `pulumi:"endDateTime"`
	StartDateTime pulumi.StringPtrInput `pulumi:"startDateTime"`
}

func (AccessReviewInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewInstance)(nil)).Elem()
}

func (i AccessReviewInstanceArgs) ToAccessReviewInstanceOutput() AccessReviewInstanceOutput {
	return i.ToAccessReviewInstanceOutputWithContext(context.Background())
}

func (i AccessReviewInstanceArgs) ToAccessReviewInstanceOutputWithContext(ctx context.Context) AccessReviewInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessReviewInstanceOutput)
}





type AccessReviewInstanceArrayInput interface {
	pulumi.Input

	ToAccessReviewInstanceArrayOutput() AccessReviewInstanceArrayOutput
	ToAccessReviewInstanceArrayOutputWithContext(context.Context) AccessReviewInstanceArrayOutput
}

type AccessReviewInstanceArray []AccessReviewInstanceInput

func (AccessReviewInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessReviewInstance)(nil)).Elem()
}

func (i AccessReviewInstanceArray) ToAccessReviewInstanceArrayOutput() AccessReviewInstanceArrayOutput {
	return i.ToAccessReviewInstanceArrayOutputWithContext(context.Background())
}

func (i AccessReviewInstanceArray) ToAccessReviewInstanceArrayOutputWithContext(ctx context.Context) AccessReviewInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessReviewInstanceArrayOutput)
}

type AccessReviewInstanceOutput struct{ *pulumi.OutputState }

func (AccessReviewInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewInstance)(nil)).Elem()
}

func (o AccessReviewInstanceOutput) ToAccessReviewInstanceOutput() AccessReviewInstanceOutput {
	return o
}

func (o AccessReviewInstanceOutput) ToAccessReviewInstanceOutputWithContext(ctx context.Context) AccessReviewInstanceOutput {
	return o
}

func (o AccessReviewInstanceOutput) EndDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewInstance) *string { return v.EndDateTime }).(pulumi.StringPtrOutput)
}

func (o AccessReviewInstanceOutput) StartDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewInstance) *string { return v.StartDateTime }).(pulumi.StringPtrOutput)
}

type AccessReviewInstanceArrayOutput struct{ *pulumi.OutputState }

func (AccessReviewInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessReviewInstance)(nil)).Elem()
}

func (o AccessReviewInstanceArrayOutput) ToAccessReviewInstanceArrayOutput() AccessReviewInstanceArrayOutput {
	return o
}

func (o AccessReviewInstanceArrayOutput) ToAccessReviewInstanceArrayOutputWithContext(ctx context.Context) AccessReviewInstanceArrayOutput {
	return o
}

func (o AccessReviewInstanceArrayOutput) Index(i pulumi.IntInput) AccessReviewInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessReviewInstance {
		return vs[0].([]AccessReviewInstance)[vs[1].(int)]
	}).(AccessReviewInstanceOutput)
}

type AccessReviewInstanceResponse struct {
	EndDateTime   *string `pulumi:"endDateTime"`
	Id            string  `pulumi:"id"`
	Name          string  `pulumi:"name"`
	StartDateTime *string `pulumi:"startDateTime"`
	Status        string  `pulumi:"status"`
	Type          string  `pulumi:"type"`
}

type AccessReviewInstanceResponseOutput struct{ *pulumi.OutputState }

func (AccessReviewInstanceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewInstanceResponse)(nil)).Elem()
}

func (o AccessReviewInstanceResponseOutput) ToAccessReviewInstanceResponseOutput() AccessReviewInstanceResponseOutput {
	return o
}

func (o AccessReviewInstanceResponseOutput) ToAccessReviewInstanceResponseOutputWithContext(ctx context.Context) AccessReviewInstanceResponseOutput {
	return o
}

func (o AccessReviewInstanceResponseOutput) EndDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewInstanceResponse) *string { return v.EndDateTime }).(pulumi.StringPtrOutput)
}

func (o AccessReviewInstanceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewInstanceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o AccessReviewInstanceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewInstanceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o AccessReviewInstanceResponseOutput) StartDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewInstanceResponse) *string { return v.StartDateTime }).(pulumi.StringPtrOutput)
}

func (o AccessReviewInstanceResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewInstanceResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o AccessReviewInstanceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewInstanceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type AccessReviewInstanceResponseArrayOutput struct{ *pulumi.OutputState }

func (AccessReviewInstanceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessReviewInstanceResponse)(nil)).Elem()
}

func (o AccessReviewInstanceResponseArrayOutput) ToAccessReviewInstanceResponseArrayOutput() AccessReviewInstanceResponseArrayOutput {
	return o
}

func (o AccessReviewInstanceResponseArrayOutput) ToAccessReviewInstanceResponseArrayOutputWithContext(ctx context.Context) AccessReviewInstanceResponseArrayOutput {
	return o
}

func (o AccessReviewInstanceResponseArrayOutput) Index(i pulumi.IntInput) AccessReviewInstanceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessReviewInstanceResponse {
		return vs[0].([]AccessReviewInstanceResponse)[vs[1].(int)]
	}).(AccessReviewInstanceResponseOutput)
}

type AccessReviewReviewer struct {
	PrincipalId *string `pulumi:"principalId"`
}





type AccessReviewReviewerInput interface {
	pulumi.Input

	ToAccessReviewReviewerOutput() AccessReviewReviewerOutput
	ToAccessReviewReviewerOutputWithContext(context.Context) AccessReviewReviewerOutput
}

type AccessReviewReviewerArgs struct {
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
}

func (AccessReviewReviewerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewReviewer)(nil)).Elem()
}

func (i AccessReviewReviewerArgs) ToAccessReviewReviewerOutput() AccessReviewReviewerOutput {
	return i.ToAccessReviewReviewerOutputWithContext(context.Background())
}

func (i AccessReviewReviewerArgs) ToAccessReviewReviewerOutputWithContext(ctx context.Context) AccessReviewReviewerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessReviewReviewerOutput)
}





type AccessReviewReviewerArrayInput interface {
	pulumi.Input

	ToAccessReviewReviewerArrayOutput() AccessReviewReviewerArrayOutput
	ToAccessReviewReviewerArrayOutputWithContext(context.Context) AccessReviewReviewerArrayOutput
}

type AccessReviewReviewerArray []AccessReviewReviewerInput

func (AccessReviewReviewerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessReviewReviewer)(nil)).Elem()
}

func (i AccessReviewReviewerArray) ToAccessReviewReviewerArrayOutput() AccessReviewReviewerArrayOutput {
	return i.ToAccessReviewReviewerArrayOutputWithContext(context.Background())
}

func (i AccessReviewReviewerArray) ToAccessReviewReviewerArrayOutputWithContext(ctx context.Context) AccessReviewReviewerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessReviewReviewerArrayOutput)
}

type AccessReviewReviewerOutput struct{ *pulumi.OutputState }

func (AccessReviewReviewerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewReviewer)(nil)).Elem()
}

func (o AccessReviewReviewerOutput) ToAccessReviewReviewerOutput() AccessReviewReviewerOutput {
	return o
}

func (o AccessReviewReviewerOutput) ToAccessReviewReviewerOutputWithContext(ctx context.Context) AccessReviewReviewerOutput {
	return o
}

func (o AccessReviewReviewerOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewReviewer) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

type AccessReviewReviewerArrayOutput struct{ *pulumi.OutputState }

func (AccessReviewReviewerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessReviewReviewer)(nil)).Elem()
}

func (o AccessReviewReviewerArrayOutput) ToAccessReviewReviewerArrayOutput() AccessReviewReviewerArrayOutput {
	return o
}

func (o AccessReviewReviewerArrayOutput) ToAccessReviewReviewerArrayOutputWithContext(ctx context.Context) AccessReviewReviewerArrayOutput {
	return o
}

func (o AccessReviewReviewerArrayOutput) Index(i pulumi.IntInput) AccessReviewReviewerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessReviewReviewer {
		return vs[0].([]AccessReviewReviewer)[vs[1].(int)]
	}).(AccessReviewReviewerOutput)
}

type AccessReviewReviewerResponse struct {
	PrincipalId   *string `pulumi:"principalId"`
	PrincipalType string  `pulumi:"principalType"`
}

type AccessReviewReviewerResponseOutput struct{ *pulumi.OutputState }

func (AccessReviewReviewerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewReviewerResponse)(nil)).Elem()
}

func (o AccessReviewReviewerResponseOutput) ToAccessReviewReviewerResponseOutput() AccessReviewReviewerResponseOutput {
	return o
}

func (o AccessReviewReviewerResponseOutput) ToAccessReviewReviewerResponseOutputWithContext(ctx context.Context) AccessReviewReviewerResponseOutput {
	return o
}

func (o AccessReviewReviewerResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessReviewReviewerResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o AccessReviewReviewerResponseOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v AccessReviewReviewerResponse) string { return v.PrincipalType }).(pulumi.StringOutput)
}

type AccessReviewReviewerResponseArrayOutput struct{ *pulumi.OutputState }

func (AccessReviewReviewerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessReviewReviewerResponse)(nil)).Elem()
}

func (o AccessReviewReviewerResponseArrayOutput) ToAccessReviewReviewerResponseArrayOutput() AccessReviewReviewerResponseArrayOutput {
	return o
}

func (o AccessReviewReviewerResponseArrayOutput) ToAccessReviewReviewerResponseArrayOutputWithContext(ctx context.Context) AccessReviewReviewerResponseArrayOutput {
	return o
}

func (o AccessReviewReviewerResponseArrayOutput) Index(i pulumi.IntInput) AccessReviewReviewerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessReviewReviewerResponse {
		return vs[0].([]AccessReviewReviewerResponse)[vs[1].(int)]
	}).(AccessReviewReviewerResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessReviewInstanceOutput{})
	pulumi.RegisterOutputType(AccessReviewInstanceArrayOutput{})
	pulumi.RegisterOutputType(AccessReviewInstanceResponseOutput{})
	pulumi.RegisterOutputType(AccessReviewInstanceResponseArrayOutput{})
	pulumi.RegisterOutputType(AccessReviewReviewerOutput{})
	pulumi.RegisterOutputType(AccessReviewReviewerArrayOutput{})
	pulumi.RegisterOutputType(AccessReviewReviewerResponseOutput{})
	pulumi.RegisterOutputType(AccessReviewReviewerResponseArrayOutput{})
}
