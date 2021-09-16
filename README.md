# scd
Change the working directory with a second name.  

`$ cd ~/dotfiles`  
/home/user/dotfiles  
`$ scd register -n dot`  
`$ cd ~`  
/home/user  
`$ scd change` Choose interactively.  
/home/user/dotfiles  

`$ scd change dot`  
/home/user/dotfiles  
`$ cd config/config/zsh`  
/home/user/dotfiles/config/config/zsh  
`$ scd register -s`  
`$ cd ~`  
/home/user  
`$ scd change -s` Choose interactively.  
/home/user/dotfiles/config/config/zsh  

# Available Commands
## register
Attach a second name to a target path.

|option|description|default|
|-|-|-|
|--name(-n)|a second name|the current working directory name|
|--path(-p)|a target path|the path of the current working directory|
|--sub(-s)|can register a sub directory|

## change
Change the current working directory with a second name.
|option|
|-|
|--sub(-s)|

## remove
Remove a second name.

## list
List a second name and a target path in JSON format.

|option|description|
|-|-|
|--name(-n)|List only a name|
|--path(-p)|List only a path|

## show
Display a target path by a second name.

## init
Initialize data.

# Setup
## Load the zsh script
Add this line to your .zshrc.
```shell
source <(scd script)
```
## Data file
Data is saved in $XDG_DATA_HOME/scd.  
If $XDG_DATA_HOME is either not set or empty, a default equal to $HOME/.local/share should be used.  
It can also be specified by $SCD_DATA_PATH.
