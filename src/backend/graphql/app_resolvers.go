package graphql

import (
	"context"
	"github.com/BlackRule/App-and-its-features-CRUD/common"
	"github.com/BlackRule/App-and-its-features-CRUD/entities/app_feature"
)

// FIXME
//func (r *Resolver) App() gen.AppResolver {
//	return nil
//}

func (r *mutationResolver) CreateApp(ctx context.Context, name string) (*app_feature.App, error) {
	//TODO
	_, err := common.Require_auth(ctx)
	if err != nil {
		return nil, err
	}
	app := &app_feature.App{
		Name: name,
	}
	err = r.db.Create(app).Error
	if err != nil {
		return nil, err
	}
	return app, err
}

func (r *queryResolver) Apps(ctx context.Context) ([]app_feature.App, error) {
	var apps []app_feature.App
	result := r.db.Find(&apps)
	return apps, result.Error
}

func (r *mutationResolver) UpdateApp(ctx context.Context, id int, name string) (*app_feature.App, error) {
	//TODO
	_, _ = common.Require_auth(ctx)
	app := app_feature.App{
		ID:   uint(id),
		Name: name,
	}
	return &app, r.db.Save(app).Error
}

func (r *mutationResolver) DeleteApp(ctx context.Context, id int) (int, error) {
	//TODO
	_, _ = common.Require_auth(ctx)
	app := app_feature.App{
		ID: uint(id),
	}
	return id, r.db.Delete(&app).Error
}
