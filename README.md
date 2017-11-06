# HK

CLI tools for hakuna.ch

## Install

`go get -u github.com/twhiston/hk`

or download the binary directly from the releases page

## Usage

Before using hk you will need to create a `.hk.yml` config file in your home directory
(or another dir but then you will need to set the --config option on startup)

This file must contain the following data

```
hakuna:
  token: your_api_access_token
  domain: your_hakuna_domain
```

Once these details are provided hk will be able to connect to Hakuna.
Note that some commands will not be available unless you have an organization api key

## Commands

Commands currently cover a large chunk of the api and are described in the manifest.yml file.
Currently you cannot update an entry or perform any absence functions. Coming soon!

### time

The time command with no option will take todays date, the option --date allows you to specify a date to query
As well as dates in the api format of yyyy-mm-dd you can use `yesterday` or `y` to get yesterdays details


## Generated Code

All the commands are generated from the manifest.yml file using the code in the generator folder
To generate the output simple run `go generate` in the root folder

## TODO

- Generate code for command tests
- Cover everything else with tests
- consolidate templates for verbs in to a single one