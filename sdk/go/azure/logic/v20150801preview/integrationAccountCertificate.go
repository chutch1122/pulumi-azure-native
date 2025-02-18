


package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationAccountCertificate struct {
	pulumi.CustomResourceState

	ChangedTime       pulumi.StringOutput                   `pulumi:"changedTime"`
	CreatedTime       pulumi.StringOutput                   `pulumi:"createdTime"`
	Key               KeyVaultKeyReferenceResponsePtrOutput `pulumi:"key"`
	Location          pulumi.StringPtrOutput                `pulumi:"location"`
	Metadata          pulumi.AnyOutput                      `pulumi:"metadata"`
	Name              pulumi.StringPtrOutput                `pulumi:"name"`
	PublicCertificate pulumi.StringPtrOutput                `pulumi:"publicCertificate"`
	Tags              pulumi.StringMapOutput                `pulumi:"tags"`
	Type              pulumi.StringPtrOutput                `pulumi:"type"`
}


func NewIntegrationAccountCertificate(ctx *pulumi.Context,
	name string, args *IntegrationAccountCertificateArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccountCertificate"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:IntegrationAccountCertificate"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccountCertificate
	err := ctx.RegisterResource("azure-native:logic/v20150801preview:IntegrationAccountCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIntegrationAccountCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountCertificateState, opts ...pulumi.ResourceOption) (*IntegrationAccountCertificate, error) {
	var resource IntegrationAccountCertificate
	err := ctx.ReadResource("azure-native:logic/v20150801preview:IntegrationAccountCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type integrationAccountCertificateState struct {
}

type IntegrationAccountCertificateState struct {
}

func (IntegrationAccountCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountCertificateState)(nil)).Elem()
}

type integrationAccountCertificateArgs struct {
	CertificateName        *string               `pulumi:"certificateName"`
	Id                     *string               `pulumi:"id"`
	IntegrationAccountName string                `pulumi:"integrationAccountName"`
	Key                    *KeyVaultKeyReference `pulumi:"key"`
	Location               *string               `pulumi:"location"`
	Metadata               interface{}           `pulumi:"metadata"`
	Name                   *string               `pulumi:"name"`
	PublicCertificate      *string               `pulumi:"publicCertificate"`
	ResourceGroupName      string                `pulumi:"resourceGroupName"`
	Tags                   map[string]string     `pulumi:"tags"`
	Type                   *string               `pulumi:"type"`
}


type IntegrationAccountCertificateArgs struct {
	CertificateName        pulumi.StringPtrInput
	Id                     pulumi.StringPtrInput
	IntegrationAccountName pulumi.StringInput
	Key                    KeyVaultKeyReferencePtrInput
	Location               pulumi.StringPtrInput
	Metadata               pulumi.Input
	Name                   pulumi.StringPtrInput
	PublicCertificate      pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
	Type                   pulumi.StringPtrInput
}

func (IntegrationAccountCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountCertificateArgs)(nil)).Elem()
}

type IntegrationAccountCertificateInput interface {
	pulumi.Input

	ToIntegrationAccountCertificateOutput() IntegrationAccountCertificateOutput
	ToIntegrationAccountCertificateOutputWithContext(ctx context.Context) IntegrationAccountCertificateOutput
}

func (*IntegrationAccountCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountCertificate)(nil)).Elem()
}

func (i *IntegrationAccountCertificate) ToIntegrationAccountCertificateOutput() IntegrationAccountCertificateOutput {
	return i.ToIntegrationAccountCertificateOutputWithContext(context.Background())
}

func (i *IntegrationAccountCertificate) ToIntegrationAccountCertificateOutputWithContext(ctx context.Context) IntegrationAccountCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountCertificateOutput)
}

type IntegrationAccountCertificateOutput struct{ *pulumi.OutputState }

func (IntegrationAccountCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountCertificate)(nil)).Elem()
}

func (o IntegrationAccountCertificateOutput) ToIntegrationAccountCertificateOutput() IntegrationAccountCertificateOutput {
	return o
}

func (o IntegrationAccountCertificateOutput) ToIntegrationAccountCertificateOutputWithContext(ctx context.Context) IntegrationAccountCertificateOutput {
	return o
}

func (o IntegrationAccountCertificateOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountCertificate) pulumi.StringOutput { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o IntegrationAccountCertificateOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountCertificate) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o IntegrationAccountCertificateOutput) Key() KeyVaultKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v *IntegrationAccountCertificate) KeyVaultKeyReferenceResponsePtrOutput { return v.Key }).(KeyVaultKeyReferenceResponsePtrOutput)
}

func (o IntegrationAccountCertificateOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountCertificate) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o IntegrationAccountCertificateOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *IntegrationAccountCertificate) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

func (o IntegrationAccountCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountCertificate) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IntegrationAccountCertificateOutput) PublicCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountCertificate) pulumi.StringPtrOutput { return v.PublicCertificate }).(pulumi.StringPtrOutput)
}

func (o IntegrationAccountCertificateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationAccountCertificate) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o IntegrationAccountCertificateOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountCertificate) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountCertificateOutput{})
}
