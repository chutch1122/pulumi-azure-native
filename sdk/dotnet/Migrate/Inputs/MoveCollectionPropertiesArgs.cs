// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Inputs
{

    /// <summary>
    /// Defines the move collection properties.
    /// </summary>
    public sealed class MoveCollectionPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the source region.
        /// </summary>
        [Input("sourceRegion", required: true)]
        public Input<string> SourceRegion { get; set; } = null!;

        /// <summary>
        /// Gets or sets the target region.
        /// </summary>
        [Input("targetRegion", required: true)]
        public Input<string> TargetRegion { get; set; } = null!;

        public MoveCollectionPropertiesArgs()
        {
        }
        public static new MoveCollectionPropertiesArgs Empty => new MoveCollectionPropertiesArgs();
    }
}
