// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.V20170901Preview.Inputs
{

    /// <summary>
    /// Self-hosted integration runtime.
    /// </summary>
    public sealed class SelfHostedIntegrationRuntimeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Integration runtime description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The base definition of a secret type.
        /// </summary>
        [Input("linkedInfo")]
        public InputUnion<Inputs.LinkedIntegrationRuntimeKeyArgs, Inputs.LinkedIntegrationRuntimeRbacArgs>? LinkedInfo { get; set; }

        /// <summary>
        /// The type of integration runtime.
        /// Expected value is 'SelfHosted'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public SelfHostedIntegrationRuntimeArgs()
        {
        }
        public static new SelfHostedIntegrationRuntimeArgs Empty => new SelfHostedIntegrationRuntimeArgs();
    }
}
