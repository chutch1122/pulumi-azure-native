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
    'ActiveDirectoryArgs',
    'ExportPolicyRuleArgs',
    'VolumePropertiesExportPolicyArgs',
]

@pulumi.input_type
class ActiveDirectoryArgs:
    def __init__(__self__, *,
                 active_directory_id: Optional[pulumi.Input[str]] = None,
                 d_ns: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 organizational_unit: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 s_mb_server_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Active Directory
        :param pulumi.Input[str] active_directory_id: Id of the Active Directory
        :param pulumi.Input[str] d_ns: Comma separated list of DNS server IP addresses for the Active Directory domain
        :param pulumi.Input[str] domain: Name of the Active Directory domain
        :param pulumi.Input[str] organizational_unit: The Organizational Unit (OU) within the Windows Active Directory
        :param pulumi.Input[str] password: Plain text password of Active Directory domain administrator
        :param pulumi.Input[str] s_mb_server_name: NetBIOS name of the SMB server. This name will be registered as a computer account in the AD and used to mount volumes
        :param pulumi.Input[str] status: Status of the Active Directory
        :param pulumi.Input[str] username: Username of Active Directory domain administrator
        """
        if active_directory_id is not None:
            pulumi.set(__self__, "active_directory_id", active_directory_id)
        if d_ns is not None:
            pulumi.set(__self__, "d_ns", d_ns)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if organizational_unit is not None:
            pulumi.set(__self__, "organizational_unit", organizational_unit)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if s_mb_server_name is not None:
            pulumi.set(__self__, "s_mb_server_name", s_mb_server_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="activeDirectoryId")
    def active_directory_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of the Active Directory
        """
        return pulumi.get(self, "active_directory_id")

    @active_directory_id.setter
    def active_directory_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "active_directory_id", value)

    @property
    @pulumi.getter(name="dNS")
    def d_ns(self) -> Optional[pulumi.Input[str]]:
        """
        Comma separated list of DNS server IP addresses for the Active Directory domain
        """
        return pulumi.get(self, "d_ns")

    @d_ns.setter
    def d_ns(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "d_ns", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Active Directory domain
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="organizationalUnit")
    def organizational_unit(self) -> Optional[pulumi.Input[str]]:
        """
        The Organizational Unit (OU) within the Windows Active Directory
        """
        return pulumi.get(self, "organizational_unit")

    @organizational_unit.setter
    def organizational_unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organizational_unit", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Plain text password of Active Directory domain administrator
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="sMBServerName")
    def s_mb_server_name(self) -> Optional[pulumi.Input[str]]:
        """
        NetBIOS name of the SMB server. This name will be registered as a computer account in the AD and used to mount volumes
        """
        return pulumi.get(self, "s_mb_server_name")

    @s_mb_server_name.setter
    def s_mb_server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s_mb_server_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the Active Directory
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        Username of Active Directory domain administrator
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class ExportPolicyRuleArgs:
    def __init__(__self__, *,
                 allowed_clients: Optional[pulumi.Input[str]] = None,
                 cifs: Optional[pulumi.Input[bool]] = None,
                 nfsv3: Optional[pulumi.Input[bool]] = None,
                 nfsv4: Optional[pulumi.Input[bool]] = None,
                 rule_index: Optional[pulumi.Input[int]] = None,
                 unix_read_only: Optional[pulumi.Input[bool]] = None,
                 unix_read_write: Optional[pulumi.Input[bool]] = None):
        """
        Volume Export Policy Rule
        :param pulumi.Input[str] allowed_clients: Client ingress specification as comma separated string with IPv4 CIDRs, IPv4 host addresses and host names
        :param pulumi.Input[bool] cifs: Allows CIFS protocol
        :param pulumi.Input[bool] nfsv3: Allows NFSv3 protocol
        :param pulumi.Input[bool] nfsv4: Allows NFSv4 protocol
        :param pulumi.Input[int] rule_index: Order index
        :param pulumi.Input[bool] unix_read_only: Read only access
        :param pulumi.Input[bool] unix_read_write: Read and write access
        """
        if allowed_clients is not None:
            pulumi.set(__self__, "allowed_clients", allowed_clients)
        if cifs is not None:
            pulumi.set(__self__, "cifs", cifs)
        if nfsv3 is not None:
            pulumi.set(__self__, "nfsv3", nfsv3)
        if nfsv4 is not None:
            pulumi.set(__self__, "nfsv4", nfsv4)
        if rule_index is not None:
            pulumi.set(__self__, "rule_index", rule_index)
        if unix_read_only is not None:
            pulumi.set(__self__, "unix_read_only", unix_read_only)
        if unix_read_write is not None:
            pulumi.set(__self__, "unix_read_write", unix_read_write)

    @property
    @pulumi.getter(name="allowedClients")
    def allowed_clients(self) -> Optional[pulumi.Input[str]]:
        """
        Client ingress specification as comma separated string with IPv4 CIDRs, IPv4 host addresses and host names
        """
        return pulumi.get(self, "allowed_clients")

    @allowed_clients.setter
    def allowed_clients(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allowed_clients", value)

    @property
    @pulumi.getter
    def cifs(self) -> Optional[pulumi.Input[bool]]:
        """
        Allows CIFS protocol
        """
        return pulumi.get(self, "cifs")

    @cifs.setter
    def cifs(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "cifs", value)

    @property
    @pulumi.getter
    def nfsv3(self) -> Optional[pulumi.Input[bool]]:
        """
        Allows NFSv3 protocol
        """
        return pulumi.get(self, "nfsv3")

    @nfsv3.setter
    def nfsv3(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "nfsv3", value)

    @property
    @pulumi.getter
    def nfsv4(self) -> Optional[pulumi.Input[bool]]:
        """
        Allows NFSv4 protocol
        """
        return pulumi.get(self, "nfsv4")

    @nfsv4.setter
    def nfsv4(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "nfsv4", value)

    @property
    @pulumi.getter(name="ruleIndex")
    def rule_index(self) -> Optional[pulumi.Input[int]]:
        """
        Order index
        """
        return pulumi.get(self, "rule_index")

    @rule_index.setter
    def rule_index(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rule_index", value)

    @property
    @pulumi.getter(name="unixReadOnly")
    def unix_read_only(self) -> Optional[pulumi.Input[bool]]:
        """
        Read only access
        """
        return pulumi.get(self, "unix_read_only")

    @unix_read_only.setter
    def unix_read_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "unix_read_only", value)

    @property
    @pulumi.getter(name="unixReadWrite")
    def unix_read_write(self) -> Optional[pulumi.Input[bool]]:
        """
        Read and write access
        """
        return pulumi.get(self, "unix_read_write")

    @unix_read_write.setter
    def unix_read_write(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "unix_read_write", value)


@pulumi.input_type
class VolumePropertiesExportPolicyArgs:
    def __init__(__self__, *,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['ExportPolicyRuleArgs']]]] = None):
        """
        Export policy rule
        """
        if rules is not None:
            pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ExportPolicyRuleArgs']]]]:
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ExportPolicyRuleArgs']]]]):
        pulumi.set(self, "rules", value)


