// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.V20200501.Inputs
{

    public sealed class TaskSchedulingPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("nodeFillType", required: true)]
        public Input<Pulumi.AzureNative.Batch.V20200501.ComputeNodeFillType> NodeFillType { get; set; } = null!;

        public TaskSchedulingPolicyArgs()
        {
        }
        public static new TaskSchedulingPolicyArgs Empty => new TaskSchedulingPolicyArgs();
    }
}
