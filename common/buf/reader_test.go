package buf_test

import (
	"io"
	"strings"
	"testing"

	"v2ray.com/core/common"
	. "v2ray.com/core/common/buf"
	"v2ray.com/core/transport/pipe"
	. "v2ray.com/ext/assert"
)

func TestBytesReaderWriteTo(t *testing.T) {
	assert := With(t)

	pReader, pWriter := pipe.New(pipe.WithSizeLimit(1024))
	reader := &BufferedReader{Reader: pReader}
	b1 := New()
	b1.WriteBytes('a', 'b', 'c')
	b2 := New()
	b2.WriteBytes('e', 'f', 'g')
	assert(pWriter.WriteMultiBuffer(NewMultiBufferValue(b1, b2)), IsNil)
	pWriter.Close()

	pReader2, pWriter2 := pipe.New(pipe.WithSizeLimit(1024))
	writer := NewBufferedWriter(pWriter2)
	writer.SetBuffered(false)

	nBytes, err := io.Copy(writer, reader)
	assert(err, IsNil)
	assert(nBytes, Equals, int64(6))

	mb, err := pReader2.ReadMultiBuffer()
	assert(err, IsNil)
	assert(len(mb), Equals, 2)
	assert(mb[0].String(), Equals, "abc")
	assert(mb[1].String(), Equals, "efg")
}

func TestBytesReaderMultiBuffer(t *testing.T) {
	assert := With(t)

	pReader, pWriter := pipe.New(pipe.WithSizeLimit(1024))
	reader := &BufferedReader{Reader: pReader}
	b1 := New()
	b1.WriteBytes('a', 'b', 'c')
	b2 := New()
	b2.WriteBytes('e', 'f', 'g')
	assert(pWriter.WriteMultiBuffer(NewMultiBufferValue(b1, b2)), IsNil)
	pWriter.Close()

	mbReader := NewReader(reader)
	mb, err := mbReader.ReadMultiBuffer()
	assert(err, IsNil)
	assert(len(mb), Equals, 2)
	assert(mb[0].String(), Equals, "abc")
	assert(mb[1].String(), Equals, "efg")
}

func TestReadByte(t *testing.T) {
	sr := strings.NewReader("abcd")
	reader := &BufferedReader{
		Reader: NewReader(sr),
	}
	b, err := reader.ReadByte()
	common.Must(err)
	if b != 'a' {
		t.Error("unexpected byte: ", b, " want a")
	}

	var mb MultiBuffer
	nBytes, err := reader.WriteTo(&mb)
	common.Must(err)
	if nBytes != 3 {
		t.Error("unexpect bytes written: ", nBytes)
	}
}

func TestReaderInterface(t *testing.T) {
	_ = (io.Reader)(new(ReadVReader))
	_ = (Reader)(new(ReadVReader))

	_ = (Reader)(new(BufferedReader))
	_ = (io.Reader)(new(BufferedReader))
	_ = (io.ByteReader)(new(BufferedReader))
	_ = (io.WriterTo)(new(BufferedReader))
}
