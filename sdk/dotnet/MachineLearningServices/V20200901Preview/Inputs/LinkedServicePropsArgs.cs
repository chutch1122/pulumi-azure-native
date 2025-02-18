// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20200901Preview.Inputs
{

    /// <summary>
    /// LinkedService specific properties.
    /// </summary>
    public sealed class LinkedServicePropsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The creation time of the linked service.
        /// </summary>
        [Input("createdTime")]
        public Input<string>? CreatedTime { get; set; }

        /// <summary>
        /// Type of the link target.
        /// </summary>
        [Input("linkType")]
        public Input<Pulumi.AzureNative.MachineLearningServices.V20200901Preview.LinkedServiceLinkType>? LinkType { get; set; }

        /// <summary>
        /// ResourceId of the link target of the linked service.
        /// </summary>
        [Input("linkedServiceResourceId", required: true)]
        public Input<string> LinkedServiceResourceId { get; set; } = null!;

        /// <summary>
        /// The last modified time of the linked service.
        /// </summary>
        [Input("modifiedTime")]
        public Input<string>? ModifiedTime { get; set; }

        public LinkedServicePropsArgs()
        {
        }
        public static new LinkedServicePropsArgs Empty => new LinkedServicePropsArgs();
    }
}
