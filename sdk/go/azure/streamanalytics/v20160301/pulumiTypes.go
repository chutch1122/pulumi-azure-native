


package v20160301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AvroSerialization struct {
	Type string `pulumi:"type"`
}

type AvroSerializationResponse struct {
	Type string `pulumi:"type"`
}

type AzureDataLakeStoreOutputDataSource struct {
	AccountName            *string `pulumi:"accountName"`
	DateFormat             *string `pulumi:"dateFormat"`
	FilePathPrefix         *string `pulumi:"filePathPrefix"`
	RefreshToken           *string `pulumi:"refreshToken"`
	TenantId               *string `pulumi:"tenantId"`
	TimeFormat             *string `pulumi:"timeFormat"`
	TokenUserDisplayName   *string `pulumi:"tokenUserDisplayName"`
	TokenUserPrincipalName *string `pulumi:"tokenUserPrincipalName"`
	Type                   string  `pulumi:"type"`
}

type AzureDataLakeStoreOutputDataSourceResponse struct {
	AccountName            *string `pulumi:"accountName"`
	DateFormat             *string `pulumi:"dateFormat"`
	FilePathPrefix         *string `pulumi:"filePathPrefix"`
	RefreshToken           *string `pulumi:"refreshToken"`
	TenantId               *string `pulumi:"tenantId"`
	TimeFormat             *string `pulumi:"timeFormat"`
	TokenUserDisplayName   *string `pulumi:"tokenUserDisplayName"`
	TokenUserPrincipalName *string `pulumi:"tokenUserPrincipalName"`
	Type                   string  `pulumi:"type"`
}

type AzureMachineLearningWebServiceFunctionBinding struct {
	ApiKey    *string                                      `pulumi:"apiKey"`
	BatchSize *int                                         `pulumi:"batchSize"`
	Endpoint  *string                                      `pulumi:"endpoint"`
	Inputs    *AzureMachineLearningWebServiceInputs        `pulumi:"inputs"`
	Outputs   []AzureMachineLearningWebServiceOutputColumn `pulumi:"outputs"`
	Type      string                                       `pulumi:"type"`
}

type AzureMachineLearningWebServiceFunctionBindingResponse struct {
	ApiKey    *string                                              `pulumi:"apiKey"`
	BatchSize *int                                                 `pulumi:"batchSize"`
	Endpoint  *string                                              `pulumi:"endpoint"`
	Inputs    *AzureMachineLearningWebServiceInputsResponse        `pulumi:"inputs"`
	Outputs   []AzureMachineLearningWebServiceOutputColumnResponse `pulumi:"outputs"`
	Type      string                                               `pulumi:"type"`
}

type AzureMachineLearningWebServiceInputColumn struct {
	DataType *string `pulumi:"dataType"`
	MapTo    *int    `pulumi:"mapTo"`
	Name     *string `pulumi:"name"`
}

type AzureMachineLearningWebServiceInputColumnResponse struct {
	DataType *string `pulumi:"dataType"`
	MapTo    *int    `pulumi:"mapTo"`
	Name     *string `pulumi:"name"`
}

type AzureMachineLearningWebServiceInputs struct {
	ColumnNames []AzureMachineLearningWebServiceInputColumn `pulumi:"columnNames"`
	Name        *string                                     `pulumi:"name"`
}

type AzureMachineLearningWebServiceInputsResponse struct {
	ColumnNames []AzureMachineLearningWebServiceInputColumnResponse `pulumi:"columnNames"`
	Name        *string                                             `pulumi:"name"`
}

type AzureMachineLearningWebServiceOutputColumn struct {
	DataType *string `pulumi:"dataType"`
	Name     *string `pulumi:"name"`
}

type AzureMachineLearningWebServiceOutputColumnResponse struct {
	DataType *string `pulumi:"dataType"`
	Name     *string `pulumi:"name"`
}

type AzureSqlDatabaseOutputDataSource struct {
	Database *string `pulumi:"database"`
	Password *string `pulumi:"password"`
	Server   *string `pulumi:"server"`
	Table    *string `pulumi:"table"`
	Type     string  `pulumi:"type"`
	User     *string `pulumi:"user"`
}

type AzureSqlDatabaseOutputDataSourceResponse struct {
	Database *string `pulumi:"database"`
	Password *string `pulumi:"password"`
	Server   *string `pulumi:"server"`
	Table    *string `pulumi:"table"`
	Type     string  `pulumi:"type"`
	User     *string `pulumi:"user"`
}

type AzureTableOutputDataSource struct {
	AccountKey      *string  `pulumi:"accountKey"`
	AccountName     *string  `pulumi:"accountName"`
	BatchSize       *int     `pulumi:"batchSize"`
	ColumnsToRemove []string `pulumi:"columnsToRemove"`
	PartitionKey    *string  `pulumi:"partitionKey"`
	RowKey          *string  `pulumi:"rowKey"`
	Table           *string  `pulumi:"table"`
	Type            string   `pulumi:"type"`
}

type AzureTableOutputDataSourceResponse struct {
	AccountKey      *string  `pulumi:"accountKey"`
	AccountName     *string  `pulumi:"accountName"`
	BatchSize       *int     `pulumi:"batchSize"`
	ColumnsToRemove []string `pulumi:"columnsToRemove"`
	PartitionKey    *string  `pulumi:"partitionKey"`
	RowKey          *string  `pulumi:"rowKey"`
	Table           *string  `pulumi:"table"`
	Type            string   `pulumi:"type"`
}

type BlobOutputDataSource struct {
	Container       *string          `pulumi:"container"`
	DateFormat      *string          `pulumi:"dateFormat"`
	PathPattern     *string          `pulumi:"pathPattern"`
	StorageAccounts []StorageAccount `pulumi:"storageAccounts"`
	TimeFormat      *string          `pulumi:"timeFormat"`
	Type            string           `pulumi:"type"`
}

type BlobOutputDataSourceResponse struct {
	Container       *string                  `pulumi:"container"`
	DateFormat      *string                  `pulumi:"dateFormat"`
	PathPattern     *string                  `pulumi:"pathPattern"`
	StorageAccounts []StorageAccountResponse `pulumi:"storageAccounts"`
	TimeFormat      *string                  `pulumi:"timeFormat"`
	Type            string                   `pulumi:"type"`
}

type BlobReferenceInputDataSource struct {
	Container       *string          `pulumi:"container"`
	DateFormat      *string          `pulumi:"dateFormat"`
	PathPattern     *string          `pulumi:"pathPattern"`
	StorageAccounts []StorageAccount `pulumi:"storageAccounts"`
	TimeFormat      *string          `pulumi:"timeFormat"`
	Type            string           `pulumi:"type"`
}

type BlobReferenceInputDataSourceResponse struct {
	Container       *string                  `pulumi:"container"`
	DateFormat      *string                  `pulumi:"dateFormat"`
	PathPattern     *string                  `pulumi:"pathPattern"`
	StorageAccounts []StorageAccountResponse `pulumi:"storageAccounts"`
	TimeFormat      *string                  `pulumi:"timeFormat"`
	Type            string                   `pulumi:"type"`
}

type BlobStreamInputDataSource struct {
	Container            *string          `pulumi:"container"`
	DateFormat           *string          `pulumi:"dateFormat"`
	PathPattern          *string          `pulumi:"pathPattern"`
	SourcePartitionCount *int             `pulumi:"sourcePartitionCount"`
	StorageAccounts      []StorageAccount `pulumi:"storageAccounts"`
	TimeFormat           *string          `pulumi:"timeFormat"`
	Type                 string           `pulumi:"type"`
}

type BlobStreamInputDataSourceResponse struct {
	Container            *string                  `pulumi:"container"`
	DateFormat           *string                  `pulumi:"dateFormat"`
	PathPattern          *string                  `pulumi:"pathPattern"`
	SourcePartitionCount *int                     `pulumi:"sourcePartitionCount"`
	StorageAccounts      []StorageAccountResponse `pulumi:"storageAccounts"`
	TimeFormat           *string                  `pulumi:"timeFormat"`
	Type                 string                   `pulumi:"type"`
}

type CsvSerialization struct {
	Encoding       *string `pulumi:"encoding"`
	FieldDelimiter *string `pulumi:"fieldDelimiter"`
	Type           string  `pulumi:"type"`
}

type CsvSerializationResponse struct {
	Encoding       *string `pulumi:"encoding"`
	FieldDelimiter *string `pulumi:"fieldDelimiter"`
	Type           string  `pulumi:"type"`
}

type DiagnosticConditionResponse struct {
	Code    string `pulumi:"code"`
	Message string `pulumi:"message"`
	Since   string `pulumi:"since"`
}

type DiagnosticConditionResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticConditionResponse)(nil)).Elem()
}

func (o DiagnosticConditionResponseOutput) ToDiagnosticConditionResponseOutput() DiagnosticConditionResponseOutput {
	return o
}

func (o DiagnosticConditionResponseOutput) ToDiagnosticConditionResponseOutputWithContext(ctx context.Context) DiagnosticConditionResponseOutput {
	return o
}

func (o DiagnosticConditionResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticConditionResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o DiagnosticConditionResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticConditionResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o DiagnosticConditionResponseOutput) Since() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticConditionResponse) string { return v.Since }).(pulumi.StringOutput)
}

type DiagnosticConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (DiagnosticConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiagnosticConditionResponse)(nil)).Elem()
}

func (o DiagnosticConditionResponseArrayOutput) ToDiagnosticConditionResponseArrayOutput() DiagnosticConditionResponseArrayOutput {
	return o
}

func (o DiagnosticConditionResponseArrayOutput) ToDiagnosticConditionResponseArrayOutputWithContext(ctx context.Context) DiagnosticConditionResponseArrayOutput {
	return o
}

func (o DiagnosticConditionResponseArrayOutput) Index(i pulumi.IntInput) DiagnosticConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiagnosticConditionResponse {
		return vs[0].([]DiagnosticConditionResponse)[vs[1].(int)]
	}).(DiagnosticConditionResponseOutput)
}

type DiagnosticsResponse struct {
	Conditions []DiagnosticConditionResponse `pulumi:"conditions"`
}

type DiagnosticsResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsResponse)(nil)).Elem()
}

func (o DiagnosticsResponseOutput) ToDiagnosticsResponseOutput() DiagnosticsResponseOutput {
	return o
}

func (o DiagnosticsResponseOutput) ToDiagnosticsResponseOutputWithContext(ctx context.Context) DiagnosticsResponseOutput {
	return o
}

func (o DiagnosticsResponseOutput) Conditions() DiagnosticConditionResponseArrayOutput {
	return o.ApplyT(func(v DiagnosticsResponse) []DiagnosticConditionResponse { return v.Conditions }).(DiagnosticConditionResponseArrayOutput)
}

type DocumentDbOutputDataSource struct {
	AccountId             *string `pulumi:"accountId"`
	AccountKey            *string `pulumi:"accountKey"`
	CollectionNamePattern *string `pulumi:"collectionNamePattern"`
	Database              *string `pulumi:"database"`
	DocumentId            *string `pulumi:"documentId"`
	PartitionKey          *string `pulumi:"partitionKey"`
	Type                  string  `pulumi:"type"`
}

type DocumentDbOutputDataSourceResponse struct {
	AccountId             *string `pulumi:"accountId"`
	AccountKey            *string `pulumi:"accountKey"`
	CollectionNamePattern *string `pulumi:"collectionNamePattern"`
	Database              *string `pulumi:"database"`
	DocumentId            *string `pulumi:"documentId"`
	PartitionKey          *string `pulumi:"partitionKey"`
	Type                  string  `pulumi:"type"`
}

type EventHubOutputDataSource struct {
	EventHubName           *string `pulumi:"eventHubName"`
	PartitionKey           *string `pulumi:"partitionKey"`
	ServiceBusNamespace    *string `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	Type                   string  `pulumi:"type"`
}

type EventHubOutputDataSourceResponse struct {
	EventHubName           *string `pulumi:"eventHubName"`
	PartitionKey           *string `pulumi:"partitionKey"`
	ServiceBusNamespace    *string `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	Type                   string  `pulumi:"type"`
}

type EventHubStreamInputDataSource struct {
	ConsumerGroupName      *string `pulumi:"consumerGroupName"`
	EventHubName           *string `pulumi:"eventHubName"`
	ServiceBusNamespace    *string `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	Type                   string  `pulumi:"type"`
}

type EventHubStreamInputDataSourceResponse struct {
	ConsumerGroupName      *string `pulumi:"consumerGroupName"`
	EventHubName           *string `pulumi:"eventHubName"`
	ServiceBusNamespace    *string `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	Type                   string  `pulumi:"type"`
}

type FunctionType struct {
	Name       *string                   `pulumi:"name"`
	Properties *ScalarFunctionProperties `pulumi:"properties"`
}





type FunctionTypeInput interface {
	pulumi.Input

	ToFunctionTypeOutput() FunctionTypeOutput
	ToFunctionTypeOutputWithContext(context.Context) FunctionTypeOutput
}

type FunctionTypeArgs struct {
	Name       pulumi.StringPtrInput            `pulumi:"name"`
	Properties ScalarFunctionPropertiesPtrInput `pulumi:"properties"`
}

func (FunctionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionType)(nil)).Elem()
}

func (i FunctionTypeArgs) ToFunctionTypeOutput() FunctionTypeOutput {
	return i.ToFunctionTypeOutputWithContext(context.Background())
}

func (i FunctionTypeArgs) ToFunctionTypeOutputWithContext(ctx context.Context) FunctionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionTypeOutput)
}





type FunctionTypeArrayInput interface {
	pulumi.Input

	ToFunctionTypeArrayOutput() FunctionTypeArrayOutput
	ToFunctionTypeArrayOutputWithContext(context.Context) FunctionTypeArrayOutput
}

type FunctionTypeArray []FunctionTypeInput

func (FunctionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionType)(nil)).Elem()
}

func (i FunctionTypeArray) ToFunctionTypeArrayOutput() FunctionTypeArrayOutput {
	return i.ToFunctionTypeArrayOutputWithContext(context.Background())
}

func (i FunctionTypeArray) ToFunctionTypeArrayOutputWithContext(ctx context.Context) FunctionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionTypeArrayOutput)
}

type FunctionTypeOutput struct{ *pulumi.OutputState }

func (FunctionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionType)(nil)).Elem()
}

func (o FunctionTypeOutput) ToFunctionTypeOutput() FunctionTypeOutput {
	return o
}

func (o FunctionTypeOutput) ToFunctionTypeOutputWithContext(ctx context.Context) FunctionTypeOutput {
	return o
}

func (o FunctionTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FunctionType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FunctionTypeOutput) Properties() ScalarFunctionPropertiesPtrOutput {
	return o.ApplyT(func(v FunctionType) *ScalarFunctionProperties { return v.Properties }).(ScalarFunctionPropertiesPtrOutput)
}

type FunctionTypeArrayOutput struct{ *pulumi.OutputState }

func (FunctionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionType)(nil)).Elem()
}

func (o FunctionTypeArrayOutput) ToFunctionTypeArrayOutput() FunctionTypeArrayOutput {
	return o
}

func (o FunctionTypeArrayOutput) ToFunctionTypeArrayOutputWithContext(ctx context.Context) FunctionTypeArrayOutput {
	return o
}

func (o FunctionTypeArrayOutput) Index(i pulumi.IntInput) FunctionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FunctionType {
		return vs[0].([]FunctionType)[vs[1].(int)]
	}).(FunctionTypeOutput)
}

type FunctionInputType struct {
	DataType                 *string `pulumi:"dataType"`
	IsConfigurationParameter *bool   `pulumi:"isConfigurationParameter"`
}





type FunctionInputTypeInput interface {
	pulumi.Input

	ToFunctionInputTypeOutput() FunctionInputTypeOutput
	ToFunctionInputTypeOutputWithContext(context.Context) FunctionInputTypeOutput
}

type FunctionInputTypeArgs struct {
	DataType                 pulumi.StringPtrInput `pulumi:"dataType"`
	IsConfigurationParameter pulumi.BoolPtrInput   `pulumi:"isConfigurationParameter"`
}

func (FunctionInputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionInputType)(nil)).Elem()
}

func (i FunctionInputTypeArgs) ToFunctionInputTypeOutput() FunctionInputTypeOutput {
	return i.ToFunctionInputTypeOutputWithContext(context.Background())
}

func (i FunctionInputTypeArgs) ToFunctionInputTypeOutputWithContext(ctx context.Context) FunctionInputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionInputTypeOutput)
}





type FunctionInputTypeArrayInput interface {
	pulumi.Input

	ToFunctionInputTypeArrayOutput() FunctionInputTypeArrayOutput
	ToFunctionInputTypeArrayOutputWithContext(context.Context) FunctionInputTypeArrayOutput
}

type FunctionInputTypeArray []FunctionInputTypeInput

func (FunctionInputTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionInputType)(nil)).Elem()
}

func (i FunctionInputTypeArray) ToFunctionInputTypeArrayOutput() FunctionInputTypeArrayOutput {
	return i.ToFunctionInputTypeArrayOutputWithContext(context.Background())
}

func (i FunctionInputTypeArray) ToFunctionInputTypeArrayOutputWithContext(ctx context.Context) FunctionInputTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionInputTypeArrayOutput)
}

type FunctionInputTypeOutput struct{ *pulumi.OutputState }

func (FunctionInputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionInputType)(nil)).Elem()
}

func (o FunctionInputTypeOutput) ToFunctionInputTypeOutput() FunctionInputTypeOutput {
	return o
}

func (o FunctionInputTypeOutput) ToFunctionInputTypeOutputWithContext(ctx context.Context) FunctionInputTypeOutput {
	return o
}

func (o FunctionInputTypeOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FunctionInputType) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o FunctionInputTypeOutput) IsConfigurationParameter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FunctionInputType) *bool { return v.IsConfigurationParameter }).(pulumi.BoolPtrOutput)
}

type FunctionInputTypeArrayOutput struct{ *pulumi.OutputState }

func (FunctionInputTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionInputType)(nil)).Elem()
}

func (o FunctionInputTypeArrayOutput) ToFunctionInputTypeArrayOutput() FunctionInputTypeArrayOutput {
	return o
}

func (o FunctionInputTypeArrayOutput) ToFunctionInputTypeArrayOutputWithContext(ctx context.Context) FunctionInputTypeArrayOutput {
	return o
}

func (o FunctionInputTypeArrayOutput) Index(i pulumi.IntInput) FunctionInputTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FunctionInputType {
		return vs[0].([]FunctionInputType)[vs[1].(int)]
	}).(FunctionInputTypeOutput)
}

type FunctionInputResponse struct {
	DataType                 *string `pulumi:"dataType"`
	IsConfigurationParameter *bool   `pulumi:"isConfigurationParameter"`
}

type FunctionInputResponseOutput struct{ *pulumi.OutputState }

func (FunctionInputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionInputResponse)(nil)).Elem()
}

func (o FunctionInputResponseOutput) ToFunctionInputResponseOutput() FunctionInputResponseOutput {
	return o
}

func (o FunctionInputResponseOutput) ToFunctionInputResponseOutputWithContext(ctx context.Context) FunctionInputResponseOutput {
	return o
}

func (o FunctionInputResponseOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FunctionInputResponse) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o FunctionInputResponseOutput) IsConfigurationParameter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FunctionInputResponse) *bool { return v.IsConfigurationParameter }).(pulumi.BoolPtrOutput)
}

type FunctionInputResponseArrayOutput struct{ *pulumi.OutputState }

func (FunctionInputResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionInputResponse)(nil)).Elem()
}

func (o FunctionInputResponseArrayOutput) ToFunctionInputResponseArrayOutput() FunctionInputResponseArrayOutput {
	return o
}

func (o FunctionInputResponseArrayOutput) ToFunctionInputResponseArrayOutputWithContext(ctx context.Context) FunctionInputResponseArrayOutput {
	return o
}

func (o FunctionInputResponseArrayOutput) Index(i pulumi.IntInput) FunctionInputResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FunctionInputResponse {
		return vs[0].([]FunctionInputResponse)[vs[1].(int)]
	}).(FunctionInputResponseOutput)
}

type FunctionOutputType struct {
	DataType *string `pulumi:"dataType"`
}





type FunctionOutputTypeInput interface {
	pulumi.Input

	ToFunctionOutputTypeOutput() FunctionOutputTypeOutput
	ToFunctionOutputTypeOutputWithContext(context.Context) FunctionOutputTypeOutput
}

type FunctionOutputTypeArgs struct {
	DataType pulumi.StringPtrInput `pulumi:"dataType"`
}

func (FunctionOutputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionOutputType)(nil)).Elem()
}

func (i FunctionOutputTypeArgs) ToFunctionOutputTypeOutput() FunctionOutputTypeOutput {
	return i.ToFunctionOutputTypeOutputWithContext(context.Background())
}

func (i FunctionOutputTypeArgs) ToFunctionOutputTypeOutputWithContext(ctx context.Context) FunctionOutputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutputTypeOutput)
}

func (i FunctionOutputTypeArgs) ToFunctionOutputTypePtrOutput() FunctionOutputTypePtrOutput {
	return i.ToFunctionOutputTypePtrOutputWithContext(context.Background())
}

func (i FunctionOutputTypeArgs) ToFunctionOutputTypePtrOutputWithContext(ctx context.Context) FunctionOutputTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutputTypeOutput).ToFunctionOutputTypePtrOutputWithContext(ctx)
}









type FunctionOutputTypePtrInput interface {
	pulumi.Input

	ToFunctionOutputTypePtrOutput() FunctionOutputTypePtrOutput
	ToFunctionOutputTypePtrOutputWithContext(context.Context) FunctionOutputTypePtrOutput
}

type functionOutputTypePtrType FunctionOutputTypeArgs

func FunctionOutputTypePtr(v *FunctionOutputTypeArgs) FunctionOutputTypePtrInput {
	return (*functionOutputTypePtrType)(v)
}

func (*functionOutputTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionOutputType)(nil)).Elem()
}

func (i *functionOutputTypePtrType) ToFunctionOutputTypePtrOutput() FunctionOutputTypePtrOutput {
	return i.ToFunctionOutputTypePtrOutputWithContext(context.Background())
}

func (i *functionOutputTypePtrType) ToFunctionOutputTypePtrOutputWithContext(ctx context.Context) FunctionOutputTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutputTypePtrOutput)
}

type FunctionOutputTypeOutput struct{ *pulumi.OutputState }

func (FunctionOutputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionOutputType)(nil)).Elem()
}

func (o FunctionOutputTypeOutput) ToFunctionOutputTypeOutput() FunctionOutputTypeOutput {
	return o
}

func (o FunctionOutputTypeOutput) ToFunctionOutputTypeOutputWithContext(ctx context.Context) FunctionOutputTypeOutput {
	return o
}

func (o FunctionOutputTypeOutput) ToFunctionOutputTypePtrOutput() FunctionOutputTypePtrOutput {
	return o.ToFunctionOutputTypePtrOutputWithContext(context.Background())
}

func (o FunctionOutputTypeOutput) ToFunctionOutputTypePtrOutputWithContext(ctx context.Context) FunctionOutputTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FunctionOutputType) *FunctionOutputType {
		return &v
	}).(FunctionOutputTypePtrOutput)
}

func (o FunctionOutputTypeOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FunctionOutputType) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

type FunctionOutputTypePtrOutput struct{ *pulumi.OutputState }

func (FunctionOutputTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionOutputType)(nil)).Elem()
}

func (o FunctionOutputTypePtrOutput) ToFunctionOutputTypePtrOutput() FunctionOutputTypePtrOutput {
	return o
}

func (o FunctionOutputTypePtrOutput) ToFunctionOutputTypePtrOutputWithContext(ctx context.Context) FunctionOutputTypePtrOutput {
	return o
}

func (o FunctionOutputTypePtrOutput) Elem() FunctionOutputTypeOutput {
	return o.ApplyT(func(v *FunctionOutputType) FunctionOutputType {
		if v != nil {
			return *v
		}
		var ret FunctionOutputType
		return ret
	}).(FunctionOutputTypeOutput)
}

func (o FunctionOutputTypePtrOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FunctionOutputType) *string {
		if v == nil {
			return nil
		}
		return v.DataType
	}).(pulumi.StringPtrOutput)
}

type FunctionOutputResponse struct {
	DataType *string `pulumi:"dataType"`
}

type FunctionOutputResponseOutput struct{ *pulumi.OutputState }

func (FunctionOutputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionOutputResponse)(nil)).Elem()
}

func (o FunctionOutputResponseOutput) ToFunctionOutputResponseOutput() FunctionOutputResponseOutput {
	return o
}

func (o FunctionOutputResponseOutput) ToFunctionOutputResponseOutputWithContext(ctx context.Context) FunctionOutputResponseOutput {
	return o
}

func (o FunctionOutputResponseOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FunctionOutputResponse) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

type FunctionOutputResponsePtrOutput struct{ *pulumi.OutputState }

func (FunctionOutputResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionOutputResponse)(nil)).Elem()
}

func (o FunctionOutputResponsePtrOutput) ToFunctionOutputResponsePtrOutput() FunctionOutputResponsePtrOutput {
	return o
}

func (o FunctionOutputResponsePtrOutput) ToFunctionOutputResponsePtrOutputWithContext(ctx context.Context) FunctionOutputResponsePtrOutput {
	return o
}

func (o FunctionOutputResponsePtrOutput) Elem() FunctionOutputResponseOutput {
	return o.ApplyT(func(v *FunctionOutputResponse) FunctionOutputResponse {
		if v != nil {
			return *v
		}
		var ret FunctionOutputResponse
		return ret
	}).(FunctionOutputResponseOutput)
}

func (o FunctionOutputResponsePtrOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FunctionOutputResponse) *string {
		if v == nil {
			return nil
		}
		return v.DataType
	}).(pulumi.StringPtrOutput)
}

type FunctionResponse struct {
	Id         string                            `pulumi:"id"`
	Name       *string                           `pulumi:"name"`
	Properties *ScalarFunctionPropertiesResponse `pulumi:"properties"`
	Type       string                            `pulumi:"type"`
}

type FunctionResponseOutput struct{ *pulumi.OutputState }

func (FunctionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionResponse)(nil)).Elem()
}

func (o FunctionResponseOutput) ToFunctionResponseOutput() FunctionResponseOutput {
	return o
}

func (o FunctionResponseOutput) ToFunctionResponseOutputWithContext(ctx context.Context) FunctionResponseOutput {
	return o
}

func (o FunctionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v FunctionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o FunctionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FunctionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FunctionResponseOutput) Properties() ScalarFunctionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v FunctionResponse) *ScalarFunctionPropertiesResponse { return v.Properties }).(ScalarFunctionPropertiesResponsePtrOutput)
}

func (o FunctionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v FunctionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type FunctionResponseArrayOutput struct{ *pulumi.OutputState }

func (FunctionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionResponse)(nil)).Elem()
}

func (o FunctionResponseArrayOutput) ToFunctionResponseArrayOutput() FunctionResponseArrayOutput {
	return o
}

func (o FunctionResponseArrayOutput) ToFunctionResponseArrayOutputWithContext(ctx context.Context) FunctionResponseArrayOutput {
	return o
}

func (o FunctionResponseArrayOutput) Index(i pulumi.IntInput) FunctionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FunctionResponse {
		return vs[0].([]FunctionResponse)[vs[1].(int)]
	}).(FunctionResponseOutput)
}

type InputType struct {
	Name       *string     `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
}





type InputTypeInput interface {
	pulumi.Input

	ToInputTypeOutput() InputTypeOutput
	ToInputTypeOutputWithContext(context.Context) InputTypeOutput
}

type InputTypeArgs struct {
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Properties pulumi.Input          `pulumi:"properties"`
}

func (InputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InputType)(nil)).Elem()
}

func (i InputTypeArgs) ToInputTypeOutput() InputTypeOutput {
	return i.ToInputTypeOutputWithContext(context.Background())
}

func (i InputTypeArgs) ToInputTypeOutputWithContext(ctx context.Context) InputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputTypeOutput)
}





type InputTypeArrayInput interface {
	pulumi.Input

	ToInputTypeArrayOutput() InputTypeArrayOutput
	ToInputTypeArrayOutputWithContext(context.Context) InputTypeArrayOutput
}

type InputTypeArray []InputTypeInput

func (InputTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InputType)(nil)).Elem()
}

func (i InputTypeArray) ToInputTypeArrayOutput() InputTypeArrayOutput {
	return i.ToInputTypeArrayOutputWithContext(context.Background())
}

func (i InputTypeArray) ToInputTypeArrayOutputWithContext(ctx context.Context) InputTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputTypeArrayOutput)
}

type InputTypeOutput struct{ *pulumi.OutputState }

func (InputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputType)(nil)).Elem()
}

func (o InputTypeOutput) ToInputTypeOutput() InputTypeOutput {
	return o
}

func (o InputTypeOutput) ToInputTypeOutputWithContext(ctx context.Context) InputTypeOutput {
	return o
}

func (o InputTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InputType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o InputTypeOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v InputType) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

type InputTypeArrayOutput struct{ *pulumi.OutputState }

func (InputTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InputType)(nil)).Elem()
}

func (o InputTypeArrayOutput) ToInputTypeArrayOutput() InputTypeArrayOutput {
	return o
}

func (o InputTypeArrayOutput) ToInputTypeArrayOutputWithContext(ctx context.Context) InputTypeArrayOutput {
	return o
}

func (o InputTypeArrayOutput) Index(i pulumi.IntInput) InputTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InputType {
		return vs[0].([]InputType)[vs[1].(int)]
	}).(InputTypeOutput)
}

type InputResponse struct {
	Id         string      `pulumi:"id"`
	Name       *string     `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}

type InputResponseOutput struct{ *pulumi.OutputState }

func (InputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputResponse)(nil)).Elem()
}

func (o InputResponseOutput) ToInputResponseOutput() InputResponseOutput {
	return o
}

func (o InputResponseOutput) ToInputResponseOutputWithContext(ctx context.Context) InputResponseOutput {
	return o
}

func (o InputResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v InputResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o InputResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InputResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o InputResponseOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v InputResponse) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o InputResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v InputResponse) string { return v.Type }).(pulumi.StringOutput)
}

type InputResponseArrayOutput struct{ *pulumi.OutputState }

func (InputResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InputResponse)(nil)).Elem()
}

func (o InputResponseArrayOutput) ToInputResponseArrayOutput() InputResponseArrayOutput {
	return o
}

func (o InputResponseArrayOutput) ToInputResponseArrayOutputWithContext(ctx context.Context) InputResponseArrayOutput {
	return o
}

func (o InputResponseArrayOutput) Index(i pulumi.IntInput) InputResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InputResponse {
		return vs[0].([]InputResponse)[vs[1].(int)]
	}).(InputResponseOutput)
}

type IoTHubStreamInputDataSource struct {
	ConsumerGroupName      *string `pulumi:"consumerGroupName"`
	Endpoint               *string `pulumi:"endpoint"`
	IotHubNamespace        *string `pulumi:"iotHubNamespace"`
	SharedAccessPolicyKey  *string `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	Type                   string  `pulumi:"type"`
}

type IoTHubStreamInputDataSourceResponse struct {
	ConsumerGroupName      *string `pulumi:"consumerGroupName"`
	Endpoint               *string `pulumi:"endpoint"`
	IotHubNamespace        *string `pulumi:"iotHubNamespace"`
	SharedAccessPolicyKey  *string `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	Type                   string  `pulumi:"type"`
}

type JavaScriptFunctionBinding struct {
	Script *string `pulumi:"script"`
	Type   string  `pulumi:"type"`
}

type JavaScriptFunctionBindingResponse struct {
	Script *string `pulumi:"script"`
	Type   string  `pulumi:"type"`
}

type JsonSerialization struct {
	Encoding *string `pulumi:"encoding"`
	Format   *string `pulumi:"format"`
	Type     string  `pulumi:"type"`
}

type JsonSerializationResponse struct {
	Encoding *string `pulumi:"encoding"`
	Format   *string `pulumi:"format"`
	Type     string  `pulumi:"type"`
}

type OutputType struct {
	Datasource    interface{} `pulumi:"datasource"`
	Name          *string     `pulumi:"name"`
	Serialization interface{} `pulumi:"serialization"`
}





type OutputTypeInput interface {
	pulumi.Input

	ToOutputTypeOutput() OutputTypeOutput
	ToOutputTypeOutputWithContext(context.Context) OutputTypeOutput
}

type OutputTypeArgs struct {
	Datasource    pulumi.Input          `pulumi:"datasource"`
	Name          pulumi.StringPtrInput `pulumi:"name"`
	Serialization pulumi.Input          `pulumi:"serialization"`
}

func (OutputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputType)(nil)).Elem()
}

func (i OutputTypeArgs) ToOutputTypeOutput() OutputTypeOutput {
	return i.ToOutputTypeOutputWithContext(context.Background())
}

func (i OutputTypeArgs) ToOutputTypeOutputWithContext(ctx context.Context) OutputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputTypeOutput)
}





type OutputTypeArrayInput interface {
	pulumi.Input

	ToOutputTypeArrayOutput() OutputTypeArrayOutput
	ToOutputTypeArrayOutputWithContext(context.Context) OutputTypeArrayOutput
}

type OutputTypeArray []OutputTypeInput

func (OutputTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputType)(nil)).Elem()
}

func (i OutputTypeArray) ToOutputTypeArrayOutput() OutputTypeArrayOutput {
	return i.ToOutputTypeArrayOutputWithContext(context.Background())
}

func (i OutputTypeArray) ToOutputTypeArrayOutputWithContext(ctx context.Context) OutputTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputTypeArrayOutput)
}

type OutputTypeOutput struct{ *pulumi.OutputState }

func (OutputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputType)(nil)).Elem()
}

func (o OutputTypeOutput) ToOutputTypeOutput() OutputTypeOutput {
	return o
}

func (o OutputTypeOutput) ToOutputTypeOutputWithContext(ctx context.Context) OutputTypeOutput {
	return o
}

func (o OutputTypeOutput) Datasource() pulumi.AnyOutput {
	return o.ApplyT(func(v OutputType) interface{} { return v.Datasource }).(pulumi.AnyOutput)
}

func (o OutputTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutputType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OutputTypeOutput) Serialization() pulumi.AnyOutput {
	return o.ApplyT(func(v OutputType) interface{} { return v.Serialization }).(pulumi.AnyOutput)
}

type OutputTypeArrayOutput struct{ *pulumi.OutputState }

func (OutputTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputType)(nil)).Elem()
}

func (o OutputTypeArrayOutput) ToOutputTypeArrayOutput() OutputTypeArrayOutput {
	return o
}

func (o OutputTypeArrayOutput) ToOutputTypeArrayOutputWithContext(ctx context.Context) OutputTypeArrayOutput {
	return o
}

func (o OutputTypeArrayOutput) Index(i pulumi.IntInput) OutputTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OutputType {
		return vs[0].([]OutputType)[vs[1].(int)]
	}).(OutputTypeOutput)
}

type OutputResponse struct {
	Datasource    interface{}         `pulumi:"datasource"`
	Diagnostics   DiagnosticsResponse `pulumi:"diagnostics"`
	Etag          string              `pulumi:"etag"`
	Id            string              `pulumi:"id"`
	Name          *string             `pulumi:"name"`
	Serialization interface{}         `pulumi:"serialization"`
	Type          string              `pulumi:"type"`
}

type OutputResponseOutput struct{ *pulumi.OutputState }

func (OutputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputResponse)(nil)).Elem()
}

func (o OutputResponseOutput) ToOutputResponseOutput() OutputResponseOutput {
	return o
}

func (o OutputResponseOutput) ToOutputResponseOutputWithContext(ctx context.Context) OutputResponseOutput {
	return o
}

func (o OutputResponseOutput) Datasource() pulumi.AnyOutput {
	return o.ApplyT(func(v OutputResponse) interface{} { return v.Datasource }).(pulumi.AnyOutput)
}

func (o OutputResponseOutput) Diagnostics() DiagnosticsResponseOutput {
	return o.ApplyT(func(v OutputResponse) DiagnosticsResponse { return v.Diagnostics }).(DiagnosticsResponseOutput)
}

func (o OutputResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v OutputResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o OutputResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v OutputResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o OutputResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutputResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OutputResponseOutput) Serialization() pulumi.AnyOutput {
	return o.ApplyT(func(v OutputResponse) interface{} { return v.Serialization }).(pulumi.AnyOutput)
}

func (o OutputResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v OutputResponse) string { return v.Type }).(pulumi.StringOutput)
}

type OutputResponseArrayOutput struct{ *pulumi.OutputState }

func (OutputResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputResponse)(nil)).Elem()
}

func (o OutputResponseArrayOutput) ToOutputResponseArrayOutput() OutputResponseArrayOutput {
	return o
}

func (o OutputResponseArrayOutput) ToOutputResponseArrayOutputWithContext(ctx context.Context) OutputResponseArrayOutput {
	return o
}

func (o OutputResponseArrayOutput) Index(i pulumi.IntInput) OutputResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OutputResponse {
		return vs[0].([]OutputResponse)[vs[1].(int)]
	}).(OutputResponseOutput)
}

type PowerBIOutputDataSource struct {
	Dataset                *string `pulumi:"dataset"`
	GroupId                *string `pulumi:"groupId"`
	GroupName              *string `pulumi:"groupName"`
	RefreshToken           *string `pulumi:"refreshToken"`
	Table                  *string `pulumi:"table"`
	TokenUserDisplayName   *string `pulumi:"tokenUserDisplayName"`
	TokenUserPrincipalName *string `pulumi:"tokenUserPrincipalName"`
	Type                   string  `pulumi:"type"`
}

type PowerBIOutputDataSourceResponse struct {
	Dataset                *string `pulumi:"dataset"`
	GroupId                *string `pulumi:"groupId"`
	GroupName              *string `pulumi:"groupName"`
	RefreshToken           *string `pulumi:"refreshToken"`
	Table                  *string `pulumi:"table"`
	TokenUserDisplayName   *string `pulumi:"tokenUserDisplayName"`
	TokenUserPrincipalName *string `pulumi:"tokenUserPrincipalName"`
	Type                   string  `pulumi:"type"`
}

type ReferenceInputProperties struct {
	Datasource    *BlobReferenceInputDataSource `pulumi:"datasource"`
	Serialization interface{}                   `pulumi:"serialization"`
	Type          string                        `pulumi:"type"`
}

type ReferenceInputPropertiesResponse struct {
	Datasource    *BlobReferenceInputDataSourceResponse `pulumi:"datasource"`
	Diagnostics   DiagnosticsResponse                   `pulumi:"diagnostics"`
	Etag          string                                `pulumi:"etag"`
	Serialization interface{}                           `pulumi:"serialization"`
	Type          string                                `pulumi:"type"`
}

type ScalarFunctionProperties struct {
	Binding interface{}         `pulumi:"binding"`
	Inputs  []FunctionInputType `pulumi:"inputs"`
	Output  *FunctionOutputType `pulumi:"output"`
	Type    string              `pulumi:"type"`
}





type ScalarFunctionPropertiesInput interface {
	pulumi.Input

	ToScalarFunctionPropertiesOutput() ScalarFunctionPropertiesOutput
	ToScalarFunctionPropertiesOutputWithContext(context.Context) ScalarFunctionPropertiesOutput
}

type ScalarFunctionPropertiesArgs struct {
	Binding pulumi.Input                `pulumi:"binding"`
	Inputs  FunctionInputTypeArrayInput `pulumi:"inputs"`
	Output  FunctionOutputTypePtrInput  `pulumi:"output"`
	Type    pulumi.StringInput          `pulumi:"type"`
}

func (ScalarFunctionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalarFunctionProperties)(nil)).Elem()
}

func (i ScalarFunctionPropertiesArgs) ToScalarFunctionPropertiesOutput() ScalarFunctionPropertiesOutput {
	return i.ToScalarFunctionPropertiesOutputWithContext(context.Background())
}

func (i ScalarFunctionPropertiesArgs) ToScalarFunctionPropertiesOutputWithContext(ctx context.Context) ScalarFunctionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalarFunctionPropertiesOutput)
}

func (i ScalarFunctionPropertiesArgs) ToScalarFunctionPropertiesPtrOutput() ScalarFunctionPropertiesPtrOutput {
	return i.ToScalarFunctionPropertiesPtrOutputWithContext(context.Background())
}

func (i ScalarFunctionPropertiesArgs) ToScalarFunctionPropertiesPtrOutputWithContext(ctx context.Context) ScalarFunctionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalarFunctionPropertiesOutput).ToScalarFunctionPropertiesPtrOutputWithContext(ctx)
}









type ScalarFunctionPropertiesPtrInput interface {
	pulumi.Input

	ToScalarFunctionPropertiesPtrOutput() ScalarFunctionPropertiesPtrOutput
	ToScalarFunctionPropertiesPtrOutputWithContext(context.Context) ScalarFunctionPropertiesPtrOutput
}

type scalarFunctionPropertiesPtrType ScalarFunctionPropertiesArgs

func ScalarFunctionPropertiesPtr(v *ScalarFunctionPropertiesArgs) ScalarFunctionPropertiesPtrInput {
	return (*scalarFunctionPropertiesPtrType)(v)
}

func (*scalarFunctionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalarFunctionProperties)(nil)).Elem()
}

func (i *scalarFunctionPropertiesPtrType) ToScalarFunctionPropertiesPtrOutput() ScalarFunctionPropertiesPtrOutput {
	return i.ToScalarFunctionPropertiesPtrOutputWithContext(context.Background())
}

func (i *scalarFunctionPropertiesPtrType) ToScalarFunctionPropertiesPtrOutputWithContext(ctx context.Context) ScalarFunctionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalarFunctionPropertiesPtrOutput)
}

type ScalarFunctionPropertiesOutput struct{ *pulumi.OutputState }

func (ScalarFunctionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalarFunctionProperties)(nil)).Elem()
}

func (o ScalarFunctionPropertiesOutput) ToScalarFunctionPropertiesOutput() ScalarFunctionPropertiesOutput {
	return o
}

func (o ScalarFunctionPropertiesOutput) ToScalarFunctionPropertiesOutputWithContext(ctx context.Context) ScalarFunctionPropertiesOutput {
	return o
}

func (o ScalarFunctionPropertiesOutput) ToScalarFunctionPropertiesPtrOutput() ScalarFunctionPropertiesPtrOutput {
	return o.ToScalarFunctionPropertiesPtrOutputWithContext(context.Background())
}

func (o ScalarFunctionPropertiesOutput) ToScalarFunctionPropertiesPtrOutputWithContext(ctx context.Context) ScalarFunctionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScalarFunctionProperties) *ScalarFunctionProperties {
		return &v
	}).(ScalarFunctionPropertiesPtrOutput)
}

func (o ScalarFunctionPropertiesOutput) Binding() pulumi.AnyOutput {
	return o.ApplyT(func(v ScalarFunctionProperties) interface{} { return v.Binding }).(pulumi.AnyOutput)
}

func (o ScalarFunctionPropertiesOutput) Inputs() FunctionInputTypeArrayOutput {
	return o.ApplyT(func(v ScalarFunctionProperties) []FunctionInputType { return v.Inputs }).(FunctionInputTypeArrayOutput)
}

func (o ScalarFunctionPropertiesOutput) Output() FunctionOutputTypePtrOutput {
	return o.ApplyT(func(v ScalarFunctionProperties) *FunctionOutputType { return v.Output }).(FunctionOutputTypePtrOutput)
}

func (o ScalarFunctionPropertiesOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ScalarFunctionProperties) string { return v.Type }).(pulumi.StringOutput)
}

type ScalarFunctionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ScalarFunctionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalarFunctionProperties)(nil)).Elem()
}

func (o ScalarFunctionPropertiesPtrOutput) ToScalarFunctionPropertiesPtrOutput() ScalarFunctionPropertiesPtrOutput {
	return o
}

func (o ScalarFunctionPropertiesPtrOutput) ToScalarFunctionPropertiesPtrOutputWithContext(ctx context.Context) ScalarFunctionPropertiesPtrOutput {
	return o
}

func (o ScalarFunctionPropertiesPtrOutput) Elem() ScalarFunctionPropertiesOutput {
	return o.ApplyT(func(v *ScalarFunctionProperties) ScalarFunctionProperties {
		if v != nil {
			return *v
		}
		var ret ScalarFunctionProperties
		return ret
	}).(ScalarFunctionPropertiesOutput)
}

func (o ScalarFunctionPropertiesPtrOutput) Binding() pulumi.AnyOutput {
	return o.ApplyT(func(v *ScalarFunctionProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.Binding
	}).(pulumi.AnyOutput)
}

func (o ScalarFunctionPropertiesPtrOutput) Inputs() FunctionInputTypeArrayOutput {
	return o.ApplyT(func(v *ScalarFunctionProperties) []FunctionInputType {
		if v == nil {
			return nil
		}
		return v.Inputs
	}).(FunctionInputTypeArrayOutput)
}

func (o ScalarFunctionPropertiesPtrOutput) Output() FunctionOutputTypePtrOutput {
	return o.ApplyT(func(v *ScalarFunctionProperties) *FunctionOutputType {
		if v == nil {
			return nil
		}
		return v.Output
	}).(FunctionOutputTypePtrOutput)
}

func (o ScalarFunctionPropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalarFunctionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ScalarFunctionPropertiesResponse struct {
	Binding interface{}             `pulumi:"binding"`
	Etag    string                  `pulumi:"etag"`
	Inputs  []FunctionInputResponse `pulumi:"inputs"`
	Output  *FunctionOutputResponse `pulumi:"output"`
	Type    string                  `pulumi:"type"`
}

type ScalarFunctionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ScalarFunctionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalarFunctionPropertiesResponse)(nil)).Elem()
}

func (o ScalarFunctionPropertiesResponseOutput) ToScalarFunctionPropertiesResponseOutput() ScalarFunctionPropertiesResponseOutput {
	return o
}

func (o ScalarFunctionPropertiesResponseOutput) ToScalarFunctionPropertiesResponseOutputWithContext(ctx context.Context) ScalarFunctionPropertiesResponseOutput {
	return o
}

func (o ScalarFunctionPropertiesResponseOutput) Binding() pulumi.AnyOutput {
	return o.ApplyT(func(v ScalarFunctionPropertiesResponse) interface{} { return v.Binding }).(pulumi.AnyOutput)
}

func (o ScalarFunctionPropertiesResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ScalarFunctionPropertiesResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o ScalarFunctionPropertiesResponseOutput) Inputs() FunctionInputResponseArrayOutput {
	return o.ApplyT(func(v ScalarFunctionPropertiesResponse) []FunctionInputResponse { return v.Inputs }).(FunctionInputResponseArrayOutput)
}

func (o ScalarFunctionPropertiesResponseOutput) Output() FunctionOutputResponsePtrOutput {
	return o.ApplyT(func(v ScalarFunctionPropertiesResponse) *FunctionOutputResponse { return v.Output }).(FunctionOutputResponsePtrOutput)
}

func (o ScalarFunctionPropertiesResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ScalarFunctionPropertiesResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ScalarFunctionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ScalarFunctionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalarFunctionPropertiesResponse)(nil)).Elem()
}

func (o ScalarFunctionPropertiesResponsePtrOutput) ToScalarFunctionPropertiesResponsePtrOutput() ScalarFunctionPropertiesResponsePtrOutput {
	return o
}

func (o ScalarFunctionPropertiesResponsePtrOutput) ToScalarFunctionPropertiesResponsePtrOutputWithContext(ctx context.Context) ScalarFunctionPropertiesResponsePtrOutput {
	return o
}

func (o ScalarFunctionPropertiesResponsePtrOutput) Elem() ScalarFunctionPropertiesResponseOutput {
	return o.ApplyT(func(v *ScalarFunctionPropertiesResponse) ScalarFunctionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ScalarFunctionPropertiesResponse
		return ret
	}).(ScalarFunctionPropertiesResponseOutput)
}

func (o ScalarFunctionPropertiesResponsePtrOutput) Binding() pulumi.AnyOutput {
	return o.ApplyT(func(v *ScalarFunctionPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Binding
	}).(pulumi.AnyOutput)
}

func (o ScalarFunctionPropertiesResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalarFunctionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o ScalarFunctionPropertiesResponsePtrOutput) Inputs() FunctionInputResponseArrayOutput {
	return o.ApplyT(func(v *ScalarFunctionPropertiesResponse) []FunctionInputResponse {
		if v == nil {
			return nil
		}
		return v.Inputs
	}).(FunctionInputResponseArrayOutput)
}

func (o ScalarFunctionPropertiesResponsePtrOutput) Output() FunctionOutputResponsePtrOutput {
	return o.ApplyT(func(v *ScalarFunctionPropertiesResponse) *FunctionOutputResponse {
		if v == nil {
			return nil
		}
		return v.Output
	}).(FunctionOutputResponsePtrOutput)
}

func (o ScalarFunctionPropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalarFunctionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ServiceBusQueueOutputDataSource struct {
	PropertyColumns        []string `pulumi:"propertyColumns"`
	QueueName              *string  `pulumi:"queueName"`
	ServiceBusNamespace    *string  `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string  `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string  `pulumi:"sharedAccessPolicyName"`
	Type                   string   `pulumi:"type"`
}

type ServiceBusQueueOutputDataSourceResponse struct {
	PropertyColumns        []string `pulumi:"propertyColumns"`
	QueueName              *string  `pulumi:"queueName"`
	ServiceBusNamespace    *string  `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string  `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string  `pulumi:"sharedAccessPolicyName"`
	Type                   string   `pulumi:"type"`
}

type ServiceBusTopicOutputDataSource struct {
	PropertyColumns        []string `pulumi:"propertyColumns"`
	ServiceBusNamespace    *string  `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string  `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string  `pulumi:"sharedAccessPolicyName"`
	TopicName              *string  `pulumi:"topicName"`
	Type                   string   `pulumi:"type"`
}

type ServiceBusTopicOutputDataSourceResponse struct {
	PropertyColumns        []string `pulumi:"propertyColumns"`
	ServiceBusNamespace    *string  `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string  `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string  `pulumi:"sharedAccessPolicyName"`
	TopicName              *string  `pulumi:"topicName"`
	Type                   string   `pulumi:"type"`
}

type Sku struct {
	Name *string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name *string `pulumi:"name"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type StorageAccount struct {
	AccountKey  *string `pulumi:"accountKey"`
	AccountName *string `pulumi:"accountName"`
}

type StorageAccountResponse struct {
	AccountKey  *string `pulumi:"accountKey"`
	AccountName *string `pulumi:"accountName"`
}

type StreamInputProperties struct {
	Datasource    interface{} `pulumi:"datasource"`
	Serialization interface{} `pulumi:"serialization"`
	Type          string      `pulumi:"type"`
}

type StreamInputPropertiesResponse struct {
	Datasource    interface{}         `pulumi:"datasource"`
	Diagnostics   DiagnosticsResponse `pulumi:"diagnostics"`
	Etag          string              `pulumi:"etag"`
	Serialization interface{}         `pulumi:"serialization"`
	Type          string              `pulumi:"type"`
}

type Transformation struct {
	Name           *string `pulumi:"name"`
	Query          *string `pulumi:"query"`
	StreamingUnits *int    `pulumi:"streamingUnits"`
}





type TransformationInput interface {
	pulumi.Input

	ToTransformationOutput() TransformationOutput
	ToTransformationOutputWithContext(context.Context) TransformationOutput
}

type TransformationArgs struct {
	Name           pulumi.StringPtrInput `pulumi:"name"`
	Query          pulumi.StringPtrInput `pulumi:"query"`
	StreamingUnits pulumi.IntPtrInput    `pulumi:"streamingUnits"`
}

func (TransformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Transformation)(nil)).Elem()
}

func (i TransformationArgs) ToTransformationOutput() TransformationOutput {
	return i.ToTransformationOutputWithContext(context.Background())
}

func (i TransformationArgs) ToTransformationOutputWithContext(ctx context.Context) TransformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationOutput)
}

func (i TransformationArgs) ToTransformationPtrOutput() TransformationPtrOutput {
	return i.ToTransformationPtrOutputWithContext(context.Background())
}

func (i TransformationArgs) ToTransformationPtrOutputWithContext(ctx context.Context) TransformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationOutput).ToTransformationPtrOutputWithContext(ctx)
}









type TransformationPtrInput interface {
	pulumi.Input

	ToTransformationPtrOutput() TransformationPtrOutput
	ToTransformationPtrOutputWithContext(context.Context) TransformationPtrOutput
}

type transformationPtrType TransformationArgs

func TransformationPtr(v *TransformationArgs) TransformationPtrInput {
	return (*transformationPtrType)(v)
}

func (*transformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Transformation)(nil)).Elem()
}

func (i *transformationPtrType) ToTransformationPtrOutput() TransformationPtrOutput {
	return i.ToTransformationPtrOutputWithContext(context.Background())
}

func (i *transformationPtrType) ToTransformationPtrOutputWithContext(ctx context.Context) TransformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationPtrOutput)
}

type TransformationOutput struct{ *pulumi.OutputState }

func (TransformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Transformation)(nil)).Elem()
}

func (o TransformationOutput) ToTransformationOutput() TransformationOutput {
	return o
}

func (o TransformationOutput) ToTransformationOutputWithContext(ctx context.Context) TransformationOutput {
	return o
}

func (o TransformationOutput) ToTransformationPtrOutput() TransformationPtrOutput {
	return o.ToTransformationPtrOutputWithContext(context.Background())
}

func (o TransformationOutput) ToTransformationPtrOutputWithContext(ctx context.Context) TransformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Transformation) *Transformation {
		return &v
	}).(TransformationPtrOutput)
}

func (o TransformationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Transformation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TransformationOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Transformation) *string { return v.Query }).(pulumi.StringPtrOutput)
}

func (o TransformationOutput) StreamingUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Transformation) *int { return v.StreamingUnits }).(pulumi.IntPtrOutput)
}

type TransformationPtrOutput struct{ *pulumi.OutputState }

func (TransformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Transformation)(nil)).Elem()
}

func (o TransformationPtrOutput) ToTransformationPtrOutput() TransformationPtrOutput {
	return o
}

func (o TransformationPtrOutput) ToTransformationPtrOutputWithContext(ctx context.Context) TransformationPtrOutput {
	return o
}

func (o TransformationPtrOutput) Elem() TransformationOutput {
	return o.ApplyT(func(v *Transformation) Transformation {
		if v != nil {
			return *v
		}
		var ret Transformation
		return ret
	}).(TransformationOutput)
}

func (o TransformationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Transformation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o TransformationPtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Transformation) *string {
		if v == nil {
			return nil
		}
		return v.Query
	}).(pulumi.StringPtrOutput)
}

func (o TransformationPtrOutput) StreamingUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Transformation) *int {
		if v == nil {
			return nil
		}
		return v.StreamingUnits
	}).(pulumi.IntPtrOutput)
}

type TransformationResponse struct {
	Etag           string  `pulumi:"etag"`
	Id             string  `pulumi:"id"`
	Name           *string `pulumi:"name"`
	Query          *string `pulumi:"query"`
	StreamingUnits *int    `pulumi:"streamingUnits"`
	Type           string  `pulumi:"type"`
}

type TransformationResponseOutput struct{ *pulumi.OutputState }

func (TransformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransformationResponse)(nil)).Elem()
}

func (o TransformationResponseOutput) ToTransformationResponseOutput() TransformationResponseOutput {
	return o
}

func (o TransformationResponseOutput) ToTransformationResponseOutputWithContext(ctx context.Context) TransformationResponseOutput {
	return o
}

func (o TransformationResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v TransformationResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o TransformationResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v TransformationResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o TransformationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransformationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TransformationResponseOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransformationResponse) *string { return v.Query }).(pulumi.StringPtrOutput)
}

func (o TransformationResponseOutput) StreamingUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TransformationResponse) *int { return v.StreamingUnits }).(pulumi.IntPtrOutput)
}

func (o TransformationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v TransformationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type TransformationResponsePtrOutput struct{ *pulumi.OutputState }

func (TransformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransformationResponse)(nil)).Elem()
}

func (o TransformationResponsePtrOutput) ToTransformationResponsePtrOutput() TransformationResponsePtrOutput {
	return o
}

func (o TransformationResponsePtrOutput) ToTransformationResponsePtrOutputWithContext(ctx context.Context) TransformationResponsePtrOutput {
	return o
}

func (o TransformationResponsePtrOutput) Elem() TransformationResponseOutput {
	return o.ApplyT(func(v *TransformationResponse) TransformationResponse {
		if v != nil {
			return *v
		}
		var ret TransformationResponse
		return ret
	}).(TransformationResponseOutput)
}

func (o TransformationResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o TransformationResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o TransformationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o TransformationResponsePtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Query
	}).(pulumi.StringPtrOutput)
}

func (o TransformationResponsePtrOutput) StreamingUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TransformationResponse) *int {
		if v == nil {
			return nil
		}
		return v.StreamingUnits
	}).(pulumi.IntPtrOutput)
}

func (o TransformationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DiagnosticConditionResponseOutput{})
	pulumi.RegisterOutputType(DiagnosticConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(DiagnosticsResponseOutput{})
	pulumi.RegisterOutputType(FunctionTypeOutput{})
	pulumi.RegisterOutputType(FunctionTypeArrayOutput{})
	pulumi.RegisterOutputType(FunctionInputTypeOutput{})
	pulumi.RegisterOutputType(FunctionInputTypeArrayOutput{})
	pulumi.RegisterOutputType(FunctionInputResponseOutput{})
	pulumi.RegisterOutputType(FunctionInputResponseArrayOutput{})
	pulumi.RegisterOutputType(FunctionOutputTypeOutput{})
	pulumi.RegisterOutputType(FunctionOutputTypePtrOutput{})
	pulumi.RegisterOutputType(FunctionOutputResponseOutput{})
	pulumi.RegisterOutputType(FunctionOutputResponsePtrOutput{})
	pulumi.RegisterOutputType(FunctionResponseOutput{})
	pulumi.RegisterOutputType(FunctionResponseArrayOutput{})
	pulumi.RegisterOutputType(InputTypeOutput{})
	pulumi.RegisterOutputType(InputTypeArrayOutput{})
	pulumi.RegisterOutputType(InputResponseOutput{})
	pulumi.RegisterOutputType(InputResponseArrayOutput{})
	pulumi.RegisterOutputType(OutputTypeOutput{})
	pulumi.RegisterOutputType(OutputTypeArrayOutput{})
	pulumi.RegisterOutputType(OutputResponseOutput{})
	pulumi.RegisterOutputType(OutputResponseArrayOutput{})
	pulumi.RegisterOutputType(ScalarFunctionPropertiesOutput{})
	pulumi.RegisterOutputType(ScalarFunctionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ScalarFunctionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ScalarFunctionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(TransformationOutput{})
	pulumi.RegisterOutputType(TransformationPtrOutput{})
	pulumi.RegisterOutputType(TransformationResponseOutput{})
	pulumi.RegisterOutputType(TransformationResponsePtrOutput{})
}
