


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlDWTableDataSetMapping struct {
	pulumi.CustomResourceState

	DataSetId            pulumi.StringOutput      `pulumi:"dataSetId"`
	DataSetMappingStatus pulumi.StringOutput      `pulumi:"dataSetMappingStatus"`
	DataWarehouseName    pulumi.StringOutput      `pulumi:"dataWarehouseName"`
	Kind                 pulumi.StringOutput      `pulumi:"kind"`
	Name                 pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState    pulumi.StringOutput      `pulumi:"provisioningState"`
	SchemaName           pulumi.StringOutput      `pulumi:"schemaName"`
	SqlServerResourceId  pulumi.StringOutput      `pulumi:"sqlServerResourceId"`
	SystemData           SystemDataResponseOutput `pulumi:"systemData"`
	TableName            pulumi.StringOutput      `pulumi:"tableName"`
	Type                 pulumi.StringOutput      `pulumi:"type"`
}


func NewSqlDWTableDataSetMapping(ctx *pulumi.Context,
	name string, args *SqlDWTableDataSetMappingArgs, opts ...pulumi.ResourceOption) (*SqlDWTableDataSetMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DataSetId == nil {
		return nil, errors.New("invalid value for required argument 'DataSetId'")
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
	if args.ShareSubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'ShareSubscriptionName'")
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
			Type: pulumi.String("azure-native:datashare:SqlDWTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:SqlDWTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:SqlDWTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:SqlDWTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:SqlDWTableDataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlDWTableDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare/v20201001preview:SqlDWTableDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlDWTableDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlDWTableDataSetMappingState, opts ...pulumi.ResourceOption) (*SqlDWTableDataSetMapping, error) {
	var resource SqlDWTableDataSetMapping
	err := ctx.ReadResource("azure-native:datashare/v20201001preview:SqlDWTableDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlDWTableDataSetMappingState struct {
}

type SqlDWTableDataSetMappingState struct {
}

func (SqlDWTableDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDWTableDataSetMappingState)(nil)).Elem()
}

type sqlDWTableDataSetMappingArgs struct {
	AccountName           string  `pulumi:"accountName"`
	DataSetId             string  `pulumi:"dataSetId"`
	DataSetMappingName    *string `pulumi:"dataSetMappingName"`
	DataWarehouseName     string  `pulumi:"dataWarehouseName"`
	Kind                  string  `pulumi:"kind"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	SchemaName            string  `pulumi:"schemaName"`
	ShareSubscriptionName string  `pulumi:"shareSubscriptionName"`
	SqlServerResourceId   string  `pulumi:"sqlServerResourceId"`
	TableName             string  `pulumi:"tableName"`
}


type SqlDWTableDataSetMappingArgs struct {
	AccountName           pulumi.StringInput
	DataSetId             pulumi.StringInput
	DataSetMappingName    pulumi.StringPtrInput
	DataWarehouseName     pulumi.StringInput
	Kind                  pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	SchemaName            pulumi.StringInput
	ShareSubscriptionName pulumi.StringInput
	SqlServerResourceId   pulumi.StringInput
	TableName             pulumi.StringInput
}

func (SqlDWTableDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDWTableDataSetMappingArgs)(nil)).Elem()
}

type SqlDWTableDataSetMappingInput interface {
	pulumi.Input

	ToSqlDWTableDataSetMappingOutput() SqlDWTableDataSetMappingOutput
	ToSqlDWTableDataSetMappingOutputWithContext(ctx context.Context) SqlDWTableDataSetMappingOutput
}

func (*SqlDWTableDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDWTableDataSetMapping)(nil)).Elem()
}

func (i *SqlDWTableDataSetMapping) ToSqlDWTableDataSetMappingOutput() SqlDWTableDataSetMappingOutput {
	return i.ToSqlDWTableDataSetMappingOutputWithContext(context.Background())
}

func (i *SqlDWTableDataSetMapping) ToSqlDWTableDataSetMappingOutputWithContext(ctx context.Context) SqlDWTableDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDWTableDataSetMappingOutput)
}

type SqlDWTableDataSetMappingOutput struct{ *pulumi.OutputState }

func (SqlDWTableDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDWTableDataSetMapping)(nil)).Elem()
}

func (o SqlDWTableDataSetMappingOutput) ToSqlDWTableDataSetMappingOutput() SqlDWTableDataSetMappingOutput {
	return o
}

func (o SqlDWTableDataSetMappingOutput) ToSqlDWTableDataSetMappingOutputWithContext(ctx context.Context) SqlDWTableDataSetMappingOutput {
	return o
}

func (o SqlDWTableDataSetMappingOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSetMapping) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

func (o SqlDWTableDataSetMappingOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSetMapping) pulumi.StringOutput { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o SqlDWTableDataSetMappingOutput) DataWarehouseName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSetMapping) pulumi.StringOutput { return v.DataWarehouseName }).(pulumi.StringOutput)
}

func (o SqlDWTableDataSetMappingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSetMapping) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o SqlDWTableDataSetMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSetMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlDWTableDataSetMappingOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSetMapping) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SqlDWTableDataSetMappingOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSetMapping) pulumi.StringOutput { return v.SchemaName }).(pulumi.StringOutput)
}

func (o SqlDWTableDataSetMappingOutput) SqlServerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSetMapping) pulumi.StringOutput { return v.SqlServerResourceId }).(pulumi.StringOutput)
}

func (o SqlDWTableDataSetMappingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqlDWTableDataSetMapping) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SqlDWTableDataSetMappingOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSetMapping) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

func (o SqlDWTableDataSetMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDWTableDataSetMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlDWTableDataSetMappingOutput{})
}
