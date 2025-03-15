#!/usr/bin/env bash

function path() {
    go run . --config tests/simple.config.yaml path
}

diff <(path) <(echo example-path)

function change() {
    go run . --config tests/simple.config.yaml change "$1"
}

function change_cwd() {
    f=$(mktemp)
    go run . --config tests/simple.config.yaml  change --cwd_file "$f" "$1"
    diff "$f" <(echo "$2")
}

diff <(change all) <(echo cd "$HOME")
change_cwd all "$HOME"
change does-not-exit 2>/dev/null && echo "change does-not-exist should fail"

function create() {
    go run . --config tests/simple.config.yaml create "$1"
}

function create_cwd() {
    f=$(mktemp)
    go run . --config tests/simple.config.yaml create --cwd_file "$f" "$1" > /dev/null
    diff "$f" <(echo "$2")
}

diff <(create all) <(echo cd $(pwd))
create_cwd all $(pwd)
diff <(create creatable) <(echo simple; echo cd "$HOME")
create_cwd creatable "$HOME"
create does-not-exit 2>/dev/null && echo "create does-not-exist should fail"

function layout() {
    go run . --config tests/simple.config.yaml layout "$1"
}

function layout_cwd() {
    f=$(mktemp)
    go run . --config tests/simple.config.yaml layout --cwd_file "$f" "$1" > /dev/null
    diff "$f" <(echo "$2")
}

diff <(layout all) <(echo look around; echo cd $(pwd))
diff <(layout creatable) <(echo tmp; echo cd /tmp)
layout_cwd all $(pwd)
layout_cwd creatable /tmp

function layout_position() {
    go run . --config tests/layout.config.yaml layout --position "$2" "$1"
}

function layout_position_cwd() {
    f=$(mktemp)
    go run . --config tests/layout.config.yaml layout --cwd_file "$f" --position "$2" "$1" > /dev/null
    diff "$f" <(echo "$2")
}

diff <(layout_position here) <(echo 4; echo cd $(pwd))
diff <(layout_position here 0) <(echo 4-up; echo cd $(pwd))
diff <(layout_position here 1) <(echo 4-right; echo cd $(pwd))
