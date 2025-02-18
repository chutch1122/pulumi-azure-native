


package datashare

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SynchronizationSetting struct {
	pulumi.CustomResourceState

	Kind       pulumi.StringOutput      `pulumi:"kind"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewSynchronizationSetting(ctx *pulumi.Context,
	name string, args *SynchronizationSettingArgs, opts ...pulumi.ResourceOption) (*SynchronizationSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:SynchronizationSetting"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:SynchronizationSetting"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:SynchronizationSetting"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:SynchronizationSetting"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:SynchronizationSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource SynchronizationSetting
	err := ctx.RegisterResource("azure-native:datashare:SynchronizationSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSynchronizationSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SynchronizationSettingState, opts ...pulumi.ResourceOption) (*SynchronizationSetting, error) {
	var resource SynchronizationSetting
	err := ctx.ReadResource("azure-native:datashare:SynchronizationSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type synchronizationSettingState struct {
}

type SynchronizationSettingState struct {
}

func (SynchronizationSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*synchronizationSettingState)(nil)).Elem()
}

type synchronizationSettingArgs struct {
	AccountName                string  `pulumi:"accountName"`
	Kind                       string  `pulumi:"kind"`
	ResourceGroupName          string  `pulumi:"resourceGroupName"`
	ShareName                  string  `pulumi:"shareName"`
	SynchronizationSettingName *string `pulumi:"synchronizationSettingName"`
}


type SynchronizationSettingArgs struct {
	AccountName                pulumi.StringInput
	Kind                       pulumi.StringInput
	ResourceGroupName          pulumi.StringInput
	ShareName                  pulumi.StringInput
	SynchronizationSettingName pulumi.StringPtrInput
}

func (SynchronizationSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*synchronizationSettingArgs)(nil)).Elem()
}

type SynchronizationSettingInput interface {
	pulumi.Input

	ToSynchronizationSettingOutput() SynchronizationSettingOutput
	ToSynchronizationSettingOutputWithContext(ctx context.Context) SynchronizationSettingOutput
}

func (*SynchronizationSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**SynchronizationSetting)(nil)).Elem()
}

func (i *SynchronizationSetting) ToSynchronizationSettingOutput() SynchronizationSettingOutput {
	return i.ToSynchronizationSettingOutputWithContext(context.Background())
}

func (i *SynchronizationSetting) ToSynchronizationSettingOutputWithContext(ctx context.Context) SynchronizationSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynchronizationSettingOutput)
}

type SynchronizationSettingOutput struct{ *pulumi.OutputState }

func (SynchronizationSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SynchronizationSetting)(nil)).Elem()
}

func (o SynchronizationSettingOutput) ToSynchronizationSettingOutput() SynchronizationSettingOutput {
	return o
}

func (o SynchronizationSettingOutput) ToSynchronizationSettingOutputWithContext(ctx context.Context) SynchronizationSettingOutput {
	return o
}

func (o SynchronizationSettingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *SynchronizationSetting) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o SynchronizationSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SynchronizationSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SynchronizationSettingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SynchronizationSetting) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SynchronizationSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SynchronizationSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SynchronizationSettingOutput{})
}
