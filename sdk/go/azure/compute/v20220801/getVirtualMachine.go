


package v20220801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineResult, error) {
	var rv LookupVirtualMachineResult
	err := ctx.Invoke("azure-native:compute/v20220801:getVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineArgs struct {
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	VmName            string  `pulumi:"vmName"`
}


type LookupVirtualMachineResult struct {
	AdditionalCapabilities  *AdditionalCapabilitiesResponse     `pulumi:"additionalCapabilities"`
	ApplicationProfile      *ApplicationProfileResponse         `pulumi:"applicationProfile"`
	AvailabilitySet         *SubResourceResponse                `pulumi:"availabilitySet"`
	BillingProfile          *BillingProfileResponse             `pulumi:"billingProfile"`
	CapacityReservation     *CapacityReservationProfileResponse `pulumi:"capacityReservation"`
	DiagnosticsProfile      *DiagnosticsProfileResponse         `pulumi:"diagnosticsProfile"`
	EvictionPolicy          *string                             `pulumi:"evictionPolicy"`
	ExtendedLocation        *ExtendedLocationResponse           `pulumi:"extendedLocation"`
	ExtensionsTimeBudget    *string                             `pulumi:"extensionsTimeBudget"`
	HardwareProfile         *HardwareProfileResponse            `pulumi:"hardwareProfile"`
	Host                    *SubResourceResponse                `pulumi:"host"`
	HostGroup               *SubResourceResponse                `pulumi:"hostGroup"`
	Id                      string                              `pulumi:"id"`
	Identity                *VirtualMachineIdentityResponse     `pulumi:"identity"`
	InstanceView            VirtualMachineInstanceViewResponse  `pulumi:"instanceView"`
	LicenseType             *string                             `pulumi:"licenseType"`
	Location                string                              `pulumi:"location"`
	Name                    string                              `pulumi:"name"`
	NetworkProfile          *NetworkProfileResponse             `pulumi:"networkProfile"`
	OsProfile               *OSProfileResponse                  `pulumi:"osProfile"`
	Plan                    *PlanResponse                       `pulumi:"plan"`
	PlatformFaultDomain     *int                                `pulumi:"platformFaultDomain"`
	Priority                *string                             `pulumi:"priority"`
	ProvisioningState       string                              `pulumi:"provisioningState"`
	ProximityPlacementGroup *SubResourceResponse                `pulumi:"proximityPlacementGroup"`
	Resources               []VirtualMachineExtensionResponse   `pulumi:"resources"`
	ScheduledEventsProfile  *ScheduledEventsProfileResponse     `pulumi:"scheduledEventsProfile"`
	SecurityProfile         *SecurityProfileResponse            `pulumi:"securityProfile"`
	StorageProfile          *StorageProfileResponse             `pulumi:"storageProfile"`
	Tags                    map[string]string                   `pulumi:"tags"`
	TimeCreated             string                              `pulumi:"timeCreated"`
	Type                    string                              `pulumi:"type"`
	UserData                *string                             `pulumi:"userData"`
	VirtualMachineScaleSet  *SubResourceResponse                `pulumi:"virtualMachineScaleSet"`
	VmId                    string                              `pulumi:"vmId"`
	Zones                   []string                            `pulumi:"zones"`
}

func LookupVirtualMachineOutput(ctx *pulumi.Context, args LookupVirtualMachineOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineResult, error) {
			args := v.(LookupVirtualMachineArgs)
			r, err := LookupVirtualMachine(ctx, &args, opts...)
			var s LookupVirtualMachineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineResultOutput)
}

type LookupVirtualMachineOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	VmName            pulumi.StringInput    `pulumi:"vmName"`
}

func (LookupVirtualMachineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineArgs)(nil)).Elem()
}


type LookupVirtualMachineResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineResult)(nil)).Elem()
}

func (o LookupVirtualMachineResultOutput) ToLookupVirtualMachineResultOutput() LookupVirtualMachineResultOutput {
	return o
}

func (o LookupVirtualMachineResultOutput) ToLookupVirtualMachineResultOutputWithContext(ctx context.Context) LookupVirtualMachineResultOutput {
	return o
}

func (o LookupVirtualMachineResultOutput) AdditionalCapabilities() AdditionalCapabilitiesResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *AdditionalCapabilitiesResponse { return v.AdditionalCapabilities }).(AdditionalCapabilitiesResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) ApplicationProfile() ApplicationProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *ApplicationProfileResponse { return v.ApplicationProfile }).(ApplicationProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) AvailabilitySet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *SubResourceResponse { return v.AvailabilitySet }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) BillingProfile() BillingProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *BillingProfileResponse { return v.BillingProfile }).(BillingProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) CapacityReservation() CapacityReservationProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *CapacityReservationProfileResponse { return v.CapacityReservation }).(CapacityReservationProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) DiagnosticsProfile() DiagnosticsProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *DiagnosticsProfileResponse { return v.DiagnosticsProfile }).(DiagnosticsProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) EvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.EvictionPolicy }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) ExtensionsTimeBudget() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.ExtensionsTimeBudget }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *HardwareProfileResponse { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) Host() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *SubResourceResponse { return v.Host }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) HostGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *SubResourceResponse { return v.HostGroup }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) Identity() VirtualMachineIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *VirtualMachineIdentityResponse { return v.Identity }).(VirtualMachineIdentityResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) InstanceView() VirtualMachineInstanceViewResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) VirtualMachineInstanceViewResponse { return v.InstanceView }).(VirtualMachineInstanceViewResponseOutput)
}

func (o LookupVirtualMachineResultOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *NetworkProfileResponse { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) OsProfile() OSProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *OSProfileResponse { return v.OsProfile }).(OSProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *PlanResponse { return v.Plan }).(PlanResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) PlatformFaultDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *int { return v.PlatformFaultDomain }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualMachineResultOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.Priority }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) ProximityPlacementGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *SubResourceResponse { return v.ProximityPlacementGroup }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) Resources() VirtualMachineExtensionResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []VirtualMachineExtensionResponse { return v.Resources }).(VirtualMachineExtensionResponseArrayOutput)
}

func (o LookupVirtualMachineResultOutput) ScheduledEventsProfile() ScheduledEventsProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *ScheduledEventsProfileResponse { return v.ScheduledEventsProfile }).(ScheduledEventsProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) SecurityProfile() SecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *SecurityProfileResponse { return v.SecurityProfile }).(SecurityProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *StorageProfileResponse { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualMachineResultOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.TimeCreated }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.UserData }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) VirtualMachineScaleSet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *SubResourceResponse { return v.VirtualMachineScaleSet }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.VmId }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineResultOutput{})
}
