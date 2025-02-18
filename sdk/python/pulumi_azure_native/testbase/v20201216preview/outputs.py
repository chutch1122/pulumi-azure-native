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
    'CommandResponse',
    'DistributionGroupListReceiverValueResponse',
    'NotificationEventReceiverResponse',
    'NotificationReceiverValueResponse',
    'PackageValidationResultResponse',
    'SubscriptionReceiverValueResponse',
    'SystemDataResponse',
    'TargetOSInfoResponse',
    'TestBaseAccountSKUCapabilityResponse',
    'TestBaseAccountSKUResponse',
    'TestResponse',
    'UserObjectReceiverValueResponse',
]

@pulumi.output_type
class CommandResponse(dict):
    """
    The command used in the test
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "contentType":
            suggest = "content_type"
        elif key == "alwaysRun":
            suggest = "always_run"
        elif key == "applyUpdateBefore":
            suggest = "apply_update_before"
        elif key == "maxRunTime":
            suggest = "max_run_time"
        elif key == "restartAfter":
            suggest = "restart_after"
        elif key == "runAsInteractive":
            suggest = "run_as_interactive"
        elif key == "runElevated":
            suggest = "run_elevated"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CommandResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CommandResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CommandResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 action: str,
                 content: str,
                 content_type: str,
                 name: str,
                 always_run: Optional[bool] = None,
                 apply_update_before: Optional[bool] = None,
                 max_run_time: Optional[int] = None,
                 restart_after: Optional[bool] = None,
                 run_as_interactive: Optional[bool] = None,
                 run_elevated: Optional[bool] = None):
        """
        The command used in the test
        :param str action: The action of the command.
        :param str content: The content of the command. The content depends on source type.
        :param str content_type: The type of command content.
        :param str name: The name of the command.
        :param bool always_run: Specifies whether to run the command even if a previous command is failed.
        :param bool apply_update_before: Specifies whether to apply update before the command.
        :param int max_run_time: Specifies the max run time of the command.
        :param bool restart_after: Specifies whether to restart the VM after the command executed.
        :param bool run_as_interactive: Specifies whether to run the command in interactive mode.
        :param bool run_elevated: Specifies whether to run the command as administrator.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "content_type", content_type)
        pulumi.set(__self__, "name", name)
        if always_run is not None:
            pulumi.set(__self__, "always_run", always_run)
        if apply_update_before is not None:
            pulumi.set(__self__, "apply_update_before", apply_update_before)
        if max_run_time is not None:
            pulumi.set(__self__, "max_run_time", max_run_time)
        if restart_after is not None:
            pulumi.set(__self__, "restart_after", restart_after)
        if run_as_interactive is not None:
            pulumi.set(__self__, "run_as_interactive", run_as_interactive)
        if run_elevated is not None:
            pulumi.set(__self__, "run_elevated", run_elevated)

    @property
    @pulumi.getter
    def action(self) -> str:
        """
        The action of the command.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def content(self) -> str:
        """
        The content of the command. The content depends on source type.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> str:
        """
        The type of command content.
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the command.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="alwaysRun")
    def always_run(self) -> Optional[bool]:
        """
        Specifies whether to run the command even if a previous command is failed.
        """
        return pulumi.get(self, "always_run")

    @property
    @pulumi.getter(name="applyUpdateBefore")
    def apply_update_before(self) -> Optional[bool]:
        """
        Specifies whether to apply update before the command.
        """
        return pulumi.get(self, "apply_update_before")

    @property
    @pulumi.getter(name="maxRunTime")
    def max_run_time(self) -> Optional[int]:
        """
        Specifies the max run time of the command.
        """
        return pulumi.get(self, "max_run_time")

    @property
    @pulumi.getter(name="restartAfter")
    def restart_after(self) -> Optional[bool]:
        """
        Specifies whether to restart the VM after the command executed.
        """
        return pulumi.get(self, "restart_after")

    @property
    @pulumi.getter(name="runAsInteractive")
    def run_as_interactive(self) -> Optional[bool]:
        """
        Specifies whether to run the command in interactive mode.
        """
        return pulumi.get(self, "run_as_interactive")

    @property
    @pulumi.getter(name="runElevated")
    def run_elevated(self) -> Optional[bool]:
        """
        Specifies whether to run the command as administrator.
        """
        return pulumi.get(self, "run_elevated")


@pulumi.output_type
class DistributionGroupListReceiverValueResponse(dict):
    """
    The user object receiver value.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "distributionGroups":
            suggest = "distribution_groups"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DistributionGroupListReceiverValueResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DistributionGroupListReceiverValueResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DistributionGroupListReceiverValueResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 distribution_groups: Optional[Sequence[str]] = None):
        """
        The user object receiver value.
        :param Sequence[str] distribution_groups: The list of distribution groups.
        """
        if distribution_groups is not None:
            pulumi.set(__self__, "distribution_groups", distribution_groups)

    @property
    @pulumi.getter(name="distributionGroups")
    def distribution_groups(self) -> Optional[Sequence[str]]:
        """
        The list of distribution groups.
        """
        return pulumi.get(self, "distribution_groups")


@pulumi.output_type
class NotificationEventReceiverResponse(dict):
    """
    A notification event receivers.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "receiverType":
            suggest = "receiver_type"
        elif key == "receiverValue":
            suggest = "receiver_value"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NotificationEventReceiverResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NotificationEventReceiverResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NotificationEventReceiverResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 receiver_type: Optional[str] = None,
                 receiver_value: Optional['outputs.NotificationReceiverValueResponse'] = None):
        """
        A notification event receivers.
        :param str receiver_type: The type of the notification event receiver.
        :param 'NotificationReceiverValueResponse' receiver_value: The notification event receiver value.
        """
        if receiver_type is not None:
            pulumi.set(__self__, "receiver_type", receiver_type)
        if receiver_value is not None:
            pulumi.set(__self__, "receiver_value", receiver_value)

    @property
    @pulumi.getter(name="receiverType")
    def receiver_type(self) -> Optional[str]:
        """
        The type of the notification event receiver.
        """
        return pulumi.get(self, "receiver_type")

    @property
    @pulumi.getter(name="receiverValue")
    def receiver_value(self) -> Optional['outputs.NotificationReceiverValueResponse']:
        """
        The notification event receiver value.
        """
        return pulumi.get(self, "receiver_value")


@pulumi.output_type
class NotificationReceiverValueResponse(dict):
    """
    A notification event receiver value.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "distributionGroupListReceiverValue":
            suggest = "distribution_group_list_receiver_value"
        elif key == "subscriptionReceiverValue":
            suggest = "subscription_receiver_value"
        elif key == "userObjectReceiverValue":
            suggest = "user_object_receiver_value"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NotificationReceiverValueResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NotificationReceiverValueResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NotificationReceiverValueResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 distribution_group_list_receiver_value: Optional['outputs.DistributionGroupListReceiverValueResponse'] = None,
                 subscription_receiver_value: Optional['outputs.SubscriptionReceiverValueResponse'] = None,
                 user_object_receiver_value: Optional['outputs.UserObjectReceiverValueResponse'] = None):
        """
        A notification event receiver value.
        :param 'DistributionGroupListReceiverValueResponse' distribution_group_list_receiver_value: The user object receiver value.
        :param 'SubscriptionReceiverValueResponse' subscription_receiver_value: The user object receiver value.
        :param 'UserObjectReceiverValueResponse' user_object_receiver_value: The user object receiver value.
        """
        if distribution_group_list_receiver_value is not None:
            pulumi.set(__self__, "distribution_group_list_receiver_value", distribution_group_list_receiver_value)
        if subscription_receiver_value is not None:
            pulumi.set(__self__, "subscription_receiver_value", subscription_receiver_value)
        if user_object_receiver_value is not None:
            pulumi.set(__self__, "user_object_receiver_value", user_object_receiver_value)

    @property
    @pulumi.getter(name="distributionGroupListReceiverValue")
    def distribution_group_list_receiver_value(self) -> Optional['outputs.DistributionGroupListReceiverValueResponse']:
        """
        The user object receiver value.
        """
        return pulumi.get(self, "distribution_group_list_receiver_value")

    @property
    @pulumi.getter(name="subscriptionReceiverValue")
    def subscription_receiver_value(self) -> Optional['outputs.SubscriptionReceiverValueResponse']:
        """
        The user object receiver value.
        """
        return pulumi.get(self, "subscription_receiver_value")

    @property
    @pulumi.getter(name="userObjectReceiverValue")
    def user_object_receiver_value(self) -> Optional['outputs.UserObjectReceiverValueResponse']:
        """
        The user object receiver value.
        """
        return pulumi.get(self, "user_object_receiver_value")


@pulumi.output_type
class PackageValidationResultResponse(dict):
    """
    The validation results. There's validation on package when it's created or updated.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "isValid":
            suggest = "is_valid"
        elif key == "validationName":
            suggest = "validation_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PackageValidationResultResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PackageValidationResultResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PackageValidationResultResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 errors: Sequence[str],
                 is_valid: bool,
                 validation_name: str):
        """
        The validation results. There's validation on package when it's created or updated.
        :param Sequence[str] errors: Error information.
        :param bool is_valid: Indicates whether the package passed the validation.
        :param str validation_name: Validation name.
        """
        pulumi.set(__self__, "errors", errors)
        pulumi.set(__self__, "is_valid", is_valid)
        pulumi.set(__self__, "validation_name", validation_name)

    @property
    @pulumi.getter
    def errors(self) -> Sequence[str]:
        """
        Error information.
        """
        return pulumi.get(self, "errors")

    @property
    @pulumi.getter(name="isValid")
    def is_valid(self) -> bool:
        """
        Indicates whether the package passed the validation.
        """
        return pulumi.get(self, "is_valid")

    @property
    @pulumi.getter(name="validationName")
    def validation_name(self) -> str:
        """
        Validation name.
        """
        return pulumi.get(self, "validation_name")


@pulumi.output_type
class SubscriptionReceiverValueResponse(dict):
    """
    The subscription role receiver value.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "subscriptionId":
            suggest = "subscription_id"
        elif key == "subscriptionName":
            suggest = "subscription_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SubscriptionReceiverValueResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SubscriptionReceiverValueResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SubscriptionReceiverValueResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 role: Optional[str] = None,
                 subscription_id: Optional[str] = None,
                 subscription_name: Optional[str] = None):
        """
        The subscription role receiver value.
        :param str role: The role of the notification receiver.
        :param str subscription_id: The subscription id of the notification receiver.
        :param str subscription_name: The subscription name of the notification receiver.
        """
        if role is not None:
            pulumi.set(__self__, "role", role)
        if subscription_id is not None:
            pulumi.set(__self__, "subscription_id", subscription_id)
        if subscription_name is not None:
            pulumi.set(__self__, "subscription_name", subscription_name)

    @property
    @pulumi.getter
    def role(self) -> Optional[str]:
        """
        The role of the notification receiver.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> Optional[str]:
        """
        The subscription id of the notification receiver.
        """
        return pulumi.get(self, "subscription_id")

    @property
    @pulumi.getter(name="subscriptionName")
    def subscription_name(self) -> Optional[str]:
        """
        The subscription name of the notification receiver.
        """
        return pulumi.get(self, "subscription_name")


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
        :param str last_modified_at: The type of identity that last modified the resource.
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
        The type of identity that last modified the resource.
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


@pulumi.output_type
class TargetOSInfoResponse(dict):
    """
    The information of the target OS to be tested.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "osUpdateType":
            suggest = "os_update_type"
        elif key == "targetOSs":
            suggest = "target_oss"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TargetOSInfoResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TargetOSInfoResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TargetOSInfoResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 os_update_type: str,
                 target_oss: Sequence[str]):
        """
        The information of the target OS to be tested.
        :param str os_update_type: Specifies the OS update type to test against, e.g., 'Security updates' or 'Feature updates'.
        :param Sequence[str] target_oss: Specifies the target OSs to be tested.
        """
        pulumi.set(__self__, "os_update_type", os_update_type)
        pulumi.set(__self__, "target_oss", target_oss)

    @property
    @pulumi.getter(name="osUpdateType")
    def os_update_type(self) -> str:
        """
        Specifies the OS update type to test against, e.g., 'Security updates' or 'Feature updates'.
        """
        return pulumi.get(self, "os_update_type")

    @property
    @pulumi.getter(name="targetOSs")
    def target_oss(self) -> Sequence[str]:
        """
        Specifies the target OSs to be tested.
        """
        return pulumi.get(self, "target_oss")


@pulumi.output_type
class TestBaseAccountSKUCapabilityResponse(dict):
    """
    Properties of the Test Base Account SKU Capability.
    """
    def __init__(__self__, *,
                 name: str,
                 value: str):
        """
        Properties of the Test Base Account SKU Capability.
        :param str name: An invariant to describe the feature, such as 'SLA'.
        :param str value: An invariant if the feature is measured by quantity, such as 99.9%.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        An invariant to describe the feature, such as 'SLA'.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        An invariant if the feature is measured by quantity, such as 99.9%.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class TestBaseAccountSKUResponse(dict):
    """
    Describes a Test Base Account SKU.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "resourceType":
            suggest = "resource_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TestBaseAccountSKUResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TestBaseAccountSKUResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TestBaseAccountSKUResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 capabilities: Sequence['outputs.TestBaseAccountSKUCapabilityResponse'],
                 name: str,
                 tier: str,
                 locations: Optional[Sequence[str]] = None,
                 resource_type: Optional[str] = None):
        """
        Describes a Test Base Account SKU.
        :param Sequence['TestBaseAccountSKUCapabilityResponse'] capabilities: The capabilities of a SKU.
        :param str name: The name of the SKU. This is typically a letter + number code, such as B0 or S0.
        :param str tier: The tier of this particular SKU.
        :param Sequence[str] locations: The locations that the SKU is available.
        :param str resource_type: The type of resource the SKU applies to.
        """
        pulumi.set(__self__, "capabilities", capabilities)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "tier", tier)
        if locations is not None:
            pulumi.set(__self__, "locations", locations)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter
    def capabilities(self) -> Sequence['outputs.TestBaseAccountSKUCapabilityResponse']:
        """
        The capabilities of a SKU.
        """
        return pulumi.get(self, "capabilities")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the SKU. This is typically a letter + number code, such as B0 or S0.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tier(self) -> str:
        """
        The tier of this particular SKU.
        """
        return pulumi.get(self, "tier")

    @property
    @pulumi.getter
    def locations(self) -> Optional[Sequence[str]]:
        """
        The locations that the SKU is available.
        """
        return pulumi.get(self, "locations")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[str]:
        """
        The type of resource the SKU applies to.
        """
        return pulumi.get(self, "resource_type")


@pulumi.output_type
class TestResponse(dict):
    """
    The definition of a Test.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "testType":
            suggest = "test_type"
        elif key == "validationRunStatus":
            suggest = "validation_run_status"
        elif key == "isActive":
            suggest = "is_active"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TestResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TestResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TestResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 commands: Sequence['outputs.CommandResponse'],
                 test_type: str,
                 validation_run_status: str,
                 is_active: Optional[bool] = None):
        """
        The definition of a Test.
        :param Sequence['CommandResponse'] commands: The commands used in the test.
        :param str test_type: The type of the test.
        :param str validation_run_status: The status of the validation run of the package.
        :param bool is_active: Indicates if this test is active.It doesn't schedule test for not active Test.
        """
        pulumi.set(__self__, "commands", commands)
        pulumi.set(__self__, "test_type", test_type)
        pulumi.set(__self__, "validation_run_status", validation_run_status)
        if is_active is not None:
            pulumi.set(__self__, "is_active", is_active)

    @property
    @pulumi.getter
    def commands(self) -> Sequence['outputs.CommandResponse']:
        """
        The commands used in the test.
        """
        return pulumi.get(self, "commands")

    @property
    @pulumi.getter(name="testType")
    def test_type(self) -> str:
        """
        The type of the test.
        """
        return pulumi.get(self, "test_type")

    @property
    @pulumi.getter(name="validationRunStatus")
    def validation_run_status(self) -> str:
        """
        The status of the validation run of the package.
        """
        return pulumi.get(self, "validation_run_status")

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> Optional[bool]:
        """
        Indicates if this test is active.It doesn't schedule test for not active Test.
        """
        return pulumi.get(self, "is_active")


@pulumi.output_type
class UserObjectReceiverValueResponse(dict):
    """
    The user object receiver value.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "userObjectIds":
            suggest = "user_object_ids"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in UserObjectReceiverValueResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        UserObjectReceiverValueResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        UserObjectReceiverValueResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 user_object_ids: Optional[Sequence[str]] = None):
        """
        The user object receiver value.
        :param Sequence[str] user_object_ids: user object ids.
        """
        if user_object_ids is not None:
            pulumi.set(__self__, "user_object_ids", user_object_ids)

    @property
    @pulumi.getter(name="userObjectIds")
    def user_object_ids(self) -> Optional[Sequence[str]]:
        """
        user object ids.
        """
        return pulumi.get(self, "user_object_ids")


