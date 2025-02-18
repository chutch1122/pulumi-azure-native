


package v20180801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Report struct {
	pulumi.CustomResourceState

	Definition   ReportDefinitionResponseOutput   `pulumi:"definition"`
	DeliveryInfo ReportDeliveryInfoResponseOutput `pulumi:"deliveryInfo"`
	Format       pulumi.StringPtrOutput           `pulumi:"format"`
	Name         pulumi.StringOutput              `pulumi:"name"`
	Schedule     ReportScheduleResponsePtrOutput  `pulumi:"schedule"`
	Tags         pulumi.StringMapOutput           `pulumi:"tags"`
	Type         pulumi.StringOutput              `pulumi:"type"`
}


func NewReport(ctx *pulumi.Context,
	name string, args *ReportArgs, opts ...pulumi.ResourceOption) (*Report, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	if args.DeliveryInfo == nil {
		return nil, errors.New("invalid value for required argument 'DeliveryInfo'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement:Report"),
		},
	})
	opts = append(opts, aliases)
	var resource Report
	err := ctx.RegisterResource("azure-native:costmanagement/v20180801preview:Report", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReportState, opts ...pulumi.ResourceOption) (*Report, error) {
	var resource Report
	err := ctx.ReadResource("azure-native:costmanagement/v20180801preview:Report", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type reportState struct {
}

type ReportState struct {
}

func (ReportState) ElementType() reflect.Type {
	return reflect.TypeOf((*reportState)(nil)).Elem()
}

type reportArgs struct {
	Definition   ReportDefinition   `pulumi:"definition"`
	DeliveryInfo ReportDeliveryInfo `pulumi:"deliveryInfo"`
	Format       *string            `pulumi:"format"`
	ReportName   *string            `pulumi:"reportName"`
	Schedule     *ReportSchedule    `pulumi:"schedule"`
}


type ReportArgs struct {
	Definition   ReportDefinitionInput
	DeliveryInfo ReportDeliveryInfoInput
	Format       pulumi.StringPtrInput
	ReportName   pulumi.StringPtrInput
	Schedule     ReportSchedulePtrInput
}

func (ReportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reportArgs)(nil)).Elem()
}

type ReportInput interface {
	pulumi.Input

	ToReportOutput() ReportOutput
	ToReportOutputWithContext(ctx context.Context) ReportOutput
}

func (*Report) ElementType() reflect.Type {
	return reflect.TypeOf((**Report)(nil)).Elem()
}

func (i *Report) ToReportOutput() ReportOutput {
	return i.ToReportOutputWithContext(context.Background())
}

func (i *Report) ToReportOutputWithContext(ctx context.Context) ReportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportOutput)
}

type ReportOutput struct{ *pulumi.OutputState }

func (ReportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Report)(nil)).Elem()
}

func (o ReportOutput) ToReportOutput() ReportOutput {
	return o
}

func (o ReportOutput) ToReportOutputWithContext(ctx context.Context) ReportOutput {
	return o
}

func (o ReportOutput) Definition() ReportDefinitionResponseOutput {
	return o.ApplyT(func(v *Report) ReportDefinitionResponseOutput { return v.Definition }).(ReportDefinitionResponseOutput)
}

func (o ReportOutput) DeliveryInfo() ReportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v *Report) ReportDeliveryInfoResponseOutput { return v.DeliveryInfo }).(ReportDeliveryInfoResponseOutput)
}

func (o ReportOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Report) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

func (o ReportOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Report) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ReportOutput) Schedule() ReportScheduleResponsePtrOutput {
	return o.ApplyT(func(v *Report) ReportScheduleResponsePtrOutput { return v.Schedule }).(ReportScheduleResponsePtrOutput)
}

func (o ReportOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Report) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ReportOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Report) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReportOutput{})
}
