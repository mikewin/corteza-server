package rdbms

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// corteza/store/reminders.yaml

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/corteza/store"
	"github.com/cortezaproject/corteza-server/system/types"
	"github.com/jmoiron/sqlx"
)

// SearchReminders returns all matching rows
//
// This function calls convertReminderFilter with the given
// types.ReminderFilter and expects to receive a working squirrel.SelectBuilder
func (s Store) SearchReminders(ctx context.Context, f types.ReminderFilter) (types.ReminderSet, types.ReminderFilter, error) {
	query, err := s.convertReminderFilter(f)
	if err != nil {
		return nil, f, err
	}

	var set = types.ReminderSet{}

	if f.Count, err = Count(ctx, s.db, query); err != nil || f.Count == 0 {
		return nil, f, err

	}

	return set, f, FetchPaged(ctx, s.db, query, f.PageFilter, &set)

}

// LookupReminderByID searches for reminder by its ID
//
// It returns reminder even if deleted or suspended
func (s Store) LookupReminderByID(ctx context.Context, id uint64) (*types.Reminder, error) {
	return s.ReminderLookup(ctx, squirrel.Eq{
		"rmd.id": id,
	})
}

// CreateReminder creates one or more rows in sys_reminder table
func (s Store) CreateReminder(ctx context.Context, rr ...*types.Reminder) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Insert(s.ReminderTable()).SetMap(s.ReminderEnc(res)))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// UpdateReminder updates one or more existing rows in sys_reminder
func (s Store) UpdateReminder(ctx context.Context, rr ...*types.Reminder) error {
	return s.PartialUpdateReminder(ctx, nil, rr...)
}

// PartialUpdateReminder updates one or more existing rows in sys_reminder
//
// It wraps the update into transaction and can perform partial update by providing list of updatable columns
func (s Store) PartialUpdateReminder(ctx context.Context, onlyColumns []string, rr ...*types.Reminder) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = s.ExecUpdateReminders(
				ctx,
				squirrel.Eq{s.preprocessColumn("rmd.id", ""): s.preprocessValue(res.ID, "")},
				s.ReminderEnc(res).Skip("id").Only(onlyColumns...))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveReminder removes one or more rows from sys_reminder table
func (s Store) RemoveReminder(ctx context.Context, rr ...*types.Reminder) error {
	if len(rr) == 0 {
		return nil
	}

	return Tx(ctx, s.db, s.config, nil, func(db *sqlx.Tx) (err error) {
		for _, res := range rr {
			err = ExecuteSqlizer(ctx, s.DB(), s.Delete(s.ReminderTable("rmd")).Where(squirrel.Eq{s.preprocessColumn("rmd.id", ""): s.preprocessValue(res.ID, "")}))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RemoveReminderByID removes row from the sys_reminder table
func (s Store) RemoveReminderByID(ctx context.Context, ID uint64) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Delete(s.ReminderTable()).Where(squirrel.Eq{s.preprocessColumn("rmd.id", ""): s.preprocessValue(ID, "")}))
}

// TruncateReminders removes all rows from the sys_reminder table
func (s Store) TruncateReminders(ctx context.Context) error {
	return Truncate(ctx, s.DB(), s.ReminderTable())
}

// ExecUpdateReminders updates all matchhed (cnd) rows in sys_reminder with given data
func (s Store) ExecUpdateReminders(ctx context.Context, cnd squirrel.Sqlizer, set store.Payload) error {
	return ExecuteSqlizer(ctx, s.DB(), s.Update(s.ReminderTable("rmd")).Where(cnd).SetMap(set))
}

// ReminderLookup calls Lookup() and returns (if found) types.Reminder
func (s Store) ReminderLookup(ctx context.Context, cnd squirrel.Sqlizer) (*types.Reminder, error) {
	var (
		u   = &types.Reminder{}
		err = s.Lookup(ctx, u, s.QueryReminders(), cnd)
	)

	if err == nil {
		return u, nil
	}

	return nil, err
	//return s.scanReminder(s.LookupWithScan(ctx, s.QueryReminders(), cnd))
}

// QueryReminders returns squirrel.SelectBuilder with set table and all columns
func (s Store) QueryReminders() squirrel.SelectBuilder {
	return s.Select(s.ReminderTable("rmd"), s.ReminderColumns("rmd")...)
}

// ReminderTable name of the db table
func (Store) ReminderTable(aa ...string) string {
	var alias string
	if len(aa) > 0 {
		alias = " AS " + aa[0]
	}

	return "sys_reminder" + alias
}

// ReminderColumns returns all defined table columns
//
// With optional string arg, all columns are returned aliased
func (Store) ReminderColumns(aa ...string) []string {
	var alias string
	if len(aa) > 0 {
		alias = aa[0] + "."
	}

	return []string{
		alias + "id",
		alias + "resource",
		alias + "payload",
		alias + "snooze_count",
		alias + "assigned_to",
		alias + "assigned_by",
		alias + "assigned_at",
		alias + "dismissed_by",
		alias + "dismissed_at",
		alias + "remind_at",
		alias + "created_at",
		alias + "updated_at",
		alias + "deleted_at",
	}
}

// ReminderEnc encodes fields from types.Reminder to store.Payload (map)
func (Store) ReminderEnc(res *types.Reminder) store.Payload {
	return store.Payload{
		"id":           res.ID,
		"resource":     res.Resource,
		"payload":      res.Payload,
		"snooze_count": res.SnoozeCount,
		"assigned_to":  res.AssignedTo,
		"assigned_by":  res.AssignedBy,
		"assigned_at":  res.AssignedAt,
		"dismissed_by": res.DismissedBy,
		"dismissed_at": res.DismissedAt,
		"remind_at":    res.RemindAt,
		"created_at":   res.CreatedAt,
		"updated_at":   res.UpdatedAt,
		"deleted_at":   res.DeletedAt,
	}
}
