


package v20191109

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EventGridDataConnection struct {
	pulumi.CustomResourceState

	ConsumerGroup            pulumi.StringOutput    `pulumi:"consumerGroup"`
	DataFormat               pulumi.StringOutput    `pulumi:"dataFormat"`
	EventHubResourceId       pulumi.StringOutput    `pulumi:"eventHubResourceId"`
	Kind                     pulumi.StringOutput    `pulumi:"kind"`
	Location                 pulumi.StringPtrOutput `pulumi:"location"`
	MappingRuleName          pulumi.StringPtrOutput `pulumi:"mappingRuleName"`
	Name                     pulumi.StringOutput    `pulumi:"name"`
	StorageAccountResourceId pulumi.StringOutput    `pulumi:"storageAccountResourceId"`
	TableName                pulumi.StringOutput    `pulumi:"tableName"`
	Type                     pulumi.StringOutput    `pulumi:"type"`
}


func NewEventGridDataConnection(ctx *pulumi.Context,
	name string, args *EventGridDataConnectionArgs, opts ...pulumi.ResourceOption) (*EventGridDataConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ConsumerGroup == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerGroup'")
	}
	if args.DataFormat == nil {
		return nil, errors.New("invalid value for required argument 'DataFormat'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.EventHubResourceId == nil {
		return nil, errors.New("invalid value for required argument 'EventHubResourceId'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccountResourceId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountResourceId'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	args.Kind = pulumi.String("EventGrid")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190121:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190515:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190907:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200918:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:EventGridDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220707:EventGridDataConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource EventGridDataConnection
	err := ctx.RegisterResource("azure-native:kusto/v20191109:EventGridDataConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEventGridDataConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventGridDataConnectionState, opts ...pulumi.ResourceOption) (*EventGridDataConnection, error) {
	var resource EventGridDataConnection
	err := ctx.ReadResource("azure-native:kusto/v20191109:EventGridDataConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type eventGridDataConnectionState struct {
}

type EventGridDataConnectionState struct {
}

func (EventGridDataConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventGridDataConnectionState)(nil)).Elem()
}

type eventGridDataConnectionArgs struct {
	ClusterName              string  `pulumi:"clusterName"`
	ConsumerGroup            string  `pulumi:"consumerGroup"`
	DataConnectionName       *string `pulumi:"dataConnectionName"`
	DataFormat               string  `pulumi:"dataFormat"`
	DatabaseName             string  `pulumi:"databaseName"`
	EventHubResourceId       string  `pulumi:"eventHubResourceId"`
	Kind                     string  `pulumi:"kind"`
	Location                 *string `pulumi:"location"`
	MappingRuleName          *string `pulumi:"mappingRuleName"`
	ResourceGroupName        string  `pulumi:"resourceGroupName"`
	StorageAccountResourceId string  `pulumi:"storageAccountResourceId"`
	TableName                string  `pulumi:"tableName"`
}


type EventGridDataConnectionArgs struct {
	ClusterName              pulumi.StringInput
	ConsumerGroup            pulumi.StringInput
	DataConnectionName       pulumi.StringPtrInput
	DataFormat               pulumi.StringInput
	DatabaseName             pulumi.StringInput
	EventHubResourceId       pulumi.StringInput
	Kind                     pulumi.StringInput
	Location                 pulumi.StringPtrInput
	MappingRuleName          pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	StorageAccountResourceId pulumi.StringInput
	TableName                pulumi.StringInput
}

func (EventGridDataConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventGridDataConnectionArgs)(nil)).Elem()
}

type EventGridDataConnectionInput interface {
	pulumi.Input

	ToEventGridDataConnectionOutput() EventGridDataConnectionOutput
	ToEventGridDataConnectionOutputWithContext(ctx context.Context) EventGridDataConnectionOutput
}

func (*EventGridDataConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**EventGridDataConnection)(nil)).Elem()
}

func (i *EventGridDataConnection) ToEventGridDataConnectionOutput() EventGridDataConnectionOutput {
	return i.ToEventGridDataConnectionOutputWithContext(context.Background())
}

func (i *EventGridDataConnection) ToEventGridDataConnectionOutputWithContext(ctx context.Context) EventGridDataConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGridDataConnectionOutput)
}

type EventGridDataConnectionOutput struct{ *pulumi.OutputState }

func (EventGridDataConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventGridDataConnection)(nil)).Elem()
}

func (o EventGridDataConnectionOutput) ToEventGridDataConnectionOutput() EventGridDataConnectionOutput {
	return o
}

func (o EventGridDataConnectionOutput) ToEventGridDataConnectionOutputWithContext(ctx context.Context) EventGridDataConnectionOutput {
	return o
}

func (o EventGridDataConnectionOutput) ConsumerGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringOutput { return v.ConsumerGroup }).(pulumi.StringOutput)
}

func (o EventGridDataConnectionOutput) DataFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringOutput { return v.DataFormat }).(pulumi.StringOutput)
}

func (o EventGridDataConnectionOutput) EventHubResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringOutput { return v.EventHubResourceId }).(pulumi.StringOutput)
}

func (o EventGridDataConnectionOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o EventGridDataConnectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o EventGridDataConnectionOutput) MappingRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringPtrOutput { return v.MappingRuleName }).(pulumi.StringPtrOutput)
}

func (o EventGridDataConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EventGridDataConnectionOutput) StorageAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringOutput { return v.StorageAccountResourceId }).(pulumi.StringOutput)
}

func (o EventGridDataConnectionOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

func (o EventGridDataConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EventGridDataConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EventGridDataConnectionOutput{})
}
