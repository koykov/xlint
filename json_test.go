package xlint

import (
	"strconv"
	"testing"

	"github.com/koykov/jsonlint"
)

var (
	exampleJSON = []struct {
		raw []byte
		off int
		err error
	}{
		{[]byte(`{"menu":{"header":"SVG Viewer","items":[{"id":"Open"},{"id":"OpenNew","label":"Open New"},null,{"id":"ZoomIn","label":"Zoom In"},{"id":"ZoomOut","label":"Zoom Out"},{"id":"OriginalView","label":"Original View"},null,{"id":"Quality"},{"id":"Pause"},{"id":"Mute"},null,{"id":"Find","label":"Find..."},{"id":"FindAgain","label":"Find Again"},{"id":"Copy"},{"id":"CopyAgain","label":"Copy Again"},{"id":"CopySVG","label":"Copy SVG"},{"id":"ViewSVG","label":"View SVG"},{"id":"ViewSource","label":"View Source"},{"id":"SaveAs","label":"Save As"},null,{"id":"Help"},{"id":"About","label":"About Adobe CVG Viewer..."}]}}`), 0, nil},
		{[]byte(`{"name":"John","age":31,"city":"New York"}`), 0, nil},
		{[]byte(`["test":123]`), 7, jsonlint.ErrUnexpId},
		{[]byte(`{"Cartoon Foxes":{{"Name":"Fox Tall","Job":"Bein' tall"},{"Name":"Fox Small","Job":"Bein' small"}}}`), 18, jsonlint.ErrUnexpId},
	}
)

func TestValidateJSON(t *testing.T) {
	for i, stage := range exampleJSON {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			o, err := ValidateJSON(stage.raw)
			if stage.err != nil {
				if err != stage.err || o != stage.off {
					t.Errorf("JSON fail %s at %d", err, o)
				}
				return
			}
			if err != nil {
				t.Errorf("JSON fail %s at %d", err, o)
			}
		})
	}
}

func BenchmarkValidateJSON(b *testing.B) {
	for i, stage := range exampleJSON {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				o, err := ValidateJSON(stage.raw)
				if stage.err != nil {
					if err != stage.err || o != stage.off {
						b.Errorf("JSON fail %s at %d", err, o)
					}
					return
				}
				if err != nil {
					b.Errorf("JSON fail %s at %d", err, o)
				}
			}
		})
	}
}
