// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NotificationHubs.V20140901.Outputs
{

    /// <summary>
    /// Description of a NotificationHub ApnsCredential.
    /// </summary>
    [OutputType]
    public sealed class ApnsCredentialPropertiesResponse
    {
        /// <summary>
        /// Gets or sets the APNS certificate.
        /// </summary>
        public readonly string? ApnsCertificate;
        /// <summary>
        /// Gets or sets the certificate key.
        /// </summary>
        public readonly string? CertificateKey;
        /// <summary>
        /// Gets or sets the endpoint of this credential.
        /// </summary>
        public readonly string? Endpoint;
        /// <summary>
        /// Gets or sets the Apns certificate Thumbprint
        /// </summary>
        public readonly string? Thumbprint;

        [OutputConstructor]
        private ApnsCredentialPropertiesResponse(
            string? apnsCertificate,

            string? certificateKey,

            string? endpoint,

            string? thumbprint)
        {
            ApnsCertificate = apnsCertificate;
            CertificateKey = certificateKey;
            Endpoint = endpoint;
            Thumbprint = thumbprint;
        }
    }
}
