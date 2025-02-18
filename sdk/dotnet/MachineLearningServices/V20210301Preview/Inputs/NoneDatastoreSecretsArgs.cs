// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20210301Preview.Inputs
{

    /// <summary>
    /// Empty/none datastore secret.
    /// </summary>
    public sealed class NoneDatastoreSecretsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enum to determine the datastore secrets type.
        /// Expected value is 'None'.
        /// </summary>
        [Input("secretsType", required: true)]
        public Input<string> SecretsType { get; set; } = null!;

        public NoneDatastoreSecretsArgs()
        {
        }
        public static new NoneDatastoreSecretsArgs Empty => new NoneDatastoreSecretsArgs();
    }
}
