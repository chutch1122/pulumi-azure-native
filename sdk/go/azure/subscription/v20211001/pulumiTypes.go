


package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PutAliasRequestAdditionalProperties struct {
	ManagementGroupId    *string           `pulumi:"managementGroupId"`
	SubscriptionOwnerId  *string           `pulumi:"subscriptionOwnerId"`
	SubscriptionTenantId *string           `pulumi:"subscriptionTenantId"`
	Tags                 map[string]string `pulumi:"tags"`
}





type PutAliasRequestAdditionalPropertiesInput interface {
	pulumi.Input

	ToPutAliasRequestAdditionalPropertiesOutput() PutAliasRequestAdditionalPropertiesOutput
	ToPutAliasRequestAdditionalPropertiesOutputWithContext(context.Context) PutAliasRequestAdditionalPropertiesOutput
}

type PutAliasRequestAdditionalPropertiesArgs struct {
	ManagementGroupId    pulumi.StringPtrInput `pulumi:"managementGroupId"`
	SubscriptionOwnerId  pulumi.StringPtrInput `pulumi:"subscriptionOwnerId"`
	SubscriptionTenantId pulumi.StringPtrInput `pulumi:"subscriptionTenantId"`
	Tags                 pulumi.StringMapInput `pulumi:"tags"`
}

func (PutAliasRequestAdditionalPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PutAliasRequestAdditionalProperties)(nil)).Elem()
}

func (i PutAliasRequestAdditionalPropertiesArgs) ToPutAliasRequestAdditionalPropertiesOutput() PutAliasRequestAdditionalPropertiesOutput {
	return i.ToPutAliasRequestAdditionalPropertiesOutputWithContext(context.Background())
}

func (i PutAliasRequestAdditionalPropertiesArgs) ToPutAliasRequestAdditionalPropertiesOutputWithContext(ctx context.Context) PutAliasRequestAdditionalPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasRequestAdditionalPropertiesOutput)
}

func (i PutAliasRequestAdditionalPropertiesArgs) ToPutAliasRequestAdditionalPropertiesPtrOutput() PutAliasRequestAdditionalPropertiesPtrOutput {
	return i.ToPutAliasRequestAdditionalPropertiesPtrOutputWithContext(context.Background())
}

func (i PutAliasRequestAdditionalPropertiesArgs) ToPutAliasRequestAdditionalPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestAdditionalPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasRequestAdditionalPropertiesOutput).ToPutAliasRequestAdditionalPropertiesPtrOutputWithContext(ctx)
}









type PutAliasRequestAdditionalPropertiesPtrInput interface {
	pulumi.Input

	ToPutAliasRequestAdditionalPropertiesPtrOutput() PutAliasRequestAdditionalPropertiesPtrOutput
	ToPutAliasRequestAdditionalPropertiesPtrOutputWithContext(context.Context) PutAliasRequestAdditionalPropertiesPtrOutput
}

type putAliasRequestAdditionalPropertiesPtrType PutAliasRequestAdditionalPropertiesArgs

func PutAliasRequestAdditionalPropertiesPtr(v *PutAliasRequestAdditionalPropertiesArgs) PutAliasRequestAdditionalPropertiesPtrInput {
	return (*putAliasRequestAdditionalPropertiesPtrType)(v)
}

func (*putAliasRequestAdditionalPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PutAliasRequestAdditionalProperties)(nil)).Elem()
}

func (i *putAliasRequestAdditionalPropertiesPtrType) ToPutAliasRequestAdditionalPropertiesPtrOutput() PutAliasRequestAdditionalPropertiesPtrOutput {
	return i.ToPutAliasRequestAdditionalPropertiesPtrOutputWithContext(context.Background())
}

func (i *putAliasRequestAdditionalPropertiesPtrType) ToPutAliasRequestAdditionalPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestAdditionalPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasRequestAdditionalPropertiesPtrOutput)
}

type PutAliasRequestAdditionalPropertiesOutput struct{ *pulumi.OutputState }

func (PutAliasRequestAdditionalPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PutAliasRequestAdditionalProperties)(nil)).Elem()
}

func (o PutAliasRequestAdditionalPropertiesOutput) ToPutAliasRequestAdditionalPropertiesOutput() PutAliasRequestAdditionalPropertiesOutput {
	return o
}

func (o PutAliasRequestAdditionalPropertiesOutput) ToPutAliasRequestAdditionalPropertiesOutputWithContext(ctx context.Context) PutAliasRequestAdditionalPropertiesOutput {
	return o
}

func (o PutAliasRequestAdditionalPropertiesOutput) ToPutAliasRequestAdditionalPropertiesPtrOutput() PutAliasRequestAdditionalPropertiesPtrOutput {
	return o.ToPutAliasRequestAdditionalPropertiesPtrOutputWithContext(context.Background())
}

func (o PutAliasRequestAdditionalPropertiesOutput) ToPutAliasRequestAdditionalPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestAdditionalPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PutAliasRequestAdditionalProperties) *PutAliasRequestAdditionalProperties {
		return &v
	}).(PutAliasRequestAdditionalPropertiesPtrOutput)
}

func (o PutAliasRequestAdditionalPropertiesOutput) ManagementGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestAdditionalProperties) *string { return v.ManagementGroupId }).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestAdditionalPropertiesOutput) SubscriptionOwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestAdditionalProperties) *string { return v.SubscriptionOwnerId }).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestAdditionalPropertiesOutput) SubscriptionTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestAdditionalProperties) *string { return v.SubscriptionTenantId }).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestAdditionalPropertiesOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v PutAliasRequestAdditionalProperties) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type PutAliasRequestAdditionalPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PutAliasRequestAdditionalPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PutAliasRequestAdditionalProperties)(nil)).Elem()
}

func (o PutAliasRequestAdditionalPropertiesPtrOutput) ToPutAliasRequestAdditionalPropertiesPtrOutput() PutAliasRequestAdditionalPropertiesPtrOutput {
	return o
}

func (o PutAliasRequestAdditionalPropertiesPtrOutput) ToPutAliasRequestAdditionalPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestAdditionalPropertiesPtrOutput {
	return o
}

func (o PutAliasRequestAdditionalPropertiesPtrOutput) Elem() PutAliasRequestAdditionalPropertiesOutput {
	return o.ApplyT(func(v *PutAliasRequestAdditionalProperties) PutAliasRequestAdditionalProperties {
		if v != nil {
			return *v
		}
		var ret PutAliasRequestAdditionalProperties
		return ret
	}).(PutAliasRequestAdditionalPropertiesOutput)
}

func (o PutAliasRequestAdditionalPropertiesPtrOutput) ManagementGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestAdditionalProperties) *string {
		if v == nil {
			return nil
		}
		return v.ManagementGroupId
	}).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestAdditionalPropertiesPtrOutput) SubscriptionOwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestAdditionalProperties) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionOwnerId
	}).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestAdditionalPropertiesPtrOutput) SubscriptionTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestAdditionalProperties) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionTenantId
	}).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestAdditionalPropertiesPtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PutAliasRequestAdditionalProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

type PutAliasRequestProperties struct {
	AdditionalProperties *PutAliasRequestAdditionalProperties `pulumi:"additionalProperties"`
	BillingScope         *string                              `pulumi:"billingScope"`
	DisplayName          *string                              `pulumi:"displayName"`
	ResellerId           *string                              `pulumi:"resellerId"`
	SubscriptionId       *string                              `pulumi:"subscriptionId"`
	Workload             *string                              `pulumi:"workload"`
}





type PutAliasRequestPropertiesInput interface {
	pulumi.Input

	ToPutAliasRequestPropertiesOutput() PutAliasRequestPropertiesOutput
	ToPutAliasRequestPropertiesOutputWithContext(context.Context) PutAliasRequestPropertiesOutput
}

type PutAliasRequestPropertiesArgs struct {
	AdditionalProperties PutAliasRequestAdditionalPropertiesPtrInput `pulumi:"additionalProperties"`
	BillingScope         pulumi.StringPtrInput                       `pulumi:"billingScope"`
	DisplayName          pulumi.StringPtrInput                       `pulumi:"displayName"`
	ResellerId           pulumi.StringPtrInput                       `pulumi:"resellerId"`
	SubscriptionId       pulumi.StringPtrInput                       `pulumi:"subscriptionId"`
	Workload             pulumi.StringPtrInput                       `pulumi:"workload"`
}

func (PutAliasRequestPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PutAliasRequestProperties)(nil)).Elem()
}

func (i PutAliasRequestPropertiesArgs) ToPutAliasRequestPropertiesOutput() PutAliasRequestPropertiesOutput {
	return i.ToPutAliasRequestPropertiesOutputWithContext(context.Background())
}

func (i PutAliasRequestPropertiesArgs) ToPutAliasRequestPropertiesOutputWithContext(ctx context.Context) PutAliasRequestPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasRequestPropertiesOutput)
}

func (i PutAliasRequestPropertiesArgs) ToPutAliasRequestPropertiesPtrOutput() PutAliasRequestPropertiesPtrOutput {
	return i.ToPutAliasRequestPropertiesPtrOutputWithContext(context.Background())
}

func (i PutAliasRequestPropertiesArgs) ToPutAliasRequestPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasRequestPropertiesOutput).ToPutAliasRequestPropertiesPtrOutputWithContext(ctx)
}









type PutAliasRequestPropertiesPtrInput interface {
	pulumi.Input

	ToPutAliasRequestPropertiesPtrOutput() PutAliasRequestPropertiesPtrOutput
	ToPutAliasRequestPropertiesPtrOutputWithContext(context.Context) PutAliasRequestPropertiesPtrOutput
}

type putAliasRequestPropertiesPtrType PutAliasRequestPropertiesArgs

func PutAliasRequestPropertiesPtr(v *PutAliasRequestPropertiesArgs) PutAliasRequestPropertiesPtrInput {
	return (*putAliasRequestPropertiesPtrType)(v)
}

func (*putAliasRequestPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PutAliasRequestProperties)(nil)).Elem()
}

func (i *putAliasRequestPropertiesPtrType) ToPutAliasRequestPropertiesPtrOutput() PutAliasRequestPropertiesPtrOutput {
	return i.ToPutAliasRequestPropertiesPtrOutputWithContext(context.Background())
}

func (i *putAliasRequestPropertiesPtrType) ToPutAliasRequestPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PutAliasRequestPropertiesPtrOutput)
}

type PutAliasRequestPropertiesOutput struct{ *pulumi.OutputState }

func (PutAliasRequestPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PutAliasRequestProperties)(nil)).Elem()
}

func (o PutAliasRequestPropertiesOutput) ToPutAliasRequestPropertiesOutput() PutAliasRequestPropertiesOutput {
	return o
}

func (o PutAliasRequestPropertiesOutput) ToPutAliasRequestPropertiesOutputWithContext(ctx context.Context) PutAliasRequestPropertiesOutput {
	return o
}

func (o PutAliasRequestPropertiesOutput) ToPutAliasRequestPropertiesPtrOutput() PutAliasRequestPropertiesPtrOutput {
	return o.ToPutAliasRequestPropertiesPtrOutputWithContext(context.Background())
}

func (o PutAliasRequestPropertiesOutput) ToPutAliasRequestPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PutAliasRequestProperties) *PutAliasRequestProperties {
		return &v
	}).(PutAliasRequestPropertiesPtrOutput)
}

func (o PutAliasRequestPropertiesOutput) AdditionalProperties() PutAliasRequestAdditionalPropertiesPtrOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) *PutAliasRequestAdditionalProperties { return v.AdditionalProperties }).(PutAliasRequestAdditionalPropertiesPtrOutput)
}

func (o PutAliasRequestPropertiesOutput) BillingScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) *string { return v.BillingScope }).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestPropertiesOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestPropertiesOutput) ResellerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) *string { return v.ResellerId }).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestPropertiesOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestPropertiesOutput) Workload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PutAliasRequestProperties) *string { return v.Workload }).(pulumi.StringPtrOutput)
}

type PutAliasRequestPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PutAliasRequestPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PutAliasRequestProperties)(nil)).Elem()
}

func (o PutAliasRequestPropertiesPtrOutput) ToPutAliasRequestPropertiesPtrOutput() PutAliasRequestPropertiesPtrOutput {
	return o
}

func (o PutAliasRequestPropertiesPtrOutput) ToPutAliasRequestPropertiesPtrOutputWithContext(ctx context.Context) PutAliasRequestPropertiesPtrOutput {
	return o
}

func (o PutAliasRequestPropertiesPtrOutput) Elem() PutAliasRequestPropertiesOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) PutAliasRequestProperties {
		if v != nil {
			return *v
		}
		var ret PutAliasRequestProperties
		return ret
	}).(PutAliasRequestPropertiesOutput)
}

func (o PutAliasRequestPropertiesPtrOutput) AdditionalProperties() PutAliasRequestAdditionalPropertiesPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) *PutAliasRequestAdditionalProperties {
		if v == nil {
			return nil
		}
		return v.AdditionalProperties
	}).(PutAliasRequestAdditionalPropertiesPtrOutput)
}

func (o PutAliasRequestPropertiesPtrOutput) BillingScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.BillingScope
	}).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestPropertiesPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestPropertiesPtrOutput) ResellerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.ResellerId
	}).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestPropertiesPtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o PutAliasRequestPropertiesPtrOutput) Workload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PutAliasRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.Workload
	}).(pulumi.StringPtrOutput)
}

type SubscriptionAliasResponsePropertiesResponse struct {
	AcceptOwnershipState string            `pulumi:"acceptOwnershipState"`
	AcceptOwnershipUrl   string            `pulumi:"acceptOwnershipUrl"`
	BillingScope         *string           `pulumi:"billingScope"`
	CreatedTime          *string           `pulumi:"createdTime"`
	DisplayName          *string           `pulumi:"displayName"`
	ManagementGroupId    *string           `pulumi:"managementGroupId"`
	ProvisioningState    *string           `pulumi:"provisioningState"`
	ResellerId           *string           `pulumi:"resellerId"`
	SubscriptionId       string            `pulumi:"subscriptionId"`
	SubscriptionOwnerId  *string           `pulumi:"subscriptionOwnerId"`
	Tags                 map[string]string `pulumi:"tags"`
	Workload             *string           `pulumi:"workload"`
}

type SubscriptionAliasResponsePropertiesResponseOutput struct{ *pulumi.OutputState }

func (SubscriptionAliasResponsePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionAliasResponsePropertiesResponse)(nil)).Elem()
}

func (o SubscriptionAliasResponsePropertiesResponseOutput) ToSubscriptionAliasResponsePropertiesResponseOutput() SubscriptionAliasResponsePropertiesResponseOutput {
	return o
}

func (o SubscriptionAliasResponsePropertiesResponseOutput) ToSubscriptionAliasResponsePropertiesResponseOutputWithContext(ctx context.Context) SubscriptionAliasResponsePropertiesResponseOutput {
	return o
}

func (o SubscriptionAliasResponsePropertiesResponseOutput) AcceptOwnershipState() pulumi.StringOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) string { return v.AcceptOwnershipState }).(pulumi.StringOutput)
}

func (o SubscriptionAliasResponsePropertiesResponseOutput) AcceptOwnershipUrl() pulumi.StringOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) string { return v.AcceptOwnershipUrl }).(pulumi.StringOutput)
}

func (o SubscriptionAliasResponsePropertiesResponseOutput) BillingScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) *string { return v.BillingScope }).(pulumi.StringPtrOutput)
}

func (o SubscriptionAliasResponsePropertiesResponseOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

func (o SubscriptionAliasResponsePropertiesResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o SubscriptionAliasResponsePropertiesResponseOutput) ManagementGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) *string { return v.ManagementGroupId }).(pulumi.StringPtrOutput)
}

func (o SubscriptionAliasResponsePropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o SubscriptionAliasResponsePropertiesResponseOutput) ResellerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) *string { return v.ResellerId }).(pulumi.StringPtrOutput)
}

func (o SubscriptionAliasResponsePropertiesResponseOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o SubscriptionAliasResponsePropertiesResponseOutput) SubscriptionOwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) *string { return v.SubscriptionOwnerId }).(pulumi.StringPtrOutput)
}

func (o SubscriptionAliasResponsePropertiesResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SubscriptionAliasResponsePropertiesResponseOutput) Workload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionAliasResponsePropertiesResponse) *string { return v.Workload }).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PutAliasRequestAdditionalPropertiesOutput{})
	pulumi.RegisterOutputType(PutAliasRequestAdditionalPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PutAliasRequestPropertiesOutput{})
	pulumi.RegisterOutputType(PutAliasRequestPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SubscriptionAliasResponsePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
