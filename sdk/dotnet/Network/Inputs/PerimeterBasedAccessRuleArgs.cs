// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Inputs
{

    public sealed class PerimeterBasedAccessRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// NSP id in the ARM id format.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public PerimeterBasedAccessRuleArgs()
        {
        }
        public static new PerimeterBasedAccessRuleArgs Empty => new PerimeterBasedAccessRuleArgs();
    }
}
