// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20160601.Inputs
{

    /// <summary>
    /// Authentication certificates of application gateway
    /// </summary>
    public sealed class ApplicationGatewayAuthenticationCertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate public data 
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the resource that is unique within a resource group. This name can be used to access the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Provisioning state of the authentication certificate resource Updating/Deleting/Failed
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        public ApplicationGatewayAuthenticationCertificateArgs()
        {
        }
        public static new ApplicationGatewayAuthenticationCertificateArgs Empty => new ApplicationGatewayAuthenticationCertificateArgs();
    }
}
