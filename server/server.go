package main

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net"
	"sync"
	"time"

	pb "example.com/go-tokenservice-grpc/proto"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
)



type TokenServiceServer struct {
	pb.UnimplementedTokenServiceServer
}


type Token struct{
	
	Id string
	Name string
	Domain [3]uint64
	State [2]uint64
	Writer string
	Reader []string
	WTimeStamp uint64
	mu sync.RWMutex
	
}

var req_id uint64 = 0
var port_address string
var reading bool = false

// type UnMarshalToken struct{
// 	Id string
// 	Name string
// 	Domain [3]uint64
// 	State [2]uint64
// }

// type UnMarshalDomain struct{
// 	Low uint64
// 	Mid uint64
// 	High uint64
// }

// type UnMarshalState struct{
// 	Partial_value uint64
// 	Final_value uint64
// }


var database []Token

func (r *TokenServiceServer) Create(ctx context.Context, req  *pb.IdRequest) (*pb.AckResponse, error){
	
	var id = req.GetId()
	for _, val := range(database){
		if id == val.Id{
			val.mu.Lock()
			defer val.mu.Unlock()

			// val.Domain[0] = 0
			// val.Domain[1] = 0
			// val.Domain[2] = 0
			// val.Name = ""
			val.State[0] = 0
			val.State[1] = 0
			return &pb.AckResponse{Ack: "Token Created"}, nil
		}
	}
	token := Token{
		Id: req.GetId(),
	}
	database = append(database, token)

	log.Printf("Token Created : ")
	created_token := fmt.Sprintf("%+v", token)
	log.Printf(created_token)
	for _ , val := range(database){
		log.Printf("%v",val.Id)
	}
	
	return &pb.AckResponse{Ack : "Token Created : " + created_token}, nil

}


func (r *TokenServiceServer) Drop(ctx context.Context, req  *pb.IdRequest) (*pb.AckResponse, error){
	var id = req.GetId()
	var token Token

	for idx, val := range database {
		if val.Id == id{
			token = val
			token.mu.Lock()
			defer token.mu.Unlock()
			
			database = append(database[:idx], database[idx+1:]...)
			log.Println("Token Deleted : ")
			deleted_token := fmt.Sprintf("%+v", token)
			log.Printf(deleted_token)
			for _ , val := range(database){
				log.Printf("%v",val.Id)
			}
			return &pb.AckResponse{Ack: "Token Deleted : " + deleted_token}, nil
			
		}
	}
	
	return &pb.AckResponse{Ack: "No such Token Exists!"}, nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func (r *TokenServiceServer) Read(ctx context.Context, req  *pb.IdRequest) (*pb.AckResponse, error){
	// The read request from the client hits this endpoint first.
	var id = req.GetId()
	reading = true
	

	for i, val := range database {
		// If the id of the token matches and the server has the authorization to read the token
		if (val.Id == id) && ((contains(val.Reader, port_address) || val.Writer == port_address)){

			val.mu.RLock()
			defer val.mu.RUnlock()
			// var result uint64 = math.MaxUint64
			// for i := val.Domain[1]; i < val.Domain[2]; i++ {
			// 	cmp_int := Hash(val.Name, i)
			// 	if result > cmp_int{
			// 		result = cmp_int
			// 	}
			// }

			// val.State[0] = result
			// val.State[1] = result
			


			// obj_str := fmt.Sprintf("%+v", val.State[1])

	
			log.Println("Token Read : ")
			read_token := fmt.Sprintf("%+v", val)
			log.Printf(read_token)
			for _ , val := range(database){
				log.Printf("%v",val.Id)
			}

			req_id += 1		// Maintaining a unique req_id for each request
			ch := make(chan pb.RBroadcastReturn)
			var ack_arr []pb.RBroadcastReturn	// Maintaining an array of acknowledgements

			// for idx := 0; idx < 2; idx++ {
			// 	go RBroadcast(id, database[i].Reader[idx],  ch, req_id)
			// }

			for _, node := range(database[i].Reader){	//Broadcasting to each and every reader node asynchronously
				// if ack_cnt > 1{
				// 	log.Printf("Hits Break")
				// 	break
				// }
				// log.Printf("Broadcasted : %v", node)
				go RBroadcast(id, node,  ch, req_id)
				
				}
			log.Printf("Read Broadcasted!")
			
			for {	// Once an ack is available from broadcast, we append it to the array
			val := <- ch
			ack_arr = append(ack_arr, val)
			// log.Printf("Acknowledgement Array : %v", ack_arr)

			if len(ack_arr) > 1{	// When the number of acks received establish majority
					break
				// go func() {
				// 	WBroadcast(i, node)
				// 	wg.Done()
				// }()
			}	
		}
		var max_write uint64 = 0
		var max_token pb.RBroadcastReturn
		flag := 0
					
					for _, token := range(ack_arr){	// Iterating through the acknowledgements to find the value corresponding to the latest timestamp
						if (token.Wts > max_write){
							flag = 1
							max_write = token.Wts
							max_token = token
						}
					}
			ack_arr = nil	// Clearing the acknowledgement array for next requests
			if flag == 1{	// If we find the latest value written to the system, we write-back to the system again
					//Writeback
					writeback_req := pb.WriteRequest{
						Id: max_token.Id,
						Name: max_token.Name,
						Low: max_token.Low,
						Mid: max_token.Mid,
						High: max_token.High,
						Wts: max_token.Wts,
					}
					log.Printf("WriteBack Trigerred")
					r.Write(ctx, &writeback_req)
			}
			// log.Printf("Flag was 0  %v", max_token)

			
		
		
			return &pb.AckResponse{Ack: fmt.Sprintf("Token Read : %v" , max_token)}, nil

			
			}
		
			
		}
	

	return &pb.AckResponse{Ack: "No such Token Exists!"}, nil

}

func RBroadcast(id string, node string, ch chan pb.RBroadcastReturn, req_id uint64){
	// This is the actual goroutine that makes a broadcast call.


	// Establish a connection with the reader nodes
	var conn *grpc.ClientConn
		conn, err := grpc.Dial(node, grpc.WithInsecure())
	
		if err != nil{
			log.Fatalf("Could not Dial: %s", err)
		}
		defer conn.Close()

		c:= pb.NewTokenServiceClient(conn)

		read_broadcast := pb.RBroadcastMessage{
			Rid: req_id,
			Id: id,
		}

		// log.Printf("%v", write_broadcast)
		// Broadcast call
			res, err := c.ReadOne(context.Background(), &read_broadcast)
			if err != nil{
				log.Fatalf("Error calling %s", err)
			}
			// log.Printf("Response from Server: %s", res.Ack)
			// Once the broadcast responds back, the request id is ensured.
			if res.Rid == req_id{
			// log.Printf("Res from Readone : %v", res)	
			ch <- *res 		// Putting the ack in the channel
			// log.Printf("In WB : %v", *ack_cnt)
			}
}

func (r *TokenServiceServer) ReadOne(ctx context.Context, req *pb.RBroadcastMessage) (*pb.RBroadcastReturn, error) {
	// Responds to the broadcast calls by passing value of the token as ack
	// Sleep to simulate crashing of the node
	if port_address == "127.0.0.32:8002"{
		log.Printf("Sleeping")
		time.Sleep(100 * time.Second)
	
	}

	var id = req.GetId()

	for _, val := range database {
		if val.Id == id{
			// log.Printf("Entered : %v", id)
		
			val.mu.RLock()
			defer val.mu.RUnlock()			

			//obj_str := fmt.Sprintf("%+v", val.State[1])
	
			log.Println("Token Read : ")
			read_token := fmt.Sprintf("%+v", val)
			log.Printf(read_token)
			for _ , val := range(database){
				log.Printf("%v",val.Id)
			}

			// Returning the token as an ack
			return &pb.RBroadcastReturn{
				Id: val.Id,
				Low: val.Domain[0],
				Mid: val.Domain[1],
				High: val.Domain[2],
				Name: val.Name, 
				Rid: req.GetRid(),
				Wts: val.WTimeStamp,
				Val: val.State[0],
		}, nil
	}
}
	return &pb.RBroadcastReturn{
		Id: "",
	}, nil
}



func Hash(name string, nonce uint64) uint64 {
	hasher := sha256.New()
	hasher.Write([]byte(fmt.Sprintf("%s %d", name, nonce)))
	return binary.BigEndian.Uint64(hasher.Sum(nil))
}
	

func (r *TokenServiceServer) Write(ctx context.Context, req  *pb.WriteRequest) (*pb.AckResponse, error){
	// This is the endpoint that is hit when the client issues a write request.
	id := req.GetId()
	// log.Printf("In Write : " + id)

	
	// for _, val := range database{
	for i := 0; i < len(database); i++{ 
		if (database[i].Id == id){	// Ensuring that the server is authorized to write to the token
			if database[i].Writer == port_address{
			database[i].mu.Lock()
			defer database[i].mu.Unlock()
			database[i].Name = req.GetName()	// Writing to the token
			database[i].Domain[0] = req.GetLow()
			database[i].Domain[1] = req.GetMid()
			database[i].Domain[2] = req.GetHigh()
			
			// log.Printf("Waiting for 20s : ")
			// time.Sleep(10 * time.Second)

			var result uint64 = math.MaxUint64

			for i := req.GetLow(); i < req.GetMid(); i++ {	// Calculating the partial value immediately after a write call
				cmp_int := Hash(req.GetName(), i)
				if result > cmp_int{
					result = cmp_int
				}
			}
			database[i].State[0] = result
			if reading == false{	// We increment the timestamp only if we encounter a write request
			database[i].WTimeStamp += 1
			}else{	// Copying the latest timestamp value obtained from the read broadcast
			database[i].WTimeStamp = req.GetWts()
			}
			
			obj_str := fmt.Sprintf("%+v", database[i])

			log.Println("Token Written : ")
			write_token := fmt.Sprintf("%+v", database[i])
			log.Printf(write_token)
			for _ , val := range(database){
				log.Printf("%v",val.Id)
			}
			req_id += 1
			
			ch := make(chan int)
			ack_cnt := 0
			// for idx := 0; idx < 2; idx++ {
			// 	go WBroadcast(i, database[i].Reader[idx], &ack_cnt, ch, req_id)
				
			// }
			for _, node := range(database[i].Reader){// Broadcasting to the reader nodes
				// if ack_cnt > 1{
				// 	log.Printf("Hits Break")
				// 	break
				// }
				go WBroadcast(i, node, &ack_cnt, ch, req_id)
				
				}
			log.Printf("Write Broadcasted!")
			for {		
			val := <- ch		// Waiting for the acknowledgements
			log.Printf("Value : %v", val)
			if val > 1{			// If the acknowledgements are greater than N/2 then we break
					// log.Printf("Hits Break")
					break
					
				// go func() {
				// 	WBroadcast(i, node)
				// 	wg.Done()
				// }()
			}	
		}
		if (reading == true){		// For read request, we return the appropriate value
			reading = false
			return &pb.AckResponse{Ack: "Token Read : " + write_token}, nil

		}else{
		return &pb.AckResponse{Ack: "Token Written : " +  obj_str}, nil
		}

		}else{
			return &pb.AckResponse{Ack: "Unauthorized to write to the token"}, nil
		}

	}
	}

	return &pb.AckResponse{Ack: "No such Token Exists!"}, nil
	
}


func WBroadcast(i int, node string, ack_cnt *int, ch chan int, req_id uint64){
		// log.Printf("In WBroadcast")
		// This goroutine is a broadcast function for the write operation

		// Initializing the grpc call
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(node, grpc.WithInsecure())
	
		if err != nil{
			log.Fatalf("Could not Dial: %s", err)
		}
		defer conn.Close()

		c:= pb.NewTokenServiceClient(conn)


		//Broadcasting the newly written token
		write_broadcast := pb.WBroadcastMessage{
			Id: database[i].Id,
			Low: database[i].Domain[0],
			Mid: database[i].Domain[1],
			High: database[i].Domain[2],
			Name: database[i].Name, 
			Rid: req_id,
			Wts: database[i].WTimeStamp,
			Val: database[i].State[0],
		}
		// log.Printf("%v", write_broadcast)
		//Making the broadcast call
			res, err := c.WriteOne(context.Background(), &write_broadcast)
			if err != nil{
				log.Fatalf("Error calling %s", err)
			}
			// log.Printf("Response from Server: %s", res.Ack)

			if res.Rid == req_id{	
			*ack_cnt = *ack_cnt + 1
			// log.Printf("%v", *ack_cnt)
			ch <- *ack_cnt 	
			// log.Printf("In WB : %v", *ack_cnt)
			}

	}


func (r *TokenServiceServer) WriteOne(ctx context.Context, req *pb.WBroadcastMessage) (*pb.WBroadcastReturn, error) {
	// if port_address == "127.0.0.32:8002"{
	// 	log.Printf("Sleeping....")
	// 	time.Sleep(1000 * time.Second)
	// }
	// Writes to a single node
	// This function acts as a receiver for the broadcast message
	

	id := req.GetId()
	// log.Printf("In WriteOne : " + id)
	obj_str := "None"

	
	// for _, val := range database{
	for i := 0; i < len(database); i++{ 
		if database[i].Id == id{
			// log.Printf("Write-back attempted")
			if req.GetWts() > database[i].WTimeStamp{
				database[i].mu.Lock()
				defer database[i].mu.Unlock()
				database[i].Name = req.GetName()
				database[i].Domain[0] = req.GetLow()
				database[i].Domain[1] = req.GetMid()
				database[i].Domain[2] = req.GetHigh()
				database[i].WTimeStamp = req.GetWts()
				// log.Printf("Waiting for 20s : ")
				// time.Sleep(10 * time.Second)

				database[i].State[0] = req.GetVal()

				obj_str = fmt.Sprintf("%+v", database[i])

				log.Println("Token Written : ")
				write_token := fmt.Sprintf("%+v", database[i])
				log.Printf(write_token)
				for _ , val := range(database){
					log.Printf("%v",val.Id)
				}	
		}
		return &pb.WBroadcastReturn{Ack: "Token Written : " +  obj_str, Rid: req.GetRid()}, nil
	}

	}
	return &pb.WBroadcastReturn{Ack: "No such Token Exists!", Rid: req.GetRid()}, nil

}



func main(){
	ipPtr := flag.String("ip", "127.0.0.1", "an ip address")
	portPtr := flag.String("port", "9000", "a port number")
	flag.Parse()
	port := *ipPtr + ":" + *portPtr
	port_address = port
	lis, err := net.Listen("tcp",  port)
	if err != nil {
		log.Fatalf("Failed to listen on port " +  port + "%v", err)
	}


	sample_tokens, err := ioutil.ReadFile("sample_tokens.yaml")

	if err != nil {

		 log.Fatal(err)
	}

	var tokens []Token

	err2 := yaml.Unmarshal(sample_tokens, &tokens)

	if err2 != nil {

		 log.Fatal(err2)
	}

	// log.Println(tokens)
	for _ , token := range(tokens){
		if token.Writer == port{
			database = append(database, token)
		}
		for _, val := range(token.Reader){
			if val == port{
				database = append(database, token)
			}
		}
	}
	log.Println(database)

	

	grpcServer := grpc.NewServer()
	
	pb.RegisterTokenServiceServer(grpcServer, &TokenServiceServer{})
	log.Printf("Server listening at port " +  port )

	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("Failed to serve gRPC server over port " + port + "%v", err)

	}
	
}
