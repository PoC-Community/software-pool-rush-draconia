// Code generated by ent, DO NOT EDIT.

package character

import (
	"draconia/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Character {
	return predicate.Character(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Character {
	return predicate.Character(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Character {
	return predicate.Character(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Character {
	return predicate.Character(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Character {
	return predicate.Character(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Character {
	return predicate.Character(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Character {
	return predicate.Character(sql.FieldLTE(FieldID, id))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v int64) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldPrice, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v int64) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v int64) predicate.Character {
	return predicate.Character(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...int64) predicate.Character {
	return predicate.Character(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...int64) predicate.Character {
	return predicate.Character(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v int64) predicate.Character {
	return predicate.Character(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v int64) predicate.Character {
	return predicate.Character(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v int64) predicate.Character {
	return predicate.Character(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v int64) predicate.Character {
	return predicate.Character(sql.FieldLTE(FieldPrice, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Character) predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Character) predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Character) predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		p(s.Not())
	})
}
