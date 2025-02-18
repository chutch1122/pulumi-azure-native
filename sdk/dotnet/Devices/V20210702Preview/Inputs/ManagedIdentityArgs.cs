// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Devices.V20210702Preview.Inputs
{

    /// <summary>
    /// The properties of the Managed identity.
    /// </summary>
    public sealed class ManagedIdentityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The user assigned identity.
        /// </summary>
        [Input("userAssignedIdentity")]
        public Input<string>? UserAssignedIdentity { get; set; }

        public ManagedIdentityArgs()
        {
        }
        public static new ManagedIdentityArgs Empty => new ManagedIdentityArgs();
    }
}
