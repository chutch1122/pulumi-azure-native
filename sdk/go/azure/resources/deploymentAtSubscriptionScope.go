


package resources

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DeploymentAtSubscriptionScope struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                     `pulumi:"location"`
	Name       pulumi.StringOutput                        `pulumi:"name"`
	Properties DeploymentPropertiesExtendedResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                     `pulumi:"tags"`
	Type       pulumi.StringOutput                        `pulumi:"type"`
}


func NewDeploymentAtSubscriptionScope(ctx *pulumi.Context,
	name string, args *DeploymentAtSubscriptionScopeArgs, opts ...pulumi.ResourceOption) (*DeploymentAtSubscriptionScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources/v20180501:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190301:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190501:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190510:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190701:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190801:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20191001:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200601:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200801:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210101:DeploymentAtSubscriptionScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210401:DeploymentAtSubscriptionScope"),
		},
	})
	opts = append(opts, aliases)
	var resource DeploymentAtSubscriptionScope
	err := ctx.RegisterResource("azure-native:resources:DeploymentAtSubscriptionScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDeploymentAtSubscriptionScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentAtSubscriptionScopeState, opts ...pulumi.ResourceOption) (*DeploymentAtSubscriptionScope, error) {
	var resource DeploymentAtSubscriptionScope
	err := ctx.ReadResource("azure-native:resources:DeploymentAtSubscriptionScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type deploymentAtSubscriptionScopeState struct {
}

type DeploymentAtSubscriptionScopeState struct {
}

func (DeploymentAtSubscriptionScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentAtSubscriptionScopeState)(nil)).Elem()
}

type deploymentAtSubscriptionScopeArgs struct {
	DeploymentName *string              `pulumi:"deploymentName"`
	Location       *string              `pulumi:"location"`
	Properties     DeploymentProperties `pulumi:"properties"`
	Tags           map[string]string    `pulumi:"tags"`
}


type DeploymentAtSubscriptionScopeArgs struct {
	DeploymentName pulumi.StringPtrInput
	Location       pulumi.StringPtrInput
	Properties     DeploymentPropertiesInput
	Tags           pulumi.StringMapInput
}

func (DeploymentAtSubscriptionScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deploymentAtSubscriptionScopeArgs)(nil)).Elem()
}

type DeploymentAtSubscriptionScopeInput interface {
	pulumi.Input

	ToDeploymentAtSubscriptionScopeOutput() DeploymentAtSubscriptionScopeOutput
	ToDeploymentAtSubscriptionScopeOutputWithContext(ctx context.Context) DeploymentAtSubscriptionScopeOutput
}

func (*DeploymentAtSubscriptionScope) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentAtSubscriptionScope)(nil)).Elem()
}

func (i *DeploymentAtSubscriptionScope) ToDeploymentAtSubscriptionScopeOutput() DeploymentAtSubscriptionScopeOutput {
	return i.ToDeploymentAtSubscriptionScopeOutputWithContext(context.Background())
}

func (i *DeploymentAtSubscriptionScope) ToDeploymentAtSubscriptionScopeOutputWithContext(ctx context.Context) DeploymentAtSubscriptionScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentAtSubscriptionScopeOutput)
}

type DeploymentAtSubscriptionScopeOutput struct{ *pulumi.OutputState }

func (DeploymentAtSubscriptionScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentAtSubscriptionScope)(nil)).Elem()
}

func (o DeploymentAtSubscriptionScopeOutput) ToDeploymentAtSubscriptionScopeOutput() DeploymentAtSubscriptionScopeOutput {
	return o
}

func (o DeploymentAtSubscriptionScopeOutput) ToDeploymentAtSubscriptionScopeOutputWithContext(ctx context.Context) DeploymentAtSubscriptionScopeOutput {
	return o
}

func (o DeploymentAtSubscriptionScopeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentAtSubscriptionScope) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DeploymentAtSubscriptionScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentAtSubscriptionScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DeploymentAtSubscriptionScopeOutput) Properties() DeploymentPropertiesExtendedResponseOutput {
	return o.ApplyT(func(v *DeploymentAtSubscriptionScope) DeploymentPropertiesExtendedResponseOutput { return v.Properties }).(DeploymentPropertiesExtendedResponseOutput)
}

func (o DeploymentAtSubscriptionScopeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DeploymentAtSubscriptionScope) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DeploymentAtSubscriptionScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DeploymentAtSubscriptionScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DeploymentAtSubscriptionScopeOutput{})
}
