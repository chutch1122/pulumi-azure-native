


package v20160515

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ArtifactSource struct {
	pulumi.CustomResourceState

	ArmTemplateFolderPath pulumi.StringPtrOutput `pulumi:"armTemplateFolderPath"`
	BranchRef             pulumi.StringPtrOutput `pulumi:"branchRef"`
	CreatedDate           pulumi.StringOutput    `pulumi:"createdDate"`
	DisplayName           pulumi.StringPtrOutput `pulumi:"displayName"`
	FolderPath            pulumi.StringPtrOutput `pulumi:"folderPath"`
	Location              pulumi.StringPtrOutput `pulumi:"location"`
	Name                  pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState     pulumi.StringPtrOutput `pulumi:"provisioningState"`
	SecurityToken         pulumi.StringPtrOutput `pulumi:"securityToken"`
	SourceType            pulumi.StringPtrOutput `pulumi:"sourceType"`
	Status                pulumi.StringPtrOutput `pulumi:"status"`
	Tags                  pulumi.StringMapOutput `pulumi:"tags"`
	Type                  pulumi.StringOutput    `pulumi:"type"`
	UniqueIdentifier      pulumi.StringPtrOutput `pulumi:"uniqueIdentifier"`
	Uri                   pulumi.StringPtrOutput `pulumi:"uri"`
}


func NewArtifactSource(ctx *pulumi.Context,
	name string, args *ArtifactSourceArgs, opts ...pulumi.ResourceOption) (*ArtifactSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:ArtifactSource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:ArtifactSource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:ArtifactSource"),
		},
	})
	opts = append(opts, aliases)
	var resource ArtifactSource
	err := ctx.RegisterResource("azure-native:devtestlab/v20160515:ArtifactSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetArtifactSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArtifactSourceState, opts ...pulumi.ResourceOption) (*ArtifactSource, error) {
	var resource ArtifactSource
	err := ctx.ReadResource("azure-native:devtestlab/v20160515:ArtifactSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type artifactSourceState struct {
}

type ArtifactSourceState struct {
}

func (ArtifactSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactSourceState)(nil)).Elem()
}

type artifactSourceArgs struct {
	ArmTemplateFolderPath *string           `pulumi:"armTemplateFolderPath"`
	BranchRef             *string           `pulumi:"branchRef"`
	DisplayName           *string           `pulumi:"displayName"`
	FolderPath            *string           `pulumi:"folderPath"`
	LabName               string            `pulumi:"labName"`
	Location              *string           `pulumi:"location"`
	Name                  *string           `pulumi:"name"`
	ProvisioningState     *string           `pulumi:"provisioningState"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	SecurityToken         *string           `pulumi:"securityToken"`
	SourceType            *string           `pulumi:"sourceType"`
	Status                *string           `pulumi:"status"`
	Tags                  map[string]string `pulumi:"tags"`
	UniqueIdentifier      *string           `pulumi:"uniqueIdentifier"`
	Uri                   *string           `pulumi:"uri"`
}


type ArtifactSourceArgs struct {
	ArmTemplateFolderPath pulumi.StringPtrInput
	BranchRef             pulumi.StringPtrInput
	DisplayName           pulumi.StringPtrInput
	FolderPath            pulumi.StringPtrInput
	LabName               pulumi.StringInput
	Location              pulumi.StringPtrInput
	Name                  pulumi.StringPtrInput
	ProvisioningState     pulumi.StringPtrInput
	ResourceGroupName     pulumi.StringInput
	SecurityToken         pulumi.StringPtrInput
	SourceType            pulumi.StringPtrInput
	Status                pulumi.StringPtrInput
	Tags                  pulumi.StringMapInput
	UniqueIdentifier      pulumi.StringPtrInput
	Uri                   pulumi.StringPtrInput
}

func (ArtifactSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactSourceArgs)(nil)).Elem()
}

type ArtifactSourceInput interface {
	pulumi.Input

	ToArtifactSourceOutput() ArtifactSourceOutput
	ToArtifactSourceOutputWithContext(ctx context.Context) ArtifactSourceOutput
}

func (*ArtifactSource) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactSource)(nil)).Elem()
}

func (i *ArtifactSource) ToArtifactSourceOutput() ArtifactSourceOutput {
	return i.ToArtifactSourceOutputWithContext(context.Background())
}

func (i *ArtifactSource) ToArtifactSourceOutputWithContext(ctx context.Context) ArtifactSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactSourceOutput)
}

type ArtifactSourceOutput struct{ *pulumi.OutputState }

func (ArtifactSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactSource)(nil)).Elem()
}

func (o ArtifactSourceOutput) ToArtifactSourceOutput() ArtifactSourceOutput {
	return o
}

func (o ArtifactSourceOutput) ToArtifactSourceOutputWithContext(ctx context.Context) ArtifactSourceOutput {
	return o
}

func (o ArtifactSourceOutput) ArmTemplateFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.ArmTemplateFolderPath }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceOutput) BranchRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.BranchRef }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o ArtifactSourceOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceOutput) FolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.FolderPath }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ArtifactSourceOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceOutput) SecurityToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.SecurityToken }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.SourceType }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ArtifactSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ArtifactSourceOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.Uri }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ArtifactSourceOutput{})
}
