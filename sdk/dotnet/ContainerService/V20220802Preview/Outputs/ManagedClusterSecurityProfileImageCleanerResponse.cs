// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20220802Preview.Outputs
{

    /// <summary>
    /// ImageCleaner removes unused images from nodes, freeing up disk space and helping to reduce attack surface area. Here are settings for the security profile.
    /// </summary>
    [OutputType]
    public sealed class ManagedClusterSecurityProfileImageCleanerResponse
    {
        /// <summary>
        /// Whether to enable ImageCleaner on AKS cluster.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// ImageCleaner scanning interval.
        /// </summary>
        public readonly int? IntervalHours;

        [OutputConstructor]
        private ManagedClusterSecurityProfileImageCleanerResponse(
            bool? enabled,

            int? intervalHours)
        {
            Enabled = enabled;
            IntervalHours = intervalHours;
        }
    }
}
