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

// ===== BEGIN of query set FlagQuerySet

// FlagQuerySet is an queryset type for Flag
type FlagQuerySet struct {
	db *gorm.DB
}

// NewFlagQuerySet constructs new FlagQuerySet
func NewFlagQuerySet(db *gorm.DB) FlagQuerySet {
	return FlagQuerySet{
		db: db.Model(&Flag{}),
	}
}

func (qs FlagQuerySet) w(db *gorm.DB) FlagQuerySet {
	return NewFlagQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) All(ret *[]Flag) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *Flag) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedAtEq(createdAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedAtGt(createdAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedAtGte(createdAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedAtLt(createdAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedAtLte(createdAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedAtNe(createdAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// CreatedByEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedByEq(createdBy string) FlagQuerySet {
	return qs.w(qs.db.Where("created_by = ?", createdBy))
}

// CreatedByIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedByIn(createdBy ...string) FlagQuerySet {
	if len(createdBy) == 0 {
		qs.db.AddError(errors.New("must at least pass one createdBy in CreatedByIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("created_by IN (?)", createdBy))
}

// CreatedByNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedByNe(createdBy string) FlagQuerySet {
	return qs.w(qs.db.Where("created_by != ?", createdBy))
}

// CreatedByNotIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) CreatedByNotIn(createdBy ...string) FlagQuerySet {
	if len(createdBy) == 0 {
		qs.db.AddError(errors.New("must at least pass one createdBy in CreatedByNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("created_by NOT IN (?)", createdBy))
}

// DataRecordsEnabledEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DataRecordsEnabledEq(dataRecordsEnabled bool) FlagQuerySet {
	return qs.w(qs.db.Where("data_records_enabled = ?", dataRecordsEnabled))
}

// DataRecordsEnabledIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DataRecordsEnabledIn(dataRecordsEnabled ...bool) FlagQuerySet {
	if len(dataRecordsEnabled) == 0 {
		qs.db.AddError(errors.New("must at least pass one dataRecordsEnabled in DataRecordsEnabledIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("data_records_enabled IN (?)", dataRecordsEnabled))
}

// DataRecordsEnabledNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DataRecordsEnabledNe(dataRecordsEnabled bool) FlagQuerySet {
	return qs.w(qs.db.Where("data_records_enabled != ?", dataRecordsEnabled))
}

// DataRecordsEnabledNotIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DataRecordsEnabledNotIn(dataRecordsEnabled ...bool) FlagQuerySet {
	if len(dataRecordsEnabled) == 0 {
		qs.db.AddError(errors.New("must at least pass one dataRecordsEnabled in DataRecordsEnabledNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("data_records_enabled NOT IN (?)", dataRecordsEnabled))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Flag) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) Delete() error {
	return qs.db.Delete(Flag{}).Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DeletedAtEq(deletedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DeletedAtGt(deletedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DeletedAtGte(deletedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DeletedAtIsNotNull() FlagQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DeletedAtIsNull() FlagQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DeletedAtLt(deletedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DeletedAtLte(deletedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DeletedAtNe(deletedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// DescriptionEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DescriptionEq(description string) FlagQuerySet {
	return qs.w(qs.db.Where("description = ?", description))
}

// DescriptionIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DescriptionIn(description ...string) FlagQuerySet {
	if len(description) == 0 {
		qs.db.AddError(errors.New("must at least pass one description in DescriptionIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("description IN (?)", description))
}

// DescriptionNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DescriptionNe(description string) FlagQuerySet {
	return qs.w(qs.db.Where("description != ?", description))
}

// DescriptionNotIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) DescriptionNotIn(description ...string) FlagQuerySet {
	if len(description) == 0 {
		qs.db.AddError(errors.New("must at least pass one description in DescriptionNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("description NOT IN (?)", description))
}

// EnabledEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) EnabledEq(enabled bool) FlagQuerySet {
	return qs.w(qs.db.Where("enabled = ?", enabled))
}

// EnabledIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) EnabledIn(enabled ...bool) FlagQuerySet {
	if len(enabled) == 0 {
		qs.db.AddError(errors.New("must at least pass one enabled in EnabledIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("enabled IN (?)", enabled))
}

// EnabledNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) EnabledNe(enabled bool) FlagQuerySet {
	return qs.w(qs.db.Where("enabled != ?", enabled))
}

// EnabledNotIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) EnabledNotIn(enabled ...bool) FlagQuerySet {
	if len(enabled) == 0 {
		qs.db.AddError(errors.New("must at least pass one enabled in EnabledNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("enabled NOT IN (?)", enabled))
}

// EntityTypeEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) EntityTypeEq(entityType string) FlagQuerySet {
	return qs.w(qs.db.Where("entity_type = ?", entityType))
}

// EntityTypeIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) EntityTypeIn(entityType ...string) FlagQuerySet {
	if len(entityType) == 0 {
		qs.db.AddError(errors.New("must at least pass one entityType in EntityTypeIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("entity_type IN (?)", entityType))
}

// EntityTypeNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) EntityTypeNe(entityType string) FlagQuerySet {
	return qs.w(qs.db.Where("entity_type != ?", entityType))
}

// EntityTypeNotIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) EntityTypeNotIn(entityType ...string) FlagQuerySet {
	if len(entityType) == 0 {
		qs.db.AddError(errors.New("must at least pass one entityType in EntityTypeNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("entity_type NOT IN (?)", entityType))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) GetUpdater() FlagUpdater {
	return NewFlagUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) IDEq(ID uint) FlagQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) IDGt(ID uint) FlagQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) IDGte(ID uint) FlagQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) IDIn(ID ...uint) FlagQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) IDLt(ID uint) FlagQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) IDLte(ID uint) FlagQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) IDNe(ID uint) FlagQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) IDNotIn(ID ...uint) FlagQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// KeyEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) KeyEq(key string) FlagQuerySet {
	return qs.w(qs.db.Where("key = ?", key))
}

// KeyIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) KeyIn(key ...string) FlagQuerySet {
	if len(key) == 0 {
		qs.db.AddError(errors.New("must at least pass one key in KeyIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("key IN (?)", key))
}

// KeyNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) KeyNe(key string) FlagQuerySet {
	return qs.w(qs.db.Where("key != ?", key))
}

// KeyNotIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) KeyNotIn(key ...string) FlagQuerySet {
	if len(key) == 0 {
		qs.db.AddError(errors.New("must at least pass one key in KeyNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("key NOT IN (?)", key))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) Limit(limit int) FlagQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) Offset(offset int) FlagQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs FlagQuerySet) One(ret *Flag) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) OrderAscByCreatedAt() FlagQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) OrderAscByDeletedAt() FlagQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) OrderAscByID() FlagQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscBySnapshotID is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) OrderAscBySnapshotID() FlagQuerySet {
	return qs.w(qs.db.Order("snapshot_id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) OrderAscByUpdatedAt() FlagQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) OrderDescByCreatedAt() FlagQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) OrderDescByDeletedAt() FlagQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) OrderDescByID() FlagQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescBySnapshotID is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) OrderDescBySnapshotID() FlagQuerySet {
	return qs.w(qs.db.Order("snapshot_id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) OrderDescByUpdatedAt() FlagQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetCreatedAt(createdAt time.Time) FlagUpdater {
	u.fields[string(FlagDBSchema.CreatedAt)] = createdAt
	return u
}

// SetCreatedBy is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetCreatedBy(createdBy string) FlagUpdater {
	u.fields[string(FlagDBSchema.CreatedBy)] = createdBy
	return u
}

// SetDataRecordsEnabled is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetDataRecordsEnabled(dataRecordsEnabled bool) FlagUpdater {
	u.fields[string(FlagDBSchema.DataRecordsEnabled)] = dataRecordsEnabled
	return u
}

// SetDeletedAt is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetDeletedAt(deletedAt *time.Time) FlagUpdater {
	u.fields[string(FlagDBSchema.DeletedAt)] = deletedAt
	return u
}

// SetDescription is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetDescription(description string) FlagUpdater {
	u.fields[string(FlagDBSchema.Description)] = description
	return u
}

// SetEnabled is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetEnabled(enabled bool) FlagUpdater {
	u.fields[string(FlagDBSchema.Enabled)] = enabled
	return u
}

// SetEntityType is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetEntityType(entityType string) FlagUpdater {
	u.fields[string(FlagDBSchema.EntityType)] = entityType
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetID(ID uint) FlagUpdater {
	u.fields[string(FlagDBSchema.ID)] = ID
	return u
}

// SetKey is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetKey(key string) FlagUpdater {
	u.fields[string(FlagDBSchema.Key)] = key
	return u
}

// SetSnapshotID is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetSnapshotID(snapshotID uint) FlagUpdater {
	u.fields[string(FlagDBSchema.SnapshotID)] = snapshotID
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetUpdatedAt(updatedAt time.Time) FlagUpdater {
	u.fields[string(FlagDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetUpdatedBy is an autogenerated method
// nolint: dupl
func (u FlagUpdater) SetUpdatedBy(updatedBy string) FlagUpdater {
	u.fields[string(FlagDBSchema.UpdatedBy)] = updatedBy
	return u
}

// SnapshotIDEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) SnapshotIDEq(snapshotID uint) FlagQuerySet {
	return qs.w(qs.db.Where("snapshot_id = ?", snapshotID))
}

// SnapshotIDGt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) SnapshotIDGt(snapshotID uint) FlagQuerySet {
	return qs.w(qs.db.Where("snapshot_id > ?", snapshotID))
}

// SnapshotIDGte is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) SnapshotIDGte(snapshotID uint) FlagQuerySet {
	return qs.w(qs.db.Where("snapshot_id >= ?", snapshotID))
}

// SnapshotIDIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) SnapshotIDIn(snapshotID ...uint) FlagQuerySet {
	if len(snapshotID) == 0 {
		qs.db.AddError(errors.New("must at least pass one snapshotID in SnapshotIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("snapshot_id IN (?)", snapshotID))
}

// SnapshotIDLt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) SnapshotIDLt(snapshotID uint) FlagQuerySet {
	return qs.w(qs.db.Where("snapshot_id < ?", snapshotID))
}

// SnapshotIDLte is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) SnapshotIDLte(snapshotID uint) FlagQuerySet {
	return qs.w(qs.db.Where("snapshot_id <= ?", snapshotID))
}

// SnapshotIDNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) SnapshotIDNe(snapshotID uint) FlagQuerySet {
	return qs.w(qs.db.Where("snapshot_id != ?", snapshotID))
}

// SnapshotIDNotIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) SnapshotIDNotIn(snapshotID ...uint) FlagQuerySet {
	if len(snapshotID) == 0 {
		qs.db.AddError(errors.New("must at least pass one snapshotID in SnapshotIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("snapshot_id NOT IN (?)", snapshotID))
}

// Update is an autogenerated method
// nolint: dupl
func (u FlagUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u FlagUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedAtEq(updatedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedAtGt(updatedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedAtGte(updatedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedAtLt(updatedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedAtLte(updatedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedAtNe(updatedAt time.Time) FlagQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// UpdatedByEq is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedByEq(updatedBy string) FlagQuerySet {
	return qs.w(qs.db.Where("updated_by = ?", updatedBy))
}

// UpdatedByIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedByIn(updatedBy ...string) FlagQuerySet {
	if len(updatedBy) == 0 {
		qs.db.AddError(errors.New("must at least pass one updatedBy in UpdatedByIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("updated_by IN (?)", updatedBy))
}

// UpdatedByNe is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedByNe(updatedBy string) FlagQuerySet {
	return qs.w(qs.db.Where("updated_by != ?", updatedBy))
}

// UpdatedByNotIn is an autogenerated method
// nolint: dupl
func (qs FlagQuerySet) UpdatedByNotIn(updatedBy ...string) FlagQuerySet {
	if len(updatedBy) == 0 {
		qs.db.AddError(errors.New("must at least pass one updatedBy in UpdatedByNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("updated_by NOT IN (?)", updatedBy))
}

// ===== END of query set FlagQuerySet

// ===== BEGIN of Flag modifiers

// FlagDBSchemaField describes database schema field. It requires for method 'Update'
type FlagDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f FlagDBSchemaField) String() string {
	return string(f)
}

// FlagDBSchema stores db field names of Flag
var FlagDBSchema = struct {
	ID                 FlagDBSchemaField
	CreatedAt          FlagDBSchemaField
	UpdatedAt          FlagDBSchemaField
	DeletedAt          FlagDBSchemaField
	Key                FlagDBSchemaField
	Description        FlagDBSchemaField
	CreatedBy          FlagDBSchemaField
	UpdatedBy          FlagDBSchemaField
	Enabled            FlagDBSchemaField
	SnapshotID         FlagDBSchemaField
	DataRecordsEnabled FlagDBSchemaField
	EntityType         FlagDBSchemaField
}{

	ID:                 FlagDBSchemaField("id"),
	CreatedAt:          FlagDBSchemaField("created_at"),
	UpdatedAt:          FlagDBSchemaField("updated_at"),
	DeletedAt:          FlagDBSchemaField("deleted_at"),
	Key:                FlagDBSchemaField("key"),
	Description:        FlagDBSchemaField("description"),
	CreatedBy:          FlagDBSchemaField("created_by"),
	UpdatedBy:          FlagDBSchemaField("updated_by"),
	Enabled:            FlagDBSchemaField("enabled"),
	SnapshotID:         FlagDBSchemaField("snapshot_id"),
	DataRecordsEnabled: FlagDBSchemaField("data_records_enabled"),
	EntityType:         FlagDBSchemaField("entity_type"),
}

// Update updates Flag fields by primary key
// nolint: dupl
func (o *Flag) Update(db *gorm.DB, fields ...FlagDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":                   o.ID,
		"created_at":           o.CreatedAt,
		"updated_at":           o.UpdatedAt,
		"deleted_at":           o.DeletedAt,
		"key":                  o.Key,
		"description":          o.Description,
		"created_by":           o.CreatedBy,
		"updated_by":           o.UpdatedBy,
		"enabled":              o.Enabled,
		"snapshot_id":          o.SnapshotID,
		"data_records_enabled": o.DataRecordsEnabled,
		"entity_type":          o.EntityType,
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

		return fmt.Errorf("can't update Flag %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// FlagUpdater is an Flag updates manager
type FlagUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewFlagUpdater creates new Flag updater
// nolint: dupl
func NewFlagUpdater(db *gorm.DB) FlagUpdater {
	return FlagUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Flag{}),
	}
}

// ===== END of Flag modifiers

// ===== END of all query sets
