// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NotificationHubs.V20140901.Inputs
{

    /// <summary>
    /// Description of a NotificationHub GcmCredential.
    /// </summary>
    public sealed class GcmCredentialPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the GCM endpoint.
        /// </summary>
        [Input("gcmEndpoint")]
        public Input<string>? GcmEndpoint { get; set; }

        /// <summary>
        /// Gets or sets the Google API key.
        /// </summary>
        [Input("googleApiKey")]
        public Input<string>? GoogleApiKey { get; set; }

        public GcmCredentialPropertiesArgs()
        {
        }
        public static new GcmCredentialPropertiesArgs Empty => new GcmCredentialPropertiesArgs();
    }
}
