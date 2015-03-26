# SynergyAI
Intro to Artificial Intelligence final project. Build teams based on learned compatibility between team members.

## Setting up the GOPATH

## Getting the code
Once the GOPATH is set up correctly, you can get the code using the `go` utility by simply running
`go get github.com/amirblum/SynergyAI` and
`go get github.com/amirblum/goutils`

## Running the code
1. Enter the SynergyAI directory (under GOPATH/src/github.com/amirblum/SynergyAI)
2. Run `go install`
3. Run the SynergyAI program from GOPATH/bin/SynergyAI.

## Configuration
The SynergyAI executable receives a config file to run the code.
There are existing configuration files located in the "tests" folder, that can run the tests we wrote about in the
report.
In addition you can generate your own configuration files by using the configurator by running
.../SynergyAI/configurator/configurator.html
Once you have chosen or create a config file, run `SynergyAI --config=config.json`

## Note:
The World and Task files to feed the configuration are located in the data folder