// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Insights.V20210901Preview.Inputs
{

    /// <summary>
    /// The specification of data sources. 
    /// This property is optional and can be omitted if the rule is meant to be used via direct calls to the provisioned endpoint.
    /// </summary>
    public sealed class DataCollectionRuleDataSourcesArgs : global::Pulumi.ResourceArgs
    {
        [Input("extensions")]
        private InputList<Inputs.ExtensionDataSourceArgs>? _extensions;

        /// <summary>
        /// The list of Azure VM extension data source configurations.
        /// </summary>
        public InputList<Inputs.ExtensionDataSourceArgs> Extensions
        {
            get => _extensions ?? (_extensions = new InputList<Inputs.ExtensionDataSourceArgs>());
            set => _extensions = value;
        }

        [Input("iisLogs")]
        private InputList<Inputs.IisLogsDataSourceArgs>? _iisLogs;

        /// <summary>
        /// The list of IIS logs source configurations.
        /// </summary>
        public InputList<Inputs.IisLogsDataSourceArgs> IisLogs
        {
            get => _iisLogs ?? (_iisLogs = new InputList<Inputs.IisLogsDataSourceArgs>());
            set => _iisLogs = value;
        }

        [Input("logFiles")]
        private InputList<Inputs.LogFilesDataSourceArgs>? _logFiles;

        /// <summary>
        /// The list of Log files source configurations.
        /// </summary>
        public InputList<Inputs.LogFilesDataSourceArgs> LogFiles
        {
            get => _logFiles ?? (_logFiles = new InputList<Inputs.LogFilesDataSourceArgs>());
            set => _logFiles = value;
        }

        [Input("performanceCounters")]
        private InputList<Inputs.PerfCounterDataSourceArgs>? _performanceCounters;

        /// <summary>
        /// The list of performance counter data source configurations.
        /// </summary>
        public InputList<Inputs.PerfCounterDataSourceArgs> PerformanceCounters
        {
            get => _performanceCounters ?? (_performanceCounters = new InputList<Inputs.PerfCounterDataSourceArgs>());
            set => _performanceCounters = value;
        }

        [Input("syslog")]
        private InputList<Inputs.SyslogDataSourceArgs>? _syslog;

        /// <summary>
        /// The list of Syslog data source configurations.
        /// </summary>
        public InputList<Inputs.SyslogDataSourceArgs> Syslog
        {
            get => _syslog ?? (_syslog = new InputList<Inputs.SyslogDataSourceArgs>());
            set => _syslog = value;
        }

        [Input("windowsEventLogs")]
        private InputList<Inputs.WindowsEventLogDataSourceArgs>? _windowsEventLogs;

        /// <summary>
        /// The list of Windows Event Log data source configurations.
        /// </summary>
        public InputList<Inputs.WindowsEventLogDataSourceArgs> WindowsEventLogs
        {
            get => _windowsEventLogs ?? (_windowsEventLogs = new InputList<Inputs.WindowsEventLogDataSourceArgs>());
            set => _windowsEventLogs = value;
        }

        public DataCollectionRuleDataSourcesArgs()
        {
        }
        public static new DataCollectionRuleDataSourcesArgs Empty => new DataCollectionRuleDataSourcesArgs();
    }
}
