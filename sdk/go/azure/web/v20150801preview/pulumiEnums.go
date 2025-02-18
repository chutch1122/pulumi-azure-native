


package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectionParameterType string

const (
	ConnectionParameterTypeString       = ConnectionParameterType("string")
	ConnectionParameterTypeSecurestring = ConnectionParameterType("securestring")
	ConnectionParameterTypeSecureobject = ConnectionParameterType("secureobject")
	ConnectionParameterTypeInt          = ConnectionParameterType("int")
	ConnectionParameterTypeBool         = ConnectionParameterType("bool")
	ConnectionParameterTypeObject       = ConnectionParameterType("object")
	ConnectionParameterTypeArray        = ConnectionParameterType("array")
	ConnectionParameterTypeOauthSetting = ConnectionParameterType("oauthSetting")
	ConnectionParameterTypeConnection   = ConnectionParameterType("connection")
)

func (ConnectionParameterType) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionParameterType)(nil)).Elem()
}

func (e ConnectionParameterType) ToConnectionParameterTypeOutput() ConnectionParameterTypeOutput {
	return pulumi.ToOutput(e).(ConnectionParameterTypeOutput)
}

func (e ConnectionParameterType) ToConnectionParameterTypeOutputWithContext(ctx context.Context) ConnectionParameterTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConnectionParameterTypeOutput)
}

func (e ConnectionParameterType) ToConnectionParameterTypePtrOutput() ConnectionParameterTypePtrOutput {
	return e.ToConnectionParameterTypePtrOutputWithContext(context.Background())
}

func (e ConnectionParameterType) ToConnectionParameterTypePtrOutputWithContext(ctx context.Context) ConnectionParameterTypePtrOutput {
	return ConnectionParameterType(e).ToConnectionParameterTypeOutputWithContext(ctx).ToConnectionParameterTypePtrOutputWithContext(ctx)
}

func (e ConnectionParameterType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionParameterType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectionParameterType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConnectionParameterType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConnectionParameterTypeOutput struct{ *pulumi.OutputState }

func (ConnectionParameterTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionParameterType)(nil)).Elem()
}

func (o ConnectionParameterTypeOutput) ToConnectionParameterTypeOutput() ConnectionParameterTypeOutput {
	return o
}

func (o ConnectionParameterTypeOutput) ToConnectionParameterTypeOutputWithContext(ctx context.Context) ConnectionParameterTypeOutput {
	return o
}

func (o ConnectionParameterTypeOutput) ToConnectionParameterTypePtrOutput() ConnectionParameterTypePtrOutput {
	return o.ToConnectionParameterTypePtrOutputWithContext(context.Background())
}

func (o ConnectionParameterTypeOutput) ToConnectionParameterTypePtrOutputWithContext(ctx context.Context) ConnectionParameterTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionParameterType) *ConnectionParameterType {
		return &v
	}).(ConnectionParameterTypePtrOutput)
}

func (o ConnectionParameterTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConnectionParameterTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionParameterType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConnectionParameterTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionParameterTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectionParameterType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConnectionParameterTypePtrOutput struct{ *pulumi.OutputState }

func (ConnectionParameterTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionParameterType)(nil)).Elem()
}

func (o ConnectionParameterTypePtrOutput) ToConnectionParameterTypePtrOutput() ConnectionParameterTypePtrOutput {
	return o
}

func (o ConnectionParameterTypePtrOutput) ToConnectionParameterTypePtrOutputWithContext(ctx context.Context) ConnectionParameterTypePtrOutput {
	return o
}

func (o ConnectionParameterTypePtrOutput) Elem() ConnectionParameterTypeOutput {
	return o.ApplyT(func(v *ConnectionParameterType) ConnectionParameterType {
		if v != nil {
			return *v
		}
		var ret ConnectionParameterType
		return ret
	}).(ConnectionParameterTypeOutput)
}

func (o ConnectionParameterTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectionParameterTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConnectionParameterType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConnectionParameterTypeInput interface {
	pulumi.Input

	ToConnectionParameterTypeOutput() ConnectionParameterTypeOutput
	ToConnectionParameterTypeOutputWithContext(context.Context) ConnectionParameterTypeOutput
}

var connectionParameterTypePtrType = reflect.TypeOf((**ConnectionParameterType)(nil)).Elem()

type ConnectionParameterTypePtrInput interface {
	pulumi.Input

	ToConnectionParameterTypePtrOutput() ConnectionParameterTypePtrOutput
	ToConnectionParameterTypePtrOutputWithContext(context.Context) ConnectionParameterTypePtrOutput
}

type connectionParameterTypePtr string

func ConnectionParameterTypePtr(v string) ConnectionParameterTypePtrInput {
	return (*connectionParameterTypePtr)(&v)
}

func (*connectionParameterTypePtr) ElementType() reflect.Type {
	return connectionParameterTypePtrType
}

func (in *connectionParameterTypePtr) ToConnectionParameterTypePtrOutput() ConnectionParameterTypePtrOutput {
	return pulumi.ToOutput(in).(ConnectionParameterTypePtrOutput)
}

func (in *connectionParameterTypePtr) ToConnectionParameterTypePtrOutputWithContext(ctx context.Context) ConnectionParameterTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConnectionParameterTypePtrOutput)
}

type PrincipalType string

const (
	PrincipalTypeActiveDirectory  = PrincipalType("ActiveDirectory")
	PrincipalTypeConnection       = PrincipalType("Connection")
	PrincipalTypeMicrosoftAccount = PrincipalType("MicrosoftAccount")
)

func (PrincipalType) ElementType() reflect.Type {
	return reflect.TypeOf((*PrincipalType)(nil)).Elem()
}

func (e PrincipalType) ToPrincipalTypeOutput() PrincipalTypeOutput {
	return pulumi.ToOutput(e).(PrincipalTypeOutput)
}

func (e PrincipalType) ToPrincipalTypeOutputWithContext(ctx context.Context) PrincipalTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrincipalTypeOutput)
}

func (e PrincipalType) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return e.ToPrincipalTypePtrOutputWithContext(context.Background())
}

func (e PrincipalType) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return PrincipalType(e).ToPrincipalTypeOutputWithContext(ctx).ToPrincipalTypePtrOutputWithContext(ctx)
}

func (e PrincipalType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrincipalType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrincipalType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrincipalType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrincipalTypeOutput struct{ *pulumi.OutputState }

func (PrincipalTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrincipalType)(nil)).Elem()
}

func (o PrincipalTypeOutput) ToPrincipalTypeOutput() PrincipalTypeOutput {
	return o
}

func (o PrincipalTypeOutput) ToPrincipalTypeOutputWithContext(ctx context.Context) PrincipalTypeOutput {
	return o
}

func (o PrincipalTypeOutput) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return o.ToPrincipalTypePtrOutputWithContext(context.Background())
}

func (o PrincipalTypeOutput) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrincipalType) *PrincipalType {
		return &v
	}).(PrincipalTypePtrOutput)
}

func (o PrincipalTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrincipalTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrincipalType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrincipalTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrincipalTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrincipalType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrincipalTypePtrOutput struct{ *pulumi.OutputState }

func (PrincipalTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrincipalType)(nil)).Elem()
}

func (o PrincipalTypePtrOutput) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return o
}

func (o PrincipalTypePtrOutput) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return o
}

func (o PrincipalTypePtrOutput) Elem() PrincipalTypeOutput {
	return o.ApplyT(func(v *PrincipalType) PrincipalType {
		if v != nil {
			return *v
		}
		var ret PrincipalType
		return ret
	}).(PrincipalTypeOutput)
}

func (o PrincipalTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrincipalTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrincipalType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrincipalTypeInput interface {
	pulumi.Input

	ToPrincipalTypeOutput() PrincipalTypeOutput
	ToPrincipalTypeOutputWithContext(context.Context) PrincipalTypeOutput
}

var principalTypePtrType = reflect.TypeOf((**PrincipalType)(nil)).Elem()

type PrincipalTypePtrInput interface {
	pulumi.Input

	ToPrincipalTypePtrOutput() PrincipalTypePtrOutput
	ToPrincipalTypePtrOutputWithContext(context.Context) PrincipalTypePtrOutput
}

type principalTypePtr string

func PrincipalTypePtr(v string) PrincipalTypePtrInput {
	return (*principalTypePtr)(&v)
}

func (*principalTypePtr) ElementType() reflect.Type {
	return principalTypePtrType
}

func (in *principalTypePtr) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return pulumi.ToOutput(in).(PrincipalTypePtrOutput)
}

func (in *principalTypePtr) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrincipalTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectionParameterTypeOutput{})
	pulumi.RegisterOutputType(ConnectionParameterTypePtrOutput{})
	pulumi.RegisterOutputType(PrincipalTypeOutput{})
	pulumi.RegisterOutputType(PrincipalTypePtrOutput{})
}
