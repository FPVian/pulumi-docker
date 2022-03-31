// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * `docker.Network` provides details about a specific Docker Network.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const main = pulumi.output(docker.getNetwork({
 *     name: "main",
 * }));
 * ```
 */
export function getNetwork(args: GetNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("docker:index/getNetwork:getNetwork", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetwork.
 */
export interface GetNetworkArgs {
    /**
     * The name of the Docker network.
     */
    name: string;
}

/**
 * A collection of values returned by getNetwork.
 */
export interface GetNetworkResult {
    /**
     * The driver of the Docker network. Possible values are `bridge`, `host`, `overlay`, `macvlan`. See [network docs](https://docs.docker.com/network/#network-drivers) for more details.
     */
    readonly driver: string;
    /**
     * The ID of this resource.
     */
    readonly id: string;
    /**
     * If `true`, the network is internal.
     */
    readonly internal: boolean;
    /**
     * The IPAM configuration options
     */
    readonly ipamConfigs: outputs.GetNetworkIpamConfig[];
    /**
     * The name of the Docker network.
     */
    readonly name: string;
    /**
     * Only available with bridge networks. See [bridge options docs](https://docs.docker.com/engine/reference/commandline/network_create/#bridge-driver-options) for more details.
     */
    readonly options: {[key: string]: any};
    /**
     * Scope of the network. One of `swarm`, `global`, or `local`.
     */
    readonly scope: string;
}

export function getNetworkOutput(args: GetNetworkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkResult> {
    return pulumi.output(args).apply(a => getNetwork(a, opts))
}

/**
 * A collection of arguments for invoking getNetwork.
 */
export interface GetNetworkOutputArgs {
    /**
     * The name of the Docker network.
     */
    name: pulumi.Input<string>;
}
