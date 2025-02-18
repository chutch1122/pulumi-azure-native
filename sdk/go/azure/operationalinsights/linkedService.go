


package operationalinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LinkedService struct {
	pulumi.CustomResourceState

	Name                  pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState     pulumi.StringPtrOutput `pulumi:"provisioningState"`
	ResourceId            pulumi.StringPtrOutput `pulumi:"resourceId"`
	Tags                  pulumi.StringMapOutput `pulumi:"tags"`
	Type                  pulumi.StringOutput    `pulumi:"type"`
	WriteAccessResourceId pulumi.StringPtrOutput `pulumi:"writeAccessResourceId"`
}


func NewLinkedService(ctx *pulumi.Context,
	name string, args *LinkedServiceArgs, opts ...pulumi.ResourceOption) (*LinkedService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights/v20151101preview:LinkedService"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20190801preview:LinkedService"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200301preview:LinkedService"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200801:LinkedService"),
		},
	})
	opts = append(opts, aliases)
	var resource LinkedService
	err := ctx.RegisterResource("azure-native:operationalinsights:LinkedService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLinkedService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceState, opts ...pulumi.ResourceOption) (*LinkedService, error) {
	var resource LinkedService
	err := ctx.ReadResource("azure-native:operationalinsights:LinkedService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type linkedServiceState struct {
}

type LinkedServiceState struct {
}

func (LinkedServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceState)(nil)).Elem()
}

type linkedServiceArgs struct {
	LinkedServiceName     *string           `pulumi:"linkedServiceName"`
	ProvisioningState     *string           `pulumi:"provisioningState"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	ResourceId            *string           `pulumi:"resourceId"`
	Tags                  map[string]string `pulumi:"tags"`
	WorkspaceName         string            `pulumi:"workspaceName"`
	WriteAccessResourceId *string           `pulumi:"writeAccessResourceId"`
}


type LinkedServiceArgs struct {
	LinkedServiceName     pulumi.StringPtrInput
	ProvisioningState     pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	ResourceId            pulumi.StringPtrInput
	Tags                  pulumi.StringMapInput
	WorkspaceName         pulumi.StringInput
	WriteAccessResourceId pulumi.StringPtrInput
}

func (LinkedServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceArgs)(nil)).Elem()
}

type LinkedServiceInput interface {
	pulumi.Input

	ToLinkedServiceOutput() LinkedServiceOutput
	ToLinkedServiceOutputWithContext(ctx context.Context) LinkedServiceOutput
}

func (*LinkedService) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedService)(nil)).Elem()
}

func (i *LinkedService) ToLinkedServiceOutput() LinkedServiceOutput {
	return i.ToLinkedServiceOutputWithContext(context.Background())
}

func (i *LinkedService) ToLinkedServiceOutputWithContext(ctx context.Context) LinkedServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceOutput)
}

type LinkedServiceOutput struct{ *pulumi.OutputState }

func (LinkedServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedService)(nil)).Elem()
}

func (o LinkedServiceOutput) ToLinkedServiceOutput() LinkedServiceOutput {
	return o
}

func (o LinkedServiceOutput) ToLinkedServiceOutputWithContext(ctx context.Context) LinkedServiceOutput {
	return o
}

func (o LinkedServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LinkedServiceOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedService) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LinkedServiceOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedService) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o LinkedServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LinkedService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LinkedServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o LinkedServiceOutput) WriteAccessResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedService) pulumi.StringPtrOutput { return v.WriteAccessResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedServiceOutput{})
}
