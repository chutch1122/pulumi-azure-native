


package v20160201preview

type Kind string

const (
	KindAcademic           = Kind("Academic")
	Kind_Bing_Autosuggest  = Kind("Bing.Autosuggest")
	Kind_Bing_Search       = Kind("Bing.Search")
	Kind_Bing_Speech       = Kind("Bing.Speech")
	Kind_Bing_SpellCheck   = Kind("Bing.SpellCheck")
	KindComputerVision     = Kind("ComputerVision")
	KindContentModerator   = Kind("ContentModerator")
	KindEmotion            = Kind("Emotion")
	KindFace               = Kind("Face")
	KindLUIS               = Kind("LUIS")
	KindRecommendations    = Kind("Recommendations")
	KindSpeakerRecognition = Kind("SpeakerRecognition")
	KindSpeech             = Kind("Speech")
	KindSpeechTranslation  = Kind("SpeechTranslation")
	KindTextAnalytics      = Kind("TextAnalytics")
	KindTextTranslation    = Kind("TextTranslation")
	KindWebLM              = Kind("WebLM")
)

type SkuName string

const (
	SkuNameF0 = SkuName("F0")
	SkuNameP0 = SkuName("P0")
	SkuNameP1 = SkuName("P1")
	SkuNameP2 = SkuName("P2")
	SkuNameS0 = SkuName("S0")
	SkuNameS1 = SkuName("S1")
	SkuNameS2 = SkuName("S2")
	SkuNameS3 = SkuName("S3")
	SkuNameS4 = SkuName("S4")
	SkuNameS5 = SkuName("S5")
	SkuNameS6 = SkuName("S6")
)

func init() {
}
