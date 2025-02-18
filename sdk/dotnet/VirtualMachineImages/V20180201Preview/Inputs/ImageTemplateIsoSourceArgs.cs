// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.VirtualMachineImages.V20180201Preview.Inputs
{

    /// <summary>
    /// Describes an image source that is an installation ISO. Currently only supports Red Hat Enterprise Linux 7.2-7.5 ISO's.
    /// </summary>
    public sealed class ImageTemplateIsoSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SHA256 Checksum of the ISO image.
        /// </summary>
        [Input("sha256Checksum", required: true)]
        public Input<string> Sha256Checksum { get; set; } = null!;

        /// <summary>
        /// URL to get the ISO image. This URL has to be accessible to the resource provider at the time of the imageTemplate creation.
        /// </summary>
        [Input("sourceURI", required: true)]
        public Input<string> SourceURI { get; set; } = null!;

        /// <summary>
        /// Specifies the type of source image you want to start with.
        /// Expected value is 'ISO'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ImageTemplateIsoSourceArgs()
        {
        }
        public static new ImageTemplateIsoSourceArgs Empty => new ImageTemplateIsoSourceArgs();
    }
}
