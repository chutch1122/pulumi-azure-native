


package v20210901preview

type EncryptionAtHost string

const (
	EncryptionAtHostDisabled = EncryptionAtHost("Disabled")
	EncryptionAtHostEnabled  = EncryptionAtHost("Enabled")
)

type SoftwareDefinedNetwork string

const (
	SoftwareDefinedNetworkOVNKubernetes = SoftwareDefinedNetwork("OVNKubernetes")
	SoftwareDefinedNetworkOpenShiftSDN  = SoftwareDefinedNetwork("OpenShiftSDN")
)

type VMSize string

const (
	VMSize_Standard_D16as_v4 = VMSize("Standard_D16as_v4")
	VMSize_Standard_D16s_v3  = VMSize("Standard_D16s_v3")
	VMSize_Standard_D2s_v3   = VMSize("Standard_D2s_v3")
	VMSize_Standard_D32as_v4 = VMSize("Standard_D32as_v4")
	VMSize_Standard_D32s_v3  = VMSize("Standard_D32s_v3")
	VMSize_Standard_D4as_v4  = VMSize("Standard_D4as_v4")
	VMSize_Standard_D4s_v3   = VMSize("Standard_D4s_v3")
	VMSize_Standard_D8as_v4  = VMSize("Standard_D8as_v4")
	VMSize_Standard_D8s_v3   = VMSize("Standard_D8s_v3")
	VMSize_Standard_E16s_v3  = VMSize("Standard_E16s_v3")
	VMSize_Standard_E32s_v3  = VMSize("Standard_E32s_v3")
	VMSize_Standard_E4s_v3   = VMSize("Standard_E4s_v3")
	VMSize_Standard_E64i_v3  = VMSize("Standard_E64i_v3")
	VMSize_Standard_E64is_v3 = VMSize("Standard_E64is_v3")
	VMSize_Standard_E8s_v3   = VMSize("Standard_E8s_v3")
	VMSize_Standard_F16s_v2  = VMSize("Standard_F16s_v2")
	VMSize_Standard_F32s_v2  = VMSize("Standard_F32s_v2")
	VMSize_Standard_F4s_v2   = VMSize("Standard_F4s_v2")
	VMSize_Standard_F72s_v2  = VMSize("Standard_F72s_v2")
	VMSize_Standard_F8s_v2   = VMSize("Standard_F8s_v2")
	VMSize_Standard_G5       = VMSize("Standard_G5")
	VMSize_Standard_GS5      = VMSize("Standard_GS5")
	VMSize_Standard_M128ms   = VMSize("Standard_M128ms")
)

type Visibility string

const (
	VisibilityPrivate = Visibility("Private")
	VisibilityPublic  = Visibility("Public")
)

func init() {
}
