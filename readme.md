# gomarkdoc

A GitHub Action that generates markdown documentation from Go source code.

## Inputs

### output-file

**Required** The output file to write the documentation to. Default: `DOC.md`

## Usage

The action is designed to run as part of your GitHub Actions workflow. An example workflow that generates documentation and commits it to the repository is as follows:

yamlCopy code

```yaml
name: Generate documentation

on:
  push:
    branches:
      - main

jobs:
  generate-docs:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v2
  
      - name: Generate documentation
        uses: ./.github/actions/gomarkdoc
        with:
          output-file: DOC.md
  
      - name: Commit documentation
        uses: EndBug/add-and-commit@v9
        with:
          commit: --signoff
          default_author: github_actor
          message: 'Update docs'`
```