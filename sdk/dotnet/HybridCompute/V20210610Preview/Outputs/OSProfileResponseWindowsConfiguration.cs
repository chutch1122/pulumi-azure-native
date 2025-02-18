// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridCompute.V20210610Preview.Outputs
{

    /// <summary>
    /// Specifies the windows configuration for update management.
    /// </summary>
    [OutputType]
    public sealed class OSProfileResponseWindowsConfiguration
    {
        /// <summary>
        /// Specifies the assessment mode.
        /// </summary>
        public readonly string? AssessmentMode;

        [OutputConstructor]
        private OSProfileResponseWindowsConfiguration(string? assessmentMode)
        {
            AssessmentMode = assessmentMode;
        }
    }
}
