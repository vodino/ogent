// Code generated by entc, DO NOT EDIT.

package ogent

import "ariga.io/ogent/example/pets/ent"

func NewCategoryCreate(e *ent.Category) *CategoryCreate {
	if e == nil {
		return nil
	}
	return &CategoryCreate{
		ID:   e.ID,
		Name: e.Name,
	}
}

func NewCategoryCreates(es []*ent.Category) []CategoryCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]CategoryCreate, len(es))
	for i, e := range es {
		r[i] = NewCategoryCreate(e).Elem()
	}
	return r
}

func (c *CategoryCreate) Elem() CategoryCreate {
	if c != nil {
		return CategoryCreate{}
	}
	return *c
}

func NewCategoryList(e *ent.Category) *CategoryList {
	if e == nil {
		return nil
	}
	return &CategoryList{
		ID:   e.ID,
		Name: e.Name,
	}
}

func NewCategoryLists(es []*ent.Category) []CategoryList {
	if len(es) == 0 {
		return nil
	}
	r := make([]CategoryList, len(es))
	for i, e := range es {
		r[i] = NewCategoryList(e).Elem()
	}
	return r
}

func (c *CategoryList) Elem() CategoryList {
	if c != nil {
		return CategoryList{}
	}
	return *c
}

func NewCategoryRead(e *ent.Category) *CategoryRead {
	if e == nil {
		return nil
	}
	return &CategoryRead{
		ID:   e.ID,
		Name: e.Name,
	}
}

func NewCategoryReads(es []*ent.Category) []CategoryRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]CategoryRead, len(es))
	for i, e := range es {
		r[i] = NewCategoryRead(e).Elem()
	}
	return r
}

func (c *CategoryRead) Elem() CategoryRead {
	if c != nil {
		return CategoryRead{}
	}
	return *c
}

func NewCategoryUpdate(e *ent.Category) *CategoryUpdate {
	if e == nil {
		return nil
	}
	return &CategoryUpdate{
		ID:   e.ID,
		Name: e.Name,
	}
}

func NewCategoryUpdates(es []*ent.Category) []CategoryUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]CategoryUpdate, len(es))
	for i, e := range es {
		r[i] = NewCategoryUpdate(e).Elem()
	}
	return r
}

func (c *CategoryUpdate) Elem() CategoryUpdate {
	if c != nil {
		return CategoryUpdate{}
	}
	return *c
}

func NewCategoryPetsList(e *ent.Pet) *CategoryPetsList {
	if e == nil {
		return nil
	}
	return &CategoryPetsList{
		ID:       e.ID,
		Name:     e.Name,
		Weight:   NewOptInt(e.Weight),
		Birthday: NewOptTime(e.Birthday),
	}
}

func NewCategoryPetsLists(es []*ent.Pet) []CategoryPetsList {
	if len(es) == 0 {
		return nil
	}
	r := make([]CategoryPetsList, len(es))
	for i, e := range es {
		r[i] = NewCategoryPetsList(e).Elem()
	}
	return r
}

func (pe *CategoryPetsList) Elem() CategoryPetsList {
	if pe != nil {
		return CategoryPetsList{}
	}
	return *pe
}

func NewPetCreate(e *ent.Pet) *PetCreate {
	if e == nil {
		return nil
	}
	return &PetCreate{
		ID:         e.ID,
		Name:       e.Name,
		Weight:     NewOptInt(e.Weight),
		Birthday:   NewOptTime(e.Birthday),
		Categories: NewPetCreateCategoriesSlice(e.Edges.Categories),
		Owner:      NewPetCreateOwner(e.Edges.Owner).Elem(),
	}
}

func NewPetCreates(es []*ent.Pet) []PetCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetCreate, len(es))
	for i, e := range es {
		r[i] = NewPetCreate(e).Elem()
	}
	return r
}

func (pe *PetCreate) Elem() PetCreate {
	if pe != nil {
		return PetCreate{}
	}
	return *pe
}

func NewPetCreateCategories(e *ent.Category) *PetCreateCategories {
	if e == nil {
		return nil
	}
	return &PetCreateCategories{
		ID:   e.ID,
		Name: e.Name,
	}
}

func NewPetCreateCategoriesSlice(es []*ent.Category) []PetCreateCategories {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetCreateCategories, len(es))
	for i, e := range es {
		r[i] = NewPetCreateCategories(e).Elem()
	}
	return r
}

func (c *PetCreateCategories) Elem() PetCreateCategories {
	if c != nil {
		return PetCreateCategories{}
	}
	return *c
}

func NewPetCreateOwner(e *ent.User) *PetCreateOwner {
	if e == nil {
		return nil
	}
	return &PetCreateOwner{
		ID:   e.ID,
		Name: e.Name,
		Age:  e.Age,
	}
}

func NewPetCreateOwners(es []*ent.User) []PetCreateOwner {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetCreateOwner, len(es))
	for i, e := range es {
		r[i] = NewPetCreateOwner(e).Elem()
	}
	return r
}

func (u *PetCreateOwner) Elem() PetCreateOwner {
	if u != nil {
		return PetCreateOwner{}
	}
	return *u
}

func NewPetList(e *ent.Pet) *PetList {
	if e == nil {
		return nil
	}
	return &PetList{
		ID:       e.ID,
		Name:     e.Name,
		Weight:   NewOptInt(e.Weight),
		Birthday: NewOptTime(e.Birthday),
	}
}

func NewPetLists(es []*ent.Pet) []PetList {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetList, len(es))
	for i, e := range es {
		r[i] = NewPetList(e).Elem()
	}
	return r
}

func (pe *PetList) Elem() PetList {
	if pe != nil {
		return PetList{}
	}
	return *pe
}

func NewPetRead(e *ent.Pet) *PetRead {
	if e == nil {
		return nil
	}
	return &PetRead{
		ID:       e.ID,
		Name:     e.Name,
		Weight:   NewOptInt(e.Weight),
		Birthday: NewOptTime(e.Birthday),
	}
}

func NewPetReads(es []*ent.Pet) []PetRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetRead, len(es))
	for i, e := range es {
		r[i] = NewPetRead(e).Elem()
	}
	return r
}

func (pe *PetRead) Elem() PetRead {
	if pe != nil {
		return PetRead{}
	}
	return *pe
}

func NewPetUpdate(e *ent.Pet) *PetUpdate {
	if e == nil {
		return nil
	}
	return &PetUpdate{
		ID:       e.ID,
		Name:     e.Name,
		Weight:   NewOptInt(e.Weight),
		Birthday: NewOptTime(e.Birthday),
	}
}

func NewPetUpdates(es []*ent.Pet) []PetUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetUpdate, len(es))
	for i, e := range es {
		r[i] = NewPetUpdate(e).Elem()
	}
	return r
}

func (pe *PetUpdate) Elem() PetUpdate {
	if pe != nil {
		return PetUpdate{}
	}
	return *pe
}

func NewPetCategoriesList(e *ent.Category) *PetCategoriesList {
	if e == nil {
		return nil
	}
	return &PetCategoriesList{
		ID:   e.ID,
		Name: e.Name,
	}
}

func NewPetCategoriesLists(es []*ent.Category) []PetCategoriesList {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetCategoriesList, len(es))
	for i, e := range es {
		r[i] = NewPetCategoriesList(e).Elem()
	}
	return r
}

func (c *PetCategoriesList) Elem() PetCategoriesList {
	if c != nil {
		return PetCategoriesList{}
	}
	return *c
}

func NewPetFriendsList(e *ent.Pet) *PetFriendsList {
	if e == nil {
		return nil
	}
	return &PetFriendsList{
		ID:       e.ID,
		Name:     e.Name,
		Weight:   NewOptInt(e.Weight),
		Birthday: NewOptTime(e.Birthday),
	}
}

func NewPetFriendsLists(es []*ent.Pet) []PetFriendsList {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetFriendsList, len(es))
	for i, e := range es {
		r[i] = NewPetFriendsList(e).Elem()
	}
	return r
}

func (pe *PetFriendsList) Elem() PetFriendsList {
	if pe != nil {
		return PetFriendsList{}
	}
	return *pe
}

func NewPetOwnerRead(e *ent.User) *PetOwnerRead {
	if e == nil {
		return nil
	}
	return &PetOwnerRead{
		ID:   e.ID,
		Name: e.Name,
		Age:  e.Age,
	}
}

func NewPetOwnerReads(es []*ent.User) []PetOwnerRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]PetOwnerRead, len(es))
	for i, e := range es {
		r[i] = NewPetOwnerRead(e).Elem()
	}
	return r
}

func (u *PetOwnerRead) Elem() PetOwnerRead {
	if u != nil {
		return PetOwnerRead{}
	}
	return *u
}

func NewUserCreate(e *ent.User) *UserCreate {
	if e == nil {
		return nil
	}
	return &UserCreate{
		ID:   e.ID,
		Name: e.Name,
		Age:  e.Age,
	}
}

func NewUserCreates(es []*ent.User) []UserCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserCreate, len(es))
	for i, e := range es {
		r[i] = NewUserCreate(e).Elem()
	}
	return r
}

func (u *UserCreate) Elem() UserCreate {
	if u != nil {
		return UserCreate{}
	}
	return *u
}

func NewUserList(e *ent.User) *UserList {
	if e == nil {
		return nil
	}
	return &UserList{
		ID:   e.ID,
		Name: e.Name,
		Age:  e.Age,
	}
}

func NewUserLists(es []*ent.User) []UserList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserList, len(es))
	for i, e := range es {
		r[i] = NewUserList(e).Elem()
	}
	return r
}

func (u *UserList) Elem() UserList {
	if u != nil {
		return UserList{}
	}
	return *u
}

func NewUserRead(e *ent.User) *UserRead {
	if e == nil {
		return nil
	}
	return &UserRead{
		ID:   e.ID,
		Name: e.Name,
		Age:  e.Age,
	}
}

func NewUserReads(es []*ent.User) []UserRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserRead, len(es))
	for i, e := range es {
		r[i] = NewUserRead(e).Elem()
	}
	return r
}

func (u *UserRead) Elem() UserRead {
	if u != nil {
		return UserRead{}
	}
	return *u
}

func NewUserUpdate(e *ent.User) *UserUpdate {
	if e == nil {
		return nil
	}
	return &UserUpdate{
		ID:   e.ID,
		Name: e.Name,
		Age:  e.Age,
	}
}

func NewUserUpdates(es []*ent.User) []UserUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserUpdate, len(es))
	for i, e := range es {
		r[i] = NewUserUpdate(e).Elem()
	}
	return r
}

func (u *UserUpdate) Elem() UserUpdate {
	if u != nil {
		return UserUpdate{}
	}
	return *u
}

func NewUserPetsList(e *ent.Pet) *UserPetsList {
	if e == nil {
		return nil
	}
	return &UserPetsList{
		ID:       e.ID,
		Name:     e.Name,
		Weight:   NewOptInt(e.Weight),
		Birthday: NewOptTime(e.Birthday),
	}
}

func NewUserPetsLists(es []*ent.Pet) []UserPetsList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserPetsList, len(es))
	for i, e := range es {
		r[i] = NewUserPetsList(e).Elem()
	}
	return r
}

func (pe *UserPetsList) Elem() UserPetsList {
	if pe != nil {
		return UserPetsList{}
	}
	return *pe
}
