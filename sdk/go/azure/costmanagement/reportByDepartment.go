


package costmanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReportByDepartment struct {
	pulumi.CustomResourceState

	Definition   ReportDefinitionResponseOutput   `pulumi:"definition"`
	DeliveryInfo ReportDeliveryInfoResponseOutput `pulumi:"deliveryInfo"`
	Format       pulumi.StringPtrOutput           `pulumi:"format"`
	Name         pulumi.StringOutput              `pulumi:"name"`
	Schedule     ReportScheduleResponsePtrOutput  `pulumi:"schedule"`
	Tags         pulumi.StringMapOutput           `pulumi:"tags"`
	Type         pulumi.StringOutput              `pulumi:"type"`
}


func NewReportByDepartment(ctx *pulumi.Context,
	name string, args *ReportByDepartmentArgs, opts ...pulumi.ResourceOption) (*ReportByDepartment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	if args.DeliveryInfo == nil {
		return nil, errors.New("invalid value for required argument 'DeliveryInfo'")
	}
	if args.DepartmentId == nil {
		return nil, errors.New("invalid value for required argument 'DepartmentId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement/v20180801preview:ReportByDepartment"),
		},
	})
	opts = append(opts, aliases)
	var resource ReportByDepartment
	err := ctx.RegisterResource("azure-native:costmanagement:ReportByDepartment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReportByDepartment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReportByDepartmentState, opts ...pulumi.ResourceOption) (*ReportByDepartment, error) {
	var resource ReportByDepartment
	err := ctx.ReadResource("azure-native:costmanagement:ReportByDepartment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type reportByDepartmentState struct {
}

type ReportByDepartmentState struct {
}

func (ReportByDepartmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*reportByDepartmentState)(nil)).Elem()
}

type reportByDepartmentArgs struct {
	Definition   ReportDefinition   `pulumi:"definition"`
	DeliveryInfo ReportDeliveryInfo `pulumi:"deliveryInfo"`
	DepartmentId string             `pulumi:"departmentId"`
	Format       *string            `pulumi:"format"`
	ReportName   *string            `pulumi:"reportName"`
	Schedule     *ReportSchedule    `pulumi:"schedule"`
}


type ReportByDepartmentArgs struct {
	Definition   ReportDefinitionInput
	DeliveryInfo ReportDeliveryInfoInput
	DepartmentId pulumi.StringInput
	Format       pulumi.StringPtrInput
	ReportName   pulumi.StringPtrInput
	Schedule     ReportSchedulePtrInput
}

func (ReportByDepartmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reportByDepartmentArgs)(nil)).Elem()
}

type ReportByDepartmentInput interface {
	pulumi.Input

	ToReportByDepartmentOutput() ReportByDepartmentOutput
	ToReportByDepartmentOutputWithContext(ctx context.Context) ReportByDepartmentOutput
}

func (*ReportByDepartment) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportByDepartment)(nil)).Elem()
}

func (i *ReportByDepartment) ToReportByDepartmentOutput() ReportByDepartmentOutput {
	return i.ToReportByDepartmentOutputWithContext(context.Background())
}

func (i *ReportByDepartment) ToReportByDepartmentOutputWithContext(ctx context.Context) ReportByDepartmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportByDepartmentOutput)
}

type ReportByDepartmentOutput struct{ *pulumi.OutputState }

func (ReportByDepartmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportByDepartment)(nil)).Elem()
}

func (o ReportByDepartmentOutput) ToReportByDepartmentOutput() ReportByDepartmentOutput {
	return o
}

func (o ReportByDepartmentOutput) ToReportByDepartmentOutputWithContext(ctx context.Context) ReportByDepartmentOutput {
	return o
}

func (o ReportByDepartmentOutput) Definition() ReportDefinitionResponseOutput {
	return o.ApplyT(func(v *ReportByDepartment) ReportDefinitionResponseOutput { return v.Definition }).(ReportDefinitionResponseOutput)
}

func (o ReportByDepartmentOutput) DeliveryInfo() ReportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v *ReportByDepartment) ReportDeliveryInfoResponseOutput { return v.DeliveryInfo }).(ReportDeliveryInfoResponseOutput)
}

func (o ReportByDepartmentOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportByDepartment) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

func (o ReportByDepartmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportByDepartment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ReportByDepartmentOutput) Schedule() ReportScheduleResponsePtrOutput {
	return o.ApplyT(func(v *ReportByDepartment) ReportScheduleResponsePtrOutput { return v.Schedule }).(ReportScheduleResponsePtrOutput)
}

func (o ReportByDepartmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ReportByDepartment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ReportByDepartmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportByDepartment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReportByDepartmentOutput{})
}
