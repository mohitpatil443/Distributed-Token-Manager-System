package main

import (
	"context"
	"flag"
	"log"

	pb "example.com/go-tokenservice-grpc/proto"
	"google.golang.org/grpc"
)

func main(){
	createPtr := flag.Bool("create", false, "Create Token")
	dropPtr := flag.Bool("drop", false, "Drop Token")
	readPtr := flag.Bool("read", false, "Read Token")
	writePtr := flag.Bool("write", false, "Write Token")
	
	idPtr := flag.String("id", "101", "Token ID")
	namePtr := flag.String("name", "test_token", "Token Name")
	lowPtr := flag.Int("low", 0, "Token Low")
	midPtr := flag.Int("mid", 5, "Token Mid")
	highPtr := flag.Int("high", 10, "Token High")

	portPtr := flag.String("port", "50051", "Port Number")
	hostPtr := flag.String("host", "localhost", "Host Name")

	
	flag.Parse()

	port := *hostPtr + ":" + *portPtr 



	var conn *grpc.ClientConn
	conn, err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil{
		log.Fatalf("Could not Dial: %s", err)
	}
	defer conn.Close()

	c:= pb.NewTokenServiceClient(conn)





	if *createPtr{
		id_request := pb.IdRequest{
			Id: *idPtr,
		}
		res, err := c.Create(context.Background(), &id_request)
		if err != nil{
			log.Fatalf("Error calling %s", err)
		}
	
		log.Printf("Response from Server: %s", res.Ack)	
	}

	if *dropPtr{
		id_request := pb.IdRequest{
			Id: *idPtr,
		}
		res, err := c.Drop(context.Background(), &id_request)
		if err != nil{
			log.Fatalf("Error calling %s", err)
		}
	
		log.Printf("Response from Server: %s", res.Ack)	
	}

	if *writePtr{

		write_request := pb.WriteRequest{
		Id: *idPtr,
		Low: uint64(*lowPtr),
		Mid: uint64(*midPtr),
		High: uint64(*highPtr),
		Name: *namePtr,}

		res, err := c.Write(context.Background(), &write_request)
		if err != nil{
			log.Fatalf("Error calling %s", err)
		}
		log.Printf("Response from Server: %s", res.Ack)		

	}

	if *readPtr{
		id_request := pb.IdRequest{
			Id: *idPtr,
		}
		res, err := c.Read(context.Background(), &id_request)
		if err != nil{
			log.Fatalf("Error calling %s", err)
		}
	
		log.Printf("Response from Server: %s", res.Ack)	

	}
	
}