package grpctools

import (
	"google.golang.org/grpc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server", func() {
	var subject *Server

	BeforeEach(func() {
		subject = NewServer("test", "127.0.0.1:8080")
	})

	AfterEach(func() {
		subject.Stop()
	})

	It("should meter load", func() {
		Expect(func() { subject.Increment(1) }).NotTo(Panic())
	})

	It("should be a gRPC server", func() {
		Expect(subject.Server).To(BeAssignableToTypeOf(&grpc.Server{}))
		Expect(subject.Server).NotTo(BeNil())
	})

})
