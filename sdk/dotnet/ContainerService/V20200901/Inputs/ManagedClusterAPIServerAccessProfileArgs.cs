// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20200901.Inputs
{

    /// <summary>
    /// Access profile for managed cluster API server.
    /// </summary>
    public sealed class ManagedClusterAPIServerAccessProfileArgs : global::Pulumi.ResourceArgs
    {
        [Input("authorizedIPRanges")]
        private InputList<string>? _authorizedIPRanges;

        /// <summary>
        /// Authorized IP Ranges to kubernetes API server.
        /// </summary>
        public InputList<string> AuthorizedIPRanges
        {
            get => _authorizedIPRanges ?? (_authorizedIPRanges = new InputList<string>());
            set => _authorizedIPRanges = value;
        }

        /// <summary>
        /// Whether to create the cluster as a private cluster or not.
        /// </summary>
        [Input("enablePrivateCluster")]
        public Input<bool>? EnablePrivateCluster { get; set; }

        public ManagedClusterAPIServerAccessProfileArgs()
        {
        }
        public static new ManagedClusterAPIServerAccessProfileArgs Empty => new ManagedClusterAPIServerAccessProfileArgs();
    }
}
