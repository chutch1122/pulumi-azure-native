


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppAzureStorageAccountsSlot struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringPtrOutput                 `pulumi:"kind"`
	Name       pulumi.StringOutput                    `pulumi:"name"`
	Properties AzureStorageInfoValueResponseMapOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                    `pulumi:"type"`
}


func NewWebAppAzureStorageAccountsSlot(ctx *pulumi.Context,
	name string, args *WebAppAzureStorageAccountsSlotArgs, opts ...pulumi.ResourceOption) (*WebAppAzureStorageAccountsSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppAzureStorageAccountsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppAzureStorageAccountsSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppAzureStorageAccountsSlot
	err := ctx.RegisterResource("azure-native:web/v20210101:WebAppAzureStorageAccountsSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppAzureStorageAccountsSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppAzureStorageAccountsSlotState, opts ...pulumi.ResourceOption) (*WebAppAzureStorageAccountsSlot, error) {
	var resource WebAppAzureStorageAccountsSlot
	err := ctx.ReadResource("azure-native:web/v20210101:WebAppAzureStorageAccountsSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppAzureStorageAccountsSlotState struct {
}

type WebAppAzureStorageAccountsSlotState struct {
}

func (WebAppAzureStorageAccountsSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAzureStorageAccountsSlotState)(nil)).Elem()
}

type webAppAzureStorageAccountsSlotArgs struct {
	Kind              *string                          `pulumi:"kind"`
	Name              string                           `pulumi:"name"`
	Properties        map[string]AzureStorageInfoValue `pulumi:"properties"`
	ResourceGroupName string                           `pulumi:"resourceGroupName"`
	Slot              string                           `pulumi:"slot"`
}


type WebAppAzureStorageAccountsSlotArgs struct {
	Kind              pulumi.StringPtrInput
	Name              pulumi.StringInput
	Properties        AzureStorageInfoValueMapInput
	ResourceGroupName pulumi.StringInput
	Slot              pulumi.StringInput
}

func (WebAppAzureStorageAccountsSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAzureStorageAccountsSlotArgs)(nil)).Elem()
}

type WebAppAzureStorageAccountsSlotInput interface {
	pulumi.Input

	ToWebAppAzureStorageAccountsSlotOutput() WebAppAzureStorageAccountsSlotOutput
	ToWebAppAzureStorageAccountsSlotOutputWithContext(ctx context.Context) WebAppAzureStorageAccountsSlotOutput
}

func (*WebAppAzureStorageAccountsSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppAzureStorageAccountsSlot)(nil)).Elem()
}

func (i *WebAppAzureStorageAccountsSlot) ToWebAppAzureStorageAccountsSlotOutput() WebAppAzureStorageAccountsSlotOutput {
	return i.ToWebAppAzureStorageAccountsSlotOutputWithContext(context.Background())
}

func (i *WebAppAzureStorageAccountsSlot) ToWebAppAzureStorageAccountsSlotOutputWithContext(ctx context.Context) WebAppAzureStorageAccountsSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppAzureStorageAccountsSlotOutput)
}

type WebAppAzureStorageAccountsSlotOutput struct{ *pulumi.OutputState }

func (WebAppAzureStorageAccountsSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppAzureStorageAccountsSlot)(nil)).Elem()
}

func (o WebAppAzureStorageAccountsSlotOutput) ToWebAppAzureStorageAccountsSlotOutput() WebAppAzureStorageAccountsSlotOutput {
	return o
}

func (o WebAppAzureStorageAccountsSlotOutput) ToWebAppAzureStorageAccountsSlotOutputWithContext(ctx context.Context) WebAppAzureStorageAccountsSlotOutput {
	return o
}

func (o WebAppAzureStorageAccountsSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccountsSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppAzureStorageAccountsSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccountsSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppAzureStorageAccountsSlotOutput) Properties() AzureStorageInfoValueResponseMapOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccountsSlot) AzureStorageInfoValueResponseMapOutput { return v.Properties }).(AzureStorageInfoValueResponseMapOutput)
}

func (o WebAppAzureStorageAccountsSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccountsSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppAzureStorageAccountsSlotOutput{})
}
