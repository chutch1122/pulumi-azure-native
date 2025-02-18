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
    'GetMediaServiceResult',
    'AwaitableGetMediaServiceResult',
    'get_media_service',
    'get_media_service_output',
]

@pulumi.output_type
class GetMediaServiceResult:
    """
    A Media Services account.
    """
    def __init__(__self__, encryption=None, id=None, identity=None, location=None, media_service_id=None, name=None, storage_accounts=None, storage_authentication=None, system_data=None, tags=None, type=None):
        if encryption and not isinstance(encryption, dict):
            raise TypeError("Expected argument 'encryption' to be a dict")
        pulumi.set(__self__, "encryption", encryption)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if media_service_id and not isinstance(media_service_id, str):
            raise TypeError("Expected argument 'media_service_id' to be a str")
        pulumi.set(__self__, "media_service_id", media_service_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if storage_accounts and not isinstance(storage_accounts, list):
            raise TypeError("Expected argument 'storage_accounts' to be a list")
        pulumi.set(__self__, "storage_accounts", storage_accounts)
        if storage_authentication and not isinstance(storage_authentication, str):
            raise TypeError("Expected argument 'storage_authentication' to be a str")
        pulumi.set(__self__, "storage_authentication", storage_authentication)
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
    def encryption(self) -> Optional['outputs.AccountEncryptionResponse']:
        """
        The account encryption properties.
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.MediaServiceIdentityResponse']:
        """
        The Managed Identity for the Media Services account.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="mediaServiceId")
    def media_service_id(self) -> str:
        """
        The Media Services account ID.
        """
        return pulumi.get(self, "media_service_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="storageAccounts")
    def storage_accounts(self) -> Optional[Sequence['outputs.StorageAccountResponse']]:
        """
        The storage accounts for this resource.
        """
        return pulumi.get(self, "storage_accounts")

    @property
    @pulumi.getter(name="storageAuthentication")
    def storage_authentication(self) -> Optional[str]:
        return pulumi.get(self, "storage_authentication")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        The system metadata relating to this resource.
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
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetMediaServiceResult(GetMediaServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMediaServiceResult(
            encryption=self.encryption,
            id=self.id,
            identity=self.identity,
            location=self.location,
            media_service_id=self.media_service_id,
            name=self.name,
            storage_accounts=self.storage_accounts,
            storage_authentication=self.storage_authentication,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type)


def get_media_service(account_name: Optional[str] = None,
                      resource_group_name: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMediaServiceResult:
    """
    A Media Services account.


    :param str account_name: The Media Services account name.
    :param str resource_group_name: The name of the resource group within the Azure subscription.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:media/v20200501:getMediaService', __args__, opts=opts, typ=GetMediaServiceResult).value

    return AwaitableGetMediaServiceResult(
        encryption=__ret__.encryption,
        id=__ret__.id,
        identity=__ret__.identity,
        location=__ret__.location,
        media_service_id=__ret__.media_service_id,
        name=__ret__.name,
        storage_accounts=__ret__.storage_accounts,
        storage_authentication=__ret__.storage_authentication,
        system_data=__ret__.system_data,
        tags=__ret__.tags,
        type=__ret__.type)


@_utilities.lift_output_func(get_media_service)
def get_media_service_output(account_name: Optional[pulumi.Input[str]] = None,
                             resource_group_name: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMediaServiceResult]:
    """
    A Media Services account.


    :param str account_name: The Media Services account name.
    :param str resource_group_name: The name of the resource group within the Azure subscription.
    """
    ...
