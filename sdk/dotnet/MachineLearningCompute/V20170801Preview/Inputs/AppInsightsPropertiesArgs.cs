// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningCompute.V20170801Preview.Inputs
{

    /// <summary>
    /// Properties of App Insights.
    /// </summary>
    public sealed class AppInsightsPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARM resource ID of the App Insights.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        public AppInsightsPropertiesArgs()
        {
        }
        public static new AppInsightsPropertiesArgs Empty => new AppInsightsPropertiesArgs();
    }
}
