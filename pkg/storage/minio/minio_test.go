package minio

import (
	"context"
	"io/ioutil"

	"github.com/gobuffalo/buffalo"
)

func (d *MinioTests) TestGetSaveListRoundTrip() {
	c := &buffalo.DefaultContext{
		Context: context.Background(),
	}
	r := d.Require()
	r.NoError(d.storage.Save(c, module, version, mod, zip, info))
	listedVersions, err := d.storage.List(module)
	r.NoError(err)
	r.Equal(1, len(listedVersions))
	retVersion := listedVersions[0]
	r.Equal(version, retVersion)
	gotten, err := d.storage.Get(module, version)
	r.NoError(err)
	defer gotten.Zip.Close()
	// TODO: test the time
	r.Equal(gotten.Mod, mod)
	zipContent, err := ioutil.ReadAll(gotten.Zip)
	r.NoError(err)
	r.Equal(zipContent, zip)
	r.Equal(gotten.Info, info)
}
