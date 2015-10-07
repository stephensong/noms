// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __mainPackageInFile_types_CachedRef = __mainPackageInFile_types_Ref()

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func __mainPackageInFile_types_Ref() ref.Ref {
	p := types.PackageDef{
		NamedTypes: types.MapOfStringToTypeRefDef{

			"Geoposition": types.MakeStructTypeRef("Geoposition",
				[]types.Field{
					types.Field{"Latitude", types.MakePrimitiveTypeRef(types.Float32Kind), false},
					types.Field{"Longitude", types.MakePrimitiveTypeRef(types.Float32Kind), false},
				},
				types.Choices{},
			),
			"Georectangle": types.MakeStructTypeRef("Georectangle",
				[]types.Field{
					types.Field{"TopLeft", types.MakeTypeRef("Geoposition", ref.Ref{}), false},
					types.Field{"BottomRight", types.MakeTypeRef("Geoposition", ref.Ref{}), false},
				},
				types.Choices{},
			),
			"Incident": types.MakeStructTypeRef("Incident",
				[]types.Field{
					types.Field{"Category", types.MakePrimitiveTypeRef(types.StringKind), false},
					types.Field{"Description", types.MakePrimitiveTypeRef(types.StringKind), false},
					types.Field{"Address", types.MakePrimitiveTypeRef(types.StringKind), false},
					types.Field{"Date", types.MakePrimitiveTypeRef(types.StringKind), false},
					types.Field{"Geoposition", types.MakeTypeRef("Geoposition", ref.Ref{}), false},
				},
				types.Choices{},
			),
			"SQuadTree": types.MakeStructTypeRef("SQuadTree",
				[]types.Field{
					types.Field{"Nodes", types.MakeCompoundTypeRef("", types.ListKind, types.MakeTypeRef("Incident", ref.Ref{})), false},
					types.Field{"Tiles", types.MakeCompoundTypeRef("", types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeTypeRef("SQuadTree", ref.Ref{})), false},
					types.Field{"Depth", types.MakePrimitiveTypeRef(types.UInt8Kind), false},
					types.Field{"NumDescendents", types.MakePrimitiveTypeRef(types.UInt32Kind), false},
					types.Field{"Path", types.MakePrimitiveTypeRef(types.StringKind), false},
					types.Field{"Georectangle", types.MakeTypeRef("Georectangle", ref.Ref{}), false},
				},
				types.Choices{},
			),
		},
	}.New()
	return types.RegisterPackage(&p)
}

// ListOfIncident

type ListOfIncident struct {
	l types.List
}

func NewListOfIncident() ListOfIncident {
	return ListOfIncident{types.NewList()}
}

type ListOfIncidentDef []IncidentDef

func (def ListOfIncidentDef) New() ListOfIncident {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = d.New().NomsValue()
	}
	return ListOfIncident{types.NewList(l...)}
}

func (l ListOfIncident) Def() ListOfIncidentDef {
	d := make([]IncidentDef, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = IncidentFromVal(l.l.Get(i)).Def()
	}
	return d
}

func ListOfIncidentFromVal(val types.Value) ListOfIncident {
	// TODO: Validate here
	return ListOfIncident{val.(types.List)}
}

func (l ListOfIncident) NomsValue() types.Value {
	return l.l
}

func (l ListOfIncident) Equals(other types.Value) bool {
	if other, ok := other.(ListOfIncident); ok {
		return l.l.Equals(other.l)
	}
	return false
}

func (l ListOfIncident) Ref() ref.Ref {
	return l.l.Ref()
}

func (l ListOfIncident) Chunks() []types.Future {
	return l.l.Chunks()
}

// A Noms Value that describes ListOfIncident.
var __typeRefForListOfIncident = types.MakeCompoundTypeRef("", types.ListKind, types.MakeTypeRef("Incident", __mainPackageInFile_types_CachedRef))

func (m ListOfIncident) TypeRef() types.TypeRef {
	return __typeRefForListOfIncident
}

func init() {
	types.RegisterFromValFunction(__typeRefForListOfIncident, func(v types.Value) types.NomsValue {
		return ListOfIncidentFromVal(v)
	})
}

func (l ListOfIncident) Len() uint64 {
	return l.l.Len()
}

func (l ListOfIncident) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfIncident) Get(i uint64) Incident {
	return IncidentFromVal(l.l.Get(i))
}

func (l ListOfIncident) Slice(idx uint64, end uint64) ListOfIncident {
	return ListOfIncident{l.l.Slice(idx, end)}
}

func (l ListOfIncident) Set(i uint64, val Incident) ListOfIncident {
	return ListOfIncident{l.l.Set(i, val.NomsValue())}
}

func (l ListOfIncident) Append(v ...Incident) ListOfIncident {
	return ListOfIncident{l.l.Append(l.fromElemSlice(v)...)}
}

func (l ListOfIncident) Insert(idx uint64, v ...Incident) ListOfIncident {
	return ListOfIncident{l.l.Insert(idx, l.fromElemSlice(v)...)}
}

func (l ListOfIncident) Remove(idx uint64, end uint64) ListOfIncident {
	return ListOfIncident{l.l.Remove(idx, end)}
}

func (l ListOfIncident) RemoveAt(idx uint64) ListOfIncident {
	return ListOfIncident{(l.l.RemoveAt(idx))}
}

func (l ListOfIncident) fromElemSlice(p []Incident) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v.NomsValue()
	}
	return r
}

type ListOfIncidentIterCallback func(v Incident, i uint64) (stop bool)

func (l ListOfIncident) Iter(cb ListOfIncidentIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(IncidentFromVal(v), i)
	})
}

type ListOfIncidentIterAllCallback func(v Incident, i uint64)

func (l ListOfIncident) IterAll(cb ListOfIncidentIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(IncidentFromVal(v), i)
	})
}

type ListOfIncidentFilterCallback func(v Incident, i uint64) (keep bool)

func (l ListOfIncident) Filter(cb ListOfIncidentFilterCallback) ListOfIncident {
	nl := NewListOfIncident()
	l.IterAll(func(v Incident, i uint64) {
		if cb(v, i) {
			nl = nl.Append(v)
		}
	})
	return nl
}

// Incident

type Incident struct {
	m types.Map
}

func NewIncident() Incident {
	return Incident{types.NewMap(
		types.NewString("$name"), types.NewString("Incident"),
		types.NewString("$type"), types.MakeTypeRef("Incident", __mainPackageInFile_types_CachedRef),
		types.NewString("Category"), types.NewString(""),
		types.NewString("Description"), types.NewString(""),
		types.NewString("Address"), types.NewString(""),
		types.NewString("Date"), types.NewString(""),
		types.NewString("Geoposition"), NewGeoposition().NomsValue(),
	)}
}

type IncidentDef struct {
	Category    string
	Description string
	Address     string
	Date        string
	Geoposition GeopositionDef
}

func (def IncidentDef) New() Incident {
	return Incident{
		types.NewMap(
			types.NewString("$name"), types.NewString("Incident"),
			types.NewString("$type"), types.MakeTypeRef("Incident", __mainPackageInFile_types_CachedRef),
			types.NewString("Category"), types.NewString(def.Category),
			types.NewString("Description"), types.NewString(def.Description),
			types.NewString("Address"), types.NewString(def.Address),
			types.NewString("Date"), types.NewString(def.Date),
			types.NewString("Geoposition"), def.Geoposition.New().NomsValue(),
		)}
}

func (s Incident) Def() (d IncidentDef) {
	d.Category = s.m.Get(types.NewString("Category")).(types.String).String()
	d.Description = s.m.Get(types.NewString("Description")).(types.String).String()
	d.Address = s.m.Get(types.NewString("Address")).(types.String).String()
	d.Date = s.m.Get(types.NewString("Date")).(types.String).String()
	d.Geoposition = GeopositionFromVal(s.m.Get(types.NewString("Geoposition"))).Def()
	return
}

var __typeRefForIncident = types.MakeTypeRef("Incident", __mainPackageInFile_types_CachedRef)

func (m Incident) TypeRef() types.TypeRef {
	return __typeRefForIncident
}

func init() {
	types.RegisterFromValFunction(__typeRefForIncident, func(v types.Value) types.NomsValue {
		return IncidentFromVal(v)
	})
}

func IncidentFromVal(val types.Value) Incident {
	// TODO: Validate here
	return Incident{val.(types.Map)}
}

func (s Incident) NomsValue() types.Value {
	return s.m
}

func (s Incident) Equals(other types.Value) bool {
	if other, ok := other.(Incident); ok {
		return s.m.Equals(other.m)
	}
	return false
}

func (s Incident) Ref() ref.Ref {
	return s.m.Ref()
}

func (s Incident) Chunks() []types.Future {
	return s.m.Chunks()
}

func (s Incident) Category() string {
	return s.m.Get(types.NewString("Category")).(types.String).String()
}

func (s Incident) SetCategory(val string) Incident {
	return Incident{s.m.Set(types.NewString("Category"), types.NewString(val))}
}

func (s Incident) Description() string {
	return s.m.Get(types.NewString("Description")).(types.String).String()
}

func (s Incident) SetDescription(val string) Incident {
	return Incident{s.m.Set(types.NewString("Description"), types.NewString(val))}
}

func (s Incident) Address() string {
	return s.m.Get(types.NewString("Address")).(types.String).String()
}

func (s Incident) SetAddress(val string) Incident {
	return Incident{s.m.Set(types.NewString("Address"), types.NewString(val))}
}

func (s Incident) Date() string {
	return s.m.Get(types.NewString("Date")).(types.String).String()
}

func (s Incident) SetDate(val string) Incident {
	return Incident{s.m.Set(types.NewString("Date"), types.NewString(val))}
}

func (s Incident) Geoposition() Geoposition {
	return GeopositionFromVal(s.m.Get(types.NewString("Geoposition")))
}

func (s Incident) SetGeoposition(val Geoposition) Incident {
	return Incident{s.m.Set(types.NewString("Geoposition"), val.NomsValue())}
}

// Geoposition

type Geoposition struct {
	m types.Map
}

func NewGeoposition() Geoposition {
	return Geoposition{types.NewMap(
		types.NewString("$name"), types.NewString("Geoposition"),
		types.NewString("$type"), types.MakeTypeRef("Geoposition", __mainPackageInFile_types_CachedRef),
		types.NewString("Latitude"), types.Float32(0),
		types.NewString("Longitude"), types.Float32(0),
	)}
}

type GeopositionDef struct {
	Latitude  float32
	Longitude float32
}

func (def GeopositionDef) New() Geoposition {
	return Geoposition{
		types.NewMap(
			types.NewString("$name"), types.NewString("Geoposition"),
			types.NewString("$type"), types.MakeTypeRef("Geoposition", __mainPackageInFile_types_CachedRef),
			types.NewString("Latitude"), types.Float32(def.Latitude),
			types.NewString("Longitude"), types.Float32(def.Longitude),
		)}
}

func (s Geoposition) Def() (d GeopositionDef) {
	d.Latitude = float32(s.m.Get(types.NewString("Latitude")).(types.Float32))
	d.Longitude = float32(s.m.Get(types.NewString("Longitude")).(types.Float32))
	return
}

var __typeRefForGeoposition = types.MakeTypeRef("Geoposition", __mainPackageInFile_types_CachedRef)

func (m Geoposition) TypeRef() types.TypeRef {
	return __typeRefForGeoposition
}

func init() {
	types.RegisterFromValFunction(__typeRefForGeoposition, func(v types.Value) types.NomsValue {
		return GeopositionFromVal(v)
	})
}

func GeopositionFromVal(val types.Value) Geoposition {
	// TODO: Validate here
	return Geoposition{val.(types.Map)}
}

func (s Geoposition) NomsValue() types.Value {
	return s.m
}

func (s Geoposition) Equals(other types.Value) bool {
	if other, ok := other.(Geoposition); ok {
		return s.m.Equals(other.m)
	}
	return false
}

func (s Geoposition) Ref() ref.Ref {
	return s.m.Ref()
}

func (s Geoposition) Chunks() []types.Future {
	return s.m.Chunks()
}

func (s Geoposition) Latitude() float32 {
	return float32(s.m.Get(types.NewString("Latitude")).(types.Float32))
}

func (s Geoposition) SetLatitude(val float32) Geoposition {
	return Geoposition{s.m.Set(types.NewString("Latitude"), types.Float32(val))}
}

func (s Geoposition) Longitude() float32 {
	return float32(s.m.Get(types.NewString("Longitude")).(types.Float32))
}

func (s Geoposition) SetLongitude(val float32) Geoposition {
	return Geoposition{s.m.Set(types.NewString("Longitude"), types.Float32(val))}
}

// Georectangle

type Georectangle struct {
	m types.Map
}

func NewGeorectangle() Georectangle {
	return Georectangle{types.NewMap(
		types.NewString("$name"), types.NewString("Georectangle"),
		types.NewString("$type"), types.MakeTypeRef("Georectangle", __mainPackageInFile_types_CachedRef),
		types.NewString("TopLeft"), NewGeoposition().NomsValue(),
		types.NewString("BottomRight"), NewGeoposition().NomsValue(),
	)}
}

type GeorectangleDef struct {
	TopLeft     GeopositionDef
	BottomRight GeopositionDef
}

func (def GeorectangleDef) New() Georectangle {
	return Georectangle{
		types.NewMap(
			types.NewString("$name"), types.NewString("Georectangle"),
			types.NewString("$type"), types.MakeTypeRef("Georectangle", __mainPackageInFile_types_CachedRef),
			types.NewString("TopLeft"), def.TopLeft.New().NomsValue(),
			types.NewString("BottomRight"), def.BottomRight.New().NomsValue(),
		)}
}

func (s Georectangle) Def() (d GeorectangleDef) {
	d.TopLeft = GeopositionFromVal(s.m.Get(types.NewString("TopLeft"))).Def()
	d.BottomRight = GeopositionFromVal(s.m.Get(types.NewString("BottomRight"))).Def()
	return
}

var __typeRefForGeorectangle = types.MakeTypeRef("Georectangle", __mainPackageInFile_types_CachedRef)

func (m Georectangle) TypeRef() types.TypeRef {
	return __typeRefForGeorectangle
}

func init() {
	types.RegisterFromValFunction(__typeRefForGeorectangle, func(v types.Value) types.NomsValue {
		return GeorectangleFromVal(v)
	})
}

func GeorectangleFromVal(val types.Value) Georectangle {
	// TODO: Validate here
	return Georectangle{val.(types.Map)}
}

func (s Georectangle) NomsValue() types.Value {
	return s.m
}

func (s Georectangle) Equals(other types.Value) bool {
	if other, ok := other.(Georectangle); ok {
		return s.m.Equals(other.m)
	}
	return false
}

func (s Georectangle) Ref() ref.Ref {
	return s.m.Ref()
}

func (s Georectangle) Chunks() []types.Future {
	return s.m.Chunks()
}

func (s Georectangle) TopLeft() Geoposition {
	return GeopositionFromVal(s.m.Get(types.NewString("TopLeft")))
}

func (s Georectangle) SetTopLeft(val Geoposition) Georectangle {
	return Georectangle{s.m.Set(types.NewString("TopLeft"), val.NomsValue())}
}

func (s Georectangle) BottomRight() Geoposition {
	return GeopositionFromVal(s.m.Get(types.NewString("BottomRight")))
}

func (s Georectangle) SetBottomRight(val Geoposition) Georectangle {
	return Georectangle{s.m.Set(types.NewString("BottomRight"), val.NomsValue())}
}

// SQuadTree

type SQuadTree struct {
	m types.Map
}

func NewSQuadTree() SQuadTree {
	return SQuadTree{types.NewMap(
		types.NewString("$name"), types.NewString("SQuadTree"),
		types.NewString("$type"), types.MakeTypeRef("SQuadTree", __mainPackageInFile_types_CachedRef),
		types.NewString("Nodes"), types.NewList(),
		types.NewString("Tiles"), types.NewMap(),
		types.NewString("Depth"), types.UInt8(0),
		types.NewString("NumDescendents"), types.UInt32(0),
		types.NewString("Path"), types.NewString(""),
		types.NewString("Georectangle"), NewGeorectangle().NomsValue(),
	)}
}

type SQuadTreeDef struct {
	Nodes          ListOfIncidentDef
	Tiles          MapOfStringToSQuadTreeDef
	Depth          uint8
	NumDescendents uint32
	Path           string
	Georectangle   GeorectangleDef
}

func (def SQuadTreeDef) New() SQuadTree {
	return SQuadTree{
		types.NewMap(
			types.NewString("$name"), types.NewString("SQuadTree"),
			types.NewString("$type"), types.MakeTypeRef("SQuadTree", __mainPackageInFile_types_CachedRef),
			types.NewString("Nodes"), def.Nodes.New().NomsValue(),
			types.NewString("Tiles"), def.Tiles.New().NomsValue(),
			types.NewString("Depth"), types.UInt8(def.Depth),
			types.NewString("NumDescendents"), types.UInt32(def.NumDescendents),
			types.NewString("Path"), types.NewString(def.Path),
			types.NewString("Georectangle"), def.Georectangle.New().NomsValue(),
		)}
}

func (s SQuadTree) Def() (d SQuadTreeDef) {
	d.Nodes = ListOfIncidentFromVal(s.m.Get(types.NewString("Nodes"))).Def()
	d.Tiles = MapOfStringToSQuadTreeFromVal(s.m.Get(types.NewString("Tiles"))).Def()
	d.Depth = uint8(s.m.Get(types.NewString("Depth")).(types.UInt8))
	d.NumDescendents = uint32(s.m.Get(types.NewString("NumDescendents")).(types.UInt32))
	d.Path = s.m.Get(types.NewString("Path")).(types.String).String()
	d.Georectangle = GeorectangleFromVal(s.m.Get(types.NewString("Georectangle"))).Def()
	return
}

var __typeRefForSQuadTree = types.MakeTypeRef("SQuadTree", __mainPackageInFile_types_CachedRef)

func (m SQuadTree) TypeRef() types.TypeRef {
	return __typeRefForSQuadTree
}

func init() {
	types.RegisterFromValFunction(__typeRefForSQuadTree, func(v types.Value) types.NomsValue {
		return SQuadTreeFromVal(v)
	})
}

func SQuadTreeFromVal(val types.Value) SQuadTree {
	// TODO: Validate here
	return SQuadTree{val.(types.Map)}
}

func (s SQuadTree) NomsValue() types.Value {
	return s.m
}

func (s SQuadTree) Equals(other types.Value) bool {
	if other, ok := other.(SQuadTree); ok {
		return s.m.Equals(other.m)
	}
	return false
}

func (s SQuadTree) Ref() ref.Ref {
	return s.m.Ref()
}

func (s SQuadTree) Chunks() []types.Future {
	return s.m.Chunks()
}

func (s SQuadTree) Nodes() ListOfIncident {
	return ListOfIncidentFromVal(s.m.Get(types.NewString("Nodes")))
}

func (s SQuadTree) SetNodes(val ListOfIncident) SQuadTree {
	return SQuadTree{s.m.Set(types.NewString("Nodes"), val.NomsValue())}
}

func (s SQuadTree) Tiles() MapOfStringToSQuadTree {
	return MapOfStringToSQuadTreeFromVal(s.m.Get(types.NewString("Tiles")))
}

func (s SQuadTree) SetTiles(val MapOfStringToSQuadTree) SQuadTree {
	return SQuadTree{s.m.Set(types.NewString("Tiles"), val.NomsValue())}
}

func (s SQuadTree) Depth() uint8 {
	return uint8(s.m.Get(types.NewString("Depth")).(types.UInt8))
}

func (s SQuadTree) SetDepth(val uint8) SQuadTree {
	return SQuadTree{s.m.Set(types.NewString("Depth"), types.UInt8(val))}
}

func (s SQuadTree) NumDescendents() uint32 {
	return uint32(s.m.Get(types.NewString("NumDescendents")).(types.UInt32))
}

func (s SQuadTree) SetNumDescendents(val uint32) SQuadTree {
	return SQuadTree{s.m.Set(types.NewString("NumDescendents"), types.UInt32(val))}
}

func (s SQuadTree) Path() string {
	return s.m.Get(types.NewString("Path")).(types.String).String()
}

func (s SQuadTree) SetPath(val string) SQuadTree {
	return SQuadTree{s.m.Set(types.NewString("Path"), types.NewString(val))}
}

func (s SQuadTree) Georectangle() Georectangle {
	return GeorectangleFromVal(s.m.Get(types.NewString("Georectangle")))
}

func (s SQuadTree) SetGeorectangle(val Georectangle) SQuadTree {
	return SQuadTree{s.m.Set(types.NewString("Georectangle"), val.NomsValue())}
}

// MapOfStringToSQuadTree

type MapOfStringToSQuadTree struct {
	m types.Map
}

func NewMapOfStringToSQuadTree() MapOfStringToSQuadTree {
	return MapOfStringToSQuadTree{types.NewMap()}
}

type MapOfStringToSQuadTreeDef map[string]SQuadTreeDef

func (def MapOfStringToSQuadTreeDef) New() MapOfStringToSQuadTree {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), v.New().NomsValue())
	}
	return MapOfStringToSQuadTree{types.NewMap(kv...)}
}

func (m MapOfStringToSQuadTree) Def() MapOfStringToSQuadTreeDef {
	def := make(map[string]SQuadTreeDef)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = SQuadTreeFromVal(v).Def()
		return false
	})
	return def
}

func MapOfStringToSQuadTreeFromVal(p types.Value) MapOfStringToSQuadTree {
	// TODO: Validate here
	return MapOfStringToSQuadTree{p.(types.Map)}
}

func (m MapOfStringToSQuadTree) NomsValue() types.Value {
	return m.m
}

func (m MapOfStringToSQuadTree) Equals(other types.Value) bool {
	if other, ok := other.(MapOfStringToSQuadTree); ok {
		return m.m.Equals(other.m)
	}
	return false
}

func (m MapOfStringToSQuadTree) Ref() ref.Ref {
	return m.m.Ref()
}

func (m MapOfStringToSQuadTree) Chunks() []types.Future {
	return m.m.Chunks()
}

// A Noms Value that describes MapOfStringToSQuadTree.
var __typeRefForMapOfStringToSQuadTree = types.MakeCompoundTypeRef("", types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeTypeRef("SQuadTree", __mainPackageInFile_types_CachedRef))

func (m MapOfStringToSQuadTree) TypeRef() types.TypeRef {
	return __typeRefForMapOfStringToSQuadTree
}

func init() {
	types.RegisterFromValFunction(__typeRefForMapOfStringToSQuadTree, func(v types.Value) types.NomsValue {
		return MapOfStringToSQuadTreeFromVal(v)
	})
}

func (m MapOfStringToSQuadTree) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToSQuadTree) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToSQuadTree) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToSQuadTree) Get(p string) SQuadTree {
	return SQuadTreeFromVal(m.m.Get(types.NewString(p)))
}

func (m MapOfStringToSQuadTree) MaybeGet(p string) (SQuadTree, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return NewSQuadTree(), false
	}
	return SQuadTreeFromVal(v), ok
}

func (m MapOfStringToSQuadTree) Set(k string, v SQuadTree) MapOfStringToSQuadTree {
	return MapOfStringToSQuadTree{m.m.Set(types.NewString(k), v.NomsValue())}
}

// TODO: Implement SetM?

func (m MapOfStringToSQuadTree) Remove(p string) MapOfStringToSQuadTree {
	return MapOfStringToSQuadTree{m.m.Remove(types.NewString(p))}
}

type MapOfStringToSQuadTreeIterCallback func(k string, v SQuadTree) (stop bool)

func (m MapOfStringToSQuadTree) Iter(cb MapOfStringToSQuadTreeIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), SQuadTreeFromVal(v))
	})
}

type MapOfStringToSQuadTreeIterAllCallback func(k string, v SQuadTree)

func (m MapOfStringToSQuadTree) IterAll(cb MapOfStringToSQuadTreeIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), SQuadTreeFromVal(v))
	})
}

type MapOfStringToSQuadTreeFilterCallback func(k string, v SQuadTree) (keep bool)

func (m MapOfStringToSQuadTree) Filter(cb MapOfStringToSQuadTreeFilterCallback) MapOfStringToSQuadTree {
	nm := NewMapOfStringToSQuadTree()
	m.IterAll(func(k string, v SQuadTree) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}
