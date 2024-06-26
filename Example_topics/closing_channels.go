//Closing a channel indicates that no more values will be sent on it.
//This can be useful to communicate completion to the channel’s receivers.

package main

import "fmt"

func main() {
	jobs := make(chan int, 5) //In this example we’ll use a jobs channel to communicate work to be done from the main() goroutine to a worker goroutine.
	//When we have no more jobs for the worker we’ll close the jobs channel.
	done := make(chan bool)

	go func() { //Here’s the worker goroutine. It repeatedly receives from jobs with j, more := <-jobs.
		//In this special 2-value form of receive, the more value will be false if jobs has been closed and all values in the channel have already been received.
		//We use this to notify on done when we’ve worked all our jobs.
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ { //This sends 3 jobs to the worker over the jobs channel, then closes it.
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done //We await the worker using the synchronization approach we saw earlier.

	_, ok := <-jobs //Reading from a closed channel succeeds immediately, returning the zero value of the underlying type.
	//The optional second return value is true if the value received was delivered by a successful send operation to the channel, or false if it was a zero value generated because the channel is closed and empty.
	fmt.Println("received more jobs:", ok)
}
