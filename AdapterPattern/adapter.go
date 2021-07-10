package adapterpattern

import "fmt"

type mediaPlayer interface {
	play(string)
}

type mp3 struct{}

func (m *mp3) play(obj string) {
	fmt.Println("mp3: " + obj)
}

type mp4 struct{}

func (m *mp4) play(obj string) {
	fmt.Println("mp4: " + obj)
}

func NewAdapter(media mediaPlayer) Adapter {
	return Adapter{
		mediaPlayer: media,
	}
}

type Adapter struct {
	mediaPlayer
}
