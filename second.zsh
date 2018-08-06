function second() {
  local second=${GOPATH}/bin/second
  [[ $1 == 'change' ]] \
    && cd $(${second} $@ || echo '.') \
    || ${second} $@
}
