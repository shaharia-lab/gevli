<h1 align="center">Gevli</h1>
<p align="center">Implementing Event Listener for Golang Application</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/shahariaazam/gevli"><img src="https://pkg.go.dev/badge/github.com/shahariaazam/gevli.svg" height="20"/></a>
</p><br/><br/>

<p align="center">
  <a href="https://github.com/shahariaazam/gevli/actions/workflows/CI.yaml"><img src="https://github.com/shahariaazam/gevli/actions/workflows/CI.yaml/badge.svg" height="20"/></a>
  <a href="https://codecov.io/gh/shahariaazam/gevli"><img src="https://codecov.io/gh/shahariaazam/gevli/branch/master/graph/badge.svg?token=NKTKQ45HDN" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_gevli"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_gevli&metric=reliability_rating" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_gevli"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_gevli&metric=vulnerabilities" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_gevli"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_gevli&metric=security_rating" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_gevli"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_gevli&metric=sqale_rating" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_gevli"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_gevli&metric=code_smells" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_gevli"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_gevli&metric=ncloc" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_gevli"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_gevli&metric=alert_status" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_gevli"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_gevli&metric=duplicated_lines_density" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_gevli"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_gevli&metric=bugs" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_gevli"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_gevli&metric=sqale_index" height="20"/></a>
</p><br/><br/>


## Usage

```shell
go get go get github.com/shahariaazam/gevli
```

And start using like -

```golang
package main

import (
	"fmt"
	"github.com/shahariaazam/gevli"
)

func main() {
	emitter := gevli.NewEventEmitter()

	// Add listener for "message" event.
	emitter.AddListener("message", func(event gevli.Event) {
		fmt.Println("Message received:", event.Data)
	})

	// Add listener for "ping" event.
	emitter.AddListener("ping", func(event gevli.Event) {
		fmt.Println("Ping received:", event.Data)
	})

	// Emit "message" event.
	emitter.EmitSync("message", "Hello, world!")

	// Emit "ping" event.
	emitter.Emit("ping", nil)
}
```

## Documentation

Full documentation is available on [pkg.go.dev/github.com/shahariaazam/gevli](https://pkg.go.dev/github.com/shahariaazam/gevli#section-documentation)

### 📝 License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/shahariaazam/gevli/blob/master/LICENSE) file for details.