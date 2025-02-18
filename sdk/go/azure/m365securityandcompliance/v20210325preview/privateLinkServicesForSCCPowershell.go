


package v20210325preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateLinkServicesForSCCPowershell struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput                    `pulumi:"etag"`
	Identity   ServicesResourceResponseIdentityPtrOutput `pulumi:"identity"`
	Kind       pulumi.StringOutput                       `pulumi:"kind"`
	Location   pulumi.StringOutput                       `pulumi:"location"`
	Name       pulumi.StringOutput                       `pulumi:"name"`
	Properties ServicesPropertiesResponseOutput          `pulumi:"properties"`
	SystemData SystemDataResponseOutput                  `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput                    `pulumi:"tags"`
	Type       pulumi.StringOutput                       `pulumi:"type"`
}


func NewPrivateLinkServicesForSCCPowershell(ctx *pulumi.Context,
	name string, args *PrivateLinkServicesForSCCPowershellArgs, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForSCCPowershell, error) {
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
			Type: pulumi.String("azure-native:m365securityandcompliance:privateLinkServicesForSCCPowershell"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateLinkServicesForSCCPowershell
	err := ctx.RegisterResource("azure-native:m365securityandcompliance/v20210325preview:privateLinkServicesForSCCPowershell", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateLinkServicesForSCCPowershell(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkServicesForSCCPowershellState, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForSCCPowershell, error) {
	var resource PrivateLinkServicesForSCCPowershell
	err := ctx.ReadResource("azure-native:m365securityandcompliance/v20210325preview:privateLinkServicesForSCCPowershell", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateLinkServicesForSCCPowershellState struct {
}

type PrivateLinkServicesForSCCPowershellState struct {
}

func (PrivateLinkServicesForSCCPowershellState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForSCCPowershellState)(nil)).Elem()
}

type privateLinkServicesForSCCPowershellArgs struct {
	Identity          *ServicesResourceIdentity `pulumi:"identity"`
	Kind              Kind                      `pulumi:"kind"`
	Location          *string                   `pulumi:"location"`
	Properties        *ServicesProperties       `pulumi:"properties"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	ResourceName      *string                   `pulumi:"resourceName"`
	Tags              map[string]string         `pulumi:"tags"`
}


type PrivateLinkServicesForSCCPowershellArgs struct {
	Identity          ServicesResourceIdentityPtrInput
	Kind              KindInput
	Location          pulumi.StringPtrInput
	Properties        ServicesPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (PrivateLinkServicesForSCCPowershellArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForSCCPowershellArgs)(nil)).Elem()
}

type PrivateLinkServicesForSCCPowershellInput interface {
	pulumi.Input

	ToPrivateLinkServicesForSCCPowershellOutput() PrivateLinkServicesForSCCPowershellOutput
	ToPrivateLinkServicesForSCCPowershellOutputWithContext(ctx context.Context) PrivateLinkServicesForSCCPowershellOutput
}

func (*PrivateLinkServicesForSCCPowershell) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServicesForSCCPowershell)(nil)).Elem()
}

func (i *PrivateLinkServicesForSCCPowershell) ToPrivateLinkServicesForSCCPowershellOutput() PrivateLinkServicesForSCCPowershellOutput {
	return i.ToPrivateLinkServicesForSCCPowershellOutputWithContext(context.Background())
}

func (i *PrivateLinkServicesForSCCPowershell) ToPrivateLinkServicesForSCCPowershellOutputWithContext(ctx context.Context) PrivateLinkServicesForSCCPowershellOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServicesForSCCPowershellOutput)
}

type PrivateLinkServicesForSCCPowershellOutput struct{ *pulumi.OutputState }

func (PrivateLinkServicesForSCCPowershellOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServicesForSCCPowershell)(nil)).Elem()
}

func (o PrivateLinkServicesForSCCPowershellOutput) ToPrivateLinkServicesForSCCPowershellOutput() PrivateLinkServicesForSCCPowershellOutput {
	return o
}

func (o PrivateLinkServicesForSCCPowershellOutput) ToPrivateLinkServicesForSCCPowershellOutputWithContext(ctx context.Context) PrivateLinkServicesForSCCPowershellOutput {
	return o
}

func (o PrivateLinkServicesForSCCPowershellOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForSCCPowershell) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServicesForSCCPowershellOutput) Identity() ServicesResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForSCCPowershell) ServicesResourceResponseIdentityPtrOutput {
		return v.Identity
	}).(ServicesResourceResponseIdentityPtrOutput)
}

func (o PrivateLinkServicesForSCCPowershellOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForSCCPowershell) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o PrivateLinkServicesForSCCPowershellOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForSCCPowershell) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PrivateLinkServicesForSCCPowershellOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForSCCPowershell) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateLinkServicesForSCCPowershellOutput) Properties() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForSCCPowershell) ServicesPropertiesResponseOutput { return v.Properties }).(ServicesPropertiesResponseOutput)
}

func (o PrivateLinkServicesForSCCPowershellOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForSCCPowershell) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateLinkServicesForSCCPowershellOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForSCCPowershell) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PrivateLinkServicesForSCCPowershellOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForSCCPowershell) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkServicesForSCCPowershellOutput{})
}
