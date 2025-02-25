// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceTaskSpecLogDriverArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceTaskSpecLogDriverArgs Empty = new ServiceTaskSpecLogDriverArgs();

    /**
     * Name of the service
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the service
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    @Import(name="options")
    private @Nullable Output<Map<String,String>> options;

    public Optional<Output<Map<String,String>>> options() {
        return Optional.ofNullable(this.options);
    }

    private ServiceTaskSpecLogDriverArgs() {}

    private ServiceTaskSpecLogDriverArgs(ServiceTaskSpecLogDriverArgs $) {
        this.name = $.name;
        this.options = $.options;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceTaskSpecLogDriverArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceTaskSpecLogDriverArgs $;

        public Builder() {
            $ = new ServiceTaskSpecLogDriverArgs();
        }

        public Builder(ServiceTaskSpecLogDriverArgs defaults) {
            $ = new ServiceTaskSpecLogDriverArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the service
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the service
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder options(@Nullable Output<Map<String,String>> options) {
            $.options = options;
            return this;
        }

        public Builder options(Map<String,String> options) {
            return options(Output.of(options));
        }

        public ServiceTaskSpecLogDriverArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
