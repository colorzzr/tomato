package rest

import (
	"reflect"
	"testing"

	"github.com/lfq7413/tomato/types"
)

func Test_Execute(t *testing.T) {
	// BuildRestWhere
	// runFind
	// runCount
	// handleInclude
	// TODO
}

func Test_BuildRestWhere(t *testing.T) {
	// getUserAndRoleACL
	// redirectClassNameForKey
	// validateClientClassCreation
	// replaceSelect
	// replaceDontSelect
	// replaceInQuery
	// replaceNotInQuery
	// TODO
}

func Test_getUserAndRoleACL(t *testing.T) {
	// TODO
}

func Test_redirectClassNameForKey(t *testing.T) {
	// TODO
}

func Test_validateClientClassCreation(t *testing.T) {
	// TODO
}

func Test_replaceSelect(t *testing.T) {
	// findObjectWithKey
	// NewQuery
	// Execute
	// transformSelect
	// TODO
}

func Test_replaceDontSelect(t *testing.T) {
	// findObjectWithKey
	// NewQuery
	// Execute
	// transformDontSelect
	// TODO
}

func Test_replaceInQuery(t *testing.T) {
	// findObjectWithKey
	// NewQuery
	// Execute
	// TODO
}

func Test_replaceNotInQuery(t *testing.T) {
	// findObjectWithKey
	// NewQuery
	// Execute
	// TODO
}

func Test_runFind(t *testing.T) {
	// TODO
}

func Test_runCount(t *testing.T) {
	// TODO
}

func Test_handleInclude(t *testing.T) {
	// includePath
	// TODO
}

/////////////////////////////////////////////////////////////////

func Test_NewQuery(t *testing.T) {
	// TODO
}

func Test_includePath(t *testing.T) {
	// findPointers
	// NewQuery
	// Execute
	// replacePointers
	// TODO
}

func Test_findPointers(t *testing.T) {
	// TODO
}

func Test_replacePointers(t *testing.T) {
	// TODO
}

func Test_findObjectWithKey(t *testing.T) {
	// TODO
}

func Test_transformSelect(t *testing.T) {
	// TODO
}

func Test_transformDontSelect(t *testing.T) {
	// TODO
}

func Test_transformInQuery(t *testing.T) {
	var inQueryObject types.M
	var className string
	var results []types.M
	var expect types.M
	/**********************************************************/
	inQueryObject = nil
	className = "user"
	results = nil
	transformInQuery(inQueryObject, className, results)
	expect = nil
	if reflect.DeepEqual(expect, inQueryObject) == false {
		t.Error("expect:", expect, "result:", inQueryObject)
	}
	/**********************************************************/
	inQueryObject = types.M{}
	className = "user"
	results = nil
	transformInQuery(inQueryObject, className, results)
	expect = types.M{}
	if reflect.DeepEqual(expect, inQueryObject) == false {
		t.Error("expect:", expect, "result:", inQueryObject)
	}
	/**********************************************************/
	inQueryObject = types.M{
		"$inQuery": "string",
	}
	className = "user"
	results = nil
	transformInQuery(inQueryObject, className, results)
	expect = types.M{
		"$in": types.S{},
	}
	if reflect.DeepEqual(expect, inQueryObject) == false {
		t.Error("expect:", expect, "result:", inQueryObject)
	}
	/**********************************************************/
	inQueryObject = types.M{
		"$inQuery": "string",
	}
	className = "user"
	results = []types.M{}
	transformInQuery(inQueryObject, className, results)
	expect = types.M{
		"$in": types.S{},
	}
	if reflect.DeepEqual(expect, inQueryObject) == false {
		t.Error("expect:", expect, "result:", inQueryObject)
	}
	/**********************************************************/
	inQueryObject = types.M{
		"$inQuery": "string",
	}
	className = "user"
	results = []types.M{
		types.M{
			"objectId": "1001",
		},
		types.M{
			"key": "1002",
		},
	}
	transformInQuery(inQueryObject, className, results)
	expect = types.M{
		"$in": types.S{
			types.M{
				"__type":    "Pointer",
				"className": "user",
				"objectId":  "1001",
			},
		},
	}
	if reflect.DeepEqual(expect, inQueryObject) == false {
		t.Error("expect:", expect, "result:", inQueryObject)
	}
	/**********************************************************/
	inQueryObject = types.M{
		"$inQuery": "string",
	}
	className = "user"
	results = []types.M{
		types.M{
			"objectId": "1001",
		},
		types.M{
			"objectId": "1002",
		},
	}
	transformInQuery(inQueryObject, className, results)
	expect = types.M{
		"$in": types.S{
			types.M{
				"__type":    "Pointer",
				"className": "user",
				"objectId":  "1001",
			},
			types.M{
				"__type":    "Pointer",
				"className": "user",
				"objectId":  "1002",
			},
		},
	}
	if reflect.DeepEqual(expect, inQueryObject) == false {
		t.Error("expect:", expect, "result:", inQueryObject)
	}
	/**********************************************************/
	inQueryObject = types.M{
		"$inQuery": "string",
		"$in": types.S{
			types.M{
				"__type":    "Pointer",
				"className": "user",
				"objectId":  "1003",
			},
		},
	}
	className = "user"
	results = []types.M{
		types.M{
			"objectId": "1001",
		},
		types.M{
			"objectId": "1002",
		},
	}
	transformInQuery(inQueryObject, className, results)
	expect = types.M{
		"$in": types.S{
			types.M{
				"__type":    "Pointer",
				"className": "user",
				"objectId":  "1003",
			},
			types.M{
				"__type":    "Pointer",
				"className": "user",
				"objectId":  "1001",
			},
			types.M{
				"__type":    "Pointer",
				"className": "user",
				"objectId":  "1002",
			},
		},
	}
	if reflect.DeepEqual(expect, inQueryObject) == false {
		t.Error("expect:", expect, "result:", inQueryObject)
	}
}

func Test_transformNotInQuery(t *testing.T) {
	var notInQueryObject types.M
	var className string
	var results []types.M
	var expect types.M
	/**********************************************************/
	notInQueryObject = nil
	className = "user"
	results = nil
	transformNotInQuery(notInQueryObject, className, results)
	expect = nil
	if reflect.DeepEqual(expect, notInQueryObject) == false {
		t.Error("expect:", expect, "result:", notInQueryObject)
	}
	/**********************************************************/
	notInQueryObject = types.M{}
	className = "user"
	results = nil
	transformNotInQuery(notInQueryObject, className, results)
	expect = types.M{}
	if reflect.DeepEqual(expect, notInQueryObject) == false {
		t.Error("expect:", expect, "result:", notInQueryObject)
	}
	/**********************************************************/
	notInQueryObject = types.M{
		"$notInQuery": "string",
	}
	className = "user"
	results = nil
	transformNotInQuery(notInQueryObject, className, results)
	expect = types.M{
		"$nin": types.S{},
	}
	if reflect.DeepEqual(expect, notInQueryObject) == false {
		t.Error("expect:", expect, "result:", notInQueryObject)
	}
	/**********************************************************/
	notInQueryObject = types.M{
		"$notInQuery": "string",
	}
	className = "user"
	results = []types.M{}
	transformNotInQuery(notInQueryObject, className, results)
	expect = types.M{
		"$nin": types.S{},
	}
	if reflect.DeepEqual(expect, notInQueryObject) == false {
		t.Error("expect:", expect, "result:", notInQueryObject)
	}
	/**********************************************************/
	notInQueryObject = types.M{
		"$notInQuery": "string",
	}
	className = "user"
	results = []types.M{
		types.M{
			"objectId": "1001",
		},
		types.M{
			"key": "1002",
		},
	}
	transformNotInQuery(notInQueryObject, className, results)
	expect = types.M{
		"$nin": types.S{
			types.M{
				"__type":    "Pointer",
				"className": "user",
				"objectId":  "1001",
			},
		},
	}
	if reflect.DeepEqual(expect, notInQueryObject) == false {
		t.Error("expect:", expect, "result:", notInQueryObject)
	}
	/**********************************************************/
	notInQueryObject = types.M{
		"$notInQuery": "string",
	}
	className = "user"
	results = []types.M{
		types.M{
			"objectId": "1001",
		},
		types.M{
			"objectId": "1002",
		},
	}
	transformNotInQuery(notInQueryObject, className, results)
	expect = types.M{
		"$nin": types.S{
			types.M{
				"__type":    "Pointer",
				"className": "user",
				"objectId":  "1001",
			},
			types.M{
				"__type":    "Pointer",
				"className": "user",
				"objectId":  "1002",
			},
		},
	}
	if reflect.DeepEqual(expect, notInQueryObject) == false {
		t.Error("expect:", expect, "result:", notInQueryObject)
	}
	/**********************************************************/
	notInQueryObject = types.M{
		"$notInQuery": "string",
		"$nin": types.S{
			types.M{
				"__type":    "Pointer",
				"className": "user",
				"objectId":  "1003",
			},
		},
	}
	className = "user"
	results = []types.M{
		types.M{
			"objectId": "1001",
		},
		types.M{
			"objectId": "1002",
		},
	}
	transformNotInQuery(notInQueryObject, className, results)
	expect = types.M{
		"$nin": types.S{
			types.M{
				"__type":    "Pointer",
				"className": "user",
				"objectId":  "1003",
			},
			types.M{
				"__type":    "Pointer",
				"className": "user",
				"objectId":  "1001",
			},
			types.M{
				"__type":    "Pointer",
				"className": "user",
				"objectId":  "1002",
			},
		},
	}
	if reflect.DeepEqual(expect, notInQueryObject) == false {
		t.Error("expect:", expect, "result:", notInQueryObject)
	}
}
