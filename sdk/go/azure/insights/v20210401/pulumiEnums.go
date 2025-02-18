


package v20210401

type KnownDataCollectionEndpointResourceKind string

const (
	KnownDataCollectionEndpointResourceKindLinux   = KnownDataCollectionEndpointResourceKind("Linux")
	KnownDataCollectionEndpointResourceKindWindows = KnownDataCollectionEndpointResourceKind("Windows")
)

type KnownDataCollectionRuleResourceKind string

const (
	KnownDataCollectionRuleResourceKindLinux   = KnownDataCollectionRuleResourceKind("Linux")
	KnownDataCollectionRuleResourceKindWindows = KnownDataCollectionRuleResourceKind("Windows")
)

type KnownDataFlowStreams string

const (
	KnownDataFlowStreams_Microsoft_Event           = KnownDataFlowStreams("Microsoft-Event")
	KnownDataFlowStreams_Microsoft_InsightsMetrics = KnownDataFlowStreams("Microsoft-InsightsMetrics")
	KnownDataFlowStreams_Microsoft_Perf            = KnownDataFlowStreams("Microsoft-Perf")
	KnownDataFlowStreams_Microsoft_Syslog          = KnownDataFlowStreams("Microsoft-Syslog")
	KnownDataFlowStreams_Microsoft_WindowsEvent    = KnownDataFlowStreams("Microsoft-WindowsEvent")
)

type KnownExtensionDataSourceStreams string

const (
	KnownExtensionDataSourceStreams_Microsoft_Event           = KnownExtensionDataSourceStreams("Microsoft-Event")
	KnownExtensionDataSourceStreams_Microsoft_InsightsMetrics = KnownExtensionDataSourceStreams("Microsoft-InsightsMetrics")
	KnownExtensionDataSourceStreams_Microsoft_Perf            = KnownExtensionDataSourceStreams("Microsoft-Perf")
	KnownExtensionDataSourceStreams_Microsoft_Syslog          = KnownExtensionDataSourceStreams("Microsoft-Syslog")
	KnownExtensionDataSourceStreams_Microsoft_WindowsEvent    = KnownExtensionDataSourceStreams("Microsoft-WindowsEvent")
)

type KnownPerfCounterDataSourceStreams string

const (
	KnownPerfCounterDataSourceStreams_Microsoft_Perf            = KnownPerfCounterDataSourceStreams("Microsoft-Perf")
	KnownPerfCounterDataSourceStreams_Microsoft_InsightsMetrics = KnownPerfCounterDataSourceStreams("Microsoft-InsightsMetrics")
)

type KnownPublicNetworkAccessOptions string

const (
	KnownPublicNetworkAccessOptionsEnabled  = KnownPublicNetworkAccessOptions("Enabled")
	KnownPublicNetworkAccessOptionsDisabled = KnownPublicNetworkAccessOptions("Disabled")
)

type KnownSyslogDataSourceFacilityNames string

const (
	KnownSyslogDataSourceFacilityNamesAuth     = KnownSyslogDataSourceFacilityNames("auth")
	KnownSyslogDataSourceFacilityNamesAuthpriv = KnownSyslogDataSourceFacilityNames("authpriv")
	KnownSyslogDataSourceFacilityNamesCron     = KnownSyslogDataSourceFacilityNames("cron")
	KnownSyslogDataSourceFacilityNamesDaemon   = KnownSyslogDataSourceFacilityNames("daemon")
	KnownSyslogDataSourceFacilityNamesKern     = KnownSyslogDataSourceFacilityNames("kern")
	KnownSyslogDataSourceFacilityNamesLpr      = KnownSyslogDataSourceFacilityNames("lpr")
	KnownSyslogDataSourceFacilityNamesMail     = KnownSyslogDataSourceFacilityNames("mail")
	KnownSyslogDataSourceFacilityNamesMark     = KnownSyslogDataSourceFacilityNames("mark")
	KnownSyslogDataSourceFacilityNamesNews     = KnownSyslogDataSourceFacilityNames("news")
	KnownSyslogDataSourceFacilityNamesSyslog   = KnownSyslogDataSourceFacilityNames("syslog")
	KnownSyslogDataSourceFacilityNamesUser     = KnownSyslogDataSourceFacilityNames("user")
	KnownSyslogDataSourceFacilityNamesUucp     = KnownSyslogDataSourceFacilityNames("uucp")
	KnownSyslogDataSourceFacilityNamesLocal0   = KnownSyslogDataSourceFacilityNames("local0")
	KnownSyslogDataSourceFacilityNamesLocal1   = KnownSyslogDataSourceFacilityNames("local1")
	KnownSyslogDataSourceFacilityNamesLocal2   = KnownSyslogDataSourceFacilityNames("local2")
	KnownSyslogDataSourceFacilityNamesLocal3   = KnownSyslogDataSourceFacilityNames("local3")
	KnownSyslogDataSourceFacilityNamesLocal4   = KnownSyslogDataSourceFacilityNames("local4")
	KnownSyslogDataSourceFacilityNamesLocal5   = KnownSyslogDataSourceFacilityNames("local5")
	KnownSyslogDataSourceFacilityNamesLocal6   = KnownSyslogDataSourceFacilityNames("local6")
	KnownSyslogDataSourceFacilityNamesLocal7   = KnownSyslogDataSourceFacilityNames("local7")
	KnownSyslogDataSourceFacilityNamesAsterisk = KnownSyslogDataSourceFacilityNames("*")
)

type KnownSyslogDataSourceLogLevels string

const (
	KnownSyslogDataSourceLogLevelsDebug     = KnownSyslogDataSourceLogLevels("Debug")
	KnownSyslogDataSourceLogLevelsInfo      = KnownSyslogDataSourceLogLevels("Info")
	KnownSyslogDataSourceLogLevelsNotice    = KnownSyslogDataSourceLogLevels("Notice")
	KnownSyslogDataSourceLogLevelsWarning   = KnownSyslogDataSourceLogLevels("Warning")
	KnownSyslogDataSourceLogLevelsError     = KnownSyslogDataSourceLogLevels("Error")
	KnownSyslogDataSourceLogLevelsCritical  = KnownSyslogDataSourceLogLevels("Critical")
	KnownSyslogDataSourceLogLevelsAlert     = KnownSyslogDataSourceLogLevels("Alert")
	KnownSyslogDataSourceLogLevelsEmergency = KnownSyslogDataSourceLogLevels("Emergency")
	KnownSyslogDataSourceLogLevelsAsterisk  = KnownSyslogDataSourceLogLevels("*")
)

type KnownSyslogDataSourceStreams string

const (
	KnownSyslogDataSourceStreams_Microsoft_Syslog = KnownSyslogDataSourceStreams("Microsoft-Syslog")
)

type KnownWindowsEventLogDataSourceStreams string

const (
	KnownWindowsEventLogDataSourceStreams_Microsoft_WindowsEvent = KnownWindowsEventLogDataSourceStreams("Microsoft-WindowsEvent")
	KnownWindowsEventLogDataSourceStreams_Microsoft_Event        = KnownWindowsEventLogDataSourceStreams("Microsoft-Event")
)

func init() {
}
