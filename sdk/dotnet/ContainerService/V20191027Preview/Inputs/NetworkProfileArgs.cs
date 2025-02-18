// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20191027Preview.Inputs
{

    /// <summary>
    /// Represents the OpenShift networking configuration
    /// </summary>
    public sealed class NetworkProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CIDR of subnet used to create PLS needed for management of the cluster
        /// </summary>
        [Input("managementSubnetCidr")]
        public Input<string>? ManagementSubnetCidr { get; set; }

        /// <summary>
        /// CIDR for the OpenShift Vnet.
        /// </summary>
        [Input("vnetCidr")]
        public Input<string>? VnetCidr { get; set; }

        /// <summary>
        /// ID of the Vnet created for OSA cluster.
        /// </summary>
        [Input("vnetId")]
        public Input<string>? VnetId { get; set; }

        public NetworkProfileArgs()
        {
            VnetCidr = "10.0.0.0/8";
        }
        public static new NetworkProfileArgs Empty => new NetworkProfileArgs();
    }
}
