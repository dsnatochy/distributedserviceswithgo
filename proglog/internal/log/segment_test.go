package log

import (
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"os"
	"testing"

	api "github.com/dsnatochy/proglog/api/v1"
	"github.com/stretchr/testify/require"
)

func TestSegment(t *testing.T) {
	dir, _ := ioutil.TempDir("","segment-test")
	defer os.RemoveAll(dir)

	want := &api.Record{
		Value:  []byte("hello world"),
		Offset: 0,
	}
	c := Config{}
	c.Segment.MaxStoreBytes = 1024
	c.Segment.MaxIndexBytes = entWidth * 3

	s, err := newSegment(dir, 16,c)
	require.NoError(t,err)
	require.Equal(t, uint64(16),s.nextOffset,s.nextOffset) //TODO why twice?
	require.False(t, s.IsMaxed())

	for i := uint64(0); i<3; i++{
		off, err := s.Append(want)
		require.NoError(t,err)
		require.Equal(t, 16+i, off)

		got, err := s.Read(off)
		require.NoError(t,err)
		assert.Equal(t, want.Value, got.Value)

	}

	_, err = s.Append(want)
	assert.Equal(t, io.EOF, err)

	// maxed index
	assert.True(t, s.IsMaxed())

	c.Segment.MaxStoreBytes = uint64(len(want.Value)*3)
	c.Segment.MaxIndexBytes = 1024

	s, err = newSegment(dir, 16, c)
	require.NoError(t, err)
	// maxed store
	assert.True(t, s.IsMaxed())

	//fmt.Println(dir)
	err = s.Remove()
	require.NoError(t, err)
	s, err = newSegment(dir, 16, c)
	require.NoError(t,err)
	assert.False(t, s.IsMaxed())
}