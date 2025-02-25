// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.docker.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceTaskSpecContainerSpecPrivilegesSeLinuxContextArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceTaskSpecContainerSpecPrivilegesSeLinuxContextArgs Empty = new ServiceTaskSpecContainerSpecPrivilegesSeLinuxContextArgs();

    @Import(name="disable")
    private @Nullable Output<Boolean> disable;

    public Optional<Output<Boolean>> disable() {
        return Optional.ofNullable(this.disable);
    }

    @Import(name="level")
    private @Nullable Output<String> level;

    public Optional<Output<String>> level() {
        return Optional.ofNullable(this.level);
    }

    @Import(name="role")
    private @Nullable Output<String> role;

    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    @Import(name="type")
    private @Nullable Output<String> type;

    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    @Import(name="user")
    private @Nullable Output<String> user;

    public Optional<Output<String>> user() {
        return Optional.ofNullable(this.user);
    }

    private ServiceTaskSpecContainerSpecPrivilegesSeLinuxContextArgs() {}

    private ServiceTaskSpecContainerSpecPrivilegesSeLinuxContextArgs(ServiceTaskSpecContainerSpecPrivilegesSeLinuxContextArgs $) {
        this.disable = $.disable;
        this.level = $.level;
        this.role = $.role;
        this.type = $.type;
        this.user = $.user;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceTaskSpecContainerSpecPrivilegesSeLinuxContextArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceTaskSpecContainerSpecPrivilegesSeLinuxContextArgs $;

        public Builder() {
            $ = new ServiceTaskSpecContainerSpecPrivilegesSeLinuxContextArgs();
        }

        public Builder(ServiceTaskSpecContainerSpecPrivilegesSeLinuxContextArgs defaults) {
            $ = new ServiceTaskSpecContainerSpecPrivilegesSeLinuxContextArgs(Objects.requireNonNull(defaults));
        }

        public Builder disable(@Nullable Output<Boolean> disable) {
            $.disable = disable;
            return this;
        }

        public Builder disable(Boolean disable) {
            return disable(Output.of(disable));
        }

        public Builder level(@Nullable Output<String> level) {
            $.level = level;
            return this;
        }

        public Builder level(String level) {
            return level(Output.of(level));
        }

        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        public Builder role(String role) {
            return role(Output.of(role));
        }

        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public Builder user(@Nullable Output<String> user) {
            $.user = user;
            return this;
        }

        public Builder user(String user) {
            return user(Output.of(user));
        }

        public ServiceTaskSpecContainerSpecPrivilegesSeLinuxContextArgs build() {
            return $;
        }
    }

}
