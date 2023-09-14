package kzobjectstorage

func (c *defaultClient) SelectBucket(bucketName string) {
	c.bucketName = bucketName
}
