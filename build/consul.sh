consul-template --once \
  -consul-addr=consul \
  -template="/go/src/pmgo/build/env.ctmpl:/go/src/pmgo/config/env:exit 0"
