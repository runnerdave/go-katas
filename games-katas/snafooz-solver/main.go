package main

func SolveSnafooz(pcs [6][6][6]int) [6][6][6]int {
	var cube [6][6][6]int
	cube = [6][6][6]int{{{0, 0, 1, 1, 0, 0},
		{0, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 1},
		{1, 0, 1, 0, 1, 1}},

		{{0, 1, 0, 0, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1},
			{0, 0, 1, 1, 0, 1}},

		{{0, 0, 1, 1, 0, 1},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1},
			{0, 0, 1, 1, 0, 0}},

		{{0, 0, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 0},
			{0, 1, 0, 0, 1, 0}},

		{{0, 0, 1, 1, 0, 1},
			{1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1},
			{1, 1, 0, 0, 1, 1}},

		{{0, 0, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1},
			{0, 1, 0, 0, 1, 0}}}
	return cube
}

type rule struct {
	position int
	rule     bool
}

var snafoozRules rule
var (
	s0 [6][6]int
	s1 [6][6]int
	s2 [6][6]int
	s3 [6][6]int
	s4 [6][6]int
	s5 [6][6]int

	r0b = rule{0, CheckSides(ExtractBottom(s0), ExtractTop(s2))} //r2t
	r0t = rule{0, CheckSides(ExtractTop(s0), ExtractBottom(s5))} //r5b
	r0r = rule{0, CheckSides(ExtractRight(s0), ExtractTop(s3))}  //r3t
	r0l = rule{0, CheckSides(ExtractLeft(s0), ExtractTop(s1))}   //r1t

	r1r = rule{1, CheckSides(ExtractRight(s1), ExtractLeft(s2))}  //r2l
	r1b = rule{1, CheckSides(ExtractBottom(s1), ExtractLeft(s4))} //r4l

	r2r = rule{2, CheckSides(ExtractRight(s2), ExtractLeft(s3))} //r3l
	r2b = rule{2, CheckSides(ExtractBottom(s2), ExtractTop(s4))} //r4t

	r3b = rule{3, CheckSides(ExtractBottom(s3), ExtractRight(s4))} //r4r

	r4b = rule{4, CheckSides(ExtractBottom(s4), ExtractTop(s5))} //r5t

	r5l = rule{5, CheckSides(ExtractLeft(s5), ExtractLeft(s1))}   //r1l
	r5r = rule{1, CheckSides(ExtractRight(s1), ExtractRight(s3))} //r3r
)

func ApplyRule0(pcs [6][6][6]int) [6][6][6]int {
	var used [6]bool
	//fix side0 and apply each rule by rotating each side
	s0 = pcs[0]
	used[0] = true
	s2 = FillUpSide(s2, used, pcs, ExtractBottom, ExtractTop)
	s1 = FillUpSide(s1, used, pcs, ExtractLeft, ExtractTop)
	s3 = FillUpSide(s3, used, pcs, ExtractRight, ExtractTop)
	s5 = FillUpSide(s5, used, pcs, ExtractTop, ExtractBottom)
	return [6][6][6]int{s0, s1, s2, s3, s4, s5}
}

func FillUpSide(side [6][6]int, usedPlaces [6]bool, pieces [6][6][6]int, f1, f2 func([6][6]int) [6]int) [6][6]int {
	for i, v := range usedPlaces {
		needToMirror := true
		if !v {
			t := pieces[i]
			for j := 0; j <= 3; j++ {
				if CheckSides(f1(s0), f2(t)) {
					side = t
					usedPlaces[i] = true
					needToMirror = false
					break
				}
				t = RotateSide(t)
			}
			if needToMirror {
				t := MirrorSide(pieces[i])
				for j := 0; j <= 3; j++ {
					if CheckSides(f1(s0), f2(t)) {
						side = t
						usedPlaces[i] = true
						break
					}
					t = RotateSide(t)
				}
				needToMirror = false
			}
		}
	}
	return side
}

func CheckSides(s1, s2 [6]int) bool {
	ok := true
	for i := range s1 {
		if s1[i] == s2[i] {
			ok = false
			break
		}
	}
	return ok
}

func ExtractTop(side [6][6]int) [6]int {
	return side[0]
}

func ExtractBottom(side [6][6]int) [6]int {
	return side[5]
}

func ExtractLeft(side [6][6]int) [6]int {
	var left [6]int
	for i, s := range side {
		left[i] = s[0]
	}
	return left
}

func ExtractRight(side [6][6]int) [6]int {
	var right [6]int
	for i, s := range side {
		right[i] = s[5]
	}
	return right
}

func RotateSide(side [6][6]int) [6][6]int {
	var sideR [6][6]int
	var top [6]int
	var bottom [6]int
	for i, v := range sideR {
		for j := range v {
			sideR[i][j] = 1
		}
	}
	for j, a := range side {
		top[j] = a[5]
	}
	for j, a := range side {
		bottom[j] = a[0]
	}
	for j, a := range side[0] {
		sideR[5-j][0] = a
	}
	for j, a := range side[5] {
		sideR[5-j][5] = a
	}
	sideR[0] = top
	sideR[5] = bottom
	return sideR
}

func MirrorSide(side [6][6]int) [6][6]int {
	sideR := [6][6]int{{1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1}}
	//top
	t := ExtractTop(side)
	for i := range t {
		sideR[0][5-i] = t[i]
	}
	//bottom
	b := ExtractBottom(side)
	for i := range side[5] {
		sideR[5][5-i] = b[i]
	}
	// right
	l := ExtractLeft(side)
	for i := range l {
		sideR[i][5] = l[i]
	}
	// left
	r := ExtractRight(side)
	for i := range r {
		sideR[i][0] = r[i]
	}
	return sideR
}
