package test

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"testing"
	"time"
)

const (
	ADDRESS = ":11111"
)

type StreamService struct {}

func init() {
	server := grpc.NewServer()
	RegisterStreamServiceServer(server, &StreamService{})
	lis, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		log.Printf("net.Listen error: %v", err)
	}
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Printf("server.Serve error: %v", err)
			return
		}
		log.Printf("grpc server start serving on %s", ADDRESS)
	}()
	go dispatch()
}

var ch chan int = make(chan int)

func dispatch() {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func (this *StreamService) Publish(r *StreamRequest, stream StreamService_PublishServer) error {
	for i := 0; i < 10; i++ {
		err := stream.Send(&StreamResponse{
			Pt: &StreamPoint{
				Name: r.Pt.Name,
				Value: int32(i),
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *StreamService) Queue(r *StreamRequest, stream StreamService_QueueServer) error {
	for {
		// channel: concurrency
		i := <-ch
		err := stream.Send(&StreamResponse{
			Pt: &StreamPoint{
				Name: r.Pt.Name,
				Value: int32(i),
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *StreamService) Record(stream StreamService_RecordServer) error {
	return nil
}

func (this *StreamService) Route(stream StreamService_RouteServer) error {
	return nil
}

func PrintPublish(number int, client StreamServiceClient, r *StreamRequest) error {
	stream, err := client.Publish(context.Background(), r)
	if err != nil {
		return err
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("resp: number = %d, name = %s, value = %d", number, resp.Pt.Name, resp.Pt.Value)
	}
	return nil
}

func PrintQueue(number int, client StreamServiceClient, r *StreamRequest) error {
	stream, err := client.Queue(context.Background(), r)
	if err != nil {
		return err
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("resp: number = %d, name = %s, value = %d", number, resp.Pt.Name, resp.Pt.Value)
	}
	return nil
}

func PrintRecord(number int, client StreamServiceClient, r *StreamRequest) error {
	return nil
}

func PrintRoute(number int, client StreamServiceClient, r *StreamRequest) error {
	return nil
}

func Client(number int) {
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Printf("grpc.Dial error: %v", err)
	}
	defer conn.Close()
	//
	client := NewStreamServiceClient(conn)
	// PrintQueue
	// PrintPublish
	err = PrintPublish(number, client, &StreamRequest{Pt: &StreamPoint{Name: "grpc stream client: list", Value: 2020}})
	if err != nil {
		log.Printf("print queue/publish error: %v", err)
	}
	//
	err = PrintRecord(number, client, &StreamRequest{Pt: &StreamPoint{Name: "grpc stream client: record", Value: 2020}})
	if err != nil {
		log.Printf("print record error: %v", err)
	}
	//
	err = PrintRoute(number, client, &StreamRequest{Pt: &StreamPoint{Name: "grpc stream client: route", Value: 2020}})
	if err != nil {
		log.Printf("print route error: %v", err)
	}
}

func TestStreamService(t *testing.T) {
	for i := 1; i < 3; i++ {
		go Client(i)
	}
	Client(0)

	time.Sleep(time.Minute)
}
