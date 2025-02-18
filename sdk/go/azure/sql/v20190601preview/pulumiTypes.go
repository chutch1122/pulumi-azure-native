


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnectionPropertiesResponse struct {
	PrivateEndpoint                   *PrivateEndpointPropertyResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                             `pulumi:"provisioningState"`
}

type PrivateEndpointConnectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateEndpoint() PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointPropertyResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointPropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateLinkServiceConnectionStatePropertyResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type PrivateEndpointPropertyResponse struct {
	Id *string `pulumi:"id"`
}

type PrivateEndpointPropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointPropertyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponsePtrOutput) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o PrivateEndpointPropertyResponsePtrOutput) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o PrivateEndpointPropertyResponsePtrOutput) Elem() PrivateEndpointPropertyResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointPropertyResponse) PrivateEndpointPropertyResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointPropertyResponse
		return ret
	}).(PrivateEndpointPropertyResponseOutput)
}

func (o PrivateEndpointPropertyResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponse struct {
	ActionsRequired string `pulumi:"actionsRequired"`
	Description     string `pulumi:"description"`
	Status          string `pulumi:"status"`
}

type PrivateLinkServiceConnectionStatePropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.ActionsRequired }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Elem() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) PrivateLinkServiceConnectionStatePropertyResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStatePropertyResponse
		return ret
	}).(PrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentity struct {
	Type *string `pulumi:"type"`
}





type ResourceIdentityInput interface {
	pulumi.Input

	ToResourceIdentityOutput() ResourceIdentityOutput
	ToResourceIdentityOutputWithContext(context.Context) ResourceIdentityOutput
}

type ResourceIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (i ResourceIdentityArgs) ToResourceIdentityOutput() ResourceIdentityOutput {
	return i.ToResourceIdentityOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput)
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput).ToResourceIdentityPtrOutputWithContext(ctx)
}









type ResourceIdentityPtrInput interface {
	pulumi.Input

	ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput
	ToResourceIdentityPtrOutputWithContext(context.Context) ResourceIdentityPtrOutput
}

type resourceIdentityPtrType ResourceIdentityArgs

func ResourceIdentityPtr(v *ResourceIdentityArgs) ResourceIdentityPtrInput {
	return (*resourceIdentityPtrType)(v)
}

func (*resourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityPtrOutput)
}

type ResourceIdentityOutput struct{ *pulumi.OutputState }

func (ResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityOutput) ToResourceIdentityOutput() ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentity) *ResourceIdentity {
		return &v
	}).(ResourceIdentityPtrOutput)
}

func (o ResourceIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) Elem() ResourceIdentityOutput {
	return o.ApplyT(func(v *ResourceIdentity) ResourceIdentity {
		if v != nil {
			return *v
		}
		var ret ResourceIdentity
		return ret
	}).(ResourceIdentityOutput)
}

func (o ResourceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type ResourceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ResourceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) Elem() ResourceIdentityResponseOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) ResourceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityResponse
		return ret
	}).(ResourceIdentityResponseOutput)
}

func (o ResourceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ServerPrivateEndpointConnectionResponse struct {
	Id         string                                      `pulumi:"id"`
	Properties PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
}

type ServerPrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (ServerPrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o ServerPrivateEndpointConnectionResponseOutput) ToServerPrivateEndpointConnectionResponseOutput() ServerPrivateEndpointConnectionResponseOutput {
	return o
}

func (o ServerPrivateEndpointConnectionResponseOutput) ToServerPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) ServerPrivateEndpointConnectionResponseOutput {
	return o
}

func (o ServerPrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ServerPrivateEndpointConnectionResponseOutput) Properties() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v ServerPrivateEndpointConnectionResponse) PrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

type ServerPrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (ServerPrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerPrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o ServerPrivateEndpointConnectionResponseArrayOutput) ToServerPrivateEndpointConnectionResponseArrayOutput() ServerPrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o ServerPrivateEndpointConnectionResponseArrayOutput) ToServerPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) ServerPrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o ServerPrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) ServerPrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerPrivateEndpointConnectionResponse {
		return vs[0].([]ServerPrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(ServerPrivateEndpointConnectionResponseOutput)
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SyncGroupSchema struct {
	MasterSyncMemberName *string                `pulumi:"masterSyncMemberName"`
	Tables               []SyncGroupSchemaTable `pulumi:"tables"`
}





type SyncGroupSchemaInput interface {
	pulumi.Input

	ToSyncGroupSchemaOutput() SyncGroupSchemaOutput
	ToSyncGroupSchemaOutputWithContext(context.Context) SyncGroupSchemaOutput
}

type SyncGroupSchemaArgs struct {
	MasterSyncMemberName pulumi.StringPtrInput          `pulumi:"masterSyncMemberName"`
	Tables               SyncGroupSchemaTableArrayInput `pulumi:"tables"`
}

func (SyncGroupSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchema)(nil)).Elem()
}

func (i SyncGroupSchemaArgs) ToSyncGroupSchemaOutput() SyncGroupSchemaOutput {
	return i.ToSyncGroupSchemaOutputWithContext(context.Background())
}

func (i SyncGroupSchemaArgs) ToSyncGroupSchemaOutputWithContext(ctx context.Context) SyncGroupSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaOutput)
}

func (i SyncGroupSchemaArgs) ToSyncGroupSchemaPtrOutput() SyncGroupSchemaPtrOutput {
	return i.ToSyncGroupSchemaPtrOutputWithContext(context.Background())
}

func (i SyncGroupSchemaArgs) ToSyncGroupSchemaPtrOutputWithContext(ctx context.Context) SyncGroupSchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaOutput).ToSyncGroupSchemaPtrOutputWithContext(ctx)
}









type SyncGroupSchemaPtrInput interface {
	pulumi.Input

	ToSyncGroupSchemaPtrOutput() SyncGroupSchemaPtrOutput
	ToSyncGroupSchemaPtrOutputWithContext(context.Context) SyncGroupSchemaPtrOutput
}

type syncGroupSchemaPtrType SyncGroupSchemaArgs

func SyncGroupSchemaPtr(v *SyncGroupSchemaArgs) SyncGroupSchemaPtrInput {
	return (*syncGroupSchemaPtrType)(v)
}

func (*syncGroupSchemaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGroupSchema)(nil)).Elem()
}

func (i *syncGroupSchemaPtrType) ToSyncGroupSchemaPtrOutput() SyncGroupSchemaPtrOutput {
	return i.ToSyncGroupSchemaPtrOutputWithContext(context.Background())
}

func (i *syncGroupSchemaPtrType) ToSyncGroupSchemaPtrOutputWithContext(ctx context.Context) SyncGroupSchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaPtrOutput)
}

type SyncGroupSchemaOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchema)(nil)).Elem()
}

func (o SyncGroupSchemaOutput) ToSyncGroupSchemaOutput() SyncGroupSchemaOutput {
	return o
}

func (o SyncGroupSchemaOutput) ToSyncGroupSchemaOutputWithContext(ctx context.Context) SyncGroupSchemaOutput {
	return o
}

func (o SyncGroupSchemaOutput) ToSyncGroupSchemaPtrOutput() SyncGroupSchemaPtrOutput {
	return o.ToSyncGroupSchemaPtrOutputWithContext(context.Background())
}

func (o SyncGroupSchemaOutput) ToSyncGroupSchemaPtrOutputWithContext(ctx context.Context) SyncGroupSchemaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncGroupSchema) *SyncGroupSchema {
		return &v
	}).(SyncGroupSchemaPtrOutput)
}

func (o SyncGroupSchemaOutput) MasterSyncMemberName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchema) *string { return v.MasterSyncMemberName }).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaOutput) Tables() SyncGroupSchemaTableArrayOutput {
	return o.ApplyT(func(v SyncGroupSchema) []SyncGroupSchemaTable { return v.Tables }).(SyncGroupSchemaTableArrayOutput)
}

type SyncGroupSchemaPtrOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGroupSchema)(nil)).Elem()
}

func (o SyncGroupSchemaPtrOutput) ToSyncGroupSchemaPtrOutput() SyncGroupSchemaPtrOutput {
	return o
}

func (o SyncGroupSchemaPtrOutput) ToSyncGroupSchemaPtrOutputWithContext(ctx context.Context) SyncGroupSchemaPtrOutput {
	return o
}

func (o SyncGroupSchemaPtrOutput) Elem() SyncGroupSchemaOutput {
	return o.ApplyT(func(v *SyncGroupSchema) SyncGroupSchema {
		if v != nil {
			return *v
		}
		var ret SyncGroupSchema
		return ret
	}).(SyncGroupSchemaOutput)
}

func (o SyncGroupSchemaPtrOutput) MasterSyncMemberName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncGroupSchema) *string {
		if v == nil {
			return nil
		}
		return v.MasterSyncMemberName
	}).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaPtrOutput) Tables() SyncGroupSchemaTableArrayOutput {
	return o.ApplyT(func(v *SyncGroupSchema) []SyncGroupSchemaTable {
		if v == nil {
			return nil
		}
		return v.Tables
	}).(SyncGroupSchemaTableArrayOutput)
}

type SyncGroupSchemaResponse struct {
	MasterSyncMemberName *string                        `pulumi:"masterSyncMemberName"`
	Tables               []SyncGroupSchemaTableResponse `pulumi:"tables"`
}

type SyncGroupSchemaResponseOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaResponse)(nil)).Elem()
}

func (o SyncGroupSchemaResponseOutput) ToSyncGroupSchemaResponseOutput() SyncGroupSchemaResponseOutput {
	return o
}

func (o SyncGroupSchemaResponseOutput) ToSyncGroupSchemaResponseOutputWithContext(ctx context.Context) SyncGroupSchemaResponseOutput {
	return o
}

func (o SyncGroupSchemaResponseOutput) MasterSyncMemberName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaResponse) *string { return v.MasterSyncMemberName }).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaResponseOutput) Tables() SyncGroupSchemaTableResponseArrayOutput {
	return o.ApplyT(func(v SyncGroupSchemaResponse) []SyncGroupSchemaTableResponse { return v.Tables }).(SyncGroupSchemaTableResponseArrayOutput)
}

type SyncGroupSchemaResponsePtrOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncGroupSchemaResponse)(nil)).Elem()
}

func (o SyncGroupSchemaResponsePtrOutput) ToSyncGroupSchemaResponsePtrOutput() SyncGroupSchemaResponsePtrOutput {
	return o
}

func (o SyncGroupSchemaResponsePtrOutput) ToSyncGroupSchemaResponsePtrOutputWithContext(ctx context.Context) SyncGroupSchemaResponsePtrOutput {
	return o
}

func (o SyncGroupSchemaResponsePtrOutput) Elem() SyncGroupSchemaResponseOutput {
	return o.ApplyT(func(v *SyncGroupSchemaResponse) SyncGroupSchemaResponse {
		if v != nil {
			return *v
		}
		var ret SyncGroupSchemaResponse
		return ret
	}).(SyncGroupSchemaResponseOutput)
}

func (o SyncGroupSchemaResponsePtrOutput) MasterSyncMemberName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncGroupSchemaResponse) *string {
		if v == nil {
			return nil
		}
		return v.MasterSyncMemberName
	}).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaResponsePtrOutput) Tables() SyncGroupSchemaTableResponseArrayOutput {
	return o.ApplyT(func(v *SyncGroupSchemaResponse) []SyncGroupSchemaTableResponse {
		if v == nil {
			return nil
		}
		return v.Tables
	}).(SyncGroupSchemaTableResponseArrayOutput)
}

type SyncGroupSchemaTable struct {
	Columns    []SyncGroupSchemaTableColumn `pulumi:"columns"`
	QuotedName *string                      `pulumi:"quotedName"`
}





type SyncGroupSchemaTableInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableOutput() SyncGroupSchemaTableOutput
	ToSyncGroupSchemaTableOutputWithContext(context.Context) SyncGroupSchemaTableOutput
}

type SyncGroupSchemaTableArgs struct {
	Columns    SyncGroupSchemaTableColumnArrayInput `pulumi:"columns"`
	QuotedName pulumi.StringPtrInput                `pulumi:"quotedName"`
}

func (SyncGroupSchemaTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTable)(nil)).Elem()
}

func (i SyncGroupSchemaTableArgs) ToSyncGroupSchemaTableOutput() SyncGroupSchemaTableOutput {
	return i.ToSyncGroupSchemaTableOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableArgs) ToSyncGroupSchemaTableOutputWithContext(ctx context.Context) SyncGroupSchemaTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableOutput)
}





type SyncGroupSchemaTableArrayInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableArrayOutput() SyncGroupSchemaTableArrayOutput
	ToSyncGroupSchemaTableArrayOutputWithContext(context.Context) SyncGroupSchemaTableArrayOutput
}

type SyncGroupSchemaTableArray []SyncGroupSchemaTableInput

func (SyncGroupSchemaTableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTable)(nil)).Elem()
}

func (i SyncGroupSchemaTableArray) ToSyncGroupSchemaTableArrayOutput() SyncGroupSchemaTableArrayOutput {
	return i.ToSyncGroupSchemaTableArrayOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableArray) ToSyncGroupSchemaTableArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableArrayOutput)
}

type SyncGroupSchemaTableOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTable)(nil)).Elem()
}

func (o SyncGroupSchemaTableOutput) ToSyncGroupSchemaTableOutput() SyncGroupSchemaTableOutput {
	return o
}

func (o SyncGroupSchemaTableOutput) ToSyncGroupSchemaTableOutputWithContext(ctx context.Context) SyncGroupSchemaTableOutput {
	return o
}

func (o SyncGroupSchemaTableOutput) Columns() SyncGroupSchemaTableColumnArrayOutput {
	return o.ApplyT(func(v SyncGroupSchemaTable) []SyncGroupSchemaTableColumn { return v.Columns }).(SyncGroupSchemaTableColumnArrayOutput)
}

func (o SyncGroupSchemaTableOutput) QuotedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTable) *string { return v.QuotedName }).(pulumi.StringPtrOutput)
}

type SyncGroupSchemaTableArrayOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTable)(nil)).Elem()
}

func (o SyncGroupSchemaTableArrayOutput) ToSyncGroupSchemaTableArrayOutput() SyncGroupSchemaTableArrayOutput {
	return o
}

func (o SyncGroupSchemaTableArrayOutput) ToSyncGroupSchemaTableArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableArrayOutput {
	return o
}

func (o SyncGroupSchemaTableArrayOutput) Index(i pulumi.IntInput) SyncGroupSchemaTableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SyncGroupSchemaTable {
		return vs[0].([]SyncGroupSchemaTable)[vs[1].(int)]
	}).(SyncGroupSchemaTableOutput)
}

type SyncGroupSchemaTableColumn struct {
	DataSize   *string `pulumi:"dataSize"`
	DataType   *string `pulumi:"dataType"`
	QuotedName *string `pulumi:"quotedName"`
}





type SyncGroupSchemaTableColumnInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableColumnOutput() SyncGroupSchemaTableColumnOutput
	ToSyncGroupSchemaTableColumnOutputWithContext(context.Context) SyncGroupSchemaTableColumnOutput
}

type SyncGroupSchemaTableColumnArgs struct {
	DataSize   pulumi.StringPtrInput `pulumi:"dataSize"`
	DataType   pulumi.StringPtrInput `pulumi:"dataType"`
	QuotedName pulumi.StringPtrInput `pulumi:"quotedName"`
}

func (SyncGroupSchemaTableColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTableColumn)(nil)).Elem()
}

func (i SyncGroupSchemaTableColumnArgs) ToSyncGroupSchemaTableColumnOutput() SyncGroupSchemaTableColumnOutput {
	return i.ToSyncGroupSchemaTableColumnOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableColumnArgs) ToSyncGroupSchemaTableColumnOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableColumnOutput)
}





type SyncGroupSchemaTableColumnArrayInput interface {
	pulumi.Input

	ToSyncGroupSchemaTableColumnArrayOutput() SyncGroupSchemaTableColumnArrayOutput
	ToSyncGroupSchemaTableColumnArrayOutputWithContext(context.Context) SyncGroupSchemaTableColumnArrayOutput
}

type SyncGroupSchemaTableColumnArray []SyncGroupSchemaTableColumnInput

func (SyncGroupSchemaTableColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTableColumn)(nil)).Elem()
}

func (i SyncGroupSchemaTableColumnArray) ToSyncGroupSchemaTableColumnArrayOutput() SyncGroupSchemaTableColumnArrayOutput {
	return i.ToSyncGroupSchemaTableColumnArrayOutputWithContext(context.Background())
}

func (i SyncGroupSchemaTableColumnArray) ToSyncGroupSchemaTableColumnArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncGroupSchemaTableColumnArrayOutput)
}

type SyncGroupSchemaTableColumnOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTableColumn)(nil)).Elem()
}

func (o SyncGroupSchemaTableColumnOutput) ToSyncGroupSchemaTableColumnOutput() SyncGroupSchemaTableColumnOutput {
	return o
}

func (o SyncGroupSchemaTableColumnOutput) ToSyncGroupSchemaTableColumnOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnOutput {
	return o
}

func (o SyncGroupSchemaTableColumnOutput) DataSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableColumn) *string { return v.DataSize }).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaTableColumnOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableColumn) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaTableColumnOutput) QuotedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableColumn) *string { return v.QuotedName }).(pulumi.StringPtrOutput)
}

type SyncGroupSchemaTableColumnArrayOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTableColumn)(nil)).Elem()
}

func (o SyncGroupSchemaTableColumnArrayOutput) ToSyncGroupSchemaTableColumnArrayOutput() SyncGroupSchemaTableColumnArrayOutput {
	return o
}

func (o SyncGroupSchemaTableColumnArrayOutput) ToSyncGroupSchemaTableColumnArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnArrayOutput {
	return o
}

func (o SyncGroupSchemaTableColumnArrayOutput) Index(i pulumi.IntInput) SyncGroupSchemaTableColumnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SyncGroupSchemaTableColumn {
		return vs[0].([]SyncGroupSchemaTableColumn)[vs[1].(int)]
	}).(SyncGroupSchemaTableColumnOutput)
}

type SyncGroupSchemaTableColumnResponse struct {
	DataSize   *string `pulumi:"dataSize"`
	DataType   *string `pulumi:"dataType"`
	QuotedName *string `pulumi:"quotedName"`
}

type SyncGroupSchemaTableColumnResponseOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableColumnResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTableColumnResponse)(nil)).Elem()
}

func (o SyncGroupSchemaTableColumnResponseOutput) ToSyncGroupSchemaTableColumnResponseOutput() SyncGroupSchemaTableColumnResponseOutput {
	return o
}

func (o SyncGroupSchemaTableColumnResponseOutput) ToSyncGroupSchemaTableColumnResponseOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnResponseOutput {
	return o
}

func (o SyncGroupSchemaTableColumnResponseOutput) DataSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableColumnResponse) *string { return v.DataSize }).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaTableColumnResponseOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableColumnResponse) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o SyncGroupSchemaTableColumnResponseOutput) QuotedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableColumnResponse) *string { return v.QuotedName }).(pulumi.StringPtrOutput)
}

type SyncGroupSchemaTableColumnResponseArrayOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableColumnResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTableColumnResponse)(nil)).Elem()
}

func (o SyncGroupSchemaTableColumnResponseArrayOutput) ToSyncGroupSchemaTableColumnResponseArrayOutput() SyncGroupSchemaTableColumnResponseArrayOutput {
	return o
}

func (o SyncGroupSchemaTableColumnResponseArrayOutput) ToSyncGroupSchemaTableColumnResponseArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableColumnResponseArrayOutput {
	return o
}

func (o SyncGroupSchemaTableColumnResponseArrayOutput) Index(i pulumi.IntInput) SyncGroupSchemaTableColumnResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SyncGroupSchemaTableColumnResponse {
		return vs[0].([]SyncGroupSchemaTableColumnResponse)[vs[1].(int)]
	}).(SyncGroupSchemaTableColumnResponseOutput)
}

type SyncGroupSchemaTableResponse struct {
	Columns    []SyncGroupSchemaTableColumnResponse `pulumi:"columns"`
	QuotedName *string                              `pulumi:"quotedName"`
}

type SyncGroupSchemaTableResponseOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncGroupSchemaTableResponse)(nil)).Elem()
}

func (o SyncGroupSchemaTableResponseOutput) ToSyncGroupSchemaTableResponseOutput() SyncGroupSchemaTableResponseOutput {
	return o
}

func (o SyncGroupSchemaTableResponseOutput) ToSyncGroupSchemaTableResponseOutputWithContext(ctx context.Context) SyncGroupSchemaTableResponseOutput {
	return o
}

func (o SyncGroupSchemaTableResponseOutput) Columns() SyncGroupSchemaTableColumnResponseArrayOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableResponse) []SyncGroupSchemaTableColumnResponse { return v.Columns }).(SyncGroupSchemaTableColumnResponseArrayOutput)
}

func (o SyncGroupSchemaTableResponseOutput) QuotedName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncGroupSchemaTableResponse) *string { return v.QuotedName }).(pulumi.StringPtrOutput)
}

type SyncGroupSchemaTableResponseArrayOutput struct{ *pulumi.OutputState }

func (SyncGroupSchemaTableResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SyncGroupSchemaTableResponse)(nil)).Elem()
}

func (o SyncGroupSchemaTableResponseArrayOutput) ToSyncGroupSchemaTableResponseArrayOutput() SyncGroupSchemaTableResponseArrayOutput {
	return o
}

func (o SyncGroupSchemaTableResponseArrayOutput) ToSyncGroupSchemaTableResponseArrayOutputWithContext(ctx context.Context) SyncGroupSchemaTableResponseArrayOutput {
	return o
}

func (o SyncGroupSchemaTableResponseArrayOutput) Index(i pulumi.IntInput) SyncGroupSchemaTableResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SyncGroupSchemaTableResponse {
		return vs[0].([]SyncGroupSchemaTableResponse)[vs[1].(int)]
	}).(SyncGroupSchemaTableResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityOutput{})
	pulumi.RegisterOutputType(ResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ServerPrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(ServerPrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaPtrOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaResponseOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaResponsePtrOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableArrayOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableColumnOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableColumnArrayOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableColumnResponseOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableColumnResponseArrayOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableResponseOutput{})
	pulumi.RegisterOutputType(SyncGroupSchemaTableResponseArrayOutput{})
}
