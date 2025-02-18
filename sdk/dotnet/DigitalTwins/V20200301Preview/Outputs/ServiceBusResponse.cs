// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DigitalTwins.V20200301Preview.Outputs
{

    /// <summary>
    /// properties related to servicebus.
    /// </summary>
    [OutputType]
    public sealed class ServiceBusResponse
    {
        /// <summary>
        /// Time when the Endpoint was added to DigitalTwinsInstance.
        /// </summary>
        public readonly string CreatedTime;
        /// <summary>
        /// The type of Digital Twins endpoint
        /// Expected value is 'ServiceBus'.
        /// </summary>
        public readonly string EndpointType;
        /// <summary>
        /// PrimaryConnectionString of the endpoint. Will be obfuscated during read
        /// </summary>
        public readonly string PrimaryConnectionString;
        /// <summary>
        /// The provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// SecondaryConnectionString of the endpoint. Will be obfuscated during read
        /// </summary>
        public readonly string SecondaryConnectionString;
        /// <summary>
        /// The resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private ServiceBusResponse(
            string createdTime,

            string endpointType,

            string primaryConnectionString,

            string provisioningState,

            string secondaryConnectionString,

            ImmutableDictionary<string, string>? tags)
        {
            CreatedTime = createdTime;
            EndpointType = endpointType;
            PrimaryConnectionString = primaryConnectionString;
            ProvisioningState = provisioningState;
            SecondaryConnectionString = secondaryConnectionString;
            Tags = tags;
        }
    }
}
