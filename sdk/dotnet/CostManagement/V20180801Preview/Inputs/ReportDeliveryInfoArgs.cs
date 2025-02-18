// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CostManagement.V20180801Preview.Inputs
{

    /// <summary>
    /// The delivery information associated with a report.
    /// </summary>
    public sealed class ReportDeliveryInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Has destination for the report being delivered.
        /// </summary>
        [Input("destination", required: true)]
        public Input<Inputs.ReportDeliveryDestinationArgs> Destination { get; set; } = null!;

        public ReportDeliveryInfoArgs()
        {
        }
        public static new ReportDeliveryInfoArgs Empty => new ReportDeliveryInfoArgs();
    }
}
