// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20180401
{
    public static class GetVirtualNetworkGatewayAdvertisedRoutes
    {
        /// <summary>
        /// List of virtual network gateway routes
        /// </summary>
        public static Task<GetVirtualNetworkGatewayAdvertisedRoutesResult> InvokeAsync(GetVirtualNetworkGatewayAdvertisedRoutesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVirtualNetworkGatewayAdvertisedRoutesResult>("azure-native:network/v20180401:getVirtualNetworkGatewayAdvertisedRoutes", args ?? new GetVirtualNetworkGatewayAdvertisedRoutesArgs(), options.WithDefaults());

        /// <summary>
        /// List of virtual network gateway routes
        /// </summary>
        public static Output<GetVirtualNetworkGatewayAdvertisedRoutesResult> Invoke(GetVirtualNetworkGatewayAdvertisedRoutesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVirtualNetworkGatewayAdvertisedRoutesResult>("azure-native:network/v20180401:getVirtualNetworkGatewayAdvertisedRoutes", args ?? new GetVirtualNetworkGatewayAdvertisedRoutesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVirtualNetworkGatewayAdvertisedRoutesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IP address of the peer
        /// </summary>
        [Input("peer", required: true)]
        public string Peer { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual network gateway.
        /// </summary>
        [Input("virtualNetworkGatewayName", required: true)]
        public string VirtualNetworkGatewayName { get; set; } = null!;

        public GetVirtualNetworkGatewayAdvertisedRoutesArgs()
        {
        }
        public static new GetVirtualNetworkGatewayAdvertisedRoutesArgs Empty => new GetVirtualNetworkGatewayAdvertisedRoutesArgs();
    }

    public sealed class GetVirtualNetworkGatewayAdvertisedRoutesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IP address of the peer
        /// </summary>
        [Input("peer", required: true)]
        public Input<string> Peer { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual network gateway.
        /// </summary>
        [Input("virtualNetworkGatewayName", required: true)]
        public Input<string> VirtualNetworkGatewayName { get; set; } = null!;

        public GetVirtualNetworkGatewayAdvertisedRoutesInvokeArgs()
        {
        }
        public static new GetVirtualNetworkGatewayAdvertisedRoutesInvokeArgs Empty => new GetVirtualNetworkGatewayAdvertisedRoutesInvokeArgs();
    }


    [OutputType]
    public sealed class GetVirtualNetworkGatewayAdvertisedRoutesResult
    {
        /// <summary>
        /// List of gateway routes
        /// </summary>
        public readonly ImmutableArray<Outputs.GatewayRouteResponse> Value;

        [OutputConstructor]
        private GetVirtualNetworkGatewayAdvertisedRoutesResult(ImmutableArray<Outputs.GatewayRouteResponse> value)
        {
            Value = value;
        }
    }
}
