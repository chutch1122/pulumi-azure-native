


package cache

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Database struct {
	pulumi.CustomResourceState

	ClientProtocol    pulumi.StringPtrOutput       `pulumi:"clientProtocol"`
	ClusteringPolicy  pulumi.StringPtrOutput       `pulumi:"clusteringPolicy"`
	EvictionPolicy    pulumi.StringPtrOutput       `pulumi:"evictionPolicy"`
	Modules           ModuleResponseArrayOutput    `pulumi:"modules"`
	Name              pulumi.StringOutput          `pulumi:"name"`
	Persistence       PersistenceResponsePtrOutput `pulumi:"persistence"`
	Port              pulumi.IntPtrOutput          `pulumi:"port"`
	ProvisioningState pulumi.StringOutput          `pulumi:"provisioningState"`
	ResourceState     pulumi.StringOutput          `pulumi:"resourceState"`
	Type              pulumi.StringOutput          `pulumi:"type"`
}


func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cache/v20201001preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210201preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210301:Database"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210801:Database"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20220101:Database"),
		},
	})
	opts = append(opts, aliases)
	var resource Database
	err := ctx.RegisterResource("azure-native:cache:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("azure-native:cache:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type databaseState struct {
}

type DatabaseState struct {
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	ClientProtocol    *string      `pulumi:"clientProtocol"`
	ClusterName       string       `pulumi:"clusterName"`
	ClusteringPolicy  *string      `pulumi:"clusteringPolicy"`
	DatabaseName      *string      `pulumi:"databaseName"`
	EvictionPolicy    *string      `pulumi:"evictionPolicy"`
	Modules           []Module     `pulumi:"modules"`
	Persistence       *Persistence `pulumi:"persistence"`
	Port              *int         `pulumi:"port"`
	ResourceGroupName string       `pulumi:"resourceGroupName"`
}


type DatabaseArgs struct {
	ClientProtocol    pulumi.StringPtrInput
	ClusterName       pulumi.StringInput
	ClusteringPolicy  pulumi.StringPtrInput
	DatabaseName      pulumi.StringPtrInput
	EvictionPolicy    pulumi.StringPtrInput
	Modules           ModuleArrayInput
	Persistence       PersistencePtrInput
	Port              pulumi.IntPtrInput
	ResourceGroupName pulumi.StringInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

type DatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

func (o DatabaseOutput) ClientProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.ClientProtocol }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) ClusteringPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.ClusteringPolicy }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) EvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.EvictionPolicy }).(pulumi.StringPtrOutput)
}

func (o DatabaseOutput) Modules() ModuleResponseArrayOutput {
	return o.ApplyT(func(v *Database) ModuleResponseArrayOutput { return v.Modules }).(ModuleResponseArrayOutput)
}

func (o DatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseOutput) Persistence() PersistenceResponsePtrOutput {
	return o.ApplyT(func(v *Database) PersistenceResponsePtrOutput { return v.Persistence }).(PersistenceResponsePtrOutput)
}

func (o DatabaseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

func (o DatabaseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DatabaseOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

func (o DatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseOutput{})
}
