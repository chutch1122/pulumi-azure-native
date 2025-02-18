


package v20170401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagementLockAtResourceLevel struct {
	pulumi.CustomResourceState

	Level  pulumi.StringOutput                    `pulumi:"level"`
	Name   pulumi.StringOutput                    `pulumi:"name"`
	Notes  pulumi.StringPtrOutput                 `pulumi:"notes"`
	Owners ManagementLockOwnerResponseArrayOutput `pulumi:"owners"`
	Type   pulumi.StringOutput                    `pulumi:"type"`
}


func NewManagementLockAtResourceLevel(ctx *pulumi.Context,
	name string, args *ManagementLockAtResourceLevelArgs, opts ...pulumi.ResourceOption) (*ManagementLockAtResourceLevel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Level == nil {
		return nil, errors.New("invalid value for required argument 'Level'")
	}
	if args.ParentResourcePath == nil {
		return nil, errors.New("invalid value for required argument 'ParentResourcePath'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.ResourceProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ResourceProviderNamespace'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization:ManagementLockAtResourceLevel"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20160901:ManagementLockAtResourceLevel"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200501:ManagementLockAtResourceLevel"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagementLockAtResourceLevel
	err := ctx.RegisterResource("azure-native:authorization/v20170401:ManagementLockAtResourceLevel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagementLockAtResourceLevel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementLockAtResourceLevelState, opts ...pulumi.ResourceOption) (*ManagementLockAtResourceLevel, error) {
	var resource ManagementLockAtResourceLevel
	err := ctx.ReadResource("azure-native:authorization/v20170401:ManagementLockAtResourceLevel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managementLockAtResourceLevelState struct {
}

type ManagementLockAtResourceLevelState struct {
}

func (ManagementLockAtResourceLevelState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementLockAtResourceLevelState)(nil)).Elem()
}

type managementLockAtResourceLevelArgs struct {
	Level                     string                `pulumi:"level"`
	LockName                  *string               `pulumi:"lockName"`
	Notes                     *string               `pulumi:"notes"`
	Owners                    []ManagementLockOwner `pulumi:"owners"`
	ParentResourcePath        string                `pulumi:"parentResourcePath"`
	ResourceGroupName         string                `pulumi:"resourceGroupName"`
	ResourceName              string                `pulumi:"resourceName"`
	ResourceProviderNamespace string                `pulumi:"resourceProviderNamespace"`
	ResourceType              string                `pulumi:"resourceType"`
}


type ManagementLockAtResourceLevelArgs struct {
	Level                     pulumi.StringInput
	LockName                  pulumi.StringPtrInput
	Notes                     pulumi.StringPtrInput
	Owners                    ManagementLockOwnerArrayInput
	ParentResourcePath        pulumi.StringInput
	ResourceGroupName         pulumi.StringInput
	ResourceName              pulumi.StringInput
	ResourceProviderNamespace pulumi.StringInput
	ResourceType              pulumi.StringInput
}

func (ManagementLockAtResourceLevelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementLockAtResourceLevelArgs)(nil)).Elem()
}

type ManagementLockAtResourceLevelInput interface {
	pulumi.Input

	ToManagementLockAtResourceLevelOutput() ManagementLockAtResourceLevelOutput
	ToManagementLockAtResourceLevelOutputWithContext(ctx context.Context) ManagementLockAtResourceLevelOutput
}

func (*ManagementLockAtResourceLevel) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementLockAtResourceLevel)(nil)).Elem()
}

func (i *ManagementLockAtResourceLevel) ToManagementLockAtResourceLevelOutput() ManagementLockAtResourceLevelOutput {
	return i.ToManagementLockAtResourceLevelOutputWithContext(context.Background())
}

func (i *ManagementLockAtResourceLevel) ToManagementLockAtResourceLevelOutputWithContext(ctx context.Context) ManagementLockAtResourceLevelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementLockAtResourceLevelOutput)
}

type ManagementLockAtResourceLevelOutput struct{ *pulumi.OutputState }

func (ManagementLockAtResourceLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementLockAtResourceLevel)(nil)).Elem()
}

func (o ManagementLockAtResourceLevelOutput) ToManagementLockAtResourceLevelOutput() ManagementLockAtResourceLevelOutput {
	return o
}

func (o ManagementLockAtResourceLevelOutput) ToManagementLockAtResourceLevelOutputWithContext(ctx context.Context) ManagementLockAtResourceLevelOutput {
	return o
}

func (o ManagementLockAtResourceLevelOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementLockAtResourceLevel) pulumi.StringOutput { return v.Level }).(pulumi.StringOutput)
}

func (o ManagementLockAtResourceLevelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementLockAtResourceLevel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagementLockAtResourceLevelOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementLockAtResourceLevel) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o ManagementLockAtResourceLevelOutput) Owners() ManagementLockOwnerResponseArrayOutput {
	return o.ApplyT(func(v *ManagementLockAtResourceLevel) ManagementLockOwnerResponseArrayOutput { return v.Owners }).(ManagementLockOwnerResponseArrayOutput)
}

func (o ManagementLockAtResourceLevelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementLockAtResourceLevel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagementLockAtResourceLevelOutput{})
}
