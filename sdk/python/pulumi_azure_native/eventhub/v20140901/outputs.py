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
    'SkuResponse',
]

@pulumi.output_type
class SkuResponse(dict):
    """
    SKU parameters supplied to the create Namespace operation
    """
    def __init__(__self__, *,
                 tier: str,
                 capacity: Optional[int] = None,
                 name: Optional[str] = None):
        """
        SKU parameters supplied to the create Namespace operation
        :param str tier: The billing tier of this particular SKU.
        :param int capacity: The Event Hubs throughput units.
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
        The Event Hubs throughput units.
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of this SKU.
        """
        return pulumi.get(self, "name")


