


package containerregistry

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ImportPipeline struct {
	pulumi.CustomResourceState

	Identity          IdentityPropertiesResponsePtrOutput          `pulumi:"identity"`
	Location          pulumi.StringPtrOutput                       `pulumi:"location"`
	Name              pulumi.StringOutput                          `pulumi:"name"`
	Options           pulumi.StringArrayOutput                     `pulumi:"options"`
	ProvisioningState pulumi.StringOutput                          `pulumi:"provisioningState"`
	Source            ImportPipelineSourcePropertiesResponseOutput `pulumi:"source"`
	SystemData        SystemDataResponseOutput                     `pulumi:"systemData"`
	Trigger           PipelineTriggerPropertiesResponsePtrOutput   `pulumi:"trigger"`
	Type              pulumi.StringOutput                          `pulumi:"type"`
}


func NewImportPipeline(ctx *pulumi.Context,
	name string, args *ImportPipelineArgs, opts ...pulumi.ResourceOption) (*ImportPipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	args.Source = args.Source.ToImportPipelineSourcePropertiesOutput().ApplyT(func(v ImportPipelineSourceProperties) ImportPipelineSourceProperties { return *v.Defaults() }).(ImportPipelineSourcePropertiesOutput)
	if args.Trigger != nil {
		args.Trigger = args.Trigger.ToPipelineTriggerPropertiesPtrOutput().ApplyT(func(v *PipelineTriggerProperties) *PipelineTriggerProperties { return v.Defaults() }).(PipelineTriggerPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry/v20191201preview:ImportPipeline"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20201101preview:ImportPipeline"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210601preview:ImportPipeline"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210801preview:ImportPipeline"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20211201preview:ImportPipeline"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20220201preview:ImportPipeline"),
		},
	})
	opts = append(opts, aliases)
	var resource ImportPipeline
	err := ctx.RegisterResource("azure-native:containerregistry:ImportPipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetImportPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImportPipelineState, opts ...pulumi.ResourceOption) (*ImportPipeline, error) {
	var resource ImportPipeline
	err := ctx.ReadResource("azure-native:containerregistry:ImportPipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type importPipelineState struct {
}

type ImportPipelineState struct {
}

func (ImportPipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*importPipelineState)(nil)).Elem()
}

type importPipelineArgs struct {
	Identity           *IdentityProperties            `pulumi:"identity"`
	ImportPipelineName *string                        `pulumi:"importPipelineName"`
	Location           *string                        `pulumi:"location"`
	Options            []string                       `pulumi:"options"`
	RegistryName       string                         `pulumi:"registryName"`
	ResourceGroupName  string                         `pulumi:"resourceGroupName"`
	Source             ImportPipelineSourceProperties `pulumi:"source"`
	Trigger            *PipelineTriggerProperties     `pulumi:"trigger"`
}


type ImportPipelineArgs struct {
	Identity           IdentityPropertiesPtrInput
	ImportPipelineName pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	Options            pulumi.StringArrayInput
	RegistryName       pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	Source             ImportPipelineSourcePropertiesInput
	Trigger            PipelineTriggerPropertiesPtrInput
}

func (ImportPipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*importPipelineArgs)(nil)).Elem()
}

type ImportPipelineInput interface {
	pulumi.Input

	ToImportPipelineOutput() ImportPipelineOutput
	ToImportPipelineOutputWithContext(ctx context.Context) ImportPipelineOutput
}

func (*ImportPipeline) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportPipeline)(nil)).Elem()
}

func (i *ImportPipeline) ToImportPipelineOutput() ImportPipelineOutput {
	return i.ToImportPipelineOutputWithContext(context.Background())
}

func (i *ImportPipeline) ToImportPipelineOutputWithContext(ctx context.Context) ImportPipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportPipelineOutput)
}

type ImportPipelineOutput struct{ *pulumi.OutputState }

func (ImportPipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportPipeline)(nil)).Elem()
}

func (o ImportPipelineOutput) ToImportPipelineOutput() ImportPipelineOutput {
	return o
}

func (o ImportPipelineOutput) ToImportPipelineOutputWithContext(ctx context.Context) ImportPipelineOutput {
	return o
}

func (o ImportPipelineOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ImportPipeline) IdentityPropertiesResponsePtrOutput { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o ImportPipelineOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportPipeline) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ImportPipelineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ImportPipeline) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ImportPipelineOutput) Options() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ImportPipeline) pulumi.StringArrayOutput { return v.Options }).(pulumi.StringArrayOutput)
}

func (o ImportPipelineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ImportPipeline) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ImportPipelineOutput) Source() ImportPipelineSourcePropertiesResponseOutput {
	return o.ApplyT(func(v *ImportPipeline) ImportPipelineSourcePropertiesResponseOutput { return v.Source }).(ImportPipelineSourcePropertiesResponseOutput)
}

func (o ImportPipelineOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ImportPipeline) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ImportPipelineOutput) Trigger() PipelineTriggerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ImportPipeline) PipelineTriggerPropertiesResponsePtrOutput { return v.Trigger }).(PipelineTriggerPropertiesResponsePtrOutput)
}

func (o ImportPipelineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ImportPipeline) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ImportPipelineOutput{})
}
