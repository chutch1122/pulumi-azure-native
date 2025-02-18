


package v20180915

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Disk struct {
	pulumi.CustomResourceState

	CreatedDate       pulumi.StringOutput    `pulumi:"createdDate"`
	DiskBlobName      pulumi.StringPtrOutput `pulumi:"diskBlobName"`
	DiskSizeGiB       pulumi.IntPtrOutput    `pulumi:"diskSizeGiB"`
	DiskType          pulumi.StringPtrOutput `pulumi:"diskType"`
	DiskUri           pulumi.StringPtrOutput `pulumi:"diskUri"`
	HostCaching       pulumi.StringPtrOutput `pulumi:"hostCaching"`
	LeasedByLabVmId   pulumi.StringPtrOutput `pulumi:"leasedByLabVmId"`
	Location          pulumi.StringPtrOutput `pulumi:"location"`
	ManagedDiskId     pulumi.StringPtrOutput `pulumi:"managedDiskId"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	StorageAccountId  pulumi.StringPtrOutput `pulumi:"storageAccountId"`
	Tags              pulumi.StringMapOutput `pulumi:"tags"`
	Type              pulumi.StringOutput    `pulumi:"type"`
	UniqueIdentifier  pulumi.StringOutput    `pulumi:"uniqueIdentifier"`
}


func NewDisk(ctx *pulumi.Context,
	name string, args *DiskArgs, opts ...pulumi.ResourceOption) (*Disk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:Disk"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:Disk"),
		},
	})
	opts = append(opts, aliases)
	var resource Disk
	err := ctx.RegisterResource("azure-native:devtestlab/v20180915:Disk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskState, opts ...pulumi.ResourceOption) (*Disk, error) {
	var resource Disk
	err := ctx.ReadResource("azure-native:devtestlab/v20180915:Disk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type diskState struct {
}

type DiskState struct {
}

func (DiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskState)(nil)).Elem()
}

type diskArgs struct {
	DiskBlobName      *string           `pulumi:"diskBlobName"`
	DiskSizeGiB       *int              `pulumi:"diskSizeGiB"`
	DiskType          *string           `pulumi:"diskType"`
	DiskUri           *string           `pulumi:"diskUri"`
	HostCaching       *string           `pulumi:"hostCaching"`
	LabName           string            `pulumi:"labName"`
	LeasedByLabVmId   *string           `pulumi:"leasedByLabVmId"`
	Location          *string           `pulumi:"location"`
	ManagedDiskId     *string           `pulumi:"managedDiskId"`
	Name              *string           `pulumi:"name"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	StorageAccountId  *string           `pulumi:"storageAccountId"`
	Tags              map[string]string `pulumi:"tags"`
	UserName          string            `pulumi:"userName"`
}


type DiskArgs struct {
	DiskBlobName      pulumi.StringPtrInput
	DiskSizeGiB       pulumi.IntPtrInput
	DiskType          pulumi.StringPtrInput
	DiskUri           pulumi.StringPtrInput
	HostCaching       pulumi.StringPtrInput
	LabName           pulumi.StringInput
	LeasedByLabVmId   pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ManagedDiskId     pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	StorageAccountId  pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	UserName          pulumi.StringInput
}

func (DiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskArgs)(nil)).Elem()
}

type DiskInput interface {
	pulumi.Input

	ToDiskOutput() DiskOutput
	ToDiskOutputWithContext(ctx context.Context) DiskOutput
}

func (*Disk) ElementType() reflect.Type {
	return reflect.TypeOf((**Disk)(nil)).Elem()
}

func (i *Disk) ToDiskOutput() DiskOutput {
	return i.ToDiskOutputWithContext(context.Background())
}

func (i *Disk) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskOutput)
}

type DiskOutput struct{ *pulumi.OutputState }

func (DiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Disk)(nil)).Elem()
}

func (o DiskOutput) ToDiskOutput() DiskOutput {
	return o
}

func (o DiskOutput) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return o
}

func (o DiskOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o DiskOutput) DiskBlobName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.DiskBlobName }).(pulumi.StringPtrOutput)
}

func (o DiskOutput) DiskSizeGiB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.IntPtrOutput { return v.DiskSizeGiB }).(pulumi.IntPtrOutput)
}

func (o DiskOutput) DiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.DiskType }).(pulumi.StringPtrOutput)
}

func (o DiskOutput) DiskUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.DiskUri }).(pulumi.StringPtrOutput)
}

func (o DiskOutput) HostCaching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.HostCaching }).(pulumi.StringPtrOutput)
}

func (o DiskOutput) LeasedByLabVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.LeasedByLabVmId }).(pulumi.StringPtrOutput)
}

func (o DiskOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DiskOutput) ManagedDiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.ManagedDiskId }).(pulumi.StringPtrOutput)
}

func (o DiskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DiskOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DiskOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

func (o DiskOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DiskOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o DiskOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DiskOutput{})
}
