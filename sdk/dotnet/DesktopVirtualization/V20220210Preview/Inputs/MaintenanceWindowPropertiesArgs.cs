// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DesktopVirtualization.V20220210Preview.Inputs
{

    /// <summary>
    /// Maintenance window starting hour and day of week.
    /// </summary>
    public sealed class MaintenanceWindowPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The day of the week (Monday-Sunday).
        /// </summary>
        [Input("dayOfWeek")]
        public Input<Pulumi.AzureNative.DesktopVirtualization.V20220210Preview.DayOfWeek>? DayOfWeek { get; set; }

        /// <summary>
        /// The starting hour of the maintenance window (0-23). Note that maintenance windows are 2 hours long. This means that updates can be applied anytime from the specified start hour to 2 hours after.
        /// </summary>
        [Input("hour")]
        public Input<int>? Hour { get; set; }

        public MaintenanceWindowPropertiesArgs()
        {
        }
        public static new MaintenanceWindowPropertiesArgs Empty => new MaintenanceWindowPropertiesArgs();
    }
}
