package minio

func (d *MinioTests) TestList() {
	r := d.Require()
	versions := []string{"v1.0.0", "v1.1.0", "v1.2.0"}
	for _, version := range versions {
		r.NoError(d.storage.Save(module, version, mod, zip, info))
	}
	retVersions, err := d.storage.List(module)
	r.NoError(err)
	r.Equal(versions, retVersions)
}
