# Available Commands
## register
Attach the second name to the target path.

|option|description|default|
|-|-|-|
|--name(-n)|the second name|the current working directory name|
|--path(-p)|the target path|the path of the current working directory|

## change
Change the current working directory with the second name.
## remove
Remove the second name.
## list
List the second name and the target path in JSON format.

|option|description|
|-|-|
|--name(-n)|List only the name|
|--path(-p)|List only the path|

## show
Display the target path by the second name.
## init
Initialize the data.

# Setup
Specify the path to save the data using $SECOND_LIST_PATH.

It is $XDG_CONFIG_HOME/second/list.json in the default.
