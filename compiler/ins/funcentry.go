package ins

type FuncEntry struct {
	rtm uint16
	vars uint8
}

func (p *FuncEntry)Bytes() []byte {
	return []byte{0x31, byte(p.rtm), byte(p.rtm>>8), p.vars}
}

func FuncEntryNew() *FuncEntry {
	return &FuncEntry{
		rtm:  5,
		vars: 0,
	}
}