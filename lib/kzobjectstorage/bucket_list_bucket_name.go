package kzobjectstorage

import "context"

func (c *defaultBucket) ListBucket() ([]string, error) {
	buckets, err := c.minioClient.ListBuckets(context.TODO())
	if err != nil {
		return nil, err
	}

	newBuckets := make([]string, 0)
	for _, b := range buckets {
		newBuckets = append(newBuckets, b.Name)
	}

	return newBuckets, nil
}
