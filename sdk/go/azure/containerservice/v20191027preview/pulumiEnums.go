


package v20191027preview

type OSType string

const (
	OSTypeLinux   = OSType("Linux")
	OSTypeWindows = OSType("Windows")
)

type OpenShiftAgentPoolProfileRole string

const (
	OpenShiftAgentPoolProfileRoleCompute = OpenShiftAgentPoolProfileRole("compute")
	OpenShiftAgentPoolProfileRoleInfra   = OpenShiftAgentPoolProfileRole("infra")
)

type OpenShiftContainerServiceVMSize string

const (
	OpenShiftContainerServiceVMSize_Standard_D2s_v3  = OpenShiftContainerServiceVMSize("Standard_D2s_v3")
	OpenShiftContainerServiceVMSize_Standard_D4s_v3  = OpenShiftContainerServiceVMSize("Standard_D4s_v3")
	OpenShiftContainerServiceVMSize_Standard_D8s_v3  = OpenShiftContainerServiceVMSize("Standard_D8s_v3")
	OpenShiftContainerServiceVMSize_Standard_D16s_v3 = OpenShiftContainerServiceVMSize("Standard_D16s_v3")
	OpenShiftContainerServiceVMSize_Standard_D32s_v3 = OpenShiftContainerServiceVMSize("Standard_D32s_v3")
	OpenShiftContainerServiceVMSize_Standard_D64s_v3 = OpenShiftContainerServiceVMSize("Standard_D64s_v3")
	OpenShiftContainerServiceVMSize_Standard_DS4_v2  = OpenShiftContainerServiceVMSize("Standard_DS4_v2")
	OpenShiftContainerServiceVMSize_Standard_DS5_v2  = OpenShiftContainerServiceVMSize("Standard_DS5_v2")
	OpenShiftContainerServiceVMSize_Standard_F8s_v2  = OpenShiftContainerServiceVMSize("Standard_F8s_v2")
	OpenShiftContainerServiceVMSize_Standard_F16s_v2 = OpenShiftContainerServiceVMSize("Standard_F16s_v2")
	OpenShiftContainerServiceVMSize_Standard_F32s_v2 = OpenShiftContainerServiceVMSize("Standard_F32s_v2")
	OpenShiftContainerServiceVMSize_Standard_F64s_v2 = OpenShiftContainerServiceVMSize("Standard_F64s_v2")
	OpenShiftContainerServiceVMSize_Standard_F72s_v2 = OpenShiftContainerServiceVMSize("Standard_F72s_v2")
	OpenShiftContainerServiceVMSize_Standard_F8s     = OpenShiftContainerServiceVMSize("Standard_F8s")
	OpenShiftContainerServiceVMSize_Standard_F16s    = OpenShiftContainerServiceVMSize("Standard_F16s")
	OpenShiftContainerServiceVMSize_Standard_E4s_v3  = OpenShiftContainerServiceVMSize("Standard_E4s_v3")
	OpenShiftContainerServiceVMSize_Standard_E8s_v3  = OpenShiftContainerServiceVMSize("Standard_E8s_v3")
	OpenShiftContainerServiceVMSize_Standard_E16s_v3 = OpenShiftContainerServiceVMSize("Standard_E16s_v3")
	OpenShiftContainerServiceVMSize_Standard_E20s_v3 = OpenShiftContainerServiceVMSize("Standard_E20s_v3")
	OpenShiftContainerServiceVMSize_Standard_E32s_v3 = OpenShiftContainerServiceVMSize("Standard_E32s_v3")
	OpenShiftContainerServiceVMSize_Standard_E64s_v3 = OpenShiftContainerServiceVMSize("Standard_E64s_v3")
	OpenShiftContainerServiceVMSize_Standard_GS2     = OpenShiftContainerServiceVMSize("Standard_GS2")
	OpenShiftContainerServiceVMSize_Standard_GS3     = OpenShiftContainerServiceVMSize("Standard_GS3")
	OpenShiftContainerServiceVMSize_Standard_GS4     = OpenShiftContainerServiceVMSize("Standard_GS4")
	OpenShiftContainerServiceVMSize_Standard_GS5     = OpenShiftContainerServiceVMSize("Standard_GS5")
	OpenShiftContainerServiceVMSize_Standard_DS12_v2 = OpenShiftContainerServiceVMSize("Standard_DS12_v2")
	OpenShiftContainerServiceVMSize_Standard_DS13_v2 = OpenShiftContainerServiceVMSize("Standard_DS13_v2")
	OpenShiftContainerServiceVMSize_Standard_DS14_v2 = OpenShiftContainerServiceVMSize("Standard_DS14_v2")
	OpenShiftContainerServiceVMSize_Standard_DS15_v2 = OpenShiftContainerServiceVMSize("Standard_DS15_v2")
	OpenShiftContainerServiceVMSize_Standard_L4s     = OpenShiftContainerServiceVMSize("Standard_L4s")
	OpenShiftContainerServiceVMSize_Standard_L8s     = OpenShiftContainerServiceVMSize("Standard_L8s")
	OpenShiftContainerServiceVMSize_Standard_L16s    = OpenShiftContainerServiceVMSize("Standard_L16s")
	OpenShiftContainerServiceVMSize_Standard_L32s    = OpenShiftContainerServiceVMSize("Standard_L32s")
)

func init() {
}
