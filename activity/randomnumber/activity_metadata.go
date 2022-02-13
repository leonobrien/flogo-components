package randomnumber

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "randomnumber",
  "type": "flogo:activity",
  "ref": "github.com/leonobrien/flogo-components/edit/master/activity/randomnumber",
  "version": "0.0.1",
  "title": "Random Number",
  "description": "Creates a random number between min and max",
  "author": "retgits",
  "homepage": "https://github.com/leonobrien/flogo-components/edit/master/activity/randomnumber",
  "inputs":[
    {
      "name": "min",
      "type": "integer"
    },
    {
      "name": "max",
      "type": "integer"
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "integer"
    }
  ]
}`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
