package schema

//
//import (
//	"entgo.io/ent"
//	"entgo.io/ent/schema/field"
//	"entgo.io/ent/schema/index"
//	"github.com/google/uuid"
//)
//
//// ModelSample holds the schema definition for the Setting entity.
//type ModelSample struct {
//	ent.Schema
//}
//
//func (ModelSample) Mixin() []ent.Mixin {
//	return []ent.Mixin{
//		AuditMixin{},
//	}
//}
//
//// Fields of the Setting.
//func (ModelSample) Fields() []ent.Field {
//	return []ent.Field{
//		field.UUID("id", uuid.UUID{}).
//			Default(uuid.New).
//			Unique().
//			Immutable(),
//		field.String("name").Comment("Name").Optional(),
//	}
//}
//
//func (ModelSample) Indexes() []ent.Index {
//	return []ent.Index{
//		// unique index.
//		index.Fields("name").
//			Unique(),
//	}
//}
