package advent

import (
	"math/big"
	"testing"
)

func TestAdvent6Lights(t *testing.T) {
	tests := []struct {
		in         string
		out1, out2 int
	}{
		{"turn on 0,0 through 4,4\nturn off 1,2 through 3,3\ntoggle 0,0 through 2,2", 14, 0},
	}

	// 0001100000
	// 0001100000
	// 0110100000
	// 1000100000
	// 1111100000
	// 0000000000
	for i, tt := range tests {
		r1, r2 := Advent06Lights(tt.in)
		if r1 != tt.out1 || r2 != tt.out2 {
			t.Errorf("Test %d: Advent06Lights(%q) => %d, %d, want %d, %d", i, tt.in, r1, r2, tt.out1, tt.out2)

		}
	}
}

func TestAdvent6ToBinary(t *testing.T) {
	bigIntBin := BigIntToBinary(big.NewInt(int64(8214555720323791245)))
	want := "111000111111111111101101110011110110001100011010111100110001101"
	if bigIntBin != want {
		t.Errorf("BigIntToBinary(6) => [%s], want [%s]", bigIntBin, want)
	}
}

func TestAdvent6ToBinaryZeroes(t *testing.T) {
	bigIntBin := BigIntToBinary(big.NewInt(int64(32)))
	want := "100000"
	if bigIntBin != want {
		t.Errorf("BigIntToBinary(32) => [%s], want [%s]", bigIntBin, want)
	}
}

func TestAdvent6PopCount(t *testing.T) {
	count1 := PopCountInt(8214555720323791245)
	want := 41
	if count1 != want {
		t.Errorf("PopCountInt() => [%d], want [%d]", count1, want)
	}

	g := LightGridInt{big.NewInt(8214555720323791245), 0}
	count2 := g.PopCount()
	// "111000111111111111101101110011110110001100011010111100110001101"
	if count2 != want {
		t.Errorf("g.PopCount() => [%d], want [%d]", count2, want)
	}
}

func TestAdvent6LightGridInt(t *testing.T) {
	g := NewLightGridInt(5)
	g.I = g.I.Add(g.I, big.NewInt(225))

	val := g.I.Int64()
	wantI := int64(225)
	if val != wantI {
		t.Errorf("Int64() => [%d], want [%d]", val, wantI)
	}

	bin := g.ToBinary()
	want := "0000000000000000011100001"
	if bin != want {
		t.Errorf("ToBinary() => [%s], want [%s]", bin, want)
	}

	binSq := g.ToBinarySquare()
	want = "00000\n00000\n00000\n00111\n00001\n"
	if binSq != want {
		t.Errorf("ToBinarySquare() => [%s], want [%s]", binSq, want)
	}

	mask := g.MaskFromRectangle(1, 2, 3, 3)
	wantd := big.NewInt(14784)
	if wantd.Cmp(mask) != 0 {
		t.Errorf("MaskFromRectangle() => [%d], want [%d]", mask, wantd)
	}
}

func TestAdvent6LightGridIntHuge(t *testing.T) {
	g := NewLightGridInt(1000)
	g.I = g.I.Add(g.I, big.NewInt(225))

	val := g.I.Int64()
	wantI64 := int64(225)
	if val != wantI64 {
		t.Errorf("Int64() => [%d], want [%d]", val, wantI64)
	}

	bin := g.ToBinary()
	wantI := 1000000
	if len(bin) != wantI {
		t.Errorf("ToBinary() => [%d], want [%d]", len(bin), wantI)
	}

	binSq := g.ToBinarySquare()
	wantI = 1000000 + 1000
	if len(binSq) != wantI {
		t.Errorf("ToBinarySquare() => [%d], want [%d]", len(binSq), wantI)
	}
}

func TestAdvent6SetOn(t *testing.T) {
	g := NewLightGridInt(5)
	g.RectSetOn(1, 2, 3, 3)
	//println("int64", g.I.Int64())
	binSq := g.ToBinarySquare()
	//println("binSq:", binSq)
	want := "00000\n00000\n01110\n01110\n00000\n"
	if binSq != want {
		t.Errorf("RectSetOn() => [%s], want [%s]", binSq, want)
	}

	c := g.PopCount()
	wantC := 6
	if c != wantC {
		t.Errorf("PopCount() => %d, want %d", c, wantC)
	}
}

func TestAdvent6SetOff(t *testing.T) {
	g := NewLightGridInt(5)
	g.RectSetOn(0, 0, 4, 4)
	g.RectSetOff(1, 2, 3, 3)
	//println("int64", g.I.Int64())
	binSq := g.ToBinarySquare()
	//println("binSq:", binSq)
	want := "11111\n11111\n10001\n10001\n11111\n"
	if binSq != want {
		t.Errorf("RectSetOff() => [%s], want [%s]", binSq, want)
	}

	c := g.PopCount()
	wantC := 19
	if c != wantC {
		t.Errorf("PopCount() => %d, want %d", c, wantC)
	}
}

func TestAdvent6Toggle(t *testing.T) {
	g := NewLightGridInt(5)
	g.RectSetOn(0, 0, 4, 4)
	g.RectSetOff(1, 2, 3, 3)
	g.RectToggle(0, 0, 2, 2)

	//println("int64", g.I.Int64())
	binSq := g.ToBinarySquare()
	//println("binSq:", binSq)
	want := "00011\n00011\n01101\n10001\n11111\n"
	if binSq != want {
		t.Errorf("RectToggle() => [%s], want [%s]", binSq, want)
	}

	c := g.PopCount()
	wantC := 14
	if c != wantC {
		t.Errorf("PopCount() => %d, want %d", c, wantC)
	}
}
