


package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type MachineLearningService struct {
	pulumi.CustomResourceState

	Identity   IdentityResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringPtrOutput    `pulumi:"location"`
	Name       pulumi.StringOutput       `pulumi:"name"`
	Properties pulumi.AnyOutput          `pulumi:"properties"`
	Sku        SkuResponsePtrOutput      `pulumi:"sku"`
	SystemData SystemDataResponseOutput  `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput    `pulumi:"tags"`
	Type       pulumi.StringOutput       `pulumi:"type"`
}


func NewMachineLearningService(ctx *pulumi.Context,
	name string, args *MachineLearningServiceArgs, opts ...pulumi.ResourceOption) (*MachineLearningService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeType == nil {
		return nil, errors.New("invalid value for required argument 'ComputeType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200501preview:MachineLearningService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200515preview:MachineLearningService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200901preview:MachineLearningService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210101:MachineLearningService"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210401:MachineLearningService"),
		},
	})
	opts = append(opts, aliases)
	var resource MachineLearningService
	err := ctx.RegisterResource("azure-native:machinelearningservices:MachineLearningService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMachineLearningService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineLearningServiceState, opts ...pulumi.ResourceOption) (*MachineLearningService, error) {
	var resource MachineLearningService
	err := ctx.ReadResource("azure-native:machinelearningservices:MachineLearningService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type machineLearningServiceState struct {
}

type MachineLearningServiceState struct {
}

func (MachineLearningServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineLearningServiceState)(nil)).Elem()
}

type machineLearningServiceArgs struct {
	ComputeType             string                                       `pulumi:"computeType"`
	Description             *string                                      `pulumi:"description"`
	EnvironmentImageRequest *CreateServiceRequestEnvironmentImageRequest `pulumi:"environmentImageRequest"`
	Keys                    *CreateServiceRequestKeys                    `pulumi:"keys"`
	KvTags                  map[string]string                            `pulumi:"kvTags"`
	Location                *string                                      `pulumi:"location"`
	Properties              map[string]string                            `pulumi:"properties"`
	ResourceGroupName       string                                       `pulumi:"resourceGroupName"`
	ServiceName             *string                                      `pulumi:"serviceName"`
	WorkspaceName           string                                       `pulumi:"workspaceName"`
}


type MachineLearningServiceArgs struct {
	ComputeType             pulumi.StringInput
	Description             pulumi.StringPtrInput
	EnvironmentImageRequest CreateServiceRequestEnvironmentImageRequestPtrInput
	Keys                    CreateServiceRequestKeysPtrInput
	KvTags                  pulumi.StringMapInput
	Location                pulumi.StringPtrInput
	Properties              pulumi.StringMapInput
	ResourceGroupName       pulumi.StringInput
	ServiceName             pulumi.StringPtrInput
	WorkspaceName           pulumi.StringInput
}

func (MachineLearningServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineLearningServiceArgs)(nil)).Elem()
}

type MachineLearningServiceInput interface {
	pulumi.Input

	ToMachineLearningServiceOutput() MachineLearningServiceOutput
	ToMachineLearningServiceOutputWithContext(ctx context.Context) MachineLearningServiceOutput
}

func (*MachineLearningService) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineLearningService)(nil)).Elem()
}

func (i *MachineLearningService) ToMachineLearningServiceOutput() MachineLearningServiceOutput {
	return i.ToMachineLearningServiceOutputWithContext(context.Background())
}

func (i *MachineLearningService) ToMachineLearningServiceOutputWithContext(ctx context.Context) MachineLearningServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineLearningServiceOutput)
}

type MachineLearningServiceOutput struct{ *pulumi.OutputState }

func (MachineLearningServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineLearningService)(nil)).Elem()
}

func (o MachineLearningServiceOutput) ToMachineLearningServiceOutput() MachineLearningServiceOutput {
	return o
}

func (o MachineLearningServiceOutput) ToMachineLearningServiceOutputWithContext(ctx context.Context) MachineLearningServiceOutput {
	return o
}

func (o MachineLearningServiceOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *MachineLearningService) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o MachineLearningServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineLearningService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o MachineLearningServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineLearningService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MachineLearningServiceOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *MachineLearningService) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

func (o MachineLearningServiceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *MachineLearningService) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o MachineLearningServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MachineLearningService) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MachineLearningServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MachineLearningService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MachineLearningServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineLearningService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MachineLearningServiceOutput{})
}
