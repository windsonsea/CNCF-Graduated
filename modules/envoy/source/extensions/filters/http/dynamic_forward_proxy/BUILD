load(
    "//bazel:envoy_build_system.bzl",
    "envoy_cc_extension",
    "envoy_cc_library",
    "envoy_extension_package",
)

licenses(["notice"])  # Apache 2

envoy_extension_package()

envoy_cc_library(
    name = "proxy_filter_lib",
    srcs = ["proxy_filter.cc"],
    hdrs = ["proxy_filter.h"],
    deps = [
        "//envoy/http:filter_interface",
        "//source/common/http:header_utility_lib",
        "//source/common/network:filter_state_proxy_info_lib",
        "//source/common/stream_info:upstream_address_lib",
        "//source/common/upstream:upstream_includes",
        "//source/extensions/clusters/dynamic_forward_proxy:cluster",
        "//source/extensions/common/dynamic_forward_proxy:dns_cache_interface",
        "//source/extensions/filters/http/common:pass_through_filter_lib",
        "@envoy_api//envoy/config/cluster/v3:pkg_cc_proto",
        "@envoy_api//envoy/config/core/v3:pkg_cc_proto",
        "@envoy_api//envoy/extensions/clusters/dynamic_forward_proxy/v3:pkg_cc_proto",
        "@envoy_api//envoy/extensions/filters/http/dynamic_forward_proxy/v3:pkg_cc_proto",
    ],
)

envoy_cc_extension(
    name = "config",
    srcs = ["config.cc"],
    hdrs = ["config.h"],
    deps = [
        "//envoy/registry",
        "//envoy/server:filter_config_interface",
        "//source/extensions/common/dynamic_forward_proxy:dns_cache_manager_impl",
        "//source/extensions/filters/http/common:factory_base_lib",
        "//source/extensions/filters/http/dynamic_forward_proxy:proxy_filter_lib",
        "@envoy_api//envoy/extensions/filters/http/dynamic_forward_proxy/v3:pkg_cc_proto",
    ],
)
