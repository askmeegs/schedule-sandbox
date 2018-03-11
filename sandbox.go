package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	calendar "google.golang.org/api/calendar/v3"
)

// return a read only google calendar API client for my account
func AuthedClient() (*calendar.Service, error) {
	// Contexts = easy passing of request-scoped values, signals, deadlines between goroutines
	ctx := context.Background()

	// Read my local key
	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := GetClient(ctx, config)
	calService, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve calendar Client %v", err)
	}
	return calService, nil
}

func main() {
	calService, err := AuthedClient()
	if err != nil {
		log.Fatalf("Could not instantiate client with auth: %v", err)
	}

	t := time.Now().Format(time.RFC3339)
	events, err := calService.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events. %v", err)
	}

	fmt.Println("Upcoming events:")
	if len(events.Items) > 0 {
		for _, i := range events.Items {
			var when string
			// If the DateTime is an empty string the Event is an all-day Event.
			// So only Date is available.
			if i.Start.DateTime != "" {
				when = i.Start.DateTime
			} else {
				when = i.Start.Date
			}
			fmt.Printf("%s (%s)\n", i.Summary, when)
		}
	} else {
		fmt.Printf("No upcoming events found.\n")
	}

}
