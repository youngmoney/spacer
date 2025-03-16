# Spacer

Workspaces for your terminal.

## Usage

### Path

The path command outputs a context depend path name, useful in command
prompts.

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

``` bash
spacer create <location>
```

It can also be used as part of change:

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

## Config

``` yaml
spacer:
  locations:
  - name: match-documents
    current_path_regex: .*/[Dd]ocuments/.*
    current_path_command: pwd | sed 's:[Dd]ocuments/:[docs]:'
```

See `examples/config.yaml` and the test configs for more.
