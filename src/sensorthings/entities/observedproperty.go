package entities

import (
	"encoding/json"
	"errors"
	gostErrors "github.com/geodan/gost/src/errors"
)

// ObservedProperty in SensorThings represents the physical phenomenon being observed by the Sensor. An ObserveProperty is
// linked to a Datatream which can only have one ObserveProperty
type ObservedProperty struct {
	BaseEntity
	Description    string        `json:"description,omitempty"`
	Name           string        `json:"name,omitempty"`
	Definition     string        `json:"definition,omitempty"`
	NavDatastreams string        `json:"Datastreams@iot.navigationLink,omitempty"`
	Datastreams    []*Datastream `json:"Datastreams,omitempty"`
}

// GetEntityType returns the EntityType for ObservedProperty
func (o ObservedProperty) GetEntityType() EntityType {
	return EntityTypeObservedProperty
}

// GetPropertyNames returns the available properties for a ObservedProperty
func (o *ObservedProperty) GetPropertyNames() []string {
	return []string{"id", "description", "name", "definition"}
}

// ParseEntity tries to parse the given json byte array into the current entity
func (o *ObservedProperty) ParseEntity(data []byte) error {
	op := &o
	err := json.Unmarshal(data, op)
	if err != nil {
		return gostErrors.NewBadRequestError(errors.New("Unable to parse ObservedProperty"))
	}

	return nil
}

// ContainsMandatoryParams checks if all mandatory params for ObservedProperty are available before posting.
func (o *ObservedProperty) ContainsMandatoryParams() (bool, []error) {
	err := []error{}
	CheckMandatoryParam(&err, o.Name, o.GetEntityType(), "name")
	CheckMandatoryParam(&err, o.Definition, o.GetEntityType(), "definition")
	CheckMandatoryParam(&err, o.Description, o.GetEntityType(), "description")

	if len(err) != 0 {
		return false, err
	}

	return true, nil
}

// SetLinks sets the entity specific navigation links if needed
func (o *ObservedProperty) SetLinks(externalURL string) {
	o.NavSelf = CreateEntitySelfLink(externalURL, EntityLinkObservedPropertys.ToString(), o.ID)
	o.NavDatastreams = CreateEntityLink(o.Datastreams == nil, externalURL, EntityLinkObservedPropertys.ToString(), EntityLinkDatastreams.ToString(), o.ID)
}

// GetSupportedEncoding returns the supported encoding tye for this entity
func (o ObservedProperty) GetSupportedEncoding() map[int]EncodingType {
	return map[int]EncodingType{}
}
