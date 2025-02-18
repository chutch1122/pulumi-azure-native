


package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ScheduleResource struct {
	pulumi.CustomResourceState

	DailyRecurrence   DayDetailsResponsePtrOutput  `pulumi:"dailyRecurrence"`
	HourlyRecurrence  HourDetailsResponsePtrOutput `pulumi:"hourlyRecurrence"`
	Location          pulumi.StringPtrOutput       `pulumi:"location"`
	Name              pulumi.StringPtrOutput       `pulumi:"name"`
	ProvisioningState pulumi.StringPtrOutput       `pulumi:"provisioningState"`
	Status            pulumi.StringPtrOutput       `pulumi:"status"`
	Tags              pulumi.StringMapOutput       `pulumi:"tags"`
	TaskType          pulumi.StringPtrOutput       `pulumi:"taskType"`
	TimeZoneId        pulumi.StringPtrOutput       `pulumi:"timeZoneId"`
	Type              pulumi.StringPtrOutput       `pulumi:"type"`
	WeeklyRecurrence  WeekDetailsResponsePtrOutput `pulumi:"weeklyRecurrence"`
}


func NewScheduleResource(ctx *pulumi.Context,
	name string, args *ScheduleResourceArgs, opts ...pulumi.ResourceOption) (*ScheduleResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:ScheduleResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:ScheduleResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:ScheduleResource"),
		},
	})
	opts = append(opts, aliases)
	var resource ScheduleResource
	err := ctx.RegisterResource("azure-native:devtestlab/v20150521preview:ScheduleResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScheduleResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduleResourceState, opts ...pulumi.ResourceOption) (*ScheduleResource, error) {
	var resource ScheduleResource
	err := ctx.ReadResource("azure-native:devtestlab/v20150521preview:ScheduleResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type scheduleResourceState struct {
}

type ScheduleResourceState struct {
}

func (ScheduleResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleResourceState)(nil)).Elem()
}

type scheduleResourceArgs struct {
	DailyRecurrence   *DayDetails       `pulumi:"dailyRecurrence"`
	HourlyRecurrence  *HourDetails      `pulumi:"hourlyRecurrence"`
	Id                *string           `pulumi:"id"`
	LabName           string            `pulumi:"labName"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Status            *string           `pulumi:"status"`
	Tags              map[string]string `pulumi:"tags"`
	TaskType          *string           `pulumi:"taskType"`
	TimeZoneId        *string           `pulumi:"timeZoneId"`
	Type              *string           `pulumi:"type"`
	WeeklyRecurrence  *WeekDetails      `pulumi:"weeklyRecurrence"`
}


type ScheduleResourceArgs struct {
	DailyRecurrence   DayDetailsPtrInput
	HourlyRecurrence  HourDetailsPtrInput
	Id                pulumi.StringPtrInput
	LabName           pulumi.StringInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	ProvisioningState pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Status            pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	TaskType          pulumi.StringPtrInput
	TimeZoneId        pulumi.StringPtrInput
	Type              pulumi.StringPtrInput
	WeeklyRecurrence  WeekDetailsPtrInput
}

func (ScheduleResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleResourceArgs)(nil)).Elem()
}

type ScheduleResourceInput interface {
	pulumi.Input

	ToScheduleResourceOutput() ScheduleResourceOutput
	ToScheduleResourceOutputWithContext(ctx context.Context) ScheduleResourceOutput
}

func (*ScheduleResource) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleResource)(nil)).Elem()
}

func (i *ScheduleResource) ToScheduleResourceOutput() ScheduleResourceOutput {
	return i.ToScheduleResourceOutputWithContext(context.Background())
}

func (i *ScheduleResource) ToScheduleResourceOutputWithContext(ctx context.Context) ScheduleResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleResourceOutput)
}

type ScheduleResourceOutput struct{ *pulumi.OutputState }

func (ScheduleResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleResource)(nil)).Elem()
}

func (o ScheduleResourceOutput) ToScheduleResourceOutput() ScheduleResourceOutput {
	return o
}

func (o ScheduleResourceOutput) ToScheduleResourceOutputWithContext(ctx context.Context) ScheduleResourceOutput {
	return o
}

func (o ScheduleResourceOutput) DailyRecurrence() DayDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ScheduleResource) DayDetailsResponsePtrOutput { return v.DailyRecurrence }).(DayDetailsResponsePtrOutput)
}

func (o ScheduleResourceOutput) HourlyRecurrence() HourDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ScheduleResource) HourDetailsResponsePtrOutput { return v.HourlyRecurrence }).(HourDetailsResponsePtrOutput)
}

func (o ScheduleResourceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResource) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ScheduleResourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResource) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ScheduleResourceOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResource) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ScheduleResourceOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResource) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ScheduleResourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ScheduleResource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ScheduleResourceOutput) TaskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResource) pulumi.StringPtrOutput { return v.TaskType }).(pulumi.StringPtrOutput)
}

func (o ScheduleResourceOutput) TimeZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResource) pulumi.StringPtrOutput { return v.TimeZoneId }).(pulumi.StringPtrOutput)
}

func (o ScheduleResourceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResource) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ScheduleResourceOutput) WeeklyRecurrence() WeekDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ScheduleResource) WeekDetailsResponsePtrOutput { return v.WeeklyRecurrence }).(WeekDetailsResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduleResourceOutput{})
}
