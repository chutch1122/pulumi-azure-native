// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevCenter.Outputs
{

    /// <summary>
    /// Image reference information
    /// </summary>
    [OutputType]
    public sealed class ImageReferenceResponse
    {
        /// <summary>
        /// The actual version of the image after use. When id references a gallery image latest version, this will indicate the actual version in use.
        /// </summary>
        public readonly string ExactVersion;
        /// <summary>
        /// Image ID, or Image version ID. When Image ID is provided, its latest version will be used.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The image offer.
        /// </summary>
        public readonly string? Offer;
        /// <summary>
        /// The image publisher.
        /// </summary>
        public readonly string? Publisher;
        /// <summary>
        /// The image sku.
        /// </summary>
        public readonly string? Sku;

        [OutputConstructor]
        private ImageReferenceResponse(
            string exactVersion,

            string? id,

            string? offer,

            string? publisher,

            string? sku)
        {
            ExactVersion = exactVersion;
            Id = id;
            Offer = offer;
            Publisher = publisher;
            Sku = sku;
        }
    }
}
