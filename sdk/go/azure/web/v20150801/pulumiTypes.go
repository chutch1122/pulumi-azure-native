


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiDefinitionInfo struct {
	Url *string `pulumi:"url"`
}





type ApiDefinitionInfoInput interface {
	pulumi.Input

	ToApiDefinitionInfoOutput() ApiDefinitionInfoOutput
	ToApiDefinitionInfoOutputWithContext(context.Context) ApiDefinitionInfoOutput
}

type ApiDefinitionInfoArgs struct {
	Url pulumi.StringPtrInput `pulumi:"url"`
}

func (ApiDefinitionInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiDefinitionInfo)(nil)).Elem()
}

func (i ApiDefinitionInfoArgs) ToApiDefinitionInfoOutput() ApiDefinitionInfoOutput {
	return i.ToApiDefinitionInfoOutputWithContext(context.Background())
}

func (i ApiDefinitionInfoArgs) ToApiDefinitionInfoOutputWithContext(ctx context.Context) ApiDefinitionInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDefinitionInfoOutput)
}

func (i ApiDefinitionInfoArgs) ToApiDefinitionInfoPtrOutput() ApiDefinitionInfoPtrOutput {
	return i.ToApiDefinitionInfoPtrOutputWithContext(context.Background())
}

func (i ApiDefinitionInfoArgs) ToApiDefinitionInfoPtrOutputWithContext(ctx context.Context) ApiDefinitionInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDefinitionInfoOutput).ToApiDefinitionInfoPtrOutputWithContext(ctx)
}









type ApiDefinitionInfoPtrInput interface {
	pulumi.Input

	ToApiDefinitionInfoPtrOutput() ApiDefinitionInfoPtrOutput
	ToApiDefinitionInfoPtrOutputWithContext(context.Context) ApiDefinitionInfoPtrOutput
}

type apiDefinitionInfoPtrType ApiDefinitionInfoArgs

func ApiDefinitionInfoPtr(v *ApiDefinitionInfoArgs) ApiDefinitionInfoPtrInput {
	return (*apiDefinitionInfoPtrType)(v)
}

func (*apiDefinitionInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDefinitionInfo)(nil)).Elem()
}

func (i *apiDefinitionInfoPtrType) ToApiDefinitionInfoPtrOutput() ApiDefinitionInfoPtrOutput {
	return i.ToApiDefinitionInfoPtrOutputWithContext(context.Background())
}

func (i *apiDefinitionInfoPtrType) ToApiDefinitionInfoPtrOutputWithContext(ctx context.Context) ApiDefinitionInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDefinitionInfoPtrOutput)
}

type ApiDefinitionInfoOutput struct{ *pulumi.OutputState }

func (ApiDefinitionInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiDefinitionInfo)(nil)).Elem()
}

func (o ApiDefinitionInfoOutput) ToApiDefinitionInfoOutput() ApiDefinitionInfoOutput {
	return o
}

func (o ApiDefinitionInfoOutput) ToApiDefinitionInfoOutputWithContext(ctx context.Context) ApiDefinitionInfoOutput {
	return o
}

func (o ApiDefinitionInfoOutput) ToApiDefinitionInfoPtrOutput() ApiDefinitionInfoPtrOutput {
	return o.ToApiDefinitionInfoPtrOutputWithContext(context.Background())
}

func (o ApiDefinitionInfoOutput) ToApiDefinitionInfoPtrOutputWithContext(ctx context.Context) ApiDefinitionInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiDefinitionInfo) *ApiDefinitionInfo {
		return &v
	}).(ApiDefinitionInfoPtrOutput)
}

func (o ApiDefinitionInfoOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiDefinitionInfo) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type ApiDefinitionInfoPtrOutput struct{ *pulumi.OutputState }

func (ApiDefinitionInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDefinitionInfo)(nil)).Elem()
}

func (o ApiDefinitionInfoPtrOutput) ToApiDefinitionInfoPtrOutput() ApiDefinitionInfoPtrOutput {
	return o
}

func (o ApiDefinitionInfoPtrOutput) ToApiDefinitionInfoPtrOutputWithContext(ctx context.Context) ApiDefinitionInfoPtrOutput {
	return o
}

func (o ApiDefinitionInfoPtrOutput) Elem() ApiDefinitionInfoOutput {
	return o.ApplyT(func(v *ApiDefinitionInfo) ApiDefinitionInfo {
		if v != nil {
			return *v
		}
		var ret ApiDefinitionInfo
		return ret
	}).(ApiDefinitionInfoOutput)
}

func (o ApiDefinitionInfoPtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiDefinitionInfo) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type ApiDefinitionInfoResponse struct {
	Url *string `pulumi:"url"`
}

type ApiDefinitionInfoResponseOutput struct{ *pulumi.OutputState }

func (ApiDefinitionInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiDefinitionInfoResponse)(nil)).Elem()
}

func (o ApiDefinitionInfoResponseOutput) ToApiDefinitionInfoResponseOutput() ApiDefinitionInfoResponseOutput {
	return o
}

func (o ApiDefinitionInfoResponseOutput) ToApiDefinitionInfoResponseOutputWithContext(ctx context.Context) ApiDefinitionInfoResponseOutput {
	return o
}

func (o ApiDefinitionInfoResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiDefinitionInfoResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type ApiDefinitionInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiDefinitionInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDefinitionInfoResponse)(nil)).Elem()
}

func (o ApiDefinitionInfoResponsePtrOutput) ToApiDefinitionInfoResponsePtrOutput() ApiDefinitionInfoResponsePtrOutput {
	return o
}

func (o ApiDefinitionInfoResponsePtrOutput) ToApiDefinitionInfoResponsePtrOutputWithContext(ctx context.Context) ApiDefinitionInfoResponsePtrOutput {
	return o
}

func (o ApiDefinitionInfoResponsePtrOutput) Elem() ApiDefinitionInfoResponseOutput {
	return o.ApplyT(func(v *ApiDefinitionInfoResponse) ApiDefinitionInfoResponse {
		if v != nil {
			return *v
		}
		var ret ApiDefinitionInfoResponse
		return ret
	}).(ApiDefinitionInfoResponseOutput)
}

func (o ApiDefinitionInfoResponsePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiDefinitionInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type ApplicationLogsConfig struct {
	AzureBlobStorage  *AzureBlobStorageApplicationLogsConfig  `pulumi:"azureBlobStorage"`
	AzureTableStorage *AzureTableStorageApplicationLogsConfig `pulumi:"azureTableStorage"`
	FileSystem        *FileSystemApplicationLogsConfig        `pulumi:"fileSystem"`
}





type ApplicationLogsConfigInput interface {
	pulumi.Input

	ToApplicationLogsConfigOutput() ApplicationLogsConfigOutput
	ToApplicationLogsConfigOutputWithContext(context.Context) ApplicationLogsConfigOutput
}

type ApplicationLogsConfigArgs struct {
	AzureBlobStorage  AzureBlobStorageApplicationLogsConfigPtrInput  `pulumi:"azureBlobStorage"`
	AzureTableStorage AzureTableStorageApplicationLogsConfigPtrInput `pulumi:"azureTableStorage"`
	FileSystem        FileSystemApplicationLogsConfigPtrInput        `pulumi:"fileSystem"`
}

func (ApplicationLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationLogsConfig)(nil)).Elem()
}

func (i ApplicationLogsConfigArgs) ToApplicationLogsConfigOutput() ApplicationLogsConfigOutput {
	return i.ToApplicationLogsConfigOutputWithContext(context.Background())
}

func (i ApplicationLogsConfigArgs) ToApplicationLogsConfigOutputWithContext(ctx context.Context) ApplicationLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLogsConfigOutput)
}

func (i ApplicationLogsConfigArgs) ToApplicationLogsConfigPtrOutput() ApplicationLogsConfigPtrOutput {
	return i.ToApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i ApplicationLogsConfigArgs) ToApplicationLogsConfigPtrOutputWithContext(ctx context.Context) ApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLogsConfigOutput).ToApplicationLogsConfigPtrOutputWithContext(ctx)
}









type ApplicationLogsConfigPtrInput interface {
	pulumi.Input

	ToApplicationLogsConfigPtrOutput() ApplicationLogsConfigPtrOutput
	ToApplicationLogsConfigPtrOutputWithContext(context.Context) ApplicationLogsConfigPtrOutput
}

type applicationLogsConfigPtrType ApplicationLogsConfigArgs

func ApplicationLogsConfigPtr(v *ApplicationLogsConfigArgs) ApplicationLogsConfigPtrInput {
	return (*applicationLogsConfigPtrType)(v)
}

func (*applicationLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLogsConfig)(nil)).Elem()
}

func (i *applicationLogsConfigPtrType) ToApplicationLogsConfigPtrOutput() ApplicationLogsConfigPtrOutput {
	return i.ToApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i *applicationLogsConfigPtrType) ToApplicationLogsConfigPtrOutputWithContext(ctx context.Context) ApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLogsConfigPtrOutput)
}

type ApplicationLogsConfigOutput struct{ *pulumi.OutputState }

func (ApplicationLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationLogsConfig)(nil)).Elem()
}

func (o ApplicationLogsConfigOutput) ToApplicationLogsConfigOutput() ApplicationLogsConfigOutput {
	return o
}

func (o ApplicationLogsConfigOutput) ToApplicationLogsConfigOutputWithContext(ctx context.Context) ApplicationLogsConfigOutput {
	return o
}

func (o ApplicationLogsConfigOutput) ToApplicationLogsConfigPtrOutput() ApplicationLogsConfigPtrOutput {
	return o.ToApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (o ApplicationLogsConfigOutput) ToApplicationLogsConfigPtrOutputWithContext(ctx context.Context) ApplicationLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationLogsConfig) *ApplicationLogsConfig {
		return &v
	}).(ApplicationLogsConfigPtrOutput)
}

func (o ApplicationLogsConfigOutput) AzureBlobStorage() AzureBlobStorageApplicationLogsConfigPtrOutput {
	return o.ApplyT(func(v ApplicationLogsConfig) *AzureBlobStorageApplicationLogsConfig { return v.AzureBlobStorage }).(AzureBlobStorageApplicationLogsConfigPtrOutput)
}

func (o ApplicationLogsConfigOutput) AzureTableStorage() AzureTableStorageApplicationLogsConfigPtrOutput {
	return o.ApplyT(func(v ApplicationLogsConfig) *AzureTableStorageApplicationLogsConfig { return v.AzureTableStorage }).(AzureTableStorageApplicationLogsConfigPtrOutput)
}

func (o ApplicationLogsConfigOutput) FileSystem() FileSystemApplicationLogsConfigPtrOutput {
	return o.ApplyT(func(v ApplicationLogsConfig) *FileSystemApplicationLogsConfig { return v.FileSystem }).(FileSystemApplicationLogsConfigPtrOutput)
}

type ApplicationLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (ApplicationLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLogsConfig)(nil)).Elem()
}

func (o ApplicationLogsConfigPtrOutput) ToApplicationLogsConfigPtrOutput() ApplicationLogsConfigPtrOutput {
	return o
}

func (o ApplicationLogsConfigPtrOutput) ToApplicationLogsConfigPtrOutputWithContext(ctx context.Context) ApplicationLogsConfigPtrOutput {
	return o
}

func (o ApplicationLogsConfigPtrOutput) Elem() ApplicationLogsConfigOutput {
	return o.ApplyT(func(v *ApplicationLogsConfig) ApplicationLogsConfig {
		if v != nil {
			return *v
		}
		var ret ApplicationLogsConfig
		return ret
	}).(ApplicationLogsConfigOutput)
}

func (o ApplicationLogsConfigPtrOutput) AzureBlobStorage() AzureBlobStorageApplicationLogsConfigPtrOutput {
	return o.ApplyT(func(v *ApplicationLogsConfig) *AzureBlobStorageApplicationLogsConfig {
		if v == nil {
			return nil
		}
		return v.AzureBlobStorage
	}).(AzureBlobStorageApplicationLogsConfigPtrOutput)
}

func (o ApplicationLogsConfigPtrOutput) AzureTableStorage() AzureTableStorageApplicationLogsConfigPtrOutput {
	return o.ApplyT(func(v *ApplicationLogsConfig) *AzureTableStorageApplicationLogsConfig {
		if v == nil {
			return nil
		}
		return v.AzureTableStorage
	}).(AzureTableStorageApplicationLogsConfigPtrOutput)
}

func (o ApplicationLogsConfigPtrOutput) FileSystem() FileSystemApplicationLogsConfigPtrOutput {
	return o.ApplyT(func(v *ApplicationLogsConfig) *FileSystemApplicationLogsConfig {
		if v == nil {
			return nil
		}
		return v.FileSystem
	}).(FileSystemApplicationLogsConfigPtrOutput)
}

type ApplicationLogsConfigResponse struct {
	AzureBlobStorage  *AzureBlobStorageApplicationLogsConfigResponse  `pulumi:"azureBlobStorage"`
	AzureTableStorage *AzureTableStorageApplicationLogsConfigResponse `pulumi:"azureTableStorage"`
	FileSystem        *FileSystemApplicationLogsConfigResponse        `pulumi:"fileSystem"`
}

type ApplicationLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (ApplicationLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationLogsConfigResponse)(nil)).Elem()
}

func (o ApplicationLogsConfigResponseOutput) ToApplicationLogsConfigResponseOutput() ApplicationLogsConfigResponseOutput {
	return o
}

func (o ApplicationLogsConfigResponseOutput) ToApplicationLogsConfigResponseOutputWithContext(ctx context.Context) ApplicationLogsConfigResponseOutput {
	return o
}

func (o ApplicationLogsConfigResponseOutput) AzureBlobStorage() AzureBlobStorageApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v ApplicationLogsConfigResponse) *AzureBlobStorageApplicationLogsConfigResponse {
		return v.AzureBlobStorage
	}).(AzureBlobStorageApplicationLogsConfigResponsePtrOutput)
}

func (o ApplicationLogsConfigResponseOutput) AzureTableStorage() AzureTableStorageApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v ApplicationLogsConfigResponse) *AzureTableStorageApplicationLogsConfigResponse {
		return v.AzureTableStorage
	}).(AzureTableStorageApplicationLogsConfigResponsePtrOutput)
}

func (o ApplicationLogsConfigResponseOutput) FileSystem() FileSystemApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v ApplicationLogsConfigResponse) *FileSystemApplicationLogsConfigResponse { return v.FileSystem }).(FileSystemApplicationLogsConfigResponsePtrOutput)
}

type ApplicationLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLogsConfigResponse)(nil)).Elem()
}

func (o ApplicationLogsConfigResponsePtrOutput) ToApplicationLogsConfigResponsePtrOutput() ApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o ApplicationLogsConfigResponsePtrOutput) ToApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) ApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o ApplicationLogsConfigResponsePtrOutput) Elem() ApplicationLogsConfigResponseOutput {
	return o.ApplyT(func(v *ApplicationLogsConfigResponse) ApplicationLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationLogsConfigResponse
		return ret
	}).(ApplicationLogsConfigResponseOutput)
}

func (o ApplicationLogsConfigResponsePtrOutput) AzureBlobStorage() AzureBlobStorageApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationLogsConfigResponse) *AzureBlobStorageApplicationLogsConfigResponse {
		if v == nil {
			return nil
		}
		return v.AzureBlobStorage
	}).(AzureBlobStorageApplicationLogsConfigResponsePtrOutput)
}

func (o ApplicationLogsConfigResponsePtrOutput) AzureTableStorage() AzureTableStorageApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationLogsConfigResponse) *AzureTableStorageApplicationLogsConfigResponse {
		if v == nil {
			return nil
		}
		return v.AzureTableStorage
	}).(AzureTableStorageApplicationLogsConfigResponsePtrOutput)
}

func (o ApplicationLogsConfigResponsePtrOutput) FileSystem() FileSystemApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationLogsConfigResponse) *FileSystemApplicationLogsConfigResponse {
		if v == nil {
			return nil
		}
		return v.FileSystem
	}).(FileSystemApplicationLogsConfigResponsePtrOutput)
}

type AutoHealActions struct {
	ActionType              AutoHealActionType    `pulumi:"actionType"`
	CustomAction            *AutoHealCustomAction `pulumi:"customAction"`
	MinProcessExecutionTime *string               `pulumi:"minProcessExecutionTime"`
}





type AutoHealActionsInput interface {
	pulumi.Input

	ToAutoHealActionsOutput() AutoHealActionsOutput
	ToAutoHealActionsOutputWithContext(context.Context) AutoHealActionsOutput
}

type AutoHealActionsArgs struct {
	ActionType              AutoHealActionTypeInput      `pulumi:"actionType"`
	CustomAction            AutoHealCustomActionPtrInput `pulumi:"customAction"`
	MinProcessExecutionTime pulumi.StringPtrInput        `pulumi:"minProcessExecutionTime"`
}

func (AutoHealActionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealActions)(nil)).Elem()
}

func (i AutoHealActionsArgs) ToAutoHealActionsOutput() AutoHealActionsOutput {
	return i.ToAutoHealActionsOutputWithContext(context.Background())
}

func (i AutoHealActionsArgs) ToAutoHealActionsOutputWithContext(ctx context.Context) AutoHealActionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealActionsOutput)
}

func (i AutoHealActionsArgs) ToAutoHealActionsPtrOutput() AutoHealActionsPtrOutput {
	return i.ToAutoHealActionsPtrOutputWithContext(context.Background())
}

func (i AutoHealActionsArgs) ToAutoHealActionsPtrOutputWithContext(ctx context.Context) AutoHealActionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealActionsOutput).ToAutoHealActionsPtrOutputWithContext(ctx)
}









type AutoHealActionsPtrInput interface {
	pulumi.Input

	ToAutoHealActionsPtrOutput() AutoHealActionsPtrOutput
	ToAutoHealActionsPtrOutputWithContext(context.Context) AutoHealActionsPtrOutput
}

type autoHealActionsPtrType AutoHealActionsArgs

func AutoHealActionsPtr(v *AutoHealActionsArgs) AutoHealActionsPtrInput {
	return (*autoHealActionsPtrType)(v)
}

func (*autoHealActionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealActions)(nil)).Elem()
}

func (i *autoHealActionsPtrType) ToAutoHealActionsPtrOutput() AutoHealActionsPtrOutput {
	return i.ToAutoHealActionsPtrOutputWithContext(context.Background())
}

func (i *autoHealActionsPtrType) ToAutoHealActionsPtrOutputWithContext(ctx context.Context) AutoHealActionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealActionsPtrOutput)
}

type AutoHealActionsOutput struct{ *pulumi.OutputState }

func (AutoHealActionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealActions)(nil)).Elem()
}

func (o AutoHealActionsOutput) ToAutoHealActionsOutput() AutoHealActionsOutput {
	return o
}

func (o AutoHealActionsOutput) ToAutoHealActionsOutputWithContext(ctx context.Context) AutoHealActionsOutput {
	return o
}

func (o AutoHealActionsOutput) ToAutoHealActionsPtrOutput() AutoHealActionsPtrOutput {
	return o.ToAutoHealActionsPtrOutputWithContext(context.Background())
}

func (o AutoHealActionsOutput) ToAutoHealActionsPtrOutputWithContext(ctx context.Context) AutoHealActionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealActions) *AutoHealActions {
		return &v
	}).(AutoHealActionsPtrOutput)
}

func (o AutoHealActionsOutput) ActionType() AutoHealActionTypeOutput {
	return o.ApplyT(func(v AutoHealActions) AutoHealActionType { return v.ActionType }).(AutoHealActionTypeOutput)
}

func (o AutoHealActionsOutput) CustomAction() AutoHealCustomActionPtrOutput {
	return o.ApplyT(func(v AutoHealActions) *AutoHealCustomAction { return v.CustomAction }).(AutoHealCustomActionPtrOutput)
}

func (o AutoHealActionsOutput) MinProcessExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealActions) *string { return v.MinProcessExecutionTime }).(pulumi.StringPtrOutput)
}

type AutoHealActionsPtrOutput struct{ *pulumi.OutputState }

func (AutoHealActionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealActions)(nil)).Elem()
}

func (o AutoHealActionsPtrOutput) ToAutoHealActionsPtrOutput() AutoHealActionsPtrOutput {
	return o
}

func (o AutoHealActionsPtrOutput) ToAutoHealActionsPtrOutputWithContext(ctx context.Context) AutoHealActionsPtrOutput {
	return o
}

func (o AutoHealActionsPtrOutput) Elem() AutoHealActionsOutput {
	return o.ApplyT(func(v *AutoHealActions) AutoHealActions {
		if v != nil {
			return *v
		}
		var ret AutoHealActions
		return ret
	}).(AutoHealActionsOutput)
}

func (o AutoHealActionsPtrOutput) ActionType() AutoHealActionTypePtrOutput {
	return o.ApplyT(func(v *AutoHealActions) *AutoHealActionType {
		if v == nil {
			return nil
		}
		return &v.ActionType
	}).(AutoHealActionTypePtrOutput)
}

func (o AutoHealActionsPtrOutput) CustomAction() AutoHealCustomActionPtrOutput {
	return o.ApplyT(func(v *AutoHealActions) *AutoHealCustomAction {
		if v == nil {
			return nil
		}
		return v.CustomAction
	}).(AutoHealCustomActionPtrOutput)
}

func (o AutoHealActionsPtrOutput) MinProcessExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealActions) *string {
		if v == nil {
			return nil
		}
		return v.MinProcessExecutionTime
	}).(pulumi.StringPtrOutput)
}

type AutoHealActionsResponse struct {
	ActionType              string                        `pulumi:"actionType"`
	CustomAction            *AutoHealCustomActionResponse `pulumi:"customAction"`
	MinProcessExecutionTime *string                       `pulumi:"minProcessExecutionTime"`
}

type AutoHealActionsResponseOutput struct{ *pulumi.OutputState }

func (AutoHealActionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealActionsResponse)(nil)).Elem()
}

func (o AutoHealActionsResponseOutput) ToAutoHealActionsResponseOutput() AutoHealActionsResponseOutput {
	return o
}

func (o AutoHealActionsResponseOutput) ToAutoHealActionsResponseOutputWithContext(ctx context.Context) AutoHealActionsResponseOutput {
	return o
}

func (o AutoHealActionsResponseOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutoHealActionsResponse) string { return v.ActionType }).(pulumi.StringOutput)
}

func (o AutoHealActionsResponseOutput) CustomAction() AutoHealCustomActionResponsePtrOutput {
	return o.ApplyT(func(v AutoHealActionsResponse) *AutoHealCustomActionResponse { return v.CustomAction }).(AutoHealCustomActionResponsePtrOutput)
}

func (o AutoHealActionsResponseOutput) MinProcessExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealActionsResponse) *string { return v.MinProcessExecutionTime }).(pulumi.StringPtrOutput)
}

type AutoHealActionsResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoHealActionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealActionsResponse)(nil)).Elem()
}

func (o AutoHealActionsResponsePtrOutput) ToAutoHealActionsResponsePtrOutput() AutoHealActionsResponsePtrOutput {
	return o
}

func (o AutoHealActionsResponsePtrOutput) ToAutoHealActionsResponsePtrOutputWithContext(ctx context.Context) AutoHealActionsResponsePtrOutput {
	return o
}

func (o AutoHealActionsResponsePtrOutput) Elem() AutoHealActionsResponseOutput {
	return o.ApplyT(func(v *AutoHealActionsResponse) AutoHealActionsResponse {
		if v != nil {
			return *v
		}
		var ret AutoHealActionsResponse
		return ret
	}).(AutoHealActionsResponseOutput)
}

func (o AutoHealActionsResponsePtrOutput) ActionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealActionsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActionType
	}).(pulumi.StringPtrOutput)
}

func (o AutoHealActionsResponsePtrOutput) CustomAction() AutoHealCustomActionResponsePtrOutput {
	return o.ApplyT(func(v *AutoHealActionsResponse) *AutoHealCustomActionResponse {
		if v == nil {
			return nil
		}
		return v.CustomAction
	}).(AutoHealCustomActionResponsePtrOutput)
}

func (o AutoHealActionsResponsePtrOutput) MinProcessExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealActionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.MinProcessExecutionTime
	}).(pulumi.StringPtrOutput)
}

type AutoHealCustomAction struct {
	Exe        *string `pulumi:"exe"`
	Parameters *string `pulumi:"parameters"`
}





type AutoHealCustomActionInput interface {
	pulumi.Input

	ToAutoHealCustomActionOutput() AutoHealCustomActionOutput
	ToAutoHealCustomActionOutputWithContext(context.Context) AutoHealCustomActionOutput
}

type AutoHealCustomActionArgs struct {
	Exe        pulumi.StringPtrInput `pulumi:"exe"`
	Parameters pulumi.StringPtrInput `pulumi:"parameters"`
}

func (AutoHealCustomActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealCustomAction)(nil)).Elem()
}

func (i AutoHealCustomActionArgs) ToAutoHealCustomActionOutput() AutoHealCustomActionOutput {
	return i.ToAutoHealCustomActionOutputWithContext(context.Background())
}

func (i AutoHealCustomActionArgs) ToAutoHealCustomActionOutputWithContext(ctx context.Context) AutoHealCustomActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealCustomActionOutput)
}

func (i AutoHealCustomActionArgs) ToAutoHealCustomActionPtrOutput() AutoHealCustomActionPtrOutput {
	return i.ToAutoHealCustomActionPtrOutputWithContext(context.Background())
}

func (i AutoHealCustomActionArgs) ToAutoHealCustomActionPtrOutputWithContext(ctx context.Context) AutoHealCustomActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealCustomActionOutput).ToAutoHealCustomActionPtrOutputWithContext(ctx)
}









type AutoHealCustomActionPtrInput interface {
	pulumi.Input

	ToAutoHealCustomActionPtrOutput() AutoHealCustomActionPtrOutput
	ToAutoHealCustomActionPtrOutputWithContext(context.Context) AutoHealCustomActionPtrOutput
}

type autoHealCustomActionPtrType AutoHealCustomActionArgs

func AutoHealCustomActionPtr(v *AutoHealCustomActionArgs) AutoHealCustomActionPtrInput {
	return (*autoHealCustomActionPtrType)(v)
}

func (*autoHealCustomActionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealCustomAction)(nil)).Elem()
}

func (i *autoHealCustomActionPtrType) ToAutoHealCustomActionPtrOutput() AutoHealCustomActionPtrOutput {
	return i.ToAutoHealCustomActionPtrOutputWithContext(context.Background())
}

func (i *autoHealCustomActionPtrType) ToAutoHealCustomActionPtrOutputWithContext(ctx context.Context) AutoHealCustomActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealCustomActionPtrOutput)
}

type AutoHealCustomActionOutput struct{ *pulumi.OutputState }

func (AutoHealCustomActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealCustomAction)(nil)).Elem()
}

func (o AutoHealCustomActionOutput) ToAutoHealCustomActionOutput() AutoHealCustomActionOutput {
	return o
}

func (o AutoHealCustomActionOutput) ToAutoHealCustomActionOutputWithContext(ctx context.Context) AutoHealCustomActionOutput {
	return o
}

func (o AutoHealCustomActionOutput) ToAutoHealCustomActionPtrOutput() AutoHealCustomActionPtrOutput {
	return o.ToAutoHealCustomActionPtrOutputWithContext(context.Background())
}

func (o AutoHealCustomActionOutput) ToAutoHealCustomActionPtrOutputWithContext(ctx context.Context) AutoHealCustomActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealCustomAction) *AutoHealCustomAction {
		return &v
	}).(AutoHealCustomActionPtrOutput)
}

func (o AutoHealCustomActionOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealCustomAction) *string { return v.Exe }).(pulumi.StringPtrOutput)
}

func (o AutoHealCustomActionOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealCustomAction) *string { return v.Parameters }).(pulumi.StringPtrOutput)
}

type AutoHealCustomActionPtrOutput struct{ *pulumi.OutputState }

func (AutoHealCustomActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealCustomAction)(nil)).Elem()
}

func (o AutoHealCustomActionPtrOutput) ToAutoHealCustomActionPtrOutput() AutoHealCustomActionPtrOutput {
	return o
}

func (o AutoHealCustomActionPtrOutput) ToAutoHealCustomActionPtrOutputWithContext(ctx context.Context) AutoHealCustomActionPtrOutput {
	return o
}

func (o AutoHealCustomActionPtrOutput) Elem() AutoHealCustomActionOutput {
	return o.ApplyT(func(v *AutoHealCustomAction) AutoHealCustomAction {
		if v != nil {
			return *v
		}
		var ret AutoHealCustomAction
		return ret
	}).(AutoHealCustomActionOutput)
}

func (o AutoHealCustomActionPtrOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealCustomAction) *string {
		if v == nil {
			return nil
		}
		return v.Exe
	}).(pulumi.StringPtrOutput)
}

func (o AutoHealCustomActionPtrOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealCustomAction) *string {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.StringPtrOutput)
}

type AutoHealCustomActionResponse struct {
	Exe        *string `pulumi:"exe"`
	Parameters *string `pulumi:"parameters"`
}

type AutoHealCustomActionResponseOutput struct{ *pulumi.OutputState }

func (AutoHealCustomActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealCustomActionResponse)(nil)).Elem()
}

func (o AutoHealCustomActionResponseOutput) ToAutoHealCustomActionResponseOutput() AutoHealCustomActionResponseOutput {
	return o
}

func (o AutoHealCustomActionResponseOutput) ToAutoHealCustomActionResponseOutputWithContext(ctx context.Context) AutoHealCustomActionResponseOutput {
	return o
}

func (o AutoHealCustomActionResponseOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealCustomActionResponse) *string { return v.Exe }).(pulumi.StringPtrOutput)
}

func (o AutoHealCustomActionResponseOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealCustomActionResponse) *string { return v.Parameters }).(pulumi.StringPtrOutput)
}

type AutoHealCustomActionResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoHealCustomActionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealCustomActionResponse)(nil)).Elem()
}

func (o AutoHealCustomActionResponsePtrOutput) ToAutoHealCustomActionResponsePtrOutput() AutoHealCustomActionResponsePtrOutput {
	return o
}

func (o AutoHealCustomActionResponsePtrOutput) ToAutoHealCustomActionResponsePtrOutputWithContext(ctx context.Context) AutoHealCustomActionResponsePtrOutput {
	return o
}

func (o AutoHealCustomActionResponsePtrOutput) Elem() AutoHealCustomActionResponseOutput {
	return o.ApplyT(func(v *AutoHealCustomActionResponse) AutoHealCustomActionResponse {
		if v != nil {
			return *v
		}
		var ret AutoHealCustomActionResponse
		return ret
	}).(AutoHealCustomActionResponseOutput)
}

func (o AutoHealCustomActionResponsePtrOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealCustomActionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Exe
	}).(pulumi.StringPtrOutput)
}

func (o AutoHealCustomActionResponsePtrOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealCustomActionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.StringPtrOutput)
}

type AutoHealRules struct {
	Actions  *AutoHealActions  `pulumi:"actions"`
	Triggers *AutoHealTriggers `pulumi:"triggers"`
}





type AutoHealRulesInput interface {
	pulumi.Input

	ToAutoHealRulesOutput() AutoHealRulesOutput
	ToAutoHealRulesOutputWithContext(context.Context) AutoHealRulesOutput
}

type AutoHealRulesArgs struct {
	Actions  AutoHealActionsPtrInput  `pulumi:"actions"`
	Triggers AutoHealTriggersPtrInput `pulumi:"triggers"`
}

func (AutoHealRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealRules)(nil)).Elem()
}

func (i AutoHealRulesArgs) ToAutoHealRulesOutput() AutoHealRulesOutput {
	return i.ToAutoHealRulesOutputWithContext(context.Background())
}

func (i AutoHealRulesArgs) ToAutoHealRulesOutputWithContext(ctx context.Context) AutoHealRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealRulesOutput)
}

func (i AutoHealRulesArgs) ToAutoHealRulesPtrOutput() AutoHealRulesPtrOutput {
	return i.ToAutoHealRulesPtrOutputWithContext(context.Background())
}

func (i AutoHealRulesArgs) ToAutoHealRulesPtrOutputWithContext(ctx context.Context) AutoHealRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealRulesOutput).ToAutoHealRulesPtrOutputWithContext(ctx)
}









type AutoHealRulesPtrInput interface {
	pulumi.Input

	ToAutoHealRulesPtrOutput() AutoHealRulesPtrOutput
	ToAutoHealRulesPtrOutputWithContext(context.Context) AutoHealRulesPtrOutput
}

type autoHealRulesPtrType AutoHealRulesArgs

func AutoHealRulesPtr(v *AutoHealRulesArgs) AutoHealRulesPtrInput {
	return (*autoHealRulesPtrType)(v)
}

func (*autoHealRulesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealRules)(nil)).Elem()
}

func (i *autoHealRulesPtrType) ToAutoHealRulesPtrOutput() AutoHealRulesPtrOutput {
	return i.ToAutoHealRulesPtrOutputWithContext(context.Background())
}

func (i *autoHealRulesPtrType) ToAutoHealRulesPtrOutputWithContext(ctx context.Context) AutoHealRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealRulesPtrOutput)
}

type AutoHealRulesOutput struct{ *pulumi.OutputState }

func (AutoHealRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealRules)(nil)).Elem()
}

func (o AutoHealRulesOutput) ToAutoHealRulesOutput() AutoHealRulesOutput {
	return o
}

func (o AutoHealRulesOutput) ToAutoHealRulesOutputWithContext(ctx context.Context) AutoHealRulesOutput {
	return o
}

func (o AutoHealRulesOutput) ToAutoHealRulesPtrOutput() AutoHealRulesPtrOutput {
	return o.ToAutoHealRulesPtrOutputWithContext(context.Background())
}

func (o AutoHealRulesOutput) ToAutoHealRulesPtrOutputWithContext(ctx context.Context) AutoHealRulesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealRules) *AutoHealRules {
		return &v
	}).(AutoHealRulesPtrOutput)
}

func (o AutoHealRulesOutput) Actions() AutoHealActionsPtrOutput {
	return o.ApplyT(func(v AutoHealRules) *AutoHealActions { return v.Actions }).(AutoHealActionsPtrOutput)
}

func (o AutoHealRulesOutput) Triggers() AutoHealTriggersPtrOutput {
	return o.ApplyT(func(v AutoHealRules) *AutoHealTriggers { return v.Triggers }).(AutoHealTriggersPtrOutput)
}

type AutoHealRulesPtrOutput struct{ *pulumi.OutputState }

func (AutoHealRulesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealRules)(nil)).Elem()
}

func (o AutoHealRulesPtrOutput) ToAutoHealRulesPtrOutput() AutoHealRulesPtrOutput {
	return o
}

func (o AutoHealRulesPtrOutput) ToAutoHealRulesPtrOutputWithContext(ctx context.Context) AutoHealRulesPtrOutput {
	return o
}

func (o AutoHealRulesPtrOutput) Elem() AutoHealRulesOutput {
	return o.ApplyT(func(v *AutoHealRules) AutoHealRules {
		if v != nil {
			return *v
		}
		var ret AutoHealRules
		return ret
	}).(AutoHealRulesOutput)
}

func (o AutoHealRulesPtrOutput) Actions() AutoHealActionsPtrOutput {
	return o.ApplyT(func(v *AutoHealRules) *AutoHealActions {
		if v == nil {
			return nil
		}
		return v.Actions
	}).(AutoHealActionsPtrOutput)
}

func (o AutoHealRulesPtrOutput) Triggers() AutoHealTriggersPtrOutput {
	return o.ApplyT(func(v *AutoHealRules) *AutoHealTriggers {
		if v == nil {
			return nil
		}
		return v.Triggers
	}).(AutoHealTriggersPtrOutput)
}

type AutoHealRulesResponse struct {
	Actions  *AutoHealActionsResponse  `pulumi:"actions"`
	Triggers *AutoHealTriggersResponse `pulumi:"triggers"`
}

type AutoHealRulesResponseOutput struct{ *pulumi.OutputState }

func (AutoHealRulesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealRulesResponse)(nil)).Elem()
}

func (o AutoHealRulesResponseOutput) ToAutoHealRulesResponseOutput() AutoHealRulesResponseOutput {
	return o
}

func (o AutoHealRulesResponseOutput) ToAutoHealRulesResponseOutputWithContext(ctx context.Context) AutoHealRulesResponseOutput {
	return o
}

func (o AutoHealRulesResponseOutput) Actions() AutoHealActionsResponsePtrOutput {
	return o.ApplyT(func(v AutoHealRulesResponse) *AutoHealActionsResponse { return v.Actions }).(AutoHealActionsResponsePtrOutput)
}

func (o AutoHealRulesResponseOutput) Triggers() AutoHealTriggersResponsePtrOutput {
	return o.ApplyT(func(v AutoHealRulesResponse) *AutoHealTriggersResponse { return v.Triggers }).(AutoHealTriggersResponsePtrOutput)
}

type AutoHealRulesResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoHealRulesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealRulesResponse)(nil)).Elem()
}

func (o AutoHealRulesResponsePtrOutput) ToAutoHealRulesResponsePtrOutput() AutoHealRulesResponsePtrOutput {
	return o
}

func (o AutoHealRulesResponsePtrOutput) ToAutoHealRulesResponsePtrOutputWithContext(ctx context.Context) AutoHealRulesResponsePtrOutput {
	return o
}

func (o AutoHealRulesResponsePtrOutput) Elem() AutoHealRulesResponseOutput {
	return o.ApplyT(func(v *AutoHealRulesResponse) AutoHealRulesResponse {
		if v != nil {
			return *v
		}
		var ret AutoHealRulesResponse
		return ret
	}).(AutoHealRulesResponseOutput)
}

func (o AutoHealRulesResponsePtrOutput) Actions() AutoHealActionsResponsePtrOutput {
	return o.ApplyT(func(v *AutoHealRulesResponse) *AutoHealActionsResponse {
		if v == nil {
			return nil
		}
		return v.Actions
	}).(AutoHealActionsResponsePtrOutput)
}

func (o AutoHealRulesResponsePtrOutput) Triggers() AutoHealTriggersResponsePtrOutput {
	return o.ApplyT(func(v *AutoHealRulesResponse) *AutoHealTriggersResponse {
		if v == nil {
			return nil
		}
		return v.Triggers
	}).(AutoHealTriggersResponsePtrOutput)
}

type AutoHealTriggers struct {
	PrivateBytesInKB *int                      `pulumi:"privateBytesInKB"`
	Requests         *RequestsBasedTrigger     `pulumi:"requests"`
	SlowRequests     *SlowRequestsBasedTrigger `pulumi:"slowRequests"`
	StatusCodes      []StatusCodesBasedTrigger `pulumi:"statusCodes"`
}





type AutoHealTriggersInput interface {
	pulumi.Input

	ToAutoHealTriggersOutput() AutoHealTriggersOutput
	ToAutoHealTriggersOutputWithContext(context.Context) AutoHealTriggersOutput
}

type AutoHealTriggersArgs struct {
	PrivateBytesInKB pulumi.IntPtrInput                `pulumi:"privateBytesInKB"`
	Requests         RequestsBasedTriggerPtrInput      `pulumi:"requests"`
	SlowRequests     SlowRequestsBasedTriggerPtrInput  `pulumi:"slowRequests"`
	StatusCodes      StatusCodesBasedTriggerArrayInput `pulumi:"statusCodes"`
}

func (AutoHealTriggersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealTriggers)(nil)).Elem()
}

func (i AutoHealTriggersArgs) ToAutoHealTriggersOutput() AutoHealTriggersOutput {
	return i.ToAutoHealTriggersOutputWithContext(context.Background())
}

func (i AutoHealTriggersArgs) ToAutoHealTriggersOutputWithContext(ctx context.Context) AutoHealTriggersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealTriggersOutput)
}

func (i AutoHealTriggersArgs) ToAutoHealTriggersPtrOutput() AutoHealTriggersPtrOutput {
	return i.ToAutoHealTriggersPtrOutputWithContext(context.Background())
}

func (i AutoHealTriggersArgs) ToAutoHealTriggersPtrOutputWithContext(ctx context.Context) AutoHealTriggersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealTriggersOutput).ToAutoHealTriggersPtrOutputWithContext(ctx)
}









type AutoHealTriggersPtrInput interface {
	pulumi.Input

	ToAutoHealTriggersPtrOutput() AutoHealTriggersPtrOutput
	ToAutoHealTriggersPtrOutputWithContext(context.Context) AutoHealTriggersPtrOutput
}

type autoHealTriggersPtrType AutoHealTriggersArgs

func AutoHealTriggersPtr(v *AutoHealTriggersArgs) AutoHealTriggersPtrInput {
	return (*autoHealTriggersPtrType)(v)
}

func (*autoHealTriggersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealTriggers)(nil)).Elem()
}

func (i *autoHealTriggersPtrType) ToAutoHealTriggersPtrOutput() AutoHealTriggersPtrOutput {
	return i.ToAutoHealTriggersPtrOutputWithContext(context.Background())
}

func (i *autoHealTriggersPtrType) ToAutoHealTriggersPtrOutputWithContext(ctx context.Context) AutoHealTriggersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealTriggersPtrOutput)
}

type AutoHealTriggersOutput struct{ *pulumi.OutputState }

func (AutoHealTriggersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealTriggers)(nil)).Elem()
}

func (o AutoHealTriggersOutput) ToAutoHealTriggersOutput() AutoHealTriggersOutput {
	return o
}

func (o AutoHealTriggersOutput) ToAutoHealTriggersOutputWithContext(ctx context.Context) AutoHealTriggersOutput {
	return o
}

func (o AutoHealTriggersOutput) ToAutoHealTriggersPtrOutput() AutoHealTriggersPtrOutput {
	return o.ToAutoHealTriggersPtrOutputWithContext(context.Background())
}

func (o AutoHealTriggersOutput) ToAutoHealTriggersPtrOutputWithContext(ctx context.Context) AutoHealTriggersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealTriggers) *AutoHealTriggers {
		return &v
	}).(AutoHealTriggersPtrOutput)
}

func (o AutoHealTriggersOutput) PrivateBytesInKB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoHealTriggers) *int { return v.PrivateBytesInKB }).(pulumi.IntPtrOutput)
}

func (o AutoHealTriggersOutput) Requests() RequestsBasedTriggerPtrOutput {
	return o.ApplyT(func(v AutoHealTriggers) *RequestsBasedTrigger { return v.Requests }).(RequestsBasedTriggerPtrOutput)
}

func (o AutoHealTriggersOutput) SlowRequests() SlowRequestsBasedTriggerPtrOutput {
	return o.ApplyT(func(v AutoHealTriggers) *SlowRequestsBasedTrigger { return v.SlowRequests }).(SlowRequestsBasedTriggerPtrOutput)
}

func (o AutoHealTriggersOutput) StatusCodes() StatusCodesBasedTriggerArrayOutput {
	return o.ApplyT(func(v AutoHealTriggers) []StatusCodesBasedTrigger { return v.StatusCodes }).(StatusCodesBasedTriggerArrayOutput)
}

type AutoHealTriggersPtrOutput struct{ *pulumi.OutputState }

func (AutoHealTriggersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealTriggers)(nil)).Elem()
}

func (o AutoHealTriggersPtrOutput) ToAutoHealTriggersPtrOutput() AutoHealTriggersPtrOutput {
	return o
}

func (o AutoHealTriggersPtrOutput) ToAutoHealTriggersPtrOutputWithContext(ctx context.Context) AutoHealTriggersPtrOutput {
	return o
}

func (o AutoHealTriggersPtrOutput) Elem() AutoHealTriggersOutput {
	return o.ApplyT(func(v *AutoHealTriggers) AutoHealTriggers {
		if v != nil {
			return *v
		}
		var ret AutoHealTriggers
		return ret
	}).(AutoHealTriggersOutput)
}

func (o AutoHealTriggersPtrOutput) PrivateBytesInKB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoHealTriggers) *int {
		if v == nil {
			return nil
		}
		return v.PrivateBytesInKB
	}).(pulumi.IntPtrOutput)
}

func (o AutoHealTriggersPtrOutput) Requests() RequestsBasedTriggerPtrOutput {
	return o.ApplyT(func(v *AutoHealTriggers) *RequestsBasedTrigger {
		if v == nil {
			return nil
		}
		return v.Requests
	}).(RequestsBasedTriggerPtrOutput)
}

func (o AutoHealTriggersPtrOutput) SlowRequests() SlowRequestsBasedTriggerPtrOutput {
	return o.ApplyT(func(v *AutoHealTriggers) *SlowRequestsBasedTrigger {
		if v == nil {
			return nil
		}
		return v.SlowRequests
	}).(SlowRequestsBasedTriggerPtrOutput)
}

func (o AutoHealTriggersPtrOutput) StatusCodes() StatusCodesBasedTriggerArrayOutput {
	return o.ApplyT(func(v *AutoHealTriggers) []StatusCodesBasedTrigger {
		if v == nil {
			return nil
		}
		return v.StatusCodes
	}).(StatusCodesBasedTriggerArrayOutput)
}

type AutoHealTriggersResponse struct {
	PrivateBytesInKB *int                              `pulumi:"privateBytesInKB"`
	Requests         *RequestsBasedTriggerResponse     `pulumi:"requests"`
	SlowRequests     *SlowRequestsBasedTriggerResponse `pulumi:"slowRequests"`
	StatusCodes      []StatusCodesBasedTriggerResponse `pulumi:"statusCodes"`
}

type AutoHealTriggersResponseOutput struct{ *pulumi.OutputState }

func (AutoHealTriggersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealTriggersResponse)(nil)).Elem()
}

func (o AutoHealTriggersResponseOutput) ToAutoHealTriggersResponseOutput() AutoHealTriggersResponseOutput {
	return o
}

func (o AutoHealTriggersResponseOutput) ToAutoHealTriggersResponseOutputWithContext(ctx context.Context) AutoHealTriggersResponseOutput {
	return o
}

func (o AutoHealTriggersResponseOutput) PrivateBytesInKB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoHealTriggersResponse) *int { return v.PrivateBytesInKB }).(pulumi.IntPtrOutput)
}

func (o AutoHealTriggersResponseOutput) Requests() RequestsBasedTriggerResponsePtrOutput {
	return o.ApplyT(func(v AutoHealTriggersResponse) *RequestsBasedTriggerResponse { return v.Requests }).(RequestsBasedTriggerResponsePtrOutput)
}

func (o AutoHealTriggersResponseOutput) SlowRequests() SlowRequestsBasedTriggerResponsePtrOutput {
	return o.ApplyT(func(v AutoHealTriggersResponse) *SlowRequestsBasedTriggerResponse { return v.SlowRequests }).(SlowRequestsBasedTriggerResponsePtrOutput)
}

func (o AutoHealTriggersResponseOutput) StatusCodes() StatusCodesBasedTriggerResponseArrayOutput {
	return o.ApplyT(func(v AutoHealTriggersResponse) []StatusCodesBasedTriggerResponse { return v.StatusCodes }).(StatusCodesBasedTriggerResponseArrayOutput)
}

type AutoHealTriggersResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoHealTriggersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealTriggersResponse)(nil)).Elem()
}

func (o AutoHealTriggersResponsePtrOutput) ToAutoHealTriggersResponsePtrOutput() AutoHealTriggersResponsePtrOutput {
	return o
}

func (o AutoHealTriggersResponsePtrOutput) ToAutoHealTriggersResponsePtrOutputWithContext(ctx context.Context) AutoHealTriggersResponsePtrOutput {
	return o
}

func (o AutoHealTriggersResponsePtrOutput) Elem() AutoHealTriggersResponseOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) AutoHealTriggersResponse {
		if v != nil {
			return *v
		}
		var ret AutoHealTriggersResponse
		return ret
	}).(AutoHealTriggersResponseOutput)
}

func (o AutoHealTriggersResponsePtrOutput) PrivateBytesInKB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) *int {
		if v == nil {
			return nil
		}
		return v.PrivateBytesInKB
	}).(pulumi.IntPtrOutput)
}

func (o AutoHealTriggersResponsePtrOutput) Requests() RequestsBasedTriggerResponsePtrOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) *RequestsBasedTriggerResponse {
		if v == nil {
			return nil
		}
		return v.Requests
	}).(RequestsBasedTriggerResponsePtrOutput)
}

func (o AutoHealTriggersResponsePtrOutput) SlowRequests() SlowRequestsBasedTriggerResponsePtrOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) *SlowRequestsBasedTriggerResponse {
		if v == nil {
			return nil
		}
		return v.SlowRequests
	}).(SlowRequestsBasedTriggerResponsePtrOutput)
}

func (o AutoHealTriggersResponsePtrOutput) StatusCodes() StatusCodesBasedTriggerResponseArrayOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) []StatusCodesBasedTriggerResponse {
		if v == nil {
			return nil
		}
		return v.StatusCodes
	}).(StatusCodesBasedTriggerResponseArrayOutput)
}

type AzureBlobStorageApplicationLogsConfig struct {
	Level           *LogLevel `pulumi:"level"`
	RetentionInDays *int      `pulumi:"retentionInDays"`
	SasUrl          *string   `pulumi:"sasUrl"`
}





type AzureBlobStorageApplicationLogsConfigInput interface {
	pulumi.Input

	ToAzureBlobStorageApplicationLogsConfigOutput() AzureBlobStorageApplicationLogsConfigOutput
	ToAzureBlobStorageApplicationLogsConfigOutputWithContext(context.Context) AzureBlobStorageApplicationLogsConfigOutput
}

type AzureBlobStorageApplicationLogsConfigArgs struct {
	Level           LogLevelPtrInput      `pulumi:"level"`
	RetentionInDays pulumi.IntPtrInput    `pulumi:"retentionInDays"`
	SasUrl          pulumi.StringPtrInput `pulumi:"sasUrl"`
}

func (AzureBlobStorageApplicationLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageApplicationLogsConfig)(nil)).Elem()
}

func (i AzureBlobStorageApplicationLogsConfigArgs) ToAzureBlobStorageApplicationLogsConfigOutput() AzureBlobStorageApplicationLogsConfigOutput {
	return i.ToAzureBlobStorageApplicationLogsConfigOutputWithContext(context.Background())
}

func (i AzureBlobStorageApplicationLogsConfigArgs) ToAzureBlobStorageApplicationLogsConfigOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageApplicationLogsConfigOutput)
}

func (i AzureBlobStorageApplicationLogsConfigArgs) ToAzureBlobStorageApplicationLogsConfigPtrOutput() AzureBlobStorageApplicationLogsConfigPtrOutput {
	return i.ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i AzureBlobStorageApplicationLogsConfigArgs) ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageApplicationLogsConfigOutput).ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(ctx)
}









type AzureBlobStorageApplicationLogsConfigPtrInput interface {
	pulumi.Input

	ToAzureBlobStorageApplicationLogsConfigPtrOutput() AzureBlobStorageApplicationLogsConfigPtrOutput
	ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(context.Context) AzureBlobStorageApplicationLogsConfigPtrOutput
}

type azureBlobStorageApplicationLogsConfigPtrType AzureBlobStorageApplicationLogsConfigArgs

func AzureBlobStorageApplicationLogsConfigPtr(v *AzureBlobStorageApplicationLogsConfigArgs) AzureBlobStorageApplicationLogsConfigPtrInput {
	return (*azureBlobStorageApplicationLogsConfigPtrType)(v)
}

func (*azureBlobStorageApplicationLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageApplicationLogsConfig)(nil)).Elem()
}

func (i *azureBlobStorageApplicationLogsConfigPtrType) ToAzureBlobStorageApplicationLogsConfigPtrOutput() AzureBlobStorageApplicationLogsConfigPtrOutput {
	return i.ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i *azureBlobStorageApplicationLogsConfigPtrType) ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageApplicationLogsConfigPtrOutput)
}

type AzureBlobStorageApplicationLogsConfigOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageApplicationLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageApplicationLogsConfig)(nil)).Elem()
}

func (o AzureBlobStorageApplicationLogsConfigOutput) ToAzureBlobStorageApplicationLogsConfigOutput() AzureBlobStorageApplicationLogsConfigOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigOutput) ToAzureBlobStorageApplicationLogsConfigOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigOutput) ToAzureBlobStorageApplicationLogsConfigPtrOutput() AzureBlobStorageApplicationLogsConfigPtrOutput {
	return o.ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (o AzureBlobStorageApplicationLogsConfigOutput) ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureBlobStorageApplicationLogsConfig) *AzureBlobStorageApplicationLogsConfig {
		return &v
	}).(AzureBlobStorageApplicationLogsConfigPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigOutput) Level() LogLevelPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageApplicationLogsConfig) *LogLevel { return v.Level }).(LogLevelPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageApplicationLogsConfig) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageApplicationLogsConfig) *string { return v.SasUrl }).(pulumi.StringPtrOutput)
}

type AzureBlobStorageApplicationLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageApplicationLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageApplicationLogsConfig)(nil)).Elem()
}

func (o AzureBlobStorageApplicationLogsConfigPtrOutput) ToAzureBlobStorageApplicationLogsConfigPtrOutput() AzureBlobStorageApplicationLogsConfigPtrOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigPtrOutput) ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigPtrOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigPtrOutput) Elem() AzureBlobStorageApplicationLogsConfigOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfig) AzureBlobStorageApplicationLogsConfig {
		if v != nil {
			return *v
		}
		var ret AzureBlobStorageApplicationLogsConfig
		return ret
	}).(AzureBlobStorageApplicationLogsConfigOutput)
}

func (o AzureBlobStorageApplicationLogsConfigPtrOutput) Level() LogLevelPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfig) *LogLevel {
		if v == nil {
			return nil
		}
		return v.Level
	}).(LogLevelPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigPtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfig) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigPtrOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfig) *string {
		if v == nil {
			return nil
		}
		return v.SasUrl
	}).(pulumi.StringPtrOutput)
}

type AzureBlobStorageApplicationLogsConfigResponse struct {
	Level           *string `pulumi:"level"`
	RetentionInDays *int    `pulumi:"retentionInDays"`
	SasUrl          *string `pulumi:"sasUrl"`
}

type AzureBlobStorageApplicationLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageApplicationLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageApplicationLogsConfigResponse)(nil)).Elem()
}

func (o AzureBlobStorageApplicationLogsConfigResponseOutput) ToAzureBlobStorageApplicationLogsConfigResponseOutput() AzureBlobStorageApplicationLogsConfigResponseOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigResponseOutput) ToAzureBlobStorageApplicationLogsConfigResponseOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigResponseOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageApplicationLogsConfigResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigResponseOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageApplicationLogsConfigResponse) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigResponseOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageApplicationLogsConfigResponse) *string { return v.SasUrl }).(pulumi.StringPtrOutput)
}

type AzureBlobStorageApplicationLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageApplicationLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageApplicationLogsConfigResponse)(nil)).Elem()
}

func (o AzureBlobStorageApplicationLogsConfigResponsePtrOutput) ToAzureBlobStorageApplicationLogsConfigResponsePtrOutput() AzureBlobStorageApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigResponsePtrOutput) ToAzureBlobStorageApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigResponsePtrOutput) Elem() AzureBlobStorageApplicationLogsConfigResponseOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfigResponse) AzureBlobStorageApplicationLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret AzureBlobStorageApplicationLogsConfigResponse
		return ret
	}).(AzureBlobStorageApplicationLogsConfigResponseOutput)
}

func (o AzureBlobStorageApplicationLogsConfigResponsePtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Level
	}).(pulumi.StringPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigResponsePtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigResponsePtrOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.SasUrl
	}).(pulumi.StringPtrOutput)
}

type AzureBlobStorageHttpLogsConfig struct {
	Enabled         *bool   `pulumi:"enabled"`
	RetentionInDays *int    `pulumi:"retentionInDays"`
	SasUrl          *string `pulumi:"sasUrl"`
}





type AzureBlobStorageHttpLogsConfigInput interface {
	pulumi.Input

	ToAzureBlobStorageHttpLogsConfigOutput() AzureBlobStorageHttpLogsConfigOutput
	ToAzureBlobStorageHttpLogsConfigOutputWithContext(context.Context) AzureBlobStorageHttpLogsConfigOutput
}

type AzureBlobStorageHttpLogsConfigArgs struct {
	Enabled         pulumi.BoolPtrInput   `pulumi:"enabled"`
	RetentionInDays pulumi.IntPtrInput    `pulumi:"retentionInDays"`
	SasUrl          pulumi.StringPtrInput `pulumi:"sasUrl"`
}

func (AzureBlobStorageHttpLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageHttpLogsConfig)(nil)).Elem()
}

func (i AzureBlobStorageHttpLogsConfigArgs) ToAzureBlobStorageHttpLogsConfigOutput() AzureBlobStorageHttpLogsConfigOutput {
	return i.ToAzureBlobStorageHttpLogsConfigOutputWithContext(context.Background())
}

func (i AzureBlobStorageHttpLogsConfigArgs) ToAzureBlobStorageHttpLogsConfigOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageHttpLogsConfigOutput)
}

func (i AzureBlobStorageHttpLogsConfigArgs) ToAzureBlobStorageHttpLogsConfigPtrOutput() AzureBlobStorageHttpLogsConfigPtrOutput {
	return i.ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (i AzureBlobStorageHttpLogsConfigArgs) ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageHttpLogsConfigOutput).ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(ctx)
}









type AzureBlobStorageHttpLogsConfigPtrInput interface {
	pulumi.Input

	ToAzureBlobStorageHttpLogsConfigPtrOutput() AzureBlobStorageHttpLogsConfigPtrOutput
	ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(context.Context) AzureBlobStorageHttpLogsConfigPtrOutput
}

type azureBlobStorageHttpLogsConfigPtrType AzureBlobStorageHttpLogsConfigArgs

func AzureBlobStorageHttpLogsConfigPtr(v *AzureBlobStorageHttpLogsConfigArgs) AzureBlobStorageHttpLogsConfigPtrInput {
	return (*azureBlobStorageHttpLogsConfigPtrType)(v)
}

func (*azureBlobStorageHttpLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageHttpLogsConfig)(nil)).Elem()
}

func (i *azureBlobStorageHttpLogsConfigPtrType) ToAzureBlobStorageHttpLogsConfigPtrOutput() AzureBlobStorageHttpLogsConfigPtrOutput {
	return i.ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (i *azureBlobStorageHttpLogsConfigPtrType) ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageHttpLogsConfigPtrOutput)
}

type AzureBlobStorageHttpLogsConfigOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageHttpLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageHttpLogsConfig)(nil)).Elem()
}

func (o AzureBlobStorageHttpLogsConfigOutput) ToAzureBlobStorageHttpLogsConfigOutput() AzureBlobStorageHttpLogsConfigOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigOutput) ToAzureBlobStorageHttpLogsConfigOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigOutput) ToAzureBlobStorageHttpLogsConfigPtrOutput() AzureBlobStorageHttpLogsConfigPtrOutput {
	return o.ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (o AzureBlobStorageHttpLogsConfigOutput) ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureBlobStorageHttpLogsConfig) *AzureBlobStorageHttpLogsConfig {
		return &v
	}).(AzureBlobStorageHttpLogsConfigPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageHttpLogsConfig) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageHttpLogsConfig) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageHttpLogsConfig) *string { return v.SasUrl }).(pulumi.StringPtrOutput)
}

type AzureBlobStorageHttpLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageHttpLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageHttpLogsConfig)(nil)).Elem()
}

func (o AzureBlobStorageHttpLogsConfigPtrOutput) ToAzureBlobStorageHttpLogsConfigPtrOutput() AzureBlobStorageHttpLogsConfigPtrOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigPtrOutput) ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigPtrOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigPtrOutput) Elem() AzureBlobStorageHttpLogsConfigOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfig) AzureBlobStorageHttpLogsConfig {
		if v != nil {
			return *v
		}
		var ret AzureBlobStorageHttpLogsConfig
		return ret
	}).(AzureBlobStorageHttpLogsConfigOutput)
}

func (o AzureBlobStorageHttpLogsConfigPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfig) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigPtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfig) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigPtrOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfig) *string {
		if v == nil {
			return nil
		}
		return v.SasUrl
	}).(pulumi.StringPtrOutput)
}

type AzureBlobStorageHttpLogsConfigResponse struct {
	Enabled         *bool   `pulumi:"enabled"`
	RetentionInDays *int    `pulumi:"retentionInDays"`
	SasUrl          *string `pulumi:"sasUrl"`
}

type AzureBlobStorageHttpLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageHttpLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageHttpLogsConfigResponse)(nil)).Elem()
}

func (o AzureBlobStorageHttpLogsConfigResponseOutput) ToAzureBlobStorageHttpLogsConfigResponseOutput() AzureBlobStorageHttpLogsConfigResponseOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigResponseOutput) ToAzureBlobStorageHttpLogsConfigResponseOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigResponseOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageHttpLogsConfigResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigResponseOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageHttpLogsConfigResponse) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigResponseOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageHttpLogsConfigResponse) *string { return v.SasUrl }).(pulumi.StringPtrOutput)
}

type AzureBlobStorageHttpLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageHttpLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageHttpLogsConfigResponse)(nil)).Elem()
}

func (o AzureBlobStorageHttpLogsConfigResponsePtrOutput) ToAzureBlobStorageHttpLogsConfigResponsePtrOutput() AzureBlobStorageHttpLogsConfigResponsePtrOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigResponsePtrOutput) ToAzureBlobStorageHttpLogsConfigResponsePtrOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigResponsePtrOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigResponsePtrOutput) Elem() AzureBlobStorageHttpLogsConfigResponseOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfigResponse) AzureBlobStorageHttpLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret AzureBlobStorageHttpLogsConfigResponse
		return ret
	}).(AzureBlobStorageHttpLogsConfigResponseOutput)
}

func (o AzureBlobStorageHttpLogsConfigResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigResponsePtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigResponsePtrOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.SasUrl
	}).(pulumi.StringPtrOutput)
}

type AzureTableStorageApplicationLogsConfig struct {
	Level  *LogLevel `pulumi:"level"`
	SasUrl *string   `pulumi:"sasUrl"`
}





type AzureTableStorageApplicationLogsConfigInput interface {
	pulumi.Input

	ToAzureTableStorageApplicationLogsConfigOutput() AzureTableStorageApplicationLogsConfigOutput
	ToAzureTableStorageApplicationLogsConfigOutputWithContext(context.Context) AzureTableStorageApplicationLogsConfigOutput
}

type AzureTableStorageApplicationLogsConfigArgs struct {
	Level  LogLevelPtrInput      `pulumi:"level"`
	SasUrl pulumi.StringPtrInput `pulumi:"sasUrl"`
}

func (AzureTableStorageApplicationLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureTableStorageApplicationLogsConfig)(nil)).Elem()
}

func (i AzureTableStorageApplicationLogsConfigArgs) ToAzureTableStorageApplicationLogsConfigOutput() AzureTableStorageApplicationLogsConfigOutput {
	return i.ToAzureTableStorageApplicationLogsConfigOutputWithContext(context.Background())
}

func (i AzureTableStorageApplicationLogsConfigArgs) ToAzureTableStorageApplicationLogsConfigOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureTableStorageApplicationLogsConfigOutput)
}

func (i AzureTableStorageApplicationLogsConfigArgs) ToAzureTableStorageApplicationLogsConfigPtrOutput() AzureTableStorageApplicationLogsConfigPtrOutput {
	return i.ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i AzureTableStorageApplicationLogsConfigArgs) ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureTableStorageApplicationLogsConfigOutput).ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(ctx)
}









type AzureTableStorageApplicationLogsConfigPtrInput interface {
	pulumi.Input

	ToAzureTableStorageApplicationLogsConfigPtrOutput() AzureTableStorageApplicationLogsConfigPtrOutput
	ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(context.Context) AzureTableStorageApplicationLogsConfigPtrOutput
}

type azureTableStorageApplicationLogsConfigPtrType AzureTableStorageApplicationLogsConfigArgs

func AzureTableStorageApplicationLogsConfigPtr(v *AzureTableStorageApplicationLogsConfigArgs) AzureTableStorageApplicationLogsConfigPtrInput {
	return (*azureTableStorageApplicationLogsConfigPtrType)(v)
}

func (*azureTableStorageApplicationLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureTableStorageApplicationLogsConfig)(nil)).Elem()
}

func (i *azureTableStorageApplicationLogsConfigPtrType) ToAzureTableStorageApplicationLogsConfigPtrOutput() AzureTableStorageApplicationLogsConfigPtrOutput {
	return i.ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i *azureTableStorageApplicationLogsConfigPtrType) ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureTableStorageApplicationLogsConfigPtrOutput)
}

type AzureTableStorageApplicationLogsConfigOutput struct{ *pulumi.OutputState }

func (AzureTableStorageApplicationLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureTableStorageApplicationLogsConfig)(nil)).Elem()
}

func (o AzureTableStorageApplicationLogsConfigOutput) ToAzureTableStorageApplicationLogsConfigOutput() AzureTableStorageApplicationLogsConfigOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigOutput) ToAzureTableStorageApplicationLogsConfigOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigOutput) ToAzureTableStorageApplicationLogsConfigPtrOutput() AzureTableStorageApplicationLogsConfigPtrOutput {
	return o.ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (o AzureTableStorageApplicationLogsConfigOutput) ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureTableStorageApplicationLogsConfig) *AzureTableStorageApplicationLogsConfig {
		return &v
	}).(AzureTableStorageApplicationLogsConfigPtrOutput)
}

func (o AzureTableStorageApplicationLogsConfigOutput) Level() LogLevelPtrOutput {
	return o.ApplyT(func(v AzureTableStorageApplicationLogsConfig) *LogLevel { return v.Level }).(LogLevelPtrOutput)
}

func (o AzureTableStorageApplicationLogsConfigOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureTableStorageApplicationLogsConfig) *string { return v.SasUrl }).(pulumi.StringPtrOutput)
}

type AzureTableStorageApplicationLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (AzureTableStorageApplicationLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureTableStorageApplicationLogsConfig)(nil)).Elem()
}

func (o AzureTableStorageApplicationLogsConfigPtrOutput) ToAzureTableStorageApplicationLogsConfigPtrOutput() AzureTableStorageApplicationLogsConfigPtrOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigPtrOutput) ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigPtrOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigPtrOutput) Elem() AzureTableStorageApplicationLogsConfigOutput {
	return o.ApplyT(func(v *AzureTableStorageApplicationLogsConfig) AzureTableStorageApplicationLogsConfig {
		if v != nil {
			return *v
		}
		var ret AzureTableStorageApplicationLogsConfig
		return ret
	}).(AzureTableStorageApplicationLogsConfigOutput)
}

func (o AzureTableStorageApplicationLogsConfigPtrOutput) Level() LogLevelPtrOutput {
	return o.ApplyT(func(v *AzureTableStorageApplicationLogsConfig) *LogLevel {
		if v == nil {
			return nil
		}
		return v.Level
	}).(LogLevelPtrOutput)
}

func (o AzureTableStorageApplicationLogsConfigPtrOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureTableStorageApplicationLogsConfig) *string {
		if v == nil {
			return nil
		}
		return v.SasUrl
	}).(pulumi.StringPtrOutput)
}

type AzureTableStorageApplicationLogsConfigResponse struct {
	Level  *string `pulumi:"level"`
	SasUrl *string `pulumi:"sasUrl"`
}

type AzureTableStorageApplicationLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (AzureTableStorageApplicationLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureTableStorageApplicationLogsConfigResponse)(nil)).Elem()
}

func (o AzureTableStorageApplicationLogsConfigResponseOutput) ToAzureTableStorageApplicationLogsConfigResponseOutput() AzureTableStorageApplicationLogsConfigResponseOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigResponseOutput) ToAzureTableStorageApplicationLogsConfigResponseOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigResponseOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureTableStorageApplicationLogsConfigResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o AzureTableStorageApplicationLogsConfigResponseOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureTableStorageApplicationLogsConfigResponse) *string { return v.SasUrl }).(pulumi.StringPtrOutput)
}

type AzureTableStorageApplicationLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureTableStorageApplicationLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureTableStorageApplicationLogsConfigResponse)(nil)).Elem()
}

func (o AzureTableStorageApplicationLogsConfigResponsePtrOutput) ToAzureTableStorageApplicationLogsConfigResponsePtrOutput() AzureTableStorageApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigResponsePtrOutput) ToAzureTableStorageApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigResponsePtrOutput) Elem() AzureTableStorageApplicationLogsConfigResponseOutput {
	return o.ApplyT(func(v *AzureTableStorageApplicationLogsConfigResponse) AzureTableStorageApplicationLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret AzureTableStorageApplicationLogsConfigResponse
		return ret
	}).(AzureTableStorageApplicationLogsConfigResponseOutput)
}

func (o AzureTableStorageApplicationLogsConfigResponsePtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureTableStorageApplicationLogsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Level
	}).(pulumi.StringPtrOutput)
}

func (o AzureTableStorageApplicationLogsConfigResponsePtrOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureTableStorageApplicationLogsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.SasUrl
	}).(pulumi.StringPtrOutput)
}

type BackupSchedule struct {
	FrequencyInterval     *int          `pulumi:"frequencyInterval"`
	FrequencyUnit         FrequencyUnit `pulumi:"frequencyUnit"`
	KeepAtLeastOneBackup  *bool         `pulumi:"keepAtLeastOneBackup"`
	LastExecutionTime     *string       `pulumi:"lastExecutionTime"`
	RetentionPeriodInDays *int          `pulumi:"retentionPeriodInDays"`
	StartTime             *string       `pulumi:"startTime"`
}





type BackupScheduleInput interface {
	pulumi.Input

	ToBackupScheduleOutput() BackupScheduleOutput
	ToBackupScheduleOutputWithContext(context.Context) BackupScheduleOutput
}

type BackupScheduleArgs struct {
	FrequencyInterval     pulumi.IntPtrInput    `pulumi:"frequencyInterval"`
	FrequencyUnit         FrequencyUnitInput    `pulumi:"frequencyUnit"`
	KeepAtLeastOneBackup  pulumi.BoolPtrInput   `pulumi:"keepAtLeastOneBackup"`
	LastExecutionTime     pulumi.StringPtrInput `pulumi:"lastExecutionTime"`
	RetentionPeriodInDays pulumi.IntPtrInput    `pulumi:"retentionPeriodInDays"`
	StartTime             pulumi.StringPtrInput `pulumi:"startTime"`
}

func (BackupScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupSchedule)(nil)).Elem()
}

func (i BackupScheduleArgs) ToBackupScheduleOutput() BackupScheduleOutput {
	return i.ToBackupScheduleOutputWithContext(context.Background())
}

func (i BackupScheduleArgs) ToBackupScheduleOutputWithContext(ctx context.Context) BackupScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupScheduleOutput)
}

func (i BackupScheduleArgs) ToBackupSchedulePtrOutput() BackupSchedulePtrOutput {
	return i.ToBackupSchedulePtrOutputWithContext(context.Background())
}

func (i BackupScheduleArgs) ToBackupSchedulePtrOutputWithContext(ctx context.Context) BackupSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupScheduleOutput).ToBackupSchedulePtrOutputWithContext(ctx)
}









type BackupSchedulePtrInput interface {
	pulumi.Input

	ToBackupSchedulePtrOutput() BackupSchedulePtrOutput
	ToBackupSchedulePtrOutputWithContext(context.Context) BackupSchedulePtrOutput
}

type backupSchedulePtrType BackupScheduleArgs

func BackupSchedulePtr(v *BackupScheduleArgs) BackupSchedulePtrInput {
	return (*backupSchedulePtrType)(v)
}

func (*backupSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupSchedule)(nil)).Elem()
}

func (i *backupSchedulePtrType) ToBackupSchedulePtrOutput() BackupSchedulePtrOutput {
	return i.ToBackupSchedulePtrOutputWithContext(context.Background())
}

func (i *backupSchedulePtrType) ToBackupSchedulePtrOutputWithContext(ctx context.Context) BackupSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupSchedulePtrOutput)
}

type BackupScheduleOutput struct{ *pulumi.OutputState }

func (BackupScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupSchedule)(nil)).Elem()
}

func (o BackupScheduleOutput) ToBackupScheduleOutput() BackupScheduleOutput {
	return o
}

func (o BackupScheduleOutput) ToBackupScheduleOutputWithContext(ctx context.Context) BackupScheduleOutput {
	return o
}

func (o BackupScheduleOutput) ToBackupSchedulePtrOutput() BackupSchedulePtrOutput {
	return o.ToBackupSchedulePtrOutputWithContext(context.Background())
}

func (o BackupScheduleOutput) ToBackupSchedulePtrOutputWithContext(ctx context.Context) BackupSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupSchedule) *BackupSchedule {
		return &v
	}).(BackupSchedulePtrOutput)
}

func (o BackupScheduleOutput) FrequencyInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackupSchedule) *int { return v.FrequencyInterval }).(pulumi.IntPtrOutput)
}

func (o BackupScheduleOutput) FrequencyUnit() FrequencyUnitOutput {
	return o.ApplyT(func(v BackupSchedule) FrequencyUnit { return v.FrequencyUnit }).(FrequencyUnitOutput)
}

func (o BackupScheduleOutput) KeepAtLeastOneBackup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BackupSchedule) *bool { return v.KeepAtLeastOneBackup }).(pulumi.BoolPtrOutput)
}

func (o BackupScheduleOutput) LastExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupSchedule) *string { return v.LastExecutionTime }).(pulumi.StringPtrOutput)
}

func (o BackupScheduleOutput) RetentionPeriodInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackupSchedule) *int { return v.RetentionPeriodInDays }).(pulumi.IntPtrOutput)
}

func (o BackupScheduleOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupSchedule) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type BackupSchedulePtrOutput struct{ *pulumi.OutputState }

func (BackupSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupSchedule)(nil)).Elem()
}

func (o BackupSchedulePtrOutput) ToBackupSchedulePtrOutput() BackupSchedulePtrOutput {
	return o
}

func (o BackupSchedulePtrOutput) ToBackupSchedulePtrOutputWithContext(ctx context.Context) BackupSchedulePtrOutput {
	return o
}

func (o BackupSchedulePtrOutput) Elem() BackupScheduleOutput {
	return o.ApplyT(func(v *BackupSchedule) BackupSchedule {
		if v != nil {
			return *v
		}
		var ret BackupSchedule
		return ret
	}).(BackupScheduleOutput)
}

func (o BackupSchedulePtrOutput) FrequencyInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupSchedule) *int {
		if v == nil {
			return nil
		}
		return v.FrequencyInterval
	}).(pulumi.IntPtrOutput)
}

func (o BackupSchedulePtrOutput) FrequencyUnit() FrequencyUnitPtrOutput {
	return o.ApplyT(func(v *BackupSchedule) *FrequencyUnit {
		if v == nil {
			return nil
		}
		return &v.FrequencyUnit
	}).(FrequencyUnitPtrOutput)
}

func (o BackupSchedulePtrOutput) KeepAtLeastOneBackup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackupSchedule) *bool {
		if v == nil {
			return nil
		}
		return v.KeepAtLeastOneBackup
	}).(pulumi.BoolPtrOutput)
}

func (o BackupSchedulePtrOutput) LastExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupSchedule) *string {
		if v == nil {
			return nil
		}
		return v.LastExecutionTime
	}).(pulumi.StringPtrOutput)
}

func (o BackupSchedulePtrOutput) RetentionPeriodInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupSchedule) *int {
		if v == nil {
			return nil
		}
		return v.RetentionPeriodInDays
	}).(pulumi.IntPtrOutput)
}

func (o BackupSchedulePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupSchedule) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

type BackupScheduleResponse struct {
	FrequencyInterval     *int    `pulumi:"frequencyInterval"`
	FrequencyUnit         string  `pulumi:"frequencyUnit"`
	KeepAtLeastOneBackup  *bool   `pulumi:"keepAtLeastOneBackup"`
	LastExecutionTime     *string `pulumi:"lastExecutionTime"`
	RetentionPeriodInDays *int    `pulumi:"retentionPeriodInDays"`
	StartTime             *string `pulumi:"startTime"`
}

type BackupScheduleResponseOutput struct{ *pulumi.OutputState }

func (BackupScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupScheduleResponse)(nil)).Elem()
}

func (o BackupScheduleResponseOutput) ToBackupScheduleResponseOutput() BackupScheduleResponseOutput {
	return o
}

func (o BackupScheduleResponseOutput) ToBackupScheduleResponseOutputWithContext(ctx context.Context) BackupScheduleResponseOutput {
	return o
}

func (o BackupScheduleResponseOutput) FrequencyInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackupScheduleResponse) *int { return v.FrequencyInterval }).(pulumi.IntPtrOutput)
}

func (o BackupScheduleResponseOutput) FrequencyUnit() pulumi.StringOutput {
	return o.ApplyT(func(v BackupScheduleResponse) string { return v.FrequencyUnit }).(pulumi.StringOutput)
}

func (o BackupScheduleResponseOutput) KeepAtLeastOneBackup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BackupScheduleResponse) *bool { return v.KeepAtLeastOneBackup }).(pulumi.BoolPtrOutput)
}

func (o BackupScheduleResponseOutput) LastExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupScheduleResponse) *string { return v.LastExecutionTime }).(pulumi.StringPtrOutput)
}

func (o BackupScheduleResponseOutput) RetentionPeriodInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BackupScheduleResponse) *int { return v.RetentionPeriodInDays }).(pulumi.IntPtrOutput)
}

func (o BackupScheduleResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupScheduleResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type BackupScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (BackupScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupScheduleResponse)(nil)).Elem()
}

func (o BackupScheduleResponsePtrOutput) ToBackupScheduleResponsePtrOutput() BackupScheduleResponsePtrOutput {
	return o
}

func (o BackupScheduleResponsePtrOutput) ToBackupScheduleResponsePtrOutputWithContext(ctx context.Context) BackupScheduleResponsePtrOutput {
	return o
}

func (o BackupScheduleResponsePtrOutput) Elem() BackupScheduleResponseOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) BackupScheduleResponse {
		if v != nil {
			return *v
		}
		var ret BackupScheduleResponse
		return ret
	}).(BackupScheduleResponseOutput)
}

func (o BackupScheduleResponsePtrOutput) FrequencyInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return v.FrequencyInterval
	}).(pulumi.IntPtrOutput)
}

func (o BackupScheduleResponsePtrOutput) FrequencyUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FrequencyUnit
	}).(pulumi.StringPtrOutput)
}

func (o BackupScheduleResponsePtrOutput) KeepAtLeastOneBackup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) *bool {
		if v == nil {
			return nil
		}
		return v.KeepAtLeastOneBackup
	}).(pulumi.BoolPtrOutput)
}

func (o BackupScheduleResponsePtrOutput) LastExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastExecutionTime
	}).(pulumi.StringPtrOutput)
}

func (o BackupScheduleResponsePtrOutput) RetentionPeriodInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetentionPeriodInDays
	}).(pulumi.IntPtrOutput)
}

func (o BackupScheduleResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

type CloningInfo struct {
	AppSettingsOverrides      map[string]string `pulumi:"appSettingsOverrides"`
	CloneCustomHostNames      *bool             `pulumi:"cloneCustomHostNames"`
	CloneSourceControl        *bool             `pulumi:"cloneSourceControl"`
	ConfigureLoadBalancing    *bool             `pulumi:"configureLoadBalancing"`
	CorrelationId             *string           `pulumi:"correlationId"`
	HostingEnvironment        *string           `pulumi:"hostingEnvironment"`
	Overwrite                 *bool             `pulumi:"overwrite"`
	SourceWebAppId            *string           `pulumi:"sourceWebAppId"`
	TrafficManagerProfileId   *string           `pulumi:"trafficManagerProfileId"`
	TrafficManagerProfileName *string           `pulumi:"trafficManagerProfileName"`
}





type CloningInfoInput interface {
	pulumi.Input

	ToCloningInfoOutput() CloningInfoOutput
	ToCloningInfoOutputWithContext(context.Context) CloningInfoOutput
}

type CloningInfoArgs struct {
	AppSettingsOverrides      pulumi.StringMapInput `pulumi:"appSettingsOverrides"`
	CloneCustomHostNames      pulumi.BoolPtrInput   `pulumi:"cloneCustomHostNames"`
	CloneSourceControl        pulumi.BoolPtrInput   `pulumi:"cloneSourceControl"`
	ConfigureLoadBalancing    pulumi.BoolPtrInput   `pulumi:"configureLoadBalancing"`
	CorrelationId             pulumi.StringPtrInput `pulumi:"correlationId"`
	HostingEnvironment        pulumi.StringPtrInput `pulumi:"hostingEnvironment"`
	Overwrite                 pulumi.BoolPtrInput   `pulumi:"overwrite"`
	SourceWebAppId            pulumi.StringPtrInput `pulumi:"sourceWebAppId"`
	TrafficManagerProfileId   pulumi.StringPtrInput `pulumi:"trafficManagerProfileId"`
	TrafficManagerProfileName pulumi.StringPtrInput `pulumi:"trafficManagerProfileName"`
}

func (CloningInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloningInfo)(nil)).Elem()
}

func (i CloningInfoArgs) ToCloningInfoOutput() CloningInfoOutput {
	return i.ToCloningInfoOutputWithContext(context.Background())
}

func (i CloningInfoArgs) ToCloningInfoOutputWithContext(ctx context.Context) CloningInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloningInfoOutput)
}

func (i CloningInfoArgs) ToCloningInfoPtrOutput() CloningInfoPtrOutput {
	return i.ToCloningInfoPtrOutputWithContext(context.Background())
}

func (i CloningInfoArgs) ToCloningInfoPtrOutputWithContext(ctx context.Context) CloningInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloningInfoOutput).ToCloningInfoPtrOutputWithContext(ctx)
}









type CloningInfoPtrInput interface {
	pulumi.Input

	ToCloningInfoPtrOutput() CloningInfoPtrOutput
	ToCloningInfoPtrOutputWithContext(context.Context) CloningInfoPtrOutput
}

type cloningInfoPtrType CloningInfoArgs

func CloningInfoPtr(v *CloningInfoArgs) CloningInfoPtrInput {
	return (*cloningInfoPtrType)(v)
}

func (*cloningInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloningInfo)(nil)).Elem()
}

func (i *cloningInfoPtrType) ToCloningInfoPtrOutput() CloningInfoPtrOutput {
	return i.ToCloningInfoPtrOutputWithContext(context.Background())
}

func (i *cloningInfoPtrType) ToCloningInfoPtrOutputWithContext(ctx context.Context) CloningInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloningInfoPtrOutput)
}

type CloningInfoOutput struct{ *pulumi.OutputState }

func (CloningInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloningInfo)(nil)).Elem()
}

func (o CloningInfoOutput) ToCloningInfoOutput() CloningInfoOutput {
	return o
}

func (o CloningInfoOutput) ToCloningInfoOutputWithContext(ctx context.Context) CloningInfoOutput {
	return o
}

func (o CloningInfoOutput) ToCloningInfoPtrOutput() CloningInfoPtrOutput {
	return o.ToCloningInfoPtrOutputWithContext(context.Background())
}

func (o CloningInfoOutput) ToCloningInfoPtrOutputWithContext(ctx context.Context) CloningInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloningInfo) *CloningInfo {
		return &v
	}).(CloningInfoPtrOutput)
}

func (o CloningInfoOutput) AppSettingsOverrides() pulumi.StringMapOutput {
	return o.ApplyT(func(v CloningInfo) map[string]string { return v.AppSettingsOverrides }).(pulumi.StringMapOutput)
}

func (o CloningInfoOutput) CloneCustomHostNames() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfo) *bool { return v.CloneCustomHostNames }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoOutput) CloneSourceControl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfo) *bool { return v.CloneSourceControl }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoOutput) ConfigureLoadBalancing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfo) *bool { return v.ConfigureLoadBalancing }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfo) *string { return v.CorrelationId }).(pulumi.StringPtrOutput)
}

func (o CloningInfoOutput) HostingEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfo) *string { return v.HostingEnvironment }).(pulumi.StringPtrOutput)
}

func (o CloningInfoOutput) Overwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfo) *bool { return v.Overwrite }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoOutput) SourceWebAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfo) *string { return v.SourceWebAppId }).(pulumi.StringPtrOutput)
}

func (o CloningInfoOutput) TrafficManagerProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfo) *string { return v.TrafficManagerProfileId }).(pulumi.StringPtrOutput)
}

func (o CloningInfoOutput) TrafficManagerProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfo) *string { return v.TrafficManagerProfileName }).(pulumi.StringPtrOutput)
}

type CloningInfoPtrOutput struct{ *pulumi.OutputState }

func (CloningInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloningInfo)(nil)).Elem()
}

func (o CloningInfoPtrOutput) ToCloningInfoPtrOutput() CloningInfoPtrOutput {
	return o
}

func (o CloningInfoPtrOutput) ToCloningInfoPtrOutputWithContext(ctx context.Context) CloningInfoPtrOutput {
	return o
}

func (o CloningInfoPtrOutput) Elem() CloningInfoOutput {
	return o.ApplyT(func(v *CloningInfo) CloningInfo {
		if v != nil {
			return *v
		}
		var ret CloningInfo
		return ret
	}).(CloningInfoOutput)
}

func (o CloningInfoPtrOutput) AppSettingsOverrides() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CloningInfo) map[string]string {
		if v == nil {
			return nil
		}
		return v.AppSettingsOverrides
	}).(pulumi.StringMapOutput)
}

func (o CloningInfoPtrOutput) CloneCustomHostNames() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *bool {
		if v == nil {
			return nil
		}
		return v.CloneCustomHostNames
	}).(pulumi.BoolPtrOutput)
}

func (o CloningInfoPtrOutput) CloneSourceControl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *bool {
		if v == nil {
			return nil
		}
		return v.CloneSourceControl
	}).(pulumi.BoolPtrOutput)
}

func (o CloningInfoPtrOutput) ConfigureLoadBalancing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *bool {
		if v == nil {
			return nil
		}
		return v.ConfigureLoadBalancing
	}).(pulumi.BoolPtrOutput)
}

func (o CloningInfoPtrOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *string {
		if v == nil {
			return nil
		}
		return v.CorrelationId
	}).(pulumi.StringPtrOutput)
}

func (o CloningInfoPtrOutput) HostingEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *string {
		if v == nil {
			return nil
		}
		return v.HostingEnvironment
	}).(pulumi.StringPtrOutput)
}

func (o CloningInfoPtrOutput) Overwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *bool {
		if v == nil {
			return nil
		}
		return v.Overwrite
	}).(pulumi.BoolPtrOutput)
}

func (o CloningInfoPtrOutput) SourceWebAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *string {
		if v == nil {
			return nil
		}
		return v.SourceWebAppId
	}).(pulumi.StringPtrOutput)
}

func (o CloningInfoPtrOutput) TrafficManagerProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *string {
		if v == nil {
			return nil
		}
		return v.TrafficManagerProfileId
	}).(pulumi.StringPtrOutput)
}

func (o CloningInfoPtrOutput) TrafficManagerProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *string {
		if v == nil {
			return nil
		}
		return v.TrafficManagerProfileName
	}).(pulumi.StringPtrOutput)
}

type CloningInfoResponse struct {
	AppSettingsOverrides      map[string]string `pulumi:"appSettingsOverrides"`
	CloneCustomHostNames      *bool             `pulumi:"cloneCustomHostNames"`
	CloneSourceControl        *bool             `pulumi:"cloneSourceControl"`
	ConfigureLoadBalancing    *bool             `pulumi:"configureLoadBalancing"`
	CorrelationId             *string           `pulumi:"correlationId"`
	HostingEnvironment        *string           `pulumi:"hostingEnvironment"`
	Overwrite                 *bool             `pulumi:"overwrite"`
	SourceWebAppId            *string           `pulumi:"sourceWebAppId"`
	TrafficManagerProfileId   *string           `pulumi:"trafficManagerProfileId"`
	TrafficManagerProfileName *string           `pulumi:"trafficManagerProfileName"`
}

type CloningInfoResponseOutput struct{ *pulumi.OutputState }

func (CloningInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloningInfoResponse)(nil)).Elem()
}

func (o CloningInfoResponseOutput) ToCloningInfoResponseOutput() CloningInfoResponseOutput {
	return o
}

func (o CloningInfoResponseOutput) ToCloningInfoResponseOutputWithContext(ctx context.Context) CloningInfoResponseOutput {
	return o
}

func (o CloningInfoResponseOutput) AppSettingsOverrides() pulumi.StringMapOutput {
	return o.ApplyT(func(v CloningInfoResponse) map[string]string { return v.AppSettingsOverrides }).(pulumi.StringMapOutput)
}

func (o CloningInfoResponseOutput) CloneCustomHostNames() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfoResponse) *bool { return v.CloneCustomHostNames }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoResponseOutput) CloneSourceControl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfoResponse) *bool { return v.CloneSourceControl }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoResponseOutput) ConfigureLoadBalancing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfoResponse) *bool { return v.ConfigureLoadBalancing }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoResponseOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfoResponse) *string { return v.CorrelationId }).(pulumi.StringPtrOutput)
}

func (o CloningInfoResponseOutput) HostingEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfoResponse) *string { return v.HostingEnvironment }).(pulumi.StringPtrOutput)
}

func (o CloningInfoResponseOutput) Overwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfoResponse) *bool { return v.Overwrite }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoResponseOutput) SourceWebAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfoResponse) *string { return v.SourceWebAppId }).(pulumi.StringPtrOutput)
}

func (o CloningInfoResponseOutput) TrafficManagerProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfoResponse) *string { return v.TrafficManagerProfileId }).(pulumi.StringPtrOutput)
}

func (o CloningInfoResponseOutput) TrafficManagerProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfoResponse) *string { return v.TrafficManagerProfileName }).(pulumi.StringPtrOutput)
}

type CloningInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (CloningInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloningInfoResponse)(nil)).Elem()
}

func (o CloningInfoResponsePtrOutput) ToCloningInfoResponsePtrOutput() CloningInfoResponsePtrOutput {
	return o
}

func (o CloningInfoResponsePtrOutput) ToCloningInfoResponsePtrOutputWithContext(ctx context.Context) CloningInfoResponsePtrOutput {
	return o
}

func (o CloningInfoResponsePtrOutput) Elem() CloningInfoResponseOutput {
	return o.ApplyT(func(v *CloningInfoResponse) CloningInfoResponse {
		if v != nil {
			return *v
		}
		var ret CloningInfoResponse
		return ret
	}).(CloningInfoResponseOutput)
}

func (o CloningInfoResponsePtrOutput) AppSettingsOverrides() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CloningInfoResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.AppSettingsOverrides
	}).(pulumi.StringMapOutput)
}

func (o CloningInfoResponsePtrOutput) CloneCustomHostNames() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloningInfoResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CloneCustomHostNames
	}).(pulumi.BoolPtrOutput)
}

func (o CloningInfoResponsePtrOutput) CloneSourceControl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloningInfoResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CloneSourceControl
	}).(pulumi.BoolPtrOutput)
}

func (o CloningInfoResponsePtrOutput) ConfigureLoadBalancing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloningInfoResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ConfigureLoadBalancing
	}).(pulumi.BoolPtrOutput)
}

func (o CloningInfoResponsePtrOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.CorrelationId
	}).(pulumi.StringPtrOutput)
}

func (o CloningInfoResponsePtrOutput) HostingEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.HostingEnvironment
	}).(pulumi.StringPtrOutput)
}

func (o CloningInfoResponsePtrOutput) Overwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloningInfoResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Overwrite
	}).(pulumi.BoolPtrOutput)
}

func (o CloningInfoResponsePtrOutput) SourceWebAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.SourceWebAppId
	}).(pulumi.StringPtrOutput)
}

func (o CloningInfoResponsePtrOutput) TrafficManagerProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.TrafficManagerProfileId
	}).(pulumi.StringPtrOutput)
}

func (o CloningInfoResponsePtrOutput) TrafficManagerProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.TrafficManagerProfileName
	}).(pulumi.StringPtrOutput)
}

type ConnStringInfo struct {
	ConnectionString *string            `pulumi:"connectionString"`
	Name             *string            `pulumi:"name"`
	Type             DatabaseServerType `pulumi:"type"`
}





type ConnStringInfoInput interface {
	pulumi.Input

	ToConnStringInfoOutput() ConnStringInfoOutput
	ToConnStringInfoOutputWithContext(context.Context) ConnStringInfoOutput
}

type ConnStringInfoArgs struct {
	ConnectionString pulumi.StringPtrInput   `pulumi:"connectionString"`
	Name             pulumi.StringPtrInput   `pulumi:"name"`
	Type             DatabaseServerTypeInput `pulumi:"type"`
}

func (ConnStringInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringInfo)(nil)).Elem()
}

func (i ConnStringInfoArgs) ToConnStringInfoOutput() ConnStringInfoOutput {
	return i.ToConnStringInfoOutputWithContext(context.Background())
}

func (i ConnStringInfoArgs) ToConnStringInfoOutputWithContext(ctx context.Context) ConnStringInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnStringInfoOutput)
}





type ConnStringInfoArrayInput interface {
	pulumi.Input

	ToConnStringInfoArrayOutput() ConnStringInfoArrayOutput
	ToConnStringInfoArrayOutputWithContext(context.Context) ConnStringInfoArrayOutput
}

type ConnStringInfoArray []ConnStringInfoInput

func (ConnStringInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnStringInfo)(nil)).Elem()
}

func (i ConnStringInfoArray) ToConnStringInfoArrayOutput() ConnStringInfoArrayOutput {
	return i.ToConnStringInfoArrayOutputWithContext(context.Background())
}

func (i ConnStringInfoArray) ToConnStringInfoArrayOutputWithContext(ctx context.Context) ConnStringInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnStringInfoArrayOutput)
}

type ConnStringInfoOutput struct{ *pulumi.OutputState }

func (ConnStringInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringInfo)(nil)).Elem()
}

func (o ConnStringInfoOutput) ToConnStringInfoOutput() ConnStringInfoOutput {
	return o
}

func (o ConnStringInfoOutput) ToConnStringInfoOutputWithContext(ctx context.Context) ConnStringInfoOutput {
	return o
}

func (o ConnStringInfoOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnStringInfo) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o ConnStringInfoOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnStringInfo) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConnStringInfoOutput) Type() DatabaseServerTypeOutput {
	return o.ApplyT(func(v ConnStringInfo) DatabaseServerType { return v.Type }).(DatabaseServerTypeOutput)
}

type ConnStringInfoArrayOutput struct{ *pulumi.OutputState }

func (ConnStringInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnStringInfo)(nil)).Elem()
}

func (o ConnStringInfoArrayOutput) ToConnStringInfoArrayOutput() ConnStringInfoArrayOutput {
	return o
}

func (o ConnStringInfoArrayOutput) ToConnStringInfoArrayOutputWithContext(ctx context.Context) ConnStringInfoArrayOutput {
	return o
}

func (o ConnStringInfoArrayOutput) Index(i pulumi.IntInput) ConnStringInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnStringInfo {
		return vs[0].([]ConnStringInfo)[vs[1].(int)]
	}).(ConnStringInfoOutput)
}

type ConnStringInfoResponse struct {
	ConnectionString *string `pulumi:"connectionString"`
	Name             *string `pulumi:"name"`
	Type             string  `pulumi:"type"`
}

type ConnStringInfoResponseOutput struct{ *pulumi.OutputState }

func (ConnStringInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringInfoResponse)(nil)).Elem()
}

func (o ConnStringInfoResponseOutput) ToConnStringInfoResponseOutput() ConnStringInfoResponseOutput {
	return o
}

func (o ConnStringInfoResponseOutput) ToConnStringInfoResponseOutputWithContext(ctx context.Context) ConnStringInfoResponseOutput {
	return o
}

func (o ConnStringInfoResponseOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnStringInfoResponse) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o ConnStringInfoResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnStringInfoResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConnStringInfoResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ConnStringInfoResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ConnStringInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (ConnStringInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnStringInfoResponse)(nil)).Elem()
}

func (o ConnStringInfoResponseArrayOutput) ToConnStringInfoResponseArrayOutput() ConnStringInfoResponseArrayOutput {
	return o
}

func (o ConnStringInfoResponseArrayOutput) ToConnStringInfoResponseArrayOutputWithContext(ctx context.Context) ConnStringInfoResponseArrayOutput {
	return o
}

func (o ConnStringInfoResponseArrayOutput) Index(i pulumi.IntInput) ConnStringInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnStringInfoResponse {
		return vs[0].([]ConnStringInfoResponse)[vs[1].(int)]
	}).(ConnStringInfoResponseOutput)
}

type ConnStringValueTypePair struct {
	Type  DatabaseServerType `pulumi:"type"`
	Value *string            `pulumi:"value"`
}





type ConnStringValueTypePairInput interface {
	pulumi.Input

	ToConnStringValueTypePairOutput() ConnStringValueTypePairOutput
	ToConnStringValueTypePairOutputWithContext(context.Context) ConnStringValueTypePairOutput
}

type ConnStringValueTypePairArgs struct {
	Type  DatabaseServerTypeInput `pulumi:"type"`
	Value pulumi.StringPtrInput   `pulumi:"value"`
}

func (ConnStringValueTypePairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringValueTypePair)(nil)).Elem()
}

func (i ConnStringValueTypePairArgs) ToConnStringValueTypePairOutput() ConnStringValueTypePairOutput {
	return i.ToConnStringValueTypePairOutputWithContext(context.Background())
}

func (i ConnStringValueTypePairArgs) ToConnStringValueTypePairOutputWithContext(ctx context.Context) ConnStringValueTypePairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnStringValueTypePairOutput)
}





type ConnStringValueTypePairMapInput interface {
	pulumi.Input

	ToConnStringValueTypePairMapOutput() ConnStringValueTypePairMapOutput
	ToConnStringValueTypePairMapOutputWithContext(context.Context) ConnStringValueTypePairMapOutput
}

type ConnStringValueTypePairMap map[string]ConnStringValueTypePairInput

func (ConnStringValueTypePairMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnStringValueTypePair)(nil)).Elem()
}

func (i ConnStringValueTypePairMap) ToConnStringValueTypePairMapOutput() ConnStringValueTypePairMapOutput {
	return i.ToConnStringValueTypePairMapOutputWithContext(context.Background())
}

func (i ConnStringValueTypePairMap) ToConnStringValueTypePairMapOutputWithContext(ctx context.Context) ConnStringValueTypePairMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnStringValueTypePairMapOutput)
}

type ConnStringValueTypePairOutput struct{ *pulumi.OutputState }

func (ConnStringValueTypePairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringValueTypePair)(nil)).Elem()
}

func (o ConnStringValueTypePairOutput) ToConnStringValueTypePairOutput() ConnStringValueTypePairOutput {
	return o
}

func (o ConnStringValueTypePairOutput) ToConnStringValueTypePairOutputWithContext(ctx context.Context) ConnStringValueTypePairOutput {
	return o
}

func (o ConnStringValueTypePairOutput) Type() DatabaseServerTypeOutput {
	return o.ApplyT(func(v ConnStringValueTypePair) DatabaseServerType { return v.Type }).(DatabaseServerTypeOutput)
}

func (o ConnStringValueTypePairOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnStringValueTypePair) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ConnStringValueTypePairMapOutput struct{ *pulumi.OutputState }

func (ConnStringValueTypePairMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnStringValueTypePair)(nil)).Elem()
}

func (o ConnStringValueTypePairMapOutput) ToConnStringValueTypePairMapOutput() ConnStringValueTypePairMapOutput {
	return o
}

func (o ConnStringValueTypePairMapOutput) ToConnStringValueTypePairMapOutputWithContext(ctx context.Context) ConnStringValueTypePairMapOutput {
	return o
}

func (o ConnStringValueTypePairMapOutput) MapIndex(k pulumi.StringInput) ConnStringValueTypePairOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ConnStringValueTypePair {
		return vs[0].(map[string]ConnStringValueTypePair)[vs[1].(string)]
	}).(ConnStringValueTypePairOutput)
}

type ConnStringValueTypePairResponse struct {
	Type  string  `pulumi:"type"`
	Value *string `pulumi:"value"`
}

type ConnStringValueTypePairResponseOutput struct{ *pulumi.OutputState }

func (ConnStringValueTypePairResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringValueTypePairResponse)(nil)).Elem()
}

func (o ConnStringValueTypePairResponseOutput) ToConnStringValueTypePairResponseOutput() ConnStringValueTypePairResponseOutput {
	return o
}

func (o ConnStringValueTypePairResponseOutput) ToConnStringValueTypePairResponseOutputWithContext(ctx context.Context) ConnStringValueTypePairResponseOutput {
	return o
}

func (o ConnStringValueTypePairResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ConnStringValueTypePairResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ConnStringValueTypePairResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnStringValueTypePairResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ConnStringValueTypePairResponseMapOutput struct{ *pulumi.OutputState }

func (ConnStringValueTypePairResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnStringValueTypePairResponse)(nil)).Elem()
}

func (o ConnStringValueTypePairResponseMapOutput) ToConnStringValueTypePairResponseMapOutput() ConnStringValueTypePairResponseMapOutput {
	return o
}

func (o ConnStringValueTypePairResponseMapOutput) ToConnStringValueTypePairResponseMapOutputWithContext(ctx context.Context) ConnStringValueTypePairResponseMapOutput {
	return o
}

func (o ConnStringValueTypePairResponseMapOutput) MapIndex(k pulumi.StringInput) ConnStringValueTypePairResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ConnStringValueTypePairResponse {
		return vs[0].(map[string]ConnStringValueTypePairResponse)[vs[1].(string)]
	}).(ConnStringValueTypePairResponseOutput)
}

type CorsSettings struct {
	AllowedOrigins []string `pulumi:"allowedOrigins"`
}





type CorsSettingsInput interface {
	pulumi.Input

	ToCorsSettingsOutput() CorsSettingsOutput
	ToCorsSettingsOutputWithContext(context.Context) CorsSettingsOutput
}

type CorsSettingsArgs struct {
	AllowedOrigins pulumi.StringArrayInput `pulumi:"allowedOrigins"`
}

func (CorsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsSettings)(nil)).Elem()
}

func (i CorsSettingsArgs) ToCorsSettingsOutput() CorsSettingsOutput {
	return i.ToCorsSettingsOutputWithContext(context.Background())
}

func (i CorsSettingsArgs) ToCorsSettingsOutputWithContext(ctx context.Context) CorsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsSettingsOutput)
}

func (i CorsSettingsArgs) ToCorsSettingsPtrOutput() CorsSettingsPtrOutput {
	return i.ToCorsSettingsPtrOutputWithContext(context.Background())
}

func (i CorsSettingsArgs) ToCorsSettingsPtrOutputWithContext(ctx context.Context) CorsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsSettingsOutput).ToCorsSettingsPtrOutputWithContext(ctx)
}









type CorsSettingsPtrInput interface {
	pulumi.Input

	ToCorsSettingsPtrOutput() CorsSettingsPtrOutput
	ToCorsSettingsPtrOutputWithContext(context.Context) CorsSettingsPtrOutput
}

type corsSettingsPtrType CorsSettingsArgs

func CorsSettingsPtr(v *CorsSettingsArgs) CorsSettingsPtrInput {
	return (*corsSettingsPtrType)(v)
}

func (*corsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CorsSettings)(nil)).Elem()
}

func (i *corsSettingsPtrType) ToCorsSettingsPtrOutput() CorsSettingsPtrOutput {
	return i.ToCorsSettingsPtrOutputWithContext(context.Background())
}

func (i *corsSettingsPtrType) ToCorsSettingsPtrOutputWithContext(ctx context.Context) CorsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsSettingsPtrOutput)
}

type CorsSettingsOutput struct{ *pulumi.OutputState }

func (CorsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsSettings)(nil)).Elem()
}

func (o CorsSettingsOutput) ToCorsSettingsOutput() CorsSettingsOutput {
	return o
}

func (o CorsSettingsOutput) ToCorsSettingsOutputWithContext(ctx context.Context) CorsSettingsOutput {
	return o
}

func (o CorsSettingsOutput) ToCorsSettingsPtrOutput() CorsSettingsPtrOutput {
	return o.ToCorsSettingsPtrOutputWithContext(context.Background())
}

func (o CorsSettingsOutput) ToCorsSettingsPtrOutputWithContext(ctx context.Context) CorsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CorsSettings) *CorsSettings {
		return &v
	}).(CorsSettingsPtrOutput)
}

func (o CorsSettingsOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsSettings) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

type CorsSettingsPtrOutput struct{ *pulumi.OutputState }

func (CorsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CorsSettings)(nil)).Elem()
}

func (o CorsSettingsPtrOutput) ToCorsSettingsPtrOutput() CorsSettingsPtrOutput {
	return o
}

func (o CorsSettingsPtrOutput) ToCorsSettingsPtrOutputWithContext(ctx context.Context) CorsSettingsPtrOutput {
	return o
}

func (o CorsSettingsPtrOutput) Elem() CorsSettingsOutput {
	return o.ApplyT(func(v *CorsSettings) CorsSettings {
		if v != nil {
			return *v
		}
		var ret CorsSettings
		return ret
	}).(CorsSettingsOutput)
}

func (o CorsSettingsPtrOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CorsSettings) []string {
		if v == nil {
			return nil
		}
		return v.AllowedOrigins
	}).(pulumi.StringArrayOutput)
}

type CorsSettingsResponse struct {
	AllowedOrigins []string `pulumi:"allowedOrigins"`
}

type CorsSettingsResponseOutput struct{ *pulumi.OutputState }

func (CorsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsSettingsResponse)(nil)).Elem()
}

func (o CorsSettingsResponseOutput) ToCorsSettingsResponseOutput() CorsSettingsResponseOutput {
	return o
}

func (o CorsSettingsResponseOutput) ToCorsSettingsResponseOutputWithContext(ctx context.Context) CorsSettingsResponseOutput {
	return o
}

func (o CorsSettingsResponseOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsSettingsResponse) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

type CorsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (CorsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CorsSettingsResponse)(nil)).Elem()
}

func (o CorsSettingsResponsePtrOutput) ToCorsSettingsResponsePtrOutput() CorsSettingsResponsePtrOutput {
	return o
}

func (o CorsSettingsResponsePtrOutput) ToCorsSettingsResponsePtrOutputWithContext(ctx context.Context) CorsSettingsResponsePtrOutput {
	return o
}

func (o CorsSettingsResponsePtrOutput) Elem() CorsSettingsResponseOutput {
	return o.ApplyT(func(v *CorsSettingsResponse) CorsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret CorsSettingsResponse
		return ret
	}).(CorsSettingsResponseOutput)
}

func (o CorsSettingsResponsePtrOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CorsSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedOrigins
	}).(pulumi.StringArrayOutput)
}

type DatabaseBackupSetting struct {
	ConnectionString     *string `pulumi:"connectionString"`
	ConnectionStringName *string `pulumi:"connectionStringName"`
	DatabaseType         *string `pulumi:"databaseType"`
	Name                 *string `pulumi:"name"`
}





type DatabaseBackupSettingInput interface {
	pulumi.Input

	ToDatabaseBackupSettingOutput() DatabaseBackupSettingOutput
	ToDatabaseBackupSettingOutputWithContext(context.Context) DatabaseBackupSettingOutput
}

type DatabaseBackupSettingArgs struct {
	ConnectionString     pulumi.StringPtrInput `pulumi:"connectionString"`
	ConnectionStringName pulumi.StringPtrInput `pulumi:"connectionStringName"`
	DatabaseType         pulumi.StringPtrInput `pulumi:"databaseType"`
	Name                 pulumi.StringPtrInput `pulumi:"name"`
}

func (DatabaseBackupSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseBackupSetting)(nil)).Elem()
}

func (i DatabaseBackupSettingArgs) ToDatabaseBackupSettingOutput() DatabaseBackupSettingOutput {
	return i.ToDatabaseBackupSettingOutputWithContext(context.Background())
}

func (i DatabaseBackupSettingArgs) ToDatabaseBackupSettingOutputWithContext(ctx context.Context) DatabaseBackupSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseBackupSettingOutput)
}





type DatabaseBackupSettingArrayInput interface {
	pulumi.Input

	ToDatabaseBackupSettingArrayOutput() DatabaseBackupSettingArrayOutput
	ToDatabaseBackupSettingArrayOutputWithContext(context.Context) DatabaseBackupSettingArrayOutput
}

type DatabaseBackupSettingArray []DatabaseBackupSettingInput

func (DatabaseBackupSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseBackupSetting)(nil)).Elem()
}

func (i DatabaseBackupSettingArray) ToDatabaseBackupSettingArrayOutput() DatabaseBackupSettingArrayOutput {
	return i.ToDatabaseBackupSettingArrayOutputWithContext(context.Background())
}

func (i DatabaseBackupSettingArray) ToDatabaseBackupSettingArrayOutputWithContext(ctx context.Context) DatabaseBackupSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseBackupSettingArrayOutput)
}

type DatabaseBackupSettingOutput struct{ *pulumi.OutputState }

func (DatabaseBackupSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseBackupSetting)(nil)).Elem()
}

func (o DatabaseBackupSettingOutput) ToDatabaseBackupSettingOutput() DatabaseBackupSettingOutput {
	return o
}

func (o DatabaseBackupSettingOutput) ToDatabaseBackupSettingOutputWithContext(ctx context.Context) DatabaseBackupSettingOutput {
	return o
}

func (o DatabaseBackupSettingOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSetting) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o DatabaseBackupSettingOutput) ConnectionStringName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSetting) *string { return v.ConnectionStringName }).(pulumi.StringPtrOutput)
}

func (o DatabaseBackupSettingOutput) DatabaseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSetting) *string { return v.DatabaseType }).(pulumi.StringPtrOutput)
}

func (o DatabaseBackupSettingOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSetting) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DatabaseBackupSettingArrayOutput struct{ *pulumi.OutputState }

func (DatabaseBackupSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseBackupSetting)(nil)).Elem()
}

func (o DatabaseBackupSettingArrayOutput) ToDatabaseBackupSettingArrayOutput() DatabaseBackupSettingArrayOutput {
	return o
}

func (o DatabaseBackupSettingArrayOutput) ToDatabaseBackupSettingArrayOutputWithContext(ctx context.Context) DatabaseBackupSettingArrayOutput {
	return o
}

func (o DatabaseBackupSettingArrayOutput) Index(i pulumi.IntInput) DatabaseBackupSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseBackupSetting {
		return vs[0].([]DatabaseBackupSetting)[vs[1].(int)]
	}).(DatabaseBackupSettingOutput)
}

type DatabaseBackupSettingResponse struct {
	ConnectionString     *string `pulumi:"connectionString"`
	ConnectionStringName *string `pulumi:"connectionStringName"`
	DatabaseType         *string `pulumi:"databaseType"`
	Name                 *string `pulumi:"name"`
}

type DatabaseBackupSettingResponseOutput struct{ *pulumi.OutputState }

func (DatabaseBackupSettingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseBackupSettingResponse)(nil)).Elem()
}

func (o DatabaseBackupSettingResponseOutput) ToDatabaseBackupSettingResponseOutput() DatabaseBackupSettingResponseOutput {
	return o
}

func (o DatabaseBackupSettingResponseOutput) ToDatabaseBackupSettingResponseOutputWithContext(ctx context.Context) DatabaseBackupSettingResponseOutput {
	return o
}

func (o DatabaseBackupSettingResponseOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSettingResponse) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o DatabaseBackupSettingResponseOutput) ConnectionStringName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSettingResponse) *string { return v.ConnectionStringName }).(pulumi.StringPtrOutput)
}

func (o DatabaseBackupSettingResponseOutput) DatabaseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSettingResponse) *string { return v.DatabaseType }).(pulumi.StringPtrOutput)
}

func (o DatabaseBackupSettingResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSettingResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DatabaseBackupSettingResponseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseBackupSettingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseBackupSettingResponse)(nil)).Elem()
}

func (o DatabaseBackupSettingResponseArrayOutput) ToDatabaseBackupSettingResponseArrayOutput() DatabaseBackupSettingResponseArrayOutput {
	return o
}

func (o DatabaseBackupSettingResponseArrayOutput) ToDatabaseBackupSettingResponseArrayOutputWithContext(ctx context.Context) DatabaseBackupSettingResponseArrayOutput {
	return o
}

func (o DatabaseBackupSettingResponseArrayOutput) Index(i pulumi.IntInput) DatabaseBackupSettingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseBackupSettingResponse {
		return vs[0].([]DatabaseBackupSettingResponse)[vs[1].(int)]
	}).(DatabaseBackupSettingResponseOutput)
}

type EnabledConfig struct {
	Enabled *bool `pulumi:"enabled"`
}





type EnabledConfigInput interface {
	pulumi.Input

	ToEnabledConfigOutput() EnabledConfigOutput
	ToEnabledConfigOutputWithContext(context.Context) EnabledConfigOutput
}

type EnabledConfigArgs struct {
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (EnabledConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledConfig)(nil)).Elem()
}

func (i EnabledConfigArgs) ToEnabledConfigOutput() EnabledConfigOutput {
	return i.ToEnabledConfigOutputWithContext(context.Background())
}

func (i EnabledConfigArgs) ToEnabledConfigOutputWithContext(ctx context.Context) EnabledConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledConfigOutput)
}

func (i EnabledConfigArgs) ToEnabledConfigPtrOutput() EnabledConfigPtrOutput {
	return i.ToEnabledConfigPtrOutputWithContext(context.Background())
}

func (i EnabledConfigArgs) ToEnabledConfigPtrOutputWithContext(ctx context.Context) EnabledConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledConfigOutput).ToEnabledConfigPtrOutputWithContext(ctx)
}









type EnabledConfigPtrInput interface {
	pulumi.Input

	ToEnabledConfigPtrOutput() EnabledConfigPtrOutput
	ToEnabledConfigPtrOutputWithContext(context.Context) EnabledConfigPtrOutput
}

type enabledConfigPtrType EnabledConfigArgs

func EnabledConfigPtr(v *EnabledConfigArgs) EnabledConfigPtrInput {
	return (*enabledConfigPtrType)(v)
}

func (*enabledConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledConfig)(nil)).Elem()
}

func (i *enabledConfigPtrType) ToEnabledConfigPtrOutput() EnabledConfigPtrOutput {
	return i.ToEnabledConfigPtrOutputWithContext(context.Background())
}

func (i *enabledConfigPtrType) ToEnabledConfigPtrOutputWithContext(ctx context.Context) EnabledConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledConfigPtrOutput)
}

type EnabledConfigOutput struct{ *pulumi.OutputState }

func (EnabledConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledConfig)(nil)).Elem()
}

func (o EnabledConfigOutput) ToEnabledConfigOutput() EnabledConfigOutput {
	return o
}

func (o EnabledConfigOutput) ToEnabledConfigOutputWithContext(ctx context.Context) EnabledConfigOutput {
	return o
}

func (o EnabledConfigOutput) ToEnabledConfigPtrOutput() EnabledConfigPtrOutput {
	return o.ToEnabledConfigPtrOutputWithContext(context.Background())
}

func (o EnabledConfigOutput) ToEnabledConfigPtrOutputWithContext(ctx context.Context) EnabledConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnabledConfig) *EnabledConfig {
		return &v
	}).(EnabledConfigPtrOutput)
}

func (o EnabledConfigOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EnabledConfig) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type EnabledConfigPtrOutput struct{ *pulumi.OutputState }

func (EnabledConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledConfig)(nil)).Elem()
}

func (o EnabledConfigPtrOutput) ToEnabledConfigPtrOutput() EnabledConfigPtrOutput {
	return o
}

func (o EnabledConfigPtrOutput) ToEnabledConfigPtrOutputWithContext(ctx context.Context) EnabledConfigPtrOutput {
	return o
}

func (o EnabledConfigPtrOutput) Elem() EnabledConfigOutput {
	return o.ApplyT(func(v *EnabledConfig) EnabledConfig {
		if v != nil {
			return *v
		}
		var ret EnabledConfig
		return ret
	}).(EnabledConfigOutput)
}

func (o EnabledConfigPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledConfig) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type EnabledConfigResponse struct {
	Enabled *bool `pulumi:"enabled"`
}

type EnabledConfigResponseOutput struct{ *pulumi.OutputState }

func (EnabledConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledConfigResponse)(nil)).Elem()
}

func (o EnabledConfigResponseOutput) ToEnabledConfigResponseOutput() EnabledConfigResponseOutput {
	return o
}

func (o EnabledConfigResponseOutput) ToEnabledConfigResponseOutputWithContext(ctx context.Context) EnabledConfigResponseOutput {
	return o
}

func (o EnabledConfigResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EnabledConfigResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type EnabledConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (EnabledConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledConfigResponse)(nil)).Elem()
}

func (o EnabledConfigResponsePtrOutput) ToEnabledConfigResponsePtrOutput() EnabledConfigResponsePtrOutput {
	return o
}

func (o EnabledConfigResponsePtrOutput) ToEnabledConfigResponsePtrOutputWithContext(ctx context.Context) EnabledConfigResponsePtrOutput {
	return o
}

func (o EnabledConfigResponsePtrOutput) Elem() EnabledConfigResponseOutput {
	return o.ApplyT(func(v *EnabledConfigResponse) EnabledConfigResponse {
		if v != nil {
			return *v
		}
		var ret EnabledConfigResponse
		return ret
	}).(EnabledConfigResponseOutput)
}

func (o EnabledConfigResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type Experiments struct {
	RampUpRules []RampUpRule `pulumi:"rampUpRules"`
}





type ExperimentsInput interface {
	pulumi.Input

	ToExperimentsOutput() ExperimentsOutput
	ToExperimentsOutputWithContext(context.Context) ExperimentsOutput
}

type ExperimentsArgs struct {
	RampUpRules RampUpRuleArrayInput `pulumi:"rampUpRules"`
}

func (ExperimentsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Experiments)(nil)).Elem()
}

func (i ExperimentsArgs) ToExperimentsOutput() ExperimentsOutput {
	return i.ToExperimentsOutputWithContext(context.Background())
}

func (i ExperimentsArgs) ToExperimentsOutputWithContext(ctx context.Context) ExperimentsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentsOutput)
}

func (i ExperimentsArgs) ToExperimentsPtrOutput() ExperimentsPtrOutput {
	return i.ToExperimentsPtrOutputWithContext(context.Background())
}

func (i ExperimentsArgs) ToExperimentsPtrOutputWithContext(ctx context.Context) ExperimentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentsOutput).ToExperimentsPtrOutputWithContext(ctx)
}









type ExperimentsPtrInput interface {
	pulumi.Input

	ToExperimentsPtrOutput() ExperimentsPtrOutput
	ToExperimentsPtrOutputWithContext(context.Context) ExperimentsPtrOutput
}

type experimentsPtrType ExperimentsArgs

func ExperimentsPtr(v *ExperimentsArgs) ExperimentsPtrInput {
	return (*experimentsPtrType)(v)
}

func (*experimentsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Experiments)(nil)).Elem()
}

func (i *experimentsPtrType) ToExperimentsPtrOutput() ExperimentsPtrOutput {
	return i.ToExperimentsPtrOutputWithContext(context.Background())
}

func (i *experimentsPtrType) ToExperimentsPtrOutputWithContext(ctx context.Context) ExperimentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentsPtrOutput)
}

type ExperimentsOutput struct{ *pulumi.OutputState }

func (ExperimentsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Experiments)(nil)).Elem()
}

func (o ExperimentsOutput) ToExperimentsOutput() ExperimentsOutput {
	return o
}

func (o ExperimentsOutput) ToExperimentsOutputWithContext(ctx context.Context) ExperimentsOutput {
	return o
}

func (o ExperimentsOutput) ToExperimentsPtrOutput() ExperimentsPtrOutput {
	return o.ToExperimentsPtrOutputWithContext(context.Background())
}

func (o ExperimentsOutput) ToExperimentsPtrOutputWithContext(ctx context.Context) ExperimentsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Experiments) *Experiments {
		return &v
	}).(ExperimentsPtrOutput)
}

func (o ExperimentsOutput) RampUpRules() RampUpRuleArrayOutput {
	return o.ApplyT(func(v Experiments) []RampUpRule { return v.RampUpRules }).(RampUpRuleArrayOutput)
}

type ExperimentsPtrOutput struct{ *pulumi.OutputState }

func (ExperimentsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Experiments)(nil)).Elem()
}

func (o ExperimentsPtrOutput) ToExperimentsPtrOutput() ExperimentsPtrOutput {
	return o
}

func (o ExperimentsPtrOutput) ToExperimentsPtrOutputWithContext(ctx context.Context) ExperimentsPtrOutput {
	return o
}

func (o ExperimentsPtrOutput) Elem() ExperimentsOutput {
	return o.ApplyT(func(v *Experiments) Experiments {
		if v != nil {
			return *v
		}
		var ret Experiments
		return ret
	}).(ExperimentsOutput)
}

func (o ExperimentsPtrOutput) RampUpRules() RampUpRuleArrayOutput {
	return o.ApplyT(func(v *Experiments) []RampUpRule {
		if v == nil {
			return nil
		}
		return v.RampUpRules
	}).(RampUpRuleArrayOutput)
}

type ExperimentsResponse struct {
	RampUpRules []RampUpRuleResponse `pulumi:"rampUpRules"`
}

type ExperimentsResponseOutput struct{ *pulumi.OutputState }

func (ExperimentsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExperimentsResponse)(nil)).Elem()
}

func (o ExperimentsResponseOutput) ToExperimentsResponseOutput() ExperimentsResponseOutput {
	return o
}

func (o ExperimentsResponseOutput) ToExperimentsResponseOutputWithContext(ctx context.Context) ExperimentsResponseOutput {
	return o
}

func (o ExperimentsResponseOutput) RampUpRules() RampUpRuleResponseArrayOutput {
	return o.ApplyT(func(v ExperimentsResponse) []RampUpRuleResponse { return v.RampUpRules }).(RampUpRuleResponseArrayOutput)
}

type ExperimentsResponsePtrOutput struct{ *pulumi.OutputState }

func (ExperimentsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExperimentsResponse)(nil)).Elem()
}

func (o ExperimentsResponsePtrOutput) ToExperimentsResponsePtrOutput() ExperimentsResponsePtrOutput {
	return o
}

func (o ExperimentsResponsePtrOutput) ToExperimentsResponsePtrOutputWithContext(ctx context.Context) ExperimentsResponsePtrOutput {
	return o
}

func (o ExperimentsResponsePtrOutput) Elem() ExperimentsResponseOutput {
	return o.ApplyT(func(v *ExperimentsResponse) ExperimentsResponse {
		if v != nil {
			return *v
		}
		var ret ExperimentsResponse
		return ret
	}).(ExperimentsResponseOutput)
}

func (o ExperimentsResponsePtrOutput) RampUpRules() RampUpRuleResponseArrayOutput {
	return o.ApplyT(func(v *ExperimentsResponse) []RampUpRuleResponse {
		if v == nil {
			return nil
		}
		return v.RampUpRules
	}).(RampUpRuleResponseArrayOutput)
}

type FileSystemApplicationLogsConfig struct {
	Level *LogLevel `pulumi:"level"`
}





type FileSystemApplicationLogsConfigInput interface {
	pulumi.Input

	ToFileSystemApplicationLogsConfigOutput() FileSystemApplicationLogsConfigOutput
	ToFileSystemApplicationLogsConfigOutputWithContext(context.Context) FileSystemApplicationLogsConfigOutput
}

type FileSystemApplicationLogsConfigArgs struct {
	Level LogLevelPtrInput `pulumi:"level"`
}

func (FileSystemApplicationLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemApplicationLogsConfig)(nil)).Elem()
}

func (i FileSystemApplicationLogsConfigArgs) ToFileSystemApplicationLogsConfigOutput() FileSystemApplicationLogsConfigOutput {
	return i.ToFileSystemApplicationLogsConfigOutputWithContext(context.Background())
}

func (i FileSystemApplicationLogsConfigArgs) ToFileSystemApplicationLogsConfigOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemApplicationLogsConfigOutput)
}

func (i FileSystemApplicationLogsConfigArgs) ToFileSystemApplicationLogsConfigPtrOutput() FileSystemApplicationLogsConfigPtrOutput {
	return i.ToFileSystemApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i FileSystemApplicationLogsConfigArgs) ToFileSystemApplicationLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemApplicationLogsConfigOutput).ToFileSystemApplicationLogsConfigPtrOutputWithContext(ctx)
}









type FileSystemApplicationLogsConfigPtrInput interface {
	pulumi.Input

	ToFileSystemApplicationLogsConfigPtrOutput() FileSystemApplicationLogsConfigPtrOutput
	ToFileSystemApplicationLogsConfigPtrOutputWithContext(context.Context) FileSystemApplicationLogsConfigPtrOutput
}

type fileSystemApplicationLogsConfigPtrType FileSystemApplicationLogsConfigArgs

func FileSystemApplicationLogsConfigPtr(v *FileSystemApplicationLogsConfigArgs) FileSystemApplicationLogsConfigPtrInput {
	return (*fileSystemApplicationLogsConfigPtrType)(v)
}

func (*fileSystemApplicationLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemApplicationLogsConfig)(nil)).Elem()
}

func (i *fileSystemApplicationLogsConfigPtrType) ToFileSystemApplicationLogsConfigPtrOutput() FileSystemApplicationLogsConfigPtrOutput {
	return i.ToFileSystemApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i *fileSystemApplicationLogsConfigPtrType) ToFileSystemApplicationLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemApplicationLogsConfigPtrOutput)
}

type FileSystemApplicationLogsConfigOutput struct{ *pulumi.OutputState }

func (FileSystemApplicationLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemApplicationLogsConfig)(nil)).Elem()
}

func (o FileSystemApplicationLogsConfigOutput) ToFileSystemApplicationLogsConfigOutput() FileSystemApplicationLogsConfigOutput {
	return o
}

func (o FileSystemApplicationLogsConfigOutput) ToFileSystemApplicationLogsConfigOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigOutput {
	return o
}

func (o FileSystemApplicationLogsConfigOutput) ToFileSystemApplicationLogsConfigPtrOutput() FileSystemApplicationLogsConfigPtrOutput {
	return o.ToFileSystemApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (o FileSystemApplicationLogsConfigOutput) ToFileSystemApplicationLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FileSystemApplicationLogsConfig) *FileSystemApplicationLogsConfig {
		return &v
	}).(FileSystemApplicationLogsConfigPtrOutput)
}

func (o FileSystemApplicationLogsConfigOutput) Level() LogLevelPtrOutput {
	return o.ApplyT(func(v FileSystemApplicationLogsConfig) *LogLevel { return v.Level }).(LogLevelPtrOutput)
}

type FileSystemApplicationLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (FileSystemApplicationLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemApplicationLogsConfig)(nil)).Elem()
}

func (o FileSystemApplicationLogsConfigPtrOutput) ToFileSystemApplicationLogsConfigPtrOutput() FileSystemApplicationLogsConfigPtrOutput {
	return o
}

func (o FileSystemApplicationLogsConfigPtrOutput) ToFileSystemApplicationLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigPtrOutput {
	return o
}

func (o FileSystemApplicationLogsConfigPtrOutput) Elem() FileSystemApplicationLogsConfigOutput {
	return o.ApplyT(func(v *FileSystemApplicationLogsConfig) FileSystemApplicationLogsConfig {
		if v != nil {
			return *v
		}
		var ret FileSystemApplicationLogsConfig
		return ret
	}).(FileSystemApplicationLogsConfigOutput)
}

func (o FileSystemApplicationLogsConfigPtrOutput) Level() LogLevelPtrOutput {
	return o.ApplyT(func(v *FileSystemApplicationLogsConfig) *LogLevel {
		if v == nil {
			return nil
		}
		return v.Level
	}).(LogLevelPtrOutput)
}

type FileSystemApplicationLogsConfigResponse struct {
	Level *string `pulumi:"level"`
}

type FileSystemApplicationLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (FileSystemApplicationLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemApplicationLogsConfigResponse)(nil)).Elem()
}

func (o FileSystemApplicationLogsConfigResponseOutput) ToFileSystemApplicationLogsConfigResponseOutput() FileSystemApplicationLogsConfigResponseOutput {
	return o
}

func (o FileSystemApplicationLogsConfigResponseOutput) ToFileSystemApplicationLogsConfigResponseOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigResponseOutput {
	return o
}

func (o FileSystemApplicationLogsConfigResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileSystemApplicationLogsConfigResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

type FileSystemApplicationLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (FileSystemApplicationLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemApplicationLogsConfigResponse)(nil)).Elem()
}

func (o FileSystemApplicationLogsConfigResponsePtrOutput) ToFileSystemApplicationLogsConfigResponsePtrOutput() FileSystemApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o FileSystemApplicationLogsConfigResponsePtrOutput) ToFileSystemApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o FileSystemApplicationLogsConfigResponsePtrOutput) Elem() FileSystemApplicationLogsConfigResponseOutput {
	return o.ApplyT(func(v *FileSystemApplicationLogsConfigResponse) FileSystemApplicationLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret FileSystemApplicationLogsConfigResponse
		return ret
	}).(FileSystemApplicationLogsConfigResponseOutput)
}

func (o FileSystemApplicationLogsConfigResponsePtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileSystemApplicationLogsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Level
	}).(pulumi.StringPtrOutput)
}

type FileSystemHttpLogsConfig struct {
	Enabled         *bool `pulumi:"enabled"`
	RetentionInDays *int  `pulumi:"retentionInDays"`
	RetentionInMb   *int  `pulumi:"retentionInMb"`
}





type FileSystemHttpLogsConfigInput interface {
	pulumi.Input

	ToFileSystemHttpLogsConfigOutput() FileSystemHttpLogsConfigOutput
	ToFileSystemHttpLogsConfigOutputWithContext(context.Context) FileSystemHttpLogsConfigOutput
}

type FileSystemHttpLogsConfigArgs struct {
	Enabled         pulumi.BoolPtrInput `pulumi:"enabled"`
	RetentionInDays pulumi.IntPtrInput  `pulumi:"retentionInDays"`
	RetentionInMb   pulumi.IntPtrInput  `pulumi:"retentionInMb"`
}

func (FileSystemHttpLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemHttpLogsConfig)(nil)).Elem()
}

func (i FileSystemHttpLogsConfigArgs) ToFileSystemHttpLogsConfigOutput() FileSystemHttpLogsConfigOutput {
	return i.ToFileSystemHttpLogsConfigOutputWithContext(context.Background())
}

func (i FileSystemHttpLogsConfigArgs) ToFileSystemHttpLogsConfigOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemHttpLogsConfigOutput)
}

func (i FileSystemHttpLogsConfigArgs) ToFileSystemHttpLogsConfigPtrOutput() FileSystemHttpLogsConfigPtrOutput {
	return i.ToFileSystemHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (i FileSystemHttpLogsConfigArgs) ToFileSystemHttpLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemHttpLogsConfigOutput).ToFileSystemHttpLogsConfigPtrOutputWithContext(ctx)
}









type FileSystemHttpLogsConfigPtrInput interface {
	pulumi.Input

	ToFileSystemHttpLogsConfigPtrOutput() FileSystemHttpLogsConfigPtrOutput
	ToFileSystemHttpLogsConfigPtrOutputWithContext(context.Context) FileSystemHttpLogsConfigPtrOutput
}

type fileSystemHttpLogsConfigPtrType FileSystemHttpLogsConfigArgs

func FileSystemHttpLogsConfigPtr(v *FileSystemHttpLogsConfigArgs) FileSystemHttpLogsConfigPtrInput {
	return (*fileSystemHttpLogsConfigPtrType)(v)
}

func (*fileSystemHttpLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemHttpLogsConfig)(nil)).Elem()
}

func (i *fileSystemHttpLogsConfigPtrType) ToFileSystemHttpLogsConfigPtrOutput() FileSystemHttpLogsConfigPtrOutput {
	return i.ToFileSystemHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (i *fileSystemHttpLogsConfigPtrType) ToFileSystemHttpLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemHttpLogsConfigPtrOutput)
}

type FileSystemHttpLogsConfigOutput struct{ *pulumi.OutputState }

func (FileSystemHttpLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemHttpLogsConfig)(nil)).Elem()
}

func (o FileSystemHttpLogsConfigOutput) ToFileSystemHttpLogsConfigOutput() FileSystemHttpLogsConfigOutput {
	return o
}

func (o FileSystemHttpLogsConfigOutput) ToFileSystemHttpLogsConfigOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigOutput {
	return o
}

func (o FileSystemHttpLogsConfigOutput) ToFileSystemHttpLogsConfigPtrOutput() FileSystemHttpLogsConfigPtrOutput {
	return o.ToFileSystemHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (o FileSystemHttpLogsConfigOutput) ToFileSystemHttpLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FileSystemHttpLogsConfig) *FileSystemHttpLogsConfig {
		return &v
	}).(FileSystemHttpLogsConfigPtrOutput)
}

func (o FileSystemHttpLogsConfigOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FileSystemHttpLogsConfig) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o FileSystemHttpLogsConfigOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FileSystemHttpLogsConfig) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o FileSystemHttpLogsConfigOutput) RetentionInMb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FileSystemHttpLogsConfig) *int { return v.RetentionInMb }).(pulumi.IntPtrOutput)
}

type FileSystemHttpLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (FileSystemHttpLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemHttpLogsConfig)(nil)).Elem()
}

func (o FileSystemHttpLogsConfigPtrOutput) ToFileSystemHttpLogsConfigPtrOutput() FileSystemHttpLogsConfigPtrOutput {
	return o
}

func (o FileSystemHttpLogsConfigPtrOutput) ToFileSystemHttpLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigPtrOutput {
	return o
}

func (o FileSystemHttpLogsConfigPtrOutput) Elem() FileSystemHttpLogsConfigOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfig) FileSystemHttpLogsConfig {
		if v != nil {
			return *v
		}
		var ret FileSystemHttpLogsConfig
		return ret
	}).(FileSystemHttpLogsConfigOutput)
}

func (o FileSystemHttpLogsConfigPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfig) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o FileSystemHttpLogsConfigPtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfig) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o FileSystemHttpLogsConfigPtrOutput) RetentionInMb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfig) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInMb
	}).(pulumi.IntPtrOutput)
}

type FileSystemHttpLogsConfigResponse struct {
	Enabled         *bool `pulumi:"enabled"`
	RetentionInDays *int  `pulumi:"retentionInDays"`
	RetentionInMb   *int  `pulumi:"retentionInMb"`
}

type FileSystemHttpLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (FileSystemHttpLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemHttpLogsConfigResponse)(nil)).Elem()
}

func (o FileSystemHttpLogsConfigResponseOutput) ToFileSystemHttpLogsConfigResponseOutput() FileSystemHttpLogsConfigResponseOutput {
	return o
}

func (o FileSystemHttpLogsConfigResponseOutput) ToFileSystemHttpLogsConfigResponseOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigResponseOutput {
	return o
}

func (o FileSystemHttpLogsConfigResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FileSystemHttpLogsConfigResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o FileSystemHttpLogsConfigResponseOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FileSystemHttpLogsConfigResponse) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o FileSystemHttpLogsConfigResponseOutput) RetentionInMb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FileSystemHttpLogsConfigResponse) *int { return v.RetentionInMb }).(pulumi.IntPtrOutput)
}

type FileSystemHttpLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (FileSystemHttpLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemHttpLogsConfigResponse)(nil)).Elem()
}

func (o FileSystemHttpLogsConfigResponsePtrOutput) ToFileSystemHttpLogsConfigResponsePtrOutput() FileSystemHttpLogsConfigResponsePtrOutput {
	return o
}

func (o FileSystemHttpLogsConfigResponsePtrOutput) ToFileSystemHttpLogsConfigResponsePtrOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigResponsePtrOutput {
	return o
}

func (o FileSystemHttpLogsConfigResponsePtrOutput) Elem() FileSystemHttpLogsConfigResponseOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfigResponse) FileSystemHttpLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret FileSystemHttpLogsConfigResponse
		return ret
	}).(FileSystemHttpLogsConfigResponseOutput)
}

func (o FileSystemHttpLogsConfigResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o FileSystemHttpLogsConfigResponsePtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o FileSystemHttpLogsConfigResponsePtrOutput) RetentionInMb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInMb
	}).(pulumi.IntPtrOutput)
}

type HandlerMapping struct {
	Arguments       *string `pulumi:"arguments"`
	Extension       *string `pulumi:"extension"`
	ScriptProcessor *string `pulumi:"scriptProcessor"`
}





type HandlerMappingInput interface {
	pulumi.Input

	ToHandlerMappingOutput() HandlerMappingOutput
	ToHandlerMappingOutputWithContext(context.Context) HandlerMappingOutput
}

type HandlerMappingArgs struct {
	Arguments       pulumi.StringPtrInput `pulumi:"arguments"`
	Extension       pulumi.StringPtrInput `pulumi:"extension"`
	ScriptProcessor pulumi.StringPtrInput `pulumi:"scriptProcessor"`
}

func (HandlerMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HandlerMapping)(nil)).Elem()
}

func (i HandlerMappingArgs) ToHandlerMappingOutput() HandlerMappingOutput {
	return i.ToHandlerMappingOutputWithContext(context.Background())
}

func (i HandlerMappingArgs) ToHandlerMappingOutputWithContext(ctx context.Context) HandlerMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HandlerMappingOutput)
}





type HandlerMappingArrayInput interface {
	pulumi.Input

	ToHandlerMappingArrayOutput() HandlerMappingArrayOutput
	ToHandlerMappingArrayOutputWithContext(context.Context) HandlerMappingArrayOutput
}

type HandlerMappingArray []HandlerMappingInput

func (HandlerMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HandlerMapping)(nil)).Elem()
}

func (i HandlerMappingArray) ToHandlerMappingArrayOutput() HandlerMappingArrayOutput {
	return i.ToHandlerMappingArrayOutputWithContext(context.Background())
}

func (i HandlerMappingArray) ToHandlerMappingArrayOutputWithContext(ctx context.Context) HandlerMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HandlerMappingArrayOutput)
}

type HandlerMappingOutput struct{ *pulumi.OutputState }

func (HandlerMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HandlerMapping)(nil)).Elem()
}

func (o HandlerMappingOutput) ToHandlerMappingOutput() HandlerMappingOutput {
	return o
}

func (o HandlerMappingOutput) ToHandlerMappingOutputWithContext(ctx context.Context) HandlerMappingOutput {
	return o
}

func (o HandlerMappingOutput) Arguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HandlerMapping) *string { return v.Arguments }).(pulumi.StringPtrOutput)
}

func (o HandlerMappingOutput) Extension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HandlerMapping) *string { return v.Extension }).(pulumi.StringPtrOutput)
}

func (o HandlerMappingOutput) ScriptProcessor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HandlerMapping) *string { return v.ScriptProcessor }).(pulumi.StringPtrOutput)
}

type HandlerMappingArrayOutput struct{ *pulumi.OutputState }

func (HandlerMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HandlerMapping)(nil)).Elem()
}

func (o HandlerMappingArrayOutput) ToHandlerMappingArrayOutput() HandlerMappingArrayOutput {
	return o
}

func (o HandlerMappingArrayOutput) ToHandlerMappingArrayOutputWithContext(ctx context.Context) HandlerMappingArrayOutput {
	return o
}

func (o HandlerMappingArrayOutput) Index(i pulumi.IntInput) HandlerMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HandlerMapping {
		return vs[0].([]HandlerMapping)[vs[1].(int)]
	}).(HandlerMappingOutput)
}

type HandlerMappingResponse struct {
	Arguments       *string `pulumi:"arguments"`
	Extension       *string `pulumi:"extension"`
	ScriptProcessor *string `pulumi:"scriptProcessor"`
}

type HandlerMappingResponseOutput struct{ *pulumi.OutputState }

func (HandlerMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HandlerMappingResponse)(nil)).Elem()
}

func (o HandlerMappingResponseOutput) ToHandlerMappingResponseOutput() HandlerMappingResponseOutput {
	return o
}

func (o HandlerMappingResponseOutput) ToHandlerMappingResponseOutputWithContext(ctx context.Context) HandlerMappingResponseOutput {
	return o
}

func (o HandlerMappingResponseOutput) Arguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HandlerMappingResponse) *string { return v.Arguments }).(pulumi.StringPtrOutput)
}

func (o HandlerMappingResponseOutput) Extension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HandlerMappingResponse) *string { return v.Extension }).(pulumi.StringPtrOutput)
}

func (o HandlerMappingResponseOutput) ScriptProcessor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HandlerMappingResponse) *string { return v.ScriptProcessor }).(pulumi.StringPtrOutput)
}

type HandlerMappingResponseArrayOutput struct{ *pulumi.OutputState }

func (HandlerMappingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HandlerMappingResponse)(nil)).Elem()
}

func (o HandlerMappingResponseArrayOutput) ToHandlerMappingResponseArrayOutput() HandlerMappingResponseArrayOutput {
	return o
}

func (o HandlerMappingResponseArrayOutput) ToHandlerMappingResponseArrayOutputWithContext(ctx context.Context) HandlerMappingResponseArrayOutput {
	return o
}

func (o HandlerMappingResponseArrayOutput) Index(i pulumi.IntInput) HandlerMappingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HandlerMappingResponse {
		return vs[0].([]HandlerMappingResponse)[vs[1].(int)]
	}).(HandlerMappingResponseOutput)
}

type HostNameSslState struct {
	Name       *string  `pulumi:"name"`
	SslState   SslState `pulumi:"sslState"`
	Thumbprint *string  `pulumi:"thumbprint"`
	ToUpdate   *bool    `pulumi:"toUpdate"`
	VirtualIP  *string  `pulumi:"virtualIP"`
}





type HostNameSslStateInput interface {
	pulumi.Input

	ToHostNameSslStateOutput() HostNameSslStateOutput
	ToHostNameSslStateOutputWithContext(context.Context) HostNameSslStateOutput
}

type HostNameSslStateArgs struct {
	Name       pulumi.StringPtrInput `pulumi:"name"`
	SslState   SslStateInput         `pulumi:"sslState"`
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
	ToUpdate   pulumi.BoolPtrInput   `pulumi:"toUpdate"`
	VirtualIP  pulumi.StringPtrInput `pulumi:"virtualIP"`
}

func (HostNameSslStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameSslState)(nil)).Elem()
}

func (i HostNameSslStateArgs) ToHostNameSslStateOutput() HostNameSslStateOutput {
	return i.ToHostNameSslStateOutputWithContext(context.Background())
}

func (i HostNameSslStateArgs) ToHostNameSslStateOutputWithContext(ctx context.Context) HostNameSslStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostNameSslStateOutput)
}





type HostNameSslStateArrayInput interface {
	pulumi.Input

	ToHostNameSslStateArrayOutput() HostNameSslStateArrayOutput
	ToHostNameSslStateArrayOutputWithContext(context.Context) HostNameSslStateArrayOutput
}

type HostNameSslStateArray []HostNameSslStateInput

func (HostNameSslStateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostNameSslState)(nil)).Elem()
}

func (i HostNameSslStateArray) ToHostNameSslStateArrayOutput() HostNameSslStateArrayOutput {
	return i.ToHostNameSslStateArrayOutputWithContext(context.Background())
}

func (i HostNameSslStateArray) ToHostNameSslStateArrayOutputWithContext(ctx context.Context) HostNameSslStateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostNameSslStateArrayOutput)
}

type HostNameSslStateOutput struct{ *pulumi.OutputState }

func (HostNameSslStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameSslState)(nil)).Elem()
}

func (o HostNameSslStateOutput) ToHostNameSslStateOutput() HostNameSslStateOutput {
	return o
}

func (o HostNameSslStateOutput) ToHostNameSslStateOutputWithContext(ctx context.Context) HostNameSslStateOutput {
	return o
}

func (o HostNameSslStateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslState) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HostNameSslStateOutput) SslState() SslStateOutput {
	return o.ApplyT(func(v HostNameSslState) SslState { return v.SslState }).(SslStateOutput)
}

func (o HostNameSslStateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslState) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func (o HostNameSslStateOutput) ToUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HostNameSslState) *bool { return v.ToUpdate }).(pulumi.BoolPtrOutput)
}

func (o HostNameSslStateOutput) VirtualIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslState) *string { return v.VirtualIP }).(pulumi.StringPtrOutput)
}

type HostNameSslStateArrayOutput struct{ *pulumi.OutputState }

func (HostNameSslStateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostNameSslState)(nil)).Elem()
}

func (o HostNameSslStateArrayOutput) ToHostNameSslStateArrayOutput() HostNameSslStateArrayOutput {
	return o
}

func (o HostNameSslStateArrayOutput) ToHostNameSslStateArrayOutputWithContext(ctx context.Context) HostNameSslStateArrayOutput {
	return o
}

func (o HostNameSslStateArrayOutput) Index(i pulumi.IntInput) HostNameSslStateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostNameSslState {
		return vs[0].([]HostNameSslState)[vs[1].(int)]
	}).(HostNameSslStateOutput)
}

type HostNameSslStateResponse struct {
	Name       *string `pulumi:"name"`
	SslState   string  `pulumi:"sslState"`
	Thumbprint *string `pulumi:"thumbprint"`
	ToUpdate   *bool   `pulumi:"toUpdate"`
	VirtualIP  *string `pulumi:"virtualIP"`
}

type HostNameSslStateResponseOutput struct{ *pulumi.OutputState }

func (HostNameSslStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameSslStateResponse)(nil)).Elem()
}

func (o HostNameSslStateResponseOutput) ToHostNameSslStateResponseOutput() HostNameSslStateResponseOutput {
	return o
}

func (o HostNameSslStateResponseOutput) ToHostNameSslStateResponseOutputWithContext(ctx context.Context) HostNameSslStateResponseOutput {
	return o
}

func (o HostNameSslStateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslStateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HostNameSslStateResponseOutput) SslState() pulumi.StringOutput {
	return o.ApplyT(func(v HostNameSslStateResponse) string { return v.SslState }).(pulumi.StringOutput)
}

func (o HostNameSslStateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslStateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func (o HostNameSslStateResponseOutput) ToUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HostNameSslStateResponse) *bool { return v.ToUpdate }).(pulumi.BoolPtrOutput)
}

func (o HostNameSslStateResponseOutput) VirtualIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslStateResponse) *string { return v.VirtualIP }).(pulumi.StringPtrOutput)
}

type HostNameSslStateResponseArrayOutput struct{ *pulumi.OutputState }

func (HostNameSslStateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostNameSslStateResponse)(nil)).Elem()
}

func (o HostNameSslStateResponseArrayOutput) ToHostNameSslStateResponseArrayOutput() HostNameSslStateResponseArrayOutput {
	return o
}

func (o HostNameSslStateResponseArrayOutput) ToHostNameSslStateResponseArrayOutputWithContext(ctx context.Context) HostNameSslStateResponseArrayOutput {
	return o
}

func (o HostNameSslStateResponseArrayOutput) Index(i pulumi.IntInput) HostNameSslStateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostNameSslStateResponse {
		return vs[0].([]HostNameSslStateResponse)[vs[1].(int)]
	}).(HostNameSslStateResponseOutput)
}

type HostingEnvironmentProfile struct {
	Id   *string `pulumi:"id"`
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type HostingEnvironmentProfileInput interface {
	pulumi.Input

	ToHostingEnvironmentProfileOutput() HostingEnvironmentProfileOutput
	ToHostingEnvironmentProfileOutputWithContext(context.Context) HostingEnvironmentProfileOutput
}

type HostingEnvironmentProfileArgs struct {
	Id   pulumi.StringPtrInput `pulumi:"id"`
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (HostingEnvironmentProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentProfile)(nil)).Elem()
}

func (i HostingEnvironmentProfileArgs) ToHostingEnvironmentProfileOutput() HostingEnvironmentProfileOutput {
	return i.ToHostingEnvironmentProfileOutputWithContext(context.Background())
}

func (i HostingEnvironmentProfileArgs) ToHostingEnvironmentProfileOutputWithContext(ctx context.Context) HostingEnvironmentProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentProfileOutput)
}

func (i HostingEnvironmentProfileArgs) ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput {
	return i.ToHostingEnvironmentProfilePtrOutputWithContext(context.Background())
}

func (i HostingEnvironmentProfileArgs) ToHostingEnvironmentProfilePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentProfileOutput).ToHostingEnvironmentProfilePtrOutputWithContext(ctx)
}









type HostingEnvironmentProfilePtrInput interface {
	pulumi.Input

	ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput
	ToHostingEnvironmentProfilePtrOutputWithContext(context.Context) HostingEnvironmentProfilePtrOutput
}

type hostingEnvironmentProfilePtrType HostingEnvironmentProfileArgs

func HostingEnvironmentProfilePtr(v *HostingEnvironmentProfileArgs) HostingEnvironmentProfilePtrInput {
	return (*hostingEnvironmentProfilePtrType)(v)
}

func (*hostingEnvironmentProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingEnvironmentProfile)(nil)).Elem()
}

func (i *hostingEnvironmentProfilePtrType) ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput {
	return i.ToHostingEnvironmentProfilePtrOutputWithContext(context.Background())
}

func (i *hostingEnvironmentProfilePtrType) ToHostingEnvironmentProfilePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentProfilePtrOutput)
}

type HostingEnvironmentProfileOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentProfile)(nil)).Elem()
}

func (o HostingEnvironmentProfileOutput) ToHostingEnvironmentProfileOutput() HostingEnvironmentProfileOutput {
	return o
}

func (o HostingEnvironmentProfileOutput) ToHostingEnvironmentProfileOutputWithContext(ctx context.Context) HostingEnvironmentProfileOutput {
	return o
}

func (o HostingEnvironmentProfileOutput) ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput {
	return o.ToHostingEnvironmentProfilePtrOutputWithContext(context.Background())
}

func (o HostingEnvironmentProfileOutput) ToHostingEnvironmentProfilePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostingEnvironmentProfile) *HostingEnvironmentProfile {
		return &v
	}).(HostingEnvironmentProfilePtrOutput)
}

func (o HostingEnvironmentProfileOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentProfile) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfileOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentProfile) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfileOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentProfile) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type HostingEnvironmentProfilePtrOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingEnvironmentProfile)(nil)).Elem()
}

func (o HostingEnvironmentProfilePtrOutput) ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput {
	return o
}

func (o HostingEnvironmentProfilePtrOutput) ToHostingEnvironmentProfilePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfilePtrOutput {
	return o
}

func (o HostingEnvironmentProfilePtrOutput) Elem() HostingEnvironmentProfileOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfile) HostingEnvironmentProfile {
		if v != nil {
			return *v
		}
		var ret HostingEnvironmentProfile
		return ret
	}).(HostingEnvironmentProfileOutput)
}

func (o HostingEnvironmentProfilePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfile) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfilePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfile) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfilePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfile) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type HostingEnvironmentProfileResponse struct {
	Id   *string `pulumi:"id"`
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type HostingEnvironmentProfileResponseOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentProfileResponse)(nil)).Elem()
}

func (o HostingEnvironmentProfileResponseOutput) ToHostingEnvironmentProfileResponseOutput() HostingEnvironmentProfileResponseOutput {
	return o
}

func (o HostingEnvironmentProfileResponseOutput) ToHostingEnvironmentProfileResponseOutputWithContext(ctx context.Context) HostingEnvironmentProfileResponseOutput {
	return o
}

func (o HostingEnvironmentProfileResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentProfileResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfileResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentProfileResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfileResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentProfileResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type HostingEnvironmentProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingEnvironmentProfileResponse)(nil)).Elem()
}

func (o HostingEnvironmentProfileResponsePtrOutput) ToHostingEnvironmentProfileResponsePtrOutput() HostingEnvironmentProfileResponsePtrOutput {
	return o
}

func (o HostingEnvironmentProfileResponsePtrOutput) ToHostingEnvironmentProfileResponsePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfileResponsePtrOutput {
	return o
}

func (o HostingEnvironmentProfileResponsePtrOutput) Elem() HostingEnvironmentProfileResponseOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) HostingEnvironmentProfileResponse {
		if v != nil {
			return *v
		}
		var ret HostingEnvironmentProfileResponse
		return ret
	}).(HostingEnvironmentProfileResponseOutput)
}

func (o HostingEnvironmentProfileResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfileResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfileResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type HttpLogsConfig struct {
	AzureBlobStorage *AzureBlobStorageHttpLogsConfig `pulumi:"azureBlobStorage"`
	FileSystem       *FileSystemHttpLogsConfig       `pulumi:"fileSystem"`
}





type HttpLogsConfigInput interface {
	pulumi.Input

	ToHttpLogsConfigOutput() HttpLogsConfigOutput
	ToHttpLogsConfigOutputWithContext(context.Context) HttpLogsConfigOutput
}

type HttpLogsConfigArgs struct {
	AzureBlobStorage AzureBlobStorageHttpLogsConfigPtrInput `pulumi:"azureBlobStorage"`
	FileSystem       FileSystemHttpLogsConfigPtrInput       `pulumi:"fileSystem"`
}

func (HttpLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpLogsConfig)(nil)).Elem()
}

func (i HttpLogsConfigArgs) ToHttpLogsConfigOutput() HttpLogsConfigOutput {
	return i.ToHttpLogsConfigOutputWithContext(context.Background())
}

func (i HttpLogsConfigArgs) ToHttpLogsConfigOutputWithContext(ctx context.Context) HttpLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpLogsConfigOutput)
}

func (i HttpLogsConfigArgs) ToHttpLogsConfigPtrOutput() HttpLogsConfigPtrOutput {
	return i.ToHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (i HttpLogsConfigArgs) ToHttpLogsConfigPtrOutputWithContext(ctx context.Context) HttpLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpLogsConfigOutput).ToHttpLogsConfigPtrOutputWithContext(ctx)
}









type HttpLogsConfigPtrInput interface {
	pulumi.Input

	ToHttpLogsConfigPtrOutput() HttpLogsConfigPtrOutput
	ToHttpLogsConfigPtrOutputWithContext(context.Context) HttpLogsConfigPtrOutput
}

type httpLogsConfigPtrType HttpLogsConfigArgs

func HttpLogsConfigPtr(v *HttpLogsConfigArgs) HttpLogsConfigPtrInput {
	return (*httpLogsConfigPtrType)(v)
}

func (*httpLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpLogsConfig)(nil)).Elem()
}

func (i *httpLogsConfigPtrType) ToHttpLogsConfigPtrOutput() HttpLogsConfigPtrOutput {
	return i.ToHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (i *httpLogsConfigPtrType) ToHttpLogsConfigPtrOutputWithContext(ctx context.Context) HttpLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpLogsConfigPtrOutput)
}

type HttpLogsConfigOutput struct{ *pulumi.OutputState }

func (HttpLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpLogsConfig)(nil)).Elem()
}

func (o HttpLogsConfigOutput) ToHttpLogsConfigOutput() HttpLogsConfigOutput {
	return o
}

func (o HttpLogsConfigOutput) ToHttpLogsConfigOutputWithContext(ctx context.Context) HttpLogsConfigOutput {
	return o
}

func (o HttpLogsConfigOutput) ToHttpLogsConfigPtrOutput() HttpLogsConfigPtrOutput {
	return o.ToHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (o HttpLogsConfigOutput) ToHttpLogsConfigPtrOutputWithContext(ctx context.Context) HttpLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpLogsConfig) *HttpLogsConfig {
		return &v
	}).(HttpLogsConfigPtrOutput)
}

func (o HttpLogsConfigOutput) AzureBlobStorage() AzureBlobStorageHttpLogsConfigPtrOutput {
	return o.ApplyT(func(v HttpLogsConfig) *AzureBlobStorageHttpLogsConfig { return v.AzureBlobStorage }).(AzureBlobStorageHttpLogsConfigPtrOutput)
}

func (o HttpLogsConfigOutput) FileSystem() FileSystemHttpLogsConfigPtrOutput {
	return o.ApplyT(func(v HttpLogsConfig) *FileSystemHttpLogsConfig { return v.FileSystem }).(FileSystemHttpLogsConfigPtrOutput)
}

type HttpLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (HttpLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpLogsConfig)(nil)).Elem()
}

func (o HttpLogsConfigPtrOutput) ToHttpLogsConfigPtrOutput() HttpLogsConfigPtrOutput {
	return o
}

func (o HttpLogsConfigPtrOutput) ToHttpLogsConfigPtrOutputWithContext(ctx context.Context) HttpLogsConfigPtrOutput {
	return o
}

func (o HttpLogsConfigPtrOutput) Elem() HttpLogsConfigOutput {
	return o.ApplyT(func(v *HttpLogsConfig) HttpLogsConfig {
		if v != nil {
			return *v
		}
		var ret HttpLogsConfig
		return ret
	}).(HttpLogsConfigOutput)
}

func (o HttpLogsConfigPtrOutput) AzureBlobStorage() AzureBlobStorageHttpLogsConfigPtrOutput {
	return o.ApplyT(func(v *HttpLogsConfig) *AzureBlobStorageHttpLogsConfig {
		if v == nil {
			return nil
		}
		return v.AzureBlobStorage
	}).(AzureBlobStorageHttpLogsConfigPtrOutput)
}

func (o HttpLogsConfigPtrOutput) FileSystem() FileSystemHttpLogsConfigPtrOutput {
	return o.ApplyT(func(v *HttpLogsConfig) *FileSystemHttpLogsConfig {
		if v == nil {
			return nil
		}
		return v.FileSystem
	}).(FileSystemHttpLogsConfigPtrOutput)
}

type HttpLogsConfigResponse struct {
	AzureBlobStorage *AzureBlobStorageHttpLogsConfigResponse `pulumi:"azureBlobStorage"`
	FileSystem       *FileSystemHttpLogsConfigResponse       `pulumi:"fileSystem"`
}

type HttpLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (HttpLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpLogsConfigResponse)(nil)).Elem()
}

func (o HttpLogsConfigResponseOutput) ToHttpLogsConfigResponseOutput() HttpLogsConfigResponseOutput {
	return o
}

func (o HttpLogsConfigResponseOutput) ToHttpLogsConfigResponseOutputWithContext(ctx context.Context) HttpLogsConfigResponseOutput {
	return o
}

func (o HttpLogsConfigResponseOutput) AzureBlobStorage() AzureBlobStorageHttpLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v HttpLogsConfigResponse) *AzureBlobStorageHttpLogsConfigResponse { return v.AzureBlobStorage }).(AzureBlobStorageHttpLogsConfigResponsePtrOutput)
}

func (o HttpLogsConfigResponseOutput) FileSystem() FileSystemHttpLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v HttpLogsConfigResponse) *FileSystemHttpLogsConfigResponse { return v.FileSystem }).(FileSystemHttpLogsConfigResponsePtrOutput)
}

type HttpLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (HttpLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpLogsConfigResponse)(nil)).Elem()
}

func (o HttpLogsConfigResponsePtrOutput) ToHttpLogsConfigResponsePtrOutput() HttpLogsConfigResponsePtrOutput {
	return o
}

func (o HttpLogsConfigResponsePtrOutput) ToHttpLogsConfigResponsePtrOutputWithContext(ctx context.Context) HttpLogsConfigResponsePtrOutput {
	return o
}

func (o HttpLogsConfigResponsePtrOutput) Elem() HttpLogsConfigResponseOutput {
	return o.ApplyT(func(v *HttpLogsConfigResponse) HttpLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret HttpLogsConfigResponse
		return ret
	}).(HttpLogsConfigResponseOutput)
}

func (o HttpLogsConfigResponsePtrOutput) AzureBlobStorage() AzureBlobStorageHttpLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v *HttpLogsConfigResponse) *AzureBlobStorageHttpLogsConfigResponse {
		if v == nil {
			return nil
		}
		return v.AzureBlobStorage
	}).(AzureBlobStorageHttpLogsConfigResponsePtrOutput)
}

func (o HttpLogsConfigResponsePtrOutput) FileSystem() FileSystemHttpLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v *HttpLogsConfigResponse) *FileSystemHttpLogsConfigResponse {
		if v == nil {
			return nil
		}
		return v.FileSystem
	}).(FileSystemHttpLogsConfigResponsePtrOutput)
}

type IpSecurityRestriction struct {
	IpAddress  *string `pulumi:"ipAddress"`
	SubnetMask *string `pulumi:"subnetMask"`
}





type IpSecurityRestrictionInput interface {
	pulumi.Input

	ToIpSecurityRestrictionOutput() IpSecurityRestrictionOutput
	ToIpSecurityRestrictionOutputWithContext(context.Context) IpSecurityRestrictionOutput
}

type IpSecurityRestrictionArgs struct {
	IpAddress  pulumi.StringPtrInput `pulumi:"ipAddress"`
	SubnetMask pulumi.StringPtrInput `pulumi:"subnetMask"`
}

func (IpSecurityRestrictionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpSecurityRestriction)(nil)).Elem()
}

func (i IpSecurityRestrictionArgs) ToIpSecurityRestrictionOutput() IpSecurityRestrictionOutput {
	return i.ToIpSecurityRestrictionOutputWithContext(context.Background())
}

func (i IpSecurityRestrictionArgs) ToIpSecurityRestrictionOutputWithContext(ctx context.Context) IpSecurityRestrictionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSecurityRestrictionOutput)
}





type IpSecurityRestrictionArrayInput interface {
	pulumi.Input

	ToIpSecurityRestrictionArrayOutput() IpSecurityRestrictionArrayOutput
	ToIpSecurityRestrictionArrayOutputWithContext(context.Context) IpSecurityRestrictionArrayOutput
}

type IpSecurityRestrictionArray []IpSecurityRestrictionInput

func (IpSecurityRestrictionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpSecurityRestriction)(nil)).Elem()
}

func (i IpSecurityRestrictionArray) ToIpSecurityRestrictionArrayOutput() IpSecurityRestrictionArrayOutput {
	return i.ToIpSecurityRestrictionArrayOutputWithContext(context.Background())
}

func (i IpSecurityRestrictionArray) ToIpSecurityRestrictionArrayOutputWithContext(ctx context.Context) IpSecurityRestrictionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSecurityRestrictionArrayOutput)
}

type IpSecurityRestrictionOutput struct{ *pulumi.OutputState }

func (IpSecurityRestrictionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpSecurityRestriction)(nil)).Elem()
}

func (o IpSecurityRestrictionOutput) ToIpSecurityRestrictionOutput() IpSecurityRestrictionOutput {
	return o
}

func (o IpSecurityRestrictionOutput) ToIpSecurityRestrictionOutputWithContext(ctx context.Context) IpSecurityRestrictionOutput {
	return o
}

func (o IpSecurityRestrictionOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionOutput) SubnetMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *string { return v.SubnetMask }).(pulumi.StringPtrOutput)
}

type IpSecurityRestrictionArrayOutput struct{ *pulumi.OutputState }

func (IpSecurityRestrictionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpSecurityRestriction)(nil)).Elem()
}

func (o IpSecurityRestrictionArrayOutput) ToIpSecurityRestrictionArrayOutput() IpSecurityRestrictionArrayOutput {
	return o
}

func (o IpSecurityRestrictionArrayOutput) ToIpSecurityRestrictionArrayOutputWithContext(ctx context.Context) IpSecurityRestrictionArrayOutput {
	return o
}

func (o IpSecurityRestrictionArrayOutput) Index(i pulumi.IntInput) IpSecurityRestrictionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpSecurityRestriction {
		return vs[0].([]IpSecurityRestriction)[vs[1].(int)]
	}).(IpSecurityRestrictionOutput)
}

type IpSecurityRestrictionResponse struct {
	IpAddress  *string `pulumi:"ipAddress"`
	SubnetMask *string `pulumi:"subnetMask"`
}

type IpSecurityRestrictionResponseOutput struct{ *pulumi.OutputState }

func (IpSecurityRestrictionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpSecurityRestrictionResponse)(nil)).Elem()
}

func (o IpSecurityRestrictionResponseOutput) ToIpSecurityRestrictionResponseOutput() IpSecurityRestrictionResponseOutput {
	return o
}

func (o IpSecurityRestrictionResponseOutput) ToIpSecurityRestrictionResponseOutputWithContext(ctx context.Context) IpSecurityRestrictionResponseOutput {
	return o
}

func (o IpSecurityRestrictionResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) SubnetMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *string { return v.SubnetMask }).(pulumi.StringPtrOutput)
}

type IpSecurityRestrictionResponseArrayOutput struct{ *pulumi.OutputState }

func (IpSecurityRestrictionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpSecurityRestrictionResponse)(nil)).Elem()
}

func (o IpSecurityRestrictionResponseArrayOutput) ToIpSecurityRestrictionResponseArrayOutput() IpSecurityRestrictionResponseArrayOutput {
	return o
}

func (o IpSecurityRestrictionResponseArrayOutput) ToIpSecurityRestrictionResponseArrayOutputWithContext(ctx context.Context) IpSecurityRestrictionResponseArrayOutput {
	return o
}

func (o IpSecurityRestrictionResponseArrayOutput) Index(i pulumi.IntInput) IpSecurityRestrictionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpSecurityRestrictionResponse {
		return vs[0].([]IpSecurityRestrictionResponse)[vs[1].(int)]
	}).(IpSecurityRestrictionResponseOutput)
}

type NameValuePair struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type NameValuePairInput interface {
	pulumi.Input

	ToNameValuePairOutput() NameValuePairOutput
	ToNameValuePairOutputWithContext(context.Context) NameValuePairOutput
}

type NameValuePairArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (NameValuePairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NameValuePair)(nil)).Elem()
}

func (i NameValuePairArgs) ToNameValuePairOutput() NameValuePairOutput {
	return i.ToNameValuePairOutputWithContext(context.Background())
}

func (i NameValuePairArgs) ToNameValuePairOutputWithContext(ctx context.Context) NameValuePairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NameValuePairOutput)
}





type NameValuePairArrayInput interface {
	pulumi.Input

	ToNameValuePairArrayOutput() NameValuePairArrayOutput
	ToNameValuePairArrayOutputWithContext(context.Context) NameValuePairArrayOutput
}

type NameValuePairArray []NameValuePairInput

func (NameValuePairArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NameValuePair)(nil)).Elem()
}

func (i NameValuePairArray) ToNameValuePairArrayOutput() NameValuePairArrayOutput {
	return i.ToNameValuePairArrayOutputWithContext(context.Background())
}

func (i NameValuePairArray) ToNameValuePairArrayOutputWithContext(ctx context.Context) NameValuePairArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NameValuePairArrayOutput)
}

type NameValuePairOutput struct{ *pulumi.OutputState }

func (NameValuePairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NameValuePair)(nil)).Elem()
}

func (o NameValuePairOutput) ToNameValuePairOutput() NameValuePairOutput {
	return o
}

func (o NameValuePairOutput) ToNameValuePairOutputWithContext(ctx context.Context) NameValuePairOutput {
	return o
}

func (o NameValuePairOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NameValuePair) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NameValuePairOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NameValuePair) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type NameValuePairArrayOutput struct{ *pulumi.OutputState }

func (NameValuePairArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NameValuePair)(nil)).Elem()
}

func (o NameValuePairArrayOutput) ToNameValuePairArrayOutput() NameValuePairArrayOutput {
	return o
}

func (o NameValuePairArrayOutput) ToNameValuePairArrayOutputWithContext(ctx context.Context) NameValuePairArrayOutput {
	return o
}

func (o NameValuePairArrayOutput) Index(i pulumi.IntInput) NameValuePairOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NameValuePair {
		return vs[0].([]NameValuePair)[vs[1].(int)]
	}).(NameValuePairOutput)
}

type NameValuePairResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type NameValuePairResponseOutput struct{ *pulumi.OutputState }

func (NameValuePairResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NameValuePairResponse)(nil)).Elem()
}

func (o NameValuePairResponseOutput) ToNameValuePairResponseOutput() NameValuePairResponseOutput {
	return o
}

func (o NameValuePairResponseOutput) ToNameValuePairResponseOutputWithContext(ctx context.Context) NameValuePairResponseOutput {
	return o
}

func (o NameValuePairResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NameValuePairResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NameValuePairResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NameValuePairResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type NameValuePairResponseArrayOutput struct{ *pulumi.OutputState }

func (NameValuePairResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NameValuePairResponse)(nil)).Elem()
}

func (o NameValuePairResponseArrayOutput) ToNameValuePairResponseArrayOutput() NameValuePairResponseArrayOutput {
	return o
}

func (o NameValuePairResponseArrayOutput) ToNameValuePairResponseArrayOutputWithContext(ctx context.Context) NameValuePairResponseArrayOutput {
	return o
}

func (o NameValuePairResponseArrayOutput) Index(i pulumi.IntInput) NameValuePairResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NameValuePairResponse {
		return vs[0].([]NameValuePairResponse)[vs[1].(int)]
	}).(NameValuePairResponseOutput)
}

type NetworkAccessControlEntry struct {
	Action       *AccessControlEntryAction `pulumi:"action"`
	Description  *string                   `pulumi:"description"`
	Order        *int                      `pulumi:"order"`
	RemoteSubnet *string                   `pulumi:"remoteSubnet"`
}





type NetworkAccessControlEntryInput interface {
	pulumi.Input

	ToNetworkAccessControlEntryOutput() NetworkAccessControlEntryOutput
	ToNetworkAccessControlEntryOutputWithContext(context.Context) NetworkAccessControlEntryOutput
}

type NetworkAccessControlEntryArgs struct {
	Action       AccessControlEntryActionPtrInput `pulumi:"action"`
	Description  pulumi.StringPtrInput            `pulumi:"description"`
	Order        pulumi.IntPtrInput               `pulumi:"order"`
	RemoteSubnet pulumi.StringPtrInput            `pulumi:"remoteSubnet"`
}

func (NetworkAccessControlEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAccessControlEntry)(nil)).Elem()
}

func (i NetworkAccessControlEntryArgs) ToNetworkAccessControlEntryOutput() NetworkAccessControlEntryOutput {
	return i.ToNetworkAccessControlEntryOutputWithContext(context.Background())
}

func (i NetworkAccessControlEntryArgs) ToNetworkAccessControlEntryOutputWithContext(ctx context.Context) NetworkAccessControlEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAccessControlEntryOutput)
}





type NetworkAccessControlEntryArrayInput interface {
	pulumi.Input

	ToNetworkAccessControlEntryArrayOutput() NetworkAccessControlEntryArrayOutput
	ToNetworkAccessControlEntryArrayOutputWithContext(context.Context) NetworkAccessControlEntryArrayOutput
}

type NetworkAccessControlEntryArray []NetworkAccessControlEntryInput

func (NetworkAccessControlEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkAccessControlEntry)(nil)).Elem()
}

func (i NetworkAccessControlEntryArray) ToNetworkAccessControlEntryArrayOutput() NetworkAccessControlEntryArrayOutput {
	return i.ToNetworkAccessControlEntryArrayOutputWithContext(context.Background())
}

func (i NetworkAccessControlEntryArray) ToNetworkAccessControlEntryArrayOutputWithContext(ctx context.Context) NetworkAccessControlEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAccessControlEntryArrayOutput)
}

type NetworkAccessControlEntryOutput struct{ *pulumi.OutputState }

func (NetworkAccessControlEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAccessControlEntry)(nil)).Elem()
}

func (o NetworkAccessControlEntryOutput) ToNetworkAccessControlEntryOutput() NetworkAccessControlEntryOutput {
	return o
}

func (o NetworkAccessControlEntryOutput) ToNetworkAccessControlEntryOutputWithContext(ctx context.Context) NetworkAccessControlEntryOutput {
	return o
}

func (o NetworkAccessControlEntryOutput) Action() AccessControlEntryActionPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntry) *AccessControlEntryAction { return v.Action }).(AccessControlEntryActionPtrOutput)
}

func (o NetworkAccessControlEntryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntry) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkAccessControlEntryOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntry) *int { return v.Order }).(pulumi.IntPtrOutput)
}

func (o NetworkAccessControlEntryOutput) RemoteSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntry) *string { return v.RemoteSubnet }).(pulumi.StringPtrOutput)
}

type NetworkAccessControlEntryArrayOutput struct{ *pulumi.OutputState }

func (NetworkAccessControlEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkAccessControlEntry)(nil)).Elem()
}

func (o NetworkAccessControlEntryArrayOutput) ToNetworkAccessControlEntryArrayOutput() NetworkAccessControlEntryArrayOutput {
	return o
}

func (o NetworkAccessControlEntryArrayOutput) ToNetworkAccessControlEntryArrayOutputWithContext(ctx context.Context) NetworkAccessControlEntryArrayOutput {
	return o
}

func (o NetworkAccessControlEntryArrayOutput) Index(i pulumi.IntInput) NetworkAccessControlEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkAccessControlEntry {
		return vs[0].([]NetworkAccessControlEntry)[vs[1].(int)]
	}).(NetworkAccessControlEntryOutput)
}

type NetworkAccessControlEntryResponse struct {
	Action       *string `pulumi:"action"`
	Description  *string `pulumi:"description"`
	Order        *int    `pulumi:"order"`
	RemoteSubnet *string `pulumi:"remoteSubnet"`
}

type NetworkAccessControlEntryResponseOutput struct{ *pulumi.OutputState }

func (NetworkAccessControlEntryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAccessControlEntryResponse)(nil)).Elem()
}

func (o NetworkAccessControlEntryResponseOutput) ToNetworkAccessControlEntryResponseOutput() NetworkAccessControlEntryResponseOutput {
	return o
}

func (o NetworkAccessControlEntryResponseOutput) ToNetworkAccessControlEntryResponseOutputWithContext(ctx context.Context) NetworkAccessControlEntryResponseOutput {
	return o
}

func (o NetworkAccessControlEntryResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntryResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o NetworkAccessControlEntryResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntryResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkAccessControlEntryResponseOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntryResponse) *int { return v.Order }).(pulumi.IntPtrOutput)
}

func (o NetworkAccessControlEntryResponseOutput) RemoteSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntryResponse) *string { return v.RemoteSubnet }).(pulumi.StringPtrOutput)
}

type NetworkAccessControlEntryResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkAccessControlEntryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkAccessControlEntryResponse)(nil)).Elem()
}

func (o NetworkAccessControlEntryResponseArrayOutput) ToNetworkAccessControlEntryResponseArrayOutput() NetworkAccessControlEntryResponseArrayOutput {
	return o
}

func (o NetworkAccessControlEntryResponseArrayOutput) ToNetworkAccessControlEntryResponseArrayOutputWithContext(ctx context.Context) NetworkAccessControlEntryResponseArrayOutput {
	return o
}

func (o NetworkAccessControlEntryResponseArrayOutput) Index(i pulumi.IntInput) NetworkAccessControlEntryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkAccessControlEntryResponse {
		return vs[0].([]NetworkAccessControlEntryResponse)[vs[1].(int)]
	}).(NetworkAccessControlEntryResponseOutput)
}

type RampUpRule struct {
	ActionHostName            *string  `pulumi:"actionHostName"`
	ChangeDecisionCallbackUrl *string  `pulumi:"changeDecisionCallbackUrl"`
	ChangeIntervalInMinutes   *int     `pulumi:"changeIntervalInMinutes"`
	ChangeStep                *float64 `pulumi:"changeStep"`
	MaxReroutePercentage      *float64 `pulumi:"maxReroutePercentage"`
	MinReroutePercentage      *float64 `pulumi:"minReroutePercentage"`
	Name                      *string  `pulumi:"name"`
	ReroutePercentage         *float64 `pulumi:"reroutePercentage"`
}





type RampUpRuleInput interface {
	pulumi.Input

	ToRampUpRuleOutput() RampUpRuleOutput
	ToRampUpRuleOutputWithContext(context.Context) RampUpRuleOutput
}

type RampUpRuleArgs struct {
	ActionHostName            pulumi.StringPtrInput  `pulumi:"actionHostName"`
	ChangeDecisionCallbackUrl pulumi.StringPtrInput  `pulumi:"changeDecisionCallbackUrl"`
	ChangeIntervalInMinutes   pulumi.IntPtrInput     `pulumi:"changeIntervalInMinutes"`
	ChangeStep                pulumi.Float64PtrInput `pulumi:"changeStep"`
	MaxReroutePercentage      pulumi.Float64PtrInput `pulumi:"maxReroutePercentage"`
	MinReroutePercentage      pulumi.Float64PtrInput `pulumi:"minReroutePercentage"`
	Name                      pulumi.StringPtrInput  `pulumi:"name"`
	ReroutePercentage         pulumi.Float64PtrInput `pulumi:"reroutePercentage"`
}

func (RampUpRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RampUpRule)(nil)).Elem()
}

func (i RampUpRuleArgs) ToRampUpRuleOutput() RampUpRuleOutput {
	return i.ToRampUpRuleOutputWithContext(context.Background())
}

func (i RampUpRuleArgs) ToRampUpRuleOutputWithContext(ctx context.Context) RampUpRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RampUpRuleOutput)
}





type RampUpRuleArrayInput interface {
	pulumi.Input

	ToRampUpRuleArrayOutput() RampUpRuleArrayOutput
	ToRampUpRuleArrayOutputWithContext(context.Context) RampUpRuleArrayOutput
}

type RampUpRuleArray []RampUpRuleInput

func (RampUpRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RampUpRule)(nil)).Elem()
}

func (i RampUpRuleArray) ToRampUpRuleArrayOutput() RampUpRuleArrayOutput {
	return i.ToRampUpRuleArrayOutputWithContext(context.Background())
}

func (i RampUpRuleArray) ToRampUpRuleArrayOutputWithContext(ctx context.Context) RampUpRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RampUpRuleArrayOutput)
}

type RampUpRuleOutput struct{ *pulumi.OutputState }

func (RampUpRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RampUpRule)(nil)).Elem()
}

func (o RampUpRuleOutput) ToRampUpRuleOutput() RampUpRuleOutput {
	return o
}

func (o RampUpRuleOutput) ToRampUpRuleOutputWithContext(ctx context.Context) RampUpRuleOutput {
	return o
}

func (o RampUpRuleOutput) ActionHostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RampUpRule) *string { return v.ActionHostName }).(pulumi.StringPtrOutput)
}

func (o RampUpRuleOutput) ChangeDecisionCallbackUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RampUpRule) *string { return v.ChangeDecisionCallbackUrl }).(pulumi.StringPtrOutput)
}

func (o RampUpRuleOutput) ChangeIntervalInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RampUpRule) *int { return v.ChangeIntervalInMinutes }).(pulumi.IntPtrOutput)
}

func (o RampUpRuleOutput) ChangeStep() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRule) *float64 { return v.ChangeStep }).(pulumi.Float64PtrOutput)
}

func (o RampUpRuleOutput) MaxReroutePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRule) *float64 { return v.MaxReroutePercentage }).(pulumi.Float64PtrOutput)
}

func (o RampUpRuleOutput) MinReroutePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRule) *float64 { return v.MinReroutePercentage }).(pulumi.Float64PtrOutput)
}

func (o RampUpRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RampUpRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RampUpRuleOutput) ReroutePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRule) *float64 { return v.ReroutePercentage }).(pulumi.Float64PtrOutput)
}

type RampUpRuleArrayOutput struct{ *pulumi.OutputState }

func (RampUpRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RampUpRule)(nil)).Elem()
}

func (o RampUpRuleArrayOutput) ToRampUpRuleArrayOutput() RampUpRuleArrayOutput {
	return o
}

func (o RampUpRuleArrayOutput) ToRampUpRuleArrayOutputWithContext(ctx context.Context) RampUpRuleArrayOutput {
	return o
}

func (o RampUpRuleArrayOutput) Index(i pulumi.IntInput) RampUpRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RampUpRule {
		return vs[0].([]RampUpRule)[vs[1].(int)]
	}).(RampUpRuleOutput)
}

type RampUpRuleResponse struct {
	ActionHostName            *string  `pulumi:"actionHostName"`
	ChangeDecisionCallbackUrl *string  `pulumi:"changeDecisionCallbackUrl"`
	ChangeIntervalInMinutes   *int     `pulumi:"changeIntervalInMinutes"`
	ChangeStep                *float64 `pulumi:"changeStep"`
	MaxReroutePercentage      *float64 `pulumi:"maxReroutePercentage"`
	MinReroutePercentage      *float64 `pulumi:"minReroutePercentage"`
	Name                      *string  `pulumi:"name"`
	ReroutePercentage         *float64 `pulumi:"reroutePercentage"`
}

type RampUpRuleResponseOutput struct{ *pulumi.OutputState }

func (RampUpRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RampUpRuleResponse)(nil)).Elem()
}

func (o RampUpRuleResponseOutput) ToRampUpRuleResponseOutput() RampUpRuleResponseOutput {
	return o
}

func (o RampUpRuleResponseOutput) ToRampUpRuleResponseOutputWithContext(ctx context.Context) RampUpRuleResponseOutput {
	return o
}

func (o RampUpRuleResponseOutput) ActionHostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *string { return v.ActionHostName }).(pulumi.StringPtrOutput)
}

func (o RampUpRuleResponseOutput) ChangeDecisionCallbackUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *string { return v.ChangeDecisionCallbackUrl }).(pulumi.StringPtrOutput)
}

func (o RampUpRuleResponseOutput) ChangeIntervalInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *int { return v.ChangeIntervalInMinutes }).(pulumi.IntPtrOutput)
}

func (o RampUpRuleResponseOutput) ChangeStep() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *float64 { return v.ChangeStep }).(pulumi.Float64PtrOutput)
}

func (o RampUpRuleResponseOutput) MaxReroutePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *float64 { return v.MaxReroutePercentage }).(pulumi.Float64PtrOutput)
}

func (o RampUpRuleResponseOutput) MinReroutePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *float64 { return v.MinReroutePercentage }).(pulumi.Float64PtrOutput)
}

func (o RampUpRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RampUpRuleResponseOutput) ReroutePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *float64 { return v.ReroutePercentage }).(pulumi.Float64PtrOutput)
}

type RampUpRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (RampUpRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RampUpRuleResponse)(nil)).Elem()
}

func (o RampUpRuleResponseArrayOutput) ToRampUpRuleResponseArrayOutput() RampUpRuleResponseArrayOutput {
	return o
}

func (o RampUpRuleResponseArrayOutput) ToRampUpRuleResponseArrayOutputWithContext(ctx context.Context) RampUpRuleResponseArrayOutput {
	return o
}

func (o RampUpRuleResponseArrayOutput) Index(i pulumi.IntInput) RampUpRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RampUpRuleResponse {
		return vs[0].([]RampUpRuleResponse)[vs[1].(int)]
	}).(RampUpRuleResponseOutput)
}

type RequestsBasedTrigger struct {
	Count        *int    `pulumi:"count"`
	TimeInterval *string `pulumi:"timeInterval"`
}





type RequestsBasedTriggerInput interface {
	pulumi.Input

	ToRequestsBasedTriggerOutput() RequestsBasedTriggerOutput
	ToRequestsBasedTriggerOutputWithContext(context.Context) RequestsBasedTriggerOutput
}

type RequestsBasedTriggerArgs struct {
	Count        pulumi.IntPtrInput    `pulumi:"count"`
	TimeInterval pulumi.StringPtrInput `pulumi:"timeInterval"`
}

func (RequestsBasedTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestsBasedTrigger)(nil)).Elem()
}

func (i RequestsBasedTriggerArgs) ToRequestsBasedTriggerOutput() RequestsBasedTriggerOutput {
	return i.ToRequestsBasedTriggerOutputWithContext(context.Background())
}

func (i RequestsBasedTriggerArgs) ToRequestsBasedTriggerOutputWithContext(ctx context.Context) RequestsBasedTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestsBasedTriggerOutput)
}

func (i RequestsBasedTriggerArgs) ToRequestsBasedTriggerPtrOutput() RequestsBasedTriggerPtrOutput {
	return i.ToRequestsBasedTriggerPtrOutputWithContext(context.Background())
}

func (i RequestsBasedTriggerArgs) ToRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) RequestsBasedTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestsBasedTriggerOutput).ToRequestsBasedTriggerPtrOutputWithContext(ctx)
}









type RequestsBasedTriggerPtrInput interface {
	pulumi.Input

	ToRequestsBasedTriggerPtrOutput() RequestsBasedTriggerPtrOutput
	ToRequestsBasedTriggerPtrOutputWithContext(context.Context) RequestsBasedTriggerPtrOutput
}

type requestsBasedTriggerPtrType RequestsBasedTriggerArgs

func RequestsBasedTriggerPtr(v *RequestsBasedTriggerArgs) RequestsBasedTriggerPtrInput {
	return (*requestsBasedTriggerPtrType)(v)
}

func (*requestsBasedTriggerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestsBasedTrigger)(nil)).Elem()
}

func (i *requestsBasedTriggerPtrType) ToRequestsBasedTriggerPtrOutput() RequestsBasedTriggerPtrOutput {
	return i.ToRequestsBasedTriggerPtrOutputWithContext(context.Background())
}

func (i *requestsBasedTriggerPtrType) ToRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) RequestsBasedTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestsBasedTriggerPtrOutput)
}

type RequestsBasedTriggerOutput struct{ *pulumi.OutputState }

func (RequestsBasedTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestsBasedTrigger)(nil)).Elem()
}

func (o RequestsBasedTriggerOutput) ToRequestsBasedTriggerOutput() RequestsBasedTriggerOutput {
	return o
}

func (o RequestsBasedTriggerOutput) ToRequestsBasedTriggerOutputWithContext(ctx context.Context) RequestsBasedTriggerOutput {
	return o
}

func (o RequestsBasedTriggerOutput) ToRequestsBasedTriggerPtrOutput() RequestsBasedTriggerPtrOutput {
	return o.ToRequestsBasedTriggerPtrOutputWithContext(context.Background())
}

func (o RequestsBasedTriggerOutput) ToRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) RequestsBasedTriggerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RequestsBasedTrigger) *RequestsBasedTrigger {
		return &v
	}).(RequestsBasedTriggerPtrOutput)
}

func (o RequestsBasedTriggerOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RequestsBasedTrigger) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o RequestsBasedTriggerOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RequestsBasedTrigger) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

type RequestsBasedTriggerPtrOutput struct{ *pulumi.OutputState }

func (RequestsBasedTriggerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestsBasedTrigger)(nil)).Elem()
}

func (o RequestsBasedTriggerPtrOutput) ToRequestsBasedTriggerPtrOutput() RequestsBasedTriggerPtrOutput {
	return o
}

func (o RequestsBasedTriggerPtrOutput) ToRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) RequestsBasedTriggerPtrOutput {
	return o
}

func (o RequestsBasedTriggerPtrOutput) Elem() RequestsBasedTriggerOutput {
	return o.ApplyT(func(v *RequestsBasedTrigger) RequestsBasedTrigger {
		if v != nil {
			return *v
		}
		var ret RequestsBasedTrigger
		return ret
	}).(RequestsBasedTriggerOutput)
}

func (o RequestsBasedTriggerPtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RequestsBasedTrigger) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o RequestsBasedTriggerPtrOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestsBasedTrigger) *string {
		if v == nil {
			return nil
		}
		return v.TimeInterval
	}).(pulumi.StringPtrOutput)
}

type RequestsBasedTriggerResponse struct {
	Count        *int    `pulumi:"count"`
	TimeInterval *string `pulumi:"timeInterval"`
}

type RequestsBasedTriggerResponseOutput struct{ *pulumi.OutputState }

func (RequestsBasedTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestsBasedTriggerResponse)(nil)).Elem()
}

func (o RequestsBasedTriggerResponseOutput) ToRequestsBasedTriggerResponseOutput() RequestsBasedTriggerResponseOutput {
	return o
}

func (o RequestsBasedTriggerResponseOutput) ToRequestsBasedTriggerResponseOutputWithContext(ctx context.Context) RequestsBasedTriggerResponseOutput {
	return o
}

func (o RequestsBasedTriggerResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RequestsBasedTriggerResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o RequestsBasedTriggerResponseOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RequestsBasedTriggerResponse) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

type RequestsBasedTriggerResponsePtrOutput struct{ *pulumi.OutputState }

func (RequestsBasedTriggerResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestsBasedTriggerResponse)(nil)).Elem()
}

func (o RequestsBasedTriggerResponsePtrOutput) ToRequestsBasedTriggerResponsePtrOutput() RequestsBasedTriggerResponsePtrOutput {
	return o
}

func (o RequestsBasedTriggerResponsePtrOutput) ToRequestsBasedTriggerResponsePtrOutputWithContext(ctx context.Context) RequestsBasedTriggerResponsePtrOutput {
	return o
}

func (o RequestsBasedTriggerResponsePtrOutput) Elem() RequestsBasedTriggerResponseOutput {
	return o.ApplyT(func(v *RequestsBasedTriggerResponse) RequestsBasedTriggerResponse {
		if v != nil {
			return *v
		}
		var ret RequestsBasedTriggerResponse
		return ret
	}).(RequestsBasedTriggerResponseOutput)
}

func (o RequestsBasedTriggerResponsePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RequestsBasedTriggerResponse) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o RequestsBasedTriggerResponsePtrOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestsBasedTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeInterval
	}).(pulumi.StringPtrOutput)
}

type SiteConfig struct {
	AlwaysOn                     *bool                   `pulumi:"alwaysOn"`
	ApiDefinition                *ApiDefinitionInfo      `pulumi:"apiDefinition"`
	AppCommandLine               *string                 `pulumi:"appCommandLine"`
	AppSettings                  []NameValuePair         `pulumi:"appSettings"`
	AutoHealEnabled              *bool                   `pulumi:"autoHealEnabled"`
	AutoHealRules                *AutoHealRules          `pulumi:"autoHealRules"`
	AutoSwapSlotName             *string                 `pulumi:"autoSwapSlotName"`
	ConnectionStrings            []ConnStringInfo        `pulumi:"connectionStrings"`
	Cors                         *CorsSettings           `pulumi:"cors"`
	DefaultDocuments             []string                `pulumi:"defaultDocuments"`
	DetailedErrorLoggingEnabled  *bool                   `pulumi:"detailedErrorLoggingEnabled"`
	DocumentRoot                 *string                 `pulumi:"documentRoot"`
	Experiments                  *Experiments            `pulumi:"experiments"`
	HandlerMappings              []HandlerMapping        `pulumi:"handlerMappings"`
	HttpLoggingEnabled           *bool                   `pulumi:"httpLoggingEnabled"`
	Id                           *string                 `pulumi:"id"`
	IpSecurityRestrictions       []IpSecurityRestriction `pulumi:"ipSecurityRestrictions"`
	JavaContainer                *string                 `pulumi:"javaContainer"`
	JavaContainerVersion         *string                 `pulumi:"javaContainerVersion"`
	JavaVersion                  *string                 `pulumi:"javaVersion"`
	Kind                         *string                 `pulumi:"kind"`
	Limits                       *SiteLimits             `pulumi:"limits"`
	LoadBalancing                *SiteLoadBalancing      `pulumi:"loadBalancing"`
	LocalMySqlEnabled            *bool                   `pulumi:"localMySqlEnabled"`
	Location                     string                  `pulumi:"location"`
	LogsDirectorySizeLimit       *int                    `pulumi:"logsDirectorySizeLimit"`
	ManagedPipelineMode          *ManagedPipelineMode    `pulumi:"managedPipelineMode"`
	Metadata                     []NameValuePair         `pulumi:"metadata"`
	Name                         *string                 `pulumi:"name"`
	NetFrameworkVersion          *string                 `pulumi:"netFrameworkVersion"`
	NodeVersion                  *string                 `pulumi:"nodeVersion"`
	NumberOfWorkers              *int                    `pulumi:"numberOfWorkers"`
	PhpVersion                   *string                 `pulumi:"phpVersion"`
	PublishingPassword           *string                 `pulumi:"publishingPassword"`
	PublishingUsername           *string                 `pulumi:"publishingUsername"`
	PythonVersion                *string                 `pulumi:"pythonVersion"`
	RemoteDebuggingEnabled       *bool                   `pulumi:"remoteDebuggingEnabled"`
	RemoteDebuggingVersion       *string                 `pulumi:"remoteDebuggingVersion"`
	RequestTracingEnabled        *bool                   `pulumi:"requestTracingEnabled"`
	RequestTracingExpirationTime *string                 `pulumi:"requestTracingExpirationTime"`
	ScmType                      *string                 `pulumi:"scmType"`
	Tags                         map[string]string       `pulumi:"tags"`
	TracingOptions               *string                 `pulumi:"tracingOptions"`
	Type                         *string                 `pulumi:"type"`
	Use32BitWorkerProcess        *bool                   `pulumi:"use32BitWorkerProcess"`
	VirtualApplications          []VirtualApplication    `pulumi:"virtualApplications"`
	VnetName                     *string                 `pulumi:"vnetName"`
	WebSocketsEnabled            *bool                   `pulumi:"webSocketsEnabled"`
}





type SiteConfigInput interface {
	pulumi.Input

	ToSiteConfigOutput() SiteConfigOutput
	ToSiteConfigOutputWithContext(context.Context) SiteConfigOutput
}

type SiteConfigArgs struct {
	AlwaysOn                     pulumi.BoolPtrInput             `pulumi:"alwaysOn"`
	ApiDefinition                ApiDefinitionInfoPtrInput       `pulumi:"apiDefinition"`
	AppCommandLine               pulumi.StringPtrInput           `pulumi:"appCommandLine"`
	AppSettings                  NameValuePairArrayInput         `pulumi:"appSettings"`
	AutoHealEnabled              pulumi.BoolPtrInput             `pulumi:"autoHealEnabled"`
	AutoHealRules                AutoHealRulesPtrInput           `pulumi:"autoHealRules"`
	AutoSwapSlotName             pulumi.StringPtrInput           `pulumi:"autoSwapSlotName"`
	ConnectionStrings            ConnStringInfoArrayInput        `pulumi:"connectionStrings"`
	Cors                         CorsSettingsPtrInput            `pulumi:"cors"`
	DefaultDocuments             pulumi.StringArrayInput         `pulumi:"defaultDocuments"`
	DetailedErrorLoggingEnabled  pulumi.BoolPtrInput             `pulumi:"detailedErrorLoggingEnabled"`
	DocumentRoot                 pulumi.StringPtrInput           `pulumi:"documentRoot"`
	Experiments                  ExperimentsPtrInput             `pulumi:"experiments"`
	HandlerMappings              HandlerMappingArrayInput        `pulumi:"handlerMappings"`
	HttpLoggingEnabled           pulumi.BoolPtrInput             `pulumi:"httpLoggingEnabled"`
	Id                           pulumi.StringPtrInput           `pulumi:"id"`
	IpSecurityRestrictions       IpSecurityRestrictionArrayInput `pulumi:"ipSecurityRestrictions"`
	JavaContainer                pulumi.StringPtrInput           `pulumi:"javaContainer"`
	JavaContainerVersion         pulumi.StringPtrInput           `pulumi:"javaContainerVersion"`
	JavaVersion                  pulumi.StringPtrInput           `pulumi:"javaVersion"`
	Kind                         pulumi.StringPtrInput           `pulumi:"kind"`
	Limits                       SiteLimitsPtrInput              `pulumi:"limits"`
	LoadBalancing                SiteLoadBalancingPtrInput       `pulumi:"loadBalancing"`
	LocalMySqlEnabled            pulumi.BoolPtrInput             `pulumi:"localMySqlEnabled"`
	Location                     pulumi.StringInput              `pulumi:"location"`
	LogsDirectorySizeLimit       pulumi.IntPtrInput              `pulumi:"logsDirectorySizeLimit"`
	ManagedPipelineMode          ManagedPipelineModePtrInput     `pulumi:"managedPipelineMode"`
	Metadata                     NameValuePairArrayInput         `pulumi:"metadata"`
	Name                         pulumi.StringPtrInput           `pulumi:"name"`
	NetFrameworkVersion          pulumi.StringPtrInput           `pulumi:"netFrameworkVersion"`
	NodeVersion                  pulumi.StringPtrInput           `pulumi:"nodeVersion"`
	NumberOfWorkers              pulumi.IntPtrInput              `pulumi:"numberOfWorkers"`
	PhpVersion                   pulumi.StringPtrInput           `pulumi:"phpVersion"`
	PublishingPassword           pulumi.StringPtrInput           `pulumi:"publishingPassword"`
	PublishingUsername           pulumi.StringPtrInput           `pulumi:"publishingUsername"`
	PythonVersion                pulumi.StringPtrInput           `pulumi:"pythonVersion"`
	RemoteDebuggingEnabled       pulumi.BoolPtrInput             `pulumi:"remoteDebuggingEnabled"`
	RemoteDebuggingVersion       pulumi.StringPtrInput           `pulumi:"remoteDebuggingVersion"`
	RequestTracingEnabled        pulumi.BoolPtrInput             `pulumi:"requestTracingEnabled"`
	RequestTracingExpirationTime pulumi.StringPtrInput           `pulumi:"requestTracingExpirationTime"`
	ScmType                      pulumi.StringPtrInput           `pulumi:"scmType"`
	Tags                         pulumi.StringMapInput           `pulumi:"tags"`
	TracingOptions               pulumi.StringPtrInput           `pulumi:"tracingOptions"`
	Type                         pulumi.StringPtrInput           `pulumi:"type"`
	Use32BitWorkerProcess        pulumi.BoolPtrInput             `pulumi:"use32BitWorkerProcess"`
	VirtualApplications          VirtualApplicationArrayInput    `pulumi:"virtualApplications"`
	VnetName                     pulumi.StringPtrInput           `pulumi:"vnetName"`
	WebSocketsEnabled            pulumi.BoolPtrInput             `pulumi:"webSocketsEnabled"`
}

func (SiteConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteConfig)(nil)).Elem()
}

func (i SiteConfigArgs) ToSiteConfigOutput() SiteConfigOutput {
	return i.ToSiteConfigOutputWithContext(context.Background())
}

func (i SiteConfigArgs) ToSiteConfigOutputWithContext(ctx context.Context) SiteConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConfigOutput)
}

func (i SiteConfigArgs) ToSiteConfigPtrOutput() SiteConfigPtrOutput {
	return i.ToSiteConfigPtrOutputWithContext(context.Background())
}

func (i SiteConfigArgs) ToSiteConfigPtrOutputWithContext(ctx context.Context) SiteConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConfigOutput).ToSiteConfigPtrOutputWithContext(ctx)
}









type SiteConfigPtrInput interface {
	pulumi.Input

	ToSiteConfigPtrOutput() SiteConfigPtrOutput
	ToSiteConfigPtrOutputWithContext(context.Context) SiteConfigPtrOutput
}

type siteConfigPtrType SiteConfigArgs

func SiteConfigPtr(v *SiteConfigArgs) SiteConfigPtrInput {
	return (*siteConfigPtrType)(v)
}

func (*siteConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteConfig)(nil)).Elem()
}

func (i *siteConfigPtrType) ToSiteConfigPtrOutput() SiteConfigPtrOutput {
	return i.ToSiteConfigPtrOutputWithContext(context.Background())
}

func (i *siteConfigPtrType) ToSiteConfigPtrOutputWithContext(ctx context.Context) SiteConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConfigPtrOutput)
}

type SiteConfigOutput struct{ *pulumi.OutputState }

func (SiteConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteConfig)(nil)).Elem()
}

func (o SiteConfigOutput) ToSiteConfigOutput() SiteConfigOutput {
	return o
}

func (o SiteConfigOutput) ToSiteConfigOutputWithContext(ctx context.Context) SiteConfigOutput {
	return o
}

func (o SiteConfigOutput) ToSiteConfigPtrOutput() SiteConfigPtrOutput {
	return o.ToSiteConfigPtrOutputWithContext(context.Background())
}

func (o SiteConfigOutput) ToSiteConfigPtrOutputWithContext(ctx context.Context) SiteConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteConfig) *SiteConfig {
		return &v
	}).(SiteConfigPtrOutput)
}

func (o SiteConfigOutput) AlwaysOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.AlwaysOn }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) ApiDefinition() ApiDefinitionInfoPtrOutput {
	return o.ApplyT(func(v SiteConfig) *ApiDefinitionInfo { return v.ApiDefinition }).(ApiDefinitionInfoPtrOutput)
}

func (o SiteConfigOutput) AppCommandLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.AppCommandLine }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) AppSettings() NameValuePairArrayOutput {
	return o.ApplyT(func(v SiteConfig) []NameValuePair { return v.AppSettings }).(NameValuePairArrayOutput)
}

func (o SiteConfigOutput) AutoHealEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.AutoHealEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) AutoHealRules() AutoHealRulesPtrOutput {
	return o.ApplyT(func(v SiteConfig) *AutoHealRules { return v.AutoHealRules }).(AutoHealRulesPtrOutput)
}

func (o SiteConfigOutput) AutoSwapSlotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.AutoSwapSlotName }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) ConnectionStrings() ConnStringInfoArrayOutput {
	return o.ApplyT(func(v SiteConfig) []ConnStringInfo { return v.ConnectionStrings }).(ConnStringInfoArrayOutput)
}

func (o SiteConfigOutput) Cors() CorsSettingsPtrOutput {
	return o.ApplyT(func(v SiteConfig) *CorsSettings { return v.Cors }).(CorsSettingsPtrOutput)
}

func (o SiteConfigOutput) DefaultDocuments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SiteConfig) []string { return v.DefaultDocuments }).(pulumi.StringArrayOutput)
}

func (o SiteConfigOutput) DetailedErrorLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.DetailedErrorLoggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) DocumentRoot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.DocumentRoot }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) Experiments() ExperimentsPtrOutput {
	return o.ApplyT(func(v SiteConfig) *Experiments { return v.Experiments }).(ExperimentsPtrOutput)
}

func (o SiteConfigOutput) HandlerMappings() HandlerMappingArrayOutput {
	return o.ApplyT(func(v SiteConfig) []HandlerMapping { return v.HandlerMappings }).(HandlerMappingArrayOutput)
}

func (o SiteConfigOutput) HttpLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.HttpLoggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) IpSecurityRestrictions() IpSecurityRestrictionArrayOutput {
	return o.ApplyT(func(v SiteConfig) []IpSecurityRestriction { return v.IpSecurityRestrictions }).(IpSecurityRestrictionArrayOutput)
}

func (o SiteConfigOutput) JavaContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.JavaContainer }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) JavaContainerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.JavaContainerVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) JavaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.JavaVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) Limits() SiteLimitsPtrOutput {
	return o.ApplyT(func(v SiteConfig) *SiteLimits { return v.Limits }).(SiteLimitsPtrOutput)
}

func (o SiteConfigOutput) LoadBalancing() SiteLoadBalancingPtrOutput {
	return o.ApplyT(func(v SiteConfig) *SiteLoadBalancing { return v.LoadBalancing }).(SiteLoadBalancingPtrOutput)
}

func (o SiteConfigOutput) LocalMySqlEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.LocalMySqlEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v SiteConfig) string { return v.Location }).(pulumi.StringOutput)
}

func (o SiteConfigOutput) LogsDirectorySizeLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfig) *int { return v.LogsDirectorySizeLimit }).(pulumi.IntPtrOutput)
}

func (o SiteConfigOutput) ManagedPipelineMode() ManagedPipelineModePtrOutput {
	return o.ApplyT(func(v SiteConfig) *ManagedPipelineMode { return v.ManagedPipelineMode }).(ManagedPipelineModePtrOutput)
}

func (o SiteConfigOutput) Metadata() NameValuePairArrayOutput {
	return o.ApplyT(func(v SiteConfig) []NameValuePair { return v.Metadata }).(NameValuePairArrayOutput)
}

func (o SiteConfigOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) NetFrameworkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.NetFrameworkVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) NodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.NodeVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfig) *int { return v.NumberOfWorkers }).(pulumi.IntPtrOutput)
}

func (o SiteConfigOutput) PhpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.PhpVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) PublishingPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.PublishingPassword }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) PublishingUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.PublishingUsername }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) PythonVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.PythonVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) RemoteDebuggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.RemoteDebuggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) RemoteDebuggingVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.RemoteDebuggingVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) RequestTracingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.RequestTracingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) RequestTracingExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.RequestTracingExpirationTime }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) ScmType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.ScmType }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v SiteConfig) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SiteConfigOutput) TracingOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.TracingOptions }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) Use32BitWorkerProcess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.Use32BitWorkerProcess }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) VirtualApplications() VirtualApplicationArrayOutput {
	return o.ApplyT(func(v SiteConfig) []VirtualApplication { return v.VirtualApplications }).(VirtualApplicationArrayOutput)
}

func (o SiteConfigOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.VnetName }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) WebSocketsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.WebSocketsEnabled }).(pulumi.BoolPtrOutput)
}

type SiteConfigPtrOutput struct{ *pulumi.OutputState }

func (SiteConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteConfig)(nil)).Elem()
}

func (o SiteConfigPtrOutput) ToSiteConfigPtrOutput() SiteConfigPtrOutput {
	return o
}

func (o SiteConfigPtrOutput) ToSiteConfigPtrOutputWithContext(ctx context.Context) SiteConfigPtrOutput {
	return o
}

func (o SiteConfigPtrOutput) Elem() SiteConfigOutput {
	return o.ApplyT(func(v *SiteConfig) SiteConfig {
		if v != nil {
			return *v
		}
		var ret SiteConfig
		return ret
	}).(SiteConfigOutput)
}

func (o SiteConfigPtrOutput) AlwaysOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.AlwaysOn
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) ApiDefinition() ApiDefinitionInfoPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *ApiDefinitionInfo {
		if v == nil {
			return nil
		}
		return v.ApiDefinition
	}).(ApiDefinitionInfoPtrOutput)
}

func (o SiteConfigPtrOutput) AppCommandLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.AppCommandLine
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) AppSettings() NameValuePairArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []NameValuePair {
		if v == nil {
			return nil
		}
		return v.AppSettings
	}).(NameValuePairArrayOutput)
}

func (o SiteConfigPtrOutput) AutoHealEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.AutoHealEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) AutoHealRules() AutoHealRulesPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *AutoHealRules {
		if v == nil {
			return nil
		}
		return v.AutoHealRules
	}).(AutoHealRulesPtrOutput)
}

func (o SiteConfigPtrOutput) AutoSwapSlotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.AutoSwapSlotName
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) ConnectionStrings() ConnStringInfoArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []ConnStringInfo {
		if v == nil {
			return nil
		}
		return v.ConnectionStrings
	}).(ConnStringInfoArrayOutput)
}

func (o SiteConfigPtrOutput) Cors() CorsSettingsPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *CorsSettings {
		if v == nil {
			return nil
		}
		return v.Cors
	}).(CorsSettingsPtrOutput)
}

func (o SiteConfigPtrOutput) DefaultDocuments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []string {
		if v == nil {
			return nil
		}
		return v.DefaultDocuments
	}).(pulumi.StringArrayOutput)
}

func (o SiteConfigPtrOutput) DetailedErrorLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.DetailedErrorLoggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) DocumentRoot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.DocumentRoot
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) Experiments() ExperimentsPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *Experiments {
		if v == nil {
			return nil
		}
		return v.Experiments
	}).(ExperimentsPtrOutput)
}

func (o SiteConfigPtrOutput) HandlerMappings() HandlerMappingArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []HandlerMapping {
		if v == nil {
			return nil
		}
		return v.HandlerMappings
	}).(HandlerMappingArrayOutput)
}

func (o SiteConfigPtrOutput) HttpLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.HttpLoggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) IpSecurityRestrictions() IpSecurityRestrictionArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []IpSecurityRestriction {
		if v == nil {
			return nil
		}
		return v.IpSecurityRestrictions
	}).(IpSecurityRestrictionArrayOutput)
}

func (o SiteConfigPtrOutput) JavaContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.JavaContainer
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) JavaContainerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.JavaContainerVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) JavaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.JavaVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) Limits() SiteLimitsPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *SiteLimits {
		if v == nil {
			return nil
		}
		return v.Limits
	}).(SiteLimitsPtrOutput)
}

func (o SiteConfigPtrOutput) LoadBalancing() SiteLoadBalancingPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *SiteLoadBalancing {
		if v == nil {
			return nil
		}
		return v.LoadBalancing
	}).(SiteLoadBalancingPtrOutput)
}

func (o SiteConfigPtrOutput) LocalMySqlEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.LocalMySqlEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) LogsDirectorySizeLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *int {
		if v == nil {
			return nil
		}
		return v.LogsDirectorySizeLimit
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigPtrOutput) ManagedPipelineMode() ManagedPipelineModePtrOutput {
	return o.ApplyT(func(v *SiteConfig) *ManagedPipelineMode {
		if v == nil {
			return nil
		}
		return v.ManagedPipelineMode
	}).(ManagedPipelineModePtrOutput)
}

func (o SiteConfigPtrOutput) Metadata() NameValuePairArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []NameValuePair {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(NameValuePairArrayOutput)
}

func (o SiteConfigPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) NetFrameworkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.NetFrameworkVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) NodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.NodeVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *int {
		if v == nil {
			return nil
		}
		return v.NumberOfWorkers
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigPtrOutput) PhpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.PhpVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) PublishingPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.PublishingPassword
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) PublishingUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.PublishingUsername
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) PythonVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.PythonVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) RemoteDebuggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.RemoteDebuggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) RemoteDebuggingVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.RemoteDebuggingVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) RequestTracingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.RequestTracingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) RequestTracingExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.RequestTracingExpirationTime
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) ScmType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.ScmType
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SiteConfig) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o SiteConfigPtrOutput) TracingOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.TracingOptions
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) Use32BitWorkerProcess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.Use32BitWorkerProcess
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) VirtualApplications() VirtualApplicationArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []VirtualApplication {
		if v == nil {
			return nil
		}
		return v.VirtualApplications
	}).(VirtualApplicationArrayOutput)
}

func (o SiteConfigPtrOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.VnetName
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) WebSocketsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.WebSocketsEnabled
	}).(pulumi.BoolPtrOutput)
}

type SiteConfigResponse struct {
	AlwaysOn                     *bool                           `pulumi:"alwaysOn"`
	ApiDefinition                *ApiDefinitionInfoResponse      `pulumi:"apiDefinition"`
	AppCommandLine               *string                         `pulumi:"appCommandLine"`
	AppSettings                  []NameValuePairResponse         `pulumi:"appSettings"`
	AutoHealEnabled              *bool                           `pulumi:"autoHealEnabled"`
	AutoHealRules                *AutoHealRulesResponse          `pulumi:"autoHealRules"`
	AutoSwapSlotName             *string                         `pulumi:"autoSwapSlotName"`
	ConnectionStrings            []ConnStringInfoResponse        `pulumi:"connectionStrings"`
	Cors                         *CorsSettingsResponse           `pulumi:"cors"`
	DefaultDocuments             []string                        `pulumi:"defaultDocuments"`
	DetailedErrorLoggingEnabled  *bool                           `pulumi:"detailedErrorLoggingEnabled"`
	DocumentRoot                 *string                         `pulumi:"documentRoot"`
	Experiments                  *ExperimentsResponse            `pulumi:"experiments"`
	HandlerMappings              []HandlerMappingResponse        `pulumi:"handlerMappings"`
	HttpLoggingEnabled           *bool                           `pulumi:"httpLoggingEnabled"`
	Id                           *string                         `pulumi:"id"`
	IpSecurityRestrictions       []IpSecurityRestrictionResponse `pulumi:"ipSecurityRestrictions"`
	JavaContainer                *string                         `pulumi:"javaContainer"`
	JavaContainerVersion         *string                         `pulumi:"javaContainerVersion"`
	JavaVersion                  *string                         `pulumi:"javaVersion"`
	Kind                         *string                         `pulumi:"kind"`
	Limits                       *SiteLimitsResponse             `pulumi:"limits"`
	LoadBalancing                *string                         `pulumi:"loadBalancing"`
	LocalMySqlEnabled            *bool                           `pulumi:"localMySqlEnabled"`
	Location                     string                          `pulumi:"location"`
	LogsDirectorySizeLimit       *int                            `pulumi:"logsDirectorySizeLimit"`
	ManagedPipelineMode          *string                         `pulumi:"managedPipelineMode"`
	Metadata                     []NameValuePairResponse         `pulumi:"metadata"`
	Name                         *string                         `pulumi:"name"`
	NetFrameworkVersion          *string                         `pulumi:"netFrameworkVersion"`
	NodeVersion                  *string                         `pulumi:"nodeVersion"`
	NumberOfWorkers              *int                            `pulumi:"numberOfWorkers"`
	PhpVersion                   *string                         `pulumi:"phpVersion"`
	PublishingPassword           *string                         `pulumi:"publishingPassword"`
	PublishingUsername           *string                         `pulumi:"publishingUsername"`
	PythonVersion                *string                         `pulumi:"pythonVersion"`
	RemoteDebuggingEnabled       *bool                           `pulumi:"remoteDebuggingEnabled"`
	RemoteDebuggingVersion       *string                         `pulumi:"remoteDebuggingVersion"`
	RequestTracingEnabled        *bool                           `pulumi:"requestTracingEnabled"`
	RequestTracingExpirationTime *string                         `pulumi:"requestTracingExpirationTime"`
	ScmType                      *string                         `pulumi:"scmType"`
	Tags                         map[string]string               `pulumi:"tags"`
	TracingOptions               *string                         `pulumi:"tracingOptions"`
	Type                         *string                         `pulumi:"type"`
	Use32BitWorkerProcess        *bool                           `pulumi:"use32BitWorkerProcess"`
	VirtualApplications          []VirtualApplicationResponse    `pulumi:"virtualApplications"`
	VnetName                     *string                         `pulumi:"vnetName"`
	WebSocketsEnabled            *bool                           `pulumi:"webSocketsEnabled"`
}

type SiteConfigResponseOutput struct{ *pulumi.OutputState }

func (SiteConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteConfigResponse)(nil)).Elem()
}

func (o SiteConfigResponseOutput) ToSiteConfigResponseOutput() SiteConfigResponseOutput {
	return o
}

func (o SiteConfigResponseOutput) ToSiteConfigResponseOutputWithContext(ctx context.Context) SiteConfigResponseOutput {
	return o
}

func (o SiteConfigResponseOutput) AlwaysOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.AlwaysOn }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) ApiDefinition() ApiDefinitionInfoResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *ApiDefinitionInfoResponse { return v.ApiDefinition }).(ApiDefinitionInfoResponsePtrOutput)
}

func (o SiteConfigResponseOutput) AppCommandLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.AppCommandLine }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) AppSettings() NameValuePairResponseArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []NameValuePairResponse { return v.AppSettings }).(NameValuePairResponseArrayOutput)
}

func (o SiteConfigResponseOutput) AutoHealEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.AutoHealEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) AutoHealRules() AutoHealRulesResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *AutoHealRulesResponse { return v.AutoHealRules }).(AutoHealRulesResponsePtrOutput)
}

func (o SiteConfigResponseOutput) AutoSwapSlotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.AutoSwapSlotName }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) ConnectionStrings() ConnStringInfoResponseArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []ConnStringInfoResponse { return v.ConnectionStrings }).(ConnStringInfoResponseArrayOutput)
}

func (o SiteConfigResponseOutput) Cors() CorsSettingsResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *CorsSettingsResponse { return v.Cors }).(CorsSettingsResponsePtrOutput)
}

func (o SiteConfigResponseOutput) DefaultDocuments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []string { return v.DefaultDocuments }).(pulumi.StringArrayOutput)
}

func (o SiteConfigResponseOutput) DetailedErrorLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.DetailedErrorLoggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) DocumentRoot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.DocumentRoot }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) Experiments() ExperimentsResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *ExperimentsResponse { return v.Experiments }).(ExperimentsResponsePtrOutput)
}

func (o SiteConfigResponseOutput) HandlerMappings() HandlerMappingResponseArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []HandlerMappingResponse { return v.HandlerMappings }).(HandlerMappingResponseArrayOutput)
}

func (o SiteConfigResponseOutput) HttpLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.HttpLoggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) IpSecurityRestrictions() IpSecurityRestrictionResponseArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []IpSecurityRestrictionResponse { return v.IpSecurityRestrictions }).(IpSecurityRestrictionResponseArrayOutput)
}

func (o SiteConfigResponseOutput) JavaContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.JavaContainer }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) JavaContainerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.JavaContainerVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) JavaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.JavaVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) Limits() SiteLimitsResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *SiteLimitsResponse { return v.Limits }).(SiteLimitsResponsePtrOutput)
}

func (o SiteConfigResponseOutput) LoadBalancing() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.LoadBalancing }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) LocalMySqlEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.LocalMySqlEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v SiteConfigResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o SiteConfigResponseOutput) LogsDirectorySizeLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *int { return v.LogsDirectorySizeLimit }).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponseOutput) ManagedPipelineMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.ManagedPipelineMode }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) Metadata() NameValuePairResponseArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []NameValuePairResponse { return v.Metadata }).(NameValuePairResponseArrayOutput)
}

func (o SiteConfigResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) NetFrameworkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.NetFrameworkVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) NodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.NodeVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *int { return v.NumberOfWorkers }).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponseOutput) PhpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.PhpVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) PublishingPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.PublishingPassword }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) PublishingUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.PublishingUsername }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) PythonVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.PythonVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) RemoteDebuggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.RemoteDebuggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) RemoteDebuggingVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.RemoteDebuggingVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) RequestTracingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.RequestTracingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) RequestTracingExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.RequestTracingExpirationTime }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) ScmType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.ScmType }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v SiteConfigResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SiteConfigResponseOutput) TracingOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.TracingOptions }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) Use32BitWorkerProcess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.Use32BitWorkerProcess }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) VirtualApplications() VirtualApplicationResponseArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []VirtualApplicationResponse { return v.VirtualApplications }).(VirtualApplicationResponseArrayOutput)
}

func (o SiteConfigResponseOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.VnetName }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) WebSocketsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.WebSocketsEnabled }).(pulumi.BoolPtrOutput)
}

type SiteConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (SiteConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteConfigResponse)(nil)).Elem()
}

func (o SiteConfigResponsePtrOutput) ToSiteConfigResponsePtrOutput() SiteConfigResponsePtrOutput {
	return o
}

func (o SiteConfigResponsePtrOutput) ToSiteConfigResponsePtrOutputWithContext(ctx context.Context) SiteConfigResponsePtrOutput {
	return o
}

func (o SiteConfigResponsePtrOutput) Elem() SiteConfigResponseOutput {
	return o.ApplyT(func(v *SiteConfigResponse) SiteConfigResponse {
		if v != nil {
			return *v
		}
		var ret SiteConfigResponse
		return ret
	}).(SiteConfigResponseOutput)
}

func (o SiteConfigResponsePtrOutput) AlwaysOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AlwaysOn
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) ApiDefinition() ApiDefinitionInfoResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *ApiDefinitionInfoResponse {
		if v == nil {
			return nil
		}
		return v.ApiDefinition
	}).(ApiDefinitionInfoResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) AppCommandLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppCommandLine
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) AppSettings() NameValuePairResponseArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []NameValuePairResponse {
		if v == nil {
			return nil
		}
		return v.AppSettings
	}).(NameValuePairResponseArrayOutput)
}

func (o SiteConfigResponsePtrOutput) AutoHealEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AutoHealEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) AutoHealRules() AutoHealRulesResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *AutoHealRulesResponse {
		if v == nil {
			return nil
		}
		return v.AutoHealRules
	}).(AutoHealRulesResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) AutoSwapSlotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.AutoSwapSlotName
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) ConnectionStrings() ConnStringInfoResponseArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []ConnStringInfoResponse {
		if v == nil {
			return nil
		}
		return v.ConnectionStrings
	}).(ConnStringInfoResponseArrayOutput)
}

func (o SiteConfigResponsePtrOutput) Cors() CorsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *CorsSettingsResponse {
		if v == nil {
			return nil
		}
		return v.Cors
	}).(CorsSettingsResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) DefaultDocuments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []string {
		if v == nil {
			return nil
		}
		return v.DefaultDocuments
	}).(pulumi.StringArrayOutput)
}

func (o SiteConfigResponsePtrOutput) DetailedErrorLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DetailedErrorLoggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) DocumentRoot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.DocumentRoot
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Experiments() ExperimentsResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *ExperimentsResponse {
		if v == nil {
			return nil
		}
		return v.Experiments
	}).(ExperimentsResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) HandlerMappings() HandlerMappingResponseArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []HandlerMappingResponse {
		if v == nil {
			return nil
		}
		return v.HandlerMappings
	}).(HandlerMappingResponseArrayOutput)
}

func (o SiteConfigResponsePtrOutput) HttpLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.HttpLoggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) IpSecurityRestrictions() IpSecurityRestrictionResponseArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []IpSecurityRestrictionResponse {
		if v == nil {
			return nil
		}
		return v.IpSecurityRestrictions
	}).(IpSecurityRestrictionResponseArrayOutput)
}

func (o SiteConfigResponsePtrOutput) JavaContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.JavaContainer
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) JavaContainerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.JavaContainerVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) JavaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.JavaVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Limits() SiteLimitsResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *SiteLimitsResponse {
		if v == nil {
			return nil
		}
		return v.Limits
	}).(SiteLimitsResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) LoadBalancing() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancing
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) LocalMySqlEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.LocalMySqlEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) LogsDirectorySizeLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.LogsDirectorySizeLimit
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponsePtrOutput) ManagedPipelineMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.ManagedPipelineMode
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Metadata() NameValuePairResponseArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []NameValuePairResponse {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(NameValuePairResponseArrayOutput)
}

func (o SiteConfigResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) NetFrameworkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.NetFrameworkVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) NodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.NodeVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NumberOfWorkers
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponsePtrOutput) PhpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.PhpVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) PublishingPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublishingPassword
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) PublishingUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublishingUsername
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) PythonVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.PythonVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) RemoteDebuggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RemoteDebuggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) RemoteDebuggingVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.RemoteDebuggingVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) RequestTracingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequestTracingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) RequestTracingExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.RequestTracingExpirationTime
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) ScmType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScmType
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SiteConfigResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o SiteConfigResponsePtrOutput) TracingOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.TracingOptions
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Use32BitWorkerProcess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Use32BitWorkerProcess
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) VirtualApplications() VirtualApplicationResponseArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []VirtualApplicationResponse {
		if v == nil {
			return nil
		}
		return v.VirtualApplications
	}).(VirtualApplicationResponseArrayOutput)
}

func (o SiteConfigResponsePtrOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.VnetName
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) WebSocketsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.WebSocketsEnabled
	}).(pulumi.BoolPtrOutput)
}

type SiteLimits struct {
	MaxDiskSizeInMb  *float64 `pulumi:"maxDiskSizeInMb"`
	MaxMemoryInMb    *float64 `pulumi:"maxMemoryInMb"`
	MaxPercentageCpu *float64 `pulumi:"maxPercentageCpu"`
}





type SiteLimitsInput interface {
	pulumi.Input

	ToSiteLimitsOutput() SiteLimitsOutput
	ToSiteLimitsOutputWithContext(context.Context) SiteLimitsOutput
}

type SiteLimitsArgs struct {
	MaxDiskSizeInMb  pulumi.Float64PtrInput `pulumi:"maxDiskSizeInMb"`
	MaxMemoryInMb    pulumi.Float64PtrInput `pulumi:"maxMemoryInMb"`
	MaxPercentageCpu pulumi.Float64PtrInput `pulumi:"maxPercentageCpu"`
}

func (SiteLimitsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteLimits)(nil)).Elem()
}

func (i SiteLimitsArgs) ToSiteLimitsOutput() SiteLimitsOutput {
	return i.ToSiteLimitsOutputWithContext(context.Background())
}

func (i SiteLimitsArgs) ToSiteLimitsOutputWithContext(ctx context.Context) SiteLimitsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteLimitsOutput)
}

func (i SiteLimitsArgs) ToSiteLimitsPtrOutput() SiteLimitsPtrOutput {
	return i.ToSiteLimitsPtrOutputWithContext(context.Background())
}

func (i SiteLimitsArgs) ToSiteLimitsPtrOutputWithContext(ctx context.Context) SiteLimitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteLimitsOutput).ToSiteLimitsPtrOutputWithContext(ctx)
}









type SiteLimitsPtrInput interface {
	pulumi.Input

	ToSiteLimitsPtrOutput() SiteLimitsPtrOutput
	ToSiteLimitsPtrOutputWithContext(context.Context) SiteLimitsPtrOutput
}

type siteLimitsPtrType SiteLimitsArgs

func SiteLimitsPtr(v *SiteLimitsArgs) SiteLimitsPtrInput {
	return (*siteLimitsPtrType)(v)
}

func (*siteLimitsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteLimits)(nil)).Elem()
}

func (i *siteLimitsPtrType) ToSiteLimitsPtrOutput() SiteLimitsPtrOutput {
	return i.ToSiteLimitsPtrOutputWithContext(context.Background())
}

func (i *siteLimitsPtrType) ToSiteLimitsPtrOutputWithContext(ctx context.Context) SiteLimitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteLimitsPtrOutput)
}

type SiteLimitsOutput struct{ *pulumi.OutputState }

func (SiteLimitsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteLimits)(nil)).Elem()
}

func (o SiteLimitsOutput) ToSiteLimitsOutput() SiteLimitsOutput {
	return o
}

func (o SiteLimitsOutput) ToSiteLimitsOutputWithContext(ctx context.Context) SiteLimitsOutput {
	return o
}

func (o SiteLimitsOutput) ToSiteLimitsPtrOutput() SiteLimitsPtrOutput {
	return o.ToSiteLimitsPtrOutputWithContext(context.Background())
}

func (o SiteLimitsOutput) ToSiteLimitsPtrOutputWithContext(ctx context.Context) SiteLimitsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteLimits) *SiteLimits {
		return &v
	}).(SiteLimitsPtrOutput)
}

func (o SiteLimitsOutput) MaxDiskSizeInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SiteLimits) *float64 { return v.MaxDiskSizeInMb }).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsOutput) MaxMemoryInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SiteLimits) *float64 { return v.MaxMemoryInMb }).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsOutput) MaxPercentageCpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SiteLimits) *float64 { return v.MaxPercentageCpu }).(pulumi.Float64PtrOutput)
}

type SiteLimitsPtrOutput struct{ *pulumi.OutputState }

func (SiteLimitsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteLimits)(nil)).Elem()
}

func (o SiteLimitsPtrOutput) ToSiteLimitsPtrOutput() SiteLimitsPtrOutput {
	return o
}

func (o SiteLimitsPtrOutput) ToSiteLimitsPtrOutputWithContext(ctx context.Context) SiteLimitsPtrOutput {
	return o
}

func (o SiteLimitsPtrOutput) Elem() SiteLimitsOutput {
	return o.ApplyT(func(v *SiteLimits) SiteLimits {
		if v != nil {
			return *v
		}
		var ret SiteLimits
		return ret
	}).(SiteLimitsOutput)
}

func (o SiteLimitsPtrOutput) MaxDiskSizeInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteLimits) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxDiskSizeInMb
	}).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsPtrOutput) MaxMemoryInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteLimits) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxMemoryInMb
	}).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsPtrOutput) MaxPercentageCpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteLimits) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxPercentageCpu
	}).(pulumi.Float64PtrOutput)
}

type SiteLimitsResponse struct {
	MaxDiskSizeInMb  *float64 `pulumi:"maxDiskSizeInMb"`
	MaxMemoryInMb    *float64 `pulumi:"maxMemoryInMb"`
	MaxPercentageCpu *float64 `pulumi:"maxPercentageCpu"`
}

type SiteLimitsResponseOutput struct{ *pulumi.OutputState }

func (SiteLimitsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteLimitsResponse)(nil)).Elem()
}

func (o SiteLimitsResponseOutput) ToSiteLimitsResponseOutput() SiteLimitsResponseOutput {
	return o
}

func (o SiteLimitsResponseOutput) ToSiteLimitsResponseOutputWithContext(ctx context.Context) SiteLimitsResponseOutput {
	return o
}

func (o SiteLimitsResponseOutput) MaxDiskSizeInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SiteLimitsResponse) *float64 { return v.MaxDiskSizeInMb }).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsResponseOutput) MaxMemoryInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SiteLimitsResponse) *float64 { return v.MaxMemoryInMb }).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsResponseOutput) MaxPercentageCpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SiteLimitsResponse) *float64 { return v.MaxPercentageCpu }).(pulumi.Float64PtrOutput)
}

type SiteLimitsResponsePtrOutput struct{ *pulumi.OutputState }

func (SiteLimitsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteLimitsResponse)(nil)).Elem()
}

func (o SiteLimitsResponsePtrOutput) ToSiteLimitsResponsePtrOutput() SiteLimitsResponsePtrOutput {
	return o
}

func (o SiteLimitsResponsePtrOutput) ToSiteLimitsResponsePtrOutputWithContext(ctx context.Context) SiteLimitsResponsePtrOutput {
	return o
}

func (o SiteLimitsResponsePtrOutput) Elem() SiteLimitsResponseOutput {
	return o.ApplyT(func(v *SiteLimitsResponse) SiteLimitsResponse {
		if v != nil {
			return *v
		}
		var ret SiteLimitsResponse
		return ret
	}).(SiteLimitsResponseOutput)
}

func (o SiteLimitsResponsePtrOutput) MaxDiskSizeInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteLimitsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxDiskSizeInMb
	}).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsResponsePtrOutput) MaxMemoryInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteLimitsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxMemoryInMb
	}).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsResponsePtrOutput) MaxPercentageCpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteLimitsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxPercentageCpu
	}).(pulumi.Float64PtrOutput)
}

type SkuDescription struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     *string `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuDescriptionInput interface {
	pulumi.Input

	ToSkuDescriptionOutput() SkuDescriptionOutput
	ToSkuDescriptionOutputWithContext(context.Context) SkuDescriptionOutput
}

type SkuDescriptionArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDescription)(nil)).Elem()
}

func (i SkuDescriptionArgs) ToSkuDescriptionOutput() SkuDescriptionOutput {
	return i.ToSkuDescriptionOutputWithContext(context.Background())
}

func (i SkuDescriptionArgs) ToSkuDescriptionOutputWithContext(ctx context.Context) SkuDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionOutput)
}

func (i SkuDescriptionArgs) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return i.ToSkuDescriptionPtrOutputWithContext(context.Background())
}

func (i SkuDescriptionArgs) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionOutput).ToSkuDescriptionPtrOutputWithContext(ctx)
}









type SkuDescriptionPtrInput interface {
	pulumi.Input

	ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput
	ToSkuDescriptionPtrOutputWithContext(context.Context) SkuDescriptionPtrOutput
}

type skuDescriptionPtrType SkuDescriptionArgs

func SkuDescriptionPtr(v *SkuDescriptionArgs) SkuDescriptionPtrInput {
	return (*skuDescriptionPtrType)(v)
}

func (*skuDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuDescription)(nil)).Elem()
}

func (i *skuDescriptionPtrType) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return i.ToSkuDescriptionPtrOutputWithContext(context.Background())
}

func (i *skuDescriptionPtrType) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionPtrOutput)
}

type SkuDescriptionOutput struct{ *pulumi.OutputState }

func (SkuDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDescription)(nil)).Elem()
}

func (o SkuDescriptionOutput) ToSkuDescriptionOutput() SkuDescriptionOutput {
	return o
}

func (o SkuDescriptionOutput) ToSkuDescriptionOutputWithContext(ctx context.Context) SkuDescriptionOutput {
	return o
}

func (o SkuDescriptionOutput) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return o.ToSkuDescriptionPtrOutputWithContext(context.Background())
}

func (o SkuDescriptionOutput) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuDescription) *SkuDescription {
		return &v
	}).(SkuDescriptionPtrOutput)
}

func (o SkuDescriptionOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuDescription) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuDescriptionPtrOutput struct{ *pulumi.OutputState }

func (SkuDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuDescription)(nil)).Elem()
}

func (o SkuDescriptionPtrOutput) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return o
}

func (o SkuDescriptionPtrOutput) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return o
}

func (o SkuDescriptionPtrOutput) Elem() SkuDescriptionOutput {
	return o.ApplyT(func(v *SkuDescription) SkuDescription {
		if v != nil {
			return *v
		}
		var ret SkuDescription
		return ret
	}).(SkuDescriptionOutput)
}

func (o SkuDescriptionPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuDescriptionResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     *string `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}

type SkuDescriptionResponseOutput struct{ *pulumi.OutputState }

func (SkuDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDescriptionResponse)(nil)).Elem()
}

func (o SkuDescriptionResponseOutput) ToSkuDescriptionResponseOutput() SkuDescriptionResponseOutput {
	return o
}

func (o SkuDescriptionResponseOutput) ToSkuDescriptionResponseOutputWithContext(ctx context.Context) SkuDescriptionResponseOutput {
	return o
}

func (o SkuDescriptionResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuDescriptionResponse)(nil)).Elem()
}

func (o SkuDescriptionResponsePtrOutput) ToSkuDescriptionResponsePtrOutput() SkuDescriptionResponsePtrOutput {
	return o
}

func (o SkuDescriptionResponsePtrOutput) ToSkuDescriptionResponsePtrOutputWithContext(ctx context.Context) SkuDescriptionResponsePtrOutput {
	return o
}

func (o SkuDescriptionResponsePtrOutput) Elem() SkuDescriptionResponseOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) SkuDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret SkuDescriptionResponse
		return ret
	}).(SkuDescriptionResponseOutput)
}

func (o SkuDescriptionResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SlowRequestsBasedTrigger struct {
	Count        *int    `pulumi:"count"`
	TimeInterval *string `pulumi:"timeInterval"`
	TimeTaken    *string `pulumi:"timeTaken"`
}





type SlowRequestsBasedTriggerInput interface {
	pulumi.Input

	ToSlowRequestsBasedTriggerOutput() SlowRequestsBasedTriggerOutput
	ToSlowRequestsBasedTriggerOutputWithContext(context.Context) SlowRequestsBasedTriggerOutput
}

type SlowRequestsBasedTriggerArgs struct {
	Count        pulumi.IntPtrInput    `pulumi:"count"`
	TimeInterval pulumi.StringPtrInput `pulumi:"timeInterval"`
	TimeTaken    pulumi.StringPtrInput `pulumi:"timeTaken"`
}

func (SlowRequestsBasedTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SlowRequestsBasedTrigger)(nil)).Elem()
}

func (i SlowRequestsBasedTriggerArgs) ToSlowRequestsBasedTriggerOutput() SlowRequestsBasedTriggerOutput {
	return i.ToSlowRequestsBasedTriggerOutputWithContext(context.Background())
}

func (i SlowRequestsBasedTriggerArgs) ToSlowRequestsBasedTriggerOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlowRequestsBasedTriggerOutput)
}

func (i SlowRequestsBasedTriggerArgs) ToSlowRequestsBasedTriggerPtrOutput() SlowRequestsBasedTriggerPtrOutput {
	return i.ToSlowRequestsBasedTriggerPtrOutputWithContext(context.Background())
}

func (i SlowRequestsBasedTriggerArgs) ToSlowRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlowRequestsBasedTriggerOutput).ToSlowRequestsBasedTriggerPtrOutputWithContext(ctx)
}









type SlowRequestsBasedTriggerPtrInput interface {
	pulumi.Input

	ToSlowRequestsBasedTriggerPtrOutput() SlowRequestsBasedTriggerPtrOutput
	ToSlowRequestsBasedTriggerPtrOutputWithContext(context.Context) SlowRequestsBasedTriggerPtrOutput
}

type slowRequestsBasedTriggerPtrType SlowRequestsBasedTriggerArgs

func SlowRequestsBasedTriggerPtr(v *SlowRequestsBasedTriggerArgs) SlowRequestsBasedTriggerPtrInput {
	return (*slowRequestsBasedTriggerPtrType)(v)
}

func (*slowRequestsBasedTriggerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SlowRequestsBasedTrigger)(nil)).Elem()
}

func (i *slowRequestsBasedTriggerPtrType) ToSlowRequestsBasedTriggerPtrOutput() SlowRequestsBasedTriggerPtrOutput {
	return i.ToSlowRequestsBasedTriggerPtrOutputWithContext(context.Background())
}

func (i *slowRequestsBasedTriggerPtrType) ToSlowRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlowRequestsBasedTriggerPtrOutput)
}

type SlowRequestsBasedTriggerOutput struct{ *pulumi.OutputState }

func (SlowRequestsBasedTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SlowRequestsBasedTrigger)(nil)).Elem()
}

func (o SlowRequestsBasedTriggerOutput) ToSlowRequestsBasedTriggerOutput() SlowRequestsBasedTriggerOutput {
	return o
}

func (o SlowRequestsBasedTriggerOutput) ToSlowRequestsBasedTriggerOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerOutput {
	return o
}

func (o SlowRequestsBasedTriggerOutput) ToSlowRequestsBasedTriggerPtrOutput() SlowRequestsBasedTriggerPtrOutput {
	return o.ToSlowRequestsBasedTriggerPtrOutputWithContext(context.Background())
}

func (o SlowRequestsBasedTriggerOutput) ToSlowRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SlowRequestsBasedTrigger) *SlowRequestsBasedTrigger {
		return &v
	}).(SlowRequestsBasedTriggerPtrOutput)
}

func (o SlowRequestsBasedTriggerOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTrigger) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o SlowRequestsBasedTriggerOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTrigger) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerOutput) TimeTaken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTrigger) *string { return v.TimeTaken }).(pulumi.StringPtrOutput)
}

type SlowRequestsBasedTriggerPtrOutput struct{ *pulumi.OutputState }

func (SlowRequestsBasedTriggerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SlowRequestsBasedTrigger)(nil)).Elem()
}

func (o SlowRequestsBasedTriggerPtrOutput) ToSlowRequestsBasedTriggerPtrOutput() SlowRequestsBasedTriggerPtrOutput {
	return o
}

func (o SlowRequestsBasedTriggerPtrOutput) ToSlowRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerPtrOutput {
	return o
}

func (o SlowRequestsBasedTriggerPtrOutput) Elem() SlowRequestsBasedTriggerOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTrigger) SlowRequestsBasedTrigger {
		if v != nil {
			return *v
		}
		var ret SlowRequestsBasedTrigger
		return ret
	}).(SlowRequestsBasedTriggerOutput)
}

func (o SlowRequestsBasedTriggerPtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTrigger) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o SlowRequestsBasedTriggerPtrOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTrigger) *string {
		if v == nil {
			return nil
		}
		return v.TimeInterval
	}).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerPtrOutput) TimeTaken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTrigger) *string {
		if v == nil {
			return nil
		}
		return v.TimeTaken
	}).(pulumi.StringPtrOutput)
}

type SlowRequestsBasedTriggerResponse struct {
	Count        *int    `pulumi:"count"`
	TimeInterval *string `pulumi:"timeInterval"`
	TimeTaken    *string `pulumi:"timeTaken"`
}

type SlowRequestsBasedTriggerResponseOutput struct{ *pulumi.OutputState }

func (SlowRequestsBasedTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SlowRequestsBasedTriggerResponse)(nil)).Elem()
}

func (o SlowRequestsBasedTriggerResponseOutput) ToSlowRequestsBasedTriggerResponseOutput() SlowRequestsBasedTriggerResponseOutput {
	return o
}

func (o SlowRequestsBasedTriggerResponseOutput) ToSlowRequestsBasedTriggerResponseOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerResponseOutput {
	return o
}

func (o SlowRequestsBasedTriggerResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTriggerResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o SlowRequestsBasedTriggerResponseOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTriggerResponse) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerResponseOutput) TimeTaken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTriggerResponse) *string { return v.TimeTaken }).(pulumi.StringPtrOutput)
}

type SlowRequestsBasedTriggerResponsePtrOutput struct{ *pulumi.OutputState }

func (SlowRequestsBasedTriggerResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SlowRequestsBasedTriggerResponse)(nil)).Elem()
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) ToSlowRequestsBasedTriggerResponsePtrOutput() SlowRequestsBasedTriggerResponsePtrOutput {
	return o
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) ToSlowRequestsBasedTriggerResponsePtrOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerResponsePtrOutput {
	return o
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) Elem() SlowRequestsBasedTriggerResponseOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTriggerResponse) SlowRequestsBasedTriggerResponse {
		if v != nil {
			return *v
		}
		var ret SlowRequestsBasedTriggerResponse
		return ret
	}).(SlowRequestsBasedTriggerResponseOutput)
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTriggerResponse) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeInterval
	}).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) TimeTaken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeTaken
	}).(pulumi.StringPtrOutput)
}

type StampCapacity struct {
	AvailableCapacity              *float64            `pulumi:"availableCapacity"`
	ComputeMode                    *ComputeModeOptions `pulumi:"computeMode"`
	ExcludeFromCapacityAllocation  *bool               `pulumi:"excludeFromCapacityAllocation"`
	IsApplicableForAllComputeModes *bool               `pulumi:"isApplicableForAllComputeModes"`
	Name                           *string             `pulumi:"name"`
	SiteMode                       *string             `pulumi:"siteMode"`
	TotalCapacity                  *float64            `pulumi:"totalCapacity"`
	Unit                           *string             `pulumi:"unit"`
	WorkerSize                     *WorkerSizeOptions  `pulumi:"workerSize"`
	WorkerSizeId                   *int                `pulumi:"workerSizeId"`
}





type StampCapacityInput interface {
	pulumi.Input

	ToStampCapacityOutput() StampCapacityOutput
	ToStampCapacityOutputWithContext(context.Context) StampCapacityOutput
}

type StampCapacityArgs struct {
	AvailableCapacity              pulumi.Float64PtrInput     `pulumi:"availableCapacity"`
	ComputeMode                    ComputeModeOptionsPtrInput `pulumi:"computeMode"`
	ExcludeFromCapacityAllocation  pulumi.BoolPtrInput        `pulumi:"excludeFromCapacityAllocation"`
	IsApplicableForAllComputeModes pulumi.BoolPtrInput        `pulumi:"isApplicableForAllComputeModes"`
	Name                           pulumi.StringPtrInput      `pulumi:"name"`
	SiteMode                       pulumi.StringPtrInput      `pulumi:"siteMode"`
	TotalCapacity                  pulumi.Float64PtrInput     `pulumi:"totalCapacity"`
	Unit                           pulumi.StringPtrInput      `pulumi:"unit"`
	WorkerSize                     WorkerSizeOptionsPtrInput  `pulumi:"workerSize"`
	WorkerSizeId                   pulumi.IntPtrInput         `pulumi:"workerSizeId"`
}

func (StampCapacityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StampCapacity)(nil)).Elem()
}

func (i StampCapacityArgs) ToStampCapacityOutput() StampCapacityOutput {
	return i.ToStampCapacityOutputWithContext(context.Background())
}

func (i StampCapacityArgs) ToStampCapacityOutputWithContext(ctx context.Context) StampCapacityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StampCapacityOutput)
}





type StampCapacityArrayInput interface {
	pulumi.Input

	ToStampCapacityArrayOutput() StampCapacityArrayOutput
	ToStampCapacityArrayOutputWithContext(context.Context) StampCapacityArrayOutput
}

type StampCapacityArray []StampCapacityInput

func (StampCapacityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StampCapacity)(nil)).Elem()
}

func (i StampCapacityArray) ToStampCapacityArrayOutput() StampCapacityArrayOutput {
	return i.ToStampCapacityArrayOutputWithContext(context.Background())
}

func (i StampCapacityArray) ToStampCapacityArrayOutputWithContext(ctx context.Context) StampCapacityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StampCapacityArrayOutput)
}

type StampCapacityOutput struct{ *pulumi.OutputState }

func (StampCapacityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StampCapacity)(nil)).Elem()
}

func (o StampCapacityOutput) ToStampCapacityOutput() StampCapacityOutput {
	return o
}

func (o StampCapacityOutput) ToStampCapacityOutputWithContext(ctx context.Context) StampCapacityOutput {
	return o
}

func (o StampCapacityOutput) AvailableCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v StampCapacity) *float64 { return v.AvailableCapacity }).(pulumi.Float64PtrOutput)
}

func (o StampCapacityOutput) ComputeMode() ComputeModeOptionsPtrOutput {
	return o.ApplyT(func(v StampCapacity) *ComputeModeOptions { return v.ComputeMode }).(ComputeModeOptionsPtrOutput)
}

func (o StampCapacityOutput) ExcludeFromCapacityAllocation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StampCapacity) *bool { return v.ExcludeFromCapacityAllocation }).(pulumi.BoolPtrOutput)
}

func (o StampCapacityOutput) IsApplicableForAllComputeModes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StampCapacity) *bool { return v.IsApplicableForAllComputeModes }).(pulumi.BoolPtrOutput)
}

func (o StampCapacityOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StampCapacity) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StampCapacityOutput) SiteMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StampCapacity) *string { return v.SiteMode }).(pulumi.StringPtrOutput)
}

func (o StampCapacityOutput) TotalCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v StampCapacity) *float64 { return v.TotalCapacity }).(pulumi.Float64PtrOutput)
}

func (o StampCapacityOutput) Unit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StampCapacity) *string { return v.Unit }).(pulumi.StringPtrOutput)
}

func (o StampCapacityOutput) WorkerSize() WorkerSizeOptionsPtrOutput {
	return o.ApplyT(func(v StampCapacity) *WorkerSizeOptions { return v.WorkerSize }).(WorkerSizeOptionsPtrOutput)
}

func (o StampCapacityOutput) WorkerSizeId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StampCapacity) *int { return v.WorkerSizeId }).(pulumi.IntPtrOutput)
}

type StampCapacityArrayOutput struct{ *pulumi.OutputState }

func (StampCapacityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StampCapacity)(nil)).Elem()
}

func (o StampCapacityArrayOutput) ToStampCapacityArrayOutput() StampCapacityArrayOutput {
	return o
}

func (o StampCapacityArrayOutput) ToStampCapacityArrayOutputWithContext(ctx context.Context) StampCapacityArrayOutput {
	return o
}

func (o StampCapacityArrayOutput) Index(i pulumi.IntInput) StampCapacityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StampCapacity {
		return vs[0].([]StampCapacity)[vs[1].(int)]
	}).(StampCapacityOutput)
}

type StampCapacityResponse struct {
	AvailableCapacity              *float64 `pulumi:"availableCapacity"`
	ComputeMode                    *string  `pulumi:"computeMode"`
	ExcludeFromCapacityAllocation  *bool    `pulumi:"excludeFromCapacityAllocation"`
	IsApplicableForAllComputeModes *bool    `pulumi:"isApplicableForAllComputeModes"`
	Name                           *string  `pulumi:"name"`
	SiteMode                       *string  `pulumi:"siteMode"`
	TotalCapacity                  *float64 `pulumi:"totalCapacity"`
	Unit                           *string  `pulumi:"unit"`
	WorkerSize                     *string  `pulumi:"workerSize"`
	WorkerSizeId                   *int     `pulumi:"workerSizeId"`
}

type StampCapacityResponseOutput struct{ *pulumi.OutputState }

func (StampCapacityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StampCapacityResponse)(nil)).Elem()
}

func (o StampCapacityResponseOutput) ToStampCapacityResponseOutput() StampCapacityResponseOutput {
	return o
}

func (o StampCapacityResponseOutput) ToStampCapacityResponseOutputWithContext(ctx context.Context) StampCapacityResponseOutput {
	return o
}

func (o StampCapacityResponseOutput) AvailableCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *float64 { return v.AvailableCapacity }).(pulumi.Float64PtrOutput)
}

func (o StampCapacityResponseOutput) ComputeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *string { return v.ComputeMode }).(pulumi.StringPtrOutput)
}

func (o StampCapacityResponseOutput) ExcludeFromCapacityAllocation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *bool { return v.ExcludeFromCapacityAllocation }).(pulumi.BoolPtrOutput)
}

func (o StampCapacityResponseOutput) IsApplicableForAllComputeModes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *bool { return v.IsApplicableForAllComputeModes }).(pulumi.BoolPtrOutput)
}

func (o StampCapacityResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StampCapacityResponseOutput) SiteMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *string { return v.SiteMode }).(pulumi.StringPtrOutput)
}

func (o StampCapacityResponseOutput) TotalCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *float64 { return v.TotalCapacity }).(pulumi.Float64PtrOutput)
}

func (o StampCapacityResponseOutput) Unit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *string { return v.Unit }).(pulumi.StringPtrOutput)
}

func (o StampCapacityResponseOutput) WorkerSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *string { return v.WorkerSize }).(pulumi.StringPtrOutput)
}

func (o StampCapacityResponseOutput) WorkerSizeId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *int { return v.WorkerSizeId }).(pulumi.IntPtrOutput)
}

type StampCapacityResponseArrayOutput struct{ *pulumi.OutputState }

func (StampCapacityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StampCapacityResponse)(nil)).Elem()
}

func (o StampCapacityResponseArrayOutput) ToStampCapacityResponseArrayOutput() StampCapacityResponseArrayOutput {
	return o
}

func (o StampCapacityResponseArrayOutput) ToStampCapacityResponseArrayOutputWithContext(ctx context.Context) StampCapacityResponseArrayOutput {
	return o
}

func (o StampCapacityResponseArrayOutput) Index(i pulumi.IntInput) StampCapacityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StampCapacityResponse {
		return vs[0].([]StampCapacityResponse)[vs[1].(int)]
	}).(StampCapacityResponseOutput)
}

type StatusCodesBasedTrigger struct {
	Count        *int    `pulumi:"count"`
	Status       *int    `pulumi:"status"`
	SubStatus    *int    `pulumi:"subStatus"`
	TimeInterval *string `pulumi:"timeInterval"`
	Win32Status  *int    `pulumi:"win32Status"`
}





type StatusCodesBasedTriggerInput interface {
	pulumi.Input

	ToStatusCodesBasedTriggerOutput() StatusCodesBasedTriggerOutput
	ToStatusCodesBasedTriggerOutputWithContext(context.Context) StatusCodesBasedTriggerOutput
}

type StatusCodesBasedTriggerArgs struct {
	Count        pulumi.IntPtrInput    `pulumi:"count"`
	Status       pulumi.IntPtrInput    `pulumi:"status"`
	SubStatus    pulumi.IntPtrInput    `pulumi:"subStatus"`
	TimeInterval pulumi.StringPtrInput `pulumi:"timeInterval"`
	Win32Status  pulumi.IntPtrInput    `pulumi:"win32Status"`
}

func (StatusCodesBasedTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusCodesBasedTrigger)(nil)).Elem()
}

func (i StatusCodesBasedTriggerArgs) ToStatusCodesBasedTriggerOutput() StatusCodesBasedTriggerOutput {
	return i.ToStatusCodesBasedTriggerOutputWithContext(context.Background())
}

func (i StatusCodesBasedTriggerArgs) ToStatusCodesBasedTriggerOutputWithContext(ctx context.Context) StatusCodesBasedTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusCodesBasedTriggerOutput)
}





type StatusCodesBasedTriggerArrayInput interface {
	pulumi.Input

	ToStatusCodesBasedTriggerArrayOutput() StatusCodesBasedTriggerArrayOutput
	ToStatusCodesBasedTriggerArrayOutputWithContext(context.Context) StatusCodesBasedTriggerArrayOutput
}

type StatusCodesBasedTriggerArray []StatusCodesBasedTriggerInput

func (StatusCodesBasedTriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusCodesBasedTrigger)(nil)).Elem()
}

func (i StatusCodesBasedTriggerArray) ToStatusCodesBasedTriggerArrayOutput() StatusCodesBasedTriggerArrayOutput {
	return i.ToStatusCodesBasedTriggerArrayOutputWithContext(context.Background())
}

func (i StatusCodesBasedTriggerArray) ToStatusCodesBasedTriggerArrayOutputWithContext(ctx context.Context) StatusCodesBasedTriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusCodesBasedTriggerArrayOutput)
}

type StatusCodesBasedTriggerOutput struct{ *pulumi.OutputState }

func (StatusCodesBasedTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusCodesBasedTrigger)(nil)).Elem()
}

func (o StatusCodesBasedTriggerOutput) ToStatusCodesBasedTriggerOutput() StatusCodesBasedTriggerOutput {
	return o
}

func (o StatusCodesBasedTriggerOutput) ToStatusCodesBasedTriggerOutputWithContext(ctx context.Context) StatusCodesBasedTriggerOutput {
	return o
}

func (o StatusCodesBasedTriggerOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTrigger) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o StatusCodesBasedTriggerOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTrigger) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func (o StatusCodesBasedTriggerOutput) SubStatus() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTrigger) *int { return v.SubStatus }).(pulumi.IntPtrOutput)
}

func (o StatusCodesBasedTriggerOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTrigger) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

func (o StatusCodesBasedTriggerOutput) Win32Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTrigger) *int { return v.Win32Status }).(pulumi.IntPtrOutput)
}

type StatusCodesBasedTriggerArrayOutput struct{ *pulumi.OutputState }

func (StatusCodesBasedTriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusCodesBasedTrigger)(nil)).Elem()
}

func (o StatusCodesBasedTriggerArrayOutput) ToStatusCodesBasedTriggerArrayOutput() StatusCodesBasedTriggerArrayOutput {
	return o
}

func (o StatusCodesBasedTriggerArrayOutput) ToStatusCodesBasedTriggerArrayOutputWithContext(ctx context.Context) StatusCodesBasedTriggerArrayOutput {
	return o
}

func (o StatusCodesBasedTriggerArrayOutput) Index(i pulumi.IntInput) StatusCodesBasedTriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StatusCodesBasedTrigger {
		return vs[0].([]StatusCodesBasedTrigger)[vs[1].(int)]
	}).(StatusCodesBasedTriggerOutput)
}

type StatusCodesBasedTriggerResponse struct {
	Count        *int    `pulumi:"count"`
	Status       *int    `pulumi:"status"`
	SubStatus    *int    `pulumi:"subStatus"`
	TimeInterval *string `pulumi:"timeInterval"`
	Win32Status  *int    `pulumi:"win32Status"`
}

type StatusCodesBasedTriggerResponseOutput struct{ *pulumi.OutputState }

func (StatusCodesBasedTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusCodesBasedTriggerResponse)(nil)).Elem()
}

func (o StatusCodesBasedTriggerResponseOutput) ToStatusCodesBasedTriggerResponseOutput() StatusCodesBasedTriggerResponseOutput {
	return o
}

func (o StatusCodesBasedTriggerResponseOutput) ToStatusCodesBasedTriggerResponseOutputWithContext(ctx context.Context) StatusCodesBasedTriggerResponseOutput {
	return o
}

func (o StatusCodesBasedTriggerResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTriggerResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o StatusCodesBasedTriggerResponseOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTriggerResponse) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func (o StatusCodesBasedTriggerResponseOutput) SubStatus() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTriggerResponse) *int { return v.SubStatus }).(pulumi.IntPtrOutput)
}

func (o StatusCodesBasedTriggerResponseOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTriggerResponse) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

func (o StatusCodesBasedTriggerResponseOutput) Win32Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTriggerResponse) *int { return v.Win32Status }).(pulumi.IntPtrOutput)
}

type StatusCodesBasedTriggerResponseArrayOutput struct{ *pulumi.OutputState }

func (StatusCodesBasedTriggerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusCodesBasedTriggerResponse)(nil)).Elem()
}

func (o StatusCodesBasedTriggerResponseArrayOutput) ToStatusCodesBasedTriggerResponseArrayOutput() StatusCodesBasedTriggerResponseArrayOutput {
	return o
}

func (o StatusCodesBasedTriggerResponseArrayOutput) ToStatusCodesBasedTriggerResponseArrayOutputWithContext(ctx context.Context) StatusCodesBasedTriggerResponseArrayOutput {
	return o
}

func (o StatusCodesBasedTriggerResponseArrayOutput) Index(i pulumi.IntInput) StatusCodesBasedTriggerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StatusCodesBasedTriggerResponse {
		return vs[0].([]StatusCodesBasedTriggerResponse)[vs[1].(int)]
	}).(StatusCodesBasedTriggerResponseOutput)
}

type VirtualApplication struct {
	PhysicalPath       *string            `pulumi:"physicalPath"`
	PreloadEnabled     *bool              `pulumi:"preloadEnabled"`
	VirtualDirectories []VirtualDirectory `pulumi:"virtualDirectories"`
	VirtualPath        *string            `pulumi:"virtualPath"`
}





type VirtualApplicationInput interface {
	pulumi.Input

	ToVirtualApplicationOutput() VirtualApplicationOutput
	ToVirtualApplicationOutputWithContext(context.Context) VirtualApplicationOutput
}

type VirtualApplicationArgs struct {
	PhysicalPath       pulumi.StringPtrInput      `pulumi:"physicalPath"`
	PreloadEnabled     pulumi.BoolPtrInput        `pulumi:"preloadEnabled"`
	VirtualDirectories VirtualDirectoryArrayInput `pulumi:"virtualDirectories"`
	VirtualPath        pulumi.StringPtrInput      `pulumi:"virtualPath"`
}

func (VirtualApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualApplication)(nil)).Elem()
}

func (i VirtualApplicationArgs) ToVirtualApplicationOutput() VirtualApplicationOutput {
	return i.ToVirtualApplicationOutputWithContext(context.Background())
}

func (i VirtualApplicationArgs) ToVirtualApplicationOutputWithContext(ctx context.Context) VirtualApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualApplicationOutput)
}





type VirtualApplicationArrayInput interface {
	pulumi.Input

	ToVirtualApplicationArrayOutput() VirtualApplicationArrayOutput
	ToVirtualApplicationArrayOutputWithContext(context.Context) VirtualApplicationArrayOutput
}

type VirtualApplicationArray []VirtualApplicationInput

func (VirtualApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualApplication)(nil)).Elem()
}

func (i VirtualApplicationArray) ToVirtualApplicationArrayOutput() VirtualApplicationArrayOutput {
	return i.ToVirtualApplicationArrayOutputWithContext(context.Background())
}

func (i VirtualApplicationArray) ToVirtualApplicationArrayOutputWithContext(ctx context.Context) VirtualApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualApplicationArrayOutput)
}

type VirtualApplicationOutput struct{ *pulumi.OutputState }

func (VirtualApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualApplication)(nil)).Elem()
}

func (o VirtualApplicationOutput) ToVirtualApplicationOutput() VirtualApplicationOutput {
	return o
}

func (o VirtualApplicationOutput) ToVirtualApplicationOutputWithContext(ctx context.Context) VirtualApplicationOutput {
	return o
}

func (o VirtualApplicationOutput) PhysicalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualApplication) *string { return v.PhysicalPath }).(pulumi.StringPtrOutput)
}

func (o VirtualApplicationOutput) PreloadEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualApplication) *bool { return v.PreloadEnabled }).(pulumi.BoolPtrOutput)
}

func (o VirtualApplicationOutput) VirtualDirectories() VirtualDirectoryArrayOutput {
	return o.ApplyT(func(v VirtualApplication) []VirtualDirectory { return v.VirtualDirectories }).(VirtualDirectoryArrayOutput)
}

func (o VirtualApplicationOutput) VirtualPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualApplication) *string { return v.VirtualPath }).(pulumi.StringPtrOutput)
}

type VirtualApplicationArrayOutput struct{ *pulumi.OutputState }

func (VirtualApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualApplication)(nil)).Elem()
}

func (o VirtualApplicationArrayOutput) ToVirtualApplicationArrayOutput() VirtualApplicationArrayOutput {
	return o
}

func (o VirtualApplicationArrayOutput) ToVirtualApplicationArrayOutputWithContext(ctx context.Context) VirtualApplicationArrayOutput {
	return o
}

func (o VirtualApplicationArrayOutput) Index(i pulumi.IntInput) VirtualApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualApplication {
		return vs[0].([]VirtualApplication)[vs[1].(int)]
	}).(VirtualApplicationOutput)
}

type VirtualApplicationResponse struct {
	PhysicalPath       *string                    `pulumi:"physicalPath"`
	PreloadEnabled     *bool                      `pulumi:"preloadEnabled"`
	VirtualDirectories []VirtualDirectoryResponse `pulumi:"virtualDirectories"`
	VirtualPath        *string                    `pulumi:"virtualPath"`
}

type VirtualApplicationResponseOutput struct{ *pulumi.OutputState }

func (VirtualApplicationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualApplicationResponse)(nil)).Elem()
}

func (o VirtualApplicationResponseOutput) ToVirtualApplicationResponseOutput() VirtualApplicationResponseOutput {
	return o
}

func (o VirtualApplicationResponseOutput) ToVirtualApplicationResponseOutputWithContext(ctx context.Context) VirtualApplicationResponseOutput {
	return o
}

func (o VirtualApplicationResponseOutput) PhysicalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualApplicationResponse) *string { return v.PhysicalPath }).(pulumi.StringPtrOutput)
}

func (o VirtualApplicationResponseOutput) PreloadEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualApplicationResponse) *bool { return v.PreloadEnabled }).(pulumi.BoolPtrOutput)
}

func (o VirtualApplicationResponseOutput) VirtualDirectories() VirtualDirectoryResponseArrayOutput {
	return o.ApplyT(func(v VirtualApplicationResponse) []VirtualDirectoryResponse { return v.VirtualDirectories }).(VirtualDirectoryResponseArrayOutput)
}

func (o VirtualApplicationResponseOutput) VirtualPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualApplicationResponse) *string { return v.VirtualPath }).(pulumi.StringPtrOutput)
}

type VirtualApplicationResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualApplicationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualApplicationResponse)(nil)).Elem()
}

func (o VirtualApplicationResponseArrayOutput) ToVirtualApplicationResponseArrayOutput() VirtualApplicationResponseArrayOutput {
	return o
}

func (o VirtualApplicationResponseArrayOutput) ToVirtualApplicationResponseArrayOutputWithContext(ctx context.Context) VirtualApplicationResponseArrayOutput {
	return o
}

func (o VirtualApplicationResponseArrayOutput) Index(i pulumi.IntInput) VirtualApplicationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualApplicationResponse {
		return vs[0].([]VirtualApplicationResponse)[vs[1].(int)]
	}).(VirtualApplicationResponseOutput)
}

type VirtualDirectory struct {
	PhysicalPath *string `pulumi:"physicalPath"`
	VirtualPath  *string `pulumi:"virtualPath"`
}





type VirtualDirectoryInput interface {
	pulumi.Input

	ToVirtualDirectoryOutput() VirtualDirectoryOutput
	ToVirtualDirectoryOutputWithContext(context.Context) VirtualDirectoryOutput
}

type VirtualDirectoryArgs struct {
	PhysicalPath pulumi.StringPtrInput `pulumi:"physicalPath"`
	VirtualPath  pulumi.StringPtrInput `pulumi:"virtualPath"`
}

func (VirtualDirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDirectory)(nil)).Elem()
}

func (i VirtualDirectoryArgs) ToVirtualDirectoryOutput() VirtualDirectoryOutput {
	return i.ToVirtualDirectoryOutputWithContext(context.Background())
}

func (i VirtualDirectoryArgs) ToVirtualDirectoryOutputWithContext(ctx context.Context) VirtualDirectoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDirectoryOutput)
}





type VirtualDirectoryArrayInput interface {
	pulumi.Input

	ToVirtualDirectoryArrayOutput() VirtualDirectoryArrayOutput
	ToVirtualDirectoryArrayOutputWithContext(context.Context) VirtualDirectoryArrayOutput
}

type VirtualDirectoryArray []VirtualDirectoryInput

func (VirtualDirectoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDirectory)(nil)).Elem()
}

func (i VirtualDirectoryArray) ToVirtualDirectoryArrayOutput() VirtualDirectoryArrayOutput {
	return i.ToVirtualDirectoryArrayOutputWithContext(context.Background())
}

func (i VirtualDirectoryArray) ToVirtualDirectoryArrayOutputWithContext(ctx context.Context) VirtualDirectoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDirectoryArrayOutput)
}

type VirtualDirectoryOutput struct{ *pulumi.OutputState }

func (VirtualDirectoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDirectory)(nil)).Elem()
}

func (o VirtualDirectoryOutput) ToVirtualDirectoryOutput() VirtualDirectoryOutput {
	return o
}

func (o VirtualDirectoryOutput) ToVirtualDirectoryOutputWithContext(ctx context.Context) VirtualDirectoryOutput {
	return o
}

func (o VirtualDirectoryOutput) PhysicalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDirectory) *string { return v.PhysicalPath }).(pulumi.StringPtrOutput)
}

func (o VirtualDirectoryOutput) VirtualPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDirectory) *string { return v.VirtualPath }).(pulumi.StringPtrOutput)
}

type VirtualDirectoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualDirectoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDirectory)(nil)).Elem()
}

func (o VirtualDirectoryArrayOutput) ToVirtualDirectoryArrayOutput() VirtualDirectoryArrayOutput {
	return o
}

func (o VirtualDirectoryArrayOutput) ToVirtualDirectoryArrayOutputWithContext(ctx context.Context) VirtualDirectoryArrayOutput {
	return o
}

func (o VirtualDirectoryArrayOutput) Index(i pulumi.IntInput) VirtualDirectoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualDirectory {
		return vs[0].([]VirtualDirectory)[vs[1].(int)]
	}).(VirtualDirectoryOutput)
}

type VirtualDirectoryResponse struct {
	PhysicalPath *string `pulumi:"physicalPath"`
	VirtualPath  *string `pulumi:"virtualPath"`
}

type VirtualDirectoryResponseOutput struct{ *pulumi.OutputState }

func (VirtualDirectoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDirectoryResponse)(nil)).Elem()
}

func (o VirtualDirectoryResponseOutput) ToVirtualDirectoryResponseOutput() VirtualDirectoryResponseOutput {
	return o
}

func (o VirtualDirectoryResponseOutput) ToVirtualDirectoryResponseOutputWithContext(ctx context.Context) VirtualDirectoryResponseOutput {
	return o
}

func (o VirtualDirectoryResponseOutput) PhysicalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDirectoryResponse) *string { return v.PhysicalPath }).(pulumi.StringPtrOutput)
}

func (o VirtualDirectoryResponseOutput) VirtualPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDirectoryResponse) *string { return v.VirtualPath }).(pulumi.StringPtrOutput)
}

type VirtualDirectoryResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualDirectoryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDirectoryResponse)(nil)).Elem()
}

func (o VirtualDirectoryResponseArrayOutput) ToVirtualDirectoryResponseArrayOutput() VirtualDirectoryResponseArrayOutput {
	return o
}

func (o VirtualDirectoryResponseArrayOutput) ToVirtualDirectoryResponseArrayOutputWithContext(ctx context.Context) VirtualDirectoryResponseArrayOutput {
	return o
}

func (o VirtualDirectoryResponseArrayOutput) Index(i pulumi.IntInput) VirtualDirectoryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualDirectoryResponse {
		return vs[0].([]VirtualDirectoryResponse)[vs[1].(int)]
	}).(VirtualDirectoryResponseOutput)
}

type VirtualIPMapping struct {
	InUse             *bool   `pulumi:"inUse"`
	InternalHttpPort  *int    `pulumi:"internalHttpPort"`
	InternalHttpsPort *int    `pulumi:"internalHttpsPort"`
	VirtualIP         *string `pulumi:"virtualIP"`
}





type VirtualIPMappingInput interface {
	pulumi.Input

	ToVirtualIPMappingOutput() VirtualIPMappingOutput
	ToVirtualIPMappingOutputWithContext(context.Context) VirtualIPMappingOutput
}

type VirtualIPMappingArgs struct {
	InUse             pulumi.BoolPtrInput   `pulumi:"inUse"`
	InternalHttpPort  pulumi.IntPtrInput    `pulumi:"internalHttpPort"`
	InternalHttpsPort pulumi.IntPtrInput    `pulumi:"internalHttpsPort"`
	VirtualIP         pulumi.StringPtrInput `pulumi:"virtualIP"`
}

func (VirtualIPMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualIPMapping)(nil)).Elem()
}

func (i VirtualIPMappingArgs) ToVirtualIPMappingOutput() VirtualIPMappingOutput {
	return i.ToVirtualIPMappingOutputWithContext(context.Background())
}

func (i VirtualIPMappingArgs) ToVirtualIPMappingOutputWithContext(ctx context.Context) VirtualIPMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualIPMappingOutput)
}





type VirtualIPMappingArrayInput interface {
	pulumi.Input

	ToVirtualIPMappingArrayOutput() VirtualIPMappingArrayOutput
	ToVirtualIPMappingArrayOutputWithContext(context.Context) VirtualIPMappingArrayOutput
}

type VirtualIPMappingArray []VirtualIPMappingInput

func (VirtualIPMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualIPMapping)(nil)).Elem()
}

func (i VirtualIPMappingArray) ToVirtualIPMappingArrayOutput() VirtualIPMappingArrayOutput {
	return i.ToVirtualIPMappingArrayOutputWithContext(context.Background())
}

func (i VirtualIPMappingArray) ToVirtualIPMappingArrayOutputWithContext(ctx context.Context) VirtualIPMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualIPMappingArrayOutput)
}

type VirtualIPMappingOutput struct{ *pulumi.OutputState }

func (VirtualIPMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualIPMapping)(nil)).Elem()
}

func (o VirtualIPMappingOutput) ToVirtualIPMappingOutput() VirtualIPMappingOutput {
	return o
}

func (o VirtualIPMappingOutput) ToVirtualIPMappingOutputWithContext(ctx context.Context) VirtualIPMappingOutput {
	return o
}

func (o VirtualIPMappingOutput) InUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualIPMapping) *bool { return v.InUse }).(pulumi.BoolPtrOutput)
}

func (o VirtualIPMappingOutput) InternalHttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualIPMapping) *int { return v.InternalHttpPort }).(pulumi.IntPtrOutput)
}

func (o VirtualIPMappingOutput) InternalHttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualIPMapping) *int { return v.InternalHttpsPort }).(pulumi.IntPtrOutput)
}

func (o VirtualIPMappingOutput) VirtualIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualIPMapping) *string { return v.VirtualIP }).(pulumi.StringPtrOutput)
}

type VirtualIPMappingArrayOutput struct{ *pulumi.OutputState }

func (VirtualIPMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualIPMapping)(nil)).Elem()
}

func (o VirtualIPMappingArrayOutput) ToVirtualIPMappingArrayOutput() VirtualIPMappingArrayOutput {
	return o
}

func (o VirtualIPMappingArrayOutput) ToVirtualIPMappingArrayOutputWithContext(ctx context.Context) VirtualIPMappingArrayOutput {
	return o
}

func (o VirtualIPMappingArrayOutput) Index(i pulumi.IntInput) VirtualIPMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualIPMapping {
		return vs[0].([]VirtualIPMapping)[vs[1].(int)]
	}).(VirtualIPMappingOutput)
}

type VirtualIPMappingResponse struct {
	InUse             *bool   `pulumi:"inUse"`
	InternalHttpPort  *int    `pulumi:"internalHttpPort"`
	InternalHttpsPort *int    `pulumi:"internalHttpsPort"`
	VirtualIP         *string `pulumi:"virtualIP"`
}

type VirtualIPMappingResponseOutput struct{ *pulumi.OutputState }

func (VirtualIPMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualIPMappingResponse)(nil)).Elem()
}

func (o VirtualIPMappingResponseOutput) ToVirtualIPMappingResponseOutput() VirtualIPMappingResponseOutput {
	return o
}

func (o VirtualIPMappingResponseOutput) ToVirtualIPMappingResponseOutputWithContext(ctx context.Context) VirtualIPMappingResponseOutput {
	return o
}

func (o VirtualIPMappingResponseOutput) InUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualIPMappingResponse) *bool { return v.InUse }).(pulumi.BoolPtrOutput)
}

func (o VirtualIPMappingResponseOutput) InternalHttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualIPMappingResponse) *int { return v.InternalHttpPort }).(pulumi.IntPtrOutput)
}

func (o VirtualIPMappingResponseOutput) InternalHttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualIPMappingResponse) *int { return v.InternalHttpsPort }).(pulumi.IntPtrOutput)
}

func (o VirtualIPMappingResponseOutput) VirtualIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualIPMappingResponse) *string { return v.VirtualIP }).(pulumi.StringPtrOutput)
}

type VirtualIPMappingResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualIPMappingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualIPMappingResponse)(nil)).Elem()
}

func (o VirtualIPMappingResponseArrayOutput) ToVirtualIPMappingResponseArrayOutput() VirtualIPMappingResponseArrayOutput {
	return o
}

func (o VirtualIPMappingResponseArrayOutput) ToVirtualIPMappingResponseArrayOutputWithContext(ctx context.Context) VirtualIPMappingResponseArrayOutput {
	return o
}

func (o VirtualIPMappingResponseArrayOutput) Index(i pulumi.IntInput) VirtualIPMappingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualIPMappingResponse {
		return vs[0].([]VirtualIPMappingResponse)[vs[1].(int)]
	}).(VirtualIPMappingResponseOutput)
}

type VirtualNetworkProfile struct {
	Id     *string `pulumi:"id"`
	Name   *string `pulumi:"name"`
	Subnet *string `pulumi:"subnet"`
	Type   *string `pulumi:"type"`
}





type VirtualNetworkProfileInput interface {
	pulumi.Input

	ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput
	ToVirtualNetworkProfileOutputWithContext(context.Context) VirtualNetworkProfileOutput
}

type VirtualNetworkProfileArgs struct {
	Id     pulumi.StringPtrInput `pulumi:"id"`
	Name   pulumi.StringPtrInput `pulumi:"name"`
	Subnet pulumi.StringPtrInput `pulumi:"subnet"`
	Type   pulumi.StringPtrInput `pulumi:"type"`
}

func (VirtualNetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfile)(nil)).Elem()
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput {
	return i.ToVirtualNetworkProfileOutputWithContext(context.Background())
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfileOutputWithContext(ctx context.Context) VirtualNetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkProfileOutput)
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput {
	return i.ToVirtualNetworkProfilePtrOutputWithContext(context.Background())
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkProfileOutput).ToVirtualNetworkProfilePtrOutputWithContext(ctx)
}









type VirtualNetworkProfilePtrInput interface {
	pulumi.Input

	ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput
	ToVirtualNetworkProfilePtrOutputWithContext(context.Context) VirtualNetworkProfilePtrOutput
}

type virtualNetworkProfilePtrType VirtualNetworkProfileArgs

func VirtualNetworkProfilePtr(v *VirtualNetworkProfileArgs) VirtualNetworkProfilePtrInput {
	return (*virtualNetworkProfilePtrType)(v)
}

func (*virtualNetworkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkProfile)(nil)).Elem()
}

func (i *virtualNetworkProfilePtrType) ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput {
	return i.ToVirtualNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *virtualNetworkProfilePtrType) ToVirtualNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkProfilePtrOutput)
}

type VirtualNetworkProfileOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfile)(nil)).Elem()
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput {
	return o
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfileOutputWithContext(ctx context.Context) VirtualNetworkProfileOutput {
	return o
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput {
	return o.ToVirtualNetworkProfilePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualNetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkProfile) *VirtualNetworkProfile {
		return &v
	}).(VirtualNetworkProfilePtrOutput)
}

func (o VirtualNetworkProfileOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfile) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfile) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfile) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfile) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VirtualNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkProfile)(nil)).Elem()
}

func (o VirtualNetworkProfilePtrOutput) ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput {
	return o
}

func (o VirtualNetworkProfilePtrOutput) ToVirtualNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualNetworkProfilePtrOutput {
	return o
}

func (o VirtualNetworkProfilePtrOutput) Elem() VirtualNetworkProfileOutput {
	return o.ApplyT(func(v *VirtualNetworkProfile) VirtualNetworkProfile {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkProfile
		return ret
	}).(VirtualNetworkProfileOutput)
}

func (o VirtualNetworkProfilePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfilePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfilePtrOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfilePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkProfileResponse struct {
	Id     *string `pulumi:"id"`
	Name   *string `pulumi:"name"`
	Subnet *string `pulumi:"subnet"`
	Type   *string `pulumi:"type"`
}

type VirtualNetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfileResponse)(nil)).Elem()
}

func (o VirtualNetworkProfileResponseOutput) ToVirtualNetworkProfileResponseOutput() VirtualNetworkProfileResponseOutput {
	return o
}

func (o VirtualNetworkProfileResponseOutput) ToVirtualNetworkProfileResponseOutputWithContext(ctx context.Context) VirtualNetworkProfileResponseOutput {
	return o
}

func (o VirtualNetworkProfileResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileResponseOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VirtualNetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkProfileResponse)(nil)).Elem()
}

func (o VirtualNetworkProfileResponsePtrOutput) ToVirtualNetworkProfileResponsePtrOutput() VirtualNetworkProfileResponsePtrOutput {
	return o
}

func (o VirtualNetworkProfileResponsePtrOutput) ToVirtualNetworkProfileResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkProfileResponsePtrOutput {
	return o
}

func (o VirtualNetworkProfileResponsePtrOutput) Elem() VirtualNetworkProfileResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkProfileResponse) VirtualNetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkProfileResponse
		return ret
	}).(VirtualNetworkProfileResponseOutput)
}

func (o VirtualNetworkProfileResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileResponsePtrOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type VnetRoute struct {
	EndAddress   *string           `pulumi:"endAddress"`
	Id           *string           `pulumi:"id"`
	Kind         *string           `pulumi:"kind"`
	Location     string            `pulumi:"location"`
	Name         *string           `pulumi:"name"`
	RouteType    *string           `pulumi:"routeType"`
	StartAddress *string           `pulumi:"startAddress"`
	Tags         map[string]string `pulumi:"tags"`
	Type         *string           `pulumi:"type"`
}





type VnetRouteInput interface {
	pulumi.Input

	ToVnetRouteOutput() VnetRouteOutput
	ToVnetRouteOutputWithContext(context.Context) VnetRouteOutput
}

type VnetRouteArgs struct {
	EndAddress   pulumi.StringPtrInput `pulumi:"endAddress"`
	Id           pulumi.StringPtrInput `pulumi:"id"`
	Kind         pulumi.StringPtrInput `pulumi:"kind"`
	Location     pulumi.StringInput    `pulumi:"location"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
	RouteType    pulumi.StringPtrInput `pulumi:"routeType"`
	StartAddress pulumi.StringPtrInput `pulumi:"startAddress"`
	Tags         pulumi.StringMapInput `pulumi:"tags"`
	Type         pulumi.StringPtrInput `pulumi:"type"`
}

func (VnetRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VnetRoute)(nil)).Elem()
}

func (i VnetRouteArgs) ToVnetRouteOutput() VnetRouteOutput {
	return i.ToVnetRouteOutputWithContext(context.Background())
}

func (i VnetRouteArgs) ToVnetRouteOutputWithContext(ctx context.Context) VnetRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnetRouteOutput)
}





type VnetRouteArrayInput interface {
	pulumi.Input

	ToVnetRouteArrayOutput() VnetRouteArrayOutput
	ToVnetRouteArrayOutputWithContext(context.Context) VnetRouteArrayOutput
}

type VnetRouteArray []VnetRouteInput

func (VnetRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VnetRoute)(nil)).Elem()
}

func (i VnetRouteArray) ToVnetRouteArrayOutput() VnetRouteArrayOutput {
	return i.ToVnetRouteArrayOutputWithContext(context.Background())
}

func (i VnetRouteArray) ToVnetRouteArrayOutputWithContext(ctx context.Context) VnetRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnetRouteArrayOutput)
}

type VnetRouteOutput struct{ *pulumi.OutputState }

func (VnetRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VnetRoute)(nil)).Elem()
}

func (o VnetRouteOutput) ToVnetRouteOutput() VnetRouteOutput {
	return o
}

func (o VnetRouteOutput) ToVnetRouteOutputWithContext(ctx context.Context) VnetRouteOutput {
	return o
}

func (o VnetRouteOutput) EndAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRoute) *string { return v.EndAddress }).(pulumi.StringPtrOutput)
}

func (o VnetRouteOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRoute) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VnetRouteOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRoute) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o VnetRouteOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v VnetRoute) string { return v.Location }).(pulumi.StringOutput)
}

func (o VnetRouteOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRoute) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VnetRouteOutput) RouteType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRoute) *string { return v.RouteType }).(pulumi.StringPtrOutput)
}

func (o VnetRouteOutput) StartAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRoute) *string { return v.StartAddress }).(pulumi.StringPtrOutput)
}

func (o VnetRouteOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VnetRoute) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VnetRouteOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRoute) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VnetRouteArrayOutput struct{ *pulumi.OutputState }

func (VnetRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VnetRoute)(nil)).Elem()
}

func (o VnetRouteArrayOutput) ToVnetRouteArrayOutput() VnetRouteArrayOutput {
	return o
}

func (o VnetRouteArrayOutput) ToVnetRouteArrayOutputWithContext(ctx context.Context) VnetRouteArrayOutput {
	return o
}

func (o VnetRouteArrayOutput) Index(i pulumi.IntInput) VnetRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VnetRoute {
		return vs[0].([]VnetRoute)[vs[1].(int)]
	}).(VnetRouteOutput)
}

type VnetRouteResponse struct {
	EndAddress   *string           `pulumi:"endAddress"`
	Id           *string           `pulumi:"id"`
	Kind         *string           `pulumi:"kind"`
	Location     string            `pulumi:"location"`
	Name         *string           `pulumi:"name"`
	RouteType    *string           `pulumi:"routeType"`
	StartAddress *string           `pulumi:"startAddress"`
	Tags         map[string]string `pulumi:"tags"`
	Type         *string           `pulumi:"type"`
}

type VnetRouteResponseOutput struct{ *pulumi.OutputState }

func (VnetRouteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VnetRouteResponse)(nil)).Elem()
}

func (o VnetRouteResponseOutput) ToVnetRouteResponseOutput() VnetRouteResponseOutput {
	return o
}

func (o VnetRouteResponseOutput) ToVnetRouteResponseOutputWithContext(ctx context.Context) VnetRouteResponseOutput {
	return o
}

func (o VnetRouteResponseOutput) EndAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRouteResponse) *string { return v.EndAddress }).(pulumi.StringPtrOutput)
}

func (o VnetRouteResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRouteResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VnetRouteResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRouteResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o VnetRouteResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v VnetRouteResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o VnetRouteResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRouteResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VnetRouteResponseOutput) RouteType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRouteResponse) *string { return v.RouteType }).(pulumi.StringPtrOutput)
}

func (o VnetRouteResponseOutput) StartAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRouteResponse) *string { return v.StartAddress }).(pulumi.StringPtrOutput)
}

func (o VnetRouteResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VnetRouteResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VnetRouteResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRouteResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VnetRouteResponseArrayOutput struct{ *pulumi.OutputState }

func (VnetRouteResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VnetRouteResponse)(nil)).Elem()
}

func (o VnetRouteResponseArrayOutput) ToVnetRouteResponseArrayOutput() VnetRouteResponseArrayOutput {
	return o
}

func (o VnetRouteResponseArrayOutput) ToVnetRouteResponseArrayOutputWithContext(ctx context.Context) VnetRouteResponseArrayOutput {
	return o
}

func (o VnetRouteResponseArrayOutput) Index(i pulumi.IntInput) VnetRouteResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VnetRouteResponse {
		return vs[0].([]VnetRouteResponse)[vs[1].(int)]
	}).(VnetRouteResponseOutput)
}

type WorkerPool struct {
	ComputeMode   *ComputeModeOptions `pulumi:"computeMode"`
	Id            *string             `pulumi:"id"`
	InstanceNames []string            `pulumi:"instanceNames"`
	Kind          *string             `pulumi:"kind"`
	Location      string              `pulumi:"location"`
	Name          *string             `pulumi:"name"`
	Sku           *SkuDescription     `pulumi:"sku"`
	Tags          map[string]string   `pulumi:"tags"`
	Type          *string             `pulumi:"type"`
	WorkerCount   *int                `pulumi:"workerCount"`
	WorkerSize    *string             `pulumi:"workerSize"`
	WorkerSizeId  *int                `pulumi:"workerSizeId"`
}





type WorkerPoolInput interface {
	pulumi.Input

	ToWorkerPoolOutput() WorkerPoolOutput
	ToWorkerPoolOutputWithContext(context.Context) WorkerPoolOutput
}

type WorkerPoolArgs struct {
	ComputeMode   ComputeModeOptionsPtrInput `pulumi:"computeMode"`
	Id            pulumi.StringPtrInput      `pulumi:"id"`
	InstanceNames pulumi.StringArrayInput    `pulumi:"instanceNames"`
	Kind          pulumi.StringPtrInput      `pulumi:"kind"`
	Location      pulumi.StringInput         `pulumi:"location"`
	Name          pulumi.StringPtrInput      `pulumi:"name"`
	Sku           SkuDescriptionPtrInput     `pulumi:"sku"`
	Tags          pulumi.StringMapInput      `pulumi:"tags"`
	Type          pulumi.StringPtrInput      `pulumi:"type"`
	WorkerCount   pulumi.IntPtrInput         `pulumi:"workerCount"`
	WorkerSize    pulumi.StringPtrInput      `pulumi:"workerSize"`
	WorkerSizeId  pulumi.IntPtrInput         `pulumi:"workerSizeId"`
}

func (WorkerPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerPool)(nil)).Elem()
}

func (i WorkerPoolArgs) ToWorkerPoolOutput() WorkerPoolOutput {
	return i.ToWorkerPoolOutputWithContext(context.Background())
}

func (i WorkerPoolArgs) ToWorkerPoolOutputWithContext(ctx context.Context) WorkerPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerPoolOutput)
}





type WorkerPoolArrayInput interface {
	pulumi.Input

	ToWorkerPoolArrayOutput() WorkerPoolArrayOutput
	ToWorkerPoolArrayOutputWithContext(context.Context) WorkerPoolArrayOutput
}

type WorkerPoolArray []WorkerPoolInput

func (WorkerPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkerPool)(nil)).Elem()
}

func (i WorkerPoolArray) ToWorkerPoolArrayOutput() WorkerPoolArrayOutput {
	return i.ToWorkerPoolArrayOutputWithContext(context.Background())
}

func (i WorkerPoolArray) ToWorkerPoolArrayOutputWithContext(ctx context.Context) WorkerPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerPoolArrayOutput)
}

type WorkerPoolOutput struct{ *pulumi.OutputState }

func (WorkerPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerPool)(nil)).Elem()
}

func (o WorkerPoolOutput) ToWorkerPoolOutput() WorkerPoolOutput {
	return o
}

func (o WorkerPoolOutput) ToWorkerPoolOutputWithContext(ctx context.Context) WorkerPoolOutput {
	return o
}

func (o WorkerPoolOutput) ComputeMode() ComputeModeOptionsPtrOutput {
	return o.ApplyT(func(v WorkerPool) *ComputeModeOptions { return v.ComputeMode }).(ComputeModeOptionsPtrOutput)
}

func (o WorkerPoolOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerPool) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o WorkerPoolOutput) InstanceNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkerPool) []string { return v.InstanceNames }).(pulumi.StringArrayOutput)
}

func (o WorkerPoolOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerPool) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WorkerPoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v WorkerPool) string { return v.Location }).(pulumi.StringOutput)
}

func (o WorkerPoolOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerPool) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WorkerPoolOutput) Sku() SkuDescriptionPtrOutput {
	return o.ApplyT(func(v WorkerPool) *SkuDescription { return v.Sku }).(SkuDescriptionPtrOutput)
}

func (o WorkerPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v WorkerPool) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o WorkerPoolOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerPool) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o WorkerPoolOutput) WorkerCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkerPool) *int { return v.WorkerCount }).(pulumi.IntPtrOutput)
}

func (o WorkerPoolOutput) WorkerSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerPool) *string { return v.WorkerSize }).(pulumi.StringPtrOutput)
}

func (o WorkerPoolOutput) WorkerSizeId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkerPool) *int { return v.WorkerSizeId }).(pulumi.IntPtrOutput)
}

type WorkerPoolArrayOutput struct{ *pulumi.OutputState }

func (WorkerPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkerPool)(nil)).Elem()
}

func (o WorkerPoolArrayOutput) ToWorkerPoolArrayOutput() WorkerPoolArrayOutput {
	return o
}

func (o WorkerPoolArrayOutput) ToWorkerPoolArrayOutputWithContext(ctx context.Context) WorkerPoolArrayOutput {
	return o
}

func (o WorkerPoolArrayOutput) Index(i pulumi.IntInput) WorkerPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkerPool {
		return vs[0].([]WorkerPool)[vs[1].(int)]
	}).(WorkerPoolOutput)
}

type WorkerPoolResponse struct {
	ComputeMode   *string                 `pulumi:"computeMode"`
	Id            *string                 `pulumi:"id"`
	InstanceNames []string                `pulumi:"instanceNames"`
	Kind          *string                 `pulumi:"kind"`
	Location      string                  `pulumi:"location"`
	Name          *string                 `pulumi:"name"`
	Sku           *SkuDescriptionResponse `pulumi:"sku"`
	Tags          map[string]string       `pulumi:"tags"`
	Type          *string                 `pulumi:"type"`
	WorkerCount   *int                    `pulumi:"workerCount"`
	WorkerSize    *string                 `pulumi:"workerSize"`
	WorkerSizeId  *int                    `pulumi:"workerSizeId"`
}

type WorkerPoolResponseOutput struct{ *pulumi.OutputState }

func (WorkerPoolResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerPoolResponse)(nil)).Elem()
}

func (o WorkerPoolResponseOutput) ToWorkerPoolResponseOutput() WorkerPoolResponseOutput {
	return o
}

func (o WorkerPoolResponseOutput) ToWorkerPoolResponseOutputWithContext(ctx context.Context) WorkerPoolResponseOutput {
	return o
}

func (o WorkerPoolResponseOutput) ComputeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerPoolResponse) *string { return v.ComputeMode }).(pulumi.StringPtrOutput)
}

func (o WorkerPoolResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerPoolResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o WorkerPoolResponseOutput) InstanceNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkerPoolResponse) []string { return v.InstanceNames }).(pulumi.StringArrayOutput)
}

func (o WorkerPoolResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerPoolResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WorkerPoolResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v WorkerPoolResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o WorkerPoolResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerPoolResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WorkerPoolResponseOutput) Sku() SkuDescriptionResponsePtrOutput {
	return o.ApplyT(func(v WorkerPoolResponse) *SkuDescriptionResponse { return v.Sku }).(SkuDescriptionResponsePtrOutput)
}

func (o WorkerPoolResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v WorkerPoolResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o WorkerPoolResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerPoolResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o WorkerPoolResponseOutput) WorkerCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkerPoolResponse) *int { return v.WorkerCount }).(pulumi.IntPtrOutput)
}

func (o WorkerPoolResponseOutput) WorkerSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerPoolResponse) *string { return v.WorkerSize }).(pulumi.StringPtrOutput)
}

func (o WorkerPoolResponseOutput) WorkerSizeId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkerPoolResponse) *int { return v.WorkerSizeId }).(pulumi.IntPtrOutput)
}

type WorkerPoolResponseArrayOutput struct{ *pulumi.OutputState }

func (WorkerPoolResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkerPoolResponse)(nil)).Elem()
}

func (o WorkerPoolResponseArrayOutput) ToWorkerPoolResponseArrayOutput() WorkerPoolResponseArrayOutput {
	return o
}

func (o WorkerPoolResponseArrayOutput) ToWorkerPoolResponseArrayOutputWithContext(ctx context.Context) WorkerPoolResponseArrayOutput {
	return o
}

func (o WorkerPoolResponseArrayOutput) Index(i pulumi.IntInput) WorkerPoolResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkerPoolResponse {
		return vs[0].([]WorkerPoolResponse)[vs[1].(int)]
	}).(WorkerPoolResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiDefinitionInfoOutput{})
	pulumi.RegisterOutputType(ApiDefinitionInfoPtrOutput{})
	pulumi.RegisterOutputType(ApiDefinitionInfoResponseOutput{})
	pulumi.RegisterOutputType(ApiDefinitionInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationLogsConfigOutput{})
	pulumi.RegisterOutputType(ApplicationLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(ApplicationLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(ApplicationLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoHealActionsOutput{})
	pulumi.RegisterOutputType(AutoHealActionsPtrOutput{})
	pulumi.RegisterOutputType(AutoHealActionsResponseOutput{})
	pulumi.RegisterOutputType(AutoHealActionsResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoHealCustomActionOutput{})
	pulumi.RegisterOutputType(AutoHealCustomActionPtrOutput{})
	pulumi.RegisterOutputType(AutoHealCustomActionResponseOutput{})
	pulumi.RegisterOutputType(AutoHealCustomActionResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoHealRulesOutput{})
	pulumi.RegisterOutputType(AutoHealRulesPtrOutput{})
	pulumi.RegisterOutputType(AutoHealRulesResponseOutput{})
	pulumi.RegisterOutputType(AutoHealRulesResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoHealTriggersOutput{})
	pulumi.RegisterOutputType(AutoHealTriggersPtrOutput{})
	pulumi.RegisterOutputType(AutoHealTriggersResponseOutput{})
	pulumi.RegisterOutputType(AutoHealTriggersResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageApplicationLogsConfigOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageApplicationLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageApplicationLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageApplicationLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageHttpLogsConfigOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageHttpLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageHttpLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageHttpLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureTableStorageApplicationLogsConfigOutput{})
	pulumi.RegisterOutputType(AzureTableStorageApplicationLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(AzureTableStorageApplicationLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(AzureTableStorageApplicationLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(BackupScheduleOutput{})
	pulumi.RegisterOutputType(BackupSchedulePtrOutput{})
	pulumi.RegisterOutputType(BackupScheduleResponseOutput{})
	pulumi.RegisterOutputType(BackupScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(CloningInfoOutput{})
	pulumi.RegisterOutputType(CloningInfoPtrOutput{})
	pulumi.RegisterOutputType(CloningInfoResponseOutput{})
	pulumi.RegisterOutputType(CloningInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnStringInfoOutput{})
	pulumi.RegisterOutputType(ConnStringInfoArrayOutput{})
	pulumi.RegisterOutputType(ConnStringInfoResponseOutput{})
	pulumi.RegisterOutputType(ConnStringInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ConnStringValueTypePairOutput{})
	pulumi.RegisterOutputType(ConnStringValueTypePairMapOutput{})
	pulumi.RegisterOutputType(ConnStringValueTypePairResponseOutput{})
	pulumi.RegisterOutputType(ConnStringValueTypePairResponseMapOutput{})
	pulumi.RegisterOutputType(CorsSettingsOutput{})
	pulumi.RegisterOutputType(CorsSettingsPtrOutput{})
	pulumi.RegisterOutputType(CorsSettingsResponseOutput{})
	pulumi.RegisterOutputType(CorsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(DatabaseBackupSettingOutput{})
	pulumi.RegisterOutputType(DatabaseBackupSettingArrayOutput{})
	pulumi.RegisterOutputType(DatabaseBackupSettingResponseOutput{})
	pulumi.RegisterOutputType(DatabaseBackupSettingResponseArrayOutput{})
	pulumi.RegisterOutputType(EnabledConfigOutput{})
	pulumi.RegisterOutputType(EnabledConfigPtrOutput{})
	pulumi.RegisterOutputType(EnabledConfigResponseOutput{})
	pulumi.RegisterOutputType(EnabledConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(ExperimentsOutput{})
	pulumi.RegisterOutputType(ExperimentsPtrOutput{})
	pulumi.RegisterOutputType(ExperimentsResponseOutput{})
	pulumi.RegisterOutputType(ExperimentsResponsePtrOutput{})
	pulumi.RegisterOutputType(FileSystemApplicationLogsConfigOutput{})
	pulumi.RegisterOutputType(FileSystemApplicationLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(FileSystemApplicationLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(FileSystemApplicationLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(FileSystemHttpLogsConfigOutput{})
	pulumi.RegisterOutputType(FileSystemHttpLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(FileSystemHttpLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(FileSystemHttpLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(HandlerMappingOutput{})
	pulumi.RegisterOutputType(HandlerMappingArrayOutput{})
	pulumi.RegisterOutputType(HandlerMappingResponseOutput{})
	pulumi.RegisterOutputType(HandlerMappingResponseArrayOutput{})
	pulumi.RegisterOutputType(HostNameSslStateOutput{})
	pulumi.RegisterOutputType(HostNameSslStateArrayOutput{})
	pulumi.RegisterOutputType(HostNameSslStateResponseOutput{})
	pulumi.RegisterOutputType(HostNameSslStateResponseArrayOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentProfileOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentProfilePtrOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentProfileResponseOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(HttpLogsConfigOutput{})
	pulumi.RegisterOutputType(HttpLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(HttpLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(HttpLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(IpSecurityRestrictionOutput{})
	pulumi.RegisterOutputType(IpSecurityRestrictionArrayOutput{})
	pulumi.RegisterOutputType(IpSecurityRestrictionResponseOutput{})
	pulumi.RegisterOutputType(IpSecurityRestrictionResponseArrayOutput{})
	pulumi.RegisterOutputType(NameValuePairOutput{})
	pulumi.RegisterOutputType(NameValuePairArrayOutput{})
	pulumi.RegisterOutputType(NameValuePairResponseOutput{})
	pulumi.RegisterOutputType(NameValuePairResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkAccessControlEntryOutput{})
	pulumi.RegisterOutputType(NetworkAccessControlEntryArrayOutput{})
	pulumi.RegisterOutputType(NetworkAccessControlEntryResponseOutput{})
	pulumi.RegisterOutputType(NetworkAccessControlEntryResponseArrayOutput{})
	pulumi.RegisterOutputType(RampUpRuleOutput{})
	pulumi.RegisterOutputType(RampUpRuleArrayOutput{})
	pulumi.RegisterOutputType(RampUpRuleResponseOutput{})
	pulumi.RegisterOutputType(RampUpRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(RequestsBasedTriggerOutput{})
	pulumi.RegisterOutputType(RequestsBasedTriggerPtrOutput{})
	pulumi.RegisterOutputType(RequestsBasedTriggerResponseOutput{})
	pulumi.RegisterOutputType(RequestsBasedTriggerResponsePtrOutput{})
	pulumi.RegisterOutputType(SiteConfigOutput{})
	pulumi.RegisterOutputType(SiteConfigPtrOutput{})
	pulumi.RegisterOutputType(SiteConfigResponseOutput{})
	pulumi.RegisterOutputType(SiteConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(SiteLimitsOutput{})
	pulumi.RegisterOutputType(SiteLimitsPtrOutput{})
	pulumi.RegisterOutputType(SiteLimitsResponseOutput{})
	pulumi.RegisterOutputType(SiteLimitsResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuDescriptionOutput{})
	pulumi.RegisterOutputType(SkuDescriptionPtrOutput{})
	pulumi.RegisterOutputType(SkuDescriptionResponseOutput{})
	pulumi.RegisterOutputType(SkuDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerPtrOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerResponseOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerResponsePtrOutput{})
	pulumi.RegisterOutputType(StampCapacityOutput{})
	pulumi.RegisterOutputType(StampCapacityArrayOutput{})
	pulumi.RegisterOutputType(StampCapacityResponseOutput{})
	pulumi.RegisterOutputType(StampCapacityResponseArrayOutput{})
	pulumi.RegisterOutputType(StatusCodesBasedTriggerOutput{})
	pulumi.RegisterOutputType(StatusCodesBasedTriggerArrayOutput{})
	pulumi.RegisterOutputType(StatusCodesBasedTriggerResponseOutput{})
	pulumi.RegisterOutputType(StatusCodesBasedTriggerResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualApplicationOutput{})
	pulumi.RegisterOutputType(VirtualApplicationArrayOutput{})
	pulumi.RegisterOutputType(VirtualApplicationResponseOutput{})
	pulumi.RegisterOutputType(VirtualApplicationResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualDirectoryOutput{})
	pulumi.RegisterOutputType(VirtualDirectoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualDirectoryResponseOutput{})
	pulumi.RegisterOutputType(VirtualDirectoryResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualIPMappingOutput{})
	pulumi.RegisterOutputType(VirtualIPMappingArrayOutput{})
	pulumi.RegisterOutputType(VirtualIPMappingResponseOutput{})
	pulumi.RegisterOutputType(VirtualIPMappingResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfileOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(VnetRouteOutput{})
	pulumi.RegisterOutputType(VnetRouteArrayOutput{})
	pulumi.RegisterOutputType(VnetRouteResponseOutput{})
	pulumi.RegisterOutputType(VnetRouteResponseArrayOutput{})
	pulumi.RegisterOutputType(WorkerPoolOutput{})
	pulumi.RegisterOutputType(WorkerPoolArrayOutput{})
	pulumi.RegisterOutputType(WorkerPoolResponseOutput{})
	pulumi.RegisterOutputType(WorkerPoolResponseArrayOutput{})
}
