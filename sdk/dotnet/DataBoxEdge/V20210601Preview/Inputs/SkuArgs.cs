// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBoxEdge.V20210601Preview.Inputs
{

    /// <summary>
    /// The SKU type.
    /// </summary>
    public sealed class SkuArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SKU name.
        /// </summary>
        [Input("name")]
        public InputUnion<string, Pulumi.AzureNative.DataBoxEdge.V20210601Preview.SkuName>? Name { get; set; }

        /// <summary>
        /// The SKU tier. This is based on the SKU name.
        /// </summary>
        [Input("tier")]
        public InputUnion<string, Pulumi.AzureNative.DataBoxEdge.V20210601Preview.SkuTier>? Tier { get; set; }

        public SkuArgs()
        {
        }
        public static new SkuArgs Empty => new SkuArgs();
    }
}
