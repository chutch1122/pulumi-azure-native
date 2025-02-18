// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media.V20180601Preview.Inputs
{

    /// <summary>
    /// The Live Event preview.
    /// </summary>
    public sealed class LiveEventPreviewArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access control for LiveEvent preview.
        /// </summary>
        [Input("accessControl")]
        public Input<Inputs.LiveEventPreviewAccessControlArgs>? AccessControl { get; set; }

        /// <summary>
        /// An Alternative Media Identifier associated with the preview url.  This identifier can be used to distinguish the preview of different live events for authorization purposes in the CustomLicenseAcquisitionUrlTemplate or the CustomKeyAcquisitionUrlTemplate of the StreamingPolicy specified in the StreamingPolicyName field.
        /// </summary>
        [Input("alternativeMediaId")]
        public Input<string>? AlternativeMediaId { get; set; }

        [Input("endpoints")]
        private InputList<Inputs.LiveEventEndpointArgs>? _endpoints;

        /// <summary>
        /// The endpoints for preview.
        /// </summary>
        public InputList<Inputs.LiveEventEndpointArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.LiveEventEndpointArgs>());
            set => _endpoints = value;
        }

        /// <summary>
        /// The preview locator Guid.
        /// </summary>
        [Input("previewLocator")]
        public Input<string>? PreviewLocator { get; set; }

        /// <summary>
        /// The name of streaming policy used for LiveEvent preview
        /// </summary>
        [Input("streamingPolicyName")]
        public Input<string>? StreamingPolicyName { get; set; }

        public LiveEventPreviewArgs()
        {
        }
        public static new LiveEventPreviewArgs Empty => new LiveEventPreviewArgs();
    }
}
