//go:build go1.18

package subjectparser

import "testing"

// Parse must not panic, and the numbers it reports are never negative.
func FuzzParse(f *testing.F) {
	f.Add(`[1/1] "debian-11.6.0-amd64-netinst.iso" (1/568)`)
	f.Add(`(debian-11.1.0-i386-DVD-1) [03/91] - "debian-11.1.0-i386-DVD-1.part01.rev" yEnc (1/328)`)
	f.Add(`File 2 of 9 - "x.zip" yEnc (1/4)`)
	f.Add(`[99999999999999999999/2] "big.bin" (1/1)`)
	f.Fuzz(func(t *testing.T, s string) {
		p, err := Parse(s)
		if err != nil {
			t.Fatal(err)
		}
		if p.File < 0 || p.TotalFiles < 0 || p.Segment < 0 || p.TotalSegments < 0 {
			t.Fatalf("%q: negative numbers %+v", s, p)
		}
	})
}
