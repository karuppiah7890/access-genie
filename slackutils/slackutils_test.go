package slackutils_test

import (
	"testing"

	"github.com/nlopes/slack"

	"github.com/karuppiah7890/access-genie/slackutils"
	"github.com/stretchr/testify/assert"
)

func TestSelectBlockForList(t *testing.T) {
	list := []string{
		"item1",
		"item2",
	}

	placeholder := slack.NewTextBlockObject("plain_text", "choose one", false, false)
	options := []*slack.OptionBlockObject{
		slack.NewOptionBlockObject("item1", slack.NewTextBlockObject("plain_text", "item1", false, false)),
		slack.NewOptionBlockObject("item2", slack.NewTextBlockObject("plain_text", "item2", false, false)),
	}

	expectedSlackMessage := &slack.SelectBlockElement{
		Type:        "static_select",
		Placeholder: placeholder,
		ActionID:    "",
		Options:     options,
	}

	slackMessage := slackutils.SelectBlockForList(list)

	assert.Equal(t, expectedSlackMessage, slackMessage)
}

func TestSectionBlock(t *testing.T) {
	buttonBlock := slack.NewButtonBlockElement("", "", slack.NewTextBlockObject("plain_text", "button", false, false))
	txtBlock := slack.NewTextBlockObject("plain_text", "section text", false, false)
	expectedSection := slack.NewSectionBlock(txtBlock, nil, slack.NewAccessory(buttonBlock))

	section := slackutils.SectionBlock("section text", buttonBlock)

	assert.Equal(t, expectedSection, section)
}
