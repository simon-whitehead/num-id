package numid

import (
	"testing"
)

func Test100EqualsCb(t *testing.T) {
	input := int64(100)
	expected := "Cb"
	actual := Encode(input)

	assertEqual(t, expected, actual)
}

func TestCbEquals100(t *testing.T) {
	input := "Cb"
	expected := int64(100)
	actual := Decode(input)

	assertEqual(t, expected, actual)
}

func TestEncodeDecode(t *testing.T) {
	input := int64(1100)
	expected := int64(1100)
	actual := Decode(Encode(input))

	assertEqual(t, expected, actual)
}

func TestFailEncodeDecode(t *testing.T) {
	input := int64(1001)
	expected := int64(101)
	actual := Decode(Encode(input))

	assertNotEqual(t, expected, actual)
}

func TestReverseBasicString(t *testing.T) {
	input := "Hello world!"
	expected := "!dlrow olleH"
	actual := reverseString(input)

	assertEqual(t, expected, actual)
}

func TestPalindrome(t *testing.T) {
	input := "kayak"
	expected := "kayak"
	actual := reverseString(input)

	assertEqual(t, expected, actual)
}

func assertNotEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected == actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
