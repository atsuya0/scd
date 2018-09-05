function second() {
  local second=${GOPATH}/bin/second
  [[ $1 == 'change' ]] \
    && eval cd $(${second} $@ || echo '.') \
    || ${second} $@
}
