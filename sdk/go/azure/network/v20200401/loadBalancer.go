


package v20200401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LoadBalancer struct {
	pulumi.CustomResourceState

	BackendAddressPools      BackendAddressPoolResponseArrayOutput      `pulumi:"backendAddressPools"`
	Etag                     pulumi.StringOutput                        `pulumi:"etag"`
	FrontendIPConfigurations FrontendIPConfigurationResponseArrayOutput `pulumi:"frontendIPConfigurations"`
	InboundNatPools          InboundNatPoolResponseArrayOutput          `pulumi:"inboundNatPools"`
	InboundNatRules          InboundNatRuleResponseArrayOutput          `pulumi:"inboundNatRules"`
	LoadBalancingRules       LoadBalancingRuleResponseArrayOutput       `pulumi:"loadBalancingRules"`
	Location                 pulumi.StringPtrOutput                     `pulumi:"location"`
	Name                     pulumi.StringOutput                        `pulumi:"name"`
	OutboundRules            OutboundRuleResponseArrayOutput            `pulumi:"outboundRules"`
	Probes                   ProbeResponseArrayOutput                   `pulumi:"probes"`
	ProvisioningState        pulumi.StringOutput                        `pulumi:"provisioningState"`
	ResourceGuid             pulumi.StringOutput                        `pulumi:"resourceGuid"`
	Sku                      LoadBalancerSkuResponsePtrOutput           `pulumi:"sku"`
	Tags                     pulumi.StringMapOutput                     `pulumi:"tags"`
	Type                     pulumi.StringOutput                        `pulumi:"type"`
}


func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:LoadBalancer"),
		},
	})
	opts = append(opts, aliases)
	var resource LoadBalancer
	err := ctx.RegisterResource("azure-native:network/v20200401:LoadBalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerState, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	var resource LoadBalancer
	err := ctx.ReadResource("azure-native:network/v20200401:LoadBalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type loadBalancerState struct {
}

type LoadBalancerState struct {
}

func (LoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerState)(nil)).Elem()
}

type loadBalancerArgs struct {
	BackendAddressPools      []BackendAddressPool      `pulumi:"backendAddressPools"`
	FrontendIPConfigurations []FrontendIPConfiguration `pulumi:"frontendIPConfigurations"`
	Id                       *string                   `pulumi:"id"`
	InboundNatPools          []InboundNatPool          `pulumi:"inboundNatPools"`
	InboundNatRules          []InboundNatRuleType      `pulumi:"inboundNatRules"`
	LoadBalancerName         *string                   `pulumi:"loadBalancerName"`
	LoadBalancingRules       []LoadBalancingRule       `pulumi:"loadBalancingRules"`
	Location                 *string                   `pulumi:"location"`
	OutboundRules            []OutboundRule            `pulumi:"outboundRules"`
	Probes                   []Probe                   `pulumi:"probes"`
	ResourceGroupName        string                    `pulumi:"resourceGroupName"`
	Sku                      *LoadBalancerSku          `pulumi:"sku"`
	Tags                     map[string]string         `pulumi:"tags"`
}


type LoadBalancerArgs struct {
	BackendAddressPools      BackendAddressPoolArrayInput
	FrontendIPConfigurations FrontendIPConfigurationArrayInput
	Id                       pulumi.StringPtrInput
	InboundNatPools          InboundNatPoolArrayInput
	InboundNatRules          InboundNatRuleTypeArrayInput
	LoadBalancerName         pulumi.StringPtrInput
	LoadBalancingRules       LoadBalancingRuleArrayInput
	Location                 pulumi.StringPtrInput
	OutboundRules            OutboundRuleArrayInput
	Probes                   ProbeArrayInput
	ResourceGroupName        pulumi.StringInput
	Sku                      LoadBalancerSkuPtrInput
	Tags                     pulumi.StringMapInput
}

func (LoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerArgs)(nil)).Elem()
}

type LoadBalancerInput interface {
	pulumi.Input

	ToLoadBalancerOutput() LoadBalancerOutput
	ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput
}

func (*LoadBalancer) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancer)(nil)).Elem()
}

func (i *LoadBalancer) ToLoadBalancerOutput() LoadBalancerOutput {
	return i.ToLoadBalancerOutputWithContext(context.Background())
}

func (i *LoadBalancer) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerOutput)
}

type LoadBalancerOutput struct{ *pulumi.OutputState }

func (LoadBalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancer)(nil)).Elem()
}

func (o LoadBalancerOutput) ToLoadBalancerOutput() LoadBalancerOutput {
	return o
}

func (o LoadBalancerOutput) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return o
}

func (o LoadBalancerOutput) BackendAddressPools() BackendAddressPoolResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) BackendAddressPoolResponseArrayOutput { return v.BackendAddressPools }).(BackendAddressPoolResponseArrayOutput)
}

func (o LoadBalancerOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o LoadBalancerOutput) FrontendIPConfigurations() FrontendIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) FrontendIPConfigurationResponseArrayOutput { return v.FrontendIPConfigurations }).(FrontendIPConfigurationResponseArrayOutput)
}

func (o LoadBalancerOutput) InboundNatPools() InboundNatPoolResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) InboundNatPoolResponseArrayOutput { return v.InboundNatPools }).(InboundNatPoolResponseArrayOutput)
}

func (o LoadBalancerOutput) InboundNatRules() InboundNatRuleResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) InboundNatRuleResponseArrayOutput { return v.InboundNatRules }).(InboundNatRuleResponseArrayOutput)
}

func (o LoadBalancerOutput) LoadBalancingRules() LoadBalancingRuleResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancingRuleResponseArrayOutput { return v.LoadBalancingRules }).(LoadBalancingRuleResponseArrayOutput)
}

func (o LoadBalancerOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LoadBalancerOutput) OutboundRules() OutboundRuleResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) OutboundRuleResponseArrayOutput { return v.OutboundRules }).(OutboundRuleResponseArrayOutput)
}

func (o LoadBalancerOutput) Probes() ProbeResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) ProbeResponseArrayOutput { return v.Probes }).(ProbeResponseArrayOutput)
}

func (o LoadBalancerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LoadBalancerOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LoadBalancerOutput) Sku() LoadBalancerSkuResponsePtrOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancerSkuResponsePtrOutput { return v.Sku }).(LoadBalancerSkuResponsePtrOutput)
}

func (o LoadBalancerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LoadBalancerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LoadBalancerOutput{})
}
