


package web

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppBackupConfigurationSlot struct {
	pulumi.CustomResourceState

	BackupName        pulumi.StringPtrOutput                   `pulumi:"backupName"`
	BackupSchedule    BackupScheduleResponsePtrOutput          `pulumi:"backupSchedule"`
	Databases         DatabaseBackupSettingResponseArrayOutput `pulumi:"databases"`
	Enabled           pulumi.BoolPtrOutput                     `pulumi:"enabled"`
	Kind              pulumi.StringPtrOutput                   `pulumi:"kind"`
	Name              pulumi.StringOutput                      `pulumi:"name"`
	StorageAccountUrl pulumi.StringOutput                      `pulumi:"storageAccountUrl"`
	Type              pulumi.StringOutput                      `pulumi:"type"`
}


func NewWebAppBackupConfigurationSlot(ctx *pulumi.Context,
	name string, args *WebAppBackupConfigurationSlotArgs, opts ...pulumi.ResourceOption) (*WebAppBackupConfigurationSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	if args.StorageAccountUrl == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountUrl'")
	}
	if args.BackupSchedule != nil {
		args.BackupSchedule = args.BackupSchedule.ToBackupSchedulePtrOutput().ApplyT(func(v *BackupSchedule) *BackupSchedule { return v.Defaults() }).(BackupSchedulePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppBackupConfigurationSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppBackupConfigurationSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppBackupConfigurationSlot
	err := ctx.RegisterResource("azure-native:web:WebAppBackupConfigurationSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppBackupConfigurationSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppBackupConfigurationSlotState, opts ...pulumi.ResourceOption) (*WebAppBackupConfigurationSlot, error) {
	var resource WebAppBackupConfigurationSlot
	err := ctx.ReadResource("azure-native:web:WebAppBackupConfigurationSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppBackupConfigurationSlotState struct {
}

type WebAppBackupConfigurationSlotState struct {
}

func (WebAppBackupConfigurationSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppBackupConfigurationSlotState)(nil)).Elem()
}

type webAppBackupConfigurationSlotArgs struct {
	BackupName        *string                 `pulumi:"backupName"`
	BackupSchedule    *BackupSchedule         `pulumi:"backupSchedule"`
	Databases         []DatabaseBackupSetting `pulumi:"databases"`
	Enabled           *bool                   `pulumi:"enabled"`
	Kind              *string                 `pulumi:"kind"`
	Name              string                  `pulumi:"name"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	Slot              string                  `pulumi:"slot"`
	StorageAccountUrl string                  `pulumi:"storageAccountUrl"`
}


type WebAppBackupConfigurationSlotArgs struct {
	BackupName        pulumi.StringPtrInput
	BackupSchedule    BackupSchedulePtrInput
	Databases         DatabaseBackupSettingArrayInput
	Enabled           pulumi.BoolPtrInput
	Kind              pulumi.StringPtrInput
	Name              pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Slot              pulumi.StringInput
	StorageAccountUrl pulumi.StringInput
}

func (WebAppBackupConfigurationSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppBackupConfigurationSlotArgs)(nil)).Elem()
}

type WebAppBackupConfigurationSlotInput interface {
	pulumi.Input

	ToWebAppBackupConfigurationSlotOutput() WebAppBackupConfigurationSlotOutput
	ToWebAppBackupConfigurationSlotOutputWithContext(ctx context.Context) WebAppBackupConfigurationSlotOutput
}

func (*WebAppBackupConfigurationSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppBackupConfigurationSlot)(nil)).Elem()
}

func (i *WebAppBackupConfigurationSlot) ToWebAppBackupConfigurationSlotOutput() WebAppBackupConfigurationSlotOutput {
	return i.ToWebAppBackupConfigurationSlotOutputWithContext(context.Background())
}

func (i *WebAppBackupConfigurationSlot) ToWebAppBackupConfigurationSlotOutputWithContext(ctx context.Context) WebAppBackupConfigurationSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppBackupConfigurationSlotOutput)
}

type WebAppBackupConfigurationSlotOutput struct{ *pulumi.OutputState }

func (WebAppBackupConfigurationSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppBackupConfigurationSlot)(nil)).Elem()
}

func (o WebAppBackupConfigurationSlotOutput) ToWebAppBackupConfigurationSlotOutput() WebAppBackupConfigurationSlotOutput {
	return o
}

func (o WebAppBackupConfigurationSlotOutput) ToWebAppBackupConfigurationSlotOutputWithContext(ctx context.Context) WebAppBackupConfigurationSlotOutput {
	return o
}

func (o WebAppBackupConfigurationSlotOutput) BackupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppBackupConfigurationSlot) pulumi.StringPtrOutput { return v.BackupName }).(pulumi.StringPtrOutput)
}

func (o WebAppBackupConfigurationSlotOutput) BackupSchedule() BackupScheduleResponsePtrOutput {
	return o.ApplyT(func(v *WebAppBackupConfigurationSlot) BackupScheduleResponsePtrOutput { return v.BackupSchedule }).(BackupScheduleResponsePtrOutput)
}

func (o WebAppBackupConfigurationSlotOutput) Databases() DatabaseBackupSettingResponseArrayOutput {
	return o.ApplyT(func(v *WebAppBackupConfigurationSlot) DatabaseBackupSettingResponseArrayOutput { return v.Databases }).(DatabaseBackupSettingResponseArrayOutput)
}

func (o WebAppBackupConfigurationSlotOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppBackupConfigurationSlot) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o WebAppBackupConfigurationSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppBackupConfigurationSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppBackupConfigurationSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppBackupConfigurationSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppBackupConfigurationSlotOutput) StorageAccountUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppBackupConfigurationSlot) pulumi.StringOutput { return v.StorageAccountUrl }).(pulumi.StringOutput)
}

func (o WebAppBackupConfigurationSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppBackupConfigurationSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppBackupConfigurationSlotOutput{})
}
