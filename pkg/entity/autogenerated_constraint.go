// Code generated by go-queryset. DO NOT EDIT.
package entity

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// notest
// ===== BEGIN of all query sets

// ===== BEGIN of query set ConstraintQuerySet

// ConstraintQuerySet is an queryset type for Constraint
type ConstraintQuerySet struct {
	db *gorm.DB
}

// NewConstraintQuerySet constructs new ConstraintQuerySet
func NewConstraintQuerySet(db *gorm.DB) ConstraintQuerySet {
	return ConstraintQuerySet{
		db: db.Model(&Constraint{}),
	}
}

func (qs ConstraintQuerySet) w(db *gorm.DB) ConstraintQuerySet {
	return NewConstraintQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) All(ret *[]Constraint) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *Constraint) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) CreatedAtEq(createdAt time.Time) ConstraintQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) CreatedAtGt(createdAt time.Time) ConstraintQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) CreatedAtGte(createdAt time.Time) ConstraintQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) CreatedAtLt(createdAt time.Time) ConstraintQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) CreatedAtLte(createdAt time.Time) ConstraintQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) CreatedAtNe(createdAt time.Time) ConstraintQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) Delete() error {
	return qs.db.Delete(Constraint{}).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Constraint) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) DeletedAtEq(deletedAt time.Time) ConstraintQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) DeletedAtGt(deletedAt time.Time) ConstraintQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) DeletedAtGte(deletedAt time.Time) ConstraintQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) DeletedAtIsNotNull() ConstraintQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) DeletedAtIsNull() ConstraintQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) DeletedAtLt(deletedAt time.Time) ConstraintQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) DeletedAtLte(deletedAt time.Time) ConstraintQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) DeletedAtNe(deletedAt time.Time) ConstraintQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) GetUpdater() ConstraintUpdater {
	return NewConstraintUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) IDEq(ID uint) ConstraintQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) IDGt(ID uint) ConstraintQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) IDGte(ID uint) ConstraintQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) IDIn(ID ...uint) ConstraintQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) IDLt(ID uint) ConstraintQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) IDLte(ID uint) ConstraintQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) IDNe(ID uint) ConstraintQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) IDNotIn(ID ...uint) ConstraintQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) Limit(limit int) ConstraintQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs ConstraintQuerySet) One(ret *Constraint) error {
	return qs.db.First(ret).Error
}

// OperatorEq is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) OperatorEq(operator string) ConstraintQuerySet {
	return qs.w(qs.db.Where("operator = ?", operator))
}

// OperatorIn is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) OperatorIn(operator ...string) ConstraintQuerySet {
	if len(operator) == 0 {
		qs.db.AddError(errors.New("must at least pass one operator in OperatorIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("operator IN (?)", operator))
}

// OperatorNe is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) OperatorNe(operator string) ConstraintQuerySet {
	return qs.w(qs.db.Where("operator != ?", operator))
}

// OperatorNotIn is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) OperatorNotIn(operator ...string) ConstraintQuerySet {
	if len(operator) == 0 {
		qs.db.AddError(errors.New("must at least pass one operator in OperatorNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("operator NOT IN (?)", operator))
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) OrderAscByCreatedAt() ConstraintQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) OrderAscByDeletedAt() ConstraintQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) OrderAscByID() ConstraintQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscBySegmentID is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) OrderAscBySegmentID() ConstraintQuerySet {
	return qs.w(qs.db.Order("segment_id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) OrderAscByUpdatedAt() ConstraintQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) OrderDescByCreatedAt() ConstraintQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) OrderDescByDeletedAt() ConstraintQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) OrderDescByID() ConstraintQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescBySegmentID is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) OrderDescBySegmentID() ConstraintQuerySet {
	return qs.w(qs.db.Order("segment_id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) OrderDescByUpdatedAt() ConstraintQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// PropertyEq is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) PropertyEq(property string) ConstraintQuerySet {
	return qs.w(qs.db.Where("property = ?", property))
}

// PropertyIn is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) PropertyIn(property ...string) ConstraintQuerySet {
	if len(property) == 0 {
		qs.db.AddError(errors.New("must at least pass one property in PropertyIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("property IN (?)", property))
}

// PropertyNe is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) PropertyNe(property string) ConstraintQuerySet {
	return qs.w(qs.db.Where("property != ?", property))
}

// PropertyNotIn is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) PropertyNotIn(property ...string) ConstraintQuerySet {
	if len(property) == 0 {
		qs.db.AddError(errors.New("must at least pass one property in PropertyNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("property NOT IN (?)", property))
}

// SegmentIDEq is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) SegmentIDEq(segmentID uint) ConstraintQuerySet {
	return qs.w(qs.db.Where("segment_id = ?", segmentID))
}

// SegmentIDGt is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) SegmentIDGt(segmentID uint) ConstraintQuerySet {
	return qs.w(qs.db.Where("segment_id > ?", segmentID))
}

// SegmentIDGte is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) SegmentIDGte(segmentID uint) ConstraintQuerySet {
	return qs.w(qs.db.Where("segment_id >= ?", segmentID))
}

// SegmentIDIn is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) SegmentIDIn(segmentID ...uint) ConstraintQuerySet {
	if len(segmentID) == 0 {
		qs.db.AddError(errors.New("must at least pass one segmentID in SegmentIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("segment_id IN (?)", segmentID))
}

// SegmentIDLt is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) SegmentIDLt(segmentID uint) ConstraintQuerySet {
	return qs.w(qs.db.Where("segment_id < ?", segmentID))
}

// SegmentIDLte is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) SegmentIDLte(segmentID uint) ConstraintQuerySet {
	return qs.w(qs.db.Where("segment_id <= ?", segmentID))
}

// SegmentIDNe is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) SegmentIDNe(segmentID uint) ConstraintQuerySet {
	return qs.w(qs.db.Where("segment_id != ?", segmentID))
}

// SegmentIDNotIn is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) SegmentIDNotIn(segmentID ...uint) ConstraintQuerySet {
	if len(segmentID) == 0 {
		qs.db.AddError(errors.New("must at least pass one segmentID in SegmentIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("segment_id NOT IN (?)", segmentID))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u ConstraintUpdater) SetCreatedAt(createdAt time.Time) ConstraintUpdater {
	u.fields[string(ConstraintDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDeletedAt is an autogenerated method
// nolint: dupl
func (u ConstraintUpdater) SetDeletedAt(deletedAt *time.Time) ConstraintUpdater {
	u.fields[string(ConstraintDBSchema.DeletedAt)] = deletedAt
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u ConstraintUpdater) SetID(ID uint) ConstraintUpdater {
	u.fields[string(ConstraintDBSchema.ID)] = ID
	return u
}

// SetOperator is an autogenerated method
// nolint: dupl
func (u ConstraintUpdater) SetOperator(operator string) ConstraintUpdater {
	u.fields[string(ConstraintDBSchema.Operator)] = operator
	return u
}

// SetProperty is an autogenerated method
// nolint: dupl
func (u ConstraintUpdater) SetProperty(property string) ConstraintUpdater {
	u.fields[string(ConstraintDBSchema.Property)] = property
	return u
}

// SetSegmentID is an autogenerated method
// nolint: dupl
func (u ConstraintUpdater) SetSegmentID(segmentID uint) ConstraintUpdater {
	u.fields[string(ConstraintDBSchema.SegmentID)] = segmentID
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u ConstraintUpdater) SetUpdatedAt(updatedAt time.Time) ConstraintUpdater {
	u.fields[string(ConstraintDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetValue is an autogenerated method
// nolint: dupl
func (u ConstraintUpdater) SetValue(value string) ConstraintUpdater {
	u.fields[string(ConstraintDBSchema.Value)] = value
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u ConstraintUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u ConstraintUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) UpdatedAtEq(updatedAt time.Time) ConstraintQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) UpdatedAtGt(updatedAt time.Time) ConstraintQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) UpdatedAtGte(updatedAt time.Time) ConstraintQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) UpdatedAtLt(updatedAt time.Time) ConstraintQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) UpdatedAtLte(updatedAt time.Time) ConstraintQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) UpdatedAtNe(updatedAt time.Time) ConstraintQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// ValueEq is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) ValueEq(value string) ConstraintQuerySet {
	return qs.w(qs.db.Where("value = ?", value))
}

// ValueIn is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) ValueIn(value ...string) ConstraintQuerySet {
	if len(value) == 0 {
		qs.db.AddError(errors.New("must at least pass one value in ValueIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("value IN (?)", value))
}

// ValueNe is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) ValueNe(value string) ConstraintQuerySet {
	return qs.w(qs.db.Where("value != ?", value))
}

// ValueNotIn is an autogenerated method
// nolint: dupl
func (qs ConstraintQuerySet) ValueNotIn(value ...string) ConstraintQuerySet {
	if len(value) == 0 {
		qs.db.AddError(errors.New("must at least pass one value in ValueNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("value NOT IN (?)", value))
}

// ===== END of query set ConstraintQuerySet

// ===== BEGIN of Constraint modifiers

// ConstraintDBSchemaField describes database schema field. It requires for method 'Update'
type ConstraintDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f ConstraintDBSchemaField) String() string {
	return string(f)
}

// ConstraintDBSchema stores db field names of Constraint
var ConstraintDBSchema = struct {
	ID        ConstraintDBSchemaField
	CreatedAt ConstraintDBSchemaField
	UpdatedAt ConstraintDBSchemaField
	DeletedAt ConstraintDBSchemaField
	SegmentID ConstraintDBSchemaField
	Property  ConstraintDBSchemaField
	Operator  ConstraintDBSchemaField
	Value     ConstraintDBSchemaField
}{

	ID:        ConstraintDBSchemaField("id"),
	CreatedAt: ConstraintDBSchemaField("created_at"),
	UpdatedAt: ConstraintDBSchemaField("updated_at"),
	DeletedAt: ConstraintDBSchemaField("deleted_at"),
	SegmentID: ConstraintDBSchemaField("segment_id"),
	Property:  ConstraintDBSchemaField("property"),
	Operator:  ConstraintDBSchemaField("operator"),
	Value:     ConstraintDBSchemaField("value"),
}

// Update updates Constraint fields by primary key
// nolint: dupl
func (o *Constraint) Update(db *gorm.DB, fields ...ConstraintDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":         o.ID,
		"created_at": o.CreatedAt,
		"updated_at": o.UpdatedAt,
		"deleted_at": o.DeletedAt,
		"segment_id": o.SegmentID,
		"property":   o.Property,
		"operator":   o.Operator,
		"value":      o.Value,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update Constraint %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// ConstraintUpdater is an Constraint updates manager
type ConstraintUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewConstraintUpdater creates new Constraint updater
// nolint: dupl
func NewConstraintUpdater(db *gorm.DB) ConstraintUpdater {
	return ConstraintUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Constraint{}),
	}
}

// ===== END of Constraint modifiers

// ===== END of all query sets
