# Spacer

Workspaces for your terminal.

## Usage

### Path

The path command outputs a context dependant path name, useful in
command prompts to clean up the current working directory.

``` bash
spacer path
```

### Change

The change command moves to the location if possible.

``` bash
spacer change <location>
```

### Create

The create command can be used for locations that need to be created.
This can be useful for conditional new directory creation, for instance.

``` bash
spacer create <location>
```

It can also be used as part of change, where it will be run if the
selected location does consider the current location to me movable.

``` bash
spacer change --create <location>
```

### Layout

The layout command will create tmux panes.

``` bash
spacer layout <location>
```

It can be used as part of change:

``` bash
spacer change --layout <location>
```

## Bash

In order to support changing directories for commands, there is a bash
wrapper.

To use it, just source it in your `.bashrc`

``` bash
. spacer-bash
```

Then the following command are available:

``` bash
spacer-change
spacer-create
spacer-layout
```

as well as completion. A suggest bash configuration is:

``` bash
. spacer-bash
alias layout=spacer-layout
alias change=spacer-change
alias create=spacer-create
alias new="spacer-change --create --layout"
complete -F spacer::completion layout
complete -F spacer::completion change
complete -F spacer::completion create
complete -F spacer::completion new
```

## Config

The configuration is based on `locations`, `creators` and `layouts`.

A simple location that does nothing except cleanup the path is:

``` yaml
spacer:
  locations:
  - name: match-documents
    current_path_regex: .*/[Dd]ocuments/.*
    current_path_command: sed 's:[Dd]ocuments/:[docs]:'
```

the full configuration spec is:

``` yaml
spacer:
  locations:
  - name: name-of-location
    current_path_regex: .*match-the-cwd.*
    current_path_command: sed 's:subsitute-path/:[via piping]:'
    change_path_regex: /place/the/command/.*.can.move.from/
    change_path_command: echo directory/to/move/to # also a pipe
    creator_name: name-of-creator
    layout_name: name-of-layout
  creators:
  - name: name-of-creator
    command: |
      complicate command to create
      that is multi-line if necessary
  layouts:
  - name: name-of-layout
    location_name: name-of-location to change to
    command: command to run after change
    children:
    - direction: up
      percent: 1-99 split percentage (default 50)
      location_name: name-of-location to change to
      command: command to run after change
      children:
      - direction: right
      - direction: left
    - direction: down
```

Almost all fields besides name are optional.

See `examples/config.yaml` and the test configs for more.
