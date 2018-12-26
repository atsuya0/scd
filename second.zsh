function second() {
  [[ $1 == 'change' ]] \
    && eval cd $(command second $@ || echo '.') 2> /dev/null \
    || command second $@
}
