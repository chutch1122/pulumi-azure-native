// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.V20150801Preview.Inputs
{

    /// <summary>
    /// API Management
    /// </summary>
    public sealed class ApiEntityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// API definition Url - url where the swagger can be downloaded from
        /// </summary>
        [Input("apiDefinitionUrl")]
        public Input<string>? ApiDefinitionUrl { get; set; }

        /// <summary>
        /// Backend service definition
        /// </summary>
        [Input("backendService")]
        public Input<Inputs.BackendServiceDefinitionArgs>? BackendService { get; set; }

        [Input("capabilities")]
        private InputList<string>? _capabilities;

        /// <summary>
        /// Capabilities
        /// </summary>
        public InputList<string> Capabilities
        {
            get => _capabilities ?? (_capabilities = new InputList<string>());
            set => _capabilities = value;
        }

        /// <summary>
        /// Timestamp of last connection change.
        /// </summary>
        [Input("changedTime")]
        public Input<string>? ChangedTime { get; set; }

        [Input("connectionParameters")]
        private InputMap<Inputs.ConnectionParameterArgs>? _connectionParameters;

        /// <summary>
        /// Connection parameters
        /// </summary>
        public InputMap<Inputs.ConnectionParameterArgs> ConnectionParameters
        {
            get => _connectionParameters ?? (_connectionParameters = new InputMap<Inputs.ConnectionParameterArgs>());
            set => _connectionParameters = value;
        }

        /// <summary>
        /// Timestamp of the connection creation
        /// </summary>
        [Input("createdTime")]
        public Input<string>? CreatedTime { get; set; }

        /// <summary>
        /// the URL path of this API when exposed via APIM
        /// </summary>
        [Input("generalInformation")]
        public Input<Inputs.GeneralApiInformationArgs>? GeneralInformation { get; set; }

        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Kind of resource
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Resource Location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Free form object for the data caller wants to store
        /// </summary>
        [Input("metadata")]
        public Input<object>? Metadata { get; set; }

        /// <summary>
        /// Resource Name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// the URL path of this API when exposed via APIM
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// API policies
        /// </summary>
        [Input("policies")]
        public Input<Inputs.ApiPoliciesArgs>? Policies { get; set; }

        [Input("protocols")]
        private InputList<string>? _protocols;

        /// <summary>
        /// Protocols supported by the front end - http/https
        /// </summary>
        public InputList<string> Protocols
        {
            get => _protocols ?? (_protocols = new InputList<string>());
            set => _protocols = value;
        }

        [Input("runtimeUrls")]
        private InputList<string>? _runtimeUrls;

        /// <summary>
        /// Read only property returning the runtime endpoints where the API can be called
        /// </summary>
        public InputList<string> RuntimeUrls
        {
            get => _runtimeUrls ?? (_runtimeUrls = new InputList<string>());
            set => _runtimeUrls = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Resource type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ApiEntityArgs()
        {
        }
        public static new ApiEntityArgs Empty => new ApiEntityArgs();
    }
}
