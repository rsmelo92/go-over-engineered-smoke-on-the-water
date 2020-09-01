// http://www.cs.uccs.edu/~cs525/midi/midi.html
// http://www.music-software-development.com/midi-tutorial.html
// https://www.pioneerdj.com/-/media/pioneerdj/software-info/controller/ddj-wego3/ddj-wego3_list_of_midi_message_e.pdf
package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"net/http"

	"github.com/algoGuy/EasyMIDI/smf"
	"github.com/algoGuy/EasyMIDI/smfio"
)

const midiFile = "static/smoke_on_the_water.mid"

func initMidi() *smf.MIDIFile {
	// Create division
	division, err := smf.NewDivision(960, smf.NOSMTPE)
	checkErr(err)

	// Create new midi struct
	midi, err := smf.NewSMF(smf.Format0, *division)
	checkErr(err)
	return midi
}

func writeMidi(midi *smf.MIDIFile)  {
	// Save to new midi source file
	outputMidi, err := os.Create(midiFile)
	checkErr(err)
	defer outputMidi.Close()

	// Create buffering stream
	writer := bufio.NewWriter(outputMidi)
	smfio.Write(writer, midi)
	writer.Flush()
}

func addNote(note uint8, delay uint32, track *smf.Track) {
	const volume = 10
	const velocity = 64
	const status = smf.NoteOnStatus

	notePartOne, err := smf.NewMIDIEvent(delay, status, volume, note, velocity)
	checkErr(err)

	err = track.AddEvent(notePartOne)
	checkErr(err)
}

func endOfTrack(delay uint32, track *smf.Track) {
	endOfTrack, err := smf.NewMetaEvent(delay, smf.MetaEndOfTrack, []byte{})
	checkErr(err)
	err = track.AddEvent(endOfTrack)
	checkErr(err)
}

func trackGenerator()  {
	// Create new midi struct
	midi := initMidi()
	
	// Create track struct
	track := &smf.Track{}
	
	// Add track to new midi struct
	midi.AddTrack(track)
	
	const delay = 500
	// Create some midi events
	addNote(79, delay*4, track)
	addNote(82, delay*2, track)
	addNote(84, delay*2, track)

	addNote(79, delay*3, track)
	addNote(82, delay*2, track)
	addNote(85, delay*2, track)
	addNote(84, delay, track)

	addNote(79, delay*4, track)
	addNote(82, delay*2, track)
	addNote(84, delay*2, track)

	addNote(82, delay*3, track)
	addNote(79, delay*2, track)
	endOfTrack(delay*5, track)

	writeMidi(midi)
}


func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	
	trackGenerator()
	http.HandleFunc("/smoke_on_the_water.mid", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/smoke_on_the_water.mid")
	})

	fmt.Println("Server has started on http://localhost:8080")
	panic(http.ListenAndServe(":8080", nil))

}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
