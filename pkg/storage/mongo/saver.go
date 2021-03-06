package mongo

import "github.com/gomods/athens/pkg/storage"

// Save stores a module in mongo storage.
func (s *ModuleStore) Save(module, version string, mod, zip, info []byte) error {
	m := &storage.Module{
		Module:  module,
		Version: version,
		Mod:     mod,
		Zip:     zip,
		Info:    info,
	}

	c := s.s.DB(s.d).C(s.c)
	return c.Insert(m)
}
