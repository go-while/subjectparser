package subjectparser

import "testing"

func TestParse(t *testing.T) {
	for _, c := range []struct {
		in                         string
		header, filename, basename string
		file, files, seg, segs     int
	}{
		// real subjects (Debian/Ubuntu NZBs)
		{`[1/1] "debian-11.6.0-amd64-netinst.iso" (1/568)`, "debian-11.6.0-amd64-netinst", "debian-11.6.0-amd64-netinst.iso", "debian-11.6.0-amd64-netinst", 1, 1, 1, 568},
		{`(debian-11.1.0-i386-DVD-1) [03/91] - "debian-11.1.0-i386-DVD-1.part01.rev" yEnc (1/328)`, "(debian-11.1.0-i386-DVD-1)", "debian-11.1.0-i386-DVD-1.part01.rev", "debian-11.1.0-i386-DVD-1", 3, 91, 1, 328},
		{`[1/2] - "debian-12.2.0-amd64-netinst.iso" yEnc (1/919)`, "debian-12.2.0-amd64-netinst", "debian-12.2.0-amd64-netinst.iso", "debian-12.2.0-amd64-netinst", 1, 2, 1, 919},
		{`"debian-12.9.0-amd64-DVD-1.iso" yEnc (1/5555)`, "debian-12.9.0-amd64-DVD-1", "debian-12.9.0-amd64-DVD-1.iso", "debian-12.9.0-amd64-DVD-1", 1, 1, 1, 5555},
		{`[3/7] "ubuntu-22.04-live-server-amd64.vol00+01.par2" (1/8)`, "ubuntu-22.04-live-server-amd64", "ubuntu-22.04-live-server-amd64.vol00+01.par2", "ubuntu-22.04-live-server-amd64", 3, 7, 1, 8},
		{`[10/10] "ubuntu-24.04-live-server-amd64.vol127+33.par2" (1/64)`, "ubuntu-24.04-live-server-amd64", "ubuntu-24.04-live-server-amd64.vol127+33.par2", "ubuntu-24.04-live-server-amd64", 10, 10, 1, 64},
		// other common forms
		{`Header [1/3] "a.part1.rar" yEnc (1/10) 50.000.000 bytes`, "Header", "a.part1.rar", "a", 1, 3, 1, 10},
		{`[01/12] - "Some.Release.S01E02.1080p.mkv" yEnc (001/250)`, "Some.Release.S01E02.1080p", "Some.Release.S01E02.1080p.mkv", "Some.Release.S01E02.1080p", 1, 12, 1, 250},
		{`File 2 of 9 - "x.zip" yEnc (1/4)`, "x", "x.zip", "x", 2, 9, 1, 4},
		{`"archive.7z.001" yEnc (1/3)`, "archive", "archive.7z.001", "archive", 1, 1, 1, 3},
		{`archive.7z.001 yEnc (1/3)`, "archive", "archive.7z.001", "archive", 1, 1, 1, 3},
		{`file.bin yEnc (1/2)`, "file", "file.bin", "file", 1, 1, 1, 2},
		// no segment numbers: the filename after the last number pair must not be lost
		{`"file.bin" yEnc`, "file", "file.bin", "file", 1, 1, 1, 1},
		{`[1/5] "file.rar" yEnc`, "file", "file.rar", "file", 1, 5, 1, 1},
		{`[1/5] - "file.rar"`, "file", "file.rar", "file", 1, 5, 1, 1},
		{`file.bin`, "file", "file.bin", "file", 1, 1, 1, 1},
		// numbers too big for an int count as missing, not as the largest int
		{`[99999999999999999999/2] "big.bin" (1/1)`, "big", "big.bin", "big", 0, 2, 1, 1},
		{`"big.bin" yEnc (1/99999999999999999999)`, "big", "big.bin", "big", 1, 1, 1, 1},
	} {
		got, err := Parse(c.in)
		if err != nil {
			t.Errorf("%q: %v", c.in, err)
			continue
		}
		if got.Header != c.header || got.Filename != c.filename || got.Basefilename != c.basename ||
			got.File != c.file || got.TotalFiles != c.files || got.Segment != c.seg || got.TotalSegments != c.segs {
			t.Errorf("%q:\n got  header=%q filename=%q base=%q file=%d/%d seg=%d/%d\n want header=%q filename=%q base=%q file=%d/%d seg=%d/%d",
				c.in, got.Header, got.Filename, got.Basefilename, got.File, got.TotalFiles, got.Segment, got.TotalSegments,
				c.header, c.filename, c.basename, c.file, c.files, c.seg, c.segs)
		}
	}
}
