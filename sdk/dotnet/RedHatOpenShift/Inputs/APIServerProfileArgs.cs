// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.RedHatOpenShift.Inputs
{

    /// <summary>
    /// APIServerProfile represents an API server profile.
    /// </summary>
    public sealed class APIServerProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP of the cluster API server (immutable).
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The URL to access the cluster API server (immutable).
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// API server visibility (immutable).
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public APIServerProfileArgs()
        {
        }
        public static new APIServerProfileArgs Empty => new APIServerProfileArgs();
    }
}
