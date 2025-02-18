# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetSapMonitorResult',
    'AwaitableGetSapMonitorResult',
    'get_sap_monitor',
    'get_sap_monitor_output',
]

@pulumi.output_type
class GetSapMonitorResult:
    """
    SAP monitor info on Azure (ARM properties and SAP monitor properties)
    """
    def __init__(__self__, enable_customer_analytics=None, id=None, location=None, log_analytics_workspace_arm_id=None, log_analytics_workspace_id=None, log_analytics_workspace_shared_key=None, managed_resource_group_name=None, monitor_subnet=None, name=None, provisioning_state=None, sap_monitor_collector_version=None, tags=None, type=None):
        if enable_customer_analytics and not isinstance(enable_customer_analytics, bool):
            raise TypeError("Expected argument 'enable_customer_analytics' to be a bool")
        pulumi.set(__self__, "enable_customer_analytics", enable_customer_analytics)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if log_analytics_workspace_arm_id and not isinstance(log_analytics_workspace_arm_id, str):
            raise TypeError("Expected argument 'log_analytics_workspace_arm_id' to be a str")
        pulumi.set(__self__, "log_analytics_workspace_arm_id", log_analytics_workspace_arm_id)
        if log_analytics_workspace_id and not isinstance(log_analytics_workspace_id, str):
            raise TypeError("Expected argument 'log_analytics_workspace_id' to be a str")
        pulumi.set(__self__, "log_analytics_workspace_id", log_analytics_workspace_id)
        if log_analytics_workspace_shared_key and not isinstance(log_analytics_workspace_shared_key, str):
            raise TypeError("Expected argument 'log_analytics_workspace_shared_key' to be a str")
        pulumi.set(__self__, "log_analytics_workspace_shared_key", log_analytics_workspace_shared_key)
        if managed_resource_group_name and not isinstance(managed_resource_group_name, str):
            raise TypeError("Expected argument 'managed_resource_group_name' to be a str")
        pulumi.set(__self__, "managed_resource_group_name", managed_resource_group_name)
        if monitor_subnet and not isinstance(monitor_subnet, str):
            raise TypeError("Expected argument 'monitor_subnet' to be a str")
        pulumi.set(__self__, "monitor_subnet", monitor_subnet)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if sap_monitor_collector_version and not isinstance(sap_monitor_collector_version, str):
            raise TypeError("Expected argument 'sap_monitor_collector_version' to be a str")
        pulumi.set(__self__, "sap_monitor_collector_version", sap_monitor_collector_version)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="enableCustomerAnalytics")
    def enable_customer_analytics(self) -> Optional[bool]:
        """
        The value indicating whether to send analytics to Microsoft
        """
        return pulumi.get(self, "enable_customer_analytics")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="logAnalyticsWorkspaceArmId")
    def log_analytics_workspace_arm_id(self) -> Optional[str]:
        """
        The ARM ID of the Log Analytics Workspace that is used for monitoring
        """
        return pulumi.get(self, "log_analytics_workspace_arm_id")

    @property
    @pulumi.getter(name="logAnalyticsWorkspaceId")
    def log_analytics_workspace_id(self) -> Optional[str]:
        """
        The workspace ID of the log analytics workspace to be used for monitoring
        """
        return pulumi.get(self, "log_analytics_workspace_id")

    @property
    @pulumi.getter(name="logAnalyticsWorkspaceSharedKey")
    def log_analytics_workspace_shared_key(self) -> Optional[str]:
        """
        The shared key of the log analytics workspace that is used for monitoring
        """
        return pulumi.get(self, "log_analytics_workspace_shared_key")

    @property
    @pulumi.getter(name="managedResourceGroupName")
    def managed_resource_group_name(self) -> str:
        """
        The name of the resource group the SAP Monitor resources get deployed into.
        """
        return pulumi.get(self, "managed_resource_group_name")

    @property
    @pulumi.getter(name="monitorSubnet")
    def monitor_subnet(self) -> Optional[str]:
        """
        The subnet which the SAP monitor will be deployed in
        """
        return pulumi.get(self, "monitor_subnet")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        State of provisioning of the HanaInstance
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="sapMonitorCollectorVersion")
    def sap_monitor_collector_version(self) -> str:
        """
        The version of the payload running in the Collector VM
        """
        return pulumi.get(self, "sap_monitor_collector_version")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetSapMonitorResult(GetSapMonitorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSapMonitorResult(
            enable_customer_analytics=self.enable_customer_analytics,
            id=self.id,
            location=self.location,
            log_analytics_workspace_arm_id=self.log_analytics_workspace_arm_id,
            log_analytics_workspace_id=self.log_analytics_workspace_id,
            log_analytics_workspace_shared_key=self.log_analytics_workspace_shared_key,
            managed_resource_group_name=self.managed_resource_group_name,
            monitor_subnet=self.monitor_subnet,
            name=self.name,
            provisioning_state=self.provisioning_state,
            sap_monitor_collector_version=self.sap_monitor_collector_version,
            tags=self.tags,
            type=self.type)


def get_sap_monitor(resource_group_name: Optional[str] = None,
                    sap_monitor_name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSapMonitorResult:
    """
    SAP monitor info on Azure (ARM properties and SAP monitor properties)
    API Version: 2020-02-07-preview.


    :param str resource_group_name: Name of the resource group.
    :param str sap_monitor_name: Name of the SAP monitor resource.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['sapMonitorName'] = sap_monitor_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:hanaonazure:getSapMonitor', __args__, opts=opts, typ=GetSapMonitorResult).value

    return AwaitableGetSapMonitorResult(
        enable_customer_analytics=__ret__.enable_customer_analytics,
        id=__ret__.id,
        location=__ret__.location,
        log_analytics_workspace_arm_id=__ret__.log_analytics_workspace_arm_id,
        log_analytics_workspace_id=__ret__.log_analytics_workspace_id,
        log_analytics_workspace_shared_key=__ret__.log_analytics_workspace_shared_key,
        managed_resource_group_name=__ret__.managed_resource_group_name,
        monitor_subnet=__ret__.monitor_subnet,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        sap_monitor_collector_version=__ret__.sap_monitor_collector_version,
        tags=__ret__.tags,
        type=__ret__.type)


@_utilities.lift_output_func(get_sap_monitor)
def get_sap_monitor_output(resource_group_name: Optional[pulumi.Input[str]] = None,
                           sap_monitor_name: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSapMonitorResult]:
    """
    SAP monitor info on Azure (ARM properties and SAP monitor properties)
    API Version: 2020-02-07-preview.


    :param str resource_group_name: Name of the resource group.
    :param str sap_monitor_name: Name of the SAP monitor resource.
    """
    ...
