// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Solutions.V20180301
{
    [Obsolete(@"Version 2018-03-01 will be removed in v2 of the provider.")]
    public static class GetApplication
    {
        /// <summary>
        /// Information about managed application.
        /// </summary>
        public static Task<GetApplicationResult> InvokeAsync(GetApplicationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetApplicationResult>("azure-native:solutions/v20180301:getApplication", args ?? new GetApplicationArgs(), options.WithDefaults());

        /// <summary>
        /// Information about managed application.
        /// </summary>
        public static Output<GetApplicationResult> Invoke(GetApplicationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetApplicationResult>("azure-native:solutions/v20180301:getApplication", args ?? new GetApplicationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the managed application.
        /// </summary>
        [Input("applicationName", required: true)]
        public string ApplicationName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetApplicationArgs()
        {
        }
        public static new GetApplicationArgs Empty => new GetApplicationArgs();
    }

    public sealed class GetApplicationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the managed application.
        /// </summary>
        [Input("applicationName", required: true)]
        public Input<string> ApplicationName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetApplicationInvokeArgs()
        {
        }
        public static new GetApplicationInvokeArgs Empty => new GetApplicationInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationResult
    {
        /// <summary>
        /// The fully qualified path of managed application definition Id.
        /// </summary>
        public readonly string? ApplicationDefinitionId;
        /// <summary>
        /// The collection of managed application artifacts.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationArtifactResponse> Artifacts;
        /// <summary>
        /// The  read-only authorizations property that is retrieved from the application package.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationAuthorizationResponse> Authorizations;
        /// <summary>
        /// The managed application billing details.
        /// </summary>
        public readonly Outputs.ApplicationBillingDetailsDefinitionResponse BillingDetails;
        /// <summary>
        /// The client entity that created the JIT request.
        /// </summary>
        public readonly Outputs.ApplicationClientDetailsResponse CreatedBy;
        /// <summary>
        /// The read-only customer support property that is retrieved from the application package.
        /// </summary>
        public readonly Outputs.ApplicationPackageContactResponse CustomerSupport;
        /// <summary>
        /// Resource ID
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The identity of the resource.
        /// </summary>
        public readonly Outputs.IdentityResponse? Identity;
        /// <summary>
        /// The managed application Jit access policy.
        /// </summary>
        public readonly Outputs.ApplicationJitAccessPolicyResponse? JitAccessPolicy;
        /// <summary>
        /// The kind of the managed application. Allowed values are MarketPlace and ServiceCatalog.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// ID of the resource that manages this resource.
        /// </summary>
        public readonly string? ManagedBy;
        /// <summary>
        /// The managed resource group Id.
        /// </summary>
        public readonly string? ManagedResourceGroupId;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Name and value pairs that define the managed application outputs.
        /// </summary>
        public readonly object Outputs;
        /// <summary>
        /// Name and value pairs that define the managed application parameters. It can be a JObject or a well formed JSON string.
        /// </summary>
        public readonly object? Parameters;
        /// <summary>
        /// The plan information.
        /// </summary>
        public readonly Outputs.PlanResponse? Plan;
        /// <summary>
        /// The managed application provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The publisher package Id.
        /// </summary>
        public readonly string? PublisherPackageId;
        /// <summary>
        /// The publisher tenant Id.
        /// </summary>
        public readonly string PublisherTenantId;
        /// <summary>
        /// The SKU of the resource.
        /// </summary>
        public readonly Outputs.SkuResponse? Sku;
        /// <summary>
        /// The read-only support URLs property that is retrieved from the application package.
        /// </summary>
        public readonly Outputs.ApplicationPackageSupportUrlsResponse SupportUrls;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The client entity that last updated the JIT request.
        /// </summary>
        public readonly Outputs.ApplicationClientDetailsResponse UpdatedBy;

        [OutputConstructor]
        private GetApplicationResult(
            string? applicationDefinitionId,

            ImmutableArray<Outputs.ApplicationArtifactResponse> artifacts,

            ImmutableArray<Outputs.ApplicationAuthorizationResponse> authorizations,

            Outputs.ApplicationBillingDetailsDefinitionResponse billingDetails,

            Outputs.ApplicationClientDetailsResponse createdBy,

            Outputs.ApplicationPackageContactResponse customerSupport,

            string id,

            Outputs.IdentityResponse? identity,

            Outputs.ApplicationJitAccessPolicyResponse? jitAccessPolicy,

            string kind,

            string? location,

            string? managedBy,

            string? managedResourceGroupId,

            string name,

            object outputs,

            object? parameters,

            Outputs.PlanResponse? plan,

            string provisioningState,

            string? publisherPackageId,

            string publisherTenantId,

            Outputs.SkuResponse? sku,

            Outputs.ApplicationPackageSupportUrlsResponse supportUrls,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.ApplicationClientDetailsResponse updatedBy)
        {
            ApplicationDefinitionId = applicationDefinitionId;
            Artifacts = artifacts;
            Authorizations = authorizations;
            BillingDetails = billingDetails;
            CreatedBy = createdBy;
            CustomerSupport = customerSupport;
            Id = id;
            Identity = identity;
            JitAccessPolicy = jitAccessPolicy;
            Kind = kind;
            Location = location;
            ManagedBy = managedBy;
            ManagedResourceGroupId = managedResourceGroupId;
            Name = name;
            Outputs = outputs;
            Parameters = parameters;
            Plan = plan;
            ProvisioningState = provisioningState;
            PublisherPackageId = publisherPackageId;
            PublisherTenantId = publisherTenantId;
            Sku = sku;
            SupportUrls = supportUrls;
            Tags = tags;
            Type = type;
            UpdatedBy = updatedBy;
        }
    }
}
