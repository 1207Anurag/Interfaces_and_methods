import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Job struct {
	id  int
	url string
}
type Result struct {
	job Job
	b   string
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func CallUrl(u string) string {
	var r string
	res, err := http.Get(u)
	if err != nil {
		r="cannot access provided url"
	} else {
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			r="cannot read body of Response"
		} else {
			r = string("done")
		}
	}
	return r
}
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, CallUrl(job.url)}
		results <- output
	}
	wg.Done()
}
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		url := "https://jsonplaceholder.typicode.com/posts/"
		job := Job{i, url}
		jobs <- job
	}
	close(jobs)
}
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, url %s , body is %s\n", result.job.id, result.job.url, result.b)
	}
	done <- true
}
func main() {
	startTime := time.Now()
	noOfJobs := 200
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 100
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "s)