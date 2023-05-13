package utils

func Contains[T comparable](s []T, str T) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func GetMapsKeys[K comparable, V any](myMap map[K]V) []K {
	keys := make([]K, len(myMap))

	i := 0
	for k := range myMap {
		keys[i] = k
		i++
	}
	return keys
}

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func Map[T any, V any](vs []T, f func(T) V) []V {
	vsm := make([]V, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
