


package v20211101preview

type StorageSkuName string

const (
	StorageSkuName_Standard_LRS    = StorageSkuName("Standard_LRS")
	StorageSkuName_Standard_GRS    = StorageSkuName("Standard_GRS")
	StorageSkuName_Standard_Ragrs  = StorageSkuName("Standard_Ragrs")
	StorageSkuName_Standard_ZRS    = StorageSkuName("Standard_ZRS")
	StorageSkuName_Premium_LRS     = StorageSkuName("Premium_LRS")
	StorageSkuName_Premium_ZRS     = StorageSkuName("Premium_ZRS")
	StorageSkuName_Standard_Gzrs   = StorageSkuName("Standard_Gzrs")
	StorageSkuName_Standard_Ragzrs = StorageSkuName("Standard_Ragzrs")
)

func init() {
}
