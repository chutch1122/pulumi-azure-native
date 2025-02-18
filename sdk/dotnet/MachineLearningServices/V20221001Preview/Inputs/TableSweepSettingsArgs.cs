// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20221001Preview.Inputs
{

    public sealed class TableSweepSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of early termination policy for the sweeping job.
        /// </summary>
        [Input("earlyTermination")]
        public object? EarlyTermination { get; set; }

        /// <summary>
        /// [Required] Type of sampling algorithm.
        /// </summary>
        [Input("samplingAlgorithm", required: true)]
        public InputUnion<string, Pulumi.AzureNative.MachineLearningServices.V20221001Preview.SamplingAlgorithmType> SamplingAlgorithm { get; set; } = null!;

        public TableSweepSettingsArgs()
        {
        }
        public static new TableSweepSettingsArgs Empty => new TableSweepSettingsArgs();
    }
}
