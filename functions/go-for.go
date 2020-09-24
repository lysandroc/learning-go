package functions

import (
	"fmt"
)

func ExampleFor() {
	coursesCompleted := []string{"docker", "golang", "k8s"}
	coursesInProgress := []string{"docker", "node", "java"}

	for _, progress := range coursesInProgress {
		for _, completed := range coursesCompleted {
			if progress == completed {
				fmt.Println(completed + "hun, i think theses course already has finished")
			} else {
				fmt.Println(progress + " its in progress")
				break
			}
		}
	}

}
