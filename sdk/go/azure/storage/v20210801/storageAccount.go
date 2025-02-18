


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StorageAccount struct {
	pulumi.CustomResourceState

	AccessTier                            pulumi.StringOutput                                    `pulumi:"accessTier"`
	AllowBlobPublicAccess                 pulumi.BoolPtrOutput                                   `pulumi:"allowBlobPublicAccess"`
	AllowCrossTenantReplication           pulumi.BoolPtrOutput                                   `pulumi:"allowCrossTenantReplication"`
	AllowSharedKeyAccess                  pulumi.BoolPtrOutput                                   `pulumi:"allowSharedKeyAccess"`
	AllowedCopyScope                      pulumi.StringPtrOutput                                 `pulumi:"allowedCopyScope"`
	AzureFilesIdentityBasedAuthentication AzureFilesIdentityBasedAuthenticationResponsePtrOutput `pulumi:"azureFilesIdentityBasedAuthentication"`
	BlobRestoreStatus                     BlobRestoreStatusResponseOutput                        `pulumi:"blobRestoreStatus"`
	CreationTime                          pulumi.StringOutput                                    `pulumi:"creationTime"`
	CustomDomain                          CustomDomainResponseOutput                             `pulumi:"customDomain"`
	DefaultToOAuthAuthentication          pulumi.BoolPtrOutput                                   `pulumi:"defaultToOAuthAuthentication"`
	EnableHttpsTrafficOnly                pulumi.BoolPtrOutput                                   `pulumi:"enableHttpsTrafficOnly"`
	EnableNfsV3                           pulumi.BoolPtrOutput                                   `pulumi:"enableNfsV3"`
	Encryption                            EncryptionResponseOutput                               `pulumi:"encryption"`
	ExtendedLocation                      ExtendedLocationResponsePtrOutput                      `pulumi:"extendedLocation"`
	FailoverInProgress                    pulumi.BoolOutput                                      `pulumi:"failoverInProgress"`
	GeoReplicationStats                   GeoReplicationStatsResponseOutput                      `pulumi:"geoReplicationStats"`
	Identity                              IdentityResponsePtrOutput                              `pulumi:"identity"`
	ImmutableStorageWithVersioning        ImmutableStorageAccountResponsePtrOutput               `pulumi:"immutableStorageWithVersioning"`
	IsHnsEnabled                          pulumi.BoolPtrOutput                                   `pulumi:"isHnsEnabled"`
	IsLocalUserEnabled                    pulumi.BoolPtrOutput                                   `pulumi:"isLocalUserEnabled"`
	IsSftpEnabled                         pulumi.BoolPtrOutput                                   `pulumi:"isSftpEnabled"`
	KeyCreationTime                       KeyCreationTimeResponseOutput                          `pulumi:"keyCreationTime"`
	KeyPolicy                             KeyPolicyResponseOutput                                `pulumi:"keyPolicy"`
	Kind                                  pulumi.StringOutput                                    `pulumi:"kind"`
	LargeFileSharesState                  pulumi.StringPtrOutput                                 `pulumi:"largeFileSharesState"`
	LastGeoFailoverTime                   pulumi.StringOutput                                    `pulumi:"lastGeoFailoverTime"`
	Location                              pulumi.StringOutput                                    `pulumi:"location"`
	MinimumTlsVersion                     pulumi.StringPtrOutput                                 `pulumi:"minimumTlsVersion"`
	Name                                  pulumi.StringOutput                                    `pulumi:"name"`
	NetworkRuleSet                        NetworkRuleSetResponseOutput                           `pulumi:"networkRuleSet"`
	PrimaryEndpoints                      EndpointsResponseOutput                                `pulumi:"primaryEndpoints"`
	PrimaryLocation                       pulumi.StringOutput                                    `pulumi:"primaryLocation"`
	PrivateEndpointConnections            PrivateEndpointConnectionResponseArrayOutput           `pulumi:"privateEndpointConnections"`
	ProvisioningState                     pulumi.StringOutput                                    `pulumi:"provisioningState"`
	PublicNetworkAccess                   pulumi.StringPtrOutput                                 `pulumi:"publicNetworkAccess"`
	RoutingPreference                     RoutingPreferenceResponsePtrOutput                     `pulumi:"routingPreference"`
	SasPolicy                             SasPolicyResponseOutput                                `pulumi:"sasPolicy"`
	SecondaryEndpoints                    EndpointsResponseOutput                                `pulumi:"secondaryEndpoints"`
	SecondaryLocation                     pulumi.StringOutput                                    `pulumi:"secondaryLocation"`
	Sku                                   SkuResponseOutput                                      `pulumi:"sku"`
	StatusOfPrimary                       pulumi.StringOutput                                    `pulumi:"statusOfPrimary"`
	StatusOfSecondary                     pulumi.StringOutput                                    `pulumi:"statusOfSecondary"`
	Tags                                  pulumi.StringMapOutput                                 `pulumi:"tags"`
	Type                                  pulumi.StringOutput                                    `pulumi:"type"`
}


func NewStorageAccount(ctx *pulumi.Context,
	name string, args *StorageAccountArgs, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.Encryption != nil {
		args.Encryption = args.Encryption.ToEncryptionPtrOutput().ApplyT(func(v *Encryption) *Encryption { return v.Defaults() }).(EncryptionPtrOutput)
	}
	if args.NetworkRuleSet != nil {
		args.NetworkRuleSet = args.NetworkRuleSet.ToNetworkRuleSetPtrOutput().ApplyT(func(v *NetworkRuleSet) *NetworkRuleSet { return v.Defaults() }).(NetworkRuleSetPtrOutput)
	}
	if args.SasPolicy != nil {
		args.SasPolicy = args.SasPolicy.ToSasPolicyPtrOutput().ApplyT(func(v *SasPolicy) *SasPolicy { return v.Defaults() }).(SasPolicyPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20150501preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20150615:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20160101:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20160501:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20161201:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20170601:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20171001:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180201:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180301preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180701:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20181101:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190401:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:StorageAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageAccount
	err := ctx.RegisterResource("azure-native:storage/v20210801:StorageAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAccountState, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	var resource StorageAccount
	err := ctx.ReadResource("azure-native:storage/v20210801:StorageAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type storageAccountState struct {
}

type StorageAccountState struct {
}

func (StorageAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountState)(nil)).Elem()
}

type storageAccountArgs struct {
	AccessTier                            *AccessTier                            `pulumi:"accessTier"`
	AccountName                           *string                                `pulumi:"accountName"`
	AllowBlobPublicAccess                 *bool                                  `pulumi:"allowBlobPublicAccess"`
	AllowCrossTenantReplication           *bool                                  `pulumi:"allowCrossTenantReplication"`
	AllowSharedKeyAccess                  *bool                                  `pulumi:"allowSharedKeyAccess"`
	AllowedCopyScope                      *string                                `pulumi:"allowedCopyScope"`
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthentication `pulumi:"azureFilesIdentityBasedAuthentication"`
	CustomDomain                          *CustomDomain                          `pulumi:"customDomain"`
	DefaultToOAuthAuthentication          *bool                                  `pulumi:"defaultToOAuthAuthentication"`
	EnableHttpsTrafficOnly                *bool                                  `pulumi:"enableHttpsTrafficOnly"`
	EnableNfsV3                           *bool                                  `pulumi:"enableNfsV3"`
	Encryption                            *Encryption                            `pulumi:"encryption"`
	ExtendedLocation                      *ExtendedLocation                      `pulumi:"extendedLocation"`
	Identity                              *Identity                              `pulumi:"identity"`
	ImmutableStorageWithVersioning        *ImmutableStorageAccount               `pulumi:"immutableStorageWithVersioning"`
	IsHnsEnabled                          *bool                                  `pulumi:"isHnsEnabled"`
	IsLocalUserEnabled                    *bool                                  `pulumi:"isLocalUserEnabled"`
	IsSftpEnabled                         *bool                                  `pulumi:"isSftpEnabled"`
	KeyPolicy                             *KeyPolicy                             `pulumi:"keyPolicy"`
	Kind                                  string                                 `pulumi:"kind"`
	LargeFileSharesState                  *string                                `pulumi:"largeFileSharesState"`
	Location                              *string                                `pulumi:"location"`
	MinimumTlsVersion                     *string                                `pulumi:"minimumTlsVersion"`
	NetworkRuleSet                        *NetworkRuleSet                        `pulumi:"networkRuleSet"`
	PublicNetworkAccess                   *string                                `pulumi:"publicNetworkAccess"`
	ResourceGroupName                     string                                 `pulumi:"resourceGroupName"`
	RoutingPreference                     *RoutingPreference                     `pulumi:"routingPreference"`
	SasPolicy                             *SasPolicy                             `pulumi:"sasPolicy"`
	Sku                                   Sku                                    `pulumi:"sku"`
	Tags                                  map[string]string                      `pulumi:"tags"`
}


type StorageAccountArgs struct {
	AccessTier                            AccessTierPtrInput
	AccountName                           pulumi.StringPtrInput
	AllowBlobPublicAccess                 pulumi.BoolPtrInput
	AllowCrossTenantReplication           pulumi.BoolPtrInput
	AllowSharedKeyAccess                  pulumi.BoolPtrInput
	AllowedCopyScope                      pulumi.StringPtrInput
	AzureFilesIdentityBasedAuthentication AzureFilesIdentityBasedAuthenticationPtrInput
	CustomDomain                          CustomDomainPtrInput
	DefaultToOAuthAuthentication          pulumi.BoolPtrInput
	EnableHttpsTrafficOnly                pulumi.BoolPtrInput
	EnableNfsV3                           pulumi.BoolPtrInput
	Encryption                            EncryptionPtrInput
	ExtendedLocation                      ExtendedLocationPtrInput
	Identity                              IdentityPtrInput
	ImmutableStorageWithVersioning        ImmutableStorageAccountPtrInput
	IsHnsEnabled                          pulumi.BoolPtrInput
	IsLocalUserEnabled                    pulumi.BoolPtrInput
	IsSftpEnabled                         pulumi.BoolPtrInput
	KeyPolicy                             KeyPolicyPtrInput
	Kind                                  pulumi.StringInput
	LargeFileSharesState                  pulumi.StringPtrInput
	Location                              pulumi.StringPtrInput
	MinimumTlsVersion                     pulumi.StringPtrInput
	NetworkRuleSet                        NetworkRuleSetPtrInput
	PublicNetworkAccess                   pulumi.StringPtrInput
	ResourceGroupName                     pulumi.StringInput
	RoutingPreference                     RoutingPreferencePtrInput
	SasPolicy                             SasPolicyPtrInput
	Sku                                   SkuInput
	Tags                                  pulumi.StringMapInput
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountArgs)(nil)).Elem()
}

type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput
}

func (*StorageAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccount)(nil)).Elem()
}

func (i *StorageAccount) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i *StorageAccount) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}

type StorageAccountOutput struct{ *pulumi.OutputState }

func (StorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccount)(nil)).Elem()
}

func (o StorageAccountOutput) ToStorageAccountOutput() StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) AccessTier() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.AccessTier }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) AllowBlobPublicAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.AllowBlobPublicAccess }).(pulumi.BoolPtrOutput)
}

func (o StorageAccountOutput) AllowCrossTenantReplication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.AllowCrossTenantReplication }).(pulumi.BoolPtrOutput)
}

func (o StorageAccountOutput) AllowSharedKeyAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.AllowSharedKeyAccess }).(pulumi.BoolPtrOutput)
}

func (o StorageAccountOutput) AllowedCopyScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.AllowedCopyScope }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) AzureFilesIdentityBasedAuthentication() AzureFilesIdentityBasedAuthenticationResponsePtrOutput {
	return o.ApplyT(func(v *StorageAccount) AzureFilesIdentityBasedAuthenticationResponsePtrOutput {
		return v.AzureFilesIdentityBasedAuthentication
	}).(AzureFilesIdentityBasedAuthenticationResponsePtrOutput)
}

func (o StorageAccountOutput) BlobRestoreStatus() BlobRestoreStatusResponseOutput {
	return o.ApplyT(func(v *StorageAccount) BlobRestoreStatusResponseOutput { return v.BlobRestoreStatus }).(BlobRestoreStatusResponseOutput)
}

func (o StorageAccountOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) CustomDomain() CustomDomainResponseOutput {
	return o.ApplyT(func(v *StorageAccount) CustomDomainResponseOutput { return v.CustomDomain }).(CustomDomainResponseOutput)
}

func (o StorageAccountOutput) DefaultToOAuthAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.DefaultToOAuthAuthentication }).(pulumi.BoolPtrOutput)
}

func (o StorageAccountOutput) EnableHttpsTrafficOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.EnableHttpsTrafficOnly }).(pulumi.BoolPtrOutput)
}

func (o StorageAccountOutput) EnableNfsV3() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.EnableNfsV3 }).(pulumi.BoolPtrOutput)
}

func (o StorageAccountOutput) Encryption() EncryptionResponseOutput {
	return o.ApplyT(func(v *StorageAccount) EncryptionResponseOutput { return v.Encryption }).(EncryptionResponseOutput)
}

func (o StorageAccountOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *StorageAccount) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o StorageAccountOutput) FailoverInProgress() pulumi.BoolOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolOutput { return v.FailoverInProgress }).(pulumi.BoolOutput)
}

func (o StorageAccountOutput) GeoReplicationStats() GeoReplicationStatsResponseOutput {
	return o.ApplyT(func(v *StorageAccount) GeoReplicationStatsResponseOutput { return v.GeoReplicationStats }).(GeoReplicationStatsResponseOutput)
}

func (o StorageAccountOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *StorageAccount) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o StorageAccountOutput) ImmutableStorageWithVersioning() ImmutableStorageAccountResponsePtrOutput {
	return o.ApplyT(func(v *StorageAccount) ImmutableStorageAccountResponsePtrOutput {
		return v.ImmutableStorageWithVersioning
	}).(ImmutableStorageAccountResponsePtrOutput)
}

func (o StorageAccountOutput) IsHnsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.IsHnsEnabled }).(pulumi.BoolPtrOutput)
}

func (o StorageAccountOutput) IsLocalUserEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.IsLocalUserEnabled }).(pulumi.BoolPtrOutput)
}

func (o StorageAccountOutput) IsSftpEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.IsSftpEnabled }).(pulumi.BoolPtrOutput)
}

func (o StorageAccountOutput) KeyCreationTime() KeyCreationTimeResponseOutput {
	return o.ApplyT(func(v *StorageAccount) KeyCreationTimeResponseOutput { return v.KeyCreationTime }).(KeyCreationTimeResponseOutput)
}

func (o StorageAccountOutput) KeyPolicy() KeyPolicyResponseOutput {
	return o.ApplyT(func(v *StorageAccount) KeyPolicyResponseOutput { return v.KeyPolicy }).(KeyPolicyResponseOutput)
}

func (o StorageAccountOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) LargeFileSharesState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.LargeFileSharesState }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) LastGeoFailoverTime() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.LastGeoFailoverTime }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) MinimumTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.MinimumTlsVersion }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) NetworkRuleSet() NetworkRuleSetResponseOutput {
	return o.ApplyT(func(v *StorageAccount) NetworkRuleSetResponseOutput { return v.NetworkRuleSet }).(NetworkRuleSetResponseOutput)
}

func (o StorageAccountOutput) PrimaryEndpoints() EndpointsResponseOutput {
	return o.ApplyT(func(v *StorageAccount) EndpointsResponseOutput { return v.PrimaryEndpoints }).(EndpointsResponseOutput)
}

func (o StorageAccountOutput) PrimaryLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.PrimaryLocation }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *StorageAccount) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o StorageAccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) RoutingPreference() RoutingPreferenceResponsePtrOutput {
	return o.ApplyT(func(v *StorageAccount) RoutingPreferenceResponsePtrOutput { return v.RoutingPreference }).(RoutingPreferenceResponsePtrOutput)
}

func (o StorageAccountOutput) SasPolicy() SasPolicyResponseOutput {
	return o.ApplyT(func(v *StorageAccount) SasPolicyResponseOutput { return v.SasPolicy }).(SasPolicyResponseOutput)
}

func (o StorageAccountOutput) SecondaryEndpoints() EndpointsResponseOutput {
	return o.ApplyT(func(v *StorageAccount) EndpointsResponseOutput { return v.SecondaryEndpoints }).(EndpointsResponseOutput)
}

func (o StorageAccountOutput) SecondaryLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.SecondaryLocation }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *StorageAccount) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

func (o StorageAccountOutput) StatusOfPrimary() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.StatusOfPrimary }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) StatusOfSecondary() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.StatusOfSecondary }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o StorageAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageAccountOutput{})
}
