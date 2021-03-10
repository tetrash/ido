package events

import (
	"flag"
	"github.com/tetrash/event-recorder/internal"
	"strings"
	"time"
)

func Add(args []string) {
	var file string
	var tag string

	defaultFileName := "~/.event-record/" + internal.DefaultFileName()
	defaultTag := "cli"

	flag.StringVar(&file, "file", defaultFileName, "File path to log events, eg.: ~/events.log")
	flag.StringVar(&file, "f", defaultFileName, "File path to log events, eg.: ~/events.log")
	flag.StringVar(&tag, "tag", defaultTag, "Event tag, eg.: cli")
	flag.StringVar(&tag, "t", defaultTag, "Event tag, eg.: cli")

	flag.Parse()

	eventBody := strings.Join(args, " ")
	event := internal.CreateEvent(eventBody, time.Now(), &tag)
	internal.AppendFile(file, event)
}
