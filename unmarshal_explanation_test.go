package main_test

import (
	"reflect"
	"testing"
)

func BenchmarkUnmarshalPOCExplanation(b *testing.B) {
	// This benchmark illustrates the improvement made the the bson.Unmarshal codepath using the changes in
	// GODRIVER-1563. This function is broken up into two sub-benchmarks:
	//
	// 1. "intDecodeValue" - This represents how the driver's current decoder system works. The intDecodeValue function
	// is similar to the driver's IntDecodeValue decoder. It takes a reflect.Value and sets it. To call it, the
	// benchmark must first allocate a reflect.Value using reflect.New(tInt).Elem().
	//
	// 2. "intDecodeType" - This represents the decoder system after making the proposed changes. The intDecodeType
	// function takes an int and uses it to construct a new reflect.Value. The corresponding function added to the
	// driver is slightly more complex as it needs to actually take the reflect.Type being used. This is because some
	// decoders are used for multiple types (e.g. the same int decoder is used for int, int8, int16, int32, and int64),
	// so it needs to know what type to use when converting the BSON value it reads.
	//
	// Results from running these benchmarks:
	//
	// BenchmarkUnmarshalPOCExplanation/intDecodeValue-12              42923340                27.8 ns/op             8 B/op          1 allocs/op
	// BenchmarkUnmarshalPOCExplanation/intDecodeType-12               73488028                15.1 ns/op

	tInt := reflect.TypeOf(int(5))

	b.Run("intDecodeValue", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			val := reflect.New(tInt).Elem()
			intDecodeValue(val, i)
		}
	})
	b.Run("intDecodeType", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			val := intDecodeType(i)
			_ = val
		}
	})
}

func intDecodeValue(val reflect.Value, i int) {
	val.SetInt(int64(i))
}

func intDecodeType(i int) reflect.Value {
	return reflect.ValueOf(i)
}
