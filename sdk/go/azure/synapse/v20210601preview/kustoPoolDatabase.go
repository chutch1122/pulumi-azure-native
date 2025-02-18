


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type KustoPoolDatabase struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringOutput      `pulumi:"kind"`
	Location   pulumi.StringPtrOutput   `pulumi:"location"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewKustoPoolDatabase(ctx *pulumi.Context,
	name string, args *KustoPoolDatabaseArgs, opts ...pulumi.ResourceOption) (*KustoPoolDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.KustoPoolName == nil {
		return nil, errors.New("invalid value for required argument 'KustoPoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:KustoPoolDatabase"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:KustoPoolDatabase"),
		},
	})
	opts = append(opts, aliases)
	var resource KustoPoolDatabase
	err := ctx.RegisterResource("azure-native:synapse/v20210601preview:KustoPoolDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKustoPoolDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoPoolDatabaseState, opts ...pulumi.ResourceOption) (*KustoPoolDatabase, error) {
	var resource KustoPoolDatabase
	err := ctx.ReadResource("azure-native:synapse/v20210601preview:KustoPoolDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type kustoPoolDatabaseState struct {
}

type KustoPoolDatabaseState struct {
}

func (KustoPoolDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoPoolDatabaseState)(nil)).Elem()
}

type kustoPoolDatabaseArgs struct {
	DatabaseName      *string `pulumi:"databaseName"`
	Kind              string  `pulumi:"kind"`
	KustoPoolName     string  `pulumi:"kustoPoolName"`
	Location          *string `pulumi:"location"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


type KustoPoolDatabaseArgs struct {
	DatabaseName      pulumi.StringPtrInput
	Kind              pulumi.StringInput
	KustoPoolName     pulumi.StringInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (KustoPoolDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoPoolDatabaseArgs)(nil)).Elem()
}

type KustoPoolDatabaseInput interface {
	pulumi.Input

	ToKustoPoolDatabaseOutput() KustoPoolDatabaseOutput
	ToKustoPoolDatabaseOutputWithContext(ctx context.Context) KustoPoolDatabaseOutput
}

func (*KustoPoolDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoPoolDatabase)(nil)).Elem()
}

func (i *KustoPoolDatabase) ToKustoPoolDatabaseOutput() KustoPoolDatabaseOutput {
	return i.ToKustoPoolDatabaseOutputWithContext(context.Background())
}

func (i *KustoPoolDatabase) ToKustoPoolDatabaseOutputWithContext(ctx context.Context) KustoPoolDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoPoolDatabaseOutput)
}

type KustoPoolDatabaseOutput struct{ *pulumi.OutputState }

func (KustoPoolDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoPoolDatabase)(nil)).Elem()
}

func (o KustoPoolDatabaseOutput) ToKustoPoolDatabaseOutput() KustoPoolDatabaseOutput {
	return o
}

func (o KustoPoolDatabaseOutput) ToKustoPoolDatabaseOutputWithContext(ctx context.Context) KustoPoolDatabaseOutput {
	return o
}

func (o KustoPoolDatabaseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolDatabase) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o KustoPoolDatabaseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KustoPoolDatabase) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o KustoPoolDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o KustoPoolDatabaseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *KustoPoolDatabase) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o KustoPoolDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KustoPoolDatabaseOutput{})
}
