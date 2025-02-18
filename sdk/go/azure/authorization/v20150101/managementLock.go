


package v20150101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ManagementLock struct {
	pulumi.CustomResourceState

	Level pulumi.StringPtrOutput `pulumi:"level"`
	Name  pulumi.StringPtrOutput `pulumi:"name"`
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	Type  pulumi.StringOutput    `pulumi:"type"`
}


func NewManagementLock(ctx *pulumi.Context,
	name string, args *ManagementLockArgs, opts ...pulumi.ResourceOption) (*ManagementLock, error) {
	if args == nil {
		args = &ManagementLockArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization:ManagementLock"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20160901:ManagementLock"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20170401:ManagementLock"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200501:ManagementLock"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagementLock
	err := ctx.RegisterResource("azure-native:authorization/v20150101:ManagementLock", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagementLock(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementLockState, opts ...pulumi.ResourceOption) (*ManagementLock, error) {
	var resource ManagementLock
	err := ctx.ReadResource("azure-native:authorization/v20150101:ManagementLock", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managementLockState struct {
}

type ManagementLockState struct {
}

func (ManagementLockState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementLockState)(nil)).Elem()
}

type managementLockArgs struct {
	Level    *string `pulumi:"level"`
	LockName *string `pulumi:"lockName"`
	Name     *string `pulumi:"name"`
	Notes    *string `pulumi:"notes"`
}


type ManagementLockArgs struct {
	Level    pulumi.StringPtrInput
	LockName pulumi.StringPtrInput
	Name     pulumi.StringPtrInput
	Notes    pulumi.StringPtrInput
}

func (ManagementLockArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementLockArgs)(nil)).Elem()
}

type ManagementLockInput interface {
	pulumi.Input

	ToManagementLockOutput() ManagementLockOutput
	ToManagementLockOutputWithContext(ctx context.Context) ManagementLockOutput
}

func (*ManagementLock) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementLock)(nil)).Elem()
}

func (i *ManagementLock) ToManagementLockOutput() ManagementLockOutput {
	return i.ToManagementLockOutputWithContext(context.Background())
}

func (i *ManagementLock) ToManagementLockOutputWithContext(ctx context.Context) ManagementLockOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockOutput)
}

type ManagementLockOutput struct{ *pulumi.OutputState }

func (ManagementLockOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementLock)(nil)).Elem()
}

func (o ManagementLockOutput) ToManagementLockOutput() ManagementLockOutput {
	return o
}

func (o ManagementLockOutput) ToManagementLockOutputWithContext(ctx context.Context) ManagementLockOutput {
	return o
}

func (o ManagementLockOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementLock) pulumi.StringPtrOutput { return v.Level }).(pulumi.StringPtrOutput)
}

func (o ManagementLockOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementLock) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ManagementLockOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementLock) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o ManagementLockOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementLock) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagementLockOutput{})
}
