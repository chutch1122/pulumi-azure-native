// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.V20150504Preview.Inputs
{

    /// <summary>
    /// A CNAME record.
    /// </summary>
    public sealed class CnameRecordArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the canonical name for this record without a terminating dot.
        /// </summary>
        [Input("cname")]
        public Input<string>? Cname { get; set; }

        public CnameRecordArgs()
        {
        }
        public static new CnameRecordArgs Empty => new CnameRecordArgs();
    }
}
