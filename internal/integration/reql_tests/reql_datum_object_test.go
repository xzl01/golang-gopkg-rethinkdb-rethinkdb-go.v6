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

// Tests conversion to and from the RQL object type
func TestDatumObjectSuite(t *testing.T) {
	suite.Run(t, new(DatumObjectSuite))
}

type DatumObjectSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *DatumObjectSuite) SetupTest() {
	suite.T().Log("Setting up DatumObjectSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("db_dobj").Exec(suite.session)
	err = r.DBCreate("db_dobj").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_dobj").Wait().Exec(suite.session)
	suite.Require().NoError(err)

}

func (suite *DatumObjectSuite) TearDownSuite() {
	suite.T().Log("Tearing down DatumObjectSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DBDrop("db_dobj").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *DatumObjectSuite) TestCases() {
	suite.T().Log("Running DatumObjectSuite: Tests conversion to and from the RQL object type")

	{
		// datum/object.yaml line #6
		/* {} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{}
		/* r.expr({}) */

		suite.T().Log("About to run line #6: r.Expr(map[interface{}]interface{}{})")

		runAndAssert(suite.Suite, expected_, r.Expr(map[interface{}]interface{}{}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #6")
	}

	{
		// datum/object.yaml line #11
		/* {'a':1} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": 1}
		/* r.expr({'a':1}) */

		suite.T().Log("About to run line #11: r.Expr(map[interface{}]interface{}{'a': 1, })")

		runAndAssert(suite.Suite, expected_, r.Expr(map[interface{}]interface{}{"a": 1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #11")
	}

	{
		// datum/object.yaml line #16
		/* {'a':1, 'b':'two', 'c':True} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": 1, "b": "two", "c": true}
		/* r.expr({'a':1, 'b':'two', 'c':True}) */

		suite.T().Log("About to run line #16: r.Expr(map[interface{}]interface{}{'a': 1, 'b': 'two', 'c': true, })")

		runAndAssert(suite.Suite, expected_, r.Expr(map[interface{}]interface{}{"a": 1, "b": "two", "c": true}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #16")
	}

	{
		// datum/object.yaml line #20
		/* {'a':1} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": 1}
		/* r.expr({'a':r.expr(1)}) */

		suite.T().Log("About to run line #20: r.Expr(map[interface{}]interface{}{'a': r.Expr(1), })")

		runAndAssert(suite.Suite, expected_, r.Expr(map[interface{}]interface{}{"a": r.Expr(1)}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #20")
	}

	{
		// datum/object.yaml line #23
		/* {'a':{'b':[{'c':2}, 'a', 4]}} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{map[interface{}]interface{}{"c": 2}, "a", 4}}}
		/* r.expr({'a':{'b':[{'c':2}, 'a', 4]}}) */

		suite.T().Log("About to run line #23: r.Expr(map[interface{}]interface{}{'a': map[interface{}]interface{}{'b': []interface{}{map[interface{}]interface{}{'c': 2, }, 'a', 4}, }, })")

		runAndAssert(suite.Suite, expected_, r.Expr(map[interface{}]interface{}{"a": map[interface{}]interface{}{"b": []interface{}{map[interface{}]interface{}{"c": 2}, "a", 4}}}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #23")
	}

	{
		// datum/object.yaml line #26
		/* 'OBJECT' */
		var expected_ string = "OBJECT"
		/* r.expr({'a':1}).type_of() */

		suite.T().Log("About to run line #26: r.Expr(map[interface{}]interface{}{'a': 1, }).TypeOf()")

		runAndAssert(suite.Suite, expected_, r.Expr(map[interface{}]interface{}{"a": 1}).TypeOf(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #26")
	}

	{
		// datum/object.yaml line #30
		/* '{"a":1}' */
		var expected_ string = "{\"a\":1}"
		/* r.expr({'a':1}).coerce_to('string') */

		suite.T().Log("About to run line #30: r.Expr(map[interface{}]interface{}{'a': 1, }).CoerceTo('string')")

		runAndAssert(suite.Suite, expected_, r.Expr(map[interface{}]interface{}{"a": 1}).CoerceTo("string"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #30")
	}

	{
		// datum/object.yaml line #34
		/* {'a':1} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": 1}
		/* r.expr({'a':1}).coerce_to('object') */

		suite.T().Log("About to run line #34: r.Expr(map[interface{}]interface{}{'a': 1, }).CoerceTo('object')")

		runAndAssert(suite.Suite, expected_, r.Expr(map[interface{}]interface{}{"a": 1}).CoerceTo("object"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #34")
	}

	{
		// datum/object.yaml line #37
		/* [['a',1]] */
		var expected_ []interface{} = []interface{}{[]interface{}{"a", 1}}
		/* r.expr({'a':1}).coerce_to('array') */

		suite.T().Log("About to run line #37: r.Expr(map[interface{}]interface{}{'a': 1, }).CoerceTo('array')")

		runAndAssert(suite.Suite, expected_, r.Expr(map[interface{}]interface{}{"a": 1}).CoerceTo("array"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #37")
	}

	{
		// datum/object.yaml line #66
		/* {} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{}
		/* r.object() */

		suite.T().Log("About to run line #66: r.Object()")

		runAndAssert(suite.Suite, expected_, r.Object(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #66")
	}

	{
		// datum/object.yaml line #69
		/* {'a':1,'b':2} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"a": 1, "b": 2}
		/* r.object('a', 1, 'b', 2) */

		suite.T().Log("About to run line #69: r.Object('a', 1, 'b', 2)")

		runAndAssert(suite.Suite, expected_, r.Object("a", 1, "b", 2), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #69")
	}

	{
		// datum/object.yaml line #72
		/* {'cd':3} */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"cd": 3}
		/* r.object('c'+'d', 3) */

		suite.T().Log("About to run line #72: r.Object(r.Add('c', 'd'), 3)")

		runAndAssert(suite.Suite, expected_, r.Object(r.Add("c", "d"), 3), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #72")
	}

	{
		// datum/object.yaml line #75
		/* err("ReqlQueryLogicError", "OBJECT expects an even number of arguments (but found 3).", []) */
		var expected_ Err = err("ReqlQueryLogicError", "OBJECT expects an even number of arguments (but found 3).")
		/* r.object('o','d','d') */

		suite.T().Log("About to run line #75: r.Object('o', 'd', 'd')")

		runAndAssert(suite.Suite, expected_, r.Object("o", "d", "d"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #75")
	}

	{
		// datum/object.yaml line #78
		/* err("ReqlQueryLogicError","Expected type STRING but found NUMBER.",[]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type STRING but found NUMBER.")
		/* r.object(1, 1) */

		suite.T().Log("About to run line #78: r.Object(1, 1)")

		runAndAssert(suite.Suite, expected_, r.Object(1, 1), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #78")
	}

	{
		// datum/object.yaml line #81
		/* err("ReqlQueryLogicError","Duplicate key \"e\" in object.  (got 4 and 5 as values)",[]) */
		var expected_ Err = err("ReqlQueryLogicError", "Duplicate key \"e\" in object.  (got 4 and 5 as values)")
		/* r.object('e', 4, 'e', 5) */

		suite.T().Log("About to run line #81: r.Object('e', 4, 'e', 5)")

		runAndAssert(suite.Suite, expected_, r.Object("e", 4, "e", 5), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #81")
	}

	{
		// datum/object.yaml line #84
		/* err("ReqlQueryLogicError","Expected type DATUM but found DATABASE:",[]) */
		var expected_ Err = err("ReqlQueryLogicError", "Expected type DATUM but found DATABASE:")
		/* r.object('g', r.db('test')) */

		suite.T().Log("About to run line #84: r.Object('g', r.DB('test'))")

		runAndAssert(suite.Suite, expected_, r.Object("g", r.DB("db_dobj")), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #84")
	}
}
