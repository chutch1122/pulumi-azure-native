// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20210301Preview.Inputs
{

    /// <summary>
    /// The workflow trigger recurrence for ComputeStartStop schedule type.
    /// </summary>
    public sealed class RecurrenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The recurrence frequency.
        /// </summary>
        [Input("frequency")]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.V20210301Preview.RecurrenceFrequency>? Frequency { get; set; }

        /// <summary>
        /// The interval.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// The recurrence schedule
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.RecurrenceScheduleArgs>? Schedule { get; set; }

        /// <summary>
        /// The start time.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// The time zone.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        public RecurrenceArgs()
        {
        }
        public static new RecurrenceArgs Empty => new RecurrenceArgs();
    }
}
