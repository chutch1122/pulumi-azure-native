// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Logic.V20160601.Inputs
{

    /// <summary>
    /// The AS2 agreement security settings.
    /// </summary>
    public sealed class AS2SecuritySettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The value indicating whether to enable NRR for inbound decoded messages.
        /// </summary>
        [Input("enableNrrForInboundDecodedMessages", required: true)]
        public Input<bool> EnableNrrForInboundDecodedMessages { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to enable NRR for inbound encoded messages.
        /// </summary>
        [Input("enableNrrForInboundEncodedMessages", required: true)]
        public Input<bool> EnableNrrForInboundEncodedMessages { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to enable NRR for inbound MDN.
        /// </summary>
        [Input("enableNrrForInboundMdn", required: true)]
        public Input<bool> EnableNrrForInboundMdn { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to enable NRR for outbound decoded messages.
        /// </summary>
        [Input("enableNrrForOutboundDecodedMessages", required: true)]
        public Input<bool> EnableNrrForOutboundDecodedMessages { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to enable NRR for outbound encoded messages.
        /// </summary>
        [Input("enableNrrForOutboundEncodedMessages", required: true)]
        public Input<bool> EnableNrrForOutboundEncodedMessages { get; set; } = null!;

        /// <summary>
        /// The value indicating whether to enable NRR for outbound MDN.
        /// </summary>
        [Input("enableNrrForOutboundMdn", required: true)]
        public Input<bool> EnableNrrForOutboundMdn { get; set; } = null!;

        /// <summary>
        /// The name of the encryption certificate.
        /// </summary>
        [Input("encryptionCertificateName")]
        public Input<string>? EncryptionCertificateName { get; set; }

        /// <summary>
        /// The value indicating whether to send or request a MDN.
        /// </summary>
        [Input("overrideGroupSigningCertificate", required: true)]
        public Input<bool> OverrideGroupSigningCertificate { get; set; } = null!;

        /// <summary>
        /// The Sha2 algorithm format. Valid values are Sha2, ShaHashSize, ShaHyphenHashSize, Sha2UnderscoreHashSize.
        /// </summary>
        [Input("sha2AlgorithmFormat")]
        public Input<string>? Sha2AlgorithmFormat { get; set; }

        /// <summary>
        /// The name of the signing certificate.
        /// </summary>
        [Input("signingCertificateName")]
        public Input<string>? SigningCertificateName { get; set; }

        public AS2SecuritySettingsArgs()
        {
        }
        public static new AS2SecuritySettingsArgs Empty => new AS2SecuritySettingsArgs();
    }
}
