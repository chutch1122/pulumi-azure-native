


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Sensor struct {
	pulumi.CustomResourceState

	ConnectivityTime   pulumi.StringOutput      `pulumi:"connectivityTime"`
	DynamicLearning    pulumi.BoolOutput        `pulumi:"dynamicLearning"`
	LearningMode       pulumi.BoolOutput        `pulumi:"learningMode"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	SensorStatus       pulumi.StringOutput      `pulumi:"sensorStatus"`
	SensorType         pulumi.StringPtrOutput   `pulumi:"sensorType"`
	SensorVersion      pulumi.StringOutput      `pulumi:"sensorVersion"`
	SystemData         SystemDataResponseOutput `pulumi:"systemData"`
	TiAutomaticUpdates pulumi.BoolPtrOutput     `pulumi:"tiAutomaticUpdates"`
	TiStatus           pulumi.StringOutput      `pulumi:"tiStatus"`
	TiVersion          pulumi.StringOutput      `pulumi:"tiVersion"`
	Type               pulumi.StringOutput      `pulumi:"type"`
	Zone               pulumi.StringPtrOutput   `pulumi:"zone"`
}


func NewSensor(ctx *pulumi.Context,
	name string, args *SensorArgs, opts ...pulumi.ResourceOption) (*Sensor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:iotsecurity:Sensor"),
		},
	})
	opts = append(opts, aliases)
	var resource Sensor
	err := ctx.RegisterResource("azure-native:iotsecurity/v20210201preview:Sensor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSensor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SensorState, opts ...pulumi.ResourceOption) (*Sensor, error) {
	var resource Sensor
	err := ctx.ReadResource("azure-native:iotsecurity/v20210201preview:Sensor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sensorState struct {
}

type SensorState struct {
}

func (SensorState) ElementType() reflect.Type {
	return reflect.TypeOf((*sensorState)(nil)).Elem()
}

type sensorArgs struct {
	Scope              string  `pulumi:"scope"`
	SensorName         *string `pulumi:"sensorName"`
	SensorType         *string `pulumi:"sensorType"`
	TiAutomaticUpdates *bool   `pulumi:"tiAutomaticUpdates"`
	Zone               *string `pulumi:"zone"`
}


type SensorArgs struct {
	Scope              pulumi.StringInput
	SensorName         pulumi.StringPtrInput
	SensorType         pulumi.StringPtrInput
	TiAutomaticUpdates pulumi.BoolPtrInput
	Zone               pulumi.StringPtrInput
}

func (SensorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sensorArgs)(nil)).Elem()
}

type SensorInput interface {
	pulumi.Input

	ToSensorOutput() SensorOutput
	ToSensorOutputWithContext(ctx context.Context) SensorOutput
}

func (*Sensor) ElementType() reflect.Type {
	return reflect.TypeOf((**Sensor)(nil)).Elem()
}

func (i *Sensor) ToSensorOutput() SensorOutput {
	return i.ToSensorOutputWithContext(context.Background())
}

func (i *Sensor) ToSensorOutputWithContext(ctx context.Context) SensorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SensorOutput)
}

type SensorOutput struct{ *pulumi.OutputState }

func (SensorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sensor)(nil)).Elem()
}

func (o SensorOutput) ToSensorOutput() SensorOutput {
	return o
}

func (o SensorOutput) ToSensorOutputWithContext(ctx context.Context) SensorOutput {
	return o
}

func (o SensorOutput) ConnectivityTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Sensor) pulumi.StringOutput { return v.ConnectivityTime }).(pulumi.StringOutput)
}

func (o SensorOutput) DynamicLearning() pulumi.BoolOutput {
	return o.ApplyT(func(v *Sensor) pulumi.BoolOutput { return v.DynamicLearning }).(pulumi.BoolOutput)
}

func (o SensorOutput) LearningMode() pulumi.BoolOutput {
	return o.ApplyT(func(v *Sensor) pulumi.BoolOutput { return v.LearningMode }).(pulumi.BoolOutput)
}

func (o SensorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Sensor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SensorOutput) SensorStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Sensor) pulumi.StringOutput { return v.SensorStatus }).(pulumi.StringOutput)
}

func (o SensorOutput) SensorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sensor) pulumi.StringPtrOutput { return v.SensorType }).(pulumi.StringPtrOutput)
}

func (o SensorOutput) SensorVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Sensor) pulumi.StringOutput { return v.SensorVersion }).(pulumi.StringOutput)
}

func (o SensorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Sensor) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SensorOutput) TiAutomaticUpdates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Sensor) pulumi.BoolPtrOutput { return v.TiAutomaticUpdates }).(pulumi.BoolPtrOutput)
}

func (o SensorOutput) TiStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Sensor) pulumi.StringOutput { return v.TiStatus }).(pulumi.StringOutput)
}

func (o SensorOutput) TiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Sensor) pulumi.StringOutput { return v.TiVersion }).(pulumi.StringOutput)
}

func (o SensorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Sensor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SensorOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sensor) pulumi.StringPtrOutput { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SensorOutput{})
}
