


package v20160901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagementLockOwner struct {
	ApplicationId *string `pulumi:"applicationId"`
}





type ManagementLockOwnerInput interface {
	pulumi.Input

	ToManagementLockOwnerOutput() ManagementLockOwnerOutput
	ToManagementLockOwnerOutputWithContext(context.Context) ManagementLockOwnerOutput
}

type ManagementLockOwnerArgs struct {
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
}

func (ManagementLockOwnerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockOwner)(nil)).Elem()
}

func (i ManagementLockOwnerArgs) ToManagementLockOwnerOutput() ManagementLockOwnerOutput {
	return i.ToManagementLockOwnerOutputWithContext(context.Background())
}

func (i ManagementLockOwnerArgs) ToManagementLockOwnerOutputWithContext(ctx context.Context) ManagementLockOwnerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockOwnerOutput)
}





type ManagementLockOwnerArrayInput interface {
	pulumi.Input

	ToManagementLockOwnerArrayOutput() ManagementLockOwnerArrayOutput
	ToManagementLockOwnerArrayOutputWithContext(context.Context) ManagementLockOwnerArrayOutput
}

type ManagementLockOwnerArray []ManagementLockOwnerInput

func (ManagementLockOwnerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementLockOwner)(nil)).Elem()
}

func (i ManagementLockOwnerArray) ToManagementLockOwnerArrayOutput() ManagementLockOwnerArrayOutput {
	return i.ToManagementLockOwnerArrayOutputWithContext(context.Background())
}

func (i ManagementLockOwnerArray) ToManagementLockOwnerArrayOutputWithContext(ctx context.Context) ManagementLockOwnerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockOwnerArrayOutput)
}

type ManagementLockOwnerOutput struct{ *pulumi.OutputState }

func (ManagementLockOwnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockOwner)(nil)).Elem()
}

func (o ManagementLockOwnerOutput) ToManagementLockOwnerOutput() ManagementLockOwnerOutput {
	return o
}

func (o ManagementLockOwnerOutput) ToManagementLockOwnerOutputWithContext(ctx context.Context) ManagementLockOwnerOutput {
	return o
}

func (o ManagementLockOwnerOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementLockOwner) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

type ManagementLockOwnerArrayOutput struct{ *pulumi.OutputState }

func (ManagementLockOwnerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementLockOwner)(nil)).Elem()
}

func (o ManagementLockOwnerArrayOutput) ToManagementLockOwnerArrayOutput() ManagementLockOwnerArrayOutput {
	return o
}

func (o ManagementLockOwnerArrayOutput) ToManagementLockOwnerArrayOutputWithContext(ctx context.Context) ManagementLockOwnerArrayOutput {
	return o
}

func (o ManagementLockOwnerArrayOutput) Index(i pulumi.IntInput) ManagementLockOwnerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementLockOwner {
		return vs[0].([]ManagementLockOwner)[vs[1].(int)]
	}).(ManagementLockOwnerOutput)
}

type ManagementLockOwnerResponse struct {
	ApplicationId *string `pulumi:"applicationId"`
}

type ManagementLockOwnerResponseOutput struct{ *pulumi.OutputState }

func (ManagementLockOwnerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementLockOwnerResponse)(nil)).Elem()
}

func (o ManagementLockOwnerResponseOutput) ToManagementLockOwnerResponseOutput() ManagementLockOwnerResponseOutput {
	return o
}

func (o ManagementLockOwnerResponseOutput) ToManagementLockOwnerResponseOutputWithContext(ctx context.Context) ManagementLockOwnerResponseOutput {
	return o
}

func (o ManagementLockOwnerResponseOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementLockOwnerResponse) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

type ManagementLockOwnerResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagementLockOwnerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementLockOwnerResponse)(nil)).Elem()
}

func (o ManagementLockOwnerResponseArrayOutput) ToManagementLockOwnerResponseArrayOutput() ManagementLockOwnerResponseArrayOutput {
	return o
}

func (o ManagementLockOwnerResponseArrayOutput) ToManagementLockOwnerResponseArrayOutputWithContext(ctx context.Context) ManagementLockOwnerResponseArrayOutput {
	return o
}

func (o ManagementLockOwnerResponseArrayOutput) Index(i pulumi.IntInput) ManagementLockOwnerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementLockOwnerResponse {
		return vs[0].([]ManagementLockOwnerResponse)[vs[1].(int)]
	}).(ManagementLockOwnerResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagementLockOwnerOutput{})
	pulumi.RegisterOutputType(ManagementLockOwnerArrayOutput{})
	pulumi.RegisterOutputType(ManagementLockOwnerResponseOutput{})
	pulumi.RegisterOutputType(ManagementLockOwnerResponseArrayOutput{})
}
