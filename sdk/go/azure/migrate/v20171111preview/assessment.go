


package v20171111preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Assessment struct {
	pulumi.CustomResourceState

	AzureHybridUseBenefit  pulumi.StringOutput    `pulumi:"azureHybridUseBenefit"`
	AzureLocation          pulumi.StringOutput    `pulumi:"azureLocation"`
	AzureOfferCode         pulumi.StringOutput    `pulumi:"azureOfferCode"`
	AzurePricingTier       pulumi.StringOutput    `pulumi:"azurePricingTier"`
	AzureStorageRedundancy pulumi.StringOutput    `pulumi:"azureStorageRedundancy"`
	CreatedTimestamp       pulumi.StringOutput    `pulumi:"createdTimestamp"`
	Currency               pulumi.StringOutput    `pulumi:"currency"`
	DiscountPercentage     pulumi.Float64Output   `pulumi:"discountPercentage"`
	ETag                   pulumi.StringPtrOutput `pulumi:"eTag"`
	MonthlyBandwidthCost   pulumi.Float64Output   `pulumi:"monthlyBandwidthCost"`
	MonthlyComputeCost     pulumi.Float64Output   `pulumi:"monthlyComputeCost"`
	MonthlyStorageCost     pulumi.Float64Output   `pulumi:"monthlyStorageCost"`
	Name                   pulumi.StringOutput    `pulumi:"name"`
	NumberOfMachines       pulumi.IntOutput       `pulumi:"numberOfMachines"`
	Percentile             pulumi.StringOutput    `pulumi:"percentile"`
	PricesTimestamp        pulumi.StringOutput    `pulumi:"pricesTimestamp"`
	ScalingFactor          pulumi.Float64Output   `pulumi:"scalingFactor"`
	Stage                  pulumi.StringOutput    `pulumi:"stage"`
	Status                 pulumi.StringOutput    `pulumi:"status"`
	TimeRange              pulumi.StringOutput    `pulumi:"timeRange"`
	Type                   pulumi.StringOutput    `pulumi:"type"`
	UpdatedTimestamp       pulumi.StringOutput    `pulumi:"updatedTimestamp"`
}


func NewAssessment(ctx *pulumi.Context,
	name string, args *AssessmentArgs, opts ...pulumi.ResourceOption) (*Assessment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AzureHybridUseBenefit == nil {
		return nil, errors.New("invalid value for required argument 'AzureHybridUseBenefit'")
	}
	if args.AzureLocation == nil {
		return nil, errors.New("invalid value for required argument 'AzureLocation'")
	}
	if args.AzureOfferCode == nil {
		return nil, errors.New("invalid value for required argument 'AzureOfferCode'")
	}
	if args.AzurePricingTier == nil {
		return nil, errors.New("invalid value for required argument 'AzurePricingTier'")
	}
	if args.AzureStorageRedundancy == nil {
		return nil, errors.New("invalid value for required argument 'AzureStorageRedundancy'")
	}
	if args.Currency == nil {
		return nil, errors.New("invalid value for required argument 'Currency'")
	}
	if args.DiscountPercentage == nil {
		return nil, errors.New("invalid value for required argument 'DiscountPercentage'")
	}
	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.Percentile == nil {
		return nil, errors.New("invalid value for required argument 'Percentile'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScalingFactor == nil {
		return nil, errors.New("invalid value for required argument 'ScalingFactor'")
	}
	if args.Stage == nil {
		return nil, errors.New("invalid value for required argument 'Stage'")
	}
	if args.TimeRange == nil {
		return nil, errors.New("invalid value for required argument 'TimeRange'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20180202:Assessment"),
		},
	})
	opts = append(opts, aliases)
	var resource Assessment
	err := ctx.RegisterResource("azure-native:migrate/v20171111preview:Assessment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAssessment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssessmentState, opts ...pulumi.ResourceOption) (*Assessment, error) {
	var resource Assessment
	err := ctx.ReadResource("azure-native:migrate/v20171111preview:Assessment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type assessmentState struct {
}

type AssessmentState struct {
}

func (AssessmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentState)(nil)).Elem()
}

type assessmentArgs struct {
	AssessmentName         *string `pulumi:"assessmentName"`
	AzureHybridUseBenefit  string  `pulumi:"azureHybridUseBenefit"`
	AzureLocation          string  `pulumi:"azureLocation"`
	AzureOfferCode         string  `pulumi:"azureOfferCode"`
	AzurePricingTier       string  `pulumi:"azurePricingTier"`
	AzureStorageRedundancy string  `pulumi:"azureStorageRedundancy"`
	Currency               string  `pulumi:"currency"`
	DiscountPercentage     float64 `pulumi:"discountPercentage"`
	ETag                   *string `pulumi:"eTag"`
	GroupName              string  `pulumi:"groupName"`
	Percentile             string  `pulumi:"percentile"`
	ProjectName            string  `pulumi:"projectName"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
	ScalingFactor          float64 `pulumi:"scalingFactor"`
	Stage                  string  `pulumi:"stage"`
	TimeRange              string  `pulumi:"timeRange"`
}


type AssessmentArgs struct {
	AssessmentName         pulumi.StringPtrInput
	AzureHybridUseBenefit  pulumi.StringInput
	AzureLocation          pulumi.StringInput
	AzureOfferCode         pulumi.StringInput
	AzurePricingTier       pulumi.StringInput
	AzureStorageRedundancy pulumi.StringInput
	Currency               pulumi.StringInput
	DiscountPercentage     pulumi.Float64Input
	ETag                   pulumi.StringPtrInput
	GroupName              pulumi.StringInput
	Percentile             pulumi.StringInput
	ProjectName            pulumi.StringInput
	ResourceGroupName      pulumi.StringInput
	ScalingFactor          pulumi.Float64Input
	Stage                  pulumi.StringInput
	TimeRange              pulumi.StringInput
}

func (AssessmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentArgs)(nil)).Elem()
}

type AssessmentInput interface {
	pulumi.Input

	ToAssessmentOutput() AssessmentOutput
	ToAssessmentOutputWithContext(ctx context.Context) AssessmentOutput
}

func (*Assessment) ElementType() reflect.Type {
	return reflect.TypeOf((**Assessment)(nil)).Elem()
}

func (i *Assessment) ToAssessmentOutput() AssessmentOutput {
	return i.ToAssessmentOutputWithContext(context.Background())
}

func (i *Assessment) ToAssessmentOutputWithContext(ctx context.Context) AssessmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentOutput)
}

type AssessmentOutput struct{ *pulumi.OutputState }

func (AssessmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Assessment)(nil)).Elem()
}

func (o AssessmentOutput) ToAssessmentOutput() AssessmentOutput {
	return o
}

func (o AssessmentOutput) ToAssessmentOutputWithContext(ctx context.Context) AssessmentOutput {
	return o
}

func (o AssessmentOutput) AzureHybridUseBenefit() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.AzureHybridUseBenefit }).(pulumi.StringOutput)
}

func (o AssessmentOutput) AzureLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.AzureLocation }).(pulumi.StringOutput)
}

func (o AssessmentOutput) AzureOfferCode() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.AzureOfferCode }).(pulumi.StringOutput)
}

func (o AssessmentOutput) AzurePricingTier() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.AzurePricingTier }).(pulumi.StringOutput)
}

func (o AssessmentOutput) AzureStorageRedundancy() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.AzureStorageRedundancy }).(pulumi.StringOutput)
}

func (o AssessmentOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

func (o AssessmentOutput) Currency() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.Currency }).(pulumi.StringOutput)
}

func (o AssessmentOutput) DiscountPercentage() pulumi.Float64Output {
	return o.ApplyT(func(v *Assessment) pulumi.Float64Output { return v.DiscountPercentage }).(pulumi.Float64Output)
}

func (o AssessmentOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o AssessmentOutput) MonthlyBandwidthCost() pulumi.Float64Output {
	return o.ApplyT(func(v *Assessment) pulumi.Float64Output { return v.MonthlyBandwidthCost }).(pulumi.Float64Output)
}

func (o AssessmentOutput) MonthlyComputeCost() pulumi.Float64Output {
	return o.ApplyT(func(v *Assessment) pulumi.Float64Output { return v.MonthlyComputeCost }).(pulumi.Float64Output)
}

func (o AssessmentOutput) MonthlyStorageCost() pulumi.Float64Output {
	return o.ApplyT(func(v *Assessment) pulumi.Float64Output { return v.MonthlyStorageCost }).(pulumi.Float64Output)
}

func (o AssessmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AssessmentOutput) NumberOfMachines() pulumi.IntOutput {
	return o.ApplyT(func(v *Assessment) pulumi.IntOutput { return v.NumberOfMachines }).(pulumi.IntOutput)
}

func (o AssessmentOutput) Percentile() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.Percentile }).(pulumi.StringOutput)
}

func (o AssessmentOutput) PricesTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.PricesTimestamp }).(pulumi.StringOutput)
}

func (o AssessmentOutput) ScalingFactor() pulumi.Float64Output {
	return o.ApplyT(func(v *Assessment) pulumi.Float64Output { return v.ScalingFactor }).(pulumi.Float64Output)
}

func (o AssessmentOutput) Stage() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.Stage }).(pulumi.StringOutput)
}

func (o AssessmentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o AssessmentOutput) TimeRange() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.TimeRange }).(pulumi.StringOutput)
}

func (o AssessmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AssessmentOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AssessmentOutput{})
}
