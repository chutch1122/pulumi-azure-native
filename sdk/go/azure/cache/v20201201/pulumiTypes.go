


package v20201201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	Type                              string                                    `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type RedisAccessKeysResponse struct {
	PrimaryKey   string `pulumi:"primaryKey"`
	SecondaryKey string `pulumi:"secondaryKey"`
}

type RedisAccessKeysResponseOutput struct{ *pulumi.OutputState }

func (RedisAccessKeysResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisAccessKeysResponse)(nil)).Elem()
}

func (o RedisAccessKeysResponseOutput) ToRedisAccessKeysResponseOutput() RedisAccessKeysResponseOutput {
	return o
}

func (o RedisAccessKeysResponseOutput) ToRedisAccessKeysResponseOutputWithContext(ctx context.Context) RedisAccessKeysResponseOutput {
	return o
}

func (o RedisAccessKeysResponseOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v RedisAccessKeysResponse) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o RedisAccessKeysResponseOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v RedisAccessKeysResponse) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

type RedisCommonPropertiesRedisConfiguration struct {
	AofBackupEnabled               *string `pulumi:"aofBackupEnabled"`
	AofStorageConnectionString0    *string `pulumi:"aofStorageConnectionString0"`
	AofStorageConnectionString1    *string `pulumi:"aofStorageConnectionString1"`
	Authnotrequired                *string `pulumi:"authnotrequired"`
	MaxfragmentationmemoryReserved *string `pulumi:"maxfragmentationmemoryReserved"`
	MaxmemoryDelta                 *string `pulumi:"maxmemoryDelta"`
	MaxmemoryPolicy                *string `pulumi:"maxmemoryPolicy"`
	MaxmemoryReserved              *string `pulumi:"maxmemoryReserved"`
	RdbBackupEnabled               *string `pulumi:"rdbBackupEnabled"`
	RdbBackupFrequency             *string `pulumi:"rdbBackupFrequency"`
	RdbBackupMaxSnapshotCount      *string `pulumi:"rdbBackupMaxSnapshotCount"`
	RdbStorageConnectionString     *string `pulumi:"rdbStorageConnectionString"`
}





type RedisCommonPropertiesRedisConfigurationInput interface {
	pulumi.Input

	ToRedisCommonPropertiesRedisConfigurationOutput() RedisCommonPropertiesRedisConfigurationOutput
	ToRedisCommonPropertiesRedisConfigurationOutputWithContext(context.Context) RedisCommonPropertiesRedisConfigurationOutput
}

type RedisCommonPropertiesRedisConfigurationArgs struct {
	AofBackupEnabled               pulumi.StringPtrInput `pulumi:"aofBackupEnabled"`
	AofStorageConnectionString0    pulumi.StringPtrInput `pulumi:"aofStorageConnectionString0"`
	AofStorageConnectionString1    pulumi.StringPtrInput `pulumi:"aofStorageConnectionString1"`
	Authnotrequired                pulumi.StringPtrInput `pulumi:"authnotrequired"`
	MaxfragmentationmemoryReserved pulumi.StringPtrInput `pulumi:"maxfragmentationmemoryReserved"`
	MaxmemoryDelta                 pulumi.StringPtrInput `pulumi:"maxmemoryDelta"`
	MaxmemoryPolicy                pulumi.StringPtrInput `pulumi:"maxmemoryPolicy"`
	MaxmemoryReserved              pulumi.StringPtrInput `pulumi:"maxmemoryReserved"`
	RdbBackupEnabled               pulumi.StringPtrInput `pulumi:"rdbBackupEnabled"`
	RdbBackupFrequency             pulumi.StringPtrInput `pulumi:"rdbBackupFrequency"`
	RdbBackupMaxSnapshotCount      pulumi.StringPtrInput `pulumi:"rdbBackupMaxSnapshotCount"`
	RdbStorageConnectionString     pulumi.StringPtrInput `pulumi:"rdbStorageConnectionString"`
}

func (RedisCommonPropertiesRedisConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisCommonPropertiesRedisConfiguration)(nil)).Elem()
}

func (i RedisCommonPropertiesRedisConfigurationArgs) ToRedisCommonPropertiesRedisConfigurationOutput() RedisCommonPropertiesRedisConfigurationOutput {
	return i.ToRedisCommonPropertiesRedisConfigurationOutputWithContext(context.Background())
}

func (i RedisCommonPropertiesRedisConfigurationArgs) ToRedisCommonPropertiesRedisConfigurationOutputWithContext(ctx context.Context) RedisCommonPropertiesRedisConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisCommonPropertiesRedisConfigurationOutput)
}

func (i RedisCommonPropertiesRedisConfigurationArgs) ToRedisCommonPropertiesRedisConfigurationPtrOutput() RedisCommonPropertiesRedisConfigurationPtrOutput {
	return i.ToRedisCommonPropertiesRedisConfigurationPtrOutputWithContext(context.Background())
}

func (i RedisCommonPropertiesRedisConfigurationArgs) ToRedisCommonPropertiesRedisConfigurationPtrOutputWithContext(ctx context.Context) RedisCommonPropertiesRedisConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisCommonPropertiesRedisConfigurationOutput).ToRedisCommonPropertiesRedisConfigurationPtrOutputWithContext(ctx)
}









type RedisCommonPropertiesRedisConfigurationPtrInput interface {
	pulumi.Input

	ToRedisCommonPropertiesRedisConfigurationPtrOutput() RedisCommonPropertiesRedisConfigurationPtrOutput
	ToRedisCommonPropertiesRedisConfigurationPtrOutputWithContext(context.Context) RedisCommonPropertiesRedisConfigurationPtrOutput
}

type redisCommonPropertiesRedisConfigurationPtrType RedisCommonPropertiesRedisConfigurationArgs

func RedisCommonPropertiesRedisConfigurationPtr(v *RedisCommonPropertiesRedisConfigurationArgs) RedisCommonPropertiesRedisConfigurationPtrInput {
	return (*redisCommonPropertiesRedisConfigurationPtrType)(v)
}

func (*redisCommonPropertiesRedisConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisCommonPropertiesRedisConfiguration)(nil)).Elem()
}

func (i *redisCommonPropertiesRedisConfigurationPtrType) ToRedisCommonPropertiesRedisConfigurationPtrOutput() RedisCommonPropertiesRedisConfigurationPtrOutput {
	return i.ToRedisCommonPropertiesRedisConfigurationPtrOutputWithContext(context.Background())
}

func (i *redisCommonPropertiesRedisConfigurationPtrType) ToRedisCommonPropertiesRedisConfigurationPtrOutputWithContext(ctx context.Context) RedisCommonPropertiesRedisConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisCommonPropertiesRedisConfigurationPtrOutput)
}

type RedisCommonPropertiesRedisConfigurationOutput struct{ *pulumi.OutputState }

func (RedisCommonPropertiesRedisConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisCommonPropertiesRedisConfiguration)(nil)).Elem()
}

func (o RedisCommonPropertiesRedisConfigurationOutput) ToRedisCommonPropertiesRedisConfigurationOutput() RedisCommonPropertiesRedisConfigurationOutput {
	return o
}

func (o RedisCommonPropertiesRedisConfigurationOutput) ToRedisCommonPropertiesRedisConfigurationOutputWithContext(ctx context.Context) RedisCommonPropertiesRedisConfigurationOutput {
	return o
}

func (o RedisCommonPropertiesRedisConfigurationOutput) ToRedisCommonPropertiesRedisConfigurationPtrOutput() RedisCommonPropertiesRedisConfigurationPtrOutput {
	return o.ToRedisCommonPropertiesRedisConfigurationPtrOutputWithContext(context.Background())
}

func (o RedisCommonPropertiesRedisConfigurationOutput) ToRedisCommonPropertiesRedisConfigurationPtrOutputWithContext(ctx context.Context) RedisCommonPropertiesRedisConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RedisCommonPropertiesRedisConfiguration) *RedisCommonPropertiesRedisConfiguration {
		return &v
	}).(RedisCommonPropertiesRedisConfigurationPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) AofBackupEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.AofBackupEnabled }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) AofStorageConnectionString0() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.AofStorageConnectionString0 }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) AofStorageConnectionString1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.AofStorageConnectionString1 }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) Authnotrequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.Authnotrequired }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) MaxfragmentationmemoryReserved() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.MaxfragmentationmemoryReserved }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) MaxmemoryDelta() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.MaxmemoryDelta }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) MaxmemoryPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.MaxmemoryPolicy }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) MaxmemoryReserved() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.MaxmemoryReserved }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) RdbBackupEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.RdbBackupEnabled }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) RdbBackupFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.RdbBackupFrequency }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) RdbBackupMaxSnapshotCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.RdbBackupMaxSnapshotCount }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) RdbStorageConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.RdbStorageConnectionString }).(pulumi.StringPtrOutput)
}

type RedisCommonPropertiesRedisConfigurationPtrOutput struct{ *pulumi.OutputState }

func (RedisCommonPropertiesRedisConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisCommonPropertiesRedisConfiguration)(nil)).Elem()
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) ToRedisCommonPropertiesRedisConfigurationPtrOutput() RedisCommonPropertiesRedisConfigurationPtrOutput {
	return o
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) ToRedisCommonPropertiesRedisConfigurationPtrOutputWithContext(ctx context.Context) RedisCommonPropertiesRedisConfigurationPtrOutput {
	return o
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) Elem() RedisCommonPropertiesRedisConfigurationOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) RedisCommonPropertiesRedisConfiguration {
		if v != nil {
			return *v
		}
		var ret RedisCommonPropertiesRedisConfiguration
		return ret
	}).(RedisCommonPropertiesRedisConfigurationOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) AofBackupEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AofBackupEnabled
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) AofStorageConnectionString0() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AofStorageConnectionString0
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) AofStorageConnectionString1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AofStorageConnectionString1
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) Authnotrequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Authnotrequired
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) MaxfragmentationmemoryReserved() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.MaxfragmentationmemoryReserved
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) MaxmemoryDelta() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.MaxmemoryDelta
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) MaxmemoryPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.MaxmemoryPolicy
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) MaxmemoryReserved() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.MaxmemoryReserved
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) RdbBackupEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RdbBackupEnabled
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) RdbBackupFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RdbBackupFrequency
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) RdbBackupMaxSnapshotCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RdbBackupMaxSnapshotCount
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) RdbStorageConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RdbStorageConnectionString
	}).(pulumi.StringPtrOutput)
}

type RedisCommonPropertiesResponseRedisConfiguration struct {
	AofBackupEnabled               *string `pulumi:"aofBackupEnabled"`
	AofStorageConnectionString0    *string `pulumi:"aofStorageConnectionString0"`
	AofStorageConnectionString1    *string `pulumi:"aofStorageConnectionString1"`
	Authnotrequired                *string `pulumi:"authnotrequired"`
	Maxclients                     string  `pulumi:"maxclients"`
	MaxfragmentationmemoryReserved *string `pulumi:"maxfragmentationmemoryReserved"`
	MaxmemoryDelta                 *string `pulumi:"maxmemoryDelta"`
	MaxmemoryPolicy                *string `pulumi:"maxmemoryPolicy"`
	MaxmemoryReserved              *string `pulumi:"maxmemoryReserved"`
	RdbBackupEnabled               *string `pulumi:"rdbBackupEnabled"`
	RdbBackupFrequency             *string `pulumi:"rdbBackupFrequency"`
	RdbBackupMaxSnapshotCount      *string `pulumi:"rdbBackupMaxSnapshotCount"`
	RdbStorageConnectionString     *string `pulumi:"rdbStorageConnectionString"`
	ZonalConfiguration             string  `pulumi:"zonalConfiguration"`
}

type RedisCommonPropertiesResponseRedisConfigurationOutput struct{ *pulumi.OutputState }

func (RedisCommonPropertiesResponseRedisConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisCommonPropertiesResponseRedisConfiguration)(nil)).Elem()
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) ToRedisCommonPropertiesResponseRedisConfigurationOutput() RedisCommonPropertiesResponseRedisConfigurationOutput {
	return o
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) ToRedisCommonPropertiesResponseRedisConfigurationOutputWithContext(ctx context.Context) RedisCommonPropertiesResponseRedisConfigurationOutput {
	return o
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) AofBackupEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.AofBackupEnabled }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) AofStorageConnectionString0() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.AofStorageConnectionString0 }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) AofStorageConnectionString1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.AofStorageConnectionString1 }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) Authnotrequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.Authnotrequired }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) Maxclients() pulumi.StringOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) string { return v.Maxclients }).(pulumi.StringOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) MaxfragmentationmemoryReserved() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string {
		return v.MaxfragmentationmemoryReserved
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) MaxmemoryDelta() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.MaxmemoryDelta }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) MaxmemoryPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.MaxmemoryPolicy }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) MaxmemoryReserved() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.MaxmemoryReserved }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) RdbBackupEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.RdbBackupEnabled }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) RdbBackupFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.RdbBackupFrequency }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) RdbBackupMaxSnapshotCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.RdbBackupMaxSnapshotCount }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) RdbStorageConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.RdbStorageConnectionString }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) ZonalConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) string { return v.ZonalConfiguration }).(pulumi.StringOutput)
}

type RedisCommonPropertiesResponseRedisConfigurationPtrOutput struct{ *pulumi.OutputState }

func (RedisCommonPropertiesResponseRedisConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisCommonPropertiesResponseRedisConfiguration)(nil)).Elem()
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) ToRedisCommonPropertiesResponseRedisConfigurationPtrOutput() RedisCommonPropertiesResponseRedisConfigurationPtrOutput {
	return o
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) ToRedisCommonPropertiesResponseRedisConfigurationPtrOutputWithContext(ctx context.Context) RedisCommonPropertiesResponseRedisConfigurationPtrOutput {
	return o
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) Elem() RedisCommonPropertiesResponseRedisConfigurationOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) RedisCommonPropertiesResponseRedisConfiguration {
		if v != nil {
			return *v
		}
		var ret RedisCommonPropertiesResponseRedisConfiguration
		return ret
	}).(RedisCommonPropertiesResponseRedisConfigurationOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) AofBackupEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AofBackupEnabled
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) AofStorageConnectionString0() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AofStorageConnectionString0
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) AofStorageConnectionString1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AofStorageConnectionString1
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) Authnotrequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Authnotrequired
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) Maxclients() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.Maxclients
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) MaxfragmentationmemoryReserved() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.MaxfragmentationmemoryReserved
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) MaxmemoryDelta() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.MaxmemoryDelta
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) MaxmemoryPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.MaxmemoryPolicy
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) MaxmemoryReserved() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.MaxmemoryReserved
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) RdbBackupEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RdbBackupEnabled
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) RdbBackupFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RdbBackupFrequency
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) RdbBackupMaxSnapshotCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RdbBackupMaxSnapshotCount
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) RdbStorageConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RdbStorageConnectionString
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) ZonalConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.ZonalConfiguration
	}).(pulumi.StringPtrOutput)
}

type RedisInstanceDetailsResponse struct {
	IsMaster   bool   `pulumi:"isMaster"`
	IsPrimary  bool   `pulumi:"isPrimary"`
	NonSslPort int    `pulumi:"nonSslPort"`
	ShardId    int    `pulumi:"shardId"`
	SslPort    int    `pulumi:"sslPort"`
	Zone       string `pulumi:"zone"`
}

type RedisInstanceDetailsResponseOutput struct{ *pulumi.OutputState }

func (RedisInstanceDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisInstanceDetailsResponse)(nil)).Elem()
}

func (o RedisInstanceDetailsResponseOutput) ToRedisInstanceDetailsResponseOutput() RedisInstanceDetailsResponseOutput {
	return o
}

func (o RedisInstanceDetailsResponseOutput) ToRedisInstanceDetailsResponseOutputWithContext(ctx context.Context) RedisInstanceDetailsResponseOutput {
	return o
}

func (o RedisInstanceDetailsResponseOutput) IsMaster() pulumi.BoolOutput {
	return o.ApplyT(func(v RedisInstanceDetailsResponse) bool { return v.IsMaster }).(pulumi.BoolOutput)
}

func (o RedisInstanceDetailsResponseOutput) IsPrimary() pulumi.BoolOutput {
	return o.ApplyT(func(v RedisInstanceDetailsResponse) bool { return v.IsPrimary }).(pulumi.BoolOutput)
}

func (o RedisInstanceDetailsResponseOutput) NonSslPort() pulumi.IntOutput {
	return o.ApplyT(func(v RedisInstanceDetailsResponse) int { return v.NonSslPort }).(pulumi.IntOutput)
}

func (o RedisInstanceDetailsResponseOutput) ShardId() pulumi.IntOutput {
	return o.ApplyT(func(v RedisInstanceDetailsResponse) int { return v.ShardId }).(pulumi.IntOutput)
}

func (o RedisInstanceDetailsResponseOutput) SslPort() pulumi.IntOutput {
	return o.ApplyT(func(v RedisInstanceDetailsResponse) int { return v.SslPort }).(pulumi.IntOutput)
}

func (o RedisInstanceDetailsResponseOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v RedisInstanceDetailsResponse) string { return v.Zone }).(pulumi.StringOutput)
}

type RedisInstanceDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (RedisInstanceDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RedisInstanceDetailsResponse)(nil)).Elem()
}

func (o RedisInstanceDetailsResponseArrayOutput) ToRedisInstanceDetailsResponseArrayOutput() RedisInstanceDetailsResponseArrayOutput {
	return o
}

func (o RedisInstanceDetailsResponseArrayOutput) ToRedisInstanceDetailsResponseArrayOutputWithContext(ctx context.Context) RedisInstanceDetailsResponseArrayOutput {
	return o
}

func (o RedisInstanceDetailsResponseArrayOutput) Index(i pulumi.IntInput) RedisInstanceDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RedisInstanceDetailsResponse {
		return vs[0].([]RedisInstanceDetailsResponse)[vs[1].(int)]
	}).(RedisInstanceDetailsResponseOutput)
}

type RedisLinkedServerResponse struct {
	Id string `pulumi:"id"`
}

type RedisLinkedServerResponseOutput struct{ *pulumi.OutputState }

func (RedisLinkedServerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisLinkedServerResponse)(nil)).Elem()
}

func (o RedisLinkedServerResponseOutput) ToRedisLinkedServerResponseOutput() RedisLinkedServerResponseOutput {
	return o
}

func (o RedisLinkedServerResponseOutput) ToRedisLinkedServerResponseOutputWithContext(ctx context.Context) RedisLinkedServerResponseOutput {
	return o
}

func (o RedisLinkedServerResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RedisLinkedServerResponse) string { return v.Id }).(pulumi.StringOutput)
}

type RedisLinkedServerResponseArrayOutput struct{ *pulumi.OutputState }

func (RedisLinkedServerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RedisLinkedServerResponse)(nil)).Elem()
}

func (o RedisLinkedServerResponseArrayOutput) ToRedisLinkedServerResponseArrayOutput() RedisLinkedServerResponseArrayOutput {
	return o
}

func (o RedisLinkedServerResponseArrayOutput) ToRedisLinkedServerResponseArrayOutputWithContext(ctx context.Context) RedisLinkedServerResponseArrayOutput {
	return o
}

func (o RedisLinkedServerResponseArrayOutput) Index(i pulumi.IntInput) RedisLinkedServerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RedisLinkedServerResponse {
		return vs[0].([]RedisLinkedServerResponse)[vs[1].(int)]
	}).(RedisLinkedServerResponseOutput)
}

type ScheduleEntry struct {
	DayOfWeek         DayOfWeek `pulumi:"dayOfWeek"`
	MaintenanceWindow *string   `pulumi:"maintenanceWindow"`
	StartHourUtc      int       `pulumi:"startHourUtc"`
}


func (val *ScheduleEntry) Defaults() *ScheduleEntry {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaintenanceWindow) {
		maintenanceWindow_ := "PT5H"
		tmp.MaintenanceWindow = &maintenanceWindow_
	}
	return &tmp
}





type ScheduleEntryInput interface {
	pulumi.Input

	ToScheduleEntryOutput() ScheduleEntryOutput
	ToScheduleEntryOutputWithContext(context.Context) ScheduleEntryOutput
}

type ScheduleEntryArgs struct {
	DayOfWeek         DayOfWeekInput        `pulumi:"dayOfWeek"`
	MaintenanceWindow pulumi.StringPtrInput `pulumi:"maintenanceWindow"`
	StartHourUtc      pulumi.IntInput       `pulumi:"startHourUtc"`
}


func (val *ScheduleEntryArgs) Defaults() *ScheduleEntryArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaintenanceWindow) {
		tmp.MaintenanceWindow = pulumi.StringPtr("PT5H")
	}
	return &tmp
}
func (ScheduleEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleEntry)(nil)).Elem()
}

func (i ScheduleEntryArgs) ToScheduleEntryOutput() ScheduleEntryOutput {
	return i.ToScheduleEntryOutputWithContext(context.Background())
}

func (i ScheduleEntryArgs) ToScheduleEntryOutputWithContext(ctx context.Context) ScheduleEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleEntryOutput)
}





type ScheduleEntryArrayInput interface {
	pulumi.Input

	ToScheduleEntryArrayOutput() ScheduleEntryArrayOutput
	ToScheduleEntryArrayOutputWithContext(context.Context) ScheduleEntryArrayOutput
}

type ScheduleEntryArray []ScheduleEntryInput

func (ScheduleEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleEntry)(nil)).Elem()
}

func (i ScheduleEntryArray) ToScheduleEntryArrayOutput() ScheduleEntryArrayOutput {
	return i.ToScheduleEntryArrayOutputWithContext(context.Background())
}

func (i ScheduleEntryArray) ToScheduleEntryArrayOutputWithContext(ctx context.Context) ScheduleEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleEntryArrayOutput)
}

type ScheduleEntryOutput struct{ *pulumi.OutputState }

func (ScheduleEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleEntry)(nil)).Elem()
}

func (o ScheduleEntryOutput) ToScheduleEntryOutput() ScheduleEntryOutput {
	return o
}

func (o ScheduleEntryOutput) ToScheduleEntryOutputWithContext(ctx context.Context) ScheduleEntryOutput {
	return o
}

func (o ScheduleEntryOutput) DayOfWeek() DayOfWeekOutput {
	return o.ApplyT(func(v ScheduleEntry) DayOfWeek { return v.DayOfWeek }).(DayOfWeekOutput)
}

func (o ScheduleEntryOutput) MaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleEntry) *string { return v.MaintenanceWindow }).(pulumi.StringPtrOutput)
}

func (o ScheduleEntryOutput) StartHourUtc() pulumi.IntOutput {
	return o.ApplyT(func(v ScheduleEntry) int { return v.StartHourUtc }).(pulumi.IntOutput)
}

type ScheduleEntryArrayOutput struct{ *pulumi.OutputState }

func (ScheduleEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleEntry)(nil)).Elem()
}

func (o ScheduleEntryArrayOutput) ToScheduleEntryArrayOutput() ScheduleEntryArrayOutput {
	return o
}

func (o ScheduleEntryArrayOutput) ToScheduleEntryArrayOutputWithContext(ctx context.Context) ScheduleEntryArrayOutput {
	return o
}

func (o ScheduleEntryArrayOutput) Index(i pulumi.IntInput) ScheduleEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScheduleEntry {
		return vs[0].([]ScheduleEntry)[vs[1].(int)]
	}).(ScheduleEntryOutput)
}

type ScheduleEntryResponse struct {
	DayOfWeek         string  `pulumi:"dayOfWeek"`
	MaintenanceWindow *string `pulumi:"maintenanceWindow"`
	StartHourUtc      int     `pulumi:"startHourUtc"`
}


func (val *ScheduleEntryResponse) Defaults() *ScheduleEntryResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaintenanceWindow) {
		maintenanceWindow_ := "PT5H"
		tmp.MaintenanceWindow = &maintenanceWindow_
	}
	return &tmp
}

type ScheduleEntryResponseOutput struct{ *pulumi.OutputState }

func (ScheduleEntryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleEntryResponse)(nil)).Elem()
}

func (o ScheduleEntryResponseOutput) ToScheduleEntryResponseOutput() ScheduleEntryResponseOutput {
	return o
}

func (o ScheduleEntryResponseOutput) ToScheduleEntryResponseOutputWithContext(ctx context.Context) ScheduleEntryResponseOutput {
	return o
}

func (o ScheduleEntryResponseOutput) DayOfWeek() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleEntryResponse) string { return v.DayOfWeek }).(pulumi.StringOutput)
}

func (o ScheduleEntryResponseOutput) MaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleEntryResponse) *string { return v.MaintenanceWindow }).(pulumi.StringPtrOutput)
}

func (o ScheduleEntryResponseOutput) StartHourUtc() pulumi.IntOutput {
	return o.ApplyT(func(v ScheduleEntryResponse) int { return v.StartHourUtc }).(pulumi.IntOutput)
}

type ScheduleEntryResponseArrayOutput struct{ *pulumi.OutputState }

func (ScheduleEntryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleEntryResponse)(nil)).Elem()
}

func (o ScheduleEntryResponseArrayOutput) ToScheduleEntryResponseArrayOutput() ScheduleEntryResponseArrayOutput {
	return o
}

func (o ScheduleEntryResponseArrayOutput) ToScheduleEntryResponseArrayOutputWithContext(ctx context.Context) ScheduleEntryResponseArrayOutput {
	return o
}

func (o ScheduleEntryResponseArrayOutput) Index(i pulumi.IntInput) ScheduleEntryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScheduleEntryResponse {
		return vs[0].([]ScheduleEntryResponse)[vs[1].(int)]
	}).(ScheduleEntryResponseOutput)
}

type Sku struct {
	Capacity int    `pulumi:"capacity"`
	Family   string `pulumi:"family"`
	Name     string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntInput    `pulumi:"capacity"`
	Family   pulumi.StringInput `pulumi:"family"`
	Name     pulumi.StringInput `pulumi:"name"`
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

func (o SkuOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v Sku) int { return v.Capacity }).(pulumi.IntOutput)
}

func (o SkuOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Family }).(pulumi.StringOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponse struct {
	Capacity int    `pulumi:"capacity"`
	Family   string `pulumi:"family"`
	Name     string `pulumi:"name"`
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

func (o SkuResponseOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v SkuResponse) int { return v.Capacity }).(pulumi.IntOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Family }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(RedisAccessKeysResponseOutput{})
	pulumi.RegisterOutputType(RedisCommonPropertiesRedisConfigurationOutput{})
	pulumi.RegisterOutputType(RedisCommonPropertiesRedisConfigurationPtrOutput{})
	pulumi.RegisterOutputType(RedisCommonPropertiesResponseRedisConfigurationOutput{})
	pulumi.RegisterOutputType(RedisCommonPropertiesResponseRedisConfigurationPtrOutput{})
	pulumi.RegisterOutputType(RedisInstanceDetailsResponseOutput{})
	pulumi.RegisterOutputType(RedisInstanceDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(RedisLinkedServerResponseOutput{})
	pulumi.RegisterOutputType(RedisLinkedServerResponseArrayOutput{})
	pulumi.RegisterOutputType(ScheduleEntryOutput{})
	pulumi.RegisterOutputType(ScheduleEntryArrayOutput{})
	pulumi.RegisterOutputType(ScheduleEntryResponseOutput{})
	pulumi.RegisterOutputType(ScheduleEntryResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
}
