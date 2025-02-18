


package servicefabricmesh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Gateway struct {
	pulumi.CustomResourceState

	Description        pulumi.StringPtrOutput        `pulumi:"description"`
	DestinationNetwork NetworkRefResponseOutput      `pulumi:"destinationNetwork"`
	Http               HttpConfigResponseArrayOutput `pulumi:"http"`
	IpAddress          pulumi.StringOutput           `pulumi:"ipAddress"`
	Location           pulumi.StringOutput           `pulumi:"location"`
	Name               pulumi.StringOutput           `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput           `pulumi:"provisioningState"`
	SourceNetwork      NetworkRefResponseOutput      `pulumi:"sourceNetwork"`
	Status             pulumi.StringOutput           `pulumi:"status"`
	StatusDetails      pulumi.StringOutput           `pulumi:"statusDetails"`
	Tags               pulumi.StringMapOutput        `pulumi:"tags"`
	Tcp                TcpConfigResponseArrayOutput  `pulumi:"tcp"`
	Type               pulumi.StringOutput           `pulumi:"type"`
}


func NewGateway(ctx *pulumi.Context,
	name string, args *GatewayArgs, opts ...pulumi.ResourceOption) (*Gateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationNetwork == nil {
		return nil, errors.New("invalid value for required argument 'DestinationNetwork'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SourceNetwork == nil {
		return nil, errors.New("invalid value for required argument 'SourceNetwork'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabricmesh/v20180901preview:Gateway"),
		},
	})
	opts = append(opts, aliases)
	var resource Gateway
	err := ctx.RegisterResource("azure-native:servicefabricmesh:Gateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayState, opts ...pulumi.ResourceOption) (*Gateway, error) {
	var resource Gateway
	err := ctx.ReadResource("azure-native:servicefabricmesh:Gateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type gatewayState struct {
}

type GatewayState struct {
}

func (GatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayState)(nil)).Elem()
}

type gatewayArgs struct {
	Description         *string           `pulumi:"description"`
	DestinationNetwork  NetworkRef        `pulumi:"destinationNetwork"`
	GatewayResourceName *string           `pulumi:"gatewayResourceName"`
	Http                []HttpConfig      `pulumi:"http"`
	Location            *string           `pulumi:"location"`
	ResourceGroupName   string            `pulumi:"resourceGroupName"`
	SourceNetwork       NetworkRef        `pulumi:"sourceNetwork"`
	Tags                map[string]string `pulumi:"tags"`
	Tcp                 []TcpConfig       `pulumi:"tcp"`
}


type GatewayArgs struct {
	Description         pulumi.StringPtrInput
	DestinationNetwork  NetworkRefInput
	GatewayResourceName pulumi.StringPtrInput
	Http                HttpConfigArrayInput
	Location            pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	SourceNetwork       NetworkRefInput
	Tags                pulumi.StringMapInput
	Tcp                 TcpConfigArrayInput
}

func (GatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayArgs)(nil)).Elem()
}

type GatewayInput interface {
	pulumi.Input

	ToGatewayOutput() GatewayOutput
	ToGatewayOutputWithContext(ctx context.Context) GatewayOutput
}

func (*Gateway) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (i *Gateway) ToGatewayOutput() GatewayOutput {
	return i.ToGatewayOutputWithContext(context.Background())
}

func (i *Gateway) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayOutput)
}

type GatewayOutput struct{ *pulumi.OutputState }

func (GatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (o GatewayOutput) ToGatewayOutput() GatewayOutput {
	return o
}

func (o GatewayOutput) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return o
}

func (o GatewayOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GatewayOutput) DestinationNetwork() NetworkRefResponseOutput {
	return o.ApplyT(func(v *Gateway) NetworkRefResponseOutput { return v.DestinationNetwork }).(NetworkRefResponseOutput)
}

func (o GatewayOutput) Http() HttpConfigResponseArrayOutput {
	return o.ApplyT(func(v *Gateway) HttpConfigResponseArrayOutput { return v.Http }).(HttpConfigResponseArrayOutput)
}

func (o GatewayOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

func (o GatewayOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o GatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GatewayOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GatewayOutput) SourceNetwork() NetworkRefResponseOutput {
	return o.ApplyT(func(v *Gateway) NetworkRefResponseOutput { return v.SourceNetwork }).(NetworkRefResponseOutput)
}

func (o GatewayOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o GatewayOutput) StatusDetails() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.StatusDetails }).(pulumi.StringOutput)
}

func (o GatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GatewayOutput) Tcp() TcpConfigResponseArrayOutput {
	return o.ApplyT(func(v *Gateway) TcpConfigResponseArrayOutput { return v.Tcp }).(TcpConfigResponseArrayOutput)
}

func (o GatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewayOutput{})
}
