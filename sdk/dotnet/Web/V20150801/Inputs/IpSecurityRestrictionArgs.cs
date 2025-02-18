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
    /// Represents an ip security restriction on a web app.
    /// </summary>
    public sealed class IpSecurityRestrictionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP address the security restriction is valid for
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Subnet mask for the range of IP addresses the restriction is valid for
        /// </summary>
        [Input("subnetMask")]
        public Input<string>? SubnetMask { get; set; }

        public IpSecurityRestrictionArgs()
        {
        }
        public static new IpSecurityRestrictionArgs Empty => new IpSecurityRestrictionArgs();
    }
}
