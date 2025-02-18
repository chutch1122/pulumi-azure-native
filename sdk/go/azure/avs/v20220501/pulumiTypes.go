


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AddonArcProperties struct {
	AddonType string  `pulumi:"addonType"`
	VCenter   *string `pulumi:"vCenter"`
}

type AddonArcPropertiesResponse struct {
	AddonType         string  `pulumi:"addonType"`
	ProvisioningState string  `pulumi:"provisioningState"`
	VCenter           *string `pulumi:"vCenter"`
}

type AddonHcxProperties struct {
	AddonType string `pulumi:"addonType"`
	Offer     string `pulumi:"offer"`
}

type AddonHcxPropertiesResponse struct {
	AddonType         string `pulumi:"addonType"`
	Offer             string `pulumi:"offer"`
	ProvisioningState string `pulumi:"provisioningState"`
}

type AddonSrmProperties struct {
	AddonType  string  `pulumi:"addonType"`
	LicenseKey *string `pulumi:"licenseKey"`
}

type AddonSrmPropertiesResponse struct {
	AddonType         string  `pulumi:"addonType"`
	LicenseKey        *string `pulumi:"licenseKey"`
	ProvisioningState string  `pulumi:"provisioningState"`
}

type AddonVrProperties struct {
	AddonType string `pulumi:"addonType"`
	VrsCount  int    `pulumi:"vrsCount"`
}

type AddonVrPropertiesResponse struct {
	AddonType         string `pulumi:"addonType"`
	ProvisioningState string `pulumi:"provisioningState"`
	VrsCount          int    `pulumi:"vrsCount"`
}

type AvailabilityProperties struct {
	SecondaryZone *int    `pulumi:"secondaryZone"`
	Strategy      *string `pulumi:"strategy"`
	Zone          *int    `pulumi:"zone"`
}





type AvailabilityPropertiesInput interface {
	pulumi.Input

	ToAvailabilityPropertiesOutput() AvailabilityPropertiesOutput
	ToAvailabilityPropertiesOutputWithContext(context.Context) AvailabilityPropertiesOutput
}

type AvailabilityPropertiesArgs struct {
	SecondaryZone pulumi.IntPtrInput    `pulumi:"secondaryZone"`
	Strategy      pulumi.StringPtrInput `pulumi:"strategy"`
	Zone          pulumi.IntPtrInput    `pulumi:"zone"`
}

func (AvailabilityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityProperties)(nil)).Elem()
}

func (i AvailabilityPropertiesArgs) ToAvailabilityPropertiesOutput() AvailabilityPropertiesOutput {
	return i.ToAvailabilityPropertiesOutputWithContext(context.Background())
}

func (i AvailabilityPropertiesArgs) ToAvailabilityPropertiesOutputWithContext(ctx context.Context) AvailabilityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityPropertiesOutput)
}

func (i AvailabilityPropertiesArgs) ToAvailabilityPropertiesPtrOutput() AvailabilityPropertiesPtrOutput {
	return i.ToAvailabilityPropertiesPtrOutputWithContext(context.Background())
}

func (i AvailabilityPropertiesArgs) ToAvailabilityPropertiesPtrOutputWithContext(ctx context.Context) AvailabilityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityPropertiesOutput).ToAvailabilityPropertiesPtrOutputWithContext(ctx)
}









type AvailabilityPropertiesPtrInput interface {
	pulumi.Input

	ToAvailabilityPropertiesPtrOutput() AvailabilityPropertiesPtrOutput
	ToAvailabilityPropertiesPtrOutputWithContext(context.Context) AvailabilityPropertiesPtrOutput
}

type availabilityPropertiesPtrType AvailabilityPropertiesArgs

func AvailabilityPropertiesPtr(v *AvailabilityPropertiesArgs) AvailabilityPropertiesPtrInput {
	return (*availabilityPropertiesPtrType)(v)
}

func (*availabilityPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailabilityProperties)(nil)).Elem()
}

func (i *availabilityPropertiesPtrType) ToAvailabilityPropertiesPtrOutput() AvailabilityPropertiesPtrOutput {
	return i.ToAvailabilityPropertiesPtrOutputWithContext(context.Background())
}

func (i *availabilityPropertiesPtrType) ToAvailabilityPropertiesPtrOutputWithContext(ctx context.Context) AvailabilityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityPropertiesPtrOutput)
}

type AvailabilityPropertiesOutput struct{ *pulumi.OutputState }

func (AvailabilityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityProperties)(nil)).Elem()
}

func (o AvailabilityPropertiesOutput) ToAvailabilityPropertiesOutput() AvailabilityPropertiesOutput {
	return o
}

func (o AvailabilityPropertiesOutput) ToAvailabilityPropertiesOutputWithContext(ctx context.Context) AvailabilityPropertiesOutput {
	return o
}

func (o AvailabilityPropertiesOutput) ToAvailabilityPropertiesPtrOutput() AvailabilityPropertiesPtrOutput {
	return o.ToAvailabilityPropertiesPtrOutputWithContext(context.Background())
}

func (o AvailabilityPropertiesOutput) ToAvailabilityPropertiesPtrOutputWithContext(ctx context.Context) AvailabilityPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AvailabilityProperties) *AvailabilityProperties {
		return &v
	}).(AvailabilityPropertiesPtrOutput)
}

func (o AvailabilityPropertiesOutput) SecondaryZone() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AvailabilityProperties) *int { return v.SecondaryZone }).(pulumi.IntPtrOutput)
}

func (o AvailabilityPropertiesOutput) Strategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AvailabilityProperties) *string { return v.Strategy }).(pulumi.StringPtrOutput)
}

func (o AvailabilityPropertiesOutput) Zone() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AvailabilityProperties) *int { return v.Zone }).(pulumi.IntPtrOutput)
}

type AvailabilityPropertiesPtrOutput struct{ *pulumi.OutputState }

func (AvailabilityPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailabilityProperties)(nil)).Elem()
}

func (o AvailabilityPropertiesPtrOutput) ToAvailabilityPropertiesPtrOutput() AvailabilityPropertiesPtrOutput {
	return o
}

func (o AvailabilityPropertiesPtrOutput) ToAvailabilityPropertiesPtrOutputWithContext(ctx context.Context) AvailabilityPropertiesPtrOutput {
	return o
}

func (o AvailabilityPropertiesPtrOutput) Elem() AvailabilityPropertiesOutput {
	return o.ApplyT(func(v *AvailabilityProperties) AvailabilityProperties {
		if v != nil {
			return *v
		}
		var ret AvailabilityProperties
		return ret
	}).(AvailabilityPropertiesOutput)
}

func (o AvailabilityPropertiesPtrOutput) SecondaryZone() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AvailabilityProperties) *int {
		if v == nil {
			return nil
		}
		return v.SecondaryZone
	}).(pulumi.IntPtrOutput)
}

func (o AvailabilityPropertiesPtrOutput) Strategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvailabilityProperties) *string {
		if v == nil {
			return nil
		}
		return v.Strategy
	}).(pulumi.StringPtrOutput)
}

func (o AvailabilityPropertiesPtrOutput) Zone() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AvailabilityProperties) *int {
		if v == nil {
			return nil
		}
		return v.Zone
	}).(pulumi.IntPtrOutput)
}

type AvailabilityPropertiesResponse struct {
	SecondaryZone *int    `pulumi:"secondaryZone"`
	Strategy      *string `pulumi:"strategy"`
	Zone          *int    `pulumi:"zone"`
}

type AvailabilityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AvailabilityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityPropertiesResponse)(nil)).Elem()
}

func (o AvailabilityPropertiesResponseOutput) ToAvailabilityPropertiesResponseOutput() AvailabilityPropertiesResponseOutput {
	return o
}

func (o AvailabilityPropertiesResponseOutput) ToAvailabilityPropertiesResponseOutputWithContext(ctx context.Context) AvailabilityPropertiesResponseOutput {
	return o
}

func (o AvailabilityPropertiesResponseOutput) SecondaryZone() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AvailabilityPropertiesResponse) *int { return v.SecondaryZone }).(pulumi.IntPtrOutput)
}

func (o AvailabilityPropertiesResponseOutput) Strategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AvailabilityPropertiesResponse) *string { return v.Strategy }).(pulumi.StringPtrOutput)
}

func (o AvailabilityPropertiesResponseOutput) Zone() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AvailabilityPropertiesResponse) *int { return v.Zone }).(pulumi.IntPtrOutput)
}

type AvailabilityPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AvailabilityPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailabilityPropertiesResponse)(nil)).Elem()
}

func (o AvailabilityPropertiesResponsePtrOutput) ToAvailabilityPropertiesResponsePtrOutput() AvailabilityPropertiesResponsePtrOutput {
	return o
}

func (o AvailabilityPropertiesResponsePtrOutput) ToAvailabilityPropertiesResponsePtrOutputWithContext(ctx context.Context) AvailabilityPropertiesResponsePtrOutput {
	return o
}

func (o AvailabilityPropertiesResponsePtrOutput) Elem() AvailabilityPropertiesResponseOutput {
	return o.ApplyT(func(v *AvailabilityPropertiesResponse) AvailabilityPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AvailabilityPropertiesResponse
		return ret
	}).(AvailabilityPropertiesResponseOutput)
}

func (o AvailabilityPropertiesResponsePtrOutput) SecondaryZone() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AvailabilityPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.SecondaryZone
	}).(pulumi.IntPtrOutput)
}

func (o AvailabilityPropertiesResponsePtrOutput) Strategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvailabilityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Strategy
	}).(pulumi.StringPtrOutput)
}

func (o AvailabilityPropertiesResponsePtrOutput) Zone() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AvailabilityPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.Zone
	}).(pulumi.IntPtrOutput)
}

type CircuitResponse struct {
	ExpressRouteID               string `pulumi:"expressRouteID"`
	ExpressRoutePrivatePeeringID string `pulumi:"expressRoutePrivatePeeringID"`
	PrimarySubnet                string `pulumi:"primarySubnet"`
	SecondarySubnet              string `pulumi:"secondarySubnet"`
}

type CircuitResponseOutput struct{ *pulumi.OutputState }

func (CircuitResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CircuitResponse)(nil)).Elem()
}

func (o CircuitResponseOutput) ToCircuitResponseOutput() CircuitResponseOutput {
	return o
}

func (o CircuitResponseOutput) ToCircuitResponseOutputWithContext(ctx context.Context) CircuitResponseOutput {
	return o
}

func (o CircuitResponseOutput) ExpressRouteID() pulumi.StringOutput {
	return o.ApplyT(func(v CircuitResponse) string { return v.ExpressRouteID }).(pulumi.StringOutput)
}

func (o CircuitResponseOutput) ExpressRoutePrivatePeeringID() pulumi.StringOutput {
	return o.ApplyT(func(v CircuitResponse) string { return v.ExpressRoutePrivatePeeringID }).(pulumi.StringOutput)
}

func (o CircuitResponseOutput) PrimarySubnet() pulumi.StringOutput {
	return o.ApplyT(func(v CircuitResponse) string { return v.PrimarySubnet }).(pulumi.StringOutput)
}

func (o CircuitResponseOutput) SecondarySubnet() pulumi.StringOutput {
	return o.ApplyT(func(v CircuitResponse) string { return v.SecondarySubnet }).(pulumi.StringOutput)
}

type CircuitResponsePtrOutput struct{ *pulumi.OutputState }

func (CircuitResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CircuitResponse)(nil)).Elem()
}

func (o CircuitResponsePtrOutput) ToCircuitResponsePtrOutput() CircuitResponsePtrOutput {
	return o
}

func (o CircuitResponsePtrOutput) ToCircuitResponsePtrOutputWithContext(ctx context.Context) CircuitResponsePtrOutput {
	return o
}

func (o CircuitResponsePtrOutput) Elem() CircuitResponseOutput {
	return o.ApplyT(func(v *CircuitResponse) CircuitResponse {
		if v != nil {
			return *v
		}
		var ret CircuitResponse
		return ret
	}).(CircuitResponseOutput)
}

func (o CircuitResponsePtrOutput) ExpressRouteID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CircuitResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ExpressRouteID
	}).(pulumi.StringPtrOutput)
}

func (o CircuitResponsePtrOutput) ExpressRoutePrivatePeeringID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CircuitResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ExpressRoutePrivatePeeringID
	}).(pulumi.StringPtrOutput)
}

func (o CircuitResponsePtrOutput) PrimarySubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CircuitResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrimarySubnet
	}).(pulumi.StringPtrOutput)
}

func (o CircuitResponsePtrOutput) SecondarySubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CircuitResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecondarySubnet
	}).(pulumi.StringPtrOutput)
}

type ClusterZoneResponse struct {
	Hosts []string `pulumi:"hosts"`
	Zone  string   `pulumi:"zone"`
}

type ClusterZoneResponseOutput struct{ *pulumi.OutputState }

func (ClusterZoneResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterZoneResponse)(nil)).Elem()
}

func (o ClusterZoneResponseOutput) ToClusterZoneResponseOutput() ClusterZoneResponseOutput {
	return o
}

func (o ClusterZoneResponseOutput) ToClusterZoneResponseOutputWithContext(ctx context.Context) ClusterZoneResponseOutput {
	return o
}

func (o ClusterZoneResponseOutput) Hosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ClusterZoneResponse) []string { return v.Hosts }).(pulumi.StringArrayOutput)
}

func (o ClusterZoneResponseOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterZoneResponse) string { return v.Zone }).(pulumi.StringOutput)
}

type ClusterZoneResponseArrayOutput struct{ *pulumi.OutputState }

func (ClusterZoneResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterZoneResponse)(nil)).Elem()
}

func (o ClusterZoneResponseArrayOutput) ToClusterZoneResponseArrayOutput() ClusterZoneResponseArrayOutput {
	return o
}

func (o ClusterZoneResponseArrayOutput) ToClusterZoneResponseArrayOutputWithContext(ctx context.Context) ClusterZoneResponseArrayOutput {
	return o
}

func (o ClusterZoneResponseArrayOutput) Index(i pulumi.IntInput) ClusterZoneResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterZoneResponse {
		return vs[0].([]ClusterZoneResponse)[vs[1].(int)]
	}).(ClusterZoneResponseOutput)
}

type DiskPoolVolume struct {
	LunName     string  `pulumi:"lunName"`
	MountOption *string `pulumi:"mountOption"`
	TargetId    string  `pulumi:"targetId"`
}


func (val *DiskPoolVolume) Defaults() *DiskPoolVolume {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MountOption) {
		mountOption_ := "MOUNT"
		tmp.MountOption = &mountOption_
	}
	return &tmp
}





type DiskPoolVolumeInput interface {
	pulumi.Input

	ToDiskPoolVolumeOutput() DiskPoolVolumeOutput
	ToDiskPoolVolumeOutputWithContext(context.Context) DiskPoolVolumeOutput
}

type DiskPoolVolumeArgs struct {
	LunName     pulumi.StringInput    `pulumi:"lunName"`
	MountOption pulumi.StringPtrInput `pulumi:"mountOption"`
	TargetId    pulumi.StringInput    `pulumi:"targetId"`
}


func (val *DiskPoolVolumeArgs) Defaults() *DiskPoolVolumeArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MountOption) {
		tmp.MountOption = pulumi.StringPtr("MOUNT")
	}
	return &tmp
}
func (DiskPoolVolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskPoolVolume)(nil)).Elem()
}

func (i DiskPoolVolumeArgs) ToDiskPoolVolumeOutput() DiskPoolVolumeOutput {
	return i.ToDiskPoolVolumeOutputWithContext(context.Background())
}

func (i DiskPoolVolumeArgs) ToDiskPoolVolumeOutputWithContext(ctx context.Context) DiskPoolVolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskPoolVolumeOutput)
}

func (i DiskPoolVolumeArgs) ToDiskPoolVolumePtrOutput() DiskPoolVolumePtrOutput {
	return i.ToDiskPoolVolumePtrOutputWithContext(context.Background())
}

func (i DiskPoolVolumeArgs) ToDiskPoolVolumePtrOutputWithContext(ctx context.Context) DiskPoolVolumePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskPoolVolumeOutput).ToDiskPoolVolumePtrOutputWithContext(ctx)
}









type DiskPoolVolumePtrInput interface {
	pulumi.Input

	ToDiskPoolVolumePtrOutput() DiskPoolVolumePtrOutput
	ToDiskPoolVolumePtrOutputWithContext(context.Context) DiskPoolVolumePtrOutput
}

type diskPoolVolumePtrType DiskPoolVolumeArgs

func DiskPoolVolumePtr(v *DiskPoolVolumeArgs) DiskPoolVolumePtrInput {
	return (*diskPoolVolumePtrType)(v)
}

func (*diskPoolVolumePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskPoolVolume)(nil)).Elem()
}

func (i *diskPoolVolumePtrType) ToDiskPoolVolumePtrOutput() DiskPoolVolumePtrOutput {
	return i.ToDiskPoolVolumePtrOutputWithContext(context.Background())
}

func (i *diskPoolVolumePtrType) ToDiskPoolVolumePtrOutputWithContext(ctx context.Context) DiskPoolVolumePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskPoolVolumePtrOutput)
}

type DiskPoolVolumeOutput struct{ *pulumi.OutputState }

func (DiskPoolVolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskPoolVolume)(nil)).Elem()
}

func (o DiskPoolVolumeOutput) ToDiskPoolVolumeOutput() DiskPoolVolumeOutput {
	return o
}

func (o DiskPoolVolumeOutput) ToDiskPoolVolumeOutputWithContext(ctx context.Context) DiskPoolVolumeOutput {
	return o
}

func (o DiskPoolVolumeOutput) ToDiskPoolVolumePtrOutput() DiskPoolVolumePtrOutput {
	return o.ToDiskPoolVolumePtrOutputWithContext(context.Background())
}

func (o DiskPoolVolumeOutput) ToDiskPoolVolumePtrOutputWithContext(ctx context.Context) DiskPoolVolumePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskPoolVolume) *DiskPoolVolume {
		return &v
	}).(DiskPoolVolumePtrOutput)
}

func (o DiskPoolVolumeOutput) LunName() pulumi.StringOutput {
	return o.ApplyT(func(v DiskPoolVolume) string { return v.LunName }).(pulumi.StringOutput)
}

func (o DiskPoolVolumeOutput) MountOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskPoolVolume) *string { return v.MountOption }).(pulumi.StringPtrOutput)
}

func (o DiskPoolVolumeOutput) TargetId() pulumi.StringOutput {
	return o.ApplyT(func(v DiskPoolVolume) string { return v.TargetId }).(pulumi.StringOutput)
}

type DiskPoolVolumePtrOutput struct{ *pulumi.OutputState }

func (DiskPoolVolumePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskPoolVolume)(nil)).Elem()
}

func (o DiskPoolVolumePtrOutput) ToDiskPoolVolumePtrOutput() DiskPoolVolumePtrOutput {
	return o
}

func (o DiskPoolVolumePtrOutput) ToDiskPoolVolumePtrOutputWithContext(ctx context.Context) DiskPoolVolumePtrOutput {
	return o
}

func (o DiskPoolVolumePtrOutput) Elem() DiskPoolVolumeOutput {
	return o.ApplyT(func(v *DiskPoolVolume) DiskPoolVolume {
		if v != nil {
			return *v
		}
		var ret DiskPoolVolume
		return ret
	}).(DiskPoolVolumeOutput)
}

func (o DiskPoolVolumePtrOutput) LunName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskPoolVolume) *string {
		if v == nil {
			return nil
		}
		return &v.LunName
	}).(pulumi.StringPtrOutput)
}

func (o DiskPoolVolumePtrOutput) MountOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskPoolVolume) *string {
		if v == nil {
			return nil
		}
		return v.MountOption
	}).(pulumi.StringPtrOutput)
}

func (o DiskPoolVolumePtrOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskPoolVolume) *string {
		if v == nil {
			return nil
		}
		return &v.TargetId
	}).(pulumi.StringPtrOutput)
}

type DiskPoolVolumeResponse struct {
	LunName     string  `pulumi:"lunName"`
	MountOption *string `pulumi:"mountOption"`
	Path        string  `pulumi:"path"`
	TargetId    string  `pulumi:"targetId"`
}


func (val *DiskPoolVolumeResponse) Defaults() *DiskPoolVolumeResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MountOption) {
		mountOption_ := "MOUNT"
		tmp.MountOption = &mountOption_
	}
	return &tmp
}

type DiskPoolVolumeResponseOutput struct{ *pulumi.OutputState }

func (DiskPoolVolumeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskPoolVolumeResponse)(nil)).Elem()
}

func (o DiskPoolVolumeResponseOutput) ToDiskPoolVolumeResponseOutput() DiskPoolVolumeResponseOutput {
	return o
}

func (o DiskPoolVolumeResponseOutput) ToDiskPoolVolumeResponseOutputWithContext(ctx context.Context) DiskPoolVolumeResponseOutput {
	return o
}

func (o DiskPoolVolumeResponseOutput) LunName() pulumi.StringOutput {
	return o.ApplyT(func(v DiskPoolVolumeResponse) string { return v.LunName }).(pulumi.StringOutput)
}

func (o DiskPoolVolumeResponseOutput) MountOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskPoolVolumeResponse) *string { return v.MountOption }).(pulumi.StringPtrOutput)
}

func (o DiskPoolVolumeResponseOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v DiskPoolVolumeResponse) string { return v.Path }).(pulumi.StringOutput)
}

func (o DiskPoolVolumeResponseOutput) TargetId() pulumi.StringOutput {
	return o.ApplyT(func(v DiskPoolVolumeResponse) string { return v.TargetId }).(pulumi.StringOutput)
}

type DiskPoolVolumeResponsePtrOutput struct{ *pulumi.OutputState }

func (DiskPoolVolumeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskPoolVolumeResponse)(nil)).Elem()
}

func (o DiskPoolVolumeResponsePtrOutput) ToDiskPoolVolumeResponsePtrOutput() DiskPoolVolumeResponsePtrOutput {
	return o
}

func (o DiskPoolVolumeResponsePtrOutput) ToDiskPoolVolumeResponsePtrOutputWithContext(ctx context.Context) DiskPoolVolumeResponsePtrOutput {
	return o
}

func (o DiskPoolVolumeResponsePtrOutput) Elem() DiskPoolVolumeResponseOutput {
	return o.ApplyT(func(v *DiskPoolVolumeResponse) DiskPoolVolumeResponse {
		if v != nil {
			return *v
		}
		var ret DiskPoolVolumeResponse
		return ret
	}).(DiskPoolVolumeResponseOutput)
}

func (o DiskPoolVolumeResponsePtrOutput) LunName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskPoolVolumeResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LunName
	}).(pulumi.StringPtrOutput)
}

func (o DiskPoolVolumeResponsePtrOutput) MountOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskPoolVolumeResponse) *string {
		if v == nil {
			return nil
		}
		return v.MountOption
	}).(pulumi.StringPtrOutput)
}

func (o DiskPoolVolumeResponsePtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskPoolVolumeResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Path
	}).(pulumi.StringPtrOutput)
}

func (o DiskPoolVolumeResponsePtrOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskPoolVolumeResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TargetId
	}).(pulumi.StringPtrOutput)
}

type Encryption struct {
	KeyVaultProperties *EncryptionKeyVaultProperties `pulumi:"keyVaultProperties"`
	Status             *string                       `pulumi:"status"`
}





type EncryptionInput interface {
	pulumi.Input

	ToEncryptionOutput() EncryptionOutput
	ToEncryptionOutputWithContext(context.Context) EncryptionOutput
}

type EncryptionArgs struct {
	KeyVaultProperties EncryptionKeyVaultPropertiesPtrInput `pulumi:"keyVaultProperties"`
	Status             pulumi.StringPtrInput                `pulumi:"status"`
}

func (EncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (i EncryptionArgs) ToEncryptionOutput() EncryptionOutput {
	return i.ToEncryptionOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput)
}

func (i EncryptionArgs) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput).ToEncryptionPtrOutputWithContext(ctx)
}









type EncryptionPtrInput interface {
	pulumi.Input

	ToEncryptionPtrOutput() EncryptionPtrOutput
	ToEncryptionPtrOutputWithContext(context.Context) EncryptionPtrOutput
}

type encryptionPtrType EncryptionArgs

func EncryptionPtr(v *EncryptionArgs) EncryptionPtrInput {
	return (*encryptionPtrType)(v)
}

func (*encryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (i *encryptionPtrType) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i *encryptionPtrType) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPtrOutput)
}

type EncryptionOutput struct{ *pulumi.OutputState }

func (EncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (o EncryptionOutput) ToEncryptionOutput() EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o.ToEncryptionPtrOutputWithContext(context.Background())
}

func (o EncryptionOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Encryption) *Encryption {
		return &v
	}).(EncryptionPtrOutput)
}

func (o EncryptionOutput) KeyVaultProperties() EncryptionKeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v Encryption) *EncryptionKeyVaultProperties { return v.KeyVaultProperties }).(EncryptionKeyVaultPropertiesPtrOutput)
}

func (o EncryptionOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Encryption) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type EncryptionPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) Elem() EncryptionOutput {
	return o.ApplyT(func(v *Encryption) Encryption {
		if v != nil {
			return *v
		}
		var ret Encryption
		return ret
	}).(EncryptionOutput)
}

func (o EncryptionPtrOutput) KeyVaultProperties() EncryptionKeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *Encryption) *EncryptionKeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(EncryptionKeyVaultPropertiesPtrOutput)
}

func (o EncryptionPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Encryption) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type EncryptionKeyVaultProperties struct {
	KeyName     *string `pulumi:"keyName"`
	KeyVaultUrl *string `pulumi:"keyVaultUrl"`
	KeyVersion  *string `pulumi:"keyVersion"`
}





type EncryptionKeyVaultPropertiesInput interface {
	pulumi.Input

	ToEncryptionKeyVaultPropertiesOutput() EncryptionKeyVaultPropertiesOutput
	ToEncryptionKeyVaultPropertiesOutputWithContext(context.Context) EncryptionKeyVaultPropertiesOutput
}

type EncryptionKeyVaultPropertiesArgs struct {
	KeyName     pulumi.StringPtrInput `pulumi:"keyName"`
	KeyVaultUrl pulumi.StringPtrInput `pulumi:"keyVaultUrl"`
	KeyVersion  pulumi.StringPtrInput `pulumi:"keyVersion"`
}

func (EncryptionKeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeyVaultProperties)(nil)).Elem()
}

func (i EncryptionKeyVaultPropertiesArgs) ToEncryptionKeyVaultPropertiesOutput() EncryptionKeyVaultPropertiesOutput {
	return i.ToEncryptionKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i EncryptionKeyVaultPropertiesArgs) ToEncryptionKeyVaultPropertiesOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyVaultPropertiesOutput)
}

func (i EncryptionKeyVaultPropertiesArgs) ToEncryptionKeyVaultPropertiesPtrOutput() EncryptionKeyVaultPropertiesPtrOutput {
	return i.ToEncryptionKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i EncryptionKeyVaultPropertiesArgs) ToEncryptionKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyVaultPropertiesOutput).ToEncryptionKeyVaultPropertiesPtrOutputWithContext(ctx)
}









type EncryptionKeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToEncryptionKeyVaultPropertiesPtrOutput() EncryptionKeyVaultPropertiesPtrOutput
	ToEncryptionKeyVaultPropertiesPtrOutputWithContext(context.Context) EncryptionKeyVaultPropertiesPtrOutput
}

type encryptionKeyVaultPropertiesPtrType EncryptionKeyVaultPropertiesArgs

func EncryptionKeyVaultPropertiesPtr(v *EncryptionKeyVaultPropertiesArgs) EncryptionKeyVaultPropertiesPtrInput {
	return (*encryptionKeyVaultPropertiesPtrType)(v)
}

func (*encryptionKeyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionKeyVaultProperties)(nil)).Elem()
}

func (i *encryptionKeyVaultPropertiesPtrType) ToEncryptionKeyVaultPropertiesPtrOutput() EncryptionKeyVaultPropertiesPtrOutput {
	return i.ToEncryptionKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *encryptionKeyVaultPropertiesPtrType) ToEncryptionKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionKeyVaultPropertiesPtrOutput)
}

type EncryptionKeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (EncryptionKeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeyVaultProperties)(nil)).Elem()
}

func (o EncryptionKeyVaultPropertiesOutput) ToEncryptionKeyVaultPropertiesOutput() EncryptionKeyVaultPropertiesOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesOutput) ToEncryptionKeyVaultPropertiesOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesOutput) ToEncryptionKeyVaultPropertiesPtrOutput() EncryptionKeyVaultPropertiesPtrOutput {
	return o.ToEncryptionKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o EncryptionKeyVaultPropertiesOutput) ToEncryptionKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionKeyVaultProperties) *EncryptionKeyVaultProperties {
		return &v
	}).(EncryptionKeyVaultPropertiesPtrOutput)
}

func (o EncryptionKeyVaultPropertiesOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyVaultProperties) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesOutput) KeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyVaultProperties) *string { return v.KeyVaultUrl }).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyVaultProperties) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

type EncryptionKeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionKeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionKeyVaultProperties)(nil)).Elem()
}

func (o EncryptionKeyVaultPropertiesPtrOutput) ToEncryptionKeyVaultPropertiesPtrOutput() EncryptionKeyVaultPropertiesPtrOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesPtrOutput) ToEncryptionKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesPtrOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesPtrOutput) Elem() EncryptionKeyVaultPropertiesOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultProperties) EncryptionKeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret EncryptionKeyVaultProperties
		return ret
	}).(EncryptionKeyVaultPropertiesOutput)
}

func (o EncryptionKeyVaultPropertiesPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesPtrOutput) KeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUrl
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

type EncryptionKeyVaultPropertiesResponse struct {
	AutoDetectedKeyVersion string  `pulumi:"autoDetectedKeyVersion"`
	KeyName                *string `pulumi:"keyName"`
	KeyState               string  `pulumi:"keyState"`
	KeyVaultUrl            *string `pulumi:"keyVaultUrl"`
	KeyVersion             *string `pulumi:"keyVersion"`
	VersionType            string  `pulumi:"versionType"`
}

type EncryptionKeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EncryptionKeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionKeyVaultPropertiesResponse)(nil)).Elem()
}

func (o EncryptionKeyVaultPropertiesResponseOutput) ToEncryptionKeyVaultPropertiesResponseOutput() EncryptionKeyVaultPropertiesResponseOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesResponseOutput) ToEncryptionKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesResponseOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesResponseOutput) AutoDetectedKeyVersion() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionKeyVaultPropertiesResponse) string { return v.AutoDetectedKeyVersion }).(pulumi.StringOutput)
}

func (o EncryptionKeyVaultPropertiesResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyVaultPropertiesResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponseOutput) KeyState() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionKeyVaultPropertiesResponse) string { return v.KeyState }).(pulumi.StringOutput)
}

func (o EncryptionKeyVaultPropertiesResponseOutput) KeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyVaultPropertiesResponse) *string { return v.KeyVaultUrl }).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionKeyVaultPropertiesResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponseOutput) VersionType() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionKeyVaultPropertiesResponse) string { return v.VersionType }).(pulumi.StringOutput)
}

type EncryptionKeyVaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionKeyVaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionKeyVaultPropertiesResponse)(nil)).Elem()
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) ToEncryptionKeyVaultPropertiesResponsePtrOutput() EncryptionKeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) ToEncryptionKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) EncryptionKeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) Elem() EncryptionKeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultPropertiesResponse) EncryptionKeyVaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionKeyVaultPropertiesResponse
		return ret
	}).(EncryptionKeyVaultPropertiesResponseOutput)
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) AutoDetectedKeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AutoDetectedKeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) KeyState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyState
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) KeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultUrl
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionKeyVaultPropertiesResponsePtrOutput) VersionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionKeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VersionType
	}).(pulumi.StringPtrOutput)
}

type EncryptionResponse struct {
	KeyVaultProperties *EncryptionKeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	Status             *string                               `pulumi:"status"`
}

type EncryptionResponseOutput struct{ *pulumi.OutputState }

func (EncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) KeyVaultProperties() EncryptionKeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *EncryptionKeyVaultPropertiesResponse { return v.KeyVaultProperties }).(EncryptionKeyVaultPropertiesResponsePtrOutput)
}

func (o EncryptionResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type EncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) Elem() EncryptionResponseOutput {
	return o.ApplyT(func(v *EncryptionResponse) EncryptionResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionResponse
		return ret
	}).(EncryptionResponseOutput)
}

func (o EncryptionResponsePtrOutput) KeyVaultProperties() EncryptionKeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *EncryptionKeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(EncryptionKeyVaultPropertiesResponsePtrOutput)
}

func (o EncryptionResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type EndpointsResponse struct {
	HcxCloudManager string `pulumi:"hcxCloudManager"`
	NsxtManager     string `pulumi:"nsxtManager"`
	Vcsa            string `pulumi:"vcsa"`
}

type EndpointsResponseOutput struct{ *pulumi.OutputState }

func (EndpointsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointsResponse)(nil)).Elem()
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutput() EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutputWithContext(ctx context.Context) EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) HcxCloudManager() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.HcxCloudManager }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) NsxtManager() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.NsxtManager }).(pulumi.StringOutput)
}

func (o EndpointsResponseOutput) Vcsa() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Vcsa }).(pulumi.StringOutput)
}

type IdentitySource struct {
	Alias           *string `pulumi:"alias"`
	BaseGroupDN     *string `pulumi:"baseGroupDN"`
	BaseUserDN      *string `pulumi:"baseUserDN"`
	Domain          *string `pulumi:"domain"`
	Name            *string `pulumi:"name"`
	Password        *string `pulumi:"password"`
	PrimaryServer   *string `pulumi:"primaryServer"`
	SecondaryServer *string `pulumi:"secondaryServer"`
	Ssl             *string `pulumi:"ssl"`
	Username        *string `pulumi:"username"`
}





type IdentitySourceInput interface {
	pulumi.Input

	ToIdentitySourceOutput() IdentitySourceOutput
	ToIdentitySourceOutputWithContext(context.Context) IdentitySourceOutput
}

type IdentitySourceArgs struct {
	Alias           pulumi.StringPtrInput `pulumi:"alias"`
	BaseGroupDN     pulumi.StringPtrInput `pulumi:"baseGroupDN"`
	BaseUserDN      pulumi.StringPtrInput `pulumi:"baseUserDN"`
	Domain          pulumi.StringPtrInput `pulumi:"domain"`
	Name            pulumi.StringPtrInput `pulumi:"name"`
	Password        pulumi.StringPtrInput `pulumi:"password"`
	PrimaryServer   pulumi.StringPtrInput `pulumi:"primaryServer"`
	SecondaryServer pulumi.StringPtrInput `pulumi:"secondaryServer"`
	Ssl             pulumi.StringPtrInput `pulumi:"ssl"`
	Username        pulumi.StringPtrInput `pulumi:"username"`
}

func (IdentitySourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentitySource)(nil)).Elem()
}

func (i IdentitySourceArgs) ToIdentitySourceOutput() IdentitySourceOutput {
	return i.ToIdentitySourceOutputWithContext(context.Background())
}

func (i IdentitySourceArgs) ToIdentitySourceOutputWithContext(ctx context.Context) IdentitySourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentitySourceOutput)
}





type IdentitySourceArrayInput interface {
	pulumi.Input

	ToIdentitySourceArrayOutput() IdentitySourceArrayOutput
	ToIdentitySourceArrayOutputWithContext(context.Context) IdentitySourceArrayOutput
}

type IdentitySourceArray []IdentitySourceInput

func (IdentitySourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentitySource)(nil)).Elem()
}

func (i IdentitySourceArray) ToIdentitySourceArrayOutput() IdentitySourceArrayOutput {
	return i.ToIdentitySourceArrayOutputWithContext(context.Background())
}

func (i IdentitySourceArray) ToIdentitySourceArrayOutputWithContext(ctx context.Context) IdentitySourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentitySourceArrayOutput)
}

type IdentitySourceOutput struct{ *pulumi.OutputState }

func (IdentitySourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentitySource)(nil)).Elem()
}

func (o IdentitySourceOutput) ToIdentitySourceOutput() IdentitySourceOutput {
	return o
}

func (o IdentitySourceOutput) ToIdentitySourceOutputWithContext(ctx context.Context) IdentitySourceOutput {
	return o
}

func (o IdentitySourceOutput) Alias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.Alias }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) BaseGroupDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.BaseGroupDN }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) BaseUserDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.BaseUserDN }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) PrimaryServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.PrimaryServer }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) SecondaryServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.SecondaryServer }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) Ssl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.Ssl }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySource) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type IdentitySourceArrayOutput struct{ *pulumi.OutputState }

func (IdentitySourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentitySource)(nil)).Elem()
}

func (o IdentitySourceArrayOutput) ToIdentitySourceArrayOutput() IdentitySourceArrayOutput {
	return o
}

func (o IdentitySourceArrayOutput) ToIdentitySourceArrayOutputWithContext(ctx context.Context) IdentitySourceArrayOutput {
	return o
}

func (o IdentitySourceArrayOutput) Index(i pulumi.IntInput) IdentitySourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IdentitySource {
		return vs[0].([]IdentitySource)[vs[1].(int)]
	}).(IdentitySourceOutput)
}

type IdentitySourceResponse struct {
	Alias           *string `pulumi:"alias"`
	BaseGroupDN     *string `pulumi:"baseGroupDN"`
	BaseUserDN      *string `pulumi:"baseUserDN"`
	Domain          *string `pulumi:"domain"`
	Name            *string `pulumi:"name"`
	Password        *string `pulumi:"password"`
	PrimaryServer   *string `pulumi:"primaryServer"`
	SecondaryServer *string `pulumi:"secondaryServer"`
	Ssl             *string `pulumi:"ssl"`
	Username        *string `pulumi:"username"`
}

type IdentitySourceResponseOutput struct{ *pulumi.OutputState }

func (IdentitySourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentitySourceResponse)(nil)).Elem()
}

func (o IdentitySourceResponseOutput) ToIdentitySourceResponseOutput() IdentitySourceResponseOutput {
	return o
}

func (o IdentitySourceResponseOutput) ToIdentitySourceResponseOutputWithContext(ctx context.Context) IdentitySourceResponseOutput {
	return o
}

func (o IdentitySourceResponseOutput) Alias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.Alias }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) BaseGroupDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.BaseGroupDN }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) BaseUserDN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.BaseUserDN }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) PrimaryServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.PrimaryServer }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) SecondaryServer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.SecondaryServer }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) Ssl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.Ssl }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentitySourceResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type IdentitySourceResponseArrayOutput struct{ *pulumi.OutputState }

func (IdentitySourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentitySourceResponse)(nil)).Elem()
}

func (o IdentitySourceResponseArrayOutput) ToIdentitySourceResponseArrayOutput() IdentitySourceResponseArrayOutput {
	return o
}

func (o IdentitySourceResponseArrayOutput) ToIdentitySourceResponseArrayOutputWithContext(ctx context.Context) IdentitySourceResponseArrayOutput {
	return o
}

func (o IdentitySourceResponseArrayOutput) Index(i pulumi.IntInput) IdentitySourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IdentitySourceResponse {
		return vs[0].([]IdentitySourceResponse)[vs[1].(int)]
	}).(IdentitySourceResponseOutput)
}

type ManagementCluster struct {
	ClusterSize int      `pulumi:"clusterSize"`
	Hosts       []string `pulumi:"hosts"`
}





type ManagementClusterInput interface {
	pulumi.Input

	ToManagementClusterOutput() ManagementClusterOutput
	ToManagementClusterOutputWithContext(context.Context) ManagementClusterOutput
}

type ManagementClusterArgs struct {
	ClusterSize pulumi.IntInput         `pulumi:"clusterSize"`
	Hosts       pulumi.StringArrayInput `pulumi:"hosts"`
}

func (ManagementClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementCluster)(nil)).Elem()
}

func (i ManagementClusterArgs) ToManagementClusterOutput() ManagementClusterOutput {
	return i.ToManagementClusterOutputWithContext(context.Background())
}

func (i ManagementClusterArgs) ToManagementClusterOutputWithContext(ctx context.Context) ManagementClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementClusterOutput)
}

type ManagementClusterOutput struct{ *pulumi.OutputState }

func (ManagementClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementCluster)(nil)).Elem()
}

func (o ManagementClusterOutput) ToManagementClusterOutput() ManagementClusterOutput {
	return o
}

func (o ManagementClusterOutput) ToManagementClusterOutputWithContext(ctx context.Context) ManagementClusterOutput {
	return o
}

func (o ManagementClusterOutput) ClusterSize() pulumi.IntOutput {
	return o.ApplyT(func(v ManagementCluster) int { return v.ClusterSize }).(pulumi.IntOutput)
}

func (o ManagementClusterOutput) Hosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagementCluster) []string { return v.Hosts }).(pulumi.StringArrayOutput)
}

type ManagementClusterResponse struct {
	ClusterId         int      `pulumi:"clusterId"`
	ClusterSize       int      `pulumi:"clusterSize"`
	Hosts             []string `pulumi:"hosts"`
	ProvisioningState string   `pulumi:"provisioningState"`
}

type ManagementClusterResponseOutput struct{ *pulumi.OutputState }

func (ManagementClusterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementClusterResponse)(nil)).Elem()
}

func (o ManagementClusterResponseOutput) ToManagementClusterResponseOutput() ManagementClusterResponseOutput {
	return o
}

func (o ManagementClusterResponseOutput) ToManagementClusterResponseOutputWithContext(ctx context.Context) ManagementClusterResponseOutput {
	return o
}

func (o ManagementClusterResponseOutput) ClusterId() pulumi.IntOutput {
	return o.ApplyT(func(v ManagementClusterResponse) int { return v.ClusterId }).(pulumi.IntOutput)
}

func (o ManagementClusterResponseOutput) ClusterSize() pulumi.IntOutput {
	return o.ApplyT(func(v ManagementClusterResponse) int { return v.ClusterSize }).(pulumi.IntOutput)
}

func (o ManagementClusterResponseOutput) Hosts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagementClusterResponse) []string { return v.Hosts }).(pulumi.StringArrayOutput)
}

func (o ManagementClusterResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementClusterResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type NetAppVolume struct {
	Id string `pulumi:"id"`
}





type NetAppVolumeInput interface {
	pulumi.Input

	ToNetAppVolumeOutput() NetAppVolumeOutput
	ToNetAppVolumeOutputWithContext(context.Context) NetAppVolumeOutput
}

type NetAppVolumeArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (NetAppVolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetAppVolume)(nil)).Elem()
}

func (i NetAppVolumeArgs) ToNetAppVolumeOutput() NetAppVolumeOutput {
	return i.ToNetAppVolumeOutputWithContext(context.Background())
}

func (i NetAppVolumeArgs) ToNetAppVolumeOutputWithContext(ctx context.Context) NetAppVolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetAppVolumeOutput)
}

func (i NetAppVolumeArgs) ToNetAppVolumePtrOutput() NetAppVolumePtrOutput {
	return i.ToNetAppVolumePtrOutputWithContext(context.Background())
}

func (i NetAppVolumeArgs) ToNetAppVolumePtrOutputWithContext(ctx context.Context) NetAppVolumePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetAppVolumeOutput).ToNetAppVolumePtrOutputWithContext(ctx)
}









type NetAppVolumePtrInput interface {
	pulumi.Input

	ToNetAppVolumePtrOutput() NetAppVolumePtrOutput
	ToNetAppVolumePtrOutputWithContext(context.Context) NetAppVolumePtrOutput
}

type netAppVolumePtrType NetAppVolumeArgs

func NetAppVolumePtr(v *NetAppVolumeArgs) NetAppVolumePtrInput {
	return (*netAppVolumePtrType)(v)
}

func (*netAppVolumePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetAppVolume)(nil)).Elem()
}

func (i *netAppVolumePtrType) ToNetAppVolumePtrOutput() NetAppVolumePtrOutput {
	return i.ToNetAppVolumePtrOutputWithContext(context.Background())
}

func (i *netAppVolumePtrType) ToNetAppVolumePtrOutputWithContext(ctx context.Context) NetAppVolumePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetAppVolumePtrOutput)
}

type NetAppVolumeOutput struct{ *pulumi.OutputState }

func (NetAppVolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetAppVolume)(nil)).Elem()
}

func (o NetAppVolumeOutput) ToNetAppVolumeOutput() NetAppVolumeOutput {
	return o
}

func (o NetAppVolumeOutput) ToNetAppVolumeOutputWithContext(ctx context.Context) NetAppVolumeOutput {
	return o
}

func (o NetAppVolumeOutput) ToNetAppVolumePtrOutput() NetAppVolumePtrOutput {
	return o.ToNetAppVolumePtrOutputWithContext(context.Background())
}

func (o NetAppVolumeOutput) ToNetAppVolumePtrOutputWithContext(ctx context.Context) NetAppVolumePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetAppVolume) *NetAppVolume {
		return &v
	}).(NetAppVolumePtrOutput)
}

func (o NetAppVolumeOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v NetAppVolume) string { return v.Id }).(pulumi.StringOutput)
}

type NetAppVolumePtrOutput struct{ *pulumi.OutputState }

func (NetAppVolumePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetAppVolume)(nil)).Elem()
}

func (o NetAppVolumePtrOutput) ToNetAppVolumePtrOutput() NetAppVolumePtrOutput {
	return o
}

func (o NetAppVolumePtrOutput) ToNetAppVolumePtrOutputWithContext(ctx context.Context) NetAppVolumePtrOutput {
	return o
}

func (o NetAppVolumePtrOutput) Elem() NetAppVolumeOutput {
	return o.ApplyT(func(v *NetAppVolume) NetAppVolume {
		if v != nil {
			return *v
		}
		var ret NetAppVolume
		return ret
	}).(NetAppVolumeOutput)
}

func (o NetAppVolumePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetAppVolume) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type NetAppVolumeResponse struct {
	Id string `pulumi:"id"`
}

type NetAppVolumeResponseOutput struct{ *pulumi.OutputState }

func (NetAppVolumeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetAppVolumeResponse)(nil)).Elem()
}

func (o NetAppVolumeResponseOutput) ToNetAppVolumeResponseOutput() NetAppVolumeResponseOutput {
	return o
}

func (o NetAppVolumeResponseOutput) ToNetAppVolumeResponseOutputWithContext(ctx context.Context) NetAppVolumeResponseOutput {
	return o
}

func (o NetAppVolumeResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v NetAppVolumeResponse) string { return v.Id }).(pulumi.StringOutput)
}

type NetAppVolumeResponsePtrOutput struct{ *pulumi.OutputState }

func (NetAppVolumeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetAppVolumeResponse)(nil)).Elem()
}

func (o NetAppVolumeResponsePtrOutput) ToNetAppVolumeResponsePtrOutput() NetAppVolumeResponsePtrOutput {
	return o
}

func (o NetAppVolumeResponsePtrOutput) ToNetAppVolumeResponsePtrOutputWithContext(ctx context.Context) NetAppVolumeResponsePtrOutput {
	return o
}

func (o NetAppVolumeResponsePtrOutput) Elem() NetAppVolumeResponseOutput {
	return o.ApplyT(func(v *NetAppVolumeResponse) NetAppVolumeResponse {
		if v != nil {
			return *v
		}
		var ret NetAppVolumeResponse
		return ret
	}).(NetAppVolumeResponseOutput)
}

func (o NetAppVolumeResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetAppVolumeResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PSCredentialExecutionParameter struct {
	Name     string  `pulumi:"name"`
	Password *string `pulumi:"password"`
	Type     string  `pulumi:"type"`
	Username *string `pulumi:"username"`
}

type PSCredentialExecutionParameterResponse struct {
	Name     string  `pulumi:"name"`
	Password *string `pulumi:"password"`
	Type     string  `pulumi:"type"`
	Username *string `pulumi:"username"`
}

type PrivateCloudIdentity struct {
	Type *string `pulumi:"type"`
}





type PrivateCloudIdentityInput interface {
	pulumi.Input

	ToPrivateCloudIdentityOutput() PrivateCloudIdentityOutput
	ToPrivateCloudIdentityOutputWithContext(context.Context) PrivateCloudIdentityOutput
}

type PrivateCloudIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (PrivateCloudIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateCloudIdentity)(nil)).Elem()
}

func (i PrivateCloudIdentityArgs) ToPrivateCloudIdentityOutput() PrivateCloudIdentityOutput {
	return i.ToPrivateCloudIdentityOutputWithContext(context.Background())
}

func (i PrivateCloudIdentityArgs) ToPrivateCloudIdentityOutputWithContext(ctx context.Context) PrivateCloudIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateCloudIdentityOutput)
}

func (i PrivateCloudIdentityArgs) ToPrivateCloudIdentityPtrOutput() PrivateCloudIdentityPtrOutput {
	return i.ToPrivateCloudIdentityPtrOutputWithContext(context.Background())
}

func (i PrivateCloudIdentityArgs) ToPrivateCloudIdentityPtrOutputWithContext(ctx context.Context) PrivateCloudIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateCloudIdentityOutput).ToPrivateCloudIdentityPtrOutputWithContext(ctx)
}









type PrivateCloudIdentityPtrInput interface {
	pulumi.Input

	ToPrivateCloudIdentityPtrOutput() PrivateCloudIdentityPtrOutput
	ToPrivateCloudIdentityPtrOutputWithContext(context.Context) PrivateCloudIdentityPtrOutput
}

type privateCloudIdentityPtrType PrivateCloudIdentityArgs

func PrivateCloudIdentityPtr(v *PrivateCloudIdentityArgs) PrivateCloudIdentityPtrInput {
	return (*privateCloudIdentityPtrType)(v)
}

func (*privateCloudIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateCloudIdentity)(nil)).Elem()
}

func (i *privateCloudIdentityPtrType) ToPrivateCloudIdentityPtrOutput() PrivateCloudIdentityPtrOutput {
	return i.ToPrivateCloudIdentityPtrOutputWithContext(context.Background())
}

func (i *privateCloudIdentityPtrType) ToPrivateCloudIdentityPtrOutputWithContext(ctx context.Context) PrivateCloudIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateCloudIdentityPtrOutput)
}

type PrivateCloudIdentityOutput struct{ *pulumi.OutputState }

func (PrivateCloudIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateCloudIdentity)(nil)).Elem()
}

func (o PrivateCloudIdentityOutput) ToPrivateCloudIdentityOutput() PrivateCloudIdentityOutput {
	return o
}

func (o PrivateCloudIdentityOutput) ToPrivateCloudIdentityOutputWithContext(ctx context.Context) PrivateCloudIdentityOutput {
	return o
}

func (o PrivateCloudIdentityOutput) ToPrivateCloudIdentityPtrOutput() PrivateCloudIdentityPtrOutput {
	return o.ToPrivateCloudIdentityPtrOutputWithContext(context.Background())
}

func (o PrivateCloudIdentityOutput) ToPrivateCloudIdentityPtrOutputWithContext(ctx context.Context) PrivateCloudIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateCloudIdentity) *PrivateCloudIdentity {
		return &v
	}).(PrivateCloudIdentityPtrOutput)
}

func (o PrivateCloudIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateCloudIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PrivateCloudIdentityPtrOutput struct{ *pulumi.OutputState }

func (PrivateCloudIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateCloudIdentity)(nil)).Elem()
}

func (o PrivateCloudIdentityPtrOutput) ToPrivateCloudIdentityPtrOutput() PrivateCloudIdentityPtrOutput {
	return o
}

func (o PrivateCloudIdentityPtrOutput) ToPrivateCloudIdentityPtrOutputWithContext(ctx context.Context) PrivateCloudIdentityPtrOutput {
	return o
}

func (o PrivateCloudIdentityPtrOutput) Elem() PrivateCloudIdentityOutput {
	return o.ApplyT(func(v *PrivateCloudIdentity) PrivateCloudIdentity {
		if v != nil {
			return *v
		}
		var ret PrivateCloudIdentity
		return ret
	}).(PrivateCloudIdentityOutput)
}

func (o PrivateCloudIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateCloudIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type PrivateCloudIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type PrivateCloudIdentityResponseOutput struct{ *pulumi.OutputState }

func (PrivateCloudIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateCloudIdentityResponse)(nil)).Elem()
}

func (o PrivateCloudIdentityResponseOutput) ToPrivateCloudIdentityResponseOutput() PrivateCloudIdentityResponseOutput {
	return o
}

func (o PrivateCloudIdentityResponseOutput) ToPrivateCloudIdentityResponseOutputWithContext(ctx context.Context) PrivateCloudIdentityResponseOutput {
	return o
}

func (o PrivateCloudIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateCloudIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o PrivateCloudIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateCloudIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o PrivateCloudIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateCloudIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PrivateCloudIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateCloudIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateCloudIdentityResponse)(nil)).Elem()
}

func (o PrivateCloudIdentityResponsePtrOutput) ToPrivateCloudIdentityResponsePtrOutput() PrivateCloudIdentityResponsePtrOutput {
	return o
}

func (o PrivateCloudIdentityResponsePtrOutput) ToPrivateCloudIdentityResponsePtrOutputWithContext(ctx context.Context) PrivateCloudIdentityResponsePtrOutput {
	return o
}

func (o PrivateCloudIdentityResponsePtrOutput) Elem() PrivateCloudIdentityResponseOutput {
	return o.ApplyT(func(v *PrivateCloudIdentityResponse) PrivateCloudIdentityResponse {
		if v != nil {
			return *v
		}
		var ret PrivateCloudIdentityResponse
		return ret
	}).(PrivateCloudIdentityResponseOutput)
}

func (o PrivateCloudIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateCloudIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o PrivateCloudIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateCloudIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o PrivateCloudIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateCloudIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ScriptSecureStringExecutionParameter struct {
	Name        string  `pulumi:"name"`
	SecureValue *string `pulumi:"secureValue"`
	Type        string  `pulumi:"type"`
}

type ScriptSecureStringExecutionParameterResponse struct {
	Name        string  `pulumi:"name"`
	SecureValue *string `pulumi:"secureValue"`
	Type        string  `pulumi:"type"`
}

type ScriptStringExecutionParameter struct {
	Name  string  `pulumi:"name"`
	Type  string  `pulumi:"type"`
	Value *string `pulumi:"value"`
}

type ScriptStringExecutionParameterResponse struct {
	Name  string  `pulumi:"name"`
	Type  string  `pulumi:"type"`
	Value *string `pulumi:"value"`
}

type Sku struct {
	Name string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
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

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
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

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type VmHostPlacementPolicyProperties struct {
	AffinityStrength       *string  `pulumi:"affinityStrength"`
	AffinityType           string   `pulumi:"affinityType"`
	AzureHybridBenefitType *string  `pulumi:"azureHybridBenefitType"`
	DisplayName            *string  `pulumi:"displayName"`
	HostMembers            []string `pulumi:"hostMembers"`
	State                  *string  `pulumi:"state"`
	Type                   string   `pulumi:"type"`
	VmMembers              []string `pulumi:"vmMembers"`
}

type VmHostPlacementPolicyPropertiesResponse struct {
	AffinityStrength       *string  `pulumi:"affinityStrength"`
	AffinityType           string   `pulumi:"affinityType"`
	AzureHybridBenefitType *string  `pulumi:"azureHybridBenefitType"`
	DisplayName            *string  `pulumi:"displayName"`
	HostMembers            []string `pulumi:"hostMembers"`
	ProvisioningState      string   `pulumi:"provisioningState"`
	State                  *string  `pulumi:"state"`
	Type                   string   `pulumi:"type"`
	VmMembers              []string `pulumi:"vmMembers"`
}

type VmVmPlacementPolicyProperties struct {
	AffinityType string   `pulumi:"affinityType"`
	DisplayName  *string  `pulumi:"displayName"`
	State        *string  `pulumi:"state"`
	Type         string   `pulumi:"type"`
	VmMembers    []string `pulumi:"vmMembers"`
}

type VmVmPlacementPolicyPropertiesResponse struct {
	AffinityType      string   `pulumi:"affinityType"`
	DisplayName       *string  `pulumi:"displayName"`
	ProvisioningState string   `pulumi:"provisioningState"`
	State             *string  `pulumi:"state"`
	Type              string   `pulumi:"type"`
	VmMembers         []string `pulumi:"vmMembers"`
}

type WorkloadNetworkDhcpRelay struct {
	DhcpType        string   `pulumi:"dhcpType"`
	DisplayName     *string  `pulumi:"displayName"`
	Revision        *float64 `pulumi:"revision"`
	ServerAddresses []string `pulumi:"serverAddresses"`
}

type WorkloadNetworkDhcpRelayResponse struct {
	DhcpType          string   `pulumi:"dhcpType"`
	DisplayName       *string  `pulumi:"displayName"`
	ProvisioningState string   `pulumi:"provisioningState"`
	Revision          *float64 `pulumi:"revision"`
	Segments          []string `pulumi:"segments"`
	ServerAddresses   []string `pulumi:"serverAddresses"`
}

type WorkloadNetworkDhcpServer struct {
	DhcpType      string   `pulumi:"dhcpType"`
	DisplayName   *string  `pulumi:"displayName"`
	LeaseTime     *float64 `pulumi:"leaseTime"`
	Revision      *float64 `pulumi:"revision"`
	ServerAddress *string  `pulumi:"serverAddress"`
}

type WorkloadNetworkDhcpServerResponse struct {
	DhcpType          string   `pulumi:"dhcpType"`
	DisplayName       *string  `pulumi:"displayName"`
	LeaseTime         *float64 `pulumi:"leaseTime"`
	ProvisioningState string   `pulumi:"provisioningState"`
	Revision          *float64 `pulumi:"revision"`
	Segments          []string `pulumi:"segments"`
	ServerAddress     *string  `pulumi:"serverAddress"`
}

type WorkloadNetworkSegmentPortVifResponse struct {
	PortName *string `pulumi:"portName"`
}

type WorkloadNetworkSegmentPortVifResponseOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkSegmentPortVifResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkSegmentPortVifResponse)(nil)).Elem()
}

func (o WorkloadNetworkSegmentPortVifResponseOutput) ToWorkloadNetworkSegmentPortVifResponseOutput() WorkloadNetworkSegmentPortVifResponseOutput {
	return o
}

func (o WorkloadNetworkSegmentPortVifResponseOutput) ToWorkloadNetworkSegmentPortVifResponseOutputWithContext(ctx context.Context) WorkloadNetworkSegmentPortVifResponseOutput {
	return o
}

func (o WorkloadNetworkSegmentPortVifResponseOutput) PortName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadNetworkSegmentPortVifResponse) *string { return v.PortName }).(pulumi.StringPtrOutput)
}

type WorkloadNetworkSegmentPortVifResponseArrayOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkSegmentPortVifResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkloadNetworkSegmentPortVifResponse)(nil)).Elem()
}

func (o WorkloadNetworkSegmentPortVifResponseArrayOutput) ToWorkloadNetworkSegmentPortVifResponseArrayOutput() WorkloadNetworkSegmentPortVifResponseArrayOutput {
	return o
}

func (o WorkloadNetworkSegmentPortVifResponseArrayOutput) ToWorkloadNetworkSegmentPortVifResponseArrayOutputWithContext(ctx context.Context) WorkloadNetworkSegmentPortVifResponseArrayOutput {
	return o
}

func (o WorkloadNetworkSegmentPortVifResponseArrayOutput) Index(i pulumi.IntInput) WorkloadNetworkSegmentPortVifResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkloadNetworkSegmentPortVifResponse {
		return vs[0].([]WorkloadNetworkSegmentPortVifResponse)[vs[1].(int)]
	}).(WorkloadNetworkSegmentPortVifResponseOutput)
}

type WorkloadNetworkSegmentSubnet struct {
	DhcpRanges     []string `pulumi:"dhcpRanges"`
	GatewayAddress *string  `pulumi:"gatewayAddress"`
}





type WorkloadNetworkSegmentSubnetInput interface {
	pulumi.Input

	ToWorkloadNetworkSegmentSubnetOutput() WorkloadNetworkSegmentSubnetOutput
	ToWorkloadNetworkSegmentSubnetOutputWithContext(context.Context) WorkloadNetworkSegmentSubnetOutput
}

type WorkloadNetworkSegmentSubnetArgs struct {
	DhcpRanges     pulumi.StringArrayInput `pulumi:"dhcpRanges"`
	GatewayAddress pulumi.StringPtrInput   `pulumi:"gatewayAddress"`
}

func (WorkloadNetworkSegmentSubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkSegmentSubnet)(nil)).Elem()
}

func (i WorkloadNetworkSegmentSubnetArgs) ToWorkloadNetworkSegmentSubnetOutput() WorkloadNetworkSegmentSubnetOutput {
	return i.ToWorkloadNetworkSegmentSubnetOutputWithContext(context.Background())
}

func (i WorkloadNetworkSegmentSubnetArgs) ToWorkloadNetworkSegmentSubnetOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkSegmentSubnetOutput)
}

func (i WorkloadNetworkSegmentSubnetArgs) ToWorkloadNetworkSegmentSubnetPtrOutput() WorkloadNetworkSegmentSubnetPtrOutput {
	return i.ToWorkloadNetworkSegmentSubnetPtrOutputWithContext(context.Background())
}

func (i WorkloadNetworkSegmentSubnetArgs) ToWorkloadNetworkSegmentSubnetPtrOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkSegmentSubnetOutput).ToWorkloadNetworkSegmentSubnetPtrOutputWithContext(ctx)
}









type WorkloadNetworkSegmentSubnetPtrInput interface {
	pulumi.Input

	ToWorkloadNetworkSegmentSubnetPtrOutput() WorkloadNetworkSegmentSubnetPtrOutput
	ToWorkloadNetworkSegmentSubnetPtrOutputWithContext(context.Context) WorkloadNetworkSegmentSubnetPtrOutput
}

type workloadNetworkSegmentSubnetPtrType WorkloadNetworkSegmentSubnetArgs

func WorkloadNetworkSegmentSubnetPtr(v *WorkloadNetworkSegmentSubnetArgs) WorkloadNetworkSegmentSubnetPtrInput {
	return (*workloadNetworkSegmentSubnetPtrType)(v)
}

func (*workloadNetworkSegmentSubnetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkSegmentSubnet)(nil)).Elem()
}

func (i *workloadNetworkSegmentSubnetPtrType) ToWorkloadNetworkSegmentSubnetPtrOutput() WorkloadNetworkSegmentSubnetPtrOutput {
	return i.ToWorkloadNetworkSegmentSubnetPtrOutputWithContext(context.Background())
}

func (i *workloadNetworkSegmentSubnetPtrType) ToWorkloadNetworkSegmentSubnetPtrOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkSegmentSubnetPtrOutput)
}

type WorkloadNetworkSegmentSubnetOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkSegmentSubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkSegmentSubnet)(nil)).Elem()
}

func (o WorkloadNetworkSegmentSubnetOutput) ToWorkloadNetworkSegmentSubnetOutput() WorkloadNetworkSegmentSubnetOutput {
	return o
}

func (o WorkloadNetworkSegmentSubnetOutput) ToWorkloadNetworkSegmentSubnetOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetOutput {
	return o
}

func (o WorkloadNetworkSegmentSubnetOutput) ToWorkloadNetworkSegmentSubnetPtrOutput() WorkloadNetworkSegmentSubnetPtrOutput {
	return o.ToWorkloadNetworkSegmentSubnetPtrOutputWithContext(context.Background())
}

func (o WorkloadNetworkSegmentSubnetOutput) ToWorkloadNetworkSegmentSubnetPtrOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkloadNetworkSegmentSubnet) *WorkloadNetworkSegmentSubnet {
		return &v
	}).(WorkloadNetworkSegmentSubnetPtrOutput)
}

func (o WorkloadNetworkSegmentSubnetOutput) DhcpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkloadNetworkSegmentSubnet) []string { return v.DhcpRanges }).(pulumi.StringArrayOutput)
}

func (o WorkloadNetworkSegmentSubnetOutput) GatewayAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadNetworkSegmentSubnet) *string { return v.GatewayAddress }).(pulumi.StringPtrOutput)
}

type WorkloadNetworkSegmentSubnetPtrOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkSegmentSubnetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkSegmentSubnet)(nil)).Elem()
}

func (o WorkloadNetworkSegmentSubnetPtrOutput) ToWorkloadNetworkSegmentSubnetPtrOutput() WorkloadNetworkSegmentSubnetPtrOutput {
	return o
}

func (o WorkloadNetworkSegmentSubnetPtrOutput) ToWorkloadNetworkSegmentSubnetPtrOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetPtrOutput {
	return o
}

func (o WorkloadNetworkSegmentSubnetPtrOutput) Elem() WorkloadNetworkSegmentSubnetOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegmentSubnet) WorkloadNetworkSegmentSubnet {
		if v != nil {
			return *v
		}
		var ret WorkloadNetworkSegmentSubnet
		return ret
	}).(WorkloadNetworkSegmentSubnetOutput)
}

func (o WorkloadNetworkSegmentSubnetPtrOutput) DhcpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegmentSubnet) []string {
		if v == nil {
			return nil
		}
		return v.DhcpRanges
	}).(pulumi.StringArrayOutput)
}

func (o WorkloadNetworkSegmentSubnetPtrOutput) GatewayAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegmentSubnet) *string {
		if v == nil {
			return nil
		}
		return v.GatewayAddress
	}).(pulumi.StringPtrOutput)
}

type WorkloadNetworkSegmentSubnetResponse struct {
	DhcpRanges     []string `pulumi:"dhcpRanges"`
	GatewayAddress *string  `pulumi:"gatewayAddress"`
}

type WorkloadNetworkSegmentSubnetResponseOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkSegmentSubnetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadNetworkSegmentSubnetResponse)(nil)).Elem()
}

func (o WorkloadNetworkSegmentSubnetResponseOutput) ToWorkloadNetworkSegmentSubnetResponseOutput() WorkloadNetworkSegmentSubnetResponseOutput {
	return o
}

func (o WorkloadNetworkSegmentSubnetResponseOutput) ToWorkloadNetworkSegmentSubnetResponseOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetResponseOutput {
	return o
}

func (o WorkloadNetworkSegmentSubnetResponseOutput) DhcpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkloadNetworkSegmentSubnetResponse) []string { return v.DhcpRanges }).(pulumi.StringArrayOutput)
}

func (o WorkloadNetworkSegmentSubnetResponseOutput) GatewayAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkloadNetworkSegmentSubnetResponse) *string { return v.GatewayAddress }).(pulumi.StringPtrOutput)
}

type WorkloadNetworkSegmentSubnetResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkSegmentSubnetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkSegmentSubnetResponse)(nil)).Elem()
}

func (o WorkloadNetworkSegmentSubnetResponsePtrOutput) ToWorkloadNetworkSegmentSubnetResponsePtrOutput() WorkloadNetworkSegmentSubnetResponsePtrOutput {
	return o
}

func (o WorkloadNetworkSegmentSubnetResponsePtrOutput) ToWorkloadNetworkSegmentSubnetResponsePtrOutputWithContext(ctx context.Context) WorkloadNetworkSegmentSubnetResponsePtrOutput {
	return o
}

func (o WorkloadNetworkSegmentSubnetResponsePtrOutput) Elem() WorkloadNetworkSegmentSubnetResponseOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegmentSubnetResponse) WorkloadNetworkSegmentSubnetResponse {
		if v != nil {
			return *v
		}
		var ret WorkloadNetworkSegmentSubnetResponse
		return ret
	}).(WorkloadNetworkSegmentSubnetResponseOutput)
}

func (o WorkloadNetworkSegmentSubnetResponsePtrOutput) DhcpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegmentSubnetResponse) []string {
		if v == nil {
			return nil
		}
		return v.DhcpRanges
	}).(pulumi.StringArrayOutput)
}

func (o WorkloadNetworkSegmentSubnetResponsePtrOutput) GatewayAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegmentSubnetResponse) *string {
		if v == nil {
			return nil
		}
		return v.GatewayAddress
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AvailabilityPropertiesOutput{})
	pulumi.RegisterOutputType(AvailabilityPropertiesPtrOutput{})
	pulumi.RegisterOutputType(AvailabilityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AvailabilityPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CircuitResponseOutput{})
	pulumi.RegisterOutputType(CircuitResponsePtrOutput{})
	pulumi.RegisterOutputType(ClusterZoneResponseOutput{})
	pulumi.RegisterOutputType(ClusterZoneResponseArrayOutput{})
	pulumi.RegisterOutputType(DiskPoolVolumeOutput{})
	pulumi.RegisterOutputType(DiskPoolVolumePtrOutput{})
	pulumi.RegisterOutputType(DiskPoolVolumeResponseOutput{})
	pulumi.RegisterOutputType(DiskPoolVolumeResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionOutput{})
	pulumi.RegisterOutputType(EncryptionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionKeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(EncryptionKeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionKeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionKeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointsResponseOutput{})
	pulumi.RegisterOutputType(IdentitySourceOutput{})
	pulumi.RegisterOutputType(IdentitySourceArrayOutput{})
	pulumi.RegisterOutputType(IdentitySourceResponseOutput{})
	pulumi.RegisterOutputType(IdentitySourceResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagementClusterOutput{})
	pulumi.RegisterOutputType(ManagementClusterResponseOutput{})
	pulumi.RegisterOutputType(NetAppVolumeOutput{})
	pulumi.RegisterOutputType(NetAppVolumePtrOutput{})
	pulumi.RegisterOutputType(NetAppVolumeResponseOutput{})
	pulumi.RegisterOutputType(NetAppVolumeResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateCloudIdentityOutput{})
	pulumi.RegisterOutputType(PrivateCloudIdentityPtrOutput{})
	pulumi.RegisterOutputType(PrivateCloudIdentityResponseOutput{})
	pulumi.RegisterOutputType(PrivateCloudIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(WorkloadNetworkSegmentPortVifResponseOutput{})
	pulumi.RegisterOutputType(WorkloadNetworkSegmentPortVifResponseArrayOutput{})
	pulumi.RegisterOutputType(WorkloadNetworkSegmentSubnetOutput{})
	pulumi.RegisterOutputType(WorkloadNetworkSegmentSubnetPtrOutput{})
	pulumi.RegisterOutputType(WorkloadNetworkSegmentSubnetResponseOutput{})
	pulumi.RegisterOutputType(WorkloadNetworkSegmentSubnetResponsePtrOutput{})
}
