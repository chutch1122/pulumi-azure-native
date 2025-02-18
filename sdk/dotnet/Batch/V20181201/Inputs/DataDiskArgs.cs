// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.V20181201.Inputs
{

    /// <summary>
    /// Data Disk settings which will be used by the data disks associated to Compute Nodes in the pool.
    /// </summary>
    public sealed class DataDiskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Values are:
        /// 
        ///  none - The caching mode for the disk is not enabled.
        ///  readOnly - The caching mode for the disk is read only.
        ///  readWrite - The caching mode for the disk is read and write.
        /// 
        ///  The default value for caching is none. For information about the caching options see: https://blogs.msdn.microsoft.com/windowsazurestorage/2012/06/27/exploring-windows-azure-drives-disks-and-images/.
        /// </summary>
        [Input("caching")]
        public Input<Pulumi.AzureNative.Batch.V20181201.CachingType>? Caching { get; set; }

        [Input("diskSizeGB", required: true)]
        public Input<int> DiskSizeGB { get; set; } = null!;

        /// <summary>
        /// The lun is used to uniquely identify each data disk. If attaching multiple disks, each should have a distinct lun.
        /// </summary>
        [Input("lun", required: true)]
        public Input<int> Lun { get; set; } = null!;

        /// <summary>
        /// If omitted, the default is "Standard_LRS". Values are:
        /// 
        ///  Standard_LRS - The data disk should use standard locally redundant storage.
        ///  Premium_LRS - The data disk should use premium locally redundant storage.
        /// </summary>
        [Input("storageAccountType")]
        public Input<Pulumi.AzureNative.Batch.V20181201.StorageAccountType>? StorageAccountType { get; set; }

        public DataDiskArgs()
        {
        }
        public static new DataDiskArgs Empty => new DataDiskArgs();
    }
}
