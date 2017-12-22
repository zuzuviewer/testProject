package handler

import (
	"ecologyServer/utils/mongo"
	"encoding/json"
	"fmt"
	"github.com/henrylee2cn/faygo"
	"io/ioutil"
)

var SaveResult = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	var r = mongo.Result{}
	//ctx.R.ParseForm()
	body, err := ioutil.ReadAll(ctx.R.Body)
	if err != nil {
		fmt.Printf("read body err,%v\n", err)
		return ctx.JSON(412, faygo.Map{"error": err.Error()}, true)
	}
	println("json bosy:", string(body))
	if err = json.Unmarshal(body, &r); err != nil {
		fmt.Printf("unmarshal err,%v\n", err)
		return ctx.JSON(412, faygo.Map{"error": err.Error()}, true)
	}
	if err = mongo.SaveExperimentResult(r); err != nil {
		fmt.Printf("insert to mongo db err,%v\n", err)
		return ctx.JSON(412, faygo.Map{"error": err.Error()}, true)
	}
	fmt.Printf("start to reply client,date time is %s\n", r.ExperimentDateTime)
	return ctx.JSON(200, faygo.Map{"experiment_datetime": r.ExperimentDateTime}, true)
	//return err
})
