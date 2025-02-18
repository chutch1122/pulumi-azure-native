// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20160901.Inputs
{

    /// <summary>
    /// VirtualNetworkGatewaySku details
    /// </summary>
    public sealed class VirtualNetworkGatewaySkuArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The capacity.
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// Gateway SKU name. Possible values are: 'Basic', 'HighPerformance','Standard', and 'UltraPerformance'.
        /// </summary>
        [Input("name", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Network.V20160901.VirtualNetworkGatewaySkuName> Name { get; set; } = null!;

        /// <summary>
        /// Gateway SKU tier. Possible values are: 'Basic', 'HighPerformance','Standard', and 'UltraPerformance'.
        /// </summary>
        [Input("tier", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Network.V20160901.VirtualNetworkGatewaySkuTier> Tier { get; set; } = null!;

        public VirtualNetworkGatewaySkuArgs()
        {
        }
        public static new VirtualNetworkGatewaySkuArgs Empty => new VirtualNetworkGatewaySkuArgs();
    }
}
