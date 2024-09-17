package main

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"

	"github.com/DinnieJ/p2p-chat/peer"
)

type PeersRegistry struct {
	mu    sync.Mutex
	peers map[string]*peer.PeerInfo
}

func (p *PeersRegistry) AddPeerFromByte(data []byte, peerAddr *net.UDPAddr) error {

	userData := &peer.PeerInfo{}
	err := json.Unmarshal(data, userData)
	if err != nil {
		return err
	}
	newPeer := &peer.PeerInfo{
		Username: userData.Username,
		Address:  peerAddr.IP.String(),
		Port:     peerAddr.Port,
	}
	p.mu.Lock()
	p.peers[newPeer.Address] = newPeer
	p.mu.Unlock()
	return nil
}

var peers *PeersRegistry = &PeersRegistry{}

func handleIncomingPeer(buffer []byte, peerAddr *net.UDPAddr) {
	if err := peers.AddPeerFromByte(buffer, peerAddr); err == nil {
		fmt.Println("Failed to add peer")
		return
	}
}

func main() {
	addr := &net.TCPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: 44321,
	}

	listener, err := net.ListenTCP("tcp", addr)

	if err != nil {
		panic("Failed to start relay server")
	}

	defer conn.Close()

	fmt.Println("Start relay server at port 44321")
	for {
		buffer := make([]byte, 1024)
		conn, err := conn.AcceptTCP()

		if err != nil {
			fmt.Println("Failed to read from peer: ", err)
			continue
		}
		conn.Read(buffer)
		go handleIncomingPeer(buffer, conn 1)
	}

}
