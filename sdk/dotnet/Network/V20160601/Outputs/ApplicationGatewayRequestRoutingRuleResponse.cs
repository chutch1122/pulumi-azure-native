// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20160601.Outputs
{

    /// <summary>
    /// Request routing rule of application gateway
    /// </summary>
    [OutputType]
    public sealed class ApplicationGatewayRequestRoutingRuleResponse
    {
        /// <summary>
        /// Backend address pool resource of application gateway 
        /// </summary>
        public readonly Outputs.SubResourceResponse? BackendAddressPool;
        /// <summary>
        /// Frontend port resource of application gateway 
        /// </summary>
        public readonly Outputs.SubResourceResponse? BackendHttpSettings;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Http listener resource of application gateway 
        /// </summary>
        public readonly Outputs.SubResourceResponse? HttpListener;
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Name of the resource that is unique within a resource group. This name can be used to access the resource
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Provisioning state of the request routing rule resource Updating/Deleting/Failed
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Rule type
        /// </summary>
        public readonly string? RuleType;
        /// <summary>
        /// Url path map resource of application gateway 
        /// </summary>
        public readonly Outputs.SubResourceResponse? UrlPathMap;

        [OutputConstructor]
        private ApplicationGatewayRequestRoutingRuleResponse(
            Outputs.SubResourceResponse? backendAddressPool,

            Outputs.SubResourceResponse? backendHttpSettings,

            string? etag,

            Outputs.SubResourceResponse? httpListener,

            string? id,

            string? name,

            string? provisioningState,

            string? ruleType,

            Outputs.SubResourceResponse? urlPathMap)
        {
            BackendAddressPool = backendAddressPool;
            BackendHttpSettings = backendHttpSettings;
            Etag = etag;
            HttpListener = httpListener;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            RuleType = ruleType;
            UrlPathMap = urlPathMap;
        }
    }
}
