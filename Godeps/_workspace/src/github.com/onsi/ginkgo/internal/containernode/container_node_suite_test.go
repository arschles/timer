package containernode_test

import (
	. "github.com/arschles/timer/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/arschles/timer/Godeps/_workspace/src/github.com/onsi/gomega"
	"testing"
)

func TestContainernode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Containernode Suite")
}
