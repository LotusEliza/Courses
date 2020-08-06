package streams

import (
	"github.com/gocraft/health"
	"os"
)

var (
	stdOutStream *health.Stream
)

func init() {
	stdOutStream = health.NewStream()
	stdOutStream.AddSink(&health.WriterSink{Writer: os.Stdout})

}

func GetStdOutStream() *health.Stream {
	return stdOutStream
}

func NewJob(name string) *health.Job {
	return stdOutStream.NewJob(name)
}
