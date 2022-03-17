package balancer

import (
	"net/http"
	"net/url"

	"github.com/AbhiswasPatil/olscheduler/httputil"
	"github.com/AbhiswasPatil/olscheduler/lambda"
)

type Balancer interface {
	SelectWorker(r *http.Request, l *lambda.Lambda) (url.URL, *httputil.HttpError)
	ReleaseWorker(workerUrl url.URL)
	AddWorker(workerUrl url.URL)
	RemoveWorker(workerUrl url.URL)
	GetAllWorkers() []url.URL
}
