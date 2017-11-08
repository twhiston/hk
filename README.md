# HK

CLI tools for hakuna.ch

[ ![Codeship Status for twhiston/hk](https://app.codeship.com/projects/c134b890-a619-0135-4ff4-16f7c16b7dca/status?branch=master)](https://app.codeship.com/projects/255432)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/5d777742d71a44679e3a513c3144c71f)](https://www.codacy.com/app/twhiston/hk?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=twhiston/hk&amp;utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/twhiston/hk)](https://goreportcard.com/report/github.com/twhiston/hk)

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

use `hk --help` or `hk {command} --help` to find out options and further sub-commands

### time

The time command has some special case options to make it more useful:
- With no option will take todays date
- As well as dates in the api format of yyyy-mm-dd you can use `yesterday` or `y` to get yesterdays details

## Additional Commands

In addition to the generated commands there are custom commands that dont only return the API functionality

### today

using the `hk today` command you can see your total for today including any active timers, and it displays a
sum of all times, this means that you can quickly use this command to get an overview of your day


## Generated Code

All the commands are generated from the manifest.yml file using the code in the generator folder
To generate the output simply run `go generate` in the root folder

## TODO

- Generate code for command tests
- Cover everything else with tests
- Consolidate templates for verbs in to a single one