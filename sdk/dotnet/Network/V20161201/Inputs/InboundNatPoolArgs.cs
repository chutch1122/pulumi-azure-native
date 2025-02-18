// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20161201.Inputs
{

    /// <summary>
    /// Inbound NAT pool of the load balancer.
    /// </summary>
    public sealed class InboundNatPoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The port used for internal connections on the endpoint. Acceptable values are between 1 and 65535.
        /// </summary>
        [Input("backendPort", required: true)]
        public Input<int> BackendPort { get; set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// A reference to frontend IP addresses.
        /// </summary>
        [Input("frontendIPConfiguration")]
        public Input<Inputs.SubResourceArgs>? FrontendIPConfiguration { get; set; }

        /// <summary>
        /// The last port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with a load balancer. Acceptable values range between 1 and 65535.
        /// </summary>
        [Input("frontendPortRangeEnd", required: true)]
        public Input<int> FrontendPortRangeEnd { get; set; } = null!;

        /// <summary>
        /// The first port number in the range of external ports that will be used to provide Inbound Nat to NICs associated with a load balancer. Acceptable values range between 1 and 65534.
        /// </summary>
        [Input("frontendPortRangeStart", required: true)]
        public Input<int> FrontendPortRangeStart { get; set; } = null!;

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The transport protocol for the endpoint. Possible values are: 'Udp' or 'Tcp'.
        /// </summary>
        [Input("protocol", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Network.V20161201.TransportProtocol> Protocol { get; set; } = null!;

        /// <summary>
        /// Gets the provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        public InboundNatPoolArgs()
        {
        }
        public static new InboundNatPoolArgs Empty => new InboundNatPoolArgs();
    }
}
