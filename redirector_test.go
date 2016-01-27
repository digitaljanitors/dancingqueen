package main_test

import (
	"math"
	"testing"

	dq "github.com/digitaljanitors/dancingqueen"
)

var redirectTests = []struct {
	Host     string
	Referer  string
	Expected string
}{
	{"launch.newsinc.com", "nydailynews.com", "launch.newsinc.com/21/test"},
}

func TestRedirector(t *testing.T) {
	t.SkipNow()

	//for _, tt := range redirectTests {
	// do something here
	//}
}

func TestSingleVersion(t *testing.T) {
	var svTest = []struct {
		in  string
		out string
	}{
		{"version1", "version1"},
		{"2", "2"},
		{"100", "100"},
	}
	for _, tt := range svTest {
		sv := &dq.SingleVersion{tt.in}
		actual := sv.GetNextVersion()
		if tt.out != actual {
			t.Error("expected:", tt.out)
			t.Error("actual:  ", actual)
		}
	}
}

func TestMultipleVersions(t *testing.T) {
	versions := [][]dq.WeightedVersion{
		[]dq.WeightedVersion{dq.WeightedVersion{0.75, "version1"}, dq.WeightedVersion{0.25, "version2"}},
		[]dq.WeightedVersion{dq.WeightedVersion{0.25, "1"}, dq.WeightedVersion{0.25, "2"}, dq.WeightedVersion{0.5, "3"}},
	}

	done := make(chan bool, len(versions))

	for _, tt := range versions {
		go func() {
			mv := dq.NewMultipleVersions(tt)
			counts := make(map[string]int, len(tt))

			const rounds = 10e6

			for i := 0; i < rounds; i++ {
				tv := mv.GetNextVersion()
				counts[tv]++
			}

			const eps = 0.001
			for i := range tt {
				g := float64(counts[tt[i].Version]) / float64(rounds)
				if math.Abs(g-tt[i].Weight) > eps {
					t.Errorf("version: %s, expected +/- 0.001:", tt[i].Weight)
					t.Errorf("version: %s, actual:            ", g)
				}
			}

			done <- true
		}()
	}
	<-done

}
