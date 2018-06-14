package mongo

import (
	"github.com/gobuffalo/buffalo"
)

func (m *MongoTests) TestList() {
	r := m.Require()
	versions := []string{"v1.0.0", "v1.1.0", "v1.2.0"}
	for _, version := range versions {
		m.storage.Save(&buffalo.DefaultContext{}, module, version, mod, zip, info)
	}
	retVersions, err := m.storage.List(module)
	r.NoError(err)
	r.Equal(versions, retVersions)
}
