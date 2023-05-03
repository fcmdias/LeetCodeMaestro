package main

import "testing"

func TestAlgoOne(t *testing.T) {
	var tests = []struct {
		wordsDict []string
		word1     string
		word2     string
		want      int
	}{
		{
			wordsDict: []string{"practice", "makes", "perfect", "coding", "makes"},
			word1:     "makes",
			word2:     "coding",
			want:      1,
		},
		{
			wordsDict: []string{"practice", "makes", "perfect", "coding", "makes"},
			word1:     "makes",
			word2:     "makes",
			want:      3,
		},
	}

	for _, test := range tests {
		got := algoOne(test.wordsDict, test.word1, test.word2)
		if test.want != got {
			t.Fatalf("got %v and wanted %v", got, test.want)
		}
	}

}
func TestAlgoTwo(t *testing.T) {
	var tests = []struct {
		wordsDict []string
		word1     string
		word2     string
		want      int
	}{
		{
			wordsDict: []string{"practice", "makes", "perfect", "coding", "makes"},
			word1:     "makes",
			word2:     "coding",
			want:      1,
		},
		{
			wordsDict: []string{"practice", "makes", "perfect", "coding", "makes"},
			word1:     "makes",
			word2:     "makes",
			want:      3,
		},
	}

	for _, test := range tests {
		got := algoTwo(test.wordsDict, test.word1, test.word2)
		if test.want != got {
			t.Fatalf("got %v and wanted %v", got, test.want)
		}
	}

}

func TestAlgoThree(t *testing.T) {
	var tests = []struct {
		wordsDict []string
		word1     string
		word2     string
		want      int
	}{
		{
			wordsDict: []string{"practice", "makes", "perfect", "coding", "makes"},
			word1:     "makes",
			word2:     "coding",
			want:      1,
		},
		{
			wordsDict: []string{"practice", "makes", "perfect", "coding", "makes"},
			word1:     "makes",
			word2:     "makes",
			want:      3,
		},
	}

	for _, test := range tests {
		got := algoThree(test.wordsDict, test.word1, test.word2)
		if test.want != got {
			t.Fatalf("got %v and wanted %v", got, test.want)
		}
	}

}

var benchmarkTest = struct {
	wordsDict []string
	word1     string
	word2     string
	want      int
}{
	wordsDict: []string{"practice", "makes", "perfect", "coding", "makes"},
	word1:     "makes",
	word2:     "perfect",
	want:      1,
}

func BenchmarkAlgoOne(b *testing.B) {

	for n := 0; n < b.N; n++ {
		algoOne(benchmarkTest.wordsDict, benchmarkTest.word1, benchmarkTest.word2)
	}
}

func BenchmarkAlgoTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		algoTwo(benchmarkTest.wordsDict, benchmarkTest.word1, benchmarkTest.word2)
	}
}

func BenchmarkAlgoThree(b *testing.B) {
	for n := 0; n < b.N; n++ {
		algoThree(benchmarkTest.wordsDict, benchmarkTest.word1, benchmarkTest.word2)
	}
}
