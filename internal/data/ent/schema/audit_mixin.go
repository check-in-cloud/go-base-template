package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"fmt"
	"github.com/go-kratos/kratos-layout/pkgs/helpers"
	"time"
)

type AuditMixin struct {
	mixin.Schema
}

func (AuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.String("created_by").
			Optional(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("updated_by").
			Optional(),
	}
}

func (AuditMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("created_by"),
		index.Fields("updated_at"),
		index.Fields("updated_by"),
	}
}

func (AuditMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		AuditHook,
	}
}

// A AuditHook is an example for audit-log hook.
func AuditHook(next ent.Mutator) ent.Mutator {
	// AuditLogger wraps the methods that are shared between all mutations of
	// schemas that embed the AuditLog mixin. The variable "exists" is true, if
	// the field already exists in the mutation (e.g. was set by a different hook).
	type AuditLogger interface {
		SetCreatedAt(time.Time)
		CreatedAt() (value time.Time, exists bool)
		SetCreatedBy(string)
		CreatedBy() (id string, exists bool)
		SetUpdatedAt(time.Time)
		UpdatedAt() (value time.Time, exists bool)
		SetUpdatedBy(string)
		UpdatedBy() (id string, exists bool)
	}
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		ml, ok := m.(AuditLogger)
		if !ok {
			return nil, fmt.Errorf("unexpected audit-log call from mutation type %T", m)
		}
		uid, err := helpers.UserIDFromContext(ctx)
		if err != nil {
			return nil, err
		}
		switch op := m.Op(); {
		case op.Is(ent.OpCreate):
			ml.SetCreatedAt(time.Now())
			if _, exists := ml.CreatedBy(); !exists {
				ml.SetCreatedBy(uid)
			}
		case op.Is(ent.OpUpdateOne | ent.OpUpdate):
			ml.SetUpdatedAt(time.Now())
			if _, exists := ml.UpdatedBy(); !exists {
				ml.SetUpdatedBy(uid)
			}
		}
		return next.Mutate(ctx, m)
	})
}
