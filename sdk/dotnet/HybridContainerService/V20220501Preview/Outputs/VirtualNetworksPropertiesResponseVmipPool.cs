// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService.V20220501Preview.Outputs
{

    [OutputType]
    public sealed class VirtualNetworksPropertiesResponseVmipPool
    {
        /// <summary>
        /// Ending IP address for the IP Pool
        /// </summary>
        public readonly string? EndIP;
        /// <summary>
        /// Starting IP address for the IP Pool
        /// </summary>
        public readonly string? StartIP;

        [OutputConstructor]
        private VirtualNetworksPropertiesResponseVmipPool(
            string? endIP,

            string? startIP)
        {
            EndIP = endIP;
            StartIP = startIP;
        }
    }
}
