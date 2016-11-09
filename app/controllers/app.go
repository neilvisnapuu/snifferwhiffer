package controllers

import "github.com/revel/revel"
import "github.com/neilvisnapuu/demomgo"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}



func (c App) Hello(myName string) revel.Result {
    c.Validation.Required(myName).Message("You have to tell us something!")
    c.Validation.MinSize(myName, 3).Message("Sorry but 3 chars doesn't cut it!")

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(App.Index)
    }

    demomgo.BlatherMe(myName)

    return c.Render(myName)
}



func (c App) Archive(myShit string) revel.Result {
    foo := demomgo.FullReport("")
    return c.Render(foo)
}