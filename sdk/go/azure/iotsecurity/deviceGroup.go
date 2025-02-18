


package iotsecurity

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DeviceGroup struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewDeviceGroup(ctx *pulumi.Context,
	name string, args *DeviceGroupArgs, opts ...pulumi.ResourceOption) (*DeviceGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IotDefenderLocation == nil {
		return nil, errors.New("invalid value for required argument 'IotDefenderLocation'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:iotsecurity/v20210201preview:DeviceGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource DeviceGroup
	err := ctx.RegisterResource("azure-native:iotsecurity:DeviceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDeviceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceGroupState, opts ...pulumi.ResourceOption) (*DeviceGroup, error) {
	var resource DeviceGroup
	err := ctx.ReadResource("azure-native:iotsecurity:DeviceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type deviceGroupState struct {
}

type DeviceGroupState struct {
}

func (DeviceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceGroupState)(nil)).Elem()
}

type deviceGroupArgs struct {
	DeviceGroupName     *string `pulumi:"deviceGroupName"`
	IotDefenderLocation string  `pulumi:"iotDefenderLocation"`
}


type DeviceGroupArgs struct {
	DeviceGroupName     pulumi.StringPtrInput
	IotDefenderLocation pulumi.StringInput
}

func (DeviceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceGroupArgs)(nil)).Elem()
}

type DeviceGroupInput interface {
	pulumi.Input

	ToDeviceGroupOutput() DeviceGroupOutput
	ToDeviceGroupOutputWithContext(ctx context.Context) DeviceGroupOutput
}

func (*DeviceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceGroup)(nil)).Elem()
}

func (i *DeviceGroup) ToDeviceGroupOutput() DeviceGroupOutput {
	return i.ToDeviceGroupOutputWithContext(context.Background())
}

func (i *DeviceGroup) ToDeviceGroupOutputWithContext(ctx context.Context) DeviceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceGroupOutput)
}

type DeviceGroupOutput struct{ *pulumi.OutputState }

func (DeviceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceGroup)(nil)).Elem()
}

func (o DeviceGroupOutput) ToDeviceGroupOutput() DeviceGroupOutput {
	return o
}

func (o DeviceGroupOutput) ToDeviceGroupOutputWithContext(ctx context.Context) DeviceGroupOutput {
	return o
}

func (o DeviceGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeviceGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DeviceGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DeviceGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DeviceGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DeviceGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeviceGroupOutput{})
}
