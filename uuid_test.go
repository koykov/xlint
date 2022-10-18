package xlint

import (
	"testing"
)

var (
	exampleUUID = []struct {
		raw string
		ok  bool
	}{
		{"f47ac10b-58cc-0372-8567-0e02b2c3d479", true},
		{"f47ac10b-58cc-1372-8567-0e02b2c3d479", true},
		{"f47ac10b-58cc-2372-8567-0e02b2c3d479", true},
		{"f47ac10b-58cc-3372-8567-0e02b2c3d479", true},
		{"f47ac10b-58cc-4372-8567-0e02b2c3d479", true},
		{"f47ac10b-58cc-5372-8567-0e02b2c3d479", true},
		{"f47ac10b-58cc-6372-8567-0e02b2c3d479", true},
		{"f47ac10b-58cc-7372-8567-0e02b2c3d479", true},
		{"f47ac10b-58cc-8372-8567-0e02b2c3d479", true},
		{"f47ac10b-58cc-9372-8567-0e02b2c3d479", true},
		{"f47ac10b-58cc-a372-8567-0e02b2c3d479", true},
		{"f47ac10b-58cc-b372-8567-0e02b2c3d479", true},
		{"f47ac10b-58cc-c372-8567-0e02b2c3d479", true},
		{"f47ac10b-58cc-d372-8567-0e02b2c3d479", true},
		{"f47ac10b-58cc-e372-8567-0e02b2c3d479", true},
		{"f47ac10b-58cc-f372-8567-0e02b2c3d479", true},

		{"urn:uuid:f47ac10b-58cc-4372-0567-0e02b2c3d479", true},
		{"URN:UUID:f47ac10b-58cc-4372-0567-0e02b2c3d479", true},
		{"f47ac10b-58cc-4372-0567-0e02b2c3d479", true},
		{"f47ac10b-58cc-4372-1567-0e02b2c3d479", true},
		{"f47ac10b-58cc-4372-2567-0e02b2c3d479", true},
		{"f47ac10b-58cc-4372-3567-0e02b2c3d479", true},
		{"f47ac10b-58cc-4372-4567-0e02b2c3d479", true},
		{"f47ac10b-58cc-4372-5567-0e02b2c3d479", true},
		{"f47ac10b-58cc-4372-6567-0e02b2c3d479", true},
		{"f47ac10b-58cc-4372-7567-0e02b2c3d479", true},
		{"f47ac10b-58cc-4372-9567-0e02b2c3d479", true},
		{"f47ac10b-58cc-4372-a567-0e02b2c3d479", true},
		{"f47ac10b-58cc-4372-b567-0e02b2c3d479", true},
		{"f47ac10b-58cc-4372-c567-0e02b2c3d479", true},
		{"f47ac10b-58cc-4372-d567-0e02b2c3d479", true},
		{"f47ac10b-58cc-4372-e567-0e02b2c3d479", true},
		{"f47ac10b-58cc-4372-f567-0e02b2c3d479", true},

		{"f47ac10b158cc-5372-a567-0e02b2c3d479", false},
		{"f47ac10b-58cc25372-a567-0e02b2c3d479", false},
		{"f47ac10b-58cc-53723a567-0e02b2c3d479", false},
		{"f47ac10b-58cc-5372-a56740e02b2c3d479", false},
		{"f47ac10b-58cc-5372-a567-0e02-2c3d479", false},
		{"g47ac10b-58cc-4372-a567-0e02b2c3d479", false},

		{"{f47ac10b-58cc-0372-8567-0e02b2c3d479}", true},
		{"{f47ac10b-58cc-0372-8567-0e02b2c3d479", false},
		{"f47ac10b-58cc-0372-8567-0e02b2c3d479}", false},

		{"f47ac10b58cc037285670e02b2c3d479", true},
		{"f47ac10b58cc037285670e02b2c3d4790", false},
		{"f47ac10b58cc037285670e02b2c3d47", false},
	}
)

func TestValidateUUID(t *testing.T) {
	for _, stage := range exampleUUID {
		t.Run(stage.raw, func(t *testing.T) {
			ok, _ := ValidateUUIDStr(stage.raw)
			if ok != stage.ok {
				t.Errorf("validation failed on %s, got %v expected %v", stage.raw, ok, stage.ok)
			}
		})
	}
}

func BenchmarkValidateUUID(b *testing.B) {
	for _, stage := range exampleUUID {
		b.Run(stage.raw, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				ok, _ := ValidateUUIDStr(stage.raw)
				if ok != stage.ok {
					b.Errorf("validation failed on %s, got %v expected %v", stage.raw, ok, stage.ok)
				}
			}
		})
	}
}
