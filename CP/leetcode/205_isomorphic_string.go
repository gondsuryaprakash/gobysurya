package leetcode

func IsIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	sp := make(map[string]string, 0)
	for i := 0; i < len(s); i++ {

		v, ok := sp[string(s[i])]
		if !ok {
			sp[string(s[i])] = string(t[i])
		} else {
			if v != string(t[i]) {
				return false
			}
		}
	}

	tp := make(map[string]string, 0)

	for i := 0; i < len(t); i++ {

		v, ok := tp[string(t[i])]
		if !ok {
			tp[string(t[i])] = string(s[i])
		} else {
			if v != string(s[i]) {
				return false
			}
		}
	}

	return true
}
