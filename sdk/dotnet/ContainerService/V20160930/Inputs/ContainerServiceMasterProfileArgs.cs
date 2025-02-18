// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20160930.Inputs
{

    /// <summary>
    /// Profile for the container service master.
    /// </summary>
    public sealed class ContainerServiceMasterProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of masters (VMs) in the container service cluster. Allowed values are 1, 3, and 5. The default value is 1.
        /// </summary>
        [Input("count")]
        public Input<int>? Count { get; set; }

        /// <summary>
        /// DNS prefix to be used to create the FQDN for master.
        /// </summary>
        [Input("dnsPrefix", required: true)]
        public Input<string> DnsPrefix { get; set; } = null!;

        public ContainerServiceMasterProfileArgs()
        {
            Count = 1;
        }
        public static new ContainerServiceMasterProfileArgs Empty => new ContainerServiceMasterProfileArgs();
    }
}
