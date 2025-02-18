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
    /// Compute start stop schedule properties
    /// </summary>
    public sealed class ComputeStartStopScheduleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The compute power action.
        /// </summary>
        [Input("action")]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.V20210301Preview.ComputePowerAction>? Action { get; set; }

        /// <summary>
        /// The workflow trigger cron for ComputeStartStop schedule type.
        /// </summary>
        [Input("cron")]
        public Input<Inputs.CronArgs>? Cron { get; set; }

        /// <summary>
        /// The workflow trigger recurrence for ComputeStartStop schedule type.
        /// </summary>
        [Input("recurrence")]
        public Input<Inputs.RecurrenceArgs>? Recurrence { get; set; }

        /// <summary>
        /// The schedule status.
        /// </summary>
        [Input("status")]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.V20210301Preview.ScheduleStatus>? Status { get; set; }

        /// <summary>
        /// The schedule trigger type.
        /// </summary>
        [Input("triggerType")]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.V20210301Preview.TriggerType>? TriggerType { get; set; }

        public ComputeStartStopScheduleArgs()
        {
        }
        public static new ComputeStartStopScheduleArgs Empty => new ComputeStartStopScheduleArgs();
    }
}
