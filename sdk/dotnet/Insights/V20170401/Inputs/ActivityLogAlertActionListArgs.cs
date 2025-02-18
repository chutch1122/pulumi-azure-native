// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Insights.V20170401.Inputs
{

    /// <summary>
    /// A list of activity log alert actions.
    /// </summary>
    public sealed class ActivityLogAlertActionListArgs : global::Pulumi.ResourceArgs
    {
        [Input("actionGroups")]
        private InputList<Inputs.ActivityLogAlertActionGroupArgs>? _actionGroups;

        /// <summary>
        /// The list of activity log alerts.
        /// </summary>
        public InputList<Inputs.ActivityLogAlertActionGroupArgs> ActionGroups
        {
            get => _actionGroups ?? (_actionGroups = new InputList<Inputs.ActivityLogAlertActionGroupArgs>());
            set => _actionGroups = value;
        }

        public ActivityLogAlertActionListArgs()
        {
        }
        public static new ActivityLogAlertActionListArgs Empty => new ActivityLogAlertActionListArgs();
    }
}
