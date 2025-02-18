


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigurationAssignment struct {
	pulumi.CustomResourceState

	Location                   pulumi.StringPtrOutput   `pulumi:"location"`
	MaintenanceConfigurationId pulumi.StringPtrOutput   `pulumi:"maintenanceConfigurationId"`
	Name                       pulumi.StringOutput      `pulumi:"name"`
	ResourceId                 pulumi.StringPtrOutput   `pulumi:"resourceId"`
	SystemData                 SystemDataResponseOutput `pulumi:"systemData"`
	Type                       pulumi.StringOutput      `pulumi:"type"`
}


func NewConfigurationAssignment(ctx *pulumi.Context,
	name string, args *ConfigurationAssignmentArgs, opts ...pulumi.ResourceOption) (*ConfigurationAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderName == nil {
		return nil, errors.New("invalid value for required argument 'ProviderName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:maintenance:ConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20210401preview:ConfigurationAssignment"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20220701preview:ConfigurationAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource ConfigurationAssignment
	err := ctx.RegisterResource("azure-native:maintenance/v20210901preview:ConfigurationAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetConfigurationAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationAssignmentState, opts ...pulumi.ResourceOption) (*ConfigurationAssignment, error) {
	var resource ConfigurationAssignment
	err := ctx.ReadResource("azure-native:maintenance/v20210901preview:ConfigurationAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type configurationAssignmentState struct {
}

type ConfigurationAssignmentState struct {
}

func (ConfigurationAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAssignmentState)(nil)).Elem()
}

type configurationAssignmentArgs struct {
	ConfigurationAssignmentName *string `pulumi:"configurationAssignmentName"`
	Location                    *string `pulumi:"location"`
	MaintenanceConfigurationId  *string `pulumi:"maintenanceConfigurationId"`
	ProviderName                string  `pulumi:"providerName"`
	ResourceGroupName           string  `pulumi:"resourceGroupName"`
	ResourceId                  *string `pulumi:"resourceId"`
	ResourceName                string  `pulumi:"resourceName"`
	ResourceType                string  `pulumi:"resourceType"`
}


type ConfigurationAssignmentArgs struct {
	ConfigurationAssignmentName pulumi.StringPtrInput
	Location                    pulumi.StringPtrInput
	MaintenanceConfigurationId  pulumi.StringPtrInput
	ProviderName                pulumi.StringInput
	ResourceGroupName           pulumi.StringInput
	ResourceId                  pulumi.StringPtrInput
	ResourceName                pulumi.StringInput
	ResourceType                pulumi.StringInput
}

func (ConfigurationAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAssignmentArgs)(nil)).Elem()
}

type ConfigurationAssignmentInput interface {
	pulumi.Input

	ToConfigurationAssignmentOutput() ConfigurationAssignmentOutput
	ToConfigurationAssignmentOutputWithContext(ctx context.Context) ConfigurationAssignmentOutput
}

func (*ConfigurationAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationAssignment)(nil)).Elem()
}

func (i *ConfigurationAssignment) ToConfigurationAssignmentOutput() ConfigurationAssignmentOutput {
	return i.ToConfigurationAssignmentOutputWithContext(context.Background())
}

func (i *ConfigurationAssignment) ToConfigurationAssignmentOutputWithContext(ctx context.Context) ConfigurationAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationAssignmentOutput)
}

type ConfigurationAssignmentOutput struct{ *pulumi.OutputState }

func (ConfigurationAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationAssignment)(nil)).Elem()
}

func (o ConfigurationAssignmentOutput) ToConfigurationAssignmentOutput() ConfigurationAssignmentOutput {
	return o
}

func (o ConfigurationAssignmentOutput) ToConfigurationAssignmentOutputWithContext(ctx context.Context) ConfigurationAssignmentOutput {
	return o
}

func (o ConfigurationAssignmentOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationAssignment) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ConfigurationAssignmentOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationAssignment) pulumi.StringPtrOutput { return v.MaintenanceConfigurationId }).(pulumi.StringPtrOutput)
}

func (o ConfigurationAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigurationAssignmentOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationAssignment) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o ConfigurationAssignmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigurationAssignment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ConfigurationAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationAssignmentOutput{})
}
