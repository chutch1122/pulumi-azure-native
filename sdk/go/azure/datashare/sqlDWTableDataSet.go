


package datashare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlDWTableDataSet struct {
	pulumi.CustomResourceState

	DataSetId           pulumi.StringOutput      `pulumi:"dataSetId"`
	DataWarehouseName   pulumi.StringOutput      `pulumi:"dataWarehouseName"`
	Kind                pulumi.StringOutput      `pulumi:"kind"`
	Name                pulumi.StringOutput      `pulumi:"name"`
	SchemaName          pulumi.StringOutput      `pulumi:"schemaName"`
	SqlServerResourceId pulumi.StringOutput      `pulumi:"sqlServerResourceId"`
	SystemData          SystemDataResponseOutput `pulumi:"systemData"`
	TableName           pulumi.StringOutput      `pulumi:"tableName"`
	Type                pulumi.StringOutput      `pulumi:"type"`
}


func NewSqlDWTableDataSet(ctx *pulumi.Context,
	name string, args *SqlDWTableDataSetArgs, opts ...pulumi.ResourceOption) (*SqlDWTableDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DataWarehouseName == nil {
		return nil, errors.New("invalid value for required argument 'DataWarehouseName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SchemaName == nil {
		return nil, errors.New("invalid value for required argument 'SchemaName'")
	}
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	if args.SqlServerResourceId == nil {
		return nil, errors.New("invalid value for required argument 'SqlServerResourceId'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	args.Kind = pulumi.String("SqlDWTable")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:SqlDWTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:SqlDWTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:SqlDWTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:SqlDWTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:SqlDWTableDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlDWTableDataSet
	err := ctx.RegisterResource("azure-native:datashare:SqlDWTableDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlDWTableDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlDWTableDataSetState, opts ...pulumi.ResourceOption) (*SqlDWTableDataSet, error) {
	var resource SqlDWTableDataSet
	err := ctx.ReadResource("azure-native:datashare:SqlDWTableDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlDWTableDataSetState struct {
}

type SqlDWTableDataSetState struct {
}

func (SqlDWTableDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDWTableDataSetState)(nil)).Elem()
}

type sqlDWTableDataSetArgs struct {
	AccountName         string  `pulumi:"accountName"`
	DataSetName         *string `pulumi:"dataSetName"`
	DataWarehouseName   string  `pulumi:"dataWarehouseName"`
	Kind                string  `pulumi:"kind"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
	SchemaName          string  `pulumi:"schemaName"`
	ShareName           string  `pulumi:"shareName"`
	SqlServerResourceId string  `pulumi:"sqlServerResourceId"`
	TableName           string  `pulumi:"tableName"`
}


type SqlDWTableDataSetArgs struct {
	AccountName         pulumi.StringInput
	DataSetName         pulumi.StringPtrInput
	DataWarehouseName   pulumi.StringInput
	Kind                pulumi.StringInput
	ResourceGroupName   pulumi.StringInput
	SchemaName          pulumi.StringInput
	ShareName           pulumi.StringInput
	SqlServerResourceId pulumi.StringInput
	TableName           pulumi.StringInput
}

func (SqlDWTableDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDWTableDataSetArgs)(nil)).Elem()
}

type SqlDWTableDataSetInput interface {
	pulumi.Input

	ToSqlDWTableDataSetOutput() SqlDWTableDataSetOutput
	ToSqlDWTableDataSetOutputWithContext(ctx context.Context) SqlDWTableDataSetOutput
}

func (*SqlDWTableDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDWTableDataSet)(nil)).Elem()
}

func (i *SqlDWTableDataSet) ToSqlDWTableDataSetOutput() SqlDWTableDataSetOutput {
	return i.ToSqlDWTableDataSetOutputWithContext(context.Background())
}

func (i *SqlDWTableDataSet) ToSqlDWTableDataSetOutputWithContext(ctx context.Context) SqlDWTableDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDWTableDataSetOutput)
}

type SqlDWTableDataSetOutput struct{ *pulumi.OutputState }

func (SqlDWTableDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDWTableDataSet)(nil)).Elem()
}

func (o SqlDWTableDataSetOutput) ToSqlDWTableDataSetOutput() SqlDWTableDataSetOutput {
	return o
}

func (o SqlDWTableDataSetOutput) ToSqlDWTableDataSetOutputWithContext(ctx context.Context) SqlDWTableDataSetOutput {
	return o
}

func (o SqlDWTableDataSetOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSet) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

func (o SqlDWTableDataSetOutput) DataWarehouseName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSet) pulumi.StringOutput { return v.DataWarehouseName }).(pulumi.StringOutput)
}

func (o SqlDWTableDataSetOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSet) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o SqlDWTableDataSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlDWTableDataSetOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSet) pulumi.StringOutput { return v.SchemaName }).(pulumi.StringOutput)
}

func (o SqlDWTableDataSetOutput) SqlServerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSet) pulumi.StringOutput { return v.SqlServerResourceId }).(pulumi.StringOutput)
}

func (o SqlDWTableDataSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqlDWTableDataSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SqlDWTableDataSetOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSet) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

func (o SqlDWTableDataSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlDWTableDataSetOutput{})
}
