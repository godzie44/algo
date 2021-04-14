package process

type Process struct {
	S int
	F int
}

func ActivitySelector(processes []Process) []Process {
	var result []Process
	result = append(result, processes[0])

	k := 0
	for m := 1; m < len(processes); m++ {
		if processes[m].S >= processes[k].F {
			result = append(result, processes[m])
			k = m
		}
	}

	return result
}
