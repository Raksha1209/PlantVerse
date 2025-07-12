import pretty_midi
import pandas as pd
import sys
import os
import subprocess
from pydub import AudioSegment
import sys
import io
sys.stdout = io.TextIOWrapper(sys.stdout.buffer, encoding='utf-8')


SOUNDFONT = r"C:\tools\fluidsynth\GeneralUser-GS.sf2"  # ✅ your downloaded SoundFont path

def detect_mood(mean, std, var):
    if mean > 2.5 and std < 0.5:
        return "calm"
    elif var > 2.0:
        return "excited"
    elif std > 1.5:
        return "anxious"
    elif mean < 1.0:
        return "sad"
    else:
        return "happy"

# 🎹 Map mood → General MIDI instrument numbers
mood_instruments = {
    "happy": 0,        # Acoustic Grand Piano
    "calm": 14,        # Tubular Bells
    "excited": 30,     # Overdriven Guitar
    "anxious": 27,     # Electric Guitar (Jazz)
    "sad": 40          # Violin
}

def csv_to_midi(csv_path, midi_path):
    if not os.path.exists(csv_path):
        print("❌ CSV file not found:", csv_path)
        return False

    df = pd.read_csv(csv_path)
    required_columns = {'mean_voltage', 'var_voltage', 'std_voltage'}
    if not required_columns.issubset(df.columns):
        print("❌ Required columns not found in CSV")
        return False

    midi = pretty_midi.PrettyMIDI()
    instruments = {}  # mood → instrument
    all_mood_notes = []  # (mood, note)

    start = 0.0
    for _, row in df.iterrows():
        try:
            mean = float(row['mean_voltage'])
            var = float(row['var_voltage'])
            std = float(row['std_voltage'])

            mood = detect_mood(mean, std, var)
            program = mood_instruments.get(mood, 0)

            if mood not in instruments:
                instruments[mood] = pretty_midi.Instrument(program=program, name=mood)

            pitch = int(60 + mean * 36)
            duration = max(0.3, min(1.5, 0.5 + std))
            velocity = int(max(50, min(127, 80 + var * 40)))

            print(f"🎼 Row {_}: mood = {mood}, pitch = {pitch}, duration = {duration:.2f}, velocity = {velocity}")

            note = pretty_midi.Note(
                velocity=velocity,
                pitch=max(30, min(96, pitch)),
                start=start,
                end=start + duration
            )

            instruments[mood].notes.append(note)
            all_mood_notes.append((mood, note))
            start += duration
        except Exception as e:
            print("⚠️ Skipping row due to error:", e)
            continue

    if not all_mood_notes:
        print("❌ No valid notes generated!")
        return False

    # ⏱️ Loop until total duration reaches at least 60 seconds
    loop_time = all_mood_notes[-1][1].end
    if loop_time < 60:
        loops_needed = int(60 / loop_time)
        for i in range(1, loops_needed + 1):
            for mood, note in all_mood_notes:
                shifted = pretty_midi.Note(
                    velocity=note.velocity,
                    pitch=note.pitch,
                    start=note.start + i * loop_time,
                    end=note.end + i * loop_time
                )
                instruments[mood].notes.append(shifted)

    for instrument in instruments.values():
        midi.instruments.append(instrument)

    os.makedirs(os.path.dirname(midi_path), exist_ok=True)
    midi.write(midi_path)
    print(f"✅ MIDI generated with moods: {midi_path}")
    return True


def midi_to_mp3(midi_path, mp3_path):
    import subprocess
    from pydub import AudioSegment
    import os

    wav_path = midi_path.replace(".mid", ".wav")
    fluidsynth_path = r"C:\tools\fluidsynth\bin\fluidsynth.exe"
    soundfont_path = r"C:\Users\raksh\Desktop\Placements-june\plantverse\server\assets\FluidR3_GM.sf2"

    print(f"➡ Converting MIDI to WAV using FluidSynth")

    result = subprocess.run([
        fluidsynth_path,
        "-F", wav_path,
        "-T", "wav",
        soundfont_path,
        midi_path
    ], capture_output=True, text=True)

    print(" FluidSynth STDOUT:\n", result.stdout)
    print(" FluidSynth STDERR:\n", result.stderr)

    if result.returncode != 0:
        print("❌ FluidSynth failed with return code", result.returncode)
        return

    if not os.path.exists(wav_path):
        print("❌ WAV not generated!")
        return

    sound = AudioSegment.from_wav(wav_path)
    print(f" WAV duration: {len(sound) / 1000:.2f} seconds")

    if len(sound) < 1000:
        print("⚠️ Warning: Very short or silent WAV!")

    sound.export(mp3_path, format="mp3")
    print(f"🎧 MP3 saved: {mp3_path}")

    os.remove(wav_path)



if __name__ == "__main__":
    if len(sys.argv) < 3:
        print("Usage: python generate_music.py <input_csv> <output_midi>")
    else:
        csv_file = sys.argv[1]
        midi_file = sys.argv[2]
        mp3_file = midi_file.replace(".mid", ".mp3")

        if csv_to_midi(csv_file, midi_file):
            try:
                midi_to_mp3(midi_file, mp3_file)
            except Exception as e:
                print("❌ Music conversion failed:", e)
