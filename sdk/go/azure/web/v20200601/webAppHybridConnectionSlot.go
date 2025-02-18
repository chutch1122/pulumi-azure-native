


package v20200601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppHybridConnectionSlot struct {
	pulumi.CustomResourceState

	Hostname            pulumi.StringPtrOutput `pulumi:"hostname"`
	Kind                pulumi.StringPtrOutput `pulumi:"kind"`
	Name                pulumi.StringOutput    `pulumi:"name"`
	Port                pulumi.IntPtrOutput    `pulumi:"port"`
	RelayArmUri         pulumi.StringPtrOutput `pulumi:"relayArmUri"`
	RelayName           pulumi.StringPtrOutput `pulumi:"relayName"`
	SendKeyName         pulumi.StringPtrOutput `pulumi:"sendKeyName"`
	SendKeyValue        pulumi.StringPtrOutput `pulumi:"sendKeyValue"`
	ServiceBusNamespace pulumi.StringPtrOutput `pulumi:"serviceBusNamespace"`
	ServiceBusSuffix    pulumi.StringPtrOutput `pulumi:"serviceBusSuffix"`
	Type                pulumi.StringOutput    `pulumi:"type"`
}


func NewWebAppHybridConnectionSlot(ctx *pulumi.Context,
	name string, args *WebAppHybridConnectionSlotArgs, opts ...pulumi.ResourceOption) (*WebAppHybridConnectionSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppHybridConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppHybridConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppHybridConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppHybridConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppHybridConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppHybridConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppHybridConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppHybridConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppHybridConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppHybridConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppHybridConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppHybridConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppHybridConnectionSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppHybridConnectionSlot
	err := ctx.RegisterResource("azure-native:web/v20200601:WebAppHybridConnectionSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppHybridConnectionSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppHybridConnectionSlotState, opts ...pulumi.ResourceOption) (*WebAppHybridConnectionSlot, error) {
	var resource WebAppHybridConnectionSlot
	err := ctx.ReadResource("azure-native:web/v20200601:WebAppHybridConnectionSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppHybridConnectionSlotState struct {
}

type WebAppHybridConnectionSlotState struct {
}

func (WebAppHybridConnectionSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppHybridConnectionSlotState)(nil)).Elem()
}

type webAppHybridConnectionSlotArgs struct {
	Hostname            *string `pulumi:"hostname"`
	Kind                *string `pulumi:"kind"`
	Name                string  `pulumi:"name"`
	NamespaceName       string  `pulumi:"namespaceName"`
	Port                *int    `pulumi:"port"`
	RelayArmUri         *string `pulumi:"relayArmUri"`
	RelayName           *string `pulumi:"relayName"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
	SendKeyName         *string `pulumi:"sendKeyName"`
	SendKeyValue        *string `pulumi:"sendKeyValue"`
	ServiceBusNamespace *string `pulumi:"serviceBusNamespace"`
	ServiceBusSuffix    *string `pulumi:"serviceBusSuffix"`
	Slot                string  `pulumi:"slot"`
}


type WebAppHybridConnectionSlotArgs struct {
	Hostname            pulumi.StringPtrInput
	Kind                pulumi.StringPtrInput
	Name                pulumi.StringInput
	NamespaceName       pulumi.StringInput
	Port                pulumi.IntPtrInput
	RelayArmUri         pulumi.StringPtrInput
	RelayName           pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	SendKeyName         pulumi.StringPtrInput
	SendKeyValue        pulumi.StringPtrInput
	ServiceBusNamespace pulumi.StringPtrInput
	ServiceBusSuffix    pulumi.StringPtrInput
	Slot                pulumi.StringInput
}

func (WebAppHybridConnectionSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppHybridConnectionSlotArgs)(nil)).Elem()
}

type WebAppHybridConnectionSlotInput interface {
	pulumi.Input

	ToWebAppHybridConnectionSlotOutput() WebAppHybridConnectionSlotOutput
	ToWebAppHybridConnectionSlotOutputWithContext(ctx context.Context) WebAppHybridConnectionSlotOutput
}

func (*WebAppHybridConnectionSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppHybridConnectionSlot)(nil)).Elem()
}

func (i *WebAppHybridConnectionSlot) ToWebAppHybridConnectionSlotOutput() WebAppHybridConnectionSlotOutput {
	return i.ToWebAppHybridConnectionSlotOutputWithContext(context.Background())
}

func (i *WebAppHybridConnectionSlot) ToWebAppHybridConnectionSlotOutputWithContext(ctx context.Context) WebAppHybridConnectionSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppHybridConnectionSlotOutput)
}

type WebAppHybridConnectionSlotOutput struct{ *pulumi.OutputState }

func (WebAppHybridConnectionSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppHybridConnectionSlot)(nil)).Elem()
}

func (o WebAppHybridConnectionSlotOutput) ToWebAppHybridConnectionSlotOutput() WebAppHybridConnectionSlotOutput {
	return o
}

func (o WebAppHybridConnectionSlotOutput) ToWebAppHybridConnectionSlotOutputWithContext(ctx context.Context) WebAppHybridConnectionSlotOutput {
	return o
}

func (o WebAppHybridConnectionSlotOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnectionSlot) pulumi.StringPtrOutput { return v.Hostname }).(pulumi.StringPtrOutput)
}

func (o WebAppHybridConnectionSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnectionSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppHybridConnectionSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppHybridConnectionSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppHybridConnectionSlotOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnectionSlot) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

func (o WebAppHybridConnectionSlotOutput) RelayArmUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnectionSlot) pulumi.StringPtrOutput { return v.RelayArmUri }).(pulumi.StringPtrOutput)
}

func (o WebAppHybridConnectionSlotOutput) RelayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnectionSlot) pulumi.StringPtrOutput { return v.RelayName }).(pulumi.StringPtrOutput)
}

func (o WebAppHybridConnectionSlotOutput) SendKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnectionSlot) pulumi.StringPtrOutput { return v.SendKeyName }).(pulumi.StringPtrOutput)
}

func (o WebAppHybridConnectionSlotOutput) SendKeyValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnectionSlot) pulumi.StringPtrOutput { return v.SendKeyValue }).(pulumi.StringPtrOutput)
}

func (o WebAppHybridConnectionSlotOutput) ServiceBusNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnectionSlot) pulumi.StringPtrOutput { return v.ServiceBusNamespace }).(pulumi.StringPtrOutput)
}

func (o WebAppHybridConnectionSlotOutput) ServiceBusSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHybridConnectionSlot) pulumi.StringPtrOutput { return v.ServiceBusSuffix }).(pulumi.StringPtrOutput)
}

func (o WebAppHybridConnectionSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppHybridConnectionSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppHybridConnectionSlotOutput{})
}
