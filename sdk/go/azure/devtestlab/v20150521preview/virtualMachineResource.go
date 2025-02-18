


package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type VirtualMachineResource struct {
	pulumi.CustomResourceState

	ArtifactDeploymentStatus   ArtifactDeploymentStatusPropertiesResponsePtrOutput `pulumi:"artifactDeploymentStatus"`
	Artifacts                  ArtifactInstallPropertiesResponseArrayOutput        `pulumi:"artifacts"`
	ComputeId                  pulumi.StringPtrOutput                              `pulumi:"computeId"`
	CreatedByUser              pulumi.StringPtrOutput                              `pulumi:"createdByUser"`
	CreatedByUserId            pulumi.StringPtrOutput                              `pulumi:"createdByUserId"`
	CustomImageId              pulumi.StringPtrOutput                              `pulumi:"customImageId"`
	DisallowPublicIpAddress    pulumi.BoolPtrOutput                                `pulumi:"disallowPublicIpAddress"`
	Fqdn                       pulumi.StringPtrOutput                              `pulumi:"fqdn"`
	GalleryImageReference      GalleryImageReferenceResponsePtrOutput              `pulumi:"galleryImageReference"`
	IsAuthenticationWithSshKey pulumi.BoolPtrOutput                                `pulumi:"isAuthenticationWithSshKey"`
	LabSubnetName              pulumi.StringPtrOutput                              `pulumi:"labSubnetName"`
	LabVirtualNetworkId        pulumi.StringPtrOutput                              `pulumi:"labVirtualNetworkId"`
	Location                   pulumi.StringPtrOutput                              `pulumi:"location"`
	Name                       pulumi.StringPtrOutput                              `pulumi:"name"`
	Notes                      pulumi.StringPtrOutput                              `pulumi:"notes"`
	OsType                     pulumi.StringPtrOutput                              `pulumi:"osType"`
	OwnerObjectId              pulumi.StringPtrOutput                              `pulumi:"ownerObjectId"`
	Password                   pulumi.StringPtrOutput                              `pulumi:"password"`
	ProvisioningState          pulumi.StringPtrOutput                              `pulumi:"provisioningState"`
	Size                       pulumi.StringPtrOutput                              `pulumi:"size"`
	SshKey                     pulumi.StringPtrOutput                              `pulumi:"sshKey"`
	Tags                       pulumi.StringMapOutput                              `pulumi:"tags"`
	Type                       pulumi.StringPtrOutput                              `pulumi:"type"`
	UserName                   pulumi.StringPtrOutput                              `pulumi:"userName"`
}


func NewVirtualMachineResource(ctx *pulumi.Context,
	name string, args *VirtualMachineResourceArgs, opts ...pulumi.ResourceOption) (*VirtualMachineResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:VirtualMachineResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:VirtualMachineResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:VirtualMachineResource"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineResource
	err := ctx.RegisterResource("azure-native:devtestlab/v20150521preview:VirtualMachineResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachineResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineResourceState, opts ...pulumi.ResourceOption) (*VirtualMachineResource, error) {
	var resource VirtualMachineResource
	err := ctx.ReadResource("azure-native:devtestlab/v20150521preview:VirtualMachineResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualMachineResourceState struct {
}

type VirtualMachineResourceState struct {
}

func (VirtualMachineResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineResourceState)(nil)).Elem()
}

type virtualMachineResourceArgs struct {
	ArtifactDeploymentStatus   *ArtifactDeploymentStatusProperties `pulumi:"artifactDeploymentStatus"`
	Artifacts                  []ArtifactInstallProperties         `pulumi:"artifacts"`
	ComputeId                  *string                             `pulumi:"computeId"`
	CreatedByUser              *string                             `pulumi:"createdByUser"`
	CreatedByUserId            *string                             `pulumi:"createdByUserId"`
	CustomImageId              *string                             `pulumi:"customImageId"`
	DisallowPublicIpAddress    *bool                               `pulumi:"disallowPublicIpAddress"`
	Fqdn                       *string                             `pulumi:"fqdn"`
	GalleryImageReference      *GalleryImageReference              `pulumi:"galleryImageReference"`
	Id                         *string                             `pulumi:"id"`
	IsAuthenticationWithSshKey *bool                               `pulumi:"isAuthenticationWithSshKey"`
	LabName                    string                              `pulumi:"labName"`
	LabSubnetName              *string                             `pulumi:"labSubnetName"`
	LabVirtualNetworkId        *string                             `pulumi:"labVirtualNetworkId"`
	Location                   *string                             `pulumi:"location"`
	Name                       *string                             `pulumi:"name"`
	Notes                      *string                             `pulumi:"notes"`
	OsType                     *string                             `pulumi:"osType"`
	OwnerObjectId              *string                             `pulumi:"ownerObjectId"`
	Password                   *string                             `pulumi:"password"`
	ProvisioningState          *string                             `pulumi:"provisioningState"`
	ResourceGroupName          string                              `pulumi:"resourceGroupName"`
	Size                       *string                             `pulumi:"size"`
	SshKey                     *string                             `pulumi:"sshKey"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       *string                             `pulumi:"type"`
	UserName                   *string                             `pulumi:"userName"`
}


type VirtualMachineResourceArgs struct {
	ArtifactDeploymentStatus   ArtifactDeploymentStatusPropertiesPtrInput
	Artifacts                  ArtifactInstallPropertiesArrayInput
	ComputeId                  pulumi.StringPtrInput
	CreatedByUser              pulumi.StringPtrInput
	CreatedByUserId            pulumi.StringPtrInput
	CustomImageId              pulumi.StringPtrInput
	DisallowPublicIpAddress    pulumi.BoolPtrInput
	Fqdn                       pulumi.StringPtrInput
	GalleryImageReference      GalleryImageReferencePtrInput
	Id                         pulumi.StringPtrInput
	IsAuthenticationWithSshKey pulumi.BoolPtrInput
	LabName                    pulumi.StringInput
	LabSubnetName              pulumi.StringPtrInput
	LabVirtualNetworkId        pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	Name                       pulumi.StringPtrInput
	Notes                      pulumi.StringPtrInput
	OsType                     pulumi.StringPtrInput
	OwnerObjectId              pulumi.StringPtrInput
	Password                   pulumi.StringPtrInput
	ProvisioningState          pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	Size                       pulumi.StringPtrInput
	SshKey                     pulumi.StringPtrInput
	Tags                       pulumi.StringMapInput
	Type                       pulumi.StringPtrInput
	UserName                   pulumi.StringPtrInput
}

func (VirtualMachineResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineResourceArgs)(nil)).Elem()
}

type VirtualMachineResourceInput interface {
	pulumi.Input

	ToVirtualMachineResourceOutput() VirtualMachineResourceOutput
	ToVirtualMachineResourceOutputWithContext(ctx context.Context) VirtualMachineResourceOutput
}

func (*VirtualMachineResource) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineResource)(nil)).Elem()
}

func (i *VirtualMachineResource) ToVirtualMachineResourceOutput() VirtualMachineResourceOutput {
	return i.ToVirtualMachineResourceOutputWithContext(context.Background())
}

func (i *VirtualMachineResource) ToVirtualMachineResourceOutputWithContext(ctx context.Context) VirtualMachineResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineResourceOutput)
}

type VirtualMachineResourceOutput struct{ *pulumi.OutputState }

func (VirtualMachineResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineResource)(nil)).Elem()
}

func (o VirtualMachineResourceOutput) ToVirtualMachineResourceOutput() VirtualMachineResourceOutput {
	return o
}

func (o VirtualMachineResourceOutput) ToVirtualMachineResourceOutputWithContext(ctx context.Context) VirtualMachineResourceOutput {
	return o
}

func (o VirtualMachineResourceOutput) ArtifactDeploymentStatus() ArtifactDeploymentStatusPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) ArtifactDeploymentStatusPropertiesResponsePtrOutput {
		return v.ArtifactDeploymentStatus
	}).(ArtifactDeploymentStatusPropertiesResponsePtrOutput)
}

func (o VirtualMachineResourceOutput) Artifacts() ArtifactInstallPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineResource) ArtifactInstallPropertiesResponseArrayOutput { return v.Artifacts }).(ArtifactInstallPropertiesResponseArrayOutput)
}

func (o VirtualMachineResourceOutput) ComputeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringPtrOutput { return v.ComputeId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceOutput) CreatedByUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringPtrOutput { return v.CreatedByUser }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceOutput) CreatedByUserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringPtrOutput { return v.CreatedByUserId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceOutput) CustomImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringPtrOutput { return v.CustomImageId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceOutput) DisallowPublicIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.BoolPtrOutput { return v.DisallowPublicIpAddress }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineResourceOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringPtrOutput { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceOutput) GalleryImageReference() GalleryImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) GalleryImageReferenceResponsePtrOutput { return v.GalleryImageReference }).(GalleryImageReferenceResponsePtrOutput)
}

func (o VirtualMachineResourceOutput) IsAuthenticationWithSshKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.BoolPtrOutput { return v.IsAuthenticationWithSshKey }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineResourceOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringPtrOutput { return v.LabSubnetName }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceOutput) LabVirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringPtrOutput { return v.LabVirtualNetworkId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceOutput) OwnerObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringPtrOutput { return v.OwnerObjectId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringPtrOutput { return v.Size }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceOutput) SshKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringPtrOutput { return v.SshKey }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualMachineResourceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineResource) pulumi.StringPtrOutput { return v.UserName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineResourceOutput{})
}
