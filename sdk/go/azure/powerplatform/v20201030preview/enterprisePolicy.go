


package v20201030preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnterprisePolicy struct {
	pulumi.CustomResourceState

	Encryption       PropertiesResponseEncryptionPtrOutput       `pulumi:"encryption"`
	Identity         EnterprisePolicyIdentityResponsePtrOutput   `pulumi:"identity"`
	Kind             pulumi.StringOutput                         `pulumi:"kind"`
	Location         pulumi.StringOutput                         `pulumi:"location"`
	Lockbox          PropertiesResponseLockboxPtrOutput          `pulumi:"lockbox"`
	Name             pulumi.StringOutput                         `pulumi:"name"`
	NetworkInjection PropertiesResponseNetworkInjectionPtrOutput `pulumi:"networkInjection"`
	SystemData       SystemDataResponseOutput                    `pulumi:"systemData"`
	SystemId         pulumi.StringOutput                         `pulumi:"systemId"`
	Tags             pulumi.StringMapOutput                      `pulumi:"tags"`
	Type             pulumi.StringOutput                         `pulumi:"type"`
}


func NewEnterprisePolicy(ctx *pulumi.Context,
	name string, args *EnterprisePolicyArgs, opts ...pulumi.ResourceOption) (*EnterprisePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:powerplatform:EnterprisePolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource EnterprisePolicy
	err := ctx.RegisterResource("azure-native:powerplatform/v20201030preview:EnterprisePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEnterprisePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterprisePolicyState, opts ...pulumi.ResourceOption) (*EnterprisePolicy, error) {
	var resource EnterprisePolicy
	err := ctx.ReadResource("azure-native:powerplatform/v20201030preview:EnterprisePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type enterprisePolicyState struct {
}

type EnterprisePolicyState struct {
}

func (EnterprisePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterprisePolicyState)(nil)).Elem()
}

type enterprisePolicyArgs struct {
	Encryption           *PropertiesEncryption       `pulumi:"encryption"`
	EnterprisePolicyName *string                     `pulumi:"enterprisePolicyName"`
	Identity             *EnterprisePolicyIdentity   `pulumi:"identity"`
	Kind                 string                      `pulumi:"kind"`
	Location             *string                     `pulumi:"location"`
	Lockbox              *PropertiesLockbox          `pulumi:"lockbox"`
	NetworkInjection     *PropertiesNetworkInjection `pulumi:"networkInjection"`
	ResourceGroupName    string                      `pulumi:"resourceGroupName"`
	Tags                 map[string]string           `pulumi:"tags"`
}


type EnterprisePolicyArgs struct {
	Encryption           PropertiesEncryptionPtrInput
	EnterprisePolicyName pulumi.StringPtrInput
	Identity             EnterprisePolicyIdentityPtrInput
	Kind                 pulumi.StringInput
	Location             pulumi.StringPtrInput
	Lockbox              PropertiesLockboxPtrInput
	NetworkInjection     PropertiesNetworkInjectionPtrInput
	ResourceGroupName    pulumi.StringInput
	Tags                 pulumi.StringMapInput
}

func (EnterprisePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterprisePolicyArgs)(nil)).Elem()
}

type EnterprisePolicyInput interface {
	pulumi.Input

	ToEnterprisePolicyOutput() EnterprisePolicyOutput
	ToEnterprisePolicyOutputWithContext(ctx context.Context) EnterprisePolicyOutput
}

func (*EnterprisePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterprisePolicy)(nil)).Elem()
}

func (i *EnterprisePolicy) ToEnterprisePolicyOutput() EnterprisePolicyOutput {
	return i.ToEnterprisePolicyOutputWithContext(context.Background())
}

func (i *EnterprisePolicy) ToEnterprisePolicyOutputWithContext(ctx context.Context) EnterprisePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterprisePolicyOutput)
}

type EnterprisePolicyOutput struct{ *pulumi.OutputState }

func (EnterprisePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterprisePolicy)(nil)).Elem()
}

func (o EnterprisePolicyOutput) ToEnterprisePolicyOutput() EnterprisePolicyOutput {
	return o
}

func (o EnterprisePolicyOutput) ToEnterprisePolicyOutputWithContext(ctx context.Context) EnterprisePolicyOutput {
	return o
}

func (o EnterprisePolicyOutput) Encryption() PropertiesResponseEncryptionPtrOutput {
	return o.ApplyT(func(v *EnterprisePolicy) PropertiesResponseEncryptionPtrOutput { return v.Encryption }).(PropertiesResponseEncryptionPtrOutput)
}

func (o EnterprisePolicyOutput) Identity() EnterprisePolicyIdentityResponsePtrOutput {
	return o.ApplyT(func(v *EnterprisePolicy) EnterprisePolicyIdentityResponsePtrOutput { return v.Identity }).(EnterprisePolicyIdentityResponsePtrOutput)
}

func (o EnterprisePolicyOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterprisePolicy) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o EnterprisePolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterprisePolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o EnterprisePolicyOutput) Lockbox() PropertiesResponseLockboxPtrOutput {
	return o.ApplyT(func(v *EnterprisePolicy) PropertiesResponseLockboxPtrOutput { return v.Lockbox }).(PropertiesResponseLockboxPtrOutput)
}

func (o EnterprisePolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterprisePolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EnterprisePolicyOutput) NetworkInjection() PropertiesResponseNetworkInjectionPtrOutput {
	return o.ApplyT(func(v *EnterprisePolicy) PropertiesResponseNetworkInjectionPtrOutput { return v.NetworkInjection }).(PropertiesResponseNetworkInjectionPtrOutput)
}

func (o EnterprisePolicyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EnterprisePolicy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o EnterprisePolicyOutput) SystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterprisePolicy) pulumi.StringOutput { return v.SystemId }).(pulumi.StringOutput)
}

func (o EnterprisePolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EnterprisePolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o EnterprisePolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterprisePolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EnterprisePolicyOutput{})
}
