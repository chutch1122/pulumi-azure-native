


package v20210401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataCollectionRule struct {
	pulumi.CustomResourceState

	DataFlows         DataFlowResponseArrayOutput                        `pulumi:"dataFlows"`
	DataSources       DataCollectionRuleResponseDataSourcesPtrOutput     `pulumi:"dataSources"`
	Description       pulumi.StringPtrOutput                             `pulumi:"description"`
	Destinations      DataCollectionRuleResponseDestinationsPtrOutput    `pulumi:"destinations"`
	Etag              pulumi.StringOutput                                `pulumi:"etag"`
	ImmutableId       pulumi.StringOutput                                `pulumi:"immutableId"`
	Kind              pulumi.StringPtrOutput                             `pulumi:"kind"`
	Location          pulumi.StringOutput                                `pulumi:"location"`
	Name              pulumi.StringOutput                                `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                                `pulumi:"provisioningState"`
	SystemData        DataCollectionRuleResourceResponseSystemDataOutput `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                             `pulumi:"tags"`
	Type              pulumi.StringOutput                                `pulumi:"type"`
}


func NewDataCollectionRule(ctx *pulumi.Context,
	name string, args *DataCollectionRuleArgs, opts ...pulumi.ResourceOption) (*DataCollectionRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:DataCollectionRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20191101preview:DataCollectionRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210901preview:DataCollectionRule"),
		},
	})
	opts = append(opts, aliases)
	var resource DataCollectionRule
	err := ctx.RegisterResource("azure-native:insights/v20210401:DataCollectionRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataCollectionRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataCollectionRuleState, opts ...pulumi.ResourceOption) (*DataCollectionRule, error) {
	var resource DataCollectionRule
	err := ctx.ReadResource("azure-native:insights/v20210401:DataCollectionRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dataCollectionRuleState struct {
}

type DataCollectionRuleState struct {
}

func (DataCollectionRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCollectionRuleState)(nil)).Elem()
}

type dataCollectionRuleArgs struct {
	DataCollectionRuleName *string                         `pulumi:"dataCollectionRuleName"`
	DataFlows              []DataFlow                      `pulumi:"dataFlows"`
	DataSources            *DataCollectionRuleDataSources  `pulumi:"dataSources"`
	Description            *string                         `pulumi:"description"`
	Destinations           *DataCollectionRuleDestinations `pulumi:"destinations"`
	Kind                   *string                         `pulumi:"kind"`
	Location               *string                         `pulumi:"location"`
	ResourceGroupName      string                          `pulumi:"resourceGroupName"`
	Tags                   map[string]string               `pulumi:"tags"`
}


type DataCollectionRuleArgs struct {
	DataCollectionRuleName pulumi.StringPtrInput
	DataFlows              DataFlowArrayInput
	DataSources            DataCollectionRuleDataSourcesPtrInput
	Description            pulumi.StringPtrInput
	Destinations           DataCollectionRuleDestinationsPtrInput
	Kind                   pulumi.StringPtrInput
	Location               pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
}

func (DataCollectionRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCollectionRuleArgs)(nil)).Elem()
}

type DataCollectionRuleInput interface {
	pulumi.Input

	ToDataCollectionRuleOutput() DataCollectionRuleOutput
	ToDataCollectionRuleOutputWithContext(ctx context.Context) DataCollectionRuleOutput
}

func (*DataCollectionRule) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRule)(nil)).Elem()
}

func (i *DataCollectionRule) ToDataCollectionRuleOutput() DataCollectionRuleOutput {
	return i.ToDataCollectionRuleOutputWithContext(context.Background())
}

func (i *DataCollectionRule) ToDataCollectionRuleOutputWithContext(ctx context.Context) DataCollectionRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectionRuleOutput)
}

type DataCollectionRuleOutput struct{ *pulumi.OutputState }

func (DataCollectionRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollectionRule)(nil)).Elem()
}

func (o DataCollectionRuleOutput) ToDataCollectionRuleOutput() DataCollectionRuleOutput {
	return o
}

func (o DataCollectionRuleOutput) ToDataCollectionRuleOutputWithContext(ctx context.Context) DataCollectionRuleOutput {
	return o
}

func (o DataCollectionRuleOutput) DataFlows() DataFlowResponseArrayOutput {
	return o.ApplyT(func(v *DataCollectionRule) DataFlowResponseArrayOutput { return v.DataFlows }).(DataFlowResponseArrayOutput)
}

func (o DataCollectionRuleOutput) DataSources() DataCollectionRuleResponseDataSourcesPtrOutput {
	return o.ApplyT(func(v *DataCollectionRule) DataCollectionRuleResponseDataSourcesPtrOutput { return v.DataSources }).(DataCollectionRuleResponseDataSourcesPtrOutput)
}

func (o DataCollectionRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataCollectionRuleOutput) Destinations() DataCollectionRuleResponseDestinationsPtrOutput {
	return o.ApplyT(func(v *DataCollectionRule) DataCollectionRuleResponseDestinationsPtrOutput { return v.Destinations }).(DataCollectionRuleResponseDestinationsPtrOutput)
}

func (o DataCollectionRuleOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DataCollectionRuleOutput) ImmutableId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringOutput { return v.ImmutableId }).(pulumi.StringOutput)
}

func (o DataCollectionRuleOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o DataCollectionRuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DataCollectionRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DataCollectionRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DataCollectionRuleOutput) SystemData() DataCollectionRuleResourceResponseSystemDataOutput {
	return o.ApplyT(func(v *DataCollectionRule) DataCollectionRuleResourceResponseSystemDataOutput { return v.SystemData }).(DataCollectionRuleResourceResponseSystemDataOutput)
}

func (o DataCollectionRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DataCollectionRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCollectionRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataCollectionRuleOutput{})
}
