


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiDiagnostic struct {
	pulumi.CustomResourceState

	AlwaysLog                    pulumi.StringPtrOutput                      `pulumi:"alwaysLog"`
	Backend                      PipelineDiagnosticSettingsResponsePtrOutput `pulumi:"backend"`
	EnableHttpCorrelationHeaders pulumi.BoolPtrOutput                        `pulumi:"enableHttpCorrelationHeaders"`
	Frontend                     PipelineDiagnosticSettingsResponsePtrOutput `pulumi:"frontend"`
	LoggerId                     pulumi.StringOutput                         `pulumi:"loggerId"`
	Name                         pulumi.StringOutput                         `pulumi:"name"`
	Sampling                     SamplingSettingsResponsePtrOutput           `pulumi:"sampling"`
	Type                         pulumi.StringOutput                         `pulumi:"type"`
}


func NewApiDiagnostic(ctx *pulumi.Context,
	name string, args *ApiDiagnosticArgs, opts ...pulumi.ResourceOption) (*ApiDiagnostic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
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
			Type: pulumi.String("azure-native:apimanagement:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:ApiDiagnostic"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiDiagnostic
	err := ctx.RegisterResource("azure-native:apimanagement/v20180601preview:ApiDiagnostic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiDiagnostic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiDiagnosticState, opts ...pulumi.ResourceOption) (*ApiDiagnostic, error) {
	var resource ApiDiagnostic
	err := ctx.ReadResource("azure-native:apimanagement/v20180601preview:ApiDiagnostic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type apiDiagnosticState struct {
}

type ApiDiagnosticState struct {
}

func (ApiDiagnosticState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiDiagnosticState)(nil)).Elem()
}

type apiDiagnosticArgs struct {
	AlwaysLog                    *string                     `pulumi:"alwaysLog"`
	ApiId                        string                      `pulumi:"apiId"`
	Backend                      *PipelineDiagnosticSettings `pulumi:"backend"`
	DiagnosticId                 *string                     `pulumi:"diagnosticId"`
	EnableHttpCorrelationHeaders *bool                       `pulumi:"enableHttpCorrelationHeaders"`
	Frontend                     *PipelineDiagnosticSettings `pulumi:"frontend"`
	LoggerId                     string                      `pulumi:"loggerId"`
	ResourceGroupName            string                      `pulumi:"resourceGroupName"`
	Sampling                     *SamplingSettings           `pulumi:"sampling"`
	ServiceName                  string                      `pulumi:"serviceName"`
}


type ApiDiagnosticArgs struct {
	AlwaysLog                    pulumi.StringPtrInput
	ApiId                        pulumi.StringInput
	Backend                      PipelineDiagnosticSettingsPtrInput
	DiagnosticId                 pulumi.StringPtrInput
	EnableHttpCorrelationHeaders pulumi.BoolPtrInput
	Frontend                     PipelineDiagnosticSettingsPtrInput
	LoggerId                     pulumi.StringInput
	ResourceGroupName            pulumi.StringInput
	Sampling                     SamplingSettingsPtrInput
	ServiceName                  pulumi.StringInput
}

func (ApiDiagnosticArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiDiagnosticArgs)(nil)).Elem()
}

type ApiDiagnosticInput interface {
	pulumi.Input

	ToApiDiagnosticOutput() ApiDiagnosticOutput
	ToApiDiagnosticOutputWithContext(ctx context.Context) ApiDiagnosticOutput
}

func (*ApiDiagnostic) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDiagnostic)(nil)).Elem()
}

func (i *ApiDiagnostic) ToApiDiagnosticOutput() ApiDiagnosticOutput {
	return i.ToApiDiagnosticOutputWithContext(context.Background())
}

func (i *ApiDiagnostic) ToApiDiagnosticOutputWithContext(ctx context.Context) ApiDiagnosticOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDiagnosticOutput)
}

type ApiDiagnosticOutput struct{ *pulumi.OutputState }

func (ApiDiagnosticOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDiagnostic)(nil)).Elem()
}

func (o ApiDiagnosticOutput) ToApiDiagnosticOutput() ApiDiagnosticOutput {
	return o
}

func (o ApiDiagnosticOutput) ToApiDiagnosticOutputWithContext(ctx context.Context) ApiDiagnosticOutput {
	return o
}

func (o ApiDiagnosticOutput) AlwaysLog() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiDiagnostic) pulumi.StringPtrOutput { return v.AlwaysLog }).(pulumi.StringPtrOutput)
}

func (o ApiDiagnosticOutput) Backend() PipelineDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ApiDiagnostic) PipelineDiagnosticSettingsResponsePtrOutput { return v.Backend }).(PipelineDiagnosticSettingsResponsePtrOutput)
}

func (o ApiDiagnosticOutput) EnableHttpCorrelationHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiDiagnostic) pulumi.BoolPtrOutput { return v.EnableHttpCorrelationHeaders }).(pulumi.BoolPtrOutput)
}

func (o ApiDiagnosticOutput) Frontend() PipelineDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ApiDiagnostic) PipelineDiagnosticSettingsResponsePtrOutput { return v.Frontend }).(PipelineDiagnosticSettingsResponsePtrOutput)
}

func (o ApiDiagnosticOutput) LoggerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiDiagnostic) pulumi.StringOutput { return v.LoggerId }).(pulumi.StringOutput)
}

func (o ApiDiagnosticOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiDiagnostic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApiDiagnosticOutput) Sampling() SamplingSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ApiDiagnostic) SamplingSettingsResponsePtrOutput { return v.Sampling }).(SamplingSettingsResponsePtrOutput)
}

func (o ApiDiagnosticOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiDiagnostic) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiDiagnosticOutput{})
}
