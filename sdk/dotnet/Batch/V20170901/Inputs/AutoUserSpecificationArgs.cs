// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.V20170901.Inputs
{

    public sealed class AutoUserSpecificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// nonAdmin - The auto user is a standard user without elevated access. admin - The auto user is a user with elevated access and operates with full Administrator permissions. The default value is nonAdmin.
        /// </summary>
        [Input("elevationLevel")]
        public Input<Pulumi.AzureNative.Batch.V20170901.ElevationLevel>? ElevationLevel { get; set; }

        /// <summary>
        /// pool - specifies that the task runs as the common auto user account which is created on every node in a pool. task - specifies that the service should create a new user for the task. The default value is task.
        /// </summary>
        [Input("scope")]
        public Input<Pulumi.AzureNative.Batch.V20170901.AutoUserScope>? Scope { get; set; }

        public AutoUserSpecificationArgs()
        {
        }
        public static new AutoUserSpecificationArgs Empty => new AutoUserSpecificationArgs();
    }
}
