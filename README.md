# HK

CLI tools for hakuna.ch

[ ![Codeship Status for twhiston/hk](https://app.codeship.com/projects/c134b890-a619-0135-4ff4-16f7c16b7dca/status?branch=master)](https://app.codeship.com/projects/255432)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/5d777742d71a44679e3a513c3144c71f)](https://www.codacy.com/app/twhiston/hk?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=twhiston/hk&amp;utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/twhiston/hk)](https://goreportcard.com/report/github.com/twhiston/hk)

## Install

`go get -u github.com/twhiston/hk`

or download the binary directly from the releases page for your architecture,
rename it to `hk` and add it to your path

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
Note that some commands will not be available unless you have an organization api key.

You can additionally add the top-level key (not nested under `hakuna:`)
```
vertical_print: true
```
to the config file to get vertical table printouts, which is useful if
on a small monitor or viewing a very large set of data which breaks the terminal line.

Keys in the `.hk.yml` file can be set/overriden with environmental variables.
Environmental variables should match the key names and be prefixed by `HK_`

## Commands

Commands currently cover the entire api and are described in the manifest.yml file.
Acting as other users are not supported currently

use `hk --help` or `hk {command} --help` to find out options and further sub-commands

### Date Parsing

`hk` uses `github.com/olebedev/when` for date parsing allowing you to either parse the date as per the api format or to use natural language such as

`hk time list --date "last wednesday"`
`hk time create --start "today 1:15am" --end "2am"`
`hk timer start --start "1 hour ago"`

## Additional Commands

In addition to the generated commands there are custom commands that don't only return the API functionality

### today

using the `hk today` command you can see your total for today including any active timers, and it displays a
sum of all times, this means that you can quickly use this command to get an overview of your day

### timer new

Stops a current timer, if running, and starts a new one, takes all of the flags that can be parsed to timer start

### timer running

Returns true or false depending on if a timer is currently running

## Using with bitbar

If you want to use hk with bitbar to start and stop timers, and see your current days total you can use the following script

Make sure that you set your gopath and path properly so that hk can be found
(thanks to [Philippe Hässig](https://github.com/neckhair) for the idea for this)
```
#!/bin/bash

export GOPATH=/Users/$(whoami)/go
export PATH=$PATH:$GOPATH/bin

# If hk is not installed get it
which hk > /dev/null
if [ $? -ne 0 ]; then
  go get -u github.com/twhiston/hk
fi


function startTimer ()
{
	hk timer start --config=/Users/$(whoami)/.hk.yml
}
function stopTimer()
{
	hk timer stop --config=/Users/$(whoami)/.hk.yml
}

## This case statement captures what was passed to param1 and decides which function to call
case "$1" in
    startTimer)
    startTimer
    ;;
    stopTimer)
    stopTimer
    ;;
    *)
esac

time=$(hk today | grep Today | cut -d "|" -f 3)

if [[ $(hk timer running) == "true" ]]; then
	echo "⏱ $time|trim=true"
else
	echo "   $time|trim=true"
fi

echo "---"
echo "Start Timer | terminal=false bash=$0 param1=startTimer refresh=true"
echo "Stop Timer  | terminal=false bash=$0 param1=stopTimer  refresh=true"

```


## Development

All the commands are generated from the manifest.yml file using the code in the generator folder
To generate the output simply run `go generate` in the root folder
