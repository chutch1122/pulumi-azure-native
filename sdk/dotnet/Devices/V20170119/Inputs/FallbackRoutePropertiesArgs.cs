// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Devices.V20170119.Inputs
{

    /// <summary>
    /// The properties related to the fallback route based on which the IoT hub routes messages to the fallback endpoint.
    /// </summary>
    public sealed class FallbackRoutePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The condition which is evaluated in order to apply the fallback route. If the condition is not provided it will evaluate to true by default. For grammar, See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language
        /// </summary>
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        [Input("endpointNames", required: true)]
        private InputList<string>? _endpointNames;

        /// <summary>
        /// The list of endpoints to which the messages that satisfy the condition are routed to. Currently only 1 endpoint is allowed.
        /// </summary>
        public InputList<string> EndpointNames
        {
            get => _endpointNames ?? (_endpointNames = new InputList<string>());
            set => _endpointNames = value;
        }

        /// <summary>
        /// Used to specify whether the fallback route is enabled or not.
        /// </summary>
        [Input("isEnabled", required: true)]
        public Input<bool> IsEnabled { get; set; } = null!;

        /// <summary>
        /// The source to which the routing rule is to be applied to. e.g. DeviceMessages
        /// </summary>
        [Input("source", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Devices.V20170119.RoutingSource> Source { get; set; } = null!;

        public FallbackRoutePropertiesArgs()
        {
        }
        public static new FallbackRoutePropertiesArgs Empty => new FallbackRoutePropertiesArgs();
    }
}
