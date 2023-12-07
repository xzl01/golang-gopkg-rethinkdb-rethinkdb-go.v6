// Code generated by gen_tests.py and process_polyglot.py.
// Do not edit this file directly.
// The template for this file is located at:
// ../template.go.tpl
package reql_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"gopkg.in/rethinkdb/rethinkdb-go.v6/internal/compare"
)

// Tests manipulation operations on objects
func TestTransformObjectSuite(t *testing.T) {
	suite.Run(t, new(TransformObjectSuite))
}

type TransformObjectSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *TransformObjectSuite) SetupTest() {
	suite.T().Log("Setting up TransformObjectSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("db_tfm_obj").Exec(suite.session)
	err = r.DBCreate("db_tfm_obj").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_tfm_obj").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *TransformObjectSuite) TearDownSuite() {
	suite.T().Log("Tearing down TransformObjectSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("db_tfm_obj").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *TransformObjectSuite) TestCases() {
	suite.T().Log("Running TransformObjectSuite: Tests manipulation operations on objects")

	// transform/object.yaml line #5
	// obj = r.expr({'a':1, 'b':2,'c':"str",'d':null,'e':{'f':'buzz'}})
	suite.T().Log("Possibly executing: var obj r.Term = r.Expr(map[interface{}]interface{}{'a': 1, 'b': 2, 'c': 'str', 'd': nil, 'e': map[interface{}]interface{}{'f': 'buzz', }, })")

	obj := r.Expr(map[interface{}]interface{}{"a": 1, "b": 2, "c": "str", "d": nil, "e": map[interface{}]interface{}{"f": "buzz"}})
	_ = obj // Prevent any noused variable errors

	{
		// transform/object.yaml line #9
		/* 1 */
		var expected_ int = 1
		/* obj['a'] */

		suite.T().Log("About to run line #9: obj.AtIndex('a')")

		runAndAssert(suite.Suite, expected_, obj.AtIndex("a"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #9")
	}

	{
		// transform/object.yaml line #14
		/* 'str' */
		var expected_ string = "str"
		/* obj['c'] */

		suite.T().Log("About to run line #14: obj.AtIndex('c')")

		runAndAssert(suite.Suite, expected_, obj.AtIndex("c"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #14")
	}

	{
		// transform/object.yaml line #22
		/* true */
		var expected_ bool = true
		/* obj.has_fields('b') */

		suite.T().Log("About to run line #22: obj.HasFields('b')")

		runAndAssert(suite.Suite, expected_, obj.HasFields("b"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #22")
	}

	{
		// transform/object.yaml line #24
		/* true */
		var expected_ bool = true
		/* obj.keys().contains('d') */

		suite.T().Log("About to run line #24: obj.Keys().Contains('d')")

		runAndAssert(suite.Suite, expected_, obj.Keys().Contains("d"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #24")
	}

	{
		// transform/object.yaml line #26
		/* false */
		var expected_ bool = false
		/* obj.has_fields('d') */

		suite.T().Log("About to run line #26: obj.HasFields('d')")

		runAndAssert(suite.Suite, expected_, obj.HasFields("d"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #26")
	}

	{
		// transform/object.yaml line #28
		/* true */
		var expected_ bool = true
		/* obj.has_fields({'e':'f'}) */

		suite.T().Log("About to run line #28: obj.HasFields(map[interface{}]interface{}{'e': 'f', })")

		runAndAssert(suite.Suite, expected_, obj.HasFields(map[interface{}]interface{}{"e": "f"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #28")
	}

	{
		// transform/object.yaml line #30
		/* false */
		var expected_ bool = false
		/* obj.has_fields({'e':'g'}) */

		suite.T().Log("About to run line #30: obj.HasFields(map[interface{}]interface{}{'e': 'g', })")

		runAndAssert(suite.Suite, expected_, obj.HasFields(map[interface{}]interface{}{"e": "g"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #30")
	}

	{
		// transform/object.yaml line #32
		/* false */
		var expected_ bool = false
		/* obj.has_fields('f') */

		suite.T().Log("About to run line #32: obj.HasFields('f')")

		runAndAssert(suite.Suite, expected_, obj.HasFields("f"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #32")
	}

	{
		// transform/object.yaml line #36
		/* true */
		var expected_ bool = true
		/* obj.has_fields('a', 'b') */

		suite.T().Log("About to run line #36: obj.HasFields('a', 'b')")

		runAndAssert(suite.Suite, expected_, obj.HasFields("a", "b"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #36")
	}

	{
		// transform/object.yaml line #38
		/* false */
		var expected_ bool = false
		/* obj.has_fields('a', 'd') */

		suite.T().Log("About to run line #38: obj.HasFields('a', 'd')")

		runAndAssert(suite.Suite, expected_, obj.HasFields("a", "d"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #38")
	}

	{
		// transform/object.yaml line #40
		/* false */
		var expected_ bool = false
		/* obj.has_fields('a', 'f') */

		suite.T().Log("About to run line #40: obj.HasFields('a', 'f')")

		runAndAssert(suite.Suite, expected_, obj.HasFields("a", "f"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #40")
	}

	{
		// transform/object.yaml line #42
		/* true */
		var expected_ bool = true
		/* obj.has_fields('a', {'e':'f'}) */

		suite.T().Log("About to run line #42: obj.HasFields('a', map[interface{}]interface{}{'e': 'f', })")

		runAndAssert(suite.Suite, expected_, obj.HasFields("a", map[interface{}]interface{}{"e": "f"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #42")
	}

	{
		// transform/object.yaml line #46
		/* 2 */
		var expected_ int = 2
		/* r.expr([obj, obj.pluck('a', 'b')]).has_fields('a', 'b').count() */

		suite.T().Log("About to run line #46: r.Expr([]interface{}{obj, obj.Pluck('a', 'b')}).HasFields('a', 'b').Count()")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{obj, obj.Pluck("a", "b")}).HasFields("a", "b").Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #46")
	}

	{
		// transform/object.yaml line #48
		/* 1 */
		var expected_ int = 1
		/* r.expr([obj, obj.pluck('a', 'b')]).has_fields('a', 'c').count() */

		suite.T().Log("About to run line #48: r.Expr([]interface{}{obj, obj.Pluck('a', 'b')}).HasFields('a', 'c').Count()")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{obj, obj.Pluck("a", "b")}).HasFields("a", "c").Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #48")
	}

	{
		// transform/object.yaml line #50
		/* 2 */
		var expected_ int = 2
		/* r.expr([obj, obj.pluck('a', 'e')]).has_fields('a', {'e':'f'}).count() */

		suite.T().Log("About to run line #50: r.Expr([]interface{}{obj, obj.Pluck('a', 'e')}).HasFields('a', map[interface{}]interface{}{'e': 'f', }).Count()")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{obj, obj.Pluck("a", "e")}).HasFields("a", map[interface{}]interface{}{"e": "f"}).Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #50")
	}

	{
		// transform/object.yaml line #55
		/* {'a':1} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": 1}
		/* obj.pluck('a') */

		suite.T().Log("About to run line #55: obj.Pluck('a')")

		runAndAssert(suite.Suite, expected_, obj.Pluck("a"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #55")
	}

	{
		// transform/object.yaml line #57
		/* {'a':1, 'b':2} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": 1, "b": 2}
		/* obj.pluck('a', 'b') */

		suite.T().Log("About to run line #57: obj.Pluck('a', 'b')")

		runAndAssert(suite.Suite, expected_, obj.Pluck("a", "b"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #57")
	}

	{
		// transform/object.yaml line #62
		/* {'b':2, 'c':'str', 'd':null, 'e':{'f':'buzz'}} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"b": 2, "c": "str", "d": nil, "e": map[interface{}]interface{}{"f": "buzz"}}
		/* obj.without('a') */

		suite.T().Log("About to run line #62: obj.Without('a')")

		runAndAssert(suite.Suite, expected_, obj.Without("a"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #62")
	}

	{
		// transform/object.yaml line #64
		/* {'c':'str', 'd':null,'e':{'f':'buzz'}} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"c": "str", "d": nil, "e": map[interface{}]interface{}{"f": "buzz"}}
		/* obj.without('a', 'b') */

		suite.T().Log("About to run line #64: obj.Without('a', 'b')")

		runAndAssert(suite.Suite, expected_, obj.Without("a", "b"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #64")
	}

	{
		// transform/object.yaml line #66
		/* {'e':{'f':'buzz'}} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"e": map[interface{}]interface{}{"f": "buzz"}}
		/* obj.without('a', 'b', 'c', 'd') */

		suite.T().Log("About to run line #66: obj.Without('a', 'b', 'c', 'd')")

		runAndAssert(suite.Suite, expected_, obj.Without("a", "b", "c", "d"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #66")
	}

	{
		// transform/object.yaml line #68
		/* {'a':1, 'b':2, 'c':'str', 'd':null, 'e':{}} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": 1, "b": 2, "c": "str", "d": nil, "e": map[interface{}]interface{}{}}
		/* obj.without({'e':'f'}) */

		suite.T().Log("About to run line #68: obj.Without(map[interface{}]interface{}{'e': 'f', })")

		runAndAssert(suite.Suite, expected_, obj.Without(map[interface{}]interface{}{"e": "f"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #68")
	}

	{
		// transform/object.yaml line #70
		/* {'a':1, 'b':2, 'c':'str', 'd':null, 'e':{'f':'buzz'}} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": 1, "b": 2, "c": "str", "d": nil, "e": map[interface{}]interface{}{"f": "buzz"}}
		/* obj.without({'e':'buzz'}) */

		suite.T().Log("About to run line #70: obj.Without(map[interface{}]interface{}{'e': 'buzz', })")

		runAndAssert(suite.Suite, expected_, obj.Without(map[interface{}]interface{}{"e": "buzz"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #70")
	}

	{
		// transform/object.yaml line #77
		/* 1 */
		var expected_ int = 1
		/* obj.merge(1) */

		suite.T().Log("About to run line #77: obj.Merge(1)")

		runAndAssert(suite.Suite, expected_, obj.Merge(1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #77")
	}

	{
		// transform/object.yaml line #81
		/* {'a':1, 'b':2, 'c':'str', 'd':null, 'e':-2} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": 1, "b": 2, "c": "str", "d": nil, "e": -2}
		/* obj.merge({'e':-2}) */

		suite.T().Log("About to run line #81: obj.Merge(map[interface{}]interface{}{'e': -2, })")

		runAndAssert(suite.Suite, expected_, obj.Merge(map[interface{}]interface{}{"e": -2}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #81")
	}

	{
		// transform/object.yaml line #85
		/* {'a':1, 'b':2, 'c':'str', 'd':null} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": 1, "b": 2, "c": "str", "d": nil}
		/* obj.merge({'e':r.literal()}) */

		suite.T().Log("About to run line #85: obj.Merge(map[interface{}]interface{}{'e': r.Literal(), })")

		runAndAssert(suite.Suite, expected_, obj.Merge(map[interface{}]interface{}{"e": r.Literal()}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #85")
	}

	{
		// transform/object.yaml line #89
		/* {'a':1, 'b':2, 'c':'str', 'd':null, 'e':{'f':'quux'}} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": 1, "b": 2, "c": "str", "d": nil, "e": map[interface{}]interface{}{"f": "quux"}}
		/* obj.merge({'e':{'f':'quux'}}) */

		suite.T().Log("About to run line #89: obj.Merge(map[interface{}]interface{}{'e': map[interface{}]interface{}{'f': 'quux', }, })")

		runAndAssert(suite.Suite, expected_, obj.Merge(map[interface{}]interface{}{"e": map[interface{}]interface{}{"f": "quux"}}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #89")
	}

	{
		// transform/object.yaml line #92
		/* {'a':1, 'b':2, 'c':'str', 'd':null, 'e':{'f':'buzz', 'g':'quux'}} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": 1, "b": 2, "c": "str", "d": nil, "e": map[interface{}]interface{}{"f": "buzz", "g": "quux"}}
		/* obj.merge({'e':{'g':'quux'}}) */

		suite.T().Log("About to run line #92: obj.Merge(map[interface{}]interface{}{'e': map[interface{}]interface{}{'g': 'quux', }, })")

		runAndAssert(suite.Suite, expected_, obj.Merge(map[interface{}]interface{}{"e": map[interface{}]interface{}{"g": "quux"}}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #92")
	}

	{
		// transform/object.yaml line #95
		/* {'a':1, 'b':2, 'c':'str', 'd':null, 'e':{'g':'quux'}} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": 1, "b": 2, "c": "str", "d": nil, "e": map[interface{}]interface{}{"g": "quux"}}
		/* obj.merge({'e':r.literal({'g':'quux'})}) */

		suite.T().Log("About to run line #95: obj.Merge(map[interface{}]interface{}{'e': r.Literal(map[interface{}]interface{}{'g': 'quux', }), })")

		runAndAssert(suite.Suite, expected_, obj.Merge(map[interface{}]interface{}{"e": r.Literal(map[interface{}]interface{}{"g": "quux"})}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #95")
	}

	{
		// transform/object.yaml line #99
		/* {'a':-1, 'b':2, 'c':'str', 'd':null, 'e':{'f':'buzz'}} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": -1, "b": 2, "c": "str", "d": nil, "e": map[interface{}]interface{}{"f": "buzz"}}
		/* obj.merge({'a':-1}) */

		suite.T().Log("About to run line #99: obj.Merge(map[interface{}]interface{}{'a': -1, })")

		runAndAssert(suite.Suite, expected_, obj.Merge(map[interface{}]interface{}{"a": -1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #99")
	}

	// transform/object.yaml line #102
	// errmsg = 'Stray literal keyword found:'+' literal is only legal inside of the object passed to merge or update and cannot nest inside other literals.'
	suite.T().Log("Possibly executing: var errmsg string = 'Stray literal keyword found:' + ' literal is only legal inside of the object passed to merge or update and cannot nest inside other literals.'")

	errmsg := "Stray literal keyword found:" + " literal is only legal inside of the object passed to merge or update and cannot nest inside other literals."
	_ = errmsg // Prevent any noused variable errors

	{
		// transform/object.yaml line #105
		/* err("ReqlQueryLogicError", errmsg, []) */
		var expected_ Err = err("ReqlQueryLogicError", errmsg)
		/* r.literal('foo') */

		suite.T().Log("About to run line #105: r.Literal('foo')")

		runAndAssert(suite.Suite, expected_, r.Literal("foo"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #105")
	}

	{
		// transform/object.yaml line #108
		/* err("ReqlQueryLogicError", errmsg, []) */
		var expected_ Err = err("ReqlQueryLogicError", errmsg)
		/* obj.merge(r.literal('foo')) */

		suite.T().Log("About to run line #108: obj.Merge(r.Literal('foo'))")

		runAndAssert(suite.Suite, expected_, obj.Merge(r.Literal("foo")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #108")
	}

	{
		// transform/object.yaml line #111
		/* err("ReqlQueryLogicError", errmsg, []) */
		var expected_ Err = err("ReqlQueryLogicError", errmsg)
		/* obj.merge({'foo':r.literal(r.literal('foo'))}) */

		suite.T().Log("About to run line #111: obj.Merge(map[interface{}]interface{}{'foo': r.Literal(r.Literal('foo')), })")

		runAndAssert(suite.Suite, expected_, obj.Merge(map[interface{}]interface{}{"foo": r.Literal(r.Literal("foo"))}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #111")
	}

	// transform/object.yaml line #114
	// o = r.expr({'a':{'b':1, 'c':2}, 'd':3})
	suite.T().Log("Possibly executing: var o r.Term = r.Expr(map[interface{}]interface{}{'a': map[interface{}]interface{}{'b': 1, 'c': 2, }, 'd': 3, })")

	o := r.Expr(map[interface{}]interface{}{"a": map[interface{}]interface{}{"b": 1, "c": 2}, "d": 3})
	_ = o // Prevent any noused variable errors

	{
		// transform/object.yaml line #116
		/* ({'a':{'b':1, 'c':2}, 'd':3, 'e':4, 'f':5}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": map[interface{}]interface{}{"b": 1, "c": 2}, "d": 3, "e": 4, "f": 5}
		/* o.merge({'e':4}, {'f':5}) */

		suite.T().Log("About to run line #116: o.Merge(map[interface{}]interface{}{'e': 4, }, map[interface{}]interface{}{'f': 5, })")

		runAndAssert(suite.Suite, expected_, o.Merge(map[interface{}]interface{}{"e": 4}, map[interface{}]interface{}{"f": 5}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #116")
	}

	{
		// transform/object.yaml line #120
		/* ([{'a':{'b':1, 'c':2}, 'd':3, 'e':3}, {'a':{'b':1, 'c':2}, 'd':4, 'e':4}]) */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"a": map[interface{}]interface{}{"b": 1, "c": 2}, "d": 3, "e": 3}, map[interface{}]interface{}{"a": map[interface{}]interface{}{"b": 1, "c": 2}, "d": 4, "e": 4}}
		/* r.expr([o, o.merge({'d':4})]).merge(lambda row:{'e':row['d']}) */

		suite.T().Log("About to run line #120: r.Expr([]interface{}{o, o.Merge(map[interface{}]interface{}{'d': 4, })}).Merge(func(row r.Term) interface{} { return map[interface{}]interface{}{'e': row.AtIndex('d'), }})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{o, o.Merge(map[interface{}]interface{}{"d": 4})}).Merge(func(row r.Term) interface{} { return map[interface{}]interface{}{"e": row.AtIndex("d")} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #120")
	}

	{
		// transform/object.yaml line #124
		/* ([{'a':{'b':1, 'c':2}, 'd':3, 'e':3}, {'a':{'b':1, 'c':2}, 'd':4, 'e':4}]) */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"a": map[interface{}]interface{}{"b": 1, "c": 2}, "d": 3, "e": 3}, map[interface{}]interface{}{"a": map[interface{}]interface{}{"b": 1, "c": 2}, "d": 4, "e": 4}}
		/* r.expr([o, o.merge({'d':4})]).merge({'e':r.row['d']}) */

		suite.T().Log("About to run line #124: r.Expr([]interface{}{o, o.Merge(map[interface{}]interface{}{'d': 4, })}).Merge(map[interface{}]interface{}{'e': r.Row.AtIndex('d'), })")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{o, o.Merge(map[interface{}]interface{}{"d": 4})}).Merge(map[interface{}]interface{}{"e": r.Row.AtIndex("d")}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #124")
	}

	{
		// transform/object.yaml line #129
		/* ([{'a':{'b':2, 'c':2}, 'd':3}, {'a':{'b':2, 'c':2}, 'd':4}]) */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"a": map[interface{}]interface{}{"b": 2, "c": 2}, "d": 3}, map[interface{}]interface{}{"a": map[interface{}]interface{}{"b": 2, "c": 2}, "d": 4}}
		/* r.expr([o, o.merge({'d':4})]).merge(lambda row:{'a':{'b':2}}) */

		suite.T().Log("About to run line #129: r.Expr([]interface{}{o, o.Merge(map[interface{}]interface{}{'d': 4, })}).Merge(func(row r.Term) interface{} { return map[interface{}]interface{}{'a': map[interface{}]interface{}{'b': 2, }, }})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{o, o.Merge(map[interface{}]interface{}{"d": 4})}).Merge(func(row r.Term) interface{} {
			return map[interface{}]interface{}{"a": map[interface{}]interface{}{"b": 2}}
		}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #129")
	}

	{
		// transform/object.yaml line #134
		/* ([{'a':{'b':2}, 'd':3}, {'a':{'b':2}, 'd':4}]) */
		var expected_ []interface{} = []interface{}{map[interface{}]interface{}{"a": map[interface{}]interface{}{"b": 2}, "d": 3}, map[interface{}]interface{}{"a": map[interface{}]interface{}{"b": 2}, "d": 4}}
		/* r.expr([o, o.merge({'d':4})]).merge(lambda row:{'a':r.literal({'b':2})}) */

		suite.T().Log("About to run line #134: r.Expr([]interface{}{o, o.Merge(map[interface{}]interface{}{'d': 4, })}).Merge(func(row r.Term) interface{} { return map[interface{}]interface{}{'a': r.Literal(map[interface{}]interface{}{'b': 2, }), }})")

		runAndAssert(suite.Suite, expected_, r.Expr([]interface{}{o, o.Merge(map[interface{}]interface{}{"d": 4})}).Merge(func(row r.Term) interface{} {
			return map[interface{}]interface{}{"a": r.Literal(map[interface{}]interface{}{"b": 2})}
		}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #134")
	}

	{
		// transform/object.yaml line #139
		/* (['a', 'b', 'c', 'd', 'e']) */
		var expected_ []interface{} = []interface{}{"a", "b", "c", "d", "e"}
		/* obj.keys() */

		suite.T().Log("About to run line #139: obj.Keys()")

		runAndAssert(suite.Suite, expected_, obj.Keys(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #139")
	}

	{
		// transform/object.yaml line #142
		/* ([1, 2, 'str', null, {'f':'buzz'}]) */
		var expected_ []interface{} = []interface{}{1, 2, "str", nil, map[interface{}]interface{}{"f": "buzz"}}
		/* obj.values() */

		suite.T().Log("About to run line #142: obj.Values()")

		runAndAssert(suite.Suite, expected_, obj.Values(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #142")
	}

	{
		// transform/object.yaml line #146
		/* 5 */
		var expected_ int = 5
		/* obj.count() */

		suite.T().Log("About to run line #146: obj.Count()")

		runAndAssert(suite.Suite, expected_, obj.Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #146")
	}
}
