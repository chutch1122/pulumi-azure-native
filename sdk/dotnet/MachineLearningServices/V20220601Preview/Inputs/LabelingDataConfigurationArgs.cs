// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20220601Preview.Inputs
{

    /// <summary>
    /// Labeling data configuration definition
    /// </summary>
    public sealed class LabelingDataConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource Id of the data asset to perform labeling.
        /// </summary>
        [Input("dataId")]
        public Input<string>? DataId { get; set; }

        /// <summary>
        /// Indicates whether to enable incremental data refresh.
        /// </summary>
        [Input("incrementalDataRefreshEnabled")]
        public Input<bool>? IncrementalDataRefreshEnabled { get; set; }

        public LabelingDataConfigurationArgs()
        {
            IncrementalDataRefreshEnabled = false;
        }
        public static new LabelingDataConfigurationArgs Empty => new LabelingDataConfigurationArgs();
    }
}
