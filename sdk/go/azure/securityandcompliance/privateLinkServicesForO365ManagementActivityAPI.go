


package securityandcompliance

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateLinkServicesForO365ManagementActivityAPI struct {
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


func NewPrivateLinkServicesForO365ManagementActivityAPI(ctx *pulumi.Context,
	name string, args *PrivateLinkServicesForO365ManagementActivityAPIArgs, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForO365ManagementActivityAPI, error) {
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
			Type: pulumi.String("azure-native:securityandcompliance/v20210111:privateLinkServicesForO365ManagementActivityAPI"),
		},
		{
			Type: pulumi.String("azure-native:securityandcompliance/v20210308:privateLinkServicesForO365ManagementActivityAPI"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateLinkServicesForO365ManagementActivityAPI
	err := ctx.RegisterResource("azure-native:securityandcompliance:privateLinkServicesForO365ManagementActivityAPI", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateLinkServicesForO365ManagementActivityAPI(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkServicesForO365ManagementActivityAPIState, opts ...pulumi.ResourceOption) (*PrivateLinkServicesForO365ManagementActivityAPI, error) {
	var resource PrivateLinkServicesForO365ManagementActivityAPI
	err := ctx.ReadResource("azure-native:securityandcompliance:privateLinkServicesForO365ManagementActivityAPI", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateLinkServicesForO365ManagementActivityAPIState struct {
}

type PrivateLinkServicesForO365ManagementActivityAPIState struct {
}

func (PrivateLinkServicesForO365ManagementActivityAPIState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForO365ManagementActivityAPIState)(nil)).Elem()
}

type privateLinkServicesForO365ManagementActivityAPIArgs struct {
	Identity          *ServicesResourceIdentity `pulumi:"identity"`
	Kind              Kind                      `pulumi:"kind"`
	Location          *string                   `pulumi:"location"`
	Properties        *ServicesProperties       `pulumi:"properties"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	ResourceName      *string                   `pulumi:"resourceName"`
	Tags              map[string]string         `pulumi:"tags"`
}


type PrivateLinkServicesForO365ManagementActivityAPIArgs struct {
	Identity          ServicesResourceIdentityPtrInput
	Kind              KindInput
	Location          pulumi.StringPtrInput
	Properties        ServicesPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (PrivateLinkServicesForO365ManagementActivityAPIArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicesForO365ManagementActivityAPIArgs)(nil)).Elem()
}

type PrivateLinkServicesForO365ManagementActivityAPIInput interface {
	pulumi.Input

	ToPrivateLinkServicesForO365ManagementActivityAPIOutput() PrivateLinkServicesForO365ManagementActivityAPIOutput
	ToPrivateLinkServicesForO365ManagementActivityAPIOutputWithContext(ctx context.Context) PrivateLinkServicesForO365ManagementActivityAPIOutput
}

func (*PrivateLinkServicesForO365ManagementActivityAPI) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServicesForO365ManagementActivityAPI)(nil)).Elem()
}

func (i *PrivateLinkServicesForO365ManagementActivityAPI) ToPrivateLinkServicesForO365ManagementActivityAPIOutput() PrivateLinkServicesForO365ManagementActivityAPIOutput {
	return i.ToPrivateLinkServicesForO365ManagementActivityAPIOutputWithContext(context.Background())
}

func (i *PrivateLinkServicesForO365ManagementActivityAPI) ToPrivateLinkServicesForO365ManagementActivityAPIOutputWithContext(ctx context.Context) PrivateLinkServicesForO365ManagementActivityAPIOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServicesForO365ManagementActivityAPIOutput)
}

type PrivateLinkServicesForO365ManagementActivityAPIOutput struct{ *pulumi.OutputState }

func (PrivateLinkServicesForO365ManagementActivityAPIOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServicesForO365ManagementActivityAPI)(nil)).Elem()
}

func (o PrivateLinkServicesForO365ManagementActivityAPIOutput) ToPrivateLinkServicesForO365ManagementActivityAPIOutput() PrivateLinkServicesForO365ManagementActivityAPIOutput {
	return o
}

func (o PrivateLinkServicesForO365ManagementActivityAPIOutput) ToPrivateLinkServicesForO365ManagementActivityAPIOutputWithContext(ctx context.Context) PrivateLinkServicesForO365ManagementActivityAPIOutput {
	return o
}

func (o PrivateLinkServicesForO365ManagementActivityAPIOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForO365ManagementActivityAPI) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServicesForO365ManagementActivityAPIOutput) Identity() ServicesResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForO365ManagementActivityAPI) ServicesResourceResponseIdentityPtrOutput {
		return v.Identity
	}).(ServicesResourceResponseIdentityPtrOutput)
}

func (o PrivateLinkServicesForO365ManagementActivityAPIOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForO365ManagementActivityAPI) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o PrivateLinkServicesForO365ManagementActivityAPIOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForO365ManagementActivityAPI) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PrivateLinkServicesForO365ManagementActivityAPIOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForO365ManagementActivityAPI) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateLinkServicesForO365ManagementActivityAPIOutput) Properties() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForO365ManagementActivityAPI) ServicesPropertiesResponseOutput {
		return v.Properties
	}).(ServicesPropertiesResponseOutput)
}

func (o PrivateLinkServicesForO365ManagementActivityAPIOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForO365ManagementActivityAPI) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateLinkServicesForO365ManagementActivityAPIOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForO365ManagementActivityAPI) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PrivateLinkServicesForO365ManagementActivityAPIOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkServicesForO365ManagementActivityAPI) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkServicesForO365ManagementActivityAPIOutput{})
}
