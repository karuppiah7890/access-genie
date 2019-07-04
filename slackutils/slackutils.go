package slackutils

import (
	"github.com/nlopes/slack"
)

// SelectBlockForList gives the slack select block for a given list of strings
func SelectBlockForList(list []string) *slack.SelectBlockElement {
	placeholder := slack.NewTextBlockObject("plain_text", "choose one", false, false)
	allOptBlocks := make([]*slack.OptionBlockObject, 0, len(list))

	for _, elem := range list {
		txtBlock := slack.NewTextBlockObject("plain_text", elem, false, false)
		optBlock := slack.NewOptionBlockObject(elem, txtBlock)
		allOptBlocks = append(allOptBlocks, optBlock)
	}

	return slack.NewOptionsSelectBlockElement("static_select", placeholder, "", allOptBlocks...)
}

// SectionBlock gives the slack section block for a given text and slack block
func SectionBlock(sectionTxt string, slackBlock slack.BlockElement) *slack.SectionBlock {
	return slack.NewSectionBlock(slack.NewTextBlockObject("plain_text", sectionTxt, false, false), nil, slack.NewAccessory(slackBlock))
}
