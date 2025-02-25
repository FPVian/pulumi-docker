// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.docker.SecretArgs;
import com.pulumi.docker.Utilities;
import com.pulumi.docker.inputs.SecretState;
import com.pulumi.docker.outputs.SecretLabel;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * #!/bin/bash Docker secret cannot be imported as the secret data, once set, is never exposed again.
 * 
 */
@ResourceType(type="docker:index/secret:Secret")
public class Secret extends com.pulumi.resources.CustomResource {
    /**
     * Base64-url-safe-encoded secret data
     * 
     */
    @Export(name="data", refs={String.class}, tree="[0]")
    private Output<String> data;

    /**
     * @return Base64-url-safe-encoded secret data
     * 
     */
    public Output<String> data() {
        return this.data;
    }
    /**
     * User-defined key/value metadata
     * 
     */
    @Export(name="labels", refs={List.class,SecretLabel.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SecretLabel>> labels;

    /**
     * @return User-defined key/value metadata
     * 
     */
    public Output<Optional<List<SecretLabel>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * User-defined name of the secret
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return User-defined name of the secret
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Secret(String name) {
        this(name, SecretArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Secret(String name, SecretArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Secret(String name, SecretArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("docker:index/secret:Secret", name, args == null ? SecretArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Secret(String name, Output<String> id, @Nullable SecretState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("docker:index/secret:Secret", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "data"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Secret get(String name, Output<String> id, @Nullable SecretState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Secret(name, id, state, options);
    }
}
