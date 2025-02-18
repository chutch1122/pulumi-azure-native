


package deploymentmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ArtifactSource struct {
	pulumi.CustomResourceState

	ArtifactRoot   pulumi.StringPtrOutput          `pulumi:"artifactRoot"`
	Authentication SasAuthenticationResponseOutput `pulumi:"authentication"`
	Location       pulumi.StringOutput             `pulumi:"location"`
	Name           pulumi.StringOutput             `pulumi:"name"`
	SourceType     pulumi.StringOutput             `pulumi:"sourceType"`
	Tags           pulumi.StringMapOutput          `pulumi:"tags"`
	Type           pulumi.StringOutput             `pulumi:"type"`
}


func NewArtifactSource(ctx *pulumi.Context,
	name string, args *ArtifactSourceArgs, opts ...pulumi.ResourceOption) (*ArtifactSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Authentication == nil {
		return nil, errors.New("invalid value for required argument 'Authentication'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SourceType == nil {
		return nil, errors.New("invalid value for required argument 'SourceType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:deploymentmanager/v20180901preview:ArtifactSource"),
		},
		{
			Type: pulumi.String("azure-native:deploymentmanager/v20191101preview:ArtifactSource"),
		},
	})
	opts = append(opts, aliases)
	var resource ArtifactSource
	err := ctx.RegisterResource("azure-native:deploymentmanager:ArtifactSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetArtifactSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArtifactSourceState, opts ...pulumi.ResourceOption) (*ArtifactSource, error) {
	var resource ArtifactSource
	err := ctx.ReadResource("azure-native:deploymentmanager:ArtifactSource", name, id, state, &resource, opts...)
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
	ArtifactRoot       *string           `pulumi:"artifactRoot"`
	ArtifactSourceName *string           `pulumi:"artifactSourceName"`
	Authentication     SasAuthentication `pulumi:"authentication"`
	Location           *string           `pulumi:"location"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	SourceType         string            `pulumi:"sourceType"`
	Tags               map[string]string `pulumi:"tags"`
}


type ArtifactSourceArgs struct {
	ArtifactRoot       pulumi.StringPtrInput
	ArtifactSourceName pulumi.StringPtrInput
	Authentication     SasAuthenticationInput
	Location           pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	SourceType         pulumi.StringInput
	Tags               pulumi.StringMapInput
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

func (o ArtifactSourceOutput) ArtifactRoot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringPtrOutput { return v.ArtifactRoot }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceOutput) Authentication() SasAuthenticationResponseOutput {
	return o.ApplyT(func(v *ArtifactSource) SasAuthenticationResponseOutput { return v.Authentication }).(SasAuthenticationResponseOutput)
}

func (o ArtifactSourceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ArtifactSourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ArtifactSourceOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

func (o ArtifactSourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ArtifactSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ArtifactSource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ArtifactSourceOutput{})
}
