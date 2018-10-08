//go:generate goqueryset -in flag.go

package entity

import (
	"fmt"
	"strings"

	"github.com/checkr/flagr/pkg/util"
	"github.com/jinzhu/gorm"
)

// Flag is the unit of flags
// gen:qs
type Flag struct {
	gorm.Model

	Key         string `gorm:"type:varchar(64);unique_index:idx_flag_key"`
	Description string `sql:"type:text"`
	CreatedBy   string
	UpdatedBy   string
	Enabled     bool
	Segments    []Segment
	Variants    []Variant
	SnapshotID  uint `json:"-"`

	DataRecordsEnabled bool
	EntityType         string

	FlagEvaluation FlagEvaluation `gorm:"-" json:"-"`
}

// FlagEvaluation is a struct that holds the necessary info for evaluation
type FlagEvaluation struct {
	VariantsMap map[uint]*Variant
}

// Preload preloads the segments and variants into flags
func (f *Flag) Preload(db *gorm.DB) error {
	// preload Segments
	ss := []Segment{}
	segmentQuery := NewSegmentQuerySet(db)
	if err := segmentQuery.FlagIDEq(f.ID).OrderAscByRank().OrderAscByID().All(&ss); err != nil {
		return err
	}
	for i, s := range ss {
		if err := s.Preload(db); err != nil {
			return err
		}
		ss[i] = s
	}
	f.Segments = ss

	// preload Variants
	vs := []Variant{}
	variantQuery := NewVariantQuerySet(db)
	err := variantQuery.FlagIDEq(f.ID).OrderAscByID().All(&vs)
	if err != nil {
		return err
	}
	f.Variants = vs
	return nil
}

// PrepareEvaluation prepares the information for evaluation
func (f *Flag) PrepareEvaluation() error {
	f.FlagEvaluation = FlagEvaluation{
		VariantsMap: make(map[uint]*Variant),
	}
	for i := range f.Segments {
		if err := f.Segments[i].PrepareEvaluation(); err != nil {
			return err
		}
	}
	for i := range f.Variants {
		f.FlagEvaluation.VariantsMap[f.Variants[i].ID] = &f.Variants[i]
	}
	return nil
}

// DescriptionLike patches the autogenerated queryset for find flags that
// partially match the description
func (qs FlagQuerySet) DescriptionLike(description string) FlagQuerySet {
	return qs.w(qs.db.Where(
		"lower(description) like ?",
		fmt.Sprintf("%%%s%%", strings.ToLower(description)),
	))
}

// CreateFlagKey creates the key based on the given key
func CreateFlagKey(key string) (string, error) {
	if key == "" {
		key = util.NewSecureRandomKey()
	} else {
		ok, reason := util.IsSafeKey(key)
		if !ok {
			return "", fmt.Errorf("cannot create flag due to invalid key. reason: %s", reason)
		}
	}
	return key, nil
}

// CreateFlagEntityType creates the FlagEntityType if not exists
func CreateFlagEntityType(db *gorm.DB, key string) error {
	ok, reason := util.IsSafeKey(key)
	if !ok && key != "" {
		return fmt.Errorf("invalid DataRecordsEntityType. reason: %s", reason)
	}
	d := FlagEntityType{Key: key}
	if err := db.Where(d).FirstOrCreate(&d).Error; err != nil {
		return err
	}
	return nil
}
