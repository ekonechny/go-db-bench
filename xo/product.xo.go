package xo

// Code generated by xo. DO NOT EDIT.

import (
	"context"

	"github.com/lib/pq"

	"github.com/google/uuid"
)

// Product represents a row from 'public.products'.
type Product struct {
	ID          uuid.UUID      `json:"id"`          // id
	Name        string         `json:"name"`        // name
	Description string         `json:"description"` // description
	Categories  pq.StringArray `json:"categories"`  // categories
	Price       float64        `json:"price"`       // price
	Features    pq.StringArray `json:"features"`    // features
	Color       string         `json:"color"`       // color
	Material    string         `json:"material"`    // material
	Upc         string         `json:"upc"`         // upc
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [Product] exists in the database.
func (p *Product) Exists() bool {
	return p._exists
}

// Deleted returns true when the [Product] has been marked for deletion
// from the database.
func (p *Product) Deleted() bool {
	return p._deleted
}

// Insert inserts the [Product] to the database.
func (p *Product) Insert(ctx context.Context, db DB) error {
	switch {
	case p._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case p._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.products (` +
		`id, name, description, categories, price, features, color, material, upc` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`)`
	// run
	logf(sqlstr, p.ID, p.Name, p.Description, p.Categories, p.Price, p.Features, p.Color, p.Material, p.Upc)
	if _, err := db.ExecContext(ctx, sqlstr, p.ID, p.Name, p.Description, p.Categories, p.Price, p.Features, p.Color, p.Material, p.Upc); err != nil {
		return logerror(err)
	}
	// set exists
	p._exists = true
	return nil
}

// Update updates a [Product] in the database.
func (p *Product) Update(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case p._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.products SET ` +
		`name = $1, description = $2, categories = $3, price = $4, features = $5, color = $6, material = $7, upc = $8 ` +
		`WHERE id = $9`
	// run
	logf(sqlstr, p.Name, p.Description, p.Categories, p.Price, p.Features, p.Color, p.Material, p.Upc, p.ID)
	if _, err := db.ExecContext(ctx, sqlstr, p.Name, p.Description, p.Categories, p.Price, p.Features, p.Color, p.Material, p.Upc, p.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [Product] to the database.
func (p *Product) Save(ctx context.Context, db DB) error {
	if p.Exists() {
		return p.Update(ctx, db)
	}
	return p.Insert(ctx, db)
}

// Upsert performs an upsert for [Product].
func (p *Product) Upsert(ctx context.Context, db DB) error {
	switch {
	case p._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.products (` +
		`id, name, description, categories, price, features, color, material, upc` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`name = EXCLUDED.name, description = EXCLUDED.description, categories = EXCLUDED.categories, price = EXCLUDED.price, features = EXCLUDED.features, color = EXCLUDED.color, material = EXCLUDED.material, upc = EXCLUDED.upc `
	// run
	logf(sqlstr, p.ID, p.Name, p.Description, p.Categories, p.Price, p.Features, p.Color, p.Material, p.Upc)
	if _, err := db.ExecContext(ctx, sqlstr, p.ID, p.Name, p.Description, p.Categories, p.Price, p.Features, p.Color, p.Material, p.Upc); err != nil {
		return logerror(err)
	}
	// set exists
	p._exists = true
	return nil
}

// Delete deletes the [Product] from the database.
func (p *Product) Delete(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return nil
	case p._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.products ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, p.ID)
	if _, err := db.ExecContext(ctx, sqlstr, p.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	p._deleted = true
	return nil
}

// ProductByID retrieves a row from 'public.products' as a [Product].
//
// Generated from index 'products_pkey'.
func ProductByID(ctx context.Context, db DB, id uuid.UUID) (*Product, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, description, categories, price, features, color, material, upc ` +
		`FROM public.products ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	p := Product{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&p.ID, &p.Name, &p.Description, &p.Categories, &p.Price, &p.Features, &p.Color, &p.Material, &p.Upc); err != nil {
		return nil, logerror(err)
	}
	return &p, nil
}
