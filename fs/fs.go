package fs

func GoTo(nextDir string, dirStack []string) []string {
	if nextDir == ".." {
		if len(dirStack) > 1 {
			return dirStack[:len(dirStack)-1]
		}
	}

	return append(dirStack, nextDir)
}
