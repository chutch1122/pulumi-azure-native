


package v20180201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppVnetConnection struct {
	pulumi.CustomResourceState

	CertBlob       pulumi.StringPtrOutput       `pulumi:"certBlob"`
	CertThumbprint pulumi.StringOutput          `pulumi:"certThumbprint"`
	DnsServers     pulumi.StringPtrOutput       `pulumi:"dnsServers"`
	IsSwift        pulumi.BoolPtrOutput         `pulumi:"isSwift"`
	Kind           pulumi.StringPtrOutput       `pulumi:"kind"`
	Name           pulumi.StringOutput          `pulumi:"name"`
	ResyncRequired pulumi.BoolOutput            `pulumi:"resyncRequired"`
	Routes         VnetRouteResponseArrayOutput `pulumi:"routes"`
	Type           pulumi.StringOutput          `pulumi:"type"`
	VnetResourceId pulumi.StringPtrOutput       `pulumi:"vnetResourceId"`
}


func NewWebAppVnetConnection(ctx *pulumi.Context,
	name string, args *WebAppVnetConnectionArgs, opts ...pulumi.ResourceOption) (*WebAppVnetConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppVnetConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppVnetConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppVnetConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppVnetConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppVnetConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppVnetConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppVnetConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppVnetConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppVnetConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppVnetConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppVnetConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppVnetConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppVnetConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppVnetConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppVnetConnection
	err := ctx.RegisterResource("azure-native:web/v20180201:WebAppVnetConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppVnetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppVnetConnectionState, opts ...pulumi.ResourceOption) (*WebAppVnetConnection, error) {
	var resource WebAppVnetConnection
	err := ctx.ReadResource("azure-native:web/v20180201:WebAppVnetConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppVnetConnectionState struct {
}

type WebAppVnetConnectionState struct {
}

func (WebAppVnetConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppVnetConnectionState)(nil)).Elem()
}

type webAppVnetConnectionArgs struct {
	CertBlob          *string `pulumi:"certBlob"`
	DnsServers        *string `pulumi:"dnsServers"`
	IsSwift           *bool   `pulumi:"isSwift"`
	Kind              *string `pulumi:"kind"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	VnetName          *string `pulumi:"vnetName"`
	VnetResourceId    *string `pulumi:"vnetResourceId"`
}


type WebAppVnetConnectionArgs struct {
	CertBlob          pulumi.StringPtrInput
	DnsServers        pulumi.StringPtrInput
	IsSwift           pulumi.BoolPtrInput
	Kind              pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	VnetName          pulumi.StringPtrInput
	VnetResourceId    pulumi.StringPtrInput
}

func (WebAppVnetConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppVnetConnectionArgs)(nil)).Elem()
}

type WebAppVnetConnectionInput interface {
	pulumi.Input

	ToWebAppVnetConnectionOutput() WebAppVnetConnectionOutput
	ToWebAppVnetConnectionOutputWithContext(ctx context.Context) WebAppVnetConnectionOutput
}

func (*WebAppVnetConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppVnetConnection)(nil)).Elem()
}

func (i *WebAppVnetConnection) ToWebAppVnetConnectionOutput() WebAppVnetConnectionOutput {
	return i.ToWebAppVnetConnectionOutputWithContext(context.Background())
}

func (i *WebAppVnetConnection) ToWebAppVnetConnectionOutputWithContext(ctx context.Context) WebAppVnetConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppVnetConnectionOutput)
}

type WebAppVnetConnectionOutput struct{ *pulumi.OutputState }

func (WebAppVnetConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppVnetConnection)(nil)).Elem()
}

func (o WebAppVnetConnectionOutput) ToWebAppVnetConnectionOutput() WebAppVnetConnectionOutput {
	return o
}

func (o WebAppVnetConnectionOutput) ToWebAppVnetConnectionOutputWithContext(ctx context.Context) WebAppVnetConnectionOutput {
	return o
}

func (o WebAppVnetConnectionOutput) CertBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppVnetConnection) pulumi.StringPtrOutput { return v.CertBlob }).(pulumi.StringPtrOutput)
}

func (o WebAppVnetConnectionOutput) CertThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppVnetConnection) pulumi.StringOutput { return v.CertThumbprint }).(pulumi.StringOutput)
}

func (o WebAppVnetConnectionOutput) DnsServers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppVnetConnection) pulumi.StringPtrOutput { return v.DnsServers }).(pulumi.StringPtrOutput)
}

func (o WebAppVnetConnectionOutput) IsSwift() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppVnetConnection) pulumi.BoolPtrOutput { return v.IsSwift }).(pulumi.BoolPtrOutput)
}

func (o WebAppVnetConnectionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppVnetConnection) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppVnetConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppVnetConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppVnetConnectionOutput) ResyncRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v *WebAppVnetConnection) pulumi.BoolOutput { return v.ResyncRequired }).(pulumi.BoolOutput)
}

func (o WebAppVnetConnectionOutput) Routes() VnetRouteResponseArrayOutput {
	return o.ApplyT(func(v *WebAppVnetConnection) VnetRouteResponseArrayOutput { return v.Routes }).(VnetRouteResponseArrayOutput)
}

func (o WebAppVnetConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppVnetConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o WebAppVnetConnectionOutput) VnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppVnetConnection) pulumi.StringPtrOutput { return v.VnetResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppVnetConnectionOutput{})
}
