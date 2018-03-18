package main

import (
	"fmt"
	"log"

	"google.golang.org/api/calendar/v3"
)

func CreateDefaultEvent(srv *calendar.Service) {
	log.Println("creating a hello world event...")

	event := &calendar.Event{
		Summary:     "Yoga",
		Location:    "My Yoga Mat",
		Description: "Set up by ScheduleMaker",
		Start: &calendar.EventDateTime{
			DateTime: "2018-03-18T16:30:00-04:00",
			TimeZone: "America/New_York",
		},
		End: &calendar.EventDateTime{
			DateTime: "2018-03-18T17:00:00-04:00",
			TimeZone: "America/New_York",
		},
	}

	calendarID := "primary"
	event, err := srv.Events.Insert(calendarID, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}
	fmt.Printf("Event created: %s\n", event.HtmlLink)
}
