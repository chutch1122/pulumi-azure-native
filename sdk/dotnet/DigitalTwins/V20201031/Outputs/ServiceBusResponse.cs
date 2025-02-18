// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DigitalTwins.V20201031.Outputs
{

    /// <summary>
    /// Properties related to ServiceBus.
    /// </summary>
    [OutputType]
    public sealed class ServiceBusResponse
    {
        /// <summary>
        /// Time when the Endpoint was added to DigitalTwinsInstance.
        /// </summary>
        public readonly string CreatedTime;
        /// <summary>
        /// Dead letter storage secret. Will be obfuscated during read.
        /// </summary>
        public readonly string? DeadLetterSecret;
        /// <summary>
        /// The type of Digital Twins endpoint
        /// Expected value is 'ServiceBus'.
        /// </summary>
        public readonly string EndpointType;
        /// <summary>
        /// PrimaryConnectionString of the endpoint. Will be obfuscated during read.
        /// </summary>
        public readonly string PrimaryConnectionString;
        /// <summary>
        /// The provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// SecondaryConnectionString of the endpoint. Will be obfuscated during read.
        /// </summary>
        public readonly string? SecondaryConnectionString;

        [OutputConstructor]
        private ServiceBusResponse(
            string createdTime,

            string? deadLetterSecret,

            string endpointType,

            string primaryConnectionString,

            string provisioningState,

            string? secondaryConnectionString)
        {
            CreatedTime = createdTime;
            DeadLetterSecret = deadLetterSecret;
            EndpointType = endpointType;
            PrimaryConnectionString = primaryConnectionString;
            ProvisioningState = provisioningState;
            SecondaryConnectionString = secondaryConnectionString;
        }
    }
}
