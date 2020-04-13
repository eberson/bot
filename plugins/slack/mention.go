package slack

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	reMention *regexp.Regexp
)

func (s *Slack) ShouldAnswer(ev *slackEvent) bool {
	if s.isDirectToBot(ev.source.Channel) {
		return true
	}

	if reMention == nil {
		reMention = regexp.MustCompile(fmt.Sprintf("^(?P<user>\\<@%s\\>)[\\s,]?(?P<message>.*)$", s.rootinha.ID))
	}

	match := reMention.FindStringSubmatch(ev.Text())

	if len(match) >= 3 {
		ev.text = strings.Trim(match[len(match)-1], " ")
		return true
	}

	return false
}

func (s *Slack) isDirectToBot(channel string) bool {
	if len(s.dms) == 0 {
		var users Users = s.users
		s.dms = users.createDMChannel(s.api)
	}

	for _, dm := range s.dms {
		if strings.EqualFold(dm.channel, channel) {
			return true
		}
	}

	return false
}
