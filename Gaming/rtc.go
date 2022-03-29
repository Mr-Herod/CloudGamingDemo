package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"

	"github.com/Mr-Herod/CloudGamingDemo/Gaming/config"

	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/examples/internals/signal"
)

func RTC(playinfo Playinfo, des string, desChan chan string) {
	fmt.Printf("rtc() working\n")
	// Create a new RTCPeerConnection
	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{config.ServiceConf.StunServer},
			},
		},
	})
	if err != nil {
		fmt.Printf("NewPeerConnection Error:%v\n", err)
		return
	}
	defer func() {
		//if cErr := peerConnection.Close(); cErr != nil {
		fmt.Printf("end of peerConnection \n")
		//}
	}()
	portInt := int(time.Now().Unix()%1000 + 5000)
	portStr := strconv.Itoa(portInt)
	go StartGame(playinfo, portStr)
	_, iceConnectedCtxCancel := context.WithCancel(context.Background())
	// Set the handler for ICE connection state
	// This will notify you when the peer has connected/disconnected
	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("Connection State has changed %s \n", connectionState.String())
		if connectionState == webrtc.ICEConnectionStateConnected {
			iceConnectedCtxCancel()
		}
	})

	// Set the handler for Peer connection state
	// This will notify you when the peer has connected/disconnected
	peerConnection.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		fmt.Printf("Peer Connection State has changed: %s\n", s.String())

		if s == webrtc.PeerConnectionStateFailed {
			// Wait until PeerConnection has had no network activity for 30 seconds or another failure. It may be reconnected using an ICE Restart.
			// Use webrtc.PeerConnectionStateDisconnected if you are interested in detecting faster timeout.
			// Note that the PeerConnection may come back from PeerConnectionStateDisconnected.
			endGame(playinfo, portStr)
			fmt.Println("Peer Connection has gone to failed exiting")
			return
		}
	})
	nowPos := 1
	// Register data channel creation handling
	score := 1000
	startTime := int64(0)
	moveCnt := 0
	peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
		fmt.Printf("New DataChannel %s %d\n", d.Label(), d.ID())
		if startTime == 0 {
			startTime = time.Now().Unix()
		}
		// Register text message handling
		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Printf("Message from DataChannel '%s': '%s'\n", d.Label(), string(msg.Data))
			if string(msg.Data) == " " {
				score = 1000
				moveCnt = 0
				startTime = time.Now().Unix()
			}
			nextPos, win := Move(nowPos, string(msg.Data), portStr)
			nowPos = nextPos
			moveCnt += 1
			if win {
				timeCost := int(time.Now().Unix() - startTime)
				score = score - timeCost*17 - moveCnt*13
				if score <= 0 {
					score = 1
				}
				saveRecord(playinfo, score)
				sendErr := d.SendText(strconv.Itoa(score))
				if sendErr != nil {
					fmt.Printf("SendText Error:%v\n", err)
					return
				}
			}
		})
	})

	// Open a UDP Listener for RTP Packets on port 5004
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("localhost"), Port: portInt})
	if err != nil {
		fmt.Printf("ListenUDP Error:%v\n", err)
		return
	}
	defer func() {
		if err = listener.Close(); err != nil {
			fmt.Printf("listener.Close Error:%v\n", err)
			return
		}
	}()

	// Create a video track
	videoTrack, err := webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeVP8}, "video", "pion")
	if err != nil {
		fmt.Printf("NewTrackLocalStaticRTP Error:%v\n", err)
		return
	}
	rtpSender, err := peerConnection.AddTrack(videoTrack)
	if err != nil {
		fmt.Printf("peerConnection.AddTrack Error:%v\n", err)
		return
	}

	// Read incoming RTCP packets
	// Before these packets are returned they are processed by interceptors. For things
	// like NACK this needs to be called.
	go func() {
		rtcpBuf := make([]byte, 1500)
		for {
			if _, _, rtcpErr := rtpSender.Read(rtcpBuf); rtcpErr != nil {
				fmt.Printf("rtcp readError:%v\n", rtcpErr)
			}
		}
	}()

	// Wait for the offer to be pasted
	offer := webrtc.SessionDescription{}
	signal.Decode(des, &offer)

	// Set the remote SessionDescription
	if err = peerConnection.SetRemoteDescription(offer); err != nil {
		fmt.Printf("peerConnection.SetRemoteDescription Error:%v\n", err)
		return
	}

	// Create answer
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		fmt.Printf("peerConnection.CreateAnswer Error:%v\n", err)
		return
	}

	// Create channel that is blocked until ICE Gathering is complete
	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	// Sets the LocalDescription, and starts our UDP listeners
	if err = peerConnection.SetLocalDescription(answer); err != nil {
		fmt.Printf("peerConnection.SetLocalDescription Error:%v\n", err)
		return
	}

	// Block until ICE Gathering is complete, disabling trickle ICE
	// we do this because we only can exchange one signaling message
	// in a production application you should exchange ICE Candidates via OnICECandidate
	<-gatherComplete

	// Output the answer in base64 so we can paste it in browser
	remoteDes := signal.Encode(*peerConnection.LocalDescription())
	fmt.Println("server des:", remoteDes[:10])
	desChan <- remoteDes
	// Block forever
	// Read RTP packets forever and send them to the WebRTC Client
	inboundRTPPacket := make([]byte, 1600) // UDP MTU
	for {
		n, _, err := listener.ReadFrom(inboundRTPPacket)
		if err != nil {
			fmt.Printf("error during read inboundRTPPacket Error:%v\n", err)
			return
		}

		if _, err = videoTrack.Write(inboundRTPPacket[:n]); err != nil {
			if errors.Is(err, io.ErrClosedPipe) {
				// The peerConnection has been closed.
				fmt.Printf("closed Pipe error,err:%v\n", err)
				return
			}
			fmt.Printf("videoTrack.Write Error:%v\n", err)
			return
		}
	}
}
