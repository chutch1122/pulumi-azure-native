// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Logic.V20160601.Inputs
{

    /// <summary>
    /// The X12 message filter for odata query.
    /// </summary>
    public sealed class X12MessageFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The message filter type.
        /// </summary>
        [Input("messageFilterType", required: true)]
        public Input<Pulumi.AzureNative.Logic.V20160601.MessageFilterType> MessageFilterType { get; set; } = null!;

        public X12MessageFilterArgs()
        {
        }
        public static new X12MessageFilterArgs Empty => new X12MessageFilterArgs();
    }
}
