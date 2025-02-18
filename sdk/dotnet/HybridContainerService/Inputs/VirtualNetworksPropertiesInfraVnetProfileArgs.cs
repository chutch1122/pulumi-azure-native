// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService.Inputs
{

    public sealed class VirtualNetworksPropertiesInfraVnetProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Infra network profile for HCI platform
        /// </summary>
        [Input("hci")]
        public Input<Inputs.VirtualNetworksPropertiesHciArgs>? Hci { get; set; }

        /// <summary>
        /// Infra network profile for KubeVirt platform
        /// </summary>
        [Input("kubevirt")]
        public Input<Inputs.VirtualNetworksPropertiesKubevirtArgs>? Kubevirt { get; set; }

        /// <summary>
        /// Infra network profile for VMware platform
        /// </summary>
        [Input("vmware")]
        public Input<Inputs.VirtualNetworksPropertiesVmwareArgs>? Vmware { get; set; }

        public VirtualNetworksPropertiesInfraVnetProfileArgs()
        {
        }
        public static new VirtualNetworksPropertiesInfraVnetProfileArgs Empty => new VirtualNetworksPropertiesInfraVnetProfileArgs();
    }
}
