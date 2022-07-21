package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestHellosNames(t *testing.T) {
	names := []string{
		"one",
		"two",
		"three",
	}
	want := map[string]*regexp.Regexp{
		"one":   regexp.MustCompile(`\b` + "one" + `\b`),
		"two":   regexp.MustCompile(`\b` + "two" + `\b`),
		"three": regexp.MustCompile(`\b` + "three" + `\b`),
	}
	msgs, err := Hellos(names)
	for name, msg := range msgs {
		if !want[name].MatchString(msg) || err != nil {
			t.Fatalf("Hellos(%v:%v) = %q, %v, want match for %#q, nil", name, msg, msg, err, want)
		}
	}
}
