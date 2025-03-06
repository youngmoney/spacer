#!/usr/bin/env bash

diff <(go run . --config tests/simple.config.yaml path) <(echo example-path)
