


package v20170901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AutoUserScope string

const (
	AutoUserScopeTask = AutoUserScope("Task")
	AutoUserScopePool = AutoUserScope("Pool")
)

func (AutoUserScope) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoUserScope)(nil)).Elem()
}

func (e AutoUserScope) ToAutoUserScopeOutput() AutoUserScopeOutput {
	return pulumi.ToOutput(e).(AutoUserScopeOutput)
}

func (e AutoUserScope) ToAutoUserScopeOutputWithContext(ctx context.Context) AutoUserScopeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutoUserScopeOutput)
}

func (e AutoUserScope) ToAutoUserScopePtrOutput() AutoUserScopePtrOutput {
	return e.ToAutoUserScopePtrOutputWithContext(context.Background())
}

func (e AutoUserScope) ToAutoUserScopePtrOutputWithContext(ctx context.Context) AutoUserScopePtrOutput {
	return AutoUserScope(e).ToAutoUserScopeOutputWithContext(ctx).ToAutoUserScopePtrOutputWithContext(ctx)
}

func (e AutoUserScope) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoUserScope) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoUserScope) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutoUserScope) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutoUserScopeOutput struct{ *pulumi.OutputState }

func (AutoUserScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoUserScope)(nil)).Elem()
}

func (o AutoUserScopeOutput) ToAutoUserScopeOutput() AutoUserScopeOutput {
	return o
}

func (o AutoUserScopeOutput) ToAutoUserScopeOutputWithContext(ctx context.Context) AutoUserScopeOutput {
	return o
}

func (o AutoUserScopeOutput) ToAutoUserScopePtrOutput() AutoUserScopePtrOutput {
	return o.ToAutoUserScopePtrOutputWithContext(context.Background())
}

func (o AutoUserScopeOutput) ToAutoUserScopePtrOutputWithContext(ctx context.Context) AutoUserScopePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoUserScope) *AutoUserScope {
		return &v
	}).(AutoUserScopePtrOutput)
}

func (o AutoUserScopeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutoUserScopeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoUserScope) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutoUserScopeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoUserScopeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoUserScope) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutoUserScopePtrOutput struct{ *pulumi.OutputState }

func (AutoUserScopePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoUserScope)(nil)).Elem()
}

func (o AutoUserScopePtrOutput) ToAutoUserScopePtrOutput() AutoUserScopePtrOutput {
	return o
}

func (o AutoUserScopePtrOutput) ToAutoUserScopePtrOutputWithContext(ctx context.Context) AutoUserScopePtrOutput {
	return o
}

func (o AutoUserScopePtrOutput) Elem() AutoUserScopeOutput {
	return o.ApplyT(func(v *AutoUserScope) AutoUserScope {
		if v != nil {
			return *v
		}
		var ret AutoUserScope
		return ret
	}).(AutoUserScopeOutput)
}

func (o AutoUserScopePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoUserScopePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutoUserScope) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AutoUserScopeInput interface {
	pulumi.Input

	ToAutoUserScopeOutput() AutoUserScopeOutput
	ToAutoUserScopeOutputWithContext(context.Context) AutoUserScopeOutput
}

var autoUserScopePtrType = reflect.TypeOf((**AutoUserScope)(nil)).Elem()

type AutoUserScopePtrInput interface {
	pulumi.Input

	ToAutoUserScopePtrOutput() AutoUserScopePtrOutput
	ToAutoUserScopePtrOutputWithContext(context.Context) AutoUserScopePtrOutput
}

type autoUserScopePtr string

func AutoUserScopePtr(v string) AutoUserScopePtrInput {
	return (*autoUserScopePtr)(&v)
}

func (*autoUserScopePtr) ElementType() reflect.Type {
	return autoUserScopePtrType
}

func (in *autoUserScopePtr) ToAutoUserScopePtrOutput() AutoUserScopePtrOutput {
	return pulumi.ToOutput(in).(AutoUserScopePtrOutput)
}

func (in *autoUserScopePtr) ToAutoUserScopePtrOutputWithContext(ctx context.Context) AutoUserScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutoUserScopePtrOutput)
}

type CachingType string

const (
	CachingTypeNone      = CachingType("None")
	CachingTypeReadOnly  = CachingType("ReadOnly")
	CachingTypeReadWrite = CachingType("ReadWrite")
)

func (CachingType) ElementType() reflect.Type {
	return reflect.TypeOf((*CachingType)(nil)).Elem()
}

func (e CachingType) ToCachingTypeOutput() CachingTypeOutput {
	return pulumi.ToOutput(e).(CachingTypeOutput)
}

func (e CachingType) ToCachingTypeOutputWithContext(ctx context.Context) CachingTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CachingTypeOutput)
}

func (e CachingType) ToCachingTypePtrOutput() CachingTypePtrOutput {
	return e.ToCachingTypePtrOutputWithContext(context.Background())
}

func (e CachingType) ToCachingTypePtrOutputWithContext(ctx context.Context) CachingTypePtrOutput {
	return CachingType(e).ToCachingTypeOutputWithContext(ctx).ToCachingTypePtrOutputWithContext(ctx)
}

func (e CachingType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CachingType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CachingType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CachingType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CachingTypeOutput struct{ *pulumi.OutputState }

func (CachingTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CachingType)(nil)).Elem()
}

func (o CachingTypeOutput) ToCachingTypeOutput() CachingTypeOutput {
	return o
}

func (o CachingTypeOutput) ToCachingTypeOutputWithContext(ctx context.Context) CachingTypeOutput {
	return o
}

func (o CachingTypeOutput) ToCachingTypePtrOutput() CachingTypePtrOutput {
	return o.ToCachingTypePtrOutputWithContext(context.Background())
}

func (o CachingTypeOutput) ToCachingTypePtrOutputWithContext(ctx context.Context) CachingTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CachingType) *CachingType {
		return &v
	}).(CachingTypePtrOutput)
}

func (o CachingTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CachingTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CachingType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CachingTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CachingTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CachingType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CachingTypePtrOutput struct{ *pulumi.OutputState }

func (CachingTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CachingType)(nil)).Elem()
}

func (o CachingTypePtrOutput) ToCachingTypePtrOutput() CachingTypePtrOutput {
	return o
}

func (o CachingTypePtrOutput) ToCachingTypePtrOutputWithContext(ctx context.Context) CachingTypePtrOutput {
	return o
}

func (o CachingTypePtrOutput) Elem() CachingTypeOutput {
	return o.ApplyT(func(v *CachingType) CachingType {
		if v != nil {
			return *v
		}
		var ret CachingType
		return ret
	}).(CachingTypeOutput)
}

func (o CachingTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CachingTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CachingType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CachingTypeInput interface {
	pulumi.Input

	ToCachingTypeOutput() CachingTypeOutput
	ToCachingTypeOutputWithContext(context.Context) CachingTypeOutput
}

var cachingTypePtrType = reflect.TypeOf((**CachingType)(nil)).Elem()

type CachingTypePtrInput interface {
	pulumi.Input

	ToCachingTypePtrOutput() CachingTypePtrOutput
	ToCachingTypePtrOutputWithContext(context.Context) CachingTypePtrOutput
}

type cachingTypePtr string

func CachingTypePtr(v string) CachingTypePtrInput {
	return (*cachingTypePtr)(&v)
}

func (*cachingTypePtr) ElementType() reflect.Type {
	return cachingTypePtrType
}

func (in *cachingTypePtr) ToCachingTypePtrOutput() CachingTypePtrOutput {
	return pulumi.ToOutput(in).(CachingTypePtrOutput)
}

func (in *cachingTypePtr) ToCachingTypePtrOutputWithContext(ctx context.Context) CachingTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CachingTypePtrOutput)
}

type CertificateFormat string

const (
	CertificateFormatPfx = CertificateFormat("Pfx")
	CertificateFormatCer = CertificateFormat("Cer")
)

func (CertificateFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateFormat)(nil)).Elem()
}

func (e CertificateFormat) ToCertificateFormatOutput() CertificateFormatOutput {
	return pulumi.ToOutput(e).(CertificateFormatOutput)
}

func (e CertificateFormat) ToCertificateFormatOutputWithContext(ctx context.Context) CertificateFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CertificateFormatOutput)
}

func (e CertificateFormat) ToCertificateFormatPtrOutput() CertificateFormatPtrOutput {
	return e.ToCertificateFormatPtrOutputWithContext(context.Background())
}

func (e CertificateFormat) ToCertificateFormatPtrOutputWithContext(ctx context.Context) CertificateFormatPtrOutput {
	return CertificateFormat(e).ToCertificateFormatOutputWithContext(ctx).ToCertificateFormatPtrOutputWithContext(ctx)
}

func (e CertificateFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CertificateFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CertificateFormatOutput struct{ *pulumi.OutputState }

func (CertificateFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateFormat)(nil)).Elem()
}

func (o CertificateFormatOutput) ToCertificateFormatOutput() CertificateFormatOutput {
	return o
}

func (o CertificateFormatOutput) ToCertificateFormatOutputWithContext(ctx context.Context) CertificateFormatOutput {
	return o
}

func (o CertificateFormatOutput) ToCertificateFormatPtrOutput() CertificateFormatPtrOutput {
	return o.ToCertificateFormatPtrOutputWithContext(context.Background())
}

func (o CertificateFormatOutput) ToCertificateFormatPtrOutputWithContext(ctx context.Context) CertificateFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateFormat) *CertificateFormat {
		return &v
	}).(CertificateFormatPtrOutput)
}

func (o CertificateFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CertificateFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CertificateFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CertificateFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CertificateFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CertificateFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CertificateFormatPtrOutput struct{ *pulumi.OutputState }

func (CertificateFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateFormat)(nil)).Elem()
}

func (o CertificateFormatPtrOutput) ToCertificateFormatPtrOutput() CertificateFormatPtrOutput {
	return o
}

func (o CertificateFormatPtrOutput) ToCertificateFormatPtrOutputWithContext(ctx context.Context) CertificateFormatPtrOutput {
	return o
}

func (o CertificateFormatPtrOutput) Elem() CertificateFormatOutput {
	return o.ApplyT(func(v *CertificateFormat) CertificateFormat {
		if v != nil {
			return *v
		}
		var ret CertificateFormat
		return ret
	}).(CertificateFormatOutput)
}

func (o CertificateFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CertificateFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CertificateFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CertificateFormatInput interface {
	pulumi.Input

	ToCertificateFormatOutput() CertificateFormatOutput
	ToCertificateFormatOutputWithContext(context.Context) CertificateFormatOutput
}

var certificateFormatPtrType = reflect.TypeOf((**CertificateFormat)(nil)).Elem()

type CertificateFormatPtrInput interface {
	pulumi.Input

	ToCertificateFormatPtrOutput() CertificateFormatPtrOutput
	ToCertificateFormatPtrOutputWithContext(context.Context) CertificateFormatPtrOutput
}

type certificateFormatPtr string

func CertificateFormatPtr(v string) CertificateFormatPtrInput {
	return (*certificateFormatPtr)(&v)
}

func (*certificateFormatPtr) ElementType() reflect.Type {
	return certificateFormatPtrType
}

func (in *certificateFormatPtr) ToCertificateFormatPtrOutput() CertificateFormatPtrOutput {
	return pulumi.ToOutput(in).(CertificateFormatPtrOutput)
}

func (in *certificateFormatPtr) ToCertificateFormatPtrOutputWithContext(ctx context.Context) CertificateFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CertificateFormatPtrOutput)
}

type CertificateStoreLocation string

const (
	CertificateStoreLocationCurrentUser  = CertificateStoreLocation("CurrentUser")
	CertificateStoreLocationLocalMachine = CertificateStoreLocation("LocalMachine")
)

func (CertificateStoreLocation) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateStoreLocation)(nil)).Elem()
}

func (e CertificateStoreLocation) ToCertificateStoreLocationOutput() CertificateStoreLocationOutput {
	return pulumi.ToOutput(e).(CertificateStoreLocationOutput)
}

func (e CertificateStoreLocation) ToCertificateStoreLocationOutputWithContext(ctx context.Context) CertificateStoreLocationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CertificateStoreLocationOutput)
}

func (e CertificateStoreLocation) ToCertificateStoreLocationPtrOutput() CertificateStoreLocationPtrOutput {
	return e.ToCertificateStoreLocationPtrOutputWithContext(context.Background())
}

func (e CertificateStoreLocation) ToCertificateStoreLocationPtrOutputWithContext(ctx context.Context) CertificateStoreLocationPtrOutput {
	return CertificateStoreLocation(e).ToCertificateStoreLocationOutputWithContext(ctx).ToCertificateStoreLocationPtrOutputWithContext(ctx)
}

func (e CertificateStoreLocation) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateStoreLocation) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateStoreLocation) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CertificateStoreLocation) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CertificateStoreLocationOutput struct{ *pulumi.OutputState }

func (CertificateStoreLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateStoreLocation)(nil)).Elem()
}

func (o CertificateStoreLocationOutput) ToCertificateStoreLocationOutput() CertificateStoreLocationOutput {
	return o
}

func (o CertificateStoreLocationOutput) ToCertificateStoreLocationOutputWithContext(ctx context.Context) CertificateStoreLocationOutput {
	return o
}

func (o CertificateStoreLocationOutput) ToCertificateStoreLocationPtrOutput() CertificateStoreLocationPtrOutput {
	return o.ToCertificateStoreLocationPtrOutputWithContext(context.Background())
}

func (o CertificateStoreLocationOutput) ToCertificateStoreLocationPtrOutputWithContext(ctx context.Context) CertificateStoreLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateStoreLocation) *CertificateStoreLocation {
		return &v
	}).(CertificateStoreLocationPtrOutput)
}

func (o CertificateStoreLocationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CertificateStoreLocationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CertificateStoreLocation) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CertificateStoreLocationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CertificateStoreLocationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CertificateStoreLocation) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CertificateStoreLocationPtrOutput struct{ *pulumi.OutputState }

func (CertificateStoreLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateStoreLocation)(nil)).Elem()
}

func (o CertificateStoreLocationPtrOutput) ToCertificateStoreLocationPtrOutput() CertificateStoreLocationPtrOutput {
	return o
}

func (o CertificateStoreLocationPtrOutput) ToCertificateStoreLocationPtrOutputWithContext(ctx context.Context) CertificateStoreLocationPtrOutput {
	return o
}

func (o CertificateStoreLocationPtrOutput) Elem() CertificateStoreLocationOutput {
	return o.ApplyT(func(v *CertificateStoreLocation) CertificateStoreLocation {
		if v != nil {
			return *v
		}
		var ret CertificateStoreLocation
		return ret
	}).(CertificateStoreLocationOutput)
}

func (o CertificateStoreLocationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CertificateStoreLocationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CertificateStoreLocation) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CertificateStoreLocationInput interface {
	pulumi.Input

	ToCertificateStoreLocationOutput() CertificateStoreLocationOutput
	ToCertificateStoreLocationOutputWithContext(context.Context) CertificateStoreLocationOutput
}

var certificateStoreLocationPtrType = reflect.TypeOf((**CertificateStoreLocation)(nil)).Elem()

type CertificateStoreLocationPtrInput interface {
	pulumi.Input

	ToCertificateStoreLocationPtrOutput() CertificateStoreLocationPtrOutput
	ToCertificateStoreLocationPtrOutputWithContext(context.Context) CertificateStoreLocationPtrOutput
}

type certificateStoreLocationPtr string

func CertificateStoreLocationPtr(v string) CertificateStoreLocationPtrInput {
	return (*certificateStoreLocationPtr)(&v)
}

func (*certificateStoreLocationPtr) ElementType() reflect.Type {
	return certificateStoreLocationPtrType
}

func (in *certificateStoreLocationPtr) ToCertificateStoreLocationPtrOutput() CertificateStoreLocationPtrOutput {
	return pulumi.ToOutput(in).(CertificateStoreLocationPtrOutput)
}

func (in *certificateStoreLocationPtr) ToCertificateStoreLocationPtrOutputWithContext(ctx context.Context) CertificateStoreLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CertificateStoreLocationPtrOutput)
}

type CertificateVisibility string

const (
	CertificateVisibilityStartTask  = CertificateVisibility("StartTask")
	CertificateVisibilityTask       = CertificateVisibility("Task")
	CertificateVisibilityRemoteUser = CertificateVisibility("RemoteUser")
)

func (CertificateVisibility) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateVisibility)(nil)).Elem()
}

func (e CertificateVisibility) ToCertificateVisibilityOutput() CertificateVisibilityOutput {
	return pulumi.ToOutput(e).(CertificateVisibilityOutput)
}

func (e CertificateVisibility) ToCertificateVisibilityOutputWithContext(ctx context.Context) CertificateVisibilityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CertificateVisibilityOutput)
}

func (e CertificateVisibility) ToCertificateVisibilityPtrOutput() CertificateVisibilityPtrOutput {
	return e.ToCertificateVisibilityPtrOutputWithContext(context.Background())
}

func (e CertificateVisibility) ToCertificateVisibilityPtrOutputWithContext(ctx context.Context) CertificateVisibilityPtrOutput {
	return CertificateVisibility(e).ToCertificateVisibilityOutputWithContext(ctx).ToCertificateVisibilityPtrOutputWithContext(ctx)
}

func (e CertificateVisibility) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateVisibility) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CertificateVisibility) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CertificateVisibility) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CertificateVisibilityOutput struct{ *pulumi.OutputState }

func (CertificateVisibilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateVisibility)(nil)).Elem()
}

func (o CertificateVisibilityOutput) ToCertificateVisibilityOutput() CertificateVisibilityOutput {
	return o
}

func (o CertificateVisibilityOutput) ToCertificateVisibilityOutputWithContext(ctx context.Context) CertificateVisibilityOutput {
	return o
}

func (o CertificateVisibilityOutput) ToCertificateVisibilityPtrOutput() CertificateVisibilityPtrOutput {
	return o.ToCertificateVisibilityPtrOutputWithContext(context.Background())
}

func (o CertificateVisibilityOutput) ToCertificateVisibilityPtrOutputWithContext(ctx context.Context) CertificateVisibilityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificateVisibility) *CertificateVisibility {
		return &v
	}).(CertificateVisibilityPtrOutput)
}

func (o CertificateVisibilityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CertificateVisibilityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CertificateVisibility) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CertificateVisibilityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CertificateVisibilityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CertificateVisibility) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CertificateVisibilityPtrOutput struct{ *pulumi.OutputState }

func (CertificateVisibilityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateVisibility)(nil)).Elem()
}

func (o CertificateVisibilityPtrOutput) ToCertificateVisibilityPtrOutput() CertificateVisibilityPtrOutput {
	return o
}

func (o CertificateVisibilityPtrOutput) ToCertificateVisibilityPtrOutputWithContext(ctx context.Context) CertificateVisibilityPtrOutput {
	return o
}

func (o CertificateVisibilityPtrOutput) Elem() CertificateVisibilityOutput {
	return o.ApplyT(func(v *CertificateVisibility) CertificateVisibility {
		if v != nil {
			return *v
		}
		var ret CertificateVisibility
		return ret
	}).(CertificateVisibilityOutput)
}

func (o CertificateVisibilityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CertificateVisibilityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CertificateVisibility) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CertificateVisibilityInput interface {
	pulumi.Input

	ToCertificateVisibilityOutput() CertificateVisibilityOutput
	ToCertificateVisibilityOutputWithContext(context.Context) CertificateVisibilityOutput
}

var certificateVisibilityPtrType = reflect.TypeOf((**CertificateVisibility)(nil)).Elem()

type CertificateVisibilityPtrInput interface {
	pulumi.Input

	ToCertificateVisibilityPtrOutput() CertificateVisibilityPtrOutput
	ToCertificateVisibilityPtrOutputWithContext(context.Context) CertificateVisibilityPtrOutput
}

type certificateVisibilityPtr string

func CertificateVisibilityPtr(v string) CertificateVisibilityPtrInput {
	return (*certificateVisibilityPtr)(&v)
}

func (*certificateVisibilityPtr) ElementType() reflect.Type {
	return certificateVisibilityPtrType
}

func (in *certificateVisibilityPtr) ToCertificateVisibilityPtrOutput() CertificateVisibilityPtrOutput {
	return pulumi.ToOutput(in).(CertificateVisibilityPtrOutput)
}

func (in *certificateVisibilityPtr) ToCertificateVisibilityPtrOutputWithContext(ctx context.Context) CertificateVisibilityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CertificateVisibilityPtrOutput)
}





type CertificateVisibilityArrayInput interface {
	pulumi.Input

	ToCertificateVisibilityArrayOutput() CertificateVisibilityArrayOutput
	ToCertificateVisibilityArrayOutputWithContext(context.Context) CertificateVisibilityArrayOutput
}

type CertificateVisibilityArray []CertificateVisibility

func (CertificateVisibilityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateVisibility)(nil)).Elem()
}

func (i CertificateVisibilityArray) ToCertificateVisibilityArrayOutput() CertificateVisibilityArrayOutput {
	return i.ToCertificateVisibilityArrayOutputWithContext(context.Background())
}

func (i CertificateVisibilityArray) ToCertificateVisibilityArrayOutputWithContext(ctx context.Context) CertificateVisibilityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateVisibilityArrayOutput)
}

type CertificateVisibilityArrayOutput struct{ *pulumi.OutputState }

func (CertificateVisibilityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateVisibility)(nil)).Elem()
}

func (o CertificateVisibilityArrayOutput) ToCertificateVisibilityArrayOutput() CertificateVisibilityArrayOutput {
	return o
}

func (o CertificateVisibilityArrayOutput) ToCertificateVisibilityArrayOutputWithContext(ctx context.Context) CertificateVisibilityArrayOutput {
	return o
}

func (o CertificateVisibilityArrayOutput) Index(i pulumi.IntInput) CertificateVisibilityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CertificateVisibility {
		return vs[0].([]CertificateVisibility)[vs[1].(int)]
	}).(CertificateVisibilityOutput)
}

type ComputeNodeDeallocationOption string

const (
	ComputeNodeDeallocationOptionRequeue        = ComputeNodeDeallocationOption("Requeue")
	ComputeNodeDeallocationOptionTerminate      = ComputeNodeDeallocationOption("Terminate")
	ComputeNodeDeallocationOptionTaskCompletion = ComputeNodeDeallocationOption("TaskCompletion")
	ComputeNodeDeallocationOptionRetainedData   = ComputeNodeDeallocationOption("RetainedData")
)

func (ComputeNodeDeallocationOption) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeNodeDeallocationOption)(nil)).Elem()
}

func (e ComputeNodeDeallocationOption) ToComputeNodeDeallocationOptionOutput() ComputeNodeDeallocationOptionOutput {
	return pulumi.ToOutput(e).(ComputeNodeDeallocationOptionOutput)
}

func (e ComputeNodeDeallocationOption) ToComputeNodeDeallocationOptionOutputWithContext(ctx context.Context) ComputeNodeDeallocationOptionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ComputeNodeDeallocationOptionOutput)
}

func (e ComputeNodeDeallocationOption) ToComputeNodeDeallocationOptionPtrOutput() ComputeNodeDeallocationOptionPtrOutput {
	return e.ToComputeNodeDeallocationOptionPtrOutputWithContext(context.Background())
}

func (e ComputeNodeDeallocationOption) ToComputeNodeDeallocationOptionPtrOutputWithContext(ctx context.Context) ComputeNodeDeallocationOptionPtrOutput {
	return ComputeNodeDeallocationOption(e).ToComputeNodeDeallocationOptionOutputWithContext(ctx).ToComputeNodeDeallocationOptionPtrOutputWithContext(ctx)
}

func (e ComputeNodeDeallocationOption) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeNodeDeallocationOption) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeNodeDeallocationOption) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ComputeNodeDeallocationOption) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ComputeNodeDeallocationOptionOutput struct{ *pulumi.OutputState }

func (ComputeNodeDeallocationOptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeNodeDeallocationOption)(nil)).Elem()
}

func (o ComputeNodeDeallocationOptionOutput) ToComputeNodeDeallocationOptionOutput() ComputeNodeDeallocationOptionOutput {
	return o
}

func (o ComputeNodeDeallocationOptionOutput) ToComputeNodeDeallocationOptionOutputWithContext(ctx context.Context) ComputeNodeDeallocationOptionOutput {
	return o
}

func (o ComputeNodeDeallocationOptionOutput) ToComputeNodeDeallocationOptionPtrOutput() ComputeNodeDeallocationOptionPtrOutput {
	return o.ToComputeNodeDeallocationOptionPtrOutputWithContext(context.Background())
}

func (o ComputeNodeDeallocationOptionOutput) ToComputeNodeDeallocationOptionPtrOutputWithContext(ctx context.Context) ComputeNodeDeallocationOptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeNodeDeallocationOption) *ComputeNodeDeallocationOption {
		return &v
	}).(ComputeNodeDeallocationOptionPtrOutput)
}

func (o ComputeNodeDeallocationOptionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ComputeNodeDeallocationOptionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComputeNodeDeallocationOption) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ComputeNodeDeallocationOptionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComputeNodeDeallocationOptionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComputeNodeDeallocationOption) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ComputeNodeDeallocationOptionPtrOutput struct{ *pulumi.OutputState }

func (ComputeNodeDeallocationOptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeNodeDeallocationOption)(nil)).Elem()
}

func (o ComputeNodeDeallocationOptionPtrOutput) ToComputeNodeDeallocationOptionPtrOutput() ComputeNodeDeallocationOptionPtrOutput {
	return o
}

func (o ComputeNodeDeallocationOptionPtrOutput) ToComputeNodeDeallocationOptionPtrOutputWithContext(ctx context.Context) ComputeNodeDeallocationOptionPtrOutput {
	return o
}

func (o ComputeNodeDeallocationOptionPtrOutput) Elem() ComputeNodeDeallocationOptionOutput {
	return o.ApplyT(func(v *ComputeNodeDeallocationOption) ComputeNodeDeallocationOption {
		if v != nil {
			return *v
		}
		var ret ComputeNodeDeallocationOption
		return ret
	}).(ComputeNodeDeallocationOptionOutput)
}

func (o ComputeNodeDeallocationOptionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComputeNodeDeallocationOptionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ComputeNodeDeallocationOption) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ComputeNodeDeallocationOptionInput interface {
	pulumi.Input

	ToComputeNodeDeallocationOptionOutput() ComputeNodeDeallocationOptionOutput
	ToComputeNodeDeallocationOptionOutputWithContext(context.Context) ComputeNodeDeallocationOptionOutput
}

var computeNodeDeallocationOptionPtrType = reflect.TypeOf((**ComputeNodeDeallocationOption)(nil)).Elem()

type ComputeNodeDeallocationOptionPtrInput interface {
	pulumi.Input

	ToComputeNodeDeallocationOptionPtrOutput() ComputeNodeDeallocationOptionPtrOutput
	ToComputeNodeDeallocationOptionPtrOutputWithContext(context.Context) ComputeNodeDeallocationOptionPtrOutput
}

type computeNodeDeallocationOptionPtr string

func ComputeNodeDeallocationOptionPtr(v string) ComputeNodeDeallocationOptionPtrInput {
	return (*computeNodeDeallocationOptionPtr)(&v)
}

func (*computeNodeDeallocationOptionPtr) ElementType() reflect.Type {
	return computeNodeDeallocationOptionPtrType
}

func (in *computeNodeDeallocationOptionPtr) ToComputeNodeDeallocationOptionPtrOutput() ComputeNodeDeallocationOptionPtrOutput {
	return pulumi.ToOutput(in).(ComputeNodeDeallocationOptionPtrOutput)
}

func (in *computeNodeDeallocationOptionPtr) ToComputeNodeDeallocationOptionPtrOutputWithContext(ctx context.Context) ComputeNodeDeallocationOptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ComputeNodeDeallocationOptionPtrOutput)
}

type ComputeNodeFillType string

const (
	ComputeNodeFillTypeSpread = ComputeNodeFillType("Spread")
	ComputeNodeFillTypePack   = ComputeNodeFillType("Pack")
)

func (ComputeNodeFillType) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeNodeFillType)(nil)).Elem()
}

func (e ComputeNodeFillType) ToComputeNodeFillTypeOutput() ComputeNodeFillTypeOutput {
	return pulumi.ToOutput(e).(ComputeNodeFillTypeOutput)
}

func (e ComputeNodeFillType) ToComputeNodeFillTypeOutputWithContext(ctx context.Context) ComputeNodeFillTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ComputeNodeFillTypeOutput)
}

func (e ComputeNodeFillType) ToComputeNodeFillTypePtrOutput() ComputeNodeFillTypePtrOutput {
	return e.ToComputeNodeFillTypePtrOutputWithContext(context.Background())
}

func (e ComputeNodeFillType) ToComputeNodeFillTypePtrOutputWithContext(ctx context.Context) ComputeNodeFillTypePtrOutput {
	return ComputeNodeFillType(e).ToComputeNodeFillTypeOutputWithContext(ctx).ToComputeNodeFillTypePtrOutputWithContext(ctx)
}

func (e ComputeNodeFillType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeNodeFillType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ComputeNodeFillType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ComputeNodeFillType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ComputeNodeFillTypeOutput struct{ *pulumi.OutputState }

func (ComputeNodeFillTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeNodeFillType)(nil)).Elem()
}

func (o ComputeNodeFillTypeOutput) ToComputeNodeFillTypeOutput() ComputeNodeFillTypeOutput {
	return o
}

func (o ComputeNodeFillTypeOutput) ToComputeNodeFillTypeOutputWithContext(ctx context.Context) ComputeNodeFillTypeOutput {
	return o
}

func (o ComputeNodeFillTypeOutput) ToComputeNodeFillTypePtrOutput() ComputeNodeFillTypePtrOutput {
	return o.ToComputeNodeFillTypePtrOutputWithContext(context.Background())
}

func (o ComputeNodeFillTypeOutput) ToComputeNodeFillTypePtrOutputWithContext(ctx context.Context) ComputeNodeFillTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeNodeFillType) *ComputeNodeFillType {
		return &v
	}).(ComputeNodeFillTypePtrOutput)
}

func (o ComputeNodeFillTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ComputeNodeFillTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComputeNodeFillType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ComputeNodeFillTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComputeNodeFillTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ComputeNodeFillType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ComputeNodeFillTypePtrOutput struct{ *pulumi.OutputState }

func (ComputeNodeFillTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeNodeFillType)(nil)).Elem()
}

func (o ComputeNodeFillTypePtrOutput) ToComputeNodeFillTypePtrOutput() ComputeNodeFillTypePtrOutput {
	return o
}

func (o ComputeNodeFillTypePtrOutput) ToComputeNodeFillTypePtrOutputWithContext(ctx context.Context) ComputeNodeFillTypePtrOutput {
	return o
}

func (o ComputeNodeFillTypePtrOutput) Elem() ComputeNodeFillTypeOutput {
	return o.ApplyT(func(v *ComputeNodeFillType) ComputeNodeFillType {
		if v != nil {
			return *v
		}
		var ret ComputeNodeFillType
		return ret
	}).(ComputeNodeFillTypeOutput)
}

func (o ComputeNodeFillTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ComputeNodeFillTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ComputeNodeFillType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ComputeNodeFillTypeInput interface {
	pulumi.Input

	ToComputeNodeFillTypeOutput() ComputeNodeFillTypeOutput
	ToComputeNodeFillTypeOutputWithContext(context.Context) ComputeNodeFillTypeOutput
}

var computeNodeFillTypePtrType = reflect.TypeOf((**ComputeNodeFillType)(nil)).Elem()

type ComputeNodeFillTypePtrInput interface {
	pulumi.Input

	ToComputeNodeFillTypePtrOutput() ComputeNodeFillTypePtrOutput
	ToComputeNodeFillTypePtrOutputWithContext(context.Context) ComputeNodeFillTypePtrOutput
}

type computeNodeFillTypePtr string

func ComputeNodeFillTypePtr(v string) ComputeNodeFillTypePtrInput {
	return (*computeNodeFillTypePtr)(&v)
}

func (*computeNodeFillTypePtr) ElementType() reflect.Type {
	return computeNodeFillTypePtrType
}

func (in *computeNodeFillTypePtr) ToComputeNodeFillTypePtrOutput() ComputeNodeFillTypePtrOutput {
	return pulumi.ToOutput(in).(ComputeNodeFillTypePtrOutput)
}

func (in *computeNodeFillTypePtr) ToComputeNodeFillTypePtrOutputWithContext(ctx context.Context) ComputeNodeFillTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ComputeNodeFillTypePtrOutput)
}

type ElevationLevel string

const (
	ElevationLevelNonAdmin = ElevationLevel("NonAdmin")
	ElevationLevelAdmin    = ElevationLevel("Admin")
)

func (ElevationLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*ElevationLevel)(nil)).Elem()
}

func (e ElevationLevel) ToElevationLevelOutput() ElevationLevelOutput {
	return pulumi.ToOutput(e).(ElevationLevelOutput)
}

func (e ElevationLevel) ToElevationLevelOutputWithContext(ctx context.Context) ElevationLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ElevationLevelOutput)
}

func (e ElevationLevel) ToElevationLevelPtrOutput() ElevationLevelPtrOutput {
	return e.ToElevationLevelPtrOutputWithContext(context.Background())
}

func (e ElevationLevel) ToElevationLevelPtrOutputWithContext(ctx context.Context) ElevationLevelPtrOutput {
	return ElevationLevel(e).ToElevationLevelOutputWithContext(ctx).ToElevationLevelPtrOutputWithContext(ctx)
}

func (e ElevationLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ElevationLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ElevationLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ElevationLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ElevationLevelOutput struct{ *pulumi.OutputState }

func (ElevationLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ElevationLevel)(nil)).Elem()
}

func (o ElevationLevelOutput) ToElevationLevelOutput() ElevationLevelOutput {
	return o
}

func (o ElevationLevelOutput) ToElevationLevelOutputWithContext(ctx context.Context) ElevationLevelOutput {
	return o
}

func (o ElevationLevelOutput) ToElevationLevelPtrOutput() ElevationLevelPtrOutput {
	return o.ToElevationLevelPtrOutputWithContext(context.Background())
}

func (o ElevationLevelOutput) ToElevationLevelPtrOutputWithContext(ctx context.Context) ElevationLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ElevationLevel) *ElevationLevel {
		return &v
	}).(ElevationLevelPtrOutput)
}

func (o ElevationLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ElevationLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ElevationLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ElevationLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ElevationLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ElevationLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ElevationLevelPtrOutput struct{ *pulumi.OutputState }

func (ElevationLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElevationLevel)(nil)).Elem()
}

func (o ElevationLevelPtrOutput) ToElevationLevelPtrOutput() ElevationLevelPtrOutput {
	return o
}

func (o ElevationLevelPtrOutput) ToElevationLevelPtrOutputWithContext(ctx context.Context) ElevationLevelPtrOutput {
	return o
}

func (o ElevationLevelPtrOutput) Elem() ElevationLevelOutput {
	return o.ApplyT(func(v *ElevationLevel) ElevationLevel {
		if v != nil {
			return *v
		}
		var ret ElevationLevel
		return ret
	}).(ElevationLevelOutput)
}

func (o ElevationLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ElevationLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ElevationLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ElevationLevelInput interface {
	pulumi.Input

	ToElevationLevelOutput() ElevationLevelOutput
	ToElevationLevelOutputWithContext(context.Context) ElevationLevelOutput
}

var elevationLevelPtrType = reflect.TypeOf((**ElevationLevel)(nil)).Elem()

type ElevationLevelPtrInput interface {
	pulumi.Input

	ToElevationLevelPtrOutput() ElevationLevelPtrOutput
	ToElevationLevelPtrOutputWithContext(context.Context) ElevationLevelPtrOutput
}

type elevationLevelPtr string

func ElevationLevelPtr(v string) ElevationLevelPtrInput {
	return (*elevationLevelPtr)(&v)
}

func (*elevationLevelPtr) ElementType() reflect.Type {
	return elevationLevelPtrType
}

func (in *elevationLevelPtr) ToElevationLevelPtrOutput() ElevationLevelPtrOutput {
	return pulumi.ToOutput(in).(ElevationLevelPtrOutput)
}

func (in *elevationLevelPtr) ToElevationLevelPtrOutputWithContext(ctx context.Context) ElevationLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ElevationLevelPtrOutput)
}

type InboundEndpointProtocol string

const (
	InboundEndpointProtocolTCP = InboundEndpointProtocol("TCP")
	InboundEndpointProtocolUDP = InboundEndpointProtocol("UDP")
)

func (InboundEndpointProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundEndpointProtocol)(nil)).Elem()
}

func (e InboundEndpointProtocol) ToInboundEndpointProtocolOutput() InboundEndpointProtocolOutput {
	return pulumi.ToOutput(e).(InboundEndpointProtocolOutput)
}

func (e InboundEndpointProtocol) ToInboundEndpointProtocolOutputWithContext(ctx context.Context) InboundEndpointProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InboundEndpointProtocolOutput)
}

func (e InboundEndpointProtocol) ToInboundEndpointProtocolPtrOutput() InboundEndpointProtocolPtrOutput {
	return e.ToInboundEndpointProtocolPtrOutputWithContext(context.Background())
}

func (e InboundEndpointProtocol) ToInboundEndpointProtocolPtrOutputWithContext(ctx context.Context) InboundEndpointProtocolPtrOutput {
	return InboundEndpointProtocol(e).ToInboundEndpointProtocolOutputWithContext(ctx).ToInboundEndpointProtocolPtrOutputWithContext(ctx)
}

func (e InboundEndpointProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InboundEndpointProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InboundEndpointProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InboundEndpointProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InboundEndpointProtocolOutput struct{ *pulumi.OutputState }

func (InboundEndpointProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundEndpointProtocol)(nil)).Elem()
}

func (o InboundEndpointProtocolOutput) ToInboundEndpointProtocolOutput() InboundEndpointProtocolOutput {
	return o
}

func (o InboundEndpointProtocolOutput) ToInboundEndpointProtocolOutputWithContext(ctx context.Context) InboundEndpointProtocolOutput {
	return o
}

func (o InboundEndpointProtocolOutput) ToInboundEndpointProtocolPtrOutput() InboundEndpointProtocolPtrOutput {
	return o.ToInboundEndpointProtocolPtrOutputWithContext(context.Background())
}

func (o InboundEndpointProtocolOutput) ToInboundEndpointProtocolPtrOutputWithContext(ctx context.Context) InboundEndpointProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InboundEndpointProtocol) *InboundEndpointProtocol {
		return &v
	}).(InboundEndpointProtocolPtrOutput)
}

func (o InboundEndpointProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InboundEndpointProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InboundEndpointProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InboundEndpointProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InboundEndpointProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InboundEndpointProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InboundEndpointProtocolPtrOutput struct{ *pulumi.OutputState }

func (InboundEndpointProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InboundEndpointProtocol)(nil)).Elem()
}

func (o InboundEndpointProtocolPtrOutput) ToInboundEndpointProtocolPtrOutput() InboundEndpointProtocolPtrOutput {
	return o
}

func (o InboundEndpointProtocolPtrOutput) ToInboundEndpointProtocolPtrOutputWithContext(ctx context.Context) InboundEndpointProtocolPtrOutput {
	return o
}

func (o InboundEndpointProtocolPtrOutput) Elem() InboundEndpointProtocolOutput {
	return o.ApplyT(func(v *InboundEndpointProtocol) InboundEndpointProtocol {
		if v != nil {
			return *v
		}
		var ret InboundEndpointProtocol
		return ret
	}).(InboundEndpointProtocolOutput)
}

func (o InboundEndpointProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InboundEndpointProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InboundEndpointProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type InboundEndpointProtocolInput interface {
	pulumi.Input

	ToInboundEndpointProtocolOutput() InboundEndpointProtocolOutput
	ToInboundEndpointProtocolOutputWithContext(context.Context) InboundEndpointProtocolOutput
}

var inboundEndpointProtocolPtrType = reflect.TypeOf((**InboundEndpointProtocol)(nil)).Elem()

type InboundEndpointProtocolPtrInput interface {
	pulumi.Input

	ToInboundEndpointProtocolPtrOutput() InboundEndpointProtocolPtrOutput
	ToInboundEndpointProtocolPtrOutputWithContext(context.Context) InboundEndpointProtocolPtrOutput
}

type inboundEndpointProtocolPtr string

func InboundEndpointProtocolPtr(v string) InboundEndpointProtocolPtrInput {
	return (*inboundEndpointProtocolPtr)(&v)
}

func (*inboundEndpointProtocolPtr) ElementType() reflect.Type {
	return inboundEndpointProtocolPtrType
}

func (in *inboundEndpointProtocolPtr) ToInboundEndpointProtocolPtrOutput() InboundEndpointProtocolPtrOutput {
	return pulumi.ToOutput(in).(InboundEndpointProtocolPtrOutput)
}

func (in *inboundEndpointProtocolPtr) ToInboundEndpointProtocolPtrOutputWithContext(ctx context.Context) InboundEndpointProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InboundEndpointProtocolPtrOutput)
}

type InterNodeCommunicationState string

const (
	InterNodeCommunicationStateEnabled  = InterNodeCommunicationState("Enabled")
	InterNodeCommunicationStateDisabled = InterNodeCommunicationState("Disabled")
)

func (InterNodeCommunicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*InterNodeCommunicationState)(nil)).Elem()
}

func (e InterNodeCommunicationState) ToInterNodeCommunicationStateOutput() InterNodeCommunicationStateOutput {
	return pulumi.ToOutput(e).(InterNodeCommunicationStateOutput)
}

func (e InterNodeCommunicationState) ToInterNodeCommunicationStateOutputWithContext(ctx context.Context) InterNodeCommunicationStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InterNodeCommunicationStateOutput)
}

func (e InterNodeCommunicationState) ToInterNodeCommunicationStatePtrOutput() InterNodeCommunicationStatePtrOutput {
	return e.ToInterNodeCommunicationStatePtrOutputWithContext(context.Background())
}

func (e InterNodeCommunicationState) ToInterNodeCommunicationStatePtrOutputWithContext(ctx context.Context) InterNodeCommunicationStatePtrOutput {
	return InterNodeCommunicationState(e).ToInterNodeCommunicationStateOutputWithContext(ctx).ToInterNodeCommunicationStatePtrOutputWithContext(ctx)
}

func (e InterNodeCommunicationState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InterNodeCommunicationState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InterNodeCommunicationState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InterNodeCommunicationState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InterNodeCommunicationStateOutput struct{ *pulumi.OutputState }

func (InterNodeCommunicationStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InterNodeCommunicationState)(nil)).Elem()
}

func (o InterNodeCommunicationStateOutput) ToInterNodeCommunicationStateOutput() InterNodeCommunicationStateOutput {
	return o
}

func (o InterNodeCommunicationStateOutput) ToInterNodeCommunicationStateOutputWithContext(ctx context.Context) InterNodeCommunicationStateOutput {
	return o
}

func (o InterNodeCommunicationStateOutput) ToInterNodeCommunicationStatePtrOutput() InterNodeCommunicationStatePtrOutput {
	return o.ToInterNodeCommunicationStatePtrOutputWithContext(context.Background())
}

func (o InterNodeCommunicationStateOutput) ToInterNodeCommunicationStatePtrOutputWithContext(ctx context.Context) InterNodeCommunicationStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InterNodeCommunicationState) *InterNodeCommunicationState {
		return &v
	}).(InterNodeCommunicationStatePtrOutput)
}

func (o InterNodeCommunicationStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InterNodeCommunicationStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InterNodeCommunicationState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InterNodeCommunicationStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InterNodeCommunicationStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InterNodeCommunicationState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InterNodeCommunicationStatePtrOutput struct{ *pulumi.OutputState }

func (InterNodeCommunicationStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InterNodeCommunicationState)(nil)).Elem()
}

func (o InterNodeCommunicationStatePtrOutput) ToInterNodeCommunicationStatePtrOutput() InterNodeCommunicationStatePtrOutput {
	return o
}

func (o InterNodeCommunicationStatePtrOutput) ToInterNodeCommunicationStatePtrOutputWithContext(ctx context.Context) InterNodeCommunicationStatePtrOutput {
	return o
}

func (o InterNodeCommunicationStatePtrOutput) Elem() InterNodeCommunicationStateOutput {
	return o.ApplyT(func(v *InterNodeCommunicationState) InterNodeCommunicationState {
		if v != nil {
			return *v
		}
		var ret InterNodeCommunicationState
		return ret
	}).(InterNodeCommunicationStateOutput)
}

func (o InterNodeCommunicationStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InterNodeCommunicationStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InterNodeCommunicationState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type InterNodeCommunicationStateInput interface {
	pulumi.Input

	ToInterNodeCommunicationStateOutput() InterNodeCommunicationStateOutput
	ToInterNodeCommunicationStateOutputWithContext(context.Context) InterNodeCommunicationStateOutput
}

var interNodeCommunicationStatePtrType = reflect.TypeOf((**InterNodeCommunicationState)(nil)).Elem()

type InterNodeCommunicationStatePtrInput interface {
	pulumi.Input

	ToInterNodeCommunicationStatePtrOutput() InterNodeCommunicationStatePtrOutput
	ToInterNodeCommunicationStatePtrOutputWithContext(context.Context) InterNodeCommunicationStatePtrOutput
}

type interNodeCommunicationStatePtr string

func InterNodeCommunicationStatePtr(v string) InterNodeCommunicationStatePtrInput {
	return (*interNodeCommunicationStatePtr)(&v)
}

func (*interNodeCommunicationStatePtr) ElementType() reflect.Type {
	return interNodeCommunicationStatePtrType
}

func (in *interNodeCommunicationStatePtr) ToInterNodeCommunicationStatePtrOutput() InterNodeCommunicationStatePtrOutput {
	return pulumi.ToOutput(in).(InterNodeCommunicationStatePtrOutput)
}

func (in *interNodeCommunicationStatePtr) ToInterNodeCommunicationStatePtrOutputWithContext(ctx context.Context) InterNodeCommunicationStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InterNodeCommunicationStatePtrOutput)
}

type NetworkSecurityGroupRuleAccess string

const (
	NetworkSecurityGroupRuleAccessAllow = NetworkSecurityGroupRuleAccess("Allow")
	NetworkSecurityGroupRuleAccessDeny  = NetworkSecurityGroupRuleAccess("Deny")
)

func (NetworkSecurityGroupRuleAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupRuleAccess)(nil)).Elem()
}

func (e NetworkSecurityGroupRuleAccess) ToNetworkSecurityGroupRuleAccessOutput() NetworkSecurityGroupRuleAccessOutput {
	return pulumi.ToOutput(e).(NetworkSecurityGroupRuleAccessOutput)
}

func (e NetworkSecurityGroupRuleAccess) ToNetworkSecurityGroupRuleAccessOutputWithContext(ctx context.Context) NetworkSecurityGroupRuleAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkSecurityGroupRuleAccessOutput)
}

func (e NetworkSecurityGroupRuleAccess) ToNetworkSecurityGroupRuleAccessPtrOutput() NetworkSecurityGroupRuleAccessPtrOutput {
	return e.ToNetworkSecurityGroupRuleAccessPtrOutputWithContext(context.Background())
}

func (e NetworkSecurityGroupRuleAccess) ToNetworkSecurityGroupRuleAccessPtrOutputWithContext(ctx context.Context) NetworkSecurityGroupRuleAccessPtrOutput {
	return NetworkSecurityGroupRuleAccess(e).ToNetworkSecurityGroupRuleAccessOutputWithContext(ctx).ToNetworkSecurityGroupRuleAccessPtrOutputWithContext(ctx)
}

func (e NetworkSecurityGroupRuleAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkSecurityGroupRuleAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkSecurityGroupRuleAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkSecurityGroupRuleAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkSecurityGroupRuleAccessOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupRuleAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupRuleAccess)(nil)).Elem()
}

func (o NetworkSecurityGroupRuleAccessOutput) ToNetworkSecurityGroupRuleAccessOutput() NetworkSecurityGroupRuleAccessOutput {
	return o
}

func (o NetworkSecurityGroupRuleAccessOutput) ToNetworkSecurityGroupRuleAccessOutputWithContext(ctx context.Context) NetworkSecurityGroupRuleAccessOutput {
	return o
}

func (o NetworkSecurityGroupRuleAccessOutput) ToNetworkSecurityGroupRuleAccessPtrOutput() NetworkSecurityGroupRuleAccessPtrOutput {
	return o.ToNetworkSecurityGroupRuleAccessPtrOutputWithContext(context.Background())
}

func (o NetworkSecurityGroupRuleAccessOutput) ToNetworkSecurityGroupRuleAccessPtrOutputWithContext(ctx context.Context) NetworkSecurityGroupRuleAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkSecurityGroupRuleAccess) *NetworkSecurityGroupRuleAccess {
		return &v
	}).(NetworkSecurityGroupRuleAccessPtrOutput)
}

func (o NetworkSecurityGroupRuleAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkSecurityGroupRuleAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkSecurityGroupRuleAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkSecurityGroupRuleAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkSecurityGroupRuleAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkSecurityGroupRuleAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkSecurityGroupRuleAccessPtrOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupRuleAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityGroupRuleAccess)(nil)).Elem()
}

func (o NetworkSecurityGroupRuleAccessPtrOutput) ToNetworkSecurityGroupRuleAccessPtrOutput() NetworkSecurityGroupRuleAccessPtrOutput {
	return o
}

func (o NetworkSecurityGroupRuleAccessPtrOutput) ToNetworkSecurityGroupRuleAccessPtrOutputWithContext(ctx context.Context) NetworkSecurityGroupRuleAccessPtrOutput {
	return o
}

func (o NetworkSecurityGroupRuleAccessPtrOutput) Elem() NetworkSecurityGroupRuleAccessOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupRuleAccess) NetworkSecurityGroupRuleAccess {
		if v != nil {
			return *v
		}
		var ret NetworkSecurityGroupRuleAccess
		return ret
	}).(NetworkSecurityGroupRuleAccessOutput)
}

func (o NetworkSecurityGroupRuleAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkSecurityGroupRuleAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkSecurityGroupRuleAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NetworkSecurityGroupRuleAccessInput interface {
	pulumi.Input

	ToNetworkSecurityGroupRuleAccessOutput() NetworkSecurityGroupRuleAccessOutput
	ToNetworkSecurityGroupRuleAccessOutputWithContext(context.Context) NetworkSecurityGroupRuleAccessOutput
}

var networkSecurityGroupRuleAccessPtrType = reflect.TypeOf((**NetworkSecurityGroupRuleAccess)(nil)).Elem()

type NetworkSecurityGroupRuleAccessPtrInput interface {
	pulumi.Input

	ToNetworkSecurityGroupRuleAccessPtrOutput() NetworkSecurityGroupRuleAccessPtrOutput
	ToNetworkSecurityGroupRuleAccessPtrOutputWithContext(context.Context) NetworkSecurityGroupRuleAccessPtrOutput
}

type networkSecurityGroupRuleAccessPtr string

func NetworkSecurityGroupRuleAccessPtr(v string) NetworkSecurityGroupRuleAccessPtrInput {
	return (*networkSecurityGroupRuleAccessPtr)(&v)
}

func (*networkSecurityGroupRuleAccessPtr) ElementType() reflect.Type {
	return networkSecurityGroupRuleAccessPtrType
}

func (in *networkSecurityGroupRuleAccessPtr) ToNetworkSecurityGroupRuleAccessPtrOutput() NetworkSecurityGroupRuleAccessPtrOutput {
	return pulumi.ToOutput(in).(NetworkSecurityGroupRuleAccessPtrOutput)
}

func (in *networkSecurityGroupRuleAccessPtr) ToNetworkSecurityGroupRuleAccessPtrOutputWithContext(ctx context.Context) NetworkSecurityGroupRuleAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkSecurityGroupRuleAccessPtrOutput)
}

type PoolAllocationMode string

const (
	PoolAllocationModeBatchService     = PoolAllocationMode("BatchService")
	PoolAllocationModeUserSubscription = PoolAllocationMode("UserSubscription")
)

func (PoolAllocationMode) ElementType() reflect.Type {
	return reflect.TypeOf((*PoolAllocationMode)(nil)).Elem()
}

func (e PoolAllocationMode) ToPoolAllocationModeOutput() PoolAllocationModeOutput {
	return pulumi.ToOutput(e).(PoolAllocationModeOutput)
}

func (e PoolAllocationMode) ToPoolAllocationModeOutputWithContext(ctx context.Context) PoolAllocationModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PoolAllocationModeOutput)
}

func (e PoolAllocationMode) ToPoolAllocationModePtrOutput() PoolAllocationModePtrOutput {
	return e.ToPoolAllocationModePtrOutputWithContext(context.Background())
}

func (e PoolAllocationMode) ToPoolAllocationModePtrOutputWithContext(ctx context.Context) PoolAllocationModePtrOutput {
	return PoolAllocationMode(e).ToPoolAllocationModeOutputWithContext(ctx).ToPoolAllocationModePtrOutputWithContext(ctx)
}

func (e PoolAllocationMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PoolAllocationMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PoolAllocationMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PoolAllocationMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PoolAllocationModeOutput struct{ *pulumi.OutputState }

func (PoolAllocationModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PoolAllocationMode)(nil)).Elem()
}

func (o PoolAllocationModeOutput) ToPoolAllocationModeOutput() PoolAllocationModeOutput {
	return o
}

func (o PoolAllocationModeOutput) ToPoolAllocationModeOutputWithContext(ctx context.Context) PoolAllocationModeOutput {
	return o
}

func (o PoolAllocationModeOutput) ToPoolAllocationModePtrOutput() PoolAllocationModePtrOutput {
	return o.ToPoolAllocationModePtrOutputWithContext(context.Background())
}

func (o PoolAllocationModeOutput) ToPoolAllocationModePtrOutputWithContext(ctx context.Context) PoolAllocationModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PoolAllocationMode) *PoolAllocationMode {
		return &v
	}).(PoolAllocationModePtrOutput)
}

func (o PoolAllocationModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PoolAllocationModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PoolAllocationMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PoolAllocationModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PoolAllocationModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PoolAllocationMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PoolAllocationModePtrOutput struct{ *pulumi.OutputState }

func (PoolAllocationModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PoolAllocationMode)(nil)).Elem()
}

func (o PoolAllocationModePtrOutput) ToPoolAllocationModePtrOutput() PoolAllocationModePtrOutput {
	return o
}

func (o PoolAllocationModePtrOutput) ToPoolAllocationModePtrOutputWithContext(ctx context.Context) PoolAllocationModePtrOutput {
	return o
}

func (o PoolAllocationModePtrOutput) Elem() PoolAllocationModeOutput {
	return o.ApplyT(func(v *PoolAllocationMode) PoolAllocationMode {
		if v != nil {
			return *v
		}
		var ret PoolAllocationMode
		return ret
	}).(PoolAllocationModeOutput)
}

func (o PoolAllocationModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PoolAllocationModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PoolAllocationMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PoolAllocationModeInput interface {
	pulumi.Input

	ToPoolAllocationModeOutput() PoolAllocationModeOutput
	ToPoolAllocationModeOutputWithContext(context.Context) PoolAllocationModeOutput
}

var poolAllocationModePtrType = reflect.TypeOf((**PoolAllocationMode)(nil)).Elem()

type PoolAllocationModePtrInput interface {
	pulumi.Input

	ToPoolAllocationModePtrOutput() PoolAllocationModePtrOutput
	ToPoolAllocationModePtrOutputWithContext(context.Context) PoolAllocationModePtrOutput
}

type poolAllocationModePtr string

func PoolAllocationModePtr(v string) PoolAllocationModePtrInput {
	return (*poolAllocationModePtr)(&v)
}

func (*poolAllocationModePtr) ElementType() reflect.Type {
	return poolAllocationModePtrType
}

func (in *poolAllocationModePtr) ToPoolAllocationModePtrOutput() PoolAllocationModePtrOutput {
	return pulumi.ToOutput(in).(PoolAllocationModePtrOutput)
}

func (in *poolAllocationModePtr) ToPoolAllocationModePtrOutputWithContext(ctx context.Context) PoolAllocationModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PoolAllocationModePtrOutput)
}

type StorageAccountType string

const (
	StorageAccountType_Standard_LRS = StorageAccountType("Standard_LRS")
	StorageAccountType_Premium_LRS  = StorageAccountType("Premium_LRS")
)

func (StorageAccountType) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountType)(nil)).Elem()
}

func (e StorageAccountType) ToStorageAccountTypeOutput() StorageAccountTypeOutput {
	return pulumi.ToOutput(e).(StorageAccountTypeOutput)
}

func (e StorageAccountType) ToStorageAccountTypeOutputWithContext(ctx context.Context) StorageAccountTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StorageAccountTypeOutput)
}

func (e StorageAccountType) ToStorageAccountTypePtrOutput() StorageAccountTypePtrOutput {
	return e.ToStorageAccountTypePtrOutputWithContext(context.Background())
}

func (e StorageAccountType) ToStorageAccountTypePtrOutputWithContext(ctx context.Context) StorageAccountTypePtrOutput {
	return StorageAccountType(e).ToStorageAccountTypeOutputWithContext(ctx).ToStorageAccountTypePtrOutputWithContext(ctx)
}

func (e StorageAccountType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageAccountType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageAccountType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StorageAccountType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StorageAccountTypeOutput struct{ *pulumi.OutputState }

func (StorageAccountTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountType)(nil)).Elem()
}

func (o StorageAccountTypeOutput) ToStorageAccountTypeOutput() StorageAccountTypeOutput {
	return o
}

func (o StorageAccountTypeOutput) ToStorageAccountTypeOutputWithContext(ctx context.Context) StorageAccountTypeOutput {
	return o
}

func (o StorageAccountTypeOutput) ToStorageAccountTypePtrOutput() StorageAccountTypePtrOutput {
	return o.ToStorageAccountTypePtrOutputWithContext(context.Background())
}

func (o StorageAccountTypeOutput) ToStorageAccountTypePtrOutputWithContext(ctx context.Context) StorageAccountTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountType) *StorageAccountType {
		return &v
	}).(StorageAccountTypePtrOutput)
}

func (o StorageAccountTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StorageAccountTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageAccountType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StorageAccountTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageAccountTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageAccountType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StorageAccountTypePtrOutput struct{ *pulumi.OutputState }

func (StorageAccountTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountType)(nil)).Elem()
}

func (o StorageAccountTypePtrOutput) ToStorageAccountTypePtrOutput() StorageAccountTypePtrOutput {
	return o
}

func (o StorageAccountTypePtrOutput) ToStorageAccountTypePtrOutputWithContext(ctx context.Context) StorageAccountTypePtrOutput {
	return o
}

func (o StorageAccountTypePtrOutput) Elem() StorageAccountTypeOutput {
	return o.ApplyT(func(v *StorageAccountType) StorageAccountType {
		if v != nil {
			return *v
		}
		var ret StorageAccountType
		return ret
	}).(StorageAccountTypeOutput)
}

func (o StorageAccountTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageAccountTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StorageAccountType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StorageAccountTypeInput interface {
	pulumi.Input

	ToStorageAccountTypeOutput() StorageAccountTypeOutput
	ToStorageAccountTypeOutputWithContext(context.Context) StorageAccountTypeOutput
}

var storageAccountTypePtrType = reflect.TypeOf((**StorageAccountType)(nil)).Elem()

type StorageAccountTypePtrInput interface {
	pulumi.Input

	ToStorageAccountTypePtrOutput() StorageAccountTypePtrOutput
	ToStorageAccountTypePtrOutputWithContext(context.Context) StorageAccountTypePtrOutput
}

type storageAccountTypePtr string

func StorageAccountTypePtr(v string) StorageAccountTypePtrInput {
	return (*storageAccountTypePtr)(&v)
}

func (*storageAccountTypePtr) ElementType() reflect.Type {
	return storageAccountTypePtrType
}

func (in *storageAccountTypePtr) ToStorageAccountTypePtrOutput() StorageAccountTypePtrOutput {
	return pulumi.ToOutput(in).(StorageAccountTypePtrOutput)
}

func (in *storageAccountTypePtr) ToStorageAccountTypePtrOutputWithContext(ctx context.Context) StorageAccountTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StorageAccountTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AutoUserScopeOutput{})
	pulumi.RegisterOutputType(AutoUserScopePtrOutput{})
	pulumi.RegisterOutputType(CachingTypeOutput{})
	pulumi.RegisterOutputType(CachingTypePtrOutput{})
	pulumi.RegisterOutputType(CertificateFormatOutput{})
	pulumi.RegisterOutputType(CertificateFormatPtrOutput{})
	pulumi.RegisterOutputType(CertificateStoreLocationOutput{})
	pulumi.RegisterOutputType(CertificateStoreLocationPtrOutput{})
	pulumi.RegisterOutputType(CertificateVisibilityOutput{})
	pulumi.RegisterOutputType(CertificateVisibilityPtrOutput{})
	pulumi.RegisterOutputType(CertificateVisibilityArrayOutput{})
	pulumi.RegisterOutputType(ComputeNodeDeallocationOptionOutput{})
	pulumi.RegisterOutputType(ComputeNodeDeallocationOptionPtrOutput{})
	pulumi.RegisterOutputType(ComputeNodeFillTypeOutput{})
	pulumi.RegisterOutputType(ComputeNodeFillTypePtrOutput{})
	pulumi.RegisterOutputType(ElevationLevelOutput{})
	pulumi.RegisterOutputType(ElevationLevelPtrOutput{})
	pulumi.RegisterOutputType(InboundEndpointProtocolOutput{})
	pulumi.RegisterOutputType(InboundEndpointProtocolPtrOutput{})
	pulumi.RegisterOutputType(InterNodeCommunicationStateOutput{})
	pulumi.RegisterOutputType(InterNodeCommunicationStatePtrOutput{})
	pulumi.RegisterOutputType(NetworkSecurityGroupRuleAccessOutput{})
	pulumi.RegisterOutputType(NetworkSecurityGroupRuleAccessPtrOutput{})
	pulumi.RegisterOutputType(PoolAllocationModeOutput{})
	pulumi.RegisterOutputType(PoolAllocationModePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountTypeOutput{})
	pulumi.RegisterOutputType(StorageAccountTypePtrOutput{})
}
