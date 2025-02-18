// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.V20150801.Inputs
{

    /// <summary>
    /// Specification for using a virtual network
    /// </summary>
    public sealed class VirtualNetworkProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource id of the virtual network
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the virtual network (read-only)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Subnet within the virtual network
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        /// <summary>
        /// Resource type of the virtual network (read-only)
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public VirtualNetworkProfileArgs()
        {
        }
        public static new VirtualNetworkProfileArgs Empty => new VirtualNetworkProfileArgs();
    }
}
