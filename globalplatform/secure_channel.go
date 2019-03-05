package globalplatform

import (
	"github.com/status-im/keycard-go/apdu"
	"github.com/status-im/keycard-go/hexutils"
)

// SecureChannel wraps another channel and sends wrapped commands using SCP02Wrapper.
type SecureChannel struct {
	session *Session
	c       Channel
	w       *SCP02Wrapper
}

// NewSecureChannel returns a new SecureChannel based on a session and wrapping a Channel c.
func NewSecureChannel(session *Session, c Channel) *SecureChannel {
	return &SecureChannel{
		session: session,
		c:       c,
		w:       NewSCP02Wrapper(session.KeyProvider().Mac()),
	}
}

// Send sends wrapped commands to the inner channel.
func (c *SecureChannel) Send(cmd *apdu.Command) (*apdu.Response, error) {
	rawCmd, err := cmd.Serialize()
	if err != nil {
		return nil, err
	}

	logger.Debug("wrapping apdu command", "hex", hexutils.BytesToHexWithSpaces(rawCmd))
	wrappedCmd, err := c.w.Wrap(cmd)
	if err != nil {
		return nil, err
	}

	return c.c.Send(wrappedCmd)
}
