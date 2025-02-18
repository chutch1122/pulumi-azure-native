


package v20220301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StaticSiteUserProvidedFunctionAppForStaticSiteBuild struct {
	pulumi.CustomResourceState

	CreatedOn             pulumi.StringOutput    `pulumi:"createdOn"`
	FunctionAppRegion     pulumi.StringPtrOutput `pulumi:"functionAppRegion"`
	FunctionAppResourceId pulumi.StringPtrOutput `pulumi:"functionAppResourceId"`
	Kind                  pulumi.StringPtrOutput `pulumi:"kind"`
	Name                  pulumi.StringOutput    `pulumi:"name"`
	Type                  pulumi.StringOutput    `pulumi:"type"`
}


func NewStaticSiteUserProvidedFunctionAppForStaticSiteBuild(ctx *pulumi.Context,
	name string, args *StaticSiteUserProvidedFunctionAppForStaticSiteBuildArgs, opts ...pulumi.ResourceOption) (*StaticSiteUserProvidedFunctionAppForStaticSiteBuild, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:StaticSiteUserProvidedFunctionAppForStaticSiteBuild"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:StaticSiteUserProvidedFunctionAppForStaticSiteBuild"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:StaticSiteUserProvidedFunctionAppForStaticSiteBuild"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:StaticSiteUserProvidedFunctionAppForStaticSiteBuild"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:StaticSiteUserProvidedFunctionAppForStaticSiteBuild"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:StaticSiteUserProvidedFunctionAppForStaticSiteBuild"),
		},
	})
	opts = append(opts, aliases)
	var resource StaticSiteUserProvidedFunctionAppForStaticSiteBuild
	err := ctx.RegisterResource("azure-native:web/v20220301:StaticSiteUserProvidedFunctionAppForStaticSiteBuild", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStaticSiteUserProvidedFunctionAppForStaticSiteBuild(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticSiteUserProvidedFunctionAppForStaticSiteBuildState, opts ...pulumi.ResourceOption) (*StaticSiteUserProvidedFunctionAppForStaticSiteBuild, error) {
	var resource StaticSiteUserProvidedFunctionAppForStaticSiteBuild
	err := ctx.ReadResource("azure-native:web/v20220301:StaticSiteUserProvidedFunctionAppForStaticSiteBuild", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type staticSiteUserProvidedFunctionAppForStaticSiteBuildState struct {
}

type StaticSiteUserProvidedFunctionAppForStaticSiteBuildState struct {
}

func (StaticSiteUserProvidedFunctionAppForStaticSiteBuildState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteUserProvidedFunctionAppForStaticSiteBuildState)(nil)).Elem()
}

type staticSiteUserProvidedFunctionAppForStaticSiteBuildArgs struct {
	EnvironmentName       string  `pulumi:"environmentName"`
	FunctionAppName       *string `pulumi:"functionAppName"`
	FunctionAppRegion     *string `pulumi:"functionAppRegion"`
	FunctionAppResourceId *string `pulumi:"functionAppResourceId"`
	IsForced              *bool   `pulumi:"isForced"`
	Kind                  *string `pulumi:"kind"`
	Name                  string  `pulumi:"name"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
}


type StaticSiteUserProvidedFunctionAppForStaticSiteBuildArgs struct {
	EnvironmentName       pulumi.StringInput
	FunctionAppName       pulumi.StringPtrInput
	FunctionAppRegion     pulumi.StringPtrInput
	FunctionAppResourceId pulumi.StringPtrInput
	IsForced              pulumi.BoolPtrInput
	Kind                  pulumi.StringPtrInput
	Name                  pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
}

func (StaticSiteUserProvidedFunctionAppForStaticSiteBuildArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteUserProvidedFunctionAppForStaticSiteBuildArgs)(nil)).Elem()
}

type StaticSiteUserProvidedFunctionAppForStaticSiteBuildInput interface {
	pulumi.Input

	ToStaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput() StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput
	ToStaticSiteUserProvidedFunctionAppForStaticSiteBuildOutputWithContext(ctx context.Context) StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput
}

func (*StaticSiteUserProvidedFunctionAppForStaticSiteBuild) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteUserProvidedFunctionAppForStaticSiteBuild)(nil)).Elem()
}

func (i *StaticSiteUserProvidedFunctionAppForStaticSiteBuild) ToStaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput() StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput {
	return i.ToStaticSiteUserProvidedFunctionAppForStaticSiteBuildOutputWithContext(context.Background())
}

func (i *StaticSiteUserProvidedFunctionAppForStaticSiteBuild) ToStaticSiteUserProvidedFunctionAppForStaticSiteBuildOutputWithContext(ctx context.Context) StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput)
}

type StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput struct{ *pulumi.OutputState }

func (StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteUserProvidedFunctionAppForStaticSiteBuild)(nil)).Elem()
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput) ToStaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput() StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput {
	return o
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput) ToStaticSiteUserProvidedFunctionAppForStaticSiteBuildOutputWithContext(ctx context.Context) StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput {
	return o
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSiteBuild) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput) FunctionAppRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSiteBuild) pulumi.StringPtrOutput {
		return v.FunctionAppRegion
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput) FunctionAppResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSiteBuild) pulumi.StringPtrOutput {
		return v.FunctionAppResourceId
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSiteBuild) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSiteBuild) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSiteUserProvidedFunctionAppForStaticSiteBuild) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StaticSiteUserProvidedFunctionAppForStaticSiteBuildOutput{})
}
