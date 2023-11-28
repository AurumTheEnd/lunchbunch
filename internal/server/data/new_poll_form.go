package data

import (
	"github.com/mozillazg/go-unidecode"
	"gitlab.fi.muni.cz/xhrdlic3/lunchbunch/internal/models"
	"net/http"
	"sort"
	"strconv"
)

type NewPollFormDataToClient struct {
	SnapshotID     uint       `validate:"required,number"`
	Checkboxes     []Checkbox `validate:"required"`
	ShouldRemember bool       `validate:"required,boolean"`
}

type Checkbox struct {
	ID      uint
	Name    string
	Checked bool
}

type NewPollFormDataToServer struct {
	SnapshotID     uint
	Checked        []uint
	ShouldRemember bool
}

func FromRequest(req *http.Request) (*NewPollFormDataToServer, error) {
	var result = NewPollFormDataToServer{}

	var parseErr = req.ParseForm()
	if parseErr != nil {
		return nil, parseErr
	}
	var formValues = req.Form

	var snapshotId, convertErr = strconv.ParseUint(formValues.Get("snapshot_id"), 10, 0)
	if convertErr != nil {
		return nil, convertErr
	}

	result.SnapshotID = uint(snapshotId)
	result.ShouldRemember = formValues.Has("should_remember")

	for _, checkedString := range formValues["checked"] {
		var checkedId, convertErr2 = strconv.ParseUint(checkedString, 10, 0)
		if convertErr2 != nil {
			return nil, convertErr2
		}
		result.Checked = append(result.Checked, uint(checkedId))
	}

	return &result, nil
}

func FromSnapshot(snapshot *models.RestaurantSnapshot) (formData NewPollFormDataToClient) {
	formData.ShouldRemember = false
	formData.SnapshotID = snapshot.ID

	for _, restaurant := range snapshot.Restaurants {
		var oneCheckbox = Checkbox{
			ID:      restaurant.ID,
			Name:    restaurant.Name,
			Checked: false,
		}
		formData.Checkboxes = append(formData.Checkboxes, oneCheckbox)
	}

	sort.Slice(formData.Checkboxes, func(i, j int) bool {
		var first = formData.Checkboxes[i].Name
		if len(first) == 0 {
			return true
		}
		var second = formData.Checkboxes[j].Name
		if len(second) == 0 {
			return false
		}

		return unidecode.Unidecode(first) < unidecode.Unidecode(second)
	})

	return
}
