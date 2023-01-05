[![Go Reference](https://pkg.go.dev/badge/github.com/tommzn/hob-core.svg)](https://pkg.go.dev/github.com/tommzn/hob-core)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tommzn/hob-core)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/tommzn/hob-core)
[![Go Report Card](https://goreportcard.com/badge/github.com/tommzn/hob-core)](https://goreportcard.com/report/github.com/tommzn/hob-core)
[![Actions Status](https://github.com/tommzn/hob-core/actions/workflows/go.pkg.auto-ci.yml/badge.svg)](https://github.com/tommzn/hob-core/actions)

# HomeOffice Button - Core Components
Tis repository contains some core components for the [HomeOffice Button - Time Tracking](https://github.com/tommzn/hob-timetracker) Project.

## Events
Events are defined using Protobuf and will be automatically generate Go sources.
### GenerateReportRequest
This event is used to trigger report generattion from different sources.

## Helper
This package provides serialize/deserialize helper to publish and consume events in Protobuf format.

## Change event schema
### Checkout new branch
```
git checkout -b proto-v1.0.x
```
### Commit and push your changes
Commit and push your changes in feature branch. This will create a PR you can merge vit GitHub. Do not forget to add a merge comment starting with "fix: " to ensure a new 
release is generated after merging.
```
git add .
git commit -m "Update/Add event schema"
git push origin proto-v1.0.x
```

# Links
[HomeOffice Button - Time Tracking](https://github.com/tommzn/hob-timetracker)  
[AWS IoT 1-Click](https://aws.amazon.com/iot-1-click/?nc1=h_ls)  
