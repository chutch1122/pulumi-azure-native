


package v20200901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ArmIdentity struct {
	IdentityType           *string                `pulumi:"identityType"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ArmIdentityInput interface {
	pulumi.Input

	ToArmIdentityOutput() ArmIdentityOutput
	ToArmIdentityOutputWithContext(context.Context) ArmIdentityOutput
}

type ArmIdentityArgs struct {
	IdentityType           pulumi.StringPtrInput `pulumi:"identityType"`
	UserAssignedIdentities pulumi.MapInput       `pulumi:"userAssignedIdentities"`
}

func (ArmIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmIdentity)(nil)).Elem()
}

func (i ArmIdentityArgs) ToArmIdentityOutput() ArmIdentityOutput {
	return i.ToArmIdentityOutputWithContext(context.Background())
}

func (i ArmIdentityArgs) ToArmIdentityOutputWithContext(ctx context.Context) ArmIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmIdentityOutput)
}

func (i ArmIdentityArgs) ToArmIdentityPtrOutput() ArmIdentityPtrOutput {
	return i.ToArmIdentityPtrOutputWithContext(context.Background())
}

func (i ArmIdentityArgs) ToArmIdentityPtrOutputWithContext(ctx context.Context) ArmIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmIdentityOutput).ToArmIdentityPtrOutputWithContext(ctx)
}









type ArmIdentityPtrInput interface {
	pulumi.Input

	ToArmIdentityPtrOutput() ArmIdentityPtrOutput
	ToArmIdentityPtrOutputWithContext(context.Context) ArmIdentityPtrOutput
}

type armIdentityPtrType ArmIdentityArgs

func ArmIdentityPtr(v *ArmIdentityArgs) ArmIdentityPtrInput {
	return (*armIdentityPtrType)(v)
}

func (*armIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmIdentity)(nil)).Elem()
}

func (i *armIdentityPtrType) ToArmIdentityPtrOutput() ArmIdentityPtrOutput {
	return i.ToArmIdentityPtrOutputWithContext(context.Background())
}

func (i *armIdentityPtrType) ToArmIdentityPtrOutputWithContext(ctx context.Context) ArmIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmIdentityPtrOutput)
}

type ArmIdentityOutput struct{ *pulumi.OutputState }

func (ArmIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmIdentity)(nil)).Elem()
}

func (o ArmIdentityOutput) ToArmIdentityOutput() ArmIdentityOutput {
	return o
}

func (o ArmIdentityOutput) ToArmIdentityOutputWithContext(ctx context.Context) ArmIdentityOutput {
	return o
}

func (o ArmIdentityOutput) ToArmIdentityPtrOutput() ArmIdentityPtrOutput {
	return o.ToArmIdentityPtrOutputWithContext(context.Background())
}

func (o ArmIdentityOutput) ToArmIdentityPtrOutputWithContext(ctx context.Context) ArmIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArmIdentity) *ArmIdentity {
		return &v
	}).(ArmIdentityPtrOutput)
}

func (o ArmIdentityOutput) IdentityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmIdentity) *string { return v.IdentityType }).(pulumi.StringPtrOutput)
}

func (o ArmIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ArmIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ArmIdentityPtrOutput struct{ *pulumi.OutputState }

func (ArmIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmIdentity)(nil)).Elem()
}

func (o ArmIdentityPtrOutput) ToArmIdentityPtrOutput() ArmIdentityPtrOutput {
	return o
}

func (o ArmIdentityPtrOutput) ToArmIdentityPtrOutputWithContext(ctx context.Context) ArmIdentityPtrOutput {
	return o
}

func (o ArmIdentityPtrOutput) Elem() ArmIdentityOutput {
	return o.ApplyT(func(v *ArmIdentity) ArmIdentity {
		if v != nil {
			return *v
		}
		var ret ArmIdentity
		return ret
	}).(ArmIdentityOutput)
}

func (o ArmIdentityPtrOutput) IdentityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmIdentity) *string {
		if v == nil {
			return nil
		}
		return v.IdentityType
	}).(pulumi.StringPtrOutput)
}

func (o ArmIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ArmIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ArmIdentityResponse struct {
	IdentityType           *string                            `pulumi:"identityType"`
	PrincipalId            string                             `pulumi:"principalId"`
	TenantId               string                             `pulumi:"tenantId"`
	UserAssignedIdentities map[string]ArmUserIdentityResponse `pulumi:"userAssignedIdentities"`
}

type ArmIdentityResponseOutput struct{ *pulumi.OutputState }

func (ArmIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmIdentityResponse)(nil)).Elem()
}

func (o ArmIdentityResponseOutput) ToArmIdentityResponseOutput() ArmIdentityResponseOutput {
	return o
}

func (o ArmIdentityResponseOutput) ToArmIdentityResponseOutputWithContext(ctx context.Context) ArmIdentityResponseOutput {
	return o
}

func (o ArmIdentityResponseOutput) IdentityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmIdentityResponse) *string { return v.IdentityType }).(pulumi.StringPtrOutput)
}

func (o ArmIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ArmIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ArmIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ArmIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ArmIdentityResponseOutput) UserAssignedIdentities() ArmUserIdentityResponseMapOutput {
	return o.ApplyT(func(v ArmIdentityResponse) map[string]ArmUserIdentityResponse { return v.UserAssignedIdentities }).(ArmUserIdentityResponseMapOutput)
}

type ArmIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ArmIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmIdentityResponse)(nil)).Elem()
}

func (o ArmIdentityResponsePtrOutput) ToArmIdentityResponsePtrOutput() ArmIdentityResponsePtrOutput {
	return o
}

func (o ArmIdentityResponsePtrOutput) ToArmIdentityResponsePtrOutputWithContext(ctx context.Context) ArmIdentityResponsePtrOutput {
	return o
}

func (o ArmIdentityResponsePtrOutput) Elem() ArmIdentityResponseOutput {
	return o.ApplyT(func(v *ArmIdentityResponse) ArmIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ArmIdentityResponse
		return ret
	}).(ArmIdentityResponseOutput)
}

func (o ArmIdentityResponsePtrOutput) IdentityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.IdentityType
	}).(pulumi.StringPtrOutput)
}

func (o ArmIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ArmIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ArmIdentityResponsePtrOutput) UserAssignedIdentities() ArmUserIdentityResponseMapOutput {
	return o.ApplyT(func(v *ArmIdentityResponse) map[string]ArmUserIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(ArmUserIdentityResponseMapOutput)
}

type ArmUserIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type ArmUserIdentityResponseOutput struct{ *pulumi.OutputState }

func (ArmUserIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmUserIdentityResponse)(nil)).Elem()
}

func (o ArmUserIdentityResponseOutput) ToArmUserIdentityResponseOutput() ArmUserIdentityResponseOutput {
	return o
}

func (o ArmUserIdentityResponseOutput) ToArmUserIdentityResponseOutputWithContext(ctx context.Context) ArmUserIdentityResponseOutput {
	return o
}

func (o ArmUserIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ArmUserIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ArmUserIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ArmUserIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type ArmUserIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (ArmUserIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ArmUserIdentityResponse)(nil)).Elem()
}

func (o ArmUserIdentityResponseMapOutput) ToArmUserIdentityResponseMapOutput() ArmUserIdentityResponseMapOutput {
	return o
}

func (o ArmUserIdentityResponseMapOutput) ToArmUserIdentityResponseMapOutputWithContext(ctx context.Context) ArmUserIdentityResponseMapOutput {
	return o
}

func (o ArmUserIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) ArmUserIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ArmUserIdentityResponse {
		return vs[0].(map[string]ArmUserIdentityResponse)[vs[1].(string)]
	}).(ArmUserIdentityResponseOutput)
}

type CertificatePropertiesResponse struct {
	Created    string `pulumi:"created"`
	Expiry     string `pulumi:"expiry"`
	IsVerified bool   `pulumi:"isVerified"`
	Subject    string `pulumi:"subject"`
	Thumbprint string `pulumi:"thumbprint"`
	Updated    string `pulumi:"updated"`
}

type CertificatePropertiesResponseOutput struct{ *pulumi.OutputState }

func (CertificatePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificatePropertiesResponse)(nil)).Elem()
}

func (o CertificatePropertiesResponseOutput) ToCertificatePropertiesResponseOutput() CertificatePropertiesResponseOutput {
	return o
}

func (o CertificatePropertiesResponseOutput) ToCertificatePropertiesResponseOutputWithContext(ctx context.Context) CertificatePropertiesResponseOutput {
	return o
}

func (o CertificatePropertiesResponseOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Created }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) Expiry() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Expiry }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) IsVerified() pulumi.BoolOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) bool { return v.IsVerified }).(pulumi.BoolOutput)
}

func (o CertificatePropertiesResponseOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Subject }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) Updated() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Updated }).(pulumi.StringOutput)
}

type EncryptionKeyIdentity struct {
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}





type EncryptionKeyIdentityInput interface {
	pulumi.Input

	ToEncryptionKeyIdentityOutput() EncryptionKeyIdentityOutput
	ToEncryptionKeyIdentityOutputWithContext(context.Context) EncryptionKeyIdentityOutput
}

type EncryptionKeyIdentityArgs struct {
	UserAssignedIdentity pulumi.StringPtrInput `pulumi:"userAssignedIdentity"`
}

func (EncryptionKeyIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeyIdentity)(nil)).Elem()
}

func (i EncryptionKeyIdentityArgs) ToEncryptionKeyIdentityOutput() EncryptionKeyIdentityOutput {
	return i.ToEncryptionKeyIdentityOutputWithContext(context.Background())
}

func (i EncryptionKeyIdentityArgs) ToEncryptionKeyIdentityOutputWithContext(ctx context.Context) EncryptionKeyIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyIdentityOutput)
}

func (i EncryptionKeyIdentityArgs) ToEncryptionKeyIdentityPtrOutput() EncryptionKeyIdentityPtrOutput {
	return i.ToEncryptionKeyIdentityPtrOutputWithContext(context.Background())
}

func (i EncryptionKeyIdentityArgs) ToEncryptionKeyIdentityPtrOutputWithContext(ctx context.Context) EncryptionKeyIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyIdentityOutput).ToEncryptionKeyIdentityPtrOutputWithContext(ctx)
}









type EncryptionKeyIdentityPtrInput interface {
	pulumi.Input

	ToEncryptionKeyIdentityPtrOutput() EncryptionKeyIdentityPtrOutput
	ToEncryptionKeyIdentityPtrOutputWithContext(context.Context) EncryptionKeyIdentityPtrOutput
}

type encryptionKeyIdentityPtrType EncryptionKeyIdentityArgs

func EncryptionKeyIdentityPtr(v *EncryptionKeyIdentityArgs) EncryptionKeyIdentityPtrInput {
	return (*encryptionKeyIdentityPtrType)(v)
}

func (*encryptionKeyIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionKeyIdentity)(nil)).Elem()
}

func (i *encryptionKeyIdentityPtrType) ToEncryptionKeyIdentityPtrOutput() EncryptionKeyIdentityPtrOutput {
	return i.ToEncryptionKeyIdentityPtrOutputWithContext(context.Background())
}

func (i *encryptionKeyIdentityPtrType) ToEncryptionKeyIdentityPtrOutputWithContext(ctx context.Context) EncryptionKeyIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyIdentityPtrOutput)
}

type EncryptionKeyIdentityOutput struct{ *pulumi.OutputState }

func (EncryptionKeyIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeyIdentity)(nil)).Elem()
}

func (o EncryptionKeyIdentityOutput) ToEncryptionKeyIdentityOutput() EncryptionKeyIdentityOutput {
	return o
}

func (o EncryptionKeyIdentityOutput) ToEncryptionKeyIdentityOutputWithContext(ctx context.Context) EncryptionKeyIdentityOutput {
	return o
}

func (o EncryptionKeyIdentityOutput) ToEncryptionKeyIdentityPtrOutput() EncryptionKeyIdentityPtrOutput {
	return o.ToEncryptionKeyIdentityPtrOutputWithContext(context.Background())
}

func (o EncryptionKeyIdentityOutput) ToEncryptionKeyIdentityPtrOutputWithContext(ctx context.Context) EncryptionKeyIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionKeyIdentity) *EncryptionKeyIdentity {
		return &v
	}).(EncryptionKeyIdentityPtrOutput)
}

func (o EncryptionKeyIdentityOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyIdentity) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type EncryptionKeyIdentityPtrOutput struct{ *pulumi.OutputState }

func (EncryptionKeyIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionKeyIdentity)(nil)).Elem()
}

func (o EncryptionKeyIdentityPtrOutput) ToEncryptionKeyIdentityPtrOutput() EncryptionKeyIdentityPtrOutput {
	return o
}

func (o EncryptionKeyIdentityPtrOutput) ToEncryptionKeyIdentityPtrOutputWithContext(ctx context.Context) EncryptionKeyIdentityPtrOutput {
	return o
}

func (o EncryptionKeyIdentityPtrOutput) Elem() EncryptionKeyIdentityOutput {
	return o.ApplyT(func(v *EncryptionKeyIdentity) EncryptionKeyIdentity {
		if v != nil {
			return *v
		}
		var ret EncryptionKeyIdentity
		return ret
	}).(EncryptionKeyIdentityOutput)
}

func (o EncryptionKeyIdentityPtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyIdentity) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type EncryptionKeyIdentityResponse struct {
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}

type EncryptionKeyIdentityResponseOutput struct{ *pulumi.OutputState }

func (EncryptionKeyIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeyIdentityResponse)(nil)).Elem()
}

func (o EncryptionKeyIdentityResponseOutput) ToEncryptionKeyIdentityResponseOutput() EncryptionKeyIdentityResponseOutput {
	return o
}

func (o EncryptionKeyIdentityResponseOutput) ToEncryptionKeyIdentityResponseOutputWithContext(ctx context.Context) EncryptionKeyIdentityResponseOutput {
	return o
}

func (o EncryptionKeyIdentityResponseOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyIdentityResponse) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type EncryptionKeyIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionKeyIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionKeyIdentityResponse)(nil)).Elem()
}

func (o EncryptionKeyIdentityResponsePtrOutput) ToEncryptionKeyIdentityResponsePtrOutput() EncryptionKeyIdentityResponsePtrOutput {
	return o
}

func (o EncryptionKeyIdentityResponsePtrOutput) ToEncryptionKeyIdentityResponsePtrOutputWithContext(ctx context.Context) EncryptionKeyIdentityResponsePtrOutput {
	return o
}

func (o EncryptionKeyIdentityResponsePtrOutput) Elem() EncryptionKeyIdentityResponseOutput {
	return o.ApplyT(func(v *EncryptionKeyIdentityResponse) EncryptionKeyIdentityResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionKeyIdentityResponse
		return ret
	}).(EncryptionKeyIdentityResponseOutput)
}

func (o EncryptionKeyIdentityResponsePtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type EncryptionPropertiesDescription struct {
	Identity           *EncryptionKeyIdentity  `pulumi:"identity"`
	KeySource          *string                 `pulumi:"keySource"`
	KeyVaultProperties []KeyVaultKeyProperties `pulumi:"keyVaultProperties"`
}





type EncryptionPropertiesDescriptionInput interface {
	pulumi.Input

	ToEncryptionPropertiesDescriptionOutput() EncryptionPropertiesDescriptionOutput
	ToEncryptionPropertiesDescriptionOutputWithContext(context.Context) EncryptionPropertiesDescriptionOutput
}

type EncryptionPropertiesDescriptionArgs struct {
	Identity           EncryptionKeyIdentityPtrInput   `pulumi:"identity"`
	KeySource          pulumi.StringPtrInput           `pulumi:"keySource"`
	KeyVaultProperties KeyVaultKeyPropertiesArrayInput `pulumi:"keyVaultProperties"`
}

func (EncryptionPropertiesDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertiesDescription)(nil)).Elem()
}

func (i EncryptionPropertiesDescriptionArgs) ToEncryptionPropertiesDescriptionOutput() EncryptionPropertiesDescriptionOutput {
	return i.ToEncryptionPropertiesDescriptionOutputWithContext(context.Background())
}

func (i EncryptionPropertiesDescriptionArgs) ToEncryptionPropertiesDescriptionOutputWithContext(ctx context.Context) EncryptionPropertiesDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesDescriptionOutput)
}

func (i EncryptionPropertiesDescriptionArgs) ToEncryptionPropertiesDescriptionPtrOutput() EncryptionPropertiesDescriptionPtrOutput {
	return i.ToEncryptionPropertiesDescriptionPtrOutputWithContext(context.Background())
}

func (i EncryptionPropertiesDescriptionArgs) ToEncryptionPropertiesDescriptionPtrOutputWithContext(ctx context.Context) EncryptionPropertiesDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesDescriptionOutput).ToEncryptionPropertiesDescriptionPtrOutputWithContext(ctx)
}









type EncryptionPropertiesDescriptionPtrInput interface {
	pulumi.Input

	ToEncryptionPropertiesDescriptionPtrOutput() EncryptionPropertiesDescriptionPtrOutput
	ToEncryptionPropertiesDescriptionPtrOutputWithContext(context.Context) EncryptionPropertiesDescriptionPtrOutput
}

type encryptionPropertiesDescriptionPtrType EncryptionPropertiesDescriptionArgs

func EncryptionPropertiesDescriptionPtr(v *EncryptionPropertiesDescriptionArgs) EncryptionPropertiesDescriptionPtrInput {
	return (*encryptionPropertiesDescriptionPtrType)(v)
}

func (*encryptionPropertiesDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertiesDescription)(nil)).Elem()
}

func (i *encryptionPropertiesDescriptionPtrType) ToEncryptionPropertiesDescriptionPtrOutput() EncryptionPropertiesDescriptionPtrOutput {
	return i.ToEncryptionPropertiesDescriptionPtrOutputWithContext(context.Background())
}

func (i *encryptionPropertiesDescriptionPtrType) ToEncryptionPropertiesDescriptionPtrOutputWithContext(ctx context.Context) EncryptionPropertiesDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesDescriptionPtrOutput)
}

type EncryptionPropertiesDescriptionOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertiesDescription)(nil)).Elem()
}

func (o EncryptionPropertiesDescriptionOutput) ToEncryptionPropertiesDescriptionOutput() EncryptionPropertiesDescriptionOutput {
	return o
}

func (o EncryptionPropertiesDescriptionOutput) ToEncryptionPropertiesDescriptionOutputWithContext(ctx context.Context) EncryptionPropertiesDescriptionOutput {
	return o
}

func (o EncryptionPropertiesDescriptionOutput) ToEncryptionPropertiesDescriptionPtrOutput() EncryptionPropertiesDescriptionPtrOutput {
	return o.ToEncryptionPropertiesDescriptionPtrOutputWithContext(context.Background())
}

func (o EncryptionPropertiesDescriptionOutput) ToEncryptionPropertiesDescriptionPtrOutputWithContext(ctx context.Context) EncryptionPropertiesDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionPropertiesDescription) *EncryptionPropertiesDescription {
		return &v
	}).(EncryptionPropertiesDescriptionPtrOutput)
}

func (o EncryptionPropertiesDescriptionOutput) Identity() EncryptionKeyIdentityPtrOutput {
	return o.ApplyT(func(v EncryptionPropertiesDescription) *EncryptionKeyIdentity { return v.Identity }).(EncryptionKeyIdentityPtrOutput)
}

func (o EncryptionPropertiesDescriptionOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionPropertiesDescription) *string { return v.KeySource }).(pulumi.StringPtrOutput)
}

func (o EncryptionPropertiesDescriptionOutput) KeyVaultProperties() KeyVaultKeyPropertiesArrayOutput {
	return o.ApplyT(func(v EncryptionPropertiesDescription) []KeyVaultKeyProperties { return v.KeyVaultProperties }).(KeyVaultKeyPropertiesArrayOutput)
}

type EncryptionPropertiesDescriptionPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertiesDescription)(nil)).Elem()
}

func (o EncryptionPropertiesDescriptionPtrOutput) ToEncryptionPropertiesDescriptionPtrOutput() EncryptionPropertiesDescriptionPtrOutput {
	return o
}

func (o EncryptionPropertiesDescriptionPtrOutput) ToEncryptionPropertiesDescriptionPtrOutputWithContext(ctx context.Context) EncryptionPropertiesDescriptionPtrOutput {
	return o
}

func (o EncryptionPropertiesDescriptionPtrOutput) Elem() EncryptionPropertiesDescriptionOutput {
	return o.ApplyT(func(v *EncryptionPropertiesDescription) EncryptionPropertiesDescription {
		if v != nil {
			return *v
		}
		var ret EncryptionPropertiesDescription
		return ret
	}).(EncryptionPropertiesDescriptionOutput)
}

func (o EncryptionPropertiesDescriptionPtrOutput) Identity() EncryptionKeyIdentityPtrOutput {
	return o.ApplyT(func(v *EncryptionPropertiesDescription) *EncryptionKeyIdentity {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(EncryptionKeyIdentityPtrOutput)
}

func (o EncryptionPropertiesDescriptionPtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionPropertiesDescription) *string {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPropertiesDescriptionPtrOutput) KeyVaultProperties() KeyVaultKeyPropertiesArrayOutput {
	return o.ApplyT(func(v *EncryptionPropertiesDescription) []KeyVaultKeyProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultKeyPropertiesArrayOutput)
}

type EncryptionPropertiesDescriptionResponse struct {
	Identity           *EncryptionKeyIdentityResponse  `pulumi:"identity"`
	KeySource          *string                         `pulumi:"keySource"`
	KeyVaultProperties []KeyVaultKeyPropertiesResponse `pulumi:"keyVaultProperties"`
}

type EncryptionPropertiesDescriptionResponseOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertiesDescriptionResponse)(nil)).Elem()
}

func (o EncryptionPropertiesDescriptionResponseOutput) ToEncryptionPropertiesDescriptionResponseOutput() EncryptionPropertiesDescriptionResponseOutput {
	return o
}

func (o EncryptionPropertiesDescriptionResponseOutput) ToEncryptionPropertiesDescriptionResponseOutputWithContext(ctx context.Context) EncryptionPropertiesDescriptionResponseOutput {
	return o
}

func (o EncryptionPropertiesDescriptionResponseOutput) Identity() EncryptionKeyIdentityResponsePtrOutput {
	return o.ApplyT(func(v EncryptionPropertiesDescriptionResponse) *EncryptionKeyIdentityResponse { return v.Identity }).(EncryptionKeyIdentityResponsePtrOutput)
}

func (o EncryptionPropertiesDescriptionResponseOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionPropertiesDescriptionResponse) *string { return v.KeySource }).(pulumi.StringPtrOutput)
}

func (o EncryptionPropertiesDescriptionResponseOutput) KeyVaultProperties() KeyVaultKeyPropertiesResponseArrayOutput {
	return o.ApplyT(func(v EncryptionPropertiesDescriptionResponse) []KeyVaultKeyPropertiesResponse {
		return v.KeyVaultProperties
	}).(KeyVaultKeyPropertiesResponseArrayOutput)
}

type EncryptionPropertiesDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertiesDescriptionResponse)(nil)).Elem()
}

func (o EncryptionPropertiesDescriptionResponsePtrOutput) ToEncryptionPropertiesDescriptionResponsePtrOutput() EncryptionPropertiesDescriptionResponsePtrOutput {
	return o
}

func (o EncryptionPropertiesDescriptionResponsePtrOutput) ToEncryptionPropertiesDescriptionResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertiesDescriptionResponsePtrOutput {
	return o
}

func (o EncryptionPropertiesDescriptionResponsePtrOutput) Elem() EncryptionPropertiesDescriptionResponseOutput {
	return o.ApplyT(func(v *EncryptionPropertiesDescriptionResponse) EncryptionPropertiesDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionPropertiesDescriptionResponse
		return ret
	}).(EncryptionPropertiesDescriptionResponseOutput)
}

func (o EncryptionPropertiesDescriptionResponsePtrOutput) Identity() EncryptionKeyIdentityResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionPropertiesDescriptionResponse) *EncryptionKeyIdentityResponse {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(EncryptionKeyIdentityResponsePtrOutput)
}

func (o EncryptionPropertiesDescriptionResponsePtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionPropertiesDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPropertiesDescriptionResponsePtrOutput) KeyVaultProperties() KeyVaultKeyPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *EncryptionPropertiesDescriptionResponse) []KeyVaultKeyPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultKeyPropertiesResponseArrayOutput)
}

type IotDpsPropertiesDescription struct {
	AllocationPolicy           *string                                                         `pulumi:"allocationPolicy"`
	AuthorizationPolicies      []SharedAccessSignatureAuthorizationRuleAccessRightsDescription `pulumi:"authorizationPolicies"`
	Encryption                 *EncryptionPropertiesDescription                                `pulumi:"encryption"`
	IotHubs                    []IotHubDefinitionDescription                                   `pulumi:"iotHubs"`
	IpFilterRules              []TargetIpFilterRule                                            `pulumi:"ipFilterRules"`
	PrivateEndpointConnections []PrivateEndpointConnection                                     `pulumi:"privateEndpointConnections"`
	ProvisioningState          *string                                                         `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                                                         `pulumi:"publicNetworkAccess"`
	State                      *string                                                         `pulumi:"state"`
}





type IotDpsPropertiesDescriptionInput interface {
	pulumi.Input

	ToIotDpsPropertiesDescriptionOutput() IotDpsPropertiesDescriptionOutput
	ToIotDpsPropertiesDescriptionOutputWithContext(context.Context) IotDpsPropertiesDescriptionOutput
}

type IotDpsPropertiesDescriptionArgs struct {
	AllocationPolicy           pulumi.StringPtrInput                                                   `pulumi:"allocationPolicy"`
	AuthorizationPolicies      SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayInput `pulumi:"authorizationPolicies"`
	Encryption                 EncryptionPropertiesDescriptionPtrInput                                 `pulumi:"encryption"`
	IotHubs                    IotHubDefinitionDescriptionArrayInput                                   `pulumi:"iotHubs"`
	IpFilterRules              TargetIpFilterRuleArrayInput                                            `pulumi:"ipFilterRules"`
	PrivateEndpointConnections PrivateEndpointConnectionArrayInput                                     `pulumi:"privateEndpointConnections"`
	ProvisioningState          pulumi.StringPtrInput                                                   `pulumi:"provisioningState"`
	PublicNetworkAccess        pulumi.StringPtrInput                                                   `pulumi:"publicNetworkAccess"`
	State                      pulumi.StringPtrInput                                                   `pulumi:"state"`
}

func (IotDpsPropertiesDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsPropertiesDescription)(nil)).Elem()
}

func (i IotDpsPropertiesDescriptionArgs) ToIotDpsPropertiesDescriptionOutput() IotDpsPropertiesDescriptionOutput {
	return i.ToIotDpsPropertiesDescriptionOutputWithContext(context.Background())
}

func (i IotDpsPropertiesDescriptionArgs) ToIotDpsPropertiesDescriptionOutputWithContext(ctx context.Context) IotDpsPropertiesDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDpsPropertiesDescriptionOutput)
}

type IotDpsPropertiesDescriptionOutput struct{ *pulumi.OutputState }

func (IotDpsPropertiesDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsPropertiesDescription)(nil)).Elem()
}

func (o IotDpsPropertiesDescriptionOutput) ToIotDpsPropertiesDescriptionOutput() IotDpsPropertiesDescriptionOutput {
	return o
}

func (o IotDpsPropertiesDescriptionOutput) ToIotDpsPropertiesDescriptionOutputWithContext(ctx context.Context) IotDpsPropertiesDescriptionOutput {
	return o
}

func (o IotDpsPropertiesDescriptionOutput) AllocationPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) *string { return v.AllocationPolicy }).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionOutput) AuthorizationPolicies() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) []SharedAccessSignatureAuthorizationRuleAccessRightsDescription {
		return v.AuthorizationPolicies
	}).(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput)
}

func (o IotDpsPropertiesDescriptionOutput) Encryption() EncryptionPropertiesDescriptionPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) *EncryptionPropertiesDescription { return v.Encryption }).(EncryptionPropertiesDescriptionPtrOutput)
}

func (o IotDpsPropertiesDescriptionOutput) IotHubs() IotHubDefinitionDescriptionArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) []IotHubDefinitionDescription { return v.IotHubs }).(IotHubDefinitionDescriptionArrayOutput)
}

func (o IotDpsPropertiesDescriptionOutput) IpFilterRules() TargetIpFilterRuleArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) []TargetIpFilterRule { return v.IpFilterRules }).(TargetIpFilterRuleArrayOutput)
}

func (o IotDpsPropertiesDescriptionOutput) PrivateEndpointConnections() PrivateEndpointConnectionArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) []PrivateEndpointConnection { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionArrayOutput)
}

func (o IotDpsPropertiesDescriptionOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescription) *string { return v.State }).(pulumi.StringPtrOutput)
}

type IotDpsPropertiesDescriptionResponse struct {
	AllocationPolicy           *string                                                                 `pulumi:"allocationPolicy"`
	AuthorizationPolicies      []SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse `pulumi:"authorizationPolicies"`
	DeviceProvisioningHostName string                                                                  `pulumi:"deviceProvisioningHostName"`
	Encryption                 *EncryptionPropertiesDescriptionResponse                                `pulumi:"encryption"`
	IdScope                    string                                                                  `pulumi:"idScope"`
	IotHubs                    []IotHubDefinitionDescriptionResponse                                   `pulumi:"iotHubs"`
	IpFilterRules              []TargetIpFilterRuleResponse                                            `pulumi:"ipFilterRules"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse                                     `pulumi:"privateEndpointConnections"`
	ProvisioningState          *string                                                                 `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                                                                 `pulumi:"publicNetworkAccess"`
	ServiceOperationsHostName  string                                                                  `pulumi:"serviceOperationsHostName"`
	State                      *string                                                                 `pulumi:"state"`
}

type IotDpsPropertiesDescriptionResponseOutput struct{ *pulumi.OutputState }

func (IotDpsPropertiesDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsPropertiesDescriptionResponse)(nil)).Elem()
}

func (o IotDpsPropertiesDescriptionResponseOutput) ToIotDpsPropertiesDescriptionResponseOutput() IotDpsPropertiesDescriptionResponseOutput {
	return o
}

func (o IotDpsPropertiesDescriptionResponseOutput) ToIotDpsPropertiesDescriptionResponseOutputWithContext(ctx context.Context) IotDpsPropertiesDescriptionResponseOutput {
	return o
}

func (o IotDpsPropertiesDescriptionResponseOutput) AllocationPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) *string { return v.AllocationPolicy }).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) AuthorizationPolicies() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) []SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse {
		return v.AuthorizationPolicies
	}).(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) DeviceProvisioningHostName() pulumi.StringOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) string { return v.DeviceProvisioningHostName }).(pulumi.StringOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) Encryption() EncryptionPropertiesDescriptionResponsePtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) *EncryptionPropertiesDescriptionResponse {
		return v.Encryption
	}).(EncryptionPropertiesDescriptionResponsePtrOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) IdScope() pulumi.StringOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) string { return v.IdScope }).(pulumi.StringOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) IotHubs() IotHubDefinitionDescriptionResponseArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) []IotHubDefinitionDescriptionResponse { return v.IotHubs }).(IotHubDefinitionDescriptionResponseArrayOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) IpFilterRules() TargetIpFilterRuleResponseArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) []TargetIpFilterRuleResponse { return v.IpFilterRules }).(TargetIpFilterRuleResponseArrayOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) ServiceOperationsHostName() pulumi.StringOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) string { return v.ServiceOperationsHostName }).(pulumi.StringOutput)
}

func (o IotDpsPropertiesDescriptionResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsPropertiesDescriptionResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type IotDpsSkuInfo struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     *string  `pulumi:"name"`
}





type IotDpsSkuInfoInput interface {
	pulumi.Input

	ToIotDpsSkuInfoOutput() IotDpsSkuInfoOutput
	ToIotDpsSkuInfoOutputWithContext(context.Context) IotDpsSkuInfoOutput
}

type IotDpsSkuInfoArgs struct {
	Capacity pulumi.Float64PtrInput `pulumi:"capacity"`
	Name     pulumi.StringPtrInput  `pulumi:"name"`
}

func (IotDpsSkuInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsSkuInfo)(nil)).Elem()
}

func (i IotDpsSkuInfoArgs) ToIotDpsSkuInfoOutput() IotDpsSkuInfoOutput {
	return i.ToIotDpsSkuInfoOutputWithContext(context.Background())
}

func (i IotDpsSkuInfoArgs) ToIotDpsSkuInfoOutputWithContext(ctx context.Context) IotDpsSkuInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDpsSkuInfoOutput)
}

type IotDpsSkuInfoOutput struct{ *pulumi.OutputState }

func (IotDpsSkuInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsSkuInfo)(nil)).Elem()
}

func (o IotDpsSkuInfoOutput) ToIotDpsSkuInfoOutput() IotDpsSkuInfoOutput {
	return o
}

func (o IotDpsSkuInfoOutput) ToIotDpsSkuInfoOutputWithContext(ctx context.Context) IotDpsSkuInfoOutput {
	return o
}

func (o IotDpsSkuInfoOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v IotDpsSkuInfo) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o IotDpsSkuInfoOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsSkuInfo) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type IotDpsSkuInfoResponse struct {
	Capacity *float64 `pulumi:"capacity"`
	Name     *string  `pulumi:"name"`
	Tier     string   `pulumi:"tier"`
}

type IotDpsSkuInfoResponseOutput struct{ *pulumi.OutputState }

func (IotDpsSkuInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsSkuInfoResponse)(nil)).Elem()
}

func (o IotDpsSkuInfoResponseOutput) ToIotDpsSkuInfoResponseOutput() IotDpsSkuInfoResponseOutput {
	return o
}

func (o IotDpsSkuInfoResponseOutput) ToIotDpsSkuInfoResponseOutputWithContext(ctx context.Context) IotDpsSkuInfoResponseOutput {
	return o
}

func (o IotDpsSkuInfoResponseOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v IotDpsSkuInfoResponse) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
}

func (o IotDpsSkuInfoResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotDpsSkuInfoResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IotDpsSkuInfoResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v IotDpsSkuInfoResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type IotHubDefinitionDescription struct {
	AllocationWeight      *int   `pulumi:"allocationWeight"`
	ApplyAllocationPolicy *bool  `pulumi:"applyAllocationPolicy"`
	ConnectionString      string `pulumi:"connectionString"`
	Location              string `pulumi:"location"`
}





type IotHubDefinitionDescriptionInput interface {
	pulumi.Input

	ToIotHubDefinitionDescriptionOutput() IotHubDefinitionDescriptionOutput
	ToIotHubDefinitionDescriptionOutputWithContext(context.Context) IotHubDefinitionDescriptionOutput
}

type IotHubDefinitionDescriptionArgs struct {
	AllocationWeight      pulumi.IntPtrInput  `pulumi:"allocationWeight"`
	ApplyAllocationPolicy pulumi.BoolPtrInput `pulumi:"applyAllocationPolicy"`
	ConnectionString      pulumi.StringInput  `pulumi:"connectionString"`
	Location              pulumi.StringInput  `pulumi:"location"`
}

func (IotHubDefinitionDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubDefinitionDescription)(nil)).Elem()
}

func (i IotHubDefinitionDescriptionArgs) ToIotHubDefinitionDescriptionOutput() IotHubDefinitionDescriptionOutput {
	return i.ToIotHubDefinitionDescriptionOutputWithContext(context.Background())
}

func (i IotHubDefinitionDescriptionArgs) ToIotHubDefinitionDescriptionOutputWithContext(ctx context.Context) IotHubDefinitionDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubDefinitionDescriptionOutput)
}





type IotHubDefinitionDescriptionArrayInput interface {
	pulumi.Input

	ToIotHubDefinitionDescriptionArrayOutput() IotHubDefinitionDescriptionArrayOutput
	ToIotHubDefinitionDescriptionArrayOutputWithContext(context.Context) IotHubDefinitionDescriptionArrayOutput
}

type IotHubDefinitionDescriptionArray []IotHubDefinitionDescriptionInput

func (IotHubDefinitionDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IotHubDefinitionDescription)(nil)).Elem()
}

func (i IotHubDefinitionDescriptionArray) ToIotHubDefinitionDescriptionArrayOutput() IotHubDefinitionDescriptionArrayOutput {
	return i.ToIotHubDefinitionDescriptionArrayOutputWithContext(context.Background())
}

func (i IotHubDefinitionDescriptionArray) ToIotHubDefinitionDescriptionArrayOutputWithContext(ctx context.Context) IotHubDefinitionDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubDefinitionDescriptionArrayOutput)
}

type IotHubDefinitionDescriptionOutput struct{ *pulumi.OutputState }

func (IotHubDefinitionDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubDefinitionDescription)(nil)).Elem()
}

func (o IotHubDefinitionDescriptionOutput) ToIotHubDefinitionDescriptionOutput() IotHubDefinitionDescriptionOutput {
	return o
}

func (o IotHubDefinitionDescriptionOutput) ToIotHubDefinitionDescriptionOutputWithContext(ctx context.Context) IotHubDefinitionDescriptionOutput {
	return o
}

func (o IotHubDefinitionDescriptionOutput) AllocationWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IotHubDefinitionDescription) *int { return v.AllocationWeight }).(pulumi.IntPtrOutput)
}

func (o IotHubDefinitionDescriptionOutput) ApplyAllocationPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IotHubDefinitionDescription) *bool { return v.ApplyAllocationPolicy }).(pulumi.BoolPtrOutput)
}

func (o IotHubDefinitionDescriptionOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubDefinitionDescription) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o IotHubDefinitionDescriptionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubDefinitionDescription) string { return v.Location }).(pulumi.StringOutput)
}

type IotHubDefinitionDescriptionArrayOutput struct{ *pulumi.OutputState }

func (IotHubDefinitionDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IotHubDefinitionDescription)(nil)).Elem()
}

func (o IotHubDefinitionDescriptionArrayOutput) ToIotHubDefinitionDescriptionArrayOutput() IotHubDefinitionDescriptionArrayOutput {
	return o
}

func (o IotHubDefinitionDescriptionArrayOutput) ToIotHubDefinitionDescriptionArrayOutputWithContext(ctx context.Context) IotHubDefinitionDescriptionArrayOutput {
	return o
}

func (o IotHubDefinitionDescriptionArrayOutput) Index(i pulumi.IntInput) IotHubDefinitionDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IotHubDefinitionDescription {
		return vs[0].([]IotHubDefinitionDescription)[vs[1].(int)]
	}).(IotHubDefinitionDescriptionOutput)
}

type IotHubDefinitionDescriptionResponse struct {
	AllocationWeight      *int   `pulumi:"allocationWeight"`
	ApplyAllocationPolicy *bool  `pulumi:"applyAllocationPolicy"`
	ConnectionString      string `pulumi:"connectionString"`
	Location              string `pulumi:"location"`
	Name                  string `pulumi:"name"`
}

type IotHubDefinitionDescriptionResponseOutput struct{ *pulumi.OutputState }

func (IotHubDefinitionDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubDefinitionDescriptionResponse)(nil)).Elem()
}

func (o IotHubDefinitionDescriptionResponseOutput) ToIotHubDefinitionDescriptionResponseOutput() IotHubDefinitionDescriptionResponseOutput {
	return o
}

func (o IotHubDefinitionDescriptionResponseOutput) ToIotHubDefinitionDescriptionResponseOutputWithContext(ctx context.Context) IotHubDefinitionDescriptionResponseOutput {
	return o
}

func (o IotHubDefinitionDescriptionResponseOutput) AllocationWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IotHubDefinitionDescriptionResponse) *int { return v.AllocationWeight }).(pulumi.IntPtrOutput)
}

func (o IotHubDefinitionDescriptionResponseOutput) ApplyAllocationPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IotHubDefinitionDescriptionResponse) *bool { return v.ApplyAllocationPolicy }).(pulumi.BoolPtrOutput)
}

func (o IotHubDefinitionDescriptionResponseOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubDefinitionDescriptionResponse) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o IotHubDefinitionDescriptionResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubDefinitionDescriptionResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o IotHubDefinitionDescriptionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubDefinitionDescriptionResponse) string { return v.Name }).(pulumi.StringOutput)
}

type IotHubDefinitionDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (IotHubDefinitionDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IotHubDefinitionDescriptionResponse)(nil)).Elem()
}

func (o IotHubDefinitionDescriptionResponseArrayOutput) ToIotHubDefinitionDescriptionResponseArrayOutput() IotHubDefinitionDescriptionResponseArrayOutput {
	return o
}

func (o IotHubDefinitionDescriptionResponseArrayOutput) ToIotHubDefinitionDescriptionResponseArrayOutputWithContext(ctx context.Context) IotHubDefinitionDescriptionResponseArrayOutput {
	return o
}

func (o IotHubDefinitionDescriptionResponseArrayOutput) Index(i pulumi.IntInput) IotHubDefinitionDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IotHubDefinitionDescriptionResponse {
		return vs[0].([]IotHubDefinitionDescriptionResponse)[vs[1].(int)]
	}).(IotHubDefinitionDescriptionResponseOutput)
}

type KeyVaultKeyProperties struct {
	KeyIdentifier *string `pulumi:"keyIdentifier"`
}





type KeyVaultKeyPropertiesInput interface {
	pulumi.Input

	ToKeyVaultKeyPropertiesOutput() KeyVaultKeyPropertiesOutput
	ToKeyVaultKeyPropertiesOutputWithContext(context.Context) KeyVaultKeyPropertiesOutput
}

type KeyVaultKeyPropertiesArgs struct {
	KeyIdentifier pulumi.StringPtrInput `pulumi:"keyIdentifier"`
}

func (KeyVaultKeyPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyProperties)(nil)).Elem()
}

func (i KeyVaultKeyPropertiesArgs) ToKeyVaultKeyPropertiesOutput() KeyVaultKeyPropertiesOutput {
	return i.ToKeyVaultKeyPropertiesOutputWithContext(context.Background())
}

func (i KeyVaultKeyPropertiesArgs) ToKeyVaultKeyPropertiesOutputWithContext(ctx context.Context) KeyVaultKeyPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyPropertiesOutput)
}





type KeyVaultKeyPropertiesArrayInput interface {
	pulumi.Input

	ToKeyVaultKeyPropertiesArrayOutput() KeyVaultKeyPropertiesArrayOutput
	ToKeyVaultKeyPropertiesArrayOutputWithContext(context.Context) KeyVaultKeyPropertiesArrayOutput
}

type KeyVaultKeyPropertiesArray []KeyVaultKeyPropertiesInput

func (KeyVaultKeyPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyVaultKeyProperties)(nil)).Elem()
}

func (i KeyVaultKeyPropertiesArray) ToKeyVaultKeyPropertiesArrayOutput() KeyVaultKeyPropertiesArrayOutput {
	return i.ToKeyVaultKeyPropertiesArrayOutputWithContext(context.Background())
}

func (i KeyVaultKeyPropertiesArray) ToKeyVaultKeyPropertiesArrayOutputWithContext(ctx context.Context) KeyVaultKeyPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultKeyPropertiesArrayOutput)
}

type KeyVaultKeyPropertiesOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyProperties)(nil)).Elem()
}

func (o KeyVaultKeyPropertiesOutput) ToKeyVaultKeyPropertiesOutput() KeyVaultKeyPropertiesOutput {
	return o
}

func (o KeyVaultKeyPropertiesOutput) ToKeyVaultKeyPropertiesOutputWithContext(ctx context.Context) KeyVaultKeyPropertiesOutput {
	return o
}

func (o KeyVaultKeyPropertiesOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyProperties) *string { return v.KeyIdentifier }).(pulumi.StringPtrOutput)
}

type KeyVaultKeyPropertiesArrayOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyVaultKeyProperties)(nil)).Elem()
}

func (o KeyVaultKeyPropertiesArrayOutput) ToKeyVaultKeyPropertiesArrayOutput() KeyVaultKeyPropertiesArrayOutput {
	return o
}

func (o KeyVaultKeyPropertiesArrayOutput) ToKeyVaultKeyPropertiesArrayOutputWithContext(ctx context.Context) KeyVaultKeyPropertiesArrayOutput {
	return o
}

func (o KeyVaultKeyPropertiesArrayOutput) Index(i pulumi.IntInput) KeyVaultKeyPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyVaultKeyProperties {
		return vs[0].([]KeyVaultKeyProperties)[vs[1].(int)]
	}).(KeyVaultKeyPropertiesOutput)
}

type KeyVaultKeyPropertiesResponse struct {
	KeyIdentifier *string `pulumi:"keyIdentifier"`
}

type KeyVaultKeyPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultKeyPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultKeyPropertiesResponseOutput) ToKeyVaultKeyPropertiesResponseOutput() KeyVaultKeyPropertiesResponseOutput {
	return o
}

func (o KeyVaultKeyPropertiesResponseOutput) ToKeyVaultKeyPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultKeyPropertiesResponseOutput {
	return o
}

func (o KeyVaultKeyPropertiesResponseOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultKeyPropertiesResponse) *string { return v.KeyIdentifier }).(pulumi.StringPtrOutput)
}

type KeyVaultKeyPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (KeyVaultKeyPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyVaultKeyPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultKeyPropertiesResponseArrayOutput) ToKeyVaultKeyPropertiesResponseArrayOutput() KeyVaultKeyPropertiesResponseArrayOutput {
	return o
}

func (o KeyVaultKeyPropertiesResponseArrayOutput) ToKeyVaultKeyPropertiesResponseArrayOutputWithContext(ctx context.Context) KeyVaultKeyPropertiesResponseArrayOutput {
	return o
}

func (o KeyVaultKeyPropertiesResponseArrayOutput) Index(i pulumi.IntInput) KeyVaultKeyPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyVaultKeyPropertiesResponse {
		return vs[0].([]KeyVaultKeyPropertiesResponse)[vs[1].(int)]
	}).(KeyVaultKeyPropertiesResponseOutput)
}

type PrivateEndpointConnection struct {
	Properties PrivateEndpointConnectionProperties `pulumi:"properties"`
}





type PrivateEndpointConnectionInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput
	ToPrivateEndpointConnectionOutputWithContext(context.Context) PrivateEndpointConnectionOutput
}

type PrivateEndpointConnectionArgs struct {
	Properties PrivateEndpointConnectionPropertiesInput `pulumi:"properties"`
}

func (PrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnection)(nil)).Elem()
}

func (i PrivateEndpointConnectionArgs) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return i.ToPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionArgs) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionOutput)
}





type PrivateEndpointConnectionArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionArrayOutput() PrivateEndpointConnectionArrayOutput
	ToPrivateEndpointConnectionArrayOutputWithContext(context.Context) PrivateEndpointConnectionArrayOutput
}

type PrivateEndpointConnectionArray []PrivateEndpointConnectionInput

func (PrivateEndpointConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnection)(nil)).Elem()
}

func (i PrivateEndpointConnectionArray) ToPrivateEndpointConnectionArrayOutput() PrivateEndpointConnectionArrayOutput {
	return i.ToPrivateEndpointConnectionArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionArray) ToPrivateEndpointConnectionArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionArrayOutput)
}

type PrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnection)(nil)).Elem()
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return o
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return o
}

func (o PrivateEndpointConnectionOutput) Properties() PrivateEndpointConnectionPropertiesOutput {
	return o.ApplyT(func(v PrivateEndpointConnection) PrivateEndpointConnectionProperties { return v.Properties }).(PrivateEndpointConnectionPropertiesOutput)
}

type PrivateEndpointConnectionArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnection)(nil)).Elem()
}

func (o PrivateEndpointConnectionArrayOutput) ToPrivateEndpointConnectionArrayOutput() PrivateEndpointConnectionArrayOutput {
	return o
}

func (o PrivateEndpointConnectionArrayOutput) ToPrivateEndpointConnectionArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionArrayOutput {
	return o
}

func (o PrivateEndpointConnectionArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnection {
		return vs[0].([]PrivateEndpointConnection)[vs[1].(int)]
	}).(PrivateEndpointConnectionOutput)
}

type PrivateEndpointConnectionProperties struct {
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionPropertiesInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput
	ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesOutput
}

type PrivateEndpointConnectionPropertiesArgs struct {
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return i.ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesOutput)
}

type PrivateEndpointConnectionPropertiesOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateEndpointConnectionPropertiesResponse struct {
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
}

type PrivateEndpointConnectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id         string                                      `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	Type       string                                      `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Properties() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     string  `pulumi:"description"`
	Status          string  `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringInput    `pulumi:"description"`
	Status          pulumi.StringInput    `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     string  `pulumi:"description"`
	Status          string  `pulumi:"status"`
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) string { return v.Status }).(pulumi.StringOutput)
}

type SharedAccessSignatureAuthorizationRuleAccessRightsDescription struct {
	KeyName      string  `pulumi:"keyName"`
	PrimaryKey   *string `pulumi:"primaryKey"`
	Rights       string  `pulumi:"rights"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}





type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionInput interface {
	pulumi.Input

	ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput
	ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutputWithContext(context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput
}

type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArgs struct {
	KeyName      pulumi.StringInput    `pulumi:"keyName"`
	PrimaryKey   pulumi.StringPtrInput `pulumi:"primaryKey"`
	Rights       pulumi.StringInput    `pulumi:"rights"`
	SecondaryKey pulumi.StringPtrInput `pulumi:"secondaryKey"`
}

func (SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedAccessSignatureAuthorizationRuleAccessRightsDescription)(nil)).Elem()
}

func (i SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArgs) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput {
	return i.ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutputWithContext(context.Background())
}

func (i SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArgs) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput)
}





type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayInput interface {
	pulumi.Input

	ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput
	ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutputWithContext(context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput
}

type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArray []SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionInput

func (SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedAccessSignatureAuthorizationRuleAccessRightsDescription)(nil)).Elem()
}

func (i SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArray) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput {
	return i.ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutputWithContext(context.Background())
}

func (i SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArray) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput)
}

type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput struct{ *pulumi.OutputState }

func (SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedAccessSignatureAuthorizationRuleAccessRightsDescription)(nil)).Elem()
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleAccessRightsDescription) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleAccessRightsDescription) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput) Rights() pulumi.StringOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleAccessRightsDescription) string { return v.Rights }).(pulumi.StringOutput)
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleAccessRightsDescription) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput struct{ *pulumi.OutputState }

func (SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedAccessSignatureAuthorizationRuleAccessRightsDescription)(nil)).Elem()
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput) Index(i pulumi.IntInput) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharedAccessSignatureAuthorizationRuleAccessRightsDescription {
		return vs[0].([]SharedAccessSignatureAuthorizationRuleAccessRightsDescription)[vs[1].(int)]
	}).(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput)
}

type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse struct {
	KeyName      string  `pulumi:"keyName"`
	PrimaryKey   *string `pulumi:"primaryKey"`
	Rights       string  `pulumi:"rights"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}

type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput struct{ *pulumi.OutputState }

func (SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse)(nil)).Elem()
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse) *string {
		return v.PrimaryKey
	}).(pulumi.StringPtrOutput)
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput) Rights() pulumi.StringOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse) string { return v.Rights }).(pulumi.StringOutput)
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse) *string {
		return v.SecondaryKey
	}).(pulumi.StringPtrOutput)
}

type SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse)(nil)).Elem()
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput() SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput) ToSharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput) Index(i pulumi.IntInput) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse {
		return vs[0].([]SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse)[vs[1].(int)]
	}).(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput)
}

type TargetIpFilterRule struct {
	Action     IpFilterActionType  `pulumi:"action"`
	FilterName string              `pulumi:"filterName"`
	IpMask     string              `pulumi:"ipMask"`
	Target     *IpFilterTargetType `pulumi:"target"`
}





type TargetIpFilterRuleInput interface {
	pulumi.Input

	ToTargetIpFilterRuleOutput() TargetIpFilterRuleOutput
	ToTargetIpFilterRuleOutputWithContext(context.Context) TargetIpFilterRuleOutput
}

type TargetIpFilterRuleArgs struct {
	Action     IpFilterActionTypeInput    `pulumi:"action"`
	FilterName pulumi.StringInput         `pulumi:"filterName"`
	IpMask     pulumi.StringInput         `pulumi:"ipMask"`
	Target     IpFilterTargetTypePtrInput `pulumi:"target"`
}

func (TargetIpFilterRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetIpFilterRule)(nil)).Elem()
}

func (i TargetIpFilterRuleArgs) ToTargetIpFilterRuleOutput() TargetIpFilterRuleOutput {
	return i.ToTargetIpFilterRuleOutputWithContext(context.Background())
}

func (i TargetIpFilterRuleArgs) ToTargetIpFilterRuleOutputWithContext(ctx context.Context) TargetIpFilterRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetIpFilterRuleOutput)
}





type TargetIpFilterRuleArrayInput interface {
	pulumi.Input

	ToTargetIpFilterRuleArrayOutput() TargetIpFilterRuleArrayOutput
	ToTargetIpFilterRuleArrayOutputWithContext(context.Context) TargetIpFilterRuleArrayOutput
}

type TargetIpFilterRuleArray []TargetIpFilterRuleInput

func (TargetIpFilterRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetIpFilterRule)(nil)).Elem()
}

func (i TargetIpFilterRuleArray) ToTargetIpFilterRuleArrayOutput() TargetIpFilterRuleArrayOutput {
	return i.ToTargetIpFilterRuleArrayOutputWithContext(context.Background())
}

func (i TargetIpFilterRuleArray) ToTargetIpFilterRuleArrayOutputWithContext(ctx context.Context) TargetIpFilterRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetIpFilterRuleArrayOutput)
}

type TargetIpFilterRuleOutput struct{ *pulumi.OutputState }

func (TargetIpFilterRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetIpFilterRule)(nil)).Elem()
}

func (o TargetIpFilterRuleOutput) ToTargetIpFilterRuleOutput() TargetIpFilterRuleOutput {
	return o
}

func (o TargetIpFilterRuleOutput) ToTargetIpFilterRuleOutputWithContext(ctx context.Context) TargetIpFilterRuleOutput {
	return o
}

func (o TargetIpFilterRuleOutput) Action() IpFilterActionTypeOutput {
	return o.ApplyT(func(v TargetIpFilterRule) IpFilterActionType { return v.Action }).(IpFilterActionTypeOutput)
}

func (o TargetIpFilterRuleOutput) FilterName() pulumi.StringOutput {
	return o.ApplyT(func(v TargetIpFilterRule) string { return v.FilterName }).(pulumi.StringOutput)
}

func (o TargetIpFilterRuleOutput) IpMask() pulumi.StringOutput {
	return o.ApplyT(func(v TargetIpFilterRule) string { return v.IpMask }).(pulumi.StringOutput)
}

func (o TargetIpFilterRuleOutput) Target() IpFilterTargetTypePtrOutput {
	return o.ApplyT(func(v TargetIpFilterRule) *IpFilterTargetType { return v.Target }).(IpFilterTargetTypePtrOutput)
}

type TargetIpFilterRuleArrayOutput struct{ *pulumi.OutputState }

func (TargetIpFilterRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetIpFilterRule)(nil)).Elem()
}

func (o TargetIpFilterRuleArrayOutput) ToTargetIpFilterRuleArrayOutput() TargetIpFilterRuleArrayOutput {
	return o
}

func (o TargetIpFilterRuleArrayOutput) ToTargetIpFilterRuleArrayOutputWithContext(ctx context.Context) TargetIpFilterRuleArrayOutput {
	return o
}

func (o TargetIpFilterRuleArrayOutput) Index(i pulumi.IntInput) TargetIpFilterRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetIpFilterRule {
		return vs[0].([]TargetIpFilterRule)[vs[1].(int)]
	}).(TargetIpFilterRuleOutput)
}

type TargetIpFilterRuleResponse struct {
	Action     string  `pulumi:"action"`
	FilterName string  `pulumi:"filterName"`
	IpMask     string  `pulumi:"ipMask"`
	Target     *string `pulumi:"target"`
}

type TargetIpFilterRuleResponseOutput struct{ *pulumi.OutputState }

func (TargetIpFilterRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetIpFilterRuleResponse)(nil)).Elem()
}

func (o TargetIpFilterRuleResponseOutput) ToTargetIpFilterRuleResponseOutput() TargetIpFilterRuleResponseOutput {
	return o
}

func (o TargetIpFilterRuleResponseOutput) ToTargetIpFilterRuleResponseOutputWithContext(ctx context.Context) TargetIpFilterRuleResponseOutput {
	return o
}

func (o TargetIpFilterRuleResponseOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v TargetIpFilterRuleResponse) string { return v.Action }).(pulumi.StringOutput)
}

func (o TargetIpFilterRuleResponseOutput) FilterName() pulumi.StringOutput {
	return o.ApplyT(func(v TargetIpFilterRuleResponse) string { return v.FilterName }).(pulumi.StringOutput)
}

func (o TargetIpFilterRuleResponseOutput) IpMask() pulumi.StringOutput {
	return o.ApplyT(func(v TargetIpFilterRuleResponse) string { return v.IpMask }).(pulumi.StringOutput)
}

func (o TargetIpFilterRuleResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetIpFilterRuleResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type TargetIpFilterRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (TargetIpFilterRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetIpFilterRuleResponse)(nil)).Elem()
}

func (o TargetIpFilterRuleResponseArrayOutput) ToTargetIpFilterRuleResponseArrayOutput() TargetIpFilterRuleResponseArrayOutput {
	return o
}

func (o TargetIpFilterRuleResponseArrayOutput) ToTargetIpFilterRuleResponseArrayOutputWithContext(ctx context.Context) TargetIpFilterRuleResponseArrayOutput {
	return o
}

func (o TargetIpFilterRuleResponseArrayOutput) Index(i pulumi.IntInput) TargetIpFilterRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetIpFilterRuleResponse {
		return vs[0].([]TargetIpFilterRuleResponse)[vs[1].(int)]
	}).(TargetIpFilterRuleResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ArmIdentityOutput{})
	pulumi.RegisterOutputType(ArmIdentityPtrOutput{})
	pulumi.RegisterOutputType(ArmIdentityResponseOutput{})
	pulumi.RegisterOutputType(ArmIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ArmUserIdentityResponseOutput{})
	pulumi.RegisterOutputType(ArmUserIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(CertificatePropertiesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionKeyIdentityOutput{})
	pulumi.RegisterOutputType(EncryptionKeyIdentityPtrOutput{})
	pulumi.RegisterOutputType(EncryptionKeyIdentityResponseOutput{})
	pulumi.RegisterOutputType(EncryptionKeyIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesDescriptionOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesDescriptionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesDescriptionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(IotDpsPropertiesDescriptionOutput{})
	pulumi.RegisterOutputType(IotDpsPropertiesDescriptionResponseOutput{})
	pulumi.RegisterOutputType(IotDpsSkuInfoOutput{})
	pulumi.RegisterOutputType(IotDpsSkuInfoResponseOutput{})
	pulumi.RegisterOutputType(IotHubDefinitionDescriptionOutput{})
	pulumi.RegisterOutputType(IotHubDefinitionDescriptionArrayOutput{})
	pulumi.RegisterOutputType(IotHubDefinitionDescriptionResponseOutput{})
	pulumi.RegisterOutputType(IotHubDefinitionDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyPropertiesArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultKeyPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArrayOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(TargetIpFilterRuleOutput{})
	pulumi.RegisterOutputType(TargetIpFilterRuleArrayOutput{})
	pulumi.RegisterOutputType(TargetIpFilterRuleResponseOutput{})
	pulumi.RegisterOutputType(TargetIpFilterRuleResponseArrayOutput{})
}
