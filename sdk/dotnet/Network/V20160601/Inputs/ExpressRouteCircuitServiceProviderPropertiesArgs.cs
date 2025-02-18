// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20160601.Inputs
{

    /// <summary>
    /// Contains ServiceProviderProperties in an ExpressRouteCircuit
    /// </summary>
    public sealed class ExpressRouteCircuitServiceProviderPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets BandwidthInMbps.
        /// </summary>
        [Input("bandwidthInMbps")]
        public Input<int>? BandwidthInMbps { get; set; }

        /// <summary>
        /// Gets or sets peering location.
        /// </summary>
        [Input("peeringLocation")]
        public Input<string>? PeeringLocation { get; set; }

        /// <summary>
        /// Gets or sets serviceProviderName.
        /// </summary>
        [Input("serviceProviderName")]
        public Input<string>? ServiceProviderName { get; set; }

        public ExpressRouteCircuitServiceProviderPropertiesArgs()
        {
        }
        public static new ExpressRouteCircuitServiceProviderPropertiesArgs Empty => new ExpressRouteCircuitServiceProviderPropertiesArgs();
    }
}
