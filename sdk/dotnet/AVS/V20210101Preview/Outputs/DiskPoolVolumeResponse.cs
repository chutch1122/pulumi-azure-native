// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AVS.V20210101Preview.Outputs
{

    /// <summary>
    /// An iSCSI volume from Microsoft.StoragePool provider
    /// </summary>
    [OutputType]
    public sealed class DiskPoolVolumeResponse
    {
        /// <summary>
        /// iSCSI provider target IP address list
        /// </summary>
        public readonly ImmutableArray<string> Endpoints;
        /// <summary>
        /// Name of the LUN to be used
        /// </summary>
        public readonly string? LunName;

        [OutputConstructor]
        private DiskPoolVolumeResponse(
            ImmutableArray<string> endpoints,

            string? lunName)
        {
            Endpoints = endpoints;
            LunName = lunName;
        }
    }
}
