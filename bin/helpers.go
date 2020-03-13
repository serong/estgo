package bin

import (
	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
	"log"
)

func AddUnderline(str string, c string) (result string) {
	result = str + "\n"

	for i := 0; i < len(str); i++ {
		result = result + c
	}

	return result + "\n"
}

func PlayMP3(rawMP3 []byte) {
	decoder, data, err := minimp3.DecodeFull(rawMP3)
	if err != nil {
		log.Fatal(err)
	}

	player, err := oto.NewPlayer(decoder.SampleRate, decoder.Channels, 2, 1024)
	if err != nil {
		log.Fatal(err)
	}

	player.Write(data)
}
