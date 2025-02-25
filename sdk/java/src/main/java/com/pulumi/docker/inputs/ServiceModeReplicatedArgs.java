// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceModeReplicatedArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceModeReplicatedArgs Empty = new ServiceModeReplicatedArgs();

    @Import(name="replicas")
    private @Nullable Output<Integer> replicas;

    public Optional<Output<Integer>> replicas() {
        return Optional.ofNullable(this.replicas);
    }

    private ServiceModeReplicatedArgs() {}

    private ServiceModeReplicatedArgs(ServiceModeReplicatedArgs $) {
        this.replicas = $.replicas;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceModeReplicatedArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceModeReplicatedArgs $;

        public Builder() {
            $ = new ServiceModeReplicatedArgs();
        }

        public Builder(ServiceModeReplicatedArgs defaults) {
            $ = new ServiceModeReplicatedArgs(Objects.requireNonNull(defaults));
        }

        public Builder replicas(@Nullable Output<Integer> replicas) {
            $.replicas = replicas;
            return this;
        }

        public Builder replicas(Integer replicas) {
            return replicas(Output.of(replicas));
        }

        public ServiceModeReplicatedArgs build() {
            return $;
        }
    }

}
