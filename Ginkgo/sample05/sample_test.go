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
	Describe("the first:", func() {
		Context("Test Add function", func() {
			By("Add")
			Specify("test add is ok", func() {
				//It("Should return the correct result given 2 ints", func() {
					result, err := Add(2, 4)
					Expect(result).Should(Equal(6))
					Expect(err).Should(BeNil())
				//})
			})
		})

		Context("Test Subtract function", func() {
			By("Subtract")
			It("Should return the correct result given 2 ints", func() {
				result, err := Subtract(8, 3)
				Expect(result).Should(Equal(5))
				Expect(err).Should(BeNil())
			})
		})
	})

	Describe("the second:", func() {
		Context("the 3 & 4:", func() {
			By("Multiply")
			It("Test Multiply function", func() {
				result, err := Multiply(2, 3)
				Expect(result).Should(Equal(6))
				Expect(err).Should(BeNil())
			})
			By("Divide")
			It("Test Divide function", func() {
				result, err := Divide(8, 4)
				//Expect(result).Should(Equal(2))
				Expect(result).To(Equal(2))
				//Expect(err).Should(BeNil())
				Expect(err).To(BeNil())
			})
			By("Divide is not ok")
			It("Test Divide function", func() {
				result, err := Divide(2, 0)
				Expect(result).Should(Equal(0))
				Expect(err).ShouldNot(BeNil())
			})
		})
	})
})
