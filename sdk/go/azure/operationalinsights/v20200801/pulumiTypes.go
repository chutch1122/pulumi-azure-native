


package v20200801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterSku struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     *string  `pulumi:"name"`
}





type ClusterSkuInput interface {
	pulumi.Input

	ToClusterSkuOutput() ClusterSkuOutput
	ToClusterSkuOutputWithContext(context.Context) ClusterSkuOutput
}

type ClusterSkuArgs struct {
	Capacity pulumi.Float64PtrInput `pulumi:"capacity"`
	Name     pulumi.StringPtrInput  `pulumi:"name"`
}

func (ClusterSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSku)(nil)).Elem()
}

func (i ClusterSkuArgs) ToClusterSkuOutput() ClusterSkuOutput {
	return i.ToClusterSkuOutputWithContext(context.Background())
}

func (i ClusterSkuArgs) ToClusterSkuOutputWithContext(ctx context.Context) ClusterSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuOutput)
}

func (i ClusterSkuArgs) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return i.ToClusterSkuPtrOutputWithContext(context.Background())
}

func (i ClusterSkuArgs) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuOutput).ToClusterSkuPtrOutputWithContext(ctx)
}









type ClusterSkuPtrInput interface {
	pulumi.Input

	ToClusterSkuPtrOutput() ClusterSkuPtrOutput
	ToClusterSkuPtrOutputWithContext(context.Context) ClusterSkuPtrOutput
}

type clusterSkuPtrType ClusterSkuArgs

func ClusterSkuPtr(v *ClusterSkuArgs) ClusterSkuPtrInput {
	return (*clusterSkuPtrType)(v)
}

func (*clusterSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSku)(nil)).Elem()
}

func (i *clusterSkuPtrType) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return i.ToClusterSkuPtrOutputWithContext(context.Background())
}

func (i *clusterSkuPtrType) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSkuPtrOutput)
}

type ClusterSkuOutput struct{ *pulumi.OutputState }

func (ClusterSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSku)(nil)).Elem()
}

func (o ClusterSkuOutput) ToClusterSkuOutput() ClusterSkuOutput {
	return o
}

func (o ClusterSkuOutput) ToClusterSkuOutputWithContext(ctx context.Context) ClusterSkuOutput {
	return o
}

func (o ClusterSkuOutput) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return o.ToClusterSkuPtrOutputWithContext(context.Background())
}

func (o ClusterSkuOutput) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterSku) *ClusterSku {
		return &v
	}).(ClusterSkuPtrOutput)
}

func (o ClusterSkuOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ClusterSku) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o ClusterSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ClusterSkuPtrOutput struct{ *pulumi.OutputState }

func (ClusterSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSku)(nil)).Elem()
}

func (o ClusterSkuPtrOutput) ToClusterSkuPtrOutput() ClusterSkuPtrOutput {
	return o
}

func (o ClusterSkuPtrOutput) ToClusterSkuPtrOutputWithContext(ctx context.Context) ClusterSkuPtrOutput {
	return o
}

func (o ClusterSkuPtrOutput) Elem() ClusterSkuOutput {
	return o.ApplyT(func(v *ClusterSku) ClusterSku {
		if v != nil {
			return *v
		}
		var ret ClusterSku
		return ret
	}).(ClusterSkuOutput)
}

func (o ClusterSkuPtrOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ClusterSku) *float64 {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.Float64PtrOutput)
}

func (o ClusterSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type ClusterSkuResponse struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     *string  `pulumi:"name"`
}

type ClusterSkuResponseOutput struct{ *pulumi.OutputState }

func (ClusterSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSkuResponse)(nil)).Elem()
}

func (o ClusterSkuResponseOutput) ToClusterSkuResponseOutput() ClusterSkuResponseOutput {
	return o
}

func (o ClusterSkuResponseOutput) ToClusterSkuResponseOutputWithContext(ctx context.Context) ClusterSkuResponseOutput {
	return o
}

func (o ClusterSkuResponseOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ClusterSkuResponse) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o ClusterSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ClusterSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSkuResponse)(nil)).Elem()
}

func (o ClusterSkuResponsePtrOutput) ToClusterSkuResponsePtrOutput() ClusterSkuResponsePtrOutput {
	return o
}

func (o ClusterSkuResponsePtrOutput) ToClusterSkuResponsePtrOutputWithContext(ctx context.Context) ClusterSkuResponsePtrOutput {
	return o
}

func (o ClusterSkuResponsePtrOutput) Elem() ClusterSkuResponseOutput {
	return o.ApplyT(func(v *ClusterSkuResponse) ClusterSkuResponse {
		if v != nil {
			return *v
		}
		var ret ClusterSkuResponse
		return ret
	}).(ClusterSkuResponseOutput)
}

func (o ClusterSkuResponsePtrOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ClusterSkuResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.Float64PtrOutput)
}

func (o ClusterSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type Identity struct {
	Type IdentityType `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type IdentityTypeInput `pulumi:"type"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() IdentityTypeOutput {
	return o.ApplyT(func(v Identity) IdentityType { return v.Type }).(IdentityTypeOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() IdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *IdentityType {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(IdentityTypePtrOutput)
}

type IdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type KeyVaultProperties struct {
	KeyName     *string `pulumi:"keyName"`
	KeyVaultUri *string `pulumi:"keyVaultUri"`
	KeyVersion  *string `pulumi:"keyVersion"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	KeyName     pulumi.StringPtrInput `pulumi:"keyName"`
	KeyVaultUri pulumi.StringPtrInput `pulumi:"keyVaultUri"`
	KeyVersion  pulumi.StringPtrInput `pulumi:"keyVersion"`
}

func (KeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return i.ToKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput)
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput).ToKeyVaultPropertiesPtrOutputWithContext(ctx)
}









type KeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput
	ToKeyVaultPropertiesPtrOutputWithContext(context.Context) KeyVaultPropertiesPtrOutput
}

type keyVaultPropertiesPtrType KeyVaultPropertiesArgs

func KeyVaultPropertiesPtr(v *KeyVaultPropertiesArgs) KeyVaultPropertiesPtrInput {
	return (*keyVaultPropertiesPtrType)(v)
}

func (*keyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesPtrOutput)
}

type KeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultProperties) *KeyVaultProperties {
		return &v
	}).(KeyVaultPropertiesPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) Elem() KeyVaultPropertiesOutput {
	return o.ApplyT(func(v *KeyVaultProperties) KeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret KeyVaultProperties
		return ret
	}).(KeyVaultPropertiesOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponse struct {
	KeyName     *string `pulumi:"keyName"`
	KeyVaultUri *string `pulumi:"keyVaultUri"`
	KeyVersion  *string `pulumi:"keyVersion"`
}

type KeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVaultUri }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) Elem() KeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) KeyVaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultPropertiesResponse
		return ret
	}).(KeyVaultPropertiesResponseOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkScopedResourceResponse struct {
	ResourceId *string `pulumi:"resourceId"`
	ScopeId    *string `pulumi:"scopeId"`
}

type PrivateLinkScopedResourceResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkScopedResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkScopedResourceResponse)(nil)).Elem()
}

func (o PrivateLinkScopedResourceResponseOutput) ToPrivateLinkScopedResourceResponseOutput() PrivateLinkScopedResourceResponseOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseOutput) ToPrivateLinkScopedResourceResponseOutputWithContext(ctx context.Context) PrivateLinkScopedResourceResponseOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkScopedResourceResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkScopedResourceResponseOutput) ScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkScopedResourceResponse) *string { return v.ScopeId }).(pulumi.StringPtrOutput)
}

type PrivateLinkScopedResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkScopedResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkScopedResourceResponse)(nil)).Elem()
}

func (o PrivateLinkScopedResourceResponseArrayOutput) ToPrivateLinkScopedResourceResponseArrayOutput() PrivateLinkScopedResourceResponseArrayOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseArrayOutput) ToPrivateLinkScopedResourceResponseArrayOutputWithContext(ctx context.Context) PrivateLinkScopedResourceResponseArrayOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseArrayOutput) Index(i pulumi.IntInput) PrivateLinkScopedResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateLinkScopedResourceResponse {
		return vs[0].([]PrivateLinkScopedResourceResponse)[vs[1].(int)]
	}).(PrivateLinkScopedResourceResponseOutput)
}

type StorageAccount struct {
	Id  string `pulumi:"id"`
	Key string `pulumi:"key"`
}





type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(context.Context) StorageAccountOutput
}

type StorageAccountArgs struct {
	Id  pulumi.StringInput `pulumi:"id"`
	Key pulumi.StringInput `pulumi:"key"`
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (i StorageAccountArgs) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}

type StorageAccountOutput struct{ *pulumi.OutputState }

func (StorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (o StorageAccountOutput) ToStorageAccountOutput() StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccount) string { return v.Id }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccount) string { return v.Key }).(pulumi.StringOutput)
}

type StorageAccountResponse struct {
	Id  string `pulumi:"id"`
	Key string `pulumi:"key"`
}

type StorageAccountResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutput() StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutputWithContext(ctx context.Context) StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o StorageAccountResponseOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountResponse) string { return v.Key }).(pulumi.StringOutput)
}

type StorageInsightStatusResponse struct {
	Description *string `pulumi:"description"`
	State       string  `pulumi:"state"`
}

type StorageInsightStatusResponseOutput struct{ *pulumi.OutputState }

func (StorageInsightStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageInsightStatusResponse)(nil)).Elem()
}

func (o StorageInsightStatusResponseOutput) ToStorageInsightStatusResponseOutput() StorageInsightStatusResponseOutput {
	return o
}

func (o StorageInsightStatusResponseOutput) ToStorageInsightStatusResponseOutputWithContext(ctx context.Context) StorageInsightStatusResponseOutput {
	return o
}

func (o StorageInsightStatusResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageInsightStatusResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o StorageInsightStatusResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v StorageInsightStatusResponse) string { return v.State }).(pulumi.StringOutput)
}

type Tag struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type TagInput interface {
	pulumi.Input

	ToTagOutput() TagOutput
	ToTagOutputWithContext(context.Context) TagOutput
}

type TagArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (TagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Tag)(nil)).Elem()
}

func (i TagArgs) ToTagOutput() TagOutput {
	return i.ToTagOutputWithContext(context.Background())
}

func (i TagArgs) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOutput)
}





type TagArrayInput interface {
	pulumi.Input

	ToTagArrayOutput() TagArrayOutput
	ToTagArrayOutputWithContext(context.Context) TagArrayOutput
}

type TagArray []TagInput

func (TagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Tag)(nil)).Elem()
}

func (i TagArray) ToTagArrayOutput() TagArrayOutput {
	return i.ToTagArrayOutputWithContext(context.Background())
}

func (i TagArray) ToTagArrayOutputWithContext(ctx context.Context) TagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagArrayOutput)
}

type TagOutput struct{ *pulumi.OutputState }

func (TagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Tag)(nil)).Elem()
}

func (o TagOutput) ToTagOutput() TagOutput {
	return o
}

func (o TagOutput) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return o
}

func (o TagOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Tag) string { return v.Name }).(pulumi.StringOutput)
}

func (o TagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v Tag) string { return v.Value }).(pulumi.StringOutput)
}

type TagArrayOutput struct{ *pulumi.OutputState }

func (TagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Tag)(nil)).Elem()
}

func (o TagArrayOutput) ToTagArrayOutput() TagArrayOutput {
	return o
}

func (o TagArrayOutput) ToTagArrayOutputWithContext(ctx context.Context) TagArrayOutput {
	return o
}

func (o TagArrayOutput) Index(i pulumi.IntInput) TagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Tag {
		return vs[0].([]Tag)[vs[1].(int)]
	}).(TagOutput)
}

type TagResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type TagResponseOutput struct{ *pulumi.OutputState }

func (TagResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagResponse)(nil)).Elem()
}

func (o TagResponseOutput) ToTagResponseOutput() TagResponseOutput {
	return o
}

func (o TagResponseOutput) ToTagResponseOutputWithContext(ctx context.Context) TagResponseOutput {
	return o
}

func (o TagResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TagResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o TagResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TagResponse) string { return v.Value }).(pulumi.StringOutput)
}

type TagResponseArrayOutput struct{ *pulumi.OutputState }

func (TagResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagResponse)(nil)).Elem()
}

func (o TagResponseArrayOutput) ToTagResponseArrayOutput() TagResponseArrayOutput {
	return o
}

func (o TagResponseArrayOutput) ToTagResponseArrayOutputWithContext(ctx context.Context) TagResponseArrayOutput {
	return o
}

func (o TagResponseArrayOutput) Index(i pulumi.IntInput) TagResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TagResponse {
		return vs[0].([]TagResponse)[vs[1].(int)]
	}).(TagResponseOutput)
}

type WorkspaceCapping struct {
	DailyQuotaGb *float64 `pulumi:"dailyQuotaGb"`
}





type WorkspaceCappingInput interface {
	pulumi.Input

	ToWorkspaceCappingOutput() WorkspaceCappingOutput
	ToWorkspaceCappingOutputWithContext(context.Context) WorkspaceCappingOutput
}

type WorkspaceCappingArgs struct {
	DailyQuotaGb pulumi.Float64PtrInput `pulumi:"dailyQuotaGb"`
}

func (WorkspaceCappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCapping)(nil)).Elem()
}

func (i WorkspaceCappingArgs) ToWorkspaceCappingOutput() WorkspaceCappingOutput {
	return i.ToWorkspaceCappingOutputWithContext(context.Background())
}

func (i WorkspaceCappingArgs) ToWorkspaceCappingOutputWithContext(ctx context.Context) WorkspaceCappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCappingOutput)
}

func (i WorkspaceCappingArgs) ToWorkspaceCappingPtrOutput() WorkspaceCappingPtrOutput {
	return i.ToWorkspaceCappingPtrOutputWithContext(context.Background())
}

func (i WorkspaceCappingArgs) ToWorkspaceCappingPtrOutputWithContext(ctx context.Context) WorkspaceCappingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCappingOutput).ToWorkspaceCappingPtrOutputWithContext(ctx)
}









type WorkspaceCappingPtrInput interface {
	pulumi.Input

	ToWorkspaceCappingPtrOutput() WorkspaceCappingPtrOutput
	ToWorkspaceCappingPtrOutputWithContext(context.Context) WorkspaceCappingPtrOutput
}

type workspaceCappingPtrType WorkspaceCappingArgs

func WorkspaceCappingPtr(v *WorkspaceCappingArgs) WorkspaceCappingPtrInput {
	return (*workspaceCappingPtrType)(v)
}

func (*workspaceCappingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCapping)(nil)).Elem()
}

func (i *workspaceCappingPtrType) ToWorkspaceCappingPtrOutput() WorkspaceCappingPtrOutput {
	return i.ToWorkspaceCappingPtrOutputWithContext(context.Background())
}

func (i *workspaceCappingPtrType) ToWorkspaceCappingPtrOutputWithContext(ctx context.Context) WorkspaceCappingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCappingPtrOutput)
}

type WorkspaceCappingOutput struct{ *pulumi.OutputState }

func (WorkspaceCappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCapping)(nil)).Elem()
}

func (o WorkspaceCappingOutput) ToWorkspaceCappingOutput() WorkspaceCappingOutput {
	return o
}

func (o WorkspaceCappingOutput) ToWorkspaceCappingOutputWithContext(ctx context.Context) WorkspaceCappingOutput {
	return o
}

func (o WorkspaceCappingOutput) ToWorkspaceCappingPtrOutput() WorkspaceCappingPtrOutput {
	return o.ToWorkspaceCappingPtrOutputWithContext(context.Background())
}

func (o WorkspaceCappingOutput) ToWorkspaceCappingPtrOutputWithContext(ctx context.Context) WorkspaceCappingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceCapping) *WorkspaceCapping {
		return &v
	}).(WorkspaceCappingPtrOutput)
}

func (o WorkspaceCappingOutput) DailyQuotaGb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v WorkspaceCapping) *float64 { return v.DailyQuotaGb }).(pulumi.Float64PtrOutput)
}

type WorkspaceCappingPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceCappingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCapping)(nil)).Elem()
}

func (o WorkspaceCappingPtrOutput) ToWorkspaceCappingPtrOutput() WorkspaceCappingPtrOutput {
	return o
}

func (o WorkspaceCappingPtrOutput) ToWorkspaceCappingPtrOutputWithContext(ctx context.Context) WorkspaceCappingPtrOutput {
	return o
}

func (o WorkspaceCappingPtrOutput) Elem() WorkspaceCappingOutput {
	return o.ApplyT(func(v *WorkspaceCapping) WorkspaceCapping {
		if v != nil {
			return *v
		}
		var ret WorkspaceCapping
		return ret
	}).(WorkspaceCappingOutput)
}

func (o WorkspaceCappingPtrOutput) DailyQuotaGb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkspaceCapping) *float64 {
		if v == nil {
			return nil
		}
		return v.DailyQuotaGb
	}).(pulumi.Float64PtrOutput)
}

type WorkspaceCappingResponse struct {
	DailyQuotaGb        *float64 `pulumi:"dailyQuotaGb"`
	DataIngestionStatus string   `pulumi:"dataIngestionStatus"`
	QuotaNextResetTime  string   `pulumi:"quotaNextResetTime"`
}

type WorkspaceCappingResponseOutput struct{ *pulumi.OutputState }

func (WorkspaceCappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCappingResponse)(nil)).Elem()
}

func (o WorkspaceCappingResponseOutput) ToWorkspaceCappingResponseOutput() WorkspaceCappingResponseOutput {
	return o
}

func (o WorkspaceCappingResponseOutput) ToWorkspaceCappingResponseOutputWithContext(ctx context.Context) WorkspaceCappingResponseOutput {
	return o
}

func (o WorkspaceCappingResponseOutput) DailyQuotaGb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v WorkspaceCappingResponse) *float64 { return v.DailyQuotaGb }).(pulumi.Float64PtrOutput)
}

func (o WorkspaceCappingResponseOutput) DataIngestionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceCappingResponse) string { return v.DataIngestionStatus }).(pulumi.StringOutput)
}

func (o WorkspaceCappingResponseOutput) QuotaNextResetTime() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceCappingResponse) string { return v.QuotaNextResetTime }).(pulumi.StringOutput)
}

type WorkspaceCappingResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkspaceCappingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCappingResponse)(nil)).Elem()
}

func (o WorkspaceCappingResponsePtrOutput) ToWorkspaceCappingResponsePtrOutput() WorkspaceCappingResponsePtrOutput {
	return o
}

func (o WorkspaceCappingResponsePtrOutput) ToWorkspaceCappingResponsePtrOutputWithContext(ctx context.Context) WorkspaceCappingResponsePtrOutput {
	return o
}

func (o WorkspaceCappingResponsePtrOutput) Elem() WorkspaceCappingResponseOutput {
	return o.ApplyT(func(v *WorkspaceCappingResponse) WorkspaceCappingResponse {
		if v != nil {
			return *v
		}
		var ret WorkspaceCappingResponse
		return ret
	}).(WorkspaceCappingResponseOutput)
}

func (o WorkspaceCappingResponsePtrOutput) DailyQuotaGb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkspaceCappingResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.DailyQuotaGb
	}).(pulumi.Float64PtrOutput)
}

func (o WorkspaceCappingResponsePtrOutput) DataIngestionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceCappingResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DataIngestionStatus
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceCappingResponsePtrOutput) QuotaNextResetTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceCappingResponse) *string {
		if v == nil {
			return nil
		}
		return &v.QuotaNextResetTime
	}).(pulumi.StringPtrOutput)
}

type WorkspaceSku struct {
	CapacityReservationLevel *int   `pulumi:"capacityReservationLevel"`
	Name                     string `pulumi:"name"`
}





type WorkspaceSkuInput interface {
	pulumi.Input

	ToWorkspaceSkuOutput() WorkspaceSkuOutput
	ToWorkspaceSkuOutputWithContext(context.Context) WorkspaceSkuOutput
}

type WorkspaceSkuArgs struct {
	CapacityReservationLevel pulumi.IntPtrInput `pulumi:"capacityReservationLevel"`
	Name                     pulumi.StringInput `pulumi:"name"`
}

func (WorkspaceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSku)(nil)).Elem()
}

func (i WorkspaceSkuArgs) ToWorkspaceSkuOutput() WorkspaceSkuOutput {
	return i.ToWorkspaceSkuOutputWithContext(context.Background())
}

func (i WorkspaceSkuArgs) ToWorkspaceSkuOutputWithContext(ctx context.Context) WorkspaceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSkuOutput)
}

func (i WorkspaceSkuArgs) ToWorkspaceSkuPtrOutput() WorkspaceSkuPtrOutput {
	return i.ToWorkspaceSkuPtrOutputWithContext(context.Background())
}

func (i WorkspaceSkuArgs) ToWorkspaceSkuPtrOutputWithContext(ctx context.Context) WorkspaceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSkuOutput).ToWorkspaceSkuPtrOutputWithContext(ctx)
}









type WorkspaceSkuPtrInput interface {
	pulumi.Input

	ToWorkspaceSkuPtrOutput() WorkspaceSkuPtrOutput
	ToWorkspaceSkuPtrOutputWithContext(context.Context) WorkspaceSkuPtrOutput
}

type workspaceSkuPtrType WorkspaceSkuArgs

func WorkspaceSkuPtr(v *WorkspaceSkuArgs) WorkspaceSkuPtrInput {
	return (*workspaceSkuPtrType)(v)
}

func (*workspaceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceSku)(nil)).Elem()
}

func (i *workspaceSkuPtrType) ToWorkspaceSkuPtrOutput() WorkspaceSkuPtrOutput {
	return i.ToWorkspaceSkuPtrOutputWithContext(context.Background())
}

func (i *workspaceSkuPtrType) ToWorkspaceSkuPtrOutputWithContext(ctx context.Context) WorkspaceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSkuPtrOutput)
}

type WorkspaceSkuOutput struct{ *pulumi.OutputState }

func (WorkspaceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSku)(nil)).Elem()
}

func (o WorkspaceSkuOutput) ToWorkspaceSkuOutput() WorkspaceSkuOutput {
	return o
}

func (o WorkspaceSkuOutput) ToWorkspaceSkuOutputWithContext(ctx context.Context) WorkspaceSkuOutput {
	return o
}

func (o WorkspaceSkuOutput) ToWorkspaceSkuPtrOutput() WorkspaceSkuPtrOutput {
	return o.ToWorkspaceSkuPtrOutputWithContext(context.Background())
}

func (o WorkspaceSkuOutput) ToWorkspaceSkuPtrOutputWithContext(ctx context.Context) WorkspaceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceSku) *WorkspaceSku {
		return &v
	}).(WorkspaceSkuPtrOutput)
}

func (o WorkspaceSkuOutput) CapacityReservationLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkspaceSku) *int { return v.CapacityReservationLevel }).(pulumi.IntPtrOutput)
}

func (o WorkspaceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceSku) string { return v.Name }).(pulumi.StringOutput)
}

type WorkspaceSkuPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceSku)(nil)).Elem()
}

func (o WorkspaceSkuPtrOutput) ToWorkspaceSkuPtrOutput() WorkspaceSkuPtrOutput {
	return o
}

func (o WorkspaceSkuPtrOutput) ToWorkspaceSkuPtrOutputWithContext(ctx context.Context) WorkspaceSkuPtrOutput {
	return o
}

func (o WorkspaceSkuPtrOutput) Elem() WorkspaceSkuOutput {
	return o.ApplyT(func(v *WorkspaceSku) WorkspaceSku {
		if v != nil {
			return *v
		}
		var ret WorkspaceSku
		return ret
	}).(WorkspaceSkuOutput)
}

func (o WorkspaceSkuPtrOutput) CapacityReservationLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WorkspaceSku) *int {
		if v == nil {
			return nil
		}
		return v.CapacityReservationLevel
	}).(pulumi.IntPtrOutput)
}

func (o WorkspaceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type WorkspaceSkuResponse struct {
	CapacityReservationLevel    *int   `pulumi:"capacityReservationLevel"`
	LastSkuUpdate               string `pulumi:"lastSkuUpdate"`
	MaxCapacityReservationLevel int    `pulumi:"maxCapacityReservationLevel"`
	Name                        string `pulumi:"name"`
}

type WorkspaceSkuResponseOutput struct{ *pulumi.OutputState }

func (WorkspaceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSkuResponse)(nil)).Elem()
}

func (o WorkspaceSkuResponseOutput) ToWorkspaceSkuResponseOutput() WorkspaceSkuResponseOutput {
	return o
}

func (o WorkspaceSkuResponseOutput) ToWorkspaceSkuResponseOutputWithContext(ctx context.Context) WorkspaceSkuResponseOutput {
	return o
}

func (o WorkspaceSkuResponseOutput) CapacityReservationLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkspaceSkuResponse) *int { return v.CapacityReservationLevel }).(pulumi.IntPtrOutput)
}

func (o WorkspaceSkuResponseOutput) LastSkuUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceSkuResponse) string { return v.LastSkuUpdate }).(pulumi.StringOutput)
}

func (o WorkspaceSkuResponseOutput) MaxCapacityReservationLevel() pulumi.IntOutput {
	return o.ApplyT(func(v WorkspaceSkuResponse) int { return v.MaxCapacityReservationLevel }).(pulumi.IntOutput)
}

func (o WorkspaceSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type WorkspaceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkspaceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceSkuResponse)(nil)).Elem()
}

func (o WorkspaceSkuResponsePtrOutput) ToWorkspaceSkuResponsePtrOutput() WorkspaceSkuResponsePtrOutput {
	return o
}

func (o WorkspaceSkuResponsePtrOutput) ToWorkspaceSkuResponsePtrOutputWithContext(ctx context.Context) WorkspaceSkuResponsePtrOutput {
	return o
}

func (o WorkspaceSkuResponsePtrOutput) Elem() WorkspaceSkuResponseOutput {
	return o.ApplyT(func(v *WorkspaceSkuResponse) WorkspaceSkuResponse {
		if v != nil {
			return *v
		}
		var ret WorkspaceSkuResponse
		return ret
	}).(WorkspaceSkuResponseOutput)
}

func (o WorkspaceSkuResponsePtrOutput) CapacityReservationLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WorkspaceSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.CapacityReservationLevel
	}).(pulumi.IntPtrOutput)
}

func (o WorkspaceSkuResponsePtrOutput) LastSkuUpdate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastSkuUpdate
	}).(pulumi.StringPtrOutput)
}

func (o WorkspaceSkuResponsePtrOutput) MaxCapacityReservationLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WorkspaceSkuResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MaxCapacityReservationLevel
	}).(pulumi.IntPtrOutput)
}

func (o WorkspaceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterSkuOutput{})
	pulumi.RegisterOutputType(ClusterSkuPtrOutput{})
	pulumi.RegisterOutputType(ClusterSkuResponseOutput{})
	pulumi.RegisterOutputType(ClusterSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkScopedResourceResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkScopedResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(StorageAccountOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseOutput{})
	pulumi.RegisterOutputType(StorageInsightStatusResponseOutput{})
	pulumi.RegisterOutputType(TagOutput{})
	pulumi.RegisterOutputType(TagArrayOutput{})
	pulumi.RegisterOutputType(TagResponseOutput{})
	pulumi.RegisterOutputType(TagResponseArrayOutput{})
	pulumi.RegisterOutputType(WorkspaceCappingOutput{})
	pulumi.RegisterOutputType(WorkspaceCappingPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceCappingResponseOutput{})
	pulumi.RegisterOutputType(WorkspaceCappingResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkspaceSkuOutput{})
	pulumi.RegisterOutputType(WorkspaceSkuPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceSkuResponseOutput{})
	pulumi.RegisterOutputType(WorkspaceSkuResponsePtrOutput{})
}
