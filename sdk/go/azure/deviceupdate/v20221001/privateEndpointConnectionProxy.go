


package v20221001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnectionProxy struct {
	pulumi.CustomResourceState

	ETag                  pulumi.StringOutput                    `pulumi:"eTag"`
	Name                  pulumi.StringOutput                    `pulumi:"name"`
	ProvisioningState     pulumi.StringOutput                    `pulumi:"provisioningState"`
	RemotePrivateEndpoint RemotePrivateEndpointResponsePtrOutput `pulumi:"remotePrivateEndpoint"`
	Status                pulumi.StringPtrOutput                 `pulumi:"status"`
	SystemData            SystemDataResponseOutput               `pulumi:"systemData"`
	Type                  pulumi.StringOutput                    `pulumi:"type"`
}


func NewPrivateEndpointConnectionProxy(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionProxyArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:deviceupdate:PrivateEndpointConnectionProxy"),
		},
		{
			Type: pulumi.String("azure-native:deviceupdate/v20200301preview:PrivateEndpointConnectionProxy"),
		},
		{
			Type: pulumi.String("azure-native:deviceupdate/v20220401preview:PrivateEndpointConnectionProxy"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnectionProxy
	err := ctx.RegisterResource("azure-native:deviceupdate/v20221001:PrivateEndpointConnectionProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnectionProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionProxyState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionProxy, error) {
	var resource PrivateEndpointConnectionProxy
	err := ctx.ReadResource("azure-native:deviceupdate/v20221001:PrivateEndpointConnectionProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateEndpointConnectionProxyState struct {
}

type PrivateEndpointConnectionProxyState struct {
}

func (PrivateEndpointConnectionProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionProxyState)(nil)).Elem()
}

type privateEndpointConnectionProxyArgs struct {
	AccountName                      string                 `pulumi:"accountName"`
	PrivateEndpointConnectionProxyId *string                `pulumi:"privateEndpointConnectionProxyId"`
	RemotePrivateEndpoint            *RemotePrivateEndpoint `pulumi:"remotePrivateEndpoint"`
	ResourceGroupName                string                 `pulumi:"resourceGroupName"`
	Status                           *string                `pulumi:"status"`
}


type PrivateEndpointConnectionProxyArgs struct {
	AccountName                      pulumi.StringInput
	PrivateEndpointConnectionProxyId pulumi.StringPtrInput
	RemotePrivateEndpoint            RemotePrivateEndpointPtrInput
	ResourceGroupName                pulumi.StringInput
	Status                           pulumi.StringPtrInput
}

func (PrivateEndpointConnectionProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionProxyArgs)(nil)).Elem()
}

type PrivateEndpointConnectionProxyInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionProxyOutput() PrivateEndpointConnectionProxyOutput
	ToPrivateEndpointConnectionProxyOutputWithContext(ctx context.Context) PrivateEndpointConnectionProxyOutput
}

func (*PrivateEndpointConnectionProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProxy)(nil)).Elem()
}

func (i *PrivateEndpointConnectionProxy) ToPrivateEndpointConnectionProxyOutput() PrivateEndpointConnectionProxyOutput {
	return i.ToPrivateEndpointConnectionProxyOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionProxy) ToPrivateEndpointConnectionProxyOutputWithContext(ctx context.Context) PrivateEndpointConnectionProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionProxyOutput)
}

type PrivateEndpointConnectionProxyOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProxy)(nil)).Elem()
}

func (o PrivateEndpointConnectionProxyOutput) ToPrivateEndpointConnectionProxyOutput() PrivateEndpointConnectionProxyOutput {
	return o
}

func (o PrivateEndpointConnectionProxyOutput) ToPrivateEndpointConnectionProxyOutputWithContext(ctx context.Context) PrivateEndpointConnectionProxyOutput {
	return o
}

func (o PrivateEndpointConnectionProxyOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxy) pulumi.StringOutput { return v.ETag }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionProxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionProxyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionProxyOutput) RemotePrivateEndpoint() RemotePrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxy) RemotePrivateEndpointResponsePtrOutput {
		return v.RemotePrivateEndpoint
	}).(RemotePrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionProxyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxy) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionProxyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionProxyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionProxyOutput{})
}
