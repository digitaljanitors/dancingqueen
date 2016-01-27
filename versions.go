package main

import (
	"math/rand"

	drand "github.com/dgryski/go-discreterand"
)

type TargetedVersion interface {
	GetNextVersion() string
}

type SingleVersion struct {
	Version string
}

func (v *SingleVersion) GetNextVersion() string {
	return v.Version
}

type MultipleVersions struct {
	versions   []string
	aliasTable drand.AliasTable
}

func (v *MultipleVersions) GetNextVersion() string {
	return v.versions[v.aliasTable.Next()]
}

func NewMultipleVersions(wv []WeightedVersion) *MultipleVersions {
	n := len(wv)

	v := make([]string, n)
	p := make([]float64, n)

	for i := 0; i < n; i++ {
		v[i] = wv[i].Version
		p[i] = wv[i].Weight
	}

	return &MultipleVersions{
		versions:   v,
		aliasTable: drand.NewAlias(p, rand.NewSource(69)),
	}
}

type WeightedVersion struct {
	Weight  float64
	Version string
}
