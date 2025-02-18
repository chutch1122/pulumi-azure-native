// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AVS.V20200320.Inputs
{

    /// <summary>
    /// The properties of a default cluster
    /// </summary>
    public sealed class ManagementClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cluster size
        /// </summary>
        [Input("clusterSize", required: true)]
        public Input<int> ClusterSize { get; set; } = null!;

        public ManagementClusterArgs()
        {
        }
        public static new ManagementClusterArgs Empty => new ManagementClusterArgs();
    }
}
