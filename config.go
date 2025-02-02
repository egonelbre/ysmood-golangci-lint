package main

const defaultConf = `
run:
  skip-dirs-use-default: false

linters:
  enable:
    - gofmt
    - golint
    - gocyclo
    - misspell

gocyclo:
  min-complexity: 15

issues:
  exclude-use-default: false

`
