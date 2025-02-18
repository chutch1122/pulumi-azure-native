


package v20210801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Logger struct {
	pulumi.CustomResourceState

	Credentials pulumi.StringMapOutput `pulumi:"credentials"`
	Description pulumi.StringPtrOutput `pulumi:"description"`
	IsBuffered  pulumi.BoolPtrOutput   `pulumi:"isBuffered"`
	LoggerType  pulumi.StringOutput    `pulumi:"loggerType"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	ResourceId  pulumi.StringPtrOutput `pulumi:"resourceId"`
	Type        pulumi.StringOutput    `pulumi:"type"`
}


func NewLogger(ctx *pulumi.Context,
	name string, args *LoggerArgs, opts ...pulumi.ResourceOption) (*Logger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoggerType == nil {
		return nil, errors.New("invalid value for required argument 'LoggerType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Logger"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:Logger"),
		},
	})
	opts = append(opts, aliases)
	var resource Logger
	err := ctx.RegisterResource("azure-native:apimanagement/v20210801:Logger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLogger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoggerState, opts ...pulumi.ResourceOption) (*Logger, error) {
	var resource Logger
	err := ctx.ReadResource("azure-native:apimanagement/v20210801:Logger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type loggerState struct {
}

type LoggerState struct {
}

func (LoggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loggerState)(nil)).Elem()
}

type loggerArgs struct {
	Credentials       map[string]string `pulumi:"credentials"`
	Description       *string           `pulumi:"description"`
	IsBuffered        *bool             `pulumi:"isBuffered"`
	LoggerId          *string           `pulumi:"loggerId"`
	LoggerType        string            `pulumi:"loggerType"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceId        *string           `pulumi:"resourceId"`
	ServiceName       string            `pulumi:"serviceName"`
}


type LoggerArgs struct {
	Credentials       pulumi.StringMapInput
	Description       pulumi.StringPtrInput
	IsBuffered        pulumi.BoolPtrInput
	LoggerId          pulumi.StringPtrInput
	LoggerType        pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ResourceId        pulumi.StringPtrInput
	ServiceName       pulumi.StringInput
}

func (LoggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loggerArgs)(nil)).Elem()
}

type LoggerInput interface {
	pulumi.Input

	ToLoggerOutput() LoggerOutput
	ToLoggerOutputWithContext(ctx context.Context) LoggerOutput
}

func (*Logger) ElementType() reflect.Type {
	return reflect.TypeOf((**Logger)(nil)).Elem()
}

func (i *Logger) ToLoggerOutput() LoggerOutput {
	return i.ToLoggerOutputWithContext(context.Background())
}

func (i *Logger) ToLoggerOutputWithContext(ctx context.Context) LoggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggerOutput)
}

type LoggerOutput struct{ *pulumi.OutputState }

func (LoggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Logger)(nil)).Elem()
}

func (o LoggerOutput) ToLoggerOutput() LoggerOutput {
	return o
}

func (o LoggerOutput) ToLoggerOutputWithContext(ctx context.Context) LoggerOutput {
	return o
}

func (o LoggerOutput) Credentials() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Logger) pulumi.StringMapOutput { return v.Credentials }).(pulumi.StringMapOutput)
}

func (o LoggerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Logger) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LoggerOutput) IsBuffered() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Logger) pulumi.BoolPtrOutput { return v.IsBuffered }).(pulumi.BoolPtrOutput)
}

func (o LoggerOutput) LoggerType() pulumi.StringOutput {
	return o.ApplyT(func(v *Logger) pulumi.StringOutput { return v.LoggerType }).(pulumi.StringOutput)
}

func (o LoggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Logger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LoggerOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Logger) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o LoggerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Logger) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LoggerOutput{})
}
