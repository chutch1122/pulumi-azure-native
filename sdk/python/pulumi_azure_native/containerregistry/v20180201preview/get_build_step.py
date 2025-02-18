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
    'GetBuildStepResult',
    'AwaitableGetBuildStepResult',
    'get_build_step',
    'get_build_step_output',
]

@pulumi.output_type
class GetBuildStepResult:
    """
    Build step resource properties
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
        The resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.DockerBuildStepResponse':
        """
        The properties of a build step.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetBuildStepResult(GetBuildStepResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBuildStepResult(
            id=self.id,
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_build_step(build_task_name: Optional[str] = None,
                   registry_name: Optional[str] = None,
                   resource_group_name: Optional[str] = None,
                   step_name: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBuildStepResult:
    """
    Build step resource properties


    :param str build_task_name: The name of the container registry build task.
    :param str registry_name: The name of the container registry.
    :param str resource_group_name: The name of the resource group to which the container registry belongs.
    :param str step_name: The name of a build step for a container registry build task.
    """
    __args__ = dict()
    __args__['buildTaskName'] = build_task_name
    __args__['registryName'] = registry_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['stepName'] = step_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:containerregistry/v20180201preview:getBuildStep', __args__, opts=opts, typ=GetBuildStepResult).value

    return AwaitableGetBuildStepResult(
        id=__ret__.id,
        name=__ret__.name,
        properties=__ret__.properties,
        type=__ret__.type)


@_utilities.lift_output_func(get_build_step)
def get_build_step_output(build_task_name: Optional[pulumi.Input[str]] = None,
                          registry_name: Optional[pulumi.Input[str]] = None,
                          resource_group_name: Optional[pulumi.Input[str]] = None,
                          step_name: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBuildStepResult]:
    """
    Build step resource properties


    :param str build_task_name: The name of the container registry build task.
    :param str registry_name: The name of the container registry.
    :param str resource_group_name: The name of the resource group to which the container registry belongs.
    :param str step_name: The name of a build step for a container registry build task.
    """
    ...
