package subjectparser

import "testing"

func BenchmarkParse(b *testing.B) {
	subjects := []string{
		`[1/1] "debian-11.6.0-amd64-netinst.iso" (1/568)`,
		`(debian-11.1.0-i386-DVD-1) [03/91] - "debian-11.1.0-i386-DVD-1.part01.rev" yEnc (1/328)`,
		`[3/7] "ubuntu-22.04-live-server-amd64.vol00+01.par2" (1/8)`,
		`file.bin yEnc (1/2)`,
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := Parse(subjects[i%len(subjects)]); err != nil {
			b.Fatal(err)
		}
	}
}
