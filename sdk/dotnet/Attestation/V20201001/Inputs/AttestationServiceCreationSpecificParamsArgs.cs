// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Attestation.V20201001.Inputs
{

    /// <summary>
    /// Client supplied parameters used to create a new attestation provider.
    /// </summary>
    public sealed class AttestationServiceCreationSpecificParamsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// JSON Web Key Set defining a set of X.509 Certificates that will represent the parent certificate for the signing certificate used for policy operations
        /// </summary>
        [Input("policySigningCertificates")]
        public Input<Inputs.JSONWebKeySetArgs>? PolicySigningCertificates { get; set; }

        public AttestationServiceCreationSpecificParamsArgs()
        {
        }
        public static new AttestationServiceCreationSpecificParamsArgs Empty => new AttestationServiceCreationSpecificParamsArgs();
    }
}
