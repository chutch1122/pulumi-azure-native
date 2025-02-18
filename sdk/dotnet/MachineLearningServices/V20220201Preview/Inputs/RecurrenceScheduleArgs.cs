// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20220201Preview.Inputs
{

    /// <summary>
    /// Recurrence schedule definition
    /// </summary>
    public sealed class RecurrenceScheduleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies end time of schedule in ISO 8601 format.
        /// If not present, the schedule will run indefinitely
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// [Required] Specifies frequency with with which to trigger schedule
        /// </summary>
        [Input("frequency", required: true)]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.V20220201Preview.RecurrenceFrequency> Frequency { get; set; } = null!;

        /// <summary>
        /// [Required] Specifies schedule interval in conjunction with frequency
        /// </summary>
        [Input("interval", required: true)]
        public Input<int> Interval { get; set; } = null!;

        /// <summary>
        /// Specifies the recurrence schedule pattern
        /// </summary>
        [Input("pattern")]
        public Input<Inputs.RecurrencePatternArgs>? Pattern { get; set; }

        /// <summary>
        /// Specifies the schedule's status
        /// </summary>
        [Input("scheduleStatus")]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.V20220201Preview.ScheduleStatus>? ScheduleStatus { get; set; }

        /// <summary>
        /// Enum to describe type of schedule
        /// Expected value is 'Recurrence'.
        /// </summary>
        [Input("scheduleType", required: true)]
        public Input<string> ScheduleType { get; set; } = null!;

        /// <summary>
        /// Specifies start time of schedule in ISO 8601 format.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// Specifies time zone in which the schedule runs.
        /// TimeZone should follow Windows time zone format.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public RecurrenceScheduleArgs()
        {
            ScheduleStatus = "Enabled";
            TimeZone = "UTC";
        }
        public static new RecurrenceScheduleArgs Empty => new RecurrenceScheduleArgs();
    }
}
