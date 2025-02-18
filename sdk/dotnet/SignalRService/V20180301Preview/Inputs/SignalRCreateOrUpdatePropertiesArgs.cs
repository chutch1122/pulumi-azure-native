// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SignalRService.V20180301Preview.Inputs
{

    /// <summary>
    /// Settings used to provision or configure the resource.
    /// </summary>
    public sealed class SignalRCreateOrUpdatePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Prefix for the hostName of the SignalR service. Retained for future use.
        /// The hostname will be of format: &amp;lt;hostNamePrefix&amp;gt;.service.signalr.net.
        /// </summary>
        [Input("hostNamePrefix")]
        public Input<string>? HostNamePrefix { get; set; }

        public SignalRCreateOrUpdatePropertiesArgs()
        {
        }
        public static new SignalRCreateOrUpdatePropertiesArgs Empty => new SignalRCreateOrUpdatePropertiesArgs();
    }
}
