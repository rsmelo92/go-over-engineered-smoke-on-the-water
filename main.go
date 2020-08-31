package main

// TODO: Fetch from this api https://pipl.ir/v1/getPerson 

import (
	"fmt"
	"bufio"
	"net/http"
)

func another(w http.ResponseWriter, r *http.Request)  {
	fmt.Printf("another \n")

	reader := bufio.NewReader(r.Body)
	for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				panic(err)
			}
			fmt.Printf("Got %s \n", string(line))
	}

	// // Create a new RTCPeerConnection
	// peerConnection, err := webrtc.NewPeerConnection(config)
	// if err != nil {
	// 	panic(err)
	// }

	// if _, err = peerConnection.AddTransceiver(webrtc.RTPCodecTypeVideo); err != nil {
	// 	panic(err)
	// }

	// peerConnection.OnTrack(func(track *webrtc.Track, receiver *webrtc.RTPReceiver) {

	// 	fmt.Printf("Track has started, of type %d: %s \n", track.PayloadType(), track.Codec().Name)
	// 	fileName := fmt.Sprintf("assets/bettyboop.mp4")

	// 	fmt.Printf("Got Opus track, saving to disk as %s (48 kHz, 2 channels) \n", fileName)

	// 	for {
	// 		rtpPacket, err := track.ReadRTP()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		if err := oggFile.WriteRTP(rtpPacket); err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// })

}

func main()  {
	fs := http.FileServer(http.Dir("./static"))
  http.Handle("/", fs)
	http.HandleFunc("/another", another)

	fmt.Println("Server has started on http://localhost:8080")
	panic(http.ListenAndServe(":8080", nil))
}

// func main()  {
// 	const songsDir = "./assets/bettyboop.mp4"
	
// 	fs := http.FileServer(http.Dir("./static"))
// 	sd := http.FileServer(http.Dir(songsDir))
//   http.Handle("/", fs)
// 	http.Handle("/videos", addHeaders(sd))

// 	port := ":8080"
// 	fmt.Printf("Running server on http://localhost%s\n", port)
	
// 	err := http.ListenAndServe(port, nil)
//   if err != nil {
//     log.Fatal(err)
//   }
// }

// // addHeaders will act as middleware to give us CORS support
// func addHeaders(h http.Handler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		h.ServeHTTP(w, r)
// 	}
// }
