package concurrencypattern

import "fmt"

var worker int = 10

func Wroker(jobsChannel chan Jobs, result chan Result, workerId int) {
	//Reciving channel value

	for ch := range jobsChannel {

		// Sending result value to the channel
		result <- Result{
			JobId:    ch.Id,
			Name:     ch.JobName,
			WorkerId: workerId,
		}
	}
}

type Result struct {
	JobId    int
	Name     string
	WorkerId int
}

type Jobs struct {
	JobName string
	Id      int
}

func WrokerPool() {

	jobList := []string{"Job1", "Job2", "Job3", "Job4", "Job5", "Job6", "Job7", "Job8", "Job9", "Job10", "Job11", "Job12", "Job13", "Job14", "Job15", "Job16", "Job17", "Job18", "Job19", "Job20", "Job21", "Job22"}
	numberOfJobs := len(jobList)
	jobsChannel := make(chan Jobs, numberOfJobs)

	result := make(chan Result, numberOfJobs)

	// assign task to Worker
	for i := 0; i < worker; i++ {
		go Wroker(jobsChannel, result, i)
	}

	// Sending Work to the GoRouting
	for i := 0; i < len(jobList); i++ {
		jobsChannel <- Jobs{
			JobName: jobList[i],
			Id:      i,
		}
	}

	// Close the channel jobsChannel
	close(jobsChannel)

	// Reciving the result channel value
	for i := 0; i < numberOfJobs; i++ {
		res := <-result
		fmt.Printf("Jobs With JobId %d, and JobsName %s done! with WorkerId %d\n", res.JobId, res.Name, res.WorkerId)
	}

}
