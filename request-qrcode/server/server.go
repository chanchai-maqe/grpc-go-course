package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/chanchai-maqe/grpc-go-course/request-qrcode/qrcodepb"
	"google.golang.org/grpc"
)

type server struct{}

/*
	Add Hi() function to *server aka *server implement QRcodeServiceServer
*/
func (*server) Hi(ctx context.Context, req *qrcodepb.QrcodeRequest) (*qrcodepb.QrcodeResponse, error) {

	fmt.Printf("\nHi function is invoked :%v\n", req)
	bizMchID := req.GetQrcode().GetBizMchId()
	mchID := req.GetQrcode().GetMchId()
	billerID := req.GetQrcode().GetBillerId()
	rsapublickey := req.GetQrcode().GetRsapublickey()
	apikey := req.GetQrcode().GetApikey()

	result := "\nbizMchId : " + bizMchID + "\n" +
		"mchId : " + mchID + "\n" +
		"billerId : " + billerID + "\n" +
		"rsapublickey : " + rsapublickey + "\n" +
		"apikey : " + apikey + "\n"

	return &qrcodepb.QrcodeResponse{
		Result: result,
	}, nil
}

func main() {
	fmt.Println("GRPC Server is running")

	// NOTE: 50051 is default port for grpc
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// create grpc server variable as NewServer()
	s := grpc.NewServer()

	// register server to greetpb variable
	qrcodepb.RegisterQrcodeServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
