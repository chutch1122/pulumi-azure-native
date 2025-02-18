


package databoxedge

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Order struct {
	pulumi.CustomResourceState

	ContactInformation   ContactDetailsResponseOutput    `pulumi:"contactInformation"`
	CurrentStatus        OrderStatusResponseOutput       `pulumi:"currentStatus"`
	DeliveryTrackingInfo TrackingInfoResponseArrayOutput `pulumi:"deliveryTrackingInfo"`
	Name                 pulumi.StringOutput             `pulumi:"name"`
	OrderHistory         OrderStatusResponseArrayOutput  `pulumi:"orderHistory"`
	ReturnTrackingInfo   TrackingInfoResponseArrayOutput `pulumi:"returnTrackingInfo"`
	SerialNumber         pulumi.StringOutput             `pulumi:"serialNumber"`
	ShipmentType         pulumi.StringPtrOutput          `pulumi:"shipmentType"`
	ShippingAddress      AddressResponsePtrOutput        `pulumi:"shippingAddress"`
	Type                 pulumi.StringOutput             `pulumi:"type"`
}


func NewOrder(ctx *pulumi.Context,
	name string, args *OrderArgs, opts ...pulumi.ResourceOption) (*Order, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContactInformation == nil {
		return nil, errors.New("invalid value for required argument 'ContactInformation'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:Order"),
		},
	})
	opts = append(opts, aliases)
	var resource Order
	err := ctx.RegisterResource("azure-native:databoxedge:Order", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOrder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrderState, opts ...pulumi.ResourceOption) (*Order, error) {
	var resource Order
	err := ctx.ReadResource("azure-native:databoxedge:Order", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type orderState struct {
}

type OrderState struct {
}

func (OrderState) ElementType() reflect.Type {
	return reflect.TypeOf((*orderState)(nil)).Elem()
}

type orderArgs struct {
	ContactInformation ContactDetails `pulumi:"contactInformation"`
	DeviceName         string         `pulumi:"deviceName"`
	ResourceGroupName  string         `pulumi:"resourceGroupName"`
	ShipmentType       *string        `pulumi:"shipmentType"`
	ShippingAddress    *Address       `pulumi:"shippingAddress"`
}


type OrderArgs struct {
	ContactInformation ContactDetailsInput
	DeviceName         pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	ShipmentType       pulumi.StringPtrInput
	ShippingAddress    AddressPtrInput
}

func (OrderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orderArgs)(nil)).Elem()
}

type OrderInput interface {
	pulumi.Input

	ToOrderOutput() OrderOutput
	ToOrderOutputWithContext(ctx context.Context) OrderOutput
}

func (*Order) ElementType() reflect.Type {
	return reflect.TypeOf((**Order)(nil)).Elem()
}

func (i *Order) ToOrderOutput() OrderOutput {
	return i.ToOrderOutputWithContext(context.Background())
}

func (i *Order) ToOrderOutputWithContext(ctx context.Context) OrderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderOutput)
}

type OrderOutput struct{ *pulumi.OutputState }

func (OrderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Order)(nil)).Elem()
}

func (o OrderOutput) ToOrderOutput() OrderOutput {
	return o
}

func (o OrderOutput) ToOrderOutputWithContext(ctx context.Context) OrderOutput {
	return o
}

func (o OrderOutput) ContactInformation() ContactDetailsResponseOutput {
	return o.ApplyT(func(v *Order) ContactDetailsResponseOutput { return v.ContactInformation }).(ContactDetailsResponseOutput)
}

func (o OrderOutput) CurrentStatus() OrderStatusResponseOutput {
	return o.ApplyT(func(v *Order) OrderStatusResponseOutput { return v.CurrentStatus }).(OrderStatusResponseOutput)
}

func (o OrderOutput) DeliveryTrackingInfo() TrackingInfoResponseArrayOutput {
	return o.ApplyT(func(v *Order) TrackingInfoResponseArrayOutput { return v.DeliveryTrackingInfo }).(TrackingInfoResponseArrayOutput)
}

func (o OrderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Order) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OrderOutput) OrderHistory() OrderStatusResponseArrayOutput {
	return o.ApplyT(func(v *Order) OrderStatusResponseArrayOutput { return v.OrderHistory }).(OrderStatusResponseArrayOutput)
}

func (o OrderOutput) ReturnTrackingInfo() TrackingInfoResponseArrayOutput {
	return o.ApplyT(func(v *Order) TrackingInfoResponseArrayOutput { return v.ReturnTrackingInfo }).(TrackingInfoResponseArrayOutput)
}

func (o OrderOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *Order) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

func (o OrderOutput) ShipmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Order) pulumi.StringPtrOutput { return v.ShipmentType }).(pulumi.StringPtrOutput)
}

func (o OrderOutput) ShippingAddress() AddressResponsePtrOutput {
	return o.ApplyT(func(v *Order) AddressResponsePtrOutput { return v.ShippingAddress }).(AddressResponsePtrOutput)
}

func (o OrderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Order) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OrderOutput{})
}
