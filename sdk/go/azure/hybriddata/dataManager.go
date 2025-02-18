


package hybriddata

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataManager struct {
	pulumi.CustomResourceState

	Etag     pulumi.StringPtrOutput `pulumi:"etag"`
	Location pulumi.StringOutput    `pulumi:"location"`
	Name     pulumi.StringOutput    `pulumi:"name"`
	Sku      SkuResponsePtrOutput   `pulumi:"sku"`
	Tags     pulumi.StringMapOutput `pulumi:"tags"`
	Type     pulumi.StringOutput    `pulumi:"type"`
}


func NewDataManager(ctx *pulumi.Context,
	name string, args *DataManagerArgs, opts ...pulumi.ResourceOption) (*DataManager, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybriddata/v20160601:DataManager"),
		},
		{
			Type: pulumi.String("azure-native:hybriddata/v20190601:DataManager"),
		},
	})
	opts = append(opts, aliases)
	var resource DataManager
	err := ctx.RegisterResource("azure-native:hybriddata:DataManager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataManager(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataManagerState, opts ...pulumi.ResourceOption) (*DataManager, error) {
	var resource DataManager
	err := ctx.ReadResource("azure-native:hybriddata:DataManager", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dataManagerState struct {
}

type DataManagerState struct {
}

func (DataManagerState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataManagerState)(nil)).Elem()
}

type dataManagerArgs struct {
	DataManagerName   *string           `pulumi:"dataManagerName"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Sku               *Sku              `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}


type DataManagerArgs struct {
	DataManagerName   pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               SkuPtrInput
	Tags              pulumi.StringMapInput
}

func (DataManagerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataManagerArgs)(nil)).Elem()
}

type DataManagerInput interface {
	pulumi.Input

	ToDataManagerOutput() DataManagerOutput
	ToDataManagerOutputWithContext(ctx context.Context) DataManagerOutput
}

func (*DataManager) ElementType() reflect.Type {
	return reflect.TypeOf((**DataManager)(nil)).Elem()
}

func (i *DataManager) ToDataManagerOutput() DataManagerOutput {
	return i.ToDataManagerOutputWithContext(context.Background())
}

func (i *DataManager) ToDataManagerOutputWithContext(ctx context.Context) DataManagerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataManagerOutput)
}

type DataManagerOutput struct{ *pulumi.OutputState }

func (DataManagerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataManager)(nil)).Elem()
}

func (o DataManagerOutput) ToDataManagerOutput() DataManagerOutput {
	return o
}

func (o DataManagerOutput) ToDataManagerOutputWithContext(ctx context.Context) DataManagerOutput {
	return o
}

func (o DataManagerOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataManager) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o DataManagerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataManager) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DataManagerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataManager) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DataManagerOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *DataManager) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o DataManagerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataManager) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DataManagerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataManager) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataManagerOutput{})
}
