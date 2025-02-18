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
    'GetAccountResult',
    'AwaitableGetAccountResult',
    'get_account',
    'get_account_output',
]

@pulumi.output_type
class GetAccountResult:
    """
    An object that represents a machine learning team account.
    """
    def __init__(__self__, account_id=None, creation_date=None, description=None, discovery_uri=None, friendly_name=None, id=None, key_vault_id=None, location=None, name=None, provisioning_state=None, seats=None, storage_account=None, tags=None, type=None, vso_account_id=None):
        if account_id and not isinstance(account_id, str):
            raise TypeError("Expected argument 'account_id' to be a str")
        pulumi.set(__self__, "account_id", account_id)
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if discovery_uri and not isinstance(discovery_uri, str):
            raise TypeError("Expected argument 'discovery_uri' to be a str")
        pulumi.set(__self__, "discovery_uri", discovery_uri)
        if friendly_name and not isinstance(friendly_name, str):
            raise TypeError("Expected argument 'friendly_name' to be a str")
        pulumi.set(__self__, "friendly_name", friendly_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key_vault_id and not isinstance(key_vault_id, str):
            raise TypeError("Expected argument 'key_vault_id' to be a str")
        pulumi.set(__self__, "key_vault_id", key_vault_id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if seats and not isinstance(seats, str):
            raise TypeError("Expected argument 'seats' to be a str")
        pulumi.set(__self__, "seats", seats)
        if storage_account and not isinstance(storage_account, dict):
            raise TypeError("Expected argument 'storage_account' to be a dict")
        pulumi.set(__self__, "storage_account", storage_account)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vso_account_id and not isinstance(vso_account_id, str):
            raise TypeError("Expected argument 'vso_account_id' to be a str")
        pulumi.set(__self__, "vso_account_id", vso_account_id)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> str:
        """
        The immutable id associated with this team account.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> str:
        """
        The creation date of the machine learning team account in ISO8601 format.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of this workspace.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="discoveryUri")
    def discovery_uri(self) -> str:
        """
        The uri for this machine learning team account.
        """
        return pulumi.get(self, "discovery_uri")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[str]:
        """
        The friendly name for this workspace. This will be the workspace name in the arm id when the workspace object gets created
        """
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> str:
        """
        The fully qualified arm id of the user key vault.
        """
        return pulumi.get(self, "key_vault_id")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location of the resource. This cannot be changed after the resource is created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The current deployment state of team account resource. The provisioningState is to indicate states for resource provisioning.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def seats(self) -> Optional[str]:
        """
        The no of users/seats who can access this team account. This property defines the charge on the team account.
        """
        return pulumi.get(self, "seats")

    @property
    @pulumi.getter(name="storageAccount")
    def storage_account(self) -> 'outputs.StorageAccountPropertiesResponse':
        """
        The properties of the storage account for the machine learning team account.
        """
        return pulumi.get(self, "storage_account")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vsoAccountId")
    def vso_account_id(self) -> str:
        """
        The fully qualified arm id of the vso account to be used for this team account.
        """
        return pulumi.get(self, "vso_account_id")


class AwaitableGetAccountResult(GetAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccountResult(
            account_id=self.account_id,
            creation_date=self.creation_date,
            description=self.description,
            discovery_uri=self.discovery_uri,
            friendly_name=self.friendly_name,
            id=self.id,
            key_vault_id=self.key_vault_id,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            seats=self.seats,
            storage_account=self.storage_account,
            tags=self.tags,
            type=self.type,
            vso_account_id=self.vso_account_id)


def get_account(account_name: Optional[str] = None,
                resource_group_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccountResult:
    """
    An object that represents a machine learning team account.
    API Version: 2017-05-01-preview.


    :param str account_name: The name of the machine learning team account.
    :param str resource_group_name: The name of the resource group to which the machine learning team account belongs.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:machinelearningexperimentation:getAccount', __args__, opts=opts, typ=GetAccountResult).value

    return AwaitableGetAccountResult(
        account_id=__ret__.account_id,
        creation_date=__ret__.creation_date,
        description=__ret__.description,
        discovery_uri=__ret__.discovery_uri,
        friendly_name=__ret__.friendly_name,
        id=__ret__.id,
        key_vault_id=__ret__.key_vault_id,
        location=__ret__.location,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        seats=__ret__.seats,
        storage_account=__ret__.storage_account,
        tags=__ret__.tags,
        type=__ret__.type,
        vso_account_id=__ret__.vso_account_id)


@_utilities.lift_output_func(get_account)
def get_account_output(account_name: Optional[pulumi.Input[str]] = None,
                       resource_group_name: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccountResult]:
    """
    An object that represents a machine learning team account.
    API Version: 2017-05-01-preview.


    :param str account_name: The name of the machine learning team account.
    :param str resource_group_name: The name of the resource group to which the machine learning team account belongs.
    """
    ...
