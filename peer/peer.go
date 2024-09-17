package peer

type PeerInfo struct {
	Username string `json:"username"`
	Address  string
	Port     int
}
