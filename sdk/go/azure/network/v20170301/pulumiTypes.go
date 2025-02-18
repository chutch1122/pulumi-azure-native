


package v20170301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AddressSpace struct {
	AddressPrefixes []string `pulumi:"addressPrefixes"`
}





type AddressSpaceInput interface {
	pulumi.Input

	ToAddressSpaceOutput() AddressSpaceOutput
	ToAddressSpaceOutputWithContext(context.Context) AddressSpaceOutput
}

type AddressSpaceArgs struct {
	AddressPrefixes pulumi.StringArrayInput `pulumi:"addressPrefixes"`
}

func (AddressSpaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressSpace)(nil)).Elem()
}

func (i AddressSpaceArgs) ToAddressSpaceOutput() AddressSpaceOutput {
	return i.ToAddressSpaceOutputWithContext(context.Background())
}

func (i AddressSpaceArgs) ToAddressSpaceOutputWithContext(ctx context.Context) AddressSpaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressSpaceOutput)
}

func (i AddressSpaceArgs) ToAddressSpacePtrOutput() AddressSpacePtrOutput {
	return i.ToAddressSpacePtrOutputWithContext(context.Background())
}

func (i AddressSpaceArgs) ToAddressSpacePtrOutputWithContext(ctx context.Context) AddressSpacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressSpaceOutput).ToAddressSpacePtrOutputWithContext(ctx)
}









type AddressSpacePtrInput interface {
	pulumi.Input

	ToAddressSpacePtrOutput() AddressSpacePtrOutput
	ToAddressSpacePtrOutputWithContext(context.Context) AddressSpacePtrOutput
}

type addressSpacePtrType AddressSpaceArgs

func AddressSpacePtr(v *AddressSpaceArgs) AddressSpacePtrInput {
	return (*addressSpacePtrType)(v)
}

func (*addressSpacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressSpace)(nil)).Elem()
}

func (i *addressSpacePtrType) ToAddressSpacePtrOutput() AddressSpacePtrOutput {
	return i.ToAddressSpacePtrOutputWithContext(context.Background())
}

func (i *addressSpacePtrType) ToAddressSpacePtrOutputWithContext(ctx context.Context) AddressSpacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressSpacePtrOutput)
}

type AddressSpaceOutput struct{ *pulumi.OutputState }

func (AddressSpaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressSpace)(nil)).Elem()
}

func (o AddressSpaceOutput) ToAddressSpaceOutput() AddressSpaceOutput {
	return o
}

func (o AddressSpaceOutput) ToAddressSpaceOutputWithContext(ctx context.Context) AddressSpaceOutput {
	return o
}

func (o AddressSpaceOutput) ToAddressSpacePtrOutput() AddressSpacePtrOutput {
	return o.ToAddressSpacePtrOutputWithContext(context.Background())
}

func (o AddressSpaceOutput) ToAddressSpacePtrOutputWithContext(ctx context.Context) AddressSpacePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AddressSpace) *AddressSpace {
		return &v
	}).(AddressSpacePtrOutput)
}

func (o AddressSpaceOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AddressSpace) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

type AddressSpacePtrOutput struct{ *pulumi.OutputState }

func (AddressSpacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressSpace)(nil)).Elem()
}

func (o AddressSpacePtrOutput) ToAddressSpacePtrOutput() AddressSpacePtrOutput {
	return o
}

func (o AddressSpacePtrOutput) ToAddressSpacePtrOutputWithContext(ctx context.Context) AddressSpacePtrOutput {
	return o
}

func (o AddressSpacePtrOutput) Elem() AddressSpaceOutput {
	return o.ApplyT(func(v *AddressSpace) AddressSpace {
		if v != nil {
			return *v
		}
		var ret AddressSpace
		return ret
	}).(AddressSpaceOutput)
}

func (o AddressSpacePtrOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AddressSpace) []string {
		if v == nil {
			return nil
		}
		return v.AddressPrefixes
	}).(pulumi.StringArrayOutput)
}

type AddressSpaceResponse struct {
	AddressPrefixes []string `pulumi:"addressPrefixes"`
}

type AddressSpaceResponseOutput struct{ *pulumi.OutputState }

func (AddressSpaceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddressSpaceResponse)(nil)).Elem()
}

func (o AddressSpaceResponseOutput) ToAddressSpaceResponseOutput() AddressSpaceResponseOutput {
	return o
}

func (o AddressSpaceResponseOutput) ToAddressSpaceResponseOutputWithContext(ctx context.Context) AddressSpaceResponseOutput {
	return o
}

func (o AddressSpaceResponseOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AddressSpaceResponse) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

type AddressSpaceResponsePtrOutput struct{ *pulumi.OutputState }

func (AddressSpaceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddressSpaceResponse)(nil)).Elem()
}

func (o AddressSpaceResponsePtrOutput) ToAddressSpaceResponsePtrOutput() AddressSpaceResponsePtrOutput {
	return o
}

func (o AddressSpaceResponsePtrOutput) ToAddressSpaceResponsePtrOutputWithContext(ctx context.Context) AddressSpaceResponsePtrOutput {
	return o
}

func (o AddressSpaceResponsePtrOutput) Elem() AddressSpaceResponseOutput {
	return o.ApplyT(func(v *AddressSpaceResponse) AddressSpaceResponse {
		if v != nil {
			return *v
		}
		var ret AddressSpaceResponse
		return ret
	}).(AddressSpaceResponseOutput)
}

func (o AddressSpaceResponsePtrOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AddressSpaceResponse) []string {
		if v == nil {
			return nil
		}
		return v.AddressPrefixes
	}).(pulumi.StringArrayOutput)
}

type ApplicationGatewayAuthenticationCertificate struct {
	Data              *string `pulumi:"data"`
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	ProvisioningState *string `pulumi:"provisioningState"`
}





type ApplicationGatewayAuthenticationCertificateInput interface {
	pulumi.Input

	ToApplicationGatewayAuthenticationCertificateOutput() ApplicationGatewayAuthenticationCertificateOutput
	ToApplicationGatewayAuthenticationCertificateOutputWithContext(context.Context) ApplicationGatewayAuthenticationCertificateOutput
}

type ApplicationGatewayAuthenticationCertificateArgs struct {
	Data              pulumi.StringPtrInput `pulumi:"data"`
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (ApplicationGatewayAuthenticationCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayAuthenticationCertificate)(nil)).Elem()
}

func (i ApplicationGatewayAuthenticationCertificateArgs) ToApplicationGatewayAuthenticationCertificateOutput() ApplicationGatewayAuthenticationCertificateOutput {
	return i.ToApplicationGatewayAuthenticationCertificateOutputWithContext(context.Background())
}

func (i ApplicationGatewayAuthenticationCertificateArgs) ToApplicationGatewayAuthenticationCertificateOutputWithContext(ctx context.Context) ApplicationGatewayAuthenticationCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayAuthenticationCertificateOutput)
}





type ApplicationGatewayAuthenticationCertificateArrayInput interface {
	pulumi.Input

	ToApplicationGatewayAuthenticationCertificateArrayOutput() ApplicationGatewayAuthenticationCertificateArrayOutput
	ToApplicationGatewayAuthenticationCertificateArrayOutputWithContext(context.Context) ApplicationGatewayAuthenticationCertificateArrayOutput
}

type ApplicationGatewayAuthenticationCertificateArray []ApplicationGatewayAuthenticationCertificateInput

func (ApplicationGatewayAuthenticationCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayAuthenticationCertificate)(nil)).Elem()
}

func (i ApplicationGatewayAuthenticationCertificateArray) ToApplicationGatewayAuthenticationCertificateArrayOutput() ApplicationGatewayAuthenticationCertificateArrayOutput {
	return i.ToApplicationGatewayAuthenticationCertificateArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayAuthenticationCertificateArray) ToApplicationGatewayAuthenticationCertificateArrayOutputWithContext(ctx context.Context) ApplicationGatewayAuthenticationCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayAuthenticationCertificateArrayOutput)
}

type ApplicationGatewayAuthenticationCertificateOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayAuthenticationCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayAuthenticationCertificate)(nil)).Elem()
}

func (o ApplicationGatewayAuthenticationCertificateOutput) ToApplicationGatewayAuthenticationCertificateOutput() ApplicationGatewayAuthenticationCertificateOutput {
	return o
}

func (o ApplicationGatewayAuthenticationCertificateOutput) ToApplicationGatewayAuthenticationCertificateOutputWithContext(ctx context.Context) ApplicationGatewayAuthenticationCertificateOutput {
	return o
}

func (o ApplicationGatewayAuthenticationCertificateOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificate) *string { return v.Data }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayAuthenticationCertificateOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificate) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayAuthenticationCertificateOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificate) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayAuthenticationCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayAuthenticationCertificateOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificate) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayAuthenticationCertificateArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayAuthenticationCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayAuthenticationCertificate)(nil)).Elem()
}

func (o ApplicationGatewayAuthenticationCertificateArrayOutput) ToApplicationGatewayAuthenticationCertificateArrayOutput() ApplicationGatewayAuthenticationCertificateArrayOutput {
	return o
}

func (o ApplicationGatewayAuthenticationCertificateArrayOutput) ToApplicationGatewayAuthenticationCertificateArrayOutputWithContext(ctx context.Context) ApplicationGatewayAuthenticationCertificateArrayOutput {
	return o
}

func (o ApplicationGatewayAuthenticationCertificateArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayAuthenticationCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayAuthenticationCertificate {
		return vs[0].([]ApplicationGatewayAuthenticationCertificate)[vs[1].(int)]
	}).(ApplicationGatewayAuthenticationCertificateOutput)
}

type ApplicationGatewayAuthenticationCertificateResponse struct {
	Data              *string `pulumi:"data"`
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	ProvisioningState *string `pulumi:"provisioningState"`
}

type ApplicationGatewayAuthenticationCertificateResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayAuthenticationCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayAuthenticationCertificateResponse)(nil)).Elem()
}

func (o ApplicationGatewayAuthenticationCertificateResponseOutput) ToApplicationGatewayAuthenticationCertificateResponseOutput() ApplicationGatewayAuthenticationCertificateResponseOutput {
	return o
}

func (o ApplicationGatewayAuthenticationCertificateResponseOutput) ToApplicationGatewayAuthenticationCertificateResponseOutputWithContext(ctx context.Context) ApplicationGatewayAuthenticationCertificateResponseOutput {
	return o
}

func (o ApplicationGatewayAuthenticationCertificateResponseOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificateResponse) *string { return v.Data }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayAuthenticationCertificateResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificateResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayAuthenticationCertificateResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificateResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayAuthenticationCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayAuthenticationCertificateResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayAuthenticationCertificateResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayAuthenticationCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayAuthenticationCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayAuthenticationCertificateResponse)(nil)).Elem()
}

func (o ApplicationGatewayAuthenticationCertificateResponseArrayOutput) ToApplicationGatewayAuthenticationCertificateResponseArrayOutput() ApplicationGatewayAuthenticationCertificateResponseArrayOutput {
	return o
}

func (o ApplicationGatewayAuthenticationCertificateResponseArrayOutput) ToApplicationGatewayAuthenticationCertificateResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayAuthenticationCertificateResponseArrayOutput {
	return o
}

func (o ApplicationGatewayAuthenticationCertificateResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayAuthenticationCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayAuthenticationCertificateResponse {
		return vs[0].([]ApplicationGatewayAuthenticationCertificateResponse)[vs[1].(int)]
	}).(ApplicationGatewayAuthenticationCertificateResponseOutput)
}

type ApplicationGatewayBackendAddress struct {
	Fqdn      *string `pulumi:"fqdn"`
	IpAddress *string `pulumi:"ipAddress"`
}





type ApplicationGatewayBackendAddressInput interface {
	pulumi.Input

	ToApplicationGatewayBackendAddressOutput() ApplicationGatewayBackendAddressOutput
	ToApplicationGatewayBackendAddressOutputWithContext(context.Context) ApplicationGatewayBackendAddressOutput
}

type ApplicationGatewayBackendAddressArgs struct {
	Fqdn      pulumi.StringPtrInput `pulumi:"fqdn"`
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
}

func (ApplicationGatewayBackendAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendAddress)(nil)).Elem()
}

func (i ApplicationGatewayBackendAddressArgs) ToApplicationGatewayBackendAddressOutput() ApplicationGatewayBackendAddressOutput {
	return i.ToApplicationGatewayBackendAddressOutputWithContext(context.Background())
}

func (i ApplicationGatewayBackendAddressArgs) ToApplicationGatewayBackendAddressOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayBackendAddressOutput)
}





type ApplicationGatewayBackendAddressArrayInput interface {
	pulumi.Input

	ToApplicationGatewayBackendAddressArrayOutput() ApplicationGatewayBackendAddressArrayOutput
	ToApplicationGatewayBackendAddressArrayOutputWithContext(context.Context) ApplicationGatewayBackendAddressArrayOutput
}

type ApplicationGatewayBackendAddressArray []ApplicationGatewayBackendAddressInput

func (ApplicationGatewayBackendAddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendAddress)(nil)).Elem()
}

func (i ApplicationGatewayBackendAddressArray) ToApplicationGatewayBackendAddressArrayOutput() ApplicationGatewayBackendAddressArrayOutput {
	return i.ToApplicationGatewayBackendAddressArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayBackendAddressArray) ToApplicationGatewayBackendAddressArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayBackendAddressArrayOutput)
}

type ApplicationGatewayBackendAddressOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendAddress)(nil)).Elem()
}

func (o ApplicationGatewayBackendAddressOutput) ToApplicationGatewayBackendAddressOutput() ApplicationGatewayBackendAddressOutput {
	return o
}

func (o ApplicationGatewayBackendAddressOutput) ToApplicationGatewayBackendAddressOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressOutput {
	return o
}

func (o ApplicationGatewayBackendAddressOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddress) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendAddressOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddress) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayBackendAddressArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendAddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendAddress)(nil)).Elem()
}

func (o ApplicationGatewayBackendAddressArrayOutput) ToApplicationGatewayBackendAddressArrayOutput() ApplicationGatewayBackendAddressArrayOutput {
	return o
}

func (o ApplicationGatewayBackendAddressArrayOutput) ToApplicationGatewayBackendAddressArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressArrayOutput {
	return o
}

func (o ApplicationGatewayBackendAddressArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayBackendAddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayBackendAddress {
		return vs[0].([]ApplicationGatewayBackendAddress)[vs[1].(int)]
	}).(ApplicationGatewayBackendAddressOutput)
}

type ApplicationGatewayBackendAddressPool struct {
	BackendAddresses        []ApplicationGatewayBackendAddress `pulumi:"backendAddresses"`
	BackendIPConfigurations []NetworkInterfaceIPConfiguration  `pulumi:"backendIPConfigurations"`
	Etag                    *string                            `pulumi:"etag"`
	Id                      *string                            `pulumi:"id"`
	Name                    *string                            `pulumi:"name"`
	ProvisioningState       *string                            `pulumi:"provisioningState"`
}





type ApplicationGatewayBackendAddressPoolInput interface {
	pulumi.Input

	ToApplicationGatewayBackendAddressPoolOutput() ApplicationGatewayBackendAddressPoolOutput
	ToApplicationGatewayBackendAddressPoolOutputWithContext(context.Context) ApplicationGatewayBackendAddressPoolOutput
}

type ApplicationGatewayBackendAddressPoolArgs struct {
	BackendAddresses        ApplicationGatewayBackendAddressArrayInput `pulumi:"backendAddresses"`
	BackendIPConfigurations NetworkInterfaceIPConfigurationArrayInput  `pulumi:"backendIPConfigurations"`
	Etag                    pulumi.StringPtrInput                      `pulumi:"etag"`
	Id                      pulumi.StringPtrInput                      `pulumi:"id"`
	Name                    pulumi.StringPtrInput                      `pulumi:"name"`
	ProvisioningState       pulumi.StringPtrInput                      `pulumi:"provisioningState"`
}

func (ApplicationGatewayBackendAddressPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendAddressPool)(nil)).Elem()
}

func (i ApplicationGatewayBackendAddressPoolArgs) ToApplicationGatewayBackendAddressPoolOutput() ApplicationGatewayBackendAddressPoolOutput {
	return i.ToApplicationGatewayBackendAddressPoolOutputWithContext(context.Background())
}

func (i ApplicationGatewayBackendAddressPoolArgs) ToApplicationGatewayBackendAddressPoolOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayBackendAddressPoolOutput)
}





type ApplicationGatewayBackendAddressPoolArrayInput interface {
	pulumi.Input

	ToApplicationGatewayBackendAddressPoolArrayOutput() ApplicationGatewayBackendAddressPoolArrayOutput
	ToApplicationGatewayBackendAddressPoolArrayOutputWithContext(context.Context) ApplicationGatewayBackendAddressPoolArrayOutput
}

type ApplicationGatewayBackendAddressPoolArray []ApplicationGatewayBackendAddressPoolInput

func (ApplicationGatewayBackendAddressPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendAddressPool)(nil)).Elem()
}

func (i ApplicationGatewayBackendAddressPoolArray) ToApplicationGatewayBackendAddressPoolArrayOutput() ApplicationGatewayBackendAddressPoolArrayOutput {
	return i.ToApplicationGatewayBackendAddressPoolArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayBackendAddressPoolArray) ToApplicationGatewayBackendAddressPoolArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayBackendAddressPoolArrayOutput)
}

type ApplicationGatewayBackendAddressPoolOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendAddressPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendAddressPool)(nil)).Elem()
}

func (o ApplicationGatewayBackendAddressPoolOutput) ToApplicationGatewayBackendAddressPoolOutput() ApplicationGatewayBackendAddressPoolOutput {
	return o
}

func (o ApplicationGatewayBackendAddressPoolOutput) ToApplicationGatewayBackendAddressPoolOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressPoolOutput {
	return o
}

func (o ApplicationGatewayBackendAddressPoolOutput) BackendAddresses() ApplicationGatewayBackendAddressArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPool) []ApplicationGatewayBackendAddress {
		return v.BackendAddresses
	}).(ApplicationGatewayBackendAddressArrayOutput)
}

func (o ApplicationGatewayBackendAddressPoolOutput) BackendIPConfigurations() NetworkInterfaceIPConfigurationArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPool) []NetworkInterfaceIPConfiguration {
		return v.BackendIPConfigurations
	}).(NetworkInterfaceIPConfigurationArrayOutput)
}

func (o ApplicationGatewayBackendAddressPoolOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPool) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendAddressPoolOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPool) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendAddressPoolOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPool) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendAddressPoolOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPool) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayBackendAddressPoolArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendAddressPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendAddressPool)(nil)).Elem()
}

func (o ApplicationGatewayBackendAddressPoolArrayOutput) ToApplicationGatewayBackendAddressPoolArrayOutput() ApplicationGatewayBackendAddressPoolArrayOutput {
	return o
}

func (o ApplicationGatewayBackendAddressPoolArrayOutput) ToApplicationGatewayBackendAddressPoolArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressPoolArrayOutput {
	return o
}

func (o ApplicationGatewayBackendAddressPoolArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayBackendAddressPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayBackendAddressPool {
		return vs[0].([]ApplicationGatewayBackendAddressPool)[vs[1].(int)]
	}).(ApplicationGatewayBackendAddressPoolOutput)
}

type ApplicationGatewayBackendAddressPoolResponse struct {
	BackendAddresses        []ApplicationGatewayBackendAddressResponse `pulumi:"backendAddresses"`
	BackendIPConfigurations []NetworkInterfaceIPConfigurationResponse  `pulumi:"backendIPConfigurations"`
	Etag                    *string                                    `pulumi:"etag"`
	Id                      *string                                    `pulumi:"id"`
	Name                    *string                                    `pulumi:"name"`
	ProvisioningState       *string                                    `pulumi:"provisioningState"`
}

type ApplicationGatewayBackendAddressPoolResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendAddressPoolResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendAddressPoolResponse)(nil)).Elem()
}

func (o ApplicationGatewayBackendAddressPoolResponseOutput) ToApplicationGatewayBackendAddressPoolResponseOutput() ApplicationGatewayBackendAddressPoolResponseOutput {
	return o
}

func (o ApplicationGatewayBackendAddressPoolResponseOutput) ToApplicationGatewayBackendAddressPoolResponseOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressPoolResponseOutput {
	return o
}

func (o ApplicationGatewayBackendAddressPoolResponseOutput) BackendAddresses() ApplicationGatewayBackendAddressResponseArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPoolResponse) []ApplicationGatewayBackendAddressResponse {
		return v.BackendAddresses
	}).(ApplicationGatewayBackendAddressResponseArrayOutput)
}

func (o ApplicationGatewayBackendAddressPoolResponseOutput) BackendIPConfigurations() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPoolResponse) []NetworkInterfaceIPConfigurationResponse {
		return v.BackendIPConfigurations
	}).(NetworkInterfaceIPConfigurationResponseArrayOutput)
}

func (o ApplicationGatewayBackendAddressPoolResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPoolResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendAddressPoolResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPoolResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendAddressPoolResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPoolResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendAddressPoolResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressPoolResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayBackendAddressPoolResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendAddressPoolResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendAddressPoolResponse)(nil)).Elem()
}

func (o ApplicationGatewayBackendAddressPoolResponseArrayOutput) ToApplicationGatewayBackendAddressPoolResponseArrayOutput() ApplicationGatewayBackendAddressPoolResponseArrayOutput {
	return o
}

func (o ApplicationGatewayBackendAddressPoolResponseArrayOutput) ToApplicationGatewayBackendAddressPoolResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressPoolResponseArrayOutput {
	return o
}

func (o ApplicationGatewayBackendAddressPoolResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayBackendAddressPoolResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayBackendAddressPoolResponse {
		return vs[0].([]ApplicationGatewayBackendAddressPoolResponse)[vs[1].(int)]
	}).(ApplicationGatewayBackendAddressPoolResponseOutput)
}

type ApplicationGatewayBackendAddressResponse struct {
	Fqdn      *string `pulumi:"fqdn"`
	IpAddress *string `pulumi:"ipAddress"`
}

type ApplicationGatewayBackendAddressResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendAddressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendAddressResponse)(nil)).Elem()
}

func (o ApplicationGatewayBackendAddressResponseOutput) ToApplicationGatewayBackendAddressResponseOutput() ApplicationGatewayBackendAddressResponseOutput {
	return o
}

func (o ApplicationGatewayBackendAddressResponseOutput) ToApplicationGatewayBackendAddressResponseOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressResponseOutput {
	return o
}

func (o ApplicationGatewayBackendAddressResponseOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressResponse) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendAddressResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendAddressResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayBackendAddressResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendAddressResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendAddressResponse)(nil)).Elem()
}

func (o ApplicationGatewayBackendAddressResponseArrayOutput) ToApplicationGatewayBackendAddressResponseArrayOutput() ApplicationGatewayBackendAddressResponseArrayOutput {
	return o
}

func (o ApplicationGatewayBackendAddressResponseArrayOutput) ToApplicationGatewayBackendAddressResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendAddressResponseArrayOutput {
	return o
}

func (o ApplicationGatewayBackendAddressResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayBackendAddressResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayBackendAddressResponse {
		return vs[0].([]ApplicationGatewayBackendAddressResponse)[vs[1].(int)]
	}).(ApplicationGatewayBackendAddressResponseOutput)
}

type ApplicationGatewayBackendHttpSettings struct {
	AuthenticationCertificates []SubResource                         `pulumi:"authenticationCertificates"`
	ConnectionDraining         *ApplicationGatewayConnectionDraining `pulumi:"connectionDraining"`
	CookieBasedAffinity        *string                               `pulumi:"cookieBasedAffinity"`
	Etag                       *string                               `pulumi:"etag"`
	Id                         *string                               `pulumi:"id"`
	Name                       *string                               `pulumi:"name"`
	Port                       *int                                  `pulumi:"port"`
	Probe                      *SubResource                          `pulumi:"probe"`
	Protocol                   *string                               `pulumi:"protocol"`
	ProvisioningState          *string                               `pulumi:"provisioningState"`
	RequestTimeout             *int                                  `pulumi:"requestTimeout"`
}





type ApplicationGatewayBackendHttpSettingsInput interface {
	pulumi.Input

	ToApplicationGatewayBackendHttpSettingsOutput() ApplicationGatewayBackendHttpSettingsOutput
	ToApplicationGatewayBackendHttpSettingsOutputWithContext(context.Context) ApplicationGatewayBackendHttpSettingsOutput
}

type ApplicationGatewayBackendHttpSettingsArgs struct {
	AuthenticationCertificates SubResourceArrayInput                        `pulumi:"authenticationCertificates"`
	ConnectionDraining         ApplicationGatewayConnectionDrainingPtrInput `pulumi:"connectionDraining"`
	CookieBasedAffinity        pulumi.StringPtrInput                        `pulumi:"cookieBasedAffinity"`
	Etag                       pulumi.StringPtrInput                        `pulumi:"etag"`
	Id                         pulumi.StringPtrInput                        `pulumi:"id"`
	Name                       pulumi.StringPtrInput                        `pulumi:"name"`
	Port                       pulumi.IntPtrInput                           `pulumi:"port"`
	Probe                      SubResourcePtrInput                          `pulumi:"probe"`
	Protocol                   pulumi.StringPtrInput                        `pulumi:"protocol"`
	ProvisioningState          pulumi.StringPtrInput                        `pulumi:"provisioningState"`
	RequestTimeout             pulumi.IntPtrInput                           `pulumi:"requestTimeout"`
}

func (ApplicationGatewayBackendHttpSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendHttpSettings)(nil)).Elem()
}

func (i ApplicationGatewayBackendHttpSettingsArgs) ToApplicationGatewayBackendHttpSettingsOutput() ApplicationGatewayBackendHttpSettingsOutput {
	return i.ToApplicationGatewayBackendHttpSettingsOutputWithContext(context.Background())
}

func (i ApplicationGatewayBackendHttpSettingsArgs) ToApplicationGatewayBackendHttpSettingsOutputWithContext(ctx context.Context) ApplicationGatewayBackendHttpSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayBackendHttpSettingsOutput)
}





type ApplicationGatewayBackendHttpSettingsArrayInput interface {
	pulumi.Input

	ToApplicationGatewayBackendHttpSettingsArrayOutput() ApplicationGatewayBackendHttpSettingsArrayOutput
	ToApplicationGatewayBackendHttpSettingsArrayOutputWithContext(context.Context) ApplicationGatewayBackendHttpSettingsArrayOutput
}

type ApplicationGatewayBackendHttpSettingsArray []ApplicationGatewayBackendHttpSettingsInput

func (ApplicationGatewayBackendHttpSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendHttpSettings)(nil)).Elem()
}

func (i ApplicationGatewayBackendHttpSettingsArray) ToApplicationGatewayBackendHttpSettingsArrayOutput() ApplicationGatewayBackendHttpSettingsArrayOutput {
	return i.ToApplicationGatewayBackendHttpSettingsArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayBackendHttpSettingsArray) ToApplicationGatewayBackendHttpSettingsArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendHttpSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayBackendHttpSettingsArrayOutput)
}

type ApplicationGatewayBackendHttpSettingsOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendHttpSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendHttpSettings)(nil)).Elem()
}

func (o ApplicationGatewayBackendHttpSettingsOutput) ToApplicationGatewayBackendHttpSettingsOutput() ApplicationGatewayBackendHttpSettingsOutput {
	return o
}

func (o ApplicationGatewayBackendHttpSettingsOutput) ToApplicationGatewayBackendHttpSettingsOutputWithContext(ctx context.Context) ApplicationGatewayBackendHttpSettingsOutput {
	return o
}

func (o ApplicationGatewayBackendHttpSettingsOutput) AuthenticationCertificates() SubResourceArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) []SubResource { return v.AuthenticationCertificates }).(SubResourceArrayOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) ConnectionDraining() ApplicationGatewayConnectionDrainingPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *ApplicationGatewayConnectionDraining {
		return v.ConnectionDraining
	}).(ApplicationGatewayConnectionDrainingPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) CookieBasedAffinity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.CookieBasedAffinity }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Probe() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *SubResource { return v.Probe }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsOutput) RequestTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettings) *int { return v.RequestTimeout }).(pulumi.IntPtrOutput)
}

type ApplicationGatewayBackendHttpSettingsArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendHttpSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendHttpSettings)(nil)).Elem()
}

func (o ApplicationGatewayBackendHttpSettingsArrayOutput) ToApplicationGatewayBackendHttpSettingsArrayOutput() ApplicationGatewayBackendHttpSettingsArrayOutput {
	return o
}

func (o ApplicationGatewayBackendHttpSettingsArrayOutput) ToApplicationGatewayBackendHttpSettingsArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendHttpSettingsArrayOutput {
	return o
}

func (o ApplicationGatewayBackendHttpSettingsArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayBackendHttpSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayBackendHttpSettings {
		return vs[0].([]ApplicationGatewayBackendHttpSettings)[vs[1].(int)]
	}).(ApplicationGatewayBackendHttpSettingsOutput)
}

type ApplicationGatewayBackendHttpSettingsResponse struct {
	AuthenticationCertificates []SubResourceResponse                         `pulumi:"authenticationCertificates"`
	ConnectionDraining         *ApplicationGatewayConnectionDrainingResponse `pulumi:"connectionDraining"`
	CookieBasedAffinity        *string                                       `pulumi:"cookieBasedAffinity"`
	Etag                       *string                                       `pulumi:"etag"`
	Id                         *string                                       `pulumi:"id"`
	Name                       *string                                       `pulumi:"name"`
	Port                       *int                                          `pulumi:"port"`
	Probe                      *SubResourceResponse                          `pulumi:"probe"`
	Protocol                   *string                                       `pulumi:"protocol"`
	ProvisioningState          *string                                       `pulumi:"provisioningState"`
	RequestTimeout             *int                                          `pulumi:"requestTimeout"`
}

type ApplicationGatewayBackendHttpSettingsResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendHttpSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayBackendHttpSettingsResponse)(nil)).Elem()
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) ToApplicationGatewayBackendHttpSettingsResponseOutput() ApplicationGatewayBackendHttpSettingsResponseOutput {
	return o
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) ToApplicationGatewayBackendHttpSettingsResponseOutputWithContext(ctx context.Context) ApplicationGatewayBackendHttpSettingsResponseOutput {
	return o
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) AuthenticationCertificates() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) []SubResourceResponse {
		return v.AuthenticationCertificates
	}).(SubResourceResponseArrayOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) ConnectionDraining() ApplicationGatewayConnectionDrainingResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *ApplicationGatewayConnectionDrainingResponse {
		return v.ConnectionDraining
	}).(ApplicationGatewayConnectionDrainingResponsePtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) CookieBasedAffinity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.CookieBasedAffinity }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Probe() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *SubResourceResponse { return v.Probe }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayBackendHttpSettingsResponseOutput) RequestTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayBackendHttpSettingsResponse) *int { return v.RequestTimeout }).(pulumi.IntPtrOutput)
}

type ApplicationGatewayBackendHttpSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayBackendHttpSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayBackendHttpSettingsResponse)(nil)).Elem()
}

func (o ApplicationGatewayBackendHttpSettingsResponseArrayOutput) ToApplicationGatewayBackendHttpSettingsResponseArrayOutput() ApplicationGatewayBackendHttpSettingsResponseArrayOutput {
	return o
}

func (o ApplicationGatewayBackendHttpSettingsResponseArrayOutput) ToApplicationGatewayBackendHttpSettingsResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayBackendHttpSettingsResponseArrayOutput {
	return o
}

func (o ApplicationGatewayBackendHttpSettingsResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayBackendHttpSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayBackendHttpSettingsResponse {
		return vs[0].([]ApplicationGatewayBackendHttpSettingsResponse)[vs[1].(int)]
	}).(ApplicationGatewayBackendHttpSettingsResponseOutput)
}

type ApplicationGatewayConnectionDraining struct {
	DrainTimeoutInSec int  `pulumi:"drainTimeoutInSec"`
	Enabled           bool `pulumi:"enabled"`
}





type ApplicationGatewayConnectionDrainingInput interface {
	pulumi.Input

	ToApplicationGatewayConnectionDrainingOutput() ApplicationGatewayConnectionDrainingOutput
	ToApplicationGatewayConnectionDrainingOutputWithContext(context.Context) ApplicationGatewayConnectionDrainingOutput
}

type ApplicationGatewayConnectionDrainingArgs struct {
	DrainTimeoutInSec pulumi.IntInput  `pulumi:"drainTimeoutInSec"`
	Enabled           pulumi.BoolInput `pulumi:"enabled"`
}

func (ApplicationGatewayConnectionDrainingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayConnectionDraining)(nil)).Elem()
}

func (i ApplicationGatewayConnectionDrainingArgs) ToApplicationGatewayConnectionDrainingOutput() ApplicationGatewayConnectionDrainingOutput {
	return i.ToApplicationGatewayConnectionDrainingOutputWithContext(context.Background())
}

func (i ApplicationGatewayConnectionDrainingArgs) ToApplicationGatewayConnectionDrainingOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayConnectionDrainingOutput)
}

func (i ApplicationGatewayConnectionDrainingArgs) ToApplicationGatewayConnectionDrainingPtrOutput() ApplicationGatewayConnectionDrainingPtrOutput {
	return i.ToApplicationGatewayConnectionDrainingPtrOutputWithContext(context.Background())
}

func (i ApplicationGatewayConnectionDrainingArgs) ToApplicationGatewayConnectionDrainingPtrOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayConnectionDrainingOutput).ToApplicationGatewayConnectionDrainingPtrOutputWithContext(ctx)
}









type ApplicationGatewayConnectionDrainingPtrInput interface {
	pulumi.Input

	ToApplicationGatewayConnectionDrainingPtrOutput() ApplicationGatewayConnectionDrainingPtrOutput
	ToApplicationGatewayConnectionDrainingPtrOutputWithContext(context.Context) ApplicationGatewayConnectionDrainingPtrOutput
}

type applicationGatewayConnectionDrainingPtrType ApplicationGatewayConnectionDrainingArgs

func ApplicationGatewayConnectionDrainingPtr(v *ApplicationGatewayConnectionDrainingArgs) ApplicationGatewayConnectionDrainingPtrInput {
	return (*applicationGatewayConnectionDrainingPtrType)(v)
}

func (*applicationGatewayConnectionDrainingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayConnectionDraining)(nil)).Elem()
}

func (i *applicationGatewayConnectionDrainingPtrType) ToApplicationGatewayConnectionDrainingPtrOutput() ApplicationGatewayConnectionDrainingPtrOutput {
	return i.ToApplicationGatewayConnectionDrainingPtrOutputWithContext(context.Background())
}

func (i *applicationGatewayConnectionDrainingPtrType) ToApplicationGatewayConnectionDrainingPtrOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayConnectionDrainingPtrOutput)
}

type ApplicationGatewayConnectionDrainingOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayConnectionDrainingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayConnectionDraining)(nil)).Elem()
}

func (o ApplicationGatewayConnectionDrainingOutput) ToApplicationGatewayConnectionDrainingOutput() ApplicationGatewayConnectionDrainingOutput {
	return o
}

func (o ApplicationGatewayConnectionDrainingOutput) ToApplicationGatewayConnectionDrainingOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingOutput {
	return o
}

func (o ApplicationGatewayConnectionDrainingOutput) ToApplicationGatewayConnectionDrainingPtrOutput() ApplicationGatewayConnectionDrainingPtrOutput {
	return o.ToApplicationGatewayConnectionDrainingPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayConnectionDrainingOutput) ToApplicationGatewayConnectionDrainingPtrOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewayConnectionDraining) *ApplicationGatewayConnectionDraining {
		return &v
	}).(ApplicationGatewayConnectionDrainingPtrOutput)
}

func (o ApplicationGatewayConnectionDrainingOutput) DrainTimeoutInSec() pulumi.IntOutput {
	return o.ApplyT(func(v ApplicationGatewayConnectionDraining) int { return v.DrainTimeoutInSec }).(pulumi.IntOutput)
}

func (o ApplicationGatewayConnectionDrainingOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ApplicationGatewayConnectionDraining) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type ApplicationGatewayConnectionDrainingPtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayConnectionDrainingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayConnectionDraining)(nil)).Elem()
}

func (o ApplicationGatewayConnectionDrainingPtrOutput) ToApplicationGatewayConnectionDrainingPtrOutput() ApplicationGatewayConnectionDrainingPtrOutput {
	return o
}

func (o ApplicationGatewayConnectionDrainingPtrOutput) ToApplicationGatewayConnectionDrainingPtrOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingPtrOutput {
	return o
}

func (o ApplicationGatewayConnectionDrainingPtrOutput) Elem() ApplicationGatewayConnectionDrainingOutput {
	return o.ApplyT(func(v *ApplicationGatewayConnectionDraining) ApplicationGatewayConnectionDraining {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewayConnectionDraining
		return ret
	}).(ApplicationGatewayConnectionDrainingOutput)
}

func (o ApplicationGatewayConnectionDrainingPtrOutput) DrainTimeoutInSec() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayConnectionDraining) *int {
		if v == nil {
			return nil
		}
		return &v.DrainTimeoutInSec
	}).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayConnectionDrainingPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayConnectionDraining) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type ApplicationGatewayConnectionDrainingResponse struct {
	DrainTimeoutInSec int  `pulumi:"drainTimeoutInSec"`
	Enabled           bool `pulumi:"enabled"`
}

type ApplicationGatewayConnectionDrainingResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayConnectionDrainingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayConnectionDrainingResponse)(nil)).Elem()
}

func (o ApplicationGatewayConnectionDrainingResponseOutput) ToApplicationGatewayConnectionDrainingResponseOutput() ApplicationGatewayConnectionDrainingResponseOutput {
	return o
}

func (o ApplicationGatewayConnectionDrainingResponseOutput) ToApplicationGatewayConnectionDrainingResponseOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingResponseOutput {
	return o
}

func (o ApplicationGatewayConnectionDrainingResponseOutput) DrainTimeoutInSec() pulumi.IntOutput {
	return o.ApplyT(func(v ApplicationGatewayConnectionDrainingResponse) int { return v.DrainTimeoutInSec }).(pulumi.IntOutput)
}

func (o ApplicationGatewayConnectionDrainingResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ApplicationGatewayConnectionDrainingResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type ApplicationGatewayConnectionDrainingResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayConnectionDrainingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayConnectionDrainingResponse)(nil)).Elem()
}

func (o ApplicationGatewayConnectionDrainingResponsePtrOutput) ToApplicationGatewayConnectionDrainingResponsePtrOutput() ApplicationGatewayConnectionDrainingResponsePtrOutput {
	return o
}

func (o ApplicationGatewayConnectionDrainingResponsePtrOutput) ToApplicationGatewayConnectionDrainingResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewayConnectionDrainingResponsePtrOutput {
	return o
}

func (o ApplicationGatewayConnectionDrainingResponsePtrOutput) Elem() ApplicationGatewayConnectionDrainingResponseOutput {
	return o.ApplyT(func(v *ApplicationGatewayConnectionDrainingResponse) ApplicationGatewayConnectionDrainingResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewayConnectionDrainingResponse
		return ret
	}).(ApplicationGatewayConnectionDrainingResponseOutput)
}

func (o ApplicationGatewayConnectionDrainingResponsePtrOutput) DrainTimeoutInSec() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayConnectionDrainingResponse) *int {
		if v == nil {
			return nil
		}
		return &v.DrainTimeoutInSec
	}).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayConnectionDrainingResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayConnectionDrainingResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type ApplicationGatewayFirewallDisabledRuleGroup struct {
	RuleGroupName string `pulumi:"ruleGroupName"`
	Rules         []int  `pulumi:"rules"`
}





type ApplicationGatewayFirewallDisabledRuleGroupInput interface {
	pulumi.Input

	ToApplicationGatewayFirewallDisabledRuleGroupOutput() ApplicationGatewayFirewallDisabledRuleGroupOutput
	ToApplicationGatewayFirewallDisabledRuleGroupOutputWithContext(context.Context) ApplicationGatewayFirewallDisabledRuleGroupOutput
}

type ApplicationGatewayFirewallDisabledRuleGroupArgs struct {
	RuleGroupName pulumi.StringInput   `pulumi:"ruleGroupName"`
	Rules         pulumi.IntArrayInput `pulumi:"rules"`
}

func (ApplicationGatewayFirewallDisabledRuleGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFirewallDisabledRuleGroup)(nil)).Elem()
}

func (i ApplicationGatewayFirewallDisabledRuleGroupArgs) ToApplicationGatewayFirewallDisabledRuleGroupOutput() ApplicationGatewayFirewallDisabledRuleGroupOutput {
	return i.ToApplicationGatewayFirewallDisabledRuleGroupOutputWithContext(context.Background())
}

func (i ApplicationGatewayFirewallDisabledRuleGroupArgs) ToApplicationGatewayFirewallDisabledRuleGroupOutputWithContext(ctx context.Context) ApplicationGatewayFirewallDisabledRuleGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayFirewallDisabledRuleGroupOutput)
}





type ApplicationGatewayFirewallDisabledRuleGroupArrayInput interface {
	pulumi.Input

	ToApplicationGatewayFirewallDisabledRuleGroupArrayOutput() ApplicationGatewayFirewallDisabledRuleGroupArrayOutput
	ToApplicationGatewayFirewallDisabledRuleGroupArrayOutputWithContext(context.Context) ApplicationGatewayFirewallDisabledRuleGroupArrayOutput
}

type ApplicationGatewayFirewallDisabledRuleGroupArray []ApplicationGatewayFirewallDisabledRuleGroupInput

func (ApplicationGatewayFirewallDisabledRuleGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFirewallDisabledRuleGroup)(nil)).Elem()
}

func (i ApplicationGatewayFirewallDisabledRuleGroupArray) ToApplicationGatewayFirewallDisabledRuleGroupArrayOutput() ApplicationGatewayFirewallDisabledRuleGroupArrayOutput {
	return i.ToApplicationGatewayFirewallDisabledRuleGroupArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayFirewallDisabledRuleGroupArray) ToApplicationGatewayFirewallDisabledRuleGroupArrayOutputWithContext(ctx context.Context) ApplicationGatewayFirewallDisabledRuleGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayFirewallDisabledRuleGroupArrayOutput)
}

type ApplicationGatewayFirewallDisabledRuleGroupOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFirewallDisabledRuleGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFirewallDisabledRuleGroup)(nil)).Elem()
}

func (o ApplicationGatewayFirewallDisabledRuleGroupOutput) ToApplicationGatewayFirewallDisabledRuleGroupOutput() ApplicationGatewayFirewallDisabledRuleGroupOutput {
	return o
}

func (o ApplicationGatewayFirewallDisabledRuleGroupOutput) ToApplicationGatewayFirewallDisabledRuleGroupOutputWithContext(ctx context.Context) ApplicationGatewayFirewallDisabledRuleGroupOutput {
	return o
}

func (o ApplicationGatewayFirewallDisabledRuleGroupOutput) RuleGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationGatewayFirewallDisabledRuleGroup) string { return v.RuleGroupName }).(pulumi.StringOutput)
}

func (o ApplicationGatewayFirewallDisabledRuleGroupOutput) Rules() pulumi.IntArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayFirewallDisabledRuleGroup) []int { return v.Rules }).(pulumi.IntArrayOutput)
}

type ApplicationGatewayFirewallDisabledRuleGroupArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFirewallDisabledRuleGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFirewallDisabledRuleGroup)(nil)).Elem()
}

func (o ApplicationGatewayFirewallDisabledRuleGroupArrayOutput) ToApplicationGatewayFirewallDisabledRuleGroupArrayOutput() ApplicationGatewayFirewallDisabledRuleGroupArrayOutput {
	return o
}

func (o ApplicationGatewayFirewallDisabledRuleGroupArrayOutput) ToApplicationGatewayFirewallDisabledRuleGroupArrayOutputWithContext(ctx context.Context) ApplicationGatewayFirewallDisabledRuleGroupArrayOutput {
	return o
}

func (o ApplicationGatewayFirewallDisabledRuleGroupArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayFirewallDisabledRuleGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayFirewallDisabledRuleGroup {
		return vs[0].([]ApplicationGatewayFirewallDisabledRuleGroup)[vs[1].(int)]
	}).(ApplicationGatewayFirewallDisabledRuleGroupOutput)
}

type ApplicationGatewayFirewallDisabledRuleGroupResponse struct {
	RuleGroupName string `pulumi:"ruleGroupName"`
	Rules         []int  `pulumi:"rules"`
}

type ApplicationGatewayFirewallDisabledRuleGroupResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFirewallDisabledRuleGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFirewallDisabledRuleGroupResponse)(nil)).Elem()
}

func (o ApplicationGatewayFirewallDisabledRuleGroupResponseOutput) ToApplicationGatewayFirewallDisabledRuleGroupResponseOutput() ApplicationGatewayFirewallDisabledRuleGroupResponseOutput {
	return o
}

func (o ApplicationGatewayFirewallDisabledRuleGroupResponseOutput) ToApplicationGatewayFirewallDisabledRuleGroupResponseOutputWithContext(ctx context.Context) ApplicationGatewayFirewallDisabledRuleGroupResponseOutput {
	return o
}

func (o ApplicationGatewayFirewallDisabledRuleGroupResponseOutput) RuleGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationGatewayFirewallDisabledRuleGroupResponse) string { return v.RuleGroupName }).(pulumi.StringOutput)
}

func (o ApplicationGatewayFirewallDisabledRuleGroupResponseOutput) Rules() pulumi.IntArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayFirewallDisabledRuleGroupResponse) []int { return v.Rules }).(pulumi.IntArrayOutput)
}

type ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFirewallDisabledRuleGroupResponse)(nil)).Elem()
}

func (o ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput) ToApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput() ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput {
	return o
}

func (o ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput) ToApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput {
	return o
}

func (o ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayFirewallDisabledRuleGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayFirewallDisabledRuleGroupResponse {
		return vs[0].([]ApplicationGatewayFirewallDisabledRuleGroupResponse)[vs[1].(int)]
	}).(ApplicationGatewayFirewallDisabledRuleGroupResponseOutput)
}

type ApplicationGatewayFrontendIPConfiguration struct {
	Etag                      *string      `pulumi:"etag"`
	Id                        *string      `pulumi:"id"`
	Name                      *string      `pulumi:"name"`
	PrivateIPAddress          *string      `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string      `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         *string      `pulumi:"provisioningState"`
	PublicIPAddress           *SubResource `pulumi:"publicIPAddress"`
	Subnet                    *SubResource `pulumi:"subnet"`
}





type ApplicationGatewayFrontendIPConfigurationInput interface {
	pulumi.Input

	ToApplicationGatewayFrontendIPConfigurationOutput() ApplicationGatewayFrontendIPConfigurationOutput
	ToApplicationGatewayFrontendIPConfigurationOutputWithContext(context.Context) ApplicationGatewayFrontendIPConfigurationOutput
}

type ApplicationGatewayFrontendIPConfigurationArgs struct {
	Etag                      pulumi.StringPtrInput `pulumi:"etag"`
	Id                        pulumi.StringPtrInput `pulumi:"id"`
	Name                      pulumi.StringPtrInput `pulumi:"name"`
	PrivateIPAddress          pulumi.StringPtrInput `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod pulumi.StringPtrInput `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         pulumi.StringPtrInput `pulumi:"provisioningState"`
	PublicIPAddress           SubResourcePtrInput   `pulumi:"publicIPAddress"`
	Subnet                    SubResourcePtrInput   `pulumi:"subnet"`
}

func (ApplicationGatewayFrontendIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFrontendIPConfiguration)(nil)).Elem()
}

func (i ApplicationGatewayFrontendIPConfigurationArgs) ToApplicationGatewayFrontendIPConfigurationOutput() ApplicationGatewayFrontendIPConfigurationOutput {
	return i.ToApplicationGatewayFrontendIPConfigurationOutputWithContext(context.Background())
}

func (i ApplicationGatewayFrontendIPConfigurationArgs) ToApplicationGatewayFrontendIPConfigurationOutputWithContext(ctx context.Context) ApplicationGatewayFrontendIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayFrontendIPConfigurationOutput)
}





type ApplicationGatewayFrontendIPConfigurationArrayInput interface {
	pulumi.Input

	ToApplicationGatewayFrontendIPConfigurationArrayOutput() ApplicationGatewayFrontendIPConfigurationArrayOutput
	ToApplicationGatewayFrontendIPConfigurationArrayOutputWithContext(context.Context) ApplicationGatewayFrontendIPConfigurationArrayOutput
}

type ApplicationGatewayFrontendIPConfigurationArray []ApplicationGatewayFrontendIPConfigurationInput

func (ApplicationGatewayFrontendIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFrontendIPConfiguration)(nil)).Elem()
}

func (i ApplicationGatewayFrontendIPConfigurationArray) ToApplicationGatewayFrontendIPConfigurationArrayOutput() ApplicationGatewayFrontendIPConfigurationArrayOutput {
	return i.ToApplicationGatewayFrontendIPConfigurationArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayFrontendIPConfigurationArray) ToApplicationGatewayFrontendIPConfigurationArrayOutputWithContext(ctx context.Context) ApplicationGatewayFrontendIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayFrontendIPConfigurationArrayOutput)
}

type ApplicationGatewayFrontendIPConfigurationOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFrontendIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFrontendIPConfiguration)(nil)).Elem()
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) ToApplicationGatewayFrontendIPConfigurationOutput() ApplicationGatewayFrontendIPConfigurationOutput {
	return o
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) ToApplicationGatewayFrontendIPConfigurationOutputWithContext(ctx context.Context) ApplicationGatewayFrontendIPConfigurationOutput {
	return o
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfiguration) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfiguration) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfiguration) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfiguration) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) PublicIPAddress() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfiguration) *SubResource { return v.PublicIPAddress }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationOutput) Subnet() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfiguration) *SubResource { return v.Subnet }).(SubResourcePtrOutput)
}

type ApplicationGatewayFrontendIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFrontendIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFrontendIPConfiguration)(nil)).Elem()
}

func (o ApplicationGatewayFrontendIPConfigurationArrayOutput) ToApplicationGatewayFrontendIPConfigurationArrayOutput() ApplicationGatewayFrontendIPConfigurationArrayOutput {
	return o
}

func (o ApplicationGatewayFrontendIPConfigurationArrayOutput) ToApplicationGatewayFrontendIPConfigurationArrayOutputWithContext(ctx context.Context) ApplicationGatewayFrontendIPConfigurationArrayOutput {
	return o
}

func (o ApplicationGatewayFrontendIPConfigurationArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayFrontendIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayFrontendIPConfiguration {
		return vs[0].([]ApplicationGatewayFrontendIPConfiguration)[vs[1].(int)]
	}).(ApplicationGatewayFrontendIPConfigurationOutput)
}

type ApplicationGatewayFrontendIPConfigurationResponse struct {
	Etag                      *string              `pulumi:"etag"`
	Id                        *string              `pulumi:"id"`
	Name                      *string              `pulumi:"name"`
	PrivateIPAddress          *string              `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string              `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         *string              `pulumi:"provisioningState"`
	PublicIPAddress           *SubResourceResponse `pulumi:"publicIPAddress"`
	Subnet                    *SubResourceResponse `pulumi:"subnet"`
}

type ApplicationGatewayFrontendIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFrontendIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFrontendIPConfigurationResponse)(nil)).Elem()
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) ToApplicationGatewayFrontendIPConfigurationResponseOutput() ApplicationGatewayFrontendIPConfigurationResponseOutput {
	return o
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) ToApplicationGatewayFrontendIPConfigurationResponseOutputWithContext(ctx context.Context) ApplicationGatewayFrontendIPConfigurationResponseOutput {
	return o
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfigurationResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfigurationResponse) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfigurationResponse) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfigurationResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) PublicIPAddress() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfigurationResponse) *SubResourceResponse {
		return v.PublicIPAddress
	}).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayFrontendIPConfigurationResponseOutput) Subnet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendIPConfigurationResponse) *SubResourceResponse { return v.Subnet }).(SubResourceResponsePtrOutput)
}

type ApplicationGatewayFrontendIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFrontendIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFrontendIPConfigurationResponse)(nil)).Elem()
}

func (o ApplicationGatewayFrontendIPConfigurationResponseArrayOutput) ToApplicationGatewayFrontendIPConfigurationResponseArrayOutput() ApplicationGatewayFrontendIPConfigurationResponseArrayOutput {
	return o
}

func (o ApplicationGatewayFrontendIPConfigurationResponseArrayOutput) ToApplicationGatewayFrontendIPConfigurationResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayFrontendIPConfigurationResponseArrayOutput {
	return o
}

func (o ApplicationGatewayFrontendIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayFrontendIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayFrontendIPConfigurationResponse {
		return vs[0].([]ApplicationGatewayFrontendIPConfigurationResponse)[vs[1].(int)]
	}).(ApplicationGatewayFrontendIPConfigurationResponseOutput)
}

type ApplicationGatewayFrontendPort struct {
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	Port              *int    `pulumi:"port"`
	ProvisioningState *string `pulumi:"provisioningState"`
}





type ApplicationGatewayFrontendPortInput interface {
	pulumi.Input

	ToApplicationGatewayFrontendPortOutput() ApplicationGatewayFrontendPortOutput
	ToApplicationGatewayFrontendPortOutputWithContext(context.Context) ApplicationGatewayFrontendPortOutput
}

type ApplicationGatewayFrontendPortArgs struct {
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	Port              pulumi.IntPtrInput    `pulumi:"port"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (ApplicationGatewayFrontendPortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFrontendPort)(nil)).Elem()
}

func (i ApplicationGatewayFrontendPortArgs) ToApplicationGatewayFrontendPortOutput() ApplicationGatewayFrontendPortOutput {
	return i.ToApplicationGatewayFrontendPortOutputWithContext(context.Background())
}

func (i ApplicationGatewayFrontendPortArgs) ToApplicationGatewayFrontendPortOutputWithContext(ctx context.Context) ApplicationGatewayFrontendPortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayFrontendPortOutput)
}





type ApplicationGatewayFrontendPortArrayInput interface {
	pulumi.Input

	ToApplicationGatewayFrontendPortArrayOutput() ApplicationGatewayFrontendPortArrayOutput
	ToApplicationGatewayFrontendPortArrayOutputWithContext(context.Context) ApplicationGatewayFrontendPortArrayOutput
}

type ApplicationGatewayFrontendPortArray []ApplicationGatewayFrontendPortInput

func (ApplicationGatewayFrontendPortArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFrontendPort)(nil)).Elem()
}

func (i ApplicationGatewayFrontendPortArray) ToApplicationGatewayFrontendPortArrayOutput() ApplicationGatewayFrontendPortArrayOutput {
	return i.ToApplicationGatewayFrontendPortArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayFrontendPortArray) ToApplicationGatewayFrontendPortArrayOutputWithContext(ctx context.Context) ApplicationGatewayFrontendPortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayFrontendPortArrayOutput)
}

type ApplicationGatewayFrontendPortOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFrontendPortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFrontendPort)(nil)).Elem()
}

func (o ApplicationGatewayFrontendPortOutput) ToApplicationGatewayFrontendPortOutput() ApplicationGatewayFrontendPortOutput {
	return o
}

func (o ApplicationGatewayFrontendPortOutput) ToApplicationGatewayFrontendPortOutputWithContext(ctx context.Context) ApplicationGatewayFrontendPortOutput {
	return o
}

func (o ApplicationGatewayFrontendPortOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPort) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendPortOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPort) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendPortOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPort) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendPortOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPort) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayFrontendPortOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPort) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayFrontendPortArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFrontendPortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFrontendPort)(nil)).Elem()
}

func (o ApplicationGatewayFrontendPortArrayOutput) ToApplicationGatewayFrontendPortArrayOutput() ApplicationGatewayFrontendPortArrayOutput {
	return o
}

func (o ApplicationGatewayFrontendPortArrayOutput) ToApplicationGatewayFrontendPortArrayOutputWithContext(ctx context.Context) ApplicationGatewayFrontendPortArrayOutput {
	return o
}

func (o ApplicationGatewayFrontendPortArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayFrontendPortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayFrontendPort {
		return vs[0].([]ApplicationGatewayFrontendPort)[vs[1].(int)]
	}).(ApplicationGatewayFrontendPortOutput)
}

type ApplicationGatewayFrontendPortResponse struct {
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	Port              *int    `pulumi:"port"`
	ProvisioningState *string `pulumi:"provisioningState"`
}

type ApplicationGatewayFrontendPortResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFrontendPortResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayFrontendPortResponse)(nil)).Elem()
}

func (o ApplicationGatewayFrontendPortResponseOutput) ToApplicationGatewayFrontendPortResponseOutput() ApplicationGatewayFrontendPortResponseOutput {
	return o
}

func (o ApplicationGatewayFrontendPortResponseOutput) ToApplicationGatewayFrontendPortResponseOutputWithContext(ctx context.Context) ApplicationGatewayFrontendPortResponseOutput {
	return o
}

func (o ApplicationGatewayFrontendPortResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPortResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendPortResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPortResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendPortResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPortResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayFrontendPortResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPortResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayFrontendPortResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayFrontendPortResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayFrontendPortResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayFrontendPortResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayFrontendPortResponse)(nil)).Elem()
}

func (o ApplicationGatewayFrontendPortResponseArrayOutput) ToApplicationGatewayFrontendPortResponseArrayOutput() ApplicationGatewayFrontendPortResponseArrayOutput {
	return o
}

func (o ApplicationGatewayFrontendPortResponseArrayOutput) ToApplicationGatewayFrontendPortResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayFrontendPortResponseArrayOutput {
	return o
}

func (o ApplicationGatewayFrontendPortResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayFrontendPortResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayFrontendPortResponse {
		return vs[0].([]ApplicationGatewayFrontendPortResponse)[vs[1].(int)]
	}).(ApplicationGatewayFrontendPortResponseOutput)
}

type ApplicationGatewayHttpListener struct {
	Etag                        *string      `pulumi:"etag"`
	FrontendIPConfiguration     *SubResource `pulumi:"frontendIPConfiguration"`
	FrontendPort                *SubResource `pulumi:"frontendPort"`
	HostName                    *string      `pulumi:"hostName"`
	Id                          *string      `pulumi:"id"`
	Name                        *string      `pulumi:"name"`
	Protocol                    *string      `pulumi:"protocol"`
	ProvisioningState           *string      `pulumi:"provisioningState"`
	RequireServerNameIndication *bool        `pulumi:"requireServerNameIndication"`
	SslCertificate              *SubResource `pulumi:"sslCertificate"`
}





type ApplicationGatewayHttpListenerInput interface {
	pulumi.Input

	ToApplicationGatewayHttpListenerOutput() ApplicationGatewayHttpListenerOutput
	ToApplicationGatewayHttpListenerOutputWithContext(context.Context) ApplicationGatewayHttpListenerOutput
}

type ApplicationGatewayHttpListenerArgs struct {
	Etag                        pulumi.StringPtrInput `pulumi:"etag"`
	FrontendIPConfiguration     SubResourcePtrInput   `pulumi:"frontendIPConfiguration"`
	FrontendPort                SubResourcePtrInput   `pulumi:"frontendPort"`
	HostName                    pulumi.StringPtrInput `pulumi:"hostName"`
	Id                          pulumi.StringPtrInput `pulumi:"id"`
	Name                        pulumi.StringPtrInput `pulumi:"name"`
	Protocol                    pulumi.StringPtrInput `pulumi:"protocol"`
	ProvisioningState           pulumi.StringPtrInput `pulumi:"provisioningState"`
	RequireServerNameIndication pulumi.BoolPtrInput   `pulumi:"requireServerNameIndication"`
	SslCertificate              SubResourcePtrInput   `pulumi:"sslCertificate"`
}

func (ApplicationGatewayHttpListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayHttpListener)(nil)).Elem()
}

func (i ApplicationGatewayHttpListenerArgs) ToApplicationGatewayHttpListenerOutput() ApplicationGatewayHttpListenerOutput {
	return i.ToApplicationGatewayHttpListenerOutputWithContext(context.Background())
}

func (i ApplicationGatewayHttpListenerArgs) ToApplicationGatewayHttpListenerOutputWithContext(ctx context.Context) ApplicationGatewayHttpListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayHttpListenerOutput)
}





type ApplicationGatewayHttpListenerArrayInput interface {
	pulumi.Input

	ToApplicationGatewayHttpListenerArrayOutput() ApplicationGatewayHttpListenerArrayOutput
	ToApplicationGatewayHttpListenerArrayOutputWithContext(context.Context) ApplicationGatewayHttpListenerArrayOutput
}

type ApplicationGatewayHttpListenerArray []ApplicationGatewayHttpListenerInput

func (ApplicationGatewayHttpListenerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayHttpListener)(nil)).Elem()
}

func (i ApplicationGatewayHttpListenerArray) ToApplicationGatewayHttpListenerArrayOutput() ApplicationGatewayHttpListenerArrayOutput {
	return i.ToApplicationGatewayHttpListenerArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayHttpListenerArray) ToApplicationGatewayHttpListenerArrayOutputWithContext(ctx context.Context) ApplicationGatewayHttpListenerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayHttpListenerArrayOutput)
}

type ApplicationGatewayHttpListenerOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayHttpListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayHttpListener)(nil)).Elem()
}

func (o ApplicationGatewayHttpListenerOutput) ToApplicationGatewayHttpListenerOutput() ApplicationGatewayHttpListenerOutput {
	return o
}

func (o ApplicationGatewayHttpListenerOutput) ToApplicationGatewayHttpListenerOutputWithContext(ctx context.Context) ApplicationGatewayHttpListenerOutput {
	return o
}

func (o ApplicationGatewayHttpListenerOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerOutput) FrontendIPConfiguration() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *SubResource { return v.FrontendIPConfiguration }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayHttpListenerOutput) FrontendPort() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *SubResource { return v.FrontendPort }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayHttpListenerOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *string { return v.HostName }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerOutput) RequireServerNameIndication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *bool { return v.RequireServerNameIndication }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayHttpListenerOutput) SslCertificate() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListener) *SubResource { return v.SslCertificate }).(SubResourcePtrOutput)
}

type ApplicationGatewayHttpListenerArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayHttpListenerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayHttpListener)(nil)).Elem()
}

func (o ApplicationGatewayHttpListenerArrayOutput) ToApplicationGatewayHttpListenerArrayOutput() ApplicationGatewayHttpListenerArrayOutput {
	return o
}

func (o ApplicationGatewayHttpListenerArrayOutput) ToApplicationGatewayHttpListenerArrayOutputWithContext(ctx context.Context) ApplicationGatewayHttpListenerArrayOutput {
	return o
}

func (o ApplicationGatewayHttpListenerArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayHttpListenerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayHttpListener {
		return vs[0].([]ApplicationGatewayHttpListener)[vs[1].(int)]
	}).(ApplicationGatewayHttpListenerOutput)
}

type ApplicationGatewayHttpListenerResponse struct {
	Etag                        *string              `pulumi:"etag"`
	FrontendIPConfiguration     *SubResourceResponse `pulumi:"frontendIPConfiguration"`
	FrontendPort                *SubResourceResponse `pulumi:"frontendPort"`
	HostName                    *string              `pulumi:"hostName"`
	Id                          *string              `pulumi:"id"`
	Name                        *string              `pulumi:"name"`
	Protocol                    *string              `pulumi:"protocol"`
	ProvisioningState           *string              `pulumi:"provisioningState"`
	RequireServerNameIndication *bool                `pulumi:"requireServerNameIndication"`
	SslCertificate              *SubResourceResponse `pulumi:"sslCertificate"`
}

type ApplicationGatewayHttpListenerResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayHttpListenerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayHttpListenerResponse)(nil)).Elem()
}

func (o ApplicationGatewayHttpListenerResponseOutput) ToApplicationGatewayHttpListenerResponseOutput() ApplicationGatewayHttpListenerResponseOutput {
	return o
}

func (o ApplicationGatewayHttpListenerResponseOutput) ToApplicationGatewayHttpListenerResponseOutputWithContext(ctx context.Context) ApplicationGatewayHttpListenerResponseOutput {
	return o
}

func (o ApplicationGatewayHttpListenerResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerResponseOutput) FrontendIPConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *SubResourceResponse { return v.FrontendIPConfiguration }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayHttpListenerResponseOutput) FrontendPort() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *SubResourceResponse { return v.FrontendPort }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayHttpListenerResponseOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *string { return v.HostName }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayHttpListenerResponseOutput) RequireServerNameIndication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *bool { return v.RequireServerNameIndication }).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayHttpListenerResponseOutput) SslCertificate() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayHttpListenerResponse) *SubResourceResponse { return v.SslCertificate }).(SubResourceResponsePtrOutput)
}

type ApplicationGatewayHttpListenerResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayHttpListenerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayHttpListenerResponse)(nil)).Elem()
}

func (o ApplicationGatewayHttpListenerResponseArrayOutput) ToApplicationGatewayHttpListenerResponseArrayOutput() ApplicationGatewayHttpListenerResponseArrayOutput {
	return o
}

func (o ApplicationGatewayHttpListenerResponseArrayOutput) ToApplicationGatewayHttpListenerResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayHttpListenerResponseArrayOutput {
	return o
}

func (o ApplicationGatewayHttpListenerResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayHttpListenerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayHttpListenerResponse {
		return vs[0].([]ApplicationGatewayHttpListenerResponse)[vs[1].(int)]
	}).(ApplicationGatewayHttpListenerResponseOutput)
}

type ApplicationGatewayIPConfiguration struct {
	Etag              *string      `pulumi:"etag"`
	Id                *string      `pulumi:"id"`
	Name              *string      `pulumi:"name"`
	ProvisioningState *string      `pulumi:"provisioningState"`
	Subnet            *SubResource `pulumi:"subnet"`
}





type ApplicationGatewayIPConfigurationInput interface {
	pulumi.Input

	ToApplicationGatewayIPConfigurationOutput() ApplicationGatewayIPConfigurationOutput
	ToApplicationGatewayIPConfigurationOutputWithContext(context.Context) ApplicationGatewayIPConfigurationOutput
}

type ApplicationGatewayIPConfigurationArgs struct {
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
	Subnet            SubResourcePtrInput   `pulumi:"subnet"`
}

func (ApplicationGatewayIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayIPConfiguration)(nil)).Elem()
}

func (i ApplicationGatewayIPConfigurationArgs) ToApplicationGatewayIPConfigurationOutput() ApplicationGatewayIPConfigurationOutput {
	return i.ToApplicationGatewayIPConfigurationOutputWithContext(context.Background())
}

func (i ApplicationGatewayIPConfigurationArgs) ToApplicationGatewayIPConfigurationOutputWithContext(ctx context.Context) ApplicationGatewayIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayIPConfigurationOutput)
}





type ApplicationGatewayIPConfigurationArrayInput interface {
	pulumi.Input

	ToApplicationGatewayIPConfigurationArrayOutput() ApplicationGatewayIPConfigurationArrayOutput
	ToApplicationGatewayIPConfigurationArrayOutputWithContext(context.Context) ApplicationGatewayIPConfigurationArrayOutput
}

type ApplicationGatewayIPConfigurationArray []ApplicationGatewayIPConfigurationInput

func (ApplicationGatewayIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayIPConfiguration)(nil)).Elem()
}

func (i ApplicationGatewayIPConfigurationArray) ToApplicationGatewayIPConfigurationArrayOutput() ApplicationGatewayIPConfigurationArrayOutput {
	return i.ToApplicationGatewayIPConfigurationArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayIPConfigurationArray) ToApplicationGatewayIPConfigurationArrayOutputWithContext(ctx context.Context) ApplicationGatewayIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayIPConfigurationArrayOutput)
}

type ApplicationGatewayIPConfigurationOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayIPConfiguration)(nil)).Elem()
}

func (o ApplicationGatewayIPConfigurationOutput) ToApplicationGatewayIPConfigurationOutput() ApplicationGatewayIPConfigurationOutput {
	return o
}

func (o ApplicationGatewayIPConfigurationOutput) ToApplicationGatewayIPConfigurationOutputWithContext(ctx context.Context) ApplicationGatewayIPConfigurationOutput {
	return o
}

func (o ApplicationGatewayIPConfigurationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfiguration) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayIPConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayIPConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayIPConfigurationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfiguration) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayIPConfigurationOutput) Subnet() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfiguration) *SubResource { return v.Subnet }).(SubResourcePtrOutput)
}

type ApplicationGatewayIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayIPConfiguration)(nil)).Elem()
}

func (o ApplicationGatewayIPConfigurationArrayOutput) ToApplicationGatewayIPConfigurationArrayOutput() ApplicationGatewayIPConfigurationArrayOutput {
	return o
}

func (o ApplicationGatewayIPConfigurationArrayOutput) ToApplicationGatewayIPConfigurationArrayOutputWithContext(ctx context.Context) ApplicationGatewayIPConfigurationArrayOutput {
	return o
}

func (o ApplicationGatewayIPConfigurationArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayIPConfiguration {
		return vs[0].([]ApplicationGatewayIPConfiguration)[vs[1].(int)]
	}).(ApplicationGatewayIPConfigurationOutput)
}

type ApplicationGatewayIPConfigurationResponse struct {
	Etag              *string              `pulumi:"etag"`
	Id                *string              `pulumi:"id"`
	Name              *string              `pulumi:"name"`
	ProvisioningState *string              `pulumi:"provisioningState"`
	Subnet            *SubResourceResponse `pulumi:"subnet"`
}

type ApplicationGatewayIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayIPConfigurationResponse)(nil)).Elem()
}

func (o ApplicationGatewayIPConfigurationResponseOutput) ToApplicationGatewayIPConfigurationResponseOutput() ApplicationGatewayIPConfigurationResponseOutput {
	return o
}

func (o ApplicationGatewayIPConfigurationResponseOutput) ToApplicationGatewayIPConfigurationResponseOutputWithContext(ctx context.Context) ApplicationGatewayIPConfigurationResponseOutput {
	return o
}

func (o ApplicationGatewayIPConfigurationResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfigurationResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayIPConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayIPConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayIPConfigurationResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfigurationResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayIPConfigurationResponseOutput) Subnet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayIPConfigurationResponse) *SubResourceResponse { return v.Subnet }).(SubResourceResponsePtrOutput)
}

type ApplicationGatewayIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayIPConfigurationResponse)(nil)).Elem()
}

func (o ApplicationGatewayIPConfigurationResponseArrayOutput) ToApplicationGatewayIPConfigurationResponseArrayOutput() ApplicationGatewayIPConfigurationResponseArrayOutput {
	return o
}

func (o ApplicationGatewayIPConfigurationResponseArrayOutput) ToApplicationGatewayIPConfigurationResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayIPConfigurationResponseArrayOutput {
	return o
}

func (o ApplicationGatewayIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayIPConfigurationResponse {
		return vs[0].([]ApplicationGatewayIPConfigurationResponse)[vs[1].(int)]
	}).(ApplicationGatewayIPConfigurationResponseOutput)
}

type ApplicationGatewayPathRule struct {
	BackendAddressPool  *SubResource `pulumi:"backendAddressPool"`
	BackendHttpSettings *SubResource `pulumi:"backendHttpSettings"`
	Etag                *string      `pulumi:"etag"`
	Id                  *string      `pulumi:"id"`
	Name                *string      `pulumi:"name"`
	Paths               []string     `pulumi:"paths"`
	ProvisioningState   *string      `pulumi:"provisioningState"`
}





type ApplicationGatewayPathRuleInput interface {
	pulumi.Input

	ToApplicationGatewayPathRuleOutput() ApplicationGatewayPathRuleOutput
	ToApplicationGatewayPathRuleOutputWithContext(context.Context) ApplicationGatewayPathRuleOutput
}

type ApplicationGatewayPathRuleArgs struct {
	BackendAddressPool  SubResourcePtrInput     `pulumi:"backendAddressPool"`
	BackendHttpSettings SubResourcePtrInput     `pulumi:"backendHttpSettings"`
	Etag                pulumi.StringPtrInput   `pulumi:"etag"`
	Id                  pulumi.StringPtrInput   `pulumi:"id"`
	Name                pulumi.StringPtrInput   `pulumi:"name"`
	Paths               pulumi.StringArrayInput `pulumi:"paths"`
	ProvisioningState   pulumi.StringPtrInput   `pulumi:"provisioningState"`
}

func (ApplicationGatewayPathRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayPathRule)(nil)).Elem()
}

func (i ApplicationGatewayPathRuleArgs) ToApplicationGatewayPathRuleOutput() ApplicationGatewayPathRuleOutput {
	return i.ToApplicationGatewayPathRuleOutputWithContext(context.Background())
}

func (i ApplicationGatewayPathRuleArgs) ToApplicationGatewayPathRuleOutputWithContext(ctx context.Context) ApplicationGatewayPathRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayPathRuleOutput)
}





type ApplicationGatewayPathRuleArrayInput interface {
	pulumi.Input

	ToApplicationGatewayPathRuleArrayOutput() ApplicationGatewayPathRuleArrayOutput
	ToApplicationGatewayPathRuleArrayOutputWithContext(context.Context) ApplicationGatewayPathRuleArrayOutput
}

type ApplicationGatewayPathRuleArray []ApplicationGatewayPathRuleInput

func (ApplicationGatewayPathRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayPathRule)(nil)).Elem()
}

func (i ApplicationGatewayPathRuleArray) ToApplicationGatewayPathRuleArrayOutput() ApplicationGatewayPathRuleArrayOutput {
	return i.ToApplicationGatewayPathRuleArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayPathRuleArray) ToApplicationGatewayPathRuleArrayOutputWithContext(ctx context.Context) ApplicationGatewayPathRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayPathRuleArrayOutput)
}

type ApplicationGatewayPathRuleOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayPathRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayPathRule)(nil)).Elem()
}

func (o ApplicationGatewayPathRuleOutput) ToApplicationGatewayPathRuleOutput() ApplicationGatewayPathRuleOutput {
	return o
}

func (o ApplicationGatewayPathRuleOutput) ToApplicationGatewayPathRuleOutputWithContext(ctx context.Context) ApplicationGatewayPathRuleOutput {
	return o
}

func (o ApplicationGatewayPathRuleOutput) BackendAddressPool() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRule) *SubResource { return v.BackendAddressPool }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayPathRuleOutput) BackendHttpSettings() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRule) *SubResource { return v.BackendHttpSettings }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayPathRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRule) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayPathRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayPathRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayPathRuleOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRule) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

func (o ApplicationGatewayPathRuleOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRule) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayPathRuleArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayPathRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayPathRule)(nil)).Elem()
}

func (o ApplicationGatewayPathRuleArrayOutput) ToApplicationGatewayPathRuleArrayOutput() ApplicationGatewayPathRuleArrayOutput {
	return o
}

func (o ApplicationGatewayPathRuleArrayOutput) ToApplicationGatewayPathRuleArrayOutputWithContext(ctx context.Context) ApplicationGatewayPathRuleArrayOutput {
	return o
}

func (o ApplicationGatewayPathRuleArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayPathRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayPathRule {
		return vs[0].([]ApplicationGatewayPathRule)[vs[1].(int)]
	}).(ApplicationGatewayPathRuleOutput)
}

type ApplicationGatewayPathRuleResponse struct {
	BackendAddressPool  *SubResourceResponse `pulumi:"backendAddressPool"`
	BackendHttpSettings *SubResourceResponse `pulumi:"backendHttpSettings"`
	Etag                *string              `pulumi:"etag"`
	Id                  *string              `pulumi:"id"`
	Name                *string              `pulumi:"name"`
	Paths               []string             `pulumi:"paths"`
	ProvisioningState   *string              `pulumi:"provisioningState"`
}

type ApplicationGatewayPathRuleResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayPathRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayPathRuleResponse)(nil)).Elem()
}

func (o ApplicationGatewayPathRuleResponseOutput) ToApplicationGatewayPathRuleResponseOutput() ApplicationGatewayPathRuleResponseOutput {
	return o
}

func (o ApplicationGatewayPathRuleResponseOutput) ToApplicationGatewayPathRuleResponseOutputWithContext(ctx context.Context) ApplicationGatewayPathRuleResponseOutput {
	return o
}

func (o ApplicationGatewayPathRuleResponseOutput) BackendAddressPool() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRuleResponse) *SubResourceResponse { return v.BackendAddressPool }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayPathRuleResponseOutput) BackendHttpSettings() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRuleResponse) *SubResourceResponse { return v.BackendHttpSettings }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayPathRuleResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRuleResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayPathRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayPathRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayPathRuleResponseOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRuleResponse) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

func (o ApplicationGatewayPathRuleResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayPathRuleResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayPathRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayPathRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayPathRuleResponse)(nil)).Elem()
}

func (o ApplicationGatewayPathRuleResponseArrayOutput) ToApplicationGatewayPathRuleResponseArrayOutput() ApplicationGatewayPathRuleResponseArrayOutput {
	return o
}

func (o ApplicationGatewayPathRuleResponseArrayOutput) ToApplicationGatewayPathRuleResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayPathRuleResponseArrayOutput {
	return o
}

func (o ApplicationGatewayPathRuleResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayPathRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayPathRuleResponse {
		return vs[0].([]ApplicationGatewayPathRuleResponse)[vs[1].(int)]
	}).(ApplicationGatewayPathRuleResponseOutput)
}

type ApplicationGatewayProbe struct {
	Etag               *string `pulumi:"etag"`
	Host               *string `pulumi:"host"`
	Id                 *string `pulumi:"id"`
	Interval           *int    `pulumi:"interval"`
	Name               *string `pulumi:"name"`
	Path               *string `pulumi:"path"`
	Protocol           *string `pulumi:"protocol"`
	ProvisioningState  *string `pulumi:"provisioningState"`
	Timeout            *int    `pulumi:"timeout"`
	UnhealthyThreshold *int    `pulumi:"unhealthyThreshold"`
}





type ApplicationGatewayProbeInput interface {
	pulumi.Input

	ToApplicationGatewayProbeOutput() ApplicationGatewayProbeOutput
	ToApplicationGatewayProbeOutputWithContext(context.Context) ApplicationGatewayProbeOutput
}

type ApplicationGatewayProbeArgs struct {
	Etag               pulumi.StringPtrInput `pulumi:"etag"`
	Host               pulumi.StringPtrInput `pulumi:"host"`
	Id                 pulumi.StringPtrInput `pulumi:"id"`
	Interval           pulumi.IntPtrInput    `pulumi:"interval"`
	Name               pulumi.StringPtrInput `pulumi:"name"`
	Path               pulumi.StringPtrInput `pulumi:"path"`
	Protocol           pulumi.StringPtrInput `pulumi:"protocol"`
	ProvisioningState  pulumi.StringPtrInput `pulumi:"provisioningState"`
	Timeout            pulumi.IntPtrInput    `pulumi:"timeout"`
	UnhealthyThreshold pulumi.IntPtrInput    `pulumi:"unhealthyThreshold"`
}

func (ApplicationGatewayProbeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayProbe)(nil)).Elem()
}

func (i ApplicationGatewayProbeArgs) ToApplicationGatewayProbeOutput() ApplicationGatewayProbeOutput {
	return i.ToApplicationGatewayProbeOutputWithContext(context.Background())
}

func (i ApplicationGatewayProbeArgs) ToApplicationGatewayProbeOutputWithContext(ctx context.Context) ApplicationGatewayProbeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayProbeOutput)
}





type ApplicationGatewayProbeArrayInput interface {
	pulumi.Input

	ToApplicationGatewayProbeArrayOutput() ApplicationGatewayProbeArrayOutput
	ToApplicationGatewayProbeArrayOutputWithContext(context.Context) ApplicationGatewayProbeArrayOutput
}

type ApplicationGatewayProbeArray []ApplicationGatewayProbeInput

func (ApplicationGatewayProbeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayProbe)(nil)).Elem()
}

func (i ApplicationGatewayProbeArray) ToApplicationGatewayProbeArrayOutput() ApplicationGatewayProbeArrayOutput {
	return i.ToApplicationGatewayProbeArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayProbeArray) ToApplicationGatewayProbeArrayOutputWithContext(ctx context.Context) ApplicationGatewayProbeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayProbeArrayOutput)
}

type ApplicationGatewayProbeOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayProbeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayProbe)(nil)).Elem()
}

func (o ApplicationGatewayProbeOutput) ToApplicationGatewayProbeOutput() ApplicationGatewayProbeOutput {
	return o
}

func (o ApplicationGatewayProbeOutput) ToApplicationGatewayProbeOutputWithContext(ctx context.Context) ApplicationGatewayProbeOutput {
	return o
}

func (o ApplicationGatewayProbeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayProbeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayProbeOutput) UnhealthyThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbe) *int { return v.UnhealthyThreshold }).(pulumi.IntPtrOutput)
}

type ApplicationGatewayProbeArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayProbeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayProbe)(nil)).Elem()
}

func (o ApplicationGatewayProbeArrayOutput) ToApplicationGatewayProbeArrayOutput() ApplicationGatewayProbeArrayOutput {
	return o
}

func (o ApplicationGatewayProbeArrayOutput) ToApplicationGatewayProbeArrayOutputWithContext(ctx context.Context) ApplicationGatewayProbeArrayOutput {
	return o
}

func (o ApplicationGatewayProbeArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayProbeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayProbe {
		return vs[0].([]ApplicationGatewayProbe)[vs[1].(int)]
	}).(ApplicationGatewayProbeOutput)
}

type ApplicationGatewayProbeResponse struct {
	Etag               *string `pulumi:"etag"`
	Host               *string `pulumi:"host"`
	Id                 *string `pulumi:"id"`
	Interval           *int    `pulumi:"interval"`
	Name               *string `pulumi:"name"`
	Path               *string `pulumi:"path"`
	Protocol           *string `pulumi:"protocol"`
	ProvisioningState  *string `pulumi:"provisioningState"`
	Timeout            *int    `pulumi:"timeout"`
	UnhealthyThreshold *int    `pulumi:"unhealthyThreshold"`
}

type ApplicationGatewayProbeResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayProbeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayProbeResponse)(nil)).Elem()
}

func (o ApplicationGatewayProbeResponseOutput) ToApplicationGatewayProbeResponseOutput() ApplicationGatewayProbeResponseOutput {
	return o
}

func (o ApplicationGatewayProbeResponseOutput) ToApplicationGatewayProbeResponseOutputWithContext(ctx context.Context) ApplicationGatewayProbeResponseOutput {
	return o
}

func (o ApplicationGatewayProbeResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewayProbeResponseOutput) UnhealthyThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayProbeResponse) *int { return v.UnhealthyThreshold }).(pulumi.IntPtrOutput)
}

type ApplicationGatewayProbeResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayProbeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayProbeResponse)(nil)).Elem()
}

func (o ApplicationGatewayProbeResponseArrayOutput) ToApplicationGatewayProbeResponseArrayOutput() ApplicationGatewayProbeResponseArrayOutput {
	return o
}

func (o ApplicationGatewayProbeResponseArrayOutput) ToApplicationGatewayProbeResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayProbeResponseArrayOutput {
	return o
}

func (o ApplicationGatewayProbeResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayProbeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayProbeResponse {
		return vs[0].([]ApplicationGatewayProbeResponse)[vs[1].(int)]
	}).(ApplicationGatewayProbeResponseOutput)
}

type ApplicationGatewayRequestRoutingRule struct {
	BackendAddressPool  *SubResource `pulumi:"backendAddressPool"`
	BackendHttpSettings *SubResource `pulumi:"backendHttpSettings"`
	Etag                *string      `pulumi:"etag"`
	HttpListener        *SubResource `pulumi:"httpListener"`
	Id                  *string      `pulumi:"id"`
	Name                *string      `pulumi:"name"`
	ProvisioningState   *string      `pulumi:"provisioningState"`
	RuleType            *string      `pulumi:"ruleType"`
	UrlPathMap          *SubResource `pulumi:"urlPathMap"`
}





type ApplicationGatewayRequestRoutingRuleInput interface {
	pulumi.Input

	ToApplicationGatewayRequestRoutingRuleOutput() ApplicationGatewayRequestRoutingRuleOutput
	ToApplicationGatewayRequestRoutingRuleOutputWithContext(context.Context) ApplicationGatewayRequestRoutingRuleOutput
}

type ApplicationGatewayRequestRoutingRuleArgs struct {
	BackendAddressPool  SubResourcePtrInput   `pulumi:"backendAddressPool"`
	BackendHttpSettings SubResourcePtrInput   `pulumi:"backendHttpSettings"`
	Etag                pulumi.StringPtrInput `pulumi:"etag"`
	HttpListener        SubResourcePtrInput   `pulumi:"httpListener"`
	Id                  pulumi.StringPtrInput `pulumi:"id"`
	Name                pulumi.StringPtrInput `pulumi:"name"`
	ProvisioningState   pulumi.StringPtrInput `pulumi:"provisioningState"`
	RuleType            pulumi.StringPtrInput `pulumi:"ruleType"`
	UrlPathMap          SubResourcePtrInput   `pulumi:"urlPathMap"`
}

func (ApplicationGatewayRequestRoutingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayRequestRoutingRule)(nil)).Elem()
}

func (i ApplicationGatewayRequestRoutingRuleArgs) ToApplicationGatewayRequestRoutingRuleOutput() ApplicationGatewayRequestRoutingRuleOutput {
	return i.ToApplicationGatewayRequestRoutingRuleOutputWithContext(context.Background())
}

func (i ApplicationGatewayRequestRoutingRuleArgs) ToApplicationGatewayRequestRoutingRuleOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayRequestRoutingRuleOutput)
}





type ApplicationGatewayRequestRoutingRuleArrayInput interface {
	pulumi.Input

	ToApplicationGatewayRequestRoutingRuleArrayOutput() ApplicationGatewayRequestRoutingRuleArrayOutput
	ToApplicationGatewayRequestRoutingRuleArrayOutputWithContext(context.Context) ApplicationGatewayRequestRoutingRuleArrayOutput
}

type ApplicationGatewayRequestRoutingRuleArray []ApplicationGatewayRequestRoutingRuleInput

func (ApplicationGatewayRequestRoutingRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayRequestRoutingRule)(nil)).Elem()
}

func (i ApplicationGatewayRequestRoutingRuleArray) ToApplicationGatewayRequestRoutingRuleArrayOutput() ApplicationGatewayRequestRoutingRuleArrayOutput {
	return i.ToApplicationGatewayRequestRoutingRuleArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayRequestRoutingRuleArray) ToApplicationGatewayRequestRoutingRuleArrayOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayRequestRoutingRuleArrayOutput)
}

type ApplicationGatewayRequestRoutingRuleOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayRequestRoutingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayRequestRoutingRule)(nil)).Elem()
}

func (o ApplicationGatewayRequestRoutingRuleOutput) ToApplicationGatewayRequestRoutingRuleOutput() ApplicationGatewayRequestRoutingRuleOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleOutput) ToApplicationGatewayRequestRoutingRuleOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleOutput) BackendAddressPool() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *SubResource { return v.BackendAddressPool }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleOutput) BackendHttpSettings() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *SubResource { return v.BackendHttpSettings }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleOutput) HttpListener() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *SubResource { return v.HttpListener }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleOutput) RuleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *string { return v.RuleType }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleOutput) UrlPathMap() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRule) *SubResource { return v.UrlPathMap }).(SubResourcePtrOutput)
}

type ApplicationGatewayRequestRoutingRuleArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayRequestRoutingRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayRequestRoutingRule)(nil)).Elem()
}

func (o ApplicationGatewayRequestRoutingRuleArrayOutput) ToApplicationGatewayRequestRoutingRuleArrayOutput() ApplicationGatewayRequestRoutingRuleArrayOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleArrayOutput) ToApplicationGatewayRequestRoutingRuleArrayOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleArrayOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayRequestRoutingRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayRequestRoutingRule {
		return vs[0].([]ApplicationGatewayRequestRoutingRule)[vs[1].(int)]
	}).(ApplicationGatewayRequestRoutingRuleOutput)
}

type ApplicationGatewayRequestRoutingRuleResponse struct {
	BackendAddressPool  *SubResourceResponse `pulumi:"backendAddressPool"`
	BackendHttpSettings *SubResourceResponse `pulumi:"backendHttpSettings"`
	Etag                *string              `pulumi:"etag"`
	HttpListener        *SubResourceResponse `pulumi:"httpListener"`
	Id                  *string              `pulumi:"id"`
	Name                *string              `pulumi:"name"`
	ProvisioningState   *string              `pulumi:"provisioningState"`
	RuleType            *string              `pulumi:"ruleType"`
	UrlPathMap          *SubResourceResponse `pulumi:"urlPathMap"`
}

type ApplicationGatewayRequestRoutingRuleResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayRequestRoutingRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayRequestRoutingRuleResponse)(nil)).Elem()
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) ToApplicationGatewayRequestRoutingRuleResponseOutput() ApplicationGatewayRequestRoutingRuleResponseOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) ToApplicationGatewayRequestRoutingRuleResponseOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleResponseOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) BackendAddressPool() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *SubResourceResponse { return v.BackendAddressPool }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) BackendHttpSettings() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *SubResourceResponse {
		return v.BackendHttpSettings
	}).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) HttpListener() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *SubResourceResponse { return v.HttpListener }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) RuleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *string { return v.RuleType }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayRequestRoutingRuleResponseOutput) UrlPathMap() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayRequestRoutingRuleResponse) *SubResourceResponse { return v.UrlPathMap }).(SubResourceResponsePtrOutput)
}

type ApplicationGatewayRequestRoutingRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayRequestRoutingRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayRequestRoutingRuleResponse)(nil)).Elem()
}

func (o ApplicationGatewayRequestRoutingRuleResponseArrayOutput) ToApplicationGatewayRequestRoutingRuleResponseArrayOutput() ApplicationGatewayRequestRoutingRuleResponseArrayOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleResponseArrayOutput) ToApplicationGatewayRequestRoutingRuleResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayRequestRoutingRuleResponseArrayOutput {
	return o
}

func (o ApplicationGatewayRequestRoutingRuleResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayRequestRoutingRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayRequestRoutingRuleResponse {
		return vs[0].([]ApplicationGatewayRequestRoutingRuleResponse)[vs[1].(int)]
	}).(ApplicationGatewayRequestRoutingRuleResponseOutput)
}

type ApplicationGatewaySku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type ApplicationGatewaySkuInput interface {
	pulumi.Input

	ToApplicationGatewaySkuOutput() ApplicationGatewaySkuOutput
	ToApplicationGatewaySkuOutputWithContext(context.Context) ApplicationGatewaySkuOutput
}

type ApplicationGatewaySkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (ApplicationGatewaySkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySku)(nil)).Elem()
}

func (i ApplicationGatewaySkuArgs) ToApplicationGatewaySkuOutput() ApplicationGatewaySkuOutput {
	return i.ToApplicationGatewaySkuOutputWithContext(context.Background())
}

func (i ApplicationGatewaySkuArgs) ToApplicationGatewaySkuOutputWithContext(ctx context.Context) ApplicationGatewaySkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySkuOutput)
}

func (i ApplicationGatewaySkuArgs) ToApplicationGatewaySkuPtrOutput() ApplicationGatewaySkuPtrOutput {
	return i.ToApplicationGatewaySkuPtrOutputWithContext(context.Background())
}

func (i ApplicationGatewaySkuArgs) ToApplicationGatewaySkuPtrOutputWithContext(ctx context.Context) ApplicationGatewaySkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySkuOutput).ToApplicationGatewaySkuPtrOutputWithContext(ctx)
}









type ApplicationGatewaySkuPtrInput interface {
	pulumi.Input

	ToApplicationGatewaySkuPtrOutput() ApplicationGatewaySkuPtrOutput
	ToApplicationGatewaySkuPtrOutputWithContext(context.Context) ApplicationGatewaySkuPtrOutput
}

type applicationGatewaySkuPtrType ApplicationGatewaySkuArgs

func ApplicationGatewaySkuPtr(v *ApplicationGatewaySkuArgs) ApplicationGatewaySkuPtrInput {
	return (*applicationGatewaySkuPtrType)(v)
}

func (*applicationGatewaySkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySku)(nil)).Elem()
}

func (i *applicationGatewaySkuPtrType) ToApplicationGatewaySkuPtrOutput() ApplicationGatewaySkuPtrOutput {
	return i.ToApplicationGatewaySkuPtrOutputWithContext(context.Background())
}

func (i *applicationGatewaySkuPtrType) ToApplicationGatewaySkuPtrOutputWithContext(ctx context.Context) ApplicationGatewaySkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySkuPtrOutput)
}

type ApplicationGatewaySkuOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySku)(nil)).Elem()
}

func (o ApplicationGatewaySkuOutput) ToApplicationGatewaySkuOutput() ApplicationGatewaySkuOutput {
	return o
}

func (o ApplicationGatewaySkuOutput) ToApplicationGatewaySkuOutputWithContext(ctx context.Context) ApplicationGatewaySkuOutput {
	return o
}

func (o ApplicationGatewaySkuOutput) ToApplicationGatewaySkuPtrOutput() ApplicationGatewaySkuPtrOutput {
	return o.ToApplicationGatewaySkuPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySkuOutput) ToApplicationGatewaySkuPtrOutputWithContext(ctx context.Context) ApplicationGatewaySkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewaySku) *ApplicationGatewaySku {
		return &v
	}).(ApplicationGatewaySkuPtrOutput)
}

func (o ApplicationGatewaySkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewaySkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySkuPtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySku)(nil)).Elem()
}

func (o ApplicationGatewaySkuPtrOutput) ToApplicationGatewaySkuPtrOutput() ApplicationGatewaySkuPtrOutput {
	return o
}

func (o ApplicationGatewaySkuPtrOutput) ToApplicationGatewaySkuPtrOutputWithContext(ctx context.Context) ApplicationGatewaySkuPtrOutput {
	return o
}

func (o ApplicationGatewaySkuPtrOutput) Elem() ApplicationGatewaySkuOutput {
	return o.ApplyT(func(v *ApplicationGatewaySku) ApplicationGatewaySku {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewaySku
		return ret
	}).(ApplicationGatewaySkuOutput)
}

func (o ApplicationGatewaySkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewaySku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewaySkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewaySku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewaySku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}

type ApplicationGatewaySkuResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySkuResponse)(nil)).Elem()
}

func (o ApplicationGatewaySkuResponseOutput) ToApplicationGatewaySkuResponseOutput() ApplicationGatewaySkuResponseOutput {
	return o
}

func (o ApplicationGatewaySkuResponseOutput) ToApplicationGatewaySkuResponseOutputWithContext(ctx context.Context) ApplicationGatewaySkuResponseOutput {
	return o
}

func (o ApplicationGatewaySkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewaySkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySkuResponse)(nil)).Elem()
}

func (o ApplicationGatewaySkuResponsePtrOutput) ToApplicationGatewaySkuResponsePtrOutput() ApplicationGatewaySkuResponsePtrOutput {
	return o
}

func (o ApplicationGatewaySkuResponsePtrOutput) ToApplicationGatewaySkuResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewaySkuResponsePtrOutput {
	return o
}

func (o ApplicationGatewaySkuResponsePtrOutput) Elem() ApplicationGatewaySkuResponseOutput {
	return o.ApplyT(func(v *ApplicationGatewaySkuResponse) ApplicationGatewaySkuResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewaySkuResponse
		return ret
	}).(ApplicationGatewaySkuResponseOutput)
}

func (o ApplicationGatewaySkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewaySkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ApplicationGatewaySkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewaySkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewaySkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySslCertificate struct {
	Data              *string `pulumi:"data"`
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	Password          *string `pulumi:"password"`
	ProvisioningState *string `pulumi:"provisioningState"`
	PublicCertData    *string `pulumi:"publicCertData"`
}





type ApplicationGatewaySslCertificateInput interface {
	pulumi.Input

	ToApplicationGatewaySslCertificateOutput() ApplicationGatewaySslCertificateOutput
	ToApplicationGatewaySslCertificateOutputWithContext(context.Context) ApplicationGatewaySslCertificateOutput
}

type ApplicationGatewaySslCertificateArgs struct {
	Data              pulumi.StringPtrInput `pulumi:"data"`
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	Password          pulumi.StringPtrInput `pulumi:"password"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
	PublicCertData    pulumi.StringPtrInput `pulumi:"publicCertData"`
}

func (ApplicationGatewaySslCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslCertificate)(nil)).Elem()
}

func (i ApplicationGatewaySslCertificateArgs) ToApplicationGatewaySslCertificateOutput() ApplicationGatewaySslCertificateOutput {
	return i.ToApplicationGatewaySslCertificateOutputWithContext(context.Background())
}

func (i ApplicationGatewaySslCertificateArgs) ToApplicationGatewaySslCertificateOutputWithContext(ctx context.Context) ApplicationGatewaySslCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySslCertificateOutput)
}





type ApplicationGatewaySslCertificateArrayInput interface {
	pulumi.Input

	ToApplicationGatewaySslCertificateArrayOutput() ApplicationGatewaySslCertificateArrayOutput
	ToApplicationGatewaySslCertificateArrayOutputWithContext(context.Context) ApplicationGatewaySslCertificateArrayOutput
}

type ApplicationGatewaySslCertificateArray []ApplicationGatewaySslCertificateInput

func (ApplicationGatewaySslCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewaySslCertificate)(nil)).Elem()
}

func (i ApplicationGatewaySslCertificateArray) ToApplicationGatewaySslCertificateArrayOutput() ApplicationGatewaySslCertificateArrayOutput {
	return i.ToApplicationGatewaySslCertificateArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewaySslCertificateArray) ToApplicationGatewaySslCertificateArrayOutputWithContext(ctx context.Context) ApplicationGatewaySslCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySslCertificateArrayOutput)
}

type ApplicationGatewaySslCertificateOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslCertificate)(nil)).Elem()
}

func (o ApplicationGatewaySslCertificateOutput) ToApplicationGatewaySslCertificateOutput() ApplicationGatewaySslCertificateOutput {
	return o
}

func (o ApplicationGatewaySslCertificateOutput) ToApplicationGatewaySslCertificateOutputWithContext(ctx context.Context) ApplicationGatewaySslCertificateOutput {
	return o
}

func (o ApplicationGatewaySslCertificateOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificate) *string { return v.Data }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificate) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificate) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificate) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificate) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateOutput) PublicCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificate) *string { return v.PublicCertData }).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySslCertificateArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewaySslCertificate)(nil)).Elem()
}

func (o ApplicationGatewaySslCertificateArrayOutput) ToApplicationGatewaySslCertificateArrayOutput() ApplicationGatewaySslCertificateArrayOutput {
	return o
}

func (o ApplicationGatewaySslCertificateArrayOutput) ToApplicationGatewaySslCertificateArrayOutputWithContext(ctx context.Context) ApplicationGatewaySslCertificateArrayOutput {
	return o
}

func (o ApplicationGatewaySslCertificateArrayOutput) Index(i pulumi.IntInput) ApplicationGatewaySslCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewaySslCertificate {
		return vs[0].([]ApplicationGatewaySslCertificate)[vs[1].(int)]
	}).(ApplicationGatewaySslCertificateOutput)
}

type ApplicationGatewaySslCertificateResponse struct {
	Data              *string `pulumi:"data"`
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	Password          *string `pulumi:"password"`
	ProvisioningState *string `pulumi:"provisioningState"`
	PublicCertData    *string `pulumi:"publicCertData"`
}

type ApplicationGatewaySslCertificateResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslCertificateResponse)(nil)).Elem()
}

func (o ApplicationGatewaySslCertificateResponseOutput) ToApplicationGatewaySslCertificateResponseOutput() ApplicationGatewaySslCertificateResponseOutput {
	return o
}

func (o ApplicationGatewaySslCertificateResponseOutput) ToApplicationGatewaySslCertificateResponseOutputWithContext(ctx context.Context) ApplicationGatewaySslCertificateResponseOutput {
	return o
}

func (o ApplicationGatewaySslCertificateResponseOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificateResponse) *string { return v.Data }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificateResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificateResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificateResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificateResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewaySslCertificateResponseOutput) PublicCertData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewaySslCertificateResponse) *string { return v.PublicCertData }).(pulumi.StringPtrOutput)
}

type ApplicationGatewaySslCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewaySslCertificateResponse)(nil)).Elem()
}

func (o ApplicationGatewaySslCertificateResponseArrayOutput) ToApplicationGatewaySslCertificateResponseArrayOutput() ApplicationGatewaySslCertificateResponseArrayOutput {
	return o
}

func (o ApplicationGatewaySslCertificateResponseArrayOutput) ToApplicationGatewaySslCertificateResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewaySslCertificateResponseArrayOutput {
	return o
}

func (o ApplicationGatewaySslCertificateResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewaySslCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewaySslCertificateResponse {
		return vs[0].([]ApplicationGatewaySslCertificateResponse)[vs[1].(int)]
	}).(ApplicationGatewaySslCertificateResponseOutput)
}

type ApplicationGatewaySslPolicy struct {
	DisabledSslProtocols []string `pulumi:"disabledSslProtocols"`
}





type ApplicationGatewaySslPolicyInput interface {
	pulumi.Input

	ToApplicationGatewaySslPolicyOutput() ApplicationGatewaySslPolicyOutput
	ToApplicationGatewaySslPolicyOutputWithContext(context.Context) ApplicationGatewaySslPolicyOutput
}

type ApplicationGatewaySslPolicyArgs struct {
	DisabledSslProtocols pulumi.StringArrayInput `pulumi:"disabledSslProtocols"`
}

func (ApplicationGatewaySslPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslPolicy)(nil)).Elem()
}

func (i ApplicationGatewaySslPolicyArgs) ToApplicationGatewaySslPolicyOutput() ApplicationGatewaySslPolicyOutput {
	return i.ToApplicationGatewaySslPolicyOutputWithContext(context.Background())
}

func (i ApplicationGatewaySslPolicyArgs) ToApplicationGatewaySslPolicyOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySslPolicyOutput)
}

func (i ApplicationGatewaySslPolicyArgs) ToApplicationGatewaySslPolicyPtrOutput() ApplicationGatewaySslPolicyPtrOutput {
	return i.ToApplicationGatewaySslPolicyPtrOutputWithContext(context.Background())
}

func (i ApplicationGatewaySslPolicyArgs) ToApplicationGatewaySslPolicyPtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySslPolicyOutput).ToApplicationGatewaySslPolicyPtrOutputWithContext(ctx)
}









type ApplicationGatewaySslPolicyPtrInput interface {
	pulumi.Input

	ToApplicationGatewaySslPolicyPtrOutput() ApplicationGatewaySslPolicyPtrOutput
	ToApplicationGatewaySslPolicyPtrOutputWithContext(context.Context) ApplicationGatewaySslPolicyPtrOutput
}

type applicationGatewaySslPolicyPtrType ApplicationGatewaySslPolicyArgs

func ApplicationGatewaySslPolicyPtr(v *ApplicationGatewaySslPolicyArgs) ApplicationGatewaySslPolicyPtrInput {
	return (*applicationGatewaySslPolicyPtrType)(v)
}

func (*applicationGatewaySslPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySslPolicy)(nil)).Elem()
}

func (i *applicationGatewaySslPolicyPtrType) ToApplicationGatewaySslPolicyPtrOutput() ApplicationGatewaySslPolicyPtrOutput {
	return i.ToApplicationGatewaySslPolicyPtrOutputWithContext(context.Background())
}

func (i *applicationGatewaySslPolicyPtrType) ToApplicationGatewaySslPolicyPtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewaySslPolicyPtrOutput)
}

type ApplicationGatewaySslPolicyOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslPolicy)(nil)).Elem()
}

func (o ApplicationGatewaySslPolicyOutput) ToApplicationGatewaySslPolicyOutput() ApplicationGatewaySslPolicyOutput {
	return o
}

func (o ApplicationGatewaySslPolicyOutput) ToApplicationGatewaySslPolicyOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyOutput {
	return o
}

func (o ApplicationGatewaySslPolicyOutput) ToApplicationGatewaySslPolicyPtrOutput() ApplicationGatewaySslPolicyPtrOutput {
	return o.ToApplicationGatewaySslPolicyPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewaySslPolicyOutput) ToApplicationGatewaySslPolicyPtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewaySslPolicy) *ApplicationGatewaySslPolicy {
		return &v
	}).(ApplicationGatewaySslPolicyPtrOutput)
}

func (o ApplicationGatewaySslPolicyOutput) DisabledSslProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApplicationGatewaySslPolicy) []string { return v.DisabledSslProtocols }).(pulumi.StringArrayOutput)
}

type ApplicationGatewaySslPolicyPtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySslPolicy)(nil)).Elem()
}

func (o ApplicationGatewaySslPolicyPtrOutput) ToApplicationGatewaySslPolicyPtrOutput() ApplicationGatewaySslPolicyPtrOutput {
	return o
}

func (o ApplicationGatewaySslPolicyPtrOutput) ToApplicationGatewaySslPolicyPtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyPtrOutput {
	return o
}

func (o ApplicationGatewaySslPolicyPtrOutput) Elem() ApplicationGatewaySslPolicyOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslPolicy) ApplicationGatewaySslPolicy {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewaySslPolicy
		return ret
	}).(ApplicationGatewaySslPolicyOutput)
}

func (o ApplicationGatewaySslPolicyPtrOutput) DisabledSslProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslPolicy) []string {
		if v == nil {
			return nil
		}
		return v.DisabledSslProtocols
	}).(pulumi.StringArrayOutput)
}

type ApplicationGatewaySslPolicyResponse struct {
	DisabledSslProtocols []string `pulumi:"disabledSslProtocols"`
}

type ApplicationGatewaySslPolicyResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewaySslPolicyResponse)(nil)).Elem()
}

func (o ApplicationGatewaySslPolicyResponseOutput) ToApplicationGatewaySslPolicyResponseOutput() ApplicationGatewaySslPolicyResponseOutput {
	return o
}

func (o ApplicationGatewaySslPolicyResponseOutput) ToApplicationGatewaySslPolicyResponseOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyResponseOutput {
	return o
}

func (o ApplicationGatewaySslPolicyResponseOutput) DisabledSslProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ApplicationGatewaySslPolicyResponse) []string { return v.DisabledSslProtocols }).(pulumi.StringArrayOutput)
}

type ApplicationGatewaySslPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewaySslPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewaySslPolicyResponse)(nil)).Elem()
}

func (o ApplicationGatewaySslPolicyResponsePtrOutput) ToApplicationGatewaySslPolicyResponsePtrOutput() ApplicationGatewaySslPolicyResponsePtrOutput {
	return o
}

func (o ApplicationGatewaySslPolicyResponsePtrOutput) ToApplicationGatewaySslPolicyResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewaySslPolicyResponsePtrOutput {
	return o
}

func (o ApplicationGatewaySslPolicyResponsePtrOutput) Elem() ApplicationGatewaySslPolicyResponseOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslPolicyResponse) ApplicationGatewaySslPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewaySslPolicyResponse
		return ret
	}).(ApplicationGatewaySslPolicyResponseOutput)
}

func (o ApplicationGatewaySslPolicyResponsePtrOutput) DisabledSslProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationGatewaySslPolicyResponse) []string {
		if v == nil {
			return nil
		}
		return v.DisabledSslProtocols
	}).(pulumi.StringArrayOutput)
}

type ApplicationGatewayUrlPathMap struct {
	DefaultBackendAddressPool  *SubResource                 `pulumi:"defaultBackendAddressPool"`
	DefaultBackendHttpSettings *SubResource                 `pulumi:"defaultBackendHttpSettings"`
	Etag                       *string                      `pulumi:"etag"`
	Id                         *string                      `pulumi:"id"`
	Name                       *string                      `pulumi:"name"`
	PathRules                  []ApplicationGatewayPathRule `pulumi:"pathRules"`
	ProvisioningState          *string                      `pulumi:"provisioningState"`
}





type ApplicationGatewayUrlPathMapInput interface {
	pulumi.Input

	ToApplicationGatewayUrlPathMapOutput() ApplicationGatewayUrlPathMapOutput
	ToApplicationGatewayUrlPathMapOutputWithContext(context.Context) ApplicationGatewayUrlPathMapOutput
}

type ApplicationGatewayUrlPathMapArgs struct {
	DefaultBackendAddressPool  SubResourcePtrInput                  `pulumi:"defaultBackendAddressPool"`
	DefaultBackendHttpSettings SubResourcePtrInput                  `pulumi:"defaultBackendHttpSettings"`
	Etag                       pulumi.StringPtrInput                `pulumi:"etag"`
	Id                         pulumi.StringPtrInput                `pulumi:"id"`
	Name                       pulumi.StringPtrInput                `pulumi:"name"`
	PathRules                  ApplicationGatewayPathRuleArrayInput `pulumi:"pathRules"`
	ProvisioningState          pulumi.StringPtrInput                `pulumi:"provisioningState"`
}

func (ApplicationGatewayUrlPathMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayUrlPathMap)(nil)).Elem()
}

func (i ApplicationGatewayUrlPathMapArgs) ToApplicationGatewayUrlPathMapOutput() ApplicationGatewayUrlPathMapOutput {
	return i.ToApplicationGatewayUrlPathMapOutputWithContext(context.Background())
}

func (i ApplicationGatewayUrlPathMapArgs) ToApplicationGatewayUrlPathMapOutputWithContext(ctx context.Context) ApplicationGatewayUrlPathMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayUrlPathMapOutput)
}





type ApplicationGatewayUrlPathMapArrayInput interface {
	pulumi.Input

	ToApplicationGatewayUrlPathMapArrayOutput() ApplicationGatewayUrlPathMapArrayOutput
	ToApplicationGatewayUrlPathMapArrayOutputWithContext(context.Context) ApplicationGatewayUrlPathMapArrayOutput
}

type ApplicationGatewayUrlPathMapArray []ApplicationGatewayUrlPathMapInput

func (ApplicationGatewayUrlPathMapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayUrlPathMap)(nil)).Elem()
}

func (i ApplicationGatewayUrlPathMapArray) ToApplicationGatewayUrlPathMapArrayOutput() ApplicationGatewayUrlPathMapArrayOutput {
	return i.ToApplicationGatewayUrlPathMapArrayOutputWithContext(context.Background())
}

func (i ApplicationGatewayUrlPathMapArray) ToApplicationGatewayUrlPathMapArrayOutputWithContext(ctx context.Context) ApplicationGatewayUrlPathMapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayUrlPathMapArrayOutput)
}

type ApplicationGatewayUrlPathMapOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayUrlPathMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayUrlPathMap)(nil)).Elem()
}

func (o ApplicationGatewayUrlPathMapOutput) ToApplicationGatewayUrlPathMapOutput() ApplicationGatewayUrlPathMapOutput {
	return o
}

func (o ApplicationGatewayUrlPathMapOutput) ToApplicationGatewayUrlPathMapOutputWithContext(ctx context.Context) ApplicationGatewayUrlPathMapOutput {
	return o
}

func (o ApplicationGatewayUrlPathMapOutput) DefaultBackendAddressPool() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMap) *SubResource { return v.DefaultBackendAddressPool }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayUrlPathMapOutput) DefaultBackendHttpSettings() SubResourcePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMap) *SubResource { return v.DefaultBackendHttpSettings }).(SubResourcePtrOutput)
}

func (o ApplicationGatewayUrlPathMapOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMap) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayUrlPathMapOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMap) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayUrlPathMapOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMap) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayUrlPathMapOutput) PathRules() ApplicationGatewayPathRuleArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMap) []ApplicationGatewayPathRule { return v.PathRules }).(ApplicationGatewayPathRuleArrayOutput)
}

func (o ApplicationGatewayUrlPathMapOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMap) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayUrlPathMapArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayUrlPathMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayUrlPathMap)(nil)).Elem()
}

func (o ApplicationGatewayUrlPathMapArrayOutput) ToApplicationGatewayUrlPathMapArrayOutput() ApplicationGatewayUrlPathMapArrayOutput {
	return o
}

func (o ApplicationGatewayUrlPathMapArrayOutput) ToApplicationGatewayUrlPathMapArrayOutputWithContext(ctx context.Context) ApplicationGatewayUrlPathMapArrayOutput {
	return o
}

func (o ApplicationGatewayUrlPathMapArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayUrlPathMapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayUrlPathMap {
		return vs[0].([]ApplicationGatewayUrlPathMap)[vs[1].(int)]
	}).(ApplicationGatewayUrlPathMapOutput)
}

type ApplicationGatewayUrlPathMapResponse struct {
	DefaultBackendAddressPool  *SubResourceResponse                 `pulumi:"defaultBackendAddressPool"`
	DefaultBackendHttpSettings *SubResourceResponse                 `pulumi:"defaultBackendHttpSettings"`
	Etag                       *string                              `pulumi:"etag"`
	Id                         *string                              `pulumi:"id"`
	Name                       *string                              `pulumi:"name"`
	PathRules                  []ApplicationGatewayPathRuleResponse `pulumi:"pathRules"`
	ProvisioningState          *string                              `pulumi:"provisioningState"`
}

type ApplicationGatewayUrlPathMapResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayUrlPathMapResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayUrlPathMapResponse)(nil)).Elem()
}

func (o ApplicationGatewayUrlPathMapResponseOutput) ToApplicationGatewayUrlPathMapResponseOutput() ApplicationGatewayUrlPathMapResponseOutput {
	return o
}

func (o ApplicationGatewayUrlPathMapResponseOutput) ToApplicationGatewayUrlPathMapResponseOutputWithContext(ctx context.Context) ApplicationGatewayUrlPathMapResponseOutput {
	return o
}

func (o ApplicationGatewayUrlPathMapResponseOutput) DefaultBackendAddressPool() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMapResponse) *SubResourceResponse { return v.DefaultBackendAddressPool }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayUrlPathMapResponseOutput) DefaultBackendHttpSettings() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMapResponse) *SubResourceResponse { return v.DefaultBackendHttpSettings }).(SubResourceResponsePtrOutput)
}

func (o ApplicationGatewayUrlPathMapResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMapResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayUrlPathMapResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMapResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayUrlPathMapResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMapResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayUrlPathMapResponseOutput) PathRules() ApplicationGatewayPathRuleResponseArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMapResponse) []ApplicationGatewayPathRuleResponse { return v.PathRules }).(ApplicationGatewayPathRuleResponseArrayOutput)
}

func (o ApplicationGatewayUrlPathMapResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationGatewayUrlPathMapResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ApplicationGatewayUrlPathMapResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayUrlPathMapResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationGatewayUrlPathMapResponse)(nil)).Elem()
}

func (o ApplicationGatewayUrlPathMapResponseArrayOutput) ToApplicationGatewayUrlPathMapResponseArrayOutput() ApplicationGatewayUrlPathMapResponseArrayOutput {
	return o
}

func (o ApplicationGatewayUrlPathMapResponseArrayOutput) ToApplicationGatewayUrlPathMapResponseArrayOutputWithContext(ctx context.Context) ApplicationGatewayUrlPathMapResponseArrayOutput {
	return o
}

func (o ApplicationGatewayUrlPathMapResponseArrayOutput) Index(i pulumi.IntInput) ApplicationGatewayUrlPathMapResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationGatewayUrlPathMapResponse {
		return vs[0].([]ApplicationGatewayUrlPathMapResponse)[vs[1].(int)]
	}).(ApplicationGatewayUrlPathMapResponseOutput)
}

type ApplicationGatewayWebApplicationFirewallConfiguration struct {
	DisabledRuleGroups []ApplicationGatewayFirewallDisabledRuleGroup `pulumi:"disabledRuleGroups"`
	Enabled            bool                                          `pulumi:"enabled"`
	FirewallMode       string                                        `pulumi:"firewallMode"`
	RuleSetType        string                                        `pulumi:"ruleSetType"`
	RuleSetVersion     string                                        `pulumi:"ruleSetVersion"`
}





type ApplicationGatewayWebApplicationFirewallConfigurationInput interface {
	pulumi.Input

	ToApplicationGatewayWebApplicationFirewallConfigurationOutput() ApplicationGatewayWebApplicationFirewallConfigurationOutput
	ToApplicationGatewayWebApplicationFirewallConfigurationOutputWithContext(context.Context) ApplicationGatewayWebApplicationFirewallConfigurationOutput
}

type ApplicationGatewayWebApplicationFirewallConfigurationArgs struct {
	DisabledRuleGroups ApplicationGatewayFirewallDisabledRuleGroupArrayInput `pulumi:"disabledRuleGroups"`
	Enabled            pulumi.BoolInput                                      `pulumi:"enabled"`
	FirewallMode       pulumi.StringInput                                    `pulumi:"firewallMode"`
	RuleSetType        pulumi.StringInput                                    `pulumi:"ruleSetType"`
	RuleSetVersion     pulumi.StringInput                                    `pulumi:"ruleSetVersion"`
}

func (ApplicationGatewayWebApplicationFirewallConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayWebApplicationFirewallConfiguration)(nil)).Elem()
}

func (i ApplicationGatewayWebApplicationFirewallConfigurationArgs) ToApplicationGatewayWebApplicationFirewallConfigurationOutput() ApplicationGatewayWebApplicationFirewallConfigurationOutput {
	return i.ToApplicationGatewayWebApplicationFirewallConfigurationOutputWithContext(context.Background())
}

func (i ApplicationGatewayWebApplicationFirewallConfigurationArgs) ToApplicationGatewayWebApplicationFirewallConfigurationOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayWebApplicationFirewallConfigurationOutput)
}

func (i ApplicationGatewayWebApplicationFirewallConfigurationArgs) ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutput() ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput {
	return i.ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutputWithContext(context.Background())
}

func (i ApplicationGatewayWebApplicationFirewallConfigurationArgs) ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayWebApplicationFirewallConfigurationOutput).ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutputWithContext(ctx)
}









type ApplicationGatewayWebApplicationFirewallConfigurationPtrInput interface {
	pulumi.Input

	ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutput() ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput
	ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutputWithContext(context.Context) ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput
}

type applicationGatewayWebApplicationFirewallConfigurationPtrType ApplicationGatewayWebApplicationFirewallConfigurationArgs

func ApplicationGatewayWebApplicationFirewallConfigurationPtr(v *ApplicationGatewayWebApplicationFirewallConfigurationArgs) ApplicationGatewayWebApplicationFirewallConfigurationPtrInput {
	return (*applicationGatewayWebApplicationFirewallConfigurationPtrType)(v)
}

func (*applicationGatewayWebApplicationFirewallConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayWebApplicationFirewallConfiguration)(nil)).Elem()
}

func (i *applicationGatewayWebApplicationFirewallConfigurationPtrType) ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutput() ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput {
	return i.ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutputWithContext(context.Background())
}

func (i *applicationGatewayWebApplicationFirewallConfigurationPtrType) ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput)
}

type ApplicationGatewayWebApplicationFirewallConfigurationOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayWebApplicationFirewallConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayWebApplicationFirewallConfiguration)(nil)).Elem()
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationOutput) ToApplicationGatewayWebApplicationFirewallConfigurationOutput() ApplicationGatewayWebApplicationFirewallConfigurationOutput {
	return o
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationOutput) ToApplicationGatewayWebApplicationFirewallConfigurationOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationOutput {
	return o
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationOutput) ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutput() ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput {
	return o.ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutputWithContext(context.Background())
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationOutput) ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationGatewayWebApplicationFirewallConfiguration) *ApplicationGatewayWebApplicationFirewallConfiguration {
		return &v
	}).(ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationOutput) DisabledRuleGroups() ApplicationGatewayFirewallDisabledRuleGroupArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfiguration) []ApplicationGatewayFirewallDisabledRuleGroup {
		return v.DisabledRuleGroups
	}).(ApplicationGatewayFirewallDisabledRuleGroupArrayOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfiguration) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationOutput) FirewallMode() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfiguration) string { return v.FirewallMode }).(pulumi.StringOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationOutput) RuleSetType() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfiguration) string { return v.RuleSetType }).(pulumi.StringOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationOutput) RuleSetVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfiguration) string { return v.RuleSetVersion }).(pulumi.StringOutput)
}

type ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayWebApplicationFirewallConfiguration)(nil)).Elem()
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput) ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutput() ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput {
	return o
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput) ToApplicationGatewayWebApplicationFirewallConfigurationPtrOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput {
	return o
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput) Elem() ApplicationGatewayWebApplicationFirewallConfigurationOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfiguration) ApplicationGatewayWebApplicationFirewallConfiguration {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewayWebApplicationFirewallConfiguration
		return ret
	}).(ApplicationGatewayWebApplicationFirewallConfigurationOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput) DisabledRuleGroups() ApplicationGatewayFirewallDisabledRuleGroupArrayOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfiguration) []ApplicationGatewayFirewallDisabledRuleGroup {
		if v == nil {
			return nil
		}
		return v.DisabledRuleGroups
	}).(ApplicationGatewayFirewallDisabledRuleGroupArrayOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfiguration) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput) FirewallMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.FirewallMode
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput) RuleSetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.RuleSetType
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput) RuleSetVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.RuleSetVersion
	}).(pulumi.StringPtrOutput)
}

type ApplicationGatewayWebApplicationFirewallConfigurationResponse struct {
	DisabledRuleGroups []ApplicationGatewayFirewallDisabledRuleGroupResponse `pulumi:"disabledRuleGroups"`
	Enabled            bool                                                  `pulumi:"enabled"`
	FirewallMode       string                                                `pulumi:"firewallMode"`
	RuleSetType        string                                                `pulumi:"ruleSetType"`
	RuleSetVersion     string                                                `pulumi:"ruleSetVersion"`
}

type ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationGatewayWebApplicationFirewallConfigurationResponse)(nil)).Elem()
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput) ToApplicationGatewayWebApplicationFirewallConfigurationResponseOutput() ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput {
	return o
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput) ToApplicationGatewayWebApplicationFirewallConfigurationResponseOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput {
	return o
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput) DisabledRuleGroups() ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfigurationResponse) []ApplicationGatewayFirewallDisabledRuleGroupResponse {
		return v.DisabledRuleGroups
	}).(ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfigurationResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput) FirewallMode() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfigurationResponse) string { return v.FirewallMode }).(pulumi.StringOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput) RuleSetType() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfigurationResponse) string { return v.RuleSetType }).(pulumi.StringOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput) RuleSetVersion() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationGatewayWebApplicationFirewallConfigurationResponse) string { return v.RuleSetVersion }).(pulumi.StringOutput)
}

type ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayWebApplicationFirewallConfigurationResponse)(nil)).Elem()
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput) ToApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput() ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput {
	return o
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput) ToApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutputWithContext(ctx context.Context) ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput {
	return o
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput) Elem() ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfigurationResponse) ApplicationGatewayWebApplicationFirewallConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationGatewayWebApplicationFirewallConfigurationResponse
		return ret
	}).(ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput) DisabledRuleGroups() ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfigurationResponse) []ApplicationGatewayFirewallDisabledRuleGroupResponse {
		if v == nil {
			return nil
		}
		return v.DisabledRuleGroups
	}).(ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput) FirewallMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FirewallMode
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput) RuleSetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RuleSetType
	}).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput) RuleSetVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayWebApplicationFirewallConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RuleSetVersion
	}).(pulumi.StringPtrOutput)
}

type BackendAddressPool struct {
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	ProvisioningState *string `pulumi:"provisioningState"`
}





type BackendAddressPoolInput interface {
	pulumi.Input

	ToBackendAddressPoolOutput() BackendAddressPoolOutput
	ToBackendAddressPoolOutputWithContext(context.Context) BackendAddressPoolOutput
}

type BackendAddressPoolArgs struct {
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (BackendAddressPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendAddressPool)(nil)).Elem()
}

func (i BackendAddressPoolArgs) ToBackendAddressPoolOutput() BackendAddressPoolOutput {
	return i.ToBackendAddressPoolOutputWithContext(context.Background())
}

func (i BackendAddressPoolArgs) ToBackendAddressPoolOutputWithContext(ctx context.Context) BackendAddressPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAddressPoolOutput)
}





type BackendAddressPoolArrayInput interface {
	pulumi.Input

	ToBackendAddressPoolArrayOutput() BackendAddressPoolArrayOutput
	ToBackendAddressPoolArrayOutputWithContext(context.Context) BackendAddressPoolArrayOutput
}

type BackendAddressPoolArray []BackendAddressPoolInput

func (BackendAddressPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendAddressPool)(nil)).Elem()
}

func (i BackendAddressPoolArray) ToBackendAddressPoolArrayOutput() BackendAddressPoolArrayOutput {
	return i.ToBackendAddressPoolArrayOutputWithContext(context.Background())
}

func (i BackendAddressPoolArray) ToBackendAddressPoolArrayOutputWithContext(ctx context.Context) BackendAddressPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackendAddressPoolArrayOutput)
}

type BackendAddressPoolOutput struct{ *pulumi.OutputState }

func (BackendAddressPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendAddressPool)(nil)).Elem()
}

func (o BackendAddressPoolOutput) ToBackendAddressPoolOutput() BackendAddressPoolOutput {
	return o
}

func (o BackendAddressPoolOutput) ToBackendAddressPoolOutputWithContext(ctx context.Context) BackendAddressPoolOutput {
	return o
}

func (o BackendAddressPoolOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPool) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o BackendAddressPoolOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPool) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o BackendAddressPoolOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPool) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o BackendAddressPoolOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPool) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type BackendAddressPoolArrayOutput struct{ *pulumi.OutputState }

func (BackendAddressPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendAddressPool)(nil)).Elem()
}

func (o BackendAddressPoolArrayOutput) ToBackendAddressPoolArrayOutput() BackendAddressPoolArrayOutput {
	return o
}

func (o BackendAddressPoolArrayOutput) ToBackendAddressPoolArrayOutputWithContext(ctx context.Context) BackendAddressPoolArrayOutput {
	return o
}

func (o BackendAddressPoolArrayOutput) Index(i pulumi.IntInput) BackendAddressPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackendAddressPool {
		return vs[0].([]BackendAddressPool)[vs[1].(int)]
	}).(BackendAddressPoolOutput)
}

type BackendAddressPoolResponse struct {
	BackendIPConfigurations []NetworkInterfaceIPConfigurationResponse `pulumi:"backendIPConfigurations"`
	Etag                    *string                                   `pulumi:"etag"`
	Id                      *string                                   `pulumi:"id"`
	LoadBalancingRules      []SubResourceResponse                     `pulumi:"loadBalancingRules"`
	Name                    *string                                   `pulumi:"name"`
	OutboundNatRule         SubResourceResponse                       `pulumi:"outboundNatRule"`
	ProvisioningState       *string                                   `pulumi:"provisioningState"`
}

type BackendAddressPoolResponseOutput struct{ *pulumi.OutputState }

func (BackendAddressPoolResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackendAddressPoolResponse)(nil)).Elem()
}

func (o BackendAddressPoolResponseOutput) ToBackendAddressPoolResponseOutput() BackendAddressPoolResponseOutput {
	return o
}

func (o BackendAddressPoolResponseOutput) ToBackendAddressPoolResponseOutputWithContext(ctx context.Context) BackendAddressPoolResponseOutput {
	return o
}

func (o BackendAddressPoolResponseOutput) BackendIPConfigurations() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v BackendAddressPoolResponse) []NetworkInterfaceIPConfigurationResponse {
		return v.BackendIPConfigurations
	}).(NetworkInterfaceIPConfigurationResponseArrayOutput)
}

func (o BackendAddressPoolResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPoolResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o BackendAddressPoolResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPoolResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o BackendAddressPoolResponseOutput) LoadBalancingRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v BackendAddressPoolResponse) []SubResourceResponse { return v.LoadBalancingRules }).(SubResourceResponseArrayOutput)
}

func (o BackendAddressPoolResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPoolResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o BackendAddressPoolResponseOutput) OutboundNatRule() SubResourceResponseOutput {
	return o.ApplyT(func(v BackendAddressPoolResponse) SubResourceResponse { return v.OutboundNatRule }).(SubResourceResponseOutput)
}

func (o BackendAddressPoolResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackendAddressPoolResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type BackendAddressPoolResponseArrayOutput struct{ *pulumi.OutputState }

func (BackendAddressPoolResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackendAddressPoolResponse)(nil)).Elem()
}

func (o BackendAddressPoolResponseArrayOutput) ToBackendAddressPoolResponseArrayOutput() BackendAddressPoolResponseArrayOutput {
	return o
}

func (o BackendAddressPoolResponseArrayOutput) ToBackendAddressPoolResponseArrayOutputWithContext(ctx context.Context) BackendAddressPoolResponseArrayOutput {
	return o
}

func (o BackendAddressPoolResponseArrayOutput) Index(i pulumi.IntInput) BackendAddressPoolResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackendAddressPoolResponse {
		return vs[0].([]BackendAddressPoolResponse)[vs[1].(int)]
	}).(BackendAddressPoolResponseOutput)
}

type BgpPeerStatusResponse struct {
	Asn               int     `pulumi:"asn"`
	ConnectedDuration string  `pulumi:"connectedDuration"`
	LocalAddress      string  `pulumi:"localAddress"`
	MessagesReceived  float64 `pulumi:"messagesReceived"`
	MessagesSent      float64 `pulumi:"messagesSent"`
	Neighbor          string  `pulumi:"neighbor"`
	RoutesReceived    float64 `pulumi:"routesReceived"`
	State             string  `pulumi:"state"`
}

type BgpPeerStatusResponseOutput struct{ *pulumi.OutputState }

func (BgpPeerStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpPeerStatusResponse)(nil)).Elem()
}

func (o BgpPeerStatusResponseOutput) ToBgpPeerStatusResponseOutput() BgpPeerStatusResponseOutput {
	return o
}

func (o BgpPeerStatusResponseOutput) ToBgpPeerStatusResponseOutputWithContext(ctx context.Context) BgpPeerStatusResponseOutput {
	return o
}

func (o BgpPeerStatusResponseOutput) Asn() pulumi.IntOutput {
	return o.ApplyT(func(v BgpPeerStatusResponse) int { return v.Asn }).(pulumi.IntOutput)
}

func (o BgpPeerStatusResponseOutput) ConnectedDuration() pulumi.StringOutput {
	return o.ApplyT(func(v BgpPeerStatusResponse) string { return v.ConnectedDuration }).(pulumi.StringOutput)
}

func (o BgpPeerStatusResponseOutput) LocalAddress() pulumi.StringOutput {
	return o.ApplyT(func(v BgpPeerStatusResponse) string { return v.LocalAddress }).(pulumi.StringOutput)
}

func (o BgpPeerStatusResponseOutput) MessagesReceived() pulumi.Float64Output {
	return o.ApplyT(func(v BgpPeerStatusResponse) float64 { return v.MessagesReceived }).(pulumi.Float64Output)
}

func (o BgpPeerStatusResponseOutput) MessagesSent() pulumi.Float64Output {
	return o.ApplyT(func(v BgpPeerStatusResponse) float64 { return v.MessagesSent }).(pulumi.Float64Output)
}

func (o BgpPeerStatusResponseOutput) Neighbor() pulumi.StringOutput {
	return o.ApplyT(func(v BgpPeerStatusResponse) string { return v.Neighbor }).(pulumi.StringOutput)
}

func (o BgpPeerStatusResponseOutput) RoutesReceived() pulumi.Float64Output {
	return o.ApplyT(func(v BgpPeerStatusResponse) float64 { return v.RoutesReceived }).(pulumi.Float64Output)
}

func (o BgpPeerStatusResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v BgpPeerStatusResponse) string { return v.State }).(pulumi.StringOutput)
}

type BgpPeerStatusResponseArrayOutput struct{ *pulumi.OutputState }

func (BgpPeerStatusResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BgpPeerStatusResponse)(nil)).Elem()
}

func (o BgpPeerStatusResponseArrayOutput) ToBgpPeerStatusResponseArrayOutput() BgpPeerStatusResponseArrayOutput {
	return o
}

func (o BgpPeerStatusResponseArrayOutput) ToBgpPeerStatusResponseArrayOutputWithContext(ctx context.Context) BgpPeerStatusResponseArrayOutput {
	return o
}

func (o BgpPeerStatusResponseArrayOutput) Index(i pulumi.IntInput) BgpPeerStatusResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BgpPeerStatusResponse {
		return vs[0].([]BgpPeerStatusResponse)[vs[1].(int)]
	}).(BgpPeerStatusResponseOutput)
}

type BgpSettings struct {
	Asn               *float64 `pulumi:"asn"`
	BgpPeeringAddress *string  `pulumi:"bgpPeeringAddress"`
	PeerWeight        *int     `pulumi:"peerWeight"`
}





type BgpSettingsInput interface {
	pulumi.Input

	ToBgpSettingsOutput() BgpSettingsOutput
	ToBgpSettingsOutputWithContext(context.Context) BgpSettingsOutput
}

type BgpSettingsArgs struct {
	Asn               pulumi.Float64PtrInput `pulumi:"asn"`
	BgpPeeringAddress pulumi.StringPtrInput  `pulumi:"bgpPeeringAddress"`
	PeerWeight        pulumi.IntPtrInput     `pulumi:"peerWeight"`
}

func (BgpSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpSettings)(nil)).Elem()
}

func (i BgpSettingsArgs) ToBgpSettingsOutput() BgpSettingsOutput {
	return i.ToBgpSettingsOutputWithContext(context.Background())
}

func (i BgpSettingsArgs) ToBgpSettingsOutputWithContext(ctx context.Context) BgpSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpSettingsOutput)
}

func (i BgpSettingsArgs) ToBgpSettingsPtrOutput() BgpSettingsPtrOutput {
	return i.ToBgpSettingsPtrOutputWithContext(context.Background())
}

func (i BgpSettingsArgs) ToBgpSettingsPtrOutputWithContext(ctx context.Context) BgpSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpSettingsOutput).ToBgpSettingsPtrOutputWithContext(ctx)
}









type BgpSettingsPtrInput interface {
	pulumi.Input

	ToBgpSettingsPtrOutput() BgpSettingsPtrOutput
	ToBgpSettingsPtrOutputWithContext(context.Context) BgpSettingsPtrOutput
}

type bgpSettingsPtrType BgpSettingsArgs

func BgpSettingsPtr(v *BgpSettingsArgs) BgpSettingsPtrInput {
	return (*bgpSettingsPtrType)(v)
}

func (*bgpSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpSettings)(nil)).Elem()
}

func (i *bgpSettingsPtrType) ToBgpSettingsPtrOutput() BgpSettingsPtrOutput {
	return i.ToBgpSettingsPtrOutputWithContext(context.Background())
}

func (i *bgpSettingsPtrType) ToBgpSettingsPtrOutputWithContext(ctx context.Context) BgpSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BgpSettingsPtrOutput)
}

type BgpSettingsOutput struct{ *pulumi.OutputState }

func (BgpSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpSettings)(nil)).Elem()
}

func (o BgpSettingsOutput) ToBgpSettingsOutput() BgpSettingsOutput {
	return o
}

func (o BgpSettingsOutput) ToBgpSettingsOutputWithContext(ctx context.Context) BgpSettingsOutput {
	return o
}

func (o BgpSettingsOutput) ToBgpSettingsPtrOutput() BgpSettingsPtrOutput {
	return o.ToBgpSettingsPtrOutputWithContext(context.Background())
}

func (o BgpSettingsOutput) ToBgpSettingsPtrOutputWithContext(ctx context.Context) BgpSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BgpSettings) *BgpSettings {
		return &v
	}).(BgpSettingsPtrOutput)
}

func (o BgpSettingsOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v BgpSettings) *float64 { return v.Asn }).(pulumi.Float64PtrOutput)
}

func (o BgpSettingsOutput) BgpPeeringAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSettings) *string { return v.BgpPeeringAddress }).(pulumi.StringPtrOutput)
}

func (o BgpSettingsOutput) PeerWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BgpSettings) *int { return v.PeerWeight }).(pulumi.IntPtrOutput)
}

type BgpSettingsPtrOutput struct{ *pulumi.OutputState }

func (BgpSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpSettings)(nil)).Elem()
}

func (o BgpSettingsPtrOutput) ToBgpSettingsPtrOutput() BgpSettingsPtrOutput {
	return o
}

func (o BgpSettingsPtrOutput) ToBgpSettingsPtrOutputWithContext(ctx context.Context) BgpSettingsPtrOutput {
	return o
}

func (o BgpSettingsPtrOutput) Elem() BgpSettingsOutput {
	return o.ApplyT(func(v *BgpSettings) BgpSettings {
		if v != nil {
			return *v
		}
		var ret BgpSettings
		return ret
	}).(BgpSettingsOutput)
}

func (o BgpSettingsPtrOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *BgpSettings) *float64 {
		if v == nil {
			return nil
		}
		return v.Asn
	}).(pulumi.Float64PtrOutput)
}

func (o BgpSettingsPtrOutput) BgpPeeringAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSettings) *string {
		if v == nil {
			return nil
		}
		return v.BgpPeeringAddress
	}).(pulumi.StringPtrOutput)
}

func (o BgpSettingsPtrOutput) PeerWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BgpSettings) *int {
		if v == nil {
			return nil
		}
		return v.PeerWeight
	}).(pulumi.IntPtrOutput)
}

type BgpSettingsResponse struct {
	Asn               *float64 `pulumi:"asn"`
	BgpPeeringAddress *string  `pulumi:"bgpPeeringAddress"`
	PeerWeight        *int     `pulumi:"peerWeight"`
}

type BgpSettingsResponseOutput struct{ *pulumi.OutputState }

func (BgpSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BgpSettingsResponse)(nil)).Elem()
}

func (o BgpSettingsResponseOutput) ToBgpSettingsResponseOutput() BgpSettingsResponseOutput {
	return o
}

func (o BgpSettingsResponseOutput) ToBgpSettingsResponseOutputWithContext(ctx context.Context) BgpSettingsResponseOutput {
	return o
}

func (o BgpSettingsResponseOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v BgpSettingsResponse) *float64 { return v.Asn }).(pulumi.Float64PtrOutput)
}

func (o BgpSettingsResponseOutput) BgpPeeringAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BgpSettingsResponse) *string { return v.BgpPeeringAddress }).(pulumi.StringPtrOutput)
}

func (o BgpSettingsResponseOutput) PeerWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BgpSettingsResponse) *int { return v.PeerWeight }).(pulumi.IntPtrOutput)
}

type BgpSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (BgpSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BgpSettingsResponse)(nil)).Elem()
}

func (o BgpSettingsResponsePtrOutput) ToBgpSettingsResponsePtrOutput() BgpSettingsResponsePtrOutput {
	return o
}

func (o BgpSettingsResponsePtrOutput) ToBgpSettingsResponsePtrOutputWithContext(ctx context.Context) BgpSettingsResponsePtrOutput {
	return o
}

func (o BgpSettingsResponsePtrOutput) Elem() BgpSettingsResponseOutput {
	return o.ApplyT(func(v *BgpSettingsResponse) BgpSettingsResponse {
		if v != nil {
			return *v
		}
		var ret BgpSettingsResponse
		return ret
	}).(BgpSettingsResponseOutput)
}

func (o BgpSettingsResponsePtrOutput) Asn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *BgpSettingsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Asn
	}).(pulumi.Float64PtrOutput)
}

func (o BgpSettingsResponsePtrOutput) BgpPeeringAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BgpSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.BgpPeeringAddress
	}).(pulumi.StringPtrOutput)
}

func (o BgpSettingsResponsePtrOutput) PeerWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BgpSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.PeerWeight
	}).(pulumi.IntPtrOutput)
}

type DhcpOptions struct {
	DnsServers []string `pulumi:"dnsServers"`
}





type DhcpOptionsInput interface {
	pulumi.Input

	ToDhcpOptionsOutput() DhcpOptionsOutput
	ToDhcpOptionsOutputWithContext(context.Context) DhcpOptionsOutput
}

type DhcpOptionsArgs struct {
	DnsServers pulumi.StringArrayInput `pulumi:"dnsServers"`
}

func (DhcpOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DhcpOptions)(nil)).Elem()
}

func (i DhcpOptionsArgs) ToDhcpOptionsOutput() DhcpOptionsOutput {
	return i.ToDhcpOptionsOutputWithContext(context.Background())
}

func (i DhcpOptionsArgs) ToDhcpOptionsOutputWithContext(ctx context.Context) DhcpOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpOptionsOutput)
}

func (i DhcpOptionsArgs) ToDhcpOptionsPtrOutput() DhcpOptionsPtrOutput {
	return i.ToDhcpOptionsPtrOutputWithContext(context.Background())
}

func (i DhcpOptionsArgs) ToDhcpOptionsPtrOutputWithContext(ctx context.Context) DhcpOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpOptionsOutput).ToDhcpOptionsPtrOutputWithContext(ctx)
}









type DhcpOptionsPtrInput interface {
	pulumi.Input

	ToDhcpOptionsPtrOutput() DhcpOptionsPtrOutput
	ToDhcpOptionsPtrOutputWithContext(context.Context) DhcpOptionsPtrOutput
}

type dhcpOptionsPtrType DhcpOptionsArgs

func DhcpOptionsPtr(v *DhcpOptionsArgs) DhcpOptionsPtrInput {
	return (*dhcpOptionsPtrType)(v)
}

func (*dhcpOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpOptions)(nil)).Elem()
}

func (i *dhcpOptionsPtrType) ToDhcpOptionsPtrOutput() DhcpOptionsPtrOutput {
	return i.ToDhcpOptionsPtrOutputWithContext(context.Background())
}

func (i *dhcpOptionsPtrType) ToDhcpOptionsPtrOutputWithContext(ctx context.Context) DhcpOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpOptionsPtrOutput)
}

type DhcpOptionsOutput struct{ *pulumi.OutputState }

func (DhcpOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DhcpOptions)(nil)).Elem()
}

func (o DhcpOptionsOutput) ToDhcpOptionsOutput() DhcpOptionsOutput {
	return o
}

func (o DhcpOptionsOutput) ToDhcpOptionsOutputWithContext(ctx context.Context) DhcpOptionsOutput {
	return o
}

func (o DhcpOptionsOutput) ToDhcpOptionsPtrOutput() DhcpOptionsPtrOutput {
	return o.ToDhcpOptionsPtrOutputWithContext(context.Background())
}

func (o DhcpOptionsOutput) ToDhcpOptionsPtrOutputWithContext(ctx context.Context) DhcpOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DhcpOptions) *DhcpOptions {
		return &v
	}).(DhcpOptionsPtrOutput)
}

func (o DhcpOptionsOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DhcpOptions) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

type DhcpOptionsPtrOutput struct{ *pulumi.OutputState }

func (DhcpOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpOptions)(nil)).Elem()
}

func (o DhcpOptionsPtrOutput) ToDhcpOptionsPtrOutput() DhcpOptionsPtrOutput {
	return o
}

func (o DhcpOptionsPtrOutput) ToDhcpOptionsPtrOutputWithContext(ctx context.Context) DhcpOptionsPtrOutput {
	return o
}

func (o DhcpOptionsPtrOutput) Elem() DhcpOptionsOutput {
	return o.ApplyT(func(v *DhcpOptions) DhcpOptions {
		if v != nil {
			return *v
		}
		var ret DhcpOptions
		return ret
	}).(DhcpOptionsOutput)
}

func (o DhcpOptionsPtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DhcpOptions) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

type DhcpOptionsResponse struct {
	DnsServers []string `pulumi:"dnsServers"`
}

type DhcpOptionsResponseOutput struct{ *pulumi.OutputState }

func (DhcpOptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DhcpOptionsResponse)(nil)).Elem()
}

func (o DhcpOptionsResponseOutput) ToDhcpOptionsResponseOutput() DhcpOptionsResponseOutput {
	return o
}

func (o DhcpOptionsResponseOutput) ToDhcpOptionsResponseOutputWithContext(ctx context.Context) DhcpOptionsResponseOutput {
	return o
}

func (o DhcpOptionsResponseOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DhcpOptionsResponse) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

type DhcpOptionsResponsePtrOutput struct{ *pulumi.OutputState }

func (DhcpOptionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpOptionsResponse)(nil)).Elem()
}

func (o DhcpOptionsResponsePtrOutput) ToDhcpOptionsResponsePtrOutput() DhcpOptionsResponsePtrOutput {
	return o
}

func (o DhcpOptionsResponsePtrOutput) ToDhcpOptionsResponsePtrOutputWithContext(ctx context.Context) DhcpOptionsResponsePtrOutput {
	return o
}

func (o DhcpOptionsResponsePtrOutput) Elem() DhcpOptionsResponseOutput {
	return o.ApplyT(func(v *DhcpOptionsResponse) DhcpOptionsResponse {
		if v != nil {
			return *v
		}
		var ret DhcpOptionsResponse
		return ret
	}).(DhcpOptionsResponseOutput)
}

func (o DhcpOptionsResponsePtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DhcpOptionsResponse) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

type DnsConfig struct {
	Fqdn         *string  `pulumi:"fqdn"`
	RelativeName *string  `pulumi:"relativeName"`
	Ttl          *float64 `pulumi:"ttl"`
}





type DnsConfigInput interface {
	pulumi.Input

	ToDnsConfigOutput() DnsConfigOutput
	ToDnsConfigOutputWithContext(context.Context) DnsConfigOutput
}

type DnsConfigArgs struct {
	Fqdn         pulumi.StringPtrInput  `pulumi:"fqdn"`
	RelativeName pulumi.StringPtrInput  `pulumi:"relativeName"`
	Ttl          pulumi.Float64PtrInput `pulumi:"ttl"`
}

func (DnsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsConfig)(nil)).Elem()
}

func (i DnsConfigArgs) ToDnsConfigOutput() DnsConfigOutput {
	return i.ToDnsConfigOutputWithContext(context.Background())
}

func (i DnsConfigArgs) ToDnsConfigOutputWithContext(ctx context.Context) DnsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsConfigOutput)
}

func (i DnsConfigArgs) ToDnsConfigPtrOutput() DnsConfigPtrOutput {
	return i.ToDnsConfigPtrOutputWithContext(context.Background())
}

func (i DnsConfigArgs) ToDnsConfigPtrOutputWithContext(ctx context.Context) DnsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsConfigOutput).ToDnsConfigPtrOutputWithContext(ctx)
}









type DnsConfigPtrInput interface {
	pulumi.Input

	ToDnsConfigPtrOutput() DnsConfigPtrOutput
	ToDnsConfigPtrOutputWithContext(context.Context) DnsConfigPtrOutput
}

type dnsConfigPtrType DnsConfigArgs

func DnsConfigPtr(v *DnsConfigArgs) DnsConfigPtrInput {
	return (*dnsConfigPtrType)(v)
}

func (*dnsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsConfig)(nil)).Elem()
}

func (i *dnsConfigPtrType) ToDnsConfigPtrOutput() DnsConfigPtrOutput {
	return i.ToDnsConfigPtrOutputWithContext(context.Background())
}

func (i *dnsConfigPtrType) ToDnsConfigPtrOutputWithContext(ctx context.Context) DnsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsConfigPtrOutput)
}

type DnsConfigOutput struct{ *pulumi.OutputState }

func (DnsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsConfig)(nil)).Elem()
}

func (o DnsConfigOutput) ToDnsConfigOutput() DnsConfigOutput {
	return o
}

func (o DnsConfigOutput) ToDnsConfigOutputWithContext(ctx context.Context) DnsConfigOutput {
	return o
}

func (o DnsConfigOutput) ToDnsConfigPtrOutput() DnsConfigPtrOutput {
	return o.ToDnsConfigPtrOutputWithContext(context.Background())
}

func (o DnsConfigOutput) ToDnsConfigPtrOutputWithContext(ctx context.Context) DnsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DnsConfig) *DnsConfig {
		return &v
	}).(DnsConfigPtrOutput)
}

func (o DnsConfigOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnsConfig) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o DnsConfigOutput) RelativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnsConfig) *string { return v.RelativeName }).(pulumi.StringPtrOutput)
}

func (o DnsConfigOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DnsConfig) *float64 { return v.Ttl }).(pulumi.Float64PtrOutput)
}

type DnsConfigPtrOutput struct{ *pulumi.OutputState }

func (DnsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsConfig)(nil)).Elem()
}

func (o DnsConfigPtrOutput) ToDnsConfigPtrOutput() DnsConfigPtrOutput {
	return o
}

func (o DnsConfigPtrOutput) ToDnsConfigPtrOutputWithContext(ctx context.Context) DnsConfigPtrOutput {
	return o
}

func (o DnsConfigPtrOutput) Elem() DnsConfigOutput {
	return o.ApplyT(func(v *DnsConfig) DnsConfig {
		if v != nil {
			return *v
		}
		var ret DnsConfig
		return ret
	}).(DnsConfigOutput)
}

func (o DnsConfigPtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsConfig) *string {
		if v == nil {
			return nil
		}
		return v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o DnsConfigPtrOutput) RelativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsConfig) *string {
		if v == nil {
			return nil
		}
		return v.RelativeName
	}).(pulumi.StringPtrOutput)
}

func (o DnsConfigPtrOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DnsConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.Ttl
	}).(pulumi.Float64PtrOutput)
}

type DnsConfigResponse struct {
	Fqdn         *string  `pulumi:"fqdn"`
	RelativeName *string  `pulumi:"relativeName"`
	Ttl          *float64 `pulumi:"ttl"`
}

type DnsConfigResponseOutput struct{ *pulumi.OutputState }

func (DnsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsConfigResponse)(nil)).Elem()
}

func (o DnsConfigResponseOutput) ToDnsConfigResponseOutput() DnsConfigResponseOutput {
	return o
}

func (o DnsConfigResponseOutput) ToDnsConfigResponseOutputWithContext(ctx context.Context) DnsConfigResponseOutput {
	return o
}

func (o DnsConfigResponseOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnsConfigResponse) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o DnsConfigResponseOutput) RelativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnsConfigResponse) *string { return v.RelativeName }).(pulumi.StringPtrOutput)
}

func (o DnsConfigResponseOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DnsConfigResponse) *float64 { return v.Ttl }).(pulumi.Float64PtrOutput)
}

type DnsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (DnsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsConfigResponse)(nil)).Elem()
}

func (o DnsConfigResponsePtrOutput) ToDnsConfigResponsePtrOutput() DnsConfigResponsePtrOutput {
	return o
}

func (o DnsConfigResponsePtrOutput) ToDnsConfigResponsePtrOutputWithContext(ctx context.Context) DnsConfigResponsePtrOutput {
	return o
}

func (o DnsConfigResponsePtrOutput) Elem() DnsConfigResponseOutput {
	return o.ApplyT(func(v *DnsConfigResponse) DnsConfigResponse {
		if v != nil {
			return *v
		}
		var ret DnsConfigResponse
		return ret
	}).(DnsConfigResponseOutput)
}

func (o DnsConfigResponsePtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o DnsConfigResponsePtrOutput) RelativeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.RelativeName
	}).(pulumi.StringPtrOutput)
}

func (o DnsConfigResponsePtrOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DnsConfigResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Ttl
	}).(pulumi.Float64PtrOutput)
}

type EndpointType struct {
	EndpointLocation      *string  `pulumi:"endpointLocation"`
	EndpointMonitorStatus *string  `pulumi:"endpointMonitorStatus"`
	EndpointStatus        *string  `pulumi:"endpointStatus"`
	GeoMapping            []string `pulumi:"geoMapping"`
	Id                    *string  `pulumi:"id"`
	MinChildEndpoints     *float64 `pulumi:"minChildEndpoints"`
	Name                  *string  `pulumi:"name"`
	Priority              *float64 `pulumi:"priority"`
	Target                *string  `pulumi:"target"`
	TargetResourceId      *string  `pulumi:"targetResourceId"`
	Type                  *string  `pulumi:"type"`
	Weight                *float64 `pulumi:"weight"`
}





type EndpointTypeInput interface {
	pulumi.Input

	ToEndpointTypeOutput() EndpointTypeOutput
	ToEndpointTypeOutputWithContext(context.Context) EndpointTypeOutput
}

type EndpointTypeArgs struct {
	EndpointLocation      pulumi.StringPtrInput   `pulumi:"endpointLocation"`
	EndpointMonitorStatus pulumi.StringPtrInput   `pulumi:"endpointMonitorStatus"`
	EndpointStatus        pulumi.StringPtrInput   `pulumi:"endpointStatus"`
	GeoMapping            pulumi.StringArrayInput `pulumi:"geoMapping"`
	Id                    pulumi.StringPtrInput   `pulumi:"id"`
	MinChildEndpoints     pulumi.Float64PtrInput  `pulumi:"minChildEndpoints"`
	Name                  pulumi.StringPtrInput   `pulumi:"name"`
	Priority              pulumi.Float64PtrInput  `pulumi:"priority"`
	Target                pulumi.StringPtrInput   `pulumi:"target"`
	TargetResourceId      pulumi.StringPtrInput   `pulumi:"targetResourceId"`
	Type                  pulumi.StringPtrInput   `pulumi:"type"`
	Weight                pulumi.Float64PtrInput  `pulumi:"weight"`
}

func (EndpointTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointType)(nil)).Elem()
}

func (i EndpointTypeArgs) ToEndpointTypeOutput() EndpointTypeOutput {
	return i.ToEndpointTypeOutputWithContext(context.Background())
}

func (i EndpointTypeArgs) ToEndpointTypeOutputWithContext(ctx context.Context) EndpointTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointTypeOutput)
}





type EndpointTypeArrayInput interface {
	pulumi.Input

	ToEndpointTypeArrayOutput() EndpointTypeArrayOutput
	ToEndpointTypeArrayOutputWithContext(context.Context) EndpointTypeArrayOutput
}

type EndpointTypeArray []EndpointTypeInput

func (EndpointTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointType)(nil)).Elem()
}

func (i EndpointTypeArray) ToEndpointTypeArrayOutput() EndpointTypeArrayOutput {
	return i.ToEndpointTypeArrayOutputWithContext(context.Background())
}

func (i EndpointTypeArray) ToEndpointTypeArrayOutputWithContext(ctx context.Context) EndpointTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointTypeArrayOutput)
}

type EndpointTypeOutput struct{ *pulumi.OutputState }

func (EndpointTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointType)(nil)).Elem()
}

func (o EndpointTypeOutput) ToEndpointTypeOutput() EndpointTypeOutput {
	return o
}

func (o EndpointTypeOutput) ToEndpointTypeOutputWithContext(ctx context.Context) EndpointTypeOutput {
	return o
}

func (o EndpointTypeOutput) EndpointLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.EndpointLocation }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) EndpointMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.EndpointMonitorStatus }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) EndpointStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.EndpointStatus }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) GeoMapping() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EndpointType) []string { return v.GeoMapping }).(pulumi.StringArrayOutput)
}

func (o EndpointTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) MinChildEndpoints() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointType) *float64 { return v.MinChildEndpoints }).(pulumi.Float64PtrOutput)
}

func (o EndpointTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) Priority() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointType) *float64 { return v.Priority }).(pulumi.Float64PtrOutput)
}

func (o EndpointTypeOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointType) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o EndpointTypeOutput) Weight() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointType) *float64 { return v.Weight }).(pulumi.Float64PtrOutput)
}

type EndpointTypeArrayOutput struct{ *pulumi.OutputState }

func (EndpointTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointType)(nil)).Elem()
}

func (o EndpointTypeArrayOutput) ToEndpointTypeArrayOutput() EndpointTypeArrayOutput {
	return o
}

func (o EndpointTypeArrayOutput) ToEndpointTypeArrayOutputWithContext(ctx context.Context) EndpointTypeArrayOutput {
	return o
}

func (o EndpointTypeArrayOutput) Index(i pulumi.IntInput) EndpointTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointType {
		return vs[0].([]EndpointType)[vs[1].(int)]
	}).(EndpointTypeOutput)
}

type EndpointResponse struct {
	EndpointLocation      *string  `pulumi:"endpointLocation"`
	EndpointMonitorStatus *string  `pulumi:"endpointMonitorStatus"`
	EndpointStatus        *string  `pulumi:"endpointStatus"`
	GeoMapping            []string `pulumi:"geoMapping"`
	Id                    *string  `pulumi:"id"`
	MinChildEndpoints     *float64 `pulumi:"minChildEndpoints"`
	Name                  *string  `pulumi:"name"`
	Priority              *float64 `pulumi:"priority"`
	Target                *string  `pulumi:"target"`
	TargetResourceId      *string  `pulumi:"targetResourceId"`
	Type                  *string  `pulumi:"type"`
	Weight                *float64 `pulumi:"weight"`
}

type EndpointResponseOutput struct{ *pulumi.OutputState }

func (EndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointResponse)(nil)).Elem()
}

func (o EndpointResponseOutput) ToEndpointResponseOutput() EndpointResponseOutput {
	return o
}

func (o EndpointResponseOutput) ToEndpointResponseOutputWithContext(ctx context.Context) EndpointResponseOutput {
	return o
}

func (o EndpointResponseOutput) EndpointLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.EndpointLocation }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) EndpointMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.EndpointMonitorStatus }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) EndpointStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.EndpointStatus }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) GeoMapping() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EndpointResponse) []string { return v.GeoMapping }).(pulumi.StringArrayOutput)
}

func (o EndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) MinChildEndpoints() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointResponse) *float64 { return v.MinChildEndpoints }).(pulumi.Float64PtrOutput)
}

func (o EndpointResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) Priority() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointResponse) *float64 { return v.Priority }).(pulumi.Float64PtrOutput)
}

func (o EndpointResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o EndpointResponseOutput) Weight() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EndpointResponse) *float64 { return v.Weight }).(pulumi.Float64PtrOutput)
}

type EndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (EndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointResponse)(nil)).Elem()
}

func (o EndpointResponseArrayOutput) ToEndpointResponseArrayOutput() EndpointResponseArrayOutput {
	return o
}

func (o EndpointResponseArrayOutput) ToEndpointResponseArrayOutputWithContext(ctx context.Context) EndpointResponseArrayOutput {
	return o
}

func (o EndpointResponseArrayOutput) Index(i pulumi.IntInput) EndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointResponse {
		return vs[0].([]EndpointResponse)[vs[1].(int)]
	}).(EndpointResponseOutput)
}

type ExpressRouteCircuitAuthorizationType struct {
	AuthorizationKey       *string `pulumi:"authorizationKey"`
	AuthorizationUseStatus *string `pulumi:"authorizationUseStatus"`
	Id                     *string `pulumi:"id"`
	Name                   *string `pulumi:"name"`
	ProvisioningState      *string `pulumi:"provisioningState"`
}





type ExpressRouteCircuitAuthorizationTypeInput interface {
	pulumi.Input

	ToExpressRouteCircuitAuthorizationTypeOutput() ExpressRouteCircuitAuthorizationTypeOutput
	ToExpressRouteCircuitAuthorizationTypeOutputWithContext(context.Context) ExpressRouteCircuitAuthorizationTypeOutput
}

type ExpressRouteCircuitAuthorizationTypeArgs struct {
	AuthorizationKey       pulumi.StringPtrInput `pulumi:"authorizationKey"`
	AuthorizationUseStatus pulumi.StringPtrInput `pulumi:"authorizationUseStatus"`
	Id                     pulumi.StringPtrInput `pulumi:"id"`
	Name                   pulumi.StringPtrInput `pulumi:"name"`
	ProvisioningState      pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (ExpressRouteCircuitAuthorizationTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitAuthorizationType)(nil)).Elem()
}

func (i ExpressRouteCircuitAuthorizationTypeArgs) ToExpressRouteCircuitAuthorizationTypeOutput() ExpressRouteCircuitAuthorizationTypeOutput {
	return i.ToExpressRouteCircuitAuthorizationTypeOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitAuthorizationTypeArgs) ToExpressRouteCircuitAuthorizationTypeOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitAuthorizationTypeOutput)
}





type ExpressRouteCircuitAuthorizationTypeArrayInput interface {
	pulumi.Input

	ToExpressRouteCircuitAuthorizationTypeArrayOutput() ExpressRouteCircuitAuthorizationTypeArrayOutput
	ToExpressRouteCircuitAuthorizationTypeArrayOutputWithContext(context.Context) ExpressRouteCircuitAuthorizationTypeArrayOutput
}

type ExpressRouteCircuitAuthorizationTypeArray []ExpressRouteCircuitAuthorizationTypeInput

func (ExpressRouteCircuitAuthorizationTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressRouteCircuitAuthorizationType)(nil)).Elem()
}

func (i ExpressRouteCircuitAuthorizationTypeArray) ToExpressRouteCircuitAuthorizationTypeArrayOutput() ExpressRouteCircuitAuthorizationTypeArrayOutput {
	return i.ToExpressRouteCircuitAuthorizationTypeArrayOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitAuthorizationTypeArray) ToExpressRouteCircuitAuthorizationTypeArrayOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitAuthorizationTypeArrayOutput)
}

type ExpressRouteCircuitAuthorizationTypeOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitAuthorizationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitAuthorizationType)(nil)).Elem()
}

func (o ExpressRouteCircuitAuthorizationTypeOutput) ToExpressRouteCircuitAuthorizationTypeOutput() ExpressRouteCircuitAuthorizationTypeOutput {
	return o
}

func (o ExpressRouteCircuitAuthorizationTypeOutput) ToExpressRouteCircuitAuthorizationTypeOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationTypeOutput {
	return o
}

func (o ExpressRouteCircuitAuthorizationTypeOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationType) *string { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitAuthorizationTypeOutput) AuthorizationUseStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationType) *string { return v.AuthorizationUseStatus }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitAuthorizationTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitAuthorizationTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitAuthorizationTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitAuthorizationTypeArrayOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitAuthorizationTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressRouteCircuitAuthorizationType)(nil)).Elem()
}

func (o ExpressRouteCircuitAuthorizationTypeArrayOutput) ToExpressRouteCircuitAuthorizationTypeArrayOutput() ExpressRouteCircuitAuthorizationTypeArrayOutput {
	return o
}

func (o ExpressRouteCircuitAuthorizationTypeArrayOutput) ToExpressRouteCircuitAuthorizationTypeArrayOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationTypeArrayOutput {
	return o
}

func (o ExpressRouteCircuitAuthorizationTypeArrayOutput) Index(i pulumi.IntInput) ExpressRouteCircuitAuthorizationTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExpressRouteCircuitAuthorizationType {
		return vs[0].([]ExpressRouteCircuitAuthorizationType)[vs[1].(int)]
	}).(ExpressRouteCircuitAuthorizationTypeOutput)
}

type ExpressRouteCircuitAuthorizationResponse struct {
	AuthorizationKey       *string `pulumi:"authorizationKey"`
	AuthorizationUseStatus *string `pulumi:"authorizationUseStatus"`
	Etag                   string  `pulumi:"etag"`
	Id                     *string `pulumi:"id"`
	Name                   *string `pulumi:"name"`
	ProvisioningState      *string `pulumi:"provisioningState"`
}

type ExpressRouteCircuitAuthorizationResponseOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitAuthorizationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitAuthorizationResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitAuthorizationResponseOutput) ToExpressRouteCircuitAuthorizationResponseOutput() ExpressRouteCircuitAuthorizationResponseOutput {
	return o
}

func (o ExpressRouteCircuitAuthorizationResponseOutput) ToExpressRouteCircuitAuthorizationResponseOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationResponseOutput {
	return o
}

func (o ExpressRouteCircuitAuthorizationResponseOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationResponse) *string { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitAuthorizationResponseOutput) AuthorizationUseStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationResponse) *string { return v.AuthorizationUseStatus }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitAuthorizationResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o ExpressRouteCircuitAuthorizationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitAuthorizationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitAuthorizationResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitAuthorizationResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitAuthorizationResponseArrayOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitAuthorizationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressRouteCircuitAuthorizationResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitAuthorizationResponseArrayOutput) ToExpressRouteCircuitAuthorizationResponseArrayOutput() ExpressRouteCircuitAuthorizationResponseArrayOutput {
	return o
}

func (o ExpressRouteCircuitAuthorizationResponseArrayOutput) ToExpressRouteCircuitAuthorizationResponseArrayOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationResponseArrayOutput {
	return o
}

func (o ExpressRouteCircuitAuthorizationResponseArrayOutput) Index(i pulumi.IntInput) ExpressRouteCircuitAuthorizationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExpressRouteCircuitAuthorizationResponse {
		return vs[0].([]ExpressRouteCircuitAuthorizationResponse)[vs[1].(int)]
	}).(ExpressRouteCircuitAuthorizationResponseOutput)
}

type ExpressRouteCircuitPeeringType struct {
	AzureASN                   *int                                  `pulumi:"azureASN"`
	GatewayManagerEtag         *string                               `pulumi:"gatewayManagerEtag"`
	Id                         *string                               `pulumi:"id"`
	Ipv6PeeringConfig          *Ipv6ExpressRouteCircuitPeeringConfig `pulumi:"ipv6PeeringConfig"`
	LastModifiedBy             *string                               `pulumi:"lastModifiedBy"`
	MicrosoftPeeringConfig     *ExpressRouteCircuitPeeringConfig     `pulumi:"microsoftPeeringConfig"`
	Name                       *string                               `pulumi:"name"`
	PeerASN                    *int                                  `pulumi:"peerASN"`
	PeeringType                *string                               `pulumi:"peeringType"`
	PrimaryAzurePort           *string                               `pulumi:"primaryAzurePort"`
	PrimaryPeerAddressPrefix   *string                               `pulumi:"primaryPeerAddressPrefix"`
	ProvisioningState          *string                               `pulumi:"provisioningState"`
	RouteFilter                *RouteFilterType                      `pulumi:"routeFilter"`
	SecondaryAzurePort         *string                               `pulumi:"secondaryAzurePort"`
	SecondaryPeerAddressPrefix *string                               `pulumi:"secondaryPeerAddressPrefix"`
	SharedKey                  *string                               `pulumi:"sharedKey"`
	State                      *string                               `pulumi:"state"`
	Stats                      *ExpressRouteCircuitStats             `pulumi:"stats"`
	VlanId                     *int                                  `pulumi:"vlanId"`
}





type ExpressRouteCircuitPeeringTypeInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringTypeOutput() ExpressRouteCircuitPeeringTypeOutput
	ToExpressRouteCircuitPeeringTypeOutputWithContext(context.Context) ExpressRouteCircuitPeeringTypeOutput
}

type ExpressRouteCircuitPeeringTypeArgs struct {
	AzureASN                   pulumi.IntPtrInput                           `pulumi:"azureASN"`
	GatewayManagerEtag         pulumi.StringPtrInput                        `pulumi:"gatewayManagerEtag"`
	Id                         pulumi.StringPtrInput                        `pulumi:"id"`
	Ipv6PeeringConfig          Ipv6ExpressRouteCircuitPeeringConfigPtrInput `pulumi:"ipv6PeeringConfig"`
	LastModifiedBy             pulumi.StringPtrInput                        `pulumi:"lastModifiedBy"`
	MicrosoftPeeringConfig     ExpressRouteCircuitPeeringConfigPtrInput     `pulumi:"microsoftPeeringConfig"`
	Name                       pulumi.StringPtrInput                        `pulumi:"name"`
	PeerASN                    pulumi.IntPtrInput                           `pulumi:"peerASN"`
	PeeringType                pulumi.StringPtrInput                        `pulumi:"peeringType"`
	PrimaryAzurePort           pulumi.StringPtrInput                        `pulumi:"primaryAzurePort"`
	PrimaryPeerAddressPrefix   pulumi.StringPtrInput                        `pulumi:"primaryPeerAddressPrefix"`
	ProvisioningState          pulumi.StringPtrInput                        `pulumi:"provisioningState"`
	RouteFilter                RouteFilterTypePtrInput                      `pulumi:"routeFilter"`
	SecondaryAzurePort         pulumi.StringPtrInput                        `pulumi:"secondaryAzurePort"`
	SecondaryPeerAddressPrefix pulumi.StringPtrInput                        `pulumi:"secondaryPeerAddressPrefix"`
	SharedKey                  pulumi.StringPtrInput                        `pulumi:"sharedKey"`
	State                      pulumi.StringPtrInput                        `pulumi:"state"`
	Stats                      ExpressRouteCircuitStatsPtrInput             `pulumi:"stats"`
	VlanId                     pulumi.IntPtrInput                           `pulumi:"vlanId"`
}

func (ExpressRouteCircuitPeeringTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringType)(nil)).Elem()
}

func (i ExpressRouteCircuitPeeringTypeArgs) ToExpressRouteCircuitPeeringTypeOutput() ExpressRouteCircuitPeeringTypeOutput {
	return i.ToExpressRouteCircuitPeeringTypeOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitPeeringTypeArgs) ToExpressRouteCircuitPeeringTypeOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringTypeOutput)
}





type ExpressRouteCircuitPeeringTypeArrayInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringTypeArrayOutput() ExpressRouteCircuitPeeringTypeArrayOutput
	ToExpressRouteCircuitPeeringTypeArrayOutputWithContext(context.Context) ExpressRouteCircuitPeeringTypeArrayOutput
}

type ExpressRouteCircuitPeeringTypeArray []ExpressRouteCircuitPeeringTypeInput

func (ExpressRouteCircuitPeeringTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressRouteCircuitPeeringType)(nil)).Elem()
}

func (i ExpressRouteCircuitPeeringTypeArray) ToExpressRouteCircuitPeeringTypeArrayOutput() ExpressRouteCircuitPeeringTypeArrayOutput {
	return i.ToExpressRouteCircuitPeeringTypeArrayOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitPeeringTypeArray) ToExpressRouteCircuitPeeringTypeArrayOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringTypeArrayOutput)
}

type ExpressRouteCircuitPeeringTypeOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringType)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringTypeOutput) ToExpressRouteCircuitPeeringTypeOutput() ExpressRouteCircuitPeeringTypeOutput {
	return o
}

func (o ExpressRouteCircuitPeeringTypeOutput) ToExpressRouteCircuitPeeringTypeOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringTypeOutput {
	return o
}

func (o ExpressRouteCircuitPeeringTypeOutput) AzureASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *int { return v.AzureASN }).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) GatewayManagerEtag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.GatewayManagerEtag }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) Ipv6PeeringConfig() Ipv6ExpressRouteCircuitPeeringConfigPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *Ipv6ExpressRouteCircuitPeeringConfig {
		return v.Ipv6PeeringConfig
	}).(Ipv6ExpressRouteCircuitPeeringConfigPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *ExpressRouteCircuitPeeringConfig {
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) PeerASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *int { return v.PeerASN }).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) PeeringType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.PeeringType }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) PrimaryAzurePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.PrimaryAzurePort }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.PrimaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) RouteFilter() RouteFilterTypePtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *RouteFilterType { return v.RouteFilter }).(RouteFilterTypePtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) SecondaryAzurePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.SecondaryAzurePort }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.SecondaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) Stats() ExpressRouteCircuitStatsPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *ExpressRouteCircuitStats { return v.Stats }).(ExpressRouteCircuitStatsPtrOutput)
}

func (o ExpressRouteCircuitPeeringTypeOutput) VlanId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringType) *int { return v.VlanId }).(pulumi.IntPtrOutput)
}

type ExpressRouteCircuitPeeringTypeArrayOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressRouteCircuitPeeringType)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringTypeArrayOutput) ToExpressRouteCircuitPeeringTypeArrayOutput() ExpressRouteCircuitPeeringTypeArrayOutput {
	return o
}

func (o ExpressRouteCircuitPeeringTypeArrayOutput) ToExpressRouteCircuitPeeringTypeArrayOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringTypeArrayOutput {
	return o
}

func (o ExpressRouteCircuitPeeringTypeArrayOutput) Index(i pulumi.IntInput) ExpressRouteCircuitPeeringTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExpressRouteCircuitPeeringType {
		return vs[0].([]ExpressRouteCircuitPeeringType)[vs[1].(int)]
	}).(ExpressRouteCircuitPeeringTypeOutput)
}

type ExpressRouteCircuitPeeringConfig struct {
	AdvertisedPublicPrefixes      []string `pulumi:"advertisedPublicPrefixes"`
	AdvertisedPublicPrefixesState *string  `pulumi:"advertisedPublicPrefixesState"`
	CustomerASN                   *int     `pulumi:"customerASN"`
	RoutingRegistryName           *string  `pulumi:"routingRegistryName"`
}





type ExpressRouteCircuitPeeringConfigInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringConfigOutput() ExpressRouteCircuitPeeringConfigOutput
	ToExpressRouteCircuitPeeringConfigOutputWithContext(context.Context) ExpressRouteCircuitPeeringConfigOutput
}

type ExpressRouteCircuitPeeringConfigArgs struct {
	AdvertisedPublicPrefixes      pulumi.StringArrayInput `pulumi:"advertisedPublicPrefixes"`
	AdvertisedPublicPrefixesState pulumi.StringPtrInput   `pulumi:"advertisedPublicPrefixesState"`
	CustomerASN                   pulumi.IntPtrInput      `pulumi:"customerASN"`
	RoutingRegistryName           pulumi.StringPtrInput   `pulumi:"routingRegistryName"`
}

func (ExpressRouteCircuitPeeringConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringConfig)(nil)).Elem()
}

func (i ExpressRouteCircuitPeeringConfigArgs) ToExpressRouteCircuitPeeringConfigOutput() ExpressRouteCircuitPeeringConfigOutput {
	return i.ToExpressRouteCircuitPeeringConfigOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitPeeringConfigArgs) ToExpressRouteCircuitPeeringConfigOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringConfigOutput)
}

func (i ExpressRouteCircuitPeeringConfigArgs) ToExpressRouteCircuitPeeringConfigPtrOutput() ExpressRouteCircuitPeeringConfigPtrOutput {
	return i.ToExpressRouteCircuitPeeringConfigPtrOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitPeeringConfigArgs) ToExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringConfigOutput).ToExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx)
}









type ExpressRouteCircuitPeeringConfigPtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringConfigPtrOutput() ExpressRouteCircuitPeeringConfigPtrOutput
	ToExpressRouteCircuitPeeringConfigPtrOutputWithContext(context.Context) ExpressRouteCircuitPeeringConfigPtrOutput
}

type expressRouteCircuitPeeringConfigPtrType ExpressRouteCircuitPeeringConfigArgs

func ExpressRouteCircuitPeeringConfigPtr(v *ExpressRouteCircuitPeeringConfigArgs) ExpressRouteCircuitPeeringConfigPtrInput {
	return (*expressRouteCircuitPeeringConfigPtrType)(v)
}

func (*expressRouteCircuitPeeringConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitPeeringConfig)(nil)).Elem()
}

func (i *expressRouteCircuitPeeringConfigPtrType) ToExpressRouteCircuitPeeringConfigPtrOutput() ExpressRouteCircuitPeeringConfigPtrOutput {
	return i.ToExpressRouteCircuitPeeringConfigPtrOutputWithContext(context.Background())
}

func (i *expressRouteCircuitPeeringConfigPtrType) ToExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringConfigPtrOutput)
}

type ExpressRouteCircuitPeeringConfigOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringConfig)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringConfigOutput) ToExpressRouteCircuitPeeringConfigOutput() ExpressRouteCircuitPeeringConfigOutput {
	return o
}

func (o ExpressRouteCircuitPeeringConfigOutput) ToExpressRouteCircuitPeeringConfigOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigOutput {
	return o
}

func (o ExpressRouteCircuitPeeringConfigOutput) ToExpressRouteCircuitPeeringConfigPtrOutput() ExpressRouteCircuitPeeringConfigPtrOutput {
	return o.ToExpressRouteCircuitPeeringConfigPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitPeeringConfigOutput) ToExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuitPeeringConfig) *ExpressRouteCircuitPeeringConfig {
		return &v
	}).(ExpressRouteCircuitPeeringConfigPtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigOutput) AdvertisedPublicPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfig) []string { return v.AdvertisedPublicPrefixes }).(pulumi.StringArrayOutput)
}

func (o ExpressRouteCircuitPeeringConfigOutput) AdvertisedPublicPrefixesState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfig) *string { return v.AdvertisedPublicPrefixesState }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigOutput) CustomerASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfig) *int { return v.CustomerASN }).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigOutput) RoutingRegistryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfig) *string { return v.RoutingRegistryName }).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitPeeringConfigPtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitPeeringConfig)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringConfigPtrOutput) ToExpressRouteCircuitPeeringConfigPtrOutput() ExpressRouteCircuitPeeringConfigPtrOutput {
	return o
}

func (o ExpressRouteCircuitPeeringConfigPtrOutput) ToExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigPtrOutput {
	return o
}

func (o ExpressRouteCircuitPeeringConfigPtrOutput) Elem() ExpressRouteCircuitPeeringConfigOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfig) ExpressRouteCircuitPeeringConfig {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitPeeringConfig
		return ret
	}).(ExpressRouteCircuitPeeringConfigOutput)
}

func (o ExpressRouteCircuitPeeringConfigPtrOutput) AdvertisedPublicPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfig) []string {
		if v == nil {
			return nil
		}
		return v.AdvertisedPublicPrefixes
	}).(pulumi.StringArrayOutput)
}

func (o ExpressRouteCircuitPeeringConfigPtrOutput) AdvertisedPublicPrefixesState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfig) *string {
		if v == nil {
			return nil
		}
		return v.AdvertisedPublicPrefixesState
	}).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigPtrOutput) CustomerASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfig) *int {
		if v == nil {
			return nil
		}
		return v.CustomerASN
	}).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigPtrOutput) RoutingRegistryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfig) *string {
		if v == nil {
			return nil
		}
		return v.RoutingRegistryName
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitPeeringConfigResponse struct {
	AdvertisedPublicPrefixes      []string `pulumi:"advertisedPublicPrefixes"`
	AdvertisedPublicPrefixesState *string  `pulumi:"advertisedPublicPrefixesState"`
	CustomerASN                   *int     `pulumi:"customerASN"`
	RoutingRegistryName           *string  `pulumi:"routingRegistryName"`
}

type ExpressRouteCircuitPeeringConfigResponseOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringConfigResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringConfigResponseOutput) ToExpressRouteCircuitPeeringConfigResponseOutput() ExpressRouteCircuitPeeringConfigResponseOutput {
	return o
}

func (o ExpressRouteCircuitPeeringConfigResponseOutput) ToExpressRouteCircuitPeeringConfigResponseOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigResponseOutput {
	return o
}

func (o ExpressRouteCircuitPeeringConfigResponseOutput) AdvertisedPublicPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfigResponse) []string { return v.AdvertisedPublicPrefixes }).(pulumi.StringArrayOutput)
}

func (o ExpressRouteCircuitPeeringConfigResponseOutput) AdvertisedPublicPrefixesState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfigResponse) *string { return v.AdvertisedPublicPrefixesState }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigResponseOutput) CustomerASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfigResponse) *int { return v.CustomerASN }).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigResponseOutput) RoutingRegistryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringConfigResponse) *string { return v.RoutingRegistryName }).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitPeeringConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitPeeringConfigResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringConfigResponsePtrOutput) ToExpressRouteCircuitPeeringConfigResponsePtrOutput() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o
}

func (o ExpressRouteCircuitPeeringConfigResponsePtrOutput) ToExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o
}

func (o ExpressRouteCircuitPeeringConfigResponsePtrOutput) Elem() ExpressRouteCircuitPeeringConfigResponseOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfigResponse) ExpressRouteCircuitPeeringConfigResponse {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitPeeringConfigResponse
		return ret
	}).(ExpressRouteCircuitPeeringConfigResponseOutput)
}

func (o ExpressRouteCircuitPeeringConfigResponsePtrOutput) AdvertisedPublicPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfigResponse) []string {
		if v == nil {
			return nil
		}
		return v.AdvertisedPublicPrefixes
	}).(pulumi.StringArrayOutput)
}

func (o ExpressRouteCircuitPeeringConfigResponsePtrOutput) AdvertisedPublicPrefixesState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdvertisedPublicPrefixesState
	}).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigResponsePtrOutput) CustomerASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.CustomerASN
	}).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitPeeringConfigResponsePtrOutput) RoutingRegistryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitPeeringConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.RoutingRegistryName
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitPeeringResponse struct {
	AzureASN                   *int                                          `pulumi:"azureASN"`
	Etag                       string                                        `pulumi:"etag"`
	GatewayManagerEtag         *string                                       `pulumi:"gatewayManagerEtag"`
	Id                         *string                                       `pulumi:"id"`
	Ipv6PeeringConfig          *Ipv6ExpressRouteCircuitPeeringConfigResponse `pulumi:"ipv6PeeringConfig"`
	LastModifiedBy             *string                                       `pulumi:"lastModifiedBy"`
	MicrosoftPeeringConfig     *ExpressRouteCircuitPeeringConfigResponse     `pulumi:"microsoftPeeringConfig"`
	Name                       *string                                       `pulumi:"name"`
	PeerASN                    *int                                          `pulumi:"peerASN"`
	PeeringType                *string                                       `pulumi:"peeringType"`
	PrimaryAzurePort           *string                                       `pulumi:"primaryAzurePort"`
	PrimaryPeerAddressPrefix   *string                                       `pulumi:"primaryPeerAddressPrefix"`
	ProvisioningState          *string                                       `pulumi:"provisioningState"`
	RouteFilter                *RouteFilterResponse                          `pulumi:"routeFilter"`
	SecondaryAzurePort         *string                                       `pulumi:"secondaryAzurePort"`
	SecondaryPeerAddressPrefix *string                                       `pulumi:"secondaryPeerAddressPrefix"`
	SharedKey                  *string                                       `pulumi:"sharedKey"`
	State                      *string                                       `pulumi:"state"`
	Stats                      *ExpressRouteCircuitStatsResponse             `pulumi:"stats"`
	VlanId                     *int                                          `pulumi:"vlanId"`
}

type ExpressRouteCircuitPeeringResponseOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeeringResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringResponseOutput) ToExpressRouteCircuitPeeringResponseOutput() ExpressRouteCircuitPeeringResponseOutput {
	return o
}

func (o ExpressRouteCircuitPeeringResponseOutput) ToExpressRouteCircuitPeeringResponseOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringResponseOutput {
	return o
}

func (o ExpressRouteCircuitPeeringResponseOutput) AzureASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *int { return v.AzureASN }).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) GatewayManagerEtag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.GatewayManagerEtag }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) Ipv6PeeringConfig() Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *Ipv6ExpressRouteCircuitPeeringConfigResponse {
		return v.Ipv6PeeringConfig
	}).(Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *ExpressRouteCircuitPeeringConfigResponse {
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) PeerASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *int { return v.PeerASN }).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) PeeringType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.PeeringType }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) PrimaryAzurePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.PrimaryAzurePort }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.PrimaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) RouteFilter() RouteFilterResponsePtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *RouteFilterResponse { return v.RouteFilter }).(RouteFilterResponsePtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) SecondaryAzurePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.SecondaryAzurePort }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.SecondaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) Stats() ExpressRouteCircuitStatsResponsePtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *ExpressRouteCircuitStatsResponse { return v.Stats }).(ExpressRouteCircuitStatsResponsePtrOutput)
}

func (o ExpressRouteCircuitPeeringResponseOutput) VlanId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitPeeringResponse) *int { return v.VlanId }).(pulumi.IntPtrOutput)
}

type ExpressRouteCircuitPeeringResponseArrayOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressRouteCircuitPeeringResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitPeeringResponseArrayOutput) ToExpressRouteCircuitPeeringResponseArrayOutput() ExpressRouteCircuitPeeringResponseArrayOutput {
	return o
}

func (o ExpressRouteCircuitPeeringResponseArrayOutput) ToExpressRouteCircuitPeeringResponseArrayOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringResponseArrayOutput {
	return o
}

func (o ExpressRouteCircuitPeeringResponseArrayOutput) Index(i pulumi.IntInput) ExpressRouteCircuitPeeringResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExpressRouteCircuitPeeringResponse {
		return vs[0].([]ExpressRouteCircuitPeeringResponse)[vs[1].(int)]
	}).(ExpressRouteCircuitPeeringResponseOutput)
}

type ExpressRouteCircuitServiceProviderProperties struct {
	BandwidthInMbps     *int    `pulumi:"bandwidthInMbps"`
	PeeringLocation     *string `pulumi:"peeringLocation"`
	ServiceProviderName *string `pulumi:"serviceProviderName"`
}





type ExpressRouteCircuitServiceProviderPropertiesInput interface {
	pulumi.Input

	ToExpressRouteCircuitServiceProviderPropertiesOutput() ExpressRouteCircuitServiceProviderPropertiesOutput
	ToExpressRouteCircuitServiceProviderPropertiesOutputWithContext(context.Context) ExpressRouteCircuitServiceProviderPropertiesOutput
}

type ExpressRouteCircuitServiceProviderPropertiesArgs struct {
	BandwidthInMbps     pulumi.IntPtrInput    `pulumi:"bandwidthInMbps"`
	PeeringLocation     pulumi.StringPtrInput `pulumi:"peeringLocation"`
	ServiceProviderName pulumi.StringPtrInput `pulumi:"serviceProviderName"`
}

func (ExpressRouteCircuitServiceProviderPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitServiceProviderProperties)(nil)).Elem()
}

func (i ExpressRouteCircuitServiceProviderPropertiesArgs) ToExpressRouteCircuitServiceProviderPropertiesOutput() ExpressRouteCircuitServiceProviderPropertiesOutput {
	return i.ToExpressRouteCircuitServiceProviderPropertiesOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitServiceProviderPropertiesArgs) ToExpressRouteCircuitServiceProviderPropertiesOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitServiceProviderPropertiesOutput)
}

func (i ExpressRouteCircuitServiceProviderPropertiesArgs) ToExpressRouteCircuitServiceProviderPropertiesPtrOutput() ExpressRouteCircuitServiceProviderPropertiesPtrOutput {
	return i.ToExpressRouteCircuitServiceProviderPropertiesPtrOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitServiceProviderPropertiesArgs) ToExpressRouteCircuitServiceProviderPropertiesPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitServiceProviderPropertiesOutput).ToExpressRouteCircuitServiceProviderPropertiesPtrOutputWithContext(ctx)
}









type ExpressRouteCircuitServiceProviderPropertiesPtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitServiceProviderPropertiesPtrOutput() ExpressRouteCircuitServiceProviderPropertiesPtrOutput
	ToExpressRouteCircuitServiceProviderPropertiesPtrOutputWithContext(context.Context) ExpressRouteCircuitServiceProviderPropertiesPtrOutput
}

type expressRouteCircuitServiceProviderPropertiesPtrType ExpressRouteCircuitServiceProviderPropertiesArgs

func ExpressRouteCircuitServiceProviderPropertiesPtr(v *ExpressRouteCircuitServiceProviderPropertiesArgs) ExpressRouteCircuitServiceProviderPropertiesPtrInput {
	return (*expressRouteCircuitServiceProviderPropertiesPtrType)(v)
}

func (*expressRouteCircuitServiceProviderPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitServiceProviderProperties)(nil)).Elem()
}

func (i *expressRouteCircuitServiceProviderPropertiesPtrType) ToExpressRouteCircuitServiceProviderPropertiesPtrOutput() ExpressRouteCircuitServiceProviderPropertiesPtrOutput {
	return i.ToExpressRouteCircuitServiceProviderPropertiesPtrOutputWithContext(context.Background())
}

func (i *expressRouteCircuitServiceProviderPropertiesPtrType) ToExpressRouteCircuitServiceProviderPropertiesPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitServiceProviderPropertiesPtrOutput)
}

type ExpressRouteCircuitServiceProviderPropertiesOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitServiceProviderPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitServiceProviderProperties)(nil)).Elem()
}

func (o ExpressRouteCircuitServiceProviderPropertiesOutput) ToExpressRouteCircuitServiceProviderPropertiesOutput() ExpressRouteCircuitServiceProviderPropertiesOutput {
	return o
}

func (o ExpressRouteCircuitServiceProviderPropertiesOutput) ToExpressRouteCircuitServiceProviderPropertiesOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesOutput {
	return o
}

func (o ExpressRouteCircuitServiceProviderPropertiesOutput) ToExpressRouteCircuitServiceProviderPropertiesPtrOutput() ExpressRouteCircuitServiceProviderPropertiesPtrOutput {
	return o.ToExpressRouteCircuitServiceProviderPropertiesPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitServiceProviderPropertiesOutput) ToExpressRouteCircuitServiceProviderPropertiesPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuitServiceProviderProperties) *ExpressRouteCircuitServiceProviderProperties {
		return &v
	}).(ExpressRouteCircuitServiceProviderPropertiesPtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesOutput) BandwidthInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitServiceProviderProperties) *int { return v.BandwidthInMbps }).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesOutput) PeeringLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitServiceProviderProperties) *string { return v.PeeringLocation }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesOutput) ServiceProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitServiceProviderProperties) *string { return v.ServiceProviderName }).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitServiceProviderPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitServiceProviderPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitServiceProviderProperties)(nil)).Elem()
}

func (o ExpressRouteCircuitServiceProviderPropertiesPtrOutput) ToExpressRouteCircuitServiceProviderPropertiesPtrOutput() ExpressRouteCircuitServiceProviderPropertiesPtrOutput {
	return o
}

func (o ExpressRouteCircuitServiceProviderPropertiesPtrOutput) ToExpressRouteCircuitServiceProviderPropertiesPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesPtrOutput {
	return o
}

func (o ExpressRouteCircuitServiceProviderPropertiesPtrOutput) Elem() ExpressRouteCircuitServiceProviderPropertiesOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitServiceProviderProperties) ExpressRouteCircuitServiceProviderProperties {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitServiceProviderProperties
		return ret
	}).(ExpressRouteCircuitServiceProviderPropertiesOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesPtrOutput) BandwidthInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitServiceProviderProperties) *int {
		if v == nil {
			return nil
		}
		return v.BandwidthInMbps
	}).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesPtrOutput) PeeringLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitServiceProviderProperties) *string {
		if v == nil {
			return nil
		}
		return v.PeeringLocation
	}).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesPtrOutput) ServiceProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitServiceProviderProperties) *string {
		if v == nil {
			return nil
		}
		return v.ServiceProviderName
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitServiceProviderPropertiesResponse struct {
	BandwidthInMbps     *int    `pulumi:"bandwidthInMbps"`
	PeeringLocation     *string `pulumi:"peeringLocation"`
	ServiceProviderName *string `pulumi:"serviceProviderName"`
}

type ExpressRouteCircuitServiceProviderPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitServiceProviderPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitServiceProviderPropertiesResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponseOutput) ToExpressRouteCircuitServiceProviderPropertiesResponseOutput() ExpressRouteCircuitServiceProviderPropertiesResponseOutput {
	return o
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponseOutput) ToExpressRouteCircuitServiceProviderPropertiesResponseOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesResponseOutput {
	return o
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponseOutput) BandwidthInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitServiceProviderPropertiesResponse) *int { return v.BandwidthInMbps }).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponseOutput) PeeringLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitServiceProviderPropertiesResponse) *string { return v.PeeringLocation }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponseOutput) ServiceProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitServiceProviderPropertiesResponse) *string { return v.ServiceProviderName }).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitServiceProviderPropertiesResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput) ToExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput() ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput {
	return o
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput) ToExpressRouteCircuitServiceProviderPropertiesResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput {
	return o
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput) Elem() ExpressRouteCircuitServiceProviderPropertiesResponseOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitServiceProviderPropertiesResponse) ExpressRouteCircuitServiceProviderPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitServiceProviderPropertiesResponse
		return ret
	}).(ExpressRouteCircuitServiceProviderPropertiesResponseOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput) BandwidthInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitServiceProviderPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.BandwidthInMbps
	}).(pulumi.IntPtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput) PeeringLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitServiceProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PeeringLocation
	}).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput) ServiceProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitServiceProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceProviderName
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitSku struct {
	Family *string `pulumi:"family"`
	Name   *string `pulumi:"name"`
	Tier   *string `pulumi:"tier"`
}





type ExpressRouteCircuitSkuInput interface {
	pulumi.Input

	ToExpressRouteCircuitSkuOutput() ExpressRouteCircuitSkuOutput
	ToExpressRouteCircuitSkuOutputWithContext(context.Context) ExpressRouteCircuitSkuOutput
}

type ExpressRouteCircuitSkuArgs struct {
	Family pulumi.StringPtrInput `pulumi:"family"`
	Name   pulumi.StringPtrInput `pulumi:"name"`
	Tier   pulumi.StringPtrInput `pulumi:"tier"`
}

func (ExpressRouteCircuitSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitSku)(nil)).Elem()
}

func (i ExpressRouteCircuitSkuArgs) ToExpressRouteCircuitSkuOutput() ExpressRouteCircuitSkuOutput {
	return i.ToExpressRouteCircuitSkuOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitSkuArgs) ToExpressRouteCircuitSkuOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitSkuOutput)
}

func (i ExpressRouteCircuitSkuArgs) ToExpressRouteCircuitSkuPtrOutput() ExpressRouteCircuitSkuPtrOutput {
	return i.ToExpressRouteCircuitSkuPtrOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitSkuArgs) ToExpressRouteCircuitSkuPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitSkuOutput).ToExpressRouteCircuitSkuPtrOutputWithContext(ctx)
}









type ExpressRouteCircuitSkuPtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitSkuPtrOutput() ExpressRouteCircuitSkuPtrOutput
	ToExpressRouteCircuitSkuPtrOutputWithContext(context.Context) ExpressRouteCircuitSkuPtrOutput
}

type expressRouteCircuitSkuPtrType ExpressRouteCircuitSkuArgs

func ExpressRouteCircuitSkuPtr(v *ExpressRouteCircuitSkuArgs) ExpressRouteCircuitSkuPtrInput {
	return (*expressRouteCircuitSkuPtrType)(v)
}

func (*expressRouteCircuitSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitSku)(nil)).Elem()
}

func (i *expressRouteCircuitSkuPtrType) ToExpressRouteCircuitSkuPtrOutput() ExpressRouteCircuitSkuPtrOutput {
	return i.ToExpressRouteCircuitSkuPtrOutputWithContext(context.Background())
}

func (i *expressRouteCircuitSkuPtrType) ToExpressRouteCircuitSkuPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitSkuPtrOutput)
}

type ExpressRouteCircuitSkuOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitSku)(nil)).Elem()
}

func (o ExpressRouteCircuitSkuOutput) ToExpressRouteCircuitSkuOutput() ExpressRouteCircuitSkuOutput {
	return o
}

func (o ExpressRouteCircuitSkuOutput) ToExpressRouteCircuitSkuOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuOutput {
	return o
}

func (o ExpressRouteCircuitSkuOutput) ToExpressRouteCircuitSkuPtrOutput() ExpressRouteCircuitSkuPtrOutput {
	return o.ToExpressRouteCircuitSkuPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitSkuOutput) ToExpressRouteCircuitSkuPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuitSku) *ExpressRouteCircuitSku {
		return &v
	}).(ExpressRouteCircuitSkuPtrOutput)
}

func (o ExpressRouteCircuitSkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitSku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitSkuPtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitSku)(nil)).Elem()
}

func (o ExpressRouteCircuitSkuPtrOutput) ToExpressRouteCircuitSkuPtrOutput() ExpressRouteCircuitSkuPtrOutput {
	return o
}

func (o ExpressRouteCircuitSkuPtrOutput) ToExpressRouteCircuitSkuPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuPtrOutput {
	return o
}

func (o ExpressRouteCircuitSkuPtrOutput) Elem() ExpressRouteCircuitSkuOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitSku) ExpressRouteCircuitSku {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitSku
		return ret
	}).(ExpressRouteCircuitSkuOutput)
}

func (o ExpressRouteCircuitSkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitSku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitSku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitSkuResponse struct {
	Family *string `pulumi:"family"`
	Name   *string `pulumi:"name"`
	Tier   *string `pulumi:"tier"`
}

type ExpressRouteCircuitSkuResponseOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitSkuResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitSkuResponseOutput) ToExpressRouteCircuitSkuResponseOutput() ExpressRouteCircuitSkuResponseOutput {
	return o
}

func (o ExpressRouteCircuitSkuResponseOutput) ToExpressRouteCircuitSkuResponseOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuResponseOutput {
	return o
}

func (o ExpressRouteCircuitSkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitSkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitSkuResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitSkuResponsePtrOutput) ToExpressRouteCircuitSkuResponsePtrOutput() ExpressRouteCircuitSkuResponsePtrOutput {
	return o
}

func (o ExpressRouteCircuitSkuResponsePtrOutput) ToExpressRouteCircuitSkuResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitSkuResponsePtrOutput {
	return o
}

func (o ExpressRouteCircuitSkuResponsePtrOutput) Elem() ExpressRouteCircuitSkuResponseOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitSkuResponse) ExpressRouteCircuitSkuResponse {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitSkuResponse
		return ret
	}).(ExpressRouteCircuitSkuResponseOutput)
}

func (o ExpressRouteCircuitSkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ExpressRouteCircuitStats struct {
	PrimarybytesIn    *float64 `pulumi:"primarybytesIn"`
	PrimarybytesOut   *float64 `pulumi:"primarybytesOut"`
	SecondarybytesIn  *float64 `pulumi:"secondarybytesIn"`
	SecondarybytesOut *float64 `pulumi:"secondarybytesOut"`
}





type ExpressRouteCircuitStatsInput interface {
	pulumi.Input

	ToExpressRouteCircuitStatsOutput() ExpressRouteCircuitStatsOutput
	ToExpressRouteCircuitStatsOutputWithContext(context.Context) ExpressRouteCircuitStatsOutput
}

type ExpressRouteCircuitStatsArgs struct {
	PrimarybytesIn    pulumi.Float64PtrInput `pulumi:"primarybytesIn"`
	PrimarybytesOut   pulumi.Float64PtrInput `pulumi:"primarybytesOut"`
	SecondarybytesIn  pulumi.Float64PtrInput `pulumi:"secondarybytesIn"`
	SecondarybytesOut pulumi.Float64PtrInput `pulumi:"secondarybytesOut"`
}

func (ExpressRouteCircuitStatsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitStats)(nil)).Elem()
}

func (i ExpressRouteCircuitStatsArgs) ToExpressRouteCircuitStatsOutput() ExpressRouteCircuitStatsOutput {
	return i.ToExpressRouteCircuitStatsOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitStatsArgs) ToExpressRouteCircuitStatsOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitStatsOutput)
}

func (i ExpressRouteCircuitStatsArgs) ToExpressRouteCircuitStatsPtrOutput() ExpressRouteCircuitStatsPtrOutput {
	return i.ToExpressRouteCircuitStatsPtrOutputWithContext(context.Background())
}

func (i ExpressRouteCircuitStatsArgs) ToExpressRouteCircuitStatsPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitStatsOutput).ToExpressRouteCircuitStatsPtrOutputWithContext(ctx)
}









type ExpressRouteCircuitStatsPtrInput interface {
	pulumi.Input

	ToExpressRouteCircuitStatsPtrOutput() ExpressRouteCircuitStatsPtrOutput
	ToExpressRouteCircuitStatsPtrOutputWithContext(context.Context) ExpressRouteCircuitStatsPtrOutput
}

type expressRouteCircuitStatsPtrType ExpressRouteCircuitStatsArgs

func ExpressRouteCircuitStatsPtr(v *ExpressRouteCircuitStatsArgs) ExpressRouteCircuitStatsPtrInput {
	return (*expressRouteCircuitStatsPtrType)(v)
}

func (*expressRouteCircuitStatsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitStats)(nil)).Elem()
}

func (i *expressRouteCircuitStatsPtrType) ToExpressRouteCircuitStatsPtrOutput() ExpressRouteCircuitStatsPtrOutput {
	return i.ToExpressRouteCircuitStatsPtrOutputWithContext(context.Background())
}

func (i *expressRouteCircuitStatsPtrType) ToExpressRouteCircuitStatsPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitStatsPtrOutput)
}

type ExpressRouteCircuitStatsOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitStatsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitStats)(nil)).Elem()
}

func (o ExpressRouteCircuitStatsOutput) ToExpressRouteCircuitStatsOutput() ExpressRouteCircuitStatsOutput {
	return o
}

func (o ExpressRouteCircuitStatsOutput) ToExpressRouteCircuitStatsOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsOutput {
	return o
}

func (o ExpressRouteCircuitStatsOutput) ToExpressRouteCircuitStatsPtrOutput() ExpressRouteCircuitStatsPtrOutput {
	return o.ToExpressRouteCircuitStatsPtrOutputWithContext(context.Background())
}

func (o ExpressRouteCircuitStatsOutput) ToExpressRouteCircuitStatsPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExpressRouteCircuitStats) *ExpressRouteCircuitStats {
		return &v
	}).(ExpressRouteCircuitStatsPtrOutput)
}

func (o ExpressRouteCircuitStatsOutput) PrimarybytesIn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStats) *float64 { return v.PrimarybytesIn }).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsOutput) PrimarybytesOut() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStats) *float64 { return v.PrimarybytesOut }).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsOutput) SecondarybytesIn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStats) *float64 { return v.SecondarybytesIn }).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsOutput) SecondarybytesOut() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStats) *float64 { return v.SecondarybytesOut }).(pulumi.Float64PtrOutput)
}

type ExpressRouteCircuitStatsPtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitStatsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitStats)(nil)).Elem()
}

func (o ExpressRouteCircuitStatsPtrOutput) ToExpressRouteCircuitStatsPtrOutput() ExpressRouteCircuitStatsPtrOutput {
	return o
}

func (o ExpressRouteCircuitStatsPtrOutput) ToExpressRouteCircuitStatsPtrOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsPtrOutput {
	return o
}

func (o ExpressRouteCircuitStatsPtrOutput) Elem() ExpressRouteCircuitStatsOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStats) ExpressRouteCircuitStats {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitStats
		return ret
	}).(ExpressRouteCircuitStatsOutput)
}

func (o ExpressRouteCircuitStatsPtrOutput) PrimarybytesIn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStats) *float64 {
		if v == nil {
			return nil
		}
		return v.PrimarybytesIn
	}).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsPtrOutput) PrimarybytesOut() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStats) *float64 {
		if v == nil {
			return nil
		}
		return v.PrimarybytesOut
	}).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsPtrOutput) SecondarybytesIn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStats) *float64 {
		if v == nil {
			return nil
		}
		return v.SecondarybytesIn
	}).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsPtrOutput) SecondarybytesOut() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStats) *float64 {
		if v == nil {
			return nil
		}
		return v.SecondarybytesOut
	}).(pulumi.Float64PtrOutput)
}

type ExpressRouteCircuitStatsResponse struct {
	PrimarybytesIn    *float64 `pulumi:"primarybytesIn"`
	PrimarybytesOut   *float64 `pulumi:"primarybytesOut"`
	SecondarybytesIn  *float64 `pulumi:"secondarybytesIn"`
	SecondarybytesOut *float64 `pulumi:"secondarybytesOut"`
}

type ExpressRouteCircuitStatsResponseOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitStatsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitStatsResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitStatsResponseOutput) ToExpressRouteCircuitStatsResponseOutput() ExpressRouteCircuitStatsResponseOutput {
	return o
}

func (o ExpressRouteCircuitStatsResponseOutput) ToExpressRouteCircuitStatsResponseOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsResponseOutput {
	return o
}

func (o ExpressRouteCircuitStatsResponseOutput) PrimarybytesIn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStatsResponse) *float64 { return v.PrimarybytesIn }).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsResponseOutput) PrimarybytesOut() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStatsResponse) *float64 { return v.PrimarybytesOut }).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsResponseOutput) SecondarybytesIn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStatsResponse) *float64 { return v.SecondarybytesIn }).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsResponseOutput) SecondarybytesOut() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ExpressRouteCircuitStatsResponse) *float64 { return v.SecondarybytesOut }).(pulumi.Float64PtrOutput)
}

type ExpressRouteCircuitStatsResponsePtrOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitStatsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitStatsResponse)(nil)).Elem()
}

func (o ExpressRouteCircuitStatsResponsePtrOutput) ToExpressRouteCircuitStatsResponsePtrOutput() ExpressRouteCircuitStatsResponsePtrOutput {
	return o
}

func (o ExpressRouteCircuitStatsResponsePtrOutput) ToExpressRouteCircuitStatsResponsePtrOutputWithContext(ctx context.Context) ExpressRouteCircuitStatsResponsePtrOutput {
	return o
}

func (o ExpressRouteCircuitStatsResponsePtrOutput) Elem() ExpressRouteCircuitStatsResponseOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStatsResponse) ExpressRouteCircuitStatsResponse {
		if v != nil {
			return *v
		}
		var ret ExpressRouteCircuitStatsResponse
		return ret
	}).(ExpressRouteCircuitStatsResponseOutput)
}

func (o ExpressRouteCircuitStatsResponsePtrOutput) PrimarybytesIn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStatsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.PrimarybytesIn
	}).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsResponsePtrOutput) PrimarybytesOut() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStatsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.PrimarybytesOut
	}).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsResponsePtrOutput) SecondarybytesIn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStatsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.SecondarybytesIn
	}).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCircuitStatsResponsePtrOutput) SecondarybytesOut() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitStatsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.SecondarybytesOut
	}).(pulumi.Float64PtrOutput)
}

type FrontendIPConfiguration struct {
	Etag                      *string              `pulumi:"etag"`
	Id                        *string              `pulumi:"id"`
	Name                      *string              `pulumi:"name"`
	PrivateIPAddress          *string              `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string              `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         *string              `pulumi:"provisioningState"`
	PublicIPAddress           *PublicIPAddressType `pulumi:"publicIPAddress"`
	Subnet                    *SubnetType          `pulumi:"subnet"`
}





type FrontendIPConfigurationInput interface {
	pulumi.Input

	ToFrontendIPConfigurationOutput() FrontendIPConfigurationOutput
	ToFrontendIPConfigurationOutputWithContext(context.Context) FrontendIPConfigurationOutput
}

type FrontendIPConfigurationArgs struct {
	Etag                      pulumi.StringPtrInput       `pulumi:"etag"`
	Id                        pulumi.StringPtrInput       `pulumi:"id"`
	Name                      pulumi.StringPtrInput       `pulumi:"name"`
	PrivateIPAddress          pulumi.StringPtrInput       `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod pulumi.StringPtrInput       `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         pulumi.StringPtrInput       `pulumi:"provisioningState"`
	PublicIPAddress           PublicIPAddressTypePtrInput `pulumi:"publicIPAddress"`
	Subnet                    SubnetTypePtrInput          `pulumi:"subnet"`
}

func (FrontendIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendIPConfiguration)(nil)).Elem()
}

func (i FrontendIPConfigurationArgs) ToFrontendIPConfigurationOutput() FrontendIPConfigurationOutput {
	return i.ToFrontendIPConfigurationOutputWithContext(context.Background())
}

func (i FrontendIPConfigurationArgs) ToFrontendIPConfigurationOutputWithContext(ctx context.Context) FrontendIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendIPConfigurationOutput)
}





type FrontendIPConfigurationArrayInput interface {
	pulumi.Input

	ToFrontendIPConfigurationArrayOutput() FrontendIPConfigurationArrayOutput
	ToFrontendIPConfigurationArrayOutputWithContext(context.Context) FrontendIPConfigurationArrayOutput
}

type FrontendIPConfigurationArray []FrontendIPConfigurationInput

func (FrontendIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendIPConfiguration)(nil)).Elem()
}

func (i FrontendIPConfigurationArray) ToFrontendIPConfigurationArrayOutput() FrontendIPConfigurationArrayOutput {
	return i.ToFrontendIPConfigurationArrayOutputWithContext(context.Background())
}

func (i FrontendIPConfigurationArray) ToFrontendIPConfigurationArrayOutputWithContext(ctx context.Context) FrontendIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendIPConfigurationArrayOutput)
}

type FrontendIPConfigurationOutput struct{ *pulumi.OutputState }

func (FrontendIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendIPConfiguration)(nil)).Elem()
}

func (o FrontendIPConfigurationOutput) ToFrontendIPConfigurationOutput() FrontendIPConfigurationOutput {
	return o
}

func (o FrontendIPConfigurationOutput) ToFrontendIPConfigurationOutputWithContext(ctx context.Context) FrontendIPConfigurationOutput {
	return o
}

func (o FrontendIPConfigurationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfiguration) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfiguration) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfiguration) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfiguration) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationOutput) PublicIPAddress() PublicIPAddressTypePtrOutput {
	return o.ApplyT(func(v FrontendIPConfiguration) *PublicIPAddressType { return v.PublicIPAddress }).(PublicIPAddressTypePtrOutput)
}

func (o FrontendIPConfigurationOutput) Subnet() SubnetTypePtrOutput {
	return o.ApplyT(func(v FrontendIPConfiguration) *SubnetType { return v.Subnet }).(SubnetTypePtrOutput)
}

type FrontendIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (FrontendIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendIPConfiguration)(nil)).Elem()
}

func (o FrontendIPConfigurationArrayOutput) ToFrontendIPConfigurationArrayOutput() FrontendIPConfigurationArrayOutput {
	return o
}

func (o FrontendIPConfigurationArrayOutput) ToFrontendIPConfigurationArrayOutputWithContext(ctx context.Context) FrontendIPConfigurationArrayOutput {
	return o
}

func (o FrontendIPConfigurationArrayOutput) Index(i pulumi.IntInput) FrontendIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontendIPConfiguration {
		return vs[0].([]FrontendIPConfiguration)[vs[1].(int)]
	}).(FrontendIPConfigurationOutput)
}

type FrontendIPConfigurationResponse struct {
	Etag                      *string                  `pulumi:"etag"`
	Id                        *string                  `pulumi:"id"`
	InboundNatPools           []SubResourceResponse    `pulumi:"inboundNatPools"`
	InboundNatRules           []SubResourceResponse    `pulumi:"inboundNatRules"`
	LoadBalancingRules        []SubResourceResponse    `pulumi:"loadBalancingRules"`
	Name                      *string                  `pulumi:"name"`
	OutboundNatRules          []SubResourceResponse    `pulumi:"outboundNatRules"`
	PrivateIPAddress          *string                  `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string                  `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         *string                  `pulumi:"provisioningState"`
	PublicIPAddress           *PublicIPAddressResponse `pulumi:"publicIPAddress"`
	Subnet                    *SubnetResponse          `pulumi:"subnet"`
}

type FrontendIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (FrontendIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontendIPConfigurationResponse)(nil)).Elem()
}

func (o FrontendIPConfigurationResponseOutput) ToFrontendIPConfigurationResponseOutput() FrontendIPConfigurationResponseOutput {
	return o
}

func (o FrontendIPConfigurationResponseOutput) ToFrontendIPConfigurationResponseOutputWithContext(ctx context.Context) FrontendIPConfigurationResponseOutput {
	return o
}

func (o FrontendIPConfigurationResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationResponseOutput) InboundNatPools() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) []SubResourceResponse { return v.InboundNatPools }).(SubResourceResponseArrayOutput)
}

func (o FrontendIPConfigurationResponseOutput) InboundNatRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) []SubResourceResponse { return v.InboundNatRules }).(SubResourceResponseArrayOutput)
}

func (o FrontendIPConfigurationResponseOutput) LoadBalancingRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) []SubResourceResponse { return v.LoadBalancingRules }).(SubResourceResponseArrayOutput)
}

func (o FrontendIPConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationResponseOutput) OutboundNatRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) []SubResourceResponse { return v.OutboundNatRules }).(SubResourceResponseArrayOutput)
}

func (o FrontendIPConfigurationResponseOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationResponseOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o FrontendIPConfigurationResponseOutput) PublicIPAddress() PublicIPAddressResponsePtrOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) *PublicIPAddressResponse { return v.PublicIPAddress }).(PublicIPAddressResponsePtrOutput)
}

func (o FrontendIPConfigurationResponseOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v FrontendIPConfigurationResponse) *SubnetResponse { return v.Subnet }).(SubnetResponsePtrOutput)
}

type FrontendIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (FrontendIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FrontendIPConfigurationResponse)(nil)).Elem()
}

func (o FrontendIPConfigurationResponseArrayOutput) ToFrontendIPConfigurationResponseArrayOutput() FrontendIPConfigurationResponseArrayOutput {
	return o
}

func (o FrontendIPConfigurationResponseArrayOutput) ToFrontendIPConfigurationResponseArrayOutputWithContext(ctx context.Context) FrontendIPConfigurationResponseArrayOutput {
	return o
}

func (o FrontendIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) FrontendIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FrontendIPConfigurationResponse {
		return vs[0].([]FrontendIPConfigurationResponse)[vs[1].(int)]
	}).(FrontendIPConfigurationResponseOutput)
}

type GatewayRouteResponse struct {
	AsPath       string `pulumi:"asPath"`
	LocalAddress string `pulumi:"localAddress"`
	Network      string `pulumi:"network"`
	NextHop      string `pulumi:"nextHop"`
	Origin       string `pulumi:"origin"`
	SourcePeer   string `pulumi:"sourcePeer"`
	Weight       int    `pulumi:"weight"`
}

type GatewayRouteResponseOutput struct{ *pulumi.OutputState }

func (GatewayRouteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayRouteResponse)(nil)).Elem()
}

func (o GatewayRouteResponseOutput) ToGatewayRouteResponseOutput() GatewayRouteResponseOutput {
	return o
}

func (o GatewayRouteResponseOutput) ToGatewayRouteResponseOutputWithContext(ctx context.Context) GatewayRouteResponseOutput {
	return o
}

func (o GatewayRouteResponseOutput) AsPath() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayRouteResponse) string { return v.AsPath }).(pulumi.StringOutput)
}

func (o GatewayRouteResponseOutput) LocalAddress() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayRouteResponse) string { return v.LocalAddress }).(pulumi.StringOutput)
}

func (o GatewayRouteResponseOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayRouteResponse) string { return v.Network }).(pulumi.StringOutput)
}

func (o GatewayRouteResponseOutput) NextHop() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayRouteResponse) string { return v.NextHop }).(pulumi.StringOutput)
}

func (o GatewayRouteResponseOutput) Origin() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayRouteResponse) string { return v.Origin }).(pulumi.StringOutput)
}

func (o GatewayRouteResponseOutput) SourcePeer() pulumi.StringOutput {
	return o.ApplyT(func(v GatewayRouteResponse) string { return v.SourcePeer }).(pulumi.StringOutput)
}

func (o GatewayRouteResponseOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v GatewayRouteResponse) int { return v.Weight }).(pulumi.IntOutput)
}

type GatewayRouteResponseArrayOutput struct{ *pulumi.OutputState }

func (GatewayRouteResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GatewayRouteResponse)(nil)).Elem()
}

func (o GatewayRouteResponseArrayOutput) ToGatewayRouteResponseArrayOutput() GatewayRouteResponseArrayOutput {
	return o
}

func (o GatewayRouteResponseArrayOutput) ToGatewayRouteResponseArrayOutputWithContext(ctx context.Context) GatewayRouteResponseArrayOutput {
	return o
}

func (o GatewayRouteResponseArrayOutput) Index(i pulumi.IntInput) GatewayRouteResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GatewayRouteResponse {
		return vs[0].([]GatewayRouteResponse)[vs[1].(int)]
	}).(GatewayRouteResponseOutput)
}

type IPConfigurationResponse struct {
	Etag                      *string                  `pulumi:"etag"`
	Id                        *string                  `pulumi:"id"`
	Name                      *string                  `pulumi:"name"`
	PrivateIPAddress          *string                  `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string                  `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         *string                  `pulumi:"provisioningState"`
	PublicIPAddress           *PublicIPAddressResponse `pulumi:"publicIPAddress"`
	Subnet                    *SubnetResponse          `pulumi:"subnet"`
}

type IPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (IPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPConfigurationResponse)(nil)).Elem()
}

func (o IPConfigurationResponseOutput) ToIPConfigurationResponseOutput() IPConfigurationResponseOutput {
	return o
}

func (o IPConfigurationResponseOutput) ToIPConfigurationResponseOutputWithContext(ctx context.Context) IPConfigurationResponseOutput {
	return o
}

func (o IPConfigurationResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponseOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponseOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponseOutput) PublicIPAddress() PublicIPAddressResponsePtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *PublicIPAddressResponse { return v.PublicIPAddress }).(PublicIPAddressResponsePtrOutput)
}

func (o IPConfigurationResponseOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v IPConfigurationResponse) *SubnetResponse { return v.Subnet }).(SubnetResponsePtrOutput)
}

type IPConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (IPConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPConfigurationResponse)(nil)).Elem()
}

func (o IPConfigurationResponsePtrOutput) ToIPConfigurationResponsePtrOutput() IPConfigurationResponsePtrOutput {
	return o
}

func (o IPConfigurationResponsePtrOutput) ToIPConfigurationResponsePtrOutputWithContext(ctx context.Context) IPConfigurationResponsePtrOutput {
	return o
}

func (o IPConfigurationResponsePtrOutput) Elem() IPConfigurationResponseOutput {
	return o.ApplyT(func(v *IPConfigurationResponse) IPConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret IPConfigurationResponse
		return ret
	}).(IPConfigurationResponseOutput)
}

func (o IPConfigurationResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponsePtrOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateIPAddress
	}).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponsePtrOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrivateIPAllocationMethod
	}).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o IPConfigurationResponsePtrOutput) PublicIPAddress() PublicIPAddressResponsePtrOutput {
	return o.ApplyT(func(v *IPConfigurationResponse) *PublicIPAddressResponse {
		if v == nil {
			return nil
		}
		return v.PublicIPAddress
	}).(PublicIPAddressResponsePtrOutput)
}

func (o IPConfigurationResponsePtrOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v *IPConfigurationResponse) *SubnetResponse {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(SubnetResponsePtrOutput)
}

type IPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (IPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPConfigurationResponse)(nil)).Elem()
}

func (o IPConfigurationResponseArrayOutput) ToIPConfigurationResponseArrayOutput() IPConfigurationResponseArrayOutput {
	return o
}

func (o IPConfigurationResponseArrayOutput) ToIPConfigurationResponseArrayOutputWithContext(ctx context.Context) IPConfigurationResponseArrayOutput {
	return o
}

func (o IPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) IPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPConfigurationResponse {
		return vs[0].([]IPConfigurationResponse)[vs[1].(int)]
	}).(IPConfigurationResponseOutput)
}

type InboundNatPool struct {
	BackendPort             int          `pulumi:"backendPort"`
	Etag                    *string      `pulumi:"etag"`
	FrontendIPConfiguration *SubResource `pulumi:"frontendIPConfiguration"`
	FrontendPortRangeEnd    int          `pulumi:"frontendPortRangeEnd"`
	FrontendPortRangeStart  int          `pulumi:"frontendPortRangeStart"`
	Id                      *string      `pulumi:"id"`
	Name                    *string      `pulumi:"name"`
	Protocol                string       `pulumi:"protocol"`
	ProvisioningState       *string      `pulumi:"provisioningState"`
}





type InboundNatPoolInput interface {
	pulumi.Input

	ToInboundNatPoolOutput() InboundNatPoolOutput
	ToInboundNatPoolOutputWithContext(context.Context) InboundNatPoolOutput
}

type InboundNatPoolArgs struct {
	BackendPort             pulumi.IntInput       `pulumi:"backendPort"`
	Etag                    pulumi.StringPtrInput `pulumi:"etag"`
	FrontendIPConfiguration SubResourcePtrInput   `pulumi:"frontendIPConfiguration"`
	FrontendPortRangeEnd    pulumi.IntInput       `pulumi:"frontendPortRangeEnd"`
	FrontendPortRangeStart  pulumi.IntInput       `pulumi:"frontendPortRangeStart"`
	Id                      pulumi.StringPtrInput `pulumi:"id"`
	Name                    pulumi.StringPtrInput `pulumi:"name"`
	Protocol                pulumi.StringInput    `pulumi:"protocol"`
	ProvisioningState       pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (InboundNatPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatPool)(nil)).Elem()
}

func (i InboundNatPoolArgs) ToInboundNatPoolOutput() InboundNatPoolOutput {
	return i.ToInboundNatPoolOutputWithContext(context.Background())
}

func (i InboundNatPoolArgs) ToInboundNatPoolOutputWithContext(ctx context.Context) InboundNatPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatPoolOutput)
}





type InboundNatPoolArrayInput interface {
	pulumi.Input

	ToInboundNatPoolArrayOutput() InboundNatPoolArrayOutput
	ToInboundNatPoolArrayOutputWithContext(context.Context) InboundNatPoolArrayOutput
}

type InboundNatPoolArray []InboundNatPoolInput

func (InboundNatPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatPool)(nil)).Elem()
}

func (i InboundNatPoolArray) ToInboundNatPoolArrayOutput() InboundNatPoolArrayOutput {
	return i.ToInboundNatPoolArrayOutputWithContext(context.Background())
}

func (i InboundNatPoolArray) ToInboundNatPoolArrayOutputWithContext(ctx context.Context) InboundNatPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatPoolArrayOutput)
}

type InboundNatPoolOutput struct{ *pulumi.OutputState }

func (InboundNatPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatPool)(nil)).Elem()
}

func (o InboundNatPoolOutput) ToInboundNatPoolOutput() InboundNatPoolOutput {
	return o
}

func (o InboundNatPoolOutput) ToInboundNatPoolOutputWithContext(ctx context.Context) InboundNatPoolOutput {
	return o
}

func (o InboundNatPoolOutput) BackendPort() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatPool) int { return v.BackendPort }).(pulumi.IntOutput)
}

func (o InboundNatPoolOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatPool) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o InboundNatPoolOutput) FrontendIPConfiguration() SubResourcePtrOutput {
	return o.ApplyT(func(v InboundNatPool) *SubResource { return v.FrontendIPConfiguration }).(SubResourcePtrOutput)
}

func (o InboundNatPoolOutput) FrontendPortRangeEnd() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatPool) int { return v.FrontendPortRangeEnd }).(pulumi.IntOutput)
}

func (o InboundNatPoolOutput) FrontendPortRangeStart() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatPool) int { return v.FrontendPortRangeStart }).(pulumi.IntOutput)
}

func (o InboundNatPoolOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatPool) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o InboundNatPoolOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatPool) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o InboundNatPoolOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v InboundNatPool) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o InboundNatPoolOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatPool) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type InboundNatPoolArrayOutput struct{ *pulumi.OutputState }

func (InboundNatPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatPool)(nil)).Elem()
}

func (o InboundNatPoolArrayOutput) ToInboundNatPoolArrayOutput() InboundNatPoolArrayOutput {
	return o
}

func (o InboundNatPoolArrayOutput) ToInboundNatPoolArrayOutputWithContext(ctx context.Context) InboundNatPoolArrayOutput {
	return o
}

func (o InboundNatPoolArrayOutput) Index(i pulumi.IntInput) InboundNatPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundNatPool {
		return vs[0].([]InboundNatPool)[vs[1].(int)]
	}).(InboundNatPoolOutput)
}

type InboundNatPoolResponse struct {
	BackendPort             int                  `pulumi:"backendPort"`
	Etag                    *string              `pulumi:"etag"`
	FrontendIPConfiguration *SubResourceResponse `pulumi:"frontendIPConfiguration"`
	FrontendPortRangeEnd    int                  `pulumi:"frontendPortRangeEnd"`
	FrontendPortRangeStart  int                  `pulumi:"frontendPortRangeStart"`
	Id                      *string              `pulumi:"id"`
	Name                    *string              `pulumi:"name"`
	Protocol                string               `pulumi:"protocol"`
	ProvisioningState       *string              `pulumi:"provisioningState"`
}

type InboundNatPoolResponseOutput struct{ *pulumi.OutputState }

func (InboundNatPoolResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatPoolResponse)(nil)).Elem()
}

func (o InboundNatPoolResponseOutput) ToInboundNatPoolResponseOutput() InboundNatPoolResponseOutput {
	return o
}

func (o InboundNatPoolResponseOutput) ToInboundNatPoolResponseOutputWithContext(ctx context.Context) InboundNatPoolResponseOutput {
	return o
}

func (o InboundNatPoolResponseOutput) BackendPort() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) int { return v.BackendPort }).(pulumi.IntOutput)
}

func (o InboundNatPoolResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o InboundNatPoolResponseOutput) FrontendIPConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) *SubResourceResponse { return v.FrontendIPConfiguration }).(SubResourceResponsePtrOutput)
}

func (o InboundNatPoolResponseOutput) FrontendPortRangeEnd() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) int { return v.FrontendPortRangeEnd }).(pulumi.IntOutput)
}

func (o InboundNatPoolResponseOutput) FrontendPortRangeStart() pulumi.IntOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) int { return v.FrontendPortRangeStart }).(pulumi.IntOutput)
}

func (o InboundNatPoolResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o InboundNatPoolResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o InboundNatPoolResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o InboundNatPoolResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatPoolResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type InboundNatPoolResponseArrayOutput struct{ *pulumi.OutputState }

func (InboundNatPoolResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatPoolResponse)(nil)).Elem()
}

func (o InboundNatPoolResponseArrayOutput) ToInboundNatPoolResponseArrayOutput() InboundNatPoolResponseArrayOutput {
	return o
}

func (o InboundNatPoolResponseArrayOutput) ToInboundNatPoolResponseArrayOutputWithContext(ctx context.Context) InboundNatPoolResponseArrayOutput {
	return o
}

func (o InboundNatPoolResponseArrayOutput) Index(i pulumi.IntInput) InboundNatPoolResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundNatPoolResponse {
		return vs[0].([]InboundNatPoolResponse)[vs[1].(int)]
	}).(InboundNatPoolResponseOutput)
}

type InboundNatRule struct {
	BackendPort             *int         `pulumi:"backendPort"`
	EnableFloatingIP        *bool        `pulumi:"enableFloatingIP"`
	Etag                    *string      `pulumi:"etag"`
	FrontendIPConfiguration *SubResource `pulumi:"frontendIPConfiguration"`
	FrontendPort            *int         `pulumi:"frontendPort"`
	Id                      *string      `pulumi:"id"`
	IdleTimeoutInMinutes    *int         `pulumi:"idleTimeoutInMinutes"`
	Name                    *string      `pulumi:"name"`
	Protocol                *string      `pulumi:"protocol"`
	ProvisioningState       *string      `pulumi:"provisioningState"`
}





type InboundNatRuleInput interface {
	pulumi.Input

	ToInboundNatRuleOutput() InboundNatRuleOutput
	ToInboundNatRuleOutputWithContext(context.Context) InboundNatRuleOutput
}

type InboundNatRuleArgs struct {
	BackendPort             pulumi.IntPtrInput    `pulumi:"backendPort"`
	EnableFloatingIP        pulumi.BoolPtrInput   `pulumi:"enableFloatingIP"`
	Etag                    pulumi.StringPtrInput `pulumi:"etag"`
	FrontendIPConfiguration SubResourcePtrInput   `pulumi:"frontendIPConfiguration"`
	FrontendPort            pulumi.IntPtrInput    `pulumi:"frontendPort"`
	Id                      pulumi.StringPtrInput `pulumi:"id"`
	IdleTimeoutInMinutes    pulumi.IntPtrInput    `pulumi:"idleTimeoutInMinutes"`
	Name                    pulumi.StringPtrInput `pulumi:"name"`
	Protocol                pulumi.StringPtrInput `pulumi:"protocol"`
	ProvisioningState       pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (InboundNatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatRule)(nil)).Elem()
}

func (i InboundNatRuleArgs) ToInboundNatRuleOutput() InboundNatRuleOutput {
	return i.ToInboundNatRuleOutputWithContext(context.Background())
}

func (i InboundNatRuleArgs) ToInboundNatRuleOutputWithContext(ctx context.Context) InboundNatRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatRuleOutput)
}





type InboundNatRuleArrayInput interface {
	pulumi.Input

	ToInboundNatRuleArrayOutput() InboundNatRuleArrayOutput
	ToInboundNatRuleArrayOutputWithContext(context.Context) InboundNatRuleArrayOutput
}

type InboundNatRuleArray []InboundNatRuleInput

func (InboundNatRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatRule)(nil)).Elem()
}

func (i InboundNatRuleArray) ToInboundNatRuleArrayOutput() InboundNatRuleArrayOutput {
	return i.ToInboundNatRuleArrayOutputWithContext(context.Background())
}

func (i InboundNatRuleArray) ToInboundNatRuleArrayOutputWithContext(ctx context.Context) InboundNatRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatRuleArrayOutput)
}

type InboundNatRuleOutput struct{ *pulumi.OutputState }

func (InboundNatRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatRule)(nil)).Elem()
}

func (o InboundNatRuleOutput) ToInboundNatRuleOutput() InboundNatRuleOutput {
	return o
}

func (o InboundNatRuleOutput) ToInboundNatRuleOutputWithContext(ctx context.Context) InboundNatRuleOutput {
	return o
}

func (o InboundNatRuleOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRule) *int { return v.BackendPort }).(pulumi.IntPtrOutput)
}

func (o InboundNatRuleOutput) EnableFloatingIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v InboundNatRule) *bool { return v.EnableFloatingIP }).(pulumi.BoolPtrOutput)
}

func (o InboundNatRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRule) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleOutput) FrontendIPConfiguration() SubResourcePtrOutput {
	return o.ApplyT(func(v InboundNatRule) *SubResource { return v.FrontendIPConfiguration }).(SubResourcePtrOutput)
}

func (o InboundNatRuleOutput) FrontendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRule) *int { return v.FrontendPort }).(pulumi.IntPtrOutput)
}

func (o InboundNatRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRule) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o InboundNatRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRule) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRule) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type InboundNatRuleArrayOutput struct{ *pulumi.OutputState }

func (InboundNatRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatRule)(nil)).Elem()
}

func (o InboundNatRuleArrayOutput) ToInboundNatRuleArrayOutput() InboundNatRuleArrayOutput {
	return o
}

func (o InboundNatRuleArrayOutput) ToInboundNatRuleArrayOutputWithContext(ctx context.Context) InboundNatRuleArrayOutput {
	return o
}

func (o InboundNatRuleArrayOutput) Index(i pulumi.IntInput) InboundNatRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundNatRule {
		return vs[0].([]InboundNatRule)[vs[1].(int)]
	}).(InboundNatRuleOutput)
}

type InboundNatRuleResponse struct {
	BackendIPConfiguration  NetworkInterfaceIPConfigurationResponse `pulumi:"backendIPConfiguration"`
	BackendPort             *int                                    `pulumi:"backendPort"`
	EnableFloatingIP        *bool                                   `pulumi:"enableFloatingIP"`
	Etag                    *string                                 `pulumi:"etag"`
	FrontendIPConfiguration *SubResourceResponse                    `pulumi:"frontendIPConfiguration"`
	FrontendPort            *int                                    `pulumi:"frontendPort"`
	Id                      *string                                 `pulumi:"id"`
	IdleTimeoutInMinutes    *int                                    `pulumi:"idleTimeoutInMinutes"`
	Name                    *string                                 `pulumi:"name"`
	Protocol                *string                                 `pulumi:"protocol"`
	ProvisioningState       *string                                 `pulumi:"provisioningState"`
}

type InboundNatRuleResponseOutput struct{ *pulumi.OutputState }

func (InboundNatRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundNatRuleResponse)(nil)).Elem()
}

func (o InboundNatRuleResponseOutput) ToInboundNatRuleResponseOutput() InboundNatRuleResponseOutput {
	return o
}

func (o InboundNatRuleResponseOutput) ToInboundNatRuleResponseOutputWithContext(ctx context.Context) InboundNatRuleResponseOutput {
	return o
}

func (o InboundNatRuleResponseOutput) BackendIPConfiguration() NetworkInterfaceIPConfigurationResponseOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) NetworkInterfaceIPConfigurationResponse {
		return v.BackendIPConfiguration
	}).(NetworkInterfaceIPConfigurationResponseOutput)
}

func (o InboundNatRuleResponseOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *int { return v.BackendPort }).(pulumi.IntPtrOutput)
}

func (o InboundNatRuleResponseOutput) EnableFloatingIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *bool { return v.EnableFloatingIP }).(pulumi.BoolPtrOutput)
}

func (o InboundNatRuleResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleResponseOutput) FrontendIPConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *SubResourceResponse { return v.FrontendIPConfiguration }).(SubResourceResponsePtrOutput)
}

func (o InboundNatRuleResponseOutput) FrontendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *int { return v.FrontendPort }).(pulumi.IntPtrOutput)
}

func (o InboundNatRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleResponseOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o InboundNatRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o InboundNatRuleResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundNatRuleResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type InboundNatRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (InboundNatRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundNatRuleResponse)(nil)).Elem()
}

func (o InboundNatRuleResponseArrayOutput) ToInboundNatRuleResponseArrayOutput() InboundNatRuleResponseArrayOutput {
	return o
}

func (o InboundNatRuleResponseArrayOutput) ToInboundNatRuleResponseArrayOutputWithContext(ctx context.Context) InboundNatRuleResponseArrayOutput {
	return o
}

func (o InboundNatRuleResponseArrayOutput) Index(i pulumi.IntInput) InboundNatRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundNatRuleResponse {
		return vs[0].([]InboundNatRuleResponse)[vs[1].(int)]
	}).(InboundNatRuleResponseOutput)
}

type IpsecPolicy struct {
	DhGroup             string `pulumi:"dhGroup"`
	IkeEncryption       string `pulumi:"ikeEncryption"`
	IkeIntegrity        string `pulumi:"ikeIntegrity"`
	IpsecEncryption     string `pulumi:"ipsecEncryption"`
	IpsecIntegrity      string `pulumi:"ipsecIntegrity"`
	PfsGroup            string `pulumi:"pfsGroup"`
	SaDataSizeKilobytes int    `pulumi:"saDataSizeKilobytes"`
	SaLifeTimeSeconds   int    `pulumi:"saLifeTimeSeconds"`
}





type IpsecPolicyInput interface {
	pulumi.Input

	ToIpsecPolicyOutput() IpsecPolicyOutput
	ToIpsecPolicyOutputWithContext(context.Context) IpsecPolicyOutput
}

type IpsecPolicyArgs struct {
	DhGroup             pulumi.StringInput `pulumi:"dhGroup"`
	IkeEncryption       pulumi.StringInput `pulumi:"ikeEncryption"`
	IkeIntegrity        pulumi.StringInput `pulumi:"ikeIntegrity"`
	IpsecEncryption     pulumi.StringInput `pulumi:"ipsecEncryption"`
	IpsecIntegrity      pulumi.StringInput `pulumi:"ipsecIntegrity"`
	PfsGroup            pulumi.StringInput `pulumi:"pfsGroup"`
	SaDataSizeKilobytes pulumi.IntInput    `pulumi:"saDataSizeKilobytes"`
	SaLifeTimeSeconds   pulumi.IntInput    `pulumi:"saLifeTimeSeconds"`
}

func (IpsecPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpsecPolicy)(nil)).Elem()
}

func (i IpsecPolicyArgs) ToIpsecPolicyOutput() IpsecPolicyOutput {
	return i.ToIpsecPolicyOutputWithContext(context.Background())
}

func (i IpsecPolicyArgs) ToIpsecPolicyOutputWithContext(ctx context.Context) IpsecPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsecPolicyOutput)
}





type IpsecPolicyArrayInput interface {
	pulumi.Input

	ToIpsecPolicyArrayOutput() IpsecPolicyArrayOutput
	ToIpsecPolicyArrayOutputWithContext(context.Context) IpsecPolicyArrayOutput
}

type IpsecPolicyArray []IpsecPolicyInput

func (IpsecPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpsecPolicy)(nil)).Elem()
}

func (i IpsecPolicyArray) ToIpsecPolicyArrayOutput() IpsecPolicyArrayOutput {
	return i.ToIpsecPolicyArrayOutputWithContext(context.Background())
}

func (i IpsecPolicyArray) ToIpsecPolicyArrayOutputWithContext(ctx context.Context) IpsecPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpsecPolicyArrayOutput)
}

type IpsecPolicyOutput struct{ *pulumi.OutputState }

func (IpsecPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpsecPolicy)(nil)).Elem()
}

func (o IpsecPolicyOutput) ToIpsecPolicyOutput() IpsecPolicyOutput {
	return o
}

func (o IpsecPolicyOutput) ToIpsecPolicyOutputWithContext(ctx context.Context) IpsecPolicyOutput {
	return o
}

func (o IpsecPolicyOutput) DhGroup() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicy) string { return v.DhGroup }).(pulumi.StringOutput)
}

func (o IpsecPolicyOutput) IkeEncryption() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicy) string { return v.IkeEncryption }).(pulumi.StringOutput)
}

func (o IpsecPolicyOutput) IkeIntegrity() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicy) string { return v.IkeIntegrity }).(pulumi.StringOutput)
}

func (o IpsecPolicyOutput) IpsecEncryption() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicy) string { return v.IpsecEncryption }).(pulumi.StringOutput)
}

func (o IpsecPolicyOutput) IpsecIntegrity() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicy) string { return v.IpsecIntegrity }).(pulumi.StringOutput)
}

func (o IpsecPolicyOutput) PfsGroup() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicy) string { return v.PfsGroup }).(pulumi.StringOutput)
}

func (o IpsecPolicyOutput) SaDataSizeKilobytes() pulumi.IntOutput {
	return o.ApplyT(func(v IpsecPolicy) int { return v.SaDataSizeKilobytes }).(pulumi.IntOutput)
}

func (o IpsecPolicyOutput) SaLifeTimeSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v IpsecPolicy) int { return v.SaLifeTimeSeconds }).(pulumi.IntOutput)
}

type IpsecPolicyArrayOutput struct{ *pulumi.OutputState }

func (IpsecPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpsecPolicy)(nil)).Elem()
}

func (o IpsecPolicyArrayOutput) ToIpsecPolicyArrayOutput() IpsecPolicyArrayOutput {
	return o
}

func (o IpsecPolicyArrayOutput) ToIpsecPolicyArrayOutputWithContext(ctx context.Context) IpsecPolicyArrayOutput {
	return o
}

func (o IpsecPolicyArrayOutput) Index(i pulumi.IntInput) IpsecPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpsecPolicy {
		return vs[0].([]IpsecPolicy)[vs[1].(int)]
	}).(IpsecPolicyOutput)
}

type IpsecPolicyResponse struct {
	DhGroup             string `pulumi:"dhGroup"`
	IkeEncryption       string `pulumi:"ikeEncryption"`
	IkeIntegrity        string `pulumi:"ikeIntegrity"`
	IpsecEncryption     string `pulumi:"ipsecEncryption"`
	IpsecIntegrity      string `pulumi:"ipsecIntegrity"`
	PfsGroup            string `pulumi:"pfsGroup"`
	SaDataSizeKilobytes int    `pulumi:"saDataSizeKilobytes"`
	SaLifeTimeSeconds   int    `pulumi:"saLifeTimeSeconds"`
}

type IpsecPolicyResponseOutput struct{ *pulumi.OutputState }

func (IpsecPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpsecPolicyResponse)(nil)).Elem()
}

func (o IpsecPolicyResponseOutput) ToIpsecPolicyResponseOutput() IpsecPolicyResponseOutput {
	return o
}

func (o IpsecPolicyResponseOutput) ToIpsecPolicyResponseOutputWithContext(ctx context.Context) IpsecPolicyResponseOutput {
	return o
}

func (o IpsecPolicyResponseOutput) DhGroup() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicyResponse) string { return v.DhGroup }).(pulumi.StringOutput)
}

func (o IpsecPolicyResponseOutput) IkeEncryption() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicyResponse) string { return v.IkeEncryption }).(pulumi.StringOutput)
}

func (o IpsecPolicyResponseOutput) IkeIntegrity() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicyResponse) string { return v.IkeIntegrity }).(pulumi.StringOutput)
}

func (o IpsecPolicyResponseOutput) IpsecEncryption() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicyResponse) string { return v.IpsecEncryption }).(pulumi.StringOutput)
}

func (o IpsecPolicyResponseOutput) IpsecIntegrity() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicyResponse) string { return v.IpsecIntegrity }).(pulumi.StringOutput)
}

func (o IpsecPolicyResponseOutput) PfsGroup() pulumi.StringOutput {
	return o.ApplyT(func(v IpsecPolicyResponse) string { return v.PfsGroup }).(pulumi.StringOutput)
}

func (o IpsecPolicyResponseOutput) SaDataSizeKilobytes() pulumi.IntOutput {
	return o.ApplyT(func(v IpsecPolicyResponse) int { return v.SaDataSizeKilobytes }).(pulumi.IntOutput)
}

func (o IpsecPolicyResponseOutput) SaLifeTimeSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v IpsecPolicyResponse) int { return v.SaLifeTimeSeconds }).(pulumi.IntOutput)
}

type IpsecPolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (IpsecPolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpsecPolicyResponse)(nil)).Elem()
}

func (o IpsecPolicyResponseArrayOutput) ToIpsecPolicyResponseArrayOutput() IpsecPolicyResponseArrayOutput {
	return o
}

func (o IpsecPolicyResponseArrayOutput) ToIpsecPolicyResponseArrayOutputWithContext(ctx context.Context) IpsecPolicyResponseArrayOutput {
	return o
}

func (o IpsecPolicyResponseArrayOutput) Index(i pulumi.IntInput) IpsecPolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpsecPolicyResponse {
		return vs[0].([]IpsecPolicyResponse)[vs[1].(int)]
	}).(IpsecPolicyResponseOutput)
}

type Ipv6ExpressRouteCircuitPeeringConfig struct {
	MicrosoftPeeringConfig     *ExpressRouteCircuitPeeringConfig `pulumi:"microsoftPeeringConfig"`
	PrimaryPeerAddressPrefix   *string                           `pulumi:"primaryPeerAddressPrefix"`
	RouteFilter                *RouteFilterType                  `pulumi:"routeFilter"`
	SecondaryPeerAddressPrefix *string                           `pulumi:"secondaryPeerAddressPrefix"`
	State                      *string                           `pulumi:"state"`
}





type Ipv6ExpressRouteCircuitPeeringConfigInput interface {
	pulumi.Input

	ToIpv6ExpressRouteCircuitPeeringConfigOutput() Ipv6ExpressRouteCircuitPeeringConfigOutput
	ToIpv6ExpressRouteCircuitPeeringConfigOutputWithContext(context.Context) Ipv6ExpressRouteCircuitPeeringConfigOutput
}

type Ipv6ExpressRouteCircuitPeeringConfigArgs struct {
	MicrosoftPeeringConfig     ExpressRouteCircuitPeeringConfigPtrInput `pulumi:"microsoftPeeringConfig"`
	PrimaryPeerAddressPrefix   pulumi.StringPtrInput                    `pulumi:"primaryPeerAddressPrefix"`
	RouteFilter                RouteFilterTypePtrInput                  `pulumi:"routeFilter"`
	SecondaryPeerAddressPrefix pulumi.StringPtrInput                    `pulumi:"secondaryPeerAddressPrefix"`
	State                      pulumi.StringPtrInput                    `pulumi:"state"`
}

func (Ipv6ExpressRouteCircuitPeeringConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Ipv6ExpressRouteCircuitPeeringConfig)(nil)).Elem()
}

func (i Ipv6ExpressRouteCircuitPeeringConfigArgs) ToIpv6ExpressRouteCircuitPeeringConfigOutput() Ipv6ExpressRouteCircuitPeeringConfigOutput {
	return i.ToIpv6ExpressRouteCircuitPeeringConfigOutputWithContext(context.Background())
}

func (i Ipv6ExpressRouteCircuitPeeringConfigArgs) ToIpv6ExpressRouteCircuitPeeringConfigOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6ExpressRouteCircuitPeeringConfigOutput)
}

func (i Ipv6ExpressRouteCircuitPeeringConfigArgs) ToIpv6ExpressRouteCircuitPeeringConfigPtrOutput() Ipv6ExpressRouteCircuitPeeringConfigPtrOutput {
	return i.ToIpv6ExpressRouteCircuitPeeringConfigPtrOutputWithContext(context.Background())
}

func (i Ipv6ExpressRouteCircuitPeeringConfigArgs) ToIpv6ExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6ExpressRouteCircuitPeeringConfigOutput).ToIpv6ExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx)
}









type Ipv6ExpressRouteCircuitPeeringConfigPtrInput interface {
	pulumi.Input

	ToIpv6ExpressRouteCircuitPeeringConfigPtrOutput() Ipv6ExpressRouteCircuitPeeringConfigPtrOutput
	ToIpv6ExpressRouteCircuitPeeringConfigPtrOutputWithContext(context.Context) Ipv6ExpressRouteCircuitPeeringConfigPtrOutput
}

type ipv6ExpressRouteCircuitPeeringConfigPtrType Ipv6ExpressRouteCircuitPeeringConfigArgs

func Ipv6ExpressRouteCircuitPeeringConfigPtr(v *Ipv6ExpressRouteCircuitPeeringConfigArgs) Ipv6ExpressRouteCircuitPeeringConfigPtrInput {
	return (*ipv6ExpressRouteCircuitPeeringConfigPtrType)(v)
}

func (*ipv6ExpressRouteCircuitPeeringConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv6ExpressRouteCircuitPeeringConfig)(nil)).Elem()
}

func (i *ipv6ExpressRouteCircuitPeeringConfigPtrType) ToIpv6ExpressRouteCircuitPeeringConfigPtrOutput() Ipv6ExpressRouteCircuitPeeringConfigPtrOutput {
	return i.ToIpv6ExpressRouteCircuitPeeringConfigPtrOutputWithContext(context.Background())
}

func (i *ipv6ExpressRouteCircuitPeeringConfigPtrType) ToIpv6ExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv6ExpressRouteCircuitPeeringConfigPtrOutput)
}

type Ipv6ExpressRouteCircuitPeeringConfigOutput struct{ *pulumi.OutputState }

func (Ipv6ExpressRouteCircuitPeeringConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ipv6ExpressRouteCircuitPeeringConfig)(nil)).Elem()
}

func (o Ipv6ExpressRouteCircuitPeeringConfigOutput) ToIpv6ExpressRouteCircuitPeeringConfigOutput() Ipv6ExpressRouteCircuitPeeringConfigOutput {
	return o
}

func (o Ipv6ExpressRouteCircuitPeeringConfigOutput) ToIpv6ExpressRouteCircuitPeeringConfigOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigOutput {
	return o
}

func (o Ipv6ExpressRouteCircuitPeeringConfigOutput) ToIpv6ExpressRouteCircuitPeeringConfigPtrOutput() Ipv6ExpressRouteCircuitPeeringConfigPtrOutput {
	return o.ToIpv6ExpressRouteCircuitPeeringConfigPtrOutputWithContext(context.Background())
}

func (o Ipv6ExpressRouteCircuitPeeringConfigOutput) ToIpv6ExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ipv6ExpressRouteCircuitPeeringConfig) *Ipv6ExpressRouteCircuitPeeringConfig {
		return &v
	}).(Ipv6ExpressRouteCircuitPeeringConfigPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigPtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfig) *ExpressRouteCircuitPeeringConfig {
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfig) *string { return v.PrimaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigOutput) RouteFilter() RouteFilterTypePtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfig) *RouteFilterType { return v.RouteFilter }).(RouteFilterTypePtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfig) *string { return v.SecondaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfig) *string { return v.State }).(pulumi.StringPtrOutput)
}

type Ipv6ExpressRouteCircuitPeeringConfigPtrOutput struct{ *pulumi.OutputState }

func (Ipv6ExpressRouteCircuitPeeringConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv6ExpressRouteCircuitPeeringConfig)(nil)).Elem()
}

func (o Ipv6ExpressRouteCircuitPeeringConfigPtrOutput) ToIpv6ExpressRouteCircuitPeeringConfigPtrOutput() Ipv6ExpressRouteCircuitPeeringConfigPtrOutput {
	return o
}

func (o Ipv6ExpressRouteCircuitPeeringConfigPtrOutput) ToIpv6ExpressRouteCircuitPeeringConfigPtrOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigPtrOutput {
	return o
}

func (o Ipv6ExpressRouteCircuitPeeringConfigPtrOutput) Elem() Ipv6ExpressRouteCircuitPeeringConfigOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfig) Ipv6ExpressRouteCircuitPeeringConfig {
		if v != nil {
			return *v
		}
		var ret Ipv6ExpressRouteCircuitPeeringConfig
		return ret
	}).(Ipv6ExpressRouteCircuitPeeringConfigOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigPtrOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigPtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfig) *ExpressRouteCircuitPeeringConfig {
		if v == nil {
			return nil
		}
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigPtrOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfig) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryPeerAddressPrefix
	}).(pulumi.StringPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigPtrOutput) RouteFilter() RouteFilterTypePtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfig) *RouteFilterType {
		if v == nil {
			return nil
		}
		return v.RouteFilter
	}).(RouteFilterTypePtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigPtrOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfig) *string {
		if v == nil {
			return nil
		}
		return v.SecondaryPeerAddressPrefix
	}).(pulumi.StringPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfig) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type Ipv6ExpressRouteCircuitPeeringConfigResponse struct {
	MicrosoftPeeringConfig     *ExpressRouteCircuitPeeringConfigResponse `pulumi:"microsoftPeeringConfig"`
	PrimaryPeerAddressPrefix   *string                                   `pulumi:"primaryPeerAddressPrefix"`
	RouteFilter                *RouteFilterResponse                      `pulumi:"routeFilter"`
	SecondaryPeerAddressPrefix *string                                   `pulumi:"secondaryPeerAddressPrefix"`
	State                      *string                                   `pulumi:"state"`
}

type Ipv6ExpressRouteCircuitPeeringConfigResponseOutput struct{ *pulumi.OutputState }

func (Ipv6ExpressRouteCircuitPeeringConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ipv6ExpressRouteCircuitPeeringConfigResponse)(nil)).Elem()
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponseOutput) ToIpv6ExpressRouteCircuitPeeringConfigResponseOutput() Ipv6ExpressRouteCircuitPeeringConfigResponseOutput {
	return o
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponseOutput) ToIpv6ExpressRouteCircuitPeeringConfigResponseOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigResponseOutput {
	return o
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponseOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfigResponse) *ExpressRouteCircuitPeeringConfigResponse {
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponseOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfigResponse) *string { return v.PrimaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponseOutput) RouteFilter() RouteFilterResponsePtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfigResponse) *RouteFilterResponse { return v.RouteFilter }).(RouteFilterResponsePtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponseOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfigResponse) *string { return v.SecondaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6ExpressRouteCircuitPeeringConfigResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv6ExpressRouteCircuitPeeringConfigResponse)(nil)).Elem()
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput) ToIpv6ExpressRouteCircuitPeeringConfigResponsePtrOutput() Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput) ToIpv6ExpressRouteCircuitPeeringConfigResponsePtrOutputWithContext(ctx context.Context) Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput) Elem() Ipv6ExpressRouteCircuitPeeringConfigResponseOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfigResponse) Ipv6ExpressRouteCircuitPeeringConfigResponse {
		if v != nil {
			return *v
		}
		var ret Ipv6ExpressRouteCircuitPeeringConfigResponse
		return ret
	}).(Ipv6ExpressRouteCircuitPeeringConfigResponseOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfigResponse) *ExpressRouteCircuitPeeringConfigResponse {
		if v == nil {
			return nil
		}
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryPeerAddressPrefix
	}).(pulumi.StringPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput) RouteFilter() RouteFilterResponsePtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfigResponse) *RouteFilterResponse {
		if v == nil {
			return nil
		}
		return v.RouteFilter
	}).(RouteFilterResponsePtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.SecondaryPeerAddressPrefix
	}).(pulumi.StringPtrOutput)
}

func (o Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv6ExpressRouteCircuitPeeringConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type LoadBalancingRule struct {
	BackendAddressPool      *SubResource `pulumi:"backendAddressPool"`
	BackendPort             *int         `pulumi:"backendPort"`
	EnableFloatingIP        *bool        `pulumi:"enableFloatingIP"`
	Etag                    *string      `pulumi:"etag"`
	FrontendIPConfiguration *SubResource `pulumi:"frontendIPConfiguration"`
	FrontendPort            int          `pulumi:"frontendPort"`
	Id                      *string      `pulumi:"id"`
	IdleTimeoutInMinutes    *int         `pulumi:"idleTimeoutInMinutes"`
	LoadDistribution        *string      `pulumi:"loadDistribution"`
	Name                    *string      `pulumi:"name"`
	Probe                   *SubResource `pulumi:"probe"`
	Protocol                string       `pulumi:"protocol"`
	ProvisioningState       *string      `pulumi:"provisioningState"`
}





type LoadBalancingRuleInput interface {
	pulumi.Input

	ToLoadBalancingRuleOutput() LoadBalancingRuleOutput
	ToLoadBalancingRuleOutputWithContext(context.Context) LoadBalancingRuleOutput
}

type LoadBalancingRuleArgs struct {
	BackendAddressPool      SubResourcePtrInput   `pulumi:"backendAddressPool"`
	BackendPort             pulumi.IntPtrInput    `pulumi:"backendPort"`
	EnableFloatingIP        pulumi.BoolPtrInput   `pulumi:"enableFloatingIP"`
	Etag                    pulumi.StringPtrInput `pulumi:"etag"`
	FrontendIPConfiguration SubResourcePtrInput   `pulumi:"frontendIPConfiguration"`
	FrontendPort            pulumi.IntInput       `pulumi:"frontendPort"`
	Id                      pulumi.StringPtrInput `pulumi:"id"`
	IdleTimeoutInMinutes    pulumi.IntPtrInput    `pulumi:"idleTimeoutInMinutes"`
	LoadDistribution        pulumi.StringPtrInput `pulumi:"loadDistribution"`
	Name                    pulumi.StringPtrInput `pulumi:"name"`
	Probe                   SubResourcePtrInput   `pulumi:"probe"`
	Protocol                pulumi.StringInput    `pulumi:"protocol"`
	ProvisioningState       pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (LoadBalancingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingRule)(nil)).Elem()
}

func (i LoadBalancingRuleArgs) ToLoadBalancingRuleOutput() LoadBalancingRuleOutput {
	return i.ToLoadBalancingRuleOutputWithContext(context.Background())
}

func (i LoadBalancingRuleArgs) ToLoadBalancingRuleOutputWithContext(ctx context.Context) LoadBalancingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingRuleOutput)
}





type LoadBalancingRuleArrayInput interface {
	pulumi.Input

	ToLoadBalancingRuleArrayOutput() LoadBalancingRuleArrayOutput
	ToLoadBalancingRuleArrayOutputWithContext(context.Context) LoadBalancingRuleArrayOutput
}

type LoadBalancingRuleArray []LoadBalancingRuleInput

func (LoadBalancingRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingRule)(nil)).Elem()
}

func (i LoadBalancingRuleArray) ToLoadBalancingRuleArrayOutput() LoadBalancingRuleArrayOutput {
	return i.ToLoadBalancingRuleArrayOutputWithContext(context.Background())
}

func (i LoadBalancingRuleArray) ToLoadBalancingRuleArrayOutputWithContext(ctx context.Context) LoadBalancingRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancingRuleArrayOutput)
}

type LoadBalancingRuleOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingRule)(nil)).Elem()
}

func (o LoadBalancingRuleOutput) ToLoadBalancingRuleOutput() LoadBalancingRuleOutput {
	return o
}

func (o LoadBalancingRuleOutput) ToLoadBalancingRuleOutputWithContext(ctx context.Context) LoadBalancingRuleOutput {
	return o
}

func (o LoadBalancingRuleOutput) BackendAddressPool() SubResourcePtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *SubResource { return v.BackendAddressPool }).(SubResourcePtrOutput)
}

func (o LoadBalancingRuleOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *int { return v.BackendPort }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingRuleOutput) EnableFloatingIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *bool { return v.EnableFloatingIP }).(pulumi.BoolPtrOutput)
}

func (o LoadBalancingRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleOutput) FrontendIPConfiguration() SubResourcePtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *SubResource { return v.FrontendIPConfiguration }).(SubResourcePtrOutput)
}

func (o LoadBalancingRuleOutput) FrontendPort() pulumi.IntOutput {
	return o.ApplyT(func(v LoadBalancingRule) int { return v.FrontendPort }).(pulumi.IntOutput)
}

func (o LoadBalancingRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingRuleOutput) LoadDistribution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *string { return v.LoadDistribution }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleOutput) Probe() SubResourcePtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *SubResource { return v.Probe }).(SubResourcePtrOutput)
}

func (o LoadBalancingRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingRule) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o LoadBalancingRuleOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRule) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type LoadBalancingRuleArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingRule)(nil)).Elem()
}

func (o LoadBalancingRuleArrayOutput) ToLoadBalancingRuleArrayOutput() LoadBalancingRuleArrayOutput {
	return o
}

func (o LoadBalancingRuleArrayOutput) ToLoadBalancingRuleArrayOutputWithContext(ctx context.Context) LoadBalancingRuleArrayOutput {
	return o
}

func (o LoadBalancingRuleArrayOutput) Index(i pulumi.IntInput) LoadBalancingRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancingRule {
		return vs[0].([]LoadBalancingRule)[vs[1].(int)]
	}).(LoadBalancingRuleOutput)
}

type LoadBalancingRuleResponse struct {
	BackendAddressPool      *SubResourceResponse `pulumi:"backendAddressPool"`
	BackendPort             *int                 `pulumi:"backendPort"`
	EnableFloatingIP        *bool                `pulumi:"enableFloatingIP"`
	Etag                    *string              `pulumi:"etag"`
	FrontendIPConfiguration *SubResourceResponse `pulumi:"frontendIPConfiguration"`
	FrontendPort            int                  `pulumi:"frontendPort"`
	Id                      *string              `pulumi:"id"`
	IdleTimeoutInMinutes    *int                 `pulumi:"idleTimeoutInMinutes"`
	LoadDistribution        *string              `pulumi:"loadDistribution"`
	Name                    *string              `pulumi:"name"`
	Probe                   *SubResourceResponse `pulumi:"probe"`
	Protocol                string               `pulumi:"protocol"`
	ProvisioningState       *string              `pulumi:"provisioningState"`
}

type LoadBalancingRuleResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancingRuleResponse)(nil)).Elem()
}

func (o LoadBalancingRuleResponseOutput) ToLoadBalancingRuleResponseOutput() LoadBalancingRuleResponseOutput {
	return o
}

func (o LoadBalancingRuleResponseOutput) ToLoadBalancingRuleResponseOutputWithContext(ctx context.Context) LoadBalancingRuleResponseOutput {
	return o
}

func (o LoadBalancingRuleResponseOutput) BackendAddressPool() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *SubResourceResponse { return v.BackendAddressPool }).(SubResourceResponsePtrOutput)
}

func (o LoadBalancingRuleResponseOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *int { return v.BackendPort }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) EnableFloatingIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *bool { return v.EnableFloatingIP }).(pulumi.BoolPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) FrontendIPConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *SubResourceResponse { return v.FrontendIPConfiguration }).(SubResourceResponsePtrOutput)
}

func (o LoadBalancingRuleResponseOutput) FrontendPort() pulumi.IntOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) int { return v.FrontendPort }).(pulumi.IntOutput)
}

func (o LoadBalancingRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) LoadDistribution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *string { return v.LoadDistribution }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancingRuleResponseOutput) Probe() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *SubResourceResponse { return v.Probe }).(SubResourceResponsePtrOutput)
}

func (o LoadBalancingRuleResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o LoadBalancingRuleResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancingRuleResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type LoadBalancingRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancingRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancingRuleResponse)(nil)).Elem()
}

func (o LoadBalancingRuleResponseArrayOutput) ToLoadBalancingRuleResponseArrayOutput() LoadBalancingRuleResponseArrayOutput {
	return o
}

func (o LoadBalancingRuleResponseArrayOutput) ToLoadBalancingRuleResponseArrayOutputWithContext(ctx context.Context) LoadBalancingRuleResponseArrayOutput {
	return o
}

func (o LoadBalancingRuleResponseArrayOutput) Index(i pulumi.IntInput) LoadBalancingRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancingRuleResponse {
		return vs[0].([]LoadBalancingRuleResponse)[vs[1].(int)]
	}).(LoadBalancingRuleResponseOutput)
}

type LocalNetworkGatewayType struct {
	BgpSettings              *BgpSettings      `pulumi:"bgpSettings"`
	Etag                     *string           `pulumi:"etag"`
	GatewayIpAddress         *string           `pulumi:"gatewayIpAddress"`
	Id                       *string           `pulumi:"id"`
	LocalNetworkAddressSpace *AddressSpace     `pulumi:"localNetworkAddressSpace"`
	Location                 *string           `pulumi:"location"`
	ResourceGuid             *string           `pulumi:"resourceGuid"`
	Tags                     map[string]string `pulumi:"tags"`
}





type LocalNetworkGatewayTypeInput interface {
	pulumi.Input

	ToLocalNetworkGatewayTypeOutput() LocalNetworkGatewayTypeOutput
	ToLocalNetworkGatewayTypeOutputWithContext(context.Context) LocalNetworkGatewayTypeOutput
}

type LocalNetworkGatewayTypeArgs struct {
	BgpSettings              BgpSettingsPtrInput   `pulumi:"bgpSettings"`
	Etag                     pulumi.StringPtrInput `pulumi:"etag"`
	GatewayIpAddress         pulumi.StringPtrInput `pulumi:"gatewayIpAddress"`
	Id                       pulumi.StringPtrInput `pulumi:"id"`
	LocalNetworkAddressSpace AddressSpacePtrInput  `pulumi:"localNetworkAddressSpace"`
	Location                 pulumi.StringPtrInput `pulumi:"location"`
	ResourceGuid             pulumi.StringPtrInput `pulumi:"resourceGuid"`
	Tags                     pulumi.StringMapInput `pulumi:"tags"`
}

func (LocalNetworkGatewayTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalNetworkGatewayType)(nil)).Elem()
}

func (i LocalNetworkGatewayTypeArgs) ToLocalNetworkGatewayTypeOutput() LocalNetworkGatewayTypeOutput {
	return i.ToLocalNetworkGatewayTypeOutputWithContext(context.Background())
}

func (i LocalNetworkGatewayTypeArgs) ToLocalNetworkGatewayTypeOutputWithContext(ctx context.Context) LocalNetworkGatewayTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNetworkGatewayTypeOutput)
}

func (i LocalNetworkGatewayTypeArgs) ToLocalNetworkGatewayTypePtrOutput() LocalNetworkGatewayTypePtrOutput {
	return i.ToLocalNetworkGatewayTypePtrOutputWithContext(context.Background())
}

func (i LocalNetworkGatewayTypeArgs) ToLocalNetworkGatewayTypePtrOutputWithContext(ctx context.Context) LocalNetworkGatewayTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNetworkGatewayTypeOutput).ToLocalNetworkGatewayTypePtrOutputWithContext(ctx)
}









type LocalNetworkGatewayTypePtrInput interface {
	pulumi.Input

	ToLocalNetworkGatewayTypePtrOutput() LocalNetworkGatewayTypePtrOutput
	ToLocalNetworkGatewayTypePtrOutputWithContext(context.Context) LocalNetworkGatewayTypePtrOutput
}

type localNetworkGatewayTypePtrType LocalNetworkGatewayTypeArgs

func LocalNetworkGatewayTypePtr(v *LocalNetworkGatewayTypeArgs) LocalNetworkGatewayTypePtrInput {
	return (*localNetworkGatewayTypePtrType)(v)
}

func (*localNetworkGatewayTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalNetworkGatewayType)(nil)).Elem()
}

func (i *localNetworkGatewayTypePtrType) ToLocalNetworkGatewayTypePtrOutput() LocalNetworkGatewayTypePtrOutput {
	return i.ToLocalNetworkGatewayTypePtrOutputWithContext(context.Background())
}

func (i *localNetworkGatewayTypePtrType) ToLocalNetworkGatewayTypePtrOutputWithContext(ctx context.Context) LocalNetworkGatewayTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNetworkGatewayTypePtrOutput)
}

type LocalNetworkGatewayTypeOutput struct{ *pulumi.OutputState }

func (LocalNetworkGatewayTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalNetworkGatewayType)(nil)).Elem()
}

func (o LocalNetworkGatewayTypeOutput) ToLocalNetworkGatewayTypeOutput() LocalNetworkGatewayTypeOutput {
	return o
}

func (o LocalNetworkGatewayTypeOutput) ToLocalNetworkGatewayTypeOutputWithContext(ctx context.Context) LocalNetworkGatewayTypeOutput {
	return o
}

func (o LocalNetworkGatewayTypeOutput) ToLocalNetworkGatewayTypePtrOutput() LocalNetworkGatewayTypePtrOutput {
	return o.ToLocalNetworkGatewayTypePtrOutputWithContext(context.Background())
}

func (o LocalNetworkGatewayTypeOutput) ToLocalNetworkGatewayTypePtrOutputWithContext(ctx context.Context) LocalNetworkGatewayTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocalNetworkGatewayType) *LocalNetworkGatewayType {
		return &v
	}).(LocalNetworkGatewayTypePtrOutput)
}

func (o LocalNetworkGatewayTypeOutput) BgpSettings() BgpSettingsPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayType) *BgpSettings { return v.BgpSettings }).(BgpSettingsPtrOutput)
}

func (o LocalNetworkGatewayTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypeOutput) GatewayIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayType) *string { return v.GatewayIpAddress }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypeOutput) LocalNetworkAddressSpace() AddressSpacePtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayType) *AddressSpace { return v.LocalNetworkAddressSpace }).(AddressSpacePtrOutput)
}

func (o LocalNetworkGatewayTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypeOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayType) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LocalNetworkGatewayType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type LocalNetworkGatewayTypePtrOutput struct{ *pulumi.OutputState }

func (LocalNetworkGatewayTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalNetworkGatewayType)(nil)).Elem()
}

func (o LocalNetworkGatewayTypePtrOutput) ToLocalNetworkGatewayTypePtrOutput() LocalNetworkGatewayTypePtrOutput {
	return o
}

func (o LocalNetworkGatewayTypePtrOutput) ToLocalNetworkGatewayTypePtrOutputWithContext(ctx context.Context) LocalNetworkGatewayTypePtrOutput {
	return o
}

func (o LocalNetworkGatewayTypePtrOutput) Elem() LocalNetworkGatewayTypeOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayType) LocalNetworkGatewayType {
		if v != nil {
			return *v
		}
		var ret LocalNetworkGatewayType
		return ret
	}).(LocalNetworkGatewayTypeOutput)
}

func (o LocalNetworkGatewayTypePtrOutput) BgpSettings() BgpSettingsPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayType) *BgpSettings {
		if v == nil {
			return nil
		}
		return v.BgpSettings
	}).(BgpSettingsPtrOutput)
}

func (o LocalNetworkGatewayTypePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypePtrOutput) GatewayIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.GatewayIpAddress
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypePtrOutput) LocalNetworkAddressSpace() AddressSpacePtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayType) *AddressSpace {
		if v == nil {
			return nil
		}
		return v.LocalNetworkAddressSpace
	}).(AddressSpacePtrOutput)
}

func (o LocalNetworkGatewayTypePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayTypePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayType) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

type LocalNetworkGatewayResponse struct {
	BgpSettings              *BgpSettingsResponse  `pulumi:"bgpSettings"`
	Etag                     *string               `pulumi:"etag"`
	GatewayIpAddress         *string               `pulumi:"gatewayIpAddress"`
	Id                       *string               `pulumi:"id"`
	LocalNetworkAddressSpace *AddressSpaceResponse `pulumi:"localNetworkAddressSpace"`
	Location                 *string               `pulumi:"location"`
	Name                     string                `pulumi:"name"`
	ProvisioningState        string                `pulumi:"provisioningState"`
	ResourceGuid             *string               `pulumi:"resourceGuid"`
	Tags                     map[string]string     `pulumi:"tags"`
	Type                     string                `pulumi:"type"`
}

type LocalNetworkGatewayResponseOutput struct{ *pulumi.OutputState }

func (LocalNetworkGatewayResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalNetworkGatewayResponse)(nil)).Elem()
}

func (o LocalNetworkGatewayResponseOutput) ToLocalNetworkGatewayResponseOutput() LocalNetworkGatewayResponseOutput {
	return o
}

func (o LocalNetworkGatewayResponseOutput) ToLocalNetworkGatewayResponseOutputWithContext(ctx context.Context) LocalNetworkGatewayResponseOutput {
	return o
}

func (o LocalNetworkGatewayResponseOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) *BgpSettingsResponse { return v.BgpSettings }).(BgpSettingsResponsePtrOutput)
}

func (o LocalNetworkGatewayResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponseOutput) GatewayIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) *string { return v.GatewayIpAddress }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponseOutput) LocalNetworkAddressSpace() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) *AddressSpaceResponse { return v.LocalNetworkAddressSpace }).(AddressSpaceResponsePtrOutput)
}

func (o LocalNetworkGatewayResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o LocalNetworkGatewayResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LocalNetworkGatewayResponseOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LocalNetworkGatewayResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LocalNetworkGatewayResponse) string { return v.Type }).(pulumi.StringOutput)
}

type LocalNetworkGatewayResponsePtrOutput struct{ *pulumi.OutputState }

func (LocalNetworkGatewayResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalNetworkGatewayResponse)(nil)).Elem()
}

func (o LocalNetworkGatewayResponsePtrOutput) ToLocalNetworkGatewayResponsePtrOutput() LocalNetworkGatewayResponsePtrOutput {
	return o
}

func (o LocalNetworkGatewayResponsePtrOutput) ToLocalNetworkGatewayResponsePtrOutputWithContext(ctx context.Context) LocalNetworkGatewayResponsePtrOutput {
	return o
}

func (o LocalNetworkGatewayResponsePtrOutput) Elem() LocalNetworkGatewayResponseOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) LocalNetworkGatewayResponse {
		if v != nil {
			return *v
		}
		var ret LocalNetworkGatewayResponse
		return ret
	}).(LocalNetworkGatewayResponseOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *BgpSettingsResponse {
		if v == nil {
			return nil
		}
		return v.BgpSettings
	}).(BgpSettingsResponsePtrOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) GatewayIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.GatewayIpAddress
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) LocalNetworkAddressSpace() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *AddressSpaceResponse {
		if v == nil {
			return nil
		}
		return v.LocalNetworkAddressSpace
	}).(AddressSpaceResponsePtrOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o LocalNetworkGatewayResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type MonitorConfig struct {
	Path                 *string  `pulumi:"path"`
	Port                 *float64 `pulumi:"port"`
	ProfileMonitorStatus *string  `pulumi:"profileMonitorStatus"`
	Protocol             *string  `pulumi:"protocol"`
}





type MonitorConfigInput interface {
	pulumi.Input

	ToMonitorConfigOutput() MonitorConfigOutput
	ToMonitorConfigOutputWithContext(context.Context) MonitorConfigOutput
}

type MonitorConfigArgs struct {
	Path                 pulumi.StringPtrInput  `pulumi:"path"`
	Port                 pulumi.Float64PtrInput `pulumi:"port"`
	ProfileMonitorStatus pulumi.StringPtrInput  `pulumi:"profileMonitorStatus"`
	Protocol             pulumi.StringPtrInput  `pulumi:"protocol"`
}

func (MonitorConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfig)(nil)).Elem()
}

func (i MonitorConfigArgs) ToMonitorConfigOutput() MonitorConfigOutput {
	return i.ToMonitorConfigOutputWithContext(context.Background())
}

func (i MonitorConfigArgs) ToMonitorConfigOutputWithContext(ctx context.Context) MonitorConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigOutput)
}

func (i MonitorConfigArgs) ToMonitorConfigPtrOutput() MonitorConfigPtrOutput {
	return i.ToMonitorConfigPtrOutputWithContext(context.Background())
}

func (i MonitorConfigArgs) ToMonitorConfigPtrOutputWithContext(ctx context.Context) MonitorConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigOutput).ToMonitorConfigPtrOutputWithContext(ctx)
}









type MonitorConfigPtrInput interface {
	pulumi.Input

	ToMonitorConfigPtrOutput() MonitorConfigPtrOutput
	ToMonitorConfigPtrOutputWithContext(context.Context) MonitorConfigPtrOutput
}

type monitorConfigPtrType MonitorConfigArgs

func MonitorConfigPtr(v *MonitorConfigArgs) MonitorConfigPtrInput {
	return (*monitorConfigPtrType)(v)
}

func (*monitorConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorConfig)(nil)).Elem()
}

func (i *monitorConfigPtrType) ToMonitorConfigPtrOutput() MonitorConfigPtrOutput {
	return i.ToMonitorConfigPtrOutputWithContext(context.Background())
}

func (i *monitorConfigPtrType) ToMonitorConfigPtrOutputWithContext(ctx context.Context) MonitorConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorConfigPtrOutput)
}

type MonitorConfigOutput struct{ *pulumi.OutputState }

func (MonitorConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfig)(nil)).Elem()
}

func (o MonitorConfigOutput) ToMonitorConfigOutput() MonitorConfigOutput {
	return o
}

func (o MonitorConfigOutput) ToMonitorConfigOutputWithContext(ctx context.Context) MonitorConfigOutput {
	return o
}

func (o MonitorConfigOutput) ToMonitorConfigPtrOutput() MonitorConfigPtrOutput {
	return o.ToMonitorConfigPtrOutputWithContext(context.Background())
}

func (o MonitorConfigOutput) ToMonitorConfigPtrOutputWithContext(ctx context.Context) MonitorConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitorConfig) *MonitorConfig {
		return &v
	}).(MonitorConfigPtrOutput)
}

func (o MonitorConfigOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfig) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o MonitorConfigOutput) Port() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonitorConfig) *float64 { return v.Port }).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigOutput) ProfileMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfig) *string { return v.ProfileMonitorStatus }).(pulumi.StringPtrOutput)
}

func (o MonitorConfigOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfig) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

type MonitorConfigPtrOutput struct{ *pulumi.OutputState }

func (MonitorConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorConfig)(nil)).Elem()
}

func (o MonitorConfigPtrOutput) ToMonitorConfigPtrOutput() MonitorConfigPtrOutput {
	return o
}

func (o MonitorConfigPtrOutput) ToMonitorConfigPtrOutputWithContext(ctx context.Context) MonitorConfigPtrOutput {
	return o
}

func (o MonitorConfigPtrOutput) Elem() MonitorConfigOutput {
	return o.ApplyT(func(v *MonitorConfig) MonitorConfig {
		if v != nil {
			return *v
		}
		var ret MonitorConfig
		return ret
	}).(MonitorConfigOutput)
}

func (o MonitorConfigPtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorConfig) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o MonitorConfigPtrOutput) Port() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitorConfig) *float64 {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigPtrOutput) ProfileMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorConfig) *string {
		if v == nil {
			return nil
		}
		return v.ProfileMonitorStatus
	}).(pulumi.StringPtrOutput)
}

func (o MonitorConfigPtrOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorConfig) *string {
		if v == nil {
			return nil
		}
		return v.Protocol
	}).(pulumi.StringPtrOutput)
}

type MonitorConfigResponse struct {
	Path                 *string  `pulumi:"path"`
	Port                 *float64 `pulumi:"port"`
	ProfileMonitorStatus *string  `pulumi:"profileMonitorStatus"`
	Protocol             *string  `pulumi:"protocol"`
}

type MonitorConfigResponseOutput struct{ *pulumi.OutputState }

func (MonitorConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigResponse)(nil)).Elem()
}

func (o MonitorConfigResponseOutput) ToMonitorConfigResponseOutput() MonitorConfigResponseOutput {
	return o
}

func (o MonitorConfigResponseOutput) ToMonitorConfigResponseOutputWithContext(ctx context.Context) MonitorConfigResponseOutput {
	return o
}

func (o MonitorConfigResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfigResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o MonitorConfigResponseOutput) Port() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MonitorConfigResponse) *float64 { return v.Port }).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigResponseOutput) ProfileMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfigResponse) *string { return v.ProfileMonitorStatus }).(pulumi.StringPtrOutput)
}

func (o MonitorConfigResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonitorConfigResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

type MonitorConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (MonitorConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorConfigResponse)(nil)).Elem()
}

func (o MonitorConfigResponsePtrOutput) ToMonitorConfigResponsePtrOutput() MonitorConfigResponsePtrOutput {
	return o
}

func (o MonitorConfigResponsePtrOutput) ToMonitorConfigResponsePtrOutputWithContext(ctx context.Context) MonitorConfigResponsePtrOutput {
	return o
}

func (o MonitorConfigResponsePtrOutput) Elem() MonitorConfigResponseOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) MonitorConfigResponse {
		if v != nil {
			return *v
		}
		var ret MonitorConfigResponse
		return ret
	}).(MonitorConfigResponseOutput)
}

func (o MonitorConfigResponsePtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o MonitorConfigResponsePtrOutput) Port() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.Float64PtrOutput)
}

func (o MonitorConfigResponsePtrOutput) ProfileMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProfileMonitorStatus
	}).(pulumi.StringPtrOutput)
}

func (o MonitorConfigResponsePtrOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitorConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Protocol
	}).(pulumi.StringPtrOutput)
}

type NetworkInterfaceDnsSettings struct {
	AppliedDnsServers        []string `pulumi:"appliedDnsServers"`
	DnsServers               []string `pulumi:"dnsServers"`
	InternalDnsNameLabel     *string  `pulumi:"internalDnsNameLabel"`
	InternalDomainNameSuffix *string  `pulumi:"internalDomainNameSuffix"`
	InternalFqdn             *string  `pulumi:"internalFqdn"`
}





type NetworkInterfaceDnsSettingsInput interface {
	pulumi.Input

	ToNetworkInterfaceDnsSettingsOutput() NetworkInterfaceDnsSettingsOutput
	ToNetworkInterfaceDnsSettingsOutputWithContext(context.Context) NetworkInterfaceDnsSettingsOutput
}

type NetworkInterfaceDnsSettingsArgs struct {
	AppliedDnsServers        pulumi.StringArrayInput `pulumi:"appliedDnsServers"`
	DnsServers               pulumi.StringArrayInput `pulumi:"dnsServers"`
	InternalDnsNameLabel     pulumi.StringPtrInput   `pulumi:"internalDnsNameLabel"`
	InternalDomainNameSuffix pulumi.StringPtrInput   `pulumi:"internalDomainNameSuffix"`
	InternalFqdn             pulumi.StringPtrInput   `pulumi:"internalFqdn"`
}

func (NetworkInterfaceDnsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceDnsSettings)(nil)).Elem()
}

func (i NetworkInterfaceDnsSettingsArgs) ToNetworkInterfaceDnsSettingsOutput() NetworkInterfaceDnsSettingsOutput {
	return i.ToNetworkInterfaceDnsSettingsOutputWithContext(context.Background())
}

func (i NetworkInterfaceDnsSettingsArgs) ToNetworkInterfaceDnsSettingsOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceDnsSettingsOutput)
}

func (i NetworkInterfaceDnsSettingsArgs) ToNetworkInterfaceDnsSettingsPtrOutput() NetworkInterfaceDnsSettingsPtrOutput {
	return i.ToNetworkInterfaceDnsSettingsPtrOutputWithContext(context.Background())
}

func (i NetworkInterfaceDnsSettingsArgs) ToNetworkInterfaceDnsSettingsPtrOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceDnsSettingsOutput).ToNetworkInterfaceDnsSettingsPtrOutputWithContext(ctx)
}









type NetworkInterfaceDnsSettingsPtrInput interface {
	pulumi.Input

	ToNetworkInterfaceDnsSettingsPtrOutput() NetworkInterfaceDnsSettingsPtrOutput
	ToNetworkInterfaceDnsSettingsPtrOutputWithContext(context.Context) NetworkInterfaceDnsSettingsPtrOutput
}

type networkInterfaceDnsSettingsPtrType NetworkInterfaceDnsSettingsArgs

func NetworkInterfaceDnsSettingsPtr(v *NetworkInterfaceDnsSettingsArgs) NetworkInterfaceDnsSettingsPtrInput {
	return (*networkInterfaceDnsSettingsPtrType)(v)
}

func (*networkInterfaceDnsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceDnsSettings)(nil)).Elem()
}

func (i *networkInterfaceDnsSettingsPtrType) ToNetworkInterfaceDnsSettingsPtrOutput() NetworkInterfaceDnsSettingsPtrOutput {
	return i.ToNetworkInterfaceDnsSettingsPtrOutputWithContext(context.Background())
}

func (i *networkInterfaceDnsSettingsPtrType) ToNetworkInterfaceDnsSettingsPtrOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceDnsSettingsPtrOutput)
}

type NetworkInterfaceDnsSettingsOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceDnsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceDnsSettings)(nil)).Elem()
}

func (o NetworkInterfaceDnsSettingsOutput) ToNetworkInterfaceDnsSettingsOutput() NetworkInterfaceDnsSettingsOutput {
	return o
}

func (o NetworkInterfaceDnsSettingsOutput) ToNetworkInterfaceDnsSettingsOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsOutput {
	return o
}

func (o NetworkInterfaceDnsSettingsOutput) ToNetworkInterfaceDnsSettingsPtrOutput() NetworkInterfaceDnsSettingsPtrOutput {
	return o.ToNetworkInterfaceDnsSettingsPtrOutputWithContext(context.Background())
}

func (o NetworkInterfaceDnsSettingsOutput) ToNetworkInterfaceDnsSettingsPtrOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkInterfaceDnsSettings) *NetworkInterfaceDnsSettings {
		return &v
	}).(NetworkInterfaceDnsSettingsPtrOutput)
}

func (o NetworkInterfaceDnsSettingsOutput) AppliedDnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettings) []string { return v.AppliedDnsServers }).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceDnsSettingsOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettings) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceDnsSettingsOutput) InternalDnsNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettings) *string { return v.InternalDnsNameLabel }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceDnsSettingsOutput) InternalDomainNameSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettings) *string { return v.InternalDomainNameSuffix }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceDnsSettingsOutput) InternalFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettings) *string { return v.InternalFqdn }).(pulumi.StringPtrOutput)
}

type NetworkInterfaceDnsSettingsPtrOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceDnsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceDnsSettings)(nil)).Elem()
}

func (o NetworkInterfaceDnsSettingsPtrOutput) ToNetworkInterfaceDnsSettingsPtrOutput() NetworkInterfaceDnsSettingsPtrOutput {
	return o
}

func (o NetworkInterfaceDnsSettingsPtrOutput) ToNetworkInterfaceDnsSettingsPtrOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsPtrOutput {
	return o
}

func (o NetworkInterfaceDnsSettingsPtrOutput) Elem() NetworkInterfaceDnsSettingsOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettings) NetworkInterfaceDnsSettings {
		if v != nil {
			return *v
		}
		var ret NetworkInterfaceDnsSettings
		return ret
	}).(NetworkInterfaceDnsSettingsOutput)
}

func (o NetworkInterfaceDnsSettingsPtrOutput) AppliedDnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettings) []string {
		if v == nil {
			return nil
		}
		return v.AppliedDnsServers
	}).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceDnsSettingsPtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettings) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceDnsSettingsPtrOutput) InternalDnsNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettings) *string {
		if v == nil {
			return nil
		}
		return v.InternalDnsNameLabel
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceDnsSettingsPtrOutput) InternalDomainNameSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettings) *string {
		if v == nil {
			return nil
		}
		return v.InternalDomainNameSuffix
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceDnsSettingsPtrOutput) InternalFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettings) *string {
		if v == nil {
			return nil
		}
		return v.InternalFqdn
	}).(pulumi.StringPtrOutput)
}

type NetworkInterfaceDnsSettingsResponse struct {
	AppliedDnsServers        []string `pulumi:"appliedDnsServers"`
	DnsServers               []string `pulumi:"dnsServers"`
	InternalDnsNameLabel     *string  `pulumi:"internalDnsNameLabel"`
	InternalDomainNameSuffix *string  `pulumi:"internalDomainNameSuffix"`
	InternalFqdn             *string  `pulumi:"internalFqdn"`
}

type NetworkInterfaceDnsSettingsResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceDnsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceDnsSettingsResponse)(nil)).Elem()
}

func (o NetworkInterfaceDnsSettingsResponseOutput) ToNetworkInterfaceDnsSettingsResponseOutput() NetworkInterfaceDnsSettingsResponseOutput {
	return o
}

func (o NetworkInterfaceDnsSettingsResponseOutput) ToNetworkInterfaceDnsSettingsResponseOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsResponseOutput {
	return o
}

func (o NetworkInterfaceDnsSettingsResponseOutput) AppliedDnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettingsResponse) []string { return v.AppliedDnsServers }).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceDnsSettingsResponseOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettingsResponse) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceDnsSettingsResponseOutput) InternalDnsNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettingsResponse) *string { return v.InternalDnsNameLabel }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceDnsSettingsResponseOutput) InternalDomainNameSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettingsResponse) *string { return v.InternalDomainNameSuffix }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceDnsSettingsResponseOutput) InternalFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceDnsSettingsResponse) *string { return v.InternalFqdn }).(pulumi.StringPtrOutput)
}

type NetworkInterfaceDnsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceDnsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceDnsSettingsResponse)(nil)).Elem()
}

func (o NetworkInterfaceDnsSettingsResponsePtrOutput) ToNetworkInterfaceDnsSettingsResponsePtrOutput() NetworkInterfaceDnsSettingsResponsePtrOutput {
	return o
}

func (o NetworkInterfaceDnsSettingsResponsePtrOutput) ToNetworkInterfaceDnsSettingsResponsePtrOutputWithContext(ctx context.Context) NetworkInterfaceDnsSettingsResponsePtrOutput {
	return o
}

func (o NetworkInterfaceDnsSettingsResponsePtrOutput) Elem() NetworkInterfaceDnsSettingsResponseOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettingsResponse) NetworkInterfaceDnsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret NetworkInterfaceDnsSettingsResponse
		return ret
	}).(NetworkInterfaceDnsSettingsResponseOutput)
}

func (o NetworkInterfaceDnsSettingsResponsePtrOutput) AppliedDnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.AppliedDnsServers
	}).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceDnsSettingsResponsePtrOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.DnsServers
	}).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceDnsSettingsResponsePtrOutput) InternalDnsNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.InternalDnsNameLabel
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceDnsSettingsResponsePtrOutput) InternalDomainNameSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.InternalDomainNameSuffix
	}).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceDnsSettingsResponsePtrOutput) InternalFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkInterfaceDnsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.InternalFqdn
	}).(pulumi.StringPtrOutput)
}

type NetworkInterfaceIPConfiguration struct {
	ApplicationGatewayBackendAddressPools []ApplicationGatewayBackendAddressPool `pulumi:"applicationGatewayBackendAddressPools"`
	Etag                                  *string                                `pulumi:"etag"`
	Id                                    *string                                `pulumi:"id"`
	LoadBalancerBackendAddressPools       []BackendAddressPool                   `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatRules           []InboundNatRule                       `pulumi:"loadBalancerInboundNatRules"`
	Name                                  *string                                `pulumi:"name"`
	Primary                               *bool                                  `pulumi:"primary"`
	PrivateIPAddress                      *string                                `pulumi:"privateIPAddress"`
	PrivateIPAddressVersion               *string                                `pulumi:"privateIPAddressVersion"`
	PrivateIPAllocationMethod             *string                                `pulumi:"privateIPAllocationMethod"`
	ProvisioningState                     *string                                `pulumi:"provisioningState"`
	PublicIPAddress                       *PublicIPAddressType                   `pulumi:"publicIPAddress"`
	Subnet                                *SubnetType                            `pulumi:"subnet"`
}





type NetworkInterfaceIPConfigurationInput interface {
	pulumi.Input

	ToNetworkInterfaceIPConfigurationOutput() NetworkInterfaceIPConfigurationOutput
	ToNetworkInterfaceIPConfigurationOutputWithContext(context.Context) NetworkInterfaceIPConfigurationOutput
}

type NetworkInterfaceIPConfigurationArgs struct {
	ApplicationGatewayBackendAddressPools ApplicationGatewayBackendAddressPoolArrayInput `pulumi:"applicationGatewayBackendAddressPools"`
	Etag                                  pulumi.StringPtrInput                          `pulumi:"etag"`
	Id                                    pulumi.StringPtrInput                          `pulumi:"id"`
	LoadBalancerBackendAddressPools       BackendAddressPoolArrayInput                   `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatRules           InboundNatRuleArrayInput                       `pulumi:"loadBalancerInboundNatRules"`
	Name                                  pulumi.StringPtrInput                          `pulumi:"name"`
	Primary                               pulumi.BoolPtrInput                            `pulumi:"primary"`
	PrivateIPAddress                      pulumi.StringPtrInput                          `pulumi:"privateIPAddress"`
	PrivateIPAddressVersion               pulumi.StringPtrInput                          `pulumi:"privateIPAddressVersion"`
	PrivateIPAllocationMethod             pulumi.StringPtrInput                          `pulumi:"privateIPAllocationMethod"`
	ProvisioningState                     pulumi.StringPtrInput                          `pulumi:"provisioningState"`
	PublicIPAddress                       PublicIPAddressTypePtrInput                    `pulumi:"publicIPAddress"`
	Subnet                                SubnetTypePtrInput                             `pulumi:"subnet"`
}

func (NetworkInterfaceIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceIPConfiguration)(nil)).Elem()
}

func (i NetworkInterfaceIPConfigurationArgs) ToNetworkInterfaceIPConfigurationOutput() NetworkInterfaceIPConfigurationOutput {
	return i.ToNetworkInterfaceIPConfigurationOutputWithContext(context.Background())
}

func (i NetworkInterfaceIPConfigurationArgs) ToNetworkInterfaceIPConfigurationOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceIPConfigurationOutput)
}





type NetworkInterfaceIPConfigurationArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceIPConfigurationArrayOutput() NetworkInterfaceIPConfigurationArrayOutput
	ToNetworkInterfaceIPConfigurationArrayOutputWithContext(context.Context) NetworkInterfaceIPConfigurationArrayOutput
}

type NetworkInterfaceIPConfigurationArray []NetworkInterfaceIPConfigurationInput

func (NetworkInterfaceIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceIPConfiguration)(nil)).Elem()
}

func (i NetworkInterfaceIPConfigurationArray) ToNetworkInterfaceIPConfigurationArrayOutput() NetworkInterfaceIPConfigurationArrayOutput {
	return i.ToNetworkInterfaceIPConfigurationArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceIPConfigurationArray) ToNetworkInterfaceIPConfigurationArrayOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceIPConfigurationArrayOutput)
}

type NetworkInterfaceIPConfigurationOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceIPConfiguration)(nil)).Elem()
}

func (o NetworkInterfaceIPConfigurationOutput) ToNetworkInterfaceIPConfigurationOutput() NetworkInterfaceIPConfigurationOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationOutput) ToNetworkInterfaceIPConfigurationOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationOutput) ApplicationGatewayBackendAddressPools() ApplicationGatewayBackendAddressPoolArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) []ApplicationGatewayBackendAddressPool {
		return v.ApplicationGatewayBackendAddressPools
	}).(ApplicationGatewayBackendAddressPoolArrayOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) LoadBalancerBackendAddressPools() BackendAddressPoolArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) []BackendAddressPool { return v.LoadBalancerBackendAddressPools }).(BackendAddressPoolArrayOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) LoadBalancerInboundNatRules() InboundNatRuleArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) []InboundNatRule { return v.LoadBalancerInboundNatRules }).(InboundNatRuleArrayOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) PrivateIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.PrivateIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) PublicIPAddress() PublicIPAddressTypePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *PublicIPAddressType { return v.PublicIPAddress }).(PublicIPAddressTypePtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) Subnet() SubnetTypePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *SubnetType { return v.Subnet }).(SubnetTypePtrOutput)
}

type NetworkInterfaceIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceIPConfiguration)(nil)).Elem()
}

func (o NetworkInterfaceIPConfigurationArrayOutput) ToNetworkInterfaceIPConfigurationArrayOutput() NetworkInterfaceIPConfigurationArrayOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationArrayOutput) ToNetworkInterfaceIPConfigurationArrayOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationArrayOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceIPConfiguration {
		return vs[0].([]NetworkInterfaceIPConfiguration)[vs[1].(int)]
	}).(NetworkInterfaceIPConfigurationOutput)
}

type NetworkInterfaceIPConfigurationResponse struct {
	ApplicationGatewayBackendAddressPools []ApplicationGatewayBackendAddressPoolResponse `pulumi:"applicationGatewayBackendAddressPools"`
	Etag                                  *string                                        `pulumi:"etag"`
	Id                                    *string                                        `pulumi:"id"`
	LoadBalancerBackendAddressPools       []BackendAddressPoolResponse                   `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerInboundNatRules           []InboundNatRuleResponse                       `pulumi:"loadBalancerInboundNatRules"`
	Name                                  *string                                        `pulumi:"name"`
	Primary                               *bool                                          `pulumi:"primary"`
	PrivateIPAddress                      *string                                        `pulumi:"privateIPAddress"`
	PrivateIPAddressVersion               *string                                        `pulumi:"privateIPAddressVersion"`
	PrivateIPAllocationMethod             *string                                        `pulumi:"privateIPAllocationMethod"`
	ProvisioningState                     *string                                        `pulumi:"provisioningState"`
	PublicIPAddress                       *PublicIPAddressResponse                       `pulumi:"publicIPAddress"`
	Subnet                                *SubnetResponse                                `pulumi:"subnet"`
}

type NetworkInterfaceIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceIPConfigurationResponse)(nil)).Elem()
}

func (o NetworkInterfaceIPConfigurationResponseOutput) ToNetworkInterfaceIPConfigurationResponseOutput() NetworkInterfaceIPConfigurationResponseOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationResponseOutput) ToNetworkInterfaceIPConfigurationResponseOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationResponseOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationResponseOutput) ApplicationGatewayBackendAddressPools() ApplicationGatewayBackendAddressPoolResponseArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) []ApplicationGatewayBackendAddressPoolResponse {
		return v.ApplicationGatewayBackendAddressPools
	}).(ApplicationGatewayBackendAddressPoolResponseArrayOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) LoadBalancerBackendAddressPools() BackendAddressPoolResponseArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) []BackendAddressPoolResponse {
		return v.LoadBalancerBackendAddressPools
	}).(BackendAddressPoolResponseArrayOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) LoadBalancerInboundNatRules() InboundNatRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) []InboundNatRuleResponse {
		return v.LoadBalancerInboundNatRules
	}).(InboundNatRuleResponseArrayOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) PrivateIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.PrivateIPAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) PrivateIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.PrivateIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) PublicIPAddress() PublicIPAddressResponsePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *PublicIPAddressResponse { return v.PublicIPAddress }).(PublicIPAddressResponsePtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *SubnetResponse { return v.Subnet }).(SubnetResponsePtrOutput)
}

type NetworkInterfaceIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceIPConfigurationResponse)(nil)).Elem()
}

func (o NetworkInterfaceIPConfigurationResponseArrayOutput) ToNetworkInterfaceIPConfigurationResponseArrayOutput() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationResponseArrayOutput) ToNetworkInterfaceIPConfigurationResponseArrayOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceIPConfigurationResponse {
		return vs[0].([]NetworkInterfaceIPConfigurationResponse)[vs[1].(int)]
	}).(NetworkInterfaceIPConfigurationResponseOutput)
}

type NetworkInterfaceResponse struct {
	DnsSettings                 *NetworkInterfaceDnsSettingsResponse      `pulumi:"dnsSettings"`
	EnableAcceleratedNetworking *bool                                     `pulumi:"enableAcceleratedNetworking"`
	EnableIPForwarding          *bool                                     `pulumi:"enableIPForwarding"`
	Etag                        *string                                   `pulumi:"etag"`
	Id                          *string                                   `pulumi:"id"`
	IpConfigurations            []NetworkInterfaceIPConfigurationResponse `pulumi:"ipConfigurations"`
	Location                    *string                                   `pulumi:"location"`
	MacAddress                  *string                                   `pulumi:"macAddress"`
	Name                        string                                    `pulumi:"name"`
	NetworkSecurityGroup        *NetworkSecurityGroupResponse             `pulumi:"networkSecurityGroup"`
	Primary                     *bool                                     `pulumi:"primary"`
	ProvisioningState           *string                                   `pulumi:"provisioningState"`
	ResourceGuid                *string                                   `pulumi:"resourceGuid"`
	Tags                        map[string]string                         `pulumi:"tags"`
	Type                        string                                    `pulumi:"type"`
	VirtualMachine              *SubResourceResponse                      `pulumi:"virtualMachine"`
}

type NetworkInterfaceResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceResponse)(nil)).Elem()
}

func (o NetworkInterfaceResponseOutput) ToNetworkInterfaceResponseOutput() NetworkInterfaceResponseOutput {
	return o
}

func (o NetworkInterfaceResponseOutput) ToNetworkInterfaceResponseOutputWithContext(ctx context.Context) NetworkInterfaceResponseOutput {
	return o
}

func (o NetworkInterfaceResponseOutput) DnsSettings() NetworkInterfaceDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *NetworkInterfaceDnsSettingsResponse { return v.DnsSettings }).(NetworkInterfaceDnsSettingsResponsePtrOutput)
}

func (o NetworkInterfaceResponseOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *bool { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceResponseOutput) EnableIPForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *bool { return v.EnableIPForwarding }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceResponseOutput) IpConfigurations() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) []NetworkInterfaceIPConfigurationResponse { return v.IpConfigurations }).(NetworkInterfaceIPConfigurationResponseArrayOutput)
}

func (o NetworkInterfaceResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceResponseOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResponseOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *NetworkSecurityGroupResponse { return v.NetworkSecurityGroup }).(NetworkSecurityGroupResponsePtrOutput)
}

func (o NetworkInterfaceResponseOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceResponseOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkInterfaceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResponseOutput) VirtualMachine() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *SubResourceResponse { return v.VirtualMachine }).(SubResourceResponsePtrOutput)
}

type NetworkInterfaceResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceResponse)(nil)).Elem()
}

func (o NetworkInterfaceResponseArrayOutput) ToNetworkInterfaceResponseArrayOutput() NetworkInterfaceResponseArrayOutput {
	return o
}

func (o NetworkInterfaceResponseArrayOutput) ToNetworkInterfaceResponseArrayOutputWithContext(ctx context.Context) NetworkInterfaceResponseArrayOutput {
	return o
}

func (o NetworkInterfaceResponseArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceResponse {
		return vs[0].([]NetworkInterfaceResponse)[vs[1].(int)]
	}).(NetworkInterfaceResponseOutput)
}

type NetworkSecurityGroupType struct {
	DefaultSecurityRules []SecurityRuleType `pulumi:"defaultSecurityRules"`
	Etag                 *string            `pulumi:"etag"`
	Id                   *string            `pulumi:"id"`
	Location             *string            `pulumi:"location"`
	ProvisioningState    *string            `pulumi:"provisioningState"`
	ResourceGuid         *string            `pulumi:"resourceGuid"`
	SecurityRules        []SecurityRuleType `pulumi:"securityRules"`
	Tags                 map[string]string  `pulumi:"tags"`
}





type NetworkSecurityGroupTypeInput interface {
	pulumi.Input

	ToNetworkSecurityGroupTypeOutput() NetworkSecurityGroupTypeOutput
	ToNetworkSecurityGroupTypeOutputWithContext(context.Context) NetworkSecurityGroupTypeOutput
}

type NetworkSecurityGroupTypeArgs struct {
	DefaultSecurityRules SecurityRuleTypeArrayInput `pulumi:"defaultSecurityRules"`
	Etag                 pulumi.StringPtrInput      `pulumi:"etag"`
	Id                   pulumi.StringPtrInput      `pulumi:"id"`
	Location             pulumi.StringPtrInput      `pulumi:"location"`
	ProvisioningState    pulumi.StringPtrInput      `pulumi:"provisioningState"`
	ResourceGuid         pulumi.StringPtrInput      `pulumi:"resourceGuid"`
	SecurityRules        SecurityRuleTypeArrayInput `pulumi:"securityRules"`
	Tags                 pulumi.StringMapInput      `pulumi:"tags"`
}

func (NetworkSecurityGroupTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupType)(nil)).Elem()
}

func (i NetworkSecurityGroupTypeArgs) ToNetworkSecurityGroupTypeOutput() NetworkSecurityGroupTypeOutput {
	return i.ToNetworkSecurityGroupTypeOutputWithContext(context.Background())
}

func (i NetworkSecurityGroupTypeArgs) ToNetworkSecurityGroupTypeOutputWithContext(ctx context.Context) NetworkSecurityGroupTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityGroupTypeOutput)
}

func (i NetworkSecurityGroupTypeArgs) ToNetworkSecurityGroupTypePtrOutput() NetworkSecurityGroupTypePtrOutput {
	return i.ToNetworkSecurityGroupTypePtrOutputWithContext(context.Background())
}

func (i NetworkSecurityGroupTypeArgs) ToNetworkSecurityGroupTypePtrOutputWithContext(ctx context.Context) NetworkSecurityGroupTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityGroupTypeOutput).ToNetworkSecurityGroupTypePtrOutputWithContext(ctx)
}









type NetworkSecurityGroupTypePtrInput interface {
	pulumi.Input

	ToNetworkSecurityGroupTypePtrOutput() NetworkSecurityGroupTypePtrOutput
	ToNetworkSecurityGroupTypePtrOutputWithContext(context.Context) NetworkSecurityGroupTypePtrOutput
}

type networkSecurityGroupTypePtrType NetworkSecurityGroupTypeArgs

func NetworkSecurityGroupTypePtr(v *NetworkSecurityGroupTypeArgs) NetworkSecurityGroupTypePtrInput {
	return (*networkSecurityGroupTypePtrType)(v)
}

func (*networkSecurityGroupTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityGroupType)(nil)).Elem()
}

func (i *networkSecurityGroupTypePtrType) ToNetworkSecurityGroupTypePtrOutput() NetworkSecurityGroupTypePtrOutput {
	return i.ToNetworkSecurityGroupTypePtrOutputWithContext(context.Background())
}

func (i *networkSecurityGroupTypePtrType) ToNetworkSecurityGroupTypePtrOutputWithContext(ctx context.Context) NetworkSecurityGroupTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityGroupTypePtrOutput)
}

type NetworkSecurityGroupTypeOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupType)(nil)).Elem()
}

func (o NetworkSecurityGroupTypeOutput) ToNetworkSecurityGroupTypeOutput() NetworkSecurityGroupTypeOutput {
	return o
}

func (o NetworkSecurityGroupTypeOutput) ToNetworkSecurityGroupTypeOutputWithContext(ctx context.Context) NetworkSecurityGroupTypeOutput {
	return o
}

func (o NetworkSecurityGroupTypeOutput) ToNetworkSecurityGroupTypePtrOutput() NetworkSecurityGroupTypePtrOutput {
	return o.ToNetworkSecurityGroupTypePtrOutputWithContext(context.Background())
}

func (o NetworkSecurityGroupTypeOutput) ToNetworkSecurityGroupTypePtrOutputWithContext(ctx context.Context) NetworkSecurityGroupTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkSecurityGroupType) *NetworkSecurityGroupType {
		return &v
	}).(NetworkSecurityGroupTypePtrOutput)
}

func (o NetworkSecurityGroupTypeOutput) DefaultSecurityRules() SecurityRuleTypeArrayOutput {
	return o.ApplyT(func(v NetworkSecurityGroupType) []SecurityRuleType { return v.DefaultSecurityRules }).(SecurityRuleTypeArrayOutput)
}

func (o NetworkSecurityGroupTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypeOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupType) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypeOutput) SecurityRules() SecurityRuleTypeArrayOutput {
	return o.ApplyT(func(v NetworkSecurityGroupType) []SecurityRuleType { return v.SecurityRules }).(SecurityRuleTypeArrayOutput)
}

func (o NetworkSecurityGroupTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v NetworkSecurityGroupType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type NetworkSecurityGroupTypePtrOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityGroupType)(nil)).Elem()
}

func (o NetworkSecurityGroupTypePtrOutput) ToNetworkSecurityGroupTypePtrOutput() NetworkSecurityGroupTypePtrOutput {
	return o
}

func (o NetworkSecurityGroupTypePtrOutput) ToNetworkSecurityGroupTypePtrOutputWithContext(ctx context.Context) NetworkSecurityGroupTypePtrOutput {
	return o
}

func (o NetworkSecurityGroupTypePtrOutput) Elem() NetworkSecurityGroupTypeOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupType) NetworkSecurityGroupType {
		if v != nil {
			return *v
		}
		var ret NetworkSecurityGroupType
		return ret
	}).(NetworkSecurityGroupTypeOutput)
}

func (o NetworkSecurityGroupTypePtrOutput) DefaultSecurityRules() SecurityRuleTypeArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupType) []SecurityRuleType {
		if v == nil {
			return nil
		}
		return v.DefaultSecurityRules
	}).(SecurityRuleTypeArrayOutput)
}

func (o NetworkSecurityGroupTypePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupType) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupType) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupType) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupType) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupTypePtrOutput) SecurityRules() SecurityRuleTypeArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupType) []SecurityRuleType {
		if v == nil {
			return nil
		}
		return v.SecurityRules
	}).(SecurityRuleTypeArrayOutput)
}

func (o NetworkSecurityGroupTypePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupType) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

type NetworkSecurityGroupResponse struct {
	DefaultSecurityRules []SecurityRuleResponse     `pulumi:"defaultSecurityRules"`
	Etag                 *string                    `pulumi:"etag"`
	Id                   *string                    `pulumi:"id"`
	Location             *string                    `pulumi:"location"`
	Name                 string                     `pulumi:"name"`
	NetworkInterfaces    []NetworkInterfaceResponse `pulumi:"networkInterfaces"`
	ProvisioningState    *string                    `pulumi:"provisioningState"`
	ResourceGuid         *string                    `pulumi:"resourceGuid"`
	SecurityRules        []SecurityRuleResponse     `pulumi:"securityRules"`
	Subnets              []SubnetResponse           `pulumi:"subnets"`
	Tags                 map[string]string          `pulumi:"tags"`
	Type                 string                     `pulumi:"type"`
}

type NetworkSecurityGroupResponseOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupResponse)(nil)).Elem()
}

func (o NetworkSecurityGroupResponseOutput) ToNetworkSecurityGroupResponseOutput() NetworkSecurityGroupResponseOutput {
	return o
}

func (o NetworkSecurityGroupResponseOutput) ToNetworkSecurityGroupResponseOutputWithContext(ctx context.Context) NetworkSecurityGroupResponseOutput {
	return o
}

func (o NetworkSecurityGroupResponseOutput) DefaultSecurityRules() SecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) []SecurityRuleResponse { return v.DefaultSecurityRules }).(SecurityRuleResponseArrayOutput)
}

func (o NetworkSecurityGroupResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkSecurityGroupResponseOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) []NetworkInterfaceResponse { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

func (o NetworkSecurityGroupResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponseOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponseOutput) SecurityRules() SecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) []SecurityRuleResponse { return v.SecurityRules }).(SecurityRuleResponseArrayOutput)
}

func (o NetworkSecurityGroupResponseOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) []SubnetResponse { return v.Subnets }).(SubnetResponseArrayOutput)
}

func (o NetworkSecurityGroupResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NetworkSecurityGroupResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResponse) string { return v.Type }).(pulumi.StringOutput)
}

type NetworkSecurityGroupResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityGroupResponse)(nil)).Elem()
}

func (o NetworkSecurityGroupResponsePtrOutput) ToNetworkSecurityGroupResponsePtrOutput() NetworkSecurityGroupResponsePtrOutput {
	return o
}

func (o NetworkSecurityGroupResponsePtrOutput) ToNetworkSecurityGroupResponsePtrOutputWithContext(ctx context.Context) NetworkSecurityGroupResponsePtrOutput {
	return o
}

func (o NetworkSecurityGroupResponsePtrOutput) Elem() NetworkSecurityGroupResponseOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) NetworkSecurityGroupResponse {
		if v != nil {
			return *v
		}
		var ret NetworkSecurityGroupResponse
		return ret
	}).(NetworkSecurityGroupResponseOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) DefaultSecurityRules() SecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) []SecurityRuleResponse {
		if v == nil {
			return nil
		}
		return v.DefaultSecurityRules
	}).(SecurityRuleResponseArrayOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) []NetworkInterfaceResponse {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(NetworkInterfaceResponseArrayOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) SecurityRules() SecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) []SecurityRuleResponse {
		if v == nil {
			return nil
		}
		return v.SecurityRules
	}).(SecurityRuleResponseArrayOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) []SubnetResponse {
		if v == nil {
			return nil
		}
		return v.Subnets
	}).(SubnetResponseArrayOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o NetworkSecurityGroupResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroupResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type OutboundNatRule struct {
	AllocatedOutboundPorts   *int          `pulumi:"allocatedOutboundPorts"`
	BackendAddressPool       SubResource   `pulumi:"backendAddressPool"`
	Etag                     *string       `pulumi:"etag"`
	FrontendIPConfigurations []SubResource `pulumi:"frontendIPConfigurations"`
	Id                       *string       `pulumi:"id"`
	Name                     *string       `pulumi:"name"`
	ProvisioningState        *string       `pulumi:"provisioningState"`
}





type OutboundNatRuleInput interface {
	pulumi.Input

	ToOutboundNatRuleOutput() OutboundNatRuleOutput
	ToOutboundNatRuleOutputWithContext(context.Context) OutboundNatRuleOutput
}

type OutboundNatRuleArgs struct {
	AllocatedOutboundPorts   pulumi.IntPtrInput    `pulumi:"allocatedOutboundPorts"`
	BackendAddressPool       SubResourceInput      `pulumi:"backendAddressPool"`
	Etag                     pulumi.StringPtrInput `pulumi:"etag"`
	FrontendIPConfigurations SubResourceArrayInput `pulumi:"frontendIPConfigurations"`
	Id                       pulumi.StringPtrInput `pulumi:"id"`
	Name                     pulumi.StringPtrInput `pulumi:"name"`
	ProvisioningState        pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (OutboundNatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OutboundNatRule)(nil)).Elem()
}

func (i OutboundNatRuleArgs) ToOutboundNatRuleOutput() OutboundNatRuleOutput {
	return i.ToOutboundNatRuleOutputWithContext(context.Background())
}

func (i OutboundNatRuleArgs) ToOutboundNatRuleOutputWithContext(ctx context.Context) OutboundNatRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutboundNatRuleOutput)
}





type OutboundNatRuleArrayInput interface {
	pulumi.Input

	ToOutboundNatRuleArrayOutput() OutboundNatRuleArrayOutput
	ToOutboundNatRuleArrayOutputWithContext(context.Context) OutboundNatRuleArrayOutput
}

type OutboundNatRuleArray []OutboundNatRuleInput

func (OutboundNatRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutboundNatRule)(nil)).Elem()
}

func (i OutboundNatRuleArray) ToOutboundNatRuleArrayOutput() OutboundNatRuleArrayOutput {
	return i.ToOutboundNatRuleArrayOutputWithContext(context.Background())
}

func (i OutboundNatRuleArray) ToOutboundNatRuleArrayOutputWithContext(ctx context.Context) OutboundNatRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutboundNatRuleArrayOutput)
}

type OutboundNatRuleOutput struct{ *pulumi.OutputState }

func (OutboundNatRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutboundNatRule)(nil)).Elem()
}

func (o OutboundNatRuleOutput) ToOutboundNatRuleOutput() OutboundNatRuleOutput {
	return o
}

func (o OutboundNatRuleOutput) ToOutboundNatRuleOutputWithContext(ctx context.Context) OutboundNatRuleOutput {
	return o
}

func (o OutboundNatRuleOutput) AllocatedOutboundPorts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OutboundNatRule) *int { return v.AllocatedOutboundPorts }).(pulumi.IntPtrOutput)
}

func (o OutboundNatRuleOutput) BackendAddressPool() SubResourceOutput {
	return o.ApplyT(func(v OutboundNatRule) SubResource { return v.BackendAddressPool }).(SubResourceOutput)
}

func (o OutboundNatRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutboundNatRule) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o OutboundNatRuleOutput) FrontendIPConfigurations() SubResourceArrayOutput {
	return o.ApplyT(func(v OutboundNatRule) []SubResource { return v.FrontendIPConfigurations }).(SubResourceArrayOutput)
}

func (o OutboundNatRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutboundNatRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o OutboundNatRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutboundNatRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OutboundNatRuleOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutboundNatRule) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type OutboundNatRuleArrayOutput struct{ *pulumi.OutputState }

func (OutboundNatRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutboundNatRule)(nil)).Elem()
}

func (o OutboundNatRuleArrayOutput) ToOutboundNatRuleArrayOutput() OutboundNatRuleArrayOutput {
	return o
}

func (o OutboundNatRuleArrayOutput) ToOutboundNatRuleArrayOutputWithContext(ctx context.Context) OutboundNatRuleArrayOutput {
	return o
}

func (o OutboundNatRuleArrayOutput) Index(i pulumi.IntInput) OutboundNatRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OutboundNatRule {
		return vs[0].([]OutboundNatRule)[vs[1].(int)]
	}).(OutboundNatRuleOutput)
}

type OutboundNatRuleResponse struct {
	AllocatedOutboundPorts   *int                  `pulumi:"allocatedOutboundPorts"`
	BackendAddressPool       SubResourceResponse   `pulumi:"backendAddressPool"`
	Etag                     *string               `pulumi:"etag"`
	FrontendIPConfigurations []SubResourceResponse `pulumi:"frontendIPConfigurations"`
	Id                       *string               `pulumi:"id"`
	Name                     *string               `pulumi:"name"`
	ProvisioningState        *string               `pulumi:"provisioningState"`
}

type OutboundNatRuleResponseOutput struct{ *pulumi.OutputState }

func (OutboundNatRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutboundNatRuleResponse)(nil)).Elem()
}

func (o OutboundNatRuleResponseOutput) ToOutboundNatRuleResponseOutput() OutboundNatRuleResponseOutput {
	return o
}

func (o OutboundNatRuleResponseOutput) ToOutboundNatRuleResponseOutputWithContext(ctx context.Context) OutboundNatRuleResponseOutput {
	return o
}

func (o OutboundNatRuleResponseOutput) AllocatedOutboundPorts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OutboundNatRuleResponse) *int { return v.AllocatedOutboundPorts }).(pulumi.IntPtrOutput)
}

func (o OutboundNatRuleResponseOutput) BackendAddressPool() SubResourceResponseOutput {
	return o.ApplyT(func(v OutboundNatRuleResponse) SubResourceResponse { return v.BackendAddressPool }).(SubResourceResponseOutput)
}

func (o OutboundNatRuleResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutboundNatRuleResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o OutboundNatRuleResponseOutput) FrontendIPConfigurations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v OutboundNatRuleResponse) []SubResourceResponse { return v.FrontendIPConfigurations }).(SubResourceResponseArrayOutput)
}

func (o OutboundNatRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutboundNatRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o OutboundNatRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutboundNatRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OutboundNatRuleResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutboundNatRuleResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type OutboundNatRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (OutboundNatRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutboundNatRuleResponse)(nil)).Elem()
}

func (o OutboundNatRuleResponseArrayOutput) ToOutboundNatRuleResponseArrayOutput() OutboundNatRuleResponseArrayOutput {
	return o
}

func (o OutboundNatRuleResponseArrayOutput) ToOutboundNatRuleResponseArrayOutputWithContext(ctx context.Context) OutboundNatRuleResponseArrayOutput {
	return o
}

func (o OutboundNatRuleResponseArrayOutput) Index(i pulumi.IntInput) OutboundNatRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OutboundNatRuleResponse {
		return vs[0].([]OutboundNatRuleResponse)[vs[1].(int)]
	}).(OutboundNatRuleResponseOutput)
}

type PacketCaptureFilter struct {
	LocalIPAddress  *string `pulumi:"localIPAddress"`
	LocalPort       *string `pulumi:"localPort"`
	Protocol        *string `pulumi:"protocol"`
	RemoteIPAddress *string `pulumi:"remoteIPAddress"`
	RemotePort      *string `pulumi:"remotePort"`
}


func (val *PacketCaptureFilter) Defaults() *PacketCaptureFilter {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Protocol) {
		protocol_ := "Any"
		tmp.Protocol = &protocol_
	}
	return &tmp
}





type PacketCaptureFilterInput interface {
	pulumi.Input

	ToPacketCaptureFilterOutput() PacketCaptureFilterOutput
	ToPacketCaptureFilterOutputWithContext(context.Context) PacketCaptureFilterOutput
}

type PacketCaptureFilterArgs struct {
	LocalIPAddress  pulumi.StringPtrInput `pulumi:"localIPAddress"`
	LocalPort       pulumi.StringPtrInput `pulumi:"localPort"`
	Protocol        pulumi.StringPtrInput `pulumi:"protocol"`
	RemoteIPAddress pulumi.StringPtrInput `pulumi:"remoteIPAddress"`
	RemotePort      pulumi.StringPtrInput `pulumi:"remotePort"`
}


func (val *PacketCaptureFilterArgs) Defaults() *PacketCaptureFilterArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Protocol) {
		tmp.Protocol = pulumi.StringPtr("Any")
	}
	return &tmp
}
func (PacketCaptureFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PacketCaptureFilter)(nil)).Elem()
}

func (i PacketCaptureFilterArgs) ToPacketCaptureFilterOutput() PacketCaptureFilterOutput {
	return i.ToPacketCaptureFilterOutputWithContext(context.Background())
}

func (i PacketCaptureFilterArgs) ToPacketCaptureFilterOutputWithContext(ctx context.Context) PacketCaptureFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketCaptureFilterOutput)
}





type PacketCaptureFilterArrayInput interface {
	pulumi.Input

	ToPacketCaptureFilterArrayOutput() PacketCaptureFilterArrayOutput
	ToPacketCaptureFilterArrayOutputWithContext(context.Context) PacketCaptureFilterArrayOutput
}

type PacketCaptureFilterArray []PacketCaptureFilterInput

func (PacketCaptureFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PacketCaptureFilter)(nil)).Elem()
}

func (i PacketCaptureFilterArray) ToPacketCaptureFilterArrayOutput() PacketCaptureFilterArrayOutput {
	return i.ToPacketCaptureFilterArrayOutputWithContext(context.Background())
}

func (i PacketCaptureFilterArray) ToPacketCaptureFilterArrayOutputWithContext(ctx context.Context) PacketCaptureFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketCaptureFilterArrayOutput)
}

type PacketCaptureFilterOutput struct{ *pulumi.OutputState }

func (PacketCaptureFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PacketCaptureFilter)(nil)).Elem()
}

func (o PacketCaptureFilterOutput) ToPacketCaptureFilterOutput() PacketCaptureFilterOutput {
	return o
}

func (o PacketCaptureFilterOutput) ToPacketCaptureFilterOutputWithContext(ctx context.Context) PacketCaptureFilterOutput {
	return o
}

func (o PacketCaptureFilterOutput) LocalIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilter) *string { return v.LocalIPAddress }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureFilterOutput) LocalPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilter) *string { return v.LocalPort }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureFilterOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilter) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureFilterOutput) RemoteIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilter) *string { return v.RemoteIPAddress }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureFilterOutput) RemotePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilter) *string { return v.RemotePort }).(pulumi.StringPtrOutput)
}

type PacketCaptureFilterArrayOutput struct{ *pulumi.OutputState }

func (PacketCaptureFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PacketCaptureFilter)(nil)).Elem()
}

func (o PacketCaptureFilterArrayOutput) ToPacketCaptureFilterArrayOutput() PacketCaptureFilterArrayOutput {
	return o
}

func (o PacketCaptureFilterArrayOutput) ToPacketCaptureFilterArrayOutputWithContext(ctx context.Context) PacketCaptureFilterArrayOutput {
	return o
}

func (o PacketCaptureFilterArrayOutput) Index(i pulumi.IntInput) PacketCaptureFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PacketCaptureFilter {
		return vs[0].([]PacketCaptureFilter)[vs[1].(int)]
	}).(PacketCaptureFilterOutput)
}

type PacketCaptureFilterResponse struct {
	LocalIPAddress  *string `pulumi:"localIPAddress"`
	LocalPort       *string `pulumi:"localPort"`
	Protocol        *string `pulumi:"protocol"`
	RemoteIPAddress *string `pulumi:"remoteIPAddress"`
	RemotePort      *string `pulumi:"remotePort"`
}


func (val *PacketCaptureFilterResponse) Defaults() *PacketCaptureFilterResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Protocol) {
		protocol_ := "Any"
		tmp.Protocol = &protocol_
	}
	return &tmp
}

type PacketCaptureFilterResponseOutput struct{ *pulumi.OutputState }

func (PacketCaptureFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PacketCaptureFilterResponse)(nil)).Elem()
}

func (o PacketCaptureFilterResponseOutput) ToPacketCaptureFilterResponseOutput() PacketCaptureFilterResponseOutput {
	return o
}

func (o PacketCaptureFilterResponseOutput) ToPacketCaptureFilterResponseOutputWithContext(ctx context.Context) PacketCaptureFilterResponseOutput {
	return o
}

func (o PacketCaptureFilterResponseOutput) LocalIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilterResponse) *string { return v.LocalIPAddress }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureFilterResponseOutput) LocalPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilterResponse) *string { return v.LocalPort }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureFilterResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilterResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureFilterResponseOutput) RemoteIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilterResponse) *string { return v.RemoteIPAddress }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureFilterResponseOutput) RemotePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureFilterResponse) *string { return v.RemotePort }).(pulumi.StringPtrOutput)
}

type PacketCaptureFilterResponseArrayOutput struct{ *pulumi.OutputState }

func (PacketCaptureFilterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PacketCaptureFilterResponse)(nil)).Elem()
}

func (o PacketCaptureFilterResponseArrayOutput) ToPacketCaptureFilterResponseArrayOutput() PacketCaptureFilterResponseArrayOutput {
	return o
}

func (o PacketCaptureFilterResponseArrayOutput) ToPacketCaptureFilterResponseArrayOutputWithContext(ctx context.Context) PacketCaptureFilterResponseArrayOutput {
	return o
}

func (o PacketCaptureFilterResponseArrayOutput) Index(i pulumi.IntInput) PacketCaptureFilterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PacketCaptureFilterResponse {
		return vs[0].([]PacketCaptureFilterResponse)[vs[1].(int)]
	}).(PacketCaptureFilterResponseOutput)
}

type PacketCaptureStorageLocation struct {
	FilePath    *string `pulumi:"filePath"`
	StorageId   *string `pulumi:"storageId"`
	StoragePath *string `pulumi:"storagePath"`
}





type PacketCaptureStorageLocationInput interface {
	pulumi.Input

	ToPacketCaptureStorageLocationOutput() PacketCaptureStorageLocationOutput
	ToPacketCaptureStorageLocationOutputWithContext(context.Context) PacketCaptureStorageLocationOutput
}

type PacketCaptureStorageLocationArgs struct {
	FilePath    pulumi.StringPtrInput `pulumi:"filePath"`
	StorageId   pulumi.StringPtrInput `pulumi:"storageId"`
	StoragePath pulumi.StringPtrInput `pulumi:"storagePath"`
}

func (PacketCaptureStorageLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PacketCaptureStorageLocation)(nil)).Elem()
}

func (i PacketCaptureStorageLocationArgs) ToPacketCaptureStorageLocationOutput() PacketCaptureStorageLocationOutput {
	return i.ToPacketCaptureStorageLocationOutputWithContext(context.Background())
}

func (i PacketCaptureStorageLocationArgs) ToPacketCaptureStorageLocationOutputWithContext(ctx context.Context) PacketCaptureStorageLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketCaptureStorageLocationOutput)
}

type PacketCaptureStorageLocationOutput struct{ *pulumi.OutputState }

func (PacketCaptureStorageLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PacketCaptureStorageLocation)(nil)).Elem()
}

func (o PacketCaptureStorageLocationOutput) ToPacketCaptureStorageLocationOutput() PacketCaptureStorageLocationOutput {
	return o
}

func (o PacketCaptureStorageLocationOutput) ToPacketCaptureStorageLocationOutputWithContext(ctx context.Context) PacketCaptureStorageLocationOutput {
	return o
}

func (o PacketCaptureStorageLocationOutput) FilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureStorageLocation) *string { return v.FilePath }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureStorageLocationOutput) StorageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureStorageLocation) *string { return v.StorageId }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureStorageLocationOutput) StoragePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureStorageLocation) *string { return v.StoragePath }).(pulumi.StringPtrOutput)
}

type PacketCaptureStorageLocationResponse struct {
	FilePath    *string `pulumi:"filePath"`
	StorageId   *string `pulumi:"storageId"`
	StoragePath *string `pulumi:"storagePath"`
}

type PacketCaptureStorageLocationResponseOutput struct{ *pulumi.OutputState }

func (PacketCaptureStorageLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PacketCaptureStorageLocationResponse)(nil)).Elem()
}

func (o PacketCaptureStorageLocationResponseOutput) ToPacketCaptureStorageLocationResponseOutput() PacketCaptureStorageLocationResponseOutput {
	return o
}

func (o PacketCaptureStorageLocationResponseOutput) ToPacketCaptureStorageLocationResponseOutputWithContext(ctx context.Context) PacketCaptureStorageLocationResponseOutput {
	return o
}

func (o PacketCaptureStorageLocationResponseOutput) FilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureStorageLocationResponse) *string { return v.FilePath }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureStorageLocationResponseOutput) StorageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureStorageLocationResponse) *string { return v.StorageId }).(pulumi.StringPtrOutput)
}

func (o PacketCaptureStorageLocationResponseOutput) StoragePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PacketCaptureStorageLocationResponse) *string { return v.StoragePath }).(pulumi.StringPtrOutput)
}

type Probe struct {
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	IntervalInSeconds *int    `pulumi:"intervalInSeconds"`
	Name              *string `pulumi:"name"`
	NumberOfProbes    *int    `pulumi:"numberOfProbes"`
	Port              int     `pulumi:"port"`
	Protocol          string  `pulumi:"protocol"`
	ProvisioningState *string `pulumi:"provisioningState"`
	RequestPath       *string `pulumi:"requestPath"`
}





type ProbeInput interface {
	pulumi.Input

	ToProbeOutput() ProbeOutput
	ToProbeOutputWithContext(context.Context) ProbeOutput
}

type ProbeArgs struct {
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	IntervalInSeconds pulumi.IntPtrInput    `pulumi:"intervalInSeconds"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	NumberOfProbes    pulumi.IntPtrInput    `pulumi:"numberOfProbes"`
	Port              pulumi.IntInput       `pulumi:"port"`
	Protocol          pulumi.StringInput    `pulumi:"protocol"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
	RequestPath       pulumi.StringPtrInput `pulumi:"requestPath"`
}

func (ProbeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Probe)(nil)).Elem()
}

func (i ProbeArgs) ToProbeOutput() ProbeOutput {
	return i.ToProbeOutputWithContext(context.Background())
}

func (i ProbeArgs) ToProbeOutputWithContext(ctx context.Context) ProbeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProbeOutput)
}





type ProbeArrayInput interface {
	pulumi.Input

	ToProbeArrayOutput() ProbeArrayOutput
	ToProbeArrayOutputWithContext(context.Context) ProbeArrayOutput
}

type ProbeArray []ProbeInput

func (ProbeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Probe)(nil)).Elem()
}

func (i ProbeArray) ToProbeArrayOutput() ProbeArrayOutput {
	return i.ToProbeArrayOutputWithContext(context.Background())
}

func (i ProbeArray) ToProbeArrayOutputWithContext(ctx context.Context) ProbeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProbeArrayOutput)
}

type ProbeOutput struct{ *pulumi.OutputState }

func (ProbeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Probe)(nil)).Elem()
}

func (o ProbeOutput) ToProbeOutput() ProbeOutput {
	return o
}

func (o ProbeOutput) ToProbeOutputWithContext(ctx context.Context) ProbeOutput {
	return o
}

func (o ProbeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Probe) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ProbeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Probe) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ProbeOutput) IntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Probe) *int { return v.IntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o ProbeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Probe) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ProbeOutput) NumberOfProbes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Probe) *int { return v.NumberOfProbes }).(pulumi.IntPtrOutput)
}

func (o ProbeOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v Probe) int { return v.Port }).(pulumi.IntOutput)
}

func (o ProbeOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v Probe) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o ProbeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Probe) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ProbeOutput) RequestPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Probe) *string { return v.RequestPath }).(pulumi.StringPtrOutput)
}

type ProbeArrayOutput struct{ *pulumi.OutputState }

func (ProbeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Probe)(nil)).Elem()
}

func (o ProbeArrayOutput) ToProbeArrayOutput() ProbeArrayOutput {
	return o
}

func (o ProbeArrayOutput) ToProbeArrayOutputWithContext(ctx context.Context) ProbeArrayOutput {
	return o
}

func (o ProbeArrayOutput) Index(i pulumi.IntInput) ProbeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Probe {
		return vs[0].([]Probe)[vs[1].(int)]
	}).(ProbeOutput)
}

type ProbeResponse struct {
	Etag               *string               `pulumi:"etag"`
	Id                 *string               `pulumi:"id"`
	IntervalInSeconds  *int                  `pulumi:"intervalInSeconds"`
	LoadBalancingRules []SubResourceResponse `pulumi:"loadBalancingRules"`
	Name               *string               `pulumi:"name"`
	NumberOfProbes     *int                  `pulumi:"numberOfProbes"`
	Port               int                   `pulumi:"port"`
	Protocol           string                `pulumi:"protocol"`
	ProvisioningState  *string               `pulumi:"provisioningState"`
	RequestPath        *string               `pulumi:"requestPath"`
}

type ProbeResponseOutput struct{ *pulumi.OutputState }

func (ProbeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProbeResponse)(nil)).Elem()
}

func (o ProbeResponseOutput) ToProbeResponseOutput() ProbeResponseOutput {
	return o
}

func (o ProbeResponseOutput) ToProbeResponseOutputWithContext(ctx context.Context) ProbeResponseOutput {
	return o
}

func (o ProbeResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProbeResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ProbeResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProbeResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ProbeResponseOutput) IntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProbeResponse) *int { return v.IntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o ProbeResponseOutput) LoadBalancingRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v ProbeResponse) []SubResourceResponse { return v.LoadBalancingRules }).(SubResourceResponseArrayOutput)
}

func (o ProbeResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProbeResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ProbeResponseOutput) NumberOfProbes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProbeResponse) *int { return v.NumberOfProbes }).(pulumi.IntPtrOutput)
}

func (o ProbeResponseOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v ProbeResponse) int { return v.Port }).(pulumi.IntOutput)
}

func (o ProbeResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v ProbeResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o ProbeResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProbeResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ProbeResponseOutput) RequestPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProbeResponse) *string { return v.RequestPath }).(pulumi.StringPtrOutput)
}

type ProbeResponseArrayOutput struct{ *pulumi.OutputState }

func (ProbeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProbeResponse)(nil)).Elem()
}

func (o ProbeResponseArrayOutput) ToProbeResponseArrayOutput() ProbeResponseArrayOutput {
	return o
}

func (o ProbeResponseArrayOutput) ToProbeResponseArrayOutputWithContext(ctx context.Context) ProbeResponseArrayOutput {
	return o
}

func (o ProbeResponseArrayOutput) Index(i pulumi.IntInput) ProbeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProbeResponse {
		return vs[0].([]ProbeResponse)[vs[1].(int)]
	}).(ProbeResponseOutput)
}

type PublicIPAddressType struct {
	DnsSettings              *PublicIPAddressDnsSettings `pulumi:"dnsSettings"`
	Etag                     *string                     `pulumi:"etag"`
	Id                       *string                     `pulumi:"id"`
	IdleTimeoutInMinutes     *int                        `pulumi:"idleTimeoutInMinutes"`
	IpAddress                *string                     `pulumi:"ipAddress"`
	Location                 *string                     `pulumi:"location"`
	ProvisioningState        *string                     `pulumi:"provisioningState"`
	PublicIPAddressVersion   *string                     `pulumi:"publicIPAddressVersion"`
	PublicIPAllocationMethod *string                     `pulumi:"publicIPAllocationMethod"`
	ResourceGuid             *string                     `pulumi:"resourceGuid"`
	Tags                     map[string]string           `pulumi:"tags"`
}





type PublicIPAddressTypeInput interface {
	pulumi.Input

	ToPublicIPAddressTypeOutput() PublicIPAddressTypeOutput
	ToPublicIPAddressTypeOutputWithContext(context.Context) PublicIPAddressTypeOutput
}

type PublicIPAddressTypeArgs struct {
	DnsSettings              PublicIPAddressDnsSettingsPtrInput `pulumi:"dnsSettings"`
	Etag                     pulumi.StringPtrInput              `pulumi:"etag"`
	Id                       pulumi.StringPtrInput              `pulumi:"id"`
	IdleTimeoutInMinutes     pulumi.IntPtrInput                 `pulumi:"idleTimeoutInMinutes"`
	IpAddress                pulumi.StringPtrInput              `pulumi:"ipAddress"`
	Location                 pulumi.StringPtrInput              `pulumi:"location"`
	ProvisioningState        pulumi.StringPtrInput              `pulumi:"provisioningState"`
	PublicIPAddressVersion   pulumi.StringPtrInput              `pulumi:"publicIPAddressVersion"`
	PublicIPAllocationMethod pulumi.StringPtrInput              `pulumi:"publicIPAllocationMethod"`
	ResourceGuid             pulumi.StringPtrInput              `pulumi:"resourceGuid"`
	Tags                     pulumi.StringMapInput              `pulumi:"tags"`
}

func (PublicIPAddressTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressType)(nil)).Elem()
}

func (i PublicIPAddressTypeArgs) ToPublicIPAddressTypeOutput() PublicIPAddressTypeOutput {
	return i.ToPublicIPAddressTypeOutputWithContext(context.Background())
}

func (i PublicIPAddressTypeArgs) ToPublicIPAddressTypeOutputWithContext(ctx context.Context) PublicIPAddressTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressTypeOutput)
}

func (i PublicIPAddressTypeArgs) ToPublicIPAddressTypePtrOutput() PublicIPAddressTypePtrOutput {
	return i.ToPublicIPAddressTypePtrOutputWithContext(context.Background())
}

func (i PublicIPAddressTypeArgs) ToPublicIPAddressTypePtrOutputWithContext(ctx context.Context) PublicIPAddressTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressTypeOutput).ToPublicIPAddressTypePtrOutputWithContext(ctx)
}









type PublicIPAddressTypePtrInput interface {
	pulumi.Input

	ToPublicIPAddressTypePtrOutput() PublicIPAddressTypePtrOutput
	ToPublicIPAddressTypePtrOutputWithContext(context.Context) PublicIPAddressTypePtrOutput
}

type publicIPAddressTypePtrType PublicIPAddressTypeArgs

func PublicIPAddressTypePtr(v *PublicIPAddressTypeArgs) PublicIPAddressTypePtrInput {
	return (*publicIPAddressTypePtrType)(v)
}

func (*publicIPAddressTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressType)(nil)).Elem()
}

func (i *publicIPAddressTypePtrType) ToPublicIPAddressTypePtrOutput() PublicIPAddressTypePtrOutput {
	return i.ToPublicIPAddressTypePtrOutputWithContext(context.Background())
}

func (i *publicIPAddressTypePtrType) ToPublicIPAddressTypePtrOutputWithContext(ctx context.Context) PublicIPAddressTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressTypePtrOutput)
}

type PublicIPAddressTypeOutput struct{ *pulumi.OutputState }

func (PublicIPAddressTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressType)(nil)).Elem()
}

func (o PublicIPAddressTypeOutput) ToPublicIPAddressTypeOutput() PublicIPAddressTypeOutput {
	return o
}

func (o PublicIPAddressTypeOutput) ToPublicIPAddressTypeOutputWithContext(ctx context.Context) PublicIPAddressTypeOutput {
	return o
}

func (o PublicIPAddressTypeOutput) ToPublicIPAddressTypePtrOutput() PublicIPAddressTypePtrOutput {
	return o.ToPublicIPAddressTypePtrOutputWithContext(context.Background())
}

func (o PublicIPAddressTypeOutput) ToPublicIPAddressTypePtrOutputWithContext(ctx context.Context) PublicIPAddressTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIPAddressType) *PublicIPAddressType {
		return &v
	}).(PublicIPAddressTypePtrOutput)
}

func (o PublicIPAddressTypeOutput) DnsSettings() PublicIPAddressDnsSettingsPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *PublicIPAddressDnsSettings { return v.DnsSettings }).(PublicIPAddressDnsSettingsPtrOutput)
}

func (o PublicIPAddressTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypeOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o PublicIPAddressTypeOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypeOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *string { return v.PublicIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypeOutput) PublicIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *string { return v.PublicIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypeOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressType) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v PublicIPAddressType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type PublicIPAddressTypePtrOutput struct{ *pulumi.OutputState }

func (PublicIPAddressTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressType)(nil)).Elem()
}

func (o PublicIPAddressTypePtrOutput) ToPublicIPAddressTypePtrOutput() PublicIPAddressTypePtrOutput {
	return o
}

func (o PublicIPAddressTypePtrOutput) ToPublicIPAddressTypePtrOutputWithContext(ctx context.Context) PublicIPAddressTypePtrOutput {
	return o
}

func (o PublicIPAddressTypePtrOutput) Elem() PublicIPAddressTypeOutput {
	return o.ApplyT(func(v *PublicIPAddressType) PublicIPAddressType {
		if v != nil {
			return *v
		}
		var ret PublicIPAddressType
		return ret
	}).(PublicIPAddressTypeOutput)
}

func (o PublicIPAddressTypePtrOutput) DnsSettings() PublicIPAddressDnsSettingsPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *PublicIPAddressDnsSettings {
		if v == nil {
			return nil
		}
		return v.DnsSettings
	}).(PublicIPAddressDnsSettingsPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *int {
		if v == nil {
			return nil
		}
		return v.IdleTimeoutInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *string {
		if v == nil {
			return nil
		}
		return v.IpAddress
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *string {
		if v == nil {
			return nil
		}
		return v.PublicIPAddressVersion
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) PublicIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *string {
		if v == nil {
			return nil
		}
		return v.PublicIPAllocationMethod
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressType) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressTypePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PublicIPAddressType) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

type PublicIPAddressDnsSettings struct {
	DomainNameLabel *string `pulumi:"domainNameLabel"`
	Fqdn            *string `pulumi:"fqdn"`
	ReverseFqdn     *string `pulumi:"reverseFqdn"`
}





type PublicIPAddressDnsSettingsInput interface {
	pulumi.Input

	ToPublicIPAddressDnsSettingsOutput() PublicIPAddressDnsSettingsOutput
	ToPublicIPAddressDnsSettingsOutputWithContext(context.Context) PublicIPAddressDnsSettingsOutput
}

type PublicIPAddressDnsSettingsArgs struct {
	DomainNameLabel pulumi.StringPtrInput `pulumi:"domainNameLabel"`
	Fqdn            pulumi.StringPtrInput `pulumi:"fqdn"`
	ReverseFqdn     pulumi.StringPtrInput `pulumi:"reverseFqdn"`
}

func (PublicIPAddressDnsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressDnsSettings)(nil)).Elem()
}

func (i PublicIPAddressDnsSettingsArgs) ToPublicIPAddressDnsSettingsOutput() PublicIPAddressDnsSettingsOutput {
	return i.ToPublicIPAddressDnsSettingsOutputWithContext(context.Background())
}

func (i PublicIPAddressDnsSettingsArgs) ToPublicIPAddressDnsSettingsOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressDnsSettingsOutput)
}

func (i PublicIPAddressDnsSettingsArgs) ToPublicIPAddressDnsSettingsPtrOutput() PublicIPAddressDnsSettingsPtrOutput {
	return i.ToPublicIPAddressDnsSettingsPtrOutputWithContext(context.Background())
}

func (i PublicIPAddressDnsSettingsArgs) ToPublicIPAddressDnsSettingsPtrOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressDnsSettingsOutput).ToPublicIPAddressDnsSettingsPtrOutputWithContext(ctx)
}









type PublicIPAddressDnsSettingsPtrInput interface {
	pulumi.Input

	ToPublicIPAddressDnsSettingsPtrOutput() PublicIPAddressDnsSettingsPtrOutput
	ToPublicIPAddressDnsSettingsPtrOutputWithContext(context.Context) PublicIPAddressDnsSettingsPtrOutput
}

type publicIPAddressDnsSettingsPtrType PublicIPAddressDnsSettingsArgs

func PublicIPAddressDnsSettingsPtr(v *PublicIPAddressDnsSettingsArgs) PublicIPAddressDnsSettingsPtrInput {
	return (*publicIPAddressDnsSettingsPtrType)(v)
}

func (*publicIPAddressDnsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressDnsSettings)(nil)).Elem()
}

func (i *publicIPAddressDnsSettingsPtrType) ToPublicIPAddressDnsSettingsPtrOutput() PublicIPAddressDnsSettingsPtrOutput {
	return i.ToPublicIPAddressDnsSettingsPtrOutputWithContext(context.Background())
}

func (i *publicIPAddressDnsSettingsPtrType) ToPublicIPAddressDnsSettingsPtrOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressDnsSettingsPtrOutput)
}

type PublicIPAddressDnsSettingsOutput struct{ *pulumi.OutputState }

func (PublicIPAddressDnsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressDnsSettings)(nil)).Elem()
}

func (o PublicIPAddressDnsSettingsOutput) ToPublicIPAddressDnsSettingsOutput() PublicIPAddressDnsSettingsOutput {
	return o
}

func (o PublicIPAddressDnsSettingsOutput) ToPublicIPAddressDnsSettingsOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsOutput {
	return o
}

func (o PublicIPAddressDnsSettingsOutput) ToPublicIPAddressDnsSettingsPtrOutput() PublicIPAddressDnsSettingsPtrOutput {
	return o.ToPublicIPAddressDnsSettingsPtrOutputWithContext(context.Background())
}

func (o PublicIPAddressDnsSettingsOutput) ToPublicIPAddressDnsSettingsPtrOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIPAddressDnsSettings) *PublicIPAddressDnsSettings {
		return &v
	}).(PublicIPAddressDnsSettingsPtrOutput)
}

func (o PublicIPAddressDnsSettingsOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressDnsSettings) *string { return v.DomainNameLabel }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressDnsSettingsOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressDnsSettings) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressDnsSettingsOutput) ReverseFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressDnsSettings) *string { return v.ReverseFqdn }).(pulumi.StringPtrOutput)
}

type PublicIPAddressDnsSettingsPtrOutput struct{ *pulumi.OutputState }

func (PublicIPAddressDnsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressDnsSettings)(nil)).Elem()
}

func (o PublicIPAddressDnsSettingsPtrOutput) ToPublicIPAddressDnsSettingsPtrOutput() PublicIPAddressDnsSettingsPtrOutput {
	return o
}

func (o PublicIPAddressDnsSettingsPtrOutput) ToPublicIPAddressDnsSettingsPtrOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsPtrOutput {
	return o
}

func (o PublicIPAddressDnsSettingsPtrOutput) Elem() PublicIPAddressDnsSettingsOutput {
	return o.ApplyT(func(v *PublicIPAddressDnsSettings) PublicIPAddressDnsSettings {
		if v != nil {
			return *v
		}
		var ret PublicIPAddressDnsSettings
		return ret
	}).(PublicIPAddressDnsSettingsOutput)
}

func (o PublicIPAddressDnsSettingsPtrOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressDnsSettings) *string {
		if v == nil {
			return nil
		}
		return v.DomainNameLabel
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressDnsSettingsPtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressDnsSettings) *string {
		if v == nil {
			return nil
		}
		return v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressDnsSettingsPtrOutput) ReverseFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressDnsSettings) *string {
		if v == nil {
			return nil
		}
		return v.ReverseFqdn
	}).(pulumi.StringPtrOutput)
}

type PublicIPAddressDnsSettingsResponse struct {
	DomainNameLabel *string `pulumi:"domainNameLabel"`
	Fqdn            *string `pulumi:"fqdn"`
	ReverseFqdn     *string `pulumi:"reverseFqdn"`
}

type PublicIPAddressDnsSettingsResponseOutput struct{ *pulumi.OutputState }

func (PublicIPAddressDnsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressDnsSettingsResponse)(nil)).Elem()
}

func (o PublicIPAddressDnsSettingsResponseOutput) ToPublicIPAddressDnsSettingsResponseOutput() PublicIPAddressDnsSettingsResponseOutput {
	return o
}

func (o PublicIPAddressDnsSettingsResponseOutput) ToPublicIPAddressDnsSettingsResponseOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsResponseOutput {
	return o
}

func (o PublicIPAddressDnsSettingsResponseOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressDnsSettingsResponse) *string { return v.DomainNameLabel }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressDnsSettingsResponseOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressDnsSettingsResponse) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressDnsSettingsResponseOutput) ReverseFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressDnsSettingsResponse) *string { return v.ReverseFqdn }).(pulumi.StringPtrOutput)
}

type PublicIPAddressDnsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (PublicIPAddressDnsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressDnsSettingsResponse)(nil)).Elem()
}

func (o PublicIPAddressDnsSettingsResponsePtrOutput) ToPublicIPAddressDnsSettingsResponsePtrOutput() PublicIPAddressDnsSettingsResponsePtrOutput {
	return o
}

func (o PublicIPAddressDnsSettingsResponsePtrOutput) ToPublicIPAddressDnsSettingsResponsePtrOutputWithContext(ctx context.Context) PublicIPAddressDnsSettingsResponsePtrOutput {
	return o
}

func (o PublicIPAddressDnsSettingsResponsePtrOutput) Elem() PublicIPAddressDnsSettingsResponseOutput {
	return o.ApplyT(func(v *PublicIPAddressDnsSettingsResponse) PublicIPAddressDnsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret PublicIPAddressDnsSettingsResponse
		return ret
	}).(PublicIPAddressDnsSettingsResponseOutput)
}

func (o PublicIPAddressDnsSettingsResponsePtrOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressDnsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DomainNameLabel
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressDnsSettingsResponsePtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressDnsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressDnsSettingsResponsePtrOutput) ReverseFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressDnsSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReverseFqdn
	}).(pulumi.StringPtrOutput)
}

type PublicIPAddressResponse struct {
	DnsSettings              *PublicIPAddressDnsSettingsResponse `pulumi:"dnsSettings"`
	Etag                     *string                             `pulumi:"etag"`
	Id                       *string                             `pulumi:"id"`
	IdleTimeoutInMinutes     *int                                `pulumi:"idleTimeoutInMinutes"`
	IpAddress                *string                             `pulumi:"ipAddress"`
	IpConfiguration          IPConfigurationResponse             `pulumi:"ipConfiguration"`
	Location                 *string                             `pulumi:"location"`
	Name                     string                              `pulumi:"name"`
	ProvisioningState        *string                             `pulumi:"provisioningState"`
	PublicIPAddressVersion   *string                             `pulumi:"publicIPAddressVersion"`
	PublicIPAllocationMethod *string                             `pulumi:"publicIPAllocationMethod"`
	ResourceGuid             *string                             `pulumi:"resourceGuid"`
	Tags                     map[string]string                   `pulumi:"tags"`
	Type                     string                              `pulumi:"type"`
}

type PublicIPAddressResponseOutput struct{ *pulumi.OutputState }

func (PublicIPAddressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressResponse)(nil)).Elem()
}

func (o PublicIPAddressResponseOutput) ToPublicIPAddressResponseOutput() PublicIPAddressResponseOutput {
	return o
}

func (o PublicIPAddressResponseOutput) ToPublicIPAddressResponseOutputWithContext(ctx context.Context) PublicIPAddressResponseOutput {
	return o
}

func (o PublicIPAddressResponseOutput) DnsSettings() PublicIPAddressDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *PublicIPAddressDnsSettingsResponse { return v.DnsSettings }).(PublicIPAddressDnsSettingsResponsePtrOutput)
}

func (o PublicIPAddressResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponseOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o PublicIPAddressResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponseOutput) IpConfiguration() IPConfigurationResponseOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) IPConfigurationResponse { return v.IpConfiguration }).(IPConfigurationResponseOutput)
}

func (o PublicIPAddressResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PublicIPAddressResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponseOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *string { return v.PublicIPAddressVersion }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponseOutput) PublicIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *string { return v.PublicIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponseOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PublicIPAddressResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PublicIPAddressResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PublicIPAddressResponsePtrOutput struct{ *pulumi.OutputState }

func (PublicIPAddressResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPAddressResponse)(nil)).Elem()
}

func (o PublicIPAddressResponsePtrOutput) ToPublicIPAddressResponsePtrOutput() PublicIPAddressResponsePtrOutput {
	return o
}

func (o PublicIPAddressResponsePtrOutput) ToPublicIPAddressResponsePtrOutputWithContext(ctx context.Context) PublicIPAddressResponsePtrOutput {
	return o
}

func (o PublicIPAddressResponsePtrOutput) Elem() PublicIPAddressResponseOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) PublicIPAddressResponse {
		if v != nil {
			return *v
		}
		var ret PublicIPAddressResponse
		return ret
	}).(PublicIPAddressResponseOutput)
}

func (o PublicIPAddressResponsePtrOutput) DnsSettings() PublicIPAddressDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *PublicIPAddressDnsSettingsResponse {
		if v == nil {
			return nil
		}
		return v.DnsSettings
	}).(PublicIPAddressDnsSettingsResponsePtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *int {
		if v == nil {
			return nil
		}
		return v.IdleTimeoutInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.IpAddress
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) IpConfiguration() IPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *IPConfigurationResponse {
		if v == nil {
			return nil
		}
		return &v.IpConfiguration
	}).(IPConfigurationResponsePtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicIPAddressVersion
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) PublicIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicIPAllocationMethod
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o PublicIPAddressResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ResourceNavigationLink struct {
	Id                 *string `pulumi:"id"`
	Link               *string `pulumi:"link"`
	LinkedResourceType *string `pulumi:"linkedResourceType"`
	Name               *string `pulumi:"name"`
}





type ResourceNavigationLinkInput interface {
	pulumi.Input

	ToResourceNavigationLinkOutput() ResourceNavigationLinkOutput
	ToResourceNavigationLinkOutputWithContext(context.Context) ResourceNavigationLinkOutput
}

type ResourceNavigationLinkArgs struct {
	Id                 pulumi.StringPtrInput `pulumi:"id"`
	Link               pulumi.StringPtrInput `pulumi:"link"`
	LinkedResourceType pulumi.StringPtrInput `pulumi:"linkedResourceType"`
	Name               pulumi.StringPtrInput `pulumi:"name"`
}

func (ResourceNavigationLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceNavigationLink)(nil)).Elem()
}

func (i ResourceNavigationLinkArgs) ToResourceNavigationLinkOutput() ResourceNavigationLinkOutput {
	return i.ToResourceNavigationLinkOutputWithContext(context.Background())
}

func (i ResourceNavigationLinkArgs) ToResourceNavigationLinkOutputWithContext(ctx context.Context) ResourceNavigationLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceNavigationLinkOutput)
}





type ResourceNavigationLinkArrayInput interface {
	pulumi.Input

	ToResourceNavigationLinkArrayOutput() ResourceNavigationLinkArrayOutput
	ToResourceNavigationLinkArrayOutputWithContext(context.Context) ResourceNavigationLinkArrayOutput
}

type ResourceNavigationLinkArray []ResourceNavigationLinkInput

func (ResourceNavigationLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceNavigationLink)(nil)).Elem()
}

func (i ResourceNavigationLinkArray) ToResourceNavigationLinkArrayOutput() ResourceNavigationLinkArrayOutput {
	return i.ToResourceNavigationLinkArrayOutputWithContext(context.Background())
}

func (i ResourceNavigationLinkArray) ToResourceNavigationLinkArrayOutputWithContext(ctx context.Context) ResourceNavigationLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceNavigationLinkArrayOutput)
}

type ResourceNavigationLinkOutput struct{ *pulumi.OutputState }

func (ResourceNavigationLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceNavigationLink)(nil)).Elem()
}

func (o ResourceNavigationLinkOutput) ToResourceNavigationLinkOutput() ResourceNavigationLinkOutput {
	return o
}

func (o ResourceNavigationLinkOutput) ToResourceNavigationLinkOutputWithContext(ctx context.Context) ResourceNavigationLinkOutput {
	return o
}

func (o ResourceNavigationLinkOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceNavigationLink) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ResourceNavigationLinkOutput) Link() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceNavigationLink) *string { return v.Link }).(pulumi.StringPtrOutput)
}

func (o ResourceNavigationLinkOutput) LinkedResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceNavigationLink) *string { return v.LinkedResourceType }).(pulumi.StringPtrOutput)
}

func (o ResourceNavigationLinkOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceNavigationLink) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ResourceNavigationLinkArrayOutput struct{ *pulumi.OutputState }

func (ResourceNavigationLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceNavigationLink)(nil)).Elem()
}

func (o ResourceNavigationLinkArrayOutput) ToResourceNavigationLinkArrayOutput() ResourceNavigationLinkArrayOutput {
	return o
}

func (o ResourceNavigationLinkArrayOutput) ToResourceNavigationLinkArrayOutputWithContext(ctx context.Context) ResourceNavigationLinkArrayOutput {
	return o
}

func (o ResourceNavigationLinkArrayOutput) Index(i pulumi.IntInput) ResourceNavigationLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceNavigationLink {
		return vs[0].([]ResourceNavigationLink)[vs[1].(int)]
	}).(ResourceNavigationLinkOutput)
}

type ResourceNavigationLinkResponse struct {
	Etag               string  `pulumi:"etag"`
	Id                 *string `pulumi:"id"`
	Link               *string `pulumi:"link"`
	LinkedResourceType *string `pulumi:"linkedResourceType"`
	Name               *string `pulumi:"name"`
	ProvisioningState  string  `pulumi:"provisioningState"`
}

type ResourceNavigationLinkResponseOutput struct{ *pulumi.OutputState }

func (ResourceNavigationLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceNavigationLinkResponse)(nil)).Elem()
}

func (o ResourceNavigationLinkResponseOutput) ToResourceNavigationLinkResponseOutput() ResourceNavigationLinkResponseOutput {
	return o
}

func (o ResourceNavigationLinkResponseOutput) ToResourceNavigationLinkResponseOutputWithContext(ctx context.Context) ResourceNavigationLinkResponseOutput {
	return o
}

func (o ResourceNavigationLinkResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceNavigationLinkResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o ResourceNavigationLinkResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceNavigationLinkResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ResourceNavigationLinkResponseOutput) Link() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceNavigationLinkResponse) *string { return v.Link }).(pulumi.StringPtrOutput)
}

func (o ResourceNavigationLinkResponseOutput) LinkedResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceNavigationLinkResponse) *string { return v.LinkedResourceType }).(pulumi.StringPtrOutput)
}

func (o ResourceNavigationLinkResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceNavigationLinkResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ResourceNavigationLinkResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceNavigationLinkResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type ResourceNavigationLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceNavigationLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceNavigationLinkResponse)(nil)).Elem()
}

func (o ResourceNavigationLinkResponseArrayOutput) ToResourceNavigationLinkResponseArrayOutput() ResourceNavigationLinkResponseArrayOutput {
	return o
}

func (o ResourceNavigationLinkResponseArrayOutput) ToResourceNavigationLinkResponseArrayOutputWithContext(ctx context.Context) ResourceNavigationLinkResponseArrayOutput {
	return o
}

func (o ResourceNavigationLinkResponseArrayOutput) Index(i pulumi.IntInput) ResourceNavigationLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceNavigationLinkResponse {
		return vs[0].([]ResourceNavigationLinkResponse)[vs[1].(int)]
	}).(ResourceNavigationLinkResponseOutput)
}

type RouteType struct {
	AddressPrefix     *string `pulumi:"addressPrefix"`
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	NextHopIpAddress  *string `pulumi:"nextHopIpAddress"`
	NextHopType       string  `pulumi:"nextHopType"`
	ProvisioningState *string `pulumi:"provisioningState"`
}





type RouteTypeInput interface {
	pulumi.Input

	ToRouteTypeOutput() RouteTypeOutput
	ToRouteTypeOutputWithContext(context.Context) RouteTypeOutput
}

type RouteTypeArgs struct {
	AddressPrefix     pulumi.StringPtrInput `pulumi:"addressPrefix"`
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	NextHopIpAddress  pulumi.StringPtrInput `pulumi:"nextHopIpAddress"`
	NextHopType       pulumi.StringInput    `pulumi:"nextHopType"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (RouteTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteType)(nil)).Elem()
}

func (i RouteTypeArgs) ToRouteTypeOutput() RouteTypeOutput {
	return i.ToRouteTypeOutputWithContext(context.Background())
}

func (i RouteTypeArgs) ToRouteTypeOutputWithContext(ctx context.Context) RouteTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTypeOutput)
}





type RouteTypeArrayInput interface {
	pulumi.Input

	ToRouteTypeArrayOutput() RouteTypeArrayOutput
	ToRouteTypeArrayOutputWithContext(context.Context) RouteTypeArrayOutput
}

type RouteTypeArray []RouteTypeInput

func (RouteTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteType)(nil)).Elem()
}

func (i RouteTypeArray) ToRouteTypeArrayOutput() RouteTypeArrayOutput {
	return i.ToRouteTypeArrayOutputWithContext(context.Background())
}

func (i RouteTypeArray) ToRouteTypeArrayOutputWithContext(ctx context.Context) RouteTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTypeArrayOutput)
}

type RouteTypeOutput struct{ *pulumi.OutputState }

func (RouteTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteType)(nil)).Elem()
}

func (o RouteTypeOutput) ToRouteTypeOutput() RouteTypeOutput {
	return o
}

func (o RouteTypeOutput) ToRouteTypeOutputWithContext(ctx context.Context) RouteTypeOutput {
	return o
}

func (o RouteTypeOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteType) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o RouteTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o RouteTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RouteTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RouteTypeOutput) NextHopIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteType) *string { return v.NextHopIpAddress }).(pulumi.StringPtrOutput)
}

func (o RouteTypeOutput) NextHopType() pulumi.StringOutput {
	return o.ApplyT(func(v RouteType) string { return v.NextHopType }).(pulumi.StringOutput)
}

func (o RouteTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type RouteTypeArrayOutput struct{ *pulumi.OutputState }

func (RouteTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteType)(nil)).Elem()
}

func (o RouteTypeArrayOutput) ToRouteTypeArrayOutput() RouteTypeArrayOutput {
	return o
}

func (o RouteTypeArrayOutput) ToRouteTypeArrayOutputWithContext(ctx context.Context) RouteTypeArrayOutput {
	return o
}

func (o RouteTypeArrayOutput) Index(i pulumi.IntInput) RouteTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RouteType {
		return vs[0].([]RouteType)[vs[1].(int)]
	}).(RouteTypeOutput)
}

type RouteFilterType struct {
	Id       *string               `pulumi:"id"`
	Location string                `pulumi:"location"`
	Rules    []RouteFilterRuleType `pulumi:"rules"`
	Tags     map[string]string     `pulumi:"tags"`
}





type RouteFilterTypeInput interface {
	pulumi.Input

	ToRouteFilterTypeOutput() RouteFilterTypeOutput
	ToRouteFilterTypeOutputWithContext(context.Context) RouteFilterTypeOutput
}

type RouteFilterTypeArgs struct {
	Id       pulumi.StringPtrInput         `pulumi:"id"`
	Location pulumi.StringInput            `pulumi:"location"`
	Rules    RouteFilterRuleTypeArrayInput `pulumi:"rules"`
	Tags     pulumi.StringMapInput         `pulumi:"tags"`
}

func (RouteFilterTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteFilterType)(nil)).Elem()
}

func (i RouteFilterTypeArgs) ToRouteFilterTypeOutput() RouteFilterTypeOutput {
	return i.ToRouteFilterTypeOutputWithContext(context.Background())
}

func (i RouteFilterTypeArgs) ToRouteFilterTypeOutputWithContext(ctx context.Context) RouteFilterTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterTypeOutput)
}

func (i RouteFilterTypeArgs) ToRouteFilterTypePtrOutput() RouteFilterTypePtrOutput {
	return i.ToRouteFilterTypePtrOutputWithContext(context.Background())
}

func (i RouteFilterTypeArgs) ToRouteFilterTypePtrOutputWithContext(ctx context.Context) RouteFilterTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterTypeOutput).ToRouteFilterTypePtrOutputWithContext(ctx)
}









type RouteFilterTypePtrInput interface {
	pulumi.Input

	ToRouteFilterTypePtrOutput() RouteFilterTypePtrOutput
	ToRouteFilterTypePtrOutputWithContext(context.Context) RouteFilterTypePtrOutput
}

type routeFilterTypePtrType RouteFilterTypeArgs

func RouteFilterTypePtr(v *RouteFilterTypeArgs) RouteFilterTypePtrInput {
	return (*routeFilterTypePtrType)(v)
}

func (*routeFilterTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteFilterType)(nil)).Elem()
}

func (i *routeFilterTypePtrType) ToRouteFilterTypePtrOutput() RouteFilterTypePtrOutput {
	return i.ToRouteFilterTypePtrOutputWithContext(context.Background())
}

func (i *routeFilterTypePtrType) ToRouteFilterTypePtrOutputWithContext(ctx context.Context) RouteFilterTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterTypePtrOutput)
}

type RouteFilterTypeOutput struct{ *pulumi.OutputState }

func (RouteFilterTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteFilterType)(nil)).Elem()
}

func (o RouteFilterTypeOutput) ToRouteFilterTypeOutput() RouteFilterTypeOutput {
	return o
}

func (o RouteFilterTypeOutput) ToRouteFilterTypeOutputWithContext(ctx context.Context) RouteFilterTypeOutput {
	return o
}

func (o RouteFilterTypeOutput) ToRouteFilterTypePtrOutput() RouteFilterTypePtrOutput {
	return o.ToRouteFilterTypePtrOutputWithContext(context.Background())
}

func (o RouteFilterTypeOutput) ToRouteFilterTypePtrOutputWithContext(ctx context.Context) RouteFilterTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RouteFilterType) *RouteFilterType {
		return &v
	}).(RouteFilterTypePtrOutput)
}

func (o RouteFilterTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteFilterType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RouteFilterTypeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterType) string { return v.Location }).(pulumi.StringOutput)
}

func (o RouteFilterTypeOutput) Rules() RouteFilterRuleTypeArrayOutput {
	return o.ApplyT(func(v RouteFilterType) []RouteFilterRuleType { return v.Rules }).(RouteFilterRuleTypeArrayOutput)
}

func (o RouteFilterTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v RouteFilterType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type RouteFilterTypePtrOutput struct{ *pulumi.OutputState }

func (RouteFilterTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteFilterType)(nil)).Elem()
}

func (o RouteFilterTypePtrOutput) ToRouteFilterTypePtrOutput() RouteFilterTypePtrOutput {
	return o
}

func (o RouteFilterTypePtrOutput) ToRouteFilterTypePtrOutputWithContext(ctx context.Context) RouteFilterTypePtrOutput {
	return o
}

func (o RouteFilterTypePtrOutput) Elem() RouteFilterTypeOutput {
	return o.ApplyT(func(v *RouteFilterType) RouteFilterType {
		if v != nil {
			return *v
		}
		var ret RouteFilterType
		return ret
	}).(RouteFilterTypeOutput)
}

func (o RouteFilterTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteFilterType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o RouteFilterTypePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteFilterType) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o RouteFilterTypePtrOutput) Rules() RouteFilterRuleTypeArrayOutput {
	return o.ApplyT(func(v *RouteFilterType) []RouteFilterRuleType {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(RouteFilterRuleTypeArrayOutput)
}

func (o RouteFilterTypePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RouteFilterType) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

type RouteFilterResponse struct {
	Etag              string                               `pulumi:"etag"`
	Id                *string                              `pulumi:"id"`
	Location          string                               `pulumi:"location"`
	Name              string                               `pulumi:"name"`
	Peerings          []ExpressRouteCircuitPeeringResponse `pulumi:"peerings"`
	ProvisioningState string                               `pulumi:"provisioningState"`
	Rules             []RouteFilterRuleResponse            `pulumi:"rules"`
	Tags              map[string]string                    `pulumi:"tags"`
	Type              string                               `pulumi:"type"`
}

type RouteFilterResponseOutput struct{ *pulumi.OutputState }

func (RouteFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteFilterResponse)(nil)).Elem()
}

func (o RouteFilterResponseOutput) ToRouteFilterResponseOutput() RouteFilterResponseOutput {
	return o
}

func (o RouteFilterResponseOutput) ToRouteFilterResponseOutputWithContext(ctx context.Context) RouteFilterResponseOutput {
	return o
}

func (o RouteFilterResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o RouteFilterResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteFilterResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RouteFilterResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o RouteFilterResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RouteFilterResponseOutput) Peerings() ExpressRouteCircuitPeeringResponseArrayOutput {
	return o.ApplyT(func(v RouteFilterResponse) []ExpressRouteCircuitPeeringResponse { return v.Peerings }).(ExpressRouteCircuitPeeringResponseArrayOutput)
}

func (o RouteFilterResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RouteFilterResponseOutput) Rules() RouteFilterRuleResponseArrayOutput {
	return o.ApplyT(func(v RouteFilterResponse) []RouteFilterRuleResponse { return v.Rules }).(RouteFilterRuleResponseArrayOutput)
}

func (o RouteFilterResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v RouteFilterResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o RouteFilterResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterResponse) string { return v.Type }).(pulumi.StringOutput)
}

type RouteFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (RouteFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteFilterResponse)(nil)).Elem()
}

func (o RouteFilterResponsePtrOutput) ToRouteFilterResponsePtrOutput() RouteFilterResponsePtrOutput {
	return o
}

func (o RouteFilterResponsePtrOutput) ToRouteFilterResponsePtrOutputWithContext(ctx context.Context) RouteFilterResponsePtrOutput {
	return o
}

func (o RouteFilterResponsePtrOutput) Elem() RouteFilterResponseOutput {
	return o.ApplyT(func(v *RouteFilterResponse) RouteFilterResponse {
		if v != nil {
			return *v
		}
		var ret RouteFilterResponse
		return ret
	}).(RouteFilterResponseOutput)
}

func (o RouteFilterResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteFilterResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o RouteFilterResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o RouteFilterResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteFilterResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o RouteFilterResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteFilterResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o RouteFilterResponsePtrOutput) Peerings() ExpressRouteCircuitPeeringResponseArrayOutput {
	return o.ApplyT(func(v *RouteFilterResponse) []ExpressRouteCircuitPeeringResponse {
		if v == nil {
			return nil
		}
		return v.Peerings
	}).(ExpressRouteCircuitPeeringResponseArrayOutput)
}

func (o RouteFilterResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteFilterResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o RouteFilterResponsePtrOutput) Rules() RouteFilterRuleResponseArrayOutput {
	return o.ApplyT(func(v *RouteFilterResponse) []RouteFilterRuleResponse {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(RouteFilterRuleResponseArrayOutput)
}

func (o RouteFilterResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RouteFilterResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o RouteFilterResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteFilterResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type RouteFilterRuleType struct {
	Access              string            `pulumi:"access"`
	Communities         []string          `pulumi:"communities"`
	Id                  *string           `pulumi:"id"`
	Location            *string           `pulumi:"location"`
	RouteFilterRuleType string            `pulumi:"routeFilterRuleType"`
	Tags                map[string]string `pulumi:"tags"`
}





type RouteFilterRuleTypeInput interface {
	pulumi.Input

	ToRouteFilterRuleTypeOutput() RouteFilterRuleTypeOutput
	ToRouteFilterRuleTypeOutputWithContext(context.Context) RouteFilterRuleTypeOutput
}

type RouteFilterRuleTypeArgs struct {
	Access              pulumi.StringInput      `pulumi:"access"`
	Communities         pulumi.StringArrayInput `pulumi:"communities"`
	Id                  pulumi.StringPtrInput   `pulumi:"id"`
	Location            pulumi.StringPtrInput   `pulumi:"location"`
	RouteFilterRuleType pulumi.StringInput      `pulumi:"routeFilterRuleType"`
	Tags                pulumi.StringMapInput   `pulumi:"tags"`
}

func (RouteFilterRuleTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteFilterRuleType)(nil)).Elem()
}

func (i RouteFilterRuleTypeArgs) ToRouteFilterRuleTypeOutput() RouteFilterRuleTypeOutput {
	return i.ToRouteFilterRuleTypeOutputWithContext(context.Background())
}

func (i RouteFilterRuleTypeArgs) ToRouteFilterRuleTypeOutputWithContext(ctx context.Context) RouteFilterRuleTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterRuleTypeOutput)
}





type RouteFilterRuleTypeArrayInput interface {
	pulumi.Input

	ToRouteFilterRuleTypeArrayOutput() RouteFilterRuleTypeArrayOutput
	ToRouteFilterRuleTypeArrayOutputWithContext(context.Context) RouteFilterRuleTypeArrayOutput
}

type RouteFilterRuleTypeArray []RouteFilterRuleTypeInput

func (RouteFilterRuleTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteFilterRuleType)(nil)).Elem()
}

func (i RouteFilterRuleTypeArray) ToRouteFilterRuleTypeArrayOutput() RouteFilterRuleTypeArrayOutput {
	return i.ToRouteFilterRuleTypeArrayOutputWithContext(context.Background())
}

func (i RouteFilterRuleTypeArray) ToRouteFilterRuleTypeArrayOutputWithContext(ctx context.Context) RouteFilterRuleTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteFilterRuleTypeArrayOutput)
}

type RouteFilterRuleTypeOutput struct{ *pulumi.OutputState }

func (RouteFilterRuleTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteFilterRuleType)(nil)).Elem()
}

func (o RouteFilterRuleTypeOutput) ToRouteFilterRuleTypeOutput() RouteFilterRuleTypeOutput {
	return o
}

func (o RouteFilterRuleTypeOutput) ToRouteFilterRuleTypeOutputWithContext(ctx context.Context) RouteFilterRuleTypeOutput {
	return o
}

func (o RouteFilterRuleTypeOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterRuleType) string { return v.Access }).(pulumi.StringOutput)
}

func (o RouteFilterRuleTypeOutput) Communities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RouteFilterRuleType) []string { return v.Communities }).(pulumi.StringArrayOutput)
}

func (o RouteFilterRuleTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteFilterRuleType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RouteFilterRuleTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteFilterRuleType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o RouteFilterRuleTypeOutput) RouteFilterRuleType() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterRuleType) string { return v.RouteFilterRuleType }).(pulumi.StringOutput)
}

func (o RouteFilterRuleTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v RouteFilterRuleType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type RouteFilterRuleTypeArrayOutput struct{ *pulumi.OutputState }

func (RouteFilterRuleTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteFilterRuleType)(nil)).Elem()
}

func (o RouteFilterRuleTypeArrayOutput) ToRouteFilterRuleTypeArrayOutput() RouteFilterRuleTypeArrayOutput {
	return o
}

func (o RouteFilterRuleTypeArrayOutput) ToRouteFilterRuleTypeArrayOutputWithContext(ctx context.Context) RouteFilterRuleTypeArrayOutput {
	return o
}

func (o RouteFilterRuleTypeArrayOutput) Index(i pulumi.IntInput) RouteFilterRuleTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RouteFilterRuleType {
		return vs[0].([]RouteFilterRuleType)[vs[1].(int)]
	}).(RouteFilterRuleTypeOutput)
}

type RouteFilterRuleResponse struct {
	Access              string            `pulumi:"access"`
	Communities         []string          `pulumi:"communities"`
	Etag                string            `pulumi:"etag"`
	Id                  *string           `pulumi:"id"`
	Location            *string           `pulumi:"location"`
	Name                string            `pulumi:"name"`
	ProvisioningState   string            `pulumi:"provisioningState"`
	RouteFilterRuleType string            `pulumi:"routeFilterRuleType"`
	Tags                map[string]string `pulumi:"tags"`
}

type RouteFilterRuleResponseOutput struct{ *pulumi.OutputState }

func (RouteFilterRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteFilterRuleResponse)(nil)).Elem()
}

func (o RouteFilterRuleResponseOutput) ToRouteFilterRuleResponseOutput() RouteFilterRuleResponseOutput {
	return o
}

func (o RouteFilterRuleResponseOutput) ToRouteFilterRuleResponseOutputWithContext(ctx context.Context) RouteFilterRuleResponseOutput {
	return o
}

func (o RouteFilterRuleResponseOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterRuleResponse) string { return v.Access }).(pulumi.StringOutput)
}

func (o RouteFilterRuleResponseOutput) Communities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RouteFilterRuleResponse) []string { return v.Communities }).(pulumi.StringArrayOutput)
}

func (o RouteFilterRuleResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterRuleResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o RouteFilterRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteFilterRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RouteFilterRuleResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteFilterRuleResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o RouteFilterRuleResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterRuleResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RouteFilterRuleResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterRuleResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RouteFilterRuleResponseOutput) RouteFilterRuleType() pulumi.StringOutput {
	return o.ApplyT(func(v RouteFilterRuleResponse) string { return v.RouteFilterRuleType }).(pulumi.StringOutput)
}

func (o RouteFilterRuleResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v RouteFilterRuleResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type RouteFilterRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (RouteFilterRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteFilterRuleResponse)(nil)).Elem()
}

func (o RouteFilterRuleResponseArrayOutput) ToRouteFilterRuleResponseArrayOutput() RouteFilterRuleResponseArrayOutput {
	return o
}

func (o RouteFilterRuleResponseArrayOutput) ToRouteFilterRuleResponseArrayOutputWithContext(ctx context.Context) RouteFilterRuleResponseArrayOutput {
	return o
}

func (o RouteFilterRuleResponseArrayOutput) Index(i pulumi.IntInput) RouteFilterRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RouteFilterRuleResponse {
		return vs[0].([]RouteFilterRuleResponse)[vs[1].(int)]
	}).(RouteFilterRuleResponseOutput)
}

type RouteResponse struct {
	AddressPrefix     *string `pulumi:"addressPrefix"`
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	NextHopIpAddress  *string `pulumi:"nextHopIpAddress"`
	NextHopType       string  `pulumi:"nextHopType"`
	ProvisioningState *string `pulumi:"provisioningState"`
}

type RouteResponseOutput struct{ *pulumi.OutputState }

func (RouteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteResponse)(nil)).Elem()
}

func (o RouteResponseOutput) ToRouteResponseOutput() RouteResponseOutput {
	return o
}

func (o RouteResponseOutput) ToRouteResponseOutputWithContext(ctx context.Context) RouteResponseOutput {
	return o
}

func (o RouteResponseOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteResponse) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o RouteResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o RouteResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RouteResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RouteResponseOutput) NextHopIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteResponse) *string { return v.NextHopIpAddress }).(pulumi.StringPtrOutput)
}

func (o RouteResponseOutput) NextHopType() pulumi.StringOutput {
	return o.ApplyT(func(v RouteResponse) string { return v.NextHopType }).(pulumi.StringOutput)
}

func (o RouteResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type RouteResponseArrayOutput struct{ *pulumi.OutputState }

func (RouteResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteResponse)(nil)).Elem()
}

func (o RouteResponseArrayOutput) ToRouteResponseArrayOutput() RouteResponseArrayOutput {
	return o
}

func (o RouteResponseArrayOutput) ToRouteResponseArrayOutputWithContext(ctx context.Context) RouteResponseArrayOutput {
	return o
}

func (o RouteResponseArrayOutput) Index(i pulumi.IntInput) RouteResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RouteResponse {
		return vs[0].([]RouteResponse)[vs[1].(int)]
	}).(RouteResponseOutput)
}

type RouteTableType struct {
	Etag              *string           `pulumi:"etag"`
	Id                *string           `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	Routes            []RouteType       `pulumi:"routes"`
	Tags              map[string]string `pulumi:"tags"`
}





type RouteTableTypeInput interface {
	pulumi.Input

	ToRouteTableTypeOutput() RouteTableTypeOutput
	ToRouteTableTypeOutputWithContext(context.Context) RouteTableTypeOutput
}

type RouteTableTypeArgs struct {
	Etag              pulumi.StringPtrInput `pulumi:"etag"`
	Id                pulumi.StringPtrInput `pulumi:"id"`
	Location          pulumi.StringPtrInput `pulumi:"location"`
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
	Routes            RouteTypeArrayInput   `pulumi:"routes"`
	Tags              pulumi.StringMapInput `pulumi:"tags"`
}

func (RouteTableTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTableType)(nil)).Elem()
}

func (i RouteTableTypeArgs) ToRouteTableTypeOutput() RouteTableTypeOutput {
	return i.ToRouteTableTypeOutputWithContext(context.Background())
}

func (i RouteTableTypeArgs) ToRouteTableTypeOutputWithContext(ctx context.Context) RouteTableTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableTypeOutput)
}

func (i RouteTableTypeArgs) ToRouteTableTypePtrOutput() RouteTableTypePtrOutput {
	return i.ToRouteTableTypePtrOutputWithContext(context.Background())
}

func (i RouteTableTypeArgs) ToRouteTableTypePtrOutputWithContext(ctx context.Context) RouteTableTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableTypeOutput).ToRouteTableTypePtrOutputWithContext(ctx)
}









type RouteTableTypePtrInput interface {
	pulumi.Input

	ToRouteTableTypePtrOutput() RouteTableTypePtrOutput
	ToRouteTableTypePtrOutputWithContext(context.Context) RouteTableTypePtrOutput
}

type routeTableTypePtrType RouteTableTypeArgs

func RouteTableTypePtr(v *RouteTableTypeArgs) RouteTableTypePtrInput {
	return (*routeTableTypePtrType)(v)
}

func (*routeTableTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTableType)(nil)).Elem()
}

func (i *routeTableTypePtrType) ToRouteTableTypePtrOutput() RouteTableTypePtrOutput {
	return i.ToRouteTableTypePtrOutputWithContext(context.Background())
}

func (i *routeTableTypePtrType) ToRouteTableTypePtrOutputWithContext(ctx context.Context) RouteTableTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableTypePtrOutput)
}

type RouteTableTypeOutput struct{ *pulumi.OutputState }

func (RouteTableTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTableType)(nil)).Elem()
}

func (o RouteTableTypeOutput) ToRouteTableTypeOutput() RouteTableTypeOutput {
	return o
}

func (o RouteTableTypeOutput) ToRouteTableTypeOutputWithContext(ctx context.Context) RouteTableTypeOutput {
	return o
}

func (o RouteTableTypeOutput) ToRouteTableTypePtrOutput() RouteTableTypePtrOutput {
	return o.ToRouteTableTypePtrOutputWithContext(context.Background())
}

func (o RouteTableTypeOutput) ToRouteTableTypePtrOutputWithContext(ctx context.Context) RouteTableTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RouteTableType) *RouteTableType {
		return &v
	}).(RouteTableTypePtrOutput)
}

func (o RouteTableTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteTableType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o RouteTableTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteTableType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RouteTableTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteTableType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o RouteTableTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteTableType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o RouteTableTypeOutput) Routes() RouteTypeArrayOutput {
	return o.ApplyT(func(v RouteTableType) []RouteType { return v.Routes }).(RouteTypeArrayOutput)
}

func (o RouteTableTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v RouteTableType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type RouteTableTypePtrOutput struct{ *pulumi.OutputState }

func (RouteTableTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTableType)(nil)).Elem()
}

func (o RouteTableTypePtrOutput) ToRouteTableTypePtrOutput() RouteTableTypePtrOutput {
	return o
}

func (o RouteTableTypePtrOutput) ToRouteTableTypePtrOutputWithContext(ctx context.Context) RouteTableTypePtrOutput {
	return o
}

func (o RouteTableTypePtrOutput) Elem() RouteTableTypeOutput {
	return o.ApplyT(func(v *RouteTableType) RouteTableType {
		if v != nil {
			return *v
		}
		var ret RouteTableType
		return ret
	}).(RouteTableTypeOutput)
}

func (o RouteTableTypePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableType) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o RouteTableTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o RouteTableTypePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableType) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o RouteTableTypePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableType) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o RouteTableTypePtrOutput) Routes() RouteTypeArrayOutput {
	return o.ApplyT(func(v *RouteTableType) []RouteType {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(RouteTypeArrayOutput)
}

func (o RouteTableTypePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RouteTableType) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

type RouteTableResponse struct {
	Etag              *string           `pulumi:"etag"`
	Id                *string           `pulumi:"id"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	Routes            []RouteResponse   `pulumi:"routes"`
	Subnets           []SubnetResponse  `pulumi:"subnets"`
	Tags              map[string]string `pulumi:"tags"`
	Type              string            `pulumi:"type"`
}

type RouteTableResponseOutput struct{ *pulumi.OutputState }

func (RouteTableResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTableResponse)(nil)).Elem()
}

func (o RouteTableResponseOutput) ToRouteTableResponseOutput() RouteTableResponseOutput {
	return o
}

func (o RouteTableResponseOutput) ToRouteTableResponseOutputWithContext(ctx context.Context) RouteTableResponseOutput {
	return o
}

func (o RouteTableResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteTableResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o RouteTableResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteTableResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RouteTableResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteTableResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o RouteTableResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RouteTableResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RouteTableResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteTableResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o RouteTableResponseOutput) Routes() RouteResponseArrayOutput {
	return o.ApplyT(func(v RouteTableResponse) []RouteResponse { return v.Routes }).(RouteResponseArrayOutput)
}

func (o RouteTableResponseOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v RouteTableResponse) []SubnetResponse { return v.Subnets }).(SubnetResponseArrayOutput)
}

func (o RouteTableResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v RouteTableResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o RouteTableResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RouteTableResponse) string { return v.Type }).(pulumi.StringOutput)
}

type RouteTableResponsePtrOutput struct{ *pulumi.OutputState }

func (RouteTableResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTableResponse)(nil)).Elem()
}

func (o RouteTableResponsePtrOutput) ToRouteTableResponsePtrOutput() RouteTableResponsePtrOutput {
	return o
}

func (o RouteTableResponsePtrOutput) ToRouteTableResponsePtrOutputWithContext(ctx context.Context) RouteTableResponsePtrOutput {
	return o
}

func (o RouteTableResponsePtrOutput) Elem() RouteTableResponseOutput {
	return o.ApplyT(func(v *RouteTableResponse) RouteTableResponse {
		if v != nil {
			return *v
		}
		var ret RouteTableResponse
		return ret
	}).(RouteTableResponseOutput)
}

func (o RouteTableResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableResponse) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o RouteTableResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o RouteTableResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o RouteTableResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o RouteTableResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o RouteTableResponsePtrOutput) Routes() RouteResponseArrayOutput {
	return o.ApplyT(func(v *RouteTableResponse) []RouteResponse {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(RouteResponseArrayOutput)
}

func (o RouteTableResponsePtrOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v *RouteTableResponse) []SubnetResponse {
		if v == nil {
			return nil
		}
		return v.Subnets
	}).(SubnetResponseArrayOutput)
}

func (o RouteTableResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RouteTableResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o RouteTableResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteTableResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type SecurityRuleType struct {
	Access                   string  `pulumi:"access"`
	Description              *string `pulumi:"description"`
	DestinationAddressPrefix string  `pulumi:"destinationAddressPrefix"`
	DestinationPortRange     *string `pulumi:"destinationPortRange"`
	Direction                string  `pulumi:"direction"`
	Etag                     *string `pulumi:"etag"`
	Id                       *string `pulumi:"id"`
	Name                     *string `pulumi:"name"`
	Priority                 *int    `pulumi:"priority"`
	Protocol                 string  `pulumi:"protocol"`
	ProvisioningState        *string `pulumi:"provisioningState"`
	SourceAddressPrefix      string  `pulumi:"sourceAddressPrefix"`
	SourcePortRange          *string `pulumi:"sourcePortRange"`
}





type SecurityRuleTypeInput interface {
	pulumi.Input

	ToSecurityRuleTypeOutput() SecurityRuleTypeOutput
	ToSecurityRuleTypeOutputWithContext(context.Context) SecurityRuleTypeOutput
}

type SecurityRuleTypeArgs struct {
	Access                   pulumi.StringInput    `pulumi:"access"`
	Description              pulumi.StringPtrInput `pulumi:"description"`
	DestinationAddressPrefix pulumi.StringInput    `pulumi:"destinationAddressPrefix"`
	DestinationPortRange     pulumi.StringPtrInput `pulumi:"destinationPortRange"`
	Direction                pulumi.StringInput    `pulumi:"direction"`
	Etag                     pulumi.StringPtrInput `pulumi:"etag"`
	Id                       pulumi.StringPtrInput `pulumi:"id"`
	Name                     pulumi.StringPtrInput `pulumi:"name"`
	Priority                 pulumi.IntPtrInput    `pulumi:"priority"`
	Protocol                 pulumi.StringInput    `pulumi:"protocol"`
	ProvisioningState        pulumi.StringPtrInput `pulumi:"provisioningState"`
	SourceAddressPrefix      pulumi.StringInput    `pulumi:"sourceAddressPrefix"`
	SourcePortRange          pulumi.StringPtrInput `pulumi:"sourcePortRange"`
}

func (SecurityRuleTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityRuleType)(nil)).Elem()
}

func (i SecurityRuleTypeArgs) ToSecurityRuleTypeOutput() SecurityRuleTypeOutput {
	return i.ToSecurityRuleTypeOutputWithContext(context.Background())
}

func (i SecurityRuleTypeArgs) ToSecurityRuleTypeOutputWithContext(ctx context.Context) SecurityRuleTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityRuleTypeOutput)
}





type SecurityRuleTypeArrayInput interface {
	pulumi.Input

	ToSecurityRuleTypeArrayOutput() SecurityRuleTypeArrayOutput
	ToSecurityRuleTypeArrayOutputWithContext(context.Context) SecurityRuleTypeArrayOutput
}

type SecurityRuleTypeArray []SecurityRuleTypeInput

func (SecurityRuleTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityRuleType)(nil)).Elem()
}

func (i SecurityRuleTypeArray) ToSecurityRuleTypeArrayOutput() SecurityRuleTypeArrayOutput {
	return i.ToSecurityRuleTypeArrayOutputWithContext(context.Background())
}

func (i SecurityRuleTypeArray) ToSecurityRuleTypeArrayOutputWithContext(ctx context.Context) SecurityRuleTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityRuleTypeArrayOutput)
}

type SecurityRuleTypeOutput struct{ *pulumi.OutputState }

func (SecurityRuleTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityRuleType)(nil)).Elem()
}

func (o SecurityRuleTypeOutput) ToSecurityRuleTypeOutput() SecurityRuleTypeOutput {
	return o
}

func (o SecurityRuleTypeOutput) ToSecurityRuleTypeOutputWithContext(ctx context.Context) SecurityRuleTypeOutput {
	return o
}

func (o SecurityRuleTypeOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleType) string { return v.Access }).(pulumi.StringOutput)
}

func (o SecurityRuleTypeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) DestinationAddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleType) string { return v.DestinationAddressPrefix }).(pulumi.StringOutput)
}

func (o SecurityRuleTypeOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleType) string { return v.Direction }).(pulumi.StringOutput)
}

func (o SecurityRuleTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o SecurityRuleTypeOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleType) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o SecurityRuleTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleTypeOutput) SourceAddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleType) string { return v.SourceAddressPrefix }).(pulumi.StringOutput)
}

func (o SecurityRuleTypeOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleType) *string { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

type SecurityRuleTypeArrayOutput struct{ *pulumi.OutputState }

func (SecurityRuleTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityRuleType)(nil)).Elem()
}

func (o SecurityRuleTypeArrayOutput) ToSecurityRuleTypeArrayOutput() SecurityRuleTypeArrayOutput {
	return o
}

func (o SecurityRuleTypeArrayOutput) ToSecurityRuleTypeArrayOutputWithContext(ctx context.Context) SecurityRuleTypeArrayOutput {
	return o
}

func (o SecurityRuleTypeArrayOutput) Index(i pulumi.IntInput) SecurityRuleTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityRuleType {
		return vs[0].([]SecurityRuleType)[vs[1].(int)]
	}).(SecurityRuleTypeOutput)
}

type SecurityRuleResponse struct {
	Access                   string  `pulumi:"access"`
	Description              *string `pulumi:"description"`
	DestinationAddressPrefix string  `pulumi:"destinationAddressPrefix"`
	DestinationPortRange     *string `pulumi:"destinationPortRange"`
	Direction                string  `pulumi:"direction"`
	Etag                     *string `pulumi:"etag"`
	Id                       *string `pulumi:"id"`
	Name                     *string `pulumi:"name"`
	Priority                 *int    `pulumi:"priority"`
	Protocol                 string  `pulumi:"protocol"`
	ProvisioningState        *string `pulumi:"provisioningState"`
	SourceAddressPrefix      string  `pulumi:"sourceAddressPrefix"`
	SourcePortRange          *string `pulumi:"sourcePortRange"`
}

type SecurityRuleResponseOutput struct{ *pulumi.OutputState }

func (SecurityRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityRuleResponse)(nil)).Elem()
}

func (o SecurityRuleResponseOutput) ToSecurityRuleResponseOutput() SecurityRuleResponseOutput {
	return o
}

func (o SecurityRuleResponseOutput) ToSecurityRuleResponseOutputWithContext(ctx context.Context) SecurityRuleResponseOutput {
	return o
}

func (o SecurityRuleResponseOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleResponse) string { return v.Access }).(pulumi.StringOutput)
}

func (o SecurityRuleResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) DestinationAddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleResponse) string { return v.DestinationAddressPrefix }).(pulumi.StringOutput)
}

func (o SecurityRuleResponseOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleResponse) string { return v.Direction }).(pulumi.StringOutput)
}

func (o SecurityRuleResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o SecurityRuleResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o SecurityRuleResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o SecurityRuleResponseOutput) SourceAddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityRuleResponse) string { return v.SourceAddressPrefix }).(pulumi.StringOutput)
}

func (o SecurityRuleResponseOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityRuleResponse) *string { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

type SecurityRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (SecurityRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityRuleResponse)(nil)).Elem()
}

func (o SecurityRuleResponseArrayOutput) ToSecurityRuleResponseArrayOutput() SecurityRuleResponseArrayOutput {
	return o
}

func (o SecurityRuleResponseArrayOutput) ToSecurityRuleResponseArrayOutputWithContext(ctx context.Context) SecurityRuleResponseArrayOutput {
	return o
}

func (o SecurityRuleResponseArrayOutput) Index(i pulumi.IntInput) SecurityRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityRuleResponse {
		return vs[0].([]SecurityRuleResponse)[vs[1].(int)]
	}).(SecurityRuleResponseOutput)
}

type SubResource struct {
	Id *string `pulumi:"id"`
}





type SubResourceInput interface {
	pulumi.Input

	ToSubResourceOutput() SubResourceOutput
	ToSubResourceOutputWithContext(context.Context) SubResourceOutput
}

type SubResourceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (i SubResourceArgs) ToSubResourceOutput() SubResourceOutput {
	return i.ToSubResourceOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput)
}

func (i SubResourceArgs) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput).ToSubResourcePtrOutputWithContext(ctx)
}









type SubResourcePtrInput interface {
	pulumi.Input

	ToSubResourcePtrOutput() SubResourcePtrOutput
	ToSubResourcePtrOutputWithContext(context.Context) SubResourcePtrOutput
}

type subResourcePtrType SubResourceArgs

func SubResourcePtr(v *SubResourceArgs) SubResourcePtrInput {
	return (*subResourcePtrType)(v)
}

func (*subResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (i *subResourcePtrType) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i *subResourcePtrType) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourcePtrOutput)
}





type SubResourceArrayInput interface {
	pulumi.Input

	ToSubResourceArrayOutput() SubResourceArrayOutput
	ToSubResourceArrayOutputWithContext(context.Context) SubResourceArrayOutput
}

type SubResourceArray []SubResourceInput

func (SubResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (i SubResourceArray) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return i.ToSubResourceArrayOutputWithContext(context.Background())
}

func (i SubResourceArray) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceArrayOutput)
}

type SubResourceOutput struct{ *pulumi.OutputState }

func (SubResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (o SubResourceOutput) ToSubResourceOutput() SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o.ToSubResourcePtrOutputWithContext(context.Background())
}

func (o SubResourceOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubResource) *SubResource {
		return &v
	}).(SubResourcePtrOutput)
}

func (o SubResourceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResource) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourcePtrOutput struct{ *pulumi.OutputState }

func (SubResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) Elem() SubResourceOutput {
	return o.ApplyT(func(v *SubResource) SubResource {
		if v != nil {
			return *v
		}
		var ret SubResource
		return ret
	}).(SubResourceOutput)
}

func (o SubResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResource) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceArrayOutput struct{ *pulumi.OutputState }

func (SubResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResource)(nil)).Elem()
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutput() SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) ToSubResourceArrayOutputWithContext(ctx context.Context) SubResourceArrayOutput {
	return o
}

func (o SubResourceArrayOutput) Index(i pulumi.IntInput) SubResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResource {
		return vs[0].([]SubResource)[vs[1].(int)]
	}).(SubResourceOutput)
}

type SubResourceResponse struct {
	Id *string `pulumi:"id"`
}

type SubResourceResponseOutput struct{ *pulumi.OutputState }

func (SubResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourceResponsePtrOutput struct{ *pulumi.OutputState }

func (SubResourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) Elem() SubResourceResponseOutput {
	return o.ApplyT(func(v *SubResourceResponse) SubResourceResponse {
		if v != nil {
			return *v
		}
		var ret SubResourceResponse
		return ret
	}).(SubResourceResponseOutput)
}

func (o SubResourceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SubResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutput() SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutputWithContext(ctx context.Context) SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) Index(i pulumi.IntInput) SubResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResourceResponse {
		return vs[0].([]SubResourceResponse)[vs[1].(int)]
	}).(SubResourceResponseOutput)
}

type SubnetType struct {
	AddressPrefix           *string                   `pulumi:"addressPrefix"`
	Etag                    *string                   `pulumi:"etag"`
	Id                      *string                   `pulumi:"id"`
	Name                    *string                   `pulumi:"name"`
	NetworkSecurityGroup    *NetworkSecurityGroupType `pulumi:"networkSecurityGroup"`
	ProvisioningState       *string                   `pulumi:"provisioningState"`
	ResourceNavigationLinks []ResourceNavigationLink  `pulumi:"resourceNavigationLinks"`
	RouteTable              *RouteTableType           `pulumi:"routeTable"`
}





type SubnetTypeInput interface {
	pulumi.Input

	ToSubnetTypeOutput() SubnetTypeOutput
	ToSubnetTypeOutputWithContext(context.Context) SubnetTypeOutput
}

type SubnetTypeArgs struct {
	AddressPrefix           pulumi.StringPtrInput            `pulumi:"addressPrefix"`
	Etag                    pulumi.StringPtrInput            `pulumi:"etag"`
	Id                      pulumi.StringPtrInput            `pulumi:"id"`
	Name                    pulumi.StringPtrInput            `pulumi:"name"`
	NetworkSecurityGroup    NetworkSecurityGroupTypePtrInput `pulumi:"networkSecurityGroup"`
	ProvisioningState       pulumi.StringPtrInput            `pulumi:"provisioningState"`
	ResourceNavigationLinks ResourceNavigationLinkArrayInput `pulumi:"resourceNavigationLinks"`
	RouteTable              RouteTableTypePtrInput           `pulumi:"routeTable"`
}

func (SubnetTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetType)(nil)).Elem()
}

func (i SubnetTypeArgs) ToSubnetTypeOutput() SubnetTypeOutput {
	return i.ToSubnetTypeOutputWithContext(context.Background())
}

func (i SubnetTypeArgs) ToSubnetTypeOutputWithContext(ctx context.Context) SubnetTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetTypeOutput)
}

func (i SubnetTypeArgs) ToSubnetTypePtrOutput() SubnetTypePtrOutput {
	return i.ToSubnetTypePtrOutputWithContext(context.Background())
}

func (i SubnetTypeArgs) ToSubnetTypePtrOutputWithContext(ctx context.Context) SubnetTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetTypeOutput).ToSubnetTypePtrOutputWithContext(ctx)
}









type SubnetTypePtrInput interface {
	pulumi.Input

	ToSubnetTypePtrOutput() SubnetTypePtrOutput
	ToSubnetTypePtrOutputWithContext(context.Context) SubnetTypePtrOutput
}

type subnetTypePtrType SubnetTypeArgs

func SubnetTypePtr(v *SubnetTypeArgs) SubnetTypePtrInput {
	return (*subnetTypePtrType)(v)
}

func (*subnetTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetType)(nil)).Elem()
}

func (i *subnetTypePtrType) ToSubnetTypePtrOutput() SubnetTypePtrOutput {
	return i.ToSubnetTypePtrOutputWithContext(context.Background())
}

func (i *subnetTypePtrType) ToSubnetTypePtrOutputWithContext(ctx context.Context) SubnetTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetTypePtrOutput)
}





type SubnetTypeArrayInput interface {
	pulumi.Input

	ToSubnetTypeArrayOutput() SubnetTypeArrayOutput
	ToSubnetTypeArrayOutputWithContext(context.Context) SubnetTypeArrayOutput
}

type SubnetTypeArray []SubnetTypeInput

func (SubnetTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetType)(nil)).Elem()
}

func (i SubnetTypeArray) ToSubnetTypeArrayOutput() SubnetTypeArrayOutput {
	return i.ToSubnetTypeArrayOutputWithContext(context.Background())
}

func (i SubnetTypeArray) ToSubnetTypeArrayOutputWithContext(ctx context.Context) SubnetTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetTypeArrayOutput)
}

type SubnetTypeOutput struct{ *pulumi.OutputState }

func (SubnetTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetType)(nil)).Elem()
}

func (o SubnetTypeOutput) ToSubnetTypeOutput() SubnetTypeOutput {
	return o
}

func (o SubnetTypeOutput) ToSubnetTypeOutputWithContext(ctx context.Context) SubnetTypeOutput {
	return o
}

func (o SubnetTypeOutput) ToSubnetTypePtrOutput() SubnetTypePtrOutput {
	return o.ToSubnetTypePtrOutputWithContext(context.Background())
}

func (o SubnetTypeOutput) ToSubnetTypePtrOutputWithContext(ctx context.Context) SubnetTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubnetType) *SubnetType {
		return &v
	}).(SubnetTypePtrOutput)
}

func (o SubnetTypeOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) NetworkSecurityGroup() NetworkSecurityGroupTypePtrOutput {
	return o.ApplyT(func(v SubnetType) *NetworkSecurityGroupType { return v.NetworkSecurityGroup }).(NetworkSecurityGroupTypePtrOutput)
}

func (o SubnetTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o SubnetTypeOutput) ResourceNavigationLinks() ResourceNavigationLinkArrayOutput {
	return o.ApplyT(func(v SubnetType) []ResourceNavigationLink { return v.ResourceNavigationLinks }).(ResourceNavigationLinkArrayOutput)
}

func (o SubnetTypeOutput) RouteTable() RouteTableTypePtrOutput {
	return o.ApplyT(func(v SubnetType) *RouteTableType { return v.RouteTable }).(RouteTableTypePtrOutput)
}

type SubnetTypePtrOutput struct{ *pulumi.OutputState }

func (SubnetTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetType)(nil)).Elem()
}

func (o SubnetTypePtrOutput) ToSubnetTypePtrOutput() SubnetTypePtrOutput {
	return o
}

func (o SubnetTypePtrOutput) ToSubnetTypePtrOutputWithContext(ctx context.Context) SubnetTypePtrOutput {
	return o
}

func (o SubnetTypePtrOutput) Elem() SubnetTypeOutput {
	return o.ApplyT(func(v *SubnetType) SubnetType {
		if v != nil {
			return *v
		}
		var ret SubnetType
		return ret
	}).(SubnetTypeOutput)
}

func (o SubnetTypePtrOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetType) *string {
		if v == nil {
			return nil
		}
		return v.AddressPrefix
	}).(pulumi.StringPtrOutput)
}

func (o SubnetTypePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetType) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o SubnetTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SubnetTypePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetType) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SubnetTypePtrOutput) NetworkSecurityGroup() NetworkSecurityGroupTypePtrOutput {
	return o.ApplyT(func(v *SubnetType) *NetworkSecurityGroupType {
		if v == nil {
			return nil
		}
		return v.NetworkSecurityGroup
	}).(NetworkSecurityGroupTypePtrOutput)
}

func (o SubnetTypePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetType) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o SubnetTypePtrOutput) ResourceNavigationLinks() ResourceNavigationLinkArrayOutput {
	return o.ApplyT(func(v *SubnetType) []ResourceNavigationLink {
		if v == nil {
			return nil
		}
		return v.ResourceNavigationLinks
	}).(ResourceNavigationLinkArrayOutput)
}

func (o SubnetTypePtrOutput) RouteTable() RouteTableTypePtrOutput {
	return o.ApplyT(func(v *SubnetType) *RouteTableType {
		if v == nil {
			return nil
		}
		return v.RouteTable
	}).(RouteTableTypePtrOutput)
}

type SubnetTypeArrayOutput struct{ *pulumi.OutputState }

func (SubnetTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetType)(nil)).Elem()
}

func (o SubnetTypeArrayOutput) ToSubnetTypeArrayOutput() SubnetTypeArrayOutput {
	return o
}

func (o SubnetTypeArrayOutput) ToSubnetTypeArrayOutputWithContext(ctx context.Context) SubnetTypeArrayOutput {
	return o
}

func (o SubnetTypeArrayOutput) Index(i pulumi.IntInput) SubnetTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetType {
		return vs[0].([]SubnetType)[vs[1].(int)]
	}).(SubnetTypeOutput)
}

type SubnetResponse struct {
	AddressPrefix           *string                          `pulumi:"addressPrefix"`
	Etag                    *string                          `pulumi:"etag"`
	Id                      *string                          `pulumi:"id"`
	IpConfigurations        []IPConfigurationResponse        `pulumi:"ipConfigurations"`
	Name                    *string                          `pulumi:"name"`
	NetworkSecurityGroup    *NetworkSecurityGroupResponse    `pulumi:"networkSecurityGroup"`
	ProvisioningState       *string                          `pulumi:"provisioningState"`
	ResourceNavigationLinks []ResourceNavigationLinkResponse `pulumi:"resourceNavigationLinks"`
	RouteTable              *RouteTableResponse              `pulumi:"routeTable"`
}

type SubnetResponseOutput struct{ *pulumi.OutputState }

func (SubnetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResponse)(nil)).Elem()
}

func (o SubnetResponseOutput) ToSubnetResponseOutput() SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) ToSubnetResponseOutputWithContext(ctx context.Context) SubnetResponseOutput {
	return o
}

func (o SubnetResponseOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) IpConfigurations() IPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v SubnetResponse) []IPConfigurationResponse { return v.IpConfigurations }).(IPConfigurationResponseArrayOutput)
}

func (o SubnetResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v SubnetResponse) *NetworkSecurityGroupResponse { return v.NetworkSecurityGroup }).(NetworkSecurityGroupResponsePtrOutput)
}

func (o SubnetResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o SubnetResponseOutput) ResourceNavigationLinks() ResourceNavigationLinkResponseArrayOutput {
	return o.ApplyT(func(v SubnetResponse) []ResourceNavigationLinkResponse { return v.ResourceNavigationLinks }).(ResourceNavigationLinkResponseArrayOutput)
}

func (o SubnetResponseOutput) RouteTable() RouteTableResponsePtrOutput {
	return o.ApplyT(func(v SubnetResponse) *RouteTableResponse { return v.RouteTable }).(RouteTableResponsePtrOutput)
}

type SubnetResponsePtrOutput struct{ *pulumi.OutputState }

func (SubnetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetResponse)(nil)).Elem()
}

func (o SubnetResponsePtrOutput) ToSubnetResponsePtrOutput() SubnetResponsePtrOutput {
	return o
}

func (o SubnetResponsePtrOutput) ToSubnetResponsePtrOutputWithContext(ctx context.Context) SubnetResponsePtrOutput {
	return o
}

func (o SubnetResponsePtrOutput) Elem() SubnetResponseOutput {
	return o.ApplyT(func(v *SubnetResponse) SubnetResponse {
		if v != nil {
			return *v
		}
		var ret SubnetResponse
		return ret
	}).(SubnetResponseOutput)
}

func (o SubnetResponsePtrOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return v.AddressPrefix
	}).(pulumi.StringPtrOutput)
}

func (o SubnetResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o SubnetResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SubnetResponsePtrOutput) IpConfigurations() IPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *SubnetResponse) []IPConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.IpConfigurations
	}).(IPConfigurationResponseArrayOutput)
}

func (o SubnetResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SubnetResponsePtrOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *NetworkSecurityGroupResponse {
		if v == nil {
			return nil
		}
		return v.NetworkSecurityGroup
	}).(NetworkSecurityGroupResponsePtrOutput)
}

func (o SubnetResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o SubnetResponsePtrOutput) ResourceNavigationLinks() ResourceNavigationLinkResponseArrayOutput {
	return o.ApplyT(func(v *SubnetResponse) []ResourceNavigationLinkResponse {
		if v == nil {
			return nil
		}
		return v.ResourceNavigationLinks
	}).(ResourceNavigationLinkResponseArrayOutput)
}

func (o SubnetResponsePtrOutput) RouteTable() RouteTableResponsePtrOutput {
	return o.ApplyT(func(v *SubnetResponse) *RouteTableResponse {
		if v == nil {
			return nil
		}
		return v.RouteTable
	}).(RouteTableResponsePtrOutput)
}

type SubnetResponseArrayOutput struct{ *pulumi.OutputState }

func (SubnetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetResponse)(nil)).Elem()
}

func (o SubnetResponseArrayOutput) ToSubnetResponseArrayOutput() SubnetResponseArrayOutput {
	return o
}

func (o SubnetResponseArrayOutput) ToSubnetResponseArrayOutputWithContext(ctx context.Context) SubnetResponseArrayOutput {
	return o
}

func (o SubnetResponseArrayOutput) Index(i pulumi.IntInput) SubnetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetResponse {
		return vs[0].([]SubnetResponse)[vs[1].(int)]
	}).(SubnetResponseOutput)
}

type TunnelConnectionHealthResponse struct {
	ConnectionStatus                 string  `pulumi:"connectionStatus"`
	EgressBytesTransferred           float64 `pulumi:"egressBytesTransferred"`
	IngressBytesTransferred          float64 `pulumi:"ingressBytesTransferred"`
	LastConnectionEstablishedUtcTime string  `pulumi:"lastConnectionEstablishedUtcTime"`
	Tunnel                           string  `pulumi:"tunnel"`
}

type TunnelConnectionHealthResponseOutput struct{ *pulumi.OutputState }

func (TunnelConnectionHealthResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TunnelConnectionHealthResponse)(nil)).Elem()
}

func (o TunnelConnectionHealthResponseOutput) ToTunnelConnectionHealthResponseOutput() TunnelConnectionHealthResponseOutput {
	return o
}

func (o TunnelConnectionHealthResponseOutput) ToTunnelConnectionHealthResponseOutputWithContext(ctx context.Context) TunnelConnectionHealthResponseOutput {
	return o
}

func (o TunnelConnectionHealthResponseOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v TunnelConnectionHealthResponse) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o TunnelConnectionHealthResponseOutput) EgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v TunnelConnectionHealthResponse) float64 { return v.EgressBytesTransferred }).(pulumi.Float64Output)
}

func (o TunnelConnectionHealthResponseOutput) IngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v TunnelConnectionHealthResponse) float64 { return v.IngressBytesTransferred }).(pulumi.Float64Output)
}

func (o TunnelConnectionHealthResponseOutput) LastConnectionEstablishedUtcTime() pulumi.StringOutput {
	return o.ApplyT(func(v TunnelConnectionHealthResponse) string { return v.LastConnectionEstablishedUtcTime }).(pulumi.StringOutput)
}

func (o TunnelConnectionHealthResponseOutput) Tunnel() pulumi.StringOutput {
	return o.ApplyT(func(v TunnelConnectionHealthResponse) string { return v.Tunnel }).(pulumi.StringOutput)
}

type TunnelConnectionHealthResponseArrayOutput struct{ *pulumi.OutputState }

func (TunnelConnectionHealthResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TunnelConnectionHealthResponse)(nil)).Elem()
}

func (o TunnelConnectionHealthResponseArrayOutput) ToTunnelConnectionHealthResponseArrayOutput() TunnelConnectionHealthResponseArrayOutput {
	return o
}

func (o TunnelConnectionHealthResponseArrayOutput) ToTunnelConnectionHealthResponseArrayOutputWithContext(ctx context.Context) TunnelConnectionHealthResponseArrayOutput {
	return o
}

func (o TunnelConnectionHealthResponseArrayOutput) Index(i pulumi.IntInput) TunnelConnectionHealthResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TunnelConnectionHealthResponse {
		return vs[0].([]TunnelConnectionHealthResponse)[vs[1].(int)]
	}).(TunnelConnectionHealthResponseOutput)
}

type VirtualNetworkGatewayType struct {
	ActiveActive           *bool                                  `pulumi:"activeActive"`
	BgpSettings            *BgpSettings                           `pulumi:"bgpSettings"`
	EnableBgp              *bool                                  `pulumi:"enableBgp"`
	Etag                   *string                                `pulumi:"etag"`
	GatewayDefaultSite     *SubResource                           `pulumi:"gatewayDefaultSite"`
	GatewayType            *string                                `pulumi:"gatewayType"`
	Id                     *string                                `pulumi:"id"`
	IpConfigurations       []VirtualNetworkGatewayIPConfiguration `pulumi:"ipConfigurations"`
	Location               *string                                `pulumi:"location"`
	ResourceGuid           *string                                `pulumi:"resourceGuid"`
	Sku                    *VirtualNetworkGatewaySku              `pulumi:"sku"`
	Tags                   map[string]string                      `pulumi:"tags"`
	VpnClientConfiguration *VpnClientConfiguration                `pulumi:"vpnClientConfiguration"`
	VpnType                *string                                `pulumi:"vpnType"`
}





type VirtualNetworkGatewayTypeInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayTypeOutput() VirtualNetworkGatewayTypeOutput
	ToVirtualNetworkGatewayTypeOutputWithContext(context.Context) VirtualNetworkGatewayTypeOutput
}

type VirtualNetworkGatewayTypeArgs struct {
	ActiveActive           pulumi.BoolPtrInput                            `pulumi:"activeActive"`
	BgpSettings            BgpSettingsPtrInput                            `pulumi:"bgpSettings"`
	EnableBgp              pulumi.BoolPtrInput                            `pulumi:"enableBgp"`
	Etag                   pulumi.StringPtrInput                          `pulumi:"etag"`
	GatewayDefaultSite     SubResourcePtrInput                            `pulumi:"gatewayDefaultSite"`
	GatewayType            pulumi.StringPtrInput                          `pulumi:"gatewayType"`
	Id                     pulumi.StringPtrInput                          `pulumi:"id"`
	IpConfigurations       VirtualNetworkGatewayIPConfigurationArrayInput `pulumi:"ipConfigurations"`
	Location               pulumi.StringPtrInput                          `pulumi:"location"`
	ResourceGuid           pulumi.StringPtrInput                          `pulumi:"resourceGuid"`
	Sku                    VirtualNetworkGatewaySkuPtrInput               `pulumi:"sku"`
	Tags                   pulumi.StringMapInput                          `pulumi:"tags"`
	VpnClientConfiguration VpnClientConfigurationPtrInput                 `pulumi:"vpnClientConfiguration"`
	VpnType                pulumi.StringPtrInput                          `pulumi:"vpnType"`
}

func (VirtualNetworkGatewayTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayType)(nil)).Elem()
}

func (i VirtualNetworkGatewayTypeArgs) ToVirtualNetworkGatewayTypeOutput() VirtualNetworkGatewayTypeOutput {
	return i.ToVirtualNetworkGatewayTypeOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewayTypeArgs) ToVirtualNetworkGatewayTypeOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayTypeOutput)
}

func (i VirtualNetworkGatewayTypeArgs) ToVirtualNetworkGatewayTypePtrOutput() VirtualNetworkGatewayTypePtrOutput {
	return i.ToVirtualNetworkGatewayTypePtrOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewayTypeArgs) ToVirtualNetworkGatewayTypePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayTypeOutput).ToVirtualNetworkGatewayTypePtrOutputWithContext(ctx)
}









type VirtualNetworkGatewayTypePtrInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayTypePtrOutput() VirtualNetworkGatewayTypePtrOutput
	ToVirtualNetworkGatewayTypePtrOutputWithContext(context.Context) VirtualNetworkGatewayTypePtrOutput
}

type virtualNetworkGatewayTypePtrType VirtualNetworkGatewayTypeArgs

func VirtualNetworkGatewayTypePtr(v *VirtualNetworkGatewayTypeArgs) VirtualNetworkGatewayTypePtrInput {
	return (*virtualNetworkGatewayTypePtrType)(v)
}

func (*virtualNetworkGatewayTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayType)(nil)).Elem()
}

func (i *virtualNetworkGatewayTypePtrType) ToVirtualNetworkGatewayTypePtrOutput() VirtualNetworkGatewayTypePtrOutput {
	return i.ToVirtualNetworkGatewayTypePtrOutputWithContext(context.Background())
}

func (i *virtualNetworkGatewayTypePtrType) ToVirtualNetworkGatewayTypePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayTypePtrOutput)
}

type VirtualNetworkGatewayTypeOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayType)(nil)).Elem()
}

func (o VirtualNetworkGatewayTypeOutput) ToVirtualNetworkGatewayTypeOutput() VirtualNetworkGatewayTypeOutput {
	return o
}

func (o VirtualNetworkGatewayTypeOutput) ToVirtualNetworkGatewayTypeOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypeOutput {
	return o
}

func (o VirtualNetworkGatewayTypeOutput) ToVirtualNetworkGatewayTypePtrOutput() VirtualNetworkGatewayTypePtrOutput {
	return o.ToVirtualNetworkGatewayTypePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewayTypeOutput) ToVirtualNetworkGatewayTypePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkGatewayType) *VirtualNetworkGatewayType {
		return &v
	}).(VirtualNetworkGatewayTypePtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) ActiveActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *bool { return v.ActiveActive }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) BgpSettings() BgpSettingsPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *BgpSettings { return v.BgpSettings }).(BgpSettingsPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) GatewayDefaultSite() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *SubResource { return v.GatewayDefaultSite }).(SubResourcePtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) GatewayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *string { return v.GatewayType }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) IpConfigurations() VirtualNetworkGatewayIPConfigurationArrayOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) []VirtualNetworkGatewayIPConfiguration { return v.IpConfigurations }).(VirtualNetworkGatewayIPConfigurationArrayOutput)
}

func (o VirtualNetworkGatewayTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) Sku() VirtualNetworkGatewaySkuPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *VirtualNetworkGatewaySku { return v.Sku }).(VirtualNetworkGatewaySkuPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkGatewayTypeOutput) VpnClientConfiguration() VpnClientConfigurationPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *VpnClientConfiguration { return v.VpnClientConfiguration }).(VpnClientConfigurationPtrOutput)
}

func (o VirtualNetworkGatewayTypeOutput) VpnType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayType) *string { return v.VpnType }).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewayTypePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayType)(nil)).Elem()
}

func (o VirtualNetworkGatewayTypePtrOutput) ToVirtualNetworkGatewayTypePtrOutput() VirtualNetworkGatewayTypePtrOutput {
	return o
}

func (o VirtualNetworkGatewayTypePtrOutput) ToVirtualNetworkGatewayTypePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayTypePtrOutput {
	return o
}

func (o VirtualNetworkGatewayTypePtrOutput) Elem() VirtualNetworkGatewayTypeOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) VirtualNetworkGatewayType {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkGatewayType
		return ret
	}).(VirtualNetworkGatewayTypeOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) ActiveActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *bool {
		if v == nil {
			return nil
		}
		return v.ActiveActive
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) BgpSettings() BgpSettingsPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *BgpSettings {
		if v == nil {
			return nil
		}
		return v.BgpSettings
	}).(BgpSettingsPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *bool {
		if v == nil {
			return nil
		}
		return v.EnableBgp
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) GatewayDefaultSite() SubResourcePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *SubResource {
		if v == nil {
			return nil
		}
		return v.GatewayDefaultSite
	}).(SubResourcePtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) GatewayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.GatewayType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) IpConfigurations() VirtualNetworkGatewayIPConfigurationArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) []VirtualNetworkGatewayIPConfiguration {
		if v == nil {
			return nil
		}
		return v.IpConfigurations
	}).(VirtualNetworkGatewayIPConfigurationArrayOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) Sku() VirtualNetworkGatewaySkuPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *VirtualNetworkGatewaySku {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(VirtualNetworkGatewaySkuPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) VpnClientConfiguration() VpnClientConfigurationPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *VpnClientConfiguration {
		if v == nil {
			return nil
		}
		return v.VpnClientConfiguration
	}).(VpnClientConfigurationPtrOutput)
}

func (o VirtualNetworkGatewayTypePtrOutput) VpnType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayType) *string {
		if v == nil {
			return nil
		}
		return v.VpnType
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewayIPConfiguration struct {
	Etag                      *string      `pulumi:"etag"`
	Id                        *string      `pulumi:"id"`
	Name                      *string      `pulumi:"name"`
	PrivateIPAllocationMethod *string      `pulumi:"privateIPAllocationMethod"`
	PublicIPAddress           *SubResource `pulumi:"publicIPAddress"`
	Subnet                    *SubResource `pulumi:"subnet"`
}





type VirtualNetworkGatewayIPConfigurationInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayIPConfigurationOutput() VirtualNetworkGatewayIPConfigurationOutput
	ToVirtualNetworkGatewayIPConfigurationOutputWithContext(context.Context) VirtualNetworkGatewayIPConfigurationOutput
}

type VirtualNetworkGatewayIPConfigurationArgs struct {
	Etag                      pulumi.StringPtrInput `pulumi:"etag"`
	Id                        pulumi.StringPtrInput `pulumi:"id"`
	Name                      pulumi.StringPtrInput `pulumi:"name"`
	PrivateIPAllocationMethod pulumi.StringPtrInput `pulumi:"privateIPAllocationMethod"`
	PublicIPAddress           SubResourcePtrInput   `pulumi:"publicIPAddress"`
	Subnet                    SubResourcePtrInput   `pulumi:"subnet"`
}

func (VirtualNetworkGatewayIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayIPConfiguration)(nil)).Elem()
}

func (i VirtualNetworkGatewayIPConfigurationArgs) ToVirtualNetworkGatewayIPConfigurationOutput() VirtualNetworkGatewayIPConfigurationOutput {
	return i.ToVirtualNetworkGatewayIPConfigurationOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewayIPConfigurationArgs) ToVirtualNetworkGatewayIPConfigurationOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayIPConfigurationOutput)
}





type VirtualNetworkGatewayIPConfigurationArrayInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayIPConfigurationArrayOutput() VirtualNetworkGatewayIPConfigurationArrayOutput
	ToVirtualNetworkGatewayIPConfigurationArrayOutputWithContext(context.Context) VirtualNetworkGatewayIPConfigurationArrayOutput
}

type VirtualNetworkGatewayIPConfigurationArray []VirtualNetworkGatewayIPConfigurationInput

func (VirtualNetworkGatewayIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkGatewayIPConfiguration)(nil)).Elem()
}

func (i VirtualNetworkGatewayIPConfigurationArray) ToVirtualNetworkGatewayIPConfigurationArrayOutput() VirtualNetworkGatewayIPConfigurationArrayOutput {
	return i.ToVirtualNetworkGatewayIPConfigurationArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewayIPConfigurationArray) ToVirtualNetworkGatewayIPConfigurationArrayOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayIPConfigurationArrayOutput)
}

type VirtualNetworkGatewayIPConfigurationOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayIPConfiguration)(nil)).Elem()
}

func (o VirtualNetworkGatewayIPConfigurationOutput) ToVirtualNetworkGatewayIPConfigurationOutput() VirtualNetworkGatewayIPConfigurationOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationOutput) ToVirtualNetworkGatewayIPConfigurationOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfiguration) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfiguration) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfiguration) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfiguration) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationOutput) PublicIPAddress() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfiguration) *SubResource { return v.PublicIPAddress }).(SubResourcePtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationOutput) Subnet() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfiguration) *SubResource { return v.Subnet }).(SubResourcePtrOutput)
}

type VirtualNetworkGatewayIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkGatewayIPConfiguration)(nil)).Elem()
}

func (o VirtualNetworkGatewayIPConfigurationArrayOutput) ToVirtualNetworkGatewayIPConfigurationArrayOutput() VirtualNetworkGatewayIPConfigurationArrayOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationArrayOutput) ToVirtualNetworkGatewayIPConfigurationArrayOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationArrayOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationArrayOutput) Index(i pulumi.IntInput) VirtualNetworkGatewayIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkGatewayIPConfiguration {
		return vs[0].([]VirtualNetworkGatewayIPConfiguration)[vs[1].(int)]
	}).(VirtualNetworkGatewayIPConfigurationOutput)
}

type VirtualNetworkGatewayIPConfigurationResponse struct {
	Etag                      *string              `pulumi:"etag"`
	Id                        *string              `pulumi:"id"`
	Name                      *string              `pulumi:"name"`
	PrivateIPAllocationMethod *string              `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         string               `pulumi:"provisioningState"`
	PublicIPAddress           *SubResourceResponse `pulumi:"publicIPAddress"`
	Subnet                    *SubResourceResponse `pulumi:"subnet"`
}

type VirtualNetworkGatewayIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayIPConfigurationResponse)(nil)).Elem()
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) ToVirtualNetworkGatewayIPConfigurationResponseOutput() VirtualNetworkGatewayIPConfigurationResponseOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) ToVirtualNetworkGatewayIPConfigurationResponseOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationResponseOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) PrivateIPAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) *string { return v.PrivateIPAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) PublicIPAddress() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) *SubResourceResponse { return v.PublicIPAddress }).(SubResourceResponsePtrOutput)
}

func (o VirtualNetworkGatewayIPConfigurationResponseOutput) Subnet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayIPConfigurationResponse) *SubResourceResponse { return v.Subnet }).(SubResourceResponsePtrOutput)
}

type VirtualNetworkGatewayIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkGatewayIPConfigurationResponse)(nil)).Elem()
}

func (o VirtualNetworkGatewayIPConfigurationResponseArrayOutput) ToVirtualNetworkGatewayIPConfigurationResponseArrayOutput() VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationResponseArrayOutput) ToVirtualNetworkGatewayIPConfigurationResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
	return o
}

func (o VirtualNetworkGatewayIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkGatewayIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkGatewayIPConfigurationResponse {
		return vs[0].([]VirtualNetworkGatewayIPConfigurationResponse)[vs[1].(int)]
	}).(VirtualNetworkGatewayIPConfigurationResponseOutput)
}

type VirtualNetworkGatewayResponse struct {
	ActiveActive           *bool                                          `pulumi:"activeActive"`
	BgpSettings            *BgpSettingsResponse                           `pulumi:"bgpSettings"`
	EnableBgp              *bool                                          `pulumi:"enableBgp"`
	Etag                   *string                                        `pulumi:"etag"`
	GatewayDefaultSite     *SubResourceResponse                           `pulumi:"gatewayDefaultSite"`
	GatewayType            *string                                        `pulumi:"gatewayType"`
	Id                     *string                                        `pulumi:"id"`
	IpConfigurations       []VirtualNetworkGatewayIPConfigurationResponse `pulumi:"ipConfigurations"`
	Location               *string                                        `pulumi:"location"`
	Name                   string                                         `pulumi:"name"`
	ProvisioningState      string                                         `pulumi:"provisioningState"`
	ResourceGuid           *string                                        `pulumi:"resourceGuid"`
	Sku                    *VirtualNetworkGatewaySkuResponse              `pulumi:"sku"`
	Tags                   map[string]string                              `pulumi:"tags"`
	Type                   string                                         `pulumi:"type"`
	VpnClientConfiguration *VpnClientConfigurationResponse                `pulumi:"vpnClientConfiguration"`
	VpnType                *string                                        `pulumi:"vpnType"`
}

type VirtualNetworkGatewayResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayResponse)(nil)).Elem()
}

func (o VirtualNetworkGatewayResponseOutput) ToVirtualNetworkGatewayResponseOutput() VirtualNetworkGatewayResponseOutput {
	return o
}

func (o VirtualNetworkGatewayResponseOutput) ToVirtualNetworkGatewayResponseOutputWithContext(ctx context.Context) VirtualNetworkGatewayResponseOutput {
	return o
}

func (o VirtualNetworkGatewayResponseOutput) ActiveActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *bool { return v.ActiveActive }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *BgpSettingsResponse { return v.BgpSettings }).(BgpSettingsResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) GatewayDefaultSite() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *SubResourceResponse { return v.GatewayDefaultSite }).(SubResourceResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) GatewayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *string { return v.GatewayType }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) IpConfigurations() VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) []VirtualNetworkGatewayIPConfigurationResponse {
		return v.IpConfigurations
	}).(VirtualNetworkGatewayIPConfigurationResponseArrayOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayResponseOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *string { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Sku() VirtualNetworkGatewaySkuResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *VirtualNetworkGatewaySkuResponse { return v.Sku }).(VirtualNetworkGatewaySkuResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkGatewayResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayResponseOutput) VpnClientConfiguration() VpnClientConfigurationResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *VpnClientConfigurationResponse { return v.VpnClientConfiguration }).(VpnClientConfigurationResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponseOutput) VpnType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewayResponse) *string { return v.VpnType }).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewayResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayResponse)(nil)).Elem()
}

func (o VirtualNetworkGatewayResponsePtrOutput) ToVirtualNetworkGatewayResponsePtrOutput() VirtualNetworkGatewayResponsePtrOutput {
	return o
}

func (o VirtualNetworkGatewayResponsePtrOutput) ToVirtualNetworkGatewayResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewayResponsePtrOutput {
	return o
}

func (o VirtualNetworkGatewayResponsePtrOutput) Elem() VirtualNetworkGatewayResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) VirtualNetworkGatewayResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkGatewayResponse
		return ret
	}).(VirtualNetworkGatewayResponseOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) ActiveActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ActiveActive
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *BgpSettingsResponse {
		if v == nil {
			return nil
		}
		return v.BgpSettings
	}).(BgpSettingsResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableBgp
	}).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) GatewayDefaultSite() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *SubResourceResponse {
		if v == nil {
			return nil
		}
		return v.GatewayDefaultSite
	}).(SubResourceResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) GatewayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.GatewayType
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) IpConfigurations() VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) []VirtualNetworkGatewayIPConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.IpConfigurations
	}).(VirtualNetworkGatewayIPConfigurationResponseArrayOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGuid
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Sku() VirtualNetworkGatewaySkuResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *VirtualNetworkGatewaySkuResponse {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(VirtualNetworkGatewaySkuResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) VpnClientConfiguration() VpnClientConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *VpnClientConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.VpnClientConfiguration
	}).(VpnClientConfigurationResponsePtrOutput)
}

func (o VirtualNetworkGatewayResponsePtrOutput) VpnType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayResponse) *string {
		if v == nil {
			return nil
		}
		return v.VpnType
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewaySku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type VirtualNetworkGatewaySkuInput interface {
	pulumi.Input

	ToVirtualNetworkGatewaySkuOutput() VirtualNetworkGatewaySkuOutput
	ToVirtualNetworkGatewaySkuOutputWithContext(context.Context) VirtualNetworkGatewaySkuOutput
}

type VirtualNetworkGatewaySkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (VirtualNetworkGatewaySkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewaySku)(nil)).Elem()
}

func (i VirtualNetworkGatewaySkuArgs) ToVirtualNetworkGatewaySkuOutput() VirtualNetworkGatewaySkuOutput {
	return i.ToVirtualNetworkGatewaySkuOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewaySkuArgs) ToVirtualNetworkGatewaySkuOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewaySkuOutput)
}

func (i VirtualNetworkGatewaySkuArgs) ToVirtualNetworkGatewaySkuPtrOutput() VirtualNetworkGatewaySkuPtrOutput {
	return i.ToVirtualNetworkGatewaySkuPtrOutputWithContext(context.Background())
}

func (i VirtualNetworkGatewaySkuArgs) ToVirtualNetworkGatewaySkuPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewaySkuOutput).ToVirtualNetworkGatewaySkuPtrOutputWithContext(ctx)
}









type VirtualNetworkGatewaySkuPtrInput interface {
	pulumi.Input

	ToVirtualNetworkGatewaySkuPtrOutput() VirtualNetworkGatewaySkuPtrOutput
	ToVirtualNetworkGatewaySkuPtrOutputWithContext(context.Context) VirtualNetworkGatewaySkuPtrOutput
}

type virtualNetworkGatewaySkuPtrType VirtualNetworkGatewaySkuArgs

func VirtualNetworkGatewaySkuPtr(v *VirtualNetworkGatewaySkuArgs) VirtualNetworkGatewaySkuPtrInput {
	return (*virtualNetworkGatewaySkuPtrType)(v)
}

func (*virtualNetworkGatewaySkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewaySku)(nil)).Elem()
}

func (i *virtualNetworkGatewaySkuPtrType) ToVirtualNetworkGatewaySkuPtrOutput() VirtualNetworkGatewaySkuPtrOutput {
	return i.ToVirtualNetworkGatewaySkuPtrOutputWithContext(context.Background())
}

func (i *virtualNetworkGatewaySkuPtrType) ToVirtualNetworkGatewaySkuPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewaySkuPtrOutput)
}

type VirtualNetworkGatewaySkuOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewaySkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewaySku)(nil)).Elem()
}

func (o VirtualNetworkGatewaySkuOutput) ToVirtualNetworkGatewaySkuOutput() VirtualNetworkGatewaySkuOutput {
	return o
}

func (o VirtualNetworkGatewaySkuOutput) ToVirtualNetworkGatewaySkuOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuOutput {
	return o
}

func (o VirtualNetworkGatewaySkuOutput) ToVirtualNetworkGatewaySkuPtrOutput() VirtualNetworkGatewaySkuPtrOutput {
	return o.ToVirtualNetworkGatewaySkuPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkGatewaySkuOutput) ToVirtualNetworkGatewaySkuPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkGatewaySku) *VirtualNetworkGatewaySku {
		return &v
	}).(VirtualNetworkGatewaySkuPtrOutput)
}

func (o VirtualNetworkGatewaySkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewaySku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkGatewaySkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewaySku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewaySkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewaySku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewaySkuPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewaySkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewaySku)(nil)).Elem()
}

func (o VirtualNetworkGatewaySkuPtrOutput) ToVirtualNetworkGatewaySkuPtrOutput() VirtualNetworkGatewaySkuPtrOutput {
	return o
}

func (o VirtualNetworkGatewaySkuPtrOutput) ToVirtualNetworkGatewaySkuPtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuPtrOutput {
	return o
}

func (o VirtualNetworkGatewaySkuPtrOutput) Elem() VirtualNetworkGatewaySkuOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySku) VirtualNetworkGatewaySku {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkGatewaySku
		return ret
	}).(VirtualNetworkGatewaySkuOutput)
}

func (o VirtualNetworkGatewaySkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkGatewaySkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewaySkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewaySkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}

type VirtualNetworkGatewaySkuResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewaySkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewaySkuResponse)(nil)).Elem()
}

func (o VirtualNetworkGatewaySkuResponseOutput) ToVirtualNetworkGatewaySkuResponseOutput() VirtualNetworkGatewaySkuResponseOutput {
	return o
}

func (o VirtualNetworkGatewaySkuResponseOutput) ToVirtualNetworkGatewaySkuResponseOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuResponseOutput {
	return o
}

func (o VirtualNetworkGatewaySkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewaySkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkGatewaySkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewaySkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewaySkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkGatewaySkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type VirtualNetworkGatewaySkuResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewaySkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewaySkuResponse)(nil)).Elem()
}

func (o VirtualNetworkGatewaySkuResponsePtrOutput) ToVirtualNetworkGatewaySkuResponsePtrOutput() VirtualNetworkGatewaySkuResponsePtrOutput {
	return o
}

func (o VirtualNetworkGatewaySkuResponsePtrOutput) ToVirtualNetworkGatewaySkuResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkGatewaySkuResponsePtrOutput {
	return o
}

func (o VirtualNetworkGatewaySkuResponsePtrOutput) Elem() VirtualNetworkGatewaySkuResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySkuResponse) VirtualNetworkGatewaySkuResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkGatewaySkuResponse
		return ret
	}).(VirtualNetworkGatewaySkuResponseOutput)
}

func (o VirtualNetworkGatewaySkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkGatewaySkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewaySkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewaySkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkPeeringType struct {
	AllowForwardedTraffic     *bool        `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit       *bool        `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess *bool        `pulumi:"allowVirtualNetworkAccess"`
	Etag                      *string      `pulumi:"etag"`
	Id                        *string      `pulumi:"id"`
	Name                      *string      `pulumi:"name"`
	PeeringState              *string      `pulumi:"peeringState"`
	ProvisioningState         *string      `pulumi:"provisioningState"`
	RemoteVirtualNetwork      *SubResource `pulumi:"remoteVirtualNetwork"`
	UseRemoteGateways         *bool        `pulumi:"useRemoteGateways"`
}





type VirtualNetworkPeeringTypeInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringTypeOutput() VirtualNetworkPeeringTypeOutput
	ToVirtualNetworkPeeringTypeOutputWithContext(context.Context) VirtualNetworkPeeringTypeOutput
}

type VirtualNetworkPeeringTypeArgs struct {
	AllowForwardedTraffic     pulumi.BoolPtrInput   `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit       pulumi.BoolPtrInput   `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess pulumi.BoolPtrInput   `pulumi:"allowVirtualNetworkAccess"`
	Etag                      pulumi.StringPtrInput `pulumi:"etag"`
	Id                        pulumi.StringPtrInput `pulumi:"id"`
	Name                      pulumi.StringPtrInput `pulumi:"name"`
	PeeringState              pulumi.StringPtrInput `pulumi:"peeringState"`
	ProvisioningState         pulumi.StringPtrInput `pulumi:"provisioningState"`
	RemoteVirtualNetwork      SubResourcePtrInput   `pulumi:"remoteVirtualNetwork"`
	UseRemoteGateways         pulumi.BoolPtrInput   `pulumi:"useRemoteGateways"`
}

func (VirtualNetworkPeeringTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringType)(nil)).Elem()
}

func (i VirtualNetworkPeeringTypeArgs) ToVirtualNetworkPeeringTypeOutput() VirtualNetworkPeeringTypeOutput {
	return i.ToVirtualNetworkPeeringTypeOutputWithContext(context.Background())
}

func (i VirtualNetworkPeeringTypeArgs) ToVirtualNetworkPeeringTypeOutputWithContext(ctx context.Context) VirtualNetworkPeeringTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringTypeOutput)
}





type VirtualNetworkPeeringTypeArrayInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringTypeArrayOutput() VirtualNetworkPeeringTypeArrayOutput
	ToVirtualNetworkPeeringTypeArrayOutputWithContext(context.Context) VirtualNetworkPeeringTypeArrayOutput
}

type VirtualNetworkPeeringTypeArray []VirtualNetworkPeeringTypeInput

func (VirtualNetworkPeeringTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkPeeringType)(nil)).Elem()
}

func (i VirtualNetworkPeeringTypeArray) ToVirtualNetworkPeeringTypeArrayOutput() VirtualNetworkPeeringTypeArrayOutput {
	return i.ToVirtualNetworkPeeringTypeArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkPeeringTypeArray) ToVirtualNetworkPeeringTypeArrayOutputWithContext(ctx context.Context) VirtualNetworkPeeringTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringTypeArrayOutput)
}

type VirtualNetworkPeeringTypeOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringType)(nil)).Elem()
}

func (o VirtualNetworkPeeringTypeOutput) ToVirtualNetworkPeeringTypeOutput() VirtualNetworkPeeringTypeOutput {
	return o
}

func (o VirtualNetworkPeeringTypeOutput) ToVirtualNetworkPeeringTypeOutputWithContext(ctx context.Context) VirtualNetworkPeeringTypeOutput {
	return o
}

func (o VirtualNetworkPeeringTypeOutput) AllowForwardedTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *bool { return v.AllowForwardedTraffic }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) AllowGatewayTransit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *bool { return v.AllowGatewayTransit }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) AllowVirtualNetworkAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *bool { return v.AllowVirtualNetworkAccess }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) PeeringState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *string { return v.PeeringState }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) RemoteVirtualNetwork() SubResourcePtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *SubResource { return v.RemoteVirtualNetwork }).(SubResourcePtrOutput)
}

func (o VirtualNetworkPeeringTypeOutput) UseRemoteGateways() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringType) *bool { return v.UseRemoteGateways }).(pulumi.BoolPtrOutput)
}

type VirtualNetworkPeeringTypeArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkPeeringType)(nil)).Elem()
}

func (o VirtualNetworkPeeringTypeArrayOutput) ToVirtualNetworkPeeringTypeArrayOutput() VirtualNetworkPeeringTypeArrayOutput {
	return o
}

func (o VirtualNetworkPeeringTypeArrayOutput) ToVirtualNetworkPeeringTypeArrayOutputWithContext(ctx context.Context) VirtualNetworkPeeringTypeArrayOutput {
	return o
}

func (o VirtualNetworkPeeringTypeArrayOutput) Index(i pulumi.IntInput) VirtualNetworkPeeringTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkPeeringType {
		return vs[0].([]VirtualNetworkPeeringType)[vs[1].(int)]
	}).(VirtualNetworkPeeringTypeOutput)
}

type VirtualNetworkPeeringResponse struct {
	AllowForwardedTraffic     *bool                `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit       *bool                `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess *bool                `pulumi:"allowVirtualNetworkAccess"`
	Etag                      *string              `pulumi:"etag"`
	Id                        *string              `pulumi:"id"`
	Name                      *string              `pulumi:"name"`
	PeeringState              *string              `pulumi:"peeringState"`
	ProvisioningState         *string              `pulumi:"provisioningState"`
	RemoteVirtualNetwork      *SubResourceResponse `pulumi:"remoteVirtualNetwork"`
	UseRemoteGateways         *bool                `pulumi:"useRemoteGateways"`
}

type VirtualNetworkPeeringResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeeringResponse)(nil)).Elem()
}

func (o VirtualNetworkPeeringResponseOutput) ToVirtualNetworkPeeringResponseOutput() VirtualNetworkPeeringResponseOutput {
	return o
}

func (o VirtualNetworkPeeringResponseOutput) ToVirtualNetworkPeeringResponseOutputWithContext(ctx context.Context) VirtualNetworkPeeringResponseOutput {
	return o
}

func (o VirtualNetworkPeeringResponseOutput) AllowForwardedTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *bool { return v.AllowForwardedTraffic }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) AllowGatewayTransit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *bool { return v.AllowGatewayTransit }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) AllowVirtualNetworkAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *bool { return v.AllowVirtualNetworkAccess }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) PeeringState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *string { return v.PeeringState }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) RemoteVirtualNetwork() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *SubResourceResponse { return v.RemoteVirtualNetwork }).(SubResourceResponsePtrOutput)
}

func (o VirtualNetworkPeeringResponseOutput) UseRemoteGateways() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkPeeringResponse) *bool { return v.UseRemoteGateways }).(pulumi.BoolPtrOutput)
}

type VirtualNetworkPeeringResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkPeeringResponse)(nil)).Elem()
}

func (o VirtualNetworkPeeringResponseArrayOutput) ToVirtualNetworkPeeringResponseArrayOutput() VirtualNetworkPeeringResponseArrayOutput {
	return o
}

func (o VirtualNetworkPeeringResponseArrayOutput) ToVirtualNetworkPeeringResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkPeeringResponseArrayOutput {
	return o
}

func (o VirtualNetworkPeeringResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkPeeringResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkPeeringResponse {
		return vs[0].([]VirtualNetworkPeeringResponse)[vs[1].(int)]
	}).(VirtualNetworkPeeringResponseOutput)
}

type VpnClientConfiguration struct {
	VpnClientAddressPool         *AddressSpace                 `pulumi:"vpnClientAddressPool"`
	VpnClientRevokedCertificates []VpnClientRevokedCertificate `pulumi:"vpnClientRevokedCertificates"`
	VpnClientRootCertificates    []VpnClientRootCertificate    `pulumi:"vpnClientRootCertificates"`
}





type VpnClientConfigurationInput interface {
	pulumi.Input

	ToVpnClientConfigurationOutput() VpnClientConfigurationOutput
	ToVpnClientConfigurationOutputWithContext(context.Context) VpnClientConfigurationOutput
}

type VpnClientConfigurationArgs struct {
	VpnClientAddressPool         AddressSpacePtrInput                  `pulumi:"vpnClientAddressPool"`
	VpnClientRevokedCertificates VpnClientRevokedCertificateArrayInput `pulumi:"vpnClientRevokedCertificates"`
	VpnClientRootCertificates    VpnClientRootCertificateArrayInput    `pulumi:"vpnClientRootCertificates"`
}

func (VpnClientConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientConfiguration)(nil)).Elem()
}

func (i VpnClientConfigurationArgs) ToVpnClientConfigurationOutput() VpnClientConfigurationOutput {
	return i.ToVpnClientConfigurationOutputWithContext(context.Background())
}

func (i VpnClientConfigurationArgs) ToVpnClientConfigurationOutputWithContext(ctx context.Context) VpnClientConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientConfigurationOutput)
}

func (i VpnClientConfigurationArgs) ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput {
	return i.ToVpnClientConfigurationPtrOutputWithContext(context.Background())
}

func (i VpnClientConfigurationArgs) ToVpnClientConfigurationPtrOutputWithContext(ctx context.Context) VpnClientConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientConfigurationOutput).ToVpnClientConfigurationPtrOutputWithContext(ctx)
}









type VpnClientConfigurationPtrInput interface {
	pulumi.Input

	ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput
	ToVpnClientConfigurationPtrOutputWithContext(context.Context) VpnClientConfigurationPtrOutput
}

type vpnClientConfigurationPtrType VpnClientConfigurationArgs

func VpnClientConfigurationPtr(v *VpnClientConfigurationArgs) VpnClientConfigurationPtrInput {
	return (*vpnClientConfigurationPtrType)(v)
}

func (*vpnClientConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnClientConfiguration)(nil)).Elem()
}

func (i *vpnClientConfigurationPtrType) ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput {
	return i.ToVpnClientConfigurationPtrOutputWithContext(context.Background())
}

func (i *vpnClientConfigurationPtrType) ToVpnClientConfigurationPtrOutputWithContext(ctx context.Context) VpnClientConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientConfigurationPtrOutput)
}

type VpnClientConfigurationOutput struct{ *pulumi.OutputState }

func (VpnClientConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientConfiguration)(nil)).Elem()
}

func (o VpnClientConfigurationOutput) ToVpnClientConfigurationOutput() VpnClientConfigurationOutput {
	return o
}

func (o VpnClientConfigurationOutput) ToVpnClientConfigurationOutputWithContext(ctx context.Context) VpnClientConfigurationOutput {
	return o
}

func (o VpnClientConfigurationOutput) ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput {
	return o.ToVpnClientConfigurationPtrOutputWithContext(context.Background())
}

func (o VpnClientConfigurationOutput) ToVpnClientConfigurationPtrOutputWithContext(ctx context.Context) VpnClientConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VpnClientConfiguration) *VpnClientConfiguration {
		return &v
	}).(VpnClientConfigurationPtrOutput)
}

func (o VpnClientConfigurationOutput) VpnClientAddressPool() AddressSpacePtrOutput {
	return o.ApplyT(func(v VpnClientConfiguration) *AddressSpace { return v.VpnClientAddressPool }).(AddressSpacePtrOutput)
}

func (o VpnClientConfigurationOutput) VpnClientRevokedCertificates() VpnClientRevokedCertificateArrayOutput {
	return o.ApplyT(func(v VpnClientConfiguration) []VpnClientRevokedCertificate { return v.VpnClientRevokedCertificates }).(VpnClientRevokedCertificateArrayOutput)
}

func (o VpnClientConfigurationOutput) VpnClientRootCertificates() VpnClientRootCertificateArrayOutput {
	return o.ApplyT(func(v VpnClientConfiguration) []VpnClientRootCertificate { return v.VpnClientRootCertificates }).(VpnClientRootCertificateArrayOutput)
}

type VpnClientConfigurationPtrOutput struct{ *pulumi.OutputState }

func (VpnClientConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnClientConfiguration)(nil)).Elem()
}

func (o VpnClientConfigurationPtrOutput) ToVpnClientConfigurationPtrOutput() VpnClientConfigurationPtrOutput {
	return o
}

func (o VpnClientConfigurationPtrOutput) ToVpnClientConfigurationPtrOutputWithContext(ctx context.Context) VpnClientConfigurationPtrOutput {
	return o
}

func (o VpnClientConfigurationPtrOutput) Elem() VpnClientConfigurationOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) VpnClientConfiguration {
		if v != nil {
			return *v
		}
		var ret VpnClientConfiguration
		return ret
	}).(VpnClientConfigurationOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnClientAddressPool() AddressSpacePtrOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) *AddressSpace {
		if v == nil {
			return nil
		}
		return v.VpnClientAddressPool
	}).(AddressSpacePtrOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnClientRevokedCertificates() VpnClientRevokedCertificateArrayOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) []VpnClientRevokedCertificate {
		if v == nil {
			return nil
		}
		return v.VpnClientRevokedCertificates
	}).(VpnClientRevokedCertificateArrayOutput)
}

func (o VpnClientConfigurationPtrOutput) VpnClientRootCertificates() VpnClientRootCertificateArrayOutput {
	return o.ApplyT(func(v *VpnClientConfiguration) []VpnClientRootCertificate {
		if v == nil {
			return nil
		}
		return v.VpnClientRootCertificates
	}).(VpnClientRootCertificateArrayOutput)
}

type VpnClientConfigurationResponse struct {
	VpnClientAddressPool         *AddressSpaceResponse                 `pulumi:"vpnClientAddressPool"`
	VpnClientRevokedCertificates []VpnClientRevokedCertificateResponse `pulumi:"vpnClientRevokedCertificates"`
	VpnClientRootCertificates    []VpnClientRootCertificateResponse    `pulumi:"vpnClientRootCertificates"`
}

type VpnClientConfigurationResponseOutput struct{ *pulumi.OutputState }

func (VpnClientConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientConfigurationResponse)(nil)).Elem()
}

func (o VpnClientConfigurationResponseOutput) ToVpnClientConfigurationResponseOutput() VpnClientConfigurationResponseOutput {
	return o
}

func (o VpnClientConfigurationResponseOutput) ToVpnClientConfigurationResponseOutputWithContext(ctx context.Context) VpnClientConfigurationResponseOutput {
	return o
}

func (o VpnClientConfigurationResponseOutput) VpnClientAddressPool() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) *AddressSpaceResponse { return v.VpnClientAddressPool }).(AddressSpaceResponsePtrOutput)
}

func (o VpnClientConfigurationResponseOutput) VpnClientRevokedCertificates() VpnClientRevokedCertificateResponseArrayOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) []VpnClientRevokedCertificateResponse {
		return v.VpnClientRevokedCertificates
	}).(VpnClientRevokedCertificateResponseArrayOutput)
}

func (o VpnClientConfigurationResponseOutput) VpnClientRootCertificates() VpnClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v VpnClientConfigurationResponse) []VpnClientRootCertificateResponse {
		return v.VpnClientRootCertificates
	}).(VpnClientRootCertificateResponseArrayOutput)
}

type VpnClientConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (VpnClientConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnClientConfigurationResponse)(nil)).Elem()
}

func (o VpnClientConfigurationResponsePtrOutput) ToVpnClientConfigurationResponsePtrOutput() VpnClientConfigurationResponsePtrOutput {
	return o
}

func (o VpnClientConfigurationResponsePtrOutput) ToVpnClientConfigurationResponsePtrOutputWithContext(ctx context.Context) VpnClientConfigurationResponsePtrOutput {
	return o
}

func (o VpnClientConfigurationResponsePtrOutput) Elem() VpnClientConfigurationResponseOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) VpnClientConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret VpnClientConfigurationResponse
		return ret
	}).(VpnClientConfigurationResponseOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnClientAddressPool() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) *AddressSpaceResponse {
		if v == nil {
			return nil
		}
		return v.VpnClientAddressPool
	}).(AddressSpaceResponsePtrOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnClientRevokedCertificates() VpnClientRevokedCertificateResponseArrayOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) []VpnClientRevokedCertificateResponse {
		if v == nil {
			return nil
		}
		return v.VpnClientRevokedCertificates
	}).(VpnClientRevokedCertificateResponseArrayOutput)
}

func (o VpnClientConfigurationResponsePtrOutput) VpnClientRootCertificates() VpnClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v *VpnClientConfigurationResponse) []VpnClientRootCertificateResponse {
		if v == nil {
			return nil
		}
		return v.VpnClientRootCertificates
	}).(VpnClientRootCertificateResponseArrayOutput)
}

type VpnClientRevokedCertificate struct {
	Etag       *string `pulumi:"etag"`
	Id         *string `pulumi:"id"`
	Name       *string `pulumi:"name"`
	Thumbprint *string `pulumi:"thumbprint"`
}





type VpnClientRevokedCertificateInput interface {
	pulumi.Input

	ToVpnClientRevokedCertificateOutput() VpnClientRevokedCertificateOutput
	ToVpnClientRevokedCertificateOutputWithContext(context.Context) VpnClientRevokedCertificateOutput
}

type VpnClientRevokedCertificateArgs struct {
	Etag       pulumi.StringPtrInput `pulumi:"etag"`
	Id         pulumi.StringPtrInput `pulumi:"id"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (VpnClientRevokedCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRevokedCertificate)(nil)).Elem()
}

func (i VpnClientRevokedCertificateArgs) ToVpnClientRevokedCertificateOutput() VpnClientRevokedCertificateOutput {
	return i.ToVpnClientRevokedCertificateOutputWithContext(context.Background())
}

func (i VpnClientRevokedCertificateArgs) ToVpnClientRevokedCertificateOutputWithContext(ctx context.Context) VpnClientRevokedCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRevokedCertificateOutput)
}





type VpnClientRevokedCertificateArrayInput interface {
	pulumi.Input

	ToVpnClientRevokedCertificateArrayOutput() VpnClientRevokedCertificateArrayOutput
	ToVpnClientRevokedCertificateArrayOutputWithContext(context.Context) VpnClientRevokedCertificateArrayOutput
}

type VpnClientRevokedCertificateArray []VpnClientRevokedCertificateInput

func (VpnClientRevokedCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRevokedCertificate)(nil)).Elem()
}

func (i VpnClientRevokedCertificateArray) ToVpnClientRevokedCertificateArrayOutput() VpnClientRevokedCertificateArrayOutput {
	return i.ToVpnClientRevokedCertificateArrayOutputWithContext(context.Background())
}

func (i VpnClientRevokedCertificateArray) ToVpnClientRevokedCertificateArrayOutputWithContext(ctx context.Context) VpnClientRevokedCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRevokedCertificateArrayOutput)
}

type VpnClientRevokedCertificateOutput struct{ *pulumi.OutputState }

func (VpnClientRevokedCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRevokedCertificate)(nil)).Elem()
}

func (o VpnClientRevokedCertificateOutput) ToVpnClientRevokedCertificateOutput() VpnClientRevokedCertificateOutput {
	return o
}

func (o VpnClientRevokedCertificateOutput) ToVpnClientRevokedCertificateOutputWithContext(ctx context.Context) VpnClientRevokedCertificateOutput {
	return o
}

func (o VpnClientRevokedCertificateOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificate) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificate) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificate) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnClientRevokedCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnClientRevokedCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRevokedCertificate)(nil)).Elem()
}

func (o VpnClientRevokedCertificateArrayOutput) ToVpnClientRevokedCertificateArrayOutput() VpnClientRevokedCertificateArrayOutput {
	return o
}

func (o VpnClientRevokedCertificateArrayOutput) ToVpnClientRevokedCertificateArrayOutputWithContext(ctx context.Context) VpnClientRevokedCertificateArrayOutput {
	return o
}

func (o VpnClientRevokedCertificateArrayOutput) Index(i pulumi.IntInput) VpnClientRevokedCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnClientRevokedCertificate {
		return vs[0].([]VpnClientRevokedCertificate)[vs[1].(int)]
	}).(VpnClientRevokedCertificateOutput)
}

type VpnClientRevokedCertificateResponse struct {
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	ProvisioningState string  `pulumi:"provisioningState"`
	Thumbprint        *string `pulumi:"thumbprint"`
}

type VpnClientRevokedCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnClientRevokedCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRevokedCertificateResponse)(nil)).Elem()
}

func (o VpnClientRevokedCertificateResponseOutput) ToVpnClientRevokedCertificateResponseOutput() VpnClientRevokedCertificateResponseOutput {
	return o
}

func (o VpnClientRevokedCertificateResponseOutput) ToVpnClientRevokedCertificateResponseOutputWithContext(ctx context.Context) VpnClientRevokedCertificateResponseOutput {
	return o
}

func (o VpnClientRevokedCertificateResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnClientRevokedCertificateResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnClientRevokedCertificateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRevokedCertificateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type VpnClientRevokedCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnClientRevokedCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRevokedCertificateResponse)(nil)).Elem()
}

func (o VpnClientRevokedCertificateResponseArrayOutput) ToVpnClientRevokedCertificateResponseArrayOutput() VpnClientRevokedCertificateResponseArrayOutput {
	return o
}

func (o VpnClientRevokedCertificateResponseArrayOutput) ToVpnClientRevokedCertificateResponseArrayOutputWithContext(ctx context.Context) VpnClientRevokedCertificateResponseArrayOutput {
	return o
}

func (o VpnClientRevokedCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnClientRevokedCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnClientRevokedCertificateResponse {
		return vs[0].([]VpnClientRevokedCertificateResponse)[vs[1].(int)]
	}).(VpnClientRevokedCertificateResponseOutput)
}

type VpnClientRootCertificate struct {
	Etag           *string `pulumi:"etag"`
	Id             *string `pulumi:"id"`
	Name           *string `pulumi:"name"`
	PublicCertData string  `pulumi:"publicCertData"`
}





type VpnClientRootCertificateInput interface {
	pulumi.Input

	ToVpnClientRootCertificateOutput() VpnClientRootCertificateOutput
	ToVpnClientRootCertificateOutputWithContext(context.Context) VpnClientRootCertificateOutput
}

type VpnClientRootCertificateArgs struct {
	Etag           pulumi.StringPtrInput `pulumi:"etag"`
	Id             pulumi.StringPtrInput `pulumi:"id"`
	Name           pulumi.StringPtrInput `pulumi:"name"`
	PublicCertData pulumi.StringInput    `pulumi:"publicCertData"`
}

func (VpnClientRootCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRootCertificate)(nil)).Elem()
}

func (i VpnClientRootCertificateArgs) ToVpnClientRootCertificateOutput() VpnClientRootCertificateOutput {
	return i.ToVpnClientRootCertificateOutputWithContext(context.Background())
}

func (i VpnClientRootCertificateArgs) ToVpnClientRootCertificateOutputWithContext(ctx context.Context) VpnClientRootCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRootCertificateOutput)
}





type VpnClientRootCertificateArrayInput interface {
	pulumi.Input

	ToVpnClientRootCertificateArrayOutput() VpnClientRootCertificateArrayOutput
	ToVpnClientRootCertificateArrayOutputWithContext(context.Context) VpnClientRootCertificateArrayOutput
}

type VpnClientRootCertificateArray []VpnClientRootCertificateInput

func (VpnClientRootCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRootCertificate)(nil)).Elem()
}

func (i VpnClientRootCertificateArray) ToVpnClientRootCertificateArrayOutput() VpnClientRootCertificateArrayOutput {
	return i.ToVpnClientRootCertificateArrayOutputWithContext(context.Background())
}

func (i VpnClientRootCertificateArray) ToVpnClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnClientRootCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnClientRootCertificateArrayOutput)
}

type VpnClientRootCertificateOutput struct{ *pulumi.OutputState }

func (VpnClientRootCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRootCertificate)(nil)).Elem()
}

func (o VpnClientRootCertificateOutput) ToVpnClientRootCertificateOutput() VpnClientRootCertificateOutput {
	return o
}

func (o VpnClientRootCertificateOutput) ToVpnClientRootCertificateOutputWithContext(ctx context.Context) VpnClientRootCertificateOutput {
	return o
}

func (o VpnClientRootCertificateOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificate) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificate) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateOutput) PublicCertData() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRootCertificate) string { return v.PublicCertData }).(pulumi.StringOutput)
}

type VpnClientRootCertificateArrayOutput struct{ *pulumi.OutputState }

func (VpnClientRootCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRootCertificate)(nil)).Elem()
}

func (o VpnClientRootCertificateArrayOutput) ToVpnClientRootCertificateArrayOutput() VpnClientRootCertificateArrayOutput {
	return o
}

func (o VpnClientRootCertificateArrayOutput) ToVpnClientRootCertificateArrayOutputWithContext(ctx context.Context) VpnClientRootCertificateArrayOutput {
	return o
}

func (o VpnClientRootCertificateArrayOutput) Index(i pulumi.IntInput) VpnClientRootCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnClientRootCertificate {
		return vs[0].([]VpnClientRootCertificate)[vs[1].(int)]
	}).(VpnClientRootCertificateOutput)
}

type VpnClientRootCertificateResponse struct {
	Etag              *string `pulumi:"etag"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	ProvisioningState string  `pulumi:"provisioningState"`
	PublicCertData    string  `pulumi:"publicCertData"`
}

type VpnClientRootCertificateResponseOutput struct{ *pulumi.OutputState }

func (VpnClientRootCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnClientRootCertificateResponseOutput) ToVpnClientRootCertificateResponseOutput() VpnClientRootCertificateResponseOutput {
	return o
}

func (o VpnClientRootCertificateResponseOutput) ToVpnClientRootCertificateResponseOutputWithContext(ctx context.Context) VpnClientRootCertificateResponseOutput {
	return o
}

func (o VpnClientRootCertificateResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VpnClientRootCertificateResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VpnClientRootCertificateResponseOutput) PublicCertData() pulumi.StringOutput {
	return o.ApplyT(func(v VpnClientRootCertificateResponse) string { return v.PublicCertData }).(pulumi.StringOutput)
}

type VpnClientRootCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (VpnClientRootCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnClientRootCertificateResponse)(nil)).Elem()
}

func (o VpnClientRootCertificateResponseArrayOutput) ToVpnClientRootCertificateResponseArrayOutput() VpnClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnClientRootCertificateResponseArrayOutput) ToVpnClientRootCertificateResponseArrayOutputWithContext(ctx context.Context) VpnClientRootCertificateResponseArrayOutput {
	return o
}

func (o VpnClientRootCertificateResponseArrayOutput) Index(i pulumi.IntInput) VpnClientRootCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpnClientRootCertificateResponse {
		return vs[0].([]VpnClientRootCertificateResponse)[vs[1].(int)]
	}).(VpnClientRootCertificateResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AddressSpaceOutput{})
	pulumi.RegisterOutputType(AddressSpacePtrOutput{})
	pulumi.RegisterOutputType(AddressSpaceResponseOutput{})
	pulumi.RegisterOutputType(AddressSpaceResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayAuthenticationCertificateOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayAuthenticationCertificateArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayAuthenticationCertificateResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayAuthenticationCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendAddressOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendAddressArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendAddressPoolOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendAddressPoolArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendAddressPoolResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendAddressPoolResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendAddressResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendAddressResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendHttpSettingsOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendHttpSettingsArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendHttpSettingsResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayBackendHttpSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayConnectionDrainingOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayConnectionDrainingPtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayConnectionDrainingResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayConnectionDrainingResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFirewallDisabledRuleGroupOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFirewallDisabledRuleGroupArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFirewallDisabledRuleGroupResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFirewallDisabledRuleGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFrontendIPConfigurationOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFrontendIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFrontendIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFrontendIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFrontendPortOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFrontendPortArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFrontendPortResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayFrontendPortResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayHttpListenerOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayHttpListenerArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayHttpListenerResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayHttpListenerResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayIPConfigurationOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayPathRuleOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayPathRuleArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayPathRuleResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayPathRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayProbeOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayProbeArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayProbeResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayProbeResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayRequestRoutingRuleOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayRequestRoutingRuleArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayRequestRoutingRuleResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayRequestRoutingRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySkuOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySkuPtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySkuResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySkuResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslCertificateOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslCertificateArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslCertificateResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslPolicyOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslPolicyPtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslPolicyResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewaySslPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayUrlPathMapOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayUrlPathMapArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayUrlPathMapResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayUrlPathMapResponseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayWebApplicationFirewallConfigurationOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayWebApplicationFirewallConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayWebApplicationFirewallConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(BackendAddressPoolOutput{})
	pulumi.RegisterOutputType(BackendAddressPoolArrayOutput{})
	pulumi.RegisterOutputType(BackendAddressPoolResponseOutput{})
	pulumi.RegisterOutputType(BackendAddressPoolResponseArrayOutput{})
	pulumi.RegisterOutputType(BgpPeerStatusResponseOutput{})
	pulumi.RegisterOutputType(BgpPeerStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(BgpSettingsOutput{})
	pulumi.RegisterOutputType(BgpSettingsPtrOutput{})
	pulumi.RegisterOutputType(BgpSettingsResponseOutput{})
	pulumi.RegisterOutputType(BgpSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(DhcpOptionsOutput{})
	pulumi.RegisterOutputType(DhcpOptionsPtrOutput{})
	pulumi.RegisterOutputType(DhcpOptionsResponseOutput{})
	pulumi.RegisterOutputType(DhcpOptionsResponsePtrOutput{})
	pulumi.RegisterOutputType(DnsConfigOutput{})
	pulumi.RegisterOutputType(DnsConfigPtrOutput{})
	pulumi.RegisterOutputType(DnsConfigResponseOutput{})
	pulumi.RegisterOutputType(DnsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointTypeOutput{})
	pulumi.RegisterOutputType(EndpointTypeArrayOutput{})
	pulumi.RegisterOutputType(EndpointResponseOutput{})
	pulumi.RegisterOutputType(EndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitAuthorizationTypeOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitAuthorizationTypeArrayOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitAuthorizationResponseOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitAuthorizationResponseArrayOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringTypeOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringTypeArrayOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringConfigOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringConfigPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringConfigResponseOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringResponseOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringResponseArrayOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitServiceProviderPropertiesOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitServiceProviderPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitServiceProviderPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitSkuOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitSkuPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitSkuResponseOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitStatsOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitStatsPtrOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitStatsResponseOutput{})
	pulumi.RegisterOutputType(ExpressRouteCircuitStatsResponsePtrOutput{})
	pulumi.RegisterOutputType(FrontendIPConfigurationOutput{})
	pulumi.RegisterOutputType(FrontendIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(FrontendIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(FrontendIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(GatewayRouteResponseOutput{})
	pulumi.RegisterOutputType(GatewayRouteResponseArrayOutput{})
	pulumi.RegisterOutputType(IPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(IPConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(IPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(InboundNatPoolOutput{})
	pulumi.RegisterOutputType(InboundNatPoolArrayOutput{})
	pulumi.RegisterOutputType(InboundNatPoolResponseOutput{})
	pulumi.RegisterOutputType(InboundNatPoolResponseArrayOutput{})
	pulumi.RegisterOutputType(InboundNatRuleOutput{})
	pulumi.RegisterOutputType(InboundNatRuleArrayOutput{})
	pulumi.RegisterOutputType(InboundNatRuleResponseOutput{})
	pulumi.RegisterOutputType(InboundNatRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(IpsecPolicyOutput{})
	pulumi.RegisterOutputType(IpsecPolicyArrayOutput{})
	pulumi.RegisterOutputType(IpsecPolicyResponseOutput{})
	pulumi.RegisterOutputType(IpsecPolicyResponseArrayOutput{})
	pulumi.RegisterOutputType(Ipv6ExpressRouteCircuitPeeringConfigOutput{})
	pulumi.RegisterOutputType(Ipv6ExpressRouteCircuitPeeringConfigPtrOutput{})
	pulumi.RegisterOutputType(Ipv6ExpressRouteCircuitPeeringConfigResponseOutput{})
	pulumi.RegisterOutputType(Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancingRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(LocalNetworkGatewayTypeOutput{})
	pulumi.RegisterOutputType(LocalNetworkGatewayTypePtrOutput{})
	pulumi.RegisterOutputType(LocalNetworkGatewayResponseOutput{})
	pulumi.RegisterOutputType(LocalNetworkGatewayResponsePtrOutput{})
	pulumi.RegisterOutputType(MonitorConfigOutput{})
	pulumi.RegisterOutputType(MonitorConfigPtrOutput{})
	pulumi.RegisterOutputType(MonitorConfigResponseOutput{})
	pulumi.RegisterOutputType(MonitorConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceDnsSettingsOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceDnsSettingsPtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceDnsSettingsResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceDnsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceIPConfigurationOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkSecurityGroupTypeOutput{})
	pulumi.RegisterOutputType(NetworkSecurityGroupTypePtrOutput{})
	pulumi.RegisterOutputType(NetworkSecurityGroupResponseOutput{})
	pulumi.RegisterOutputType(NetworkSecurityGroupResponsePtrOutput{})
	pulumi.RegisterOutputType(OutboundNatRuleOutput{})
	pulumi.RegisterOutputType(OutboundNatRuleArrayOutput{})
	pulumi.RegisterOutputType(OutboundNatRuleResponseOutput{})
	pulumi.RegisterOutputType(OutboundNatRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(PacketCaptureFilterOutput{})
	pulumi.RegisterOutputType(PacketCaptureFilterArrayOutput{})
	pulumi.RegisterOutputType(PacketCaptureFilterResponseOutput{})
	pulumi.RegisterOutputType(PacketCaptureFilterResponseArrayOutput{})
	pulumi.RegisterOutputType(PacketCaptureStorageLocationOutput{})
	pulumi.RegisterOutputType(PacketCaptureStorageLocationResponseOutput{})
	pulumi.RegisterOutputType(ProbeOutput{})
	pulumi.RegisterOutputType(ProbeArrayOutput{})
	pulumi.RegisterOutputType(ProbeResponseOutput{})
	pulumi.RegisterOutputType(ProbeResponseArrayOutput{})
	pulumi.RegisterOutputType(PublicIPAddressTypeOutput{})
	pulumi.RegisterOutputType(PublicIPAddressTypePtrOutput{})
	pulumi.RegisterOutputType(PublicIPAddressDnsSettingsOutput{})
	pulumi.RegisterOutputType(PublicIPAddressDnsSettingsPtrOutput{})
	pulumi.RegisterOutputType(PublicIPAddressDnsSettingsResponseOutput{})
	pulumi.RegisterOutputType(PublicIPAddressDnsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(PublicIPAddressResponseOutput{})
	pulumi.RegisterOutputType(PublicIPAddressResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceNavigationLinkOutput{})
	pulumi.RegisterOutputType(ResourceNavigationLinkArrayOutput{})
	pulumi.RegisterOutputType(ResourceNavigationLinkResponseOutput{})
	pulumi.RegisterOutputType(ResourceNavigationLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(RouteTypeOutput{})
	pulumi.RegisterOutputType(RouteTypeArrayOutput{})
	pulumi.RegisterOutputType(RouteFilterTypeOutput{})
	pulumi.RegisterOutputType(RouteFilterTypePtrOutput{})
	pulumi.RegisterOutputType(RouteFilterResponseOutput{})
	pulumi.RegisterOutputType(RouteFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(RouteFilterRuleTypeOutput{})
	pulumi.RegisterOutputType(RouteFilterRuleTypeArrayOutput{})
	pulumi.RegisterOutputType(RouteFilterRuleResponseOutput{})
	pulumi.RegisterOutputType(RouteFilterRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(RouteResponseOutput{})
	pulumi.RegisterOutputType(RouteResponseArrayOutput{})
	pulumi.RegisterOutputType(RouteTableTypeOutput{})
	pulumi.RegisterOutputType(RouteTableTypePtrOutput{})
	pulumi.RegisterOutputType(RouteTableResponseOutput{})
	pulumi.RegisterOutputType(RouteTableResponsePtrOutput{})
	pulumi.RegisterOutputType(SecurityRuleTypeOutput{})
	pulumi.RegisterOutputType(SecurityRuleTypeArrayOutput{})
	pulumi.RegisterOutputType(SecurityRuleResponseOutput{})
	pulumi.RegisterOutputType(SecurityRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourcePtrOutput{})
	pulumi.RegisterOutputType(SubResourceArrayOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SubnetTypeOutput{})
	pulumi.RegisterOutputType(SubnetTypePtrOutput{})
	pulumi.RegisterOutputType(SubnetTypeArrayOutput{})
	pulumi.RegisterOutputType(SubnetResponseOutput{})
	pulumi.RegisterOutputType(SubnetResponsePtrOutput{})
	pulumi.RegisterOutputType(SubnetResponseArrayOutput{})
	pulumi.RegisterOutputType(TunnelConnectionHealthResponseOutput{})
	pulumi.RegisterOutputType(TunnelConnectionHealthResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayTypeOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayTypePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayIPConfigurationOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewayResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkGatewaySkuResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringTypeOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringTypeArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnClientConfigurationOutput{})
	pulumi.RegisterOutputType(VpnClientConfigurationPtrOutput{})
	pulumi.RegisterOutputType(VpnClientConfigurationResponseOutput{})
	pulumi.RegisterOutputType(VpnClientConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(VpnClientRevokedCertificateOutput{})
	pulumi.RegisterOutputType(VpnClientRevokedCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnClientRevokedCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnClientRevokedCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(VpnClientRootCertificateOutput{})
	pulumi.RegisterOutputType(VpnClientRootCertificateArrayOutput{})
	pulumi.RegisterOutputType(VpnClientRootCertificateResponseOutput{})
	pulumi.RegisterOutputType(VpnClientRootCertificateResponseArrayOutput{})
}
