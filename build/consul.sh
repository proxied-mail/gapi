consul-template --once \
  -consul-addr=pm-consul:8500 \
  -template="/app/build/env.ctmpl:/app/.env:exit 0"
