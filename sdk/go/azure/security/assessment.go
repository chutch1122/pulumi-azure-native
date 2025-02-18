


package security

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Assessment struct {
	pulumi.CustomResourceState

	AdditionalData  pulumi.StringMapOutput                                `pulumi:"additionalData"`
	DisplayName     pulumi.StringOutput                                   `pulumi:"displayName"`
	Links           AssessmentLinksResponseOutput                         `pulumi:"links"`
	Metadata        SecurityAssessmentMetadataPropertiesResponsePtrOutput `pulumi:"metadata"`
	Name            pulumi.StringOutput                                   `pulumi:"name"`
	PartnersData    SecurityAssessmentPartnerDataResponsePtrOutput        `pulumi:"partnersData"`
	ResourceDetails pulumi.AnyOutput                                      `pulumi:"resourceDetails"`
	Status          AssessmentStatusResponseOutput                        `pulumi:"status"`
	Type            pulumi.StringOutput                                   `pulumi:"type"`
}


func NewAssessment(ctx *pulumi.Context,
	name string, args *AssessmentArgs, opts ...pulumi.ResourceOption) (*Assessment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceDetails == nil {
		return nil, errors.New("invalid value for required argument 'ResourceDetails'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security/v20190101preview:Assessment"),
		},
		{
			Type: pulumi.String("azure-native:security/v20200101:Assessment"),
		},
		{
			Type: pulumi.String("azure-native:security/v20210601:Assessment"),
		},
	})
	opts = append(opts, aliases)
	var resource Assessment
	err := ctx.RegisterResource("azure-native:security:Assessment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAssessment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssessmentState, opts ...pulumi.ResourceOption) (*Assessment, error) {
	var resource Assessment
	err := ctx.ReadResource("azure-native:security:Assessment", name, id, state, &resource, opts...)
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
	AdditionalData  map[string]string                     `pulumi:"additionalData"`
	AssessmentName  *string                               `pulumi:"assessmentName"`
	Metadata        *SecurityAssessmentMetadataProperties `pulumi:"metadata"`
	PartnersData    *SecurityAssessmentPartnerData        `pulumi:"partnersData"`
	ResourceDetails interface{}                           `pulumi:"resourceDetails"`
	ResourceId      string                                `pulumi:"resourceId"`
	Status          AssessmentStatus                      `pulumi:"status"`
}


type AssessmentArgs struct {
	AdditionalData  pulumi.StringMapInput
	AssessmentName  pulumi.StringPtrInput
	Metadata        SecurityAssessmentMetadataPropertiesPtrInput
	PartnersData    SecurityAssessmentPartnerDataPtrInput
	ResourceDetails pulumi.Input
	ResourceId      pulumi.StringInput
	Status          AssessmentStatusInput
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

func (o AssessmentOutput) AdditionalData() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringMapOutput { return v.AdditionalData }).(pulumi.StringMapOutput)
}

func (o AssessmentOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o AssessmentOutput) Links() AssessmentLinksResponseOutput {
	return o.ApplyT(func(v *Assessment) AssessmentLinksResponseOutput { return v.Links }).(AssessmentLinksResponseOutput)
}

func (o AssessmentOutput) Metadata() SecurityAssessmentMetadataPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Assessment) SecurityAssessmentMetadataPropertiesResponsePtrOutput { return v.Metadata }).(SecurityAssessmentMetadataPropertiesResponsePtrOutput)
}

func (o AssessmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AssessmentOutput) PartnersData() SecurityAssessmentPartnerDataResponsePtrOutput {
	return o.ApplyT(func(v *Assessment) SecurityAssessmentPartnerDataResponsePtrOutput { return v.PartnersData }).(SecurityAssessmentPartnerDataResponsePtrOutput)
}

func (o AssessmentOutput) ResourceDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v *Assessment) pulumi.AnyOutput { return v.ResourceDetails }).(pulumi.AnyOutput)
}

func (o AssessmentOutput) Status() AssessmentStatusResponseOutput {
	return o.ApplyT(func(v *Assessment) AssessmentStatusResponseOutput { return v.Status }).(AssessmentStatusResponseOutput)
}

func (o AssessmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AssessmentOutput{})
}
