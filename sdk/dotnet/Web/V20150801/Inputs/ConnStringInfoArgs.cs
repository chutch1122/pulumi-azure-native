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
    /// Represents database connection string information
    /// </summary>
    public sealed class ConnStringInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Connection string value
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        /// <summary>
        /// Name of connection string
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Type of database
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AzureNative.Web.V20150801.DatabaseServerType> Type { get; set; } = null!;

        public ConnStringInfoArgs()
        {
        }
        public static new ConnStringInfoArgs Empty => new ConnStringInfoArgs();
    }
}
