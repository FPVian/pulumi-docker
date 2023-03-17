// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Docker
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("docker");

        private static readonly __Value<string?> _caMaterial = new __Value<string?>(() => __config.Get("caMaterial"));
        /// <summary>
        /// PEM-encoded content of Docker host CA certificate
        /// </summary>
        public static string? CaMaterial
        {
            get => _caMaterial.Get();
            set => _caMaterial.Set(value);
        }

        private static readonly __Value<string?> _certMaterial = new __Value<string?>(() => __config.Get("certMaterial"));
        /// <summary>
        /// PEM-encoded content of Docker client certificate
        /// </summary>
        public static string? CertMaterial
        {
            get => _certMaterial.Get();
            set => _certMaterial.Set(value);
        }

        private static readonly __Value<string?> _certPath = new __Value<string?>(() => __config.Get("certPath"));
        /// <summary>
        /// Path to directory with Docker TLS config
        /// </summary>
        public static string? CertPath
        {
            get => _certPath.Get();
            set => _certPath.Set(value);
        }

        private static readonly __Value<string?> _host = new __Value<string?>(() => __config.Get("host") ?? Utilities.GetEnv("DOCKER_HOST"));
        /// <summary>
        /// The Docker daemon address
        /// </summary>
        public static string? Host
        {
            get => _host.Get();
            set => _host.Set(value);
        }

        private static readonly __Value<string?> _keyMaterial = new __Value<string?>(() => __config.Get("keyMaterial"));
        /// <summary>
        /// PEM-encoded content of Docker client private key
        /// </summary>
        public static string? KeyMaterial
        {
            get => _keyMaterial.Get();
            set => _keyMaterial.Set(value);
        }

        private static readonly __Value<ImmutableArray<Pulumi.Docker.Config.Types.RegistryAuth>> _registryAuth = new __Value<ImmutableArray<Pulumi.Docker.Config.Types.RegistryAuth>>(() => __config.GetObject<ImmutableArray<Pulumi.Docker.Config.Types.RegistryAuth>>("registryAuth"));
        public static ImmutableArray<Pulumi.Docker.Config.Types.RegistryAuth> RegistryAuth
        {
            get => _registryAuth.Get();
            set => _registryAuth.Set(value);
        }

        private static readonly __Value<ImmutableArray<string>> _sshOpts = new __Value<ImmutableArray<string>>(() => __config.GetObject<ImmutableArray<string>>("sshOpts"));
        /// <summary>
        /// Additional SSH option flags to be appended when using `ssh://` protocol
        /// </summary>
        public static ImmutableArray<string> SshOpts
        {
            get => _sshOpts.Get();
            set => _sshOpts.Set(value);
        }

        public static class Types
        {

             public class RegistryAuth
             {
                public string Address { get; set; }
                public bool? AuthDisabled { get; set; }
                public string? ConfigFile { get; set; } = null!;
                public string? ConfigFileContent { get; set; } = null!;
                public string? Password { get; set; } = null!;
                public string? Username { get; set; } = null!;
            }
        }
    }
}
