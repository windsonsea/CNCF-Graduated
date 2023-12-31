#include "source/extensions/upstreams/http/tcp/config.h"

#include "source/extensions/upstreams/http/tcp/upstream_request.h"

namespace Envoy {
namespace Extensions {
namespace Upstreams {
namespace Http {
namespace Tcp {

Router::GenericConnPoolPtr TcpGenericConnPoolFactory::createGenericConnPool(
    Upstream::ThreadLocalCluster& thread_local_cluster,
    Router::GenericConnPoolFactory::UpstreamProtocol, const Router::RouteEntry& route_entry,
    absl::optional<Envoy::Http::Protocol>, Upstream::LoadBalancerContext* ctx) const {
  auto ret = std::make_unique<TcpConnPool>(thread_local_cluster, route_entry, ctx);
  return (ret->valid() ? std::move(ret) : nullptr);
}

REGISTER_FACTORY(TcpGenericConnPoolFactory, Router::GenericConnPoolFactory);

} // namespace Tcp
} // namespace Http
} // namespace Upstreams
} // namespace Extensions
} // namespace Envoy
