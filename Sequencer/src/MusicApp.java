import javax.sound.midi.*;

public class MusicApp {
	public void play()
	{
		try
		{
			Sequencer player = MidiSystem.getSequencer();
			player.open();
			
			Sequence seq = new Sequence(Sequence.PPQ, 4);
			Track track = seq.createTrack();
			
			ShortMessage msg = new ShortMessage();
			msg.setMessage(144, 1, 60, 100);
			MidiEvent noteOn = new MidiEvent(msg, 1);
			track.add(noteOn);
			
			ShortMessage msg2 = new ShortMessage();
			msg2.setMessage(128, 1, 60, 100);
			MidiEvent noteOff = new MidiEvent(msg2, 16);
			track.add(noteOff);
			
			player.setSequence(seq);
			player.start();
		}
		catch (Exception ex)
		{
			ex.printStackTrace();
		}	
	}
}
