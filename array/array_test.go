package array

import (
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {

	// resultadoTeste := func(t *testing.T, want int, got int) {
	// 	t.Helper()
	// 	if want != got {
	// 		t.Errorf("Esperava %d mas obteve %d", want, got)
	// 	}
	// }

	// t.Run("Somando array", func(t *testing.T) {
	// 	got := Somar([]int{1, 2, 3, 4})
	// 	want := 10
	// 	resultadoTeste(t, want, got)
	// })

	t.Run("Somando varios arrays", func(t *testing.T) {
		got := SomaTudo([]int{1, 2, 3, 4}, []int{1, 2}, []int{4, 2})
		want := []int{10, 3, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Esperava %v e recebeu %v ", want, got)
		}
	})

}

func BenchmarkSomar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Somar([]int{1, 2, 3, 4})
	}
}
