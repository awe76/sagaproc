package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	pb "github.com/awe76/sagaproc/proto"
)

type Sagaproc struct{}

func (e *Sagaproc) HandleOperation(ctx context.Context, req *pb.OperationPayload, rsp *pb.OperationPayload) error {

	if req.IsRollback {
		fmt.Printf("%s operation rollback is started\n", req.Operation.Name)
	} else {
		fmt.Printf("%s operation is started\n", req.Operation.Name)
	}

	rand.Seed(time.Now().UnixNano())
	pause := rand.Intn(100)

	// sleep for some random time
	time.Sleep(time.Duration(pause) * time.Millisecond)

	payload, err := json.Marshal(rand.Float32())
	if err != nil {
		return err
	}

	rsp.Id = req.Id
	rsp.Operation = req.Operation
	rsp.Payload = string(payload)
	rsp.Name = req.Name
	rsp.IsRollback = req.IsRollback

	return nil
}
