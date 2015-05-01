# nix-tools
Various unix / linux small command line tools

## Navigation stack
Inspired by the csh *pushd* builtin commands, the navigation stack commands allows :
* pushing as many paths as you want in the stack
* showing the stack content at any moment
* jumping to any item of the stack, provided its number
* replacing a given item in the stack by another path
* resetting the path (= clearing the stack)
* stack persistence over current session (logout, login again, the stack is still there :)

## enhanced "cd"
This command accepts a path with only the first characters of each directory level provided. It then tries to autocomplete the path. If there is only one path matching, it jumps directly to this path. Otherwise it lists the paths that match the provided "pattern" and allows to choose which one to jump to.



