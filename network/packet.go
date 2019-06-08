package network

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

type Packet struct {
	Buffer []uint8
	Cursor uint16
}

func NewPacket() *Packet {
	var p Packet
	p.Buffer = make([]uint8, 2)
	p.Cursor = 2
	return &p
}

func (p *Packet) ReadUint8() uint8 {
	v := p.Buffer[p.Cursor]
	p.Cursor++
	return v
}

func (p *Packet) ReadUint16() uint16 {
	v := binary.LittleEndian.Uint16(p.Buffer[p.Cursor : p.Cursor+2])
	p.Cursor += 2
	return v
}

func (p *Packet) ReadUint32() uint32 {
	v := binary.LittleEndian.Uint32(p.Buffer[p.Cursor : p.Cursor+4])
	p.Cursor += 4
	return v
}

func (p *Packet) ReadString() string {
	var str string
	strlen := p.ReadUint16()
	for i := (uint16)(0); i < strlen; i++ {
		str += (string)(p.ReadUint8())
	}
	return str
}

func (p *Packet) WriteUint8(v uint8) {
	p.Buffer = append(p.Buffer, v)
	binary.LittleEndian.PutUint16(p.Buffer[0:2], (uint16)(len(p.Buffer)-2))
	p.Cursor++
}

func (p *Packet) WriteUint16(v uint16) {
	bytes := make([]uint8, 2)
	binary.LittleEndian.PutUint16(bytes, v)
	p.WriteUint8(bytes[0])
	p.WriteUint8(bytes[1])
}

func (p *Packet) WriteString(str string) {
	p.WriteUint16((uint16)(len(str)))
	for i := 0; i < len(str); i++ {
		p.WriteUint8((uint8)(str[i]))
	}
}

func (p *Packet) Length() uint16 {
	return binary.LittleEndian.Uint16(p.Buffer[0:2])
}

func (p *Packet) SkipBytes(n uint16) {
	p.Cursor += n
}

func (p *Packet) HexDump(prefix string) {
	fmt.Printf("\n[%s]\n%s", prefix, hex.Dump(p.Buffer))
}
