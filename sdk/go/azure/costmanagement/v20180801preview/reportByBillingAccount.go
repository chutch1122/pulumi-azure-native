


package v20180801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReportByBillingAccount struct {
	pulumi.CustomResourceState

	Definition   ReportDefinitionResponseOutput   `pulumi:"definition"`
	DeliveryInfo ReportDeliveryInfoResponseOutput `pulumi:"deliveryInfo"`
	Format       pulumi.StringPtrOutput           `pulumi:"format"`
	Name         pulumi.StringOutput              `pulumi:"name"`
	Schedule     ReportScheduleResponsePtrOutput  `pulumi:"schedule"`
	Tags         pulumi.StringMapOutput           `pulumi:"tags"`
	Type         pulumi.StringOutput              `pulumi:"type"`
}


func NewReportByBillingAccount(ctx *pulumi.Context,
	name string, args *ReportByBillingAccountArgs, opts ...pulumi.ResourceOption) (*ReportByBillingAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountId == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountId'")
	}
	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	if args.DeliveryInfo == nil {
		return nil, errors.New("invalid value for required argument 'DeliveryInfo'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement:ReportByBillingAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource ReportByBillingAccount
	err := ctx.RegisterResource("azure-native:costmanagement/v20180801preview:ReportByBillingAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetReportByBillingAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReportByBillingAccountState, opts ...pulumi.ResourceOption) (*ReportByBillingAccount, error) {
	var resource ReportByBillingAccount
	err := ctx.ReadResource("azure-native:costmanagement/v20180801preview:ReportByBillingAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type reportByBillingAccountState struct {
}

type ReportByBillingAccountState struct {
}

func (ReportByBillingAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*reportByBillingAccountState)(nil)).Elem()
}

type reportByBillingAccountArgs struct {
	BillingAccountId string             `pulumi:"billingAccountId"`
	Definition       ReportDefinition   `pulumi:"definition"`
	DeliveryInfo     ReportDeliveryInfo `pulumi:"deliveryInfo"`
	Format           *string            `pulumi:"format"`
	ReportName       *string            `pulumi:"reportName"`
	Schedule         *ReportSchedule    `pulumi:"schedule"`
}


type ReportByBillingAccountArgs struct {
	BillingAccountId pulumi.StringInput
	Definition       ReportDefinitionInput
	DeliveryInfo     ReportDeliveryInfoInput
	Format           pulumi.StringPtrInput
	ReportName       pulumi.StringPtrInput
	Schedule         ReportSchedulePtrInput
}

func (ReportByBillingAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reportByBillingAccountArgs)(nil)).Elem()
}

type ReportByBillingAccountInput interface {
	pulumi.Input

	ToReportByBillingAccountOutput() ReportByBillingAccountOutput
	ToReportByBillingAccountOutputWithContext(ctx context.Context) ReportByBillingAccountOutput
}

func (*ReportByBillingAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportByBillingAccount)(nil)).Elem()
}

func (i *ReportByBillingAccount) ToReportByBillingAccountOutput() ReportByBillingAccountOutput {
	return i.ToReportByBillingAccountOutputWithContext(context.Background())
}

func (i *ReportByBillingAccount) ToReportByBillingAccountOutputWithContext(ctx context.Context) ReportByBillingAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportByBillingAccountOutput)
}

type ReportByBillingAccountOutput struct{ *pulumi.OutputState }

func (ReportByBillingAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportByBillingAccount)(nil)).Elem()
}

func (o ReportByBillingAccountOutput) ToReportByBillingAccountOutput() ReportByBillingAccountOutput {
	return o
}

func (o ReportByBillingAccountOutput) ToReportByBillingAccountOutputWithContext(ctx context.Context) ReportByBillingAccountOutput {
	return o
}

func (o ReportByBillingAccountOutput) Definition() ReportDefinitionResponseOutput {
	return o.ApplyT(func(v *ReportByBillingAccount) ReportDefinitionResponseOutput { return v.Definition }).(ReportDefinitionResponseOutput)
}

func (o ReportByBillingAccountOutput) DeliveryInfo() ReportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v *ReportByBillingAccount) ReportDeliveryInfoResponseOutput { return v.DeliveryInfo }).(ReportDeliveryInfoResponseOutput)
}

func (o ReportByBillingAccountOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportByBillingAccount) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

func (o ReportByBillingAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportByBillingAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ReportByBillingAccountOutput) Schedule() ReportScheduleResponsePtrOutput {
	return o.ApplyT(func(v *ReportByBillingAccount) ReportScheduleResponsePtrOutput { return v.Schedule }).(ReportScheduleResponsePtrOutput)
}

func (o ReportByBillingAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ReportByBillingAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ReportByBillingAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportByBillingAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReportByBillingAccountOutput{})
}
