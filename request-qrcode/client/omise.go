package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chanchai-maqe/grpc-go-course/request-qrcode/qrcodepb"
	"google.golang.org/grpc"
)

const endpoint = "localhost:50051"

func main() {

	fmt.Print("This message is sent from Omise-client\n")
	/*
		Uses WithInsecure() b/c it is the testing and there is no SSL
		required at this moment.
	*/
	cc, err := grpc.Dial(endpoint, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	/*cc will close after greetpb.NewGreetServiceClient*/
	defer cc.Close()

	c := qrcodepb.NewQrcodeServiceClient(cc)

	//fmt.Printf("Created client %v :", c)
	fmt.Println()

	doUnary(c)

}

func doUnary(c qrcodepb.QrcodeServiceClient) {
	fmt.Println("Starting Unary RPC from Omise client ...")
	req := &qrcodepb.QrcodeRequest{
		/*Request Message*/
		Qrcode: &qrcodepb.QRCode{
			BizMchId:     "000001",
			MchId:        "001",
			BillerId:     "1000000001",
			Rsapublickey: "asadi2j9fljK#&$HJIDPVKSNVLKSDNVL(C)581628368127DLKLDVLKSDLFKJKDKj;clPIDLJPGUJELKJGIDUOIJW*(*93248320948092504t-0ifpodjfpoiut03ht09240942u0t492u thisis@example.com ",
			Apikey:       "788899kdjk8ok32jo09uvojpy2g38i3",
		},
	}

	res, err := c.Hi(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Hi function from Omise: %v", err)
	}

	log.Printf("Response : %v", res.Result)
	log.Println("SHOW RESULT:...")
}
