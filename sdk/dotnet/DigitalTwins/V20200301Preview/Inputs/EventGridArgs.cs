// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DigitalTwins.V20200301Preview.Inputs
{

    /// <summary>
    /// properties related to eventgrid.
    /// </summary>
    public sealed class EventGridArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// EventGrid secondary accesskey. Will be obfuscated during read
        /// </summary>
        [Input("accessKey1", required: true)]
        public Input<string> AccessKey1 { get; set; } = null!;

        /// <summary>
        /// EventGrid secondary accesskey. Will be obfuscated during read
        /// </summary>
        [Input("accessKey2", required: true)]
        public Input<string> AccessKey2 { get; set; } = null!;

        /// <summary>
        /// The type of Digital Twins endpoint
        /// Expected value is 'EventGrid'.
        /// </summary>
        [Input("endpointType", required: true)]
        public Input<string> EndpointType { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// EventGrid Topic Endpoint
        /// </summary>
        [Input("topicEndpoint")]
        public Input<string>? TopicEndpoint { get; set; }

        public EventGridArgs()
        {
        }
        public static new EventGridArgs Empty => new EventGridArgs();
    }
}
