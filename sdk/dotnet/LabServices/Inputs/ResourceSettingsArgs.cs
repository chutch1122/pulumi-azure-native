// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.LabServices.Inputs
{

    /// <summary>
    /// Represents resource specific settings
    /// </summary>
    public sealed class ResourceSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource id of the gallery image used for creating the virtual machine
        /// </summary>
        [Input("galleryImageResourceId")]
        public Input<string>? GalleryImageResourceId { get; set; }

        /// <summary>
        /// Details specific to Reference Vm
        /// </summary>
        [Input("referenceVm", required: true)]
        public Input<Inputs.ReferenceVmArgs> ReferenceVm { get; set; } = null!;

        /// <summary>
        /// The size of the virtual machine
        /// </summary>
        [Input("size")]
        public InputUnion<string, Pulumi.AzureNative.LabServices.ManagedLabVmSize>? Size { get; set; }

        public ResourceSettingsArgs()
        {
        }
        public static new ResourceSettingsArgs Empty => new ResourceSettingsArgs();
    }
}
