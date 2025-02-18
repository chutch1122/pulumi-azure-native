


package v20220125

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GuestConfigurationHCRPAssignment struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                               `pulumi:"location"`
	Name       pulumi.StringPtrOutput                               `pulumi:"name"`
	Properties GuestConfigurationAssignmentPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                             `pulumi:"systemData"`
	Type       pulumi.StringOutput                                  `pulumi:"type"`
}


func NewGuestConfigurationHCRPAssignment(ctx *pulumi.Context,
	name string, args *GuestConfigurationHCRPAssignmentArgs, opts ...pulumi.ResourceOption) (*GuestConfigurationHCRPAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MachineName == nil {
		return nil, errors.New("invalid value for required argument 'MachineName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:guestconfiguration:GuestConfigurationHCRPAssignment"),
		},
		{
			Type: pulumi.String("azure-native:guestconfiguration/v20181120:GuestConfigurationHCRPAssignment"),
		},
		{
			Type: pulumi.String("azure-native:guestconfiguration/v20200625:GuestConfigurationHCRPAssignment"),
		},
		{
			Type: pulumi.String("azure-native:guestconfiguration/v20210125:GuestConfigurationHCRPAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource GuestConfigurationHCRPAssignment
	err := ctx.RegisterResource("azure-native:guestconfiguration/v20220125:GuestConfigurationHCRPAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGuestConfigurationHCRPAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GuestConfigurationHCRPAssignmentState, opts ...pulumi.ResourceOption) (*GuestConfigurationHCRPAssignment, error) {
	var resource GuestConfigurationHCRPAssignment
	err := ctx.ReadResource("azure-native:guestconfiguration/v20220125:GuestConfigurationHCRPAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type guestConfigurationHCRPAssignmentState struct {
}

type GuestConfigurationHCRPAssignmentState struct {
}

func (GuestConfigurationHCRPAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*guestConfigurationHCRPAssignmentState)(nil)).Elem()
}

type guestConfigurationHCRPAssignmentArgs struct {
	GuestConfigurationAssignmentName *string                                 `pulumi:"guestConfigurationAssignmentName"`
	Location                         *string                                 `pulumi:"location"`
	MachineName                      string                                  `pulumi:"machineName"`
	Name                             *string                                 `pulumi:"name"`
	Properties                       *GuestConfigurationAssignmentProperties `pulumi:"properties"`
	ResourceGroupName                string                                  `pulumi:"resourceGroupName"`
}


type GuestConfigurationHCRPAssignmentArgs struct {
	GuestConfigurationAssignmentName pulumi.StringPtrInput
	Location                         pulumi.StringPtrInput
	MachineName                      pulumi.StringInput
	Name                             pulumi.StringPtrInput
	Properties                       GuestConfigurationAssignmentPropertiesPtrInput
	ResourceGroupName                pulumi.StringInput
}

func (GuestConfigurationHCRPAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*guestConfigurationHCRPAssignmentArgs)(nil)).Elem()
}

type GuestConfigurationHCRPAssignmentInput interface {
	pulumi.Input

	ToGuestConfigurationHCRPAssignmentOutput() GuestConfigurationHCRPAssignmentOutput
	ToGuestConfigurationHCRPAssignmentOutputWithContext(ctx context.Context) GuestConfigurationHCRPAssignmentOutput
}

func (*GuestConfigurationHCRPAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationHCRPAssignment)(nil)).Elem()
}

func (i *GuestConfigurationHCRPAssignment) ToGuestConfigurationHCRPAssignmentOutput() GuestConfigurationHCRPAssignmentOutput {
	return i.ToGuestConfigurationHCRPAssignmentOutputWithContext(context.Background())
}

func (i *GuestConfigurationHCRPAssignment) ToGuestConfigurationHCRPAssignmentOutputWithContext(ctx context.Context) GuestConfigurationHCRPAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationHCRPAssignmentOutput)
}

type GuestConfigurationHCRPAssignmentOutput struct{ *pulumi.OutputState }

func (GuestConfigurationHCRPAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationHCRPAssignment)(nil)).Elem()
}

func (o GuestConfigurationHCRPAssignmentOutput) ToGuestConfigurationHCRPAssignmentOutput() GuestConfigurationHCRPAssignmentOutput {
	return o
}

func (o GuestConfigurationHCRPAssignmentOutput) ToGuestConfigurationHCRPAssignmentOutputWithContext(ctx context.Context) GuestConfigurationHCRPAssignmentOutput {
	return o
}

func (o GuestConfigurationHCRPAssignmentOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationHCRPAssignment) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationHCRPAssignmentOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationHCRPAssignment) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationHCRPAssignmentOutput) Properties() GuestConfigurationAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v *GuestConfigurationHCRPAssignment) GuestConfigurationAssignmentPropertiesResponseOutput {
		return v.Properties
	}).(GuestConfigurationAssignmentPropertiesResponseOutput)
}

func (o GuestConfigurationHCRPAssignmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *GuestConfigurationHCRPAssignment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GuestConfigurationHCRPAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GuestConfigurationHCRPAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GuestConfigurationHCRPAssignmentOutput{})
}
