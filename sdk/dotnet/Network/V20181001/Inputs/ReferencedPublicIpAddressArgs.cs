// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20181001.Inputs
{

    public sealed class ReferencedPublicIpAddressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The PublicIPAddress Reference
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public ReferencedPublicIpAddressArgs()
        {
        }
        public static new ReferencedPublicIpAddressArgs Empty => new ReferencedPublicIpAddressArgs();
    }
}
