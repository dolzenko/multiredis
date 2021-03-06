package multiredis

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {

	It("should init simple clients", func() {
		client := New(&Options{
			Addrs: []string{"127.0.0.1:6379"},
			DB:    9,
		})
		defer client.Close()

		Expect(client).To(BeAssignableToTypeOf(simpleClient{}))
	})

	It("should init cluster clients", func() {
		client := New(&Options{
			Addrs:       []string{"10.0.0.1:6379", "10.0.0.2:6379"},
			DialTimeout: time.Millisecond,
		})
		defer client.Close()

		Expect(client).To(BeAssignableToTypeOf(clusterClient{}))
	})

})

// --------------------------------------------------------------------

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "multiredis")
}
