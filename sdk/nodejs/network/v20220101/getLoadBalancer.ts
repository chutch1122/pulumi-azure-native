// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * LoadBalancer resource.
 */
export function getLoadBalancer(args: GetLoadBalancerArgs, opts?: pulumi.InvokeOptions): Promise<GetLoadBalancerResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("azure-native:network/v20220101:getLoadBalancer", {
        "expand": args.expand,
        "loadBalancerName": args.loadBalancerName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetLoadBalancerArgs {
    /**
     * Expands referenced resources.
     */
    expand?: string;
    /**
     * The name of the load balancer.
     */
    loadBalancerName: string;
    /**
     * The name of the resource group.
     */
    resourceGroupName: string;
}

/**
 * LoadBalancer resource.
 */
export interface GetLoadBalancerResult {
    /**
     * Collection of backend address pools used by a load balancer.
     */
    readonly backendAddressPools?: outputs.network.v20220101.BackendAddressPoolResponse[];
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * The extended location of the load balancer.
     */
    readonly extendedLocation?: outputs.network.v20220101.ExtendedLocationResponse;
    /**
     * Object representing the frontend IPs to be used for the load balancer.
     */
    readonly frontendIPConfigurations?: outputs.network.v20220101.FrontendIPConfigurationResponse[];
    /**
     * Resource ID.
     */
    readonly id?: string;
    /**
     * Defines an port range to be used by inbound NAT Pools. Inbound NAT pools are used to define a range of NAT ports to be used by a VMSS cluster. After the creation of an inbound NAT pool, individual inbound NAT rules are automatically created for every VM in a VMSS cluster.  Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are associated with VMSS, while inbound NAT rules are associated with individual VMs.
     */
    readonly inboundNatPools?: outputs.network.v20220101.InboundNatPoolResponse[];
    /**
     * collection of inbound NAT Rules used by a load balancer. An inbound NAT rule is used to forward traffic from a load balancer frontend to one or more instances in the backend pool. Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are associated with VMSS, while inbound NAT rules are associated with individual VMs.
     */
    readonly inboundNatRules?: outputs.network.v20220101.InboundNatRuleResponse[];
    /**
     * Object collection representing the load balancing rules Gets the provisioning.
     */
    readonly loadBalancingRules?: outputs.network.v20220101.LoadBalancingRuleResponse[];
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The outbound rules.
     */
    readonly outboundRules?: outputs.network.v20220101.OutboundRuleResponse[];
    /**
     * Collection of probe objects used in the load balancer.
     */
    readonly probes?: outputs.network.v20220101.ProbeResponse[];
    /**
     * The provisioning state of the load balancer resource.
     */
    readonly provisioningState: string;
    /**
     * The resource GUID property of the load balancer resource.
     */
    readonly resourceGuid: string;
    /**
     * The load balancer SKU.
     */
    readonly sku?: outputs.network.v20220101.LoadBalancerSkuResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}

export function getLoadBalancerOutput(args: GetLoadBalancerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLoadBalancerResult> {
    return pulumi.output(args).apply(a => getLoadBalancer(a, opts))
}

export interface GetLoadBalancerOutputArgs {
    /**
     * Expands referenced resources.
     */
    expand?: pulumi.Input<string>;
    /**
     * The name of the load balancer.
     */
    loadBalancerName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    resourceGroupName: pulumi.Input<string>;
}
