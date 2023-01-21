package tx

import (
	"context"
	"fmt"
)

func (t *Tx) WithTx(ctx context.Context, fn func(tx *Tx) error) error {
	tx, err := t.Client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	q := NewTx(t.Client)
	if err := fn(q); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
