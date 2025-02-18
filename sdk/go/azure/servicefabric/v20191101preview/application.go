


package v20191101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Application struct {
	pulumi.CustomResourceState

	Etag                      pulumi.StringOutput                                `pulumi:"etag"`
	Identity                  ManagedIdentityResponsePtrOutput                   `pulumi:"identity"`
	Location                  pulumi.StringPtrOutput                             `pulumi:"location"`
	ManagedIdentities         ApplicationUserAssignedIdentityResponseArrayOutput `pulumi:"managedIdentities"`
	MaximumNodes              pulumi.Float64PtrOutput                            `pulumi:"maximumNodes"`
	Metrics                   ApplicationMetricDescriptionResponseArrayOutput    `pulumi:"metrics"`
	MinimumNodes              pulumi.Float64PtrOutput                            `pulumi:"minimumNodes"`
	Name                      pulumi.StringOutput                                `pulumi:"name"`
	Parameters                pulumi.StringMapOutput                             `pulumi:"parameters"`
	ProvisioningState         pulumi.StringOutput                                `pulumi:"provisioningState"`
	RemoveApplicationCapacity pulumi.BoolPtrOutput                               `pulumi:"removeApplicationCapacity"`
	Tags                      pulumi.StringMapOutput                             `pulumi:"tags"`
	Type                      pulumi.StringOutput                                `pulumi:"type"`
	TypeName                  pulumi.StringPtrOutput                             `pulumi:"typeName"`
	TypeVersion               pulumi.StringPtrOutput                             `pulumi:"typeVersion"`
	UpgradePolicy             ApplicationUpgradePolicyResponsePtrOutput          `pulumi:"upgradePolicy"`
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
	if isZero(args.MaximumNodes) {
		args.MaximumNodes = pulumi.Float64Ptr(0.0)
	}
	if args.UpgradePolicy != nil {
		args.UpgradePolicy = args.UpgradePolicy.ToApplicationUpgradePolicyPtrOutput().ApplyT(func(v *ApplicationUpgradePolicy) *ApplicationUpgradePolicy { return v.Defaults() }).(ApplicationUpgradePolicyPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20170701preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190301:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190301preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190601preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20200301:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20201201preview:Application"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210601:Application"),
		},
	})
	opts = append(opts, aliases)
	var resource Application
	err := ctx.RegisterResource("azure-native:servicefabric/v20191101preview:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("azure-native:servicefabric/v20191101preview:Application", name, id, state, &resource, opts...)
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
	ApplicationName           *string                           `pulumi:"applicationName"`
	ClusterName               string                            `pulumi:"clusterName"`
	Identity                  *ManagedIdentity                  `pulumi:"identity"`
	Location                  *string                           `pulumi:"location"`
	ManagedIdentities         []ApplicationUserAssignedIdentity `pulumi:"managedIdentities"`
	MaximumNodes              *float64                          `pulumi:"maximumNodes"`
	Metrics                   []ApplicationMetricDescription    `pulumi:"metrics"`
	MinimumNodes              *float64                          `pulumi:"minimumNodes"`
	Parameters                map[string]string                 `pulumi:"parameters"`
	RemoveApplicationCapacity *bool                             `pulumi:"removeApplicationCapacity"`
	ResourceGroupName         string                            `pulumi:"resourceGroupName"`
	Tags                      map[string]string                 `pulumi:"tags"`
	TypeName                  *string                           `pulumi:"typeName"`
	TypeVersion               *string                           `pulumi:"typeVersion"`
	UpgradePolicy             *ApplicationUpgradePolicy         `pulumi:"upgradePolicy"`
}


type ApplicationArgs struct {
	ApplicationName           pulumi.StringPtrInput
	ClusterName               pulumi.StringInput
	Identity                  ManagedIdentityPtrInput
	Location                  pulumi.StringPtrInput
	ManagedIdentities         ApplicationUserAssignedIdentityArrayInput
	MaximumNodes              pulumi.Float64PtrInput
	Metrics                   ApplicationMetricDescriptionArrayInput
	MinimumNodes              pulumi.Float64PtrInput
	Parameters                pulumi.StringMapInput
	RemoveApplicationCapacity pulumi.BoolPtrInput
	ResourceGroupName         pulumi.StringInput
	Tags                      pulumi.StringMapInput
	TypeName                  pulumi.StringPtrInput
	TypeVersion               pulumi.StringPtrInput
	UpgradePolicy             ApplicationUpgradePolicyPtrInput
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

func (o ApplicationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ApplicationOutput) Identity() ManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Application) ManagedIdentityResponsePtrOutput { return v.Identity }).(ManagedIdentityResponsePtrOutput)
}

func (o ApplicationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) ManagedIdentities() ApplicationUserAssignedIdentityResponseArrayOutput {
	return o.ApplyT(func(v *Application) ApplicationUserAssignedIdentityResponseArrayOutput { return v.ManagedIdentities }).(ApplicationUserAssignedIdentityResponseArrayOutput)
}

func (o ApplicationOutput) MaximumNodes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Application) pulumi.Float64PtrOutput { return v.MaximumNodes }).(pulumi.Float64PtrOutput)
}

func (o ApplicationOutput) Metrics() ApplicationMetricDescriptionResponseArrayOutput {
	return o.ApplyT(func(v *Application) ApplicationMetricDescriptionResponseArrayOutput { return v.Metrics }).(ApplicationMetricDescriptionResponseArrayOutput)
}

func (o ApplicationOutput) MinimumNodes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Application) pulumi.Float64PtrOutput { return v.MinimumNodes }).(pulumi.Float64PtrOutput)
}

func (o ApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Application) pulumi.StringMapOutput { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o ApplicationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ApplicationOutput) RemoveApplicationCapacity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.BoolPtrOutput { return v.RemoveApplicationCapacity }).(pulumi.BoolPtrOutput)
}

func (o ApplicationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Application) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApplicationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ApplicationOutput) TypeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.TypeName }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) TypeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.TypeVersion }).(pulumi.StringPtrOutput)
}

func (o ApplicationOutput) UpgradePolicy() ApplicationUpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v *Application) ApplicationUpgradePolicyResponsePtrOutput { return v.UpgradePolicy }).(ApplicationUpgradePolicyResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationOutput{})
}
