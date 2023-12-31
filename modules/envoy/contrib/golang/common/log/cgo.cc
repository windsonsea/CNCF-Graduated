#include "contrib/golang/common/log/cgo.h"

namespace Envoy {
namespace Extensions {
namespace Common {
namespace Golang {

/* FilterLogger */
void FilterLogger::log(uint32_t level, absl::string_view message) const {
  switch (static_cast<spdlog::level::level_enum>(level)) {
  case spdlog::level::trace:
    ENVOY_LOG(trace, "{}", message);
    return;
  case spdlog::level::debug:
    ENVOY_LOG(debug, "{}", message);
    return;
  case spdlog::level::info:
    ENVOY_LOG(info, "{}", message);
    return;
  case spdlog::level::warn:
    ENVOY_LOG(warn, "{}", message);
    return;
  case spdlog::level::err:
    ENVOY_LOG(error, "{}", message);
    return;
  case spdlog::level::critical:
    ENVOY_LOG(critical, "{}", message);
    return;
  case spdlog::level::off:
    // means not logging
    return;
  case spdlog::level::n_levels:
    PANIC("not implemented");
  }

  ENVOY_LOG(error, "undefined log level {} with message '{}'", level, message);

  PANIC_DUE_TO_CORRUPT_ENUM;
}

uint32_t FilterLogger::level() const { return static_cast<uint32_t>(ENVOY_LOGGER().level()); }

const FilterLogger& getFilterLogger() { CONSTRUCT_ON_FIRST_USE(FilterLogger); }

// The returned absl::string_view only refer to the GoString, won't copy the string content into
// C++, should not use it after the current cgo call returns.
absl::string_view referGoString(void* str) {
  if (str == nullptr) {
    return "";
  }
  auto go_str = reinterpret_cast<GoString*>(str);
  return {go_str->p, static_cast<size_t>(go_str->n)};
}

#ifdef __cplusplus
extern "C" {
#endif

void envoyGoFilterLog(uint32_t level, void* message) {
  auto mesg = referGoString(message);
  getFilterLogger().log(level, mesg);
}

uint32_t envoyGoFilterLogLevel() { return getFilterLogger().level(); }

#ifdef __cplusplus
}
#endif
} // namespace Golang
} // namespace Common
} // namespace Extensions
} // namespace Envoy
