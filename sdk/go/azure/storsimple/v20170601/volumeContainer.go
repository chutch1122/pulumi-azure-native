


package v20170601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VolumeContainer struct {
	pulumi.CustomResourceState

	BandWidthRateInMbps           pulumi.IntPtrOutput                        `pulumi:"bandWidthRateInMbps"`
	BandwidthSettingId            pulumi.StringPtrOutput                     `pulumi:"bandwidthSettingId"`
	EncryptionKey                 AsymmetricEncryptedSecretResponsePtrOutput `pulumi:"encryptionKey"`
	EncryptionStatus              pulumi.StringOutput                        `pulumi:"encryptionStatus"`
	Kind                          pulumi.StringPtrOutput                     `pulumi:"kind"`
	Name                          pulumi.StringOutput                        `pulumi:"name"`
	OwnerShipStatus               pulumi.StringOutput                        `pulumi:"ownerShipStatus"`
	StorageAccountCredentialId    pulumi.StringOutput                        `pulumi:"storageAccountCredentialId"`
	TotalCloudStorageUsageInBytes pulumi.Float64Output                       `pulumi:"totalCloudStorageUsageInBytes"`
	Type                          pulumi.StringOutput                        `pulumi:"type"`
	VolumeCount                   pulumi.IntOutput                           `pulumi:"volumeCount"`
}


func NewVolumeContainer(ctx *pulumi.Context,
	name string, args *VolumeContainerArgs, opts ...pulumi.ResourceOption) (*VolumeContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccountCredentialId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountCredentialId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storsimple:VolumeContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource VolumeContainer
	err := ctx.RegisterResource("azure-native:storsimple/v20170601:VolumeContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVolumeContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeContainerState, opts ...pulumi.ResourceOption) (*VolumeContainer, error) {
	var resource VolumeContainer
	err := ctx.ReadResource("azure-native:storsimple/v20170601:VolumeContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type volumeContainerState struct {
}

type VolumeContainerState struct {
}

func (VolumeContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeContainerState)(nil)).Elem()
}

type volumeContainerArgs struct {
	BandWidthRateInMbps        *int                       `pulumi:"bandWidthRateInMbps"`
	BandwidthSettingId         *string                    `pulumi:"bandwidthSettingId"`
	DeviceName                 string                     `pulumi:"deviceName"`
	EncryptionKey              *AsymmetricEncryptedSecret `pulumi:"encryptionKey"`
	Kind                       *Kind                      `pulumi:"kind"`
	ManagerName                string                     `pulumi:"managerName"`
	ResourceGroupName          string                     `pulumi:"resourceGroupName"`
	StorageAccountCredentialId string                     `pulumi:"storageAccountCredentialId"`
	VolumeContainerName        *string                    `pulumi:"volumeContainerName"`
}


type VolumeContainerArgs struct {
	BandWidthRateInMbps        pulumi.IntPtrInput
	BandwidthSettingId         pulumi.StringPtrInput
	DeviceName                 pulumi.StringInput
	EncryptionKey              AsymmetricEncryptedSecretPtrInput
	Kind                       KindPtrInput
	ManagerName                pulumi.StringInput
	ResourceGroupName          pulumi.StringInput
	StorageAccountCredentialId pulumi.StringInput
	VolumeContainerName        pulumi.StringPtrInput
}

func (VolumeContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeContainerArgs)(nil)).Elem()
}

type VolumeContainerInput interface {
	pulumi.Input

	ToVolumeContainerOutput() VolumeContainerOutput
	ToVolumeContainerOutputWithContext(ctx context.Context) VolumeContainerOutput
}

func (*VolumeContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeContainer)(nil)).Elem()
}

func (i *VolumeContainer) ToVolumeContainerOutput() VolumeContainerOutput {
	return i.ToVolumeContainerOutputWithContext(context.Background())
}

func (i *VolumeContainer) ToVolumeContainerOutputWithContext(ctx context.Context) VolumeContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeContainerOutput)
}

type VolumeContainerOutput struct{ *pulumi.OutputState }

func (VolumeContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeContainer)(nil)).Elem()
}

func (o VolumeContainerOutput) ToVolumeContainerOutput() VolumeContainerOutput {
	return o
}

func (o VolumeContainerOutput) ToVolumeContainerOutputWithContext(ctx context.Context) VolumeContainerOutput {
	return o
}

func (o VolumeContainerOutput) BandWidthRateInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VolumeContainer) pulumi.IntPtrOutput { return v.BandWidthRateInMbps }).(pulumi.IntPtrOutput)
}

func (o VolumeContainerOutput) BandwidthSettingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeContainer) pulumi.StringPtrOutput { return v.BandwidthSettingId }).(pulumi.StringPtrOutput)
}

func (o VolumeContainerOutput) EncryptionKey() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v *VolumeContainer) AsymmetricEncryptedSecretResponsePtrOutput { return v.EncryptionKey }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

func (o VolumeContainerOutput) EncryptionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeContainer) pulumi.StringOutput { return v.EncryptionStatus }).(pulumi.StringOutput)
}

func (o VolumeContainerOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeContainer) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o VolumeContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VolumeContainerOutput) OwnerShipStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeContainer) pulumi.StringOutput { return v.OwnerShipStatus }).(pulumi.StringOutput)
}

func (o VolumeContainerOutput) StorageAccountCredentialId() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeContainer) pulumi.StringOutput { return v.StorageAccountCredentialId }).(pulumi.StringOutput)
}

func (o VolumeContainerOutput) TotalCloudStorageUsageInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *VolumeContainer) pulumi.Float64Output { return v.TotalCloudStorageUsageInBytes }).(pulumi.Float64Output)
}

func (o VolumeContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VolumeContainerOutput) VolumeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *VolumeContainer) pulumi.IntOutput { return v.VolumeCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeContainerOutput{})
}
