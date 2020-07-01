package main

import (
	"testing"
)

// Benchmarks ported from https://github.com/mongodb/mongonet/blob/Benchmarks/bsonutil_test.go

func BenchmarkMarshal(b *testing.B) {
	b.Run("flat", func(b *testing.B) {
		testCases := []struct {
			name    string
			docSize int
		}{
			{"empty", 0},
			{"small", 1},
			{"size 10", 10},
			{"size 50", 50},
			{"size 100", 100},
			{"size 500", 500},
			{"size 1000", 1000},
		}
		for _, tc := range testCases {
			b.Run(tc.name, func(b *testing.B) {
				doc := buildDoc(tc.docSize)

				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_, err := marshal(doc)
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	})

	b.Run("nested", func(b *testing.B) {
		testCases := []struct {
			name  string
			depth int
		}{
			// No "empty" test because a doc without any nesting is already covered by BenchmarkMarshal
			{"small", 1},
			{"size 10", 10},
			{"size 50", 50},
			{"size 100", 100},
			{"size 500", 500},
			{"size 1000", 1000},
		}
		for _, tc := range testCases {
			b.Run(tc.name, func(b *testing.B) {
				doc := buildNestedDoc(tc.depth)

				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_, err := marshal(doc)
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	})

	b.Run("commands", func(b *testing.B) {
		testCases := []struct {
			name string
			doc  interface{}
		}{
			{"isMaster response", getIsMasterResponse},
			{"findOne request", getFindOneRequest},
			{"findOne response", getFindOneRequest},
		}
		for _, tc := range testCases {
			b.Run(tc.name, func(b *testing.B) {
				b.ReportAllocs()

				for i := 0; i < b.N; i++ {
					_, err := marshal(tc.doc)
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	})
}

func BenchmarkUnmarshal(b *testing.B) {
	b.Run("flat", func(b *testing.B) {
		testCases := []struct {
			name    string
			docSize int
		}{
			{"empty", 0},
			{"small", 1},
			{"size 10", 10},
			{"size 50", 50},
			{"size 100", 100},
			{"size 500", 500},
			{"size 1000", 1000},
		}
		for _, tc := range testCases {
			b.Run(tc.name, func(b *testing.B) {
				doc := buildRawDoc(tc.docSize)

				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_, err := unmarshal(doc)
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	})

	b.Run("nested", func(b *testing.B) {
		testCases := []struct {
			name  string
			depth int
		}{
			// No "empty" test because a doc without any nesting is already covered by BenchmarkUnmarshal
			{"small", 1},
			{"size 10", 10},
			{"size 50", 50},
			{"size 100", 100},
			{"size 500", 500},
			{"size 1000", 1000},
		}
		for _, tc := range testCases {
			b.Run(tc.name, func(b *testing.B) {
				doc := buildNestedRawDoc(tc.depth)

				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_, err := unmarshal(doc)
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	})

	b.Run("commands", func(b *testing.B) {
		testCases := []struct {
			name string
			doc  interface{}
		}{
			{"isMaster response", getIsMasterResponse()},
			{"findOne request", getFindOneRequest()},
			{"findOne response", getFindOneResponse()},
		}
		for _, tc := range testCases {
			b.Run(tc.name, func(b *testing.B) {
				docBytes, err := marshal(tc.doc)
				if err != nil {
					b.Fatal(err)
				}

				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_, err := unmarshal(docBytes)
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	})
}
