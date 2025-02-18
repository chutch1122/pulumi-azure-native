


package v20150615

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type VirtualNetworkGatewayConnection struct {
	pulumi.CustomResourceState

	AuthorizationKey        pulumi.StringPtrOutput                 `pulumi:"authorizationKey"`
	ConnectionStatus        pulumi.StringPtrOutput                 `pulumi:"connectionStatus"`
	ConnectionType          pulumi.StringPtrOutput                 `pulumi:"connectionType"`
	EgressBytesTransferred  pulumi.Float64PtrOutput                `pulumi:"egressBytesTransferred"`
	EnableBgp               pulumi.BoolPtrOutput                   `pulumi:"enableBgp"`
	Etag                    pulumi.StringPtrOutput                 `pulumi:"etag"`
	IngressBytesTransferred pulumi.Float64PtrOutput                `pulumi:"ingressBytesTransferred"`
	LocalNetworkGateway2    LocalNetworkGatewayResponsePtrOutput   `pulumi:"localNetworkGateway2"`
	Location                pulumi.StringPtrOutput                 `pulumi:"location"`
	Name                    pulumi.StringOutput                    `pulumi:"name"`
	Peer                    SubResourceResponsePtrOutput           `pulumi:"peer"`
	ProvisioningState       pulumi.StringPtrOutput                 `pulumi:"provisioningState"`
	ResourceGuid            pulumi.StringPtrOutput                 `pulumi:"resourceGuid"`
	RoutingWeight           pulumi.IntPtrOutput                    `pulumi:"routingWeight"`
	SharedKey               pulumi.StringPtrOutput                 `pulumi:"sharedKey"`
	Tags                    pulumi.StringMapOutput                 `pulumi:"tags"`
	Type                    pulumi.StringOutput                    `pulumi:"type"`
	VirtualNetworkGateway1  VirtualNetworkGatewayResponsePtrOutput `pulumi:"virtualNetworkGateway1"`
	VirtualNetworkGateway2  VirtualNetworkGatewayResponsePtrOutput `pulumi:"virtualNetworkGateway2"`
}


func NewVirtualNetworkGatewayConnection(ctx *pulumi.Context,
	name string, args *VirtualNetworkGatewayConnectionArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkGatewayConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VirtualNetworkGatewayConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkGatewayConnection
	err := ctx.RegisterResource("azure-native:network/v20150615:VirtualNetworkGatewayConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetworkGatewayConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkGatewayConnectionState, opts ...pulumi.ResourceOption) (*VirtualNetworkGatewayConnection, error) {
	var resource VirtualNetworkGatewayConnection
	err := ctx.ReadResource("azure-native:network/v20150615:VirtualNetworkGatewayConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualNetworkGatewayConnectionState struct {
}

type VirtualNetworkGatewayConnectionState struct {
}

func (VirtualNetworkGatewayConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayConnectionState)(nil)).Elem()
}

type virtualNetworkGatewayConnectionArgs struct {
	AuthorizationKey                    *string                    `pulumi:"authorizationKey"`
	ConnectionStatus                    *string                    `pulumi:"connectionStatus"`
	ConnectionType                      *string                    `pulumi:"connectionType"`
	EgressBytesTransferred              *float64                   `pulumi:"egressBytesTransferred"`
	EnableBgp                           *bool                      `pulumi:"enableBgp"`
	Id                                  *string                    `pulumi:"id"`
	IngressBytesTransferred             *float64                   `pulumi:"ingressBytesTransferred"`
	LocalNetworkGateway2                *LocalNetworkGatewayType   `pulumi:"localNetworkGateway2"`
	Location                            *string                    `pulumi:"location"`
	Peer                                *SubResource               `pulumi:"peer"`
	ProvisioningState                   *string                    `pulumi:"provisioningState"`
	ResourceGroupName                   string                     `pulumi:"resourceGroupName"`
	ResourceGuid                        *string                    `pulumi:"resourceGuid"`
	RoutingWeight                       *int                       `pulumi:"routingWeight"`
	SharedKey                           *string                    `pulumi:"sharedKey"`
	Tags                                map[string]string          `pulumi:"tags"`
	VirtualNetworkGateway1              *VirtualNetworkGatewayType `pulumi:"virtualNetworkGateway1"`
	VirtualNetworkGateway2              *VirtualNetworkGatewayType `pulumi:"virtualNetworkGateway2"`
	VirtualNetworkGatewayConnectionName *string                    `pulumi:"virtualNetworkGatewayConnectionName"`
}


type VirtualNetworkGatewayConnectionArgs struct {
	AuthorizationKey                    pulumi.StringPtrInput
	ConnectionStatus                    pulumi.StringPtrInput
	ConnectionType                      pulumi.StringPtrInput
	EgressBytesTransferred              pulumi.Float64PtrInput
	EnableBgp                           pulumi.BoolPtrInput
	Id                                  pulumi.StringPtrInput
	IngressBytesTransferred             pulumi.Float64PtrInput
	LocalNetworkGateway2                LocalNetworkGatewayTypePtrInput
	Location                            pulumi.StringPtrInput
	Peer                                SubResourcePtrInput
	ProvisioningState                   pulumi.StringPtrInput
	ResourceGroupName                   pulumi.StringInput
	ResourceGuid                        pulumi.StringPtrInput
	RoutingWeight                       pulumi.IntPtrInput
	SharedKey                           pulumi.StringPtrInput
	Tags                                pulumi.StringMapInput
	VirtualNetworkGateway1              VirtualNetworkGatewayTypePtrInput
	VirtualNetworkGateway2              VirtualNetworkGatewayTypePtrInput
	VirtualNetworkGatewayConnectionName pulumi.StringPtrInput
}

func (VirtualNetworkGatewayConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayConnectionArgs)(nil)).Elem()
}

type VirtualNetworkGatewayConnectionInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayConnectionOutput() VirtualNetworkGatewayConnectionOutput
	ToVirtualNetworkGatewayConnectionOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionOutput
}

func (*VirtualNetworkGatewayConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayConnection)(nil)).Elem()
}

func (i *VirtualNetworkGatewayConnection) ToVirtualNetworkGatewayConnectionOutput() VirtualNetworkGatewayConnectionOutput {
	return i.ToVirtualNetworkGatewayConnectionOutputWithContext(context.Background())
}

func (i *VirtualNetworkGatewayConnection) ToVirtualNetworkGatewayConnectionOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayConnectionOutput)
}

type VirtualNetworkGatewayConnectionOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayConnection)(nil)).Elem()
}

func (o VirtualNetworkGatewayConnectionOutput) ToVirtualNetworkGatewayConnectionOutput() VirtualNetworkGatewayConnectionOutput {
	return o
}

func (o VirtualNetworkGatewayConnectionOutput) ToVirtualNetworkGatewayConnectionOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionOutput {
	return o
}

func (o VirtualNetworkGatewayConnectionOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringPtrOutput { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) ConnectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringPtrOutput { return v.ConnectionStatus }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) ConnectionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringPtrOutput { return v.ConnectionType }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) EgressBytesTransferred() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.Float64PtrOutput { return v.EgressBytesTransferred }).(pulumi.Float64PtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.BoolPtrOutput { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) IngressBytesTransferred() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.Float64PtrOutput { return v.IngressBytesTransferred }).(pulumi.Float64PtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) LocalNetworkGateway2() LocalNetworkGatewayResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) LocalNetworkGatewayResponsePtrOutput {
		return v.LocalNetworkGateway2
	}).(LocalNetworkGatewayResponsePtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) Peer() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) SubResourceResponsePtrOutput { return v.Peer }).(SubResourceResponsePtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringPtrOutput { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.IntPtrOutput { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringPtrOutput { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) VirtualNetworkGateway1() VirtualNetworkGatewayResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) VirtualNetworkGatewayResponsePtrOutput {
		return v.VirtualNetworkGateway1
	}).(VirtualNetworkGatewayResponsePtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) VirtualNetworkGateway2() VirtualNetworkGatewayResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) VirtualNetworkGatewayResponsePtrOutput {
		return v.VirtualNetworkGateway2
	}).(VirtualNetworkGatewayResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkGatewayConnectionOutput{})
}
