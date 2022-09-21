package Bench

import (
	"math/rand"
	"testing"
	"time"

	"github.com/samber/lo"
)

type Serialized struct {
	Times []int64
}

type Data struct {
	Times []time.Time
}

func BenchmarkMemoryByValues(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	input := Serialized{Times: createTimes(10000)}

	for n := 0; n < b.N; n++ {
		// Run the operation b.N times
		indexToDelete := rand.Intn(10000) //nolint:gosec

		_ = deleteByValuesHandler(input, indexToDelete)
	}
}

func BenchmarkMemoryByRefs(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	input := Serialized{Times: createTimes(10000)}

	for n := 0; n < b.N; n++ {
		// Run the operation b.N times
		indexToDelete := rand.Intn(10000) //nolint:gosec

		_ = deleteByRefsHandler(input, indexToDelete)
	}
}

func createTimes(count int) []int64 {
	return lo.Map(lo.Range(count), func(i int, _ int) int64 {
		return int64(rand.Intn(int(time.Now().UnixMilli()))) //nolint:gosec
	})
}

func timeFromUnixMilli(ms int64) time.Time {
	return time.Unix(0, ms*int64(time.Millisecond)).UTC()
}

// By Values
// ---------

func deleteByValuesHandler(input Serialized, i int) Serialized {
	data := fromInputByValues(input)

	result := deleteTimeAtByValues(data, i)

	return toOutputByValues(result)
}

func deleteTimeAtByValues(data Data, i int) Data {
	times := data.Times

	times[i] = times[len(times)-1]
	times = times[:len(times)-1]

	return Data{Times: times}
}

func fromInputByValues(input Serialized) Data {
	return Data{
		Times: lo.Map(input.Times, func(t int64, _ int) time.Time {
			return timeFromUnixMilli(t)
		}),
	}
}

func toOutputByValues(input Data) Serialized {
	return Serialized{
		Times: lo.Map(input.Times, func(t time.Time, _ int) int64 {
			return t.UnixMilli()
		}),
	}
}

// By Refs
// -------

func deleteByRefsHandler(input Serialized, i int) Serialized {
	data := Data{
		Times: make([]time.Time, len(input.Times), len(input.Times)),
	}
	fromInputByRefs(&input, &data)

	deleteTimeAtByRefs(i, &data)

	output := Serialized{
		Times: make([]int64, len(input.Times), len(input.Times)),
	}
	toOutputByRefs(&data, &output)

	return output
}

func deleteTimeAtByRefs(i int, data *Data) {
	data.Times[i] = data.Times[len(data.Times)-1]
	data.Times = data.Times[:len(data.Times)-1]
}

func fromInputByRefs(input *Serialized, output *Data) {
	if output == nil {
		return
	}

	for i, time := range input.Times {
		output.Times[i] = timeFromUnixMilli(time)
	}
}

func toOutputByRefs(input *Data, output *Serialized) {
	if output == nil {
		return
	}

	for i, time := range input.Times {
		output.Times[i] = time.UnixMilli()
	}
}
