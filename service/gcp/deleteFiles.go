package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"time"
)

func (g *gcpService) DeleteFile(ctx context.Context, Object string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	objectData := g.Client.Bucket(g.BUCKET).Object(Object)

	attrs, err := objectData.Attrs(ctx)
	if err != nil {
		return err
	}
	objectData = objectData.If(storage.Conditions{GenerationMatch: attrs.Generation})

	if err = objectData.Delete(ctx); err != nil {
		return err
	}

	return nil
}
