// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DeploymentManager.V20180901Preview.Inputs
{

    /// <summary>
    /// The parameters for the wait step.
    /// </summary>
    public sealed class WaitStepAttributesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The duration in ISO 8601 format of how long the wait should be.
        /// </summary>
        [Input("duration", required: true)]
        public Input<string> Duration { get; set; } = null!;

        public WaitStepAttributesArgs()
        {
        }
        public static new WaitStepAttributesArgs Empty => new WaitStepAttributesArgs();
    }
}
