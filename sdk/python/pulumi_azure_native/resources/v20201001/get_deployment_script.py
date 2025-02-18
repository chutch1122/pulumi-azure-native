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
    'GetDeploymentScriptResult',
    'AwaitableGetDeploymentScriptResult',
    'get_deployment_script',
    'get_deployment_script_output',
]

warnings.warn("""Please use one of the variants: AzureCliScript, AzurePowerShellScript.""", DeprecationWarning)

@pulumi.output_type
class GetDeploymentScriptResult:
    """
    Deployment script object.
    """
    def __init__(__self__, id=None, identity=None, kind=None, location=None, name=None, system_data=None, tags=None, type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        String Id used to locate any resource on Azure.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.ManagedServiceIdentityResponse']:
        """
        Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the script.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location of the ACI and the storage account for the deployment script.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of this resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        The system metadata related to this resource.
        """
        return pulumi.get(self, "system_data")

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
        Type of this resource.
        """
        return pulumi.get(self, "type")


class AwaitableGetDeploymentScriptResult(GetDeploymentScriptResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeploymentScriptResult(
            id=self.id,
            identity=self.identity,
            kind=self.kind,
            location=self.location,
            name=self.name,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type)


def get_deployment_script(resource_group_name: Optional[str] = None,
                          script_name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeploymentScriptResult:
    """
    Deployment script object.


    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str script_name: Name of the deployment script.
    """
    pulumi.log.warn("""get_deployment_script is deprecated: Please use one of the variants: AzureCliScript, AzurePowerShellScript.""")
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['scriptName'] = script_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:resources/v20201001:getDeploymentScript', __args__, opts=opts, typ=GetDeploymentScriptResult).value

    return AwaitableGetDeploymentScriptResult(
        id=__ret__.id,
        identity=__ret__.identity,
        kind=__ret__.kind,
        location=__ret__.location,
        name=__ret__.name,
        system_data=__ret__.system_data,
        tags=__ret__.tags,
        type=__ret__.type)


@_utilities.lift_output_func(get_deployment_script)
def get_deployment_script_output(resource_group_name: Optional[pulumi.Input[str]] = None,
                                 script_name: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDeploymentScriptResult]:
    """
    Deployment script object.


    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str script_name: Name of the deployment script.
    """
    pulumi.log.warn("""get_deployment_script is deprecated: Please use one of the variants: AzureCliScript, AzurePowerShellScript.""")
    ...
