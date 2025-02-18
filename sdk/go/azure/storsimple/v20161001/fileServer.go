


package v20161001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type FileServer struct {
	pulumi.CustomResourceState

	BackupScheduleGroupId pulumi.StringOutput    `pulumi:"backupScheduleGroupId"`
	Description           pulumi.StringPtrOutput `pulumi:"description"`
	DomainName            pulumi.StringOutput    `pulumi:"domainName"`
	Name                  pulumi.StringOutput    `pulumi:"name"`
	StorageDomainId       pulumi.StringOutput    `pulumi:"storageDomainId"`
	Type                  pulumi.StringOutput    `pulumi:"type"`
}


func NewFileServer(ctx *pulumi.Context,
	name string, args *FileServerArgs, opts ...pulumi.ResourceOption) (*FileServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupScheduleGroupId == nil {
		return nil, errors.New("invalid value for required argument 'BackupScheduleGroupId'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageDomainId == nil {
		return nil, errors.New("invalid value for required argument 'StorageDomainId'")
	}
	var resource FileServer
	err := ctx.RegisterResource("azure-native:storsimple/v20161001:FileServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFileServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileServerState, opts ...pulumi.ResourceOption) (*FileServer, error) {
	var resource FileServer
	err := ctx.ReadResource("azure-native:storsimple/v20161001:FileServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type fileServerState struct {
}

type FileServerState struct {
}

func (FileServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileServerState)(nil)).Elem()
}

type fileServerArgs struct {
	BackupScheduleGroupId string  `pulumi:"backupScheduleGroupId"`
	Description           *string `pulumi:"description"`
	DeviceName            string  `pulumi:"deviceName"`
	DomainName            string  `pulumi:"domainName"`
	FileServerName        *string `pulumi:"fileServerName"`
	ManagerName           string  `pulumi:"managerName"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	StorageDomainId       string  `pulumi:"storageDomainId"`
}


type FileServerArgs struct {
	BackupScheduleGroupId pulumi.StringInput
	Description           pulumi.StringPtrInput
	DeviceName            pulumi.StringInput
	DomainName            pulumi.StringInput
	FileServerName        pulumi.StringPtrInput
	ManagerName           pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	StorageDomainId       pulumi.StringInput
}

func (FileServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileServerArgs)(nil)).Elem()
}

type FileServerInput interface {
	pulumi.Input

	ToFileServerOutput() FileServerOutput
	ToFileServerOutputWithContext(ctx context.Context) FileServerOutput
}

func (*FileServer) ElementType() reflect.Type {
	return reflect.TypeOf((**FileServer)(nil)).Elem()
}

func (i *FileServer) ToFileServerOutput() FileServerOutput {
	return i.ToFileServerOutputWithContext(context.Background())
}

func (i *FileServer) ToFileServerOutputWithContext(ctx context.Context) FileServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileServerOutput)
}

type FileServerOutput struct{ *pulumi.OutputState }

func (FileServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileServer)(nil)).Elem()
}

func (o FileServerOutput) ToFileServerOutput() FileServerOutput {
	return o
}

func (o FileServerOutput) ToFileServerOutputWithContext(ctx context.Context) FileServerOutput {
	return o
}

func (o FileServerOutput) BackupScheduleGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *FileServer) pulumi.StringOutput { return v.BackupScheduleGroupId }).(pulumi.StringOutput)
}

func (o FileServerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileServer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o FileServerOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *FileServer) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

func (o FileServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FileServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FileServerOutput) StorageDomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *FileServer) pulumi.StringOutput { return v.StorageDomainId }).(pulumi.StringOutput)
}

func (o FileServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FileServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FileServerOutput{})
}
