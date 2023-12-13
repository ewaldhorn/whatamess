func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Buffer(make([]byte, 1000), 1000)
	var z []string
	var r int64
	n := 0
	f := 1
	s.Scan()
	Sscan(s.Text(), &n)
	s.Scan()
	z = Split(s.Text(), " ")
	for i := 0; i < n; i++ {
		t, _ := strconv.ParseInt(z[i], 10, 32)
		if f > 0 {
			r = t
			f = 0
		}
		if Abs(float64(t)) < Abs(float64(r)) {
			r = t
		} else if Abs(float64(t)) == Abs(float64(r)) {
			if r < 0 && t < 0 {
				r = t
			} else {
				r = int64(Abs(float64(t)))
			}
		}
	}
	Println(r)
}