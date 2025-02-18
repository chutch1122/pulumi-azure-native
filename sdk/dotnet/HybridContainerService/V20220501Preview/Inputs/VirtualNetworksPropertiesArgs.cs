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
    /// HybridAKSNetworkSpec defines the desired state of HybridAKSNetwork
    /// </summary>
    public sealed class VirtualNetworksPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("infraVnetProfile")]
        public Input<Inputs.VirtualNetworksPropertiesInfraVnetProfileArgs>? InfraVnetProfile { get; set; }

        [Input("vipPool")]
        private InputList<Inputs.VirtualNetworksPropertiesVipPoolArgs>? _vipPool;

        /// <summary>
        /// Virtual IP Pool for Kubernetes
        /// </summary>
        public InputList<Inputs.VirtualNetworksPropertiesVipPoolArgs> VipPool
        {
            get => _vipPool ?? (_vipPool = new InputList<Inputs.VirtualNetworksPropertiesVipPoolArgs>());
            set => _vipPool = value;
        }

        [Input("vmipPool")]
        private InputList<Inputs.VirtualNetworksPropertiesVmipPoolArgs>? _vmipPool;

        /// <summary>
        /// IP Pool for Virtual Machines
        /// </summary>
        public InputList<Inputs.VirtualNetworksPropertiesVmipPoolArgs> VmipPool
        {
            get => _vmipPool ?? (_vmipPool = new InputList<Inputs.VirtualNetworksPropertiesVmipPoolArgs>());
            set => _vmipPool = value;
        }

        public VirtualNetworksPropertiesArgs()
        {
        }
        public static new VirtualNetworksPropertiesArgs Empty => new VirtualNetworksPropertiesArgs();
    }
}
