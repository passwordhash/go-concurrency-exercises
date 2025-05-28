//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, out chan<- Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			break
		}
		out <- *tweet
	}
	close(out)
}

func consumer(in <-chan Tweet) {
	for tweet := range in {
		if tweet.IsTalkingAboutGo() {
			fmt.Println(tweet.Username, "\ttweets about golang")
		} else {
			fmt.Println(tweet.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	tweets := make(chan Tweet)

	// Producer
	go producer(stream, tweets)

	// Consumer
	consumer(tweets)

	fmt.Printf("Process took %s\n", time.Since(start))
}
