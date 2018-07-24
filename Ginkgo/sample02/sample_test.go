package sample_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "sample"
)
/*
var _ = Describe("Sample", func() {
	Context("Valid inputs for add function", func() {
		It("Should return the correct result given 2 integers", func() {
			result, err := Add(2, 2)
			Expect(result).Should(Equal(4))
			Expect(err).Should(BeNil())
		})
	})
})
*/
var _ = Describe("Sample2", func() {
	Context("Test Add function", func() {
		It("Should return the correct result given 2 ints", func() {
			result, err := Add(2, 4)
			Expect(result).Should(Equal(6))
			Expect(err).Should(BeNil())
		})
	})

	Context("Test Del function", func() {
		It("Should return the correct result given 2 ints", func() {
			result, err := Del(8, 3)
			Expect(result).Should(Equal(5))
			Expect(err).Should(BeNil())
		})
	})
})
