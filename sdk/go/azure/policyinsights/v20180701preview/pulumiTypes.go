


package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ErrorDefinitionResponse struct {
	AdditionalInfo []TypedErrorInfoResponse  `pulumi:"additionalInfo"`
	Code           string                    `pulumi:"code"`
	Details        []ErrorDefinitionResponse `pulumi:"details"`
	Message        string                    `pulumi:"message"`
	Target         string                    `pulumi:"target"`
}

type ErrorDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ErrorDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDefinitionResponse)(nil)).Elem()
}

func (o ErrorDefinitionResponseOutput) ToErrorDefinitionResponseOutput() ErrorDefinitionResponseOutput {
	return o
}

func (o ErrorDefinitionResponseOutput) ToErrorDefinitionResponseOutputWithContext(ctx context.Context) ErrorDefinitionResponseOutput {
	return o
}

func (o ErrorDefinitionResponseOutput) AdditionalInfo() TypedErrorInfoResponseArrayOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) []TypedErrorInfoResponse { return v.AdditionalInfo }).(TypedErrorInfoResponseArrayOutput)
}

func (o ErrorDefinitionResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDefinitionResponseOutput) Details() ErrorDefinitionResponseArrayOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) []ErrorDefinitionResponse { return v.Details }).(ErrorDefinitionResponseArrayOutput)
}

func (o ErrorDefinitionResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ErrorDefinitionResponseOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDefinitionResponse) string { return v.Target }).(pulumi.StringOutput)
}

type ErrorDefinitionResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorDefinitionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDefinitionResponse)(nil)).Elem()
}

func (o ErrorDefinitionResponseArrayOutput) ToErrorDefinitionResponseArrayOutput() ErrorDefinitionResponseArrayOutput {
	return o
}

func (o ErrorDefinitionResponseArrayOutput) ToErrorDefinitionResponseArrayOutputWithContext(ctx context.Context) ErrorDefinitionResponseArrayOutput {
	return o
}

func (o ErrorDefinitionResponseArrayOutput) Index(i pulumi.IntInput) ErrorDefinitionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorDefinitionResponse {
		return vs[0].([]ErrorDefinitionResponse)[vs[1].(int)]
	}).(ErrorDefinitionResponseOutput)
}

type RemediationDeploymentResponse struct {
	CreatedOn            string                  `pulumi:"createdOn"`
	DeploymentId         string                  `pulumi:"deploymentId"`
	Error                ErrorDefinitionResponse `pulumi:"error"`
	LastUpdatedOn        string                  `pulumi:"lastUpdatedOn"`
	RemediatedResourceId string                  `pulumi:"remediatedResourceId"`
	ResourceLocation     string                  `pulumi:"resourceLocation"`
	Status               string                  `pulumi:"status"`
}

type RemediationDeploymentResponseOutput struct{ *pulumi.OutputState }

func (RemediationDeploymentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationDeploymentResponse)(nil)).Elem()
}

func (o RemediationDeploymentResponseOutput) ToRemediationDeploymentResponseOutput() RemediationDeploymentResponseOutput {
	return o
}

func (o RemediationDeploymentResponseOutput) ToRemediationDeploymentResponseOutputWithContext(ctx context.Context) RemediationDeploymentResponseOutput {
	return o
}

func (o RemediationDeploymentResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o RemediationDeploymentResponseOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) string { return v.DeploymentId }).(pulumi.StringOutput)
}

func (o RemediationDeploymentResponseOutput) Error() ErrorDefinitionResponseOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) ErrorDefinitionResponse { return v.Error }).(ErrorDefinitionResponseOutput)
}

func (o RemediationDeploymentResponseOutput) LastUpdatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) string { return v.LastUpdatedOn }).(pulumi.StringOutput)
}

func (o RemediationDeploymentResponseOutput) RemediatedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) string { return v.RemediatedResourceId }).(pulumi.StringOutput)
}

func (o RemediationDeploymentResponseOutput) ResourceLocation() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) string { return v.ResourceLocation }).(pulumi.StringOutput)
}

func (o RemediationDeploymentResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v RemediationDeploymentResponse) string { return v.Status }).(pulumi.StringOutput)
}

type RemediationDeploymentResponseArrayOutput struct{ *pulumi.OutputState }

func (RemediationDeploymentResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RemediationDeploymentResponse)(nil)).Elem()
}

func (o RemediationDeploymentResponseArrayOutput) ToRemediationDeploymentResponseArrayOutput() RemediationDeploymentResponseArrayOutput {
	return o
}

func (o RemediationDeploymentResponseArrayOutput) ToRemediationDeploymentResponseArrayOutputWithContext(ctx context.Context) RemediationDeploymentResponseArrayOutput {
	return o
}

func (o RemediationDeploymentResponseArrayOutput) Index(i pulumi.IntInput) RemediationDeploymentResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RemediationDeploymentResponse {
		return vs[0].([]RemediationDeploymentResponse)[vs[1].(int)]
	}).(RemediationDeploymentResponseOutput)
}

type RemediationDeploymentSummary struct {
	FailedDeployments     *int `pulumi:"failedDeployments"`
	SuccessfulDeployments *int `pulumi:"successfulDeployments"`
	TotalDeployments      *int `pulumi:"totalDeployments"`
}





type RemediationDeploymentSummaryInput interface {
	pulumi.Input

	ToRemediationDeploymentSummaryOutput() RemediationDeploymentSummaryOutput
	ToRemediationDeploymentSummaryOutputWithContext(context.Context) RemediationDeploymentSummaryOutput
}

type RemediationDeploymentSummaryArgs struct {
	FailedDeployments     pulumi.IntPtrInput `pulumi:"failedDeployments"`
	SuccessfulDeployments pulumi.IntPtrInput `pulumi:"successfulDeployments"`
	TotalDeployments      pulumi.IntPtrInput `pulumi:"totalDeployments"`
}

func (RemediationDeploymentSummaryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationDeploymentSummary)(nil)).Elem()
}

func (i RemediationDeploymentSummaryArgs) ToRemediationDeploymentSummaryOutput() RemediationDeploymentSummaryOutput {
	return i.ToRemediationDeploymentSummaryOutputWithContext(context.Background())
}

func (i RemediationDeploymentSummaryArgs) ToRemediationDeploymentSummaryOutputWithContext(ctx context.Context) RemediationDeploymentSummaryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationDeploymentSummaryOutput)
}

func (i RemediationDeploymentSummaryArgs) ToRemediationDeploymentSummaryPtrOutput() RemediationDeploymentSummaryPtrOutput {
	return i.ToRemediationDeploymentSummaryPtrOutputWithContext(context.Background())
}

func (i RemediationDeploymentSummaryArgs) ToRemediationDeploymentSummaryPtrOutputWithContext(ctx context.Context) RemediationDeploymentSummaryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationDeploymentSummaryOutput).ToRemediationDeploymentSummaryPtrOutputWithContext(ctx)
}









type RemediationDeploymentSummaryPtrInput interface {
	pulumi.Input

	ToRemediationDeploymentSummaryPtrOutput() RemediationDeploymentSummaryPtrOutput
	ToRemediationDeploymentSummaryPtrOutputWithContext(context.Context) RemediationDeploymentSummaryPtrOutput
}

type remediationDeploymentSummaryPtrType RemediationDeploymentSummaryArgs

func RemediationDeploymentSummaryPtr(v *RemediationDeploymentSummaryArgs) RemediationDeploymentSummaryPtrInput {
	return (*remediationDeploymentSummaryPtrType)(v)
}

func (*remediationDeploymentSummaryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationDeploymentSummary)(nil)).Elem()
}

func (i *remediationDeploymentSummaryPtrType) ToRemediationDeploymentSummaryPtrOutput() RemediationDeploymentSummaryPtrOutput {
	return i.ToRemediationDeploymentSummaryPtrOutputWithContext(context.Background())
}

func (i *remediationDeploymentSummaryPtrType) ToRemediationDeploymentSummaryPtrOutputWithContext(ctx context.Context) RemediationDeploymentSummaryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationDeploymentSummaryPtrOutput)
}

type RemediationDeploymentSummaryOutput struct{ *pulumi.OutputState }

func (RemediationDeploymentSummaryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationDeploymentSummary)(nil)).Elem()
}

func (o RemediationDeploymentSummaryOutput) ToRemediationDeploymentSummaryOutput() RemediationDeploymentSummaryOutput {
	return o
}

func (o RemediationDeploymentSummaryOutput) ToRemediationDeploymentSummaryOutputWithContext(ctx context.Context) RemediationDeploymentSummaryOutput {
	return o
}

func (o RemediationDeploymentSummaryOutput) ToRemediationDeploymentSummaryPtrOutput() RemediationDeploymentSummaryPtrOutput {
	return o.ToRemediationDeploymentSummaryPtrOutputWithContext(context.Background())
}

func (o RemediationDeploymentSummaryOutput) ToRemediationDeploymentSummaryPtrOutputWithContext(ctx context.Context) RemediationDeploymentSummaryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RemediationDeploymentSummary) *RemediationDeploymentSummary {
		return &v
	}).(RemediationDeploymentSummaryPtrOutput)
}

func (o RemediationDeploymentSummaryOutput) FailedDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RemediationDeploymentSummary) *int { return v.FailedDeployments }).(pulumi.IntPtrOutput)
}

func (o RemediationDeploymentSummaryOutput) SuccessfulDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RemediationDeploymentSummary) *int { return v.SuccessfulDeployments }).(pulumi.IntPtrOutput)
}

func (o RemediationDeploymentSummaryOutput) TotalDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RemediationDeploymentSummary) *int { return v.TotalDeployments }).(pulumi.IntPtrOutput)
}

type RemediationDeploymentSummaryPtrOutput struct{ *pulumi.OutputState }

func (RemediationDeploymentSummaryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationDeploymentSummary)(nil)).Elem()
}

func (o RemediationDeploymentSummaryPtrOutput) ToRemediationDeploymentSummaryPtrOutput() RemediationDeploymentSummaryPtrOutput {
	return o
}

func (o RemediationDeploymentSummaryPtrOutput) ToRemediationDeploymentSummaryPtrOutputWithContext(ctx context.Context) RemediationDeploymentSummaryPtrOutput {
	return o
}

func (o RemediationDeploymentSummaryPtrOutput) Elem() RemediationDeploymentSummaryOutput {
	return o.ApplyT(func(v *RemediationDeploymentSummary) RemediationDeploymentSummary {
		if v != nil {
			return *v
		}
		var ret RemediationDeploymentSummary
		return ret
	}).(RemediationDeploymentSummaryOutput)
}

func (o RemediationDeploymentSummaryPtrOutput) FailedDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RemediationDeploymentSummary) *int {
		if v == nil {
			return nil
		}
		return v.FailedDeployments
	}).(pulumi.IntPtrOutput)
}

func (o RemediationDeploymentSummaryPtrOutput) SuccessfulDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RemediationDeploymentSummary) *int {
		if v == nil {
			return nil
		}
		return v.SuccessfulDeployments
	}).(pulumi.IntPtrOutput)
}

func (o RemediationDeploymentSummaryPtrOutput) TotalDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RemediationDeploymentSummary) *int {
		if v == nil {
			return nil
		}
		return v.TotalDeployments
	}).(pulumi.IntPtrOutput)
}

type RemediationDeploymentSummaryResponse struct {
	FailedDeployments     *int `pulumi:"failedDeployments"`
	SuccessfulDeployments *int `pulumi:"successfulDeployments"`
	TotalDeployments      *int `pulumi:"totalDeployments"`
}

type RemediationDeploymentSummaryResponseOutput struct{ *pulumi.OutputState }

func (RemediationDeploymentSummaryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationDeploymentSummaryResponse)(nil)).Elem()
}

func (o RemediationDeploymentSummaryResponseOutput) ToRemediationDeploymentSummaryResponseOutput() RemediationDeploymentSummaryResponseOutput {
	return o
}

func (o RemediationDeploymentSummaryResponseOutput) ToRemediationDeploymentSummaryResponseOutputWithContext(ctx context.Context) RemediationDeploymentSummaryResponseOutput {
	return o
}

func (o RemediationDeploymentSummaryResponseOutput) FailedDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RemediationDeploymentSummaryResponse) *int { return v.FailedDeployments }).(pulumi.IntPtrOutput)
}

func (o RemediationDeploymentSummaryResponseOutput) SuccessfulDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RemediationDeploymentSummaryResponse) *int { return v.SuccessfulDeployments }).(pulumi.IntPtrOutput)
}

func (o RemediationDeploymentSummaryResponseOutput) TotalDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RemediationDeploymentSummaryResponse) *int { return v.TotalDeployments }).(pulumi.IntPtrOutput)
}

type RemediationDeploymentSummaryResponsePtrOutput struct{ *pulumi.OutputState }

func (RemediationDeploymentSummaryResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationDeploymentSummaryResponse)(nil)).Elem()
}

func (o RemediationDeploymentSummaryResponsePtrOutput) ToRemediationDeploymentSummaryResponsePtrOutput() RemediationDeploymentSummaryResponsePtrOutput {
	return o
}

func (o RemediationDeploymentSummaryResponsePtrOutput) ToRemediationDeploymentSummaryResponsePtrOutputWithContext(ctx context.Context) RemediationDeploymentSummaryResponsePtrOutput {
	return o
}

func (o RemediationDeploymentSummaryResponsePtrOutput) Elem() RemediationDeploymentSummaryResponseOutput {
	return o.ApplyT(func(v *RemediationDeploymentSummaryResponse) RemediationDeploymentSummaryResponse {
		if v != nil {
			return *v
		}
		var ret RemediationDeploymentSummaryResponse
		return ret
	}).(RemediationDeploymentSummaryResponseOutput)
}

func (o RemediationDeploymentSummaryResponsePtrOutput) FailedDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RemediationDeploymentSummaryResponse) *int {
		if v == nil {
			return nil
		}
		return v.FailedDeployments
	}).(pulumi.IntPtrOutput)
}

func (o RemediationDeploymentSummaryResponsePtrOutput) SuccessfulDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RemediationDeploymentSummaryResponse) *int {
		if v == nil {
			return nil
		}
		return v.SuccessfulDeployments
	}).(pulumi.IntPtrOutput)
}

func (o RemediationDeploymentSummaryResponsePtrOutput) TotalDeployments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RemediationDeploymentSummaryResponse) *int {
		if v == nil {
			return nil
		}
		return v.TotalDeployments
	}).(pulumi.IntPtrOutput)
}

type RemediationFilters struct {
	Locations []string `pulumi:"locations"`
}





type RemediationFiltersInput interface {
	pulumi.Input

	ToRemediationFiltersOutput() RemediationFiltersOutput
	ToRemediationFiltersOutputWithContext(context.Context) RemediationFiltersOutput
}

type RemediationFiltersArgs struct {
	Locations pulumi.StringArrayInput `pulumi:"locations"`
}

func (RemediationFiltersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationFilters)(nil)).Elem()
}

func (i RemediationFiltersArgs) ToRemediationFiltersOutput() RemediationFiltersOutput {
	return i.ToRemediationFiltersOutputWithContext(context.Background())
}

func (i RemediationFiltersArgs) ToRemediationFiltersOutputWithContext(ctx context.Context) RemediationFiltersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationFiltersOutput)
}

func (i RemediationFiltersArgs) ToRemediationFiltersPtrOutput() RemediationFiltersPtrOutput {
	return i.ToRemediationFiltersPtrOutputWithContext(context.Background())
}

func (i RemediationFiltersArgs) ToRemediationFiltersPtrOutputWithContext(ctx context.Context) RemediationFiltersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationFiltersOutput).ToRemediationFiltersPtrOutputWithContext(ctx)
}









type RemediationFiltersPtrInput interface {
	pulumi.Input

	ToRemediationFiltersPtrOutput() RemediationFiltersPtrOutput
	ToRemediationFiltersPtrOutputWithContext(context.Context) RemediationFiltersPtrOutput
}

type remediationFiltersPtrType RemediationFiltersArgs

func RemediationFiltersPtr(v *RemediationFiltersArgs) RemediationFiltersPtrInput {
	return (*remediationFiltersPtrType)(v)
}

func (*remediationFiltersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationFilters)(nil)).Elem()
}

func (i *remediationFiltersPtrType) ToRemediationFiltersPtrOutput() RemediationFiltersPtrOutput {
	return i.ToRemediationFiltersPtrOutputWithContext(context.Background())
}

func (i *remediationFiltersPtrType) ToRemediationFiltersPtrOutputWithContext(ctx context.Context) RemediationFiltersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationFiltersPtrOutput)
}

type RemediationFiltersOutput struct{ *pulumi.OutputState }

func (RemediationFiltersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationFilters)(nil)).Elem()
}

func (o RemediationFiltersOutput) ToRemediationFiltersOutput() RemediationFiltersOutput {
	return o
}

func (o RemediationFiltersOutput) ToRemediationFiltersOutputWithContext(ctx context.Context) RemediationFiltersOutput {
	return o
}

func (o RemediationFiltersOutput) ToRemediationFiltersPtrOutput() RemediationFiltersPtrOutput {
	return o.ToRemediationFiltersPtrOutputWithContext(context.Background())
}

func (o RemediationFiltersOutput) ToRemediationFiltersPtrOutputWithContext(ctx context.Context) RemediationFiltersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RemediationFilters) *RemediationFilters {
		return &v
	}).(RemediationFiltersPtrOutput)
}

func (o RemediationFiltersOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RemediationFilters) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

type RemediationFiltersPtrOutput struct{ *pulumi.OutputState }

func (RemediationFiltersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationFilters)(nil)).Elem()
}

func (o RemediationFiltersPtrOutput) ToRemediationFiltersPtrOutput() RemediationFiltersPtrOutput {
	return o
}

func (o RemediationFiltersPtrOutput) ToRemediationFiltersPtrOutputWithContext(ctx context.Context) RemediationFiltersPtrOutput {
	return o
}

func (o RemediationFiltersPtrOutput) Elem() RemediationFiltersOutput {
	return o.ApplyT(func(v *RemediationFilters) RemediationFilters {
		if v != nil {
			return *v
		}
		var ret RemediationFilters
		return ret
	}).(RemediationFiltersOutput)
}

func (o RemediationFiltersPtrOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RemediationFilters) []string {
		if v == nil {
			return nil
		}
		return v.Locations
	}).(pulumi.StringArrayOutput)
}

type RemediationFiltersResponse struct {
	Locations []string `pulumi:"locations"`
}

type RemediationFiltersResponseOutput struct{ *pulumi.OutputState }

func (RemediationFiltersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemediationFiltersResponse)(nil)).Elem()
}

func (o RemediationFiltersResponseOutput) ToRemediationFiltersResponseOutput() RemediationFiltersResponseOutput {
	return o
}

func (o RemediationFiltersResponseOutput) ToRemediationFiltersResponseOutputWithContext(ctx context.Context) RemediationFiltersResponseOutput {
	return o
}

func (o RemediationFiltersResponseOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RemediationFiltersResponse) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

type RemediationFiltersResponsePtrOutput struct{ *pulumi.OutputState }

func (RemediationFiltersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationFiltersResponse)(nil)).Elem()
}

func (o RemediationFiltersResponsePtrOutput) ToRemediationFiltersResponsePtrOutput() RemediationFiltersResponsePtrOutput {
	return o
}

func (o RemediationFiltersResponsePtrOutput) ToRemediationFiltersResponsePtrOutputWithContext(ctx context.Context) RemediationFiltersResponsePtrOutput {
	return o
}

func (o RemediationFiltersResponsePtrOutput) Elem() RemediationFiltersResponseOutput {
	return o.ApplyT(func(v *RemediationFiltersResponse) RemediationFiltersResponse {
		if v != nil {
			return *v
		}
		var ret RemediationFiltersResponse
		return ret
	}).(RemediationFiltersResponseOutput)
}

func (o RemediationFiltersResponsePtrOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RemediationFiltersResponse) []string {
		if v == nil {
			return nil
		}
		return v.Locations
	}).(pulumi.StringArrayOutput)
}

type TypedErrorInfoResponse struct {
	Info interface{} `pulumi:"info"`
	Type string      `pulumi:"type"`
}

type TypedErrorInfoResponseOutput struct{ *pulumi.OutputState }

func (TypedErrorInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TypedErrorInfoResponse)(nil)).Elem()
}

func (o TypedErrorInfoResponseOutput) ToTypedErrorInfoResponseOutput() TypedErrorInfoResponseOutput {
	return o
}

func (o TypedErrorInfoResponseOutput) ToTypedErrorInfoResponseOutputWithContext(ctx context.Context) TypedErrorInfoResponseOutput {
	return o
}

func (o TypedErrorInfoResponseOutput) Info() pulumi.AnyOutput {
	return o.ApplyT(func(v TypedErrorInfoResponse) interface{} { return v.Info }).(pulumi.AnyOutput)
}

func (o TypedErrorInfoResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v TypedErrorInfoResponse) string { return v.Type }).(pulumi.StringOutput)
}

type TypedErrorInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (TypedErrorInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TypedErrorInfoResponse)(nil)).Elem()
}

func (o TypedErrorInfoResponseArrayOutput) ToTypedErrorInfoResponseArrayOutput() TypedErrorInfoResponseArrayOutput {
	return o
}

func (o TypedErrorInfoResponseArrayOutput) ToTypedErrorInfoResponseArrayOutputWithContext(ctx context.Context) TypedErrorInfoResponseArrayOutput {
	return o
}

func (o TypedErrorInfoResponseArrayOutput) Index(i pulumi.IntInput) TypedErrorInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TypedErrorInfoResponse {
		return vs[0].([]TypedErrorInfoResponse)[vs[1].(int)]
	}).(TypedErrorInfoResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ErrorDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ErrorDefinitionResponseArrayOutput{})
	pulumi.RegisterOutputType(RemediationDeploymentResponseOutput{})
	pulumi.RegisterOutputType(RemediationDeploymentResponseArrayOutput{})
	pulumi.RegisterOutputType(RemediationDeploymentSummaryOutput{})
	pulumi.RegisterOutputType(RemediationDeploymentSummaryPtrOutput{})
	pulumi.RegisterOutputType(RemediationDeploymentSummaryResponseOutput{})
	pulumi.RegisterOutputType(RemediationDeploymentSummaryResponsePtrOutput{})
	pulumi.RegisterOutputType(RemediationFiltersOutput{})
	pulumi.RegisterOutputType(RemediationFiltersPtrOutput{})
	pulumi.RegisterOutputType(RemediationFiltersResponseOutput{})
	pulumi.RegisterOutputType(RemediationFiltersResponsePtrOutput{})
	pulumi.RegisterOutputType(TypedErrorInfoResponseOutput{})
	pulumi.RegisterOutputType(TypedErrorInfoResponseArrayOutput{})
}
