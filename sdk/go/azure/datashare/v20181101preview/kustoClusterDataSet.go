


package v20181101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KustoClusterDataSet struct {
	pulumi.CustomResourceState

	DataSetId              pulumi.StringOutput `pulumi:"dataSetId"`
	Kind                   pulumi.StringOutput `pulumi:"kind"`
	KustoClusterResourceId pulumi.StringOutput `pulumi:"kustoClusterResourceId"`
	Location               pulumi.StringOutput `pulumi:"location"`
	Name                   pulumi.StringOutput `pulumi:"name"`
	ProvisioningState      pulumi.StringOutput `pulumi:"provisioningState"`
	Type                   pulumi.StringOutput `pulumi:"type"`
}


func NewKustoClusterDataSet(ctx *pulumi.Context,
	name string, args *KustoClusterDataSetArgs, opts ...pulumi.ResourceOption) (*KustoClusterDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.KustoClusterResourceId == nil {
		return nil, errors.New("invalid value for required argument 'KustoClusterResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	args.Kind = pulumi.String("KustoCluster")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:KustoClusterDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:KustoClusterDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:KustoClusterDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:KustoClusterDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:KustoClusterDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource KustoClusterDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20181101preview:KustoClusterDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetKustoClusterDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoClusterDataSetState, opts ...pulumi.ResourceOption) (*KustoClusterDataSet, error) {
	var resource KustoClusterDataSet
	err := ctx.ReadResource("azure-native:datashare/v20181101preview:KustoClusterDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type kustoClusterDataSetState struct {
}

type KustoClusterDataSetState struct {
}

func (KustoClusterDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoClusterDataSetState)(nil)).Elem()
}

type kustoClusterDataSetArgs struct {
	AccountName            string  `pulumi:"accountName"`
	DataSetName            *string `pulumi:"dataSetName"`
	Kind                   string  `pulumi:"kind"`
	KustoClusterResourceId string  `pulumi:"kustoClusterResourceId"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
	ShareName              string  `pulumi:"shareName"`
}


type KustoClusterDataSetArgs struct {
	AccountName            pulumi.StringInput
	DataSetName            pulumi.StringPtrInput
	Kind                   pulumi.StringInput
	KustoClusterResourceId pulumi.StringInput
	ResourceGroupName      pulumi.StringInput
	ShareName              pulumi.StringInput
}

func (KustoClusterDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoClusterDataSetArgs)(nil)).Elem()
}

type KustoClusterDataSetInput interface {
	pulumi.Input

	ToKustoClusterDataSetOutput() KustoClusterDataSetOutput
	ToKustoClusterDataSetOutputWithContext(ctx context.Context) KustoClusterDataSetOutput
}

func (*KustoClusterDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoClusterDataSet)(nil)).Elem()
}

func (i *KustoClusterDataSet) ToKustoClusterDataSetOutput() KustoClusterDataSetOutput {
	return i.ToKustoClusterDataSetOutputWithContext(context.Background())
}

func (i *KustoClusterDataSet) ToKustoClusterDataSetOutputWithContext(ctx context.Context) KustoClusterDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoClusterDataSetOutput)
}

type KustoClusterDataSetOutput struct{ *pulumi.OutputState }

func (KustoClusterDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoClusterDataSet)(nil)).Elem()
}

func (o KustoClusterDataSetOutput) ToKustoClusterDataSetOutput() KustoClusterDataSetOutput {
	return o
}

func (o KustoClusterDataSetOutput) ToKustoClusterDataSetOutputWithContext(ctx context.Context) KustoClusterDataSetOutput {
	return o
}

func (o KustoClusterDataSetOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSet) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

func (o KustoClusterDataSetOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSet) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o KustoClusterDataSetOutput) KustoClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSet) pulumi.StringOutput { return v.KustoClusterResourceId }).(pulumi.StringOutput)
}

func (o KustoClusterDataSetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o KustoClusterDataSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o KustoClusterDataSetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSet) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o KustoClusterDataSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoClusterDataSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KustoClusterDataSetOutput{})
}
