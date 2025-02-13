package calc

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	// Arrange
	a, b := 5, 7
	// Act
	got := Add(a, b)
	// Assert
	want := 12

	if got != want {
		t.Errorf("Want %v, got %v", want, got)
	}
}

func TestMinus(t *testing.T) {
	// Arrange
	a, b := 7, 5
	// Act
	got := Minus(a, b)
	// Assert
	want := 2

	if got != want {
		t.Errorf("Want %v, got %v", want, got)
	}
}

// func TestDivide(t *testing.T) {
// 	t.Run("Divide 8 by 4 returns 2", func(t *testing.T) {
// 		// Arrange
// 		a, b := 8, 4
// 		// Act
// 		got, _ := Divide(a, b)
// 		// Assert
// 		want := 2

// 		if got != want {
// 			t.Errorf("Want %v, got %v", want, got)
// 		}
// 	})

// 	t.Run("Divide by 0 throws an error", func(t *testing.T) {
// 		// Arrange
// 		a, b := 8, 0
// 		// Act
// 		_, gotErr := Divide(a, b)
// 		// Assert
// 		wantErr := errors.New("impossible to divide by 0")

// 		if errors.Is(gotErr, wantErr) {
// 			t.Errorf("WantErr %v, gotErr %v", wantErr, gotErr)
// 		}
// 	})
// }

func TestDivideFico(t *testing.T) {
	testCases := []struct {
		name string
		a, b int
		want int
		err  error
	}{
		{name: "Divide 8 by 4 returns 2", a: 8, b: 4, want: 2, err: nil},
		{"Divide by 0 throws an error", 8, 0, -1, errors.New("impossible to divide by 0")},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := Divide(test.a, test.b)
			if err != nil {
				if errors.Is(err, test.err) {
					t.Errorf("WantErr %v, err %v", test.err, err)
				}
			}

			if got != test.want {
				t.Errorf("Want %v, got %v", test.want, got)
			}

		})
	}
}
