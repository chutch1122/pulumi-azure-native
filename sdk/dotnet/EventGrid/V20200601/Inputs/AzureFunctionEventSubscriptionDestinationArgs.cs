// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid.V20200601.Inputs
{

    /// <summary>
    /// Information about the azure function destination for an event subscription.
    /// </summary>
    public sealed class AzureFunctionEventSubscriptionDestinationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of the endpoint for the event subscription destination.
        /// Expected value is 'AzureFunction'.
        /// </summary>
        [Input("endpointType", required: true)]
        public Input<string> EndpointType { get; set; } = null!;

        /// <summary>
        /// Maximum number of events per batch.
        /// </summary>
        [Input("maxEventsPerBatch")]
        public Input<int>? MaxEventsPerBatch { get; set; }

        /// <summary>
        /// Preferred batch size in Kilobytes.
        /// </summary>
        [Input("preferredBatchSizeInKilobytes")]
        public Input<int>? PreferredBatchSizeInKilobytes { get; set; }

        /// <summary>
        /// The Azure Resource Id that represents the endpoint of the Azure Function destination of an event subscription.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        public AzureFunctionEventSubscriptionDestinationArgs()
        {
            MaxEventsPerBatch = 1;
            PreferredBatchSizeInKilobytes = 64;
        }
        public static new AzureFunctionEventSubscriptionDestinationArgs Empty => new AzureFunctionEventSubscriptionDestinationArgs();
    }
}
