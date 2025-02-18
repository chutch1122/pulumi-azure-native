// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.WebPubSub.V20210901Preview.Outputs
{

    /// <summary>
    /// live trace category configuration of a Microsoft.SignalRService resource.
    /// </summary>
    [OutputType]
    public sealed class LiveTraceCategoryResponse
    {
        /// <summary>
        /// Indicates whether or the log category is enabled.
        /// Available values: true, false.
        /// Case insensitive.
        /// </summary>
        public readonly string? Enabled;
        /// <summary>
        /// Gets or sets the log category's name.
        /// Available values: ConnectivityLogs, MessagingLogs.
        /// Case insensitive.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private LiveTraceCategoryResponse(
            string? enabled,

            string? name)
        {
            Enabled = enabled;
            Name = name;
        }
    }
}
