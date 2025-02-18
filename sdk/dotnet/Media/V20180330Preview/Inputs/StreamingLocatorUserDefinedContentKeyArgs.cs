// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media.V20180330Preview.Inputs
{

    /// <summary>
    /// Describes the properties of a user-defined content key in the Streaming Locator
    /// </summary>
    public sealed class StreamingLocatorUserDefinedContentKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of Content Key
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The Content Key description
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// The Content Key secret
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public StreamingLocatorUserDefinedContentKeyArgs()
        {
        }
        public static new StreamingLocatorUserDefinedContentKeyArgs Empty => new StreamingLocatorUserDefinedContentKeyArgs();
    }
}
