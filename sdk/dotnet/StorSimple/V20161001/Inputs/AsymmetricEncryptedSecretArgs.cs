// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.StorSimple.V20161001.Inputs
{

    /// <summary>
    /// This class can be used as the Type for any secret entity represented as Password, CertThumbprint, Algorithm. This class is intended to be used when the secret is encrypted with an asymmetric key pair. The encryptionAlgorithm field is mainly for future usage to potentially allow different entities encrypted using different algorithms.
    /// </summary>
    public sealed class AsymmetricEncryptedSecretArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Algorithm used to encrypt "Value"
        /// </summary>
        [Input("encryptionAlgorithm", required: true)]
        public Input<Pulumi.AzureNative.StorSimple.V20161001.EncryptionAlgorithm> EncryptionAlgorithm { get; set; } = null!;

        /// <summary>
        /// Thumbprint certificate that was used to encrypt "Value"
        /// </summary>
        [Input("encryptionCertificateThumbprint")]
        public Input<string>? EncryptionCertificateThumbprint { get; set; }

        /// <summary>
        /// The value of the secret itself. If the secret is in plaintext then EncryptionAlgorithm will be none and EncryptionCertThumbprint will be null.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public AsymmetricEncryptedSecretArgs()
        {
        }
        public static new AsymmetricEncryptedSecretArgs Empty => new AsymmetricEncryptedSecretArgs();
    }
}
