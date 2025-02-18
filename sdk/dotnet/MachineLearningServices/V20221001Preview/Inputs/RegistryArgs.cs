// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20221001Preview.Inputs
{

    /// <summary>
    /// Details of the Registry
    /// </summary>
    public sealed class RegistryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The asset description text.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("discoveryUrl")]
        public Input<string>? DiscoveryUrl { get; set; }

        [Input("intellectualPropertyPublisher")]
        public Input<string>? IntellectualPropertyPublisher { get; set; }

        /// <summary>
        /// Managed resource group created for the registry
        /// </summary>
        [Input("managedResourceGroup")]
        public Input<Inputs.ArmResourceIdArgs>? ManagedResourceGroup { get; set; }

        [Input("mlFlowRegistryUri")]
        public Input<string>? MlFlowRegistryUri { get; set; }

        [Input("privateLinkCount")]
        public Input<int>? PrivateLinkCount { get; set; }

        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// The asset property dictionary.
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        [Input("publicNetworkAccess")]
        public Input<string>? PublicNetworkAccess { get; set; }

        [Input("regionDetails")]
        private InputList<Inputs.RegistryRegionArmDetailsArgs>? _regionDetails;

        /// <summary>
        /// Details of each region the registry is in
        /// </summary>
        public InputList<Inputs.RegistryRegionArmDetailsArgs> RegionDetails
        {
            get => _regionDetails ?? (_regionDetails = new InputList<Inputs.RegistryRegionArmDetailsArgs>());
            set => _regionDetails = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tag dictionary. Tags can be added, removed, and updated.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public RegistryArgs()
        {
        }
        public static new RegistryArgs Empty => new RegistryArgs();
    }
}
