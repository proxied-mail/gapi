consul-template --once \
  -consul-addr=localhost:8502 \
  -template="../config/env.ctmpl:../config/.env:exit 0"
