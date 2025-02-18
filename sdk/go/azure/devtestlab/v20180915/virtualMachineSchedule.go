


package v20180915

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachineSchedule struct {
	pulumi.CustomResourceState

	CreatedDate          pulumi.StringOutput                   `pulumi:"createdDate"`
	DailyRecurrence      DayDetailsResponsePtrOutput           `pulumi:"dailyRecurrence"`
	HourlyRecurrence     HourDetailsResponsePtrOutput          `pulumi:"hourlyRecurrence"`
	Location             pulumi.StringPtrOutput                `pulumi:"location"`
	Name                 pulumi.StringOutput                   `pulumi:"name"`
	NotificationSettings NotificationSettingsResponsePtrOutput `pulumi:"notificationSettings"`
	ProvisioningState    pulumi.StringOutput                   `pulumi:"provisioningState"`
	Status               pulumi.StringPtrOutput                `pulumi:"status"`
	Tags                 pulumi.StringMapOutput                `pulumi:"tags"`
	TargetResourceId     pulumi.StringPtrOutput                `pulumi:"targetResourceId"`
	TaskType             pulumi.StringPtrOutput                `pulumi:"taskType"`
	TimeZoneId           pulumi.StringPtrOutput                `pulumi:"timeZoneId"`
	Type                 pulumi.StringOutput                   `pulumi:"type"`
	UniqueIdentifier     pulumi.StringOutput                   `pulumi:"uniqueIdentifier"`
	WeeklyRecurrence     WeekDetailsResponsePtrOutput          `pulumi:"weeklyRecurrence"`
}


func NewVirtualMachineSchedule(ctx *pulumi.Context,
	name string, args *VirtualMachineScheduleArgs, opts ...pulumi.ResourceOption) (*VirtualMachineSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualMachineName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineName'")
	}
	if args.NotificationSettings != nil {
		args.NotificationSettings = args.NotificationSettings.ToNotificationSettingsPtrOutput().ApplyT(func(v *NotificationSettings) *NotificationSettings { return v.Defaults() }).(NotificationSettingsPtrOutput)
	}
	if isZero(args.Status) {
		args.Status = pulumi.StringPtr("Disabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:VirtualMachineSchedule"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:VirtualMachineSchedule"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineSchedule
	err := ctx.RegisterResource("azure-native:devtestlab/v20180915:VirtualMachineSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachineSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScheduleState, opts ...pulumi.ResourceOption) (*VirtualMachineSchedule, error) {
	var resource VirtualMachineSchedule
	err := ctx.ReadResource("azure-native:devtestlab/v20180915:VirtualMachineSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualMachineScheduleState struct {
}

type VirtualMachineScheduleState struct {
}

func (VirtualMachineScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScheduleState)(nil)).Elem()
}

type virtualMachineScheduleArgs struct {
	DailyRecurrence      *DayDetails           `pulumi:"dailyRecurrence"`
	HourlyRecurrence     *HourDetails          `pulumi:"hourlyRecurrence"`
	LabName              string                `pulumi:"labName"`
	Location             *string               `pulumi:"location"`
	Name                 *string               `pulumi:"name"`
	NotificationSettings *NotificationSettings `pulumi:"notificationSettings"`
	ResourceGroupName    string                `pulumi:"resourceGroupName"`
	Status               *string               `pulumi:"status"`
	Tags                 map[string]string     `pulumi:"tags"`
	TargetResourceId     *string               `pulumi:"targetResourceId"`
	TaskType             *string               `pulumi:"taskType"`
	TimeZoneId           *string               `pulumi:"timeZoneId"`
	VirtualMachineName   string                `pulumi:"virtualMachineName"`
	WeeklyRecurrence     *WeekDetails          `pulumi:"weeklyRecurrence"`
}


type VirtualMachineScheduleArgs struct {
	DailyRecurrence      DayDetailsPtrInput
	HourlyRecurrence     HourDetailsPtrInput
	LabName              pulumi.StringInput
	Location             pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	NotificationSettings NotificationSettingsPtrInput
	ResourceGroupName    pulumi.StringInput
	Status               pulumi.StringPtrInput
	Tags                 pulumi.StringMapInput
	TargetResourceId     pulumi.StringPtrInput
	TaskType             pulumi.StringPtrInput
	TimeZoneId           pulumi.StringPtrInput
	VirtualMachineName   pulumi.StringInput
	WeeklyRecurrence     WeekDetailsPtrInput
}

func (VirtualMachineScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScheduleArgs)(nil)).Elem()
}

type VirtualMachineScheduleInput interface {
	pulumi.Input

	ToVirtualMachineScheduleOutput() VirtualMachineScheduleOutput
	ToVirtualMachineScheduleOutputWithContext(ctx context.Context) VirtualMachineScheduleOutput
}

func (*VirtualMachineSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineSchedule)(nil)).Elem()
}

func (i *VirtualMachineSchedule) ToVirtualMachineScheduleOutput() VirtualMachineScheduleOutput {
	return i.ToVirtualMachineScheduleOutputWithContext(context.Background())
}

func (i *VirtualMachineSchedule) ToVirtualMachineScheduleOutputWithContext(ctx context.Context) VirtualMachineScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScheduleOutput)
}

type VirtualMachineScheduleOutput struct{ *pulumi.OutputState }

func (VirtualMachineScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineSchedule)(nil)).Elem()
}

func (o VirtualMachineScheduleOutput) ToVirtualMachineScheduleOutput() VirtualMachineScheduleOutput {
	return o
}

func (o VirtualMachineScheduleOutput) ToVirtualMachineScheduleOutputWithContext(ctx context.Context) VirtualMachineScheduleOutput {
	return o
}

func (o VirtualMachineScheduleOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineSchedule) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o VirtualMachineScheduleOutput) DailyRecurrence() DayDetailsResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineSchedule) DayDetailsResponsePtrOutput { return v.DailyRecurrence }).(DayDetailsResponsePtrOutput)
}

func (o VirtualMachineScheduleOutput) HourlyRecurrence() HourDetailsResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineSchedule) HourDetailsResponsePtrOutput { return v.HourlyRecurrence }).(HourDetailsResponsePtrOutput)
}

func (o VirtualMachineScheduleOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSchedule) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineSchedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScheduleOutput) NotificationSettings() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineSchedule) NotificationSettingsResponsePtrOutput { return v.NotificationSettings }).(NotificationSettingsResponsePtrOutput)
}

func (o VirtualMachineScheduleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineSchedule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineScheduleOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSchedule) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScheduleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachineSchedule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualMachineScheduleOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSchedule) pulumi.StringPtrOutput { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScheduleOutput) TaskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSchedule) pulumi.StringPtrOutput { return v.TaskType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScheduleOutput) TimeZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineSchedule) pulumi.StringPtrOutput { return v.TimeZoneId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScheduleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineSchedule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualMachineScheduleOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineSchedule) pulumi.StringOutput { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func (o VirtualMachineScheduleOutput) WeeklyRecurrence() WeekDetailsResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineSchedule) WeekDetailsResponsePtrOutput { return v.WeeklyRecurrence }).(WeekDetailsResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineScheduleOutput{})
}
