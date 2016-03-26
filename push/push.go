package push

import "github.com/lfq7413/tomato/rest"
import "github.com/lfq7413/tomato/config"
import "github.com/lfq7413/tomato/utils"
import "github.com/lfq7413/tomato/orm"
import "strconv"

var adapter pushAdapter

func init() {
	a := config.TConfig.PushAdapter
	if a == "tomato" {
		adapter = newTomatoPush()
	} else {
		adapter = newTomatoPush()
	}
}

// SendPush ...
func SendPush(body map[string]interface{}, where map[string]interface{}, auth *rest.Auth) error {
	err := validatePushType(where, adapter.getValidPushTypes())
	if err != nil {
		return err
	}

	if body["expiration_time"] != nil {
		body["expiration_time"], err = getExpirationTime(body)
		if err != nil {
			return err
		}
	}

	var op map[string]interface{}
	var updateWhere map[string]interface{}
	data := utils.MapInterface(body["data"])
	if data != nil && data["badge"] != nil {
		badge := data["badge"]
		op = map[string]interface{}{}
		if utils.String(badge) == "Increment" {
			inc := map[string]interface{}{
				"badge": 1,
			}
			op["$inc"] = inc
		} else if v, ok := badge.(float64); ok {
			set := map[string]interface{}{
				"badge": v,
			}
			op["$set"] = set
		} else {
			// TODO badge 值不符合要求
			return nil
		}
		updateWhere = utils.CopyMap(where)
	}

	if op != nil && updateWhere != nil {
		badgeQuery := rest.NewQuery(auth, "_Installation", updateWhere, map[string]interface{}{})
		badgeQuery.BuildRestWhere()
		restWhere := utils.CopyMap(badgeQuery.Where)
		and := utils.SliceInterface(restWhere["$and"])
		if and == nil {
			restWhere["$and"] = []interface{}{badgeQuery.Where}
		}
		tp := map[string]interface{}{
			"deviceType": "ios",
		}
		and = append(and, tp)
		restWhere["$and"] = and
		orm.UpdateAll("_Installation", restWhere, op, map[string]interface{}{})
	}

	response := rest.Find(auth, "_Installation", where, map[string]interface{}{})
	if utils.HasResults(response) == false {
		return nil
	}
	results := utils.SliceInterface(response["results"])

	if data != nil && data["badge"] != nil && utils.String(data["badge"]) == "Increment" {
		badgeInstallationsMap := map[string]interface{}{}
		for _, v := range results {
			installation := utils.MapInterface(v)
			var badge string
			if v, ok := installation["badge"].(float64); ok {
				badge = strconv.Itoa(int(v))
			} else {
				continue
			}
			if utils.String(installation["deviceType"]) != "ios" {
				badge = "unsupported"
			}
			installations := []interface{}{}
			if badgeInstallationsMap[badge] != nil {
				installations = append(installations, utils.SliceInterface(badgeInstallationsMap[badge])...)
			}
			installations = append(installations, installation)
			badgeInstallationsMap[badge] = installations
		}

		for k, v := range badgeInstallationsMap {
			payload := utils.CopyMap(body)
			paydata := utils.MapInterface(payload["data"])
			if k == "unsupported" {
				delete(paydata, "badge")
			} else {
				paydata["badge"], _ = strconv.Atoi(k)
			}
			adapter.send(payload, utils.SliceInterface(v))
		}
		return nil
	}
	adapter.send(body, results)

	return nil
}

func validatePushType(where map[string]interface{}, validPushTypes []string) error {
	return nil
}

func getExpirationTime(body map[string]interface{}) (string, error) {
	return "", nil
}

type pushAdapter interface {
	send(data map[string]interface{}, installations []interface{})
	getValidPushTypes() []string
}
