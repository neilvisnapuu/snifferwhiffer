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
    c.Validation.Required(myName).Message("Your name is required!")
    c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(App.Index)
    }

    demomgo.BlatherMe(myName)

    return c.Render(myName)
}