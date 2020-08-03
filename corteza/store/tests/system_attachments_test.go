package test_store

import (
	"context"
	"github.com/cortezaproject/corteza-server/corteza/store"
	"github.com/cortezaproject/corteza-server/system/types"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func testAttachment(t *testing.T, s store.Attachments) {
	var (
		ctx = context.Background()
		req = require.New(t)

		//err        error
		attachment *types.Attachment
	)

	t.Run("create", func(t *testing.T) {
		attachment = &types.Attachment{
			ID:        42,
			CreatedAt: time.Now(),
		}
		req.NoError(s.CreateAttachment(ctx, attachment))
	})

	t.Run("lookup by ID", func(t *testing.T) {
		fetched, err := s.LookupAttachmentByID(ctx, attachment.ID)
		req.NoError(err)
		req.Equal(attachment.ID, fetched.ID)
		req.NotNil(fetched.CreatedAt)
		req.Nil(fetched.UpdatedAt)
		req.Nil(fetched.DeletedAt)
	})

	t.Run("update", func(t *testing.T) {
		attachment = &types.Attachment{
			ID:        42,
			CreatedAt: time.Now(),
		}
		req.NoError(s.UpdateAttachment(ctx, attachment))
	})

	//t.Run("delete/undelete", func(t *testing.T) {
	//	ID := attachment.ID
	//	attachment, err = s.LookupAttachmentByID(ctx, ID)
	//	req.NoError(err)
	//
	//	req.NoError(s.DeleteAttachmentByID(ctx, ID))
	//	attachment, err = s.LookupAttachmentByID(ctx, ID)
	//	req.NoError(err)
	//	req.NotNil(attachment.DeletedAt)
	//
	//	req.NoError(s.UndeleteAttachmentByID(ctx, ID))
	//	attachment, err = s.LookupAttachmentByID(ctx, ID)
	//	req.NoError(err)
	//	req.Nil(attachment.DeletedAt)
	//})

	t.Run("search", func(t *testing.T) {
		t.Skip("not implemented")
		//set, f, err := s.SearchAttachments(ctx, types.AttachmentFilter{})
		//req.NoError(err)
		//req.Len(set, 1)
		//req.Equal(uint(1), f.Count)
	})

	t.Run("search by *", func(t *testing.T) {
		t.Skip("not implemented")
	})

	t.Run("ordered search", func(t *testing.T) {
		t.Skip("not implemented")
	})
}
