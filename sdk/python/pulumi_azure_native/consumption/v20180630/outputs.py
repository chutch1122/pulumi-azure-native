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
    'BudgetTimePeriodResponse',
    'CurrentSpendResponse',
    'FiltersResponse',
    'NotificationResponse',
]

@pulumi.output_type
class BudgetTimePeriodResponse(dict):
    """
    The start and end date for a budget.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "startDate":
            suggest = "start_date"
        elif key == "endDate":
            suggest = "end_date"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BudgetTimePeriodResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BudgetTimePeriodResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BudgetTimePeriodResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 start_date: str,
                 end_date: Optional[str] = None):
        """
        The start and end date for a budget.
        :param str start_date: The start date for the budget.
        :param str end_date: The end date for the budget. If not provided, we default this to 10 years from the start date.
        """
        pulumi.set(__self__, "start_date", start_date)
        if end_date is not None:
            pulumi.set(__self__, "end_date", end_date)

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> str:
        """
        The start date for the budget.
        """
        return pulumi.get(self, "start_date")

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> Optional[str]:
        """
        The end date for the budget. If not provided, we default this to 10 years from the start date.
        """
        return pulumi.get(self, "end_date")


@pulumi.output_type
class CurrentSpendResponse(dict):
    """
    The current amount of cost which is being tracked for a budget.
    """
    def __init__(__self__, *,
                 amount: float,
                 unit: str):
        """
        The current amount of cost which is being tracked for a budget.
        :param float amount: The total amount of cost which is being tracked by the budget.
        :param str unit: The unit of measure for the budget amount.
        """
        pulumi.set(__self__, "amount", amount)
        pulumi.set(__self__, "unit", unit)

    @property
    @pulumi.getter
    def amount(self) -> float:
        """
        The total amount of cost which is being tracked by the budget.
        """
        return pulumi.get(self, "amount")

    @property
    @pulumi.getter
    def unit(self) -> str:
        """
        The unit of measure for the budget amount.
        """
        return pulumi.get(self, "unit")


@pulumi.output_type
class FiltersResponse(dict):
    """
    May be used to filter budgets by resource group, resource, or meter.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "resourceGroups":
            suggest = "resource_groups"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FiltersResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FiltersResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FiltersResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 meters: Optional[Sequence[str]] = None,
                 resource_groups: Optional[Sequence[str]] = None,
                 resources: Optional[Sequence[str]] = None,
                 tags: Optional[Mapping[str, Sequence[str]]] = None):
        """
        May be used to filter budgets by resource group, resource, or meter.
        :param Sequence[str] meters: The list of filters on meters (GUID), mandatory for budgets of usage category. 
        :param Sequence[str] resource_groups: The list of filters on resource groups, allowed at subscription level only.
        :param Sequence[str] resources: The list of filters on resources.
        :param Mapping[str, Sequence[str]] tags: The dictionary of filters on tags.
        """
        if meters is not None:
            pulumi.set(__self__, "meters", meters)
        if resource_groups is not None:
            pulumi.set(__self__, "resource_groups", resource_groups)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def meters(self) -> Optional[Sequence[str]]:
        """
        The list of filters on meters (GUID), mandatory for budgets of usage category. 
        """
        return pulumi.get(self, "meters")

    @property
    @pulumi.getter(name="resourceGroups")
    def resource_groups(self) -> Optional[Sequence[str]]:
        """
        The list of filters on resource groups, allowed at subscription level only.
        """
        return pulumi.get(self, "resource_groups")

    @property
    @pulumi.getter
    def resources(self) -> Optional[Sequence[str]]:
        """
        The list of filters on resources.
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Sequence[str]]]:
        """
        The dictionary of filters on tags.
        """
        return pulumi.get(self, "tags")


@pulumi.output_type
class NotificationResponse(dict):
    """
    The notification associated with a budget.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "contactEmails":
            suggest = "contact_emails"
        elif key == "contactGroups":
            suggest = "contact_groups"
        elif key == "contactRoles":
            suggest = "contact_roles"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NotificationResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NotificationResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NotificationResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 contact_emails: Sequence[str],
                 enabled: bool,
                 operator: str,
                 threshold: float,
                 contact_groups: Optional[Sequence[str]] = None,
                 contact_roles: Optional[Sequence[str]] = None):
        """
        The notification associated with a budget.
        :param Sequence[str] contact_emails: Email addresses to send the budget notification to when the threshold is exceeded.
        :param bool enabled: The notification is enabled or not.
        :param str operator: The comparison operator.
        :param float threshold: Threshold value associated with a notification. Notification is sent when the cost exceeded the threshold. It is always percent and has to be between 0 and 1000.
        :param Sequence[str] contact_groups: Action groups to send the budget notification to when the threshold is exceeded.
        :param Sequence[str] contact_roles: Contact roles to send the budget notification to when the threshold is exceeded.
        """
        pulumi.set(__self__, "contact_emails", contact_emails)
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "operator", operator)
        pulumi.set(__self__, "threshold", threshold)
        if contact_groups is not None:
            pulumi.set(__self__, "contact_groups", contact_groups)
        if contact_roles is not None:
            pulumi.set(__self__, "contact_roles", contact_roles)

    @property
    @pulumi.getter(name="contactEmails")
    def contact_emails(self) -> Sequence[str]:
        """
        Email addresses to send the budget notification to when the threshold is exceeded.
        """
        return pulumi.get(self, "contact_emails")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        The notification is enabled or not.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def operator(self) -> str:
        """
        The comparison operator.
        """
        return pulumi.get(self, "operator")

    @property
    @pulumi.getter
    def threshold(self) -> float:
        """
        Threshold value associated with a notification. Notification is sent when the cost exceeded the threshold. It is always percent and has to be between 0 and 1000.
        """
        return pulumi.get(self, "threshold")

    @property
    @pulumi.getter(name="contactGroups")
    def contact_groups(self) -> Optional[Sequence[str]]:
        """
        Action groups to send the budget notification to when the threshold is exceeded.
        """
        return pulumi.get(self, "contact_groups")

    @property
    @pulumi.getter(name="contactRoles")
    def contact_roles(self) -> Optional[Sequence[str]]:
        """
        Contact roles to send the budget notification to when the threshold is exceeded.
        """
        return pulumi.get(self, "contact_roles")


