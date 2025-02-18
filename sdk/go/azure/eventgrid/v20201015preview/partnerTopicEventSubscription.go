


package v20201015preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PartnerTopicEventSubscription struct {
	pulumi.CustomResourceState

	DeadLetterDestination          StorageBlobDeadLetterDestinationResponsePtrOutput `pulumi:"deadLetterDestination"`
	DeadLetterWithResourceIdentity DeadLetterWithResourceIdentityResponsePtrOutput   `pulumi:"deadLetterWithResourceIdentity"`
	DeliveryWithResourceIdentity   DeliveryWithResourceIdentityResponsePtrOutput     `pulumi:"deliveryWithResourceIdentity"`
	Destination                    pulumi.AnyOutput                                  `pulumi:"destination"`
	EventDeliverySchema            pulumi.StringPtrOutput                            `pulumi:"eventDeliverySchema"`
	ExpirationTimeUtc              pulumi.StringPtrOutput                            `pulumi:"expirationTimeUtc"`
	Filter                         EventSubscriptionFilterResponsePtrOutput          `pulumi:"filter"`
	Labels                         pulumi.StringArrayOutput                          `pulumi:"labels"`
	Name                           pulumi.StringOutput                               `pulumi:"name"`
	ProvisioningState              pulumi.StringOutput                               `pulumi:"provisioningState"`
	RetryPolicy                    RetryPolicyResponsePtrOutput                      `pulumi:"retryPolicy"`
	SystemData                     SystemDataResponseOutput                          `pulumi:"systemData"`
	Topic                          pulumi.StringOutput                               `pulumi:"topic"`
	Type                           pulumi.StringOutput                               `pulumi:"type"`
}


func NewPartnerTopicEventSubscription(ctx *pulumi.Context,
	name string, args *PartnerTopicEventSubscriptionArgs, opts ...pulumi.ResourceOption) (*PartnerTopicEventSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PartnerTopicName == nil {
		return nil, errors.New("invalid value for required argument 'PartnerTopicName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.EventDeliverySchema) {
		args.EventDeliverySchema = pulumi.StringPtr("EventGridSchema")
	}
	if args.Filter != nil {
		args.Filter = args.Filter.ToEventSubscriptionFilterPtrOutput().ApplyT(func(v *EventSubscriptionFilter) *EventSubscriptionFilter { return v.Defaults() }).(EventSubscriptionFilterPtrOutput)
	}
	if args.RetryPolicy != nil {
		args.RetryPolicy = args.RetryPolicy.ToRetryPolicyPtrOutput().ApplyT(func(v *RetryPolicy) *RetryPolicy { return v.Defaults() }).(RetryPolicyPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:PartnerTopicEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:PartnerTopicEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20210601preview:PartnerTopicEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:PartnerTopicEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20220615:PartnerTopicEventSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource PartnerTopicEventSubscription
	err := ctx.RegisterResource("azure-native:eventgrid/v20201015preview:PartnerTopicEventSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPartnerTopicEventSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartnerTopicEventSubscriptionState, opts ...pulumi.ResourceOption) (*PartnerTopicEventSubscription, error) {
	var resource PartnerTopicEventSubscription
	err := ctx.ReadResource("azure-native:eventgrid/v20201015preview:PartnerTopicEventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type partnerTopicEventSubscriptionState struct {
}

type PartnerTopicEventSubscriptionState struct {
}

func (PartnerTopicEventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerTopicEventSubscriptionState)(nil)).Elem()
}

type partnerTopicEventSubscriptionArgs struct {
	DeadLetterDestination          *StorageBlobDeadLetterDestination `pulumi:"deadLetterDestination"`
	DeadLetterWithResourceIdentity *DeadLetterWithResourceIdentity   `pulumi:"deadLetterWithResourceIdentity"`
	DeliveryWithResourceIdentity   *DeliveryWithResourceIdentity     `pulumi:"deliveryWithResourceIdentity"`
	Destination                    interface{}                       `pulumi:"destination"`
	EventDeliverySchema            *string                           `pulumi:"eventDeliverySchema"`
	EventSubscriptionName          *string                           `pulumi:"eventSubscriptionName"`
	ExpirationTimeUtc              *string                           `pulumi:"expirationTimeUtc"`
	Filter                         *EventSubscriptionFilter          `pulumi:"filter"`
	Labels                         []string                          `pulumi:"labels"`
	PartnerTopicName               string                            `pulumi:"partnerTopicName"`
	ResourceGroupName              string                            `pulumi:"resourceGroupName"`
	RetryPolicy                    *RetryPolicy                      `pulumi:"retryPolicy"`
}


type PartnerTopicEventSubscriptionArgs struct {
	DeadLetterDestination          StorageBlobDeadLetterDestinationPtrInput
	DeadLetterWithResourceIdentity DeadLetterWithResourceIdentityPtrInput
	DeliveryWithResourceIdentity   DeliveryWithResourceIdentityPtrInput
	Destination                    pulumi.Input
	EventDeliverySchema            pulumi.StringPtrInput
	EventSubscriptionName          pulumi.StringPtrInput
	ExpirationTimeUtc              pulumi.StringPtrInput
	Filter                         EventSubscriptionFilterPtrInput
	Labels                         pulumi.StringArrayInput
	PartnerTopicName               pulumi.StringInput
	ResourceGroupName              pulumi.StringInput
	RetryPolicy                    RetryPolicyPtrInput
}

func (PartnerTopicEventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerTopicEventSubscriptionArgs)(nil)).Elem()
}

type PartnerTopicEventSubscriptionInput interface {
	pulumi.Input

	ToPartnerTopicEventSubscriptionOutput() PartnerTopicEventSubscriptionOutput
	ToPartnerTopicEventSubscriptionOutputWithContext(ctx context.Context) PartnerTopicEventSubscriptionOutput
}

func (*PartnerTopicEventSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerTopicEventSubscription)(nil)).Elem()
}

func (i *PartnerTopicEventSubscription) ToPartnerTopicEventSubscriptionOutput() PartnerTopicEventSubscriptionOutput {
	return i.ToPartnerTopicEventSubscriptionOutputWithContext(context.Background())
}

func (i *PartnerTopicEventSubscription) ToPartnerTopicEventSubscriptionOutputWithContext(ctx context.Context) PartnerTopicEventSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerTopicEventSubscriptionOutput)
}

type PartnerTopicEventSubscriptionOutput struct{ *pulumi.OutputState }

func (PartnerTopicEventSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerTopicEventSubscription)(nil)).Elem()
}

func (o PartnerTopicEventSubscriptionOutput) ToPartnerTopicEventSubscriptionOutput() PartnerTopicEventSubscriptionOutput {
	return o
}

func (o PartnerTopicEventSubscriptionOutput) ToPartnerTopicEventSubscriptionOutputWithContext(ctx context.Context) PartnerTopicEventSubscriptionOutput {
	return o
}

func (o PartnerTopicEventSubscriptionOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyT(func(v *PartnerTopicEventSubscription) StorageBlobDeadLetterDestinationResponsePtrOutput {
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

func (o PartnerTopicEventSubscriptionOutput) DeadLetterWithResourceIdentity() DeadLetterWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *PartnerTopicEventSubscription) DeadLetterWithResourceIdentityResponsePtrOutput {
		return v.DeadLetterWithResourceIdentity
	}).(DeadLetterWithResourceIdentityResponsePtrOutput)
}

func (o PartnerTopicEventSubscriptionOutput) DeliveryWithResourceIdentity() DeliveryWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *PartnerTopicEventSubscription) DeliveryWithResourceIdentityResponsePtrOutput {
		return v.DeliveryWithResourceIdentity
	}).(DeliveryWithResourceIdentityResponsePtrOutput)
}

func (o PartnerTopicEventSubscriptionOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v *PartnerTopicEventSubscription) pulumi.AnyOutput { return v.Destination }).(pulumi.AnyOutput)
}

func (o PartnerTopicEventSubscriptionOutput) EventDeliverySchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopicEventSubscription) pulumi.StringPtrOutput { return v.EventDeliverySchema }).(pulumi.StringPtrOutput)
}

func (o PartnerTopicEventSubscriptionOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopicEventSubscription) pulumi.StringPtrOutput { return v.ExpirationTimeUtc }).(pulumi.StringPtrOutput)
}

func (o PartnerTopicEventSubscriptionOutput) Filter() EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyT(func(v *PartnerTopicEventSubscription) EventSubscriptionFilterResponsePtrOutput { return v.Filter }).(EventSubscriptionFilterResponsePtrOutput)
}

func (o PartnerTopicEventSubscriptionOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PartnerTopicEventSubscription) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o PartnerTopicEventSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerTopicEventSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PartnerTopicEventSubscriptionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerTopicEventSubscription) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PartnerTopicEventSubscriptionOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v *PartnerTopicEventSubscription) RetryPolicyResponsePtrOutput { return v.RetryPolicy }).(RetryPolicyResponsePtrOutput)
}

func (o PartnerTopicEventSubscriptionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PartnerTopicEventSubscription) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PartnerTopicEventSubscriptionOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerTopicEventSubscription) pulumi.StringOutput { return v.Topic }).(pulumi.StringOutput)
}

func (o PartnerTopicEventSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerTopicEventSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PartnerTopicEventSubscriptionOutput{})
}
