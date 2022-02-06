# adp

[![Go Reference](https://pkg.go.dev/badge/EmeraldCityCircusArts/adp.svg)](https://pkg.go.dev/EmeraldCityCircusArts/adp)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Package adp implements a simple Golang library for working with the adp.com API.

Based on ADP API documentation available at [ADP API Explorer](https://developers.adp.com/articles/api/all/apiexplorer)

## Installation

```bash
$ go get -u github.com/EmeraldCityCircusArts/adp
```

## Quick Start

Add this import line to the file you're working in:

```Go
import "github.com/EmeraldCityCircusArts/adp"
```

## Schema and Model sources
ADP provided model and schema files are located in the models/ and schemas/ directories.

Schema and model file source locations:

- [table](https://developers.adp.com/articles/api/table-v1-api)
- [workers](https://developers.adp.com/articles/api/workers-v2-api)
- [workSchedules](https://developers.adp.com/articles/api/work-schedules-v1-api)

## Table usage
https://developers.adp.com/articles/preview/guide-validation-table-code-list-api-guide-for-adp-workforce-now-8?chapter=3

For example, the specific codelist URL for workSchedules Departments can be retrieved from workSchedulesMeta


## License

EmeraldCityCircusArts/adp is licensed under the [MIT](https://opensource.org/licenses/MIT) License.
