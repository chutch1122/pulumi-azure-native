// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService.V20220501Preview.Inputs
{

    /// <summary>
    /// Infra network profile for KubeVirt platform
    /// </summary>
    public sealed class VirtualNetworksPropertiesKubevirtArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the network in KubeVirt
        /// </summary>
        [Input("vnetName")]
        public Input<string>? VnetName { get; set; }

        public VirtualNetworksPropertiesKubevirtArgs()
        {
        }
        public static new VirtualNetworksPropertiesKubevirtArgs Empty => new VirtualNetworksPropertiesKubevirtArgs();
    }
}
