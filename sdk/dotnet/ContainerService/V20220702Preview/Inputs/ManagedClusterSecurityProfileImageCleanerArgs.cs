// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20220702Preview.Inputs
{

    /// <summary>
    /// ImageCleaner removes unused images from nodes, freeing up disk space and helping to reduce attack surface area. Here are settings for the security profile.
    /// </summary>
    public sealed class ManagedClusterSecurityProfileImageCleanerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable ImageCleaner on AKS cluster.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// ImageCleaner scanning interval.
        /// </summary>
        [Input("intervalHours")]
        public Input<int>? IntervalHours { get; set; }

        public ManagedClusterSecurityProfileImageCleanerArgs()
        {
        }
        public static new ManagedClusterSecurityProfileImageCleanerArgs Empty => new ManagedClusterSecurityProfileImageCleanerArgs();
    }
}
