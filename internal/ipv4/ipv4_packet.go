package ipv4

type ipv4 struct{
	version uint8 //first 4 bits, is always 4 for ipv4
	IHL uint8 //second 4 bits, minimum value 5, max value 15
	DSCP byte
	ENC byte
	TotalLength int
	Identification uint8
	Flags int
	TTL int
	Protocol int
	SAddress string
	DAddress string
	Options int
}
const (
	version_size = 4
	IHL_max_byte_size = 60
	IHL_minimum_byte_size = 15
	ENC_TYPES_byte_max = 3
	ENC_TYPES_byte_min = 0
)
func validate (packet *ipv4) bool {
	return 
	packet == nil || 
	packet.version != version_size ||
	packet.ENC < ENC_TYPES_byte_min ||
	packet.ENC > ENC_TYPES_byte_max ||
	packet.IHL > IHL_max_byte_size ||
	packet.IHL < IHL_minimum_byte_size

}
func sanitize(packet *ipv4) string{
	if(packet == nil){
		return ""
	}

	if(packet)
}