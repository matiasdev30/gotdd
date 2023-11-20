package benchmarking

import "testing"

func TestRepetir(t * testing.T){
	t.Run("testando respetições de uma caracter", func(t *testing.T) {
		resultado := Repetir("a")
		esperado := "aaaa"

		if resultado != esperado {
			t.Errorf("resultado -> '%s' esperado '%s'", resultado, esperado)
		}
	})
}


//Benchmark perfomatic test
func BenchmarkRepetir (b *testing.B){
	for c := 0; c < b.N; c ++{
		Repetir("a")
	}
}
