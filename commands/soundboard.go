package command

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

//SoundBoard returns connects to voice channel and plays a sound
func SoundBoard(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	for i := range content {
		if i != 0 {
			fmt.Println(content[i])
			folder := "./commands/sounds"
			fileName := content[i] + ".mp3"
			dgv, err := joinUserVoiceChannel(s, m.Author.ID)
			if err != nil {
				fmt.Println(err)
				return
			}
			// Start loop and attempt to play all files in the given folder
			fmt.Println("Reading Folder: ", folder)
			files, _ := ioutil.ReadDir(folder)
			for _, f := range files {
				if strings.Contains(f.Name(), fileName) {
					fmt.Println("PlayAudioFile:", f.Name())
					s.UpdateStatus(0, f.Name())
					dgvoice.PlayAudioFile(dgv, fmt.Sprintf("%s/%s", folder, f.Name()), make(chan bool))
					// ToDo: Disconnect may be replaced by s.ChannelVoiceLeave()
					dgv.Disconnect()
				}
			}
			// Close connections
			dgv.Close()
		}
	}
}

// below functions taken from https://github.com/bwmarrin/discordgo/wiki/FAQ#playing-audio-over-a-voice-connection

func findUserVoiceState(session *discordgo.Session, userid string) (*discordgo.VoiceState, error) {
	for _, guild := range session.State.Guilds {
		for _, vs := range guild.VoiceStates {
			if vs.UserID == userid {
				return vs, nil
			}
		}
	}
	return nil, errors.New("Could not find user's voice state")
}

func joinUserVoiceChannel(session *discordgo.Session, userID string) (*discordgo.VoiceConnection, error) {
	// Find a user's current voice channel
	vs, err := findUserVoiceState(session, userID)
	if err != nil {
		return nil, err
	}

	// Join the user's channel and start unmuted and deafened.
	return session.ChannelVoiceJoin(vs.GuildID, vs.ChannelID, false, true)
}

// Reads an opus packet to send over the vc.OpusSend channel
// func readOpus(source io.Reader) ([]byte, error) {
// 	var opuslen int16
// 	err := binary.Read(source, binary.LittleEndian, &opuslen)
// 	if err != nil {
// 		if err == io.EOF || err == io.ErrUnexpectedEOF {
// 			return nil, err
// 		}
// 		return nil, errors.New("ERR reading opus header")
// 	}

// 	var opusframe = make([]byte, opuslen)
// 	err = binary.Read(source, binary.LittleEndian, &opusframe)
// 	if err != nil {
// 		if err == io.EOF || err == io.ErrUnexpectedEOF {
// 			return nil, err
// 		}
// 		return nil, errors.New("ERR reading opus frame")
// 	}

// 	return opusframe, nil
// }
