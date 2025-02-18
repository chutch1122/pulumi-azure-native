


package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateLinkForAzureAd struct {
	pulumi.CustomResourceState

	AllTenants     pulumi.BoolPtrOutput     `pulumi:"allTenants"`
	Name           pulumi.StringPtrOutput   `pulumi:"name"`
	OwnerTenantId  pulumi.StringPtrOutput   `pulumi:"ownerTenantId"`
	ResourceGroup  pulumi.StringPtrOutput   `pulumi:"resourceGroup"`
	ResourceName   pulumi.StringPtrOutput   `pulumi:"resourceName"`
	SubscriptionId pulumi.StringPtrOutput   `pulumi:"subscriptionId"`
	Tags           pulumi.StringMapOutput   `pulumi:"tags"`
	Tenants        pulumi.StringArrayOutput `pulumi:"tenants"`
	Type           pulumi.StringOutput      `pulumi:"type"`
}


func NewPrivateLinkForAzureAd(ctx *pulumi.Context,
	name string, args *PrivateLinkForAzureAdArgs, opts ...pulumi.ResourceOption) (*PrivateLinkForAzureAd, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:aadiam:privateLinkForAzureAd"),
		},
		{
			Type: pulumi.String("azure-native:aadiam/v20200301:privateLinkForAzureAd"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateLinkForAzureAd
	err := ctx.RegisterResource("azure-native:aadiam/v20200301preview:privateLinkForAzureAd", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateLinkForAzureAd(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkForAzureAdState, opts ...pulumi.ResourceOption) (*PrivateLinkForAzureAd, error) {
	var resource PrivateLinkForAzureAd
	err := ctx.ReadResource("azure-native:aadiam/v20200301preview:privateLinkForAzureAd", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateLinkForAzureAdState struct {
}

type PrivateLinkForAzureAdState struct {
}

func (PrivateLinkForAzureAdState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkForAzureAdState)(nil)).Elem()
}

type privateLinkForAzureAdArgs struct {
	AllTenants        *bool             `pulumi:"allTenants"`
	Name              *string           `pulumi:"name"`
	OwnerTenantId     *string           `pulumi:"ownerTenantId"`
	PolicyName        *string           `pulumi:"policyName"`
	ResourceGroup     *string           `pulumi:"resourceGroup"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      *string           `pulumi:"resourceName"`
	SubscriptionId    *string           `pulumi:"subscriptionId"`
	Tags              map[string]string `pulumi:"tags"`
	Tenants           []string          `pulumi:"tenants"`
}


type PrivateLinkForAzureAdArgs struct {
	AllTenants        pulumi.BoolPtrInput
	Name              pulumi.StringPtrInput
	OwnerTenantId     pulumi.StringPtrInput
	PolicyName        pulumi.StringPtrInput
	ResourceGroup     pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	SubscriptionId    pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	Tenants           pulumi.StringArrayInput
}

func (PrivateLinkForAzureAdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkForAzureAdArgs)(nil)).Elem()
}

type PrivateLinkForAzureAdInput interface {
	pulumi.Input

	ToPrivateLinkForAzureAdOutput() PrivateLinkForAzureAdOutput
	ToPrivateLinkForAzureAdOutputWithContext(ctx context.Context) PrivateLinkForAzureAdOutput
}

func (*PrivateLinkForAzureAd) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkForAzureAd)(nil)).Elem()
}

func (i *PrivateLinkForAzureAd) ToPrivateLinkForAzureAdOutput() PrivateLinkForAzureAdOutput {
	return i.ToPrivateLinkForAzureAdOutputWithContext(context.Background())
}

func (i *PrivateLinkForAzureAd) ToPrivateLinkForAzureAdOutputWithContext(ctx context.Context) PrivateLinkForAzureAdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkForAzureAdOutput)
}

type PrivateLinkForAzureAdOutput struct{ *pulumi.OutputState }

func (PrivateLinkForAzureAdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkForAzureAd)(nil)).Elem()
}

func (o PrivateLinkForAzureAdOutput) ToPrivateLinkForAzureAdOutput() PrivateLinkForAzureAdOutput {
	return o
}

func (o PrivateLinkForAzureAdOutput) ToPrivateLinkForAzureAdOutputWithContext(ctx context.Context) PrivateLinkForAzureAdOutput {
	return o
}

func (o PrivateLinkForAzureAdOutput) AllTenants() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PrivateLinkForAzureAd) pulumi.BoolPtrOutput { return v.AllTenants }).(pulumi.BoolPtrOutput)
}

func (o PrivateLinkForAzureAdOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkForAzureAd) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkForAzureAdOutput) OwnerTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkForAzureAd) pulumi.StringPtrOutput { return v.OwnerTenantId }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkForAzureAdOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkForAzureAd) pulumi.StringPtrOutput { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkForAzureAdOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkForAzureAd) pulumi.StringPtrOutput { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkForAzureAdOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkForAzureAd) pulumi.StringPtrOutput { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkForAzureAdOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateLinkForAzureAd) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PrivateLinkForAzureAdOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateLinkForAzureAd) pulumi.StringArrayOutput { return v.Tenants }).(pulumi.StringArrayOutput)
}

func (o PrivateLinkForAzureAdOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkForAzureAd) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkForAzureAdOutput{})
}
