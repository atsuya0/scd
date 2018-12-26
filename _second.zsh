function _second() {
  function sub_commands() {
    _values 'Commands' \
      'change' \
      'register' \
      'list' \
      'delete' \
      'init'
  }

  _arguments -C \
    '(-h --help)'{-h,--help}'[show help]' \
    '1: :sub_commands' \
    '*:: :->args'

  case "${state}" in
    (args)
      case "${words[1]}" in
        (register)
          _arguments \
            '(-n --name)'{-n,--name}'[Second name]' \
            '(-p --path)'{-p,--path}'[Target path]'
        ;;
        (change)
          _values \
            'Second names' \
            $(second list --name)
        ;;
        (list)
          _arguments \
            '(-n --name)'{-n,--name}'[Second name]' \
            '(-p --path)'{-p,--path}'[Target path]'
        ;;
        (delete)
          _values \
            'Second names' \
            $(second list --name)
        ;;
        (init)
        ;;
      esac
  esac
}
compdef _second second
