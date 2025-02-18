// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media.V20200201Preview.Inputs
{

    /// <summary>
    /// RTSP source.
    /// </summary>
    public sealed class MediaGraphRtspSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// RTSP endpoint of the stream being connected to.
        /// </summary>
        [Input("endpoint", required: true)]
        public InputUnion<Inputs.MediaGraphClearEndpointArgs, Inputs.MediaGraphTlsEndpointArgs> Endpoint { get; set; } = null!;

        /// <summary>
        /// Source name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The discriminator for derived types.
        /// Expected value is '#Microsoft.Media.MediaGraphRtspSource'.
        /// </summary>
        [Input("odataType", required: true)]
        public Input<string> OdataType { get; set; } = null!;

        /// <summary>
        /// Underlying RTSP transport. This can be used to enable or disable HTTP tunneling.
        /// </summary>
        [Input("transport", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Media.V20200201Preview.MediaGraphRtspTransport> Transport { get; set; } = null!;

        public MediaGraphRtspSourceArgs()
        {
        }
        public static new MediaGraphRtspSourceArgs Empty => new MediaGraphRtspSourceArgs();
    }
}
