


package v20190401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerServiceLinuxProfile struct {
	AdminUsername string                           `pulumi:"adminUsername"`
	Ssh           ContainerServiceSshConfiguration `pulumi:"ssh"`
}





type ContainerServiceLinuxProfileInput interface {
	pulumi.Input

	ToContainerServiceLinuxProfileOutput() ContainerServiceLinuxProfileOutput
	ToContainerServiceLinuxProfileOutputWithContext(context.Context) ContainerServiceLinuxProfileOutput
}

type ContainerServiceLinuxProfileArgs struct {
	AdminUsername pulumi.StringInput                    `pulumi:"adminUsername"`
	Ssh           ContainerServiceSshConfigurationInput `pulumi:"ssh"`
}

func (ContainerServiceLinuxProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceLinuxProfile)(nil)).Elem()
}

func (i ContainerServiceLinuxProfileArgs) ToContainerServiceLinuxProfileOutput() ContainerServiceLinuxProfileOutput {
	return i.ToContainerServiceLinuxProfileOutputWithContext(context.Background())
}

func (i ContainerServiceLinuxProfileArgs) ToContainerServiceLinuxProfileOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceLinuxProfileOutput)
}

func (i ContainerServiceLinuxProfileArgs) ToContainerServiceLinuxProfilePtrOutput() ContainerServiceLinuxProfilePtrOutput {
	return i.ToContainerServiceLinuxProfilePtrOutputWithContext(context.Background())
}

func (i ContainerServiceLinuxProfileArgs) ToContainerServiceLinuxProfilePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceLinuxProfileOutput).ToContainerServiceLinuxProfilePtrOutputWithContext(ctx)
}









type ContainerServiceLinuxProfilePtrInput interface {
	pulumi.Input

	ToContainerServiceLinuxProfilePtrOutput() ContainerServiceLinuxProfilePtrOutput
	ToContainerServiceLinuxProfilePtrOutputWithContext(context.Context) ContainerServiceLinuxProfilePtrOutput
}

type containerServiceLinuxProfilePtrType ContainerServiceLinuxProfileArgs

func ContainerServiceLinuxProfilePtr(v *ContainerServiceLinuxProfileArgs) ContainerServiceLinuxProfilePtrInput {
	return (*containerServiceLinuxProfilePtrType)(v)
}

func (*containerServiceLinuxProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceLinuxProfile)(nil)).Elem()
}

func (i *containerServiceLinuxProfilePtrType) ToContainerServiceLinuxProfilePtrOutput() ContainerServiceLinuxProfilePtrOutput {
	return i.ToContainerServiceLinuxProfilePtrOutputWithContext(context.Background())
}

func (i *containerServiceLinuxProfilePtrType) ToContainerServiceLinuxProfilePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceLinuxProfilePtrOutput)
}

type ContainerServiceLinuxProfileOutput struct{ *pulumi.OutputState }

func (ContainerServiceLinuxProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceLinuxProfile)(nil)).Elem()
}

func (o ContainerServiceLinuxProfileOutput) ToContainerServiceLinuxProfileOutput() ContainerServiceLinuxProfileOutput {
	return o
}

func (o ContainerServiceLinuxProfileOutput) ToContainerServiceLinuxProfileOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileOutput {
	return o
}

func (o ContainerServiceLinuxProfileOutput) ToContainerServiceLinuxProfilePtrOutput() ContainerServiceLinuxProfilePtrOutput {
	return o.ToContainerServiceLinuxProfilePtrOutputWithContext(context.Background())
}

func (o ContainerServiceLinuxProfileOutput) ToContainerServiceLinuxProfilePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceLinuxProfile) *ContainerServiceLinuxProfile {
		return &v
	}).(ContainerServiceLinuxProfilePtrOutput)
}

func (o ContainerServiceLinuxProfileOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceLinuxProfile) string { return v.AdminUsername }).(pulumi.StringOutput)
}

func (o ContainerServiceLinuxProfileOutput) Ssh() ContainerServiceSshConfigurationOutput {
	return o.ApplyT(func(v ContainerServiceLinuxProfile) ContainerServiceSshConfiguration { return v.Ssh }).(ContainerServiceSshConfigurationOutput)
}

type ContainerServiceLinuxProfilePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceLinuxProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceLinuxProfile)(nil)).Elem()
}

func (o ContainerServiceLinuxProfilePtrOutput) ToContainerServiceLinuxProfilePtrOutput() ContainerServiceLinuxProfilePtrOutput {
	return o
}

func (o ContainerServiceLinuxProfilePtrOutput) ToContainerServiceLinuxProfilePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfilePtrOutput {
	return o
}

func (o ContainerServiceLinuxProfilePtrOutput) Elem() ContainerServiceLinuxProfileOutput {
	return o.ApplyT(func(v *ContainerServiceLinuxProfile) ContainerServiceLinuxProfile {
		if v != nil {
			return *v
		}
		var ret ContainerServiceLinuxProfile
		return ret
	}).(ContainerServiceLinuxProfileOutput)
}

func (o ContainerServiceLinuxProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceLinuxProfile) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceLinuxProfilePtrOutput) Ssh() ContainerServiceSshConfigurationPtrOutput {
	return o.ApplyT(func(v *ContainerServiceLinuxProfile) *ContainerServiceSshConfiguration {
		if v == nil {
			return nil
		}
		return &v.Ssh
	}).(ContainerServiceSshConfigurationPtrOutput)
}

type ContainerServiceLinuxProfileResponse struct {
	AdminUsername string                                   `pulumi:"adminUsername"`
	Ssh           ContainerServiceSshConfigurationResponse `pulumi:"ssh"`
}

type ContainerServiceLinuxProfileResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceLinuxProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceLinuxProfileResponse)(nil)).Elem()
}

func (o ContainerServiceLinuxProfileResponseOutput) ToContainerServiceLinuxProfileResponseOutput() ContainerServiceLinuxProfileResponseOutput {
	return o
}

func (o ContainerServiceLinuxProfileResponseOutput) ToContainerServiceLinuxProfileResponseOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileResponseOutput {
	return o
}

func (o ContainerServiceLinuxProfileResponseOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceLinuxProfileResponse) string { return v.AdminUsername }).(pulumi.StringOutput)
}

func (o ContainerServiceLinuxProfileResponseOutput) Ssh() ContainerServiceSshConfigurationResponseOutput {
	return o.ApplyT(func(v ContainerServiceLinuxProfileResponse) ContainerServiceSshConfigurationResponse { return v.Ssh }).(ContainerServiceSshConfigurationResponseOutput)
}

type ContainerServiceLinuxProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceLinuxProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceLinuxProfileResponse)(nil)).Elem()
}

func (o ContainerServiceLinuxProfileResponsePtrOutput) ToContainerServiceLinuxProfileResponsePtrOutput() ContainerServiceLinuxProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceLinuxProfileResponsePtrOutput) ToContainerServiceLinuxProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceLinuxProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceLinuxProfileResponsePtrOutput) Elem() ContainerServiceLinuxProfileResponseOutput {
	return o.ApplyT(func(v *ContainerServiceLinuxProfileResponse) ContainerServiceLinuxProfileResponse {
		if v != nil {
			return *v
		}
		var ret ContainerServiceLinuxProfileResponse
		return ret
	}).(ContainerServiceLinuxProfileResponseOutput)
}

func (o ContainerServiceLinuxProfileResponsePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceLinuxProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceLinuxProfileResponsePtrOutput) Ssh() ContainerServiceSshConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ContainerServiceLinuxProfileResponse) *ContainerServiceSshConfigurationResponse {
		if v == nil {
			return nil
		}
		return &v.Ssh
	}).(ContainerServiceSshConfigurationResponsePtrOutput)
}

type ContainerServiceNetworkProfile struct {
	DnsServiceIP     *string `pulumi:"dnsServiceIP"`
	DockerBridgeCidr *string `pulumi:"dockerBridgeCidr"`
	LoadBalancerSku  *string `pulumi:"loadBalancerSku"`
	NetworkPlugin    *string `pulumi:"networkPlugin"`
	NetworkPolicy    *string `pulumi:"networkPolicy"`
	PodCidr          *string `pulumi:"podCidr"`
	ServiceCidr      *string `pulumi:"serviceCidr"`
}


func (val *ContainerServiceNetworkProfile) Defaults() *ContainerServiceNetworkProfile {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DnsServiceIP) {
		dnsServiceIP_ := "10.0.0.10"
		tmp.DnsServiceIP = &dnsServiceIP_
	}
	if isZero(tmp.DockerBridgeCidr) {
		dockerBridgeCidr_ := "172.17.0.1/16"
		tmp.DockerBridgeCidr = &dockerBridgeCidr_
	}
	if isZero(tmp.NetworkPlugin) {
		networkPlugin_ := "kubenet"
		tmp.NetworkPlugin = &networkPlugin_
	}
	if isZero(tmp.PodCidr) {
		podCidr_ := "10.244.0.0/16"
		tmp.PodCidr = &podCidr_
	}
	if isZero(tmp.ServiceCidr) {
		serviceCidr_ := "10.0.0.0/16"
		tmp.ServiceCidr = &serviceCidr_
	}
	return &tmp
}





type ContainerServiceNetworkProfileInput interface {
	pulumi.Input

	ToContainerServiceNetworkProfileOutput() ContainerServiceNetworkProfileOutput
	ToContainerServiceNetworkProfileOutputWithContext(context.Context) ContainerServiceNetworkProfileOutput
}

type ContainerServiceNetworkProfileArgs struct {
	DnsServiceIP     pulumi.StringPtrInput `pulumi:"dnsServiceIP"`
	DockerBridgeCidr pulumi.StringPtrInput `pulumi:"dockerBridgeCidr"`
	LoadBalancerSku  pulumi.StringPtrInput `pulumi:"loadBalancerSku"`
	NetworkPlugin    pulumi.StringPtrInput `pulumi:"networkPlugin"`
	NetworkPolicy    pulumi.StringPtrInput `pulumi:"networkPolicy"`
	PodCidr          pulumi.StringPtrInput `pulumi:"podCidr"`
	ServiceCidr      pulumi.StringPtrInput `pulumi:"serviceCidr"`
}


func (val *ContainerServiceNetworkProfileArgs) Defaults() *ContainerServiceNetworkProfileArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DnsServiceIP) {
		tmp.DnsServiceIP = pulumi.StringPtr("10.0.0.10")
	}
	if isZero(tmp.DockerBridgeCidr) {
		tmp.DockerBridgeCidr = pulumi.StringPtr("172.17.0.1/16")
	}
	if isZero(tmp.NetworkPlugin) {
		tmp.NetworkPlugin = pulumi.StringPtr("kubenet")
	}
	if isZero(tmp.PodCidr) {
		tmp.PodCidr = pulumi.StringPtr("10.244.0.0/16")
	}
	if isZero(tmp.ServiceCidr) {
		tmp.ServiceCidr = pulumi.StringPtr("10.0.0.0/16")
	}
	return &tmp
}
func (ContainerServiceNetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceNetworkProfile)(nil)).Elem()
}

func (i ContainerServiceNetworkProfileArgs) ToContainerServiceNetworkProfileOutput() ContainerServiceNetworkProfileOutput {
	return i.ToContainerServiceNetworkProfileOutputWithContext(context.Background())
}

func (i ContainerServiceNetworkProfileArgs) ToContainerServiceNetworkProfileOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceNetworkProfileOutput)
}

func (i ContainerServiceNetworkProfileArgs) ToContainerServiceNetworkProfilePtrOutput() ContainerServiceNetworkProfilePtrOutput {
	return i.ToContainerServiceNetworkProfilePtrOutputWithContext(context.Background())
}

func (i ContainerServiceNetworkProfileArgs) ToContainerServiceNetworkProfilePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceNetworkProfileOutput).ToContainerServiceNetworkProfilePtrOutputWithContext(ctx)
}









type ContainerServiceNetworkProfilePtrInput interface {
	pulumi.Input

	ToContainerServiceNetworkProfilePtrOutput() ContainerServiceNetworkProfilePtrOutput
	ToContainerServiceNetworkProfilePtrOutputWithContext(context.Context) ContainerServiceNetworkProfilePtrOutput
}

type containerServiceNetworkProfilePtrType ContainerServiceNetworkProfileArgs

func ContainerServiceNetworkProfilePtr(v *ContainerServiceNetworkProfileArgs) ContainerServiceNetworkProfilePtrInput {
	return (*containerServiceNetworkProfilePtrType)(v)
}

func (*containerServiceNetworkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceNetworkProfile)(nil)).Elem()
}

func (i *containerServiceNetworkProfilePtrType) ToContainerServiceNetworkProfilePtrOutput() ContainerServiceNetworkProfilePtrOutput {
	return i.ToContainerServiceNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *containerServiceNetworkProfilePtrType) ToContainerServiceNetworkProfilePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceNetworkProfilePtrOutput)
}

type ContainerServiceNetworkProfileOutput struct{ *pulumi.OutputState }

func (ContainerServiceNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceNetworkProfile)(nil)).Elem()
}

func (o ContainerServiceNetworkProfileOutput) ToContainerServiceNetworkProfileOutput() ContainerServiceNetworkProfileOutput {
	return o
}

func (o ContainerServiceNetworkProfileOutput) ToContainerServiceNetworkProfileOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileOutput {
	return o
}

func (o ContainerServiceNetworkProfileOutput) ToContainerServiceNetworkProfilePtrOutput() ContainerServiceNetworkProfilePtrOutput {
	return o.ToContainerServiceNetworkProfilePtrOutputWithContext(context.Background())
}

func (o ContainerServiceNetworkProfileOutput) ToContainerServiceNetworkProfilePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceNetworkProfile) *ContainerServiceNetworkProfile {
		return &v
	}).(ContainerServiceNetworkProfilePtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.DnsServiceIP }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.DockerBridgeCidr }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.LoadBalancerSku }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) NetworkPlugin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.NetworkPlugin }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) NetworkPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.NetworkPolicy }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.PodCidr }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfile) *string { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

type ContainerServiceNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceNetworkProfile)(nil)).Elem()
}

func (o ContainerServiceNetworkProfilePtrOutput) ToContainerServiceNetworkProfilePtrOutput() ContainerServiceNetworkProfilePtrOutput {
	return o
}

func (o ContainerServiceNetworkProfilePtrOutput) ToContainerServiceNetworkProfilePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfilePtrOutput {
	return o
}

func (o ContainerServiceNetworkProfilePtrOutput) Elem() ContainerServiceNetworkProfileOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) ContainerServiceNetworkProfile {
		if v != nil {
			return *v
		}
		var ret ContainerServiceNetworkProfile
		return ret
	}).(ContainerServiceNetworkProfileOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.DnsServiceIP
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.DockerBridgeCidr
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancerSku
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) NetworkPlugin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.NetworkPlugin
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) NetworkPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.NetworkPolicy
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.PodCidr
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfilePtrOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCidr
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceNetworkProfileResponse struct {
	DnsServiceIP     *string `pulumi:"dnsServiceIP"`
	DockerBridgeCidr *string `pulumi:"dockerBridgeCidr"`
	LoadBalancerSku  *string `pulumi:"loadBalancerSku"`
	NetworkPlugin    *string `pulumi:"networkPlugin"`
	NetworkPolicy    *string `pulumi:"networkPolicy"`
	PodCidr          *string `pulumi:"podCidr"`
	ServiceCidr      *string `pulumi:"serviceCidr"`
}


func (val *ContainerServiceNetworkProfileResponse) Defaults() *ContainerServiceNetworkProfileResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DnsServiceIP) {
		dnsServiceIP_ := "10.0.0.10"
		tmp.DnsServiceIP = &dnsServiceIP_
	}
	if isZero(tmp.DockerBridgeCidr) {
		dockerBridgeCidr_ := "172.17.0.1/16"
		tmp.DockerBridgeCidr = &dockerBridgeCidr_
	}
	if isZero(tmp.NetworkPlugin) {
		networkPlugin_ := "kubenet"
		tmp.NetworkPlugin = &networkPlugin_
	}
	if isZero(tmp.PodCidr) {
		podCidr_ := "10.244.0.0/16"
		tmp.PodCidr = &podCidr_
	}
	if isZero(tmp.ServiceCidr) {
		serviceCidr_ := "10.0.0.0/16"
		tmp.ServiceCidr = &serviceCidr_
	}
	return &tmp
}

type ContainerServiceNetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceNetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceNetworkProfileResponse)(nil)).Elem()
}

func (o ContainerServiceNetworkProfileResponseOutput) ToContainerServiceNetworkProfileResponseOutput() ContainerServiceNetworkProfileResponseOutput {
	return o
}

func (o ContainerServiceNetworkProfileResponseOutput) ToContainerServiceNetworkProfileResponseOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileResponseOutput {
	return o
}

func (o ContainerServiceNetworkProfileResponseOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.DnsServiceIP }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.DockerBridgeCidr }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.LoadBalancerSku }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) NetworkPlugin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.NetworkPlugin }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) NetworkPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.NetworkPolicy }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.PodCidr }).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponseOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerServiceNetworkProfileResponse) *string { return v.ServiceCidr }).(pulumi.StringPtrOutput)
}

type ContainerServiceNetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceNetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceNetworkProfileResponse)(nil)).Elem()
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) ToContainerServiceNetworkProfileResponsePtrOutput() ContainerServiceNetworkProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) ToContainerServiceNetworkProfileResponsePtrOutputWithContext(ctx context.Context) ContainerServiceNetworkProfileResponsePtrOutput {
	return o
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) Elem() ContainerServiceNetworkProfileResponseOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) ContainerServiceNetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret ContainerServiceNetworkProfileResponse
		return ret
	}).(ContainerServiceNetworkProfileResponseOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) DnsServiceIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.DnsServiceIP
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.DockerBridgeCidr
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) LoadBalancerSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancerSku
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) NetworkPlugin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.NetworkPlugin
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) NetworkPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.NetworkPolicy
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) PodCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.PodCidr
	}).(pulumi.StringPtrOutput)
}

func (o ContainerServiceNetworkProfileResponsePtrOutput) ServiceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerServiceNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceCidr
	}).(pulumi.StringPtrOutput)
}

type ContainerServiceSshConfiguration struct {
	PublicKeys []ContainerServiceSshPublicKey `pulumi:"publicKeys"`
}





type ContainerServiceSshConfigurationInput interface {
	pulumi.Input

	ToContainerServiceSshConfigurationOutput() ContainerServiceSshConfigurationOutput
	ToContainerServiceSshConfigurationOutputWithContext(context.Context) ContainerServiceSshConfigurationOutput
}

type ContainerServiceSshConfigurationArgs struct {
	PublicKeys ContainerServiceSshPublicKeyArrayInput `pulumi:"publicKeys"`
}

func (ContainerServiceSshConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshConfiguration)(nil)).Elem()
}

func (i ContainerServiceSshConfigurationArgs) ToContainerServiceSshConfigurationOutput() ContainerServiceSshConfigurationOutput {
	return i.ToContainerServiceSshConfigurationOutputWithContext(context.Background())
}

func (i ContainerServiceSshConfigurationArgs) ToContainerServiceSshConfigurationOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshConfigurationOutput)
}

func (i ContainerServiceSshConfigurationArgs) ToContainerServiceSshConfigurationPtrOutput() ContainerServiceSshConfigurationPtrOutput {
	return i.ToContainerServiceSshConfigurationPtrOutputWithContext(context.Background())
}

func (i ContainerServiceSshConfigurationArgs) ToContainerServiceSshConfigurationPtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshConfigurationOutput).ToContainerServiceSshConfigurationPtrOutputWithContext(ctx)
}









type ContainerServiceSshConfigurationPtrInput interface {
	pulumi.Input

	ToContainerServiceSshConfigurationPtrOutput() ContainerServiceSshConfigurationPtrOutput
	ToContainerServiceSshConfigurationPtrOutputWithContext(context.Context) ContainerServiceSshConfigurationPtrOutput
}

type containerServiceSshConfigurationPtrType ContainerServiceSshConfigurationArgs

func ContainerServiceSshConfigurationPtr(v *ContainerServiceSshConfigurationArgs) ContainerServiceSshConfigurationPtrInput {
	return (*containerServiceSshConfigurationPtrType)(v)
}

func (*containerServiceSshConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceSshConfiguration)(nil)).Elem()
}

func (i *containerServiceSshConfigurationPtrType) ToContainerServiceSshConfigurationPtrOutput() ContainerServiceSshConfigurationPtrOutput {
	return i.ToContainerServiceSshConfigurationPtrOutputWithContext(context.Background())
}

func (i *containerServiceSshConfigurationPtrType) ToContainerServiceSshConfigurationPtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshConfigurationPtrOutput)
}

type ContainerServiceSshConfigurationOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshConfiguration)(nil)).Elem()
}

func (o ContainerServiceSshConfigurationOutput) ToContainerServiceSshConfigurationOutput() ContainerServiceSshConfigurationOutput {
	return o
}

func (o ContainerServiceSshConfigurationOutput) ToContainerServiceSshConfigurationOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationOutput {
	return o
}

func (o ContainerServiceSshConfigurationOutput) ToContainerServiceSshConfigurationPtrOutput() ContainerServiceSshConfigurationPtrOutput {
	return o.ToContainerServiceSshConfigurationPtrOutputWithContext(context.Background())
}

func (o ContainerServiceSshConfigurationOutput) ToContainerServiceSshConfigurationPtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerServiceSshConfiguration) *ContainerServiceSshConfiguration {
		return &v
	}).(ContainerServiceSshConfigurationPtrOutput)
}

func (o ContainerServiceSshConfigurationOutput) PublicKeys() ContainerServiceSshPublicKeyArrayOutput {
	return o.ApplyT(func(v ContainerServiceSshConfiguration) []ContainerServiceSshPublicKey { return v.PublicKeys }).(ContainerServiceSshPublicKeyArrayOutput)
}

type ContainerServiceSshConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceSshConfiguration)(nil)).Elem()
}

func (o ContainerServiceSshConfigurationPtrOutput) ToContainerServiceSshConfigurationPtrOutput() ContainerServiceSshConfigurationPtrOutput {
	return o
}

func (o ContainerServiceSshConfigurationPtrOutput) ToContainerServiceSshConfigurationPtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationPtrOutput {
	return o
}

func (o ContainerServiceSshConfigurationPtrOutput) Elem() ContainerServiceSshConfigurationOutput {
	return o.ApplyT(func(v *ContainerServiceSshConfiguration) ContainerServiceSshConfiguration {
		if v != nil {
			return *v
		}
		var ret ContainerServiceSshConfiguration
		return ret
	}).(ContainerServiceSshConfigurationOutput)
}

func (o ContainerServiceSshConfigurationPtrOutput) PublicKeys() ContainerServiceSshPublicKeyArrayOutput {
	return o.ApplyT(func(v *ContainerServiceSshConfiguration) []ContainerServiceSshPublicKey {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(ContainerServiceSshPublicKeyArrayOutput)
}

type ContainerServiceSshConfigurationResponse struct {
	PublicKeys []ContainerServiceSshPublicKeyResponse `pulumi:"publicKeys"`
}

type ContainerServiceSshConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshConfigurationResponse)(nil)).Elem()
}

func (o ContainerServiceSshConfigurationResponseOutput) ToContainerServiceSshConfigurationResponseOutput() ContainerServiceSshConfigurationResponseOutput {
	return o
}

func (o ContainerServiceSshConfigurationResponseOutput) ToContainerServiceSshConfigurationResponseOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationResponseOutput {
	return o
}

func (o ContainerServiceSshConfigurationResponseOutput) PublicKeys() ContainerServiceSshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v ContainerServiceSshConfigurationResponse) []ContainerServiceSshPublicKeyResponse {
		return v.PublicKeys
	}).(ContainerServiceSshPublicKeyResponseArrayOutput)
}

type ContainerServiceSshConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerServiceSshConfigurationResponse)(nil)).Elem()
}

func (o ContainerServiceSshConfigurationResponsePtrOutput) ToContainerServiceSshConfigurationResponsePtrOutput() ContainerServiceSshConfigurationResponsePtrOutput {
	return o
}

func (o ContainerServiceSshConfigurationResponsePtrOutput) ToContainerServiceSshConfigurationResponsePtrOutputWithContext(ctx context.Context) ContainerServiceSshConfigurationResponsePtrOutput {
	return o
}

func (o ContainerServiceSshConfigurationResponsePtrOutput) Elem() ContainerServiceSshConfigurationResponseOutput {
	return o.ApplyT(func(v *ContainerServiceSshConfigurationResponse) ContainerServiceSshConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ContainerServiceSshConfigurationResponse
		return ret
	}).(ContainerServiceSshConfigurationResponseOutput)
}

func (o ContainerServiceSshConfigurationResponsePtrOutput) PublicKeys() ContainerServiceSshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v *ContainerServiceSshConfigurationResponse) []ContainerServiceSshPublicKeyResponse {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(ContainerServiceSshPublicKeyResponseArrayOutput)
}

type ContainerServiceSshPublicKey struct {
	KeyData string `pulumi:"keyData"`
}





type ContainerServiceSshPublicKeyInput interface {
	pulumi.Input

	ToContainerServiceSshPublicKeyOutput() ContainerServiceSshPublicKeyOutput
	ToContainerServiceSshPublicKeyOutputWithContext(context.Context) ContainerServiceSshPublicKeyOutput
}

type ContainerServiceSshPublicKeyArgs struct {
	KeyData pulumi.StringInput `pulumi:"keyData"`
}

func (ContainerServiceSshPublicKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshPublicKey)(nil)).Elem()
}

func (i ContainerServiceSshPublicKeyArgs) ToContainerServiceSshPublicKeyOutput() ContainerServiceSshPublicKeyOutput {
	return i.ToContainerServiceSshPublicKeyOutputWithContext(context.Background())
}

func (i ContainerServiceSshPublicKeyArgs) ToContainerServiceSshPublicKeyOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshPublicKeyOutput)
}





type ContainerServiceSshPublicKeyArrayInput interface {
	pulumi.Input

	ToContainerServiceSshPublicKeyArrayOutput() ContainerServiceSshPublicKeyArrayOutput
	ToContainerServiceSshPublicKeyArrayOutputWithContext(context.Context) ContainerServiceSshPublicKeyArrayOutput
}

type ContainerServiceSshPublicKeyArray []ContainerServiceSshPublicKeyInput

func (ContainerServiceSshPublicKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerServiceSshPublicKey)(nil)).Elem()
}

func (i ContainerServiceSshPublicKeyArray) ToContainerServiceSshPublicKeyArrayOutput() ContainerServiceSshPublicKeyArrayOutput {
	return i.ToContainerServiceSshPublicKeyArrayOutputWithContext(context.Background())
}

func (i ContainerServiceSshPublicKeyArray) ToContainerServiceSshPublicKeyArrayOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerServiceSshPublicKeyArrayOutput)
}

type ContainerServiceSshPublicKeyOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshPublicKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshPublicKey)(nil)).Elem()
}

func (o ContainerServiceSshPublicKeyOutput) ToContainerServiceSshPublicKeyOutput() ContainerServiceSshPublicKeyOutput {
	return o
}

func (o ContainerServiceSshPublicKeyOutput) ToContainerServiceSshPublicKeyOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyOutput {
	return o
}

func (o ContainerServiceSshPublicKeyOutput) KeyData() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceSshPublicKey) string { return v.KeyData }).(pulumi.StringOutput)
}

type ContainerServiceSshPublicKeyArrayOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshPublicKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerServiceSshPublicKey)(nil)).Elem()
}

func (o ContainerServiceSshPublicKeyArrayOutput) ToContainerServiceSshPublicKeyArrayOutput() ContainerServiceSshPublicKeyArrayOutput {
	return o
}

func (o ContainerServiceSshPublicKeyArrayOutput) ToContainerServiceSshPublicKeyArrayOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyArrayOutput {
	return o
}

func (o ContainerServiceSshPublicKeyArrayOutput) Index(i pulumi.IntInput) ContainerServiceSshPublicKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerServiceSshPublicKey {
		return vs[0].([]ContainerServiceSshPublicKey)[vs[1].(int)]
	}).(ContainerServiceSshPublicKeyOutput)
}

type ContainerServiceSshPublicKeyResponse struct {
	KeyData string `pulumi:"keyData"`
}

type ContainerServiceSshPublicKeyResponseOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshPublicKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerServiceSshPublicKeyResponse)(nil)).Elem()
}

func (o ContainerServiceSshPublicKeyResponseOutput) ToContainerServiceSshPublicKeyResponseOutput() ContainerServiceSshPublicKeyResponseOutput {
	return o
}

func (o ContainerServiceSshPublicKeyResponseOutput) ToContainerServiceSshPublicKeyResponseOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyResponseOutput {
	return o
}

func (o ContainerServiceSshPublicKeyResponseOutput) KeyData() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerServiceSshPublicKeyResponse) string { return v.KeyData }).(pulumi.StringOutput)
}

type ContainerServiceSshPublicKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerServiceSshPublicKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerServiceSshPublicKeyResponse)(nil)).Elem()
}

func (o ContainerServiceSshPublicKeyResponseArrayOutput) ToContainerServiceSshPublicKeyResponseArrayOutput() ContainerServiceSshPublicKeyResponseArrayOutput {
	return o
}

func (o ContainerServiceSshPublicKeyResponseArrayOutput) ToContainerServiceSshPublicKeyResponseArrayOutputWithContext(ctx context.Context) ContainerServiceSshPublicKeyResponseArrayOutput {
	return o
}

func (o ContainerServiceSshPublicKeyResponseArrayOutput) Index(i pulumi.IntInput) ContainerServiceSshPublicKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerServiceSshPublicKeyResponse {
		return vs[0].([]ContainerServiceSshPublicKeyResponse)[vs[1].(int)]
	}).(ContainerServiceSshPublicKeyResponseOutput)
}

type CredentialResultResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type CredentialResultResponseOutput struct{ *pulumi.OutputState }

func (CredentialResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CredentialResultResponse)(nil)).Elem()
}

func (o CredentialResultResponseOutput) ToCredentialResultResponseOutput() CredentialResultResponseOutput {
	return o
}

func (o CredentialResultResponseOutput) ToCredentialResultResponseOutputWithContext(ctx context.Context) CredentialResultResponseOutput {
	return o
}

func (o CredentialResultResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CredentialResultResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CredentialResultResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v CredentialResultResponse) string { return v.Value }).(pulumi.StringOutput)
}

type CredentialResultResponseArrayOutput struct{ *pulumi.OutputState }

func (CredentialResultResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CredentialResultResponse)(nil)).Elem()
}

func (o CredentialResultResponseArrayOutput) ToCredentialResultResponseArrayOutput() CredentialResultResponseArrayOutput {
	return o
}

func (o CredentialResultResponseArrayOutput) ToCredentialResultResponseArrayOutputWithContext(ctx context.Context) CredentialResultResponseArrayOutput {
	return o
}

func (o CredentialResultResponseArrayOutput) Index(i pulumi.IntInput) CredentialResultResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CredentialResultResponse {
		return vs[0].([]CredentialResultResponse)[vs[1].(int)]
	}).(CredentialResultResponseOutput)
}

type ManagedClusterAADProfile struct {
	ClientAppID     string  `pulumi:"clientAppID"`
	ServerAppID     string  `pulumi:"serverAppID"`
	ServerAppSecret *string `pulumi:"serverAppSecret"`
	TenantID        *string `pulumi:"tenantID"`
}





type ManagedClusterAADProfileInput interface {
	pulumi.Input

	ToManagedClusterAADProfileOutput() ManagedClusterAADProfileOutput
	ToManagedClusterAADProfileOutputWithContext(context.Context) ManagedClusterAADProfileOutput
}

type ManagedClusterAADProfileArgs struct {
	ClientAppID     pulumi.StringInput    `pulumi:"clientAppID"`
	ServerAppID     pulumi.StringInput    `pulumi:"serverAppID"`
	ServerAppSecret pulumi.StringPtrInput `pulumi:"serverAppSecret"`
	TenantID        pulumi.StringPtrInput `pulumi:"tenantID"`
}

func (ManagedClusterAADProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAADProfile)(nil)).Elem()
}

func (i ManagedClusterAADProfileArgs) ToManagedClusterAADProfileOutput() ManagedClusterAADProfileOutput {
	return i.ToManagedClusterAADProfileOutputWithContext(context.Background())
}

func (i ManagedClusterAADProfileArgs) ToManagedClusterAADProfileOutputWithContext(ctx context.Context) ManagedClusterAADProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAADProfileOutput)
}

func (i ManagedClusterAADProfileArgs) ToManagedClusterAADProfilePtrOutput() ManagedClusterAADProfilePtrOutput {
	return i.ToManagedClusterAADProfilePtrOutputWithContext(context.Background())
}

func (i ManagedClusterAADProfileArgs) ToManagedClusterAADProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAADProfileOutput).ToManagedClusterAADProfilePtrOutputWithContext(ctx)
}









type ManagedClusterAADProfilePtrInput interface {
	pulumi.Input

	ToManagedClusterAADProfilePtrOutput() ManagedClusterAADProfilePtrOutput
	ToManagedClusterAADProfilePtrOutputWithContext(context.Context) ManagedClusterAADProfilePtrOutput
}

type managedClusterAADProfilePtrType ManagedClusterAADProfileArgs

func ManagedClusterAADProfilePtr(v *ManagedClusterAADProfileArgs) ManagedClusterAADProfilePtrInput {
	return (*managedClusterAADProfilePtrType)(v)
}

func (*managedClusterAADProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAADProfile)(nil)).Elem()
}

func (i *managedClusterAADProfilePtrType) ToManagedClusterAADProfilePtrOutput() ManagedClusterAADProfilePtrOutput {
	return i.ToManagedClusterAADProfilePtrOutputWithContext(context.Background())
}

func (i *managedClusterAADProfilePtrType) ToManagedClusterAADProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAADProfilePtrOutput)
}

type ManagedClusterAADProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterAADProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAADProfile)(nil)).Elem()
}

func (o ManagedClusterAADProfileOutput) ToManagedClusterAADProfileOutput() ManagedClusterAADProfileOutput {
	return o
}

func (o ManagedClusterAADProfileOutput) ToManagedClusterAADProfileOutputWithContext(ctx context.Context) ManagedClusterAADProfileOutput {
	return o
}

func (o ManagedClusterAADProfileOutput) ToManagedClusterAADProfilePtrOutput() ManagedClusterAADProfilePtrOutput {
	return o.ToManagedClusterAADProfilePtrOutputWithContext(context.Background())
}

func (o ManagedClusterAADProfileOutput) ToManagedClusterAADProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterAADProfile) *ManagedClusterAADProfile {
		return &v
	}).(ManagedClusterAADProfilePtrOutput)
}

func (o ManagedClusterAADProfileOutput) ClientAppID() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAADProfile) string { return v.ClientAppID }).(pulumi.StringOutput)
}

func (o ManagedClusterAADProfileOutput) ServerAppID() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAADProfile) string { return v.ServerAppID }).(pulumi.StringOutput)
}

func (o ManagedClusterAADProfileOutput) ServerAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfile) *string { return v.ServerAppSecret }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileOutput) TenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfile) *string { return v.TenantID }).(pulumi.StringPtrOutput)
}

type ManagedClusterAADProfilePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterAADProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAADProfile)(nil)).Elem()
}

func (o ManagedClusterAADProfilePtrOutput) ToManagedClusterAADProfilePtrOutput() ManagedClusterAADProfilePtrOutput {
	return o
}

func (o ManagedClusterAADProfilePtrOutput) ToManagedClusterAADProfilePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfilePtrOutput {
	return o
}

func (o ManagedClusterAADProfilePtrOutput) Elem() ManagedClusterAADProfileOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfile) ManagedClusterAADProfile {
		if v != nil {
			return *v
		}
		var ret ManagedClusterAADProfile
		return ret
	}).(ManagedClusterAADProfileOutput)
}

func (o ManagedClusterAADProfilePtrOutput) ClientAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfile) *string {
		if v == nil {
			return nil
		}
		return &v.ClientAppID
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfilePtrOutput) ServerAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfile) *string {
		if v == nil {
			return nil
		}
		return &v.ServerAppID
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfilePtrOutput) ServerAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfile) *string {
		if v == nil {
			return nil
		}
		return v.ServerAppSecret
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfilePtrOutput) TenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfile) *string {
		if v == nil {
			return nil
		}
		return v.TenantID
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterAADProfileResponse struct {
	ClientAppID     string  `pulumi:"clientAppID"`
	ServerAppID     string  `pulumi:"serverAppID"`
	ServerAppSecret *string `pulumi:"serverAppSecret"`
	TenantID        *string `pulumi:"tenantID"`
}

type ManagedClusterAADProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterAADProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAADProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAADProfileResponseOutput) ToManagedClusterAADProfileResponseOutput() ManagedClusterAADProfileResponseOutput {
	return o
}

func (o ManagedClusterAADProfileResponseOutput) ToManagedClusterAADProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAADProfileResponseOutput {
	return o
}

func (o ManagedClusterAADProfileResponseOutput) ClientAppID() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAADProfileResponse) string { return v.ClientAppID }).(pulumi.StringOutput)
}

func (o ManagedClusterAADProfileResponseOutput) ServerAppID() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAADProfileResponse) string { return v.ServerAppID }).(pulumi.StringOutput)
}

func (o ManagedClusterAADProfileResponseOutput) ServerAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfileResponse) *string { return v.ServerAppSecret }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileResponseOutput) TenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAADProfileResponse) *string { return v.TenantID }).(pulumi.StringPtrOutput)
}

type ManagedClusterAADProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterAADProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterAADProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAADProfileResponsePtrOutput) ToManagedClusterAADProfileResponsePtrOutput() ManagedClusterAADProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterAADProfileResponsePtrOutput) ToManagedClusterAADProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterAADProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterAADProfileResponsePtrOutput) Elem() ManagedClusterAADProfileResponseOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfileResponse) ManagedClusterAADProfileResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterAADProfileResponse
		return ret
	}).(ManagedClusterAADProfileResponseOutput)
}

func (o ManagedClusterAADProfileResponsePtrOutput) ClientAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClientAppID
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileResponsePtrOutput) ServerAppID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ServerAppID
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileResponsePtrOutput) ServerAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServerAppSecret
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAADProfileResponsePtrOutput) TenantID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterAADProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantID
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterAddonProfile struct {
	Config  map[string]string `pulumi:"config"`
	Enabled bool              `pulumi:"enabled"`
}





type ManagedClusterAddonProfileInput interface {
	pulumi.Input

	ToManagedClusterAddonProfileOutput() ManagedClusterAddonProfileOutput
	ToManagedClusterAddonProfileOutputWithContext(context.Context) ManagedClusterAddonProfileOutput
}

type ManagedClusterAddonProfileArgs struct {
	Config  pulumi.StringMapInput `pulumi:"config"`
	Enabled pulumi.BoolInput      `pulumi:"enabled"`
}

func (ManagedClusterAddonProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAddonProfile)(nil)).Elem()
}

func (i ManagedClusterAddonProfileArgs) ToManagedClusterAddonProfileOutput() ManagedClusterAddonProfileOutput {
	return i.ToManagedClusterAddonProfileOutputWithContext(context.Background())
}

func (i ManagedClusterAddonProfileArgs) ToManagedClusterAddonProfileOutputWithContext(ctx context.Context) ManagedClusterAddonProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAddonProfileOutput)
}





type ManagedClusterAddonProfileMapInput interface {
	pulumi.Input

	ToManagedClusterAddonProfileMapOutput() ManagedClusterAddonProfileMapOutput
	ToManagedClusterAddonProfileMapOutputWithContext(context.Context) ManagedClusterAddonProfileMapOutput
}

type ManagedClusterAddonProfileMap map[string]ManagedClusterAddonProfileInput

func (ManagedClusterAddonProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterAddonProfile)(nil)).Elem()
}

func (i ManagedClusterAddonProfileMap) ToManagedClusterAddonProfileMapOutput() ManagedClusterAddonProfileMapOutput {
	return i.ToManagedClusterAddonProfileMapOutputWithContext(context.Background())
}

func (i ManagedClusterAddonProfileMap) ToManagedClusterAddonProfileMapOutputWithContext(ctx context.Context) ManagedClusterAddonProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAddonProfileMapOutput)
}

type ManagedClusterAddonProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterAddonProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAddonProfile)(nil)).Elem()
}

func (o ManagedClusterAddonProfileOutput) ToManagedClusterAddonProfileOutput() ManagedClusterAddonProfileOutput {
	return o
}

func (o ManagedClusterAddonProfileOutput) ToManagedClusterAddonProfileOutputWithContext(ctx context.Context) ManagedClusterAddonProfileOutput {
	return o
}

func (o ManagedClusterAddonProfileOutput) Config() pulumi.StringMapOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfile) map[string]string { return v.Config }).(pulumi.StringMapOutput)
}

func (o ManagedClusterAddonProfileOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfile) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type ManagedClusterAddonProfileMapOutput struct{ *pulumi.OutputState }

func (ManagedClusterAddonProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterAddonProfile)(nil)).Elem()
}

func (o ManagedClusterAddonProfileMapOutput) ToManagedClusterAddonProfileMapOutput() ManagedClusterAddonProfileMapOutput {
	return o
}

func (o ManagedClusterAddonProfileMapOutput) ToManagedClusterAddonProfileMapOutputWithContext(ctx context.Context) ManagedClusterAddonProfileMapOutput {
	return o
}

func (o ManagedClusterAddonProfileMapOutput) MapIndex(k pulumi.StringInput) ManagedClusterAddonProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManagedClusterAddonProfile {
		return vs[0].(map[string]ManagedClusterAddonProfile)[vs[1].(string)]
	}).(ManagedClusterAddonProfileOutput)
}

type ManagedClusterAddonProfileResponse struct {
	Config  map[string]string `pulumi:"config"`
	Enabled bool              `pulumi:"enabled"`
}

type ManagedClusterAddonProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterAddonProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAddonProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAddonProfileResponseOutput) ToManagedClusterAddonProfileResponseOutput() ManagedClusterAddonProfileResponseOutput {
	return o
}

func (o ManagedClusterAddonProfileResponseOutput) ToManagedClusterAddonProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAddonProfileResponseOutput {
	return o
}

func (o ManagedClusterAddonProfileResponseOutput) Config() pulumi.StringMapOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfileResponse) map[string]string { return v.Config }).(pulumi.StringMapOutput)
}

func (o ManagedClusterAddonProfileResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ManagedClusterAddonProfileResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type ManagedClusterAddonProfileResponseMapOutput struct{ *pulumi.OutputState }

func (ManagedClusterAddonProfileResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedClusterAddonProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAddonProfileResponseMapOutput) ToManagedClusterAddonProfileResponseMapOutput() ManagedClusterAddonProfileResponseMapOutput {
	return o
}

func (o ManagedClusterAddonProfileResponseMapOutput) ToManagedClusterAddonProfileResponseMapOutputWithContext(ctx context.Context) ManagedClusterAddonProfileResponseMapOutput {
	return o
}

func (o ManagedClusterAddonProfileResponseMapOutput) MapIndex(k pulumi.StringInput) ManagedClusterAddonProfileResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManagedClusterAddonProfileResponse {
		return vs[0].(map[string]ManagedClusterAddonProfileResponse)[vs[1].(string)]
	}).(ManagedClusterAddonProfileResponseOutput)
}

type ManagedClusterAgentPoolProfile struct {
	AvailabilityZones   []string `pulumi:"availabilityZones"`
	Count               int      `pulumi:"count"`
	EnableAutoScaling   *bool    `pulumi:"enableAutoScaling"`
	MaxCount            *int     `pulumi:"maxCount"`
	MaxPods             *int     `pulumi:"maxPods"`
	MinCount            *int     `pulumi:"minCount"`
	Name                string   `pulumi:"name"`
	OrchestratorVersion *string  `pulumi:"orchestratorVersion"`
	OsDiskSizeGB        *int     `pulumi:"osDiskSizeGB"`
	OsType              *string  `pulumi:"osType"`
	Type                *string  `pulumi:"type"`
	VmSize              string   `pulumi:"vmSize"`
	VnetSubnetID        *string  `pulumi:"vnetSubnetID"`
}


func (val *ManagedClusterAgentPoolProfile) Defaults() *ManagedClusterAgentPoolProfile {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		tmp.Count = 1
	}
	return &tmp
}





type ManagedClusterAgentPoolProfileInput interface {
	pulumi.Input

	ToManagedClusterAgentPoolProfileOutput() ManagedClusterAgentPoolProfileOutput
	ToManagedClusterAgentPoolProfileOutputWithContext(context.Context) ManagedClusterAgentPoolProfileOutput
}

type ManagedClusterAgentPoolProfileArgs struct {
	AvailabilityZones   pulumi.StringArrayInput `pulumi:"availabilityZones"`
	Count               pulumi.IntInput         `pulumi:"count"`
	EnableAutoScaling   pulumi.BoolPtrInput     `pulumi:"enableAutoScaling"`
	MaxCount            pulumi.IntPtrInput      `pulumi:"maxCount"`
	MaxPods             pulumi.IntPtrInput      `pulumi:"maxPods"`
	MinCount            pulumi.IntPtrInput      `pulumi:"minCount"`
	Name                pulumi.StringInput      `pulumi:"name"`
	OrchestratorVersion pulumi.StringPtrInput   `pulumi:"orchestratorVersion"`
	OsDiskSizeGB        pulumi.IntPtrInput      `pulumi:"osDiskSizeGB"`
	OsType              pulumi.StringPtrInput   `pulumi:"osType"`
	Type                pulumi.StringPtrInput   `pulumi:"type"`
	VmSize              pulumi.StringInput      `pulumi:"vmSize"`
	VnetSubnetID        pulumi.StringPtrInput   `pulumi:"vnetSubnetID"`
}


func (val *ManagedClusterAgentPoolProfileArgs) Defaults() *ManagedClusterAgentPoolProfileArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		tmp.Count = pulumi.Int(1)
	}
	return &tmp
}
func (ManagedClusterAgentPoolProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAgentPoolProfile)(nil)).Elem()
}

func (i ManagedClusterAgentPoolProfileArgs) ToManagedClusterAgentPoolProfileOutput() ManagedClusterAgentPoolProfileOutput {
	return i.ToManagedClusterAgentPoolProfileOutputWithContext(context.Background())
}

func (i ManagedClusterAgentPoolProfileArgs) ToManagedClusterAgentPoolProfileOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAgentPoolProfileOutput)
}





type ManagedClusterAgentPoolProfileArrayInput interface {
	pulumi.Input

	ToManagedClusterAgentPoolProfileArrayOutput() ManagedClusterAgentPoolProfileArrayOutput
	ToManagedClusterAgentPoolProfileArrayOutputWithContext(context.Context) ManagedClusterAgentPoolProfileArrayOutput
}

type ManagedClusterAgentPoolProfileArray []ManagedClusterAgentPoolProfileInput

func (ManagedClusterAgentPoolProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterAgentPoolProfile)(nil)).Elem()
}

func (i ManagedClusterAgentPoolProfileArray) ToManagedClusterAgentPoolProfileArrayOutput() ManagedClusterAgentPoolProfileArrayOutput {
	return i.ToManagedClusterAgentPoolProfileArrayOutputWithContext(context.Background())
}

func (i ManagedClusterAgentPoolProfileArray) ToManagedClusterAgentPoolProfileArrayOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterAgentPoolProfileArrayOutput)
}

type ManagedClusterAgentPoolProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterAgentPoolProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAgentPoolProfile)(nil)).Elem()
}

func (o ManagedClusterAgentPoolProfileOutput) ToManagedClusterAgentPoolProfileOutput() ManagedClusterAgentPoolProfileOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileOutput) ToManagedClusterAgentPoolProfileOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) int { return v.Count }).(pulumi.IntOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) EnableAutoScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *bool { return v.EnableAutoScaling }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *int { return v.MaxCount }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *int { return v.MaxPods }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *int { return v.MinCount }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) string { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) OrchestratorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.OrchestratorVersion }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *int { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) VmSize() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) string { return v.VmSize }).(pulumi.StringOutput)
}

func (o ManagedClusterAgentPoolProfileOutput) VnetSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfile) *string { return v.VnetSubnetID }).(pulumi.StringPtrOutput)
}

type ManagedClusterAgentPoolProfileArrayOutput struct{ *pulumi.OutputState }

func (ManagedClusterAgentPoolProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterAgentPoolProfile)(nil)).Elem()
}

func (o ManagedClusterAgentPoolProfileArrayOutput) ToManagedClusterAgentPoolProfileArrayOutput() ManagedClusterAgentPoolProfileArrayOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileArrayOutput) ToManagedClusterAgentPoolProfileArrayOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileArrayOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileArrayOutput) Index(i pulumi.IntInput) ManagedClusterAgentPoolProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedClusterAgentPoolProfile {
		return vs[0].([]ManagedClusterAgentPoolProfile)[vs[1].(int)]
	}).(ManagedClusterAgentPoolProfileOutput)
}

type ManagedClusterAgentPoolProfileResponse struct {
	AvailabilityZones   []string `pulumi:"availabilityZones"`
	Count               int      `pulumi:"count"`
	EnableAutoScaling   *bool    `pulumi:"enableAutoScaling"`
	MaxCount            *int     `pulumi:"maxCount"`
	MaxPods             *int     `pulumi:"maxPods"`
	MinCount            *int     `pulumi:"minCount"`
	Name                string   `pulumi:"name"`
	OrchestratorVersion *string  `pulumi:"orchestratorVersion"`
	OsDiskSizeGB        *int     `pulumi:"osDiskSizeGB"`
	OsType              *string  `pulumi:"osType"`
	ProvisioningState   string   `pulumi:"provisioningState"`
	Type                *string  `pulumi:"type"`
	VmSize              string   `pulumi:"vmSize"`
	VnetSubnetID        *string  `pulumi:"vnetSubnetID"`
}


func (val *ManagedClusterAgentPoolProfileResponse) Defaults() *ManagedClusterAgentPoolProfileResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Count) {
		tmp.Count = 1
	}
	return &tmp
}

type ManagedClusterAgentPoolProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterAgentPoolProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterAgentPoolProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAgentPoolProfileResponseOutput) ToManagedClusterAgentPoolProfileResponseOutput() ManagedClusterAgentPoolProfileResponseOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileResponseOutput) ToManagedClusterAgentPoolProfileResponseOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileResponseOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileResponseOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) int { return v.Count }).(pulumi.IntOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) EnableAutoScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *bool { return v.EnableAutoScaling }).(pulumi.BoolPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *int { return v.MaxCount }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *int { return v.MaxPods }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *int { return v.MinCount }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) OrchestratorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.OrchestratorVersion }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *int { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) VmSize() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) string { return v.VmSize }).(pulumi.StringOutput)
}

func (o ManagedClusterAgentPoolProfileResponseOutput) VnetSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterAgentPoolProfileResponse) *string { return v.VnetSubnetID }).(pulumi.StringPtrOutput)
}

type ManagedClusterAgentPoolProfileResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagedClusterAgentPoolProfileResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagedClusterAgentPoolProfileResponse)(nil)).Elem()
}

func (o ManagedClusterAgentPoolProfileResponseArrayOutput) ToManagedClusterAgentPoolProfileResponseArrayOutput() ManagedClusterAgentPoolProfileResponseArrayOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileResponseArrayOutput) ToManagedClusterAgentPoolProfileResponseArrayOutputWithContext(ctx context.Context) ManagedClusterAgentPoolProfileResponseArrayOutput {
	return o
}

func (o ManagedClusterAgentPoolProfileResponseArrayOutput) Index(i pulumi.IntInput) ManagedClusterAgentPoolProfileResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagedClusterAgentPoolProfileResponse {
		return vs[0].([]ManagedClusterAgentPoolProfileResponse)[vs[1].(int)]
	}).(ManagedClusterAgentPoolProfileResponseOutput)
}

type ManagedClusterIdentity struct {
	Type *ResourceIdentityType `pulumi:"type"`
}





type ManagedClusterIdentityInput interface {
	pulumi.Input

	ToManagedClusterIdentityOutput() ManagedClusterIdentityOutput
	ToManagedClusterIdentityOutputWithContext(context.Context) ManagedClusterIdentityOutput
}

type ManagedClusterIdentityArgs struct {
	Type ResourceIdentityTypePtrInput `pulumi:"type"`
}

func (ManagedClusterIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterIdentity)(nil)).Elem()
}

func (i ManagedClusterIdentityArgs) ToManagedClusterIdentityOutput() ManagedClusterIdentityOutput {
	return i.ToManagedClusterIdentityOutputWithContext(context.Background())
}

func (i ManagedClusterIdentityArgs) ToManagedClusterIdentityOutputWithContext(ctx context.Context) ManagedClusterIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterIdentityOutput)
}

func (i ManagedClusterIdentityArgs) ToManagedClusterIdentityPtrOutput() ManagedClusterIdentityPtrOutput {
	return i.ToManagedClusterIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedClusterIdentityArgs) ToManagedClusterIdentityPtrOutputWithContext(ctx context.Context) ManagedClusterIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterIdentityOutput).ToManagedClusterIdentityPtrOutputWithContext(ctx)
}









type ManagedClusterIdentityPtrInput interface {
	pulumi.Input

	ToManagedClusterIdentityPtrOutput() ManagedClusterIdentityPtrOutput
	ToManagedClusterIdentityPtrOutputWithContext(context.Context) ManagedClusterIdentityPtrOutput
}

type managedClusterIdentityPtrType ManagedClusterIdentityArgs

func ManagedClusterIdentityPtr(v *ManagedClusterIdentityArgs) ManagedClusterIdentityPtrInput {
	return (*managedClusterIdentityPtrType)(v)
}

func (*managedClusterIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterIdentity)(nil)).Elem()
}

func (i *managedClusterIdentityPtrType) ToManagedClusterIdentityPtrOutput() ManagedClusterIdentityPtrOutput {
	return i.ToManagedClusterIdentityPtrOutputWithContext(context.Background())
}

func (i *managedClusterIdentityPtrType) ToManagedClusterIdentityPtrOutputWithContext(ctx context.Context) ManagedClusterIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterIdentityPtrOutput)
}

type ManagedClusterIdentityOutput struct{ *pulumi.OutputState }

func (ManagedClusterIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterIdentity)(nil)).Elem()
}

func (o ManagedClusterIdentityOutput) ToManagedClusterIdentityOutput() ManagedClusterIdentityOutput {
	return o
}

func (o ManagedClusterIdentityOutput) ToManagedClusterIdentityOutputWithContext(ctx context.Context) ManagedClusterIdentityOutput {
	return o
}

func (o ManagedClusterIdentityOutput) ToManagedClusterIdentityPtrOutput() ManagedClusterIdentityPtrOutput {
	return o.ToManagedClusterIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedClusterIdentityOutput) ToManagedClusterIdentityPtrOutputWithContext(ctx context.Context) ManagedClusterIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterIdentity) *ManagedClusterIdentity {
		return &v
	}).(ManagedClusterIdentityPtrOutput)
}

func (o ManagedClusterIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v ManagedClusterIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

type ManagedClusterIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterIdentity)(nil)).Elem()
}

func (o ManagedClusterIdentityPtrOutput) ToManagedClusterIdentityPtrOutput() ManagedClusterIdentityPtrOutput {
	return o
}

func (o ManagedClusterIdentityPtrOutput) ToManagedClusterIdentityPtrOutputWithContext(ctx context.Context) ManagedClusterIdentityPtrOutput {
	return o
}

func (o ManagedClusterIdentityPtrOutput) Elem() ManagedClusterIdentityOutput {
	return o.ApplyT(func(v *ManagedClusterIdentity) ManagedClusterIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedClusterIdentity
		return ret
	}).(ManagedClusterIdentityOutput)
}

func (o ManagedClusterIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ManagedClusterIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

type ManagedClusterIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type ManagedClusterIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterIdentityResponse)(nil)).Elem()
}

func (o ManagedClusterIdentityResponseOutput) ToManagedClusterIdentityResponseOutput() ManagedClusterIdentityResponseOutput {
	return o
}

func (o ManagedClusterIdentityResponseOutput) ToManagedClusterIdentityResponseOutputWithContext(ctx context.Context) ManagedClusterIdentityResponseOutput {
	return o
}

func (o ManagedClusterIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedClusterIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedClusterIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ManagedClusterIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterIdentityResponse)(nil)).Elem()
}

func (o ManagedClusterIdentityResponsePtrOutput) ToManagedClusterIdentityResponsePtrOutput() ManagedClusterIdentityResponsePtrOutput {
	return o
}

func (o ManagedClusterIdentityResponsePtrOutput) ToManagedClusterIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedClusterIdentityResponsePtrOutput {
	return o
}

func (o ManagedClusterIdentityResponsePtrOutput) Elem() ManagedClusterIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedClusterIdentityResponse) ManagedClusterIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterIdentityResponse
		return ret
	}).(ManagedClusterIdentityResponseOutput)
}

func (o ManagedClusterIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterServicePrincipalProfile struct {
	ClientId string  `pulumi:"clientId"`
	Secret   *string `pulumi:"secret"`
}





type ManagedClusterServicePrincipalProfileInput interface {
	pulumi.Input

	ToManagedClusterServicePrincipalProfileOutput() ManagedClusterServicePrincipalProfileOutput
	ToManagedClusterServicePrincipalProfileOutputWithContext(context.Context) ManagedClusterServicePrincipalProfileOutput
}

type ManagedClusterServicePrincipalProfileArgs struct {
	ClientId pulumi.StringInput    `pulumi:"clientId"`
	Secret   pulumi.StringPtrInput `pulumi:"secret"`
}

func (ManagedClusterServicePrincipalProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterServicePrincipalProfile)(nil)).Elem()
}

func (i ManagedClusterServicePrincipalProfileArgs) ToManagedClusterServicePrincipalProfileOutput() ManagedClusterServicePrincipalProfileOutput {
	return i.ToManagedClusterServicePrincipalProfileOutputWithContext(context.Background())
}

func (i ManagedClusterServicePrincipalProfileArgs) ToManagedClusterServicePrincipalProfileOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterServicePrincipalProfileOutput)
}

func (i ManagedClusterServicePrincipalProfileArgs) ToManagedClusterServicePrincipalProfilePtrOutput() ManagedClusterServicePrincipalProfilePtrOutput {
	return i.ToManagedClusterServicePrincipalProfilePtrOutputWithContext(context.Background())
}

func (i ManagedClusterServicePrincipalProfileArgs) ToManagedClusterServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterServicePrincipalProfileOutput).ToManagedClusterServicePrincipalProfilePtrOutputWithContext(ctx)
}









type ManagedClusterServicePrincipalProfilePtrInput interface {
	pulumi.Input

	ToManagedClusterServicePrincipalProfilePtrOutput() ManagedClusterServicePrincipalProfilePtrOutput
	ToManagedClusterServicePrincipalProfilePtrOutputWithContext(context.Context) ManagedClusterServicePrincipalProfilePtrOutput
}

type managedClusterServicePrincipalProfilePtrType ManagedClusterServicePrincipalProfileArgs

func ManagedClusterServicePrincipalProfilePtr(v *ManagedClusterServicePrincipalProfileArgs) ManagedClusterServicePrincipalProfilePtrInput {
	return (*managedClusterServicePrincipalProfilePtrType)(v)
}

func (*managedClusterServicePrincipalProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterServicePrincipalProfile)(nil)).Elem()
}

func (i *managedClusterServicePrincipalProfilePtrType) ToManagedClusterServicePrincipalProfilePtrOutput() ManagedClusterServicePrincipalProfilePtrOutput {
	return i.ToManagedClusterServicePrincipalProfilePtrOutputWithContext(context.Background())
}

func (i *managedClusterServicePrincipalProfilePtrType) ToManagedClusterServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterServicePrincipalProfilePtrOutput)
}

type ManagedClusterServicePrincipalProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterServicePrincipalProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterServicePrincipalProfile)(nil)).Elem()
}

func (o ManagedClusterServicePrincipalProfileOutput) ToManagedClusterServicePrincipalProfileOutput() ManagedClusterServicePrincipalProfileOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfileOutput) ToManagedClusterServicePrincipalProfileOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfileOutput) ToManagedClusterServicePrincipalProfilePtrOutput() ManagedClusterServicePrincipalProfilePtrOutput {
	return o.ToManagedClusterServicePrincipalProfilePtrOutputWithContext(context.Background())
}

func (o ManagedClusterServicePrincipalProfileOutput) ToManagedClusterServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterServicePrincipalProfile) *ManagedClusterServicePrincipalProfile {
		return &v
	}).(ManagedClusterServicePrincipalProfilePtrOutput)
}

func (o ManagedClusterServicePrincipalProfileOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterServicePrincipalProfile) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ManagedClusterServicePrincipalProfileOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterServicePrincipalProfile) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

type ManagedClusterServicePrincipalProfilePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterServicePrincipalProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterServicePrincipalProfile)(nil)).Elem()
}

func (o ManagedClusterServicePrincipalProfilePtrOutput) ToManagedClusterServicePrincipalProfilePtrOutput() ManagedClusterServicePrincipalProfilePtrOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfilePtrOutput) ToManagedClusterServicePrincipalProfilePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfilePtrOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfilePtrOutput) Elem() ManagedClusterServicePrincipalProfileOutput {
	return o.ApplyT(func(v *ManagedClusterServicePrincipalProfile) ManagedClusterServicePrincipalProfile {
		if v != nil {
			return *v
		}
		var ret ManagedClusterServicePrincipalProfile
		return ret
	}).(ManagedClusterServicePrincipalProfileOutput)
}

func (o ManagedClusterServicePrincipalProfilePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterServicePrincipalProfile) *string {
		if v == nil {
			return nil
		}
		return &v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterServicePrincipalProfilePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterServicePrincipalProfile) *string {
		if v == nil {
			return nil
		}
		return v.Secret
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterServicePrincipalProfileResponse struct {
	ClientId string  `pulumi:"clientId"`
	Secret   *string `pulumi:"secret"`
}

type ManagedClusterServicePrincipalProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterServicePrincipalProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterServicePrincipalProfileResponse)(nil)).Elem()
}

func (o ManagedClusterServicePrincipalProfileResponseOutput) ToManagedClusterServicePrincipalProfileResponseOutput() ManagedClusterServicePrincipalProfileResponseOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfileResponseOutput) ToManagedClusterServicePrincipalProfileResponseOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileResponseOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfileResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterServicePrincipalProfileResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ManagedClusterServicePrincipalProfileResponseOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterServicePrincipalProfileResponse) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

type ManagedClusterServicePrincipalProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterServicePrincipalProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterServicePrincipalProfileResponse)(nil)).Elem()
}

func (o ManagedClusterServicePrincipalProfileResponsePtrOutput) ToManagedClusterServicePrincipalProfileResponsePtrOutput() ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfileResponsePtrOutput) ToManagedClusterServicePrincipalProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterServicePrincipalProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterServicePrincipalProfileResponsePtrOutput) Elem() ManagedClusterServicePrincipalProfileResponseOutput {
	return o.ApplyT(func(v *ManagedClusterServicePrincipalProfileResponse) ManagedClusterServicePrincipalProfileResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterServicePrincipalProfileResponse
		return ret
	}).(ManagedClusterServicePrincipalProfileResponseOutput)
}

func (o ManagedClusterServicePrincipalProfileResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterServicePrincipalProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterServicePrincipalProfileResponsePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterServicePrincipalProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Secret
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterWindowsProfile struct {
	AdminPassword *string `pulumi:"adminPassword"`
	AdminUsername string  `pulumi:"adminUsername"`
}





type ManagedClusterWindowsProfileInput interface {
	pulumi.Input

	ToManagedClusterWindowsProfileOutput() ManagedClusterWindowsProfileOutput
	ToManagedClusterWindowsProfileOutputWithContext(context.Context) ManagedClusterWindowsProfileOutput
}

type ManagedClusterWindowsProfileArgs struct {
	AdminPassword pulumi.StringPtrInput `pulumi:"adminPassword"`
	AdminUsername pulumi.StringInput    `pulumi:"adminUsername"`
}

func (ManagedClusterWindowsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterWindowsProfile)(nil)).Elem()
}

func (i ManagedClusterWindowsProfileArgs) ToManagedClusterWindowsProfileOutput() ManagedClusterWindowsProfileOutput {
	return i.ToManagedClusterWindowsProfileOutputWithContext(context.Background())
}

func (i ManagedClusterWindowsProfileArgs) ToManagedClusterWindowsProfileOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterWindowsProfileOutput)
}

func (i ManagedClusterWindowsProfileArgs) ToManagedClusterWindowsProfilePtrOutput() ManagedClusterWindowsProfilePtrOutput {
	return i.ToManagedClusterWindowsProfilePtrOutputWithContext(context.Background())
}

func (i ManagedClusterWindowsProfileArgs) ToManagedClusterWindowsProfilePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterWindowsProfileOutput).ToManagedClusterWindowsProfilePtrOutputWithContext(ctx)
}









type ManagedClusterWindowsProfilePtrInput interface {
	pulumi.Input

	ToManagedClusterWindowsProfilePtrOutput() ManagedClusterWindowsProfilePtrOutput
	ToManagedClusterWindowsProfilePtrOutputWithContext(context.Context) ManagedClusterWindowsProfilePtrOutput
}

type managedClusterWindowsProfilePtrType ManagedClusterWindowsProfileArgs

func ManagedClusterWindowsProfilePtr(v *ManagedClusterWindowsProfileArgs) ManagedClusterWindowsProfilePtrInput {
	return (*managedClusterWindowsProfilePtrType)(v)
}

func (*managedClusterWindowsProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterWindowsProfile)(nil)).Elem()
}

func (i *managedClusterWindowsProfilePtrType) ToManagedClusterWindowsProfilePtrOutput() ManagedClusterWindowsProfilePtrOutput {
	return i.ToManagedClusterWindowsProfilePtrOutputWithContext(context.Background())
}

func (i *managedClusterWindowsProfilePtrType) ToManagedClusterWindowsProfilePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterWindowsProfilePtrOutput)
}

type ManagedClusterWindowsProfileOutput struct{ *pulumi.OutputState }

func (ManagedClusterWindowsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterWindowsProfile)(nil)).Elem()
}

func (o ManagedClusterWindowsProfileOutput) ToManagedClusterWindowsProfileOutput() ManagedClusterWindowsProfileOutput {
	return o
}

func (o ManagedClusterWindowsProfileOutput) ToManagedClusterWindowsProfileOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileOutput {
	return o
}

func (o ManagedClusterWindowsProfileOutput) ToManagedClusterWindowsProfilePtrOutput() ManagedClusterWindowsProfilePtrOutput {
	return o.ToManagedClusterWindowsProfilePtrOutputWithContext(context.Background())
}

func (o ManagedClusterWindowsProfileOutput) ToManagedClusterWindowsProfilePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedClusterWindowsProfile) *ManagedClusterWindowsProfile {
		return &v
	}).(ManagedClusterWindowsProfilePtrOutput)
}

func (o ManagedClusterWindowsProfileOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterWindowsProfile) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterWindowsProfileOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterWindowsProfile) string { return v.AdminUsername }).(pulumi.StringOutput)
}

type ManagedClusterWindowsProfilePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterWindowsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterWindowsProfile)(nil)).Elem()
}

func (o ManagedClusterWindowsProfilePtrOutput) ToManagedClusterWindowsProfilePtrOutput() ManagedClusterWindowsProfilePtrOutput {
	return o
}

func (o ManagedClusterWindowsProfilePtrOutput) ToManagedClusterWindowsProfilePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfilePtrOutput {
	return o
}

func (o ManagedClusterWindowsProfilePtrOutput) Elem() ManagedClusterWindowsProfileOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfile) ManagedClusterWindowsProfile {
		if v != nil {
			return *v
		}
		var ret ManagedClusterWindowsProfile
		return ret
	}).(ManagedClusterWindowsProfileOutput)
}

func (o ManagedClusterWindowsProfilePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterWindowsProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfile) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

type ManagedClusterWindowsProfileResponse struct {
	AdminPassword *string `pulumi:"adminPassword"`
	AdminUsername string  `pulumi:"adminUsername"`
}

type ManagedClusterWindowsProfileResponseOutput struct{ *pulumi.OutputState }

func (ManagedClusterWindowsProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedClusterWindowsProfileResponse)(nil)).Elem()
}

func (o ManagedClusterWindowsProfileResponseOutput) ToManagedClusterWindowsProfileResponseOutput() ManagedClusterWindowsProfileResponseOutput {
	return o
}

func (o ManagedClusterWindowsProfileResponseOutput) ToManagedClusterWindowsProfileResponseOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileResponseOutput {
	return o
}

func (o ManagedClusterWindowsProfileResponseOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedClusterWindowsProfileResponse) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o ManagedClusterWindowsProfileResponseOutput) AdminUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedClusterWindowsProfileResponse) string { return v.AdminUsername }).(pulumi.StringOutput)
}

type ManagedClusterWindowsProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedClusterWindowsProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterWindowsProfileResponse)(nil)).Elem()
}

func (o ManagedClusterWindowsProfileResponsePtrOutput) ToManagedClusterWindowsProfileResponsePtrOutput() ManagedClusterWindowsProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterWindowsProfileResponsePtrOutput) ToManagedClusterWindowsProfileResponsePtrOutputWithContext(ctx context.Context) ManagedClusterWindowsProfileResponsePtrOutput {
	return o
}

func (o ManagedClusterWindowsProfileResponsePtrOutput) Elem() ManagedClusterWindowsProfileResponseOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfileResponse) ManagedClusterWindowsProfileResponse {
		if v != nil {
			return *v
		}
		var ret ManagedClusterWindowsProfileResponse
		return ret
	}).(ManagedClusterWindowsProfileResponseOutput)
}

func (o ManagedClusterWindowsProfileResponsePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o ManagedClusterWindowsProfileResponsePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterWindowsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ContainerServiceLinuxProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceLinuxProfilePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceLinuxProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceLinuxProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceNetworkProfileOutput{})
	pulumi.RegisterOutputType(ContainerServiceNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceNetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceNetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshConfigurationOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyArrayOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyResponseOutput{})
	pulumi.RegisterOutputType(ContainerServiceSshPublicKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(CredentialResultResponseOutput{})
	pulumi.RegisterOutputType(CredentialResultResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedClusterAADProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterAADProfilePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterAADProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterAADProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterAddonProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterAddonProfileMapOutput{})
	pulumi.RegisterOutputType(ManagedClusterAddonProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterAddonProfileResponseMapOutput{})
	pulumi.RegisterOutputType(ManagedClusterAgentPoolProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterAgentPoolProfileArrayOutput{})
	pulumi.RegisterOutputType(ManagedClusterAgentPoolProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterAgentPoolProfileResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedClusterIdentityOutput{})
	pulumi.RegisterOutputType(ManagedClusterIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterServicePrincipalProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterServicePrincipalProfilePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterServicePrincipalProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterServicePrincipalProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterWindowsProfileOutput{})
	pulumi.RegisterOutputType(ManagedClusterWindowsProfilePtrOutput{})
	pulumi.RegisterOutputType(ManagedClusterWindowsProfileResponseOutput{})
	pulumi.RegisterOutputType(ManagedClusterWindowsProfileResponsePtrOutput{})
}
