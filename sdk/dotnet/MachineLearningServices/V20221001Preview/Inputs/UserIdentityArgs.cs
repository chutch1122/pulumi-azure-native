// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20221001Preview.Inputs
{

    /// <summary>
    /// User identity configuration.
    /// </summary>
    public sealed class UserIdentityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enum to determine identity framework.
        /// Expected value is 'UserIdentity'.
        /// </summary>
        [Input("identityType", required: true)]
        public Input<string> IdentityType { get; set; } = null!;

        public UserIdentityArgs()
        {
        }
        public static new UserIdentityArgs Empty => new UserIdentityArgs();
    }
}
