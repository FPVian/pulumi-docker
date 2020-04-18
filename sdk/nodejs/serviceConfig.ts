// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages the configuration of a Docker service in a swarm.
 * 
 * 
 * ## Basic
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 * 
 * // Creates a config
 * const fooConfig = new docker.ServiceConfig("fooConfig", {
 *     data: "ewogICJzZXJIfQo=",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-docker/blob/master/website/docs/r/config.html.markdown.
 */
export class ServiceConfig extends pulumi.CustomResource {
    /**
     * Get an existing ServiceConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceConfigState, opts?: pulumi.CustomResourceOptions): ServiceConfig {
        return new ServiceConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'docker:index/serviceConfig:ServiceConfig';

    /**
     * Returns true if the given object is an instance of ServiceConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceConfig.__pulumiType;
    }

    /**
     * The base64 encoded data of the config.
     */
    public readonly data!: pulumi.Output<string>;
    /**
     * The name of the Docker config.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a ServiceConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceConfigArgs | ServiceConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServiceConfigState | undefined;
            inputs["data"] = state ? state.data : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ServiceConfigArgs | undefined;
            if (!args || args.data === undefined) {
                throw new Error("Missing required property 'data'");
            }
            inputs["data"] = args ? args.data : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ServiceConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceConfig resources.
 */
export interface ServiceConfigState {
    /**
     * The base64 encoded data of the config.
     */
    readonly data?: pulumi.Input<string>;
    /**
     * The name of the Docker config.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceConfig resource.
 */
export interface ServiceConfigArgs {
    /**
     * The base64 encoded data of the config.
     */
    readonly data: pulumi.Input<string>;
    /**
     * The name of the Docker config.
     */
    readonly name?: pulumi.Input<string>;
}
