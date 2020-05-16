package binchunk

import (
	"encoding/binary"
	"math"
)

type reader struct {
	data []byte
}

func (self *reader) readByte() byte {
	b := self.data[0]
	self.data = self.data[1:]
	return b
}

func (self *reader) readUint32() uint32 {
	i := binary.LittleEndian.Uint32(self.data)
	self.data = self.data[4:]
	return i
}

func (self *reader) readUint64() uint64 {
	i := binary.LittleEndian.Uint64(self.data)
	self.data = self.data[8:]
	return i
}

func (self *reader) readLuaInteger() int64 {
	return int64(self.readUint64())
}

func (self *reader) readLuaNumber() float64 {
	return math.Float64frombits(self.readUint64())
}

func (self *reader) readString() string {
	size := uint(self.readByte())
	if size == 0 {
		return ""
	}
	if size == 0xFF {
		size = uint(self.readUint64())
	}
	bytes := self.readBytes(size - 1)
	return string(bytes)
}

func (self *reader) readBytes(n uint) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}

func (self *reader) checkHeader() {
	if string(self.readBytes(4)) != LUA_SIGNATURE {
		panic("not a precompiled chunk!")
	}
	if self.readByte() != LUAC_VERSION {
		panic("version mismatch!")
	}
	if self.readByte() != LUAC_FORMAT {
		panic("format mismatch!")
	}
	if string(self.readBytes(6)) != LUAC_DATA {
		panic("corrupted!")
	}
	if self.readByte() != CINT_SIZE {
		panic("int size mismatch!")
	}
	if self.readByte() != CSIZET_SIZE {
		panic("size_t size mismatch!")
	}
	if self.readByte() != INSTRUCTION_SIZE {
		panic("instruction size mismatch!")
	}
	if self.readByte() != LUA_INTEGER_SIZE {
		panic("lua_Integer size mismatch!")
	}
	if self.readByte() != LUA_NUMBER_SIZE {
		panic("lua_Number size mismatch!")
	}
	if self.readLuaInteger() != LUAC_INT {
		panic("endianness mismatch!")
	}
	if self.readLuaNumber() != LUAC_NUM {
		panic("float format mismatch!")
	}
}
