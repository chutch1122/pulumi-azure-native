


package v20210501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlPool struct {
	pulumi.CustomResourceState

	Collation             pulumi.StringPtrOutput  `pulumi:"collation"`
	CreateMode            pulumi.StringPtrOutput  `pulumi:"createMode"`
	CreationDate          pulumi.StringPtrOutput  `pulumi:"creationDate"`
	Location              pulumi.StringOutput     `pulumi:"location"`
	MaxSizeBytes          pulumi.Float64PtrOutput `pulumi:"maxSizeBytes"`
	Name                  pulumi.StringOutput     `pulumi:"name"`
	ProvisioningState     pulumi.StringPtrOutput  `pulumi:"provisioningState"`
	RecoverableDatabaseId pulumi.StringPtrOutput  `pulumi:"recoverableDatabaseId"`
	RestorePointInTime    pulumi.StringPtrOutput  `pulumi:"restorePointInTime"`
	Sku                   SkuResponsePtrOutput    `pulumi:"sku"`
	SourceDatabaseId      pulumi.StringPtrOutput  `pulumi:"sourceDatabaseId"`
	Status                pulumi.StringPtrOutput  `pulumi:"status"`
	StorageAccountType    pulumi.StringPtrOutput  `pulumi:"storageAccountType"`
	Tags                  pulumi.StringMapOutput  `pulumi:"tags"`
	Type                  pulumi.StringOutput     `pulumi:"type"`
}


func NewSqlPool(ctx *pulumi.Context,
	name string, args *SqlPoolArgs, opts ...pulumi.ResourceOption) (*SqlPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:SqlPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:SqlPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20200401preview:SqlPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:SqlPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:SqlPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:SqlPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:SqlPool"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:SqlPool"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlPool
	err := ctx.RegisterResource("azure-native:synapse/v20210501:SqlPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlPoolState, opts ...pulumi.ResourceOption) (*SqlPool, error) {
	var resource SqlPool
	err := ctx.ReadResource("azure-native:synapse/v20210501:SqlPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlPoolState struct {
}

type SqlPoolState struct {
}

func (SqlPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolState)(nil)).Elem()
}

type sqlPoolArgs struct {
	Collation             *string           `pulumi:"collation"`
	CreateMode            *string           `pulumi:"createMode"`
	CreationDate          *string           `pulumi:"creationDate"`
	Location              *string           `pulumi:"location"`
	MaxSizeBytes          *float64          `pulumi:"maxSizeBytes"`
	ProvisioningState     *string           `pulumi:"provisioningState"`
	RecoverableDatabaseId *string           `pulumi:"recoverableDatabaseId"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	RestorePointInTime    *string           `pulumi:"restorePointInTime"`
	Sku                   *Sku              `pulumi:"sku"`
	SourceDatabaseId      *string           `pulumi:"sourceDatabaseId"`
	SqlPoolName           *string           `pulumi:"sqlPoolName"`
	Status                *string           `pulumi:"status"`
	StorageAccountType    *string           `pulumi:"storageAccountType"`
	Tags                  map[string]string `pulumi:"tags"`
	WorkspaceName         string            `pulumi:"workspaceName"`
}


type SqlPoolArgs struct {
	Collation             pulumi.StringPtrInput
	CreateMode            pulumi.StringPtrInput
	CreationDate          pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	MaxSizeBytes          pulumi.Float64PtrInput
	ProvisioningState     pulumi.StringPtrInput
	RecoverableDatabaseId pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	RestorePointInTime    pulumi.StringPtrInput
	Sku                   SkuPtrInput
	SourceDatabaseId      pulumi.StringPtrInput
	SqlPoolName           pulumi.StringPtrInput
	Status                pulumi.StringPtrInput
	StorageAccountType    pulumi.StringPtrInput
	Tags                  pulumi.StringMapInput
	WorkspaceName         pulumi.StringInput
}

func (SqlPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolArgs)(nil)).Elem()
}

type SqlPoolInput interface {
	pulumi.Input

	ToSqlPoolOutput() SqlPoolOutput
	ToSqlPoolOutputWithContext(ctx context.Context) SqlPoolOutput
}

func (*SqlPool) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlPool)(nil)).Elem()
}

func (i *SqlPool) ToSqlPoolOutput() SqlPoolOutput {
	return i.ToSqlPoolOutputWithContext(context.Background())
}

func (i *SqlPool) ToSqlPoolOutputWithContext(ctx context.Context) SqlPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlPoolOutput)
}

type SqlPoolOutput struct{ *pulumi.OutputState }

func (SqlPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlPool)(nil)).Elem()
}

func (o SqlPoolOutput) ToSqlPoolOutput() SqlPoolOutput {
	return o
}

func (o SqlPoolOutput) ToSqlPoolOutputWithContext(ctx context.Context) SqlPoolOutput {
	return o
}

func (o SqlPoolOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlPool) pulumi.StringPtrOutput { return v.Collation }).(pulumi.StringPtrOutput)
}

func (o SqlPoolOutput) CreateMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlPool) pulumi.StringPtrOutput { return v.CreateMode }).(pulumi.StringPtrOutput)
}

func (o SqlPoolOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlPool) pulumi.StringPtrOutput { return v.CreationDate }).(pulumi.StringPtrOutput)
}

func (o SqlPoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlPool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SqlPoolOutput) MaxSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SqlPool) pulumi.Float64PtrOutput { return v.MaxSizeBytes }).(pulumi.Float64PtrOutput)
}

func (o SqlPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlPoolOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlPool) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o SqlPoolOutput) RecoverableDatabaseId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlPool) pulumi.StringPtrOutput { return v.RecoverableDatabaseId }).(pulumi.StringPtrOutput)
}

func (o SqlPoolOutput) RestorePointInTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlPool) pulumi.StringPtrOutput { return v.RestorePointInTime }).(pulumi.StringPtrOutput)
}

func (o SqlPoolOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *SqlPool) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o SqlPoolOutput) SourceDatabaseId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlPool) pulumi.StringPtrOutput { return v.SourceDatabaseId }).(pulumi.StringPtrOutput)
}

func (o SqlPoolOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlPool) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o SqlPoolOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlPool) pulumi.StringPtrOutput { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

func (o SqlPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SqlPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlPoolOutput{})
}
