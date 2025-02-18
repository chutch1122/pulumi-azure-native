// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql.V20180601Preview.Inputs
{

    /// <summary>
    /// Azure Active Directory identity configuration for a resource.
    /// </summary>
    public sealed class ResourceIdentityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identity type. Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active Directory principal for the resource.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNative.Sql.V20180601Preview.IdentityType>? Type { get; set; }

        public ResourceIdentityArgs()
        {
        }
        public static new ResourceIdentityArgs Empty => new ResourceIdentityArgs();
    }
}
