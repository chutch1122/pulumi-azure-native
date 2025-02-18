


package v20180907preview

type AzureSkuName string

const (
	AzureSkuNameKC8     = AzureSkuName("KC8")
	AzureSkuNameKC16    = AzureSkuName("KC16")
	AzureSkuNameKS8     = AzureSkuName("KS8")
	AzureSkuNameKS16    = AzureSkuName("KS16")
	AzureSkuName_D13_v2 = AzureSkuName("D13_v2")
	AzureSkuName_D14_v2 = AzureSkuName("D14_v2")
	AzureSkuNameL8      = AzureSkuName("L8")
	AzureSkuNameL16     = AzureSkuName("L16")
)

type AzureSkuTier string

const (
	AzureSkuTierStandard = AzureSkuTier("Standard")
)

type DataFormat string

const (
	DataFormatMULTIJSON = DataFormat("MULTIJSON")
	DataFormatJSON      = DataFormat("JSON")
	DataFormatCSV       = DataFormat("CSV")
)

func init() {
}
