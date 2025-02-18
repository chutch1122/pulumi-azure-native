


package v20150601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Vault struct {
	pulumi.CustomResourceState

	Location   pulumi.StringOutput           `pulumi:"location"`
	Name       pulumi.StringOutput           `pulumi:"name"`
	Properties VaultPropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput        `pulumi:"tags"`
	Type       pulumi.StringOutput           `pulumi:"type"`
}


func NewVault(ctx *pulumi.Context,
	name string, args *VaultArgs, opts ...pulumi.ResourceOption) (*Vault, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:keyvault:Vault"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20161001:Vault"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20180214:Vault"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20180214preview:Vault"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20190901:Vault"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20200401preview:Vault"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20210401preview:Vault"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20210601preview:Vault"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20211001:Vault"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20211101preview:Vault"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20220701:Vault"),
		},
	})
	opts = append(opts, aliases)
	var resource Vault
	err := ctx.RegisterResource("azure-native:keyvault/v20150601:Vault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VaultState, opts ...pulumi.ResourceOption) (*Vault, error) {
	var resource Vault
	err := ctx.ReadResource("azure-native:keyvault/v20150601:Vault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type vaultState struct {
}

type VaultState struct {
}

func (VaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultState)(nil)).Elem()
}

type vaultArgs struct {
	Location          *string           `pulumi:"location"`
	Properties        VaultProperties   `pulumi:"properties"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	VaultName         *string           `pulumi:"vaultName"`
}


type VaultArgs struct {
	Location          pulumi.StringPtrInput
	Properties        VaultPropertiesInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VaultName         pulumi.StringPtrInput
}

func (VaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vaultArgs)(nil)).Elem()
}

type VaultInput interface {
	pulumi.Input

	ToVaultOutput() VaultOutput
	ToVaultOutputWithContext(ctx context.Context) VaultOutput
}

func (*Vault) ElementType() reflect.Type {
	return reflect.TypeOf((**Vault)(nil)).Elem()
}

func (i *Vault) ToVaultOutput() VaultOutput {
	return i.ToVaultOutputWithContext(context.Background())
}

func (i *Vault) ToVaultOutputWithContext(ctx context.Context) VaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultOutput)
}

type VaultOutput struct{ *pulumi.OutputState }

func (VaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vault)(nil)).Elem()
}

func (o VaultOutput) ToVaultOutput() VaultOutput {
	return o
}

func (o VaultOutput) ToVaultOutputWithContext(ctx context.Context) VaultOutput {
	return o
}

func (o VaultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Vault) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VaultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Vault) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VaultOutput) Properties() VaultPropertiesResponseOutput {
	return o.ApplyT(func(v *Vault) VaultPropertiesResponseOutput { return v.Properties }).(VaultPropertiesResponseOutput)
}

func (o VaultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Vault) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VaultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Vault) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VaultOutput{})
}
