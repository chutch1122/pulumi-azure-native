


package v20210308

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnectionsAdtAPI struct {
	pulumi.CustomResourceState

	Name                              pulumi.StringOutput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrOutput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                             `pulumi:"provisioningState"`
	SystemData                        SystemDataResponseOutput                        `pulumi:"systemData"`
	Type                              pulumi.StringOutput                             `pulumi:"type"`
}


func NewPrivateEndpointConnectionsAdtAPI(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionsAdtAPIArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsAdtAPI, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateLinkServiceConnectionState == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkServiceConnectionState'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityandcompliance:PrivateEndpointConnectionsAdtAPI"),
		},
		{
			Type: pulumi.String("azure-native:securityandcompliance/v20210111:PrivateEndpointConnectionsAdtAPI"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnectionsAdtAPI
	err := ctx.RegisterResource("azure-native:securityandcompliance/v20210308:PrivateEndpointConnectionsAdtAPI", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnectionsAdtAPI(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionsAdtAPIState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsAdtAPI, error) {
	var resource PrivateEndpointConnectionsAdtAPI
	err := ctx.ReadResource("azure-native:securityandcompliance/v20210308:PrivateEndpointConnectionsAdtAPI", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateEndpointConnectionsAdtAPIState struct {
}

type PrivateEndpointConnectionsAdtAPIState struct {
}

func (PrivateEndpointConnectionsAdtAPIState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsAdtAPIState)(nil)).Elem()
}

type privateEndpointConnectionsAdtAPIArgs struct {
	PrivateEndpointConnectionName     *string                           `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                            `pulumi:"resourceGroupName"`
	ResourceName                      string                            `pulumi:"resourceName"`
}


type PrivateEndpointConnectionsAdtAPIArgs struct {
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	ResourceGroupName                 pulumi.StringInput
	ResourceName                      pulumi.StringInput
}

func (PrivateEndpointConnectionsAdtAPIArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsAdtAPIArgs)(nil)).Elem()
}

type PrivateEndpointConnectionsAdtAPIInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionsAdtAPIOutput() PrivateEndpointConnectionsAdtAPIOutput
	ToPrivateEndpointConnectionsAdtAPIOutputWithContext(ctx context.Context) PrivateEndpointConnectionsAdtAPIOutput
}

func (*PrivateEndpointConnectionsAdtAPI) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionsAdtAPI)(nil)).Elem()
}

func (i *PrivateEndpointConnectionsAdtAPI) ToPrivateEndpointConnectionsAdtAPIOutput() PrivateEndpointConnectionsAdtAPIOutput {
	return i.ToPrivateEndpointConnectionsAdtAPIOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionsAdtAPI) ToPrivateEndpointConnectionsAdtAPIOutputWithContext(ctx context.Context) PrivateEndpointConnectionsAdtAPIOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionsAdtAPIOutput)
}

type PrivateEndpointConnectionsAdtAPIOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionsAdtAPIOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionsAdtAPI)(nil)).Elem()
}

func (o PrivateEndpointConnectionsAdtAPIOutput) ToPrivateEndpointConnectionsAdtAPIOutput() PrivateEndpointConnectionsAdtAPIOutput {
	return o
}

func (o PrivateEndpointConnectionsAdtAPIOutput) ToPrivateEndpointConnectionsAdtAPIOutputWithContext(ctx context.Context) PrivateEndpointConnectionsAdtAPIOutput {
	return o
}

func (o PrivateEndpointConnectionsAdtAPIOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsAdtAPI) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionsAdtAPIOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsAdtAPI) PrivateEndpointResponsePtrOutput { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionsAdtAPIOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsAdtAPI) PrivateLinkServiceConnectionStateResponseOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionsAdtAPIOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsAdtAPI) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionsAdtAPIOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsAdtAPI) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionsAdtAPIOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsAdtAPI) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionsAdtAPIOutput{})
}
