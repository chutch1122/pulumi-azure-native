


package v20151201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationPackageResponse struct {
	Format             *string `pulumi:"format"`
	Id                 *string `pulumi:"id"`
	LastActivationTime *string `pulumi:"lastActivationTime"`
	State              *string `pulumi:"state"`
	StorageUrl         *string `pulumi:"storageUrl"`
	StorageUrlExpiry   *string `pulumi:"storageUrlExpiry"`
	Version            *string `pulumi:"version"`
}

type ApplicationPackageResponseOutput struct{ *pulumi.OutputState }

func (ApplicationPackageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackageResponse)(nil)).Elem()
}

func (o ApplicationPackageResponseOutput) ToApplicationPackageResponseOutput() ApplicationPackageResponseOutput {
	return o
}

func (o ApplicationPackageResponseOutput) ToApplicationPackageResponseOutputWithContext(ctx context.Context) ApplicationPackageResponseOutput {
	return o
}

func (o ApplicationPackageResponseOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o ApplicationPackageResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationPackageResponseOutput) LastActivationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) *string { return v.LastActivationTime }).(pulumi.StringPtrOutput)
}

func (o ApplicationPackageResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o ApplicationPackageResponseOutput) StorageUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) *string { return v.StorageUrl }).(pulumi.StringPtrOutput)
}

func (o ApplicationPackageResponseOutput) StorageUrlExpiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) *string { return v.StorageUrlExpiry }).(pulumi.StringPtrOutput)
}

func (o ApplicationPackageResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ApplicationPackageResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationPackageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationPackageResponse)(nil)).Elem()
}

func (o ApplicationPackageResponseArrayOutput) ToApplicationPackageResponseArrayOutput() ApplicationPackageResponseArrayOutput {
	return o
}

func (o ApplicationPackageResponseArrayOutput) ToApplicationPackageResponseArrayOutputWithContext(ctx context.Context) ApplicationPackageResponseArrayOutput {
	return o
}

func (o ApplicationPackageResponseArrayOutput) Index(i pulumi.IntInput) ApplicationPackageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationPackageResponse {
		return vs[0].([]ApplicationPackageResponse)[vs[1].(int)]
	}).(ApplicationPackageResponseOutput)
}

type AutoStorageBaseProperties struct {
	StorageAccountId string `pulumi:"storageAccountId"`
}





type AutoStorageBasePropertiesInput interface {
	pulumi.Input

	ToAutoStorageBasePropertiesOutput() AutoStorageBasePropertiesOutput
	ToAutoStorageBasePropertiesOutputWithContext(context.Context) AutoStorageBasePropertiesOutput
}

type AutoStorageBasePropertiesArgs struct {
	StorageAccountId pulumi.StringInput `pulumi:"storageAccountId"`
}

func (AutoStorageBasePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoStorageBaseProperties)(nil)).Elem()
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesOutput() AutoStorageBasePropertiesOutput {
	return i.ToAutoStorageBasePropertiesOutputWithContext(context.Background())
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesOutputWithContext(ctx context.Context) AutoStorageBasePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoStorageBasePropertiesOutput)
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return i.ToAutoStorageBasePropertiesPtrOutputWithContext(context.Background())
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoStorageBasePropertiesOutput).ToAutoStorageBasePropertiesPtrOutputWithContext(ctx)
}









type AutoStorageBasePropertiesPtrInput interface {
	pulumi.Input

	ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput
	ToAutoStorageBasePropertiesPtrOutputWithContext(context.Context) AutoStorageBasePropertiesPtrOutput
}

type autoStorageBasePropertiesPtrType AutoStorageBasePropertiesArgs

func AutoStorageBasePropertiesPtr(v *AutoStorageBasePropertiesArgs) AutoStorageBasePropertiesPtrInput {
	return (*autoStorageBasePropertiesPtrType)(v)
}

func (*autoStorageBasePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoStorageBaseProperties)(nil)).Elem()
}

func (i *autoStorageBasePropertiesPtrType) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return i.ToAutoStorageBasePropertiesPtrOutputWithContext(context.Background())
}

func (i *autoStorageBasePropertiesPtrType) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoStorageBasePropertiesPtrOutput)
}

type AutoStorageBasePropertiesOutput struct{ *pulumi.OutputState }

func (AutoStorageBasePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoStorageBaseProperties)(nil)).Elem()
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesOutput() AutoStorageBasePropertiesOutput {
	return o
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesOutputWithContext(ctx context.Context) AutoStorageBasePropertiesOutput {
	return o
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return o.ToAutoStorageBasePropertiesPtrOutputWithContext(context.Background())
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoStorageBaseProperties) *AutoStorageBaseProperties {
		return &v
	}).(AutoStorageBasePropertiesPtrOutput)
}

func (o AutoStorageBasePropertiesOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v AutoStorageBaseProperties) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

type AutoStorageBasePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AutoStorageBasePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoStorageBaseProperties)(nil)).Elem()
}

func (o AutoStorageBasePropertiesPtrOutput) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return o
}

func (o AutoStorageBasePropertiesPtrOutput) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return o
}

func (o AutoStorageBasePropertiesPtrOutput) Elem() AutoStorageBasePropertiesOutput {
	return o.ApplyT(func(v *AutoStorageBaseProperties) AutoStorageBaseProperties {
		if v != nil {
			return *v
		}
		var ret AutoStorageBaseProperties
		return ret
	}).(AutoStorageBasePropertiesOutput)
}

func (o AutoStorageBasePropertiesPtrOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoStorageBaseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountId
	}).(pulumi.StringPtrOutput)
}

type AutoStoragePropertiesResponse struct {
	LastKeySync      string `pulumi:"lastKeySync"`
	StorageAccountId string `pulumi:"storageAccountId"`
}

type AutoStoragePropertiesResponseOutput struct{ *pulumi.OutputState }

func (AutoStoragePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoStoragePropertiesResponse)(nil)).Elem()
}

func (o AutoStoragePropertiesResponseOutput) ToAutoStoragePropertiesResponseOutput() AutoStoragePropertiesResponseOutput {
	return o
}

func (o AutoStoragePropertiesResponseOutput) ToAutoStoragePropertiesResponseOutputWithContext(ctx context.Context) AutoStoragePropertiesResponseOutput {
	return o
}

func (o AutoStoragePropertiesResponseOutput) LastKeySync() pulumi.StringOutput {
	return o.ApplyT(func(v AutoStoragePropertiesResponse) string { return v.LastKeySync }).(pulumi.StringOutput)
}

func (o AutoStoragePropertiesResponseOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v AutoStoragePropertiesResponse) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

type AutoStoragePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoStoragePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoStoragePropertiesResponse)(nil)).Elem()
}

func (o AutoStoragePropertiesResponsePtrOutput) ToAutoStoragePropertiesResponsePtrOutput() AutoStoragePropertiesResponsePtrOutput {
	return o
}

func (o AutoStoragePropertiesResponsePtrOutput) ToAutoStoragePropertiesResponsePtrOutputWithContext(ctx context.Context) AutoStoragePropertiesResponsePtrOutput {
	return o
}

func (o AutoStoragePropertiesResponsePtrOutput) Elem() AutoStoragePropertiesResponseOutput {
	return o.ApplyT(func(v *AutoStoragePropertiesResponse) AutoStoragePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AutoStoragePropertiesResponse
		return ret
	}).(AutoStoragePropertiesResponseOutput)
}

func (o AutoStoragePropertiesResponsePtrOutput) LastKeySync() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoStoragePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastKeySync
	}).(pulumi.StringPtrOutput)
}

func (o AutoStoragePropertiesResponsePtrOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoStoragePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationPackageResponseOutput{})
	pulumi.RegisterOutputType(ApplicationPackageResponseArrayOutput{})
	pulumi.RegisterOutputType(AutoStorageBasePropertiesOutput{})
	pulumi.RegisterOutputType(AutoStorageBasePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AutoStoragePropertiesResponseOutput{})
	pulumi.RegisterOutputType(AutoStoragePropertiesResponsePtrOutput{})
}
