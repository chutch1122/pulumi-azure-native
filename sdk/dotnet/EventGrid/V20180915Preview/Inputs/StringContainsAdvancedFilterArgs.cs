// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid.V20180915Preview.Inputs
{

    /// <summary>
    /// StringContains Filter
    /// </summary>
    public sealed class StringContainsAdvancedFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The filter key. Represents an event property with up to two levels of nesting.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Represents the filter operator
        /// Expected value is 'StringContains'.
        /// </summary>
        [Input("operatorType", required: true)]
        public Input<string> OperatorType { get; set; } = null!;

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// The set of filter values
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public StringContainsAdvancedFilterArgs()
        {
        }
        public static new StringContainsAdvancedFilterArgs Empty => new StringContainsAdvancedFilterArgs();
    }
}
