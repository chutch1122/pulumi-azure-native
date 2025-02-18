// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventHub.V20150801.Inputs
{

    /// <summary>
    /// SKU parameters supplied to the create Namespace operation
    /// </summary>
    public sealed class SkuArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Event Hubs throughput units.
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// Name of this SKU.
        /// </summary>
        [Input("name")]
        public InputUnion<string, Pulumi.AzureNative.EventHub.V20150801.SkuName>? Name { get; set; }

        /// <summary>
        /// The billing tier of this particular SKU.
        /// </summary>
        [Input("tier", required: true)]
        public InputUnion<string, Pulumi.AzureNative.EventHub.V20150801.SkuTier> Tier { get; set; } = null!;

        public SkuArgs()
        {
        }
        public static new SkuArgs Empty => new SkuArgs();
    }
}
