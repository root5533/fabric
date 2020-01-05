package etcdraft

// DetermineRepID canopus
func DetermineRepID(id uint64) uint64 {
	if id < 4 {
		return uint64(1)
	}

	return uint64(4)

}

// LeafNodesFromID canopus
func LeafNodesFromID(id uint64, iter int) []uint64 {
	if id < 4 && iter == 1 {
		return []uint64{1, 2, 3}
	}

	return []uint64{4, 5}
}
