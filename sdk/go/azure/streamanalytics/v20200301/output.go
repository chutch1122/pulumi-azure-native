


package v20200301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Output struct {
	pulumi.CustomResourceState

	Datasource    pulumi.AnyOutput          `pulumi:"datasource"`
	Diagnostics   DiagnosticsResponseOutput `pulumi:"diagnostics"`
	Etag          pulumi.StringOutput       `pulumi:"etag"`
	Name          pulumi.StringPtrOutput    `pulumi:"name"`
	Serialization pulumi.AnyOutput          `pulumi:"serialization"`
	SizeWindow    pulumi.Float64PtrOutput   `pulumi:"sizeWindow"`
	TimeWindow    pulumi.StringPtrOutput    `pulumi:"timeWindow"`
	Type          pulumi.StringOutput       `pulumi:"type"`
}


func NewOutput(ctx *pulumi.Context,
	name string, args *OutputArgs, opts ...pulumi.ResourceOption) (*Output, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobName == nil {
		return nil, errors.New("invalid value for required argument 'JobName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:streamanalytics:Output"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20160301:Output"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20170401preview:Output"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20211001preview:Output"),
		},
	})
	opts = append(opts, aliases)
	var resource Output
	err := ctx.RegisterResource("azure-native:streamanalytics/v20200301:Output", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOutput(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutputState, opts ...pulumi.ResourceOption) (*Output, error) {
	var resource Output
	err := ctx.ReadResource("azure-native:streamanalytics/v20200301:Output", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type outputState struct {
}

type OutputState struct {
}

func (OutputState) ElementType() reflect.Type {
	return reflect.TypeOf((*outputState)(nil)).Elem()
}

type outputArgs struct {
	Datasource        interface{} `pulumi:"datasource"`
	JobName           string      `pulumi:"jobName"`
	Name              *string     `pulumi:"name"`
	OutputName        *string     `pulumi:"outputName"`
	ResourceGroupName string      `pulumi:"resourceGroupName"`
	Serialization     interface{} `pulumi:"serialization"`
	SizeWindow        *float64    `pulumi:"sizeWindow"`
	TimeWindow        *string     `pulumi:"timeWindow"`
}


type OutputArgs struct {
	Datasource        pulumi.Input
	JobName           pulumi.StringInput
	Name              pulumi.StringPtrInput
	OutputName        pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Serialization     pulumi.Input
	SizeWindow        pulumi.Float64PtrInput
	TimeWindow        pulumi.StringPtrInput
}

func (OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*outputArgs)(nil)).Elem()
}

type OutputInput interface {
	pulumi.Input

	ToOutputOutput() OutputOutput
	ToOutputOutputWithContext(ctx context.Context) OutputOutput
}

func (*Output) ElementType() reflect.Type {
	return reflect.TypeOf((**Output)(nil)).Elem()
}

func (i *Output) ToOutputOutput() OutputOutput {
	return i.ToOutputOutputWithContext(context.Background())
}

func (i *Output) ToOutputOutputWithContext(ctx context.Context) OutputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputOutput)
}

type OutputOutput struct{ *pulumi.OutputState }

func (OutputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Output)(nil)).Elem()
}

func (o OutputOutput) ToOutputOutput() OutputOutput {
	return o
}

func (o OutputOutput) ToOutputOutputWithContext(ctx context.Context) OutputOutput {
	return o
}

func (o OutputOutput) Datasource() pulumi.AnyOutput {
	return o.ApplyT(func(v *Output) pulumi.AnyOutput { return v.Datasource }).(pulumi.AnyOutput)
}

func (o OutputOutput) Diagnostics() DiagnosticsResponseOutput {
	return o.ApplyT(func(v *Output) DiagnosticsResponseOutput { return v.Diagnostics }).(DiagnosticsResponseOutput)
}

func (o OutputOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Output) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o OutputOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Output) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OutputOutput) Serialization() pulumi.AnyOutput {
	return o.ApplyT(func(v *Output) pulumi.AnyOutput { return v.Serialization }).(pulumi.AnyOutput)
}

func (o OutputOutput) SizeWindow() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Output) pulumi.Float64PtrOutput { return v.SizeWindow }).(pulumi.Float64PtrOutput)
}

func (o OutputOutput) TimeWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Output) pulumi.StringPtrOutput { return v.TimeWindow }).(pulumi.StringPtrOutput)
}

func (o OutputOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Output) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OutputOutput{})
}
