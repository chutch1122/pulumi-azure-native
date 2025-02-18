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
from ._enums import *

__all__ = [
    'AccessControlResponse',
    'AccountEncryptionResponse',
    'EdgeUsageDataCollectionPolicyResponse',
    'EdgeUsageDataEventHubResponse',
    'KeyDeliveryResponse',
    'KeyVaultPropertiesResponse',
    'MediaServiceIdentityResponse',
    'PrivateEndpointResponse',
    'PrivateLinkServiceConnectionStateResponse',
    'StorageAccountResponse',
    'SystemDataResponse',
]

@pulumi.output_type
class AccessControlResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "defaultAction":
            suggest = "default_action"
        elif key == "ipAllowList":
            suggest = "ip_allow_list"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AccessControlResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AccessControlResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AccessControlResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 default_action: Optional[str] = None,
                 ip_allow_list: Optional[Sequence[str]] = None):
        """
        :param str default_action: The behavior for IP access control in Key Delivery.
        :param Sequence[str] ip_allow_list: The IP allow list for access control in Key Delivery. If the default action is set to 'Allow', the IP allow list must be empty.
        """
        if default_action is not None:
            pulumi.set(__self__, "default_action", default_action)
        if ip_allow_list is not None:
            pulumi.set(__self__, "ip_allow_list", ip_allow_list)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional[str]:
        """
        The behavior for IP access control in Key Delivery.
        """
        return pulumi.get(self, "default_action")

    @property
    @pulumi.getter(name="ipAllowList")
    def ip_allow_list(self) -> Optional[Sequence[str]]:
        """
        The IP allow list for access control in Key Delivery. If the default action is set to 'Allow', the IP allow list must be empty.
        """
        return pulumi.get(self, "ip_allow_list")


@pulumi.output_type
class AccountEncryptionResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "keyVaultProperties":
            suggest = "key_vault_properties"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AccountEncryptionResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AccountEncryptionResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AccountEncryptionResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 type: str,
                 key_vault_properties: Optional['outputs.KeyVaultPropertiesResponse'] = None):
        """
        :param str type: The type of key used to encrypt the Account Key.
        :param 'KeyVaultPropertiesResponse' key_vault_properties: The properties of the key used to encrypt the account.
        """
        pulumi.set(__self__, "type", type)
        if key_vault_properties is not None:
            pulumi.set(__self__, "key_vault_properties", key_vault_properties)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of key used to encrypt the Account Key.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="keyVaultProperties")
    def key_vault_properties(self) -> Optional['outputs.KeyVaultPropertiesResponse']:
        """
        The properties of the key used to encrypt the account.
        """
        return pulumi.get(self, "key_vault_properties")


@pulumi.output_type
class EdgeUsageDataCollectionPolicyResponse(dict):
    def __init__(__self__, *,
                 data_collection_frequency: Optional[str] = None,
                 data_reporting_frequency: Optional[str] = None,
                 event_hub_details: Optional['outputs.EdgeUsageDataEventHubResponse'] = None,
                 max_allowed_unreported_usage_duration: Optional[str] = None):
        """
        :param str data_collection_frequency: Usage data collection frequency in ISO 8601 duration format e.g. PT10M , PT5H.
        :param str data_reporting_frequency: Usage data reporting frequency in ISO 8601 duration format e.g. PT10M , PT5H.
        :param 'EdgeUsageDataEventHubResponse' event_hub_details: Details of Event Hub where the usage will be reported.
        :param str max_allowed_unreported_usage_duration: Maximum time for which the functionality of the device will not be hampered for not reporting the usage data.
        """
        if data_collection_frequency is not None:
            pulumi.set(__self__, "data_collection_frequency", data_collection_frequency)
        if data_reporting_frequency is not None:
            pulumi.set(__self__, "data_reporting_frequency", data_reporting_frequency)
        if event_hub_details is not None:
            pulumi.set(__self__, "event_hub_details", event_hub_details)
        if max_allowed_unreported_usage_duration is not None:
            pulumi.set(__self__, "max_allowed_unreported_usage_duration", max_allowed_unreported_usage_duration)

    @property
    @pulumi.getter(name="dataCollectionFrequency")
    def data_collection_frequency(self) -> Optional[str]:
        """
        Usage data collection frequency in ISO 8601 duration format e.g. PT10M , PT5H.
        """
        return pulumi.get(self, "data_collection_frequency")

    @property
    @pulumi.getter(name="dataReportingFrequency")
    def data_reporting_frequency(self) -> Optional[str]:
        """
        Usage data reporting frequency in ISO 8601 duration format e.g. PT10M , PT5H.
        """
        return pulumi.get(self, "data_reporting_frequency")

    @property
    @pulumi.getter(name="eventHubDetails")
    def event_hub_details(self) -> Optional['outputs.EdgeUsageDataEventHubResponse']:
        """
        Details of Event Hub where the usage will be reported.
        """
        return pulumi.get(self, "event_hub_details")

    @property
    @pulumi.getter(name="maxAllowedUnreportedUsageDuration")
    def max_allowed_unreported_usage_duration(self) -> Optional[str]:
        """
        Maximum time for which the functionality of the device will not be hampered for not reporting the usage data.
        """
        return pulumi.get(self, "max_allowed_unreported_usage_duration")


@pulumi.output_type
class EdgeUsageDataEventHubResponse(dict):
    def __init__(__self__, *,
                 name: Optional[str] = None,
                 namespace: Optional[str] = None,
                 token: Optional[str] = None):
        """
        :param str name: Name of the Event Hub where usage will be reported.
        :param str namespace: Namespace of the Event Hub where usage will be reported.
        :param str token: SAS token needed to interact with Event Hub.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the Event Hub where usage will be reported.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        """
        Namespace of the Event Hub where usage will be reported.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def token(self) -> Optional[str]:
        """
        SAS token needed to interact with Event Hub.
        """
        return pulumi.get(self, "token")


@pulumi.output_type
class KeyDeliveryResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accessControl":
            suggest = "access_control"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in KeyDeliveryResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        KeyDeliveryResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        KeyDeliveryResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 access_control: Optional['outputs.AccessControlResponse'] = None):
        """
        :param 'AccessControlResponse' access_control: The access control properties for Key Delivery.
        """
        if access_control is not None:
            pulumi.set(__self__, "access_control", access_control)

    @property
    @pulumi.getter(name="accessControl")
    def access_control(self) -> Optional['outputs.AccessControlResponse']:
        """
        The access control properties for Key Delivery.
        """
        return pulumi.get(self, "access_control")


@pulumi.output_type
class KeyVaultPropertiesResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "currentKeyIdentifier":
            suggest = "current_key_identifier"
        elif key == "keyIdentifier":
            suggest = "key_identifier"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in KeyVaultPropertiesResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        KeyVaultPropertiesResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        KeyVaultPropertiesResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 current_key_identifier: str,
                 key_identifier: Optional[str] = None):
        """
        :param str current_key_identifier: The current key used to encrypt the Media Services account, including the key version.
        :param str key_identifier: The URL of the Key Vault key used to encrypt the account. The key may either be versioned (for example https://vault/keys/mykey/version1) or reference a key without a version (for example https://vault/keys/mykey).
        """
        pulumi.set(__self__, "current_key_identifier", current_key_identifier)
        if key_identifier is not None:
            pulumi.set(__self__, "key_identifier", key_identifier)

    @property
    @pulumi.getter(name="currentKeyIdentifier")
    def current_key_identifier(self) -> str:
        """
        The current key used to encrypt the Media Services account, including the key version.
        """
        return pulumi.get(self, "current_key_identifier")

    @property
    @pulumi.getter(name="keyIdentifier")
    def key_identifier(self) -> Optional[str]:
        """
        The URL of the Key Vault key used to encrypt the account. The key may either be versioned (for example https://vault/keys/mykey/version1) or reference a key without a version (for example https://vault/keys/mykey).
        """
        return pulumi.get(self, "key_identifier")


@pulumi.output_type
class MediaServiceIdentityResponse(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "principalId":
            suggest = "principal_id"
        elif key == "tenantId":
            suggest = "tenant_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in MediaServiceIdentityResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        MediaServiceIdentityResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        MediaServiceIdentityResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 principal_id: str,
                 tenant_id: str,
                 type: str):
        """
        :param str principal_id: The Principal ID of the identity.
        :param str tenant_id: The Tenant ID of the identity.
        :param str type: The identity type.
        """
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "tenant_id", tenant_id)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        The Principal ID of the identity.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The Tenant ID of the identity.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The identity type.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class PrivateEndpointResponse(dict):
    """
    The Private Endpoint resource.
    """
    def __init__(__self__, *,
                 id: str):
        """
        The Private Endpoint resource.
        :param str id: The ARM identifier for Private Endpoint
        """
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ARM identifier for Private Endpoint
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class PrivateLinkServiceConnectionStateResponse(dict):
    """
    A collection of information about the state of the connection between service consumer and provider.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "actionsRequired":
            suggest = "actions_required"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PrivateLinkServiceConnectionStateResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PrivateLinkServiceConnectionStateResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PrivateLinkServiceConnectionStateResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 actions_required: Optional[str] = None,
                 description: Optional[str] = None,
                 status: Optional[str] = None):
        """
        A collection of information about the state of the connection between service consumer and provider.
        :param str actions_required: A message indicating if changes on the service provider require any updates on the consumer.
        :param str description: The reason for approval/rejection of the connection.
        :param str status: Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
        """
        if actions_required is not None:
            pulumi.set(__self__, "actions_required", actions_required)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="actionsRequired")
    def actions_required(self) -> Optional[str]:
        """
        A message indicating if changes on the service provider require any updates on the consumer.
        """
        return pulumi.get(self, "actions_required")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The reason for approval/rejection of the connection.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class StorageAccountResponse(dict):
    """
    The storage account details.
    """
    def __init__(__self__, *,
                 type: str,
                 id: Optional[str] = None):
        """
        The storage account details.
        :param str type: The type of the storage account.
        :param str id: The ID of the storage account resource. Media Services relies on tables and queues as well as blobs, so the primary storage account must be a Standard Storage account (either Microsoft.ClassicStorage or Microsoft.Storage). Blob only storage accounts can be added as secondary storage accounts.
        """
        pulumi.set(__self__, "type", type)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the storage account.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of the storage account resource. Media Services relies on tables and queues as well as blobs, so the primary storage account must be a Standard Storage account (either Microsoft.ClassicStorage or Microsoft.Storage). Blob only storage accounts can be added as secondary storage accounts.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class SystemDataResponse(dict):
    """
    Metadata pertaining to creation and last modification of the resource.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "createdBy":
            suggest = "created_by"
        elif key == "createdByType":
            suggest = "created_by_type"
        elif key == "lastModifiedAt":
            suggest = "last_modified_at"
        elif key == "lastModifiedBy":
            suggest = "last_modified_by"
        elif key == "lastModifiedByType":
            suggest = "last_modified_by_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SystemDataResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SystemDataResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SystemDataResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created_at: Optional[str] = None,
                 created_by: Optional[str] = None,
                 created_by_type: Optional[str] = None,
                 last_modified_at: Optional[str] = None,
                 last_modified_by: Optional[str] = None,
                 last_modified_by_type: Optional[str] = None):
        """
        Metadata pertaining to creation and last modification of the resource.
        :param str created_at: The timestamp of resource creation (UTC).
        :param str created_by: The identity that created the resource.
        :param str created_by_type: The type of identity that created the resource.
        :param str last_modified_at: The timestamp of resource last modification (UTC)
        :param str last_modified_by: The identity that last modified the resource.
        :param str last_modified_by_type: The type of identity that last modified the resource.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if created_by is not None:
            pulumi.set(__self__, "created_by", created_by)
        if created_by_type is not None:
            pulumi.set(__self__, "created_by_type", created_by_type)
        if last_modified_at is not None:
            pulumi.set(__self__, "last_modified_at", last_modified_at)
        if last_modified_by is not None:
            pulumi.set(__self__, "last_modified_by", last_modified_by)
        if last_modified_by_type is not None:
            pulumi.set(__self__, "last_modified_by_type", last_modified_by_type)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[str]:
        """
        The timestamp of resource creation (UTC).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="createdBy")
    def created_by(self) -> Optional[str]:
        """
        The identity that created the resource.
        """
        return pulumi.get(self, "created_by")

    @property
    @pulumi.getter(name="createdByType")
    def created_by_type(self) -> Optional[str]:
        """
        The type of identity that created the resource.
        """
        return pulumi.get(self, "created_by_type")

    @property
    @pulumi.getter(name="lastModifiedAt")
    def last_modified_at(self) -> Optional[str]:
        """
        The timestamp of resource last modification (UTC)
        """
        return pulumi.get(self, "last_modified_at")

    @property
    @pulumi.getter(name="lastModifiedBy")
    def last_modified_by(self) -> Optional[str]:
        """
        The identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by")

    @property
    @pulumi.getter(name="lastModifiedByType")
    def last_modified_by_type(self) -> Optional[str]:
        """
        The type of identity that last modified the resource.
        """
        return pulumi.get(self, "last_modified_by_type")


