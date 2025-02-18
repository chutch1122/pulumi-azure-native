


package v20200214

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachineImageTemplate struct {
	pulumi.CustomResourceState

	BuildTimeoutInMinutes pulumi.IntPtrOutput                      `pulumi:"buildTimeoutInMinutes"`
	Customize             pulumi.ArrayOutput                       `pulumi:"customize"`
	Distribute            pulumi.ArrayOutput                       `pulumi:"distribute"`
	Identity              ImageTemplateIdentityResponseOutput      `pulumi:"identity"`
	LastRunStatus         ImageTemplateLastRunStatusResponseOutput `pulumi:"lastRunStatus"`
	Location              pulumi.StringOutput                      `pulumi:"location"`
	Name                  pulumi.StringOutput                      `pulumi:"name"`
	ProvisioningError     ProvisioningErrorResponseOutput          `pulumi:"provisioningError"`
	ProvisioningState     pulumi.StringOutput                      `pulumi:"provisioningState"`
	Source                pulumi.AnyOutput                         `pulumi:"source"`
	Tags                  pulumi.StringMapOutput                   `pulumi:"tags"`
	Type                  pulumi.StringOutput                      `pulumi:"type"`
	VmProfile             ImageTemplateVmProfileResponsePtrOutput  `pulumi:"vmProfile"`
}


func NewVirtualMachineImageTemplate(ctx *pulumi.Context,
	name string, args *VirtualMachineImageTemplateArgs, opts ...pulumi.ResourceOption) (*VirtualMachineImageTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Distribute == nil {
		return nil, errors.New("invalid value for required argument 'Distribute'")
	}
	if args.Identity == nil {
		return nil, errors.New("invalid value for required argument 'Identity'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	if isZero(args.BuildTimeoutInMinutes) {
		args.BuildTimeoutInMinutes = pulumi.IntPtr(0)
	}
	if args.VmProfile != nil {
		args.VmProfile = args.VmProfile.ToImageTemplateVmProfilePtrOutput().ApplyT(func(v *ImageTemplateVmProfile) *ImageTemplateVmProfile { return v.Defaults() }).(ImageTemplateVmProfilePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:virtualmachineimages:VirtualMachineImageTemplate"),
		},
		{
			Type: pulumi.String("azure-native:virtualmachineimages/v20180201preview:VirtualMachineImageTemplate"),
		},
		{
			Type: pulumi.String("azure-native:virtualmachineimages/v20190201preview:VirtualMachineImageTemplate"),
		},
		{
			Type: pulumi.String("azure-native:virtualmachineimages/v20190501preview:VirtualMachineImageTemplate"),
		},
		{
			Type: pulumi.String("azure-native:virtualmachineimages/v20211001:VirtualMachineImageTemplate"),
		},
		{
			Type: pulumi.String("azure-native:virtualmachineimages/v20220214:VirtualMachineImageTemplate"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineImageTemplate
	err := ctx.RegisterResource("azure-native:virtualmachineimages/v20200214:VirtualMachineImageTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachineImageTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineImageTemplateState, opts ...pulumi.ResourceOption) (*VirtualMachineImageTemplate, error) {
	var resource VirtualMachineImageTemplate
	err := ctx.ReadResource("azure-native:virtualmachineimages/v20200214:VirtualMachineImageTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualMachineImageTemplateState struct {
}

type VirtualMachineImageTemplateState struct {
}

func (VirtualMachineImageTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineImageTemplateState)(nil)).Elem()
}

type virtualMachineImageTemplateArgs struct {
	BuildTimeoutInMinutes *int                    `pulumi:"buildTimeoutInMinutes"`
	Customize             []interface{}           `pulumi:"customize"`
	Distribute            []interface{}           `pulumi:"distribute"`
	Identity              ImageTemplateIdentity   `pulumi:"identity"`
	ImageTemplateName     *string                 `pulumi:"imageTemplateName"`
	Location              *string                 `pulumi:"location"`
	ResourceGroupName     string                  `pulumi:"resourceGroupName"`
	Source                interface{}             `pulumi:"source"`
	Tags                  map[string]string       `pulumi:"tags"`
	VmProfile             *ImageTemplateVmProfile `pulumi:"vmProfile"`
}


type VirtualMachineImageTemplateArgs struct {
	BuildTimeoutInMinutes pulumi.IntPtrInput
	Customize             pulumi.ArrayInput
	Distribute            pulumi.ArrayInput
	Identity              ImageTemplateIdentityInput
	ImageTemplateName     pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	Source                pulumi.Input
	Tags                  pulumi.StringMapInput
	VmProfile             ImageTemplateVmProfilePtrInput
}

func (VirtualMachineImageTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineImageTemplateArgs)(nil)).Elem()
}

type VirtualMachineImageTemplateInput interface {
	pulumi.Input

	ToVirtualMachineImageTemplateOutput() VirtualMachineImageTemplateOutput
	ToVirtualMachineImageTemplateOutputWithContext(ctx context.Context) VirtualMachineImageTemplateOutput
}

func (*VirtualMachineImageTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineImageTemplate)(nil)).Elem()
}

func (i *VirtualMachineImageTemplate) ToVirtualMachineImageTemplateOutput() VirtualMachineImageTemplateOutput {
	return i.ToVirtualMachineImageTemplateOutputWithContext(context.Background())
}

func (i *VirtualMachineImageTemplate) ToVirtualMachineImageTemplateOutputWithContext(ctx context.Context) VirtualMachineImageTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineImageTemplateOutput)
}

type VirtualMachineImageTemplateOutput struct{ *pulumi.OutputState }

func (VirtualMachineImageTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineImageTemplate)(nil)).Elem()
}

func (o VirtualMachineImageTemplateOutput) ToVirtualMachineImageTemplateOutput() VirtualMachineImageTemplateOutput {
	return o
}

func (o VirtualMachineImageTemplateOutput) ToVirtualMachineImageTemplateOutputWithContext(ctx context.Context) VirtualMachineImageTemplateOutput {
	return o
}

func (o VirtualMachineImageTemplateOutput) BuildTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) pulumi.IntPtrOutput { return v.BuildTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineImageTemplateOutput) Customize() pulumi.ArrayOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) pulumi.ArrayOutput { return v.Customize }).(pulumi.ArrayOutput)
}

func (o VirtualMachineImageTemplateOutput) Distribute() pulumi.ArrayOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) pulumi.ArrayOutput { return v.Distribute }).(pulumi.ArrayOutput)
}

func (o VirtualMachineImageTemplateOutput) Identity() ImageTemplateIdentityResponseOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) ImageTemplateIdentityResponseOutput { return v.Identity }).(ImageTemplateIdentityResponseOutput)
}

func (o VirtualMachineImageTemplateOutput) LastRunStatus() ImageTemplateLastRunStatusResponseOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) ImageTemplateLastRunStatusResponseOutput { return v.LastRunStatus }).(ImageTemplateLastRunStatusResponseOutput)
}

func (o VirtualMachineImageTemplateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualMachineImageTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineImageTemplateOutput) ProvisioningError() ProvisioningErrorResponseOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) ProvisioningErrorResponseOutput { return v.ProvisioningError }).(ProvisioningErrorResponseOutput)
}

func (o VirtualMachineImageTemplateOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineImageTemplateOutput) Source() pulumi.AnyOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) pulumi.AnyOutput { return v.Source }).(pulumi.AnyOutput)
}

func (o VirtualMachineImageTemplateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualMachineImageTemplateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualMachineImageTemplateOutput) VmProfile() ImageTemplateVmProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineImageTemplate) ImageTemplateVmProfileResponsePtrOutput { return v.VmProfile }).(ImageTemplateVmProfileResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineImageTemplateOutput{})
}
