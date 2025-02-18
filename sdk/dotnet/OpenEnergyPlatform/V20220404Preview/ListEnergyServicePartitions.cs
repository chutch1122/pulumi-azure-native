// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.OpenEnergyPlatform.V20220404Preview
{
    public static class ListEnergyServicePartitions
    {
        /// <summary>
        /// List of data partitions.
        /// </summary>
        public static Task<ListEnergyServicePartitionsResult> InvokeAsync(ListEnergyServicePartitionsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<ListEnergyServicePartitionsResult>("azure-native:openenergyplatform/v20220404preview:listEnergyServicePartitions", args ?? new ListEnergyServicePartitionsArgs(), options.WithDefaults());

        /// <summary>
        /// List of data partitions.
        /// </summary>
        public static Output<ListEnergyServicePartitionsResult> Invoke(ListEnergyServicePartitionsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<ListEnergyServicePartitionsResult>("azure-native:openenergyplatform/v20220404preview:listEnergyServicePartitions", args ?? new ListEnergyServicePartitionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class ListEnergyServicePartitionsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public ListEnergyServicePartitionsArgs()
        {
        }
        public static new ListEnergyServicePartitionsArgs Empty => new ListEnergyServicePartitionsArgs();
    }

    public sealed class ListEnergyServicePartitionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public ListEnergyServicePartitionsInvokeArgs()
        {
        }
        public static new ListEnergyServicePartitionsInvokeArgs Empty => new ListEnergyServicePartitionsInvokeArgs();
    }


    [OutputType]
    public sealed class ListEnergyServicePartitionsResult
    {
        /// <summary>
        /// List of data partitions along with their properties in a given OEP resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.DataPartitionPropertiesResponse> DataPartitionInfo;

        [OutputConstructor]
        private ListEnergyServicePartitionsResult(ImmutableArray<Outputs.DataPartitionPropertiesResponse> dataPartitionInfo)
        {
            DataPartitionInfo = dataPartitionInfo;
        }
    }
}
