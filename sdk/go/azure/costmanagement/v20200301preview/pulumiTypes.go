


package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CostAllocationProportion struct {
	Name       string  `pulumi:"name"`
	Percentage float64 `pulumi:"percentage"`
}





type CostAllocationProportionInput interface {
	pulumi.Input

	ToCostAllocationProportionOutput() CostAllocationProportionOutput
	ToCostAllocationProportionOutputWithContext(context.Context) CostAllocationProportionOutput
}

type CostAllocationProportionArgs struct {
	Name       pulumi.StringInput  `pulumi:"name"`
	Percentage pulumi.Float64Input `pulumi:"percentage"`
}

func (CostAllocationProportionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationProportion)(nil)).Elem()
}

func (i CostAllocationProportionArgs) ToCostAllocationProportionOutput() CostAllocationProportionOutput {
	return i.ToCostAllocationProportionOutputWithContext(context.Background())
}

func (i CostAllocationProportionArgs) ToCostAllocationProportionOutputWithContext(ctx context.Context) CostAllocationProportionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostAllocationProportionOutput)
}





type CostAllocationProportionArrayInput interface {
	pulumi.Input

	ToCostAllocationProportionArrayOutput() CostAllocationProportionArrayOutput
	ToCostAllocationProportionArrayOutputWithContext(context.Context) CostAllocationProportionArrayOutput
}

type CostAllocationProportionArray []CostAllocationProportionInput

func (CostAllocationProportionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CostAllocationProportion)(nil)).Elem()
}

func (i CostAllocationProportionArray) ToCostAllocationProportionArrayOutput() CostAllocationProportionArrayOutput {
	return i.ToCostAllocationProportionArrayOutputWithContext(context.Background())
}

func (i CostAllocationProportionArray) ToCostAllocationProportionArrayOutputWithContext(ctx context.Context) CostAllocationProportionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostAllocationProportionArrayOutput)
}

type CostAllocationProportionOutput struct{ *pulumi.OutputState }

func (CostAllocationProportionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationProportion)(nil)).Elem()
}

func (o CostAllocationProportionOutput) ToCostAllocationProportionOutput() CostAllocationProportionOutput {
	return o
}

func (o CostAllocationProportionOutput) ToCostAllocationProportionOutputWithContext(ctx context.Context) CostAllocationProportionOutput {
	return o
}

func (o CostAllocationProportionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CostAllocationProportion) string { return v.Name }).(pulumi.StringOutput)
}

func (o CostAllocationProportionOutput) Percentage() pulumi.Float64Output {
	return o.ApplyT(func(v CostAllocationProportion) float64 { return v.Percentage }).(pulumi.Float64Output)
}

type CostAllocationProportionArrayOutput struct{ *pulumi.OutputState }

func (CostAllocationProportionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CostAllocationProportion)(nil)).Elem()
}

func (o CostAllocationProportionArrayOutput) ToCostAllocationProportionArrayOutput() CostAllocationProportionArrayOutput {
	return o
}

func (o CostAllocationProportionArrayOutput) ToCostAllocationProportionArrayOutputWithContext(ctx context.Context) CostAllocationProportionArrayOutput {
	return o
}

func (o CostAllocationProportionArrayOutput) Index(i pulumi.IntInput) CostAllocationProportionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CostAllocationProportion {
		return vs[0].([]CostAllocationProportion)[vs[1].(int)]
	}).(CostAllocationProportionOutput)
}

type CostAllocationProportionResponse struct {
	Name       string  `pulumi:"name"`
	Percentage float64 `pulumi:"percentage"`
}

type CostAllocationProportionResponseOutput struct{ *pulumi.OutputState }

func (CostAllocationProportionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationProportionResponse)(nil)).Elem()
}

func (o CostAllocationProportionResponseOutput) ToCostAllocationProportionResponseOutput() CostAllocationProportionResponseOutput {
	return o
}

func (o CostAllocationProportionResponseOutput) ToCostAllocationProportionResponseOutputWithContext(ctx context.Context) CostAllocationProportionResponseOutput {
	return o
}

func (o CostAllocationProportionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CostAllocationProportionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CostAllocationProportionResponseOutput) Percentage() pulumi.Float64Output {
	return o.ApplyT(func(v CostAllocationProportionResponse) float64 { return v.Percentage }).(pulumi.Float64Output)
}

type CostAllocationProportionResponseArrayOutput struct{ *pulumi.OutputState }

func (CostAllocationProportionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CostAllocationProportionResponse)(nil)).Elem()
}

func (o CostAllocationProportionResponseArrayOutput) ToCostAllocationProportionResponseArrayOutput() CostAllocationProportionResponseArrayOutput {
	return o
}

func (o CostAllocationProportionResponseArrayOutput) ToCostAllocationProportionResponseArrayOutputWithContext(ctx context.Context) CostAllocationProportionResponseArrayOutput {
	return o
}

func (o CostAllocationProportionResponseArrayOutput) Index(i pulumi.IntInput) CostAllocationProportionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CostAllocationProportionResponse {
		return vs[0].([]CostAllocationProportionResponse)[vs[1].(int)]
	}).(CostAllocationProportionResponseOutput)
}

type CostAllocationRuleDetails struct {
	SourceResources []SourceCostAllocationResource `pulumi:"sourceResources"`
	TargetResources []TargetCostAllocationResource `pulumi:"targetResources"`
}





type CostAllocationRuleDetailsInput interface {
	pulumi.Input

	ToCostAllocationRuleDetailsOutput() CostAllocationRuleDetailsOutput
	ToCostAllocationRuleDetailsOutputWithContext(context.Context) CostAllocationRuleDetailsOutput
}

type CostAllocationRuleDetailsArgs struct {
	SourceResources SourceCostAllocationResourceArrayInput `pulumi:"sourceResources"`
	TargetResources TargetCostAllocationResourceArrayInput `pulumi:"targetResources"`
}

func (CostAllocationRuleDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationRuleDetails)(nil)).Elem()
}

func (i CostAllocationRuleDetailsArgs) ToCostAllocationRuleDetailsOutput() CostAllocationRuleDetailsOutput {
	return i.ToCostAllocationRuleDetailsOutputWithContext(context.Background())
}

func (i CostAllocationRuleDetailsArgs) ToCostAllocationRuleDetailsOutputWithContext(ctx context.Context) CostAllocationRuleDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostAllocationRuleDetailsOutput)
}

func (i CostAllocationRuleDetailsArgs) ToCostAllocationRuleDetailsPtrOutput() CostAllocationRuleDetailsPtrOutput {
	return i.ToCostAllocationRuleDetailsPtrOutputWithContext(context.Background())
}

func (i CostAllocationRuleDetailsArgs) ToCostAllocationRuleDetailsPtrOutputWithContext(ctx context.Context) CostAllocationRuleDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostAllocationRuleDetailsOutput).ToCostAllocationRuleDetailsPtrOutputWithContext(ctx)
}









type CostAllocationRuleDetailsPtrInput interface {
	pulumi.Input

	ToCostAllocationRuleDetailsPtrOutput() CostAllocationRuleDetailsPtrOutput
	ToCostAllocationRuleDetailsPtrOutputWithContext(context.Context) CostAllocationRuleDetailsPtrOutput
}

type costAllocationRuleDetailsPtrType CostAllocationRuleDetailsArgs

func CostAllocationRuleDetailsPtr(v *CostAllocationRuleDetailsArgs) CostAllocationRuleDetailsPtrInput {
	return (*costAllocationRuleDetailsPtrType)(v)
}

func (*costAllocationRuleDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CostAllocationRuleDetails)(nil)).Elem()
}

func (i *costAllocationRuleDetailsPtrType) ToCostAllocationRuleDetailsPtrOutput() CostAllocationRuleDetailsPtrOutput {
	return i.ToCostAllocationRuleDetailsPtrOutputWithContext(context.Background())
}

func (i *costAllocationRuleDetailsPtrType) ToCostAllocationRuleDetailsPtrOutputWithContext(ctx context.Context) CostAllocationRuleDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostAllocationRuleDetailsPtrOutput)
}

type CostAllocationRuleDetailsOutput struct{ *pulumi.OutputState }

func (CostAllocationRuleDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationRuleDetails)(nil)).Elem()
}

func (o CostAllocationRuleDetailsOutput) ToCostAllocationRuleDetailsOutput() CostAllocationRuleDetailsOutput {
	return o
}

func (o CostAllocationRuleDetailsOutput) ToCostAllocationRuleDetailsOutputWithContext(ctx context.Context) CostAllocationRuleDetailsOutput {
	return o
}

func (o CostAllocationRuleDetailsOutput) ToCostAllocationRuleDetailsPtrOutput() CostAllocationRuleDetailsPtrOutput {
	return o.ToCostAllocationRuleDetailsPtrOutputWithContext(context.Background())
}

func (o CostAllocationRuleDetailsOutput) ToCostAllocationRuleDetailsPtrOutputWithContext(ctx context.Context) CostAllocationRuleDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CostAllocationRuleDetails) *CostAllocationRuleDetails {
		return &v
	}).(CostAllocationRuleDetailsPtrOutput)
}

func (o CostAllocationRuleDetailsOutput) SourceResources() SourceCostAllocationResourceArrayOutput {
	return o.ApplyT(func(v CostAllocationRuleDetails) []SourceCostAllocationResource { return v.SourceResources }).(SourceCostAllocationResourceArrayOutput)
}

func (o CostAllocationRuleDetailsOutput) TargetResources() TargetCostAllocationResourceArrayOutput {
	return o.ApplyT(func(v CostAllocationRuleDetails) []TargetCostAllocationResource { return v.TargetResources }).(TargetCostAllocationResourceArrayOutput)
}

type CostAllocationRuleDetailsPtrOutput struct{ *pulumi.OutputState }

func (CostAllocationRuleDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CostAllocationRuleDetails)(nil)).Elem()
}

func (o CostAllocationRuleDetailsPtrOutput) ToCostAllocationRuleDetailsPtrOutput() CostAllocationRuleDetailsPtrOutput {
	return o
}

func (o CostAllocationRuleDetailsPtrOutput) ToCostAllocationRuleDetailsPtrOutputWithContext(ctx context.Context) CostAllocationRuleDetailsPtrOutput {
	return o
}

func (o CostAllocationRuleDetailsPtrOutput) Elem() CostAllocationRuleDetailsOutput {
	return o.ApplyT(func(v *CostAllocationRuleDetails) CostAllocationRuleDetails {
		if v != nil {
			return *v
		}
		var ret CostAllocationRuleDetails
		return ret
	}).(CostAllocationRuleDetailsOutput)
}

func (o CostAllocationRuleDetailsPtrOutput) SourceResources() SourceCostAllocationResourceArrayOutput {
	return o.ApplyT(func(v *CostAllocationRuleDetails) []SourceCostAllocationResource {
		if v == nil {
			return nil
		}
		return v.SourceResources
	}).(SourceCostAllocationResourceArrayOutput)
}

func (o CostAllocationRuleDetailsPtrOutput) TargetResources() TargetCostAllocationResourceArrayOutput {
	return o.ApplyT(func(v *CostAllocationRuleDetails) []TargetCostAllocationResource {
		if v == nil {
			return nil
		}
		return v.TargetResources
	}).(TargetCostAllocationResourceArrayOutput)
}

type CostAllocationRuleDetailsResponse struct {
	SourceResources []SourceCostAllocationResourceResponse `pulumi:"sourceResources"`
	TargetResources []TargetCostAllocationResourceResponse `pulumi:"targetResources"`
}

type CostAllocationRuleDetailsResponseOutput struct{ *pulumi.OutputState }

func (CostAllocationRuleDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationRuleDetailsResponse)(nil)).Elem()
}

func (o CostAllocationRuleDetailsResponseOutput) ToCostAllocationRuleDetailsResponseOutput() CostAllocationRuleDetailsResponseOutput {
	return o
}

func (o CostAllocationRuleDetailsResponseOutput) ToCostAllocationRuleDetailsResponseOutputWithContext(ctx context.Context) CostAllocationRuleDetailsResponseOutput {
	return o
}

func (o CostAllocationRuleDetailsResponseOutput) SourceResources() SourceCostAllocationResourceResponseArrayOutput {
	return o.ApplyT(func(v CostAllocationRuleDetailsResponse) []SourceCostAllocationResourceResponse {
		return v.SourceResources
	}).(SourceCostAllocationResourceResponseArrayOutput)
}

func (o CostAllocationRuleDetailsResponseOutput) TargetResources() TargetCostAllocationResourceResponseArrayOutput {
	return o.ApplyT(func(v CostAllocationRuleDetailsResponse) []TargetCostAllocationResourceResponse {
		return v.TargetResources
	}).(TargetCostAllocationResourceResponseArrayOutput)
}

type CostAllocationRuleProperties struct {
	Description *string                   `pulumi:"description"`
	Details     CostAllocationRuleDetails `pulumi:"details"`
	Status      string                    `pulumi:"status"`
}





type CostAllocationRulePropertiesInput interface {
	pulumi.Input

	ToCostAllocationRulePropertiesOutput() CostAllocationRulePropertiesOutput
	ToCostAllocationRulePropertiesOutputWithContext(context.Context) CostAllocationRulePropertiesOutput
}

type CostAllocationRulePropertiesArgs struct {
	Description pulumi.StringPtrInput          `pulumi:"description"`
	Details     CostAllocationRuleDetailsInput `pulumi:"details"`
	Status      pulumi.StringInput             `pulumi:"status"`
}

func (CostAllocationRulePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationRuleProperties)(nil)).Elem()
}

func (i CostAllocationRulePropertiesArgs) ToCostAllocationRulePropertiesOutput() CostAllocationRulePropertiesOutput {
	return i.ToCostAllocationRulePropertiesOutputWithContext(context.Background())
}

func (i CostAllocationRulePropertiesArgs) ToCostAllocationRulePropertiesOutputWithContext(ctx context.Context) CostAllocationRulePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostAllocationRulePropertiesOutput)
}

func (i CostAllocationRulePropertiesArgs) ToCostAllocationRulePropertiesPtrOutput() CostAllocationRulePropertiesPtrOutput {
	return i.ToCostAllocationRulePropertiesPtrOutputWithContext(context.Background())
}

func (i CostAllocationRulePropertiesArgs) ToCostAllocationRulePropertiesPtrOutputWithContext(ctx context.Context) CostAllocationRulePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostAllocationRulePropertiesOutput).ToCostAllocationRulePropertiesPtrOutputWithContext(ctx)
}









type CostAllocationRulePropertiesPtrInput interface {
	pulumi.Input

	ToCostAllocationRulePropertiesPtrOutput() CostAllocationRulePropertiesPtrOutput
	ToCostAllocationRulePropertiesPtrOutputWithContext(context.Context) CostAllocationRulePropertiesPtrOutput
}

type costAllocationRulePropertiesPtrType CostAllocationRulePropertiesArgs

func CostAllocationRulePropertiesPtr(v *CostAllocationRulePropertiesArgs) CostAllocationRulePropertiesPtrInput {
	return (*costAllocationRulePropertiesPtrType)(v)
}

func (*costAllocationRulePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CostAllocationRuleProperties)(nil)).Elem()
}

func (i *costAllocationRulePropertiesPtrType) ToCostAllocationRulePropertiesPtrOutput() CostAllocationRulePropertiesPtrOutput {
	return i.ToCostAllocationRulePropertiesPtrOutputWithContext(context.Background())
}

func (i *costAllocationRulePropertiesPtrType) ToCostAllocationRulePropertiesPtrOutputWithContext(ctx context.Context) CostAllocationRulePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostAllocationRulePropertiesPtrOutput)
}

type CostAllocationRulePropertiesOutput struct{ *pulumi.OutputState }

func (CostAllocationRulePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationRuleProperties)(nil)).Elem()
}

func (o CostAllocationRulePropertiesOutput) ToCostAllocationRulePropertiesOutput() CostAllocationRulePropertiesOutput {
	return o
}

func (o CostAllocationRulePropertiesOutput) ToCostAllocationRulePropertiesOutputWithContext(ctx context.Context) CostAllocationRulePropertiesOutput {
	return o
}

func (o CostAllocationRulePropertiesOutput) ToCostAllocationRulePropertiesPtrOutput() CostAllocationRulePropertiesPtrOutput {
	return o.ToCostAllocationRulePropertiesPtrOutputWithContext(context.Background())
}

func (o CostAllocationRulePropertiesOutput) ToCostAllocationRulePropertiesPtrOutputWithContext(ctx context.Context) CostAllocationRulePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CostAllocationRuleProperties) *CostAllocationRuleProperties {
		return &v
	}).(CostAllocationRulePropertiesPtrOutput)
}

func (o CostAllocationRulePropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CostAllocationRuleProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CostAllocationRulePropertiesOutput) Details() CostAllocationRuleDetailsOutput {
	return o.ApplyT(func(v CostAllocationRuleProperties) CostAllocationRuleDetails { return v.Details }).(CostAllocationRuleDetailsOutput)
}

func (o CostAllocationRulePropertiesOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v CostAllocationRuleProperties) string { return v.Status }).(pulumi.StringOutput)
}

type CostAllocationRulePropertiesPtrOutput struct{ *pulumi.OutputState }

func (CostAllocationRulePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CostAllocationRuleProperties)(nil)).Elem()
}

func (o CostAllocationRulePropertiesPtrOutput) ToCostAllocationRulePropertiesPtrOutput() CostAllocationRulePropertiesPtrOutput {
	return o
}

func (o CostAllocationRulePropertiesPtrOutput) ToCostAllocationRulePropertiesPtrOutputWithContext(ctx context.Context) CostAllocationRulePropertiesPtrOutput {
	return o
}

func (o CostAllocationRulePropertiesPtrOutput) Elem() CostAllocationRulePropertiesOutput {
	return o.ApplyT(func(v *CostAllocationRuleProperties) CostAllocationRuleProperties {
		if v != nil {
			return *v
		}
		var ret CostAllocationRuleProperties
		return ret
	}).(CostAllocationRulePropertiesOutput)
}

func (o CostAllocationRulePropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CostAllocationRuleProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o CostAllocationRulePropertiesPtrOutput) Details() CostAllocationRuleDetailsPtrOutput {
	return o.ApplyT(func(v *CostAllocationRuleProperties) *CostAllocationRuleDetails {
		if v == nil {
			return nil
		}
		return &v.Details
	}).(CostAllocationRuleDetailsPtrOutput)
}

func (o CostAllocationRulePropertiesPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CostAllocationRuleProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type CostAllocationRulePropertiesResponse struct {
	CreatedDate string                            `pulumi:"createdDate"`
	Description *string                           `pulumi:"description"`
	Details     CostAllocationRuleDetailsResponse `pulumi:"details"`
	Status      string                            `pulumi:"status"`
	UpdatedDate string                            `pulumi:"updatedDate"`
}

type CostAllocationRulePropertiesResponseOutput struct{ *pulumi.OutputState }

func (CostAllocationRulePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationRulePropertiesResponse)(nil)).Elem()
}

func (o CostAllocationRulePropertiesResponseOutput) ToCostAllocationRulePropertiesResponseOutput() CostAllocationRulePropertiesResponseOutput {
	return o
}

func (o CostAllocationRulePropertiesResponseOutput) ToCostAllocationRulePropertiesResponseOutputWithContext(ctx context.Context) CostAllocationRulePropertiesResponseOutput {
	return o
}

func (o CostAllocationRulePropertiesResponseOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v CostAllocationRulePropertiesResponse) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o CostAllocationRulePropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CostAllocationRulePropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CostAllocationRulePropertiesResponseOutput) Details() CostAllocationRuleDetailsResponseOutput {
	return o.ApplyT(func(v CostAllocationRulePropertiesResponse) CostAllocationRuleDetailsResponse { return v.Details }).(CostAllocationRuleDetailsResponseOutput)
}

func (o CostAllocationRulePropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v CostAllocationRulePropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o CostAllocationRulePropertiesResponseOutput) UpdatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v CostAllocationRulePropertiesResponse) string { return v.UpdatedDate }).(pulumi.StringOutput)
}

type SourceCostAllocationResource struct {
	Name         string   `pulumi:"name"`
	ResourceType string   `pulumi:"resourceType"`
	Values       []string `pulumi:"values"`
}





type SourceCostAllocationResourceInput interface {
	pulumi.Input

	ToSourceCostAllocationResourceOutput() SourceCostAllocationResourceOutput
	ToSourceCostAllocationResourceOutputWithContext(context.Context) SourceCostAllocationResourceOutput
}

type SourceCostAllocationResourceArgs struct {
	Name         pulumi.StringInput      `pulumi:"name"`
	ResourceType pulumi.StringInput      `pulumi:"resourceType"`
	Values       pulumi.StringArrayInput `pulumi:"values"`
}

func (SourceCostAllocationResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceCostAllocationResource)(nil)).Elem()
}

func (i SourceCostAllocationResourceArgs) ToSourceCostAllocationResourceOutput() SourceCostAllocationResourceOutput {
	return i.ToSourceCostAllocationResourceOutputWithContext(context.Background())
}

func (i SourceCostAllocationResourceArgs) ToSourceCostAllocationResourceOutputWithContext(ctx context.Context) SourceCostAllocationResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCostAllocationResourceOutput)
}





type SourceCostAllocationResourceArrayInput interface {
	pulumi.Input

	ToSourceCostAllocationResourceArrayOutput() SourceCostAllocationResourceArrayOutput
	ToSourceCostAllocationResourceArrayOutputWithContext(context.Context) SourceCostAllocationResourceArrayOutput
}

type SourceCostAllocationResourceArray []SourceCostAllocationResourceInput

func (SourceCostAllocationResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SourceCostAllocationResource)(nil)).Elem()
}

func (i SourceCostAllocationResourceArray) ToSourceCostAllocationResourceArrayOutput() SourceCostAllocationResourceArrayOutput {
	return i.ToSourceCostAllocationResourceArrayOutputWithContext(context.Background())
}

func (i SourceCostAllocationResourceArray) ToSourceCostAllocationResourceArrayOutputWithContext(ctx context.Context) SourceCostAllocationResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCostAllocationResourceArrayOutput)
}

type SourceCostAllocationResourceOutput struct{ *pulumi.OutputState }

func (SourceCostAllocationResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceCostAllocationResource)(nil)).Elem()
}

func (o SourceCostAllocationResourceOutput) ToSourceCostAllocationResourceOutput() SourceCostAllocationResourceOutput {
	return o
}

func (o SourceCostAllocationResourceOutput) ToSourceCostAllocationResourceOutputWithContext(ctx context.Context) SourceCostAllocationResourceOutput {
	return o
}

func (o SourceCostAllocationResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SourceCostAllocationResource) string { return v.Name }).(pulumi.StringOutput)
}

func (o SourceCostAllocationResourceOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v SourceCostAllocationResource) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o SourceCostAllocationResourceOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SourceCostAllocationResource) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type SourceCostAllocationResourceArrayOutput struct{ *pulumi.OutputState }

func (SourceCostAllocationResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SourceCostAllocationResource)(nil)).Elem()
}

func (o SourceCostAllocationResourceArrayOutput) ToSourceCostAllocationResourceArrayOutput() SourceCostAllocationResourceArrayOutput {
	return o
}

func (o SourceCostAllocationResourceArrayOutput) ToSourceCostAllocationResourceArrayOutputWithContext(ctx context.Context) SourceCostAllocationResourceArrayOutput {
	return o
}

func (o SourceCostAllocationResourceArrayOutput) Index(i pulumi.IntInput) SourceCostAllocationResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SourceCostAllocationResource {
		return vs[0].([]SourceCostAllocationResource)[vs[1].(int)]
	}).(SourceCostAllocationResourceOutput)
}

type SourceCostAllocationResourceResponse struct {
	Name         string   `pulumi:"name"`
	ResourceType string   `pulumi:"resourceType"`
	Values       []string `pulumi:"values"`
}

type SourceCostAllocationResourceResponseOutput struct{ *pulumi.OutputState }

func (SourceCostAllocationResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceCostAllocationResourceResponse)(nil)).Elem()
}

func (o SourceCostAllocationResourceResponseOutput) ToSourceCostAllocationResourceResponseOutput() SourceCostAllocationResourceResponseOutput {
	return o
}

func (o SourceCostAllocationResourceResponseOutput) ToSourceCostAllocationResourceResponseOutputWithContext(ctx context.Context) SourceCostAllocationResourceResponseOutput {
	return o
}

func (o SourceCostAllocationResourceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SourceCostAllocationResourceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SourceCostAllocationResourceResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v SourceCostAllocationResourceResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o SourceCostAllocationResourceResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SourceCostAllocationResourceResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type SourceCostAllocationResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SourceCostAllocationResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SourceCostAllocationResourceResponse)(nil)).Elem()
}

func (o SourceCostAllocationResourceResponseArrayOutput) ToSourceCostAllocationResourceResponseArrayOutput() SourceCostAllocationResourceResponseArrayOutput {
	return o
}

func (o SourceCostAllocationResourceResponseArrayOutput) ToSourceCostAllocationResourceResponseArrayOutputWithContext(ctx context.Context) SourceCostAllocationResourceResponseArrayOutput {
	return o
}

func (o SourceCostAllocationResourceResponseArrayOutput) Index(i pulumi.IntInput) SourceCostAllocationResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SourceCostAllocationResourceResponse {
		return vs[0].([]SourceCostAllocationResourceResponse)[vs[1].(int)]
	}).(SourceCostAllocationResourceResponseOutput)
}

type TargetCostAllocationResource struct {
	Name         string                     `pulumi:"name"`
	PolicyType   string                     `pulumi:"policyType"`
	ResourceType string                     `pulumi:"resourceType"`
	Values       []CostAllocationProportion `pulumi:"values"`
}





type TargetCostAllocationResourceInput interface {
	pulumi.Input

	ToTargetCostAllocationResourceOutput() TargetCostAllocationResourceOutput
	ToTargetCostAllocationResourceOutputWithContext(context.Context) TargetCostAllocationResourceOutput
}

type TargetCostAllocationResourceArgs struct {
	Name         pulumi.StringInput                 `pulumi:"name"`
	PolicyType   pulumi.StringInput                 `pulumi:"policyType"`
	ResourceType pulumi.StringInput                 `pulumi:"resourceType"`
	Values       CostAllocationProportionArrayInput `pulumi:"values"`
}

func (TargetCostAllocationResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetCostAllocationResource)(nil)).Elem()
}

func (i TargetCostAllocationResourceArgs) ToTargetCostAllocationResourceOutput() TargetCostAllocationResourceOutput {
	return i.ToTargetCostAllocationResourceOutputWithContext(context.Background())
}

func (i TargetCostAllocationResourceArgs) ToTargetCostAllocationResourceOutputWithContext(ctx context.Context) TargetCostAllocationResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetCostAllocationResourceOutput)
}





type TargetCostAllocationResourceArrayInput interface {
	pulumi.Input

	ToTargetCostAllocationResourceArrayOutput() TargetCostAllocationResourceArrayOutput
	ToTargetCostAllocationResourceArrayOutputWithContext(context.Context) TargetCostAllocationResourceArrayOutput
}

type TargetCostAllocationResourceArray []TargetCostAllocationResourceInput

func (TargetCostAllocationResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetCostAllocationResource)(nil)).Elem()
}

func (i TargetCostAllocationResourceArray) ToTargetCostAllocationResourceArrayOutput() TargetCostAllocationResourceArrayOutput {
	return i.ToTargetCostAllocationResourceArrayOutputWithContext(context.Background())
}

func (i TargetCostAllocationResourceArray) ToTargetCostAllocationResourceArrayOutputWithContext(ctx context.Context) TargetCostAllocationResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetCostAllocationResourceArrayOutput)
}

type TargetCostAllocationResourceOutput struct{ *pulumi.OutputState }

func (TargetCostAllocationResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetCostAllocationResource)(nil)).Elem()
}

func (o TargetCostAllocationResourceOutput) ToTargetCostAllocationResourceOutput() TargetCostAllocationResourceOutput {
	return o
}

func (o TargetCostAllocationResourceOutput) ToTargetCostAllocationResourceOutputWithContext(ctx context.Context) TargetCostAllocationResourceOutput {
	return o
}

func (o TargetCostAllocationResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TargetCostAllocationResource) string { return v.Name }).(pulumi.StringOutput)
}

func (o TargetCostAllocationResourceOutput) PolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v TargetCostAllocationResource) string { return v.PolicyType }).(pulumi.StringOutput)
}

func (o TargetCostAllocationResourceOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v TargetCostAllocationResource) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o TargetCostAllocationResourceOutput) Values() CostAllocationProportionArrayOutput {
	return o.ApplyT(func(v TargetCostAllocationResource) []CostAllocationProportion { return v.Values }).(CostAllocationProportionArrayOutput)
}

type TargetCostAllocationResourceArrayOutput struct{ *pulumi.OutputState }

func (TargetCostAllocationResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetCostAllocationResource)(nil)).Elem()
}

func (o TargetCostAllocationResourceArrayOutput) ToTargetCostAllocationResourceArrayOutput() TargetCostAllocationResourceArrayOutput {
	return o
}

func (o TargetCostAllocationResourceArrayOutput) ToTargetCostAllocationResourceArrayOutputWithContext(ctx context.Context) TargetCostAllocationResourceArrayOutput {
	return o
}

func (o TargetCostAllocationResourceArrayOutput) Index(i pulumi.IntInput) TargetCostAllocationResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetCostAllocationResource {
		return vs[0].([]TargetCostAllocationResource)[vs[1].(int)]
	}).(TargetCostAllocationResourceOutput)
}

type TargetCostAllocationResourceResponse struct {
	Name         string                             `pulumi:"name"`
	PolicyType   string                             `pulumi:"policyType"`
	ResourceType string                             `pulumi:"resourceType"`
	Values       []CostAllocationProportionResponse `pulumi:"values"`
}

type TargetCostAllocationResourceResponseOutput struct{ *pulumi.OutputState }

func (TargetCostAllocationResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetCostAllocationResourceResponse)(nil)).Elem()
}

func (o TargetCostAllocationResourceResponseOutput) ToTargetCostAllocationResourceResponseOutput() TargetCostAllocationResourceResponseOutput {
	return o
}

func (o TargetCostAllocationResourceResponseOutput) ToTargetCostAllocationResourceResponseOutputWithContext(ctx context.Context) TargetCostAllocationResourceResponseOutput {
	return o
}

func (o TargetCostAllocationResourceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TargetCostAllocationResourceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o TargetCostAllocationResourceResponseOutput) PolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v TargetCostAllocationResourceResponse) string { return v.PolicyType }).(pulumi.StringOutput)
}

func (o TargetCostAllocationResourceResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v TargetCostAllocationResourceResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o TargetCostAllocationResourceResponseOutput) Values() CostAllocationProportionResponseArrayOutput {
	return o.ApplyT(func(v TargetCostAllocationResourceResponse) []CostAllocationProportionResponse { return v.Values }).(CostAllocationProportionResponseArrayOutput)
}

type TargetCostAllocationResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (TargetCostAllocationResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetCostAllocationResourceResponse)(nil)).Elem()
}

func (o TargetCostAllocationResourceResponseArrayOutput) ToTargetCostAllocationResourceResponseArrayOutput() TargetCostAllocationResourceResponseArrayOutput {
	return o
}

func (o TargetCostAllocationResourceResponseArrayOutput) ToTargetCostAllocationResourceResponseArrayOutputWithContext(ctx context.Context) TargetCostAllocationResourceResponseArrayOutput {
	return o
}

func (o TargetCostAllocationResourceResponseArrayOutput) Index(i pulumi.IntInput) TargetCostAllocationResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetCostAllocationResourceResponse {
		return vs[0].([]TargetCostAllocationResourceResponse)[vs[1].(int)]
	}).(TargetCostAllocationResourceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(CostAllocationProportionOutput{})
	pulumi.RegisterOutputType(CostAllocationProportionArrayOutput{})
	pulumi.RegisterOutputType(CostAllocationProportionResponseOutput{})
	pulumi.RegisterOutputType(CostAllocationProportionResponseArrayOutput{})
	pulumi.RegisterOutputType(CostAllocationRuleDetailsOutput{})
	pulumi.RegisterOutputType(CostAllocationRuleDetailsPtrOutput{})
	pulumi.RegisterOutputType(CostAllocationRuleDetailsResponseOutput{})
	pulumi.RegisterOutputType(CostAllocationRulePropertiesOutput{})
	pulumi.RegisterOutputType(CostAllocationRulePropertiesPtrOutput{})
	pulumi.RegisterOutputType(CostAllocationRulePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SourceCostAllocationResourceOutput{})
	pulumi.RegisterOutputType(SourceCostAllocationResourceArrayOutput{})
	pulumi.RegisterOutputType(SourceCostAllocationResourceResponseOutput{})
	pulumi.RegisterOutputType(SourceCostAllocationResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(TargetCostAllocationResourceOutput{})
	pulumi.RegisterOutputType(TargetCostAllocationResourceArrayOutput{})
	pulumi.RegisterOutputType(TargetCostAllocationResourceResponseOutput{})
	pulumi.RegisterOutputType(TargetCostAllocationResourceResponseArrayOutput{})
}
