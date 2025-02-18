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
    'DiskResponse',
    'HardwareProfileResponse',
    'IpAddressResponse',
    'NetworkProfileResponse',
    'OSProfileResponse',
    'SAPSystemIDResponse',
    'StorageProfileResponse',
]

@pulumi.output_type
class DiskResponse(dict):
    """
    Specifies the disk information for the HANA instance
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "diskSizeGB":
            suggest = "disk_size_gb"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DiskResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DiskResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DiskResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 lun: int,
                 disk_size_gb: Optional[int] = None,
                 name: Optional[str] = None):
        """
        Specifies the disk information for the HANA instance
        :param int lun: Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and therefore must be unique for each data disk attached to a VM.
        :param int disk_size_gb: Specifies the size of an empty data disk in gigabytes.
        :param str name: The disk name.
        """
        pulumi.set(__self__, "lun", lun)
        if disk_size_gb is not None:
            pulumi.set(__self__, "disk_size_gb", disk_size_gb)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def lun(self) -> int:
        """
        Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and therefore must be unique for each data disk attached to a VM.
        """
        return pulumi.get(self, "lun")

    @property
    @pulumi.getter(name="diskSizeGB")
    def disk_size_gb(self) -> Optional[int]:
        """
        Specifies the size of an empty data disk in gigabytes.
        """
        return pulumi.get(self, "disk_size_gb")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The disk name.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class HardwareProfileResponse(dict):
    """
    Specifies the hardware settings for the HANA instance.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "hanaInstanceSize":
            suggest = "hana_instance_size"
        elif key == "hardwareType":
            suggest = "hardware_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in HardwareProfileResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        HardwareProfileResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        HardwareProfileResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 hana_instance_size: str,
                 hardware_type: str):
        """
        Specifies the hardware settings for the HANA instance.
        :param str hana_instance_size: Specifies the HANA instance SKU.
        :param str hardware_type: Name of the hardware type (vendor and/or their product name)
        """
        pulumi.set(__self__, "hana_instance_size", hana_instance_size)
        pulumi.set(__self__, "hardware_type", hardware_type)

    @property
    @pulumi.getter(name="hanaInstanceSize")
    def hana_instance_size(self) -> str:
        """
        Specifies the HANA instance SKU.
        """
        return pulumi.get(self, "hana_instance_size")

    @property
    @pulumi.getter(name="hardwareType")
    def hardware_type(self) -> str:
        """
        Name of the hardware type (vendor and/or their product name)
        """
        return pulumi.get(self, "hardware_type")


@pulumi.output_type
class IpAddressResponse(dict):
    """
    Specifies the IP address of the network interface.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "ipAddress":
            suggest = "ip_address"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IpAddressResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IpAddressResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IpAddressResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 ip_address: Optional[str] = None):
        """
        Specifies the IP address of the network interface.
        :param str ip_address: Specifies the IP address of the network interface.
        """
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[str]:
        """
        Specifies the IP address of the network interface.
        """
        return pulumi.get(self, "ip_address")


@pulumi.output_type
class NetworkProfileResponse(dict):
    """
    Specifies the network settings for the HANA instance disks.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "circuitId":
            suggest = "circuit_id"
        elif key == "networkInterfaces":
            suggest = "network_interfaces"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NetworkProfileResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NetworkProfileResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NetworkProfileResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 circuit_id: str,
                 network_interfaces: Optional[Sequence['outputs.IpAddressResponse']] = None):
        """
        Specifies the network settings for the HANA instance disks.
        :param str circuit_id: Specifies the circuit id for connecting to express route.
        :param Sequence['IpAddressResponse'] network_interfaces: Specifies the network interfaces for the HANA instance.
        """
        pulumi.set(__self__, "circuit_id", circuit_id)
        if network_interfaces is not None:
            pulumi.set(__self__, "network_interfaces", network_interfaces)

    @property
    @pulumi.getter(name="circuitId")
    def circuit_id(self) -> str:
        """
        Specifies the circuit id for connecting to express route.
        """
        return pulumi.get(self, "circuit_id")

    @property
    @pulumi.getter(name="networkInterfaces")
    def network_interfaces(self) -> Optional[Sequence['outputs.IpAddressResponse']]:
        """
        Specifies the network interfaces for the HANA instance.
        """
        return pulumi.get(self, "network_interfaces")


@pulumi.output_type
class OSProfileResponse(dict):
    """
    Specifies the operating system settings for the HANA instance.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "osType":
            suggest = "os_type"
        elif key == "computerName":
            suggest = "computer_name"
        elif key == "sshPublicKey":
            suggest = "ssh_public_key"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in OSProfileResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        OSProfileResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        OSProfileResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 os_type: str,
                 version: str,
                 computer_name: Optional[str] = None,
                 ssh_public_key: Optional[str] = None):
        """
        Specifies the operating system settings for the HANA instance.
        :param str os_type: This property allows you to specify the type of the OS.
        :param str version: Specifies version of operating system.
        :param str computer_name: Specifies the host OS name of the HANA instance.
        :param str ssh_public_key: Specifies the SSH public key used to access the operating system.
        """
        pulumi.set(__self__, "os_type", os_type)
        pulumi.set(__self__, "version", version)
        if computer_name is not None:
            pulumi.set(__self__, "computer_name", computer_name)
        if ssh_public_key is not None:
            pulumi.set(__self__, "ssh_public_key", ssh_public_key)

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> str:
        """
        This property allows you to specify the type of the OS.
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        Specifies version of operating system.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="computerName")
    def computer_name(self) -> Optional[str]:
        """
        Specifies the host OS name of the HANA instance.
        """
        return pulumi.get(self, "computer_name")

    @property
    @pulumi.getter(name="sshPublicKey")
    def ssh_public_key(self) -> Optional[str]:
        """
        Specifies the SSH public key used to access the operating system.
        """
        return pulumi.get(self, "ssh_public_key")


@pulumi.output_type
class SAPSystemIDResponse(dict):
    """
    Specifies information related to a SAP system ID
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "memoryAllocation":
            suggest = "memory_allocation"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SAPSystemIDResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SAPSystemIDResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SAPSystemIDResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 memory_allocation: str,
                 gid: Optional[str] = None,
                 sid: Optional[str] = None,
                 uid: Optional[str] = None,
                 username: Optional[str] = None):
        """
        Specifies information related to a SAP system ID
        :param str memory_allocation: Percent of memory to allocate to this SID.
        :param str gid: Group ID of the HANA database user.
        :param str sid: SAP system ID as database identifier.
        :param str uid: User ID of the HANA database user.
        :param str username: Name of the HANA database user.
        """
        pulumi.set(__self__, "memory_allocation", memory_allocation)
        if gid is not None:
            pulumi.set(__self__, "gid", gid)
        if sid is not None:
            pulumi.set(__self__, "sid", sid)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="memoryAllocation")
    def memory_allocation(self) -> str:
        """
        Percent of memory to allocate to this SID.
        """
        return pulumi.get(self, "memory_allocation")

    @property
    @pulumi.getter
    def gid(self) -> Optional[str]:
        """
        Group ID of the HANA database user.
        """
        return pulumi.get(self, "gid")

    @property
    @pulumi.getter
    def sid(self) -> Optional[str]:
        """
        SAP system ID as database identifier.
        """
        return pulumi.get(self, "sid")

    @property
    @pulumi.getter
    def uid(self) -> Optional[str]:
        """
        User ID of the HANA database user.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter
    def username(self) -> Optional[str]:
        """
        Name of the HANA database user.
        """
        return pulumi.get(self, "username")


@pulumi.output_type
class StorageProfileResponse(dict):
    """
    Specifies the storage settings for the HANA instance disks.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "nfsIpAddress":
            suggest = "nfs_ip_address"
        elif key == "hanaSids":
            suggest = "hana_sids"
        elif key == "osDisks":
            suggest = "os_disks"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StorageProfileResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StorageProfileResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StorageProfileResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 nfs_ip_address: str,
                 hana_sids: Optional[Sequence['outputs.SAPSystemIDResponse']] = None,
                 os_disks: Optional[Sequence['outputs.DiskResponse']] = None):
        """
        Specifies the storage settings for the HANA instance disks.
        :param str nfs_ip_address: IP Address to connect to storage.
        :param Sequence['SAPSystemIDResponse'] hana_sids: Specifies information related to SAP system IDs for the hana instance.
        :param Sequence['DiskResponse'] os_disks: Specifies information about the operating system disk used by the hana instance.
        """
        pulumi.set(__self__, "nfs_ip_address", nfs_ip_address)
        if hana_sids is not None:
            pulumi.set(__self__, "hana_sids", hana_sids)
        if os_disks is not None:
            pulumi.set(__self__, "os_disks", os_disks)

    @property
    @pulumi.getter(name="nfsIpAddress")
    def nfs_ip_address(self) -> str:
        """
        IP Address to connect to storage.
        """
        return pulumi.get(self, "nfs_ip_address")

    @property
    @pulumi.getter(name="hanaSids")
    def hana_sids(self) -> Optional[Sequence['outputs.SAPSystemIDResponse']]:
        """
        Specifies information related to SAP system IDs for the hana instance.
        """
        return pulumi.get(self, "hana_sids")

    @property
    @pulumi.getter(name="osDisks")
    def os_disks(self) -> Optional[Sequence['outputs.DiskResponse']]:
        """
        Specifies information about the operating system disk used by the hana instance.
        """
        return pulumi.get(self, "os_disks")


