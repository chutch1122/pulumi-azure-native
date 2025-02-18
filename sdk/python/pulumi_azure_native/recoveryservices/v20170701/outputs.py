# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = [
    'AzureRecoveryServiceVaultProtectionIntentResponse',
    'AzureResourceProtectionIntentResponse',
    'AzureWorkloadAutoProtectionIntentResponse',
    'AzureWorkloadSQLAutoProtectionIntentResponse',
]

@pulumi.output_type
class AzureRecoveryServiceVaultProtectionIntentResponse(dict):
    """
    Azure Recovery Services Vault specific protection intent item.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "backupManagementType":
            suggest = "backup_management_type"
        elif key == "itemId":
            suggest = "item_id"
        elif key == "policyId":
            suggest = "policy_id"
        elif key == "protectionIntentItemType":
            suggest = "protection_intent_item_type"
        elif key == "protectionState":
            suggest = "protection_state"
        elif key == "sourceResourceId":
            suggest = "source_resource_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AzureRecoveryServiceVaultProtectionIntentResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AzureRecoveryServiceVaultProtectionIntentResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AzureRecoveryServiceVaultProtectionIntentResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 backup_management_type: Optional[str] = None,
                 item_id: Optional[str] = None,
                 policy_id: Optional[str] = None,
                 protection_intent_item_type: Optional[str] = None,
                 protection_state: Optional[str] = None,
                 source_resource_id: Optional[str] = None):
        """
        Azure Recovery Services Vault specific protection intent item.
        :param str backup_management_type: Type of backup management for the backed up item.
        :param str item_id: ID of the item which is getting protected, In case of Azure Vm , it is ProtectedItemId
        :param str policy_id: ID of the backup policy with which this item is backed up.
        :param str protection_intent_item_type: backup protectionIntent type.
               Expected value is 'RecoveryServiceVaultItem'.
        :param str protection_state: Backup state of this backup item.
        :param str source_resource_id: ARM ID of the resource to be backed up.
        """
        if backup_management_type is not None:
            pulumi.set(__self__, "backup_management_type", backup_management_type)
        if item_id is not None:
            pulumi.set(__self__, "item_id", item_id)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if protection_intent_item_type is not None:
            pulumi.set(__self__, "protection_intent_item_type", 'RecoveryServiceVaultItem')
        if protection_state is not None:
            pulumi.set(__self__, "protection_state", protection_state)
        if source_resource_id is not None:
            pulumi.set(__self__, "source_resource_id", source_resource_id)

    @property
    @pulumi.getter(name="backupManagementType")
    def backup_management_type(self) -> Optional[str]:
        """
        Type of backup management for the backed up item.
        """
        return pulumi.get(self, "backup_management_type")

    @property
    @pulumi.getter(name="itemId")
    def item_id(self) -> Optional[str]:
        """
        ID of the item which is getting protected, In case of Azure Vm , it is ProtectedItemId
        """
        return pulumi.get(self, "item_id")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[str]:
        """
        ID of the backup policy with which this item is backed up.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="protectionIntentItemType")
    def protection_intent_item_type(self) -> Optional[str]:
        """
        backup protectionIntent type.
        Expected value is 'RecoveryServiceVaultItem'.
        """
        return pulumi.get(self, "protection_intent_item_type")

    @property
    @pulumi.getter(name="protectionState")
    def protection_state(self) -> Optional[str]:
        """
        Backup state of this backup item.
        """
        return pulumi.get(self, "protection_state")

    @property
    @pulumi.getter(name="sourceResourceId")
    def source_resource_id(self) -> Optional[str]:
        """
        ARM ID of the resource to be backed up.
        """
        return pulumi.get(self, "source_resource_id")


@pulumi.output_type
class AzureResourceProtectionIntentResponse(dict):
    """
    IaaS VM specific backup protection intent item.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "backupManagementType":
            suggest = "backup_management_type"
        elif key == "friendlyName":
            suggest = "friendly_name"
        elif key == "itemId":
            suggest = "item_id"
        elif key == "policyId":
            suggest = "policy_id"
        elif key == "protectionIntentItemType":
            suggest = "protection_intent_item_type"
        elif key == "protectionState":
            suggest = "protection_state"
        elif key == "sourceResourceId":
            suggest = "source_resource_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AzureResourceProtectionIntentResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AzureResourceProtectionIntentResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AzureResourceProtectionIntentResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 backup_management_type: Optional[str] = None,
                 friendly_name: Optional[str] = None,
                 item_id: Optional[str] = None,
                 policy_id: Optional[str] = None,
                 protection_intent_item_type: Optional[str] = None,
                 protection_state: Optional[str] = None,
                 source_resource_id: Optional[str] = None):
        """
        IaaS VM specific backup protection intent item.
        :param str backup_management_type: Type of backup management for the backed up item.
        :param str friendly_name: Friendly name of the VM represented by this backup item.
        :param str item_id: ID of the item which is getting protected, In case of Azure Vm , it is ProtectedItemId
        :param str policy_id: ID of the backup policy with which this item is backed up.
        :param str protection_intent_item_type: backup protectionIntent type.
               Expected value is 'AzureResourceItem'.
        :param str protection_state: Backup state of this backup item.
        :param str source_resource_id: ARM ID of the resource to be backed up.
        """
        if backup_management_type is not None:
            pulumi.set(__self__, "backup_management_type", backup_management_type)
        if friendly_name is not None:
            pulumi.set(__self__, "friendly_name", friendly_name)
        if item_id is not None:
            pulumi.set(__self__, "item_id", item_id)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if protection_intent_item_type is not None:
            pulumi.set(__self__, "protection_intent_item_type", 'AzureResourceItem')
        if protection_state is not None:
            pulumi.set(__self__, "protection_state", protection_state)
        if source_resource_id is not None:
            pulumi.set(__self__, "source_resource_id", source_resource_id)

    @property
    @pulumi.getter(name="backupManagementType")
    def backup_management_type(self) -> Optional[str]:
        """
        Type of backup management for the backed up item.
        """
        return pulumi.get(self, "backup_management_type")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[str]:
        """
        Friendly name of the VM represented by this backup item.
        """
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter(name="itemId")
    def item_id(self) -> Optional[str]:
        """
        ID of the item which is getting protected, In case of Azure Vm , it is ProtectedItemId
        """
        return pulumi.get(self, "item_id")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[str]:
        """
        ID of the backup policy with which this item is backed up.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="protectionIntentItemType")
    def protection_intent_item_type(self) -> Optional[str]:
        """
        backup protectionIntent type.
        Expected value is 'AzureResourceItem'.
        """
        return pulumi.get(self, "protection_intent_item_type")

    @property
    @pulumi.getter(name="protectionState")
    def protection_state(self) -> Optional[str]:
        """
        Backup state of this backup item.
        """
        return pulumi.get(self, "protection_state")

    @property
    @pulumi.getter(name="sourceResourceId")
    def source_resource_id(self) -> Optional[str]:
        """
        ARM ID of the resource to be backed up.
        """
        return pulumi.get(self, "source_resource_id")


@pulumi.output_type
class AzureWorkloadAutoProtectionIntentResponse(dict):
    """
    Azure Recovery Services Vault specific protection intent item.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "backupManagementType":
            suggest = "backup_management_type"
        elif key == "itemId":
            suggest = "item_id"
        elif key == "policyId":
            suggest = "policy_id"
        elif key == "protectionIntentItemType":
            suggest = "protection_intent_item_type"
        elif key == "protectionState":
            suggest = "protection_state"
        elif key == "sourceResourceId":
            suggest = "source_resource_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AzureWorkloadAutoProtectionIntentResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AzureWorkloadAutoProtectionIntentResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AzureWorkloadAutoProtectionIntentResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 backup_management_type: Optional[str] = None,
                 item_id: Optional[str] = None,
                 policy_id: Optional[str] = None,
                 protection_intent_item_type: Optional[str] = None,
                 protection_state: Optional[str] = None,
                 source_resource_id: Optional[str] = None):
        """
        Azure Recovery Services Vault specific protection intent item.
        :param str backup_management_type: Type of backup management for the backed up item.
        :param str item_id: ID of the item which is getting protected, In case of Azure Vm , it is ProtectedItemId
        :param str policy_id: ID of the backup policy with which this item is backed up.
        :param str protection_intent_item_type: backup protectionIntent type.
               Expected value is 'AzureWorkloadAutoProtectionIntent'.
        :param str protection_state: Backup state of this backup item.
        :param str source_resource_id: ARM ID of the resource to be backed up.
        """
        if backup_management_type is not None:
            pulumi.set(__self__, "backup_management_type", backup_management_type)
        if item_id is not None:
            pulumi.set(__self__, "item_id", item_id)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if protection_intent_item_type is not None:
            pulumi.set(__self__, "protection_intent_item_type", 'AzureWorkloadAutoProtectionIntent')
        if protection_state is not None:
            pulumi.set(__self__, "protection_state", protection_state)
        if source_resource_id is not None:
            pulumi.set(__self__, "source_resource_id", source_resource_id)

    @property
    @pulumi.getter(name="backupManagementType")
    def backup_management_type(self) -> Optional[str]:
        """
        Type of backup management for the backed up item.
        """
        return pulumi.get(self, "backup_management_type")

    @property
    @pulumi.getter(name="itemId")
    def item_id(self) -> Optional[str]:
        """
        ID of the item which is getting protected, In case of Azure Vm , it is ProtectedItemId
        """
        return pulumi.get(self, "item_id")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[str]:
        """
        ID of the backup policy with which this item is backed up.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="protectionIntentItemType")
    def protection_intent_item_type(self) -> Optional[str]:
        """
        backup protectionIntent type.
        Expected value is 'AzureWorkloadAutoProtectionIntent'.
        """
        return pulumi.get(self, "protection_intent_item_type")

    @property
    @pulumi.getter(name="protectionState")
    def protection_state(self) -> Optional[str]:
        """
        Backup state of this backup item.
        """
        return pulumi.get(self, "protection_state")

    @property
    @pulumi.getter(name="sourceResourceId")
    def source_resource_id(self) -> Optional[str]:
        """
        ARM ID of the resource to be backed up.
        """
        return pulumi.get(self, "source_resource_id")


@pulumi.output_type
class AzureWorkloadSQLAutoProtectionIntentResponse(dict):
    """
    Azure Workload SQL Auto Protection intent item.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "backupManagementType":
            suggest = "backup_management_type"
        elif key == "itemId":
            suggest = "item_id"
        elif key == "policyId":
            suggest = "policy_id"
        elif key == "protectionIntentItemType":
            suggest = "protection_intent_item_type"
        elif key == "protectionState":
            suggest = "protection_state"
        elif key == "sourceResourceId":
            suggest = "source_resource_id"
        elif key == "workloadItemType":
            suggest = "workload_item_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AzureWorkloadSQLAutoProtectionIntentResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AzureWorkloadSQLAutoProtectionIntentResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AzureWorkloadSQLAutoProtectionIntentResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 backup_management_type: Optional[str] = None,
                 item_id: Optional[str] = None,
                 policy_id: Optional[str] = None,
                 protection_intent_item_type: Optional[str] = None,
                 protection_state: Optional[str] = None,
                 source_resource_id: Optional[str] = None,
                 workload_item_type: Optional[str] = None):
        """
        Azure Workload SQL Auto Protection intent item.
        :param str backup_management_type: Type of backup management for the backed up item.
        :param str item_id: ID of the item which is getting protected, In case of Azure Vm , it is ProtectedItemId
        :param str policy_id: ID of the backup policy with which this item is backed up.
        :param str protection_intent_item_type: backup protectionIntent type.
               Expected value is 'AzureWorkloadSQLAutoProtectionIntent'.
        :param str protection_state: Backup state of this backup item.
        :param str source_resource_id: ARM ID of the resource to be backed up.
        :param str workload_item_type: Workload item type of the item for which intent is to be set
        """
        if backup_management_type is not None:
            pulumi.set(__self__, "backup_management_type", backup_management_type)
        if item_id is not None:
            pulumi.set(__self__, "item_id", item_id)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if protection_intent_item_type is not None:
            pulumi.set(__self__, "protection_intent_item_type", 'AzureWorkloadSQLAutoProtectionIntent')
        if protection_state is not None:
            pulumi.set(__self__, "protection_state", protection_state)
        if source_resource_id is not None:
            pulumi.set(__self__, "source_resource_id", source_resource_id)
        if workload_item_type is not None:
            pulumi.set(__self__, "workload_item_type", workload_item_type)

    @property
    @pulumi.getter(name="backupManagementType")
    def backup_management_type(self) -> Optional[str]:
        """
        Type of backup management for the backed up item.
        """
        return pulumi.get(self, "backup_management_type")

    @property
    @pulumi.getter(name="itemId")
    def item_id(self) -> Optional[str]:
        """
        ID of the item which is getting protected, In case of Azure Vm , it is ProtectedItemId
        """
        return pulumi.get(self, "item_id")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[str]:
        """
        ID of the backup policy with which this item is backed up.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="protectionIntentItemType")
    def protection_intent_item_type(self) -> Optional[str]:
        """
        backup protectionIntent type.
        Expected value is 'AzureWorkloadSQLAutoProtectionIntent'.
        """
        return pulumi.get(self, "protection_intent_item_type")

    @property
    @pulumi.getter(name="protectionState")
    def protection_state(self) -> Optional[str]:
        """
        Backup state of this backup item.
        """
        return pulumi.get(self, "protection_state")

    @property
    @pulumi.getter(name="sourceResourceId")
    def source_resource_id(self) -> Optional[str]:
        """
        ARM ID of the resource to be backed up.
        """
        return pulumi.get(self, "source_resource_id")

    @property
    @pulumi.getter(name="workloadItemType")
    def workload_item_type(self) -> Optional[str]:
        """
        Workload item type of the item for which intent is to be set
        """
        return pulumi.get(self, "workload_item_type")


