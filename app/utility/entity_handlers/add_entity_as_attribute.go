package entity_handlers

import "GIG/app/models"

func AddEntityAsAttribute(entity models.Entity, attributeName string, attributeEntity models.Entity) (models.Entity, models.Entity, error) {
	entity, linkEntity, err := AddEntityAsLink(entity, attributeEntity)
	if err != nil {
		return entity, attributeEntity, err
	}
	entity = entity.SetAttribute(attributeName, models.Value{
		Type:     "objectId",
		RawValue: linkEntity.ID.Hex(),
	})

	return entity, linkEntity, nil
}