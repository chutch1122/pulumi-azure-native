


package v20190401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type FileShare struct {
	pulumi.CustomResourceState

	Etag             pulumi.StringOutput    `pulumi:"etag"`
	LastModifiedTime pulumi.StringOutput    `pulumi:"lastModifiedTime"`
	Metadata         pulumi.StringMapOutput `pulumi:"metadata"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	ShareQuota       pulumi.IntPtrOutput    `pulumi:"shareQuota"`
	Type             pulumi.StringOutput    `pulumi:"type"`
}


func NewFileShare(ctx *pulumi.Context,
	name string, args *FileShareArgs, opts ...pulumi.ResourceOption) (*FileShare, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:FileShare"),
		},
	})
	opts = append(opts, aliases)
	var resource FileShare
	err := ctx.RegisterResource("azure-native:storage/v20190401:FileShare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFileShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileShareState, opts ...pulumi.ResourceOption) (*FileShare, error) {
	var resource FileShare
	err := ctx.ReadResource("azure-native:storage/v20190401:FileShare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type fileShareState struct {
}

type FileShareState struct {
}

func (FileShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileShareState)(nil)).Elem()
}

type fileShareArgs struct {
	AccountName       string            `pulumi:"accountName"`
	Metadata          map[string]string `pulumi:"metadata"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ShareName         *string           `pulumi:"shareName"`
	ShareQuota        *int              `pulumi:"shareQuota"`
}


type FileShareArgs struct {
	AccountName       pulumi.StringInput
	Metadata          pulumi.StringMapInput
	ResourceGroupName pulumi.StringInput
	ShareName         pulumi.StringPtrInput
	ShareQuota        pulumi.IntPtrInput
}

func (FileShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileShareArgs)(nil)).Elem()
}

type FileShareInput interface {
	pulumi.Input

	ToFileShareOutput() FileShareOutput
	ToFileShareOutputWithContext(ctx context.Context) FileShareOutput
}

func (*FileShare) ElementType() reflect.Type {
	return reflect.TypeOf((**FileShare)(nil)).Elem()
}

func (i *FileShare) ToFileShareOutput() FileShareOutput {
	return i.ToFileShareOutputWithContext(context.Background())
}

func (i *FileShare) ToFileShareOutputWithContext(ctx context.Context) FileShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileShareOutput)
}

type FileShareOutput struct{ *pulumi.OutputState }

func (FileShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileShare)(nil)).Elem()
}

func (o FileShareOutput) ToFileShareOutput() FileShareOutput {
	return o
}

func (o FileShareOutput) ToFileShareOutputWithContext(ctx context.Context) FileShareOutput {
	return o
}

func (o FileShareOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o FileShareOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o FileShareOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o FileShareOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FileShareOutput) ShareQuota() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FileShare) pulumi.IntPtrOutput { return v.ShareQuota }).(pulumi.IntPtrOutput)
}

func (o FileShareOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FileShareOutput{})
}
