package main

import (
	"fmt"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func (p *guard) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {

	if post.IsSystemMessage() {
		return post, ""
	}

	if p.channelRoleChecker(post) {
		return post, ""
	}

	postUser, _ := p.API.GetUser(post.UserId)
	if postUser.IsBot == true {
		return post, ""
	}

	guards := p.getGuards()
	allowedUsers, ok := guards[post.ChannelId]
	if ok == false {
		return post, ""
	}

	channel, _ := p.API.GetChannel(post.ChannelId)
	if p.isTeamAdmin(post.UserId, channel.TeamId) {
		return post, ""
	}

	users, _ := p.API.GetUsersByUsernames(allowedUsers)
	if len(users) != 0 {
		for _, user := range users {
			if post.UserId == user.Id {
				return post, ""
			}
		}
	}
	p.API.SendEphemeralPost(post.UserId, &model.Post{
		UserId:    p.botUserID,
		ChannelId: post.ChannelId,
		Message:   p.message,
	})

	str := fmt.Sprintf("%s attempted to post in channel %s", post.UserId, post.ChannelId)
	return nil, str

}

func (p *guard) isTeamAdmin(userId string, teamId string) bool {

	teamMember, _ := p.API.GetTeamMember(teamId, userId)

	teamRoles := teamMember.GetRoles()

	for _, a := range teamRoles {
		if a == "team_admin" {
			return true
		}
	}

	return false

}

func (p *guard) channelRoleChecker(post *model.Post) bool {

	channelMember, _ := p.API.GetChannelMember(post.ChannelId, post.UserId)

	channelRoles := channelMember.GetRoles()

	for _, a := range channelRoles {
		if a == "channel_admin" {
			return true
		}
	}

	return false

}
