


package hdinsight

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Application struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput              `pulumi:"etag"`
	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties ApplicationPropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput              `pulumi:"tags"`
	Type       pulumi.StringOutput                 `pulumi:"type"`
}


func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hdinsight/v20150301preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:hdinsight/v20180601preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:hdinsight/v20210601:Application"),
		},
	})
	opts = append(opts, aliases)
	var resource Application
	err := ctx.RegisterResource("azure-native:hdinsight:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("azure-native:hdinsight:Application", name, id, state, &resource, opts...)
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
	ApplicationName   *string                `pulumi:"applicationName"`
	ClusterName       string                 `pulumi:"clusterName"`
	Properties        *ApplicationProperties `pulumi:"properties"`
	ResourceGroupName string                 `pulumi:"resourceGroupName"`
	Tags              map[string]string      `pulumi:"tags"`
}


type ApplicationArgs struct {
	ApplicationName   pulumi.StringPtrInput
	ClusterName       pulumi.StringInput
	Properties        ApplicationPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
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

func (o ApplicationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationOutput) Properties() ApplicationPropertiesResponseOutput {
	return o.ApplyT(func(v *Application) ApplicationPropertiesResponseOutput { return v.Properties }).(ApplicationPropertiesResponseOutput)
}

func (o ApplicationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Application) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApplicationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationOutput{})
}
