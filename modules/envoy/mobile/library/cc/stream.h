#pragma once

#include <vector>

#include "library/common/types/c_types.h"
#include "request_headers.h"
#include "request_trailers.h"
#include "stream_callbacks.h"

namespace Envoy {
namespace Platform {

class Engine;
using EngineSharedPtr = std::shared_ptr<Engine>;

class Stream {
public:
  Stream(envoy_engine_t engine_handle, envoy_stream_t handle);

  Stream& sendHeaders(RequestHeadersSharedPtr headers, bool end_stream);
  Stream& sendData(envoy_data data);
  Stream& readData(size_t bytes_to_read);
  void close(RequestTrailersSharedPtr trailers);
  void close(envoy_data data);
  void cancel();

private:
  envoy_engine_t engine_handle_;
  envoy_stream_t handle_;
};

using StreamSharedPtr = std::shared_ptr<Stream>;

} // namespace Platform
} // namespace Envoy
