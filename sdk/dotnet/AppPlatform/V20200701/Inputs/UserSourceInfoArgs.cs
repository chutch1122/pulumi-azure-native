// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform.V20200701.Inputs
{

    /// <summary>
    /// Source information for a deployment
    /// </summary>
    public sealed class UserSourceInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Selector for the artifact to be used for the deployment for multi-module projects. This should be
        /// the relative path to the target module/project.
        /// </summary>
        [Input("artifactSelector")]
        public Input<string>? ArtifactSelector { get; set; }

        /// <summary>
        /// Relative path of the storage which stores the source
        /// </summary>
        [Input("relativePath")]
        public Input<string>? RelativePath { get; set; }

        /// <summary>
        /// Type of the source uploaded
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNative.AppPlatform.V20200701.UserSourceType>? Type { get; set; }

        /// <summary>
        /// Version of the source
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public UserSourceInfoArgs()
        {
        }
        public static new UserSourceInfoArgs Empty => new UserSourceInfoArgs();
    }
}
