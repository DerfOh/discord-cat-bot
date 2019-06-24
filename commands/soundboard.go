package command

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var buffer = make([][]byte, 0)

// SoundBoard makes bot join voice channel and play audio file
//		dca files are created using combindation of dca golang and ffmpeg 'ffmpeg -i test.mp3 -f s16le -ar 48000 -ac 2 pipe:1 | dca > test.dca'
func SoundBoard(content []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	for i := range content {
		if i != 0 {
			// fmt.Println(content[i])
			folder := "./commands/sounds/"
			fileName := content[i] + ".dca"
			// Start loop and attempt to play all files in the given folder
			fmt.Println("Reading Folder: ", folder)
			files, _ := ioutil.ReadDir(folder)
			for _, f := range files {
				if strings.Contains(f.Name(), fileName) {
					fmt.Println(fileName + " and " + f.Name() + " match ")
					// Load the sound file.
					err := loadSound(fileName)
					if err != nil {
						fmt.Println("Error loading sound: ", err)
						return
					}
				}
			}
			// Find the channel that the message came from.
			c, err := s.State.Channel(m.ChannelID)
			if err != nil {
				// Could not find channel.
				return
			}

			// Find the guild for that channel.
			g, err := s.State.Guild(c.GuildID)
			if err != nil {
				// Could not find guild.
				return
			}

			// Look for the message sender in that guild's current voice states.
			for _, vs := range g.VoiceStates {
				if vs.UserID == m.Author.ID {
					err = playSound(s, g.ID, vs.ChannelID)
					if err != nil {
						fmt.Println("Error playing sound:", err)
					}

					return
				}
			}
		}
	}
}

// loadSound attempts to load an encoded sound file from disk.
func loadSound(fileName string) error {

	file, err := os.Open("./commands/sounds/" + fileName)
	if err != nil {
		fmt.Println("Error opening dca file :", err)
		return err
	}

	var opuslen int16

	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				return err
			}
			return nil
		}

		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		// Read encoded pcm from dca file.
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}

		// Append encoded pcm data to the buffer.
		buffer = append(buffer, InBuf)
	}
}

// playSound plays the current buffer to the provided channel.
func playSound(s *discordgo.Session, guildID, channelID string) (err error) {

	// Join the provided voice channel.
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return err
	}

	// Sleep for a specified amount of time before playing the sound
	time.Sleep(250 * time.Millisecond)

	// Start speaking.
	vc.Speaking(true)

	// Send the buffer data.
	for _, buff := range buffer {
		vc.OpusSend <- buff
	}

	// Stop speaking
	vc.Speaking(false)

	// Sleep for a specificed amount of time before ending.
	time.Sleep(250 * time.Millisecond)

	// Disconnect from the provided voice channel.
	vc.Disconnect()

	// Clear the buffer
	buffer = nil
	return nil
}

// leaveChannel explicitly makes the bot leave
func leaveChannel(s *discordgo.Session, guildID, channelID string) (err error) {

	// Join the provided voice channel.
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return err
	}

	// Stop speaking
	vc.Speaking(false)

	// Sleep for a specificed amount of time before ending.
	time.Sleep(250 * time.Millisecond)

	// Disconnect from the provided voice channel.
	vc.Disconnect()

	// Clear the buffer
	buffer = nil
	return nil
}
