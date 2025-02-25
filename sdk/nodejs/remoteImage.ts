// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * <!-- Bug: Type and Name are switched -->
 * Pulls a Docker image to a given Docker host from a Docker Registry.
 *  This resource will *not* pull new layers of the image automatically unless used in conjunction with docker.RegistryImage data source to update the `pullTriggers` field.
 *
 * ## Example Usage
 * ### Basic
 *
 * Finds and downloads the latest `ubuntu:precise` image but does not check
 * for further updates of the image
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const ubuntu = new docker.RemoteImage("ubuntu", {name: "ubuntu:precise"});
 * ```
 * ### Dynamic updates
 *
 * To be able to update an image dynamically when the `sha256` sum changes,
 * you need to use it in combination with `docker.RegistryImage` as follows:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as docker from "@pulumi/docker";
 *
 * const ubuntuRegistryImage = docker.getRegistryImage({
 *     name: "ubuntu:precise",
 * });
 * const ubuntuRemoteImage = new docker.RemoteImage("ubuntuRemoteImage", {
 *     name: ubuntuRegistryImage.then(ubuntuRegistryImage => ubuntuRegistryImage.name),
 *     pullTriggers: [ubuntuRegistryImage.then(ubuntuRegistryImage => ubuntuRegistryImage.sha256Digest)],
 * });
 * ```
 */
export class RemoteImage extends pulumi.CustomResource {
    /**
     * Get an existing RemoteImage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RemoteImageState, opts?: pulumi.CustomResourceOptions): RemoteImage {
        return new RemoteImage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'docker:index/remoteImage:RemoteImage';

    /**
     * Returns true if the given object is an instance of RemoteImage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RemoteImage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RemoteImage.__pulumiType;
    }

    /**
     * Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
     */
    public readonly build!: pulumi.Output<outputs.RemoteImageBuild | undefined>;
    /**
     * If true, then the image is removed forcibly when the resource is destroyed.
     */
    public readonly forceRemove!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the image (as seen when executing `docker inspect` on the image). Can be used to reference the image via its ID in other resources.
     */
    public /*out*/ readonly imageId!: pulumi.Output<string>;
    /**
     * If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
     */
    public readonly keepLocally!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the Docker image, including any tags or SHA256 repo digests.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The platform to use when pulling the image. Defaults to the platform of the current machine.
     */
    public readonly platform!: pulumi.Output<string | undefined>;
    /**
     * List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
     */
    public readonly pullTriggers!: pulumi.Output<string[] | undefined>;
    /**
     * The image sha256 digest in the form of `repo[:tag]@sha256:<hash>`.
     */
    public /*out*/ readonly repoDigest!: pulumi.Output<string>;
    /**
     * A map of arbitrary strings that, when changed, will force the `docker.RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
     */
    public readonly triggers!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a RemoteImage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RemoteImageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RemoteImageArgs | RemoteImageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RemoteImageState | undefined;
            resourceInputs["build"] = state ? state.build : undefined;
            resourceInputs["forceRemove"] = state ? state.forceRemove : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["keepLocally"] = state ? state.keepLocally : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["platform"] = state ? state.platform : undefined;
            resourceInputs["pullTriggers"] = state ? state.pullTriggers : undefined;
            resourceInputs["repoDigest"] = state ? state.repoDigest : undefined;
            resourceInputs["triggers"] = state ? state.triggers : undefined;
        } else {
            const args = argsOrState as RemoteImageArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["build"] = args ? args.build : undefined;
            resourceInputs["forceRemove"] = args ? args.forceRemove : undefined;
            resourceInputs["keepLocally"] = args ? args.keepLocally : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["platform"] = args ? args.platform : undefined;
            resourceInputs["pullTriggers"] = args ? args.pullTriggers : undefined;
            resourceInputs["triggers"] = args ? args.triggers : undefined;
            resourceInputs["imageId"] = undefined /*out*/;
            resourceInputs["repoDigest"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RemoteImage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RemoteImage resources.
 */
export interface RemoteImageState {
    /**
     * Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
     */
    build?: pulumi.Input<inputs.RemoteImageBuild>;
    /**
     * If true, then the image is removed forcibly when the resource is destroyed.
     */
    forceRemove?: pulumi.Input<boolean>;
    /**
     * The ID of the image (as seen when executing `docker inspect` on the image). Can be used to reference the image via its ID in other resources.
     */
    imageId?: pulumi.Input<string>;
    /**
     * If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
     */
    keepLocally?: pulumi.Input<boolean>;
    /**
     * The name of the Docker image, including any tags or SHA256 repo digests.
     */
    name?: pulumi.Input<string>;
    /**
     * The platform to use when pulling the image. Defaults to the platform of the current machine.
     */
    platform?: pulumi.Input<string>;
    /**
     * List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
     */
    pullTriggers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The image sha256 digest in the form of `repo[:tag]@sha256:<hash>`.
     */
    repoDigest?: pulumi.Input<string>;
    /**
     * A map of arbitrary strings that, when changed, will force the `docker.RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
     */
    triggers?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a RemoteImage resource.
 */
export interface RemoteImageArgs {
    /**
     * Configuration to build an image. Please see [docker build command reference](https://docs.docker.com/engine/reference/commandline/build/#options) too.
     */
    build?: pulumi.Input<inputs.RemoteImageBuild>;
    /**
     * If true, then the image is removed forcibly when the resource is destroyed.
     */
    forceRemove?: pulumi.Input<boolean>;
    /**
     * If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.
     */
    keepLocally?: pulumi.Input<boolean>;
    /**
     * The name of the Docker image, including any tags or SHA256 repo digests.
     */
    name: pulumi.Input<string>;
    /**
     * The platform to use when pulling the image. Defaults to the platform of the current machine.
     */
    platform?: pulumi.Input<string>;
    /**
     * List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the docker*registry*image.
     */
    pullTriggers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of arbitrary strings that, when changed, will force the `docker.RemoteImage` resource to be replaced. This can be used to rebuild an image when contents of source code folders change
     */
    triggers?: pulumi.Input<{[key: string]: any}>;
}
