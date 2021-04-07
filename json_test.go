package xlint

import (
	"testing"

	"github.com/koykov/jsonlint"
)

var (
	validJSON0 = []byte(`{"menu":{"header":"SVG Viewer","items":[{"id":"Open"},{"id":"OpenNew","label":"Open New"},null,{"id":"ZoomIn","label":"Zoom In"},{"id":"ZoomOut","label":"Zoom Out"},{"id":"OriginalView","label":"Original View"},null,{"id":"Quality"},{"id":"Pause"},{"id":"Mute"},null,{"id":"Find","label":"Find..."},{"id":"FindAgain","label":"Find Again"},{"id":"Copy"},{"id":"CopyAgain","label":"Copy Again"},{"id":"CopySVG","label":"Copy SVG"},{"id":"ViewSVG","label":"View SVG"},{"id":"ViewSource","label":"View Source"},{"id":"SaveAs","label":"Save As"},null,{"id":"Help"},{"id":"About","label":"About Adobe CVG Viewer..."}]}}`)
	validJSON1 = []byte(`{"name":"John","age":31,"city":"New York"}`)

	invalidJSON0 = []byte(`["test":123]`)
	invalidJSON1 = []byte(`{"Cartoon Foxes":{{"Name":"Fox Tall","Job":"Bein' tall"},{"Name":"Fox Small","Job":"Bein' small"}}}`)
)

func TestValidateJSON(t *testing.T) {
	if o, err := ValidateJSON(validJSON0); err != nil {
		t.Errorf("JSON fail %s at %d", err, o)
	}
	if o, err := ValidateJSON(validJSON1); err != nil {
		t.Errorf("JSON fail %s at %d", err, o)
	}
	if o, err := ValidateJSON(invalidJSON0); err != jsonlint.ErrUnexpId || o != 7 {
		t.Errorf("JSON fail %s at %d", err, o)
	}
	if o, err := ValidateJSON(invalidJSON1); err != jsonlint.ErrUnexpId || o != 18 {
		t.Errorf("JSON fail %s at %d", err, o)
	}
}

func BenchmarkValidateValidJSON0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if o, err := ValidateJSON(validJSON0); err != nil {
			b.Errorf("JSON fail %s at %d", err, o)
		}
	}
}

func BenchmarkValidateValidJSON1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if o, err := ValidateJSON(validJSON1); err != nil {
			b.Errorf("JSON fail %s at %d", err, o)
		}
	}
}

func BenchmarkValidateInvalidJSON0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if o, err := ValidateJSON(invalidJSON0); err != jsonlint.ErrUnexpId || o != 7 {
			b.Errorf("JSON fail %s at %d", err, o)
		}
	}
}

func BenchmarkValidateInvalidJSON1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if o, err := ValidateJSON(invalidJSON1); err != jsonlint.ErrUnexpId || o != 18 {
			b.Errorf("JSON fail %s at %d", err, o)
		}
	}
}
