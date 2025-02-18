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
    'GetManagedClusterResult',
    'AwaitableGetManagedClusterResult',
    'get_managed_cluster',
    'get_managed_cluster_output',
]

@pulumi.output_type
class GetManagedClusterResult:
    """
    The manged cluster resource
    """
    def __init__(__self__, addon_features=None, admin_password=None, admin_user_name=None, azure_active_directory=None, client_connection_port=None, clients=None, cluster_certificate_thumbprint=None, cluster_code_version=None, cluster_id=None, cluster_state=None, dns_name=None, etag=None, fabric_settings=None, fqdn=None, http_gateway_connection_port=None, id=None, load_balancing_rules=None, location=None, name=None, provisioning_state=None, sku=None, tags=None, type=None):
        if addon_features and not isinstance(addon_features, list):
            raise TypeError("Expected argument 'addon_features' to be a list")
        pulumi.set(__self__, "addon_features", addon_features)
        if admin_password and not isinstance(admin_password, str):
            raise TypeError("Expected argument 'admin_password' to be a str")
        pulumi.set(__self__, "admin_password", admin_password)
        if admin_user_name and not isinstance(admin_user_name, str):
            raise TypeError("Expected argument 'admin_user_name' to be a str")
        pulumi.set(__self__, "admin_user_name", admin_user_name)
        if azure_active_directory and not isinstance(azure_active_directory, dict):
            raise TypeError("Expected argument 'azure_active_directory' to be a dict")
        pulumi.set(__self__, "azure_active_directory", azure_active_directory)
        if client_connection_port and not isinstance(client_connection_port, int):
            raise TypeError("Expected argument 'client_connection_port' to be a int")
        pulumi.set(__self__, "client_connection_port", client_connection_port)
        if clients and not isinstance(clients, list):
            raise TypeError("Expected argument 'clients' to be a list")
        pulumi.set(__self__, "clients", clients)
        if cluster_certificate_thumbprint and not isinstance(cluster_certificate_thumbprint, str):
            raise TypeError("Expected argument 'cluster_certificate_thumbprint' to be a str")
        pulumi.set(__self__, "cluster_certificate_thumbprint", cluster_certificate_thumbprint)
        if cluster_code_version and not isinstance(cluster_code_version, str):
            raise TypeError("Expected argument 'cluster_code_version' to be a str")
        pulumi.set(__self__, "cluster_code_version", cluster_code_version)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if cluster_state and not isinstance(cluster_state, str):
            raise TypeError("Expected argument 'cluster_state' to be a str")
        pulumi.set(__self__, "cluster_state", cluster_state)
        if dns_name and not isinstance(dns_name, str):
            raise TypeError("Expected argument 'dns_name' to be a str")
        pulumi.set(__self__, "dns_name", dns_name)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if fabric_settings and not isinstance(fabric_settings, list):
            raise TypeError("Expected argument 'fabric_settings' to be a list")
        pulumi.set(__self__, "fabric_settings", fabric_settings)
        if fqdn and not isinstance(fqdn, str):
            raise TypeError("Expected argument 'fqdn' to be a str")
        pulumi.set(__self__, "fqdn", fqdn)
        if http_gateway_connection_port and not isinstance(http_gateway_connection_port, int):
            raise TypeError("Expected argument 'http_gateway_connection_port' to be a int")
        pulumi.set(__self__, "http_gateway_connection_port", http_gateway_connection_port)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if load_balancing_rules and not isinstance(load_balancing_rules, list):
            raise TypeError("Expected argument 'load_balancing_rules' to be a list")
        pulumi.set(__self__, "load_balancing_rules", load_balancing_rules)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="addonFeatures")
    def addon_features(self) -> Optional[Sequence[str]]:
        """
        client certificates for the cluster.
        """
        return pulumi.get(self, "addon_features")

    @property
    @pulumi.getter(name="adminPassword")
    def admin_password(self) -> Optional[str]:
        """
        vm admin user password.
        """
        return pulumi.get(self, "admin_password")

    @property
    @pulumi.getter(name="adminUserName")
    def admin_user_name(self) -> str:
        """
        vm admin user name.
        """
        return pulumi.get(self, "admin_user_name")

    @property
    @pulumi.getter(name="azureActiveDirectory")
    def azure_active_directory(self) -> Optional['outputs.AzureActiveDirectoryResponse']:
        """
        Azure active directory.
        """
        return pulumi.get(self, "azure_active_directory")

    @property
    @pulumi.getter(name="clientConnectionPort")
    def client_connection_port(self) -> Optional[int]:
        """
        The port used for client connections to the cluster.
        """
        return pulumi.get(self, "client_connection_port")

    @property
    @pulumi.getter
    def clients(self) -> Optional[Sequence['outputs.ClientCertificateResponse']]:
        """
        client certificates for the cluster.
        """
        return pulumi.get(self, "clients")

    @property
    @pulumi.getter(name="clusterCertificateThumbprint")
    def cluster_certificate_thumbprint(self) -> str:
        """
        The cluster certificate thumbprint used node to node communication.
        """
        return pulumi.get(self, "cluster_certificate_thumbprint")

    @property
    @pulumi.getter(name="clusterCodeVersion")
    def cluster_code_version(self) -> Optional[str]:
        """
        The Service Fabric runtime version of the cluster. This property can only by set the user when **upgradeMode** is set to 'Manual'. To get list of available Service Fabric versions for new clusters use [ClusterVersion API](./ClusterVersion.md). To get the list of available version for existing clusters use **availableClusterVersions**.
        """
        return pulumi.get(self, "cluster_code_version")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        A service generated unique identifier for the cluster resource.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterState")
    def cluster_state(self) -> str:
        """
        The current state of the cluster.
        """
        return pulumi.get(self, "cluster_state")

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> str:
        """
        The cluster dns name.
        """
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        Azure resource etag.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="fabricSettings")
    def fabric_settings(self) -> Optional[Sequence['outputs.SettingsSectionDescriptionResponse']]:
        """
        The list of custom fabric settings to configure the cluster.
        """
        return pulumi.get(self, "fabric_settings")

    @property
    @pulumi.getter
    def fqdn(self) -> str:
        """
        the cluster Fully qualified domain name.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter(name="httpGatewayConnectionPort")
    def http_gateway_connection_port(self) -> Optional[int]:
        """
        The port used for http connections to the cluster.
        """
        return pulumi.get(self, "http_gateway_connection_port")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Azure resource identifier.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="loadBalancingRules")
    def load_balancing_rules(self) -> Optional[Sequence['outputs.LoadBalancingRuleResponse']]:
        """
        Describes load balancing rules.
        """
        return pulumi.get(self, "load_balancing_rules")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Azure resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Azure resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the managed cluster resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.SkuResponse']:
        """
        The sku of the managed cluster
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Azure resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Azure resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetManagedClusterResult(GetManagedClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetManagedClusterResult(
            addon_features=self.addon_features,
            admin_password=self.admin_password,
            admin_user_name=self.admin_user_name,
            azure_active_directory=self.azure_active_directory,
            client_connection_port=self.client_connection_port,
            clients=self.clients,
            cluster_certificate_thumbprint=self.cluster_certificate_thumbprint,
            cluster_code_version=self.cluster_code_version,
            cluster_id=self.cluster_id,
            cluster_state=self.cluster_state,
            dns_name=self.dns_name,
            etag=self.etag,
            fabric_settings=self.fabric_settings,
            fqdn=self.fqdn,
            http_gateway_connection_port=self.http_gateway_connection_port,
            id=self.id,
            load_balancing_rules=self.load_balancing_rules,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            sku=self.sku,
            tags=self.tags,
            type=self.type)


def get_managed_cluster(cluster_name: Optional[str] = None,
                        resource_group_name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetManagedClusterResult:
    """
    The manged cluster resource

    API Version: 2020-01-01-preview.


    :param str cluster_name: The name of the cluster resource.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:servicefabric:getManagedCluster', __args__, opts=opts, typ=GetManagedClusterResult).value

    return AwaitableGetManagedClusterResult(
        addon_features=__ret__.addon_features,
        admin_password=__ret__.admin_password,
        admin_user_name=__ret__.admin_user_name,
        azure_active_directory=__ret__.azure_active_directory,
        client_connection_port=__ret__.client_connection_port,
        clients=__ret__.clients,
        cluster_certificate_thumbprint=__ret__.cluster_certificate_thumbprint,
        cluster_code_version=__ret__.cluster_code_version,
        cluster_id=__ret__.cluster_id,
        cluster_state=__ret__.cluster_state,
        dns_name=__ret__.dns_name,
        etag=__ret__.etag,
        fabric_settings=__ret__.fabric_settings,
        fqdn=__ret__.fqdn,
        http_gateway_connection_port=__ret__.http_gateway_connection_port,
        id=__ret__.id,
        load_balancing_rules=__ret__.load_balancing_rules,
        location=__ret__.location,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        sku=__ret__.sku,
        tags=__ret__.tags,
        type=__ret__.type)


@_utilities.lift_output_func(get_managed_cluster)
def get_managed_cluster_output(cluster_name: Optional[pulumi.Input[str]] = None,
                               resource_group_name: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetManagedClusterResult]:
    """
    The manged cluster resource

    API Version: 2020-01-01-preview.


    :param str cluster_name: The name of the cluster resource.
    :param str resource_group_name: The name of the resource group.
    """
    ...
