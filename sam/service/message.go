package service

import (
	"context"
	"github.com/crusttech/crust/internal/auth"
	"github.com/crusttech/crust/sam/repository"
	"github.com/crusttech/crust/sam/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

type (
	message struct {
		channel repository.Channel
		message repository.Message
		reaction repository.Reaction

		att AttachmentService
	}

	MessageService interface {
		Find(ctx context.Context, filter *types.MessageFilter) (types.MessageSet, error)

		Create(ctx context.Context, messages *types.Message) (*types.Message, error)
		Update(ctx context.Context, messages *types.Message) (*types.Message, error)

		React(ctx context.Context, messageID uint64, reaction string) error
		Unreact(ctx context.Context, messageID uint64, reaction string) error

		Pin(ctx context.Context, messageID uint64) error
		Unpin(ctx context.Context, messageID uint64) error

		Flag(ctx context.Context, messageID uint64) error
		Unflag(ctx context.Context, messageID uint64) error

		Direct(ctx context.Context, recipientID uint64, in *types.Message) (out *types.Message, err error)

		deleter
	}
)

func Message(attSvc AttachmentService) *message {
	m := &message{
		att: attSvc,
		message: repository.NewMessage(context.Background()),
	}
	return m
}

func (svc message) Find(ctx context.Context, filter *types.MessageFilter) (mm types.MessageSet, err error) {
	// @todo get user from context
	var currentUserID uint64 = auth.GetIdentityFromContext(ctx).Identity()

	// @todo verify if current user can access & read from this channel
	_ = currentUserID
	_ = filter.ChannelID

	mm, err = svc.message.FindMessages(filter)
	if err != nil {
		return nil, err
	}

	return mm, svc.att.LoadFromMessages(ctx, mm)
}

func (svc message) Direct(ctx context.Context, recipientID uint64, in *types.Message) (out *types.Message, err error) {
	return out, repository.DB().Transaction(func() (err error) {
		var currentUserID = auth.GetIdentityFromContext(ctx).Identity()

		// @todo [SECURITY] verify if current user can send direct messages to anyone?
		if false {
			return errors.New("Not allowed to send direct messages")
		}

		// @todo [SECURITY] verify if current user can send direct messages to this user
		if false {
			return errors.New("Not allowed to send direct messages to this user")
		}

		dch, err := svc.channel.FindDirectChannelByUserID(currentUserID, recipientID)
		if err == repository.ErrChannelNotFound {
			dch, err = svc.channel.CreateChannel(&types.Channel{
				Type: types.ChannelTypeDirect,
			})

			if err != nil {
				return
			}

			membership := &types.ChannelMember{ChannelID: dch.ID, Type: types.ChannelMembershipTypeOwner}

			membership.UserID = currentUserID
			spew.Dump(membership)
			if _, err = svc.channel.AddChannelMember(membership); err != nil {
				return
			}

			spew.Dump(membership)
			membership.UserID = recipientID
			if _, err = svc.channel.AddChannelMember(membership); err != nil {
				return
			}

		} else if err != nil {
			return errors.Wrap(err, "Could not send direct message")
		}

		// Make sure our message is sent to the right channel
		in.ChannelID = dch.ID
		in.UserID = currentUserID
		in.Type = types.MessageTypeSimpleMessage

		spew.Dump(in)

		// @todo send new msg to the event-loop
		out, err = svc.message.CreateMessage(in)
		return
	})
}

func (svc message) Create(ctx context.Context, mod *types.Message) (*types.Message, error) {
	// @todo get user from context
	var currentUserID uint64 = auth.GetIdentityFromContext(ctx).Identity()

	// @todo verify if current user can access & write to this channel

	mod.UserID = currentUserID

	message, err := svc.message.CreateMessage(mod)
	if err == nil {
		PubSub().Event(ctx, "new message added")
	}
	return message, err
}

func (svc message) Update(ctx context.Context, mod *types.Message) (*types.Message, error) {
	// @todo get user from context
	var currentUserID uint64 = auth.GetIdentityFromContext(ctx).Identity()

	// @todo verify if current user can access & write to this channel
	_ = currentUserID

	// @todo load current message

	// @todo verify ownership

	return svc.message.UpdateMessage(mod)
}

func (svc message) Delete(ctx context.Context, id uint64) error {
	// @todo get user from context
	var currentUserID uint64 = auth.GetIdentityFromContext(ctx).Identity()

	// @todo verify if current user can access & write to this channel
	_ = currentUserID

	// @todo load current message

	// @todo verify ownership

	return svc.message.DeleteMessageByID(id)
}

func (svc message) React(ctx context.Context, messageID uint64, reaction string) error {
	// @todo get user from context
	var currentUserID uint64 = auth.GetIdentityFromContext(ctx).Identity()

	// @todo verify if current user can access & write to this channel
	var m *types.Message

	// @todo validate reaction

	r := &types.Reaction{
		UserID:    currentUserID,
		MessageID: messageID,
		ChannelID: m.ChannelID,
		Reaction:  reaction,
	}

	if _, err := svc.reaction.CreateReaction(r); err != nil {
		return err
	}

	return nil
}

func (svc message) Unreact(ctx context.Context, messageID uint64, reaction string) error {
	// @todo get user from context
	var currentUserID uint64 = auth.GetIdentityFromContext(ctx).Identity()

	// @todo verify if current user can access & write to this channel
	_ = currentUserID

	// @todo load reaction and verify ownership
	var r *types.Reaction

	return svc.reaction.DeleteReactionByID(r.ID)
}

func (svc message) Pin(ctx context.Context, messageID uint64) error {
	// @todo get user from context
	var currentUserID uint64 = auth.GetIdentityFromContext(ctx).Identity()

	// @todo verify if current user can access & write to this channel
	_ = currentUserID

	return nil
}

func (svc message) Unpin(ctx context.Context, messageID uint64) error {
	// @todo get user from context
	var currentUserID uint64 = auth.GetIdentityFromContext(ctx).Identity()

	// @todo verify if current user can access & write to this channel
	_ = currentUserID

	return nil
}

func (svc message) Flag(ctx context.Context, messageID uint64) error {
	// @todo get user from context
	var currentUserID uint64 = auth.GetIdentityFromContext(ctx).Identity()

	// @todo verify if current user can access & write to this channel
	_ = currentUserID

	return nil
}

func (svc message) Unflag(ctx context.Context, messageID uint64) error {
	// @todo get user from context
	var currentUserID uint64 = auth.GetIdentityFromContext(ctx).Identity()

	// @todo verify if current user can access & write to this channel
	_ = currentUserID

	return nil
}

var _ MessageService = &message{}
