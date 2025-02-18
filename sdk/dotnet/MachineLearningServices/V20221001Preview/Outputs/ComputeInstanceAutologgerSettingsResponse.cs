// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20221001Preview.Outputs
{

    /// <summary>
    /// Specifies settings for autologger.
    /// </summary>
    [OutputType]
    public sealed class ComputeInstanceAutologgerSettingsResponse
    {
        /// <summary>
        /// Indicates whether mlflow autologger is enabled for notebooks.
        /// </summary>
        public readonly string? MlflowAutologger;

        [OutputConstructor]
        private ComputeInstanceAutologgerSettingsResponse(string? mlflowAutologger)
        {
            MlflowAutologger = mlflowAutologger;
        }
    }
}
