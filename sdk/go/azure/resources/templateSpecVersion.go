


package resources

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TemplateSpecVersion struct {
	pulumi.CustomResourceState

	Description      pulumi.StringPtrOutput                    `pulumi:"description"`
	LinkedTemplates  LinkedTemplateArtifactResponseArrayOutput `pulumi:"linkedTemplates"`
	Location         pulumi.StringOutput                       `pulumi:"location"`
	MainTemplate     pulumi.AnyOutput                          `pulumi:"mainTemplate"`
	Metadata         pulumi.AnyOutput                          `pulumi:"metadata"`
	Name             pulumi.StringOutput                       `pulumi:"name"`
	SystemData       SystemDataResponseOutput                  `pulumi:"systemData"`
	Tags             pulumi.StringMapOutput                    `pulumi:"tags"`
	Type             pulumi.StringOutput                       `pulumi:"type"`
	UiFormDefinition pulumi.AnyOutput                          `pulumi:"uiFormDefinition"`
}


func NewTemplateSpecVersion(ctx *pulumi.Context,
	name string, args *TemplateSpecVersionArgs, opts ...pulumi.ResourceOption) (*TemplateSpecVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TemplateSpecName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateSpecName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources/v20190601preview:TemplateSpecVersion"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210301preview:TemplateSpecVersion"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210501:TemplateSpecVersion"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20220201:TemplateSpecVersion"),
		},
	})
	opts = append(opts, aliases)
	var resource TemplateSpecVersion
	err := ctx.RegisterResource("azure-native:resources:TemplateSpecVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTemplateSpecVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateSpecVersionState, opts ...pulumi.ResourceOption) (*TemplateSpecVersion, error) {
	var resource TemplateSpecVersion
	err := ctx.ReadResource("azure-native:resources:TemplateSpecVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type templateSpecVersionState struct {
}

type TemplateSpecVersionState struct {
}

func (TemplateSpecVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateSpecVersionState)(nil)).Elem()
}

type templateSpecVersionArgs struct {
	Description         *string                  `pulumi:"description"`
	LinkedTemplates     []LinkedTemplateArtifact `pulumi:"linkedTemplates"`
	Location            *string                  `pulumi:"location"`
	MainTemplate        interface{}              `pulumi:"mainTemplate"`
	Metadata            interface{}              `pulumi:"metadata"`
	ResourceGroupName   string                   `pulumi:"resourceGroupName"`
	Tags                map[string]string        `pulumi:"tags"`
	TemplateSpecName    string                   `pulumi:"templateSpecName"`
	TemplateSpecVersion *string                  `pulumi:"templateSpecVersion"`
	UiFormDefinition    interface{}              `pulumi:"uiFormDefinition"`
}


type TemplateSpecVersionArgs struct {
	Description         pulumi.StringPtrInput
	LinkedTemplates     LinkedTemplateArtifactArrayInput
	Location            pulumi.StringPtrInput
	MainTemplate        pulumi.Input
	Metadata            pulumi.Input
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
	TemplateSpecName    pulumi.StringInput
	TemplateSpecVersion pulumi.StringPtrInput
	UiFormDefinition    pulumi.Input
}

func (TemplateSpecVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateSpecVersionArgs)(nil)).Elem()
}

type TemplateSpecVersionInput interface {
	pulumi.Input

	ToTemplateSpecVersionOutput() TemplateSpecVersionOutput
	ToTemplateSpecVersionOutputWithContext(ctx context.Context) TemplateSpecVersionOutput
}

func (*TemplateSpecVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateSpecVersion)(nil)).Elem()
}

func (i *TemplateSpecVersion) ToTemplateSpecVersionOutput() TemplateSpecVersionOutput {
	return i.ToTemplateSpecVersionOutputWithContext(context.Background())
}

func (i *TemplateSpecVersion) ToTemplateSpecVersionOutputWithContext(ctx context.Context) TemplateSpecVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateSpecVersionOutput)
}

type TemplateSpecVersionOutput struct{ *pulumi.OutputState }

func (TemplateSpecVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateSpecVersion)(nil)).Elem()
}

func (o TemplateSpecVersionOutput) ToTemplateSpecVersionOutput() TemplateSpecVersionOutput {
	return o
}

func (o TemplateSpecVersionOutput) ToTemplateSpecVersionOutputWithContext(ctx context.Context) TemplateSpecVersionOutput {
	return o
}

func (o TemplateSpecVersionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TemplateSpecVersionOutput) LinkedTemplates() LinkedTemplateArtifactResponseArrayOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) LinkedTemplateArtifactResponseArrayOutput { return v.LinkedTemplates }).(LinkedTemplateArtifactResponseArrayOutput)
}

func (o TemplateSpecVersionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o TemplateSpecVersionOutput) MainTemplate() pulumi.AnyOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.AnyOutput { return v.MainTemplate }).(pulumi.AnyOutput)
}

func (o TemplateSpecVersionOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

func (o TemplateSpecVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TemplateSpecVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o TemplateSpecVersionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o TemplateSpecVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o TemplateSpecVersionOutput) UiFormDefinition() pulumi.AnyOutput {
	return o.ApplyT(func(v *TemplateSpecVersion) pulumi.AnyOutput { return v.UiFormDefinition }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(TemplateSpecVersionOutput{})
}
