// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Blueprint.V20171111Preview.Inputs
{

    /// <summary>
    /// Base class for ParameterValue.
    /// </summary>
    public sealed class ParameterValueBaseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional property, just to establish ParameterValueBase as a BaseClass.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        public ParameterValueBaseArgs()
        {
        }
        public static new ParameterValueBaseArgs Empty => new ParameterValueBaseArgs();
    }
}
