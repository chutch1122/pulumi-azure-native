


package v20191201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Diagnostic struct {
	pulumi.CustomResourceState

	AlwaysLog               pulumi.StringPtrOutput                      `pulumi:"alwaysLog"`
	Backend                 PipelineDiagnosticSettingsResponsePtrOutput `pulumi:"backend"`
	Frontend                PipelineDiagnosticSettingsResponsePtrOutput `pulumi:"frontend"`
	HttpCorrelationProtocol pulumi.StringPtrOutput                      `pulumi:"httpCorrelationProtocol"`
	LogClientIp             pulumi.BoolPtrOutput                        `pulumi:"logClientIp"`
	LoggerId                pulumi.StringOutput                         `pulumi:"loggerId"`
	Name                    pulumi.StringOutput                         `pulumi:"name"`
	Sampling                SamplingSettingsResponsePtrOutput           `pulumi:"sampling"`
	Type                    pulumi.StringOutput                         `pulumi:"type"`
	Verbosity               pulumi.StringPtrOutput                      `pulumi:"verbosity"`
}


func NewDiagnostic(ctx *pulumi.Context,
	name string, args *DiagnosticArgs, opts ...pulumi.ResourceOption) (*Diagnostic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoggerId == nil {
		return nil, errors.New("invalid value for required argument 'LoggerId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:Diagnostic"),
		},
	})
	opts = append(opts, aliases)
	var resource Diagnostic
	err := ctx.RegisterResource("azure-native:apimanagement/v20191201:Diagnostic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDiagnostic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiagnosticState, opts ...pulumi.ResourceOption) (*Diagnostic, error) {
	var resource Diagnostic
	err := ctx.ReadResource("azure-native:apimanagement/v20191201:Diagnostic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type diagnosticState struct {
}

type DiagnosticState struct {
}

func (DiagnosticState) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticState)(nil)).Elem()
}

type diagnosticArgs struct {
	AlwaysLog               *string                     `pulumi:"alwaysLog"`
	Backend                 *PipelineDiagnosticSettings `pulumi:"backend"`
	DiagnosticId            *string                     `pulumi:"diagnosticId"`
	Frontend                *PipelineDiagnosticSettings `pulumi:"frontend"`
	HttpCorrelationProtocol *string                     `pulumi:"httpCorrelationProtocol"`
	LogClientIp             *bool                       `pulumi:"logClientIp"`
	LoggerId                string                      `pulumi:"loggerId"`
	ResourceGroupName       string                      `pulumi:"resourceGroupName"`
	Sampling                *SamplingSettings           `pulumi:"sampling"`
	ServiceName             string                      `pulumi:"serviceName"`
	Verbosity               *string                     `pulumi:"verbosity"`
}


type DiagnosticArgs struct {
	AlwaysLog               pulumi.StringPtrInput
	Backend                 PipelineDiagnosticSettingsPtrInput
	DiagnosticId            pulumi.StringPtrInput
	Frontend                PipelineDiagnosticSettingsPtrInput
	HttpCorrelationProtocol pulumi.StringPtrInput
	LogClientIp             pulumi.BoolPtrInput
	LoggerId                pulumi.StringInput
	ResourceGroupName       pulumi.StringInput
	Sampling                SamplingSettingsPtrInput
	ServiceName             pulumi.StringInput
	Verbosity               pulumi.StringPtrInput
}

func (DiagnosticArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticArgs)(nil)).Elem()
}

type DiagnosticInput interface {
	pulumi.Input

	ToDiagnosticOutput() DiagnosticOutput
	ToDiagnosticOutputWithContext(ctx context.Context) DiagnosticOutput
}

func (*Diagnostic) ElementType() reflect.Type {
	return reflect.TypeOf((**Diagnostic)(nil)).Elem()
}

func (i *Diagnostic) ToDiagnosticOutput() DiagnosticOutput {
	return i.ToDiagnosticOutputWithContext(context.Background())
}

func (i *Diagnostic) ToDiagnosticOutputWithContext(ctx context.Context) DiagnosticOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticOutput)
}

type DiagnosticOutput struct{ *pulumi.OutputState }

func (DiagnosticOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Diagnostic)(nil)).Elem()
}

func (o DiagnosticOutput) ToDiagnosticOutput() DiagnosticOutput {
	return o
}

func (o DiagnosticOutput) ToDiagnosticOutputWithContext(ctx context.Context) DiagnosticOutput {
	return o
}

func (o DiagnosticOutput) AlwaysLog() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Diagnostic) pulumi.StringPtrOutput { return v.AlwaysLog }).(pulumi.StringPtrOutput)
}

func (o DiagnosticOutput) Backend() PipelineDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Diagnostic) PipelineDiagnosticSettingsResponsePtrOutput { return v.Backend }).(PipelineDiagnosticSettingsResponsePtrOutput)
}

func (o DiagnosticOutput) Frontend() PipelineDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Diagnostic) PipelineDiagnosticSettingsResponsePtrOutput { return v.Frontend }).(PipelineDiagnosticSettingsResponsePtrOutput)
}

func (o DiagnosticOutput) HttpCorrelationProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Diagnostic) pulumi.StringPtrOutput { return v.HttpCorrelationProtocol }).(pulumi.StringPtrOutput)
}

func (o DiagnosticOutput) LogClientIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Diagnostic) pulumi.BoolPtrOutput { return v.LogClientIp }).(pulumi.BoolPtrOutput)
}

func (o DiagnosticOutput) LoggerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Diagnostic) pulumi.StringOutput { return v.LoggerId }).(pulumi.StringOutput)
}

func (o DiagnosticOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Diagnostic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DiagnosticOutput) Sampling() SamplingSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Diagnostic) SamplingSettingsResponsePtrOutput { return v.Sampling }).(SamplingSettingsResponsePtrOutput)
}

func (o DiagnosticOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Diagnostic) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o DiagnosticOutput) Verbosity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Diagnostic) pulumi.StringPtrOutput { return v.Verbosity }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DiagnosticOutput{})
}
