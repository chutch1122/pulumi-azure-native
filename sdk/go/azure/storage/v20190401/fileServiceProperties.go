


package v20190401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type FileServiceProperties struct {
	pulumi.CustomResourceState

	Cors CorsRulesResponsePtrOutput `pulumi:"cors"`
	Name pulumi.StringOutput        `pulumi:"name"`
	Type pulumi.StringOutput        `pulumi:"type"`
}


func NewFileServiceProperties(ctx *pulumi.Context,
	name string, args *FileServicePropertiesArgs, opts ...pulumi.ResourceOption) (*FileServiceProperties, error) {
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
			Type: pulumi.String("azure-native:storage:FileServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:FileServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:FileServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:FileServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:FileServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:FileServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:FileServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:FileServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:FileServiceProperties"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:FileServiceProperties"),
		},
	})
	opts = append(opts, aliases)
	var resource FileServiceProperties
	err := ctx.RegisterResource("azure-native:storage/v20190401:FileServiceProperties", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFileServiceProperties(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileServicePropertiesState, opts ...pulumi.ResourceOption) (*FileServiceProperties, error) {
	var resource FileServiceProperties
	err := ctx.ReadResource("azure-native:storage/v20190401:FileServiceProperties", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type fileServicePropertiesState struct {
}

type FileServicePropertiesState struct {
}

func (FileServicePropertiesState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileServicePropertiesState)(nil)).Elem()
}

type fileServicePropertiesArgs struct {
	AccountName       string     `pulumi:"accountName"`
	Cors              *CorsRules `pulumi:"cors"`
	FileServicesName  *string    `pulumi:"fileServicesName"`
	ResourceGroupName string     `pulumi:"resourceGroupName"`
}


type FileServicePropertiesArgs struct {
	AccountName       pulumi.StringInput
	Cors              CorsRulesPtrInput
	FileServicesName  pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (FileServicePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileServicePropertiesArgs)(nil)).Elem()
}

type FileServicePropertiesInput interface {
	pulumi.Input

	ToFileServicePropertiesOutput() FileServicePropertiesOutput
	ToFileServicePropertiesOutputWithContext(ctx context.Context) FileServicePropertiesOutput
}

func (*FileServiceProperties) ElementType() reflect.Type {
	return reflect.TypeOf((**FileServiceProperties)(nil)).Elem()
}

func (i *FileServiceProperties) ToFileServicePropertiesOutput() FileServicePropertiesOutput {
	return i.ToFileServicePropertiesOutputWithContext(context.Background())
}

func (i *FileServiceProperties) ToFileServicePropertiesOutputWithContext(ctx context.Context) FileServicePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileServicePropertiesOutput)
}

type FileServicePropertiesOutput struct{ *pulumi.OutputState }

func (FileServicePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileServiceProperties)(nil)).Elem()
}

func (o FileServicePropertiesOutput) ToFileServicePropertiesOutput() FileServicePropertiesOutput {
	return o
}

func (o FileServicePropertiesOutput) ToFileServicePropertiesOutputWithContext(ctx context.Context) FileServicePropertiesOutput {
	return o
}

func (o FileServicePropertiesOutput) Cors() CorsRulesResponsePtrOutput {
	return o.ApplyT(func(v *FileServiceProperties) CorsRulesResponsePtrOutput { return v.Cors }).(CorsRulesResponsePtrOutput)
}

func (o FileServicePropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FileServiceProperties) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FileServicePropertiesOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FileServiceProperties) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FileServicePropertiesOutput{})
}
