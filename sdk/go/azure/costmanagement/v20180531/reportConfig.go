


package v20180531

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ReportConfig struct {
	pulumi.CustomResourceState

	Definition   ReportConfigDefinitionResponseOutput   `pulumi:"definition"`
	DeliveryInfo ReportConfigDeliveryInfoResponseOutput `pulumi:"deliveryInfo"`
	Format       pulumi.StringPtrOutput                 `pulumi:"format"`
	Name         pulumi.StringOutput                    `pulumi:"name"`
	Schedule     ReportConfigScheduleResponsePtrOutput  `pulumi:"schedule"`
	Tags         pulumi.StringMapOutput                 `pulumi:"tags"`
	Type         pulumi.StringOutput                    `pulumi:"type"`
}


func NewReportConfig(ctx *pulumi.Context,
	name string, args *ReportConfigArgs, opts ...pulumi.ResourceOption) (*ReportConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	if args.DeliveryInfo == nil {
		return nil, errors.New("invalid value for required argument 'DeliveryInfo'")
	}
	var resource ReportConfig
	err := ctx.RegisterResource("azure-native:costmanagement/v20180531:ReportConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReportConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReportConfigState, opts ...pulumi.ResourceOption) (*ReportConfig, error) {
	var resource ReportConfig
	err := ctx.ReadResource("azure-native:costmanagement/v20180531:ReportConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type reportConfigState struct {
}

type ReportConfigState struct {
}

func (ReportConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*reportConfigState)(nil)).Elem()
}

type reportConfigArgs struct {
	Definition       ReportConfigDefinition   `pulumi:"definition"`
	DeliveryInfo     ReportConfigDeliveryInfo `pulumi:"deliveryInfo"`
	Format           *string                  `pulumi:"format"`
	ReportConfigName *string                  `pulumi:"reportConfigName"`
	Schedule         *ReportConfigSchedule    `pulumi:"schedule"`
}


type ReportConfigArgs struct {
	Definition       ReportConfigDefinitionInput
	DeliveryInfo     ReportConfigDeliveryInfoInput
	Format           pulumi.StringPtrInput
	ReportConfigName pulumi.StringPtrInput
	Schedule         ReportConfigSchedulePtrInput
}

func (ReportConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reportConfigArgs)(nil)).Elem()
}

type ReportConfigInput interface {
	pulumi.Input

	ToReportConfigOutput() ReportConfigOutput
	ToReportConfigOutputWithContext(ctx context.Context) ReportConfigOutput
}

func (*ReportConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfig)(nil)).Elem()
}

func (i *ReportConfig) ToReportConfigOutput() ReportConfigOutput {
	return i.ToReportConfigOutputWithContext(context.Background())
}

func (i *ReportConfig) ToReportConfigOutputWithContext(ctx context.Context) ReportConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigOutput)
}

type ReportConfigOutput struct{ *pulumi.OutputState }

func (ReportConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfig)(nil)).Elem()
}

func (o ReportConfigOutput) ToReportConfigOutput() ReportConfigOutput {
	return o
}

func (o ReportConfigOutput) ToReportConfigOutputWithContext(ctx context.Context) ReportConfigOutput {
	return o
}

func (o ReportConfigOutput) Definition() ReportConfigDefinitionResponseOutput {
	return o.ApplyT(func(v *ReportConfig) ReportConfigDefinitionResponseOutput { return v.Definition }).(ReportConfigDefinitionResponseOutput)
}

func (o ReportConfigOutput) DeliveryInfo() ReportConfigDeliveryInfoResponseOutput {
	return o.ApplyT(func(v *ReportConfig) ReportConfigDeliveryInfoResponseOutput { return v.DeliveryInfo }).(ReportConfigDeliveryInfoResponseOutput)
}

func (o ReportConfigOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfig) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

func (o ReportConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ReportConfigOutput) Schedule() ReportConfigScheduleResponsePtrOutput {
	return o.ApplyT(func(v *ReportConfig) ReportConfigScheduleResponsePtrOutput { return v.Schedule }).(ReportConfigScheduleResponsePtrOutput)
}

func (o ReportConfigOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ReportConfig) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ReportConfigOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportConfig) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReportConfigOutput{})
}
