


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ADLSGen2StorageAccountDataSet struct {
	pulumi.CustomResourceState

	DataSetId                pulumi.StringOutput                           `pulumi:"dataSetId"`
	Kind                     pulumi.StringOutput                           `pulumi:"kind"`
	Location                 pulumi.StringOutput                           `pulumi:"location"`
	Name                     pulumi.StringOutput                           `pulumi:"name"`
	Paths                    ADLSGen2StorageAccountPathResponseArrayOutput `pulumi:"paths"`
	StorageAccountResourceId pulumi.StringOutput                           `pulumi:"storageAccountResourceId"`
	SystemData               SystemDataResponseOutput                      `pulumi:"systemData"`
	Type                     pulumi.StringOutput                           `pulumi:"type"`
}


func NewADLSGen2StorageAccountDataSet(ctx *pulumi.Context,
	name string, args *ADLSGen2StorageAccountDataSetArgs, opts ...pulumi.ResourceOption) (*ADLSGen2StorageAccountDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.Paths == nil {
		return nil, errors.New("invalid value for required argument 'Paths'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	if args.StorageAccountResourceId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountResourceId'")
	}
	args.Kind = pulumi.String("AdlsGen2StorageAccount")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare:ADLSGen2StorageAccountDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:ADLSGen2StorageAccountDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:ADLSGen2StorageAccountDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:ADLSGen2StorageAccountDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:ADLSGen2StorageAccountDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource ADLSGen2StorageAccountDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20201001preview:ADLSGen2StorageAccountDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetADLSGen2StorageAccountDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADLSGen2StorageAccountDataSetState, opts ...pulumi.ResourceOption) (*ADLSGen2StorageAccountDataSet, error) {
	var resource ADLSGen2StorageAccountDataSet
	err := ctx.ReadResource("azure-native:datashare/v20201001preview:ADLSGen2StorageAccountDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type adlsgen2StorageAccountDataSetState struct {
}

type ADLSGen2StorageAccountDataSetState struct {
}

func (ADLSGen2StorageAccountDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2StorageAccountDataSetState)(nil)).Elem()
}

type adlsgen2StorageAccountDataSetArgs struct {
	AccountName              string                       `pulumi:"accountName"`
	DataSetName              *string                      `pulumi:"dataSetName"`
	Kind                     string                       `pulumi:"kind"`
	Paths                    []ADLSGen2StorageAccountPath `pulumi:"paths"`
	ResourceGroupName        string                       `pulumi:"resourceGroupName"`
	ShareName                string                       `pulumi:"shareName"`
	StorageAccountResourceId string                       `pulumi:"storageAccountResourceId"`
}


type ADLSGen2StorageAccountDataSetArgs struct {
	AccountName              pulumi.StringInput
	DataSetName              pulumi.StringPtrInput
	Kind                     pulumi.StringInput
	Paths                    ADLSGen2StorageAccountPathArrayInput
	ResourceGroupName        pulumi.StringInput
	ShareName                pulumi.StringInput
	StorageAccountResourceId pulumi.StringInput
}

func (ADLSGen2StorageAccountDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2StorageAccountDataSetArgs)(nil)).Elem()
}

type ADLSGen2StorageAccountDataSetInput interface {
	pulumi.Input

	ToADLSGen2StorageAccountDataSetOutput() ADLSGen2StorageAccountDataSetOutput
	ToADLSGen2StorageAccountDataSetOutputWithContext(ctx context.Context) ADLSGen2StorageAccountDataSetOutput
}

func (*ADLSGen2StorageAccountDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2StorageAccountDataSet)(nil)).Elem()
}

func (i *ADLSGen2StorageAccountDataSet) ToADLSGen2StorageAccountDataSetOutput() ADLSGen2StorageAccountDataSetOutput {
	return i.ToADLSGen2StorageAccountDataSetOutputWithContext(context.Background())
}

func (i *ADLSGen2StorageAccountDataSet) ToADLSGen2StorageAccountDataSetOutputWithContext(ctx context.Context) ADLSGen2StorageAccountDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADLSGen2StorageAccountDataSetOutput)
}

type ADLSGen2StorageAccountDataSetOutput struct{ *pulumi.OutputState }

func (ADLSGen2StorageAccountDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2StorageAccountDataSet)(nil)).Elem()
}

func (o ADLSGen2StorageAccountDataSetOutput) ToADLSGen2StorageAccountDataSetOutput() ADLSGen2StorageAccountDataSetOutput {
	return o
}

func (o ADLSGen2StorageAccountDataSetOutput) ToADLSGen2StorageAccountDataSetOutputWithContext(ctx context.Context) ADLSGen2StorageAccountDataSetOutput {
	return o
}

func (o ADLSGen2StorageAccountDataSetOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2StorageAccountDataSet) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

func (o ADLSGen2StorageAccountDataSetOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2StorageAccountDataSet) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ADLSGen2StorageAccountDataSetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2StorageAccountDataSet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ADLSGen2StorageAccountDataSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2StorageAccountDataSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ADLSGen2StorageAccountDataSetOutput) Paths() ADLSGen2StorageAccountPathResponseArrayOutput {
	return o.ApplyT(func(v *ADLSGen2StorageAccountDataSet) ADLSGen2StorageAccountPathResponseArrayOutput { return v.Paths }).(ADLSGen2StorageAccountPathResponseArrayOutput)
}

func (o ADLSGen2StorageAccountDataSetOutput) StorageAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2StorageAccountDataSet) pulumi.StringOutput { return v.StorageAccountResourceId }).(pulumi.StringOutput)
}

func (o ADLSGen2StorageAccountDataSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ADLSGen2StorageAccountDataSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ADLSGen2StorageAccountDataSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2StorageAccountDataSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ADLSGen2StorageAccountDataSetOutput{})
}
