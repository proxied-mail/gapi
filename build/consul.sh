consul-template --once \
  -consul-addr=consul \
  -template="env.ctmpl:../env:exit 0"
