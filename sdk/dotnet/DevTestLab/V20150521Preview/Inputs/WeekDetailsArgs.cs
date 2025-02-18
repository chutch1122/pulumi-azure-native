// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab.V20150521Preview.Inputs
{

    /// <summary>
    /// Properties of a weekly schedule.
    /// </summary>
    public sealed class WeekDetailsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time of the day.
        /// </summary>
        [Input("time")]
        public Input<string>? Time { get; set; }

        [Input("weekdays")]
        private InputList<string>? _weekdays;

        /// <summary>
        /// The days of the week.
        /// </summary>
        public InputList<string> Weekdays
        {
            get => _weekdays ?? (_weekdays = new InputList<string>());
            set => _weekdays = value;
        }

        public WeekDetailsArgs()
        {
        }
        public static new WeekDetailsArgs Empty => new WeekDetailsArgs();
    }
}
