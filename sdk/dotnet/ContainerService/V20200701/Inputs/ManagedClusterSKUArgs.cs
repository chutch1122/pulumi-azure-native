// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20200701.Inputs
{

    public sealed class ManagedClusterSKUArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of a managed cluster SKU.
        /// </summary>
        [Input("name")]
        public InputUnion<string, Pulumi.AzureNative.ContainerService.V20200701.ManagedClusterSKUName>? Name { get; set; }

        /// <summary>
        /// Tier of a managed cluster SKU.
        /// </summary>
        [Input("tier")]
        public InputUnion<string, Pulumi.AzureNative.ContainerService.V20200701.ManagedClusterSKUTier>? Tier { get; set; }

        public ManagedClusterSKUArgs()
        {
        }
        public static new ManagedClusterSKUArgs Empty => new ManagedClusterSKUArgs();
    }
}
