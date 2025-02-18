


package datashare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlDBTableDataSetMapping struct {
	pulumi.CustomResourceState

	DataSetId            pulumi.StringOutput      `pulumi:"dataSetId"`
	DataSetMappingStatus pulumi.StringOutput      `pulumi:"dataSetMappingStatus"`
	DatabaseName         pulumi.StringOutput      `pulumi:"databaseName"`
	Kind                 pulumi.StringOutput      `pulumi:"kind"`
	Name                 pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState    pulumi.StringOutput      `pulumi:"provisioningState"`
	SchemaName           pulumi.StringOutput      `pulumi:"schemaName"`
	SqlServerResourceId  pulumi.StringOutput      `pulumi:"sqlServerResourceId"`
	SystemData           SystemDataResponseOutput `pulumi:"systemData"`
	TableName            pulumi.StringOutput      `pulumi:"tableName"`
	Type                 pulumi.StringOutput      `pulumi:"type"`
}


func NewSqlDBTableDataSetMapping(ctx *pulumi.Context,
	name string, args *SqlDBTableDataSetMappingArgs, opts ...pulumi.ResourceOption) (*SqlDBTableDataSetMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DataSetId == nil {
		return nil, errors.New("invalid value for required argument 'DataSetId'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
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
	args.Kind = pulumi.String("SqlDBTable")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:SqlDBTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:SqlDBTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:SqlDBTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:SqlDBTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:SqlDBTableDataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlDBTableDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare:SqlDBTableDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlDBTableDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlDBTableDataSetMappingState, opts ...pulumi.ResourceOption) (*SqlDBTableDataSetMapping, error) {
	var resource SqlDBTableDataSetMapping
	err := ctx.ReadResource("azure-native:datashare:SqlDBTableDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlDBTableDataSetMappingState struct {
}

type SqlDBTableDataSetMappingState struct {
}

func (SqlDBTableDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDBTableDataSetMappingState)(nil)).Elem()
}

type sqlDBTableDataSetMappingArgs struct {
	AccountName           string  `pulumi:"accountName"`
	DataSetId             string  `pulumi:"dataSetId"`
	DataSetMappingName    *string `pulumi:"dataSetMappingName"`
	DatabaseName          string  `pulumi:"databaseName"`
	Kind                  string  `pulumi:"kind"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	SchemaName            string  `pulumi:"schemaName"`
	ShareSubscriptionName string  `pulumi:"shareSubscriptionName"`
	SqlServerResourceId   string  `pulumi:"sqlServerResourceId"`
	TableName             string  `pulumi:"tableName"`
}


type SqlDBTableDataSetMappingArgs struct {
	AccountName           pulumi.StringInput
	DataSetId             pulumi.StringInput
	DataSetMappingName    pulumi.StringPtrInput
	DatabaseName          pulumi.StringInput
	Kind                  pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	SchemaName            pulumi.StringInput
	ShareSubscriptionName pulumi.StringInput
	SqlServerResourceId   pulumi.StringInput
	TableName             pulumi.StringInput
}

func (SqlDBTableDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDBTableDataSetMappingArgs)(nil)).Elem()
}

type SqlDBTableDataSetMappingInput interface {
	pulumi.Input

	ToSqlDBTableDataSetMappingOutput() SqlDBTableDataSetMappingOutput
	ToSqlDBTableDataSetMappingOutputWithContext(ctx context.Context) SqlDBTableDataSetMappingOutput
}

func (*SqlDBTableDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDBTableDataSetMapping)(nil)).Elem()
}

func (i *SqlDBTableDataSetMapping) ToSqlDBTableDataSetMappingOutput() SqlDBTableDataSetMappingOutput {
	return i.ToSqlDBTableDataSetMappingOutputWithContext(context.Background())
}

func (i *SqlDBTableDataSetMapping) ToSqlDBTableDataSetMappingOutputWithContext(ctx context.Context) SqlDBTableDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDBTableDataSetMappingOutput)
}

type SqlDBTableDataSetMappingOutput struct{ *pulumi.OutputState }

func (SqlDBTableDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDBTableDataSetMapping)(nil)).Elem()
}

func (o SqlDBTableDataSetMappingOutput) ToSqlDBTableDataSetMappingOutput() SqlDBTableDataSetMappingOutput {
	return o
}

func (o SqlDBTableDataSetMappingOutput) ToSqlDBTableDataSetMappingOutputWithContext(ctx context.Context) SqlDBTableDataSetMappingOutput {
	return o
}

func (o SqlDBTableDataSetMappingOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDBTableDataSetMapping) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

func (o SqlDBTableDataSetMappingOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDBTableDataSetMapping) pulumi.StringOutput { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

func (o SqlDBTableDataSetMappingOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDBTableDataSetMapping) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o SqlDBTableDataSetMappingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDBTableDataSetMapping) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o SqlDBTableDataSetMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDBTableDataSetMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlDBTableDataSetMappingOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDBTableDataSetMapping) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SqlDBTableDataSetMappingOutput) SchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDBTableDataSetMapping) pulumi.StringOutput { return v.SchemaName }).(pulumi.StringOutput)
}

func (o SqlDBTableDataSetMappingOutput) SqlServerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDBTableDataSetMapping) pulumi.StringOutput { return v.SqlServerResourceId }).(pulumi.StringOutput)
}

func (o SqlDBTableDataSetMappingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqlDBTableDataSetMapping) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SqlDBTableDataSetMappingOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDBTableDataSetMapping) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

func (o SqlDBTableDataSetMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlDBTableDataSetMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlDBTableDataSetMappingOutput{})
}
