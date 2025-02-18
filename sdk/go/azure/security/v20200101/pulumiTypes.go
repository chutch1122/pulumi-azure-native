


package v20200101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdaptiveApplicationControlIssueSummaryResponse struct {
	Issue       *string  `pulumi:"issue"`
	NumberOfVms *float64 `pulumi:"numberOfVms"`
}

type AdaptiveApplicationControlIssueSummaryResponseOutput struct{ *pulumi.OutputState }

func (AdaptiveApplicationControlIssueSummaryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdaptiveApplicationControlIssueSummaryResponse)(nil)).Elem()
}

func (o AdaptiveApplicationControlIssueSummaryResponseOutput) ToAdaptiveApplicationControlIssueSummaryResponseOutput() AdaptiveApplicationControlIssueSummaryResponseOutput {
	return o
}

func (o AdaptiveApplicationControlIssueSummaryResponseOutput) ToAdaptiveApplicationControlIssueSummaryResponseOutputWithContext(ctx context.Context) AdaptiveApplicationControlIssueSummaryResponseOutput {
	return o
}

func (o AdaptiveApplicationControlIssueSummaryResponseOutput) Issue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdaptiveApplicationControlIssueSummaryResponse) *string { return v.Issue }).(pulumi.StringPtrOutput)
}

func (o AdaptiveApplicationControlIssueSummaryResponseOutput) NumberOfVms() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AdaptiveApplicationControlIssueSummaryResponse) *float64 { return v.NumberOfVms }).(pulumi.Float64PtrOutput)
}

type AdaptiveApplicationControlIssueSummaryResponseArrayOutput struct{ *pulumi.OutputState }

func (AdaptiveApplicationControlIssueSummaryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdaptiveApplicationControlIssueSummaryResponse)(nil)).Elem()
}

func (o AdaptiveApplicationControlIssueSummaryResponseArrayOutput) ToAdaptiveApplicationControlIssueSummaryResponseArrayOutput() AdaptiveApplicationControlIssueSummaryResponseArrayOutput {
	return o
}

func (o AdaptiveApplicationControlIssueSummaryResponseArrayOutput) ToAdaptiveApplicationControlIssueSummaryResponseArrayOutputWithContext(ctx context.Context) AdaptiveApplicationControlIssueSummaryResponseArrayOutput {
	return o
}

func (o AdaptiveApplicationControlIssueSummaryResponseArrayOutput) Index(i pulumi.IntInput) AdaptiveApplicationControlIssueSummaryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdaptiveApplicationControlIssueSummaryResponse {
		return vs[0].([]AdaptiveApplicationControlIssueSummaryResponse)[vs[1].(int)]
	}).(AdaptiveApplicationControlIssueSummaryResponseOutput)
}

type AssessmentLinksResponse struct {
	AzurePortalUri string `pulumi:"azurePortalUri"`
}

type AssessmentLinksResponseOutput struct{ *pulumi.OutputState }

func (AssessmentLinksResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentLinksResponse)(nil)).Elem()
}

func (o AssessmentLinksResponseOutput) ToAssessmentLinksResponseOutput() AssessmentLinksResponseOutput {
	return o
}

func (o AssessmentLinksResponseOutput) ToAssessmentLinksResponseOutputWithContext(ctx context.Context) AssessmentLinksResponseOutput {
	return o
}

func (o AssessmentLinksResponseOutput) AzurePortalUri() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentLinksResponse) string { return v.AzurePortalUri }).(pulumi.StringOutput)
}

type AssessmentStatus struct {
	Cause       *string `pulumi:"cause"`
	Code        string  `pulumi:"code"`
	Description *string `pulumi:"description"`
}





type AssessmentStatusInput interface {
	pulumi.Input

	ToAssessmentStatusOutput() AssessmentStatusOutput
	ToAssessmentStatusOutputWithContext(context.Context) AssessmentStatusOutput
}

type AssessmentStatusArgs struct {
	Cause       pulumi.StringPtrInput `pulumi:"cause"`
	Code        pulumi.StringInput    `pulumi:"code"`
	Description pulumi.StringPtrInput `pulumi:"description"`
}

func (AssessmentStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentStatus)(nil)).Elem()
}

func (i AssessmentStatusArgs) ToAssessmentStatusOutput() AssessmentStatusOutput {
	return i.ToAssessmentStatusOutputWithContext(context.Background())
}

func (i AssessmentStatusArgs) ToAssessmentStatusOutputWithContext(ctx context.Context) AssessmentStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentStatusOutput)
}

type AssessmentStatusOutput struct{ *pulumi.OutputState }

func (AssessmentStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentStatus)(nil)).Elem()
}

func (o AssessmentStatusOutput) ToAssessmentStatusOutput() AssessmentStatusOutput {
	return o
}

func (o AssessmentStatusOutput) ToAssessmentStatusOutputWithContext(ctx context.Context) AssessmentStatusOutput {
	return o
}

func (o AssessmentStatusOutput) Cause() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssessmentStatus) *string { return v.Cause }).(pulumi.StringPtrOutput)
}

func (o AssessmentStatusOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentStatus) string { return v.Code }).(pulumi.StringOutput)
}

func (o AssessmentStatusOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssessmentStatus) *string { return v.Description }).(pulumi.StringPtrOutput)
}

type AssessmentStatusResponse struct {
	Cause       *string `pulumi:"cause"`
	Code        string  `pulumi:"code"`
	Description *string `pulumi:"description"`
}

type AssessmentStatusResponseOutput struct{ *pulumi.OutputState }

func (AssessmentStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentStatusResponse)(nil)).Elem()
}

func (o AssessmentStatusResponseOutput) ToAssessmentStatusResponseOutput() AssessmentStatusResponseOutput {
	return o
}

func (o AssessmentStatusResponseOutput) ToAssessmentStatusResponseOutputWithContext(ctx context.Context) AssessmentStatusResponseOutput {
	return o
}

func (o AssessmentStatusResponseOutput) Cause() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssessmentStatusResponse) *string { return v.Cause }).(pulumi.StringPtrOutput)
}

func (o AssessmentStatusResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentStatusResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o AssessmentStatusResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssessmentStatusResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

type AzureResourceDetails struct {
	Source string `pulumi:"source"`
}

type AzureResourceDetailsResponse struct {
	Id     string `pulumi:"id"`
	Source string `pulumi:"source"`
}

type JitNetworkAccessPolicyVirtualMachine struct {
	Id              string                     `pulumi:"id"`
	Ports           []JitNetworkAccessPortRule `pulumi:"ports"`
	PublicIpAddress *string                    `pulumi:"publicIpAddress"`
}





type JitNetworkAccessPolicyVirtualMachineInput interface {
	pulumi.Input

	ToJitNetworkAccessPolicyVirtualMachineOutput() JitNetworkAccessPolicyVirtualMachineOutput
	ToJitNetworkAccessPolicyVirtualMachineOutputWithContext(context.Context) JitNetworkAccessPolicyVirtualMachineOutput
}

type JitNetworkAccessPolicyVirtualMachineArgs struct {
	Id              pulumi.StringInput                 `pulumi:"id"`
	Ports           JitNetworkAccessPortRuleArrayInput `pulumi:"ports"`
	PublicIpAddress pulumi.StringPtrInput              `pulumi:"publicIpAddress"`
}

func (JitNetworkAccessPolicyVirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessPolicyVirtualMachine)(nil)).Elem()
}

func (i JitNetworkAccessPolicyVirtualMachineArgs) ToJitNetworkAccessPolicyVirtualMachineOutput() JitNetworkAccessPolicyVirtualMachineOutput {
	return i.ToJitNetworkAccessPolicyVirtualMachineOutputWithContext(context.Background())
}

func (i JitNetworkAccessPolicyVirtualMachineArgs) ToJitNetworkAccessPolicyVirtualMachineOutputWithContext(ctx context.Context) JitNetworkAccessPolicyVirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessPolicyVirtualMachineOutput)
}





type JitNetworkAccessPolicyVirtualMachineArrayInput interface {
	pulumi.Input

	ToJitNetworkAccessPolicyVirtualMachineArrayOutput() JitNetworkAccessPolicyVirtualMachineArrayOutput
	ToJitNetworkAccessPolicyVirtualMachineArrayOutputWithContext(context.Context) JitNetworkAccessPolicyVirtualMachineArrayOutput
}

type JitNetworkAccessPolicyVirtualMachineArray []JitNetworkAccessPolicyVirtualMachineInput

func (JitNetworkAccessPolicyVirtualMachineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessPolicyVirtualMachine)(nil)).Elem()
}

func (i JitNetworkAccessPolicyVirtualMachineArray) ToJitNetworkAccessPolicyVirtualMachineArrayOutput() JitNetworkAccessPolicyVirtualMachineArrayOutput {
	return i.ToJitNetworkAccessPolicyVirtualMachineArrayOutputWithContext(context.Background())
}

func (i JitNetworkAccessPolicyVirtualMachineArray) ToJitNetworkAccessPolicyVirtualMachineArrayOutputWithContext(ctx context.Context) JitNetworkAccessPolicyVirtualMachineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessPolicyVirtualMachineArrayOutput)
}

type JitNetworkAccessPolicyVirtualMachineOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessPolicyVirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessPolicyVirtualMachine)(nil)).Elem()
}

func (o JitNetworkAccessPolicyVirtualMachineOutput) ToJitNetworkAccessPolicyVirtualMachineOutput() JitNetworkAccessPolicyVirtualMachineOutput {
	return o
}

func (o JitNetworkAccessPolicyVirtualMachineOutput) ToJitNetworkAccessPolicyVirtualMachineOutputWithContext(ctx context.Context) JitNetworkAccessPolicyVirtualMachineOutput {
	return o
}

func (o JitNetworkAccessPolicyVirtualMachineOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessPolicyVirtualMachine) string { return v.Id }).(pulumi.StringOutput)
}

func (o JitNetworkAccessPolicyVirtualMachineOutput) Ports() JitNetworkAccessPortRuleArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessPolicyVirtualMachine) []JitNetworkAccessPortRule { return v.Ports }).(JitNetworkAccessPortRuleArrayOutput)
}

func (o JitNetworkAccessPolicyVirtualMachineOutput) PublicIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessPolicyVirtualMachine) *string { return v.PublicIpAddress }).(pulumi.StringPtrOutput)
}

type JitNetworkAccessPolicyVirtualMachineArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessPolicyVirtualMachineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessPolicyVirtualMachine)(nil)).Elem()
}

func (o JitNetworkAccessPolicyVirtualMachineArrayOutput) ToJitNetworkAccessPolicyVirtualMachineArrayOutput() JitNetworkAccessPolicyVirtualMachineArrayOutput {
	return o
}

func (o JitNetworkAccessPolicyVirtualMachineArrayOutput) ToJitNetworkAccessPolicyVirtualMachineArrayOutputWithContext(ctx context.Context) JitNetworkAccessPolicyVirtualMachineArrayOutput {
	return o
}

func (o JitNetworkAccessPolicyVirtualMachineArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessPolicyVirtualMachineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessPolicyVirtualMachine {
		return vs[0].([]JitNetworkAccessPolicyVirtualMachine)[vs[1].(int)]
	}).(JitNetworkAccessPolicyVirtualMachineOutput)
}

type JitNetworkAccessPolicyVirtualMachineResponse struct {
	Id              string                             `pulumi:"id"`
	Ports           []JitNetworkAccessPortRuleResponse `pulumi:"ports"`
	PublicIpAddress *string                            `pulumi:"publicIpAddress"`
}

type JitNetworkAccessPolicyVirtualMachineResponseOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessPolicyVirtualMachineResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessPolicyVirtualMachineResponse)(nil)).Elem()
}

func (o JitNetworkAccessPolicyVirtualMachineResponseOutput) ToJitNetworkAccessPolicyVirtualMachineResponseOutput() JitNetworkAccessPolicyVirtualMachineResponseOutput {
	return o
}

func (o JitNetworkAccessPolicyVirtualMachineResponseOutput) ToJitNetworkAccessPolicyVirtualMachineResponseOutputWithContext(ctx context.Context) JitNetworkAccessPolicyVirtualMachineResponseOutput {
	return o
}

func (o JitNetworkAccessPolicyVirtualMachineResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessPolicyVirtualMachineResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o JitNetworkAccessPolicyVirtualMachineResponseOutput) Ports() JitNetworkAccessPortRuleResponseArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessPolicyVirtualMachineResponse) []JitNetworkAccessPortRuleResponse {
		return v.Ports
	}).(JitNetworkAccessPortRuleResponseArrayOutput)
}

func (o JitNetworkAccessPolicyVirtualMachineResponseOutput) PublicIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessPolicyVirtualMachineResponse) *string { return v.PublicIpAddress }).(pulumi.StringPtrOutput)
}

type JitNetworkAccessPolicyVirtualMachineResponseArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessPolicyVirtualMachineResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessPolicyVirtualMachineResponse)(nil)).Elem()
}

func (o JitNetworkAccessPolicyVirtualMachineResponseArrayOutput) ToJitNetworkAccessPolicyVirtualMachineResponseArrayOutput() JitNetworkAccessPolicyVirtualMachineResponseArrayOutput {
	return o
}

func (o JitNetworkAccessPolicyVirtualMachineResponseArrayOutput) ToJitNetworkAccessPolicyVirtualMachineResponseArrayOutputWithContext(ctx context.Context) JitNetworkAccessPolicyVirtualMachineResponseArrayOutput {
	return o
}

func (o JitNetworkAccessPolicyVirtualMachineResponseArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessPolicyVirtualMachineResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessPolicyVirtualMachineResponse {
		return vs[0].([]JitNetworkAccessPolicyVirtualMachineResponse)[vs[1].(int)]
	}).(JitNetworkAccessPolicyVirtualMachineResponseOutput)
}

type JitNetworkAccessPortRule struct {
	AllowedSourceAddressPrefix   *string  `pulumi:"allowedSourceAddressPrefix"`
	AllowedSourceAddressPrefixes []string `pulumi:"allowedSourceAddressPrefixes"`
	MaxRequestAccessDuration     string   `pulumi:"maxRequestAccessDuration"`
	Number                       int      `pulumi:"number"`
	Protocol                     string   `pulumi:"protocol"`
}





type JitNetworkAccessPortRuleInput interface {
	pulumi.Input

	ToJitNetworkAccessPortRuleOutput() JitNetworkAccessPortRuleOutput
	ToJitNetworkAccessPortRuleOutputWithContext(context.Context) JitNetworkAccessPortRuleOutput
}

type JitNetworkAccessPortRuleArgs struct {
	AllowedSourceAddressPrefix   pulumi.StringPtrInput   `pulumi:"allowedSourceAddressPrefix"`
	AllowedSourceAddressPrefixes pulumi.StringArrayInput `pulumi:"allowedSourceAddressPrefixes"`
	MaxRequestAccessDuration     pulumi.StringInput      `pulumi:"maxRequestAccessDuration"`
	Number                       pulumi.IntInput         `pulumi:"number"`
	Protocol                     pulumi.StringInput      `pulumi:"protocol"`
}

func (JitNetworkAccessPortRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessPortRule)(nil)).Elem()
}

func (i JitNetworkAccessPortRuleArgs) ToJitNetworkAccessPortRuleOutput() JitNetworkAccessPortRuleOutput {
	return i.ToJitNetworkAccessPortRuleOutputWithContext(context.Background())
}

func (i JitNetworkAccessPortRuleArgs) ToJitNetworkAccessPortRuleOutputWithContext(ctx context.Context) JitNetworkAccessPortRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessPortRuleOutput)
}





type JitNetworkAccessPortRuleArrayInput interface {
	pulumi.Input

	ToJitNetworkAccessPortRuleArrayOutput() JitNetworkAccessPortRuleArrayOutput
	ToJitNetworkAccessPortRuleArrayOutputWithContext(context.Context) JitNetworkAccessPortRuleArrayOutput
}

type JitNetworkAccessPortRuleArray []JitNetworkAccessPortRuleInput

func (JitNetworkAccessPortRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessPortRule)(nil)).Elem()
}

func (i JitNetworkAccessPortRuleArray) ToJitNetworkAccessPortRuleArrayOutput() JitNetworkAccessPortRuleArrayOutput {
	return i.ToJitNetworkAccessPortRuleArrayOutputWithContext(context.Background())
}

func (i JitNetworkAccessPortRuleArray) ToJitNetworkAccessPortRuleArrayOutputWithContext(ctx context.Context) JitNetworkAccessPortRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessPortRuleArrayOutput)
}

type JitNetworkAccessPortRuleOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessPortRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessPortRule)(nil)).Elem()
}

func (o JitNetworkAccessPortRuleOutput) ToJitNetworkAccessPortRuleOutput() JitNetworkAccessPortRuleOutput {
	return o
}

func (o JitNetworkAccessPortRuleOutput) ToJitNetworkAccessPortRuleOutputWithContext(ctx context.Context) JitNetworkAccessPortRuleOutput {
	return o
}

func (o JitNetworkAccessPortRuleOutput) AllowedSourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRule) *string { return v.AllowedSourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o JitNetworkAccessPortRuleOutput) AllowedSourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRule) []string { return v.AllowedSourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o JitNetworkAccessPortRuleOutput) MaxRequestAccessDuration() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRule) string { return v.MaxRequestAccessDuration }).(pulumi.StringOutput)
}

func (o JitNetworkAccessPortRuleOutput) Number() pulumi.IntOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRule) int { return v.Number }).(pulumi.IntOutput)
}

func (o JitNetworkAccessPortRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRule) string { return v.Protocol }).(pulumi.StringOutput)
}

type JitNetworkAccessPortRuleArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessPortRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessPortRule)(nil)).Elem()
}

func (o JitNetworkAccessPortRuleArrayOutput) ToJitNetworkAccessPortRuleArrayOutput() JitNetworkAccessPortRuleArrayOutput {
	return o
}

func (o JitNetworkAccessPortRuleArrayOutput) ToJitNetworkAccessPortRuleArrayOutputWithContext(ctx context.Context) JitNetworkAccessPortRuleArrayOutput {
	return o
}

func (o JitNetworkAccessPortRuleArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessPortRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessPortRule {
		return vs[0].([]JitNetworkAccessPortRule)[vs[1].(int)]
	}).(JitNetworkAccessPortRuleOutput)
}

type JitNetworkAccessPortRuleResponse struct {
	AllowedSourceAddressPrefix   *string  `pulumi:"allowedSourceAddressPrefix"`
	AllowedSourceAddressPrefixes []string `pulumi:"allowedSourceAddressPrefixes"`
	MaxRequestAccessDuration     string   `pulumi:"maxRequestAccessDuration"`
	Number                       int      `pulumi:"number"`
	Protocol                     string   `pulumi:"protocol"`
}

type JitNetworkAccessPortRuleResponseOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessPortRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessPortRuleResponse)(nil)).Elem()
}

func (o JitNetworkAccessPortRuleResponseOutput) ToJitNetworkAccessPortRuleResponseOutput() JitNetworkAccessPortRuleResponseOutput {
	return o
}

func (o JitNetworkAccessPortRuleResponseOutput) ToJitNetworkAccessPortRuleResponseOutputWithContext(ctx context.Context) JitNetworkAccessPortRuleResponseOutput {
	return o
}

func (o JitNetworkAccessPortRuleResponseOutput) AllowedSourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRuleResponse) *string { return v.AllowedSourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o JitNetworkAccessPortRuleResponseOutput) AllowedSourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRuleResponse) []string { return v.AllowedSourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o JitNetworkAccessPortRuleResponseOutput) MaxRequestAccessDuration() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRuleResponse) string { return v.MaxRequestAccessDuration }).(pulumi.StringOutput)
}

func (o JitNetworkAccessPortRuleResponseOutput) Number() pulumi.IntOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRuleResponse) int { return v.Number }).(pulumi.IntOutput)
}

func (o JitNetworkAccessPortRuleResponseOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessPortRuleResponse) string { return v.Protocol }).(pulumi.StringOutput)
}

type JitNetworkAccessPortRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessPortRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessPortRuleResponse)(nil)).Elem()
}

func (o JitNetworkAccessPortRuleResponseArrayOutput) ToJitNetworkAccessPortRuleResponseArrayOutput() JitNetworkAccessPortRuleResponseArrayOutput {
	return o
}

func (o JitNetworkAccessPortRuleResponseArrayOutput) ToJitNetworkAccessPortRuleResponseArrayOutputWithContext(ctx context.Context) JitNetworkAccessPortRuleResponseArrayOutput {
	return o
}

func (o JitNetworkAccessPortRuleResponseArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessPortRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessPortRuleResponse {
		return vs[0].([]JitNetworkAccessPortRuleResponse)[vs[1].(int)]
	}).(JitNetworkAccessPortRuleResponseOutput)
}

type JitNetworkAccessRequest struct {
	Justification   *string                                 `pulumi:"justification"`
	Requestor       string                                  `pulumi:"requestor"`
	StartTimeUtc    string                                  `pulumi:"startTimeUtc"`
	VirtualMachines []JitNetworkAccessRequestVirtualMachine `pulumi:"virtualMachines"`
}





type JitNetworkAccessRequestInput interface {
	pulumi.Input

	ToJitNetworkAccessRequestOutput() JitNetworkAccessRequestOutput
	ToJitNetworkAccessRequestOutputWithContext(context.Context) JitNetworkAccessRequestOutput
}

type JitNetworkAccessRequestArgs struct {
	Justification   pulumi.StringPtrInput                           `pulumi:"justification"`
	Requestor       pulumi.StringInput                              `pulumi:"requestor"`
	StartTimeUtc    pulumi.StringInput                              `pulumi:"startTimeUtc"`
	VirtualMachines JitNetworkAccessRequestVirtualMachineArrayInput `pulumi:"virtualMachines"`
}

func (JitNetworkAccessRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequest)(nil)).Elem()
}

func (i JitNetworkAccessRequestArgs) ToJitNetworkAccessRequestOutput() JitNetworkAccessRequestOutput {
	return i.ToJitNetworkAccessRequestOutputWithContext(context.Background())
}

func (i JitNetworkAccessRequestArgs) ToJitNetworkAccessRequestOutputWithContext(ctx context.Context) JitNetworkAccessRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessRequestOutput)
}





type JitNetworkAccessRequestArrayInput interface {
	pulumi.Input

	ToJitNetworkAccessRequestArrayOutput() JitNetworkAccessRequestArrayOutput
	ToJitNetworkAccessRequestArrayOutputWithContext(context.Context) JitNetworkAccessRequestArrayOutput
}

type JitNetworkAccessRequestArray []JitNetworkAccessRequestInput

func (JitNetworkAccessRequestArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequest)(nil)).Elem()
}

func (i JitNetworkAccessRequestArray) ToJitNetworkAccessRequestArrayOutput() JitNetworkAccessRequestArrayOutput {
	return i.ToJitNetworkAccessRequestArrayOutputWithContext(context.Background())
}

func (i JitNetworkAccessRequestArray) ToJitNetworkAccessRequestArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessRequestArrayOutput)
}

type JitNetworkAccessRequestOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequest)(nil)).Elem()
}

func (o JitNetworkAccessRequestOutput) ToJitNetworkAccessRequestOutput() JitNetworkAccessRequestOutput {
	return o
}

func (o JitNetworkAccessRequestOutput) ToJitNetworkAccessRequestOutputWithContext(ctx context.Context) JitNetworkAccessRequestOutput {
	return o
}

func (o JitNetworkAccessRequestOutput) Justification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessRequest) *string { return v.Justification }).(pulumi.StringPtrOutput)
}

func (o JitNetworkAccessRequestOutput) Requestor() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequest) string { return v.Requestor }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestOutput) StartTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequest) string { return v.StartTimeUtc }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestOutput) VirtualMachines() JitNetworkAccessRequestVirtualMachineArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessRequest) []JitNetworkAccessRequestVirtualMachine { return v.VirtualMachines }).(JitNetworkAccessRequestVirtualMachineArrayOutput)
}

type JitNetworkAccessRequestArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequest)(nil)).Elem()
}

func (o JitNetworkAccessRequestArrayOutput) ToJitNetworkAccessRequestArrayOutput() JitNetworkAccessRequestArrayOutput {
	return o
}

func (o JitNetworkAccessRequestArrayOutput) ToJitNetworkAccessRequestArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestArrayOutput {
	return o
}

func (o JitNetworkAccessRequestArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessRequestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessRequest {
		return vs[0].([]JitNetworkAccessRequest)[vs[1].(int)]
	}).(JitNetworkAccessRequestOutput)
}

type JitNetworkAccessRequestPort struct {
	AllowedSourceAddressPrefix   *string  `pulumi:"allowedSourceAddressPrefix"`
	AllowedSourceAddressPrefixes []string `pulumi:"allowedSourceAddressPrefixes"`
	EndTimeUtc                   string   `pulumi:"endTimeUtc"`
	MappedPort                   *int     `pulumi:"mappedPort"`
	Number                       int      `pulumi:"number"`
	Status                       string   `pulumi:"status"`
	StatusReason                 string   `pulumi:"statusReason"`
}





type JitNetworkAccessRequestPortInput interface {
	pulumi.Input

	ToJitNetworkAccessRequestPortOutput() JitNetworkAccessRequestPortOutput
	ToJitNetworkAccessRequestPortOutputWithContext(context.Context) JitNetworkAccessRequestPortOutput
}

type JitNetworkAccessRequestPortArgs struct {
	AllowedSourceAddressPrefix   pulumi.StringPtrInput   `pulumi:"allowedSourceAddressPrefix"`
	AllowedSourceAddressPrefixes pulumi.StringArrayInput `pulumi:"allowedSourceAddressPrefixes"`
	EndTimeUtc                   pulumi.StringInput      `pulumi:"endTimeUtc"`
	MappedPort                   pulumi.IntPtrInput      `pulumi:"mappedPort"`
	Number                       pulumi.IntInput         `pulumi:"number"`
	Status                       pulumi.StringInput      `pulumi:"status"`
	StatusReason                 pulumi.StringInput      `pulumi:"statusReason"`
}

func (JitNetworkAccessRequestPortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequestPort)(nil)).Elem()
}

func (i JitNetworkAccessRequestPortArgs) ToJitNetworkAccessRequestPortOutput() JitNetworkAccessRequestPortOutput {
	return i.ToJitNetworkAccessRequestPortOutputWithContext(context.Background())
}

func (i JitNetworkAccessRequestPortArgs) ToJitNetworkAccessRequestPortOutputWithContext(ctx context.Context) JitNetworkAccessRequestPortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessRequestPortOutput)
}





type JitNetworkAccessRequestPortArrayInput interface {
	pulumi.Input

	ToJitNetworkAccessRequestPortArrayOutput() JitNetworkAccessRequestPortArrayOutput
	ToJitNetworkAccessRequestPortArrayOutputWithContext(context.Context) JitNetworkAccessRequestPortArrayOutput
}

type JitNetworkAccessRequestPortArray []JitNetworkAccessRequestPortInput

func (JitNetworkAccessRequestPortArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequestPort)(nil)).Elem()
}

func (i JitNetworkAccessRequestPortArray) ToJitNetworkAccessRequestPortArrayOutput() JitNetworkAccessRequestPortArrayOutput {
	return i.ToJitNetworkAccessRequestPortArrayOutputWithContext(context.Background())
}

func (i JitNetworkAccessRequestPortArray) ToJitNetworkAccessRequestPortArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestPortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessRequestPortArrayOutput)
}

type JitNetworkAccessRequestPortOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestPortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequestPort)(nil)).Elem()
}

func (o JitNetworkAccessRequestPortOutput) ToJitNetworkAccessRequestPortOutput() JitNetworkAccessRequestPortOutput {
	return o
}

func (o JitNetworkAccessRequestPortOutput) ToJitNetworkAccessRequestPortOutputWithContext(ctx context.Context) JitNetworkAccessRequestPortOutput {
	return o
}

func (o JitNetworkAccessRequestPortOutput) AllowedSourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPort) *string { return v.AllowedSourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o JitNetworkAccessRequestPortOutput) AllowedSourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPort) []string { return v.AllowedSourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o JitNetworkAccessRequestPortOutput) EndTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPort) string { return v.EndTimeUtc }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestPortOutput) MappedPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPort) *int { return v.MappedPort }).(pulumi.IntPtrOutput)
}

func (o JitNetworkAccessRequestPortOutput) Number() pulumi.IntOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPort) int { return v.Number }).(pulumi.IntOutput)
}

func (o JitNetworkAccessRequestPortOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPort) string { return v.Status }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestPortOutput) StatusReason() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPort) string { return v.StatusReason }).(pulumi.StringOutput)
}

type JitNetworkAccessRequestPortArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestPortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequestPort)(nil)).Elem()
}

func (o JitNetworkAccessRequestPortArrayOutput) ToJitNetworkAccessRequestPortArrayOutput() JitNetworkAccessRequestPortArrayOutput {
	return o
}

func (o JitNetworkAccessRequestPortArrayOutput) ToJitNetworkAccessRequestPortArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestPortArrayOutput {
	return o
}

func (o JitNetworkAccessRequestPortArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessRequestPortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessRequestPort {
		return vs[0].([]JitNetworkAccessRequestPort)[vs[1].(int)]
	}).(JitNetworkAccessRequestPortOutput)
}

type JitNetworkAccessRequestPortResponse struct {
	AllowedSourceAddressPrefix   *string  `pulumi:"allowedSourceAddressPrefix"`
	AllowedSourceAddressPrefixes []string `pulumi:"allowedSourceAddressPrefixes"`
	EndTimeUtc                   string   `pulumi:"endTimeUtc"`
	MappedPort                   *int     `pulumi:"mappedPort"`
	Number                       int      `pulumi:"number"`
	Status                       string   `pulumi:"status"`
	StatusReason                 string   `pulumi:"statusReason"`
}

type JitNetworkAccessRequestPortResponseOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestPortResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequestPortResponse)(nil)).Elem()
}

func (o JitNetworkAccessRequestPortResponseOutput) ToJitNetworkAccessRequestPortResponseOutput() JitNetworkAccessRequestPortResponseOutput {
	return o
}

func (o JitNetworkAccessRequestPortResponseOutput) ToJitNetworkAccessRequestPortResponseOutputWithContext(ctx context.Context) JitNetworkAccessRequestPortResponseOutput {
	return o
}

func (o JitNetworkAccessRequestPortResponseOutput) AllowedSourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPortResponse) *string { return v.AllowedSourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o JitNetworkAccessRequestPortResponseOutput) AllowedSourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPortResponse) []string { return v.AllowedSourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

func (o JitNetworkAccessRequestPortResponseOutput) EndTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPortResponse) string { return v.EndTimeUtc }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestPortResponseOutput) MappedPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPortResponse) *int { return v.MappedPort }).(pulumi.IntPtrOutput)
}

func (o JitNetworkAccessRequestPortResponseOutput) Number() pulumi.IntOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPortResponse) int { return v.Number }).(pulumi.IntOutput)
}

func (o JitNetworkAccessRequestPortResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPortResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestPortResponseOutput) StatusReason() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestPortResponse) string { return v.StatusReason }).(pulumi.StringOutput)
}

type JitNetworkAccessRequestPortResponseArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestPortResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequestPortResponse)(nil)).Elem()
}

func (o JitNetworkAccessRequestPortResponseArrayOutput) ToJitNetworkAccessRequestPortResponseArrayOutput() JitNetworkAccessRequestPortResponseArrayOutput {
	return o
}

func (o JitNetworkAccessRequestPortResponseArrayOutput) ToJitNetworkAccessRequestPortResponseArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestPortResponseArrayOutput {
	return o
}

func (o JitNetworkAccessRequestPortResponseArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessRequestPortResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessRequestPortResponse {
		return vs[0].([]JitNetworkAccessRequestPortResponse)[vs[1].(int)]
	}).(JitNetworkAccessRequestPortResponseOutput)
}

type JitNetworkAccessRequestResponse struct {
	Justification   *string                                         `pulumi:"justification"`
	Requestor       string                                          `pulumi:"requestor"`
	StartTimeUtc    string                                          `pulumi:"startTimeUtc"`
	VirtualMachines []JitNetworkAccessRequestVirtualMachineResponse `pulumi:"virtualMachines"`
}

type JitNetworkAccessRequestResponseOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequestResponse)(nil)).Elem()
}

func (o JitNetworkAccessRequestResponseOutput) ToJitNetworkAccessRequestResponseOutput() JitNetworkAccessRequestResponseOutput {
	return o
}

func (o JitNetworkAccessRequestResponseOutput) ToJitNetworkAccessRequestResponseOutputWithContext(ctx context.Context) JitNetworkAccessRequestResponseOutput {
	return o
}

func (o JitNetworkAccessRequestResponseOutput) Justification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestResponse) *string { return v.Justification }).(pulumi.StringPtrOutput)
}

func (o JitNetworkAccessRequestResponseOutput) Requestor() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestResponse) string { return v.Requestor }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestResponseOutput) StartTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestResponse) string { return v.StartTimeUtc }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestResponseOutput) VirtualMachines() JitNetworkAccessRequestVirtualMachineResponseArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestResponse) []JitNetworkAccessRequestVirtualMachineResponse {
		return v.VirtualMachines
	}).(JitNetworkAccessRequestVirtualMachineResponseArrayOutput)
}

type JitNetworkAccessRequestResponseArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequestResponse)(nil)).Elem()
}

func (o JitNetworkAccessRequestResponseArrayOutput) ToJitNetworkAccessRequestResponseArrayOutput() JitNetworkAccessRequestResponseArrayOutput {
	return o
}

func (o JitNetworkAccessRequestResponseArrayOutput) ToJitNetworkAccessRequestResponseArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestResponseArrayOutput {
	return o
}

func (o JitNetworkAccessRequestResponseArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessRequestResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessRequestResponse {
		return vs[0].([]JitNetworkAccessRequestResponse)[vs[1].(int)]
	}).(JitNetworkAccessRequestResponseOutput)
}

type JitNetworkAccessRequestVirtualMachine struct {
	Id    string                        `pulumi:"id"`
	Ports []JitNetworkAccessRequestPort `pulumi:"ports"`
}





type JitNetworkAccessRequestVirtualMachineInput interface {
	pulumi.Input

	ToJitNetworkAccessRequestVirtualMachineOutput() JitNetworkAccessRequestVirtualMachineOutput
	ToJitNetworkAccessRequestVirtualMachineOutputWithContext(context.Context) JitNetworkAccessRequestVirtualMachineOutput
}

type JitNetworkAccessRequestVirtualMachineArgs struct {
	Id    pulumi.StringInput                    `pulumi:"id"`
	Ports JitNetworkAccessRequestPortArrayInput `pulumi:"ports"`
}

func (JitNetworkAccessRequestVirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequestVirtualMachine)(nil)).Elem()
}

func (i JitNetworkAccessRequestVirtualMachineArgs) ToJitNetworkAccessRequestVirtualMachineOutput() JitNetworkAccessRequestVirtualMachineOutput {
	return i.ToJitNetworkAccessRequestVirtualMachineOutputWithContext(context.Background())
}

func (i JitNetworkAccessRequestVirtualMachineArgs) ToJitNetworkAccessRequestVirtualMachineOutputWithContext(ctx context.Context) JitNetworkAccessRequestVirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessRequestVirtualMachineOutput)
}





type JitNetworkAccessRequestVirtualMachineArrayInput interface {
	pulumi.Input

	ToJitNetworkAccessRequestVirtualMachineArrayOutput() JitNetworkAccessRequestVirtualMachineArrayOutput
	ToJitNetworkAccessRequestVirtualMachineArrayOutputWithContext(context.Context) JitNetworkAccessRequestVirtualMachineArrayOutput
}

type JitNetworkAccessRequestVirtualMachineArray []JitNetworkAccessRequestVirtualMachineInput

func (JitNetworkAccessRequestVirtualMachineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequestVirtualMachine)(nil)).Elem()
}

func (i JitNetworkAccessRequestVirtualMachineArray) ToJitNetworkAccessRequestVirtualMachineArrayOutput() JitNetworkAccessRequestVirtualMachineArrayOutput {
	return i.ToJitNetworkAccessRequestVirtualMachineArrayOutputWithContext(context.Background())
}

func (i JitNetworkAccessRequestVirtualMachineArray) ToJitNetworkAccessRequestVirtualMachineArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestVirtualMachineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JitNetworkAccessRequestVirtualMachineArrayOutput)
}

type JitNetworkAccessRequestVirtualMachineOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestVirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequestVirtualMachine)(nil)).Elem()
}

func (o JitNetworkAccessRequestVirtualMachineOutput) ToJitNetworkAccessRequestVirtualMachineOutput() JitNetworkAccessRequestVirtualMachineOutput {
	return o
}

func (o JitNetworkAccessRequestVirtualMachineOutput) ToJitNetworkAccessRequestVirtualMachineOutputWithContext(ctx context.Context) JitNetworkAccessRequestVirtualMachineOutput {
	return o
}

func (o JitNetworkAccessRequestVirtualMachineOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestVirtualMachine) string { return v.Id }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestVirtualMachineOutput) Ports() JitNetworkAccessRequestPortArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestVirtualMachine) []JitNetworkAccessRequestPort { return v.Ports }).(JitNetworkAccessRequestPortArrayOutput)
}

type JitNetworkAccessRequestVirtualMachineArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestVirtualMachineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequestVirtualMachine)(nil)).Elem()
}

func (o JitNetworkAccessRequestVirtualMachineArrayOutput) ToJitNetworkAccessRequestVirtualMachineArrayOutput() JitNetworkAccessRequestVirtualMachineArrayOutput {
	return o
}

func (o JitNetworkAccessRequestVirtualMachineArrayOutput) ToJitNetworkAccessRequestVirtualMachineArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestVirtualMachineArrayOutput {
	return o
}

func (o JitNetworkAccessRequestVirtualMachineArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessRequestVirtualMachineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessRequestVirtualMachine {
		return vs[0].([]JitNetworkAccessRequestVirtualMachine)[vs[1].(int)]
	}).(JitNetworkAccessRequestVirtualMachineOutput)
}

type JitNetworkAccessRequestVirtualMachineResponse struct {
	Id    string                                `pulumi:"id"`
	Ports []JitNetworkAccessRequestPortResponse `pulumi:"ports"`
}

type JitNetworkAccessRequestVirtualMachineResponseOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestVirtualMachineResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JitNetworkAccessRequestVirtualMachineResponse)(nil)).Elem()
}

func (o JitNetworkAccessRequestVirtualMachineResponseOutput) ToJitNetworkAccessRequestVirtualMachineResponseOutput() JitNetworkAccessRequestVirtualMachineResponseOutput {
	return o
}

func (o JitNetworkAccessRequestVirtualMachineResponseOutput) ToJitNetworkAccessRequestVirtualMachineResponseOutputWithContext(ctx context.Context) JitNetworkAccessRequestVirtualMachineResponseOutput {
	return o
}

func (o JitNetworkAccessRequestVirtualMachineResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestVirtualMachineResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o JitNetworkAccessRequestVirtualMachineResponseOutput) Ports() JitNetworkAccessRequestPortResponseArrayOutput {
	return o.ApplyT(func(v JitNetworkAccessRequestVirtualMachineResponse) []JitNetworkAccessRequestPortResponse {
		return v.Ports
	}).(JitNetworkAccessRequestPortResponseArrayOutput)
}

type JitNetworkAccessRequestVirtualMachineResponseArrayOutput struct{ *pulumi.OutputState }

func (JitNetworkAccessRequestVirtualMachineResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JitNetworkAccessRequestVirtualMachineResponse)(nil)).Elem()
}

func (o JitNetworkAccessRequestVirtualMachineResponseArrayOutput) ToJitNetworkAccessRequestVirtualMachineResponseArrayOutput() JitNetworkAccessRequestVirtualMachineResponseArrayOutput {
	return o
}

func (o JitNetworkAccessRequestVirtualMachineResponseArrayOutput) ToJitNetworkAccessRequestVirtualMachineResponseArrayOutputWithContext(ctx context.Context) JitNetworkAccessRequestVirtualMachineResponseArrayOutput {
	return o
}

func (o JitNetworkAccessRequestVirtualMachineResponseArrayOutput) Index(i pulumi.IntInput) JitNetworkAccessRequestVirtualMachineResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JitNetworkAccessRequestVirtualMachineResponse {
		return vs[0].([]JitNetworkAccessRequestVirtualMachineResponse)[vs[1].(int)]
	}).(JitNetworkAccessRequestVirtualMachineResponseOutput)
}

type OnPremiseResourceDetails struct {
	MachineName      string `pulumi:"machineName"`
	Source           string `pulumi:"source"`
	SourceComputerId string `pulumi:"sourceComputerId"`
	Vmuuid           string `pulumi:"vmuuid"`
	WorkspaceId      string `pulumi:"workspaceId"`
}

type OnPremiseResourceDetailsResponse struct {
	MachineName      string `pulumi:"machineName"`
	Source           string `pulumi:"source"`
	SourceComputerId string `pulumi:"sourceComputerId"`
	Vmuuid           string `pulumi:"vmuuid"`
	WorkspaceId      string `pulumi:"workspaceId"`
}

type OnPremiseSqlResourceDetails struct {
	DatabaseName     string `pulumi:"databaseName"`
	MachineName      string `pulumi:"machineName"`
	ServerName       string `pulumi:"serverName"`
	Source           string `pulumi:"source"`
	SourceComputerId string `pulumi:"sourceComputerId"`
	Vmuuid           string `pulumi:"vmuuid"`
	WorkspaceId      string `pulumi:"workspaceId"`
}

type OnPremiseSqlResourceDetailsResponse struct {
	DatabaseName     string `pulumi:"databaseName"`
	MachineName      string `pulumi:"machineName"`
	ServerName       string `pulumi:"serverName"`
	Source           string `pulumi:"source"`
	SourceComputerId string `pulumi:"sourceComputerId"`
	Vmuuid           string `pulumi:"vmuuid"`
	WorkspaceId      string `pulumi:"workspaceId"`
}

type PathRecommendation struct {
	Action              *string              `pulumi:"action"`
	Common              *bool                `pulumi:"common"`
	ConfigurationStatus *string              `pulumi:"configurationStatus"`
	FileType            *string              `pulumi:"fileType"`
	Path                *string              `pulumi:"path"`
	PublisherInfo       *PublisherInfo       `pulumi:"publisherInfo"`
	Type                *string              `pulumi:"type"`
	UserSids            []string             `pulumi:"userSids"`
	Usernames           []UserRecommendation `pulumi:"usernames"`
}





type PathRecommendationInput interface {
	pulumi.Input

	ToPathRecommendationOutput() PathRecommendationOutput
	ToPathRecommendationOutputWithContext(context.Context) PathRecommendationOutput
}

type PathRecommendationArgs struct {
	Action              pulumi.StringPtrInput        `pulumi:"action"`
	Common              pulumi.BoolPtrInput          `pulumi:"common"`
	ConfigurationStatus pulumi.StringPtrInput        `pulumi:"configurationStatus"`
	FileType            pulumi.StringPtrInput        `pulumi:"fileType"`
	Path                pulumi.StringPtrInput        `pulumi:"path"`
	PublisherInfo       PublisherInfoPtrInput        `pulumi:"publisherInfo"`
	Type                pulumi.StringPtrInput        `pulumi:"type"`
	UserSids            pulumi.StringArrayInput      `pulumi:"userSids"`
	Usernames           UserRecommendationArrayInput `pulumi:"usernames"`
}

func (PathRecommendationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PathRecommendation)(nil)).Elem()
}

func (i PathRecommendationArgs) ToPathRecommendationOutput() PathRecommendationOutput {
	return i.ToPathRecommendationOutputWithContext(context.Background())
}

func (i PathRecommendationArgs) ToPathRecommendationOutputWithContext(ctx context.Context) PathRecommendationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PathRecommendationOutput)
}





type PathRecommendationArrayInput interface {
	pulumi.Input

	ToPathRecommendationArrayOutput() PathRecommendationArrayOutput
	ToPathRecommendationArrayOutputWithContext(context.Context) PathRecommendationArrayOutput
}

type PathRecommendationArray []PathRecommendationInput

func (PathRecommendationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PathRecommendation)(nil)).Elem()
}

func (i PathRecommendationArray) ToPathRecommendationArrayOutput() PathRecommendationArrayOutput {
	return i.ToPathRecommendationArrayOutputWithContext(context.Background())
}

func (i PathRecommendationArray) ToPathRecommendationArrayOutputWithContext(ctx context.Context) PathRecommendationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PathRecommendationArrayOutput)
}

type PathRecommendationOutput struct{ *pulumi.OutputState }

func (PathRecommendationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PathRecommendation)(nil)).Elem()
}

func (o PathRecommendationOutput) ToPathRecommendationOutput() PathRecommendationOutput {
	return o
}

func (o PathRecommendationOutput) ToPathRecommendationOutputWithContext(ctx context.Context) PathRecommendationOutput {
	return o
}

func (o PathRecommendationOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendation) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationOutput) Common() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PathRecommendation) *bool { return v.Common }).(pulumi.BoolPtrOutput)
}

func (o PathRecommendationOutput) ConfigurationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendation) *string { return v.ConfigurationStatus }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationOutput) FileType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendation) *string { return v.FileType }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendation) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationOutput) PublisherInfo() PublisherInfoPtrOutput {
	return o.ApplyT(func(v PathRecommendation) *PublisherInfo { return v.PublisherInfo }).(PublisherInfoPtrOutput)
}

func (o PathRecommendationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationOutput) UserSids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PathRecommendation) []string { return v.UserSids }).(pulumi.StringArrayOutput)
}

func (o PathRecommendationOutput) Usernames() UserRecommendationArrayOutput {
	return o.ApplyT(func(v PathRecommendation) []UserRecommendation { return v.Usernames }).(UserRecommendationArrayOutput)
}

type PathRecommendationArrayOutput struct{ *pulumi.OutputState }

func (PathRecommendationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PathRecommendation)(nil)).Elem()
}

func (o PathRecommendationArrayOutput) ToPathRecommendationArrayOutput() PathRecommendationArrayOutput {
	return o
}

func (o PathRecommendationArrayOutput) ToPathRecommendationArrayOutputWithContext(ctx context.Context) PathRecommendationArrayOutput {
	return o
}

func (o PathRecommendationArrayOutput) Index(i pulumi.IntInput) PathRecommendationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PathRecommendation {
		return vs[0].([]PathRecommendation)[vs[1].(int)]
	}).(PathRecommendationOutput)
}

type PathRecommendationResponse struct {
	Action              *string                      `pulumi:"action"`
	Common              *bool                        `pulumi:"common"`
	ConfigurationStatus *string                      `pulumi:"configurationStatus"`
	FileType            *string                      `pulumi:"fileType"`
	Path                *string                      `pulumi:"path"`
	PublisherInfo       *PublisherInfoResponse       `pulumi:"publisherInfo"`
	Type                *string                      `pulumi:"type"`
	UserSids            []string                     `pulumi:"userSids"`
	Usernames           []UserRecommendationResponse `pulumi:"usernames"`
}

type PathRecommendationResponseOutput struct{ *pulumi.OutputState }

func (PathRecommendationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PathRecommendationResponse)(nil)).Elem()
}

func (o PathRecommendationResponseOutput) ToPathRecommendationResponseOutput() PathRecommendationResponseOutput {
	return o
}

func (o PathRecommendationResponseOutput) ToPathRecommendationResponseOutputWithContext(ctx context.Context) PathRecommendationResponseOutput {
	return o
}

func (o PathRecommendationResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendationResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationResponseOutput) Common() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PathRecommendationResponse) *bool { return v.Common }).(pulumi.BoolPtrOutput)
}

func (o PathRecommendationResponseOutput) ConfigurationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendationResponse) *string { return v.ConfigurationStatus }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationResponseOutput) FileType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendationResponse) *string { return v.FileType }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendationResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationResponseOutput) PublisherInfo() PublisherInfoResponsePtrOutput {
	return o.ApplyT(func(v PathRecommendationResponse) *PublisherInfoResponse { return v.PublisherInfo }).(PublisherInfoResponsePtrOutput)
}

func (o PathRecommendationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PathRecommendationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o PathRecommendationResponseOutput) UserSids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PathRecommendationResponse) []string { return v.UserSids }).(pulumi.StringArrayOutput)
}

func (o PathRecommendationResponseOutput) Usernames() UserRecommendationResponseArrayOutput {
	return o.ApplyT(func(v PathRecommendationResponse) []UserRecommendationResponse { return v.Usernames }).(UserRecommendationResponseArrayOutput)
}

type PathRecommendationResponseArrayOutput struct{ *pulumi.OutputState }

func (PathRecommendationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PathRecommendationResponse)(nil)).Elem()
}

func (o PathRecommendationResponseArrayOutput) ToPathRecommendationResponseArrayOutput() PathRecommendationResponseArrayOutput {
	return o
}

func (o PathRecommendationResponseArrayOutput) ToPathRecommendationResponseArrayOutputWithContext(ctx context.Context) PathRecommendationResponseArrayOutput {
	return o
}

func (o PathRecommendationResponseArrayOutput) Index(i pulumi.IntInput) PathRecommendationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PathRecommendationResponse {
		return vs[0].([]PathRecommendationResponse)[vs[1].(int)]
	}).(PathRecommendationResponseOutput)
}

type ProtectionMode struct {
	Exe        *string `pulumi:"exe"`
	Executable *string `pulumi:"executable"`
	Msi        *string `pulumi:"msi"`
	Script     *string `pulumi:"script"`
}





type ProtectionModeInput interface {
	pulumi.Input

	ToProtectionModeOutput() ProtectionModeOutput
	ToProtectionModeOutputWithContext(context.Context) ProtectionModeOutput
}

type ProtectionModeArgs struct {
	Exe        pulumi.StringPtrInput `pulumi:"exe"`
	Executable pulumi.StringPtrInput `pulumi:"executable"`
	Msi        pulumi.StringPtrInput `pulumi:"msi"`
	Script     pulumi.StringPtrInput `pulumi:"script"`
}

func (ProtectionModeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionMode)(nil)).Elem()
}

func (i ProtectionModeArgs) ToProtectionModeOutput() ProtectionModeOutput {
	return i.ToProtectionModeOutputWithContext(context.Background())
}

func (i ProtectionModeArgs) ToProtectionModeOutputWithContext(ctx context.Context) ProtectionModeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionModeOutput)
}

func (i ProtectionModeArgs) ToProtectionModePtrOutput() ProtectionModePtrOutput {
	return i.ToProtectionModePtrOutputWithContext(context.Background())
}

func (i ProtectionModeArgs) ToProtectionModePtrOutputWithContext(ctx context.Context) ProtectionModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionModeOutput).ToProtectionModePtrOutputWithContext(ctx)
}









type ProtectionModePtrInput interface {
	pulumi.Input

	ToProtectionModePtrOutput() ProtectionModePtrOutput
	ToProtectionModePtrOutputWithContext(context.Context) ProtectionModePtrOutput
}

type protectionModePtrType ProtectionModeArgs

func ProtectionModePtr(v *ProtectionModeArgs) ProtectionModePtrInput {
	return (*protectionModePtrType)(v)
}

func (*protectionModePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionMode)(nil)).Elem()
}

func (i *protectionModePtrType) ToProtectionModePtrOutput() ProtectionModePtrOutput {
	return i.ToProtectionModePtrOutputWithContext(context.Background())
}

func (i *protectionModePtrType) ToProtectionModePtrOutputWithContext(ctx context.Context) ProtectionModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionModePtrOutput)
}

type ProtectionModeOutput struct{ *pulumi.OutputState }

func (ProtectionModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionMode)(nil)).Elem()
}

func (o ProtectionModeOutput) ToProtectionModeOutput() ProtectionModeOutput {
	return o
}

func (o ProtectionModeOutput) ToProtectionModeOutputWithContext(ctx context.Context) ProtectionModeOutput {
	return o
}

func (o ProtectionModeOutput) ToProtectionModePtrOutput() ProtectionModePtrOutput {
	return o.ToProtectionModePtrOutputWithContext(context.Background())
}

func (o ProtectionModeOutput) ToProtectionModePtrOutputWithContext(ctx context.Context) ProtectionModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProtectionMode) *ProtectionMode {
		return &v
	}).(ProtectionModePtrOutput)
}

func (o ProtectionModeOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionMode) *string { return v.Exe }).(pulumi.StringPtrOutput)
}

func (o ProtectionModeOutput) Executable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionMode) *string { return v.Executable }).(pulumi.StringPtrOutput)
}

func (o ProtectionModeOutput) Msi() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionMode) *string { return v.Msi }).(pulumi.StringPtrOutput)
}

func (o ProtectionModeOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionMode) *string { return v.Script }).(pulumi.StringPtrOutput)
}

type ProtectionModePtrOutput struct{ *pulumi.OutputState }

func (ProtectionModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionMode)(nil)).Elem()
}

func (o ProtectionModePtrOutput) ToProtectionModePtrOutput() ProtectionModePtrOutput {
	return o
}

func (o ProtectionModePtrOutput) ToProtectionModePtrOutputWithContext(ctx context.Context) ProtectionModePtrOutput {
	return o
}

func (o ProtectionModePtrOutput) Elem() ProtectionModeOutput {
	return o.ApplyT(func(v *ProtectionMode) ProtectionMode {
		if v != nil {
			return *v
		}
		var ret ProtectionMode
		return ret
	}).(ProtectionModeOutput)
}

func (o ProtectionModePtrOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionMode) *string {
		if v == nil {
			return nil
		}
		return v.Exe
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionModePtrOutput) Executable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionMode) *string {
		if v == nil {
			return nil
		}
		return v.Executable
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionModePtrOutput) Msi() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionMode) *string {
		if v == nil {
			return nil
		}
		return v.Msi
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionModePtrOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionMode) *string {
		if v == nil {
			return nil
		}
		return v.Script
	}).(pulumi.StringPtrOutput)
}

type ProtectionModeResponse struct {
	Exe        *string `pulumi:"exe"`
	Executable *string `pulumi:"executable"`
	Msi        *string `pulumi:"msi"`
	Script     *string `pulumi:"script"`
}

type ProtectionModeResponseOutput struct{ *pulumi.OutputState }

func (ProtectionModeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionModeResponse)(nil)).Elem()
}

func (o ProtectionModeResponseOutput) ToProtectionModeResponseOutput() ProtectionModeResponseOutput {
	return o
}

func (o ProtectionModeResponseOutput) ToProtectionModeResponseOutputWithContext(ctx context.Context) ProtectionModeResponseOutput {
	return o
}

func (o ProtectionModeResponseOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionModeResponse) *string { return v.Exe }).(pulumi.StringPtrOutput)
}

func (o ProtectionModeResponseOutput) Executable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionModeResponse) *string { return v.Executable }).(pulumi.StringPtrOutput)
}

func (o ProtectionModeResponseOutput) Msi() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionModeResponse) *string { return v.Msi }).(pulumi.StringPtrOutput)
}

func (o ProtectionModeResponseOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionModeResponse) *string { return v.Script }).(pulumi.StringPtrOutput)
}

type ProtectionModeResponsePtrOutput struct{ *pulumi.OutputState }

func (ProtectionModeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionModeResponse)(nil)).Elem()
}

func (o ProtectionModeResponsePtrOutput) ToProtectionModeResponsePtrOutput() ProtectionModeResponsePtrOutput {
	return o
}

func (o ProtectionModeResponsePtrOutput) ToProtectionModeResponsePtrOutputWithContext(ctx context.Context) ProtectionModeResponsePtrOutput {
	return o
}

func (o ProtectionModeResponsePtrOutput) Elem() ProtectionModeResponseOutput {
	return o.ApplyT(func(v *ProtectionModeResponse) ProtectionModeResponse {
		if v != nil {
			return *v
		}
		var ret ProtectionModeResponse
		return ret
	}).(ProtectionModeResponseOutput)
}

func (o ProtectionModeResponsePtrOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionModeResponse) *string {
		if v == nil {
			return nil
		}
		return v.Exe
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionModeResponsePtrOutput) Executable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionModeResponse) *string {
		if v == nil {
			return nil
		}
		return v.Executable
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionModeResponsePtrOutput) Msi() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionModeResponse) *string {
		if v == nil {
			return nil
		}
		return v.Msi
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionModeResponsePtrOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionModeResponse) *string {
		if v == nil {
			return nil
		}
		return v.Script
	}).(pulumi.StringPtrOutput)
}

type PublisherInfo struct {
	BinaryName    *string `pulumi:"binaryName"`
	ProductName   *string `pulumi:"productName"`
	PublisherName *string `pulumi:"publisherName"`
	Version       *string `pulumi:"version"`
}





type PublisherInfoInput interface {
	pulumi.Input

	ToPublisherInfoOutput() PublisherInfoOutput
	ToPublisherInfoOutputWithContext(context.Context) PublisherInfoOutput
}

type PublisherInfoArgs struct {
	BinaryName    pulumi.StringPtrInput `pulumi:"binaryName"`
	ProductName   pulumi.StringPtrInput `pulumi:"productName"`
	PublisherName pulumi.StringPtrInput `pulumi:"publisherName"`
	Version       pulumi.StringPtrInput `pulumi:"version"`
}

func (PublisherInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublisherInfo)(nil)).Elem()
}

func (i PublisherInfoArgs) ToPublisherInfoOutput() PublisherInfoOutput {
	return i.ToPublisherInfoOutputWithContext(context.Background())
}

func (i PublisherInfoArgs) ToPublisherInfoOutputWithContext(ctx context.Context) PublisherInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublisherInfoOutput)
}

func (i PublisherInfoArgs) ToPublisherInfoPtrOutput() PublisherInfoPtrOutput {
	return i.ToPublisherInfoPtrOutputWithContext(context.Background())
}

func (i PublisherInfoArgs) ToPublisherInfoPtrOutputWithContext(ctx context.Context) PublisherInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublisherInfoOutput).ToPublisherInfoPtrOutputWithContext(ctx)
}









type PublisherInfoPtrInput interface {
	pulumi.Input

	ToPublisherInfoPtrOutput() PublisherInfoPtrOutput
	ToPublisherInfoPtrOutputWithContext(context.Context) PublisherInfoPtrOutput
}

type publisherInfoPtrType PublisherInfoArgs

func PublisherInfoPtr(v *PublisherInfoArgs) PublisherInfoPtrInput {
	return (*publisherInfoPtrType)(v)
}

func (*publisherInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublisherInfo)(nil)).Elem()
}

func (i *publisherInfoPtrType) ToPublisherInfoPtrOutput() PublisherInfoPtrOutput {
	return i.ToPublisherInfoPtrOutputWithContext(context.Background())
}

func (i *publisherInfoPtrType) ToPublisherInfoPtrOutputWithContext(ctx context.Context) PublisherInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublisherInfoPtrOutput)
}

type PublisherInfoOutput struct{ *pulumi.OutputState }

func (PublisherInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublisherInfo)(nil)).Elem()
}

func (o PublisherInfoOutput) ToPublisherInfoOutput() PublisherInfoOutput {
	return o
}

func (o PublisherInfoOutput) ToPublisherInfoOutputWithContext(ctx context.Context) PublisherInfoOutput {
	return o
}

func (o PublisherInfoOutput) ToPublisherInfoPtrOutput() PublisherInfoPtrOutput {
	return o.ToPublisherInfoPtrOutputWithContext(context.Background())
}

func (o PublisherInfoOutput) ToPublisherInfoPtrOutputWithContext(ctx context.Context) PublisherInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublisherInfo) *PublisherInfo {
		return &v
	}).(PublisherInfoPtrOutput)
}

func (o PublisherInfoOutput) BinaryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublisherInfo) *string { return v.BinaryName }).(pulumi.StringPtrOutput)
}

func (o PublisherInfoOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublisherInfo) *string { return v.ProductName }).(pulumi.StringPtrOutput)
}

func (o PublisherInfoOutput) PublisherName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublisherInfo) *string { return v.PublisherName }).(pulumi.StringPtrOutput)
}

func (o PublisherInfoOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublisherInfo) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type PublisherInfoPtrOutput struct{ *pulumi.OutputState }

func (PublisherInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublisherInfo)(nil)).Elem()
}

func (o PublisherInfoPtrOutput) ToPublisherInfoPtrOutput() PublisherInfoPtrOutput {
	return o
}

func (o PublisherInfoPtrOutput) ToPublisherInfoPtrOutputWithContext(ctx context.Context) PublisherInfoPtrOutput {
	return o
}

func (o PublisherInfoPtrOutput) Elem() PublisherInfoOutput {
	return o.ApplyT(func(v *PublisherInfo) PublisherInfo {
		if v != nil {
			return *v
		}
		var ret PublisherInfo
		return ret
	}).(PublisherInfoOutput)
}

func (o PublisherInfoPtrOutput) BinaryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublisherInfo) *string {
		if v == nil {
			return nil
		}
		return v.BinaryName
	}).(pulumi.StringPtrOutput)
}

func (o PublisherInfoPtrOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublisherInfo) *string {
		if v == nil {
			return nil
		}
		return v.ProductName
	}).(pulumi.StringPtrOutput)
}

func (o PublisherInfoPtrOutput) PublisherName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublisherInfo) *string {
		if v == nil {
			return nil
		}
		return v.PublisherName
	}).(pulumi.StringPtrOutput)
}

func (o PublisherInfoPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublisherInfo) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type PublisherInfoResponse struct {
	BinaryName    *string `pulumi:"binaryName"`
	ProductName   *string `pulumi:"productName"`
	PublisherName *string `pulumi:"publisherName"`
	Version       *string `pulumi:"version"`
}

type PublisherInfoResponseOutput struct{ *pulumi.OutputState }

func (PublisherInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublisherInfoResponse)(nil)).Elem()
}

func (o PublisherInfoResponseOutput) ToPublisherInfoResponseOutput() PublisherInfoResponseOutput {
	return o
}

func (o PublisherInfoResponseOutput) ToPublisherInfoResponseOutputWithContext(ctx context.Context) PublisherInfoResponseOutput {
	return o
}

func (o PublisherInfoResponseOutput) BinaryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublisherInfoResponse) *string { return v.BinaryName }).(pulumi.StringPtrOutput)
}

func (o PublisherInfoResponseOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublisherInfoResponse) *string { return v.ProductName }).(pulumi.StringPtrOutput)
}

func (o PublisherInfoResponseOutput) PublisherName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublisherInfoResponse) *string { return v.PublisherName }).(pulumi.StringPtrOutput)
}

func (o PublisherInfoResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublisherInfoResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type PublisherInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (PublisherInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublisherInfoResponse)(nil)).Elem()
}

func (o PublisherInfoResponsePtrOutput) ToPublisherInfoResponsePtrOutput() PublisherInfoResponsePtrOutput {
	return o
}

func (o PublisherInfoResponsePtrOutput) ToPublisherInfoResponsePtrOutputWithContext(ctx context.Context) PublisherInfoResponsePtrOutput {
	return o
}

func (o PublisherInfoResponsePtrOutput) Elem() PublisherInfoResponseOutput {
	return o.ApplyT(func(v *PublisherInfoResponse) PublisherInfoResponse {
		if v != nil {
			return *v
		}
		var ret PublisherInfoResponse
		return ret
	}).(PublisherInfoResponseOutput)
}

func (o PublisherInfoResponsePtrOutput) BinaryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublisherInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.BinaryName
	}).(pulumi.StringPtrOutput)
}

func (o PublisherInfoResponsePtrOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublisherInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProductName
	}).(pulumi.StringPtrOutput)
}

func (o PublisherInfoResponsePtrOutput) PublisherName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublisherInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublisherName
	}).(pulumi.StringPtrOutput)
}

func (o PublisherInfoResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublisherInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type SecurityAssessmentMetadataPartnerData struct {
	PartnerName string  `pulumi:"partnerName"`
	ProductName *string `pulumi:"productName"`
	Secret      string  `pulumi:"secret"`
}





type SecurityAssessmentMetadataPartnerDataInput interface {
	pulumi.Input

	ToSecurityAssessmentMetadataPartnerDataOutput() SecurityAssessmentMetadataPartnerDataOutput
	ToSecurityAssessmentMetadataPartnerDataOutputWithContext(context.Context) SecurityAssessmentMetadataPartnerDataOutput
}

type SecurityAssessmentMetadataPartnerDataArgs struct {
	PartnerName pulumi.StringInput    `pulumi:"partnerName"`
	ProductName pulumi.StringPtrInput `pulumi:"productName"`
	Secret      pulumi.StringInput    `pulumi:"secret"`
}

func (SecurityAssessmentMetadataPartnerDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataPartnerData)(nil)).Elem()
}

func (i SecurityAssessmentMetadataPartnerDataArgs) ToSecurityAssessmentMetadataPartnerDataOutput() SecurityAssessmentMetadataPartnerDataOutput {
	return i.ToSecurityAssessmentMetadataPartnerDataOutputWithContext(context.Background())
}

func (i SecurityAssessmentMetadataPartnerDataArgs) ToSecurityAssessmentMetadataPartnerDataOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPartnerDataOutput)
}

func (i SecurityAssessmentMetadataPartnerDataArgs) ToSecurityAssessmentMetadataPartnerDataPtrOutput() SecurityAssessmentMetadataPartnerDataPtrOutput {
	return i.ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(context.Background())
}

func (i SecurityAssessmentMetadataPartnerDataArgs) ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPartnerDataOutput).ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(ctx)
}









type SecurityAssessmentMetadataPartnerDataPtrInput interface {
	pulumi.Input

	ToSecurityAssessmentMetadataPartnerDataPtrOutput() SecurityAssessmentMetadataPartnerDataPtrOutput
	ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(context.Context) SecurityAssessmentMetadataPartnerDataPtrOutput
}

type securityAssessmentMetadataPartnerDataPtrType SecurityAssessmentMetadataPartnerDataArgs

func SecurityAssessmentMetadataPartnerDataPtr(v *SecurityAssessmentMetadataPartnerDataArgs) SecurityAssessmentMetadataPartnerDataPtrInput {
	return (*securityAssessmentMetadataPartnerDataPtrType)(v)
}

func (*securityAssessmentMetadataPartnerDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataPartnerData)(nil)).Elem()
}

func (i *securityAssessmentMetadataPartnerDataPtrType) ToSecurityAssessmentMetadataPartnerDataPtrOutput() SecurityAssessmentMetadataPartnerDataPtrOutput {
	return i.ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(context.Background())
}

func (i *securityAssessmentMetadataPartnerDataPtrType) ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPartnerDataPtrOutput)
}

type SecurityAssessmentMetadataPartnerDataOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPartnerDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataPartnerData)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPartnerDataOutput) ToSecurityAssessmentMetadataPartnerDataOutput() SecurityAssessmentMetadataPartnerDataOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataOutput) ToSecurityAssessmentMetadataPartnerDataOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataOutput) ToSecurityAssessmentMetadataPartnerDataPtrOutput() SecurityAssessmentMetadataPartnerDataPtrOutput {
	return o.ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(context.Background())
}

func (o SecurityAssessmentMetadataPartnerDataOutput) ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityAssessmentMetadataPartnerData) *SecurityAssessmentMetadataPartnerData {
		return &v
	}).(SecurityAssessmentMetadataPartnerDataPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataOutput) PartnerName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPartnerData) string { return v.PartnerName }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPartnerDataOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPartnerData) *string { return v.ProductName }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPartnerData) string { return v.Secret }).(pulumi.StringOutput)
}

type SecurityAssessmentMetadataPartnerDataPtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPartnerDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataPartnerData)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPartnerDataPtrOutput) ToSecurityAssessmentMetadataPartnerDataPtrOutput() SecurityAssessmentMetadataPartnerDataPtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataPtrOutput) ToSecurityAssessmentMetadataPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataPtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataPtrOutput) Elem() SecurityAssessmentMetadataPartnerDataOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerData) SecurityAssessmentMetadataPartnerData {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentMetadataPartnerData
		return ret
	}).(SecurityAssessmentMetadataPartnerDataOutput)
}

func (o SecurityAssessmentMetadataPartnerDataPtrOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerData) *string {
		if v == nil {
			return nil
		}
		return &v.PartnerName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataPtrOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerData) *string {
		if v == nil {
			return nil
		}
		return v.ProductName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataPtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerData) *string {
		if v == nil {
			return nil
		}
		return &v.Secret
	}).(pulumi.StringPtrOutput)
}

type SecurityAssessmentMetadataPartnerDataResponse struct {
	PartnerName string  `pulumi:"partnerName"`
	ProductName *string `pulumi:"productName"`
	Secret      string  `pulumi:"secret"`
}

type SecurityAssessmentMetadataPartnerDataResponseOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPartnerDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataPartnerDataResponse)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPartnerDataResponseOutput) ToSecurityAssessmentMetadataPartnerDataResponseOutput() SecurityAssessmentMetadataPartnerDataResponseOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataResponseOutput) ToSecurityAssessmentMetadataPartnerDataResponseOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataResponseOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataResponseOutput) PartnerName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPartnerDataResponse) string { return v.PartnerName }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPartnerDataResponseOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPartnerDataResponse) *string { return v.ProductName }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataResponseOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPartnerDataResponse) string { return v.Secret }).(pulumi.StringOutput)
}

type SecurityAssessmentMetadataPartnerDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPartnerDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataPartnerDataResponse)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPartnerDataResponsePtrOutput) ToSecurityAssessmentMetadataPartnerDataResponsePtrOutput() SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataResponsePtrOutput) ToSecurityAssessmentMetadataPartnerDataResponsePtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPartnerDataResponsePtrOutput) Elem() SecurityAssessmentMetadataPartnerDataResponseOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerDataResponse) SecurityAssessmentMetadataPartnerDataResponse {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentMetadataPartnerDataResponse
		return ret
	}).(SecurityAssessmentMetadataPartnerDataResponseOutput)
}

func (o SecurityAssessmentMetadataPartnerDataResponsePtrOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PartnerName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataResponsePtrOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProductName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPartnerDataResponsePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPartnerDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Secret
	}).(pulumi.StringPtrOutput)
}

type SecurityAssessmentMetadataProperties struct {
	AssessmentType         string                                 `pulumi:"assessmentType"`
	Categories             []string                               `pulumi:"categories"`
	Description            *string                                `pulumi:"description"`
	DisplayName            string                                 `pulumi:"displayName"`
	ImplementationEffort   *string                                `pulumi:"implementationEffort"`
	PartnerData            *SecurityAssessmentMetadataPartnerData `pulumi:"partnerData"`
	Preview                *bool                                  `pulumi:"preview"`
	RemediationDescription *string                                `pulumi:"remediationDescription"`
	Severity               string                                 `pulumi:"severity"`
	Threats                []string                               `pulumi:"threats"`
	UserImpact             *string                                `pulumi:"userImpact"`
}





type SecurityAssessmentMetadataPropertiesInput interface {
	pulumi.Input

	ToSecurityAssessmentMetadataPropertiesOutput() SecurityAssessmentMetadataPropertiesOutput
	ToSecurityAssessmentMetadataPropertiesOutputWithContext(context.Context) SecurityAssessmentMetadataPropertiesOutput
}

type SecurityAssessmentMetadataPropertiesArgs struct {
	AssessmentType         pulumi.StringInput                            `pulumi:"assessmentType"`
	Categories             pulumi.StringArrayInput                       `pulumi:"categories"`
	Description            pulumi.StringPtrInput                         `pulumi:"description"`
	DisplayName            pulumi.StringInput                            `pulumi:"displayName"`
	ImplementationEffort   pulumi.StringPtrInput                         `pulumi:"implementationEffort"`
	PartnerData            SecurityAssessmentMetadataPartnerDataPtrInput `pulumi:"partnerData"`
	Preview                pulumi.BoolPtrInput                           `pulumi:"preview"`
	RemediationDescription pulumi.StringPtrInput                         `pulumi:"remediationDescription"`
	Severity               pulumi.StringInput                            `pulumi:"severity"`
	Threats                pulumi.StringArrayInput                       `pulumi:"threats"`
	UserImpact             pulumi.StringPtrInput                         `pulumi:"userImpact"`
}

func (SecurityAssessmentMetadataPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataProperties)(nil)).Elem()
}

func (i SecurityAssessmentMetadataPropertiesArgs) ToSecurityAssessmentMetadataPropertiesOutput() SecurityAssessmentMetadataPropertiesOutput {
	return i.ToSecurityAssessmentMetadataPropertiesOutputWithContext(context.Background())
}

func (i SecurityAssessmentMetadataPropertiesArgs) ToSecurityAssessmentMetadataPropertiesOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPropertiesOutput)
}

func (i SecurityAssessmentMetadataPropertiesArgs) ToSecurityAssessmentMetadataPropertiesPtrOutput() SecurityAssessmentMetadataPropertiesPtrOutput {
	return i.ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(context.Background())
}

func (i SecurityAssessmentMetadataPropertiesArgs) ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPropertiesOutput).ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(ctx)
}









type SecurityAssessmentMetadataPropertiesPtrInput interface {
	pulumi.Input

	ToSecurityAssessmentMetadataPropertiesPtrOutput() SecurityAssessmentMetadataPropertiesPtrOutput
	ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(context.Context) SecurityAssessmentMetadataPropertiesPtrOutput
}

type securityAssessmentMetadataPropertiesPtrType SecurityAssessmentMetadataPropertiesArgs

func SecurityAssessmentMetadataPropertiesPtr(v *SecurityAssessmentMetadataPropertiesArgs) SecurityAssessmentMetadataPropertiesPtrInput {
	return (*securityAssessmentMetadataPropertiesPtrType)(v)
}

func (*securityAssessmentMetadataPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataProperties)(nil)).Elem()
}

func (i *securityAssessmentMetadataPropertiesPtrType) ToSecurityAssessmentMetadataPropertiesPtrOutput() SecurityAssessmentMetadataPropertiesPtrOutput {
	return i.ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(context.Background())
}

func (i *securityAssessmentMetadataPropertiesPtrType) ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentMetadataPropertiesPtrOutput)
}

type SecurityAssessmentMetadataPropertiesOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataProperties)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPropertiesOutput) ToSecurityAssessmentMetadataPropertiesOutput() SecurityAssessmentMetadataPropertiesOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesOutput) ToSecurityAssessmentMetadataPropertiesOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesOutput) ToSecurityAssessmentMetadataPropertiesPtrOutput() SecurityAssessmentMetadataPropertiesPtrOutput {
	return o.ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(context.Background())
}

func (o SecurityAssessmentMetadataPropertiesOutput) ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityAssessmentMetadataProperties) *SecurityAssessmentMetadataProperties {
		return &v
	}).(SecurityAssessmentMetadataPropertiesPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) AssessmentType() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) string { return v.AssessmentType }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) ImplementationEffort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) *string { return v.ImplementationEffort }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) PartnerData() SecurityAssessmentMetadataPartnerDataPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) *SecurityAssessmentMetadataPartnerData {
		return v.PartnerData
	}).(SecurityAssessmentMetadataPartnerDataPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) *bool { return v.Preview }).(pulumi.BoolPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) *string { return v.RemediationDescription }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) string { return v.Severity }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) Threats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) []string { return v.Threats }).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesOutput) UserImpact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataProperties) *string { return v.UserImpact }).(pulumi.StringPtrOutput)
}

type SecurityAssessmentMetadataPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataProperties)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) ToSecurityAssessmentMetadataPropertiesPtrOutput() SecurityAssessmentMetadataPropertiesPtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) ToSecurityAssessmentMetadataPropertiesPtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesPtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) Elem() SecurityAssessmentMetadataPropertiesOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) SecurityAssessmentMetadataProperties {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentMetadataProperties
		return ret
	}).(SecurityAssessmentMetadataPropertiesOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) AssessmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AssessmentType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) []string {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) ImplementationEffort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return v.ImplementationEffort
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) PartnerData() SecurityAssessmentMetadataPartnerDataPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *SecurityAssessmentMetadataPartnerData {
		if v == nil {
			return nil
		}
		return v.PartnerData
	}).(SecurityAssessmentMetadataPartnerDataPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Preview
	}).(pulumi.BoolPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return v.RemediationDescription
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) Severity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Severity
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) Threats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) []string {
		if v == nil {
			return nil
		}
		return v.Threats
	}).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesPtrOutput) UserImpact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataProperties) *string {
		if v == nil {
			return nil
		}
		return v.UserImpact
	}).(pulumi.StringPtrOutput)
}

type SecurityAssessmentMetadataPropertiesResponse struct {
	AssessmentType         string                                         `pulumi:"assessmentType"`
	Categories             []string                                       `pulumi:"categories"`
	Description            *string                                        `pulumi:"description"`
	DisplayName            string                                         `pulumi:"displayName"`
	ImplementationEffort   *string                                        `pulumi:"implementationEffort"`
	PartnerData            *SecurityAssessmentMetadataPartnerDataResponse `pulumi:"partnerData"`
	PolicyDefinitionId     string                                         `pulumi:"policyDefinitionId"`
	Preview                *bool                                          `pulumi:"preview"`
	RemediationDescription *string                                        `pulumi:"remediationDescription"`
	Severity               string                                         `pulumi:"severity"`
	Threats                []string                                       `pulumi:"threats"`
	UserImpact             *string                                        `pulumi:"userImpact"`
}

type SecurityAssessmentMetadataPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentMetadataPropertiesResponse)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) ToSecurityAssessmentMetadataPropertiesResponseOutput() SecurityAssessmentMetadataPropertiesResponseOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) ToSecurityAssessmentMetadataPropertiesResponseOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesResponseOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) AssessmentType() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) string { return v.AssessmentType }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) ImplementationEffort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) *string { return v.ImplementationEffort }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) PartnerData() SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) *SecurityAssessmentMetadataPartnerDataResponse {
		return v.PartnerData
	}).(SecurityAssessmentMetadataPartnerDataResponsePtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) PolicyDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) string { return v.PolicyDefinitionId }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) *bool { return v.Preview }).(pulumi.BoolPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) *string { return v.RemediationDescription }).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) string { return v.Severity }).(pulumi.StringOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) Threats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) []string { return v.Threats }).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponseOutput) UserImpact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAssessmentMetadataPropertiesResponse) *string { return v.UserImpact }).(pulumi.StringPtrOutput)
}

type SecurityAssessmentMetadataPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentMetadataPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentMetadataPropertiesResponse)(nil)).Elem()
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) ToSecurityAssessmentMetadataPropertiesResponsePtrOutput() SecurityAssessmentMetadataPropertiesResponsePtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) ToSecurityAssessmentMetadataPropertiesResponsePtrOutputWithContext(ctx context.Context) SecurityAssessmentMetadataPropertiesResponsePtrOutput {
	return o
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) Elem() SecurityAssessmentMetadataPropertiesResponseOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) SecurityAssessmentMetadataPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentMetadataPropertiesResponse
		return ret
	}).(SecurityAssessmentMetadataPropertiesResponseOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) AssessmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AssessmentType
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) ImplementationEffort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ImplementationEffort
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) PartnerData() SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *SecurityAssessmentMetadataPartnerDataResponse {
		if v == nil {
			return nil
		}
		return v.PartnerData
	}).(SecurityAssessmentMetadataPartnerDataResponsePtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) PolicyDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PolicyDefinitionId
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Preview
	}).(pulumi.BoolPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RemediationDescription
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) Severity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Severity
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) Threats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.Threats
	}).(pulumi.StringArrayOutput)
}

func (o SecurityAssessmentMetadataPropertiesResponsePtrOutput) UserImpact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentMetadataPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserImpact
	}).(pulumi.StringPtrOutput)
}

type SecurityAssessmentPartnerData struct {
	PartnerName string `pulumi:"partnerName"`
	Secret      string `pulumi:"secret"`
}





type SecurityAssessmentPartnerDataInput interface {
	pulumi.Input

	ToSecurityAssessmentPartnerDataOutput() SecurityAssessmentPartnerDataOutput
	ToSecurityAssessmentPartnerDataOutputWithContext(context.Context) SecurityAssessmentPartnerDataOutput
}

type SecurityAssessmentPartnerDataArgs struct {
	PartnerName pulumi.StringInput `pulumi:"partnerName"`
	Secret      pulumi.StringInput `pulumi:"secret"`
}

func (SecurityAssessmentPartnerDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentPartnerData)(nil)).Elem()
}

func (i SecurityAssessmentPartnerDataArgs) ToSecurityAssessmentPartnerDataOutput() SecurityAssessmentPartnerDataOutput {
	return i.ToSecurityAssessmentPartnerDataOutputWithContext(context.Background())
}

func (i SecurityAssessmentPartnerDataArgs) ToSecurityAssessmentPartnerDataOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentPartnerDataOutput)
}

func (i SecurityAssessmentPartnerDataArgs) ToSecurityAssessmentPartnerDataPtrOutput() SecurityAssessmentPartnerDataPtrOutput {
	return i.ToSecurityAssessmentPartnerDataPtrOutputWithContext(context.Background())
}

func (i SecurityAssessmentPartnerDataArgs) ToSecurityAssessmentPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentPartnerDataOutput).ToSecurityAssessmentPartnerDataPtrOutputWithContext(ctx)
}









type SecurityAssessmentPartnerDataPtrInput interface {
	pulumi.Input

	ToSecurityAssessmentPartnerDataPtrOutput() SecurityAssessmentPartnerDataPtrOutput
	ToSecurityAssessmentPartnerDataPtrOutputWithContext(context.Context) SecurityAssessmentPartnerDataPtrOutput
}

type securityAssessmentPartnerDataPtrType SecurityAssessmentPartnerDataArgs

func SecurityAssessmentPartnerDataPtr(v *SecurityAssessmentPartnerDataArgs) SecurityAssessmentPartnerDataPtrInput {
	return (*securityAssessmentPartnerDataPtrType)(v)
}

func (*securityAssessmentPartnerDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentPartnerData)(nil)).Elem()
}

func (i *securityAssessmentPartnerDataPtrType) ToSecurityAssessmentPartnerDataPtrOutput() SecurityAssessmentPartnerDataPtrOutput {
	return i.ToSecurityAssessmentPartnerDataPtrOutputWithContext(context.Background())
}

func (i *securityAssessmentPartnerDataPtrType) ToSecurityAssessmentPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAssessmentPartnerDataPtrOutput)
}

type SecurityAssessmentPartnerDataOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentPartnerDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentPartnerData)(nil)).Elem()
}

func (o SecurityAssessmentPartnerDataOutput) ToSecurityAssessmentPartnerDataOutput() SecurityAssessmentPartnerDataOutput {
	return o
}

func (o SecurityAssessmentPartnerDataOutput) ToSecurityAssessmentPartnerDataOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataOutput {
	return o
}

func (o SecurityAssessmentPartnerDataOutput) ToSecurityAssessmentPartnerDataPtrOutput() SecurityAssessmentPartnerDataPtrOutput {
	return o.ToSecurityAssessmentPartnerDataPtrOutputWithContext(context.Background())
}

func (o SecurityAssessmentPartnerDataOutput) ToSecurityAssessmentPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecurityAssessmentPartnerData) *SecurityAssessmentPartnerData {
		return &v
	}).(SecurityAssessmentPartnerDataPtrOutput)
}

func (o SecurityAssessmentPartnerDataOutput) PartnerName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentPartnerData) string { return v.PartnerName }).(pulumi.StringOutput)
}

func (o SecurityAssessmentPartnerDataOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentPartnerData) string { return v.Secret }).(pulumi.StringOutput)
}

type SecurityAssessmentPartnerDataPtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentPartnerDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentPartnerData)(nil)).Elem()
}

func (o SecurityAssessmentPartnerDataPtrOutput) ToSecurityAssessmentPartnerDataPtrOutput() SecurityAssessmentPartnerDataPtrOutput {
	return o
}

func (o SecurityAssessmentPartnerDataPtrOutput) ToSecurityAssessmentPartnerDataPtrOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataPtrOutput {
	return o
}

func (o SecurityAssessmentPartnerDataPtrOutput) Elem() SecurityAssessmentPartnerDataOutput {
	return o.ApplyT(func(v *SecurityAssessmentPartnerData) SecurityAssessmentPartnerData {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentPartnerData
		return ret
	}).(SecurityAssessmentPartnerDataOutput)
}

func (o SecurityAssessmentPartnerDataPtrOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentPartnerData) *string {
		if v == nil {
			return nil
		}
		return &v.PartnerName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentPartnerDataPtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentPartnerData) *string {
		if v == nil {
			return nil
		}
		return &v.Secret
	}).(pulumi.StringPtrOutput)
}

type SecurityAssessmentPartnerDataResponse struct {
	PartnerName string `pulumi:"partnerName"`
	Secret      string `pulumi:"secret"`
}

type SecurityAssessmentPartnerDataResponseOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentPartnerDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAssessmentPartnerDataResponse)(nil)).Elem()
}

func (o SecurityAssessmentPartnerDataResponseOutput) ToSecurityAssessmentPartnerDataResponseOutput() SecurityAssessmentPartnerDataResponseOutput {
	return o
}

func (o SecurityAssessmentPartnerDataResponseOutput) ToSecurityAssessmentPartnerDataResponseOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataResponseOutput {
	return o
}

func (o SecurityAssessmentPartnerDataResponseOutput) PartnerName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentPartnerDataResponse) string { return v.PartnerName }).(pulumi.StringOutput)
}

func (o SecurityAssessmentPartnerDataResponseOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAssessmentPartnerDataResponse) string { return v.Secret }).(pulumi.StringOutput)
}

type SecurityAssessmentPartnerDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SecurityAssessmentPartnerDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityAssessmentPartnerDataResponse)(nil)).Elem()
}

func (o SecurityAssessmentPartnerDataResponsePtrOutput) ToSecurityAssessmentPartnerDataResponsePtrOutput() SecurityAssessmentPartnerDataResponsePtrOutput {
	return o
}

func (o SecurityAssessmentPartnerDataResponsePtrOutput) ToSecurityAssessmentPartnerDataResponsePtrOutputWithContext(ctx context.Context) SecurityAssessmentPartnerDataResponsePtrOutput {
	return o
}

func (o SecurityAssessmentPartnerDataResponsePtrOutput) Elem() SecurityAssessmentPartnerDataResponseOutput {
	return o.ApplyT(func(v *SecurityAssessmentPartnerDataResponse) SecurityAssessmentPartnerDataResponse {
		if v != nil {
			return *v
		}
		var ret SecurityAssessmentPartnerDataResponse
		return ret
	}).(SecurityAssessmentPartnerDataResponseOutput)
}

func (o SecurityAssessmentPartnerDataResponsePtrOutput) PartnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentPartnerDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PartnerName
	}).(pulumi.StringPtrOutput)
}

func (o SecurityAssessmentPartnerDataResponsePtrOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityAssessmentPartnerDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Secret
	}).(pulumi.StringPtrOutput)
}

type UserRecommendation struct {
	RecommendationAction *string `pulumi:"recommendationAction"`
	Username             *string `pulumi:"username"`
}





type UserRecommendationInput interface {
	pulumi.Input

	ToUserRecommendationOutput() UserRecommendationOutput
	ToUserRecommendationOutputWithContext(context.Context) UserRecommendationOutput
}

type UserRecommendationArgs struct {
	RecommendationAction pulumi.StringPtrInput `pulumi:"recommendationAction"`
	Username             pulumi.StringPtrInput `pulumi:"username"`
}

func (UserRecommendationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserRecommendation)(nil)).Elem()
}

func (i UserRecommendationArgs) ToUserRecommendationOutput() UserRecommendationOutput {
	return i.ToUserRecommendationOutputWithContext(context.Background())
}

func (i UserRecommendationArgs) ToUserRecommendationOutputWithContext(ctx context.Context) UserRecommendationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRecommendationOutput)
}





type UserRecommendationArrayInput interface {
	pulumi.Input

	ToUserRecommendationArrayOutput() UserRecommendationArrayOutput
	ToUserRecommendationArrayOutputWithContext(context.Context) UserRecommendationArrayOutput
}

type UserRecommendationArray []UserRecommendationInput

func (UserRecommendationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserRecommendation)(nil)).Elem()
}

func (i UserRecommendationArray) ToUserRecommendationArrayOutput() UserRecommendationArrayOutput {
	return i.ToUserRecommendationArrayOutputWithContext(context.Background())
}

func (i UserRecommendationArray) ToUserRecommendationArrayOutputWithContext(ctx context.Context) UserRecommendationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserRecommendationArrayOutput)
}

type UserRecommendationOutput struct{ *pulumi.OutputState }

func (UserRecommendationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserRecommendation)(nil)).Elem()
}

func (o UserRecommendationOutput) ToUserRecommendationOutput() UserRecommendationOutput {
	return o
}

func (o UserRecommendationOutput) ToUserRecommendationOutputWithContext(ctx context.Context) UserRecommendationOutput {
	return o
}

func (o UserRecommendationOutput) RecommendationAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserRecommendation) *string { return v.RecommendationAction }).(pulumi.StringPtrOutput)
}

func (o UserRecommendationOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserRecommendation) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type UserRecommendationArrayOutput struct{ *pulumi.OutputState }

func (UserRecommendationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserRecommendation)(nil)).Elem()
}

func (o UserRecommendationArrayOutput) ToUserRecommendationArrayOutput() UserRecommendationArrayOutput {
	return o
}

func (o UserRecommendationArrayOutput) ToUserRecommendationArrayOutputWithContext(ctx context.Context) UserRecommendationArrayOutput {
	return o
}

func (o UserRecommendationArrayOutput) Index(i pulumi.IntInput) UserRecommendationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserRecommendation {
		return vs[0].([]UserRecommendation)[vs[1].(int)]
	}).(UserRecommendationOutput)
}

type UserRecommendationResponse struct {
	RecommendationAction *string `pulumi:"recommendationAction"`
	Username             *string `pulumi:"username"`
}

type UserRecommendationResponseOutput struct{ *pulumi.OutputState }

func (UserRecommendationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserRecommendationResponse)(nil)).Elem()
}

func (o UserRecommendationResponseOutput) ToUserRecommendationResponseOutput() UserRecommendationResponseOutput {
	return o
}

func (o UserRecommendationResponseOutput) ToUserRecommendationResponseOutputWithContext(ctx context.Context) UserRecommendationResponseOutput {
	return o
}

func (o UserRecommendationResponseOutput) RecommendationAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserRecommendationResponse) *string { return v.RecommendationAction }).(pulumi.StringPtrOutput)
}

func (o UserRecommendationResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserRecommendationResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type UserRecommendationResponseArrayOutput struct{ *pulumi.OutputState }

func (UserRecommendationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserRecommendationResponse)(nil)).Elem()
}

func (o UserRecommendationResponseArrayOutput) ToUserRecommendationResponseArrayOutput() UserRecommendationResponseArrayOutput {
	return o
}

func (o UserRecommendationResponseArrayOutput) ToUserRecommendationResponseArrayOutputWithContext(ctx context.Context) UserRecommendationResponseArrayOutput {
	return o
}

func (o UserRecommendationResponseArrayOutput) Index(i pulumi.IntInput) UserRecommendationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserRecommendationResponse {
		return vs[0].([]UserRecommendationResponse)[vs[1].(int)]
	}).(UserRecommendationResponseOutput)
}

type VmRecommendation struct {
	ConfigurationStatus  *string `pulumi:"configurationStatus"`
	EnforcementSupport   *string `pulumi:"enforcementSupport"`
	RecommendationAction *string `pulumi:"recommendationAction"`
	ResourceId           *string `pulumi:"resourceId"`
}





type VmRecommendationInput interface {
	pulumi.Input

	ToVmRecommendationOutput() VmRecommendationOutput
	ToVmRecommendationOutputWithContext(context.Context) VmRecommendationOutput
}

type VmRecommendationArgs struct {
	ConfigurationStatus  pulumi.StringPtrInput `pulumi:"configurationStatus"`
	EnforcementSupport   pulumi.StringPtrInput `pulumi:"enforcementSupport"`
	RecommendationAction pulumi.StringPtrInput `pulumi:"recommendationAction"`
	ResourceId           pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (VmRecommendationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmRecommendation)(nil)).Elem()
}

func (i VmRecommendationArgs) ToVmRecommendationOutput() VmRecommendationOutput {
	return i.ToVmRecommendationOutputWithContext(context.Background())
}

func (i VmRecommendationArgs) ToVmRecommendationOutputWithContext(ctx context.Context) VmRecommendationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmRecommendationOutput)
}





type VmRecommendationArrayInput interface {
	pulumi.Input

	ToVmRecommendationArrayOutput() VmRecommendationArrayOutput
	ToVmRecommendationArrayOutputWithContext(context.Context) VmRecommendationArrayOutput
}

type VmRecommendationArray []VmRecommendationInput

func (VmRecommendationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmRecommendation)(nil)).Elem()
}

func (i VmRecommendationArray) ToVmRecommendationArrayOutput() VmRecommendationArrayOutput {
	return i.ToVmRecommendationArrayOutputWithContext(context.Background())
}

func (i VmRecommendationArray) ToVmRecommendationArrayOutputWithContext(ctx context.Context) VmRecommendationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmRecommendationArrayOutput)
}

type VmRecommendationOutput struct{ *pulumi.OutputState }

func (VmRecommendationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmRecommendation)(nil)).Elem()
}

func (o VmRecommendationOutput) ToVmRecommendationOutput() VmRecommendationOutput {
	return o
}

func (o VmRecommendationOutput) ToVmRecommendationOutputWithContext(ctx context.Context) VmRecommendationOutput {
	return o
}

func (o VmRecommendationOutput) ConfigurationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmRecommendation) *string { return v.ConfigurationStatus }).(pulumi.StringPtrOutput)
}

func (o VmRecommendationOutput) EnforcementSupport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmRecommendation) *string { return v.EnforcementSupport }).(pulumi.StringPtrOutput)
}

func (o VmRecommendationOutput) RecommendationAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmRecommendation) *string { return v.RecommendationAction }).(pulumi.StringPtrOutput)
}

func (o VmRecommendationOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmRecommendation) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type VmRecommendationArrayOutput struct{ *pulumi.OutputState }

func (VmRecommendationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmRecommendation)(nil)).Elem()
}

func (o VmRecommendationArrayOutput) ToVmRecommendationArrayOutput() VmRecommendationArrayOutput {
	return o
}

func (o VmRecommendationArrayOutput) ToVmRecommendationArrayOutputWithContext(ctx context.Context) VmRecommendationArrayOutput {
	return o
}

func (o VmRecommendationArrayOutput) Index(i pulumi.IntInput) VmRecommendationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VmRecommendation {
		return vs[0].([]VmRecommendation)[vs[1].(int)]
	}).(VmRecommendationOutput)
}

type VmRecommendationResponse struct {
	ConfigurationStatus  *string `pulumi:"configurationStatus"`
	EnforcementSupport   *string `pulumi:"enforcementSupport"`
	RecommendationAction *string `pulumi:"recommendationAction"`
	ResourceId           *string `pulumi:"resourceId"`
}

type VmRecommendationResponseOutput struct{ *pulumi.OutputState }

func (VmRecommendationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmRecommendationResponse)(nil)).Elem()
}

func (o VmRecommendationResponseOutput) ToVmRecommendationResponseOutput() VmRecommendationResponseOutput {
	return o
}

func (o VmRecommendationResponseOutput) ToVmRecommendationResponseOutputWithContext(ctx context.Context) VmRecommendationResponseOutput {
	return o
}

func (o VmRecommendationResponseOutput) ConfigurationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmRecommendationResponse) *string { return v.ConfigurationStatus }).(pulumi.StringPtrOutput)
}

func (o VmRecommendationResponseOutput) EnforcementSupport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmRecommendationResponse) *string { return v.EnforcementSupport }).(pulumi.StringPtrOutput)
}

func (o VmRecommendationResponseOutput) RecommendationAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmRecommendationResponse) *string { return v.RecommendationAction }).(pulumi.StringPtrOutput)
}

func (o VmRecommendationResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmRecommendationResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type VmRecommendationResponseArrayOutput struct{ *pulumi.OutputState }

func (VmRecommendationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmRecommendationResponse)(nil)).Elem()
}

func (o VmRecommendationResponseArrayOutput) ToVmRecommendationResponseArrayOutput() VmRecommendationResponseArrayOutput {
	return o
}

func (o VmRecommendationResponseArrayOutput) ToVmRecommendationResponseArrayOutputWithContext(ctx context.Context) VmRecommendationResponseArrayOutput {
	return o
}

func (o VmRecommendationResponseArrayOutput) Index(i pulumi.IntInput) VmRecommendationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VmRecommendationResponse {
		return vs[0].([]VmRecommendationResponse)[vs[1].(int)]
	}).(VmRecommendationResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AdaptiveApplicationControlIssueSummaryResponseOutput{})
	pulumi.RegisterOutputType(AdaptiveApplicationControlIssueSummaryResponseArrayOutput{})
	pulumi.RegisterOutputType(AssessmentLinksResponseOutput{})
	pulumi.RegisterOutputType(AssessmentStatusOutput{})
	pulumi.RegisterOutputType(AssessmentStatusResponseOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessPolicyVirtualMachineOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessPolicyVirtualMachineArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessPolicyVirtualMachineResponseOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessPolicyVirtualMachineResponseArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessPortRuleOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessPortRuleArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessPortRuleResponseOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessPortRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestPortOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestPortArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestPortResponseOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestPortResponseArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestResponseOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestResponseArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestVirtualMachineOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestVirtualMachineArrayOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestVirtualMachineResponseOutput{})
	pulumi.RegisterOutputType(JitNetworkAccessRequestVirtualMachineResponseArrayOutput{})
	pulumi.RegisterOutputType(PathRecommendationOutput{})
	pulumi.RegisterOutputType(PathRecommendationArrayOutput{})
	pulumi.RegisterOutputType(PathRecommendationResponseOutput{})
	pulumi.RegisterOutputType(PathRecommendationResponseArrayOutput{})
	pulumi.RegisterOutputType(ProtectionModeOutput{})
	pulumi.RegisterOutputType(ProtectionModePtrOutput{})
	pulumi.RegisterOutputType(ProtectionModeResponseOutput{})
	pulumi.RegisterOutputType(ProtectionModeResponsePtrOutput{})
	pulumi.RegisterOutputType(PublisherInfoOutput{})
	pulumi.RegisterOutputType(PublisherInfoPtrOutput{})
	pulumi.RegisterOutputType(PublisherInfoResponseOutput{})
	pulumi.RegisterOutputType(PublisherInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPartnerDataOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPartnerDataPtrOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPartnerDataResponseOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPartnerDataResponsePtrOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPropertiesOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentMetadataPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentPartnerDataOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentPartnerDataPtrOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentPartnerDataResponseOutput{})
	pulumi.RegisterOutputType(SecurityAssessmentPartnerDataResponsePtrOutput{})
	pulumi.RegisterOutputType(UserRecommendationOutput{})
	pulumi.RegisterOutputType(UserRecommendationArrayOutput{})
	pulumi.RegisterOutputType(UserRecommendationResponseOutput{})
	pulumi.RegisterOutputType(UserRecommendationResponseArrayOutput{})
	pulumi.RegisterOutputType(VmRecommendationOutput{})
	pulumi.RegisterOutputType(VmRecommendationArrayOutput{})
	pulumi.RegisterOutputType(VmRecommendationResponseOutput{})
	pulumi.RegisterOutputType(VmRecommendationResponseArrayOutput{})
}
