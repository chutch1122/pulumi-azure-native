


package v20200315preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IscsiTarget struct {
	pulumi.CustomResourceState

	Name              pulumi.StringOutput                  `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                  `pulumi:"provisioningState"`
	Status            pulumi.StringOutput                  `pulumi:"status"`
	TargetIqn         pulumi.StringOutput                  `pulumi:"targetIqn"`
	Tpgs              TargetPortalGroupResponseArrayOutput `pulumi:"tpgs"`
	Type              pulumi.StringOutput                  `pulumi:"type"`
}


func NewIscsiTarget(ctx *pulumi.Context,
	name string, args *IscsiTargetArgs, opts ...pulumi.ResourceOption) (*IscsiTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskPoolName == nil {
		return nil, errors.New("invalid value for required argument 'DiskPoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Tpgs == nil {
		return nil, errors.New("invalid value for required argument 'Tpgs'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagepool:IscsiTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagepool/v20210401preview:IscsiTarget"),
		},
		{
			Type: pulumi.String("azure-native:storagepool/v20210801:IscsiTarget"),
		},
	})
	opts = append(opts, aliases)
	var resource IscsiTarget
	err := ctx.RegisterResource("azure-native:storagepool/v20200315preview:IscsiTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIscsiTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IscsiTargetState, opts ...pulumi.ResourceOption) (*IscsiTarget, error) {
	var resource IscsiTarget
	err := ctx.ReadResource("azure-native:storagepool/v20200315preview:IscsiTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type iscsiTargetState struct {
}

type IscsiTargetState struct {
}

func (IscsiTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*iscsiTargetState)(nil)).Elem()
}

type iscsiTargetArgs struct {
	DiskPoolName      string                    `pulumi:"diskPoolName"`
	IscsiTargetName   *string                   `pulumi:"iscsiTargetName"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	TargetIqn         *string                   `pulumi:"targetIqn"`
	Tpgs              []TargetPortalGroupCreate `pulumi:"tpgs"`
}


type IscsiTargetArgs struct {
	DiskPoolName      pulumi.StringInput
	IscsiTargetName   pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	TargetIqn         pulumi.StringPtrInput
	Tpgs              TargetPortalGroupCreateArrayInput
}

func (IscsiTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iscsiTargetArgs)(nil)).Elem()
}

type IscsiTargetInput interface {
	pulumi.Input

	ToIscsiTargetOutput() IscsiTargetOutput
	ToIscsiTargetOutputWithContext(ctx context.Context) IscsiTargetOutput
}

func (*IscsiTarget) ElementType() reflect.Type {
	return reflect.TypeOf((**IscsiTarget)(nil)).Elem()
}

func (i *IscsiTarget) ToIscsiTargetOutput() IscsiTargetOutput {
	return i.ToIscsiTargetOutputWithContext(context.Background())
}

func (i *IscsiTarget) ToIscsiTargetOutputWithContext(ctx context.Context) IscsiTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IscsiTargetOutput)
}

type IscsiTargetOutput struct{ *pulumi.OutputState }

func (IscsiTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IscsiTarget)(nil)).Elem()
}

func (o IscsiTargetOutput) ToIscsiTargetOutput() IscsiTargetOutput {
	return o
}

func (o IscsiTargetOutput) ToIscsiTargetOutputWithContext(ctx context.Context) IscsiTargetOutput {
	return o
}

func (o IscsiTargetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IscsiTargetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o IscsiTargetOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o IscsiTargetOutput) TargetIqn() pulumi.StringOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringOutput { return v.TargetIqn }).(pulumi.StringOutput)
}

func (o IscsiTargetOutput) Tpgs() TargetPortalGroupResponseArrayOutput {
	return o.ApplyT(func(v *IscsiTarget) TargetPortalGroupResponseArrayOutput { return v.Tpgs }).(TargetPortalGroupResponseArrayOutput)
}

func (o IscsiTargetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IscsiTarget) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IscsiTargetOutput{})
}
