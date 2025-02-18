// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20210301.Inputs
{

    /// <summary>
    /// The time span with start and end properties.
    /// </summary>
    public sealed class TimeSpanArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The end of a time span
        /// </summary>
        [Input("end")]
        public Input<string>? End { get; set; }

        /// <summary>
        /// The start of a time span
        /// </summary>
        [Input("start")]
        public Input<string>? Start { get; set; }

        public TimeSpanArgs()
        {
        }
        public static new TimeSpanArgs Empty => new TimeSpanArgs();
    }
}
