package main

import (
	"reflect"
	"testing"
)

func TestConvertStrToInt(t *testing.T) {
	tests := []struct {
		name      string
		input     []string
		expected  []int
		shouldErr bool
	}{
		{
			name:      "valid numbers",
			input:     []string{"1", "2", "3"},
			expected:  []int{1, 2, 3},
			shouldErr: false,
		},
		{
			name:      "invalid number",
			input:     []string{"1", "a"},
			expected:  nil,
			shouldErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := convertStrToInt(tt.input)

			if tt.shouldErr && err == nil {
				t.Fatalf("expected error, got nil")
			}

			if !tt.shouldErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tt.shouldErr && !reflect.DeepEqual(result, tt.expected) {
				t.Fatalf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPerformOperation_Add(t *testing.T) {
	result, err := performOperation(ADD, 1, 2, 3)
	if err != nil {
		t.Fatal(err)
	}
	if result != 6 {
		t.Fatalf("got %.1f, want 6.0", result)
	}
}

func TestPerformOperation_Diff(t *testing.T) {
	result, err := performOperation(DIFF, 10, 3, 2)
	if err != nil {
		t.Fatal(err)
	}
	if result != 5 {
		t.Fatalf("got %.1f, want 5.0", result)
	}
}

func TestPerformOperation_Multiply(t *testing.T) {
	result, err := performOperation(MULITPLY, 2, 3, 4)
	if err != nil {
		t.Fatal(err)
	}
	if result != 24 {
		t.Fatalf("got %.1f, want 24.0", result)
	}
}

func TestPerformOperation_Divide(t *testing.T) {
	result, err := performOperation(DIVIDE, 10, 2)
	if err != nil {
		t.Fatal(err)
	}
	if result != 5 {
		t.Fatalf("got %.1f, want 5.0", result)
	}
}

func TestPerformOperation_DivideByZero_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic, but did not panic")
		}
	}()

	performOperation(DIVIDE, 10, 0)
}

func TestPerformOperation_Divide_InvalidArgs(t *testing.T) {
	_, err := performOperation(DIVIDE, 10, 2, 3)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
