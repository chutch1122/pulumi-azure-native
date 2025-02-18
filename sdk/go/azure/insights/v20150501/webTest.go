


package v20150501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebTest struct {
	pulumi.CustomResourceState

	Configuration      WebTestPropertiesResponseConfigurationPtrOutput `pulumi:"configuration"`
	Description        pulumi.StringPtrOutput                          `pulumi:"description"`
	Enabled            pulumi.BoolPtrOutput                            `pulumi:"enabled"`
	Frequency          pulumi.IntPtrOutput                             `pulumi:"frequency"`
	Kind               pulumi.StringPtrOutput                          `pulumi:"kind"`
	Location           pulumi.StringOutput                             `pulumi:"location"`
	Locations          WebTestGeolocationResponseArrayOutput           `pulumi:"locations"`
	Name               pulumi.StringOutput                             `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput                             `pulumi:"provisioningState"`
	RetryEnabled       pulumi.BoolPtrOutput                            `pulumi:"retryEnabled"`
	SyntheticMonitorId pulumi.StringOutput                             `pulumi:"syntheticMonitorId"`
	Tags               pulumi.StringMapOutput                          `pulumi:"tags"`
	Timeout            pulumi.IntPtrOutput                             `pulumi:"timeout"`
	Type               pulumi.StringOutput                             `pulumi:"type"`
	WebTestKind        pulumi.StringOutput                             `pulumi:"webTestKind"`
	WebTestName        pulumi.StringOutput                             `pulumi:"webTestName"`
}


func NewWebTest(ctx *pulumi.Context,
	name string, args *WebTestArgs, opts ...pulumi.ResourceOption) (*WebTest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Locations == nil {
		return nil, errors.New("invalid value for required argument 'Locations'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SyntheticMonitorId == nil {
		return nil, errors.New("invalid value for required argument 'SyntheticMonitorId'")
	}
	if isZero(args.Frequency) {
		args.Frequency = pulumi.IntPtr(300)
	}
	if isZero(args.Kind) {
		args.Kind = WebTestKind("ping")
	}
	if isZero(args.Timeout) {
		args.Timeout = pulumi.IntPtr(30)
	}
	if isZero(args.WebTestKind) {
		args.WebTestKind = WebTestKind("ping")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights:WebTest"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20180501preview:WebTest"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20201005preview:WebTest"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20220615:WebTest"),
		},
	})
	opts = append(opts, aliases)
	var resource WebTest
	err := ctx.RegisterResource("azure-native:insights/v20150501:WebTest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebTest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebTestState, opts ...pulumi.ResourceOption) (*WebTest, error) {
	var resource WebTest
	err := ctx.ReadResource("azure-native:insights/v20150501:WebTest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webTestState struct {
}

type WebTestState struct {
}

func (WebTestState) ElementType() reflect.Type {
	return reflect.TypeOf((*webTestState)(nil)).Elem()
}

type webTestArgs struct {
	Configuration      *WebTestPropertiesConfiguration `pulumi:"configuration"`
	Description        *string                         `pulumi:"description"`
	Enabled            *bool                           `pulumi:"enabled"`
	Frequency          *int                            `pulumi:"frequency"`
	Kind               *WebTestKind                    `pulumi:"kind"`
	Location           *string                         `pulumi:"location"`
	Locations          []WebTestGeolocation            `pulumi:"locations"`
	ResourceGroupName  string                          `pulumi:"resourceGroupName"`
	RetryEnabled       *bool                           `pulumi:"retryEnabled"`
	SyntheticMonitorId string                          `pulumi:"syntheticMonitorId"`
	Tags               map[string]string               `pulumi:"tags"`
	Timeout            *int                            `pulumi:"timeout"`
	WebTestKind        WebTestKind                     `pulumi:"webTestKind"`
	WebTestName        *string                         `pulumi:"webTestName"`
}


type WebTestArgs struct {
	Configuration      WebTestPropertiesConfigurationPtrInput
	Description        pulumi.StringPtrInput
	Enabled            pulumi.BoolPtrInput
	Frequency          pulumi.IntPtrInput
	Kind               WebTestKindPtrInput
	Location           pulumi.StringPtrInput
	Locations          WebTestGeolocationArrayInput
	ResourceGroupName  pulumi.StringInput
	RetryEnabled       pulumi.BoolPtrInput
	SyntheticMonitorId pulumi.StringInput
	Tags               pulumi.StringMapInput
	Timeout            pulumi.IntPtrInput
	WebTestKind        WebTestKindInput
	WebTestName        pulumi.StringPtrInput
}

func (WebTestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webTestArgs)(nil)).Elem()
}

type WebTestInput interface {
	pulumi.Input

	ToWebTestOutput() WebTestOutput
	ToWebTestOutputWithContext(ctx context.Context) WebTestOutput
}

func (*WebTest) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTest)(nil)).Elem()
}

func (i *WebTest) ToWebTestOutput() WebTestOutput {
	return i.ToWebTestOutputWithContext(context.Background())
}

func (i *WebTest) ToWebTestOutputWithContext(ctx context.Context) WebTestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTestOutput)
}

type WebTestOutput struct{ *pulumi.OutputState }

func (WebTestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTest)(nil)).Elem()
}

func (o WebTestOutput) ToWebTestOutput() WebTestOutput {
	return o
}

func (o WebTestOutput) ToWebTestOutputWithContext(ctx context.Context) WebTestOutput {
	return o
}

func (o WebTestOutput) Configuration() WebTestPropertiesResponseConfigurationPtrOutput {
	return o.ApplyT(func(v *WebTest) WebTestPropertiesResponseConfigurationPtrOutput { return v.Configuration }).(WebTestPropertiesResponseConfigurationPtrOutput)
}

func (o WebTestOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebTest) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WebTestOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebTest) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o WebTestOutput) Frequency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebTest) pulumi.IntPtrOutput { return v.Frequency }).(pulumi.IntPtrOutput)
}

func (o WebTestOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebTest) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebTestOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WebTest) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o WebTestOutput) Locations() WebTestGeolocationResponseArrayOutput {
	return o.ApplyT(func(v *WebTest) WebTestGeolocationResponseArrayOutput { return v.Locations }).(WebTestGeolocationResponseArrayOutput)
}

func (o WebTestOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebTest) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebTestOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebTest) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WebTestOutput) RetryEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebTest) pulumi.BoolPtrOutput { return v.RetryEnabled }).(pulumi.BoolPtrOutput)
}

func (o WebTestOutput) SyntheticMonitorId() pulumi.StringOutput {
	return o.ApplyT(func(v *WebTest) pulumi.StringOutput { return v.SyntheticMonitorId }).(pulumi.StringOutput)
}

func (o WebTestOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebTest) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o WebTestOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebTest) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o WebTestOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebTest) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o WebTestOutput) WebTestKind() pulumi.StringOutput {
	return o.ApplyT(func(v *WebTest) pulumi.StringOutput { return v.WebTestKind }).(pulumi.StringOutput)
}

func (o WebTestOutput) WebTestName() pulumi.StringOutput {
	return o.ApplyT(func(v *WebTest) pulumi.StringOutput { return v.WebTestName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebTestOutput{})
}
