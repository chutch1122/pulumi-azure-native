


package v20180701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type StorageAccount struct {
	pulumi.CustomResourceState

	AccessTier                     pulumi.StringOutput               `pulumi:"accessTier"`
	CreationTime                   pulumi.StringOutput               `pulumi:"creationTime"`
	CustomDomain                   CustomDomainResponseOutput        `pulumi:"customDomain"`
	EnableAzureFilesAadIntegration pulumi.BoolPtrOutput              `pulumi:"enableAzureFilesAadIntegration"`
	EnableHttpsTrafficOnly         pulumi.BoolPtrOutput              `pulumi:"enableHttpsTrafficOnly"`
	Encryption                     EncryptionResponseOutput          `pulumi:"encryption"`
	FailoverInProgress             pulumi.BoolOutput                 `pulumi:"failoverInProgress"`
	GeoReplicationStats            GeoReplicationStatsResponseOutput `pulumi:"geoReplicationStats"`
	Identity                       IdentityResponsePtrOutput         `pulumi:"identity"`
	IsHnsEnabled                   pulumi.BoolPtrOutput              `pulumi:"isHnsEnabled"`
	Kind                           pulumi.StringOutput               `pulumi:"kind"`
	LastGeoFailoverTime            pulumi.StringOutput               `pulumi:"lastGeoFailoverTime"`
	Location                       pulumi.StringOutput               `pulumi:"location"`
	Name                           pulumi.StringOutput               `pulumi:"name"`
	NetworkRuleSet                 NetworkRuleSetResponseOutput      `pulumi:"networkRuleSet"`
	PrimaryEndpoints               EndpointsResponseOutput           `pulumi:"primaryEndpoints"`
	PrimaryLocation                pulumi.StringOutput               `pulumi:"primaryLocation"`
	ProvisioningState              pulumi.StringOutput               `pulumi:"provisioningState"`
	SecondaryEndpoints             EndpointsResponseOutput           `pulumi:"secondaryEndpoints"`
	SecondaryLocation              pulumi.StringOutput               `pulumi:"secondaryLocation"`
	Sku                            SkuResponseOutput                 `pulumi:"sku"`
	StatusOfPrimary                pulumi.StringOutput               `pulumi:"statusOfPrimary"`
	StatusOfSecondary              pulumi.StringOutput               `pulumi:"statusOfSecondary"`
	Tags                           pulumi.StringMapOutput            `pulumi:"tags"`
	Type                           pulumi.StringOutput               `pulumi:"type"`
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
			Type: pulumi.String("azure-native:storage/v20210801:StorageAccount"),
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
	err := ctx.RegisterResource("azure-native:storage/v20180701:StorageAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAccountState, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	var resource StorageAccount
	err := ctx.ReadResource("azure-native:storage/v20180701:StorageAccount", name, id, state, &resource, opts...)
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
	AccessTier                     *AccessTier       `pulumi:"accessTier"`
	AccountName                    *string           `pulumi:"accountName"`
	CustomDomain                   *CustomDomain     `pulumi:"customDomain"`
	EnableAzureFilesAadIntegration *bool             `pulumi:"enableAzureFilesAadIntegration"`
	EnableHttpsTrafficOnly         *bool             `pulumi:"enableHttpsTrafficOnly"`
	Encryption                     *Encryption       `pulumi:"encryption"`
	Identity                       *Identity         `pulumi:"identity"`
	IsHnsEnabled                   *bool             `pulumi:"isHnsEnabled"`
	Kind                           Kind              `pulumi:"kind"`
	Location                       *string           `pulumi:"location"`
	NetworkRuleSet                 *NetworkRuleSet   `pulumi:"networkRuleSet"`
	ResourceGroupName              string            `pulumi:"resourceGroupName"`
	Sku                            Sku               `pulumi:"sku"`
	Tags                           map[string]string `pulumi:"tags"`
}


type StorageAccountArgs struct {
	AccessTier                     AccessTierPtrInput
	AccountName                    pulumi.StringPtrInput
	CustomDomain                   CustomDomainPtrInput
	EnableAzureFilesAadIntegration pulumi.BoolPtrInput
	EnableHttpsTrafficOnly         pulumi.BoolPtrInput
	Encryption                     EncryptionPtrInput
	Identity                       IdentityPtrInput
	IsHnsEnabled                   pulumi.BoolPtrInput
	Kind                           KindInput
	Location                       pulumi.StringPtrInput
	NetworkRuleSet                 NetworkRuleSetPtrInput
	ResourceGroupName              pulumi.StringInput
	Sku                            SkuInput
	Tags                           pulumi.StringMapInput
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

func (o StorageAccountOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) CustomDomain() CustomDomainResponseOutput {
	return o.ApplyT(func(v *StorageAccount) CustomDomainResponseOutput { return v.CustomDomain }).(CustomDomainResponseOutput)
}

func (o StorageAccountOutput) EnableAzureFilesAadIntegration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.EnableAzureFilesAadIntegration }).(pulumi.BoolPtrOutput)
}

func (o StorageAccountOutput) EnableHttpsTrafficOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.EnableHttpsTrafficOnly }).(pulumi.BoolPtrOutput)
}

func (o StorageAccountOutput) Encryption() EncryptionResponseOutput {
	return o.ApplyT(func(v *StorageAccount) EncryptionResponseOutput { return v.Encryption }).(EncryptionResponseOutput)
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

func (o StorageAccountOutput) IsHnsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.BoolPtrOutput { return v.IsHnsEnabled }).(pulumi.BoolPtrOutput)
}

func (o StorageAccountOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) LastGeoFailoverTime() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.LastGeoFailoverTime }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
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

func (o StorageAccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
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
