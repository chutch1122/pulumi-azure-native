// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20150501Preview.Inputs
{

    /// <summary>
    /// Contains sku in an ExpressRouteCircuit
    /// </summary>
    public sealed class ExpressRouteCircuitSkuArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets family of the sku.
        /// </summary>
        [Input("family")]
        public InputUnion<string, Pulumi.AzureNative.Network.V20150501Preview.ExpressRouteCircuitSkuFamily>? Family { get; set; }

        /// <summary>
        /// Gets or sets name of the sku.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Gets or sets tier of the sku.
        /// </summary>
        [Input("tier")]
        public InputUnion<string, Pulumi.AzureNative.Network.V20150501Preview.ExpressRouteCircuitSkuTier>? Tier { get; set; }

        public ExpressRouteCircuitSkuArgs()
        {
        }
        public static new ExpressRouteCircuitSkuArgs Empty => new ExpressRouteCircuitSkuArgs();
    }
}
