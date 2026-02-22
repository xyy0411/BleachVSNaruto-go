package audio

const (
	EventJump = "jump"
)

func DefaultSFX() map[string]string {
	return map[string]string{
		EventJump: "assets/public_audio/jump/0.wav",
	}
}
