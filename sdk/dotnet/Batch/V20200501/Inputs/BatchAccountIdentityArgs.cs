// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.V20200501.Inputs
{

    /// <summary>
    /// The identity of the Batch account, if configured. This is only used when the user specifies 'Microsoft.KeyVault' as their Batch account encryption configuration.
    /// </summary>
    public sealed class BatchAccountIdentityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of identity used for the Batch account.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AzureNative.Batch.V20200501.ResourceIdentityType> Type { get; set; } = null!;

        public BatchAccountIdentityArgs()
        {
        }
        public static new BatchAccountIdentityArgs Empty => new BatchAccountIdentityArgs();
    }
}
