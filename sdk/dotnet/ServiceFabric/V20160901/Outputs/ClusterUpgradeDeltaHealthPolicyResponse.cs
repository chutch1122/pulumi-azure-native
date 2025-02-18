// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ServiceFabric.V20160901.Outputs
{

    /// <summary>
    /// Delta health policy for the cluster
    /// </summary>
    [OutputType]
    public sealed class ClusterUpgradeDeltaHealthPolicyResponse
    {
        /// <summary>
        /// Additional unhealthy applications percentage
        /// </summary>
        public readonly int MaxPercentDeltaUnhealthyApplications;
        /// <summary>
        /// Additional unhealthy nodes percentage
        /// </summary>
        public readonly int MaxPercentDeltaUnhealthyNodes;
        /// <summary>
        /// Additional unhealthy nodes percentage per upgrade domain 
        /// </summary>
        public readonly int MaxPercentUpgradeDomainDeltaUnhealthyNodes;

        [OutputConstructor]
        private ClusterUpgradeDeltaHealthPolicyResponse(
            int maxPercentDeltaUnhealthyApplications,

            int maxPercentDeltaUnhealthyNodes,

            int maxPercentUpgradeDomainDeltaUnhealthyNodes)
        {
            MaxPercentDeltaUnhealthyApplications = maxPercentDeltaUnhealthyApplications;
            MaxPercentDeltaUnhealthyNodes = maxPercentDeltaUnhealthyNodes;
            MaxPercentUpgradeDomainDeltaUnhealthyNodes = maxPercentUpgradeDomainDeltaUnhealthyNodes;
        }
    }
}
