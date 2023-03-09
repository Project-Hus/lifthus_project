// Code generated by ent, DO NOT EDIT.

package user

import (
	"hus-auth/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// GoogleSub applies equality check predicate on the "google_sub" field. It's identical to GoogleSubEQ.
func GoogleSub(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGoogleSub, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailVerified applies equality check predicate on the "email_verified" field. It's identical to EmailVerifiedEQ.
func EmailVerified(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmailVerified, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// GivenName applies equality check predicate on the "given_name" field. It's identical to GivenNameEQ.
func GivenName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGivenName, v))
}

// FamilyName applies equality check predicate on the "family_name" field. It's identical to FamilyNameEQ.
func FamilyName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFamilyName, v))
}

// Birthdate applies equality check predicate on the "birthdate" field. It's identical to BirthdateEQ.
func Birthdate(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBirthdate, v))
}

// ProfilePictureURL applies equality check predicate on the "profile_picture_url" field. It's identical to ProfilePictureURLEQ.
func ProfilePictureURL(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldProfilePictureURL, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// ProviderEQ applies the EQ predicate on the "provider" field.
func ProviderEQ(v Provider) predicate.User {
	return predicate.User(sql.FieldEQ(FieldProvider, v))
}

// ProviderNEQ applies the NEQ predicate on the "provider" field.
func ProviderNEQ(v Provider) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldProvider, v))
}

// ProviderIn applies the In predicate on the "provider" field.
func ProviderIn(vs ...Provider) predicate.User {
	return predicate.User(sql.FieldIn(FieldProvider, vs...))
}

// ProviderNotIn applies the NotIn predicate on the "provider" field.
func ProviderNotIn(vs ...Provider) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldProvider, vs...))
}

// GoogleSubEQ applies the EQ predicate on the "google_sub" field.
func GoogleSubEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGoogleSub, v))
}

// GoogleSubNEQ applies the NEQ predicate on the "google_sub" field.
func GoogleSubNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldGoogleSub, v))
}

// GoogleSubIn applies the In predicate on the "google_sub" field.
func GoogleSubIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldGoogleSub, vs...))
}

// GoogleSubNotIn applies the NotIn predicate on the "google_sub" field.
func GoogleSubNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldGoogleSub, vs...))
}

// GoogleSubGT applies the GT predicate on the "google_sub" field.
func GoogleSubGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldGoogleSub, v))
}

// GoogleSubGTE applies the GTE predicate on the "google_sub" field.
func GoogleSubGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldGoogleSub, v))
}

// GoogleSubLT applies the LT predicate on the "google_sub" field.
func GoogleSubLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldGoogleSub, v))
}

// GoogleSubLTE applies the LTE predicate on the "google_sub" field.
func GoogleSubLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldGoogleSub, v))
}

// GoogleSubContains applies the Contains predicate on the "google_sub" field.
func GoogleSubContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldGoogleSub, v))
}

// GoogleSubHasPrefix applies the HasPrefix predicate on the "google_sub" field.
func GoogleSubHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldGoogleSub, v))
}

// GoogleSubHasSuffix applies the HasSuffix predicate on the "google_sub" field.
func GoogleSubHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldGoogleSub, v))
}

// GoogleSubEqualFold applies the EqualFold predicate on the "google_sub" field.
func GoogleSubEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldGoogleSub, v))
}

// GoogleSubContainsFold applies the ContainsFold predicate on the "google_sub" field.
func GoogleSubContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldGoogleSub, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// EmailVerifiedEQ applies the EQ predicate on the "email_verified" field.
func EmailVerifiedEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmailVerified, v))
}

// EmailVerifiedNEQ applies the NEQ predicate on the "email_verified" field.
func EmailVerifiedNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmailVerified, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// GivenNameEQ applies the EQ predicate on the "given_name" field.
func GivenNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGivenName, v))
}

// GivenNameNEQ applies the NEQ predicate on the "given_name" field.
func GivenNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldGivenName, v))
}

// GivenNameIn applies the In predicate on the "given_name" field.
func GivenNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldGivenName, vs...))
}

// GivenNameNotIn applies the NotIn predicate on the "given_name" field.
func GivenNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldGivenName, vs...))
}

// GivenNameGT applies the GT predicate on the "given_name" field.
func GivenNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldGivenName, v))
}

// GivenNameGTE applies the GTE predicate on the "given_name" field.
func GivenNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldGivenName, v))
}

// GivenNameLT applies the LT predicate on the "given_name" field.
func GivenNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldGivenName, v))
}

// GivenNameLTE applies the LTE predicate on the "given_name" field.
func GivenNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldGivenName, v))
}

// GivenNameContains applies the Contains predicate on the "given_name" field.
func GivenNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldGivenName, v))
}

// GivenNameHasPrefix applies the HasPrefix predicate on the "given_name" field.
func GivenNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldGivenName, v))
}

// GivenNameHasSuffix applies the HasSuffix predicate on the "given_name" field.
func GivenNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldGivenName, v))
}

// GivenNameEqualFold applies the EqualFold predicate on the "given_name" field.
func GivenNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldGivenName, v))
}

// GivenNameContainsFold applies the ContainsFold predicate on the "given_name" field.
func GivenNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldGivenName, v))
}

// FamilyNameEQ applies the EQ predicate on the "family_name" field.
func FamilyNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFamilyName, v))
}

// FamilyNameNEQ applies the NEQ predicate on the "family_name" field.
func FamilyNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldFamilyName, v))
}

// FamilyNameIn applies the In predicate on the "family_name" field.
func FamilyNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldFamilyName, vs...))
}

// FamilyNameNotIn applies the NotIn predicate on the "family_name" field.
func FamilyNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldFamilyName, vs...))
}

// FamilyNameGT applies the GT predicate on the "family_name" field.
func FamilyNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldFamilyName, v))
}

// FamilyNameGTE applies the GTE predicate on the "family_name" field.
func FamilyNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldFamilyName, v))
}

// FamilyNameLT applies the LT predicate on the "family_name" field.
func FamilyNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldFamilyName, v))
}

// FamilyNameLTE applies the LTE predicate on the "family_name" field.
func FamilyNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldFamilyName, v))
}

// FamilyNameContains applies the Contains predicate on the "family_name" field.
func FamilyNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldFamilyName, v))
}

// FamilyNameHasPrefix applies the HasPrefix predicate on the "family_name" field.
func FamilyNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldFamilyName, v))
}

// FamilyNameHasSuffix applies the HasSuffix predicate on the "family_name" field.
func FamilyNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldFamilyName, v))
}

// FamilyNameEqualFold applies the EqualFold predicate on the "family_name" field.
func FamilyNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldFamilyName, v))
}

// FamilyNameContainsFold applies the ContainsFold predicate on the "family_name" field.
func FamilyNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldFamilyName, v))
}

// BirthdateEQ applies the EQ predicate on the "birthdate" field.
func BirthdateEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldBirthdate, v))
}

// BirthdateNEQ applies the NEQ predicate on the "birthdate" field.
func BirthdateNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldBirthdate, v))
}

// BirthdateIn applies the In predicate on the "birthdate" field.
func BirthdateIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldBirthdate, vs...))
}

// BirthdateNotIn applies the NotIn predicate on the "birthdate" field.
func BirthdateNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldBirthdate, vs...))
}

// BirthdateGT applies the GT predicate on the "birthdate" field.
func BirthdateGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldBirthdate, v))
}

// BirthdateGTE applies the GTE predicate on the "birthdate" field.
func BirthdateGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldBirthdate, v))
}

// BirthdateLT applies the LT predicate on the "birthdate" field.
func BirthdateLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldBirthdate, v))
}

// BirthdateLTE applies the LTE predicate on the "birthdate" field.
func BirthdateLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldBirthdate, v))
}

// BirthdateIsNil applies the IsNil predicate on the "birthdate" field.
func BirthdateIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldBirthdate))
}

// BirthdateNotNil applies the NotNil predicate on the "birthdate" field.
func BirthdateNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldBirthdate))
}

// ProfilePictureURLEQ applies the EQ predicate on the "profile_picture_url" field.
func ProfilePictureURLEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldProfilePictureURL, v))
}

// ProfilePictureURLNEQ applies the NEQ predicate on the "profile_picture_url" field.
func ProfilePictureURLNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldProfilePictureURL, v))
}

// ProfilePictureURLIn applies the In predicate on the "profile_picture_url" field.
func ProfilePictureURLIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldProfilePictureURL, vs...))
}

// ProfilePictureURLNotIn applies the NotIn predicate on the "profile_picture_url" field.
func ProfilePictureURLNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldProfilePictureURL, vs...))
}

// ProfilePictureURLGT applies the GT predicate on the "profile_picture_url" field.
func ProfilePictureURLGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldProfilePictureURL, v))
}

// ProfilePictureURLGTE applies the GTE predicate on the "profile_picture_url" field.
func ProfilePictureURLGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldProfilePictureURL, v))
}

// ProfilePictureURLLT applies the LT predicate on the "profile_picture_url" field.
func ProfilePictureURLLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldProfilePictureURL, v))
}

// ProfilePictureURLLTE applies the LTE predicate on the "profile_picture_url" field.
func ProfilePictureURLLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldProfilePictureURL, v))
}

// ProfilePictureURLContains applies the Contains predicate on the "profile_picture_url" field.
func ProfilePictureURLContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldProfilePictureURL, v))
}

// ProfilePictureURLHasPrefix applies the HasPrefix predicate on the "profile_picture_url" field.
func ProfilePictureURLHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldProfilePictureURL, v))
}

// ProfilePictureURLHasSuffix applies the HasSuffix predicate on the "profile_picture_url" field.
func ProfilePictureURLHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldProfilePictureURL, v))
}

// ProfilePictureURLIsNil applies the IsNil predicate on the "profile_picture_url" field.
func ProfilePictureURLIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldProfilePictureURL))
}

// ProfilePictureURLNotNil applies the NotNil predicate on the "profile_picture_url" field.
func ProfilePictureURLNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldProfilePictureURL))
}

// ProfilePictureURLEqualFold applies the EqualFold predicate on the "profile_picture_url" field.
func ProfilePictureURLEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldProfilePictureURL, v))
}

// ProfilePictureURLContainsFold applies the ContainsFold predicate on the "profile_picture_url" field.
func ProfilePictureURLContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldProfilePictureURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasHusSessions applies the HasEdge predicate on the "hus_sessions" edge.
func HasHusSessions() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, HusSessionsTable, HusSessionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHusSessionsWith applies the HasEdge predicate on the "hus_sessions" edge with a given conditions (other predicates).
func HasHusSessionsWith(preds ...predicate.HusSession) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HusSessionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, HusSessionsTable, HusSessionsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
