// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerService.V20170131.Inputs
{

    /// <summary>
    /// Properties to configure a custom container service cluster.
    /// </summary>
    public sealed class ContainerServiceCustomProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the custom orchestrator to use.
        /// </summary>
        [Input("orchestrator", required: true)]
        public Input<string> Orchestrator { get; set; } = null!;

        public ContainerServiceCustomProfileArgs()
        {
        }
        public static new ContainerServiceCustomProfileArgs Empty => new ContainerServiceCustomProfileArgs();
    }
}
