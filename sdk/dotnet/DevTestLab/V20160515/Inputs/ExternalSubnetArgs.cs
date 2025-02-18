// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevTestLab.V20160515.Inputs
{

    /// <summary>
    /// Subnet information as returned by the Microsoft.Network API.
    /// </summary>
    public sealed class ExternalSubnetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the identifier.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Gets or sets the name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ExternalSubnetArgs()
        {
        }
        public static new ExternalSubnetArgs Empty => new ExternalSubnetArgs();
    }
}
