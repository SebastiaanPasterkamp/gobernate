Package `gobernate` provides an easy HTTP Handler containing all end-points
required to run a golang service in Kubernetes. This code is roughly based on:
https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/

To use the `gobernate` package to Kubernetes enable your service simply use:

```go
import (
	"github.com/SebastiaanPasterkamp/gobernate"
)

func main() {
	g := gobernate.New("8080", "example", "1.0.0", "commit-sha", "build-time")

	shutdown := g.Launch()
	g.Ready()
	<-shutdown
}
```

The `gobernate` http service provides:
* `GET /version` to return a JSON structure with the service.
* `GET /health` to return `http.StatusOK` to indicate the (web) service is
  running.
* `GET /readiness` to signal when the service is ready. Expects `Ready()` to be
  called before signaling `http.StatusOK`, and will report if the service is
  already shutting down.
* `GET /metrics` to return Prometheus formatted metric data.
