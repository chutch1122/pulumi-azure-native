


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Addon struct {
	pulumi.CustomResourceState

	AddonType         pulumi.StringPtrOutput `pulumi:"addonType"`
	LicenseKey        pulumi.StringPtrOutput `pulumi:"licenseKey"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewAddon(ctx *pulumi.Context,
	name string, args *AddonArgs, opts ...pulumi.ResourceOption) (*Addon, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs:Addon"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:Addon"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:Addon"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:Addon"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20220501:Addon"),
		},
	})
	opts = append(opts, aliases)
	var resource Addon
	err := ctx.RegisterResource("azure-native:avs/v20210101preview:Addon", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAddon(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddonState, opts ...pulumi.ResourceOption) (*Addon, error) {
	var resource Addon
	err := ctx.ReadResource("azure-native:avs/v20210101preview:Addon", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type addonState struct {
}

type AddonState struct {
}

func (AddonState) ElementType() reflect.Type {
	return reflect.TypeOf((*addonState)(nil)).Elem()
}

type addonArgs struct {
	AddonName         *string `pulumi:"addonName"`
	AddonType         *string `pulumi:"addonType"`
	LicenseKey        *string `pulumi:"licenseKey"`
	PrivateCloudName  string  `pulumi:"privateCloudName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type AddonArgs struct {
	AddonName         pulumi.StringPtrInput
	AddonType         pulumi.StringPtrInput
	LicenseKey        pulumi.StringPtrInput
	PrivateCloudName  pulumi.StringInput
	ResourceGroupName pulumi.StringInput
}

func (AddonArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addonArgs)(nil)).Elem()
}

type AddonInput interface {
	pulumi.Input

	ToAddonOutput() AddonOutput
	ToAddonOutputWithContext(ctx context.Context) AddonOutput
}

func (*Addon) ElementType() reflect.Type {
	return reflect.TypeOf((**Addon)(nil)).Elem()
}

func (i *Addon) ToAddonOutput() AddonOutput {
	return i.ToAddonOutputWithContext(context.Background())
}

func (i *Addon) ToAddonOutputWithContext(ctx context.Context) AddonOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddonOutput)
}

type AddonOutput struct{ *pulumi.OutputState }

func (AddonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Addon)(nil)).Elem()
}

func (o AddonOutput) ToAddonOutput() AddonOutput {
	return o
}

func (o AddonOutput) ToAddonOutputWithContext(ctx context.Context) AddonOutput {
	return o
}

func (o AddonOutput) AddonType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Addon) pulumi.StringPtrOutput { return v.AddonType }).(pulumi.StringPtrOutput)
}

func (o AddonOutput) LicenseKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Addon) pulumi.StringPtrOutput { return v.LicenseKey }).(pulumi.StringPtrOutput)
}

func (o AddonOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Addon) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AddonOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Addon) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AddonOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Addon) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AddonOutput{})
}
