


package v20191210preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Application struct {
	pulumi.CustomResourceState

	CommandLineArguments pulumi.StringPtrOutput `pulumi:"commandLineArguments"`
	CommandLineSetting   pulumi.StringOutput    `pulumi:"commandLineSetting"`
	Description          pulumi.StringPtrOutput `pulumi:"description"`
	FilePath             pulumi.StringPtrOutput `pulumi:"filePath"`
	FriendlyName         pulumi.StringPtrOutput `pulumi:"friendlyName"`
	IconContent          pulumi.StringOutput    `pulumi:"iconContent"`
	IconHash             pulumi.StringOutput    `pulumi:"iconHash"`
	IconIndex            pulumi.IntPtrOutput    `pulumi:"iconIndex"`
	IconPath             pulumi.StringPtrOutput `pulumi:"iconPath"`
	Name                 pulumi.StringOutput    `pulumi:"name"`
	ShowInPortal         pulumi.BoolPtrOutput   `pulumi:"showInPortal"`
	Type                 pulumi.StringOutput    `pulumi:"type"`
}


func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationGroupName'")
	}
	if args.CommandLineSetting == nil {
		return nil, errors.New("invalid value for required argument 'CommandLineSetting'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:desktopvirtualization:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190123preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190924preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20200921preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201019preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201102preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201110preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210114preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210201preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210309preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210401preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210712:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210903preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220210preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220401preview:Application"),
		},
	})
	opts = append(opts, aliases)
	var resource Application
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20191210preview:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20191210preview:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applicationState struct {
}

type ApplicationState struct {
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	ApplicationGroupName string  `pulumi:"applicationGroupName"`
	ApplicationName      *string `pulumi:"applicationName"`
	CommandLineArguments *string `pulumi:"commandLineArguments"`
	CommandLineSetting   string  `pulumi:"commandLineSetting"`
	Description          *string `pulumi:"description"`
	FilePath             *string `pulumi:"filePath"`
	FriendlyName         *string `pulumi:"friendlyName"`
	IconIndex            *int    `pulumi:"iconIndex"`
	IconPath             *string `pulumi:"iconPath"`
	ResourceGroupName    string  `pulumi:"resourceGroupName"`
	ShowInPortal         *bool   `pulumi:"showInPortal"`
}


type ApplicationArgs struct {
	ApplicationGroupName pulumi.StringInput
	ApplicationName      pulumi.StringPtrInput
	CommandLineArguments pulumi.StringPtrInput
	CommandLineSetting   pulumi.StringInput
	Description          pulumi.StringPtrInput
	FilePath             pulumi.StringPtrInput
	FriendlyName         pulumi.StringPtrInput
	IconIndex            pulumi.IntPtrInput
	IconPath             pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	ShowInPortal         pulumi.BoolPtrInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

func (o ApplicationOutput) CommandLineArguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.CommandLineArguments }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) CommandLineSetting() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.CommandLineSetting }).(pulumi.StringOutput)
}

func (o ApplicationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) FilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.FilePath }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) IconContent() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.IconContent }).(pulumi.StringOutput)
}

func (o ApplicationOutput) IconHash() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.IconHash }).(pulumi.StringOutput)
}

func (o ApplicationOutput) IconIndex() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.IntPtrOutput { return v.IconIndex }).(pulumi.IntPtrOutput)
}

func (o ApplicationOutput) IconPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.IconPath }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationOutput) ShowInPortal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.BoolPtrOutput { return v.ShowInPortal }).(pulumi.BoolPtrOutput)
}

func (o ApplicationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationOutput{})
}
