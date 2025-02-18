


package machinelearningcompute

type AgentVMSizeTypes string

const (
	AgentVMSizeTypes_Standard_A0     = AgentVMSizeTypes("Standard_A0")
	AgentVMSizeTypes_Standard_A1     = AgentVMSizeTypes("Standard_A1")
	AgentVMSizeTypes_Standard_A2     = AgentVMSizeTypes("Standard_A2")
	AgentVMSizeTypes_Standard_A3     = AgentVMSizeTypes("Standard_A3")
	AgentVMSizeTypes_Standard_A4     = AgentVMSizeTypes("Standard_A4")
	AgentVMSizeTypes_Standard_A5     = AgentVMSizeTypes("Standard_A5")
	AgentVMSizeTypes_Standard_A6     = AgentVMSizeTypes("Standard_A6")
	AgentVMSizeTypes_Standard_A7     = AgentVMSizeTypes("Standard_A7")
	AgentVMSizeTypes_Standard_A8     = AgentVMSizeTypes("Standard_A8")
	AgentVMSizeTypes_Standard_A9     = AgentVMSizeTypes("Standard_A9")
	AgentVMSizeTypes_Standard_A10    = AgentVMSizeTypes("Standard_A10")
	AgentVMSizeTypes_Standard_A11    = AgentVMSizeTypes("Standard_A11")
	AgentVMSizeTypes_Standard_D1     = AgentVMSizeTypes("Standard_D1")
	AgentVMSizeTypes_Standard_D2     = AgentVMSizeTypes("Standard_D2")
	AgentVMSizeTypes_Standard_D3     = AgentVMSizeTypes("Standard_D3")
	AgentVMSizeTypes_Standard_D4     = AgentVMSizeTypes("Standard_D4")
	AgentVMSizeTypes_Standard_D11    = AgentVMSizeTypes("Standard_D11")
	AgentVMSizeTypes_Standard_D12    = AgentVMSizeTypes("Standard_D12")
	AgentVMSizeTypes_Standard_D13    = AgentVMSizeTypes("Standard_D13")
	AgentVMSizeTypes_Standard_D14    = AgentVMSizeTypes("Standard_D14")
	AgentVMSizeTypes_Standard_D1_v2  = AgentVMSizeTypes("Standard_D1_v2")
	AgentVMSizeTypes_Standard_D2_v2  = AgentVMSizeTypes("Standard_D2_v2")
	AgentVMSizeTypes_Standard_D3_v2  = AgentVMSizeTypes("Standard_D3_v2")
	AgentVMSizeTypes_Standard_D4_v2  = AgentVMSizeTypes("Standard_D4_v2")
	AgentVMSizeTypes_Standard_D5_v2  = AgentVMSizeTypes("Standard_D5_v2")
	AgentVMSizeTypes_Standard_D11_v2 = AgentVMSizeTypes("Standard_D11_v2")
	AgentVMSizeTypes_Standard_D12_v2 = AgentVMSizeTypes("Standard_D12_v2")
	AgentVMSizeTypes_Standard_D13_v2 = AgentVMSizeTypes("Standard_D13_v2")
	AgentVMSizeTypes_Standard_D14_v2 = AgentVMSizeTypes("Standard_D14_v2")
	AgentVMSizeTypes_Standard_G1     = AgentVMSizeTypes("Standard_G1")
	AgentVMSizeTypes_Standard_G2     = AgentVMSizeTypes("Standard_G2")
	AgentVMSizeTypes_Standard_G3     = AgentVMSizeTypes("Standard_G3")
	AgentVMSizeTypes_Standard_G4     = AgentVMSizeTypes("Standard_G4")
	AgentVMSizeTypes_Standard_G5     = AgentVMSizeTypes("Standard_G5")
	AgentVMSizeTypes_Standard_DS1    = AgentVMSizeTypes("Standard_DS1")
	AgentVMSizeTypes_Standard_DS2    = AgentVMSizeTypes("Standard_DS2")
	AgentVMSizeTypes_Standard_DS3    = AgentVMSizeTypes("Standard_DS3")
	AgentVMSizeTypes_Standard_DS4    = AgentVMSizeTypes("Standard_DS4")
	AgentVMSizeTypes_Standard_DS11   = AgentVMSizeTypes("Standard_DS11")
	AgentVMSizeTypes_Standard_DS12   = AgentVMSizeTypes("Standard_DS12")
	AgentVMSizeTypes_Standard_DS13   = AgentVMSizeTypes("Standard_DS13")
	AgentVMSizeTypes_Standard_DS14   = AgentVMSizeTypes("Standard_DS14")
	AgentVMSizeTypes_Standard_GS1    = AgentVMSizeTypes("Standard_GS1")
	AgentVMSizeTypes_Standard_GS2    = AgentVMSizeTypes("Standard_GS2")
	AgentVMSizeTypes_Standard_GS3    = AgentVMSizeTypes("Standard_GS3")
	AgentVMSizeTypes_Standard_GS4    = AgentVMSizeTypes("Standard_GS4")
	AgentVMSizeTypes_Standard_GS5    = AgentVMSizeTypes("Standard_GS5")
)

type ClusterType string

const (
	ClusterTypeACS   = ClusterType("ACS")
	ClusterTypeLocal = ClusterType("Local")
)

type OrchestratorType string

const (
	OrchestratorTypeKubernetes = OrchestratorType("Kubernetes")
	OrchestratorTypeNone       = OrchestratorType("None")
)

type Status string

const (
	StatusEnabled  = Status("Enabled")
	StatusDisabled = Status("Disabled")
)

type SystemServiceType string

const (
	SystemServiceTypeNone            = SystemServiceType("None")
	SystemServiceTypeScoringFrontEnd = SystemServiceType("ScoringFrontEnd")
	SystemServiceTypeBatchFrontEnd   = SystemServiceType("BatchFrontEnd")
)

func init() {
}
