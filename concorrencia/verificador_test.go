package concorrencia

import (
	"reflect"
	"testing"
	"time"
)

func mockVerificadorWebsites(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func slowStubVerificadorWebsites(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkVerificadorWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "uma url"
	}

	for i := 0; i < b.N; i++ {
		VerificadorWebsites(slowStubVerificadorWebsites, urls)
	}
}

func TestVerificadorWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	wait := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := VerificadorWebsites(mockVerificadorWebsites, websites)

	if !reflect.DeepEqual(wait, got) {
		t.Fatalf("Esperava %v Recebeu %v", wait, got)
	}

}
