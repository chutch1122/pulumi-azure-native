


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Job struct {
	pulumi.CustomResourceState

	CorrelationData pulumi.StringMapOutput            `pulumi:"correlationData"`
	Created         pulumi.StringOutput               `pulumi:"created"`
	Description     pulumi.StringPtrOutput            `pulumi:"description"`
	Input           pulumi.AnyOutput                  `pulumi:"input"`
	LastModified    pulumi.StringOutput               `pulumi:"lastModified"`
	Name            pulumi.StringOutput               `pulumi:"name"`
	Outputs         JobOutputAssetResponseArrayOutput `pulumi:"outputs"`
	Priority        pulumi.StringPtrOutput            `pulumi:"priority"`
	State           pulumi.StringOutput               `pulumi:"state"`
	Type            pulumi.StringOutput               `pulumi:"type"`
}


func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Input == nil {
		return nil, errors.New("invalid value for required argument 'Input'")
	}
	if args.Outputs == nil {
		return nil, errors.New("invalid value for required argument 'Outputs'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TransformName == nil {
		return nil, errors.New("invalid value for required argument 'TransformName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media:Job"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:Job"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:Job"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:Job"),
		},
		{
			Type: pulumi.String("azure-native:media/v20211101:Job"),
		},
	})
	opts = append(opts, aliases)
	var resource Job
	err := ctx.RegisterResource("azure-native:media/v20180601preview:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("azure-native:media/v20180601preview:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type jobState struct {
}

type JobState struct {
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	AccountName       string            `pulumi:"accountName"`
	CorrelationData   map[string]string `pulumi:"correlationData"`
	Description       *string           `pulumi:"description"`
	Input             interface{}       `pulumi:"input"`
	JobName           *string           `pulumi:"jobName"`
	Outputs           []JobOutputAsset  `pulumi:"outputs"`
	Priority          *string           `pulumi:"priority"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	TransformName     string            `pulumi:"transformName"`
}


type JobArgs struct {
	AccountName       pulumi.StringInput
	CorrelationData   pulumi.StringMapInput
	Description       pulumi.StringPtrInput
	Input             pulumi.Input
	JobName           pulumi.StringPtrInput
	Outputs           JobOutputAssetArrayInput
	Priority          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	TransformName     pulumi.StringInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

func (o JobOutput) CorrelationData() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Job) pulumi.StringMapOutput { return v.CorrelationData }).(pulumi.StringMapOutput)
}

func (o JobOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

func (o JobOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o JobOutput) Input() pulumi.AnyOutput {
	return o.ApplyT(func(v *Job) pulumi.AnyOutput { return v.Input }).(pulumi.AnyOutput)
}

func (o JobOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

func (o JobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o JobOutput) Outputs() JobOutputAssetResponseArrayOutput {
	return o.ApplyT(func(v *Job) JobOutputAssetResponseArrayOutput { return v.Outputs }).(JobOutputAssetResponseArrayOutput)
}

func (o JobOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.Priority }).(pulumi.StringPtrOutput)
}

func (o JobOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o JobOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(JobOutput{})
}
