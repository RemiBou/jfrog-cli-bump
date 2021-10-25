# bump 

## About this plugin
This plugin provides a way to automatically bump your dependencies

## Installation with JFrog CLI
Installing the latest version:

`$ jfrog plugin install bump`

Installing a specific version:

`$ jfrog plugin install bump@version`

Uninstalling a plugin

`$ jfrog plugin uninstall bump`

## Usage
### Commands
* vcs : configure your vcs provider (only Bitbucket cloud available now)
    - Arguments:
        - url : your bitbucket cloud rest url
        - token : your personnal access token
    - Example:
    ```
  $ jfrog bump vcp https://mybitbucket.com/rest

## Additional info
None.

## Release Notes
The release notes are available [here](RELEASE.md).
