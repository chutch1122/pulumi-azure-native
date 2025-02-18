


package v20180915

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Lab struct {
	pulumi.CustomResourceState

	Announcement                         LabAnnouncementPropertiesResponsePtrOutput `pulumi:"announcement"`
	ArtifactsStorageAccount              pulumi.StringOutput                        `pulumi:"artifactsStorageAccount"`
	CreatedDate                          pulumi.StringOutput                        `pulumi:"createdDate"`
	DefaultPremiumStorageAccount         pulumi.StringOutput                        `pulumi:"defaultPremiumStorageAccount"`
	DefaultStorageAccount                pulumi.StringOutput                        `pulumi:"defaultStorageAccount"`
	EnvironmentPermission                pulumi.StringPtrOutput                     `pulumi:"environmentPermission"`
	ExtendedProperties                   pulumi.StringMapOutput                     `pulumi:"extendedProperties"`
	LabStorageType                       pulumi.StringPtrOutput                     `pulumi:"labStorageType"`
	LoadBalancerId                       pulumi.StringOutput                        `pulumi:"loadBalancerId"`
	Location                             pulumi.StringPtrOutput                     `pulumi:"location"`
	MandatoryArtifactsResourceIdsLinux   pulumi.StringArrayOutput                   `pulumi:"mandatoryArtifactsResourceIdsLinux"`
	MandatoryArtifactsResourceIdsWindows pulumi.StringArrayOutput                   `pulumi:"mandatoryArtifactsResourceIdsWindows"`
	Name                                 pulumi.StringOutput                        `pulumi:"name"`
	NetworkSecurityGroupId               pulumi.StringOutput                        `pulumi:"networkSecurityGroupId"`
	PremiumDataDiskStorageAccount        pulumi.StringOutput                        `pulumi:"premiumDataDiskStorageAccount"`
	PremiumDataDisks                     pulumi.StringPtrOutput                     `pulumi:"premiumDataDisks"`
	ProvisioningState                    pulumi.StringOutput                        `pulumi:"provisioningState"`
	PublicIpId                           pulumi.StringOutput                        `pulumi:"publicIpId"`
	Support                              LabSupportPropertiesResponsePtrOutput      `pulumi:"support"`
	Tags                                 pulumi.StringMapOutput                     `pulumi:"tags"`
	Type                                 pulumi.StringOutput                        `pulumi:"type"`
	UniqueIdentifier                     pulumi.StringOutput                        `pulumi:"uniqueIdentifier"`
	VaultName                            pulumi.StringOutput                        `pulumi:"vaultName"`
	VmCreationResourceGroup              pulumi.StringOutput                        `pulumi:"vmCreationResourceGroup"`
}


func NewLab(ctx *pulumi.Context,
	name string, args *LabArgs, opts ...pulumi.ResourceOption) (*Lab, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.LabStorageType) {
		args.LabStorageType = pulumi.StringPtr("Premium")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:Lab"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:Lab"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:Lab"),
		},
	})
	opts = append(opts, aliases)
	var resource Lab
	err := ctx.RegisterResource("azure-native:devtestlab/v20180915:Lab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLab(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabState, opts ...pulumi.ResourceOption) (*Lab, error) {
	var resource Lab
	err := ctx.ReadResource("azure-native:devtestlab/v20180915:Lab", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type labState struct {
}

type LabState struct {
}

func (LabState) ElementType() reflect.Type {
	return reflect.TypeOf((*labState)(nil)).Elem()
}

type labArgs struct {
	Announcement                         *LabAnnouncementProperties `pulumi:"announcement"`
	EnvironmentPermission                *string                    `pulumi:"environmentPermission"`
	ExtendedProperties                   map[string]string          `pulumi:"extendedProperties"`
	LabStorageType                       *string                    `pulumi:"labStorageType"`
	Location                             *string                    `pulumi:"location"`
	MandatoryArtifactsResourceIdsLinux   []string                   `pulumi:"mandatoryArtifactsResourceIdsLinux"`
	MandatoryArtifactsResourceIdsWindows []string                   `pulumi:"mandatoryArtifactsResourceIdsWindows"`
	Name                                 *string                    `pulumi:"name"`
	PremiumDataDisks                     *string                    `pulumi:"premiumDataDisks"`
	ResourceGroupName                    string                     `pulumi:"resourceGroupName"`
	Support                              *LabSupportProperties      `pulumi:"support"`
	Tags                                 map[string]string          `pulumi:"tags"`
}


type LabArgs struct {
	Announcement                         LabAnnouncementPropertiesPtrInput
	EnvironmentPermission                pulumi.StringPtrInput
	ExtendedProperties                   pulumi.StringMapInput
	LabStorageType                       pulumi.StringPtrInput
	Location                             pulumi.StringPtrInput
	MandatoryArtifactsResourceIdsLinux   pulumi.StringArrayInput
	MandatoryArtifactsResourceIdsWindows pulumi.StringArrayInput
	Name                                 pulumi.StringPtrInput
	PremiumDataDisks                     pulumi.StringPtrInput
	ResourceGroupName                    pulumi.StringInput
	Support                              LabSupportPropertiesPtrInput
	Tags                                 pulumi.StringMapInput
}

func (LabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labArgs)(nil)).Elem()
}

type LabInput interface {
	pulumi.Input

	ToLabOutput() LabOutput
	ToLabOutputWithContext(ctx context.Context) LabOutput
}

func (*Lab) ElementType() reflect.Type {
	return reflect.TypeOf((**Lab)(nil)).Elem()
}

func (i *Lab) ToLabOutput() LabOutput {
	return i.ToLabOutputWithContext(context.Background())
}

func (i *Lab) ToLabOutputWithContext(ctx context.Context) LabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabOutput)
}

type LabOutput struct{ *pulumi.OutputState }

func (LabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Lab)(nil)).Elem()
}

func (o LabOutput) ToLabOutput() LabOutput {
	return o
}

func (o LabOutput) ToLabOutputWithContext(ctx context.Context) LabOutput {
	return o
}

func (o LabOutput) Announcement() LabAnnouncementPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Lab) LabAnnouncementPropertiesResponsePtrOutput { return v.Announcement }).(LabAnnouncementPropertiesResponsePtrOutput)
}

func (o LabOutput) ArtifactsStorageAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.ArtifactsStorageAccount }).(pulumi.StringOutput)
}

func (o LabOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LabOutput) DefaultPremiumStorageAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.DefaultPremiumStorageAccount }).(pulumi.StringOutput)
}

func (o LabOutput) DefaultStorageAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.DefaultStorageAccount }).(pulumi.StringOutput)
}

func (o LabOutput) EnvironmentPermission() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.EnvironmentPermission }).(pulumi.StringPtrOutput)
}

func (o LabOutput) ExtendedProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringMapOutput { return v.ExtendedProperties }).(pulumi.StringMapOutput)
}

func (o LabOutput) LabStorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.LabStorageType }).(pulumi.StringPtrOutput)
}

func (o LabOutput) LoadBalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.LoadBalancerId }).(pulumi.StringOutput)
}

func (o LabOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LabOutput) MandatoryArtifactsResourceIdsLinux() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringArrayOutput { return v.MandatoryArtifactsResourceIdsLinux }).(pulumi.StringArrayOutput)
}

func (o LabOutput) MandatoryArtifactsResourceIdsWindows() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringArrayOutput { return v.MandatoryArtifactsResourceIdsWindows }).(pulumi.StringArrayOutput)
}

func (o LabOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LabOutput) NetworkSecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.NetworkSecurityGroupId }).(pulumi.StringOutput)
}

func (o LabOutput) PremiumDataDiskStorageAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.PremiumDataDiskStorageAccount }).(pulumi.StringOutput)
}

func (o LabOutput) PremiumDataDisks() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringPtrOutput { return v.PremiumDataDisks }).(pulumi.StringPtrOutput)
}

func (o LabOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LabOutput) PublicIpId() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.PublicIpId }).(pulumi.StringOutput)
}

func (o LabOutput) Support() LabSupportPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Lab) LabSupportPropertiesResponsePtrOutput { return v.Support }).(LabSupportPropertiesResponsePtrOutput)
}

func (o LabOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LabOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o LabOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func (o LabOutput) VaultName() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.VaultName }).(pulumi.StringOutput)
}

func (o LabOutput) VmCreationResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Lab) pulumi.StringOutput { return v.VmCreationResourceGroup }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LabOutput{})
}
