package graphql

import (
	"context"
	"github.com/BlackRule/App-and-its-features-CRUD/common"
	"github.com/BlackRule/App-and-its-features-CRUD/entities/app_feature"
	"github.com/BlackRule/App-and-its-features-CRUD/graphql/gen"
)

func (r *Resolver) Features(ctx context.Context, obj *app_feature.App) ([]app_feature.Feature, error) {
	var features []app_feature.Feature
	result := r.db.Model(obj).Association("Features").Find(&features)
	//Fixme handle Find error!
	return features, result.Error

}

type appResolver struct{ *Resolver }

func (r *appResolver) Features(ctx context.Context, obj *app_feature.App) ([]app_feature.Feature, error) {
	var features []app_feature.Feature
	result1 := r.db.Model(&obj)
	result2 := result1.Association("Features")
	result3 := result2.Find(&features)
	return features, result3.Error
}

func (r *Resolver) App() gen.AppResolver {
	return &appResolver{r}
}

func (r *mutationResolver) CreateApp(ctx context.Context, name string) (*app_feature.App, error) {
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
	_, err := common.Require_auth(ctx)
	if err != nil {
		return nil, err
	}
	app := app_feature.App{
		ID:   uint(id),
		Name: name,
	}
	return &app, r.db.Save(app).Error
}

func (r *mutationResolver) DeleteApp(ctx context.Context, id int) (int, error) {
	_, err := common.Require_auth(ctx)
	if err != nil {
		return -1, err
	}
	app := app_feature.App{
		ID: uint(id),
	}
	return id, r.db.Delete(&app).Error
}

func (r *mutationResolver) AddFeatureToApp(ctx context.Context, featureID int, appID int) (*bool, error) {
	//fixme maybe return bool instead of *bool everywhere?
	var res bool
	_, err := common.Require_auth(ctx)
	if err != nil {
		return &res, err
	}
	var app app_feature.App
	if err := r.db.First(&app, appID).Error; err != nil {
		return &res, err
	}
	var feature app_feature.Feature
	if err := r.db.First(&feature, featureID).Error; err != nil {
		return &res, err
	}
	result := r.db.Model(&app).Association("Features").Append(&feature)
	if result.Error == nil {
		res = true
	}
	return &res, result.Error
}

func (r *mutationResolver) RemoveFeatureFromApp(ctx context.Context, featureID int, appID int) (*bool, error) {
	var res bool
	_, err := common.Require_auth(ctx)
	if err != nil {
		return &res, err
	}
	var app app_feature.App
	if err := r.db.First(&app, appID).Error; err != nil {
		res = false
		return &res, err
	}
	var feature app_feature.Feature
	if err := r.db.First(&feature, featureID).Error; err != nil {
		res = false
		return &res, err
	}
	result := r.db.Model(&app).Association("Features").Delete(feature)
	if result.Error == nil {
		res = true
	}
	return &res, result.Error
}
