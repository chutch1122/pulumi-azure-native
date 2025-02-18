// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RecoveryServices.V20160601.Inputs
{

    /// <summary>
    /// Retention duration.
    /// </summary>
    public sealed class RetentionDurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Count of the duration types. Retention duration is determined by the combining the Count times and durationType. 
        ///    For example, if Count = 3 and durationType = Weeks, then the retention duration is three weeks.
        /// </summary>
        [Input("count")]
        public Input<int>? Count { get; set; }

        /// <summary>
        /// The retention duration type of the retention policy.
        /// </summary>
        [Input("durationType")]
        public Input<Pulumi.AzureNative.RecoveryServices.V20160601.RetentionDurationType>? DurationType { get; set; }

        public RetentionDurationArgs()
        {
        }
        public static new RetentionDurationArgs Empty => new RetentionDurationArgs();
    }
}
