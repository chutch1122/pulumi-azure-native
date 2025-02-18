


package datashare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SynapseWorkspaceSqlPoolTableDataSet struct {
	pulumi.CustomResourceState

	DataSetId                              pulumi.StringOutput      `pulumi:"dataSetId"`
	Kind                                   pulumi.StringOutput      `pulumi:"kind"`
	Name                                   pulumi.StringOutput      `pulumi:"name"`
	SynapseWorkspaceSqlPoolTableResourceId pulumi.StringOutput      `pulumi:"synapseWorkspaceSqlPoolTableResourceId"`
	SystemData                             SystemDataResponseOutput `pulumi:"systemData"`
	Type                                   pulumi.StringOutput      `pulumi:"type"`
}


func NewSynapseWorkspaceSqlPoolTableDataSet(ctx *pulumi.Context,
	name string, args *SynapseWorkspaceSqlPoolTableDataSetArgs, opts ...pulumi.ResourceOption) (*SynapseWorkspaceSqlPoolTableDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	if args.SynapseWorkspaceSqlPoolTableResourceId == nil {
		return nil, errors.New("invalid value for required argument 'SynapseWorkspaceSqlPoolTableResourceId'")
	}
	args.Kind = pulumi.String("SynapseWorkspaceSqlPoolTable")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:SynapseWorkspaceSqlPoolTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:SynapseWorkspaceSqlPoolTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:SynapseWorkspaceSqlPoolTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:SynapseWorkspaceSqlPoolTableDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:SynapseWorkspaceSqlPoolTableDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource SynapseWorkspaceSqlPoolTableDataSet
	err := ctx.RegisterResource("azure-native:datashare:SynapseWorkspaceSqlPoolTableDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSynapseWorkspaceSqlPoolTableDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SynapseWorkspaceSqlPoolTableDataSetState, opts ...pulumi.ResourceOption) (*SynapseWorkspaceSqlPoolTableDataSet, error) {
	var resource SynapseWorkspaceSqlPoolTableDataSet
	err := ctx.ReadResource("azure-native:datashare:SynapseWorkspaceSqlPoolTableDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type synapseWorkspaceSqlPoolTableDataSetState struct {
}

type SynapseWorkspaceSqlPoolTableDataSetState struct {
}

func (SynapseWorkspaceSqlPoolTableDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*synapseWorkspaceSqlPoolTableDataSetState)(nil)).Elem()
}

type synapseWorkspaceSqlPoolTableDataSetArgs struct {
	AccountName                            string  `pulumi:"accountName"`
	DataSetName                            *string `pulumi:"dataSetName"`
	Kind                                   string  `pulumi:"kind"`
	ResourceGroupName                      string  `pulumi:"resourceGroupName"`
	ShareName                              string  `pulumi:"shareName"`
	SynapseWorkspaceSqlPoolTableResourceId string  `pulumi:"synapseWorkspaceSqlPoolTableResourceId"`
}


type SynapseWorkspaceSqlPoolTableDataSetArgs struct {
	AccountName                            pulumi.StringInput
	DataSetName                            pulumi.StringPtrInput
	Kind                                   pulumi.StringInput
	ResourceGroupName                      pulumi.StringInput
	ShareName                              pulumi.StringInput
	SynapseWorkspaceSqlPoolTableResourceId pulumi.StringInput
}

func (SynapseWorkspaceSqlPoolTableDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*synapseWorkspaceSqlPoolTableDataSetArgs)(nil)).Elem()
}

type SynapseWorkspaceSqlPoolTableDataSetInput interface {
	pulumi.Input

	ToSynapseWorkspaceSqlPoolTableDataSetOutput() SynapseWorkspaceSqlPoolTableDataSetOutput
	ToSynapseWorkspaceSqlPoolTableDataSetOutputWithContext(ctx context.Context) SynapseWorkspaceSqlPoolTableDataSetOutput
}

func (*SynapseWorkspaceSqlPoolTableDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**SynapseWorkspaceSqlPoolTableDataSet)(nil)).Elem()
}

func (i *SynapseWorkspaceSqlPoolTableDataSet) ToSynapseWorkspaceSqlPoolTableDataSetOutput() SynapseWorkspaceSqlPoolTableDataSetOutput {
	return i.ToSynapseWorkspaceSqlPoolTableDataSetOutputWithContext(context.Background())
}

func (i *SynapseWorkspaceSqlPoolTableDataSet) ToSynapseWorkspaceSqlPoolTableDataSetOutputWithContext(ctx context.Context) SynapseWorkspaceSqlPoolTableDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynapseWorkspaceSqlPoolTableDataSetOutput)
}

type SynapseWorkspaceSqlPoolTableDataSetOutput struct{ *pulumi.OutputState }

func (SynapseWorkspaceSqlPoolTableDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SynapseWorkspaceSqlPoolTableDataSet)(nil)).Elem()
}

func (o SynapseWorkspaceSqlPoolTableDataSetOutput) ToSynapseWorkspaceSqlPoolTableDataSetOutput() SynapseWorkspaceSqlPoolTableDataSetOutput {
	return o
}

func (o SynapseWorkspaceSqlPoolTableDataSetOutput) ToSynapseWorkspaceSqlPoolTableDataSetOutputWithContext(ctx context.Context) SynapseWorkspaceSqlPoolTableDataSetOutput {
	return o
}

func (o SynapseWorkspaceSqlPoolTableDataSetOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *SynapseWorkspaceSqlPoolTableDataSet) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

func (o SynapseWorkspaceSqlPoolTableDataSetOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *SynapseWorkspaceSqlPoolTableDataSet) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o SynapseWorkspaceSqlPoolTableDataSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SynapseWorkspaceSqlPoolTableDataSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SynapseWorkspaceSqlPoolTableDataSetOutput) SynapseWorkspaceSqlPoolTableResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SynapseWorkspaceSqlPoolTableDataSet) pulumi.StringOutput {
		return v.SynapseWorkspaceSqlPoolTableResourceId
	}).(pulumi.StringOutput)
}

func (o SynapseWorkspaceSqlPoolTableDataSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SynapseWorkspaceSqlPoolTableDataSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SynapseWorkspaceSqlPoolTableDataSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SynapseWorkspaceSqlPoolTableDataSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SynapseWorkspaceSqlPoolTableDataSetOutput{})
}
