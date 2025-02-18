// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.V20170901.Inputs
{

    public sealed class ResourceFileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This URL must be readable using anonymous access; that is, the Batch service does not present any credentials when downloading the blob. There are two ways to get such a URL for a blob in Azure storage: include a Shared Access Signature (SAS) granting read permissions on the blob, or set the ACL for the blob or its container to allow public access.
        /// </summary>
        [Input("blobSource", required: true)]
        public Input<string> BlobSource { get; set; } = null!;

        /// <summary>
        /// This property applies only to files being downloaded to Linux compute nodes. It will be ignored if it is specified for a resourceFile which will be downloaded to a Windows node. If this property is not specified for a Linux node, then a default value of 0770 is applied to the file.
        /// </summary>
        [Input("fileMode")]
        public Input<string>? FileMode { get; set; }

        [Input("filePath", required: true)]
        public Input<string> FilePath { get; set; } = null!;

        public ResourceFileArgs()
        {
        }
        public static new ResourceFileArgs Empty => new ResourceFileArgs();
    }
}
