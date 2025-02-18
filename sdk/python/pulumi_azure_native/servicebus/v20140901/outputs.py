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
    'MessageCountDetailsResponse',
    'SkuResponse',
]

@pulumi.output_type
class MessageCountDetailsResponse(dict):
    """
    Message Count Details.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "activeMessageCount":
            suggest = "active_message_count"
        elif key == "deadLetterMessageCount":
            suggest = "dead_letter_message_count"
        elif key == "scheduledMessageCount":
            suggest = "scheduled_message_count"
        elif key == "transferDeadLetterMessageCount":
            suggest = "transfer_dead_letter_message_count"
        elif key == "transferMessageCount":
            suggest = "transfer_message_count"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in MessageCountDetailsResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        MessageCountDetailsResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        MessageCountDetailsResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 active_message_count: float,
                 dead_letter_message_count: float,
                 scheduled_message_count: float,
                 transfer_dead_letter_message_count: float,
                 transfer_message_count: float):
        """
        Message Count Details.
        :param float active_message_count: Number of active messages in the queue, topic, or subscription.
        :param float dead_letter_message_count: Number of messages that are dead lettered.
        :param float scheduled_message_count: Number of scheduled messages.
        :param float transfer_dead_letter_message_count: Number of messages transferred into dead letters.
        :param float transfer_message_count: Number of messages transferred to another queue, topic, or subscription.
        """
        pulumi.set(__self__, "active_message_count", active_message_count)
        pulumi.set(__self__, "dead_letter_message_count", dead_letter_message_count)
        pulumi.set(__self__, "scheduled_message_count", scheduled_message_count)
        pulumi.set(__self__, "transfer_dead_letter_message_count", transfer_dead_letter_message_count)
        pulumi.set(__self__, "transfer_message_count", transfer_message_count)

    @property
    @pulumi.getter(name="activeMessageCount")
    def active_message_count(self) -> float:
        """
        Number of active messages in the queue, topic, or subscription.
        """
        return pulumi.get(self, "active_message_count")

    @property
    @pulumi.getter(name="deadLetterMessageCount")
    def dead_letter_message_count(self) -> float:
        """
        Number of messages that are dead lettered.
        """
        return pulumi.get(self, "dead_letter_message_count")

    @property
    @pulumi.getter(name="scheduledMessageCount")
    def scheduled_message_count(self) -> float:
        """
        Number of scheduled messages.
        """
        return pulumi.get(self, "scheduled_message_count")

    @property
    @pulumi.getter(name="transferDeadLetterMessageCount")
    def transfer_dead_letter_message_count(self) -> float:
        """
        Number of messages transferred into dead letters.
        """
        return pulumi.get(self, "transfer_dead_letter_message_count")

    @property
    @pulumi.getter(name="transferMessageCount")
    def transfer_message_count(self) -> float:
        """
        Number of messages transferred to another queue, topic, or subscription.
        """
        return pulumi.get(self, "transfer_message_count")


@pulumi.output_type
class SkuResponse(dict):
    """
    SKU of the namespace.
    """
    def __init__(__self__, *,
                 tier: str,
                 capacity: Optional[int] = None,
                 name: Optional[str] = None):
        """
        SKU of the namespace.
        :param str tier: The billing tier of this particular SKU.
        :param int capacity: The specified messaging units for the tier.
        :param str name: Name of this SKU.
        """
        pulumi.set(__self__, "tier", tier)
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def tier(self) -> str:
        """
        The billing tier of this particular SKU.
        """
        return pulumi.get(self, "tier")

    @property
    @pulumi.getter
    def capacity(self) -> Optional[int]:
        """
        The specified messaging units for the tier.
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of this SKU.
        """
        return pulumi.get(self, "name")


