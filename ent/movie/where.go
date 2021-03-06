// Code generated by entc, DO NOT EDIT.

package movie

import (
	"imdb-movie/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Rated applies equality check predicate on the "rated" field. It's identical to RatedEQ.
func Rated(v float32) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRated), v))
	})
}

// RealeaseDate applies equality check predicate on the "realease_date" field. It's identical to RealeaseDateEQ.
func RealeaseDate(v time.Time) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRealeaseDate), v))
	})
}

// Genre applies equality check predicate on the "genre" field. It's identical to GenreEQ.
func Genre(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGenre), v))
	})
}

// Poster applies equality check predicate on the "poster" field. It's identical to PosterEQ.
func Poster(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPoster), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// Plot applies equality check predicate on the "plot" field. It's identical to PlotEQ.
func Plot(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlot), v))
	})
}

// Stars applies equality check predicate on the "stars" field. It's identical to StarsEQ.
func Stars(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStars), v))
	})
}

// ImdbRating applies equality check predicate on the "imdb_rating" field. It's identical to ImdbRatingEQ.
func ImdbRating(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImdbRating), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Movie {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Movie {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// RatedEQ applies the EQ predicate on the "rated" field.
func RatedEQ(v float32) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRated), v))
	})
}

// RatedNEQ applies the NEQ predicate on the "rated" field.
func RatedNEQ(v float32) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRated), v))
	})
}

// RatedIn applies the In predicate on the "rated" field.
func RatedIn(vs ...float32) predicate.Movie {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRated), v...))
	})
}

// RatedNotIn applies the NotIn predicate on the "rated" field.
func RatedNotIn(vs ...float32) predicate.Movie {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRated), v...))
	})
}

// RatedGT applies the GT predicate on the "rated" field.
func RatedGT(v float32) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRated), v))
	})
}

// RatedGTE applies the GTE predicate on the "rated" field.
func RatedGTE(v float32) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRated), v))
	})
}

// RatedLT applies the LT predicate on the "rated" field.
func RatedLT(v float32) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRated), v))
	})
}

// RatedLTE applies the LTE predicate on the "rated" field.
func RatedLTE(v float32) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRated), v))
	})
}

// RealeaseDateEQ applies the EQ predicate on the "realease_date" field.
func RealeaseDateEQ(v time.Time) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRealeaseDate), v))
	})
}

// RealeaseDateNEQ applies the NEQ predicate on the "realease_date" field.
func RealeaseDateNEQ(v time.Time) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRealeaseDate), v))
	})
}

// RealeaseDateIn applies the In predicate on the "realease_date" field.
func RealeaseDateIn(vs ...time.Time) predicate.Movie {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRealeaseDate), v...))
	})
}

// RealeaseDateNotIn applies the NotIn predicate on the "realease_date" field.
func RealeaseDateNotIn(vs ...time.Time) predicate.Movie {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRealeaseDate), v...))
	})
}

// RealeaseDateGT applies the GT predicate on the "realease_date" field.
func RealeaseDateGT(v time.Time) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRealeaseDate), v))
	})
}

// RealeaseDateGTE applies the GTE predicate on the "realease_date" field.
func RealeaseDateGTE(v time.Time) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRealeaseDate), v))
	})
}

// RealeaseDateLT applies the LT predicate on the "realease_date" field.
func RealeaseDateLT(v time.Time) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRealeaseDate), v))
	})
}

// RealeaseDateLTE applies the LTE predicate on the "realease_date" field.
func RealeaseDateLTE(v time.Time) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRealeaseDate), v))
	})
}

// GenreEQ applies the EQ predicate on the "genre" field.
func GenreEQ(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGenre), v))
	})
}

// GenreNEQ applies the NEQ predicate on the "genre" field.
func GenreNEQ(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGenre), v))
	})
}

// GenreIn applies the In predicate on the "genre" field.
func GenreIn(vs ...string) predicate.Movie {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGenre), v...))
	})
}

// GenreNotIn applies the NotIn predicate on the "genre" field.
func GenreNotIn(vs ...string) predicate.Movie {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGenre), v...))
	})
}

// GenreGT applies the GT predicate on the "genre" field.
func GenreGT(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGenre), v))
	})
}

// GenreGTE applies the GTE predicate on the "genre" field.
func GenreGTE(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGenre), v))
	})
}

// GenreLT applies the LT predicate on the "genre" field.
func GenreLT(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGenre), v))
	})
}

// GenreLTE applies the LTE predicate on the "genre" field.
func GenreLTE(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGenre), v))
	})
}

// GenreContains applies the Contains predicate on the "genre" field.
func GenreContains(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGenre), v))
	})
}

// GenreHasPrefix applies the HasPrefix predicate on the "genre" field.
func GenreHasPrefix(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGenre), v))
	})
}

// GenreHasSuffix applies the HasSuffix predicate on the "genre" field.
func GenreHasSuffix(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGenre), v))
	})
}

// GenreEqualFold applies the EqualFold predicate on the "genre" field.
func GenreEqualFold(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGenre), v))
	})
}

// GenreContainsFold applies the ContainsFold predicate on the "genre" field.
func GenreContainsFold(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGenre), v))
	})
}

// PosterEQ applies the EQ predicate on the "poster" field.
func PosterEQ(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPoster), v))
	})
}

// PosterNEQ applies the NEQ predicate on the "poster" field.
func PosterNEQ(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPoster), v))
	})
}

// PosterIn applies the In predicate on the "poster" field.
func PosterIn(vs ...string) predicate.Movie {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPoster), v...))
	})
}

// PosterNotIn applies the NotIn predicate on the "poster" field.
func PosterNotIn(vs ...string) predicate.Movie {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPoster), v...))
	})
}

// PosterGT applies the GT predicate on the "poster" field.
func PosterGT(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPoster), v))
	})
}

// PosterGTE applies the GTE predicate on the "poster" field.
func PosterGTE(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPoster), v))
	})
}

// PosterLT applies the LT predicate on the "poster" field.
func PosterLT(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPoster), v))
	})
}

// PosterLTE applies the LTE predicate on the "poster" field.
func PosterLTE(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPoster), v))
	})
}

// PosterContains applies the Contains predicate on the "poster" field.
func PosterContains(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPoster), v))
	})
}

// PosterHasPrefix applies the HasPrefix predicate on the "poster" field.
func PosterHasPrefix(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPoster), v))
	})
}

// PosterHasSuffix applies the HasSuffix predicate on the "poster" field.
func PosterHasSuffix(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPoster), v))
	})
}

// PosterEqualFold applies the EqualFold predicate on the "poster" field.
func PosterEqualFold(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPoster), v))
	})
}

// PosterContainsFold applies the ContainsFold predicate on the "poster" field.
func PosterContainsFold(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPoster), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Movie {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Movie {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// PlotEQ applies the EQ predicate on the "plot" field.
func PlotEQ(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlot), v))
	})
}

// PlotNEQ applies the NEQ predicate on the "plot" field.
func PlotNEQ(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPlot), v))
	})
}

// PlotIn applies the In predicate on the "plot" field.
func PlotIn(vs ...string) predicate.Movie {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPlot), v...))
	})
}

// PlotNotIn applies the NotIn predicate on the "plot" field.
func PlotNotIn(vs ...string) predicate.Movie {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPlot), v...))
	})
}

// PlotGT applies the GT predicate on the "plot" field.
func PlotGT(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPlot), v))
	})
}

// PlotGTE applies the GTE predicate on the "plot" field.
func PlotGTE(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPlot), v))
	})
}

// PlotLT applies the LT predicate on the "plot" field.
func PlotLT(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPlot), v))
	})
}

// PlotLTE applies the LTE predicate on the "plot" field.
func PlotLTE(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPlot), v))
	})
}

// PlotContains applies the Contains predicate on the "plot" field.
func PlotContains(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPlot), v))
	})
}

// PlotHasPrefix applies the HasPrefix predicate on the "plot" field.
func PlotHasPrefix(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPlot), v))
	})
}

// PlotHasSuffix applies the HasSuffix predicate on the "plot" field.
func PlotHasSuffix(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPlot), v))
	})
}

// PlotEqualFold applies the EqualFold predicate on the "plot" field.
func PlotEqualFold(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPlot), v))
	})
}

// PlotContainsFold applies the ContainsFold predicate on the "plot" field.
func PlotContainsFold(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPlot), v))
	})
}

// StarsEQ applies the EQ predicate on the "stars" field.
func StarsEQ(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStars), v))
	})
}

// StarsNEQ applies the NEQ predicate on the "stars" field.
func StarsNEQ(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStars), v))
	})
}

// StarsIn applies the In predicate on the "stars" field.
func StarsIn(vs ...string) predicate.Movie {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStars), v...))
	})
}

// StarsNotIn applies the NotIn predicate on the "stars" field.
func StarsNotIn(vs ...string) predicate.Movie {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStars), v...))
	})
}

// StarsGT applies the GT predicate on the "stars" field.
func StarsGT(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStars), v))
	})
}

// StarsGTE applies the GTE predicate on the "stars" field.
func StarsGTE(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStars), v))
	})
}

// StarsLT applies the LT predicate on the "stars" field.
func StarsLT(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStars), v))
	})
}

// StarsLTE applies the LTE predicate on the "stars" field.
func StarsLTE(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStars), v))
	})
}

// StarsContains applies the Contains predicate on the "stars" field.
func StarsContains(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStars), v))
	})
}

// StarsHasPrefix applies the HasPrefix predicate on the "stars" field.
func StarsHasPrefix(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStars), v))
	})
}

// StarsHasSuffix applies the HasSuffix predicate on the "stars" field.
func StarsHasSuffix(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStars), v))
	})
}

// StarsEqualFold applies the EqualFold predicate on the "stars" field.
func StarsEqualFold(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStars), v))
	})
}

// StarsContainsFold applies the ContainsFold predicate on the "stars" field.
func StarsContainsFold(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStars), v))
	})
}

// ImdbRatingEQ applies the EQ predicate on the "imdb_rating" field.
func ImdbRatingEQ(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImdbRating), v))
	})
}

// ImdbRatingNEQ applies the NEQ predicate on the "imdb_rating" field.
func ImdbRatingNEQ(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImdbRating), v))
	})
}

// ImdbRatingIn applies the In predicate on the "imdb_rating" field.
func ImdbRatingIn(vs ...string) predicate.Movie {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldImdbRating), v...))
	})
}

// ImdbRatingNotIn applies the NotIn predicate on the "imdb_rating" field.
func ImdbRatingNotIn(vs ...string) predicate.Movie {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Movie(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldImdbRating), v...))
	})
}

// ImdbRatingGT applies the GT predicate on the "imdb_rating" field.
func ImdbRatingGT(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldImdbRating), v))
	})
}

// ImdbRatingGTE applies the GTE predicate on the "imdb_rating" field.
func ImdbRatingGTE(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldImdbRating), v))
	})
}

// ImdbRatingLT applies the LT predicate on the "imdb_rating" field.
func ImdbRatingLT(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldImdbRating), v))
	})
}

// ImdbRatingLTE applies the LTE predicate on the "imdb_rating" field.
func ImdbRatingLTE(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldImdbRating), v))
	})
}

// ImdbRatingContains applies the Contains predicate on the "imdb_rating" field.
func ImdbRatingContains(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldImdbRating), v))
	})
}

// ImdbRatingHasPrefix applies the HasPrefix predicate on the "imdb_rating" field.
func ImdbRatingHasPrefix(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldImdbRating), v))
	})
}

// ImdbRatingHasSuffix applies the HasSuffix predicate on the "imdb_rating" field.
func ImdbRatingHasSuffix(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldImdbRating), v))
	})
}

// ImdbRatingEqualFold applies the EqualFold predicate on the "imdb_rating" field.
func ImdbRatingEqualFold(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldImdbRating), v))
	})
}

// ImdbRatingContainsFold applies the ContainsFold predicate on the "imdb_rating" field.
func ImdbRatingContainsFold(v string) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldImdbRating), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Movie) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Movie) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
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
func Not(p predicate.Movie) predicate.Movie {
	return predicate.Movie(func(s *sql.Selector) {
		p(s.Not())
	})
}
