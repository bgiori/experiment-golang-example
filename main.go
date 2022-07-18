package main

import (
	"fmt"
	"github.com/amplitude/experiment-go-server/pkg/experiment"
	"github.com/amplitude/experiment-go-server/pkg/experiment/local"
	"github.com/amplitude/experiment-go-server/pkg/experiment/remote"
)

func main() {
	remoteEvaluation()
	localEvaluation()
}

func remoteEvaluation() {
	// (1) Initialize the local evaluation client with a server deployment key.
	client := remote.Initialize("server-VY0FufBsdITI1Gv9y7RyUopLzk9m8t0n", &remote.Config{Debug: true})

	// (2) Fetch variants for a user
	user := &experiment.User{
		UserId:   "user@company.com",
		DeviceId: "abcdefg",
		UserProperties: map[string]interface{}{
			"premium": true,
		},
	}
	variants, err := client.Fetch(user)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", variants)

}

func localEvaluation() {
	client := local.Initialize("server-VY0FufBsdITI1Gv9y7RyUopLzk9m8t0n", &local.Config{Debug: true})
	err := client.Start()
	if err != nil {
		panic(err)
	}
	user := &experiment.User{UserId: "user@company.com"}
	variants, err := client.Evaluate(user, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", variants)
}
