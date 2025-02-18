


package v20210630preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EventHubEventSource struct {
	pulumi.CustomResourceState

	ConsumerGroupName     pulumi.StringOutput             `pulumi:"consumerGroupName"`
	CreationTime          pulumi.StringOutput             `pulumi:"creationTime"`
	EventHubName          pulumi.StringOutput             `pulumi:"eventHubName"`
	EventSourceResourceId pulumi.StringOutput             `pulumi:"eventSourceResourceId"`
	KeyName               pulumi.StringOutput             `pulumi:"keyName"`
	Kind                  pulumi.StringOutput             `pulumi:"kind"`
	LocalTimestamp        LocalTimestampResponsePtrOutput `pulumi:"localTimestamp"`
	Location              pulumi.StringOutput             `pulumi:"location"`
	Name                  pulumi.StringOutput             `pulumi:"name"`
	ProvisioningState     pulumi.StringOutput             `pulumi:"provisioningState"`
	ServiceBusNamespace   pulumi.StringOutput             `pulumi:"serviceBusNamespace"`
	Tags                  pulumi.StringMapOutput          `pulumi:"tags"`
	Time                  pulumi.StringPtrOutput          `pulumi:"time"`
	TimestampPropertyName pulumi.StringPtrOutput          `pulumi:"timestampPropertyName"`
	Type                  pulumi.StringOutput             `pulumi:"type"`
}


func NewEventHubEventSource(ctx *pulumi.Context,
	name string, args *EventHubEventSourceArgs, opts ...pulumi.ResourceOption) (*EventHubEventSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsumerGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerGroupName'")
	}
	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	if args.EventHubName == nil {
		return nil, errors.New("invalid value for required argument 'EventHubName'")
	}
	if args.EventSourceResourceId == nil {
		return nil, errors.New("invalid value for required argument 'EventSourceResourceId'")
	}
	if args.KeyName == nil {
		return nil, errors.New("invalid value for required argument 'KeyName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceBusNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ServiceBusNamespace'")
	}
	if args.SharedAccessKey == nil {
		return nil, errors.New("invalid value for required argument 'SharedAccessKey'")
	}
	args.Kind = pulumi.String("Microsoft.EventHub")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:timeseriesinsights:EventHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20170228preview:EventHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20171115:EventHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20180815preview:EventHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20200515:EventHubEventSource"),
		},
		{
			Type: pulumi.String("azure-native:timeseriesinsights/v20210331preview:EventHubEventSource"),
		},
	})
	opts = append(opts, aliases)
	var resource EventHubEventSource
	err := ctx.RegisterResource("azure-native:timeseriesinsights/v20210630preview:EventHubEventSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEventHubEventSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubEventSourceState, opts ...pulumi.ResourceOption) (*EventHubEventSource, error) {
	var resource EventHubEventSource
	err := ctx.ReadResource("azure-native:timeseriesinsights/v20210630preview:EventHubEventSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type eventHubEventSourceState struct {
}

type EventHubEventSourceState struct {
}

func (EventHubEventSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubEventSourceState)(nil)).Elem()
}

type eventHubEventSourceArgs struct {
	ConsumerGroupName     string            `pulumi:"consumerGroupName"`
	EnvironmentName       string            `pulumi:"environmentName"`
	EventHubName          string            `pulumi:"eventHubName"`
	EventSourceName       *string           `pulumi:"eventSourceName"`
	EventSourceResourceId string            `pulumi:"eventSourceResourceId"`
	KeyName               string            `pulumi:"keyName"`
	Kind                  string            `pulumi:"kind"`
	LocalTimestamp        *LocalTimestamp   `pulumi:"localTimestamp"`
	Location              *string           `pulumi:"location"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	ServiceBusNamespace   string            `pulumi:"serviceBusNamespace"`
	SharedAccessKey       string            `pulumi:"sharedAccessKey"`
	Tags                  map[string]string `pulumi:"tags"`
	Time                  *string           `pulumi:"time"`
	TimestampPropertyName *string           `pulumi:"timestampPropertyName"`
	Type                  *string           `pulumi:"type"`
}


type EventHubEventSourceArgs struct {
	ConsumerGroupName     pulumi.StringInput
	EnvironmentName       pulumi.StringInput
	EventHubName          pulumi.StringInput
	EventSourceName       pulumi.StringPtrInput
	EventSourceResourceId pulumi.StringInput
	KeyName               pulumi.StringInput
	Kind                  pulumi.StringInput
	LocalTimestamp        LocalTimestampPtrInput
	Location              pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	ServiceBusNamespace   pulumi.StringInput
	SharedAccessKey       pulumi.StringInput
	Tags                  pulumi.StringMapInput
	Time                  pulumi.StringPtrInput
	TimestampPropertyName pulumi.StringPtrInput
	Type                  pulumi.StringPtrInput
}

func (EventHubEventSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubEventSourceArgs)(nil)).Elem()
}

type EventHubEventSourceInput interface {
	pulumi.Input

	ToEventHubEventSourceOutput() EventHubEventSourceOutput
	ToEventHubEventSourceOutputWithContext(ctx context.Context) EventHubEventSourceOutput
}

func (*EventHubEventSource) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubEventSource)(nil)).Elem()
}

func (i *EventHubEventSource) ToEventHubEventSourceOutput() EventHubEventSourceOutput {
	return i.ToEventHubEventSourceOutputWithContext(context.Background())
}

func (i *EventHubEventSource) ToEventHubEventSourceOutputWithContext(ctx context.Context) EventHubEventSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubEventSourceOutput)
}

type EventHubEventSourceOutput struct{ *pulumi.OutputState }

func (EventHubEventSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubEventSource)(nil)).Elem()
}

func (o EventHubEventSourceOutput) ToEventHubEventSourceOutput() EventHubEventSourceOutput {
	return o
}

func (o EventHubEventSourceOutput) ToEventHubEventSourceOutputWithContext(ctx context.Context) EventHubEventSourceOutput {
	return o
}

func (o EventHubEventSourceOutput) ConsumerGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubEventSource) pulumi.StringOutput { return v.ConsumerGroupName }).(pulumi.StringOutput)
}

func (o EventHubEventSourceOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubEventSource) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o EventHubEventSourceOutput) EventHubName() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubEventSource) pulumi.StringOutput { return v.EventHubName }).(pulumi.StringOutput)
}

func (o EventHubEventSourceOutput) EventSourceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubEventSource) pulumi.StringOutput { return v.EventSourceResourceId }).(pulumi.StringOutput)
}

func (o EventHubEventSourceOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubEventSource) pulumi.StringOutput { return v.KeyName }).(pulumi.StringOutput)
}

func (o EventHubEventSourceOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubEventSource) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o EventHubEventSourceOutput) LocalTimestamp() LocalTimestampResponsePtrOutput {
	return o.ApplyT(func(v *EventHubEventSource) LocalTimestampResponsePtrOutput { return v.LocalTimestamp }).(LocalTimestampResponsePtrOutput)
}

func (o EventHubEventSourceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubEventSource) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o EventHubEventSourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubEventSource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EventHubEventSourceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubEventSource) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o EventHubEventSourceOutput) ServiceBusNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubEventSource) pulumi.StringOutput { return v.ServiceBusNamespace }).(pulumi.StringOutput)
}

func (o EventHubEventSourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EventHubEventSource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o EventHubEventSourceOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHubEventSource) pulumi.StringPtrOutput { return v.Time }).(pulumi.StringPtrOutput)
}

func (o EventHubEventSourceOutput) TimestampPropertyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHubEventSource) pulumi.StringPtrOutput { return v.TimestampPropertyName }).(pulumi.StringPtrOutput)
}

func (o EventHubEventSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubEventSource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EventHubEventSourceOutput{})
}
