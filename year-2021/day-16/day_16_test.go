package day_16

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test16Part1(t *testing.T) {
	file, _ := os.Open("day_16.in")
	defer file.Close()
	bytes, _ := io.ReadAll(file)

	bin := convertToBin(string(bytes))
	packets := decode(bin, 0, len(bin), true)

	sum := 0
	for _, p := range flatPackets(packets) {
		sum += p.version
	}
	assert.Equal(t, 984, sum)
}

func Test16Part2(t *testing.T) {
	file, _ := os.Open("day_16.in")
	defer file.Close()
	bytes, _ := io.ReadAll(file)

	bin := convertToBin(string(bytes))
	pockets := decode(bin, 0, len(bin), true)
	assert.Equal(t, int64(1015320896946), eval(pockets[0]))
}

func convertToBin(str string) string {
	sb := strings.Builder{}
	for _, r := range str {
		num, _ := strconv.ParseInt(string(r), 16, 32)
		sb.WriteString(fmt.Sprintf("%04b", num))
	}
	return sb.String()
}

type packet struct {
	header
	literal
	operator
}

type header struct {
	version, typeId int
}

type literal struct {
	num int64
}

type operator struct {
	subPackets []packet
}

func decode(bin string, idx int, length int, applyPadding bool) []packet {
	packets := make([]packet, 0)
	for idx < length {
		p, newIdx := decodePacket(bin, idx, applyPadding)
		packets = append(packets, p)
		idx = newIdx
	}
	return packets
}

func decodePacket(bin string, idx int, applyPadding bool) (packet, int) {
	p := packet{}

	version, typeId := decodeHeader(bin, idx)
	p.version = version
	p.typeId = typeId
	idx += 6

	switch typeId {
	case 4:
		{
			str := ""
			for grp := "1"; grp[0] != '0'; idx += 5 {
				grp = bin[idx : idx+5]
				str += grp[1:]
			}
			num, _ := strconv.ParseInt(str, 2, 64)
			p.num = num
		}
	default:
		{
			lengthTypeId := bin[idx] - '0'
			idx += 1

			if lengthTypeId == 0 {
				totalLengthInBits, _ := strconv.ParseInt(bin[idx:idx+15], 2, 32)
				idx += 15

				p.subPackets = decode(bin, idx, idx+int(totalLengthInBits), false)
				idx += int(totalLengthInBits)
			} else {
				numberOfSubPackets, _ := strconv.ParseInt(bin[idx:idx+11], 2, 32)
				idx += 11

				packets := make([]packet, 0)
				for n := 0; n < int(numberOfSubPackets); n++ {
					pack, newIdx := decodePacket(bin, idx, false)
					packets = append(packets, pack)
					idx = newIdx
				}
				p.subPackets = packets
			}
		}
	}

	if applyPadding {
		idx = len(bin)
	}

	return p, idx
}

func decodeHeader(bin string, idx int) (int, int) {
	version, _ := strconv.ParseInt(bin[idx:idx+3], 2, 32)
	typeId, _ := strconv.ParseInt(bin[idx+3:idx+6], 2, 32)
	return int(version), int(typeId)
}

func flatPackets(xs []packet) []packet {
	result := make([]packet, 0)
	for _, p := range xs {
		result = append(result, p)
		if p.subPackets != nil && len(p.subPackets) > 0 {
			result = append(result, flatPackets(p.subPackets)...)
		}
	}
	return result
}

func eval(p packet) int64 {
	switch p.typeId {
	case 0:
		{
			result := int64(0)
			for _, sp := range p.subPackets {
				result += eval(sp)
			}
			return result
		}
	case 1:
		{
			result := int64(1)
			for _, sp := range p.subPackets {
				result *= eval(sp)
			}
			return result
		}
	case 2:
		{
			result := int64(9223372036854775807)
			for _, sp := range p.subPackets {
				val := eval(sp)
				if val < result {
					result = val
				}
			}
			return result
		}
	case 3:
		{
			result := int64(-9223372036854775808)
			for _, sp := range p.subPackets {
				val := eval(sp)
				if val > result {
					result = val
				}
			}
			return result
		}
	case 4:
		{
			return p.num
		}
	case 5:
		{
			if eval(p.subPackets[0]) > eval(p.subPackets[1]) {
				return 1
			}
			return 0
		}
	case 6:
		{
			if eval(p.subPackets[0]) < eval(p.subPackets[1]) {
				return 1
			}
			return 0
		}
	case 7:
		{
			if eval(p.subPackets[0]) == eval(p.subPackets[1]) {
				return 1
			}
			return 0
		}
	default:
		{
			panic("invalid type")
		}
	}
}
