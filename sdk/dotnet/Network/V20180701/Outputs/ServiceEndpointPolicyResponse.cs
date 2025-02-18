// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20180701.Outputs
{

    /// <summary>
    /// Service End point policy resource.
    /// </summary>
    [OutputType]
    public sealed class ServiceEndpointPolicyResponse
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the service endpoint policy. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// The resource GUID property of the service endpoint policy resource.
        /// </summary>
        public readonly string? ResourceGuid;
        /// <summary>
        /// A collection of service endpoint policy definitions of the service endpoint policy.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceEndpointPolicyDefinitionResponse> ServiceEndpointPolicyDefinitions;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ServiceEndpointPolicyResponse(
            string? etag,

            string? id,

            string? location,

            string name,

            string? provisioningState,

            string? resourceGuid,

            ImmutableArray<Outputs.ServiceEndpointPolicyDefinitionResponse> serviceEndpointPolicyDefinitions,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Etag = etag;
            Id = id;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            ResourceGuid = resourceGuid;
            ServiceEndpointPolicyDefinitions = serviceEndpointPolicyDefinitions;
            Tags = tags;
            Type = type;
        }
    }
}
