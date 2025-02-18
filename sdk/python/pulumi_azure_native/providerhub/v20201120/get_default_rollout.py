# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetDefaultRolloutResult',
    'AwaitableGetDefaultRolloutResult',
    'get_default_rollout',
    'get_default_rollout_output',
]

@pulumi.output_type
class GetDefaultRolloutResult:
    """
    Default rollout definition.
    """
    def __init__(__self__, id=None, name=None, properties=None, type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)
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
    @pulumi.getter
    def properties(self) -> 'outputs.DefaultRolloutResponseProperties':
        """
        Properties of the rollout.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetDefaultRolloutResult(GetDefaultRolloutResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDefaultRolloutResult(
            id=self.id,
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_default_rollout(provider_namespace: Optional[str] = None,
                        rollout_name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDefaultRolloutResult:
    """
    Default rollout definition.


    :param str provider_namespace: The name of the resource provider hosted within ProviderHub.
    :param str rollout_name: The rollout name.
    """
    __args__ = dict()
    __args__['providerNamespace'] = provider_namespace
    __args__['rolloutName'] = rollout_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:providerhub/v20201120:getDefaultRollout', __args__, opts=opts, typ=GetDefaultRolloutResult).value

    return AwaitableGetDefaultRolloutResult(
        id=__ret__.id,
        name=__ret__.name,
        properties=__ret__.properties,
        type=__ret__.type)


@_utilities.lift_output_func(get_default_rollout)
def get_default_rollout_output(provider_namespace: Optional[pulumi.Input[str]] = None,
                               rollout_name: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDefaultRolloutResult]:
    """
    Default rollout definition.


    :param str provider_namespace: The name of the resource provider hosted within ProviderHub.
    :param str rollout_name: The rollout name.
    """
    ...
