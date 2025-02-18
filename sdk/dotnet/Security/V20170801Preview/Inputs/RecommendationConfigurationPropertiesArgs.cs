// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.V20170801Preview.Inputs
{

    /// <summary>
    /// Recommendation configuration
    /// </summary>
    public sealed class RecommendationConfigurationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The recommendation type.
        /// </summary>
        [Input("recommendationType", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Security.V20170801Preview.RecommendationType> RecommendationType { get; set; } = null!;

        /// <summary>
        /// Recommendation status. The recommendation is not generated when the status is disabled
        /// </summary>
        [Input("status", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Security.V20170801Preview.RecommendationConfigStatus> Status { get; set; } = null!;

        public RecommendationConfigurationPropertiesArgs()
        {
            Status = "Enabled";
        }
        public static new RecommendationConfigurationPropertiesArgs Empty => new RecommendationConfigurationPropertiesArgs();
    }
}
