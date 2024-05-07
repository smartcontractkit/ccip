package main

import (
	"flag"
	"strings"
)

type CommandLineArgs struct {
	Sender    string
	First     int
	Offset    int
	Receiver  string
	Source    string
	Dest      string
	FeeToken  string
	MessageId string
}

// Helper function to translate network names
func translateNetworkName(shorthand string) string {
	if fullName, exists := networkNameMapping[shorthand]; exists {
		return fullName
	}
	return shorthand // Return the input if no mapping found
}

func ParseFlags() CommandLineArgs {
	args := CommandLineArgs{}
	// Define temporary variables to hold the pointer values
	sender := flag.String("sender", "", "Sender address")
	first := flag.Int("first", 100, "First N results")
	offset := flag.Int("offset", 0, "Offset for results")
	receiver := flag.String("receiver", "", "Receiver address")
	source := flag.String("source", "", "Source network name")
	dest := flag.String("dest", "", "Destination network name")
	feeToken := flag.String("feeToken", "", "Fee Token")
	messageId := flag.String("messageId", "", "Message ID")

	flag.Parse()

	// Translate shorthand network names to full names
	fullSource := translateNetworkName(strings.ToUpper(*source))
	fullDest := translateNetworkName(strings.ToUpper(*dest))

	// Assign and convert to lowercase where applicable
	args.Sender = strings.ToLower(*sender)
	args.First = *first
	args.Offset = *offset
	args.Receiver = strings.ToLower(*receiver)
	args.Source = strings.ToLower(fullSource)
	args.Dest = strings.ToLower(fullDest)
	args.FeeToken = strings.ToLower(*feeToken)
	args.MessageId = strings.ToLower(*messageId)
	
	return args
}
