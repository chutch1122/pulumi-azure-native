// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DigitalTwins.V20201031.Inputs
{

    /// <summary>
    /// Properties related to ServiceBus.
    /// </summary>
    public sealed class ServiceBusArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dead letter storage secret. Will be obfuscated during read.
        /// </summary>
        [Input("deadLetterSecret")]
        public Input<string>? DeadLetterSecret { get; set; }

        /// <summary>
        /// The type of Digital Twins endpoint
        /// Expected value is 'ServiceBus'.
        /// </summary>
        [Input("endpointType", required: true)]
        public Input<string> EndpointType { get; set; } = null!;

        /// <summary>
        /// PrimaryConnectionString of the endpoint. Will be obfuscated during read.
        /// </summary>
        [Input("primaryConnectionString", required: true)]
        public Input<string> PrimaryConnectionString { get; set; } = null!;

        /// <summary>
        /// SecondaryConnectionString of the endpoint. Will be obfuscated during read.
        /// </summary>
        [Input("secondaryConnectionString")]
        public Input<string>? SecondaryConnectionString { get; set; }

        public ServiceBusArgs()
        {
        }
        public static new ServiceBusArgs Empty => new ServiceBusArgs();
    }
}
