


package v20211130

type SkuName string

const (
	// The dedicated HSM is a Safenet Luna Network HSM A790 device.
	SkuName_SafeNet_Luna_Network_HSM_A790 = SkuName("SafeNet Luna Network HSM A790")
	// The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device with 1 local master key which supports up to 60 calls per second.
	SkuName_PayShield10K_LMK1_CPS60 = SkuName("payShield10K_LMK1_CPS60")
	// The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device with 1 local master key which supports up to 250 calls per second.
	SkuName_PayShield10K_LMK1_CPS250 = SkuName("payShield10K_LMK1_CPS250")
	// The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device with 1 local master key which supports up to 2500 calls per second.
	SkuName_PayShield10K_LMK1_CPS2500 = SkuName("payShield10K_LMK1_CPS2500")
	// The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device with 2 local master keys which supports up to 60 calls per second.
	SkuName_PayShield10K_LMK2_CPS60 = SkuName("payShield10K_LMK2_CPS60")
	// The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device with 2 local master keys which supports up to 250 calls per second.
	SkuName_PayShield10K_LMK2_CPS250 = SkuName("payShield10K_LMK2_CPS250")
	// The dedicated HSM is a payShield 10K, model PS10-D, 10Gb Ethernet Hardware Platform device with 2 local master keys which supports up to 2500 calls per second.
	SkuName_PayShield10K_LMK2_CPS2500 = SkuName("payShield10K_LMK2_CPS2500")
)

func init() {
}
