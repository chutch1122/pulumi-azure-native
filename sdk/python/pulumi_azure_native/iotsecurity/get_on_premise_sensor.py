# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetOnPremiseSensorResult',
    'AwaitableGetOnPremiseSensorResult',
    'get_on_premise_sensor',
    'get_on_premise_sensor_output',
]

@pulumi.output_type
class GetOnPremiseSensorResult:
    """
    On-premise IoT sensor
    """
    def __init__(__self__, id=None, name=None, system_data=None, type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetOnPremiseSensorResult(GetOnPremiseSensorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOnPremiseSensorResult(
            id=self.id,
            name=self.name,
            system_data=self.system_data,
            type=self.type)


def get_on_premise_sensor(on_premise_sensor_name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOnPremiseSensorResult:
    """
    On-premise IoT sensor
    API Version: 2021-02-01-preview.


    :param str on_premise_sensor_name: Name of the on-premise IoT sensor
    """
    __args__ = dict()
    __args__['onPremiseSensorName'] = on_premise_sensor_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:iotsecurity:getOnPremiseSensor', __args__, opts=opts, typ=GetOnPremiseSensorResult).value

    return AwaitableGetOnPremiseSensorResult(
        id=__ret__.id,
        name=__ret__.name,
        system_data=__ret__.system_data,
        type=__ret__.type)


@_utilities.lift_output_func(get_on_premise_sensor)
def get_on_premise_sensor_output(on_premise_sensor_name: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOnPremiseSensorResult]:
    """
    On-premise IoT sensor
    API Version: 2021-02-01-preview.


    :param str on_premise_sensor_name: Name of the on-premise IoT sensor
    """
    ...
