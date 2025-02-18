// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.V20220601Preview.Outputs
{

    /// <summary>
    /// Container App Ingress configuration.
    /// </summary>
    [OutputType]
    public sealed class IngressResponse
    {
        /// <summary>
        /// Bool indicating if HTTP connections to is allowed. If set to false HTTP connections are automatically redirected to HTTPS connections
        /// </summary>
        public readonly bool? AllowInsecure;
        /// <summary>
        /// custom domain bindings for Container Apps' hostnames.
        /// </summary>
        public readonly ImmutableArray<Outputs.CustomDomainResponse> CustomDomains;
        /// <summary>
        /// Exposed Port in containers for TCP traffic from ingress
        /// </summary>
        public readonly int? ExposedPort;
        /// <summary>
        /// Bool indicating if app exposes an external http endpoint
        /// </summary>
        public readonly bool? External;
        /// <summary>
        /// Hostname.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// Rules to restrict incoming IP address.
        /// </summary>
        public readonly ImmutableArray<Outputs.IpSecurityRestrictionRuleResponse> IpSecurityRestrictions;
        /// <summary>
        /// Target Port in containers for traffic from ingress
        /// </summary>
        public readonly int? TargetPort;
        /// <summary>
        /// Traffic weights for app's revisions
        /// </summary>
        public readonly ImmutableArray<Outputs.TrafficWeightResponse> Traffic;
        /// <summary>
        /// Ingress transport protocol
        /// </summary>
        public readonly string? Transport;

        [OutputConstructor]
        private IngressResponse(
            bool? allowInsecure,

            ImmutableArray<Outputs.CustomDomainResponse> customDomains,

            int? exposedPort,

            bool? external,

            string fqdn,

            ImmutableArray<Outputs.IpSecurityRestrictionRuleResponse> ipSecurityRestrictions,

            int? targetPort,

            ImmutableArray<Outputs.TrafficWeightResponse> traffic,

            string? transport)
        {
            AllowInsecure = allowInsecure;
            CustomDomains = customDomains;
            ExposedPort = exposedPort;
            External = external;
            Fqdn = fqdn;
            IpSecurityRestrictions = ipSecurityRestrictions;
            TargetPort = targetPort;
            Traffic = traffic;
            Transport = transport;
        }
    }
}
