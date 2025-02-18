


package v20171001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Webhook struct {
	pulumi.CustomResourceState

	Actions           pulumi.StringArrayOutput `pulumi:"actions"`
	Location          pulumi.StringOutput      `pulumi:"location"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	Scope             pulumi.StringPtrOutput   `pulumi:"scope"`
	Status            pulumi.StringPtrOutput   `pulumi:"status"`
	Tags              pulumi.StringMapOutput   `pulumi:"tags"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewWebhook(ctx *pulumi.Context,
	name string, args *WebhookArgs, opts ...pulumi.ResourceOption) (*Webhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceUri == nil {
		return nil, errors.New("invalid value for required argument 'ServiceUri'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20170601preview:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20190501:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20191201preview:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20201101preview:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210601preview:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210801preview:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210901:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20211201preview:Webhook"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20220201preview:Webhook"),
		},
	})
	opts = append(opts, aliases)
	var resource Webhook
	err := ctx.RegisterResource("azure-native:containerregistry/v20171001:Webhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebhookState, opts ...pulumi.ResourceOption) (*Webhook, error) {
	var resource Webhook
	err := ctx.ReadResource("azure-native:containerregistry/v20171001:Webhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webhookState struct {
}

type WebhookState struct {
}

func (WebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookState)(nil)).Elem()
}

type webhookArgs struct {
	Actions           []string          `pulumi:"actions"`
	CustomHeaders     map[string]string `pulumi:"customHeaders"`
	Location          *string           `pulumi:"location"`
	RegistryName      string            `pulumi:"registryName"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Scope             *string           `pulumi:"scope"`
	ServiceUri        string            `pulumi:"serviceUri"`
	Status            *string           `pulumi:"status"`
	Tags              map[string]string `pulumi:"tags"`
	WebhookName       *string           `pulumi:"webhookName"`
}


type WebhookArgs struct {
	Actions           pulumi.StringArrayInput
	CustomHeaders     pulumi.StringMapInput
	Location          pulumi.StringPtrInput
	RegistryName      pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Scope             pulumi.StringPtrInput
	ServiceUri        pulumi.StringInput
	Status            pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	WebhookName       pulumi.StringPtrInput
}

func (WebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookArgs)(nil)).Elem()
}

type WebhookInput interface {
	pulumi.Input

	ToWebhookOutput() WebhookOutput
	ToWebhookOutputWithContext(ctx context.Context) WebhookOutput
}

func (*Webhook) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhook)(nil)).Elem()
}

func (i *Webhook) ToWebhookOutput() WebhookOutput {
	return i.ToWebhookOutputWithContext(context.Background())
}

func (i *Webhook) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookOutput)
}

type WebhookOutput struct{ *pulumi.OutputState }

func (WebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhook)(nil)).Elem()
}

func (o WebhookOutput) ToWebhookOutput() WebhookOutput {
	return o
}

func (o WebhookOutput) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return o
}

func (o WebhookOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringArrayOutput { return v.Actions }).(pulumi.StringArrayOutput)
}

func (o WebhookOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o WebhookOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebhookOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WebhookOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o WebhookOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o WebhookOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o WebhookOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhook) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebhookOutput{})
}
