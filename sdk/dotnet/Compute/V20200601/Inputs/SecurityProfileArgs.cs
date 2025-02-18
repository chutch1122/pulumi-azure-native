// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.V20200601.Inputs
{

    /// <summary>
    /// Specifies the Security profile settings for the virtual machine or virtual machine scale set.
    /// </summary>
    public sealed class SecurityProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This property can be used by user in the request to enable or disable the Host Encryption for the virtual machine or virtual machine scale set. This will enable the encryption for all the disks including Resource/Temp disk at host itself. &lt;br&gt;&lt;br&gt; Default: The Encryption at host will be disabled unless this property is set to true for the resource.
        /// </summary>
        [Input("encryptionAtHost")]
        public Input<bool>? EncryptionAtHost { get; set; }

        public SecurityProfileArgs()
        {
        }
        public static new SecurityProfileArgs Empty => new SecurityProfileArgs();
    }
}
