// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cdn.V20190415.Inputs
{

    /// <summary>
    /// Defines the parameters for RequestMethod match conditions
    /// </summary>
    public sealed class RequestMethodMatchConditionParametersArgs : global::Pulumi.ResourceArgs
    {
        [Input("matchValues")]
        private InputList<string>? _matchValues;

        /// <summary>
        /// The match value for the condition of the delivery rule
        /// </summary>
        public InputList<string> MatchValues
        {
            get => _matchValues ?? (_matchValues = new InputList<string>());
            set => _matchValues = value;
        }

        /// <summary>
        /// Describes if this is negate condition or not
        /// </summary>
        [Input("negateCondition")]
        public Input<bool>? NegateCondition { get; set; }

        [Input("odataType", required: true)]
        public Input<string> OdataType { get; set; } = null!;

        /// <summary>
        /// Describes operator to be matched
        /// </summary>
        [Input("operator", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Cdn.V20190415.RequestMethodOperator> Operator { get; set; } = null!;

        public RequestMethodMatchConditionParametersArgs()
        {
        }
        public static new RequestMethodMatchConditionParametersArgs Empty => new RequestMethodMatchConditionParametersArgs();
    }
}
