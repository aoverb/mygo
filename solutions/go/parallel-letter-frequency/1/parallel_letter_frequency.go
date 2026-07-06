package parallelletterfrequency

import (
    "unicode"
    "strings"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
    text = strings.ToLower(text)
    var fmap FreqMap
    fmap = make(map[rune]int, 64)
	for _, r := range text {
        if !unicode.IsLetter(r) {
            continue
        }
        fmap[r]++
    }
    return fmap
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
    var fmap FreqMap
    fmap = make(map[rune]int, 64)
    channel := make(chan FreqMap)
	for _, text := range texts {
        go func() {
            channel <- Frequency(text)
        }()
    }
	for _ = range texts {
        submap := <- channel
        for k, v := range submap {
            fmap[k] += v
        }
    }
    return fmap
}
