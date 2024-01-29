package rangex

func Backward[E any](s []E) func(func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

func Chunck[E any](s []E, size int) func(func(int, []E) bool) {
	return func(yield func(int, []E) bool) {
		c := make([]E, 0, size)
		for i, e := range s {
			if len(c) >= size {
				if !yield(i-1, c) {
					return
				}
				c = make([]E, 0, size)
			}
			c = append(c, e)
		}

		yield(len(s)-1, c)
	}
}

func Filter[E any](s []E, f func(E) bool) func(func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i, e := range s {
			if f(e) {
				if !yield(i, e) {
					return
				}
			}
		}
	}
}
