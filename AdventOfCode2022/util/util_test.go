package util

import (
    "testing"
)

func TestMaxIntALtB(t *testing.T) {
    expected := 2
    actual := Max(1, 2)
    if actual != expected {
        t.Fatalf(`expected %v(%T) but got %v(%T)`, expected, expected, actual, actual)
    }
}

func TestMaxIntAEqB(t *testing.T) {
    expected := 2
    actual := Max(2, 2)
    if actual != expected {
        t.Fatalf(`expected %v but got %v`, expected, actual)
    }
}

func TestMaxIntAGtB(t *testing.T) {
    expected := 2
    actual := Max(2, 1)
    if actual != expected {
        t.Fatalf(`expected %v but got %v`, expected, actual)
    }
}

func TestMaxFloatALtB(t *testing.T) {
    expected := 2.0
    actual := Max(1.0, 2.0)
    if actual != expected {
        t.Fatalf(`expected %v(%T) but got %v(%T)`, expected, expected, actual, actual)
    }
}

func TestMaxAFloatEqB(t *testing.T) {
    expected := 2.0
    actual := Max(2.0, 2.0)
    if actual != expected {
        t.Fatalf(`expected %v but got %v`, expected, actual)
    }
}

func TestMaxFloatAGtB(t *testing.T) {
    expected := 2.0
    actual := Max(2.0, 1.0)
    if actual != expected {
        t.Fatalf(`expected %v but got %v`, expected, actual)
    }
}

func TestAbsPositiveInt(t *testing.T) {
    expected := 2
    actual := Abs(2)
    if actual != expected {
        t.Fatalf(`expected %v but got %v`, expected, actual)
    }
}

func TestAbsNegativeInt(t *testing.T) {
    expected := 2
    actual := Abs(-2)
    if actual != expected {
        t.Fatalf(`expected %v but got %v`, expected, actual)
    }
}

func TestAbsPositiveFloat(t *testing.T) {
    expected := 2.0
    actual := Abs(2.0)
    if actual != expected {
        t.Fatalf(`expected %v but got %v`, expected, actual)
    }
}

func TestAbsNegativeFloat(t *testing.T) {
    expected := 2.0
    actual := Abs(-2.0)
    if actual != expected {
        t.Fatalf(`expected %v but got %v`, expected, actual)
    }
}
