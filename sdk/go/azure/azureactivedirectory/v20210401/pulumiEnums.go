


package v20210401

type B2CResourceSKUName string

const (
	// Azure AD B2C usage is billed to a linked Azure subscription and uses a monthly active users (MAU) billing model.
	B2CResourceSKUNameStandard = B2CResourceSKUName("Standard")
	// Azure AD B2C usage is billed to a linked Azure subscription and uses number of authentications based billing.
	B2CResourceSKUNamePremiumP1 = B2CResourceSKUName("PremiumP1")
	// Azure AD B2C usage is billed to a linked Azure subscription and uses number of authentications based billing.
	B2CResourceSKUNamePremiumP2 = B2CResourceSKUName("PremiumP2")
)

type B2CResourceSKUTier string

const (
	// The SKU tier used for all Azure AD B2C tenants.
	B2CResourceSKUTierA0 = B2CResourceSKUTier("A0")
)

func init() {
}
