package sqltestutil

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/liquidata-inc/ld/dolt/go/libraries/doltcore/dtestutils"
	"github.com/liquidata-inc/ld/dolt/go/libraries/doltcore/env"
	"github.com/liquidata-inc/ld/dolt/go/libraries/doltcore/row"
	"github.com/liquidata-inc/ld/dolt/go/libraries/doltcore/schema"
	"github.com/liquidata-inc/ld/dolt/go/libraries/doltcore/table/untyped"
	"github.com/liquidata-inc/ld/dolt/go/store/types"
	"reflect"
	"testing"
)

// This file collects useful test table definitions and functions for SQL tests to use. It primarily defines a table
// name, schema, and some sample rows to use in tests, as well as functions for creating and seeding a test database,
// transforming row results, and so on.

const (
	IdTag = iota
	FirstTag
	LastTag
	IsMarriedTag
	AgeTag
	emptyTag
	RatingTag
	UuidTag
	NumEpisodesTag
	firstUnusedTag // keep at end
)

const (
	episodeIdTag = iota
	epNameTag
	epAirDateTag
	epRatingTag
)

const (
	appCharacterTag = iota
	appEpTag
	appCommentsTag
)

const (
	homerId = iota
	margeId
	bartId
	lisaId
	moeId
	barneyId
)

var PeopleTestSchema = createPeopleTestSchema()
var untypedPeopleSch = untyped.UntypeUnkeySchema(PeopleTestSchema)
var PeopleTableName = "people"

var EpisodesTestSchema = createEpisodesTestSchema()
var untypedEpisodesSch = untyped.UntypeUnkeySchema(EpisodesTestSchema)
var episodesTableName = "episodes"

var AppearancesTestSchema = createAppearancesTestSchema()
var untypedAppearacesSch = untyped.UntypeUnkeySchema(AppearancesTestSchema)
var appearancesTableName = "appearances"

func createPeopleTestSchema() schema.Schema {
	colColl, _ := schema.NewColCollection(
		schema.NewColumn("id", IdTag, types.IntKind, true, schema.NotNullConstraint{}),
		schema.NewColumn("first", FirstTag, types.StringKind, false, schema.NotNullConstraint{}),
		schema.NewColumn("last", LastTag, types.StringKind, false, schema.NotNullConstraint{}),
		schema.NewColumn("is_married", IsMarriedTag, types.BoolKind, false),
		schema.NewColumn("age", AgeTag, types.IntKind, false),
		//		schema.NewColumn("empty", emptyTag, types.IntKind, false),
		schema.NewColumn("rating", RatingTag, types.FloatKind, false),
		schema.NewColumn("uuid", UuidTag, types.UUIDKind, false),
		schema.NewColumn("num_episodes", NumEpisodesTag, types.UintKind, false),
	)
	return schema.SchemaFromCols(colColl)
}

func createEpisodesTestSchema() schema.Schema {
	colColl, _ := schema.NewColCollection(
		schema.NewColumn("id", episodeIdTag, types.IntKind, true, schema.NotNullConstraint{}),
		schema.NewColumn("name", epNameTag, types.StringKind, false, schema.NotNullConstraint{}),
		schema.NewColumn("air_date", epAirDateTag, types.IntKind, false),
		schema.NewColumn("rating", epRatingTag, types.FloatKind, false),
	)
	return schema.SchemaFromCols(colColl)
}

func createAppearancesTestSchema() schema.Schema {
	colColl, _ := schema.NewColCollection(
		schema.NewColumn("character_id", appCharacterTag, types.IntKind, true, schema.NotNullConstraint{}),
		schema.NewColumn("episode_id", appEpTag, types.IntKind, true, schema.NotNullConstraint{}),
		schema.NewColumn("comments", appCommentsTag, types.StringKind, false),
	)
	return schema.SchemaFromCols(colColl)
}

func NewPeopleRow(id int, first, last string, isMarried bool, age int, rating float32) row.Row {
	vals := row.TaggedValues{
		IdTag:        types.Int(id),
		FirstTag:     types.String(first),
		LastTag:      types.String(last),
		IsMarriedTag: types.Bool(isMarried),
		AgeTag:       types.Int(age),
		RatingTag:    types.Float(rating),
	}

	return row.New(PeopleTestSchema, vals)
}

func newEpsRow(id int, name string, airdate int, rating float32) row.Row {
	vals := row.TaggedValues{
		episodeIdTag: types.Int(id),
		epNameTag:    types.String(name),
		epAirDateTag: types.Int(airdate),
		epRatingTag:  types.Float(rating),
	}

	return row.New(EpisodesTestSchema, vals)
}

func newAppsRow(charId, epId int, comment string) row.Row {
	vals := row.TaggedValues{
		appCharacterTag: types.Int(charId),
		appEpTag:        types.Int(epId),
		appCommentsTag:  types.String(comment),
	}

	return row.New(AppearancesTestSchema, vals)
}

// Most rows don't have these optional fields set, as they aren't needed for basic testing
func NewPeopleRowWithOptionalFields(id int, first, last string, isMarried bool, age int, rating float32, uid uuid.UUID, numEpisodes uint64) row.Row {
	vals := row.TaggedValues{
		IdTag:          types.Int(id),
		FirstTag:       types.String(first),
		LastTag:        types.String(last),
		IsMarriedTag:   types.Bool(isMarried),
		AgeTag:         types.Int(age),
		RatingTag:      types.Float(rating),
		UuidTag:        types.UUID(uid),
		NumEpisodesTag: types.Uint(numEpisodes),
	}

	return row.New(PeopleTestSchema, vals)
}

// 6 characters
var Homer = NewPeopleRow(homerId, "Homer", "Simpson", true, 40, 8.5)
var Marge = NewPeopleRowWithOptionalFields(margeId, "Marge", "Simpson", true, 38, 8, uuid.MustParse("00000000-0000-0000-0000-000000000001"), 111)
var Bart = NewPeopleRowWithOptionalFields(bartId, "Bart", "Simpson", false, 10, 9, uuid.MustParse("00000000-0000-0000-0000-000000000002"), 222)
var Lisa = NewPeopleRowWithOptionalFields(lisaId, "Lisa", "Simpson", false, 8, 10, uuid.MustParse("00000000-0000-0000-0000-000000000003"), 333)
var Moe = NewPeopleRowWithOptionalFields(moeId, "Moe", "Szyslak", false, 48, 6.5, uuid.MustParse("00000000-0000-0000-0000-000000000004"), 444)
var Barney = NewPeopleRowWithOptionalFields(barneyId, "Barney", "Gumble", false, 40, 4, uuid.MustParse("00000000-0000-0000-0000-000000000005"), 555)
var AllPeopleRows = Rs(Homer, Marge, Bart, Lisa, Moe, Barney)

// Actually the first 4 episodes of the show
var Ep1 = newEpsRow(1, "Simpsons Roasting On an Open Fire", 629953200, 8.0)
var Ep2 = newEpsRow(2, "Bart the Genius", 632372400, 9.0)
var Ep3 = newEpsRow(3, "Homer's Odyssey", 632977200, 7.0)
var Ep4 = newEpsRow(4, "There's No Disgrace Like Home", 633582000, 8.5)
var allEpsRows = Rs(Ep1, Ep2, Ep3, Ep4)

// These are made up, not the actual show data
var app1 = newAppsRow(homerId, 1, "Homer is great in this one")
var app2 = newAppsRow(margeId, 1, "Marge is here too")
var app3 = newAppsRow(homerId, 2, "Homer is great in this one too")
var app4 = newAppsRow(bartId, 2, "This episode is named after Bart")
var app5 = newAppsRow(lisaId, 2, "Lisa is here too")
var app6 = newAppsRow(moeId, 2, "I think there's a prank call scene")
var app7 = newAppsRow(homerId, 3, "Homer is in every episode")
var app8 = newAppsRow(margeId, 3, "Marge shows up a lot too")
var app9 = newAppsRow(lisaId, 3, "Lisa is the best Simpson")
var app10 = newAppsRow(barneyId, 3, "I'm making this all up")

// nobody in episode 4, that one was terrible
// Unlike the other tables, you can't count on the order of these rows matching the insertion order.
var AllAppsRows = Rs(app1, app2, app3, app4, app5, app6, app7, app8, app9, app10)

// Convenience func to avoid the boilerplate of typing []row.Row{} all the time
func Rs(rows ...row.Row) []row.Row {
	if rows == nil {
		return make([]row.Row, 0)
	}
	return rows
}

// Returns the index of the first row in the list that has the same primary key as the one given, or -1 otherwise.
func FindRowIndex(find row.Row, rows []row.Row) int {
	idx := -1
	for i, updatedRow := range rows {
		rowId, _ := find.GetColVal(IdTag)
		updatedId, _ := updatedRow.GetColVal(IdTag)
		if rowId.Equals(updatedId) {
			idx = i
			break
		}
	}
	return idx
}

// Mutates the row given with pairs of {tag,value} given in the varargs param. Converts built-in types to noms types.
func MutateRow(r row.Row, tagsAndVals ...interface{}) row.Row {
	if len(tagsAndVals)%2 != 0 {
		panic("expected pairs of tags and values")
	}

	var mutated row.Row = r
	var err error

	for i := 0; i < len(tagsAndVals); i += 2 {
		tag := tagsAndVals[i].(int)
		val := tagsAndVals[i+1]
		var nomsVal types.Value
		if val != nil {
			switch v := val.(type) {
			case uint64:
				nomsVal = types.Uint(v)
			case int:
				nomsVal = types.Int(v)
			case int32:
				nomsVal = types.Int(v)
			case int64:
				nomsVal = types.Int(v)
			case float32:
				nomsVal = types.Float(v)
			case float64:
				nomsVal = types.Float(v)
			case string:
				nomsVal = types.String(v)
			case uuid.UUID:
				nomsVal = types.UUID(v)
			case bool:
				nomsVal = types.Bool(v)
			default:
				panic("Unhandled type " + reflect.TypeOf(val).String())
			}
		} else {
			nomsVal = nil
		}

		mutated, err = mutated.SetColVal(uint64(tag), nomsVal, PeopleTestSchema)
		if err != nil {
			panic(err.Error())
		}
	}

	return mutated
}

// Creates a new row for a result set specified by the given values
func NewResultSetRow(colVals ...types.Value) row.Row {

	taggedVals := make(row.TaggedValues)
	cols := make([]schema.Column, len(colVals))
	for i := 0; i < len(colVals); i++ {
		taggedVals[uint64(i)] = colVals[i]
		nomsKind := colVals[i].Kind()
		cols[i] = schema.NewColumn(fmt.Sprintf("%v", i), uint64(i), nomsKind, false)
	}

	collection, err := schema.NewColCollection(cols...)
	if err != nil {
		panic("unexpected error " + err.Error())
	}
	sch := schema.UnkeyedSchemaFromCols(collection)

	return row.New(sch, taggedVals)
}

// Creates a test database with the test data set in it
func CreateTestDatabase(dEnv *env.DoltEnv, t *testing.T) {
	dtestutils.CreateTestTable(t, dEnv, PeopleTableName, PeopleTestSchema, AllPeopleRows...)
	dtestutils.CreateTestTable(t, dEnv, episodesTableName, EpisodesTestSchema, allEpsRows...)
	dtestutils.CreateTestTable(t, dEnv, appearancesTableName, AppearancesTestSchema, AllAppsRows...)
}