


package databoxedge

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StorageAccount struct {
	pulumi.CustomResourceState

	BlobEndpoint               pulumi.StringOutput      `pulumi:"blobEndpoint"`
	ContainerCount             pulumi.IntOutput         `pulumi:"containerCount"`
	DataPolicy                 pulumi.StringOutput      `pulumi:"dataPolicy"`
	Description                pulumi.StringPtrOutput   `pulumi:"description"`
	Name                       pulumi.StringOutput      `pulumi:"name"`
	StorageAccountCredentialId pulumi.StringPtrOutput   `pulumi:"storageAccountCredentialId"`
	StorageAccountStatus       pulumi.StringPtrOutput   `pulumi:"storageAccountStatus"`
	SystemData                 SystemDataResponseOutput `pulumi:"systemData"`
	Type                       pulumi.StringOutput      `pulumi:"type"`
}


func NewStorageAccount(ctx *pulumi.Context,
	name string, args *StorageAccountArgs, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataPolicy == nil {
		return nil, errors.New("invalid value for required argument 'DataPolicy'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:StorageAccount"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:StorageAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageAccount
	err := ctx.RegisterResource("azure-native:databoxedge:StorageAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStorageAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAccountState, opts ...pulumi.ResourceOption) (*StorageAccount, error) {
	var resource StorageAccount
	err := ctx.ReadResource("azure-native:databoxedge:StorageAccount", name, id, state, &resource, opts...)
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
	DataPolicy                 string  `pulumi:"dataPolicy"`
	Description                *string `pulumi:"description"`
	DeviceName                 string  `pulumi:"deviceName"`
	ResourceGroupName          string  `pulumi:"resourceGroupName"`
	StorageAccountCredentialId *string `pulumi:"storageAccountCredentialId"`
	StorageAccountName         *string `pulumi:"storageAccountName"`
	StorageAccountStatus       *string `pulumi:"storageAccountStatus"`
}


type StorageAccountArgs struct {
	DataPolicy                 pulumi.StringInput
	Description                pulumi.StringPtrInput
	DeviceName                 pulumi.StringInput
	ResourceGroupName          pulumi.StringInput
	StorageAccountCredentialId pulumi.StringPtrInput
	StorageAccountName         pulumi.StringPtrInput
	StorageAccountStatus       pulumi.StringPtrInput
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

func (o StorageAccountOutput) BlobEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.BlobEndpoint }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) ContainerCount() pulumi.IntOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.IntOutput { return v.ContainerCount }).(pulumi.IntOutput)
}

func (o StorageAccountOutput) DataPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.DataPolicy }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StorageAccountOutput) StorageAccountCredentialId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.StorageAccountCredentialId }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) StorageAccountStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringPtrOutput { return v.StorageAccountStatus }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StorageAccount) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o StorageAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageAccountOutput{})
}
