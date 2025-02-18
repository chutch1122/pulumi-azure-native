


package v20210401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type B2CTenant struct {
	pulumi.CustomResourceState

	BillingConfig B2CTenantResourcePropertiesResponseBillingConfigPtrOutput `pulumi:"billingConfig"`
	Location      pulumi.StringOutput                                       `pulumi:"location"`
	Name          pulumi.StringOutput                                       `pulumi:"name"`
	Sku           B2CResourceSKUResponseOutput                              `pulumi:"sku"`
	SystemData    SystemDataResponseOutput                                  `pulumi:"systemData"`
	Tags          pulumi.StringMapOutput                                    `pulumi:"tags"`
	TenantId      pulumi.StringPtrOutput                                    `pulumi:"tenantId"`
	Type          pulumi.StringOutput                                       `pulumi:"type"`
}


func NewB2CTenant(ctx *pulumi.Context,
	name string, args *B2CTenantArgs, opts ...pulumi.ResourceOption) (*B2CTenant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azureactivedirectory:B2CTenant"),
		},
		{
			Type: pulumi.String("azure-native:azureactivedirectory/v20190101preview:B2CTenant"),
		},
	})
	opts = append(opts, aliases)
	var resource B2CTenant
	err := ctx.RegisterResource("azure-native:azureactivedirectory/v20210401:B2CTenant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetB2CTenant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *B2CTenantState, opts ...pulumi.ResourceOption) (*B2CTenant, error) {
	var resource B2CTenant
	err := ctx.ReadResource("azure-native:azureactivedirectory/v20210401:B2CTenant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type b2ctenantState struct {
}

type B2CTenantState struct {
}

func (B2CTenantState) ElementType() reflect.Type {
	return reflect.TypeOf((*b2ctenantState)(nil)).Elem()
}

type b2ctenantArgs struct {
	CountryCode       *string           `pulumi:"countryCode"`
	DisplayName       *string           `pulumi:"displayName"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      *string           `pulumi:"resourceName"`
	Sku               B2CResourceSKU    `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
}


type B2CTenantArgs struct {
	CountryCode       pulumi.StringPtrInput
	DisplayName       pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Sku               B2CResourceSKUInput
	Tags              pulumi.StringMapInput
}

func (B2CTenantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*b2ctenantArgs)(nil)).Elem()
}

type B2CTenantInput interface {
	pulumi.Input

	ToB2CTenantOutput() B2CTenantOutput
	ToB2CTenantOutputWithContext(ctx context.Context) B2CTenantOutput
}

func (*B2CTenant) ElementType() reflect.Type {
	return reflect.TypeOf((**B2CTenant)(nil)).Elem()
}

func (i *B2CTenant) ToB2CTenantOutput() B2CTenantOutput {
	return i.ToB2CTenantOutputWithContext(context.Background())
}

func (i *B2CTenant) ToB2CTenantOutputWithContext(ctx context.Context) B2CTenantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(B2CTenantOutput)
}

type B2CTenantOutput struct{ *pulumi.OutputState }

func (B2CTenantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**B2CTenant)(nil)).Elem()
}

func (o B2CTenantOutput) ToB2CTenantOutput() B2CTenantOutput {
	return o
}

func (o B2CTenantOutput) ToB2CTenantOutputWithContext(ctx context.Context) B2CTenantOutput {
	return o
}

func (o B2CTenantOutput) BillingConfig() B2CTenantResourcePropertiesResponseBillingConfigPtrOutput {
	return o.ApplyT(func(v *B2CTenant) B2CTenantResourcePropertiesResponseBillingConfigPtrOutput { return v.BillingConfig }).(B2CTenantResourcePropertiesResponseBillingConfigPtrOutput)
}

func (o B2CTenantOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *B2CTenant) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o B2CTenantOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *B2CTenant) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o B2CTenantOutput) Sku() B2CResourceSKUResponseOutput {
	return o.ApplyT(func(v *B2CTenant) B2CResourceSKUResponseOutput { return v.Sku }).(B2CResourceSKUResponseOutput)
}

func (o B2CTenantOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *B2CTenant) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o B2CTenantOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *B2CTenant) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o B2CTenantOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *B2CTenant) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o B2CTenantOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *B2CTenant) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(B2CTenantOutput{})
}
