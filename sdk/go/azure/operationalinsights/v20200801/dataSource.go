


package v20200801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataSource struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput `pulumi:"etag"`
	Kind       pulumi.StringOutput    `pulumi:"kind"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	Properties pulumi.AnyOutput       `pulumi:"properties"`
	Tags       pulumi.StringMapOutput `pulumi:"tags"`
	Type       pulumi.StringOutput    `pulumi:"type"`
}


func NewDataSource(ctx *pulumi.Context,
	name string, args *DataSourceArgs, opts ...pulumi.ResourceOption) (*DataSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights:DataSource"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20151101preview:DataSource"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200301preview:DataSource"),
		},
	})
	opts = append(opts, aliases)
	var resource DataSource
	err := ctx.RegisterResource("azure-native:operationalinsights/v20200801:DataSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataSourceState, opts ...pulumi.ResourceOption) (*DataSource, error) {
	var resource DataSource
	err := ctx.ReadResource("azure-native:operationalinsights/v20200801:DataSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dataSourceState struct {
}

type DataSourceState struct {
}

func (DataSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceState)(nil)).Elem()
}

type dataSourceArgs struct {
	DataSourceName    *string           `pulumi:"dataSourceName"`
	Kind              string            `pulumi:"kind"`
	Properties        interface{}       `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	WorkspaceName     string            `pulumi:"workspaceName"`
}


type DataSourceArgs struct {
	DataSourceName    pulumi.StringPtrInput
	Kind              pulumi.StringInput
	Properties        pulumi.Input
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	WorkspaceName     pulumi.StringInput
}

func (DataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceArgs)(nil)).Elem()
}

type DataSourceInput interface {
	pulumi.Input

	ToDataSourceOutput() DataSourceOutput
	ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput
}

func (*DataSource) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSource)(nil)).Elem()
}

func (i *DataSource) ToDataSourceOutput() DataSourceOutput {
	return i.ToDataSourceOutputWithContext(context.Background())
}

func (i *DataSource) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceOutput)
}

type DataSourceOutput struct{ *pulumi.OutputState }

func (DataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSource)(nil)).Elem()
}

func (o DataSourceOutput) ToDataSourceOutput() DataSourceOutput {
	return o
}

func (o DataSourceOutput) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return o
}

func (o DataSourceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o DataSourceOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o DataSourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DataSourceOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *DataSource) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

func (o DataSourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataSourceOutput{})
}
