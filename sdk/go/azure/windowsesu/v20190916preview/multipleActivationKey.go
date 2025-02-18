


package v20190916preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MultipleActivationKey struct {
	pulumi.CustomResourceState

	AgreementNumber       pulumi.StringPtrOutput `pulumi:"agreementNumber"`
	ExpirationDate        pulumi.StringOutput    `pulumi:"expirationDate"`
	InstalledServerNumber pulumi.IntPtrOutput    `pulumi:"installedServerNumber"`
	IsEligible            pulumi.BoolPtrOutput   `pulumi:"isEligible"`
	Location              pulumi.StringOutput    `pulumi:"location"`
	MultipleActivationKey pulumi.StringOutput    `pulumi:"multipleActivationKey"`
	Name                  pulumi.StringOutput    `pulumi:"name"`
	OsType                pulumi.StringPtrOutput `pulumi:"osType"`
	ProvisioningState     pulumi.StringOutput    `pulumi:"provisioningState"`
	SupportType           pulumi.StringPtrOutput `pulumi:"supportType"`
	Tags                  pulumi.StringMapOutput `pulumi:"tags"`
	Type                  pulumi.StringOutput    `pulumi:"type"`
}


func NewMultipleActivationKey(ctx *pulumi.Context,
	name string, args *MultipleActivationKeyArgs, opts ...pulumi.ResourceOption) (*MultipleActivationKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.SupportType) {
		args.SupportType = pulumi.StringPtr("SupplementalServicing")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:windowsesu:MultipleActivationKey"),
		},
	})
	opts = append(opts, aliases)
	var resource MultipleActivationKey
	err := ctx.RegisterResource("azure-native:windowsesu/v20190916preview:MultipleActivationKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMultipleActivationKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MultipleActivationKeyState, opts ...pulumi.ResourceOption) (*MultipleActivationKey, error) {
	var resource MultipleActivationKey
	err := ctx.ReadResource("azure-native:windowsesu/v20190916preview:MultipleActivationKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type multipleActivationKeyState struct {
}

type MultipleActivationKeyState struct {
}

func (MultipleActivationKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*multipleActivationKeyState)(nil)).Elem()
}

type multipleActivationKeyArgs struct {
	AgreementNumber           *string           `pulumi:"agreementNumber"`
	InstalledServerNumber     *int              `pulumi:"installedServerNumber"`
	IsEligible                *bool             `pulumi:"isEligible"`
	Location                  *string           `pulumi:"location"`
	MultipleActivationKeyName *string           `pulumi:"multipleActivationKeyName"`
	OsType                    *string           `pulumi:"osType"`
	ResourceGroupName         string            `pulumi:"resourceGroupName"`
	SupportType               *string           `pulumi:"supportType"`
	Tags                      map[string]string `pulumi:"tags"`
}


type MultipleActivationKeyArgs struct {
	AgreementNumber           pulumi.StringPtrInput
	InstalledServerNumber     pulumi.IntPtrInput
	IsEligible                pulumi.BoolPtrInput
	Location                  pulumi.StringPtrInput
	MultipleActivationKeyName pulumi.StringPtrInput
	OsType                    pulumi.StringPtrInput
	ResourceGroupName         pulumi.StringInput
	SupportType               pulumi.StringPtrInput
	Tags                      pulumi.StringMapInput
}

func (MultipleActivationKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*multipleActivationKeyArgs)(nil)).Elem()
}

type MultipleActivationKeyInput interface {
	pulumi.Input

	ToMultipleActivationKeyOutput() MultipleActivationKeyOutput
	ToMultipleActivationKeyOutputWithContext(ctx context.Context) MultipleActivationKeyOutput
}

func (*MultipleActivationKey) ElementType() reflect.Type {
	return reflect.TypeOf((**MultipleActivationKey)(nil)).Elem()
}

func (i *MultipleActivationKey) ToMultipleActivationKeyOutput() MultipleActivationKeyOutput {
	return i.ToMultipleActivationKeyOutputWithContext(context.Background())
}

func (i *MultipleActivationKey) ToMultipleActivationKeyOutputWithContext(ctx context.Context) MultipleActivationKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultipleActivationKeyOutput)
}

type MultipleActivationKeyOutput struct{ *pulumi.OutputState }

func (MultipleActivationKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MultipleActivationKey)(nil)).Elem()
}

func (o MultipleActivationKeyOutput) ToMultipleActivationKeyOutput() MultipleActivationKeyOutput {
	return o
}

func (o MultipleActivationKeyOutput) ToMultipleActivationKeyOutputWithContext(ctx context.Context) MultipleActivationKeyOutput {
	return o
}

func (o MultipleActivationKeyOutput) AgreementNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MultipleActivationKey) pulumi.StringPtrOutput { return v.AgreementNumber }).(pulumi.StringPtrOutput)
}

func (o MultipleActivationKeyOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *MultipleActivationKey) pulumi.StringOutput { return v.ExpirationDate }).(pulumi.StringOutput)
}

func (o MultipleActivationKeyOutput) InstalledServerNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MultipleActivationKey) pulumi.IntPtrOutput { return v.InstalledServerNumber }).(pulumi.IntPtrOutput)
}

func (o MultipleActivationKeyOutput) IsEligible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MultipleActivationKey) pulumi.BoolPtrOutput { return v.IsEligible }).(pulumi.BoolPtrOutput)
}

func (o MultipleActivationKeyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MultipleActivationKey) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o MultipleActivationKeyOutput) MultipleActivationKey() pulumi.StringOutput {
	return o.ApplyT(func(v *MultipleActivationKey) pulumi.StringOutput { return v.MultipleActivationKey }).(pulumi.StringOutput)
}

func (o MultipleActivationKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MultipleActivationKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MultipleActivationKeyOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MultipleActivationKey) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o MultipleActivationKeyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *MultipleActivationKey) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MultipleActivationKeyOutput) SupportType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MultipleActivationKey) pulumi.StringPtrOutput { return v.SupportType }).(pulumi.StringPtrOutput)
}

func (o MultipleActivationKeyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MultipleActivationKey) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MultipleActivationKeyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MultipleActivationKey) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MultipleActivationKeyOutput{})
}
