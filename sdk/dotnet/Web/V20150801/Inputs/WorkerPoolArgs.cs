// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.V20150801.Inputs
{

    /// <summary>
    /// Worker pool of a hostingEnvironment (App Service Environment)
    /// </summary>
    public sealed class WorkerPoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Shared or dedicated web app hosting
        /// </summary>
        [Input("computeMode")]
        public Input<Pulumi.AzureNative.Web.V20150801.ComputeModeOptions>? ComputeMode { get; set; }

        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("instanceNames")]
        private InputList<string>? _instanceNames;

        /// <summary>
        /// Names of all instances in the worker pool (read only)
        /// </summary>
        public InputList<string> InstanceNames
        {
            get => _instanceNames ?? (_instanceNames = new InputList<string>());
            set => _instanceNames = value;
        }

        /// <summary>
        /// Kind of resource
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Resource Location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Resource Name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Describes a sku for a scalable resource
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuDescriptionArgs>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Resource type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Number of instances in the worker pool
        /// </summary>
        [Input("workerCount")]
        public Input<int>? WorkerCount { get; set; }

        /// <summary>
        /// VM size of the worker pool instances
        /// </summary>
        [Input("workerSize")]
        public Input<string>? WorkerSize { get; set; }

        /// <summary>
        /// Worker size id for referencing this worker pool
        /// </summary>
        [Input("workerSizeId")]
        public Input<int>? WorkerSizeId { get; set; }

        public WorkerPoolArgs()
        {
        }
        public static new WorkerPoolArgs Empty => new WorkerPoolArgs();
    }
}
