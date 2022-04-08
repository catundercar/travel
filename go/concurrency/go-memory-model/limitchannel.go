package main

var limit = make(chan struct{}, 3)

func main() {
	work := make([]func(), 0, 10)
	for _, w := range work {
		go func(w func()) {
			limit <- struct{}{}
			w()
			// The kth receive on a channel with capacity C happens before the k+Cth send from that channel completes.
			<-limit
		}(w)
	}
	select{}
}
