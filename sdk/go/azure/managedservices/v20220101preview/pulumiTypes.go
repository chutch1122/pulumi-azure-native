


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Authorization struct {
	DelegatedRoleDefinitionIds []string `pulumi:"delegatedRoleDefinitionIds"`
	PrincipalId                string   `pulumi:"principalId"`
	PrincipalIdDisplayName     *string  `pulumi:"principalIdDisplayName"`
	RoleDefinitionId           string   `pulumi:"roleDefinitionId"`
}





type AuthorizationInput interface {
	pulumi.Input

	ToAuthorizationOutput() AuthorizationOutput
	ToAuthorizationOutputWithContext(context.Context) AuthorizationOutput
}

type AuthorizationArgs struct {
	DelegatedRoleDefinitionIds pulumi.StringArrayInput `pulumi:"delegatedRoleDefinitionIds"`
	PrincipalId                pulumi.StringInput      `pulumi:"principalId"`
	PrincipalIdDisplayName     pulumi.StringPtrInput   `pulumi:"principalIdDisplayName"`
	RoleDefinitionId           pulumi.StringInput      `pulumi:"roleDefinitionId"`
}

func (AuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Authorization)(nil)).Elem()
}

func (i AuthorizationArgs) ToAuthorizationOutput() AuthorizationOutput {
	return i.ToAuthorizationOutputWithContext(context.Background())
}

func (i AuthorizationArgs) ToAuthorizationOutputWithContext(ctx context.Context) AuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationOutput)
}





type AuthorizationArrayInput interface {
	pulumi.Input

	ToAuthorizationArrayOutput() AuthorizationArrayOutput
	ToAuthorizationArrayOutputWithContext(context.Context) AuthorizationArrayOutput
}

type AuthorizationArray []AuthorizationInput

func (AuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Authorization)(nil)).Elem()
}

func (i AuthorizationArray) ToAuthorizationArrayOutput() AuthorizationArrayOutput {
	return i.ToAuthorizationArrayOutputWithContext(context.Background())
}

func (i AuthorizationArray) ToAuthorizationArrayOutputWithContext(ctx context.Context) AuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationArrayOutput)
}

type AuthorizationOutput struct{ *pulumi.OutputState }

func (AuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Authorization)(nil)).Elem()
}

func (o AuthorizationOutput) ToAuthorizationOutput() AuthorizationOutput {
	return o
}

func (o AuthorizationOutput) ToAuthorizationOutputWithContext(ctx context.Context) AuthorizationOutput {
	return o
}

func (o AuthorizationOutput) DelegatedRoleDefinitionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Authorization) []string { return v.DelegatedRoleDefinitionIds }).(pulumi.StringArrayOutput)
}

func (o AuthorizationOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v Authorization) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o AuthorizationOutput) PrincipalIdDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Authorization) *string { return v.PrincipalIdDisplayName }).(pulumi.StringPtrOutput)
}

func (o AuthorizationOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v Authorization) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type AuthorizationArrayOutput struct{ *pulumi.OutputState }

func (AuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Authorization)(nil)).Elem()
}

func (o AuthorizationArrayOutput) ToAuthorizationArrayOutput() AuthorizationArrayOutput {
	return o
}

func (o AuthorizationArrayOutput) ToAuthorizationArrayOutputWithContext(ctx context.Context) AuthorizationArrayOutput {
	return o
}

func (o AuthorizationArrayOutput) Index(i pulumi.IntInput) AuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Authorization {
		return vs[0].([]Authorization)[vs[1].(int)]
	}).(AuthorizationOutput)
}

type AuthorizationResponse struct {
	DelegatedRoleDefinitionIds []string `pulumi:"delegatedRoleDefinitionIds"`
	PrincipalId                string   `pulumi:"principalId"`
	PrincipalIdDisplayName     *string  `pulumi:"principalIdDisplayName"`
	RoleDefinitionId           string   `pulumi:"roleDefinitionId"`
}

type AuthorizationResponseOutput struct{ *pulumi.OutputState }

func (AuthorizationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthorizationResponse)(nil)).Elem()
}

func (o AuthorizationResponseOutput) ToAuthorizationResponseOutput() AuthorizationResponseOutput {
	return o
}

func (o AuthorizationResponseOutput) ToAuthorizationResponseOutputWithContext(ctx context.Context) AuthorizationResponseOutput {
	return o
}

func (o AuthorizationResponseOutput) DelegatedRoleDefinitionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AuthorizationResponse) []string { return v.DelegatedRoleDefinitionIds }).(pulumi.StringArrayOutput)
}

func (o AuthorizationResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v AuthorizationResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o AuthorizationResponseOutput) PrincipalIdDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AuthorizationResponse) *string { return v.PrincipalIdDisplayName }).(pulumi.StringPtrOutput)
}

func (o AuthorizationResponseOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v AuthorizationResponse) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type AuthorizationResponseArrayOutput struct{ *pulumi.OutputState }

func (AuthorizationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthorizationResponse)(nil)).Elem()
}

func (o AuthorizationResponseArrayOutput) ToAuthorizationResponseArrayOutput() AuthorizationResponseArrayOutput {
	return o
}

func (o AuthorizationResponseArrayOutput) ToAuthorizationResponseArrayOutputWithContext(ctx context.Context) AuthorizationResponseArrayOutput {
	return o
}

func (o AuthorizationResponseArrayOutput) Index(i pulumi.IntInput) AuthorizationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AuthorizationResponse {
		return vs[0].([]AuthorizationResponse)[vs[1].(int)]
	}).(AuthorizationResponseOutput)
}

type EligibleApprover struct {
	PrincipalId            string  `pulumi:"principalId"`
	PrincipalIdDisplayName *string `pulumi:"principalIdDisplayName"`
}





type EligibleApproverInput interface {
	pulumi.Input

	ToEligibleApproverOutput() EligibleApproverOutput
	ToEligibleApproverOutputWithContext(context.Context) EligibleApproverOutput
}

type EligibleApproverArgs struct {
	PrincipalId            pulumi.StringInput    `pulumi:"principalId"`
	PrincipalIdDisplayName pulumi.StringPtrInput `pulumi:"principalIdDisplayName"`
}

func (EligibleApproverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EligibleApprover)(nil)).Elem()
}

func (i EligibleApproverArgs) ToEligibleApproverOutput() EligibleApproverOutput {
	return i.ToEligibleApproverOutputWithContext(context.Background())
}

func (i EligibleApproverArgs) ToEligibleApproverOutputWithContext(ctx context.Context) EligibleApproverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EligibleApproverOutput)
}





type EligibleApproverArrayInput interface {
	pulumi.Input

	ToEligibleApproverArrayOutput() EligibleApproverArrayOutput
	ToEligibleApproverArrayOutputWithContext(context.Context) EligibleApproverArrayOutput
}

type EligibleApproverArray []EligibleApproverInput

func (EligibleApproverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EligibleApprover)(nil)).Elem()
}

func (i EligibleApproverArray) ToEligibleApproverArrayOutput() EligibleApproverArrayOutput {
	return i.ToEligibleApproverArrayOutputWithContext(context.Background())
}

func (i EligibleApproverArray) ToEligibleApproverArrayOutputWithContext(ctx context.Context) EligibleApproverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EligibleApproverArrayOutput)
}

type EligibleApproverOutput struct{ *pulumi.OutputState }

func (EligibleApproverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EligibleApprover)(nil)).Elem()
}

func (o EligibleApproverOutput) ToEligibleApproverOutput() EligibleApproverOutput {
	return o
}

func (o EligibleApproverOutput) ToEligibleApproverOutputWithContext(ctx context.Context) EligibleApproverOutput {
	return o
}

func (o EligibleApproverOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v EligibleApprover) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o EligibleApproverOutput) PrincipalIdDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EligibleApprover) *string { return v.PrincipalIdDisplayName }).(pulumi.StringPtrOutput)
}

type EligibleApproverArrayOutput struct{ *pulumi.OutputState }

func (EligibleApproverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EligibleApprover)(nil)).Elem()
}

func (o EligibleApproverArrayOutput) ToEligibleApproverArrayOutput() EligibleApproverArrayOutput {
	return o
}

func (o EligibleApproverArrayOutput) ToEligibleApproverArrayOutputWithContext(ctx context.Context) EligibleApproverArrayOutput {
	return o
}

func (o EligibleApproverArrayOutput) Index(i pulumi.IntInput) EligibleApproverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EligibleApprover {
		return vs[0].([]EligibleApprover)[vs[1].(int)]
	}).(EligibleApproverOutput)
}

type EligibleApproverResponse struct {
	PrincipalId            string  `pulumi:"principalId"`
	PrincipalIdDisplayName *string `pulumi:"principalIdDisplayName"`
}

type EligibleApproverResponseOutput struct{ *pulumi.OutputState }

func (EligibleApproverResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EligibleApproverResponse)(nil)).Elem()
}

func (o EligibleApproverResponseOutput) ToEligibleApproverResponseOutput() EligibleApproverResponseOutput {
	return o
}

func (o EligibleApproverResponseOutput) ToEligibleApproverResponseOutputWithContext(ctx context.Context) EligibleApproverResponseOutput {
	return o
}

func (o EligibleApproverResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v EligibleApproverResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o EligibleApproverResponseOutput) PrincipalIdDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EligibleApproverResponse) *string { return v.PrincipalIdDisplayName }).(pulumi.StringPtrOutput)
}

type EligibleApproverResponseArrayOutput struct{ *pulumi.OutputState }

func (EligibleApproverResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EligibleApproverResponse)(nil)).Elem()
}

func (o EligibleApproverResponseArrayOutput) ToEligibleApproverResponseArrayOutput() EligibleApproverResponseArrayOutput {
	return o
}

func (o EligibleApproverResponseArrayOutput) ToEligibleApproverResponseArrayOutputWithContext(ctx context.Context) EligibleApproverResponseArrayOutput {
	return o
}

func (o EligibleApproverResponseArrayOutput) Index(i pulumi.IntInput) EligibleApproverResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EligibleApproverResponse {
		return vs[0].([]EligibleApproverResponse)[vs[1].(int)]
	}).(EligibleApproverResponseOutput)
}

type EligibleAuthorization struct {
	JustInTimeAccessPolicy *JustInTimeAccessPolicy `pulumi:"justInTimeAccessPolicy"`
	PrincipalId            string                  `pulumi:"principalId"`
	PrincipalIdDisplayName *string                 `pulumi:"principalIdDisplayName"`
	RoleDefinitionId       string                  `pulumi:"roleDefinitionId"`
}


func (val *EligibleAuthorization) Defaults() *EligibleAuthorization {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.JustInTimeAccessPolicy = tmp.JustInTimeAccessPolicy.Defaults()

	return &tmp
}





type EligibleAuthorizationInput interface {
	pulumi.Input

	ToEligibleAuthorizationOutput() EligibleAuthorizationOutput
	ToEligibleAuthorizationOutputWithContext(context.Context) EligibleAuthorizationOutput
}

type EligibleAuthorizationArgs struct {
	JustInTimeAccessPolicy JustInTimeAccessPolicyPtrInput `pulumi:"justInTimeAccessPolicy"`
	PrincipalId            pulumi.StringInput             `pulumi:"principalId"`
	PrincipalIdDisplayName pulumi.StringPtrInput          `pulumi:"principalIdDisplayName"`
	RoleDefinitionId       pulumi.StringInput             `pulumi:"roleDefinitionId"`
}


func (val *EligibleAuthorizationArgs) Defaults() *EligibleAuthorizationArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	return &tmp
}
func (EligibleAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EligibleAuthorization)(nil)).Elem()
}

func (i EligibleAuthorizationArgs) ToEligibleAuthorizationOutput() EligibleAuthorizationOutput {
	return i.ToEligibleAuthorizationOutputWithContext(context.Background())
}

func (i EligibleAuthorizationArgs) ToEligibleAuthorizationOutputWithContext(ctx context.Context) EligibleAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EligibleAuthorizationOutput)
}





type EligibleAuthorizationArrayInput interface {
	pulumi.Input

	ToEligibleAuthorizationArrayOutput() EligibleAuthorizationArrayOutput
	ToEligibleAuthorizationArrayOutputWithContext(context.Context) EligibleAuthorizationArrayOutput
}

type EligibleAuthorizationArray []EligibleAuthorizationInput

func (EligibleAuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EligibleAuthorization)(nil)).Elem()
}

func (i EligibleAuthorizationArray) ToEligibleAuthorizationArrayOutput() EligibleAuthorizationArrayOutput {
	return i.ToEligibleAuthorizationArrayOutputWithContext(context.Background())
}

func (i EligibleAuthorizationArray) ToEligibleAuthorizationArrayOutputWithContext(ctx context.Context) EligibleAuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EligibleAuthorizationArrayOutput)
}

type EligibleAuthorizationOutput struct{ *pulumi.OutputState }

func (EligibleAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EligibleAuthorization)(nil)).Elem()
}

func (o EligibleAuthorizationOutput) ToEligibleAuthorizationOutput() EligibleAuthorizationOutput {
	return o
}

func (o EligibleAuthorizationOutput) ToEligibleAuthorizationOutputWithContext(ctx context.Context) EligibleAuthorizationOutput {
	return o
}

func (o EligibleAuthorizationOutput) JustInTimeAccessPolicy() JustInTimeAccessPolicyPtrOutput {
	return o.ApplyT(func(v EligibleAuthorization) *JustInTimeAccessPolicy { return v.JustInTimeAccessPolicy }).(JustInTimeAccessPolicyPtrOutput)
}

func (o EligibleAuthorizationOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v EligibleAuthorization) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o EligibleAuthorizationOutput) PrincipalIdDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EligibleAuthorization) *string { return v.PrincipalIdDisplayName }).(pulumi.StringPtrOutput)
}

func (o EligibleAuthorizationOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v EligibleAuthorization) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type EligibleAuthorizationArrayOutput struct{ *pulumi.OutputState }

func (EligibleAuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EligibleAuthorization)(nil)).Elem()
}

func (o EligibleAuthorizationArrayOutput) ToEligibleAuthorizationArrayOutput() EligibleAuthorizationArrayOutput {
	return o
}

func (o EligibleAuthorizationArrayOutput) ToEligibleAuthorizationArrayOutputWithContext(ctx context.Context) EligibleAuthorizationArrayOutput {
	return o
}

func (o EligibleAuthorizationArrayOutput) Index(i pulumi.IntInput) EligibleAuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EligibleAuthorization {
		return vs[0].([]EligibleAuthorization)[vs[1].(int)]
	}).(EligibleAuthorizationOutput)
}

type EligibleAuthorizationResponse struct {
	JustInTimeAccessPolicy *JustInTimeAccessPolicyResponse `pulumi:"justInTimeAccessPolicy"`
	PrincipalId            string                          `pulumi:"principalId"`
	PrincipalIdDisplayName *string                         `pulumi:"principalIdDisplayName"`
	RoleDefinitionId       string                          `pulumi:"roleDefinitionId"`
}


func (val *EligibleAuthorizationResponse) Defaults() *EligibleAuthorizationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.JustInTimeAccessPolicy = tmp.JustInTimeAccessPolicy.Defaults()

	return &tmp
}

type EligibleAuthorizationResponseOutput struct{ *pulumi.OutputState }

func (EligibleAuthorizationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EligibleAuthorizationResponse)(nil)).Elem()
}

func (o EligibleAuthorizationResponseOutput) ToEligibleAuthorizationResponseOutput() EligibleAuthorizationResponseOutput {
	return o
}

func (o EligibleAuthorizationResponseOutput) ToEligibleAuthorizationResponseOutputWithContext(ctx context.Context) EligibleAuthorizationResponseOutput {
	return o
}

func (o EligibleAuthorizationResponseOutput) JustInTimeAccessPolicy() JustInTimeAccessPolicyResponsePtrOutput {
	return o.ApplyT(func(v EligibleAuthorizationResponse) *JustInTimeAccessPolicyResponse { return v.JustInTimeAccessPolicy }).(JustInTimeAccessPolicyResponsePtrOutput)
}

func (o EligibleAuthorizationResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v EligibleAuthorizationResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o EligibleAuthorizationResponseOutput) PrincipalIdDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EligibleAuthorizationResponse) *string { return v.PrincipalIdDisplayName }).(pulumi.StringPtrOutput)
}

func (o EligibleAuthorizationResponseOutput) RoleDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v EligibleAuthorizationResponse) string { return v.RoleDefinitionId }).(pulumi.StringOutput)
}

type EligibleAuthorizationResponseArrayOutput struct{ *pulumi.OutputState }

func (EligibleAuthorizationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EligibleAuthorizationResponse)(nil)).Elem()
}

func (o EligibleAuthorizationResponseArrayOutput) ToEligibleAuthorizationResponseArrayOutput() EligibleAuthorizationResponseArrayOutput {
	return o
}

func (o EligibleAuthorizationResponseArrayOutput) ToEligibleAuthorizationResponseArrayOutputWithContext(ctx context.Context) EligibleAuthorizationResponseArrayOutput {
	return o
}

func (o EligibleAuthorizationResponseArrayOutput) Index(i pulumi.IntInput) EligibleAuthorizationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EligibleAuthorizationResponse {
		return vs[0].([]EligibleAuthorizationResponse)[vs[1].(int)]
	}).(EligibleAuthorizationResponseOutput)
}

type JustInTimeAccessPolicy struct {
	ManagedByTenantApprovers  []EligibleApprover `pulumi:"managedByTenantApprovers"`
	MaximumActivationDuration *string            `pulumi:"maximumActivationDuration"`
	MultiFactorAuthProvider   string             `pulumi:"multiFactorAuthProvider"`
}


func (val *JustInTimeAccessPolicy) Defaults() *JustInTimeAccessPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaximumActivationDuration) {
		maximumActivationDuration_ := "PT8H"
		tmp.MaximumActivationDuration = &maximumActivationDuration_
	}
	if isZero(tmp.MultiFactorAuthProvider) {
		tmp.MultiFactorAuthProvider = "None"
	}
	return &tmp
}





type JustInTimeAccessPolicyInput interface {
	pulumi.Input

	ToJustInTimeAccessPolicyOutput() JustInTimeAccessPolicyOutput
	ToJustInTimeAccessPolicyOutputWithContext(context.Context) JustInTimeAccessPolicyOutput
}

type JustInTimeAccessPolicyArgs struct {
	ManagedByTenantApprovers  EligibleApproverArrayInput `pulumi:"managedByTenantApprovers"`
	MaximumActivationDuration pulumi.StringPtrInput      `pulumi:"maximumActivationDuration"`
	MultiFactorAuthProvider   pulumi.StringInput         `pulumi:"multiFactorAuthProvider"`
}


func (val *JustInTimeAccessPolicyArgs) Defaults() *JustInTimeAccessPolicyArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaximumActivationDuration) {
		tmp.MaximumActivationDuration = pulumi.StringPtr("PT8H")
	}
	if isZero(tmp.MultiFactorAuthProvider) {
		tmp.MultiFactorAuthProvider = pulumi.String("None")
	}
	return &tmp
}
func (JustInTimeAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JustInTimeAccessPolicy)(nil)).Elem()
}

func (i JustInTimeAccessPolicyArgs) ToJustInTimeAccessPolicyOutput() JustInTimeAccessPolicyOutput {
	return i.ToJustInTimeAccessPolicyOutputWithContext(context.Background())
}

func (i JustInTimeAccessPolicyArgs) ToJustInTimeAccessPolicyOutputWithContext(ctx context.Context) JustInTimeAccessPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JustInTimeAccessPolicyOutput)
}

func (i JustInTimeAccessPolicyArgs) ToJustInTimeAccessPolicyPtrOutput() JustInTimeAccessPolicyPtrOutput {
	return i.ToJustInTimeAccessPolicyPtrOutputWithContext(context.Background())
}

func (i JustInTimeAccessPolicyArgs) ToJustInTimeAccessPolicyPtrOutputWithContext(ctx context.Context) JustInTimeAccessPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JustInTimeAccessPolicyOutput).ToJustInTimeAccessPolicyPtrOutputWithContext(ctx)
}









type JustInTimeAccessPolicyPtrInput interface {
	pulumi.Input

	ToJustInTimeAccessPolicyPtrOutput() JustInTimeAccessPolicyPtrOutput
	ToJustInTimeAccessPolicyPtrOutputWithContext(context.Context) JustInTimeAccessPolicyPtrOutput
}

type justInTimeAccessPolicyPtrType JustInTimeAccessPolicyArgs

func JustInTimeAccessPolicyPtr(v *JustInTimeAccessPolicyArgs) JustInTimeAccessPolicyPtrInput {
	return (*justInTimeAccessPolicyPtrType)(v)
}

func (*justInTimeAccessPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JustInTimeAccessPolicy)(nil)).Elem()
}

func (i *justInTimeAccessPolicyPtrType) ToJustInTimeAccessPolicyPtrOutput() JustInTimeAccessPolicyPtrOutput {
	return i.ToJustInTimeAccessPolicyPtrOutputWithContext(context.Background())
}

func (i *justInTimeAccessPolicyPtrType) ToJustInTimeAccessPolicyPtrOutputWithContext(ctx context.Context) JustInTimeAccessPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JustInTimeAccessPolicyPtrOutput)
}

type JustInTimeAccessPolicyOutput struct{ *pulumi.OutputState }

func (JustInTimeAccessPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JustInTimeAccessPolicy)(nil)).Elem()
}

func (o JustInTimeAccessPolicyOutput) ToJustInTimeAccessPolicyOutput() JustInTimeAccessPolicyOutput {
	return o
}

func (o JustInTimeAccessPolicyOutput) ToJustInTimeAccessPolicyOutputWithContext(ctx context.Context) JustInTimeAccessPolicyOutput {
	return o
}

func (o JustInTimeAccessPolicyOutput) ToJustInTimeAccessPolicyPtrOutput() JustInTimeAccessPolicyPtrOutput {
	return o.ToJustInTimeAccessPolicyPtrOutputWithContext(context.Background())
}

func (o JustInTimeAccessPolicyOutput) ToJustInTimeAccessPolicyPtrOutputWithContext(ctx context.Context) JustInTimeAccessPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JustInTimeAccessPolicy) *JustInTimeAccessPolicy {
		return &v
	}).(JustInTimeAccessPolicyPtrOutput)
}

func (o JustInTimeAccessPolicyOutput) ManagedByTenantApprovers() EligibleApproverArrayOutput {
	return o.ApplyT(func(v JustInTimeAccessPolicy) []EligibleApprover { return v.ManagedByTenantApprovers }).(EligibleApproverArrayOutput)
}

func (o JustInTimeAccessPolicyOutput) MaximumActivationDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JustInTimeAccessPolicy) *string { return v.MaximumActivationDuration }).(pulumi.StringPtrOutput)
}

func (o JustInTimeAccessPolicyOutput) MultiFactorAuthProvider() pulumi.StringOutput {
	return o.ApplyT(func(v JustInTimeAccessPolicy) string { return v.MultiFactorAuthProvider }).(pulumi.StringOutput)
}

type JustInTimeAccessPolicyPtrOutput struct{ *pulumi.OutputState }

func (JustInTimeAccessPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JustInTimeAccessPolicy)(nil)).Elem()
}

func (o JustInTimeAccessPolicyPtrOutput) ToJustInTimeAccessPolicyPtrOutput() JustInTimeAccessPolicyPtrOutput {
	return o
}

func (o JustInTimeAccessPolicyPtrOutput) ToJustInTimeAccessPolicyPtrOutputWithContext(ctx context.Context) JustInTimeAccessPolicyPtrOutput {
	return o
}

func (o JustInTimeAccessPolicyPtrOutput) Elem() JustInTimeAccessPolicyOutput {
	return o.ApplyT(func(v *JustInTimeAccessPolicy) JustInTimeAccessPolicy {
		if v != nil {
			return *v
		}
		var ret JustInTimeAccessPolicy
		return ret
	}).(JustInTimeAccessPolicyOutput)
}

func (o JustInTimeAccessPolicyPtrOutput) ManagedByTenantApprovers() EligibleApproverArrayOutput {
	return o.ApplyT(func(v *JustInTimeAccessPolicy) []EligibleApprover {
		if v == nil {
			return nil
		}
		return v.ManagedByTenantApprovers
	}).(EligibleApproverArrayOutput)
}

func (o JustInTimeAccessPolicyPtrOutput) MaximumActivationDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JustInTimeAccessPolicy) *string {
		if v == nil {
			return nil
		}
		return v.MaximumActivationDuration
	}).(pulumi.StringPtrOutput)
}

func (o JustInTimeAccessPolicyPtrOutput) MultiFactorAuthProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JustInTimeAccessPolicy) *string {
		if v == nil {
			return nil
		}
		return &v.MultiFactorAuthProvider
	}).(pulumi.StringPtrOutput)
}

type JustInTimeAccessPolicyResponse struct {
	ManagedByTenantApprovers  []EligibleApproverResponse `pulumi:"managedByTenantApprovers"`
	MaximumActivationDuration *string                    `pulumi:"maximumActivationDuration"`
	MultiFactorAuthProvider   string                     `pulumi:"multiFactorAuthProvider"`
}


func (val *JustInTimeAccessPolicyResponse) Defaults() *JustInTimeAccessPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaximumActivationDuration) {
		maximumActivationDuration_ := "PT8H"
		tmp.MaximumActivationDuration = &maximumActivationDuration_
	}
	if isZero(tmp.MultiFactorAuthProvider) {
		tmp.MultiFactorAuthProvider = "None"
	}
	return &tmp
}

type JustInTimeAccessPolicyResponseOutput struct{ *pulumi.OutputState }

func (JustInTimeAccessPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JustInTimeAccessPolicyResponse)(nil)).Elem()
}

func (o JustInTimeAccessPolicyResponseOutput) ToJustInTimeAccessPolicyResponseOutput() JustInTimeAccessPolicyResponseOutput {
	return o
}

func (o JustInTimeAccessPolicyResponseOutput) ToJustInTimeAccessPolicyResponseOutputWithContext(ctx context.Context) JustInTimeAccessPolicyResponseOutput {
	return o
}

func (o JustInTimeAccessPolicyResponseOutput) ManagedByTenantApprovers() EligibleApproverResponseArrayOutput {
	return o.ApplyT(func(v JustInTimeAccessPolicyResponse) []EligibleApproverResponse { return v.ManagedByTenantApprovers }).(EligibleApproverResponseArrayOutput)
}

func (o JustInTimeAccessPolicyResponseOutput) MaximumActivationDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JustInTimeAccessPolicyResponse) *string { return v.MaximumActivationDuration }).(pulumi.StringPtrOutput)
}

func (o JustInTimeAccessPolicyResponseOutput) MultiFactorAuthProvider() pulumi.StringOutput {
	return o.ApplyT(func(v JustInTimeAccessPolicyResponse) string { return v.MultiFactorAuthProvider }).(pulumi.StringOutput)
}

type JustInTimeAccessPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (JustInTimeAccessPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JustInTimeAccessPolicyResponse)(nil)).Elem()
}

func (o JustInTimeAccessPolicyResponsePtrOutput) ToJustInTimeAccessPolicyResponsePtrOutput() JustInTimeAccessPolicyResponsePtrOutput {
	return o
}

func (o JustInTimeAccessPolicyResponsePtrOutput) ToJustInTimeAccessPolicyResponsePtrOutputWithContext(ctx context.Context) JustInTimeAccessPolicyResponsePtrOutput {
	return o
}

func (o JustInTimeAccessPolicyResponsePtrOutput) Elem() JustInTimeAccessPolicyResponseOutput {
	return o.ApplyT(func(v *JustInTimeAccessPolicyResponse) JustInTimeAccessPolicyResponse {
		if v != nil {
			return *v
		}
		var ret JustInTimeAccessPolicyResponse
		return ret
	}).(JustInTimeAccessPolicyResponseOutput)
}

func (o JustInTimeAccessPolicyResponsePtrOutput) ManagedByTenantApprovers() EligibleApproverResponseArrayOutput {
	return o.ApplyT(func(v *JustInTimeAccessPolicyResponse) []EligibleApproverResponse {
		if v == nil {
			return nil
		}
		return v.ManagedByTenantApprovers
	}).(EligibleApproverResponseArrayOutput)
}

func (o JustInTimeAccessPolicyResponsePtrOutput) MaximumActivationDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JustInTimeAccessPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.MaximumActivationDuration
	}).(pulumi.StringPtrOutput)
}

func (o JustInTimeAccessPolicyResponsePtrOutput) MultiFactorAuthProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JustInTimeAccessPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.MultiFactorAuthProvider
	}).(pulumi.StringPtrOutput)
}

type Plan struct {
	Name      string `pulumi:"name"`
	Product   string `pulumi:"product"`
	Publisher string `pulumi:"publisher"`
	Version   string `pulumi:"version"`
}





type PlanInput interface {
	pulumi.Input

	ToPlanOutput() PlanOutput
	ToPlanOutputWithContext(context.Context) PlanOutput
}

type PlanArgs struct {
	Name      pulumi.StringInput `pulumi:"name"`
	Product   pulumi.StringInput `pulumi:"product"`
	Publisher pulumi.StringInput `pulumi:"publisher"`
	Version   pulumi.StringInput `pulumi:"version"`
}

func (PlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (i PlanArgs) ToPlanOutput() PlanOutput {
	return i.ToPlanOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput)
}

func (i PlanArgs) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput).ToPlanPtrOutputWithContext(ctx)
}









type PlanPtrInput interface {
	pulumi.Input

	ToPlanPtrOutput() PlanPtrOutput
	ToPlanPtrOutputWithContext(context.Context) PlanPtrOutput
}

type planPtrType PlanArgs

func PlanPtr(v *PlanArgs) PlanPtrInput {
	return (*planPtrType)(v)
}

func (*planPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (i *planPtrType) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i *planPtrType) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanPtrOutput)
}

type PlanOutput struct{ *pulumi.OutputState }

func (PlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (o PlanOutput) ToPlanOutput() PlanOutput {
	return o
}

func (o PlanOutput) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return o
}

func (o PlanOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o.ToPlanPtrOutputWithContext(context.Background())
}

func (o PlanOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Plan) *Plan {
		return &v
	}).(PlanPtrOutput)
}

func (o PlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Name }).(pulumi.StringOutput)
}

func (o PlanOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Product }).(pulumi.StringOutput)
}

func (o PlanOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o PlanOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v Plan) string { return v.Version }).(pulumi.StringOutput)
}

type PlanPtrOutput struct{ *pulumi.OutputState }

func (PlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil)).Elem()
}

func (o PlanPtrOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) Elem() PlanOutput {
	return o.ApplyT(func(v *Plan) Plan {
		if v != nil {
			return *v
		}
		var ret Plan
		return ret
	}).(PlanOutput)
}

func (o PlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o PlanPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Plan) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type PlanResponse struct {
	Name      string `pulumi:"name"`
	Product   string `pulumi:"product"`
	Publisher string `pulumi:"publisher"`
	Version   string `pulumi:"version"`
}

type PlanResponseOutput struct{ *pulumi.OutputState }

func (PlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanResponse)(nil)).Elem()
}

func (o PlanResponseOutput) ToPlanResponseOutput() PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) ToPlanResponseOutputWithContext(ctx context.Context) PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Product }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.Version }).(pulumi.StringOutput)
}

type PlanResponsePtrOutput struct{ *pulumi.OutputState }

func (PlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlanResponse)(nil)).Elem()
}

func (o PlanResponsePtrOutput) ToPlanResponsePtrOutput() PlanResponsePtrOutput {
	return o
}

func (o PlanResponsePtrOutput) ToPlanResponsePtrOutputWithContext(ctx context.Context) PlanResponsePtrOutput {
	return o
}

func (o PlanResponsePtrOutput) Elem() PlanResponseOutput {
	return o.ApplyT(func(v *PlanResponse) PlanResponse {
		if v != nil {
			return *v
		}
		var ret PlanResponse
		return ret
	}).(PlanResponseOutput)
}

func (o PlanResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o PlanResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type RegistrationAssignmentProperties struct {
	RegistrationDefinitionId string `pulumi:"registrationDefinitionId"`
}





type RegistrationAssignmentPropertiesInput interface {
	pulumi.Input

	ToRegistrationAssignmentPropertiesOutput() RegistrationAssignmentPropertiesOutput
	ToRegistrationAssignmentPropertiesOutputWithContext(context.Context) RegistrationAssignmentPropertiesOutput
}

type RegistrationAssignmentPropertiesArgs struct {
	RegistrationDefinitionId pulumi.StringInput `pulumi:"registrationDefinitionId"`
}

func (RegistrationAssignmentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationAssignmentProperties)(nil)).Elem()
}

func (i RegistrationAssignmentPropertiesArgs) ToRegistrationAssignmentPropertiesOutput() RegistrationAssignmentPropertiesOutput {
	return i.ToRegistrationAssignmentPropertiesOutputWithContext(context.Background())
}

func (i RegistrationAssignmentPropertiesArgs) ToRegistrationAssignmentPropertiesOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationAssignmentPropertiesOutput)
}

func (i RegistrationAssignmentPropertiesArgs) ToRegistrationAssignmentPropertiesPtrOutput() RegistrationAssignmentPropertiesPtrOutput {
	return i.ToRegistrationAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (i RegistrationAssignmentPropertiesArgs) ToRegistrationAssignmentPropertiesPtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationAssignmentPropertiesOutput).ToRegistrationAssignmentPropertiesPtrOutputWithContext(ctx)
}









type RegistrationAssignmentPropertiesPtrInput interface {
	pulumi.Input

	ToRegistrationAssignmentPropertiesPtrOutput() RegistrationAssignmentPropertiesPtrOutput
	ToRegistrationAssignmentPropertiesPtrOutputWithContext(context.Context) RegistrationAssignmentPropertiesPtrOutput
}

type registrationAssignmentPropertiesPtrType RegistrationAssignmentPropertiesArgs

func RegistrationAssignmentPropertiesPtr(v *RegistrationAssignmentPropertiesArgs) RegistrationAssignmentPropertiesPtrInput {
	return (*registrationAssignmentPropertiesPtrType)(v)
}

func (*registrationAssignmentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationAssignmentProperties)(nil)).Elem()
}

func (i *registrationAssignmentPropertiesPtrType) ToRegistrationAssignmentPropertiesPtrOutput() RegistrationAssignmentPropertiesPtrOutput {
	return i.ToRegistrationAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (i *registrationAssignmentPropertiesPtrType) ToRegistrationAssignmentPropertiesPtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationAssignmentPropertiesPtrOutput)
}

type RegistrationAssignmentPropertiesOutput struct{ *pulumi.OutputState }

func (RegistrationAssignmentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationAssignmentProperties)(nil)).Elem()
}

func (o RegistrationAssignmentPropertiesOutput) ToRegistrationAssignmentPropertiesOutput() RegistrationAssignmentPropertiesOutput {
	return o
}

func (o RegistrationAssignmentPropertiesOutput) ToRegistrationAssignmentPropertiesOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesOutput {
	return o
}

func (o RegistrationAssignmentPropertiesOutput) ToRegistrationAssignmentPropertiesPtrOutput() RegistrationAssignmentPropertiesPtrOutput {
	return o.ToRegistrationAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (o RegistrationAssignmentPropertiesOutput) ToRegistrationAssignmentPropertiesPtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegistrationAssignmentProperties) *RegistrationAssignmentProperties {
		return &v
	}).(RegistrationAssignmentPropertiesPtrOutput)
}

func (o RegistrationAssignmentPropertiesOutput) RegistrationDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationAssignmentProperties) string { return v.RegistrationDefinitionId }).(pulumi.StringOutput)
}

type RegistrationAssignmentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (RegistrationAssignmentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationAssignmentProperties)(nil)).Elem()
}

func (o RegistrationAssignmentPropertiesPtrOutput) ToRegistrationAssignmentPropertiesPtrOutput() RegistrationAssignmentPropertiesPtrOutput {
	return o
}

func (o RegistrationAssignmentPropertiesPtrOutput) ToRegistrationAssignmentPropertiesPtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesPtrOutput {
	return o
}

func (o RegistrationAssignmentPropertiesPtrOutput) Elem() RegistrationAssignmentPropertiesOutput {
	return o.ApplyT(func(v *RegistrationAssignmentProperties) RegistrationAssignmentProperties {
		if v != nil {
			return *v
		}
		var ret RegistrationAssignmentProperties
		return ret
	}).(RegistrationAssignmentPropertiesOutput)
}

func (o RegistrationAssignmentPropertiesPtrOutput) RegistrationDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentProperties) *string {
		if v == nil {
			return nil
		}
		return &v.RegistrationDefinitionId
	}).(pulumi.StringPtrOutput)
}

type RegistrationAssignmentPropertiesResponse struct {
	ProvisioningState        string                                                         `pulumi:"provisioningState"`
	RegistrationDefinition   RegistrationAssignmentPropertiesResponseRegistrationDefinition `pulumi:"registrationDefinition"`
	RegistrationDefinitionId string                                                         `pulumi:"registrationDefinitionId"`
}

type RegistrationAssignmentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RegistrationAssignmentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationAssignmentPropertiesResponse)(nil)).Elem()
}

func (o RegistrationAssignmentPropertiesResponseOutput) ToRegistrationAssignmentPropertiesResponseOutput() RegistrationAssignmentPropertiesResponseOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponseOutput) ToRegistrationAssignmentPropertiesResponseOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponseOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RegistrationAssignmentPropertiesResponseOutput) RegistrationDefinition() RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponse) RegistrationAssignmentPropertiesResponseRegistrationDefinition {
		return v.RegistrationDefinition
	}).(RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput)
}

func (o RegistrationAssignmentPropertiesResponseOutput) RegistrationDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponse) string { return v.RegistrationDefinitionId }).(pulumi.StringOutput)
}

type RegistrationAssignmentPropertiesResponseProperties struct {
	Authorizations             []AuthorizationResponse         `pulumi:"authorizations"`
	Description                *string                         `pulumi:"description"`
	EligibleAuthorizations     []EligibleAuthorizationResponse `pulumi:"eligibleAuthorizations"`
	ManagedByTenantId          *string                         `pulumi:"managedByTenantId"`
	ManagedByTenantName        *string                         `pulumi:"managedByTenantName"`
	ManageeTenantId            *string                         `pulumi:"manageeTenantId"`
	ManageeTenantName          *string                         `pulumi:"manageeTenantName"`
	ProvisioningState          *string                         `pulumi:"provisioningState"`
	RegistrationDefinitionName *string                         `pulumi:"registrationDefinitionName"`
}

type RegistrationAssignmentPropertiesResponsePropertiesOutput struct{ *pulumi.OutputState }

func (RegistrationAssignmentPropertiesResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationAssignmentPropertiesResponseProperties)(nil)).Elem()
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) ToRegistrationAssignmentPropertiesResponsePropertiesOutput() RegistrationAssignmentPropertiesResponsePropertiesOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) ToRegistrationAssignmentPropertiesResponsePropertiesOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponsePropertiesOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) Authorizations() AuthorizationResponseArrayOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseProperties) []AuthorizationResponse {
		return v.Authorizations
	}).(AuthorizationResponseArrayOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) EligibleAuthorizations() EligibleAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseProperties) []EligibleAuthorizationResponse {
		return v.EligibleAuthorizations
	}).(EligibleAuthorizationResponseArrayOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) ManagedByTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseProperties) *string { return v.ManagedByTenantId }).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) ManagedByTenantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseProperties) *string { return v.ManagedByTenantName }).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) ManageeTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseProperties) *string { return v.ManageeTenantId }).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) ManageeTenantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseProperties) *string { return v.ManageeTenantName }).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesOutput) RegistrationDefinitionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseProperties) *string {
		return v.RegistrationDefinitionName
	}).(pulumi.StringPtrOutput)
}

type RegistrationAssignmentPropertiesResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationAssignmentPropertiesResponseProperties)(nil)).Elem()
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) ToRegistrationAssignmentPropertiesResponsePropertiesPtrOutput() RegistrationAssignmentPropertiesResponsePropertiesPtrOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) ToRegistrationAssignmentPropertiesResponsePropertiesPtrOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponsePropertiesPtrOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) Elem() RegistrationAssignmentPropertiesResponsePropertiesOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) RegistrationAssignmentPropertiesResponseProperties {
		if v != nil {
			return *v
		}
		var ret RegistrationAssignmentPropertiesResponseProperties
		return ret
	}).(RegistrationAssignmentPropertiesResponsePropertiesOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) Authorizations() AuthorizationResponseArrayOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) []AuthorizationResponse {
		if v == nil {
			return nil
		}
		return v.Authorizations
	}).(AuthorizationResponseArrayOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) EligibleAuthorizations() EligibleAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) []EligibleAuthorizationResponse {
		if v == nil {
			return nil
		}
		return v.EligibleAuthorizations
	}).(EligibleAuthorizationResponseArrayOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) ManagedByTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ManagedByTenantId
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) ManagedByTenantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ManagedByTenantName
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) ManageeTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ManageeTenantId
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) ManageeTenantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ManageeTenantName
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponsePropertiesPtrOutput) RegistrationDefinitionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationAssignmentPropertiesResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.RegistrationDefinitionName
	}).(pulumi.StringPtrOutput)
}

type RegistrationAssignmentPropertiesResponseRegistrationDefinition struct {
	Id         string                                              `pulumi:"id"`
	Name       string                                              `pulumi:"name"`
	Plan       *PlanResponse                                       `pulumi:"plan"`
	Properties *RegistrationAssignmentPropertiesResponseProperties `pulumi:"properties"`
	SystemData SystemDataResponse                                  `pulumi:"systemData"`
	Type       string                                              `pulumi:"type"`
}

type RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput struct{ *pulumi.OutputState }

func (RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationAssignmentPropertiesResponseRegistrationDefinition)(nil)).Elem()
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput() RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) ToRegistrationAssignmentPropertiesResponseRegistrationDefinitionOutputWithContext(ctx context.Context) RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput {
	return o
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseRegistrationDefinition) string { return v.Id }).(pulumi.StringOutput)
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseRegistrationDefinition) string { return v.Name }).(pulumi.StringOutput)
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseRegistrationDefinition) *PlanResponse { return v.Plan }).(PlanResponsePtrOutput)
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) Properties() RegistrationAssignmentPropertiesResponsePropertiesPtrOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseRegistrationDefinition) *RegistrationAssignmentPropertiesResponseProperties {
		return v.Properties
	}).(RegistrationAssignmentPropertiesResponsePropertiesPtrOutput)
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseRegistrationDefinition) SystemDataResponse {
		return v.SystemData
	}).(SystemDataResponseOutput)
}

func (o RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationAssignmentPropertiesResponseRegistrationDefinition) string { return v.Type }).(pulumi.StringOutput)
}

type RegistrationDefinitionProperties struct {
	Authorizations             []Authorization         `pulumi:"authorizations"`
	Description                *string                 `pulumi:"description"`
	EligibleAuthorizations     []EligibleAuthorization `pulumi:"eligibleAuthorizations"`
	ManagedByTenantId          string                  `pulumi:"managedByTenantId"`
	RegistrationDefinitionName *string                 `pulumi:"registrationDefinitionName"`
}





type RegistrationDefinitionPropertiesInput interface {
	pulumi.Input

	ToRegistrationDefinitionPropertiesOutput() RegistrationDefinitionPropertiesOutput
	ToRegistrationDefinitionPropertiesOutputWithContext(context.Context) RegistrationDefinitionPropertiesOutput
}

type RegistrationDefinitionPropertiesArgs struct {
	Authorizations             AuthorizationArrayInput         `pulumi:"authorizations"`
	Description                pulumi.StringPtrInput           `pulumi:"description"`
	EligibleAuthorizations     EligibleAuthorizationArrayInput `pulumi:"eligibleAuthorizations"`
	ManagedByTenantId          pulumi.StringInput              `pulumi:"managedByTenantId"`
	RegistrationDefinitionName pulumi.StringPtrInput           `pulumi:"registrationDefinitionName"`
}

func (RegistrationDefinitionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationDefinitionProperties)(nil)).Elem()
}

func (i RegistrationDefinitionPropertiesArgs) ToRegistrationDefinitionPropertiesOutput() RegistrationDefinitionPropertiesOutput {
	return i.ToRegistrationDefinitionPropertiesOutputWithContext(context.Background())
}

func (i RegistrationDefinitionPropertiesArgs) ToRegistrationDefinitionPropertiesOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationDefinitionPropertiesOutput)
}

func (i RegistrationDefinitionPropertiesArgs) ToRegistrationDefinitionPropertiesPtrOutput() RegistrationDefinitionPropertiesPtrOutput {
	return i.ToRegistrationDefinitionPropertiesPtrOutputWithContext(context.Background())
}

func (i RegistrationDefinitionPropertiesArgs) ToRegistrationDefinitionPropertiesPtrOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationDefinitionPropertiesOutput).ToRegistrationDefinitionPropertiesPtrOutputWithContext(ctx)
}









type RegistrationDefinitionPropertiesPtrInput interface {
	pulumi.Input

	ToRegistrationDefinitionPropertiesPtrOutput() RegistrationDefinitionPropertiesPtrOutput
	ToRegistrationDefinitionPropertiesPtrOutputWithContext(context.Context) RegistrationDefinitionPropertiesPtrOutput
}

type registrationDefinitionPropertiesPtrType RegistrationDefinitionPropertiesArgs

func RegistrationDefinitionPropertiesPtr(v *RegistrationDefinitionPropertiesArgs) RegistrationDefinitionPropertiesPtrInput {
	return (*registrationDefinitionPropertiesPtrType)(v)
}

func (*registrationDefinitionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationDefinitionProperties)(nil)).Elem()
}

func (i *registrationDefinitionPropertiesPtrType) ToRegistrationDefinitionPropertiesPtrOutput() RegistrationDefinitionPropertiesPtrOutput {
	return i.ToRegistrationDefinitionPropertiesPtrOutputWithContext(context.Background())
}

func (i *registrationDefinitionPropertiesPtrType) ToRegistrationDefinitionPropertiesPtrOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationDefinitionPropertiesPtrOutput)
}

type RegistrationDefinitionPropertiesOutput struct{ *pulumi.OutputState }

func (RegistrationDefinitionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationDefinitionProperties)(nil)).Elem()
}

func (o RegistrationDefinitionPropertiesOutput) ToRegistrationDefinitionPropertiesOutput() RegistrationDefinitionPropertiesOutput {
	return o
}

func (o RegistrationDefinitionPropertiesOutput) ToRegistrationDefinitionPropertiesOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesOutput {
	return o
}

func (o RegistrationDefinitionPropertiesOutput) ToRegistrationDefinitionPropertiesPtrOutput() RegistrationDefinitionPropertiesPtrOutput {
	return o.ToRegistrationDefinitionPropertiesPtrOutputWithContext(context.Background())
}

func (o RegistrationDefinitionPropertiesOutput) ToRegistrationDefinitionPropertiesPtrOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RegistrationDefinitionProperties) *RegistrationDefinitionProperties {
		return &v
	}).(RegistrationDefinitionPropertiesPtrOutput)
}

func (o RegistrationDefinitionPropertiesOutput) Authorizations() AuthorizationArrayOutput {
	return o.ApplyT(func(v RegistrationDefinitionProperties) []Authorization { return v.Authorizations }).(AuthorizationArrayOutput)
}

func (o RegistrationDefinitionPropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationDefinitionProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RegistrationDefinitionPropertiesOutput) EligibleAuthorizations() EligibleAuthorizationArrayOutput {
	return o.ApplyT(func(v RegistrationDefinitionProperties) []EligibleAuthorization { return v.EligibleAuthorizations }).(EligibleAuthorizationArrayOutput)
}

func (o RegistrationDefinitionPropertiesOutput) ManagedByTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationDefinitionProperties) string { return v.ManagedByTenantId }).(pulumi.StringOutput)
}

func (o RegistrationDefinitionPropertiesOutput) RegistrationDefinitionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationDefinitionProperties) *string { return v.RegistrationDefinitionName }).(pulumi.StringPtrOutput)
}

type RegistrationDefinitionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (RegistrationDefinitionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationDefinitionProperties)(nil)).Elem()
}

func (o RegistrationDefinitionPropertiesPtrOutput) ToRegistrationDefinitionPropertiesPtrOutput() RegistrationDefinitionPropertiesPtrOutput {
	return o
}

func (o RegistrationDefinitionPropertiesPtrOutput) ToRegistrationDefinitionPropertiesPtrOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesPtrOutput {
	return o
}

func (o RegistrationDefinitionPropertiesPtrOutput) Elem() RegistrationDefinitionPropertiesOutput {
	return o.ApplyT(func(v *RegistrationDefinitionProperties) RegistrationDefinitionProperties {
		if v != nil {
			return *v
		}
		var ret RegistrationDefinitionProperties
		return ret
	}).(RegistrationDefinitionPropertiesOutput)
}

func (o RegistrationDefinitionPropertiesPtrOutput) Authorizations() AuthorizationArrayOutput {
	return o.ApplyT(func(v *RegistrationDefinitionProperties) []Authorization {
		if v == nil {
			return nil
		}
		return v.Authorizations
	}).(AuthorizationArrayOutput)
}

func (o RegistrationDefinitionPropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationDefinitionProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationDefinitionPropertiesPtrOutput) EligibleAuthorizations() EligibleAuthorizationArrayOutput {
	return o.ApplyT(func(v *RegistrationDefinitionProperties) []EligibleAuthorization {
		if v == nil {
			return nil
		}
		return v.EligibleAuthorizations
	}).(EligibleAuthorizationArrayOutput)
}

func (o RegistrationDefinitionPropertiesPtrOutput) ManagedByTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationDefinitionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ManagedByTenantId
	}).(pulumi.StringPtrOutput)
}

func (o RegistrationDefinitionPropertiesPtrOutput) RegistrationDefinitionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationDefinitionProperties) *string {
		if v == nil {
			return nil
		}
		return v.RegistrationDefinitionName
	}).(pulumi.StringPtrOutput)
}

type RegistrationDefinitionPropertiesResponse struct {
	Authorizations             []AuthorizationResponse         `pulumi:"authorizations"`
	Description                *string                         `pulumi:"description"`
	EligibleAuthorizations     []EligibleAuthorizationResponse `pulumi:"eligibleAuthorizations"`
	ManagedByTenantId          string                          `pulumi:"managedByTenantId"`
	ManagedByTenantName        string                          `pulumi:"managedByTenantName"`
	ManageeTenantId            string                          `pulumi:"manageeTenantId"`
	ManageeTenantName          string                          `pulumi:"manageeTenantName"`
	ProvisioningState          string                          `pulumi:"provisioningState"`
	RegistrationDefinitionName *string                         `pulumi:"registrationDefinitionName"`
}

type RegistrationDefinitionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RegistrationDefinitionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationDefinitionPropertiesResponse)(nil)).Elem()
}

func (o RegistrationDefinitionPropertiesResponseOutput) ToRegistrationDefinitionPropertiesResponseOutput() RegistrationDefinitionPropertiesResponseOutput {
	return o
}

func (o RegistrationDefinitionPropertiesResponseOutput) ToRegistrationDefinitionPropertiesResponseOutputWithContext(ctx context.Context) RegistrationDefinitionPropertiesResponseOutput {
	return o
}

func (o RegistrationDefinitionPropertiesResponseOutput) Authorizations() AuthorizationResponseArrayOutput {
	return o.ApplyT(func(v RegistrationDefinitionPropertiesResponse) []AuthorizationResponse { return v.Authorizations }).(AuthorizationResponseArrayOutput)
}

func (o RegistrationDefinitionPropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationDefinitionPropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RegistrationDefinitionPropertiesResponseOutput) EligibleAuthorizations() EligibleAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v RegistrationDefinitionPropertiesResponse) []EligibleAuthorizationResponse {
		return v.EligibleAuthorizations
	}).(EligibleAuthorizationResponseArrayOutput)
}

func (o RegistrationDefinitionPropertiesResponseOutput) ManagedByTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationDefinitionPropertiesResponse) string { return v.ManagedByTenantId }).(pulumi.StringOutput)
}

func (o RegistrationDefinitionPropertiesResponseOutput) ManagedByTenantName() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationDefinitionPropertiesResponse) string { return v.ManagedByTenantName }).(pulumi.StringOutput)
}

func (o RegistrationDefinitionPropertiesResponseOutput) ManageeTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationDefinitionPropertiesResponse) string { return v.ManageeTenantId }).(pulumi.StringOutput)
}

func (o RegistrationDefinitionPropertiesResponseOutput) ManageeTenantName() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationDefinitionPropertiesResponse) string { return v.ManageeTenantName }).(pulumi.StringOutput)
}

func (o RegistrationDefinitionPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationDefinitionPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RegistrationDefinitionPropertiesResponseOutput) RegistrationDefinitionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationDefinitionPropertiesResponse) *string { return v.RegistrationDefinitionName }).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorizationOutput{})
	pulumi.RegisterOutputType(AuthorizationArrayOutput{})
	pulumi.RegisterOutputType(AuthorizationResponseOutput{})
	pulumi.RegisterOutputType(AuthorizationResponseArrayOutput{})
	pulumi.RegisterOutputType(EligibleApproverOutput{})
	pulumi.RegisterOutputType(EligibleApproverArrayOutput{})
	pulumi.RegisterOutputType(EligibleApproverResponseOutput{})
	pulumi.RegisterOutputType(EligibleApproverResponseArrayOutput{})
	pulumi.RegisterOutputType(EligibleAuthorizationOutput{})
	pulumi.RegisterOutputType(EligibleAuthorizationArrayOutput{})
	pulumi.RegisterOutputType(EligibleAuthorizationResponseOutput{})
	pulumi.RegisterOutputType(EligibleAuthorizationResponseArrayOutput{})
	pulumi.RegisterOutputType(JustInTimeAccessPolicyOutput{})
	pulumi.RegisterOutputType(JustInTimeAccessPolicyPtrOutput{})
	pulumi.RegisterOutputType(JustInTimeAccessPolicyResponseOutput{})
	pulumi.RegisterOutputType(JustInTimeAccessPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(PlanOutput{})
	pulumi.RegisterOutputType(PlanPtrOutput{})
	pulumi.RegisterOutputType(PlanResponseOutput{})
	pulumi.RegisterOutputType(PlanResponsePtrOutput{})
	pulumi.RegisterOutputType(RegistrationAssignmentPropertiesOutput{})
	pulumi.RegisterOutputType(RegistrationAssignmentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(RegistrationAssignmentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RegistrationAssignmentPropertiesResponsePropertiesOutput{})
	pulumi.RegisterOutputType(RegistrationAssignmentPropertiesResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(RegistrationAssignmentPropertiesResponseRegistrationDefinitionOutput{})
	pulumi.RegisterOutputType(RegistrationDefinitionPropertiesOutput{})
	pulumi.RegisterOutputType(RegistrationDefinitionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(RegistrationDefinitionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
