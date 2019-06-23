package notification

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/derfoh/discord-cat-bot/config"
)

var buffer = make([][]byte, 0)

// Notify makes bot join voice channel and play audio file in the owner's channel notifying someone joined the channel
//		dca files are created using combindation of dca golang and ffmpeg 'ffmpeg -i test.mp3 -f s16le -ar 48000 -ac 2 pipe:1 | dca > test.dca'
func Notify(action string, s *discordgo.Session, channelID string) {
	folder := "./notifications/sounds/"
	fileName := action + ".dca"
	// Start loop and attempt to play all files in the given folder
	// fmt.Println("File Name: ", fileName)
	// fmt.Println("Reading Folder: ", folder)
	files, _ := ioutil.ReadDir(folder)
	for _, f := range files {
		// fmt.Println("Found: ", f.Name())
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
	c, err := s.State.Channel(channelID)
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

	// Look for the notification target in that guild's current voice states.
	for _, vs := range g.VoiceStates {
		if vs.UserID == config.BotOwner {
			fmt.Println("notifying owner")
			err = playSound(s, g.ID, vs.ChannelID)
			if err != nil {
				fmt.Println("Error playing sound:", err)
			}

			return
		}
	}

}

// loadSound attempts to load an encoded sound file from disk.
func loadSound(fileName string) error {

	file, err := os.Open("./notifications/sounds/" + fileName)
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
