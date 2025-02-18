


package v20170301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DiagnosticLogger struct {
	pulumi.CustomResourceState

	Credentials pulumi.StringMapOutput                  `pulumi:"credentials"`
	Description pulumi.StringPtrOutput                  `pulumi:"description"`
	IsBuffered  pulumi.BoolPtrOutput                    `pulumi:"isBuffered"`
	LoggerType  pulumi.StringOutput                     `pulumi:"loggerType"`
	Name        pulumi.StringOutput                     `pulumi:"name"`
	Sampling    LoggerSamplingContractResponsePtrOutput `pulumi:"sampling"`
	Type        pulumi.StringOutput                     `pulumi:"type"`
}


func NewDiagnosticLogger(ctx *pulumi.Context,
	name string, args *DiagnosticLoggerArgs, opts ...pulumi.ResourceOption) (*DiagnosticLogger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiagnosticId == nil {
		return nil, errors.New("invalid value for required argument 'DiagnosticId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:DiagnosticLogger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:DiagnosticLogger"),
		},
	})
	opts = append(opts, aliases)
	var resource DiagnosticLogger
	err := ctx.RegisterResource("azure-native:apimanagement/v20170301:DiagnosticLogger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDiagnosticLogger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiagnosticLoggerState, opts ...pulumi.ResourceOption) (*DiagnosticLogger, error) {
	var resource DiagnosticLogger
	err := ctx.ReadResource("azure-native:apimanagement/v20170301:DiagnosticLogger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type diagnosticLoggerState struct {
}

type DiagnosticLoggerState struct {
}

func (DiagnosticLoggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticLoggerState)(nil)).Elem()
}

type diagnosticLoggerArgs struct {
	DiagnosticId      string  `pulumi:"diagnosticId"`
	Loggerid          *string `pulumi:"loggerid"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type DiagnosticLoggerArgs struct {
	DiagnosticId      pulumi.StringInput
	Loggerid          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (DiagnosticLoggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticLoggerArgs)(nil)).Elem()
}

type DiagnosticLoggerInput interface {
	pulumi.Input

	ToDiagnosticLoggerOutput() DiagnosticLoggerOutput
	ToDiagnosticLoggerOutputWithContext(ctx context.Context) DiagnosticLoggerOutput
}

func (*DiagnosticLogger) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticLogger)(nil)).Elem()
}

func (i *DiagnosticLogger) ToDiagnosticLoggerOutput() DiagnosticLoggerOutput {
	return i.ToDiagnosticLoggerOutputWithContext(context.Background())
}

func (i *DiagnosticLogger) ToDiagnosticLoggerOutputWithContext(ctx context.Context) DiagnosticLoggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticLoggerOutput)
}

type DiagnosticLoggerOutput struct{ *pulumi.OutputState }

func (DiagnosticLoggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticLogger)(nil)).Elem()
}

func (o DiagnosticLoggerOutput) ToDiagnosticLoggerOutput() DiagnosticLoggerOutput {
	return o
}

func (o DiagnosticLoggerOutput) ToDiagnosticLoggerOutputWithContext(ctx context.Context) DiagnosticLoggerOutput {
	return o
}

func (o DiagnosticLoggerOutput) Credentials() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DiagnosticLogger) pulumi.StringMapOutput { return v.Credentials }).(pulumi.StringMapOutput)
}

func (o DiagnosticLoggerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiagnosticLogger) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DiagnosticLoggerOutput) IsBuffered() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiagnosticLogger) pulumi.BoolPtrOutput { return v.IsBuffered }).(pulumi.BoolPtrOutput)
}

func (o DiagnosticLoggerOutput) LoggerType() pulumi.StringOutput {
	return o.ApplyT(func(v *DiagnosticLogger) pulumi.StringOutput { return v.LoggerType }).(pulumi.StringOutput)
}

func (o DiagnosticLoggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DiagnosticLogger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DiagnosticLoggerOutput) Sampling() LoggerSamplingContractResponsePtrOutput {
	return o.ApplyT(func(v *DiagnosticLogger) LoggerSamplingContractResponsePtrOutput { return v.Sampling }).(LoggerSamplingContractResponsePtrOutput)
}

func (o DiagnosticLoggerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DiagnosticLogger) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DiagnosticLoggerOutput{})
}
