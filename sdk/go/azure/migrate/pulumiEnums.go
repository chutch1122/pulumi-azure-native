


package migrate

type AssessmentSizingCriterion string

const (
	AssessmentSizingCriterionPerformanceBased = AssessmentSizingCriterion("PerformanceBased")
	AssessmentSizingCriterionAsOnPremises     = AssessmentSizingCriterion("AsOnPremises")
)

type AssessmentStage string

const (
	AssessmentStageInProgress  = AssessmentStage("InProgress")
	AssessmentStageUnderReview = AssessmentStage("UnderReview")
	AssessmentStageApproved    = AssessmentStage("Approved")
)

type AzureDiskType string

const (
	AzureDiskTypeUnknown           = AzureDiskType("Unknown")
	AzureDiskTypeStandard          = AzureDiskType("Standard")
	AzureDiskTypePremium           = AzureDiskType("Premium")
	AzureDiskTypeStandardSSD       = AzureDiskType("StandardSSD")
	AzureDiskTypeStandardOrPremium = AzureDiskType("StandardOrPremium")
)

type AzureHybridUseBenefit string

const (
	AzureHybridUseBenefitUnknown = AzureHybridUseBenefit("Unknown")
	AzureHybridUseBenefitYes     = AzureHybridUseBenefit("Yes")
	AzureHybridUseBenefitNo      = AzureHybridUseBenefit("No")
)

type AzureLocation string

const (
	AzureLocationUnknown            = AzureLocation("Unknown")
	AzureLocationEastAsia           = AzureLocation("EastAsia")
	AzureLocationSoutheastAsia      = AzureLocation("SoutheastAsia")
	AzureLocationAustraliaEast      = AzureLocation("AustraliaEast")
	AzureLocationAustraliaSoutheast = AzureLocation("AustraliaSoutheast")
	AzureLocationBrazilSouth        = AzureLocation("BrazilSouth")
	AzureLocationCanadaCentral      = AzureLocation("CanadaCentral")
	AzureLocationCanadaEast         = AzureLocation("CanadaEast")
	AzureLocationWestEurope         = AzureLocation("WestEurope")
	AzureLocationNorthEurope        = AzureLocation("NorthEurope")
	AzureLocationCentralIndia       = AzureLocation("CentralIndia")
	AzureLocationSouthIndia         = AzureLocation("SouthIndia")
	AzureLocationWestIndia          = AzureLocation("WestIndia")
	AzureLocationJapanEast          = AzureLocation("JapanEast")
	AzureLocationJapanWest          = AzureLocation("JapanWest")
	AzureLocationKoreaCentral       = AzureLocation("KoreaCentral")
	AzureLocationKoreaSouth         = AzureLocation("KoreaSouth")
	AzureLocationUkWest             = AzureLocation("UkWest")
	AzureLocationUkSouth            = AzureLocation("UkSouth")
	AzureLocationNorthCentralUs     = AzureLocation("NorthCentralUs")
	AzureLocationEastUs             = AzureLocation("EastUs")
	AzureLocationWestUs2            = AzureLocation("WestUs2")
	AzureLocationSouthCentralUs     = AzureLocation("SouthCentralUs")
	AzureLocationCentralUs          = AzureLocation("CentralUs")
	AzureLocationEastUs2            = AzureLocation("EastUs2")
	AzureLocationWestUs             = AzureLocation("WestUs")
	AzureLocationWestCentralUs      = AzureLocation("WestCentralUs")
	AzureLocationGermanyCentral     = AzureLocation("GermanyCentral")
	AzureLocationGermanyNortheast   = AzureLocation("GermanyNortheast")
	AzureLocationChinaNorth         = AzureLocation("ChinaNorth")
	AzureLocationChinaEast          = AzureLocation("ChinaEast")
	AzureLocationUSGovArizona       = AzureLocation("USGovArizona")
	AzureLocationUSGovTexas         = AzureLocation("USGovTexas")
	AzureLocationUSGovIowa          = AzureLocation("USGovIowa")
	AzureLocationUSGovVirginia      = AzureLocation("USGovVirginia")
	AzureLocationUSDoDCentral       = AzureLocation("USDoDCentral")
	AzureLocationUSDoDEast          = AzureLocation("USDoDEast")
)

type AzureOfferCode string

const (
	AzureOfferCodeUnknown         = AzureOfferCode("Unknown")
	AzureOfferCodeMSAZR0003P      = AzureOfferCode("MSAZR0003P")
	AzureOfferCodeMSAZR0044P      = AzureOfferCode("MSAZR0044P")
	AzureOfferCodeMSAZR0059P      = AzureOfferCode("MSAZR0059P")
	AzureOfferCodeMSAZR0060P      = AzureOfferCode("MSAZR0060P")
	AzureOfferCodeMSAZR0062P      = AzureOfferCode("MSAZR0062P")
	AzureOfferCodeMSAZR0063P      = AzureOfferCode("MSAZR0063P")
	AzureOfferCodeMSAZR0064P      = AzureOfferCode("MSAZR0064P")
	AzureOfferCodeMSAZR0029P      = AzureOfferCode("MSAZR0029P")
	AzureOfferCodeMSAZR0022P      = AzureOfferCode("MSAZR0022P")
	AzureOfferCodeMSAZR0023P      = AzureOfferCode("MSAZR0023P")
	AzureOfferCodeMSAZR0148P      = AzureOfferCode("MSAZR0148P")
	AzureOfferCodeMSAZR0025P      = AzureOfferCode("MSAZR0025P")
	AzureOfferCodeMSAZR0036P      = AzureOfferCode("MSAZR0036P")
	AzureOfferCodeMSAZR0120P      = AzureOfferCode("MSAZR0120P")
	AzureOfferCodeMSAZR0121P      = AzureOfferCode("MSAZR0121P")
	AzureOfferCodeMSAZR0122P      = AzureOfferCode("MSAZR0122P")
	AzureOfferCodeMSAZR0123P      = AzureOfferCode("MSAZR0123P")
	AzureOfferCodeMSAZR0124P      = AzureOfferCode("MSAZR0124P")
	AzureOfferCodeMSAZR0125P      = AzureOfferCode("MSAZR0125P")
	AzureOfferCodeMSAZR0126P      = AzureOfferCode("MSAZR0126P")
	AzureOfferCodeMSAZR0127P      = AzureOfferCode("MSAZR0127P")
	AzureOfferCodeMSAZR0128P      = AzureOfferCode("MSAZR0128P")
	AzureOfferCodeMSAZR0129P      = AzureOfferCode("MSAZR0129P")
	AzureOfferCodeMSAZR0130P      = AzureOfferCode("MSAZR0130P")
	AzureOfferCodeMSAZR0111P      = AzureOfferCode("MSAZR0111P")
	AzureOfferCodeMSAZR0144P      = AzureOfferCode("MSAZR0144P")
	AzureOfferCodeMSAZR0149P      = AzureOfferCode("MSAZR0149P")
	AzureOfferCodeMSMCAZR0044P    = AzureOfferCode("MSMCAZR0044P")
	AzureOfferCodeMSMCAZR0059P    = AzureOfferCode("MSMCAZR0059P")
	AzureOfferCodeMSMCAZR0060P    = AzureOfferCode("MSMCAZR0060P")
	AzureOfferCodeMSMCAZR0063P    = AzureOfferCode("MSMCAZR0063P")
	AzureOfferCodeMSMCAZR0120P    = AzureOfferCode("MSMCAZR0120P")
	AzureOfferCodeMSMCAZR0121P    = AzureOfferCode("MSMCAZR0121P")
	AzureOfferCodeMSMCAZR0125P    = AzureOfferCode("MSMCAZR0125P")
	AzureOfferCodeMSMCAZR0128P    = AzureOfferCode("MSMCAZR0128P")
	AzureOfferCodeMSAZRDE0003P    = AzureOfferCode("MSAZRDE0003P")
	AzureOfferCodeMSAZRDE0044P    = AzureOfferCode("MSAZRDE0044P")
	AzureOfferCodeMSAZRUSGOV0003P = AzureOfferCode("MSAZRUSGOV0003P")
	AzureOfferCodeEA              = AzureOfferCode("EA")
)

type AzurePricingTier string

const (
	AzurePricingTierStandard = AzurePricingTier("Standard")
	AzurePricingTierBasic    = AzurePricingTier("Basic")
)

type AzureStorageRedundancy string

const (
	AzureStorageRedundancyUnknown                = AzureStorageRedundancy("Unknown")
	AzureStorageRedundancyLocallyRedundant       = AzureStorageRedundancy("LocallyRedundant")
	AzureStorageRedundancyZoneRedundant          = AzureStorageRedundancy("ZoneRedundant")
	AzureStorageRedundancyGeoRedundant           = AzureStorageRedundancy("GeoRedundant")
	AzureStorageRedundancyReadAccessGeoRedundant = AzureStorageRedundancy("ReadAccessGeoRedundant")
)

type AzureVmFamily string

const (
	AzureVmFamilyUnknown          = AzureVmFamily("Unknown")
	AzureVmFamily_Basic_A0_A4     = AzureVmFamily("Basic_A0_A4")
	AzureVmFamily_Standard_A0_A7  = AzureVmFamily("Standard_A0_A7")
	AzureVmFamily_Standard_A8_A11 = AzureVmFamily("Standard_A8_A11")
	AzureVmFamily_Av2_series      = AzureVmFamily("Av2_series")
	AzureVmFamily_D_series        = AzureVmFamily("D_series")
	AzureVmFamily_Dv2_series      = AzureVmFamily("Dv2_series")
	AzureVmFamily_DS_series       = AzureVmFamily("DS_series")
	AzureVmFamily_DSv2_series     = AzureVmFamily("DSv2_series")
	AzureVmFamily_F_series        = AzureVmFamily("F_series")
	AzureVmFamily_Fs_series       = AzureVmFamily("Fs_series")
	AzureVmFamily_G_series        = AzureVmFamily("G_series")
	AzureVmFamily_GS_series       = AzureVmFamily("GS_series")
	AzureVmFamily_H_series        = AzureVmFamily("H_series")
	AzureVmFamily_Ls_series       = AzureVmFamily("Ls_series")
	AzureVmFamily_Dsv3_series     = AzureVmFamily("Dsv3_series")
	AzureVmFamily_Dv3_series      = AzureVmFamily("Dv3_series")
	AzureVmFamily_Fsv2_series     = AzureVmFamily("Fsv2_series")
	AzureVmFamily_Ev3_series      = AzureVmFamily("Ev3_series")
	AzureVmFamily_Esv3_series     = AzureVmFamily("Esv3_series")
	AzureVmFamily_M_series        = AzureVmFamily("M_series")
	AzureVmFamily_DC_Series       = AzureVmFamily("DC_Series")
)

type Currency string

const (
	CurrencyUnknown = Currency("Unknown")
	CurrencyUSD     = Currency("USD")
	CurrencyDKK     = Currency("DKK")
	CurrencyCAD     = Currency("CAD")
	CurrencyIDR     = Currency("IDR")
	CurrencyJPY     = Currency("JPY")
	CurrencyKRW     = Currency("KRW")
	CurrencyNZD     = Currency("NZD")
	CurrencyNOK     = Currency("NOK")
	CurrencyRUB     = Currency("RUB")
	CurrencySAR     = Currency("SAR")
	CurrencyZAR     = Currency("ZAR")
	CurrencySEK     = Currency("SEK")
	CurrencyTRY     = Currency("TRY")
	CurrencyGBP     = Currency("GBP")
	CurrencyMXN     = Currency("MXN")
	CurrencyMYR     = Currency("MYR")
	CurrencyINR     = Currency("INR")
	CurrencyHKD     = Currency("HKD")
	CurrencyBRL     = Currency("BRL")
	CurrencyTWD     = Currency("TWD")
	CurrencyEUR     = Currency("EUR")
	CurrencyCHF     = Currency("CHF")
	CurrencyARS     = Currency("ARS")
	CurrencyAUD     = Currency("AUD")
	CurrencyCNY     = Currency("CNY")
)

type Percentile string

const (
	PercentilePercentile50 = Percentile("Percentile50")
	PercentilePercentile90 = Percentile("Percentile90")
	PercentilePercentile95 = Percentile("Percentile95")
	PercentilePercentile99 = Percentile("Percentile99")
)

type ProjectStatus string

const (
	ProjectStatusActive   = ProjectStatus("Active")
	ProjectStatusInactive = ProjectStatus("Inactive")
)

type ProvisioningState string

const (
	ProvisioningStateAccepted  = ProvisioningState("Accepted")
	ProvisioningStateCreating  = ProvisioningState("Creating")
	ProvisioningStateDeleting  = ProvisioningState("Deleting")
	ProvisioningStateFailed    = ProvisioningState("Failed")
	ProvisioningStateMoving    = ProvisioningState("Moving")
	ProvisioningStateSucceeded = ProvisioningState("Succeeded")
)

type ReservedInstance string

const (
	ReservedInstanceNone    = ReservedInstance("None")
	ReservedInstanceRI1Year = ReservedInstance("RI1Year")
	ReservedInstanceRI3Year = ReservedInstance("RI3Year")
)

type ResourceIdentityType string

const (
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned   = ResourceIdentityType("UserAssigned")
)

type TargetAvailabilityZone string

const (
	TargetAvailabilityZoneOne   = TargetAvailabilityZone("1")
	TargetAvailabilityZoneTwo   = TargetAvailabilityZone("2")
	TargetAvailabilityZoneThree = TargetAvailabilityZone("3")
	TargetAvailabilityZoneNA    = TargetAvailabilityZone("NA")
)

type TimeRange string

const (
	TimeRangeDay    = TimeRange("Day")
	TimeRangeWeek   = TimeRange("Week")
	TimeRangeMonth  = TimeRange("Month")
	TimeRangeCustom = TimeRange("Custom")
)

type ZoneRedundant string

const (
	ZoneRedundantEnable  = ZoneRedundant("Enable")
	ZoneRedundantDisable = ZoneRedundant("Disable")
)

func init() {
}
