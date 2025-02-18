// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DeploymentManager.Inputs
{

    /// <summary>
    /// The properties that make up a REST request
    /// </summary>
    public sealed class RestRequestArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authentication information required in the request to the health provider.
        /// </summary>
        [Input("authentication", required: true)]
        public InputUnion<Inputs.ApiKeyAuthenticationArgs, Inputs.RolloutIdentityAuthenticationArgs> Authentication { get; set; } = null!;

        /// <summary>
        /// The HTTP method to use for the request.
        /// </summary>
        [Input("method", required: true)]
        public Input<Pulumi.AzureNative.DeploymentManager.RestRequestMethod> Method { get; set; } = null!;

        /// <summary>
        /// The HTTP URI to use for the request.
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public RestRequestArgs()
        {
        }
        public static new RestRequestArgs Empty => new RestRequestArgs();
    }
}
