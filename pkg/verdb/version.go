package verdb

type Version []int

func (new Version) IsNewerThan(old Version) bool {
	if len(old) != len(new) {
		return false
	}

	for i, _ := range old {
		if old[i] < new[i] {
			return true
		}
		if old[i] > new[i] {
			return false
		}
	}
	return false
}
