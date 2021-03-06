package controllers

import (
        "github.com/astaxie/beego"
        "github.com/lenfree/awsLaCapa/models"
)

type IAMController struct {
        beego.Controller
}

// @Title GetAll
// @Description Return all IAM users
// @Success 200 {object} models.IAMUsers
// @router / [get]
func (c *IAMController) GetAll() {
        IAMUsers, err := models.IAMUserList()
        if err != nil {
                c.Data["json"] = err.Error()
        } else {
                c.Data["json"] = IAMUsers
        }
        c.ServeJSON()
}
