


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppServiceCertificateOrderCertificate struct {
	pulumi.CustomResourceState

	KeyVaultId         pulumi.StringPtrOutput `pulumi:"keyVaultId"`
	KeyVaultSecretName pulumi.StringPtrOutput `pulumi:"keyVaultSecretName"`
	Kind               pulumi.StringPtrOutput `pulumi:"kind"`
	Location           pulumi.StringOutput    `pulumi:"location"`
	Name               pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput    `pulumi:"provisioningState"`
	Tags               pulumi.StringMapOutput `pulumi:"tags"`
	Type               pulumi.StringOutput    `pulumi:"type"`
}


func NewAppServiceCertificateOrderCertificate(ctx *pulumi.Context,
	name string, args *AppServiceCertificateOrderCertificateArgs, opts ...pulumi.ResourceOption) (*AppServiceCertificateOrderCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateOrderName == nil {
		return nil, errors.New("invalid value for required argument 'CertificateOrderName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:certificateregistration:AppServiceCertificateOrderCertificate"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20150801:AppServiceCertificateOrderCertificate"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20180201:AppServiceCertificateOrderCertificate"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20190801:AppServiceCertificateOrderCertificate"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20200601:AppServiceCertificateOrderCertificate"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20200901:AppServiceCertificateOrderCertificate"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20201001:AppServiceCertificateOrderCertificate"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20201201:AppServiceCertificateOrderCertificate"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20210115:AppServiceCertificateOrderCertificate"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20210201:AppServiceCertificateOrderCertificate"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20210301:AppServiceCertificateOrderCertificate"),
		},
		{
			Type: pulumi.String("azure-native:certificateregistration/v20220301:AppServiceCertificateOrderCertificate"),
		},
	})
	opts = append(opts, aliases)
	var resource AppServiceCertificateOrderCertificate
	err := ctx.RegisterResource("azure-native:certificateregistration/v20210101:AppServiceCertificateOrderCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAppServiceCertificateOrderCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServiceCertificateOrderCertificateState, opts ...pulumi.ResourceOption) (*AppServiceCertificateOrderCertificate, error) {
	var resource AppServiceCertificateOrderCertificate
	err := ctx.ReadResource("azure-native:certificateregistration/v20210101:AppServiceCertificateOrderCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type appServiceCertificateOrderCertificateState struct {
}

type AppServiceCertificateOrderCertificateState struct {
}

func (AppServiceCertificateOrderCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceCertificateOrderCertificateState)(nil)).Elem()
}

type appServiceCertificateOrderCertificateArgs struct {
	CertificateOrderName string            `pulumi:"certificateOrderName"`
	KeyVaultId           *string           `pulumi:"keyVaultId"`
	KeyVaultSecretName   *string           `pulumi:"keyVaultSecretName"`
	Kind                 *string           `pulumi:"kind"`
	Location             *string           `pulumi:"location"`
	Name                 *string           `pulumi:"name"`
	ResourceGroupName    string            `pulumi:"resourceGroupName"`
	Tags                 map[string]string `pulumi:"tags"`
}


type AppServiceCertificateOrderCertificateArgs struct {
	CertificateOrderName pulumi.StringInput
	KeyVaultId           pulumi.StringPtrInput
	KeyVaultSecretName   pulumi.StringPtrInput
	Kind                 pulumi.StringPtrInput
	Location             pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	Tags                 pulumi.StringMapInput
}

func (AppServiceCertificateOrderCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceCertificateOrderCertificateArgs)(nil)).Elem()
}

type AppServiceCertificateOrderCertificateInput interface {
	pulumi.Input

	ToAppServiceCertificateOrderCertificateOutput() AppServiceCertificateOrderCertificateOutput
	ToAppServiceCertificateOrderCertificateOutputWithContext(ctx context.Context) AppServiceCertificateOrderCertificateOutput
}

func (*AppServiceCertificateOrderCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**AppServiceCertificateOrderCertificate)(nil)).Elem()
}

func (i *AppServiceCertificateOrderCertificate) ToAppServiceCertificateOrderCertificateOutput() AppServiceCertificateOrderCertificateOutput {
	return i.ToAppServiceCertificateOrderCertificateOutputWithContext(context.Background())
}

func (i *AppServiceCertificateOrderCertificate) ToAppServiceCertificateOrderCertificateOutputWithContext(ctx context.Context) AppServiceCertificateOrderCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServiceCertificateOrderCertificateOutput)
}

type AppServiceCertificateOrderCertificateOutput struct{ *pulumi.OutputState }

func (AppServiceCertificateOrderCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppServiceCertificateOrderCertificate)(nil)).Elem()
}

func (o AppServiceCertificateOrderCertificateOutput) ToAppServiceCertificateOrderCertificateOutput() AppServiceCertificateOrderCertificateOutput {
	return o
}

func (o AppServiceCertificateOrderCertificateOutput) ToAppServiceCertificateOrderCertificateOutputWithContext(ctx context.Context) AppServiceCertificateOrderCertificateOutput {
	return o
}

func (o AppServiceCertificateOrderCertificateOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrderCertificate) pulumi.StringPtrOutput { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

func (o AppServiceCertificateOrderCertificateOutput) KeyVaultSecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrderCertificate) pulumi.StringPtrOutput { return v.KeyVaultSecretName }).(pulumi.StringPtrOutput)
}

func (o AppServiceCertificateOrderCertificateOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrderCertificate) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o AppServiceCertificateOrderCertificateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrderCertificate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AppServiceCertificateOrderCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrderCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AppServiceCertificateOrderCertificateOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrderCertificate) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AppServiceCertificateOrderCertificateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrderCertificate) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AppServiceCertificateOrderCertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AppServiceCertificateOrderCertificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AppServiceCertificateOrderCertificateOutput{})
}
