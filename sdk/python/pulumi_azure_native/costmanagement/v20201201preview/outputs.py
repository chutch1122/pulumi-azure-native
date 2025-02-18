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
    'CommonExportPropertiesResponse',
    'ErrorDetailsResponse',
    'ExportDatasetConfigurationResponse',
    'ExportDatasetResponse',
    'ExportDefinitionResponse',
    'ExportDeliveryDestinationResponse',
    'ExportDeliveryInfoResponse',
    'ExportExecutionListResultResponse',
    'ExportExecutionResponse',
    'ExportRecurrencePeriodResponse',
    'ExportScheduleResponse',
    'ExportTimePeriodResponse',
]

@pulumi.output_type
class CommonExportPropertiesResponse(dict):
    """
    The common properties of the export.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "deliveryInfo":
            suggest = "delivery_info"
        elif key == "nextRunTimeEstimate":
            suggest = "next_run_time_estimate"
        elif key == "runHistory":
            suggest = "run_history"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in CommonExportPropertiesResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        CommonExportPropertiesResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        CommonExportPropertiesResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 definition: 'outputs.ExportDefinitionResponse',
                 delivery_info: 'outputs.ExportDeliveryInfoResponse',
                 next_run_time_estimate: str,
                 format: Optional[str] = None,
                 run_history: Optional['outputs.ExportExecutionListResultResponse'] = None):
        """
        The common properties of the export.
        :param 'ExportDefinitionResponse' definition: Has the definition for the export.
        :param 'ExportDeliveryInfoResponse' delivery_info: Has delivery information for the export.
        :param str next_run_time_estimate: If the export has an active schedule, provides an estimate of the next execution time.
        :param str format: The format of the export being delivered. Currently only 'Csv' is supported.
        :param 'ExportExecutionListResultResponse' run_history: If requested, has the most recent execution history for the export.
        """
        pulumi.set(__self__, "definition", definition)
        pulumi.set(__self__, "delivery_info", delivery_info)
        pulumi.set(__self__, "next_run_time_estimate", next_run_time_estimate)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if run_history is not None:
            pulumi.set(__self__, "run_history", run_history)

    @property
    @pulumi.getter
    def definition(self) -> 'outputs.ExportDefinitionResponse':
        """
        Has the definition for the export.
        """
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter(name="deliveryInfo")
    def delivery_info(self) -> 'outputs.ExportDeliveryInfoResponse':
        """
        Has delivery information for the export.
        """
        return pulumi.get(self, "delivery_info")

    @property
    @pulumi.getter(name="nextRunTimeEstimate")
    def next_run_time_estimate(self) -> str:
        """
        If the export has an active schedule, provides an estimate of the next execution time.
        """
        return pulumi.get(self, "next_run_time_estimate")

    @property
    @pulumi.getter
    def format(self) -> Optional[str]:
        """
        The format of the export being delivered. Currently only 'Csv' is supported.
        """
        return pulumi.get(self, "format")

    @property
    @pulumi.getter(name="runHistory")
    def run_history(self) -> Optional['outputs.ExportExecutionListResultResponse']:
        """
        If requested, has the most recent execution history for the export.
        """
        return pulumi.get(self, "run_history")


@pulumi.output_type
class ErrorDetailsResponse(dict):
    """
    The details of the error.
    """
    def __init__(__self__, *,
                 code: str,
                 message: str):
        """
        The details of the error.
        :param str code: Error code.
        :param str message: Error message indicating why the operation failed.
        """
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "message", message)

    @property
    @pulumi.getter
    def code(self) -> str:
        """
        Error code.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def message(self) -> str:
        """
        Error message indicating why the operation failed.
        """
        return pulumi.get(self, "message")


@pulumi.output_type
class ExportDatasetConfigurationResponse(dict):
    """
    The export dataset configuration. Allows columns to be selected for the export. If not provided then the export will include all available columns.
    """
    def __init__(__self__, *,
                 columns: Optional[Sequence[str]] = None):
        """
        The export dataset configuration. Allows columns to be selected for the export. If not provided then the export will include all available columns.
        :param Sequence[str] columns: Array of column names to be included in the export. If not provided then the export will include all available columns. The available columns can vary by customer channel (see examples).
        """
        if columns is not None:
            pulumi.set(__self__, "columns", columns)

    @property
    @pulumi.getter
    def columns(self) -> Optional[Sequence[str]]:
        """
        Array of column names to be included in the export. If not provided then the export will include all available columns. The available columns can vary by customer channel (see examples).
        """
        return pulumi.get(self, "columns")


@pulumi.output_type
class ExportDatasetResponse(dict):
    """
    The definition for data in the export.
    """
    def __init__(__self__, *,
                 configuration: Optional['outputs.ExportDatasetConfigurationResponse'] = None,
                 granularity: Optional[str] = None):
        """
        The definition for data in the export.
        :param 'ExportDatasetConfigurationResponse' configuration: The export dataset configuration.
        :param str granularity: The granularity of rows in the export. Currently only 'Daily' is supported.
        """
        if configuration is not None:
            pulumi.set(__self__, "configuration", configuration)
        if granularity is not None:
            pulumi.set(__self__, "granularity", granularity)

    @property
    @pulumi.getter
    def configuration(self) -> Optional['outputs.ExportDatasetConfigurationResponse']:
        """
        The export dataset configuration.
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter
    def granularity(self) -> Optional[str]:
        """
        The granularity of rows in the export. Currently only 'Daily' is supported.
        """
        return pulumi.get(self, "granularity")


@pulumi.output_type
class ExportDefinitionResponse(dict):
    """
    The definition of an export.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dataSet":
            suggest = "data_set"
        elif key == "timePeriod":
            suggest = "time_period"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ExportDefinitionResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ExportDefinitionResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ExportDefinitionResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 timeframe: str,
                 type: str,
                 data_set: Optional['outputs.ExportDatasetResponse'] = None,
                 time_period: Optional['outputs.ExportTimePeriodResponse'] = None):
        """
        The definition of an export.
        :param str timeframe: The time frame for pulling data for the export. If custom, then a specific time period must be provided.
        :param str type: The type of the export. Note that 'Usage' is equivalent to 'ActualCost' and is applicable to exports that do not yet provide data for charges or amortization for service reservations.
        :param 'ExportDatasetResponse' data_set: The definition for data in the export.
        :param 'ExportTimePeriodResponse' time_period: Has time period for pulling data for the export.
        """
        pulumi.set(__self__, "timeframe", timeframe)
        pulumi.set(__self__, "type", type)
        if data_set is not None:
            pulumi.set(__self__, "data_set", data_set)
        if time_period is not None:
            pulumi.set(__self__, "time_period", time_period)

    @property
    @pulumi.getter
    def timeframe(self) -> str:
        """
        The time frame for pulling data for the export. If custom, then a specific time period must be provided.
        """
        return pulumi.get(self, "timeframe")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the export. Note that 'Usage' is equivalent to 'ActualCost' and is applicable to exports that do not yet provide data for charges or amortization for service reservations.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="dataSet")
    def data_set(self) -> Optional['outputs.ExportDatasetResponse']:
        """
        The definition for data in the export.
        """
        return pulumi.get(self, "data_set")

    @property
    @pulumi.getter(name="timePeriod")
    def time_period(self) -> Optional['outputs.ExportTimePeriodResponse']:
        """
        Has time period for pulling data for the export.
        """
        return pulumi.get(self, "time_period")


@pulumi.output_type
class ExportDeliveryDestinationResponse(dict):
    """
    This represents the blob storage account location where exports of costs will be delivered. There are two ways to configure the destination. The approach recommended for most customers is to specify the resourceId of the storage account. This requires a one-time registration of the account's subscription with the Microsoft.CostManagementExports resource provider in order to give Azure Cost Management services access to the storage. When creating an export in the Azure portal this registration is performed automatically but API users may need to register the subscription explicitly (for more information see https://docs.microsoft.com/en-us/azure/azure-resource-manager/resource-manager-supported-services ). Another way to configure the destination is available ONLY to Partners with a Microsoft Partner Agreement plan who are global admins of their billing account. These Partners, instead of specifying the resourceId of a storage account, can specify the storage account name along with a SAS token for the account. This allows exports of costs to a storage account in any tenant. The SAS token should be created for the blob service with Service/Container/Object resource types and with Read/Write/Delete/List/Add/Create permissions (for more information see https://docs.microsoft.com/en-us/azure/cost-management-billing/costs/export-cost-data-storage-account-sas-key ).
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "resourceId":
            suggest = "resource_id"
        elif key == "rootFolderPath":
            suggest = "root_folder_path"
        elif key == "storageAccount":
            suggest = "storage_account"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ExportDeliveryDestinationResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ExportDeliveryDestinationResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ExportDeliveryDestinationResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 container: str,
                 resource_id: Optional[str] = None,
                 root_folder_path: Optional[str] = None,
                 storage_account: Optional[str] = None):
        """
        This represents the blob storage account location where exports of costs will be delivered. There are two ways to configure the destination. The approach recommended for most customers is to specify the resourceId of the storage account. This requires a one-time registration of the account's subscription with the Microsoft.CostManagementExports resource provider in order to give Azure Cost Management services access to the storage. When creating an export in the Azure portal this registration is performed automatically but API users may need to register the subscription explicitly (for more information see https://docs.microsoft.com/en-us/azure/azure-resource-manager/resource-manager-supported-services ). Another way to configure the destination is available ONLY to Partners with a Microsoft Partner Agreement plan who are global admins of their billing account. These Partners, instead of specifying the resourceId of a storage account, can specify the storage account name along with a SAS token for the account. This allows exports of costs to a storage account in any tenant. The SAS token should be created for the blob service with Service/Container/Object resource types and with Read/Write/Delete/List/Add/Create permissions (for more information see https://docs.microsoft.com/en-us/azure/cost-management-billing/costs/export-cost-data-storage-account-sas-key ).
        :param str container: The name of the container where exports will be uploaded. If the container does not exist it will be created.
        :param str resource_id: The resource id of the storage account where exports will be delivered. This is not required if a sasToken and storageAccount are specified.
        :param str root_folder_path: The name of the directory where exports will be uploaded.
        :param str storage_account: The storage account where exports will be uploaded. For a restricted set of Azure customers this together with sasToken can be specified instead of resourceId.
        """
        pulumi.set(__self__, "container", container)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if root_folder_path is not None:
            pulumi.set(__self__, "root_folder_path", root_folder_path)
        if storage_account is not None:
            pulumi.set(__self__, "storage_account", storage_account)

    @property
    @pulumi.getter
    def container(self) -> str:
        """
        The name of the container where exports will be uploaded. If the container does not exist it will be created.
        """
        return pulumi.get(self, "container")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[str]:
        """
        The resource id of the storage account where exports will be delivered. This is not required if a sasToken and storageAccount are specified.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="rootFolderPath")
    def root_folder_path(self) -> Optional[str]:
        """
        The name of the directory where exports will be uploaded.
        """
        return pulumi.get(self, "root_folder_path")

    @property
    @pulumi.getter(name="storageAccount")
    def storage_account(self) -> Optional[str]:
        """
        The storage account where exports will be uploaded. For a restricted set of Azure customers this together with sasToken can be specified instead of resourceId.
        """
        return pulumi.get(self, "storage_account")


@pulumi.output_type
class ExportDeliveryInfoResponse(dict):
    """
    The delivery information associated with a export.
    """
    def __init__(__self__, *,
                 destination: 'outputs.ExportDeliveryDestinationResponse'):
        """
        The delivery information associated with a export.
        :param 'ExportDeliveryDestinationResponse' destination: Has destination for the export being delivered.
        """
        pulumi.set(__self__, "destination", destination)

    @property
    @pulumi.getter
    def destination(self) -> 'outputs.ExportDeliveryDestinationResponse':
        """
        Has destination for the export being delivered.
        """
        return pulumi.get(self, "destination")


@pulumi.output_type
class ExportExecutionListResultResponse(dict):
    """
    Result of listing the execution history of an export.
    """
    def __init__(__self__, *,
                 value: Sequence['outputs.ExportExecutionResponse']):
        """
        Result of listing the execution history of an export.
        :param Sequence['ExportExecutionResponse'] value: A list of export executions.
        """
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> Sequence['outputs.ExportExecutionResponse']:
        """
        A list of export executions.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class ExportExecutionResponse(dict):
    """
    An export execution.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "eTag":
            suggest = "e_tag"
        elif key == "executionType":
            suggest = "execution_type"
        elif key == "fileName":
            suggest = "file_name"
        elif key == "processingEndTime":
            suggest = "processing_end_time"
        elif key == "processingStartTime":
            suggest = "processing_start_time"
        elif key == "runSettings":
            suggest = "run_settings"
        elif key == "submittedBy":
            suggest = "submitted_by"
        elif key == "submittedTime":
            suggest = "submitted_time"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ExportExecutionResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ExportExecutionResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ExportExecutionResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 id: str,
                 name: str,
                 type: str,
                 e_tag: Optional[str] = None,
                 error: Optional['outputs.ErrorDetailsResponse'] = None,
                 execution_type: Optional[str] = None,
                 file_name: Optional[str] = None,
                 processing_end_time: Optional[str] = None,
                 processing_start_time: Optional[str] = None,
                 run_settings: Optional['outputs.CommonExportPropertiesResponse'] = None,
                 status: Optional[str] = None,
                 submitted_by: Optional[str] = None,
                 submitted_time: Optional[str] = None):
        """
        An export execution.
        :param str id: Resource Id.
        :param str name: Resource name.
        :param str type: Resource type.
        :param str e_tag: eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        :param 'ErrorDetailsResponse' error: The details of any error.
        :param str execution_type: The type of the export execution.
        :param str file_name: The name of the exported file.
        :param str processing_end_time: The time when the export execution finished.
        :param str processing_start_time: The time when export was picked up to be executed.
        :param 'CommonExportPropertiesResponse' run_settings: The export settings that were in effect for this execution.
        :param str status: The last known status of the export execution.
        :param str submitted_by: The identifier for the entity that executed the export. For OnDemand executions it is the user email. For scheduled executions it is 'System'.
        :param str submitted_time: The time when export was queued to be executed.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        if e_tag is not None:
            pulumi.set(__self__, "e_tag", e_tag)
        if error is not None:
            pulumi.set(__self__, "error", error)
        if execution_type is not None:
            pulumi.set(__self__, "execution_type", execution_type)
        if file_name is not None:
            pulumi.set(__self__, "file_name", file_name)
        if processing_end_time is not None:
            pulumi.set(__self__, "processing_end_time", processing_end_time)
        if processing_start_time is not None:
            pulumi.set(__self__, "processing_start_time", processing_start_time)
        if run_settings is not None:
            pulumi.set(__self__, "run_settings", run_settings)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if submitted_by is not None:
            pulumi.set(__self__, "submitted_by", submitted_by)
        if submitted_time is not None:
            pulumi.set(__self__, "submitted_time", submitted_time)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Resource Id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="eTag")
    def e_tag(self) -> Optional[str]:
        """
        eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        """
        return pulumi.get(self, "e_tag")

    @property
    @pulumi.getter
    def error(self) -> Optional['outputs.ErrorDetailsResponse']:
        """
        The details of any error.
        """
        return pulumi.get(self, "error")

    @property
    @pulumi.getter(name="executionType")
    def execution_type(self) -> Optional[str]:
        """
        The type of the export execution.
        """
        return pulumi.get(self, "execution_type")

    @property
    @pulumi.getter(name="fileName")
    def file_name(self) -> Optional[str]:
        """
        The name of the exported file.
        """
        return pulumi.get(self, "file_name")

    @property
    @pulumi.getter(name="processingEndTime")
    def processing_end_time(self) -> Optional[str]:
        """
        The time when the export execution finished.
        """
        return pulumi.get(self, "processing_end_time")

    @property
    @pulumi.getter(name="processingStartTime")
    def processing_start_time(self) -> Optional[str]:
        """
        The time when export was picked up to be executed.
        """
        return pulumi.get(self, "processing_start_time")

    @property
    @pulumi.getter(name="runSettings")
    def run_settings(self) -> Optional['outputs.CommonExportPropertiesResponse']:
        """
        The export settings that were in effect for this execution.
        """
        return pulumi.get(self, "run_settings")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The last known status of the export execution.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="submittedBy")
    def submitted_by(self) -> Optional[str]:
        """
        The identifier for the entity that executed the export. For OnDemand executions it is the user email. For scheduled executions it is 'System'.
        """
        return pulumi.get(self, "submitted_by")

    @property
    @pulumi.getter(name="submittedTime")
    def submitted_time(self) -> Optional[str]:
        """
        The time when export was queued to be executed.
        """
        return pulumi.get(self, "submitted_time")


@pulumi.output_type
class ExportRecurrencePeriodResponse(dict):
    """
    The start and end date for recurrence schedule.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "from":
            suggest = "from_"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ExportRecurrencePeriodResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ExportRecurrencePeriodResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ExportRecurrencePeriodResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 from_: str,
                 to: Optional[str] = None):
        """
        The start and end date for recurrence schedule.
        :param str from_: The start date of recurrence.
        :param str to: The end date of recurrence.
        """
        pulumi.set(__self__, "from_", from_)
        if to is not None:
            pulumi.set(__self__, "to", to)

    @property
    @pulumi.getter(name="from")
    def from_(self) -> str:
        """
        The start date of recurrence.
        """
        return pulumi.get(self, "from_")

    @property
    @pulumi.getter
    def to(self) -> Optional[str]:
        """
        The end date of recurrence.
        """
        return pulumi.get(self, "to")


@pulumi.output_type
class ExportScheduleResponse(dict):
    """
    The schedule associated with the export.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "recurrencePeriod":
            suggest = "recurrence_period"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ExportScheduleResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ExportScheduleResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ExportScheduleResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 recurrence: Optional[str] = None,
                 recurrence_period: Optional['outputs.ExportRecurrencePeriodResponse'] = None,
                 status: Optional[str] = None):
        """
        The schedule associated with the export.
        :param str recurrence: The schedule recurrence.
        :param 'ExportRecurrencePeriodResponse' recurrence_period: Has start and end date of the recurrence. The start date must be in future. If present, the end date must be greater than start date.
        :param str status: The status of the export's schedule. If 'Inactive', the export's schedule is paused.
        """
        if recurrence is not None:
            pulumi.set(__self__, "recurrence", recurrence)
        if recurrence_period is not None:
            pulumi.set(__self__, "recurrence_period", recurrence_period)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def recurrence(self) -> Optional[str]:
        """
        The schedule recurrence.
        """
        return pulumi.get(self, "recurrence")

    @property
    @pulumi.getter(name="recurrencePeriod")
    def recurrence_period(self) -> Optional['outputs.ExportRecurrencePeriodResponse']:
        """
        Has start and end date of the recurrence. The start date must be in future. If present, the end date must be greater than start date.
        """
        return pulumi.get(self, "recurrence_period")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the export's schedule. If 'Inactive', the export's schedule is paused.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class ExportTimePeriodResponse(dict):
    """
    The date range for data in the export. This should only be specified with timeFrame set to 'Custom'. The maximum date range is 3 months.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "from":
            suggest = "from_"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ExportTimePeriodResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ExportTimePeriodResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ExportTimePeriodResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 from_: str,
                 to: str):
        """
        The date range for data in the export. This should only be specified with timeFrame set to 'Custom'. The maximum date range is 3 months.
        :param str from_: The start date for export data.
        :param str to: The end date for export data.
        """
        pulumi.set(__self__, "from_", from_)
        pulumi.set(__self__, "to", to)

    @property
    @pulumi.getter(name="from")
    def from_(self) -> str:
        """
        The start date for export data.
        """
        return pulumi.get(self, "from_")

    @property
    @pulumi.getter
    def to(self) -> str:
        """
        The end date for export data.
        """
        return pulumi.get(self, "to")


