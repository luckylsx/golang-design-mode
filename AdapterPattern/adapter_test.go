package adapterpattern

import "testing"

func TestAdapter(t *testing.T) {
	adapterMp3 := NewAdapter(new(mp3))
	adapterMp3.play("music")
	adapterMp4 := NewAdapter(new(mp4))
	adapterMp4.play("video")
}
