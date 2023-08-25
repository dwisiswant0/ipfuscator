package ipfuscator

import (
	"net"
	"testing"
)

func TestNewIPFuscator(t *testing.T) {
	testCases := []struct {
		name      string
		ip        net.IP
		expect    *IPFuscator
		expectErr bool
	}{
		{
			name: "ValidIPv4",
			ip:   net.ParseIP("192.168.0.1"),
			expect: &IPFuscator{
				hex:   []string{"0xc0", "0xa8", "0x0", "0x1"},
				ips:   []string{"192", "168", "0", "1"},
				octal: []string{"300", "250", "1", "1"},
			},
			expectErr: false,
		},
		{
			name:      "InvalidIPv6",
			ip:        net.ParseIP("2001:db8::1"),
			expect:    nil,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := New(tc.ip)

			if tc.expectErr && err == nil {
				t.Error("Expected error but got nil")
			}

			if !tc.expectErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestTo(t *testing.T) {
	ip := net.ParseIP("192.168.0.1")
	ipf, _ := New(ip)

	t.Run("BaseWithPadding", func(t *testing.T) {
		s := ipf.ToBaseWithPadding()
		if s == "" {
			t.Error("Expected non-empty string")
		}
	})

	t.Run("CircledDigits", func(t *testing.T) {
		s := ipf.ToCircledDigits()
		if s == "" {
			t.Error("Expected non-empty string")
		}
	})

	t.Run("Decimal", func(t *testing.T) {
		s := ipf.ToDecimal()
		if s == "" {
			t.Error("Expected non-empty string")
		}
	})

	t.Run("Hex", func(t *testing.T) {
		s := ipf.ToHex()
		if s == "" {
			t.Error("Expected non-empty string")
		}
	})

	t.Run("HexWithPadding", func(t *testing.T) {
		s := ipf.ToHexWithPadding()
		if s == "" {
			t.Error("Expected non-empty string")
		}
	})

	t.Run("IPv6CompatibleV4", func(t *testing.T) {
		s := ipf.ToIPv6CompatibleV4()
		if s == "" {
			t.Error("Expected non-empty string")
		}
	})

	t.Run("NoZeros", func(t *testing.T) {
		s := ipf.ToNoZeros()
		if s == "" {
			t.Error("Expected non-empty string")
		}
	})

	t.Run("Octal", func(t *testing.T) {
		s := ipf.ToOctal()
		if s == "" {
			t.Error("Expected non-empty string")
		}
	})

	t.Run("OctalWithPadding", func(t *testing.T) {
		s := ipf.ToOctalWithPadding()
		if s == "" {
			t.Error("Expected non-empty string")
		}
	})

	t.Run("Rand8Bits", func(t *testing.T) {
		s := ipf.ToRand8Bits()
		if s == "" {
			t.Error("Expected non-empty string")
		}
	})

	t.Run("RandBase", func(t *testing.T) {
		s := ipf.ToRandBase()
		if s == "" {
			t.Error("Expected non-empty string")
		}
	})

	t.Run("RandBaseWithPadding", func(t *testing.T) {
		s := ipf.ToRandBaseWithPadding()
		if s == "" {
			t.Error("Expected non-empty string")
		}
	})
}

func BenchmarkNew(b *testing.B) {
	ip := net.ParseIP("192.168.0.1")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = New(ip)
	}
}

func BenchmarkToBaseWithPadding(b *testing.B) {
	ip := net.ParseIP("192.168.0.1")
	ipf, _ := New(ip)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ipf.ToBaseWithPadding()
	}
}

func BenchmarkToCircledDigits(b *testing.B) {
	ip := net.ParseIP("192.168.0.1")
	ipf, _ := New(ip)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ipf.ToCircledDigits()
	}
}

func BenchmarkToDecimal(b *testing.B) {
	ip := net.ParseIP("192.168.0.1")
	ipf, _ := New(ip)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ipf.ToDecimal()
	}
}

func BenchmarkToHex(b *testing.B) {
	ip := net.ParseIP("192.168.0.1")
	ipf, _ := New(ip)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ipf.ToHex()
	}
}

func BenchmarkToHexWithPadding(b *testing.B) {
	ip := net.ParseIP("192.168.0.1")
	ipf, _ := New(ip)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ipf.ToHexWithPadding()
	}
}

func BenchmarkToIPv6CompatibleV4(b *testing.B) {
	ip := net.ParseIP("192.168.0.1")
	ipf, _ := New(ip)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ipf.ToIPv6CompatibleV4()
	}
}

func BenchmarkToNoZeros(b *testing.B) {
	ip := net.ParseIP("192.168.0.1")
	ipf, _ := New(ip)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ipf.ToNoZeros()
	}
}

func BenchmarkToOctal(b *testing.B) {
	ip := net.ParseIP("192.168.0.1")
	ipf, _ := New(ip)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ipf.ToOctal()
	}
}

func BenchmarkToOctalWithPadding(b *testing.B) {
	ip := net.ParseIP("192.168.0.1")
	ipf, _ := New(ip)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ipf.ToOctalWithPadding()
	}
}

func BenchmarkToRand8Bits(b *testing.B) {
	ip := net.ParseIP("192.168.0.1")
	ipf, _ := New(ip)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ipf.ToRand8Bits()
	}
}

func BenchmarkToRandBase(b *testing.B) {
	ip := net.ParseIP("192.168.0.1")
	ipf, _ := New(ip)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ipf.ToRandBase()
	}
}

func BenchmarkToRandBaseWithPadding(b *testing.B) {
	ip := net.ParseIP("192.168.0.1")
	ipf, _ := New(ip)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ipf.ToRandBaseWithPadding()
	}
}
