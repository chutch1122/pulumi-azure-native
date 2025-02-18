


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AKSIdentityType string

const (
	AKSIdentityTypeSystemAssigned = AKSIdentityType("SystemAssigned")
	AKSIdentityTypeUserAssigned   = AKSIdentityType("UserAssigned")
)

func (AKSIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSIdentityType)(nil)).Elem()
}

func (e AKSIdentityType) ToAKSIdentityTypeOutput() AKSIdentityTypeOutput {
	return pulumi.ToOutput(e).(AKSIdentityTypeOutput)
}

func (e AKSIdentityType) ToAKSIdentityTypeOutputWithContext(ctx context.Context) AKSIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AKSIdentityTypeOutput)
}

func (e AKSIdentityType) ToAKSIdentityTypePtrOutput() AKSIdentityTypePtrOutput {
	return e.ToAKSIdentityTypePtrOutputWithContext(context.Background())
}

func (e AKSIdentityType) ToAKSIdentityTypePtrOutputWithContext(ctx context.Context) AKSIdentityTypePtrOutput {
	return AKSIdentityType(e).ToAKSIdentityTypeOutputWithContext(ctx).ToAKSIdentityTypePtrOutputWithContext(ctx)
}

func (e AKSIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AKSIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AKSIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AKSIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AKSIdentityTypeOutput struct{ *pulumi.OutputState }

func (AKSIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AKSIdentityType)(nil)).Elem()
}

func (o AKSIdentityTypeOutput) ToAKSIdentityTypeOutput() AKSIdentityTypeOutput {
	return o
}

func (o AKSIdentityTypeOutput) ToAKSIdentityTypeOutputWithContext(ctx context.Context) AKSIdentityTypeOutput {
	return o
}

func (o AKSIdentityTypeOutput) ToAKSIdentityTypePtrOutput() AKSIdentityTypePtrOutput {
	return o.ToAKSIdentityTypePtrOutputWithContext(context.Background())
}

func (o AKSIdentityTypeOutput) ToAKSIdentityTypePtrOutputWithContext(ctx context.Context) AKSIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AKSIdentityType) *AKSIdentityType {
		return &v
	}).(AKSIdentityTypePtrOutput)
}

func (o AKSIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AKSIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AKSIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AKSIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AKSIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AKSIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AKSIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (AKSIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AKSIdentityType)(nil)).Elem()
}

func (o AKSIdentityTypePtrOutput) ToAKSIdentityTypePtrOutput() AKSIdentityTypePtrOutput {
	return o
}

func (o AKSIdentityTypePtrOutput) ToAKSIdentityTypePtrOutputWithContext(ctx context.Context) AKSIdentityTypePtrOutput {
	return o
}

func (o AKSIdentityTypePtrOutput) Elem() AKSIdentityTypeOutput {
	return o.ApplyT(func(v *AKSIdentityType) AKSIdentityType {
		if v != nil {
			return *v
		}
		var ret AKSIdentityType
		return ret
	}).(AKSIdentityTypeOutput)
}

func (o AKSIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AKSIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AKSIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AKSIdentityTypeInput interface {
	pulumi.Input

	ToAKSIdentityTypeOutput() AKSIdentityTypeOutput
	ToAKSIdentityTypeOutputWithContext(context.Context) AKSIdentityTypeOutput
}

var aksidentityTypePtrType = reflect.TypeOf((**AKSIdentityType)(nil)).Elem()

type AKSIdentityTypePtrInput interface {
	pulumi.Input

	ToAKSIdentityTypePtrOutput() AKSIdentityTypePtrOutput
	ToAKSIdentityTypePtrOutputWithContext(context.Context) AKSIdentityTypePtrOutput
}

type aksidentityTypePtr string

func AKSIdentityTypePtr(v string) AKSIdentityTypePtrInput {
	return (*aksidentityTypePtr)(&v)
}

func (*aksidentityTypePtr) ElementType() reflect.Type {
	return aksidentityTypePtrType
}

func (in *aksidentityTypePtr) ToAKSIdentityTypePtrOutput() AKSIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(AKSIdentityTypePtrOutput)
}

func (in *aksidentityTypePtr) ToAKSIdentityTypePtrOutputWithContext(ctx context.Context) AKSIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AKSIdentityTypePtrOutput)
}

type LevelType string

const (
	LevelTypeError       = LevelType("Error")
	LevelTypeWarning     = LevelType("Warning")
	LevelTypeInformation = LevelType("Information")
)

type OperatorScopeType string

const (
	OperatorScopeTypeCluster   = OperatorScopeType("cluster")
	OperatorScopeTypeNamespace = OperatorScopeType("namespace")
)

type OperatorType string

const (
	OperatorTypeFlux = OperatorType("Flux")
)

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

type ScopeType string

const (
	ScopeTypeCluster   = ScopeType("cluster")
	ScopeTypeNamespace = ScopeType("namespace")
)

type SourceKindType string

const (
	SourceKindTypeGitRepository = SourceKindType("GitRepository")
	SourceKindTypeBucket        = SourceKindType("Bucket")
)

func init() {
	pulumi.RegisterOutputType(AKSIdentityTypeOutput{})
	pulumi.RegisterOutputType(AKSIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
