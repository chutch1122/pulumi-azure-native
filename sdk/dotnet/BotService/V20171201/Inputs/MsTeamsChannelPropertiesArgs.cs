// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BotService.V20171201.Inputs
{

    /// <summary>
    /// The parameters to provide for the Microsoft Teams channel.
    /// </summary>
    public sealed class MsTeamsChannelPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable messaging for Microsoft Teams channel
        /// </summary>
        [Input("callMode")]
        public Input<string>? CallMode { get; set; }

        /// <summary>
        /// Enable calling for Microsoft Teams channel
        /// </summary>
        [Input("enableCalling")]
        public Input<bool>? EnableCalling { get; set; }

        /// <summary>
        /// Enable media cards for Microsoft Teams channel
        /// </summary>
        [Input("enableMediaCards")]
        public Input<bool>? EnableMediaCards { get; set; }

        /// <summary>
        /// Enable messaging for Microsoft Teams channel
        /// </summary>
        [Input("enableMessaging")]
        public Input<bool>? EnableMessaging { get; set; }

        /// <summary>
        /// Enable video for Microsoft Teams channel
        /// </summary>
        [Input("enableVideo")]
        public Input<bool>? EnableVideo { get; set; }

        /// <summary>
        /// Whether this channel is enabled for the bot
        /// </summary>
        [Input("isEnabled", required: true)]
        public Input<bool> IsEnabled { get; set; } = null!;

        public MsTeamsChannelPropertiesArgs()
        {
        }
        public static new MsTeamsChannelPropertiesArgs Empty => new MsTeamsChannelPropertiesArgs();
    }
}
