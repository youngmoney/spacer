# Spacer

Workspaces for your terminal.

## Usage

### Path

The path command outputs a context depend path name, useful in command
prompts.

``` bash
spacer path
```

## Config

``` yaml
spacer:
  locations:
  - name: match-documents
    current_path_regex: .*/[Dd]ocuments/.*
    current_path_command: pwd | sed 's:[Dd]ocuments/:[docs]:'
```

See `examples/config.yaml`
