package item_test

import (
	. "github.com/hamster21/minion/item"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"
)

var _ = Describe("Task", func() {
	var (
		parent = new(Collection)
	)

	It("remembers the creation time", func() {
		t := NewTask("Summary", "Detail", parent)

		Expect(t).NotTo(BeNil())
		Expect(t.CreatedAt()).To(BeTemporally("~", time.Now()))
	})
})
