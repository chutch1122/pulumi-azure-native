// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.OffAzure.V20200707.Inputs
{

    /// <summary>
    /// Class for site agent properties.
    /// </summary>
    public sealed class SiteAgentPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Key vault ARM Id.
        /// </summary>
        [Input("keyVaultId")]
        public Input<string>? KeyVaultId { get; set; }

        /// <summary>
        /// Key vault URI.
        /// </summary>
        [Input("keyVaultUri")]
        public Input<string>? KeyVaultUri { get; set; }

        public SiteAgentPropertiesArgs()
        {
        }
        public static new SiteAgentPropertiesArgs Empty => new SiteAgentPropertiesArgs();
    }
}
