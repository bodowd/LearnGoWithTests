package concurrency

type WebsiteChecker func(string) bool

type result struct {
	url  string
	isUp bool
}

func CheckWebistes(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// by giving each anonymous function a parameter for the url `u` and then
		// calling the anonymous function with the url as the argument, we make sure that the value
		// of `u` is fixed as the value of url for the iteration of the loop that we're launching the go routin in.
		// `u` is a __copy__ of the value of url, so it can't be changed

		// otherwise, it all writes to the same copy of url, which overwrites
		// each other

		// solve race condition with channels -- controls the timing of each write into
		// the results map, ensuring that it happens one at a time (in the for loop below)
		go func(u string) {
			// Send statement
			resultChannel <- result{url: u, isUp: wc(u)}
		}(url)
	}

	// iterate through each url, get the value from the channel and assign it to the results map
	for i := 0; i < len(urls); i++ {
		// Receive expression -- assigns a value received from a channel to a variable
		r := <-resultChannel
		results[r.url] = r.isUp

	}

	return results
}
