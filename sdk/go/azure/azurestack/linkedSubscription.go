


package azurestack

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LinkedSubscription struct {
	pulumi.CustomResourceState

	DeviceConnectionStatus pulumi.StringOutput      `pulumi:"deviceConnectionStatus"`
	DeviceId               pulumi.StringOutput      `pulumi:"deviceId"`
	DeviceLinkState        pulumi.StringOutput      `pulumi:"deviceLinkState"`
	DeviceObjectId         pulumi.StringOutput      `pulumi:"deviceObjectId"`
	Etag                   pulumi.StringPtrOutput   `pulumi:"etag"`
	Kind                   pulumi.StringOutput      `pulumi:"kind"`
	LastConnectedTime      pulumi.StringOutput      `pulumi:"lastConnectedTime"`
	LinkedSubscriptionId   pulumi.StringPtrOutput   `pulumi:"linkedSubscriptionId"`
	Location               pulumi.StringOutput      `pulumi:"location"`
	Name                   pulumi.StringOutput      `pulumi:"name"`
	RegistrationResourceId pulumi.StringPtrOutput   `pulumi:"registrationResourceId"`
	SystemData             SystemDataResponseOutput `pulumi:"systemData"`
	Tags                   pulumi.StringMapOutput   `pulumi:"tags"`
	Type                   pulumi.StringOutput      `pulumi:"type"`
}


func NewLinkedSubscription(ctx *pulumi.Context,
	name string, args *LinkedSubscriptionArgs, opts ...pulumi.ResourceOption) (*LinkedSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LinkedSubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'LinkedSubscriptionId'")
	}
	if args.RegistrationResourceId == nil {
		return nil, errors.New("invalid value for required argument 'RegistrationResourceId'")
	}
	if args.ResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroup'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestack/v20200601preview:LinkedSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource LinkedSubscription
	err := ctx.RegisterResource("azure-native:azurestack:LinkedSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLinkedSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedSubscriptionState, opts ...pulumi.ResourceOption) (*LinkedSubscription, error) {
	var resource LinkedSubscription
	err := ctx.ReadResource("azure-native:azurestack:LinkedSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type linkedSubscriptionState struct {
}

type LinkedSubscriptionState struct {
}

func (LinkedSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedSubscriptionState)(nil)).Elem()
}

type linkedSubscriptionArgs struct {
	LinkedSubscriptionId   string  `pulumi:"linkedSubscriptionId"`
	LinkedSubscriptionName *string `pulumi:"linkedSubscriptionName"`
	Location               *string `pulumi:"location"`
	RegistrationResourceId string  `pulumi:"registrationResourceId"`
	ResourceGroup          string  `pulumi:"resourceGroup"`
}


type LinkedSubscriptionArgs struct {
	LinkedSubscriptionId   pulumi.StringInput
	LinkedSubscriptionName pulumi.StringPtrInput
	Location               pulumi.StringPtrInput
	RegistrationResourceId pulumi.StringInput
	ResourceGroup          pulumi.StringInput
}

func (LinkedSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedSubscriptionArgs)(nil)).Elem()
}

type LinkedSubscriptionInput interface {
	pulumi.Input

	ToLinkedSubscriptionOutput() LinkedSubscriptionOutput
	ToLinkedSubscriptionOutputWithContext(ctx context.Context) LinkedSubscriptionOutput
}

func (*LinkedSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedSubscription)(nil)).Elem()
}

func (i *LinkedSubscription) ToLinkedSubscriptionOutput() LinkedSubscriptionOutput {
	return i.ToLinkedSubscriptionOutputWithContext(context.Background())
}

func (i *LinkedSubscription) ToLinkedSubscriptionOutputWithContext(ctx context.Context) LinkedSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedSubscriptionOutput)
}

type LinkedSubscriptionOutput struct{ *pulumi.OutputState }

func (LinkedSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedSubscription)(nil)).Elem()
}

func (o LinkedSubscriptionOutput) ToLinkedSubscriptionOutput() LinkedSubscriptionOutput {
	return o
}

func (o LinkedSubscriptionOutput) ToLinkedSubscriptionOutputWithContext(ctx context.Context) LinkedSubscriptionOutput {
	return o
}

func (o LinkedSubscriptionOutput) DeviceConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringOutput { return v.DeviceConnectionStatus }).(pulumi.StringOutput)
}

func (o LinkedSubscriptionOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringOutput { return v.DeviceId }).(pulumi.StringOutput)
}

func (o LinkedSubscriptionOutput) DeviceLinkState() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringOutput { return v.DeviceLinkState }).(pulumi.StringOutput)
}

func (o LinkedSubscriptionOutput) DeviceObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringOutput { return v.DeviceObjectId }).(pulumi.StringOutput)
}

func (o LinkedSubscriptionOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LinkedSubscriptionOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o LinkedSubscriptionOutput) LastConnectedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringOutput { return v.LastConnectedTime }).(pulumi.StringOutput)
}

func (o LinkedSubscriptionOutput) LinkedSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringPtrOutput { return v.LinkedSubscriptionId }).(pulumi.StringPtrOutput)
}

func (o LinkedSubscriptionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o LinkedSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LinkedSubscriptionOutput) RegistrationResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringPtrOutput { return v.RegistrationResourceId }).(pulumi.StringPtrOutput)
}

func (o LinkedSubscriptionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LinkedSubscription) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LinkedSubscriptionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LinkedSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedSubscriptionOutput{})
}
