// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ContainerRegistry.V20190501Preview.Outputs
{

    /// <summary>
    /// The properties of the credentials that can be used for authenticating the token.
    /// </summary>
    [OutputType]
    public sealed class TokenCredentialsPropertiesResponse
    {
        /// <summary>
        /// The Active Directory Object that will be used for authenticating the token of a container registry.
        /// </summary>
        public readonly Outputs.ActiveDirectoryObjectResponse? ActiveDirectoryObject;
        public readonly ImmutableArray<Outputs.TokenCertificateResponse> Certificates;
        public readonly ImmutableArray<Outputs.TokenPasswordResponse> Passwords;

        [OutputConstructor]
        private TokenCredentialsPropertiesResponse(
            Outputs.ActiveDirectoryObjectResponse? activeDirectoryObject,

            ImmutableArray<Outputs.TokenCertificateResponse> certificates,

            ImmutableArray<Outputs.TokenPasswordResponse> passwords)
        {
            ActiveDirectoryObject = activeDirectoryObject;
            Certificates = certificates;
            Passwords = passwords;
        }
    }
}
