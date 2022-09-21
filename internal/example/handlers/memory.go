// Package handlers provides Gin handler methods
//
// Adapted from:
// https://github.com/pixie-io/pixie-demos/tree/main/go-garbage-collector
package handlers

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"

	"github.com/bkonkle/go-garbage/internal/example/io"
	"github.com/gin-gonic/gin"
)

// Allocate allocates memory
func Allocate(_ *gin.Context, input *io.AllocateInput) (*io.MessageOutput, error) {
	arrayLength := parseOr(input.ArrayLength, 1000)
	bytesPerElement := parseOr(input.BytesPerElement, 8192)

	arr := generateRandomStringArray(arrayLength, bytesPerElement)

	message := fmt.Sprintf("Generated string array with %d bytes of data", len(arr)*len(arr[0]))

	return &io.MessageOutput{
		Success: true,
		Message: message,
	}, nil
}

// RunGC runs the garbage collector
func RunGC(_ *gin.Context) (*io.MessageOutput, error) {
	runtime.GC()

	message := fmt.Sprintf("Ran garbage collector")

	return &io.MessageOutput{
		Success: true,
		Message: message,
	}, nil
}

// AllocateAndRunGC allocates memory and runs the garbage collector
func AllocateAndRunGC(_ *gin.Context, input *io.AllocateInput) (*io.MessageOutput, error) {
	arrayLength := parseOr(input.ArrayLength, 1000)
	bytesPerElement := parseOr(input.BytesPerElement, 8192)

	arr := generateRandomStringArray(arrayLength, bytesPerElement)

	runtime.GC()

	message := fmt.Sprintf(
		"Generated string array with %d bytes of data and ran garbage collector",
		len(arr)*len(arr[0]),
	)

	return &io.MessageOutput{
		Success: true,
		Message: message,
	}, nil
}

func parseOr(value *string, fallback int) int {
	if value == nil {
		return fallback
	}

	result, err := strconv.Atoi(*value)
	if err != nil {
		return fallback
	}

	return result
}

func launchTimeAfter() {
	go func() {
		<-time.After(10 * time.Second)
		fmt.Println("Waited for 10 seconds")
	}()
}

func launchSleep() {
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("Slept for 10 seconds")
	}()
}

// generateRandomString function adapted from
// https://kpbird.medium.com/golang-generate-fixed-size-random-string-dd6dbd5e63c0
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func generateRandomString(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// Allocate a bunch of memory
func generateRandomStringArray(arrLength, strBytes int) []string {
	arr := make([]string, arrLength)
	for i := 0; i < arrLength; i++ {
		arr[i] = generateRandomString(strBytes)
	}
	return arr
}
